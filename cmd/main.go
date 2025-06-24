package main

import (
	"fmt"
	"log"

	"github.com/shahid-io/urban-waddle/cmd/api"
)

func main() {
	fmt.Println("urban waddle")

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
