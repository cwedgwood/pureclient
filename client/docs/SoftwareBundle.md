# SoftwareBundle

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A non-modifiable, globally unique ID chosen by the system. | [optional] [default to null]
**Source** | **string** | The location of the upgrade bundle. | [optional] [default to null]
**Created** | **float64** | Creation time in milliseconds since the UNIX epoch. | [optional] [default to null]
**Details** | **string** | The detailed reason for the &#x60;status&#x60;. | [optional] [default to null]
**DownloadProgress** | **float32** | The progress of the download. Displayed in decimal format. | [optional] [default to null]
**Status** | **string** | The status of the software bundle. Valid values are &#x60;downloading&#x60;, &#x60;failed&#x60;, &#x60;ready&#x60;, and &#x60;verifying&#x60;. A status of &#x60;downloading&#x60; indicates that the array is downloading the upgrade bundle. A status of &#x60;failed&#x60; indicates that the array has failed to download the upgrade bundle. A status of &#x60;ready&#x60; indicates that the upgrade bundle is ready to be installed. A status of &#x60;verifying&#x60; indicates that the array is verifying the upgrade bundle. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

