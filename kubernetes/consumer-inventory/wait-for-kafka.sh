#!/bin/bash
set -e

host="$1"
shift
port="$1"
shift

until nc -z "$host" "$port"; do
  echo "Warte auf Kafka bei $host:$port..."
  sleep 2
done

echo "Kafka ist bereit - führe Befehl aus"
exec "$@"