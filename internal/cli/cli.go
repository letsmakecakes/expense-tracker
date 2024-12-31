package cli

import (
	"expense-tracker/internal/models"
	"expense-tracker/internal/service"
	"expense-tracker/pkg/utils"
	"flag"
	"fmt"
	"strconv"
	"time"
)

// CLI handles command line interactions for the expense tracker.
type CLI struct {
	service *service.ExpenseService
}

// NewCLI creates a new instance of CLI.
func NewCLI(svc *service.ExpenseService) *CLI {
	return &CLI{
		service: svc,
	}
}

// Run processes CLI commands and dispatches them to the appropriate handlers.
func (c *CLI) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("command required")
	}

	switch args[1] {
	case "add":
		return c.handleAdd(args[2:])
	case "list":
		return c.handleList()
	case "delete":
		return c.handleDelete(args[2:])
	case "update":
		return c.handleUpdate(args[2:])
	case "summary":
		return c.handleSummary(args[2:])
	default:
		return fmt.Errorf("unknown command: %s", args[1])
	}
}

func (c *CLI) handleAdd(args []string) error {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	desc := addCmd.String("description", "", "Expense description")
	amount := addCmd.Float64("amount", 0, "Expense amount")
	category := addCmd.String("category", "misc", "Expense category")

	if err := addCmd.Parse(args); err != nil {
		return err
	}

	if *desc == "" || *amount <= 0 {
		return fmt.Errorf("description and amount are required")
	}

	expense := &models.Expense{
		Description: *desc,
		Amount:      *amount,
		Category:    *category,
		Date:        utils.FormatDate(time.Now()),
	}

	id, err := c.service.AddExpense(expense)
	if err != nil {
		return err
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", id)
	return nil
}

func (c *CLI) handleList() error {
	expenses, err := c.service.GetAllExpense()
	if err != nil {
		return err
	}

	fmt.Println("ID\tDate\t\tDescription\tAmount\tCategory")
	for _, e := range expenses {
		fmt.Printf("%d\t%s\t%s\t\t$%.2f\t%s\n", e.ID, e.Date, e.Description, e.Amount, e.Category)
	}
	return nil
}

func (c *CLI) handleDelete(args []string) error {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Int("id", 0, "Expense ID to delete")

	if err := deleteCmd.Parse(args); err != nil {
		return err
	}

	if *id <= 0 {
		return fmt.Errorf("valid ID is required")
	}

	if err := c.service.DeleteExpense(*id); err != nil {
		return err
	}

	fmt.Println("Expense deleted successfully")
	return nil
}

func (c *CLI) handleSummary(args []string) error {
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	month := summaryCmd.String("month", "", "Month in YYYY-MM format")

	if err := summaryCmd.Parse(args); err != nil {
		return err
	}

	if *month == "" {
		total, err := c.service.GetTotalExpenses()
		if err != nil {
			return err
		}
		fmt.Printf("Total expenses: $%.2f\n", total)
		return nil
	}

	monthInt, err := strconv.Atoi(*month)
	if err != nil {
		return fmt.Errorf("invalid month format: %s", *month)
	}

	if !utils.IsValidMonth(monthInt) {
		return fmt.Errorf("inavlid month: %s", *month)
	}

	date, err := utils.ParseDate(*month + "-01")
	if err != nil {
		return fmt.Errorf("invalid month format: %s", *month)
	}

	total, err := c.service.GetMonthlyExpenses(date.Year(), int(date.Month()))
	if err != nil {
		return err
	}

	fmt.Printf("Total expenses for %s: $%.2f\n", date.Format("January 2006"), total)
	return nil
}

func (c *CLI) handleUpdate(args []string) error {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	id := updateCmd.Int("id", 0, "Expense ID")
	desc := updateCmd.String("description", "", "New description")
	amount := updateCmd.Float64("amount", 0, "New amount")
	category := updateCmd.String("category", "", "New category")

	if err := updateCmd.Parse(args); err != nil {
		return err
	}

	if *id <= 0 {
		return fmt.Errorf("valid ID is required")
	}

	expense, err := c.service.GetExpense(*id)
	if err != nil {
		return err
	}

	if *desc != "" {
		expense.Description = *desc
	}
	if *amount > 0 {
		expense.Amount = *amount
	}
	if *category != "" {
		expense.Category = *category
	}

	if err := c.service.UpdateExpense(expense); err != nil {
		return err
	}

	fmt.Println("Expense updated successfully")
	return nil
}
