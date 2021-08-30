#!/bin/bash
set -eu

action=apply

if [[ $# -eq 0 ]] ; then
    echo 'Usage: $0 (remove) NODE_ID'
    exit 3
fi

if [[ $# -eq 2 && "$1" == "remove" ]] ; then
    action=delete
    shift
fi

kubectl $action -f- <<- YAML
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-mariadb-traffic-node-$1
spec:
  podSelector:
    matchLabels:
      app.kubernetes.io/name: mariadb-galera
      statefulset.kubernetes.io/pod-name: mariadb-$1
  # Exception for test-client
  ingress:
    - from:
      - podSelector:
          matchLabels:
            role: client
  egress:
    - to:
      - podSelector:
          matchLabels:
            role: client

  policyTypes:
  - Ingress
  - Egress
YAML
