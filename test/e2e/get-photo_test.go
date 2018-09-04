package e2e

import (
	"testing"

	baloo "gopkg.in/h2non/baloo.v3"
)

func getPhoto(t *testing.T) {
	// Setup
	readFeed("photos", &photos)
	test = baloo.New(apiURL())

	t.Run("oneOk", oneOk)
	t.Run("oneNil", oneNil)
	// t.Run("projection", projection)

}

func oneOk(t *testing.T) {
	test.Get("/" + photos[0].ID).
		Expect(t).
		Status(200).
		Type("json").
		JSON(photos[0]).
		Done()
}

func oneNil(t *testing.T) {
	test.Get("/wrongID").
		Expect(t).
		Status(404).
		Done()
}
