// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/auth/login": {
            "post": {
                "description": "Make login and get access-token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/auth.JwtToken"
                            }
                        }
                    },
                    "401": {
                        "description": "when email or password is invalid",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            }
        },
        "/auth/refresh-token": {
            "post": {
                "description": "Get a new access-token, this action will be expire the last one",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh access-token",
                "parameters": [
                    {
                        "description": "the refresh token",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/auth.JwtToken"
                            }
                        }
                    },
                    "401": {
                        "description": "when token is invalid",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Check app and dependencies status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health check"
                ],
                "summary": "Health Cehck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/products/": {
            "get": {
                "description": "Get all products data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/products.Product"
                            }
                        }
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new product and return data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create a New Product",
                "parameters": [
                    {
                        "description": "the product data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/products.Product"
                        }
                    },
                    "422": {
                        "description": "when data is invalid",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get product Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/products.Product"
                        }
                    },
                    "422": {
                        "description": "when id is invalid",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete Product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            }
        },
        "/sales": {
            "get": {
                "description": "get sales",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sales"
                ],
                "summary": "List sales",
                "parameters": [
                    {
                        "type": "string",
                        "format": "dateTime",
                        "description": "name search by q",
                        "name": "startAt",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "dateTime",
                        "description": "name search by q",
                        "name": "endAt",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/sales.Sale"
                            }
                        }
                    },
                    "422": {
                        "description": "when start or end param is invalid",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            },
            "post": {
                "description": "Register a sale and return sale data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sales"
                ],
                "summary": "Register a new sale",
                "parameters": [
                    {
                        "description": "the payload data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterSalePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/sales.Sale"
                            }
                        }
                    },
                    "422": {
                        "description": "when payload is invalid",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create a new user and return user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "the user data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserPayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/users.User"
                        }
                    },
                    "422": {
                        "description": "when payload is invalid",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "description": "Get current user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.User"
                        }
                    },
                    "500": {
                        "description": "when an error occurs",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageJSON"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.JwtToken": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expiration": {
                    "type": "integer"
                },
                "grant_type": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.CreateProductPayload": {
            "type": "object",
            "properties": {
                "atacado_amount": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price_atacado": {
                    "type": "number"
                },
                "price_varejo": {
                    "type": "number"
                }
            }
        },
        "dto.CreateUserPayload": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.LoginPayload": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.MessageJSON": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.RefreshTokenPayload": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterSalePayload": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sales.CartItem"
                    }
                },
                "payment_type": {
                    "type": "string"
                }
            }
        },
        "products.Product": {
            "type": "object",
            "required": [
                "name",
                "price_atacado",
                "price_varejo"
            ],
            "properties": {
                "atacado_amount": {
                    "type": "integer",
                    "minimum": 1
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price_atacado": {
                    "type": "number"
                },
                "price_varejo": {
                    "type": "number"
                }
            }
        },
        "sales.CartItem": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "item_id": {
                    "type": "string"
                }
            }
        },
        "sales.Item": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "unit_price": {
                    "type": "number"
                }
            }
        },
        "sales.Sale": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sales.Item"
                    }
                },
                "payment_type": {
                    "type": "string"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "users.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
	Title:            "Sorveteria três estrelas - API",
	Description:      "API para o cadastro de produtos, controle de vendas e fluxo de caixa para a sorveteria três estrelas",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
