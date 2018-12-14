#!/bin/sh

echo "Starting simulator..."
./simulator_amd64 127.0.0.1:7575 2> /dev/null &
SIM_PID=$!
sleep 0.3

echo "Starting Judge..."
./game_amd64 127.0.0.1:7575 2> /dev/null & 
JUDGE_PID=$!
sleep 0.3

echo "Starting random agent..."
./random_player_amd64 127.0.0.1:7575 Random01 2> /dev/null & 
RND_PID=$!

echo "Start GUI..."
./gui_amd64 127.0.0.1:7575 You 2> /dev/null

echo 'DONE...'
echo 'Closing all:'
echo "Killing random player: $(kill $RND_PID)"
echo "Killing judge: $(kill $JUDGE_PID)"
echo "Killing simulator $(kill $SIM_PID)"
