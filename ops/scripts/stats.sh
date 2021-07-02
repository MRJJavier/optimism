#!/bin/bash

# set up the stats file
mkdir ~/logs
touch ~/logs/stats

while true; do
  {
    echo "$(date) ----------------";
    echo "location ---------------";
    pwd;
    echo "total memory usage -----";
    free -m;
    echo "docker stats -----------";
    docker stats --no-stream;
    echo "memory munchers --------";
    ps aux --sort=-%mem | head;
  } >> ~/logs/stats
  sleep 1;
done
