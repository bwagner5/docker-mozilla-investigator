# MIG Core

Right now the MIG Core container will start with some manual interaction (this will soon be scripted out):
  
  - `docker run --link=rabbitmq --link=postgres -h api.mig.example.com -it mig/core /bin/bash`
  - Once it starts, you'll need an editor so execute `apt-get update` and `apt-get install -y vim`
  - Edit the hosts file to add the fully qualified domain name to the rabbitmq host.
  - `vi /etc/hosts` and add `mig.example.com` to the line that ends with `mig`

