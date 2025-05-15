package persistence

import (
	"ddd-structure/internal/domain/payment"

	"github.com/google/uuid"
)

type PaymentRepository struct {
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

func (r *PaymentRepository) FindByID(id payment.PaymentID) (*payment.Payment, error) {
	return payment.NewPayment(
		payment.PaymentID(uuid.New().String()),
		payment.Amount(1000),
		payment.Quantity(1),
		"customer_id",
	), nil
}

func (r *PaymentRepository) Save(p *payment.Payment) (*payment.Payment, error) {
	return payment.NewPayment(
		p.ID,
		p.Amount,
		p.Quantity,
		p.CustomerID,
	), nil
}
