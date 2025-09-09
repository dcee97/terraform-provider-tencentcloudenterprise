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

package v20180808

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-08-08"

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

func NewModifySubDomainRequest() (request *ModifySubDomainRequest) {
	request = &ModifySubDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifySubDomain")
	return
}

func NewModifySubDomainResponse() (response *ModifySubDomainResponse) {
	response = &ModifySubDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifySubDomain）用于修改服务的自定义域名设置中的路径映射，可以修改绑定自定义域名之前的路径映射规则。
func (c *Client) ModifySubDomain(request *ModifySubDomainRequest) (response *ModifySubDomainResponse, err error) {
	if request == nil {
		request = NewModifySubDomainRequest()
	}
	response = NewModifySubDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTcbScfApiRequest() (request *DeleteTcbScfApiRequest) {
	request = &DeleteTcbScfApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteTcbScfApi")
	return
}

func NewDeleteTcbScfApiResponse() (response *DeleteTcbScfApiResponse) {
	response = &DeleteTcbScfApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于删除TCB-SCF-HTTP触发器
func (c *Client) DeleteTcbScfApi(request *DeleteTcbScfApiRequest) (response *DeleteTcbScfApiResponse, err error) {
	if request == nil {
		request = NewDeleteTcbScfApiRequest()
	}
	response = NewDeleteTcbScfApiResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiAppRequest() (request *DescribeApiAppRequest) {
	request = &DescribeApiAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiApp")
	return
}

func NewDescribeApiAppResponse() (response *DescribeApiAppResponse) {
	response = &DescribeApiAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApiApp）用于根据应用ID搜索应用。
func (c *Client) DescribeApiApp(request *DescribeApiAppRequest) (response *DescribeApiAppResponse, err error) {
	if request == nil {
		request = NewDescribeApiAppRequest()
	}
	response = NewDescribeApiAppResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePluginRequest() (request *CreatePluginRequest) {
	request = &CreatePluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreatePlugin")
	return
}

func NewCreatePluginResponse() (response *CreatePluginResponse) {
	response = &CreatePluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建API网关插件。
func (c *Client) CreatePlugin(request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
	if request == nil {
		request = NewCreatePluginRequest()
	}
	response = NewCreatePluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiKeyRequest() (request *DescribeApiKeyRequest) {
	request = &DescribeApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiKey")
	return
}

func NewDescribeApiKeyResponse() (response *DescribeApiKeyResponse) {
	response = &DescribeApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApiKey）用于查询密钥详情。
// 用户在创建密钥后，可用此接口查询一个 API 密钥的详情，该接口会显示密钥 Key。
func (c *Client) DescribeApiKey(request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
	if request == nil {
		request = NewDescribeApiKeyRequest()
	}
	response = NewDescribeApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewImportOpenApiRequest() (request *ImportOpenApiRequest) {
	request = &ImportOpenApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ImportOpenApi")
	return
}

func NewImportOpenApiResponse() (response *ImportOpenApiResponse) {
	response = &ImportOpenApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ImportOpenApi）用于将OpenAPI规范定义的API导入到API网关。
func (c *Client) ImportOpenApi(request *ImportOpenApiRequest) (response *ImportOpenApiResponse, err error) {
	if request == nil {
		request = NewImportOpenApiRequest()
	}
	response = NewImportOpenApiResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPluginRequest() (request *ModifyPluginRequest) {
	request = &ModifyPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyPlugin")
	return
}

func NewModifyPluginResponse() (response *ModifyPluginResponse) {
	response = &ModifyPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改API网关插件。
func (c *Client) ModifyPlugin(request *ModifyPluginRequest) (response *ModifyPluginResponse, err error) {
	if request == nil {
		request = NewModifyPluginRequest()
	}
	response = NewModifyPluginResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUpstreamRequest() (request *CreateUpstreamRequest) {
	request = &CreateUpstreamRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateUpstream")
	return
}

func NewCreateUpstreamResponse() (response *CreateUpstreamResponse) {
	response = &CreateUpstreamResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于创建创建VPC通道
func (c *Client) CreateUpstream(request *CreateUpstreamRequest) (response *CreateUpstreamResponse, err error) {
	if request == nil {
		request = NewCreateUpstreamRequest()
	}
	response = NewCreateUpstreamResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsagePlansStatusRequest() (request *DescribeUsagePlansStatusRequest) {
	request = &DescribeUsagePlansStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlansStatus")
	return
}

func NewDescribeUsagePlansStatusResponse() (response *DescribeUsagePlansStatusResponse) {
	response = &DescribeUsagePlansStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeUsagePlanStatus）用于查询使用计划的列表。
func (c *Client) DescribeUsagePlansStatus(request *DescribeUsagePlansStatusRequest) (response *DescribeUsagePlansStatusResponse, err error) {
	if request == nil {
		request = NewDescribeUsagePlansStatusRequest()
	}
	response = NewDescribeUsagePlansStatusResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRequest() (request *InquiryPriceRequest) {
	request = &InquiryPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "InquiryPrice")
	return
}

func NewInquiryPriceResponse() (response *InquiryPriceResponse) {
	response = &InquiryPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建apigateway时，用于询价。
func (c *Client) InquiryPrice(request *InquiryPriceRequest) (response *InquiryPriceResponse, err error) {
	if request == nil {
		request = NewInquiryPriceRequest()
	}
	response = NewInquiryPriceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailablePluginApisRequest() (request *DescribeAvailablePluginApisRequest) {
	request = &DescribeAvailablePluginApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAvailablePluginApis")
	return
}

func NewDescribeAvailablePluginApisResponse() (response *DescribeAvailablePluginApisResponse) {
	response = &DescribeAvailablePluginApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示插件可绑定的API列表。
func (c *Client) DescribeAvailablePluginApis(request *DescribeAvailablePluginApisRequest) (response *DescribeAvailablePluginApisResponse, err error) {
	if request == nil {
		request = NewDescribeAvailablePluginApisRequest()
	}
	response = NewDescribeAvailablePluginApisResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceForApiAppRequest() (request *DescribeServiceForApiAppRequest) {
	request = &DescribeServiceForApiAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceForApiApp")
	return
}

func NewDescribeServiceForApiAppResponse() (response *DescribeServiceForApiAppResponse) {
	response = &DescribeServiceForApiAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceForApiApp）用于应用使用者查询一个服务的详细信息、包括服务的描述、域名、协议等信息。
func (c *Client) DescribeServiceForApiApp(request *DescribeServiceForApiAppRequest) (response *DescribeServiceForApiAppResponse, err error) {
	if request == nil {
		request = NewDescribeServiceForApiAppRequest()
	}
	response = NewDescribeServiceForApiAppResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePluginRequest() (request *DescribePluginRequest) {
	request = &DescribePluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribePlugin")
	return
}

func NewDescribePluginResponse() (response *DescribePluginResponse) {
	response = &DescribePluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示插件详情，支持按照插件ID进行。
func (c *Client) DescribePlugin(request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
	if request == nil {
		request = NewDescribePluginRequest()
	}
	response = NewDescribePluginResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiEnvironmentStrategyRequest() (request *ModifyApiEnvironmentStrategyRequest) {
	request = &ModifyApiEnvironmentStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiEnvironmentStrategy")
	return
}

func NewModifyApiEnvironmentStrategyResponse() (response *ModifyApiEnvironmentStrategyResponse) {
	response = &ModifyApiEnvironmentStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyApiEnvironmentStrategy）用于修改API限流策略
func (c *Client) ModifyApiEnvironmentStrategy(request *ModifyApiEnvironmentStrategyRequest) (response *ModifyApiEnvironmentStrategyResponse, err error) {
	if request == nil {
		request = NewModifyApiEnvironmentStrategyRequest()
	}
	response = NewModifyApiEnvironmentStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUsagePlanRequest() (request *CreateUsagePlanRequest) {
	request = &CreateUsagePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateUsagePlan")
	return
}

func NewCreateUsagePlanResponse() (response *CreateUsagePlanResponse) {
	response = &CreateUsagePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateUsagePlan）用于创建使用计划。
// 用户在使用 API 网关时，需要创建使用计划并将其绑定到服务的环境中使用。
func (c *Client) CreateUsagePlan(request *CreateUsagePlanRequest) (response *CreateUsagePlanResponse, err error) {
	if request == nil {
		request = NewCreateUsagePlanRequest()
	}
	response = NewCreateUsagePlanResponse()
	err = c.Send(request, response)
	return
}

func NewBuildAPIDocRequest() (request *BuildAPIDocRequest) {
	request = &BuildAPIDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "BuildAPIDoc")
	return
}

func NewBuildAPIDocResponse() (response *BuildAPIDocResponse) {
	response = &BuildAPIDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 构建 API 文档
func (c *Client) BuildAPIDoc(request *BuildAPIDocRequest) (response *BuildAPIDocResponse, err error) {
	if request == nil {
		request = NewBuildAPIDocRequest()
	}
	response = NewBuildAPIDocResponse()
	err = c.Send(request, response)
	return
}

func NewCheckActionToInstanceRequest() (request *CheckActionToInstanceRequest) {
	request = &CheckActionToInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CheckActionToInstance")
	return
}

func NewCheckActionToInstanceResponse() (response *CheckActionToInstanceResponse) {
	response = &CheckActionToInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CheckActionToInstance）用于检验实例操作是否允许。​
func (c *Client) CheckActionToInstance(request *CheckActionToInstanceRequest) (response *CheckActionToInstanceResponse, err error) {
	if request == nil {
		request = NewCheckActionToInstanceRequest()
	}
	response = NewCheckActionToInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCheckServiceConfigRequest() (request *CheckServiceConfigRequest) {
	request = &CheckServiceConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CheckServiceConfig")
	return
}

func NewCheckServiceConfigResponse() (response *CheckServiceConfigResponse) {
	response = &CheckServiceConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查服务配置合法性
func (c *Client) CheckServiceConfig(request *CheckServiceConfigRequest) (response *CheckServiceConfigResponse, err error) {
	if request == nil {
		request = NewCheckServiceConfigRequest()
	}
	response = NewCheckServiceConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUnBindEnvironmentRequest() (request *UnBindEnvironmentRequest) {
	request = &UnBindEnvironmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UnBindEnvironment")
	return
}

func NewUnBindEnvironmentResponse() (response *UnBindEnvironmentResponse) {
	response = &UnBindEnvironmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnBindEnvironment）用于将使用计划从特定环境解绑。
func (c *Client) UnBindEnvironment(request *UnBindEnvironmentRequest) (response *UnBindEnvironmentResponse, err error) {
	if request == nil {
		request = NewUnBindEnvironmentRequest()
	}
	response = NewUnBindEnvironmentResponse()
	err = c.Send(request, response)
	return
}

func NewCloneApisRequest() (request *CloneApisRequest) {
	request = &CloneApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CloneApis")
	return
}

func NewCloneApisResponse() (response *CloneApisResponse) {
	response = &CloneApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 克隆api
func (c *Client) CloneApis(request *CloneApisRequest) (response *CloneApisResponse, err error) {
	if request == nil {
		request = NewCloneApisRequest()
	}
	response = NewCloneApisResponse()
	err = c.Send(request, response)
	return
}

func NewRunApiForMarketRequest() (request *RunApiForMarketRequest) {
	request = &RunApiForMarketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "RunApiForMarket")
	return
}

func NewRunApiForMarketResponse() (response *RunApiForMarketResponse) {
	response = &RunApiForMarketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行API（云市场专用）
func (c *Client) RunApiForMarket(request *RunApiForMarketRequest) (response *RunApiForMarketResponse, err error) {
	if request == nil {
		request = NewRunApiForMarketRequest()
	}
	response = NewRunApiForMarketResponse()
	err = c.Send(request, response)
	return
}

func NewBindIPStrategyRequest() (request *BindIPStrategyRequest) {
	request = &BindIPStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "BindIPStrategy")
	return
}

func NewBindIPStrategyResponse() (response *BindIPStrategyResponse) {
	response = &BindIPStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BindIPStrategy）用于API绑定IP策略。
func (c *Client) BindIPStrategy(request *BindIPStrategyRequest) (response *BindIPStrategyResponse, err error) {
	if request == nil {
		request = NewBindIPStrategyRequest()
	}
	response = NewBindIPStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExclusiveSetRequest() (request *DescribeExclusiveSetRequest) {
	request = &DescribeExclusiveSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeExclusiveSet")
	return
}

func NewDescribeExclusiveSetResponse() (response *DescribeExclusiveSetResponse) {
	response = &DescribeExclusiveSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询独立集群列表
func (c *Client) DescribeExclusiveSet(request *DescribeExclusiveSetRequest) (response *DescribeExclusiveSetResponse, err error) {
	if request == nil {
		request = NewDescribeExclusiveSetRequest()
	}
	response = NewDescribeExclusiveSetResponse()
	err = c.Send(request, response)
	return
}

func NewUnBindSecretIdsRequest() (request *UnBindSecretIdsRequest) {
	request = &UnBindSecretIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSecretIds")
	return
}

func NewUnBindSecretIdsResponse() (response *UnBindSecretIdsResponse) {
	response = &UnBindSecretIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnBindSecretIds）用于为使用计划解绑密钥。
func (c *Client) UnBindSecretIds(request *UnBindSecretIdsRequest) (response *UnBindSecretIdsResponse, err error) {
	if request == nil {
		request = NewUnBindSecretIdsRequest()
	}
	response = NewUnBindSecretIdsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUpstreamRequest() (request *DeleteUpstreamRequest) {
	request = &DeleteUpstreamRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteUpstream")
	return
}

func NewDeleteUpstreamResponse() (response *DeleteUpstreamResponse) {
	response = &DeleteUpstreamResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除VPC通道，需要注意有api绑定时，不允许删除
func (c *Client) DeleteUpstream(request *DeleteUpstreamRequest) (response *DeleteUpstreamResponse, err error) {
	if request == nil {
		request = NewDeleteUpstreamRequest()
	}
	response = NewDeleteUpstreamResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePluginApisRequest() (request *DescribePluginApisRequest) {
	request = &DescribePluginApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribePluginApis")
	return
}

func NewDescribePluginApisResponse() (response *DescribePluginApisResponse) {
	response = &DescribePluginApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定插件下绑定的API信息
func (c *Client) DescribePluginApis(request *DescribePluginApisRequest) (response *DescribePluginApisResponse, err error) {
	if request == nil {
		request = NewDescribePluginApisRequest()
	}
	response = NewDescribePluginApisResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceReleaseVersionRequest() (request *DescribeServiceReleaseVersionRequest) {
	request = &DescribeServiceReleaseVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceReleaseVersion")
	return
}

func NewDescribeServiceReleaseVersionResponse() (response *DescribeServiceReleaseVersionResponse) {
	response = &DescribeServiceReleaseVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceReleaseVersion）查询一个服务下面所有已经发布的版本列表。
// 用户在发布服务时，常有多个版本发布，可使用本接口查询已发布的版本。
func (c *Client) DescribeServiceReleaseVersion(request *DescribeServiceReleaseVersionRequest) (response *DescribeServiceReleaseVersionResponse, err error) {
	if request == nil {
		request = NewDescribeServiceReleaseVersionRequest()
	}
	response = NewDescribeServiceReleaseVersionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApiKeyRequest() (request *CreateApiKeyRequest) {
	request = &CreateApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateApiKey")
	return
}

func NewCreateApiKeyResponse() (response *CreateApiKeyResponse) {
	response = &CreateApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateApiKey）用于创建一对新的 API 密钥。
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
	if request == nil {
		request = NewCreateApiKeyRequest()
	}
	response = NewCreateApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApiAppRequest() (request *DeleteApiAppRequest) {
	request = &DeleteApiAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApiApp")
	return
}

func NewDeleteApiAppResponse() (response *DeleteApiAppResponse) {
	response = &DeleteApiAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteApiApp）用于删除已经创建的应用。
func (c *Client) DeleteApiApp(request *DeleteApiAppRequest) (response *DeleteApiAppResponse, err error) {
	if request == nil {
		request = NewDeleteApiAppRequest()
	}
	response = NewDeleteApiAppResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogRulesRequest() (request *DescribeLogRulesRequest) {
	request = &DescribeLogRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeLogRules")
	return
}

func NewDescribeLogRulesResponse() (response *DescribeLogRulesResponse) {
	response = &DescribeLogRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志规则列表
func (c *Client) DescribeLogRules(request *DescribeLogRulesRequest) (response *DescribeLogRulesResponse, err error) {
	if request == nil {
		request = NewDescribeLogRulesRequest()
	}
	response = NewDescribeLogRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogRulesRequest() (request *DeleteLogRulesRequest) {
	request = &DeleteLogRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteLogRules")
	return
}

func NewDeleteLogRulesResponse() (response *DeleteLogRulesResponse) {
	response = &DeleteLogRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志规则。
func (c *Client) DeleteLogRules(request *DeleteLogRulesRequest) (response *DeleteLogRulesResponse, err error) {
	if request == nil {
		request = NewDeleteLogRulesRequest()
	}
	response = NewDeleteLogRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
	request = &DescribeServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeService")
	return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
	response = &DescribeServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeService）用于查询一个服务的详细信息、包括服务的描述、域名、协议、创建时间、发布情况等信息。
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
	if request == nil {
		request = NewDescribeServiceRequest()
	}
	response = NewDescribeServiceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServiceEnvironmentKeyMonitorUploadRequest() (request *ModifyServiceEnvironmentKeyMonitorUploadRequest) {
	request = &ModifyServiceEnvironmentKeyMonitorUploadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyServiceEnvironmentKeyMonitorUpload")
	return
}

func NewModifyServiceEnvironmentKeyMonitorUploadResponse() (response *ModifyServiceEnvironmentKeyMonitorUploadResponse) {
	response = &ModifyServiceEnvironmentKeyMonitorUploadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置服务的环境是否进行key维度的监控数据上报
func (c *Client) ModifyServiceEnvironmentKeyMonitorUpload(request *ModifyServiceEnvironmentKeyMonitorUploadRequest) (response *ModifyServiceEnvironmentKeyMonitorUploadResponse, err error) {
	if request == nil {
		request = NewModifyServiceEnvironmentKeyMonitorUploadRequest()
	}
	response = NewModifyServiceEnvironmentKeyMonitorUploadResponse()
	err = c.Send(request, response)
	return
}

func NewBindApiAppRequest() (request *BindApiAppRequest) {
	request = &BindApiAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "BindApiApp")
	return
}

func NewBindApiAppResponse() (response *BindApiAppResponse) {
	response = &BindApiAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BindApiApp）用于绑定应用到API。
func (c *Client) BindApiApp(request *BindApiAppRequest) (response *BindApiAppResponse, err error) {
	if request == nil {
		request = NewBindApiAppRequest()
	}
	response = NewBindApiAppResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiRequest() (request *ModifyApiRequest) {
	request = &ModifyApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApi")
	return
}

func NewModifyApiResponse() (response *ModifyApiResponse) {
	response = &ModifyApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyApi）用于修改 API 接口，可调用此接口对已经配置的 API 接口进行编辑修改。修改后的 API 需要重新发布 API 所在的服务到对应环境方能生效。
func (c *Client) ModifyApi(request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
	if request == nil {
		request = NewModifyApiRequest()
	}
	response = NewModifyApiResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTcbScfApiRequest() (request *CreateTcbScfApiRequest) {
	request = &CreateTcbScfApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateTcbScfApi")
	return
}

func NewCreateTcbScfApiResponse() (response *CreateTcbScfApiResponse) {
	response = &CreateTcbScfApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建TCB-SCF-HTTP触发器
func (c *Client) CreateTcbScfApi(request *CreateTcbScfApiRequest) (response *CreateTcbScfApiResponse, err error) {
	if request == nil {
		request = NewCreateTcbScfApiRequest()
	}
	response = NewCreateTcbScfApiResponse()
	err = c.Send(request, response)
	return
}

func NewBindEnvironmentRequest() (request *BindEnvironmentRequest) {
	request = &BindEnvironmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "BindEnvironment")
	return
}

func NewBindEnvironmentResponse() (response *BindEnvironmentResponse) {
	response = &BindEnvironmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BindEnvironment）用于绑定使用计划到服务或API。
// 用户在发布服务到某个环境中后，如果 API 需要鉴权，还需要绑定使用计划才能进行调用，此接口用户将使用计划绑定到特定环境。
// 目前支持绑定使用计划到API，但是同一个服务不能同时存在绑定到服务的使用计划和绑定到API的使用计划，所以对已经绑定过服务级别使用计划的环境，请先使用 服务级别使用计划降级 接口进行降级操作。
func (c *Client) BindEnvironment(request *BindEnvironmentRequest) (response *BindEnvironmentResponse, err error) {
	if request == nil {
		request = NewBindEnvironmentRequest()
	}
	response = NewBindEnvironmentResponse()
	err = c.Send(request, response)
	return
}

func NewEnableApiKeyRequest() (request *EnableApiKeyRequest) {
	request = &EnableApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "EnableApiKey")
	return
}

func NewEnableApiKeyResponse() (response *EnableApiKeyResponse) {
	response = &EnableApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（EnableApiKey）用于启动一对被禁用的 API 密钥。
func (c *Client) EnableApiKey(request *EnableApiKeyRequest) (response *EnableApiKeyResponse, err error) {
	if request == nil {
		request = NewEnableApiKeyRequest()
	}
	response = NewEnableApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSubDomainRequest() (request *UpdateSubDomainRequest) {
	request = &UpdateSubDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UpdateSubDomain")
	return
}

func NewUpdateSubDomainResponse() (response *UpdateSubDomainResponse) {
	response = &UpdateSubDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BindSubDomain）用于更新自定义域名状态。
// 定时检查自定义域名注册备案状态，下线未备案或已过期的域名，域名重新备案后在上线。
func (c *Client) UpdateSubDomain(request *UpdateSubDomainRequest) (response *UpdateSubDomainResponse, err error) {
	if request == nil {
		request = NewUpdateSubDomainRequest()
	}
	response = NewUpdateSubDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApiRequest() (request *DeleteApiRequest) {
	request = &DeleteApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApi")
	return
}

func NewDeleteApiResponse() (response *DeleteApiResponse) {
	response = &DeleteApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteApi）用于删除已经创建的API。
func (c *Client) DeleteApi(request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
	if request == nil {
		request = NewDeleteApiRequest()
	}
	response = NewDeleteApiResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServiceRequest() (request *ModifyServiceRequest) {
	request = &ModifyServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyService")
	return
}

func NewModifyServiceResponse() (response *ModifyServiceResponse) {
	response = &ModifyServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyService）用于修改服务的相关信息。当服务创建后，服务的名称、描述和服务类型均可被修改。
func (c *Client) ModifyService(request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
	if request == nil {
		request = NewModifyServiceRequest()
	}
	response = NewModifyServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiAppBindApisStatusRequest() (request *DescribeApiAppBindApisStatusRequest) {
	request = &DescribeApiAppBindApisStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiAppBindApisStatus")
	return
}

func NewDescribeApiAppBindApisStatusResponse() (response *DescribeApiAppBindApisStatusResponse) {
	response = &DescribeApiAppBindApisStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApiAppBindApisStatus）查询应用绑定的Api列表。
func (c *Client) DescribeApiAppBindApisStatus(request *DescribeApiAppBindApisStatusRequest) (response *DescribeApiAppBindApisStatusResponse, err error) {
	if request == nil {
		request = NewDescribeApiAppBindApisStatusRequest()
	}
	response = NewDescribeApiAppBindApisStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourcePackStatusRequest() (request *DescribeResourcePackStatusRequest) {
	request = &DescribeResourcePackStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeResourcePackStatus")
	return
}

func NewDescribeResourcePackStatusResponse() (response *DescribeResourcePackStatusResponse) {
	response = &DescribeResourcePackStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本DescribeResourcePackStatus用户展示资源包
func (c *Client) DescribeResourcePackStatus(request *DescribeResourcePackStatusRequest) (response *DescribeResourcePackStatusResponse, err error) {
	if request == nil {
		request = NewDescribeResourcePackStatusRequest()
	}
	response = NewDescribeResourcePackStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllPluginApisRequest() (request *DescribeAllPluginApisRequest) {
	request = &DescribeAllPluginApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAllPluginApis")
	return
}

func NewDescribeAllPluginApisResponse() (response *DescribeAllPluginApisResponse) {
	response = &DescribeAllPluginApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示插件相关的API列表，包括已绑定的和未绑定的API信息。
func (c *Client) DescribeAllPluginApis(request *DescribeAllPluginApisRequest) (response *DescribeAllPluginApisResponse, err error) {
	if request == nil {
		request = NewDescribeAllPluginApisRequest()
	}
	response = NewDescribeAllPluginApisResponse()
	err = c.Send(request, response)
	return
}

func NewMonitorTopsRequest() (request *MonitorTopsRequest) {
	request = &MonitorTopsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "MonitorTops")
	return
}

func NewMonitorTopsResponse() (response *MonitorTopsResponse) {
	response = &MonitorTopsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询监控top数据
func (c *Client) MonitorTops(request *MonitorTopsRequest) (response *MonitorTopsResponse, err error) {
	if request == nil {
		request = NewMonitorTopsRequest()
	}
	response = NewMonitorTopsResponse()
	err = c.Send(request, response)
	return
}

func NewResetAPIDocPasswordRequest() (request *ResetAPIDocPasswordRequest) {
	request = &ResetAPIDocPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ResetAPIDocPassword")
	return
}

func NewResetAPIDocPasswordResponse() (response *ResetAPIDocPasswordResponse) {
	response = &ResetAPIDocPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置API文档密码
func (c *Client) ResetAPIDocPassword(request *ResetAPIDocPasswordRequest) (response *ResetAPIDocPasswordResponse, err error) {
	if request == nil {
		request = NewResetAPIDocPasswordRequest()
	}
	response = NewResetAPIDocPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEIAMApisRequest() (request *DescribeEIAMApisRequest) {
	request = &DescribeEIAMApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeEIAMApis")
	return
}

func NewDescribeEIAMApisResponse() (response *DescribeEIAMApisResponse) {
	response = &DescribeEIAMApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询绑定EIAM应用的API列表
func (c *Client) DescribeEIAMApis(request *DescribeEIAMApisRequest) (response *DescribeEIAMApisResponse, err error) {
	if request == nil {
		request = NewDescribeEIAMApisRequest()
	}
	response = NewDescribeEIAMApisResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
	request = &CreateServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateService")
	return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
	response = &CreateServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateService）用于创建服务。
// API 网关使用的最大单元为服务，每个服务中可创建多个 API 接口。每个服务有一个默认域名供客户调用，用户也可绑定自定义域名到此服务中。
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
	if request == nil {
		request = NewCreateServiceRequest()
	}
	response = NewCreateServiceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLogRuleRequest() (request *CreateLogRuleRequest) {
	request = &CreateLogRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateLogRule")
	return
}

func NewCreateLogRuleResponse() (response *CreateLogRuleResponse) {
	response = &CreateLogRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用于日志上报的规则。
func (c *Client) CreateLogRule(request *CreateLogRuleRequest) (response *CreateLogRuleResponse, err error) {
	if request == nil {
		request = NewCreateLogRuleRequest()
	}
	response = NewCreateLogRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUnBindSubDomainRequest() (request *UnBindSubDomainRequest) {
	request = &UnBindSubDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSubDomain")
	return
}

func NewUnBindSubDomainResponse() (response *UnBindSubDomainResponse) {
	response = &UnBindSubDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnBindSubDomain）用于解绑自定义域名。
// 用户使用 API 网关绑定了自定义域名到服务中后，若想要解绑此自定义域名，可使用此接口。
func (c *Client) UnBindSubDomain(request *UnBindSubDomainRequest) (response *UnBindSubDomainResponse, err error) {
	if request == nil {
		request = NewUnBindSubDomainRequest()
	}
	response = NewUnBindSubDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePluginsRequest() (request *DeletePluginsRequest) {
	request = &DeletePluginsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeletePlugins")
	return
}

func NewDeletePluginsResponse() (response *DeletePluginsResponse) {
	response = &DeletePluginsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除API网关插件。
func (c *Client) DeletePlugins(request *DeletePluginsRequest) (response *DeletePluginsResponse, err error) {
	if request == nil {
		request = NewDeletePluginsRequest()
	}
	response = NewDeletePluginsResponse()
	err = c.Send(request, response)
	return
}

func NewUnBindIPStrategyRequest() (request *UnBindIPStrategyRequest) {
	request = &UnBindIPStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UnBindIPStrategy")
	return
}

func NewUnBindIPStrategyResponse() (response *UnBindIPStrategyResponse) {
	response = &UnBindIPStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnBindIPStrategy）用于服务解绑IP策略。
func (c *Client) UnBindIPStrategy(request *UnBindIPStrategyRequest) (response *UnBindIPStrategyResponse, err error) {
	if request == nil {
		request = NewUnBindIPStrategyRequest()
	}
	response = NewUnBindIPStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApiRequest() (request *CreateApiRequest) {
	request = &CreateApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateApi")
	return
}

func NewCreateApiResponse() (response *CreateApiResponse) {
	response = &CreateApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateApi）用于创建 API 接口，创建 API 前，用户需要先创建服务，每个 API 都有自己归属的服务。
func (c *Client) CreateApi(request *CreateApiRequest) (response *CreateApiResponse, err error) {
	if request == nil {
		request = NewCreateApiRequest()
	}
	response = NewCreateApiResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApisRequest() (request *DeleteApisRequest) {
	request = &DeleteApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApis")
	return
}

func NewDeleteApisResponse() (response *DeleteApisResponse) {
	response = &DeleteApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteApis）用于批量删除API。
func (c *Client) DeleteApis(request *DeleteApisRequest) (response *DeleteApisResponse, err error) {
	if request == nil {
		request = NewDeleteApisRequest()
	}
	response = NewDeleteApisResponse()
	err = c.Send(request, response)
	return
}

func NewMonitorTrendRequest() (request *MonitorTrendRequest) {
	request = &MonitorTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "MonitorTrend")
	return
}

func NewMonitorTrendResponse() (response *MonitorTrendResponse) {
	response = &MonitorTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取监控指标趋势，粒度为分钟
func (c *Client) MonitorTrend(request *MonitorTrendRequest) (response *MonitorTrendResponse, err error) {
	if request == nil {
		request = NewMonitorTrendRequest()
	}
	response = NewMonitorTrendResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApisAuthRelationApiRequest() (request *ModifyApisAuthRelationApiRequest) {
	request = &ModifyApisAuthRelationApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApisAuthRelationApi")
	return
}

func NewModifyApisAuthRelationApiResponse() (response *ModifyApisAuthRelationApiResponse) {
	response = &ModifyApisAuthRelationApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改关联授权API
func (c *Client) ModifyApisAuthRelationApi(request *ModifyApisAuthRelationApiRequest) (response *ModifyApisAuthRelationApiResponse, err error) {
	if request == nil {
		request = NewModifyApisAuthRelationApiRequest()
	}
	response = NewModifyApisAuthRelationApiResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindTriggerRequest() (request *UnbindTriggerRequest) {
	request = &UnbindTriggerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UnbindTrigger")
	return
}

func NewUnbindTriggerResponse() (response *UnbindTriggerResponse) {
	response = &UnbindTriggerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnbindTrigger）用于SCF侧来解绑网关触发器。
func (c *Client) UnbindTrigger(request *UnbindTriggerRequest) (response *UnbindTriggerResponse, err error) {
	if request == nil {
		request = NewUnbindTriggerRequest()
	}
	response = NewUnbindTriggerResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApiKeyRequest() (request *UpdateApiKeyRequest) {
	request = &UpdateApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UpdateApiKey")
	return
}

func NewUpdateApiKeyResponse() (response *UpdateApiKeyResponse) {
	response = &UpdateApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateApiKey）用于更换用户已创建的一对 API 密钥。
func (c *Client) UpdateApiKey(request *UpdateApiKeyRequest) (response *UpdateApiKeyResponse, err error) {
	if request == nil {
		request = NewUpdateApiKeyRequest()
	}
	response = NewUpdateApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPStrategysStatusRequest() (request *DescribeIPStrategysStatusRequest) {
	request = &DescribeIPStrategysStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategysStatus")
	return
}

func NewDescribeIPStrategysStatusResponse() (response *DescribeIPStrategysStatusResponse) {
	response = &DescribeIPStrategysStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeIPStrategysStatus）用于查询服务IP策略列表。
func (c *Client) DescribeIPStrategysStatus(request *DescribeIPStrategysStatusRequest) (response *DescribeIPStrategysStatusResponse, err error) {
	if request == nil {
		request = NewDescribeIPStrategysStatusRequest()
	}
	response = NewDescribeIPStrategysStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAPIDocRequest() (request *ModifyAPIDocRequest) {
	request = &ModifyAPIDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyAPIDoc")
	return
}

func NewModifyAPIDocResponse() (response *ModifyAPIDocResponse) {
	response = &ModifyAPIDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改 API 文档
func (c *Client) ModifyAPIDoc(request *ModifyAPIDocRequest) (response *ModifyAPIDocResponse, err error) {
	if request == nil {
		request = NewModifyAPIDocRequest()
	}
	response = NewModifyAPIDocResponse()
	err = c.Send(request, response)
	return
}

func NewModifyExclusiveInstanceRequest() (request *ModifyExclusiveInstanceRequest) {
	request = &ModifyExclusiveInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyExclusiveInstance")
	return
}

func NewModifyExclusiveInstanceResponse() (response *ModifyExclusiveInstanceResponse) {
	response = &ModifyExclusiveInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyExclusiveInstance）用于修改独享实例信息。​
func (c *Client) ModifyExclusiveInstance(request *ModifyExclusiveInstanceRequest) (response *ModifyExclusiveInstanceResponse, err error) {
	if request == nil {
		request = NewModifyExclusiveInstanceRequest()
	}
	response = NewModifyExclusiveInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDisableApiKeyRequest() (request *DisableApiKeyRequest) {
	request = &DisableApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DisableApiKey")
	return
}

func NewDisableApiKeyResponse() (response *DisableApiKeyResponse) {
	response = &DisableApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DisableApiKey）用于禁用一对 API 密钥。
func (c *Client) DisableApiKey(request *DisableApiKeyRequest) (response *DisableApiKeyResponse, err error) {
	if request == nil {
		request = NewDisableApiKeyRequest()
	}
	response = NewDisableApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiRequest() (request *DescribeApiRequest) {
	request = &DescribeApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApi")
	return
}

func NewDescribeApiResponse() (response *DescribeApiResponse) {
	response = &DescribeApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApi）用于查询用户 API 网关的 API 接口的详细信息。​
func (c *Client) DescribeApi(request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
	if request == nil {
		request = NewDescribeApiRequest()
	}
	response = NewDescribeApiResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateServiceRequest() (request *UpdateServiceRequest) {
	request = &UpdateServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UpdateService")
	return
}

func NewUpdateServiceResponse() (response *UpdateServiceResponse) {
	response = &UpdateServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateService）用于从服务发布的环境中运行版本切换到特定版本。用户在使用 API 网关创建服务并发布服务到某个环境后，多因为开发过程会产生多个版本，此时可调用本接口。
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
	if request == nil {
		request = NewUpdateServiceRequest()
	}
	response = NewUpdateServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPStrategyApisStatusRequest() (request *DescribeIPStrategyApisStatusRequest) {
	request = &DescribeIPStrategyApisStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategyApisStatus")
	return
}

func NewDescribeIPStrategyApisStatusResponse() (response *DescribeIPStrategyApisStatusResponse) {
	response = &DescribeIPStrategyApisStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeIPStrategyApisStatus）用于查询IP策略可以绑定的API列表。即服务下所有API和该策略已绑定API的差集。
func (c *Client) DescribeIPStrategyApisStatus(request *DescribeIPStrategyApisStatusRequest) (response *DescribeIPStrategyApisStatusResponse, err error) {
	if request == nil {
		request = NewDescribeIPStrategyApisStatusRequest()
	}
	response = NewDescribeIPStrategyApisStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTcbScfApisRequest() (request *DescribeTcbScfApisRequest) {
	request = &DescribeTcbScfApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeTcbScfApis")
	return
}

func NewDescribeTcbScfApisResponse() (response *DescribeTcbScfApisResponse) {
	response = &DescribeTcbScfApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询TCB-SCF-HTTP触发器列表
func (c *Client) DescribeTcbScfApis(request *DescribeTcbScfApisRequest) (response *DescribeTcbScfApisResponse, err error) {
	if request == nil {
		request = NewDescribeTcbScfApisRequest()
	}
	response = NewDescribeTcbScfApisResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPStrategyRequest() (request *DescribeIPStrategyRequest) {
	request = &DescribeIPStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategy")
	return
}

func NewDescribeIPStrategyResponse() (response *DescribeIPStrategyResponse) {
	response = &DescribeIPStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeIPStrategy）用于查询IP策略详情。
func (c *Client) DescribeIPStrategy(request *DescribeIPStrategyRequest) (response *DescribeIPStrategyResponse, err error) {
	if request == nil {
		request = NewDescribeIPStrategyRequest()
	}
	response = NewDescribeIPStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiAppsStatusRequest() (request *DescribeApiAppsStatusRequest) {
	request = &DescribeApiAppsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiAppsStatus")
	return
}

func NewDescribeApiAppsStatusResponse() (response *DescribeApiAppsStatusResponse) {
	response = &DescribeApiAppsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApiAppsStatus）查询应用列表。
func (c *Client) DescribeApiAppsStatus(request *DescribeApiAppsStatusRequest) (response *DescribeApiAppsStatusResponse, err error) {
	if request == nil {
		request = NewDescribeApiAppsStatusRequest()
	}
	response = NewDescribeApiAppsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiEnvironmentStrategyRequest() (request *DescribeApiEnvironmentStrategyRequest) {
	request = &DescribeApiEnvironmentStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiEnvironmentStrategy")
	return
}

func NewDescribeApiEnvironmentStrategyResponse() (response *DescribeApiEnvironmentStrategyResponse) {
	response = &DescribeApiEnvironmentStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApiEnvironmentStrategy）用于展示API绑定的限流策略。
func (c *Client) DescribeApiEnvironmentStrategy(request *DescribeApiEnvironmentStrategyRequest) (response *DescribeApiEnvironmentStrategyResponse, err error) {
	if request == nil {
		request = NewDescribeApiEnvironmentStrategyRequest()
	}
	response = NewDescribeApiEnvironmentStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePluginsRequest() (request *DescribePluginsRequest) {
	request = &DescribePluginsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribePlugins")
	return
}

func NewDescribePluginsResponse() (response *DescribePluginsResponse) {
	response = &DescribePluginsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示插件列表和详情，支持分页，支持按照插件类型查询，支持按照插件ID批量查询，支持按照插件名称查询。
func (c *Client) DescribePlugins(request *DescribePluginsRequest) (response *DescribePluginsResponse, err error) {
	if request == nil {
		request = NewDescribePluginsRequest()
	}
	response = NewDescribePluginsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiBindApiAppsStatusRequest() (request *DescribeApiBindApiAppsStatusRequest) {
	request = &DescribeApiBindApiAppsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiBindApiAppsStatus")
	return
}

func NewDescribeApiBindApiAppsStatusResponse() (response *DescribeApiBindApiAppsStatusResponse) {
	response = &DescribeApiBindApiAppsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApiBindApiAppsStatus）查询Api绑定的应用列表。
func (c *Client) DescribeApiBindApiAppsStatus(request *DescribeApiBindApiAppsStatusRequest) (response *DescribeApiBindApiAppsStatusResponse, err error) {
	if request == nil {
		request = NewDescribeApiBindApiAppsStatusRequest()
	}
	response = NewDescribeApiBindApiAppsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsagePlanSecretIdsRequest() (request *DescribeUsagePlanSecretIdsRequest) {
	request = &DescribeUsagePlanSecretIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanSecretIds")
	return
}

func NewDescribeUsagePlanSecretIdsResponse() (response *DescribeUsagePlanSecretIdsResponse) {
	response = &DescribeUsagePlanSecretIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeUsagePlanSecretIds）用于查询使用计划绑定的密钥列表。
// 在 API 网关中，一个使用计划可绑定多个密钥对，可使用本接口查询使用计划绑定的密钥列表。
func (c *Client) DescribeUsagePlanSecretIds(request *DescribeUsagePlanSecretIdsRequest) (response *DescribeUsagePlanSecretIdsResponse, err error) {
	if request == nil {
		request = NewDescribeUsagePlanSecretIdsRequest()
	}
	response = NewDescribeUsagePlanSecretIdsResponse()
	err = c.Send(request, response)
	return
}

func NewUnReleaseServiceRequest() (request *UnReleaseServiceRequest) {
	request = &UnReleaseServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UnReleaseService")
	return
}

func NewUnReleaseServiceResponse() (response *UnReleaseServiceResponse) {
	response = &UnReleaseServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnReleaseService）用于下线服务。
// 用户发布服务到某个环境后，此服务中的 API 方可被调用者进行调用，当用户需要将此服务从发布环境中下线时，可调用此 API。下线后的服务不可被调用。
func (c *Client) UnReleaseService(request *UnReleaseServiceRequest) (response *UnReleaseServiceResponse, err error) {
	if request == nil {
		request = NewUnReleaseServiceRequest()
	}
	response = NewUnReleaseServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiForApiAppRequest() (request *DescribeApiForApiAppRequest) {
	request = &DescribeApiForApiAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiForApiApp")
	return
}

func NewDescribeApiForApiAppResponse() (response *DescribeApiForApiAppResponse) {
	response = &DescribeApiForApiAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApiForApiApp）用于应用使用者查询部署于 API 网关的 API 接口的详细信息。​
func (c *Client) DescribeApiForApiApp(request *DescribeApiForApiAppRequest) (response *DescribeApiForApiAppResponse, err error) {
	if request == nil {
		request = NewDescribeApiForApiAppRequest()
	}
	response = NewDescribeApiForApiAppResponse()
	err = c.Send(request, response)
	return
}

func NewBindSecretIdsRequest() (request *BindSecretIdsRequest) {
	request = &BindSecretIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "BindSecretIds")
	return
}

func NewBindSecretIdsResponse() (response *BindSecretIdsResponse) {
	response = &BindSecretIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BindSecretIds）用于为使用计划绑定密钥。
// 将密钥绑定到某个使用计划，并将此使用计划绑定到某个服务发布的环境上，调用者方可使用此密钥调用这个服务中的 API，可使用本接口为使用计划绑定密钥。
func (c *Client) BindSecretIds(request *BindSecretIdsRequest) (response *BindSecretIdsResponse, err error) {
	if request == nil {
		request = NewBindSecretIdsRequest()
	}
	response = NewBindSecretIdsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiEnvironmentApiKeysRequest() (request *DescribeApiEnvironmentApiKeysRequest) {
	request = &DescribeApiEnvironmentApiKeysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiEnvironmentApiKeys")
	return
}

func NewDescribeApiEnvironmentApiKeysResponse() (response *DescribeApiEnvironmentApiKeysResponse) {
	response = &DescribeApiEnvironmentApiKeysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询已发布API密钥
func (c *Client) DescribeApiEnvironmentApiKeys(request *DescribeApiEnvironmentApiKeysRequest) (response *DescribeApiEnvironmentApiKeysResponse, err error) {
	if request == nil {
		request = NewDescribeApiEnvironmentApiKeysRequest()
	}
	response = NewDescribeApiEnvironmentApiKeysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceSubDomainMappingsRequest() (request *DescribeServiceSubDomainMappingsRequest) {
	request = &DescribeServiceSubDomainMappingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomainMappings")
	return
}

func NewDescribeServiceSubDomainMappingsResponse() (response *DescribeServiceSubDomainMappingsResponse) {
	response = &DescribeServiceSubDomainMappingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceSubDomainMappings）用于查询自定义域名的路径映射。
// API 网关可绑定自定义域名到服务，并且可以对自定义域名的路径进行映射，可自定义不同的路径映射到服务中的三个环境，本接口用于查询绑定服务的自定义域名的路径映射列表。
func (c *Client) DescribeServiceSubDomainMappings(request *DescribeServiceSubDomainMappingsRequest) (response *DescribeServiceSubDomainMappingsResponse, err error) {
	if request == nil {
		request = NewDescribeServiceSubDomainMappingsRequest()
	}
	response = NewDescribeServiceSubDomainMappingsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogRuleRequest() (request *ModifyLogRuleRequest) {
	request = &ModifyLogRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyLogRule")
	return
}

func NewModifyLogRuleResponse() (response *ModifyLogRuleResponse) {
	response = &ModifyLogRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志规则
func (c *Client) ModifyLogRule(request *ModifyLogRuleRequest) (response *ModifyLogRuleResponse, err error) {
	if request == nil {
		request = NewModifyLogRuleRequest()
	}
	response = NewModifyLogRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePluginRequest() (request *DeletePluginRequest) {
	request = &DeletePluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeletePlugin")
	return
}

func NewDeletePluginResponse() (response *DeletePluginResponse) {
	response = &DeletePluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除API网关插件
func (c *Client) DeletePlugin(request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
	if request == nil {
		request = NewDeletePluginRequest()
	}
	response = NewDeletePluginResponse()
	err = c.Send(request, response)
	return
}

func NewGenerateApiDocumentRequest() (request *GenerateApiDocumentRequest) {
	request = &GenerateApiDocumentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "GenerateApiDocument")
	return
}

func NewGenerateApiDocumentResponse() (response *GenerateApiDocumentResponse) {
	response = &GenerateApiDocumentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GenerateApiDocument）用于自动生成 API 文档和 SDK，一个服务的一个环境生成一份文档和 SDK。
func (c *Client) GenerateApiDocument(request *GenerateApiDocumentRequest) (response *GenerateApiDocumentResponse, err error) {
	if request == nil {
		request = NewGenerateApiDocumentRequest()
	}
	response = NewGenerateApiDocumentResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseServiceRequest() (request *ReleaseServiceRequest) {
	request = &ReleaseServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ReleaseService")
	return
}

func NewReleaseServiceResponse() (response *ReleaseServiceResponse) {
	response = &ReleaseServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ReleaseService）用于发布服务。
// API 网关的服务创建后，需要发布到某个环境方生效后，使用者才能进行调用，此接口用于发布服务到环境，如 release 环境。
func (c *Client) ReleaseService(request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
	if request == nil {
		request = NewReleaseServiceRequest()
	}
	response = NewReleaseServiceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAPIDocRequest() (request *CreateAPIDocRequest) {
	request = &CreateAPIDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateAPIDoc")
	return
}

func NewCreateAPIDocResponse() (response *CreateAPIDocResponse) {
	response = &CreateAPIDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建 API 文档
func (c *Client) CreateAPIDoc(request *CreateAPIDocRequest) (response *CreateAPIDocResponse, err error) {
	if request == nil {
		request = NewCreateAPIDocRequest()
	}
	response = NewCreateAPIDocResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIPStrategyRequest() (request *DeleteIPStrategyRequest) {
	request = &DeleteIPStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteIPStrategy")
	return
}

func NewDeleteIPStrategyResponse() (response *DeleteIPStrategyResponse) {
	response = &DeleteIPStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteIPStrategy）用于删除服务IP策略。
func (c *Client) DeleteIPStrategy(request *DeleteIPStrategyRequest) (response *DeleteIPStrategyResponse, err error) {
	if request == nil {
		request = NewDeleteIPStrategyRequest()
	}
	response = NewDeleteIPStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAPIDocDetailRequest() (request *DescribeAPIDocDetailRequest) {
	request = &DescribeAPIDocDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAPIDocDetail")
	return
}

func NewDescribeAPIDocDetailResponse() (response *DescribeAPIDocDetailResponse) {
	response = &DescribeAPIDocDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 API 文档详情
func (c *Client) DescribeAPIDocDetail(request *DescribeAPIDocDetailRequest) (response *DescribeAPIDocDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAPIDocDetailRequest()
	}
	response = NewDescribeAPIDocDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceEnvironmentKeyMonitorUploadRequest() (request *DescribeServiceEnvironmentKeyMonitorUploadRequest) {
	request = &DescribeServiceEnvironmentKeyMonitorUploadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentKeyMonitorUpload")
	return
}

func NewDescribeServiceEnvironmentKeyMonitorUploadResponse() (response *DescribeServiceEnvironmentKeyMonitorUploadResponse) {
	response = &DescribeServiceEnvironmentKeyMonitorUploadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务的环境是否进行key维度的监控数据上报
func (c *Client) DescribeServiceEnvironmentKeyMonitorUpload(request *DescribeServiceEnvironmentKeyMonitorUploadRequest) (response *DescribeServiceEnvironmentKeyMonitorUploadResponse, err error) {
	if request == nil {
		request = NewDescribeServiceEnvironmentKeyMonitorUploadRequest()
	}
	response = NewDescribeServiceEnvironmentKeyMonitorUploadResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpstreamBindApisRequest() (request *DescribeUpstreamBindApisRequest) {
	request = &DescribeUpstreamBindApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUpstreamBindApis")
	return
}

func NewDescribeUpstreamBindApisResponse() (response *DescribeUpstreamBindApisResponse) {
	response = &DescribeUpstreamBindApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VPC通道绑定的api列表
func (c *Client) DescribeUpstreamBindApis(request *DescribeUpstreamBindApisRequest) (response *DescribeUpstreamBindApisResponse, err error) {
	if request == nil {
		request = NewDescribeUpstreamBindApisRequest()
	}
	response = NewDescribeUpstreamBindApisResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainRequest() (request *DescribeDomainRequest) {
	request = &DescribeDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeDomain")
	return
}

func NewDescribeDomainResponse() (response *DescribeDomainResponse) {
	response = &DescribeDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询域名状态
func (c *Client) DescribeDomain(request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
	if request == nil {
		request = NewDescribeDomainRequest()
	}
	response = NewDescribeDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsagePlanEnvironmentsRequest() (request *DescribeUsagePlanEnvironmentsRequest) {
	request = &DescribeUsagePlanEnvironmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanEnvironments")
	return
}

func NewDescribeUsagePlanEnvironmentsResponse() (response *DescribeUsagePlanEnvironmentsResponse) {
	response = &DescribeUsagePlanEnvironmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeUsagePlanEnvironments）用于查询使用计划绑定的环境列表。
// 用户在绑定了某个使用计划到环境后，可使用本接口查询这个使用计划绑定的所有服务的环境。
func (c *Client) DescribeUsagePlanEnvironments(request *DescribeUsagePlanEnvironmentsRequest) (response *DescribeUsagePlanEnvironmentsResponse, err error) {
	if request == nil {
		request = NewDescribeUsagePlanEnvironmentsRequest()
	}
	response = NewDescribeUsagePlanEnvironmentsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiIncrementRequest() (request *ModifyApiIncrementRequest) {
	request = &ModifyApiIncrementRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiIncrement")
	return
}

func NewModifyApiIncrementResponse() (response *ModifyApiIncrementResponse) {
	response = &ModifyApiIncrementResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供增量更新API能力，主要是给程序调用（区别于ModifyApi，该接口是需要传入API的全量参数，对console使用较友好）
func (c *Client) ModifyApiIncrement(request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
	if request == nil {
		request = NewModifyApiIncrementRequest()
	}
	response = NewModifyApiIncrementResponse()
	err = c.Send(request, response)
	return
}

func NewCheckCloneApisRequest() (request *CheckCloneApisRequest) {
	request = &CheckCloneApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CheckCloneApis")
	return
}

func NewCheckCloneApisResponse() (response *CheckCloneApisResponse) {
	response = &CheckCloneApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验克隆api 是否满足要求
func (c *Client) CheckCloneApis(request *CheckCloneApisRequest) (response *CheckCloneApisResponse, err error) {
	if request == nil {
		request = NewCheckCloneApisRequest()
	}
	response = NewCheckCloneApisResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceSubDomainMappingRequest() (request *DeleteServiceSubDomainMappingRequest) {
	request = &DeleteServiceSubDomainMappingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteServiceSubDomainMapping")
	return
}

func NewDeleteServiceSubDomainMappingResponse() (response *DeleteServiceSubDomainMappingResponse) {
	response = &DeleteServiceSubDomainMappingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteServiceSubDomainMapping）用于删除服务中某个环境的自定义域名映射。
// 当用户使用自定义域名，并使用了自定义映射时，可使用此接口。但需注意，若删除了所有环境的映射时，调用此 API 均会返回失败。
func (c *Client) DeleteServiceSubDomainMapping(request *DeleteServiceSubDomainMappingRequest) (response *DeleteServiceSubDomainMappingResponse, err error) {
	if request == nil {
		request = NewDeleteServiceSubDomainMappingRequest()
	}
	response = NewDeleteServiceSubDomainMappingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApisStatusRequest() (request *DescribeApisStatusRequest) {
	request = &DescribeApisStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApisStatus")
	return
}

func NewDescribeApisStatusResponse() (response *DescribeApisStatusResponse) {
	response = &DescribeApisStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApisStatus）用于查看一个服务下的某个 API 或所有 API 列表及其相关信息。
func (c *Client) DescribeApisStatus(request *DescribeApisStatusRequest) (response *DescribeApisStatusResponse, err error) {
	if request == nil {
		request = NewDescribeApisStatusRequest()
	}
	response = NewDescribeApisStatusResponse()
	err = c.Send(request, response)
	return
}

func NewExportOpenApiRequest() (request *ExportOpenApiRequest) {
	request = &ExportOpenApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ExportOpenApi")
	return
}

func NewExportOpenApiResponse() (response *ExportOpenApiResponse) {
	response = &ExportOpenApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ExportOpenApi）用于将API网关创建的API按OpenAPI规范导出。
func (c *Client) ExportOpenApi(request *ExportOpenApiRequest) (response *ExportOpenApiResponse, err error) {
	if request == nil {
		request = NewExportOpenApiRequest()
	}
	response = NewExportOpenApiResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindTriggersRequest() (request *UnbindTriggersRequest) {
	request = &UnbindTriggersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UnbindTriggers")
	return
}

func NewUnbindTriggersResponse() (response *UnbindTriggersResponse) {
	response = &UnbindTriggersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnbindTrigger）用于SCF侧批量解绑网关触发器。
func (c *Client) UnbindTriggers(request *UnbindTriggersRequest) (response *UnbindTriggersResponse, err error) {
	if request == nil {
		request = NewUnbindTriggersRequest()
	}
	response = NewUnbindTriggersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServicesStatusRequest() (request *DescribeServicesStatusRequest) {
	request = &DescribeServicesStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServicesStatus")
	return
}

func NewDescribeServicesStatusResponse() (response *DescribeServicesStatusResponse) {
	response = &DescribeServicesStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServicesStatus）用于搜索查询某一个服务或多个服务的列表，并返回服务相关的域名、时间等信息。
func (c *Client) DescribeServicesStatus(request *DescribeServicesStatusRequest) (response *DescribeServicesStatusResponse, err error) {
	if request == nil {
		request = NewDescribeServicesStatusRequest()
	}
	response = NewDescribeServicesStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
	request = &DeleteServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteService")
	return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
	response = &DeleteServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteService）用于删除 API 网关中某个服务。
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
	if request == nil {
		request = NewDeleteServiceRequest()
	}
	response = NewDeleteServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceEnvironmentStrategyRequest() (request *DescribeServiceEnvironmentStrategyRequest) {
	request = &DescribeServiceEnvironmentStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentStrategy")
	return
}

func NewDescribeServiceEnvironmentStrategyResponse() (response *DescribeServiceEnvironmentStrategyResponse) {
	response = &DescribeServiceEnvironmentStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceEnvironmentStrategy）用于展示服务限流策略。
func (c *Client) DescribeServiceEnvironmentStrategy(request *DescribeServiceEnvironmentStrategyRequest) (response *DescribeServiceEnvironmentStrategyResponse, err error) {
	if request == nil {
		request = NewDescribeServiceEnvironmentStrategyRequest()
	}
	response = NewDescribeServiceEnvironmentStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDetachPluginRequest() (request *DetachPluginRequest) {
	request = &DetachPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DetachPlugin")
	return
}

func NewDetachPluginResponse() (response *DetachPluginResponse) {
	response = &DetachPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除插件与API绑定
func (c *Client) DetachPlugin(request *DetachPluginRequest) (response *DetachPluginResponse, err error) {
	if request == nil {
		request = NewDetachPluginRequest()
	}
	response = NewDetachPluginResponse()
	err = c.Send(request, response)
	return
}

func NewAttachPluginRequest() (request *AttachPluginRequest) {
	request = &AttachPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "AttachPlugin")
	return
}

func NewAttachPluginResponse() (response *AttachPluginResponse) {
	response = &AttachPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定插件到API上。
func (c *Client) AttachPlugin(request *AttachPluginRequest) (response *AttachPluginResponse, err error) {
	if request == nil {
		request = NewAttachPluginRequest()
	}
	response = NewAttachPluginResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUsagePlanRequest() (request *ModifyUsagePlanRequest) {
	request = &ModifyUsagePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyUsagePlan")
	return
}

func NewModifyUsagePlanResponse() (response *ModifyUsagePlanResponse) {
	response = &ModifyUsagePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyUsagePlan）用于修改使用计划的名称，描述及 QPS。
func (c *Client) ModifyUsagePlan(request *ModifyUsagePlanRequest) (response *ModifyUsagePlanResponse, err error) {
	if request == nil {
		request = NewModifyUsagePlanRequest()
	}
	response = NewModifyUsagePlanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAPIDocsRequest() (request *DescribeAPIDocsRequest) {
	request = &DescribeAPIDocsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAPIDocs")
	return
}

func NewDescribeAPIDocsResponse() (response *DescribeAPIDocsResponse) {
	response = &DescribeAPIDocsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 API 文档列表
func (c *Client) DescribeAPIDocs(request *DescribeAPIDocsRequest) (response *DescribeAPIDocsResponse, err error) {
	if request == nil {
		request = NewDescribeAPIDocsRequest()
	}
	response = NewDescribeAPIDocsResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateExclusiveInstanceRequest() (request *TerminateExclusiveInstanceRequest) {
	request = &TerminateExclusiveInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "TerminateExclusiveInstance")
	return
}

func NewTerminateExclusiveInstanceResponse() (response *TerminateExclusiveInstanceResponse) {
	response = &TerminateExclusiveInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例销毁接口，只有按量计费的实例才允许销毁，且实例下不存在服务
func (c *Client) TerminateExclusiveInstance(request *TerminateExclusiveInstanceRequest) (response *TerminateExclusiveInstanceResponse, err error) {
	if request == nil {
		request = NewTerminateExclusiveInstanceRequest()
	}
	response = NewTerminateExclusiveInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceUsagePlanRequest() (request *DescribeServiceUsagePlanRequest) {
	request = &DescribeServiceUsagePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceUsagePlan")
	return
}

func NewDescribeServiceUsagePlanResponse() (response *DescribeServiceUsagePlanResponse) {
	response = &DescribeServiceUsagePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceUsagePlan）用于查询服务使用计划详情。
// 服务若需要鉴权限流生效，则需要绑定使用计划到此服务中，本接口用于查询绑定到一个服务的所有使用计划。
func (c *Client) DescribeServiceUsagePlan(request *DescribeServiceUsagePlanRequest) (response *DescribeServiceUsagePlanResponse, err error) {
	if request == nil {
		request = NewDescribeServiceUsagePlanRequest()
	}
	response = NewDescribeServiceUsagePlanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiUsagePlanRequest() (request *DescribeApiUsagePlanRequest) {
	request = &DescribeApiUsagePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiUsagePlan")
	return
}

func NewDescribeApiUsagePlanResponse() (response *DescribeApiUsagePlanResponse) {
	response = &DescribeApiUsagePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApiUsagePlan）用于查询服务中 API 使用计划详情。
// 服务若需要鉴权限流生效，则需要绑定使用计划到此服务中，本接口用于查询绑定到一个服务及其中 API 的所有使用计划。
func (c *Client) DescribeApiUsagePlan(request *DescribeApiUsagePlanRequest) (response *DescribeApiUsagePlanResponse, err error) {
	if request == nil {
		request = NewDescribeApiUsagePlanRequest()
	}
	response = NewDescribeApiUsagePlanResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUsagePlanRequest() (request *DeleteUsagePlanRequest) {
	request = &DeleteUsagePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteUsagePlan")
	return
}

func NewDeleteUsagePlanResponse() (response *DeleteUsagePlanResponse) {
	response = &DeleteUsagePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteUsagePlan）用于删除使用计划。
func (c *Client) DeleteUsagePlan(request *DeleteUsagePlanRequest) (response *DeleteUsagePlanResponse, err error) {
	if request == nil {
		request = NewDeleteUsagePlanRequest()
	}
	response = NewDeleteUsagePlanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiKeysStatusRequest() (request *DescribeApiKeysStatusRequest) {
	request = &DescribeApiKeysStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiKeysStatus")
	return
}

func NewDescribeApiKeysStatusResponse() (response *DescribeApiKeysStatusResponse) {
	response = &DescribeApiKeysStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeApiKeysStatus）用于查询密钥列表。
// 当用户创建了多个密钥对时，可使用本接口查询一个或多个 API 密钥信息，本接口不会显示密钥 Key。
func (c *Client) DescribeApiKeysStatus(request *DescribeApiKeysStatusRequest) (response *DescribeApiKeysStatusResponse, err error) {
	if request == nil {
		request = NewDescribeApiKeysStatusRequest()
	}
	response = NewDescribeApiKeysStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpstreamsRequest() (request *DescribeUpstreamsRequest) {
	request = &DescribeUpstreamsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUpstreams")
	return
}

func NewDescribeUpstreamsResponse() (response *DescribeUpstreamsResponse) {
	response = &DescribeUpstreamsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VPC通道列表详情
func (c *Client) DescribeUpstreams(request *DescribeUpstreamsRequest) (response *DescribeUpstreamsResponse, err error) {
	if request == nil {
		request = NewDescribeUpstreamsRequest()
	}
	response = NewDescribeUpstreamsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIPStrategyRequest() (request *CreateIPStrategyRequest) {
	request = &CreateIPStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateIPStrategy")
	return
}

func NewCreateIPStrategyResponse() (response *CreateIPStrategyResponse) {
	response = &CreateIPStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateIPStrategy）用于创建服务IP策略。
func (c *Client) CreateIPStrategy(request *CreateIPStrategyRequest) (response *CreateIPStrategyResponse, err error) {
	if request == nil {
		request = NewCreateIPStrategyRequest()
	}
	response = NewCreateIPStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiAppRequest() (request *ModifyApiAppRequest) {
	request = &ModifyApiAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiApp")
	return
}

func NewModifyApiAppResponse() (response *ModifyApiAppResponse) {
	response = &ModifyApiAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyApiApp）用于修改已经创建的应用。
func (c *Client) ModifyApiApp(request *ModifyApiAppRequest) (response *ModifyApiAppResponse, err error) {
	if request == nil {
		request = NewModifyApiAppRequest()
	}
	response = NewModifyApiAppResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePluginsByApiRequest() (request *DescribePluginsByApiRequest) {
	request = &DescribePluginsByApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribePluginsByApi")
	return
}

func NewDescribePluginsByApiResponse() (response *DescribePluginsByApiResponse) {
	response = &DescribePluginsByApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示API上已绑定的插件列表。
func (c *Client) DescribePluginsByApi(request *DescribePluginsByApiRequest) (response *DescribePluginsByApiResponse, err error) {
	if request == nil {
		request = NewDescribePluginsByApiRequest()
	}
	response = NewDescribePluginsByApiResponse()
	err = c.Send(request, response)
	return
}

func NewDemoteServiceUsagePlanRequest() (request *DemoteServiceUsagePlanRequest) {
	request = &DemoteServiceUsagePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DemoteServiceUsagePlan")
	return
}

func NewDemoteServiceUsagePlanResponse() (response *DemoteServiceUsagePlanResponse) {
	response = &DemoteServiceUsagePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DemoteServiceUsagePlan）用于将某个服务在某个环境的使用计划，降级到API上。
// 如果服务内没有API不允许进行此操作。
// 如果当前环境没有发布，不允许进行此操作。
func (c *Client) DemoteServiceUsagePlan(request *DemoteServiceUsagePlanRequest) (response *DemoteServiceUsagePlanResponse, err error) {
	if request == nil {
		request = NewDemoteServiceUsagePlanRequest()
	}
	response = NewDemoteServiceUsagePlanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceEnvironmentReleaseHistoryRequest() (request *DescribeServiceEnvironmentReleaseHistoryRequest) {
	request = &DescribeServiceEnvironmentReleaseHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentReleaseHistory")
	return
}

func NewDescribeServiceEnvironmentReleaseHistoryResponse() (response *DescribeServiceEnvironmentReleaseHistoryResponse) {
	response = &DescribeServiceEnvironmentReleaseHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceEnvironmentReleaseHistory）用于查询服务环境的发布历史。
// 用户在创建好服务后需要发布到某个环境中才能进行使用，本接口用于查询一个服务某个环境的发布记录。
func (c *Client) DescribeServiceEnvironmentReleaseHistory(request *DescribeServiceEnvironmentReleaseHistoryRequest) (response *DescribeServiceEnvironmentReleaseHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeServiceEnvironmentReleaseHistoryRequest()
	}
	response = NewDescribeServiceEnvironmentReleaseHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExclusiveInstancesRequest() (request *DescribeExclusiveInstancesRequest) {
	request = &DescribeExclusiveInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeExclusiveInstances")
	return
}

func NewDescribeExclusiveInstancesResponse() (response *DescribeExclusiveInstancesResponse) {
	response = &DescribeExclusiveInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeExclusiveInstances）用于查询独享实例列表信息。​
func (c *Client) DescribeExclusiveInstances(request *DescribeExclusiveInstancesRequest) (response *DescribeExclusiveInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeExclusiveInstancesRequest()
	}
	response = NewDescribeExclusiveInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApiKeyRequest() (request *DeleteApiKeyRequest) {
	request = &DeleteApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApiKey")
	return
}

func NewDeleteApiKeyResponse() (response *DeleteApiKeyResponse) {
	response = &DeleteApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteApiKey）用于删除一对 API 密钥。
func (c *Client) DeleteApiKey(request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
	if request == nil {
		request = NewDeleteApiKeyRequest()
	}
	response = NewDeleteApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExclusiveInstanceDetailRequest() (request *DescribeExclusiveInstanceDetailRequest) {
	request = &DescribeExclusiveInstanceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeExclusiveInstanceDetail")
	return
}

func NewDescribeExclusiveInstanceDetailResponse() (response *DescribeExclusiveInstanceDetailResponse) {
	response = &DescribeExclusiveInstanceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeExclusiveInstanceDetail）用于查询独享实例详情信息。​
func (c *Client) DescribeExclusiveInstanceDetail(request *DescribeExclusiveInstanceDetailRequest) (response *DescribeExclusiveInstanceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeExclusiveInstanceDetailRequest()
	}
	response = NewDescribeExclusiveInstanceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUpstreamRequest() (request *ModifyUpstreamRequest) {
	request = &ModifyUpstreamRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyUpstream")
	return
}

func NewModifyUpstreamResponse() (response *ModifyUpstreamResponse) {
	response = &ModifyUpstreamResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改VPC通道
func (c *Client) ModifyUpstream(request *ModifyUpstreamRequest) (response *ModifyUpstreamResponse, err error) {
	if request == nil {
		request = NewModifyUpstreamRequest()
	}
	response = NewModifyUpstreamResponse()
	err = c.Send(request, response)
	return
}

func NewBindSubDomainRequest() (request *BindSubDomainRequest) {
	request = &BindSubDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "BindSubDomain")
	return
}

func NewBindSubDomainResponse() (response *BindSubDomainResponse) {
	response = &BindSubDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BindSubDomain）用于绑定自定义域名到服务。
// API 网关中每个服务都会提供一个默认的域名供用户调用，但当用户想使用自己的已有域名时，也可以将自定义域名绑定到此服务，在做好备案、与默认域名的 CNAME 后，可直接调用自定义域名。
func (c *Client) BindSubDomain(request *BindSubDomainRequest) (response *BindSubDomainResponse, err error) {
	if request == nil {
		request = NewBindSubDomainRequest()
	}
	response = NewBindSubDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAPIDocRequest() (request *DeleteAPIDocRequest) {
	request = &DeleteAPIDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DeleteAPIDoc")
	return
}

func NewDeleteAPIDocResponse() (response *DeleteAPIDocResponse) {
	response = &DeleteAPIDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除 API 文档
func (c *Client) DeleteAPIDoc(request *DeleteAPIDocRequest) (response *DeleteAPIDocResponse, err error) {
	if request == nil {
		request = NewDeleteAPIDocRequest()
	}
	response = NewDeleteAPIDocResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFreePkgResourceRequest() (request *CreateFreePkgResourceRequest) {
	request = &CreateFreePkgResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateFreePkgResource")
	return
}

func NewCreateFreePkgResourceResponse() (response *CreateFreePkgResourceResponse) {
	response = &CreateFreePkgResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口CreateFreePkgResource用于开通免费额度资源包
func (c *Client) CreateFreePkgResource(request *CreateFreePkgResourceRequest) (response *CreateFreePkgResourceResponse, err error) {
	if request == nil {
		request = NewCreateFreePkgResourceRequest()
	}
	response = NewCreateFreePkgResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogSearchRequest() (request *DescribeLogSearchRequest) {
	request = &DescribeLogSearchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeLogSearch")
	return
}

func NewDescribeLogSearchResponse() (response *DescribeLogSearchResponse) {
	response = &DescribeLogSearchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口DescribeLogSearch用于搜索日志
func (c *Client) DescribeLogSearch(request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
	if request == nil {
		request = NewDescribeLogSearchRequest()
	}
	response = NewDescribeLogSearchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIPStrategyRequest() (request *ModifyIPStrategyRequest) {
	request = &ModifyIPStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyIPStrategy")
	return
}

func NewModifyIPStrategyResponse() (response *ModifyIPStrategyResponse) {
	response = &ModifyIPStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyIPStrategy）用于修改服务IP策略。
func (c *Client) ModifyIPStrategy(request *ModifyIPStrategyRequest) (response *ModifyIPStrategyResponse, err error) {
	if request == nil {
		request = NewModifyIPStrategyRequest()
	}
	response = NewModifyIPStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewRunApiRequest() (request *RunApiRequest) {
	request = &RunApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "RunApi")
	return
}

func NewRunApiResponse() (response *RunApiResponse) {
	response = &RunApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RunApi）用于调试 API 接口。用户在配置完成后可调用此接口进行调试，无需等到发布后走正式的调用流程。
func (c *Client) RunApi(request *RunApiRequest) (response *RunApiResponse, err error) {
	if request == nil {
		request = NewRunApiRequest()
	}
	response = NewRunApiResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApiAppRequest() (request *CreateApiAppRequest) {
	request = &CreateApiAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "CreateApiApp")
	return
}

func NewCreateApiAppResponse() (response *CreateApiAppResponse) {
	response = &CreateApiAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateApiApp）用于创建应用。
func (c *Client) CreateApiApp(request *CreateApiAppRequest) (response *CreateApiAppResponse, err error) {
	if request == nil {
		request = NewCreateApiAppRequest()
	}
	response = NewCreateApiAppResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceEnvironmentListRequest() (request *DescribeServiceEnvironmentListRequest) {
	request = &DescribeServiceEnvironmentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentList")
	return
}

func NewDescribeServiceEnvironmentListResponse() (response *DescribeServiceEnvironmentListResponse) {
	response = &DescribeServiceEnvironmentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceEnvironmentList）用于查询一个服务的环境列表，可查询到此服务下所有环境及其状态。
func (c *Client) DescribeServiceEnvironmentList(request *DescribeServiceEnvironmentListRequest) (response *DescribeServiceEnvironmentListResponse, err error) {
	if request == nil {
		request = NewDescribeServiceEnvironmentListRequest()
	}
	response = NewDescribeServiceEnvironmentListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceSubDomainsRequest() (request *DescribeServiceSubDomainsRequest) {
	request = &DescribeServiceSubDomainsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomains")
	return
}

func NewDescribeServiceSubDomainsResponse() (response *DescribeServiceSubDomainsResponse) {
	response = &DescribeServiceSubDomainsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceSubDomains）用于查询自定义域名列表。
// API 网关可绑定自定义域名到服务，用于服务调用。此接口用于查询用户绑定在服务的自定义域名列表。
func (c *Client) DescribeServiceSubDomains(request *DescribeServiceSubDomainsRequest) (response *DescribeServiceSubDomainsResponse, err error) {
	if request == nil {
		request = NewDescribeServiceSubDomainsRequest()
	}
	response = NewDescribeServiceSubDomainsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApiAppKeyRequest() (request *UpdateApiAppKeyRequest) {
	request = &UpdateApiAppKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UpdateApiAppKey")
	return
}

func NewUpdateApiAppKeyResponse() (response *UpdateApiAppKeyResponse) {
	response = &UpdateApiAppKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateApiAppKey）用于更新应用秘钥。
func (c *Client) UpdateApiAppKey(request *UpdateApiAppKeyRequest) (response *UpdateApiAppKeyResponse, err error) {
	if request == nil {
		request = NewUpdateApiAppKeyRequest()
	}
	response = NewUpdateApiAppKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsagePlanRequest() (request *DescribeUsagePlanRequest) {
	request = &DescribeUsagePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlan")
	return
}

func NewDescribeUsagePlanResponse() (response *DescribeUsagePlanResponse) {
	response = &DescribeUsagePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeUsagePlan）用于查询一个使用计划的详细信息，包括名称、QPS、创建时间绑定的环境等。
func (c *Client) DescribeUsagePlan(request *DescribeUsagePlanRequest) (response *DescribeUsagePlanResponse, err error) {
	if request == nil {
		request = NewDescribeUsagePlanRequest()
	}
	response = NewDescribeUsagePlanResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServiceEnvironmentStrategyRequest() (request *ModifyServiceEnvironmentStrategyRequest) {
	request = &ModifyServiceEnvironmentStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "ModifyServiceEnvironmentStrategy")
	return
}

func NewModifyServiceEnvironmentStrategyResponse() (response *ModifyServiceEnvironmentStrategyResponse) {
	response = &ModifyServiceEnvironmentStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyServiceEnvironmentStrategy）用于修改服务限流策略
func (c *Client) ModifyServiceEnvironmentStrategy(request *ModifyServiceEnvironmentStrategyRequest) (response *ModifyServiceEnvironmentStrategyResponse, err error) {
	if request == nil {
		request = NewModifyServiceEnvironmentStrategyRequest()
	}
	response = NewModifyServiceEnvironmentStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewGetAccountSettingsRequest() (request *GetAccountSettingsRequest) {
	request = &GetAccountSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "GetAccountSettings")
	return
}

func NewGetAccountSettingsResponse() (response *GetAccountSettingsResponse) {
	response = &GetAccountSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户配额
func (c *Client) GetAccountSettings(request *GetAccountSettingsRequest) (response *GetAccountSettingsResponse, err error) {
	if request == nil {
		request = NewGetAccountSettingsRequest()
	}
	response = NewGetAccountSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindApiAppRequest() (request *UnbindApiAppRequest) {
	request = &UnbindApiAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apigateway", APIVersion, "UnbindApiApp")
	return
}

func NewUnbindApiAppResponse() (response *UnbindApiAppResponse) {
	response = &UnbindApiAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnbindApiApp）用于解除应用和API绑定。
func (c *Client) UnbindApiApp(request *UnbindApiAppRequest) (response *UnbindApiAppResponse, err error) {
	if request == nil {
		request = NewUnbindApiAppRequest()
	}
	response = NewUnbindApiAppResponse()
	err = c.Send(request, response)
	return
}
