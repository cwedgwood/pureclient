# SoftwareInstallationStepsChecks

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | **string** | Detailed result of the check used to diagnose check failures. | [optional] [default to null]
**Name** | **string** | Name of the upgrade check. | [optional] [default to null]
**Overridable** | **bool** | Whether the check failure can be overridden. | [optional] [default to null]
**Status** | **string** | Status of the check. Valid values are &#x60;running&#x60;, &#x60;failed&#x60;, &#x60;passed&#x60;, and &#x60;overridden&#x60;. A status of &#x60;running&#x60; indicates that the check has not finished. A status of &#x60;failed&#x60; indicates that the check has failed. A status of &#x60;passed&#x60; indicates that the check has passed. A status of &#x60;overridden&#x60; indicates that the check has failed, but the failure has been overridden. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

