/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type GetSnapshotSchedulingHistoryResponseVo struct {
	CurrentPage int32 `json:"currentPage,omitempty"`
	IsFirst bool `json:"isFirst,omitempty"`
	IsLast bool `json:"isLast,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
	SnapshotSchedulingHistoryList []SnapshotSchedulingHistoryResponseVo `json:"snapshotSchedulingHistoryList,omitempty"`
	TotalCount int64 `json:"totalCount,omitempty"`
	TotalPage int32 `json:"totalPage,omitempty"`
}