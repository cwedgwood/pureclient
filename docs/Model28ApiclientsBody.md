# Model28ApiclientsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxRole** | **string** | The maximum role allowed for ID Tokens issued by this API client. The bearer of an access token will be authorized to perform actions within the intersection of this &#x60;max_role&#x60; and the role of the array user specified as the &#x60;sub&#x60;. Valid values are &#x60;array_admin&#x60;, &#x60;storage_admin&#x60;, &#x60;ops_admin&#x60;, and &#x60;readonly&#x60;. Users with the &#x60;readonly&#x60; (Read Only) role can perform operations that convey the state of the array. Read Only users cannot alter the state of the array. Users with the &#x60;ops_admin&#x60; (Ops Admin) role can perform the same operations as Read Only users plus enable and disable remote assistance sessions. Ops Admin users cannot alter the state of the array. Users with the &#x60;storage_admin&#x60; (Storage Admin) role can perform the same operations as Read Only users plus storage related operations, such as administering volumes, hosts, and host groups. Storage Admin users cannot perform operations that deal with global and system configurations. Users with the &#x60;array_admin&#x60; (Array Admin) role can perform the same operations as Storage Admin users plus array-wide changes dealing with global and system configurations. In other words, Array Admin users can perform all operations. | [optional] [default to null]
**Issuer** | **string** | The name of the identity provider that will be issuing ID Tokens for this API client. The &#x60;iss&#x60; claim in the JWT issued must match this string. If not specified, defaults to the API client name. | [optional] [default to null]
**PublicKey** | **string** | The API client&#x27;s PEM formatted (Base64 encoded) RSA public key. Include the &#x60;-----BEGIN PUBLIC KEY-----&#x60; and &#x60;-----END PUBLIC KEY-----&#x60; lines. | [optional] [default to null]
**AccessTokenTtlInMs** | **int64** | The TTL (Time To Live) length of time for the exchanged access token. Measured in milliseconds. If not specified, defaults to &#x60;86400000&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

