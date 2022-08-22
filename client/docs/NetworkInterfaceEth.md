# NetworkInterfaceEth

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The IPv4 or IPv6 address to be associated with the specified network interface. | [optional] [default to null]
**Gateway** | **string** | The IPv4 or IPv6 address of the gateway through which the specified network interface is to communicate with the network. | [optional] [default to null]
**MacAddress** | **string** | The media access control address associated with the specified network interface. | [optional] [default to null]
**Mtu** | **int32** | Maximum message transfer unit (packet) size for the network interface, in bytes. MTU setting cannot exceed the MTU of the corresponding physical interface. | [optional] [default to null]
**Netmask** | **string** | Netmask of the specified network interface that, when combined with the address of the interface, determines the network address of the interface. | [optional] [default to null]
**Subinterfaces** | [**[]Api28networkinterfacesEthSubinterfaces**](api2.8networkinterfaces_eth_subinterfaces.md) | List of network interfaces configured to be a subinterface of the specified network interface. | [optional] [default to null]
**Subnet** | [***Api28networkinterfacesEthSubnet**](api2.8networkinterfaces_eth_subnet.md) |  | [optional] [default to null]
**Subtype** | **string** | The subtype of the specified network interface. Only interfaces of subtype &#x60;virtual&#x60; can be created. Configurable on POST only. Valid values are &#x60;failover_bond&#x60;, &#x60;lacp_bond&#x60;, &#x60;physical&#x60;, and &#x60;virtual&#x60;. | [optional] [default to null]
**Vlan** | **int32** | VLAN ID | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

