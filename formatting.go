package gohelpers

import (
	"encoding/json"
	"fmt"
)

// OutputStruct marshals a given struct into a pretty-printed JSON format and prints it to the standard output.
// It takes any struct or map as input and formats it with indentation for better readability.
//
// If JSON marshaling fails, an error message is printed.
// If there is no json tags in the structure, it will print the property name
//
// Example usage:
//
//	type Person struct {
//		Name  string `json:"name"`
//		Age   int    `json:"age"`
//		Email string `json:"email"`
//	}
//
//	func main() {
//		p := Person{Name: "Alice", Age: 25, Email: "alice@example.com"}
//		OutputStruct(p)
//	}
//
// Output (formatted JSON):
//
//	{
//	  "name": "Alice",
//	  "age": 25,
//	  "email": "alice@example.com"
//	}
func OutputStruct(structObj interface{}) {
	b, err := json.MarshalIndent(structObj, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
