package main

import (
	"log"
	"streaming.service/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
