package allegro

import (
	"container/list"
	"github.com/dradtke/go-allegro/allegro/tools"
	"os"
	"sync"
	"testing"
)

func TestCoverage(t *testing.T) {
	packageRoot, err := os.Getwd()
	if err != nil {
		t.Fatal(err.Error())
	}

	missingFuncs := make(chan *tools.MissingFunc)
	errs := make(chan error)
	go tools.ScanHeaders(packageRoot, missingFuncs, errs)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for f := range missingFuncs {
			if f.Module == "" {
				t.Errorf("Missing allegro function '%s'", f.Name)
			} else {
				t.Errorf("Module '%s' missing function '%s'", f.Module, f.Name)
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
