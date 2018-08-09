package dao

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DB interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemInput, error)
}

func New() (DB, error) {
	sess, err := session.NewSession(getConfig())
	if err != nil {
		return nil, err
	}
	return db.New(sess)
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

// Add a Photo record to DynamoDB.
func (db *DynamoDb) putItem(ph *photo) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Photo"),
		Item: map[string]*dynamodb.AttributeValue{
			"AlbumID": {
				S: aws.String(ph.AlbumID),
			},
			"ID": {
				S: aws.String(ph.ID),
			},
			"URL": {
				S: aws.String(ph.URL),
			},
			"Tags": {
				S: aws.String(strings.Join(ph.Tags, " ")),
			},
			"Description": {
				S: aws.String(ph.Description),
			},
			"Date": {
				S: aws.String(ph.Date),
			},
		},
	}

	output, err := db.PutItem(input)
	fmt.Printf("-----> %v", output)
	return err
}
