package database

import "github.com/maleck13/onions/pkg/shop"

// OrderRepo handles orders in the database. Implementor of the shop.OrderRepo interface
type OrderRepo struct{}

func (or *OrderRepo) FindByID(id string) (*shop.Order, error) {
	return &shop.Order{}, nil
}
func (or *OrderRepo) Save(*Order) error {
	return nil
}
