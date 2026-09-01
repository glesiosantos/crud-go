package cliente

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRegisterModule(db *pgxpool.Pool, router chi.Router) {
	repository := NewRepository(db)
	service    := NewService(repository)
	handler    := NewHandler(service)

	router.Get("/clientes", handler.ListarTodosClientes,) 
	router.Post("/clientes", handler.AddCliente,) 
	router.Get("/clientes/{id}", handler.BuscarClientePorId,)
}