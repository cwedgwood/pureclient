/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Certificate struct {
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// The text of the certificate.
	Certificate string `json:"certificate,omitempty"`
	// The common name field listed in the certificate.
	CommonName string `json:"common_name,omitempty"`
	// Two-letter country (ISO) code listed in the certificate.
	Country string `json:"country,omitempty"`
	// The email field listed in the certificate.
	Email string `json:"email,omitempty"`
	// The text of the intermediate certificate chains.
	IntermediateCertificate string `json:"intermediate_certificate,omitempty"`
	// The party that issued the certificate.
	IssuedBy string `json:"issued_by,omitempty"`
	// The party to whom the certificate is issued.
	IssuedTo string `json:"issued_to,omitempty"`
	// The size of the private key for the certificate in bits. Default is 2048 bits.
	KeySize int32 `json:"key_size,omitempty"`
	// The locality field listed in the certificate.
	Locality string `json:"locality,omitempty"`
	// The organization field listed in the certificate.
	Organization string `json:"organization,omitempty"`
	// The organizational unit field listed in the certificate.
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	// The state/province field listed in the certificate.
	State string `json:"state,omitempty"`
	// The type of certificate. Valid values are `self-signed` and `imported`.
	Status string `json:"status,omitempty"`
	// The date when the certificate starts being valid.
	ValidFrom int64 `json:"valid_from,omitempty"`
	// The date of when the certificate stops being valid.
	ValidTo int64 `json:"valid_to,omitempty"`
}