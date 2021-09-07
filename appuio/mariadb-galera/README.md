# mariadb-galera

![Version: 1.2.0](https://img.shields.io/badge/Version-1.2.0-informational?style=flat-square) ![AppVersion: 10.5.12](https://img.shields.io/badge/AppVersion-10.5.12-informational?style=flat-square)

MariaDB Galera is a multi-master database cluster solution for synchronous replication and high availability.

**Homepage:** <https://github.com/appuio/charts/tree/master/appuio/mariadb-galera>

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install mariadb-galera appuio/mariadb-galera
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | Affinity for pod assignment |
| clusterDomain | string | `"cluster.local"` | Kubernetes DNS Domain name to use |
| configurationConfigMap | string | `""` | ConfigMap with the MariaDB configuration files (Note: Overrides `mariadbConfiguration`). The value is evaluated as a template. |
| customPasswordFiles | object | `{}` | Use custom password files when `usePasswordFiles` is set to `true`. Define path for keys `root`, `user`, and `mariabackup`. |
| db.forcePassword | bool | `false` | Option to force users to specify a password. That is required for 'helm upgrade' to work properly. |
| db.name | string | `"my_database"` | Name for new database to create |
| db.password | string | `""` | Password for the new user. Ignored if existing secret is provided. |
| db.user | string | `""` | Username of new user to create |
| diagnosticMode.args | list | `["infinity"]` | Args to override all containers in the deployment |
| diagnosticMode.command | list | `["sleep"]` | Command to override all containers in the deployment |
| diagnosticMode.enabled | bool | `false` | Enable diagnostic mode (all probes will be disabled and the command will be overridden) |
| existingSecret | string | `""` | Use existing secret for password details (`rootUser.password`, `db.password`, `galera.mariabackup.password` will be ignored and picked up from this secret) The secret has to contain the keys mariadb-root-password, mariadb-galera-mariabackup-password and mariadb-password. |
| extraContainers | list | `[]` | Additional containers (this value is evaluated as a template) |
| extraDeploy | list | `[]` | Array of extra objects to deploy with the release (evaluated as a template) |
| extraEnvVars | list | `[]` | Array containing extra env vars to configure MariaDB Galera replicas |
| extraEnvVarsCM | string | `""` | ConfigMap containing extra env vars to configure MariaDB Galera replicas |
| extraEnvVarsSecret | string | `""` | Secret containing extra env vars to configure MariaDB Galera replicas |
| extraFlags | string | `""` | MariaDB additional command line flags |
| extraInitContainers | list | `[]` | Additional init containers (this value is evaluated as a template) |
| extraVolumeMounts | list | `[]` |  |
| extraVolumes | list | `[]` |  |
| forceUpdate | bool | `false` | Force update the StatefulSet. If enabled the chart will recreate the StatefulSet without touching the Pods (cascade orphan), allowing you to update locked field, such as PVC size. |
| forceUpdateKubectlImage | string | `"quay.io/bitnami/kubectl"` | Image used to recreate the StatefulSet. Needs to contain kubectl. |
| fullnameOverride | string | `""` | String to fully override common.names.fullname template with a string |
| galera.bootstrap.bootstrapFromNode | string | `""` | Node to bootstrap from, you will need to change this parameter in case you want to bootstrap from other node |
| galera.bootstrap.forceSafeToBootstrap | bool | `false` | Force `safe_to_bootstrap: 1` in `grastate.date` file. This will set safe_to_bootstrap=1 in the node indicated by bootstrapFromNode. |
| galera.mariabackup.forcePassword | bool | `false` | Option to force users to specify a password. That is required for 'helm upgrade' to work properly. |
| galera.mariabackup.password | string | `""` | MariaBackup password. Password is ignored if existingSecret is specified. |
| galera.mariabackup.user | string | `"mariabackup"` | MariaBackup username |
| galera.name | string | `"galera"` | Galera cluster name |
| global.imagePullSecrets | list | `[]` | Global Docker registry secret names as an array |
| global.imageRegistry | string | `""` | Global Docker image registry |
| global.storageClass | string | `""` | Global StorageClass for Persistent Volume(s) |
| hostAliases | list | `[]` | Add deployment host aliases |
| image.debug | bool | `false` | Specify if debug logs should be enabled |
| image.pullPolicy | string | `"IfNotPresent"` | MariaDB Galera image pull policy. Defaults to 'Always' if image tag is 'latest', else set to 'IfNotPresent' |
| image.pullSecrets | list | `[]` |  |
| image.registry | string | `"docker.io"` | MariaDB Galera image registry |
| image.repository | string | `"bitnami/mariadb-galera"` | MariaDB Galera image repository |
| image.tag | string | `"10.5.12-debian-10-r1"` | MariaDB Galera image tag (immutable tags are recommended) |
| initdbScripts | object | `{}` | Specify dictionary of scripts to be run at first boot |
| initdbScriptsConfigMap | string | `""` | ConfigMap with the initdb scripts (Note: Overrides `initdbScripts`) |
| ldap.base | string | `""` | LDAP base DN |
| ldap.binddn | string | `""` |  |
| ldap.bindpw | string | `""` | LDAP bind password |
| ldap.bslookup | string | `""` | LDAP base lookup |
| ldap.enabled | bool | `false` | Enable LDAP support |
| ldap.filter | string | `""` | LDAP custom filter |
| ldap.map | string | `""` | LDAP custom map |
| ldap.nss_initgroups_ignoreusers | string | `"root,nslcd"` | LDAP ignored users |
| ldap.scope | string | `""` | LDAP search scope |
| ldap.tls_reqcert | string | `""` | LDAP TLS check on server certificates |
| ldap.uri | string | `""` | LDAP URL beginning in the form `ldap |
| livenessProbe.enabled | bool | `true` | Turn on and off liveness probe |
| livenessProbe.failureThreshold | int | `3` | Minimum consecutive failures for the probe |
| livenessProbe.initialDelaySeconds | int | `120` |  |
| livenessProbe.periodSeconds | int | `10` | How often to perform the probe |
| livenessProbe.successThreshold | int | `1` | consecutive successes for the probe |
| livenessProbe.timeoutSeconds | int | `1` | When the probe times out |
| mariadbConfiguration | string | `"[client]\nport=3306\nsocket=/opt/bitnami/mariadb/tmp/mysql.sock\nplugin_dir=/opt/bitnami/mariadb/plugin\n\n[mysqld]\ndefault_storage_engine=InnoDB\nbasedir=/opt/bitnami/mariadb\ndatadir=/bitnami/mariadb/data\nplugin_dir=/opt/bitnami/mariadb/plugin\ntmpdir=/opt/bitnami/mariadb/tmp\nsocket=/opt/bitnami/mariadb/tmp/mysql.sock\npid_file=/opt/bitnami/mariadb/tmp/mysqld.pid\nbind_address=0.0.0.0\n\n## Character set\n##\ncollation_server=utf8_unicode_ci\ninit_connect='SET NAMES utf8'\ncharacter_set_server=utf8\n\n## MyISAM\n##\nkey_buffer_size=32M\nmyisam_recover_options=FORCE,BACKUP\n\n## Safety\n##\nskip_host_cache\nskip_name_resolve\nmax_allowed_packet=16M\nmax_connect_errors=1000000\nsql_mode=STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_AUTO_VALUE_ON_ZERO,NO_ENGINE_SUBSTITUTION,NO_ZERO_DATE,NO_ZERO_IN_DATE,ONLY_FULL_GROUP_BY\nsysdate_is_now=1\n\n## Binary Logging\n##\nlog_bin=mysql-bin\nexpire_logs_days=14\n# Disabling for performance per http://severalnines.com/blog/9-tips-going-production-galera-cluster-mysql\nsync_binlog=0\n# Required for Galera\nbinlog_format=row\n\n## Caches and Limits\n##\ntmp_table_size=32M\nmax_heap_table_size=32M\n# Re-enabling as now works with Maria 10.1.2\nquery_cache_type=1\nquery_cache_limit=4M\nquery_cache_size=256M\nmax_connections=500\nthread_cache_size=50\nopen_files_limit=65535\ntable_definition_cache=4096\ntable_open_cache=4096\n\n## InnoDB\n##\ninnodb=FORCE\ninnodb_strict_mode=1\n# Mandatory per https://github.com/codership/documentation/issues/25\ninnodb_autoinc_lock_mode=2\n# Per https://www.percona.com/blog/2006/08/04/innodb-double-write/\ninnodb_doublewrite=1\ninnodb_flush_method=O_DIRECT\ninnodb_log_files_in_group=2\ninnodb_log_file_size=128M\ninnodb_flush_log_at_trx_commit=1\ninnodb_file_per_table=1\n# 80% Memory is default reco.\n# Need to re-evaluate when DB size grows\ninnodb_buffer_pool_size=2G\ninnodb_file_format=Barracuda\n\n## Logging\n##\nlog_error=/opt/bitnami/mariadb/logs/mysqld.log\nslow_query_log_file=/opt/bitnami/mariadb/logs/mysqld.log\nlog_queries_not_using_indexes=1\nslow_query_log=1\n\n## SSL\n## Use extraVolumes and extraVolumeMounts to mount /certs filesystem\n# ssl_ca=/certs/ca.pem\n# ssl_cert=/certs/server-cert.pem\n# ssl_key=/certs/server-key.pem\n\n[galera]\nwsrep_on=ON\nwsrep_provider=/opt/bitnami/mariadb/lib/libgalera_smm.so\nwsrep_sst_method=mariabackup\nwsrep_slave_threads=4\nwsrep_cluster_address=gcomm://\nwsrep_cluster_name=galera\nwsrep_sst_auth=\"root:\"\n# Enabled for performance per https://mariadb.com/kb/en/innodb-system-variables/#innodb_flush_log_at_trx_commit\ninnodb_flush_log_at_trx_commit=2\n# MYISAM REPLICATION SUPPORT #\nwsrep_replicate_myisam=ON\n\n[mariadb]\nplugin_load_add=auth_pam\n\n## Data-at-Rest Encryption\n## Use extraVolumes and extraVolumeMounts to mount /encryption filesystem\n# plugin_load_add=file_key_management\n# file_key_management_filename=/encryption/keyfile.enc\n# file_key_management_filekey=FILE:/encryption/keyfile.key\n# file_key_management_encryption_algorithm=AES_CTR\n# encrypt_binlog=ON\n# encrypt_tmp_files=ON\n\n## InnoDB/XtraDB Encryption\n# innodb_encrypt_tables=ON\n# innodb_encrypt_temporary_tables=ON\n# innodb_encrypt_log=ON\n# innodb_encryption_threads=4\n# innodb_encryption_rotate_key_age=1\n\n## Aria Encryption\n# aria_encrypt_tables=ON\n# encrypt_tmp_disk_tables=ON"` | Configuration for the MariaDB server |
| metrics.enabled | bool | `false` | Start a side-car prometheus exporter |
| metrics.extraFlags | list | `[]` | MariaDB Prometheus exporter additional command line flags |
| metrics.image.pullPolicy | string | `"IfNotPresent"` | MariaDB Prometheus exporter image pull policy |
| metrics.image.pullSecrets | list | `[]` | MariaDB Prometheus exporter image pull secrets. Optionally specify an array of imagePullSecrets (secrets must be manually created in the namespace) |
| metrics.image.registry | string | `"docker.io"` | MariaDB Prometheus exporter image registry |
| metrics.image.repository | string | `"bitnami/mysqld-exporter"` | MariaDB Prometheus exporter image repository |
| metrics.image.tag | string | `"0.13.0-debian-10-r75"` | MariaDB Prometheus exporter image tag (immutable tags are recommended) |
| metrics.prometheusRules.enabled | bool | `false` | if `true`, creates a Prometheus Operator PrometheusRule (also requires `metrics.enabled` to be `true`, and makes little sense without ServiceMonitor) |
| metrics.prometheusRules.rules | object | `{}` | PrometheusRule rules to configure |
| metrics.prometheusRules.selector | object | `{"app":"prometheus-operator","release":"prometheus"}` | [object] Additional labels to add to the PrometheusRule so it is picked up by the operator If using the [Helm Chart](https://github.com/helm/charts/tree/master/stable/prometheus-operator) this is the name of the Helm release and 'app: prometheus-operator' |
| metrics.resources.limits | object | `{}` | The resources limits for the container |
| metrics.resources.requests | object | `{}` | The requested resources for the container |
| metrics.service.annotations | object | `{"prometheus.io/port":"9104","prometheus.io/scrape":"true"}` | Prometheus exporter service annotations |
| metrics.service.port | int | `9104` | Prometheus exporter service port |
| metrics.service.type | string | `"ClusterIP"` | Prometheus exporter service type |
| metrics.serviceMonitor.enabled | bool | `false` | if `true`, creates a Prometheus Operator ServiceMonitor (also requires `metrics.enabled` to be `true`) |
| metrics.serviceMonitor.interval | string | `""` | How frequently to scrape metrics (use by default, falling back to Prometheus' default) |
| metrics.serviceMonitor.metricRelabelings | list | `[]` | MetricRelabelConfigs to apply to samples before ingestion |
| metrics.serviceMonitor.namespace | string | `""` | Optional namespace which Prometheus is running in |
| metrics.serviceMonitor.relabelings | list | `[]` | RelabelConfigs to apply to samples before scraping |
| metrics.serviceMonitor.scrapeTimeout | string | `""` | Timeout after which the scrape is ended |
| metrics.serviceMonitor.selector | object | `{"prometheus":"kube-prometheus"}` | ServiceMonitor selector labels. Default to kube-prometheus install (CoreOS recommended), but should be set according to Prometheus install. |
| nameOverride | string | `""` |  |
| nodeAffinityPreset.key | string | `""` | Node label key to match. Ignored if `affinity` is set. |
| nodeAffinityPreset.type | string | `""` | Node affinity preset type. Ignored if `affinity` is set. Allowed values: `soft` or `hard` |
| nodeAffinityPreset.values | list | `[]` | Node label values to match. Ignored if `affinity` is set. |
| nodeSelector | object | `{}` | Node labels for pod assignment |
| persistence.accessModes | list | `["ReadWriteOnce"]` | Persistent Volume Access Modes |
| persistence.annotations | object | `{}` | Persistent Volume Claim annotations |
| persistence.enabled | bool | `true` | Enable persistence using PVC |
| persistence.existingClaim | string | `""` | Provide an existing `PersistentVolumeClaim` |
| persistence.mountPath | string | `"/bitnami/mariadb"` | Path to mount the volume at |
| persistence.selector | object | `{}` | Selector to match an existing Persistent Volume (this value is evaluated as a template) |
| persistence.size | string | `"8Gi"` | Persistent Volume Size |
| persistence.storageClass | string | `""` | Persistent Volume Storage Class |
| persistence.subPath | string | `""` | Subdirectory of the volume to mount |
| podAffinityPreset | string | `""` | Pod affinity preset. Ignored if `affinity` is set. Allowed values: `soft` or `hard` |
| podAnnotations | object | `{}` | Annotations for MariaDB Galera  pods |
| podAntiAffinityPreset | string | `"soft"` | Pod anti-affinity preset. Ignored if `affinity` is set. Allowed values: `soft` or `hard` |
| podDisruptionBudget.create | bool | `false` | Specifies whether a Pod disruption budget should be created |
| podDisruptionBudget.maxUnavailable | string | `""` | Maximum number / percentage of pods that may be made unavailable |
| podDisruptionBudget.minAvailable | int | `1` | Minimum number / percentage of pods that should remain scheduled |
| podLabels | object | `{}` | Extra labels for MariaDB Galera pods |
| podManagementPolicy | string | `"OrderedReady"` | StatefulSet controller supports relax its ordering guarantees while preserving its uniqueness and identity guarantees.  There are two valid pod management policies: OrderedReady and Parallel |
| priorityClassName | string | `""` | Priority Class Name for Statefulset |
| rbac.create | bool | `false` | Specify whether RBAC resources should be created and used |
| readinessProbe.enabled | bool | `true` | Turn on and off readiness probe |
| readinessProbe.failureThreshold | int | `3` | Minimum consecutive failures for the probe |
| readinessProbe.initialDelaySeconds | int | `30` | Delay before readiness probe is initiated |
| readinessProbe.periodSeconds | int | `10` | How often to perform the probe |
| readinessProbe.successThreshold | int | `1` | Minimum consecutive successes for the probe |
| readinessProbe.timeoutSeconds | int | `1` | When the probe times out |
| replicaCount | int | `3` | Desired number of cluster nodes |
| resources.limits | object | `{}` |  |
| resources.requests | object | `{}` |  |
| rootUser.forcePassword | bool | `false` | Option to force users to specify a password. That is required for 'helm upgrade' to work properly. If it is not force, a random password will be generated. |
| rootUser.password | string | `""` | Password for the admin user. Ignored if existing secret is provided. Password is ignored if existingSecret is specified. |
| rootUser.user | string | `"root"` | Username for the admin user. |
| schedulerName | string | `""` | Name of the Kubernetes scheduler (other than default) |
| securityContext.enabled | bool | `true` | Enable security context |
| securityContext.fsGroup | int | `1001` | Group ID for the container filesystem |
| securityContext.runAsUser | int | `1001` | User ID for the container |
| service.annotations | object | `{}` | Additional annotations for MariaDB Galera service |
| service.clusterIP | string | `""` | Specific cluster IP when service type is cluster IP. Use `None` for headless service |
| service.externalIPs | list | `[]` | External IP list to use with ClusterIP service type |
| service.headless.annotations | object | `{}` | Annotations for the headless service. May be useful for setting `service.alpha.kubernetes.io/tolerate-unready-endpoints="true"` when using peer-finder. |
| service.loadBalancerIP | string | `""` | `loadBalancerIP` if service type is `LoadBalancer` |
| service.loadBalancerSourceRanges | list | `[]` | Addresses that are allowed when svc is `LoadBalancer` |
| service.nodePort | string | `""` | Specify the nodePort value for the LoadBalancer and NodePort service types. |
| service.port | int | `3306` | MariaDB service port |
| service.type | string | `"ClusterIP"` | Kubernetes service type |
| serviceAccount.create | bool | `false` | Specify whether a ServiceAccount should be created |
| serviceAccount.name | string | `""` | The name of the ServiceAccount to create If not set and create is true, a name is generated using the common.names.fullname template |
| startupProbe.enabled | bool | `false` | Turn on and off startup probe |
| startupProbe.failureThreshold | int | `48` |  |
| startupProbe.initialDelaySeconds | int | `120` |  |
| startupProbe.periodSeconds | int | `10` | How often to perform the probe |
| startupProbe.successThreshold | int | `1` | Minimum consecutive successes for the probe |
| startupProbe.timeoutSeconds | int | `1` | When the probe times out |
| tls.autoGenerated | bool | `false` | Generate automatically self-signed TLS certificates |
| tls.certCAFilename | string | `""` | CA Certificate filename |
| tls.certFilename | string | `""` | Certificate filename |
| tls.certKeyFilename | string | `""` | Certificate key filename |
| tls.certificatesSecret | string | `""` | Name of the secret that contains the certificates |
| tls.enabled | bool | `false` | Enable TLS support for replication traffic |
| tolerations | list | `[]` | Tolerations for pod assignment |
| updateStrategy | object | `{"type":"RollingUpdate"}` | updateStrategy for MariaDB Master StatefulSet |
| usePasswordFiles | bool | `false` | Mount credentials as a files instead of using an environment variable. |

## Source Code

* <https://github.com/bitnami/bitnami-docker-mariadb-galera>
* <https://github.com/prometheus/mysqld_exporter>
* <https://mariadb.org>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | common | 1.x.x |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
