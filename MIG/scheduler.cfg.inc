[agent]
        ; timeout controls the inactivity period after which
        ; agents are marked offline
        timeout = "60m"

        ; heartbeatfreq maps to the agent configuration and helps
        ; the scheduler detect duplicate agents, and some other things
        heartbeatfreq = "5m"

        ; whitelist contains a list of agent queues that are allowed
        ; to send heartbeats and receive commands
        whitelist = "/var/cache/mig/agents_whitelist.txt"

        ; detect endpoints that are running multiple agents
        detectmultiagents = true

        ; issue kill orders to duplicate agents running on the same endpoint
        killdupagents = true

; the collector continuously pulls
; pending messages from the spool
[collector]
        ; frequency at which the collector runs,
        ; default is to run every second
        freq = "1s"

; the periodic runs less often that
; the collector and does cleanup and DB updates
[periodic]
        ; frequency at which the periodic jobs run
        freq = "87s"

        ; delete finished actions, commands and invalids after
        ; this period has passed
        deleteafter = "360h"

        ; run a rabbitmq unused queues cleanup job at this frequency
        ; this is DB & amqp intensive so don't run it too often
        queuescleanupfreq = "24h"

[directories]
        spool = "/var/cache/mig/"
        tmp = "/var/tmp/"

[postgres]
        host = "postgres"
        port = 5432
        dbname = "mig"
        user = "migscheduler"
        password = "password"
        sslmode = "disable"
        maxconn = 10

[mq]
        host  = "mig"
        port  = 5671
        user  = "scheduler"
        pass  = "password"
        vhost = "mig"

; TLS options
        usetls  = true
        cacert  = "/etc/mig/ca.crt"
        tlscert = "/etc/mig/scheduler.crt"
        tlskey  = "/etc/mig/scheduler.key"

; AMQP options
; timeout defaults to 10 minutes
; keep this higher than the agent heartbeat value
        timeout = "10m"

[logging]
        mode = "stdout" ; stdout | file | syslog
        level = "debug"

; for file logging
   file = "mig_scheduler.log"

; for syslog, logs go into local3
;    host = "localhost"
;    port = 514
;    protocol = "udp"
