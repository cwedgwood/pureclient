# Model28ActivedirectoryBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputerName** | **string** | The name of the computer account to be created in the Active Directory domain. If not specified, defaults to the name of the Active Directory configuration. | [optional] [default to null]
**DirectoryServers** | **[]string** | A list of directory servers used for lookups related to user authorization. Servers must be specified in FQDN format. All specified servers must be registered to the domain appropriately in the configured DNS of the array and are only communicated with over the secure LDAP (LDAPS) protocol. If not specified, servers are resolved for the domain in DNS. | [optional] [default to null]
**Domain** | **string** | The Active Directory domain to join. | [optional] [default to null]
**KerberosServers** | **[]string** | A list of key distribution servers to use for Kerberos protocol. Servers must be specified in FQDN format. All specified servers must be registered to the domain appropriately in the configured DNS of the array. If not specified, servers are resolved for the domain in DNS. | [optional] [default to null]
**Password** | **string** | The login password of the user with privileges to create the computer account in the domain. This is not persisted on the array. | [optional] [default to null]
**User** | **string** | The login name of the user with privileges to create the computer account in the domain. This is not persisted on the array. | [optional] [default to null]
**JoinOu** | **string** | The distinguished name of the organizational unit in which the computer account should be created when joining the domain. The &#x60;DC&#x3D;...&#x60; components of the distinguished name can be optionally omitted. If not specified, defaults to &#x60;CN&#x3D;Computers&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

