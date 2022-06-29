/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Controllers struct {
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// Mode of controller. Values include `not present`, `offline`, `primary`, and `secondary`.
	Mode string `json:"mode,omitempty"`
	Model string `json:"model,omitempty"`
	// Status of controller. Values include `not ready`, `ready`, `unknown`, and `updating`.
	Status string `json:"status,omitempty"`
	// Type of controller. Values include `array_controller` and `shelf_controller`.
	Type_ string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
}