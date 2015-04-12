package calc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	cases := [][3]int{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 42, 41},
		{42, -1, 41},
		{999, 1234, 2233},
		{-5, -12, -17},
	}
	for _, c := range cases {
		assert.Equal(
			Add(c[0], c[1]),
			c[2],
			"%d + %d = %d",
			c[0], c[1], c[2],
		)
	}
}

func TestSub(t *testing.T) {
	assert := assert.New(t)

	cases := [][3]int{
		{3, 1, 2},
		{0, 0, 0},
		{1, 42, -41},
		{42, -1, 43},
		{-1, 42, -43},
		{1234, 100, 1134},
		{-5, -12, 7},
	}
	for _, c := range cases {
		assert.Equal(
			Sub(c[0], c[1]),
			c[2],
			"%d - %d = %d",
			c[0], c[1], c[2],
		)
	}
}

func TestMult(t *testing.T) {
	assert := assert.New(t)

	cases := [][3]int{
		{3, 1, 3},
		{0, 0, 0},
		{2, 42, 84},
		{42, -2, -84},
		{-1, 42, -42},
		{1234, 100, 123400},
		{-5, -12, 60},
	}
	for _, c := range cases {
		assert.Equal(
			Mult(c[0], c[1]),
			c[2],
			"%d * %d = %d",
			c[0], c[1], c[2],
		)
	}
}

func TestDiv(t *testing.T) {
	assert := assert.New(t)

	type divCase struct {
		a   int
		b   int
		r   float64
		err error
	}

	cases := []divCase{
		divCase{3, 1, 3, nil},
		divCase{0, 0, 0, errors.New("Cannot divide by 0.")},
		divCase{42, 2, 21, nil},
		divCase{42, -2, -21, nil},
		divCase{-1, 2, -0.5, nil},
		divCase{1234567, 100, 12345.67, nil},
		divCase{5, 4, 1.25, nil},
	}
	for _, c := range cases {
		r, err := Div(c.a, c.b)
		assert.Equal(err, c.err, "got expected error or nil from %d / %d", c.a, c.b)
		if c.err == nil {
			assert.Equal(
				r,
				c.r,
				"%d / %d = %f",
				c.a, c.b, c.r,
			)
		}
	}
}

func TestSum(t *testing.T) {
	assert := assert.New(t)

	cases := [][]int{
		{0},
		{1, 1},
		{1, 2, 3},
		{42, -2, 3, 43},
		{-1, -2, -3, -4, -10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 55},
	}

	for _, c := range cases {
		v := c[0 : len(c)-1]
		r := c[len(c)-1]
		assert.Equal(
			Sum(v...),
			r,
			"sum of %v = %d",
			v,
			r,
		)
	}
}
