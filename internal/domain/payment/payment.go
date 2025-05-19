package payment

import "fmt"

type PaymentID string

func (p PaymentID) String() string {
	return string(p)
}

func (p PaymentID) validate() error {
	if len(p) == 0 {
		return fmt.Errorf("payment id is empty")
	}
	return nil
}

type Amount float64

func (a Amount) validate() error {
	if a <= 0 {
		return fmt.Errorf("amount is invalid")
	}
	return nil
}

type Quantity int

func (q Quantity) validate() error {
	if q <= 0 {
		return fmt.Errorf("quantity is invalid")
	}
	return nil
}

type RefUserID string
type Payment struct {
	ID        PaymentID
	Amount    Amount
	Quantity  Quantity
	RefUserID RefUserID
}

func (p *Payment) validate() error {
	if err := p.ID.validate(); err != nil {
		return fmt.Errorf("payment id is invalid: %w", err)
	}
	if err := p.Amount.validate(); err != nil {
		return fmt.Errorf("amount is invalid: %w", err)
	}
	if err := p.Quantity.validate(); err != nil {
		return fmt.Errorf("quantity is invalid: %w", err)
	}
	return nil
}

func NewPayment(id PaymentID, amount Amount, quantity Quantity, refUserID RefUserID) *Payment {
	return &Payment{
		ID:        id,
		Amount:    amount,
		Quantity:  quantity,
		RefUserID: refUserID,
	}
}
