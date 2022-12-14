/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type ProtectionGroupTarget struct {
	ProtectionGroup *Api28certificatescertificatesigningrequestsCertificate `json:"protection_group,omitempty"`
	Target *Api28certificatescertificatesigningrequestsCertificate `json:"target,omitempty"`
	// If set to `true`, the target array has allowed the source array to replicate protection group data to the target array. If set to `false`, the target array has not allowed the source array to replicate protection group data to the target.
	Allowed bool `json:"allowed,omitempty"`
}
