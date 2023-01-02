# Go API client for sourcecommit

    <br/>https://sourcecommit.apigw.ntruss.com/api/v1

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2022-04-22T08:24:55Z
- Package version: 
- Build package: io.swagger.codegen.languages.NcpGoForVnksClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
"./sourcecommit"
```

## Documentation for API Endpoints

All URIs are relative to *https://sourcecommit.apigw.ntruss.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V1Api* | [**GetRepositories**](docs/V1Api.md#GetRepositories) | **Get** /repository | 
*V1Api* | [**GetRepository**](docs/V1Api.md#GetRepository) | **Get** /repository/{repositoryName} | 
*V1Api* | [**GetRepositoryById**](docs/V1Api.md#GetRepositoryById) | **Get** /repository/{repositoryId} | 
*V1Api* | [**CreateRepository**](docs/V1Api.md#CreateRepository) | **Post** /repository | 
*V1Api* | [**DeleteRepository**](docs/V1Api.md#DeleteRepository) | **Delete** /repository/id/{repositoryId} | 
*V1Api* | [**ChangeRepository**](docs/V1Api.md#ChangeRepository) | **Patch** /repository/id/{repositoryId} | 


## Documentation For Models

 - [ChangeRepository](docs/ChangeRepository.md)
 - [CreateRepository](docs/CreateRepository.md)
 - [CreateRepositoryLinked](docs/CreateRepositoryLinked.md)
 - [DefaultBranch](docs/DefaultBranch.md)
 - [GetRepositoryBranchListResponse](docs/GetRepositoryBranchListResponse.md)
 - [GetRepositoryBranchListResponseResult](docs/GetRepositoryBranchListResponseResult.md)
 - [GetRepositoryDetailResponse](docs/GetRepositoryDetailResponse.md)
 - [GetRepositoryDetailResponseResult](docs/GetRepositoryDetailResponseResult.md)
 - [GetRepositoryDetailResponseResultCreated](docs/GetRepositoryDetailResponseResultCreated.md)
 - [GetRepositoryDetailResponseResultGit](docs/GetRepositoryDetailResponseResultGit.md)
 - [GetRepositoryDetailResponseResultLinked](docs/GetRepositoryDetailResponseResultLinked.md)
 - [GetRepositoryListResponse](docs/GetRepositoryListResponse.md)
 - [GetRepositoryListResponseResult](docs/GetRepositoryListResponseResult.md)
 - [GetRepositoryListResponseResultRepository](docs/GetRepositoryListResponseResultRepository.md)
 - [GetRepositoryTagListResponse](docs/GetRepositoryTagListResponse.md)
 - [GetRepositoryTagListResponseResult](docs/GetRepositoryTagListResponseResult.md)
 - [OkResponse](docs/OkResponse.md)
