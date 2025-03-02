package controllers

import (
	
	"net/http"
)
//CriariUsuario cria um usuario no database
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando Usuário"))
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