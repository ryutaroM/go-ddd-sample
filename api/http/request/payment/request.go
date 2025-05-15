package payment

import (
	"ddd-structure/internal/domain/payment"
	"encoding/json"
	"io"
)

type CreatePaymentRequest struct {
	ID         string  `json:"id"`
	Amount     float64 `json:"amount"`
	Quantity   int64   `json:"quantity"`
	CustomerID string  `json:"customer_id"`
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

// ToDomain はリクエストからドメインモデルを生成します
func (r *CreatePaymentRequest) ToDomain() *payment.Payment {
	return payment.NewPayment(payment.PaymentID(r.ID), payment.Amount(r.Amount), payment.Quantity(r.Quantity), r.CustomerID)
}
