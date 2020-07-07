package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fileName := os.Args[1]

	if fileName != "" {
		f, err := os.Open(fileName)

		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		} else {
			io.Copy(os.Stdout, f)
		}

	}

}
