#!/bin/bash
# Kill any process using the 8090 port

PORT=8080

while true; do
    echo "Checking for running process on port $PORT..."
    PID=$(lsof -ti:$PORT)

    if [ ! -z "$PID" ]; then
      echo "Killing process $PID on port $PORT..."
      kill -9 $PID
      sleep 1  # Wait a bit for the port to be freed
    else
      echo "No process running on port $PORT."
      break  # Exit the loop if no process is found
    fi
done
