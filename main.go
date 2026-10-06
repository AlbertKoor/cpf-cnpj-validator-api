package main

// ESCOLHA: o main.go só MONTA a aplicação.
// Ele não sabe validar CPF nem ler JSON; só diz "esta rota vai para
// aquele handler" e liga o servidor. Assim, quando entrar o CNPJ, aqui
// muda uma linha só (a rota nova), e a lógica fica nas pastas certas.

// ESCOLHA: imports em dois grupos, separados por linha em branco.
// Em cima, a biblioteca padrão do Go; embaixo, os pacotes do próprio
// projeto. É a convenção da comunidade e o `gofmt` mantém a ordem.
import (
	"fmt"
	"net/http"

	"github.com/AlbertKoor/cpf-validator-api/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	// ESCOLHA: lendo as rotas, dá para ver a API inteira em 2 linhas.
	// `handler.ValidarCPF` diz de onde vem (pacote handler) e o que faz.
	mux.HandleFunc("GET /status", handler.Status)
	mux.HandleFunc("POST /api/v1/validate_cpf", handler.ValidarCPF)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
