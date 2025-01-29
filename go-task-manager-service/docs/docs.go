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
            "name": "Rybakov Dmitry",
            "email": "dimryb@bk.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Авторизация пользователя",
                "parameters": [
                    {
                        "description": "Данные пользователя",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Регистрация пользователя",
                "parameters": [
                    {
                        "description": "Данные пользователя",
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
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Возвращает задачи, отфильтрованные по статусу, приоритету и дате",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Получение списка задач с фильтрацией",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Статус (pending, in_progress, done)",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Приоритет (low, medium, high)",
                        "name": "priority",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата выполнения формат: 2020-01-01T12:00:00Z",
                        "name": "due_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Название задачи",
                        "name": "title",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "400": {
                        "description": "Некорректные параметры запроса",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Создает новую задачу",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Создание задачи",
                "parameters": [
                    {
                        "description": "Создание задачи",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "get": {
                "description": "Возвращает задачу по её ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Получение задачи по ID (вспомогательная)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID задачи",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "400": {
                        "description": "Некорректный ID",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "404": {
                        "description": "Задача не найдена",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновляет информацию о задаче",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Обновление задачи",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID задачи",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "400": {
                        "description": "Некорректные данные",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "404": {
                        "description": "Задача не найдена",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет задачу по указанному ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Удаление задачи",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID задачи",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Задача успешно удалена",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "400": {
                        "description": "Некорректный ID задачи",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "404": {
                        "description": "Задача не найдена",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateTaskRequest": {
            "type": "object",
            "required": [
                "due_date",
                "priority",
                "status",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Description"
                },
                "due_date": {
                    "type": "string",
                    "example": "2025-01-28T12:00:00Z"
                },
                "priority": {
                    "type": "string",
                    "enum": [
                        "low",
                        "medium",
                        "high"
                    ],
                    "example": "medium"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "pending",
                        "in_progress",
                        "done"
                    ],
                    "example": "pending"
                },
                "title": {
                    "type": "string",
                    "example": "Title"
                }
            }
        },
        "models.UpdateTaskRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Updated Description"
                },
                "due_date": {
                    "type": "string",
                    "example": "2025-02-01T15:00:00Z"
                },
                "priority": {
                    "type": "string",
                    "enum": [
                        "low",
                        "medium",
                        "high"
                    ],
                    "example": "high"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "pending",
                        "in_progress",
                        "done"
                    ],
                    "example": "in_progress"
                },
                "title": {
                    "type": "string",
                    "example": "Updated Title"
                }
            }
        },
        "models.User": {
            "type": "object"
        },
        "rest.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Task Manager Service",
	Description:      "This is a service for managing tasks.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
