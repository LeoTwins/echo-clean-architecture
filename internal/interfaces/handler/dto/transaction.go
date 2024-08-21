package dto

type TransactionResponse struct {
	ID     uint   `json:"id"`
	Type   string `json:"type"`
	Amount uint   `json:"uint"`
	Date   string `json:"string"`
}
