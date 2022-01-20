/*
 * vmysql
 *
 * <br/>https://ncloud.beta-apigw.gov-ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type CloudMysqlInstance struct {

	// CloudMysql인스턴스번호
	CloudMysqlInstanceNo *string `json:"cloudMysqlInstanceNo,omitempty"`

	// CloudMysql서비스이름
	CloudMysqlServiceName *string `json:"cloudMysqlServiceName,omitempty"`

	// CloudMysql인스턴스상태이름
	CloudMysqlInstanceStatusName *string `json:"cloudMysqlInstanceStatusName,omitempty"`

	// CloudMysql인스턴스상태
	CloudMysqlInstanceStatus *CommonCode `json:"cloudMysqlInstanceStatus,omitempty"`

	// CloudMysql인스턴스OP
	CloudMysqlInstanceOperation *CommonCode `json:"cloudMysqlInstanceOperation,omitempty"`

	// CloudMysql이미지상품코드
	CloudMysqlImageProductCode *string `json:"cloudMysqlImageProductCode,omitempty"`

	// CloudMysql엔진버전
	EngineVersion *string `json:"engineVersion,omitempty"`

	// CloudMysql라이선스
	License *CommonCode `json:"license,omitempty"`

	// CloudMysql포트
	CloudMysqlPort *int32 `json:"cloudMysqlPort,omitempty"`

	// 고가용성여부
	IsHa *bool `json:"isHa,omitempty"`

	// 멀티존여부
	IsMultiZone *bool `json:"isMultiZone,omitempty"`

	// 백업여부
	IsBackup *bool `json:"isBackup,omitempty"`

	// 백업파일보관기간
	BackupFileRetentionPeriod *int32 `json:"backupFileRetentionPeriod,omitempty"`

	// 백업시간
	BackupTime *string `json:"backupTime,omitempty"`

	// 생성일자
	CreateDate *string `json:"createDate,omitempty"`

	// ACG번호리스트
	AccessControlGroupNoList []*string `json:"accessControlGroupNoList,omitempty"`

	// CloudMysqlConfig리스트
	CloudMysqlConfigList []*string `json:"cloudMysqlConfigList,omitempty"`

	// CloudMysql서버인스턴스리스트
	CloudMysqlServerInstanceList []*CloudMysqlServerInstance `json:"cloudMysqlServerInstanceList,omitempty"`
}
