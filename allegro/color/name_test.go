// This test iterates through all of the defined Name constants
// in names.go and reports an error for each one unrecognized
// by Allegro's al_color_name_to_rgb() function.
package color

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)

func TestNames(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "names.go", nil, parser.Mode(0))
	if err != nil {
		t.Fatal(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		var (
			decl *ast.GenDecl
			ok   bool
		)
		if decl, ok = n.(*ast.GenDecl); !ok {
			return true
		}
		if decl.Tok != token.CONST {
			return false
		}
		for _, spec := range decl.Specs {
			v := spec.(*ast.ValueSpec)
			if v.Type.(*ast.Ident).Name == "Name" {
				for _, lit := range v.Values {
					n, err := strconv.Unquote(lit.(*ast.BasicLit).Value)
					if err != nil {
						t.Error(err)
					}
					_, _, _, err = NameToRgb(Name(n))
					if err != nil {
						t.Error(err)
					}
				}
			}
		}
		return false
	})
}
