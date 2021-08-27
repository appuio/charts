#!/bin/bash
set -eu

kubectl run mariadb-client \
    --rm --tty -i --restart='Never' \
    --namespace default \
    --labels="role=client" \
    --image docker.io/bitnami/mariadb-galera:10.5.12-debian-10-r1 \
    --command -- \
        mysql -h ${1:-mariadb} -P 3306 -uroot \
        -p$(kubectl get secret --namespace default mariadb -o jsonpath="{.data.mariadb-root-password}" | base64 --decode) \
        my_database \
||:

kubectl delete pod/mariadb-client ||:
