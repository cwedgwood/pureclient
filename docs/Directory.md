# Directory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Created** | **int64** | The managed directory creation time, measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**Destroyed** | **bool** | Returns a value of &#x60;true&#x60; if the managed directory has been destroyed and is pending eradication. The &#x60;time_remaining&#x60; value displays the amount of time left until the destroyed managed directory is permanently eradicated. Once the &#x60;time_remaining&#x60; period has elapsed, the managed directory is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**DirectoryName** | **string** | The managed directory name without the file system name prefix. A full managed directory name is constructed in the form of &#x60;FILE_SYSTEM:DIR&#x60; where &#x60;FILE_SYSTEM&#x60; is the file system name and &#x60;DIR&#x60; is the value of this field. | [optional] [default to null]
**FileSystem** | [***interface{}**](interface{}.md) | The file system that this managed directory is in. | [optional] [default to null]
**Path** | **string** | Absolute path of the managed directory in the file system. | [optional] [default to null]
**Space** | [***interface{}**](interface{}.md) | Displays size and space consumption information. | [optional] [default to null]
**TimeRemaining** | **int64** | The amount of time left, measured in milliseconds until the destroyed managed directory is permanently eradicated. | [optional] [default to null]
**LimitedBy** | [***DirectoryLimitedBy**](Directory_limited_by.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

