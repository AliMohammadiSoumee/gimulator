.PHONY: simulator game gui random-player logger build-all run-all agent

ROOT = github.com/alidadar7676/gimulator

vendor:
	dep ensure

simulator: vendor
	go build -o bin/simulator_amd64 $(ROOT)/simulator/cmd/

game: vendor
	go build -o bin/game_amd64 $(ROOT)/game/cmd

gui: vendor
	go build -o bin/gui_amd64 $(ROOT)/gui

random-player: vendor
	go build -o bin/random_player_amd64 $(ROOT)/random_player

agent: vendor
	go build -o bin/agent_amd64 $(ROOT)/agent/cmd

logger: vendor
	go build -o bin/logger_amd64 $(ROOT)/logger

build-all: simulator game gui agent logger

run-all: build-all
	cd bin && sh run.sh
      

