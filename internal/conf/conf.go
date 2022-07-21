package conf

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var URL string = "http://localhost:3000/"
var SaveDir string = "./replays"
var Port int = 8080

func Parse(args []string) {
	// Env variables
	URL = os.Getenv("BACK_URL")
	SaveDir = os.Getenv("BACK_SAVEDIR")
	port, err := strconv.Atoi(os.Getenv("BACK_PORT"))
	if err == nil {
		Port = port
	}

	// Arguments
	for i, arg := range args[1:] {
		switch arg {
		case "-u":
			URL = args[i+1]
			break

		case "-d":
			SaveDir = args[i+1]
			break

		case "-p":
			port, err := strconv.Atoi(args[i+1])
			if err != nil {
				log.Fatal("Port muss be an number. Given port:", args[i+1])
			}
			Port = port
			break

		case "--help":
			fmt.Println(
				"Arguments:\n" +
					"     -u     sets the osd-perm url\n" +
					"     -d     sets the replay save dir\n" +
					"     -p     sets the port\n" +
					"     --help Shows the help menu")
			os.Exit(0)
			break

		default:
			log.Fatalf("Argument %s does not exists. --help to the all commands", arg)
			return
		}
	}
}
