/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Model28OffloadsBody struct {
	Azure *Api28offloadsAzure `json:"azure,omitempty"`
	GoogleCloud *Api28offloadsGooglecloud `json:"google-cloud,omitempty"`
	Nfs *Api28offloadsNfs `json:"nfs,omitempty"`
	S3 *Api28offloadsS3 `json:"s3,omitempty"`
}
