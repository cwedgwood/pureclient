# Port

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Iqn** | **string** | The iSCSI Qualified Name (or &#x60;null&#x60; if target is not iSCSI). | [optional] [default to null]
**Nqn** | **string** | NVMe Qualified Name (or &#x60;null&#x60; if target is not NVMeoF). | [optional] [default to null]
**Portal** | **string** | IP and port number (or &#x60;null&#x60; if target is not iSCSI). | [optional] [default to null]
**Wwn** | **string** | Fibre Channel World Wide Name (or &#x60;null&#x60; if target is not Fibre Channel). | [optional] [default to null]
**Failover** | **string** | If the array port has failed over, returns the name of the port to which this port has failed over. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

