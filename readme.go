//go:build readme
// +build readme

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	// Get Template filename
	t, err := template.ParseFiles(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get Target filename
	f, err := os.Create(os.Args[2])
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	chartsDir := os.Args[3]

	files, _ := filepath.Glob("**/*/Chart.yaml")
	var charts []map[string]string
	for _, file := range files {
		charts = append(charts, map[string]string{
			"name":    strings.Split(file, "/")[1],
			"dir":     chartsDir,
			"version": extractVersion(file),
		})
	}
	err = t.Execute(f, map[string]interface{}{
		"charts":       charts,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func extractVersion(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "version: ") {
			version := strings.TrimPrefix(line, "version: ")
			return version
		}
	}
	return ""
}
