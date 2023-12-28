package main

import "fmt"

func nombre() {
	nombre, direccion := "Agustin de la Rosa -", " Av corrientes 939"
	fmt.Println(nombre + direccion)
}

func clima(temperatura float32, presion float32, humedad int) {
	fmt.Println("Temperatura = ", temperatura)
	fmt.Println("Presion = ", presion)
	fmt.Println("humedad = ", humedad)
}

func ej0() {
	fmt.Println("Ejercicio 0:")
	separadorCorto()
	nombre()
	separadorCorto()
	clima(27.5, 104.7, 64)
	separador()
}
