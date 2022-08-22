# Software

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified. | [optional] [default to null]
**Name** | **string** | Name of the resource. The name cannot be modified. | [optional] [default to null]
**Details** | **string** | The detailed reason of the &#x60;status&#x60;. | [optional] [default to null]
**PayloadId** | **string** | A checksum hash referring to the update bundle. | [optional] [default to null]
**Progress** | **float32** | The progress of the software upgrade. Displayed in decimal format. | [optional] [default to null]
**Status** | **string** | The status of the software package. Valid values are &#x60;available&#x60;, &#x60;downloaded&#x60;, &#x60;downloading&#x60;, &#x60;download_failed&#x60;, &#x60;checking&#x60;, &#x60;installing&#x60;, &#x60;paused&#x60;, &#x60;aborting&#x60;, &#x60;abort&#x60;, &#x60;canceled&#x60;, &#x60;partially_installed&#x60;, and &#x60;installed&#x60;. A status of &#x60;available&#x60; indicates that the package is available for download. This only applies if &#x60;automatic-download&#x60; is not enabled. A status of &#x60;downloaded&#x60; indicates that the package is downloaded and ready for installation. A status of &#x60;downloading&#x60; indicates that the package is currently downloading. A status of &#x60;download_failed&#x60; indicates that the download of the package failed. A status of &#x60;checking&#x60; indicates that the package is currently running in &#x60;check-only&#x60; mode. A status of &#x60;installing&#x60; indicates that the package is currently installing. A status of &#x60;paused&#x60; indicates that the upgrade is paused and waiting for user input to proceed. A status of &#x60;aborting&#x60; indicates that the upgrade is being aborted, due to an unrecoverable error or an &#x60;abort&#x60; command issued by the user. A status of &#x60;canceled&#x60; indicates that the upgrade has been canceled. A status of &#x60;partially_installed&#x60; indicates that the upgrade has been partially installed due to an &#x60;abort&#x60;. The array has been upgraded to an intermediate version and the &#x60;software&#x60; is no longer available for installation. A status of &#x60;installed&#x60; indicates that the upgrade has finished. | [optional] [default to null]
**UpgradeHops** | **[]string** | By which plan the upgrade will be conducted. The first element is the current version, the last element is the destination version, and the elements in between are intermediate versions. | [optional] [default to null]
**Version** | **string** | The version of the software package. | [optional] [default to null]
**UpgradePlan** | [**[]SoftwareUpgradePlan**](Software_upgrade_plan.md) | A list of steps that are planned to run during the upgrade in an optimal scenario (i.e., all upgrade checks pass, no step is retried, and the upgrade is not aborted). Steps are listed in the order that they should occur. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

