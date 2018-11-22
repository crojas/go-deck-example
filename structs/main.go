package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	// Imprimi el atributo y el valor asignado
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	// NO! por que puede cambiar el orden de los atributos
	// yo := person{"Carlos", "Rojas"}

	// Correcto
	// var yo person
	// yo.firstName = "Carlos"
	// yo.lastName = "Rojas"
	// yo.contact.email = "c.rojas.latorre@gmail.com"

	// Correcto
	yo := person{
		firstName: "Carlos",
		lastName:  "Rojas",
		contactInfo: contactInfo{
			email:   "c.rojas.latorre@gmail.com",
			zipCode: 55555,
		},
	}
	yo.updateName("Jesús")
	yo.print()
}
