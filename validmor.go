package validmor

import (
	"fmt"
	"reflect"
	"strings"
)

type (
	Validator interface {
		Validate(interface{}) (bool, error)
	}

	DefaultValidator struct{}
)

func (d *DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "enum":
		validator := enum{}
		validator.Values = getEnumValues(args[1])
		return &validator
	case "string":
		validator := stringText{}
		min, max := getMinMax(args[1:])
		validator.Min = min
		validator.Max = max
		return &validator
	case "number":
		validator := number{}
		min, max := getMinMax(args[1:])
		validator.Min = min
		validator.Max = max
		return &validator
	case "mail":
		validator := mail{}
		return &validator
	}

	for _, v := range args {
		if v == "required" {
			validator := required{}
			return &validator
		}
	}

	return &DefaultValidator{}
}

// validate the structure depending on the tags you have added to your structure
func ValidateStruct(s interface{}) []error {
	errs := []error{}

	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("validmor")

		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidatorFromTag(tag)

		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}

func getMinMax(args []string) (min, max uint) {
	switch len(args) {
	case 2:
		fmt.Sscanf(strings.Join(args, ","), "min=%d,max=%d", &min, &max)
	case 1:
		if strings.Contains(args[0], "min=") {
			fmt.Sscanf(args[0], "min=%d", &min)
		}
		if strings.Contains(args[0], "max=") {
			fmt.Sscanf(args[0], "max=%d", &max)
		}
	}

	return
}

func getEnumValues(arg string) []string {
	values := strings.Split(arg, "=")
	if len(values) != 2 {
		return []string{}
	}

	return strings.Split(values[1][1:len(values[1])-1], ";")
}
