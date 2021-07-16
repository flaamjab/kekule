# Kekule

A very basic RESTful API for a fictional online store that
implements CRUD operations for store items.

An item is composed of

1. ID,
2. SKU,
3. name,
4. price,
5. and category.

## Compiling and Running

To build and run the app Go compiler 1.16 is required.
Please refer to the [official installation steps](https://golang.org/doc/install) for
instructions on setting it up.

To start the API server use the following command:

    go run cmd/kekule/server.go

## Methods

The API supports basic CRUD operations.

### `GET /api/item`

Fetches a single item.

Parameters:

1. `id` of type `integer`: the ID of the item.

On success, returns a JSON-object like:

```json
{
  "result": "success",
  "item": {
    "id": 1,
    "sku": "c2n1",
    "name": "Item1",
    "price": 20.0,
    "category": 2
  }
}
```

### `GET /api/item/list`

Fetches a list of items.

Parameters:

1. `page` of type `integer`, must be larger or equal to 1.
2. `limit` of type `integer`:
how many items to fetch at most, must be less than 100.
3. `lowest_price` of type `double`: the lowest price
an item can have in the the result set, must be positive.
4. `highest_price` of type `double`: the highest price
an item can have in the result set, must be positive.
5. `category` of type `integer`: specifies the category
for items in the result set, must be positive and valid.

On success, returns a JSON-object like:

```jsonc
{
  "result": "success",
  "items": [
    {
      "id": 1,
      "sku": "c2n1",
      "name": "Item1",
      "price": 20.0,
      "category": 2
    },
    // ...
  ]
}
```

### `POST /api/item`

Creates a new item.

Parameters:

1. `name` of type `string`: the name of the item, length
cannot exceed 64 characters.
2. `price` of type `double`: the price of the item.
3. `category` of type `int`: the category of the item.

On success, returns a JSON-object like:

```jsonc
{
  "result": "success",
  "id": 12 // ID of the newly created item.
}
```

### `PUT /api/item`

Updates an existing item.

Parameters:

1. `id` of type `integer`: the ID of the item.
2. `name` of type `string`: the new name of the item, length
cannot exceed 64 characters, optional.
3. `price` of type `double`: the new price of the item, optional.
4. `category` of type `int`: the new category of the item, optional.

On success, returns a JSON-object like:

```jsonc
{
  "result": "success",
  "description": "item updated"
}
```

### `DELETE /api/item`

Deletes an existing item.

Parameters:

1. `id` of type `integer`: the ID of the item to delete.

On success, returns a JSON-object like:

```jsonc
{
  "result": "success",
  "description": "item deleted"
}
```

### `GET /api/category`

Fetches a category.

1. `id` of type `integer`: the ID of the category to fetch.

On success, returns JSON-object like:

```jsonc
{
  "result": "success",
  "category": {
    "id": 7,
    "name": "category name"
  }
}
```

### Errors

All methods return a JSON-object of the following
form if an error occurred when handling a request:

```jsonc
{
  "result": "error message",
  "description": "error details"
}
```
