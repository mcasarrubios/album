package dao

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mcasarrubios/album/test"
)

type MockDB struct{}

var dao = New(&MockDB{})
var photos = []Photo{
	{
		URL: "http://my-photo.jpg",
		QueryInput: QueryInput{
			ID: "111",
			CreateInput: CreateInput{
				Tags:        []string{"tag-111-A", "tag-111-B"},
				Description: "Awesome description 1",
				Date:        "2008-09-15T15:53:00+05:00",
			},
		},
	}, {
		URL: "http://my-photo2.jpg",
		QueryInput: QueryInput{
			ID: "222",
			CreateInput: CreateInput{
				Tags:        []string{"tag-222-A", "tag-222-B"},
				Description: "Awesome description 2",
				Date:        "2009-09-15T15:53:00+05:00",
			},
		},
	},
}

func (db *MockDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return new(dynamodb.PutItemOutput), nil
}

func TestCreatePhoto(t *testing.T) {
	input := CreateInput{
		AlbumID:     photos[0].AlbumID,
		Tags:        photos[0].Tags,
		Description: photos[0].Description,
		Date:        photos[0].Date,
	}
	photo, err := dao.Create(input, photos[0].URL)
	test.Ok(t, err)
	expected := input.photo(photo.ID, photos[0].URL)
	expected.ID = photo.ID
	test.Equals(t, photos[0], photo)
}

// func TestListPhoto(t *testing.T) {
// 	query := QueryInput{
// 		AlbumID:     "1",
// 		Tags:        []string{"tag-1", "tag-2"},
// 		Description: "description",
// 		Date:        "2008-09-15T15:53:00+05:00",
// 	}
// 	photo, err := dao.Get(input, "http://my-photo.jpg")
// 	test.Ok(t, err)
// 	expected := input.photo(photo.ID, "http://my-photo.jpg")
// 	expected.ID = photo.ID
// 	test.Equals(t, expected, photo)
// }
