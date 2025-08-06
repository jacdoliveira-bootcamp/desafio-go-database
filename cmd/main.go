package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/jacdoliveira/bw7/desafio-go-database/internal/application"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Arquivo .env não encontrado ou não pôde ser carregado: %v\n", err)
	}

	cfg := &application.ConfigApplicationDefault{
		Db: &mysql.Config{
			User:   os.Getenv("DB_USER"),
			Passwd: os.Getenv("DB_PASSWORD"),
			Net:    "tcp",
			Addr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
			DBName: os.Getenv("DB_NAME"),
		},
		Addr: fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
	}

	app := application.NewApplicationDefault(cfg)

	if err := app.SetUp(); err != nil {
		log.Fatalf("Erro durante setup da aplicação: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Erro durante execução da aplicação: %v", err)
	}
}
