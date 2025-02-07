#!/bin/bash
redis-server --port 6379 --cluster-enabled yes --cluster-config-file /data/nodes-6379.conf --cluster-node-timeout 5000 --appendonly yes &
redis-server --port 6380 --cluster-enabled yes --cluster-config-file /data/nodes-6380.conf --cluster-node-timeout 5000 --appendonly yes &
redis-server --port 6381 --cluster-enabled yes --cluster-config-file /data/nodes-6381.conf --cluster-node-timeout 5000 --appendonly yes &

sleep 5

redis-cli --cluster create 127.0.0.1:6379 127.0.0.1:6380 127.0.0.1:6381 --cluster-replicas 0 --cluster-yes

wait
