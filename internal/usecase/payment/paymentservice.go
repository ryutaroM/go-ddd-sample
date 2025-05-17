package payment

import (
	"context"
	useracl "ddd-structure/internal/acl/payment/user"
	"ddd-structure/internal/domain/payment"
	"ddd-structure/internal/usecase/payment/command"
	"ddd-structure/internal/usecase/payment/query"
)

type Service interface {
	Find(ctx context.Context, query *query.FindPaymentQuery) (*payment.Payment, error)
	Create(ctx context.Context, payment *command.CreatePaymentCommand) (*payment.Payment, error)
}

type paymentService struct {
	repo        payment.Repository
	userAdapter useracl.Service
}

func NewPaymentService(repo payment.Repository, userACL useracl.Service) Service {
	return &paymentService{
		repo:        repo,
		userAdapter: userACL,
	}
}

func (s *paymentService) Find(ctx context.Context, query *query.FindPaymentQuery) (*payment.Payment, error) {
	payment, err := s.repo.FindByID(payment.PaymentID(query.PaymentID))
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *paymentService) Create(ctx context.Context, cmd *command.CreatePaymentCommand) (*payment.Payment, error) {
	p := payment.NewPayment(
		payment.PaymentID(cmd.PaymentID),
		payment.Amount(cmd.Amount),
		payment.Quantity(cmd.Quantity),
		payment.RefCustomerID(cmd.CustomerID),
	)

	if payment, err := s.repo.Save(p); err != nil {
		return nil, err
	} else {
		return payment, nil
	}
}
