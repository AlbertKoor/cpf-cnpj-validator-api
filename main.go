package main

import "fmt"

func main() {

	var cpf string = "52998224725" //cria a variável de forma longa de declaração
	//cpf := "..."	cria a variável	forma curta de declaração
	
	var soma = 0 //cria a variável somar e atribui o valor 0

	for i := 0; i < 9; i++ {
		digito := int(cpf[i] - '0') //converte o valor do dígito de ASCII para número

		peso := 10 - i //calcula o peso do dígito, que é 10 menos o índice do dígito

		soma += digito * peso //acumula o resultado da multiplicação na variável somar
	}

	var resto = soma % 11 //calcula o resto da divisão da soma por 11


	var digito1 = 0 //inicializa a variável digito1 com 0

	if resto < 2 {
		digito1 = 0 //se o resto for menor que 2, o dígito verificador é 0
	} else {
		digito1 = 11 - resto //se o resto for maior ou igual a 2, o dígito verificador é 11 menos o resto
	}

	fmt.Println("Soma total:", soma)
	fmt.Println("digito 1", digito1)


	if int(cpf[9]-'0') == digito1 { //compara o dígito verificador calculado com o dígito verificador do CPF
		fmt.Println("O primeiro dígito verificador está correto")
	} else {
		fmt.Println("O primeiro dígito verificador está incorreto")
	}
}


