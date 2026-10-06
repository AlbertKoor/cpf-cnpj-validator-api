package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AlbertKoor/cpf-validator-api/internal/cpf"
)

// ESCOLHA: `pedido` virou `pedidoCPF`. Todos os arquivos da pasta
// handler são o MESMO pacote, então dois tipos não podem ter o mesmo
// nome. Quando vier o CNPJ, ele vai ter o `pedidoCNPJ` dele, com o
// campo `cnpj`, e os dois convivem sem conflito.
type pedidoCPF struct {
	CPF string `json:"cpf"`
}

// ESCOLHA: `validarCPF` virou `ValidarCPF` (maiúscula) pelo mesmo motivo
// do Status: o main, que está em outro pacote, precisa chamar.
//
// ESCOLHA: o handler não sabe COMO se valida um CPF. Ele só traduz:
// JSON → string → cpf.Validar → bool → JSON. Se a regra do CPF mudar,
// este arquivo nem é aberto.
func ValidarCPF(w http.ResponseWriter, r *http.Request) {
	var p pedidoCPF

	// ESCOLHA: early return. JSON quebrado sai na hora com 400, antes de
	// gastar qualquer trabalho com validação.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	valido := cpf.Validar(p.CPF)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta{Valid: valido})
}
