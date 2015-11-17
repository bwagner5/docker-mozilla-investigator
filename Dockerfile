FROM golang

RUN go get mig.ninja/mig

RUN cd /go/src/mig.ninja/mig && make mig-scheduler && make mig-api && make worker-agent-intel && make worker-compliance-item

