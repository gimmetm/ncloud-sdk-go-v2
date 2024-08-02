/*
 * vpostgresql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vpostgresql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpostgresql

type CloudPostgresqlUserParameter struct {

	// DBUserID
Name *string `json:"name"`

	// clientCidr
ClientCidr *string `json:"clientCidr"`

	// DBUserPassword
Password *string `json:"password"`

	// ReplicationRole여부
IsReplicationRole *bool `json:"isReplicationRole"`
}
