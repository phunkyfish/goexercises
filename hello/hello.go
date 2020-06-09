package main

import "fmt"

func main() {
	fmt.Printf("hello world\n")

	fmt.Printf("Total: %d\n", add(2, 3))
	z, _ := add2(2, 3)
	fmt.Printf("Total: %d\n", z)

	a, e := add2(-1, 3)
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
	} else {
		fmt.Printf("Total: %d\n", a)
	}
}

func add(x int, y int) int {
	return x + y
}

func add2(x int, y int) (int, error) {
	if x < 0 {
		return 0, fmt.Errorf("Negative value")
	}

	return x + y, nil
}
