/*
 * vnks
 *
 * <br/>https://nks.apigw.ntruss.com/vnks/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnks

type NodePoolCreationBody struct {

	// 노드풀 이름
	Name *string `json:"name"`

	// 등록 될 노드 개수
	NodeCount *int32 `json:"nodeCount"`

	// Subnet 번호
	SubnetNo *int32 `json:"subnetNo,omitempty"`

	// Subnet 번호
	SubnetNoList []*int32 `json:"subnetNoList,omitempty"`

	// Server image code
	SoftwareCode *string `json:"softwareCode,omitempty"`

	// 상품 코드 [서버 스펙 목록](/docs/compute-vserver-server-common-getserverproductlist)
	ProductCode *string `json:"productCode,omitempty"`

	// Server spec code
	ServerSpecCode *string `json:"serverSpecCode,omitempty"`

	// Storage size
	StorageSize *int32 `json:"storageSize,omitempty"`

	//
	Autoscale *AutoscalerUpdate `json:"autoscale,omitempty"`
}
