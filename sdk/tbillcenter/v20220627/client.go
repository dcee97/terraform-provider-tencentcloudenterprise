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

package v20220627

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2022-06-27"

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

func NewQueryCustomDiscountRequest() (request *QueryCustomDiscountRequest) {
	request = &QueryCustomDiscountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbillcenter", APIVersion, "QueryCustomDiscount")
	return
}

func NewQueryCustomDiscountResponse() (response *QueryCustomDiscountResponse) {
	response = &QueryCustomDiscountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户关联折扣，只查询高优折扣
func (c *Client) QueryCustomDiscount(request *QueryCustomDiscountRequest) (response *QueryCustomDiscountResponse, err error) {
	if request == nil {
		request = NewQueryCustomDiscountRequest()
	}
	response = NewQueryCustomDiscountResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductConfigRequest() (request *QueryProductConfigRequest) {
	request = &QueryProductConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbillcenter", APIVersion, "QueryProductConfig")
	return
}

func NewQueryProductConfigResponse() (response *QueryProductConfigResponse) {
	response = &QueryProductConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品配置
func (c *Client) QueryProductConfig(request *QueryProductConfigRequest) (response *QueryProductConfigResponse, err error) {
	if request == nil {
		request = NewQueryProductConfigRequest()
	}
	response = NewQueryProductConfigResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBalanceRequest() (request *QueryBalanceRequest) {
	request = &QueryBalanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbillcenter", APIVersion, "QueryBalance")
	return
}

func NewQueryBalanceResponse() (response *QueryBalanceResponse) {
	response = &QueryBalanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账户余额
func (c *Client) QueryBalance(request *QueryBalanceRequest) (response *QueryBalanceResponse, err error) {
	if request == nil {
		request = NewQueryBalanceRequest()
	}
	response = NewQueryBalanceResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductPriceRequest() (request *QueryProductPriceRequest) {
	request = &QueryProductPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbillcenter", APIVersion, "QueryProductPrice")
	return
}

func NewQueryProductPriceResponse() (response *QueryProductPriceResponse) {
	response = &QueryProductPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品定价，返回结果照运营端同名接口
func (c *Client) QueryProductPrice(request *QueryProductPriceRequest) (response *QueryProductPriceResponse, err error) {
	if request == nil {
		request = NewQueryProductPriceRequest()
	}
	response = NewQueryProductPriceResponse()
	err = c.Send(request, response)
	return
}

func NewQuerySubBillingItemsRequest() (request *QuerySubBillingItemsRequest) {
	request = &QuerySubBillingItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbillcenter", APIVersion, "QuerySubBillingItems")
	return
}

func NewQuerySubBillingItemsResponse() (response *QuerySubBillingItemsResponse) {
	response = &QuerySubBillingItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据二层编码查询三四层树，返回结果照运营端同名接口
func (c *Client) QuerySubBillingItems(request *QuerySubBillingItemsRequest) (response *QuerySubBillingItemsResponse, err error) {
	if request == nil {
		request = NewQuerySubBillingItemsRequest()
	}
	response = NewQuerySubBillingItemsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryAccountDescRequest() (request *QueryAccountDescRequest) {
	request = &QueryAccountDescRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbillcenter", APIVersion, "QueryAccountDesc")
	return
}

func NewQueryAccountDescResponse() (response *QueryAccountDescResponse) {
	response = &QueryAccountDescResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账户说明
func (c *Client) QueryAccountDesc(request *QueryAccountDescRequest) (response *QueryAccountDescResponse, err error) {
	if request == nil {
		request = NewQueryAccountDescRequest()
	}
	response = NewQueryAccountDescResponse()
	err = c.Send(request, response)
	return
}

func NewQueryInstructionRequest() (request *QueryInstructionRequest) {
	request = &QueryInstructionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbillcenter", APIVersion, "QueryInstruction")
	return
}

func NewQueryInstructionResponse() (response *QueryInstructionResponse) {
	response = &QueryInstructionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询计费说明
func (c *Client) QueryInstruction(request *QueryInstructionRequest) (response *QueryInstructionResponse, err error) {
	if request == nil {
		request = NewQueryInstructionRequest()
	}
	response = NewQueryInstructionResponse()
	err = c.Send(request, response)
	return
}

func NewQuerySubProductConsoleRequest() (request *QuerySubProductConsoleRequest) {
	request = &QuerySubProductConsoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbillcenter", APIVersion, "QuerySubProductConsole")
	return
}

func NewQuerySubProductConsoleResponse() (response *QuerySubProductConsoleResponse) {
	response = &QuerySubProductConsoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户端支持询价一二层树
func (c *Client) QuerySubProductConsole(request *QuerySubProductConsoleRequest) (response *QuerySubProductConsoleResponse, err error) {
	if request == nil {
		request = NewQuerySubProductConsoleRequest()
	}
	response = NewQuerySubProductConsoleResponse()
	err = c.Send(request, response)
	return
}
