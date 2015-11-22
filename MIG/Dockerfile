FROM golang

RUN go get mig.ninja/mig

COPY scheduler.cfg.inc /etc/mig/scheduler.cfg
COPY api.cfg.inc /etc/mig/api.cfg

COPY ca.crt /etc/mig/
COPY scheduler.key /etc/mig/
COPY scheduler.crt /etc/mig/

COPY start.sh /

COPY agents_whitelist.txt /var/cache/mig/agents_whitelist.txt

RUN cd /go/src/mig.ninja/mig && make mig-scheduler && make mig-api

EXPOSE 51664

ENTRYPOINT ["/start.sh"]

#CMD ["src/mig.ninja/mig/bin/linux/amd64/mig-scheduler"]

