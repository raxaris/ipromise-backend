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
        "/admin/promises/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Позволяет администратору удалить обещание по ID",
                "tags": [
                    "admin"
                ],
                "summary": "Удаление обещания",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID обещания",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Обещание удалено",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "error: Ошибка при удалении обещания",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "error: У вас нет прав на удаление обещания",
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
        "/admin/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает список всех зарегистрированных пользователей",
                "tags": [
                    "admin"
                ],
                "summary": "Получение всех пользователей",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "500": {
                        "description": "error: Ошибка сервера",
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
        "/auth/login": {
            "post": {
                "description": "Логин по email и паролю, выдаёт JWT токены",
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
                        "description": "Данные для входа",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "access_token: токен, refresh_token: токен",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "error: Ошибка валидации",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "error: Неверный email или пароль",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "error: Ошибка сервера",
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
        "/auth/refresh": {
            "post": {
                "description": "Использует Refresh Token для выдачи нового Access Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Обновление Access Token",
                "parameters": [
                    {
                        "description": "Refresh Token",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "access_token: новый access-токен",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "error: Ошибка валидации",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "error: Недействительный Refresh-токен",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "error: Ошибка сервера",
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
        "/auth/signup": {
            "post": {
                "description": "Создаёт нового пользователя по email, имени и паролю",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Регистрация нового пользователя",
                "parameters": [
                    {
                        "description": "Данные для регистрации пользователя",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignupRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "message: Пользователь успешно зарегистрирован",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "error: Неверные данные запроса",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "409": {
                        "description": "error: Email или имя пользователя уже занято",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "error: Внутренняя ошибка сервера",
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
        "/promises": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает список всех обещаний",
                "tags": [
                    "promises"
                ],
                "summary": "Получение всех обещаний",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Promise"
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
                "description": "Позволяет пользователю создать новое обещание",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "promises"
                ],
                "summary": "Создание нового обещания",
                "parameters": [
                    {
                        "description": "Данные обещания",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePromiseRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "message: Обещание успешно создано",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "error: Ошибка валидации",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "error: Ошибка сервера",
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
        "/promises/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Позволяет обновить информацию о обещании",
                "tags": [
                    "promises"
                ],
                "summary": "Обновление обещания",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID обещания",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePromiseRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Обещание обновлено",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "error: Ошибка валидации",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "error: Нет прав на редактирование",
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
        "/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает данные текущего пользователя",
                "tags": [
                    "users"
                ],
                "summary": "Получение информации о себе",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "error: Пользователь не найден",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
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
                "description": "Позволяет изменить username (доступно только самому пользователю)",
                "tags": [
                    "users"
                ],
                "summary": "Обновление профиля пользователя",
                "parameters": [
                    {
                        "description": "Данные для обновления",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Данные пользователя обновлены",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "error: Ошибка валидации",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "error: Нет прав на редактирование",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "error: Ошибка сервера",
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
                "description": "Удаляет аккаунт текущего пользователя",
                "tags": [
                    "users"
                ],
                "summary": "Удаление аккаунта",
                "responses": {
                    "200": {
                        "description": "message: Аккаунт удалён",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "error: Пользователь не найден",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "error: Ошибка удаления",
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
        "/users/username/{username}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает данные пользователя по username",
                "tags": [
                    "users"
                ],
                "summary": "Получение пользователя по username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Имя пользователя",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "error: Пользователь не найден",
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
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает данные пользователя по ID",
                "tags": [
                    "users"
                ],
                "summary": "Получение пользователя по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "error: Неверный формат ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "error: Пользователь не найден",
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
        "/users/{user_id}/promises": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает список обещаний, созданных пользователем",
                "tags": [
                    "promises"
                ],
                "summary": "Получение обещаний пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID пользователя",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Promise"
                            }
                        }
                    },
                    "400": {
                        "description": "error: Неверный формат ID пользователя",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "error: Ошибка получения обещаний",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreatePromiseRequest": {
            "type": "object",
            "required": [
                "status",
                "title"
            ],
            "properties": {
                "deadline": {
                    "description": "Можно передавать только для основного промиса",
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "parent_id": {
                    "description": "Если null, это основное обещание",
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRequest": {
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
        "dto.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.SignupRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UpdatePromiseRequest": {
            "type": "object",
            "properties": {
                "deadline": {
                    "description": "Только для основного обещания",
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "role": {
                    "description": "Только для админов",
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Promise": {
            "type": "object",
            "properties": {
                "deadline": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "parentID": {
                    "description": "NULL, если это основной Promise",
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "userID": {
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
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "iPromise API",
	Description:      "API для отслеживания обещаний пользователей.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
