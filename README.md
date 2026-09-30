# Validador de CPF — API em Go

## O que o app faz?

É uma API que recebe um CPF em JSON e responde se ele é válido ou não.

O CPF é considerado válido quando:

- tem 11 dígitos (com ou sem a máscara `000.000.000-00`);
- não tem letras nem outros caracteres além de números, `.` e `-`;
- não tem todos os dígitos iguais (ex.: `111.111.111-11`);
- os dois dígitos verificadores (os dois números depois do traço) batem com o cálculo feito a partir dos 9 primeiros dígitos.

O endpoint é:
`/api/v1/validate_cpf`

O método de requisição é `POST`.

Pedido:

```json
{ "cpf": "529.982.247-25" }
```

Resposta:

```json
{ "valid": true }
```

Se o JSON enviado estiver mal formado, a API responde `400 Bad Request` com a mensagem `JSON inválido`.

## Como executar o app localmente?

1. Instale o Go, versão 1.27.1 ou mais nova (a mesma do `go.mod`) -> https://go.dev/dl
2. Clone este repositório:
   ```bash
   git clone https://github.com/AlbertKoor/cpf-validator-api.git
   cd cpf-validator-api
   ```
3. Suba o servidor:
   ```bash
   go run .
   ```
4. O servidor fica rodando em `http://localhost:8080`.

Pronto! Agora dá para fazer o `POST` localmente.

No PowerShell (Windows):

```powershell
Invoke-RestMethod -Method Post -Uri http://localhost:8080/api/v1/validate_cpf -Body '{"cpf": "529.982.247-25"}'
```

No Linux/Mac, com curl:

```bash
curl -X POST http://localhost:8080/api/v1/validate_cpf -d '{"cpf": "529.982.247-25"}'
```

Também existe a rota `GET /status`, que responde `API está funcionando!`.

## Testes

Para rodar todos os testes:

```bash
go test ./...
```

Os testes ficam em `internal/cpf/cpf_test.go`. Cada caso tem uma entrada e o resultado esperado:

- CPF válido sem máscara → `true`
- CPF válido com máscara → `true`
- 1º dígito verificador errado → `false`
- 2º dígito verificador errado → `false`
- todos os dígitos iguais → `false`
- menos de 11 dígitos → `false`
- vazio → `false`
- letra no lugar de um dígito → `false`
- letra extra junto com 11 dígitos → `false`

## Decisões importantes

### - Por que Go?

É uma linguagem moderna, com tipagem forte: o compilador não deixa misturar tipos sem conversão. Durante o projeto isso me obrigou a converter o caractere do CPF (`byte`) para número (`int`) antes de fazer as contas, o que evita erro escondido. Além disso, a biblioteca padrão já tem servidor HTTP (`net/http`), JSON (`encoding/json`) e testes (`testing`), então o projeto não precisa de nenhuma dependência externa.

### - Por que a lógica fica em `internal/cpf`, separada do `main`?

O `main.go` cuida só da API: sobe o servidor, lê o JSON e devolve a resposta. Toda a regra de validação e a limpeza do CPF ficam no pacote `internal/cpf`. Assim dá para testar a validação sem subir servidor nenhum, e se um dia a regra mudar, a API não precisa mudar.

Só a função `Validar` é pública (letra maiúscula). As funções de apoio (`limpar`, `todosIguais`, `calcularDigito`) ficam privadas dentro do pacote.

### - O que acontece com máscara, letras ou dígitos iguais?

- **Máscara:** é aceita. Pontos e traço são removidos antes da validação, então `529.982.247-25` e `52998224725` dão o mesmo resultado.
- **Letras ou outros caracteres:** o CPF é rejeitado. Decidi não simplesmente ignorar a letra, porque `529982247A25` viraria um CPF válido depois de removida, e letra não faz parte de um CPF.
- **Todos os dígitos iguais:** é rejeitado. `111.111.111-11` passa na conta dos dígitos verificadores, mas não é um CPF de verdade.

### - Por que JSON inválido responde 400 e não `{"valid": false}`?

Porque o problema não está no CPF, está no pedido. No começo a API respondia `{"valid": false}` para um JSON quebrado, mesmo quando o CPF enviado era válido, e quem chamou acharia que o CPF estava errado. Com o `400` fica claro que o erro foi na montagem do pedido.

### - Como os testes foram escritos?

A ideia do projeto veio do [PasswordValidatorAPI](https://github.com/JoaoFRSeixas/PasswordValidatorAPI), que valida senhas. Os testes seguem o padrão de tabela do Go: uma lista de casos e um único loop que testa todos.

Parte deles foi escrita com TDD: para a regra das letras, escrevi primeiro o caso `529982247A25` esperando `false`, rodei e vi o teste falhar (o código aceitava esse CPF). Só depois mudei a função `limpar` até o teste passar.

## O que eu melhoraria

Hoje a resposta diz só se o CPF é válido ou não. Eu retornaria também qual regra falhou, por exemplo:

```json
{ "valid": false, "reason": "segundo dígito verificador inválido" }
```

Assim quem usa a API sabe o que corrigir. Também adicionaria testes para a camada HTTP (com `net/http/httptest`), cobrindo a rota, o `400` e o formato da resposta.
