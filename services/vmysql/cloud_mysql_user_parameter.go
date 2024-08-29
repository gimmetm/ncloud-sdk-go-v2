/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type CloudMysqlUserParameter struct {

	// DBUserID
Name *string `json:"name"`

	// 접근IP
HostIp *string `json:"hostIp"`

	// DBUserPassword
Password *string `json:"password"`

	// DB권한
Authority *string `json:"authority"`

	// 시스템테이블 접근 가능여부
IsSystemTableAccess *bool `json:"isSystemTableAccess,omitempty"`
}
