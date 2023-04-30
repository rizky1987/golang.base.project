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
        "/api/cms/v1/floor/create": {
            "post": {
                "description": "Floor Create",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Floor"
                ],
                "operationId": "FloorCreate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer %",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "number",
                        "name": "number",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "roomTypeId",
                        "name": "roomTypeId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "price",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.CommonBaseResponse"
                        }
                    }
                }
            }
        },
        "/api/cms/v1/room-type/create": {
            "post": {
                "description": "RoomType Create",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "RoomType"
                ],
                "operationId": "RoomTypeCreate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer %",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "code",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.CommonBaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.AlertResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "inner_message": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "responses.CommonBaseResponse": {
            "type": "object",
            "properties": {
                "alert": {
                    "$ref": "#/definitions/responses.AlertResponse"
                },
                "data": {}
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
