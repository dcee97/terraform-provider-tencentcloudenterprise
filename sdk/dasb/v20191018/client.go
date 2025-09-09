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

package v20191018

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-10-18"

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

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
	request = &TerminateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "TerminateInstances")
	return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
	response = &TerminateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁、退还堡垒机实例
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
	if request == nil {
		request = NewTerminateInstancesRequest()
	}
	response = NewTerminateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDasbCvmConfigRequest() (request *DescribeDasbCvmConfigRequest) {
	request = &DescribeDasbCvmConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbCvmConfig")
	return
}

func NewDescribeDasbCvmConfigResponse() (response *DescribeDasbCvmConfigResponse) {
	response = &DescribeDasbCvmConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询CVM配置信息
func (c *Client) DescribeDasbCvmConfig(request *DescribeDasbCvmConfigRequest) (response *DescribeDasbCvmConfigResponse, err error) {
	if request == nil {
		request = NewDescribeDasbCvmConfigRequest()
	}
	response = NewDescribeDasbCvmConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDasbCvmInstanceRequest() (request *DescribeDasbCvmInstanceRequest) {
	request = &DescribeDasbCvmInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbCvmInstance")
	return
}

func NewDescribeDasbCvmInstanceResponse() (response *DescribeDasbCvmInstanceResponse) {
	response = &DescribeDasbCvmInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询DASB产品对应的CVM实例信息
func (c *Client) DescribeDasbCvmInstance(request *DescribeDasbCvmInstanceRequest) (response *DescribeDasbCvmInstanceResponse, err error) {
	if request == nil {
		request = NewDescribeDasbCvmInstanceRequest()
	}
	response = NewDescribeDasbCvmInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewGetDasbUserInfosRequest() (request *GetDasbUserInfosRequest) {
	request = &GetDasbUserInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "GetDasbUserInfos")
	return
}

func NewGetDasbUserInfosResponse() (response *GetDasbUserInfosResponse) {
	response = &GetDasbUserInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户信息接口
func (c *Client) GetDasbUserInfos(request *GetDasbUserInfosRequest) (response *GetDasbUserInfosResponse, err error) {
	if request == nil {
		request = NewGetDasbUserInfosRequest()
	}
	response = NewGetDasbUserInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDasbServiceRequest() (request *DescribeDasbServiceRequest) {
	request = &DescribeDasbServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbService")
	return
}

func NewDescribeDasbServiceResponse() (response *DescribeDasbServiceResponse) {
	response = &DescribeDasbServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询DASB服务列表
func (c *Client) DescribeDasbService(request *DescribeDasbServiceRequest) (response *DescribeDasbServiceResponse, err error) {
	if request == nil {
		request = NewDescribeDasbServiceRequest()
	}
	response = NewDescribeDasbServiceResponse()
	err = c.Send(request, response)
	return
}

func NewSetDasbSecretInfosRequest() (request *SetDasbSecretInfosRequest) {
	request = &SetDasbSecretInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "SetDasbSecretInfos")
	return
}

func NewSetDasbSecretInfosResponse() (response *SetDasbSecretInfosResponse) {
	response = &SetDasbSecretInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DASB设置密钥信息接口
func (c *Client) SetDasbSecretInfos(request *SetDasbSecretInfosRequest) (response *SetDasbSecretInfosResponse, err error) {
	if request == nil {
		request = NewSetDasbSecretInfosRequest()
	}
	response = NewSetDasbSecretInfosResponse()
	err = c.Send(request, response)
	return
}

func NewSendMessageRequest() (request *SendMessageRequest) {
	request = &SendMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "SendMessage")
	return
}

func NewSendMessageResponse() (response *SendMessageResponse) {
	response = &SendMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送短信
func (c *Client) SendMessage(request *SendMessageRequest) (response *SendMessageResponse, err error) {
	if request == nil {
		request = NewSendMessageRequest()
	}
	response = NewSendMessageResponse()
	err = c.Send(request, response)
	return
}

func NewSetNameInstancesRequest() (request *SetNameInstancesRequest) {
	request = &SetNameInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "SetNameInstances")
	return
}

func NewSetNameInstancesResponse() (response *SetNameInstancesResponse) {
	response = &SetNameInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例命名
func (c *Client) SetNameInstances(request *SetNameInstancesRequest) (response *SetNameInstancesResponse, err error) {
	if request == nil {
		request = NewSetNameInstancesRequest()
	}
	response = NewSetNameInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDasbSystemVersionRequest() (request *DescribeDasbSystemVersionRequest) {
	request = &DescribeDasbSystemVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbSystemVersion")
	return
}

func NewDescribeDasbSystemVersionResponse() (response *DescribeDasbSystemVersionResponse) {
	response = &DescribeDasbSystemVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取系统版本列表
func (c *Client) DescribeDasbSystemVersion(request *DescribeDasbSystemVersionRequest) (response *DescribeDasbSystemVersionResponse, err error) {
	if request == nil {
		request = NewDescribeDasbSystemVersionRequest()
	}
	response = NewDescribeDasbSystemVersionResponse()
	err = c.Send(request, response)
	return
}

func NewRunInstanceRequest() (request *RunInstanceRequest) {
	request = &RunInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "RunInstance")
	return
}

func NewRunInstanceResponse() (response *RunInstanceResponse) {
	response = &RunInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建堡垒机实例
func (c *Client) RunInstance(request *RunInstanceRequest) (response *RunInstanceResponse, err error) {
	if request == nil {
		request = NewRunInstanceRequest()
	}
	response = NewRunInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDasbImageIdsRequest() (request *DescribeDasbImageIdsRequest) {
	request = &DescribeDasbImageIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbImageIds")
	return
}

func NewDescribeDasbImageIdsResponse() (response *DescribeDasbImageIdsResponse) {
	response = &DescribeDasbImageIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像列表
func (c *Client) DescribeDasbImageIds(request *DescribeDasbImageIdsRequest) (response *DescribeDasbImageIdsResponse, err error) {
	if request == nil {
		request = NewDescribeDasbImageIdsRequest()
	}
	response = NewDescribeDasbImageIdsResponse()
	err = c.Send(request, response)
	return
}

func NewInitDasbInstanceRequest() (request *InitDasbInstanceRequest) {
	request = &InitDasbInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "InitDasbInstance")
	return
}

func NewInitDasbInstanceResponse() (response *InitDasbInstanceResponse) {
	response = &InitDasbInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 初始化用户购买的实例
func (c *Client) InitDasbInstance(request *InitDasbInstanceRequest) (response *InitDasbInstanceResponse, err error) {
	if request == nil {
		request = NewInitDasbInstanceRequest()
	}
	response = NewInitDasbInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewAsyncInstanceRequest() (request *AsyncInstanceRequest) {
	request = &AsyncInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "AsyncInstance")
	return
}

func NewAsyncInstanceResponse() (response *AsyncInstanceResponse) {
	response = &AsyncInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步堡垒机实例
func (c *Client) AsyncInstance(request *AsyncInstanceRequest) (response *AsyncInstanceResponse, err error) {
	if request == nil {
		request = NewAsyncInstanceRequest()
	}
	response = NewAsyncInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDasbSystemConfigRequest() (request *DescribeDasbSystemConfigRequest) {
	request = &DescribeDasbSystemConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbSystemConfig")
	return
}

func NewDescribeDasbSystemConfigResponse() (response *DescribeDasbSystemConfigResponse) {
	response = &DescribeDasbSystemConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取系统配置
func (c *Client) DescribeDasbSystemConfig(request *DescribeDasbSystemConfigRequest) (response *DescribeDasbSystemConfigResponse, err error) {
	if request == nil {
		request = NewDescribeDasbSystemConfigRequest()
	}
	response = NewDescribeDasbSystemConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDasbInstanceAlarmRequest() (request *DescribeDasbInstanceAlarmRequest) {
	request = &DescribeDasbInstanceAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dasb", APIVersion, "DescribeDasbInstanceAlarm")
	return
}

func NewDescribeDasbInstanceAlarmResponse() (response *DescribeDasbInstanceAlarmResponse) {
	response = &DescribeDasbInstanceAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DASB实例告警列表
func (c *Client) DescribeDasbInstanceAlarm(request *DescribeDasbInstanceAlarmRequest) (response *DescribeDasbInstanceAlarmResponse, err error) {
	if request == nil {
		request = NewDescribeDasbInstanceAlarmRequest()
	}
	response = NewDescribeDasbInstanceAlarmResponse()
	err = c.Send(request, response)
	return
}
