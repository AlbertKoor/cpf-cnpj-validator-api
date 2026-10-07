package cnpj

func calcularDigito(cnpj string, quantidade int) int {
	soma := 0

	for i := 0; i < quantidade; i++ {
		digito := int(cnpj[i] - '0')
		posDireita := quantidade - 1 - i
		peso := 2 + posDireita%8
		soma += digito * peso
	}

	resto := soma % 11

	if resto < 2 {
		return 0
	}

	return 11 - resto
}
