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
        "/api/v1/": {
            "get": {
                "description": "根路由",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "根路由",
                "responses": {
                    "200": {
                        "description": "欢迎使用OPC-UA OpenAPI",
                        "schema": {
                            "$ref": "#/definitions/routers.ApiResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/ping": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "ping 路由",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "ping 路由",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/publish/:topic": {
            "post": {
                "description": "Publish a message to a topic",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Publish"
                ],
                "summary": "Publish a message to a topic",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Topic name",
                        "name": "topic",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Message to publish",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.ApiOKResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/subscribe/{topic}": {
            "get": {
                "description": "# 数据订阅接口\n## 使用方法\n` + "`" + `` + "`" + `` + "`" + `javascript\nvar ws = new WebSocket(\"ws://10.30.24.115:8060/api/v1/data-out/topic\");\n\nws.onmessage = function(event) {\nconsole.log(event.data);\n};\n\nws.onopen = function() {\nconsole.log(\"Connection established\");\n};\n\nws.onclose = function() {\nconsole.log(\"Connection closed\");\n};\n` + "`" + `` + "`" + `` + "`" + `",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Out"
                ],
                "summary": "Send To Front End by websocket",
                "responses": {
                    "200": {
                        "description": "Hello, World!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "healthCheck 路由",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "healthCheck 路由",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "core.ApiOKResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "200",
                    "type": "integer"
                },
                "data": {
                    "type": "string"
                },
                "message": {
                    "description": "\"success\"",
                    "type": "string"
                }
            }
        },
        "routers.ApiResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "欢迎使用OPC-UA OpenAPI"
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
