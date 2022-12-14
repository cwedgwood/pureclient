/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type SyslogServer struct {
	// A user-specified name. The name must be locally unique and cannot be changed.
	Name string `json:"name,omitempty"`
	// The URI of the syslog server in the format `PROTOCOL://HOSTNAME:PORT`.
	Uri string `json:"uri,omitempty"`
}
