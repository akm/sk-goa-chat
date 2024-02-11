#!/bin/bash

max_iteration=10
interval=5

for i in $(seq 1 $max_iteration)
do
  mysql ${APP_MYSQL_USER_OPTS} ${APP_MYSQL_DB_NAME} -e 'SHOW TABLES;'
  result=$?
  if [[ $result -eq 0 ]]
  then
    echo "Database is ready"
    exit 0
  else
    sleep ${interval}
  fi
done

exit 1
