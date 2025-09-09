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

package v20191112

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-11-12"

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

func NewDescribeVsmsRequest() (request *DescribeVsmsRequest) {
	request = &DescribeVsmsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeVsms")
	return
}

func NewDescribeVsmsResponse() (response *DescribeVsmsResponse) {
	response = &DescribeVsmsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户VSM列表
func (c *Client) DescribeVsms(request *DescribeVsmsRequest) (response *DescribeVsmsResponse, err error) {
	if request == nil {
		request = NewDescribeVsmsRequest()
	}
	response = NewDescribeVsmsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateResourceRequest() (request *CreateResourceRequest) {
	request = &CreateResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "CreateResource")
	return
}

func NewCreateResourceResponse() (response *CreateResourceResponse) {
	response = &CreateResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建VSM
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
	if request == nil {
		request = NewCreateResourceRequest()
	}
	response = NewCreateResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVsmPropertyRequest() (request *DescribeVsmPropertyRequest) {
	request = &DescribeVsmPropertyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeVsmProperty")
	return
}

func NewDescribeVsmPropertyResponse() (response *DescribeVsmPropertyResponse) {
	response = &DescribeVsmPropertyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VSM参数信息列表
func (c *Client) DescribeVsmProperty(request *DescribeVsmPropertyRequest) (response *DescribeVsmPropertyResponse, err error) {
	if request == nil {
		request = NewDescribeVsmPropertyRequest()
	}
	response = NewDescribeVsmPropertyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsgRequest() (request *DescribeUsgRequest) {
	request = &DescribeUsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeUsg")
	return
}

func NewDescribeUsgResponse() (response *DescribeUsgResponse) {
	response = &DescribeUsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据用户的AppId获取用户安全组列表
func (c *Client) DescribeUsg(request *DescribeUsgRequest) (response *DescribeUsgResponse, err error) {
	if request == nil {
		request = NewDescribeUsgRequest()
	}
	response = NewDescribeUsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSupportedHsmRequest() (request *DescribeSupportedHsmRequest) {
	request = &DescribeSupportedHsmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeSupportedHsm")
	return
}

func NewDescribeSupportedHsmResponse() (response *DescribeSupportedHsmResponse) {
	response = &DescribeSupportedHsmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前地域所支持的设备列表
func (c *Client) DescribeSupportedHsm(request *DescribeSupportedHsmRequest) (response *DescribeSupportedHsmResponse, err error) {
	if request == nil {
		request = NewDescribeSupportedHsmRequest()
	}
	response = NewDescribeSupportedHsmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHSMBySubnetIdRequest() (request *DescribeHSMBySubnetIdRequest) {
	request = &DescribeHSMBySubnetIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeHSMBySubnetId")
	return
}

func NewDescribeHSMBySubnetIdResponse() (response *DescribeHSMBySubnetIdResponse) {
	response = &DescribeHSMBySubnetIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过SubnetId获取Hsm资源数
func (c *Client) DescribeHSMBySubnetId(request *DescribeHSMBySubnetIdRequest) (response *DescribeHSMBySubnetIdResponse, err error) {
	if request == nil {
		request = NewDescribeHSMBySubnetIdRequest()
	}
	response = NewDescribeHSMBySubnetIdResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateVsmRequest() (request *TerminateVsmRequest) {
	request = &TerminateVsmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "TerminateVsm")
	return
}

func NewTerminateVsmResponse() (response *TerminateVsmResponse) {
	response = &TerminateVsmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 退还虚拟加密机
func (c *Client) TerminateVsm(request *TerminateVsmRequest) (response *TerminateVsmResponse, err error) {
	if request == nil {
		request = NewTerminateVsmRequest()
	}
	response = NewTerminateVsmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeRegions")
	return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取地域列表
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	response = NewDescribeRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVsmAttributesRequest() (request *ModifyVsmAttributesRequest) {
	request = &ModifyVsmAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "ModifyVsmAttributes")
	return
}

func NewModifyVsmAttributesResponse() (response *ModifyVsmAttributesResponse) {
	response = &ModifyVsmAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改VSM属性
func (c *Client) ModifyVsmAttributes(request *ModifyVsmAttributesRequest) (response *ModifyVsmAttributesResponse, err error) {
	if request == nil {
		request = NewModifyVsmAttributesRequest()
	}
	response = NewModifyVsmAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVsmAttributesRequest() (request *DescribeVsmAttributesRequest) {
	request = &DescribeVsmAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeVsmAttributes")
	return
}

func NewDescribeVsmAttributesResponse() (response *DescribeVsmAttributesResponse) {
	response = &DescribeVsmAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VSM属性
func (c *Client) DescribeVsmAttributes(request *DescribeVsmAttributesRequest) (response *DescribeVsmAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeVsmAttributesRequest()
	}
	response = NewDescribeVsmAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubnetRequest() (request *DescribeSubnetRequest) {
	request = &DescribeSubnetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeSubnet")
	return
}

func NewDescribeSubnetResponse() (response *DescribeSubnetResponse) {
	response = &DescribeSubnetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子网列表
func (c *Client) DescribeSubnet(request *DescribeSubnetRequest) (response *DescribeSubnetResponse, err error) {
	if request == nil {
		request = NewDescribeSubnetRequest()
	}
	response = NewDescribeSubnetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsgRuleRequest() (request *DescribeUsgRuleRequest) {
	request = &DescribeUsgRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeUsgRule")
	return
}

func NewDescribeUsgRuleResponse() (response *DescribeUsgRuleResponse) {
	response = &DescribeUsgRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全组详情
func (c *Client) DescribeUsgRule(request *DescribeUsgRuleRequest) (response *DescribeUsgRuleResponse, err error) {
	if request == nil {
		request = NewDescribeUsgRuleRequest()
	}
	response = NewDescribeUsgRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExtraFlagsRequest() (request *DescribeExtraFlagsRequest) {
	request = &DescribeExtraFlagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeExtraFlags")
	return
}

func NewDescribeExtraFlagsResponse() (response *DescribeExtraFlagsResponse) {
	response = &DescribeExtraFlagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取辅助参数
func (c *Client) DescribeExtraFlags(request *DescribeExtraFlagsRequest) (response *DescribeExtraFlagsResponse, err error) {
	if request == nil {
		request = NewDescribeExtraFlagsRequest()
	}
	response = NewDescribeExtraFlagsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceBuyVsmRequest() (request *InquiryPriceBuyVsmRequest) {
	request = &InquiryPriceBuyVsmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "InquiryPriceBuyVsm")
	return
}

func NewInquiryPriceBuyVsmResponse() (response *InquiryPriceBuyVsmResponse) {
	response = &InquiryPriceBuyVsmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 购买询价接口
func (c *Client) InquiryPriceBuyVsm(request *InquiryPriceBuyVsmRequest) (response *InquiryPriceBuyVsmResponse, err error) {
	if request == nil {
		request = NewInquiryPriceBuyVsmRequest()
	}
	response = NewInquiryPriceBuyVsmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHSMByVpcIdRequest() (request *DescribeHSMByVpcIdRequest) {
	request = &DescribeHSMByVpcIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeHSMByVpcId")
	return
}

func NewDescribeHSMByVpcIdResponse() (response *DescribeHSMByVpcIdResponse) {
	response = &DescribeHSMByVpcIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过VpcId获取Hsm资源数
func (c *Client) DescribeHSMByVpcId(request *DescribeHSMByVpcIdRequest) (response *DescribeHSMByVpcIdResponse, err error) {
	if request == nil {
		request = NewDescribeHSMByVpcIdRequest()
	}
	response = NewDescribeHSMByVpcIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcRequest() (request *DescribeVpcRequest) {
	request = &DescribeVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeVpc")
	return
}

func NewDescribeVpcResponse() (response *DescribeVpcResponse) {
	response = &DescribeVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户的私有网络列表
func (c *Client) DescribeVpc(request *DescribeVpcRequest) (response *DescribeVpcResponse, err error) {
	if request == nil {
		request = NewDescribeVpcRequest()
	}
	response = NewDescribeVpcResponse()
	err = c.Send(request, response)
	return
}
