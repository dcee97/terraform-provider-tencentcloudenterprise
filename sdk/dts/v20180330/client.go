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

package v20180330

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-03-30"

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

func NewCreateMigrateCheckJobRequest() (request *CreateMigrateCheckJobRequest) {
	request = &CreateMigrateCheckJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateCheckJob")
	return
}

func NewCreateMigrateCheckJobResponse() (response *CreateMigrateCheckJobResponse) {
	response = &CreateMigrateCheckJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建校验迁移任务
// 在开始迁移前, 必须调用本接口创建校验, 且校验成功后才能开始迁移. 校验的结果可以通过DescribeMigrateCheckJob查看.
// 校验成功后,迁移任务若有修改, 则必须重新创建校验并通过后, 才能开始迁移.
func (c *Client) CreateMigrateCheckJob(request *CreateMigrateCheckJobRequest) (response *CreateMigrateCheckJobResponse, err error) {
	if request == nil {
		request = NewCreateMigrateCheckJobRequest()
	}
	response = NewCreateMigrateCheckJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOffsetByTimeRequest() (request *DescribeOffsetByTimeRequest) {
	request = &DescribeOffsetByTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeOffsetByTime")
	return
}

func NewDescribeOffsetByTimeResponse() (response *DescribeOffsetByTimeResponse) {
	response = &DescribeOffsetByTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeOffsetByTime）用于通过时间获取偏移量信息。
func (c *Client) DescribeOffsetByTime(request *DescribeOffsetByTimeRequest) (response *DescribeOffsetByTimeResponse, err error) {
	if request == nil {
		request = NewDescribeOffsetByTimeRequest()
	}
	response = NewDescribeOffsetByTimeResponse()
	err = c.Send(request, response)
	return
}

func NewModifySubscribeAutoRenewFlagRequest() (request *ModifySubscribeAutoRenewFlagRequest) {
	request = &ModifySubscribeAutoRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeAutoRenewFlag")
	return
}

func NewModifySubscribeAutoRenewFlagResponse() (response *ModifySubscribeAutoRenewFlagResponse) {
	response = &ModifySubscribeAutoRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改订阅实例自动续费标识
func (c *Client) ModifySubscribeAutoRenewFlag(request *ModifySubscribeAutoRenewFlagRequest) (response *ModifySubscribeAutoRenewFlagResponse, err error) {
	if request == nil {
		request = NewModifySubscribeAutoRenewFlagRequest()
	}
	response = NewModifySubscribeAutoRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscribeRunningJobRequest() (request *DescribeSubscribeRunningJobRequest) {
	request = &DescribeSubscribeRunningJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeRunningJob")
	return
}

func NewDescribeSubscribeRunningJobResponse() (response *DescribeSubscribeRunningJobResponse) {
	response = &DescribeSubscribeRunningJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSubscribeRunningJob）用于查询任务运行情况。
func (c *Client) DescribeSubscribeRunningJob(request *DescribeSubscribeRunningJobRequest) (response *DescribeSubscribeRunningJobResponse, err error) {
	if request == nil {
		request = NewDescribeSubscribeRunningJobRequest()
	}
	response = NewDescribeSubscribeRunningJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscribeReturnableRequest() (request *DescribeSubscribeReturnableRequest) {
	request = &DescribeSubscribeReturnableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeReturnable")
	return
}

func NewDescribeSubscribeReturnableResponse() (response *DescribeSubscribeReturnableResponse) {
	response = &DescribeSubscribeReturnableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeSubscribeReturnable)用于查询订阅实例是否可以退换
func (c *Client) DescribeSubscribeReturnable(request *DescribeSubscribeReturnableRequest) (response *DescribeSubscribeReturnableResponse, err error) {
	if request == nil {
		request = NewDescribeSubscribeReturnableRequest()
	}
	response = NewDescribeSubscribeReturnableResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConsumerGroupPasswordRequest() (request *ModifyConsumerGroupPasswordRequest) {
	request = &ModifyConsumerGroupPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ModifyConsumerGroupPassword")
	return
}

func NewModifyConsumerGroupPasswordResponse() (response *ModifyConsumerGroupPasswordResponse) {
	response = &ModifyConsumerGroupPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyConsumerGroupPassword）用于修改消费组密码。
func (c *Client) ModifyConsumerGroupPassword(request *ModifyConsumerGroupPasswordRequest) (response *ModifyConsumerGroupPasswordResponse, err error) {
	if request == nil {
		request = NewModifyConsumerGroupPasswordRequest()
	}
	response = NewModifyConsumerGroupPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewModifySubscribeNameRequest() (request *ModifySubscribeNameRequest) {
	request = &ModifySubscribeNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeName")
	return
}

func NewModifySubscribeNameResponse() (response *ModifySubscribeNameResponse) {
	response = &ModifySubscribeNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifySubscribeName)用于修改数据订阅实例的名称
func (c *Client) ModifySubscribeName(request *ModifySubscribeNameRequest) (response *ModifySubscribeNameResponse, err error) {
	if request == nil {
		request = NewModifySubscribeNameRequest()
	}
	response = NewModifySubscribeNameResponse()
	err = c.Send(request, response)
	return
}

func NewModifySubscribeObjectsRequest() (request *ModifySubscribeObjectsRequest) {
	request = &ModifySubscribeObjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeObjects")
	return
}

func NewModifySubscribeObjectsResponse() (response *ModifySubscribeObjectsResponse) {
	response = &ModifySubscribeObjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifySubscribeObjects)用于修改数据订阅通道的订阅规则
func (c *Client) ModifySubscribeObjects(request *ModifySubscribeObjectsRequest) (response *ModifySubscribeObjectsResponse, err error) {
	if request == nil {
		request = NewModifySubscribeObjectsRequest()
	}
	response = NewModifySubscribeObjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaSubscribeConfRequest() (request *DescribeKafkaSubscribeConfRequest) {
	request = &DescribeKafkaSubscribeConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeKafkaSubscribeConf")
	return
}

func NewDescribeKafkaSubscribeConfResponse() (response *DescribeKafkaSubscribeConfResponse) {
	response = &DescribeKafkaSubscribeConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeKafkaSubscribeConf)用于查询数据订阅kafka版本配置信息。
func (c *Client) DescribeKafkaSubscribeConf(request *DescribeKafkaSubscribeConfRequest) (response *DescribeKafkaSubscribeConfResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaSubscribeConfRequest()
	}
	response = NewDescribeKafkaSubscribeConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionConfRequest() (request *DescribeRegionConfRequest) {
	request = &DescribeRegionConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeRegionConf")
	return
}

func NewDescribeRegionConfResponse() (response *DescribeRegionConfResponse) {
	response = &DescribeRegionConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeRegionConf）用于查询可售卖订阅实例的地域
func (c *Client) DescribeRegionConf(request *DescribeRegionConfRequest) (response *DescribeRegionConfResponse, err error) {
	if request == nil {
		request = NewDescribeRegionConfRequest()
	}
	response = NewDescribeRegionConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscribeConfRequest() (request *DescribeSubscribeConfRequest) {
	request = &DescribeSubscribeConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeConf")
	return
}

func NewDescribeSubscribeConfResponse() (response *DescribeSubscribeConfResponse) {
	response = &DescribeSubscribeConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSubscribeConf）用于查询订阅实例配置
func (c *Client) DescribeSubscribeConf(request *DescribeSubscribeConfRequest) (response *DescribeSubscribeConfResponse, err error) {
	if request == nil {
		request = NewDescribeSubscribeConfRequest()
	}
	response = NewDescribeSubscribeConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscribeDBInstancesRequest() (request *DescribeSubscribeDBInstancesRequest) {
	request = &DescribeSubscribeDBInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeDBInstances")
	return
}

func NewDescribeSubscribeDBInstancesResponse() (response *DescribeSubscribeDBInstancesResponse) {
	response = &DescribeSubscribeDBInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSubscribeDBInstances）用于查询支持订阅的云数据库
func (c *Client) DescribeSubscribeDBInstances(request *DescribeSubscribeDBInstancesRequest) (response *DescribeSubscribeDBInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeSubscribeDBInstancesRequest()
	}
	response = NewDescribeSubscribeDBInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewStopMigrateJobRequest() (request *StopMigrateJobRequest) {
	request = &StopMigrateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "StopMigrateJob")
	return
}

func NewStopMigrateJobResponse() (response *StopMigrateJobResponse) {
	response = &StopMigrateJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（StopMigrateJob）用于撤销数据迁移任务。
// 在迁移过程中允许调用该接口撤销迁移, 撤销迁移的任务会失败。通过DescribeMigrateJobs接口查询到任务状态为运行中（status=7）或准备完成（status=8）时，才能撤销数据迁移任务。
func (c *Client) StopMigrateJob(request *StopMigrateJobRequest) (response *StopMigrateJobResponse, err error) {
	if request == nil {
		request = NewStopMigrateJobRequest()
	}
	response = NewStopMigrateJobResponse()
	err = c.Send(request, response)
	return
}

func NewAuthenticateSubscribeSDKRequest() (request *AuthenticateSubscribeSDKRequest) {
	request = &AuthenticateSubscribeSDKRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "AuthenticateSubscribeSDK")
	return
}

func NewAuthenticateSubscribeSDKResponse() (response *AuthenticateSubscribeSDKResponse) {
	response = &AuthenticateSubscribeSDKResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// binlogSDK鉴权
func (c *Client) AuthenticateSubscribeSDK(request *AuthenticateSubscribeSDKRequest) (response *AuthenticateSubscribeSDKResponse, err error) {
	if request == nil {
		request = NewAuthenticateSubscribeSDKRequest()
	}
	response = NewAuthenticateSubscribeSDKResponse()
	err = c.Send(request, response)
	return
}

func NewModifySubscribeVipVportRequest() (request *ModifySubscribeVipVportRequest) {
	request = &ModifySubscribeVipVportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeVipVport")
	return
}

func NewModifySubscribeVipVportResponse() (response *ModifySubscribeVipVportResponse) {
	response = &ModifySubscribeVipVportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifySubscribeVipVport)用于修改数据订阅实例的IP和端口号
func (c *Client) ModifySubscribeVipVport(request *ModifySubscribeVipVportRequest) (response *ModifySubscribeVipVportResponse, err error) {
	if request == nil {
		request = NewModifySubscribeVipVportRequest()
	}
	response = NewModifySubscribeVipVportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMigrateNameRequest() (request *ModifyMigrateNameRequest) {
	request = &ModifyMigrateNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateName")
	return
}

func NewModifyMigrateNameResponse() (response *ModifyMigrateNameResponse) {
	response = &ModifyMigrateNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改迁移任务名称
func (c *Client) ModifyMigrateName(request *ModifyMigrateNameRequest) (response *ModifyMigrateNameResponse, err error) {
	if request == nil {
		request = NewModifyMigrateNameRequest()
	}
	response = NewModifyMigrateNameResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConsumerGroupRequest() (request *DeleteConsumerGroupRequest) {
	request = &DeleteConsumerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DeleteConsumerGroup")
	return
}

func NewDeleteConsumerGroupResponse() (response *DeleteConsumerGroupResponse) {
	response = &DeleteConsumerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DeleteConsumerGroup)用于删除数据订阅消费组。
func (c *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
	if request == nil {
		request = NewDeleteConsumerGroupRequest()
	}
	response = NewDeleteConsumerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJobStatusRequest() (request *DescribeJobStatusRequest) {
	request = &DescribeJobStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeJobStatus")
	return
}

func NewDescribeJobStatusResponse() (response *DescribeJobStatusResponse) {
	response = &DescribeJobStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取迁移任务状态日志列表
func (c *Client) DescribeJobStatus(request *DescribeJobStatusRequest) (response *DescribeJobStatusResponse, err error) {
	if request == nil {
		request = NewDescribeJobStatusRequest()
	}
	response = NewDescribeJobStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRdsSupportRegionRequest() (request *DescribeRdsSupportRegionRequest) {
	request = &DescribeRdsSupportRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeRdsSupportRegion")
	return
}

func NewDescribeRdsSupportRegionResponse() (response *DescribeRdsSupportRegionResponse) {
	response = &DescribeRdsSupportRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持Rds区域信息
func (c *Client) DescribeRdsSupportRegion(request *DescribeRdsSupportRegionRequest) (response *DescribeRdsSupportRegionResponse, err error) {
	if request == nil {
		request = NewDescribeRdsSupportRegionRequest()
	}
	response = NewDescribeRdsSupportRegionResponse()
	err = c.Send(request, response)
	return
}

func NewActivateSubscribeRequest() (request *ActivateSubscribeRequest) {
	request = &ActivateSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ActivateSubscribe")
	return
}

func NewActivateSubscribeResponse() (response *ActivateSubscribeResponse) {
	response = &ActivateSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于配置数据订阅，只有在未配置状态的订阅实例才能调用此接口。
func (c *Client) ActivateSubscribe(request *ActivateSubscribeRequest) (response *ActivateSubscribeResponse, err error) {
	if request == nil {
		request = NewActivateSubscribeRequest()
	}
	response = NewActivateSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMigrateObjectRequest() (request *DescribeMigrateObjectRequest) {
	request = &DescribeMigrateObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateObject")
	return
}

func NewDescribeMigrateObjectResponse() (response *DescribeMigrateObjectResponse) {
	response = &DescribeMigrateObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeMigrateObject
func (c *Client) DescribeMigrateObject(request *DescribeMigrateObjectRequest) (response *DescribeMigrateObjectResponse, err error) {
	if request == nil {
		request = NewDescribeMigrateObjectRequest()
	}
	response = NewDescribeMigrateObjectResponse()
	err = c.Send(request, response)
	return
}

func NewResumeSubscribeRequest() (request *ResumeSubscribeRequest) {
	request = &ResumeSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ResumeSubscribe")
	return
}

func NewResumeSubscribeResponse() (response *ResumeSubscribeResponse) {
	response = &ResumeSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResumeSubscribe）用于恢复数据订阅
func (c *Client) ResumeSubscribe(request *ResumeSubscribeRequest) (response *ResumeSubscribeResponse, err error) {
	if request == nil {
		request = NewResumeSubscribeRequest()
	}
	response = NewResumeSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscribeDBTablesRequest() (request *DescribeSubscribeDBTablesRequest) {
	request = &DescribeSubscribeDBTablesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeDBTables")
	return
}

func NewDescribeSubscribeDBTablesResponse() (response *DescribeSubscribeDBTablesResponse) {
	response = &DescribeSubscribeDBTablesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSubscribeDBTables）用于查询云数据库某个库所含表
func (c *Client) DescribeSubscribeDBTables(request *DescribeSubscribeDBTablesRequest) (response *DescribeSubscribeDBTablesResponse, err error) {
	if request == nil {
		request = NewDescribeSubscribeDBTablesRequest()
	}
	response = NewDescribeSubscribeDBTablesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMigrateJobs4DevOpsRequest() (request *DescribeMigrateJobs4DevOpsRequest) {
	request = &DescribeMigrateJobs4DevOpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateJobs4DevOps")
	return
}

func NewDescribeMigrateJobs4DevOpsResponse() (response *DescribeMigrateJobs4DevOpsResponse) {
	response = &DescribeMigrateJobs4DevOpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询迁移任务列表，用于面板展示所有在运行状态的迁移任务
func (c *Client) DescribeMigrateJobs4DevOps(request *DescribeMigrateJobs4DevOpsRequest) (response *DescribeMigrateJobs4DevOpsResponse, err error) {
	if request == nil {
		request = NewDescribeMigrateJobs4DevOpsRequest()
	}
	response = NewDescribeMigrateJobs4DevOpsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConnectTestResultRequest() (request *DescribeConnectTestResultRequest) {
	request = &DescribeConnectTestResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeConnectTestResult")
	return
}

func NewDescribeConnectTestResultResponse() (response *DescribeConnectTestResultResponse) {
	response = &DescribeConnectTestResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询连接性检查任务结果
func (c *Client) DescribeConnectTestResult(request *DescribeConnectTestResultRequest) (response *DescribeConnectTestResultResponse, err error) {
	if request == nil {
		request = NewDescribeConnectTestResultRequest()
	}
	response = NewDescribeConnectTestResultResponse()
	err = c.Send(request, response)
	return
}

func NewResetSubscribeRequest() (request *ResetSubscribeRequest) {
	request = &ResetSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ResetSubscribe")
	return
}

func NewResetSubscribeResponse() (response *ResetSubscribeResponse) {
	response = &ResetSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ResetSubscribe)用于重置数据订阅实例，已经激活的数据订阅实例，重置后可以使用ActivateSubscribe接口绑定其他的数据库实例
func (c *Client) ResetSubscribe(request *ResetSubscribeRequest) (response *ResetSubscribeResponse, err error) {
	if request == nil {
		request = NewResetSubscribeRequest()
	}
	response = NewResetSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewIsolateSubscribeRequest() (request *IsolateSubscribeRequest) {
	request = &IsolateSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "IsolateSubscribe")
	return
}

func NewIsolateSubscribeResponse() (response *IsolateSubscribeResponse) {
	response = &IsolateSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（IsolateSubscribe）用于隔离小时计费的订阅实例。调用后，订阅实例将不能使用，同时停止计费。
func (c *Client) IsolateSubscribe(request *IsolateSubscribeRequest) (response *IsolateSubscribeResponse, err error) {
	if request == nil {
		request = NewIsolateSubscribeRequest()
	}
	response = NewIsolateSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConsumerGroupDescriptionRequest() (request *ModifyConsumerGroupDescriptionRequest) {
	request = &ModifyConsumerGroupDescriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ModifyConsumerGroupDescription")
	return
}

func NewModifyConsumerGroupDescriptionResponse() (response *ModifyConsumerGroupDescriptionResponse) {
	response = &ModifyConsumerGroupDescriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyConsumerGroupDescription）用于修改消费组描述信息
func (c *Client) ModifyConsumerGroupDescription(request *ModifyConsumerGroupDescriptionRequest) (response *ModifyConsumerGroupDescriptionResponse, err error) {
	if request == nil {
		request = NewModifyConsumerGroupDescriptionRequest()
	}
	response = NewModifyConsumerGroupDescriptionResponse()
	err = c.Send(request, response)
	return
}

func NewCompleteMigrateJobRequest() (request *CompleteMigrateJobRequest) {
	request = &CompleteMigrateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CompleteMigrateJob")
	return
}

func NewCompleteMigrateJobResponse() (response *CompleteMigrateJobResponse) {
	response = &CompleteMigrateJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CompleteMigrateJob）用于完成数据迁移任务。
// 选择采用增量迁移方式的任务, 需要在迁移进度进入准备完成阶段后, 调用本接口, 停止迁移增量数据。
// 通过DescribeMigrateJobs接口查询到任务的状态为准备完成（status=8）时，此时可以调用本接口完成迁移任务。
func (c *Client) CompleteMigrateJob(request *CompleteMigrateJobRequest) (response *CompleteMigrateJobResponse, err error) {
	if request == nil {
		request = NewCompleteMigrateJobRequest()
	}
	response = NewCompleteMigrateJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMigrateJobsRequest() (request *DescribeMigrateJobsRequest) {
	request = &DescribeMigrateJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateJobs")
	return
}

func NewDescribeMigrateJobsResponse() (response *DescribeMigrateJobsResponse) {
	response = &DescribeMigrateJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询数据迁移任务.
// 如果是金融区链路, 请使用域名: https://dts.ap-shenzhen-fsi.tencentcloudapi.com
func (c *Client) DescribeMigrateJobs(request *DescribeMigrateJobsRequest) (response *DescribeMigrateJobsResponse, err error) {
	if request == nil {
		request = NewDescribeMigrateJobsRequest()
	}
	response = NewDescribeMigrateJobsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConnectTestJobRequest() (request *CreateConnectTestJobRequest) {
	request = &CreateConnectTestJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateConnectTestJob")
	return
}

func NewCreateConnectTestJobResponse() (response *CreateConnectTestJobResponse) {
	response = &CreateConnectTestJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于检测源或者目标连接性问题
func (c *Client) CreateConnectTestJob(request *CreateConnectTestJobRequest) (response *CreateConnectTestJobResponse, err error) {
	if request == nil {
		request = NewCreateConnectTestJobRequest()
	}
	response = NewCreateConnectTestJobResponse()
	err = c.Send(request, response)
	return
}

func NewStartSubscribeRequest() (request *StartSubscribeRequest) {
	request = &StartSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "StartSubscribe")
	return
}

func NewStartSubscribeResponse() (response *StartSubscribeResponse) {
	response = &StartSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（StartSubscribe）用于启动数据订阅
func (c *Client) StartSubscribe(request *StartSubscribeRequest) (response *StartSubscribeResponse, err error) {
	if request == nil {
		request = NewStartSubscribeRequest()
	}
	response = NewStartSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewModifySubscribeConsumeTimeRequest() (request *ModifySubscribeConsumeTimeRequest) {
	request = &ModifySubscribeConsumeTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeConsumeTime")
	return
}

func NewModifySubscribeConsumeTimeResponse() (response *ModifySubscribeConsumeTimeResponse) {
	response = &ModifySubscribeConsumeTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifySubscribeConsumeTime)用于修改数据订阅通道的消费时间点
func (c *Client) ModifySubscribeConsumeTime(request *ModifySubscribeConsumeTimeRequest) (response *ModifySubscribeConsumeTimeResponse, err error) {
	if request == nil {
		request = NewModifySubscribeConsumeTimeRequest()
	}
	response = NewModifySubscribeConsumeTimeResponse()
	err = c.Send(request, response)
	return
}

func NewStartMigrateJobRequest() (request *StartMigrateJobRequest) {
	request = &StartMigrateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "StartMigrateJob")
	return
}

func NewStartMigrateJobResponse() (response *StartMigrateJobResponse) {
	response = &StartMigrateJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（StartMigrationJob）用于启动迁移任务。非定时迁移任务会在调用后立即开始迁移，定时任务则会开始倒计时。
// 调用此接口前，请务必先使用CreateMigrateCheckJob校验数据迁移任务，并通过DescribeMigrateJobs接口查询到任务状态为校验通过（status=4）时，才能启动数据迁移任务。
func (c *Client) StartMigrateJob(request *StartMigrateJobRequest) (response *StartMigrateJobResponse, err error) {
	if request == nil {
		request = NewStartMigrateJobRequest()
	}
	response = NewStartMigrateJobResponse()
	err = c.Send(request, response)
	return
}

func NewConfigureSubscribeRequest() (request *ConfigureSubscribeRequest) {
	request = &ConfigureSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ConfigureSubscribe")
	return
}

func NewConfigureSubscribeResponse() (response *ConfigureSubscribeResponse) {
	response = &ConfigureSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ConfigureSubscribe)用于配置数据订阅
func (c *Client) ConfigureSubscribe(request *ConfigureSubscribeRequest) (response *ConfigureSubscribeResponse, err error) {
	if request == nil {
		request = NewConfigureSubscribeRequest()
	}
	response = NewConfigureSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsumerGroupsRequest() (request *DescribeConsumerGroupsRequest) {
	request = &DescribeConsumerGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeConsumerGroups")
	return
}

func NewDescribeConsumerGroupsResponse() (response *DescribeConsumerGroupsResponse) {
	response = &DescribeConsumerGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeConsumerGroups)用于查询数据订阅消费组信息。
func (c *Client) DescribeConsumerGroups(request *DescribeConsumerGroupsRequest) (response *DescribeConsumerGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeConsumerGroupsRequest()
	}
	response = NewDescribeConsumerGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscribeCheckJobRequest() (request *DescribeSubscribeCheckJobRequest) {
	request = &DescribeSubscribeCheckJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeCheckJob")
	return
}

func NewDescribeSubscribeCheckJobResponse() (response *DescribeSubscribeCheckJobResponse) {
	response = &DescribeSubscribeCheckJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSubscribeCheckJob）用于查询异步作业检查结果信息。
func (c *Client) DescribeSubscribeCheckJob(request *DescribeSubscribeCheckJobRequest) (response *DescribeSubscribeCheckJobResponse, err error) {
	if request == nil {
		request = NewDescribeSubscribeCheckJobRequest()
	}
	response = NewDescribeSubscribeCheckJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
	request = &DescribeAsyncRequestInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeAsyncRequestInfo")
	return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
	response = &DescribeAsyncRequestInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAsyncRequestInfo）用于查询任务执行结果
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAsyncRequestInfoRequest()
	}
	response = NewDescribeAsyncRequestInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMigrateJobRequest() (request *DeleteMigrateJobRequest) {
	request = &DeleteMigrateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DeleteMigrateJob")
	return
}

func NewDeleteMigrateJobResponse() (response *DeleteMigrateJobResponse) {
	response = &DeleteMigrateJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteMigrationJob）用于删除数据迁移任务。当通过DescribeMigrateJobs接口查询到任务的状态为：检验中（status=3）、运行中（status=7）、准备完成（status=8）、撤销中（status=11）或者完成中（status=12）时，不允许删除任务。
func (c *Client) DeleteMigrateJob(request *DeleteMigrateJobRequest) (response *DeleteMigrateJobResponse, err error) {
	if request == nil {
		request = NewDeleteMigrateJobRequest()
	}
	response = NewDeleteMigrateJobResponse()
	err = c.Send(request, response)
	return
}

func NewOfflineIsolatedSubscribeRequest() (request *OfflineIsolatedSubscribeRequest) {
	request = &OfflineIsolatedSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "OfflineIsolatedSubscribe")
	return
}

func NewOfflineIsolatedSubscribeResponse() (response *OfflineIsolatedSubscribeResponse) {
	response = &OfflineIsolatedSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（OfflineIsolatedSubscribe）用于下线已隔离的数据订阅实例
func (c *Client) OfflineIsolatedSubscribe(request *OfflineIsolatedSubscribeRequest) (response *OfflineIsolatedSubscribeResponse, err error) {
	if request == nil {
		request = NewOfflineIsolatedSubscribeRequest()
	}
	response = NewOfflineIsolatedSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConsumerGroupRequest() (request *CreateConsumerGroupRequest) {
	request = &CreateConsumerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateConsumerGroup")
	return
}

func NewCreateConsumerGroupResponse() (response *CreateConsumerGroupResponse) {
	response = &CreateConsumerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateConsumerGroup)用于创建一个数据订阅消费组。
func (c *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
	if request == nil {
		request = NewCreateConsumerGroupRequest()
	}
	response = NewCreateConsumerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscribesRequest() (request *DescribeSubscribesRequest) {
	request = &DescribeSubscribesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribes")
	return
}

func NewDescribeSubscribesResponse() (response *DescribeSubscribesResponse) {
	response = &DescribeSubscribesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeSubscribes)获取数据订阅实例信息列表，默认分页，每次返回20条
func (c *Client) DescribeSubscribes(request *DescribeSubscribesRequest) (response *DescribeSubscribesResponse, err error) {
	if request == nil {
		request = NewDescribeSubscribesRequest()
	}
	response = NewDescribeSubscribesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMigrateJobRequest() (request *ModifyMigrateJobRequest) {
	request = &ModifyMigrateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateJob")
	return
}

func NewModifyMigrateJobResponse() (response *ModifyMigrateJobResponse) {
	response = &ModifyMigrateJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyMigrateJob）用于修改数据迁移任务。
// 当迁移任务处于下述状态时，允许调用本接口修改迁移任务：迁移创建中（status=1）、 校验成功(status=4)、校验失败(status=5)、迁移失败(status=10)。但源实例、目标实例类型和目标实例地域不允许修改。
//
// 如果是金融区链路, 请使用域名: dts.ap-shenzhen-fsi.tencentcloudapi.com
func (c *Client) ModifyMigrateJob(request *ModifyMigrateJobRequest) (response *ModifyMigrateJobResponse, err error) {
	if request == nil {
		request = NewModifyMigrateJobRequest()
	}
	response = NewModifyMigrateJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSubscribeCheckJobRequest() (request *CreateSubscribeCheckJobRequest) {
	request = &CreateSubscribeCheckJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateSubscribeCheckJob")
	return
}

func NewCreateSubscribeCheckJobResponse() (response *CreateSubscribeCheckJobResponse) {
	response = &CreateSubscribeCheckJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateSubscribeCheckJob)用于创建一个数据订阅检查作业。
func (c *Client) CreateSubscribeCheckJob(request *CreateSubscribeCheckJobRequest) (response *CreateSubscribeCheckJobResponse, err error) {
	if request == nil {
		request = NewCreateSubscribeCheckJobRequest()
	}
	response = NewCreateSubscribeCheckJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMigrateJobRequest() (request *CreateMigrateJobRequest) {
	request = &CreateMigrateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateJob")
	return
}

func NewCreateMigrateJobResponse() (response *CreateMigrateJobResponse) {
	response = &CreateMigrateJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateMigrateJob）用于创建数据迁移任务。
//
// 如果是金融区链路, 请使用域名: dts.ap-shenzhen-fsi.tencentcloudapi.com
func (c *Client) CreateMigrateJob(request *CreateMigrateJobRequest) (response *CreateMigrateJobResponse, err error) {
	if request == nil {
		request = NewCreateMigrateJobRequest()
	}
	response = NewCreateMigrateJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBSupportRegionRequest() (request *DescribeDBSupportRegionRequest) {
	request = &DescribeDBSupportRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeDBSupportRegion")
	return
}

func NewDescribeDBSupportRegionResponse() (response *DescribeDBSupportRegionResponse) {
	response = &DescribeDBSupportRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持DB区域信息
func (c *Client) DescribeDBSupportRegion(request *DescribeDBSupportRegionRequest) (response *DescribeDBSupportRegionResponse, err error) {
	if request == nil {
		request = NewDescribeDBSupportRegionRequest()
	}
	response = NewDescribeDBSupportRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMigrateCheckJobRequest() (request *DescribeMigrateCheckJobRequest) {
	request = &DescribeMigrateCheckJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateCheckJob")
	return
}

func NewDescribeMigrateCheckJobResponse() (response *DescribeMigrateCheckJobResponse) {
	response = &DescribeMigrateCheckJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建校验后,获取校验的结果. 能查询到当前校验的状态和进度.
// 若通过校验, 则可调用'StartMigrateJob' 开始迁移.
// 若未通过校验, 则能查询到校验失败的原因. 请按照报错, 通过'ModifyMigrateJob'修改迁移配置或是调整源/目标实例的相关参数.
func (c *Client) DescribeMigrateCheckJob(request *DescribeMigrateCheckJobRequest) (response *DescribeMigrateCheckJobResponse, err error) {
	if request == nil {
		request = NewDescribeMigrateCheckJobRequest()
	}
	response = NewDescribeMigrateCheckJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMigrateSurveyRequest() (request *DescribeMigrateSurveyRequest) {
	request = &DescribeMigrateSurveyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateSurvey")
	return
}

func NewDescribeMigrateSurveyResponse() (response *DescribeMigrateSurveyResponse) {
	response = &DescribeMigrateSurveyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 迁移概览
func (c *Client) DescribeMigrateSurvey(request *DescribeMigrateSurveyRequest) (response *DescribeMigrateSurveyResponse, err error) {
	if request == nil {
		request = NewDescribeMigrateSurveyRequest()
	}
	response = NewDescribeMigrateSurveyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMigrateDBInstancesRequest() (request *DescribeMigrateDBInstancesRequest) {
	request = &DescribeMigrateDBInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateDBInstances")
	return
}

func NewDescribeMigrateDBInstancesResponse() (response *DescribeMigrateDBInstancesResponse) {
	response = &DescribeMigrateDBInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取可迁移的实例列表
func (c *Client) DescribeMigrateDBInstances(request *DescribeMigrateDBInstancesRequest) (response *DescribeMigrateDBInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeMigrateDBInstancesRequest()
	}
	response = NewDescribeMigrateDBInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProxySupportRegionRequest() (request *DescribeProxySupportRegionRequest) {
	request = &DescribeProxySupportRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeProxySupportRegion")
	return
}

func NewDescribeProxySupportRegionResponse() (response *DescribeProxySupportRegionResponse) {
	response = &DescribeProxySupportRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持Proxy区域信息
func (c *Client) DescribeProxySupportRegion(request *DescribeProxySupportRegionRequest) (response *DescribeProxySupportRegionResponse, err error) {
	if request == nil {
		request = NewDescribeProxySupportRegionRequest()
	}
	response = NewDescribeProxySupportRegionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSubscribeRequest() (request *CreateSubscribeRequest) {
	request = &CreateSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "CreateSubscribe")
	return
}

func NewCreateSubscribeResponse() (response *CreateSubscribeResponse) {
	response = &CreateSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateSubscribe)用于创建一个数据订阅实例。
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
	if request == nil {
		request = NewCreateSubscribeRequest()
	}
	response = NewCreateSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscribeDBDatabasesRequest() (request *DescribeSubscribeDBDatabasesRequest) {
	request = &DescribeSubscribeDBDatabasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeDBDatabases")
	return
}

func NewDescribeSubscribeDBDatabasesResponse() (response *DescribeSubscribeDBDatabasesResponse) {
	response = &DescribeSubscribeDBDatabasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSubscribeDBDatabases）用于查询云数据库实例的数据库
func (c *Client) DescribeSubscribeDBDatabases(request *DescribeSubscribeDBDatabasesRequest) (response *DescribeSubscribeDBDatabasesResponse, err error) {
	if request == nil {
		request = NewDescribeSubscribeDBDatabasesRequest()
	}
	response = NewDescribeSubscribeDBDatabasesResponse()
	err = c.Send(request, response)
	return
}
