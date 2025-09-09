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

package v20200622

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-06-22"

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

func NewQueryClusterFormRequest() (request *QueryClusterFormRequest) {
	request = &QueryClusterFormRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "QueryClusterForm")
	return
}

func NewQueryClusterFormResponse() (response *QueryClusterFormResponse) {
	response = &QueryClusterFormResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群部署单
func (c *Client) QueryClusterForm(request *QueryClusterFormRequest) (response *QueryClusterFormResponse, err error) {
	if request == nil {
		request = NewQueryClusterFormRequest()
	}
	response = NewQueryClusterFormResponse()
	err = c.Send(request, response)
	return
}

func NewGetServiceConfRequest() (request *GetServiceConfRequest) {
	request = &GetServiceConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetServiceConf")
	return
}

func NewGetServiceConfResponse() (response *GetServiceConfResponse) {
	response = &GetServiceConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定模块的配置信息
func (c *Client) GetServiceConf(request *GetServiceConfRequest) (response *GetServiceConfResponse, err error) {
	if request == nil {
		request = NewGetServiceConfRequest()
	}
	response = NewGetServiceConfResponse()
	err = c.Send(request, response)
	return
}

func NewKickDiskRequest() (request *KickDiskRequest) {
	request = &KickDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "KickDisk")
	return
}

func NewKickDiskResponse() (response *KickDiskResponse) {
	response = &KickDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 手动踢盘，主要用于模拟线上故障
func (c *Client) KickDisk(request *KickDiskRequest) (response *KickDiskResponse, err error) {
	if request == nil {
		request = NewKickDiskRequest()
	}
	response = NewKickDiskResponse()
	err = c.Send(request, response)
	return
}

func NewGetScalRequest() (request *GetScalRequest) {
	request = &GetScalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetScal")
	return
}

func NewGetScalResponse() (response *GetScalResponse) {
	response = &GetScalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取scal信息
func (c *Client) GetScal(request *GetScalRequest) (response *GetScalResponse, err error) {
	if request == nil {
		request = NewGetScalRequest()
	}
	response = NewGetScalResponse()
	err = c.Send(request, response)
	return
}

func NewGetSpaceTemplateByIdRequest() (request *GetSpaceTemplateByIdRequest) {
	request = &GetSpaceTemplateByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetSpaceTemplateById")
	return
}

func NewGetSpaceTemplateByIdResponse() (response *GetSpaceTemplateByIdResponse) {
	response = &GetSpaceTemplateByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Id获取Space模版
func (c *Client) GetSpaceTemplateById(request *GetSpaceTemplateByIdRequest) (response *GetSpaceTemplateByIdResponse, err error) {
	if request == nil {
		request = NewGetSpaceTemplateByIdRequest()
	}
	response = NewGetSpaceTemplateByIdResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUpgradeChangeRequest() (request *DeleteUpgradeChangeRequest) {
	request = &DeleteUpgradeChangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "DeleteUpgradeChange")
	return
}

func NewDeleteUpgradeChangeResponse() (response *DeleteUpgradeChangeResponse) {
	response = &DeleteUpgradeChangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除升级变更配置
func (c *Client) DeleteUpgradeChange(request *DeleteUpgradeChangeRequest) (response *DeleteUpgradeChangeResponse, err error) {
	if request == nil {
		request = NewDeleteUpgradeChangeRequest()
	}
	response = NewDeleteUpgradeChangeResponse()
	err = c.Send(request, response)
	return
}

func NewGetModuleVersionRequest() (request *GetModuleVersionRequest) {
	request = &GetModuleVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetModuleVersion")
	return
}

func NewGetModuleVersionResponse() (response *GetModuleVersionResponse) {
	response = &GetModuleVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定模块版本信息
func (c *Client) GetModuleVersion(request *GetModuleVersionRequest) (response *GetModuleVersionResponse, err error) {
	if request == nil {
		request = NewGetModuleVersionRequest()
	}
	response = NewGetModuleVersionResponse()
	err = c.Send(request, response)
	return
}

func NewGetSpaceByIdRequest() (request *GetSpaceByIdRequest) {
	request = &GetSpaceByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetSpaceById")
	return
}

func NewGetSpaceByIdResponse() (response *GetSpaceByIdResponse) {
	response = &GetSpaceByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Id获取Space
func (c *Client) GetSpaceById(request *GetSpaceByIdRequest) (response *GetSpaceByIdResponse, err error) {
	if request == nil {
		request = NewGetSpaceByIdRequest()
	}
	response = NewGetSpaceByIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetGroupInfosRequest() (request *GetGroupInfosRequest) {
	request = &GetGroupInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetGroupInfos")
	return
}

func NewGetGroupInfosResponse() (response *GetGroupInfosResponse) {
	response = &GetGroupInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取分组信息
func (c *Client) GetGroupInfos(request *GetGroupInfosRequest) (response *GetGroupInfosResponse, err error) {
	if request == nil {
		request = NewGetGroupInfosRequest()
	}
	response = NewGetGroupInfosResponse()
	err = c.Send(request, response)
	return
}

func NewGetOnlineVersionRequest() (request *GetOnlineVersionRequest) {
	request = &GetOnlineVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetOnlineVersion")
	return
}

func NewGetOnlineVersionResponse() (response *GetOnlineVersionResponse) {
	response = &GetOnlineVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取在线版本
func (c *Client) GetOnlineVersion(request *GetOnlineVersionRequest) (response *GetOnlineVersionResponse, err error) {
	if request == nil {
		request = NewGetOnlineVersionRequest()
	}
	response = NewGetOnlineVersionResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationTasksRequest() (request *ListOperationTasksRequest) {
	request = &ListOperationTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListOperationTasks")
	return
}

func NewListOperationTasksResponse() (response *ListOperationTasksResponse) {
	response = &ListOperationTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取操作任务
func (c *Client) ListOperationTasks(request *ListOperationTasksRequest) (response *ListOperationTasksResponse, err error) {
	if request == nil {
		request = NewListOperationTasksRequest()
	}
	response = NewListOperationTasksResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSpacePermissionRequest() (request *UpdateSpacePermissionRequest) {
	request = &UpdateSpacePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "UpdateSpacePermission")
	return
}

func NewUpdateSpacePermissionResponse() (response *UpdateSpacePermissionResponse) {
	response = &UpdateSpacePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Space权限
func (c *Client) UpdateSpacePermission(request *UpdateSpacePermissionRequest) (response *UpdateSpacePermissionResponse, err error) {
	if request == nil {
		request = NewUpdateSpacePermissionRequest()
	}
	response = NewUpdateSpacePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewControlNodeStatsRequest() (request *ControlNodeStatsRequest) {
	request = &ControlNodeStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ControlNodeStats")
	return
}

func NewControlNodeStatsResponse() (response *ControlNodeStatsResponse) {
	response = &ControlNodeStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取控制模块服务状态
func (c *Client) ControlNodeStats(request *ControlNodeStatsRequest) (response *ControlNodeStatsResponse, err error) {
	if request == nil {
		request = NewControlNodeStatsRequest()
	}
	response = NewControlNodeStatsResponse()
	err = c.Send(request, response)
	return
}

func NewListDiskRequest() (request *ListDiskRequest) {
	request = &ListDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListDisk")
	return
}

func NewListDiskResponse() (response *ListDiskResponse) {
	response = &ListDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取硬盘列表
func (c *Client) ListDisk(request *ListDiskRequest) (response *ListDiskResponse, err error) {
	if request == nil {
		request = NewListDiskRequest()
	}
	response = NewListDiskResponse()
	err = c.Send(request, response)
	return
}

func NewListDiskByIpsRequest() (request *ListDiskByIpsRequest) {
	request = &ListDiskByIpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListDiskByIps")
	return
}

func NewListDiskByIpsResponse() (response *ListDiskByIpsResponse) {
	response = &ListDiskByIpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机器上的硬盘信息
func (c *Client) ListDiskByIps(request *ListDiskByIpsRequest) (response *ListDiskByIpsResponse, err error) {
	if request == nil {
		request = NewListDiskByIpsRequest()
	}
	response = NewListDiskByIpsResponse()
	err = c.Send(request, response)
	return
}

func NewListLibraJobsRequest() (request *ListLibraJobsRequest) {
	request = &ListLibraJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListLibraJobs")
	return
}

func NewListLibraJobsResponse() (response *ListLibraJobsResponse) {
	response = &ListLibraJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Libra任务查询
func (c *Client) ListLibraJobs(request *ListLibraJobsRequest) (response *ListLibraJobsResponse, err error) {
	if request == nil {
		request = NewListLibraJobsRequest()
	}
	response = NewListLibraJobsResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterDashboardRequest() (request *GetClusterDashboardRequest) {
	request = &GetClusterDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetClusterDashboard")
	return
}

func NewGetClusterDashboardResponse() (response *GetClusterDashboardResponse) {
	response = &GetClusterDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群总览数据
func (c *Client) GetClusterDashboard(request *GetClusterDashboardRequest) (response *GetClusterDashboardResponse, err error) {
	if request == nil {
		request = NewGetClusterDashboardRequest()
	}
	response = NewGetClusterDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewAddErrorDiskRequest() (request *AddErrorDiskRequest) {
	request = &AddErrorDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "AddErrorDisk")
	return
}

func NewAddErrorDiskResponse() (response *AddErrorDiskResponse) {
	response = &AddErrorDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加坏盘
func (c *Client) AddErrorDisk(request *AddErrorDiskRequest) (response *AddErrorDiskResponse, err error) {
	if request == nil {
		request = NewAddErrorDiskRequest()
	}
	response = NewAddErrorDiskResponse()
	err = c.Send(request, response)
	return
}

func NewEnableNodesRequest() (request *EnableNodesRequest) {
	request = &EnableNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "EnableNodes")
	return
}

func NewEnableNodesResponse() (response *EnableNodesResponse) {
	response = &EnableNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 使能节点
func (c *Client) EnableNodes(request *EnableNodesRequest) (response *EnableNodesResponse, err error) {
	if request == nil {
		request = NewEnableNodesRequest()
	}
	response = NewEnableNodesResponse()
	err = c.Send(request, response)
	return
}

func NewGetErrorDiskChangeLogRequest() (request *GetErrorDiskChangeLogRequest) {
	request = &GetErrorDiskChangeLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetErrorDiskChangeLog")
	return
}

func NewGetErrorDiskChangeLogResponse() (response *GetErrorDiskChangeLogResponse) {
	response = &GetErrorDiskChangeLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看坏盘流程详情
func (c *Client) GetErrorDiskChangeLog(request *GetErrorDiskChangeLogRequest) (response *GetErrorDiskChangeLogResponse, err error) {
	if request == nil {
		request = NewGetErrorDiskChangeLogRequest()
	}
	response = NewGetErrorDiskChangeLogResponse()
	err = c.Send(request, response)
	return
}

func NewStartBalanceRequest() (request *StartBalanceRequest) {
	request = &StartBalanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "StartBalance")
	return
}

func NewStartBalanceResponse() (response *StartBalanceResponse) {
	response = &StartBalanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发起均衡
func (c *Client) StartBalance(request *StartBalanceRequest) (response *StartBalanceResponse, err error) {
	if request == nil {
		request = NewStartBalanceRequest()
	}
	response = NewStartBalanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVolumesRequest() (request *CreateVolumesRequest) {
	request = &CreateVolumesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "CreateVolumes")
	return
}

func NewCreateVolumesResponse() (response *CreateVolumesResponse) {
	response = &CreateVolumesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Volume
func (c *Client) CreateVolumes(request *CreateVolumesRequest) (response *CreateVolumesResponse, err error) {
	if request == nil {
		request = NewCreateVolumesRequest()
	}
	response = NewCreateVolumesResponse()
	err = c.Send(request, response)
	return
}

func NewGetRequestTrendRequest() (request *GetRequestTrendRequest) {
	request = &GetRequestTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetRequestTrend")
	return
}

func NewGetRequestTrendResponse() (response *GetRequestTrendResponse) {
	response = &GetRequestTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取请求趋势
func (c *Client) GetRequestTrend(request *GetRequestTrendRequest) (response *GetRequestTrendResponse, err error) {
	if request == nil {
		request = NewGetRequestTrendRequest()
	}
	response = NewGetRequestTrendResponse()
	err = c.Send(request, response)
	return
}

func NewGetDeployableHostRequest() (request *GetDeployableHostRequest) {
	request = &GetDeployableHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetDeployableHost")
	return
}

func NewGetDeployableHostResponse() (response *GetDeployableHostResponse) {
	response = &GetDeployableHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetDeployableHost
func (c *Client) GetDeployableHost(request *GetDeployableHostRequest) (response *GetDeployableHostResponse, err error) {
	if request == nil {
		request = NewGetDeployableHostRequest()
	}
	response = NewGetDeployableHostResponse()
	err = c.Send(request, response)
	return
}

func NewSubmitHandleErrorDiskTaskRequest() (request *SubmitHandleErrorDiskTaskRequest) {
	request = &SubmitHandleErrorDiskTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "SubmitHandleErrorDiskTask")
	return
}

func NewSubmitHandleErrorDiskTaskResponse() (response *SubmitHandleErrorDiskTaskResponse) {
	response = &SubmitHandleErrorDiskTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发起坏盘任务
func (c *Client) SubmitHandleErrorDiskTask(request *SubmitHandleErrorDiskTaskRequest) (response *SubmitHandleErrorDiskTaskResponse, err error) {
	if request == nil {
		request = NewSubmitHandleErrorDiskTaskRequest()
	}
	response = NewSubmitHandleErrorDiskTaskResponse()
	err = c.Send(request, response)
	return
}

func NewNewClusterConfirmUpdateByIdRequest() (request *NewClusterConfirmUpdateByIdRequest) {
	request = &NewClusterConfirmUpdateByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "NewClusterConfirmUpdateById")
	return
}

func NewNewClusterConfirmUpdateByIdResponse() (response *NewClusterConfirmUpdateByIdResponse) {
	response = &NewClusterConfirmUpdateByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上线集群
func (c *Client) NewClusterConfirmUpdateById(request *NewClusterConfirmUpdateByIdRequest) (response *NewClusterConfirmUpdateByIdResponse, err error) {
	if request == nil {
		request = NewNewClusterConfirmUpdateByIdRequest()
	}
	response = NewNewClusterConfirmUpdateByIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetTaskByUuidRequest() (request *GetTaskByUuidRequest) {
	request = &GetTaskByUuidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetTaskByUuid")
	return
}

func NewGetTaskByUuidResponse() (response *GetTaskByUuidResponse) {
	response = &GetTaskByUuidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID获取任务信息
func (c *Client) GetTaskByUuid(request *GetTaskByUuidRequest) (response *GetTaskByUuidResponse, err error) {
	if request == nil {
		request = NewGetTaskByUuidRequest()
	}
	response = NewGetTaskByUuidResponse()
	err = c.Send(request, response)
	return
}

func NewListUpgradeChangeFormsRequest() (request *ListUpgradeChangeFormsRequest) {
	request = &ListUpgradeChangeFormsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListUpgradeChangeForms")
	return
}

func NewListUpgradeChangeFormsResponse() (response *ListUpgradeChangeFormsResponse) {
	response = &ListUpgradeChangeFormsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取变更升级表
func (c *Client) ListUpgradeChangeForms(request *ListUpgradeChangeFormsRequest) (response *ListUpgradeChangeFormsResponse, err error) {
	if request == nil {
		request = NewListUpgradeChangeFormsRequest()
	}
	response = NewListUpgradeChangeFormsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterByIdRequest() (request *DeleteClusterByIdRequest) {
	request = &DeleteClusterByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "DeleteClusterById")
	return
}

func NewDeleteClusterByIdResponse() (response *DeleteClusterByIdResponse) {
	response = &DeleteClusterByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据ID删除Cluster
func (c *Client) DeleteClusterById(request *DeleteClusterByIdRequest) (response *DeleteClusterByIdResponse, err error) {
	if request == nil {
		request = NewDeleteClusterByIdRequest()
	}
	response = NewDeleteClusterByIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubTaskInfoRequest() (request *GetSubTaskInfoRequest) {
	request = &GetSubTaskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetSubTaskInfo")
	return
}

func NewGetSubTaskInfoResponse() (response *GetSubTaskInfoResponse) {
	response = &GetSubTaskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子任务详情
func (c *Client) GetSubTaskInfo(request *GetSubTaskInfoRequest) (response *GetSubTaskInfoResponse, err error) {
	if request == nil {
		request = NewGetSubTaskInfoRequest()
	}
	response = NewGetSubTaskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetGlobalDashboardRequest() (request *GetGlobalDashboardRequest) {
	request = &GetGlobalDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetGlobalDashboard")
	return
}

func NewGetGlobalDashboardResponse() (response *GetGlobalDashboardResponse) {
	response = &GetGlobalDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取大盘展示数据
func (c *Client) GetGlobalDashboard(request *GetGlobalDashboardRequest) (response *GetGlobalDashboardResponse, err error) {
	if request == nil {
		request = NewGetGlobalDashboardRequest()
	}
	response = NewGetGlobalDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterByIdRequest() (request *GetClusterByIdRequest) {
	request = &GetClusterByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetClusterById")
	return
}

func NewGetClusterByIdResponse() (response *GetClusterByIdResponse) {
	response = &GetClusterByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Id获取集群信息
func (c *Client) GetClusterById(request *GetClusterByIdRequest) (response *GetClusterByIdResponse, err error) {
	if request == nil {
		request = NewGetClusterByIdRequest()
	}
	response = NewGetClusterByIdResponse()
	err = c.Send(request, response)
	return
}

func NewListSpacesRequest() (request *ListSpacesRequest) {
	request = &ListSpacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListSpaces")
	return
}

func NewListSpacesResponse() (response *ListSpacesResponse) {
	response = &ListSpacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Space
func (c *Client) ListSpaces(request *ListSpacesRequest) (response *ListSpacesResponse, err error) {
	if request == nil {
		request = NewListSpacesRequest()
	}
	response = NewListSpacesResponse()
	err = c.Send(request, response)
	return
}

func NewRedoUpgradeSubTaskRequest() (request *RedoUpgradeSubTaskRequest) {
	request = &RedoUpgradeSubTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "RedoUpgradeSubTask")
	return
}

func NewRedoUpgradeSubTaskResponse() (response *RedoUpgradeSubTaskResponse) {
	response = &RedoUpgradeSubTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重做升级子任务
func (c *Client) RedoUpgradeSubTask(request *RedoUpgradeSubTaskRequest) (response *RedoUpgradeSubTaskResponse, err error) {
	if request == nil {
		request = NewRedoUpgradeSubTaskRequest()
	}
	response = NewRedoUpgradeSubTaskResponse()
	err = c.Send(request, response)
	return
}

func NewNewClusterConfirmRequest() (request *NewClusterConfirmRequest) {
	request = &NewClusterConfirmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "NewClusterConfirm")
	return
}

func NewNewClusterConfirmResponse() (response *NewClusterConfirmResponse) {
	response = &NewClusterConfirmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 确认部署新集群
func (c *Client) NewClusterConfirm(request *NewClusterConfirmRequest) (response *NewClusterConfirmResponse, err error) {
	if request == nil {
		request = NewNewClusterConfirmRequest()
	}
	response = NewNewClusterConfirmResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateClusterByIdRequest() (request *UpdateClusterByIdRequest) {
	request = &UpdateClusterByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "UpdateClusterById")
	return
}

func NewUpdateClusterByIdResponse() (response *UpdateClusterByIdResponse) {
	response = &UpdateClusterByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Id更新Cluster
func (c *Client) UpdateClusterById(request *UpdateClusterByIdRequest) (response *UpdateClusterByIdResponse, err error) {
	if request == nil {
		request = NewUpdateClusterByIdRequest()
	}
	response = NewUpdateClusterByIdResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeNodesRequest() (request *UpgradeNodesRequest) {
	request = &UpgradeNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "UpgradeNodes")
	return
}

func NewUpgradeNodesResponse() (response *UpgradeNodesResponse) {
	response = &UpgradeNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级服务Nodes
func (c *Client) UpgradeNodes(request *UpgradeNodesRequest) (response *UpgradeNodesResponse, err error) {
	if request == nil {
		request = NewUpgradeNodesRequest()
	}
	response = NewUpgradeNodesResponse()
	err = c.Send(request, response)
	return
}

func NewGetNodeVletsRequest() (request *GetNodeVletsRequest) {
	request = &GetNodeVletsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetNodeVlets")
	return
}

func NewGetNodeVletsResponse() (response *GetNodeVletsResponse) {
	response = &GetNodeVletsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点Vlets信息
func (c *Client) GetNodeVlets(request *GetNodeVletsRequest) (response *GetNodeVletsResponse, err error) {
	if request == nil {
		request = NewGetNodeVletsRequest()
	}
	response = NewGetNodeVletsResponse()
	err = c.Send(request, response)
	return
}

func NewSubmitClusterStageRequest() (request *SubmitClusterStageRequest) {
	request = &SubmitClusterStageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "SubmitClusterStage")
	return
}

func NewSubmitClusterStageResponse() (response *SubmitClusterStageResponse) {
	response = &SubmitClusterStageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交表单信息
func (c *Client) SubmitClusterStage(request *SubmitClusterStageRequest) (response *SubmitClusterStageResponse, err error) {
	if request == nil {
		request = NewSubmitClusterStageRequest()
	}
	response = NewSubmitClusterStageResponse()
	err = c.Send(request, response)
	return
}

func NewNewClusterStage3Request() (request *NewClusterStage3Request) {
	request = &NewClusterStage3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "NewClusterStage3")
	return
}

func NewNewClusterStage3Response() (response *NewClusterStage3Response) {
	response = &NewClusterStage3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取初始化表单信息阶段3
func (c *Client) NewClusterStage3(request *NewClusterStage3Request) (response *NewClusterStage3Response, err error) {
	if request == nil {
		request = NewNewClusterStage3Request()
	}
	response = NewNewClusterStage3Response()
	err = c.Send(request, response)
	return
}

func NewOfflineNodesRequest() (request *OfflineNodesRequest) {
	request = &OfflineNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "OfflineNodes")
	return
}

func NewOfflineNodesResponse() (response *OfflineNodesResponse) {
	response = &OfflineNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务Node下线
func (c *Client) OfflineNodes(request *OfflineNodesRequest) (response *OfflineNodesResponse, err error) {
	if request == nil {
		request = NewOfflineNodesRequest()
	}
	response = NewOfflineNodesResponse()
	err = c.Send(request, response)
	return
}

func NewGetDiskStatsRequest() (request *GetDiskStatsRequest) {
	request = &GetDiskStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetDiskStats")
	return
}

func NewGetDiskStatsResponse() (response *GetDiskStatsResponse) {
	response = &GetDiskStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取硬盘统计数据
func (c *Client) GetDiskStats(request *GetDiskStatsRequest) (response *GetDiskStatsResponse, err error) {
	if request == nil {
		request = NewGetDiskStatsRequest()
	}
	response = NewGetDiskStatsResponse()
	err = c.Send(request, response)
	return
}

func NewListLibraRelatedJobsRequest() (request *ListLibraRelatedJobsRequest) {
	request = &ListLibraRelatedJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListLibraRelatedJobs")
	return
}

func NewListLibraRelatedJobsResponse() (response *ListLibraRelatedJobsResponse) {
	response = &ListLibraRelatedJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Libra关联任务查询
func (c *Client) ListLibraRelatedJobs(request *ListLibraRelatedJobsRequest) (response *ListLibraRelatedJobsResponse, err error) {
	if request == nil {
		request = NewListLibraRelatedJobsRequest()
	}
	response = NewListLibraRelatedJobsResponse()
	err = c.Send(request, response)
	return
}

func NewNewClusterStage1Request() (request *NewClusterStage1Request) {
	request = &NewClusterStage1Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "NewClusterStage1")
	return
}

func NewNewClusterStage1Response() (response *NewClusterStage1Response) {
	response = &NewClusterStage1Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取初始化表单信息阶段1
func (c *Client) NewClusterStage1(request *NewClusterStage1Request) (response *NewClusterStage1Response, err error) {
	if request == nil {
		request = NewNewClusterStage1Request()
	}
	response = NewNewClusterStage1Response()
	err = c.Send(request, response)
	return
}

func NewListClustersRequest() (request *ListClustersRequest) {
	request = &ListClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListClusters")
	return
}

func NewListClustersResponse() (response *ListClustersResponse) {
	response = &ListClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Cluster列表
func (c *Client) ListClusters(request *ListClustersRequest) (response *ListClustersResponse, err error) {
	if request == nil {
		request = NewListClustersRequest()
	}
	response = NewListClustersResponse()
	err = c.Send(request, response)
	return
}

func NewBlackNodesRequest() (request *BlackNodesRequest) {
	request = &BlackNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "BlackNodes")
	return
}

func NewBlackNodesResponse() (response *BlackNodesResponse) {
	response = &BlackNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 屏蔽服务Nodes
func (c *Client) BlackNodes(request *BlackNodesRequest) (response *BlackNodesResponse, err error) {
	if request == nil {
		request = NewBlackNodesRequest()
	}
	response = NewBlackNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDeployNodesRequest() (request *DeployNodesRequest) {
	request = &DeployNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "DeployNodes")
	return
}

func NewDeployNodesResponse() (response *DeployNodesResponse) {
	response = &DeployNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署服务Nodes
func (c *Client) DeployNodes(request *DeployNodesRequest) (response *DeployNodesResponse, err error) {
	if request == nil {
		request = NewDeployNodesRequest()
	}
	response = NewDeployNodesResponse()
	err = c.Send(request, response)
	return
}

func NewRetryChainTaskByUuidRequest() (request *RetryChainTaskByUuidRequest) {
	request = &RetryChainTaskByUuidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "RetryChainTaskByUuid")
	return
}

func NewRetryChainTaskByUuidResponse() (response *RetryChainTaskByUuidResponse) {
	response = &RetryChainTaskByUuidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试序列任务
func (c *Client) RetryChainTaskByUuid(request *RetryChainTaskByUuidRequest) (response *RetryChainTaskByUuidResponse, err error) {
	if request == nil {
		request = NewRetryChainTaskByUuidRequest()
	}
	response = NewRetryChainTaskByUuidResponse()
	err = c.Send(request, response)
	return
}

func NewGetStorageTrendRequest() (request *GetStorageTrendRequest) {
	request = &GetStorageTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetStorageTrend")
	return
}

func NewGetStorageTrendResponse() (response *GetStorageTrendResponse) {
	response = &GetStorageTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储趋势
func (c *Client) GetStorageTrend(request *GetStorageTrendRequest) (response *GetStorageTrendResponse, err error) {
	if request == nil {
		request = NewGetStorageTrendRequest()
	}
	response = NewGetStorageTrendResponse()
	err = c.Send(request, response)
	return
}

func NewRollBackSubTaskRequest() (request *RollBackSubTaskRequest) {
	request = &RollBackSubTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "RollBackSubTask")
	return
}

func NewRollBackSubTaskResponse() (response *RollBackSubTaskResponse) {
	response = &RollBackSubTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚某个子任务
func (c *Client) RollBackSubTask(request *RollBackSubTaskRequest) (response *RollBackSubTaskResponse, err error) {
	if request == nil {
		request = NewRollBackSubTaskRequest()
	}
	response = NewRollBackSubTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGinDoPlanRequest() (request *GinDoPlanRequest) {
	request = &GinDoPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GinDoPlan")
	return
}

func NewGinDoPlanResponse() (response *GinDoPlanResponse) {
	response = &GinDoPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行扩容计划
func (c *Client) GinDoPlan(request *GinDoPlanRequest) (response *GinDoPlanResponse, err error) {
	if request == nil {
		request = NewGinDoPlanRequest()
	}
	response = NewGinDoPlanResponse()
	err = c.Send(request, response)
	return
}

func NewAllClustersOverviewRequest() (request *AllClustersOverviewRequest) {
	request = &AllClustersOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "AllClustersOverview")
	return
}

func NewAllClustersOverviewResponse() (response *AllClustersOverviewResponse) {
	response = &AllClustersOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 全集群概览
func (c *Client) AllClustersOverview(request *AllClustersOverviewRequest) (response *AllClustersOverviewResponse, err error) {
	if request == nil {
		request = NewAllClustersOverviewRequest()
	}
	response = NewAllClustersOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewGinMakePackageInfoRequest() (request *GinMakePackageInfoRequest) {
	request = &GinMakePackageInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GinMakePackageInfo")
	return
}

func NewGinMakePackageInfoResponse() (response *GinMakePackageInfoResponse) {
	response = &GinMakePackageInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安装包版本
func (c *Client) GinMakePackageInfo(request *GinMakePackageInfoRequest) (response *GinMakePackageInfoResponse, err error) {
	if request == nil {
		request = NewGinMakePackageInfoRequest()
	}
	response = NewGinMakePackageInfoResponse()
	err = c.Send(request, response)
	return
}

func NewNewClusterStage2Request() (request *NewClusterStage2Request) {
	request = &NewClusterStage2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "NewClusterStage2")
	return
}

func NewNewClusterStage2Response() (response *NewClusterStage2Response) {
	response = &NewClusterStage2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取初始化表单信息阶段2
func (c *Client) NewClusterStage2(request *NewClusterStage2Request) (response *NewClusterStage2Response, err error) {
	if request == nil {
		request = NewNewClusterStage2Request()
	}
	response = NewNewClusterStage2Response()
	err = c.Send(request, response)
	return
}

func NewCreateSpaceRequest() (request *CreateSpaceRequest) {
	request = &CreateSpaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "CreateSpace")
	return
}

func NewCreateSpaceResponse() (response *CreateSpaceResponse) {
	response = &CreateSpaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Space
func (c *Client) CreateSpace(request *CreateSpaceRequest) (response *CreateSpaceResponse, err error) {
	if request == nil {
		request = NewCreateSpaceRequest()
	}
	response = NewCreateSpaceResponse()
	err = c.Send(request, response)
	return
}

func NewGetChainTaskByUuidRequest() (request *GetChainTaskByUuidRequest) {
	request = &GetChainTaskByUuidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetChainTaskByUuid")
	return
}

func NewGetChainTaskByUuidResponse() (response *GetChainTaskByUuidResponse) {
	response = &GetChainTaskByUuidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取序列任务信息
func (c *Client) GetChainTaskByUuid(request *GetChainTaskByUuidRequest) (response *GetChainTaskByUuidResponse, err error) {
	if request == nil {
		request = NewGetChainTaskByUuidRequest()
	}
	response = NewGetChainTaskByUuidResponse()
	err = c.Send(request, response)
	return
}

func NewListSpaceTemplateRequest() (request *ListSpaceTemplateRequest) {
	request = &ListSpaceTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListSpaceTemplate")
	return
}

func NewListSpaceTemplateResponse() (response *ListSpaceTemplateResponse) {
	response = &ListSpaceTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Space模版
func (c *Client) ListSpaceTemplate(request *ListSpaceTemplateRequest) (response *ListSpaceTemplateResponse, err error) {
	if request == nil {
		request = NewListSpaceTemplateRequest()
	}
	response = NewListSpaceTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewUnregisteredNodesRequest() (request *UnregisteredNodesRequest) {
	request = &UnregisteredNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "UnregisteredNodes")
	return
}

func NewUnregisteredNodesResponse() (response *UnregisteredNodesResponse) {
	response = &UnregisteredNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 注销服务节点
func (c *Client) UnregisteredNodes(request *UnregisteredNodesRequest) (response *UnregisteredNodesResponse, err error) {
	if request == nil {
		request = NewUnregisteredNodesRequest()
	}
	response = NewUnregisteredNodesResponse()
	err = c.Send(request, response)
	return
}

func NewAddHostsRequest() (request *AddHostsRequest) {
	request = &AddHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "AddHosts")
	return
}

func NewAddHostsResponse() (response *AddHostsResponse) {
	response = &AddHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// AddHosts
func (c *Client) AddHosts(request *AddHostsRequest) (response *AddHostsResponse, err error) {
	if request == nil {
		request = NewAddHostsRequest()
	}
	response = NewAddHostsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryClusterFormByIdRequest() (request *QueryClusterFormByIdRequest) {
	request = &QueryClusterFormByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "QueryClusterFormById")
	return
}

func NewQueryClusterFormByIdResponse() (response *QueryClusterFormByIdResponse) {
	response = &QueryClusterFormByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Id查询集群部署单
func (c *Client) QueryClusterFormById(request *QueryClusterFormByIdRequest) (response *QueryClusterFormByIdResponse, err error) {
	if request == nil {
		request = NewQueryClusterFormByIdRequest()
	}
	response = NewQueryClusterFormByIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetAvailabilityTrendRequest() (request *GetAvailabilityTrendRequest) {
	request = &GetAvailabilityTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetAvailabilityTrend")
	return
}

func NewGetAvailabilityTrendResponse() (response *GetAvailabilityTrendResponse) {
	response = &GetAvailabilityTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取可用性趋势
func (c *Client) GetAvailabilityTrend(request *GetAvailabilityTrendRequest) (response *GetAvailabilityTrendResponse, err error) {
	if request == nil {
		request = NewGetAvailabilityTrendRequest()
	}
	response = NewGetAvailabilityTrendResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpgradeChangesRequest() (request *GetUpgradeChangesRequest) {
	request = &GetUpgradeChangesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetUpgradeChanges")
	return
}

func NewGetUpgradeChangesResponse() (response *GetUpgradeChangesResponse) {
	response = &GetUpgradeChangesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据升级变更名称获取详细信息
func (c *Client) GetUpgradeChanges(request *GetUpgradeChangesRequest) (response *GetUpgradeChangesResponse, err error) {
	if request == nil {
		request = NewGetUpgradeChangesRequest()
	}
	response = NewGetUpgradeChangesResponse()
	err = c.Send(request, response)
	return
}

func NewRetryErrorDiskRequest() (request *RetryErrorDiskRequest) {
	request = &RetryErrorDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "RetryErrorDisk")
	return
}

func NewRetryErrorDiskResponse() (response *RetryErrorDiskResponse) {
	response = &RetryErrorDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试坏盘任务
func (c *Client) RetryErrorDisk(request *RetryErrorDiskRequest) (response *RetryErrorDiskResponse, err error) {
	if request == nil {
		request = NewRetryErrorDiskRequest()
	}
	response = NewRetryErrorDiskResponse()
	err = c.Send(request, response)
	return
}

func NewQueryChainTaskRequest() (request *QueryChainTaskRequest) {
	request = &QueryChainTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "QueryChainTask")
	return
}

func NewQueryChainTaskResponse() (response *QueryChainTaskResponse) {
	response = &QueryChainTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务信息
func (c *Client) QueryChainTask(request *QueryChainTaskRequest) (response *QueryChainTaskResponse, err error) {
	if request == nil {
		request = NewQueryChainTaskRequest()
	}
	response = NewQueryChainTaskResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveHostsRequest() (request *RemoveHostsRequest) {
	request = &RemoveHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "RemoveHosts")
	return
}

func NewRemoveHostsResponse() (response *RemoveHostsResponse) {
	response = &RemoveHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 机器下架
func (c *Client) RemoveHosts(request *RemoveHostsRequest) (response *RemoveHostsResponse, err error) {
	if request == nil {
		request = NewRemoveHostsRequest()
	}
	response = NewRemoveHostsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSpaceThresholdRequest() (request *UpdateSpaceThresholdRequest) {
	request = &UpdateSpaceThresholdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "UpdateSpaceThreshold")
	return
}

func NewUpdateSpaceThresholdResponse() (response *UpdateSpaceThresholdResponse) {
	response = &UpdateSpaceThresholdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Space阀值
func (c *Client) UpdateSpaceThreshold(request *UpdateSpaceThresholdRequest) (response *UpdateSpaceThresholdResponse, err error) {
	if request == nil {
		request = NewUpdateSpaceThresholdRequest()
	}
	response = NewUpdateSpaceThresholdResponse()
	err = c.Send(request, response)
	return
}

func NewListLibraSummaryRequest() (request *ListLibraSummaryRequest) {
	request = &ListLibraSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListLibraSummary")
	return
}

func NewListLibraSummaryResponse() (response *ListLibraSummaryResponse) {
	response = &ListLibraSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Libra任务详情总览
func (c *Client) ListLibraSummary(request *ListLibraSummaryRequest) (response *ListLibraSummaryResponse, err error) {
	if request == nil {
		request = NewListLibraSummaryRequest()
	}
	response = NewListLibraSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewNewClusterApproveRequest() (request *NewClusterApproveRequest) {
	request = &NewClusterApproveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "NewClusterApprove")
	return
}

func NewNewClusterApproveResponse() (response *NewClusterApproveResponse) {
	response = &NewClusterApproveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 审批新集群部署单
func (c *Client) NewClusterApprove(request *NewClusterApproveRequest) (response *NewClusterApproveResponse, err error) {
	if request == nil {
		request = NewNewClusterApproveRequest()
	}
	response = NewNewClusterApproveResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostStatsRequest() (request *GetHostStatsRequest) {
	request = &GetHostStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetHostStats")
	return
}

func NewGetHostStatsResponse() (response *GetHostStatsResponse) {
	response = &GetHostStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机统计数据
func (c *Client) GetHostStats(request *GetHostStatsRequest) (response *GetHostStatsResponse, err error) {
	if request == nil {
		request = NewGetHostStatsRequest()
	}
	response = NewGetHostStatsResponse()
	err = c.Send(request, response)
	return
}

func NewGetVersionInfoRequest() (request *GetVersionInfoRequest) {
	request = &GetVersionInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetVersionInfo")
	return
}

func NewGetVersionInfoResponse() (response *GetVersionInfoResponse) {
	response = &GetVersionInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务版本信息
func (c *Client) GetVersionInfo(request *GetVersionInfoRequest) (response *GetVersionInfoResponse, err error) {
	if request == nil {
		request = NewGetVersionInfoRequest()
	}
	response = NewGetVersionInfoResponse()
	err = c.Send(request, response)
	return
}

func NewListHostsRequest() (request *ListHostsRequest) {
	request = &ListHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListHosts")
	return
}

func NewListHostsResponse() (response *ListHostsResponse) {
	response = &ListHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Hosts
func (c *Client) ListHosts(request *ListHostsRequest) (response *ListHostsResponse, err error) {
	if request == nil {
		request = NewListHostsRequest()
	}
	response = NewListHostsResponse()
	err = c.Send(request, response)
	return
}

func NewNewClusterStage0Request() (request *NewClusterStage0Request) {
	request = &NewClusterStage0Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "NewClusterStage0")
	return
}

func NewNewClusterStage0Response() (response *NewClusterStage0Response) {
	response = &NewClusterStage0Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取初始化表单信息阶段0
func (c *Client) NewClusterStage0(request *NewClusterStage0Request) (response *NewClusterStage0Response, err error) {
	if request == nil {
		request = NewNewClusterStage0Request()
	}
	response = NewNewClusterStage0Response()
	err = c.Send(request, response)
	return
}

func NewResetClusterStageRequest() (request *ResetClusterStageRequest) {
	request = &ResetClusterStageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ResetClusterStage")
	return
}

func NewResetClusterStageResponse() (response *ResetClusterStageResponse) {
	response = &ResetClusterStageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置表单信息
func (c *Client) ResetClusterStage(request *ResetClusterStageRequest) (response *ResetClusterStageResponse, err error) {
	if request == nil {
		request = NewResetClusterStageRequest()
	}
	response = NewResetClusterStageResponse()
	err = c.Send(request, response)
	return
}

func NewAbortBalanceRequest() (request *AbortBalanceRequest) {
	request = &AbortBalanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "AbortBalance")
	return
}

func NewAbortBalanceResponse() (response *AbortBalanceResponse) {
	response = &AbortBalanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 终止均衡
func (c *Client) AbortBalance(request *AbortBalanceRequest) (response *AbortBalanceResponse, err error) {
	if request == nil {
		request = NewAbortBalanceRequest()
	}
	response = NewAbortBalanceResponse()
	err = c.Send(request, response)
	return
}

func NewGetColumbusListRequest() (request *GetColumbusListRequest) {
	request = &GetColumbusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetColumbusList")
	return
}

func NewGetColumbusListResponse() (response *GetColumbusListResponse) {
	response = &GetColumbusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取columbus列表
func (c *Client) GetColumbusList(request *GetColumbusListRequest) (response *GetColumbusListResponse, err error) {
	if request == nil {
		request = NewGetColumbusListRequest()
	}
	response = NewGetColumbusListResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubChangeTaskRequest() (request *GetSubChangeTaskRequest) {
	request = &GetSubChangeTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetSubChangeTask")
	return
}

func NewGetSubChangeTaskResponse() (response *GetSubChangeTaskResponse) {
	response = &GetSubChangeTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取升级变更子任务
func (c *Client) GetSubChangeTask(request *GetSubChangeTaskRequest) (response *GetSubChangeTaskResponse, err error) {
	if request == nil {
		request = NewGetSubChangeTaskRequest()
	}
	response = NewGetSubChangeTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetTasksRequest() (request *GetTasksRequest) {
	request = &GetTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetTasks")
	return
}

func NewGetTasksResponse() (response *GetTasksResponse) {
	response = &GetTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务信息
func (c *Client) GetTasks(request *GetTasksRequest) (response *GetTasksResponse, err error) {
	if request == nil {
		request = NewGetTasksRequest()
	}
	response = NewGetTasksResponse()
	err = c.Send(request, response)
	return
}

func NewGinMakeDeployPlanRequest() (request *GinMakeDeployPlanRequest) {
	request = &GinMakeDeployPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GinMakeDeployPlan")
	return
}

func NewGinMakeDeployPlanResponse() (response *GinMakeDeployPlanResponse) {
	response = &GinMakeDeployPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成扩容计划
func (c *Client) GinMakeDeployPlan(request *GinMakeDeployPlanRequest) (response *GinMakeDeployPlanResponse, err error) {
	if request == nil {
		request = NewGinMakeDeployPlanRequest()
	}
	response = NewGinMakeDeployPlanResponse()
	err = c.Send(request, response)
	return
}

func NewNewColumbusRequest() (request *NewColumbusRequest) {
	request = &NewColumbusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "NewColumbus")
	return
}

func NewNewColumbusResponse() (response *NewColumbusResponse) {
	response = &NewColumbusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// columbus部署
func (c *Client) NewColumbus(request *NewColumbusRequest) (response *NewColumbusResponse, err error) {
	if request == nil {
		request = NewNewColumbusRequest()
	}
	response = NewNewColumbusResponse()
	err = c.Send(request, response)
	return
}

func NewCapacityReportRequest() (request *CapacityReportRequest) {
	request = &CapacityReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "CapacityReport")
	return
}

func NewCapacityReportResponse() (response *CapacityReportResponse) {
	response = &CapacityReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容量总览
func (c *Client) CapacityReport(request *CapacityReportRequest) (response *CapacityReportResponse, err error) {
	if request == nil {
		request = NewCapacityReportRequest()
	}
	response = NewCapacityReportResponse()
	err = c.Send(request, response)
	return
}

func NewDropSpaceRequest() (request *DropSpaceRequest) {
	request = &DropSpaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "DropSpace")
	return
}

func NewDropSpaceResponse() (response *DropSpaceResponse) {
	response = &DropSpaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Space
func (c *Client) DropSpace(request *DropSpaceRequest) (response *DropSpaceResponse, err error) {
	if request == nil {
		request = NewDropSpaceRequest()
	}
	response = NewDropSpaceResponse()
	err = c.Send(request, response)
	return
}

func NewSubmitUpgradeChangesRequest() (request *SubmitUpgradeChangesRequest) {
	request = &SubmitUpgradeChangesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "SubmitUpgradeChanges")
	return
}

func NewSubmitUpgradeChangesResponse() (response *SubmitUpgradeChangesResponse) {
	response = &SubmitUpgradeChangesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交升级变更配置
func (c *Client) SubmitUpgradeChanges(request *SubmitUpgradeChangesRequest) (response *SubmitUpgradeChangesResponse, err error) {
	if request == nil {
		request = NewSubmitUpgradeChangesRequest()
	}
	response = NewSubmitUpgradeChangesResponse()
	err = c.Send(request, response)
	return
}

func NewListBalanceJobsRequest() (request *ListBalanceJobsRequest) {
	request = &ListBalanceJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListBalanceJobs")
	return
}

func NewListBalanceJobsResponse() (response *ListBalanceJobsResponse) {
	response = &ListBalanceJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Balance任务
func (c *Client) ListBalanceJobs(request *ListBalanceJobsRequest) (response *ListBalanceJobsResponse, err error) {
	if request == nil {
		request = NewListBalanceJobsRequest()
	}
	response = NewListBalanceJobsResponse()
	err = c.Send(request, response)
	return
}

func NewListErrorDiskRequest() (request *ListErrorDiskRequest) {
	request = &ListErrorDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListErrorDisk")
	return
}

func NewListErrorDiskResponse() (response *ListErrorDiskResponse) {
	response = &ListErrorDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取坏盘列表
func (c *Client) ListErrorDisk(request *ListErrorDiskRequest) (response *ListErrorDiskResponse, err error) {
	if request == nil {
		request = NewListErrorDiskRequest()
	}
	response = NewListErrorDiskResponse()
	err = c.Send(request, response)
	return
}

func NewListNodesRequest() (request *ListNodesRequest) {
	request = &ListNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ListNodes")
	return
}

func NewListNodesResponse() (response *ListNodesResponse) {
	response = &ListNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有服务Nodes
func (c *Client) ListNodes(request *ListNodesRequest) (response *ListNodesResponse, err error) {
	if request == nil {
		request = NewListNodesRequest()
	}
	response = NewListNodesResponse()
	err = c.Send(request, response)
	return
}

func NewGenerateHandleDiskTaskRequest() (request *GenerateHandleDiskTaskRequest) {
	request = &GenerateHandleDiskTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GenerateHandleDiskTask")
	return
}

func NewGenerateHandleDiskTaskResponse() (response *GenerateHandleDiskTaskResponse) {
	response = &GenerateHandleDiskTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发起换盘任务
func (c *Client) GenerateHandleDiskTask(request *GenerateHandleDiskTaskRequest) (response *GenerateHandleDiskTaskResponse, err error) {
	if request == nil {
		request = NewGenerateHandleDiskTaskRequest()
	}
	response = NewGenerateHandleDiskTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetBlobSpaceRequest() (request *GetBlobSpaceRequest) {
	request = &GetBlobSpaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetBlobSpace")
	return
}

func NewGetBlobSpaceResponse() (response *GetBlobSpaceResponse) {
	response = &GetBlobSpaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物理容量
func (c *Client) GetBlobSpace(request *GetBlobSpaceRequest) (response *GetBlobSpaceResponse, err error) {
	if request == nil {
		request = NewGetBlobSpaceRequest()
	}
	response = NewGetBlobSpaceResponse()
	err = c.Send(request, response)
	return
}

func NewGetRegionListRequest() (request *GetRegionListRequest) {
	request = &GetRegionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetRegionList")
	return
}

func NewGetRegionListResponse() (response *GetRegionListResponse) {
	response = &GetRegionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取地域信息
func (c *Client) GetRegionList(request *GetRegionListRequest) (response *GetRegionListResponse, err error) {
	if request == nil {
		request = NewGetRegionListRequest()
	}
	response = NewGetRegionListResponse()
	err = c.Send(request, response)
	return
}

func NewUnBlackNodesRequest() (request *UnBlackNodesRequest) {
	request = &UnBlackNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "UnBlackNodes")
	return
}

func NewUnBlackNodesResponse() (response *UnBlackNodesResponse) {
	response = &UnBlackNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除服务Nodes屏蔽
func (c *Client) UnBlackNodes(request *UnBlackNodesRequest) (response *UnBlackNodesResponse, err error) {
	if request == nil {
		request = NewUnBlackNodesRequest()
	}
	response = NewUnBlackNodesResponse()
	err = c.Send(request, response)
	return
}

func NewGetNodeStatsRequest() (request *GetNodeStatsRequest) {
	request = &GetNodeStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "GetNodeStats")
	return
}

func NewGetNodeStatsResponse() (response *GetNodeStatsResponse) {
	response = &GetNodeStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务状态
func (c *Client) GetNodeStats(request *GetNodeStatsRequest) (response *GetNodeStatsResponse, err error) {
	if request == nil {
		request = NewGetNodeStatsRequest()
	}
	response = NewGetNodeStatsResponse()
	err = c.Send(request, response)
	return
}

func NewConfirmHandleErrorDiskCompleteRequest() (request *ConfirmHandleErrorDiskCompleteRequest) {
	request = &ConfirmHandleErrorDiskCompleteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cosyotta", APIVersion, "ConfirmHandleErrorDiskComplete")
	return
}

func NewConfirmHandleErrorDiskCompleteResponse() (response *ConfirmHandleErrorDiskCompleteResponse) {
	response = &ConfirmHandleErrorDiskCompleteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 确认换盘完成
func (c *Client) ConfirmHandleErrorDiskComplete(request *ConfirmHandleErrorDiskCompleteRequest) (response *ConfirmHandleErrorDiskCompleteResponse, err error) {
	if request == nil {
		request = NewConfirmHandleErrorDiskCompleteRequest()
	}
	response = NewConfirmHandleErrorDiskCompleteResponse()
	err = c.Send(request, response)
	return
}
