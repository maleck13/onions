package shop

import ( //OMIT
	"time" //OMIT
) //OMIT

type Order struct {
	CustomerRef string
	Location    string
	Ordered     time.Time
	Dispatched  time.Time
}

type Onion struct {
	Species   string
	Origin    string
	Harvested time.Time
}

type Customer struct{}
