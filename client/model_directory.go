/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Directory struct {
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
	// The managed directory creation time, measured in milliseconds since the UNIX epoch.
	Created int64 `json:"created,omitempty"`
	// Returns a value of `true` if the managed directory has been destroyed and is pending eradication. The `time_remaining` value displays the amount of time left until the destroyed managed directory is permanently eradicated. Once the `time_remaining` period has elapsed, the managed directory is permanently eradicated and can no longer be recovered.
	Destroyed bool `json:"destroyed,omitempty"`
	// The managed directory name without the file system name prefix. A full managed directory name is constructed in the form of `FILE_SYSTEM:DIR` where `FILE_SYSTEM` is the file system name and `DIR` is the value of this field.
	DirectoryName string `json:"directory_name,omitempty"`
	// The file system that this managed directory is in.
	FileSystem *interface{} `json:"file_system,omitempty"`
	// Absolute path of the managed directory in the file system.
	Path string `json:"path,omitempty"`
	// Displays size and space consumption information.
	Space *interface{} `json:"space,omitempty"`
	// The amount of time left, measured in milliseconds until the destroyed managed directory is permanently eradicated.
	TimeRemaining int64 `json:"time_remaining,omitempty"`
	LimitedBy *DirectoryLimitedBy `json:"limited_by,omitempty"`
}
