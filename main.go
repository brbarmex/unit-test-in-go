package main

import (
	"store/api"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	api.Execute()
}
