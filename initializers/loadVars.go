package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("The vars were loaded successfully")
	}

}
