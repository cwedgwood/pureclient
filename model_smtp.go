/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Smtp struct {
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// Password for the relay host, if needed.
	Password string `json:"password,omitempty"`
	// Relay server used as a forwarding point for email sent from the array. Can be set as a hostname, IPv4 address, or IPv6 address, with optional port numbers. The expected format for IPv4 is `ddd.ddd.ddd.ddd:PORT`. The expected format for IPv6 is `xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx` or, if a port number is specified, `[xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx]:PORT`.
	RelayHost string `json:"relay_host,omitempty"`
	// Domain name appended to alert email messages.
	SenderDomain string `json:"sender_domain,omitempty"`
	// User name for the relay host, if needed.
	UserName string `json:"user_name,omitempty"`
}