package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[1:]
	c := b[1:2]
	//d := []int{1, 2, 3, 4}

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)

	//a[2] = 500
	//c = append(c, 600)
	b = append(b, 600)
	c = append(c, 600)

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v - %d - %d\n", b, len(b), cap(b))
	fmt.Printf("%#v\n", c)
}
