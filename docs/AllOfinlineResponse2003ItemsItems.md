# AllOfinlineResponse2003ItemsItems

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**ComputerName** | **string** | The name of the computer account in the Active Directory domain. | [optional] [default to null]
**DirectoryServers** | **[]string** | A list of directory servers used for lookups related to user authorization. Servers must be specified in FQDN format. All specified servers must be registered to the domain appropriately in the configured DNS of the array and are only communicated with over the secure LDAP (LDAPS) protocol. If this field is &#x60;null&#x60;, the servers are resolved for the domain in DNS. | [optional] [default to null]
**Domain** | **string** | The Active Directory domain joined. | [optional] [default to null]
**KerberosServers** | **[]string** | A list of key distribution servers to use for Kerberos protocol. Servers must be specified in FQDN format. All specified servers must be registered to the domain appropriately in the configured DNS of the array. If this field is &#x60;null&#x60;, the servers are resolved for the domain in DNS. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

