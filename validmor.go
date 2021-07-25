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
	case "string":
		validator := stringText{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return &validator
	case "number":
		validator := number{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
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
