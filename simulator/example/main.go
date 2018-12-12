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

	hello.Key.Name = "hello2"
	hello.Value = map[string]string{"hello": "world", "best": "size"}
	err = c.Set(hello)
	fmt.Println(ans, err)

	ans, err = c.Find(simulator.Object{
		Key:   simulator.Key{},
		Value: map[string]string{"hello": "world"},
	})
	fmt.Println(ans, err)

	fmt.Scanln()
}
