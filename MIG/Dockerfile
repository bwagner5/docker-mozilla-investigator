FROM golang

RUN go get mig.ninja/mig

COPY scheduler.cfg.inc /go/src/mig.ninja/mig/conf/
COPY api.cfg.inc /go/src/mig.ninja/mig/conf/

COPY ca.crt /etc/mig/
COPY scheduler.key /etc/mig/
COPY scheduler.crt /etc/mig/

RUN cd /go/src/mig.ninja/mig && make mig-scheduler && make mig-api && make worker-agent-intel && make worker-compliance-item

CMD ["/opt/mig/bin/mig-scheduler"]

