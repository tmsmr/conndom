#!/usr/bin/env bash

# first arg expected to be user@target

REMOTE_SOCK=/var/run/dbus/system_bus_socket
LOCAL_SOCK=/tmp/conndom_dbus_system.sock

trap 'rm -f $LOCAL_SOCK' EXIT INT TERM

ssh -n -N -T -L $LOCAL_SOCK:$REMOTE_SOCK "$1"
