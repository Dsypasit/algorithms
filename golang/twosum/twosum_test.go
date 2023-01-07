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
		want          *[]int
	}{
		{
			name:          "input [1,2,3] expected 5 should be [2,3]",
			inputList:     []int{1, 2, 3},
			inputExpected: 5,
			want:          &[]int{2, 3},
		},
		{
			name:          "input [7,9,2,3,1] expected 10 should be [7,3]",
			inputList:     []int{7, 9, 2, 3, 1},
			inputExpected: 10,
			want:          &[]int{7, 3},
		},
	}

	for _, tt := range input {
		t.Run(tt.name, func(t *testing.T) {
			output := twosum(tt.inputList, tt.inputExpected)
			assert.Equal(t, tt.want, output)
		})
	}
}

func TestTwoSumError(t *testing.T) {
	input := []struct {
		name          string
		inputList     []int
		inputExpected int
		want          *[]int
	}{
		{
			name:          "input [] expected 10 should be [0,0]",
			inputList:     []int{},
			inputExpected: 10,
			want:          nil,
		},
	}

	for _, tt := range input {
		t.Run(tt.name, func(t *testing.T) {
			output := twosum(tt.inputList, tt.inputExpected)
			assert.Nil(t, tt.want, output)
		})
	}
}
