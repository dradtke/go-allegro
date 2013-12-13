package main

import (
	"bytes"
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"testing"
)

var modules = []mod{
	mod{name: "acodec", decl: buildRegex("ALLEGRO_ACODEC_FUNC")},
	//mod{name: "audio", decl: buildRegex("ALLEGRO_KCM_AUDIO_FUNC")},
	mod{name: "color", decl: buildRegex("ALLEGRO_COLOR_FUNC")},
	mod{name: "dialog", decl: buildRegex("ALLEGRO_DIALOG_FUNC"), header: "native_dialog"},
	mod{name: "font", decl: buildRegex("ALLEGRO_FONT_FUNC")},
	mod{name: "image", decl: buildRegex("ALLEGRO_IIO_FUNC")},
	mod{name: "memfile", decl: buildRegex("ALLEGRO_MEMFILE_FUNC")},
	mod{name: "physfs", decl: buildRegex("ALLEGRO_PHYSFS_FUNC")},
	mod{name: "ttf", decl: buildRegex("ALLEGRO_TTF_FUNC"), path: "font/ttf"},
}

type mod struct {
	name, path, header string
	decl               *decl
}

func (m *mod) Header() string {
	if m.header != "" {
		return "allegro_" + m.header + ".h"
	} else {
		return "allegro_" + m.name + ".h"
	}
}

func (m *mod) Path() string {
	if m.path != "" {
		return m.path
	} else {
		return m.name
	}
}

type decl struct {
	macro string
	regex *regexp.Regexp
}

// regexes for various function macros
var alFunc = buildRegex("AL_FUNC")

func buildRegex(macro string) *decl {
	return &decl{macro, regexp.MustCompile(macro + `\((?P<type>.*), (?P<name>.*), \((?P<params>.*)\)\)`)}
}

func getSource(packageRoot string) ([]byte, error) {
	var buf bytes.Buffer
	err := filepath.Walk(packageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != packageRoot {
			return filepath.SkipDir
		} else if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}
		data, err2 := ioutil.ReadFile(path)
		if err2 != nil {
			return fmt.Errorf("can't read Go source file \"%s\": %s", path, err2.Error())
		}
		buf.Write(data)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return buf.Bytes(), err
}

type missingFunc struct {
	Name   string
	Type   string
	Params string
	Header string
	Module string
}

func scanHeaders(packageRoot string, missingFuncs chan *missingFunc, errs chan error) {
	var (
		source     []byte
		sourceErr  error
		headerRoot = filepath.Join("/", "usr", "include", "allegro5")
	)

	defer func() {
		close(missingFuncs)
		close(errs)
	}()

	// First walk the full root, looking for standard allegro functions.
	source, sourceErr = getSource(packageRoot)
	if sourceErr != nil {
		errs <- sourceErr
		return
	}
	filepath.Walk(headerRoot, func(header string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == "internal" {
			return filepath.SkipDir
		} else if info.IsDir() || !strings.HasSuffix(info.Name(), ".h") {
			return nil
		}
		data, err2 := ioutil.ReadFile(header)
		if err2 != nil {
			errs <- err2
			return nil
		}
		findMissingFuncs(data, source, header, alFunc, "", missingFuncs)
		return nil
	})

	// Now iterate through all known modules.
	for _, m := range modules {
		root := filepath.Join(packageRoot, m.Path())
		header := filepath.Join(headerRoot, m.Header())
		if _, err := os.Stat(header); os.IsNotExist(err) {
			errs <- fmt.Errorf("Module header not found at '%s'", header)
			continue
		}
		if info, err := os.Stat(root); os.IsNotExist(err) || !info.IsDir() {
			errs <- fmt.Errorf("Source not found at '%s'", root)
			continue
		}
		data, err := ioutil.ReadFile(header)
		if err != nil {
			errs <- err
			continue
		}
		source, sourceErr = getSource(root)
		if sourceErr != nil {
			errs <- sourceErr
			return
		}
		findMissingFuncs(data, source, header, m.decl, m.name, missingFuncs)
	}
}

// findMissingFuncs() is a customized iteration method used to ensure that multi-line function declarations are found.
func findMissingFuncs(data, source []byte, header string, d *decl, modName string, missingFuncs chan *missingFunc) {
	ch := make(chan string)
	go func() {
		defer func() {
			close(ch)
			if r := recover(); r != nil {
				fmt.Fprintf(os.Stderr, "Header file '%s' is malformed.\n", header)
				os.Exit(1)
			}
		}()
		var buf bytes.Buffer
		lines := strings.Split(string(data), "\n")
		for i := 0; i < len(lines); i++ {
			line := strings.TrimSpace(lines[i])
			buf.WriteString(line)
			if strings.HasPrefix(line, d.macro) {
				for !strings.HasSuffix(line, ";") {
					i++
					line = strings.TrimSpace(lines[i])
					buf.WriteString(line)
				}
			}
			ch <- buf.String()
			buf.Reset()
		}
	}()
	for line := range ch {
		vals := d.regex.FindStringSubmatch(line)
		if vals == nil {
			// no match
			continue
		}
		name := strings.TrimSpace(vals[2])
		if strings.HasPrefix(name, "_") {
			// function names starting with an underscore are private
			continue
		}
		typ := strings.TrimSpace(vals[1])
		params := strings.TrimSpace(vals[3])
		if !bytes.Contains(source, []byte("C."+name)) {
			missingFuncs <- &missingFunc{Name: name, Type: typ, Params: params, Module: modName, Header: header}
		}
	}
}

// Ignore functions that I either don't know how to implement,
// are redundant, or are based on Allegro features that Go
// already provides, like file I/O and fixed point math.
func getBlacklist() map[string]bool {
	blacklist := make(map[string]bool)
	data, err := ioutil.ReadFile("blacklist")
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		i := strings.IndexRune(line, '#')
		if i != -1 {
			line = line[0:i]
		}
		blacklist[strings.TrimSpace(line)] = true
	}
	return blacklist
}

func TestCoverage(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err.Error())
	}
	packageRoot := filepath.Join(cwd, "allegro")

	blacklist := getBlacklist()

	missingFuncs := make(chan *missingFunc)
	errs := make(chan error)
	go scanHeaders(packageRoot, missingFuncs, errs)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for f := range missingFuncs {
			if blacklist[f.Name] {
				continue
			}
			if f.Module == "" {
				t.Errorf("Missing allegro function '%s' in file '%s'", f.Name, f.Header)
			} else {
				t.Errorf("Module '%s' missing function '%s' [%s %s(%s)]", f.Module, f.Name, f.Type, f.Name, f.Params)
			}
		}
		wg.Done()
	}()

	go func() {
		errorList := list.New()
		for err := range errs {
			errorList.PushBack(err)
		}
		for e := errorList.Front(); e != nil; e = e.Next() {
			t.Error("Error: " + e.Value.(error).Error())
		}
		wg.Done()
	}()

	wg.Wait()
}
