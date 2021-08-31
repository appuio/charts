# mariadb-galera

![Version: 1.1.0](https://img.shields.io/badge/Version-1.1.0-informational?style=flat-square) ![AppVersion: 10.5.12](https://img.shields.io/badge/AppVersion-10.5.12-informational?style=flat-square)

MariaDB Galera is a multi-master database cluster solution for synchronous replication and high availability.

**Homepage:** <https://github.com/appuio/charts/tree/master/appuio/mariadb-galera>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| APPUiO Team | info@appuio.ch |  |

## Source Code

* <https://github.com/bitnami/bitnami-docker-mariadb-galera>
* <https://github.com/prometheus/mysqld_exporter>
* <https://mariadb.org>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | common | 1.x.x |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| clusterDomain | string | `"cluster.local"` |  |
| configurationConfigMap | string | `""` |  |
| customPasswordFiles | object | `{}` |  |
| db.forcePassword | bool | `false` |  |
| db.name | string | `"my_database"` |  |
| db.password | string | `""` |  |
| db.user | string | `""` |  |
| diagnosticMode.args[0] | string | `"infinity"` |  |
| diagnosticMode.command[0] | string | `"sleep"` |  |
| diagnosticMode.enabled | bool | `false` |  |
| existingSecret | string | `""` |  |
| extraContainers | list | `[]` |  |
| extraDeploy | list | `[]` |  |
| extraEnvVars | list | `[]` |  |
| extraEnvVarsCM | string | `""` |  |
| extraEnvVarsSecret | string | `""` |  |
| extraFlags | string | `""` |  |
| extraInitContainers | list | `[]` |  |
| extraVolumeMounts | list | `[]` |  |
| extraVolumes | list | `[]` |  |
| forceUpdate | bool | `false` |  |
| fullnameOverride | string | `""` |  |
| galera.bootstrap.bootstrapFromNode | string | `""` |  |
| galera.bootstrap.forceSafeToBootstrap | bool | `false` |  |
| galera.mariabackup.forcePassword | bool | `false` |  |
| galera.mariabackup.password | string | `""` |  |
| galera.mariabackup.user | string | `"mariabackup"` |  |
| galera.name | string | `"galera"` |  |
| global.imagePullSecrets | list | `[]` |  |
| global.imageRegistry | string | `""` |  |
| global.storageClass | string | `""` |  |
| hostAliases | list | `[]` |  |
| image.debug | bool | `false` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.pullSecrets | list | `[]` |  |
| image.registry | string | `"docker.io"` |  |
| image.repository | string | `"bitnami/mariadb-galera"` |  |
| image.tag | string | `"10.5.12-debian-10-r1"` |  |
| initdbScripts | object | `{}` |  |
| initdbScriptsConfigMap | string | `""` |  |
| ldap.base | string | `""` |  |
| ldap.binddn | string | `""` |  |
| ldap.bindpw | string | `""` |  |
| ldap.bslookup | string | `""` |  |
| ldap.enabled | bool | `false` |  |
| ldap.filter | string | `""` |  |
| ldap.map | string | `""` |  |
| ldap.nss_initgroups_ignoreusers | string | `"root,nslcd"` |  |
| ldap.scope | string | `""` |  |
| ldap.tls_reqcert | string | `""` |  |
| ldap.uri | string | `""` |  |
| livenessProbe.enabled | bool | `true` |  |
| livenessProbe.failureThreshold | int | `3` |  |
| livenessProbe.initialDelaySeconds | int | `120` |  |
| livenessProbe.periodSeconds | int | `10` |  |
| livenessProbe.successThreshold | int | `1` |  |
| livenessProbe.timeoutSeconds | int | `1` |  |
| mariadbConfiguration | string | `"[client]\nport=3306\nsocket=/opt/bitnami/mariadb/tmp/mysql.sock\nplugin_dir=/opt/bitnami/mariadb/plugin\n\n[mysqld]\ndefault_storage_engine=InnoDB\nbasedir=/opt/bitnami/mariadb\ndatadir=/bitnami/mariadb/data\nplugin_dir=/opt/bitnami/mariadb/plugin\ntmpdir=/opt/bitnami/mariadb/tmp\nsocket=/opt/bitnami/mariadb/tmp/mysql.sock\npid_file=/opt/bitnami/mariadb/tmp/mysqld.pid\nbind_address=0.0.0.0\n\n## Character set\n##\ncollation_server=utf8_unicode_ci\ninit_connect='SET NAMES utf8'\ncharacter_set_server=utf8\n\n## MyISAM\n##\nkey_buffer_size=32M\nmyisam_recover_options=FORCE,BACKUP\n\n## Safety\n##\nskip_host_cache\nskip_name_resolve\nmax_allowed_packet=16M\nmax_connect_errors=1000000\nsql_mode=STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_AUTO_VALUE_ON_ZERO,NO_ENGINE_SUBSTITUTION,NO_ZERO_DATE,NO_ZERO_IN_DATE,ONLY_FULL_GROUP_BY\nsysdate_is_now=1\n\n## Binary Logging\n##\nlog_bin=mysql-bin\nexpire_logs_days=14\n# Disabling for performance per http://severalnines.com/blog/9-tips-going-production-galera-cluster-mysql\nsync_binlog=0\n# Required for Galera\nbinlog_format=row\n\n## Caches and Limits\n##\ntmp_table_size=32M\nmax_heap_table_size=32M\n# Re-enabling as now works with Maria 10.1.2\nquery_cache_type=1\nquery_cache_limit=4M\nquery_cache_size=256M\nmax_connections=500\nthread_cache_size=50\nopen_files_limit=65535\ntable_definition_cache=4096\ntable_open_cache=4096\n\n## InnoDB\n##\ninnodb=FORCE\ninnodb_strict_mode=1\n# Mandatory per https://github.com/codership/documentation/issues/25\ninnodb_autoinc_lock_mode=2\n# Per https://www.percona.com/blog/2006/08/04/innodb-double-write/\ninnodb_doublewrite=1\ninnodb_flush_method=O_DIRECT\ninnodb_log_files_in_group=2\ninnodb_log_file_size=128M\ninnodb_flush_log_at_trx_commit=1\ninnodb_file_per_table=1\n# 80% Memory is default reco.\n# Need to re-evaluate when DB size grows\ninnodb_buffer_pool_size=2G\ninnodb_file_format=Barracuda\n\n## Logging\n##\nlog_error=/opt/bitnami/mariadb/logs/mysqld.log\nslow_query_log_file=/opt/bitnami/mariadb/logs/mysqld.log\nlog_queries_not_using_indexes=1\nslow_query_log=1\n\n## SSL\n## Use extraVolumes and extraVolumeMounts to mount /certs filesystem\n# ssl_ca=/certs/ca.pem\n# ssl_cert=/certs/server-cert.pem\n# ssl_key=/certs/server-key.pem\n\n[galera]\nwsrep_on=ON\nwsrep_provider=/opt/bitnami/mariadb/lib/libgalera_smm.so\nwsrep_sst_method=mariabackup\nwsrep_slave_threads=4\nwsrep_cluster_address=gcomm://\nwsrep_cluster_name=galera\nwsrep_sst_auth=\"root:\"\n# Enabled for performance per https://mariadb.com/kb/en/innodb-system-variables/#innodb_flush_log_at_trx_commit\ninnodb_flush_log_at_trx_commit=2\n# MYISAM REPLICATION SUPPORT #\nwsrep_replicate_myisam=ON\n\n[mariadb]\nplugin_load_add=auth_pam\n\n## Data-at-Rest Encryption\n## Use extraVolumes and extraVolumeMounts to mount /encryption filesystem\n# plugin_load_add=file_key_management\n# file_key_management_filename=/encryption/keyfile.enc\n# file_key_management_filekey=FILE:/encryption/keyfile.key\n# file_key_management_encryption_algorithm=AES_CTR\n# encrypt_binlog=ON\n# encrypt_tmp_files=ON\n\n## InnoDB/XtraDB Encryption\n# innodb_encrypt_tables=ON\n# innodb_encrypt_temporary_tables=ON\n# innodb_encrypt_log=ON\n# innodb_encryption_threads=4\n# innodb_encryption_rotate_key_age=1\n\n## Aria Encryption\n# aria_encrypt_tables=ON\n# encrypt_tmp_disk_tables=ON"` |  |
| metrics.enabled | bool | `false` |  |
| metrics.extraFlags | list | `[]` |  |
| metrics.image.pullPolicy | string | `"IfNotPresent"` |  |
| metrics.image.pullSecrets | list | `[]` |  |
| metrics.image.registry | string | `"docker.io"` |  |
| metrics.image.repository | string | `"bitnami/mysqld-exporter"` |  |
| metrics.image.tag | string | `"0.13.0-debian-10-r75"` |  |
| metrics.prometheusRules.enabled | bool | `false` |  |
| metrics.prometheusRules.rules | object | `{}` |  |
| metrics.prometheusRules.selector.app | string | `"prometheus-operator"` |  |
| metrics.prometheusRules.selector.release | string | `"prometheus"` |  |
| metrics.resources.limits | object | `{}` |  |
| metrics.resources.requests | object | `{}` |  |
| metrics.service.annotations."prometheus.io/port" | string | `"9104"` |  |
| metrics.service.annotations."prometheus.io/scrape" | string | `"true"` |  |
| metrics.service.port | int | `9104` |  |
| metrics.service.type | string | `"ClusterIP"` |  |
| metrics.serviceMonitor.enabled | bool | `false` |  |
| metrics.serviceMonitor.interval | string | `""` |  |
| metrics.serviceMonitor.metricRelabelings | list | `[]` |  |
| metrics.serviceMonitor.namespace | string | `""` |  |
| metrics.serviceMonitor.relabelings | list | `[]` |  |
| metrics.serviceMonitor.scrapeTimeout | string | `""` |  |
| metrics.serviceMonitor.selector.prometheus | string | `"kube-prometheus"` |  |
| nameOverride | string | `""` |  |
| nodeAffinityPreset.key | string | `""` |  |
| nodeAffinityPreset.type | string | `""` |  |
| nodeAffinityPreset.values | list | `[]` |  |
| nodeSelector | object | `{}` |  |
| persistence.accessModes[0] | string | `"ReadWriteOnce"` |  |
| persistence.annotations | object | `{}` |  |
| persistence.enabled | bool | `true` |  |
| persistence.existingClaim | string | `""` |  |
| persistence.mountPath | string | `"/bitnami/mariadb"` |  |
| persistence.selector | object | `{}` |  |
| persistence.size | string | `"8Gi"` |  |
| persistence.storageClass | string | `""` |  |
| persistence.subPath | string | `""` |  |
| podAffinityPreset | string | `""` |  |
| podAnnotations | object | `{}` |  |
| podAntiAffinityPreset | string | `"soft"` |  |
| podDisruptionBudget.create | bool | `false` |  |
| podDisruptionBudget.maxUnavailable | string | `""` |  |
| podDisruptionBudget.minAvailable | int | `1` |  |
| podLabels | object | `{}` |  |
| podManagementPolicy | string | `"OrderedReady"` |  |
| priorityClassName | string | `""` |  |
| rbac.create | bool | `false` |  |
| readinessProbe.enabled | bool | `true` |  |
| readinessProbe.failureThreshold | int | `3` |  |
| readinessProbe.initialDelaySeconds | int | `30` |  |
| readinessProbe.periodSeconds | int | `10` |  |
| readinessProbe.successThreshold | int | `1` |  |
| readinessProbe.timeoutSeconds | int | `1` |  |
| replicaCount | int | `3` |  |
| resources.limits | object | `{}` |  |
| resources.requests | object | `{}` |  |
| rootUser.forcePassword | bool | `false` |  |
| rootUser.password | string | `""` |  |
| rootUser.user | string | `"root"` |  |
| schedulerName | string | `""` |  |
| securityContext.enabled | bool | `true` |  |
| securityContext.fsGroup | int | `1001` |  |
| securityContext.runAsUser | int | `1001` |  |
| service.annotations | object | `{}` |  |
| service.clusterIP | string | `""` |  |
| service.externalIPs | list | `[]` |  |
| service.headless.annotations | object | `{}` |  |
| service.loadBalancerIP | string | `""` |  |
| service.loadBalancerSourceRanges | list | `[]` |  |
| service.nodePort | string | `""` |  |
| service.port | int | `3306` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.create | bool | `false` |  |
| serviceAccount.name | string | `""` |  |
| startupProbe.enabled | bool | `false` |  |
| startupProbe.failureThreshold | int | `48` |  |
| startupProbe.initialDelaySeconds | int | `120` |  |
| startupProbe.periodSeconds | int | `10` |  |
| startupProbe.successThreshold | int | `1` |  |
| startupProbe.timeoutSeconds | int | `1` |  |
| tls.autoGenerated | bool | `false` |  |
| tls.certCAFilename | string | `""` |  |
| tls.certFilename | string | `""` |  |
| tls.certKeyFilename | string | `""` |  |
| tls.certificatesSecret | string | `""` |  |
| tls.enabled | bool | `false` |  |
| tolerations | list | `[]` |  |
| updateStrategy.type | string | `"RollingUpdate"` |  |
| usePasswordFiles | bool | `false` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
