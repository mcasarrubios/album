package schemas

import (
	"path/filepath"
	"strings"

	"github.com/mcasarrubios/album/common"
)

// Placeholders for replacing schema
type Placeholders struct {
	RootFields []string
	ItemFields []string
	MinItems   string
	MaxItems   string
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
	byteValue, err := common.ReadFile(absPath)
	if err != nil {
		return "", err
	}
	return string(byteValue), nil
}

func replace(schema string, opts Placeholders) string {
	schema = strings.Replace(schema, "\"{{rootFields}}\"", stringify(opts.RootFields), 1)
	schema = strings.Replace(schema, "\"{{itemFields}}\"", stringify(opts.ItemFields), 1)
	schema = strings.Replace(schema, "\"{{minItems}}\"", opts.MinItems, 1)
	schema = strings.Replace(schema, "\"{{maxItems}}\"", opts.MaxItems, 1)
	return schema
}

func stringify(fields []string) string {
	var output string
	if len(fields) > 0 {
		output = "\"" + strings.Join(fields, "\",\"") + "\""
	} else {
		output = ""
	}
	return output
}
