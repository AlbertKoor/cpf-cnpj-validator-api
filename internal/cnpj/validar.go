package cnpj

import "regexp"

var formatoCNPJ = regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$|^\d{14}$`)

func Validar(cnpj string) bool {

	if !formatoCNPJ.MatchString(cnpj) {
		return false
	}

	cnpj = cnpjSemCaracteresEspeciais(cnpj)

	if todosIguais(cnpj) {
		return false
	}

	digito1 := calcularDigito(cnpj, 12)
	if int(cnpj[12]-'0') != digito1 {
		return false
	}

	digito2 := calcularDigito(cnpj, 13)

	return int(cnpj[13]-'0') == digito2

}

func todosIguais(cnpj string) bool {
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[0] {
			return false
		}
	}
	return true
}
