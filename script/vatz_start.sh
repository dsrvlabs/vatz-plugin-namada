#!/bin/bash
set -e
set -v

VATZ_PATH="/root/vatz"
VATZ_PLUGIN_PATH="/root/.vatz"
CHAIN_ID="your chain id"
BASE_DIR="$HOME/.namada"
NODE_INFO="rpc node info"
NODE_GOVERNANCE_ALARM_PORT="10001"
PFD_STATUS_PORT="10006"

cd $VATZ_PATH
./vatz start --config default.yaml >> /var/log/vatz/vatz.log 2>&1 &
cd $VATZ_PLUGIN_PATH
./node_governance_alarm -port $NODE_GOVERNANCE_ALARM_PORT -apiPort $API_PORT -voterAddr $VOTER_ADDRESS >> /var/log/vatz/node_governance_alarm.log 2>&1 &

echo "true"
