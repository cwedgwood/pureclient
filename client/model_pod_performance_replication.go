/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type PodPerformanceReplication struct {
	ContinuousBytesPerSec *PodPerformanceReplicationContinuousBytesPerSec `json:"continuous_bytes_per_sec,omitempty"`
	PeriodicBytesPerSec *PodPerformanceReplicationPeriodicBytesPerSec `json:"periodic_bytes_per_sec,omitempty"`
	Pod *PodPerformanceReplicationPod `json:"pod,omitempty"`
	ResyncBytesPerSec *PodPerformanceReplicationResyncBytesPerSec `json:"resync_bytes_per_sec,omitempty"`
	SyncBytesPerSec *PodPerformanceReplicationSyncBytesPerSec `json:"sync_bytes_per_sec,omitempty"`
	// Sample time in milliseconds since the UNIX epoch.
	Time int64 `json:"time,omitempty"`
	// Total bytes transmitted and received per second for all types of replication.
	TotalBytesPerSec int64 `json:"total_bytes_per_sec,omitempty"`
}
