/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

type GetLoadBalancedServerInstanceListResponse struct {
	RequestId *string `json:"requestId,omitempty"`

	ReturnCode *string `json:"returnCode,omitempty"`

	ReturnMessage *string `json:"returnMessage,omitempty"`

	TotalRows *int32 `json:"totalRows,omitempty"`

	ServerInstanceList []*ServerInstance `json:"serverInstanceList,omitempty"`
}
