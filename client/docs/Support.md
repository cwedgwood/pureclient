# Support

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhonehomeEnabled** | **bool** | The status of phonehome. If set to &#x60;true&#x60;, enables phonehome. If set to &#x60;false&#x60;, disables phonehome. | [optional] [default to null]
**Proxy** | **string** | The value of the current proxy, which is used to connect to cloud services such as phonehome and remote assist. Specify the server name, including the scheme and proxy port number. | [optional] [default to null]
**RemoteAssistActive** | **bool** | The status of the remote assist session. If set to &#x60;true&#x60;, enables the remote assist session. If set to &#x60;false&#x60;, disables the remote assist session. | [optional] [default to null]
**RemoteAssistOpened** | **int64** | The timestamp when the session opened, measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**RemoteAssistExpires** | **int64** | The timestamp when the session expires, measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**RemoteAssistStatus** | **string** | The status of the remote assist session. Values include &#x60;connected&#x60;, &#x60;connecting&#x60;, &#x60;disconnected&#x60;, and &#x60;session-active&#x60;. | [optional] [default to null]
**RemoteAssistPaths** | [**[]SupportRemoteAssistPaths**](Support_remote_assist_paths.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

