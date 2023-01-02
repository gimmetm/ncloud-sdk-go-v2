/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type GetOpenApiServiceGroupResponseVo struct {
	ComputeList []OpenApiGetComputeInstanceListResponseDetailVo `json:"computeList,omitempty"`
	ContractNo int32 `json:"contractNo,omitempty"`
	CreatedDate *DateTimeVo `json:"createdDate,omitempty"`
	ProductCode string `json:"productCode,omitempty"`
	RegionNo int32 `json:"regionNo,omitempty"`
	ServiceGroupInstanceNo int32 `json:"serviceGroupInstanceNo,omitempty"`
	ServiceGroupInstanceTypeCode string `json:"serviceGroupInstanceTypeCode,omitempty"`
	ServiceGroupName string `json:"serviceGroupName,omitempty"`
	StatusCode string `json:"statusCode,omitempty"`
	VpcNo int32 `json:"vpcNo,omitempty"`
}