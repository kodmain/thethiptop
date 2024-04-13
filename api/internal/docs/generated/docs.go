// Package generated Code generated by swaggo/swag. DO NOT EDIT
package generated

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
        "/client": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Models"
                ],
                "operationId": "client.Find",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "put": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Models"
                ],
                "operationId": "client.UpdateComplete",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "delete": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Models"
                ],
                "operationId": "client.Delete",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Models"
                ],
                "operationId": "client.UpdatePartial",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/client/:id": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Models"
                ],
                "operationId": "client.FindOne",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/client/password/renew": {
            "post": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Password"
                ],
                "operationId": "client.Renew",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/client/password/reset": {
            "post": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Password"
                ],
                "operationId": "client.Reset",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/sign/in": {
            "post": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sign"
                ],
                "operationId": "client.SignIn",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/sign/out": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sign"
                ],
                "operationId": "jwt.Auth =\u003e client.SignOut",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/sign/renew": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sign"
                ],
                "operationId": "client.SignRenew",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/sign/up": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sign"
                ],
                "operationId": "client.SignUp",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Client created"
                    },
                    "400": {
                        "description": "Invalid email or password"
                    },
                    "409": {
                        "description": "Client already exists"
                    }
                }
            }
        },
        "/status/healthcheck": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Show the status of server.",
                "operationId": "status.HealthCheck",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/status/ip": {
            "get": {
                "description": "get the ip of user.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Show the ip of user.",
                "operationId": "status.IP",
                "responses": {
                    "200": {
                        "description": "OK"
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
