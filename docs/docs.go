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
        "/bank/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create new bank",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "banks"
                ],
                "summary": "create bank",
                "parameters": [
                    {
                        "description": "bank",
                        "name": "bank",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBankDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bank/q": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all the existing banks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "banks"
                ],
                "summary": "List existing banks",
                "parameters": [
                    {
                        "description": "bank",
                        "name": "bankDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBankDto"
                        }
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "page Index",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "page Size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/common.DefaultResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.DefaultResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "developerMessage": {
                    "type": "string"
                },
                "http_status": {
                    "type": "string"
                },
                "resultCode": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateBankDto": {
            "type": "object",
            "required": [
                "bankCode",
                "bankName"
            ],
            "properties": {
                "bankCode": {
                    "type": "string"
                },
                "bankName": {
                    "type": "string"
                },
                "userInserted": {
                    "type": "string"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
