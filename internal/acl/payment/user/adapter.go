package user

import (
	"ddd-structure/internal/domain/payment"
	"ddd-structure/internal/domain/user"
)

type Adapter struct{}

func NewAdapter() Service {
	return &Adapter{}
}

func (a *Adapter) ToCustomerID(userID user.UserID) (payment.CustomerID, error) {
	ref := payment.CustomerID(userID.String())
	return ref, nil
}
