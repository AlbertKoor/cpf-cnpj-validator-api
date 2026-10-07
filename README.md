# Validador de CPF e CNPJ — API em Go

## O que o app faz?

É uma API que recebe um CPF ou um CNPJ em JSON e responde se ele é válido ou não.

O **CPF** é considerado válido quando:

- tem 11 dígitos, no formato `52998224725` ou `529.982.247-25`;
- não tem todos os dígitos iguais (ex.: `111.111.111-11`);
- os dois dígitos verificadores batem com o cálculo feito a partir dos 9 primeiros dígitos.

O **CNPJ** é considerado válido quando:

- tem 14 dígitos, no formato `11222333000181` ou `11.222.333/0001-81`;
- não tem todos os dígitos iguais (ex.: `11.111.111/1111-11`);
- os dois dígitos verificadores batem com o cálculo feito a partir dos 12 primeiros dígitos.

Qualquer outro formato é rejeitado: letras, espaços, `#`, máscara pela metade etc.

### Rotas

| Método | Rota                    | Pedido                               |
| ------ | ----------------------- | ------------------------------------ |
| `POST` | `/api/v1/validate_cpf`  | `{ "cpf": "529.982.247-25" }`        |
| `POST` | `/api/v1/validate_cnpj` | `{ "cnpj": "11.222.333/0001-81" }`   |
| `GET`  | `/status`               | responde `API está funcionando!`     |

As duas rotas de validação respondem no mesmo formato:

```json
{ "valid": true }
```

Se o JSON enviado estiver mal formado, a API responde `400 Bad Request` com a mensagem `JSON inválido`.

## Como executar o app localmente?

1. Instale o Go, versão 1.27.1 ou mais nova (a mesma do `go.mod`) -> https://go.dev/dl
2. Clone este repositório:
   ```bash
   git clone https://github.com/AlbertKoor/cpf-cnpj-validator-api.git
   cd cpf-cnpj-validator-api
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
Invoke-RestMethod -Method Post -Uri http://localhost:8080/api/v1/validate_cnpj -Body '{"cnpj": "11.222.333/0001-81"}'
```

No Linux/Mac, com curl:

```bash
curl -X POST http://localhost:8080/api/v1/validate_cpf -d '{"cpf": "529.982.247-25"}'
curl -X POST http://localhost:8080/api/v1/validate_cnpj -d '{"cnpj": "11.222.333/0001-81"}'
```

Também dá para testar pelo Postman: método `POST`, a URL da rota e, em *Body → raw → JSON*, o pedido da tabela acima.

## Estrutura do projeto

```
cpf-cnpj-validator-api/
├── main.go               sobe o servidor e liga cada rota ao seu handler
└── internal/
    ├── cpf/              regra do CPF
    │   ├── validar.go        Validar (pública) + todosIguais
    │   ├── calculardigito.go cálculo dos dígitos verificadores
    │   ├── limpar.go         tira a máscara
    │   └── cpf_test.go
    ├── cnpj/             regra do CNPJ (mesma organização do cpf)
    │   ├── validar.go
    │   ├── calculardigito.go
    │   ├── limpar.go
    │   └── cnpj_test.go
    └── handler/          tudo que fala HTTP
        ├── cpf.go            POST /api/v1/validate_cpf
        ├── cnpj.go           POST /api/v1/validate_cnpj
        ├── resposta.go       formato da resposta, usado pelos dois
        └── status.go         GET /status
```

## Testes

Para rodar todos os testes:

```bash
go test ./...
```

Os testes ficam em `internal/cpf/cpf_test.go` (11 casos) e `internal/cnpj/cnpj_test.go` (9 casos). Cada caso tem uma entrada e o resultado esperado:

- válido sem máscara → `true`
- válido com máscara → `true`
- 1º dígito verificador errado → `false`
- 2º dígito verificador errado → `false`
- todos os dígitos iguais → `false`
- tamanho errado → `false`
- vazio → `false`
- letra no lugar de um dígito → `false`
- caractere especial (ex.: `#`) → `false`
- espaços no lugar da máscara (só CPF) → `false`

## Decisões importantes

### - Por que Go?

