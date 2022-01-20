/*
 * vmysql
 *
 * <br/>https://ncloud.beta-apigw.gov-ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type CreateCloudMysqlRecoveryInstanceRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// 백업파일이름
	FileName *string `json:"fileName,omitempty"`

	// 복원시간
	RecoveryTime *string `json:"recoveryTime,omitempty"`

	// CloudMysql인스턴스번호
	CloudMysqlInstanceNo *string `json:"cloudMysqlInstanceNo"`

	// CloudMysqlRecovery서버이름
	CloudMysqlRecoveryServerName *string `json:"cloudMysqlRecoveryServerName"`
}
