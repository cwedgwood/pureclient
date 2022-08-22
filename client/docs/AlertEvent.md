# AlertEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Actual** | **string** | Actual condition at the time the alert is created. | [optional] [default to null]
**Alert** | [***AlertEventAlert**](AlertEvent_alert.md) |  | [optional] [default to null]
**Code** | **int64** | The parent alert number. | [optional] [default to null]
**ComponentName** | **string** | The component type of the alert. | [optional] [default to null]
**ComponentType** | **string** | The component name of the alert. | [optional] [default to null]
**Created** | **int64** | The time the parent alert was created. | [optional] [default to null]
**Expected** | **string** | Expected state and threshold under normal conditions. | [optional] [default to null]
**Issue** | **string** | Information about the alert cause. | [optional] [default to null]
**KnowledgeBaseUrl** | **string** | The knowledge base URL of the alert. | [optional] [default to null]
**Severity** | **string** | The severity level of the alert. Valid values include &#x60;info&#x60;, &#x60;warning&#x60;, &#x60;critical&#x60;, and &#x60;hidden&#x60;. | [optional] [default to null]
**State** | **string** | The state of the alert. Valid values include &#x60;open&#x60;, &#x60;closing&#x60;, and &#x60;closed&#x60;. | [optional] [default to null]
**Summary** | **string** | A summary of the alert. | [optional] [default to null]
**Time** | **int64** | The time the event occurred. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

