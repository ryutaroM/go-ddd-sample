package payment

import "ddd-structure/internal/domain/payment"

type PaymentResponse struct {
	ID       string  `json:"id"`
	Amount   float64 `json:"amount"`
	Quantity int64   `json:"quantity"`
	UserID   string  `json:"user_id"`
}

// ToResponse はドメインモデルからレスポンスDTOへの変換を行う
func ToResponse(p *payment.Payment) *PaymentResponse {
	return &PaymentResponse{
		ID:       string(p.ID),
		Amount:   float64(p.Amount),
		Quantity: int64(p.Quantity),
		UserID:   string(p.RefUserID),
	}
}

// ToResponseList は複数のドメインモデルをレスポンスDTOリストに変換する
func ToResponseList(payments []*payment.Payment) []*PaymentResponse {
	result := make([]*PaymentResponse, len(payments))
	for i, p := range payments {
		result[i] = ToResponse(p)
	}
	return result
}
