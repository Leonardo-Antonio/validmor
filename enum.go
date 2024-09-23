package validmor

import (
	"fmt"
	"strings"
)

type enum struct {
	Values []string
}

func (n *enum) Validate(val interface{}) (bool, error) {
	ok := false
	value := val.(string)
	for i := 0; i < len(n.Values); i++ {
		if value == n.Values[i] {
			ok = true
			break
		}
	}

	if ok {
		return ok, nil
	}

	return ok, fmt.Errorf("is not a valid value [%s] (options: %s)", value, strings.Join(n.Values, ", "))
}
