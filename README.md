# Expense Tracker (CLI)

A command-line expense tracker built in Go, backed by SQLite. Add, view, filter, and delete expenses — all data persists across runs.

## Features

- **Add an expense** — records name, amount, category, and date (defaults to today if left blank)
- **View all expenses** — lists every recorded expense
- **View expenses by category** — filter by category (case-insensitive)
- **View expenses by date range** — filter by a start and end date (`YYYY-MM-DD`)
- **Delete an expense** — remove an entry by its ID
- **Persistent storage** — all data is stored in a local SQLite database (`expenses.db`), so nothing is lost between runs
- **Looping menu** — the app returns to the main menu after each action instead of exiting

## Tech Stack

- **Language:** Go
- **Database:** SQLite (via [`modernc.org/sqlite`](https://pkg.go.dev/modernc.org/sqlite), a pure-Go driver — no CGo required)

## Project Structure

```
expense-tracker/
├── main.go        # Application logic and CLI menu
├── schema.sql      # Database schema, applied on startup
├── go.mod
├── go.sum
└── expenses.db     # SQLite database (created/used at runtime)
```

## Getting Started

### Prerequisites

- Go 1.20+ installed

### Run

```bash
git clone <your-repo-url>
cd expense-tracker
go run main.go
```

The database schema is applied automatically on startup — no manual setup needed.

## Usage

On launch, you'll see:

```
Welcome to Kartike's Expense Tracker
1. Add an Expense
2. View Expenses
3. Delete an Expense
4. Exit
```

- **Add an Expense** — prompts for name, amount, category, and date
- **View Expenses** — opens a submenu to view all expenses, filter by category, or filter by date range
- **Delete an Expense** — prompts for the expense ID to remove
- **Exit** — closes the application

## Database Schema

```sql
create table if not exists expenses (
    id integer primary key autoincrement,
    name string not null,
    amount real not null,
    category string not null,
    date text not null default current_timestamp
);
```

## Roadmap / Known Limitations

- No validation currently prevents negative expense amounts
- No strict date-format validation on input
- All logic currently lives in a single file; a layered structure (storage/service/CLI) is planned

## License

MIT