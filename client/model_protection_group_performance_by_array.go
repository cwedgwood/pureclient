/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type ProtectionGroupPerformanceByArray struct {
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// The total number of bytes of replication data transmitted and received per second.
	BytesPerSec int64 `json:"bytes_per_sec,omitempty"`
	// The source array from where the data is replicated.
	Source string `json:"source,omitempty"`
	// The target to where the data is replicated.
	Target string `json:"target,omitempty"`
	// The time when the sample performance data was taken. Measured in milliseconds since the UNIX epoch.
	Time int64 `json:"time,omitempty"`
}
