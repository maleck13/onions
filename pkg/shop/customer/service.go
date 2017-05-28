package customer

import "github.com/maleck13/onions/pkg/shop"

type Service struct{}

func (s *Service) Find(id, city string) (*shop.Customer, error) {
	return nil, &shop.Customer{}
}
