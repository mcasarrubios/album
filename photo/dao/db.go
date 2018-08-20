package dao

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
	_, err = dao.db.PutItem(ph.putItemInput())
	if err != nil {
		return nil, err
	}
	return ph, nil
}

// List query photos
func (dao *DAO) List(query QueryInput) ([]Photo, error) {
	// var resp1, err1 = dao.db.Query(query)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// } else {
	// 	personObj := []Person{}
	// 	err = dynamodbattribute.UnmarshalListOfMaps(resp1.Items, &personObj)
	// 	log.Println(personObj)
	// }
	return nil, nil
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

// MAPPING METHODS

func (ph *Photo) putItemInput() *dynamodb.PutItemInput {
	return &dynamodb.PutItemInput{
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
}

func (in CreateInput) photo(id string, URL string) *Photo {
	p := new(Photo)
	p.AlbumID = in.AlbumID
	p.Tags = in.Tags
	p.Description = in.Description
	p.Date = in.Date
	p.ID = id
	p.URL = URL
	return p
}

func (in QueryInput) query() *dynamodb.QueryInput {
	return &dynamodb.QueryInput{
		TableName: aws.String("Photo"),
		IndexName: aws.String("AlbumID"),
		KeyConditions: map[string]*dynamodb.Condition{
			"AlbumID": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(in.AlbumID),
					},
				},
			},
		},
	}
}
