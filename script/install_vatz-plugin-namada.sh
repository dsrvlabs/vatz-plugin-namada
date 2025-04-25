#!/bin/bash
set -v

#Create vatz log folder
mkdir -p /var/log/vatz

# Clone VATZ repository
cd $HOME
git clone https://github.com/dsrvlabs/vatz.git

# Compile VATZ
cd $HOME/vatz
make build

## You will see binary named vatz

# Initialize VATZ
./vatz init

# Install vatz-plugin-namada
./vatz plugin install github.com/dsrvlabs/vatz-plugin-namada/plugins/node_governance_alarm node_governance_alarm

# Check plugin list
./vatz plugin list
