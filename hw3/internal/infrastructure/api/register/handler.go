package register

import (
	"encoding/json"
	"net/http"

	"hw3/internal/application/usecase"
)

type Handler struct {
	usecase usecase.RegisterUsecase
}

func NewHandler(uc usecase.RegisterUsecase) *Handler {
	return &Handler{
		usecase: uc,
	}
}
func (handler *Handler) RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, err := handler.usecase.Register(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	var resp struct {
		UserId int `json:"user_id"`
	}
	resp.UserId = userId
	json.NewEncoder(writer).Encode(resp)
}
