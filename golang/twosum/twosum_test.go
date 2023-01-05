package twosum

import (
	"errors"
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
		{
			name:          "input [7,9,2,3,1] expected 10 should be [7,3]",
			inputList:     []int{7, 9, 2, 3, 1},
			inputExpected: 10,
			want:          [2]int{7, 3},
		},
	}

	for _, tt := range input {
		t.Run(tt.name, func(t *testing.T) {
			output, err := twosum(tt.inputList, tt.inputExpected)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, output)
		})
	}
}

func TestTwoSumError(t *testing.T) {
	input := []struct {
		name          string
		inputList     []int
		inputExpected int
		want          error
	}{
		{
			name:          "input [] expected 10 should be [0,0]",
			inputList:     []int{},
			inputExpected: 10,
			want:          errors.New("no value"),
		},
	}

	for _, tt := range input {
		t.Run(tt.name, func(t *testing.T) {
			_, err := twosum(tt.inputList, tt.inputExpected)
			assert.Equal(t, tt.want, err)
		})
	}
}
