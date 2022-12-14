# PolicyRuleQuota

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enforced** | **bool** | Defines whether the quota rule is enforced or unenforced. If the quota rule is enforced and logical space usage exceeds the quota limit, any modification operations that result in a need for more space are blocked. If the quota rule is unenforced and logical space usage exceeds the quota limit, notification emails are sent to targets that are specified using the &#x60;notification&#x60; parameter. No client operations are blocked when an unenforced limit is exceeded. If set to &#x60;true&#x60;, the limit is enforced. If set to &#x60;false&#x60;, notification targets are informed when the usage exceeds 80 percent of the limit. | [optional] [default to null]
**Name** | **string** | Name of this rule. The name is automatically generated by the system. | [optional] [default to null]
**Notifications** | **string** | Targets to notify when usage approaches the quota limit. The list of notification targets is a comma-separated string. Valid values are &#x60;user&#x60;, and &#x60;group&#x60;. If not specified, notification targets are not assigned for the rule. | [optional] [default to null]
**Policy** | [***AllOfPolicyRuleQuotaPolicy**](AllOfPolicyRuleQuotaPolicy.md) | The policy to which this rule belongs. | [optional] [default to null]
**QuotaLimit** | **int64** | Logical space limit of the quota assigned by the rule, measured in bytes. This value cannot be set to 0. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

