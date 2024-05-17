
Cart API
========

## Overview

This repository contains a Go programming exercise for interview candidates. 
You'll be developing an API for an online shopping cart in the Go programming language.


## Requirements

This is a REST API for basic CRUD operations for an online shopping cart. Data
should be persisted in a storage layer which can use Postgres.

You should use default `net/http` package for REST implementation; `sqlx` or `sqlc` for interact with postgres;
all the queries should be wrote mannually (no ORM, no `Select` methods and etc.); your repo should be private.

#### Additional requirements

Cover your code with the unit tests (you could use `testify`).

Create a `Dockerfile` for your application.


### Domain Types

The Cart API consists of two simple types: `Cart` and `CartItem`. The `Cart`
holds zero or more `CartItem` objects.

`CartItem` objects should be created in DB exactly (not from application).


### Create Cart

A new cart should be created and an ID generated. The new empty cart should be returned.

```sh
POST http://localhost:3000/carts -d '{}'
```

```json
{
	"id": 1,
	"items": []
}
```

### Add to cart

A new item should added to an existing cart. Should fail if the cart does not
exist, if the product name is blank, or if the quantity is non-positive. The
new item should be returned.

```sh
POST http://localhost:3000/carts/1/items -d '{
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

An existing item should be removed from a cart. Should fail if the cart does not
exist or if the item does not exist.

```sh
DELETE http://localhost:3000/carts/1/items/1
```

```json
{}
```


### View cart

An existing cart should be able to be viewed with its items. Should fail if the
cart does not exist.

```sh
GET http://localhost:3000/carts/1
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

