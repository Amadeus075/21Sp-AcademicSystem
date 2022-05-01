// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/course": {
            "get": {
                "description": "获取课程 可以自由添加筛选属性",
                "tags": [
                    "course"
                ],
                "summary": "获取课程",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 4,
                        "description": "学分",
                        "name": "credit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "计算机",
                        "description": "所属院系",
                        "name": "department",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "数据库原理",
                        "description": "课名",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "0121",
                        "description": "课号",
                        "name": "number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "老师A",
                        "description": "教师姓名",
                        "name": "teacher_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "22-冬季学期",
                        "description": "学期",
                        "name": "term",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "创建课程",
                "tags": [
                    "course"
                ],
                "summary": "创建课程",
                "parameters": [
                    {
                        "description": "course 实例",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/course/{id}": {
            "get": {
                "description": "根据id获取课程信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "根据id获取课程信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    }
                }
            },
            "put": {
                "description": "整体更新课程信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "整体更新课程信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "course 实例",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login/admin": {
            "post": {
                "description": "以管理员身份登录",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "以管理员身份登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "number",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/login/student": {
            "post": {
                "description": "以学生身份登录",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "以学生身份登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学号",
                        "name": "number",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Student"
                        }
                    }
                }
            }
        },
        "/selection": {
            "post": {
                "description": "创建选课",
                "tags": [
                    "selection"
                ],
                "summary": "创建选课",
                "parameters": [
                    {
                        "description": "选课情况",
                        "name": "selection",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Selection"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除选课",
                "tags": [
                    "selection"
                ],
                "summary": "删除选课",
                "parameters": [
                    {
                        "description": "选课情况",
                        "name": "selection",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Selection"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/student": {
            "get": {
                "description": "获取学生 可以自由添加筛选属性",
                "tags": [
                    "student"
                ],
                "summary": "获取学生",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 21,
                        "name": "age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "计算机",
                        "description": "所属院系",
                        "name": "department",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "王二",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "0198",
                        "name": "number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "123",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "男",
                        "name": "sex",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "创建学生",
                "tags": [
                    "student"
                ],
                "summary": "创建学生",
                "parameters": [
                    {
                        "description": "student 实例",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Student"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/student/{id}": {
            "get": {
                "description": "根据id获取学生信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "student"
                ],
                "summary": "根据id获取学生信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Student"
                        }
                    }
                }
            }
        },
        "/student/{id}/course": {
            "get": {
                "description": "获取指定学生的所有课程",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "student",
                    "course"
                ],
                "summary": "获取指定学生的所有课程",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "是否有成绩 不写即全部返回",
                        "name": "hasScore",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.CourseByStuResponse"
                            }
                        }
                    }
                }
            }
        },
        "/student/{student_id}/course/{course_id}": {
            "put": {
                "description": "更新课程成绩",
                "tags": [
                    "selection",
                    "admin"
                ],
                "summary": "更新课程成绩",
                "parameters": [
                    {
                        "type": "string",
                        "description": "student ID",
                        "name": "student_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "course ID",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "score",
                        "name": "score",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/test/generate-examples": {
            "get": {
                "description": "生成样例数据",
                "tags": [
                    "test"
                ],
                "summary": "生成样例数据",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/test/ping": {
            "get": {
                "description": "pong me!",
                "tags": [
                    "test"
                ],
                "summary": "ping!",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Course": {
            "type": "object",
            "properties": {
                "credit": {
                    "description": "学分",
                    "type": "integer",
                    "example": 4
                },
                "department": {
                    "description": "所属院系",
                    "type": "string",
                    "example": "计算机"
                },
                "name": {
                    "description": "课名",
                    "type": "string",
                    "example": "数据库原理"
                },
                "number": {
                    "description": "课号",
                    "type": "string",
                    "example": "0121"
                },
                "selections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Selection"
                    }
                },
                "teacher_name": {
                    "description": "教师姓名",
                    "type": "string",
                    "example": "老师A"
                },
                "term": {
                    "description": "学期",
                    "type": "string",
                    "example": "22-冬季学期"
                }
            }
        },
        "model.CourseByStuResponse": {
            "type": "object",
            "properties": {
                "credit": {
                    "description": "学分",
                    "type": "integer"
                },
                "department": {
                    "description": "所属院系",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "课名",
                    "type": "string"
                },
                "number": {
                    "description": "课号",
                    "type": "string"
                },
                "score": {
                    "description": "成绩",
                    "type": "integer"
                },
                "student_name": {
                    "description": "学生名",
                    "type": "string"
                },
                "teacher_name": {
                    "description": "老师名",
                    "type": "string"
                },
                "term": {
                    "description": "学期",
                    "type": "string"
                }
            }
        },
        "model.Selection": {
            "type": "object",
            "properties": {
                "courseID": {
                    "type": "integer"
                },
                "score": {
                    "description": "分数, -1表示未评分",
                    "type": "integer",
                    "example": 75
                },
                "studentID": {
                    "type": "integer"
                }
            }
        },
        "model.Student": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 21
                },
                "department": {
                    "description": "所属院系",
                    "type": "string",
                    "example": "计算机"
                },
                "name": {
                    "type": "string",
                    "example": "王二"
                },
                "number": {
                    "type": "string",
                    "example": "0198"
                },
                "password": {
                    "type": "string",
                    "example": "123"
                },
                "selections": {
                    "description": "选课情况",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Selection"
                    }
                },
                "sex": {
                    "type": "string",
                    "example": "男"
                }
            }
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
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Academic System",
	Description: "学生选课系统",
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
	swag.Register(swag.Name, &s{})
}
