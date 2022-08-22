# NetworkInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Enabled** | **bool** | Returns a value of &#x60;true&#x60; if the specified network interface or Fibre Channel port is enabled. Returns a value of &#x60;false&#x60; if the specified network interface or Fibre Channel port is disabled. | [optional] [default to null]
**Eth** | [***NetworkInterfaceEth**](NetworkInterface_eth.md) |  | [optional] [default to null]
**Fc** | [***NetworkInterfaceFc**](NetworkInterface_fc.md) |  | [optional] [default to null]
**InterfaceType** | **string** | The interface type. Valid values are &#x60;eth&#x60; and &#x60;fc&#x60;. | [optional] [default to null]
**Services** | **[]string** | The services provided by the specified network interface or Fibre Channel port. | [optional] [default to null]
**Speed** | **int64** | Configured speed of the specified network interface or Fibre Channel port (in Gb/s). Typically this is the maximum speed of the port or bond represented by the network interface. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

