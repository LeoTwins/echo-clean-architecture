package dto

type OpenAccountRequest struct {
	Name    string `json:"name"`
	Balance uint   `json:"balance"`
}

type OpenAccountResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Balance uint   `json:"balance"`
}

type DepositAndWithdrawRequest struct {
	ID     uint `json:"id"`
	Amount uint `json:"amount"`
}

type TransferRequest struct {
	ID          uint `json:"id"`
	ToAccountID uint `json:"to_account_id"`
	Amount      uint `json:"amount"`
}
