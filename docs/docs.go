// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/user": {
            "get": {
                "security": [
                    {
                        "BaseAuth": []
                    }
                ],
                "description": "get user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userdata"
                ],
                "summary": "Get User",
                "operationId": "get-user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/userdata": {
            "get": {
                "security": [
                    {
                        "BaseAuth": []
                    }
                ],
                "description": "get userdata by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userdata"
                ],
                "summary": "Get UserData By Id",
                "operationId": "get-userdata-by-id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BaseAuth": []
                    }
                ],
                "description": "create userdata",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userdata"
                ],
                "summary": "Create userdata",
                "operationId": "create-userdata",
                "parameters": [
                    {
                        "description": "userdata",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/workouts": {
            "get": {
                "security": [
                    {
                        "BaseAuth": []
                    }
                ],
                "description": "get workouts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workouts"
                ],
                "summary": "Get Workouts",
                "operationId": "get-workouts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Workout"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignIn",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "Input request structure",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "session",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "Input request structure",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "Пароль пользователя.",
                    "type": "string"
                },
                "username": {
                    "description": "Имя пользователя.",
                    "type": "string"
                }
            }
        },
        "handler.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Exercise": {
            "type": "object",
            "properties": {
                "calories": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "reps": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "sets": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "weights": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "passwordHash": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.UserData": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "calories": {
                    "type": "integer"
                },
                "goal": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "place": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "models.Workout": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "exercises": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Exercise"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "workoutExercises": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.WorkoutExercise"
                    }
                }
            }
        },
        "models.WorkoutExercise": {
            "type": "object",
            "properties": {
                "exerciseID": {
                    "type": "integer"
                },
                "workoutID": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BaseAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "SmartFit App API",
	Description:      "API Server for SmartFit Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
