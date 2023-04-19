/*
 * vloadbalancer
 *
 * VPC Load Balancer 관련 API<br/>https://ncloud.apigw.ntruss.com/vloadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancer

type GetLoadBalancerInstanceListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 로드밸런서인스턴스번호리스트
LoadBalancerInstanceNoList []*string `json:"loadBalancerInstanceNoList,omitempty"`

	// 로드밸런서네트워크유형코드
LoadBalancerNetworkTypeCode *string `json:"loadBalancerNetworkTypeCode,omitempty"`

	// 로드밸런서유형코드
LoadBalancerTypeCode *string `json:"loadBalancerTypeCode,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 정렬리스트
SortList *string `json:"sortList,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`
}
