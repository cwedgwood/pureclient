# ApiClient

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the associated API client. The ID represents the JWT &#x60;aud&#x60; (audience) claim in ID Tokens issued for this API client. | [optional] [default to null]
**Name** | **string** | The API client name. | [optional] [default to null]
**MaxRole** | **string** | The maximum role allowed for ID Tokens issued by this API client. The bearer of an access token will be authorized to perform actions within the intersection of this &#x60;max_role&#x60; and the role of the array user specified as the JWT &#x60;sub&#x60; (subject) claim. Valid &#x60;max_role&#x60; values are &#x60;readonly&#x60;, &#x60;ops_admin&#x60;, &#x60;array_admin&#x60;, and &#x60;storage_admin&#x60;. Users with the &#x60;readonly&#x60; (Read Only) role can perform operations that convey the state of the array. Read Only users cannot alter the state of the array. Users with the &#x60;ops_admin&#x60; (Ops Admin) role can perform the same operations as Read Only users plus enable and disable remote assistance sessions. Ops Admin users cannot alter the state of the array. Users with the &#x60;storage_admin&#x60; (Storage Admin) role can perform the same operations as Read Only users plus storage related operations, such as administering volumes, hosts, and host groups. Storage Admin users cannot perform operations that deal with global and system configurations. Users with the &#x60;array_admin&#x60; (Array Admin) role can perform the same operations as Storage Admin users plus array-wide changes dealing with global and system configurations. In other words, Array Admin users can perform all operations. | [optional] [default to null]
**Issuer** | **string** | The name of the identity provider that will be issuing ID Tokens for this API client. This string represents the JWT &#x60;iss&#x60; (issuer) claim in ID Tokens issued for this API client. | [optional] [default to null]
**PublicKey** | **string** | The API client&#x27;s PEM formatted (Base64 encoded) RSA public key. | [optional] [default to null]
**KeyId** | **string** | The unique identifier for the associated public key of this API client. This string must match the JWT &#x60;kid&#x60; (key ID) claim in ID Tokens issued for this API client. | [optional] [default to null]
**Enabled** | **bool** | If &#x60;true&#x60;, the API client is permitted to exchange ID Tokens for access tokens. API clients are disabled by default. | [optional] [default to null]
**AccessTokenTtlInMs** | **int64** | The requested TTL (Time To Live) length of time for the exchanged access token. Measured in milliseconds. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

