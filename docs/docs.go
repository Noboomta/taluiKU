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
    "basePath": "/",
    "consumes": [
      "application/json"
    ],
    "info": {
      "license": {
        "name": "ISC"
      },
      "title": "TaluiKU Api docs",
      "version": "0.1.0"
    },
    "paths": {
      "/complete/getAllTalui": {
        "get": {
          "operationId": "completeAlle",
          "produces": [
            "application/json"
          ],
          "responses": {
            "200": {
              "description": "Ok",
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/definitions/taluicomplete"
                }
              }
            }
          },
          "description": "get array of TaluiComplete"
        }
      },
      "/on/getAllTalui": {
        "get": {
          "operationId": "onAll",
          "produces": [
            "application/json"
          ],
          "responses": {
            "200": {
              "description": "Ok",
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/definitions/taluion"
                }
              }
            }
          },
          "description": "get array of TaluiOn"
        }
      },
      "/complete/insertTalui": {
        "post": {
          "operationId": "Insert Complete",
          "produces": [
            "application/json"
          ],
          "parameters": [
            {
              "name": "entry",
              "in": "query",
              "description": "entry",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            },
            {
              "name": "dest",
              "in": "query",
              "description": "dest",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            },
            {
              "name": "line",
              "in": "query",
              "description": "line",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            }
          ],
          "responses": {
            "200": {
              "description": "Ok"
            },
            "400": {
              "description": "Error"
            }
          },
          "description": "Insert TaluiComplete into database"
        }
      },
      "/on/insertTalui": {
        "post": {
          "operationId": "InsertOn",
          "produces": [
            "application/json"
          ],
          "parameters": [
            {
              "name": "entry",
              "in": "query",
              "description": "entry",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            },
            {
              "name": "dest",
              "in": "query",
              "description": "dest",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            },
            {
              "name": "line",
              "in": "query",
              "description": "line",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            }
          ],
          "responses": {
            "200": {
              "description": "Ok"
            },
            "400": {
              "description": "Error"
            }
          },
          "description": "Insert TaluiOn into database"
        }
      },
      "/arriveAt": {
        "post": {
          "operationId": "arr",
          "produces": [
            "application/json"
          ],
          "parameters": [
            {
              "name": "dest",
              "in": "query",
              "description": "dest",
              "required": false,
              "type": "string",
              "items": {
                "type": "string"
              }
            },
            {
              "name": "line",
              "in": "query",
              "description": "line",
              "required": false,
              "type": "string",
              "items": {
                "type": "string"
              }
            }
          ],
          "responses": {
            "200": {
              "description": "Ok"
            },
            "400": {
              "description": "Error"
            }
          },
          "description": "arrive at for talui when arrive at station"
        }
      },
      "/station/using/entry": {
        "get": {
          "operationId": "getE",
          "produces": [
            "application/json"
          ],
          "responses": {
            "200": {
              "description": "Ok",
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/definitions/StationEntryUsing"
                }
              }
            }
          },
          "description": "get array of sum of station using (entry) separate by hour"
        }
      },
      "/station/using/dest": {
        "get": {
          "operationId": "getD",
          "produces": [
            "application/json"
          ],
          "responses": {
            "200": {
              "description": "Ok",
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/definitions/StationDestUsing"
                }
              }
            }
          },
          "description": "get array of sum of station using (dest) separate by hour"
        }
      },
      "/station/find/entry": {
        "get": {
          "operationId": "getUE",
          "produces": [
            "application/json"
          ],
          "parameters": [
            {
              "name": "station",
              "in": "query",
              "description": "station",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            },
            {
              "name": "time",
              "in": "query",
              "description": "time",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            }
          ],
          "responses": {
            "200": {
              "description": "Ok",
              "schema": {
                "$ref": "#/definitions/StationEntryUsing"
              }
            }
          },
          "description": "find one of station using (entry) separate by hour"
        }
      },
      "/station/find/dest": {
        "get": {
          "operationId": "getUD",
          "produces": [
            "application/json"
          ],
          "parameters": [
            {
              "name": "station",
              "in": "query",
              "description": "station",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            },
            {
              "name": "time",
              "in": "query",
              "description": "time",
              "required": true,
              "type": "string",
              "items": {
                "type": "string"
              }
            }
          ],
          "responses": {
            "200": {
              "description": "Ok",
              "schema": {
                "$ref": "#/definitions/StationDestUsing"
              }
            }
          },
          "description": "find one of station using (dest) separate by hour"
        }
      }
    },
    "produces": [
      "application/json"
    ],
    "swagger": "2.0",
    "securityDefinitions": {},
    "host": "talui-ku-server.herokuapp.com",
    "definitions": {
      "taluicomplete": {
        "type": "object",
        "properties": {
          "id": {
            "description": "id of talui",
            "type": "integer"
          },
          "entry": {
            "description": "entry station.",
            "type": "string"
          },
          "entry_ts": {
            "description": "entry time.",
            "type": "string"
          },
          "dest": {
            "description": "destination station.",
            "type": "string"
          },
          "dest_ts": {
            "description": "destination time.",
            "type": "string"
          },
          "line": {
            "description": "line of talui.",
            "type": "string"
          }
        }
      },
      "taluion": {
        "type": "object",
        "properties": {
          "id": {
            "description": "id of talui",
            "type": "integer"
          },
          "entry": {
            "description": "entry station.",
            "type": "string"
          },
          "entry_ts": {
            "description": "entry time.",
            "type": "string"
          },
          "dest": {
            "description": "destination station.",
            "type": "string"
          },
          "line": {
            "description": "line of talui.",
            "type": "string"
          }
        }
      },
      "StationEntryUsing": {
        "type": "object",
        "properties": {
          "station": {
            "description": "station name.",
            "type": "string"
          },
          "value": {
            "description": "number of using",
            "type": "integer"
          },
          "time": {
            "description": "entry time.",
            "type": "string"
          }
        }
      },
      "StationDestUsing": {
        "type": "object",
        "properties": {
          "station": {
            "description": "station name.",
            "type": "string"
          },
          "value": {
            "description": "number of using",
            "type": "integer"
          },
          "time": {
            "description": "dest time.",
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
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "talui-ku-server.herokuapp.com/",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Talui Api docs",
	Description: "This is Talui api-docs",
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
