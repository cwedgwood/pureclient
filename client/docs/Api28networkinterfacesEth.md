# Api28networkinterfacesEth

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The IPv4 or IPv6 address to be associated with the specified network interface. | [optional] [default to null]
**Subinterfaces** | [**[]Api28networkinterfacesEthSubinterfaces**](api2.8networkinterfaces_eth_subinterfaces.md) | List of network interfaces configured to be a subinterface of the specified network interface. | [optional] [default to null]
**Subnet** | [***Api28networkinterfacesEthSubnet**](api2.8networkinterfaces_eth_subnet.md) |  | [optional] [default to null]
**Subtype** | **string** | The subtype of the specified network interface. Only interfaces of subtype &#x60;virtual&#x60; and &#x60;lacp_bond&#x60; can be created. Configurable on POST only. Valid values are &#x60;failover_bond&#x60;, &#x60;lacp_bond&#x60;, &#x60;physical&#x60;, and &#x60;virtual&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

