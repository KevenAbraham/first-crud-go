package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("variables.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	fmt.Println(os.Getenv("TEST"))
}
