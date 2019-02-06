package main

import (
	"encoding/json"
	"fmt"
	"github.com/alecthomas/jsonschema"
	dt "github.com/trustnetworks/analytics-common/datatypes"
)

func main() {

	schema := jsonschema.Reflect(&dt.Event{})

	json, err := json.MarshalIndent(schema, "", "  ")

	if err != nil {
		fmt.Println("Failed: ", err)
	} else {
		fmt.Print(string(json))
	}
}
