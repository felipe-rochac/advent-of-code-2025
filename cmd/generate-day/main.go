package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dayNum := flag.Int("day", 0, "Day number to generate")
	flag.Parse()

	if *dayNum <= 0 || *dayNum > 25 {
		fmt.Println("Error: day must be between 1 and 25")
		os.Exit(1)
	}

	dayFolder := fmt.Sprintf("day%d", *dayNum)
	if err := generateDay(dayFolder, *dayNum); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %s\n", dayFolder)
}

func generateDay(dayFolder string, dayNum int) error {
	// Create day directory
	if err := os.MkdirAll(dayFolder, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Files to copy from template
	files := []string{"puzzle.go", "puzzle_test.go", "puzzle.md", "input.txt", "test.txt"}

	templateDir := "template"

	for _, file := range files {
		srcPath := filepath.Join(templateDir, file)
		destPath := filepath.Join(dayFolder, file)

		// Read template file
		content, err := os.ReadFile(srcPath)
		if err != nil {
			return fmt.Errorf("failed to read template file %s: %w", file, err)
		}

		// Replace package name in Go files
		if strings.HasSuffix(file, ".go") {
			contentStr := string(content)
			contentStr = strings.ReplaceAll(contentStr, "package template", fmt.Sprintf("package day%d", dayNum))
			content = []byte(contentStr)
		}

		// Write to destination
		if err := os.WriteFile(destPath, content, 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", file, err)
		}
	}

	return nil
}
