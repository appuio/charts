#!/bin/bash

set -eu

node=redis-test-cluster-node-$1

kubectl cp scratchspace/random_keys.lua ${node}:/tmp/random_keys.lua -credis;

for i in {0..1000}
do
    kubectl exec ${node} -it -c redis -- 2>/dev/null redis-cli -h localhost -p 6379 -a $REDIS_PASSWORD --eval /tmp/random_keys.lua
    kubectl exec ${node} -it -c redis -- 2>/dev/null redis-cli -h localhost -p 6379 -a $REDIS_PASSWORD dbsize
done
