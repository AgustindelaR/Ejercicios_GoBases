package main

import (
	"fmt"
)

func letrasDeUnaPalabra(palabra string) {
	println("La palabra es: ", palabra)
	i := 0
	for i < len(palabra) {
		fmt.Printf("%c,", palabra[i])
		i++
	}
	println()
	print("La palabra tiene ", i) // se rompe con 0 letras
	print(" letras")
	println()
}

func prestamosBanco(edad int, empleado bool, antiguedad int, sueldo int) {
	fmt.Printf("Caso de edad %d, empleado = %t, antiguedad = %d años, sueldo = %d \n",
		edad, empleado, antiguedad, sueldo)

	if edad >= 22 && empleado && antiguedad >= 1 {
		if sueldo >= 100000 {
			println("El banco cede prestamo sin interes")
		} else {
			println("El banco cede prestamo con interes")
		}
	} else {
		println("El banco no cede prestamo")
	}
}

func numeroDelMes(numero int) {
	meses := []string{"Enero", "Febrero", "marzo", "Abril", "Mayo", "Junio",
		"Julio", "Agosto", "Septiemre", "Octubre", "Noviembre", "Diciembre"}

	fmt.Printf("El numero %d, corresponde al mes %s\n", numero, meses[numero])
}

func queEdadTiene() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26,
		"Brenda": 19, "Darío": 44, "Pedro": 30}
	println("La edad de benjamin es: ", employees["Benjamin"])

	mayoresDeVeintiuno := 0
	for _, edad := range employees {
		if edad > 21 {
			mayoresDeVeintiuno++
		}
	}
	println("Los mayores a 21 son:", mayoresDeVeintiuno)
	fmt.Println(employees)
	println()

	println("Agrego a federico al mapa")
	employees["Federico"] = 25
	fmt.Println(employees)

	println()

	println("Elimino a Pedro")
	delete(employees, "Pedro")
	fmt.Println(employees)
}

func ej1() {
	println("Ejercicio 1: Estructuras de control")
	separadorCorto()

	letrasDeUnaPalabra("Agustin")
	separadorCorto()

	prestamosBanco(23, true, 3, 100001)
	separadorCorto()

	numeroDelMes(6)
	separadorCorto()

	queEdadTiene()
}
