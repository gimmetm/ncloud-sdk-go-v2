# Go API client for vsourcedeploy

    <br/>https://vpcsourcedeploy.apigw.ntruss.com/api/v1

## Overview

This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project. By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

-   API version: 2022-04-22T07:37:28Z
-   Package version:
-   Build package: io.swagger.codegen.languages.NcpGoForVnksClientCodegen

## Installation

Put the package under your project folder and add the following in import:

```
"./vsourcedeploy"
```

## Documentation for API Endpoints

All URIs are relative to *https://vpcsourcedeploy.apigw.ntruss.com/api/v1*

| Class   | Method                                                                             | HTTP request                                                                       | Description |
| ------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ----------- |
| _V1Api_ | [**GetAutoscalingGroupTargetGroup**](docs/V1Api.md#GetAutoscalingGroupTargetGroup) | **Get** /autoscaling/{autoscalingName}                                             |
| _V1Api_ | [**GetAutoscalingGroups**](docs/V1Api.md#GetAutoscalingGroups)                     | **Get** /autoscaling                                                               |
| _V1Api_ | [**GetKubernetesClusters**](docs/V1Api.md#GetKubernetesClusters)                   | **Get** /kubernetes/cluster                                                        |
| _V1Api_ | [**GetObjectstorageObjects**](docs/V1Api.md#GetObjectstorageObjects)               | **Get** /objectstorage/bucket/{bucketName}                                         |
| _V1Api_ | [**GetObjectstorageBuckets**](docs/V1Api.md#GetObjectstorageBuckets)               | **Get** /objectstorage/bucket                                                      |
| _V1Api_ | [**GetProjects**](docs/V1Api.md#GetProjects)                                       | **Get** /project                                                                   |
| _V1Api_ | [**CreateProject**](docs/V1Api.md#CreateProject)                                   | **Post** /project                                                                  |
| _V1Api_ | [**DeleteProject**](docs/V1Api.md#DeleteProject)                                   | **Delete** /project/{projectId}                                                    |
| _V1Api_ | [**GetHistories**](docs/V1Api.md#GetHistories)                                     | **Get** /project/{projectId}/history                                               |
| _V1Api_ | [**AcceptDeployApproval**](docs/V1Api.md#AcceptDeployApproval)                     | **Post** /project/{projectId}/history/{historyId}/approval/accept                  |
| _V1Api_ | [**RejectDeployApproval**](docs/V1Api.md#RejectDeployApproval)                     | **Post** /project/{projectId}/history/{historyId}/approval/reject                  |
| _V1Api_ | [**AcceptDeployCanary**](docs/V1Api.md#AcceptDeployCanary)                         | **Post** /project/{projectId}/history/{historyId}/canary/accept                    |
| _V1Api_ | [**RejectDeployCanary**](docs/V1Api.md#RejectDeployCanary)                         | **Post** /project/{projectId}/history/{historyId}/canary/reject                    |
| _V1Api_ | [**CancelDeploy**](docs/V1Api.md#CancelDeploy)                                     | **Post** /project/{projectId}/history/{historyId}/cancel                           |
| _V1Api_ | [**GetHistory**](docs/V1Api.md#GetHistory)                                         | **Get** /project/{projectId}/history/{historyId}                                   |
| _V1Api_ | [**GetCanaryReport**](docs/V1Api.md#GetCanaryReport)                               | **Get** /project/{projectId}/history/{historyId}/report/{endtime}                  |
| _V1Api_ | [**GetCanaryReportEndtime**](docs/V1Api.md#GetCanaryReportEndtime)                 | **Get** /project/{projectId}/history/{historyId}/report                            |
| _V1Api_ | [**GetStages**](docs/V1Api.md#GetStages)                                           | **Get** /project/{projectId}/stage                                                 |
| _V1Api_ | [**CreateStage**](docs/V1Api.md#CreateStage)                                       | **Post** /project/{projectId}/stage                                                |
| _V1Api_ | [**DeleteStage**](docs/V1Api.md#DeleteStage)                                       | **Delete** /project/{projectId}/stage/{stageId}                                    |
| _V1Api_ | [**GetStage**](docs/V1Api.md#GetStage)                                             | **Get** /project/{projectId}/stage/{stageId}                                       |
| _V1Api_ | [**ChangeStage**](docs/V1Api.md#ChangeStage)                                       | **Patch** /project/{projectId}/stage/{stageId}                                     |
| _V1Api_ | [**GetScenarioes**](docs/V1Api.md#GetScenarioes)                                   | **Get** /project/{projectId}/stage/{stageId}/scenario                              |
| _V1Api_ | [**CreateScenario**](docs/V1Api.md#CreateScenario)                                 | **Post** /project/{projectId}/stage/{stageId}/scenario                             |
| _V1Api_ | [**DeleteScenario**](docs/V1Api.md#DeleteScenario)                                 | **Delete** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}              |
| _V1Api_ | [**Deploy**](docs/V1Api.md#Deploy)                                                 | **Post** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}/deploy         |
| _V1Api_ | [**DeployRequest**](docs/V1Api.md#DeployRequest)                                   | **Post** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}/deploy/request |
| _V1Api_ | [**GetScenario**](docs/V1Api.md#GetScenario)                                       | **Get** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}                 |
| _V1Api_ | [**ChangeScenario**](docs/V1Api.md#ChangeScenario)                                 | **Patch** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}               |
| _V1Api_ | [**GetServers**](docs/V1Api.md#GetServers)                                         | **Get** /server                                                                    |
| _V1Api_ | [**GetSourcebuildProjects**](docs/V1Api.md#GetSourcebuildProjects)                 | **Get** /sourcebuild/project                                                       |
| _V1Api_ | [**GetSourcecommitRepositories**](docs/V1Api.md#GetSourcecommitRepositories)       | **Get** /sourcecommit/repository                                                   |
| _V1Api_ | [**GetSourceCommitBranches**](docs/V1Api.md#GetSourceCommitBranches)               | **Get** /sourcecommit/repository/{repositoryName}/branch                           |

## Documentation For Models

-   [ChangeScenario](docs/ChangeScenario.md)
-   [ChangeStage](docs/ChangeStage.md)
-   [CreateProject](docs/CreateProject.md)
-   [CreateScenario](docs/CreateScenario.md)
-   [CreateStage](docs/CreateStage.md)
-   [DeployRequestResponse](docs/DeployRequestResponse.md)
-   [GetAutoScalingGroupListResponse](docs/GetAutoScalingGroupListResponse.md)
-   [GetAutoScalingGroupListResponseAutoscalingList](docs/GetAutoScalingGroupListResponseAutoscalingList.md)
-   [GetCanaryAnalysisReportResponse](docs/GetCanaryAnalysisReportResponse.md)
-   [GetCanaryAnalysisStageListResponse](docs/GetCanaryAnalysisStageListResponse.md)
-   [GetDeployHistoryDetailResponse](docs/GetDeployHistoryDetailResponse.md)
-   [GetDeployHistoryDetailResponseConfig](docs/GetDeployHistoryDetailResponseConfig.md)
-   [GetDeployHistoryDetailResponseFile](docs/GetDeployHistoryDetailResponseFile.md)
-   [GetDeployHistoryDetailResponseServer](docs/GetDeployHistoryDetailResponseServer.md)
-   [GetDeployHistoryDetailResponseStep](docs/GetDeployHistoryDetailResponseStep.md)
-   [GetDeployHistoryDetailResponseTargets](docs/GetDeployHistoryDetailResponseTargets.md)
-   [GetDeployHistoryDetailResponseTime](docs/GetDeployHistoryDetailResponseTime.md)
-   [GetDeployHistoryListResponse](docs/GetDeployHistoryListResponse.md)
-   [GetDeployHistoryListResponseHistoryList](docs/GetDeployHistoryListResponseHistoryList.md)
-   [GetIdNameResponse](docs/GetIdNameResponse.md)
-   [GetKubernetesServiceClusterListResponse](docs/GetKubernetesServiceClusterListResponse.md)
-   [GetObjectStorageBucketListResponse](docs/GetObjectStorageBucketListResponse.md)
-   [GetObjectStorageObjectListResponse](docs/GetObjectStorageObjectListResponse.md)
-   [GetObjectStorageObjectListResponseObjectList](docs/GetObjectStorageObjectListResponseObjectList.md)
-   [GetProjectListResponse](docs/GetProjectListResponse.md)
-   [GetScenarioConfig](docs/GetScenarioConfig.md)
-   [GetScenarioConfigFile](docs/GetScenarioConfigFile.md)
-   [GetScenarioConfigLoadBalancer](docs/GetScenarioConfigLoadBalancer.md)
-   [GetScenarioDetailResponse](docs/GetScenarioDetailResponse.md)
-   [GetScenarioListResponse](docs/GetScenarioListResponse.md)
-   [GetServerListResponse](docs/GetServerListResponse.md)
-   [GetServerListResponseServerList](docs/GetServerListResponseServerList.md)
-   [GetSourceBuildProjectListResponse](docs/GetSourceBuildProjectListResponse.md)
-   [GetSourceCommitBranchListResponse](docs/GetSourceCommitBranchListResponse.md)
-   [GetSourceCommitRepositoryListResponse](docs/GetSourceCommitRepositoryListResponse.md)
-   [GetSourceCommitRepositoryListResponseRepositoryList](docs/GetSourceCommitRepositoryListResponseRepositoryList.md)
-   [GetStageDetailResponse](docs/GetStageDetailResponse.md)
-   [GetStageDetailResponseConfig](docs/GetStageDetailResponseConfig.md)
-   [GetStageListResponse](docs/GetStageListResponse.md)
-   [GetTargetGroupListResponse](docs/GetTargetGroupListResponse.md)
-   [OkResponse](docs/OkResponse.md)
-   [ScenarioConfig](docs/ScenarioConfig.md)
-   [ScenarioConfigCanaryConfig](docs/ScenarioConfigCanaryConfig.md)
-   [ScenarioConfigCanaryConfigAnalysisConfig](docs/ScenarioConfigCanaryConfigAnalysisConfig.md)
-   [ScenarioConfigCanaryConfigEnv](docs/ScenarioConfigCanaryConfigEnv.md)
-   [ScenarioConfigCanaryConfigMetrics](docs/ScenarioConfigCanaryConfigMetrics.md)
-   [ScenarioConfigCmd](docs/ScenarioConfigCmd.md)
-   [ScenarioConfigCmdDeploy](docs/ScenarioConfigCmdDeploy.md)
-   [ScenarioConfigCmdPrePost](docs/ScenarioConfigCmdPrePost.md)
-   [ScenarioConfigFile](docs/ScenarioConfigFile.md)
-   [ScenarioConfigFileObjectStorage](docs/ScenarioConfigFileObjectStorage.md)
-   [ScenarioConfigFileSourceBuild](docs/ScenarioConfigFileSourceBuild.md)
-   [ScenarioConfigLoadBalancer](docs/ScenarioConfigLoadBalancer.md)
-   [ScenarioConfigManifest](docs/ScenarioConfigManifest.md)
-   [StageConfig](docs/StageConfig.md)