# API Usage

## `GET /`

Returns a basic welcome message. This route can be used to verify that the server is running.

### Response

```{json}
{
  "body": "Welcome to the Infinity Ledger",
  "error": null
}
```
- `body` (string) – Welcome message.
- `error` (null or string) – Any error encountered. `null` means success.

## `POST /groups/create`

Creates a new expense group with a fresh ledger.

### Request Body

```{json}
{
  "name": "Grocery",
  "people": ["Alice", "Bob", "Charlie"],
  "currency": "USD"
}
```

- `name` (string, required) – The name of the group.
- `people` (array of strings, required) – Names of the individuals participating in the group.
- `currency` (string, required) – Preferred currency for the ledger (e.g., "USD", "EUR"). All transactions in the group must use this currency.

### Successful Response

```{json}
{
  "body": "group created",
  "error": null
}
```

- `body` (string) – Confirmation message.
- `error` (null) – Indicates no error occurred.

### Error Response

```{json}
{
  "error": "failed to bind JSON: missing required field 'currency'",
  "body": null
}
```

- `error` (string) – Error message from JSON binding or internal logic.
- `body` (null) – No data returned on error.

### Status Codes
- `200 OK` – Group created successfully.
- `400 Bad Request` – Malformed request body or internal error during group creation.

## `GET /groups/view`

Fetches group information from the database. Returns either all groups or details about a specific group if a name is provided via query parameter.

### Query Parameters

- `name` (string, optional) – The name of the group to fetch. If omitted, all groups will be returned.

#### Examples

- `GET /groups/view` -- returns all groups

- `GET /groups/view?name=grocery` -- returns only the specified 
group

### Successful Response

```{json}
{
    "body": {
        "groups": [
            {
                "created_at": "2025-06-01T23:06:44.723Z",
                "group_name": "house",
                "people": [
                    "A",
                    "B",
                    "C",
                    "D",
                    "E"
                ],
                "currency": "USD"
            },
            {
                "created_at": "2025-06-02T00:32:42.584Z",
                "group_name": "grocery",
                "people": [
                    "A",
                    "B",
                    "C",
                    "D",
                    "E"
                ],
                "currency": "USD"
            },
        ]
    },
    "error": null
}
```

- `groups` (array) – List of groups.

### Error Response

```{json}
{
  "error": "failed to query database: connection timeout",
  "body": null
}
```

- `error` (string) – Query failure or other server-side error.
- `body` (null) – No group data returned on failure.

## `DELETE /groups/delete`

Deletes a group and its associated expense ledger.

### Query Parameters

- `name` (string, required) – The name of the group to be deleted.

### Successful Response

```{json}
{
  "error": null,
  "body": "group dropped"
}
```

- `body` (string) – Confirmation that the group and its ledger were successfully deleted.
- `error` (null) – Indicates success.

### Error Response

```{json}
{
  "error": "no group found with given name",
  "body": null
}
```

### Status Codes
- `200 OK` – Group deleted successfully.
- `400 Bad Request` – Group name is missing or no matching group found.
`500 Internal Server Error` – An error occurred while deleting group data.

## `GET /ledger/view`

Fetches all the recorded expenses for a particular group ledger.

### Query Parameters

- (none required) – This route may later be extended with filters, but currently it returns all expenses.

### Successful Response

```{json}
{
  "error": null,
  "body": {
    "expenses": [
      {
        "updated_at": "2025-06-01T23:07:40.987Z",
        "item_id": 1,
        "item": "Milk",
        "price": 3.53,
        "lent": "A",
        "involved": ["B", "C", "E"],
        "type": "equal",
        "splits": null
      },
      {
        "updated_at": "2025-06-01T23:08:30.111Z",
        "item_id": 3,
        "item": "Apples",
        "price": 5.89,
        "lent": "B",
        "involved": ["C", "E"],
        "type": "percentages",
        "splits": [0.25, 0.75]
      }
    ]
  }
}

```

#### Field Breakdown

- `item_id` (integer) – Unique ID of the expense.
- `item` (string) – Description of the item or expense.
- `price` (float) – Total amount spent.
- `lent` (string) – Person who paid for the item.
- `involved` (array of strings) – People who are sharing the expense.
- `type` (string) – How the expense is split. One of: "equal", "manual", "percentages".
- `splits` (array or null) – Split amounts or ratios depending on the type. null for equal split.
- `updated_at` (timestamp) – Last time the expense was modified.

### Error Response

```{json}
{
  "error": "failed to retrieve expenses: database timeout",
  "body": null
}
```

### Status Codes
- `200 OK` – Expenses returned successfully.
- `500 Internal Server Error` – Failed to fetch ledger data from the database.

