package common

import "encoding/json"

// MapStruct converts struct in a map
func MapStruct(in interface{}) (map[string]interface{}, error) {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(in)
	err := json.Unmarshal(inrec, &inInterface)
	return inInterface, err
}
