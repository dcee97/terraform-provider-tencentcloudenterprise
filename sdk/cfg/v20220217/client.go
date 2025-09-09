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

package v20220217

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2022-02-17"

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

func NewDescribeCustomActionRequest() (request *DescribeCustomActionRequest) {
	request = &DescribeCustomActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeCustomAction")
	return
}

func NewDescribeCustomActionResponse() (response *DescribeCustomActionResponse) {
	response = &DescribeCustomActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口使用入口：
// 1. 动作库点击自定义动作后根据参数展示动作详情
// 2. 创建演练后查询用户自定义动作用来展示
func (c *Client) DescribeCustomAction(request *DescribeCustomActionRequest) (response *DescribeCustomActionResponse, err error) {
	if request == nil {
		request = NewDescribeCustomActionRequest()
	}
	response = NewDescribeCustomActionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTemplateRequest() (request *DeleteTemplateRequest) {
	request = &DeleteTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DeleteTemplate")
	return
}

func NewDeleteTemplateResponse() (response *DeleteTemplateResponse) {
	response = &DeleteTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除经验库
func (c *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteTemplateRequest()
	}
	response = NewDeleteTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskStatisticsRequest() (request *DescribeTaskStatisticsRequest) {
	request = &DescribeTaskStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeTaskStatistics")
	return
}

func NewDescribeTaskStatisticsResponse() (response *DescribeTaskStatisticsResponse) {
	response = &DescribeTaskStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询演练统计
func (c *Client) DescribeTaskStatistics(request *DescribeTaskStatisticsRequest) (response *DescribeTaskStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskStatisticsRequest()
	}
	response = NewDescribeTaskStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTaskStatusRequest() (request *ModifyTaskStatusRequest) {
	request = &ModifyTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "ModifyTaskStatus")
	return
}

func NewModifyTaskStatusResponse() (response *ModifyTaskStatusResponse) {
	response = &ModifyTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改任务状态
func (c *Client) ModifyTaskStatus(request *ModifyTaskStatusRequest) (response *ModifyTaskStatusResponse, err error) {
	if request == nil {
		request = NewModifyTaskStatusRequest()
	}
	response = NewModifyTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTaskResultRequest() (request *ModifyTaskResultRequest) {
	request = &ModifyTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "ModifyTaskResult")
	return
}

func NewModifyTaskResultResponse() (response *ModifyTaskResultResponse) {
	response = &ModifyTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 记录任务结论
func (c *Client) ModifyTaskResult(request *ModifyTaskResultRequest) (response *ModifyTaskResultResponse, err error) {
	if request == nil {
		request = NewModifyTaskResultRequest()
	}
	response = NewModifyTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewEditTemplateRequest() (request *EditTemplateRequest) {
	request = &EditTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "EditTemplate")
	return
}

func NewEditTemplateResponse() (response *EditTemplateResponse) {
	response = &EditTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑经验库
func (c *Client) EditTemplate(request *EditTemplateRequest) (response *EditTemplateResponse, err error) {
	if request == nil {
		request = NewEditTemplateRequest()
	}
	response = NewEditTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeTask")
	return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
	response = &DescribeTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	response = NewDescribeTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeActionFieldConfigListRequest() (request *DescribeActionFieldConfigListRequest) {
	request = &DescribeActionFieldConfigListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeActionFieldConfigList")
	return
}

func NewDescribeActionFieldConfigListResponse() (response *DescribeActionFieldConfigListResponse) {
	response = &DescribeActionFieldConfigListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据动作ID获取动作栏位动态配置参数信息，里面包含动作自有和通用两部分参数。
func (c *Client) DescribeActionFieldConfigList(request *DescribeActionFieldConfigListRequest) (response *DescribeActionFieldConfigListResponse, err error) {
	if request == nil {
		request = NewDescribeActionFieldConfigListRequest()
	}
	response = NewDescribeActionFieldConfigListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeObjectTypeListRequest() (request *DescribeObjectTypeListRequest) {
	request = &DescribeObjectTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeObjectTypeList")
	return
}

func NewDescribeObjectTypeListResponse() (response *DescribeObjectTypeListResponse) {
	response = &DescribeObjectTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询对象类型列表
func (c *Client) DescribeObjectTypeList(request *DescribeObjectTypeListRequest) (response *DescribeObjectTypeListResponse, err error) {
	if request == nil {
		request = NewDescribeObjectTypeListRequest()
	}
	response = NewDescribeObjectTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplateRequest() (request *DescribeTemplateRequest) {
	request = &DescribeTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeTemplate")
	return
}

func NewDescribeTemplateResponse() (response *DescribeTemplateResponse) {
	response = &DescribeTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询经验库
func (c *Client) DescribeTemplate(request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeTemplateRequest()
	}
	response = NewDescribeTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeObjectTypeSearchListRequest() (request *DescribeObjectTypeSearchListRequest) {
	request = &DescribeObjectTypeSearchListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeObjectTypeSearchList")
	return
}

func NewDescribeObjectTypeSearchListResponse() (response *DescribeObjectTypeSearchListResponse) {
	response = &DescribeObjectTypeSearchListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示对象类型搜索键的列表
func (c *Client) DescribeObjectTypeSearchList(request *DescribeObjectTypeSearchListRequest) (response *DescribeObjectTypeSearchListResponse, err error) {
	if request == nil {
		request = NewDescribeObjectTypeSearchListRequest()
	}
	response = NewDescribeObjectTypeSearchListResponse()
	err = c.Send(request, response)
	return
}

func NewEditTaskRequest() (request *EditTaskRequest) {
	request = &EditTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "EditTask")
	return
}

func NewEditTaskResponse() (response *EditTaskResponse) {
	response = &EditTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑任务
func (c *Client) EditTask(request *EditTaskRequest) (response *EditTaskResponse, err error) {
	if request == nil {
		request = NewEditTaskRequest()
	}
	response = NewEditTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCamIdentityRequest() (request *DescribeCamIdentityRequest) {
	request = &DescribeCamIdentityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeCamIdentity")
	return
}

func NewDescribeCamIdentityResponse() (response *DescribeCamIdentityResponse) {
	response = &DescribeCamIdentityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户CAM授权信息
func (c *Client) DescribeCamIdentity(request *DescribeCamIdentityRequest) (response *DescribeCamIdentityResponse, err error) {
	if request == nil {
		request = NewDescribeCamIdentityRequest()
	}
	response = NewDescribeCamIdentityResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteTaskInstanceRequest() (request *ExecuteTaskInstanceRequest) {
	request = &ExecuteTaskInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "ExecuteTaskInstance")
	return
}

func NewExecuteTaskInstanceResponse() (response *ExecuteTaskInstanceResponse) {
	response = &ExecuteTaskInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 触发混沌演练任务的动作，对于实例进行演练操作
func (c *Client) ExecuteTaskInstance(request *ExecuteTaskInstanceRequest) (response *ExecuteTaskInstanceResponse, err error) {
	if request == nil {
		request = NewExecuteTaskInstanceRequest()
	}
	response = NewExecuteTaskInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDummyRequest() (request *DescribeDummyRequest) {
	request = &DescribeDummyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeDummy")
	return
}

func NewDescribeDummyResponse() (response *DescribeDummyResponse) {
	response = &DescribeDummyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 空接口，用于测试
func (c *Client) DescribeDummy(request *DescribeDummyRequest) (response *DescribeDummyResponse, err error) {
	if request == nil {
		request = NewDescribeDummyRequest()
	}
	response = NewDescribeDummyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplateListRequest() (request *DescribeTemplateListRequest) {
	request = &DescribeTemplateListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeTemplateList")
	return
}

func NewDescribeTemplateListResponse() (response *DescribeTemplateListResponse) {
	response = &DescribeTemplateListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询经验库列表
func (c *Client) DescribeTemplateList(request *DescribeTemplateListRequest) (response *DescribeTemplateListResponse, err error) {
	if request == nil {
		request = NewDescribeTemplateListRequest()
	}
	response = NewDescribeTemplateListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTemplateRequest() (request *CreateTemplateRequest) {
	request = &CreateTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "CreateTemplate")
	return
}

func NewCreateTemplateResponse() (response *CreateTemplateResponse) {
	response = &CreateTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建经验库
func (c *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
	if request == nil {
		request = NewCreateTemplateRequest()
	}
	response = NewCreateTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
	request = &CreateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "CreateTask")
	return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
	response = &CreateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建任务
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	response = NewCreateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCustomActionRequest() (request *UpdateCustomActionRequest) {
	request = &UpdateCustomActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "UpdateCustomAction")
	return
}

func NewUpdateCustomActionResponse() (response *UpdateCustomActionResponse) {
	response = &UpdateCustomActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口使用入口：
// 1.动作库编辑自定义动作
// 2.创建演练后对已保存到动作库中的动作进行再次编辑保存
func (c *Client) UpdateCustomAction(request *UpdateCustomActionRequest) (response *UpdateCustomActionResponse, err error) {
	if request == nil {
		request = NewUpdateCustomActionRequest()
	}
	response = NewUpdateCustomActionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeActionLibraryListRequest() (request *DescribeActionLibraryListRequest) {
	request = &DescribeActionLibraryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeActionLibraryList")
	return
}

func NewDescribeActionLibraryListResponse() (response *DescribeActionLibraryListResponse) {
	response = &DescribeActionLibraryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取混沌演练平台的动作库列表
func (c *Client) DescribeActionLibraryList(request *DescribeActionLibraryListRequest) (response *DescribeActionLibraryListResponse, err error) {
	if request == nil {
		request = NewDescribeActionLibraryListRequest()
	}
	response = NewDescribeActionLibraryListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskLogContentRequest() (request *DescribeTaskLogContentRequest) {
	request = &DescribeTaskLogContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeTaskLogContent")
	return
}

func NewDescribeTaskLogContentResponse() (response *DescribeTaskLogContentResponse) {
	response = &DescribeTaskLogContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志下载的内容
func (c *Client) DescribeTaskLogContent(request *DescribeTaskLogContentRequest) (response *DescribeTaskLogContentResponse, err error) {
	if request == nil {
		request = NewDescribeTaskLogContentRequest()
	}
	response = NewDescribeTaskLogContentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTemplateIsUsedRequest() (request *ModifyTemplateIsUsedRequest) {
	request = &ModifyTemplateIsUsedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "ModifyTemplateIsUsed")
	return
}

func NewModifyTemplateIsUsedResponse() (response *ModifyTemplateIsUsedResponse) {
	response = &ModifyTemplateIsUsedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改经验库使用状态
func (c *Client) ModifyTemplateIsUsed(request *ModifyTemplateIsUsedRequest) (response *ModifyTemplateIsUsedResponse, err error) {
	if request == nil {
		request = NewModifyTemplateIsUsedRequest()
	}
	response = NewModifyTemplateIsUsedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionListRequest() (request *DescribeRegionListRequest) {
	request = &DescribeRegionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeRegionList")
	return
}

func NewDescribeRegionListResponse() (response *DescribeRegionListResponse) {
	response = &DescribeRegionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询地域列表
func (c *Client) DescribeRegionList(request *DescribeRegionListRequest) (response *DescribeRegionListResponse, err error) {
	if request == nil {
		request = NewDescribeRegionListRequest()
	}
	response = NewDescribeRegionListResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteTaskRequest() (request *ExecuteTaskRequest) {
	request = &ExecuteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "ExecuteTask")
	return
}

func NewExecuteTaskResponse() (response *ExecuteTaskResponse) {
	response = &ExecuteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行任务
func (c *Client) ExecuteTask(request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
	if request == nil {
		request = NewExecuteTaskRequest()
	}
	response = NewExecuteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomActionRequest() (request *DeleteCustomActionRequest) {
	request = &DeleteCustomActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DeleteCustomAction")
	return
}

func NewDeleteCustomActionResponse() (response *DeleteCustomActionResponse) {
	response = &DeleteCustomActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口使用入口：
// 1. 动作库删除自定义动作
// 2. 创建演练时将自定义动作保存到动作库后可以选择移除
func (c *Client) DeleteCustomAction(request *DeleteCustomActionRequest) (response *DeleteCustomActionResponse, err error) {
	if request == nil {
		request = NewDeleteCustomActionRequest()
	}
	response = NewDeleteCustomActionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNoticeIdRequest() (request *DescribeNoticeIdRequest) {
	request = &DescribeNoticeIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeNoticeId")
	return
}

func NewDescribeNoticeIdResponse() (response *DescribeNoticeIdResponse) {
	response = &DescribeNoticeIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户通知模板ID
func (c *Client) DescribeNoticeId(request *DescribeNoticeIdRequest) (response *DescribeNoticeIdResponse, err error) {
	if request == nil {
		request = NewDescribeNoticeIdRequest()
	}
	response = NewDescribeNoticeIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskListRequest() (request *DescribeTaskListRequest) {
	request = &DescribeTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeTaskList")
	return
}

func NewDescribeTaskListResponse() (response *DescribeTaskListResponse) {
	response = &DescribeTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务列表
func (c *Client) DescribeTaskList(request *DescribeTaskListRequest) (response *DescribeTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeTaskListRequest()
	}
	response = NewDescribeTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeObjectMetricsRequest() (request *DescribeObjectMetricsRequest) {
	request = &DescribeObjectMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeObjectMetrics")
	return
}

func NewDescribeObjectMetricsResponse() (response *DescribeObjectMetricsResponse) {
	response = &DescribeObjectMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据对象ID获取对象类型指标信息
func (c *Client) DescribeObjectMetrics(request *DescribeObjectMetricsRequest) (response *DescribeObjectMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeObjectMetricsRequest()
	}
	response = NewDescribeObjectMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskExecuteLogsRequest() (request *DescribeTaskExecuteLogsRequest) {
	request = &DescribeTaskExecuteLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeTaskExecuteLogs")
	return
}

func NewDescribeTaskExecuteLogsResponse() (response *DescribeTaskExecuteLogsResponse) {
	response = &DescribeTaskExecuteLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取演练过程中的所有日志
func (c *Client) DescribeTaskExecuteLogs(request *DescribeTaskExecuteLogsRequest) (response *DescribeTaskExecuteLogsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskExecuteLogsRequest()
	}
	response = NewDescribeTaskExecuteLogsResponse()
	err = c.Send(request, response)
	return
}

func NewCheckResourceByRoleNameRequest() (request *CheckResourceByRoleNameRequest) {
	request = &CheckResourceByRoleNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "CheckResourceByRoleName")
	return
}

func NewCheckResourceByRoleNameResponse() (response *CheckResourceByRoleNameResponse) {
	response = &CheckResourceByRoleNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务相关角色使用用户资源接口
func (c *Client) CheckResourceByRoleName(request *CheckResourceByRoleNameRequest) (response *CheckResourceByRoleNameResponse, err error) {
	if request == nil {
		request = NewCheckResourceByRoleNameRequest()
	}
	response = NewCheckResourceByRoleNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskLogListRequest() (request *DescribeTaskLogListRequest) {
	request = &DescribeTaskLogListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeTaskLogList")
	return
}

func NewDescribeTaskLogListResponse() (response *DescribeTaskLogListResponse) {
	response = &DescribeTaskLogListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志下载的演练名称列表
func (c *Client) DescribeTaskLogList(request *DescribeTaskLogListRequest) (response *DescribeTaskLogListResponse, err error) {
	if request == nil {
		request = NewDescribeTaskLogListRequest()
	}
	response = NewDescribeTaskLogListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskStatisticsOperateConditionRequest() (request *DescribeTaskStatisticsOperateConditionRequest) {
	request = &DescribeTaskStatisticsOperateConditionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DescribeTaskStatisticsOperateCondition")
	return
}

func NewDescribeTaskStatisticsOperateConditionResponse() (response *DescribeTaskStatisticsOperateConditionResponse) {
	response = &DescribeTaskStatisticsOperateConditionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据用户的运行情况进行统计
func (c *Client) DescribeTaskStatisticsOperateCondition(request *DescribeTaskStatisticsOperateConditionRequest) (response *DescribeTaskStatisticsOperateConditionResponse, err error) {
	if request == nil {
		request = NewDescribeTaskStatisticsOperateConditionRequest()
	}
	response = NewDescribeTaskStatisticsOperateConditionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCustomActionRequest() (request *CreateCustomActionRequest) {
	request = &CreateCustomActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "CreateCustomAction")
	return
}

func NewCreateCustomActionResponse() (response *CreateCustomActionResponse) {
	response = &CreateCustomActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建自定义动作，调用入口：
// 1. 动作库 新增自定义动作
// 2. 创建演练时将某次自定义动作保存到动作库
func (c *Client) CreateCustomAction(request *CreateCustomActionRequest) (response *CreateCustomActionResponse, err error) {
	if request == nil {
		request = NewCreateCustomActionRequest()
	}
	response = NewCreateCustomActionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
	request = &DeleteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfg", APIVersion, "DeleteTask")
	return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
	response = &DeleteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除任务
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
	if request == nil {
		request = NewDeleteTaskRequest()
	}
	response = NewDeleteTaskResponse()
	err = c.Send(request, response)
	return
}
