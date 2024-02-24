#!/bin/bash
set -e
set -v

pid0=`pidof node_governance_alarm`

# Kill process
kill -15 $pid0

# wait Kill process
sleep 1

while (( `lsof -p $pid0 | wc -l` > 0 ))
do
        echo `lsof -p $pid0 | wc -l`
        sleep 1
done

echo "seid $pid0 is killed."
