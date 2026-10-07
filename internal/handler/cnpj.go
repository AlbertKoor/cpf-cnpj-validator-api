package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AlbertKoor/cpf-validator-api/internal/cnpj"
)

type pedidoCNPJ struct {
	CNPJ string `json:"cnpj"`
}

func ValidarCNPJ(w http.ResponseWriter, r *http.Request) {
	var p pedidoCNPJ

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	valido := cnpj.Validar(p.CNPJ)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta{Valid: valido})
}
