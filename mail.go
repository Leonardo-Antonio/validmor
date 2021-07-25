package validmor

import (
	"fmt"
	"strings"
)

type mail struct{}

func (m *mail) Validate(val interface{}) (bool, error) {
	if !strings.Contains(val.(string), "@") {
		if currentTypeErrLang == ERR_ES {
			return false, fmt.Errorf("no es un email valido")
		}
		return false, fmt.Errorf("it is not a valid email")
	}

	return true, nil
}
