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
            "name": "Johnathan W",
            "email": "johnathan@w4r.dev"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/produce": {
            "get": {
                "description": "Get all produce entities.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produce"
                ],
                "summary": "Returns all produce entities",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Produce"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Uploads produce entities.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produce"
                ],
                "summary": "Uploads produce entities.",
                "parameters": [
                    {
                        "description": "Account Info",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Produce"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Produce"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/produce/{id}": {
            "get": {
                "description": "Get a produce entity.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produce"
                ],
                "summary": "Returns a produce entity based on the id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Produce ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Produce"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a produce entity.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produce"
                ],
                "summary": "Deletes a produce entity based on the id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Produce ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Produce"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal service error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/search": {
            "get": {
                "description": "Searches for a produce entity.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produce"
                ],
                "summary": "Searches for a produce entity based on the name or code.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name search",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "code search",
                        "name": "code",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Produce"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Produce": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "A12T-4GH7-QPL9-3N4M"
                },
                "id": {
                    "type": "integer",
                    "format": "int32",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "bannana"
                },
                "price": {
                    "type": "number",
                    "format": "float32",
                    "example": 3.32
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:23234",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Produce Store API",
	Description:      "This is a sample produce store server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
