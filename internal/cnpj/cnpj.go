// ESCOLHA: CNPJ ganhou pacote próprio em vez de entrar em internal/cpf.
// São documentos diferentes, com tamanho, máscara e pesos diferentes.
// Bug no CNPJ? Abre esta pasta. Os testes também ficam separados:
// `go test ./internal/cnpj` roda só os dele.
package cnpj