## `POST /ledger/add`

Adds a new expense to the specified group’s ledger.

### Query Parameters

- `name` (string, required) – The name of the group.

### Request Body

```{json}
{
  "item": "Dinner",
  "price": 60.0,
  "lent": "Alice",
  "involved": ["Bob", "Charlie"],
  "type": "percentages",
  "splits": [0.4, 0.6]
}
```

- `item` (string, required) – Description of the expense.
- `price` (float, required) – Total cost.
- `lent` (string, required) – Person who paid.
- `involved` (array of strings, required) – People sharing the expense.
- `type` (string, required) – Split type: "equal", "percentages", or "manual".
- `splits` (array or null) – Must be null for "equal", else:
    - For "percentages": values must sum to 1.0
    - For "manual": values must sum to the price

### Successful Response

```{json}
{
  "error": null,
  "body": "expense added"
}
```

### Error Response

```{json}
{
  "error": "group not found",
  "body": null
}
```

### Status Codes
- `200 OK` – Expense added successfully.
- `400 Bad Request` – Invalid group or logical error.
- `500 Internal Server Error` – JSON binding or balance calculation failed.

## `PATCH /ledger/update`

Updates an existing expense in a group's ledger, based on the item name.

**Note:** Currently, matching is done by item name alone — avoid duplicate item names in the same group. Future versions will support unique identifiers.

### Query Parameters
- `name` (string, required) – The name of the group.
- `item` (string, required) – The name of the item to update.

### Request Body

Only include the fields you want to update. For example:

```{json}
{
  "price": 5.00,
  "type": "manual",
  "splits": [3.00, 2.00]
}
```

- `item` (string, optional) – Updated item name.
- `price` (float, optional) – New price.
- `lent` (string, optional) – Updated payer.
- `involved` (array of strings, optional) – New people involved.
- `type` (string, optional) – Must be one of "equal", "percentages", "manual".
- `splits` (array or null, optional) – Must match the new type if provided: null for "equal"
    - Sums to 1.0 for "percentages"
    - Sums to new price for "manual"

### Successful Response

```{json}
{
  "error": null,
  "body": "expense updated"
}
```

### Error Response

```{json}
{
  "error": "failed to bind JSON",
  "body": null
}
```

### Status Codes
- `200 OK` – Expense updated successfully.
- `400 Bad Request` – Invalid or missing request data.
- `500 Internal Server Error` – Error during group access or update logic.

## `GET /ledger/balances`

Retrieves the current net balances for each individual in a group. A positive value means the person is owed money; a negative value means they owe money.

### Query Parameters

- `name` (string, required) – The name of the group whose ledger balances you want to view.

### Successful Response

```{json}
{
  "body": {
    "balances": {
      "group": "house",
      "balances": {
        "A": 1.7799999999999998,
        "B": -2.1799999999999997,
        "C": -0.7900000000000005,
        "D": -3.9299999999999997,
        "E": 5.12
      },
      "updated_at": "2025-06-01T23:11:25.325Z"
    }
  },
  "error": null
}
```

#### Field Breakdown

- `group` (string) – Name of the group.
- `balances` (object) – Map of person names to their net balance.
    - Positive values = amount they’re owed.
    - Negative values = amount they owe.
- `updated_at` (timestamp) – When balances were last recalculated.

### Error Response

```{json}
{
  "error": "could not fetch balances: group not found",
  "body": null
}
```

### Status Codes
- `200 OK` – Balances returned successfully.
- `500 Internal Server Error` – Failure in balance calculation or group lookup.

## `GET /ledger/transactions`

Returns the recommended transactions required to settle all debts in a group. Each transaction indicates who should pay whom and how much, based on the current ledger balances.

### Query Parameters
- `name` (string, required) – The name of the group.

### Successful Response

```{json}
{
  "body": {
    "transactions": {
      "_id": "683cdd3d6f2a62b538883788",
      "group": "house",
      "transactions": [
        {
          "amount": 3.9299999999999997,
          "from": "D",
          "to": "E"
        },
        {
          "amount": 1.7799999999999998,
          "from": "B",
          "to": "A"
        },
        {
          "amount": 0.7900000000000005,
          "from": "C",
          "to": "E"
        },
        {
          "amount": 0.3999999999999999,
          "from": "B",
          "to": "E"
        }
      ],
      "updated_at": "2025-06-01T23:11:25.325Z"
    }
  },
  "error": null
}

```
### Field Breakdown

- `group` (string) – Group name.
- `transactions` (array) – List of transactions required to settle up:
    - `from`: person who owes money
    - `to`: person who is owed
    - `amount`: value of transaction
- `updated_at` (timestamp) – Timestamp of last update