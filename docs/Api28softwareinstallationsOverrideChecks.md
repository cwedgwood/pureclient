# Api28softwareinstallationsOverrideChecks

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the upgrade check to be overridden so the software upgrade can continue if the check failed or is anticipated to fail during the upgrade process. Overriding the check forces the system to ignore the check failure and continue with the upgrade. If the check includes more specific checks that failed or are anticipated to fail, set them using the &#x60;args&#x60; parameter. For example, the HostIOCheck check may include a list of hosts that have failed or are anticipated to fail the upgrade check. | [optional] [default to null]
**Args** | **string** | The name of the specific check within the override check to ignore so that the system can continue with the software upgrade. The &#x60;name&#x60; parameter of the override check must be specified with the &#x60;args&#x60; parameter. For example, if the HostIOCheck check fails on hosts host01 and host02, the system displays a list of these host names in the failed check. To override the HostIOCheck checks for host01 and host02, set &#x60;name&#x3D;HostIOCheck&#x60;, and set &#x60;args&#x3D;host01,host02&#x60;. Enter multiple &#x60;args&#x60; in comma-separated format. Note that not all checks have &#x60;args&#x60; values. | [optional] [default to null]
**Persistent** | **bool** | If set to &#x60;true&#x60;, the system always ignores the failure of the specified upgrade check and continues with the upgrade process. If set to &#x60;false&#x60;, the system ignores the failure of the specified upgrade check until the upgrade check finishes and the upgrade process is continued. For example, the &#x60;continue&#x60; command is successfully issued in an &#x60;interactive&#x60; mode, or the first upgrade check step successfully finishes in a &#x60;one-click&#x60; mode. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

