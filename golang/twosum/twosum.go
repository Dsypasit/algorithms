package twosum

func twosum(list []int, expected int) *[]int {
	temp := make(map[int]int)
	for i, v := range list {
		if _, ok := temp[expected-v]; !ok {
			temp[expected-v] = i
		}
		if r, ok := temp[v]; ok {
			return &[]int{list[r], v}
		}
	}
	return nil
}
