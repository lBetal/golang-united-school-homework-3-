package homework

func reverse(input []int64) (result []int64) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	result = append(result, input...)
	return
}
