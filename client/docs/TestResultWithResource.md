# TestResultWithResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentAddress** | **string** | Address of the component running the test. | [optional] [default to null]
**ComponentName** | **string** | Name of the component running the test. | [optional] [default to null]
**Description** | **string** | What the test is doing. | [optional] [default to null]
**Destination** | **string** | The URI of the target server being tested. | [optional] [default to null]
**Enabled** | **bool** | Whether the object being tested is enabled or not. Returns a value of &#x60;true&#x60; if the the service is enabled. Returns a value of &#x60;false&#x60; if the service is disabled. | [optional] [default to null]
**ResultDetails** | **string** | Additional information about the test result. | [optional] [default to null]
**Success** | **bool** | Whether the object being tested passed the test or not. Returns a value of &#x60;true&#x60; if the specified test has succeeded. Returns a value of &#x60;false&#x60; if the specified test has failed. | [optional] [default to null]
**TestType** | **string** | Displays the type of test being performed. The returned values are determined by the &#x60;resource&#x60; being tested and its configuration. Values include &#x60;array-admin-group-searching&#x60;, &#x60;binding&#x60;, &#x60;connecting&#x60;, &#x60;phonehome&#x60;, &#x60;phonehome-ping&#x60;, &#x60;remote-assist&#x60;, &#x60;rootdse-searching&#x60;, &#x60;read-only-group-searching&#x60;, &#x60;storage-admin-group-searching&#x60;, and &#x60;validate-ntp-configuration&#x60;. | [optional] [default to null]
**Resource** | [***TestResultWithResourceResource**](TestResultWithResource_resource.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

