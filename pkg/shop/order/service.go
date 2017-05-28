package order

//depends on inner domain model layer uses interfaces to define its dependencies //OMIT
import "github.com/maleck13/onions/pkg/shop"

type Service struct {
	orderRepo shop.OrderRepo //dependency defined as interface
}

func (s *Service) Process(c *shop.Customer, o *shop.Order) error {
	//obviously there would be more logic than this check order fullfilment etc
	return s.orderRepo.Save(o)
}
