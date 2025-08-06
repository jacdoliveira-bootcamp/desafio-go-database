package repository

import (
	"database/sql"
	"fmt"
	"github.com/jacdoliveira/bw7/desafio-go-database/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*c).Id = int(id)

	return
}

func (r *CustomersMySQL) GetTotalByCondition() ([]internal.CustomerGetTotal, error) {
	const query = `SELECT 
    	CASE c.condition
        	WHEN 1 THEN 'Activo ( 1 )'
     		ELSE 'Inactivo ( 0 )'
    	END AS ConditionStatus,
    	ROUND(SUM(i.total), 2) AS Total
	FROM invoices i
	JOIN customers c ON c.id = i.customer_id
	GROUP BY c.condition;
`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query GetTotalByCondition: %w", err)
	}
	defer rows.Close()

	var totals []internal.CustomerGetTotal
	for rows.Next() {
		var totalCustomer internal.CustomerGetTotal
		if scanErr := rows.Scan(&totalCustomer.Condition, &totalCustomer.Total); scanErr != nil {
			return nil, fmt.Errorf("error scanning row in GetTotalByCondition: %w", scanErr)
		}
		totals = append(totals, totalCustomer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error in GetTotalByCondition: %w", err)
	}

	return totals, nil
}

func (r *CustomersMySQL) GetTopActive() ([]internal.CustomerTopActive, error) {
	const query = `
		SELECT c.first_name, c.last_name, ROUND(SUM(i.total), 2) AS total
		FROM invoices i
		JOIN customers c ON i.customer_id = c.id
		GROUP BY customer_id
		ORDER BY total DESC
		LIMIT 5;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query for most active customers: %w", err)
	}
	defer rows.Close()

	var customers []internal.CustomerTopActive

	for rows.Next() {
		var customer internal.CustomerTopActive
		if scanErr := rows.Scan(&customer.FirstName, &customer.LastName, &customer.Amount); scanErr != nil {
			return nil, fmt.Errorf("failed to scan customer row: %w", scanErr)
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %w", err)
	}

	return customers, nil
}
