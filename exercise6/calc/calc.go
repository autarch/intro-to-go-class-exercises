package calc

import "errors"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	if a < 0 {
		a *= -1
	}

	return a * b
}

func Div(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0.")
	}
	return float64(a) / float64(b), nil
}

func Sum(v ...int) int {
	a := 0
	for _, i := range v {
		a += i
	}
	return a
}
