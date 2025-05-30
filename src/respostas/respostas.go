package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}){
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if dados != nil{
		if err := json.NewEncoder(w).Encode(dados); err != nil{
			log.Fatal(err)
		}
	}
	
}

func Erro(w http.ResponseWriter, statusCode int, err error){
	JSON(w, statusCode, struct{
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
}