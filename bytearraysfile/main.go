package main

import (
	"fmt"
	"os"
)

func main() {
	r, err := os.Open("main.go")
	if err != nil {
		panic(err.Error())
	}
	b := make([]byte, 1024)
	n, err := r.Read(b)
	fmt.Printf("%d, error: %v - %v", n, err, string(b))

	// Homework create writer
}
