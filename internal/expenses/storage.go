package expenses

import (
	"encoding/json"
	"os"
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
