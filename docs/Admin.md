# Admin

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-specified name. The name must be locally unique and cannot be changed. | [optional] [default to null]
**ApiToken** | [***AdminApiToken**](Admin_api_token.md) |  | [optional] [default to null]
**IsLocal** | **bool** | Returns a value of &#x60;true&#x60; if the user is local to the machine, otherwise &#x60;false&#x60;. | [optional] [default to null]
**Locked** | **bool** | Returns a value of &#x60;true&#x60; if the user is currently locked out, otherwise &#x60;false&#x60;. Can be patched to &#x60;false&#x60; to unlock a user. This field is only visible to &#x60;array_admin&#x60; roles. For all other users, the value is always &#x60;null&#x60;. | [optional] [default to null]
**LockoutRemaining** | **int64** | The remaining lockout period, in milliseconds, if the user is locked out. This field is only visible to &#x60;array_admin&#x60; roles. For all other users, the value is always &#x60;null&#x60;. | [optional] [default to null]
**Password** | **string** | Password associated with the account. | [optional] [default to null]
**PublicKey** | **string** | Public key for SSH access. | [optional] [default to null]
**Role** | [***Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

