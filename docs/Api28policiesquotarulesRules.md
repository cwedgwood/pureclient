# Api28policiesquotarulesRules

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enforced** | **bool** | If set to &#x60;true&#x60;, this rule describes an enforced quota. An out-of-space warning is issued if logical space usage exceeds the limit value described in this rule. If set to &#x60;false&#x60;, this rule describes an unenforced quota. Alerts and/or notifications are issued when logical space usage exceeds the limit value described in this rule. If not specified, defaults to &#x60;false&#x60;. | [optional] [default to null]
**QuotaLimit** | **int64** | Logical space limit of the quota (in bytes) assigned by the rule. This value cannot be set to 0. | [optional] [default to null]
**Notifications** | **string** | Targets to notify when usage approaches the quota limit. The list of notification targets is a comma-separated string. Valid values are &#x60;user&#x60;, and &#x60;group&#x60;. If not specified, notification targets are not assigned for the rule. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

