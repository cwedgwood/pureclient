/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Api28arrayseulaSignature struct {
	// Name of the person who accepted the End User Agreement.
	Name string `json:"name,omitempty"`
	// Title of the person who accepted the End User Agreement.
	Title string `json:"title,omitempty"`
	// Company of the person who accepted the End User Agreement.
	Company string `json:"company,omitempty"`
	// Accepted time in milliseconds since the UNIX epoch.
	Accepted int64 `json:"accepted,omitempty"`
}
