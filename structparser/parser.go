package structparser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ParseFolder(inputFolder string, outputFolder string) error {

	var goFiles []string
	err := filepath.Walk(inputFolder, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {

			goFiles = append(goFiles, path)
		}
		return nil
	})
	if err != nil {

		return err
	}

	for _, goFile := range goFiles {

		parsedStructs, err := parseFile(goFile)
		if err != nil {

			return fmt.Errorf("Error parsing file: %v", err)
		}

		for _, s := range parsedStructs {

			mdContent := generateMarkdown([]structInfo{s})

			relPath, err := filepath.Rel(inputFolder, goFile)
			if err != nil {

				return fmt.Errorf("Error determining relative path: %v", err)
			}

			outputFile := filepath.Join(outputFolder, filepath.Dir(relPath), fmt.Sprintf("%s.md", s.Name))

			err = os.MkdirAll(filepath.Dir(outputFile), os.ModePerm)
			if err != nil {

				return fmt.Errorf("Error creating output directory: %v", err)
			}

			err = writeMarkdownToFile(outputFile, mdContent)
			if err != nil {

				return fmt.Errorf("Error writing markdown file: %v", err)
			}

			fmt.Printf("Markdown file generated: %s\n", outputFile)
		}
	}

	return nil
}
