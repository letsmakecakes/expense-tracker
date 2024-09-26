package expenses

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const expenseFilePath = "../../data/expenses.json"

func LoadExpenses() ([]Expense, error) {
	if _, err := os.Stat(expenseFilePath); os.IsNotExist(err) {
		return []Expense{}, nil
	}

	data, err := os.ReadFile(expenseFilePath)
	if err != nil {
		return nil, err
	}

	var expenses []Expense
	if err = json.Unmarshal(data, &expenses); err != nil {
		return nil, err
	}

	return expenses, nil
}

func SaveExpenses(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(expenseFilePath, data, 0644)
}

func AddExpense(description string, amount int) (Expense, error) {
	expenses, err := LoadExpenses()
	if err != nil {
		return Expense{}, err
	}

	if amount < 0 {
		return Expense{}, fmt.Errorf("entered an invalid negative amount")
	}

	if len(description) == 0 {
		return Expense{}, fmt.Errorf("description is empty")
	}

	newExpense := Expense{
		ID:          len(expenses) + 1,
		Date:        time.Now().Format("2006-01-02"),
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, newExpense)
	if err := SaveExpenses(expenses); err != nil {
		return Expense{}, err
	}

	return newExpense, nil
}

func DeleteExpense(id int) error {
	expenses, err := LoadExpenses()
	if err != nil {
		return err
	}

	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			return SaveExpenses(expenses)
		}
	}

	return errors.New("expense not found")
}
