package cpf

import "regexp"

var formatoCPF = regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$|^\d{11}$`)

func Validar(cpf string) bool {

	if !formatoCPF.MatchString(cpf) {
		return false
	}

	cpf = removerPontosETracos(cpf)

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

func todosIguais(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}
