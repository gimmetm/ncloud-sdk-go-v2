# Go API client for vhadoop

    <br/>https://ncloud.apigw.ntruss.com/vhadoop/v2

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2021-11-25T11:33:57Z
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.NcpGoForNcloudClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
"./vhadoop"
```

## Documentation for API Endpoints

All URIs are relative to *https://ncloud.apigw.ntruss.com/vhadoop/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V2Api* | [**BackupClusterConfiguration**](docs/V2Api.md#backupclusterconfiguration) | **Post** /backupClusterConfiguration | 
*V2Api* | [**ChangeCloudHadoopNodeCount**](docs/V2Api.md#changecloudhadoopnodecount) | **Post** /changeCloudHadoopNodeCount | 
*V2Api* | [**ChangeCloudHadoopNodeSpec**](docs/V2Api.md#changecloudhadoopnodespec) | **Post** /changeCloudHadoopNodeSpec | 
*V2Api* | [**CreateCloudHadoopInstance**](docs/V2Api.md#createcloudhadoopinstance) | **Post** /createCloudHadoopInstance | 
*V2Api* | [**DeleteCloudHadoopInstance**](docs/V2Api.md#deletecloudhadoopinstance) | **Post** /deleteCloudHadoopInstance | 
*V2Api* | [**GetCloudHadoopAddOnList**](docs/V2Api.md#getcloudhadoopaddonlist) | **Post** /getCloudHadoopAddOnList | 
*V2Api* | [**GetCloudHadoopBucketList**](docs/V2Api.md#getcloudhadoopbucketlist) | **Post** /getCloudHadoopBucketList | 
*V2Api* | [**GetCloudHadoopClusterTypeList**](docs/V2Api.md#getcloudhadoopclustertypelist) | **Post** /getCloudHadoopClusterTypeList | 
*V2Api* | [**GetCloudHadoopImageProductList**](docs/V2Api.md#getcloudhadoopimageproductlist) | **Post** /getCloudHadoopImageProductList | 
*V2Api* | [**GetCloudHadoopInstanceDetail**](docs/V2Api.md#getcloudhadoopinstancedetail) | **Post** /getCloudHadoopInstanceDetail | 
*V2Api* | [**GetCloudHadoopInstanceList**](docs/V2Api.md#getcloudhadoopinstancelist) | **Post** /getCloudHadoopInstanceList | 
*V2Api* | [**GetCloudHadoopLoginKeyList**](docs/V2Api.md#getcloudhadooploginkeylist) | **Post** /getCloudHadoopLoginKeyList | 
*V2Api* | [**GetCloudHadoopProductList**](docs/V2Api.md#getcloudhadoopproductlist) | **Post** /getCloudHadoopProductList | 


## Documentation For Models

 - [AccessControlGroupNoList](docs/AccessControlGroupNoList.md)
 - [BackupClusterConfigurationRequest](docs/BackupClusterConfigurationRequest.md)
 - [BackupClusterConfigurationResponse](docs/BackupClusterConfigurationResponse.md)
 - [ChangeCloudHadoopNodeCountRequest](docs/ChangeCloudHadoopNodeCountRequest.md)
 - [ChangeCloudHadoopNodeCountResponse](docs/ChangeCloudHadoopNodeCountResponse.md)
 - [ChangeCloudHadoopNodeSpecRequest](docs/ChangeCloudHadoopNodeSpecRequest.md)
 - [ChangeCloudHadoopNodeSpecResponse](docs/ChangeCloudHadoopNodeSpecResponse.md)
 - [CloudHadoopAddOn](docs/CloudHadoopAddOn.md)
 - [CloudHadoopAddOnList](docs/CloudHadoopAddOnList.md)
 - [CloudHadoopBucket](docs/CloudHadoopBucket.md)
 - [CloudHadoopBucketList](docs/CloudHadoopBucketList.md)
 - [CloudHadoopClusterType](docs/CloudHadoopClusterType.md)
 - [CloudHadoopClusterTypeList](docs/CloudHadoopClusterTypeList.md)
 - [CloudHadoopInstance](docs/CloudHadoopInstance.md)
 - [CloudHadoopInstanceList](docs/CloudHadoopInstanceList.md)
 - [CloudHadoopLoginKey](docs/CloudHadoopLoginKey.md)
 - [CloudHadoopLoginKeyList](docs/CloudHadoopLoginKeyList.md)
 - [CloudHadoopServerInstance](docs/CloudHadoopServerInstance.md)
 - [CloudHadoopVersion](docs/CloudHadoopVersion.md)
 - [CommonCode](docs/CommonCode.md)
 - [CreateCloudHadoopInstanceRequest](docs/CreateCloudHadoopInstanceRequest.md)
 - [CreateCloudHadoopInstanceResponse](docs/CreateCloudHadoopInstanceResponse.md)
 - [DeleteCloudHadoopInstanceRequest](docs/DeleteCloudHadoopInstanceRequest.md)
 - [DeleteCloudHadoopInstanceResponse](docs/DeleteCloudHadoopInstanceResponse.md)
 - [GetCloudHadoopAddOnListRequest](docs/GetCloudHadoopAddOnListRequest.md)
 - [GetCloudHadoopAddOnListResponse](docs/GetCloudHadoopAddOnListResponse.md)
 - [GetCloudHadoopBucketListRequest](docs/GetCloudHadoopBucketListRequest.md)
 - [GetCloudHadoopBucketListResponse](docs/GetCloudHadoopBucketListResponse.md)
 - [GetCloudHadoopClusterTypeListRequest](docs/GetCloudHadoopClusterTypeListRequest.md)
 - [GetCloudHadoopClusterTypeListResponse](docs/GetCloudHadoopClusterTypeListResponse.md)
 - [GetCloudHadoopImageProductListRequest](docs/GetCloudHadoopImageProductListRequest.md)
 - [GetCloudHadoopImageProductListResponse](docs/GetCloudHadoopImageProductListResponse.md)
 - [GetCloudHadoopInstanceDetailRequest](docs/GetCloudHadoopInstanceDetailRequest.md)
 - [GetCloudHadoopInstanceDetailResponse](docs/GetCloudHadoopInstanceDetailResponse.md)
 - [GetCloudHadoopInstanceListRequest](docs/GetCloudHadoopInstanceListRequest.md)
 - [GetCloudHadoopInstanceListResponse](docs/GetCloudHadoopInstanceListResponse.md)
 - [GetCloudHadoopLoginKeyListRequest](docs/GetCloudHadoopLoginKeyListRequest.md)
 - [GetCloudHadoopLoginKeyListResponse](docs/GetCloudHadoopLoginKeyListResponse.md)
 - [GetCloudHadoopProductListRequest](docs/GetCloudHadoopProductListRequest.md)
 - [GetCloudHadoopProductListResponse](docs/GetCloudHadoopProductListResponse.md)
 - [Product](docs/Product.md)
 - [ProductList](docs/ProductList.md)
