package validmor

import "fmt"

type stringText struct {
	Min uint
	Max uint
}

func (s *stringText) Validate(val interface{}) (bool, error) {
	numberCharacters := len(val.(string))

	if numberCharacters < int(s.Min) {
		if currentTypeErrLang == ERR_ES {
			return false, fmt.Errorf("debe tener %v, como minimo", s.Min)
		}
		return false, fmt.Errorf("the name must be at least %v characters long", s.Min)
	}

	if s.Max >= s.Min && uint(numberCharacters) > s.Max {
		if currentTypeErrLang == ERR_ES {
			return false, fmt.Errorf("debe tener menos de %v caracteres de longitud", s.Max)
		}
		return false, fmt.Errorf("should be no more than %v characters", s.Max)
	}

	return true, nil
}
