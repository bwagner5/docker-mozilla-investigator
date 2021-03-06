// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Contributor: Julien Vehent jvehent@mozilla.com [:ulfr]
package main

import(
	"mig.ninja/mig"
	"time"
)

// some tags that are useful to differentiate agents. You can add whatever
// you want in this struct, it will be sent by the agent in each heartbeat
var TAGS = struct {
	Operator string `json:"operator"`
}{
	"CFRS780",
}

// restart the agent on failures, don't let it die
var ISIMMORTAL bool = true

// request installing of a service to start the agent at boot
var MUSTINSTALLSERVICE bool = true

// attempt to discover the public IP of the endpoint by querying the api
var DISCOVERPUBLICIP = true

// in check-in mode, the agent connects to the relay, runs all pending commands
// and exits. this mode is used to run the agent as a cron job, not a daemon.
var CHECKIN = false

var LOGGINGCONF = mig.Logging{
	Mode:	"stdout",	// stdout | file | syslog
	Level:	"debug",	// debug | info | ...
	//File:	"/tmp/migagt.log",
	//Host:	"syslog_hostname",
	//Port:	514,
	//Protocol: "udp",
}

// location of the rabbitmq server
// if a direct connection fails, the agent will look for the environment
// variables HTTP_PROXY and HTTPS_PROXY, and retry the connection using
// HTTP CONNECT proxy tunneling
var AMQPBROKER string = "amqp://agent:password@rabbitmq:5672/mig"

// location of the MIG API, used for discovering the public IP
var APIURL string = "http://mig/api/v1/"

// if the connection still fails after looking for a HTTP_PROXY, try to use the
// proxies listed below
var PROXIES = [...]string{`proxy.example.net:3128`, `proxy2.example.net:8080`}

// local socket used to retrieve stat information from a running agent
var SOCKET = ""//"127.0.0.1:51664"

// frequency at which the agent sends heartbeat messages
var HEARTBEATFREQ time.Duration = 300 * time.Second

// timeout after which a module run is killed
var MODULETIMEOUT time.Duration = 300 * time.Second

// Control modules permissions by PGP keys
var AGENTACL = [...]string{
`{
    "default": {
        "minimumweight": 2,
        "investigators": {
            "Mig Investigator": {
                "fingerprint": "44D53F962B29C278E82B7769C5026C27F594783B",
                "weight": 2
            }
        }
    }
}`,
`{
    "agentdestroy": {
        "minimumweight": 1,
        "investigators": {
            "MIG Scheduler": {
                "fingerprint": "",
                "weight": 1
            }
        }
    }
}`,
}


