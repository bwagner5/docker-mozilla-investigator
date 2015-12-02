# Docker-Mozilla-InvestiGator (DMIG)

## By Brandon Wagner


### _THIS SHOULD NOT BE USED IN PRODUCTION WITHOUT MODIFICATION OF CREDENTIALS AND PKI!_


### THIS IS A SIMPLE WAY TO EVALUATE MOZILLA INVESTIGATOR.

 *You can use this repository as a base to build your own docker Mozilla InvestiGator production containers. Most notably, TLS should be used on all connections to RabbitMQ and PostGres. TLS should also be terminated at Nginx for the MIG API. 
 
 ------------------
 
#### About:
 
 This repository contains multiple docker containers that load preconfigured software to run Mozilla's InvestiGator (MIG) distributed forensic and incident response system. The containers and a convenience script to bootstrap the full system on a single host helps to quickly evaluate the software or train on it.  
 
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
The MIG Agent container is meant to be used as a base container to evaluate the agent in this bootstrapped environment. 

I plan on creating another container to build the MIG Agent binaries for different platforms and then include those in other more complex images that extend other non-MIG dockerfiles or just base images.


#### Getting Started:

- Execute `start_mig.sh`
- `cd client` and run `docker run -i mig/client`
  - Enter password `mig`
  - Type `create investigator`
  - Give the investigator any name (e.g. Mig Investigator)
  - Give the investigator a public key (or use the one provided in the container at path `/root/mig_investigator_pubkey.asc`) 
  - Confirm with `y`

** You can start more agents by cd'ing to the agent dir and executing `docker run -d mig/agent`






