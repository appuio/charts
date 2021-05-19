# MaxScale

This Chart deploys [MaxScale](https://mariadb.com/products/technology/maxscale)

## Configuration

The following table lists the configurable parameters of the BaaS Operator chart. For defaults consult `values.yaml`

| Parameter                       | Description                          |
| ---                             | ---                                  |
| `db1Address`                    | IP or Hostname of the 1. Galera node |
| `db2Address`                    | IP or Hostname of the 2. Galera node |
| `db3Address`                    | IP or Hostname of the 3. Galera node |
| `serviceUser`                   | MaxScale service user name           |
| `servicePassword`               | MaxScale service user password       |
| `monitorUser`                   | MaxScale monitor user name           |
| `monitorPassword`               | MaxScale monitor user password       |
| `service.masteronly.type`       | Maxscale master service type         |
| `service.rwsplit.type`          | Maxscale rwsplit service type        |
| `service.masteronly.annotations`| Maxscale master service annotations  |
| `service.rwsplit.annotations`   | Maxscale rwsplit service annotations |
