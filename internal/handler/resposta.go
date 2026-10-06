package handler

// ESCOLHA: a resposta ganhou um arquivo só dela porque vai ser
// COMPARTILHADA: o CPF e o CNPJ respondem igual, `{"valid": true}`.
// Se ficasse dentro de cpf.go, o handler de CNPJ ia depender de um
// arquivo que não tem nada a ver com ele.
//
// ESCOLHA: `resposta` com minúscula, porque só os handlers (deste mesmo
// pacote) usam. Ninguém de fora precisa enxergar, então não exporta.
//
// ATENÇÃO: o CAMPO `Valid` precisa ser maiúsculo mesmo com a struct
// minúscula. Quem transforma a struct em JSON é o pacote
// `encoding/json`, que é OUTRO pacote, e ele só enxerga campos
// exportados. Com `valid` minúsculo, a resposta sairia `{}`.
type resposta struct {
	Valid bool `json:"valid"`
}
