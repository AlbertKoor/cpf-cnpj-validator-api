package cpf

func removerPontosETracos(cpf string) string {
	cpf_sem_pontos_e_tracos := ""

	for _, caracter := range cpf {
		if caracter >= '0' && caracter <= '9' {
			cpf_sem_pontos_e_tracos += string(caracter)
		}
	}
	return cpf_sem_pontos_e_tracos
}
