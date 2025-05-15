package handler

import (
	request "ddd-structure/api/http/request/payment"
	"ddd-structure/api/http/response"
	paymentdomain "ddd-structure/internal/domain/payment"
	"ddd-structure/internal/usecase/payment"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PaymentHandler struct {
	service payment.Service
}

func NewPaymentHandler(service payment.Service) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) Register(r chi.Router) {
	r.Route("/payments", func(r chi.Router) {
		r.Get("/{id}", h.FindPayment)
		r.Post("/", h.CreatePayment)
	})
}

func (h *PaymentHandler) FindPayment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	payment, err := h.service.Find(r.Context(), paymentdomain.PaymentID(id))
	if err != nil {
		response.WriteError(w, http.StatusNotFound, err)
		return
	}

	response.WriteResponse(w, http.StatusOK, payment)
}

func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewCreatePaymentRequest(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment := req.ToDomain()

	if payment, err := h.service.Create(r.Context(), payment); err != nil {
		response.WriteError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.WriteResponse(w, http.StatusCreated, payment)
	}
}
