/*
 * vmysql
 *
 * <br/>https://ncloud.beta-apigw.gov-ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type AddCloudMysqlDatabaseListRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// CloudMysql인스턴스번호
	CloudMysqlInstanceNo *string `json:"cloudMysqlInstanceNo"`

	// CloudMysqlDatabase리스트
	CloudMysqlDatabaseNameList []*string `json:"cloudMysqlDatabaseNameList"`
}
