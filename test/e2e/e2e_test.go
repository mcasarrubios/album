package e2e

import (
	"fmt"
	"os/exec"
	"testing"
	"time"

	"gopkg.in/h2non/baloo.v3"
)

func init() {
	cmd := exec.Command("up", "start")
	cmd.Dir = "../.."
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(300 * time.Millisecond)
}

// test stores the HTTP testing client preconfigured
var test = baloo.New("http://localhost:3000")

const schema = `{
	"title": "Query Output",
	"type": "object",
	"properties": {
	  "items": {
		"type": "array",
		"properties": {
			"id": {
				"type": "string"
			},
			"description": {
				"type": "string"
			}
		},
		"required": ["id", "description"]
	  },
	  "lastKey": {
		"type": "string"
	  }
	},
	"required": ["items"]
  }`

func TestGetPhotos(t *testing.T) {
	test.Get("/").
		AddQuery("albumId", "1").
		AddQuery("project", "description").
		// AddQuery("project", "tags").
		Expect(t).
		Status(200).
		Type("json").
		// JSON(map[string]string{"description": "foo"}).
		JSONSchema(schema).
		Done()
}
