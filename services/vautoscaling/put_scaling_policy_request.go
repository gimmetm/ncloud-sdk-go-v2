/*
 * vautoscaling
 *
 * VPC Auto Scaling 관련 API<br/>https://ncloud.apigw.ntruss.com/vautoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vautoscaling

type PutScalingPolicyRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// 오토스케일링그룹번호
	AutoScalingGroupNo *string `json:"autoScalingGroupNo"`

	// 정책번호
	PolicyNo *string `json:"policyNo,omitempty"`

	// 정책이름
	PolicyName *string `json:"policyName,omitempty"`

	// 조정유형코드
	AdjustmentTypeCode *string `json:"adjustmentTypeCode"`

	// 조정값
	ScalingAdjustment *int32 `json:"scalingAdjustment"`

	// 최소조정폭
	MinAdjustmentStep *int32 `json:"minAdjustmentStep,omitempty"`

	// 쿨다운
	CoolDown *int32 `json:"coolDown,omitempty"`
}
