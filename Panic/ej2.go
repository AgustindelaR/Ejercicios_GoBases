package main

import (
	"fmt"
	"os"
)

func ej2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("No se encontró el atchivo indicado")
		}
		println("Ejecución finalizada")
	}()

	rutaArchivo := "customers.txt"
	datos, errorRead := os.ReadFile(rutaArchivo)
	if errorRead != nil {
		panic(errorRead)
	}

	println(string(datos))

}
