# Redis Failover Experiments

Scripts and random notes for redis sentinel failover experiments.

## Files

- `values-sentinel.yaml` Helm values for a simple redis cluster with sentinels enabled.
- `monitor.sh` Reports status changes for all nodes in a cluster.
- `fill_cluster.sh` Fills a redis cluster with random keys. Takes the id of the current master node as the first argument.
    - `random_keys.lua` Used for `fill_cluster.sh`.
- `deny-redis-traffic-to-node-*-networkpolicy.yaml` Deny network traffic to a single node.

## Resources

>> Is it OK to wait until 'master_link_status' becomes 'up', and 'master_sync_in_progress' becomes '0' and 'master_last_io_seconds' becomes >= 0?
> If you have no reason to believe something has gone haywire, this ought to tell you that the initial sync process has completed, yes.
- https://groups.google.com/g/redis-db/c/JPvnyfUWx_Q?pli=1

- https://lzone.de/cheat-sheet/Redis%20Sentinel
