package payment

import (
	"ddd-structure/internal/usecase/payment/command"
	"encoding/json"
	"io"
)

type CreatePaymentRequest struct {
	ID       string  `json:"id"`
	Amount   float64 `json:"amount"`
	Quantity int64   `json:"quantity"`
	UserID   string  `json:"user_id"`
}

// UnmarshalJSON
func NewCreatePaymentRequest(body io.ReadCloser) (*CreatePaymentRequest, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var req CreatePaymentRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *CreatePaymentRequest) ToCommand() (*command.CreatePaymentCommand, error) {
	return &command.CreatePaymentCommand{
		PaymentID: r.ID,
		Amount:    r.Amount,
		Quantity:  r.Quantity,
		UserID:    r.UserID,
	}, nil
}
