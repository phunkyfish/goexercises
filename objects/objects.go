package main

import "fmt"

// Dog s;lkdjf;kals
type Dog struct {
	name string
}

func main() {
	d := Dog{
		"Ben"}
	e := &d
	e.name = "Jerry"
	fmt.Printf("Dog name: %s\n", d.name)
	//fmt.Printf("Noise: %s\n", d.noise())
	//d.noise()
	react(&d)
}

func (d *Dog) noise() string {
	return "Barks"
}

// Pet s;lkdjf;kals
type Pet interface {
	noise() string
}

func react(p Pet) {
	fmt.Printf("Noise: %s", p.noise())
}
