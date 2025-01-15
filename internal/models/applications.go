package models

import (
	"database/sql"
	"time"
)

type Application struct {
	ID              int
	Name            string
	CompanyName     string
	Position        string
	Status          string
	ApplicationDate time.Time
	InterviewDate   *time.Time // Use a pointer for nullable dates
	Notes           string
	CreatedAt       time.Time
}

type ApplicationModel struct {
	DB *sql.DB
}

// Insert inserts a new application into the DB and returns the inserted ID.
func (m *ApplicationModel) Insert(name string, companyName string, position string, status string, applicationDate time.Time, interviewDate *time.Time, notes string) (int, error) {
	stmt := `
		INSERT INTO applications (name, company_name, position, status, application_date, interview_date, notes)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	result, err := m.DB.Exec(stmt, name, companyName, position, status, applicationDate, interviewDate, notes)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Return specific application based on ID
func (m *ApplicationModel) Get(id int) (*Application, error) {
	return nil, nil
}

// Return 10 most recently created applications
func (m *ApplicationModel) Latest() ([]*Application, error) {
	return nil, nil
}
