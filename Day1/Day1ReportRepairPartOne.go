package day1

func reportRepairPartOne(input []int, requestedSum int) int {
	m := make(map[int]int)
	for i := 0; i < cap(input); i++ {
		searched := requestedSum - input[i]
		_, ok := m[searched]
		if ok {
			result := searched * input[i]
			return result
		}
		m[input[i]] = searched
	}
	return 0
}
