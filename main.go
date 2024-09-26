package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Usage: rpl 'old_string' 'new_string' [file_or_directory]")
		os.Exit(1)
	}

	oldString := os.Args[1]
	newString := os.Args[2]

	var startPath string
	if len(os.Args) == 4 {
		startPath = os.Args[3]
	} else {
		startPath = "."
	}

	fileInfo, err := os.Stat(startPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var modifiedCount int
	if fileInfo.IsDir() {
		err = filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				modified, err := replaceInFile(path, oldString, newString)
				if err != nil {
					return err
				}
				if modified {
					fmt.Println("Modified:", path)
					modifiedCount++
				}
			}
			return nil
		})

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	} else {
		// Jeśli to plik, dokonaj zamiany tylko w tym pliku
		modified, err := replaceInFile(startPath, oldString, newString)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if modified {
			fmt.Println("Modified:", startPath)
			modifiedCount++
		} else {
			fmt.Println("No modifications made to:", startPath)
		}
	}

	fmt.Println("Total files modified:", modifiedCount)
}

func replaceInFile(filePath, oldString, newString string) (bool, error) {
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	if isBinary(input) {
		return false, nil
	}

	content := string(input)
	if !strings.Contains(content, oldString) {
		return false, nil
	}

	output := strings.ReplaceAll(content, oldString, newString)
	err = ioutil.WriteFile(filePath, []byte(output), 0666)
	if err != nil {
		return false, err
	}

	return true, nil
}

func isBinary(data []byte) bool {
	return bytes.Contains(data, []byte{0})
}
