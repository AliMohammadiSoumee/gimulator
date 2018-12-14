package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alidadar7676/gimulator/simulator"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: random_player <ip> <playername>")
		os.Exit(1)
	}
	ip := os.Args[1]
	playername := os.Args[2]

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	controller := NewController(
		playername,
		"default",
		&simulator.Client{Addr: ip},
	)
	controller.Run()

	select {
	case <-sigs:
	}
}
