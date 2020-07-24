package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ayyaruq/zanarkand"
)

func main() {
	// Open flags for debugging if wanted (-assembly_debug_log)
	flag.Parse()

	// Setup the Sniffer
	//fmt.Println(devices.ListDeviceNames(false, false))
	//fmt.Println("\\Device\\NPF_{D767704F-EED4-4830-ABC6-EBC368D020C5}")
	sniffer, err := zanarkand.NewSniffer("pcap", "\\Device\\NPF_{D767704F-EED4-4830-ABC6-EBC368D020C5}")
	if err != nil {
		log.Fatal(err)
	}

	// Create a channel to receive Messages on
	subscriber := zanarkand.NewGameEventSubscriber()

	// Don't block the Sniffer, but capture errors
	go func() {
		err := subscriber.Subscribe(sniffer)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Capture the first 10 Messages sent from the server
	// This ignores Messages sent by the client to the server
	for i := 0; i < 10; i++ {
		imessage := <-subscriber.IngressEvents
		fmt.Println("in" + imessage.String())
		emessage := <-subscriber.EgressEvents
		fmt.Println("out" + emessage.String())
	}

	// Stop the sniffer
	subscriber.Close(sniffer)
}
