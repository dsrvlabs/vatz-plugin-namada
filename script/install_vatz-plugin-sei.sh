#!/bin/bash
set -e
set -v

#Create vatz log folder
mkdir /var/log/vatz

# Clone VATZ repository
cd /root
git clone git@github.com:dsrvlabs/vatz.git

# Compile VATZ
cd /root/vatz
make

## You will see binary named vatz

# Initialize VATZ
./vatz init

# Install vatz-plugin-namada
./vatz plugin install github.com/dsrvlabs/vatz-plugin-namada/plugins/node_governance_alarm node_governance_alarm

# Check plugin list
./vatz plugin list
