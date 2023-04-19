/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type DetachNetworkInterfaceRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`

	// 네트워크인터페이스번호
NetworkInterfaceNo *string `json:"networkInterfaceNo"`

	// 서브넷번호
SubnetNo *string `json:"subnetNo"`
}
