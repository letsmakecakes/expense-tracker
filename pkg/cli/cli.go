package cli

import (
	"expense-tracker/internal/expenses"
	"fmt"
	"strconv"
)

func HandleCommand(args []string) {
	switch args[0] {
	case "add":
		if len(args) < 5 {
			fmt.Print("please provide the complete expense input, for e.g; add --description \"Lunch\" --amount 20")
			return
		}

		description := args[2]
		amount, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Printf("error parsing amount: %v", err)
		}
		expense, err := expenses.AddExpense(description, amount)
		if err != nil {
			fmt.Printf("error adding expense: %v", err)
		} else {
			fmt.Printf("expense added successfully (ID: %d)\n", expense.ID)
		}
	case "list":
		ListAllExpenses()
	}
}

func ListAllExpenses() {
	expenseList, err := expenses.LoadExpenses()
	if err != nil {
		fmt.Printf("error loading tasks: %v", err)
		return
	}

	fmt.Println("ID  Date       Description  Amount")

	for _, expense := range expenseList {
		fmt.Printf("%d   %v  %s        $%d", expense.ID, expense.Date, expense.Description, expense.Amount)
	}
}
