# Session

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**EndTime** | **int64** | Date and time the user logged out of the Purity//FA interface. Not set if the session is still active. | [optional] [default to null]
**Event** | **string** | Description of session events such as login and user session. Valid values include &#x60;failed authentication&#x60;, &#x60;user session&#x60;, &#x60;login&#x60;, &#x60;logout&#x60;, &#x60;API token obtained&#x60;, and &#x60;request without session&#x60;. | [optional] [default to null]
**EventCount** | **int32** | Number of session events. | [optional] [default to null]
**Location** | **string** | IP address of the user client connecting to the array or console if connected through local console. | [optional] [default to null]
**Method** | **string** | Method by which the user attempted to log in. Valid values include &#x60;API token&#x60;, &#x60;JWT&#x60;, &#x60;password&#x60;, and &#x60;public key&#x60;. | [optional] [default to null]
**StartTime** | **int64** | Date and time the user logged in to the Purity//FA interface. | [optional] [default to null]
**User** | **string** | Username of the Purity//FA user who triggered the user session event. | [optional] [default to null]
**UserInterface** | **string** | The user interface through which the user session event was performed. Valid values include &#x60;CLI&#x60;, &#x60;GUI&#x60;, and &#x60;REST&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

