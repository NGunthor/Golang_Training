package main

import (
	"fmt"
)

type Inter interface {
	Make()
}

type elementB struct {
	surname string
}

type elementA struct {
	name string
	elementB
}

func (a *elementA) GetName() {
	fmt.Println("My name is: ", a.name)
}

func (a *elementA) GetSurname() {
	fmt.Println("My surname is: ", a.surname)
}

func (b *elementB) Lals() {
	fmt.Println("LALS!")
}

func (e *elementA) Make() {
	fmt.Println("make make make")
}

func NewPerson() elementA {
	return elementA{
		name:     "John",
		elementB: elementB{"Govard"},
	}
}

func main() {
	fmt.Println("Test")

	nel := NewPerson()

	nel.GetName()
	nel.GetSurname()
	nel.Lals()

}
