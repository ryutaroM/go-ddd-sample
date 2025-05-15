package payment

type Repository interface {
	FindByID(id PaymentID) (*Payment, error)
	Save(payment *Payment) (*Payment, error)
}
