package main

import (
	"context"
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/go-chi/chi/v5"
	"crud-go/internal/cliente"
	"crud-go/infra/database"
	"basico-crud-go/infra/keycloack"
	"net/http"
)

func main() {
	context := context.Background()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("DATABASE_URL não foi definida")
	}

	db, err := database.NewPostgresPool(url, context)

	if err != nil {
		panic("Erro ao conectar")
		// log.Fatal("Erro ao conectar ", err )
	}

	defer db.Close()

	keycloakClient,err := keycloack.NewKeycloak(context)
	router := chi.NewRouter()

	cliente.NewRegisterModule(db, router, keycloakClient)
	
	log.Println(
        "Servidor executando em http://localhost:8082",
    )

	err = http.ListenAndServe(
		":8082",
		router,
	)

	if err != nil {
		log.Fatal(err)
	}
}	