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

	http.HandleFunc("/users/me", api.SaveReplay)

	fmt.Println("Server started on port:", 3000)
	http.ListenAndServe(":3000", nil)
}
