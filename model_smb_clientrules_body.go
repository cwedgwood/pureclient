/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type SmbClientrulesBody struct {
	// A list of SMB policy rules to create.
	Rules []Api28policiessmbclientrulesRules `json:"rules,omitempty"`
}