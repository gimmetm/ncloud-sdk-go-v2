/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type DeletePortForwardingRulesRequest struct {

	// 포트포워딩설정번호
PortForwardingConfigurationNo *string `json:"portForwardingConfigurationNo"`

	// 포트포워딩RULE리스트
PortForwardingRuleList []*PortForwardingRuleParameter `json:"portForwardingRuleList"`
}
