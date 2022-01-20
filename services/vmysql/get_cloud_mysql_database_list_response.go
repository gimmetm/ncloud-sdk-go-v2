/*
 * vmysql
 *
 * <br/>https://ncloud.beta-apigw.gov-ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type GetCloudMysqlDatabaseListResponse struct {
	RequestId *string `json:"requestId,omitempty"`

	ReturnCode *string `json:"returnCode,omitempty"`

	ReturnMessage *string `json:"returnMessage,omitempty"`

	TotalRows *int32 `json:"totalRows,omitempty"`

	// CloudMysqlDatabase리스트
	CloudMysqlInstanceList *CloudMysqlDatabaseList `json:"cloudMysqlInstanceList,omitempty"`
}
