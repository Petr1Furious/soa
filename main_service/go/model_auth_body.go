/*
 * Main Service API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AuthBody struct {
	Login string `json:"login,omitempty"`

	Password string `json:"password,omitempty"`
}
