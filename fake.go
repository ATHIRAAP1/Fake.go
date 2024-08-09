package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bxcodec/faker/v3"
)

type Person struct {
	Name  string `faker:"name"`
	Age   int    `faker:"-"`
	Email string `faker:"email"`
}

func generateFakeData(jsonStructure string) (Person, error) {
	var person Person

	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStructure), &jsonData); err != nil {
		return person, err
	}

	if err := faker.FakeData(&person); err != nil {
		return person, err
	}

	return person, nil
}

func main() {
	jsonStructure := `{
        "name": "string",
        "age": "int",
        "email": "string"
    }`

	person, err := generateFakeData(jsonStructure)
	if err != nil {
		log.Fatalf("Error generating fake data: %v", err)
	}

	fmt.Printf("Generated Fake Data: %+v\n", person)
}
