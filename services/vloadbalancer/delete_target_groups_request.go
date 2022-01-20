/*
 * vloadbalancer
 *
 * VPC Load Balancer 관련 API<br/>https://ncloud.apigw.ntruss.com/vloadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancer

type DeleteTargetGroupsRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// 타겟그룹번호리스트
	TargetGroupNoList []*string `json:"targetGroupNoList"`
}
