package main

import (
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/go-chi/chi/v5"
	"crud-go/internal/cliente"
	"crud-go/infra/database"
	"net/http"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("DATABASE_URL não foi definida")
	}

	db, err := database.NewPostgresPool(url)

	if err != nil {
		panic("Erro ao conectar")
		// log.Fatal("Erro ao conectar ", err )
	}

	defer db.Close()
	router := chi.NewRouter()

	cliente.NewRegisterModule(db, router)
	
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