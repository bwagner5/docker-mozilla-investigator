#! /bin/bash

docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)

cd postgres
docker build -t mig/postgres .
docker run -d -p 5432:5432 --name=postgres --hostname=postgres mig/postgres

cd ../rabbitmq
docker build -t mig/rabbitmq .
docker run -d -p 8080:15672 --name=rabbitmq --hostname=mig mig/rabbitmq

cd ../MIG
docker build -t mig/core .
sleep 30
docker run -d --name=core --privileged --hostname=api --link=postgres --link=rabbitmq mig/core

cd ../nginx
docker build -t mig/nginx .
docker run -d -p 80:80 --name=mig --link=core --hostname=mig mig/nginx

cd ../client
docker build -t mig/client .
#docker run -it --name=client --hostname=client mig/client

cd ../agent
docker build -t mig/agent .
docker run -d mig/agent
