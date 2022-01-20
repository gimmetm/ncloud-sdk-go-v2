/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreateInitScriptRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// 초기화스크립트내용
	InitScriptContent *string `json:"initScriptContent"`

	// 초기화스크립트이름
	InitScriptName *string `json:"initScriptName,omitempty"`

	// 초기화스크립트설명
	InitScriptDescription *string `json:"initScriptDescription,omitempty"`

	// OS유형코드
	OsTypeCode *string `json:"osTypeCode"`
}
