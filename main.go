package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mertdogan12/osd-back/internal/api"
)

func main() {
	// .env
	godotenv.Load()

	http.HandleFunc("/replay/save", api.SaveReplay)

	fmt.Println("Server started on port:", 8080)
	http.ListenAndServe(":8080", nil)
}
