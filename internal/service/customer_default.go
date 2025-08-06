package service

import "github.com/jacdoliveira/bw7/desafio-go-database/internal"

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp internal.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}

func (s *CustomersDefault) GetTotalByCondition() (d []internal.CustomerGetTotal, err error) {
	d, err = s.rp.GetTotalByCondition()
	return
}

func (s *CustomersDefault) GetTopActive() (d []internal.CustomerTopActive, err error) {
	d, err = s.rp.GetTopActive()
	return
}
