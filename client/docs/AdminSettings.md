# AdminSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockoutDuration** | **int64** | The lockout duration, in milliseconds, if a user is locked out after reaching the maximum number of login attempts. Ranges from 1 second to 90 days. | [optional] [default to null]
**MaxLoginAttempts** | **int32** | Maximum number of failed login attempts allowed before the user is locked out. | [optional] [default to null]
**MinPasswordLength** | **int32** | Minimum password length. If not specified, defaults to 1. | [optional] [default to null]
**SingleSignOnEnabled** | **bool** | If &#x60;true&#x60;, then single sign-on is enabled for the array. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

