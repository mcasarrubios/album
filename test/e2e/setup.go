package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mcasarrubios/album/photo/dao"
)

func apiURL() string {
	return "http://localhost:3000"
}

func setup() {
	setEnv()
	err := startApp()
	if err != nil {
		fmt.Println(err)
		return
	}
	setupDB()
	time.Sleep(300 * time.Millisecond)
}

func setEnv() {
	os.Setenv("UP_STAGE", "test")
	os.Setenv("AWS_ACCESS_KEY_ID", "TEST-KEY")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "TEST-SECRET")
}

func startApp() error {
	cmd := exec.Command("up", "start")
	cmd.Dir = "../.."
	return cmd.Start()
}

func setupDB() {
	db, _ := dao.OpenDB()
	deleteTables(db)
	createTables(db)
}

func createTables(db *dynamodb.DynamoDB) {
	_, err := createTable(db, "Album-test", "albumId", "created")
	_, err = createTable(db, "Photo-test", "albumId", "date")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func deleteTables(db *dynamodb.DynamoDB) {
	_, err := deleteTable(db, "Album-test")
	_, err = deleteTable(db, "Photo-test")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createTable(db *dynamodb.DynamoDB, name string, hash string, sort string) (*dynamodb.CreateTableOutput, error) {
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

	return db.CreateTable(input)

}

func deleteTable(db *dynamodb.DynamoDB, name string) (*dynamodb.DeleteTableOutput, error) {
	return db.DeleteTable(&dynamodb.DeleteTableInput{TableName: &name})
}
