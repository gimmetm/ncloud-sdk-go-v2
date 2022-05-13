/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type SnapshotHistoryResponseVo struct {
	ActionDate string `json:"actionDate,omitempty"`
	BucketName string `json:"bucketName,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	Duration string `json:"duration,omitempty"`
	HasErrorLog bool `json:"hasErrorLog,omitempty"`
	Idx int32 `json:"idx,omitempty"`
	MemberNo string `json:"memberNo,omitempty"`
	Message string `json:"message,omitempty"`
	RequestType string `json:"requestType,omitempty"`
	ServiceGroupInstanceNo int32 `json:"serviceGroupInstanceNo,omitempty"`
	SnapshotName string `json:"snapshotName,omitempty"`
}
