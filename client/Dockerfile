FROM golang

RUN apt-get update && apt-get install -y libreadline-dev

RUN go get mig.ninja/mig/client/mig-console

RUN cd /go/src/mig.ninja/mig/client

COPY migrc /root/.migrc

COPY gnupg /root/.gnupg/

RUN gpg --import /root/.gnupg/pubring.gpg

#RUN gpg --import /root/.gnupg/secring.gpg

#RUN gpg --import /root/.gnupg/trustdb.gpg
