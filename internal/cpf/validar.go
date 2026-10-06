package cpf

func Validar(cpf string) bool {

	if temLetras(cpf) {
		return false
	}

	cpf = removerPontosETracos(cpf)

	if len(cpf) != 11 {
		return false
	}
	if todosIguais(cpf) {
		return false
	}

	digito1 := calcularDigitoCPF(cpf, 9)
	digito2 := calcularDigitoCPF(cpf, 10)

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
