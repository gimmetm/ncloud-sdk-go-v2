/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type GetRouteTableListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 라우트테이블번호리스트
RouteTableNoList []*string `json:"routeTableNoList,omitempty"`

	// 라우트테이블이름
RouteTableName *string `json:"routeTableName,omitempty"`

	// 지원하는서브넷유형코드
SupportedSubnetTypeCode *string `json:"supportedSubnetTypeCode,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 정렬대상
SortedBy *string `json:"sortedBy,omitempty"`

	// 정렬순서
SortingOrder *string `json:"sortingOrder,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`
}
