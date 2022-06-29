# Model28ArrayconnectionsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementAddress** | **string** | Management IP address of the target array. | [optional] [default to null]
**ReplicationAddresses** | **[]string** | IP addresses and FQDNs of the target arrays. Configurable only when &#x60;replication_transport&#x60; is set to &#x60;ip&#x60;. If not configured, will be set to all the replication addresses available on the target array at the time of the POST. | [optional] [default to null]
**Type_** | **string** | The type of replication. Valid values are &#x60;async-replication&#x60; and &#x60;sync-replication&#x60;. | [optional] [default to null]
**ReplicationTransport** | **string** | The protocol used to transport data betwen the local array and the remote array. Valid values are &#x60;ip&#x60; and &#x60;fc&#x60;. The default is &#x60;ip&#x60;. | [optional] [default to null]
**ConnectionKey** | **string** | The connection key of the target array. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

