package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/validmor"
)

type User struct {
	Id    int    `validmor:"number,min=5,max=1000"`
	Name  string `validmor:"string,min=2,max=5"`
	Bio   int    `validmor:"number,required"`
	Email string `validmor:"mail,required"`
}

func main() {
	validmor.Errors(validmor.ERR_ES)
	user := User{
		Id:    4,
		Name:  "Leonardo",
		Bio:   0,
		Email: "novalid",
	}

	for i, err := range validmor.ValidateStruct(user) {
		fmt.Printf("%v.- %v\n", i+1, err.Error())
	}

}
