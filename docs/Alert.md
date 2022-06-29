# Alert

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Actual** | **string** | Actual condition at the time the alert is created. | [optional] [default to null]
**Category** | **string** | The category of the alert. Valid values include &#x60;array&#x60;, &#x60;hardware&#x60; and &#x60;software&#x60;. | [optional] [default to null]
**Closed** | **int64** | The time the alert was closed in milliseconds since the UNIX epoch. | [optional] [default to null]
**Code** | **int64** | The code number of the alert. | [optional] [default to null]
**ComponentName** | **string** | The name of the component that generated the alert. | [optional] [default to null]
**ComponentType** | **string** | The type of component that generated the alert. | [optional] [default to null]
**Created** | **int64** | The time the alert was created in milliseconds since the UNIX epoch. | [optional] [default to null]
**Description** | **string** | A short description of the alert. | [optional] [default to null]
**Expected** | **string** | Expected state or threshold under normal conditions. | [optional] [default to null]
**Flagged** | **bool** | If set to &#x60;true&#x60;, the message is flagged. Important messages can can be flagged and listed separately. | [optional] [default to null]
**Issue** | **string** | Information about the alert cause. | [optional] [default to null]
**KnowledgeBaseUrl** | **string** | The URL of the relevant knowledge base page. | [optional] [default to null]
**Notified** | **int64** | The time the most recent alert notification was sent in milliseconds since the UNIX epoch. | [optional] [default to null]
**Severity** | **string** | The severity level of the alert. Valid values include &#x60;info&#x60;, &#x60;warning&#x60;, &#x60;critical&#x60;, and &#x60;hidden&#x60;. | [optional] [default to null]
**State** | **string** | The current state of the alert. Valid values include &#x60;open&#x60;, &#x60;closing&#x60;, and &#x60;closed&#x60;. | [optional] [default to null]
**Summary** | **string** | A summary of the alert. | [optional] [default to null]
**Updated** | **int64** | The time the alert was last updated in milliseconds since the UNIX epoch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

