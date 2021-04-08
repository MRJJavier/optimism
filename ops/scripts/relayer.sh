#!/bin/bash
RETRIES=${RETRIES:-20}

# get the addrs from the URL provided
ADDRESSES=$(curl --retry-connrefused --retry $RETRIES --retry-delay 3 $URL)
# set the env
export ADDRESS_MANAGER_ADDRESS=$(echo $ADDRESSES | jq -r '.AddressManager')

# waits for l2geth to be up
curl --retry-connrefused --retry $RETRIES --retry-delay 1 $L2_NODE_WEB3_URL

# go
node ./exec/run-message-relayer.js