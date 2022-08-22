# HostPerformanceBalance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**FractionRelativeToMax** | **float32** | The path with the highest number of op counts is displayed with a fraction_relative_to_max of 1.0. The fraction values of all other paths in the host are then calculated relative to the path with the highest number of op counts. | [optional] [default to null]
**Initiator** | [***HostPerformanceBalanceInitiator**](HostPerformanceBalance_initiator.md) |  | [optional] [default to null]
**OpCount** | **int64** | Count of I/O operations for the host path, over the specified resolution. | [optional] [default to null]
**Target** | [***Port**](Port.md) |  | [optional] [default to null]
**Time** | **int64** | Sample time in milliseconds since UNIX epoch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