// PGP public keys that are authorized to sign actions
// this is an array of strings, put each public key block
// into its own array entry, as shown below
var PUBLICPGPKEYS = [...]string{
`-----BEGIN PGP PUBLIC KEY BLOCK-----

xsJuBFZPyVURCADwnZRF0+oLpCVN8+XiRYJtDSYIa3EQNWT4uU/gJ4844hbcvdiF
NNRz6KGfFMMEaSQpcb00k7JDTGn7QBsAEJ3+QEt41PBvRmJLpiFaAqgBiu0eLzSx
wmAu1LM2KYskCut5pI7Gv7MA4ZawQKwSxmDBZEgXQ5v9fsLGjAnQwCFB7M8LJS64
PubG+EDXWhrQUjW6ieRP3FvhuSta1VSy+HhOhH8ICgyOj6/wqPFvyXIaqRStvAWX
h08BVjmCM7Cmq/+3CuNMaPMmgQNWweXE9UaZFbGLCXD9TKjIQTHYt9SIxY9zN0mh
UwmmwYGUeIt01HVGm3+Td0jvSnEqHXjGilM3AQDs/97YmJtvuAmkcEZuiBqsloXh
D0cGT9cscfa5UPzQRQgAmJfRwS3bXjR/SEDn9xgwKbKihIiZc+FlmCpuVt+f9aaB
kTAxzXqY8Gz/1vNErQturf3r7BHkXivwYhNT7ub09+zcXHAolmO2SzaKTr5QNzej
wen+aY8cKGGn6bOOtJdeVV8jIxv8NFpemx8b5Z9mipQ/J7jVc/p7uJWWH8vvocAK
8CtWNSgwy3VYt7yJWefLPARGsyisl6uBIZ6nyKGIAXrwh1CYZQ6sQhHrLXNgWPpj
N7PHbqhJXLpjz+NAfSyzMHnTtXm36MQipu1sGAwuvBQiQukfTXCbUC7x/O4rNYL/
xZS66CDxXHIdFA3IKpaZuO1CuE7DV3Moa2dT3VKJowf/dlj7x6t7q09VpMYhYbMx
DVhaXyC1AIHU5JUpejPCXTZea/oGf+/tgsWAZLjQ86owVJiWEogFpQu3QHKF9cEz
2JmuIFvIxMfsR7eZ0JjXqtFFu2MBOM6L4TcxyGb3yX3UYfdGhsow1gyRTByxeIVp
8KMQ/cWFccLDsIzGo+C/EaeklfxgFXWwKHs1tooW9+Y85VgTyRTZMFLrg/TbQTRZ
YmT/RZFoP9x6cMh64ymahzdH1Aznl+eDTdx21M3pUk3zafNAcSr+TFGh+LxDebkO
gyS9q2SMjTf394nzx1OfGEi+OHb6EZiMkbxkvayjnMnND10RecRZbyQQEmU5ghQO
pc0vTW96aWxsYSBJbnZlc3RpR2F0b3IgKERvY2tlciBNSUcpIDxtaWdAbWlnLmNv
bT7CegQTEQgAIgUCVk/JVQIbAwYLCQgHAwIGFQgCCQoLBBYCAwECHgECF4AACgkQ
xQJsJ/WUeDtskAD+P1oUJKvPRWSZnQ9Nk+tWYQSQSsarxgk6jHfetcU945gA/ij5
e8Q1Uq3j6Q6fNqVQpsZBP+KuF8ko/0yiw41gAcXV
=/1QA
-----END PGP PUBLIC KEY BLOCK-----
`,
`-----BEGIN PGP PUBLIC KEY BLOCK-----

xsBNBFZfTFEBCACn4qC6agqjo1rVvidvpmOnqKEfXK6F/ypXEqgeL/XtP4jBXn0K
Gpv/82rb/8Zd6DHHJhPfwSvkMl8CNKeyhnd2PjZIphyAohYKLTE920ihmbEaLuDP
SjdJ6KiMevazORrLr1m3wY267kebUR+90Pcm6mwSfESB1CEXUuL9EhpXWZ/XO/x3
GBhkBH61LMg2AytRkevO8RZ+Ib5mRjAPINR72j3TqdtdNzaYTDBQqEqWOI8f1F59
Ev05CvQ7ps9vpNCVEJkMOxgkHsuho34O+ksV4BqsnhPvP569IAYQG3xTbuB68r8b
Tq1qakrXdX6o8sF+FNrhM8FJfSeLJv1CztcpABEBAAHNP21pZ3NjaGVkdWxlciAo
TUlHIFNjaGVkdWxlciBhY3Rpb24gc2lnbmluZyBrZXkpIDxzY2hlZHVsZXJAbWln
PsLAYgQTAQgAFgUCVl9MUQkQs04bO1ONascCGwMCGQEAAH3TCABVsl2J7Xzb7168
O0nMIHt2RmJFCBim8R1tCcyjwcM98IzdY0TeXB2wVkRwrmgPcvz2ntu+dyFitZql
mBAq3H3kwT4jhreCnAOfoyeutUPwXtvygqjujDXryFEgAsImVi07c+xrAFXULvuf
CV596ghh2I55tnsV++8vl+wKGGJRYWw5vmXCbd4MTXfkOX68JTQkDT8n4gGPWHCL
KwYvTokl3z/OQHQ/tjcm/pINgcJEi8aC8uf6ZYxLVFFdXHPe1bw6SSphhORw2LKr
4t2CtTK2rkQCebS98uSCdzp8xKbhmO0QzH+vzENJ1ps5CvG2qPLfRHHR1GF8KYRu
nT/yHle/zsBNBFZfTFEBCADr4j4TRo6A5bentK8ru3kgZ79Qxh6pfKFQ/zdE9jmH
HdC9nBp2MRlVCQhoF5uBhtae+qmjiEYdlhJejOSzuk7HdIzlRZ4KnorzN8fMwGDn
buWMnW7saTNtLm2MNnhJGgnRzk6U4n8I8a0c376OQdfHqjfsDdvuNav4yAhDk40g
LbY8tQEKKyOWE1oeQW0tedtBACSs5PFS7ajEeuBA+Q9wzqqs0g+CvZDxM9o1EmZa
RtCb4wwq/AO4okF1dHhWrBVXZslr0LbnBWAnHB0wv05Uf+74V/T6FZ8lVnqeoRED
sVjrqivoK4fPbLh8hx2gIPNIbeRjgJ4fodp9bpzTcCaxABEBAAHCwF8EGAEIABMF
AlZfTFEJELNOGztTjWrHAhsMAADkDwgAmp5214O9U7aKXGO/uJVzfZKyUf5upG3Z
CHlktgnavb6Ul/5BT6Px7M8U6dvpeoMNQ7qxxg5yuLHgbhwk6vEvp7sTkKe0b1sG
bXqSq/ChebLtBuYLu7a7xbYyswILUQplpe2Qu6c0/U/g8FL/0+mJ1JI5vW65q/BE
2HxSg97shNphWHpJCK8SJM7bwMfbnH11RPEYjlOlIRvtgqfaX+X+wp04/VeehlaP
nltWvEL3+th8MlXR3YWcppSnmK2v9cuaZznIjjgFeK5WXcyP9/I2yXx69bVI+bgw
yQFvRVU62xX66HCin0Xd+o6hfsA/uRYaaE6nxWwyo5eKqi1U6wyaog==
=sTPM
-----END PGP PUBLIC KEY BLOCK-----
`}


