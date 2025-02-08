package main

import (
	"log"

	"github.com/eyvazceferov/yepyeniverson/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
