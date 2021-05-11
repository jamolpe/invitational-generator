package main

import (
	"fmt"
	"os"

	"github.com/jamolpe/invitational-generator/internal/api"
	"github.com/jamolpe/invitational-generator/internal/invitational"
	"github.com/jamolpe/invitational-generator/internal/repository"
	"github.com/jamolpe/invitational-generator/internal/store"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("NOTDEV") == "" {
		err := godotenv.Load("./.env")
		if err != nil {
			fmt.Println(err.Error())
			panic("error loading environment")
		}
	}
}

func main() {
	repo := repository.ConfigureAndConnect()
	store := store.New(repo)
	invitationalService := invitational.New(store)
	handler := api.New(invitationalService)
	handler.Router()
}
