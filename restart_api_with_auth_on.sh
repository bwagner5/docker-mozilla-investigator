#! /bin/bash

docker kill core
docker rm core
docker kill mig
docker rm mig

cd MIG
mv api.cfg.inc api_no_auth.cfg.inc
mv api_auth_on.cfg.inc api.cfg.inc

docker build -t mig/core .
sleep 10
docker run -d --name=core --privileged --hostname=api --link=postgres --link=rabbitmq mig/core
docker run -d -p 80:80 --name=mig --link=core --hostname=mig mig/nginx

mv api.cfg.inc api_auth_on.cfg.inc
mv api_no_auth.cfg.inc api.cfg.inc

