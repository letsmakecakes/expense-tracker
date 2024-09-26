# Expense Tracker CLI

A simple command-line application built in Go to manage your personal finances. The Expense Tracker CLI allows you to add, update, delete, and view expenses, along with providing a summary of your expenses. It also supports tracking monthly budgets (optional).

## Features

- **Add Expense**: Add a new expense with a description and amount.
- **Update Expense**: Modify an existing expense's description or amount.
- **Delete Expense**: Remove an expense by its ID.
- **List Expenses**: View all recorded expenses.
- **Expense Summary**: Get the total of all your expenses.
- **Monthly Summary**: View a summary of expenses for a specific month (of the current year).
- **Budget Tracking (Optional)**: Set a monthly budget and get a warning if exceeded.

## Commands and Usage

Here’s how to use the Expense Tracker CLI:

### Add a New Expense
```bash
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)
```

### List All Expenses
```bash
$ expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10
```

### Update an Existing Expense
```bash
$ expense-tracker update --id 1 --description "Brunch" --amount 25
# Expense updated successfully (ID: 1)
```

### Delete an Expense
```bash
$ expense-tracker delete --id 1
# Expense deleted successfully (ID: 1)
```

### View Total Summary of Expenses
```bash
$ expense-tracker summary
# Total expenses: $30
```

### View Summary for a Specific Month
```bash
$ expense-tracker summary --month 8
# Total expenses for August: $20
```

## Installation

### Prerequisites
- **Go**: Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

### Steps to Install and Run
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/expense-tracker-cli.git
   cd expense-tracker-cli
   ```

2. Build the application:
   ```bash
   go build -o expense-tracker ./cmd
   ```

3. Run the Expense Tracker CLI:
   ```bash
   ./expense-tracker
   ```

## File Storage

The expense data is stored locally in a JSON file (`data/expenses.json`). You can customize the file path by modifying the relevant section in the code.

## Future Enhancements

- Allow filtering expenses by categories (e.g., food, travel, etc.).
- Add a feature to export expenses to a CSV file.
- Improve user interface for better usability.
- Implement database storage (e.g., SQLite) for larger datasets.
