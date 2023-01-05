package twosum

import "errors"

func twosum(list []int, expected int) ([2]int, error) {
	temp := make(map[int]int)
	for i, v := range list {
		if _, ok := temp[expected-v]; !ok {
			temp[expected-v] = i
		}
		if r, ok := temp[v]; ok {
			return [2]int{list[r], v}, nil
		}
	}
	return [2]int{}, errors.New("no value")
}
