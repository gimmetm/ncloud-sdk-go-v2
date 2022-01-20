/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type CreateRouteTableRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// 라우트테이블설명
	RouteTableDescription *string `json:"routeTableDescription,omitempty"`

	// 지원하는서브넷유형코드
	SupportedSubnetTypeCode *string `json:"supportedSubnetTypeCode"`

	// 라우트테이블이름
	RouteTableName *string `json:"routeTableName,omitempty"`

	// VPC번호
	VpcNo *string `json:"vpcNo"`
}
