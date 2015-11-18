#! /bin/bash

docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)

cd postgres
docker build -t mig/postgres .
docker run -d --name=postgres --hostname=postgres mig/postgres

cd ../rabbitmq
docker build -t mig/rabbitmq .
docker run -d --name=rabbitmq --hostname=mig mig/rabbitmq

cd ../MIG
docker build -t mig/core .
sleep 30
docker run --name=core --hostname=api --link=postgres --link=rabbitmq mig/core

