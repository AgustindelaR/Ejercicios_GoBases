package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImpuestosChocolateria_EmpleadoGanaCincuentaMil(t *testing.T) {
	//arrange
	salario := 50000
	exoectedResult := float32(salario) * (1 - 0.27) // 73 % salario
	//act
	obtainedResult := impuestosDeSalarioChocolateria(salario)
	//assert
	require.Equal(t, exoectedResult, obtainedResult)
}

func TestImpuestosChocolateria_EmpleadoGanaMenosDeCincuentaMil(t *testing.T) {
	//arrange
	salario := 10000
	exoectedResult := float32(salario) * (1 - 0.17) // 83 % salario
	//act
	obtainedResult := impuestosDeSalarioChocolateria(salario)
	//assert
	require.Equal(t, exoectedResult, obtainedResult)
}

func TestImpuestosChocolateria_EmpleadoGanaMasDeCincuentaMil(t *testing.T) {
	//arrange
	salario := 500000
	exoectedResult := float32(salario) * (1 - 0.27) // 73 % salario
	//act
	obtainedResult := impuestosDeSalarioChocolateria(salario)
	//assert
	require.Equal(t, exoectedResult, obtainedResult)
}

func TestCaclularPromedio_EscuelaRecibeListaVacia(t *testing.T) {
	//arrange
	var estudiantes []int
	expectedResult := 0.0
	//act
	obtainedResult := calcularPromedioNotas(estudiantes)
	//assert
	require.Equal(t, expectedResult, obtainedResult)
}
func TestCaclularPromedio_EscuelaRecibeListaNormal(t *testing.T) {
	//arrange
	estudiantes := []int{10, 8, 9, 4, 2, 9}
	expectedResult := 7.0
	//act
	obtainedResult := calcularPromedioNotas(estudiantes)
	//assert
	require.Equal(t, expectedResult, obtainedResult)
}

func TestCalcularSalarioCategorias_CategoriaA(t *testing.T) {
	//arrange
	horas := 20
	salarioBase := horas * 3000
	expectedResult := (salarioBase) + (salarioBase * (50 / 100))
	//act
	obtainedResult := calcularSalarioCategorias(horas, "A")
	//assert
	require.Equal(t, expectedResult, obtainedResult)
}
func TestCalcularSalarioCategorias_CategoriaB(t *testing.T) {
	//arrange
	horas := 20
	salarioBase := horas * 1500
	expectedResult := (salarioBase) + (salarioBase * (50 / 100))
	//act
	obtainedResult := calcularSalarioCategorias(horas, "B")
	//assert
	require.Equal(t, expectedResult, obtainedResult)
}

func TestCalcularSalarioCategorias_CategoriaC(t *testing.T) {
	//arrange
	horas := 20
	salarioBase := horas * 1000
	expectedResult := (salarioBase)
	//act
	obtainedResult := calcularSalarioCategorias(horas, "C")
	//assert
	require.Equal(t, expectedResult, obtainedResult)
}
