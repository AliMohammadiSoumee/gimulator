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

	ans, err = c.Get("hello")
	fmt.Println("GET key HELLO:", ans, err)

	err = c.Set("hello", map[string]string{"hello": "world"})
	fmt.Println("SET key HELLO:", err)

	ans, err = c.Get("hello")
	fmt.Println("GET key HELLO:", ans, err)

	ch := make(chan simulator.Reconcile, 32)
	err = c.Watch("hello", ch)
	fmt.Println("WATCH key hello:", err)
	go func() {
		for r := range ch {
			fmt.Println("Watched", r)
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		err = c.Set("hello", map[string]int{"hello": i})
		fmt.Println("SET key hello:", err)
	}

	err = c.Delete("hello")
	fmt.Println("DELETE key hello:", err)

	ans, err = c.Get("hello")
	fmt.Println("GET key HELLO:", ans, err)

	fmt.Scanln()
}
