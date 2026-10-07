package cnpj

func cnpjSemCaracteresEspeciais(cpf string) string {
	cnpj_sem_caracteres_especiais := ""

	for _, caracter := range cpf {
		if caracter >= '0' && caracter <= '9' {
			cnpj_sem_caracteres_especiais += string(caracter)
		}
	}
	return cnpj_sem_caracteres_especiais
}
