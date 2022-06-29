# SnmpManagers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Host** | **string** | DNS hostname or IP address of a computer that hosts an SNMP manager to which Purity//FA is to send trap messages when it generates alerts. | [optional] [default to null]
**Notification** | **string** | The type of notification the agent will send. Valid values are &#x60;inform&#x60; and &#x60;trap&#x60;. | [optional] [default to null]
**V2c** | [***Api28snmpagentsV2c**](api2.8snmpagents_v2c.md) |  | [optional] [default to null]
**V3** | [***SnmpAgentV3**](SNMPAgent_v3.md) |  | [optional] [default to null]
**Version** | **string** | Version of the SNMP protocol to be used by Purity//FA to communicate with the specified manager. Valid values are &#x60;v2c&#x60; and &#x60;v3&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

