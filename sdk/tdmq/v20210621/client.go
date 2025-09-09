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

package v20210621

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2021-06-21"

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

func NewGetPriceRequest() (request *GetPriceRequest) {
	request = &GetPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "GetPrice")
	return
}

func NewGetPriceResponse() (response *GetPriceResponse) {
	response = &GetPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TDMQ计费询价接口
func (c *Client) GetPrice(request *GetPriceRequest) (response *GetPriceResponse, err error) {
	if request == nil {
		request = NewGetPriceRequest()
	}
	response = NewGetPriceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateResourceRuleOptRequest() (request *UpdateResourceRuleOptRequest) {
	request = &UpdateResourceRuleOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdateResourceRuleOpt")
	return
}

func NewUpdateResourceRuleOptResponse() (response *UpdateResourceRuleOptResponse) {
	response = &UpdateResourceRuleOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端创建集群物理资源分配规则
func (c *Client) UpdateResourceRuleOpt(request *UpdateResourceRuleOptRequest) (response *UpdateResourceRuleOptResponse, err error) {
	if request == nil {
		request = NewUpdateResourceRuleOptRequest()
	}
	response = NewUpdateResourceRuleOptResponse()
	err = c.Send(request, response)
	return
}
