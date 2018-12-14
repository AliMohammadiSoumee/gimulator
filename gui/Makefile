IP ?= "127.0.0.1"
PLAYERNAME ?= "Player"

gen:
	go run gen.go

build: gen
	go build .

run-gui: build
	./gui $[IP] 

run-agent: build
	./gui $[IP] $[PLAYERNAME]
