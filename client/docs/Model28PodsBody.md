# Model28PodsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**FailoverPreferences** | [**[]Api28hostsPreferredArrays**](api2.8hosts_preferred_arrays.md) | Determines which array within a stretched pod should be given priority to stay online should the arrays ever lose contact with each other. The current array and any peer arrays that are connected to the current array for synchronous replication can be added to a pod for failover preference. By default, &#x60;failover_preferences&#x3D;null&#x60;, meaning no arrays have been configured for failover preference. Enter multiple arrays in comma-separated format. To clear the list of failover preferences, set to an empty list &#x60;[]&#x60;. | [optional] [default to null]
**Source** | [***Api28podsSource**](api2.8pods_source.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

