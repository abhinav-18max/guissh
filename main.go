package main

import (
	"fmt"
	// "github.com/joho/godotenv"
	"log"

	"github.com/melbahja/goph"
)

func main() {
	// godotenv.Load(".env")
	// auth, err := goph.Key("/home/rahul/.ssh/id_rsa", )

	auth := goph.Password("D#YPb4chIp")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	client, err := goph.New("root", "74.220.20.200", auth)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	for {
		dir, _ := client.Run("pwd")
		fmt.Printf("%s:%s$ ", client.User(), dir[:len(dir)-1])
		var comm string
		fmt.Scan((&comm))

		out, _ := client.Run(comm)
		fmt.Printf("%s", out)
		if comm == "exit" {
			break
		}
	}

	// fmt.Println(string(out))

	// var in int = 0
	// fmt.Println("Did you need to close the connection")

	// fmt.Scan(&in)
	// for in != 1 {
	// }

	if err != nil {
		log.Fatal(err)
	}

}
