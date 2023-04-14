package main

import "fmt"

func main() {
	// make an empty map
	ages := make(map[string]int)

	ages["eric"] = 25

	fmt.Println(ages["eric"])

	//
	names := map[string]int{
		"eric": 25,
		"mine": 22,
		"give": 1,
	}

	fmt.Println(names)

	delete(ages, "eric")

	for name, value := range names {
		fmt.Printf("%s \t%d\n", name, value)
	}

	// names := make([]string, 0, len(ages))
}
