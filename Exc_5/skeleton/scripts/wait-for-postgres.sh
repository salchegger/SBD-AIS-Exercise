#!/bin/sh
# wait-for-postgres.sh
set -e

host="$DB_HOST" #reads host from env variables
port="$POSTGRES_TCP_PORT" #reads port from env variables

echo "Waiting for Postgres at $host:$port..." #prints a status message

until nc -z "$host" "$port"; do #uses nc (netcat) to check if the port is open
  echo "Postgres is unavailable - sleeping" 
  sleep 1  #repeats every second until Postgres responds
done

echo "Postgres is up - executing ordersystem"
exec /app/ordersystem
