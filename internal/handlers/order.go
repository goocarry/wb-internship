package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/goocarry/wb-internship/internal/model"
	"github.com/goocarry/wb-internship/internal/store"
)

// Order ...
type Order struct {
	logger *log.Logger
	store  *store.Store
}

// NewOrder ...
func NewOrder(l *log.Logger, s *store.Store) *Order {
	return &Order{
		logger: l,
		store:  s,
	}
}

// CreateOrder ...
func (o *Order) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	o.logger.Println("info-b23e7145: creating new order")

	order := &model.Order{}
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		o.logger.Printf("err-44344769: can't parse data: %s", err)
		http.Error(rw, "err-44344769: can't parse data", http.StatusBadRequest)
		return
	}

	o.logger.Printf("info-b9e75800: order data: %v", order)
	_, err := o.store.Order().Create(order)
	if err != nil {
		o.logger.Printf("err-b9e75800: can't create order: %s", err)
		http.Error(rw, "err-b9e75800: can't create order", http.StatusBadRequest)
		return
	}

	o.logger.Println("info-f8a835a5: new order created")
	fmt.Fprint(rw, "new order created")
}