/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// The configuration of eradication feature.
type ArrayEradicationConfig struct {
	// The eradication delay in milliseconds. Automatically eradicate destroyed items after the delay time. Valid values are `86400000` and any multiple of `86400000` in the range of `86400000` and `2592000000`. Any other values will be rounded down to the nearest multiple of `86400000`.
	EradicationDelay int64 `json:"eradication_delay,omitempty"`
}
