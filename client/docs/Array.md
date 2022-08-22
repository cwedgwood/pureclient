# Array

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Banner** | **string** |  | [optional] [default to null]
**Capacity** | **int64** | The usable capacity in bytes. | [optional] [default to null]
**ConsoleLockEnabled** | **bool** |  | [optional] [default to null]
**Encryption** | [***ArrayEncryption**](Array_encryption.md) |  | [optional] [default to null]
**EradicationConfig** | [***ArrayEradicationConfig**](Array_eradication_config.md) |  | [optional] [default to null]
**IdleTimeout** | **int32** | The idle timeout in milliseconds. Valid values include &#x60;0&#x60; and any multiple of &#x60;60000&#x60; in the range of &#x60;300000&#x60; and &#x60;10800000&#x60;. Any other values will be rounded down to the nearest multiple of &#x60;60000&#x60;. | [optional] [default to null]
**NtpServers** | **[]string** |  | [optional] [default to null]
**Os** | **string** | Specifies the operating system. Valid values are &#x60;Purity&#x60;, &#x60;Purity//FA&#x60;, and &#x60;Purity//FB&#x60;. | [optional] [default to null]
**Parity** | **float32** | A representation of data redundancy on the array. Data redundancy is rebuilt automatically by the system whenever parity is less than &#x60;1.0&#x60;. | [optional] [default to null]
**ScsiTimeout** | **int32** | The SCSI timeout. If not specified, defaults to &#x60;60s&#x60;. | [optional] [default to null]
**Space** | [***ArraySpace**](ArraySpace.md) |  | [optional] [default to null]
**Version** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

