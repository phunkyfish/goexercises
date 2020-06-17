package main

import "fmt"

func main() {
	// a := []byte{65, 66, 67}
	// s := string(a)
	// fmt.Printf("%s\n", s)

	s := "to ✂️ Copy and 📋 Paste 👌"
	b := []byte(s)
	r := []rune(s)

	fmt.Printf("%#v\n", s)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", r)
}
