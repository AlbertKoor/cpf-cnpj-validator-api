// ESCOLHA: a pasta se chama `handler` e todo arquivo dela começa com
// `package handler`. Em Go, pasta = pacote. Tudo que fala HTTP (recebe
// pedido, devolve resposta) mora aqui; as regras de negócio (CPF, CNPJ)
// moram nas pastas delas e não sabem que HTTP existe.
package handler

import (
	"fmt"
	"net/http"
)

// ESCOLHA: o nome mudou de `apiStatus` para `Status`.
//   - Maiúscula: em Go, nome com letra maiúscula é EXPORTADO (público).
//     Agora quem chama é o main, que está em OUTRO pacote, então
//     com minúscula ele não enxergaria a função.
//   - Sem o "api": quem lê já vê `handler.Status` no main. O nome do
//     pacote faz parte do nome, então repetir contexto é ruído.
//
// Todo handler recebe os mesmos dois parâmetros:
//
//	w → onde você ESCREVE a resposta
//	r → o PEDIDO que chegou (rota, método, corpo...)
func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API está funcionando!")
}
