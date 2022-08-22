# Subnet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Enabled** | **bool** | Returns a value of &#x60;true&#x60; if subnet is enabled. Returns a value of &#x60;false&#x60; if subnet is disabled. If not specified, defaults to &#x60;true&#x60;. | [optional] [default to null]
**Gateway** | **string** | The IPv4 or IPv6 address of the gateway through which the specified subnet is to communicate with the network. | [optional] [default to null]
**Mtu** | **int32** | Maximum message transfer unit (packet) size for the subnet in bytes. MTU setting cannot exceed the MTU of the corresponding physical interface. If not specified, defaults to &#x60;1500&#x60;. | [optional] [default to null]
**Prefix** | **string** | The IPv4 or IPv6 address to be associated with the specified subnet. | [optional] [default to null]
**Vlan** | **int32** | VLAN ID number. | [optional] [default to null]
**Interfaces** | [**[]Api28networkinterfacesEthSubinterfaces**](api2.8networkinterfaces_eth_subinterfaces.md) | List of network interfaces associated with this subnet. | [optional] [default to null]
**Services** | **[]string** | The services provided by this subnet, as inherited from all of its interfaces. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

