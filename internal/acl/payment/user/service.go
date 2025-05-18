package user

import (
	"ddd-structure/internal/domain/payment"
	"ddd-structure/internal/domain/user"
)

type Service interface {
	ToCustomerID(userID user.UserID) (payment.CustomerID, error)
}
