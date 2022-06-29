# Audit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Arguments** | **string** | The arguments provided to the command. | [optional] [default to null]
**Command** | **string** | The top level command that starts with the string \&quot;pure\&quot; as a convention. | [optional] [default to null]
**Origin** | [***AuditOrigin**](Audit_origin.md) |  | [optional] [default to null]
**Subcommand** | **string** | The &#x60;command&#x60; and &#x60;subcommand&#x60; combination determines which action the user attempted to perform. | [optional] [default to null]
**Time** | **int64** | The time at which the command was run in milliseconds since the UNIX epoch. | [optional] [default to null]
**User** | **string** | The user who ran the command. | [optional] [default to null]
**UserInterface** | **string** | The user interface through which the user session event was performed. Valid values are &#x60;CLI&#x60;, &#x60;GUI&#x60;, and &#x60;REST&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

