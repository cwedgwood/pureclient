# ProtectionGroupSnapshotSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | **int64** | The time of day the snapshot is scheduled to be taken and retained on the local array or immediately replicated to the target(s). Measured in milliseconds since midnight. The &#x60;at&#x60; value is only used if the &#x60;frequency&#x60; parameter is in days (e.g., &#x60;259200000&#x60;, which is equal to 3 days). | [optional] [default to null]
**Enabled** | **bool** | If set to &#x60;true&#x60;, the policy is enabled. | [optional] [default to null]
**Frequency** | **int64** | The frequency of the scheduled action. Measured in milliseconds. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

