package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/validmor"
)

type User struct {
	Id    int    `validmor:"number,min=5,max=1000"`
	Name  string `validmor:"string,min=2,max=100"`
	Bio   string `validmor:"string"`
	Email string `validmor:"mail"`
}

func main() {
	user := User{
		Id:    4,
		Name:  "superlongstring",
		Bio:   "",
		Email: "foobar",
	}

	fmt.Println("Errors:")
	for i, err := range validmor.ValidateStruct(user) {
		fmt.Printf("\t%d. %s\n", i+1, err.Error())
	}

}
