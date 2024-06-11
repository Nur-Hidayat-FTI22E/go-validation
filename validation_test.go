package main

import (
	"encoding/json"
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

func TestValidationErrors(t *testing.T) {
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
		ValidatiErrors := err.(validator.ValidationErrors)
		for _, ValidatiError := range ValidatiErrors {
			fmt.Println("error in Field", ValidatiError.Field(), "on tag", ValidatiError.Tag(), "with error", ValidatiError.Error())
		}
	}
}

//sama dengan func TestStruct bedanya disini kita membuat agar pesan errornya lebih jelas terbaca dengan menggunakan ValidationError

func TestRegisterPass(t *testing.T) {

	type PassUser struct {
		Username        string `validate:"required"`
		Password        string `validate:"required,min=8"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
	}
	validation := validator.New()

	UserPass := PassUser{
		Username:        "Username",
		Password:        "12343564",
		ConfirmPassword: "1234354",
	}

	if err := validation.Struct(UserPass); err != nil {
		fmt.Println(err.Error())
	}
}

// NO KOMEN..!!!

func TestLoopingStruct(t *testing.T) {
	ValidasiStruct := validator.New()

	type Alamat struct {
		Provinsi  string `validate:"required"`
		Kabupaten string `validate:"required"`
		Kota      string `validate:"required"`
		Kecamatan string `validate:"required"`
		Jalan     string `validate:"required"`
	}

	type Kontak struct {
		NoHp  string `validate:"required,numeric,min=10,max=13"`
		NoWa  string `validate:"required,numeric,min=10,max=13"`
		Email string `validate:"required,email"`
	}

	type Info struct {
		NoId   int    `validate:"required"`
		Nama   string `validate:"required,uppercase"`
		Umur   string `validate:"required,number"`
		Alamat Alamat `validate:"required"`
		Kontak Kontak `validate:"required"`
	}

	DataDiri := Info{
		NoId: 12,
		Nama: "YOURNAME",
		Umur: "12",
		Alamat: Alamat{
			Provinsi:  "SulTra",
			Kabupaten: "Kolaka Utara",
			Kota:      "Lasusua",
			Kecamatan: "Lasusua",
			Jalan:     "Jln. Anonim",
		},
		Kontak: Kontak{
			NoHp:  "0988799081",
			NoWa:  "5478927391",
			Email: "example@example.com",
		},
	}

	if err := ValidasiStruct.Struct(DataDiri); err != nil {
		fmt.Println(err.Error())
		return
	}

	Mydata, err := json.MarshalIndent(DataDiri, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(Mydata))
}
