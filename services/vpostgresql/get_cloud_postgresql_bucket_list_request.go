/*
 * vpostgresql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vpostgresql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpostgresql

type GetCloudPostgresqlBucketListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// CloudPostgresql서버인스턴스번호
CloudPostgresqlServerInstanceNo *string `json:"cloudPostgresqlServerInstanceNo,omitempty"`
}
