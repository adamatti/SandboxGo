// FROM https://goplay.space/#i1RsX1k8n8o
package main

import "fmt"

type A struct{}

type MockA struct {
	A
	anotherField string
}

func (me A) Procedure() {
	me.SecondProcedure()
}

func (me A) SecondProcedure() {
	fmt.Println("Original implementation triggered.")
}

/**
// If this was commented-out, it would use ALL methods from the imbeded object.
func (me MockA) Procedure() {
	me.SecondProcedure()
}
*/

func (me MockA) SecondProcedure() {
	fmt.Println("Mocked implementation triggered.")
}

func main() {
	obj := MockA{}

	// I want this to return the Mocked implementation...
	obj.Procedure()
}
