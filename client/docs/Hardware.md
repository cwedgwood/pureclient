# Hardware

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Details** | **string** | Details about the status of the component if not healthy. | [optional] [default to null]
**IdentifyEnabled** | **bool** | If &#x60;true&#x60;, the ID LED is lit to visually identify the component. | [optional] [default to null]
**Index** | **int32** | Number that identifies the relative position of a hardware component within the array. | [optional] [default to null]
**Model** | **string** | Model number of the hardware component. | [optional] [default to null]
**Serial** | **string** | Serial number of the hardware component. | [optional] [default to null]
**Slot** | **int32** | Slot number occupied by the PCI Express card that hosts the component. | [optional] [default to null]
**Speed** | **int64** | Speed (in bytes per second) at which the component is operating. | [optional] [default to null]
**Status** | **string** | Component status. Values include &#x60;critical&#x60;, &#x60;healthy&#x60;, &#x60;identifying&#x60;, &#x60;unhealthy&#x60;, &#x60;unknown&#x60;, and &#x60;unused&#x60;. | [optional] [default to null]
**Temperature** | **int32** | Temperature (in degrees Celsius) reported by the component. | [optional] [default to null]
**Type_** | **string** | Type of hardware component. Values include &#x60;bay&#x60;, &#x60;ct&#x60;, &#x60;ch&#x60;, &#x60;eth&#x60;, &#x60;fan&#x60;, &#x60;fb&#x60;, &#x60;fc&#x60;, &#x60;fm&#x60;, &#x60;ib&#x60;, &#x60;iom&#x60;, &#x60;nvb&#x60;, &#x60;pwr&#x60;, &#x60;sas&#x60;, &#x60;sh&#x60;, and &#x60;tmp&#x60;. | [optional] [default to null]
**Voltage** | **int32** | Voltage (in Volts) reported by the component. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

