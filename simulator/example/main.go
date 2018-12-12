package main

import (
	"fmt"
	"time"

	"github.com/alidadar7676/gimulator/simulator"
)

func runServer() error {
	s := simulator.NewSimulator()
	s.Run()
	h := simulator.NewHTTPSimulator(s)
	return h.ListenAndServe("127.0.0.1:8585")
}

func main() {
	go func() {
		panic(runServer())
	}()

	// ensure server is running
	time.Sleep(time.Second * 1)

	var (
		c   = &simulator.Client{Addr: "127.0.0.1:8585"}
		ans interface{}
		err error
	)

	ch := make(chan simulator.Reconcile)
	err = c.Watch(simulator.Object{
		Key: simulator.Key{
			Namespace: "ns",
			Type:      "integer",
		},
		Value: map[string]interface{}{
			"hello": "world",
		}}, ch)
	fmt.Println("HUH?", err)

	go func() {
		for r := range ch {
			fmt.Println("Watched", r)
		}
	}()

	hellokey := simulator.Key{
		Namespace: "ns",
		Type:      "integer",
		Name:      "hello",
	}
	hello := simulator.Object{
		Key:   hellokey,
		Value: map[string]string{"hello": "world"},
	}
	err = c.Set(hello)
	fmt.Println(ans, err)

	ans, err = c.Get(simulator.Key{Name: "hello", Type: "integer", Namespace: "ns"})
	fmt.Println(ans, err)

	fmt.Scanln()
}
