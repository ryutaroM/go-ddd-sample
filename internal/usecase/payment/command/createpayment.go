package command

type CreatePaymentCommand struct {
	PaymentID string  `json:"payment_id"`
	Amount    float64 `json:"amount"`
	Quantity  int64   `json:"quantity"`
	UserID    string  `json:"user_id"`
}
