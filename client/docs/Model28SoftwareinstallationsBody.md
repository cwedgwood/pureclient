# Model28SoftwareinstallationsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Mode the upgrade is in. Valid values are &#x60;check-only&#x60;, &#x60;interactive&#x60;, &#x60;semi-interactive&#x60;, and &#x60;one-click&#x60;. In &#x60;check-only&#x60; mode, the upgrade only runs pre-upgrade checks and returns. In &#x60;interactive&#x60; mode, the upgrade process pauses at several points, at which users must enter certain commands to proceed. In &#x60;semi-interactive&#x60; mode, the upgrade pauses if there are any upgrade check failures and functions like &#x60;one-click&#x60; mode otherwise. In &#x60;one-click&#x60; mode, the upgrade proceeds automatically without pausing. | [optional] [default to null]
**OverrideChecks** | [**[]Api28softwareinstallationsOverrideChecks**](api2.8softwareinstallations_override_checks.md) | A list of upgrade checks whose failure is overridden during the upgrade. If any optional &#x60;args&#x60; are provided, they are validated later when the corresponding check script runs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

