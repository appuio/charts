# MariaDB Galera Failover Experiments

Scripts and random notes for galera failover experiments.

## Files

- `client.sh (ip)` Start a mysql client shell to the cluster.
- `monitor.sh` Reports status changes for all nodes in a cluster.
- `values.yaml` Helm values for a simple galera cluster.
- `fill.sql` Adds a table with 100'000'000 (one hundred million) rows.
- `deny-traffic-to-node.sh (remove) node_id` Deny network traffic to a single node.

## Resources

- https://mariadb.com/kb/en/what-is-mariadb-galera-cluster/
- https://github.com/bitnami/charts/pull/6769 Optimize mariadb galera cluster split-brain detection
