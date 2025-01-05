package models

import (
	"database/sql"
	"time"
)

type Application struct {
	ID          int
	Name        string
	CompanyName string
	Position    string
	CreatedAt   time.Time
}

type ApplicationModel struct {
	DB *sql.DB
}

// Insert new application to DB
func (m *ApplicationModel) Insert(name string, companyName string, position int) (int, error) {
	return 0, nil
}

// Return specific application based on ID
func (m *ApplicationModel) Get(id int) (*Application, error) {
	return nil, nil
}

// Return 10 most recently created applications
func (m *ApplicationModel) Latest() ([]*Application, error) {
	return nil, nil
}
