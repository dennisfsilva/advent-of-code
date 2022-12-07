package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("\nFailed to open file: %s, error: %v", filePath, err)
		return []byte{}
	}
	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("\nFailed to read file, error: %v", err)
		return []byte{}
	}

	return fileData
}
