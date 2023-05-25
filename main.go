package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAdrs := flag.String("listenAdrs", ":3000", "the port server address")
	flag.Parse()

	server := NewServer(*listenAdrs)
	go func() {
		for msg := range server.msgChan {
			fmt.Printf("[from: %s] %s\n", string(msg.from), string(msg.payload))
		}
	}()
	log.Fatal(server.Start())
}
