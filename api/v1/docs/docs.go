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
        "/coins": {
            "post": {
                "description": "PostCoinUrl",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "coins"
                ],
                "summary": "Get a list of the best investors of crypto coin",
                "parameters": [
                    {
                        "description": "url of coin page on cypherhunter.com",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.CoinUrl"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/main.Investor"
                        }
                    },
                    "500": {
                        "description": "fail",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.CoinUrl": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "main.Investor": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tier": {
                    "type": "string"
                },
                "tierId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Follow Smart Money API",
	Description:      "Get a list of the best investors of crypto coin. Built on top of github.com/feynmaz/cypherhunterscrapper",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
