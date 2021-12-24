// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://swagger.io/terms/",
        "contact": {
            "name": "MyPoints Team Support",
            "email": "zenhanprogram@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/:id/agent": {
            "post": {
                "description": "create agent account by admins.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Agent"
                ],
                "summary": "admins can create agent account with this api",
                "parameters": [
                    {
                        "description": "body request",
                        "name": "signUp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_yossdev_mypoints-rest-api_src_agents_dto.SignUpReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github.com_yossdev_mypoints-rest-api_src_agents_dto.AccountCreated"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "check admins by checking given email and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "check admins by given email return jwt token if successfully signIn",
                "parameters": [
                    {
                        "description": "body request",
                        "name": "signIn",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_yossdev_mypoints-rest-api_src_admins_dto.SignInReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.Token"
                        }
                    }
                }
            }
        },
        "/admin/signup": {
            "post": {
                "description": "create admin account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "admins can create from register page",
                "parameters": [
                    {
                        "description": "body request",
                        "name": "signUp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_yossdev_mypoints-rest-api_src_admins_dto.SignUpReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github.com_yossdev_mypoints-rest-api_src_admins_dto.AccountCreated"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "check agent by checking given email and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Agent"
                ],
                "summary": "check agent by given email return jwt token if successfully signIn",
                "parameters": [
                    {
                        "description": "body request",
                        "name": "signIn",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_yossdev_mypoints-rest-api_src_agents_dto.SignInReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.Token"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.Token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "sub": {
                    "type": "string"
                }
            }
        },
        "github.com_yossdev_mypoints-rest-api_src_admins_dto.AccountCreated": {
            "type": "object",
            "properties": {
                "rows_affected": {
                    "type": "integer"
                }
            }
        },
        "github.com_yossdev_mypoints-rest-api_src_admins_dto.SignInReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "github.com_yossdev_mypoints-rest-api_src_admins_dto.SignUpReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "img": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "github.com_yossdev_mypoints-rest-api_src_agents_dto.AccountCreated": {
            "type": "object",
            "properties": {
                "rows_affected": {
                    "type": "integer"
                }
            }
        },
        "github.com_yossdev_mypoints-rest-api_src_agents_dto.SignInReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "github.com_yossdev_mypoints-rest-api_src_agents_dto.SignUpReq": {
            "type": "object",
            "properties": {
                "admin_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "img": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
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

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "MyPoints API",
	Description: "This is an auto-generated API Docs.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
