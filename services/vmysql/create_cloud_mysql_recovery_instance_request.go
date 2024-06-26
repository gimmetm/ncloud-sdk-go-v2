/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type CreateCloudMysqlRecoveryInstanceRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// Subnet번호
SubnetNo *string `json:"subnetNo,omitempty"`

	// 백업파일이름
FileName *string `json:"fileName,omitempty"`

	// 복원시간
RecoveryTime *string `json:"recoveryTime,omitempty"`

	// CloudMysql인스턴스번호
CloudMysqlInstanceNo *string `json:"cloudMysqlInstanceNo"`

	// CloudMysqlRecovery서버이름
CloudMysqlRecoveryServerName *string `json:"cloudMysqlRecoveryServerName"`
}
