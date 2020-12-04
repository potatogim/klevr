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
        "/agents/handshake": {
            "put": {
                "description": "에이전트 프로세스가 기동시 최초 한번 handshake를 요청하여 에이전트 정보 등록 및 에이전트 실행에 필요한 실행 정보를 반환한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "에이전트의 handshake 요청을 받아 처리한다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API KEY",
                        "name": "X-API-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "AGENT KEY",
                        "name": "X-AGENT-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ZONE ID",
                        "name": "X-ZONE-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "agent 정보",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.Body"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Body"
                        }
                    }
                }
            }
        },
        "/agents/reports/{agentKey}": {
            "get": {
                "description": "secondary 에이전트의 primary 에이전트 상태 확인 요청을 받아 primary 재선출 및 primary 정보를 반환한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "secondary 에이전트의 primary 상태 확인 요청을 처리한다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API KEY",
                        "name": "X-API-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "AGENT KEY",
                        "name": "X-AGENT-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ZONE ID",
                        "name": "X-ZONE-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "agent key",
                        "name": "agentKey",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "agent 정보",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.Body"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Body"
                        }
                    }
                }
            }
        },
        "/agents/{agentKey}": {
            "put": {
                "description": "primary 에이전트의 polling 요청을 받아 primary 에이전트의 실행정보 갱신, nodes 정보 갱신, task 할당 및 상태 업데이트를 수행한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "primary 에이전트의 polling 요청을 받아 처리한다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API KEY",
                        "name": "X-API-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "AGENT KEY",
                        "name": "X-AGENT-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ZONE ID",
                        "name": "X-ZONE-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "agent key",
                        "name": "agentKey",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "agent 정보",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.Body"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Body"
                        }
                    }
                }
            }
        },
        "/inner/commands": {
            "get": {
                "description": "Klevr에서 사용할 수 있는 예약어 커맨드 정보를 반환한다. 사용자는 이 정보를 토대로 task를 생성하여 요청할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "예약어 커맨드 정보를 반환한다.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/manager.ReservedCommand"
                            }
                        }
                    }
                }
            }
        },
        "/inner/groups": {
            "get": {
                "description": "KLEVR ZONE 목록을 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "ZONE 목록을 조회한다.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/manager.AgentGroups"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "KLEVR ZONE을 생성한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "ZONE을 추가한다.",
                "parameters": [
                    {
                        "description": "AgentGroups model",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/manager.AgentGroups"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/manager.AgentGroups"
                        }
                    }
                }
            }
        },
        "/inner/groups/{groupID}": {
            "get": {
                "description": "KLEVR ZONE을 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "ZONE을 조회한다.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ZONE ID",
                        "name": "groupID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/manager.AgentGroups"
                        }
                    }
                }
            }
        },
        "/inner/groups/{groupID}/agents": {
            "get": {
                "description": "groupID에 해당하는 klevr zone의 모든 agent 정보를 반환한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "zone의 agent 목록을 반환한다.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ZONE ID",
                        "name": "groupID",
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
                                "$ref": "#/definitions/common.Agent"
                            }
                        }
                    }
                }
            }
        },
        "/inner/groups/{groupID}/apikey": {
            "get": {
                "description": "agent가 zone에 접속할 수 있는 API KEY를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "사용자 그룹의 API key를 조회한다.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ZONE ID",
                        "name": "groupID",
                        "in": "path",
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
            },
            "put": {
                "description": "agent가 zone에 접속할 수 있는 API KEY를 수정한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "사용자 그룹의 API key를 수정한다.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ZONE ID",
                        "name": "groupID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "API KEY",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            },
            "post": {
                "description": "agent가 zone에 접속할 수 있는 API KEY를 등록한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "사용자 그룹에 API key를 등록한다.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ZONE ID",
                        "name": "groupID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "API KEY",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/inner/groups/{groupID}/primary": {
            "get": {
                "description": "groupID에 해당하는 klevr zone의 primary agent 정보를 반환한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "primary agent 정보를 반환한다.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ZONE ID",
                        "name": "groupID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Agent"
                        }
                    }
                }
            }
        },
        "/inner/tasks": {
            "get": {
                "description": "검색조건에 해당하는 TASK 목록을 반환한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "TASK 목록을 반환한다.",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "description": "ZONE ID 배열",
                        "name": "groupID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "STATUS 배열",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "AGENT KEY 배열",
                        "name": "agentKey",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "TASK NAME 배열",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/common.KlevrTask"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "KlevrTask 모델에 기입된 ZONE의 AGENT에서 실행할 TASK를 등록한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "TASK를 등록한다.",
                "parameters": [
                    {
                        "description": "TASK",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.KlevrTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.KlevrTask"
                        }
                    }
                }
            }
        },
        "/inner/tasks/{taskID}": {
            "get": {
                "description": "taskID에 해당하는 TASK를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "TASK를 조회한다.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task id",
                        "name": "taskID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.KlevrTask"
                        }
                    }
                }
            },
            "delete": {
                "description": "agent에 전달되지 않은(hand-over 이전) task를 취소한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "TASK를 취소한다.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task id",
                        "name": "taskID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"canceld\\\":true/false}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/inner/variables": {
            "get": {
                "description": "TASK inline command에서 사용할 수 있는 시스템 변수 목록을 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "Klevr에서 제공하는 시스템 변수 목록을 조회한다.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/manager.KlevrVariable"
                            }
                        }
                    }
                }
            }
        },
        "/temp": {
            "get": {
                "description": "에이전트 프로세스가 기동시 최초 한번 handshake를 요청하여 에이전트 정보 등록 및 에이전트 실행에 필요한 실행 정보를 반환한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "temp"
                ],
                "summary": "에이전트의 handshake 요청을 받아 처리한다."
            }
        }
    },
    "definitions": {
        "common.Agent": {
            "type": "object",
            "properties": {
                "agentKey": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "lastAliveCheckTime": {
                    "type": "JSONTime"
                },
                "port": {
                    "type": "integer"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "common.Body": {
            "type": "object",
            "properties": {
                "agent": {
                    "type": "BodyAgent"
                },
                "me": {
                    "type": "Me"
                },
                "task": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.KlevrTask"
                    }
                }
            }
        },
        "common.KlevrTask": {
            "type": "object",
            "properties": {
                "agentKey": {
                    "type": "string"
                },
                "callbackUrl": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "JSONTime"
                },
                "cron": {
                    "type": "string"
                },
                "currentStep": {
                    "type": "integer"
                },
                "exeAgentChangeable": {
                    "type": "boolean"
                },
                "exeAgentKey": {
                    "type": "string"
                },
                "failedStep": {
                    "type": "integer"
                },
                "hasRecover": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "isFailedRecover": {
                    "type": "boolean"
                },
                "log": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parameter": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                },
                "schedule": {
                    "type": "JSONTime"
                },
                "showLog": {
                    "type": "boolean"
                },
                "status": {
                    "type": "TaskStatus"
                },
                "steps": {
                    "type": "array",
                    "items": {
                        "type": "KlevrTaskStep"
                    }
                },
                "taskType": {
                    "type": "TaskType"
                },
                "timeout": {
                    "type": "integer"
                },
                "totalStepCount": {
                    "type": "integer"
                },
                "untilRun": {
                    "type": "JSONTime"
                },
                "updatedAt": {
                    "type": "JSONTime"
                },
                "zoneId": {
                    "type": "integer"
                }
            }
        },
        "manager.AgentGroups": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "groupName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "platform": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "manager.KlevrVariable": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "example": {
                    "type": "string"
                },
                "length": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "manager.ReservedCommand": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "hasRecover": {
                    "type": "boolean"
                },
                "parameterModel": {
                    "type": "object"
                },
                "resultModel": {
                    "type": "object"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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