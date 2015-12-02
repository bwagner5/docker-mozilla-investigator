FROM golang

RUN apt-get update && apt-get install -y libreadline-dev #gnupg-agent pinentry-curses

RUN go get mig.ninja/mig/client/mig-console
RUN go get mig.ninja/mig/client/mig

RUN cd /go/src/mig.ninja/mig/client

COPY migrc /root/.migrc

COPY gnupg /root/.gnupg/

COPY shellshock_action.maf /actions/shellshock_action.maf
COPY simple.maf /actions/simple.maf

RUN gpg --import /root/.gnupg/pubring.gpg

RUN gpg --export -a mig@mig.com > /root/mig_investigator_pubkey.asc 

#CMD ["bin/mig"]
CMD ["bin/mig-console"]
