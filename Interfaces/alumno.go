package main

import "fmt"

type Alumno struct {
	Name    string
	Surname string
	Fecha   string //fecha de ingreso
	Dni     int
}

func (a Alumno) Detalle() {
	fmt.Printf("Alumno: %s %s con dni: %d,", a.Name, a.Surname, a.Dni)
	fmt.Printf("fecha de ingreso: %s\n", a.Fecha)
}

func (a *Alumno) ModificarNombre(nombreNuevo string) {
	(*a).Name = nombreNuevo
}

func (a *Alumno) ModificarApellido(apellidoNuevo string) {
	(*a).Surname = apellidoNuevo
}

func (a *Alumno) ModificarFecha(fechaNueva string) {
	(*a).Name = fechaNueva
}

func (a *Alumno) ModificarDNI(dniNuevo int) {
	(*a).Dni = dniNuevo
}
