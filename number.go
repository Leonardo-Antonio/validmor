package validmor

import "fmt"

type number struct {
	Min uint
	Max uint
}

func (n *number) Validate(val interface{}) (bool, error) {

	num := val.(int)

	if uint(num) < n.Min {
		return false, fmt.Errorf("deberia tener como minimo %v", n.Min)
	}

	if n.Max >= n.Min && uint(num) > n.Max {
		return false, fmt.Errorf("deberia ser menor %v", n.Max)
	}

	return true, nil
}