É uma linguagem moderna, com tipagem forte: o compilador não deixa misturar tipos sem conversão. Durante o projeto isso me obrigou a converter o caractere do CPF (`byte`) para número (`int`) antes de fazer as contas, o que evita erro escondido. Além disso, a biblioteca padrão já tem servidor HTTP (`net/http`), JSON (`encoding/json`), regex (`regexp`) e testes (`testing`), então o projeto não precisa de nenhuma dependência externa.

### - Por que essa divisão de pastas?

Cada pasta é um pacote, e cada pacote cuida de um assunto só. A regra é: se der bug no X, tem que ser óbvio em qual pasta procurar.

- `internal/cpf` e `internal/cnpj` só sabem validar. Não sabem que HTTP existe, então dá para testar a validação sem subir servidor nenhum.
- `internal/handler` só traduz: lê o JSON, chama o `Validar` do pacote certo e escreve a resposta.
- `main.go` só monta a API. Para adicionar o CNPJ, aqui mudou uma linha só (a rota nova).

Só as funções que outro pacote precisa chamar são públicas (letra maiúscula), como `cpf.Validar` e `handler.ValidarCNPJ`. As funções de apoio (`todosIguais`, `calcularDigito` etc.) ficam privadas dentro do pacote.

O CPF e o CNPJ têm algumas funções parecidas (`todosIguais`, a regra do resto), mas cada pacote tem a sua cópia. Preferi repetir umas poucas linhas a criar um pacote compartilhado que ligaria os dois: se a regra de um mudar, a do outro não quebra.

### - Por que usar regex logo no começo?

A primeira coisa que o `Validar` faz é conferir o formato com uma expressão regular. Se não bater, ele retorna `false` na hora (*early return*), sem fazer conta nenhuma.

Antes disso, o código só barrava letras e ignorava o resto, então `529#982#247#25` era aceito como válido. O regex resolveu esse bug e ainda deixou o código mais curto, porque a função que procurava letras deixou de ser necessária.

O regex fica numa variável fora da função, para ser montado uma vez só quando o programa liga, e não a cada pedido.

### - Por que a máscara pela metade é rejeitada?

O formato aceito é rígido: ou só números, ou a máscara completa. `529.982247-25` é rejeitado. Achei mais fácil de explicar e de testar uma regra clara do que aceitar qualquer combinação de pontos e traços.

### - Como o cálculo do CNPJ difere do CPF?

A regra do resto é a mesma (`soma % 11`, e o dígito é `0` se o resto for menor que 2, senão `11 - resto`). O que muda são os pesos: no CPF eles descem em linha reta (10, 9, 8...), no CNPJ eles vão de 2 a 9 da direita para a esquerda e depois voltam para o 2 (`5 4 3 2 9 8 7 6 5 4 3 2`). Para isso usei o resto da divisão por 8:

```go
posDireita := quantidade - 1 - i
peso := 2 + posDireita%8
```

### - Por que JSON inválido responde 400 e não `{"valid": false}`?

Porque o problema não está no documento, está no pedido. No começo a API respondia `{"valid": false}` para um JSON quebrado, mesmo quando o CPF enviado era válido, e quem chamou acharia que o CPF estava errado. Com o `400` fica claro que o erro foi na montagem do pedido.

### - Como os testes foram escritos?

A ideia do projeto veio do [PasswordValidatorAPI](https://github.com/JoaoFRSeixas/PasswordValidatorAPI), que valida senhas. Os testes seguem o padrão de tabela do Go: uma lista de casos e um único loop que testa todos.

Usei TDD: primeiro escrevo o caso que deveria passar, rodo e vejo o teste falhar, e só depois mexo no código. Foi assim que apareceu o bug do `#` no CPF, e foi assim que o CNPJ foi feito, com os 9 casos escritos antes da função `Validar` existir.

## Próximos passos

- **CNPJ alfanumérico:** desde julho de 2026 a Receita Federal emite CNPJs com letras nas 12 primeiras posições. O cálculo é o mesmo, usando o valor do caractere na tabela ASCII menos 48, então a mudança principal é no regex.
- **Dizer qual regra falhou:** hoje a resposta diz só se é válido ou não. Eu retornaria também o motivo, por exemplo `{ "valid": false, "reason": "segundo dígito verificador inválido" }`.
- **Testes da camada HTTP:** com `net/http/httptest`, cobrindo as rotas, o `400` e o formato da resposta.
- **Roteador:** testar o [chi](https://go-chi.io/) no lugar do `http.ServeMux` da biblioteca padrão, para comparar.
