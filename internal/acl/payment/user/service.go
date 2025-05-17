package user

import (
	"ddd-structure/internal/domain/user"
	"ddd-structure/internal/domain/payment"
)

type Service interface {
	ToRefCustomerID(userID user.UserID) (payment.RefCustomerID, error)
}
