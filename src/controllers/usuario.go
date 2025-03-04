package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

//CriariUsuario cria um usuario no database
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario modelos.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil{
		log.Fatal(err)
	}

	db, err := banco.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	repositorio.Criar(usuario)

}
//BuscarUsuarios busca todos os usuarios no database
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando todos Usuário"))
}
//BuscarUsuario busca um usuário no database
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Busacando um Usuário"))
}
//AtualizarUsuario atualiza usuario no database
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando Usuário"))
}
//Deletar usuario exclui um usuário no database
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando Usuário"))
}