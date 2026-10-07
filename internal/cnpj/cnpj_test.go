package cnpj

import "testing"

func TestValidar(t *testing.T) {
	casos := []struct {
		nome     string
		entrada  string
		esperado bool
	}{
		{"válido sem máscara", "11222333000181", true},
		{"válido com máscara", "11.222.333/0001-81", true},
		{"dígito 1 errado", "11222333000191", false},
		{"dígito 2 errado", "11222333000182", false},
		{"todos os dígitos iguais", "11111111111111", false},
		{"tamanho errado", "1122233300018", false},
		{"vazio", "", false},
		{"caractere especial", "11.222.333#0001-81", false},
		{"letra minuscula", "11.222.333/0001-8a", false},
	}

	for _, c := range casos {
		resultado := Validar(c.entrada)
		if resultado != c.esperado {
			t.Errorf("%s: Validar(%q) = %v, esperado %v", c.nome, c.entrada, resultado, c.esperado)
		}
	}
}
