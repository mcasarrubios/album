package dao

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DBMock struct {
}

func (db *DBMock) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemInput, error) {

}

func TestCreatePhoto(t *testing.T) {
	input := CreateInput{
		AlbumID:     "1",
		Tags:        []string{"tag-1", "tag-2"},
		Description: "description",
		Date:        "2008-09-15T15:53:00+05:00",
	}

	dao := New()
	photo, err := dao.Create(input)
	test.Ok(t, err)
	expected := input.photo(photo.ID)
	expected.ID = photo.ID
	test.Equals(t, expected, photo)
}

func (in CreateInput) photo(id string) *Photo {
	p := new(Photo)
	p.AlbumID = in.AlbumID
	p.Tags = in.Tags
	p.Description = in.Description
	p.Date = in.Date
	p.ID = id
	return p
}
