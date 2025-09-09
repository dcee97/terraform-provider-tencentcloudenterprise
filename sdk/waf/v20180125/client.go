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

package v20180125

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-01-25"

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

func NewDescribeAntiInfoLeakageRulesRequest() (request *DescribeAntiInfoLeakageRulesRequest) {
	request = &DescribeAntiInfoLeakageRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiInfoLeakageRules")
	return
}

func NewDescribeAntiInfoLeakageRulesResponse() (response *DescribeAntiInfoLeakageRulesResponse) {
	response = &DescribeAntiInfoLeakageRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取得信息防泄漏规则列表
func (c *Client) DescribeAntiInfoLeakageRules(request *DescribeAntiInfoLeakageRulesRequest) (response *DescribeAntiInfoLeakageRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAntiInfoLeakageRulesRequest()
	}
	response = NewDescribeAntiInfoLeakageRulesResponse()
	err = c.Send(request, response)
	return
}

func NewEnableUserSignaturePolicyRequest() (request *EnableUserSignaturePolicyRequest) {
	request = &EnableUserSignaturePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "EnableUserSignaturePolicy")
	return
}

func NewEnableUserSignaturePolicyResponse() (response *EnableUserSignaturePolicyResponse) {
	response = &EnableUserSignaturePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启用户防护规则开关
func (c *Client) EnableUserSignaturePolicy(request *EnableUserSignaturePolicyRequest) (response *EnableUserSignaturePolicyResponse, err error) {
	if request == nil {
		request = NewEnableUserSignaturePolicyRequest()
	}
	response = NewEnableUserSignaturePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSceneUCBRuleRequest() (request *ModifyBotSceneUCBRuleRequest) {
	request = &ModifyBotSceneUCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotSceneUCBRule")
	return
}

func NewModifyBotSceneUCBRuleResponse() (response *ModifyBotSceneUCBRuleResponse) {
	response = &ModifyBotSceneUCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【接口复用】场景化后更新Bot的UCB自定义规则，两个调用位置：1.BOT全局白名单 2.BOT场景配置
func (c *Client) ModifyBotSceneUCBRule(request *ModifyBotSceneUCBRuleRequest) (response *ModifyBotSceneUCBRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotSceneUCBRuleRequest()
	}
	response = NewModifyBotSceneUCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserDefaultEngineRequest() (request *DescribeUserDefaultEngineRequest) {
	request = &DescribeUserDefaultEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserDefaultEngine")
	return
}

func NewDescribeUserDefaultEngineResponse() (response *DescribeUserDefaultEngineResponse) {
	response = &DescribeUserDefaultEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户缺省规则引擎类型
func (c *Client) DescribeUserDefaultEngine(request *DescribeUserDefaultEngineRequest) (response *DescribeUserDefaultEngineResponse, err error) {
	if request == nil {
		request = NewDescribeUserDefaultEngineRequest()
	}
	response = NewDescribeUserDefaultEngineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessExportRequest() (request *CreateAccessExportRequest) {
	request = &CreateAccessExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateAccessExport")
	return
}

func NewCreateAccessExportResponse() (response *CreateAccessExportResponse) {
	response = &CreateAccessExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建访问日志导出
func (c *Client) CreateAccessExport(request *CreateAccessExportRequest) (response *CreateAccessExportResponse, err error) {
	if request == nil {
		request = NewCreateAccessExportRequest()
	}
	response = NewCreateAccessExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceEventsRequest() (request *DescribeInstanceEventsRequest) {
	request = &DescribeInstanceEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeInstanceEvents")
	return
}

func NewDescribeInstanceEventsResponse() (response *DescribeInstanceEventsResponse) {
	response = &DescribeInstanceEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户所有实例事件的详细信息
func (c *Client) DescribeInstanceEvents(request *DescribeInstanceEventsRequest) (response *DescribeInstanceEventsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceEventsRequest()
	}
	response = NewDescribeInstanceEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainInstanceIDsRequest() (request *DescribeDomainInstanceIDsRequest) {
	request = &DescribeDomainInstanceIDsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainInstanceIDs")
	return
}

func NewDescribeDomainInstanceIDsResponse() (response *DescribeDomainInstanceIDsResponse) {
	response = &DescribeDomainInstanceIDsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取添加某个域名的所有实例id
func (c *Client) DescribeDomainInstanceIDs(request *DescribeDomainInstanceIDsRequest) (response *DescribeDomainInstanceIDsResponse, err error) {
	if request == nil {
		request = NewDescribeDomainInstanceIDsRequest()
	}
	response = NewDescribeDomainInstanceIDsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackOverviewRequest() (request *DescribeAttackOverviewRequest) {
	request = &DescribeAttackOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackOverview")
	return
}

func NewDescribeAttackOverviewResponse() (response *DescribeAttackOverviewResponse) {
	response = &DescribeAttackOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 攻击总览
func (c *Client) DescribeAttackOverview(request *DescribeAttackOverviewRequest) (response *DescribeAttackOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeAttackOverviewRequest()
	}
	response = NewDescribeAttackOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomPayloadsRequest() (request *DescribeCustomPayloadsRequest) {
	request = &DescribeCustomPayloadsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomPayloads")
	return
}

func NewDescribeCustomPayloadsResponse() (response *DescribeCustomPayloadsResponse) {
	response = &DescribeCustomPayloadsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询自定义payload列表
func (c *Client) DescribeCustomPayloads(request *DescribeCustomPayloadsRequest) (response *DescribeCustomPayloadsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomPayloadsRequest()
	}
	response = NewDescribeCustomPayloadsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotLabelRequest() (request *DescribeBotLabelRequest) {
	request = &DescribeBotLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotLabel")
	return
}

func NewDescribeBotLabelResponse() (response *DescribeBotLabelResponse) {
	response = &DescribeBotLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot标签
func (c *Client) DescribeBotLabel(request *DescribeBotLabelRequest) (response *DescribeBotLabelResponse, err error) {
	if request == nil {
		request = NewDescribeBotLabelRequest()
	}
	response = NewDescribeBotLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotSceneUCBRuleRequest() (request *DeleteBotSceneUCBRuleRequest) {
	request = &DeleteBotSceneUCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteBotSceneUCBRule")
	return
}

func NewDeleteBotSceneUCBRuleResponse() (response *DeleteBotSceneUCBRuleResponse) {
	response = &DeleteBotSceneUCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 场景化后删除Bot的UCB自定义规则
func (c *Client) DeleteBotSceneUCBRule(request *DeleteBotSceneUCBRuleRequest) (response *DeleteBotSceneUCBRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBotSceneUCBRuleRequest()
	}
	response = NewDeleteBotSceneUCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoDenyIPRequest() (request *DescribeAutoDenyIPRequest) {
	request = &DescribeAutoDenyIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAutoDenyIP")
	return
}

func NewDescribeAutoDenyIPResponse() (response *DescribeAutoDenyIPResponse) {
	response = &DescribeAutoDenyIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述WAF自动封禁IP详情,对齐自动封堵状态
func (c *Client) DescribeAutoDenyIP(request *DescribeAutoDenyIPRequest) (response *DescribeAutoDenyIPResponse, err error) {
	if request == nil {
		request = NewDescribeAutoDenyIPRequest()
	}
	response = NewDescribeAutoDenyIPResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackPointsRequest() (request *GetAttackPointsRequest) {
	request = &GetAttackPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetAttackPoints")
	return
}

func NewGetAttackPointsResponse() (response *GetAttackPointsResponse) {
	response = &GetAttackPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意量趋势图
func (c *Client) GetAttackPoints(request *GetAttackPointsRequest) (response *GetAttackPointsResponse, err error) {
	if request == nil {
		request = NewGetAttackPointsRequest()
	}
	response = NewGetAttackPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOpRuleUpdateLogRequest() (request *DeleteOpRuleUpdateLogRequest) {
	request = &DeleteOpRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteOpRuleUpdateLog")
	return
}

func NewDeleteOpRuleUpdateLogResponse() (response *DeleteOpRuleUpdateLogResponse) {
	response = &DeleteOpRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：删除规则更新日志
func (c *Client) DeleteOpRuleUpdateLog(request *DeleteOpRuleUpdateLogRequest) (response *DeleteOpRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewDeleteOpRuleUpdateLogRequest()
	}
	response = NewDeleteOpRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAntiFakeUrlRequest() (request *ModifyAntiFakeUrlRequest) {
	request = &ModifyAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiFakeUrl")
	return
}

func NewModifyAntiFakeUrlResponse() (response *ModifyAntiFakeUrlResponse) {
	response = &ModifyAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑防篡改url
func (c *Client) ModifyAntiFakeUrl(request *ModifyAntiFakeUrlRequest) (response *ModifyAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewModifyAntiFakeUrlRequest()
	}
	response = NewModifyAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewWafCreateResourceAfterPayRequest() (request *WafCreateResourceAfterPayRequest) {
	request = &WafCreateResourceAfterPayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "WafCreateResourceAfterPay")
	return
}

func NewWafCreateResourceAfterPayResponse() (response *WafCreateResourceAfterPayResponse) {
	response = &WafCreateResourceAfterPayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 购买WAF
func (c *Client) WafCreateResourceAfterPay(request *WafCreateResourceAfterPayRequest) (response *WafCreateResourceAfterPayResponse, err error) {
	if request == nil {
		request = NewWafCreateResourceAfterPayRequest()
	}
	response = NewWafCreateResourceAfterPayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotDetailCountTopNRequest() (request *DescribeBotDetailCountTopNRequest) {
	request = &DescribeBotDetailCountTopNRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotDetailCountTopN")
	return
}

func NewDescribeBotDetailCountTopNResponse() (response *DescribeBotDetailCountTopNResponse) {
	response = &DescribeBotDetailCountTopNResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bot详情中请求数量topN
func (c *Client) DescribeBotDetailCountTopN(request *DescribeBotDetailCountTopNRequest) (response *DescribeBotDetailCountTopNResponse, err error) {
	if request == nil {
		request = NewDescribeBotDetailCountTopNRequest()
	}
	response = NewDescribeBotDetailCountTopNResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiSecSensitiveRuleRequest() (request *ModifyApiSecSensitiveRuleRequest) {
	request = &ModifyApiSecSensitiveRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyApiSecSensitiveRule")
	return
}

func NewModifyApiSecSensitiveRuleResponse() (response *ModifyApiSecSensitiveRuleResponse) {
	response = &ModifyApiSecSensitiveRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改api安全敏感检测规则
func (c *Client) ModifyApiSecSensitiveRule(request *ModifyApiSecSensitiveRuleRequest) (response *ModifyApiSecSensitiveRuleResponse, err error) {
	if request == nil {
		request = NewModifyApiSecSensitiveRuleRequest()
	}
	response = NewModifyApiSecSensitiveRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessCursorRequest() (request *DescribeAccessCursorRequest) {
	request = &DescribeAccessCursorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessCursor")
	return
}

func NewDescribeAccessCursorResponse() (response *DescribeAccessCursorResponse) {
	response = &DescribeAccessCursorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据指定的时间来获取访问日志的游标
func (c *Client) DescribeAccessCursor(request *DescribeAccessCursorRequest) (response *DescribeAccessCursorResponse, err error) {
	if request == nil {
		request = NewDescribeAccessCursorRequest()
	}
	response = NewDescribeAccessCursorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessFastAnalysisRequest() (request *DescribeAccessFastAnalysisRequest) {
	request = &DescribeAccessFastAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessFastAnalysis")
	return
}

func NewDescribeAccessFastAnalysisResponse() (response *DescribeAccessFastAnalysisResponse) {
	response = &DescribeAccessFastAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于访问日志的快速分析
func (c *Client) DescribeAccessFastAnalysis(request *DescribeAccessFastAnalysisRequest) (response *DescribeAccessFastAnalysisResponse, err error) {
	if request == nil {
		request = NewDescribeAccessFastAnalysisRequest()
	}
	response = NewDescribeAccessFastAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewFreshAntiFakeUrlRequest() (request *FreshAntiFakeUrlRequest) {
	request = &FreshAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "FreshAntiFakeUrl")
	return
}

func NewFreshAntiFakeUrlResponse() (response *FreshAntiFakeUrlResponse) {
	response = &FreshAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 刷新防篡改url
func (c *Client) FreshAntiFakeUrl(request *FreshAntiFakeUrlRequest) (response *FreshAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewFreshAntiFakeUrlRequest()
	}
	response = NewFreshAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostsRequest() (request *DescribeHostsRequest) {
	request = &DescribeHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeHosts")
	return
}

func NewDescribeHostsResponse() (response *DescribeHostsResponse) {
	response = &DescribeHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf中获取防护域名列表
func (c *Client) DescribeHosts(request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
	if request == nil {
		request = NewDescribeHostsRequest()
	}
	response = NewDescribeHostsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBypassAllStatusRequest() (request *QueryBypassAllStatusRequest) {
	request = &QueryBypassAllStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "QueryBypassAllStatus")
	return
}

func NewQueryBypassAllStatusResponse() (response *QueryBypassAllStatusResponse) {
	response = &QueryBypassAllStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询该用户是否被加入了全局的bypass列表
func (c *Client) QueryBypassAllStatus(request *QueryBypassAllStatusRequest) (response *QueryBypassAllStatusResponse, err error) {
	if request == nil {
		request = NewQueryBypassAllStatusRequest()
	}
	response = NewQueryBypassAllStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackTypeRequest() (request *DescribeAttackTypeRequest) {
	request = &DescribeAttackTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackType")
	return
}

func NewDescribeAttackTypeResponse() (response *DescribeAttackTypeResponse) {
	response = &DescribeAttackTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定域名TOP N攻击类型
func (c *Client) DescribeAttackType(request *DescribeAttackTypeRequest) (response *DescribeAttackTypeResponse, err error) {
	if request == nil {
		request = NewDescribeAttackTypeRequest()
	}
	response = NewDescribeAttackTypeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWhiteListRequest() (request *CreateWhiteListRequest) {
	request = &CreateWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateWhiteList")
	return
}

func NewCreateWhiteListResponse() (response *CreateWhiteListResponse) {
	response = &CreateWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增业务白名单，目前是前端对抗规则中的白名单
func (c *Client) CreateWhiteList(request *CreateWhiteListRequest) (response *CreateWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateWhiteListRequest()
	}
	response = NewCreateWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWafAutoDenyStatusRequest() (request *ModifyWafAutoDenyStatusRequest) {
	request = &ModifyWafAutoDenyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyWafAutoDenyStatus")
	return
}

func NewModifyWafAutoDenyStatusResponse() (response *ModifyWafAutoDenyStatusResponse) {
	response = &ModifyWafAutoDenyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置WAF自动封禁模块状态
func (c *Client) ModifyWafAutoDenyStatus(request *ModifyWafAutoDenyStatusRequest) (response *ModifyWafAutoDenyStatusResponse, err error) {
	if request == nil {
		request = NewModifyWafAutoDenyStatusRequest()
	}
	response = NewModifyWafAutoDenyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDetailRequest() (request *DescribeAssetDetailRequest) {
	request = &DescribeAssetDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAssetDetail")
	return
}

func NewDescribeAssetDetailResponse() (response *DescribeAssetDetailResponse) {
	response = &DescribeAssetDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的资产信息
func (c *Client) DescribeAssetDetail(request *DescribeAssetDetailRequest) (response *DescribeAssetDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDetailRequest()
	}
	response = NewDescribeAssetDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotActionRuleRequest() (request *DescribeBotActionRuleRequest) {
	request = &DescribeBotActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotActionRule")
	return
}

func NewDescribeBotActionRuleResponse() (response *DescribeBotActionRuleResponse) {
	response = &DescribeBotActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot用户配置的动作规则分布
func (c *Client) DescribeBotActionRule(request *DescribeBotActionRuleRequest) (response *DescribeBotActionRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotActionRuleRequest()
	}
	response = NewDescribeBotActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCertificatedDomainRequest() (request *DescribeCertificatedDomainRequest) {
	request = &DescribeCertificatedDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCertificatedDomain")
	return
}

func NewDescribeCertificatedDomainResponse() (response *DescribeCertificatedDomainResponse) {
	response = &DescribeCertificatedDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过ssl证书id获取可以绑定的waf中的域名
func (c *Client) DescribeCertificatedDomain(request *DescribeCertificatedDomainRequest) (response *DescribeCertificatedDomainResponse, err error) {
	if request == nil {
		request = NewDescribeCertificatedDomainRequest()
	}
	response = NewDescribeCertificatedDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafThreatenIntelligenceRequest() (request *DescribeWafThreatenIntelligenceRequest) {
	request = &DescribeWafThreatenIntelligenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeWafThreatenIntelligence")
	return
}

func NewDescribeWafThreatenIntelligenceResponse() (response *DescribeWafThreatenIntelligenceResponse) {
	response = &DescribeWafThreatenIntelligenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述WAF威胁情报封禁模块配置详情
func (c *Client) DescribeWafThreatenIntelligence(request *DescribeWafThreatenIntelligenceRequest) (response *DescribeWafThreatenIntelligenceResponse, err error) {
	if request == nil {
		request = NewDescribeWafThreatenIntelligenceRequest()
	}
	response = NewDescribeWafThreatenIntelligenceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyModuleStatusRequest() (request *ModifyModuleStatusRequest) {
	request = &ModifyModuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyModuleStatus")
	return
}

func NewModifyModuleStatusResponse() (response *ModifyModuleStatusResponse) {
	response = &ModifyModuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置某个domain下基础安全模块的开关
func (c *Client) ModifyModuleStatus(request *ModifyModuleStatusRequest) (response *ModifyModuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyModuleStatusRequest()
	}
	response = NewModifyModuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTlsVersionRequest() (request *DescribeTlsVersionRequest) {
	request = &DescribeTlsVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeTlsVersion")
	return
}

func NewDescribeTlsVersionResponse() (response *DescribeTlsVersionResponse) {
	response = &DescribeTlsVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户TLS版本
func (c *Client) DescribeTlsVersion(request *DescribeTlsVersionRequest) (response *DescribeTlsVersionResponse, err error) {
	if request == nil {
		request = NewDescribeTlsVersionRequest()
	}
	response = NewDescribeTlsVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePanDomainDetailRequest() (request *DescribePanDomainDetailRequest) {
	request = &DescribePanDomainDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePanDomainDetail")
	return
}

func NewDescribePanDomainDetailResponse() (response *DescribePanDomainDetailResponse) {
	response = &DescribePanDomainDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询泛域名非标支持详情
func (c *Client) DescribePanDomainDetail(request *DescribePanDomainDetailRequest) (response *DescribePanDomainDetailResponse, err error) {
	if request == nil {
		request = NewDescribePanDomainDetailRequest()
	}
	response = NewDescribePanDomainDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRolesRequest() (request *DescribeRolesRequest) {
	request = &DescribeRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeRoles")
	return
}

func NewDescribeRolesResponse() (response *DescribeRolesResponse) {
	response = &DescribeRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查CAM角色服务是否开通
func (c *Client) DescribeRoles(request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
	if request == nil {
		request = NewDescribeRolesRequest()
	}
	response = NewDescribeRolesResponse()
	err = c.Send(request, response)
	return
}

func NewAddBlockPageRequest() (request *AddBlockPageRequest) {
	request = &AddBlockPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddBlockPage")
	return
}

func NewAddBlockPageResponse() (response *AddBlockPageResponse) {
	response = &AddBlockPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加自定义拦截页面，每个用户最多允许添加5个自定义拦截页面，添加完后需要审核通过才可以应用到域名。
func (c *Client) AddBlockPage(request *AddBlockPageRequest) (response *AddBlockPageResponse, err error) {
	if request == nil {
		request = NewAddBlockPageRequest()
	}
	response = NewAddBlockPageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProtectionLevelRequest() (request *ModifyProtectionLevelRequest) {
	request = &ModifyProtectionLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyProtectionLevel")
	return
}

func NewModifyProtectionLevelResponse() (response *ModifyProtectionLevelResponse) {
	response = &ModifyProtectionLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改防护等级
func (c *Client) ModifyProtectionLevel(request *ModifyProtectionLevelRequest) (response *ModifyProtectionLevelResponse, err error) {
	if request == nil {
		request = NewModifyProtectionLevelRequest()
	}
	response = NewModifyProtectionLevelResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserSignatureRuleRequest() (request *ModifyUserSignatureRuleRequest) {
	request = &ModifyUserSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyUserSignatureRule")
	return
}

func NewModifyUserSignatureRuleResponse() (response *ModifyUserSignatureRuleResponse) {
	response = &ModifyUserSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户防护规则，开启关闭具体的某条规则
func (c *Client) ModifyUserSignatureRule(request *ModifyUserSignatureRuleRequest) (response *ModifyUserSignatureRuleResponse, err error) {
	if request == nil {
		request = NewModifyUserSignatureRuleRequest()
	}
	response = NewModifyUserSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewRefreshAccessCheckResultRequest() (request *RefreshAccessCheckResultRequest) {
	request = &RefreshAccessCheckResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "RefreshAccessCheckResult")
	return
}

func NewRefreshAccessCheckResultResponse() (response *RefreshAccessCheckResultResponse) {
	response = &RefreshAccessCheckResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 刷新接入检查的结果，后台会生成接入检查任务
func (c *Client) RefreshAccessCheckResult(request *RefreshAccessCheckResultRequest) (response *RefreshAccessCheckResultResponse, err error) {
	if request == nil {
		request = NewRefreshAccessCheckResultRequest()
	}
	response = NewRefreshAccessCheckResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTokenRequest() (request *DescribeBotTokenRequest) {
	request = &DescribeBotTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotToken")
	return
}

func NewDescribeBotTokenResponse() (response *DescribeBotTokenResponse) {
	response = &DescribeBotTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot用户自定义会话配置列表
func (c *Client) DescribeBotToken(request *DescribeBotTokenRequest) (response *DescribeBotTokenResponse, err error) {
	if request == nil {
		request = NewDescribeBotTokenRequest()
	}
	response = NewDescribeBotTokenResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotAutoStatisticRuleRequest() (request *ModifyBotAutoStatisticRuleRequest) {
	request = &ModifyBotAutoStatisticRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotAutoStatisticRule")
	return
}

func NewModifyBotAutoStatisticRuleResponse() (response *ModifyBotAutoStatisticRuleResponse) {
	response = &ModifyBotAutoStatisticRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot-智能统计分析模块配置
func (c *Client) ModifyBotAutoStatisticRule(request *ModifyBotAutoStatisticRuleRequest) (response *ModifyBotAutoStatisticRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotAutoStatisticRuleRequest()
	}
	response = NewModifyBotAutoStatisticRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomPayloadsRequest() (request *DeleteCustomPayloadsRequest) {
	request = &DeleteCustomPayloadsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteCustomPayloads")
	return
}

func NewDeleteCustomPayloadsResponse() (response *DeleteCustomPayloadsResponse) {
	response = &DeleteCustomPayloadsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义payload
func (c *Client) DeleteCustomPayloads(request *DeleteCustomPayloadsRequest) (response *DeleteCustomPayloadsResponse, err error) {
	if request == nil {
		request = NewDeleteCustomPayloadsRequest()
	}
	response = NewDeleteCustomPayloadsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpAttackWhiteRuleRequest() (request *DescribeOpAttackWhiteRuleRequest) {
	request = &DescribeOpAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeOpAttackWhiteRule")
	return
}

func NewDescribeOpAttackWhiteRuleResponse() (response *DescribeOpAttackWhiteRuleResponse) {
	response = &DescribeOpAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营平台接口：获取用户规则白名单列表
func (c *Client) DescribeOpAttackWhiteRule(request *DescribeOpAttackWhiteRuleRequest) (response *DescribeOpAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeOpAttackWhiteRuleRequest()
	}
	response = NewDescribeOpAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostAccessLogStatusRequest() (request *ModifyHostAccessLogStatusRequest) {
	request = &ModifyHostAccessLogStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyHostAccessLogStatus")
	return
}

func NewModifyHostAccessLogStatusResponse() (response *ModifyHostAccessLogStatusResponse) {
	response = &ModifyHostAccessLogStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置单个域名是否开启记录访问日志功能的开关状态
func (c *Client) ModifyHostAccessLogStatus(request *ModifyHostAccessLogStatusRequest) (response *ModifyHostAccessLogStatusResponse, err error) {
	if request == nil {
		request = NewModifyHostAccessLogStatusRequest()
	}
	response = NewModifyHostAccessLogStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserSceneInfoRequest() (request *DescribeUserSceneInfoRequest) {
	request = &DescribeUserSceneInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserSceneInfo")
	return
}

func NewDescribeUserSceneInfoResponse() (response *DescribeUserSceneInfoResponse) {
	response = &DescribeUserSceneInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名下的场景and动作策略列表
func (c *Client) DescribeUserSceneInfo(request *DescribeUserSceneInfoRequest) (response *DescribeUserSceneInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserSceneInfoRequest()
	}
	response = NewDescribeUserSceneInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFindDomainListRequest() (request *DescribeFindDomainListRequest) {
	request = &DescribeFindDomainListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeFindDomainList")
	return
}

func NewDescribeFindDomainListResponse() (response *DescribeFindDomainListResponse) {
	response = &DescribeFindDomainListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取发现域名列表接口
func (c *Client) DescribeFindDomainList(request *DescribeFindDomainListRequest) (response *DescribeFindDomainListResponse, err error) {
	if request == nil {
		request = NewDescribeFindDomainListRequest()
	}
	response = NewDescribeFindDomainListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainWhiteRulesRequest() (request *DescribeDomainWhiteRulesRequest) {
	request = &DescribeDomainWhiteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainWhiteRules")
	return
}

func NewDescribeDomainWhiteRulesResponse() (response *DescribeDomainWhiteRulesResponse) {
	response = &DescribeDomainWhiteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名的规则白名单
func (c *Client) DescribeDomainWhiteRules(request *DescribeDomainWhiteRulesRequest) (response *DescribeDomainWhiteRulesResponse, err error) {
	if request == nil {
		request = NewDescribeDomainWhiteRulesRequest()
	}
	response = NewDescribeDomainWhiteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDNSDetectDomainRequest() (request *ModifyDNSDetectDomainRequest) {
	request = &ModifyDNSDetectDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyDNSDetectDomain")
	return
}

func NewModifyDNSDetectDomainResponse() (response *ModifyDNSDetectDomainResponse) {
	response = &ModifyDNSDetectDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑DNS劫持检测的域名，对应【Web安全检测】【DNS劫持检测】【基础设置】【编辑】这个功能。
func (c *Client) ModifyDNSDetectDomain(request *ModifyDNSDetectDomainRequest) (response *ModifyDNSDetectDomainResponse, err error) {
	if request == nil {
		request = NewModifyDNSDetectDomainRequest()
	}
	response = NewModifyDNSDetectDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSceneJsInjectRuleRequest() (request *DescribeSceneJsInjectRuleRequest) {
	request = &DescribeSceneJsInjectRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeSceneJsInjectRule")
	return
}

func NewDescribeSceneJsInjectRuleResponse() (response *DescribeSceneJsInjectRuleResponse) {
	response = &DescribeSceneJsInjectRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【接口复用】获取对应域名的前端对抗规则，查询全局的前端对抗配置时IsGlobal==1，返回的JsInjectRule中只含有全局设置对应的字段；查询场景的前端对抗配置时IsGlobal==0，返回的JsInjectRule中只含有场景设置对应的字段。
func (c *Client) DescribeSceneJsInjectRule(request *DescribeSceneJsInjectRuleRequest) (response *DescribeSceneJsInjectRuleResponse, err error) {
	if request == nil {
		request = NewDescribeSceneJsInjectRuleRequest()
	}
	response = NewDescribeSceneJsInjectRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSDetectDataMapRequest() (request *DescribeDNSDetectDataMapRequest) {
	request = &DescribeDNSDetectDataMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDNSDetectDataMap")
	return
}

func NewDescribeDNSDetectDataMapResponse() (response *DescribeDNSDetectDataMapResponse) {
	response = &DescribeDNSDetectDataMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DNS劫持-获取地图数据
func (c *Client) DescribeDNSDetectDataMap(request *DescribeDNSDetectDataMapRequest) (response *DescribeDNSDetectDataMapResponse, err error) {
	if request == nil {
		request = NewDescribeDNSDetectDataMapRequest()
	}
	response = NewDescribeDNSDetectDataMapResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowTrendRequest() (request *DescribeFlowTrendRequest) {
	request = &DescribeFlowTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeFlowTrend")
	return
}

func NewDescribeFlowTrendResponse() (response *DescribeFlowTrendResponse) {
	response = &DescribeFlowTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取waf流量访问趋势
func (c *Client) DescribeFlowTrend(request *DescribeFlowTrendRequest) (response *DescribeFlowTrendResponse, err error) {
	if request == nil {
		request = NewDescribeFlowTrendRequest()
	}
	response = NewDescribeFlowTrendResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpartUserInfoRequest() (request *DescribeSpartUserInfoRequest) {
	request = &DescribeSpartUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeSpartUserInfo")
	return
}

func NewDescribeSpartUserInfoResponse() (response *DescribeSpartUserInfoResponse) {
	response = &DescribeSpartUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户信息，判断用户是否开通waf，套餐购买情况等信息
func (c *Client) DescribeSpartUserInfo(request *DescribeSpartUserInfoRequest) (response *DescribeSpartUserInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSpartUserInfoRequest()
	}
	response = NewDescribeSpartUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomRulesRequest() (request *DescribeCustomRulesRequest) {
	request = &DescribeCustomRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomRules")
	return
}

func NewDescribeCustomRulesResponse() (response *DescribeCustomRulesResponse) {
	response = &DescribeCustomRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防护配置中的自定义策略列表
func (c *Client) DescribeCustomRules(request *DescribeCustomRulesRequest) (response *DescribeCustomRulesResponse, err error) {
	if request == nil {
		request = NewDescribeCustomRulesRequest()
	}
	response = NewDescribeCustomRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTrafficMarkingRequest() (request *ModifyTrafficMarkingRequest) {
	request = &ModifyTrafficMarkingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyTrafficMarking")
	return
}

func NewModifyTrafficMarkingResponse() (response *ModifyTrafficMarkingResponse) {
	response = &ModifyTrafficMarkingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新流量标记配置
func (c *Client) ModifyTrafficMarking(request *ModifyTrafficMarkingRequest) (response *ModifyTrafficMarkingResponse, err error) {
	if request == nil {
		request = NewModifyTrafficMarkingRequest()
	}
	response = NewModifyTrafficMarkingResponse()
	err = c.Send(request, response)
	return
}

func NewAddBypassAllRuleRequest() (request *AddBypassAllRuleRequest) {
	request = &AddBypassAllRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddBypassAllRule")
	return
}

func NewAddBypassAllRuleResponse() (response *AddBypassAllRuleResponse) {
	response = &AddBypassAllRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加一键bypass能力支持,直接添加APPID
func (c *Client) AddBypassAllRule(request *AddBypassAllRuleRequest) (response *AddBypassAllRuleResponse, err error) {
	if request == nil {
		request = NewAddBypassAllRuleRequest()
	}
	response = NewAddBypassAllRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCiphersDetailRequest() (request *DescribeCiphersDetailRequest) {
	request = &DescribeCiphersDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCiphersDetail")
	return
}

func NewDescribeCiphersDetailResponse() (response *DescribeCiphersDetailResponse) {
	response = &DescribeCiphersDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询加密套件信息
func (c *Client) DescribeCiphersDetail(request *DescribeCiphersDetailRequest) (response *DescribeCiphersDetailResponse, err error) {
	if request == nil {
		request = NewDescribeCiphersDetailRequest()
	}
	response = NewDescribeCiphersDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpUserSignaturePolicyRequest() (request *DescribeOpUserSignaturePolicyRequest) {
	request = &DescribeOpUserSignaturePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeOpUserSignaturePolicy")
	return
}

func NewDescribeOpUserSignaturePolicyResponse() (response *DescribeOpUserSignaturePolicyResponse) {
	response = &DescribeOpUserSignaturePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：获取规则状态
func (c *Client) DescribeOpUserSignaturePolicy(request *DescribeOpUserSignaturePolicyRequest) (response *DescribeOpUserSignaturePolicyResponse, err error) {
	if request == nil {
		request = NewDescribeOpUserSignaturePolicyRequest()
	}
	response = NewDescribeOpUserSignaturePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAccessStatPointsRequest() (request *DescribeBotAccessStatPointsRequest) {
	request = &DescribeBotAccessStatPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotAccessStatPoints")
	return
}

func NewDescribeBotAccessStatPointsResponse() (response *DescribeBotAccessStatPointsResponse) {
	response = &DescribeBotAccessStatPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot4.0的趋势图接口
func (c *Client) DescribeBotAccessStatPoints(request *DescribeBotAccessStatPointsRequest) (response *DescribeBotAccessStatPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotAccessStatPointsRequest()
	}
	response = NewDescribeBotAccessStatPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessHistogramRequest() (request *DescribeAccessHistogramRequest) {
	request = &DescribeAccessHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessHistogram")
	return
}

func NewDescribeAccessHistogramResponse() (response *DescribeAccessHistogramResponse) {
	response = &DescribeAccessHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于访问日志柱状趋势图
func (c *Client) DescribeAccessHistogram(request *DescribeAccessHistogramRequest) (response *DescribeAccessHistogramResponse, err error) {
	if request == nil {
		request = NewDescribeAccessHistogramRequest()
	}
	response = NewDescribeAccessHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSessionKeyRequest() (request *ModifyBotSessionKeyRequest) {
	request = &ModifyBotSessionKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotSessionKey")
	return
}

func NewModifyBotSessionKeyResponse() (response *ModifyBotSessionKeyResponse) {
	response = &ModifyBotSessionKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot session 配置
func (c *Client) ModifyBotSessionKey(request *ModifyBotSessionKeyRequest) (response *ModifyBotSessionKeyResponse, err error) {
	if request == nil {
		request = NewModifyBotSessionKeyRequest()
	}
	response = NewModifyBotSessionKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteInstance")
	return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 释放已经过期的实例
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	response = NewDeleteInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWhiteListRequest() (request *DeleteWhiteListRequest) {
	request = &DeleteWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteWhiteList")
	return
}

func NewDeleteWhiteListResponse() (response *DeleteWhiteListResponse) {
	response = &DeleteWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除业务白名单，目前是前端对抗规则中的白名单
func (c *Client) DeleteWhiteList(request *DeleteWhiteListRequest) (response *DeleteWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteWhiteListRequest()
	}
	response = NewDeleteWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPocTestAuthorizationRequest() (request *ModifyPocTestAuthorizationRequest) {
	request = &ModifyPocTestAuthorizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyPocTestAuthorization")
	return
}

func NewModifyPocTestAuthorizationResponse() (response *ModifyPocTestAuthorizationResponse) {
	response = &ModifyPocTestAuthorizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改APPID授权状态
func (c *Client) ModifyPocTestAuthorization(request *ModifyPocTestAuthorizationRequest) (response *ModifyPocTestAuthorizationResponse, err error) {
	if request == nil {
		request = NewModifyPocTestAuthorizationRequest()
	}
	response = NewModifyPocTestAuthorizationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpRuleUpdateLogRequest() (request *DescribeOpRuleUpdateLogRequest) {
	request = &DescribeOpRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeOpRuleUpdateLog")
	return
}

func NewDescribeOpRuleUpdateLogResponse() (response *DescribeOpRuleUpdateLogResponse) {
	response = &DescribeOpRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台：获取特征规则更新动态
func (c *Client) DescribeOpRuleUpdateLog(request *DescribeOpRuleUpdateLogRequest) (response *DescribeOpRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewDescribeOpRuleUpdateLogRequest()
	}
	response = NewDescribeOpRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotUCBRecordsRequest() (request *DescribeBotUCBRecordsRequest) {
	request = &DescribeBotUCBRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotUCBRecords")
	return
}

func NewDescribeBotUCBRecordsResponse() (response *DescribeBotUCBRecordsResponse) {
	response = &DescribeBotUCBRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取UCB类型Bots
func (c *Client) DescribeBotUCBRecords(request *DescribeBotUCBRecordsRequest) (response *DescribeBotUCBRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeBotUCBRecordsRequest()
	}
	response = NewDescribeBotUCBRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
	request = &ModifyInstanceNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceName")
	return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
	response = &ModifyInstanceNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例的名称
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
	if request == nil {
		request = NewModifyInstanceNameRequest()
	}
	response = NewModifyInstanceNameResponse()
	err = c.Send(request, response)
	return
}

func NewAddAntiFakeUrlRequest() (request *AddAntiFakeUrlRequest) {
	request = &AddAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddAntiFakeUrl")
	return
}

func NewAddAntiFakeUrlResponse() (response *AddAntiFakeUrlResponse) {
	response = &AddAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加防篡改url
func (c *Client) AddAntiFakeUrl(request *AddAntiFakeUrlRequest) (response *AddAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewAddAntiFakeUrlRequest()
	}
	response = NewAddAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebshellStatusRequest() (request *DescribeWebshellStatusRequest) {
	request = &DescribeWebshellStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeWebshellStatus")
	return
}

func NewDescribeWebshellStatusResponse() (response *DescribeWebshellStatusResponse) {
	response = &DescribeWebshellStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名的webshell状态
func (c *Client) DescribeWebshellStatus(request *DescribeWebshellStatusRequest) (response *DescribeWebshellStatusResponse, err error) {
	if request == nil {
		request = NewDescribeWebshellStatusRequest()
	}
	response = NewDescribeWebshellStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetStracePointsRequest() (request *GetStracePointsRequest) {
	request = &GetStracePointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetStracePoints")
	return
}

func NewGetStracePointsResponse() (response *GetStracePointsResponse) {
	response = &GetStracePointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防护量趋势图
func (c *Client) GetStracePoints(request *GetStracePointsRequest) (response *GetStracePointsResponse, err error) {
	if request == nil {
		request = NewGetStracePointsRequest()
	}
	response = NewGetStracePointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackCountRequest() (request *DescribeAttackCountRequest) {
	request = &DescribeAttackCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackCount")
	return
}

func NewDescribeAttackCountResponse() (response *DescribeAttackCountResponse) {
	response = &DescribeAttackCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定域名，指定时间段的攻击数量
func (c *Client) DescribeAttackCount(request *DescribeAttackCountRequest) (response *DescribeAttackCountResponse, err error) {
	if request == nil {
		request = NewDescribeAttackCountRequest()
	}
	response = NewDescribeAttackCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotUCBPreinstallRuleRequest() (request *DescribeBotUCBPreinstallRuleRequest) {
	request = &DescribeBotUCBPreinstallRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotUCBPreinstallRule")
	return
}

func NewDescribeBotUCBPreinstallRuleResponse() (response *DescribeBotUCBPreinstallRuleResponse) {
	response = &DescribeBotUCBPreinstallRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取UCB预设策略
func (c *Client) DescribeBotUCBPreinstallRule(request *DescribeBotUCBPreinstallRuleRequest) (response *DescribeBotUCBPreinstallRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotUCBPreinstallRuleRequest()
	}
	response = NewDescribeBotUCBPreinstallRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomRuleListRequest() (request *DescribeCustomRuleListRequest) {
	request = &DescribeCustomRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomRuleList")
	return
}

func NewDescribeCustomRuleListResponse() (response *DescribeCustomRuleListResponse) {
	response = &DescribeCustomRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防护配置中的访问控制策略列表
func (c *Client) DescribeCustomRuleList(request *DescribeCustomRuleListRequest) (response *DescribeCustomRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeCustomRuleListRequest()
	}
	response = NewDescribeCustomRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewEnableClientMsgRequest() (request *EnableClientMsgRequest) {
	request = &EnableClientMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "EnableClientMsg")
	return
}

func NewEnableClientMsgResponse() (response *EnableClientMsgResponse) {
	response = &EnableClientMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开关开启后，会将客户端的ip和port透传到后端
func (c *Client) EnableClientMsg(request *EnableClientMsgRequest) (response *EnableClientMsgResponse, err error) {
	if request == nil {
		request = NewEnableClientMsgRequest()
	}
	response = NewEnableClientMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRequestCountRequest() (request *DescribeBotRequestCountRequest) {
	request = &DescribeBotRequestCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotRequestCount")
	return
}

func NewDescribeBotRequestCountResponse() (response *DescribeBotRequestCountResponse) {
	response = &DescribeBotRequestCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Bot访问日志的总量
func (c *Client) DescribeBotRequestCount(request *DescribeBotRequestCountRequest) (response *DescribeBotRequestCountResponse, err error) {
	if request == nil {
		request = NewDescribeBotRequestCountRequest()
	}
	response = NewDescribeBotRequestCountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDNSDetectDomainRequest() (request *DeleteDNSDetectDomainRequest) {
	request = &DeleteDNSDetectDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteDNSDetectDomain")
	return
}

func NewDeleteDNSDetectDomainResponse() (response *DeleteDNSDetectDomainResponse) {
	response = &DeleteDNSDetectDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除DNS劫持检测的域名，对应【Web安全检测】【DNS劫持检测】【基础设置】【删除】这个功能。
func (c *Client) DeleteDNSDetectDomain(request *DeleteDNSDetectDomainRequest) (response *DeleteDNSDetectDomainResponse, err error) {
	if request == nil {
		request = NewDeleteDNSDetectDomainRequest()
	}
	response = NewDeleteDNSDetectDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePocTestTasksRequest() (request *DescribePocTestTasksRequest) {
	request = &DescribePocTestTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePocTestTasks")
	return
}

func NewDescribePocTestTasksResponse() (response *DescribePocTestTasksResponse) {
	response = &DescribePocTestTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询POC测试任务
func (c *Client) DescribePocTestTasks(request *DescribePocTestTasksRequest) (response *DescribePocTestTasksResponse, err error) {
	if request == nil {
		request = NewDescribePocTestTasksRequest()
	}
	response = NewDescribePocTestTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAreaBanSupportAreasRequest() (request *DescribeAreaBanSupportAreasRequest) {
	request = &DescribeAreaBanSupportAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAreaBanSupportAreas")
	return
}

func NewDescribeAreaBanSupportAreasResponse() (response *DescribeAreaBanSupportAreasResponse) {
	response = &DescribeAreaBanSupportAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取WAF地域封禁支持的地域列表
func (c *Client) DescribeAreaBanSupportAreas(request *DescribeAreaBanSupportAreasRequest) (response *DescribeAreaBanSupportAreasResponse, err error) {
	if request == nil {
		request = NewDescribeAreaBanSupportAreasRequest()
	}
	response = NewDescribeAreaBanSupportAreasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpartUserRequest() (request *DescribeSpartUserRequest) {
	request = &DescribeSpartUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeSpartUser")
	return
}

func NewDescribeSpartUserResponse() (response *DescribeSpartUserResponse) {
	response = &DescribeSpartUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户信息，判断用户是否开通waf，套餐购买情况等信息
func (c *Client) DescribeSpartUser(request *DescribeSpartUserRequest) (response *DescribeSpartUserResponse, err error) {
	if request == nil {
		request = NewDescribeSpartUserRequest()
	}
	response = NewDescribeSpartUserResponse()
	err = c.Send(request, response)
	return
}

func NewHavePostFlowServiceRequest() (request *HavePostFlowServiceRequest) {
	request = &HavePostFlowServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "HavePostFlowService")
	return
}

func NewHavePostFlowServiceResponse() (response *HavePostFlowServiceResponse) {
	response = &HavePostFlowServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 是否开通投递服务接口
func (c *Client) HavePostFlowService(request *HavePostFlowServiceRequest) (response *HavePostFlowServiceResponse, err error) {
	if request == nil {
		request = NewHavePostFlowServiceRequest()
	}
	response = NewHavePostFlowServiceResponse()
	err = c.Send(request, response)
	return
}

func NewPostAccessDownloadTaskRequest() (request *PostAccessDownloadTaskRequest) {
	request = &PostAccessDownloadTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "PostAccessDownloadTask")
	return
}

func NewPostAccessDownloadTaskResponse() (response *PostAccessDownloadTaskResponse) {
	response = &PostAccessDownloadTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建搜索下载访问日志任务
func (c *Client) PostAccessDownloadTask(request *PostAccessDownloadTaskRequest) (response *PostAccessDownloadTaskResponse, err error) {
	if request == nil {
		request = NewPostAccessDownloadTaskRequest()
	}
	response = NewPostAccessDownloadTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSDetectDomainListRequest() (request *DescribeDNSDetectDomainListRequest) {
	request = &DescribeDNSDetectDomainListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDNSDetectDomainList")
	return
}

func NewDescribeDNSDetectDomainListResponse() (response *DescribeDNSDetectDomainListResponse) {
	response = &DescribeDNSDetectDomainListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DNS劫持检测的域名列表，对应【Web安全检测】【DNS劫持检测】【基础设置】这个功能。
func (c *Client) DescribeDNSDetectDomainList(request *DescribeDNSDetectDomainListRequest) (response *DescribeDNSDetectDomainListResponse, err error) {
	if request == nil {
		request = NewDescribeDNSDetectDomainListRequest()
	}
	response = NewDescribeDNSDetectDomainListResponse()
	err = c.Send(request, response)
	return
}

func NewGetAccessDownloadRecordsRequest() (request *GetAccessDownloadRecordsRequest) {
	request = &GetAccessDownloadRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetAccessDownloadRecords")
	return
}

func NewGetAccessDownloadRecordsResponse() (response *GetAccessDownloadRecordsResponse) {
	response = &GetAccessDownloadRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询下载访问日志任务记录列表
func (c *Client) GetAccessDownloadRecords(request *GetAccessDownloadRecordsRequest) (response *GetAccessDownloadRecordsResponse, err error) {
	if request == nil {
		request = NewGetAccessDownloadRecordsRequest()
	}
	response = NewGetAccessDownloadRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewGetAllQpsConfigRequest() (request *GetAllQpsConfigRequest) {
	request = &GetAllQpsConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetAllQpsConfig")
	return
}

func NewGetAllQpsConfigResponse() (response *GetAllQpsConfigResponse) {
	response = &GetAllQpsConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有套餐类型对应的弹性qps配置
func (c *Client) GetAllQpsConfig(request *GetAllQpsConfigRequest) (response *GetAllQpsConfigResponse, err error) {
	if request == nil {
		request = NewGetAllQpsConfigRequest()
	}
	response = NewGetAllQpsConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopAttackIPRequest() (request *GetTopAttackIPRequest) {
	request = &GetTopAttackIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetTopAttackIP")
	return
}

func NewGetTopAttackIPResponse() (response *GetTopAttackIPResponse) {
	response = &GetTopAttackIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意IP的top10
func (c *Client) GetTopAttackIP(request *GetTopAttackIPRequest) (response *GetTopAttackIPResponse, err error) {
	if request == nil {
		request = NewGetTopAttackIPRequest()
	}
	response = NewGetTopAttackIPResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCopyCustomRuleRequest() (request *CreateCopyCustomRuleRequest) {
	request = &CreateCopyCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateCopyCustomRule")
	return
}

func NewCreateCopyCustomRuleResponse() (response *CreateCopyCustomRuleResponse) {
	response = &CreateCopyCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拷贝自定义规则到其他域名
func (c *Client) CreateCopyCustomRule(request *CreateCopyCustomRuleRequest) (response *CreateCopyCustomRuleResponse, err error) {
	if request == nil {
		request = NewCreateCopyCustomRuleRequest()
	}
	response = NewCreateCopyCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewEditOpRuleUpdateLogStatusRequest() (request *EditOpRuleUpdateLogStatusRequest) {
	request = &EditOpRuleUpdateLogStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "EditOpRuleUpdateLogStatus")
	return
}

func NewEditOpRuleUpdateLogStatusResponse() (response *EditOpRuleUpdateLogStatusResponse) {
	response = &EditOpRuleUpdateLogStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：编辑规则更新日志的发布状态
func (c *Client) EditOpRuleUpdateLogStatus(request *EditOpRuleUpdateLogStatusRequest) (response *EditOpRuleUpdateLogStatusResponse, err error) {
	if request == nil {
		request = NewEditOpRuleUpdateLogStatusRequest()
	}
	response = NewEditOpRuleUpdateLogStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTsResourceRequest() (request *ModifyTsResourceRequest) {
	request = &ModifyTsResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyTsResource")
	return
}

func NewModifyTsResourceResponse() (response *ModifyTsResourceResponse) {
	response = &ModifyTsResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Tiga资源列表，如果id=0则添加，>0则修改
func (c *Client) ModifyTsResource(request *ModifyTsResourceRequest) (response *ModifyTsResourceResponse, err error) {
	if request == nil {
		request = NewModifyTsResourceRequest()
	}
	response = NewModifyTsResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotUCBFeatureRuleRequest() (request *DescribeBotUCBFeatureRuleRequest) {
	request = &DescribeBotUCBFeatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotUCBFeatureRule")
	return
}

func NewDescribeBotUCBFeatureRuleResponse() (response *DescribeBotUCBFeatureRuleResponse) {
	response = &DescribeBotUCBFeatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取UCB自定义策略
func (c *Client) DescribeBotUCBFeatureRule(request *DescribeBotUCBFeatureRuleRequest) (response *DescribeBotUCBFeatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotUCBFeatureRuleRequest()
	}
	response = NewDescribeBotUCBFeatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAntiFakeUrlRequest() (request *DeleteAntiFakeUrlRequest) {
	request = &DeleteAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteAntiFakeUrl")
	return
}

func NewDeleteAntiFakeUrlResponse() (response *DeleteAntiFakeUrlResponse) {
	response = &DeleteAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除防篡改url
func (c *Client) DeleteAntiFakeUrl(request *DeleteAntiFakeUrlRequest) (response *DeleteAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewDeleteAntiFakeUrlRequest()
	}
	response = NewDeleteAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGenerateDealsRequest() (request *ModifyGenerateDealsRequest) {
	request = &ModifyGenerateDealsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyGenerateDeals")
	return
}

func NewModifyGenerateDealsResponse() (response *ModifyGenerateDealsResponse) {
	response = &ModifyGenerateDealsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给clb等使用的waf实例下单接口，目前只支持clb旗舰版实例的下单，该接口会进行入参校验，然后调用是否为收购用户，然后调用计费接口下单。目前只支持预付费下单
func (c *Client) ModifyGenerateDeals(request *ModifyGenerateDealsRequest) (response *ModifyGenerateDealsResponse, err error) {
	if request == nil {
		request = NewModifyGenerateDealsRequest()
	}
	response = NewModifyGenerateDealsResponse()
	err = c.Send(request, response)
	return
}

func NewAddAttackWhiteRuleRequest() (request *AddAttackWhiteRuleRequest) {
	request = &AddAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddAttackWhiteRule")
	return
}

func NewAddAttackWhiteRuleResponse() (response *AddAttackWhiteRuleResponse) {
	response = &AddAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供用户控制台调用，增加Tiga规则引擎白名单。
func (c *Client) AddAttackWhiteRule(request *AddAttackWhiteRuleRequest) (response *AddAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddAttackWhiteRuleRequest()
	}
	response = NewAddAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddPocTestTaskRequest() (request *AddPocTestTaskRequest) {
	request = &AddPocTestTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddPocTestTask")
	return
}

func NewAddPocTestTaskResponse() (response *AddPocTestTaskResponse) {
	response = &AddPocTestTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建POC测试任务
func (c *Client) AddPocTestTask(request *AddPocTestTaskRequest) (response *AddPocTestTaskResponse, err error) {
	if request == nil {
		request = NewAddPocTestTaskRequest()
	}
	response = NewAddPocTestTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotTiRuleRequest() (request *ModifyBotTiRuleRequest) {
	request = &ModifyBotTiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotTiRule")
	return
}

func NewModifyBotTiRuleResponse() (response *ModifyBotTiRuleResponse) {
	response = &ModifyBotTiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot-威胁情报配置
func (c *Client) ModifyBotTiRule(request *ModifyBotTiRuleRequest) (response *ModifyBotTiRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotTiRuleRequest()
	}
	response = NewModifyBotTiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserEditionRequest() (request *DescribeUserEditionRequest) {
	request = &DescribeUserEditionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserEdition")
	return
}

func NewDescribeUserEditionResponse() (response *DescribeUserEditionResponse) {
	response = &DescribeUserEditionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户购买的WAF的类型，客户是购买的SaaS WAF还是CLB WAF。
func (c *Client) DescribeUserEdition(request *DescribeUserEditionRequest) (response *DescribeUserEditionResponse, err error) {
	if request == nil {
		request = NewDescribeUserEditionRequest()
	}
	response = NewDescribeUserEditionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户所有实例的详细信息
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllDomainRequest() (request *DescribeAllDomainRequest) {
	request = &DescribeAllDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAllDomain")
	return
}

func NewDescribeAllDomainResponse() (response *DescribeAllDomainResponse) {
	response = &DescribeAllDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取clbwaf和saaswaf所有可用域名列表及域名cls开启状态
func (c *Client) DescribeAllDomain(request *DescribeAllDomainRequest) (response *DescribeAllDomainResponse, err error) {
	if request == nil {
		request = NewDescribeAllDomainRequest()
	}
	response = NewDescribeAllDomainResponse()
	err = c.Send(request, response)
	return
}

func NewAddApiRuleRequest() (request *AddApiRuleRequest) {
	request = &AddApiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddApiRule")
	return
}

func NewAddApiRuleResponse() (response *AddApiRuleResponse) {
	response = &AddApiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户手动录入api规则
func (c *Client) AddApiRule(request *AddApiRuleRequest) (response *AddApiRuleResponse, err error) {
	if request == nil {
		request = NewAddApiRuleRequest()
	}
	response = NewAddApiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApiRulesRequest() (request *DeleteApiRulesRequest) {
	request = &DeleteApiRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteApiRules")
	return
}

func NewDeleteApiRulesResponse() (response *DeleteApiRulesResponse) {
	response = &DeleteApiRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户删除api规则,支持多条删除
func (c *Client) DeleteApiRules(request *DeleteApiRulesRequest) (response *DeleteApiRulesResponse, err error) {
	if request == nil {
		request = NewDeleteApiRulesRequest()
	}
	response = NewDeleteApiRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrafficMarkingRequest() (request *DescribeTrafficMarkingRequest) {
	request = &DescribeTrafficMarkingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeTrafficMarking")
	return
}

func NewDescribeTrafficMarkingResponse() (response *DescribeTrafficMarkingResponse) {
	response = &DescribeTrafficMarkingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询流量标记配置
func (c *Client) DescribeTrafficMarking(request *DescribeTrafficMarkingRequest) (response *DescribeTrafficMarkingResponse, err error) {
	if request == nil {
		request = NewDescribeTrafficMarkingRequest()
	}
	response = NewDescribeTrafficMarkingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCCOpenCustomCountRequest() (request *DescribeCCOpenCustomCountRequest) {
	request = &DescribeCCOpenCustomCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCCOpenCustomCount")
	return
}

func NewDescribeCCOpenCustomCountResponse() (response *DescribeCCOpenCustomCountResponse) {
	response = &DescribeCCOpenCustomCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 斯巴达版本查询自定义cc开启状态数目
func (c *Client) DescribeCCOpenCustomCount(request *DescribeCCOpenCustomCountRequest) (response *DescribeCCOpenCustomCountResponse, err error) {
	if request == nil {
		request = NewDescribeCCOpenCustomCountRequest()
	}
	response = NewDescribeCCOpenCustomCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAccessLogTwoRequest() (request *DescribeBotAccessLogTwoRequest) {
	request = &DescribeBotAccessLogTwoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotAccessLogTwo")
	return
}

func NewDescribeBotAccessLogTwoResponse() (response *DescribeBotAccessLogTwoResponse) {
	response = &DescribeBotAccessLogTwoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot详情页日志列表显示v2版本
func (c *Client) DescribeBotAccessLogTwo(request *DescribeBotAccessLogTwoRequest) (response *DescribeBotAccessLogTwoResponse, err error) {
	if request == nil {
		request = NewDescribeBotAccessLogTwoRequest()
	}
	response = NewDescribeBotAccessLogTwoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainIpv6StatusRequest() (request *ModifyDomainIpv6StatusRequest) {
	request = &ModifyDomainIpv6StatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainIpv6Status")
	return
}

func NewModifyDomainIpv6StatusResponse() (response *ModifyDomainIpv6StatusResponse) {
	response = &ModifyDomainIpv6StatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改ipv6开关
func (c *Client) ModifyDomainIpv6Status(request *ModifyDomainIpv6StatusRequest) (response *ModifyDomainIpv6StatusResponse, err error) {
	if request == nil {
		request = NewModifyDomainIpv6StatusRequest()
	}
	response = NewModifyDomainIpv6StatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHostRequest() (request *DeleteHostRequest) {
	request = &DeleteHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteHost")
	return
}

func NewDeleteHostResponse() (response *DeleteHostResponse) {
	response = &DeleteHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除CLB-WAF防护域名
// 支持批量操作
func (c *Client) DeleteHost(request *DeleteHostRequest) (response *DeleteHostResponse, err error) {
	if request == nil {
		request = NewDeleteHostRequest()
	}
	response = NewDeleteHostResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAttackDownloadRecordRequest() (request *DeleteAttackDownloadRecordRequest) {
	request = &DeleteAttackDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteAttackDownloadRecord")
	return
}

func NewDeleteAttackDownloadRecordResponse() (response *DeleteAttackDownloadRecordResponse) {
	response = &DeleteAttackDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除攻击日志下载任务记录
func (c *Client) DeleteAttackDownloadRecord(request *DeleteAttackDownloadRecordRequest) (response *DeleteAttackDownloadRecordResponse, err error) {
	if request == nil {
		request = NewDeleteAttackDownloadRecordRequest()
	}
	response = NewDeleteAttackDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDownloadRecordRequest() (request *DeleteDownloadRecordRequest) {
	request = &DeleteDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteDownloadRecord")
	return
}

func NewDeleteDownloadRecordResponse() (response *DeleteDownloadRecordResponse) {
	response = &DeleteDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问日志下载记录
func (c *Client) DeleteDownloadRecord(request *DeleteDownloadRecordRequest) (response *DeleteDownloadRecordResponse, err error) {
	if request == nil {
		request = NewDeleteDownloadRecordRequest()
	}
	response = NewDeleteDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTsRegionRequest() (request *ModifyTsRegionRequest) {
	request = &ModifyTsRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyTsRegion")
	return
}

func NewModifyTsRegionResponse() (response *ModifyTsRegionResponse) {
	response = &ModifyTsRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改ts接入区域
func (c *Client) ModifyTsRegion(request *ModifyTsRegionRequest) (response *ModifyTsRegionResponse, err error) {
	if request == nil {
		request = NewModifyTsRegionRequest()
	}
	response = NewModifyTsRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAttackWhiteRuleRequest() (request *DeleteAttackWhiteRuleRequest) {
	request = &DeleteAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteAttackWhiteRule")
	return
}

func NewDeleteAttackWhiteRuleResponse() (response *DeleteAttackWhiteRuleResponse) {
	response = &DeleteAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供用户控制台调用，删除Tiga规则引擎白名单。
func (c *Client) DeleteAttackWhiteRule(request *DeleteAttackWhiteRuleRequest) (response *DeleteAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteAttackWhiteRuleRequest()
	}
	response = NewDeleteAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCachePathRequest() (request *DescribeCachePathRequest) {
	request = &DescribeCachePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCachePath")
	return
}

func NewDescribeCachePathResponse() (response *DescribeCachePathResponse) {
	response = &DescribeCachePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-查找缓存路径
func (c *Client) DescribeCachePath(request *DescribeCachePathRequest) (response *DescribeCachePathResponse, err error) {
	if request == nil {
		request = NewDescribeCachePathRequest()
	}
	response = NewDescribeCachePathResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHistogramRequest() (request *DescribeHistogramRequest) {
	request = &DescribeHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeHistogram")
	return
}

func NewDescribeHistogramResponse() (response *DescribeHistogramResponse) {
	response = &DescribeHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询多种条件的聚类分析
func (c *Client) DescribeHistogram(request *DescribeHistogramRequest) (response *DescribeHistogramResponse, err error) {
	if request == nil {
		request = NewDescribeHistogramRequest()
	}
	response = NewDescribeHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessIndexRequest() (request *DescribeAccessIndexRequest) {
	request = &DescribeAccessIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessIndex")
	return
}

func NewDescribeAccessIndexResponse() (response *DescribeAccessIndexResponse) {
	response = &DescribeAccessIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取访问日志索引配置信息
func (c *Client) DescribeAccessIndex(request *DescribeAccessIndexRequest) (response *DescribeAccessIndexResponse, err error) {
	if request == nil {
		request = NewDescribeAccessIndexRequest()
	}
	response = NewDescribeAccessIndexResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomRuleStatusRequest() (request *ModifyCustomRuleStatusRequest) {
	request = &ModifyCustomRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomRuleStatus")
	return
}

func NewModifyCustomRuleStatusResponse() (response *ModifyCustomRuleStatusResponse) {
	response = &ModifyCustomRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启或禁用访问控制（自定义策略）
func (c *Client) ModifyCustomRuleStatus(request *ModifyCustomRuleStatusRequest) (response *ModifyCustomRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyCustomRuleStatusRequest()
	}
	response = NewModifyCustomRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserSignatureRuleRequest() (request *DescribeUserSignatureRuleRequest) {
	request = &DescribeUserSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserSignatureRule")
	return
}

func NewDescribeUserSignatureRuleResponse() (response *DescribeUserSignatureRuleResponse) {
	response = &DescribeUserSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户特征规则列表
func (c *Client) DescribeUserSignatureRule(request *DescribeUserSignatureRuleRequest) (response *DescribeUserSignatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeUserSignatureRuleRequest()
	}
	response = NewDescribeUserSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessLogsRequest() (request *DescribeAccessLogsRequest) {
	request = &DescribeAccessLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessLogs")
	return
}

func NewDescribeAccessLogsResponse() (response *DescribeAccessLogsResponse) {
	response = &DescribeAccessLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询访问日志
func (c *Client) DescribeAccessLogs(request *DescribeAccessLogsRequest) (response *DescribeAccessLogsResponse, err error) {
	if request == nil {
		request = NewDescribeAccessLogsRequest()
	}
	response = NewDescribeAccessLogsResponse()
	err = c.Send(request, response)
	return
}

func NewApplyBlockPageRequest() (request *ApplyBlockPageRequest) {
	request = &ApplyBlockPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ApplyBlockPage")
	return
}

func NewApplyBlockPageResponse() (response *ApplyBlockPageResponse) {
	response = &ApplyBlockPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给域名应用自定义拦截页面，对于CLB-WAF类型的域名只能应用HTTP状态码为624的自定义拦截页面。
func (c *Client) ApplyBlockPage(request *ApplyBlockPageRequest) (response *ApplyBlockPageResponse, err error) {
	if request == nil {
		request = NewApplyBlockPageRequest()
	}
	response = NewApplyBlockPageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostRequest() (request *DescribeHostRequest) {
	request = &DescribeHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeHost")
	return
}

func NewDescribeHostResponse() (response *DescribeHostResponse) {
	response = &DescribeHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf获取防护域名详情
func (c *Client) DescribeHost(request *DescribeHostRequest) (response *DescribeHostResponse, err error) {
	if request == nil {
		request = NewDescribeHostRequest()
	}
	response = NewDescribeHostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVIPStateRequest() (request *DescribeVIPStateRequest) {
	request = &DescribeVIPStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeVIPState")
	return
}

func NewDescribeVIPStateResponse() (response *DescribeVIPStateResponse) {
	response = &DescribeVIPStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询vip表
func (c *Client) DescribeVIPState(request *DescribeVIPStateRequest) (response *DescribeVIPStateResponse, err error) {
	if request == nil {
		request = NewDescribeVIPStateRequest()
	}
	response = NewDescribeVIPStateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotUCBFeatureRuleRequest() (request *DeleteBotUCBFeatureRuleRequest) {
	request = &DeleteBotUCBFeatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteBotUCBFeatureRule")
	return
}

func NewDeleteBotUCBFeatureRuleResponse() (response *DeleteBotUCBFeatureRuleResponse) {
	response = &DeleteBotUCBFeatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB自定义策略删除
func (c *Client) DeleteBotUCBFeatureRule(request *DeleteBotUCBFeatureRuleRequest) (response *DeleteBotUCBFeatureRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBotUCBFeatureRuleRequest()
	}
	response = NewDeleteBotUCBFeatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiRuleRequest() (request *DescribeApiRuleRequest) {
	request = &DescribeApiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiRule")
	return
}

func NewDescribeApiRuleResponse() (response *DescribeApiRuleResponse) {
	response = &DescribeApiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户查询一条api规则
func (c *Client) DescribeApiRule(request *DescribeApiRuleRequest) (response *DescribeApiRuleResponse, err error) {
	if request == nil {
		request = NewDescribeApiRuleRequest()
	}
	response = NewDescribeApiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessDownloadRecordRequest() (request *CreateAccessDownloadRecordRequest) {
	request = &CreateAccessDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateAccessDownloadRecord")
	return
}

func NewCreateAccessDownloadRecordResponse() (response *CreateAccessDownloadRecordResponse) {
	response = &CreateAccessDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载访问日志，必须购买日志服务
func (c *Client) CreateAccessDownloadRecord(request *CreateAccessDownloadRecordRequest) (response *CreateAccessDownloadRecordResponse, err error) {
	if request == nil {
		request = NewCreateAccessDownloadRecordRequest()
	}
	response = NewCreateAccessDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotItemsRequest() (request *DescribeBotItemsRequest) {
	request = &DescribeBotItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotItems")
	return
}

func NewDescribeBotItemsResponse() (response *DescribeBotItemsResponse) {
	response = &DescribeBotItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志详情对应的字段
func (c *Client) DescribeBotItems(request *DescribeBotItemsRequest) (response *DescribeBotItemsResponse, err error) {
	if request == nil {
		request = NewDescribeBotItemsRequest()
	}
	response = NewDescribeBotItemsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackDetailRequest() (request *DescribeAttackDetailRequest) {
	request = &DescribeAttackDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackDetail")
	return
}

func NewDescribeAttackDetailResponse() (response *DescribeAttackDetailResponse) {
	response = &DescribeAttackDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取攻击日志详情
func (c *Client) DescribeAttackDetail(request *DescribeAttackDetailRequest) (response *DescribeAttackDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAttackDetailRequest()
	}
	response = NewDescribeAttackDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRecordDetailRequest() (request *DescribeBotRecordDetailRequest) {
	request = &DescribeBotRecordDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotRecordDetail")
	return
}

func NewDescribeBotRecordDetailResponse() (response *DescribeBotRecordDetailResponse) {
	response = &DescribeBotRecordDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bot详情
func (c *Client) DescribeBotRecordDetail(request *DescribeBotRecordDetailRequest) (response *DescribeBotRecordDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBotRecordDetailRequest()
	}
	response = NewDescribeBotRecordDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpMainClassRequest() (request *DescribeOpMainClassRequest) {
	request = &DescribeOpMainClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeOpMainClass")
	return
}

func NewDescribeOpMainClassResponse() (response *DescribeOpMainClassResponse) {
	response = &DescribeOpMainClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：获取主类及子类信息
func (c *Client) DescribeOpMainClass(request *DescribeOpMainClassRequest) (response *DescribeOpMainClassResponse, err error) {
	if request == nil {
		request = NewDescribeOpMainClassRequest()
	}
	response = NewDescribeOpMainClassResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCopyAreaBanRequest() (request *CreateCopyAreaBanRequest) {
	request = &CreateCopyAreaBanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateCopyAreaBan")
	return
}

func NewCreateCopyAreaBanResponse() (response *CreateCopyAreaBanResponse) {
	response = &CreateCopyAreaBanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制域名的地域封禁到其他域名，其他域名的地域封禁规则会被清空
func (c *Client) CreateCopyAreaBan(request *CreateCopyAreaBanRequest) (response *CreateCopyAreaBanResponse, err error) {
	if request == nil {
		request = NewCreateCopyAreaBanRequest()
	}
	response = NewCreateCopyAreaBanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSessionKeyRequest() (request *DescribeBotSessionKeyRequest) {
	request = &DescribeBotSessionKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotSessionKey")
	return
}

func NewDescribeBotSessionKeyResponse() (response *DescribeBotSessionKeyResponse) {
	response = &DescribeBotSessionKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot的session配置
func (c *Client) DescribeBotSessionKey(request *DescribeBotSessionKeyRequest) (response *DescribeBotSessionKeyResponse, err error) {
	if request == nil {
		request = NewDescribeBotSessionKeyRequest()
	}
	response = NewDescribeBotSessionKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotThreatIntelligenceRequest() (request *DescribeBotThreatIntelligenceRequest) {
	request = &DescribeBotThreatIntelligenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotThreatIntelligence")
	return
}

func NewDescribeBotThreatIntelligenceResponse() (response *DescribeBotThreatIntelligenceResponse) {
	response = &DescribeBotThreatIntelligenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT的威胁情报配置列表
func (c *Client) DescribeBotThreatIntelligence(request *DescribeBotThreatIntelligenceRequest) (response *DescribeBotThreatIntelligenceResponse, err error) {
	if request == nil {
		request = NewDescribeBotThreatIntelligenceRequest()
	}
	response = NewDescribeBotThreatIntelligenceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeActionedIpRequest() (request *DescribeActionedIpRequest) {
	request = &DescribeActionedIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeActionedIp")
	return
}

func NewDescribeActionedIpResponse() (response *DescribeActionedIpResponse) {
	response = &DescribeActionedIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf Ip状态统一查询接口
func (c *Client) DescribeActionedIp(request *DescribeActionedIpRequest) (response *DescribeActionedIpResponse, err error) {
	if request == nil {
		request = NewDescribeActionedIpRequest()
	}
	response = NewDescribeActionedIpResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiRequestParameterRequest() (request *ModifyApiRequestParameterRequest) {
	request = &ModifyApiRequestParameterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyApiRequestParameter")
	return
}

func NewModifyApiRequestParameterResponse() (response *ModifyApiRequestParameterResponse) {
	response = &ModifyApiRequestParameterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改api自动发现的参数
func (c *Client) ModifyApiRequestParameter(request *ModifyApiRequestParameterRequest) (response *ModifyApiRequestParameterResponse, err error) {
	if request == nil {
		request = NewModifyApiRequestParameterRequest()
	}
	response = NewModifyApiRequestParameterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackDownloadRecordRequest() (request *DescribeAttackDownloadRecordRequest) {
	request = &DescribeAttackDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackDownloadRecord")
	return
}

func NewDescribeAttackDownloadRecordResponse() (response *DescribeAttackDownloadRecordResponse) {
	response = &DescribeAttackDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询下载记录
func (c *Client) DescribeAttackDownloadRecord(request *DescribeAttackDownloadRecordRequest) (response *DescribeAttackDownloadRecordResponse, err error) {
	if request == nil {
		request = NewDescribeAttackDownloadRecordRequest()
	}
	response = NewDescribeAttackDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewModifySceneJsInjectRuleRequest() (request *ModifySceneJsInjectRuleRequest) {
	request = &ModifySceneJsInjectRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifySceneJsInjectRule")
	return
}

func NewModifySceneJsInjectRuleResponse() (response *ModifySceneJsInjectRuleResponse) {
	response = &ModifySceneJsInjectRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【接口复用】前端对抗在场景化版本中，将配置拆分为全局和场景两部分，[SceneId,Status,RuleAction,Redirect]属于场景配置，[AutomatedTool,CrackBehavior,GoodBot,RasSiteID]属于全局配置，字段RuleID变成复用字段，全局时表示tb_waf_jsinject_rule的id，场景时表示tb_waf_jsinject_scene_rule的id
func (c *Client) ModifySceneJsInjectRule(request *ModifySceneJsInjectRuleRequest) (response *ModifySceneJsInjectRuleResponse, err error) {
	if request == nil {
		request = NewModifySceneJsInjectRuleRequest()
	}
	response = NewModifySceneJsInjectRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostFlowModeInnerRequest() (request *ModifyHostFlowModeInnerRequest) {
	request = &ModifyHostFlowModeInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyHostFlowModeInner")
	return
}

func NewModifyHostFlowModeInnerResponse() (response *ModifyHostFlowModeInnerResponse) {
	response = &ModifyHostFlowModeInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf 设置防护域名的流量模式
func (c *Client) ModifyHostFlowModeInner(request *ModifyHostFlowModeInnerRequest) (response *ModifyHostFlowModeInnerResponse, err error) {
	if request == nil {
		request = NewModifyHostFlowModeInnerRequest()
	}
	response = NewModifyHostFlowModeInnerResponse()
	err = c.Send(request, response)
	return
}

func NewCopyBotTCBRuleRequest() (request *CopyBotTCBRuleRequest) {
	request = &CopyBotTCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CopyBotTCBRule")
	return
}

func NewCopyBotTCBRuleResponse() (response *CopyBotTCBRuleResponse) {
	response = &CopyBotTCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 TCB策略域名拷贝
func (c *Client) CopyBotTCBRule(request *CopyBotTCBRuleRequest) (response *CopyBotTCBRuleResponse, err error) {
	if request == nil {
		request = NewCopyBotTCBRuleRequest()
	}
	response = NewCopyBotTCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHostRequest() (request *CreateHostRequest) {
	request = &CreateHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateHost")
	return
}

func NewCreateHostResponse() (response *CreateHostResponse) {
	response = &CreateHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf中添加防护的域名
func (c *Client) CreateHost(request *CreateHostRequest) (response *CreateHostResponse, err error) {
	if request == nil {
		request = NewCreateHostRequest()
	}
	response = NewCreateHostResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCachePathRequest() (request *ModifyCachePathRequest) {
	request = &ModifyCachePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyCachePath")
	return
}

func NewModifyCachePathResponse() (response *ModifyCachePathResponse) {
	response = &ModifyCachePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-编辑缓存路径
func (c *Client) ModifyCachePath(request *ModifyCachePathRequest) (response *ModifyCachePathResponse, err error) {
	if request == nil {
		request = NewModifyCachePathRequest()
	}
	response = NewModifyCachePathResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCdcResourceRequest() (request *ModifyCdcResourceRequest) {
	request = &ModifyCdcResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyCdcResource")
	return
}

func NewModifyCdcResourceResponse() (response *ModifyCdcResourceResponse) {
	response = &ModifyCdcResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改CDC资源列表，如果id=0则添加，>0则修改
func (c *Client) ModifyCdcResource(request *ModifyCdcResourceRequest) (response *ModifyCdcResourceResponse, err error) {
	if request == nil {
		request = NewModifyCdcResourceRequest()
	}
	response = NewModifyCdcResourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotTCBRuleRequest() (request *ModifyBotTCBRuleRequest) {
	request = &ModifyBotTCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotTCBRule")
	return
}

func NewModifyBotTCBRuleResponse() (response *ModifyBotTCBRuleResponse) {
	response = &ModifyBotTCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 TCB策略更新
func (c *Client) ModifyBotTCBRule(request *ModifyBotTCBRuleRequest) (response *ModifyBotTCBRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotTCBRuleRequest()
	}
	response = NewModifyBotTCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserClbWafRegionsRequest() (request *DescribeUserClbWafRegionsRequest) {
	request = &DescribeUserClbWafRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserClbWafRegions")
	return
}

func NewDescribeUserClbWafRegionsResponse() (response *DescribeUserClbWafRegionsResponse) {
	response = &DescribeUserClbWafRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在负载均衡型WAF的添加、编辑域名配置的时候，需要展示负载均衡型WAF（clb-waf)支持的地域列表，通过DescribeUserClbWafRegions既可以获得当前对客户已经开放的地域列表
func (c *Client) DescribeUserClbWafRegions(request *DescribeUserClbWafRegionsRequest) (response *DescribeUserClbWafRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeUserClbWafRegionsRequest()
	}
	response = NewDescribeUserClbWafRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDisableUserSignaturePolicyRequest() (request *DisableUserSignaturePolicyRequest) {
	request = &DisableUserSignaturePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DisableUserSignaturePolicy")
	return
}

func NewDisableUserSignaturePolicyResponse() (response *DisableUserSignaturePolicyResponse) {
	response = &DisableUserSignaturePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭用户防护规则开关
func (c *Client) DisableUserSignaturePolicy(request *DisableUserSignaturePolicyRequest) (response *DisableUserSignaturePolicyResponse, err error) {
	if request == nil {
		request = NewDisableUserSignaturePolicyRequest()
	}
	response = NewDisableUserSignaturePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRequestLogRequest() (request *DescribeBotRequestLogRequest) {
	request = &DescribeBotRequestLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotRequestLog")
	return
}

func NewDescribeBotRequestLogResponse() (response *DescribeBotRequestLogResponse) {
	response = &DescribeBotRequestLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot攻击日志
func (c *Client) DescribeBotRequestLog(request *DescribeBotRequestLogRequest) (response *DescribeBotRequestLogResponse, err error) {
	if request == nil {
		request = NewDescribeBotRequestLogRequest()
	}
	response = NewDescribeBotRequestLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleLimitRequest() (request *DescribeRuleLimitRequest) {
	request = &DescribeRuleLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeRuleLimit")
	return
}

func NewDescribeRuleLimitResponse() (response *DescribeRuleLimitResponse) {
	response = &DescribeRuleLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各个模块具体的规格限制
func (c *Client) DescribeRuleLimit(request *DescribeRuleLimitRequest) (response *DescribeRuleLimitResponse, err error) {
	if request == nil {
		request = NewDescribeRuleLimitRequest()
	}
	response = NewDescribeRuleLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserLevelRequest() (request *DescribeUserLevelRequest) {
	request = &DescribeUserLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserLevel")
	return
}

func NewDescribeUserLevelResponse() (response *DescribeUserLevelResponse) {
	response = &DescribeUserLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户防护规则等级
func (c *Client) DescribeUserLevel(request *DescribeUserLevelRequest) (response *DescribeUserLevelResponse, err error) {
	if request == nil {
		request = NewDescribeUserLevelRequest()
	}
	response = NewDescribeUserLevelResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiSecEventChangeRequest() (request *ModifyApiSecEventChangeRequest) {
	request = &ModifyApiSecEventChangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyApiSecEventChange")
	return
}

func NewModifyApiSecEventChangeResponse() (response *ModifyApiSecEventChangeResponse) {
	response = &ModifyApiSecEventChangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全状态变更接口
func (c *Client) ModifyApiSecEventChange(request *ModifyApiSecEventChangeRequest) (response *ModifyApiSecEventChangeResponse, err error) {
	if request == nil {
		request = NewModifyApiSecEventChangeRequest()
	}
	response = NewModifyApiSecEventChangeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBatchIpAccessControlRequest() (request *DescribeBatchIpAccessControlRequest) {
	request = &DescribeBatchIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBatchIpAccessControl")
	return
}

func NewDescribeBatchIpAccessControlResponse() (response *DescribeBatchIpAccessControlResponse) {
	response = &DescribeBatchIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 多域名ip黑白名单查询
func (c *Client) DescribeBatchIpAccessControl(request *DescribeBatchIpAccessControlRequest) (response *DescribeBatchIpAccessControlResponse, err error) {
	if request == nil {
		request = NewDescribeBatchIpAccessControlRequest()
	}
	response = NewDescribeBatchIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceIdListRequest() (request *DescribeInstanceIdListRequest) {
	request = &DescribeInstanceIdListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeInstanceIdList")
	return
}

func NewDescribeInstanceIdListResponse() (response *DescribeInstanceIdListResponse) {
	response = &DescribeInstanceIdListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按地域，返回实例id的列表；key为地域，值为实例列表
func (c *Client) DescribeInstanceIdList(request *DescribeInstanceIdListRequest) (response *DescribeInstanceIdListResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceIdListRequest()
	}
	response = NewDescribeInstanceIdListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductInfoRequest() (request *DescribeProductInfoRequest) {
	request = &DescribeProductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeProductInfo")
	return
}

func NewDescribeProductInfoResponse() (response *DescribeProductInfoResponse) {
	response = &DescribeProductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客户产品信息
func (c *Client) DescribeProductInfo(request *DescribeProductInfoRequest) (response *DescribeProductInfoResponse, err error) {
	if request == nil {
		request = NewDescribeProductInfoRequest()
	}
	response = NewDescribeProductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostStatusRequest() (request *ModifyHostStatusRequest) {
	request = &ModifyHostStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyHostStatus")
	return
}

func NewModifyHostStatusResponse() (response *ModifyHostStatusResponse) {
	response = &ModifyHostStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf 设置防护域名WAF开关
// 支持批量操作。
func (c *Client) ModifyHostStatus(request *ModifyHostStatusRequest) (response *ModifyHostStatusResponse, err error) {
	if request == nil {
		request = NewModifyHostStatusRequest()
	}
	response = NewModifyHostStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAttackDownloadTaskRequest() (request *CreateAttackDownloadTaskRequest) {
	request = &CreateAttackDownloadTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateAttackDownloadTask")
	return
}

func NewCreateAttackDownloadTaskResponse() (response *CreateAttackDownloadTaskResponse) {
	response = &CreateAttackDownloadTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建攻击日志下载任务
func (c *Client) CreateAttackDownloadTask(request *CreateAttackDownloadTaskRequest) (response *CreateAttackDownloadTaskResponse, err error) {
	if request == nil {
		request = NewCreateAttackDownloadTaskRequest()
	}
	response = NewCreateAttackDownloadTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccessDownloadRecordRequest() (request *DeleteAccessDownloadRecordRequest) {
	request = &DeleteAccessDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteAccessDownloadRecord")
	return
}

func NewDeleteAccessDownloadRecordResponse() (response *DeleteAccessDownloadRecordResponse) {
	response = &DeleteAccessDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问日志下载任务记录
func (c *Client) DeleteAccessDownloadRecord(request *DeleteAccessDownloadRecordRequest) (response *DeleteAccessDownloadRecordResponse, err error) {
	if request == nil {
		request = NewDeleteAccessDownloadRecordRequest()
	}
	response = NewDeleteAccessDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCdcResourceRequest() (request *DescribeCdcResourceRequest) {
	request = &DescribeCdcResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCdcResource")
	return
}

func NewDescribeCdcResourceResponse() (response *DescribeCdcResourceResponse) {
	response = &DescribeCdcResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供运营后台获取Cdc资源列表
func (c *Client) DescribeCdcResource(request *DescribeCdcResourceRequest) (response *DescribeCdcResourceResponse, err error) {
	if request == nil {
		request = NewDescribeCdcResourceRequest()
	}
	response = NewDescribeCdcResourceResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackCountRequest() (request *GetAttackCountRequest) {
	request = &GetAttackCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetAttackCount")
	return
}

func NewGetAttackCountResponse() (response *GetAttackCountResponse) {
	response = &GetAttackCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意攻击总量
func (c *Client) GetAttackCount(request *GetAttackCountRequest) (response *GetAttackCountResponse, err error) {
	if request == nil {
		request = NewGetAttackCountRequest()
	}
	response = NewGetAttackCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBusinessRiskRuleRequest() (request *ModifyBusinessRiskRuleRequest) {
	request = &ModifyBusinessRiskRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBusinessRiskRule")
	return
}

func NewModifyBusinessRiskRuleResponse() (response *ModifyBusinessRiskRuleResponse) {
	response = &ModifyBusinessRiskRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改业务安全规则
func (c *Client) ModifyBusinessRiskRule(request *ModifyBusinessRiskRuleRequest) (response *ModifyBusinessRiskRuleResponse, err error) {
	if request == nil {
		request = NewModifyBusinessRiskRuleRequest()
	}
	response = NewModifyBusinessRiskRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackWhiteRuleRequest() (request *DescribeAttackWhiteRuleRequest) {
	request = &DescribeAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackWhiteRule")
	return
}

func NewDescribeAttackWhiteRuleResponse() (response *DescribeAttackWhiteRuleResponse) {
	response = &DescribeAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户规则白名单列表
func (c *Client) DescribeAttackWhiteRule(request *DescribeAttackWhiteRuleRequest) (response *DescribeAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeAttackWhiteRuleRequest()
	}
	response = NewDescribeAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomWhiteRuleRequest() (request *DescribeCustomWhiteRuleRequest) {
	request = &DescribeCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomWhiteRule")
	return
}

func NewDescribeCustomWhiteRuleResponse() (response *DescribeCustomWhiteRuleResponse) {
	response = &DescribeCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防护配置中的精准白名单策略列表
func (c *Client) DescribeCustomWhiteRule(request *DescribeCustomWhiteRuleRequest) (response *DescribeCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeCustomWhiteRuleRequest()
	}
	response = NewDescribeCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAttackWhiteRuleRequest() (request *ModifyAttackWhiteRuleRequest) {
	request = &ModifyAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAttackWhiteRule")
	return
}

func NewModifyAttackWhiteRuleResponse() (response *ModifyAttackWhiteRuleResponse) {
	response = &ModifyAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供用户控制台调用，修改Tiga规则引擎白名单。
func (c *Client) ModifyAttackWhiteRule(request *ModifyAttackWhiteRuleRequest) (response *ModifyAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyAttackWhiteRuleRequest()
	}
	response = NewModifyAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddDomainWhiteRuleRequest() (request *AddDomainWhiteRuleRequest) {
	request = &AddDomainWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddDomainWhiteRule")
	return
}

func NewAddDomainWhiteRuleResponse() (response *AddDomainWhiteRuleResponse) {
	response = &AddDomainWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加域名规则白名单
func (c *Client) AddDomainWhiteRule(request *AddDomainWhiteRuleRequest) (response *AddDomainWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddDomainWhiteRuleRequest()
	}
	response = NewAddDomainWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotSessionKeyRequest() (request *DeleteBotSessionKeyRequest) {
	request = &DeleteBotSessionKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteBotSessionKey")
	return
}

func NewDeleteBotSessionKeyResponse() (response *DeleteBotSessionKeyResponse) {
	response = &DeleteBotSessionKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除bot session配置
func (c *Client) DeleteBotSessionKey(request *DeleteBotSessionKeyRequest) (response *DeleteBotSessionKeyResponse, err error) {
	if request == nil {
		request = NewDeleteBotSessionKeyRequest()
	}
	response = NewDeleteBotSessionKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackHistogramRequest() (request *GetAttackHistogramRequest) {
	request = &GetAttackHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetAttackHistogram")
	return
}

func NewGetAttackHistogramResponse() (response *GetAttackHistogramResponse) {
	response = &GetAttackHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成攻击日志的产生时间柱状图
func (c *Client) GetAttackHistogram(request *GetAttackHistogramRequest) (response *GetAttackHistogramResponse, err error) {
	if request == nil {
		request = NewGetAttackHistogramRequest()
	}
	response = NewGetAttackHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSessionRequest() (request *DescribeSessionRequest) {
	request = &DescribeSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeSession")
	return
}

func NewDescribeSessionResponse() (response *DescribeSessionResponse) {
	response = &DescribeSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 会话定义查询接口
func (c *Client) DescribeSession(request *DescribeSessionRequest) (response *DescribeSessionResponse, err error) {
	if request == nil {
		request = NewDescribeSessionRequest()
	}
	response = NewDescribeSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCachePathRequest() (request *DeleteCachePathRequest) {
	request = &DeleteCachePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteCachePath")
	return
}

func NewDeleteCachePathResponse() (response *DeleteCachePathResponse) {
	response = &DeleteCachePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-删除缓存路径
func (c *Client) DeleteCachePath(request *DeleteCachePathRequest) (response *DeleteCachePathResponse, err error) {
	if request == nil {
		request = NewDeleteCachePathRequest()
	}
	response = NewDeleteCachePathResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostStatusInnerRequest() (request *ModifyHostStatusInnerRequest) {
	request = &ModifyHostStatusInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyHostStatusInner")
	return
}

func NewModifyHostStatusInnerResponse() (response *ModifyHostStatusInnerResponse) {
	response = &ModifyHostStatusInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf 设置防护域名WAF开关
// 支持批量操作。
func (c *Client) ModifyHostStatusInner(request *ModifyHostStatusInnerRequest) (response *ModifyHostStatusInnerResponse, err error) {
	if request == nil {
		request = NewModifyHostStatusInnerRequest()
	}
	response = NewModifyHostStatusInnerResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertSessionRequest() (request *UpsertSessionRequest) {
	request = &UpsertSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "UpsertSession")
	return
}

func NewUpsertSessionResponse() (response *UpsertSessionResponse) {
	response = &UpsertSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  会话定义 Upsert接口
func (c *Client) UpsertSession(request *UpsertSessionRequest) (response *UpsertSessionResponse, err error) {
	if request == nil {
		request = NewUpsertSessionRequest()
	}
	response = NewUpsertSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiRuleListRequest() (request *DescribeApiRuleListRequest) {
	request = &DescribeApiRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiRuleList")
	return
}

func NewDescribeApiRuleListResponse() (response *DescribeApiRuleListResponse) {
	response = &DescribeApiRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据多条件查询API规则，包括规则ID（精准搜索）/API参数（模糊搜索）/接口名称（模糊搜索）/执行动作（精准搜索）/请求方法（精准搜索）/来源（精准搜索）/开关状态（精准搜索）
func (c *Client) DescribeApiRuleList(request *DescribeApiRuleListRequest) (response *DescribeApiRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeApiRuleListRequest()
	}
	response = NewDescribeApiRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExpiredInstancesRequest() (request *DescribeExpiredInstancesRequest) {
	request = &DescribeExpiredInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeExpiredInstances")
	return
}

func NewDescribeExpiredInstancesResponse() (response *DescribeExpiredInstancesResponse) {
	response = &DescribeExpiredInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户所有销毁实例的详细信息
func (c *Client) DescribeExpiredInstances(request *DescribeExpiredInstancesRequest) (response *DescribeExpiredInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeExpiredInstancesRequest()
	}
	response = NewDescribeExpiredInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePocTestTaskRequest() (request *DeletePocTestTaskRequest) {
	request = &DeletePocTestTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeletePocTestTask")
	return
}

func NewDeletePocTestTaskResponse() (response *DeletePocTestTaskResponse) {
	response = &DeletePocTestTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除POC测试任务
func (c *Client) DeletePocTestTask(request *DeletePocTestTaskRequest) (response *DeletePocTestTaskResponse, err error) {
	if request == nil {
		request = NewDeletePocTestTaskRequest()
	}
	response = NewDeletePocTestTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAIModelRequest() (request *ModifyAIModelRequest) {
	request = &ModifyAIModelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAIModel")
	return
}

func NewModifyAIModelResponse() (response *ModifyAIModelResponse) {
	response = &ModifyAIModelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启ai模型训练
func (c *Client) ModifyAIModel(request *ModifyAIModelRequest) (response *ModifyAIModelResponse, err error) {
	if request == nil {
		request = NewModifyAIModelRequest()
	}
	response = NewModifyAIModelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpartaProtectionListRequest() (request *DescribeSpartaProtectionListRequest) {
	request = &DescribeSpartaProtectionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeSpartaProtectionList")
	return
}

func NewDescribeSpartaProtectionListResponse() (response *DescribeSpartaProtectionListResponse) {
	response = &DescribeSpartaProtectionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询域名列表
func (c *Client) DescribeSpartaProtectionList(request *DescribeSpartaProtectionListRequest) (response *DescribeSpartaProtectionListResponse, err error) {
	if request == nil {
		request = NewDescribeSpartaProtectionListRequest()
	}
	response = NewDescribeSpartaProtectionListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotAiRuleRequest() (request *ModifyBotAiRuleRequest) {
	request = &ModifyBotAiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotAiRule")
	return
}

func NewModifyBotAiRuleResponse() (response *ModifyBotAiRuleResponse) {
	response = &ModifyBotAiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot-ai模块配置
func (c *Client) ModifyBotAiRule(request *ModifyBotAiRuleRequest) (response *ModifyBotAiRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotAiRuleRequest()
	}
	response = NewModifyBotAiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecAttackSourceListRequest() (request *DescribeApiSecAttackSourceListRequest) {
	request = &DescribeApiSecAttackSourceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiSecAttackSourceList")
	return
}

func NewDescribeApiSecAttackSourceListResponse() (response *DescribeApiSecAttackSourceListResponse) {
	response = &DescribeApiSecAttackSourceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取api安全事件攻击源详情
func (c *Client) DescribeApiSecAttackSourceList(request *DescribeApiSecAttackSourceListRequest) (response *DescribeApiSecAttackSourceListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecAttackSourceListRequest()
	}
	response = NewDescribeApiSecAttackSourceListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDNSDetectDomainRequest() (request *CreateDNSDetectDomainRequest) {
	request = &CreateDNSDetectDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateDNSDetectDomain")
	return
}

func NewCreateDNSDetectDomainResponse() (response *CreateDNSDetectDomainResponse) {
	response = &CreateDNSDetectDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加DNS劫持检测的域名，对应【Web安全检测】【DNS劫持检测】【基础设置】【添加域名】这个功能。
func (c *Client) CreateDNSDetectDomain(request *CreateDNSDetectDomainRequest) (response *CreateDNSDetectDomainResponse, err error) {
	if request == nil {
		request = NewCreateDNSDetectDomainRequest()
	}
	response = NewCreateDNSDetectDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotSceneRequest() (request *DeleteBotSceneRequest) {
	request = &DeleteBotSceneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteBotScene")
	return
}

func NewDeleteBotSceneResponse() (response *DeleteBotSceneResponse) {
	response = &DeleteBotSceneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除BOT场景配置
func (c *Client) DeleteBotScene(request *DeleteBotSceneRequest) (response *DeleteBotSceneResponse, err error) {
	if request == nil {
		request = NewDeleteBotSceneRequest()
	}
	response = NewDeleteBotSceneResponse()
	err = c.Send(request, response)
	return
}

func NewPostAttackDownloadTaskRequest() (request *PostAttackDownloadTaskRequest) {
	request = &PostAttackDownloadTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "PostAttackDownloadTask")
	return
}

func NewPostAttackDownloadTaskResponse() (response *PostAttackDownloadTaskResponse) {
	response = &PostAttackDownloadTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建搜索下载攻击日志任务，使用CLS新版本的搜索下载getlog接口
func (c *Client) PostAttackDownloadTask(request *PostAttackDownloadTaskRequest) (response *PostAttackDownloadTaskResponse, err error) {
	if request == nil {
		request = NewPostAttackDownloadTaskRequest()
	}
	response = NewPostAttackDownloadTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainPostActionRequest() (request *ModifyDomainPostActionRequest) {
	request = &ModifyDomainPostActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainPostAction")
	return
}

func NewModifyDomainPostActionResponse() (response *ModifyDomainPostActionResponse) {
	response = &ModifyDomainPostActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名投递状态
func (c *Client) ModifyDomainPostAction(request *ModifyDomainPostActionRequest) (response *ModifyDomainPostActionResponse, err error) {
	if request == nil {
		request = NewModifyDomainPostActionRequest()
	}
	response = NewModifyDomainPostActionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeParametersListRequest() (request *DescribeParametersListRequest) {
	request = &DescribeParametersListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeParametersList")
	return
}

func NewDescribeParametersListResponse() (response *DescribeParametersListResponse) {
	response = &DescribeParametersListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持的参数类型列表
func (c *Client) DescribeParametersList(request *DescribeParametersListRequest) (response *DescribeParametersListResponse, err error) {
	if request == nil {
		request = NewDescribeParametersListRequest()
	}
	response = NewDescribeParametersListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCLSRequest() (request *DescribeCLSRequest) {
	request = &DescribeCLSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCLS")
	return
}

func NewDescribeCLSResponse() (response *DescribeCLSResponse) {
	response = &DescribeCLSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前购买的日志服务的使用状况，用户必须已经购买了日志服务才能调用此接口
func (c *Client) DescribeCLS(request *DescribeCLSRequest) (response *DescribeCLSResponse, err error) {
	if request == nil {
		request = NewDescribeCLSRequest()
	}
	response = NewDescribeCLSResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchDomainWhiteRulesRequest() (request *SwitchDomainWhiteRulesRequest) {
	request = &SwitchDomainWhiteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "SwitchDomainWhiteRules")
	return
}

func NewSwitchDomainWhiteRulesResponse() (response *SwitchDomainWhiteRulesResponse) {
	response = &SwitchDomainWhiteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换域名白名单的开关
func (c *Client) SwitchDomainWhiteRules(request *SwitchDomainWhiteRulesRequest) (response *SwitchDomainWhiteRulesResponse, err error) {
	if request == nil {
		request = NewSwitchDomainWhiteRulesRequest()
	}
	response = NewSwitchDomainWhiteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCCAutoStatusRequest() (request *DescribeCCAutoStatusRequest) {
	request = &DescribeCCAutoStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCCAutoStatus")
	return
}

func NewDescribeCCAutoStatusResponse() (response *DescribeCCAutoStatusResponse) {
	response = &DescribeCCAutoStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 斯巴达版本查询cc自动封堵状态
func (c *Client) DescribeCCAutoStatus(request *DescribeCCAutoStatusRequest) (response *DescribeCCAutoStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCCAutoStatusRequest()
	}
	response = NewDescribeCCAutoStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostFlowModeRequest() (request *ModifyHostFlowModeRequest) {
	request = &ModifyHostFlowModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyHostFlowMode")
	return
}

func NewModifyHostFlowModeResponse() (response *ModifyHostFlowModeResponse) {
	response = &ModifyHostFlowModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf 设置防护域名的流量模式
func (c *Client) ModifyHostFlowMode(request *ModifyHostFlowModeRequest) (response *ModifyHostFlowModeResponse, err error) {
	if request == nil {
		request = NewModifyHostFlowModeRequest()
	}
	response = NewModifyHostFlowModeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryClientMsgRequest() (request *QueryClientMsgRequest) {
	request = &QueryClientMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "QueryClientMsg")
	return
}

func NewQueryClientMsgResponse() (response *QueryClientMsgResponse) {
	response = &QueryClientMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询是否开启透传客户端信息
func (c *Client) QueryClientMsg(request *QueryClientMsgRequest) (response *QueryClientMsgResponse, err error) {
	if request == nil {
		request = NewQueryClientMsgRequest()
	}
	response = NewQueryClientMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotGlobalConfigRequest() (request *DescribeBotGlobalConfigRequest) {
	request = &DescribeBotGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotGlobalConfig")
	return
}

func NewDescribeBotGlobalConfigResponse() (response *DescribeBotGlobalConfigResponse) {
	response = &DescribeBotGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT全局设置信息
func (c *Client) DescribeBotGlobalConfig(request *DescribeBotGlobalConfigRequest) (response *DescribeBotGlobalConfigResponse, err error) {
	if request == nil {
		request = NewDescribeBotGlobalConfigRequest()
	}
	response = NewDescribeBotGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTypeStatRequest() (request *DescribeBotTypeStatRequest) {
	request = &DescribeBotTypeStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotTypeStat")
	return
}

func NewDescribeBotTypeStatResponse() (response *DescribeBotTypeStatResponse) {
	response = &DescribeBotTypeStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 Bot类别统计
func (c *Client) DescribeBotTypeStat(request *DescribeBotTypeStatRequest) (response *DescribeBotTypeStatResponse, err error) {
	if request == nil {
		request = NewDescribeBotTypeStatRequest()
	}
	response = NewDescribeBotTypeStatResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertCCRuleRequest() (request *UpsertCCRuleRequest) {
	request = &UpsertCCRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "UpsertCCRule")
	return
}

func NewUpsertCCRuleResponse() (response *UpsertCCRuleResponse) {
	response = &UpsertCCRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  CC V2 Upsert接口
func (c *Client) UpsertCCRule(request *UpsertCCRuleRequest) (response *UpsertCCRuleResponse, err error) {
	if request == nil {
		request = NewUpsertCCRuleRequest()
	}
	response = NewUpsertCCRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertBotUCBFeatureRuleRequest() (request *UpsertBotUCBFeatureRuleRequest) {
	request = &UpsertBotUCBFeatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "UpsertBotUCBFeatureRule")
	return
}

func NewUpsertBotUCBFeatureRuleResponse() (response *UpsertBotUCBFeatureRuleResponse) {
	response = &UpsertBotUCBFeatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB策略更新
func (c *Client) UpsertBotUCBFeatureRule(request *UpsertBotUCBFeatureRuleRequest) (response *UpsertBotUCBFeatureRuleResponse, err error) {
	if request == nil {
		request = NewUpsertBotUCBFeatureRuleRequest()
	}
	response = NewUpsertBotUCBFeatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotDetailRequest() (request *DescribeBotDetailRequest) {
	request = &DescribeBotDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotDetail")
	return
}

func NewDescribeBotDetailResponse() (response *DescribeBotDetailResponse) {
	response = &DescribeBotDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot的具体详情
func (c *Client) DescribeBotDetail(request *DescribeBotDetailRequest) (response *DescribeBotDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBotDetailRequest()
	}
	response = NewDescribeBotDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAreaBanAreasRequest() (request *DescribeAreaBanAreasRequest) {
	request = &DescribeAreaBanAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAreaBanAreas")
	return
}

func NewDescribeAreaBanAreasResponse() (response *DescribeAreaBanAreasResponse) {
	response = &DescribeAreaBanAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取地域封禁配置包括地域封禁开关，设置封禁的地区信息
func (c *Client) DescribeAreaBanAreas(request *DescribeAreaBanAreasRequest) (response *DescribeAreaBanAreasResponse, err error) {
	if request == nil {
		request = NewDescribeAreaBanAreasRequest()
	}
	response = NewDescribeAreaBanAreasResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessPeriodRequest() (request *ModifyAccessPeriodRequest) {
	request = &ModifyAccessPeriodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAccessPeriod")
	return
}

func NewModifyAccessPeriodResponse() (response *ModifyAccessPeriodResponse) {
	response = &ModifyAccessPeriodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改访问日志保存期限及大字段是否存储
func (c *Client) ModifyAccessPeriod(request *ModifyAccessPeriodRequest) (response *ModifyAccessPeriodResponse, err error) {
	if request == nil {
		request = NewModifyAccessPeriodRequest()
	}
	response = NewModifyAccessPeriodResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSceneActionRuleRequest() (request *ModifyBotSceneActionRuleRequest) {
	request = &ModifyBotSceneActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotSceneActionRule")
	return
}

func NewModifyBotSceneActionRuleResponse() (response *ModifyBotSceneActionRuleResponse) {
	response = &ModifyBotSceneActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增or修改BOT场景的动作策略配置
func (c *Client) ModifyBotSceneActionRule(request *ModifyBotSceneActionRuleRequest) (response *ModifyBotSceneActionRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotSceneActionRuleRequest()
	}
	response = NewModifyBotSceneActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
	request = &DescribeDomainsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomains")
	return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
	response = &DescribeDomainsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户所有域名的详细信息
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
	if request == nil {
		request = NewDescribeDomainsRequest()
	}
	response = NewDescribeDomainsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGlobalWhiteRuleRequest() (request *DescribeGlobalWhiteRuleRequest) {
	request = &DescribeGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeGlobalWhiteRule")
	return
}

func NewDescribeGlobalWhiteRuleResponse() (response *DescribeGlobalWhiteRuleResponse) {
	response = &DescribeGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供内部运营系统调用，获取全局白名单列表。
func (c *Client) DescribeGlobalWhiteRule(request *DescribeGlobalWhiteRuleRequest) (response *DescribeGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeGlobalWhiteRuleRequest()
	}
	response = NewDescribeGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpHitItemsRequest() (request *DescribeIpHitItemsRequest) {
	request = &DescribeIpHitItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeIpHitItems")
	return
}

func NewDescribeIpHitItemsResponse() (response *DescribeIpHitItemsResponse) {
	response = &DescribeIpHitItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  IP封堵状态查询
func (c *Client) DescribeIpHitItems(request *DescribeIpHitItemsRequest) (response *DescribeIpHitItemsResponse, err error) {
	if request == nil {
		request = NewDescribeIpHitItemsRequest()
	}
	response = NewDescribeIpHitItemsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotSceneActionRuleRequest() (request *DeleteBotSceneActionRuleRequest) {
	request = &DeleteBotSceneActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteBotSceneActionRule")
	return
}

func NewDeleteBotSceneActionRuleResponse() (response *DeleteBotSceneActionRuleResponse) {
	response = &DeleteBotSceneActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除BOT场景的动作策略
func (c *Client) DeleteBotSceneActionRule(request *DeleteBotSceneActionRuleRequest) (response *DeleteBotSceneActionRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBotSceneActionRuleRequest()
	}
	response = NewDeleteBotSceneActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSDetectHijackDataRequest() (request *DescribeDNSDetectHijackDataRequest) {
	request = &DescribeDNSDetectHijackDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDNSDetectHijackData")
	return
}

func NewDescribeDNSDetectHijackDataResponse() (response *DescribeDNSDetectHijackDataResponse) {
	response = &DescribeDNSDetectHijackDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DNS劫持检测的劫持记录，对应【Web安全检测】【DNS劫持检测】【劫持记录】这个功能。
func (c *Client) DescribeDNSDetectHijackData(request *DescribeDNSDetectHijackDataRequest) (response *DescribeDNSDetectHijackDataResponse, err error) {
	if request == nil {
		request = NewDescribeDNSDetectHijackDataRequest()
	}
	response = NewDescribeDNSDetectHijackDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskInfoRequest() (request *DescribeRiskInfoRequest) {
	request = &DescribeRiskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeRiskInfo")
	return
}

func NewDescribeRiskInfoResponse() (response *DescribeRiskInfoResponse) {
	response = &DescribeRiskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务安全风险信息查询
func (c *Client) DescribeRiskInfo(request *DescribeRiskInfoRequest) (response *DescribeRiskInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRiskInfoRequest()
	}
	response = NewDescribeRiskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAntiInfoLeakRuleStatusRequest() (request *ModifyAntiInfoLeakRuleStatusRequest) {
	request = &ModifyAntiInfoLeakRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiInfoLeakRuleStatus")
	return
}

func NewModifyAntiInfoLeakRuleStatusResponse() (response *ModifyAntiInfoLeakRuleStatusResponse) {
	response = &ModifyAntiInfoLeakRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 信息防泄漏切换规则开关
func (c *Client) ModifyAntiInfoLeakRuleStatus(request *ModifyAntiInfoLeakRuleStatusRequest) (response *ModifyAntiInfoLeakRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyAntiInfoLeakRuleStatusRequest()
	}
	response = NewModifyAntiInfoLeakRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSceneRequest() (request *ModifyBotSceneRequest) {
	request = &ModifyBotSceneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotScene")
	return
}

func NewModifyBotSceneResponse() (response *ModifyBotSceneResponse) {
	response = &ModifyBotSceneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增or修改BOT场景配置
func (c *Client) ModifyBotScene(request *ModifyBotSceneRequest) (response *ModifyBotSceneResponse, err error) {
	if request == nil {
		request = NewModifyBotSceneRequest()
	}
	response = NewModifyBotSceneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotWhiteStatusRequest() (request *DescribeBotWhiteStatusRequest) {
	request = &DescribeBotWhiteStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotWhiteStatus")
	return
}

func NewDescribeBotWhiteStatusResponse() (response *DescribeBotWhiteStatusResponse) {
	response = &DescribeBotWhiteStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取goodbot的开关列表
func (c *Client) DescribeBotWhiteStatus(request *DescribeBotWhiteStatusRequest) (response *DescribeBotWhiteStatusResponse, err error) {
	if request == nil {
		request = NewDescribeBotWhiteStatusRequest()
	}
	response = NewDescribeBotWhiteStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCopyBotUCBPreinstallRuleRequest() (request *CopyBotUCBPreinstallRuleRequest) {
	request = &CopyBotUCBPreinstallRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CopyBotUCBPreinstallRule")
	return
}

func NewCopyBotUCBPreinstallRuleResponse() (response *CopyBotUCBPreinstallRuleResponse) {
	response = &CopyBotUCBPreinstallRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB预设策略拷贝
func (c *Client) CopyBotUCBPreinstallRule(request *CopyBotUCBPreinstallRuleRequest) (response *CopyBotUCBPreinstallRuleResponse, err error) {
	if request == nil {
		request = NewCopyBotUCBPreinstallRuleRequest()
	}
	response = NewCopyBotUCBPreinstallRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVipInfoRequest() (request *DescribeVipInfoRequest) {
	request = &DescribeVipInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeVipInfo")
	return
}

func NewDescribeVipInfoResponse() (response *DescribeVipInfoResponse) {
	response = &DescribeVipInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据过滤条件查询VIP信息
func (c *Client) DescribeVipInfo(request *DescribeVipInfoRequest) (response *DescribeVipInfoResponse, err error) {
	if request == nil {
		request = NewDescribeVipInfoRequest()
	}
	response = NewDescribeVipInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainRulesEnRequest() (request *DescribeDomainRulesEnRequest) {
	request = &DescribeDomainRulesEnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainRulesEn")
	return
}

func NewDescribeDomainRulesEnResponse() (response *DescribeDomainRulesEnResponse) {
	response = &DescribeDomainRulesEnResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取域名的防护规则列表支持英文
func (c *Client) DescribeDomainRulesEn(request *DescribeDomainRulesEnRequest) (response *DescribeDomainRulesEnResponse, err error) {
	if request == nil {
		request = NewDescribeDomainRulesEnRequest()
	}
	response = NewDescribeDomainRulesEnResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGlobalWhiteRuleRequest() (request *DeleteGlobalWhiteRuleRequest) {
	request = &DeleteGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteGlobalWhiteRule")
	return
}

func NewDeleteGlobalWhiteRuleResponse() (response *DeleteGlobalWhiteRuleResponse) {
	response = &DeleteGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供内部运营系统调用，删除Tiga规则引擎全局白名单。
func (c *Client) DeleteGlobalWhiteRule(request *DeleteGlobalWhiteRuleRequest) (response *DeleteGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteGlobalWhiteRuleRequest()
	}
	response = NewDeleteGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSessionRequest() (request *DeleteSessionRequest) {
	request = &DeleteSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteSession")
	return
}

func NewDeleteSessionResponse() (response *DeleteSessionResponse) {
	response = &DeleteSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除CC攻击的session设置
func (c *Client) DeleteSession(request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
	if request == nil {
		request = NewDeleteSessionRequest()
	}
	response = NewDeleteSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAutoStatisticRuleRequest() (request *DescribeBotAutoStatisticRuleRequest) {
	request = &DescribeBotAutoStatisticRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotAutoStatisticRule")
	return
}

func NewDescribeBotAutoStatisticRuleResponse() (response *DescribeBotAutoStatisticRuleResponse) {
	response = &DescribeBotAutoStatisticRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot-智能统计模块配置
func (c *Client) DescribeBotAutoStatisticRule(request *DescribeBotAutoStatisticRuleRequest) (response *DescribeBotAutoStatisticRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotAutoStatisticRuleRequest()
	}
	response = NewDescribeBotAutoStatisticRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJsInjectRuleRequest() (request *DescribeJsInjectRuleRequest) {
	request = &DescribeJsInjectRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeJsInjectRule")
	return
}

func NewDescribeJsInjectRuleResponse() (response *DescribeJsInjectRuleResponse) {
	response = &DescribeJsInjectRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对应域名的前端对抗规则。
func (c *Client) DescribeJsInjectRule(request *DescribeJsInjectRuleRequest) (response *DescribeJsInjectRuleResponse, err error) {
	if request == nil {
		request = NewDescribeJsInjectRuleRequest()
	}
	response = NewDescribeJsInjectRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTsRegionRequest() (request *DescribeTsRegionRequest) {
	request = &DescribeTsRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeTsRegion")
	return
}

func NewDescribeTsRegionResponse() (response *DescribeTsRegionResponse) {
	response = &DescribeTsRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ts接入区域
func (c *Client) DescribeTsRegion(request *DescribeTsRegionRequest) (response *DescribeTsRegionResponse, err error) {
	if request == nil {
		request = NewDescribeTsRegionRequest()
	}
	response = NewDescribeTsRegionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotUCBPreinstallRuleRequest() (request *ModifyBotUCBPreinstallRuleRequest) {
	request = &ModifyBotUCBPreinstallRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotUCBPreinstallRule")
	return
}

func NewModifyBotUCBPreinstallRuleResponse() (response *ModifyBotUCBPreinstallRuleResponse) {
	response = &ModifyBotUCBPreinstallRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB预设规则更新
func (c *Client) ModifyBotUCBPreinstallRule(request *ModifyBotUCBPreinstallRuleRequest) (response *ModifyBotUCBPreinstallRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotUCBPreinstallRuleRequest()
	}
	response = NewModifyBotUCBPreinstallRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomWhiteRuleRequest() (request *DeleteCustomWhiteRuleRequest) {
	request = &DeleteCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteCustomWhiteRule")
	return
}

func NewDeleteCustomWhiteRuleResponse() (response *DeleteCustomWhiteRuleResponse) {
	response = &DeleteCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除精准白名单规则
func (c *Client) DeleteCustomWhiteRule(request *DeleteCustomWhiteRuleRequest) (response *DeleteCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCustomWhiteRuleRequest()
	}
	response = NewDeleteCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchApiRulesRequest() (request *SwitchApiRulesRequest) {
	request = &SwitchApiRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "SwitchApiRules")
	return
}

func NewSwitchApiRulesResponse() (response *SwitchApiRulesResponse) {
	response = &SwitchApiRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用或者停用Api规则
func (c *Client) SwitchApiRules(request *SwitchApiRulesRequest) (response *SwitchApiRulesResponse, err error) {
	if request == nil {
		request = NewSwitchApiRulesRequest()
	}
	response = NewSwitchApiRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostFlowServiceRequest() (request *DescribePostFlowServiceRequest) {
	request = &DescribePostFlowServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePostFlowService")
	return
}

func NewDescribePostFlowServiceResponse() (response *DescribePostFlowServiceResponse) {
	response = &DescribePostFlowServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查开通投递服务接口
func (c *Client) DescribePostFlowService(request *DescribePostFlowServiceRequest) (response *DescribePostFlowServiceResponse, err error) {
	if request == nil {
		request = NewDescribePostFlowServiceRequest()
	}
	response = NewDescribePostFlowServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserDomainInfoRequest() (request *DescribeUserDomainInfoRequest) {
	request = &DescribeUserDomainInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserDomainInfo")
	return
}

func NewDescribeUserDomainInfoResponse() (response *DescribeUserDomainInfoResponse) {
	response = &DescribeUserDomainInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询saas和clb的域名信息
func (c *Client) DescribeUserDomainInfo(request *DescribeUserDomainInfoRequest) (response *DescribeUserDomainInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserDomainInfoRequest()
	}
	response = NewDescribeUserDomainInfoResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchDomainRulesRequest() (request *SwitchDomainRulesRequest) {
	request = &SwitchDomainRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "SwitchDomainRules")
	return
}

func NewSwitchDomainRulesResponse() (response *SwitchDomainRulesResponse) {
	response = &SwitchDomainRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换域名的规则开关
func (c *Client) SwitchDomainRules(request *SwitchDomainRulesRequest) (response *SwitchDomainRulesResponse, err error) {
	if request == nil {
		request = NewSwitchDomainRulesRequest()
	}
	response = NewSwitchDomainRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSDetectDataChartRequest() (request *DescribeDNSDetectDataChartRequest) {
	request = &DescribeDNSDetectDataChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDNSDetectDataChart")
	return
}

func NewDescribeDNSDetectDataChartResponse() (response *DescribeDNSDetectDataChartResponse) {
	response = &DescribeDNSDetectDataChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DNS劫持检测的被劫持用户数的数据，对应【Web安全检测】【DNS劫持检测】【数据统计】【被劫持用户数】这个功能。
func (c *Client) DescribeDNSDetectDataChart(request *DescribeDNSDetectDataChartRequest) (response *DescribeDNSDetectDataChartResponse, err error) {
	if request == nil {
		request = NewDescribeDNSDetectDataChartRequest()
	}
	response = NewDescribeDNSDetectDataChartResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccessExportRequest() (request *DeleteAccessExportRequest) {
	request = &DeleteAccessExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteAccessExport")
	return
}

func NewDeleteAccessExportResponse() (response *DeleteAccessExportResponse) {
	response = &DeleteAccessExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除访问日志导出
func (c *Client) DeleteAccessExport(request *DeleteAccessExportRequest) (response *DeleteAccessExportResponse, err error) {
	if request == nil {
		request = NewDeleteAccessExportRequest()
	}
	response = NewDeleteAccessExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessDownloadRecordsRequest() (request *DescribeAccessDownloadRecordsRequest) {
	request = &DescribeAccessDownloadRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessDownloadRecords")
	return
}

func NewDescribeAccessDownloadRecordsResponse() (response *DescribeAccessDownloadRecordsResponse) {
	response = &DescribeAccessDownloadRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询以创建的访问日志下载任务
func (c *Client) DescribeAccessDownloadRecords(request *DescribeAccessDownloadRecordsRequest) (response *DescribeAccessDownloadRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeAccessDownloadRecordsRequest()
	}
	response = NewDescribeAccessDownloadRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessMetaRequest() (request *DescribeAccessMetaRequest) {
	request = &DescribeAccessMetaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessMeta")
	return
}

func NewDescribeAccessMetaResponse() (response *DescribeAccessMetaResponse) {
	response = &DescribeAccessMetaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取访问日志元数据信息
func (c *Client) DescribeAccessMeta(request *DescribeAccessMetaRequest) (response *DescribeAccessMetaResponse, err error) {
	if request == nil {
		request = NewDescribeAccessMetaRequest()
	}
	response = NewDescribeAccessMetaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiStatPointsRequest() (request *DescribeApiStatPointsRequest) {
	request = &DescribeApiStatPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiStatPoints")
	return
}

func NewDescribeApiStatPointsResponse() (response *DescribeApiStatPointsResponse) {
	response = &DescribeApiStatPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取api的访问趋势图
func (c *Client) DescribeApiStatPoints(request *DescribeApiStatPointsRequest) (response *DescribeApiStatPointsResponse, err error) {
	if request == nil {
		request = NewDescribeApiStatPointsRequest()
	}
	response = NewDescribeApiStatPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTsResourceRequest() (request *DescribeTsResourceRequest) {
	request = &DescribeTsResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeTsResource")
	return
}

func NewDescribeTsResourceResponse() (response *DescribeTsResourceResponse) {
	response = &DescribeTsResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供运营后台获取自研引擎服务资源列表
func (c *Client) DescribeTsResource(request *DescribeTsResourceRequest) (response *DescribeTsResourceResponse, err error) {
	if request == nil {
		request = NewDescribeTsResourceRequest()
	}
	response = NewDescribeTsResourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceAttackLogPostRequest() (request *ModifyInstanceAttackLogPostRequest) {
	request = &ModifyInstanceAttackLogPostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceAttackLogPost")
	return
}

func NewModifyInstanceAttackLogPostResponse() (response *ModifyInstanceAttackLogPostResponse) {
	response = &ModifyInstanceAttackLogPostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例攻击日志投递开关，企业版及以上版本可以开通，否则返回错误
func (c *Client) ModifyInstanceAttackLogPost(request *ModifyInstanceAttackLogPostRequest) (response *ModifyInstanceAttackLogPostResponse, err error) {
	if request == nil {
		request = NewModifyInstanceAttackLogPostRequest()
	}
	response = NewModifyInstanceAttackLogPostResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserDefaultEngineRequest() (request *ModifyUserDefaultEngineRequest) {
	request = &ModifyUserDefaultEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyUserDefaultEngine")
	return
}

func NewModifyUserDefaultEngineResponse() (response *ModifyUserDefaultEngineResponse) {
	response = &ModifyUserDefaultEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户默认规则引擎类型
func (c *Client) ModifyUserDefaultEngine(request *ModifyUserDefaultEngineRequest) (response *ModifyUserDefaultEngineResponse, err error) {
	if request == nil {
		request = NewModifyUserDefaultEngineRequest()
	}
	response = NewModifyUserDefaultEngineResponse()
	err = c.Send(request, response)
	return
}

func NewAddAreaBanAreasRequest() (request *AddAreaBanAreasRequest) {
	request = &AddAreaBanAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddAreaBanAreas")
	return
}

func NewAddAreaBanAreasResponse() (response *AddAreaBanAreasResponse) {
	response = &AddAreaBanAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加地域封禁中的地域信息
func (c *Client) AddAreaBanAreas(request *AddAreaBanAreasRequest) (response *AddAreaBanAreasResponse, err error) {
	if request == nil {
		request = NewAddAreaBanAreasRequest()
	}
	response = NewAddAreaBanAreasResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDomainWhiteRulesRequest() (request *DeleteDomainWhiteRulesRequest) {
	request = &DeleteDomainWhiteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteDomainWhiteRules")
	return
}

func NewDeleteDomainWhiteRulesResponse() (response *DeleteDomainWhiteRulesResponse) {
	response = &DeleteDomainWhiteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除域名规则白名单
func (c *Client) DeleteDomainWhiteRules(request *DeleteDomainWhiteRulesRequest) (response *DeleteDomainWhiteRulesResponse, err error) {
	if request == nil {
		request = NewDeleteDomainWhiteRulesRequest()
	}
	response = NewDeleteDomainWhiteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotDetailVersionTwoRequest() (request *DescribeBotDetailVersionTwoRequest) {
	request = &DescribeBotDetailVersionTwoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotDetailVersionTwo")
	return
}

func NewDescribeBotDetailVersionTwoResponse() (response *DescribeBotDetailVersionTwoResponse) {
	response = &DescribeBotDetailVersionTwoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot详情显示v2版本
func (c *Client) DescribeBotDetailVersionTwo(request *DescribeBotDetailVersionTwoRequest) (response *DescribeBotDetailVersionTwoResponse, err error) {
	if request == nil {
		request = NewDescribeBotDetailVersionTwoRequest()
	}
	response = NewDescribeBotDetailVersionTwoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecEventListRequest() (request *DescribeApiSecEventListRequest) {
	request = &DescribeApiSecEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiSecEventList")
	return
}

func NewDescribeApiSecEventListResponse() (response *DescribeApiSecEventListResponse) {
	response = &DescribeApiSecEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全事件列表
func (c *Client) DescribeApiSecEventList(request *DescribeApiSecEventListRequest) (response *DescribeApiSecEventListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecEventListRequest()
	}
	response = NewDescribeApiSecEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSceneListRequest() (request *DescribeBotSceneListRequest) {
	request = &DescribeBotSceneListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotSceneList")
	return
}

func NewDescribeBotSceneListResponse() (response *DescribeBotSceneListResponse) {
	response = &DescribeBotSceneListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT场景列表与概览
func (c *Client) DescribeBotSceneList(request *DescribeBotSceneListRequest) (response *DescribeBotSceneListResponse, err error) {
	if request == nil {
		request = NewDescribeBotSceneListRequest()
	}
	response = NewDescribeBotSceneListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainCountInfoRequest() (request *DescribeDomainCountInfoRequest) {
	request = &DescribeDomainCountInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainCountInfo")
	return
}

func NewDescribeDomainCountInfoResponse() (response *DescribeDomainCountInfoResponse) {
	response = &DescribeDomainCountInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名概况
func (c *Client) DescribeDomainCountInfo(request *DescribeDomainCountInfoRequest) (response *DescribeDomainCountInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDomainCountInfoRequest()
	}
	response = NewDescribeDomainCountInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainVerifyResultRequest() (request *DescribeDomainVerifyResultRequest) {
	request = &DescribeDomainVerifyResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainVerifyResult")
	return
}

func NewDescribeDomainVerifyResultResponse() (response *DescribeDomainVerifyResultResponse) {
	response = &DescribeDomainVerifyResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取添加域名操作的结果
func (c *Client) DescribeDomainVerifyResult(request *DescribeDomainVerifyResultRequest) (response *DescribeDomainVerifyResultResponse, err error) {
	if request == nil {
		request = NewDescribeDomainVerifyResultRequest()
	}
	response = NewDescribeDomainVerifyResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiDetailRequest() (request *DescribeApiDetailRequest) {
	request = &DescribeApiDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiDetail")
	return
}

func NewDescribeApiDetailResponse() (response *DescribeApiDetailResponse) {
	response = &DescribeApiDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Api请求详情信息
func (c *Client) DescribeApiDetail(request *DescribeApiDetailRequest) (response *DescribeApiDetailResponse, err error) {
	if request == nil {
		request = NewDescribeApiDetailRequest()
	}
	response = NewDescribeApiDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCopyCustomWhiteRuleRequest() (request *CreateCopyCustomWhiteRuleRequest) {
	request = &CreateCopyCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateCopyCustomWhiteRule")
	return
}

func NewCreateCopyCustomWhiteRuleResponse() (response *CreateCopyCustomWhiteRuleResponse) {
	response = &CreateCopyCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制精准白名单到其他域名
func (c *Client) CreateCopyCustomWhiteRule(request *CreateCopyCustomWhiteRuleRequest) (response *CreateCopyCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewCreateCopyCustomWhiteRuleRequest()
	}
	response = NewCreateCopyCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserInstanceCountRequest() (request *DescribeUserInstanceCountRequest) {
	request = &DescribeUserInstanceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserInstanceCount")
	return
}

func NewDescribeUserInstanceCountResponse() (response *DescribeUserInstanceCountResponse) {
	response = &DescribeUserInstanceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过客户appid获取客户在所有支持地域的实例数量，返回一个list，key为地域，value为数量
func (c *Client) DescribeUserInstanceCount(request *DescribeUserInstanceCountRequest) (response *DescribeUserInstanceCountResponse, err error) {
	if request == nil {
		request = NewDescribeUserInstanceCountRequest()
	}
	response = NewDescribeUserInstanceCountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHostInnerRequest() (request *CreateHostInnerRequest) {
	request = &CreateHostInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateHostInner")
	return
}

func NewCreateHostInnerResponse() (response *CreateHostInnerResponse) {
	response = &CreateHostInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf中添加防护的域名
func (c *Client) CreateHostInner(request *CreateHostInnerRequest) (response *CreateHostInnerResponse, err error) {
	if request == nil {
		request = NewCreateHostInnerRequest()
	}
	response = NewCreateHostInnerResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWafAutoDenyRulesRequest() (request *ModifyWafAutoDenyRulesRequest) {
	request = &ModifyWafAutoDenyRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyWafAutoDenyRules")
	return
}

func NewModifyWafAutoDenyRulesResponse() (response *ModifyWafAutoDenyRulesResponse) {
	response = &ModifyWafAutoDenyRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改ip惩罚规则
func (c *Client) ModifyWafAutoDenyRules(request *ModifyWafAutoDenyRulesRequest) (response *ModifyWafAutoDenyRulesResponse, err error) {
	if request == nil {
		request = NewModifyWafAutoDenyRulesRequest()
	}
	response = NewModifyWafAutoDenyRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCCRuleRequest() (request *DescribeCCRuleRequest) {
	request = &DescribeCCRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCCRule")
	return
}

func NewDescribeCCRuleResponse() (response *DescribeCCRuleResponse) {
	response = &DescribeCCRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  CC V2 Query接口
func (c *Client) DescribeCCRule(request *DescribeCCRuleRequest) (response *DescribeCCRuleResponse, err error) {
	if request == nil {
		request = NewDescribeCCRuleRequest()
	}
	response = NewDescribeCCRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafRequestCountRequest() (request *DescribeWafRequestCountRequest) {
	request = &DescribeWafRequestCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeWafRequestCount")
	return
}

func NewDescribeWafRequestCountResponse() (response *DescribeWafRequestCountResponse) {
	response = &DescribeWafRequestCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取waf的访问日志总条数
func (c *Client) DescribeWafRequestCount(request *DescribeWafRequestCountRequest) (response *DescribeWafRequestCountResponse, err error) {
	if request == nil {
		request = NewDescribeWafRequestCountRequest()
	}
	response = NewDescribeWafRequestCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTiRuleRequest() (request *DescribeBotTiRuleRequest) {
	request = &DescribeBotTiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotTiRule")
	return
}

func NewDescribeBotTiRuleResponse() (response *DescribeBotTiRuleResponse) {
	response = &DescribeBotTiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot模块的威胁情报配置
func (c *Client) DescribeBotTiRule(request *DescribeBotTiRuleRequest) (response *DescribeBotTiRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotTiRuleRequest()
	}
	response = NewDescribeBotTiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVIPRSRequest() (request *DescribeVIPRSRequest) {
	request = &DescribeVIPRSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeVIPRS")
	return
}

func NewDescribeVIPRSResponse() (response *DescribeVIPRSResponse) {
	response = &DescribeVIPRSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询详细的vip信息
func (c *Client) DescribeVIPRS(request *DescribeVIPRSRequest) (response *DescribeVIPRSResponse, err error) {
	if request == nil {
		request = NewDescribeVIPRSRequest()
	}
	response = NewDescribeVIPRSResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopStraceIPRequest() (request *GetTopStraceIPRequest) {
	request = &GetTopStraceIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetTopStraceIP")
	return
}

func NewGetTopStraceIPResponse() (response *GetTopStraceIPResponse) {
	response = &GetTopStraceIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防护IP的TOP10
func (c *Client) GetTopStraceIP(request *GetTopStraceIPRequest) (response *GetTopStraceIPResponse, err error) {
	if request == nil {
		request = NewGetTopStraceIPRequest()
	}
	response = NewGetTopStraceIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotStatisticPointsRequest() (request *DescribeBotStatisticPointsRequest) {
	request = &DescribeBotStatisticPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotStatisticPoints")
	return
}

func NewDescribeBotStatisticPointsResponse() (response *DescribeBotStatisticPointsResponse) {
	response = &DescribeBotStatisticPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 Bot流量统计
func (c *Client) DescribeBotStatisticPoints(request *DescribeBotStatisticPointsRequest) (response *DescribeBotStatisticPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotStatisticPointsRequest()
	}
	response = NewDescribeBotStatisticPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainDetailsSaasRequest() (request *DescribeDomainDetailsSaasRequest) {
	request = &DescribeDomainDetailsSaasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainDetailsSaas")
	return
}

func NewDescribeDomainDetailsSaasResponse() (response *DescribeDomainDetailsSaasResponse) {
	response = &DescribeDomainDetailsSaasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单个saas域名详情
func (c *Client) DescribeDomainDetailsSaas(request *DescribeDomainDetailsSaasRequest) (response *DescribeDomainDetailsSaasResponse, err error) {
	if request == nil {
		request = NewDescribeDomainDetailsSaasRequest()
	}
	response = NewDescribeDomainDetailsSaasResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAntiFakeUrlStatusRequest() (request *ModifyAntiFakeUrlStatusRequest) {
	request = &ModifyAntiFakeUrlStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiFakeUrlStatus")
	return
}

func NewModifyAntiFakeUrlStatusResponse() (response *ModifyAntiFakeUrlStatusResponse) {
	response = &ModifyAntiFakeUrlStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换防篡改开关
func (c *Client) ModifyAntiFakeUrlStatus(request *ModifyAntiFakeUrlStatusRequest) (response *ModifyAntiFakeUrlStatusResponse, err error) {
	if request == nil {
		request = NewModifyAntiFakeUrlStatusRequest()
	}
	response = NewModifyAntiFakeUrlStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAntiInfoLeakRulesRequest() (request *ModifyAntiInfoLeakRulesRequest) {
	request = &ModifyAntiInfoLeakRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiInfoLeakRules")
	return
}

func NewModifyAntiInfoLeakRulesResponse() (response *ModifyAntiInfoLeakRulesResponse) {
	response = &ModifyAntiInfoLeakRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑信息防泄漏规则
func (c *Client) ModifyAntiInfoLeakRules(request *ModifyAntiInfoLeakRulesRequest) (response *ModifyAntiInfoLeakRulesResponse, err error) {
	if request == nil {
		request = NewModifyAntiInfoLeakRulesRequest()
	}
	response = NewModifyAntiInfoLeakRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpAccessControlRequest() (request *DescribeIpAccessControlRequest) {
	request = &DescribeIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeIpAccessControl")
	return
}

func NewDescribeIpAccessControlResponse() (response *DescribeIpAccessControlResponse) {
	response = &DescribeIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf ip黑白名单查询
func (c *Client) DescribeIpAccessControl(request *DescribeIpAccessControlRequest) (response *DescribeIpAccessControlResponse, err error) {
	if request == nil {
		request = NewDescribeIpAccessControlRequest()
	}
	response = NewDescribeIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWhiteAllDomainRequest() (request *DescribeWhiteAllDomainRequest) {
	request = &DescribeWhiteAllDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeWhiteAllDomain")
	return
}

func NewDescribeWhiteAllDomainResponse() (response *DescribeWhiteAllDomainResponse) {
	response = &DescribeWhiteAllDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询saas和clb的域名信息v2
func (c *Client) DescribeWhiteAllDomain(request *DescribeWhiteAllDomainRequest) (response *DescribeWhiteAllDomainResponse, err error) {
	if request == nil {
		request = NewDescribeWhiteAllDomainRequest()
	}
	response = NewDescribeWhiteAllDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAntiInfoLeakRuleRequest() (request *DeleteAntiInfoLeakRuleRequest) {
	request = &DeleteAntiInfoLeakRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteAntiInfoLeakRule")
	return
}

func NewDeleteAntiInfoLeakRuleResponse() (response *DeleteAntiInfoLeakRuleResponse) {
	response = &DeleteAntiInfoLeakRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 信息防泄漏删除规则
func (c *Client) DeleteAntiInfoLeakRule(request *DeleteAntiInfoLeakRuleRequest) (response *DeleteAntiInfoLeakRuleResponse, err error) {
	if request == nil {
		request = NewDeleteAntiInfoLeakRuleRequest()
	}
	response = NewDeleteAntiInfoLeakRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackTotalCountRequest() (request *GetAttackTotalCountRequest) {
	request = &GetAttackTotalCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetAttackTotalCount")
	return
}

func NewGetAttackTotalCountResponse() (response *GetAttackTotalCountResponse) {
	response = &GetAttackTotalCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照条件查询展示攻击总次数
func (c *Client) GetAttackTotalCount(request *GetAttackTotalCountRequest) (response *GetAttackTotalCountResponse, err error) {
	if request == nil {
		request = NewGetAttackTotalCountRequest()
	}
	response = NewGetAttackTotalCountResponse()
	err = c.Send(request, response)
	return
}

func NewAddSpartaProtectionAutoRequest() (request *AddSpartaProtectionAutoRequest) {
	request = &AddSpartaProtectionAutoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddSpartaProtectionAuto")
	return
}

func NewAddSpartaProtectionAutoResponse() (response *AddSpartaProtectionAutoResponse) {
	response = &AddSpartaProtectionAutoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键接入
func (c *Client) AddSpartaProtectionAuto(request *AddSpartaProtectionAutoRequest) (response *AddSpartaProtectionAutoResponse, err error) {
	if request == nil {
		request = NewAddSpartaProtectionAutoRequest()
	}
	response = NewAddSpartaProtectionAutoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRegionsStatRequest() (request *DescribeBotRegionsStatRequest) {
	request = &DescribeBotRegionsStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotRegionsStat")
	return
}

func NewDescribeBotRegionsStatResponse() (response *DescribeBotRegionsStatResponse) {
	response = &DescribeBotRegionsStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 bot地理纬度统计
func (c *Client) DescribeBotRegionsStat(request *DescribeBotRegionsStatRequest) (response *DescribeBotRegionsStatResponse, err error) {
	if request == nil {
		request = NewDescribeBotRegionsStatRequest()
	}
	response = NewDescribeBotRegionsStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRolePlayStateRequest() (request *DescribeRolePlayStateRequest) {
	request = &DescribeRolePlayStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeRolePlayState")
	return
}

func NewDescribeRolePlayStateResponse() (response *DescribeRolePlayStateResponse) {
	response = &DescribeRolePlayStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户授权的信息
func (c *Client) DescribeRolePlayState(request *DescribeRolePlayStateRequest) (response *DescribeRolePlayStateResponse, err error) {
	if request == nil {
		request = NewDescribeRolePlayStateRequest()
	}
	response = NewDescribeRolePlayStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePeakPointsRequest() (request *DescribePeakPointsRequest) {
	request = &DescribePeakPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePeakPoints")
	return
}

func NewDescribePeakPointsResponse() (response *DescribePeakPointsResponse) {
	response = &DescribePeakPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询业务和攻击概要趋势
func (c *Client) DescribePeakPoints(request *DescribePeakPointsRequest) (response *DescribePeakPointsResponse, err error) {
	if request == nil {
		request = NewDescribePeakPointsRequest()
	}
	response = NewDescribePeakPointsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostInnerRequest() (request *ModifyHostInnerRequest) {
	request = &ModifyHostInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyHostInner")
	return
}

func NewModifyHostInnerResponse() (response *ModifyHostInnerResponse) {
	response = &ModifyHostInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf编辑防护域名配置内网
func (c *Client) ModifyHostInner(request *ModifyHostInnerRequest) (response *ModifyHostInnerResponse, err error) {
	if request == nil {
		request = NewModifyHostInnerRequest()
	}
	response = NewModifyHostInnerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostPointsRequest() (request *DescribePostPointsRequest) {
	request = &DescribePostPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePostPoints")
	return
}

func NewDescribePostPointsResponse() (response *DescribePostPointsResponse) {
	response = &DescribePostPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取投递流量趋势图
func (c *Client) DescribePostPoints(request *DescribePostPointsRequest) (response *DescribePostPointsResponse, err error) {
	if request == nil {
		request = NewDescribePostPointsRequest()
	}
	response = NewDescribePostPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotFirstPurchaseRequest() (request *DescribeBotFirstPurchaseRequest) {
	request = &DescribeBotFirstPurchaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotFirstPurchase")
	return
}

func NewDescribeBotFirstPurchaseResponse() (response *DescribeBotFirstPurchaseResponse) {
	response = &DescribeBotFirstPurchaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户是否为第一次购买waf实例，api接口会调用计费侧查询订单接口
func (c *Client) DescribeBotFirstPurchase(request *DescribeBotFirstPurchaseRequest) (response *DescribeBotFirstPurchaseResponse, err error) {
	if request == nil {
		request = NewDescribeBotFirstPurchaseRequest()
	}
	response = NewDescribeBotFirstPurchaseResponse()
	err = c.Send(request, response)
	return
}

func NewSearchAttackLogRequest() (request *SearchAttackLogRequest) {
	request = &SearchAttackLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "SearchAttackLog")
	return
}

func NewSearchAttackLogResponse() (response *SearchAttackLogResponse) {
	response = &SearchAttackLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新版本CLS接口存在参数变化，query改成了query_string支持lucence语法接口搜索查询。
func (c *Client) SearchAttackLog(request *SearchAttackLogRequest) (response *SearchAttackLogResponse, err error) {
	if request == nil {
		request = NewSearchAttackLogRequest()
	}
	response = NewSearchAttackLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotStatusRequest() (request *ModifyBotStatusRequest) {
	request = &ModifyBotStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotStatus")
	return
}

func NewModifyBotStatusResponse() (response *ModifyBotStatusResponse) {
	response = &ModifyBotStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 bot总开关更新
func (c *Client) ModifyBotStatus(request *ModifyBotStatusRequest) (response *ModifyBotStatusResponse, err error) {
	if request == nil {
		request = NewModifyBotStatusRequest()
	}
	response = NewModifyBotStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClsForAfterpayRequest() (request *CreateClsForAfterpayRequest) {
	request = &CreateClsForAfterpayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateClsForAfterpay")
	return
}

func NewCreateClsForAfterpayResponse() (response *CreateClsForAfterpayResponse) {
	response = &CreateClsForAfterpayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费创建日志服务
func (c *Client) CreateClsForAfterpay(request *CreateClsForAfterpayRequest) (response *CreateClsForAfterpayResponse, err error) {
	if request == nil {
		request = NewCreateClsForAfterpayRequest()
	}
	response = NewCreateClsForAfterpayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCertificatedDomainRequest() (request *ModifyCertificatedDomainRequest) {
	request = &ModifyCertificatedDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyCertificatedDomain")
	return
}

func NewModifyCertificatedDomainResponse() (response *ModifyCertificatedDomainResponse) {
	response = &ModifyCertificatedDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名的证书
func (c *Client) ModifyCertificatedDomain(request *ModifyCertificatedDomainRequest) (response *ModifyCertificatedDomainResponse, err error) {
	if request == nil {
		request = NewModifyCertificatedDomainRequest()
	}
	response = NewModifyCertificatedDomainResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchElasticModeRequest() (request *SwitchElasticModeRequest) {
	request = &SwitchElasticModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "SwitchElasticMode")
	return
}

func NewSwitchElasticModeResponse() (response *SwitchElasticModeResponse) {
	response = &SwitchElasticModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换弹性的开关
func (c *Client) SwitchElasticMode(request *SwitchElasticModeRequest) (response *SwitchElasticModeResponse, err error) {
	if request == nil {
		request = NewSwitchElasticModeRequest()
	}
	response = NewSwitchElasticModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMainClassRequest() (request *DescribeMainClassRequest) {
	request = &DescribeMainClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeMainClass")
	return
}

func NewDescribeMainClassResponse() (response *DescribeMainClassResponse) {
	response = &DescribeMainClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主类及子类信息
func (c *Client) DescribeMainClass(request *DescribeMainClassRequest) (response *DescribeMainClassResponse, err error) {
	if request == nil {
		request = NewDescribeMainClassRequest()
	}
	response = NewDescribeMainClassResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserSignaturePolicyStatusRequest() (request *GetUserSignaturePolicyStatusRequest) {
	request = &GetUserSignaturePolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetUserSignaturePolicyStatus")
	return
}

func NewGetUserSignaturePolicyStatusResponse() (response *GetUserSignaturePolicyStatusResponse) {
	response = &GetUserSignaturePolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户防护规则开关
func (c *Client) GetUserSignaturePolicyStatus(request *GetUserSignaturePolicyStatusRequest) (response *GetUserSignaturePolicyStatusResponse, err error) {
	if request == nil {
		request = NewGetUserSignaturePolicyStatusRequest()
	}
	response = NewGetUserSignaturePolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostRequest() (request *ModifyHostRequest) {
	request = &ModifyHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyHost")
	return
}

func NewModifyHostResponse() (response *ModifyHostResponse) {
	response = &ModifyHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf编辑防护域名配置
func (c *Client) ModifyHost(request *ModifyHostRequest) (response *ModifyHostResponse, err error) {
	if request == nil {
		request = NewModifyHostRequest()
	}
	response = NewModifyHostResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAreaBanStatusRequest() (request *ModifyAreaBanStatusRequest) {
	request = &ModifyAreaBanStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAreaBanStatus")
	return
}

func NewModifyAreaBanStatusResponse() (response *ModifyAreaBanStatusResponse) {
	response = &ModifyAreaBanStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改防护域名的地域封禁状态
func (c *Client) ModifyAreaBanStatus(request *ModifyAreaBanStatusRequest) (response *ModifyAreaBanStatusResponse, err error) {
	if request == nil {
		request = NewModifyAreaBanStatusRequest()
	}
	response = NewModifyAreaBanStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotUBRecordsRequest() (request *DescribeBotUBRecordsRequest) {
	request = &DescribeBotUBRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotUBRecords")
	return
}

func NewDescribeBotUBRecordsResponse() (response *DescribeBotUBRecordsResponse) {
	response = &DescribeBotUBRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取UB类型Bots
func (c *Client) DescribeBotUBRecords(request *DescribeBotUBRecordsRequest) (response *DescribeBotUBRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeBotUBRecordsRequest()
	}
	response = NewDescribeBotUBRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRequestCountRequest() (request *DescribeRequestCountRequest) {
	request = &DescribeRequestCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeRequestCount")
	return
}

func NewDescribeRequestCountResponse() (response *DescribeRequestCountResponse) {
	response = &DescribeRequestCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定时间段内请求总数，这个接口不需要客户购买日志服务
func (c *Client) DescribeRequestCount(request *DescribeRequestCountRequest) (response *DescribeRequestCountResponse, err error) {
	if request == nil {
		request = NewDescribeRequestCountRequest()
	}
	response = NewDescribeRequestCountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBusinessRiskRequest() (request *DeleteBusinessRiskRequest) {
	request = &DeleteBusinessRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteBusinessRisk")
	return
}

func NewDeleteBusinessRiskResponse() (response *DeleteBusinessRiskResponse) {
	response = &DeleteBusinessRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除业务安全策略
func (c *Client) DeleteBusinessRisk(request *DeleteBusinessRiskRequest) (response *DeleteBusinessRiskResponse, err error) {
	if request == nil {
		request = NewDeleteBusinessRiskRequest()
	}
	response = NewDeleteBusinessRiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockPagesRequest() (request *DescribeBlockPagesRequest) {
	request = &DescribeBlockPagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBlockPages")
	return
}

func NewDescribeBlockPagesResponse() (response *DescribeBlockPagesResponse) {
	response = &DescribeBlockPagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示自定义拦截页面列表信息。
func (c *Client) DescribeBlockPages(request *DescribeBlockPagesRequest) (response *DescribeBlockPagesResponse, err error) {
	if request == nil {
		request = NewDescribeBlockPagesRequest()
	}
	response = NewDescribeBlockPagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeModuleStatusRequest() (request *DescribeModuleStatusRequest) {
	request = &DescribeModuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeModuleStatus")
	return
}

func NewDescribeModuleStatusResponse() (response *DescribeModuleStatusResponse) {
	response = &DescribeModuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询各个waf基础安全模块的开关状态，看每个模块是否开启
func (c *Client) DescribeModuleStatus(request *DescribeModuleStatusRequest) (response *DescribeModuleStatusResponse, err error) {
	if request == nil {
		request = NewDescribeModuleStatusRequest()
	}
	response = NewDescribeModuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAiRuleRequest() (request *DescribeBotAiRuleRequest) {
	request = &DescribeBotAiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotAiRule")
	return
}

func NewDescribeBotAiRuleResponse() (response *DescribeBotAiRuleResponse) {
	response = &DescribeBotAiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot-ai的配置信息
func (c *Client) DescribeBotAiRule(request *DescribeBotAiRuleRequest) (response *DescribeBotAiRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotAiRuleRequest()
	}
	response = NewDescribeBotAiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainsCLSStatusRequest() (request *ModifyDomainsCLSStatusRequest) {
	request = &ModifyDomainsCLSStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainsCLSStatus")
	return
}

func NewModifyDomainsCLSStatusResponse() (response *ModifyDomainsCLSStatusResponse) {
	response = &ModifyDomainsCLSStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名列表的访问日志开关
func (c *Client) ModifyDomainsCLSStatus(request *ModifyDomainsCLSStatusRequest) (response *ModifyDomainsCLSStatusResponse, err error) {
	if request == nil {
		request = NewModifyDomainsCLSStatusRequest()
	}
	response = NewModifyDomainsCLSStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyThreatIntelligenceStatusRequest() (request *ModifyThreatIntelligenceStatusRequest) {
	request = &ModifyThreatIntelligenceStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyThreatIntelligenceStatus")
	return
}

func NewModifyThreatIntelligenceStatusResponse() (response *ModifyThreatIntelligenceStatusResponse) {
	response = &ModifyThreatIntelligenceStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置威胁情报的开关配置
func (c *Client) ModifyThreatIntelligenceStatus(request *ModifyThreatIntelligenceStatusRequest) (response *ModifyThreatIntelligenceStatusResponse, err error) {
	if request == nil {
		request = NewModifyThreatIntelligenceStatusRequest()
	}
	response = NewModifyThreatIntelligenceStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotTokenRequest() (request *ModifyBotTokenRequest) {
	request = &ModifyBotTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotToken")
	return
}

func NewModifyBotTokenResponse() (response *ModifyBotTokenResponse) {
	response = &ModifyBotTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot会话token配置v2
func (c *Client) ModifyBotToken(request *ModifyBotTokenRequest) (response *ModifyBotTokenResponse, err error) {
	if request == nil {
		request = NewModifyBotTokenRequest()
	}
	response = NewModifyBotTokenResponse()
	err = c.Send(request, response)
	return
}

func NewGetStraceCountRequest() (request *GetStraceCountRequest) {
	request = &GetStraceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetStraceCount")
	return
}

func NewGetStraceCountResponse() (response *GetStraceCountResponse) {
	response = &GetStraceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防护总量
func (c *Client) GetStraceCount(request *GetStraceCountRequest) (response *GetStraceCountResponse, err error) {
	if request == nil {
		request = NewGetStraceCountRequest()
	}
	response = NewGetStraceCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonitorDomainsRequest() (request *DescribeMonitorDomainsRequest) {
	request = &DescribeMonitorDomainsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeMonitorDomains")
	return
}

func NewDescribeMonitorDomainsResponse() (response *DescribeMonitorDomainsResponse) {
	response = &DescribeMonitorDomainsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Barad云监控的域名列表
func (c *Client) DescribeMonitorDomains(request *DescribeMonitorDomainsRequest) (response *DescribeMonitorDomainsResponse, err error) {
	if request == nil {
		request = NewDescribeMonitorDomainsRequest()
	}
	response = NewDescribeMonitorDomainsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceRegionRequest() (request *ModifyInstanceRegionRequest) {
	request = &ModifyInstanceRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceRegion")
	return
}

func NewModifyInstanceRegionResponse() (response *ModifyInstanceRegionResponse) {
	response = &ModifyInstanceRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例的地域属性
func (c *Client) ModifyInstanceRegion(request *ModifyInstanceRegionRequest) (response *ModifyInstanceRegionResponse, err error) {
	if request == nil {
		request = NewModifyInstanceRegionRequest()
	}
	response = NewModifyInstanceRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyPostCLSFlowRequest() (request *DestroyPostCLSFlowRequest) {
	request = &DestroyPostCLSFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DestroyPostCLSFlow")
	return
}

func NewDestroyPostCLSFlowResponse() (response *DestroyPostCLSFlowResponse) {
	response = &DestroyPostCLSFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁CLS投递流任务
func (c *Client) DestroyPostCLSFlow(request *DestroyPostCLSFlowRequest) (response *DestroyPostCLSFlowResponse, err error) {
	if request == nil {
		request = NewDestroyPostCLSFlowRequest()
	}
	response = NewDestroyPostCLSFlowResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackLogCountRequest() (request *DescribeAttackLogCountRequest) {
	request = &DescribeAttackLogCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackLogCount")
	return
}

func NewDescribeAttackLogCountResponse() (response *DescribeAttackLogCountResponse) {
	response = &DescribeAttackLogCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询符合搜索条件攻击日志数量
func (c *Client) DescribeAttackLogCount(request *DescribeAttackLogCountRequest) (response *DescribeAttackLogCountResponse, err error) {
	if request == nil {
		request = NewDescribeAttackLogCountRequest()
	}
	response = NewDescribeAttackLogCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserEngineTypeRequest() (request *DescribeUserEngineTypeRequest) {
	request = &DescribeUserEngineTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserEngineType")
	return
}

func NewDescribeUserEngineTypeResponse() (response *DescribeUserEngineTypeResponse) {
	response = &DescribeUserEngineTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户规则引擎类型
func (c *Client) DescribeUserEngineType(request *DescribeUserEngineTypeRequest) (response *DescribeUserEngineTypeResponse, err error) {
	if request == nil {
		request = NewDescribeUserEngineTypeRequest()
	}
	response = NewDescribeUserEngineTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomWhiteRuleStatusRequest() (request *ModifyCustomWhiteRuleStatusRequest) {
	request = &ModifyCustomWhiteRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomWhiteRuleStatus")
	return
}

func NewModifyCustomWhiteRuleStatusResponse() (response *ModifyCustomWhiteRuleStatusResponse) {
	response = &ModifyCustomWhiteRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启或禁用精准白名单
func (c *Client) ModifyCustomWhiteRuleStatus(request *ModifyCustomWhiteRuleStatusRequest) (response *ModifyCustomWhiteRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyCustomWhiteRuleStatusRequest()
	}
	response = NewModifyCustomWhiteRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotWhiteStatusRequest() (request *ModifyBotWhiteStatusRequest) {
	request = &ModifyBotWhiteStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotWhiteStatus")
	return
}

func NewModifyBotWhiteStatusResponse() (response *ModifyBotWhiteStatusResponse) {
	response = &ModifyBotWhiteStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改goodbot的开关
func (c *Client) ModifyBotWhiteStatus(request *ModifyBotWhiteStatusRequest) (response *ModifyBotWhiteStatusResponse, err error) {
	if request == nil {
		request = NewModifyBotWhiteStatusRequest()
	}
	response = NewModifyBotWhiteStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAddSpartaProtectionRequest() (request *AddSpartaProtectionRequest) {
	request = &AddSpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddSpartaProtection")
	return
}

func NewAddSpartaProtectionResponse() (response *AddSpartaProtectionResponse) {
	response = &AddSpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加Spart防护域名
func (c *Client) AddSpartaProtection(request *AddSpartaProtectionRequest) (response *AddSpartaProtectionResponse, err error) {
	if request == nil {
		request = NewAddSpartaProtectionRequest()
	}
	response = NewAddSpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewAddAntiInfoLeakRulesRequest() (request *AddAntiInfoLeakRulesRequest) {
	request = &AddAntiInfoLeakRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddAntiInfoLeakRules")
	return
}

func NewAddAntiInfoLeakRulesResponse() (response *AddAntiInfoLeakRulesResponse) {
	response = &AddAntiInfoLeakRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加信息防泄漏规则
func (c *Client) AddAntiInfoLeakRules(request *AddAntiInfoLeakRulesRequest) (response *AddAntiInfoLeakRulesResponse, err error) {
	if request == nil {
		request = NewAddAntiInfoLeakRulesRequest()
	}
	response = NewAddAntiInfoLeakRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRuleOverviewRequest() (request *DescribeBotRuleOverviewRequest) {
	request = &DescribeBotRuleOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotRuleOverview")
	return
}

func NewDescribeBotRuleOverviewResponse() (response *DescribeBotRuleOverviewResponse) {
	response = &DescribeBotRuleOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot的规则概览
func (c *Client) DescribeBotRuleOverview(request *DescribeBotRuleOverviewRequest) (response *DescribeBotRuleOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBotRuleOverviewRequest()
	}
	response = NewDescribeBotRuleOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVIPConfigsRequest() (request *ModifyVIPConfigsRequest) {
	request = &ModifyVIPConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyVIPConfigs")
	return
}

func NewModifyVIPConfigsResponse() (response *ModifyVIPConfigsResponse) {
	response = &ModifyVIPConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改配置
func (c *Client) ModifyVIPConfigs(request *ModifyVIPConfigsRequest) (response *ModifyVIPConfigsResponse, err error) {
	if request == nil {
		request = NewModifyVIPConfigsRequest()
	}
	response = NewModifyVIPConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecEventTypeListRequest() (request *DescribeApiSecEventTypeListRequest) {
	request = &DescribeApiSecEventTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiSecEventTypeList")
	return
}

func NewDescribeApiSecEventTypeListResponse() (response *DescribeApiSecEventTypeListResponse) {
	response = &DescribeApiSecEventTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全事件分类列表
func (c *Client) DescribeApiSecEventTypeList(request *DescribeApiSecEventTypeListRequest) (response *DescribeApiSecEventTypeListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecEventTypeListRequest()
	}
	response = NewDescribeApiSecEventTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleUpdateLogRequest() (request *DescribeRuleUpdateLogRequest) {
	request = &DescribeRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeRuleUpdateLog")
	return
}

func NewDescribeRuleUpdateLogResponse() (response *DescribeRuleUpdateLogResponse) {
	response = &DescribeRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取特征规则更新动态
func (c *Client) DescribeRuleUpdateLog(request *DescribeRuleUpdateLogRequest) (response *DescribeRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewDescribeRuleUpdateLogRequest()
	}
	response = NewDescribeRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveBypassAllRuleRequest() (request *RemoveBypassAllRuleRequest) {
	request = &RemoveBypassAllRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "RemoveBypassAllRule")
	return
}

func NewRemoveBypassAllRuleResponse() (response *RemoveBypassAllRuleResponse) {
	response = &RemoveBypassAllRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除一键bypass规则
func (c *Client) RemoveBypassAllRule(request *RemoveBypassAllRuleRequest) (response *RemoveBypassAllRuleResponse, err error) {
	if request == nil {
		request = NewRemoveBypassAllRuleRequest()
	}
	response = NewRemoveBypassAllRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBlockPageRequest() (request *DeleteBlockPageRequest) {
	request = &DeleteBlockPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteBlockPage")
	return
}

func NewDeleteBlockPageResponse() (response *DeleteBlockPageResponse) {
	response = &DeleteBlockPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义拦截页面，所有应用了该拦截页面的域名都会自动删除该拦截页面。
func (c *Client) DeleteBlockPage(request *DeleteBlockPageRequest) (response *DeleteBlockPageResponse, err error) {
	if request == nil {
		request = NewDeleteBlockPageRequest()
	}
	response = NewDeleteBlockPageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGlobalWhiteRuleRequest() (request *ModifyGlobalWhiteRuleRequest) {
	request = &ModifyGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyGlobalWhiteRule")
	return
}

func NewModifyGlobalWhiteRuleResponse() (response *ModifyGlobalWhiteRuleResponse) {
	response = &ModifyGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供内部运营系统调用，增加规则引擎全局白名单列表。
func (c *Client) ModifyGlobalWhiteRule(request *ModifyGlobalWhiteRuleRequest) (response *ModifyGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyGlobalWhiteRuleRequest()
	}
	response = NewModifyGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostModeRequest() (request *ModifyHostModeRequest) {
	request = &ModifyHostModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyHostMode")
	return
}

func NewModifyHostModeResponse() (response *ModifyHostModeResponse) {
	response = &ModifyHostModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf设置防护域名防护状态
func (c *Client) ModifyHostMode(request *ModifyHostModeRequest) (response *ModifyHostModeResponse, err error) {
	if request == nil {
		request = NewModifyHostModeRequest()
	}
	response = NewModifyHostModeResponse()
	err = c.Send(request, response)
	return
}

func NewAddGlobalWhiteRuleRequest() (request *AddGlobalWhiteRuleRequest) {
	request = &AddGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddGlobalWhiteRule")
	return
}

func NewAddGlobalWhiteRuleResponse() (response *AddGlobalWhiteRuleResponse) {
	response = &AddGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供内部运营系统调用，增加Tiga规则引擎全局白名单列表。
func (c *Client) AddGlobalWhiteRule(request *AddGlobalWhiteRuleRequest) (response *AddGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddGlobalWhiteRuleRequest()
	}
	response = NewAddGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecEventChangeListRequest() (request *DescribeApiSecEventChangeListRequest) {
	request = &DescribeApiSecEventChangeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiSecEventChangeList")
	return
}

func NewDescribeApiSecEventChangeListResponse() (response *DescribeApiSecEventChangeListResponse) {
	response = &DescribeApiSecEventChangeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取api安全事件变更历史
func (c *Client) DescribeApiSecEventChangeList(request *DescribeApiSecEventChangeListRequest) (response *DescribeApiSecEventChangeListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecEventChangeListRequest()
	}
	response = NewDescribeApiSecEventChangeListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyJsInjectRuleStatusRequest() (request *ModifyJsInjectRuleStatusRequest) {
	request = &ModifyJsInjectRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyJsInjectRuleStatus")
	return
}

func NewModifyJsInjectRuleStatusResponse() (response *ModifyJsInjectRuleStatusResponse) {
	response = &ModifyJsInjectRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新前端对抗规则的状态，可以开启和关闭动态前端对抗规则。
func (c *Client) ModifyJsInjectRuleStatus(request *ModifyJsInjectRuleStatusRequest) (response *ModifyJsInjectRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyJsInjectRuleStatusRequest()
	}
	response = NewModifyJsInjectRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyObjectRequest() (request *ModifyObjectRequest) {
	request = &ModifyObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyObject")
	return
}

func NewModifyObjectResponse() (response *ModifyObjectResponse) {
	response = &ModifyObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改防护对象
func (c *Client) ModifyObject(request *ModifyObjectRequest) (response *ModifyObjectResponse, err error) {
	if request == nil {
		request = NewModifyObjectRequest()
	}
	response = NewModifyObjectResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCachePathRequest() (request *CreateCachePathRequest) {
	request = &CreateCachePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateCachePath")
	return
}

func NewCreateCachePathResponse() (response *CreateCachePathResponse) {
	response = &CreateCachePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-添加缓存路径
func (c *Client) CreateCachePath(request *CreateCachePathRequest) (response *CreateCachePathResponse, err error) {
	if request == nil {
		request = NewCreateCachePathRequest()
	}
	response = NewCreateCachePathResponse()
	err = c.Send(request, response)
	return
}

func NewAddGroupProtectionRequest() (request *AddGroupProtectionRequest) {
	request = &AddGroupProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddGroupProtection")
	return
}

func NewAddGroupProtectionResponse() (response *AddGroupProtectionResponse) {
	response = &AddGroupProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 收到qps自动告警，完成vip自动限速、迁移、回滚等操作
func (c *Client) AddGroupProtection(request *AddGroupProtectionRequest) (response *AddGroupProtectionResponse, err error) {
	if request == nil {
		request = NewAddGroupProtectionRequest()
	}
	response = NewAddGroupProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleTypesEnRequest() (request *DescribeRuleTypesEnRequest) {
	request = &DescribeRuleTypesEnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeRuleTypesEn")
	return
}

func NewDescribeRuleTypesEnResponse() (response *DescribeRuleTypesEnResponse) {
	response = &DescribeRuleTypesEnResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取规则类型描述支持英文
func (c *Client) DescribeRuleTypesEn(request *DescribeRuleTypesEnRequest) (response *DescribeRuleTypesEnResponse, err error) {
	if request == nil {
		request = NewDescribeRuleTypesEnRequest()
	}
	response = NewDescribeRuleTypesEnResponse()
	err = c.Send(request, response)
	return
}

func NewCheckGlobalWhiteRuleNameRequest() (request *CheckGlobalWhiteRuleNameRequest) {
	request = &CheckGlobalWhiteRuleNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CheckGlobalWhiteRuleName")
	return
}

func NewCheckGlobalWhiteRuleNameResponse() (response *CheckGlobalWhiteRuleNameResponse) {
	response = &CheckGlobalWhiteRuleNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台：检测全局白名单名字是否重复
func (c *Client) CheckGlobalWhiteRuleName(request *CheckGlobalWhiteRuleNameRequest) (response *CheckGlobalWhiteRuleNameResponse, err error) {
	if request == nil {
		request = NewCheckGlobalWhiteRuleNameRequest()
	}
	response = NewCheckGlobalWhiteRuleNameResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackAnalysisRequest() (request *GetAttackAnalysisRequest) {
	request = &GetAttackAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetAttackAnalysis")
	return
}

func NewGetAttackAnalysisResponse() (response *GetAttackAnalysisResponse) {
	response = &GetAttackAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CLS新版本攻击日志聚合分析排序
func (c *Client) GetAttackAnalysis(request *GetAttackAnalysisRequest) (response *GetAttackAnalysisResponse, err error) {
	if request == nil {
		request = NewGetAttackAnalysisRequest()
	}
	response = NewGetAttackAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAreaBanAreasRequest() (request *ModifyAreaBanAreasRequest) {
	request = &ModifyAreaBanAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyAreaBanAreas")
	return
}

func NewModifyAreaBanAreasResponse() (response *ModifyAreaBanAreasResponse) {
	response = &ModifyAreaBanAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改地域封禁中的地域信息
func (c *Client) ModifyAreaBanAreas(request *ModifyAreaBanAreasRequest) (response *ModifyAreaBanAreasResponse, err error) {
	if request == nil {
		request = NewModifyAreaBanAreasRequest()
	}
	response = NewModifyAreaBanAreasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiAnalyzeStatusRequest() (request *DescribeApiAnalyzeStatusRequest) {
	request = &DescribeApiAnalyzeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiAnalyzeStatus")
	return
}

func NewDescribeApiAnalyzeStatusResponse() (response *DescribeApiAnalyzeStatusResponse) {
	response = &DescribeApiAnalyzeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看api资产分析开关
func (c *Client) DescribeApiAnalyzeStatus(request *DescribeApiAnalyzeStatusRequest) (response *DescribeApiAnalyzeStatusResponse, err error) {
	if request == nil {
		request = NewDescribeApiAnalyzeStatusRequest()
	}
	response = NewDescribeApiAnalyzeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAddCustomWhiteRuleRequest() (request *AddCustomWhiteRuleRequest) {
	request = &AddCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddCustomWhiteRule")
	return
}

func NewAddCustomWhiteRuleResponse() (response *AddCustomWhiteRuleResponse) {
	response = &AddCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加精准白名单规则
func (c *Client) AddCustomWhiteRule(request *AddCustomWhiteRuleRequest) (response *AddCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddCustomWhiteRuleRequest()
	}
	response = NewAddCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAntiInfoLeakRulesRequest() (request *DescribeAntiInfoLeakRulesRequest) {
	request = &DescribeAntiInfoLeakRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiInfoLeakRules")
	return
}

func NewDescribeAntiInfoLeakRulesResponse() (response *DescribeAntiInfoLeakRulesResponse) {
	response = &DescribeAntiInfoLeakRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取信息防泄漏规则列表
func (c *Client) DescribeAntiInfoLeakRules(request *DescribeAntiInfoLeakRulesRequest) (response *DescribeAntiInfoLeakRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAntiInfoLeakRulesRequest()
	}
	response = NewDescribeAntiInfoLeakRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotOverviewRequest() (request *DescribeBotOverviewRequest) {
	request = &DescribeBotOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotOverview")
	return
}

func NewDescribeBotOverviewResponse() (response *DescribeBotOverviewResponse) {
	response = &DescribeBotOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT的概览数据
func (c *Client) DescribeBotOverview(request *DescribeBotOverviewRequest) (response *DescribeBotOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBotOverviewRequest()
	}
	response = NewDescribeBotOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiAnalyzeStatusRequest() (request *ModifyApiAnalyzeStatusRequest) {
	request = &ModifyApiAnalyzeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyApiAnalyzeStatus")
	return
}

func NewModifyApiAnalyzeStatusResponse() (response *ModifyApiAnalyzeStatusResponse) {
	response = &ModifyApiAnalyzeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api分析页面开关
func (c *Client) ModifyApiAnalyzeStatus(request *ModifyApiAnalyzeStatusRequest) (response *ModifyApiAnalyzeStatusResponse, err error) {
	if request == nil {
		request = NewModifyApiAnalyzeStatusRequest()
	}
	response = NewModifyApiAnalyzeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomRuleRequest() (request *ModifyCustomRuleRequest) {
	request = &ModifyCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomRule")
	return
}

func NewModifyCustomRuleResponse() (response *ModifyCustomRuleResponse) {
	response = &ModifyCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑自定义规则
func (c *Client) ModifyCustomRule(request *ModifyCustomRuleRequest) (response *ModifyCustomRuleResponse, err error) {
	if request == nil {
		request = NewModifyCustomRuleRequest()
	}
	response = NewModifyCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCCRuleListRequest() (request *DescribeCCRuleListRequest) {
	request = &DescribeCCRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCCRuleList")
	return
}

func NewDescribeCCRuleListResponse() (response *DescribeCCRuleListResponse) {
	response = &DescribeCCRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据多条件查询CC规则
func (c *Client) DescribeCCRuleList(request *DescribeCCRuleListRequest) (response *DescribeCCRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeCCRuleListRequest()
	}
	response = NewDescribeCCRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackDownloadRecordsRequest() (request *GetAttackDownloadRecordsRequest) {
	request = &GetAttackDownloadRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetAttackDownloadRecords")
	return
}

func NewGetAttackDownloadRecordsResponse() (response *GetAttackDownloadRecordsResponse) {
	response = &GetAttackDownloadRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询下载攻击日志任务记录列表
func (c *Client) GetAttackDownloadRecords(request *GetAttackDownloadRecordsRequest) (response *GetAttackDownloadRecordsResponse, err error) {
	if request == nil {
		request = NewGetAttackDownloadRecordsRequest()
	}
	response = NewGetAttackDownloadRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopAttackDomainRequest() (request *DescribeTopAttackDomainRequest) {
	request = &DescribeTopAttackDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeTopAttackDomain")
	return
}

func NewDescribeTopAttackDomainResponse() (response *DescribeTopAttackDomainResponse) {
	response = &DescribeTopAttackDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Top5的攻击域名
func (c *Client) DescribeTopAttackDomain(request *DescribeTopAttackDomainRequest) (response *DescribeTopAttackDomainResponse, err error) {
	if request == nil {
		request = NewDescribeTopAttackDomainRequest()
	}
	response = NewDescribeTopAttackDomainResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceRenewFlagRequest() (request *ModifyInstanceRenewFlagRequest) {
	request = &ModifyInstanceRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceRenewFlag")
	return
}

func NewModifyInstanceRenewFlagResponse() (response *ModifyInstanceRenewFlagResponse) {
	response = &ModifyInstanceRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例的自动续费开关
func (c *Client) ModifyInstanceRenewFlag(request *ModifyInstanceRenewFlagRequest) (response *ModifyInstanceRenewFlagResponse, err error) {
	if request == nil {
		request = NewModifyInstanceRenewFlagRequest()
	}
	response = NewModifyInstanceRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonitorInstancesRequest() (request *DescribeMonitorInstancesRequest) {
	request = &DescribeMonitorInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeMonitorInstances")
	return
}

func NewDescribeMonitorInstancesResponse() (response *DescribeMonitorInstancesResponse) {
	response = &DescribeMonitorInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云监控的实例
func (c *Client) DescribeMonitorInstances(request *DescribeMonitorInstancesRequest) (response *DescribeMonitorInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeMonitorInstancesRequest()
	}
	response = NewDescribeMonitorInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotStatPointsRequest() (request *DescribeBotStatPointsRequest) {
	request = &DescribeBotStatPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotStatPoints")
	return
}

func NewDescribeBotStatPointsResponse() (response *DescribeBotStatPointsResponse) {
	response = &DescribeBotStatPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot的趋势图接口
func (c *Client) DescribeBotStatPoints(request *DescribeBotStatPointsRequest) (response *DescribeBotStatPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotStatPointsRequest()
	}
	response = NewDescribeBotStatPointsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomWhiteRuleRequest() (request *ModifyCustomWhiteRuleRequest) {
	request = &ModifyCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomWhiteRule")
	return
}

func NewModifyCustomWhiteRuleResponse() (response *ModifyCustomWhiteRuleResponse) {
	response = &ModifyCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑精准白名单
func (c *Client) ModifyCustomWhiteRule(request *ModifyCustomWhiteRuleRequest) (response *ModifyCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyCustomWhiteRuleRequest()
	}
	response = NewModifyCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetInstanceQpsLimitRequest() (request *GetInstanceQpsLimitRequest) {
	request = &GetInstanceQpsLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "GetInstanceQpsLimit")
	return
}

func NewGetInstanceQpsLimitResponse() (response *GetInstanceQpsLimitResponse) {
	response = &GetInstanceQpsLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取套餐实例的弹性qps上限
func (c *Client) GetInstanceQpsLimit(request *GetInstanceQpsLimitRequest) (response *GetInstanceQpsLimitResponse, err error) {
	if request == nil {
		request = NewGetInstanceQpsLimitRequest()
	}
	response = NewGetInstanceQpsLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecEventOverviewRequest() (request *DescribeApiSecEventOverviewRequest) {
	request = &DescribeApiSecEventOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiSecEventOverview")
	return
}

func NewDescribeApiSecEventOverviewResponse() (response *DescribeApiSecEventOverviewResponse) {
	response = &DescribeApiSecEventOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// API安全事件概览
func (c *Client) DescribeApiSecEventOverview(request *DescribeApiSecEventOverviewRequest) (response *DescribeApiSecEventOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecEventOverviewRequest()
	}
	response = NewDescribeApiSecEventOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRulesInfoRequest() (request *DescribeRulesInfoRequest) {
	request = &DescribeRulesInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeRulesInfo")
	return
}

func NewDescribeRulesInfoResponse() (response *DescribeRulesInfoResponse) {
	response = &DescribeRulesInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取规则库的描述信息
func (c *Client) DescribeRulesInfo(request *DescribeRulesInfoRequest) (response *DescribeRulesInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRulesInfoRequest()
	}
	response = NewDescribeRulesInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainWhiteRuleRequest() (request *ModifyDomainWhiteRuleRequest) {
	request = &ModifyDomainWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainWhiteRule")
	return
}

func NewModifyDomainWhiteRuleResponse() (response *ModifyDomainWhiteRuleResponse) {
	response = &ModifyDomainWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改某一条规则
func (c *Client) ModifyDomainWhiteRule(request *ModifyDomainWhiteRuleRequest) (response *ModifyDomainWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyDomainWhiteRuleRequest()
	}
	response = NewModifyDomainWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAIModelStatusRequest() (request *DescribeAIModelStatusRequest) {
	request = &DescribeAIModelStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAIModelStatus")
	return
}

func NewDescribeAIModelStatusResponse() (response *DescribeAIModelStatusResponse) {
	response = &DescribeAIModelStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ai模型状态
func (c *Client) DescribeAIModelStatus(request *DescribeAIModelStatusRequest) (response *DescribeAIModelStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAIModelStatusRequest()
	}
	response = NewDescribeAIModelStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBusinessRiskRequest() (request *CreateBusinessRiskRequest) {
	request = &CreateBusinessRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreateBusinessRisk")
	return
}

func NewCreateBusinessRiskResponse() (response *CreateBusinessRiskResponse) {
	response = &CreateBusinessRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建业务安全规则
func (c *Client) CreateBusinessRisk(request *CreateBusinessRiskRequest) (response *CreateBusinessRiskResponse, err error) {
	if request == nil {
		request = NewCreateBusinessRiskRequest()
	}
	response = NewCreateBusinessRiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeObjectsRequest() (request *DescribeObjectsRequest) {
	request = &DescribeObjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeObjects")
	return
}

func NewDescribeObjectsResponse() (response *DescribeObjectsResponse) {
	response = &DescribeObjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看防护对象列表
func (c *Client) DescribeObjects(request *DescribeObjectsRequest) (response *DescribeObjectsResponse, err error) {
	if request == nil {
		request = NewDescribeObjectsRequest()
	}
	response = NewDescribeObjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomRuleRequest() (request *DeleteCustomRuleRequest) {
	request = &DeleteCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteCustomRule")
	return
}

func NewDeleteCustomRuleResponse() (response *DeleteCustomRuleResponse) {
	response = &DeleteCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义规则
func (c *Client) DeleteCustomRule(request *DeleteCustomRuleRequest) (response *DeleteCustomRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCustomRuleRequest()
	}
	response = NewDeleteCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAggregateDomainStatRequest() (request *DescribeBotAggregateDomainStatRequest) {
	request = &DescribeBotAggregateDomainStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotAggregateDomainStat")
	return
}

func NewDescribeBotAggregateDomainStatResponse() (response *DescribeBotAggregateDomainStatResponse) {
	response = &DescribeBotAggregateDomainStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 域名bot统计
func (c *Client) DescribeBotAggregateDomainStat(request *DescribeBotAggregateDomainStatRequest) (response *DescribeBotAggregateDomainStatResponse, err error) {
	if request == nil {
		request = NewDescribeBotAggregateDomainStatRequest()
	}
	response = NewDescribeBotAggregateDomainStatResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePostCKafkaFlowRequest() (request *CreatePostCKafkaFlowRequest) {
	request = &CreatePostCKafkaFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreatePostCKafkaFlow")
	return
}

func NewCreatePostCKafkaFlowResponse() (response *CreatePostCKafkaFlowResponse) {
	response = &CreatePostCKafkaFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建CKafka投递流任务
func (c *Client) CreatePostCKafkaFlow(request *CreatePostCKafkaFlowRequest) (response *CreatePostCKafkaFlowResponse, err error) {
	if request == nil {
		request = NewCreatePostCKafkaFlowRequest()
	}
	response = NewCreatePostCKafkaFlowResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCCRuleRequest() (request *DeleteCCRuleRequest) {
	request = &DeleteCCRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteCCRule")
	return
}

func NewDeleteCCRuleResponse() (response *DeleteCCRuleResponse) {
	response = &DeleteCCRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  CC V2 Delete接口
func (c *Client) DeleteCCRule(request *DeleteCCRuleRequest) (response *DeleteCCRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCCRuleRequest()
	}
	response = NewDeleteCCRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCdcResourceRequest() (request *DeleteCdcResourceRequest) {
	request = &DeleteCdcResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteCdcResource")
	return
}

func NewDeleteCdcResourceResponse() (response *DeleteCdcResourceResponse) {
	response = &DeleteCdcResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除cdc资源数据
func (c *Client) DeleteCdcResource(request *DeleteCdcResourceRequest) (response *DeleteCdcResourceResponse, err error) {
	if request == nil {
		request = NewDeleteCdcResourceRequest()
	}
	response = NewDeleteCdcResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainBlockPageRequest() (request *DescribeDomainBlockPageRequest) {
	request = &DescribeDomainBlockPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainBlockPage")
	return
}

func NewDescribeDomainBlockPageResponse() (response *DescribeDomainBlockPageResponse) {
	response = &DescribeDomainBlockPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示指定域名的拦截页面信息
func (c *Client) DescribeDomainBlockPage(request *DescribeDomainBlockPageRequest) (response *DescribeDomainBlockPageResponse, err error) {
	if request == nil {
		request = NewDescribeDomainBlockPageRequest()
	}
	response = NewDescribeDomainBlockPageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostLimitRequest() (request *DescribeHostLimitRequest) {
	request = &DescribeHostLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeHostLimit")
	return
}

func NewDescribeHostLimitResponse() (response *DescribeHostLimitResponse) {
	response = &DescribeHostLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加域名的首先验证是否购买了套餐，是否没有达到购买套餐的限制，域名是否已经添加
func (c *Client) DescribeHostLimit(request *DescribeHostLimitRequest) (response *DescribeHostLimitResponse, err error) {
	if request == nil {
		request = NewDescribeHostLimitRequest()
	}
	response = NewDescribeHostLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserInstancesRequest() (request *DescribeUserInstancesRequest) {
	request = &DescribeUserInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserInstances")
	return
}

func NewDescribeUserInstancesResponse() (response *DescribeUserInstancesResponse) {
	response = &DescribeUserInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户实例展示，展示通过地域参数展示特定地域用户购买的所有实例
func (c *Client) DescribeUserInstances(request *DescribeUserInstancesRequest) (response *DescribeUserInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeUserInstancesRequest()
	}
	response = NewDescribeUserInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessExportsRequest() (request *DescribeAccessExportsRequest) {
	request = &DescribeAccessExportsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessExports")
	return
}

func NewDescribeAccessExportsResponse() (response *DescribeAccessExportsResponse) {
	response = &DescribeAccessExportsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取访问日志导出列表
func (c *Client) DescribeAccessExports(request *DescribeAccessExportsRequest) (response *DescribeAccessExportsResponse, err error) {
	if request == nil {
		request = NewDescribeAccessExportsRequest()
	}
	response = NewDescribeAccessExportsResponse()
	err = c.Send(request, response)
	return
}

func NewCopyBotUCBFeatureRulesRequest() (request *CopyBotUCBFeatureRulesRequest) {
	request = &CopyBotUCBFeatureRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CopyBotUCBFeatureRules")
	return
}

func NewCopyBotUCBFeatureRulesResponse() (response *CopyBotUCBFeatureRulesResponse) {
	response = &CopyBotUCBFeatureRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB自定义策略拷贝
func (c *Client) CopyBotUCBFeatureRules(request *CopyBotUCBFeatureRulesRequest) (response *CopyBotUCBFeatureRulesResponse, err error) {
	if request == nil {
		request = NewCopyBotUCBFeatureRulesRequest()
	}
	response = NewCopyBotUCBFeatureRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePriceRequest() (request *DescribePriceRequest) {
	request = &DescribePriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePrice")
	return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给clb等使用的waf实例询价接口，目前只支持clb旗舰版实例的下单，该接口会进行入参校验，然后调用是否为收购用户，然后调用计费询价接口。
func (c *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
	if request == nil {
		request = NewDescribePriceRequest()
	}
	response = NewDescribePriceResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertIpAccessControlRequest() (request *UpsertIpAccessControlRequest) {
	request = &UpsertIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "UpsertIpAccessControl")
	return
}

func NewUpsertIpAccessControlResponse() (response *UpsertIpAccessControlResponse) {
	response = &UpsertIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf IP黑白名单Upsert接口
func (c *Client) UpsertIpAccessControl(request *UpsertIpAccessControlRequest) (response *UpsertIpAccessControlResponse, err error) {
	if request == nil {
		request = NewUpsertIpAccessControlRequest()
	}
	response = NewUpsertIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSpartaProtectionRequest() (request *DeleteSpartaProtectionRequest) {
	request = &DeleteSpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteSpartaProtection")
	return
}

func NewDeleteSpartaProtectionResponse() (response *DeleteSpartaProtectionResponse) {
	response = &DeleteSpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-删除防护域名
func (c *Client) DeleteSpartaProtection(request *DeleteSpartaProtectionRequest) (response *DeleteSpartaProtectionResponse, err error) {
	if request == nil {
		request = NewDeleteSpartaProtectionRequest()
	}
	response = NewDeleteSpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWhiteListRequest() (request *ModifyWhiteListRequest) {
	request = &ModifyWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyWhiteList")
	return
}

func NewModifyWhiteListResponse() (response *ModifyWhiteListResponse) {
	response = &ModifyWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新业务白名单，目前是前端对抗中的白名单
func (c *Client) ModifyWhiteList(request *ModifyWhiteListRequest) (response *ModifyWhiteListResponse, err error) {
	if request == nil {
		request = NewModifyWhiteListRequest()
	}
	response = NewModifyWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiAggregateTopNRequest() (request *DescribeApiAggregateTopNRequest) {
	request = &DescribeApiAggregateTopNRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiAggregateTopN")
	return
}

func NewDescribeApiAggregateTopNResponse() (response *DescribeApiAggregateTopNResponse) {
	response = &DescribeApiAggregateTopNResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Api安全模块的访问日志聚合topN
func (c *Client) DescribeApiAggregateTopN(request *DescribeApiAggregateTopNRequest) (response *DescribeApiAggregateTopNResponse, err error) {
	if request == nil {
		request = NewDescribeApiAggregateTopNRequest()
	}
	response = NewDescribeApiAggregateTopNResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCertificateVerifyResultRequest() (request *DescribeCertificateVerifyResultRequest) {
	request = &DescribeCertificateVerifyResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCertificateVerifyResult")
	return
}

func NewDescribeCertificateVerifyResultResponse() (response *DescribeCertificateVerifyResultResponse) {
	response = &DescribeCertificateVerifyResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取证书的检查结果
func (c *Client) DescribeCertificateVerifyResult(request *DescribeCertificateVerifyResultRequest) (response *DescribeCertificateVerifyResultResponse, err error) {
	if request == nil {
		request = NewDescribeCertificateVerifyResultRequest()
	}
	response = NewDescribeCertificateVerifyResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafAutoDenyStatusRequest() (request *DescribeWafAutoDenyStatusRequest) {
	request = &DescribeWafAutoDenyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeWafAutoDenyStatus")
	return
}

func NewDescribeWafAutoDenyStatusResponse() (response *DescribeWafAutoDenyStatusResponse) {
	response = &DescribeWafAutoDenyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述WAF自动封禁模块详情
func (c *Client) DescribeWafAutoDenyStatus(request *DescribeWafAutoDenyStatusRequest) (response *DescribeWafAutoDenyStatusResponse, err error) {
	if request == nil {
		request = NewDescribeWafAutoDenyStatusRequest()
	}
	response = NewDescribeWafAutoDenyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAIVerifyResultRequest() (request *DescribeAIVerifyResultRequest) {
	request = &DescribeAIVerifyResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAIVerifyResult")
	return
}

func NewDescribeAIVerifyResultResponse() (response *DescribeAIVerifyResultResponse) {
	response = &DescribeAIVerifyResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf ai在线验证
func (c *Client) DescribeAIVerifyResult(request *DescribeAIVerifyResultRequest) (response *DescribeAIVerifyResultResponse, err error) {
	if request == nil {
		request = NewDescribeAIVerifyResultRequest()
	}
	response = NewDescribeAIVerifyResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePieChartRequest() (request *DescribePieChartRequest) {
	request = &DescribePieChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePieChart")
	return
}

func NewDescribePieChartResponse() (response *DescribePieChartResponse) {
	response = &DescribePieChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定聚类条件的饼图
func (c *Client) DescribePieChart(request *DescribePieChartRequest) (response *DescribePieChartResponse, err error) {
	if request == nil {
		request = NewDescribePieChartRequest()
	}
	response = NewDescribePieChartResponse()
	err = c.Send(request, response)
	return
}

func NewSearchAccessLogRequest() (request *SearchAccessLogRequest) {
	request = &SearchAccessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "SearchAccessLog")
	return
}

func NewSearchAccessLogResponse() (response *SearchAccessLogResponse) {
	response = &SearchAccessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于搜索WAF访问日志
func (c *Client) SearchAccessLog(request *SearchAccessLogRequest) (response *SearchAccessLogResponse, err error) {
	if request == nil {
		request = NewSearchAccessLogRequest()
	}
	response = NewSearchAccessLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceElasticModeRequest() (request *ModifyInstanceElasticModeRequest) {
	request = &ModifyInstanceElasticModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceElasticMode")
	return
}

func NewModifyInstanceElasticModeResponse() (response *ModifyInstanceElasticModeResponse) {
	response = &ModifyInstanceElasticModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例的QPS弹性计费开关
func (c *Client) ModifyInstanceElasticMode(request *ModifyInstanceElasticModeRequest) (response *ModifyInstanceElasticModeResponse, err error) {
	if request == nil {
		request = NewModifyInstanceElasticModeRequest()
	}
	response = NewModifyInstanceElasticModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotCountRequest() (request *DescribeBotCountRequest) {
	request = &DescribeBotCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotCount")
	return
}

func NewDescribeBotCountResponse() (response *DescribeBotCountResponse) {
	response = &DescribeBotCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定域名
func (c *Client) DescribeBotCount(request *DescribeBotCountRequest) (response *DescribeBotCountResponse, err error) {
	if request == nil {
		request = NewDescribeBotCountRequest()
	}
	response = NewDescribeBotCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainRulesRequest() (request *DescribeDomainRulesRequest) {
	request = &DescribeDomainRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainRules")
	return
}

func NewDescribeDomainRulesResponse() (response *DescribeDomainRulesResponse) {
	response = &DescribeDomainRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取域名的防护规则列表
func (c *Client) DescribeDomainRules(request *DescribeDomainRulesRequest) (response *DescribeDomainRulesResponse, err error) {
	if request == nil {
		request = NewDescribeDomainRulesRequest()
	}
	response = NewDescribeDomainRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSceneActionRuleRequest() (request *DescribeBotSceneActionRuleRequest) {
	request = &DescribeBotSceneActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotSceneActionRule")
	return
}

func NewDescribeBotSceneActionRuleResponse() (response *DescribeBotSceneActionRuleResponse) {
	response = &DescribeBotSceneActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT场景的动作策略列表
func (c *Client) DescribeBotSceneActionRule(request *DescribeBotSceneActionRuleRequest) (response *DescribeBotSceneActionRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotSceneActionRuleRequest()
	}
	response = NewDescribeBotSceneActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserLevelRequest() (request *ModifyUserLevelRequest) {
	request = &ModifyUserLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyUserLevel")
	return
}

func NewModifyUserLevelResponse() (response *ModifyUserLevelResponse) {
	response = &ModifyUserLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户防护规则等级
func (c *Client) ModifyUserLevel(request *ModifyUserLevelRequest) (response *ModifyUserLevelResponse, err error) {
	if request == nil {
		request = NewModifyUserLevelRequest()
	}
	response = NewModifyUserLevelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAntiFakeUrlRequest() (request *DescribeAntiFakeUrlRequest) {
	request = &DescribeAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiFakeUrl")
	return
}

func NewDescribeAntiFakeUrlResponse() (response *DescribeAntiFakeUrlResponse) {
	response = &DescribeAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防篡改url
func (c *Client) DescribeAntiFakeUrl(request *DescribeAntiFakeUrlRequest) (response *DescribeAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewDescribeAntiFakeUrlRequest()
	}
	response = NewDescribeAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiListVersionTwoRequest() (request *DescribeApiListVersionTwoRequest) {
	request = &DescribeApiListVersionTwoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiListVersionTwo")
	return
}

func NewDescribeApiListVersionTwoResponse() (response *DescribeApiListVersionTwoResponse) {
	response = &DescribeApiListVersionTwoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api资产列表
func (c *Client) DescribeApiListVersionTwo(request *DescribeApiListVersionTwoRequest) (response *DescribeApiListVersionTwoResponse, err error) {
	if request == nil {
		request = NewDescribeApiListVersionTwoRequest()
	}
	response = NewDescribeApiListVersionTwoResponse()
	err = c.Send(request, response)
	return
}

func NewAddCustomRuleRequest() (request *AddCustomRuleRequest) {
	request = &AddCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddCustomRule")
	return
}

func NewAddCustomRuleResponse() (response *AddCustomRuleResponse) {
	response = &AddCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加访问控制（自定义策略）
func (c *Client) AddCustomRule(request *AddCustomRuleRequest) (response *AddCustomRuleResponse, err error) {
	if request == nil {
		request = NewAddCustomRuleRequest()
	}
	response = NewAddCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSceneUCBRuleRequest() (request *DescribeBotSceneUCBRuleRequest) {
	request = &DescribeBotSceneUCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotSceneUCBRule")
	return
}

func NewDescribeBotSceneUCBRuleResponse() (response *DescribeBotSceneUCBRuleResponse) {
	response = &DescribeBotSceneUCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 场景化后Bot获取UCB自定义规则策略
func (c *Client) DescribeBotSceneUCBRule(request *DescribeBotSceneUCBRuleRequest) (response *DescribeBotSceneUCBRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotSceneUCBRuleRequest()
	}
	response = NewDescribeBotSceneUCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePeakValueRequest() (request *DescribePeakValueRequest) {
	request = &DescribePeakValueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePeakValue")
	return
}

func NewDescribePeakValueResponse() (response *DescribePeakValueResponse) {
	response = &DescribePeakValueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取业务和攻击概览峰值
func (c *Client) DescribePeakValue(request *DescribePeakValueRequest) (response *DescribePeakValueResponse, err error) {
	if request == nil {
		request = NewDescribePeakValueRequest()
	}
	response = NewDescribePeakValueResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostCLSFlowsRequest() (request *DescribePostCLSFlowsRequest) {
	request = &DescribePostCLSFlowsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePostCLSFlows")
	return
}

func NewDescribePostCLSFlowsResponse() (response *DescribePostCLSFlowsResponse) {
	response = &DescribePostCLSFlowsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CLS投递流任务列表
func (c *Client) DescribePostCLSFlows(request *DescribePostCLSFlowsRequest) (response *DescribePostCLSFlowsResponse, err error) {
	if request == nil {
		request = NewDescribePostCLSFlowsRequest()
	}
	response = NewDescribePostCLSFlowsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpartaProtectionInfoRequest() (request *DescribeSpartaProtectionInfoRequest) {
	request = &DescribeSpartaProtectionInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeSpartaProtectionInfo")
	return
}

func NewDescribeSpartaProtectionInfoResponse() (response *DescribeSpartaProtectionInfoResponse) {
	response = &DescribeSpartaProtectionInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-获取防护域名信息
func (c *Client) DescribeSpartaProtectionInfo(request *DescribeSpartaProtectionInfoRequest) (response *DescribeSpartaProtectionInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSpartaProtectionInfoRequest()
	}
	response = NewDescribeSpartaProtectionInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIpAccessControlRequest() (request *DeleteIpAccessControlRequest) {
	request = &DeleteIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteIpAccessControl")
	return
}

func NewDeleteIpAccessControlResponse() (response *DeleteIpAccessControlResponse) {
	response = &DeleteIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf IP黑白名单Delete接口
func (c *Client) DeleteIpAccessControl(request *DeleteIpAccessControlRequest) (response *DeleteIpAccessControlResponse, err error) {
	if request == nil {
		request = NewDeleteIpAccessControlRequest()
	}
	response = NewDeleteIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAggregateTopNRequest() (request *DescribeBotAggregateTopNRequest) {
	request = &DescribeBotAggregateTopNRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotAggregateTopN")
	return
}

func NewDescribeBotAggregateTopNResponse() (response *DescribeBotAggregateTopNResponse) {
	response = &DescribeBotAggregateTopNResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT的日志聚合topN
func (c *Client) DescribeBotAggregateTopN(request *DescribeBotAggregateTopNRequest) (response *DescribeBotAggregateTopNResponse, err error) {
	if request == nil {
		request = NewDescribeBotAggregateTopNRequest()
	}
	response = NewDescribeBotAggregateTopNResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCLSPackageRequest() (request *DescribeCLSPackageRequest) {
	request = &DescribeCLSPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeCLSPackage")
	return
}

func NewDescribeCLSPackageResponse() (response *DescribeCLSPackageResponse) {
	response = &DescribeCLSPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的CLS信息
func (c *Client) DescribeCLSPackage(request *DescribeCLSPackageRequest) (response *DescribeCLSPackageResponse, err error) {
	if request == nil {
		request = NewDescribeCLSPackageRequest()
	}
	response = NewDescribeCLSPackageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleTypesRequest() (request *DescribeRuleTypesRequest) {
	request = &DescribeRuleTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeRuleTypes")
	return
}

func NewDescribeRuleTypesResponse() (response *DescribeRuleTypesResponse) {
	response = &DescribeRuleTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取规则类型描述
func (c *Client) DescribeRuleTypes(request *DescribeRuleTypesRequest) (response *DescribeRuleTypesResponse, err error) {
	if request == nil {
		request = NewDescribeRuleTypesRequest()
	}
	response = NewDescribeRuleTypesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceQpsLimitRequest() (request *ModifyInstanceQpsLimitRequest) {
	request = &ModifyInstanceQpsLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceQpsLimit")
	return
}

func NewModifyInstanceQpsLimitResponse() (response *ModifyInstanceQpsLimitResponse) {
	response = &ModifyInstanceQpsLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置套餐实例的弹性qps上限
func (c *Client) ModifyInstanceQpsLimit(request *ModifyInstanceQpsLimitRequest) (response *ModifyInstanceQpsLimitResponse, err error) {
	if request == nil {
		request = NewModifyInstanceQpsLimitRequest()
	}
	response = NewModifyInstanceQpsLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotActionStatRequest() (request *DescribeBotActionStatRequest) {
	request = &DescribeBotActionStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotActionStat")
	return
}

func NewDescribeBotActionStatResponse() (response *DescribeBotActionStatResponse) {
	response = &DescribeBotActionStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取bot动作统计
func (c *Client) DescribeBotActionStat(request *DescribeBotActionStatRequest) (response *DescribeBotActionStatResponse, err error) {
	if request == nil {
		request = NewDescribeBotActionStatRequest()
	}
	response = NewDescribeBotActionStatResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProtectionStatusRequest() (request *ModifyProtectionStatusRequest) {
	request = &ModifyProtectionStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyProtectionStatus")
	return
}

func NewModifyProtectionStatusResponse() (response *ModifyProtectionStatusResponse) {
	response = &ModifyProtectionStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-waf开关
func (c *Client) ModifyProtectionStatus(request *ModifyProtectionStatusRequest) (response *ModifyProtectionStatusResponse, err error) {
	if request == nil {
		request = NewModifyProtectionStatusRequest()
	}
	response = NewModifyProtectionStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotLevelRequest() (request *DescribeBotLevelRequest) {
	request = &DescribeBotLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotLevel")
	return
}

func NewDescribeBotLevelResponse() (response *DescribeBotLevelResponse) {
	response = &DescribeBotLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot等级配置
func (c *Client) DescribeBotLevel(request *DescribeBotLevelRequest) (response *DescribeBotLevelResponse, err error) {
	if request == nil {
		request = NewDescribeBotLevelRequest()
	}
	response = NewDescribeBotLevelResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpartaPackageRenewRequest() (request *ModifySpartaPackageRenewRequest) {
	request = &ModifySpartaPackageRenewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifySpartaPackageRenew")
	return
}

func NewModifySpartaPackageRenewResponse() (response *ModifySpartaPackageRenewResponse) {
	response = &ModifySpartaPackageRenewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Sparta版本WAF自动续费开关设置
func (c *Client) ModifySpartaPackageRenew(request *ModifySpartaPackageRenewRequest) (response *ModifySpartaPackageRenewResponse, err error) {
	if request == nil {
		request = NewModifySpartaPackageRenewRequest()
	}
	response = NewModifySpartaPackageRenewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVIPRSRequest() (request *ModifyVIPRSRequest) {
	request = &ModifyVIPRSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyVIPRS")
	return
}

func NewModifyVIPRSResponse() (response *ModifyVIPRSResponse) {
	response = &ModifyVIPRSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 迁移流量
func (c *Client) ModifyVIPRS(request *ModifyVIPRSRequest) (response *ModifyVIPRSResponse, err error) {
	if request == nil {
		request = NewModifyVIPRSRequest()
	}
	response = NewModifyVIPRSResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWafThreatenIntelligenceRequest() (request *ModifyWafThreatenIntelligenceRequest) {
	request = &ModifyWafThreatenIntelligenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyWafThreatenIntelligence")
	return
}

func NewModifyWafThreatenIntelligenceResponse() (response *ModifyWafThreatenIntelligenceResponse) {
	response = &ModifyWafThreatenIntelligenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置WAF威胁情报封禁模块详情
func (c *Client) ModifyWafThreatenIntelligence(request *ModifyWafThreatenIntelligenceRequest) (response *ModifyWafThreatenIntelligenceResponse, err error) {
	if request == nil {
		request = NewModifyWafThreatenIntelligenceRequest()
	}
	response = NewModifyWafThreatenIntelligenceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePortsRequest() (request *DescribePortsRequest) {
	request = &DescribePortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePorts")
	return
}

func NewDescribePortsResponse() (response *DescribePortsResponse) {
	response = &DescribePortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取非标端口列表
func (c *Client) DescribePorts(request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
	if request == nil {
		request = NewDescribePortsRequest()
	}
	response = NewDescribePortsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserGradeRequest() (request *DescribeUserGradeRequest) {
	request = &DescribeUserGradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserGrade")
	return
}

func NewDescribeUserGradeResponse() (response *DescribeUserGradeResponse) {
	response = &DescribeUserGradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据用户Uin查询用户的等级和类型
func (c *Client) DescribeUserGrade(request *DescribeUserGradeRequest) (response *DescribeUserGradeResponse, err error) {
	if request == nil {
		request = NewDescribeUserGradeRequest()
	}
	response = NewDescribeUserGradeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafInfoRequest() (request *DescribeWafInfoRequest) {
	request = &DescribeWafInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeWafInfo")
	return
}

func NewDescribeWafInfoResponse() (response *DescribeWafInfoResponse) {
	response = &DescribeWafInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取负载均衡绑定的WAF信息，可以根据租户负载均衡实例ID、负载均衡监听器ID、负载均衡的域名信息来查询对应绑定的 Waf的状态信息。
// 查询的范围：负载均衡实例ID、负载均衡实例ID+监听器ID、负载均衡实例ID+监听器ID+域名。
// 可能的错误码：ResourceNotFound（没有找到对应的资源）、UnsupportedRegion（目前clb-waf只支持北京、广州、上海、成都、重庆、香港地域）。
func (c *Client) DescribeWafInfo(request *DescribeWafInfoRequest) (response *DescribeWafInfoResponse, err error) {
	if request == nil {
		request = NewDescribeWafInfoRequest()
	}
	response = NewDescribeWafInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescirbeCCRuleRequest() (request *DescirbeCCRuleRequest) {
	request = &DescirbeCCRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescirbeCCRule")
	return
}

func NewDescirbeCCRuleResponse() (response *DescirbeCCRuleResponse) {
	response = &DescirbeCCRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  CC V2 Query接口
func (c *Client) DescirbeCCRule(request *DescirbeCCRuleRequest) (response *DescirbeCCRuleResponse, err error) {
	if request == nil {
		request = NewDescirbeCCRuleRequest()
	}
	response = NewDescirbeCCRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPublishOrCancelRuleRequest() (request *ModifyPublishOrCancelRuleRequest) {
	request = &ModifyPublishOrCancelRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyPublishOrCancelRule")
	return
}

func NewModifyPublishOrCancelRuleResponse() (response *ModifyPublishOrCancelRuleResponse) {
	response = &ModifyPublishOrCancelRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改业务安全规则上线or下线
func (c *Client) ModifyPublishOrCancelRule(request *ModifyPublishOrCancelRuleRequest) (response *ModifyPublishOrCancelRuleResponse, err error) {
	if request == nil {
		request = NewModifyPublishOrCancelRuleRequest()
	}
	response = NewModifyPublishOrCancelRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFirstPurchaseRequest() (request *DescribeFirstPurchaseRequest) {
	request = &DescribeFirstPurchaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeFirstPurchase")
	return
}

func NewDescribeFirstPurchaseResponse() (response *DescribeFirstPurchaseResponse) {
	response = &DescribeFirstPurchaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户是否为第一次购买waf实例，api接口会调用计费侧查询订单接口，如果订单数量totalNum大于0，则认为不是首购用户，计费侧根据ownerUin查询订单
func (c *Client) DescribeFirstPurchase(request *DescribeFirstPurchaseRequest) (response *DescribeFirstPurchaseResponse, err error) {
	if request == nil {
		request = NewDescribeFirstPurchaseRequest()
	}
	response = NewDescribeFirstPurchaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserInfoRequest() (request *DescribeUserInfoRequest) {
	request = &DescribeUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserInfo")
	return
}

func NewDescribeUserInfoResponse() (response *DescribeUserInfoResponse) {
	response = &DescribeUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取clb-waf的套餐信息
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserInfoRequest()
	}
	response = NewDescribeUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewAddCustomPayloadRequest() (request *AddCustomPayloadRequest) {
	request = &AddCustomPayloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "AddCustomPayload")
	return
}

func NewAddCustomPayloadResponse() (response *AddCustomPayloadResponse) {
	response = &AddCustomPayloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加自定义载荷。
func (c *Client) AddCustomPayload(request *AddCustomPayloadRequest) (response *AddCustomPayloadResponse, err error) {
	if request == nil {
		request = NewAddCustomPayloadRequest()
	}
	response = NewAddCustomPayloadResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSceneOverviewRequest() (request *DescribeBotSceneOverviewRequest) {
	request = &DescribeBotSceneOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotSceneOverview")
	return
}

func NewDescribeBotSceneOverviewResponse() (response *DescribeBotSceneOverviewResponse) {
	response = &DescribeBotSceneOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Bot场景全局概览
func (c *Client) DescribeBotSceneOverview(request *DescribeBotSceneOverviewRequest) (response *DescribeBotSceneOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBotSceneOverviewRequest()
	}
	response = NewDescribeBotSceneOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackDownloadRecordsRequest() (request *DescribeAttackDownloadRecordsRequest) {
	request = &DescribeAttackDownloadRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackDownloadRecords")
	return
}

func NewDescribeAttackDownloadRecordsResponse() (response *DescribeAttackDownloadRecordsResponse) {
	response = &DescribeAttackDownloadRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询攻击日志的下载记录
func (c *Client) DescribeAttackDownloadRecords(request *DescribeAttackDownloadRecordsRequest) (response *DescribeAttackDownloadRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeAttackDownloadRecordsRequest()
	}
	response = NewDescribeAttackDownloadRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiListRequest() (request *DescribeApiListRequest) {
	request = &DescribeApiListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiList")
	return
}

func NewDescribeApiListResponse() (response *DescribeApiListResponse) {
	response = &DescribeApiListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示api规则列表信息
func (c *Client) DescribeApiList(request *DescribeApiListRequest) (response *DescribeApiListResponse, err error) {
	if request == nil {
		request = NewDescribeApiListRequest()
	}
	response = NewDescribeApiListResponse()
	err = c.Send(request, response)
	return
}

func NewEditOpRuleUpdateLogRequest() (request *EditOpRuleUpdateLogRequest) {
	request = &EditOpRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "EditOpRuleUpdateLog")
	return
}

func NewEditOpRuleUpdateLogResponse() (response *EditOpRuleUpdateLogResponse) {
	response = &EditOpRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：编辑规则更新日志
func (c *Client) EditOpRuleUpdateLog(request *EditOpRuleUpdateLogRequest) (response *EditOpRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewEditOpRuleUpdateLogRequest()
	}
	response = NewEditOpRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewNoticeInstanceChangeRequest() (request *NoticeInstanceChangeRequest) {
	request = &NoticeInstanceChangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "NoticeInstanceChange")
	return
}

func NewNoticeInstanceChangeResponse() (response *NoticeInstanceChangeResponse) {
	response = &NoticeInstanceChangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对新购实例进行带宽限速，客户在购买、降配业务扩展包/开启关闭弹性计费/弹性计费上限调整等操作后通知接入
func (c *Client) NoticeInstanceChange(request *NoticeInstanceChangeRequest) (response *NoticeInstanceChangeResponse, err error) {
	if request == nil {
		request = NewNoticeInstanceChangeRequest()
	}
	response = NewNoticeInstanceChangeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBusinessRiskRequest() (request *DescribeBusinessRiskRequest) {
	request = &DescribeBusinessRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBusinessRisk")
	return
}

func NewDescribeBusinessRiskResponse() (response *DescribeBusinessRiskResponse) {
	response = &DescribeBusinessRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务安全规则查询
func (c *Client) DescribeBusinessRisk(request *DescribeBusinessRiskRequest) (response *DescribeBusinessRiskResponse, err error) {
	if request == nil {
		request = NewDescribeBusinessRiskRequest()
	}
	response = NewDescribeBusinessRiskResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertCCAutoStatusRequest() (request *UpsertCCAutoStatusRequest) {
	request = &UpsertCCAutoStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "UpsertCCAutoStatus")
	return
}

func NewUpsertCCAutoStatusResponse() (response *UpsertCCAutoStatusResponse) {
	response = &UpsertCCAutoStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 斯巴达版本更新cc自动封堵状态
func (c *Client) UpsertCCAutoStatus(request *UpsertCCAutoStatusRequest) (response *UpsertCCAutoStatusResponse, err error) {
	if request == nil {
		request = NewUpsertCCAutoStatusRequest()
	}
	response = NewUpsertCCAutoStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCopyCustomRuleRequest() (request *CopyCustomRuleRequest) {
	request = &CopyCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CopyCustomRule")
	return
}

func NewCopyCustomRuleResponse() (response *CopyCustomRuleResponse) {
	response = &CopyCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从一个域名复制一条自定义策略到其他域名
func (c *Client) CopyCustomRule(request *CopyCustomRuleRequest) (response *CopyCustomRuleResponse, err error) {
	if request == nil {
		request = NewCopyCustomRuleRequest()
	}
	response = NewCopyCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAccessAggregateTopNRequest() (request *DescribeBotAccessAggregateTopNRequest) {
	request = &DescribeBotAccessAggregateTopNRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotAccessAggregateTopN")
	return
}

func NewDescribeBotAccessAggregateTopNResponse() (response *DescribeBotAccessAggregateTopNResponse) {
	response = &DescribeBotAccessAggregateTopNResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT模块的访问日志聚合topN
func (c *Client) DescribeBotAccessAggregateTopN(request *DescribeBotAccessAggregateTopNRequest) (response *DescribeBotAccessAggregateTopNResponse, err error) {
	if request == nil {
		request = NewDescribeBotAccessAggregateTopNRequest()
	}
	response = NewDescribeBotAccessAggregateTopNResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessLogCountRequest() (request *DescribeAccessLogCountRequest) {
	request = &DescribeAccessLogCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessLogCount")
	return
}

func NewDescribeAccessLogCountResponse() (response *DescribeAccessLogCountResponse) {
	response = &DescribeAccessLogCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询满足搜索条件的访问日志数量
func (c *Client) DescribeAccessLogCount(request *DescribeAccessLogCountRequest) (response *DescribeAccessLogCountResponse, err error) {
	if request == nil {
		request = NewDescribeAccessLogCountRequest()
	}
	response = NewDescribeAccessLogCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRecordPointsRequest() (request *DescribeBotRecordPointsRequest) {
	request = &DescribeBotRecordPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotRecordPoints")
	return
}

func NewDescribeBotRecordPointsResponse() (response *DescribeBotRecordPointsResponse) {
	response = &DescribeBotRecordPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bot记录访问趋势图
func (c *Client) DescribeBotRecordPoints(request *DescribeBotRecordPointsRequest) (response *DescribeBotRecordPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotRecordPointsRequest()
	}
	response = NewDescribeBotRecordPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainDetailsClbRequest() (request *DescribeDomainDetailsClbRequest) {
	request = &DescribeDomainDetailsClbRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainDetailsClb")
	return
}

func NewDescribeDomainDetailsClbResponse() (response *DescribeDomainDetailsClbResponse) {
	response = &DescribeDomainDetailsClbResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取一个clb域名详情
func (c *Client) DescribeDomainDetailsClb(request *DescribeDomainDetailsClbRequest) (response *DescribeDomainDetailsClbResponse, err error) {
	if request == nil {
		request = NewDescribeDomainDetailsClbRequest()
	}
	response = NewDescribeDomainDetailsClbResponse()
	err = c.Send(request, response)
	return
}

func NewExportAccessLogsRequest() (request *ExportAccessLogsRequest) {
	request = &ExportAccessLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ExportAccessLogs")
	return
}

func NewExportAccessLogsResponse() (response *ExportAccessLogsResponse) {
	response = &ExportAccessLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 访问日志批量下载接口，不能带搜索条件
func (c *Client) ExportAccessLogs(request *ExportAccessLogsRequest) (response *ExportAccessLogsResponse, err error) {
	if request == nil {
		request = NewExportAccessLogsRequest()
	}
	response = NewExportAccessLogsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceWafInstanceRequest() (request *InquiryPriceWafInstanceRequest) {
	request = &InquiryPriceWafInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "InquiryPriceWafInstance")
	return
}

func NewInquiryPriceWafInstanceResponse() (response *InquiryPriceWafInstanceResponse) {
	response = &InquiryPriceWafInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询价格
func (c *Client) InquiryPriceWafInstance(request *InquiryPriceWafInstanceRequest) (response *InquiryPriceWafInstanceResponse, err error) {
	if request == nil {
		request = NewInquiryPriceWafInstanceRequest()
	}
	response = NewInquiryPriceWafInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSceneStatusRequest() (request *ModifyBotSceneStatusRequest) {
	request = &ModifyBotSceneStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotSceneStatus")
	return
}

func NewModifyBotSceneStatusResponse() (response *ModifyBotSceneStatusResponse) {
	response = &ModifyBotSceneStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bot子场景开关
func (c *Client) ModifyBotSceneStatus(request *ModifyBotSceneStatusRequest) (response *ModifyBotSceneStatusResponse, err error) {
	if request == nil {
		request = NewModifyBotSceneStatusRequest()
	}
	response = NewModifyBotSceneStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafAutoDenyRulesRequest() (request *DescribeWafAutoDenyRulesRequest) {
	request = &DescribeWafAutoDenyRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeWafAutoDenyRules")
	return
}

func NewDescribeWafAutoDenyRulesResponse() (response *DescribeWafAutoDenyRulesResponse) {
	response = &DescribeWafAutoDenyRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回ip惩罚规则详细信息
func (c *Client) DescribeWafAutoDenyRules(request *DescribeWafAutoDenyRulesRequest) (response *DescribeWafAutoDenyRulesResponse, err error) {
	if request == nil {
		request = NewDescribeWafAutoDenyRulesRequest()
	}
	response = NewDescribeWafAutoDenyRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPackageRenewRequest() (request *ModifyPackageRenewRequest) {
	request = &ModifyPackageRenewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyPackageRenew")
	return
}

func NewModifyPackageRenewResponse() (response *ModifyPackageRenewResponse) {
	response = &ModifyPackageRenewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置套餐自动续费
func (c *Client) ModifyPackageRenew(request *ModifyPackageRenewRequest) (response *ModifyPackageRenewResponse, err error) {
	if request == nil {
		request = NewModifyPackageRenewRequest()
	}
	response = NewModifyPackageRenewResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpartaProtectionRequest() (request *ModifySpartaProtectionRequest) {
	request = &ModifySpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifySpartaProtection")
	return
}

func NewModifySpartaProtectionResponse() (response *ModifySpartaProtectionResponse) {
	response = &ModifySpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名配置
func (c *Client) ModifySpartaProtection(request *ModifySpartaProtectionRequest) (response *ModifySpartaProtectionResponse, err error) {
	if request == nil {
		request = NewModifySpartaProtectionRequest()
	}
	response = NewModifySpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotPointsRequest() (request *DescribeBotPointsRequest) {
	request = &DescribeBotPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotPoints")
	return
}

func NewDescribeBotPointsResponse() (response *DescribeBotPointsResponse) {
	response = &DescribeBotPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Bot趋势数据
func (c *Client) DescribeBotPoints(request *DescribeBotPointsRequest) (response *DescribeBotPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotPointsRequest()
	}
	response = NewDescribeBotPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainSecurityGroupTipsCountRequest() (request *DescribeDomainSecurityGroupTipsCountRequest) {
	request = &DescribeDomainSecurityGroupTipsCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainSecurityGroupTipsCount")
	return
}

func NewDescribeDomainSecurityGroupTipsCountResponse() (response *DescribeDomainSecurityGroupTipsCountResponse) {
	response = &DescribeDomainSecurityGroupTipsCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询全量域名安全组绑定失败或安全组变更的数量
func (c *Client) DescribeDomainSecurityGroupTipsCount(request *DescribeDomainSecurityGroupTipsCountRequest) (response *DescribeDomainSecurityGroupTipsCountResponse, err error) {
	if request == nil {
		request = NewDescribeDomainSecurityGroupTipsCountRequest()
	}
	response = NewDescribeDomainSecurityGroupTipsCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDnsRuleRequest() (request *ModifyDnsRuleRequest) {
	request = &ModifyDnsRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyDnsRule")
	return
}

func NewModifyDnsRuleResponse() (response *ModifyDnsRuleResponse) {
	response = &ModifyDnsRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将接入WAF的域名切回源站/或者切回WAF
func (c *Client) ModifyDnsRule(request *ModifyDnsRuleRequest) (response *ModifyDnsRuleResponse, err error) {
	if request == nil {
		request = NewModifyDnsRuleRequest()
	}
	response = NewModifyDnsRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityRuleCountRequest() (request *DescribeSecurityRuleCountRequest) {
	request = &DescribeSecurityRuleCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeSecurityRuleCount")
	return
}

func NewDescribeSecurityRuleCountResponse() (response *DescribeSecurityRuleCountResponse) {
	response = &DescribeSecurityRuleCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 得到基础安全模块中，每个模块的规则数目
func (c *Client) DescribeSecurityRuleCount(request *DescribeSecurityRuleCountRequest) (response *DescribeSecurityRuleCountResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityRuleCountRequest()
	}
	response = NewDescribeSecurityRuleCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTCBRuleRequest() (request *DescribeBotTCBRuleRequest) {
	request = &DescribeBotTCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotTCBRule")
	return
}

func NewDescribeBotTCBRuleResponse() (response *DescribeBotTCBRuleResponse) {
	response = &DescribeBotTCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取TCB规则
func (c *Client) DescribeBotTCBRule(request *DescribeBotTCBRuleRequest) (response *DescribeBotTCBRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotTCBRuleRequest()
	}
	response = NewDescribeBotTCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTsResourceRequest() (request *DeleteTsResourceRequest) {
	request = &DeleteTsResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteTsResource")
	return
}

func NewDeleteTsResourceResponse() (response *DeleteTsResourceResponse) {
	response = &DeleteTsResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自研引擎服务资源
func (c *Client) DeleteTsResource(request *DeleteTsResourceRequest) (response *DeleteTsResourceResponse, err error) {
	if request == nil {
		request = NewDeleteTsResourceRequest()
	}
	response = NewDeleteTsResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePolicyStatusRequest() (request *DescribePolicyStatusRequest) {
	request = &DescribePolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePolicyStatus")
	return
}

func NewDescribePolicyStatusResponse() (response *DescribePolicyStatusResponse) {
	response = &DescribePolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防护状态以及生效的实例id
func (c *Client) DescribePolicyStatus(request *DescribePolicyStatusRequest) (response *DescribePolicyStatusResponse, err error) {
	if request == nil {
		request = NewDescribePolicyStatusRequest()
	}
	response = NewDescribePolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotGenerateDealsRequest() (request *ModifyBotGenerateDealsRequest) {
	request = &ModifyBotGenerateDealsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotGenerateDeals")
	return
}

func NewModifyBotGenerateDealsResponse() (response *ModifyBotGenerateDealsResponse) {
	response = &ModifyBotGenerateDealsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给使用的waf实例 bot下单接口
func (c *Client) ModifyBotGenerateDeals(request *ModifyBotGenerateDealsRequest) (response *ModifyBotGenerateDealsResponse, err error) {
	if request == nil {
		request = NewModifyBotGenerateDealsRequest()
	}
	response = NewModifyBotGenerateDealsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClsUsedRequest() (request *DescribeClsUsedRequest) {
	request = &DescribeClsUsedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeClsUsed")
	return
}

func NewDescribeClsUsedResponse() (response *DescribeClsUsedResponse) {
	response = &DescribeClsUsedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回日志用量
func (c *Client) DescribeClsUsed(request *DescribeClsUsedRequest) (response *DescribeClsUsedResponse, err error) {
	if request == nil {
		request = NewDescribeClsUsedRequest()
	}
	response = NewDescribeClsUsedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRecordItemsRequest() (request *DescribeBotRecordItemsRequest) {
	request = &DescribeBotRecordItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotRecordItems")
	return
}

func NewDescribeBotRecordItemsResponse() (response *DescribeBotRecordItemsResponse) {
	response = &DescribeBotRecordItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 bot记录访问详情
func (c *Client) DescribeBotRecordItems(request *DescribeBotRecordItemsRequest) (response *DescribeBotRecordItemsResponse, err error) {
	if request == nil {
		request = NewDescribeBotRecordItemsRequest()
	}
	response = NewDescribeBotRecordItemsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyJsInjectRuleRequest() (request *ModifyJsInjectRuleRequest) {
	request = &ModifyJsInjectRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyJsInjectRule")
	return
}

func NewModifyJsInjectRuleResponse() (response *ModifyJsInjectRuleResponse) {
	response = &ModifyJsInjectRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新域名的前端对抗规则
func (c *Client) ModifyJsInjectRule(request *ModifyJsInjectRuleRequest) (response *ModifyJsInjectRuleResponse, err error) {
	if request == nil {
		request = NewModifyJsInjectRuleRequest()
	}
	response = NewModifyJsInjectRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAccessLogRequest() (request *DescribeBotAccessLogRequest) {
	request = &DescribeBotAccessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotAccessLog")
	return
}

func NewDescribeBotAccessLogResponse() (response *DescribeBotAccessLogResponse) {
	response = &DescribeBotAccessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot访问日志列表接口
func (c *Client) DescribeBotAccessLog(request *DescribeBotAccessLogRequest) (response *DescribeBotAccessLogResponse, err error) {
	if request == nil {
		request = NewDescribeBotAccessLogRequest()
	}
	response = NewDescribeBotAccessLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotStatusRequest() (request *DescribeBotStatusRequest) {
	request = &DescribeBotStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotStatus")
	return
}

func NewDescribeBotStatusResponse() (response *DescribeBotStatusResponse) {
	response = &DescribeBotStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取bot开关
func (c *Client) DescribeBotStatus(request *DescribeBotStatusRequest) (response *DescribeBotStatusResponse, err error) {
	if request == nil {
		request = NewDescribeBotStatusRequest()
	}
	response = NewDescribeBotStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpSignatureRuleRequest() (request *DescribeOpSignatureRuleRequest) {
	request = &DescribeOpSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeOpSignatureRule")
	return
}

func NewDescribeOpSignatureRuleResponse() (response *DescribeOpSignatureRuleResponse) {
	response = &DescribeOpSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口，获取全量特征规则列表
func (c *Client) DescribeOpSignatureRule(request *DescribeOpSignatureRuleRequest) (response *DescribeOpSignatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeOpSignatureRuleRequest()
	}
	response = NewDescribeOpSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserCdcClbWafRegionsRequest() (request *DescribeUserCdcClbWafRegionsRequest) {
	request = &DescribeUserCdcClbWafRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserCdcClbWafRegions")
	return
}

func NewDescribeUserCdcClbWafRegionsResponse() (response *DescribeUserCdcClbWafRegionsResponse) {
	response = &DescribeUserCdcClbWafRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在CDC场景下，负载均衡型WAF的添加、编辑域名配置的时候，需要展示CDC负载均衡型WAF（cdc-clb-waf)支持的地域列表，通过DescribeUserCdcClbWafRegions既可以获得当前对客户已经开放的地域列表
func (c *Client) DescribeUserCdcClbWafRegions(request *DescribeUserCdcClbWafRegionsRequest) (response *DescribeUserCdcClbWafRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeUserCdcClbWafRegionsRequest()
	}
	response = NewDescribeUserCdcClbWafRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserQPSRequest() (request *DescribeUserQPSRequest) {
	request = &DescribeUserQPSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeUserQPS")
	return
}

func NewDescribeUserQPSResponse() (response *DescribeUserQPSResponse) {
	response = &DescribeUserQPSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取客户的QPS峰值
func (c *Client) DescribeUserQPS(request *DescribeUserQPSRequest) (response *DescribeUserQPSResponse, err error) {
	if request == nil {
		request = NewDescribeUserQPSRequest()
	}
	response = NewDescribeUserQPSResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTsRegionRequest() (request *DeleteTsRegionRequest) {
	request = &DeleteTsRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteTsRegion")
	return
}

func NewDeleteTsRegionResponse() (response *DeleteTsRegionResponse) {
	response = &DeleteTsRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除ts接入区域
func (c *Client) DeleteTsRegion(request *DeleteTsRegionRequest) (response *DeleteTsRegionResponse, err error) {
	if request == nil {
		request = NewDeleteTsRegionRequest()
	}
	response = NewDeleteTsRegionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainEngineRequest() (request *ModifyDomainEngineRequest) {
	request = &ModifyDomainEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainEngine")
	return
}

func NewModifyDomainEngineResponse() (response *ModifyDomainEngineResponse) {
	response = &ModifyDomainEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名引擎类型
func (c *Client) ModifyDomainEngine(request *ModifyDomainEngineRequest) (response *ModifyDomainEngineResponse, err error) {
	if request == nil {
		request = NewModifyDomainEngineRequest()
	}
	response = NewModifyDomainEngineResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpartaProtectionModeRequest() (request *ModifySpartaProtectionModeRequest) {
	request = &ModifySpartaProtectionModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifySpartaProtectionMode")
	return
}

func NewModifySpartaProtectionModeResponse() (response *ModifySpartaProtectionModeResponse) {
	response = &ModifySpartaProtectionModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置waf防护状态
func (c *Client) ModifySpartaProtectionMode(request *ModifySpartaProtectionModeRequest) (response *ModifySpartaProtectionModeResponse, err error) {
	if request == nil {
		request = NewModifySpartaProtectionModeRequest()
	}
	response = NewModifySpartaProtectionModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpAttackWhiteRuleNewRequest() (request *DescribeOpAttackWhiteRuleNewRequest) {
	request = &DescribeOpAttackWhiteRuleNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeOpAttackWhiteRuleNew")
	return
}

func NewDescribeOpAttackWhiteRuleNewResponse() (response *DescribeOpAttackWhiteRuleNewResponse) {
	response = &DescribeOpAttackWhiteRuleNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营平台接口：获取用户规则白名单列表
func (c *Client) DescribeOpAttackWhiteRuleNew(request *DescribeOpAttackWhiteRuleNewRequest) (response *DescribeOpAttackWhiteRuleNewResponse, err error) {
	if request == nil {
		request = NewDescribeOpAttackWhiteRuleNewRequest()
	}
	response = NewDescribeOpAttackWhiteRuleNewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebshellStatusRequest() (request *ModifyWebshellStatusRequest) {
	request = &ModifyWebshellStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyWebshellStatus")
	return
}

func NewModifyWebshellStatusResponse() (response *ModifyWebshellStatusResponse) {
	response = &ModifyWebshellStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置域名的webshell状态。
func (c *Client) ModifyWebshellStatus(request *ModifyWebshellStatusRequest) (response *ModifyWebshellStatusResponse, err error) {
	if request == nil {
		request = NewModifyWebshellStatusRequest()
	}
	response = NewModifyWebshellStatusResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchTrafficMarkingRequest() (request *SwitchTrafficMarkingRequest) {
	request = &SwitchTrafficMarkingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "SwitchTrafficMarking")
	return
}

func NewSwitchTrafficMarkingResponse() (response *SwitchTrafficMarkingResponse) {
	response = &SwitchTrafficMarkingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换流量标记开关
func (c *Client) SwitchTrafficMarking(request *SwitchTrafficMarkingRequest) (response *SwitchTrafficMarkingResponse, err error) {
	if request == nil {
		request = NewSwitchTrafficMarkingRequest()
	}
	response = NewSwitchTrafficMarkingResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePostCLSFlowRequest() (request *CreatePostCLSFlowRequest) {
	request = &CreatePostCLSFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "CreatePostCLSFlow")
	return
}

func NewCreatePostCLSFlowResponse() (response *CreatePostCLSFlowResponse) {
	response = &CreatePostCLSFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建CLS投递流任务
func (c *Client) CreatePostCLSFlow(request *CreatePostCLSFlowRequest) (response *CreatePostCLSFlowResponse, err error) {
	if request == nil {
		request = NewCreatePostCLSFlowRequest()
	}
	response = NewCreatePostCLSFlowResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotActionRuleRequest() (request *ModifyBotActionRuleRequest) {
	request = &ModifyBotActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyBotActionRule")
	return
}

func NewModifyBotActionRuleResponse() (response *ModifyBotActionRuleResponse) {
	response = &ModifyBotActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot-动作分布配置
func (c *Client) ModifyBotActionRule(request *ModifyBotActionRuleRequest) (response *ModifyBotActionRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotActionRuleRequest()
	}
	response = NewModifyBotActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStatisticTypesRequest() (request *DescribeStatisticTypesRequest) {
	request = &DescribeStatisticTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeStatisticTypes")
	return
}

func NewDescribeStatisticTypesResponse() (response *DescribeStatisticTypesResponse) {
	response = &DescribeStatisticTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取攻击类型统计数据
func (c *Client) DescribeStatisticTypes(request *DescribeStatisticTypesRequest) (response *DescribeStatisticTypesResponse, err error) {
	if request == nil {
		request = NewDescribeStatisticTypesRequest()
	}
	response = NewDescribeStatisticTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTCBRecordsRequest() (request *DescribeBotTCBRecordsRequest) {
	request = &DescribeBotTCBRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeBotTCBRecords")
	return
}

func NewDescribeBotTCBRecordsResponse() (response *DescribeBotTCBRecordsResponse) {
	response = &DescribeBotTCBRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取TCB类型Bots
func (c *Client) DescribeBotTCBRecords(request *DescribeBotTCBRecordsRequest) (response *DescribeBotTCBRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeBotTCBRecordsRequest()
	}
	response = NewDescribeBotTCBRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpUserSignatureRuleRequest() (request *DescribeOpUserSignatureRuleRequest) {
	request = &DescribeOpUserSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeOpUserSignatureRule")
	return
}

func NewDescribeOpUserSignatureRuleResponse() (response *DescribeOpUserSignatureRuleResponse) {
	response = &DescribeOpUserSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为运营后台获取用户规则
func (c *Client) DescribeOpUserSignatureRule(request *DescribeOpUserSignatureRuleRequest) (response *DescribeOpUserSignatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeOpUserSignatureRuleRequest()
	}
	response = NewDescribeOpUserSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOpSignatureRuleRequest() (request *ModifyOpSignatureRuleRequest) {
	request = &ModifyOpSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ModifyOpSignatureRule")
	return
}

func NewModifyOpSignatureRuleResponse() (response *ModifyOpSignatureRuleResponse) {
	response = &ModifyOpSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：修改规则
func (c *Client) ModifyOpSignatureRule(request *ModifyOpSignatureRuleRequest) (response *ModifyOpSignatureRuleResponse, err error) {
	if request == nil {
		request = NewModifyOpSignatureRuleRequest()
	}
	response = NewModifyOpSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiOverviewRequest() (request *DescribeApiOverviewRequest) {
	request = &DescribeApiOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiOverview")
	return
}

func NewDescribeApiOverviewResponse() (response *DescribeApiOverviewResponse) {
	response = &DescribeApiOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api资产概览数据
func (c *Client) DescribeApiOverview(request *DescribeApiOverviewRequest) (response *DescribeApiOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeApiOverviewRequest()
	}
	response = NewDescribeApiOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackWorldMapRequest() (request *DescribeAttackWorldMapRequest) {
	request = &DescribeAttackWorldMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackWorldMap")
	return
}

func NewDescribeAttackWorldMapResponse() (response *DescribeAttackWorldMapResponse) {
	response = &DescribeAttackWorldMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定时间段对域名攻击
func (c *Client) DescribeAttackWorldMap(request *DescribeAttackWorldMapRequest) (response *DescribeAttackWorldMapResponse, err error) {
	if request == nil {
		request = NewDescribeAttackWorldMapRequest()
	}
	response = NewDescribeAttackWorldMapResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAntiFakeRulesRequest() (request *DescribeAntiFakeRulesRequest) {
	request = &DescribeAntiFakeRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiFakeRules")
	return
}

func NewDescribeAntiFakeRulesResponse() (response *DescribeAntiFakeRulesResponse) {
	response = &DescribeAntiFakeRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防篡改url
func (c *Client) DescribeAntiFakeRules(request *DescribeAntiFakeRulesRequest) (response *DescribeAntiFakeRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAntiFakeRulesRequest()
	}
	response = NewDescribeAntiFakeRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainEngineRequest() (request *DescribeDomainEngineRequest) {
	request = &DescribeDomainEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainEngine")
	return
}

func NewDescribeDomainEngineResponse() (response *DescribeDomainEngineResponse) {
	response = &DescribeDomainEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询域名引擎类型
func (c *Client) DescribeDomainEngine(request *DescribeDomainEngineRequest) (response *DescribeDomainEngineResponse, err error) {
	if request == nil {
		request = NewDescribeDomainEngineRequest()
	}
	response = NewDescribeDomainEngineResponse()
	err = c.Send(request, response)
	return
}

func NewImportApiRulesRequest() (request *ImportApiRulesRequest) {
	request = &ImportApiRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "ImportApiRules")
	return
}

func NewImportApiRulesResponse() (response *ImportApiRulesResponse) {
	response = &ImportApiRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户手动导入api规则
func (c *Client) ImportApiRules(request *ImportApiRulesRequest) (response *ImportApiRulesResponse, err error) {
	if request == nil {
		request = NewImportApiRulesRequest()
	}
	response = NewImportApiRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackIndexRequest() (request *DescribeAttackIndexRequest) {
	request = &DescribeAttackIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackIndex")
	return
}

func NewDescribeAttackIndexResponse() (response *DescribeAttackIndexResponse) {
	response = &DescribeAttackIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 攻击日志索引描述详情
func (c *Client) DescribeAttackIndex(request *DescribeAttackIndexRequest) (response *DescribeAttackIndexResponse, err error) {
	if request == nil {
		request = NewDescribeAttackIndexRequest()
	}
	response = NewDescribeAttackIndexResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePocTestAuthorizationStatusRequest() (request *DescribePocTestAuthorizationStatusRequest) {
	request = &DescribePocTestAuthorizationStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePocTestAuthorizationStatus")
	return
}

func NewDescribePocTestAuthorizationStatusResponse() (response *DescribePocTestAuthorizationStatusResponse) {
	response = &DescribePocTestAuthorizationStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询APPID授权状态
func (c *Client) DescribePocTestAuthorizationStatus(request *DescribePocTestAuthorizationStatusRequest) (response *DescribePocTestAuthorizationStatusResponse, err error) {
	if request == nil {
		request = NewDescribePocTestAuthorizationStatusRequest()
	}
	response = NewDescribePocTestAuthorizationStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostCKafkaFlowsRequest() (request *DescribePostCKafkaFlowsRequest) {
	request = &DescribePostCKafkaFlowsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribePostCKafkaFlows")
	return
}

func NewDescribePostCKafkaFlowsResponse() (response *DescribePostCKafkaFlowsResponse) {
	response = &DescribePostCKafkaFlowsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CKafka投递流任务列表
func (c *Client) DescribePostCKafkaFlows(request *DescribePostCKafkaFlowsRequest) (response *DescribePostCKafkaFlowsResponse, err error) {
	if request == nil {
		request = NewDescribePostCKafkaFlowsRequest()
	}
	response = NewDescribePostCKafkaFlowsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecSensitiveRuleListRequest() (request *DescribeApiSecSensitiveRuleListRequest) {
	request = &DescribeApiSecSensitiveRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DescribeApiSecSensitiveRuleList")
	return
}

func NewDescribeApiSecSensitiveRuleListResponse() (response *DescribeApiSecSensitiveRuleListResponse) {
	response = &DescribeApiSecSensitiveRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取api安全敏感规则列表
func (c *Client) DescribeApiSecSensitiveRuleList(request *DescribeApiSecSensitiveRuleListRequest) (response *DescribeApiSecSensitiveRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecSensitiveRuleListRequest()
	}
	response = NewDescribeApiSecSensitiveRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyPostCKafkaFlowRequest() (request *DestroyPostCKafkaFlowRequest) {
	request = &DestroyPostCKafkaFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DestroyPostCKafkaFlow")
	return
}

func NewDestroyPostCKafkaFlowResponse() (response *DestroyPostCKafkaFlowResponse) {
	response = &DestroyPostCKafkaFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁CKafka投递流任务
func (c *Client) DestroyPostCKafkaFlow(request *DestroyPostCKafkaFlowRequest) (response *DestroyPostCKafkaFlowResponse, err error) {
	if request == nil {
		request = NewDestroyPostCKafkaFlowRequest()
	}
	response = NewDestroyPostCKafkaFlowResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHostInnerRequest() (request *DeleteHostInnerRequest) {
	request = &DeleteHostInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("waf", APIVersion, "DeleteHostInner")
	return
}

func NewDeleteHostInnerResponse() (response *DeleteHostInnerResponse) {
	response = &DeleteHostInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除CLB-WAF防护域名
// 支持批量操作
func (c *Client) DeleteHostInner(request *DeleteHostInnerRequest) (response *DeleteHostInnerResponse, err error) {
	if request == nil {
		request = NewDeleteHostInnerRequest()
	}
	response = NewDeleteHostInnerResponse()
	err = c.Send(request, response)
	return
}
