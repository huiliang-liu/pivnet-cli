#!/bin/bash

if [ -z $1 ]; then
    echo "usage: clear_release.sh [version]"
    exit 1
fi

#delete all product files
./pivnet-cli product-files -p greenplum-streaming-server -r $1 --format yaml | grep id | awk '{ print $3; }' | xargs -n 1 ./pivnet-cli delete-product-file -p greenplum-streaming-server -i

#delete all file groups
./pivnet-cli file-groups -p greenplum-streaming-server -r $1 --format yaml | grep '\- id' | awk '{ print $3; }' | xargs -n 1 ./pivnet-cli delete-file-group -p greenplum-streaming-server -i

#delete all releases
./pivnet-cli delete-release -p greenplum-streaming-server -r $1
