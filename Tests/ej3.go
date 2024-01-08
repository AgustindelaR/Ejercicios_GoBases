package main

/*
Reescribo las funciones pero sin los prints
*/

func impuestosDeSalarioChocolateria(salario int) (resultado float32) {
	resultado = float32(salario) * (1 - 0.17)
	if salario >= 50000 {
		resultado = float32(salario) * (1 - 0.27)
	}
	return
}

func calcularPromedioNotas(estudiantes []int) (promedio float64) {
	if len(estudiantes) == 0 {
		promedio = 0
		return
	}

	suma := 0
	for _, notaEstudiante := range estudiantes {
		suma += notaEstudiante
	}
	promedio = float64(suma) / float64(len(estudiantes))
	return
}

func calcularSalarioCategorias(horas int, categoria string) (salario int) {
	switch categoria {
	case "C":
		salario = horas * 1000
	case "B":
		salario = (horas * 1500) + (salario * (50 / 100))
	case "A":
		salario = (horas * 3000) + (salario * (50 / 100))
	default:
		salario = 0
	}
	return
}
