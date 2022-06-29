# Model28VolumesBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destroyed** | **bool** | If set to &#x60;true&#x60;, destroys a resource. Once set to &#x60;true&#x60;, the &#x60;time_remaining&#x60; value will display the amount of time left until the destroyed resource is permanently eradicated. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed resource can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the &#x60;time_remaining&#x60; period has elapsed, the resource is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**Provisioned** | **int64** | Sets the virtual size of the volume. Measured in bytes. | [optional] [default to null]
**Qos** | [***Api28volumesQos**](api2.8volumes_qos.md) |  | [optional] [default to null]
**Source** | [***Api28volumesSource**](api2.8volumes_source.md) |  | [optional] [default to null]
**Subtype** | **string** | The type of volume. Valid values are &#x60;protocol_endpoint&#x60; and &#x60;regular&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

