package file

import (
	"encoding/json"
	"expense-tracker/internal/models"
	"expense-tracker/pkg/utils"
	"fmt"
	"log"
	"os"
	"sync"
)

// Repository handles CRUD operations for expenses stored in a file.
type Repository struct {
	filename string
	mutex    sync.RWMutex
}

// storage represents the structure of the data stored in the file.
type storage struct {
	LastID   int              `json:"last_id"`
	Expenses []models.Expense `json:"expenses"`
}

// NewRepository creates a new instance of the Repository.
func NewRepository(filename string) *Repository {
	return &Repository{
		filename: filename,
	}
}

// readStorage reads the storage data from the file.
func (r *Repository) readStorage() (*storage, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	data, err := os.ReadFile(r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &storage{}, nil
		}
		return nil, err
	}

	var s storage
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	return &s, err
}

// writeStorage writes the storage data to the file.
func (r *Repository) writeStorage(s *storage) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	data, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filename, data, 0644)
}

// Add adds a new expense to the storage.
func (r *Repository) Add(expense *models.Expense) error {
	s, err := r.readStorage()
	if err != nil {
		return err
	}

	s.LastID++
	expense.ID = s.LastID
	s.Expenses = append(s.Expenses, *expense)

	return r.writeStorage(s)
}

// Update updates an existing expense in the storage.
func (r *Repository) Update(expense *models.Expense) error {
	s, err := r.readStorage()
	if err != nil {
		return err
	}

	for i, e := range s.Expenses {
		if e.ID == expense.ID {
			s.Expenses[i] = *expense
			return r.writeStorage(s)
		}
	}

	return fmt.Errorf("expense not found")
}

// Delete removes the expense from the storage by its ID.
func (r *Repository) Delete(id int) error {
	s, err := r.readStorage()
	if err != nil {
		return err
	}

	for i, e := range s.Expenses {
		if e.ID == id {
			s.Expenses = append(s.Expenses[:i], s.Expenses[i+1:]...)
			return r.writeStorage(s)
		}
	}

	return fmt.Errorf("expense not found")
}

// GetAll retrieves all expenses from the storage.
func (r *Repository) GetAll() ([]models.Expense, error) {
	s, err := r.readStorage()
	if err != nil {
		return nil, err
	}
	return s.Expenses, nil
}

// GetByID retrieves an expense by its ID from the storage.
func (r *Repository) GetByID(id int) (*models.Expense, error) {
	s, err := r.readStorage()
	if err != nil {
		return nil, err
	}

	for _, e := range s.Expenses {
		if e.ID == id {
			return &e, nil
		}
	}

	return nil, fmt.Errorf("expense not found")
}

// GetByMonth retrieves expense for a specific year and month from the storage.
func (r *Repository) GetByMonth(year, month int) ([]models.Expense, error) {
	s, err := r.readStorage()
	if err != nil {
		return nil, err
	}

	var expenses []models.Expense
	for _, e := range s.Expenses {
		date, err := utils.ParseDate(e.Date)
		if err != nil {
			log.Fatalf("error parsing date: %v", err)
		}
		if date.Year() == year && int(date.Month()) == month {
			expenses = append(expenses, e)
		}
	}

	return expenses, nil
}
