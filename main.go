package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlbertKoor/cpf-validator-api/internal/cpf"
	"net/http"
)

func main() {

	// o "roteador": decide qual handler atende cada rota
	mux := http.NewServeMux()

	// quando chegar um GET em /status, quem atende é a função apiStatus
	mux.HandleFunc("GET /status", apiStatus)
	mux.HandleFunc("POST /api/v1/validate_cpf", validarCPF)

	fmt.Println("Servidor rodando em http://localhost:8080")

	// liga o servidor na porta 8080 e fica esperando pedidos
	http.ListenAndServe(":8080", mux)
}

type pedido struct {
	CPF string `json:"cpf"`
}

type resposta struct {
	Valid bool `json:"valid"`
}

func validarCPF(w http.ResponseWriter, r *http.Request) {
	var p pedido
	
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil{
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	

	valido := cpf.Validar(p.CPF) // chama a função que valida o CPF

	w.Header().Set("Content-Type", "application/json") // diz que a resposta é JSON
	json.NewEncoder(w).Encode(resposta{Valid: valido})

}

// apiStatus é um handler: todo handler recebe esses dois parâmetros
//
//	w → onde você ESCREVE a resposta
//	r → o PEDIDO que chegou (rota, método, corpo...)
func apiStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API está funcionando!")
}
