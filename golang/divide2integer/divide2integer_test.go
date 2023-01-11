package divide2integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide2integer(t *testing.T) {
	input := []struct {
		name   string
		i1     int
		i2     int
		result int
	}{
		{
			name:   "dvidend 10, divisor 3 result should be 3",
			i1:     10,
			i2:     3,
			result: 3,
		},
	}
	for _, i := range input {
		t.Run(i.name, func(tt *testing.T) {
			assert.Equal(tt, i.result, divide2integer(i.i1, i.i2))
		})
	}
}
