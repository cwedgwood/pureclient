/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type DirectoryQuota struct {
	// The directory to which the quota applies.
	Directory *AllOfDirectoryQuotaDirectory `json:"directory,omitempty"`
	// Returns a value of `true` if the policy is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Defines whether the quota rule is enforced or unenforced. If the quota rule is enforced and logical space usage exceeds the quota limit, any modification operations that result in a need for more space are blocked. If the quota rule is unenforced and logical space usage exceeds the quota limit, notification emails are sent to targets that are specified using the `notification` parameter. No client operations are blocked when an unenforced limit is exceeded. If set to `true`, the limit is enforced. If set to `false`, notification targets are informed when the usage exceeds 80 percent of the limit.
	Enforced bool `json:"enforced,omitempty"`
	// Absolute path of the directory in the file system.
	Path string `json:"path,omitempty"`
	// The percentage of the space used in the directory with respect to the quota limit.
	PercentageUsed float32 `json:"percentage_used,omitempty"`
	// The effective quota policy that imposes the limit. This is the policy with the lowest limit.
	Policy *AllOfDirectoryQuotaPolicy `json:"policy,omitempty"`
	// Effective quota limit imposed by the quota policy rule attached to the directory, measured in bytes.
	QuotaLimit int64 `json:"quota_limit,omitempty"`
	// Name of the rule that results in this quota and behavior being applied to this directory.
	RuleName string `json:"rule_name,omitempty"`
	// The amount of logically written data for the directory, measured in bytes.
	Usage int64 `json:"usage,omitempty"`
}
