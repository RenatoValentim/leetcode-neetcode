package arrayshashing

func TwoSum(nums []int, target int) []int {
	invalidInput := len(nums) == 0 || len(nums) < 2
	if invalidInput {
		return []int{}
	}

	m := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		idx, ok := m[complement]
		if ok {
			return []int{idx, i}
		}
		m[v] = i
	}

	return []int{}
}
