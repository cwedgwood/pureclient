# Api28offloadsNfs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The hostname or IP address of the NFS server to where the data will be offloaded. An array can be connected to one offload target at a time, while multiple arrays can be connected to the same offload target. If the protection group is in a stretched pod, for high availability, connect both arrays in the stretched pod to the offload target. | [optional] [default to null]
**MountOptions** | **string** | The custom mount options on the NFS server. Enter multiple mount options in comma-separated format. Valid values include &#x60;port&#x60;, &#x60;rsize&#x60;, &#x60;wsize&#x60;, &#x60;nfsvers&#x60;, and &#x60;tcp&#x60; or &#x60;udp&#x60;. These mount options are common to all NFS file systems. | [optional] [default to null]
**MountPoint** | **string** | The mount point of the NFS export on the NFS server. For example, &#x60;/mnt&#x60;. The &#x60;access_key_id&#x60;, &#x60;secret_access_key&#x60;, and &#x60;bucket&#x60; parameters must be set together. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

