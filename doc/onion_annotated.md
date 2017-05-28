onion/  
├── cli  # User Interface
├── database # Infrastucture concrete implemenation of Domain Services interfaces
│   ├── orderRepository.go  
├── shop # Business Domains
│   ├── orders 
│   │   └── service.go # Domain Service / Business Logic
│   ├── order.go # Domain Model
│   ├── customer.go # Domain Model
│   ├── interfaces.go # Domain Service (interface definitions)
│   ├── order.go # Domain Model
|   ├── usecase
│   │   └── placeOrder.go # Application Service. Orchestrates a particular use case.
└── web # User Interface
    ├── catalog.go
    ├── catalog_test.go 
    └── order.go