// CA cert that signs the rabbitmq server certificate, for verification
// of the chain of trust. If rabbitmq uses a self-signed cert, add this
// cert below
var CACERT = []byte(``)//`-----BEGIN CERTIFICATE-----
//MIIHyTCCBbGgAwIBAgIBATANBgkqhkiG9w0BAQUFADB9MQswCQYDVQQGEwJJTDEW
//........
//NOsF/5oirpt9P/FlUQqmMGqz9IgcgA38corog14=
//-----END CERTIFICATE-----`)

// All clients share a single X509 certificate, for TLS auth on the
// rabbitmq server. Add the public client cert below.
var AGENTCERT = []byte(``)//`-----BEGIN CERTIFICATE-----
//MIIGYjCCBUqgAwIBAgIDDD5PMA0GCSqGSIb3DQEBBQUAMIGMMQswCQYDVQQGEwJJ
//........
//04lr0kZCZTYpIQ5KFFe/s+3n0A3RDu4qzhrxOf3BMHyAITB+/Nh4IlRCZu2ygv2X
//ej2w/mPv
//-----END CERTIFICATE-----`)

// Add the private client key below.
var AGENTKEY = []byte(``)//`-----BEGIN RSA PRIVATE KEY-----
//MIIEpAIBAAKCAQEAvJQqCjE4I63S3kR9KV0EG9e/lX/bZxa/2QVvZGi9/Suj65nD
//........
//RMSEpg+wuIVnKUi6KThiMKyXfZaTX7BDuR/ezE/JHs1TN5Hkw43TCQ==
//-----END RSA PRIVATE KEY-----`)
