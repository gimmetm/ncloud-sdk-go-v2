/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type DeleteMemberServerImagesRequest struct {

	// 회원서버이미지번호리스트
	MemberServerImageNoList []*string `json:"memberServerImageNoList"`
}
