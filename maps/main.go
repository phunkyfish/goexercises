package main

import "fmt"

func main() {
	a := map[string]string{}

	a["apple"] = "green"
	a["orange"] = "orange"
	a["pear"] = "green"

	for key, value := range a {
		fmt.Printf("%s : %s\n", key, value)
	}

	v, found := a["carrot"]
	fmt.Printf("%s, %v\n", v, found)
}
