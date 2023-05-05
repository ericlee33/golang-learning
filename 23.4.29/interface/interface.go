package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/no/file")

	if e, ok := err.(*os.PathError); ok {
		fmt.Println(e)
	}

	fmt.Println(f)
}
