/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v2)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses2

type SearchEngineVersion struct {
	SearchEngineDashboardVersionName string `json:"searchEngineDashboardVersionName,omitempty"`
	SearchEngineVersionCode          string `json:"searchEngineVersionCode,omitempty"`
	SearchEngineVersionName          string `json:"searchEngineVersionName,omitempty"`
	Type_                            string `json:"type,omitempty"`
}
