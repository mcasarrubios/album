package dao

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

// DBProvider provider to DB
type DBProvider interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	Query(query *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

// New Creates a DAO
func New(db DBProvider) DataAccessor {
	return &DAO{db}
}

// OpenDB opens a session in the DB
func OpenDB() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(getConfig().session)
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

// Create a photo
func (dao *DAO) Create(input CreateInput, URL string) (*Photo, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	ph := input.photo(id.String(), URL)
	putItemInput, err := ph.dbPutItemInput()
	fmt.Println(putItemInput)
	if err != nil {
		return nil, err
	}

	_, err = dao.db.PutItem(putItemInput)
	if err != nil {
		return nil, err
	}
	return ph, nil
}

// List query photos
func (dao *DAO) List(query QueryInput) ([]Photo, error) {
	queryInput, err := query.dbQueryInput()
	if err != nil {
		return nil, err
	}

	resp, err := dao.db.Query(queryInput)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	photoObj := []Photo{}
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &photoObj)
	log.Println(photoObj)
	return photoObj, err
}

// func listItems() ([]photo, error) {
// 	var queryInput = &dynamodb.QueryInput{
// 		TableName: aws.String("Photo"),
// 		IndexName: aws.String("AlbumID"),
// 		KeyConditions: map[string]*dynamodb.Condition{
// 			"modifier": {
// 				ComparisonOperator: aws.String("EQ"),
// 				AttributeValueList: []*dynamodb.AttributeValue{
// 					{
// 						S: aws.String("David"),
// 					},
// 				},
// 			},
// 		},
// 	}

// 	expressionAttrs := make(map[string]*dynamodb.AttributeValue)
// 	expressionAttrs[":albumID"] = &dynamodb.AttributeValue{S: aws.String("albumID")}

// 	input := &dynamodb.QueryInput{
// 		TableName: aws.String("Photo"),
// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
// 			":albumID": {
// 				S: aws.String(albumID),
// 			},
// 			":albumID": {
// 				S: aws.String(albumID),
// 			},
// 			":albumID": {
// 				S: aws.String(albumID),
// 			},
// 		},
// 		KeyConditionExpression: aws.String("Artist = :v1"),
// 		ProjectionExpression:   aws.String("SongTitle"),
// 	}
// }

// func makeQuery(q queryInput) {
// 	expressionAttrs := make(map[string]*dynamodb.AttributeValue)

// 	if q.AlbumID != "" {

// 	}
// }

// func getItem(id string) (*book, error) {

//     input := &dynamodb.GetItemInput{
//         TableName: aws.String("Photo"),
//         Key: map[string]*dynamodb.AttributeValue{
//             "ID": {
//                 S: aws.String(id),
//             },
//         },
//     }

//     result, err := db.GetItem(input)
//     if err != nil {
//         return nil, err
//     }
//     if result.Item == nil {
//         return nil, nil
//     }

//     bk := new(book)
//     err = dynamodbattribute.UnmarshalMap(result.Item, bk)
//     if err != nil {
//         return nil, err
//     }

//     return bk, nil
// }

// func createItem(ph *photo) error {
// 	id, err := uuid.NewV4()
// 	if err != nil {
// 		return err
// 	}
// 	ph.ID = id.String()
// 	return putItem(ph)
// }
