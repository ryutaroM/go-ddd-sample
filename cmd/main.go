package main

import (
	handler "ddd-structure/api/http/handler/payment"
	"ddd-structure/internal/acl/payment/user"
	"ddd-structure/internal/infrastructure/persistence"
	"ddd-structure/internal/usecase/payment"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	paymentRepo := persistence.NewPaymentRepository()

	userAcl := user.NewAdapter()

	paymentService := payment.NewPaymentService(paymentRepo, userAcl)

	paymentHandler := handler.NewPaymentHandler(paymentService)

	r := chi.NewRouter()
	paymentHandler.Register(r)
	fmt.Print("Server is running on port 9000...\n")
	http.ListenAndServe(":9000", r)
}
