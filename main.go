package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rmsrob/stcp/stcp"
)

func main() {
	listenAdrs := flag.String("listenAdrs", ":3000", "the port server address")
	flag.Parse()

	server := stcp.NewServer(*listenAdrs)
	go func() {
		for msg := range server.MsgChan {
			fmt.Printf("[from: %s] %s\n", string(msg.From), string(msg.Payload))
		}
	}()
	log.Fatal(server.Start())
}
