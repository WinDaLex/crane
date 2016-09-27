#!/bin/sh
## name: crane.sh
set -e -x

: ${bridge=crane0}
: ${base=$HOME/.crane}

# Set up bridge network
if ! ip link show $bridge > /dev/null 2>&1
then
   sudo brctl addbr $bridge
   sudo ip addr add ${net:-"10.20.30.1/24"} dev $bridge
   sudo ip link set dev $bridge up
fi

# Bootup Docker engine
sudo docker daemon \
  --bridge=$bridge \
  --exec-root=$base/var/run/docker \
  --graph=$base/var/lib/docker \
  --host=unix://$base/var/run/docker.sock \
  --pidfile=$base/var/run/docker.pid \
  -D
