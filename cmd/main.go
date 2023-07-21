package main

import (
	"github.com/Avyukth/streaming.service/internal/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
