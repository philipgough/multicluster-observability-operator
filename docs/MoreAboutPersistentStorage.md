# Persistent Stores used in Open Cluster Management Observability

Open Cluster Management Observability is a stateful application. 
It creates the following PersistentVolumes (the number of copies depend on the number of replicas).

## List of PersistentVolumes

| Name                   | Purpose                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| alertmanager           | Alertmanager stores the `nflog` data and silenced alerts in its storage. `nflog` is an append-only log of active and resolved notifications along with the notified receiver, and a hash digest of contents that the notificationn identified.                                                                                                                                                                                                                                                                                                                             |
| thanos-compact         | The compactor needs local disk space to store intermediate data for its processing, as well as bucket state cache. The required space depends on the size of the underlying blocks. The compactor must have enough space to download all of the source blocks, then build the compacted blocks on the disk. On-disk data is safe to delete between restarts and should be the first attempt to get crash-looping compactors unstuck. However, it is recommended to give the compactor persistent disks in order to effectively use bucket state cache in between restarts. |
| thanos-rule            | The thanos ruler evaluates Prometheus recording and alerting rules against a chosen query API by issuing queries at a fixed interval. Rule results are written back to the disk in the Prometheus 2.0 storage format. Rule results are written back to disk in the Prometheus 2.0 storage format. The amount of hours or days of data retained is exposed as an API parameter in `observability.open-cluster-management.io/v1beta2`: `_RetentionInLocal_`                                                                                                                  |
| thanos-receive-default | Thanos receiver accepts incoming data (Prometheus remote-write requests) and writes these into a local instance of the Prometheus TSDB. Periodically (every 2 hours), TSDB blocks are uploaded to the object storage for long term storage and compaction. The amount of hours or days of data retained is exposed as an API parameter in `observability.open-cluster-management.io/v1beta2`: `_RetentionInLocal_`                                                                                                                                                         |
| thanos-store-shard     | It acts primarily as an API gateway and therefore does not need significant amounts of local disk space. It joins a Thanos cluster on startup and advertises the data it can access. It keeps a small amount of information about all remote blocks on local disk and keeps it in sync with the bucket. This data is generally safe to delete across restarts at the cost of increased startup times.                                                                                                                                                                      |


## Configuring the StatefulSets

You can update the storage configuration for the stateful workloads individually in the
`observability.open-cluster-management.io/v1beta2` API, as shown in the following example:

```
    //defaults shown below
    StorageClass: gp2
    AlertmanagerStorageSize: 1Gi 
    RuleStorageSize: 1Gi
    CompactStorageSize: 100 Gi
    ReceiveStorageSize: 100 Gi
    StoreStorageSize: 10 Gi

```

**Note**: The default storage class, as configured in the system, is used for configuring the PersistentVolumes unless 
a different storage class is specified in the CustomResource specification. 
If no StorageClass exists (for example in default OpenShift bare metal installations), a StorageClass must be created 
and specified for a successful installation.

## Object Store

In addition to the PersistentVolumes previously mentioned, the time series historical data is stored in object stores.
Thanos uses object storage as the primary storage for metrics and metadata related to them.
