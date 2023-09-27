#!/bin/bash

#====================================================================================================
### This script is used to check the status of the Galera cluster by HAProxy
### Maintainer: https://github.com/wejdross 
### It's sh/dash compatible, because it's used in a haproxy container and bash shell has serious issues to spawn 
### logic is simple, it takes 4 parameters and check if they are equal to the expected value
### parameters are returned always in the same manner thanks to ORDER BY, if any of the parameters is not equal to the expected value, script will exit with 1
#====================================================================================================
export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
pass=$(/bin/cat /secrets/mariadb-root-password)

if [ -z "$pass" ]; then
  echo "Error: password is empty, please check mount and permission declared in haproxy deployment"
  exit 1
fi

# sample return -> "WSREP_CLUSTER_STATUS Primary WSREP_CONNECTED ON WSREP_LOCAL_STATE 4 WSREP_READY ON"
return_val=$(/usr/bin/mysql -h $3 -u root -p"$pass" -e "SELECT VARIABLE_NAME, VARIABLE_VALUE FROM information_schema.global_status WHERE VARIABLE_NAME IN ('wsrep_cluster_status','wsrep_ready','wsrep_local_state','wsrep_connected') ORDER BY VARIABLE_NAME ASC;" -Ns)
if [ $? -ne 0 ]; then
  echo Error: Unable to connect to MySQL or query failed
  exit 1
else
  to_check=$(echo $return_val | cut -d' ' -f2)

  if [ $to_check != "Primary" ]; then
    echo "Error: WSREP_CLUSTER_STATUS is not 'Primary'"
    exit 1
  fi

  to_check=$(echo $return_val | cut -d' ' -f4)

  if [ $to_check != "ON" ]; then
    echo "Error: WSREP_CONNECTED is not 'ON'"
    exit 1
  fi

  to_check=$(echo $return_val | cut -d' ' -f6)

  if [ $to_check != 4 ]; then
    echo "Error: WSREP_LOCAL_STATE is not '4' it's: $to_check"
    exit 1
  fi

  to_check=$(echo $return_val | cut -d' ' -f8)
  if [ $to_check != "ON" ]; then
    echo "Error: WSREP_READY is not 'ON'"
    exit 1
  fi
fi
exit 0 
