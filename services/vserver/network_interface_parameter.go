/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type NetworkInterfaceParameter struct {

	// 네트워크인터페이스순서
NetworkInterfaceOrder *int32 `json:"networkInterfaceOrder"`

	// 네트워크인터페이스번호
NetworkInterfaceNo *string `json:"networkInterfaceNo,omitempty"`

	// 서브넷번호
SubnetNo *string `json:"subnetNo,omitempty"`

	// IP주소
Ip *string `json:"ip,omitempty"`

	// ACG번호리스트
AccessControlGroupNoList []*string `json:"accessControlGroupNoList"`
}
