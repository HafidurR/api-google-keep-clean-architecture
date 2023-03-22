package main

import (
	"api-google-keep/libs/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := config.Get()
	config.Init(cfg)
}