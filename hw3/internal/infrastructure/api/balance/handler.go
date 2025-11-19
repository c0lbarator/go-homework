package balance

import (
	"encoding/json"
	"log"

	"net/http"

	"hw3/internal/application/dto/transfer"
	"hw3/internal/application/usecase"
)

type Handler struct {
	usecase usecase.BalanceUsecase
}

func NewHandler(uc usecase.BalanceUsecase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (handler *Handler) BalanceHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		UserId int `json:"UserId"`
	}
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		log.Printf("ERROR: failed to decode request body: %v", err)
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}

	balance, err := handler.usecase.GetBalance(request.Context(), req.UserId)

	if err != nil {
		log.Printf("ERROR: GetBalance failed: %v", err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	var resp struct {
		Balance float64 `json:"balance"`
	}
	resp.Balance = balance
	json.NewEncoder(writer).Encode(resp)
}

func (handler *Handler) DepositHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		UserId int     `json:"UserId"`
		Amount float64 `json:"Amount"`
	}
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		log.Printf("ERROR: failed to decode request body: %v", err)
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}
	err = handler.usecase.Deposit(request.Context(), req.UserId, req.Amount)
	if err != nil {
		log.Printf("ERROR: Deposit failed: %v", err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (handler *Handler) TransferHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req = transfer.TransferRequestDTO{}
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		log.Printf("ERROR: failed to decode request body: %v", err)
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}
	err = handler.usecase.Transfer(request.Context(), req)
	if err != nil {
		log.Printf("ERROR: Transfer failed: %v", err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
