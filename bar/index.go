package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Filename is the __filename equivalent
func Filename() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

func main() {
	filename, _ := Filename()

	path := filepath(filename, "../day1/dup2.go")

	content, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf(string(content))
	}
}
