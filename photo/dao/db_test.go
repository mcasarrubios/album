package dao

import (
	"testing"

	"github.com/mcasarrubios/album/test"
)

func TestCreatePhoto(t *testing.T) {
	input := CreateInput{
		AlbumID:     "1",
		Tags:        []string{"tag-1", "tag-2"},
		Description: "description",
		Date:        "2008-09-15T15:53:00+05:00",
	}

	dao, err := New()
	test.Ok(t, err)
	photo, err := dao.Create(input)
	test.Ok(t, err)
	expected := input.photo(photo.ID)
	expected.ID = photo.ID
	test.Equals(t, expected, photo)
}
