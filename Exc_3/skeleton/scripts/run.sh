  # todo
  # docker build
  # docker run db
  # docker run orderservice


#!/bin/sh
set -e #exit immediately on error

#building the GO app image
docker build -t orderservice .

#create network if not exists (added later on during 8.8)
docker network create SBDnetwork 2>/dev/null || true

#stopping and removing existing containers (added later on during 8.8)
docker rm -f ordersystem-db orderservice 2>/dev/null || true

#starting the PostgreSQL database
docker volume create pgdata_sbd
docker run -d \
  --name ordersystem-db \
  --network SBDnetwork \
  --env-file debug.env \
  -v pgdata_sbd:/var/lib/postgresql/18/docker \
  -p 5432:5432 \
  postgres:18


#waiting a few seconds for DB to initialize (added later on during 8.8)
echo "Waiting 5 seconds for Postgres to start..."
sleep 5


#starting the GO backend
docker run -d --name orderservice \
  --network SBDnetwork \
  --env-file debug.env \
  -p 3000:3000 \
  orderservice

#Final message (added later on during 8.8)
echo "All containers started! Visit http://localhost:3000"