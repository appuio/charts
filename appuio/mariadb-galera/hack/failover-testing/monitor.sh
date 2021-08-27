#!/bin/bash

echo Monitoring nodes

declare -A last_state=()

while :
do
    for i in {0..2}
    do
        node=mariadb-$i
        response=$(
            # kubectl exec ${node} -- sh -c 'mysqladmin status -u"${MARIADB_ROOT_USER}" -p"${MARIADB_ROOT_PASSWORD}"' 2>&1
            (kubectl exec ${node} -- sh -c 'mysql -Nb -u"${MARIADB_ROOT_USER}" -p"${MARIADB_ROOT_PASSWORD}" -e "select node_name from mysql.wsrep_cluster_members;"' | tr '\n' ',') 2>&1
        )
        if [ "$response" != "${last_state[$node]}" ]
        then
            echo "### $(date +%R:%S): Node $node state changed"
            echo "'${last_state[$node]}' -> '$response'"
            last_state[$node]=$response
        fi
    done
done
