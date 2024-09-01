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
        "/api/blogs": {
            "post": {
                "description": "Create a new blog with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "Create a new blog",
                "parameters": [
                    {
                        "description": "Blog data",
                        "name": "blog",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BlogCreateRequestDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/blogs-admin": {
            "get": {
                "description": "Get the list of all blogs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "Get all blogs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Blog"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new blog with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "Create a new blog",
                "parameters": [
                    {
                        "description": "Blog data",
                        "name": "blog",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Blog"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Blog"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/blogs-admin/{id}": {
            "get": {
                "description": "Get details of a blog by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "Get a blog by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Blog ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Blog"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a blog by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "Update an existing blog",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Blog ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Blog data",
                        "name": "blog",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Blog"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Blog"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a blog by its ID",
                "tags": [
                    "Blog"
                ],
                "summary": "Delete a blog",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Blog ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/blogs/": {
            "get": {
                "description": "List all blogs under a specific category identified by its slug",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "List blogs by category slug",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Blog"
                            }
                        }
                    }
                }
            }
        },
        "/api/blogs/recent-post": {
            "get": {
                "description": "Get the most recent and popular blog posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "blog recent post",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.RecentPostBlogDto"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/blogs/{categoriesSlug}": {
            "get": {
                "description": "List all blogs under a specific category identified by its slug",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "List blogs by category slug",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category Slug",
                        "name": "categoriesSlug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Blog"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BlogCreateRequestDto": {
            "type": "object",
            "required": [
                "blogContent",
                "blogTitle",
                "minRead",
                "published"
            ],
            "properties": {
                "blogContent": {
                    "type": "string"
                },
                "blogTitle": {
                    "type": "string",
                    "maxLength": 500
                },
                "categoryIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "isPin": {
                    "type": "boolean"
                },
                "minRead": {
                    "type": "integer",
                    "minimum": 1
                },
                "published": {
                    "type": "boolean"
                },
                "slug": {
                    "type": "string"
                },
                "summary": {
                    "type": "string",
                    "maxLength": 500
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "thumbnail": {
                    "type": "string",
                    "maxLength": 255
                }
            }
        },
        "dto.RecentPostBlogDto": {
            "type": "object",
            "properties": {
                "blogTitle": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "timeAgo": {
                    "type": "string"
                }
            }
        },
        "handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "fields": {
                    "description": "This field will hold the detailed validation errors",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Blog": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/models.User"
                },
                "authorID": {
                    "type": "integer"
                },
                "blogContent": {
                    "type": "string"
                },
                "blogTitle": {
                    "type": "string"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Category"
                    }
                },
                "countViewer": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "isPin": {
                    "type": "boolean"
                },
                "minRead": {
                    "type": "integer"
                },
                "parent": {
                    "$ref": "#/definitions/models.Blog"
                },
                "parentID": {
                    "type": "integer"
                },
                "published": {
                    "type": "boolean"
                },
                "slug": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Tag"
                    }
                },
                "thumbnail": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "blogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Blog"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "slug": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Tag": {
            "type": "object",
            "properties": {
                "blogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Blog"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "title": {
                    "description": "Title     string    ` + "`" + `gorm:\"size:100;not null;unique\"` + "`" + `",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "bio": {
                    "type": "string"
                },
                "confirmationToken": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isVerified": {
                    "type": "boolean"
                },
                "password": {
                    "description": "Provide a default value",
                    "type": "string"
                },
                "profileImage": {
                    "type": "string"
                },
                "resetToken": {
                    "type": "string"
                },
                "top3Count": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userName": {
                    "description": "Provide a default value",
                    "type": "string"
                },
                "verifiedByAdmin": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "backend service for blog api",
	Description:      "backend service api restfull using Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
