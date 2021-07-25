package validmor

import "fmt"

type number struct {
	Min uint
	Max uint
}

func (n *number) Validate(val interface{}) (bool, error) {

	num := val.(int)

	if uint(num) < n.Min {
		if currentTypeErrLang == ERR_ES {
			return false, fmt.Errorf("debería tener como mínimo %v", n.Min)
		}
		return false, fmt.Errorf("should have at least %v", n.Min)
	}

	if n.Max >= n.Min && uint(num) > n.Max {
		if currentTypeErrLang == ERR_ES {
			return false, fmt.Errorf("deberia ser menor %v", n.Max)
		}
		return false, fmt.Errorf("should be less %v", n.Max)
	}

	return true, nil
}
