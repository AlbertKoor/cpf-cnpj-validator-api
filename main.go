package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlbertKoor/cpf-validator-api/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /status", handler.Status)
	mux.HandleFunc("POST /api/v1/validate_cnpj", handler.ValidarCNPJ)
	mux.HandleFunc("POST /api/v1/validate_cpf", handler.ValidarCPF)

	fmt.Println("Servidor rodando em http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", mux))

}
