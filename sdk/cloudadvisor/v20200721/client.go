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

package v20200721

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-07-21"

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

func NewDescribeRoleStatusRequest() (request *DescribeRoleStatusRequest) {
	request = &DescribeRoleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeRoleStatus")
	return
}

func NewDescribeRoleStatusResponse() (response *DescribeRoleStatusResponse) {
	response = &DescribeRoleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询角色授权状态
func (c *Client) DescribeRoleStatus(request *DescribeRoleStatusRequest) (response *DescribeRoleStatusResponse, err error) {
	if request == nil {
		request = NewDescribeRoleStatusRequest()
	}
	response = NewDescribeRoleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskResultRequest() (request *DescribeTaskResultRequest) {
	request = &DescribeTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeTaskResult")
	return
}

func NewDescribeTaskResultResponse() (response *DescribeTaskResultResponse) {
	response = &DescribeTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务执行的结果
func (c *Client) DescribeTaskResult(request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeTaskResultRequest()
	}
	response = NewDescribeTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskStrategyIgnoresRequest() (request *DescribeTaskStrategyIgnoresRequest) {
	request = &DescribeTaskStrategyIgnoresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeTaskStrategyIgnores")
	return
}

func NewDescribeTaskStrategyIgnoresResponse() (response *DescribeTaskStrategyIgnoresResponse) {
	response = &DescribeTaskStrategyIgnoresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询评估项忽略实例列表
func (c *Client) DescribeTaskStrategyIgnores(request *DescribeTaskStrategyIgnoresRequest) (response *DescribeTaskStrategyIgnoresResponse, err error) {
	if request == nil {
		request = NewDescribeTaskStrategyIgnoresRequest()
	}
	response = NewDescribeTaskStrategyIgnoresResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadReportFileAsyncRequest() (request *DownloadReportFileAsyncRequest) {
	request = &DownloadReportFileAsyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DownloadReportFileAsync")
	return
}

func NewDownloadReportFileAsyncResponse() (response *DownloadReportFileAsyncResponse) {
	response = &DownloadReportFileAsyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建下载扫描报告任务
func (c *Client) DownloadReportFileAsync(request *DownloadReportFileAsyncRequest) (response *DownloadReportFileAsyncResponse, err error) {
	if request == nil {
		request = NewDownloadReportFileAsyncRequest()
	}
	response = NewDownloadReportFileAsyncResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSubscriptionRequest() (request *UpdateSubscriptionRequest) {
	request = &UpdateSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "UpdateSubscription")
	return
}

func NewUpdateSubscriptionResponse() (response *UpdateSubscriptionResponse) {
	response = &UpdateSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于创建或修改订阅信息。
func (c *Client) UpdateSubscription(request *UpdateSubscriptionRequest) (response *UpdateSubscriptionResponse, err error) {
	if request == nil {
		request = NewUpdateSubscriptionRequest()
	}
	response = NewUpdateSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewListAllIgnoreInstancesRequest() (request *ListAllIgnoreInstancesRequest) {
	request = &ListAllIgnoreInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "ListAllIgnoreInstances")
	return
}

func NewListAllIgnoreInstancesResponse() (response *ListAllIgnoreInstancesResponse) {
	response = &ListAllIgnoreInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取所有的忽略实例的列表
func (c *Client) ListAllIgnoreInstances(request *ListAllIgnoreInstancesRequest) (response *ListAllIgnoreInstancesResponse, err error) {
	if request == nil {
		request = NewListAllIgnoreInstancesRequest()
	}
	response = NewListAllIgnoreInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateScanTaskRequest() (request *CreateScanTaskRequest) {
	request = &CreateScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "CreateScanTask")
	return
}

func NewCreateScanTaskResponse() (response *CreateScanTaskResponse) {
	response = &CreateScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建扫描任务
func (c *Client) CreateScanTask(request *CreateScanTaskRequest) (response *CreateScanTaskResponse, err error) {
	if request == nil {
		request = NewCreateScanTaskRequest()
	}
	response = NewCreateScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadReportFileRequest() (request *DownloadReportFileRequest) {
	request = &DownloadReportFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DownloadReportFile")
	return
}

func NewDownloadReportFileResponse() (response *DownloadReportFileResponse) {
	response = &DownloadReportFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载扫描报告内容
func (c *Client) DownloadReportFile(request *DownloadReportFileRequest) (response *DownloadReportFileResponse, err error) {
	if request == nil {
		request = NewDownloadReportFileRequest()
	}
	response = NewDownloadReportFileResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSubscriptionRequest() (request *DeleteSubscriptionRequest) {
	request = &DeleteSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DeleteSubscription")
	return
}

func NewDeleteSubscriptionResponse() (response *DeleteSubscriptionResponse) {
	response = &DeleteSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除订阅
func (c *Client) DeleteSubscription(request *DeleteSubscriptionRequest) (response *DeleteSubscriptionResponse, err error) {
	if request == nil {
		request = NewDeleteSubscriptionRequest()
	}
	response = NewDeleteSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskProgressRequest() (request *DescribeTaskProgressRequest) {
	request = &DescribeTaskProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeTaskProgress")
	return
}

func NewDescribeTaskProgressResponse() (response *DescribeTaskProgressResponse) {
	response = &DescribeTaskProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务扫描进度
func (c *Client) DescribeTaskProgress(request *DescribeTaskProgressRequest) (response *DescribeTaskProgressResponse, err error) {
	if request == nil {
		request = NewDescribeTaskProgressRequest()
	}
	response = NewDescribeTaskProgressResponse()
	err = c.Send(request, response)
	return
}

func NewListIgnoreInstancesRequest() (request *ListIgnoreInstancesRequest) {
	request = &ListIgnoreInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "ListIgnoreInstances")
	return
}

func NewListIgnoreInstancesResponse() (response *ListIgnoreInstancesResponse) {
	response = &ListIgnoreInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取忽略的实例列表
func (c *Client) ListIgnoreInstances(request *ListIgnoreInstancesRequest) (response *ListIgnoreInstancesResponse, err error) {
	if request == nil {
		request = NewListIgnoreInstancesRequest()
	}
	response = NewListIgnoreInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionsRequest() (request *DescribeSubscriptionsRequest) {
	request = &DescribeSubscriptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeSubscriptions")
	return
}

func NewDescribeSubscriptionsResponse() (response *DescribeSubscriptionsResponse) {
	response = &DescribeSubscriptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前用户的订阅信息
func (c *Client) DescribeSubscriptions(request *DescribeSubscriptionsRequest) (response *DescribeSubscriptionsResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionsRequest()
	}
	response = NewDescribeSubscriptionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGlobalIgnoreTagsRequest() (request *DeleteGlobalIgnoreTagsRequest) {
	request = &DeleteGlobalIgnoreTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DeleteGlobalIgnoreTags")
	return
}

func NewDeleteGlobalIgnoreTagsResponse() (response *DeleteGlobalIgnoreTagsResponse) {
	response = &DeleteGlobalIgnoreTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除全局忽略的Tag
func (c *Client) DeleteGlobalIgnoreTags(request *DeleteGlobalIgnoreTagsRequest) (response *DeleteGlobalIgnoreTagsResponse, err error) {
	if request == nil {
		request = NewDeleteGlobalIgnoreTagsRequest()
	}
	response = NewDeleteGlobalIgnoreTagsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOverviewRequest() (request *DescribeOverviewRequest) {
	request = &DescribeOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeOverview")
	return
}

func NewDescribeOverviewResponse() (response *DescribeOverviewResponse) {
	response = &DescribeOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 概览数据
func (c *Client) DescribeOverview(request *DescribeOverviewRequest) (response *DescribeOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeOverviewRequest()
	}
	response = NewDescribeOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskStrategyIgnoredDetailRequest() (request *DescribeTaskStrategyIgnoredDetailRequest) {
	request = &DescribeTaskStrategyIgnoredDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeTaskStrategyIgnoredDetail")
	return
}

func NewDescribeTaskStrategyIgnoredDetailResponse() (response *DescribeTaskStrategyIgnoredDetailResponse) {
	response = &DescribeTaskStrategyIgnoredDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取忽略实例详情
func (c *Client) DescribeTaskStrategyIgnoredDetail(request *DescribeTaskStrategyIgnoredDetailRequest) (response *DescribeTaskStrategyIgnoredDetailResponse, err error) {
	if request == nil {
		request = NewDescribeTaskStrategyIgnoredDetailRequest()
	}
	response = NewDescribeTaskStrategyIgnoredDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSupportLanguageRequest() (request *DescribeSupportLanguageRequest) {
	request = &DescribeSupportLanguageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeSupportLanguage")
	return
}

func NewDescribeSupportLanguageResponse() (response *DescribeSupportLanguageResponse) {
	response = &DescribeSupportLanguageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持语言列表
func (c *Client) DescribeSupportLanguage(request *DescribeSupportLanguageRequest) (response *DescribeSupportLanguageResponse, err error) {
	if request == nil {
		request = NewDescribeSupportLanguageRequest()
	}
	response = NewDescribeSupportLanguageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
	request = &DescribeTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeTags")
	return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
	response = &DescribeTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品Tag
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
	if request == nil {
		request = NewDescribeTagsRequest()
	}
	response = NewDescribeTagsResponse()
	err = c.Send(request, response)
	return
}

func NewModifySubscriptionTemplateRequest() (request *ModifySubscriptionTemplateRequest) {
	request = &ModifySubscriptionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "ModifySubscriptionTemplate")
	return
}

func NewModifySubscriptionTemplateResponse() (response *ModifySubscriptionTemplateResponse) {
	response = &ModifySubscriptionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改订阅模板
func (c *Client) ModifySubscriptionTemplate(request *ModifySubscriptionTemplateRequest) (response *ModifySubscriptionTemplateResponse, err error) {
	if request == nil {
		request = NewModifySubscriptionTemplateRequest()
	}
	response = NewModifySubscriptionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGlobalIgnoreTagsRequest() (request *DescribeGlobalIgnoreTagsRequest) {
	request = &DescribeGlobalIgnoreTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeGlobalIgnoreTags")
	return
}

func NewDescribeGlobalIgnoreTagsResponse() (response *DescribeGlobalIgnoreTagsResponse) {
	response = &DescribeGlobalIgnoreTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询全局忽略的Tag
func (c *Client) DescribeGlobalIgnoreTags(request *DescribeGlobalIgnoreTagsRequest) (response *DescribeGlobalIgnoreTagsResponse, err error) {
	if request == nil {
		request = NewDescribeGlobalIgnoreTagsRequest()
	}
	response = NewDescribeGlobalIgnoreTagsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductConfigListRequest() (request *DescribeProductConfigListRequest) {
	request = &DescribeProductConfigListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeProductConfigList")
	return
}

func NewDescribeProductConfigListResponse() (response *DescribeProductConfigListResponse) {
	response = &DescribeProductConfigListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品配置列表
func (c *Client) DescribeProductConfigList(request *DescribeProductConfigListRequest) (response *DescribeProductConfigListResponse, err error) {
	if request == nil {
		request = NewDescribeProductConfigListRequest()
	}
	response = NewDescribeProductConfigListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionTemplatesRequest() (request *DescribeSubscriptionTemplatesRequest) {
	request = &DescribeSubscriptionTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeSubscriptionTemplates")
	return
}

func NewDescribeSubscriptionTemplatesResponse() (response *DescribeSubscriptionTemplatesResponse) {
	response = &DescribeSubscriptionTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订阅模板
func (c *Client) DescribeSubscriptionTemplates(request *DescribeSubscriptionTemplatesRequest) (response *DescribeSubscriptionTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionTemplatesRequest()
	}
	response = NewDescribeSubscriptionTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewListRegionCodesRequest() (request *ListRegionCodesRequest) {
	request = &ListRegionCodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "ListRegionCodes")
	return
}

func NewListRegionCodesResponse() (response *ListRegionCodesResponse) {
	response = &ListRegionCodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询地区对应的编码
func (c *Client) ListRegionCodes(request *ListRegionCodesRequest) (response *ListRegionCodesResponse, err error) {
	if request == nil {
		request = NewListRegionCodesRequest()
	}
	response = NewListRegionCodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLastTaskRequest() (request *DescribeLastTaskRequest) {
	request = &DescribeLastTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeLastTask")
	return
}

func NewDescribeLastTaskResponse() (response *DescribeLastTaskResponse) {
	response = &DescribeLastTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取上一个最新的任务的ID
func (c *Client) DescribeLastTask(request *DescribeLastTaskRequest) (response *DescribeLastTaskResponse, err error) {
	if request == nil {
		request = NewDescribeLastTaskRequest()
	}
	response = NewDescribeLastTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStrategiesRequest() (request *DescribeStrategiesRequest) {
	request = &DescribeStrategiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeStrategies")
	return
}

func NewDescribeStrategiesResponse() (response *DescribeStrategiesResponse) {
	response = &DescribeStrategiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询评估项的信息
func (c *Client) DescribeStrategies(request *DescribeStrategiesRequest) (response *DescribeStrategiesResponse, err error) {
	if request == nil {
		request = NewDescribeStrategiesRequest()
	}
	response = NewDescribeStrategiesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSubscriptionTemplateRequest() (request *CreateSubscriptionTemplateRequest) {
	request = &CreateSubscriptionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "CreateSubscriptionTemplate")
	return
}

func NewCreateSubscriptionTemplateResponse() (response *CreateSubscriptionTemplateResponse) {
	response = &CreateSubscriptionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建订阅模板
func (c *Client) CreateSubscriptionTemplate(request *CreateSubscriptionTemplateRequest) (response *CreateSubscriptionTemplateResponse, err error) {
	if request == nil {
		request = NewCreateSubscriptionTemplateRequest()
	}
	response = NewCreateSubscriptionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIgnoreStrategiesRequest() (request *ModifyIgnoreStrategiesRequest) {
	request = &ModifyIgnoreStrategiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "ModifyIgnoreStrategies")
	return
}

func NewModifyIgnoreStrategiesResponse() (response *ModifyIgnoreStrategiesResponse) {
	response = &ModifyIgnoreStrategiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改(删除或者增加)忽略的策略ID
func (c *Client) ModifyIgnoreStrategies(request *ModifyIgnoreStrategiesRequest) (response *ModifyIgnoreStrategiesResponse, err error) {
	if request == nil {
		request = NewModifyIgnoreStrategiesRequest()
	}
	response = NewModifyIgnoreStrategiesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateReportStatusRequest() (request *UpdateReportStatusRequest) {
	request = &UpdateReportStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "UpdateReportStatus")
	return
}

func NewUpdateReportStatusResponse() (response *UpdateReportStatusResponse) {
	response = &UpdateReportStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于更新对架构师的扫描报告解读授权
func (c *Client) UpdateReportStatus(request *UpdateReportStatusRequest) (response *UpdateReportStatusResponse, err error) {
	if request == nil {
		request = NewUpdateReportStatusRequest()
	}
	response = NewUpdateReportStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGlobalIgnoreTagsRequest() (request *ModifyGlobalIgnoreTagsRequest) {
	request = &ModifyGlobalIgnoreTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "ModifyGlobalIgnoreTags")
	return
}

func NewModifyGlobalIgnoreTagsResponse() (response *ModifyGlobalIgnoreTagsResponse) {
	response = &ModifyGlobalIgnoreTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改全局忽略的Tag
func (c *Client) ModifyGlobalIgnoreTags(request *ModifyGlobalIgnoreTagsRequest) (response *ModifyGlobalIgnoreTagsResponse, err error) {
	if request == nil {
		request = NewModifyGlobalIgnoreTagsRequest()
	}
	response = NewModifyGlobalIgnoreTagsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIgnoreInstanceRequest() (request *ModifyIgnoreInstanceRequest) {
	request = &ModifyIgnoreInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "ModifyIgnoreInstance")
	return
}

func NewModifyIgnoreInstanceResponse() (response *ModifyIgnoreInstanceResponse) {
	response = &ModifyIgnoreInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改单个忽略的实例
func (c *Client) ModifyIgnoreInstance(request *ModifyIgnoreInstanceRequest) (response *ModifyIgnoreInstanceResponse, err error) {
	if request == nil {
		request = NewModifyIgnoreInstanceRequest()
	}
	response = NewModifyIgnoreInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskStrategyRisksRequest() (request *DescribeTaskStrategyRisksRequest) {
	request = &DescribeTaskStrategyRisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeTaskStrategyRisks")
	return
}

func NewDescribeTaskStrategyRisksResponse() (response *DescribeTaskStrategyRisksResponse) {
	response = &DescribeTaskStrategyRisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询评估项风险实例列表
func (c *Client) DescribeTaskStrategyRisks(request *DescribeTaskStrategyRisksRequest) (response *DescribeTaskStrategyRisksResponse, err error) {
	if request == nil {
		request = NewDescribeTaskStrategyRisksRequest()
	}
	response = NewDescribeTaskStrategyRisksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIgnoreInstancesRequest() (request *ModifyIgnoreInstancesRequest) {
	request = &ModifyIgnoreInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "ModifyIgnoreInstances")
	return
}

func NewModifyIgnoreInstancesResponse() (response *ModifyIgnoreInstancesResponse) {
	response = &ModifyIgnoreInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改忽略的实例列表
func (c *Client) ModifyIgnoreInstances(request *ModifyIgnoreInstancesRequest) (response *ModifyIgnoreInstancesResponse, err error) {
	if request == nil {
		request = NewModifyIgnoreInstancesRequest()
	}
	response = NewModifyIgnoreInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskSummaryRequest() (request *DescribeTaskSummaryRequest) {
	request = &DescribeTaskSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeTaskSummary")
	return
}

func NewDescribeTaskSummaryResponse() (response *DescribeTaskSummaryResponse) {
	response = &DescribeTaskSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务执行的结果总览
func (c *Client) DescribeTaskSummary(request *DescribeTaskSummaryRequest) (response *DescribeTaskSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeTaskSummaryRequest()
	}
	response = NewDescribeTaskSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSubscriptionTemplateRequest() (request *DeleteSubscriptionTemplateRequest) {
	request = &DeleteSubscriptionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DeleteSubscriptionTemplate")
	return
}

func NewDeleteSubscriptionTemplateResponse() (response *DeleteSubscriptionTemplateResponse) {
	response = &DeleteSubscriptionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除订阅模板
func (c *Client) DeleteSubscriptionTemplate(request *DeleteSubscriptionTemplateRequest) (response *DeleteSubscriptionTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteSubscriptionTemplateRequest()
	}
	response = NewDeleteSubscriptionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRoleStatusRequest() (request *UpdateRoleStatusRequest) {
	request = &UpdateRoleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "UpdateRoleStatus")
	return
}

func NewUpdateRoleStatusResponse() (response *UpdateRoleStatusResponse) {
	response = &UpdateRoleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新角色授权状态
func (c *Client) UpdateRoleStatus(request *UpdateRoleStatusRequest) (response *UpdateRoleStatusResponse, err error) {
	if request == nil {
		request = NewUpdateRoleStatusRequest()
	}
	response = NewUpdateRoleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadTaskRequest() (request *DescribeDownloadTaskRequest) {
	request = &DescribeDownloadTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeDownloadTask")
	return
}

func NewDescribeDownloadTaskResponse() (response *DescribeDownloadTaskResponse) {
	response = &DescribeDownloadTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取扫描报告下载链接
func (c *Client) DescribeDownloadTask(request *DescribeDownloadTaskRequest) (response *DescribeDownloadTaskResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadTaskRequest()
	}
	response = NewDescribeDownloadTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskStrategyUnsafeDetailRequest() (request *DescribeTaskStrategyUnsafeDetailRequest) {
	request = &DescribeTaskStrategyUnsafeDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeTaskStrategyUnsafeDetail")
	return
}

func NewDescribeTaskStrategyUnsafeDetailResponse() (response *DescribeTaskStrategyUnsafeDetailResponse) {
	response = &DescribeTaskStrategyUnsafeDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取有风险实例详情
func (c *Client) DescribeTaskStrategyUnsafeDetail(request *DescribeTaskStrategyUnsafeDetailRequest) (response *DescribeTaskStrategyUnsafeDetailResponse, err error) {
	if request == nil {
		request = NewDescribeTaskStrategyUnsafeDetailRequest()
	}
	response = NewDescribeTaskStrategyUnsafeDetailResponse()
	err = c.Send(request, response)
	return
}

func NewListIgnoreStrategiesRequest() (request *ListIgnoreStrategiesRequest) {
	request = &ListIgnoreStrategiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "ListIgnoreStrategies")
	return
}

func NewListIgnoreStrategiesResponse() (response *ListIgnoreStrategiesResponse) {
	response = &ListIgnoreStrategiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取忽略扫描的策略的ID列表
func (c *Client) ListIgnoreStrategies(request *ListIgnoreStrategiesRequest) (response *ListIgnoreStrategiesResponse, err error) {
	if request == nil {
		request = NewListIgnoreStrategiesRequest()
	}
	response = NewListIgnoreStrategiesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
	request = &DescribeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeConfig")
	return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
	response = &DescribeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述平台的策略配置信息
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
	if request == nil {
		request = NewDescribeConfigRequest()
	}
	response = NewDescribeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReportStatusRequest() (request *DescribeReportStatusRequest) {
	request = &DescribeReportStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeReportStatus")
	return
}

func NewDescribeReportStatusResponse() (response *DescribeReportStatusResponse) {
	response = &DescribeReportStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订阅报告授权状态
func (c *Client) DescribeReportStatus(request *DescribeReportStatusRequest) (response *DescribeReportStatusResponse, err error) {
	if request == nil {
		request = NewDescribeReportStatusRequest()
	}
	response = NewDescribeReportStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDisplayRequest() (request *DescribeRiskDisplayRequest) {
	request = &DescribeRiskDisplayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudadvisor", APIVersion, "DescribeRiskDisplay")
	return
}

func NewDescribeRiskDisplayResponse() (response *DescribeRiskDisplayResponse) {
	response = &DescribeRiskDisplayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询评估项表头显示信息
func (c *Client) DescribeRiskDisplay(request *DescribeRiskDisplayRequest) (response *DescribeRiskDisplayResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDisplayRequest()
	}
	response = NewDescribeRiskDisplayResponse()
	err = c.Send(request, response)
	return
}
