#!/usr/bin/env bash
cd $(dirname $0)

. ./_params.sh

NETWORK=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.NetworkID}}{{end}}' ${NAME}-1)


NETWORK=${NETWORK} sentry/stop.sh

docker ps -a -q --filter "network=${NETWORK}" | while read id
do
    docker stop $id 2> /dev/null # fine if stopped already 
    docker rm $id 2> /dev/null # fine if removed already 
    echo "stopped/removed $id"
done

blockade destroy 2> /dev/null # fine if not found

docker network rm ${NETWORK} 2> /dev/null # fine if no network
