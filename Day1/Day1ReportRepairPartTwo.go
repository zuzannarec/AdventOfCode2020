package day1

func reportRepairPartTwo(input []int, requestedSum int) int {
	m := make(map[int]bool)
	for i := 0; i < cap(input); i++ {
		searched := requestedSum - input[i]
		for j := 0; j < cap(input); j++ {
			searched2 := searched - input[j]
			_, ok := m[searched2]
			if ok {
				result := searched2 * input[i] * input[j]
				return result
			}
			m[input[j]] = true
		}
		m[input[i]] = true
	}
	return 0
}
