package cpf

func calcularDigitoCPF(cpf string, quantidade int) int {
	soma := 0

	for i := 0; i < quantidade; i++ {
		digito := int(cpf[i] - '0')
		peso := quantidade + 1 - i
		soma += digito * peso
	}

	resto := soma % 11

	verificador := 0
	if resto < 2 {
		return 0
	} else {
		verificador = 11 - resto
	}

	return verificador
}
