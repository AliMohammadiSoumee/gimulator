package main

import(
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(2)
	print(rnd)
}

