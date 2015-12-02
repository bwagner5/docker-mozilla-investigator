#! /bin/bash

cd agent

docker build -t mig/agent .
docker run -d mig/agent

