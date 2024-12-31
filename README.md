# Expense Tracker CLI

A command-line expense tracking application written in Go that helps you manage and analyze your personal expenses.

## Features

- Add expenses with description, amount, and category
- Update existing expenses
- Delete expenses
- View all expenses in a formatted list
- Get expense summaries (total and monthly)
- JSON-based persistent storage
- Concurrent-safe operations

## Installation

### Prerequisites

- Go 1.20 or higher
- Git

### Building from Source

1. Clone the repository
```bash
git clone https://github.com/yourusername/expense-tracker.git
cd expense-tracker
```

2. Build the application
```bash
go build -o expense-tracker cmd/expense-tracker/main.go
```

3. Set the data file path
```bash
export EXPENSES_FILE_PATH=/path/to/custom/expenses.json
```

4. (Optional) Add to PATH
```bash
sudo mv expense-tracker /usr/local/bin/
```

## Usage

### Adding an Expense
```bash
expense-tracker add --description "Lunch" --amount 20 --category "food"
# Expense added successfully (ID: 1)
```

### Listing All Expenses
```bash
expense-tracker list
# ID  Date       Description  Amount  Category
# 1   2024-08-06  Lunch       $20.00  food
# 2   2024-08-06  Dinner      $30.00  food
```

### Getting a Summary
```bash
# Total summary
expense-tracker summary
# Total expenses: $50.00

# Monthly summary
expense-tracker summary --month 2024-08
# Total expenses for August: $50.00
```

### Updating an Expense
```bash
expense-tracker update --id 1 --description "Late Lunch" --amount 25
# Expense updated successfully
```

### Deleting an Expense
```bash
expense-tracker delete --id 1
# Expense deleted successfully
```

## Command Reference

### Add Command
```bash
expense-tracker add --description <desc> --amount <amount> [--category <category>]
```
- `--description`: Description of the expense (required)
- `--amount`: Amount spent (required, must be positive)
- `--category`: Category of expense (optional, defaults to "misc")

### List Command
```bash
expense-tracker list
```
Shows all expenses in a tabulated format.

### Update Command
```bash
expense-tracker update --id <id> [--description <desc>] [--amount <amount>] [--category <category>]
```
- `--id`: ID of the expense to update (required)
- `--description`: New description (optional)
- `--amount`: New amount (optional)
- `--category`: New category (optional)

### Delete Command
```bash
expense-tracker delete --id <id>
```
- `--id`: ID of the expense to delete (required)

### Summary Command
```bash
expense-tracker summary [--month <month>]
```
- `--month`: Month number (1-12) for monthly summary (optional)

## Project Structure

```
.
├── cmd
│   └── expense-tracker
│       └── main.go
├── internal
│   ├── cli
│   │   └── cli.go
│   ├── models
│   │   └── expense.go
│   ├── repository
│   │   ├── interface.go
│   │   └── file
│   │       └── repository.go
│   └── service
│       └── expense_service.go
├── pkg
│   └── utils
│       └── date.go
├── go.mod
└── README.md
```

## Data Storage

Expenses are stored in a JSON file (`expenses.json`) in the current directory. The file is created automatically when the first expense is added. The storage is concurrent-safe and can handle multiple operations simultaneously.

## Error Handling

The application includes comprehensive error handling for:
- Invalid input values (negative amounts, invalid dates)
- Non-existent expense IDs
- File operations
- Data validation

## Development

### Running Tests
```bash
go test ./...
```

### Adding New Features

1. Fork the repository
2. Create a feature branch
3. Implement your changes
4. Add tests
5. Submit a pull request

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Thanks to the Go community for the excellent standard library
- Inspired by various CLI applications and expense tracking tools