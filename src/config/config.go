package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexao = ""
	Porta = 0
)

//Carregar vai inicializar variáveis de ambiente
func Carregar()  {
	var err error

	if err = godotenv.Load();err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 9000
	}

	StringConexao = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	os.Getenv("DB_USUARIO"),
	os.Getenv("DB_SENHA"),
	os.Getenv("DB_NOME"),
)
}