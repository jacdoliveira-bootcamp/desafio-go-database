package repository

import (
	"database/sql"
	"fmt"
	"github.com/jacdoliveira/bw7/desafio-go-database/internal"
)

// NewSalesMySQL creates new mysql repository for sale.go entity.
func NewSalesMySQL(db *sql.DB) *SalesMySQL {
	return &SalesMySQL{db}
}

// SalesMySQL is the MySQL repository implementation for sale.go entity.
type SalesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all sales from the database.
func (r *SalesMySQL) FindAll() (s []internal.Sale, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `quantity`, `product_id`, `invoice_id` FROM sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var sa internal.Sale
		// scan the row into the sale.go
		err := rows.Scan(&sa.Id, &sa.Quantity, &sa.ProductId, &sa.InvoiceId)
		if err != nil {
			return nil, err
		}
		// append the sale.go to the slice
		s = append(s, sa)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the sale.go into the database.
func (r *SalesMySQL) Save(s *internal.Sale) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO sales (`quantity`, `product_id`, `invoice_id`) VALUES (?, ?, ?)",
		(*s).Quantity, (*s).ProductId, (*s).InvoiceId,
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
	(*s).Id = int(id)

	return
}

func (r *SalesMySQL) GetTopProducts() ([]internal.SaleTopProducts, error) {
	const query = `
		SELECT 
			p.description, 
			SUM(s.quantity) AS total  
		FROM sales s
		INNER JOIN products p ON p.id = s.product_id
		GROUP BY product_id
		ORDER BY total DESC
		LIMIT 5
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query top 5 products: %w", err)
	}
	defer rows.Close()

	var products []internal.SaleTopProducts

	for rows.Next() {
		var prod internal.SaleTopProducts
		if err := rows.Scan(&prod.Description, &prod.Total); err != nil {
			return nil, fmt.Errorf("failed to scan top 5 product row: %w", err)
		}
		products = append(products, prod)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after reading top 5 product rows: %w", err)
	}

	return products, nil
}
