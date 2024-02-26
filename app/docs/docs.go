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
                "description": "Get user by ID",
                "summary": "Create User Entry",
                "operationId": "post-customer",
                "parameters": [
                    {
                        "description": "Customer ID",
                        "name": "customerid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User Agent",
                        "name": "useragent",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Remote IP",
                        "name": "remoteip",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Timestamp",
                        "name": "timestamp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Server Error"
                    }
                }
            }
        },
        "/statistics/": {
            "post": {
                "description": "Get statistics based on Day and customer",
                "produces": [
                    "application/json"
                ],
                "summary": "Get statistics based on day and customer",
                "operationId": "get-stats",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Customer ID",
                        "name": "customerid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Day",
                        "name": "day",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Timestamp",
                        "name": "timestamp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Customer data not found"
                    },
                    "500": {
                        "description": "Server Error"
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