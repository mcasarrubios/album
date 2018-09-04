package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mcasarrubios/album/common"
	"github.com/mcasarrubios/album/config"
	"github.com/mcasarrubios/album/photo/dao"
)

type dbSetup struct {
	srv  *dynamodb.DynamoDB
	conf *config.Config
}

func apiURL() string {
	return config.GetConfig().APIURL
}

func setup() {
	setEnv()
	setupDB(config.GetConfig())
}

func setEnv() {
	os.Setenv("UP_STAGE", "test")
	os.Setenv("AWS_ACCESS_KEY_ID", "TEST-KEY")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "TEST-SECRET")
}

func setupDB(conf *config.Config) {
	database, _ := dao.OpenDB()
	db := &dbSetup{
		srv:  database,
		conf: conf,
	}

	db.deleteTables()
	db.createTables()
	db.feedTables()
}

func (db *dbSetup) createTables() {
	_, err := db.createTable(db.conf.DB.AlbumTable, "albumId", "created")
	_, err = db.createTable(db.conf.DB.PhotoTable, "albumId", "date")
	if err != nil {
		panic(err)
	}
}

func (db *dbSetup) deleteTables() {
	_, err := db.deleteTable(db.conf.DB.AlbumTable)
	_, err = db.deleteTable(db.conf.DB.PhotoTable)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (db *dbSetup) createTable(name string, hash string, sort string) (*dynamodb.CreateTableOutput, error) {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(hash),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String(sort),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(hash),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String(sort),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String(name),
	}

	return db.srv.CreateTable(input)

}

func (db *dbSetup) deleteTable(name string) (*dynamodb.DeleteTableOutput, error) {
	return db.srv.DeleteTable(&dynamodb.DeleteTableInput{TableName: &name})
}

func (db *dbSetup) feedTables() {
	photos := []dao.Photo{}
	readFeed("photos", &photos)
	db.createPhotos(photos)
}

func readFeed(fileName string, v interface{}) {
	absPath, _ := filepath.Abs("./data/" + fileName + ".json")
	byteValue, err := common.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteValue, &v)
	if err != nil {
		panic(err)
	}
}

func (db *dbSetup) createPhotos(photos []dao.Photo) {
	for _, ph := range photos {
		db.createPhoto(ph)
	}
}

func (db *dbSetup) createPhoto(ph dao.Photo) {
	item, err := dynamodbattribute.MarshalMap(ph)
	if err != nil {
		panic(err)
	}

	_, err = db.srv.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(db.conf.DB.PhotoTable),
		Item:      item,
	})
	if err != nil {
		panic(err)
	}
}
