/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// The snapshot retention policy.
type ProtectionGroupSourceRetention struct {
	// The length of time to keep the specified snapshots. Measured in seconds.
	AllForSec int32 `json:"all_for_sec,omitempty"`
	// The number of days to keep the snapshots after the `all_for_sec` period has passed.
	Days int32 `json:"days,omitempty"`
	// The number of snapshots to keep per day after the `all_for_sec` period has passed.
	PerDay int32 `json:"per_day,omitempty"`
}
