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
		ID:  "111",
		BasicPhoto: BasicPhoto{
			AlbumID:     "1",
			Tags:        []string{"tag-111-A", "tag-111-B"},
			Description: "Awesome description 1",
			Date:        "2008-09-15T15:53:00+05:00",
		},
	}, {
		URL: "http://my-photo2.jpg",
		ID:  "222",
		BasicPhoto: BasicPhoto{
			AlbumID:     "1",
			Tags:        []string{"tag-222-A", "tag-222-B"},
			Description: "Awesome description 2",
			Date:        "2009-09-15T15:53:00+05:00",
		},
	},
}

// func queryOutput() *dynamodb.QueryOutput {
// 	var items []map[string]*dynamodb.AttributeValue

// 	for i, photo := photos {
// 		aaa := common.MapStruct(photo)
// 		item := make(map[string]*dynamodb.AttributeValue)
// 		item["URL"] =
// 		items = items.append(make)
// 	}

// 	items
// 	return &dynamodb.QueryOutput{
// 		Items:
// 	}
// }

func (db *MockDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return new(dynamodb.PutItemOutput), nil
}

func (db *MockDB) Query(query *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return new(dynamodb.QueryOutput), nil
}

func TestCreatePhoto(t *testing.T) {
	input := CreateInput{
		BasicPhoto: BasicPhoto{
			AlbumID:     photos[0].AlbumID,
			Tags:        photos[0].Tags,
			Description: photos[0].Description,
			Date:        photos[0].Date,
		},
	}
	photo, err := dao.Create(input, photos[0].URL)
	test.Ok(t, err)
	expected := input.photo(photo.ID, photos[0].URL)
	test.Equals(t, expected, photo)
}

// func TestListPhoto(t *testing.T) {
// 	query := QueryInput{
// 		BasicPhoto: BasicPhoto{AlbumID: "1"},
// 	}
// 	actual, err := dao.List(query)
// 	test.Ok(t, err)
// 	test.Equals(t, photos, actual)
// }
