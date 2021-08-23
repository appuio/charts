#!/bin/bash

echo Monitoring nodes

declare -A last_state=()

while :
do
    for i in {0..2}
    do
        node=redis-test-cluster-node-$i
        response=$(
            kubectl exec ${node} -c redis -- redis-cli -h localhost -p 6379 --no-auth-warning -a $REDIS_PASSWORD --eval /health/node_ready.lua 2>&1
        )
        if [ "$response" != "${last_state[$node]}" ]
        then
            echo "### $(date +%R:%S): Node $node state changed"
            echo "'${last_state[$node]}' -> '$response'"
            last_state[$node]=$response
        fi
    done
done
