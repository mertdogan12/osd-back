package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mertdogan12/osd-back/internal/api"
	"github.com/mertdogan12/osd-back/internal/conf"
)

func main() {
	// .env
	godotenv.Load()

	conf.Parse(os.Args)

	http.HandleFunc("/replay/save", api.SaveReplay)

	fmt.Println("Server started on port:", conf.Port)
	http.ListenAndServe(":"+strconv.Itoa(conf.Port), nil)
}
