package examples

import "fmt"

type Animal struct {
	name string
}

func (a Animal) speak() string {
	switch a.name {
	case "cat":
		return "meow"
	case "dog":
		return "woof"
	default:
		return "nondescript animal noise?"
	}
}

func StructMethods() {
	a := Animal{
		name: "cat",
	}
	fmt.Println(a.speak())

	a.name = "dog"
	fmt.Println(a.speak())

	a.name = "llama"
	fmt.Println(a.speak())
}
