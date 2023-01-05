package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	input := []struct {
		name          string
		inputList     []int
		inputExpected int
		want          [2]int
	}{
		{
			name:          "input [1,2,3] expected 5 should be [2,3]",
			inputList:     []int{1, 2, 3},
			inputExpected: 5,
			want:          [2]int{2, 3},
		},
	}

	for _, tt := range input {
		t.Run(tt.name, func(t *testing.T) {
			output := twosum(tt.inputList, tt.inputExpected)
			assert.Equal(t, tt.want, output)
		})
	}
}
