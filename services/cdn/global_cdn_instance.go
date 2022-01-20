/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type GlobalCdnInstance struct {

	// CDN인스턴스번호
	CdnInstanceNo *string `json:"cdnInstanceNo,omitempty"`

	// CDN인스턴스상태
	CdnInstanceStatus *CommonCode `json:"cdnInstanceStatus,omitempty"`

	// CDN인스턴스OP
	CdnInstanceOperation *CommonCode `json:"cdnInstanceOperation,omitempty"`

	// CDN인스턴스상태명
	CdnInstanceStatusName *string `json:"cdnInstanceStatusName,omitempty"`

	// 생성일자
	CreateDate *string `json:"createDate,omitempty"`

	// UPTIME
	LastModifiedDate *string `json:"lastModifiedDate,omitempty"`

	// CDN인스턴스설명
	CdnInstanceDescription *string `json:"cdnInstanceDescription,omitempty"`

	// 서비스이름
	ServiceName *string `json:"serviceName,omitempty"`

	IsAvailablePartialDomainPurge *bool `json:"isAvailablePartialDomainPurge,omitempty"`

	GlobalCdnServiceDomainList []*GlobalCdnServiceDomain `json:"globalCdnServiceDomainList,omitempty"`

	GlobalCdnRule *GlobalCdnRule `json:"globalCdnRule,omitempty"`
}
