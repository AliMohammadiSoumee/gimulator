package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alidadar7676/gimulator/simulator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: logger <ip>")
		os.Exit(1)
	}
	ip := os.Args[1]

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	controller := NewController(
		os.Stdout,
		&simulator.Client{Addr: ip},
	)

	controller.Run()
	select {
	case <-controller.Wait:
	case <-sigs:
	}
}
