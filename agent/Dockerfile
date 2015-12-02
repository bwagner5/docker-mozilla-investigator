FROM golang

RUN go get mig.ninja/mig

COPY mig-agent-conf.go src/mig.ninja/mig/conf/mig-agent-conf.go

RUN cd src/mig.ninja/mig && make mig-agent AGTCONF=conf/mig-agent-conf.go

COPY start.sh /start.sh
COPY legend.txt /tmp/legend.txt

#ENTRYPOINT ["/start.sh"]

CMD ["/bin/sh", "-c", "./src/mig.ninja/mig/bin/linux/amd64/mig-agent-latest -d"]

