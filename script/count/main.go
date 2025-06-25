package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "."
	dirEntries, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}

	problemsDir := []os.DirEntry{}
	for _, dir := range dirEntries {
		if dir.IsDir() && dir.Name() == "problems" || dir.Name() == "extra" || dir.Name() == "top-75" {
			problemsDir = append(problemsDir, dir)
		}
	}

	totalProblems := 0
	for _, dir := range problemsDir {
		subDirPath := filepath.Join(root, dir.Name())

		subDirProblems, err := os.ReadDir(subDirPath)
		if err != nil {
			panic(err)
		}

		for range subDirProblems {
			totalProblems++
		}
	}

	fmt.Printf("I've done %d leetcodes.\n", totalProblems)
}
