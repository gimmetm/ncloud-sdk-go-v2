/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type CreateServerInstancesRequest struct {

	// 서버이미지상품코드
ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`

	// 서버상품코드
ServerProductCode *string `json:"serverProductCode,omitempty"`

	// 회원서버이미지번호
MemberServerImageNo *string `json:"memberServerImageNo,omitempty"`

	// 서버명
ServerName *string `json:"serverName,omitempty"`

	// 서버설명
ServerDescription *string `json:"serverDescription,omitempty"`

	// 로그인키명
LoginKeyName *string `json:"loginKeyName,omitempty"`

	// 반납보호여부
IsProtectServerTermination *bool `json:"isProtectServerTermination,omitempty"`

	// 서버생성갯수
ServerCreateCount *int32 `json:"serverCreateCount,omitempty"`

	// 서버생성시작번호
ServerCreateStartNo *int32 `json:"serverCreateStartNo,omitempty"`

	// 요금제구분코드
FeeSystemTypeCode *string `json:"feeSystemTypeCode,omitempty"`

	// 사용자데이터
UserData *string `json:"userData,omitempty"`

	// 초기화스크립트번호
InitScriptNo *string `json:"initScriptNo,omitempty"`

	// ZONE번호
ZoneNo *string `json:"zoneNo,omitempty"`

	// ACG설정번호리스트
AccessControlGroupConfigurationNoList []*string `json:"accessControlGroupConfigurationNoList,omitempty"`

	// RAID구분이름
RaidTypeName *string `json:"raidTypeName,omitempty"`

	// 인스턴스태그리스트
InstanceTagList []*InstanceTagParameter `json:"instanceTagList,omitempty"`

	// 백신설치여부
IsVaccineInstall *bool `json:"isVaccineInstall,omitempty"`

	// 블록디바이스파티션리스트
BlockDevicePartitionList []*BlockDevicePartition `json:"blockDevicePartitionList,omitempty"`
}
