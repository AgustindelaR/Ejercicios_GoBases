package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcularEstadisticasEstudiantes_Minimo(t *testing.T) {
	expectedResult := 2

	obtainedResult := calcularEstadisticasEstudiantes("minimum")

	require.Equal(t, expectedResult, obtainedResult)
}

func TestCalcularEstadisticasEstudiantes_Maximum(t *testing.T) {
	expectedResult := 5

	obtainedResult := calcularEstadisticasEstudiantes("maximum")

	require.Equal(t, expectedResult, obtainedResult)
}

func TestCalcularEstadisticasEstudiantes_average(t *testing.T) {
	expectedResult := 3

	obtainedResult := calcularEstadisticasEstudiantes("average")

	require.Equal(t, expectedResult, obtainedResult)
}
