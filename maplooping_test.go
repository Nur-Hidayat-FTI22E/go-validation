package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type Info struct {
	NoId   int    `validate:"required"`
	Nama   string `validate:"required,uppercase"`
	Umur   string `validate:"required,number"`
	Alamat Alamat `validate:"required"`
	Kontak Kontak `validate:"required"`
}

type Alamat struct {
	Provinsi  string `validate:"required"`
	Kabupaten string `valdeate:"required"`
	Kota      string `valdeate:"required"`
	Kecamatan string `valdeate:"required"`
	Jalan     string `valdeate:"required"`
}

type Kontak struct {
	NoHp  string `valdeate:"required,numeric,min=10,max=13"`
	NoWa  string `valdeate:"required,numeric,min=10,max13"`
	Email string `validate:"required,email"`
}

func TestMapinMap(t *testing.T) {
	ValidasiStruct := validator.New()

	DataDiri := Info{
		NoId: 12,
		Nama: "DAYAT",
		Umur: "12",
		Alamat: Alamat{
			Provinsi:  "Sultra",
			Kabupaten: "Kolak Utara",
			Kota:      "Lasusua",
			Kecamatan: "Lasusua",
			Jalan:     "Jln. Basuki Rahmat",
		},
		Kontak: Kontak{
			NoHp:  "0988799081",
			NoWa:  "5478927391",
			Email: "dayat@gmail.com",
		},
	}

	if err := ValidasiStruct.Struct(DataDiri); err != nil {
		fmt.Println(err.Error())
	}

	Mydata, err := json.MarshalIndent(DataDiri, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(Mydata))
}
