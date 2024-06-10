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

//membuat variabel validate dengan pointer validator
//memeriksa var validate apakah nilainya kosong atau tidak

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := ""

	if err := validate.Var(user, "required"); err != nil {
		fmt.Println(err.Error())
	}
}

//menambah variabel user bertipe data String
//memeriksa kondisi dimana var user harus memiliki value, jika kosong maka error

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()

	Password := 0
	ConfirmPassword := 0

	if err := validate.VarWithValue(Password, ConfirmPassword, "eqfield"); err != nil {
		fmt.Println(err.Error())
	}
}

//menambah variabel password dan confirmpassword bertipe data int
//memeriksa kondisi var Password dan ConfirmPassword dimana harus mempunyai value/nilai yang sama, jika tidak maka error

func TestValidationMultipleTag(t *testing.T) {
	validation := validator.New()

	noHp := "123"

	if err := validation.Var(noHp, "required,numeric"); err != nil {
		fmt.Println(err.Error())
	}
}

//membuat variabel noHp bertipe data string
//memeriksa var noHp yg dimana harus hanya memiliki value/nilai numeric walaupun di set dengan tipe data string, jika lain dari numeric atau valuenya kosong maka error

func TestValidationParameter(t *testing.T) {
	validation := validator.New()

	var myMoney int = 123456

	if err := validation.Var(myMoney, "required,min=5,max=10"); err != nil {
		fmt.Println(err.Error())
	}
}

//membuat variabel myMoney dengan tipe data string
//memeriksa apakah var myMoney memiliki value minimal 5 digit dan maximal 10 digit?, jika tidak maka akan error

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	RequestLogin := LoginRequest{
		Username: "struct@gmail.com",
		Password: "struct",
	}

	if err := validate.Struct(RequestLogin); err != nil {
		fmt.Println(err.Error())
	}
}

//membuat struct LoginRequest serta membuat map RequestLogin yg menangkap LoginRequest struct value
//memvalidasi vield struct yg ditangkap ke RequestLogin, jika memiliki value type email & password minimal 5 digit maka sukses, sebaliknya maka gagal
