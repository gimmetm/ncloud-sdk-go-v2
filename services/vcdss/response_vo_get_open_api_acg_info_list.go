/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type ResponseVoGetOpenApiAcgInfoList struct {
	Code int32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	RequestId string `json:"requestId,omitempty"`
	Result *GetOpenApiAcgInfoList `json:"result,omitempty"`
}
