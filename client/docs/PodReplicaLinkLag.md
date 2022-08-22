# PodReplicaLinkLag

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A non-modifiable, globally unique ID chosen by the system. | [optional] [default to null]
**Direction** | **string** | The direction of replication. Valid values include &#x60;inbound&#x60; and &#x60;outbound&#x60;. | [optional] [default to null]
**Lag** | [***PodReplicaLinkLagLag**](PodReplicaLinkLag_lag.md) |  | [optional] [default to null]
**LocalPod** | [***PodReplicaLinkLocalPod**](PodReplicaLink_local_pod.md) |  | [optional] [default to null]
**RecoveryPoint** | **int64** | Time when the last piece of data was replicated, in milliseconds since the UNIX epoch, and the recovery point if the pod is promoted. If the pod is baselining then the value is &#x60;null&#x60;. | [optional] [default to null]
**RemotePod** | [***PodReplicaLinkRemotePod**](PodReplicaLink_remote_pod.md) |  | [optional] [default to null]
**Remotes** | [**[]AlertEventAlert**](AlertEvent_alert.md) | Reference to a remote array. | [optional] [default to null]
**Status** | **string** | Status of the replica-link. Valid values include &#x60;replicating&#x60;, &#x60;baselining&#x60;, &#x60;paused&#x60;, &#x60;quiescing&#x60;, &#x60;quiesced&#x60;, &#x60;idle&#x60;, and &#x60;unhealthy&#x60;. | [optional] [default to null]
**Time** | **int64** | Sample time in milliseconds since the UNIX epoch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

