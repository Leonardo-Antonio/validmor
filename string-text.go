package validmor

import "fmt"

type stringText struct {
	Min uint
	Max uint
}

func (s *stringText) Validate(val interface{}) (bool, error) {
	numberCharacters := len(val.(string))

	if numberCharacters == 0 {
		return false, fmt.Errorf("no puede estar vacio")
	}

	if numberCharacters < int(s.Min) {
		return false, fmt.Errorf("debe tener %v, como minimo", s.Min)
	}

	if s.Max >= s.Min && uint(numberCharacters) > s.Max {
		return false, fmt.Errorf("debe tener menos de %v caracteres de longitud", s.Max)
	}

	return true, nil
}
