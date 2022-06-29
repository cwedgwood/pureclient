# SoftwareInstallations

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified. | [optional] [default to null]
**Name** | **string** | Name of the resource. The name cannot be modified. | [optional] [default to null]
**StartTime** | **int64** | Start time in milliseconds since the UNIX epoch. | [optional] [default to null]
**EndTime** | **int64** | End time in milliseconds since the UNIX epoch. | [optional] [default to null]
**CurrentStepId** | **string** | The &#x60;id&#x60; of the current step, or &#x60;null&#x60; if the upgrade is not active. | [optional] [default to null]
**Details** | **string** | The detailed reason for the &#x60;status&#x60;. | [optional] [default to null]
**Mode** | **string** | Mode that the upgrade is in. Valid values are &#x60;check-only&#x60;, &#x60;interactive&#x60;, &#x60;semi-interactive&#x60;, and &#x60;one-click&#x60;. In &#x60;check-only&#x60; mode, the upgrade only runs pre-upgrade checks and returns. In &#x60;interactive&#x60; mode, the upgrade process pauses at several points, at which users must enter certain commands to proceed. In &#x60;semi-interactive&#x60; mode, the upgrade pauses if there are any upgrade check failures, and functions like &#x60;one-click&#x60; mode otherwise. In &#x60;one-click&#x60; mode, the upgrade proceeds automatically without pausing. | [optional] [default to null]
**OverrideChecks** | [**[]Api28softwareinstallationsOverrideChecks**](api2.8softwareinstallations_override_checks.md) | A list of upgrade checks whose failure is overridden during the upgrade. If any optional &#x60;args&#x60; are provided, they are validated later when the corresponding check script runs. | [optional] [default to null]
**Software** | [***SoftwareInstallationsSoftware**](SoftwareInstallations_software.md) |  | [optional] [default to null]
**Status** | **string** | Status of the upgrade. Valid values are &#x60;installing&#x60;, &#x60;paused&#x60;, &#x60;aborting&#x60;, &#x60;aborted&#x60;, and &#x60;finished&#x60;. A status of &#x60;installing&#x60; indicates that the upgrade is running. A status of &#x60;paused&#x60; indicates that the upgrade is paused and waiting for user input. A status of &#x60;aborting&#x60; indicates that the upgrade is being aborted. A status of &#x60;aborted&#x60; indicates that the upgrade has stopped due to an abort. A status of &#x60;finished&#x60; indicates that the upgrade has finished successfully. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

