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

package v20201016

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-10-16"

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

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
	request = &DescribeMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeMachines")
	return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
	response = &DescribeMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取制定机器组下的机器状态
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeMachinesRequest()
	}
	response = NewDescribeMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewMergePartitionRequest() (request *MergePartitionRequest) {
	request = &MergePartitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "MergePartition")
	return
}

func NewMergePartitionResponse() (response *MergePartitionResponse) {
	response = &MergePartitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于合并一个读写态的主题分区，合并时指定一个主题分区 ID，日志服务会自动合并范围右相邻的分区。
func (c *Client) MergePartition(request *MergePartitionRequest) (response *MergePartitionResponse, err error) {
	if request == nil {
		request = NewMergePartitionRequest()
	}
	response = NewMergePartitionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDashboardRequest() (request *DeleteDashboardRequest) {
	request = &DeleteDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteDashboard")
	return
}

func NewDeleteDashboardResponse() (response *DeleteDashboardResponse) {
	response = &DeleteDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除仪表盘
func (c *Client) DeleteDashboard(request *DeleteDashboardRequest) (response *DeleteDashboardResponse, err error) {
	if request == nil {
		request = NewDeleteDashboardRequest()
	}
	response = NewDeleteDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmsRequest() (request *DescribeAlarmsRequest) {
	request = &DescribeAlarmsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarms")
	return
}

func NewDescribeAlarmsResponse() (response *DescribeAlarmsResponse) {
	response = &DescribeAlarmsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取告警策略。
func (c *Client) DescribeAlarms(request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmsRequest()
	}
	response = NewDescribeAlarmsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMachineGroupRequest() (request *CreateMachineGroupRequest) {
	request = &CreateMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateMachineGroup")
	return
}

func NewCreateMachineGroupResponse() (response *CreateMachineGroupResponse) {
	response = &CreateMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建机器组
func (c *Client) CreateMachineGroup(request *CreateMachineGroupRequest) (response *CreateMachineGroupResponse, err error) {
	if request == nil {
		request = NewCreateMachineGroupRequest()
	}
	response = NewCreateMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserConfigRequest() (request *DeleteUserConfigRequest) {
	request = &DeleteUserConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteUserConfig")
	return
}

func NewDeleteUserConfigResponse() (response *DeleteUserConfigResponse) {
	response = &DeleteUserConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户配置
func (c *Client) DeleteUserConfig(request *DeleteUserConfigRequest) (response *DeleteUserConfigResponse, err error) {
	if request == nil {
		request = NewDeleteUserConfigRequest()
	}
	response = NewDeleteUserConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFunctionsRequest() (request *DescribeFunctionsRequest) {
	request = &DescribeFunctionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeFunctions")
	return
}

func NewDescribeFunctionsResponse() (response *DescribeFunctionsResponse) {
	response = &DescribeFunctionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取可使用的函数列表。
func (c *Client) DescribeFunctions(request *DescribeFunctionsRequest) (response *DescribeFunctionsResponse, err error) {
	if request == nil {
		request = NewDescribeFunctionsRequest()
	}
	response = NewDescribeFunctionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogsetsRequest() (request *DescribeLogsetsRequest) {
	request = &DescribeLogsetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLogsets")
	return
}

func NewDescribeLogsetsResponse() (response *DescribeLogsetsResponse) {
	response = &DescribeLogsetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取日志集信息列表。
func (c *Client) DescribeLogsets(request *DescribeLogsetsRequest) (response *DescribeLogsetsResponse, err error) {
	if request == nil {
		request = NewDescribeLogsetsRequest()
	}
	response = NewDescribeLogsetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRebuildIndexTasksRequest() (request *DescribeRebuildIndexTasksRequest) {
	request = &DescribeRebuildIndexTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeRebuildIndexTasks")
	return
}

func NewDescribeRebuildIndexTasksResponse() (response *DescribeRebuildIndexTasksResponse) {
	response = &DescribeRebuildIndexTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取重建索引任务列表
func (c *Client) DescribeRebuildIndexTasks(request *DescribeRebuildIndexTasksRequest) (response *DescribeRebuildIndexTasksResponse, err error) {
	if request == nil {
		request = NewDescribeRebuildIndexTasksRequest()
	}
	response = NewDescribeRebuildIndexTasksResponse()
	err = c.Send(request, response)
	return
}

func NewOpenKafkaConsumerRequest() (request *OpenKafkaConsumerRequest) {
	request = &OpenKafkaConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "OpenKafkaConsumer")
	return
}

func NewOpenKafkaConsumerResponse() (response *OpenKafkaConsumerResponse) {
	response = &OpenKafkaConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 打开Kafka协议消费功能
func (c *Client) OpenKafkaConsumer(request *OpenKafkaConsumerRequest) (response *OpenKafkaConsumerResponse, err error) {
	if request == nil {
		request = NewOpenKafkaConsumerRequest()
	}
	response = NewOpenKafkaConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigFromMachineGroupRequest() (request *DeleteConfigFromMachineGroupRequest) {
	request = &DeleteConfigFromMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConfigFromMachineGroup")
	return
}

func NewDeleteConfigFromMachineGroupResponse() (response *DeleteConfigFromMachineGroupResponse) {
	response = &DeleteConfigFromMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用到机器组的采集配置
func (c *Client) DeleteConfigFromMachineGroup(request *DeleteConfigFromMachineGroupRequest) (response *DeleteConfigFromMachineGroupResponse, err error) {
	if request == nil {
		request = NewDeleteConfigFromMachineGroupRequest()
	}
	response = NewDeleteConfigFromMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndexsRequest() (request *DescribeIndexsRequest) {
	request = &DescribeIndexsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeIndexs")
	return
}

func NewDescribeIndexsResponse() (response *DescribeIndexsResponse) {
	response = &DescribeIndexsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取索引配置列表
func (c *Client) DescribeIndexs(request *DescribeIndexsRequest) (response *DescribeIndexsResponse, err error) {
	if request == nil {
		request = NewDescribeIndexsRequest()
	}
	response = NewDescribeIndexsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCosRechargesRequest() (request *DescribeCosRechargesRequest) {
	request = &DescribeCosRechargesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeCosRecharges")
	return
}

func NewDescribeCosRechargesResponse() (response *DescribeCosRechargesResponse) {
	response = &DescribeCosRechargesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取cos导入配置
func (c *Client) DescribeCosRecharges(request *DescribeCosRechargesRequest) (response *DescribeCosRechargesResponse, err error) {
	if request == nil {
		request = NewDescribeCosRechargesRequest()
	}
	response = NewDescribeCosRechargesResponse()
	err = c.Send(request, response)
	return
}

func NewGenKVRegexRequest() (request *GenKVRegexRequest) {
	request = &GenKVRegexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GenKVRegex")
	return
}

func NewGenKVRegexResponse() (response *GenKVRegexResponse) {
	response = &GenKVRegexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成提取K-V形式的正则表达式
func (c *Client) GenKVRegex(request *GenKVRegexRequest) (response *GenKVRegexResponse, err error) {
	if request == nil {
		request = NewGenKVRegexRequest()
	}
	response = NewGenKVRegexResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveMachineRequest() (request *RemoveMachineRequest) {
	request = &RemoveMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "RemoveMachine")
	return
}

func NewRemoveMachineResponse() (response *RemoveMachineResponse) {
	response = &RemoveMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于剔除标签机器组中机器
func (c *Client) RemoveMachine(request *RemoveMachineRequest) (response *RemoveMachineResponse, err error) {
	if request == nil {
		request = NewRemoveMachineRequest()
	}
	response = NewRemoveMachineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAsyncContextTaskRequest() (request *CreateAsyncContextTaskRequest) {
	request = &CreateAsyncContextTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateAsyncContextTask")
	return
}

func NewCreateAsyncContextTaskResponse() (response *CreateAsyncContextTaskResponse) {
	response = &CreateAsyncContextTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建离线上下文任务
func (c *Client) CreateAsyncContextTask(request *CreateAsyncContextTaskRequest) (response *CreateAsyncContextTaskResponse, err error) {
	if request == nil {
		request = NewCreateAsyncContextTaskRequest()
	}
	response = NewCreateAsyncContextTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaConsumerRequest() (request *DescribeKafkaConsumerRequest) {
	request = &DescribeKafkaConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumer")
	return
}

func NewDescribeKafkaConsumerResponse() (response *DescribeKafkaConsumerResponse) {
	response = &DescribeKafkaConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Kafka协议消费信息
func (c *Client) DescribeKafkaConsumer(request *DescribeKafkaConsumerRequest) (response *DescribeKafkaConsumerResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaConsumerRequest()
	}
	response = NewDescribeKafkaConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewSplitPartitionRequest() (request *SplitPartitionRequest) {
	request = &SplitPartitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "SplitPartition")
	return
}

func NewSplitPartitionResponse() (response *SplitPartitionResponse) {
	response = &SplitPartitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于分裂主题分区
func (c *Client) SplitPartition(request *SplitPartitionRequest) (response *SplitPartitionResponse, err error) {
	if request == nil {
		request = NewSplitPartitionRequest()
	}
	response = NewSplitPartitionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDemonstrationRequest() (request *DeleteDemonstrationRequest) {
	request = &DeleteDemonstrationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteDemonstration")
	return
}

func NewDeleteDemonstrationResponse() (response *DeleteDemonstrationResponse) {
	response = &DeleteDemonstrationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除演示示例
func (c *Client) DeleteDemonstration(request *DeleteDemonstrationRequest) (response *DeleteDemonstrationResponse, err error) {
	if request == nil {
		request = NewDeleteDemonstrationRequest()
	}
	response = NewDeleteDemonstrationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountRequest() (request *DescribeAccountRequest) {
	request = &DescribeAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAccount")
	return
}

func NewDescribeAccountResponse() (response *DescribeAccountResponse) {
	response = &DescribeAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账户状态
func (c *Client) DescribeAccount(request *DescribeAccountRequest) (response *DescribeAccountResponse, err error) {
	if request == nil {
		request = NewDescribeAccountRequest()
	}
	response = NewDescribeAccountResponse()
	err = c.Send(request, response)
	return
}

func NewGenBeginRegexRequest() (request *GenBeginRegexRequest) {
	request = &GenBeginRegexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GenBeginRegex")
	return
}

func NewGenBeginRegexResponse() (response *GenBeginRegexResponse) {
	response = &GenBeginRegexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成首行正则表达式
func (c *Client) GenBeginRegex(request *GenBeginRegexRequest) (response *GenBeginRegexResponse, err error) {
	if request == nil {
		request = NewGenBeginRegexRequest()
	}
	response = NewGenBeginRegexResponse()
	err = c.Send(request, response)
	return
}

func NewCancelRebuildIndexTaskRequest() (request *CancelRebuildIndexTaskRequest) {
	request = &CancelRebuildIndexTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CancelRebuildIndexTask")
	return
}

func NewCancelRebuildIndexTaskResponse() (response *CancelRebuildIndexTaskResponse) {
	response = &CancelRebuildIndexTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消重建索引任务
func (c *Client) CancelRebuildIndexTask(request *CancelRebuildIndexTaskRequest) (response *CancelRebuildIndexTaskResponse, err error) {
	if request == nil {
		request = NewCancelRebuildIndexTaskRequest()
	}
	response = NewCancelRebuildIndexTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
	request = &CreateTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateTopic")
	return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
	response = &CreateTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建日志主题。
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	if request == nil {
		request = NewCreateTopicRequest()
	}
	response = NewCreateTopicResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeAgentNormalRequest() (request *UpgradeAgentNormalRequest) {
	request = &UpgradeAgentNormalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "UpgradeAgentNormal")
	return
}

func NewUpgradeAgentNormalResponse() (response *UpgradeAgentNormalResponse) {
	response = &UpgradeAgentNormalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 基于Agent粒度的单次升级任务启动
func (c *Client) UpgradeAgentNormal(request *UpgradeAgentNormalRequest) (response *UpgradeAgentNormalResponse, err error) {
	if request == nil {
		request = NewUpgradeAgentNormalRequest()
	}
	response = NewUpgradeAgentNormalResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAsyncContextTasksRequest() (request *DescribeAsyncContextTasksRequest) {
	request = &DescribeAsyncContextTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAsyncContextTasks")
	return
}

func NewDescribeAsyncContextTasksResponse() (response *DescribeAsyncContextTasksResponse) {
	response = &DescribeAsyncContextTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取离线上下文任务列表
func (c *Client) DescribeAsyncContextTasks(request *DescribeAsyncContextTasksRequest) (response *DescribeAsyncContextTasksResponse, err error) {
	if request == nil {
		request = NewDescribeAsyncContextTasksRequest()
	}
	response = NewDescribeAsyncContextTasksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConfigRequest() (request *ModifyConfigRequest) {
	request = &ModifyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyConfig")
	return
}

func NewModifyConfigResponse() (response *ModifyConfigResponse) {
	response = &ModifyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改采集规则配置
func (c *Client) ModifyConfig(request *ModifyConfigRequest) (response *ModifyConfigResponse, err error) {
	if request == nil {
		request = NewModifyConfigRequest()
	}
	response = NewModifyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDemonstrationResourcesRequest() (request *ModifyDemonstrationResourcesRequest) {
	request = &ModifyDemonstrationResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyDemonstrationResources")
	return
}

func NewModifyDemonstrationResourcesResponse() (response *ModifyDemonstrationResourcesResponse) {
	response = &ModifyDemonstrationResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于编辑演示示例资源
func (c *Client) ModifyDemonstrationResources(request *ModifyDemonstrationResourcesRequest) (response *ModifyDemonstrationResourcesResponse, err error) {
	if request == nil {
		request = NewModifyDemonstrationResourcesRequest()
	}
	response = NewModifyDemonstrationResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsumerRequest() (request *DescribeConsumerRequest) {
	request = &DescribeConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumer")
	return
}

func NewDescribeConsumerResponse() (response *DescribeConsumerResponse) {
	response = &DescribeConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取投递配置
func (c *Client) DescribeConsumer(request *DescribeConsumerRequest) (response *DescribeConsumerResponse, err error) {
	if request == nil {
		request = NewDescribeConsumerRequest()
	}
	response = NewDescribeConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaConsumeRequest() (request *DescribeKafkaConsumeRequest) {
	request = &DescribeKafkaConsumeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsume")
	return
}

func NewDescribeKafkaConsumeResponse() (response *DescribeKafkaConsumeResponse) {
	response = &DescribeKafkaConsumeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Kakfa消费信息
func (c *Client) DescribeKafkaConsume(request *DescribeKafkaConsumeRequest) (response *DescribeKafkaConsumeResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaConsumeRequest()
	}
	response = NewDescribeKafkaConsumeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIndexRequest() (request *CreateIndexRequest) {
	request = &CreateIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateIndex")
	return
}

func NewCreateIndexResponse() (response *CreateIndexResponse) {
	response = &CreateIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建索引
func (c *Client) CreateIndex(request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
	if request == nil {
		request = NewCreateIndexRequest()
	}
	response = NewCreateIndexResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAsyncContextTaskRequest() (request *DeleteAsyncContextTaskRequest) {
	request = &DeleteAsyncContextTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteAsyncContextTask")
	return
}

func NewDeleteAsyncContextTaskResponse() (response *DeleteAsyncContextTaskResponse) {
	response = &DeleteAsyncContextTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除离线上下文任务
func (c *Client) DeleteAsyncContextTask(request *DeleteAsyncContextTaskRequest) (response *DeleteAsyncContextTaskResponse, err error) {
	if request == nil {
		request = NewDeleteAsyncContextTaskRequest()
	}
	response = NewDeleteAsyncContextTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeliverCloudFunctionRequest() (request *DescribeDeliverCloudFunctionRequest) {
	request = &DescribeDeliverCloudFunctionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDeliverCloudFunction")
	return
}

func NewDescribeDeliverCloudFunctionResponse() (response *DescribeDeliverCloudFunctionResponse) {
	response = &DescribeDeliverCloudFunctionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取投递SCF任务
func (c *Client) DescribeDeliverCloudFunction(request *DescribeDeliverCloudFunctionRequest) (response *DescribeDeliverCloudFunctionResponse, err error) {
	if request == nil {
		request = NewDescribeDeliverCloudFunctionRequest()
	}
	response = NewDescribeDeliverCloudFunctionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigExtraRequest() (request *CreateConfigExtraRequest) {
	request = &CreateConfigExtraRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateConfigExtra")
	return
}

func NewCreateConfigExtraResponse() (response *CreateConfigExtraResponse) {
	response = &CreateConfigExtraResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建特殊采集配置任务
func (c *Client) CreateConfigExtra(request *CreateConfigExtraRequest) (response *CreateConfigExtraResponse, err error) {
	if request == nil {
		request = NewCreateConfigExtraRequest()
	}
	response = NewCreateConfigExtraResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDashboardRequest() (request *ModifyDashboardRequest) {
	request = &ModifyDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyDashboard")
	return
}

func NewModifyDashboardResponse() (response *ModifyDashboardResponse) {
	response = &ModifyDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改仪表盘
func (c *Client) ModifyDashboard(request *ModifyDashboardRequest) (response *ModifyDashboardResponse, err error) {
	if request == nil {
		request = NewModifyDashboardRequest()
	}
	response = NewModifyDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDashboardRequest() (request *CreateDashboardRequest) {
	request = &CreateDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateDashboard")
	return
}

func NewCreateDashboardResponse() (response *CreateDashboardResponse) {
	response = &CreateDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建仪表盘
func (c *Client) CreateDashboard(request *CreateDashboardRequest) (response *CreateDashboardResponse, err error) {
	if request == nil {
		request = NewCreateDashboardRequest()
	}
	response = NewCreateDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCosRechargeRequest() (request *ModifyCosRechargeRequest) {
	request = &ModifyCosRechargeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyCosRecharge")
	return
}

func NewModifyCosRechargeResponse() (response *ModifyCosRechargeResponse) {
	response = &ModifyCosRechargeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改cos导入任务
func (c *Client) ModifyCosRecharge(request *ModifyCosRechargeRequest) (response *ModifyCosRechargeResponse, err error) {
	if request == nil {
		request = NewModifyCosRechargeRequest()
	}
	response = NewModifyCosRechargeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteShipperRequest() (request *DeleteShipperRequest) {
	request = &DeleteShipperRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteShipper")
	return
}

func NewDeleteShipperResponse() (response *DeleteShipperResponse) {
	response = &DeleteShipperResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除投递规则
func (c *Client) DeleteShipper(request *DeleteShipperRequest) (response *DeleteShipperResponse, err error) {
	if request == nil {
		request = NewDeleteShipperRequest()
	}
	response = NewDeleteShipperResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmNoticesRequest() (request *DescribeAlarmNoticesRequest) {
	request = &DescribeAlarmNoticesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarmNotices")
	return
}

func NewDescribeAlarmNoticesResponse() (response *DescribeAlarmNoticesResponse) {
	response = &DescribeAlarmNoticesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于获取告警通知模板列表
func (c *Client) DescribeAlarmNotices(request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmNoticesRequest()
	}
	response = NewDescribeAlarmNoticesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigMachineGroupsRequest() (request *DescribeConfigMachineGroupsRequest) {
	request = &DescribeConfigMachineGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigMachineGroups")
	return
}

func NewDescribeConfigMachineGroupsResponse() (response *DescribeConfigMachineGroupsResponse) {
	response = &DescribeConfigMachineGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集规则配置所绑定的机器组
func (c *Client) DescribeConfigMachineGroups(request *DescribeConfigMachineGroupsRequest) (response *DescribeConfigMachineGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeConfigMachineGroupsRequest()
	}
	response = NewDescribeConfigMachineGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewUploadLogRequest() (request *UploadLogRequest) {
	request = &UploadLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "UploadLog")
	return
}

func NewUploadLogResponse() (response *UploadLogResponse) {
	response = &UploadLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为了保障您日志数据的可靠性以及更高效地使用日志服务，建议您使用CLS优化后的接口[上传结构化日志](https://cloud.tencent.com/document/api/614/16873)上传日志。
func (c *Client) UploadLog(request *UploadLogRequest) (response *UploadLogResponse, err error) {
	if request == nil {
		request = NewUploadLogRequest()
	}
	response = NewUploadLogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateQcloudHourRequest() (request *CreateQcloudHourRequest) {
	request = &CreateQcloudHourRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateQcloudHour")
	return
}

func NewCreateQcloudHourResponse() (response *CreateQcloudHourResponse) {
	response = &CreateQcloudHourResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateQcloudHour
func (c *Client) CreateQcloudHour(request *CreateQcloudHourRequest) (response *CreateQcloudHourResponse, err error) {
	if request == nil {
		request = NewCreateQcloudHourRequest()
	}
	response = NewCreateQcloudHourResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConsumerRequest() (request *DeleteConsumerRequest) {
	request = &DeleteConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConsumer")
	return
}

func NewDeleteConsumerResponse() (response *DeleteConsumerResponse) {
	response = &DeleteConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除投递配置
func (c *Client) DeleteConsumer(request *DeleteConsumerRequest) (response *DeleteConsumerResponse, err error) {
	if request == nil {
		request = NewDeleteConsumerRequest()
	}
	response = NewDeleteConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeliverCloudFunctionRequest() (request *CreateDeliverCloudFunctionRequest) {
	request = &CreateDeliverCloudFunctionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateDeliverCloudFunction")
	return
}

func NewCreateDeliverCloudFunctionResponse() (response *CreateDeliverCloudFunctionResponse) {
	response = &CreateDeliverCloudFunctionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建投递SCF任务
func (c *Client) CreateDeliverCloudFunction(request *CreateDeliverCloudFunctionRequest) (response *CreateDeliverCloudFunctionResponse, err error) {
	if request == nil {
		request = NewCreateDeliverCloudFunctionRequest()
	}
	response = NewCreateDeliverCloudFunctionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
	request = &DescribeResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeResources")
	return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
	response = &DescribeResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全局或指定地域指标资源
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeResourcesRequest()
	}
	response = NewDescribeResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateShipperRequest() (request *CreateShipperRequest) {
	request = &CreateShipperRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateShipper")
	return
}

func NewCreateShipperResponse() (response *CreateShipperResponse) {
	response = &CreateShipperResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建新的投递规则，客户如果使用此接口，需要自行处理CLS对指定bucket的写权限。
func (c *Client) CreateShipper(request *CreateShipperRequest) (response *CreateShipperResponse, err error) {
	if request == nil {
		request = NewCreateShipperRequest()
	}
	response = NewCreateShipperResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTopicExtendConfigRequest() (request *DeleteTopicExtendConfigRequest) {
	request = &DeleteTopicExtendConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteTopicExtendConfig")
	return
}

func NewDeleteTopicExtendConfigResponse() (response *DeleteTopicExtendConfigResponse) {
	response = &DeleteTopicExtendConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除采集配置(clb专用)。
func (c *Client) DeleteTopicExtendConfig(request *DeleteTopicExtendConfigRequest) (response *DeleteTopicExtendConfigResponse, err error) {
	if request == nil {
		request = NewDeleteTopicExtendConfigRequest()
	}
	response = NewDeleteTopicExtendConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDemonstrationsRequest() (request *DescribeDemonstrationsRequest) {
	request = &DescribeDemonstrationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDemonstrations")
	return
}

func NewDescribeDemonstrationsResponse() (response *DescribeDemonstrationsResponse) {
	response = &DescribeDemonstrationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取日志服务演示示例列表
func (c *Client) DescribeDemonstrations(request *DescribeDemonstrationsRequest) (response *DescribeDemonstrationsResponse, err error) {
	if request == nil {
		request = NewDescribeDemonstrationsRequest()
	}
	response = NewDescribeDemonstrationsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmRequest() (request *ModifyAlarmRequest) {
	request = &ModifyAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarm")
	return
}

func NewModifyAlarmResponse() (response *ModifyAlarmResponse) {
	response = &ModifyAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改告警策略。需要至少修改一项有效内容。
func (c *Client) ModifyAlarm(request *ModifyAlarmRequest) (response *ModifyAlarmResponse, err error) {
	if request == nil {
		request = NewModifyAlarmRequest()
	}
	response = NewModifyAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDeliverCloudFunctionRequest() (request *ModifyDeliverCloudFunctionRequest) {
	request = &ModifyDeliverCloudFunctionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyDeliverCloudFunction")
	return
}

func NewModifyDeliverCloudFunctionResponse() (response *ModifyDeliverCloudFunctionResponse) {
	response = &ModifyDeliverCloudFunctionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改投递SCF任务
func (c *Client) ModifyDeliverCloudFunction(request *ModifyDeliverCloudFunctionRequest) (response *ModifyDeliverCloudFunctionResponse, err error) {
	if request == nil {
		request = NewModifyDeliverCloudFunctionRequest()
	}
	response = NewModifyDeliverCloudFunctionResponse()
	err = c.Send(request, response)
	return
}

func NewRetryShipperTaskRequest() (request *RetryShipperTaskRequest) {
	request = &RetryShipperTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "RetryShipperTask")
	return
}

func NewRetryShipperTaskResponse() (response *RetryShipperTaskResponse) {
	response = &RetryShipperTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试失败的投递任务
func (c *Client) RetryShipperTask(request *RetryShipperTaskRequest) (response *RetryShipperTaskResponse, err error) {
	if request == nil {
		request = NewRetryShipperTaskRequest()
	}
	response = NewRetryShipperTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogsetRequest() (request *DeleteLogsetRequest) {
	request = &DeleteLogsetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteLogset")
	return
}

func NewDeleteLogsetResponse() (response *DeleteLogsetResponse) {
	response = &DeleteLogsetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除日志集。
func (c *Client) DeleteLogset(request *DeleteLogsetRequest) (response *DeleteLogsetResponse, err error) {
	if request == nil {
		request = NewDeleteLogsetRequest()
	}
	response = NewDeleteLogsetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlertRecordHistoryRequest() (request *DescribeAlertRecordHistoryRequest) {
	request = &DescribeAlertRecordHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAlertRecordHistory")
	return
}

func NewDescribeAlertRecordHistoryResponse() (response *DescribeAlertRecordHistoryResponse) {
	response = &DescribeAlertRecordHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警历史记录
func (c *Client) DescribeAlertRecordHistory(request *DescribeAlertRecordHistoryRequest) (response *DescribeAlertRecordHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeAlertRecordHistoryRequest()
	}
	response = NewDescribeAlertRecordHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCosRechargeRequest() (request *CreateCosRechargeRequest) {
	request = &CreateCosRechargeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateCosRecharge")
	return
}

func NewCreateCosRechargeResponse() (response *CreateCosRechargeResponse) {
	response = &CreateCosRechargeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建cos导入任务
func (c *Client) CreateCosRecharge(request *CreateCosRechargeRequest) (response *CreateCosRechargeResponse, err error) {
	if request == nil {
		request = NewCreateCosRechargeRequest()
	}
	response = NewCreateCosRechargeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAlarmNoticeRequest() (request *DeleteAlarmNoticeRequest) {
	request = &DeleteAlarmNoticeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarmNotice")
	return
}

func NewDeleteAlarmNoticeResponse() (response *DeleteAlarmNoticeResponse) {
	response = &DeleteAlarmNoticeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于删除告警通知模板
func (c *Client) DeleteAlarmNotice(request *DeleteAlarmNoticeRequest) (response *DeleteAlarmNoticeResponse, err error) {
	if request == nil {
		request = NewDeleteAlarmNoticeRequest()
	}
	response = NewDeleteAlarmNoticeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConfigExtraRequest() (request *ModifyConfigExtraRequest) {
	request = &ModifyConfigExtraRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyConfigExtra")
	return
}

func NewModifyConfigExtraResponse() (response *ModifyConfigExtraResponse) {
	response = &ModifyConfigExtraResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改特殊采集配置任务
func (c *Client) ModifyConfigExtra(request *ModifyConfigExtraRequest) (response *ModifyConfigExtraResponse, err error) {
	if request == nil {
		request = NewModifyConfigExtraRequest()
	}
	response = NewModifyConfigExtraResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAlarmRequest() (request *CreateAlarmRequest) {
	request = &CreateAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateAlarm")
	return
}

func NewCreateAlarmResponse() (response *CreateAlarmResponse) {
	response = &CreateAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建告警策略。
func (c *Client) CreateAlarm(request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
	if request == nil {
		request = NewCreateAlarmRequest()
	}
	response = NewCreateAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewModifyShipperRequest() (request *ModifyShipperRequest) {
	request = &ModifyShipperRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyShipper")
	return
}

func NewModifyShipperResponse() (response *ModifyShipperResponse) {
	response = &ModifyShipperResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改现有的投递规则，客户如果使用此接口，需要自行处理CLS对指定bucket的写权限。
func (c *Client) ModifyShipper(request *ModifyShipperRequest) (response *ModifyShipperResponse, err error) {
	if request == nil {
		request = NewModifyShipperRequest()
	}
	response = NewModifyShipperResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDemonstrationRequest() (request *CreateDemonstrationRequest) {
	request = &CreateDemonstrationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateDemonstration")
	return
}

func NewCreateDemonstrationResponse() (response *CreateDemonstrationResponse) {
	response = &CreateDemonstrationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建演示示例
func (c *Client) CreateDemonstration(request *CreateDemonstrationRequest) (response *CreateDemonstrationResponse, err error) {
	if request == nil {
		request = NewCreateDemonstrationRequest()
	}
	response = NewCreateDemonstrationResponse()
	err = c.Send(request, response)
	return
}

func NewHeartBeatRequest() (request *HeartBeatRequest) {
	request = &HeartBeatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "HeartBeat")
	return
}

func NewHeartBeatResponse() (response *HeartBeatResponse) {
	response = &HeartBeatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上报当前采集机器的心跳
func (c *Client) HeartBeat(request *HeartBeatRequest) (response *HeartBeatResponse, err error) {
	if request == nil {
		request = NewHeartBeatRequest()
	}
	response = NewHeartBeatResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAsyncSearchTaskRequest() (request *CreateAsyncSearchTaskRequest) {
	request = &CreateAsyncSearchTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateAsyncSearchTask")
	return
}

func NewCreateAsyncSearchTaskResponse() (response *CreateAsyncSearchTaskResponse) {
	response = &CreateAsyncSearchTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建离线检索任务
func (c *Client) CreateAsyncSearchTask(request *CreateAsyncSearchTaskRequest) (response *CreateAsyncSearchTaskResponse, err error) {
	if request == nil {
		request = NewCreateAsyncSearchTaskRequest()
	}
	response = NewCreateAsyncSearchTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDemonstrationsRequest() (request *DeleteDemonstrationsRequest) {
	request = &DeleteDemonstrationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteDemonstrations")
	return
}

func NewDeleteDemonstrationsResponse() (response *DeleteDemonstrationsResponse) {
	response = &DeleteDemonstrationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除演示示例
func (c *Client) DeleteDemonstrations(request *DeleteDemonstrationsRequest) (response *DeleteDemonstrationsResponse, err error) {
	if request == nil {
		request = NewDeleteDemonstrationsRequest()
	}
	response = NewDeleteDemonstrationsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogsetRequest() (request *ModifyLogsetRequest) {
	request = &ModifyLogsetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyLogset")
	return
}

func NewModifyLogsetResponse() (response *ModifyLogsetResponse) {
	response = &ModifyLogsetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改日志集信息
func (c *Client) ModifyLogset(request *ModifyLogsetRequest) (response *ModifyLogsetResponse, err error) {
	if request == nil {
		request = NewModifyLogsetRequest()
	}
	response = NewModifyLogsetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyKafkaConsumerRequest() (request *ModifyKafkaConsumerRequest) {
	request = &ModifyKafkaConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyKafkaConsumer")
	return
}

func NewModifyKafkaConsumerResponse() (response *ModifyKafkaConsumerResponse) {
	response = &ModifyKafkaConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Kafka协议消费信息
func (c *Client) ModifyKafkaConsumer(request *ModifyKafkaConsumerRequest) (response *ModifyKafkaConsumerResponse, err error) {
	if request == nil {
		request = NewModifyKafkaConsumerRequest()
	}
	response = NewModifyKafkaConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
	request = &DescribeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigs")
	return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
	response = &DescribeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集规则配置
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeConfigsRequest()
	}
	response = NewDescribeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDemonstrationsRequest() (request *CreateDemonstrationsRequest) {
	request = &CreateDemonstrationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateDemonstrations")
	return
}

func NewCreateDemonstrationsResponse() (response *CreateDemonstrationsResponse) {
	response = &CreateDemonstrationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于批量创建演示示例
func (c *Client) CreateDemonstrations(request *CreateDemonstrationsRequest) (response *CreateDemonstrationsResponse, err error) {
	if request == nil {
		request = NewCreateDemonstrationsRequest()
	}
	response = NewCreateDemonstrationsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
	request = &DescribeTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeTemplates")
	return
}

func NewDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
	response = &DescribeTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取模版列表
func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeTemplatesRequest()
	}
	response = NewDescribeTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserConfigRequest() (request *ModifyUserConfigRequest) {
	request = &ModifyUserConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyUserConfig")
	return
}

func NewModifyUserConfigResponse() (response *ModifyUserConfigResponse) {
	response = &ModifyUserConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户配置
func (c *Client) ModifyUserConfig(request *ModifyUserConfigRequest) (response *ModifyUserConfigResponse, err error) {
	if request == nil {
		request = NewModifyUserConfigRequest()
	}
	response = NewModifyUserConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAlarmChannelRequest() (request *CheckAlarmChannelRequest) {
	request = &CheckAlarmChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckAlarmChannel")
	return
}

func NewCheckAlarmChannelResponse() (response *CheckAlarmChannelResponse) {
	response = &CheckAlarmChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警渠道检测接口
func (c *Client) CheckAlarmChannel(request *CheckAlarmChannelRequest) (response *CheckAlarmChannelResponse, err error) {
	if request == nil {
		request = NewCheckAlarmChannelRequest()
	}
	response = NewCheckAlarmChannelResponse()
	err = c.Send(request, response)
	return
}

func NewCheckFunctionRequest() (request *CheckFunctionRequest) {
	request = &CheckFunctionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckFunction")
	return
}

func NewCheckFunctionResponse() (response *CheckFunctionResponse) {
	response = &CheckFunctionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于语法校验接口。
func (c *Client) CheckFunction(request *CheckFunctionRequest) (response *CheckFunctionResponse, err error) {
	if request == nil {
		request = NewCheckFunctionRequest()
	}
	response = NewCheckFunctionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineGroupConfigsRequest() (request *DescribeMachineGroupConfigsRequest) {
	request = &DescribeMachineGroupConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeMachineGroupConfigs")
	return
}

func NewDescribeMachineGroupConfigsResponse() (response *DescribeMachineGroupConfigsResponse) {
	response = &DescribeMachineGroupConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机器组绑定的采集规则配置
func (c *Client) DescribeMachineGroupConfigs(request *DescribeMachineGroupConfigsRequest) (response *DescribeMachineGroupConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeMachineGroupConfigsRequest()
	}
	response = NewDescribeMachineGroupConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicsRequest() (request *DescribeTopicsRequest) {
	request = &DescribeTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeTopics")
	return
}

func NewDescribeTopicsResponse() (response *DescribeTopicsResponse) {
	response = &DescribeTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取日志主题列表，支持分页
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeTopicsRequest()
	}
	response = NewDescribeTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindDeliverCloudFunctionRequest() (request *UnbindDeliverCloudFunctionRequest) {
	request = &UnbindDeliverCloudFunctionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "UnbindDeliverCloudFunction")
	return
}

func NewUnbindDeliverCloudFunctionResponse() (response *UnbindDeliverCloudFunctionResponse) {
	response = &UnbindDeliverCloudFunctionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于trigger删除投递SCF任务
func (c *Client) UnbindDeliverCloudFunction(request *UnbindDeliverCloudFunctionRequest) (response *UnbindDeliverCloudFunctionResponse, err error) {
	if request == nil {
		request = NewUnbindDeliverCloudFunctionRequest()
	}
	response = NewUnbindDeliverCloudFunctionResponse()
	err = c.Send(request, response)
	return
}

func NewActionRequest() (request *ActionRequest) {
	request = &ActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "Action")
	return
}

func NewActionResponse() (response *ActionResponse) {
	response = &ActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// dvsd
func (c *Client) Action(request *ActionRequest) (response *ActionResponse, err error) {
	if request == nil {
		request = NewActionRequest()
	}
	response = NewActionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAlarmNoticeRequest() (request *CreateAlarmNoticeRequest) {
	request = &CreateAlarmNoticeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateAlarmNotice")
	return
}

func NewCreateAlarmNoticeResponse() (response *CreateAlarmNoticeResponse) {
	response = &CreateAlarmNoticeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用户创建告警通知模板。
func (c *Client) CreateAlarmNotice(request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
	if request == nil {
		request = NewCreateAlarmNoticeRequest()
	}
	response = NewCreateAlarmNoticeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserConfigRequest() (request *DescribeUserConfigRequest) {
	request = &DescribeUserConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeUserConfig")
	return
}

func NewDescribeUserConfigResponse() (response *DescribeUserConfigResponse) {
	response = &DescribeUserConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户配置信息
func (c *Client) DescribeUserConfig(request *DescribeUserConfigRequest) (response *DescribeUserConfigResponse, err error) {
	if request == nil {
		request = NewDescribeUserConfigRequest()
	}
	response = NewDescribeUserConfigResponse()
	err = c.Send(request, response)
	return
}

func NewAddMachineGroupInfoRequest() (request *AddMachineGroupInfoRequest) {
	request = &AddMachineGroupInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "AddMachineGroupInfo")
	return
}

func NewAddMachineGroupInfoResponse() (response *AddMachineGroupInfoResponse) {
	response = &AddMachineGroupInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于添加机器组信息
func (c *Client) AddMachineGroupInfo(request *AddMachineGroupInfoRequest) (response *AddMachineGroupInfoResponse, err error) {
	if request == nil {
		request = NewAddMachineGroupInfoRequest()
	}
	response = NewAddMachineGroupInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCloseKafkaConsumerRequest() (request *CloseKafkaConsumerRequest) {
	request = &CloseKafkaConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CloseKafkaConsumer")
	return
}

func NewCloseKafkaConsumerResponse() (response *CloseKafkaConsumerResponse) {
	response = &CloseKafkaConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭Kafka协议消费
func (c *Client) CloseKafkaConsumer(request *CloseKafkaConsumerRequest) (response *CloseKafkaConsumerResponse, err error) {
	if request == nil {
		request = NewCloseKafkaConsumerRequest()
	}
	response = NewCloseKafkaConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShippersRequest() (request *DescribeShippersRequest) {
	request = &DescribeShippersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeShippers")
	return
}

func NewDescribeShippersResponse() (response *DescribeShippersResponse) {
	response = &DescribeShippersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取投递规则信息列表
func (c *Client) DescribeShippers(request *DescribeShippersRequest) (response *DescribeShippersResponse, err error) {
	if request == nil {
		request = NewDescribeShippersRequest()
	}
	response = NewDescribeShippersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConsumerRequest() (request *CreateConsumerRequest) {
	request = &CreateConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateConsumer")
	return
}

func NewCreateConsumerResponse() (response *CreateConsumerResponse) {
	response = &CreateConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建投递任务
func (c *Client) CreateConsumer(request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
	if request == nil {
		request = NewCreateConsumerRequest()
	}
	response = NewCreateConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentConfigsRequest() (request *DescribeAgentConfigsRequest) {
	request = &DescribeAgentConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAgentConfigs")
	return
}

func NewDescribeAgentConfigsResponse() (response *DescribeAgentConfigsResponse) {
	response = &DescribeAgentConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取agent对应的采集配置
func (c *Client) DescribeAgentConfigs(request *DescribeAgentConfigsRequest) (response *DescribeAgentConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeAgentConfigsRequest()
	}
	response = NewDescribeAgentConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDataTransformRequest() (request *CreateDataTransformRequest) {
	request = &CreateDataTransformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateDataTransform")
	return
}

func NewCreateDataTransformResponse() (response *CreateDataTransformResponse) {
	response = &CreateDataTransformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建数据加工任务。
func (c *Client) CreateDataTransform(request *CreateDataTransformRequest) (response *CreateDataTransformResponse, err error) {
	if request == nil {
		request = NewCreateDataTransformRequest()
	}
	response = NewCreateDataTransformResponse()
	err = c.Send(request, response)
	return
}

func NewOpenKafkaConsumeRequest() (request *OpenKafkaConsumeRequest) {
	request = &OpenKafkaConsumeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "OpenKafkaConsume")
	return
}

func NewOpenKafkaConsumeResponse() (response *OpenKafkaConsumeResponse) {
	response = &OpenKafkaConsumeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 打开kafka消费功能
func (c *Client) OpenKafkaConsume(request *OpenKafkaConsumeRequest) (response *OpenKafkaConsumeResponse, err error) {
	if request == nil {
		request = NewOpenKafkaConsumeRequest()
	}
	response = NewOpenKafkaConsumeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmNoticeRequest() (request *ModifyAlarmNoticeRequest) {
	request = &ModifyAlarmNoticeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarmNotice")
	return
}

func NewModifyAlarmNoticeResponse() (response *ModifyAlarmNoticeResponse) {
	response = &ModifyAlarmNoticeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于修改告警通知模板。
func (c *Client) ModifyAlarmNotice(request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
	if request == nil {
		request = NewModifyAlarmNoticeRequest()
	}
	response = NewModifyAlarmNoticeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConsumerRequest() (request *ModifyConsumerRequest) {
	request = &ModifyConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyConsumer")
	return
}

func NewModifyConsumerResponse() (response *ModifyConsumerResponse) {
	response = &ModifyConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改投递任务
func (c *Client) ModifyConsumer(request *ModifyConsumerRequest) (response *ModifyConsumerResponse, err error) {
	if request == nil {
		request = NewModifyConsumerRequest()
	}
	response = NewModifyConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogFastAnalysisRequest() (request *DescribeLogFastAnalysisRequest) {
	request = &DescribeLogFastAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLogFastAnalysis")
	return
}

func NewDescribeLogFastAnalysisResponse() (response *DescribeLogFastAnalysisResponse) {
	response = &DescribeLogFastAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 针对日志某个字段可以做快速分析其占比情况
func (c *Client) DescribeLogFastAnalysis(request *DescribeLogFastAnalysisRequest) (response *DescribeLogFastAnalysisResponse, err error) {
	if request == nil {
		request = NewDescribeLogFastAnalysisRequest()
	}
	response = NewDescribeLogFastAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewApplyConfigToMachineGroupRequest() (request *ApplyConfigToMachineGroupRequest) {
	request = &ApplyConfigToMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ApplyConfigToMachineGroup")
	return
}

func NewApplyConfigToMachineGroupResponse() (response *ApplyConfigToMachineGroupResponse) {
	response = &ApplyConfigToMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 应用采集配置到指定机器组
func (c *Client) ApplyConfigToMachineGroup(request *ApplyConfigToMachineGroupRequest) (response *ApplyConfigToMachineGroupResponse, err error) {
	if request == nil {
		request = NewApplyConfigToMachineGroupRequest()
	}
	response = NewApplyConfigToMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewRealtimeProducerRequest() (request *RealtimeProducerRequest) {
	request = &RealtimeProducerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "RealtimeProducer")
	return
}

func NewRealtimeProducerResponse() (response *RealtimeProducerResponse) {
	response = &RealtimeProducerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// kafka协议消费接口
func (c *Client) RealtimeProducer(request *RealtimeProducerRequest) (response *RealtimeProducerResponse, err error) {
	if request == nil {
		request = NewRealtimeProducerRequest()
	}
	response = NewRealtimeProducerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformProcessInfoRequest() (request *DescribeDataTransformProcessInfoRequest) {
	request = &DescribeDataTransformProcessInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformProcessInfo")
	return
}

func NewDescribeDataTransformProcessInfoResponse() (response *DescribeDataTransformProcessInfoResponse) {
	response = &DescribeDataTransformProcessInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取数据加工任务执行进度详情
func (c *Client) DescribeDataTransformProcessInfo(request *DescribeDataTransformProcessInfoRequest) (response *DescribeDataTransformProcessInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformProcessInfoRequest()
	}
	response = NewDescribeDataTransformProcessInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndexRequest() (request *DescribeIndexRequest) {
	request = &DescribeIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeIndex")
	return
}

func NewDescribeIndexResponse() (response *DescribeIndexResponse) {
	response = &DescribeIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取索引配置信息
func (c *Client) DescribeIndex(request *DescribeIndexRequest) (response *DescribeIndexResponse, err error) {
	if request == nil {
		request = NewDescribeIndexRequest()
	}
	response = NewDescribeIndexResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLatestJsonLogRequest() (request *DescribeLatestJsonLogRequest) {
	request = &DescribeLatestJsonLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLatestJsonLog")
	return
}

func NewDescribeLatestJsonLogResponse() (response *DescribeLatestJsonLogResponse) {
	response = &DescribeLatestJsonLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取json格式日志
func (c *Client) DescribeLatestJsonLog(request *DescribeLatestJsonLogRequest) (response *DescribeLatestJsonLogResponse, err error) {
	if request == nil {
		request = NewDescribeLatestJsonLogRequest()
	}
	response = NewDescribeLatestJsonLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDataTransformRequest() (request *ModifyDataTransformRequest) {
	request = &ModifyDataTransformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyDataTransform")
	return
}

func NewModifyDataTransformResponse() (response *ModifyDataTransformResponse) {
	response = &ModifyDataTransformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改数据加工任务
func (c *Client) ModifyDataTransform(request *ModifyDataTransformRequest) (response *ModifyDataTransformResponse, err error) {
	if request == nil {
		request = NewModifyDataTransformRequest()
	}
	response = NewModifyDataTransformResponse()
	err = c.Send(request, response)
	return
}

func NewSearchCosRechargeInfoRequest() (request *SearchCosRechargeInfoRequest) {
	request = &SearchCosRechargeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "SearchCosRechargeInfo")
	return
}

func NewSearchCosRechargeInfoResponse() (response *SearchCosRechargeInfoResponse) {
	response = &SearchCosRechargeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于预览cos导入信息
func (c *Client) SearchCosRechargeInfo(request *SearchCosRechargeInfoRequest) (response *SearchCosRechargeInfoResponse, err error) {
	if request == nil {
		request = NewSearchCosRechargeInfoRequest()
	}
	response = NewSearchCosRechargeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMachineGroupRequest() (request *DeleteMachineGroupRequest) {
	request = &DeleteMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteMachineGroup")
	return
}

func NewDeleteMachineGroupResponse() (response *DeleteMachineGroupResponse) {
	response = &DeleteMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机器组
func (c *Client) DeleteMachineGroup(request *DeleteMachineGroupRequest) (response *DeleteMachineGroupResponse, err error) {
	if request == nil {
		request = NewDeleteMachineGroupRequest()
	}
	response = NewDeleteMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicExtendConfigRequest() (request *DescribeTopicExtendConfigRequest) {
	request = &DescribeTopicExtendConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeTopicExtendConfig")
	return
}

func NewDescribeTopicExtendConfigResponse() (response *DescribeTopicExtendConfigResponse) {
	response = &DescribeTopicExtendConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集配置(clb专用)。
func (c *Client) DescribeTopicExtendConfig(request *DescribeTopicExtendConfigRequest) (response *DescribeTopicExtendConfigResponse, err error) {
	if request == nil {
		request = NewDescribeTopicExtendConfigRequest()
	}
	response = NewDescribeTopicExtendConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePartitionsRequest() (request *DescribePartitionsRequest) {
	request = &DescribePartitionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribePartitions")
	return
}

func NewDescribePartitionsResponse() (response *DescribePartitionsResponse) {
	response = &DescribePartitionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取主题分区列表。
func (c *Client) DescribePartitions(request *DescribePartitionsRequest) (response *DescribePartitionsResponse, err error) {
	if request == nil {
		request = NewDescribePartitionsRequest()
	}
	response = NewDescribePartitionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDeliverCloudFunctionRequest() (request *DeleteDeliverCloudFunctionRequest) {
	request = &DeleteDeliverCloudFunctionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteDeliverCloudFunction")
	return
}

func NewDeleteDeliverCloudFunctionResponse() (response *DeleteDeliverCloudFunctionResponse) {
	response = &DeleteDeliverCloudFunctionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除投递SCF任务
func (c *Client) DeleteDeliverCloudFunction(request *DeleteDeliverCloudFunctionRequest) (response *DeleteDeliverCloudFunctionResponse, err error) {
	if request == nil {
		request = NewDeleteDeliverCloudFunctionRequest()
	}
	response = NewDeleteDeliverCloudFunctionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAsyncSearchResultRequest() (request *DescribeAsyncSearchResultRequest) {
	request = &DescribeAsyncSearchResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAsyncSearchResult")
	return
}

func NewDescribeAsyncSearchResultResponse() (response *DescribeAsyncSearchResultResponse) {
	response = &DescribeAsyncSearchResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用户获取离线检索任务的结果
func (c *Client) DescribeAsyncSearchResult(request *DescribeAsyncSearchResultRequest) (response *DescribeAsyncSearchResultResponse, err error) {
	if request == nil {
		request = NewDescribeAsyncSearchResultRequest()
	}
	response = NewDescribeAsyncSearchResultResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAlarmRuleRequest() (request *CheckAlarmRuleRequest) {
	request = &CheckAlarmRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckAlarmRule")
	return
}

func NewCheckAlarmRuleResponse() (response *CheckAlarmRuleResponse) {
	response = &CheckAlarmRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警策略检测接口
func (c *Client) CheckAlarmRule(request *CheckAlarmRuleRequest) (response *CheckAlarmRuleResponse, err error) {
	if request == nil {
		request = NewCheckAlarmRuleRequest()
	}
	response = NewCheckAlarmRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMachineGroupRequest() (request *ModifyMachineGroupRequest) {
	request = &ModifyMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyMachineGroup")
	return
}

func NewModifyMachineGroupResponse() (response *ModifyMachineGroupResponse) {
	response = &ModifyMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机器组
func (c *Client) ModifyMachineGroup(request *ModifyMachineGroupRequest) (response *ModifyMachineGroupResponse, err error) {
	if request == nil {
		request = NewModifyMachineGroupRequest()
	}
	response = NewModifyMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportsRequest() (request *DescribeExportsRequest) {
	request = &DescribeExportsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeExports")
	return
}

func NewDescribeExportsResponse() (response *DescribeExportsResponse) {
	response = &DescribeExportsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取日志下载任务列表
func (c *Client) DescribeExports(request *DescribeExportsRequest) (response *DescribeExportsResponse, err error) {
	if request == nil {
		request = NewDescribeExportsRequest()
	}
	response = NewDescribeExportsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRebuildIndexTaskRequest() (request *CreateRebuildIndexTaskRequest) {
	request = &CreateRebuildIndexTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateRebuildIndexTask")
	return
}

func NewCreateRebuildIndexTaskResponse() (response *CreateRebuildIndexTaskResponse) {
	response = &CreateRebuildIndexTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建重建索引任务
func (c *Client) CreateRebuildIndexTask(request *CreateRebuildIndexTaskRequest) (response *CreateRebuildIndexTaskResponse, err error) {
	if request == nil {
		request = NewCreateRebuildIndexTaskRequest()
	}
	response = NewCreateRebuildIndexTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAlarmRequest() (request *DeleteAlarmRequest) {
	request = &DeleteAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarm")
	return
}

func NewDeleteAlarmResponse() (response *DeleteAlarmResponse) {
	response = &DeleteAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除告警策略。
func (c *Client) DeleteAlarm(request *DeleteAlarmRequest) (response *DeleteAlarmResponse, err error) {
	if request == nil {
		request = NewDeleteAlarmRequest()
	}
	response = NewDeleteAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigRequest() (request *DeleteConfigRequest) {
	request = &DeleteConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConfig")
	return
}

func NewDeleteConfigResponse() (response *DeleteConfigResponse) {
	response = &DeleteConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除采集规则配置
func (c *Client) DeleteConfig(request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
	if request == nil {
		request = NewDeleteConfigRequest()
	}
	response = NewDeleteConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIndexRequest() (request *ModifyIndexRequest) {
	request = &ModifyIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyIndex")
	return
}

func NewModifyIndexResponse() (response *ModifyIndexResponse) {
	response = &ModifyIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改索引配置
func (c *Client) ModifyIndex(request *ModifyIndexRequest) (response *ModifyIndexResponse, err error) {
	if request == nil {
		request = NewModifyIndexRequest()
	}
	response = NewModifyIndexResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformPreviewDataInfoRequest() (request *DescribeDataTransformPreviewDataInfoRequest) {
	request = &DescribeDataTransformPreviewDataInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformPreviewDataInfo")
	return
}

func NewDescribeDataTransformPreviewDataInfoResponse() (response *DescribeDataTransformPreviewDataInfoResponse) {
	response = &DescribeDataTransformPreviewDataInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取数据加工预览详细数据。
func (c *Client) DescribeDataTransformPreviewDataInfo(request *DescribeDataTransformPreviewDataInfoRequest) (response *DescribeDataTransformPreviewDataInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformPreviewDataInfoRequest()
	}
	response = NewDescribeDataTransformPreviewDataInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformFailLogInfoRequest() (request *DescribeDataTransformFailLogInfoRequest) {
	request = &DescribeDataTransformFailLogInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformFailLogInfo")
	return
}

func NewDescribeDataTransformFailLogInfoResponse() (response *DescribeDataTransformFailLogInfoResponse) {
	response = &DescribeDataTransformFailLogInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取数据加工任务失败日志详请
func (c *Client) DescribeDataTransformFailLogInfo(request *DescribeDataTransformFailLogInfoRequest) (response *DescribeDataTransformFailLogInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformFailLogInfoRequest()
	}
	response = NewDescribeDataTransformFailLogInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigRequest() (request *CreateConfigRequest) {
	request = &CreateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateConfig")
	return
}

func NewCreateConfigResponse() (response *CreateConfigResponse) {
	response = &CreateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建采集规则配置
func (c *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
	if request == nil {
		request = NewCreateConfigRequest()
	}
	response = NewCreateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExportRequest() (request *CreateExportRequest) {
	request = &CreateExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateExport")
	return
}

func NewCreateExportResponse() (response *CreateExportResponse) {
	response = &CreateExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口仅创建下载任务，任务返回的下载地址，请用户调用DescribeExports查看任务列表。其中有下载地址CosPath参数。参考文档https://cloud.tencent.com/document/product/614/56449
func (c *Client) CreateExport(request *CreateExportRequest) (response *CreateExportResponse, err error) {
	if request == nil {
		request = NewCreateExportRequest()
	}
	response = NewCreateExportResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMachineGroupInfoRequest() (request *DeleteMachineGroupInfoRequest) {
	request = &DeleteMachineGroupInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteMachineGroupInfo")
	return
}

func NewDeleteMachineGroupInfoResponse() (response *DeleteMachineGroupInfoResponse) {
	response = &DeleteMachineGroupInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于删除机器组信息
func (c *Client) DeleteMachineGroupInfo(request *DeleteMachineGroupInfoRequest) (response *DeleteMachineGroupInfoResponse, err error) {
	if request == nil {
		request = NewDeleteMachineGroupInfoRequest()
	}
	response = NewDeleteMachineGroupInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaUserRequest() (request *DescribeKafkaUserRequest) {
	request = &DescribeKafkaUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaUser")
	return
}

func NewDescribeKafkaUserResponse() (response *DescribeKafkaUserResponse) {
	response = &DescribeKafkaUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取kafka用户信息
func (c *Client) DescribeKafkaUser(request *DescribeKafkaUserRequest) (response *DescribeKafkaUserResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaUserRequest()
	}
	response = NewDescribeKafkaUserResponse()
	err = c.Send(request, response)
	return
}

func NewUploadServiceLogRequest() (request *UploadServiceLogRequest) {
	request = &UploadServiceLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "UploadServiceLog")
	return
}

func NewUploadServiceLogResponse() (response *UploadServiceLogResponse) {
	response = &UploadServiceLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于将Agent服务日志写入到用户名下的服务日志Topic。
func (c *Client) UploadServiceLog(request *UploadServiceLogRequest) (response *UploadServiceLogResponse, err error) {
	if request == nil {
		request = NewUploadServiceLogRequest()
	}
	response = NewUploadServiceLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
	request = &ModifyTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyTopic")
	return
}

func NewModifyTopicResponse() (response *ModifyTopicResponse) {
	response = &ModifyTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改日志主题。
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
	if request == nil {
		request = NewModifyTopicRequest()
	}
	response = NewModifyTopicResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExportRequest() (request *DeleteExportRequest) {
	request = &DeleteExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteExport")
	return
}

func NewDeleteExportResponse() (response *DeleteExportResponse) {
	response = &DeleteExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除日志下载任务
func (c *Client) DeleteExport(request *DeleteExportRequest) (response *DeleteExportResponse, err error) {
	if request == nil {
		request = NewDeleteExportRequest()
	}
	response = NewDeleteExportResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLogsetRequest() (request *CreateLogsetRequest) {
	request = &CreateLogsetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateLogset")
	return
}

func NewCreateLogsetResponse() (response *CreateLogsetResponse) {
	response = &CreateLogsetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建日志集，返回新创建的日志集的 ID。
func (c *Client) CreateLogset(request *CreateLogsetRequest) (response *CreateLogsetResponse, err error) {
	if request == nil {
		request = NewCreateLogsetRequest()
	}
	response = NewCreateLogsetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNonBillingRequest() (request *CreateNonBillingRequest) {
	request = &CreateNonBillingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateNonBilling")
	return
}

func NewCreateNonBillingResponse() (response *CreateNonBillingResponse) {
	response = &CreateNonBillingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建体验账户
func (c *Client) CreateNonBilling(request *CreateNonBillingRequest) (response *CreateNonBillingResponse, err error) {
	if request == nil {
		request = NewCreateNonBillingRequest()
	}
	response = NewCreateNonBillingResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAsyncSearchTaskRequest() (request *DeleteAsyncSearchTaskRequest) {
	request = &DeleteAsyncSearchTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteAsyncSearchTask")
	return
}

func NewDeleteAsyncSearchTaskResponse() (response *DeleteAsyncSearchTaskResponse) {
	response = &DeleteAsyncSearchTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除离线检索任务
func (c *Client) DeleteAsyncSearchTask(request *DeleteAsyncSearchTaskRequest) (response *DeleteAsyncSearchTaskResponse, err error) {
	if request == nil {
		request = NewDeleteAsyncSearchTaskRequest()
	}
	response = NewDeleteAsyncSearchTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogRequest() (request *SearchLogRequest) {
	request = &SearchLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "SearchLog")
	return
}

func NewSearchLogResponse() (response *SearchLogResponse) {
	response = &SearchLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于检索分析日志, 该接口除受默认接口请求频率限制外，针对单个日志主题，查询并发数不能超过15。
func (c *Client) SearchLog(request *SearchLogRequest) (response *SearchLogResponse, err error) {
	if request == nil {
		request = NewSearchLogRequest()
	}
	response = NewSearchLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineGroupsRequest() (request *DescribeMachineGroupsRequest) {
	request = &DescribeMachineGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeMachineGroups")
	return
}

func NewDescribeMachineGroupsResponse() (response *DescribeMachineGroupsResponse) {
	response = &DescribeMachineGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机器组信息列表
func (c *Client) DescribeMachineGroups(request *DescribeMachineGroupsRequest) (response *DescribeMachineGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeMachineGroupsRequest()
	}
	response = NewDescribeMachineGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigExtrasRequest() (request *DescribeConfigExtrasRequest) {
	request = &DescribeConfigExtrasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigExtras")
	return
}

func NewDescribeConfigExtrasResponse() (response *DescribeConfigExtrasResponse) {
	response = &DescribeConfigExtrasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取特殊采集配置
func (c *Client) DescribeConfigExtras(request *DescribeConfigExtrasRequest) (response *DescribeConfigExtrasResponse, err error) {
	if request == nil {
		request = NewDescribeConfigExtrasRequest()
	}
	response = NewDescribeConfigExtrasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogContextRequest() (request *DescribeLogContextRequest) {
	request = &DescribeLogContextRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLogContext")
	return
}

func NewDescribeLogContextResponse() (response *DescribeLogContextResponse) {
	response = &DescribeLogContextResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于搜索日志上下文附近的内容
func (c *Client) DescribeLogContext(request *DescribeLogContextRequest) (response *DescribeLogContextResponse, err error) {
	if request == nil {
		request = NewDescribeLogContextRequest()
	}
	response = NewDescribeLogContextResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlarmLogRequest() (request *GetAlarmLogRequest) {
	request = &GetAlarmLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GetAlarmLog")
	return
}

func NewGetAlarmLogResponse() (response *GetAlarmLogResponse) {
	response = &GetAlarmLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取告警任务历史
func (c *Client) GetAlarmLog(request *GetAlarmLogRequest) (response *GetAlarmLogResponse, err error) {
	if request == nil {
		request = NewGetAlarmLogRequest()
	}
	response = NewGetAlarmLogResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
	request = &DeleteTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteTopic")
	return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
	response = &DeleteTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除日志主题。
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
	if request == nil {
		request = NewDeleteTopicRequest()
	}
	response = NewDeleteTopicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformPreviewInfoRequest() (request *DescribeDataTransformPreviewInfoRequest) {
	request = &DescribeDataTransformPreviewInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformPreviewInfo")
	return
}

func NewDescribeDataTransformPreviewInfoResponse() (response *DescribeDataTransformPreviewInfoResponse) {
	response = &DescribeDataTransformPreviewInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取数据加工预览任务基本信息
func (c *Client) DescribeDataTransformPreviewInfo(request *DescribeDataTransformPreviewInfoRequest) (response *DescribeDataTransformPreviewInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformPreviewInfoRequest()
	}
	response = NewDescribeDataTransformPreviewInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTopicExtendConfigRequest() (request *ModifyTopicExtendConfigRequest) {
	request = &ModifyTopicExtendConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyTopicExtendConfig")
	return
}

func NewModifyTopicExtendConfigResponse() (response *ModifyTopicExtendConfigResponse) {
	response = &ModifyTopicExtendConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改采集配置(clb专用)。
func (c *Client) ModifyTopicExtendConfig(request *ModifyTopicExtendConfigRequest) (response *ModifyTopicExtendConfigResponse, err error) {
	if request == nil {
		request = NewModifyTopicExtendConfigRequest()
	}
	response = NewModifyTopicExtendConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAgentStatusRequest() (request *UpdateAgentStatusRequest) {
	request = &UpdateAgentStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "UpdateAgentStatus")
	return
}

func NewUpdateAgentStatusResponse() (response *UpdateAgentStatusResponse) {
	response = &UpdateAgentStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上报采集机器Agent的升级状态
func (c *Client) UpdateAgentStatus(request *UpdateAgentStatusRequest) (response *UpdateAgentStatusResponse, err error) {
	if request == nil {
		request = NewUpdateAgentStatusRequest()
	}
	response = NewUpdateAgentStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAsyncSearchTasksRequest() (request *DescribeAsyncSearchTasksRequest) {
	request = &DescribeAsyncSearchTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAsyncSearchTasks")
	return
}

func NewDescribeAsyncSearchTasksResponse() (response *DescribeAsyncSearchTasksResponse) {
	response = &DescribeAsyncSearchTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用户获取离线检索任务列表
func (c *Client) DescribeAsyncSearchTasks(request *DescribeAsyncSearchTasksRequest) (response *DescribeAsyncSearchTasksResponse, err error) {
	if request == nil {
		request = NewDescribeAsyncSearchTasksRequest()
	}
	response = NewDescribeAsyncSearchTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformInfoRequest() (request *DescribeDataTransformInfoRequest) {
	request = &DescribeDataTransformInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformInfo")
	return
}

func NewDescribeDataTransformInfoResponse() (response *DescribeDataTransformInfoResponse) {
	response = &DescribeDataTransformInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取数据加工任务列表基本信息
func (c *Client) DescribeDataTransformInfo(request *DescribeDataTransformInfoRequest) (response *DescribeDataTransformInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformInfoRequest()
	}
	response = NewDescribeDataTransformInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigExtraRequest() (request *DeleteConfigExtraRequest) {
	request = &DeleteConfigExtraRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConfigExtra")
	return
}

func NewDeleteConfigExtraResponse() (response *DeleteConfigExtraResponse) {
	response = &DeleteConfigExtraResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除特殊采集规则配置
func (c *Client) DeleteConfigExtra(request *DeleteConfigExtraRequest) (response *DeleteConfigExtraResponse, err error) {
	if request == nil {
		request = NewDeleteConfigExtraRequest()
	}
	response = NewDeleteConfigExtraResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogHistogramRequest() (request *DescribeLogHistogramRequest) {
	request = &DescribeLogHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLogHistogram")
	return
}

func NewDescribeLogHistogramResponse() (response *DescribeLogHistogramResponse) {
	response = &DescribeLogHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于构建直方图
func (c *Client) DescribeLogHistogram(request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
	if request == nil {
		request = NewDescribeLogHistogramRequest()
	}
	response = NewDescribeLogHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewCloseKafkaConsumeRequest() (request *CloseKafkaConsumeRequest) {
	request = &CloseKafkaConsumeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CloseKafkaConsume")
	return
}

func NewCloseKafkaConsumeResponse() (response *CloseKafkaConsumeResponse) {
	response = &CloseKafkaConsumeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭Kafka消费
func (c *Client) CloseKafkaConsume(request *CloseKafkaConsumeRequest) (response *CloseKafkaConsumeResponse, err error) {
	if request == nil {
		request = NewCloseKafkaConsumeRequest()
	}
	response = NewCloseKafkaConsumeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAsyncContextResultRequest() (request *DescribeAsyncContextResultRequest) {
	request = &DescribeAsyncContextResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAsyncContextResult")
	return
}

func NewDescribeAsyncContextResultResponse() (response *DescribeAsyncContextResultResponse) {
	response = &DescribeAsyncContextResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用户获取离线上下文任务的结果
func (c *Client) DescribeAsyncContextResult(request *DescribeAsyncContextResultRequest) (response *DescribeAsyncContextResultResponse, err error) {
	if request == nil {
		request = NewDescribeAsyncContextResultRequest()
	}
	response = NewDescribeAsyncContextResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDashboardsRequest() (request *DescribeDashboardsRequest) {
	request = &DescribeDashboardsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDashboards")
	return
}

func NewDescribeDashboardsResponse() (response *DescribeDashboardsResponse) {
	response = &DescribeDashboardsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取仪表盘
func (c *Client) DescribeDashboards(request *DescribeDashboardsRequest) (response *DescribeDashboardsResponse, err error) {
	if request == nil {
		request = NewDescribeDashboardsRequest()
	}
	response = NewDescribeDashboardsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDataTransformRequest() (request *DeleteDataTransformRequest) {
	request = &DeleteDataTransformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteDataTransform")
	return
}

func NewDeleteDataTransformResponse() (response *DeleteDataTransformResponse) {
	response = &DeleteDataTransformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除数据加工任务
func (c *Client) DeleteDataTransform(request *DeleteDataTransformRequest) (response *DeleteDataTransformResponse, err error) {
	if request == nil {
		request = NewDeleteDataTransformRequest()
	}
	response = NewDeleteDataTransformResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIndexRequest() (request *DeleteIndexRequest) {
	request = &DeleteIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteIndex")
	return
}

func NewDeleteIndexResponse() (response *DeleteIndexResponse) {
	response = &DeleteIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于日志主题的索引配置
func (c *Client) DeleteIndex(request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
	if request == nil {
		request = NewDeleteIndexRequest()
	}
	response = NewDeleteIndexResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTopicExtendConfigRequest() (request *CreateTopicExtendConfigRequest) {
	request = &CreateTopicExtendConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateTopicExtendConfig")
	return
}

func NewCreateTopicExtendConfigResponse() (response *CreateTopicExtendConfigResponse) {
	response = &CreateTopicExtendConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建采集配置, 不包含临时密钥信息(clb专用)。
func (c *Client) CreateTopicExtendConfig(request *CreateTopicExtendConfigRequest) (response *CreateTopicExtendConfigResponse, err error) {
	if request == nil {
		request = NewCreateTopicExtendConfigRequest()
	}
	response = NewCreateTopicExtendConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShipperTasksRequest() (request *DescribeShipperTasksRequest) {
	request = &DescribeShipperTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeShipperTasks")
	return
}

func NewDescribeShipperTasksResponse() (response *DescribeShipperTasksResponse) {
	response = &DescribeShipperTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取投递任务列表
func (c *Client) DescribeShipperTasks(request *DescribeShipperTasksRequest) (response *DescribeShipperTasksResponse, err error) {
	if request == nil {
		request = NewDescribeShipperTasksRequest()
	}
	response = NewDescribeShipperTasksResponse()
	err = c.Send(request, response)
	return
}
