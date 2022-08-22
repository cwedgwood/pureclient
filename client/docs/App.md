# App

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Description** | **string** | A description of the app. | [optional] [default to null]
**Enabled** | **bool** | If set to &#x60;true&#x60;, the app is enabled. By default, apps are disabled. | [optional] [default to null]
**Status** | **string** | The status of the app. Values include &#x60;healthy&#x60; and &#x60;unhealthy&#x60;. For cluster apps, this represents the aggregate status of the app. The aggregate status is only &#x60;healthy&#x60; if all nodes are &#x60;healthy&#x60;&amp;#59; otherwise, it is &#x60;unhealthy&#x60;. | [optional] [default to null]
**Version** | **string** | The app version. For cluster apps, this represents the node version if all nodes are of the same version. If the node versions differ, a value of &#x60;null&#x60; is returned. | [optional] [default to null]
**Details** | **string** | Details of the status of the app. | [optional] [default to null]
**VncEnabled** | **bool** | If set to &#x60;true&#x60;, VNC access is enabled. By default, VNC access is disabled. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

