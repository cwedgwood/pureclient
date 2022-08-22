# Api28snmpagentsV3

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPassphrase** | **string** | Passphrase used by Purity//FA to authenticate the array with the specified managers. | [optional] [default to null]
**AuthProtocol** | **string** | Hash algorithm used to validate the authentication passphrase. Valid values are &#x60;MD5&#x60; and &#x60;SHA&#x60;. | [optional] [default to null]
**PrivacyPassphrase** | **string** | Passphrase used to encrypt SNMP messages. | [optional] [default to null]
**PrivacyProtocol** | **string** | Encryption protocol for SNMP messages. Valid values are &#x60;AES&#x60; and &#x60;DES&#x60;. | [optional] [default to null]
**User** | **string** | User ID recognized by the specified SNMP managers which Purity//FA is to use in communications with them. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

