package dao

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

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
func (dao *DAO) List(query QueryInput) (*QueryOutput, error) {
	queryInput, err := query.dbQueryInput()
	if err != nil {
		return nil, err
	}
	resp, err := dao.db.Query(queryInput)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return queryOutput(resp)
}

func queryOutput(dbOutput *dynamodb.QueryOutput) (*QueryOutput, error) {
	photos := []Photo{}
	err := dynamodbattribute.UnmarshalListOfMaps(dbOutput.Items, &photos)
	if err != nil {
		return nil, err
	}

	encoded, err := encodeLastKey(dbOutput.LastEvaluatedKey)
	if err != nil {
		return nil, err
	}
	return &QueryOutput{
		Items:   photos,
		LastKey: encoded,
	}, nil
}

func encodeLastKey(lastKey map[string]*dynamodb.AttributeValue) (string, error) {
	key := KeyPhoto{}
	err := dynamodbattribute.UnmarshalMap(lastKey, &key)
	if err != nil || key.AlbumID == "" {
		return "", err
	}
	js, err := json.Marshal(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(js), nil
}

func decodeStartKey(encoded string) (map[string]*dynamodb.AttributeValue, error) {
	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	lastKey := KeyPhoto{}
	err = json.Unmarshal(decoded, &lastKey)
	if err != nil {
		return nil, err
	}
	return dynamodbattribute.MarshalMap(lastKey)
}

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
