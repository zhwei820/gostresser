// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-03-03 00:00:47.498669651 +0800 CST m=+0.029323573

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "Swagger Example API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "localhost:8179",
    "basePath": "/",
    "paths": {
        "/conf": {
            "get": {
                "description": "ListConf",
                "produces": [
                    "application/json"
                ],
                "summary": "ListConf",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "CreateConf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreateConf",
                "parameters": [
                    {
                        "description": "create BaseConf",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/config.BaseConf"
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
        "/conf/{id}": {
            "get": {
                "description": "DetailConf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "DetailConf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "id",
                        "in": "path",
                        "required": true
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
            },
            "put": {
                "description": "UpdateConf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "UpdateConf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "create BaseConf",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/config.BaseConf"
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
            },
            "delete": {
                "description": "RemoveConf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "RemoveConf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "id",
                        "in": "path",
                        "required": true
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
        }
    },
    "definitions": {
        "config.BaseConf": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "concurrency": {
                    "type": "integer"
                },
                "cpus": {
                    "type": "integer"
                },
                "disablecompression": {
                    "type": "boolean"
                },
                "disablekeepalives": {
                    "type": "boolean"
                },
                "disableredirects": {
                    "type": "boolean"
                },
                "duration": {
                    "type": "integer"
                },
                "h2": {
                    "type": "boolean"
                },
                "proxyaddr": {
                    "type": "string"
                },
                "reqconfs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.ReqConf"
                    }
                },
                "timeout": {
                    "type": "integer"
                }
            }
        },
        "config.HeaderSlice": {
            "type": "array",
            "items": {}
        },
        "config.ReqConf": {
            "type": "object",
            "properties": {
                "accept": {
                    "type": "string"
                },
                "authheader": {
                    "type": "string"
                },
                "body": {
                    "type": "string"
                },
                "contenttype": {
                    "type": "string"
                },
                "headers": {
                    "type": "object",
                    "$ref": "#/definitions/config.HeaderSlice"
                },
                "hostheader": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "reqFreq": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}