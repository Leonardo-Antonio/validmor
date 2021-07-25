package validmor

import (
	"fmt"
	"strings"
)

type mail struct{}

func (m *mail) Validate(val interface{}) (bool, error) {
	if !strings.Contains(val.(string), "@") {
		return false, fmt.Errorf("no es un mail valido")
	}

	return true, nil
}
