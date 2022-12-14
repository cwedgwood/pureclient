/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Model28VolumegroupsBody struct {
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
	// Has this volume group been destroyed? To destroy a volume group, patch to `true`. To recover a destroyed volume group, patch to `false`. If not specified, defaults to `false`.
	Destroyed bool `json:"destroyed,omitempty"`
	Qos *VolumeGroupQos `json:"qos,omitempty"`
	Space *OffloadSpace `json:"space,omitempty"`
	// The amount of time left until the destroyed volume group is permanently eradicated. Measured in milliseconds. Before the `time_remaining` period has elapsed, the destroyed volume group can be recovered by setting `destroyed=false`.
	TimeRemaining int64 `json:"time_remaining,omitempty"`
	// The number of volumes in the volume group.
	VolumeCount int64 `json:"volume_count,omitempty"`
}
