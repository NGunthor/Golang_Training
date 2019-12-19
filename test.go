package main

import "fmt"


type Inter interface {}


type elementA struct {
}

func (e *elementA) Make() {
	fmt.Println("make make make")
}


func main() {
	fmt.Println("Test")

	var i int = 5

}