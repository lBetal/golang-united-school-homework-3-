package homework

func average(input [15]float32) (result float32) {
	for _, val := range input {
		result += val
	}
	result = result / float32(len(input))
	return
}
