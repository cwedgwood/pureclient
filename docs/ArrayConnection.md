# ArrayConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**ManagementAddress** | **string** | Management IP address or FQDN of the target array. | [optional] [default to null]
**ReplicationAddresses** | **[]string** | IP addresses of the target arrays when &#x60;replication_transport&#x60; is &#x60;ip&#x60;. WWNs of the target arrays when &#x60;replication_transport&#x60; is &#x60;fc&#x60;. | [optional] [default to null]
**Status** | **string** | Status of the connection. Valid values are &#x60;connected&#x60;, &#x60;connecting&#x60;, &#x60;partially_connected&#x60;, and &#x60;unbalanced&#x60;. A status of &#x60;connected&#x60; indicates that arrays are communicating. A status of &#x60;connecting&#x60; indicates that the array is trying to establish a connection. A status of &#x60;partially_connected&#x60; indicates that some replication addresses are communicating but others are not. A status of &#x60;unbalanced&#x60; indicates that the arrays are communicating, but the set of paths is either not redundant or not symmetric. | [optional] [default to null]
**Type_** | **string** | The type of replication. Valid values are &#x60;async-replication&#x60; and &#x60;sync-replication&#x60;. | [optional] [default to null]
**ReplicationTransport** | **string** | The protocol used to transport data betwen the local array and the remote array. Valid values are &#x60;ip&#x60; and &#x60;fc&#x60;. The default is &#x60;ip&#x60;. | [optional] [default to null]
**Version** | **string** | The version of the target array. | [optional] [default to null]
**Throttle** | [***ArrayConnectionThrottle**](ArrayConnection_throttle.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

