// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-07-06 17:15:08.77267891 -0500 -05 m=+13.988328398

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Leverage of authentication functionality",
        "title": "Go-Authenticator",
        "contact": {
            "name": "API Support",
            "url": "https://bixlabs.com/",
            "email": "jarrieta@bixlabs.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "/",
    "paths": {
        "/user/change-password": {
            "put": {
                "description": "It changes the password provided the old one and a new password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Change password functionality",
                "parameters": [
                    {
                        "description": "Change password Request",
                        "name": "changePassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/change_password.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/change_password.SwaggerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Attempts to authenticate the user with the given credentials.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login functionality",
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/login.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/login.SwaggerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    }
                }
            }
        },
        "/user/reset-password": {
            "put": {
                "description": "It resets your password given the correct code and new password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Reset password functionality",
                "parameters": [
                    {
                        "description": "Reset password Request",
                        "name": "resetPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/reset_password.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/reset_password.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "408": {
                        "description": "Request Timeout",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "504": {
                        "description": "Gateway Timeout",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    }
                }
            }
        },
        "/user/reset-password-request": {
            "put": {
                "description": "It enters into the flow of reset password sending an email with instructions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Forgot password request functionality",
                "parameters": [
                    {
                        "description": "Forgot password request",
                        "name": "resetPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/forgot_password.Request"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/forgot_password.SwaggerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Attempts to create a user provided the correct information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Signup functionality",
                "parameters": [
                    {
                        "description": "Signup Request",
                        "name": "signup",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/signup.Request"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/signup.SwaggerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ResponseWrapper"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "change_password.Request": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "newPassword": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "change_password.Response": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "object",
                    "$ref": "#/definitions/change_password.Result"
                }
            }
        },
        "change_password.Result": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "change_password.SwaggerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "result": {
                    "type": "object",
                    "$ref": "#/definitions/change_password.Response"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "forgot_password.Request": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "forgot_password.Result": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "forgot_password.SwaggerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "result": {
                    "type": "object",
                    "$ref": "#/definitions/forgot_password.Result"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "login.Request": {
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
        "login.Response": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "object",
                    "$ref": "#/definitions/login.Result"
                }
            }
        },
        "login.Result": {
            "type": "object"
        },
        "login.SwaggerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "result": {
                    "type": "object",
                    "$ref": "#/definitions/login.Response"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "reset_password.Request": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "newPassword": {
                    "type": "string"
                }
            }
        },
        "reset_password.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "result": {
                    "type": "object",
                    "$ref": "#/definitions/reset_password.Result"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "reset_password.Result": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "rest.ResponseWrapper": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "signup.Request": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "familyName": {
                    "type": "string"
                },
                "givenName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "secondFamilyName": {
                    "type": "string"
                },
                "secondName": {
                    "type": "string"
                }
            }
        },
        "signup.Response": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "object",
                    "$ref": "#/definitions/signup.Result"
                }
            }
        },
        "signup.Result": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "signup.SwaggerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "result": {
                    "type": "object",
                    "$ref": "#/definitions/signup.Response"
                },
                "status": {
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
