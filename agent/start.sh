#!/bin/bash

/bin/bash &

exec "$GOPATH/src/mig.ninja/mig/bin/linux/amd64/mig-agent-latest"

