package main

import (
	"os"
)

// DirectoryExist is to check if a directory exist
func DirectoryExist(path string) bool {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true
		}
	}
	return false
}

// EnsureDirectory is to create a directory if it don't exist
func EnsureDirectory(path string) {
	if !DirectoryExist(path) {
		os.Mkdir(path, os.ModePerm)
	}
}
