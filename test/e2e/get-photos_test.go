package e2e

import (
	"testing"

	testUtils "github.com/mcasarrubios/album/test"
	"github.com/mcasarrubios/album/test/e2e/schemas"
	"gopkg.in/h2non/baloo.v3"
)

func init() {
	setup()
}

// test stores the HTTP testing client preconfigured
var test = baloo.New(apiURL())

func TestGetPhotosProjection(t *testing.T) {
	opts := schemas.Placeholders{
		Fields:   []string{"description", "tags"},
		MinItems: "1",
		MaxItems: "100",
	}
	schema, err := schemas.Schema("query-output", opts)
	testUtils.Ok(t, err)
	test.Get("/").
		AddQuery("fields", "description").
		AddQuery("fields", "tags").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		Done()
}

func TestGetPhotosLimit(t *testing.T) {
	opts := schemas.Placeholders{
		Fields:   []string{},
		MinItems: "1",
		MaxItems: "1",
	}
	schema, err := schemas.Schema("query-output", opts)
	testUtils.Ok(t, err)
	test.Get("/").
		AddQuery("limit", "1").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		Done()
}
