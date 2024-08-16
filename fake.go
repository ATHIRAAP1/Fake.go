package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
)

type FieldType string

const (
	String FieldType = "string"
	Int    FieldType = "int"
)

type Schema struct {
	Fields map[string]FieldType `json:"fields"`
}

func loadSchema(filename string) (*Schema, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var schema Schema
	err = json.Unmarshal(data, &schema)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

func generateFakeData(schema *Schema, numRecords int) ([]map[string]interface{}, error) {
	var records []map[string]interface{}

	for i := 0; i < numRecords; i++ {
		record := make(map[string]interface{})
		for field, fieldType := range schema.Fields {
			switch fieldType {
			case String:
				record[field] = faker.Word()
			case Int:
				record[field] = rand.Intn(100)
			}
		}
		records = append(records, record)
	}

	return records, nil
}

func getFakeData(c *gin.Context) {
	schema, err := loadSchema("schema.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load schema"})
		return
	}

	numRecords := 20
	records, err := generateFakeData(schema, numRecords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate fake data"})
		return
	}

	c.JSON(http.StatusOK, records)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := gin.Default()

	router.GET("/fake-data", getFakeData)

	log.Println("Starting server on :8083")
	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
