# PodReplicaLinkPerformance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A non-modifiable, globally unique ID chosen by the system. | [optional] [default to null]
**BytesPerSecFromRemote** | **int64** | The number of bytes received per second from a remote array. | [optional] [default to null]
**BytesPerSecToRemote** | **int64** | The number of bytes transmitted per second to a remote array. | [optional] [default to null]
**BytesPerSecTotal** | **int64** | Total bytes transmitted and received per second. | [optional] [default to null]
**Direction** | **string** | The direction of replication. Valid values are &#x60;inbound&#x60; and &#x60;outbound&#x60;. | [optional] [default to null]
**LocalPod** | [***PodReplicaLinkLocalPod**](PodReplicaLink_local_pod.md) |  | [optional] [default to null]
**RemotePod** | [***PodReplicaLinkRemotePod**](PodReplicaLink_remote_pod.md) |  | [optional] [default to null]
**Remotes** | [**[]AlertEventAlert**](AlertEvent_alert.md) | Reference to a remote array. | [optional] [default to null]
**Time** | **int64** | Sample time in milliseconds since the UNIX epoch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

