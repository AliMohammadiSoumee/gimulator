package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alidadar7676/gimulator/simulator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: simulator <bind>")
		os.Exit(1)
	}
	bind := os.Args[1]

	simul := simulator.NewSimulator()
	simul.Run()

	http := simulator.NewHTTPSimulator(simul)
	log.Fatalln(
		http.ListenAndServe(bind))
}
