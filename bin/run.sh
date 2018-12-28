#!/bin/sh

echo "Starting simulator..."
./simulator_amd64 0.0.0.0:7575 2> /dev/null &
SIM_PID=$!
sleep 0.3

echo "Starting Judge..."
./game_amd64 0.0.0.0:7575 2> /dev/null & 
JUDGE_PID=$!
sleep 0.3

echo "Starting random agent..."
./agent_amd64 0.0.0.0:7575 Agent01 & #1> log1.txt & 
RND_PID=$!

echo "Starting random agent..."
./agent_amd64 0.0.0.0:7575 Agent02 & #2> log2.txt & 
RNDD_PID=$!

echo "Start GUI..."
./gui_amd64 0.0.0.0:7575 2> /dev/null

echo 'DONE...'
echo 'Closing all:'
echo "Killing agent1 player: $(kill $RND_PID)"
echo "Killing agent2 player: $(kill $RNDD_PID)"
echo "Killing judge: $(kill $JUDGE_PID)"
echo "Killing simulator $(kill $SIM_PID)"
