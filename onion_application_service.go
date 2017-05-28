/*
	Application Services provide core application domain logic
	They depend on the interfaces defined by domain service layer and domain model
*/
package order 

import(
	"github.com/maleck13/onion/pkg/shop"
)
type Service struct{
	OrderRepo shop.OrderRepository
}

func (s *Service)Process(ID string)error{
	s.OrderRepo.FindByID(ID)
	//do business logic
	return nil 
}