/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetServerImageProductListResponse struct {
	RequestId *string `json:"requestId,omitempty"`

	ReturnCode *string `json:"returnCode,omitempty"`

	ReturnMessage *string `json:"returnMessage,omitempty"`

	ProductList []*Product `json:"productList,omitempty"`

	TotalRows *int32 `json:"totalRows,omitempty"`
}
