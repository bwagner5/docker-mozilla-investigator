#! /bin/bash

umount /etc/hosts
sed '/.*mig/ s/$/ mig.example.com/' /etc/hosts > hosts.bk && mv hosts.bk /etc/hosts

echo "printing hosts file"
cat /etc/hosts

src/mig.ninja/mig/bin/linux/amd64/mig-scheduler

