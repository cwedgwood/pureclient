/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type PolicyRuleQuota struct {
	// Defines whether the quota rule is enforced or unenforced. If the quota rule is enforced and logical space usage exceeds the quota limit, any modification operations that result in a need for more space are blocked. If the quota rule is unenforced and logical space usage exceeds the quota limit, notification emails are sent to targets that are specified using the `notification` parameter. No client operations are blocked when an unenforced limit is exceeded. If set to `true`, the limit is enforced. If set to `false`, notification targets are informed when the usage exceeds 80 percent of the limit.
	Enforced bool `json:"enforced,omitempty"`
	// Name of this rule. The name is automatically generated by the system.
	Name string `json:"name,omitempty"`
	// Targets to notify when usage approaches the quota limit. The list of notification targets is a comma-separated string. Valid values are `user`, and `group`. If not specified, notification targets are not assigned for the rule.
	Notifications string `json:"notifications,omitempty"`
	// The policy to which this rule belongs.
	Policy *AllOfPolicyRuleQuotaPolicy `json:"policy,omitempty"`
	// Logical space limit of the quota assigned by the rule, measured in bytes. This value cannot be set to 0.
	QuotaLimit int64 `json:"quota_limit,omitempty"`
}
