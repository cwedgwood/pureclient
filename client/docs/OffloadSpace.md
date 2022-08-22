# OffloadSpace

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataReduction** | **float32** | The ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5&amp;#58;1 means that for every 5 MB the host writes to the array, 1 MB is stored on the array&#x27;s flash modules. | [optional] [default to null]
**Shared** | **int64** | The physical space occupied by deduplicated data, meaning that the space is shared with other volumes and snapshots as a result of data deduplication. Measured in bytes. | [optional] [default to null]
**Snapshots** | **int64** | The physical space occupied by data unique to one or more snapshots. Measured in bytes. | [optional] [default to null]
**System** | **int64** | The physical space occupied by internal array metadata. Measured in bytes. | [optional] [default to null]
**ThinProvisioning** | **float32** | The percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed. | [optional] [default to null]
**TotalPhysical** | **int64** | The total physical space occupied by system, shared space, volume, and snapshot data. Measured in bytes. | [optional] [default to null]
**TotalProvisioned** | **int64** | For a single volume, the provisioned size of the volume. For all other resources, the total provisioned size of all volumes that are connected to or are inside the resource. Represents storage capacity reported to hosts. Measured in bytes. | [optional] [default to null]
**TotalReduction** | **float32** | The ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10&amp;#58;1 means that for every 10 MB of provisioned space, 1 MB is stored on the array&#x27;s flash modules. | [optional] [default to null]
**Unique** | **int64** | The unique physical space occupied by customer data. Unique physical space does not include shared space, snapshots, and internal array metadata. Measured in bytes. | [optional] [default to null]
**Virtual** | **int64** | The amount of logically written data that a volume or a snapshot references. Measured in bytes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

