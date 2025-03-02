package rotas

import "net/http"

//Rotas representa todas as rotas da API
type Rota struct{
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}