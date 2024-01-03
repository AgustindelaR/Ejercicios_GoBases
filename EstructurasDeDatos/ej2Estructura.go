package main

import "fmt"

type Person struct {
	id          int
	name        string
	dateOfBirth string
}

type Employee struct {
	id       int
	position string
	person   Person
}

func (this Employee) PrintEmployee() {
	fmt.Printf("El empleado %s, con id %d", this.person.name, this.id)
	fmt.Printf(" en la posicion %s\n", this.position)
	fmt.Printf("tiene un id personal %d, y su cumpleaños es %s", this.person.id, this.person.dateOfBirth)
	fmt.Printf("\n")
}
