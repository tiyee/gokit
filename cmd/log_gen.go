package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
	}
}
func run() error {
	pkg := "go.uber.org/zap"
	files := []string{"logger.go", "field.go"}
	p, err := build.Import(pkg, "", build.FindOnly)
	if err != nil {
		return err
	}
	filter := func(info os.FileInfo) bool {
		for _, file := range files {
			if strings.HasSuffix(info.Name(), file) {
				return true
			}
		}
		return false
	}

	if m, err := parser.ParseDir(token.NewFileSet(), p.Dir, filter, parser.Mode(0)); err == nil {
		for _, pack := range m {
			//ast.Print(token.NewFileSet(), pack)
			for name, f := range pack.Files {
				fmt.Printf("filename=%s, packageg_name=%s\n", name, f.Name.Name)
				decls := make([]ast.Decl, 0)
				for _, dec := range f.Decls {
					v := &vs{fileName: name, packageName: f.Name.Name}
					ast.Walk(v, dec)
					for _, v := range v.funcs {
						fmt.Println("func:", v.Name.String())
						if strings.HasSuffix(name, "field.go") {
							decls = append(decls, generateFieldFunc(v))
						} else {
							decls = append(decls, generateLoggerFunc(v))
						}

					}

					//if err := printer.Fprint(os.Stdout, token.NewFileSet(), dec); err != nil {
					//	fmt.Println(err.Error())
					//}
				}
				fmt.Println("funcs num: ", len(decls))
				if pwd, err := os.Getwd(); err == nil {
					var filename string
					var imps []string
					if strings.HasSuffix(name, "field.go") {
						imps = []string{"fmt", "time", "go.uber.org/zap/zapcore", "go.uber.org/zap"}
						filename = fmt.Sprintf("%s/pkg/component/log/field.go", pwd)
					} else {
						imps = []string{"go.uber.org/zap/zapcore"}
						filename = fmt.Sprintf("%s/pkg/component/log/logger.go", pwd)
					}

					if fs, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err == nil {
						if err := writeGoFile(fs, decls, imps); err == nil {
							fs.WriteString("// End\n")

						} else {
							fmt.Println(err.Error())
						}
						fs.Close()
					}
				}
			}

		}

	} else {
		return err
	}
	return nil
}

type vs struct {
	funcs       []*ast.FuncDecl
	fileName    string
	packageName string
}

func (v *vs) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.FuncDecl:
		remain := false
		if !n.Name.IsExported() {
			return nil
		}
		if n.Recv != nil && n.Recv.NumFields() > 0 {
			if len(n.Recv.List[0].Names) > 0 {
				for _, field := range n.Recv.List {
					if t, ok := field.Type.(*ast.StarExpr); ok {
						// receiver is pointer type
						fmt.Print("point receiver type: ", t.X.(*ast.Ident).String())
						//ast.Print(token.NewFileSet(), field)
						if t.X.(*ast.Ident).String() == "Logger" {
							remain = true
						}
					}
					if t, ok := field.Type.(*ast.Ident); ok {
						fmt.Print("func receiver type: ", t.Name, remain)
					}

				}

			}

		}
		if n.Recv == nil && n.Name.IsExported() {
			//ast.Print(token.NewFileSet(), n)

			for _, field := range n.Type.Results.List {

				ident, ok := field.Type.(*ast.Ident)
				if ok && ident.Name == "Field" {
					remain = true
					//fmt.Print("pure func: ", n.Name.Name)
				}

			}
		}
		if remain {
			v.funcs = append(v.funcs, n)
		} else {
			return nil
		}

	}
	return v
}

func generateLoggerFunc(fn *ast.FuncDecl) *ast.FuncDecl {
	fn.Recv = nil

	fnName := fn.Name.String()
	var args []string
	for _, field := range fn.Type.Params.List {
		for _, id := range field.Names {
			idStr := id.String()
			_, ok := field.Type.(*ast.Ellipsis)
			if ok {
				// Ellipsis args
				idStr += "..."
			}
			args = append(args, idStr)
		}
	}

	exprStr := fmt.Sprintf(`logger.%s(%s)`, fnName, strings.Join(args, ","))
	expr, err := parser.ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}

	var body []ast.Stmt
	if fn.Type.Results != nil {
		body = []ast.Stmt{
			&ast.ReturnStmt{
				// Return:
				Results: []ast.Expr{expr},
			},
		}
	} else {
		body = []ast.Stmt{
			&ast.ExprStmt{
				X: expr,
			},
		}
	}

	fn.Body.List = body

	return fn
}
func generateFieldFunc(fn *ast.FuncDecl) *ast.FuncDecl {
	fn.Recv = nil

	fnName := fn.Name.String()
	// handle return: *Logger -> *zap.Logger
	//if fn.Type.Results != nil {
	//	for _, result := range fn.Type.Results.List {
	//		if ident, ok := result.Type.(*ast.Ident); ok {
	//			if strings.Contains(ident.Name, "Field") {
	//				exprStr := fmt.Sprintf(`zapcore.%s`, ident.Name)
	//				expr, err := parser.ParseExpr(exprStr)
	//				if err != nil {
	//					panic(err)
	//				}
	//				result.Type = expr
	//			}
	//		}
	//	}
	//}

	// handle body: _globalS.Info(args...)
	var args []string
	for _, field := range fn.Type.Params.List {
		for _, id := range field.Names {
			idStr := id.String()
			_, ok := field.Type.(*ast.Ellipsis)
			if ok {
				// Ellipsis args
				idStr += "..."
			}
			args = append(args, idStr)
		}
	}

	exprStr := fmt.Sprintf(`zap.%s(%s)`, fnName, strings.Join(args, ","))
	expr, err := parser.ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}

	var body []ast.Stmt
	if fn.Type.Results != nil {
		body = []ast.Stmt{
			&ast.ReturnStmt{
				// Return:
				Results: []ast.Expr{expr},
			},
		}
	} else {
		body = []ast.Stmt{
			&ast.ExprStmt{
				X: expr,
			},
		}
	}

	fn.Body.List = body

	return fn
}

// Output Go code
func writeGoFile(wr io.Writer, funcs []ast.Decl, imps []string) error {
	f := &ast.File{
		Name: &ast.Ident{
			Name: "log",
		},

		Decls: []ast.Decl{

			//&ast.GenDecl{
			//
			//	Tok: token.IMPORT,
			//	Specs: []ast.Spec{
			//		&ast.ImportSpec{
			//			Path: &ast.BasicLit{
			//				Kind:  token.STRING,
			//				Value: "\"go.uber.org/zap\"",
			//			},
			//		},
			//	},
			//},

		},
	}
	for _, imp := range imps {
		f.Decls = append(f.Decls, &ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{
				&ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: fmt.Sprintf("\"%s\"", imp),
					},
				},
			},
		})
	}
	f.Decls = append(f.Decls, funcs...)
	bs := new(bytes.Buffer)
	cm := "// Code generated by log-gen. DO NOT EDIT.\n"
	bs.WriteString(cm)
	f.Package = token.Pos(len(cm))
	if err := format.Node(bs, token.NewFileSet(), f); err == nil {
		if _, err := wr.Write(bs.Bytes()); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		fmt.Println(err.Error())
		return err
	}

}
