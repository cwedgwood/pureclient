# Drive

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Capacity** | **int64** | Physical storage capacity of the module in bytes. | [optional] [default to null]
**Details** | **string** | Details about the status of the module if not healthy. | [optional] [default to null]
**Protocol** | **string** | Storage protocol of the module. Valid values are &#x60;NVMe&#x60; and &#x60;SAS&#x60;. | [optional] [default to null]
**Status** | **string** | Current status of the module. Valid values are &#x60;empty&#x60;, &#x60;failed&#x60;, &#x60;healthy&#x60;, &#x60;identifying&#x60;, &#x60;missing&#x60;, &#x60;recovering&#x60;, &#x60;unadmitted&#x60;, &#x60;unhealthy&#x60;, &#x60;unrecognized&#x60;, and &#x60;updating&#x60;. | [optional] [default to null]
**Type_** | **string** | The type of the module. Valid values are &#x60;cache&#x60;, &#x60;NVRAM&#x60;, &#x60;SSD&#x60;, and &#x60;virtual&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

