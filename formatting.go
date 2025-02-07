package gohelpers

import (
	"encoding/json"
	"fmt"
)

func OutputStruct(structObj interface{}) {
	b, err := json.MarshalIndent(structObj, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
