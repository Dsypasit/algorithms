package mergetwolist

import (
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	input := []struct {
		name   string
		input1 *Node
		input2 *Node
		want   *Node
	}{
		{
			name:   "l1 [1,2,3] l2 [1,2,3] should be [1,1,2,2,3,3]",
			input1: addAll(newNode(1, nil), newNode(2, nil), newNode(3, nil)),
			input2: addAll(newNode(1, nil), newNode(2, nil), newNode(3, nil)),
			want: addAll(
				newNode(1, nil),
				newNode(1, nil),
				newNode(2, nil),
				newNode(2, nil),
				newNode(3, nil),
				newNode(3, nil),
			),
		},
	}

	for _, tt := range input {
		t.Run(tt.name, func(t *testing.T) {
			result := mergetwolist(tt.input1, tt.input2)
			if !equal(tt.want, result) {
				t.Errorf("Not equal")
			}
		})
	}
}
