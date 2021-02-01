#!/bin/bash

if [[ -z $RUNMODE ]]; then
    RUNMODE="prod"
fi

sed -i "s/dev/${RUNMODE}/g" /data/starcoin-explorer-api/conf/app.conf

cd /data/starcoin-explorer-api

while true
do
#	./starcoin-explorer-api_linux_amd64 $@ >> ./starcoin-explorer-api.log
  ./starcoin-explorer-api_linux_amd64
	echo "$(date) starcoin-explorer-api start"
	sleep 3
done
