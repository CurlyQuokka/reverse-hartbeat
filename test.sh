#!/bin/bash
# Filename: index.sh
PUSH_URL="https://demo.kuma.pet/api/push/IYFpBFrGb1?status=up&msg=OK&ping="
INTERVAL=10

while true; do
    curl -s -o /dev/null $PUSH_URL
    echo "Pushed!"
    sleep $INTERVAL
done