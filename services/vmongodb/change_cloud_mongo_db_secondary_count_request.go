/*
 * vmongodb
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmongodb/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmongodb

type ChangeCloudMongoDbSecondaryCountRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// CloudMongoDb 인스턴스번호
CloudMongoDbInstanceNo *string `json:"cloudMongoDbInstanceNo"`

	// Cloud DB for MongoDB config 서버 대수
MemberServerCount *string `json:"memberServerCount"`

	// Cloud DB for MongoDB arbiter 서버 대수
ArbiterServerCount *string `json:"arbiterServerCount"`
}