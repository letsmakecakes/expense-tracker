package repository

import "expense-tracker/internal/models"

// ExpenseRepository defines the interface for expense data operations.
type ExpenseRepository interface {
	// Add inserts a new expense record.
	Add(expense *models.Expense) error

	// Update modifies an existing expense record.
	Update(expense *models.Expense) error

	// Delete removes an expense record by its ID.
	Delete(id int) error

	// GetAll retrieves all expense records.
	GetAll() ([]models.Expense, error)

	// GetByID retrieves all expense records.
	GetByID(id int) (*models.Expense, error)

	// GetByMonth retrieves expense records for a specific year and month.
	GetByMonth(year, month int) ([]models.Expense, error)
}
