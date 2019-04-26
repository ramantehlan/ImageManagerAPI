package main

import (
	"database/sql"
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

// EnsureAllDirectory is to make sure all the important paths exist
func EnsureAllDirectory(dirs Dirs) {
	for _, dir := range dirs {
		EnsureDirectory(dir.path)
	}
}

// ErrorHandler is to handle errors and log them
func ErrorHandler(err error) {
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			Log("warn", "No rows found in table")
		default:
			Log("error", err.Error())
		}
	}
}
