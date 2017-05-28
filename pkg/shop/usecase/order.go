/*
Order use cases. Orchestrates the different domain models and business logic. Controls the flow.
*/

package usecase

//OMIT
import ( //OMIT
	"github.com/maleck13/onions/pkg/shop"          //OMIT
	"github.com/maleck13/onions/pkg/shop/customer" //OMIT
	"github.com/maleck13/onions/pkg/shop/order"    //OMIT
) //OMIT

type Orders struct {
	customers *customer.Service
	orders    *order.Service
}

func (o *Orders) Place(order *shop.Order) error {
	cust, _ := o.customers.Find(order.CustomerRef, order.Location)
	if err := o.orders.Process(cust, order); err != nil {
		return err
	}
}
