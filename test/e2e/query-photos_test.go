package e2e

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/mcasarrubios/album/photo/dao"
	testUtils "github.com/mcasarrubios/album/test"
	"github.com/mcasarrubios/album/test/e2e/schemas"
	baloo "gopkg.in/h2non/baloo.v3"
)

func encodeKey(ph dao.Photo) string {
	key := dao.KeyPhoto{
		AlbumID: ph.AlbumID,
		Date:    ph.Date,
	}
	js, _ := json.Marshal(key)
	return base64.URLEncoding.EncodeToString(js)
}

func getPhotos(t *testing.T) {
	// Setup
	readFeed("photos", &photos)
	test = baloo.New(apiURL())

	t.Run("limit", limit)
	t.Run("projection", projection)
	t.Run("lastKey", lastKey)
	t.Run("startKey", startKey)
	t.Run("sort", sort)

}

func limit(t *testing.T) {
	opts := schemas.Placeholders{
		MinItems: "1",
		MaxItems: "1",
	}
	schema, err := schemas.Schema("query-output", opts)
	testUtils.Ok(t, err)

	data := dao.QueryOutput{
		Items:   []dao.Photo{photos[0]},
		LastKey: encodeKey(photos[0]),
	}

	test.Get("/").
		AddQuery("limit", "1").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		JSON(data).
		Done()
}

func projection(t *testing.T) {
	opts := schemas.Placeholders{
		ItemFields: []string{"description", "tags", "id"},
		MinItems:   "1",
		MaxItems:   "100",
	}
	schema, err := schemas.Schema("query-output", opts)
	testUtils.Ok(t, err)

	ph := dao.Photo{
		BasicPhoto: dao.BasicPhoto{
			Description: photos[0].Description,
			Tags:        photos[0].Tags,
		},
		ID: photos[0].ID,
	}

	data := dao.QueryOutput{
		Items:   []dao.Photo{ph},
		LastKey: encodeKey(photos[0]),
	}
	test.Get("/").
		AddQuery("limit", "1").
		AddQuery("fields", "description").
		AddQuery("fields", "tags").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		JSON(data).
		Done()
}

func lastKey(t *testing.T) {
	opts := schemas.Placeholders{
		RootFields: []string{"lastKey"},
		MinItems:   "1",
		MaxItems:   "100",
	}
	schema, err := schemas.Schema("query-output", opts)
	testUtils.Ok(t, err)

	data := dao.QueryOutput{
		Items:   []dao.Photo{photos[0]},
		LastKey: encodeKey(photos[0]),
	}

	test.Get("/").
		AddQuery("limit", "1").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		JSON(data).
		Done()
}

func startKey(t *testing.T) {
	opts := schemas.Placeholders{
		RootFields: []string{"lastKey"},
		MinItems:   "1",
		MaxItems:   "100",
	}
	schema, err := schemas.Schema("query-output", opts)
	testUtils.Ok(t, err)

	data := dao.QueryOutput{
		Items:   []dao.Photo{photos[1]},
		LastKey: encodeKey(photos[1]),
	}
	test.Get("/").
		AddQuery("startKey", encodeKey(photos[0])).
		AddQuery("limit", "1").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		JSON(data).
		Done()
}

func sort(t *testing.T) {
	opts := schemas.Placeholders{
		MinItems: "1",
		MaxItems: "100",
	}
	schema, err := schemas.Schema("query-output", opts)
	testUtils.Ok(t, err)

	data := dao.QueryOutput{
		Items: []dao.Photo{photos[1], photos[0]},
	}

	test.Get("/").
		AddQuery("sortDesc", "true").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		JSON(data).
		Done()
}
