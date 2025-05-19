package query

type FindPaymentQuery struct {
	PaymentID string `json:"payment_id"`
	UserID    string `json:"user_id"`
}

func NewFindPaymentQuery(paymentID, userID string) *FindPaymentQuery {
	return &FindPaymentQuery{
		PaymentID: paymentID,
		UserID:    userID,
	}
}
