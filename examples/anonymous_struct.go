package examples

import "fmt"

func AnonymousStruct() {
	animal := struct {
		name  string
		speak func() string
	}{
		name: "cat",
		speak: func() string {
			return "meow"
		},
	}
	fmt.Printf("Our animal's name is %s and it says %s", animal.name, animal.speak())
}
