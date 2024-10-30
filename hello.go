package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaracion de constante
const Pi float32 = 3.14

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Hello())

	//Declaracion de Variables
	var firstName, lastName string
	var age uint8

	firstName = "Ignacio"
	lastName = "Palazzi"
	age = 27

	fmt.Println(firstName, lastName, "Edad:", age)
	fmt.Println(Pi)
	fmt.Println(Viernes)
}
