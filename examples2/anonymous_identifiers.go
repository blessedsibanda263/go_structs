package examples2

import "log"

type Animal struct {
	string
}

func (a Animal) speak() {
	log.Println(a.string)
}

func AnonymousIdentifiers() {
	a := Animal{"cat"}
	a.speak()

	a.string = "dog"
	a.speak()
}
