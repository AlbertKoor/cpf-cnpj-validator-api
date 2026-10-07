package cnpj

func cnpjSemCaracteresEspeciais(cnpj string) string {
	cnpj_sem_caracteres_especiais := ""

	for _, caracter := range cnpj {
		if caracter >= '0' && caracter <= '9' {
			cnpj_sem_caracteres_especiais += string(caracter)
		}
	}
	return cnpj_sem_caracteres_especiais
}
