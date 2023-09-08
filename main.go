package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/melbahja/goph"
	"log"
	
)

func main() {
	godotenv.Load(".env")
	auth, err := goph.Key("./aksc.pem", "")

	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New("root","3.210.248.187", auth)

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
