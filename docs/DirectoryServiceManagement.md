# DirectoryServiceManagement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserLoginAttribute** | **string** | User login attribute in the structure of the configured LDAP servers. Typically the attribute field that holds the user&#x27;s unique login name. Default value is &#x60;sAMAccountName&#x60; for Active Directory or &#x60;uid&#x60; for all other directory services. | [optional] [default to null]
**UserObjectClass** | **string** | Value of the object class for a management LDAP user. Defaults to &#x60;User&#x60; for Active Directory servers, &#x60;posixAccount&#x60; or &#x60;shadowAccount&#x60; for OpenLDAP servers dependent on the group type of the server, or &#x60;person&#x60; for all other directory servers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

