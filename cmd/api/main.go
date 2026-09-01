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

	// Factory
	repository := cliente.NewRepository(db)
	service    := cliente.NewService(repository)
	handler    := cliente.NewHandler(service)

	router := chi.NewRouter()

	router.Get("/clientes", handler.ListarTodosClientes,) 
	router.Post("/clientes", handler.AddCliente,) 
	router.Get("/clientes/{id}", handler.BuscarClientePorId,) 

	log.Println(
        "Servidor executando em http://localhost:8081",
    )

	err = http.ListenAndServe(
		":8081",
		router,
	)

	if err != nil {
		log.Fatal(err)
	}
}	