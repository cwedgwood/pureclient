/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// The quota policy that is limiting usage on this managed directory. This policy defines the total amount of space provisioned to this managed directory and its descendants. The returned value contains two parts&#58; the name of the policy and the managed directory to which the policy is attached.
type DirectoryLimitedBy struct {
	// Reference to the resource to which the effective quota policy is attached.
	Member *AllOfDirectoryLimitedByMember `json:"member,omitempty"`
	// Reference to the effective quota policy.
	Policy *AllOfDirectoryLimitedByPolicy `json:"policy,omitempty"`
}
