/*
 * vmongodb
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmongodb/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmongodb

type ChangeCloudMongoDbShardCountRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// CloudMongoDb 인스턴스번호
CloudMongoDbInstanceNo *string `json:"cloudMongoDbInstanceNo"`

	// Cloud DB for MongoDB 샤드 대수
ShardCount *string `json:"shardCount"`
}
