package gohelpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LoadStructFromJsonFile reads a JSON file and unmarshals its content into a given struct type `T`.
// This function is generic, allowing it to load JSON data into any struct type.
//
// Parameters:
// - filename (string): The path to the JSON file.
//
// Returns:
// - (T, error): The populated struct of type `T`, and an error if any occurred during the file read or JSON decoding process.
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
//		person, err := LoadStructFromJsonFile[Person]("person.json")
//		if err != nil {
//			fmt.Println("Error loading JSON:", err)
//			return
//		}
//		fmt.Println("Loaded Struct:", person)
//	}
//
// JSON File (person.json):
//
//	{
//	  "name": "Alice",
//	  "age": 25,
//	  "email": "alice@example.com"
//	}
//
// Expected Output:
//
//	Loaded Struct: {Alice 25 alice@example.com}
func LoadStructFromJsonFile[T any](filename string) (T, error) {
	var result T

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return result, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	// Read the file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return result, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	// Unmarshal JSON into the struct
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("failed to parse JSON from file %s: %w", filename, err)
	}

	return result, nil
}
