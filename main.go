package main

import (
	"fmt"
)

func main() {
	// for i := 0; i <= 5; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"dog", "cat", "cow"}

	// for i, animal := range animals {
	// 	log.Println(i, animal)
	// }
	dog := Dog{
		Name: "Cat",
	}

	gorilla := Gorilla{
		Name: "Dog",
	}

	PrintInfor(dog, "dog")
	PrintInfor(gorilla, "gorilla")
}

func PrintInfor(a Animal, animalType string) {
	fmt.Println(animalType, "Says", a.Says(), "and has:", a.NumberOfLegs())
}

type Dog struct {
	Name string
}

func (d Dog) Says() string {
	return "Woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name string
}

func (g Gorilla) Says() string {
	return "ghummm"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

type Animal interface {
	Says() string
	NumberOfLegs() int
}
