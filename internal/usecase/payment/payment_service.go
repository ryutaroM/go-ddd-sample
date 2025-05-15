package payment

import (
	"context"
	"ddd-structure/internal/domain/payment"
)

type Service interface {
	Find(ctx context.Context, id payment.PaymentID) (*payment.Payment, error)
	Create(ctx context.Context, payment *payment.Payment) (*payment.Payment, error)
}

type paymentService struct {
	repo payment.Repository
}

func NewPaymentService(repo payment.Repository) Service {
	return &paymentService{
		repo: repo,
	}
}

func (s *paymentService) Find(ctx context.Context, id payment.PaymentID) (*payment.Payment, error) {
	payment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *paymentService) Create(ctx context.Context, payment *payment.Payment) (*payment.Payment, error) {
	if payment, err := s.repo.Save(payment); err != nil {
		return nil, err
	} else {
		return payment, nil
	}
}
