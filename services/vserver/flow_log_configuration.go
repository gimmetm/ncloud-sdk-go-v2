/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type FlowLogConfiguration struct {

NetworkInterfaceNo *string `json:"networkInterfaceNo"`

CollectActionType *CommonCode `json:"collectActionType"`

CollectIntervalMinute *int32 `json:"collectIntervalMinute,omitempty"`

StorageType *CommonCode `json:"storageType,omitempty"`

StorageBucketName *string `json:"storageBucketName"`

StorageBucketDirectoryName *string `json:"storageBucketDirectoryName,omitempty"`
}
