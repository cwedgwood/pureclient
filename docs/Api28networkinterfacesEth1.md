# Api28networkinterfacesEth1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddSubinterfaces** | [**[]Api28networkinterfacesEthSubinterfaces**](api2.8networkinterfaces_eth_subinterfaces.md) | Child devices to be added to the specified bond interface. | [optional] [default to null]
**Address** | **string** | The IPv4 or IPv6 address to be associated with the specified network interface. | [optional] [default to null]
**Gateway** | **string** | The IPv4 or IPv6 address of the gateway through which the specified network interface is to communicate with the network. | [optional] [default to null]
**Mtu** | **int32** | Maximum message transfer unit (packet) size for the network interface in bytes. MTU setting cannot exceed the MTU of the corresponding physical interface. | [optional] [default to null]
**Netmask** | **string** | Netmask of the specified network interface that, when combined with the address of the interface, determines the network address of the interface. | [optional] [default to null]
**RemoveSubinterfaces** | [**[]Api28networkinterfacesEthSubinterfaces**](api2.8networkinterfaces_eth_subinterfaces.md) | Child devices to be removed from the specified bond interface. | [optional] [default to null]
**Subinterfaces** | [**[]Api28networkinterfacesEthSubinterfaces**](api2.8networkinterfaces_eth_subinterfaces.md) | Child devices to be added to the specified bond interface. | [optional] [default to null]
**Subnet** | [***Api28networkinterfacesEthSubnet**](api2.8networkinterfaces_eth_subnet.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

