#!/bin/bash

set -e

host="$1"
shift
# shellcheck disable=SC2124
cmd="$@"

until mysql -h"$host" -P3306 -uroot -p"$MYSQL_ROOT_PASSWORD" -e 'SELECT 1'; do
  >&2 echo "MySQL is unavailable - sleeping"
  sleep 1
done

>&2 echo "MySQL is up - executing command"
exec $cmd
