#! /usr/bin/env bash

service rabbitmq-server start

adminpass=password
schedpass=password
agentpass=password
workrpass=password

echo "creating rabbitmq users"
rabbitmqctl add_user admin $adminpass
rabbitmqctl set_user_tags admin administrator
rabbitmqctl add_user scheduler $schedpass
rabbitmqctl add_user agent $agentpass
rabbitmqctl add_user worker $workrpass

echo "deleting guest user"
rabbitmqctl delete_user guest

echo "creating 'mig' vhost"
rabbitmqctl add_vhost mig

echo "creating ACLs for scheduler user"
rabbitmqctl set_permissions -p mig scheduler \
        '^(toagents|toschedulers|toworkers|mig\.agt\..*)$' \
        '^(toagents|toworkers|mig\.agt\.(heartbeats|results))$' \
	'^(toagents|toschedulers|toworkers|mig\.agt\.(heartbeats|results))$'

echo "creating ACLs for agent user"
rabbitmqctl set_permissions -p mig agent \
        '^mig\.agt\..*$' \
        '^(toschedulers|mig\.agt\..*)$' \
        '^(toagents|mig\.agt\..*)$'

echo "creating ACLs for worker user"
rabbitmqctl set_permissions -p mig worker \
	'^migevent\..*$' \
	'^migevent(|\..*)$' \
	'^(toworkers|migevent\..*)$'


service rabbitmq-server stop

       
        
                    
                   
                     
                        
                       
                      
                               
  
                                 
                                  

echo "set mirroring policy"
rabbitmqctl -p mig set_policy mig-mirror-all "^(toschedulers|toagents|toworkers|mig(|event))\." '{"ha-mode":"all"}'

chown rabbitmq /etc/rabbitmq/*
echo
echo "rabbitmq configured with the following users:"
echo "  admin       $adminpass"
echo "  scheduler   $schedpass"
echo "  agent       $agentpass"
echo "  worker      $workrpass"
echo
echo "copy ca.crt and rabbitmq.{crt,key} into /etc/rabbitmq/"
echo "then run $ service rabbitmq-server restart"
