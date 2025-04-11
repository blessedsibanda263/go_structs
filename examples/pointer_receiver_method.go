package examples

import "fmt"

type character struct {
	name string
}

func (ch *character) fixName() {
	ch.name = "Blessed Sibanda"
}

func PointerReceiverMethod() {
	ch := new(character)
	ch.name = "Sheldon Cooper"
	fmt.Println("my name is", ch.name)
	ch.fixName()
	fmt.Println("Just kidding, my name is", ch.name)
}
