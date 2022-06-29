# ArrayConnectionPath

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**LocalPort** | **string** | The local port of the path. | [optional] [default to null]
**LocalAddress** | **string** | IP address or WWN of the local port. | [optional] [default to null]
**RemotePort** | **string** | The remote port of the path. | [optional] [default to null]
**RemoteAddress** | **string** | IP address or WWN of the remote port. | [optional] [default to null]
**Status** | **string** | Status of the connection. Valid values are &#x60;connected&#x60; and &#x60;connecting&#x60;. A status of &#x60;connected&#x60; indicates that the arrays are communicating. A status of &#x60;connecting&#x60; indicates that the array is trying to establish a connection. | [optional] [default to null]
**ReplicationTransport** | **string** | The protocol used to transport data between the local array and the remote array. Valid values are &#x60;ip&#x60; and &#x60;fc&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

