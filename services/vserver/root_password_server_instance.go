/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type RootPasswordServerInstance struct {

	// 루트패스워드
	RootPassword *string `json:"rootPassword,omitempty"`

	// 서버인스턴스번호
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
}
