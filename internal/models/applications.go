package models

import (
	"database/sql"
	"errors"
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

	stmt := `SELECT id, name, company_name, position, status, application_date, interview_date, notes FROM applications
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Application struct.
	a := &Application{}

	// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Application struct.
	err := row.Scan(&a.ID, &a.Name, &a.CompanyName, &a.Position, &a.Status, &a.ApplicationDate, &a.InterviewDate, &a.Notes)

	if err != nil {
		// If the query returns no rows, then row.Scan() will return a
		// sql.ErrNoRows error. We use the errors.Is() function check for that
		// error specifically, and return our own ErrNoRecord error
		// instead (we'll create this in a moment).
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	// If everything went OK then return the Application object.
	return a, nil
}

// Return 10 most recently created applications
func (m *ApplicationModel) Latest() ([]*Application, error) {
	return nil, nil
}
