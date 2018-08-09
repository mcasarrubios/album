package photo

import (
	"reflect"
	"testing"

	"github.com/mcasarrubios/my-album/photo/dao"
)

func TestCreatePhoto(t *testing.T) {
	input := dao.CreateInput{
		AlbumID:     "1",
		Tags:        []string{"tag-1", "tag-2"},
		Description: "description",
		Date:        "2008-09-15T15:53:00+05:00",
	}
	photo, err := dao.Create(input)

	if err != nil {
		t.Fatalf("Error creating photo %v", err.Error())
	}

	expected := dao.Photo(input)
	expected.ID = photo.ID

	if !reflect.DeepEqual(expected, photo) {
		t.Fatalf("Error creating photo. Expected: %v, Actual: %v", expected, photo)
	}
}
