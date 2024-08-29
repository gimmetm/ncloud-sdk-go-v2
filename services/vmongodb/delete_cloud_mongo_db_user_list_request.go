/*
 * vmongodb
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmongodb/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmongodb

type DeleteCloudMongoDbUserListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// Cloud DB for MongoDB 인스턴스 번호
CloudMongoDbInstanceNo *string `json:"cloudMongoDbInstanceNo"`

	// Cloud DB for MongoDB User 리스트
CloudMongoDbUserList []*DeleteCloudMongoDbUserParameter `json:"cloudMongoDbUserList"`
}
