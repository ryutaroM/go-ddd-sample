package command

type CreatePaymentCommand struct {
	PaymentID  string  `json:"payment_id"`
	Amount     float64 `json:"amount"`
	Quantity   int64   `json:"quantity"`
	CustomerID string  `json:"customer_id"`
}
