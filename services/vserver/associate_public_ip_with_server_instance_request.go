/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type AssociatePublicIpWithServerInstanceRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 공인IP인스턴스번호
PublicIpInstanceNo *string `json:"publicIpInstanceNo"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`
}
