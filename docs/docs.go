// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "This method allows user to log-in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public routes. Registration and Authentication"
                ],
                "summary": "Log-In user",
                "parameters": [
                    {
                        "description": "Input data for user log-in",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LogInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User logged-in successfully",
                        "schema": {
                            "$ref": "#/definitions/models.LogInResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "This method allows user to release new access and refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public routes. Registration and Authentication"
                ],
                "summary": "Refresh user tokens (Access and Refresh)",
                "parameters": [
                    {
                        "description": "Input data for token refresh",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RefreshRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tokens refreshed successfully",
                        "schema": {
                            "$ref": "#/definitions/models.RefreshResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/revoke": {
            "post": {
                "description": "This method allows user to revoke access and refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public routes. Registration and Authentication"
                ],
                "summary": "Revoke user tokens (Access and Refresh)",
                "parameters": [
                    {
                        "description": "Input data for token revoke",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RevokeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tokens revoked successfully",
                        "schema": {
                            "$ref": "#/definitions/models.RevokeResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "This method allows user to create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public routes. Registration and Authentication"
                ],
                "summary": "Sign-Up new user and authorize him",
                "parameters": [
                    {
                        "description": "Input data for user registration",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cart/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This method allows user to create cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart Operations. This is a simple API for online shopping cart"
                ],
                "summary": "Creates shopping cart",
                "responses": {
                    "201": {
                        "description": "Cart created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.CreateCartResponse"
                        }
                    },
                    "401": {
                        "description": "Empty claims",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cart/{cart_id}/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This method allows user to add item to cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CartItem Operations. This is a simple API for online shopping cart"
                ],
                "summary": "Add item to shopping cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cart ID",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Item",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Item added to cart successfully",
                        "schema": {
                            "$ref": "#/definitions/models.CartItem"
                        }
                    },
                    "400": {
                        "description": "Invalid cart ID",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Empty claims",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cart/{cart_id}/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This method allows user to get cart by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart Operations. This is a simple API for online shopping cart"
                ],
                "summary": "Get shopping cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cart ID",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cart retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.GetCartResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid cart ID",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Empty claims",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cart/{cart_id}/remove/{item_id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This method allows user to remove item from cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CartItem Operations. This is a simple API for online shopping cart"
                ],
                "summary": "Remove item from shopping cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cart ID",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "item_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Item removed from cart successfully",
                        "schema": {
                            "$ref": "#/definitions/models.RemoveItemMessageResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid cart ID",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Empty claims",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddItemRequest": {
            "type": "object",
            "properties": {
                "product": {
                    "type": "string",
                    "example": "apple"
                },
                "quantity": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "models.CartItem": {
            "type": "object",
            "properties": {
                "cart_id": {
                    "type": "integer",
                    "example": 1
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "product": {
                    "type": "string",
                    "example": "item1"
                },
                "quantity": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.CreateCartResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CartItem"
                    }
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.GetCartResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CartItem"
                    }
                }
            }
        },
        "models.LogInRequest": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "type": "string",
                    "example": "+1111111111"
                }
            }
        },
        "models.LogInResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "User logged-in successfully!"
                }
            }
        },
        "models.RefreshRequest": {
            "type": "object",
            "required": [
                "refreshToken"
            ],
            "properties": {
                "refreshToken": {
                    "type": "string",
                    "example": "...(numeric-letter string)"
                }
            }
        },
        "models.RefreshResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Token refreshed successfully!"
                }
            }
        },
        "models.RemoveItemMessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Item removed from cart successfully"
                }
            }
        },
        "models.RevokeRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "johndoe1@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "type": "string",
                    "example": "+1111111111"
                }
            }
        },
        "models.RevokeResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Tokens revoked!"
                }
            }
        },
        "models.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "johndoe1@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "type": "string",
                    "example": "+1111111111"
                }
            }
        },
        "models.SignUpResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "User created successfully!"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Online shopping cart API",
	Description:      "This is a simple API for online shopping cart",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
