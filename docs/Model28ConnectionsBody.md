# Model28ConnectionsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lun** | **int32** | The logical unit number (LUN) by which the specified hosts are to address the specified volume. If the LUN is not specified, the system automatically assigns a LUN to the connection. To automatically assign a LUN to a private connection, the system starts at LUN &#x60;1&#x60; and counts up to the maximum LUN &#x60;4095&#x60;, assigning the first available LUN to the connection. For shared connections, the system starts at LUN &#x60;254&#x60; and counts down to the minimum LUN &#x60;1&#x60;, assigning the first available LUN to the connection. If all LUNs in the &#x60;[1...254]&#x60; range are taken, the system starts at LUN &#x60;255&#x60; and counts up to the maximum LUN &#x60;4095&#x60;, assigning the first available LUN to the connection. | [optional] [default to null]
**ProtocolEndpoint** | [***Api28connectionsProtocolEndpoint**](api2.8connections_protocol_endpoint.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

