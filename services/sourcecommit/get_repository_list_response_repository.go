/*
 * api
 *
 * <br/>https://sourcecommit.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcecommit

type GetRepositoryListResponseRepository struct {

	// repository id
	Id *int `json:"id,omitempty"`

	// repository name
	Name *string `json:"name"`

	// repository permission
	Permission *string `json:"permission,omitempty"`

	// repository detail action name
	ActionName *string `json:"actionName,omitempty"`
}