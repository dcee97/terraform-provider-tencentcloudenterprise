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

package v20210622

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2021-06-22"

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

func NewCreateApmInstanceRequest() (request *CreateApmInstanceRequest) {
	request = &CreateApmInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "CreateApmInstance")
	return
}

func NewCreateApmInstanceResponse() (response *CreateApmInstanceResponse) {
	response = &CreateApmInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务购买APM实例，调用该接口创建
func (c *Client) CreateApmInstance(request *CreateApmInstanceRequest) (response *CreateApmInstanceResponse, err error) {
	if request == nil {
		request = NewCreateApmInstanceRequest()
	}
	response = NewCreateApmInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateApmInstanceRequest() (request *TerminateApmInstanceRequest) {
	request = &TerminateApmInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "TerminateApmInstance")
	return
}

func NewTerminateApmInstanceResponse() (response *TerminateApmInstanceResponse) {
	response = &TerminateApmInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// apm销毁实例
func (c *Client) TerminateApmInstance(request *TerminateApmInstanceRequest) (response *TerminateApmInstanceResponse, err error) {
	if request == nil {
		request = NewTerminateApmInstanceRequest()
	}
	response = NewTerminateApmInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopologyNewRequest() (request *DescribeTopologyNewRequest) {
	request = &DescribeTopologyNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTopologyNew")
	return
}

func NewDescribeTopologyNewResponse() (response *DescribeTopologyNewResponse) {
	response = &DescribeTopologyNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据服务名查询服务拓扑图
func (c *Client) DescribeTopologyNew(request *DescribeTopologyNewRequest) (response *DescribeTopologyNewResponse, err error) {
	if request == nil {
		request = NewDescribeTopologyNewRequest()
	}
	response = NewDescribeTopologyNewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogMonitorDataRequest() (request *DescribeLogMonitorDataRequest) {
	request = &DescribeLogMonitorDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeLogMonitorData")
	return
}

func NewDescribeLogMonitorDataResponse() (response *DescribeLogMonitorDataResponse) {
	response = &DescribeLogMonitorDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志监控上报
func (c *Client) DescribeLogMonitorData(request *DescribeLogMonitorDataRequest) (response *DescribeLogMonitorDataResponse, err error) {
	if request == nil {
		request = NewDescribeLogMonitorDataRequest()
	}
	response = NewDescribeLogMonitorDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmConsumptionRequest() (request *DescribeApmConsumptionRequest) {
	request = &DescribeApmConsumptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmConsumption")
	return
}

func NewDescribeApmConsumptionResponse() (response *DescribeApmConsumptionResponse) {
	response = &DescribeApmConsumptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用量统计
func (c *Client) DescribeApmConsumption(request *DescribeApmConsumptionRequest) (response *DescribeApmConsumptionResponse, err error) {
	if request == nil {
		request = NewDescribeApmConsumptionRequest()
	}
	response = NewDescribeApmConsumptionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmPAASDeduplicateSpanListRequest() (request *DescribeApmPAASDeduplicateSpanListRequest) {
	request = &DescribeApmPAASDeduplicateSpanListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmPAASDeduplicateSpanList")
	return
}

func NewDescribeApmPAASDeduplicateSpanListResponse() (response *DescribeApmPAASDeduplicateSpanListResponse) {
	response = &DescribeApmPAASDeduplicateSpanListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// span去除id重复列表接口
func (c *Client) DescribeApmPAASDeduplicateSpanList(request *DescribeApmPAASDeduplicateSpanListRequest) (response *DescribeApmPAASDeduplicateSpanListResponse, err error) {
	if request == nil {
		request = NewDescribeApmPAASDeduplicateSpanListRequest()
	}
	response = NewDescribeApmPAASDeduplicateSpanListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceBriefsRequest() (request *DescribeInstanceBriefsRequest) {
	request = &DescribeInstanceBriefsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeInstanceBriefs")
	return
}

func NewDescribeInstanceBriefsResponse() (response *DescribeInstanceBriefsResponse) {
	response = &DescribeInstanceBriefsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Apm 实例简介列表
func (c *Client) DescribeInstanceBriefs(request *DescribeInstanceBriefsRequest) (response *DescribeInstanceBriefsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceBriefsRequest()
	}
	response = NewDescribeInstanceBriefsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMetricRecordsRequest() (request *DescribeMetricRecordsRequest) {
	request = &DescribeMetricRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeMetricRecords")
	return
}

func NewDescribeMetricRecordsResponse() (response *DescribeMetricRecordsResponse) {
	response = &DescribeMetricRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取通用指标列表
func (c *Client) DescribeMetricRecords(request *DescribeMetricRecordsRequest) (response *DescribeMetricRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeMetricRecordsRequest()
	}
	response = NewDescribeMetricRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApmApplicationConfigRequest() (request *ModifyApmApplicationConfigRequest) {
	request = &ModifyApmApplicationConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "ModifyApmApplicationConfig")
	return
}

func NewModifyApmApplicationConfigResponse() (response *ModifyApmApplicationConfigResponse) {
	response = &ModifyApmApplicationConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改应用配置接口
func (c *Client) ModifyApmApplicationConfig(request *ModifyApmApplicationConfigRequest) (response *ModifyApmApplicationConfigResponse, err error) {
	if request == nil {
		request = NewModifyApmApplicationConfigRequest()
	}
	response = NewModifyApmApplicationConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTraceListRequest() (request *DescribeTraceListRequest) {
	request = &DescribeTraceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTraceList")
	return
}

func NewDescribeTraceListResponse() (response *DescribeTraceListResponse) {
	response = &DescribeTraceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用示例下的trace分页列表数据
func (c *Client) DescribeTraceList(request *DescribeTraceListRequest) (response *DescribeTraceListResponse, err error) {
	if request == nil {
		request = NewDescribeTraceListRequest()
	}
	response = NewDescribeTraceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmSampleConfigRequest() (request *DescribeApmSampleConfigRequest) {
	request = &DescribeApmSampleConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmSampleConfig")
	return
}

func NewDescribeApmSampleConfigResponse() (response *DescribeApmSampleConfigResponse) {
	response = &DescribeApmSampleConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询采样配置接口
func (c *Client) DescribeApmSampleConfig(request *DescribeApmSampleConfigRequest) (response *DescribeApmSampleConfigResponse, err error) {
	if request == nil {
		request = NewDescribeApmSampleConfigRequest()
	}
	response = NewDescribeApmSampleConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralSpanDetailRequest() (request *DescribeGeneralSpanDetailRequest) {
	request = &DescribeGeneralSpanDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralSpanDetail")
	return
}

func NewDescribeGeneralSpanDetailResponse() (response *DescribeGeneralSpanDetailResponse) {
	response = &DescribeGeneralSpanDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Span明细
func (c *Client) DescribeGeneralSpanDetail(request *DescribeGeneralSpanDetailRequest) (response *DescribeGeneralSpanDetailResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralSpanDetailRequest()
	}
	response = NewDescribeGeneralSpanDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSqlMetricDataRequest() (request *DescribeSqlMetricDataRequest) {
	request = &DescribeSqlMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeSqlMetricData")
	return
}

func NewDescribeSqlMetricDataResponse() (response *DescribeSqlMetricDataResponse) {
	response = &DescribeSqlMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取数据库相关指标接口。用户根据需要上送请求参数，返回对应的指标数据。
// 接口调用频率限制为：20次/秒，1200次/分钟。单请求的数据点数限制为1440个。
func (c *Client) DescribeSqlMetricData(request *DescribeSqlMetricDataRequest) (response *DescribeSqlMetricDataResponse, err error) {
	if request == nil {
		request = NewDescribeSqlMetricDataRequest()
	}
	response = NewDescribeSqlMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopologyRequest() (request *DescribeTopologyRequest) {
	request = &DescribeTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTopology")
	return
}

func NewDescribeTopologyResponse() (response *DescribeTopologyResponse) {
	response = &DescribeTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据服务名查询服务拓扑图
func (c *Client) DescribeTopology(request *DescribeTopologyRequest) (response *DescribeTopologyResponse, err error) {
	if request == nil {
		request = NewDescribeTopologyRequest()
	}
	response = NewDescribeTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMetricLineDataRequest() (request *DescribeMetricLineDataRequest) {
	request = &DescribeMetricLineDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeMetricLineData")
	return
}

func NewDescribeMetricLineDataResponse() (response *DescribeMetricLineDataResponse) {
	response = &DescribeMetricLineDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询不同维度聚合的指标数据，返回曲线图数据
func (c *Client) DescribeMetricLineData(request *DescribeMetricLineDataRequest) (response *DescribeMetricLineDataResponse, err error) {
	if request == nil {
		request = NewDescribeMetricLineDataRequest()
	}
	response = NewDescribeMetricLineDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmInstanceOverviewRequest() (request *DescribeApmInstanceOverviewRequest) {
	request = &DescribeApmInstanceOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmInstanceOverview")
	return
}

func NewDescribeApmInstanceOverviewResponse() (response *DescribeApmInstanceOverviewResponse) {
	response = &DescribeApmInstanceOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取实例相关信息，并返回同环比
func (c *Client) DescribeApmInstanceOverview(request *DescribeApmInstanceOverviewRequest) (response *DescribeApmInstanceOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeApmInstanceOverviewRequest()
	}
	response = NewDescribeApmInstanceOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmInfoByAppIdRequest() (request *DescribeApmInfoByAppIdRequest) {
	request = &DescribeApmInfoByAppIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmInfoByAppId")
	return
}

func NewDescribeApmInfoByAppIdResponse() (response *DescribeApmInfoByAppIdResponse) {
	response = &DescribeApmInfoByAppIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回appId对应的剩余上报额度信息
func (c *Client) DescribeApmInfoByAppId(request *DescribeApmInfoByAppIdRequest) (response *DescribeApmInfoByAppIdResponse, err error) {
	if request == nil {
		request = NewDescribeApmInfoByAppIdRequest()
	}
	response = NewDescribeApmInfoByAppIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
	request = &DescribeServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeService")
	return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
	response = &DescribeServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务实例信息
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
	if request == nil {
		request = NewDescribeServiceRequest()
	}
	response = NewDescribeServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmComboInfoRequest() (request *DescribeApmComboInfoRequest) {
	request = &DescribeApmComboInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmComboInfo")
	return
}

func NewDescribeApmComboInfoResponse() (response *DescribeApmComboInfoResponse) {
	response = &DescribeApmComboInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取预付费套餐包相关信息
func (c *Client) DescribeApmComboInfo(request *DescribeApmComboInfoRequest) (response *DescribeApmComboInfoResponse, err error) {
	if request == nil {
		request = NewDescribeApmComboInfoRequest()
	}
	response = NewDescribeApmComboInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceNodesRequest() (request *DescribeServiceNodesRequest) {
	request = &DescribeServiceNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeServiceNodes")
	return
}

func NewDescribeServiceNodesResponse() (response *DescribeServiceNodesResponse) {
	response = &DescribeServiceNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取实例下所有服务节点列表
func (c *Client) DescribeServiceNodes(request *DescribeServiceNodesRequest) (response *DescribeServiceNodesResponse, err error) {
	if request == nil {
		request = NewDescribeServiceNodesRequest()
	}
	response = NewDescribeServiceNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmAgentRequest() (request *DescribeApmAgentRequest) {
	request = &DescribeApmAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmAgent")
	return
}

func NewDescribeApmAgentResponse() (response *DescribeApmAgentResponse) {
	response = &DescribeApmAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Apm Agent信息
func (c *Client) DescribeApmAgent(request *DescribeApmAgentRequest) (response *DescribeApmAgentResponse, err error) {
	if request == nil {
		request = NewDescribeApmAgentRequest()
	}
	response = NewDescribeApmAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTraceAnalyzeSpanListRequest() (request *DescribeTraceAnalyzeSpanListRequest) {
	request = &DescribeTraceAnalyzeSpanListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTraceAnalyzeSpanList")
	return
}

func NewDescribeTraceAnalyzeSpanListResponse() (response *DescribeTraceAnalyzeSpanListResponse) {
	response = &DescribeTraceAnalyzeSpanListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询trace耗时分析信息
func (c *Client) DescribeTraceAnalyzeSpanList(request *DescribeTraceAnalyzeSpanListRequest) (response *DescribeTraceAnalyzeSpanListResponse, err error) {
	if request == nil {
		request = NewDescribeTraceAnalyzeSpanListRequest()
	}
	response = NewDescribeTraceAnalyzeSpanListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralSpanListRequest() (request *DescribeGeneralSpanListRequest) {
	request = &DescribeGeneralSpanListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralSpanList")
	return
}

func NewDescribeGeneralSpanListResponse() (response *DescribeGeneralSpanListResponse) {
	response = &DescribeGeneralSpanListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通用查询调用链列表
func (c *Client) DescribeGeneralSpanList(request *DescribeGeneralSpanListRequest) (response *DescribeGeneralSpanListResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralSpanListRequest()
	}
	response = NewDescribeGeneralSpanListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApmSampleConfigRequest() (request *CreateApmSampleConfigRequest) {
	request = &CreateApmSampleConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "CreateApmSampleConfig")
	return
}

func NewCreateApmSampleConfigResponse() (response *CreateApmSampleConfigResponse) {
	response = &CreateApmSampleConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建采样配置接口
func (c *Client) CreateApmSampleConfig(request *CreateApmSampleConfigRequest) (response *CreateApmSampleConfigResponse, err error) {
	if request == nil {
		request = NewCreateApmSampleConfigRequest()
	}
	response = NewCreateApmSampleConfigResponse()
	err = c.Send(request, response)
	return
}

func NewOpenApmPaidVersionRequest() (request *OpenApmPaidVersionRequest) {
	request = &OpenApmPaidVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "OpenApmPaidVersion")
	return
}

func NewOpenApmPaidVersionResponse() (response *OpenApmPaidVersionResponse) {
	response = &OpenApmPaidVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开通付费版本
func (c *Client) OpenApmPaidVersion(request *OpenApmPaidVersionRequest) (response *OpenApmPaidVersionResponse, err error) {
	if request == nil {
		request = NewOpenApmPaidVersionRequest()
	}
	response = NewOpenApmPaidVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagCountValuesRequest() (request *DescribeTagCountValuesRequest) {
	request = &DescribeTagCountValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTagCountValues")
	return
}

func NewDescribeTagCountValuesResponse() (response *DescribeTagCountValuesResponse) {
	response = &DescribeTagCountValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据维度名和过滤条件，查询维度分组数据
func (c *Client) DescribeTagCountValues(request *DescribeTagCountValuesRequest) (response *DescribeTagCountValuesResponse, err error) {
	if request == nil {
		request = NewDescribeTagCountValuesRequest()
	}
	response = NewDescribeTagCountValuesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagValuesRequest() (request *DescribeTagValuesRequest) {
	request = &DescribeTagValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTagValues")
	return
}

func NewDescribeTagValuesResponse() (response *DescribeTagValuesResponse) {
	response = &DescribeTagValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据维度名和过滤条件，查询维度数据.
func (c *Client) DescribeTagValues(request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
	if request == nil {
		request = NewDescribeTagValuesRequest()
	}
	response = NewDescribeTagValuesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmApplicationConfigRequest() (request *DescribeApmApplicationConfigRequest) {
	request = &DescribeApmApplicationConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmApplicationConfig")
	return
}

func NewDescribeApmApplicationConfigResponse() (response *DescribeApmApplicationConfigResponse) {
	response = &DescribeApmApplicationConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用配置接口
func (c *Client) DescribeApmApplicationConfig(request *DescribeApmApplicationConfigRequest) (response *DescribeApmApplicationConfigResponse, err error) {
	if request == nil {
		request = NewDescribeApmApplicationConfigRequest()
	}
	response = NewDescribeApmApplicationConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmRegionsRequest() (request *DescribeApmRegionsRequest) {
	request = &DescribeApmRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmRegions")
	return
}

func NewDescribeApmRegionsResponse() (response *DescribeApmRegionsResponse) {
	response = &DescribeApmRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询apm可用地区列表
func (c *Client) DescribeApmRegions(request *DescribeApmRegionsRequest) (response *DescribeApmRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeApmRegionsRequest()
	}
	response = NewDescribeApmRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTopologySnapshotRequest() (request *CreateTopologySnapshotRequest) {
	request = &CreateTopologySnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "CreateTopologySnapshot")
	return
}

func NewCreateTopologySnapshotResponse() (response *CreateTopologySnapshotResponse) {
	response = &CreateTopologySnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存用户的拓扑图信息
func (c *Client) CreateTopologySnapshot(request *CreateTopologySnapshotRequest) (response *CreateTopologySnapshotResponse, err error) {
	if request == nil {
		request = NewCreateTopologySnapshotRequest()
	}
	response = NewCreateTopologySnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMetricPointDataRequest() (request *DescribeMetricPointDataRequest) {
	request = &DescribeMetricPointDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeMetricPointData")
	return
}

func NewDescribeMetricPointDataResponse() (response *DescribeMetricPointDataResponse) {
	response = &DescribeMetricPointDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询不同维度聚合的指标数据，返回单点数据
func (c *Client) DescribeMetricPointData(request *DescribeMetricPointDataRequest) (response *DescribeMetricPointDataResponse, err error) {
	if request == nil {
		request = NewDescribeMetricPointDataRequest()
	}
	response = NewDescribeMetricPointDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpanTagListRequest() (request *DescribeSpanTagListRequest) {
	request = &DescribeSpanTagListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeSpanTagList")
	return
}

func NewDescribeSpanTagListResponse() (response *DescribeSpanTagListResponse) {
	response = &DescribeSpanTagListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 最新异常等Tag相关字段展示接口
func (c *Client) DescribeSpanTagList(request *DescribeSpanTagListRequest) (response *DescribeSpanTagListResponse, err error) {
	if request == nil {
		request = NewDescribeSpanTagListRequest()
	}
	response = NewDescribeSpanTagListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopologyMetricLineDataRequest() (request *DescribeTopologyMetricLineDataRequest) {
	request = &DescribeTopologyMetricLineDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTopologyMetricLineData")
	return
}

func NewDescribeTopologyMetricLineDataResponse() (response *DescribeTopologyMetricLineDataResponse) {
	response = &DescribeTopologyMetricLineDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拓扑图相关折线拉取
func (c *Client) DescribeTopologyMetricLineData(request *DescribeTopologyMetricLineDataRequest) (response *DescribeTopologyMetricLineDataResponse, err error) {
	if request == nil {
		request = NewDescribeTopologyMetricLineDataRequest()
	}
	response = NewDescribeTopologyMetricLineDataResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApmSampleConfigRequest() (request *DeleteApmSampleConfigRequest) {
	request = &DeleteApmSampleConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DeleteApmSampleConfig")
	return
}

func NewDeleteApmSampleConfigResponse() (response *DeleteApmSampleConfigResponse) {
	response = &DeleteApmSampleConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除采样配置接口
func (c *Client) DeleteApmSampleConfig(request *DeleteApmSampleConfigRequest) (response *DeleteApmSampleConfigResponse, err error) {
	if request == nil {
		request = NewDeleteApmSampleConfigRequest()
	}
	response = NewDeleteApmSampleConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApmInstanceRequest() (request *ModifyApmInstanceRequest) {
	request = &ModifyApmInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "ModifyApmInstance")
	return
}

func NewModifyApmInstanceResponse() (response *ModifyApmInstanceResponse) {
	response = &ModifyApmInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Apm实例接口
func (c *Client) ModifyApmInstance(request *ModifyApmInstanceRequest) (response *ModifyApmInstanceResponse, err error) {
	if request == nil {
		request = NewModifyApmInstanceRequest()
	}
	response = NewModifyApmInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceOverviewRequest() (request *DescribeServiceOverviewRequest) {
	request = &DescribeServiceOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeServiceOverview")
	return
}

func NewDescribeServiceOverviewResponse() (response *DescribeServiceOverviewResponse) {
	response = &DescribeServiceOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务概览数据拉取
func (c *Client) DescribeServiceOverview(request *DescribeServiceOverviewRequest) (response *DescribeServiceOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeServiceOverviewRequest()
	}
	response = NewDescribeServiceOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpanTreeByIDRequest() (request *DescribeSpanTreeByIDRequest) {
	request = &DescribeSpanTreeByIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeSpanTreeByID")
	return
}

func NewDescribeSpanTreeByIDResponse() (response *DescribeSpanTreeByIDResponse) {
	response = &DescribeSpanTreeByIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询trace详情接口信息
func (c *Client) DescribeSpanTreeByID(request *DescribeSpanTreeByIDRequest) (response *DescribeSpanTreeByIDResponse, err error) {
	if request == nil {
		request = NewDescribeSpanTreeByIDRequest()
	}
	response = NewDescribeSpanTreeByIDResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmInstancesRequest() (request *DescribeApmInstancesRequest) {
	request = &DescribeApmInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmInstances")
	return
}

func NewDescribeApmInstancesResponse() (response *DescribeApmInstancesResponse) {
	response = &DescribeApmInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// APM实例列表拉取
func (c *Client) DescribeApmInstances(request *DescribeApmInstancesRequest) (response *DescribeApmInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeApmInstancesRequest()
	}
	response = NewDescribeApmInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralMetricDataRequest() (request *DescribeGeneralMetricDataRequest) {
	request = &DescribeGeneralMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralMetricData")
	return
}

func NewDescribeGeneralMetricDataResponse() (response *DescribeGeneralMetricDataResponse) {
	response = &DescribeGeneralMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指标数据通用接口。用户根据需要上送请求参数，返回对应的指标数据。
// 接口调用频率限制为：20次/秒，1200次/分钟。单请求的数据点数限制为1440个。
func (c *Client) DescribeGeneralMetricData(request *DescribeGeneralMetricDataRequest) (response *DescribeGeneralMetricDataResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralMetricDataRequest()
	}
	response = NewDescribeGeneralMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMetricDetailRequest() (request *DescribeMetricDetailRequest) {
	request = &DescribeMetricDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeMetricDetail")
	return
}

func NewDescribeMetricDetailResponse() (response *DescribeMetricDetailResponse) {
	response = &DescribeMetricDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口详情/应用详情等场景使用
func (c *Client) DescribeMetricDetail(request *DescribeMetricDetailRequest) (response *DescribeMetricDetailResponse, err error) {
	if request == nil {
		request = NewDescribeMetricDetailRequest()
	}
	response = NewDescribeMetricDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApmSampleConfigRequest() (request *ModifyApmSampleConfigRequest) {
	request = &ModifyApmSampleConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "ModifyApmSampleConfig")
	return
}

func NewModifyApmSampleConfigResponse() (response *ModifyApmSampleConfigResponse) {
	response = &ModifyApmSampleConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改采样配置接口
func (c *Client) ModifyApmSampleConfig(request *ModifyApmSampleConfigRequest) (response *ModifyApmSampleConfigResponse, err error) {
	if request == nil {
		request = NewModifyApmSampleConfigRequest()
	}
	response = NewModifyApmSampleConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBTagValuesRequest() (request *DescribeDBTagValuesRequest) {
	request = &DescribeDBTagValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeDBTagValues")
	return
}

func NewDescribeDBTagValuesResponse() (response *DescribeDBTagValuesResponse) {
	response = &DescribeDBTagValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据数据库查询标签信息
func (c *Client) DescribeDBTagValues(request *DescribeDBTagValuesRequest) (response *DescribeDBTagValuesResponse, err error) {
	if request == nil {
		request = NewDescribeDBTagValuesRequest()
	}
	response = NewDescribeDBTagValuesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExistInstancesRequest() (request *DescribeExistInstancesRequest) {
	request = &DescribeExistInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeExistInstances")
	return
}

func NewDescribeExistInstancesResponse() (response *DescribeExistInstancesResponse) {
	response = &DescribeExistInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户是否创建过APM实例
func (c *Client) DescribeExistInstances(request *DescribeExistInstancesRequest) (response *DescribeExistInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeExistInstancesRequest()
	}
	response = NewDescribeExistInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePAASTopologyRequest() (request *DescribePAASTopologyRequest) {
	request = &DescribePAASTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribePAASTopology")
	return
}

func NewDescribePAASTopologyResponse() (response *DescribePAASTopologyResponse) {
	response = &DescribePAASTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TAW 通用拓扑接口（PAAS）
func (c *Client) DescribePAASTopology(request *DescribePAASTopologyRequest) (response *DescribePAASTopologyResponse, err error) {
	if request == nil {
		request = NewDescribePAASTopologyRequest()
	}
	response = NewDescribePAASTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceLinkRequest() (request *DescribeServiceLinkRequest) {
	request = &DescribeServiceLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeServiceLink")
	return
}

func NewDescribeServiceLinkResponse() (response *DescribeServiceLinkResponse) {
	response = &DescribeServiceLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过接口获取应用的上游或者下游应用，并且获取它们的对应指标
func (c *Client) DescribeServiceLink(request *DescribeServiceLinkRequest) (response *DescribeServiceLinkResponse, err error) {
	if request == nil {
		request = NewDescribeServiceLinkRequest()
	}
	response = NewDescribeServiceLinkResponse()
	err = c.Send(request, response)
	return
}
