/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type InterruptServerInstanceResponse struct {

	// 서버인스턴스번호리스트
ServerInstanceNoList []*string `json:"serverInstanceNoList"`
}