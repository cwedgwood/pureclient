# DirectoryQuota

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | [***AllOfDirectoryQuotaDirectory**](AllOfDirectoryQuotaDirectory.md) | The directory to which the quota applies. | [optional] [default to null]
**Enabled** | **bool** | Returns a value of &#x60;true&#x60; if the policy is enabled. | [optional] [default to null]
**Enforced** | **bool** | Defines whether the quota rule is enforced or unenforced. If the quota rule is enforced and logical space usage exceeds the quota limit, any modification operations that result in a need for more space are blocked. If the quota rule is unenforced and logical space usage exceeds the quota limit, notification emails are sent to targets that are specified using the &#x60;notification&#x60; parameter. No client operations are blocked when an unenforced limit is exceeded. If set to &#x60;true&#x60;, the limit is enforced. If set to &#x60;false&#x60;, notification targets are informed when the usage exceeds 80 percent of the limit. | [optional] [default to null]
**Path** | **string** | Absolute path of the directory in the file system. | [optional] [default to null]
**PercentageUsed** | **float32** | The percentage of the space used in the directory with respect to the quota limit. | [optional] [default to null]
**Policy** | [***AllOfDirectoryQuotaPolicy**](AllOfDirectoryQuotaPolicy.md) | The effective quota policy that imposes the limit. This is the policy with the lowest limit. | [optional] [default to null]
**QuotaLimit** | **int64** | Effective quota limit imposed by the quota policy rule attached to the directory, measured in bytes. | [optional] [default to null]
**RuleName** | **string** | Name of the rule that results in this quota and behavior being applied to this directory. | [optional] [default to null]
**Usage** | **int64** | The amount of logically written data for the directory, measured in bytes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

