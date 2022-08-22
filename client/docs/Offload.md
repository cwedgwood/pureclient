# Offload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Azure** | [***Api28offloadsAzure**](api2.8offloads_azure.md) |  | [optional] [default to null]
**GoogleCloud** | [***Api28offloadsGooglecloud**](api2.8offloads_googlecloud.md) |  | [optional] [default to null]
**Nfs** | [***Api28offloadsNfs**](api2.8offloads_nfs.md) |  | [optional] [default to null]
**S3** | [***Api28offloadsS3**](api2.8offloads_s3.md) |  | [optional] [default to null]
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Protocol** | **string** | Type of offload. Valid values include &#x60;azure&#x60;, &#x60;google-cloud&#x60;, &#x60;nfs&#x60;, and &#x60;s3&#x60;. | [optional] [default to null]
**TargetId** | **string** | Unique ID for the offload target. When multiple connections to one offload target are created, they each have distinct IDs but share the same &#x60;target_id&#x60;. | [optional] [default to null]
**Space** | [***OffloadSpace**](Offload_space.md) |  | [optional] [default to null]
**Status** | **string** | Offload status. Valid values are &#x60;connecting&#x60;, &#x60;connected&#x60;, &#x60;disconnecting&#x60;, &#x60;not connected&#x60;, and &#x60;scanning&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

