package main

func operation(criteria string) (int, string) {
	switch criteria {
	case "minimum":
		return minFunc(2, 3, 3, 4, 10, 2, 4, 5), ""
	case "average":
		return averageFunc(2, 3, 3, 4, 1, 2, 4, 5), ""
	case "maximum":
		return maxFunc(2, 3, 3, 4, 1, 2, 4, 5), ""
	default:
		return 0, "Criterio Inexistente"
	}
}

func minFunc(nums ...int) (min int) {
	min = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return
}
func averageFunc(nums ...int) (average int) {
	sum := 0
	for num := range nums {
		sum += num
	}
	average = sum / len(nums)
	return
}

func maxFunc(nums ...int) (max int) {
	max = 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return
}
func calcularEstadisticasEstudiantes(criteria string) (resultado int) {
	resultado, _ = operation(criteria)
	return
}
