#!/bin/bash
set -e
set -v

VATZ_PATH="/root/vatz"
VATZ_PLUGIN_PATH="/root/.vatz"
CHAIN_ID="namada chain id"
BASE_DIR="$HOME/.namada"
NODE_INFO="http://localhost:26657"
NODE_GOVERNANCE_ALARM_PORT="10001"

cd $VATZ_PATH
./vatz start --config default.yaml >> /var/log/vatz/vatz.log 2>&1 &
cd $VATZ_PLUGIN_PATH
./node_governance_alarm -port $NODE_GOVERNANCE_ALARM_PORT -chainId $CHAIN_ID -baseDir $BASE_DIR -nodeInfo $NODE_INFO >> /var/log/vatz/node_governance_alarm.log 2>&1 &

echo "true"
