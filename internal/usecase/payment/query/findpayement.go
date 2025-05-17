package query

type FindPaymentQuery struct {
	PaymentID  string `json:"payment_id"`
	CustomerID string `json:"customer_id"`
}

func NewFindPaymentQuery(paymentID, customerID string) *FindPaymentQuery {
	return &FindPaymentQuery{
		PaymentID:  paymentID,
		CustomerID: customerID,
	}
}
