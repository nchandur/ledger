# Infinity Ledger

Infinity Ledger is a backend web application built with the Gin framework in Go. It simplifies group debt tracking and expense sharing by allowing users to create groups, log shared expenses, and automatically calculate who owes whom.

## Features

- Create and manage expense groups
- Add, update, and delete expenses in a group
- View all expense groups
- Get detailed ledgers for each group
- Automatically calculate balances per member
- Generate optimal transactions to settle debts within a group

## API Endpoints

### Group Management
- `POST /groups/create` — Create a new group
- `GET /groups/view` — View all existing groups
- `DELETE /groups/delete` — Delete a group

### Expense Ledger
- `POST /ledger/add` — Add an expense to a group
- `GET /ledger/view` — View all expenses for a group
- `PATCH /ledger/update` — Update an existing expense
- `DELETE /ledger/delete` — Delete an expense

### Balances & Settlements
- `GET /ledger/balances` — Get net balance for each person (owed/lent)
- `GET /ledger/transactions` — Get simplified transactions to settle balances

### Misc
- `GET /` — Landing route (health check or welcome page)

## Tech Stack

- **Backend:** Go (Gin framework)
- **Database:** MongoDB
- **Frontend:** None (API-only application)

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.22 or higher
- [MongoDB](https://www.mongodb.com/try/download/community)

### Running Locally

```
go run cmd/server.go
```