package main

type Person struct {
	name string
	age  int
}

func main() {
	// person := Person{name: "a", age: 1}

	// changeName(&person)

	// fmt.Printf("%v", person)

	createMing()
}

func changeName(person *Person) {
	person.name = person.name + "b"
}
