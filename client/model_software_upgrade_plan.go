/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type SoftwareUpgradePlan struct {
	// Name of the upgrade step.
	StepName string `json:"step_name,omitempty"`
	// Description of the upgrade step.
	Description string `json:"description,omitempty"`
	// The version to which the step is upgrading.
	HopVersion string `json:"hop_version,omitempty"`
}
