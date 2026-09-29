package cpf

import "testing"

func TestValidar(t *testing.T) {
	casos := []struct {
		nome     string
		entrada  string
		esperado bool
	}{
		{"válido sem máscara", "52998224725", true},
		{"válido com máscara", "529.982.247-25", true},
		{"dígito 1 errado", "52998224735", false},
		{"dígito 2 errado", "52998224726", false},
		{"todos os dígitos iguais", "11111111111", false},
		{"menos de 11 dígitos", "123", false},
		{"vazio", "", false},
		{"letra no lugar de um dígito", "529.98A.247-25", false},
		{"letra extra junto com 11 dígitos", "529982247A25", false},
		
	}

	for _, c := range casos {
		resultado := Validar(c.entrada)
		if resultado != c.esperado {
			t.Errorf("%s: Validar(%q) = %v, esperado %v", c.nome, c.entrada, resultado, c.esperado)
		}
	}
}