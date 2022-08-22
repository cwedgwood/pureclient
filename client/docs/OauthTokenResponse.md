# OauthTokenResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The serialized OAuth 2.0 Bearer token used to perform authenticated requests. The access token must be added to the Authorization header of all API calls. | [optional] [default to null]
**IssuedTokenType** | **string** | The type of token that is issued. The Pure Storage REST API supports OAuth 2.0 access tokens. | [optional] [default to null]
**TokenType** | **string** | Indicates how the API client can use the access token issued. The Pure Storage REST API supports the &#x60;Bearer&#x60; token. | [optional] [default to null]
**ExpiresIn** | **int32** | The duration after which the access token will expire. Measured in seconds. This differs from other duration fields that are expressed in milliseconds. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

