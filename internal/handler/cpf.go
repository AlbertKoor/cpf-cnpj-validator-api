package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AlbertKoor/cpf-cnpj-validator-api/internal/cpf"
)

type pedidoCPF struct {
	CPF string `json:"cpf"`
}

func ValidarCPF(w http.ResponseWriter, r *http.Request) {
	var p pedidoCPF

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	valido := cpf.Validar(p.CPF)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta{Valid: valido})
}
