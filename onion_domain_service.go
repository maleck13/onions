/*
	Often the Domain Service layer defines the business logics required interfaces
*/
package shop

import (
	"github.com/maleck13/onion/pkg/shop"
)

type OrderRepository interface {
	FindById(id string) (*shop.Order, error)
}
