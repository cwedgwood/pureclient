# Model28VolumesPatchBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destroyed** | **bool** | If set to &#x60;true&#x60;, destroys a resource. Once set to &#x60;true&#x60;, the &#x60;time_remaining&#x60; value will display the amount of time left until the destroyed resource is permanently eradicated. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed resource can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the &#x60;time_remaining&#x60; period has elapsed, the resource is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**Name** | **string** | The new name for the resource. | [optional] [default to null]
**Pod** | [***interface{}**](interface{}.md) | Moves the volume into the specified pod. | [optional] [default to null]
**Provisioned** | **int64** | Updates the virtual size of the volume. Measured in bytes. | [optional] [default to null]
**Qos** | [***interface{}**](interface{}.md) | Sets QoS limits. | [optional] [default to null]
**VolumeGroup** | [***interface{}**](interface{}.md) | Adds the volume to the specified volume group. | [optional] [default to null]
**RequestedPromotionState** | **string** | Valid values are &#x60;promoted&#x60; and &#x60;demoted&#x60;. Patch &#x60;requested_promotion_state&#x60; to &#x60;demoted&#x60; to demote the volume so that the volume stops accepting write requests. Patch &#x60;requested_promotion_state&#x60; to &#x60;promoted&#x60; to promote the volume so that the volume starts accepting write requests. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

