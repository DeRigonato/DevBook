package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

//CriariUsuario cria um usuario no database
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil{
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar(); err != nil{
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuario.ID, err = repositorio.Criar(usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)

}
//BuscarUsuarios busca todos os usuarios no database
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorios := repositorios.NovoRepositorioDeUsuario(db)
	usuarios, err := repositorios.Buscar(nomeOuNick)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, usuarios)

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