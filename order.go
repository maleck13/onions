/*
	Order is a a Domain Model that represents an order in the shop.
	It has no dependencies because it is at the very core
*/
package shop

type Order struct {
	ID          string
	Items       []string
	CustomerRef string
}
