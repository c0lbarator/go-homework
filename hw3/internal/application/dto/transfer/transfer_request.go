package transfer

type TransferRequestDTO struct {
	FromId int     `json:"from_id"`
	ToId   int     `json:"to_id"`
	Amount float64 `json:"amount"`
}
