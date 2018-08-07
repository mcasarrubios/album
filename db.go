package main

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("eu-west-3"))

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

// Add a Photo record to DynamoDB.
func putItem(ph *photo) error {
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

	_, err := db.PutItem(input)
	return err
}
