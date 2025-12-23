package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	// Busca o host da variável de ambiente ou usa localhost por padrão
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	// Monta a string de conexão usando a variável host
	stringDeConexao := fmt.Sprintf("host=%s user=root password=root dbname=root port=5432 sslmode=disable", host)
	
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}

	DB.AutoMigrate(&models.Aluno{})
}