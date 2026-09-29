package main

import "fmt"

func main() {
	fmt.Println(validar("52998224725"))
	fmt.Println(validar("529.982.247-25"))
	fmt.Println(validar("11111111111"))
	fmt.Println(validar("123"))
	fmt.Println(validar(""))
}

func limpar(cpf string) string {
	limpo := ""

	for _, c := range cpf {
		if c >= '0' && c <= '9' {
			limpo += string(c)
		}
	}
	return limpo
}

func todosIguais(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

func validar(cpf string) bool {
	cpf = limpar(cpf)

	if len(cpf) != 11 {
		return false
	}
	if todosIguais(cpf) {
		return false
	}

	digito1 := calcularDigito(cpf, 9)
	digito2 := calcularDigito(cpf, 10)

	primeiroOk := int(cpf[9]-'0') == digito1
	segundoOk := int(cpf[10]-'0') == digito2

	return primeiroOk && segundoOk
}

// calcularDigito calcula um dígito verificador do CPF.
// quantidade = quantos dígitos entram na conta (9 para o 1º, 10 para o 2º).
func calcularDigito(cpf string, quantidade int) int {
	soma := 0

	for i := 0; i < quantidade; i++ {
		digito := int(cpf[i] - '0')
		peso := quantidade + 1 - i
		soma += digito * peso
	}

	resto := soma % 11

	verificador := 0
	if resto < 2 {
		verificador = 0
	} else {
		verificador = 11 - resto
	}

	return verificador
}
