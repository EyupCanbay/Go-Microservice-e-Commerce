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
        "/customer/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new customer with the given data",
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
                        "description": "Customer to create",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateCustomerRequestModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Returns created customer ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    }
                }
            }
        },
        "/customer/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve a paginated list of customers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "List customers with pagination",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns list of customers",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    }
                }
            }
        },
        "/customer/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a customer by its unique ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Get customer by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CustomerResponseModel"
                        }
                    },
                    "400": {
                        "description": "Invalid ID format",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    },
                    "404": {
                        "description": "Customer not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a customer with the given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Update an existing customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Customer data to update",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UpdateCustomerRequestModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CustomerResponseModel"
                        }
                    },
                    "400": {
                        "description": "Invalid ID format or request body",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    },
                    "404": {
                        "description": "Customer not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a customer from the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Delete a customer by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid ID format",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    },
                    "404": {
                        "description": "Customer not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "pkg.AppError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "types.Address": {
            "type": "object",
            "properties": {
                "address_id": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "zip_code": {
                    "type": "string"
                }
            }
        },
        "types.CreateCustomerRequestModel": {
            "type": "object",
            "required": [
                "FirstName",
                "address",
                "email",
                "last_name",
                "password",
                "phone"
            ],
            "properties": {
                "FirstName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "address": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Address"
                    }
                },
                "email": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "password": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "phone": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Phone"
                    }
                }
            }
        },
        "types.CustomerResponseModel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Address"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Phone"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "types.Phone": {
            "type": "object",
            "properties": {
                "phone_id": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "integer"
                }
            }
        },
        "types.UpdateCustomerRequestModel": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "address": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Address"
                    }
                },
                "email": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "password": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "phone": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Phone"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
