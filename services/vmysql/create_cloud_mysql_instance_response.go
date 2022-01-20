/*
 * vmysql
 *
 * <br/>https://ncloud.beta-apigw.gov-ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type CreateCloudMysqlInstanceResponse struct {
	RequestId *string `json:"requestId,omitempty"`

	ReturnCode *string `json:"returnCode,omitempty"`

	ReturnMessage *string `json:"returnMessage,omitempty"`

	TotalRows *int32 `json:"totalRows,omitempty"`

	// CloudMysql인스턴스리스트
	CloudMysqlInstanceList []*CloudMysqlInstance `json:"cloudMysqlInstanceList,omitempty"`
}
