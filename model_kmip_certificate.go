/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// The certificate used to verify FlashArray authenticity to the KMIP servers.
type KmipCertificate struct {
	// The resource name, such as volume name, pod name, snapshot name, and so on.
	Name string `json:"name,omitempty"`
}