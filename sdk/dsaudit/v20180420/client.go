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

package v20180420

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-04-20"

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
	request.Init().WithApiInfo("dsaudit", APIVersion, "TerminateInstances")
	return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
	response = &TerminateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁、退还数审实例
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
	if request == nil {
		request = NewTerminateInstancesRequest()
	}
	response = NewTerminateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewSetNameInstancesRequest() (request *SetNameInstancesRequest) {
	request = &SetNameInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dsaudit", APIVersion, "SetNameInstances")
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

func NewDescribeProductTypesRequest() (request *DescribeProductTypesRequest) {
	request = &DescribeProductTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dsaudit", APIVersion, "DescribeProductTypes")
	return
}

func NewDescribeProductTypesResponse() (response *DescribeProductTypesResponse) {
	response = &DescribeProductTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取数据安全审计产品规格列表
func (c *Client) DescribeProductTypes(request *DescribeProductTypesRequest) (response *DescribeProductTypesResponse, err error) {
	if request == nil {
		request = NewDescribeProductTypesRequest()
	}
	response = NewDescribeProductTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSaleInfoRequest() (request *DescribeSaleInfoRequest) {
	request = &DescribeSaleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dsaudit", APIVersion, "DescribeSaleInfo")
	return
}

func NewDescribeSaleInfoResponse() (response *DescribeSaleInfoResponse) {
	response = &DescribeSaleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品的销售信息，规格，时长
func (c *Client) DescribeSaleInfo(request *DescribeSaleInfoRequest) (response *DescribeSaleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSaleInfoRequest()
	}
	response = NewDescribeSaleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dsaudit", APIVersion, "CreateInstance")
	return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建数据安全审计实例
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
	if request == nil {
		request = NewCreateInstanceRequest()
	}
	response = NewCreateInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dsaudit", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例列表
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceAddressRequest() (request *ModifyInstanceAddressRequest) {
	request = &ModifyInstanceAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dsaudit", APIVersion, "ModifyInstanceAddress")
	return
}

func NewModifyInstanceAddressResponse() (response *ModifyInstanceAddressResponse) {
	response = &ModifyInstanceAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例网络地址
func (c *Client) ModifyInstanceAddress(request *ModifyInstanceAddressRequest) (response *ModifyInstanceAddressResponse, err error) {
	if request == nil {
		request = NewModifyInstanceAddressRequest()
	}
	response = NewModifyInstanceAddressResponse()
	err = c.Send(request, response)
	return
}

func NewStartInstanceRequest() (request *StartInstanceRequest) {
	request = &StartInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dsaudit", APIVersion, "StartInstance")
	return
}

func NewStartInstanceResponse() (response *StartInstanceResponse) {
	response = &StartInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动实例
func (c *Client) StartInstance(request *StartInstanceRequest) (response *StartInstanceResponse, err error) {
	if request == nil {
		request = NewStartInstanceRequest()
	}
	response = NewStartInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePriceRequest() (request *DescribePriceRequest) {
	request = &DescribePriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dsaudit", APIVersion, "DescribePrice")
	return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询价格
func (c *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
	if request == nil {
		request = NewDescribePriceRequest()
	}
	response = NewDescribePriceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegisterUrlRequest() (request *DescribeRegisterUrlRequest) {
	request = &DescribeRegisterUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dsaudit", APIVersion, "DescribeRegisterUrl")
	return
}

func NewDescribeRegisterUrlResponse() (response *DescribeRegisterUrlResponse) {
	response = &DescribeRegisterUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取注册链接
func (c *Client) DescribeRegisterUrl(request *DescribeRegisterUrlRequest) (response *DescribeRegisterUrlResponse, err error) {
	if request == nil {
		request = NewDescribeRegisterUrlRequest()
	}
	response = NewDescribeRegisterUrlResponse()
	err = c.Send(request, response)
	return
}
