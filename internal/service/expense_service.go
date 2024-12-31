package service

import (
	"expense-tracker/internal/models"
	"expense-tracker/internal/repository"
)

// ExpenseService provides methods to manage expenses.
type ExpenseService struct {
	repo repository.ExpenseRepository
}

// NewExpenseService creates a new instance of ExpenseService.
func NewExpenseService(repo repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

// AddExpense adds a new expense and returns its ID.
func (s *ExpenseService) AddExpense(expense *models.Expense) (int, error) {
	if err := s.repo.Add(expense); err != nil {
		return 0, err
	}
	return expense.ID, nil
}

// UpdateExpense updates an existing expense.
func (s *ExpenseService) UpdateExpense(expense *models.Expense) error {
	return s.repo.Update(expense)
}

// DeleteExpense deletes an expense by its ID.
func (s *ExpenseService) DeleteExpense(id int) error {
	return s.repo.Delete(id)
}

// GetExpense retrieves an expense by its ID.
func (s *ExpenseService) GetExpense(id int) (*models.Expense, error) {
	return s.repo.GetByID(id)
}

// GetAllExpense retrieves all expenses.
func (s *ExpenseService) GetAllExpense() ([]models.Expense, error) {
	return s.repo.GetAll()
}

// GetTotalExpenses calculates the total amount of all expenses.
func (s *ExpenseService) GetTotalExpenses() (float64, error) {
	expenses, err := s.repo.GetAll()
	if err != nil {
		return 0, err
	}

	var total float64
	for _, e := range expenses {
		total += e.Amount
	}
	return total, nil
}

// GetMonthlyExpenses calculates the total amount of expenses for a specific month and year.
func (s *ExpenseService) GetMonthlyExpenses(year, month int) (float64, error) {
	expenses, err := s.repo.GetByMonth(year, month)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, e := range expenses {
		total += e.Amount
	}
	return total, nil
}
