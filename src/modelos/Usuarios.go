package modelos

import (
	"errors"
	"strings"
	"time"
)

//Usuario representa um user usando a rede social
type Usuario struct{
	ID uint64 `json:"id,omitempty"` //omit empty não retorna valor vazio
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEM,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error{
	if err := usuario.validar(etapa); err != nil{
		return err
	}

	usuario.formatar()
	return nil

}

func (usuario *Usuario) validar(etapa string) error{
	if usuario.Nome == ""{
		return errors.New("O campo nome não pode estar em branco")
	}

	if usuario.Nick == ""{
		return errors.New("O campo nick não pode estar em branco")
	}

	if usuario.Email == ""{
		return errors.New("O campo email não pode estar em branco")
	}

	if etapa == "cadastro" && usuario.Senha == ""{
		return errors.New("O campo senha não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(){
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}


