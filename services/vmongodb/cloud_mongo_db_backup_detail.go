/*
 * vmongodb
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmongodb/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmongodb

type CloudMongoDbBackupDetail struct {

	// 백업시작시간
StartTime *string `json:"startTime,omitempty"`

	// 백업종료시간
EndTime *string `json:"endTime,omitempty"`

	// 백업사이즈
BackupSize *int64 `json:"backupSize,omitempty"`

	// 데이터스토리지사이즈
DataStorageSize *int64 `json:"dataStorageSize,omitempty"`

	// 백업속도
BackupParallel *int32 `json:"backupParallel,omitempty"`

	// 샤드(Replica Set)이름
Shard *string `json:"shard,omitempty"`
}
