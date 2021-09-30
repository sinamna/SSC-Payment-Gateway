package models

type TransactionReq struct {
	OrderID string `json:"order_id"`
	TrxID string `json:"id"`
}
