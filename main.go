package main

import (
	"fmt"
	"go_structs/examples"
	"go_structs/examples2"
	"go_structs/examples2/tags"
)

func main() {
	examples.AnonymousStruct()
	fmt.Println()
	examples.StructMethods()
	fmt.Println()
	examples.PointerReceiverMethod()
	fmt.Println()
	examples2.AnonymousIdentifiers()
	fmt.Println()
	tags.FormattingDataAsJson()
}
