# Smtp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Password** | **string** | Password for the relay host, if needed. | [optional] [default to null]
**RelayHost** | **string** | Relay server used as a forwarding point for email sent from the array. Can be set as a hostname, IPv4 address, or IPv6 address, with optional port numbers. The expected format for IPv4 is &#x60;ddd.ddd.ddd.ddd:PORT&#x60;. The expected format for IPv6 is &#x60;xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx&#x60; or, if a port number is specified, &#x60;[xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx]:PORT&#x60;. | [optional] [default to null]
**SenderDomain** | **string** | Domain name appended to alert email messages. | [optional] [default to null]
**UserName** | **string** | User name for the relay host, if needed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

