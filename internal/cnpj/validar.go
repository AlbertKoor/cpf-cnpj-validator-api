package cnpj

import "regexp"

var formatoCNPJ = regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$|^\d{14}$`)

func Validar(cnpj string) bool {

	if !formatoCNPJ.MatchString(cnpj) {
		return false
	}

	cnpj = cnpjSemCaracteresEspeciais(cnpj)

	if len(cnpj) != 14 {
		return false
	}
	if todosIguais(cnpj) {
		return false
	}

	digito1 := calcularDigito(cnpj, 12)
	digito2 := calcularDigito(cnpj, 13)

	primeiroOk := int(cnpj[12]-'0') == digito1
	segundoOk := int(cnpj[13]-'0') == digito2

	return primeiroOk && segundoOk
}

func todosIguais(cnpj string) bool {
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[0] {
			return false
		}
	}
	return true
}
