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
- [`POST /groups/create`](api/API.md#post-groupscreate) — Create a new group
- [`GET /groups/view`](api/API.md#get-groupsview) — View all existing groups
- [`DELETE /groups/delete`](api/API.md#delete-groupsdelete) — Delete a group

### Expense Ledger
- [`GET /ledger/view`](api/API.md#get-ledgerview) — View all expenses for a group
- [`POST /ledger/add`](api/API.md#post-ledgeradd) — Add an expense to a group
- [`PATCH /ledger/update`](api/API.md#patch-ledgerupdate) — Update an existing expense
- [`DELETE /ledger/delete`](api/API.md#delete-ledgerdelete) — Delete an expense

### Balances & Settlements
- [`GET /ledger/balances`](api/API.md#get-ledgerbalances) — Get net balance for each person (owed/lent)
- [`GET /ledger/transactions`](api/API.md#get-ledgertransactions) — Get simplified transactions to settle balances

### Misc
- [`GET /`](api/API.md#get-) — Landing route (health check or welcome page)

## Tech Stack

- **Backend:** Go (Gin framework)
- **Database:** MongoDB

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.22 or higher
- [MongoDB](https://www.mongodb.com/try/download/community)

### Running Locally

```
go run cmd/server.go
```

The server will start on http://localhost:8080.
