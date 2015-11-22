# Docker-Mozilla-InvestiGator (DMIG)

## By Brandon Wagner


### _THIS SHOULD NOT BE USED IN PRODUCTION!_


### THIS IS A SIMPLE WAY TO EVALUATE MOZILLA INVESTIGATOR WITHOUT HAVING TO GO THROUGH THE RATHER LONG CONFIGURATION PROCESS.

 *You can use this repository as a base to build your own docker Mozilla InvestiGator production containers. Most notably, TLS should be used on all connections to RabbitMQ and PostGres. TLS should also be terminated at Nginx for the MIG API. 
 
 ------------------
 
#### About:
 
 This repository contains multiple docker containers that load preconfigured software to run Mozilla's InvestiGator (MIG) distributed forensic and incident response system. The containers and a convenience script to bootstrap the full system on a single host helps to quickly evaluate the software or train on it. Without these containers, MIG is extremely time-consuming to setup and evaluate. Probably the easiest method of setup is using a cloud service which can be expensive for evaluation or training purposes. 
 
 The repository contains 6 containers:
   - MIG Core (includes MIG API and MIG Scheduler)
   - Nginix Reverse Proxy (Reverse proxies the MIG API)
   - PostGres RDMS 
   - RabbitMQ
   - MIG Client (to interact with MIG API)
   - MIG Agent (the base agent setup to be deployed on all hosts where you want to have remote forensics abilities)

##### start_mig.sh :
The script `start_mig.sh` will for the most part bootstrap the full enviroment. It will start MIG Core (API and Scheduler), PostGres, RabbitMQ, and Nginix. MIG Client is meant to be a convenience utility container that doesn't need to be operating all the time, so it is not started on the initial bootstrap (NOTE: You can use the MIG Client container on a production system with some reconfiguration for security. This gives you the convenience of not having to install the client on your laptop.). 

##### Mig Agent:
A MIG Agent is not bootstrapped right now because of an issue with starting the process with Docker (still investigating). The MIG Agent container is meant to be used as a base container (i.e. using docker's `FROM mig/agent`) to build a more complex image. I've also thought about just using the MIG Agent container to build the binaries for different platforms and then include those in other more complex images that extend other non-MIG dockerfiles.


#### Evaluation Process:

- Execute `start_mig.sh`
- `cd client` and run `docker run -i mig/client`
  - Enter password `mig`
  - Type `create investigator`
  - Give the investigator any name (e.g. Mig Investigator)
  - Give the investigator a public key (or use the one provided in the container at path `/root/mig_investigator_pubkey.asc`) 
  - Confirm with `y`
- `cd ../agent` and run `docker run -it mig/agent /bin/bash`
  - execute `src/mig.ninja/mig/bin/linux/amd64/mig-agent-latest`
- `cd ../client` and run `docker run -i mig/client`
  - Enter password `mig` and you should see one (1) agent active

** You can start more agents with the same commands






