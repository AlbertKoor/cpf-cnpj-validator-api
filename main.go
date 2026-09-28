package main

import "fmt"

func main() {

	var cpf string = "52998224725" //cria a variável de forma longa de declaração
	//cpf := "..."	cria a variável	forma curta de declaração
	cpf = "52998224725" //troca o valor de uma variável que já existe - atribuição
	

	for i := 0; i < 9; i++ {
		fmt.Println("CPF:", cpf[i] - '0') //imprime o valor de cada dígito do CPF, subtraindo o valor de '0' para converter de ASCII para número
	}

}


