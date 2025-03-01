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
        "/categories": {
            "get": {
                "description": "Fetches all available categories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Options"
                ],
                "summary": "Get all categories",
                "responses": {
                    "200": {
                        "description": "List of categories",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Category"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch categories",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/cities": {
            "get": {
                "description": "Fetches all available cities",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Countries/Cities"
                ],
                "summary": "Get all cities",
                "responses": {
                    "200": {
                        "description": "List of cities",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.City"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch cities",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/conveniences": {
            "get": {
                "description": "Fetches all available conveniences",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Options"
                ],
                "summary": "Get all conveniences",
                "responses": {
                    "200": {
                        "description": "List of conveniences",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Category"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch conveniences",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/countries": {
            "get": {
                "description": "Fetches all available countries",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Countries/Cities"
                ],
                "summary": "Get all countries",
                "responses": {
                    "200": {
                        "description": "List of countries",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Country"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch countries",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/country/{id}/cities": {
            "get": {
                "description": "Get all cities in the specified country by their countryID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Countries/Cities"
                ],
                "summary": "Get all cities in the specified country",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Country ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of cities",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.City"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Server error",
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
        "/products": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Returns a paginated list of products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get a list of products",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 20,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Server error",
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
        "/products/{id}/like": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Increments the product's like count",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Like a product by slug",
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
                        "description": "Product likeed successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
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
                "description": "Increments the product's like count",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Dislike a product by id",
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
                        "description": "Product disliked successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Server error",
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
        "/products/{slug}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Returns product details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get a product by slug",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Server error",
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
        "/types": {
            "get": {
                "description": "Fetches all available types",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Options"
                ],
                "summary": "Get all types",
                "responses": {
                    "200": {
                        "description": "List of types",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ProductType"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch types",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/user/favorite/products": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Returns a paginated list of user favorite products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserProducts"
                ],
                "summary": "Get a list of user favorite products",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 20,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Server error",
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
        "/users/create": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User registration",
                "parameters": [
                    {
                        "description": "User registration data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Registration successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or registration error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Returns user details for the authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Get user details",
                "responses": {
                    "200": {
                        "description": "User details",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/users/token": {
            "post": {
                "description": "Authenticate user and return access tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.AuthLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully authenticated",
                        "schema": {
                            "$ref": "#/definitions/schema.AuthLoginRes"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or credentials",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "fiber.Map": {
            "type": "object",
            "additionalProperties": true
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "models.City": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "postall_code": {
                    "type": "string"
                }
            }
        },
        "models.Convenience": {
            "type": "object",
            "properties": {
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "models.Country": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.OwnerProduct": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "bath_qty": {
                    "type": "integer"
                },
                "bed_qty": {
                    "type": "integer"
                },
                "bedroom_qty": {
                    "type": "integer"
                },
                "best_product": {
                    "type": "boolean"
                },
                "bookings": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "city": {
                    "type": "string"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ProductComment"
                    }
                },
                "convenience": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Convenience"
                    }
                },
                "country": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "guest_qty": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ProductImages"
                    }
                },
                "is_new": {
                    "type": "boolean"
                },
                "lat": {
                    "type": "string"
                },
                "like_count": {
                    "type": "integer"
                },
                "lng": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/models.OwnerProduct"
                },
                "phone_number": {
                    "type": "string"
                },
                "price_per_month": {
                    "type": "integer"
                },
                "price_per_night": {
                    "type": "integer"
                },
                "price_per_week": {
                    "type": "integer"
                },
                "promotion": {
                    "type": "boolean"
                },
                "rating": {
                    "type": "number"
                },
                "rooms_qty": {
                    "type": "integer"
                },
                "slug": {
                    "type": "string"
                },
                "toilet_qty": {
                    "type": "integer"
                },
                "type": {
                    "$ref": "#/definitions/models.ProductType"
                },
                "type_id": {
                    "type": "integer"
                }
            }
        },
        "models.ProductComment": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/models.ProductCommentUser"
                }
            }
        },
        "models.ProductCommentUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "models.ProductImages": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_label": {
                    "type": "boolean"
                },
                "mimetype": {
                    "type": "string"
                },
                "original": {
                    "type": "string"
                },
                "thumbnail": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "models.ProductType": {
            "type": "object",
            "properties": {
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "date_joined": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "iin": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "schema.AuthLoginReq": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "schema.AuthLoginRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "schema.RegisterReq": {
            "type": "object",
            "required": [
                "date_of_birth",
                "email",
                "first_name",
                "last_name",
                "password",
                "phone_number"
            ],
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "iin": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "phone_number": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "SuperApp API",
	Description:      "API documentation for SuperApp",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
