/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// Duration in milliseconds that indicates how far behind the replication target is from the source.
type PodReplicaLinkLagLag struct {
	// The average time difference between the current time and `recovery_point` for a period of time.
	Avg int64 `json:"avg,omitempty"`
	// The maximum time difference between the current time and `recovery_point` for a period of time.
	Maximum int64 `json:"maximum,omitempty"`
}