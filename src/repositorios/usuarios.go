package repositorios

import (
	"api/src/modelos"
	
	"database/sql"
	"fmt"
)

type usuarios struct{
	db *sql.DB
}

func NovoRepositorioDeUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}
	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	return uint64(ultimoIDInserido), nil
}

func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error){
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",  nomeOuNick, nomeOuNick,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next(){
		var usuario modelos.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil{
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}