/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type BlockStorageMapping struct {

	// 블록스토리지순서
Order *int32 `json:"order,omitempty"`

	// 블록스토리지스냅샷인스턴스번호
BlockStorageSnapshotInstanceNo *string `json:"blockStorageSnapshotInstanceNo,omitempty"`

	// 블록스토리지스냅샷이름
BlockStorageSnapshotName *string `json:"blockStorageSnapshotName,omitempty"`

	// 블록스토리지사이즈
BlockStorageSize *float32 `json:"blockStorageSize,omitempty"`

	// 블록스토리지이름
BlockStorageName *string `json:"blockStorageName,omitempty"`

	// 블록스토리지볼륨유형
BlockStorageVolumeType *CommonCode `json:"blockStorageVolumeType,omitempty"`

	// IOPS
Iops *float32 `json:"iops,omitempty"`

	// 부하처리성능
Throughput *float32 `json:"throughput,omitempty"`

	// 볼륨암호화여부
IsEncryptedVolume *bool `json:"isEncryptedVolume,omitempty"`
}
