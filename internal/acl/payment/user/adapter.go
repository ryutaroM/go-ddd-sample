package user

import (
	"ddd-structure/internal/domain/payment"
	"ddd-structure/internal/domain/user"
)

type Adapter struct{}

func NewAdapter() Service {
	return &Adapter{}
}

func (a *Adapter) ToRefCustomerID(userID user.UserID) (payment.RefCustomerID, error) {
	ref := payment.RefCustomerID(userID.String())
	return ref, nil
}
