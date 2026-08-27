package main

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-chi/chi/v5"
	"crud-go/internal/cliente"
	"net/http"
)

func main() {
	url := "postgres://postgres:102030@localhost:5432/cruddb"
	db, err := pgxpool.New(context.Background(), url)

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