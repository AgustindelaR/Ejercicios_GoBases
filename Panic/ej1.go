package main

import (
	"fmt"
	"os"
)

func ej1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("The indicated file was not found nor is damaged")
		}
		println("Ejecución finalizada")
	}()

	rutaArchivo := ""
	_, errorRead := os.ReadFile(rutaArchivo)
	if errorRead != nil {
		panic(errorRead)
	}

}
