package command

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"runtime"
	"testing"
)

func TestSubprocessesUseNewContext(t *testing.T) {
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate test source")
	}
	internalDir := filepath.Dir(filepath.Dir(testFile))

	tests := []struct {
		name      string
		path      string
		wantCalls int
	}{
		{name: "project detection", path: filepath.Join("project", "detect.go"), wantCalls: 2},
		{name: "Claude CLI", path: filepath.Join("llm", "claude.go"), wantCalls: 1},
		{name: "setup commands", path: filepath.Join("setup", "setup.go"), wantCalls: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := parser.ParseFile(token.NewFileSet(), filepath.Join(internalDir, tt.path), nil, 0)
			if err != nil {
				t.Fatalf("parse %s: %v", tt.path, err)
			}

			calls := 0
			ast.Inspect(file, func(node ast.Node) bool {
				selector, ok := node.(*ast.SelectorExpr)
				if !ok || selector.Sel.Name != "NewContext" {
					return true
				}
				identifier, ok := selector.X.(*ast.Ident)
				if ok && identifier.Name == "command" {
					calls++
				}
				return true
			})
			if calls != tt.wantCalls {
				t.Fatalf("%s uses command.NewContext %d times; want %d", tt.path, calls, tt.wantCalls)
			}
		})
	}
}
