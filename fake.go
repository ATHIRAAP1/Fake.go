package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v6"
)

type Config struct {
	Count int `json:"count"`
}

type person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func GenerateFakeData(count int) []person {
	var data []person
	for i := 0; i < count; i++ {
		data = append(data, person{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			Age:   gofakeit.Number(18, 60),
		})
	}
	return data
}

func LoadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func main() {
	configFile := "config.json"
	config, err := LoadConfig(configFile)
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	fakeData := GenerateFakeData(config.Count)
	jsonData, err := json.MarshalIndent(fakeData, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
