package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/structtag"
	"github.com/spf13/cobra"
)

var addTagCmd = &cobra.Command{
	Use:   "add-tag [directory]",
	Short: "Add tags to struct members in .pb.go files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(path, ".pb.go") {
				err := processFile(path)
				if err != nil {
					fmt.Println("Error processing file:", err)
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println("Error walking directory:", err)
		}
	},
}

func processFile(filePath string) error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	ast.Inspect(file, func(node ast.Node) bool {
		if typeSpec, ok := node.(*ast.TypeSpec); ok {
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				if !containsRequiredFields(structType.Fields) {
					return true
				}
				for _, field := range structType.Fields.List {
					if field.Tag == nil {
						continue
					}
					fmt.Println("Processing struct:", typeSpec.Name.Name, field.Names[0].Name)
					tags, err := structtag.Parse(field.Tag.Value[1 : len(field.Tag.Value)-1])
					if err != nil {
						log.Panic(err)
					}
					jsonTag, err := tags.Get("json")
					if err != nil {
						log.Panic(err)
					}
					keys := []string{
						"json", "form", "query", "params", "xml",
					}
					for _, key := range keys {
						_ = tags.Set(&structtag.Tag{
							Key:  key,
							Name: jsonTag.Name,
						})
					}
					field.Tag = &ast.BasicLit{Value: "`" + tags.String() + "`"}
				}
			}
		}
		return true
	})

	fileWriter, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fileWriter.Close()

	err = format.Node(fileWriter, fileSet, file)
	if err != nil {
		return err
	}

	return nil
}

func containsRequiredFields(fields *ast.FieldList) bool {
	requiredFields := map[string]struct{}{
		"state":         {},
		"sizeCache":     {},
		"unknownFields": {},
	}
	for _, field := range fields.List {
		for _, name := range field.Names {
			if _, ok := requiredFields[name.Name]; ok {
				return true
			}
		}
	}
	return false
}
