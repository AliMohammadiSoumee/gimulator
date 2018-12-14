package main

import (
	"os"
	"github.com/alidadar7676/gimulator/simulator"
	"fmt"
	"os/signal"
	"github.com/alidadar7676/gimulator/game"
	"log"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gui <IP> [PlayerName]")
		os.Exit(1)
	}
	ip := os.Args[1]

	controllerName := fmt.Sprintf("judge-controller")
	controller := game.NewController(controllerName, "default", &simulator.Client{Addr: ip})
	controller.Run()

	log.Println("Start game...")

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	}
}
