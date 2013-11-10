package tools

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type mod struct {
	name  string
	regex *regexp.Regexp
}

// regexes for various function macros
var alFunc = buildRegex("AL_FUNC")

var modules = []mod{
	mod{name: "image", regex: buildRegex("ALLEGRO_IIO_FUNC")},
	//mod{name: "audio", regex: buildRegex("ALLEGRO_KCM_AUDIO_FUNC")},
	//mod{name: "ttf", regex: buildRegex("ALLEGRO_TTF_FUNC")},
	// TODO: add whatever other modules need to be included
}

func buildRegex(macro string) *regexp.Regexp {
	return regexp.MustCompile(macro + `\((?P<type>.*), (?P<name>.*), \((?P<params>.*)\)\)`)
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

type MissingFunc struct {
	Name   string
	Header string
	Module string
}

// ScanHeaders() scans all header files found in /usr/include/allegro5, collects a list
// of AL_FUNC's, and greps the Go code found in packageRoot for each one, writing a
// newline-separated list of those that aren't found to out.
func ScanHeaders(packageRoot string, missingFuncs chan *MissingFunc, errs chan error) {
	var (
		source []byte
		sourceErr error
		headerRoot = filepath.Join("/", "usr", "include", "allegro5")
	)

	defer func() {
		close(missingFuncs)
		close(errs)
	}()

	// first walk the full root, looking for standard allegro functions
	source, sourceErr = getSource(packageRoot)
	if sourceErr != nil {
		errs <- sourceErr
		return
	}
	filepath.Walk(headerRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == "internal" {
			return filepath.SkipDir
		} else if info.IsDir() || !strings.HasSuffix(info.Name(), ".h") {
			return nil
		}
		data, err2 := ioutil.ReadFile(path)
		if err2 != nil {
			errs <- err2
			return nil
		}
		for _, line := range strings.Split(string(data), "\n") {
			vals := alFunc.FindStringSubmatch(line)
			if vals == nil {
				// no match
				continue
			}
			name := strings.TrimSpace(vals[2])
			if strings.HasPrefix(name, "_") {
				// function names starting with an underscore are private
				continue
			}
			if !bytes.Contains(source, []byte("C."+name)) {
				missingFuncs <- &MissingFunc{Name: name, Module: "", Header: path}
			}
		}
		return nil
	})

	// now iterate through all known modules
	for _, m := range modules {
		var (
			header = filepath.Join(headerRoot, "allegro_"+m.name+".h")
			root = filepath.Join(packageRoot, m.name)
		)
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
		source, sourceErr = getSource(filepath.Join(packageRoot, m.name))
		if sourceErr != nil {
			errs <- sourceErr
			return
		}
		for _, line := range strings.Split(string(data), "\n") {
			vals := m.regex.FindStringSubmatch(line)
			if vals == nil {
				// no match
				continue
			}
			name := strings.TrimSpace(vals[2])
			if strings.HasPrefix(name, "_") {
				// function names starting with an underscore are private
				continue
			}
			if !bytes.Contains(source, []byte("C."+name)) {
				missingFuncs <- &MissingFunc{Name: name, Module: m.name, Header: header}
			}
		}
	}
}
