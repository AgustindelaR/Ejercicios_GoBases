package Introdcucion

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
