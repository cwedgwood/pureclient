/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type ProtectionGroup struct {
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
	// Has this protection group been destroyed? To destroy a protection group, patch to `true`. To recover a destroyed protection group, patch to `false`. If not specified, defaults to `false`.
	Destroyed bool `json:"destroyed,omitempty"`
	// Number of hosts in this protection group.
	HostCount int64 `json:"host_count,omitempty"`
	// Number of host groups in this protection group.
	HostGroupCount int64 `json:"host_group_count,omitempty"`
	// If set to `true`, the protection group belongs to the local array. If set to `false`, the protection group belongs to the remote array.
	IsLocal bool `json:"is_local,omitempty"`
	Pod *ProtectionGroupPod `json:"pod,omitempty"`
	//REDDY: commented below
	// The schedule settings for asynchronous replication.
	//ReplicationSchedule *Replic `json:"replication_schedule,omitempty"`
	SnapshotSchedule *ProtectionGroupSnapshotSchedule `json:"snapshot_schedule,omitempty"`
	Source *ProtectionGroupSource `json:"source,omitempty"`
	SourceRetention *ProtectionGroupSourceRetention `json:"source_retention,omitempty"`
	Space *ProtectionGroupSpace `json:"space,omitempty"`
	// The number of targets to where this protection group replicates.
	TargetCount int64 `json:"target_count,omitempty"`
	TargetRetention *ProtectionGroupSourceRetention `json:"target_retention,omitempty"`
	// The amount of time left until the destroyed protection group is permanently eradicated. Measured in milliseconds. Before the `time_remaining` period has elapsed, the destroyed protection group can be recovered by setting `destroyed=false`.
	TimeRemaining int64 `json:"time_remaining,omitempty"`
	// The number of volumes in the protection group.
	VolumeCount int64 `json:"volume_count,omitempty"`
}
