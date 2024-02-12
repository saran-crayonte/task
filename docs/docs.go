// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Saran",
            "url": "github.com/saran-crayonte/",
            "email": "saran.kumaresan@crayonte.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/user": {
            "post": {
                "description": "Register a new user with username, name, email, and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Username already exists",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/login": {
            "post": {
                "description": "Logs in a user with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User authenticated",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "User not found / Password doesn't match",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/alluser": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Get all users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            }
        },
        "/api/v2/holiday": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve all task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Management"
                ],
                "summary": "Get all task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new holiday with provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Holiday Management"
                ],
                "summary": "Create a new holiday",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Holiday details",
                        "name": "holiday",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Holiday"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Holiday created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Holiday already defined",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/holiday/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve a holiday by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Holiday Management"
                ],
                "summary": "Get a holiday by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Holiday ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Holiday retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Holiday"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Holiday not found",
                        "schema": {
                            "type": "string"
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
                "description": "Update an existing holiday by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Holiday Management"
                ],
                "summary": "Update a holiday by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Updated holiday details",
                        "name": "holiday",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Holiday"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Holiday updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Holiday not found",
                        "schema": {
                            "type": "string"
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
                "description": "Delete an existing holiday by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Holiday Management"
                ],
                "summary": "Delete a holiday by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Holiday ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Holiday deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Holiday not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/refreshToken": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Refreshes the authentication token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Refresh authentication token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Refresh user token",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token refreshed successfully",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/task": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve all holidays",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Holiday Management"
                ],
                "summary": "Get all holidays",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Holiday retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Holiday"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new task with provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Management"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Task details",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Task created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/task/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve a task by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Management"
                ],
                "summary": "Get a task by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "type": "string"
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
                "description": "Update an existing task by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Management"
                ],
                "summary": "Update a task by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Updated task details",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "type": "string"
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
                "description": "Delete an existing task by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Management"
                ],
                "summary": "Delete a task by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/taskAssignment": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve all task assignments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Assignment"
                ],
                "summary": "Get all task assignments",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task Assignment retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.TaskAssignment"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new task assignment with provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Assignment"
                ],
                "summary": "Create a new task assignment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Task assignment details",
                        "name": "taskAssignment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TaskAssignment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Task assignment created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Username doesn't exist / Task not found / Task is already assigned",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/taskAssignment/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve a task assignment by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Assignment"
                ],
                "summary": "Get a task assignment by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Task Assignment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task assignment retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.TaskAssignment"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Task assignment not found",
                        "schema": {
                            "type": "string"
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
                "description": "Update an existing task assignment by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Assignment"
                ],
                "summary": "Update a task assignment by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Updated task assignment details",
                        "name": "taskAssignment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TaskAssignment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task assignment updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Username doesn't exist / Task not found / Task assignment not found",
                        "schema": {
                            "type": "string"
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
                "description": "Delete an existing task assignment by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task Assignment"
                ],
                "summary": "Delete a task assignment by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Task Assignment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task assignment deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Task assignment not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/user": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates the password of the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Update user password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Update Password Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Username doesn't exist",
                        "schema": {
                            "type": "string"
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
                "description": "Deletes the account of the authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Delete user account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User deletion request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Username doesn't exist",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Holiday": {
            "type": "object",
            "properties": {
                "holidayDate": {
                    "type": "string"
                },
                "holidayName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.Task": {
            "type": "object",
            "properties": {
                "estimatedHours": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.TaskAssignment": {
            "type": "object",
            "properties": {
                "endDate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "startDate": {
                    "type": "string"
                },
                "taskid": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.User": {
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
                "username": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Task Management API",
	Description:      "This is a sample API for managing tasks.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
