package main

func impuestosDeSalario(salario int) (resultado float32) {
	resultado = float32(salario) * (1 - 0.17)
	if salario >= 50000 {
		resultado = float32(salario) * (1 - 0.27)
	}
	print("El salario bruto es ", salario)
	println("y su total es ", resultado)
	return
}

func calcularPromedio(estudiantes []int) (promedio float64) {
	if len(estudiantes) == 0 {
		promedio = 0
		return
	}
	suma := 0
	for _, notaEstudiante := range estudiantes {
		suma += notaEstudiante
	}
	promedio = float64(suma) / float64(len(estudiantes))
	print("Los estudiantes son ", estudiantes)
	println(" y su promedio es ", promedio)
	return
}

func calcularSalario(horas int, categoria string) (salario int) {
	switch categoria {
	case "C":
		salario = horas * 1000
	case "B":
		salario = (horas * 1500) + (salario * (50 / 100))
	case "A":
		salario = (horas * 3000) + (salario * (50 / 100))
	default:
		salario = 0
		println("Categoria invalida.")
	}
	print("Su salario es ", salario)
	print(" con ", horas, " horas")
	println(" y categoria", categoria)
	return
}

func ej2() {
	println("Ejercicio 2: Funcionesl")
	separadorCorto()

	impuestosDeSalario(100000)
	separadorCorto()

	estudiantes := []int{10, 6, 8, 4, 5, 6}
	calcularPromedio(estudiantes)
	separadorCorto()

	calcularSalario(20, "A")
	separadorCorto()

}
