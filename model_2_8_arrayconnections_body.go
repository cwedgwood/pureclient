/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Model28ArrayconnectionsBody struct {
	// Management IP address of the target array.
	ManagementAddress string `json:"management_address,omitempty"`
	// IP addresses and FQDNs of the target arrays. Configurable only when `replication_transport` is set to `ip`. If not configured, will be set to all the replication addresses available on the target array at the time of the POST.
	ReplicationAddresses []string `json:"replication_addresses,omitempty"`
	// The type of replication. Valid values are `async-replication` and `sync-replication`.
	Type_ string `json:"type,omitempty"`
	// The protocol used to transport data betwen the local array and the remote array. Valid values are `ip` and `fc`. The default is `ip`.
	ReplicationTransport string `json:"replication_transport,omitempty"`
	// The connection key of the target array.
	ConnectionKey string `json:"connection_key,omitempty"`
}