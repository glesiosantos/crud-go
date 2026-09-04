package cliente

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"crud-go/infra/keycloak"
	appMiddleware "crud-go/internal/middleware"
)

func NewRegisterModule(
	db *pgxpool.Pool, 
	router chi.Router, 
	keycloakClient *keycloak.Client
) {
	repository := NewRepository(db)
	service    := NewService(repository)
	handler    := NewHandler(service)

	router.Route("/clientes", func(r chi.Router){
		r.Group(func(private chi.Router){
			private.Use(appMiddleware.Auth(keycloakClient))
			private.Get("/", handler.ListarTodosClientes,)
			private.Post("/", handler.AddCliente,) 
			private.Get("/{id}", handler.BuscarClientePorId,) 
		})
	})
}