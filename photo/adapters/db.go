package adapters

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mcasarrubios/album/config"
	"github.com/mcasarrubios/album/photo/entities"
	uuid "github.com/satori/go.uuid"
)

// DBProvider provider to DB
type DBProvider interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	Query(query *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
}

// DB struct
type DB struct {
	db DBProvider
}

// New Creates a db adapter
func NewDB(db DBProvider) entities.DataAccessor {
	return &DB{db}
}

// OpenDB opens a session in the DB
func OpenDB() (*dynamodb.DynamoDB, error) {

	sess, err := session.NewSession(getConfig())
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

func getConfig() *aws.Config {
	awsConfig := config.GetConfig().AWS
	return &aws.Config{
		Endpoint: aws.String(awsConfig.Endpoint),
		Region:   aws.String(awsConfig.Region),
	}
}

// Create a photo
func (dao *DB) Create(input CreateInput) (*Photo, error) {
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

// // List query photos
// func (dao *DB) List(query QueryInput) (*QueryOutput, error) {
// 	queryInput, err := query.dbQueryInput()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return dao.query(queryInput)
// }

// // Get a photo
// func (dao *DB) Get(input GetInput) (*Photo, error) {
// 	queryInput, err := input.dbQueryInput()
// 	if err != nil {
// 		return nil, err
// 	}
// 	output, err := dao.query(queryInput)

// 	if len(output.Items) > 0 {
// 		return &output.Items[0], nil
// 	}
// 	return nil, nil
// }

// func (dao *DB) query(queryInput *dynamodb.QueryInput) (*QueryOutput, error) {
// 	output, err := dao.db.Query(queryInput)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	return queryOutput(output)
// }

// // Delete a photo
// func (dao *DB) Delete(input DeleteInput) error {
// 	deleteInput, err := input.dbDeleteInput()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = dao.db.DeleteItem(deleteInput)
// 	return err
// }
