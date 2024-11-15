// purpose parse golang struct files to markdown

package structparser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"reflect"
	"strings"
)

// parseFile parses a Go source file and returns a list of structs with field info
func parseFile(filename string) ([]structInfo, error) {

	var structs []structInfo

	// Parse the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {

		return nil, err
	}

	// Walk through the AST to find structs
	for _, decl := range node.Decls {

		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {

			continue
		}

		for _, spec := range genDecl.Specs {

			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {

				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {

				continue
			}

			structComment := ""
			if genDecl.Doc != nil {

				structComment = genDecl.Doc.Text()
			}

			var fields []fieldInfo
			for _, field := range structType.Fields.List {

				fieldType := exprToString(field.Type)
				fieldName := ""
				if len(field.Names) > 0 {

					fieldName = field.Names[0].Name
				}

				// Extract the comment from the field
				comment := ""
				if field.Doc != nil {

					comment = field.Doc.Text()
				}

				jsonTag := ""
				if field.Tag != nil {

					tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
					jsonTag = tag.Get("json")
				}

				fields = append(fields, fieldInfo{
					Name:    fieldName,
					Type:    fieldType,
					JSONTag: jsonTag,
					Comment: comment,
				})
			}

			structs = append(structs, structInfo{
				Name:    typeSpec.Name.Name,
				Comment: structComment,
				Fields:  fields,
			})
		}
	}

	return structs, nil
}

// exprToString converts an ast.Expr to a string representation (type name)
func exprToString(expr ast.Expr) string {

	switch v := expr.(type) {

	case *ast.Ident:
		return v.Name

	case *ast.ArrayType:

		return "[]" + exprToString(v.Elt)

	case *ast.MapType:

		return "map[" + exprToString(v.Key) + "]" + exprToString(v.Value)

	case *ast.ChanType:
		return "chan " + exprToString(v.Value)

	default:
		var buf bytes.Buffer
		ast.Fprint(&buf, token.NewFileSet(), expr, nil)
		return buf.String()
	}
}

// generateMarkdown generates a Markdown string for a list of StructInfo
func generateMarkdown(structs []structInfo) string {

	var md strings.Builder

	for _, s := range structs {

		if s.Name != "" {

			md.WriteString(fmt.Sprintf("# %s\n\n", s.Name))
		}

		if s.Comment != "" {

			md.WriteString(fmt.Sprintf("%s\n\n", s.Comment))
		}

		if len(s.Fields) > 0 {

			md.WriteString("| Field | Type | Json | Description |\n")
			md.WriteString("|-------|------|------|-------------|\n")
			for _, field := range s.Fields {

				md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
					field.Name, field.Type, field.JSONTag, strings.TrimSpace(field.Comment)))
			}
			md.WriteString("\n")
		}
	}

	return md.String()
}

// writeMarkdownToFile writes the generated Markdown content to a .md file
func writeMarkdownToFile(filename, content string) error {

	return ioutil.WriteFile(filename, []byte(content), 0644)
}
