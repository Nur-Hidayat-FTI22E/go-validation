package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	var (
		Password        string
		ConfirmPassword string
	)
	validate := validator.New()

	fmt.Println("Create Your Password")
	fmt.Scan(&Password)

	fmt.Println("Confirm Your Password")
	fmt.Scan(&ConfirmPassword)

	if err := validate.VarWithValue(Password, ConfirmPassword, "eqfield"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Succes create password")
	}

}
