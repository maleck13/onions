package shop

//OrderRepo defines the interface for crudl actions on the Order domain object //OMIT
type OrderRepo interface {
	FindByID(id string) (*Order, error)
	Save(*Order) error
}

//PaymentGateway defines the interface for working with external payment gateways //OMIT
type PaymentGateway interface { //OMIT
	AcceptPayment(amount int) error //OMIT
} //OMIT
