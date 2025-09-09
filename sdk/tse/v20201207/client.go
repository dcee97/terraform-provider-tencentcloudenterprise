// All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20201207

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-12-07"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewUpdateInstanceSpecForApolloRequest() (request *UpdateInstanceSpecForApolloRequest) {
	request = &UpdateInstanceSpecForApolloRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateInstanceSpecForApollo")
	return
}

func NewUpdateInstanceSpecForApolloResponse() (response *UpdateInstanceSpecForApolloResponse) {
	response = &UpdateInstanceSpecForApolloResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Apollo实例某环境的规格和节点数
func (c *Client) UpdateInstanceSpecForApollo(request *UpdateInstanceSpecForApolloRequest) (response *UpdateInstanceSpecForApolloResponse, err error) {
	if request == nil {
		request = NewUpdateInstanceSpecForApolloRequest()
	}
	response = NewUpdateInstanceSpecForApolloResponse()
	err = c.Send(request, response)
	return
}

func NewUploadCloudNativeAPIGatewayPluginRequest() (request *UploadCloudNativeAPIGatewayPluginRequest) {
	request = &UploadCloudNativeAPIGatewayPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UploadCloudNativeAPIGatewayPlugin")
	return
}

func NewUploadCloudNativeAPIGatewayPluginResponse() (response *UploadCloudNativeAPIGatewayPluginResponse) {
	response = &UploadCloudNativeAPIGatewayPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上传云原生网关插件
func (c *Client) UploadCloudNativeAPIGatewayPlugin(request *UploadCloudNativeAPIGatewayPluginRequest) (response *UploadCloudNativeAPIGatewayPluginResponse, err error) {
	if request == nil {
		request = NewUploadCloudNativeAPIGatewayPluginRequest()
	}
	response = NewUploadCloudNativeAPIGatewayPluginResponse()
	err = c.Send(request, response)
	return
}

func NewManageConsulConsoleRequest() (request *ManageConsulConsoleRequest) {
	request = &ManageConsulConsoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ManageConsulConsole")
	return
}

func NewManageConsulConsoleResponse() (response *ManageConsulConsoleResponse) {
	response = &ManageConsulConsoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理Consul控制台接口
func (c *Client) ManageConsulConsole(request *ManageConsulConsoleRequest) (response *ManageConsulConsoleResponse, err error) {
	if request == nil {
		request = NewManageConsulConsoleRequest()
	}
	response = NewManageConsulConsoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExecTaskBatchJobInfoRequest() (request *DescribeExecTaskBatchJobInfoRequest) {
	request = &DescribeExecTaskBatchJobInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeExecTaskBatchJobInfo")
	return
}

func NewDescribeExecTaskBatchJobInfoResponse() (response *DescribeExecTaskBatchJobInfoResponse) {
	response = &DescribeExecTaskBatchJobInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看运营平台可执行任务批次信息
func (c *Client) DescribeExecTaskBatchJobInfo(request *DescribeExecTaskBatchJobInfoRequest) (response *DescribeExecTaskBatchJobInfoResponse, err error) {
	if request == nil {
		request = NewDescribeExecTaskBatchJobInfoRequest()
	}
	response = NewDescribeExecTaskBatchJobInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceUserLoginInfoRequest() (request *DescribeGovernanceUserLoginInfoRequest) {
	request = &DescribeGovernanceUserLoginInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceUserLoginInfo")
	return
}

func NewDescribeGovernanceUserLoginInfoResponse() (response *DescribeGovernanceUserLoginInfoResponse) {
	response = &DescribeGovernanceUserLoginInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心引擎的用户登录名称
func (c *Client) DescribeGovernanceUserLoginInfo(request *DescribeGovernanceUserLoginInfoRequest) (response *DescribeGovernanceUserLoginInfoResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceUserLoginInfoRequest()
	}
	response = NewDescribeGovernanceUserLoginInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperDigestTokenRequest() (request *DescribeZookeeperDigestTokenRequest) {
	request = &DescribeZookeeperDigestTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperDigestToken")
	return
}

func NewDescribeZookeeperDigestTokenResponse() (response *DescribeZookeeperDigestTokenResponse) {
	response = &DescribeZookeeperDigestTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看超级管理员摘要
func (c *Client) DescribeZookeeperDigestToken(request *DescribeZookeeperDigestTokenRequest) (response *DescribeZookeeperDigestTokenResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperDigestTokenRequest()
	}
	response = NewDescribeZookeeperDigestTokenResponse()
	err = c.Send(request, response)
	return
}

func NewGetCloudNativeAPIGatewayDomainsRequest() (request *GetCloudNativeAPIGatewayDomainsRequest) {
	request = &GetCloudNativeAPIGatewayDomainsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "GetCloudNativeAPIGatewayDomains")
	return
}

func NewGetCloudNativeAPIGatewayDomainsResponse() (response *GetCloudNativeAPIGatewayDomainsResponse) {
	response = &GetCloudNativeAPIGatewayDomainsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Nginx Ingress 域名信息
func (c *Client) GetCloudNativeAPIGatewayDomains(request *GetCloudNativeAPIGatewayDomainsRequest) (response *GetCloudNativeAPIGatewayDomainsResponse, err error) {
	if request == nil {
		request = NewGetCloudNativeAPIGatewayDomainsRequest()
	}
	response = NewGetCloudNativeAPIGatewayDomainsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayIngressRequest() (request *DescribeCloudNativeAPIGatewayIngressRequest) {
	request = &DescribeCloudNativeAPIGatewayIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayIngress")
	return
}

func NewDescribeCloudNativeAPIGatewayIngressResponse() (response *DescribeCloudNativeAPIGatewayIngressResponse) {
	response = &DescribeCloudNativeAPIGatewayIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Ingress详细信息
func (c *Client) DescribeCloudNativeAPIGatewayIngress(request *DescribeCloudNativeAPIGatewayIngressRequest) (response *DescribeCloudNativeAPIGatewayIngressResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayIngressRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayIngressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceTagInfosRequest() (request *DescribeInstanceTagInfosRequest) {
	request = &DescribeInstanceTagInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeInstanceTagInfos")
	return
}

func NewDescribeInstanceTagInfosResponse() (response *DescribeInstanceTagInfosResponse) {
	response = &DescribeInstanceTagInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看实例的标签信息
func (c *Client) DescribeInstanceTagInfos(request *DescribeInstanceTagInfosRequest) (response *DescribeInstanceTagInfosResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceTagInfosRequest()
	}
	response = NewDescribeInstanceTagInfosResponse()
	err = c.Send(request, response)
	return
}

func NewManageConfigRequest() (request *ManageConfigRequest) {
	request = &ManageConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ManageConfig")
	return
}

func NewManageConfigResponse() (response *ManageConfigResponse) {
	response = &ManageConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理配置
func (c *Client) ManageConfig(request *ManageConfigRequest) (response *ManageConfigResponse, err error) {
	if request == nil {
		request = NewManageConfigRequest()
	}
	response = NewManageConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusRequest() (request *DescribePrometheusRequest) {
	request = &DescribePrometheusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribePrometheus")
	return
}

func NewDescribePrometheusResponse() (response *DescribePrometheusResponse) {
	response = &DescribePrometheusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Prometheus指标推送状态
func (c *Client) DescribePrometheus(request *DescribePrometheusRequest) (response *DescribePrometheusResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusRequest()
	}
	response = NewDescribePrometheusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceStatusListRequest() (request *DescribeServiceStatusListRequest) {
	request = &DescribeServiceStatusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeServiceStatusList")
	return
}

func NewDescribeServiceStatusListResponse() (response *DescribeServiceStatusListResponse) {
	response = &DescribeServiceStatusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询注册服务状态列表配置信息
func (c *Client) DescribeServiceStatusList(request *DescribeServiceStatusListRequest) (response *DescribeServiceStatusListResponse, err error) {
	if request == nil {
		request = NewDescribeServiceStatusListRequest()
	}
	response = NewDescribeServiceStatusListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableZoneRequest() (request *DescribeAvailableZoneRequest) {
	request = &DescribeAvailableZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeAvailableZone")
	return
}

func NewDescribeAvailableZoneResponse() (response *DescribeAvailableZoneResponse) {
	response = &DescribeAvailableZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询地域可用区
func (c *Client) DescribeAvailableZone(request *DescribeAvailableZoneRequest) (response *DescribeAvailableZoneResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableZoneRequest()
	}
	response = NewDescribeAvailableZoneResponse()
	err = c.Send(request, response)
	return
}

func NewRestartSREInstanceRequest() (request *RestartSREInstanceRequest) {
	request = &RestartSREInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "RestartSREInstance")
	return
}

func NewRestartSREInstanceResponse() (response *RestartSREInstanceResponse) {
	response = &RestartSREInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重启微服务引擎实例
func (c *Client) RestartSREInstance(request *RestartSREInstanceRequest) (response *RestartSREInstanceResponse, err error) {
	if request == nil {
		request = NewRestartSREInstanceRequest()
	}
	response = NewRestartSREInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNativeGatewayServiceSourceRequest() (request *CreateNativeGatewayServiceSourceRequest) {
	request = &CreateNativeGatewayServiceSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateNativeGatewayServiceSource")
	return
}

func NewCreateNativeGatewayServiceSourceResponse() (response *CreateNativeGatewayServiceSourceResponse) {
	response = &CreateNativeGatewayServiceSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网关服务来源
func (c *Client) CreateNativeGatewayServiceSource(request *CreateNativeGatewayServiceSourceRequest) (response *CreateNativeGatewayServiceSourceResponse, err error) {
	if request == nil {
		request = NewCreateNativeGatewayServiceSourceRequest()
	}
	response = NewCreateNativeGatewayServiceSourceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewaySourceRequest() (request *DeleteCloudNativeAPIGatewaySourceRequest) {
	request = &DeleteCloudNativeAPIGatewaySourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewaySource")
	return
}

func NewDeleteCloudNativeAPIGatewaySourceResponse() (response *DeleteCloudNativeAPIGatewaySourceResponse) {
	response = &DeleteCloudNativeAPIGatewaySourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除云原生网关的服务来源
func (c *Client) DeleteCloudNativeAPIGatewaySource(request *DeleteCloudNativeAPIGatewaySourceRequest) (response *DeleteCloudNativeAPIGatewaySourceResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewaySourceRequest()
	}
	response = NewDeleteCloudNativeAPIGatewaySourceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEngineRequest() (request *DeleteEngineRequest) {
	request = &DeleteEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteEngine")
	return
}

func NewDeleteEngineResponse() (response *DeleteEngineResponse) {
	response = &DeleteEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除引擎实例
func (c *Client) DeleteEngine(request *DeleteEngineRequest) (response *DeleteEngineResponse, err error) {
	if request == nil {
		request = NewDeleteEngineRequest()
	}
	response = NewDeleteEngineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperConfigRequest() (request *DescribeZookeeperConfigRequest) {
	request = &DescribeZookeeperConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperConfig")
	return
}

func NewDescribeZookeeperConfigResponse() (response *DescribeZookeeperConfigResponse) {
	response = &DescribeZookeeperConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看zookeeper配置项
func (c *Client) DescribeZookeeperConfig(request *DescribeZookeeperConfigRequest) (response *DescribeZookeeperConfigResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperConfigRequest()
	}
	response = NewDescribeZookeeperConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionInfosRequest() (request *DescribeRegionInfosRequest) {
	request = &DescribeRegionInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeRegionInfos")
	return
}

func NewDescribeRegionInfosResponse() (response *DescribeRegionInfosResponse) {
	response = &DescribeRegionInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询区域相关信息
func (c *Client) DescribeRegionInfos(request *DescribeRegionInfosRequest) (response *DescribeRegionInfosResponse, err error) {
	if request == nil {
		request = NewDescribeRegionInfosRequest()
	}
	response = NewDescribeRegionInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigFileGroupsRequest() (request *DescribeConfigFileGroupsRequest) {
	request = &DescribeConfigFileGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFileGroups")
	return
}

func NewDescribeConfigFileGroupsResponse() (response *DescribeConfigFileGroupsResponse) {
	response = &DescribeConfigFileGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件分页查询配置文件组
func (c *Client) DescribeConfigFileGroups(request *DescribeConfigFileGroupsRequest) (response *DescribeConfigFileGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeConfigFileGroupsRequest()
	}
	response = NewDescribeConfigFileGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateEngineSpecListRequest() (request *DescribeOperateEngineSpecListRequest) {
	request = &DescribeOperateEngineSpecListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateEngineSpecList")
	return
}

func NewDescribeOperateEngineSpecListResponse() (response *DescribeOperateEngineSpecListResponse) {
	response = &DescribeOperateEngineSpecListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端引擎规格列表
func (c *Client) DescribeOperateEngineSpecList(request *DescribeOperateEngineSpecListRequest) (response *DescribeOperateEngineSpecListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateEngineSpecListRequest()
	}
	response = NewDescribeOperateEngineSpecListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceGroupRequest() (request *ModifyGovernanceGroupRequest) {
	request = &ModifyGovernanceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceGroup")
	return
}

func NewModifyGovernanceGroupResponse() (response *ModifyGovernanceGroupResponse) {
	response = &ModifyGovernanceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改治理中心用户组信息
func (c *Client) ModifyGovernanceGroup(request *ModifyGovernanceGroupRequest) (response *ModifyGovernanceGroupResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceGroupRequest()
	}
	response = NewModifyGovernanceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperMigrateClusterInfoRequest() (request *DescribeZookeeperMigrateClusterInfoRequest) {
	request = &DescribeZookeeperMigrateClusterInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperMigrateClusterInfo")
	return
}

func NewDescribeZookeeperMigrateClusterInfoResponse() (response *DescribeZookeeperMigrateClusterInfoResponse) {
	response = &DescribeZookeeperMigrateClusterInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看迁移集群信息
func (c *Client) DescribeZookeeperMigrateClusterInfo(request *DescribeZookeeperMigrateClusterInfoRequest) (response *DescribeZookeeperMigrateClusterInfoResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperMigrateClusterInfoRequest()
	}
	response = NewDescribeZookeeperMigrateClusterInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEngineSyncJobRequest() (request *CreateEngineSyncJobRequest) {
	request = &CreateEngineSyncJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateEngineSyncJob")
	return
}

func NewCreateEngineSyncJobResponse() (response *CreateEngineSyncJobResponse) {
	response = &CreateEngineSyncJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建引擎数据同步任务
func (c *Client) CreateEngineSyncJob(request *CreateEngineSyncJobRequest) (response *CreateEngineSyncJobResponse, err error) {
	if request == nil {
		request = NewCreateEngineSyncJobRequest()
	}
	response = NewCreateEngineSyncJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayCodeNamesRequest() (request *DescribeCloudNativeAPIGatewayCodeNamesRequest) {
	request = &DescribeCloudNativeAPIGatewayCodeNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayCodeNames")
	return
}

func NewDescribeCloudNativeAPIGatewayCodeNamesResponse() (response *DescribeCloudNativeAPIGatewayCodeNamesResponse) {
	response = &DescribeCloudNativeAPIGatewayCodeNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云原生网关文案名称列表
func (c *Client) DescribeCloudNativeAPIGatewayCodeNames(request *DescribeCloudNativeAPIGatewayCodeNamesRequest) (response *DescribeCloudNativeAPIGatewayCodeNamesResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayCodeNamesRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayCodeNamesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudNativeAPIGatewayRequest() (request *ModifyCloudNativeAPIGatewayRequest) {
	request = &ModifyCloudNativeAPIGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGateway")
	return
}

func NewModifyCloudNativeAPIGatewayResponse() (response *ModifyCloudNativeAPIGatewayResponse) {
	response = &ModifyCloudNativeAPIGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生API网关实例基础信息
func (c *Client) ModifyCloudNativeAPIGateway(request *ModifyCloudNativeAPIGatewayRequest) (response *ModifyCloudNativeAPIGatewayResponse, err error) {
	if request == nil {
		request = NewModifyCloudNativeAPIGatewayRequest()
	}
	response = NewModifyCloudNativeAPIGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateInstanceRequest() (request *ModifyOperateInstanceRequest) {
	request = &ModifyOperateInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyOperateInstance")
	return
}

func NewModifyOperateInstanceResponse() (response *ModifyOperateInstanceResponse) {
	response = &ModifyOperateInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新运营平台实例信息
func (c *Client) ModifyOperateInstance(request *ModifyOperateInstanceRequest) (response *ModifyOperateInstanceResponse, err error) {
	if request == nil {
		request = NewModifyOperateInstanceRequest()
	}
	response = NewModifyOperateInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicensesRequest() (request *DescribeLicensesRequest) {
	request = &DescribeLicensesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeLicenses")
	return
}

func NewDescribeLicensesResponse() (response *DescribeLicensesResponse) {
	response = &DescribeLicensesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询License列表
func (c *Client) DescribeLicenses(request *DescribeLicensesRequest) (response *DescribeLicensesResponse, err error) {
	if request == nil {
		request = NewDescribeLicensesRequest()
	}
	response = NewDescribeLicensesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsulServerInterfacesRequest() (request *DescribeConsulServerInterfacesRequest) {
	request = &DescribeConsulServerInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConsulServerInterfaces")
	return
}

func NewDescribeConsulServerInterfacesResponse() (response *DescribeConsulServerInterfacesResponse) {
	response = &DescribeConsulServerInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询consul服务接口列表
func (c *Client) DescribeConsulServerInterfaces(request *DescribeConsulServerInterfacesRequest) (response *DescribeConsulServerInterfacesResponse, err error) {
	if request == nil {
		request = NewDescribeConsulServerInterfacesRequest()
	}
	response = NewDescribeConsulServerInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNacosMonitorDashboardRequest() (request *DescribeNacosMonitorDashboardRequest) {
	request = &DescribeNacosMonitorDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeNacosMonitorDashboard")
	return
}

func NewDescribeNacosMonitorDashboardResponse() (response *DescribeNacosMonitorDashboardResponse) {
	response = &DescribeNacosMonitorDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 nacos大盘指标
func (c *Client) DescribeNacosMonitorDashboard(request *DescribeNacosMonitorDashboardRequest) (response *DescribeNacosMonitorDashboardResponse, err error) {
	if request == nil {
		request = NewDescribeNacosMonitorDashboardRequest()
	}
	response = NewDescribeNacosMonitorDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewGetZookeeperMigrateLoadDataStatusRequest() (request *GetZookeeperMigrateLoadDataStatusRequest) {
	request = &GetZookeeperMigrateLoadDataStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "GetZookeeperMigrateLoadDataStatus")
	return
}

func NewGetZookeeperMigrateLoadDataStatusResponse() (response *GetZookeeperMigrateLoadDataStatusResponse) {
	response = &GetZookeeperMigrateLoadDataStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看数据导入状态
func (c *Client) GetZookeeperMigrateLoadDataStatus(request *GetZookeeperMigrateLoadDataStatusRequest) (response *GetZookeeperMigrateLoadDataStatusResponse, err error) {
	if request == nil {
		request = NewGetZookeeperMigrateLoadDataStatusRequest()
	}
	response = NewGetZookeeperMigrateLoadDataStatusResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateZooKeeperConfigInfoRequest() (request *UpdateZooKeeperConfigInfoRequest) {
	request = &UpdateZooKeeperConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateZooKeeperConfigInfo")
	return
}

func NewUpdateZooKeeperConfigInfoResponse() (response *UpdateZooKeeperConfigInfoResponse) {
	response = &UpdateZooKeeperConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新zookeeper系统参数
func (c *Client) UpdateZooKeeperConfigInfo(request *UpdateZooKeeperConfigInfoRequest) (response *UpdateZooKeeperConfigInfoResponse, err error) {
	if request == nil {
		request = NewUpdateZooKeeperConfigInfoRequest()
	}
	response = NewUpdateZooKeeperConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigFilesRequest() (request *DeleteConfigFilesRequest) {
	request = &DeleteConfigFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteConfigFiles")
	return
}

func NewDeleteConfigFilesResponse() (response *DeleteConfigFilesResponse) {
	response = &DeleteConfigFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除配置文件
func (c *Client) DeleteConfigFiles(request *DeleteConfigFilesRequest) (response *DeleteConfigFilesResponse, err error) {
	if request == nil {
		request = NewDeleteConfigFilesRequest()
	}
	response = NewDeleteConfigFilesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperateFeatureVersionRequest() (request *DeleteOperateFeatureVersionRequest) {
	request = &DeleteOperateFeatureVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteOperateFeatureVersion")
	return
}

func NewDeleteOperateFeatureVersionResponse() (response *DeleteOperateFeatureVersionResponse) {
	response = &DeleteOperateFeatureVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除功能版本
func (c *Client) DeleteOperateFeatureVersion(request *DeleteOperateFeatureVersionRequest) (response *DeleteOperateFeatureVersionResponse, err error) {
	if request == nil {
		request = NewDeleteOperateFeatureVersionRequest()
	}
	response = NewDeleteOperateFeatureVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceStrategyDetailRequest() (request *DescribeGovernanceStrategyDetailRequest) {
	request = &DescribeGovernanceStrategyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceStrategyDetail")
	return
}

func NewDescribeGovernanceStrategyDetailResponse() (response *DescribeGovernanceStrategyDetailResponse) {
	response = &DescribeGovernanceStrategyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心鉴权策略详细
func (c *Client) DescribeGovernanceStrategyDetail(request *DescribeGovernanceStrategyDetailRequest) (response *DescribeGovernanceStrategyDetailResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceStrategyDetailRequest()
	}
	response = NewDescribeGovernanceStrategyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConfigFileGroupRequest() (request *ModifyConfigFileGroupRequest) {
	request = &ModifyConfigFileGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyConfigFileGroup")
	return
}

func NewModifyConfigFileGroupResponse() (response *ModifyConfigFileGroupResponse) {
	response = &ModifyConfigFileGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件批量修改配置文件组
func (c *Client) ModifyConfigFileGroup(request *ModifyConfigFileGroupRequest) (response *ModifyConfigFileGroupResponse, err error) {
	if request == nil {
		request = NewModifyConfigFileGroupRequest()
	}
	response = NewModifyConfigFileGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEngineSpecRequest() (request *ModifyEngineSpecRequest) {
	request = &ModifyEngineSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyEngineSpec")
	return
}

func NewModifyEngineSpecResponse() (response *ModifyEngineSpecResponse) {
	response = &ModifyEngineSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改引擎实例的规格或节点数
func (c *Client) ModifyEngineSpec(request *ModifyEngineSpecRequest) (response *ModifyEngineSpecResponse, err error) {
	if request == nil {
		request = NewModifyEngineSpecRequest()
	}
	response = NewModifyEngineSpecResponse()
	err = c.Send(request, response)
	return
}

func NewCheckZookeeperMigrateTargetStatusRequest() (request *CheckZookeeperMigrateTargetStatusRequest) {
	request = &CheckZookeeperMigrateTargetStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CheckZookeeperMigrateTargetStatus")
	return
}

func NewCheckZookeeperMigrateTargetStatusResponse() (response *CheckZookeeperMigrateTargetStatusResponse) {
	response = &CheckZookeeperMigrateTargetStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 迁移准备阶段检查集群信息
func (c *Client) CheckZookeeperMigrateTargetStatus(request *CheckZookeeperMigrateTargetStatusRequest) (response *CheckZookeeperMigrateTargetStatusResponse, err error) {
	if request == nil {
		request = NewCheckZookeeperMigrateTargetStatusRequest()
	}
	response = NewCheckZookeeperMigrateTargetStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceServicesRequest() (request *DescribeGovernanceServicesRequest) {
	request = &DescribeGovernanceServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceServices")
	return
}

func NewDescribeGovernanceServicesResponse() (response *DescribeGovernanceServicesResponse) {
	response = &DescribeGovernanceServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心服务列表
func (c *Client) DescribeGovernanceServices(request *DescribeGovernanceServicesRequest) (response *DescribeGovernanceServicesResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceServicesRequest()
	}
	response = NewDescribeGovernanceServicesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApolloEnvInfoRequest() (request *UpdateApolloEnvInfoRequest) {
	request = &UpdateApolloEnvInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateApolloEnvInfo")
	return
}

func NewUpdateApolloEnvInfoResponse() (response *UpdateApolloEnvInfoResponse) {
	response = &UpdateApolloEnvInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Apollo环境信息
func (c *Client) UpdateApolloEnvInfo(request *UpdateApolloEnvInfoRequest) (response *UpdateApolloEnvInfoResponse, err error) {
	if request == nil {
		request = NewUpdateApolloEnvInfoRequest()
	}
	response = NewUpdateApolloEnvInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperateWhiteListRequest() (request *CreateOperateWhiteListRequest) {
	request = &CreateOperateWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateOperateWhiteList")
	return
}

func NewCreateOperateWhiteListResponse() (response *CreateOperateWhiteListResponse) {
	response = &CreateOperateWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运营平台白名单
func (c *Client) CreateOperateWhiteList(request *CreateOperateWhiteListRequest) (response *CreateOperateWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateOperateWhiteListRequest()
	}
	response = NewCreateOperateWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperMigrateMergeStatusRequest() (request *DescribeZookeeperMigrateMergeStatusRequest) {
	request = &DescribeZookeeperMigrateMergeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperMigrateMergeStatus")
	return
}

func NewDescribeZookeeperMigrateMergeStatusResponse() (response *DescribeZookeeperMigrateMergeStatusResponse) {
	response = &DescribeZookeeperMigrateMergeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询合并状态
func (c *Client) DescribeZookeeperMigrateMergeStatus(request *DescribeZookeeperMigrateMergeStatusRequest) (response *DescribeZookeeperMigrateMergeStatusResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperMigrateMergeStatusRequest()
	}
	response = NewDescribeZookeeperMigrateMergeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSREInstanceSubnetRequest() (request *UpdateSREInstanceSubnetRequest) {
	request = &UpdateSREInstanceSubnetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateSREInstanceSubnet")
	return
}

func NewUpdateSREInstanceSubnetResponse() (response *UpdateSREInstanceSubnetResponse) {
	response = &UpdateSREInstanceSubnetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新实例部署子网信息
func (c *Client) UpdateSREInstanceSubnet(request *UpdateSREInstanceSubnetRequest) (response *UpdateSREInstanceSubnetResponse, err error) {
	if request == nil {
		request = NewUpdateSREInstanceSubnetRequest()
	}
	response = NewUpdateSREInstanceSubnetResponse()
	err = c.Send(request, response)
	return
}

func NewGetZookeeperMigrateUploadMetaInfoRequest() (request *GetZookeeperMigrateUploadMetaInfoRequest) {
	request = &GetZookeeperMigrateUploadMetaInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "GetZookeeperMigrateUploadMetaInfo")
	return
}

func NewGetZookeeperMigrateUploadMetaInfoResponse() (response *GetZookeeperMigrateUploadMetaInfoResponse) {
	response = &GetZookeeperMigrateUploadMetaInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取cos信息
func (c *Client) GetZookeeperMigrateUploadMetaInfo(request *GetZookeeperMigrateUploadMetaInfoRequest) (response *GetZookeeperMigrateUploadMetaInfoResponse, err error) {
	if request == nil {
		request = NewGetZookeeperMigrateUploadMetaInfoRequest()
	}
	response = NewGetZookeeperMigrateUploadMetaInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayBasicLogRequest() (request *DescribeCloudNativeAPIGatewayBasicLogRequest) {
	request = &DescribeCloudNativeAPIGatewayBasicLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayBasicLog")
	return
}

func NewDescribeCloudNativeAPIGatewayBasicLogResponse() (response *DescribeCloudNativeAPIGatewayBasicLogResponse) {
	response = &DescribeCloudNativeAPIGatewayBasicLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云原生网关日志。
func (c *Client) DescribeCloudNativeAPIGatewayBasicLog(request *DescribeCloudNativeAPIGatewayBasicLogRequest) (response *DescribeCloudNativeAPIGatewayBasicLogResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayBasicLogRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayBasicLogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperateEngineVersionRequest() (request *CreateOperateEngineVersionRequest) {
	request = &CreateOperateEngineVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateOperateEngineVersion")
	return
}

func NewCreateOperateEngineVersionResponse() (response *CreateOperateEngineVersionResponse) {
	response = &CreateOperateEngineVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运营端引擎版本
func (c *Client) CreateOperateEngineVersion(request *CreateOperateEngineVersionRequest) (response *CreateOperateEngineVersionResponse, err error) {
	if request == nil {
		request = NewCreateOperateEngineVersionRequest()
	}
	response = NewCreateOperateEngineVersionResponse()
	err = c.Send(request, response)
	return
}

func NewEnableInternetRequest() (request *EnableInternetRequest) {
	request = &EnableInternetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "EnableInternet")
	return
}

func NewEnableInternetResponse() (response *EnableInternetResponse) {
	response = &EnableInternetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群公网开关
func (c *Client) EnableInternet(request *EnableInternetRequest) (response *EnableInternetResponse, err error) {
	if request == nil {
		request = NewEnableInternetRequest()
	}
	response = NewEnableInternetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudNativeAPIGatewayServiceRequest() (request *ModifyCloudNativeAPIGatewayServiceRequest) {
	request = &ModifyCloudNativeAPIGatewayServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayService")
	return
}

func NewModifyCloudNativeAPIGatewayServiceResponse() (response *ModifyCloudNativeAPIGatewayServiceResponse) {
	response = &ModifyCloudNativeAPIGatewayServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生网关服务
func (c *Client) ModifyCloudNativeAPIGatewayService(request *ModifyCloudNativeAPIGatewayServiceRequest) (response *ModifyCloudNativeAPIGatewayServiceResponse, err error) {
	if request == nil {
		request = NewModifyCloudNativeAPIGatewayServiceRequest()
	}
	response = NewModifyCloudNativeAPIGatewayServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceLogRequest() (request *DescribeInstanceLogRequest) {
	request = &DescribeInstanceLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeInstanceLog")
	return
}

func NewDescribeInstanceLogResponse() (response *DescribeInstanceLogResponse) {
	response = &DescribeInstanceLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例日志信息
func (c *Client) DescribeInstanceLog(request *DescribeInstanceLogRequest) (response *DescribeInstanceLogResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceLogRequest()
	}
	response = NewDescribeInstanceLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayClsConfigRequest() (request *DescribeCloudNativeAPIGatewayClsConfigRequest) {
	request = &DescribeCloudNativeAPIGatewayClsConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayClsConfig")
	return
}

func NewDescribeCloudNativeAPIGatewayClsConfigResponse() (response *DescribeCloudNativeAPIGatewayClsConfigResponse) {
	response = &DescribeCloudNativeAPIGatewayClsConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关实例关联的CLS配置信息
func (c *Client) DescribeCloudNativeAPIGatewayClsConfig(request *DescribeCloudNativeAPIGatewayClsConfigRequest) (response *DescribeCloudNativeAPIGatewayClsConfigResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayClsConfigRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayClsConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEngineIntranetAccessAddressListRequest() (request *DescribeEngineIntranetAccessAddressListRequest) {
	request = &DescribeEngineIntranetAccessAddressListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEngineIntranetAccessAddressList")
	return
}

func NewDescribeEngineIntranetAccessAddressListResponse() (response *DescribeEngineIntranetAccessAddressListResponse) {
	response = &DescribeEngineIntranetAccessAddressListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询引擎实例网络打通地址列表
func (c *Client) DescribeEngineIntranetAccessAddressList(request *DescribeEngineIntranetAccessAddressListRequest) (response *DescribeEngineIntranetAccessAddressListResponse, err error) {
	if request == nil {
		request = NewDescribeEngineIntranetAccessAddressListRequest()
	}
	response = NewDescribeEngineIntranetAccessAddressListResponse()
	err = c.Send(request, response)
	return
}

func NewStartEngineSyncJobRequest() (request *StartEngineSyncJobRequest) {
	request = &StartEngineSyncJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "StartEngineSyncJob")
	return
}

func NewStartEngineSyncJobResponse() (response *StartEngineSyncJobResponse) {
	response = &StartEngineSyncJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动引擎数据同步任务
func (c *Client) StartEngineSyncJob(request *StartEngineSyncJobRequest) (response *StartEngineSyncJobResponse, err error) {
	if request == nil {
		request = NewStartEngineSyncJobRequest()
	}
	response = NewStartEngineSyncJobResponse()
	err = c.Send(request, response)
	return
}

func NewStopEngineSyncJobRequest() (request *StopEngineSyncJobRequest) {
	request = &StopEngineSyncJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "StopEngineSyncJob")
	return
}

func NewStopEngineSyncJobResponse() (response *StopEngineSyncJobResponse) {
	response = &StopEngineSyncJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 终止引擎数据同步任务
func (c *Client) StopEngineSyncJob(request *StopEngineSyncJobRequest) (response *StopEngineSyncJobResponse, err error) {
	if request == nil {
		request = NewStopEngineSyncJobRequest()
	}
	response = NewStopEngineSyncJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEurekaReplicasRequest() (request *DescribeEurekaReplicasRequest) {
	request = &DescribeEurekaReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEurekaReplicas")
	return
}

func NewDescribeEurekaReplicasResponse() (response *DescribeEurekaReplicasResponse) {
	response = &DescribeEurekaReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Eureka类型注册引擎实例副本信息
func (c *Client) DescribeEurekaReplicas(request *DescribeEurekaReplicasRequest) (response *DescribeEurekaReplicasResponse, err error) {
	if request == nil {
		request = NewDescribeEurekaReplicasRequest()
	}
	response = NewDescribeEurekaReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewayPublicNetworkRequest() (request *DeleteCloudNativeAPIGatewayPublicNetworkRequest) {
	request = &DeleteCloudNativeAPIGatewayPublicNetworkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayPublicNetwork")
	return
}

func NewDeleteCloudNativeAPIGatewayPublicNetworkResponse() (response *DeleteCloudNativeAPIGatewayPublicNetworkResponse) {
	response = &DeleteCloudNativeAPIGatewayPublicNetworkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除公网网络配置
func (c *Client) DeleteCloudNativeAPIGatewayPublicNetwork(request *DeleteCloudNativeAPIGatewayPublicNetworkRequest) (response *DeleteCloudNativeAPIGatewayPublicNetworkResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewayPublicNetworkRequest()
	}
	response = NewDeleteCloudNativeAPIGatewayPublicNetworkResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateEngineVersionRequest() (request *ModifyOperateEngineVersionRequest) {
	request = &ModifyOperateEngineVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyOperateEngineVersion")
	return
}

func NewModifyOperateEngineVersionResponse() (response *ModifyOperateEngineVersionResponse) {
	response = &ModifyOperateEngineVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改运营端引擎版本
func (c *Client) ModifyOperateEngineVersion(request *ModifyOperateEngineVersionRequest) (response *ModifyOperateEngineVersionResponse, err error) {
	if request == nil {
		request = NewModifyOperateEngineVersionRequest()
	}
	response = NewModifyOperateEngineVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewayPluginRequest() (request *DeleteCloudNativeAPIGatewayPluginRequest) {
	request = &DeleteCloudNativeAPIGatewayPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayPlugin")
	return
}

func NewDeleteCloudNativeAPIGatewayPluginResponse() (response *DeleteCloudNativeAPIGatewayPluginResponse) {
	response = &DeleteCloudNativeAPIGatewayPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除云原生API网关插件
func (c *Client) DeleteCloudNativeAPIGatewayPlugin(request *DeleteCloudNativeAPIGatewayPluginRequest) (response *DeleteCloudNativeAPIGatewayPluginResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewayPluginRequest()
	}
	response = NewDeleteCloudNativeAPIGatewayPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSREInstanceHealthStatusRequest() (request *DescribeSREInstanceHealthStatusRequest) {
	request = &DescribeSREInstanceHealthStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeSREInstanceHealthStatus")
	return
}

func NewDescribeSREInstanceHealthStatusResponse() (response *DescribeSREInstanceHealthStatusResponse) {
	response = &DescribeSREInstanceHealthStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询微服务注册引擎实例健康状态
func (c *Client) DescribeSREInstanceHealthStatus(request *DescribeSREInstanceHealthStatusRequest) (response *DescribeSREInstanceHealthStatusResponse, err error) {
	if request == nil {
		request = NewDescribeSREInstanceHealthStatusRequest()
	}
	response = NewDescribeSREInstanceHealthStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCloudNativeAPIGatewaySecretRequest() (request *CreateCloudNativeAPIGatewaySecretRequest) {
	request = &CreateCloudNativeAPIGatewaySecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewaySecret")
	return
}

func NewCreateCloudNativeAPIGatewaySecretResponse() (response *CreateCloudNativeAPIGatewaySecretResponse) {
	response = &CreateCloudNativeAPIGatewaySecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建 Nginx Ingress Secret 证书
func (c *Client) CreateCloudNativeAPIGatewaySecret(request *CreateCloudNativeAPIGatewaySecretRequest) (response *CreateCloudNativeAPIGatewaySecretResponse, err error) {
	if request == nil {
		request = NewCreateCloudNativeAPIGatewaySecretRequest()
	}
	response = NewCreateCloudNativeAPIGatewaySecretResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNacosConfigInfoRequest() (request *UpdateNacosConfigInfoRequest) {
	request = &UpdateNacosConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateNacosConfigInfo")
	return
}

func NewUpdateNacosConfigInfoResponse() (response *UpdateNacosConfigInfoResponse) {
	response = &UpdateNacosConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Nacos系统参数
func (c *Client) UpdateNacosConfigInfo(request *UpdateNacosConfigInfoRequest) (response *UpdateNacosConfigInfoResponse, err error) {
	if request == nil {
		request = NewUpdateNacosConfigInfoRequest()
	}
	response = NewUpdateNacosConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewayCanaryRuleRequest() (request *DeleteCloudNativeAPIGatewayCanaryRuleRequest) {
	request = &DeleteCloudNativeAPIGatewayCanaryRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayCanaryRule")
	return
}

func NewDeleteCloudNativeAPIGatewayCanaryRuleResponse() (response *DeleteCloudNativeAPIGatewayCanaryRuleResponse) {
	response = &DeleteCloudNativeAPIGatewayCanaryRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除云原生网关的灰度规则
func (c *Client) DeleteCloudNativeAPIGatewayCanaryRule(request *DeleteCloudNativeAPIGatewayCanaryRuleRequest) (response *DeleteCloudNativeAPIGatewayCanaryRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewayCanaryRuleRequest()
	}
	response = NewDeleteCloudNativeAPIGatewayCanaryRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateEventListRequest() (request *DescribeOperateEventListRequest) {
	request = &DescribeOperateEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateEventList")
	return
}

func NewDescribeOperateEventListResponse() (response *DescribeOperateEventListResponse) {
	response = &DescribeOperateEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台事件列表
func (c *Client) DescribeOperateEventList(request *DescribeOperateEventListRequest) (response *DescribeOperateEventListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateEventListRequest()
	}
	response = NewDescribeOperateEventListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEurekaConfigInfoRequest() (request *UpdateEurekaConfigInfoRequest) {
	request = &UpdateEurekaConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateEurekaConfigInfo")
	return
}

func NewUpdateEurekaConfigInfoResponse() (response *UpdateEurekaConfigInfoResponse) {
	response = &UpdateEurekaConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Eureka参数配置
func (c *Client) UpdateEurekaConfigInfo(request *UpdateEurekaConfigInfoRequest) (response *UpdateEurekaConfigInfoResponse, err error) {
	if request == nil {
		request = NewUpdateEurekaConfigInfoRequest()
	}
	response = NewUpdateEurekaConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewOpenPrometheusRequest() (request *OpenPrometheusRequest) {
	request = &OpenPrometheusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "OpenPrometheus")
	return
}

func NewOpenPrometheusResponse() (response *OpenPrometheusResponse) {
	response = &OpenPrometheusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启Prometheus指标推送
func (c *Client) OpenPrometheus(request *OpenPrometheusRequest) (response *OpenPrometheusResponse, err error) {
	if request == nil {
		request = NewOpenPrometheusRequest()
	}
	response = NewOpenPrometheusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGovernanceInstancesRequest() (request *CreateGovernanceInstancesRequest) {
	request = &CreateGovernanceInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceInstances")
	return
}

func NewCreateGovernanceInstancesResponse() (response *CreateGovernanceInstancesResponse) {
	response = &CreateGovernanceInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建治理中心服务实例
func (c *Client) CreateGovernanceInstances(request *CreateGovernanceInstancesRequest) (response *CreateGovernanceInstancesResponse, err error) {
	if request == nil {
		request = NewCreateGovernanceInstancesRequest()
	}
	response = NewCreateGovernanceInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayNodesRequest() (request *DescribeCloudNativeAPIGatewayNodesRequest) {
	request = &DescribeCloudNativeAPIGatewayNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayNodes")
	return
}

func NewDescribeCloudNativeAPIGatewayNodesResponse() (response *DescribeCloudNativeAPIGatewayNodesResponse) {
	response = &DescribeCloudNativeAPIGatewayNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关节点列表
func (c *Client) DescribeCloudNativeAPIGatewayNodes(request *DescribeCloudNativeAPIGatewayNodesRequest) (response *DescribeCloudNativeAPIGatewayNodesResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayNodesRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSREInstancesRequest() (request *DescribeSREInstancesRequest) {
	request = &DescribeSREInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeSREInstances")
	return
}

func NewDescribeSREInstancesResponse() (response *DescribeSREInstancesResponse) {
	response = &DescribeSREInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询引擎实例列表
func (c *Client) DescribeSREInstances(request *DescribeSREInstancesRequest) (response *DescribeSREInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeSREInstancesRequest()
	}
	response = NewDescribeSREInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSREInstanceFeatureRequest() (request *CheckSREInstanceFeatureRequest) {
	request = &CheckSREInstanceFeatureRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CheckSREInstanceFeature")
	return
}

func NewCheckSREInstanceFeatureResponse() (response *CheckSREInstanceFeatureResponse) {
	response = &CheckSREInstanceFeatureResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查引擎实例的版本特性
func (c *Client) CheckSREInstanceFeature(request *CheckSREInstanceFeatureRequest) (response *CheckSREInstanceFeatureResponse, err error) {
	if request == nil {
		request = NewCheckSREInstanceFeatureRequest()
	}
	response = NewCheckSREInstanceFeatureResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEksMetricsRequest() (request *DescribeEksMetricsRequest) {
	request = &DescribeEksMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEksMetrics")
	return
}

func NewDescribeEksMetricsResponse() (response *DescribeEksMetricsResponse) {
	response = &DescribeEksMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 eks 指标
func (c *Client) DescribeEksMetrics(request *DescribeEksMetricsRequest) (response *DescribeEksMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeEksMetricsRequest()
	}
	response = NewDescribeEksMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNativeGatewayServicesRequest() (request *DescribeNativeGatewayServicesRequest) {
	request = &DescribeNativeGatewayServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeNativeGatewayServices")
	return
}

func NewDescribeNativeGatewayServicesResponse() (response *DescribeNativeGatewayServicesResponse) {
	response = &DescribeNativeGatewayServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关服务列表
func (c *Client) DescribeNativeGatewayServices(request *DescribeNativeGatewayServicesRequest) (response *DescribeNativeGatewayServicesResponse, err error) {
	if request == nil {
		request = NewDescribeNativeGatewayServicesRequest()
	}
	response = NewDescribeNativeGatewayServicesResponse()
	err = c.Send(request, response)
	return
}

func NewLoadZookeeperMigrateDataRequest() (request *LoadZookeeperMigrateDataRequest) {
	request = &LoadZookeeperMigrateDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "LoadZookeeperMigrateData")
	return
}

func NewLoadZookeeperMigrateDataResponse() (response *LoadZookeeperMigrateDataResponse) {
	response = &LoadZookeeperMigrateDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行导入数据
func (c *Client) LoadZookeeperMigrateData(request *LoadZookeeperMigrateDataRequest) (response *LoadZookeeperMigrateDataResponse, err error) {
	if request == nil {
		request = NewLoadZookeeperMigrateDataRequest()
	}
	response = NewLoadZookeeperMigrateDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEurekaServicesRequest() (request *DescribeEurekaServicesRequest) {
	request = &DescribeEurekaServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEurekaServices")
	return
}

func NewDescribeEurekaServicesResponse() (response *DescribeEurekaServicesResponse) {
	response = &DescribeEurekaServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定实例下的eureka注册服务列表
func (c *Client) DescribeEurekaServices(request *DescribeEurekaServicesRequest) (response *DescribeEurekaServicesResponse, err error) {
	if request == nil {
		request = NewDescribeEurekaServicesRequest()
	}
	response = NewDescribeEurekaServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceUserTokenRequest() (request *DescribeGovernanceUserTokenRequest) {
	request = &DescribeGovernanceUserTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceUserToken")
	return
}

func NewDescribeGovernanceUserTokenResponse() (response *DescribeGovernanceUserTokenResponse) {
	response = &DescribeGovernanceUserTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心用户Token
func (c *Client) DescribeGovernanceUserToken(request *DescribeGovernanceUserTokenRequest) (response *DescribeGovernanceUserTokenResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceUserTokenRequest()
	}
	response = NewDescribeGovernanceUserTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateConsoleInstanceDetailRequest() (request *DescribeOperateConsoleInstanceDetailRequest) {
	request = &DescribeOperateConsoleInstanceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateConsoleInstanceDetail")
	return
}

func NewDescribeOperateConsoleInstanceDetailResponse() (response *DescribeOperateConsoleInstanceDetailResponse) {
	response = &DescribeOperateConsoleInstanceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台-平台层实例详情
func (c *Client) DescribeOperateConsoleInstanceDetail(request *DescribeOperateConsoleInstanceDetailRequest) (response *DescribeOperateConsoleInstanceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeOperateConsoleInstanceDetailRequest()
	}
	response = NewDescribeOperateConsoleInstanceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperServerInterfacesRequest() (request *DescribeZookeeperServerInterfacesRequest) {
	request = &DescribeZookeeperServerInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperServerInterfaces")
	return
}

func NewDescribeZookeeperServerInterfacesResponse() (response *DescribeZookeeperServerInterfacesResponse) {
	response = &DescribeZookeeperServerInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询zookeeper服务接口列表
func (c *Client) DescribeZookeeperServerInterfaces(request *DescribeZookeeperServerInterfacesRequest) (response *DescribeZookeeperServerInterfacesResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperServerInterfacesRequest()
	}
	response = NewDescribeZookeeperServerInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateConsoleInstanceLogsRequest() (request *DescribeOperateConsoleInstanceLogsRequest) {
	request = &DescribeOperateConsoleInstanceLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateConsoleInstanceLogs")
	return
}

func NewDescribeOperateConsoleInstanceLogsResponse() (response *DescribeOperateConsoleInstanceLogsResponse) {
	response = &DescribeOperateConsoleInstanceLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台-平台层实例日志
func (c *Client) DescribeOperateConsoleInstanceLogs(request *DescribeOperateConsoleInstanceLogsRequest) (response *DescribeOperateConsoleInstanceLogsResponse, err error) {
	if request == nil {
		request = NewDescribeOperateConsoleInstanceLogsRequest()
	}
	response = NewDescribeOperateConsoleInstanceLogsResponse()
	err = c.Send(request, response)
	return
}

func NewMergeZookeeperMigrateClusterRequest() (request *MergeZookeeperMigrateClusterRequest) {
	request = &MergeZookeeperMigrateClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "MergeZookeeperMigrateCluster")
	return
}

func NewMergeZookeeperMigrateClusterResponse() (response *MergeZookeeperMigrateClusterResponse) {
	response = &MergeZookeeperMigrateClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 合并集群
func (c *Client) MergeZookeeperMigrateCluster(request *MergeZookeeperMigrateClusterRequest) (response *MergeZookeeperMigrateClusterResponse, err error) {
	if request == nil {
		request = NewMergeZookeeperMigrateClusterRequest()
	}
	response = NewMergeZookeeperMigrateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateAgentDetailRequest() (request *DescribeOperateAgentDetailRequest) {
	request = &DescribeOperateAgentDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateAgentDetail")
	return
}

func NewDescribeOperateAgentDetailResponse() (response *DescribeOperateAgentDetailResponse) {
	response = &DescribeOperateAgentDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台Agent详情
func (c *Client) DescribeOperateAgentDetail(request *DescribeOperateAgentDetailRequest) (response *DescribeOperateAgentDetailResponse, err error) {
	if request == nil {
		request = NewDescribeOperateAgentDetailRequest()
	}
	response = NewDescribeOperateAgentDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNativeGatewayServicesRequest() (request *CreateNativeGatewayServicesRequest) {
	request = &CreateNativeGatewayServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateNativeGatewayServices")
	return
}

func NewCreateNativeGatewayServicesResponse() (response *CreateNativeGatewayServicesResponse) {
	response = &CreateNativeGatewayServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网关服务
func (c *Client) CreateNativeGatewayServices(request *CreateNativeGatewayServicesRequest) (response *CreateNativeGatewayServicesResponse, err error) {
	if request == nil {
		request = NewCreateNativeGatewayServicesRequest()
	}
	response = NewCreateNativeGatewayServicesResponse()
	err = c.Send(request, response)
	return
}

func NewStartZookeeperMigratePrepareRequest() (request *StartZookeeperMigratePrepareRequest) {
	request = &StartZookeeperMigratePrepareRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "StartZookeeperMigratePrepare")
	return
}

func NewStartZookeeperMigratePrepareResponse() (response *StartZookeeperMigratePrepareResponse) {
	response = &StartZookeeperMigratePrepareResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开始迁移准备阶段
func (c *Client) StartZookeeperMigratePrepare(request *StartZookeeperMigratePrepareRequest) (response *StartZookeeperMigratePrepareResponse, err error) {
	if request == nil {
		request = NewStartZookeeperMigratePrepareRequest()
	}
	response = NewStartZookeeperMigratePrepareResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePolarisBindK8sClustersRequest() (request *DescribePolarisBindK8sClustersRequest) {
	request = &DescribePolarisBindK8sClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribePolarisBindK8sClusters")
	return
}

func NewDescribePolarisBindK8sClustersResponse() (response *DescribePolarisBindK8sClustersResponse) {
	response = &DescribePolarisBindK8sClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Polaris 查询关联的 kubernetes 集群
func (c *Client) DescribePolarisBindK8sClusters(request *DescribePolarisBindK8sClustersRequest) (response *DescribePolarisBindK8sClustersResponse, err error) {
	if request == nil {
		request = NewDescribePolarisBindK8sClustersRequest()
	}
	response = NewDescribePolarisBindK8sClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSRECodeNamesRequest() (request *DescribeSRECodeNamesRequest) {
	request = &DescribeSRECodeNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeSRECodeNames")
	return
}

func NewDescribeSRECodeNamesResponse() (response *DescribeSRECodeNamesResponse) {
	response = &DescribeSRECodeNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文案名称列表
func (c *Client) DescribeSRECodeNames(request *DescribeSRECodeNamesRequest) (response *DescribeSRECodeNamesResponse, err error) {
	if request == nil {
		request = NewDescribeSRECodeNamesRequest()
	}
	response = NewDescribeSRECodeNamesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceInstancesRequest() (request *ModifyGovernanceInstancesRequest) {
	request = &ModifyGovernanceInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceInstances")
	return
}

func NewModifyGovernanceInstancesResponse() (response *ModifyGovernanceInstancesResponse) {
	response = &ModifyGovernanceInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改治理中心服务实例
func (c *Client) ModifyGovernanceInstances(request *ModifyGovernanceInstancesRequest) (response *ModifyGovernanceInstancesResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceInstancesRequest()
	}
	response = NewModifyGovernanceInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCheckEngineSyncJobRequest() (request *CheckEngineSyncJobRequest) {
	request = &CheckEngineSyncJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CheckEngineSyncJob")
	return
}

func NewCheckEngineSyncJobResponse() (response *CheckEngineSyncJobResponse) {
	response = &CheckEngineSyncJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验引擎数据同步任务
func (c *Client) CheckEngineSyncJob(request *CheckEngineSyncJobRequest) (response *CheckEngineSyncJobResponse, err error) {
	if request == nil {
		request = NewCheckEngineSyncJobRequest()
	}
	response = NewCheckEngineSyncJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayPluginsRequest() (request *DescribeCloudNativeAPIGatewayPluginsRequest) {
	request = &DescribeCloudNativeAPIGatewayPluginsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayPlugins")
	return
}

func NewDescribeCloudNativeAPIGatewayPluginsResponse() (response *DescribeCloudNativeAPIGatewayPluginsResponse) {
	response = &DescribeCloudNativeAPIGatewayPluginsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关插件列表
func (c *Client) DescribeCloudNativeAPIGatewayPlugins(request *DescribeCloudNativeAPIGatewayPluginsRequest) (response *DescribeCloudNativeAPIGatewayPluginsResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayPluginsRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayPluginsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudNativeAPIGatewayClsConfigRequest() (request *ModifyCloudNativeAPIGatewayClsConfigRequest) {
	request = &ModifyCloudNativeAPIGatewayClsConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayClsConfig")
	return
}

func NewModifyCloudNativeAPIGatewayClsConfigResponse() (response *ModifyCloudNativeAPIGatewayClsConfigResponse) {
	response = &ModifyCloudNativeAPIGatewayClsConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生API网关实例CLS配置
func (c *Client) ModifyCloudNativeAPIGatewayClsConfig(request *ModifyCloudNativeAPIGatewayClsConfigRequest) (response *ModifyCloudNativeAPIGatewayClsConfigResponse, err error) {
	if request == nil {
		request = NewModifyCloudNativeAPIGatewayClsConfigRequest()
	}
	response = NewModifyCloudNativeAPIGatewayClsConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperateEksRequest() (request *CreateOperateEksRequest) {
	request = &CreateOperateEksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateOperateEks")
	return
}

func NewCreateOperateEksResponse() (response *CreateOperateEksResponse) {
	response = &CreateOperateEksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运营平台EKS集群
func (c *Client) CreateOperateEks(request *CreateOperateEksRequest) (response *CreateOperateEksResponse, err error) {
	if request == nil {
		request = NewCreateOperateEksRequest()
	}
	response = NewCreateOperateEksResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrModifyCloudNativeAPIGatewayIngressRequest() (request *CreateOrModifyCloudNativeAPIGatewayIngressRequest) {
	request = &CreateOrModifyCloudNativeAPIGatewayIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateOrModifyCloudNativeAPIGatewayIngress")
	return
}

func NewCreateOrModifyCloudNativeAPIGatewayIngressResponse() (response *CreateOrModifyCloudNativeAPIGatewayIngressResponse) {
	response = &CreateOrModifyCloudNativeAPIGatewayIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建或修改Ingress
func (c *Client) CreateOrModifyCloudNativeAPIGatewayIngress(request *CreateOrModifyCloudNativeAPIGatewayIngressRequest) (response *CreateOrModifyCloudNativeAPIGatewayIngressResponse, err error) {
	if request == nil {
		request = NewCreateOrModifyCloudNativeAPIGatewayIngressRequest()
	}
	response = NewCreateOrModifyCloudNativeAPIGatewayIngressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewaySimpleInfosRequest() (request *DescribeCloudNativeAPIGatewaySimpleInfosRequest) {
	request = &DescribeCloudNativeAPIGatewaySimpleInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewaySimpleInfos")
	return
}

func NewDescribeCloudNativeAPIGatewaySimpleInfosResponse() (response *DescribeCloudNativeAPIGatewaySimpleInfosResponse) {
	response = &DescribeCloudNativeAPIGatewaySimpleInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看所有云原生网关简略实例
func (c *Client) DescribeCloudNativeAPIGatewaySimpleInfos(request *DescribeCloudNativeAPIGatewaySimpleInfosRequest) (response *DescribeCloudNativeAPIGatewaySimpleInfosResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewaySimpleInfosRequest()
	}
	response = NewDescribeCloudNativeAPIGatewaySimpleInfosResponse()
	err = c.Send(request, response)
	return
}

func NewGetNginxIngressGlobalConfigRequest() (request *GetNginxIngressGlobalConfigRequest) {
	request = &GetNginxIngressGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "GetNginxIngressGlobalConfig")
	return
}

func NewGetNginxIngressGlobalConfigResponse() (response *GetNginxIngressGlobalConfigResponse) {
	response = &GetNginxIngressGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取 Nginx Ingress 全部配置参数
func (c *Client) GetNginxIngressGlobalConfig(request *GetNginxIngressGlobalConfigRequest) (response *GetNginxIngressGlobalConfigResponse, err error) {
	if request == nil {
		request = NewGetNginxIngressGlobalConfigRequest()
	}
	response = NewGetNginxIngressGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperateFeatureVersionRequest() (request *CreateOperateFeatureVersionRequest) {
	request = &CreateOperateFeatureVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateOperateFeatureVersion")
	return
}

func NewCreateOperateFeatureVersionResponse() (response *CreateOperateFeatureVersionResponse) {
	response = &CreateOperateFeatureVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建功能版本
func (c *Client) CreateOperateFeatureVersion(request *CreateOperateFeatureVersionRequest) (response *CreateOperateFeatureVersionResponse, err error) {
	if request == nil {
		request = NewCreateOperateFeatureVersionRequest()
	}
	response = NewCreateOperateFeatureVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKongIngressRequest() (request *DescribeKongIngressRequest) {
	request = &DescribeKongIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeKongIngress")
	return
}

func NewDescribeKongIngressResponse() (response *DescribeKongIngressResponse) {
	response = &DescribeKongIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询KongIngress状态
func (c *Client) DescribeKongIngress(request *DescribeKongIngressRequest) (response *DescribeKongIngressResponse, err error) {
	if request == nil {
		request = NewDescribeKongIngressRequest()
	}
	response = NewDescribeKongIngressResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateZookeeperMigrateUploadMetaInfoRequest() (request *UpdateZookeeperMigrateUploadMetaInfoRequest) {
	request = &UpdateZookeeperMigrateUploadMetaInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateZookeeperMigrateUploadMetaInfo")
	return
}

func NewUpdateZookeeperMigrateUploadMetaInfoResponse() (response *UpdateZookeeperMigrateUploadMetaInfoResponse) {
	response = &UpdateZookeeperMigrateUploadMetaInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新md5值
func (c *Client) UpdateZookeeperMigrateUploadMetaInfo(request *UpdateZookeeperMigrateUploadMetaInfoRequest) (response *UpdateZookeeperMigrateUploadMetaInfoResponse, err error) {
	if request == nil {
		request = NewUpdateZookeeperMigrateUploadMetaInfoRequest()
	}
	response = NewUpdateZookeeperMigrateUploadMetaInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCredentialRequest() (request *DescribeCredentialRequest) {
	request = &DescribeCredentialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCredential")
	return
}

func NewDescribeCredentialResponse() (response *DescribeCredentialResponse) {
	response = &DescribeCredentialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看临时凭证
func (c *Client) DescribeCredential(request *DescribeCredentialRequest) (response *DescribeCredentialResponse, err error) {
	if request == nil {
		request = NewDescribeCredentialRequest()
	}
	response = NewDescribeCredentialResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperServiceInstancesRequest() (request *DescribeZookeeperServiceInstancesRequest) {
	request = &DescribeZookeeperServiceInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperServiceInstances")
	return
}

func NewDescribeZookeeperServiceInstancesResponse() (response *DescribeZookeeperServiceInstancesResponse) {
	response = &DescribeZookeeperServiceInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定实例下zookeeper注册服务的实例列表
func (c *Client) DescribeZookeeperServiceInstances(request *DescribeZookeeperServiceInstancesRequest) (response *DescribeZookeeperServiceInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperServiceInstancesRequest()
	}
	response = NewDescribeZookeeperServiceInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifySREInstanceSpecRequest() (request *ModifySREInstanceSpecRequest) {
	request = &ModifySREInstanceSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifySREInstanceSpec")
	return
}

func NewModifySREInstanceSpecResponse() (response *ModifySREInstanceSpecResponse) {
	response = &ModifySREInstanceSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改微服务引擎实例的规格
func (c *Client) ModifySREInstanceSpec(request *ModifySREInstanceSpecRequest) (response *ModifySREInstanceSpecResponse, err error) {
	if request == nil {
		request = NewModifySREInstanceSpecRequest()
	}
	response = NewModifySREInstanceSpecResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGovernanceStrategiesRequest() (request *DeleteGovernanceStrategiesRequest) {
	request = &DeleteGovernanceStrategiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceStrategies")
	return
}

func NewDeleteGovernanceStrategiesResponse() (response *DeleteGovernanceStrategiesResponse) {
	response = &DeleteGovernanceStrategiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除治理中心鉴权策略
func (c *Client) DeleteGovernanceStrategies(request *DeleteGovernanceStrategiesRequest) (response *DeleteGovernanceStrategiesResponse, err error) {
	if request == nil {
		request = NewDeleteGovernanceStrategiesRequest()
	}
	response = NewDeleteGovernanceStrategiesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateWhiteListRequest() (request *ModifyOperateWhiteListRequest) {
	request = &ModifyOperateWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyOperateWhiteList")
	return
}

func NewModifyOperateWhiteListResponse() (response *ModifyOperateWhiteListResponse) {
	response = &ModifyOperateWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新运营平台白名单
func (c *Client) ModifyOperateWhiteList(request *ModifyOperateWhiteListRequest) (response *ModifyOperateWhiteListResponse, err error) {
	if request == nil {
		request = NewModifyOperateWhiteListRequest()
	}
	response = NewModifyOperateWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperateRegionRequest() (request *CreateOperateRegionRequest) {
	request = &CreateOperateRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateOperateRegion")
	return
}

func NewCreateOperateRegionResponse() (response *CreateOperateRegionResponse) {
	response = &CreateOperateRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运营平台地域
func (c *Client) CreateOperateRegion(request *CreateOperateRegionRequest) (response *CreateOperateRegionResponse, err error) {
	if request == nil {
		request = NewCreateOperateRegionRequest()
	}
	response = NewCreateOperateRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewayRouteRequest() (request *DeleteCloudNativeAPIGatewayRouteRequest) {
	request = &DeleteCloudNativeAPIGatewayRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayRoute")
	return
}

func NewDeleteCloudNativeAPIGatewayRouteResponse() (response *DeleteCloudNativeAPIGatewayRouteResponse) {
	response = &DeleteCloudNativeAPIGatewayRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除云原生网关路由
func (c *Client) DeleteCloudNativeAPIGatewayRoute(request *DeleteCloudNativeAPIGatewayRouteRequest) (response *DeleteCloudNativeAPIGatewayRouteResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewayRouteRequest()
	}
	response = NewDeleteCloudNativeAPIGatewayRouteResponse()
	err = c.Send(request, response)
	return
}

func NewOpenGovernanceRequest() (request *OpenGovernanceRequest) {
	request = &OpenGovernanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "OpenGovernance")
	return
}

func NewOpenGovernanceResponse() (response *OpenGovernanceResponse) {
	response = &OpenGovernanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开通服务治理中心
func (c *Client) OpenGovernance(request *OpenGovernanceRequest) (response *OpenGovernanceResponse, err error) {
	if request == nil {
		request = NewOpenGovernanceRequest()
	}
	response = NewOpenGovernanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudAlarmPolarisPodsRequest() (request *DescribeCloudAlarmPolarisPodsRequest) {
	request = &DescribeCloudAlarmPolarisPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudAlarmPolarisPods")
	return
}

func NewDescribeCloudAlarmPolarisPodsResponse() (response *DescribeCloudAlarmPolarisPodsResponse) {
	response = &DescribeCloudAlarmPolarisPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询北极星的实例列表
func (c *Client) DescribeCloudAlarmPolarisPods(request *DescribeCloudAlarmPolarisPodsRequest) (response *DescribeCloudAlarmPolarisPodsResponse, err error) {
	if request == nil {
		request = NewDescribeCloudAlarmPolarisPodsRequest()
	}
	response = NewDescribeCloudAlarmPolarisPodsResponse()
	err = c.Send(request, response)
	return
}

func NewResetZookeeperMigrateRequest() (request *ResetZookeeperMigrateRequest) {
	request = &ResetZookeeperMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ResetZookeeperMigrate")
	return
}

func NewResetZookeeperMigrateResponse() (response *ResetZookeeperMigrateResponse) {
	response = &ResetZookeeperMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置迁移任务
func (c *Client) ResetZookeeperMigrate(request *ResetZookeeperMigrateRequest) (response *ResetZookeeperMigrateResponse, err error) {
	if request == nil {
		request = NewResetZookeeperMigrateRequest()
	}
	response = NewResetZookeeperMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewBindAutoScalerResourceStrategyToGroupsRequest() (request *BindAutoScalerResourceStrategyToGroupsRequest) {
	request = &BindAutoScalerResourceStrategyToGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "BindAutoScalerResourceStrategyToGroups")
	return
}

func NewBindAutoScalerResourceStrategyToGroupsResponse() (response *BindAutoScalerResourceStrategyToGroupsResponse) {
	response = &BindAutoScalerResourceStrategyToGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 弹性伸缩策略批量绑定网关分组
func (c *Client) BindAutoScalerResourceStrategyToGroups(request *BindAutoScalerResourceStrategyToGroupsRequest) (response *BindAutoScalerResourceStrategyToGroupsResponse, err error) {
	if request == nil {
		request = NewBindAutoScalerResourceStrategyToGroupsRequest()
	}
	response = NewBindAutoScalerResourceStrategyToGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceUserPasswordRequest() (request *ModifyGovernanceUserPasswordRequest) {
	request = &ModifyGovernanceUserPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceUserPassword")
	return
}

func NewModifyGovernanceUserPasswordResponse() (response *ModifyGovernanceUserPasswordResponse) {
	response = &ModifyGovernanceUserPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户密码
func (c *Client) ModifyGovernanceUserPassword(request *ModifyGovernanceUserPasswordRequest) (response *ModifyGovernanceUserPasswordResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceUserPasswordRequest()
	}
	response = NewModifyGovernanceUserPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewaySystemParametersRequest() (request *DescribeCloudNativeAPIGatewaySystemParametersRequest) {
	request = &DescribeCloudNativeAPIGatewaySystemParametersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewaySystemParameters")
	return
}

func NewDescribeCloudNativeAPIGatewaySystemParametersResponse() (response *DescribeCloudNativeAPIGatewaySystemParametersResponse) {
	response = &DescribeCloudNativeAPIGatewaySystemParametersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关系统参数列表
func (c *Client) DescribeCloudNativeAPIGatewaySystemParameters(request *DescribeCloudNativeAPIGatewaySystemParametersRequest) (response *DescribeCloudNativeAPIGatewaySystemParametersResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewaySystemParametersRequest()
	}
	response = NewDescribeCloudNativeAPIGatewaySystemParametersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateRegionsRequest() (request *DescribeOperateRegionsRequest) {
	request = &DescribeOperateRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateRegions")
	return
}

func NewDescribeOperateRegionsResponse() (response *DescribeOperateRegionsResponse) {
	response = &DescribeOperateRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台地域列表
func (c *Client) DescribeOperateRegions(request *DescribeOperateRegionsRequest) (response *DescribeOperateRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeOperateRegionsRequest()
	}
	response = NewDescribeOperateRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateInstanceUpgradeRequest() (request *ModifyOperateInstanceUpgradeRequest) {
	request = &ModifyOperateInstanceUpgradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyOperateInstanceUpgrade")
	return
}

func NewModifyOperateInstanceUpgradeResponse() (response *ModifyOperateInstanceUpgradeResponse) {
	response = &ModifyOperateInstanceUpgradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// tse运营端引擎升级
func (c *Client) ModifyOperateInstanceUpgrade(request *ModifyOperateInstanceUpgradeRequest) (response *ModifyOperateInstanceUpgradeResponse, err error) {
	if request == nil {
		request = NewModifyOperateInstanceUpgradeRequest()
	}
	response = NewModifyOperateInstanceUpgradeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateSecurityGroupRequest() (request *ModifyOperateSecurityGroupRequest) {
	request = &ModifyOperateSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyOperateSecurityGroup")
	return
}

func NewModifyOperateSecurityGroupResponse() (response *ModifyOperateSecurityGroupResponse) {
	response = &ModifyOperateSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新运营平台安全组
func (c *Client) ModifyOperateSecurityGroup(request *ModifyOperateSecurityGroupRequest) (response *ModifyOperateSecurityGroupResponse, err error) {
	if request == nil {
		request = NewModifyOperateSecurityGroupRequest()
	}
	response = NewModifyOperateSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApolloEnvReplicasRequest() (request *DescribeApolloEnvReplicasRequest) {
	request = &DescribeApolloEnvReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeApolloEnvReplicas")
	return
}

func NewDescribeApolloEnvReplicasResponse() (response *DescribeApolloEnvReplicasResponse) {
	response = &DescribeApolloEnvReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Apollo具体某个环境的所有节点信息
func (c *Client) DescribeApolloEnvReplicas(request *DescribeApolloEnvReplicasRequest) (response *DescribeApolloEnvReplicasResponse, err error) {
	if request == nil {
		request = NewDescribeApolloEnvReplicasRequest()
	}
	response = NewDescribeApolloEnvReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPolarisAuthStatusRequest() (request *ModifyPolarisAuthStatusRequest) {
	request = &ModifyPolarisAuthStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyPolarisAuthStatus")
	return
}

func NewModifyPolarisAuthStatusResponse() (response *ModifyPolarisAuthStatusResponse) {
	response = &ModifyPolarisAuthStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更治理中心鉴权开关状态
func (c *Client) ModifyPolarisAuthStatus(request *ModifyPolarisAuthStatusRequest) (response *ModifyPolarisAuthStatusResponse, err error) {
	if request == nil {
		request = NewModifyPolarisAuthStatusRequest()
	}
	response = NewModifyPolarisAuthStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceAuthStatusRequest() (request *DescribeGovernanceAuthStatusRequest) {
	request = &DescribeGovernanceAuthStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceAuthStatus")
	return
}

func NewDescribeGovernanceAuthStatusResponse() (response *DescribeGovernanceAuthStatusResponse) {
	response = &DescribeGovernanceAuthStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心鉴权功能是否开启
func (c *Client) DescribeGovernanceAuthStatus(request *DescribeGovernanceAuthStatusRequest) (response *DescribeGovernanceAuthStatusResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceAuthStatusRequest()
	}
	response = NewDescribeGovernanceAuthStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNacosConfigInfoRequest() (request *DescribeNacosConfigInfoRequest) {
	request = &DescribeNacosConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeNacosConfigInfo")
	return
}

func NewDescribeNacosConfigInfoResponse() (response *DescribeNacosConfigInfoResponse) {
	response = &DescribeNacosConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看Nacos系统参数
func (c *Client) DescribeNacosConfigInfo(request *DescribeNacosConfigInfoRequest) (response *DescribeNacosConfigInfoResponse, err error) {
	if request == nil {
		request = NewDescribeNacosConfigInfoRequest()
	}
	response = NewDescribeNacosConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayUserAuthRequest() (request *DescribeCloudNativeAPIGatewayUserAuthRequest) {
	request = &DescribeCloudNativeAPIGatewayUserAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayUserAuth")
	return
}

func NewDescribeCloudNativeAPIGatewayUserAuthResponse() (response *DescribeCloudNativeAPIGatewayUserAuthResponse) {
	response = &DescribeCloudNativeAPIGatewayUserAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云原生网关用户的权限
func (c *Client) DescribeCloudNativeAPIGatewayUserAuth(request *DescribeCloudNativeAPIGatewayUserAuthRequest) (response *DescribeCloudNativeAPIGatewayUserAuthResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayUserAuthRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayUserAuthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsulReplicasRequest() (request *DescribeConsulReplicasRequest) {
	request = &DescribeConsulReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConsulReplicas")
	return
}

func NewDescribeConsulReplicasResponse() (response *DescribeConsulReplicasResponse) {
	response = &DescribeConsulReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Consul类型注册引擎实例副本信息
func (c *Client) DescribeConsulReplicas(request *DescribeConsulReplicasRequest) (response *DescribeConsulReplicasResponse, err error) {
	if request == nil {
		request = NewDescribeConsulReplicasRequest()
	}
	response = NewDescribeConsulReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEnginePodAccessAddressRequest() (request *DescribeEnginePodAccessAddressRequest) {
	request = &DescribeEnginePodAccessAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEnginePodAccessAddress")
	return
}

func NewDescribeEnginePodAccessAddressResponse() (response *DescribeEnginePodAccessAddressResponse) {
	response = &DescribeEnginePodAccessAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回engine每个节点的地址
func (c *Client) DescribeEnginePodAccessAddress(request *DescribeEnginePodAccessAddressRequest) (response *DescribeEnginePodAccessAddressResponse, err error) {
	if request == nil {
		request = NewDescribeEnginePodAccessAddressRequest()
	}
	response = NewDescribeEnginePodAccessAddressResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSREInstanceRequest() (request *CreateSREInstanceRequest) {
	request = &CreateSREInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateSREInstance")
	return
}

func NewCreateSREInstanceResponse() (response *CreateSREInstanceResponse) {
	response = &CreateSREInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建微服务引擎实例
func (c *Client) CreateSREInstance(request *CreateSREInstanceRequest) (response *CreateSREInstanceResponse, err error) {
	if request == nil {
		request = NewCreateSREInstanceRequest()
	}
	response = NewCreateSREInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewStartZookeeperMigrateMergeClusterRequest() (request *StartZookeeperMigrateMergeClusterRequest) {
	request = &StartZookeeperMigrateMergeClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "StartZookeeperMigrateMergeCluster")
	return
}

func NewStartZookeeperMigrateMergeClusterResponse() (response *StartZookeeperMigrateMergeClusterResponse) {
	response = &StartZookeeperMigrateMergeClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开始迁移合并阶段
func (c *Client) StartZookeeperMigrateMergeCluster(request *StartZookeeperMigrateMergeClusterRequest) (response *StartZookeeperMigrateMergeClusterResponse, err error) {
	if request == nil {
		request = NewStartZookeeperMigrateMergeClusterRequest()
	}
	response = NewStartZookeeperMigrateMergeClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGovernanceNamespacesRequest() (request *DeleteGovernanceNamespacesRequest) {
	request = &DeleteGovernanceNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceNamespaces")
	return
}

func NewDeleteGovernanceNamespacesResponse() (response *DeleteGovernanceNamespacesResponse) {
	response = &DeleteGovernanceNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除治理中心命名空间
func (c *Client) DeleteGovernanceNamespaces(request *DeleteGovernanceNamespacesRequest) (response *DeleteGovernanceNamespacesResponse, err error) {
	if request == nil {
		request = NewDeleteGovernanceNamespacesRequest()
	}
	response = NewDeleteGovernanceNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewayCanaryRulesRequest() (request *ListCloudNativeAPIGatewayCanaryRulesRequest) {
	request = &ListCloudNativeAPIGatewayCanaryRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGatewayCanaryRules")
	return
}

func NewListCloudNativeAPIGatewayCanaryRulesResponse() (response *ListCloudNativeAPIGatewayCanaryRulesResponse) {
	response = &ListCloudNativeAPIGatewayCanaryRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关灰度规则列表
func (c *Client) ListCloudNativeAPIGatewayCanaryRules(request *ListCloudNativeAPIGatewayCanaryRulesRequest) (response *ListCloudNativeAPIGatewayCanaryRulesResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewayCanaryRulesRequest()
	}
	response = NewListCloudNativeAPIGatewayCanaryRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadLicenseRequest() (request *DownloadLicenseRequest) {
	request = &DownloadLicenseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DownloadLicense")
	return
}

func NewDownloadLicenseResponse() (response *DownloadLicenseResponse) {
	response = &DownloadLicenseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载License文件
func (c *Client) DownloadLicense(request *DownloadLicenseRequest) (response *DownloadLicenseResponse, err error) {
	if request == nil {
		request = NewDownloadLicenseRequest()
	}
	response = NewDownloadLicenseResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateConsulConfigInfoRequest() (request *UpdateConsulConfigInfoRequest) {
	request = &UpdateConsulConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateConsulConfigInfo")
	return
}

func NewUpdateConsulConfigInfoResponse() (response *UpdateConsulConfigInfoResponse) {
	response = &UpdateConsulConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Consul参数配置
func (c *Client) UpdateConsulConfigInfo(request *UpdateConsulConfigInfoRequest) (response *UpdateConsulConfigInfoResponse, err error) {
	if request == nil {
		request = NewUpdateConsulConfigInfoRequest()
	}
	response = NewUpdateConsulConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateEngineVersionListRequest() (request *DescribeOperateEngineVersionListRequest) {
	request = &DescribeOperateEngineVersionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateEngineVersionList")
	return
}

func NewDescribeOperateEngineVersionListResponse() (response *DescribeOperateEngineVersionListResponse) {
	response = &DescribeOperateEngineVersionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端引擎版本
func (c *Client) DescribeOperateEngineVersionList(request *DescribeOperateEngineVersionListRequest) (response *DescribeOperateEngineVersionListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateEngineVersionListRequest()
	}
	response = NewDescribeOperateEngineVersionListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApolloAdminInfoRequest() (request *ModifyApolloAdminInfoRequest) {
	request = &ModifyApolloAdminInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyApolloAdminInfo")
	return
}

func NewModifyApolloAdminInfoResponse() (response *ModifyApolloAdminInfoResponse) {
	response = &ModifyApolloAdminInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改apollo管理员信息
func (c *Client) ModifyApolloAdminInfo(request *ModifyApolloAdminInfoRequest) (response *ModifyApolloAdminInfoResponse, err error) {
	if request == nil {
		request = NewModifyApolloAdminInfoRequest()
	}
	response = NewModifyApolloAdminInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApolloReplicasRequest() (request *DescribeApolloReplicasRequest) {
	request = &DescribeApolloReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeApolloReplicas")
	return
}

func NewDescribeApolloReplicasResponse() (response *DescribeApolloReplicasResponse) {
	response = &DescribeApolloReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Apollo 集群实例信息
func (c *Client) DescribeApolloReplicas(request *DescribeApolloReplicasRequest) (response *DescribeApolloReplicasResponse, err error) {
	if request == nil {
		request = NewDescribeApolloReplicasRequest()
	}
	response = NewDescribeApolloReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewOpenKongIngressRequest() (request *OpenKongIngressRequest) {
	request = &OpenKongIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "OpenKongIngress")
	return
}

func NewOpenKongIngressResponse() (response *OpenKongIngressResponse) {
	response = &OpenKongIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启KongIngress
func (c *Client) OpenKongIngress(request *OpenKongIngressRequest) (response *OpenKongIngressResponse, err error) {
	if request == nil {
		request = NewOpenKongIngressRequest()
	}
	response = NewOpenKongIngressResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGovernanceAliasesRequest() (request *DeleteGovernanceAliasesRequest) {
	request = &DeleteGovernanceAliasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceAliases")
	return
}

func NewDeleteGovernanceAliasesResponse() (response *DeleteGovernanceAliasesResponse) {
	response = &DeleteGovernanceAliasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除治理中心服务别名
func (c *Client) DeleteGovernanceAliases(request *DeleteGovernanceAliasesRequest) (response *DeleteGovernanceAliasesResponse, err error) {
	if request == nil {
		request = NewDeleteGovernanceAliasesRequest()
	}
	response = NewDeleteGovernanceAliasesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCLBSecurityGroupRequest() (request *DescribeCLBSecurityGroupRequest) {
	request = &DescribeCLBSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCLBSecurityGroup")
	return
}

func NewDescribeCLBSecurityGroupResponse() (response *DescribeCLBSecurityGroupResponse) {
	response = &DescribeCLBSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 clb 安全组信息
func (c *Client) DescribeCLBSecurityGroup(request *DescribeCLBSecurityGroupRequest) (response *DescribeCLBSecurityGroupResponse, err error) {
	if request == nil {
		request = NewDescribeCLBSecurityGroupRequest()
	}
	response = NewDescribeCLBSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayVersionRequest() (request *DescribeCloudNativeAPIGatewayVersionRequest) {
	request = &DescribeCloudNativeAPIGatewayVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayVersion")
	return
}

func NewDescribeCloudNativeAPIGatewayVersionResponse() (response *DescribeCloudNativeAPIGatewayVersionResponse) {
	response = &DescribeCloudNativeAPIGatewayVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关实例版本列表
func (c *Client) DescribeCloudNativeAPIGatewayVersion(request *DescribeCloudNativeAPIGatewayVersionRequest) (response *DescribeCloudNativeAPIGatewayVersionResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayVersionRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateInstanceLogsRequest() (request *DescribeOperateInstanceLogsRequest) {
	request = &DescribeOperateInstanceLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateInstanceLogs")
	return
}

func NewDescribeOperateInstanceLogsResponse() (response *DescribeOperateInstanceLogsResponse) {
	response = &DescribeOperateInstanceLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台引擎实例日志
func (c *Client) DescribeOperateInstanceLogs(request *DescribeOperateInstanceLogsRequest) (response *DescribeOperateInstanceLogsResponse, err error) {
	if request == nil {
		request = NewDescribeOperateInstanceLogsRequest()
	}
	response = NewDescribeOperateInstanceLogsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCloudNativeAPIGatewayRequest() (request *CreateCloudNativeAPIGatewayRequest) {
	request = &CreateCloudNativeAPIGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGateway")
	return
}

func NewCreateCloudNativeAPIGatewayResponse() (response *CreateCloudNativeAPIGatewayResponse) {
	response = &CreateCloudNativeAPIGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建云原生API网关实例
func (c *Client) CreateCloudNativeAPIGateway(request *CreateCloudNativeAPIGatewayRequest) (response *CreateCloudNativeAPIGatewayResponse, err error) {
	if request == nil {
		request = NewCreateCloudNativeAPIGatewayRequest()
	}
	response = NewCreateCloudNativeAPIGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGovernanceGroupRequest() (request *CreateGovernanceGroupRequest) {
	request = &CreateGovernanceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceGroup")
	return
}

func NewCreateGovernanceGroupResponse() (response *CreateGovernanceGroupResponse) {
	response = &CreateGovernanceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建治理中心用户组
func (c *Client) CreateGovernanceGroup(request *CreateGovernanceGroupRequest) (response *CreateGovernanceGroupResponse, err error) {
	if request == nil {
		request = NewCreateGovernanceGroupRequest()
	}
	response = NewCreateGovernanceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNativeGatewayServerGroupRequest() (request *CreateNativeGatewayServerGroupRequest) {
	request = &CreateNativeGatewayServerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateNativeGatewayServerGroup")
	return
}

func NewCreateNativeGatewayServerGroupResponse() (response *CreateNativeGatewayServerGroupResponse) {
	response = &CreateNativeGatewayServerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建云原生网关引擎分组
func (c *Client) CreateNativeGatewayServerGroup(request *CreateNativeGatewayServerGroupRequest) (response *CreateNativeGatewayServerGroupResponse, err error) {
	if request == nil {
		request = NewCreateNativeGatewayServerGroupRequest()
	}
	response = NewCreateNativeGatewayServerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNativeGatewayServiceGroupRequest() (request *CreateNativeGatewayServiceGroupRequest) {
	request = &CreateNativeGatewayServiceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateNativeGatewayServiceGroup")
	return
}

func NewCreateNativeGatewayServiceGroupResponse() (response *CreateNativeGatewayServiceGroupResponse) {
	response = &CreateNativeGatewayServiceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网关服务实例分组
func (c *Client) CreateNativeGatewayServiceGroup(request *CreateNativeGatewayServiceGroupRequest) (response *CreateNativeGatewayServiceGroupResponse, err error) {
	if request == nil {
		request = NewCreateNativeGatewayServiceGroupRequest()
	}
	response = NewCreateNativeGatewayServiceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateConsoleInstancesRequest() (request *DescribeOperateConsoleInstancesRequest) {
	request = &DescribeOperateConsoleInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateConsoleInstances")
	return
}

func NewDescribeOperateConsoleInstancesResponse() (response *DescribeOperateConsoleInstancesResponse) {
	response = &DescribeOperateConsoleInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台-平台层实例列表
func (c *Client) DescribeOperateConsoleInstances(request *DescribeOperateConsoleInstancesRequest) (response *DescribeOperateConsoleInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeOperateConsoleInstancesRequest()
	}
	response = NewDescribeOperateConsoleInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewRebuildPolarisCLSTopicRequest() (request *RebuildPolarisCLSTopicRequest) {
	request = &RebuildPolarisCLSTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "RebuildPolarisCLSTopic")
	return
}

func NewRebuildPolarisCLSTopicResponse() (response *RebuildPolarisCLSTopicResponse) {
	response = &RebuildPolarisCLSTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修复Polaris的CLS日志topic列表
func (c *Client) RebuildPolarisCLSTopic(request *RebuildPolarisCLSTopicRequest) (response *RebuildPolarisCLSTopicResponse, err error) {
	if request == nil {
		request = NewRebuildPolarisCLSTopicRequest()
	}
	response = NewRebuildPolarisCLSTopicResponse()
	err = c.Send(request, response)
	return
}

func NewFinishMigrateRequest() (request *FinishMigrateRequest) {
	request = &FinishMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "FinishMigrate")
	return
}

func NewFinishMigrateResponse() (response *FinishMigrateResponse) {
	response = &FinishMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 结束迁移任务
func (c *Client) FinishMigrate(request *FinishMigrateRequest) (response *FinishMigrateResponse, err error) {
	if request == nil {
		request = NewFinishMigrateRequest()
	}
	response = NewFinishMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewEnableIntranetRequest() (request *EnableIntranetRequest) {
	request = &EnableIntranetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "EnableIntranet")
	return
}

func NewEnableIntranetResponse() (response *EnableIntranetResponse) {
	response = &EnableIntranetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群客户端内网开关
func (c *Client) EnableIntranet(request *EnableIntranetRequest) (response *EnableIntranetResponse, err error) {
	if request == nil {
		request = NewEnableIntranetRequest()
	}
	response = NewEnableIntranetResponse()
	err = c.Send(request, response)
	return
}

func NewStartZookeeperMigrateImportDataRequest() (request *StartZookeeperMigrateImportDataRequest) {
	request = &StartZookeeperMigrateImportDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "StartZookeeperMigrateImportData")
	return
}

func NewStartZookeeperMigrateImportDataResponse() (response *StartZookeeperMigrateImportDataResponse) {
	response = &StartZookeeperMigrateImportDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开始迁移导入数据阶段
func (c *Client) StartZookeeperMigrateImportData(request *StartZookeeperMigrateImportDataRequest) (response *StartZookeeperMigrateImportDataResponse, err error) {
	if request == nil {
		request = NewStartZookeeperMigrateImportDataRequest()
	}
	response = NewStartZookeeperMigrateImportDataResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSREInstanceBasicInfoRequest() (request *UpdateSREInstanceBasicInfoRequest) {
	request = &UpdateSREInstanceBasicInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateSREInstanceBasicInfo")
	return
}

func NewUpdateSREInstanceBasicInfoResponse() (response *UpdateSREInstanceBasicInfoResponse) {
	response = &UpdateSREInstanceBasicInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新实例基本信息
func (c *Client) UpdateSREInstanceBasicInfo(request *UpdateSREInstanceBasicInfoRequest) (response *UpdateSREInstanceBasicInfoResponse, err error) {
	if request == nil {
		request = NewUpdateSREInstanceBasicInfoRequest()
	}
	response = NewUpdateSREInstanceBasicInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsulServicesRequest() (request *DescribeConsulServicesRequest) {
	request = &DescribeConsulServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConsulServices")
	return
}

func NewDescribeConsulServicesResponse() (response *DescribeConsulServicesResponse) {
	response = &DescribeConsulServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定实例下的consul注册服务列表
func (c *Client) DescribeConsulServices(request *DescribeConsulServicesRequest) (response *DescribeConsulServicesResponse, err error) {
	if request == nil {
		request = NewDescribeConsulServicesRequest()
	}
	response = NewDescribeConsulServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMainControlRegionInfosRequest() (request *DescribeMainControlRegionInfosRequest) {
	request = &DescribeMainControlRegionInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeMainControlRegionInfos")
	return
}

func NewDescribeMainControlRegionInfosResponse() (response *DescribeMainControlRegionInfosResponse) {
	response = &DescribeMainControlRegionInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主控区域相关信息
func (c *Client) DescribeMainControlRegionInfos(request *DescribeMainControlRegionInfosRequest) (response *DescribeMainControlRegionInfosResponse, err error) {
	if request == nil {
		request = NewDescribeMainControlRegionInfosRequest()
	}
	response = NewDescribeMainControlRegionInfosResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceServicesRequest() (request *ModifyGovernanceServicesRequest) {
	request = &ModifyGovernanceServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceServices")
	return
}

func NewModifyGovernanceServicesResponse() (response *ModifyGovernanceServicesResponse) {
	response = &ModifyGovernanceServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改治理中心服务
func (c *Client) ModifyGovernanceServices(request *ModifyGovernanceServicesRequest) (response *ModifyGovernanceServicesResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceServicesRequest()
	}
	response = NewModifyGovernanceServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNativeGatewayServiceSourceRequest() (request *DeleteNativeGatewayServiceSourceRequest) {
	request = &DeleteNativeGatewayServiceSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteNativeGatewayServiceSource")
	return
}

func NewDeleteNativeGatewayServiceSourceResponse() (response *DeleteNativeGatewayServiceSourceResponse) {
	response = &DeleteNativeGatewayServiceSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网关服务来源实例
func (c *Client) DeleteNativeGatewayServiceSource(request *DeleteNativeGatewayServiceSourceRequest) (response *DeleteNativeGatewayServiceSourceResponse, err error) {
	if request == nil {
		request = NewDeleteNativeGatewayServiceSourceRequest()
	}
	response = NewDeleteNativeGatewayServiceSourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSREInstanceAccessAddressRequest() (request *DescribeSREInstanceAccessAddressRequest) {
	request = &DescribeSREInstanceAccessAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeSREInstanceAccessAddress")
	return
}

func NewDescribeSREInstanceAccessAddressResponse() (response *DescribeSREInstanceAccessAddressResponse) {
	response = &DescribeSREInstanceAccessAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询引擎实例访问地址
func (c *Client) DescribeSREInstanceAccessAddress(request *DescribeSREInstanceAccessAddressRequest) (response *DescribeSREInstanceAccessAddressResponse, err error) {
	if request == nil {
		request = NewDescribeSREInstanceAccessAddressRequest()
	}
	response = NewDescribeSREInstanceAccessAddressResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewayPluginsRequest() (request *ListCloudNativeAPIGatewayPluginsRequest) {
	request = &ListCloudNativeAPIGatewayPluginsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGatewayPlugins")
	return
}

func NewListCloudNativeAPIGatewayPluginsResponse() (response *ListCloudNativeAPIGatewayPluginsResponse) {
	response = &ListCloudNativeAPIGatewayPluginsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关插件列表
func (c *Client) ListCloudNativeAPIGatewayPlugins(request *ListCloudNativeAPIGatewayPluginsRequest) (response *ListCloudNativeAPIGatewayPluginsResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewayPluginsRequest()
	}
	response = NewListCloudNativeAPIGatewayPluginsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNginxIngressConfigMapRequest() (request *ModifyNginxIngressConfigMapRequest) {
	request = &ModifyNginxIngressConfigMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyNginxIngressConfigMap")
	return
}

func NewModifyNginxIngressConfigMapResponse() (response *ModifyNginxIngressConfigMapResponse) {
	response = &ModifyNginxIngressConfigMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑 Nginx Ingress ConfigMap
func (c *Client) ModifyNginxIngressConfigMap(request *ModifyNginxIngressConfigMapRequest) (response *ModifyNginxIngressConfigMapResponse, err error) {
	if request == nil {
		request = NewModifyNginxIngressConfigMapRequest()
	}
	response = NewModifyNginxIngressConfigMapResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayLatestTaskPhasesRequest() (request *DescribeCloudNativeAPIGatewayLatestTaskPhasesRequest) {
	request = &DescribeCloudNativeAPIGatewayLatestTaskPhasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayLatestTaskPhases")
	return
}

func NewDescribeCloudNativeAPIGatewayLatestTaskPhasesResponse() (response *DescribeCloudNativeAPIGatewayLatestTaskPhasesResponse) {
	response = &DescribeCloudNativeAPIGatewayLatestTaskPhasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当执行部分流程较长的任务时，任务会被拆分为多个阶段进行异步执行。此接口用于查询最新提交的任务的执行状态。
func (c *Client) DescribeCloudNativeAPIGatewayLatestTaskPhases(request *DescribeCloudNativeAPIGatewayLatestTaskPhasesRequest) (response *DescribeCloudNativeAPIGatewayLatestTaskPhasesResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayLatestTaskPhasesRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayLatestTaskPhasesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNacosReplicasRequest() (request *DescribeNacosReplicasRequest) {
	request = &DescribeNacosReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeNacosReplicas")
	return
}

func NewDescribeNacosReplicasResponse() (response *DescribeNacosReplicasResponse) {
	response = &DescribeNacosReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Nacos类型引擎实例副本信息
func (c *Client) DescribeNacosReplicas(request *DescribeNacosReplicasRequest) (response *DescribeNacosReplicasResponse, err error) {
	if request == nil {
		request = NewDescribeNacosReplicasRequest()
	}
	response = NewDescribeNacosReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNativeGatewayServiceGroupRequest() (request *DeleteNativeGatewayServiceGroupRequest) {
	request = &DeleteNativeGatewayServiceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteNativeGatewayServiceGroup")
	return
}

func NewDeleteNativeGatewayServiceGroupResponse() (response *DeleteNativeGatewayServiceGroupResponse) {
	response = &DeleteNativeGatewayServiceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网关服务实例分组
func (c *Client) DeleteNativeGatewayServiceGroup(request *DeleteNativeGatewayServiceGroupRequest) (response *DeleteNativeGatewayServiceGroupResponse, err error) {
	if request == nil {
		request = NewDeleteNativeGatewayServiceGroupRequest()
	}
	response = NewDeleteNativeGatewayServiceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperateWhiteListRequest() (request *DeleteOperateWhiteListRequest) {
	request = &DeleteOperateWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteOperateWhiteList")
	return
}

func NewDeleteOperateWhiteListResponse() (response *DeleteOperateWhiteListResponse) {
	response = &DeleteOperateWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运营平台白名单
func (c *Client) DeleteOperateWhiteList(request *DeleteOperateWhiteListRequest) (response *DeleteOperateWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteOperateWhiteListRequest()
	}
	response = NewDeleteOperateWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKongIngressControllerK8sVersionRequest() (request *DescribeKongIngressControllerK8sVersionRequest) {
	request = &DescribeKongIngressControllerK8sVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeKongIngressControllerK8sVersion")
	return
}

func NewDescribeKongIngressControllerK8sVersionResponse() (response *DescribeKongIngressControllerK8sVersionResponse) {
	response = &DescribeKongIngressControllerK8sVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 Kong Ingress Controller 与 K8s 版本映射
func (c *Client) DescribeKongIngressControllerK8sVersion(request *DescribeKongIngressControllerK8sVersionRequest) (response *DescribeKongIngressControllerK8sVersionResponse, err error) {
	if request == nil {
		request = NewDescribeKongIngressControllerK8sVersionRequest()
	}
	response = NewDescribeKongIngressControllerK8sVersionResponse()
	err = c.Send(request, response)
	return
}

func NewChangeOperateApolloDeployTypeRequest() (request *ChangeOperateApolloDeployTypeRequest) {
	request = &ChangeOperateApolloDeployTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ChangeOperateApolloDeployType")
	return
}

func NewChangeOperateApolloDeployTypeResponse() (response *ChangeOperateApolloDeployTypeResponse) {
	response = &ChangeOperateApolloDeployTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过运营后台扩容Apollo引擎statefulset类型，实现日志持久化
func (c *Client) ChangeOperateApolloDeployType(request *ChangeOperateApolloDeployTypeRequest) (response *ChangeOperateApolloDeployTypeResponse, err error) {
	if request == nil {
		request = NewChangeOperateApolloDeployTypeRequest()
	}
	response = NewChangeOperateApolloDeployTypeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCloudNativeAPIGatewayRouteRequest() (request *CreateCloudNativeAPIGatewayRouteRequest) {
	request = &CreateCloudNativeAPIGatewayRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayRoute")
	return
}

func NewCreateCloudNativeAPIGatewayRouteResponse() (response *CreateCloudNativeAPIGatewayRouteResponse) {
	response = &CreateCloudNativeAPIGatewayRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建云原生网关路由
func (c *Client) CreateCloudNativeAPIGatewayRoute(request *CreateCloudNativeAPIGatewayRouteRequest) (response *CreateCloudNativeAPIGatewayRouteResponse, err error) {
	if request == nil {
		request = NewCreateCloudNativeAPIGatewayRouteRequest()
	}
	response = NewCreateCloudNativeAPIGatewayRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFeatureVersionInfosRequest() (request *DescribeFeatureVersionInfosRequest) {
	request = &DescribeFeatureVersionInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeFeatureVersionInfos")
	return
}

func NewDescribeFeatureVersionInfosResponse() (response *DescribeFeatureVersionInfosResponse) {
	response = &DescribeFeatureVersionInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询功能版本列表
func (c *Client) DescribeFeatureVersionInfos(request *DescribeFeatureVersionInfosRequest) (response *DescribeFeatureVersionInfosResponse, err error) {
	if request == nil {
		request = NewDescribeFeatureVersionInfosRequest()
	}
	response = NewDescribeFeatureVersionInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceStrategiesRequest() (request *DescribeGovernanceStrategiesRequest) {
	request = &DescribeGovernanceStrategiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceStrategies")
	return
}

func NewDescribeGovernanceStrategiesResponse() (response *DescribeGovernanceStrategiesResponse) {
	response = &DescribeGovernanceStrategiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心鉴权策略列表
func (c *Client) DescribeGovernanceStrategies(request *DescribeGovernanceStrategiesRequest) (response *DescribeGovernanceStrategiesResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceStrategiesRequest()
	}
	response = NewDescribeGovernanceStrategiesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCloudNativeAPIGatewayCertificateRequest() (request *CreateCloudNativeAPIGatewayCertificateRequest) {
	request = &CreateCloudNativeAPIGatewayCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayCertificate")
	return
}

func NewCreateCloudNativeAPIGatewayCertificateResponse() (response *CreateCloudNativeAPIGatewayCertificateResponse) {
	response = &CreateCloudNativeAPIGatewayCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建云原生网关证书
func (c *Client) CreateCloudNativeAPIGatewayCertificate(request *CreateCloudNativeAPIGatewayCertificateRequest) (response *CreateCloudNativeAPIGatewayCertificateResponse, err error) {
	if request == nil {
		request = NewCreateCloudNativeAPIGatewayCertificateRequest()
	}
	response = NewCreateCloudNativeAPIGatewayCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNativeGatewayServerGroupsRequest() (request *DescribeNativeGatewayServerGroupsRequest) {
	request = &DescribeNativeGatewayServerGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeNativeGatewayServerGroups")
	return
}

func NewDescribeNativeGatewayServerGroupsResponse() (response *DescribeNativeGatewayServerGroupsResponse) {
	response = &DescribeNativeGatewayServerGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云原生网关分组信息
func (c *Client) DescribeNativeGatewayServerGroups(request *DescribeNativeGatewayServerGroupsRequest) (response *DescribeNativeGatewayServerGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeNativeGatewayServerGroupsRequest()
	}
	response = NewDescribeNativeGatewayServerGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayConfigRequest() (request *DescribeCloudNativeAPIGatewayConfigRequest) {
	request = &DescribeCloudNativeAPIGatewayConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayConfig")
	return
}

func NewDescribeCloudNativeAPIGatewayConfigResponse() (response *DescribeCloudNativeAPIGatewayConfigResponse) {
	response = &DescribeCloudNativeAPIGatewayConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关实例配置信息
func (c *Client) DescribeCloudNativeAPIGatewayConfig(request *DescribeCloudNativeAPIGatewayConfigRequest) (response *DescribeCloudNativeAPIGatewayConfigResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayConfigRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigFileRequest() (request *DescribeConfigFileRequest) {
	request = &DescribeConfigFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFile")
	return
}

func NewDescribeConfigFileResponse() (response *DescribeConfigFileResponse) {
	response = &DescribeConfigFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据命名空间、组、名字查找配置文件
func (c *Client) DescribeConfigFile(request *DescribeConfigFileRequest) (response *DescribeConfigFileResponse, err error) {
	if request == nil {
		request = NewDescribeConfigFileRequest()
	}
	response = NewDescribeConfigFileResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGovernedK8SRequest() (request *DeleteGovernedK8SRequest) {
	request = &DeleteGovernedK8SRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernedK8S")
	return
}

func NewDeleteGovernedK8SResponse() (response *DeleteGovernedK8SResponse) {
	response = &DeleteGovernedK8SResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除服务治理的kubernetes集群
func (c *Client) DeleteGovernedK8S(request *DeleteGovernedK8SRequest) (response *DeleteGovernedK8SResponse, err error) {
	if request == nil {
		request = NewDeleteGovernedK8SRequest()
	}
	response = NewDeleteGovernedK8SResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayRequest() (request *DescribeCloudNativeAPIGatewayRequest) {
	request = &DescribeCloudNativeAPIGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGateway")
	return
}

func NewDescribeCloudNativeAPIGatewayResponse() (response *DescribeCloudNativeAPIGatewayResponse) {
	response = &DescribeCloudNativeAPIGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关实例信息
func (c *Client) DescribeCloudNativeAPIGateway(request *DescribeCloudNativeAPIGatewayRequest) (response *DescribeCloudNativeAPIGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewaySystemPluginsRequest() (request *DescribeCloudNativeAPIGatewaySystemPluginsRequest) {
	request = &DescribeCloudNativeAPIGatewaySystemPluginsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewaySystemPlugins")
	return
}

func NewDescribeCloudNativeAPIGatewaySystemPluginsResponse() (response *DescribeCloudNativeAPIGatewaySystemPluginsResponse) {
	response = &DescribeCloudNativeAPIGatewaySystemPluginsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云原生网关系统插件列表
func (c *Client) DescribeCloudNativeAPIGatewaySystemPlugins(request *DescribeCloudNativeAPIGatewaySystemPluginsRequest) (response *DescribeCloudNativeAPIGatewaySystemPluginsResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewaySystemPluginsRequest()
	}
	response = NewDescribeCloudNativeAPIGatewaySystemPluginsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigFileReleaseRequest() (request *DescribeConfigFileReleaseRequest) {
	request = &DescribeConfigFileReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFileRelease")
	return
}

func NewDescribeConfigFileReleaseResponse() (response *DescribeConfigFileReleaseResponse) {
	response = &DescribeConfigFileReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置文件发布
func (c *Client) DescribeConfigFileRelease(request *DescribeConfigFileReleaseRequest) (response *DescribeConfigFileReleaseResponse, err error) {
	if request == nil {
		request = NewDescribeConfigFileReleaseRequest()
	}
	response = NewDescribeConfigFileReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoScalerResourceStrategyRequest() (request *ModifyAutoScalerResourceStrategyRequest) {
	request = &ModifyAutoScalerResourceStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyAutoScalerResourceStrategy")
	return
}

func NewModifyAutoScalerResourceStrategyResponse() (response *ModifyAutoScalerResourceStrategyResponse) {
	response = &ModifyAutoScalerResourceStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新弹性伸缩策略
func (c *Client) ModifyAutoScalerResourceStrategy(request *ModifyAutoScalerResourceStrategyRequest) (response *ModifyAutoScalerResourceStrategyResponse, err error) {
	if request == nil {
		request = NewModifyAutoScalerResourceStrategyRequest()
	}
	response = NewModifyAutoScalerResourceStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewManageApolloEnvRequest() (request *ManageApolloEnvRequest) {
	request = &ManageApolloEnvRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ManageApolloEnv")
	return
}

func NewManageApolloEnvResponse() (response *ManageApolloEnvResponse) {
	response = &ManageApolloEnvResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理Apollo环境
func (c *Client) ManageApolloEnv(request *ManageApolloEnvRequest) (response *ManageApolloEnvResponse, err error) {
	if request == nil {
		request = NewManageApolloEnvRequest()
	}
	response = NewManageApolloEnvResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGovernanceUsersRequest() (request *CreateGovernanceUsersRequest) {
	request = &CreateGovernanceUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceUsers")
	return
}

func NewCreateGovernanceUsersResponse() (response *CreateGovernanceUsersResponse) {
	response = &CreateGovernanceUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量创建治理中心用户
func (c *Client) CreateGovernanceUsers(request *CreateGovernanceUsersRequest) (response *CreateGovernanceUsersResponse, err error) {
	if request == nil {
		request = NewCreateGovernanceUsersRequest()
	}
	response = NewCreateGovernanceUsersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNativeGatewayServiceGroupRequest() (request *ModifyNativeGatewayServiceGroupRequest) {
	request = &ModifyNativeGatewayServiceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyNativeGatewayServiceGroup")
	return
}

func NewModifyNativeGatewayServiceGroupResponse() (response *ModifyNativeGatewayServiceGroupResponse) {
	response = &ModifyNativeGatewayServiceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关服务实例分组
func (c *Client) ModifyNativeGatewayServiceGroup(request *ModifyNativeGatewayServiceGroupRequest) (response *ModifyNativeGatewayServiceGroupResponse, err error) {
	if request == nil {
		request = NewModifyNativeGatewayServiceGroupRequest()
	}
	response = NewModifyNativeGatewayServiceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoScalerResourceStrategyBindingGroupsRequest() (request *DescribeAutoScalerResourceStrategyBindingGroupsRequest) {
	request = &DescribeAutoScalerResourceStrategyBindingGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeAutoScalerResourceStrategyBindingGroups")
	return
}

func NewDescribeAutoScalerResourceStrategyBindingGroupsResponse() (response *DescribeAutoScalerResourceStrategyBindingGroupsResponse) {
	response = &DescribeAutoScalerResourceStrategyBindingGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看弹性伸缩策略绑定的网关分组
func (c *Client) DescribeAutoScalerResourceStrategyBindingGroups(request *DescribeAutoScalerResourceStrategyBindingGroupsRequest) (response *DescribeAutoScalerResourceStrategyBindingGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeAutoScalerResourceStrategyBindingGroupsRequest()
	}
	response = NewDescribeAutoScalerResourceStrategyBindingGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudNativeAPIGatewayCertificateRequest() (request *ModifyCloudNativeAPIGatewayCertificateRequest) {
	request = &ModifyCloudNativeAPIGatewayCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayCertificate")
	return
}

func NewModifyCloudNativeAPIGatewayCertificateResponse() (response *ModifyCloudNativeAPIGatewayCertificateResponse) {
	response = &ModifyCloudNativeAPIGatewayCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生网关证书
func (c *Client) ModifyCloudNativeAPIGatewayCertificate(request *ModifyCloudNativeAPIGatewayCertificateRequest) (response *ModifyCloudNativeAPIGatewayCertificateResponse, err error) {
	if request == nil {
		request = NewModifyCloudNativeAPIGatewayCertificateRequest()
	}
	response = NewModifyCloudNativeAPIGatewayCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudNativeAPIGatewaySystemParametersRequest() (request *ModifyCloudNativeAPIGatewaySystemParametersRequest) {
	request = &ModifyCloudNativeAPIGatewaySystemParametersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewaySystemParameters")
	return
}

func NewModifyCloudNativeAPIGatewaySystemParametersResponse() (response *ModifyCloudNativeAPIGatewaySystemParametersResponse) {
	response = &ModifyCloudNativeAPIGatewaySystemParametersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生API网关系统参数
func (c *Client) ModifyCloudNativeAPIGatewaySystemParameters(request *ModifyCloudNativeAPIGatewaySystemParametersRequest) (response *ModifyCloudNativeAPIGatewaySystemParametersResponse, err error) {
	if request == nil {
		request = NewModifyCloudNativeAPIGatewaySystemParametersRequest()
	}
	response = NewModifyCloudNativeAPIGatewaySystemParametersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCloudNativeAPIGatewayServiceRequest() (request *CreateCloudNativeAPIGatewayServiceRequest) {
	request = &CreateCloudNativeAPIGatewayServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayService")
	return
}

func NewCreateCloudNativeAPIGatewayServiceResponse() (response *CreateCloudNativeAPIGatewayServiceResponse) {
	response = &CreateCloudNativeAPIGatewayServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建云原生网关服务
func (c *Client) CreateCloudNativeAPIGatewayService(request *CreateCloudNativeAPIGatewayServiceRequest) (response *CreateCloudNativeAPIGatewayServiceResponse, err error) {
	if request == nil {
		request = NewCreateCloudNativeAPIGatewayServiceRequest()
	}
	response = NewCreateCloudNativeAPIGatewayServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSREGlobalConfigsRequest() (request *DescribeSREGlobalConfigsRequest) {
	request = &DescribeSREGlobalConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeSREGlobalConfigs")
	return
}

func NewDescribeSREGlobalConfigsResponse() (response *DescribeSREGlobalConfigsResponse) {
	response = &DescribeSREGlobalConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询引擎全局配置信息，例如支持的规格，版本信息等
func (c *Client) DescribeSREGlobalConfigs(request *DescribeSREGlobalConfigsRequest) (response *DescribeSREGlobalConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeSREGlobalConfigsRequest()
	}
	response = NewDescribeSREGlobalConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigFileGroupRequest() (request *CreateConfigFileGroupRequest) {
	request = &CreateConfigFileGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateConfigFileGroup")
	return
}

func NewCreateConfigFileGroupResponse() (response *CreateConfigFileGroupResponse) {
	response = &CreateConfigFileGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建服务治理中心配置文件组
func (c *Client) CreateConfigFileGroup(request *CreateConfigFileGroupRequest) (response *CreateConfigFileGroupResponse, err error) {
	if request == nil {
		request = NewCreateConfigFileGroupRequest()
	}
	response = NewCreateConfigFileGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSRELatestTaskPhasesRequest() (request *DescribeSRELatestTaskPhasesRequest) {
	request = &DescribeSRELatestTaskPhasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeSRELatestTaskPhases")
	return
}

func NewDescribeSRELatestTaskPhasesResponse() (response *DescribeSRELatestTaskPhasesResponse) {
	response = &DescribeSRELatestTaskPhasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当执行部分流程较长的任务时，任务会被拆分为多个阶段进行异步执行。此接口用于查询最新提交的任务的执行状态。
func (c *Client) DescribeSRELatestTaskPhases(request *DescribeSRELatestTaskPhasesRequest) (response *DescribeSRELatestTaskPhasesResponse, err error) {
	if request == nil {
		request = NewDescribeSRELatestTaskPhasesRequest()
	}
	response = NewDescribeSRELatestTaskPhasesResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewayPluginVersionsRequest() (request *ListCloudNativeAPIGatewayPluginVersionsRequest) {
	request = &ListCloudNativeAPIGatewayPluginVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGatewayPluginVersions")
	return
}

func NewListCloudNativeAPIGatewayPluginVersionsResponse() (response *ListCloudNativeAPIGatewayPluginVersionsResponse) {
	response = &ListCloudNativeAPIGatewayPluginVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关插件版本信息列表
func (c *Client) ListCloudNativeAPIGatewayPluginVersions(request *ListCloudNativeAPIGatewayPluginVersionsRequest) (response *ListCloudNativeAPIGatewayPluginVersionsResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewayPluginVersionsRequest()
	}
	response = NewListCloudNativeAPIGatewayPluginVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEurekaServiceInstancesRequest() (request *DescribeEurekaServiceInstancesRequest) {
	request = &DescribeEurekaServiceInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEurekaServiceInstances")
	return
}

func NewDescribeEurekaServiceInstancesResponse() (response *DescribeEurekaServiceInstancesResponse) {
	response = &DescribeEurekaServiceInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定实例下eureka注册服务的实例列表
func (c *Client) DescribeEurekaServiceInstances(request *DescribeEurekaServiceInstancesRequest) (response *DescribeEurekaServiceInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeEurekaServiceInstancesRequest()
	}
	response = NewDescribeEurekaServiceInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceTagInfosRequest() (request *ModifyInstanceTagInfosRequest) {
	request = &ModifyInstanceTagInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyInstanceTagInfos")
	return
}

func NewModifyInstanceTagInfosResponse() (response *ModifyInstanceTagInfosResponse) {
	response = &ModifyInstanceTagInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例的标签信息
func (c *Client) ModifyInstanceTagInfos(request *ModifyInstanceTagInfosRequest) (response *ModifyInstanceTagInfosResponse, err error) {
	if request == nil {
		request = NewModifyInstanceTagInfosRequest()
	}
	response = NewModifyInstanceTagInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayUserStatusRequest() (request *DescribeCloudNativeAPIGatewayUserStatusRequest) {
	request = &DescribeCloudNativeAPIGatewayUserStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayUserStatus")
	return
}

func NewDescribeCloudNativeAPIGatewayUserStatusResponse() (response *DescribeCloudNativeAPIGatewayUserStatusResponse) {
	response = &DescribeCloudNativeAPIGatewayUserStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关用户状态
func (c *Client) DescribeCloudNativeAPIGatewayUserStatus(request *DescribeCloudNativeAPIGatewayUserStatusRequest) (response *DescribeCloudNativeAPIGatewayUserStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayUserStatusRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayUserStatusResponse()
	err = c.Send(request, response)
	return
}

func NewManageZookeeperConfigRequest() (request *ManageZookeeperConfigRequest) {
	request = &ManageZookeeperConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ManageZookeeperConfig")
	return
}

func NewManageZookeeperConfigResponse() (response *ManageZookeeperConfigResponse) {
	response = &ManageZookeeperConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理zookeeper配置
func (c *Client) ManageZookeeperConfig(request *ManageZookeeperConfigRequest) (response *ManageZookeeperConfigResponse, err error) {
	if request == nil {
		request = NewManageZookeeperConfigRequest()
	}
	response = NewManageZookeeperConfigResponse()
	err = c.Send(request, response)
	return
}

func NewBindCLBSecurityGroupRequest() (request *BindCLBSecurityGroupRequest) {
	request = &BindCLBSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "BindCLBSecurityGroup")
	return
}

func NewBindCLBSecurityGroupResponse() (response *BindCLBSecurityGroupResponse) {
	response = &BindCLBSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给 clb 绑定安全组
func (c *Client) BindCLBSecurityGroup(request *BindCLBSecurityGroupRequest) (response *BindCLBSecurityGroupResponse, err error) {
	if request == nil {
		request = NewBindCLBSecurityGroupRequest()
	}
	response = NewBindCLBSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCheckZookeeperMigrateConnectionRequest() (request *CheckZookeeperMigrateConnectionRequest) {
	request = &CheckZookeeperMigrateConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CheckZookeeperMigrateConnection")
	return
}

func NewCheckZookeeperMigrateConnectionResponse() (response *CheckZookeeperMigrateConnectionResponse) {
	response = &CheckZookeeperMigrateConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 迁移准备阶段检查连通性
func (c *Client) CheckZookeeperMigrateConnection(request *CheckZookeeperMigrateConnectionRequest) (response *CheckZookeeperMigrateConnectionResponse, err error) {
	if request == nil {
		request = NewCheckZookeeperMigrateConnectionRequest()
	}
	response = NewCheckZookeeperMigrateConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewEnableEngineInternetRequest() (request *EnableEngineInternetRequest) {
	request = &EnableEngineInternetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "EnableEngineInternet")
	return
}

func NewEnableEngineInternetResponse() (response *EnableEngineInternetResponse) {
	response = &EnableEngineInternetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 引擎公网开关
func (c *Client) EnableEngineInternet(request *EnableEngineInternetRequest) (response *EnableEngineInternetResponse, err error) {
	if request == nil {
		request = NewEnableEngineInternetRequest()
	}
	response = NewEnableEngineInternetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceUserRequest() (request *ModifyGovernanceUserRequest) {
	request = &ModifyGovernanceUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceUser")
	return
}

func NewModifyGovernanceUserResponse() (response *ModifyGovernanceUserResponse) {
	response = &ModifyGovernanceUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改治理中心用户信息
func (c *Client) ModifyGovernanceUser(request *ModifyGovernanceUserRequest) (response *ModifyGovernanceUserResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceUserRequest()
	}
	response = NewModifyGovernanceUserResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNativeGatewayServerGroupRequest() (request *ModifyNativeGatewayServerGroupRequest) {
	request = &ModifyNativeGatewayServerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyNativeGatewayServerGroup")
	return
}

func NewModifyNativeGatewayServerGroupResponse() (response *ModifyNativeGatewayServerGroupResponse) {
	response = &ModifyNativeGatewayServerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生API网关实例分组基础信息
func (c *Client) ModifyNativeGatewayServerGroup(request *ModifyNativeGatewayServerGroupRequest) (response *ModifyNativeGatewayServerGroupResponse, err error) {
	if request == nil {
		request = NewModifyNativeGatewayServerGroupRequest()
	}
	response = NewModifyNativeGatewayServerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceNamespacesRequest() (request *ModifyGovernanceNamespacesRequest) {
	request = &ModifyGovernanceNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceNamespaces")
	return
}

func NewModifyGovernanceNamespacesResponse() (response *ModifyGovernanceNamespacesResponse) {
	response = &ModifyGovernanceNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改治理中心命名空间
func (c *Client) ModifyGovernanceNamespaces(request *ModifyGovernanceNamespacesRequest) (response *ModifyGovernanceNamespacesResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceNamespacesRequest()
	}
	response = NewModifyGovernanceNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAutoScalerResourceStrategyRequest() (request *DeleteAutoScalerResourceStrategyRequest) {
	request = &DeleteAutoScalerResourceStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteAutoScalerResourceStrategy")
	return
}

func NewDeleteAutoScalerResourceStrategyResponse() (response *DeleteAutoScalerResourceStrategyResponse) {
	response = &DeleteAutoScalerResourceStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除弹性伸缩策略
func (c *Client) DeleteAutoScalerResourceStrategy(request *DeleteAutoScalerResourceStrategyRequest) (response *DeleteAutoScalerResourceStrategyResponse, err error) {
	if request == nil {
		request = NewDeleteAutoScalerResourceStrategyRequest()
	}
	response = NewDeleteAutoScalerResourceStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateSpecListRequest() (request *DescribeOperateSpecListRequest) {
	request = &DescribeOperateSpecListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateSpecList")
	return
}

func NewDescribeOperateSpecListResponse() (response *DescribeOperateSpecListResponse) {
	response = &DescribeOperateSpecListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台规格列表
func (c *Client) DescribeOperateSpecList(request *DescribeOperateSpecListRequest) (response *DescribeOperateSpecListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateSpecListRequest()
	}
	response = NewDescribeOperateSpecListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperServiceInstanceCountRequest() (request *DescribeZookeeperServiceInstanceCountRequest) {
	request = &DescribeZookeeperServiceInstanceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperServiceInstanceCount")
	return
}

func NewDescribeZookeeperServiceInstanceCountResponse() (response *DescribeZookeeperServiceInstanceCountResponse) {
	response = &DescribeZookeeperServiceInstanceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询zookeeper服务实例个数
func (c *Client) DescribeZookeeperServiceInstanceCount(request *DescribeZookeeperServiceInstanceCountRequest) (response *DescribeZookeeperServiceInstanceCountResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperServiceInstanceCountRequest()
	}
	response = NewDescribeZookeeperServiceInstanceCountResponse()
	err = c.Send(request, response)
	return
}

func NewGetCloudNativeAPIGatewayPluginUploadInfoRequest() (request *GetCloudNativeAPIGatewayPluginUploadInfoRequest) {
	request = &GetCloudNativeAPIGatewayPluginUploadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "GetCloudNativeAPIGatewayPluginUploadInfo")
	return
}

func NewGetCloudNativeAPIGatewayPluginUploadInfoResponse() (response *GetCloudNativeAPIGatewayPluginUploadInfoResponse) {
	response = &GetCloudNativeAPIGatewayPluginUploadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上传插件前置操作，获取COS相关信息
func (c *Client) GetCloudNativeAPIGatewayPluginUploadInfo(request *GetCloudNativeAPIGatewayPluginUploadInfoRequest) (response *GetCloudNativeAPIGatewayPluginUploadInfoResponse, err error) {
	if request == nil {
		request = NewGetCloudNativeAPIGatewayPluginUploadInfoRequest()
	}
	response = NewGetCloudNativeAPIGatewayPluginUploadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCheckUserCloudProductRequest() (request *CheckUserCloudProductRequest) {
	request = &CheckUserCloudProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CheckUserCloudProduct")
	return
}

func NewCheckUserCloudProductResponse() (response *CheckUserCloudProductResponse) {
	response = &CheckUserCloudProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查用户是否开启了相关云产品
func (c *Client) CheckUserCloudProduct(request *CheckUserCloudProductRequest) (response *CheckUserCloudProductResponse, err error) {
	if request == nil {
		request = NewCheckUserCloudProductRequest()
	}
	response = NewCheckUserCloudProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllConfigFileTemplatesRequest() (request *DescribeAllConfigFileTemplatesRequest) {
	request = &DescribeAllConfigFileTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeAllConfigFileTemplates")
	return
}

func NewDescribeAllConfigFileTemplatesResponse() (response *DescribeAllConfigFileTemplatesResponse) {
	response = &DescribeAllConfigFileTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全量配置文件模板列表
func (c *Client) DescribeAllConfigFileTemplates(request *DescribeAllConfigFileTemplatesRequest) (response *DescribeAllConfigFileTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeAllConfigFileTemplatesRequest()
	}
	response = NewDescribeAllConfigFileTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewEnableConsoleAccessAddressRequest() (request *EnableConsoleAccessAddressRequest) {
	request = &EnableConsoleAccessAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "EnableConsoleAccessAddress")
	return
}

func NewEnableConsoleAccessAddressResponse() (response *EnableConsoleAccessAddressResponse) {
	response = &EnableConsoleAccessAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 引擎控制台访问地址开关
func (c *Client) EnableConsoleAccessAddress(request *EnableConsoleAccessAddressRequest) (response *EnableConsoleAccessAddressResponse, err error) {
	if request == nil {
		request = NewEnableConsoleAccessAddressRequest()
	}
	response = NewEnableConsoleAccessAddressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNativeGatewayServiceSourcesRequest() (request *DescribeNativeGatewayServiceSourcesRequest) {
	request = &DescribeNativeGatewayServiceSourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeNativeGatewayServiceSources")
	return
}

func NewDescribeNativeGatewayServiceSourcesResponse() (response *DescribeNativeGatewayServiceSourcesResponse) {
	response = &DescribeNativeGatewayServiceSourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关服务来源实例列表
func (c *Client) DescribeNativeGatewayServiceSources(request *DescribeNativeGatewayServiceSourcesRequest) (response *DescribeNativeGatewayServiceSourcesResponse, err error) {
	if request == nil {
		request = NewDescribeNativeGatewayServiceSourcesRequest()
	}
	response = NewDescribeNativeGatewayServiceSourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGovernanceServicesRequest() (request *DeleteGovernanceServicesRequest) {
	request = &DeleteGovernanceServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceServices")
	return
}

func NewDeleteGovernanceServicesResponse() (response *DeleteGovernanceServicesResponse) {
	response = &DeleteGovernanceServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除治理中心服务
func (c *Client) DeleteGovernanceServices(request *DeleteGovernanceServicesRequest) (response *DeleteGovernanceServicesResponse, err error) {
	if request == nil {
		request = NewDeleteGovernanceServicesRequest()
	}
	response = NewDeleteGovernanceServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateInstanceDetailRequest() (request *DescribeOperateInstanceDetailRequest) {
	request = &DescribeOperateInstanceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateInstanceDetail")
	return
}

func NewDescribeOperateInstanceDetailResponse() (response *DescribeOperateInstanceDetailResponse) {
	response = &DescribeOperateInstanceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台引擎实例详情
func (c *Client) DescribeOperateInstanceDetail(request *DescribeOperateInstanceDetailRequest) (response *DescribeOperateInstanceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeOperateInstanceDetailRequest()
	}
	response = NewDescribeOperateInstanceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetNginxIngressNginxRawConfRequest() (request *GetNginxIngressNginxRawConfRequest) {
	request = &GetNginxIngressNginxRawConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "GetNginxIngressNginxRawConf")
	return
}

func NewGetNginxIngressNginxRawConfResponse() (response *GetNginxIngressNginxRawConfResponse) {
	response = &GetNginxIngressNginxRawConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取 nginx 原始的 nginx.conf
func (c *Client) GetNginxIngressNginxRawConf(request *GetNginxIngressNginxRawConfRequest) (response *GetNginxIngressNginxRawConfResponse, err error) {
	if request == nil {
		request = NewGetNginxIngressNginxRawConfRequest()
	}
	response = NewGetNginxIngressNginxRawConfResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudNativeAPIGatewayNetworkRequest() (request *ModifyCloudNativeAPIGatewayNetworkRequest) {
	request = &ModifyCloudNativeAPIGatewayNetworkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayNetwork")
	return
}

func NewModifyCloudNativeAPIGatewayNetworkResponse() (response *ModifyCloudNativeAPIGatewayNetworkResponse) {
	response = &ModifyCloudNativeAPIGatewayNetworkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生API网关实例网络配置，比如带宽以及网络访问策略。
func (c *Client) ModifyCloudNativeAPIGatewayNetwork(request *ModifyCloudNativeAPIGatewayNetworkRequest) (response *ModifyCloudNativeAPIGatewayNetworkResponse, err error) {
	if request == nil {
		request = NewModifyCloudNativeAPIGatewayNetworkRequest()
	}
	response = NewModifyCloudNativeAPIGatewayNetworkResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateRegionRequest() (request *ModifyOperateRegionRequest) {
	request = &ModifyOperateRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyOperateRegion")
	return
}

func NewModifyOperateRegionResponse() (response *ModifyOperateRegionResponse) {
	response = &ModifyOperateRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新运营平台地域
func (c *Client) ModifyOperateRegion(request *ModifyOperateRegionRequest) (response *ModifyOperateRegionResponse, err error) {
	if request == nil {
		request = NewModifyOperateRegionRequest()
	}
	response = NewModifyOperateRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewayCertificateRequest() (request *DeleteCloudNativeAPIGatewayCertificateRequest) {
	request = &DeleteCloudNativeAPIGatewayCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayCertificate")
	return
}

func NewDeleteCloudNativeAPIGatewayCertificateResponse() (response *DeleteCloudNativeAPIGatewayCertificateResponse) {
	response = &DeleteCloudNativeAPIGatewayCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除云原生网关证书
func (c *Client) DeleteCloudNativeAPIGatewayCertificate(request *DeleteCloudNativeAPIGatewayCertificateRequest) (response *DeleteCloudNativeAPIGatewayCertificateResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewayCertificateRequest()
	}
	response = NewDeleteCloudNativeAPIGatewayCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceGroupTokenRequest() (request *DescribeGovernanceGroupTokenRequest) {
	request = &DescribeGovernanceGroupTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceGroupToken")
	return
}

func NewDescribeGovernanceGroupTokenResponse() (response *DescribeGovernanceGroupTokenResponse) {
	response = &DescribeGovernanceGroupTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心用户组Token
func (c *Client) DescribeGovernanceGroupToken(request *DescribeGovernanceGroupTokenRequest) (response *DescribeGovernanceGroupTokenResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceGroupTokenRequest()
	}
	response = NewDescribeGovernanceGroupTokenResponse()
	err = c.Send(request, response)
	return
}

func NewManageVpcEndPointRequest() (request *ManageVpcEndPointRequest) {
	request = &ManageVpcEndPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ManageVpcEndPoint")
	return
}

func NewManageVpcEndPointResponse() (response *ManageVpcEndPointResponse) {
	response = &ManageVpcEndPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理VPC终端节点
func (c *Client) ManageVpcEndPoint(request *ManageVpcEndPointRequest) (response *ManageVpcEndPointResponse, err error) {
	if request == nil {
		request = NewManageVpcEndPointRequest()
	}
	response = NewManageVpcEndPointResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAutoScalerResourceStrategyRequest() (request *CreateAutoScalerResourceStrategyRequest) {
	request = &CreateAutoScalerResourceStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateAutoScalerResourceStrategy")
	return
}

func NewCreateAutoScalerResourceStrategyResponse() (response *CreateAutoScalerResourceStrategyResponse) {
	response = &CreateAutoScalerResourceStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建弹性伸缩策略
func (c *Client) CreateAutoScalerResourceStrategy(request *CreateAutoScalerResourceStrategyRequest) (response *CreateAutoScalerResourceStrategyResponse, err error) {
	if request == nil {
		request = NewCreateAutoScalerResourceStrategyRequest()
	}
	response = NewCreateAutoScalerResourceStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExecTaskRequest() (request *CreateExecTaskRequest) {
	request = &CreateExecTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateExecTask")
	return
}

func NewCreateExecTaskResponse() (response *CreateExecTaskResponse) {
	response = &CreateExecTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运营平台可执行任务
func (c *Client) CreateExecTask(request *CreateExecTaskRequest) (response *CreateExecTaskResponse, err error) {
	if request == nil {
		request = NewCreateExecTaskRequest()
	}
	response = NewCreateExecTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEngineSyncJobsRequest() (request *DescribeEngineSyncJobsRequest) {
	request = &DescribeEngineSyncJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEngineSyncJobs")
	return
}

func NewDescribeEngineSyncJobsResponse() (response *DescribeEngineSyncJobsResponse) {
	response = &DescribeEngineSyncJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询引擎数据同步任务
func (c *Client) DescribeEngineSyncJobs(request *DescribeEngineSyncJobsRequest) (response *DescribeEngineSyncJobsResponse, err error) {
	if request == nil {
		request = NewDescribeEngineSyncJobsRequest()
	}
	response = NewDescribeEngineSyncJobsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateSecurityGroupListRequest() (request *DescribeOperateSecurityGroupListRequest) {
	request = &DescribeOperateSecurityGroupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateSecurityGroupList")
	return
}

func NewDescribeOperateSecurityGroupListResponse() (response *DescribeOperateSecurityGroupListResponse) {
	response = &DescribeOperateSecurityGroupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台安全组列表
func (c *Client) DescribeOperateSecurityGroupList(request *DescribeOperateSecurityGroupListRequest) (response *DescribeOperateSecurityGroupListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateSecurityGroupListRequest()
	}
	response = NewDescribeOperateSecurityGroupListResponse()
	err = c.Send(request, response)
	return
}

func NewEnablePolarisCLSRequest() (request *EnablePolarisCLSRequest) {
	request = &EnablePolarisCLSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "EnablePolarisCLS")
	return
}

func NewEnablePolarisCLSResponse() (response *EnablePolarisCLSResponse) {
	response = &EnablePolarisCLSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 控制 polaris 日志上报到 CLS 的能力
func (c *Client) EnablePolarisCLS(request *EnablePolarisCLSRequest) (response *EnablePolarisCLSResponse, err error) {
	if request == nil {
		request = NewEnablePolarisCLSRequest()
	}
	response = NewEnablePolarisCLSResponse()
	err = c.Send(request, response)
	return
}

func NewResetGovernanceGroupTokenRequest() (request *ResetGovernanceGroupTokenRequest) {
	request = &ResetGovernanceGroupTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ResetGovernanceGroupToken")
	return
}

func NewResetGovernanceGroupTokenResponse() (response *ResetGovernanceGroupTokenResponse) {
	response = &ResetGovernanceGroupTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置治理中心用户组Token
func (c *Client) ResetGovernanceGroupToken(request *ResetGovernanceGroupTokenRequest) (response *ResetGovernanceGroupTokenResponse, err error) {
	if request == nil {
		request = NewResetGovernanceGroupTokenRequest()
	}
	response = NewResetGovernanceGroupTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayPluginVersionsRequest() (request *DescribeCloudNativeAPIGatewayPluginVersionsRequest) {
	request = &DescribeCloudNativeAPIGatewayPluginVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayPluginVersions")
	return
}

func NewDescribeCloudNativeAPIGatewayPluginVersionsResponse() (response *DescribeCloudNativeAPIGatewayPluginVersionsResponse) {
	response = &DescribeCloudNativeAPIGatewayPluginVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关插件版本信息列表
func (c *Client) DescribeCloudNativeAPIGatewayPluginVersions(request *DescribeCloudNativeAPIGatewayPluginVersionsRequest) (response *DescribeCloudNativeAPIGatewayPluginVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayPluginVersionsRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayPluginVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserStatusRequest() (request *DescribeUserStatusRequest) {
	request = &DescribeUserStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeUserStatus")
	return
}

func NewDescribeUserStatusResponse() (response *DescribeUserStatusResponse) {
	response = &DescribeUserStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户状态
func (c *Client) DescribeUserStatus(request *DescribeUserStatusRequest) (response *DescribeUserStatusResponse, err error) {
	if request == nil {
		request = NewDescribeUserStatusRequest()
	}
	response = NewDescribeUserStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDestroySREInstanceRequest() (request *DestroySREInstanceRequest) {
	request = &DestroySREInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DestroySREInstance")
	return
}

func NewDestroySREInstanceResponse() (response *DestroySREInstanceResponse) {
	response = &DestroySREInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁微服务注册引擎实例
func (c *Client) DestroySREInstance(request *DestroySREInstanceRequest) (response *DestroySREInstanceResponse, err error) {
	if request == nil {
		request = NewDestroySREInstanceRequest()
	}
	response = NewDestroySREInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewayRoutesRequest() (request *ListCloudNativeAPIGatewayRoutesRequest) {
	request = &ListCloudNativeAPIGatewayRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGatewayRoutes")
	return
}

func NewListCloudNativeAPIGatewayRoutesResponse() (response *ListCloudNativeAPIGatewayRoutesResponse) {
	response = &ListCloudNativeAPIGatewayRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关路由列表
func (c *Client) ListCloudNativeAPIGatewayRoutes(request *ListCloudNativeAPIGatewayRoutesRequest) (response *ListCloudNativeAPIGatewayRoutesResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewayRoutesRequest()
	}
	response = NewListCloudNativeAPIGatewayRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNativeGatewayServerGroupRequest() (request *DeleteNativeGatewayServerGroupRequest) {
	request = &DeleteNativeGatewayServerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteNativeGatewayServerGroup")
	return
}

func NewDeleteNativeGatewayServerGroupResponse() (response *DeleteNativeGatewayServerGroupResponse) {
	response = &DeleteNativeGatewayServerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网关实例分组
func (c *Client) DeleteNativeGatewayServerGroup(request *DeleteNativeGatewayServerGroupRequest) (response *DeleteNativeGatewayServerGroupResponse, err error) {
	if request == nil {
		request = NewDeleteNativeGatewayServerGroupRequest()
	}
	response = NewDeleteNativeGatewayServerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConfigFilesRequest() (request *ModifyConfigFilesRequest) {
	request = &ModifyConfigFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyConfigFiles")
	return
}

func NewModifyConfigFilesResponse() (response *ModifyConfigFilesResponse) {
	response = &ModifyConfigFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改配置文件
func (c *Client) ModifyConfigFiles(request *ModifyConfigFilesRequest) (response *ModifyConfigFilesResponse, err error) {
	if request == nil {
		request = NewModifyConfigFilesRequest()
	}
	response = NewModifyConfigFilesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateFeatureVersionRequest() (request *ModifyOperateFeatureVersionRequest) {
	request = &ModifyOperateFeatureVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyOperateFeatureVersion")
	return
}

func NewModifyOperateFeatureVersionResponse() (response *ModifyOperateFeatureVersionResponse) {
	response = &ModifyOperateFeatureVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改功能版本
func (c *Client) ModifyOperateFeatureVersion(request *ModifyOperateFeatureVersionRequest) (response *ModifyOperateFeatureVersionResponse, err error) {
	if request == nil {
		request = NewModifyOperateFeatureVersionRequest()
	}
	response = NewModifyOperateFeatureVersionResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewayRequest() (request *ListCloudNativeAPIGatewayRequest) {
	request = &ListCloudNativeAPIGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGateway")
	return
}

func NewListCloudNativeAPIGatewayResponse() (response *ListCloudNativeAPIGatewayResponse) {
	response = &ListCloudNativeAPIGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关实例列表
func (c *Client) ListCloudNativeAPIGateway(request *ListCloudNativeAPIGatewayRequest) (response *ListCloudNativeAPIGatewayResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewayRequest()
	}
	response = NewListCloudNativeAPIGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceGroupTokenRequest() (request *ModifyGovernanceGroupTokenRequest) {
	request = &ModifyGovernanceGroupTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceGroupToken")
	return
}

func NewModifyGovernanceGroupTokenResponse() (response *ModifyGovernanceGroupTokenResponse) {
	response = &ModifyGovernanceGroupTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新治理中心用户组Token
func (c *Client) ModifyGovernanceGroupToken(request *ModifyGovernanceGroupTokenRequest) (response *ModifyGovernanceGroupTokenResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceGroupTokenRequest()
	}
	response = NewModifyGovernanceGroupTokenResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperateEngineSpecRequest() (request *CreateOperateEngineSpecRequest) {
	request = &CreateOperateEngineSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateOperateEngineSpec")
	return
}

func NewCreateOperateEngineSpecResponse() (response *CreateOperateEngineSpecResponse) {
	response = &CreateOperateEngineSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运营端引擎规格
func (c *Client) CreateOperateEngineSpec(request *CreateOperateEngineSpecRequest) (response *CreateOperateEngineSpecResponse, err error) {
	if request == nil {
		request = NewCreateOperateEngineSpecRequest()
	}
	response = NewCreateOperateEngineSpecResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewayServicesRequest() (request *ListCloudNativeAPIGatewayServicesRequest) {
	request = &ListCloudNativeAPIGatewayServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGatewayServices")
	return
}

func NewListCloudNativeAPIGatewayServicesResponse() (response *ListCloudNativeAPIGatewayServicesResponse) {
	response = &ListCloudNativeAPIGatewayServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关服务列表
func (c *Client) ListCloudNativeAPIGatewayServices(request *ListCloudNativeAPIGatewayServicesRequest) (response *ListCloudNativeAPIGatewayServicesResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewayServicesRequest()
	}
	response = NewListCloudNativeAPIGatewayServicesResponse()
	err = c.Send(request, response)
	return
}

func NewAddGovernedK8SRequest() (request *AddGovernedK8SRequest) {
	request = &AddGovernedK8SRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "AddGovernedK8S")
	return
}

func NewAddGovernedK8SResponse() (response *AddGovernedK8SResponse) {
	response = &AddGovernedK8SResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加服务治理的kubernetes集群
func (c *Client) AddGovernedK8S(request *AddGovernedK8SRequest) (response *AddGovernedK8SResponse, err error) {
	if request == nil {
		request = NewAddGovernedK8SRequest()
	}
	response = NewAddGovernedK8SResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExecTaskRequest() (request *DescribeExecTaskRequest) {
	request = &DescribeExecTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeExecTask")
	return
}

func NewDescribeExecTaskResponse() (response *DescribeExecTaskResponse) {
	response = &DescribeExecTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看运营平台可执行任务信息
func (c *Client) DescribeExecTask(request *DescribeExecTaskRequest) (response *DescribeExecTaskResponse, err error) {
	if request == nil {
		request = NewDescribeExecTaskRequest()
	}
	response = NewDescribeExecTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceStrategyRequest() (request *ModifyGovernanceStrategyRequest) {
	request = &ModifyGovernanceStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceStrategy")
	return
}

func NewModifyGovernanceStrategyResponse() (response *ModifyGovernanceStrategyResponse) {
	response = &ModifyGovernanceStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改治理中心鉴权策略
func (c *Client) ModifyGovernanceStrategy(request *ModifyGovernanceStrategyRequest) (response *ModifyGovernanceStrategyResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceStrategyRequest()
	}
	response = NewModifyGovernanceStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNativeGatewayServiceRequest() (request *DeleteNativeGatewayServiceRequest) {
	request = &DeleteNativeGatewayServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteNativeGatewayService")
	return
}

func NewDeleteNativeGatewayServiceResponse() (response *DeleteNativeGatewayServiceResponse) {
	response = &DeleteNativeGatewayServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网关服务
func (c *Client) DeleteNativeGatewayService(request *DeleteNativeGatewayServiceRequest) (response *DeleteNativeGatewayServiceResponse, err error) {
	if request == nil {
		request = NewDeleteNativeGatewayServiceRequest()
	}
	response = NewDeleteNativeGatewayServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperateApolloDeploymentRequest() (request *DeleteOperateApolloDeploymentRequest) {
	request = &DeleteOperateApolloDeploymentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteOperateApolloDeployment")
	return
}

func NewDeleteOperateApolloDeploymentResponse() (response *DeleteOperateApolloDeploymentResponse) {
	response = &DeleteOperateApolloDeploymentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过运营后台销毁Apollo引擎deployment类型（只销毁原3个组件）
func (c *Client) DeleteOperateApolloDeployment(request *DeleteOperateApolloDeploymentRequest) (response *DeleteOperateApolloDeploymentResponse, err error) {
	if request == nil {
		request = NewDeleteOperateApolloDeploymentRequest()
	}
	response = NewDeleteOperateApolloDeploymentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEngineExportPortsRequest() (request *DescribeEngineExportPortsRequest) {
	request = &DescribeEngineExportPortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEngineExportPorts")
	return
}

func NewDescribeEngineExportPortsResponse() (response *DescribeEngineExportPortsResponse) {
	response = &DescribeEngineExportPortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询引擎暴露的端口信息
func (c *Client) DescribeEngineExportPorts(request *DescribeEngineExportPortsRequest) (response *DescribeEngineExportPortsResponse, err error) {
	if request == nil {
		request = NewDescribeEngineExportPortsRequest()
	}
	response = NewDescribeEngineExportPortsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGovernanceAliasRequest() (request *CreateGovernanceAliasRequest) {
	request = &CreateGovernanceAliasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceAlias")
	return
}

func NewCreateGovernanceAliasResponse() (response *CreateGovernanceAliasResponse) {
	response = &CreateGovernanceAliasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建治理中心服务别名
func (c *Client) CreateGovernanceAlias(request *CreateGovernanceAliasRequest) (response *CreateGovernanceAliasResponse, err error) {
	if request == nil {
		request = NewCreateGovernanceAliasRequest()
	}
	response = NewCreateGovernanceAliasResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCloudNativeAPIGatewayCanaryRuleRequest() (request *CreateCloudNativeAPIGatewayCanaryRuleRequest) {
	request = &CreateCloudNativeAPIGatewayCanaryRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayCanaryRule")
	return
}

func NewCreateCloudNativeAPIGatewayCanaryRuleResponse() (response *CreateCloudNativeAPIGatewayCanaryRuleResponse) {
	response = &CreateCloudNativeAPIGatewayCanaryRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建云原生网关的灰度规则
func (c *Client) CreateCloudNativeAPIGatewayCanaryRule(request *CreateCloudNativeAPIGatewayCanaryRuleRequest) (response *CreateCloudNativeAPIGatewayCanaryRuleResponse, err error) {
	if request == nil {
		request = NewCreateCloudNativeAPIGatewayCanaryRuleRequest()
	}
	response = NewCreateCloudNativeAPIGatewayCanaryRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceGroupDetailRequest() (request *DescribeGovernanceGroupDetailRequest) {
	request = &DescribeGovernanceGroupDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceGroupDetail")
	return
}

func NewDescribeGovernanceGroupDetailResponse() (response *DescribeGovernanceGroupDetailResponse) {
	response = &DescribeGovernanceGroupDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心用户组详细
func (c *Client) DescribeGovernanceGroupDetail(request *DescribeGovernanceGroupDetailRequest) (response *DescribeGovernanceGroupDetailResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceGroupDetailRequest()
	}
	response = NewDescribeGovernanceGroupDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewaySecretRequest() (request *DescribeCloudNativeAPIGatewaySecretRequest) {
	request = &DescribeCloudNativeAPIGatewaySecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewaySecret")
	return
}

func NewDescribeCloudNativeAPIGatewaySecretResponse() (response *DescribeCloudNativeAPIGatewaySecretResponse) {
	response = &DescribeCloudNativeAPIGatewaySecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个TLS证书
func (c *Client) DescribeCloudNativeAPIGatewaySecret(request *DescribeCloudNativeAPIGatewaySecretRequest) (response *DescribeCloudNativeAPIGatewaySecretResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewaySecretRequest()
	}
	response = NewDescribeCloudNativeAPIGatewaySecretResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewayCertificatesRequest() (request *ListCloudNativeAPIGatewayCertificatesRequest) {
	request = &ListCloudNativeAPIGatewayCertificatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGatewayCertificates")
	return
}

func NewListCloudNativeAPIGatewayCertificatesResponse() (response *ListCloudNativeAPIGatewayCertificatesResponse) {
	response = &ListCloudNativeAPIGatewayCertificatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关证书列表
func (c *Client) ListCloudNativeAPIGatewayCertificates(request *ListCloudNativeAPIGatewayCertificatesRequest) (response *ListCloudNativeAPIGatewayCertificatesResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewayCertificatesRequest()
	}
	response = NewListCloudNativeAPIGatewayCertificatesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplyLicenseRequest() (request *CreateApplyLicenseRequest) {
	request = &CreateApplyLicenseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateApplyLicense")
	return
}

func NewCreateApplyLicenseResponse() (response *CreateApplyLicenseResponse) {
	response = &CreateApplyLicenseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建License申请请求
func (c *Client) CreateApplyLicense(request *CreateApplyLicenseRequest) (response *CreateApplyLicenseResponse, err error) {
	if request == nil {
		request = NewCreateApplyLicenseRequest()
	}
	response = NewCreateApplyLicenseResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewayServiceRequest() (request *DeleteCloudNativeAPIGatewayServiceRequest) {
	request = &DeleteCloudNativeAPIGatewayServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayService")
	return
}

func NewDeleteCloudNativeAPIGatewayServiceResponse() (response *DeleteCloudNativeAPIGatewayServiceResponse) {
	response = &DeleteCloudNativeAPIGatewayServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除云原生网关服务
func (c *Client) DeleteCloudNativeAPIGatewayService(request *DeleteCloudNativeAPIGatewayServiceRequest) (response *DeleteCloudNativeAPIGatewayServiceResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewayServiceRequest()
	}
	response = NewDeleteCloudNativeAPIGatewayServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePolarisCLSTopicRequest() (request *DescribePolarisCLSTopicRequest) {
	request = &DescribePolarisCLSTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribePolarisCLSTopic")
	return
}

func NewDescribePolarisCLSTopicResponse() (response *DescribePolarisCLSTopicResponse) {
	response = &DescribePolarisCLSTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Polaris的CLS日志topic列表
func (c *Client) DescribePolarisCLSTopic(request *DescribePolarisCLSTopicRequest) (response *DescribePolarisCLSTopicResponse, err error) {
	if request == nil {
		request = NewDescribePolarisCLSTopicRequest()
	}
	response = NewDescribePolarisCLSTopicResponse()
	err = c.Send(request, response)
	return
}

func NewGetCloudNativeAPIGatewaySecretsRequest() (request *GetCloudNativeAPIGatewaySecretsRequest) {
	request = &GetCloudNativeAPIGatewaySecretsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "GetCloudNativeAPIGatewaySecrets")
	return
}

func NewGetCloudNativeAPIGatewaySecretsResponse() (response *GetCloudNativeAPIGatewaySecretsResponse) {
	response = &GetCloudNativeAPIGatewaySecretsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TLS证书
func (c *Client) GetCloudNativeAPIGatewaySecrets(request *GetCloudNativeAPIGatewaySecretsRequest) (response *GetCloudNativeAPIGatewaySecretsResponse, err error) {
	if request == nil {
		request = NewGetCloudNativeAPIGatewaySecretsRequest()
	}
	response = NewGetCloudNativeAPIGatewaySecretsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGovernanceGroupsRequest() (request *DeleteGovernanceGroupsRequest) {
	request = &DeleteGovernanceGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceGroups")
	return
}

func NewDeleteGovernanceGroupsResponse() (response *DeleteGovernanceGroupsResponse) {
	response = &DeleteGovernanceGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除治理中心的用户组
func (c *Client) DeleteGovernanceGroups(request *DeleteGovernanceGroupsRequest) (response *DeleteGovernanceGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteGovernanceGroupsRequest()
	}
	response = NewDeleteGovernanceGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEngineClbAccessAddressRequest() (request *DescribeEngineClbAccessAddressRequest) {
	request = &DescribeEngineClbAccessAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEngineClbAccessAddress")
	return
}

func NewDescribeEngineClbAccessAddressResponse() (response *DescribeEngineClbAccessAddressResponse) {
	response = &DescribeEngineClbAccessAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询引擎实例的CLB访问地址
func (c *Client) DescribeEngineClbAccessAddress(request *DescribeEngineClbAccessAddressRequest) (response *DescribeEngineClbAccessAddressResponse, err error) {
	if request == nil {
		request = NewDescribeEngineClbAccessAddressRequest()
	}
	response = NewDescribeEngineClbAccessAddressResponse()
	err = c.Send(request, response)
	return
}

func NewGetNginxIngressConfigMapRequest() (request *GetNginxIngressConfigMapRequest) {
	request = &GetNginxIngressConfigMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "GetNginxIngressConfigMap")
	return
}

func NewGetNginxIngressConfigMapResponse() (response *GetNginxIngressConfigMapResponse) {
	response = &GetNginxIngressConfigMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取 Nginx Ingress 全局 ConfigMap 配置
func (c *Client) GetNginxIngressConfigMap(request *GetNginxIngressConfigMapRequest) (response *GetNginxIngressConfigMapResponse, err error) {
	if request == nil {
		request = NewGetNginxIngressConfigMapRequest()
	}
	response = NewGetNginxIngressConfigMapResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateEksListRequest() (request *DescribeOperateEksListRequest) {
	request = &DescribeOperateEksListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateEksList")
	return
}

func NewDescribeOperateEksListResponse() (response *DescribeOperateEksListResponse) {
	response = &DescribeOperateEksListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台EKS列表
func (c *Client) DescribeOperateEksList(request *DescribeOperateEksListRequest) (response *DescribeOperateEksListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateEksListRequest()
	}
	response = NewDescribeOperateEksListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateWhiteListRequest() (request *DescribeOperateWhiteListRequest) {
	request = &DescribeOperateWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateWhiteList")
	return
}

func NewDescribeOperateWhiteListResponse() (response *DescribeOperateWhiteListResponse) {
	response = &DescribeOperateWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台白名单列表
func (c *Client) DescribeOperateWhiteList(request *DescribeOperateWhiteListRequest) (response *DescribeOperateWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateWhiteListRequest()
	}
	response = NewDescribeOperateWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateEngineSpecRequest() (request *ModifyOperateEngineSpecRequest) {
	request = &ModifyOperateEngineSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyOperateEngineSpec")
	return
}

func NewModifyOperateEngineSpecResponse() (response *ModifyOperateEngineSpecResponse) {
	response = &ModifyOperateEngineSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改运营端引擎规格
func (c *Client) ModifyOperateEngineSpec(request *ModifyOperateEngineSpecRequest) (response *ModifyOperateEngineSpecResponse, err error) {
	if request == nil {
		request = NewModifyOperateEngineSpecRequest()
	}
	response = NewModifyOperateEngineSpecResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperateEksRequest() (request *DeleteOperateEksRequest) {
	request = &DeleteOperateEksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteOperateEks")
	return
}

func NewDeleteOperateEksResponse() (response *DeleteOperateEksResponse) {
	response = &DeleteOperateEksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运营平台EKS集群
func (c *Client) DeleteOperateEks(request *DeleteOperateEksRequest) (response *DeleteOperateEksResponse, err error) {
	if request == nil {
		request = NewDeleteOperateEksRequest()
	}
	response = NewDeleteOperateEksResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewayRequest() (request *DeleteCloudNativeAPIGatewayRequest) {
	request = &DeleteCloudNativeAPIGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGateway")
	return
}

func NewDeleteCloudNativeAPIGatewayResponse() (response *DeleteCloudNativeAPIGatewayResponse) {
	response = &DeleteCloudNativeAPIGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除云原生API网关实例
func (c *Client) DeleteCloudNativeAPIGateway(request *DeleteCloudNativeAPIGatewayRequest) (response *DeleteCloudNativeAPIGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewayRequest()
	}
	response = NewDeleteCloudNativeAPIGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewResetGovernanceUserTokenRequest() (request *ResetGovernanceUserTokenRequest) {
	request = &ResetGovernanceUserTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ResetGovernanceUserToken")
	return
}

func NewResetGovernanceUserTokenResponse() (response *ResetGovernanceUserTokenResponse) {
	response = &ResetGovernanceUserTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置治理中心用户Token
func (c *Client) ResetGovernanceUserToken(request *ResetGovernanceUserTokenRequest) (response *ResetGovernanceUserTokenResponse, err error) {
	if request == nil {
		request = NewResetGovernanceUserTokenRequest()
	}
	response = NewResetGovernanceUserTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoScalerResourceStrategiesRequest() (request *DescribeAutoScalerResourceStrategiesRequest) {
	request = &DescribeAutoScalerResourceStrategiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeAutoScalerResourceStrategies")
	return
}

func NewDescribeAutoScalerResourceStrategiesResponse() (response *DescribeAutoScalerResourceStrategiesResponse) {
	response = &DescribeAutoScalerResourceStrategiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看弹性伸缩策略列表
func (c *Client) DescribeAutoScalerResourceStrategies(request *DescribeAutoScalerResourceStrategiesRequest) (response *DescribeAutoScalerResourceStrategiesResponse, err error) {
	if request == nil {
		request = NewDescribeAutoScalerResourceStrategiesRequest()
	}
	response = NewDescribeAutoScalerResourceStrategiesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEngineInternetAccessRequest() (request *UpdateEngineInternetAccessRequest) {
	request = &UpdateEngineInternetAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateEngineInternetAccess")
	return
}

func NewUpdateEngineInternetAccessResponse() (response *UpdateEngineInternetAccessResponse) {
	response = &UpdateEngineInternetAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改引擎公网访问配置
func (c *Client) UpdateEngineInternetAccess(request *UpdateEngineInternetAccessRequest) (response *UpdateEngineInternetAccessResponse, err error) {
	if request == nil {
		request = NewUpdateEngineInternetAccessRequest()
	}
	response = NewUpdateEngineInternetAccessResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigFileReleaseHistoriesRequest() (request *DescribeConfigFileReleaseHistoriesRequest) {
	request = &DescribeConfigFileReleaseHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFileReleaseHistories")
	return
}

func NewDescribeConfigFileReleaseHistoriesResponse() (response *DescribeConfigFileReleaseHistoriesResponse) {
	response = &DescribeConfigFileReleaseHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置文件发布历史列表
func (c *Client) DescribeConfigFileReleaseHistories(request *DescribeConfigFileReleaseHistoriesRequest) (response *DescribeConfigFileReleaseHistoriesResponse, err error) {
	if request == nil {
		request = NewDescribeConfigFileReleaseHistoriesRequest()
	}
	response = NewDescribeConfigFileReleaseHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallCloudNativeAPIGatewayPluginRequest() (request *UninstallCloudNativeAPIGatewayPluginRequest) {
	request = &UninstallCloudNativeAPIGatewayPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UninstallCloudNativeAPIGatewayPlugin")
	return
}

func NewUninstallCloudNativeAPIGatewayPluginResponse() (response *UninstallCloudNativeAPIGatewayPluginResponse) {
	response = &UninstallCloudNativeAPIGatewayPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云原生网关卸载插件
func (c *Client) UninstallCloudNativeAPIGatewayPlugin(request *UninstallCloudNativeAPIGatewayPluginRequest) (response *UninstallCloudNativeAPIGatewayPluginResponse, err error) {
	if request == nil {
		request = NewUninstallCloudNativeAPIGatewayPluginRequest()
	}
	response = NewUninstallCloudNativeAPIGatewayPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperServicesRequest() (request *DescribeZookeeperServicesRequest) {
	request = &DescribeZookeeperServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperServices")
	return
}

func NewDescribeZookeeperServicesResponse() (response *DescribeZookeeperServicesResponse) {
	response = &DescribeZookeeperServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定实例下的zookeeper注册服务列表
func (c *Client) DescribeZookeeperServices(request *DescribeZookeeperServicesRequest) (response *DescribeZookeeperServicesResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperServicesRequest()
	}
	response = NewDescribeZookeeperServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigFileGroupRequest() (request *DeleteConfigFileGroupRequest) {
	request = &DeleteConfigFileGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteConfigFileGroup")
	return
}

func NewDeleteConfigFileGroupResponse() (response *DeleteConfigFileGroupResponse) {
	response = &DeleteConfigFileGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件批量删除配置文件组列表
func (c *Client) DeleteConfigFileGroup(request *DeleteConfigFileGroupRequest) (response *DeleteConfigFileGroupResponse, err error) {
	if request == nil {
		request = NewDeleteConfigFileGroupRequest()
	}
	response = NewDeleteConfigFileGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPolarisCLSTopicRequest() (request *ModifyPolarisCLSTopicRequest) {
	request = &ModifyPolarisCLSTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyPolarisCLSTopic")
	return
}

func NewModifyPolarisCLSTopicResponse() (response *ModifyPolarisCLSTopicResponse) {
	response = &ModifyPolarisCLSTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑Polaris的CLS日志topic列表
func (c *Client) ModifyPolarisCLSTopic(request *ModifyPolarisCLSTopicRequest) (response *ModifyPolarisCLSTopicResponse, err error) {
	if request == nil {
		request = NewModifyPolarisCLSTopicRequest()
	}
	response = NewModifyPolarisCLSTopicResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewayCertificatesDomainRequest() (request *ListCloudNativeAPIGatewayCertificatesDomainRequest) {
	request = &ListCloudNativeAPIGatewayCertificatesDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGatewayCertificatesDomain")
	return
}

func NewListCloudNativeAPIGatewayCertificatesDomainResponse() (response *ListCloudNativeAPIGatewayCertificatesDomainResponse) {
	response = &ListCloudNativeAPIGatewayCertificatesDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关证书域名列表
func (c *Client) ListCloudNativeAPIGatewayCertificatesDomain(request *ListCloudNativeAPIGatewayCertificatesDomainRequest) (response *ListCloudNativeAPIGatewayCertificatesDomainResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewayCertificatesDomainRequest()
	}
	response = NewListCloudNativeAPIGatewayCertificatesDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePolarisReplicasRequest() (request *DescribePolarisReplicasRequest) {
	request = &DescribePolarisReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribePolarisReplicas")
	return
}

func NewDescribePolarisReplicasResponse() (response *DescribePolarisReplicasResponse) {
	response = &DescribePolarisReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Polaris引擎实例副本信息
func (c *Client) DescribePolarisReplicas(request *DescribePolarisReplicasRequest) (response *DescribePolarisReplicasResponse, err error) {
	if request == nil {
		request = NewDescribePolarisReplicasRequest()
	}
	response = NewDescribePolarisReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudNativeAPIGatewayCanaryRuleRequest() (request *ModifyCloudNativeAPIGatewayCanaryRuleRequest) {
	request = &ModifyCloudNativeAPIGatewayCanaryRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayCanaryRule")
	return
}

func NewModifyCloudNativeAPIGatewayCanaryRuleResponse() (response *ModifyCloudNativeAPIGatewayCanaryRuleResponse) {
	response = &ModifyCloudNativeAPIGatewayCanaryRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生网关的灰度规则
func (c *Client) ModifyCloudNativeAPIGatewayCanaryRule(request *ModifyCloudNativeAPIGatewayCanaryRuleRequest) (response *ModifyCloudNativeAPIGatewayCanaryRuleResponse, err error) {
	if request == nil {
		request = NewModifyCloudNativeAPIGatewayCanaryRuleRequest()
	}
	response = NewModifyCloudNativeAPIGatewayCanaryRuleResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewaySystemParametersRequest() (request *ListCloudNativeAPIGatewaySystemParametersRequest) {
	request = &ListCloudNativeAPIGatewaySystemParametersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGatewaySystemParameters")
	return
}

func NewListCloudNativeAPIGatewaySystemParametersResponse() (response *ListCloudNativeAPIGatewaySystemParametersResponse) {
	response = &ListCloudNativeAPIGatewaySystemParametersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关系统参数列表
func (c *Client) ListCloudNativeAPIGatewaySystemParameters(request *ListCloudNativeAPIGatewaySystemParametersRequest) (response *ListCloudNativeAPIGatewaySystemParametersResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewaySystemParametersRequest()
	}
	response = NewListCloudNativeAPIGatewaySystemParametersResponse()
	err = c.Send(request, response)
	return
}

func NewBreakZookeeperMigrateClusterRequest() (request *BreakZookeeperMigrateClusterRequest) {
	request = &BreakZookeeperMigrateClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "BreakZookeeperMigrateCluster")
	return
}

func NewBreakZookeeperMigrateClusterResponse() (response *BreakZookeeperMigrateClusterResponse) {
	response = &BreakZookeeperMigrateClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分离集群
func (c *Client) BreakZookeeperMigrateCluster(request *BreakZookeeperMigrateClusterRequest) (response *BreakZookeeperMigrateClusterResponse, err error) {
	if request == nil {
		request = NewBreakZookeeperMigrateClusterRequest()
	}
	response = NewBreakZookeeperMigrateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceInstancesRequest() (request *DescribeGovernanceInstancesRequest) {
	request = &DescribeGovernanceInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceInstances")
	return
}

func NewDescribeGovernanceInstancesResponse() (response *DescribeGovernanceInstancesResponse) {
	response = &DescribeGovernanceInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心服务实例
func (c *Client) DescribeGovernanceInstances(request *DescribeGovernanceInstancesRequest) (response *DescribeGovernanceInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceInstancesRequest()
	}
	response = NewDescribeGovernanceInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEngineSyncJobCheckResultRequest() (request *DescribeEngineSyncJobCheckResultRequest) {
	request = &DescribeEngineSyncJobCheckResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEngineSyncJobCheckResult")
	return
}

func NewDescribeEngineSyncJobCheckResultResponse() (response *DescribeEngineSyncJobCheckResultResponse) {
	response = &DescribeEngineSyncJobCheckResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询引擎数据同步任务校验结果
func (c *Client) DescribeEngineSyncJobCheckResult(request *DescribeEngineSyncJobCheckResultRequest) (response *DescribeEngineSyncJobCheckResultResponse, err error) {
	if request == nil {
		request = NewDescribeEngineSyncJobCheckResultRequest()
	}
	response = NewDescribeEngineSyncJobCheckResultResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNacosAdminPasswordRequest() (request *ModifyNacosAdminPasswordRequest) {
	request = &ModifyNacosAdminPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyNacosAdminPassword")
	return
}

func NewModifyNacosAdminPasswordResponse() (response *ModifyNacosAdminPasswordResponse) {
	response = &ModifyNacosAdminPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置nacos管理员密码
func (c *Client) ModifyNacosAdminPassword(request *ModifyNacosAdminPasswordRequest) (response *ModifyNacosAdminPasswordResponse, err error) {
	if request == nil {
		request = NewModifyNacosAdminPasswordRequest()
	}
	response = NewModifyNacosAdminPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGovernanceNamespacesRequest() (request *CreateGovernanceNamespacesRequest) {
	request = &CreateGovernanceNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceNamespaces")
	return
}

func NewCreateGovernanceNamespacesResponse() (response *CreateGovernanceNamespacesResponse) {
	response = &CreateGovernanceNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建治理中心命名空间
func (c *Client) CreateGovernanceNamespaces(request *CreateGovernanceNamespacesRequest) (response *CreateGovernanceNamespacesResponse, err error) {
	if request == nil {
		request = NewCreateGovernanceNamespacesRequest()
	}
	response = NewCreateGovernanceNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigFilesByGroupRequest() (request *DescribeConfigFilesByGroupRequest) {
	request = &DescribeConfigFilesByGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFilesByGroup")
	return
}

func NewDescribeConfigFilesByGroupResponse() (response *DescribeConfigFilesByGroupResponse) {
	response = &DescribeConfigFilesByGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据group查询配置文件列表
func (c *Client) DescribeConfigFilesByGroup(request *DescribeConfigFilesByGroupRequest) (response *DescribeConfigFilesByGroupResponse, err error) {
	if request == nil {
		request = NewDescribeConfigFilesByGroupRequest()
	}
	response = NewDescribeConfigFilesByGroupResponse()
	err = c.Send(request, response)
	return
}

func NewInstallCloudNativeAPIGatewayPluginRequest() (request *InstallCloudNativeAPIGatewayPluginRequest) {
	request = &InstallCloudNativeAPIGatewayPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "InstallCloudNativeAPIGatewayPlugin")
	return
}

func NewInstallCloudNativeAPIGatewayPluginResponse() (response *InstallCloudNativeAPIGatewayPluginResponse) {
	response = &InstallCloudNativeAPIGatewayPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安装云原生API网关插件
func (c *Client) InstallCloudNativeAPIGatewayPlugin(request *InstallCloudNativeAPIGatewayPluginRequest) (response *InstallCloudNativeAPIGatewayPluginResponse, err error) {
	if request == nil {
		request = NewInstallCloudNativeAPIGatewayPluginRequest()
	}
	response = NewInstallCloudNativeAPIGatewayPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewaySecretRequest() (request *DeleteCloudNativeAPIGatewaySecretRequest) {
	request = &DeleteCloudNativeAPIGatewaySecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewaySecret")
	return
}

func NewDeleteCloudNativeAPIGatewaySecretResponse() (response *DeleteCloudNativeAPIGatewaySecretResponse) {
	response = &DeleteCloudNativeAPIGatewaySecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除TLS证书
func (c *Client) DeleteCloudNativeAPIGatewaySecret(request *DeleteCloudNativeAPIGatewaySecretRequest) (response *DeleteCloudNativeAPIGatewaySecretResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewaySecretRequest()
	}
	response = NewDeleteCloudNativeAPIGatewaySecretResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudNativeAPIGatewayCertificateDetailsRequest() (request *ListCloudNativeAPIGatewayCertificateDetailsRequest) {
	request = &ListCloudNativeAPIGatewayCertificateDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ListCloudNativeAPIGatewayCertificateDetails")
	return
}

func NewListCloudNativeAPIGatewayCertificateDetailsResponse() (response *ListCloudNativeAPIGatewayCertificateDetailsResponse) {
	response = &ListCloudNativeAPIGatewayCertificateDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关单个证书详情
func (c *Client) ListCloudNativeAPIGatewayCertificateDetails(request *ListCloudNativeAPIGatewayCertificateDetailsRequest) (response *ListCloudNativeAPIGatewayCertificateDetailsResponse, err error) {
	if request == nil {
		request = NewListCloudNativeAPIGatewayCertificateDetailsRequest()
	}
	response = NewListCloudNativeAPIGatewayCertificateDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNacosDefaultConfigInfoRequest() (request *DescribeNacosDefaultConfigInfoRequest) {
	request = &DescribeNacosDefaultConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeNacosDefaultConfigInfo")
	return
}

func NewDescribeNacosDefaultConfigInfoResponse() (response *DescribeNacosDefaultConfigInfoResponse) {
	response = &DescribeNacosDefaultConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看Nacos默认系统参数
func (c *Client) DescribeNacosDefaultConfigInfo(request *DescribeNacosDefaultConfigInfoRequest) (response *DescribeNacosDefaultConfigInfoResponse, err error) {
	if request == nil {
		request = NewDescribeNacosDefaultConfigInfoRequest()
	}
	response = NewDescribeNacosDefaultConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewInstallCloudNativeAPIGatewaySystemPluginRequest() (request *InstallCloudNativeAPIGatewaySystemPluginRequest) {
	request = &InstallCloudNativeAPIGatewaySystemPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "InstallCloudNativeAPIGatewaySystemPlugin")
	return
}

func NewInstallCloudNativeAPIGatewaySystemPluginResponse() (response *InstallCloudNativeAPIGatewaySystemPluginResponse) {
	response = &InstallCloudNativeAPIGatewaySystemPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安装云原生API网关系统插件
func (c *Client) InstallCloudNativeAPIGatewaySystemPlugin(request *InstallCloudNativeAPIGatewaySystemPluginRequest) (response *InstallCloudNativeAPIGatewaySystemPluginResponse, err error) {
	if request == nil {
		request = NewInstallCloudNativeAPIGatewaySystemPluginRequest()
	}
	response = NewInstallCloudNativeAPIGatewaySystemPluginResponse()
	err = c.Send(request, response)
	return
}

func NewChangeNacosAuthStatusRequest() (request *ChangeNacosAuthStatusRequest) {
	request = &ChangeNacosAuthStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ChangeNacosAuthStatus")
	return
}

func NewChangeNacosAuthStatusResponse() (response *ChangeNacosAuthStatusResponse) {
	response = &ChangeNacosAuthStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更nacos接口和sdk鉴权开关状态
func (c *Client) ChangeNacosAuthStatus(request *ChangeNacosAuthStatusRequest) (response *ChangeNacosAuthStatusResponse, err error) {
	if request == nil {
		request = NewChangeNacosAuthStatusRequest()
	}
	response = NewChangeNacosAuthStatusResponse()
	err = c.Send(request, response)
	return
}

func NewObtainApolloAdminTokenRequest() (request *ObtainApolloAdminTokenRequest) {
	request = &ObtainApolloAdminTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ObtainApolloAdminToken")
	return
}

func NewObtainApolloAdminTokenResponse() (response *ObtainApolloAdminTokenResponse) {
	response = &ObtainApolloAdminTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取apollo管理员token
func (c *Client) ObtainApolloAdminToken(request *ObtainApolloAdminTokenRequest) (response *ObtainApolloAdminTokenResponse, err error) {
	if request == nil {
		request = NewObtainApolloAdminTokenRequest()
	}
	response = NewObtainApolloAdminTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigFilesRequest() (request *DescribeConfigFilesRequest) {
	request = &DescribeConfigFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConfigFiles")
	return
}

func NewDescribeConfigFilesResponse() (response *DescribeConfigFilesResponse) {
	response = &DescribeConfigFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据命名空间、组名、名称、标签查询配置文件列表
func (c *Client) DescribeConfigFiles(request *DescribeConfigFilesRequest) (response *DescribeConfigFilesResponse, err error) {
	if request == nil {
		request = NewDescribeConfigFilesRequest()
	}
	response = NewDescribeConfigFilesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceNamespacesRequest() (request *DescribeGovernanceNamespacesRequest) {
	request = &DescribeGovernanceNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceNamespaces")
	return
}

func NewDescribeGovernanceNamespacesResponse() (response *DescribeGovernanceNamespacesResponse) {
	response = &DescribeGovernanceNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务治理中心命名空间列表
func (c *Client) DescribeGovernanceNamespaces(request *DescribeGovernanceNamespacesRequest) (response *DescribeGovernanceNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceNamespacesRequest()
	}
	response = NewDescribeGovernanceNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceGroupsRequest() (request *DescribeGovernanceGroupsRequest) {
	request = &DescribeGovernanceGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceGroups")
	return
}

func NewDescribeGovernanceGroupsResponse() (response *DescribeGovernanceGroupsResponse) {
	response = &DescribeGovernanceGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心用户组列表
func (c *Client) DescribeGovernanceGroups(request *DescribeGovernanceGroupsRequest) (response *DescribeGovernanceGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceGroupsRequest()
	}
	response = NewDescribeGovernanceGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsulServiceInstancesRequest() (request *DescribeConsulServiceInstancesRequest) {
	request = &DescribeConsulServiceInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConsulServiceInstances")
	return
}

func NewDescribeConsulServiceInstancesResponse() (response *DescribeConsulServiceInstancesResponse) {
	response = &DescribeConsulServiceInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定实例下查询consul注册服务的实例列表
func (c *Client) DescribeConsulServiceInstances(request *DescribeConsulServiceInstancesRequest) (response *DescribeConsulServiceInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeConsulServiceInstancesRequest()
	}
	response = NewDescribeConsulServiceInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEurekaConfigInfoRequest() (request *DescribeEurekaConfigInfoRequest) {
	request = &DescribeEurekaConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEurekaConfigInfo")
	return
}

func NewDescribeEurekaConfigInfoResponse() (response *DescribeEurekaConfigInfoResponse) {
	response = &DescribeEurekaConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Eureka参数配置
func (c *Client) DescribeEurekaConfigInfo(request *DescribeEurekaConfigInfoRequest) (response *DescribeEurekaConfigInfoResponse, err error) {
	if request == nil {
		request = NewDescribeEurekaConfigInfoRequest()
	}
	response = NewDescribeEurekaConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNativeGatewayServiceGroupRequest() (request *DescribeNativeGatewayServiceGroupRequest) {
	request = &DescribeNativeGatewayServiceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeNativeGatewayServiceGroup")
	return
}

func NewDescribeNativeGatewayServiceGroupResponse() (response *DescribeNativeGatewayServiceGroupResponse) {
	response = &DescribeNativeGatewayServiceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关服务实例分组列表
func (c *Client) DescribeNativeGatewayServiceGroup(request *DescribeNativeGatewayServiceGroupRequest) (response *DescribeNativeGatewayServiceGroupResponse, err error) {
	if request == nil {
		request = NewDescribeNativeGatewayServiceGroupRequest()
	}
	response = NewDescribeNativeGatewayServiceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudNativeAPIGatewayIngressRequest() (request *DeleteCloudNativeAPIGatewayIngressRequest) {
	request = &DeleteCloudNativeAPIGatewayIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayIngress")
	return
}

func NewDeleteCloudNativeAPIGatewayIngressResponse() (response *DeleteCloudNativeAPIGatewayIngressResponse) {
	response = &DeleteCloudNativeAPIGatewayIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Ingress
func (c *Client) DeleteCloudNativeAPIGatewayIngress(request *DeleteCloudNativeAPIGatewayIngressRequest) (response *DeleteCloudNativeAPIGatewayIngressResponse, err error) {
	if request == nil {
		request = NewDeleteCloudNativeAPIGatewayIngressRequest()
	}
	response = NewDeleteCloudNativeAPIGatewayIngressResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGovernanceInstancesRequest() (request *DeleteGovernanceInstancesRequest) {
	request = &DeleteGovernanceInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceInstances")
	return
}

func NewDeleteGovernanceInstancesResponse() (response *DeleteGovernanceInstancesResponse) {
	response = &DeleteGovernanceInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除治理中心服务实例
func (c *Client) DeleteGovernanceInstances(request *DeleteGovernanceInstancesRequest) (response *DeleteGovernanceInstancesResponse, err error) {
	if request == nil {
		request = NewDeleteGovernanceInstancesRequest()
	}
	response = NewDeleteGovernanceInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayUpstreamRequest() (request *DescribeCloudNativeAPIGatewayUpstreamRequest) {
	request = &DescribeCloudNativeAPIGatewayUpstreamRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayUpstream")
	return
}

func NewDescribeCloudNativeAPIGatewayUpstreamResponse() (response *DescribeCloudNativeAPIGatewayUpstreamResponse) {
	response = &DescribeCloudNativeAPIGatewayUpstreamResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关服务详情下的Upstream列表
func (c *Client) DescribeCloudNativeAPIGatewayUpstream(request *DescribeCloudNativeAPIGatewayUpstreamRequest) (response *DescribeCloudNativeAPIGatewayUpstreamResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayUpstreamRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayUpstreamResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceAliasRequest() (request *ModifyGovernanceAliasRequest) {
	request = &ModifyGovernanceAliasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceAlias")
	return
}

func NewModifyGovernanceAliasResponse() (response *ModifyGovernanceAliasResponse) {
	response = &ModifyGovernanceAliasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改治理中心服务别名
func (c *Client) ModifyGovernanceAlias(request *ModifyGovernanceAliasRequest) (response *ModifyGovernanceAliasResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceAliasRequest()
	}
	response = NewModifyGovernanceAliasResponse()
	err = c.Send(request, response)
	return
}

func NewPublishConfigFilesRequest() (request *PublishConfigFilesRequest) {
	request = &PublishConfigFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "PublishConfigFiles")
	return
}

func NewPublishConfigFilesResponse() (response *PublishConfigFilesResponse) {
	response = &PublishConfigFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布配置文件
func (c *Client) PublishConfigFiles(request *PublishConfigFilesRequest) (response *PublishConfigFilesResponse, err error) {
	if request == nil {
		request = NewPublishConfigFilesRequest()
	}
	response = NewPublishConfigFilesResponse()
	err = c.Send(request, response)
	return
}

func NewCloseKongIngressRequest() (request *CloseKongIngressRequest) {
	request = &CloseKongIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CloseKongIngress")
	return
}

func NewCloseKongIngressResponse() (response *CloseKongIngressResponse) {
	response = &CloseKongIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭KongIngress
func (c *Client) CloseKongIngress(request *CloseKongIngressRequest) (response *CloseKongIngressResponse, err error) {
	if request == nil {
		request = NewCloseKongIngressRequest()
	}
	response = NewCloseKongIngressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateAgentListRequest() (request *DescribeOperateAgentListRequest) {
	request = &DescribeOperateAgentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateAgentList")
	return
}

func NewDescribeOperateAgentListResponse() (response *DescribeOperateAgentListResponse) {
	response = &DescribeOperateAgentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台Agent列表
func (c *Client) DescribeOperateAgentList(request *DescribeOperateAgentListRequest) (response *DescribeOperateAgentListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateAgentListRequest()
	}
	response = NewDescribeOperateAgentListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigFileRequest() (request *CreateConfigFileRequest) {
	request = &CreateConfigFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateConfigFile")
	return
}

func NewCreateConfigFileResponse() (response *CreateConfigFileResponse) {
	response = &CreateConfigFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建配置文件
func (c *Client) CreateConfigFile(request *CreateConfigFileRequest) (response *CreateConfigFileResponse, err error) {
	if request == nil {
		request = NewCreateConfigFileRequest()
	}
	response = NewCreateConfigFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewaysRequest() (request *DescribeCloudNativeAPIGatewaysRequest) {
	request = &DescribeCloudNativeAPIGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGateways")
	return
}

func NewDescribeCloudNativeAPIGatewaysResponse() (response *DescribeCloudNativeAPIGatewaysResponse) {
	response = &DescribeCloudNativeAPIGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生API网关实例列表
func (c *Client) DescribeCloudNativeAPIGateways(request *DescribeCloudNativeAPIGatewaysRequest) (response *DescribeCloudNativeAPIGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewaysRequest()
	}
	response = NewDescribeCloudNativeAPIGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGovernanceStrategyRequest() (request *CreateGovernanceStrategyRequest) {
	request = &CreateGovernanceStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceStrategy")
	return
}

func NewCreateGovernanceStrategyResponse() (response *CreateGovernanceStrategyResponse) {
	response = &CreateGovernanceStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建治理中心鉴权策略
func (c *Client) CreateGovernanceStrategy(request *CreateGovernanceStrategyRequest) (response *CreateGovernanceStrategyResponse, err error) {
	if request == nil {
		request = NewCreateGovernanceStrategyRequest()
	}
	response = NewCreateGovernanceStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperateRegionRequest() (request *DeleteOperateRegionRequest) {
	request = &DeleteOperateRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteOperateRegion")
	return
}

func NewDeleteOperateRegionResponse() (response *DeleteOperateRegionResponse) {
	response = &DeleteOperateRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运营平台地域
func (c *Client) DeleteOperateRegion(request *DeleteOperateRegionRequest) (response *DeleteOperateRegionResponse, err error) {
	if request == nil {
		request = NewDeleteOperateRegionRequest()
	}
	response = NewDeleteOperateRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceAliasesRequest() (request *DescribeGovernanceAliasesRequest) {
	request = &DescribeGovernanceAliasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceAliases")
	return
}

func NewDescribeGovernanceAliasesResponse() (response *DescribeGovernanceAliasesResponse) {
	response = &DescribeGovernanceAliasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心服务别名列表
func (c *Client) DescribeGovernanceAliases(request *DescribeGovernanceAliasesRequest) (response *DescribeGovernanceAliasesResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceAliasesRequest()
	}
	response = NewDescribeGovernanceAliasesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateInstancesRequest() (request *DescribeOperateInstancesRequest) {
	request = &DescribeOperateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateInstances")
	return
}

func NewDescribeOperateInstancesResponse() (response *DescribeOperateInstancesResponse) {
	response = &DescribeOperateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营平台引擎实例列表
func (c *Client) DescribeOperateInstances(request *DescribeOperateInstancesRequest) (response *DescribeOperateInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeOperateInstancesRequest()
	}
	response = NewDescribeOperateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewClosePrometheusRequest() (request *ClosePrometheusRequest) {
	request = &ClosePrometheusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ClosePrometheus")
	return
}

func NewClosePrometheusResponse() (response *ClosePrometheusResponse) {
	response = &ClosePrometheusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭Prometheus指标推送
func (c *Client) ClosePrometheus(request *ClosePrometheusRequest) (response *ClosePrometheusResponse, err error) {
	if request == nil {
		request = NewClosePrometheusRequest()
	}
	response = NewClosePrometheusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEngineQuotasRequest() (request *DescribeEngineQuotasRequest) {
	request = &DescribeEngineQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeEngineQuotas")
	return
}

func NewDescribeEngineQuotasResponse() (response *DescribeEngineQuotasResponse) {
	response = &DescribeEngineQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取引擎剩余额度
func (c *Client) DescribeEngineQuotas(request *DescribeEngineQuotasRequest) (response *DescribeEngineQuotasResponse, err error) {
	if request == nil {
		request = NewDescribeEngineQuotasRequest()
	}
	response = NewDescribeEngineQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewRunExecTaskBatchJobRequest() (request *RunExecTaskBatchJobRequest) {
	request = &RunExecTaskBatchJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "RunExecTaskBatchJob")
	return
}

func NewRunExecTaskBatchJobResponse() (response *RunExecTaskBatchJobResponse) {
	response = &RunExecTaskBatchJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行运营平台可执行任务的批次任务
func (c *Client) RunExecTaskBatchJob(request *RunExecTaskBatchJobRequest) (response *RunExecTaskBatchJobResponse, err error) {
	if request == nil {
		request = NewRunExecTaskBatchJobRequest()
	}
	response = NewRunExecTaskBatchJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGovernanceServicesRequest() (request *CreateGovernanceServicesRequest) {
	request = &CreateGovernanceServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceServices")
	return
}

func NewCreateGovernanceServicesResponse() (response *CreateGovernanceServicesResponse) {
	response = &CreateGovernanceServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建治理中心服务
func (c *Client) CreateGovernanceServices(request *CreateGovernanceServicesRequest) (response *CreateGovernanceServicesResponse, err error) {
	if request == nil {
		request = NewCreateGovernanceServicesRequest()
	}
	response = NewCreateGovernanceServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGovernanceUsersRequest() (request *DeleteGovernanceUsersRequest) {
	request = &DeleteGovernanceUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceUsers")
	return
}

func NewDeleteGovernanceUsersResponse() (response *DeleteGovernanceUsersResponse) {
	response = &DeleteGovernanceUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除治理中心的用户
func (c *Client) DeleteGovernanceUsers(request *DeleteGovernanceUsersRequest) (response *DeleteGovernanceUsersResponse, err error) {
	if request == nil {
		request = NewDeleteGovernanceUsersRequest()
	}
	response = NewDeleteGovernanceUsersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperateEngineSpecRequest() (request *DeleteOperateEngineSpecRequest) {
	request = &DeleteOperateEngineSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteOperateEngineSpec")
	return
}

func NewDeleteOperateEngineSpecResponse() (response *DeleteOperateEngineSpecResponse) {
	response = &DeleteOperateEngineSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除引擎规格
func (c *Client) DeleteOperateEngineSpec(request *DeleteOperateEngineSpecRequest) (response *DeleteOperateEngineSpecResponse, err error) {
	if request == nil {
		request = NewDeleteOperateEngineSpecRequest()
	}
	response = NewDeleteOperateEngineSpecResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperateEngineVersionRequest() (request *DeleteOperateEngineVersionRequest) {
	request = &DeleteOperateEngineVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteOperateEngineVersion")
	return
}

func NewDeleteOperateEngineVersionResponse() (response *DeleteOperateEngineVersionResponse) {
	response = &DeleteOperateEngineVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运营端引擎版本
func (c *Client) DeleteOperateEngineVersion(request *DeleteOperateEngineVersionRequest) (response *DeleteOperateEngineVersionResponse, err error) {
	if request == nil {
		request = NewDeleteOperateEngineVersionRequest()
	}
	response = NewDeleteOperateEngineVersionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperateEksRequest() (request *ModifyOperateEksRequest) {
	request = &ModifyOperateEksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyOperateEks")
	return
}

func NewModifyOperateEksResponse() (response *ModifyOperateEksResponse) {
	response = &ModifyOperateEksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新运营平台EKS集群
func (c *Client) ModifyOperateEks(request *ModifyOperateEksRequest) (response *ModifyOperateEksResponse, err error) {
	if request == nil {
		request = NewModifyOperateEksRequest()
	}
	response = NewModifyOperateEksResponse()
	err = c.Send(request, response)
	return
}

func NewRefreshOperateAgentRequest() (request *RefreshOperateAgentRequest) {
	request = &RefreshOperateAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "RefreshOperateAgent")
	return
}

func NewRefreshOperateAgentResponse() (response *RefreshOperateAgentResponse) {
	response = &RefreshOperateAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营平台-刷新全部agent版本记录
func (c *Client) RefreshOperateAgent(request *RefreshOperateAgentRequest) (response *RefreshOperateAgentResponse, err error) {
	if request == nil {
		request = NewRefreshOperateAgentRequest()
	}
	response = NewRefreshOperateAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEngineSyncJobRequest() (request *DeleteEngineSyncJobRequest) {
	request = &DeleteEngineSyncJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DeleteEngineSyncJob")
	return
}

func NewDeleteEngineSyncJobResponse() (response *DeleteEngineSyncJobResponse) {
	response = &DeleteEngineSyncJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除引擎数据同步任务
func (c *Client) DeleteEngineSyncJob(request *DeleteEngineSyncJobRequest) (response *DeleteEngineSyncJobResponse, err error) {
	if request == nil {
		request = NewDeleteEngineSyncJobRequest()
	}
	response = NewDeleteEngineSyncJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceDiscoverEventsRequest() (request *DescribeGovernanceDiscoverEventsRequest) {
	request = &DescribeGovernanceDiscoverEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceDiscoverEvents")
	return
}

func NewDescribeGovernanceDiscoverEventsResponse() (response *DescribeGovernanceDiscoverEventsResponse) {
	response = &DescribeGovernanceDiscoverEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心服务变更历史事件列表
func (c *Client) DescribeGovernanceDiscoverEvents(request *DescribeGovernanceDiscoverEventsRequest) (response *DescribeGovernanceDiscoverEventsResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceDiscoverEventsRequest()
	}
	response = NewDescribeGovernanceDiscoverEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExecTaskBatchJobExeInfoRequest() (request *DescribeExecTaskBatchJobExeInfoRequest) {
	request = &DescribeExecTaskBatchJobExeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeExecTaskBatchJobExeInfo")
	return
}

func NewDescribeExecTaskBatchJobExeInfoResponse() (response *DescribeExecTaskBatchJobExeInfoResponse) {
	response = &DescribeExecTaskBatchJobExeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看运营平台可执行任务的批次任务的描述
func (c *Client) DescribeExecTaskBatchJobExeInfo(request *DescribeExecTaskBatchJobExeInfoRequest) (response *DescribeExecTaskBatchJobExeInfoResponse, err error) {
	if request == nil {
		request = NewDescribeExecTaskBatchJobExeInfoRequest()
	}
	response = NewDescribeExecTaskBatchJobExeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOneCloudNativeAPIGatewayServiceRequest() (request *DescribeOneCloudNativeAPIGatewayServiceRequest) {
	request = &DescribeOneCloudNativeAPIGatewayServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOneCloudNativeAPIGatewayService")
	return
}

func NewDescribeOneCloudNativeAPIGatewayServiceResponse() (response *DescribeOneCloudNativeAPIGatewayServiceResponse) {
	response = &DescribeOneCloudNativeAPIGatewayServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关服务详情
func (c *Client) DescribeOneCloudNativeAPIGatewayService(request *DescribeOneCloudNativeAPIGatewayServiceRequest) (response *DescribeOneCloudNativeAPIGatewayServiceResponse, err error) {
	if request == nil {
		request = NewDescribeOneCloudNativeAPIGatewayServiceRequest()
	}
	response = NewDescribeOneCloudNativeAPIGatewayServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperReplicasRequest() (request *DescribeZookeeperReplicasRequest) {
	request = &DescribeZookeeperReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperReplicas")
	return
}

func NewDescribeZookeeperReplicasResponse() (response *DescribeZookeeperReplicasResponse) {
	response = &DescribeZookeeperReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Zookeeper类型注册引擎实例副本信息
func (c *Client) DescribeZookeeperReplicas(request *DescribeZookeeperReplicasRequest) (response *DescribeZookeeperReplicasResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperReplicasRequest()
	}
	response = NewDescribeZookeeperReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewSyncOperateAgentRequest() (request *SyncOperateAgentRequest) {
	request = &SyncOperateAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "SyncOperateAgent")
	return
}

func NewSyncOperateAgentResponse() (response *SyncOperateAgentResponse) {
	response = &SyncOperateAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营平台-同步agent版本记录
func (c *Client) SyncOperateAgent(request *SyncOperateAgentRequest) (response *SyncOperateAgentResponse, err error) {
	if request == nil {
		request = NewSyncOperateAgentRequest()
	}
	response = NewSyncOperateAgentResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSREInstanceLimitRequest() (request *CheckSREInstanceLimitRequest) {
	request = &CheckSREInstanceLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CheckSREInstanceLimit")
	return
}

func NewCheckSREInstanceLimitResponse() (response *CheckSREInstanceLimitResponse) {
	response = &CheckSREInstanceLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查注册引擎实例数量，判断是否可以创建更多的实例。
func (c *Client) CheckSREInstanceLimit(request *CheckSREInstanceLimitRequest) (response *CheckSREInstanceLimitResponse, err error) {
	if request == nil {
		request = NewCheckSREInstanceLimitRequest()
	}
	response = NewCheckSREInstanceLimitResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCloudNativeAPIGatewayPublicNetworkRequest() (request *CreateCloudNativeAPIGatewayPublicNetworkRequest) {
	request = &CreateCloudNativeAPIGatewayPublicNetworkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayPublicNetwork")
	return
}

func NewCreateCloudNativeAPIGatewayPublicNetworkResponse() (response *CreateCloudNativeAPIGatewayPublicNetworkResponse) {
	response = &CreateCloudNativeAPIGatewayPublicNetworkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建公网网络配置
func (c *Client) CreateCloudNativeAPIGatewayPublicNetwork(request *CreateCloudNativeAPIGatewayPublicNetworkRequest) (response *CreateCloudNativeAPIGatewayPublicNetworkResponse, err error) {
	if request == nil {
		request = NewCreateCloudNativeAPIGatewayPublicNetworkRequest()
	}
	response = NewCreateCloudNativeAPIGatewayPublicNetworkResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
	request = &DescribeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConfig")
	return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
	response = &DescribeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看配置项
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
	if request == nil {
		request = NewDescribeConfigRequest()
	}
	response = NewDescribeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceVersionRequest() (request *ModifyInstanceVersionRequest) {
	request = &ModifyInstanceVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyInstanceVersion")
	return
}

func NewModifyInstanceVersionResponse() (response *ModifyInstanceVersionResponse) {
	response = &ModifyInstanceVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新实例版本
func (c *Client) ModifyInstanceVersion(request *ModifyInstanceVersionRequest) (response *ModifyInstanceVersionResponse, err error) {
	if request == nil {
		request = NewModifyInstanceVersionRequest()
	}
	response = NewModifyInstanceVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceSimpleInfosRequest() (request *DescribeInstanceSimpleInfosRequest) {
	request = &DescribeInstanceSimpleInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeInstanceSimpleInfos")
	return
}

func NewDescribeInstanceSimpleInfosResponse() (response *DescribeInstanceSimpleInfosResponse) {
	response = &DescribeInstanceSimpleInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看所有引擎简略实例
func (c *Client) DescribeInstanceSimpleInfos(request *DescribeInstanceSimpleInfosRequest) (response *DescribeInstanceSimpleInfosResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceSimpleInfosRequest()
	}
	response = NewDescribeInstanceSimpleInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsulConfigInfoRequest() (request *DescribeConsulConfigInfoRequest) {
	request = &DescribeConsulConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeConsulConfigInfo")
	return
}

func NewDescribeConsulConfigInfoResponse() (response *DescribeConsulConfigInfoResponse) {
	response = &DescribeConsulConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Consul参数配置
func (c *Client) DescribeConsulConfigInfo(request *DescribeConsulConfigInfoRequest) (response *DescribeConsulConfigInfoResponse, err error) {
	if request == nil {
		request = NewDescribeConsulConfigInfoRequest()
	}
	response = NewDescribeConsulConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindAutoScalerResourceStrategyFromGroupsRequest() (request *UnbindAutoScalerResourceStrategyFromGroupsRequest) {
	request = &UnbindAutoScalerResourceStrategyFromGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UnbindAutoScalerResourceStrategyFromGroups")
	return
}

func NewUnbindAutoScalerResourceStrategyFromGroupsResponse() (response *UnbindAutoScalerResourceStrategyFromGroupsResponse) {
	response = &UnbindAutoScalerResourceStrategyFromGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 弹性伸缩策略批量解绑网关分组
func (c *Client) UnbindAutoScalerResourceStrategyFromGroups(request *UnbindAutoScalerResourceStrategyFromGroupsRequest) (response *UnbindAutoScalerResourceStrategyFromGroupsResponse, err error) {
	if request == nil {
		request = NewUnbindAutoScalerResourceStrategyFromGroupsRequest()
	}
	response = NewUnbindAutoScalerResourceStrategyFromGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInternetRequest() (request *ModifyInternetRequest) {
	request = &ModifyInternetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyInternet")
	return
}

func NewModifyInternetResponse() (response *ModifyInternetResponse) {
	response = &ModifyInternetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改公网带宽
func (c *Client) ModifyInternet(request *ModifyInternetRequest) (response *ModifyInternetResponse, err error) {
	if request == nil {
		request = NewModifyInternetRequest()
	}
	response = NewModifyInternetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEngineRequest() (request *CreateEngineRequest) {
	request = &CreateEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateEngine")
	return
}

func NewCreateEngineResponse() (response *CreateEngineResponse) {
	response = &CreateEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建引擎实例
func (c *Client) CreateEngine(request *CreateEngineRequest) (response *CreateEngineResponse, err error) {
	if request == nil {
		request = NewCreateEngineRequest()
	}
	response = NewCreateEngineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZooKeeperConfigInfoRequest() (request *DescribeZooKeeperConfigInfoRequest) {
	request = &DescribeZooKeeperConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZooKeeperConfigInfo")
	return
}

func NewDescribeZooKeeperConfigInfoResponse() (response *DescribeZooKeeperConfigInfoResponse) {
	response = &DescribeZooKeeperConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看zookeeper系统参数
func (c *Client) DescribeZooKeeperConfigInfo(request *DescribeZooKeeperConfigInfoRequest) (response *DescribeZooKeeperConfigInfoResponse, err error) {
	if request == nil {
		request = NewDescribeZooKeeperConfigInfoRequest()
	}
	response = NewDescribeZooKeeperConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewaySourcesRequest() (request *DescribeCloudNativeAPIGatewaySourcesRequest) {
	request = &DescribeCloudNativeAPIGatewaySourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewaySources")
	return
}

func NewDescribeCloudNativeAPIGatewaySourcesResponse() (response *DescribeCloudNativeAPIGatewaySourcesResponse) {
	response = &DescribeCloudNativeAPIGatewaySourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云原生网关的服务来源列表
func (c *Client) DescribeCloudNativeAPIGatewaySources(request *DescribeCloudNativeAPIGatewaySourcesRequest) (response *DescribeCloudNativeAPIGatewaySourcesResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewaySourcesRequest()
	}
	response = NewDescribeCloudNativeAPIGatewaySourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAuthRequest() (request *DescribeUserAuthRequest) {
	request = &DescribeUserAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeUserAuth")
	return
}

func NewDescribeUserAuthResponse() (response *DescribeUserAuthResponse) {
	response = &DescribeUserAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户的权限
func (c *Client) DescribeUserAuth(request *DescribeUserAuthRequest) (response *DescribeUserAuthResponse, err error) {
	if request == nil {
		request = NewDescribeUserAuthRequest()
	}
	response = NewDescribeUserAuthResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCloudNativeAPIGatewaySourceRequest() (request *CreateCloudNativeAPIGatewaySourceRequest) {
	request = &CreateCloudNativeAPIGatewaySourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewaySource")
	return
}

func NewCreateCloudNativeAPIGatewaySourceResponse() (response *CreateCloudNativeAPIGatewaySourceResponse) {
	response = &CreateCloudNativeAPIGatewaySourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在云原生网关中，新建服务来源
func (c *Client) CreateCloudNativeAPIGatewaySource(request *CreateCloudNativeAPIGatewaySourceRequest) (response *CreateCloudNativeAPIGatewaySourceResponse, err error) {
	if request == nil {
		request = NewCreateCloudNativeAPIGatewaySourceRequest()
	}
	response = NewCreateCloudNativeAPIGatewaySourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayOriginalPluginsRequest() (request *DescribeCloudNativeAPIGatewayOriginalPluginsRequest) {
	request = &DescribeCloudNativeAPIGatewayOriginalPluginsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayOriginalPlugins")
	return
}

func NewDescribeCloudNativeAPIGatewayOriginalPluginsResponse() (response *DescribeCloudNativeAPIGatewayOriginalPluginsResponse) {
	response = &DescribeCloudNativeAPIGatewayOriginalPluginsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云原生网关原生插件列表
func (c *Client) DescribeCloudNativeAPIGatewayOriginalPlugins(request *DescribeCloudNativeAPIGatewayOriginalPluginsRequest) (response *DescribeCloudNativeAPIGatewayOriginalPluginsResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayOriginalPluginsRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayOriginalPluginsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZooKeeperDefaultConfigInfoRequest() (request *DescribeZooKeeperDefaultConfigInfoRequest) {
	request = &DescribeZooKeeperDefaultConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZooKeeperDefaultConfigInfo")
	return
}

func NewDescribeZooKeeperDefaultConfigInfoResponse() (response *DescribeZooKeeperDefaultConfigInfoResponse) {
	response = &DescribeZooKeeperDefaultConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看zookeeper默认系统参数
func (c *Client) DescribeZooKeeperDefaultConfigInfo(request *DescribeZooKeeperDefaultConfigInfoRequest) (response *DescribeZooKeeperDefaultConfigInfoResponse, err error) {
	if request == nil {
		request = NewDescribeZooKeeperDefaultConfigInfoRequest()
	}
	response = NewDescribeZooKeeperDefaultConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePolarisServerInterfacesRequest() (request *DescribePolarisServerInterfacesRequest) {
	request = &DescribePolarisServerInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribePolarisServerInterfaces")
	return
}

func NewDescribePolarisServerInterfacesResponse() (response *DescribePolarisServerInterfacesResponse) {
	response = &DescribePolarisServerInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询北极星服务接口列表
func (c *Client) DescribePolarisServerInterfaces(request *DescribePolarisServerInterfacesRequest) (response *DescribePolarisServerInterfacesResponse, err error) {
	if request == nil {
		request = NewDescribePolarisServerInterfacesRequest()
	}
	response = NewDescribePolarisServerInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGovernanceUserTokenRequest() (request *ModifyGovernanceUserTokenRequest) {
	request = &ModifyGovernanceUserTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceUserToken")
	return
}

func NewModifyGovernanceUserTokenResponse() (response *ModifyGovernanceUserTokenResponse) {
	response = &ModifyGovernanceUserTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新治理中心用户Token
func (c *Client) ModifyGovernanceUserToken(request *ModifyGovernanceUserTokenRequest) (response *ModifyGovernanceUserTokenResponse, err error) {
	if request == nil {
		request = NewModifyGovernanceUserTokenRequest()
	}
	response = NewModifyGovernanceUserTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNacosServerInterfacesRequest() (request *DescribeNacosServerInterfacesRequest) {
	request = &DescribeNacosServerInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeNacosServerInterfaces")
	return
}

func NewDescribeNacosServerInterfacesResponse() (response *DescribeNacosServerInterfacesResponse) {
	response = &DescribeNacosServerInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询nacos服务接口列表
func (c *Client) DescribeNacosServerInterfaces(request *DescribeNacosServerInterfacesRequest) (response *DescribeNacosServerInterfacesResponse, err error) {
	if request == nil {
		request = NewDescribeNacosServerInterfacesRequest()
	}
	response = NewDescribeNacosServerInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperMigratePhaseRequest() (request *DescribeZookeeperMigratePhaseRequest) {
	request = &DescribeZookeeperMigratePhaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperMigratePhase")
	return
}

func NewDescribeZookeeperMigratePhaseResponse() (response *DescribeZookeeperMigratePhaseResponse) {
	response = &DescribeZookeeperMigratePhaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取zookeeper迁移流程信息
func (c *Client) DescribeZookeeperMigratePhase(request *DescribeZookeeperMigratePhaseRequest) (response *DescribeZookeeperMigratePhaseResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperMigratePhaseRequest()
	}
	response = NewDescribeZookeeperMigratePhaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZookeeperServiceListRequest() (request *DescribeZookeeperServiceListRequest) {
	request = &DescribeZookeeperServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeZookeeperServiceList")
	return
}

func NewDescribeZookeeperServiceListResponse() (response *DescribeZookeeperServiceListResponse) {
	response = &DescribeZookeeperServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定实例下的zookeeper注册服务列表
func (c *Client) DescribeZookeeperServiceList(request *DescribeZookeeperServiceListRequest) (response *DescribeZookeeperServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeZookeeperServiceListRequest()
	}
	response = NewDescribeZookeeperServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudNativeAPIGatewayRouteRequest() (request *ModifyCloudNativeAPIGatewayRouteRequest) {
	request = &ModifyCloudNativeAPIGatewayRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayRoute")
	return
}

func NewModifyCloudNativeAPIGatewayRouteResponse() (response *ModifyCloudNativeAPIGatewayRouteResponse) {
	response = &ModifyCloudNativeAPIGatewayRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生网关路由
func (c *Client) ModifyCloudNativeAPIGatewayRoute(request *ModifyCloudNativeAPIGatewayRouteRequest) (response *ModifyCloudNativeAPIGatewayRouteResponse, err error) {
	if request == nil {
		request = NewModifyCloudNativeAPIGatewayRouteRequest()
	}
	response = NewModifyCloudNativeAPIGatewayRouteResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperateApolloInstancesRequest() (request *UpdateOperateApolloInstancesRequest) {
	request = &UpdateOperateApolloInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateOperateApolloInstances")
	return
}

func NewUpdateOperateApolloInstancesResponse() (response *UpdateOperateApolloInstancesResponse) {
	response = &UpdateOperateApolloInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过运营后台更新 Apollo 集群信息，例如镜像地址、副本数等
func (c *Client) UpdateOperateApolloInstances(request *UpdateOperateApolloInstancesRequest) (response *UpdateOperateApolloInstancesResponse, err error) {
	if request == nil {
		request = NewUpdateOperateApolloInstancesRequest()
	}
	response = NewUpdateOperateApolloInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateFeatureVersionListRequest() (request *DescribeOperateFeatureVersionListRequest) {
	request = &DescribeOperateFeatureVersionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeOperateFeatureVersionList")
	return
}

func NewDescribeOperateFeatureVersionListResponse() (response *DescribeOperateFeatureVersionListResponse) {
	response = &DescribeOperateFeatureVersionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端功能版本列表
func (c *Client) DescribeOperateFeatureVersionList(request *DescribeOperateFeatureVersionListRequest) (response *DescribeOperateFeatureVersionListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateFeatureVersionListRequest()
	}
	response = NewDescribeOperateFeatureVersionListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCloudNativeAPIGatewaySpecRequest() (request *UpdateCloudNativeAPIGatewaySpecRequest) {
	request = &UpdateCloudNativeAPIGatewaySpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "UpdateCloudNativeAPIGatewaySpec")
	return
}

func NewUpdateCloudNativeAPIGatewaySpecResponse() (response *UpdateCloudNativeAPIGatewaySpecResponse) {
	response = &UpdateCloudNativeAPIGatewaySpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改云原生API网关实例的节点规格信息，比如节点扩缩容或者升降配
func (c *Client) UpdateCloudNativeAPIGatewaySpec(request *UpdateCloudNativeAPIGatewaySpecRequest) (response *UpdateCloudNativeAPIGatewaySpecResponse, err error) {
	if request == nil {
		request = NewUpdateCloudNativeAPIGatewaySpecRequest()
	}
	response = NewUpdateCloudNativeAPIGatewaySpecResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudNativeAPIGatewayIngressesRequest() (request *DescribeCloudNativeAPIGatewayIngressesRequest) {
	request = &DescribeCloudNativeAPIGatewayIngressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayIngresses")
	return
}

func NewDescribeCloudNativeAPIGatewayIngressesResponse() (response *DescribeCloudNativeAPIGatewayIngressesResponse) {
	response = &DescribeCloudNativeAPIGatewayIngressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Ingress列表
func (c *Client) DescribeCloudNativeAPIGatewayIngresses(request *DescribeCloudNativeAPIGatewayIngressesRequest) (response *DescribeCloudNativeAPIGatewayIngressesResponse, err error) {
	if request == nil {
		request = NewDescribeCloudNativeAPIGatewayIngressesRequest()
	}
	response = NewDescribeCloudNativeAPIGatewayIngressesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGovernanceUsersRequest() (request *DescribeGovernanceUsersRequest) {
	request = &DescribeGovernanceUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceUsers")
	return
}

func NewDescribeGovernanceUsersResponse() (response *DescribeGovernanceUsersResponse) {
	response = &DescribeGovernanceUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询治理中心用户列表
func (c *Client) DescribeGovernanceUsers(request *DescribeGovernanceUsersRequest) (response *DescribeGovernanceUsersResponse, err error) {
	if request == nil {
		request = NewDescribeGovernanceUsersRequest()
	}
	response = NewDescribeGovernanceUsersResponse()
	err = c.Send(request, response)
	return
}
