package user

import (
	"ddd-structure/internal/domain/payment"
	"ddd-structure/internal/domain/user"
)

type Service interface {
	ToRefUserID(userID user.UserID) (payment.RefUserID, error)
}
