#!/bin/sh

exec podman run --rm --interactive         \
     --security-opt label=disable          \
     --volume "${PWD}":/pwd --workdir /pwd \
     quay.io/coreos/butane:release         \
     "${@}"
