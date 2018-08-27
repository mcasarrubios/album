package dao

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mcasarrubios/album/test"
)

type MockDB struct{}

var dao = New(&MockDB{})
var photos = []Photo{
	{
		URL: "http://my-photo.jpg",
		ID:  "111",
		BasicPhoto: BasicPhoto{
			KeyPhoto: KeyPhoto{
				AlbumID: "1",
				Date:    "2008-09-15T15:53:00+05:00",
			},
			Tags:        []string{"tag-111-A", "tag-111-B"},
			Description: "Awesome description 1",
		},
	}, {
		URL: "http://my-photo2.jpg",
		ID:  "222",
		BasicPhoto: BasicPhoto{
			KeyPhoto: KeyPhoto{
				AlbumID: "1",
				Date:    "2008-09-25T15:53:00+05:00",
			},
			Tags:        []string{"tag-222-A", "tag-222-B"},
			Description: "Awesome description 2",
		},
	},
}

func dbQueryOutput() []map[string]*dynamodb.AttributeValue {
	var items []map[string]*dynamodb.AttributeValue
	for _, ph := range photos {
		item, _ := dynamodbattribute.MarshalMap(ph)
		items = append(items, item)
	}
	return items
}

func (db *MockDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return new(dynamodb.PutItemOutput), nil
}

func (db *MockDB) Query(query *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	items := dbQueryOutput()
	return &dynamodb.QueryOutput{
		Items: items,
	}, nil
}

func TestCreate(t *testing.T) {
	input := CreateInput{
		BasicPhoto: BasicPhoto{
			KeyPhoto: KeyPhoto{
				AlbumID: photos[0].AlbumID,
				Date:    photos[0].Date,
			},
			Tags:        photos[0].Tags,
			Description: photos[0].Description,
		},
	}
	photo, err := dao.Create(input, photos[0].URL)
	test.Ok(t, err)
	expected := input.photo(photo.ID, photos[0].URL)
	test.Equals(t, expected, photo)
}

func TestListRequiredFields(t *testing.T) {
	query := QueryInput{}
	_, err := dao.List(query)
	test.Equals(t, "Missing required fields", err.Error())
}

func TestListPhoto(t *testing.T) {
	query := QueryInput{
		Filter: FilterInput{AlbumID: "1"},
	}
	actual, err := dao.List(query)
	test.Ok(t, err)
	test.Equals(t, photos, actual.Items)
}
func TestGetPhoto(t *testing.T) {
	query := QueryInput{
		Filter: FilterInput{AlbumID: "1"},
	}
	actual, err := dao.List(query)
	test.Ok(t, err)
	input := GetInput{
		AlbumID: actual.Items[0].AlbumID,
		ID:      actual.Items[0].ID,
	}
	photo, err := dao.Get(input)
	test.Ok(t, err)
	test.Equals(t, input.ID, photo.ID)
}

// func TestUpdateDescriptionPhoto(t *testing.T) {
// 	photo := BasicPhoto{
// 		KeyPhoto: KeyPhoto{
// 			AlbumID: photos[0].AlbumID,
// 			Date:    photos[0].Date,
// 		},
// 		Tags:        photos[0].Tags,
// 		Description: "New Description",
// 	}
// 	photo, err := dao.Update(photo)
// 	test.Ok(t, err)
// 	expected := input.photo(photo.ID, photos[0].URL)
// 	test.Equals(t, expected, photo)
// }
