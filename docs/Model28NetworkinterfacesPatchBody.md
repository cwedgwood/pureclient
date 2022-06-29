# Model28NetworkinterfacesPatchBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Returns a value of &#x60;true&#x60; if the specified network interface or Fibre Channel port is enabled. Returns a value of &#x60;false&#x60; if the specified network interface or Fibre Channel port is disabled. | [optional] [default to null]
**Eth** | [***Api28networkinterfacesEth1**](api2.8networkinterfaces_eth_1.md) |  | [optional] [default to null]
**OverrideNpivCheck** | **bool** | N-Port ID Virtualization (NPIV) requires a balanced configuration of Fibre Channel ports configured for SCSI on both controllers. Enabling or Disabling a Fibre Channel port configured for SCSI might cause the NPIV status to change from enabled to disabled or vice versa. Set this option to proceed with enabling or disabling the port. | [optional] [default to null]
**Services** | **[]string** | The services provided by the specified network interface or Fibre Channel port. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

