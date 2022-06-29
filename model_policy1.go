/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Policy1 struct {
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
	// Returns a value of `true` if the policy is enabled.
	Enabled bool `json:"enabled,omitempty"`
}