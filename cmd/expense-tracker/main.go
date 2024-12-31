package main

import (
	"expense-tracker/internal/cli"
	"expense-tracker/internal/repository/file"
	"expense-tracker/internal/service"
	"log"
	"os"
)

func main() {
	// Get the file path from the environment variable.
	filePath := os.Getenv("EXPENSES_FILE_PATH")
	if filePath == "" {
		// Default to a specific path if the environment variable is not set.
		filePath = "./data/expenses.json"
	}

	// Initialize the file repository with the specified filename.
	repo := file.NewRepository(filePath)

	// Create a new instance of ExpenseService with the repository.
	svc := service.NewExpenseService(repo)

	// Create a new instance of CLI with the service.
	app := cli.NewCLI(svc)

	// Run the CLI application with the provided command-line arguments.
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
