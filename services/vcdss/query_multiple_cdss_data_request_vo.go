/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type QueryMultipleCdssDataRequestVo struct {
	ComputeInstanceNo string `json:"computeInstanceNo,omitempty"`
	Interval string `json:"interval,omitempty"`
	Metric string `json:"metric,omitempty"`
	MetricInfoList []MetricInfoVo `json:"metricInfoList,omitempty"`
	ServiceGroupInstanceNo string `json:"serviceGroupInstanceNo,omitempty"`
	TimeEnd int64 `json:"timeEnd,omitempty"`
	TimeStart int64 `json:"timeStart,omitempty"`
}
