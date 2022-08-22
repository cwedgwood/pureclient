# PolicySmb

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Enabled** | **bool** | Returns a value of &#x60;true&#x60; if the policy is enabled. | [optional] [default to null]
**PolicyType** | **string** | Type of the policy. Valid values include &#x60;nfs&#x60;, &#x60;smb&#x60;, &#x60;snapshot&#x60;, and &#x60;quota&#x60;. | [optional] [default to null]
**AccessBasedEnumerationEnabled** | **bool** | Returns a value of &#x60;true&#x60; if access based enumeration is enabled on the policy. When access based enumeration is enabled on a policy, files and folders within exports that are attached to the policy will be hidden from users who do not have permission to view them. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

