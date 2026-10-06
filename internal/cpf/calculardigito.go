package cpf

func calcularDigito(cpf string, quantidade int) int {
	soma := 0

	for i := 0; i < quantidade; i++ {
		digito := int(cpf[i] - '0')
		peso := quantidade + 1 - i
		soma += digito * peso
	}

	resto := soma % 11

	if resto < 2 {
		return 0
	}

	return 11 - resto
}
