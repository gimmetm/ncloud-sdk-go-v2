# Go API client for vmysql

    <br/>https://ncloud.apigw.ntruss.com/vmysql/v2

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2024-08-29T16:10:02Z
- Package version: 1.0.2
- Build package: io.swagger.codegen.languages.NcpGoForNcloudClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
"./vmysql"
```

## Documentation for API Endpoints

All URIs are relative to *https://ncloud.apigw.ntruss.com/vmysql/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V2Api* | [**AddCloudMysqlDatabaseList**](docs/V2Api.md#addcloudmysqldatabaselist) | **Post** /addCloudMysqlDatabaseList | 
*V2Api* | [**AddCloudMysqlUserList**](docs/V2Api.md#addcloudmysqluserlist) | **Post** /addCloudMysqlUserList | 
*V2Api* | [**ChangeCloudMysqlUserList**](docs/V2Api.md#changecloudmysqluserlist) | **Post** /changeCloudMysqlUserList | 
*V2Api* | [**CreateCloudMysqlInstance**](docs/V2Api.md#createcloudmysqlinstance) | **Post** /createCloudMysqlInstance | 
*V2Api* | [**CreateCloudMysqlRecoveryInstance**](docs/V2Api.md#createcloudmysqlrecoveryinstance) | **Post** /createCloudMysqlRecoveryInstance | 
*V2Api* | [**CreateCloudMysqlSlaveInstance**](docs/V2Api.md#createcloudmysqlslaveinstance) | **Post** /createCloudMysqlSlaveInstance | 
*V2Api* | [**DeleteCloudMysqlDatabaseList**](docs/V2Api.md#deletecloudmysqldatabaselist) | **Post** /deleteCloudMysqlDatabaseList | 
*V2Api* | [**DeleteCloudMysqlInstance**](docs/V2Api.md#deletecloudmysqlinstance) | **Post** /deleteCloudMysqlInstance | 
*V2Api* | [**DeleteCloudMysqlServerInstance**](docs/V2Api.md#deletecloudmysqlserverinstance) | **Post** /deleteCloudMysqlServerInstance | 
*V2Api* | [**DeleteCloudMysqlUserList**](docs/V2Api.md#deletecloudmysqluserlist) | **Post** /deleteCloudMysqlUserList | 
*V2Api* | [**ExportBackupToObjectStorage**](docs/V2Api.md#exportbackuptoobjectstorage) | **Post** /exportBackupToObjectStorage | 
*V2Api* | [**ExportDbServerLogToObjectStorage**](docs/V2Api.md#exportdbserverlogtoobjectstorage) | **Post** /exportDbServerLogToObjectStorage | 
*V2Api* | [**GetCloudMysqlBackupDetailList**](docs/V2Api.md#getcloudmysqlbackupdetaillist) | **Post** /getCloudMysqlBackupDetailList | 
*V2Api* | [**GetCloudMysqlBackupList**](docs/V2Api.md#getcloudmysqlbackuplist) | **Post** /getCloudMysqlBackupList | 
*V2Api* | [**GetCloudMysqlDatabaseList**](docs/V2Api.md#getcloudmysqldatabaselist) | **Post** /getCloudMysqlDatabaseList | 
*V2Api* | [**GetCloudMysqlImageProductList**](docs/V2Api.md#getcloudmysqlimageproductlist) | **Post** /getCloudMysqlImageProductList | 
*V2Api* | [**GetCloudMysqlInstanceDetail**](docs/V2Api.md#getcloudmysqlinstancedetail) | **Post** /getCloudMysqlInstanceDetail | 
*V2Api* | [**GetCloudMysqlInstanceList**](docs/V2Api.md#getcloudmysqlinstancelist) | **Post** /getCloudMysqlInstanceList | 
*V2Api* | [**GetCloudMysqlProductList**](docs/V2Api.md#getcloudmysqlproductlist) | **Post** /getCloudMysqlProductList | 
*V2Api* | [**GetCloudMysqlRecoveryTime**](docs/V2Api.md#getcloudmysqlrecoverytime) | **Post** /getCloudMysqlRecoveryTime | 
*V2Api* | [**GetCloudMysqlTargetSubnetList**](docs/V2Api.md#getcloudmysqltargetsubnetlist) | **Post** /getCloudMysqlTargetSubnetList | 
*V2Api* | [**GetCloudMysqlTargetVpcList**](docs/V2Api.md#getcloudmysqltargetvpclist) | **Post** /getCloudMysqlTargetVpcList | 
*V2Api* | [**GetCloudMysqlUserList**](docs/V2Api.md#getcloudmysqluserlist) | **Post** /getCloudMysqlUserList | 
*V2Api* | [**GetDbServerLogList**](docs/V2Api.md#getdbserverloglist) | **Post** /getDbServerLogList | 
*V2Api* | [**RebootCloudMysqlServerInstance**](docs/V2Api.md#rebootcloudmysqlserverinstance) | **Post** /rebootCloudMysqlServerInstance | 


## Documentation For Models

 - [AccessControlGroupNoList](docs/AccessControlGroupNoList.md)
 - [AddCloudMysqlDatabaseListRequest](docs/AddCloudMysqlDatabaseListRequest.md)
 - [AddCloudMysqlDatabaseListResponse](docs/AddCloudMysqlDatabaseListResponse.md)
 - [AddCloudMysqlUserListRequest](docs/AddCloudMysqlUserListRequest.md)
 - [AddCloudMysqlUserListResponse](docs/AddCloudMysqlUserListResponse.md)
 - [ChangeCloudMysqlUserListRequest](docs/ChangeCloudMysqlUserListRequest.md)
 - [ChangeCloudMysqlUserListResponse](docs/ChangeCloudMysqlUserListResponse.md)
 - [CloudDbProduct](docs/CloudDbProduct.md)
 - [CloudDbProductList](docs/CloudDbProductList.md)
 - [CloudMysqlBackup](docs/CloudMysqlBackup.md)
 - [CloudMysqlBackupDetail](docs/CloudMysqlBackupDetail.md)
 - [CloudMysqlConfigList](docs/CloudMysqlConfigList.md)
 - [CloudMysqlDatabase](docs/CloudMysqlDatabase.md)
 - [CloudMysqlDbServerLog](docs/CloudMysqlDbServerLog.md)
 - [CloudMysqlInstance](docs/CloudMysqlInstance.md)
 - [CloudMysqlRecoveryTime](docs/CloudMysqlRecoveryTime.md)
 - [CloudMysqlServerInstance](docs/CloudMysqlServerInstance.md)
 - [CloudMysqlUser](docs/CloudMysqlUser.md)
 - [CloudMysqlUserKeyParameter](docs/CloudMysqlUserKeyParameter.md)
 - [CloudMysqlUserParameter](docs/CloudMysqlUserParameter.md)
 - [CommonCode](docs/CommonCode.md)
 - [CreateCloudMysqlInstanceRequest](docs/CreateCloudMysqlInstanceRequest.md)
 - [CreateCloudMysqlInstanceResponse](docs/CreateCloudMysqlInstanceResponse.md)
 - [CreateCloudMysqlRecoveryInstanceRequest](docs/CreateCloudMysqlRecoveryInstanceRequest.md)
 - [CreateCloudMysqlRecoveryInstanceResponse](docs/CreateCloudMysqlRecoveryInstanceResponse.md)
 - [CreateCloudMysqlSlaveInstanceRequest](docs/CreateCloudMysqlSlaveInstanceRequest.md)
 - [CreateCloudMysqlSlaveInstanceResponse](docs/CreateCloudMysqlSlaveInstanceResponse.md)
 - [DeleteCloudMysqlDatabaseListRequest](docs/DeleteCloudMysqlDatabaseListRequest.md)
 - [DeleteCloudMysqlDatabaseListResponse](docs/DeleteCloudMysqlDatabaseListResponse.md)
 - [DeleteCloudMysqlInstanceRequest](docs/DeleteCloudMysqlInstanceRequest.md)
 - [DeleteCloudMysqlInstanceResponse](docs/DeleteCloudMysqlInstanceResponse.md)
 - [DeleteCloudMysqlServerInstanceRequest](docs/DeleteCloudMysqlServerInstanceRequest.md)
 - [DeleteCloudMysqlServerInstanceResponse](docs/DeleteCloudMysqlServerInstanceResponse.md)
 - [DeleteCloudMysqlUserListRequest](docs/DeleteCloudMysqlUserListRequest.md)
 - [DeleteCloudMysqlUserListResponse](docs/DeleteCloudMysqlUserListResponse.md)
 - [ExportBackupToObjectStorageRequest](docs/ExportBackupToObjectStorageRequest.md)
 - [ExportBackupToObjectStorageResponse](docs/ExportBackupToObjectStorageResponse.md)
 - [ExportDbServerLogToObjectStorageRequest](docs/ExportDbServerLogToObjectStorageRequest.md)
 - [ExportDbServerLogToObjectStorageResponse](docs/ExportDbServerLogToObjectStorageResponse.md)
 - [GetCloudMysqlBackupDetailListRequest](docs/GetCloudMysqlBackupDetailListRequest.md)
 - [GetCloudMysqlBackupDetailListResponse](docs/GetCloudMysqlBackupDetailListResponse.md)
 - [GetCloudMysqlBackupListRequest](docs/GetCloudMysqlBackupListRequest.md)
 - [GetCloudMysqlBackupListResponse](docs/GetCloudMysqlBackupListResponse.md)
 - [GetCloudMysqlDatabaseListRequest](docs/GetCloudMysqlDatabaseListRequest.md)
 - [GetCloudMysqlDatabaseListResponse](docs/GetCloudMysqlDatabaseListResponse.md)
 - [GetCloudMysqlImageProductListRequest](docs/GetCloudMysqlImageProductListRequest.md)
 - [GetCloudMysqlImageProductListResponse](docs/GetCloudMysqlImageProductListResponse.md)
 - [GetCloudMysqlInstanceDetailRequest](docs/GetCloudMysqlInstanceDetailRequest.md)
 - [GetCloudMysqlInstanceDetailResponse](docs/GetCloudMysqlInstanceDetailResponse.md)
 - [GetCloudMysqlInstanceListRequest](docs/GetCloudMysqlInstanceListRequest.md)
 - [GetCloudMysqlInstanceListResponse](docs/GetCloudMysqlInstanceListResponse.md)
 - [GetCloudMysqlProductListRequest](docs/GetCloudMysqlProductListRequest.md)
 - [GetCloudMysqlProductListResponse](docs/GetCloudMysqlProductListResponse.md)
 - [GetCloudMysqlRecoveryTimeRequest](docs/GetCloudMysqlRecoveryTimeRequest.md)
 - [GetCloudMysqlRecoveryTimeResponse](docs/GetCloudMysqlRecoveryTimeResponse.md)
 - [GetCloudMysqlTargetSubnetListRequest](docs/GetCloudMysqlTargetSubnetListRequest.md)
 - [GetCloudMysqlTargetSubnetListResponse](docs/GetCloudMysqlTargetSubnetListResponse.md)
 - [GetCloudMysqlTargetVpcListRequest](docs/GetCloudMysqlTargetVpcListRequest.md)
 - [GetCloudMysqlTargetVpcListResponse](docs/GetCloudMysqlTargetVpcListResponse.md)
 - [GetCloudMysqlUserListRequest](docs/GetCloudMysqlUserListRequest.md)
 - [GetCloudMysqlUserListResponse](docs/GetCloudMysqlUserListResponse.md)
 - [GetDbServerLogListRequest](docs/GetDbServerLogListRequest.md)
 - [GetDbServerLogListResponse](docs/GetDbServerLogListResponse.md)
 - [RebootCloudMysqlServerInstanceRequest](docs/RebootCloudMysqlServerInstanceRequest.md)
 - [RebootCloudMysqlServerInstanceResponse](docs/RebootCloudMysqlServerInstanceResponse.md)
 - [TargetSubnet](docs/TargetSubnet.md)
 - [TargetVpc](docs/TargetVpc.md)

