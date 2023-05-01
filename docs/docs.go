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
        "/api/cms/v1/booking/create": {
            "post": {
                "description": "Booking Create",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "operationId": "BookingCreate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer %",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateBookingRequest"
                        }
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
        "/api/cms/v1/room-price/create": {
            "post": {
                "description": "RoomPrice Create",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "RoomPrice"
                ],
                "operationId": "RoomPriceCreate",
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
                        "description": "type",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "price",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "floorId",
                        "name": "floorId",
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
        "/api/cms/v1/room/availibility-room": {
            "get": {
                "description": "Room Get Availibility Room",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Room"
                ],
                "operationId": "RoomGetAvailibilityRoom",
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
                        "description": "startDate",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "endDate",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "floorNumber",
                        "name": "floorNumber",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "roomNumber",
                        "name": "roomNumber",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "roomTypeName",
                        "name": "roomTypeName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "startFloorPrice",
                        "name": "startPrice",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "endfloorPrice",
                        "name": "endPrice",
                        "in": "query"
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
        "/api/cms/v1/room/create": {
            "post": {
                "description": "Room Create",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Room"
                ],
                "operationId": "RoomCreate",
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
                        "description": "floorId",
                        "name": "floorId",
                        "in": "formData",
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
                        "type": "integer",
                        "description": "number",
                        "name": "number",
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
        "/api/cms/v1/user/create": {
            "post": {
                "description": "User Create",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "User"
                ],
                "operationId": "UserCreate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
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
        "/api/cms/v1/user/login": {
            "post": {
                "description": "User Login",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "User"
                ],
                "operationId": "UserLogin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
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
        "requests.CreateBookingDetailRequest": {
            "type": "object",
            "required": [
                "roomId"
            ],
            "properties": {
                "roomId": {
                    "type": "string"
                }
            }
        },
        "requests.CreateBookingRequest": {
            "type": "object",
            "required": [
                "BookedBy",
                "bookingDetails",
                "downPayment",
                "endDate",
                "startDate"
            ],
            "properties": {
                "BookedBy": {
                    "type": "string"
                },
                "bookingDetails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/requests.CreateBookingDetailRequest"
                    }
                },
                "downPayment": {
                    "type": "integer"
                },
                "endDate": {
                    "type": "string"
                },
                "isTimeRulesAgree": {
                    "type": "boolean"
                },
                "startDate": {
                    "type": "string"
                }
            }
        },
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
