package schemas

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Placeholders for replacing schema
type Placeholders struct {
	Fields   []string
	MinItems string
	MaxItems string
}

// Schema gets the JSON schema
func Schema(schemaName string, opts Placeholders) (string, error) {
	schema, err := read(schemaName)
	if err != nil {
		return "", err
	}
	return replace(schema, opts), nil
}

func read(fileName string) (string, error) {

	absPath, _ := filepath.Abs("./schemas/" + fileName + ".json")
	// Open our jsonFile
	jsonFile, err := os.Open(absPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return "", err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened json as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return string(byteValue), nil
}

func replace(schema string, opts Placeholders) string {
	var fields string
	if len(opts.Fields) > 0 {
		fields = "\"" + strings.Join(opts.Fields, "\",\"") + "\""
	} else {
		fields = ""
	}

	schema = strings.Replace(schema, "\"{{requiredFields}}\"", fields, 1)
	schema = strings.Replace(schema, "\"{{minItems}}\"", opts.MinItems, 1)
	schema = strings.Replace(schema, "\"{{maxItems}}\"", opts.MaxItems, 1)
	return schema
}
