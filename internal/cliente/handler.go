package cliente

import(
	"encoding/json"
    "net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddCliente(
	response http.ResponseWriter, 
	request *http.Request){

	var cliente Cliente

	err := json.NewDecoder(request.Body).Decode(&cliente)

	if err != nil {
		http.Error(
			response,
			"JSON invalido",
			http.StatusBadRequest,
		)

		return 
	}

	err = h.service.CadastrarCliente(cliente)

	if err != nil {
		http.Error(
			response,
			err.Error(),
			http.StatusBadRequest,
		)

		return 
	}

	response.WriteHeader(http.StatusCreated)
}

func(h *Handler) ListarTodosClientes(
	response http.ResponseWriter, 
	request *http.Request){

	clientes, err := h.service.ListarClientes()	

	if err != nil {
		http.Error(
			response,
			"Error ao buscar cliente",
			http.StatusInternalServerError,
		)

		return
	}

	response.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(response).Encode(clientes)
}
