package main

import "fmt"

type Animal struct {
	color string
}

func (animal Animal) getColor() string {
	return animal.color
}

type Dog struct {
	Animal
	name string
	age  int
}

func main() {
	const dogName = "Buck"
	const dogAge = 1
	var myDog = Dog{
		Animal: Animal{
			color: "Black",
		},
		name: dogName,
		age:  dogAge,
	}
	fmt.Printf("myDog.Animal.getColor(): %v\n", myDog.Animal.getColor())
	fmt.Printf("myDog.Animal.color: %v\n", myDog.Animal.color)
}
