package cli

import (
	"expense-tracker/internal/expenses"
	"fmt"
	"strconv"
	"strings"
)

func HandleCommand(args []string) {
	switch args[0] {
	case "add":
		if len(args) < 5 {
			fmt.Print("please provide the complete expense input, usage: add --description <description> --amount <amount>")
			return
		}

		description := args[2]
		amount, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Printf("error parsing amount: %v\n", err)
			return
		}
		expense, err := expenses.AddExpense(description, amount)
		if err != nil {
			fmt.Printf("error adding expense: %v\n", err)
		} else {
			fmt.Printf("expense added successfully (ID: %d)\n", expense.ID)
		}
	case "list":
		ListAllExpenses()

	case "delete":
		if len(args) < 3 {
			fmt.Println("usage: delete --id <id>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("invalid ID: %v\n", err)
			return
		}

		if err := expenses.DeleteExpense(id); err != nil {
			fmt.Printf("error deleting expense: %v\n", err)
		} else {
			fmt.Println("expense deleted successfully")
		}
	case "summary":
		if len(args) == 3 {
			filterMonth, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("failed to parse month: %v\n", err)
				return
			}

			ExpenseSummaryByMonth(filterMonth)
		} else {
			ExpenseSummary()
		}
	}
}

func ListAllExpenses() {
	expenseList, err := expenses.LoadExpenses()
	if err != nil {
		fmt.Printf("error loading expenses: %v\n", err)
		return
	}

	fmt.Println("ID  Date       Description  Amount")

	for _, expense := range expenseList {
		fmt.Printf("%d   %v  %s        $%d", expense.ID, expense.Date, expense.Description, expense.Amount)
	}
}

func ExpenseSummary() {
	expenseList, err := expenses.LoadExpenses()
	if err != nil {
		fmt.Printf("error loading expenses: %v\n", err)
		return
	}

	totalAmount := 0
	for _, expense := range expenseList {
		totalAmount += expense.Amount
	}

	fmt.Printf("Total expenses: $%d\n", totalAmount)
}

func ExpenseSummaryByMonth(filterMonth int) {
	expenseList, err := expenses.LoadExpenses()
	if err != nil {
		fmt.Printf("error loading expenses: %v\n", err)
		return
	}

	totalAmount := 0
	for _, expense := range expenseList {
		month, err := strconv.Atoi(strings.Split(expense.Date, "-")[1])
		if err != nil {
			fmt.Printf("error parsing month: %v", err)
			return
		}

		if month == filterMonth {
			totalAmount += expense.Amount
		}
	}

	fmt.Printf("Total expenses for August: $%d", totalAmount)
}
