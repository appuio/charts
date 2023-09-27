#!/bin/bash
    
### This script is used to check the status of the Galera cluster
### Maintainer: https://github.com/wejdross 
### It takes hash of checks and expected values as input and returns 0 if all checks are successful, 1 otherwise 
pass=$(/bin/cat /secrets/mariadb-root-password)

if [ -z "$pass" ]; then
  echo "Error: password is empty, please check mount and permission declared in haproxy deployment"
  exit 1
fi

declare -A mysql_checks

mysql_checks=( 
  ["wsrep_cluster_status"]="Primary"
  ["wsrep_ready"]="ON"
  ["wsrep_local_state"]="4"
  ["wsrep_connected"]="ON"
)

for check in "${!mysql_checks[@]}"; do
  return_val=$(/usr/bin/mysql -h $3 -u root -p"$pass" -e "SELECT VARIABLE_NAME, VARIABLE_VALUE FROM information_schema.global_status WHERE VARIABLE_NAME = '$check';" -Ns)

  if [ $? -ne 0 ]; then
    echo Error: Unable to connect to MySQL or query failed
    exit 1
  else
    strarray=($return_val) # convert string to array
    if [ "${strarray[1]}" != "${mysql_checks[$check]}" ]; then
      echo Error: "$check" is not "${mysql_checks[$check]}", actual status is: "${strarray[1]}"
      exit 1
    fi
  fi
done

exit 0 