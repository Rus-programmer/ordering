// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customers": {
            "post": {
                "description": "Register a new customer with username, password, and role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Create a new customer",
                "parameters": [
                    {
                        "description": "Customer details",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCustomerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CustomerResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate user with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginResponse"
                        }
                    }
                }
            }
        },
        "/metrics": {
            "get": {
                "description": "Retrieve various system metrics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "Get system metrics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.MetricsResponse"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a list of orders with optional filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "List orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Minimum price",
                        "name": "min_price",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Maximum price",
                        "name": "max_price",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.OrderResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new order with the provided products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create an order",
                "parameters": [
                    {
                        "description": "Order details",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateOrderRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.OrderResponse"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve an order's details by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get an order by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.OrderResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an existing order's details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Update an order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated order details",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateOrderRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.OrderResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete an order by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete an order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order deleted successfully"
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Retrieve a list of products with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "List products",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page ID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ProductResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new product with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create a product",
                "parameters": [
                    {
                        "description": "Product details",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.createProductRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProductResponse"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Retrieve a product's details by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get a product by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProductResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing product's details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated product details",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.updateProductRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProductResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a product by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Delete a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product deleted successfully"
                    }
                }
            }
        },
        "/renew_access": {
            "post": {
                "description": "Register a new customer with username, password, and role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Create a new customer",
                "parameters": [
                    {
                        "description": "Customer details",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCustomerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CustomerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.createProductRequestBody": {
            "type": "object",
            "required": [
                "name",
                "price",
                "quantity"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer",
                    "minimum": 1
                },
                "quantity": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "api.updateProductRequestBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer",
                    "minimum": 1
                },
                "quantity": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "db.OrderStatus": {
            "type": "string",
            "enum": [
                "pending",
                "confirmed",
                "cancelled"
            ],
            "x-enum-varnames": [
                "OrderStatusPending",
                "OrderStatusConfirmed",
                "OrderStatusCancelled"
            ]
        },
        "db.UserRole": {
            "type": "string",
            "enum": [
                "user",
                "admin"
            ],
            "x-enum-varnames": [
                "UserRoleUser",
                "UserRoleAdmin"
            ]
        },
        "dto.CreateCustomerRequest": {
            "type": "object",
            "required": [
                "password",
                "role",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "role": {
                    "$ref": "#/definitions/db.UserRole"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.CreateOrderItem": {
            "type": "object",
            "required": [
                "ordered_amount",
                "product_id"
            ],
            "properties": {
                "ordered_amount": {
                    "type": "integer",
                    "minimum": 1
                },
                "product_id": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "dto.CreateOrderRequestBody": {
            "type": "object",
            "required": [
                "products"
            ],
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateOrderItem"
                    }
                }
            }
        },
        "dto.CustomerResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "role": {
                    "$ref": "#/definitions/db.UserRole"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "access_token_expires_at": {
                    "type": "string"
                },
                "customer": {
                    "$ref": "#/definitions/dto.CustomerResponse"
                },
                "refresh_token": {
                    "type": "string"
                },
                "refresh_token_expires_at": {
                    "type": "string"
                },
                "session_id": {
                    "type": "string"
                }
            }
        },
        "dto.MetricsResponse": {
            "type": "object",
            "properties": {
                "error_rate": {
                    "type": "number"
                },
                "requests_by_method": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "requests_by_path": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "requests_by_status_code": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "total_requests": {
                    "type": "integer"
                }
            }
        },
        "dto.OrderProductResponse": {
            "type": "object",
            "properties": {
                "ordered_amount": {
                    "type": "integer"
                },
                "product": {
                    "$ref": "#/definitions/dto.ProductResponse"
                }
            }
        },
        "dto.OrderResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "id_deleted": {
                    "type": "boolean"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.OrderProductResponse"
                    }
                },
                "status": {
                    "$ref": "#/definitions/db.OrderStatus"
                },
                "total_price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.ProductResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateOrderRequestBody": {
            "type": "object",
            "required": [
                "products"
            ],
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateOrderItem"
                    }
                },
                "status": {
                    "$ref": "#/definitions/db.OrderStatus"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Ordering project",
	Description:      "This project is a backend service designed for managing orders",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
