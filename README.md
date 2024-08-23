Cart API
========

## Overview

This repository contains a Go-based implementation of a REST API for an online shopping cart. The application is designed to fulfill a coding exercise, demonstrating key competencies in Go, REST API design, and database interaction.

## Implementation Details

The API supports basic CRUD operations for an online shopping cart. The data is persisted using Postgres, leveraging the `sqlx` library for database interactions. All SQL queries are manually written to ensure clarity and precision. The project is built with a focus on simplicity, security, and performance, utilizing the default `gin` package for RESTful implementations.

### Domain Types

The API manages two main types: `Cart` and `CartItem`. Each `Cart` can contain multiple `CartItem` objects. The `CartItem` entities are created directly in the database to ensure data consistency.

### API Endpoints

The following endpoints have been implemented:


### Create Cart

Creates a new cart with a unique ID and returns the newly created cart.

```sh
POST http://localhost:3000/cart/create -d '{}'
```

```json
{
	"id": 1,
	"items": []
}
```

### Add to cart

Adds a new item to an existing cart. The operation fails if the cart does not exist, if the product name is blank, or if the quantity is non-positive.

```sh
POST http://localhost:3000/cart/{cart_id}/add -d '{
	"product": "Shoes",
	"quantity": 10
}'
```

```json
{
	"id": 1,
	"cart_id": 1,
	"product": "Shoes",
	"quantity": 10
}
```

### Remove from cart

Removes an existing item from a cart. The operation fails if the cart or the item does not exist.

```sh
DELETE http://localhost:3000/cart/{cart_id}/remove/{item_id}
```

```json
{
  "message": "Item removed from cart successfully"
}
```


### View cart

Retrieves an existing cart along with its items. The operation fails if the cart does not exist.

```sh
GET http://localhost:3000/cart/{cart_id}/get
```

```json
{
	"id": 1,
	"items": [
		{
			"id": 1,
			"cart_id": 1,
			"product": "Shoes",
			"quantity": 10
		},
		{
			"id": 2,
			"cart_id": 1,
			"product": "Socks",
			"quantity": 5
		}
	]
}
```

