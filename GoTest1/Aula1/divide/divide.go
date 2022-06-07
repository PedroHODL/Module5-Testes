package divide

import "fmt"

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("O denominador não pode ser 0")
	}
	return a / b, nil
}
