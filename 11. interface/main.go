package main

import "fmt"

type Action interface {
	speak()
	eat()
}

type Dog struct {
	name string
}

func (dog Dog) speak() {
	fmt.Println("Go Go")
}

func (dog Dog) eat() {
	fmt.Println("Yumi Yumi")
}

type Cat struct {
	name string
}

func (cat Cat) speak() {
	fmt.Println("Meow Meow")
}

func (cat Cat) eat() {
	fmt.Println("gumgum")
}

func action(action Action) {
	action.speak()
	action.eat()
}

func main() {
	myDog := Dog{
		name: "Buck",
	}
	myCat := Cat{
		name: "Catty",
	}
	action(myDog)
	action(myCat)
}
