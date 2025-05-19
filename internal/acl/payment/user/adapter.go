package user

import (
	"ddd-structure/internal/domain/payment"
	"ddd-structure/internal/domain/user"
)

type Adapter struct{}

func NewAdapter() Service {
	return &Adapter{}
}

func (a *Adapter) ToRefUserID(userID user.UserID) (payment.RefUserID, error) {
	ref := payment.RefUserID(userID.String())
	return ref, nil
}
