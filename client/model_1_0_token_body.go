/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Model10TokenBody struct {
	// The method by which the access token will be obtained. The Pure Storage REST API supports the OAuth 2.0 \"token exchange\" grant type, which indicates that a token exchange is being performed. Set `grant_type` to `urn:ietf:params:oauth:grant-type:token-exchange`.
	GrantType string `json:"grant_type"`
	// An encoded security ID Token representing the identity of the party on behalf of whom the request is being made. The token must be issued by a trusted identity provider which must be either a registered application in Pure1 or an enabled API client on the array. The token must be a JSON Web Token and must contain the following claims: > | JWT claim | Location | API Client Field | Description | Required By | > |-|-|-|-|-| > | kid | Header | key_id | Key ID of the API client that issues the identity token. | FlashArray and FlashBlade only. | > | aud | Payload | id | Client ID of the API client that issues the identity token. | FlashArray and FlashBlade only. | > | sub | Payload | | Login name of the array user for whom the token should be issued. This must be a valid user in the system. | FlashArray and FlashBlade only. | > | iss | Payload | issuer | Application ID for the Pure1 or API client's trusted identity issuer on the array. | All products. | > | iat | Payload | | Timestamp of when the identity token was issued. Measured in milliseconds since the UNIX epoch. | All products. | > | exp | Payload | | Timestamp of when the identity token will expire. Measured in milliseconds since the UNIX epoch. | All products. |  Each token must also be signed with the private key that is paired with the API client's public key.
	SubjectToken string `json:"subject_token"`
	// An identifier that indicates the type of security token specifed in the `subject_token` parameter. The Pure Storage REST API supports the JSON Web Token (JWT) as the means for requesting the access token. Set `subject_token_type` to `urn:ietf:params:oauth:token-type:jwt`.
	SubjectTokenType string `json:"subject_token_type"`
}
