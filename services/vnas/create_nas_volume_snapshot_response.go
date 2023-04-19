/*
 * vnas
 *
 * VPC NAS 관련 API<br/>https://ncloud.apigw.ntruss.com/vnas/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnas

type CreateNasVolumeSnapshotResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

NasVolumeSnapshotList []*NasVolumeSnapshot `json:"nasVolumeSnapshotList,omitempty"`
}
