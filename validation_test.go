package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	_ "golang.org/x/tools/go/analysis/passes/printf"
)

func TestValidation(t *testing.T) {
	validate := validator.New()

	if validate == nil {
		t.Error("Validasinya kosong")
	}

}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := ""

	if err := validate.Var(user, "required"); err != nil {
		fmt.Println(err.Error())
	}

}

func TestValidationTwoVariables(t *testing.T) {
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
	}

}
