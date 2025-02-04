# Expense-Tracker-CLI


## Overview
The Expense Tracker CLI is a command-line tool that helps users manage their expenses efficiently. Users can add, update, delete, and view expenses, as well as generate summaries for better financial tracking.

## Features
- Add an expense with a description and amount.
- Update an existing expense.
- Delete an expense.
- View all recorded expenses.
- Get a summary of all expenses.
- Get a summary of expenses for a specific month (of the current year).

[.](https://roadmap.sh/projects/expense-tracker)

## Installation
1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd expense-tracker-cli
   ```
2. Build the project:
   ```sh
   go build -o expense-tracker
   ```
3. Run the CLI tool:
   ```sh
   ./expense-tracker
   ```

## Usage
### Add an Expense
```sh
./expense-tracker add -d "Lunch with friends" -a 15.50
```

### Update an Expense
```sh
./expense-tracker update -id 1 -d "Lunch with colleagues" -a 20.00
```

### Delete an Expense
```sh
./expense-tracker delete -id 1
```

### View All Expenses
```sh
./expense-tracker list
```

### View Summary of All Expenses
```sh
./expense-tracker summary
```

### View Summary for a Specific Month
```sh
./expense-tracker summary -m JAnuary # For January
```


