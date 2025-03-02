package modelos

import "time"

//Usuario representa um user usando a rede social
type Usuario struct{
	ID uint64 `json:"id,omitempty"` //omit empty não retorna valor vazio
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEM,omitempty"`
}

