# SoftwareInstallationSteps

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified. | [optional] [default to null]
**Name** | **string** | Name of the resource. The name cannot be modified. | [optional] [default to null]
**StartTime** | **int64** | Start time in milliseconds since the UNIX epoch. | [optional] [default to null]
**EndTime** | **int64** | End time in milliseconds since the UNIX epoch. | [optional] [default to null]
**Checks** | [**[]SoftwareInstallationStepsChecks**](SoftwareInstallationSteps_checks.md) | A list of checks in this upgrade step. | [optional] [default to null]
**Description** | **string** | Detailed description of the step. | [optional] [default to null]
**Details** | **string** | Detailed result of the step used to diagnose step failures. | [optional] [default to null]
**HopVersion** | **string** | The version to which the current hop is upgrading. | [optional] [default to null]
**Installation** | [***SoftwareInstallationStepsInstallation**](SoftwareInstallationSteps_installation.md) |  | [optional] [default to null]
**Status** | **string** | Status of the step. Valid values are &#x60;running&#x60; and &#x60;finished&#x60;. A status of &#x60;running&#x60; indicates that the step has not finished. A status of &#x60;finished&#x60; indicates that the check has finished. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

