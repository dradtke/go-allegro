package tools

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var alFunc = regexp.MustCompile(`AL_FUNC\((?P<type>.*), (?P<name>.*), \((?P<params>.*)\)\)`)

func ScanHeaders(packageRoot string, out io.Writer) {
	var sourceBuf bytes.Buffer
	headerRoot := "/usr/include/allegro5"

	filepath.Walk(packageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}
		data, err2 := ioutil.ReadFile(path)
		if err2 != nil {
			return fmt.Errorf("can't read Go source file \"%s\": %s", path, err2.Error())
		}
		sourceBuf.Write(data)
		return nil
	})

	source := sourceBuf.Bytes()
	var buf bytes.Buffer

	filepath.Walk(headerRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".h") {
			return nil
		}
		data, err2 := exec.Command("clang", "-E", path).Output()
		if err2 != nil {
			// failed to run preprocessor
			//return fmt.Errorf("failed to run preprocessor on \"%s\": %s", path, err2.Error())
			return nil
		}
		for _, line := range strings.Split(string(data), "\n") {
			fmt.Fprintln(os.Stderr, line)
			vals := alFunc.FindStringSubmatch(line)
			if vals == nil {
				// no match
				continue
			}
			name := vals[2]
			if strings.HasPrefix(name, "_") {
				// function names starting with an underscore are private
				continue
			}
			if !bytes.Contains(source, []byte("C."+name)) {
				buf.WriteString(name + "\n")
			}
		}
		return nil
	})

	out.Write(buf.Bytes())
}
