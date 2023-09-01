package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/melbahja/goph"
	"log"
	"os"
)

func main() {
	godotenv.Load(".env")
	auth, err := goph.Key(os.Getenv("PATH"), "")

	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New("root", os.Getenv("IP_ADDRESS"), auth)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	out, err := client.Run("ls -la")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
