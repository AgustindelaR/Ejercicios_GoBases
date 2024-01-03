package main

import "fmt"

func ejercicio_alumnos() {
	alumnoNuevo := Alumno{
		Name:    "Agustin",
		Surname: "de la Rosa",
		Dni:     123456789,
		Fecha:   "20/12/2023",
	}

	alumnoNuevo.Detalle()
	alumnoNuevo.ModificarNombre("AntiAgustin")
	alumnoNuevo.Detalle()
}

func ejercicio_productos() {
	smallProduct := CreateProduct("Small", 100)
	mediumProduct := CreateProduct("Medium", 100)
	largeProduct := CreateProduct("Large", 100)

	fmt.Printf("Precio total del producto Small: $%.2f\n", smallProduct.Price())
	fmt.Printf("Precio total del producto Medium: $%.2f\n", mediumProduct.Price())
	fmt.Printf("Precio total del producto Large: $%.2f\n", largeProduct.Price())
}

func main() {
	ejercicio_alumnos()
	println("------------------")
	ejercicio_productos()
}
