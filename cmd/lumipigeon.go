package main

import (
	"log"
	"lumipigeon/link"
)

func main() {
	link.NewServer(link.PORT, link.URL)

	log.Println("[main] waiting select{}")
	select {}
}
