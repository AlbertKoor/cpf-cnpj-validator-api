package main


import (
	"fmt"

	"github.com/AlbertKoor/cpf-validator-api/internal/cpf"
)

func main() {
	fmt.Println(cpf.Validar("52998224725"))
	fmt.Println(cpf.Validar("529.982.247-25"))
	fmt.Println(cpf.Validar("11111111111"))
	fmt.Println(cpf.Validar("123"))
	fmt.Println(cpf.Validar(""))
}

