# Model28HostsBody1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The new name for the resource. | [optional] [default to null]
**AddIqns** | **[]string** | Adds the specified iSCSI Qualified Names (IQNs) to those already associated with the specified host. | [optional] [default to null]
**AddNqns** | **[]string** | Adds the specified NVMe Qualified Names (NQNs) to those already associated with the specified host. | [optional] [default to null]
**AddWwns** | **[]string** | Adds the specified Fibre Channel World Wide Names (WWNs) to those already associated with the specified host. | [optional] [default to null]
**Chap** | [***Api28hostsChap**](api2.8hosts_chap.md) |  | [optional] [default to null]
**HostGroup** | [***Api28hostsHostGroup**](api2.8hosts_host_group.md) |  | [optional] [default to null]
**Iqns** | **[]string** | The iSCSI qualified name (IQN) associated with the host. | [optional] [default to null]
**Nqns** | **[]string** | The NVMe Qualified Name (NQN) associated with the host. | [optional] [default to null]
**Personality** | **string** | Determines how the system tunes the array to ensure that it works optimally with the host. Set &#x60;personality&#x60; to the name of the host operating system or virtual memory system. Valid values are &#x60;aix&#x60;, &#x60;esxi&#x60;, &#x60;hitachi-vsp&#x60;, &#x60;hpux&#x60;, &#x60;oracle-vm-server&#x60;, &#x60;solaris&#x60;, and &#x60;vms&#x60;. If your system is not listed as one of the valid host personalities, do not set the option. By default, the personality is not set. | [optional] [default to null]
**PreferredArrays** | [**[]Api28hostsPreferredArrays**](api2.8hosts_preferred_arrays.md) | For synchronous replication configurations, sets a host&#x27;s preferred array to specify which array exposes active/optimized paths to that host. Enter multiple preferred arrays in comma-separated format. If a preferred array is set for a host, then the other arrays in the same pod will expose active/non-optimized paths to that host. If the host is in a host group, &#x60;preferred_arrays&#x60; cannot be set because host groups have their own preferred arrays. On a preferred array of a certain host, all the paths on all the ports (for both the primary and secondary controllers) are set up as A/O (active/optimized) paths, while on a non-preferred array, all the paths are A/N (Active/Non-optimized) paths. | [optional] [default to null]
**RemoveIqns** | **[]string** | Disassociates the specified iSCSI Qualified Names (IQNs) from the specified host. | [optional] [default to null]
**RemoveNqns** | **[]string** | Disassociates the specified NVMe Qualified Names (NQNs) from the specified host. | [optional] [default to null]
**RemoveWwns** | **[]string** | Disassociates the specified Fibre Channel World Wide Names (WWNs) from the specified host. | [optional] [default to null]
**Wwns** | **[]string** | The Fibre Channel World Wide Name (WWN) associated with the host. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

