#!/bin/sh

echo "Starting simulator..."
./simulator_amd64 127.0.0.1:7575 > /dev/null &
sleep 0.3

echo "Starting Judge..."
./game_amd64 127.0.0.1:7575 > /dev/null & 
sleep 0.3

echo "Starting random agent..."
./random_player_amd64 127.0.0.1:7575 Random01 > /dev/null & 

echo "Start GUI..."
./gui_amd64 127.0.0.1:7575 You > /dev/null &

echo 'DONE...'

