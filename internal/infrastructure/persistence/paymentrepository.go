package persistence

import (
	"ddd-structure/internal/domain/payment"
)

type PaymentRepository struct {
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

func (r *PaymentRepository) FindByID(id payment.PaymentID) (*payment.Payment, error) {
	return payment.NewPayment( // mocking the payment object
		id,
		payment.Amount(1000),
		payment.Quantity(1),
		"user_id_1",
	), nil
}

func (r *PaymentRepository) Save(p *payment.Payment) (*payment.Payment, error) {
	return &payment.Payment{ // mocking the save operation
		ID:        p.ID,
		Amount:    p.Amount,
		Quantity:  p.Quantity,
		RefUserID: p.RefUserID,
	}, nil
}
