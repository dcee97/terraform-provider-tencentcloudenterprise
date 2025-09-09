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

package v20190904

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-09-04"

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

func NewModifyNatFwVpcDnsSwitchRequest() (request *ModifyNatFwVpcDnsSwitchRequest) {
	request = &ModifyNatFwVpcDnsSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwVpcDnsSwitch")
	return
}

func NewModifyNatFwVpcDnsSwitchResponse() (response *ModifyNatFwVpcDnsSwitchResponse) {
	response = &ModifyNatFwVpcDnsSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// nat 防火墙VPC DNS 开关切换
func (c *Client) ModifyNatFwVpcDnsSwitch(request *ModifyNatFwVpcDnsSwitchRequest) (response *ModifyNatFwVpcDnsSwitchResponse, err error) {
	if request == nil {
		request = NewModifyNatFwVpcDnsSwitchRequest()
	}
	response = NewModifyNatFwVpcDnsSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCFWInfoRequest() (request *DescribeCFWInfoRequest) {
	request = &DescribeCFWInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCFWInfo")
	return
}

func NewDescribeCFWInfoResponse() (response *DescribeCFWInfoResponse) {
	response = &DescribeCFWInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产讯防火墙用户的详细信息
func (c *Client) DescribeCFWInfo(request *DescribeCFWInfoRequest) (response *DescribeCFWInfoResponse, err error) {
	if request == nil {
		request = NewDescribeCFWInfoRequest()
	}
	response = NewDescribeCFWInfoResponse()
	err = c.Send(request, response)
	return
}

func NewQueryVpcFwSwitchModeRequest() (request *QueryVpcFwSwitchModeRequest) {
	request = &QueryVpcFwSwitchModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "QueryVpcFwSwitchMode")
	return
}

func NewQueryVpcFwSwitchModeResponse() (response *QueryVpcFwSwitchModeResponse) {
	response = &QueryVpcFwSwitchModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VPC间防火墙开关模式，云联网模式的防火墙实例根据实际情况可能仅支持同开同关模式
func (c *Client) QueryVpcFwSwitchMode(request *QueryVpcFwSwitchModeRequest) (response *QueryVpcFwSwitchModeResponse, err error) {
	if request == nil {
		request = NewQueryVpcFwSwitchModeRequest()
	}
	response = NewQueryVpcFwSwitchModeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateChooseVpcsRequest() (request *CreateChooseVpcsRequest) {
	request = &CreateChooseVpcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateChooseVpcs")
	return
}

func NewCreateChooseVpcsResponse() (response *CreateChooseVpcsResponse) {
	response = &CreateChooseVpcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建、选择vpc
func (c *Client) CreateChooseVpcs(request *CreateChooseVpcsRequest) (response *CreateChooseVpcsResponse, err error) {
	if request == nil {
		request = NewCreateChooseVpcsRequest()
	}
	response = NewCreateChooseVpcsResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourceGroupChangeImpactRequest() (request *GetResourceGroupChangeImpactRequest) {
	request = &GetResourceGroupChangeImpactRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "GetResourceGroupChangeImpact")
	return
}

func NewGetResourceGroupChangeImpactResponse() (response *GetResourceGroupChangeImpactResponse) {
	response = &GetResourceGroupChangeImpactResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源分组变更影响
func (c *Client) GetResourceGroupChangeImpact(request *GetResourceGroupChangeImpactRequest) (response *GetResourceGroupChangeImpactResponse, err error) {
	if request == nil {
		request = NewGetResourceGroupChangeImpactRequest()
	}
	response = NewGetResourceGroupChangeImpactResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpStatLstRequest() (request *DescribeIpStatLstRequest) {
	request = &DescribeIpStatLstRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeIpStatLst")
	return
}

func NewDescribeIpStatLstResponse() (response *DescribeIpStatLstResponse) {
	response = &DescribeIpStatLstResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取nat vpc 的单台 cvm的峰值带宽
func (c *Client) DescribeIpStatLst(request *DescribeIpStatLstRequest) (response *DescribeIpStatLstResponse, err error) {
	if request == nil {
		request = NewDescribeIpStatLstRequest()
	}
	response = NewDescribeIpStatLstResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatFwInstancesInfoRequest() (request *DescribeNatFwInstancesInfoRequest) {
	request = &DescribeNatFwInstancesInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInstancesInfo")
	return
}

func NewDescribeNatFwInstancesInfoResponse() (response *DescribeNatFwInstancesInfoResponse) {
	response = &DescribeNatFwInstancesInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetNatInstance 获取租户所有NAT实例及实例卡片信息
func (c *Client) DescribeNatFwInstancesInfo(request *DescribeNatFwInstancesInfoRequest) (response *DescribeNatFwInstancesInfoResponse, err error) {
	if request == nil {
		request = NewDescribeNatFwInstancesInfoRequest()
	}
	response = NewDescribeNatFwInstancesInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatFwTcRuleRequest() (request *DescribeNatFwTcRuleRequest) {
	request = &DescribeNatFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwTcRule")
	return
}

func NewDescribeNatFwTcRuleResponse() (response *DescribeNatFwTcRuleResponse) {
	response = &DescribeNatFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询流控策略列表
func (c *Client) DescribeNatFwTcRule(request *DescribeNatFwTcRuleRequest) (response *DescribeNatFwTcRuleResponse, err error) {
	if request == nil {
		request = NewDescribeNatFwTcRuleRequest()
	}
	response = NewDescribeNatFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHoneyPotAttackerRequest() (request *DescribeHoneyPotAttackerRequest) {
	request = &DescribeHoneyPotAttackerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeHoneyPotAttacker")
	return
}

func NewDescribeHoneyPotAttackerResponse() (response *DescribeHoneyPotAttackerResponse) {
	response = &DescribeHoneyPotAttackerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取攻击者汇总列表
func (c *Client) DescribeHoneyPotAttacker(request *DescribeHoneyPotAttackerRequest) (response *DescribeHoneyPotAttackerResponse, err error) {
	if request == nil {
		request = NewDescribeHoneyPotAttackerRequest()
	}
	response = NewDescribeHoneyPotAttackerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSourceAssetRequest() (request *DescribeSourceAssetRequest) {
	request = &DescribeSourceAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeSourceAsset")
	return
}

func NewDescribeSourceAssetResponse() (response *DescribeSourceAssetResponse) {
	response = &DescribeSourceAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeSourceAsset-查询资产组全部资产信息
func (c *Client) DescribeSourceAsset(request *DescribeSourceAssetRequest) (response *DescribeSourceAssetResponse, err error) {
	if request == nil {
		request = NewDescribeSourceAssetRequest()
	}
	response = NewDescribeSourceAssetResponse()
	err = c.Send(request, response)
	return
}

func NewAddNatFwObjRequest() (request *AddNatFwObjRequest) {
	request = &AddNatFwObjRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddNatFwObj")
	return
}

func NewAddNatFwObjResponse() (response *AddNatFwObjResponse) {
	response = &AddNatFwObjResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加NAT防火墙vpc或nat接入对象
func (c *Client) AddNatFwObj(request *AddNatFwObjRequest) (response *AddNatFwObjResponse, err error) {
	if request == nil {
		request = NewAddNatFwObjRequest()
	}
	response = NewAddNatFwObjResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatAcRuleRequest() (request *DescribeNatAcRuleRequest) {
	request = &DescribeNatAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatAcRule")
	return
}

func NewDescribeNatAcRuleResponse() (response *DescribeNatAcRuleResponse) {
	response = &DescribeNatAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询NAT访问控制列表
func (c *Client) DescribeNatAcRule(request *DescribeNatAcRuleRequest) (response *DescribeNatAcRuleResponse, err error) {
	if request == nil {
		request = NewDescribeNatAcRuleRequest()
	}
	response = NewDescribeNatAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcAcRuleRequest() (request *ModifyVpcAcRuleRequest) {
	request = &ModifyVpcAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcAcRule")
	return
}

func NewModifyVpcAcRuleResponse() (response *ModifyVpcAcRuleResponse) {
	response = &ModifyVpcAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改内网间访问控制规则
func (c *Client) ModifyVpcAcRule(request *ModifyVpcAcRuleRequest) (response *ModifyVpcAcRuleResponse, err error) {
	if request == nil {
		request = NewModifyVpcAcRuleRequest()
	}
	response = NewModifyVpcAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfwVersionRequest() (request *DescribeCfwVersionRequest) {
	request = &DescribeCfwVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwVersion")
	return
}

func NewDescribeCfwVersionResponse() (response *DescribeCfwVersionResponse) {
	response = &DescribeCfwVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeCfwVersion 查询NAT VPC 防火墙的引擎模式和版本号
func (c *Client) DescribeCfwVersion(request *DescribeCfwVersionRequest) (response *DescribeCfwVersionResponse, err error) {
	if request == nil {
		request = NewDescribeCfwVersionRequest()
	}
	response = NewDescribeCfwVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlertFeedBackListRequest() (request *DescribeAlertFeedBackListRequest) {
	request = &DescribeAlertFeedBackListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAlertFeedBackList")
	return
}

func NewDescribeAlertFeedBackListResponse() (response *DescribeAlertFeedBackListResponse) {
	response = &DescribeAlertFeedBackListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警误报反馈列表查询
func (c *Client) DescribeAlertFeedBackList(request *DescribeAlertFeedBackListRequest) (response *DescribeAlertFeedBackListResponse, err error) {
	if request == nil {
		request = NewDescribeAlertFeedBackListRequest()
	}
	response = NewDescribeAlertFeedBackListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventNameListRequest() (request *DescribeEventNameListRequest) {
	request = &DescribeEventNameListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeEventNameList")
	return
}

func NewDescribeEventNameListResponse() (response *DescribeEventNameListResponse) {
	response = &DescribeEventNameListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全事件名称列表
func (c *Client) DescribeEventNameList(request *DescribeEventNameListRequest) (response *DescribeEventNameListResponse, err error) {
	if request == nil {
		request = NewDescribeEventNameListRequest()
	}
	response = NewDescribeEventNameListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatAcRuleRequest() (request *ModifyNatAcRuleRequest) {
	request = &ModifyNatAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatAcRule")
	return
}

func NewModifyNatAcRuleResponse() (response *ModifyNatAcRuleResponse) {
	response = &ModifyNatAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改NAT访问控制规则
func (c *Client) ModifyNatAcRule(request *ModifyNatAcRuleRequest) (response *ModifyNatAcRuleResponse, err error) {
	if request == nil {
		request = NewModifyNatAcRuleRequest()
	}
	response = NewModifyNatAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfwExportBindRequest() (request *DescribeCfwExportBindRequest) {
	request = &DescribeCfwExportBindRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwExportBind")
	return
}

func NewDescribeCfwExportBindResponse() (response *DescribeCfwExportBindResponse) {
	response = &DescribeCfwExportBindResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询NAT防火墙出口绑定列表
func (c *Client) DescribeCfwExportBind(request *DescribeCfwExportBindRequest) (response *DescribeCfwExportBindResponse, err error) {
	if request == nil {
		request = NewDescribeCfwExportBindRequest()
	}
	response = NewDescribeCfwExportBindResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTiWhiteRequest() (request *ModifyTiWhiteRequest) {
	request = &ModifyTiWhiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyTiWhite")
	return
}

func NewModifyTiWhiteResponse() (response *ModifyTiWhiteResponse) {
	response = &ModifyTiWhiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ti白名单修改接口
func (c *Client) ModifyTiWhite(request *ModifyTiWhiteRequest) (response *ModifyTiWhiteResponse, err error) {
	if request == nil {
		request = NewModifyTiWhiteRequest()
	}
	response = NewModifyTiWhiteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllRegionListRequest() (request *DescribeAllRegionListRequest) {
	request = &DescribeAllRegionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAllRegionList")
	return
}

func NewDescribeAllRegionListResponse() (response *DescribeAllRegionListResponse) {
	response = &DescribeAllRegionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询地域配置信息-DescribeAllRegionList
func (c *Client) DescribeAllRegionList(request *DescribeAllRegionListRequest) (response *DescribeAllRegionListResponse, err error) {
	if request == nil {
		request = NewDescribeAllRegionListRequest()
	}
	response = NewDescribeAllRegionListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBetaTaskAclRuleListRequest() (request *DescribeBetaTaskAclRuleListRequest) {
	request = &DescribeBetaTaskAclRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBetaTaskAclRuleList")
	return
}

func NewDescribeBetaTaskAclRuleListResponse() (response *DescribeBetaTaskAclRuleListResponse) {
	response = &DescribeBetaTaskAclRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询beta规则
func (c *Client) DescribeBetaTaskAclRuleList(request *DescribeBetaTaskAclRuleListRequest) (response *DescribeBetaTaskAclRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeBetaTaskAclRuleListRequest()
	}
	response = NewDescribeBetaTaskAclRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcProtectListRequest() (request *DescribeVpcProtectListRequest) {
	request = &DescribeVpcProtectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcProtectList")
	return
}

func NewDescribeVpcProtectListResponse() (response *DescribeVpcProtectListResponse) {
	response = &DescribeVpcProtectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// VPC防火墙防护模式列表
func (c *Client) DescribeVpcProtectList(request *DescribeVpcProtectListRequest) (response *DescribeVpcProtectListResponse, err error) {
	if request == nil {
		request = NewDescribeVpcProtectListRequest()
	}
	response = NewDescribeVpcProtectListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNatFwInstanceRequest() (request *CreateNatFwInstanceRequest) {
	request = &CreateNatFwInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateNatFwInstance")
	return
}

func NewCreateNatFwInstanceResponse() (response *CreateNatFwInstanceResponse) {
	response = &CreateNatFwInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建防火墙实例
func (c *Client) CreateNatFwInstance(request *CreateNatFwInstanceRequest) (response *CreateNatFwInstanceResponse, err error) {
	if request == nil {
		request = NewCreateNatFwInstanceRequest()
	}
	response = NewCreateNatFwInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowApiDispatchRequest() (request *DescribeFlowApiDispatchRequest) {
	request = &DescribeFlowApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFlowApiDispatch")
	return
}

func NewDescribeFlowApiDispatchResponse() (response *DescribeFlowApiDispatchResponse) {
	response = &DescribeFlowApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量中心读接口请求中转
func (c *Client) DescribeFlowApiDispatch(request *DescribeFlowApiDispatchRequest) (response *DescribeFlowApiDispatchResponse, err error) {
	if request == nil {
		request = NewDescribeFlowApiDispatchRequest()
	}
	response = NewDescribeFlowApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBakRuleListRequest() (request *CreateBakRuleListRequest) {
	request = &CreateBakRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateBakRuleList")
	return
}

func NewCreateBakRuleListResponse() (response *CreateBakRuleListResponse) {
	response = &CreateBakRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建备份规则列表
func (c *Client) CreateBakRuleList(request *CreateBakRuleListRequest) (response *CreateBakRuleListResponse, err error) {
	if request == nil {
		request = NewCreateBakRuleListRequest()
	}
	response = NewCreateBakRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewModifySwitchCommonRequest() (request *ModifySwitchCommonRequest) {
	request = &ModifySwitchCommonRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifySwitchCommon")
	return
}

func NewModifySwitchCommonResponse() (response *ModifySwitchCommonResponse) {
	response = &ModifySwitchCommonResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通用开关接口状态修改
func (c *Client) ModifySwitchCommon(request *ModifySwitchCommonRequest) (response *ModifySwitchCommonResponse, err error) {
	if request == nil {
		request = NewModifySwitchCommonRequest()
	}
	response = NewModifySwitchCommonResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveVpcFwTcRuleRequest() (request *RemoveVpcFwTcRuleRequest) {
	request = &RemoveVpcFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "RemoveVpcFwTcRule")
	return
}

func NewRemoveVpcFwTcRuleResponse() (response *RemoveVpcFwTcRuleResponse) {
	response = &RemoveVpcFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除流量控制策略
func (c *Client) RemoveVpcFwTcRule(request *RemoveVpcFwTcRuleRequest) (response *RemoveVpcFwTcRuleResponse, err error) {
	if request == nil {
		request = NewRemoveVpcFwTcRuleRequest()
	}
	response = NewRemoveVpcFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwJoinInstancesRequest() (request *DescribeVpcFwJoinInstancesRequest) {
	request = &DescribeVpcFwJoinInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwJoinInstances")
	return
}

func NewDescribeVpcFwJoinInstancesResponse() (response *DescribeVpcFwJoinInstancesResponse) {
	response = &DescribeVpcFwJoinInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户所有VPC防火墙接入网络实例列表
func (c *Client) DescribeVpcFwJoinInstances(request *DescribeVpcFwJoinInstancesRequest) (response *DescribeVpcFwJoinInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwJoinInstancesRequest()
	}
	response = NewDescribeVpcFwJoinInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWidthCheckRequest() (request *ModifyWidthCheckRequest) {
	request = &ModifyWidthCheckRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyWidthCheck")
	return
}

func NewModifyWidthCheckResponse() (response *ModifyWidthCheckResponse) {
	response = &ModifyWidthCheckResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 带宽调整配置检查
func (c *Client) ModifyWidthCheck(request *ModifyWidthCheckRequest) (response *ModifyWidthCheckResponse, err error) {
	if request == nil {
		request = NewModifyWidthCheckRequest()
	}
	response = NewModifyWidthCheckResponse()
	err = c.Send(request, response)
	return
}

func NewSetCfwInsBypassRequest() (request *SetCfwInsBypassRequest) {
	request = &SetCfwInsBypassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "SetCfwInsBypass")
	return
}

func NewSetCfwInsBypassResponse() (response *SetCfwInsBypassResponse) {
	response = &SetCfwInsBypassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置防火墙实例bypass模式
func (c *Client) SetCfwInsBypass(request *SetCfwInsBypassRequest) (response *SetCfwInsBypassResponse, err error) {
	if request == nil {
		request = NewSetCfwInsBypassRequest()
	}
	response = NewSetCfwInsBypassResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEngineRequest() (request *UpdateEngineRequest) {
	request = &UpdateEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "UpdateEngine")
	return
}

func NewUpdateEngineResponse() (response *UpdateEngineResponse) {
	response = &UpdateEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级防火墙引擎
func (c *Client) UpdateEngine(request *UpdateEngineRequest) (response *UpdateEngineResponse, err error) {
	if request == nil {
		request = NewUpdateEngineRequest()
	}
	response = NewUpdateEngineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogApiDispatchRequest() (request *DescribeLogApiDispatchRequest) {
	request = &DescribeLogApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeLogApiDispatch")
	return
}

func NewDescribeLogApiDispatchResponse() (response *DescribeLogApiDispatchResponse) {
	response = &DescribeLogApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志审计读接口请求中转
func (c *Client) DescribeLogApiDispatch(request *DescribeLogApiDispatchRequest) (response *DescribeLogApiDispatchResponse, err error) {
	if request == nil {
		request = NewDescribeLogApiDispatchRequest()
	}
	response = NewDescribeLogApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOfflineExportTaskRequest() (request *DescribeOfflineExportTaskRequest) {
	request = &DescribeOfflineExportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeOfflineExportTask")
	return
}

func NewDescribeOfflineExportTaskResponse() (response *DescribeOfflineExportTaskResponse) {
	response = &DescribeOfflineExportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志离线导出任务列表
func (c *Client) DescribeOfflineExportTask(request *DescribeOfflineExportTaskRequest) (response *DescribeOfflineExportTaskResponse, err error) {
	if request == nil {
		request = NewDescribeOfflineExportTaskRequest()
	}
	response = NewDescribeOfflineExportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockTimesListRequest() (request *DescribeBlockTimesListRequest) {
	request = &DescribeBlockTimesListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockTimesList")
	return
}

func NewDescribeBlockTimesListResponse() (response *DescribeBlockTimesListResponse) {
	response = &DescribeBlockTimesListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeBlockTimesList告警中心-阻断折线图
func (c *Client) DescribeBlockTimesList(request *DescribeBlockTimesListRequest) (response *DescribeBlockTimesListResponse, err error) {
	if request == nil {
		request = NewDescribeBlockTimesListRequest()
	}
	response = NewDescribeBlockTimesListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBetaTaskDetailRequest() (request *DescribeBetaTaskDetailRequest) {
	request = &DescribeBetaTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBetaTaskDetail")
	return
}

func NewDescribeBetaTaskDetailResponse() (response *DescribeBetaTaskDetailResponse) {
	response = &DescribeBetaTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询自动化任务详情
func (c *Client) DescribeBetaTaskDetail(request *DescribeBetaTaskDetailRequest) (response *DescribeBetaTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBetaTaskDetailRequest()
	}
	response = NewDescribeBetaTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserStatusRequest() (request *DescribeUserStatusRequest) {
	request = &DescribeUserStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeUserStatus")
	return
}

func NewDescribeUserStatusResponse() (response *DescribeUserStatusResponse) {
	response = &DescribeUserStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询是否防火墙用户
func (c *Client) DescribeUserStatus(request *DescribeUserStatusRequest) (response *DescribeUserStatusResponse, err error) {
	if request == nil {
		request = NewDescribeUserStatusRequest()
	}
	response = NewDescribeUserStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogsRequest() (request *DescribeLogsRequest) {
	request = &DescribeLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeLogs")
	return
}

func NewDescribeLogsResponse() (response *DescribeLogsResponse) {
	response = &DescribeLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志审计日志查询
func (c *Client) DescribeLogs(request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
	if request == nil {
		request = NewDescribeLogsRequest()
	}
	response = NewDescribeLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStorageSettingRequest() (request *DescribeStorageSettingRequest) {
	request = &DescribeStorageSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeStorageSetting")
	return
}

func NewDescribeStorageSettingResponse() (response *DescribeStorageSettingResponse) {
	response = &DescribeStorageSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户存储配置查询接口
func (c *Client) DescribeStorageSetting(request *DescribeStorageSettingRequest) (response *DescribeStorageSettingResponse, err error) {
	if request == nil {
		request = NewDescribeStorageSettingRequest()
	}
	response = NewDescribeStorageSettingResponse()
	err = c.Send(request, response)
	return
}

func NewAddZBTiNoticeRequest() (request *AddZBTiNoticeRequest) {
	request = &AddZBTiNoticeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddZBTiNotice")
	return
}

func NewAddZBTiNoticeResponse() (response *AddZBTiNoticeResponse) {
	response = &AddZBTiNoticeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重保情报通知
func (c *Client) AddZBTiNotice(request *AddZBTiNoticeRequest) (response *AddZBTiNoticeResponse, err error) {
	if request == nil {
		request = NewAddZBTiNoticeRequest()
	}
	response = NewAddZBTiNoticeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDefenceApiDispatchRequest() (request *ModifyDefenceApiDispatchRequest) {
	request = &ModifyDefenceApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyDefenceApiDispatch")
	return
}

func NewModifyDefenceApiDispatchResponse() (response *ModifyDefenceApiDispatchResponse) {
	response = &ModifyDefenceApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 入侵防御写接口请求中转
func (c *Client) ModifyDefenceApiDispatch(request *ModifyDefenceApiDispatchRequest) (response *ModifyDefenceApiDispatchResponse, err error) {
	if request == nil {
		request = NewModifyDefenceApiDispatchRequest()
	}
	response = NewModifyDefenceApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetOverviewNewRequest() (request *DescribeAssetOverviewNewRequest) {
	request = &DescribeAssetOverviewNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssetOverviewNew")
	return
}

func NewDescribeAssetOverviewNewResponse() (response *DescribeAssetOverviewNewResponse) {
	response = &DescribeAssetOverviewNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAssetOverview 资产防护概览
func (c *Client) DescribeAssetOverviewNew(request *DescribeAssetOverviewNewRequest) (response *DescribeAssetOverviewNewResponse, err error) {
	if request == nil {
		request = NewDescribeAssetOverviewNewRequest()
	}
	response = NewDescribeAssetOverviewNewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatProtectListRequest() (request *DescribeNatProtectListRequest) {
	request = &DescribeNatProtectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatProtectList")
	return
}

func NewDescribeNatProtectListResponse() (response *DescribeNatProtectListResponse) {
	response = &DescribeNatProtectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// NAT防火墙防护模式列表
func (c *Client) DescribeNatProtectList(request *DescribeNatProtectListRequest) (response *DescribeNatProtectListResponse, err error) {
	if request == nil {
		request = NewDescribeNatProtectListRequest()
	}
	response = NewDescribeNatProtectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcSwitchStatusRequest() (request *DescribeVpcSwitchStatusRequest) {
	request = &DescribeVpcSwitchStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcSwitchStatus")
	return
}

func NewDescribeVpcSwitchStatusResponse() (response *DescribeVpcSwitchStatusResponse) {
	response = &DescribeVpcSwitchStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙开关-VPC开关状态
func (c *Client) DescribeVpcSwitchStatus(request *DescribeVpcSwitchStatusRequest) (response *DescribeVpcSwitchStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVpcSwitchStatusRequest()
	}
	response = NewDescribeVpcSwitchStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockIPBySGSwitchRequest() (request *DescribeBlockIPBySGSwitchRequest) {
	request = &DescribeBlockIPBySGSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockIPBySGSwitch")
	return
}

func NewDescribeBlockIPBySGSwitchResponse() (response *DescribeBlockIPBySGSwitchResponse) {
	response = &DescribeBlockIPBySGSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询使用安全组封禁云内ip的开关
func (c *Client) DescribeBlockIPBySGSwitch(request *DescribeBlockIPBySGSwitchRequest) (response *DescribeBlockIPBySGSwitchResponse, err error) {
	if request == nil {
		request = NewDescribeBlockIPBySGSwitchRequest()
	}
	response = NewDescribeBlockIPBySGSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskAssetsRequest() (request *DescribeRiskAssetsRequest) {
	request = &DescribeRiskAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeRiskAssets")
	return
}

func NewDescribeRiskAssetsResponse() (response *DescribeRiskAssetsResponse) {
	response = &DescribeRiskAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 替换以前的node上的DescribeRiskAssets接口
func (c *Client) DescribeRiskAssets(request *DescribeRiskAssetsRequest) (response *DescribeRiskAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeRiskAssetsRequest()
	}
	response = NewDescribeRiskAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBorderProtectListRequest() (request *DescribeBorderProtectListRequest) {
	request = &DescribeBorderProtectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBorderProtectList")
	return
}

func NewDescribeBorderProtectListResponse() (response *DescribeBorderProtectListResponse) {
	response = &DescribeBorderProtectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 互联网边界防护模式列表
func (c *Client) DescribeBorderProtectList(request *DescribeBorderProtectListRequest) (response *DescribeBorderProtectListResponse, err error) {
	if request == nil {
		request = NewDescribeBorderProtectListRequest()
	}
	response = NewDescribeBorderProtectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogsAsyncRequest() (request *DescribeLogsAsyncRequest) {
	request = &DescribeLogsAsyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeLogsAsync")
	return
}

func NewDescribeLogsAsyncResponse() (response *DescribeLogsAsyncResponse) {
	response = &DescribeLogsAsyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志审计日志异步查询
func (c *Client) DescribeLogsAsync(request *DescribeLogsAsyncRequest) (response *DescribeLogsAsyncResponse, err error) {
	if request == nil {
		request = NewDescribeLogsAsyncRequest()
	}
	response = NewDescribeLogsAsyncResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetScanRequest() (request *ModifyAssetScanRequest) {
	request = &ModifyAssetScanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyAssetScan")
	return
}

func NewModifyAssetScanResponse() (response *ModifyAssetScanResponse) {
	response = &ModifyAssetScanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产扫描
func (c *Client) ModifyAssetScan(request *ModifyAssetScanRequest) (response *ModifyAssetScanResponse, err error) {
	if request == nil {
		request = NewModifyAssetScanRequest()
	}
	response = NewModifyAssetScanResponse()
	err = c.Send(request, response)
	return
}

func NewAddIocFeedbackRequest() (request *AddIocFeedbackRequest) {
	request = &AddIocFeedbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddIocFeedback")
	return
}

func NewAddIocFeedbackResponse() (response *AddIocFeedbackResponse) {
	response = &AddIocFeedbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 误报反馈接口
func (c *Client) AddIocFeedback(request *AddIocFeedbackRequest) (response *AddIocFeedbackResponse, err error) {
	if request == nil {
		request = NewAddIocFeedbackRequest()
	}
	response = NewAddIocFeedbackResponse()
	err = c.Send(request, response)
	return
}

func NewSetNatFwEipRequest() (request *SetNatFwEipRequest) {
	request = &SetNatFwEipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "SetNatFwEip")
	return
}

func NewSetNatFwEipResponse() (response *SetNatFwEipResponse) {
	response = &SetNatFwEipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置防火墙实例弹性公网ip，目前仅支持新增模式的防火墙实例
func (c *Client) SetNatFwEip(request *SetNatFwEipRequest) (response *SetNatFwEipResponse, err error) {
	if request == nil {
		request = NewSetNatFwEipRequest()
	}
	response = NewSetNatFwEipResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfwEipsRequest() (request *DescribeCfwEipsRequest) {
	request = &DescribeCfwEipsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwEips")
	return
}

func NewDescribeCfwEipsResponse() (response *DescribeCfwEipsResponse) {
	response = &DescribeCfwEipsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防火墙弹性公网IP
func (c *Client) DescribeCfwEips(request *DescribeCfwEipsRequest) (response *DescribeCfwEipsResponse, err error) {
	if request == nil {
		request = NewDescribeCfwEipsRequest()
	}
	response = NewDescribeCfwEipsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcInstanceRequest() (request *DescribeVpcInstanceRequest) {
	request = &DescribeVpcInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcInstance")
	return
}

func NewDescribeVpcInstanceResponse() (response *DescribeVpcInstanceResponse) {
	response = &DescribeVpcInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户所有VPC防火墙实例列表
func (c *Client) DescribeVpcInstance(request *DescribeVpcInstanceRequest) (response *DescribeVpcInstanceResponse, err error) {
	if request == nil {
		request = NewDescribeVpcInstanceRequest()
	}
	response = NewDescribeVpcInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIsolateListRequest() (request *DescribeIsolateListRequest) {
	request = &DescribeIsolateListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeIsolateList")
	return
}

func NewDescribeIsolateListResponse() (response *DescribeIsolateListResponse) {
	response = &DescribeIsolateListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户配置的隔离列表
func (c *Client) DescribeIsolateList(request *DescribeIsolateListRequest) (response *DescribeIsolateListResponse, err error) {
	if request == nil {
		request = NewDescribeIsolateListRequest()
	}
	response = NewDescribeIsolateListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBorderACLListRequest() (request *DescribeBorderACLListRequest) {
	request = &DescribeBorderACLListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBorderACLList")
	return
}

func NewDescribeBorderACLListResponse() (response *DescribeBorderACLListResponse) {
	response = &DescribeBorderACLListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 互联网边界规则列表
func (c *Client) DescribeBorderACLList(request *DescribeBorderACLListRequest) (response *DescribeBorderACLListResponse, err error) {
	if request == nil {
		request = NewDescribeBorderACLListRequest()
	}
	response = NewDescribeBorderACLListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFwGroupInstanceInfoRequest() (request *DescribeFwGroupInstanceInfoRequest) {
	request = &DescribeFwGroupInstanceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFwGroupInstanceInfo")
	return
}

func NewDescribeFwGroupInstanceInfoResponse() (response *DescribeFwGroupInstanceInfoResponse) {
	response = &DescribeFwGroupInstanceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户所有VPC防火墙(组)及VPC防火墙实例卡片信息
func (c *Client) DescribeFwGroupInstanceInfo(request *DescribeFwGroupInstanceInfoRequest) (response *DescribeFwGroupInstanceInfoResponse, err error) {
	if request == nil {
		request = NewDescribeFwGroupInstanceInfoRequest()
	}
	response = NewDescribeFwGroupInstanceInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogApiDispatchRequest() (request *ModifyLogApiDispatchRequest) {
	request = &ModifyLogApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyLogApiDispatch")
	return
}

func NewModifyLogApiDispatchResponse() (response *ModifyLogApiDispatchResponse) {
	response = &ModifyLogApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志审计写接口请求中转
func (c *Client) ModifyLogApiDispatch(request *ModifyLogApiDispatchRequest) (response *ModifyLogApiDispatchResponse, err error) {
	if request == nil {
		request = NewModifyLogApiDispatchRequest()
	}
	response = NewModifyLogApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwRuleHitDetailRequest() (request *DescribeVpcFwRuleHitDetailRequest) {
	request = &DescribeVpcFwRuleHitDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwRuleHitDetail")
	return
}

func NewDescribeVpcFwRuleHitDetailResponse() (response *DescribeVpcFwRuleHitDetailResponse) {
	response = &DescribeVpcFwRuleHitDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询内网间访问控制规则命中详情
func (c *Client) DescribeVpcFwRuleHitDetail(request *DescribeVpcFwRuleHitDetailRequest) (response *DescribeVpcFwRuleHitDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwRuleHitDetailRequest()
	}
	response = NewDescribeVpcFwRuleHitDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTableStatusRequest() (request *ModifyTableStatusRequest) {
	request = &ModifyTableStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyTableStatus")
	return
}

func NewModifyTableStatusResponse() (response *ModifyTableStatusResponse) {
	response = &ModifyTableStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改规则表状态
func (c *Client) ModifyTableStatus(request *ModifyTableStatusRequest) (response *ModifyTableStatusResponse, err error) {
	if request == nil {
		request = NewModifyTableStatusRequest()
	}
	response = NewModifyTableStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatSwitchStatusRequest() (request *DescribeNatSwitchStatusRequest) {
	request = &DescribeNatSwitchStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatSwitchStatus")
	return
}

func NewDescribeNatSwitchStatusResponse() (response *DescribeNatSwitchStatusResponse) {
	response = &DescribeNatSwitchStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于定时查询 NAT 防火墙开关状态
func (c *Client) DescribeNatSwitchStatus(request *DescribeNatSwitchStatusRequest) (response *DescribeNatSwitchStatusResponse, err error) {
	if request == nil {
		request = NewDescribeNatSwitchStatusRequest()
	}
	response = NewDescribeNatSwitchStatusResponse()
	err = c.Send(request, response)
	return
}

func NewExportLogsOfflineRequest() (request *ExportLogsOfflineRequest) {
	request = &ExportLogsOfflineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ExportLogsOffline")
	return
}

func NewExportLogsOfflineResponse() (response *ExportLogsOfflineResponse) {
	response = &ExportLogsOfflineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志审计日志离线导出
func (c *Client) ExportLogsOffline(request *ExportLogsOfflineRequest) (response *ExportLogsOfflineResponse, err error) {
	if request == nil {
		request = NewExportLogsOfflineRequest()
	}
	response = NewExportLogsOfflineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNatFwInstanceWithDomainRequest() (request *CreateNatFwInstanceWithDomainRequest) {
	request = &CreateNatFwInstanceWithDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateNatFwInstanceWithDomain")
	return
}

func NewCreateNatFwInstanceWithDomainResponse() (response *CreateNatFwInstanceWithDomainResponse) {
	response = &CreateNatFwInstanceWithDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建防火墙实例和接入域名（Region参数必填）
func (c *Client) CreateNatFwInstanceWithDomain(request *CreateNatFwInstanceWithDomainRequest) (response *CreateNatFwInstanceWithDomainResponse, err error) {
	if request == nil {
		request = NewCreateNatFwInstanceWithDomainRequest()
	}
	response = NewCreateNatFwInstanceWithDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEsLogRequest() (request *DescribeEsLogRequest) {
	request = &DescribeEsLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeEsLog")
	return
}

func NewDescribeEsLogResponse() (response *DescribeEsLogResponse) {
	response = &DescribeEsLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeEsLog 分析日志查询
func (c *Client) DescribeEsLog(request *DescribeEsLogRequest) (response *DescribeEsLogResponse, err error) {
	if request == nil {
		request = NewDescribeEsLogRequest()
	}
	response = NewDescribeEsLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBetaTaskRequest() (request *DescribeBetaTaskRequest) {
	request = &DescribeBetaTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBetaTask")
	return
}

func NewDescribeBetaTaskResponse() (response *DescribeBetaTaskResponse) {
	response = &DescribeBetaTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询beta任务
func (c *Client) DescribeBetaTask(request *DescribeBetaTaskRequest) (response *DescribeBetaTaskResponse, err error) {
	if request == nil {
		request = NewDescribeBetaTaskRequest()
	}
	response = NewDescribeBetaTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnSASEEdgeRequest() (request *DescribeCcnSASEEdgeRequest) {
	request = &DescribeCcnSASEEdgeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCcnSASEEdge")
	return
}

func NewDescribeCcnSASEEdgeResponse() (response *DescribeCcnSASEEdgeResponse) {
	response = &DescribeCcnSASEEdgeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看CCN中SD-WAN的EDGE信息
func (c *Client) DescribeCcnSASEEdge(request *DescribeCcnSASEEdgeRequest) (response *DescribeCcnSASEEdgeResponse, err error) {
	if request == nil {
		request = NewDescribeCcnSASEEdgeRequest()
	}
	response = NewDescribeCcnSASEEdgeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwTcRuleRequest() (request *DescribeVpcFwTcRuleRequest) {
	request = &DescribeVpcFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwTcRule")
	return
}

func NewDescribeVpcFwTcRuleResponse() (response *DescribeVpcFwTcRuleResponse) {
	response = &DescribeVpcFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询流控策略列表
func (c *Client) DescribeVpcFwTcRule(request *DescribeVpcFwTcRuleRequest) (response *DescribeVpcFwTcRuleResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwTcRuleRequest()
	}
	response = NewDescribeVpcFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyChooseResourceGroupRequest() (request *ModifyChooseResourceGroupRequest) {
	request = &ModifyChooseResourceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyChooseResourceGroup")
	return
}

func NewModifyChooseResourceGroupResponse() (response *ModifyChooseResourceGroupResponse) {
	response = &ModifyChooseResourceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyChooseResourceGroup资产中心-资产组-资产操作
func (c *Client) ModifyChooseResourceGroup(request *ModifyChooseResourceGroupRequest) (response *ModifyChooseResourceGroupResponse, err error) {
	if request == nil {
		request = NewModifyChooseResourceGroupRequest()
	}
	response = NewModifyChooseResourceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCidrRelatedInstancesRequest() (request *DescribeCidrRelatedInstancesRequest) {
	request = &DescribeCidrRelatedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCidrRelatedInstances")
	return
}

func NewDescribeCidrRelatedInstancesResponse() (response *DescribeCidrRelatedInstancesResponse) {
	response = &DescribeCidrRelatedInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取cidr关联的网络实例
func (c *Client) DescribeCidrRelatedInstances(request *DescribeCidrRelatedInstancesRequest) (response *DescribeCidrRelatedInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeCidrRelatedInstancesRequest()
	}
	response = NewDescribeCidrRelatedInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserStorageTypeRequest() (request *DescribeUserStorageTypeRequest) {
	request = &DescribeUserStorageTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeUserStorageType")
	return
}

func NewDescribeUserStorageTypeResponse() (response *DescribeUserStorageTypeResponse) {
	response = &DescribeUserStorageTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户日志存储类型查询
func (c *Client) DescribeUserStorageType(request *DescribeUserStorageTypeRequest) (response *DescribeUserStorageTypeResponse, err error) {
	if request == nil {
		request = NewDescribeUserStorageTypeRequest()
	}
	response = NewDescribeUserStorageTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFwGroupIdNamesRequest() (request *DescribeFwGroupIdNamesRequest) {
	request = &DescribeFwGroupIdNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFwGroupIdNames")
	return
}

func NewDescribeFwGroupIdNamesResponse() (response *DescribeFwGroupIdNamesResponse) {
	response = &DescribeFwGroupIdNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户防火墙(组)的ID名称列表
func (c *Client) DescribeFwGroupIdNames(request *DescribeFwGroupIdNamesRequest) (response *DescribeFwGroupIdNamesResponse, err error) {
	if request == nil {
		request = NewDescribeFwGroupIdNamesRequest()
	}
	response = NewDescribeFwGroupIdNamesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcFwNameRequest() (request *ModifyVpcFwNameRequest) {
	request = &ModifyVpcFwNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcFwName")
	return
}

func NewModifyVpcFwNameResponse() (response *ModifyVpcFwNameResponse) {
	response = &ModifyVpcFwNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改VPC防火墙的名称
func (c *Client) ModifyVpcFwName(request *ModifyVpcFwNameRequest) (response *ModifyVpcFwNameResponse, err error) {
	if request == nil {
		request = NewModifyVpcFwNameRequest()
	}
	response = NewModifyVpcFwNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTableStatusRequest() (request *DescribeTableStatusRequest) {
	request = &DescribeTableStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeTableStatus")
	return
}

func NewDescribeTableStatusResponse() (response *DescribeTableStatusResponse) {
	response = &DescribeTableStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询规则表状态
func (c *Client) DescribeTableStatus(request *DescribeTableStatusRequest) (response *DescribeTableStatusResponse, err error) {
	if request == nil {
		request = NewDescribeTableStatusRequest()
	}
	response = NewDescribeTableStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFwSyncStatusRequest() (request *DescribeFwSyncStatusRequest) {
	request = &DescribeFwSyncStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFwSyncStatus")
	return
}

func NewDescribeFwSyncStatusResponse() (response *DescribeFwSyncStatusResponse) {
	response = &DescribeFwSyncStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防火墙同步状态，一般在执行同步操作后查询
func (c *Client) DescribeFwSyncStatus(request *DescribeFwSyncStatusRequest) (response *DescribeFwSyncStatusResponse, err error) {
	if request == nil {
		request = NewDescribeFwSyncStatusRequest()
	}
	response = NewDescribeFwSyncStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTiWhiteListRequest() (request *DescribeTiWhiteListRequest) {
	request = &DescribeTiWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeTiWhiteList")
	return
}

func NewDescribeTiWhiteListResponse() (response *DescribeTiWhiteListResponse) {
	response = &DescribeTiWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户级-情报白名单
func (c *Client) DescribeTiWhiteList(request *DescribeTiWhiteListRequest) (response *DescribeTiWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeTiWhiteListRequest()
	}
	response = NewDescribeTiWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewQueryRuleLimitsForWidthRequest() (request *QueryRuleLimitsForWidthRequest) {
	request = &QueryRuleLimitsForWidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "QueryRuleLimitsForWidth")
	return
}

func NewQueryRuleLimitsForWidthResponse() (response *QueryRuleLimitsForWidthResponse) {
	response = &QueryRuleLimitsForWidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定防火墙指定带宽下，最大支持能支持的ACL规则条数
func (c *Client) QueryRuleLimitsForWidth(request *QueryRuleLimitsForWidthRequest) (response *QueryRuleLimitsForWidthResponse, err error) {
	if request == nil {
		request = NewQueryRuleLimitsForWidthRequest()
	}
	response = NewQueryRuleLimitsForWidthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetflowBorderUsedRequest() (request *DescribeNetflowBorderUsedRequest) {
	request = &DescribeNetflowBorderUsedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNetflowBorderUsed")
	return
}

func NewDescribeNetflowBorderUsedResponse() (response *DescribeNetflowBorderUsedResponse) {
	response = &DescribeNetflowBorderUsedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 互联网边界用量超量告警
func (c *Client) DescribeNetflowBorderUsed(request *DescribeNetflowBorderUsedRequest) (response *DescribeNetflowBorderUsedResponse, err error) {
	if request == nil {
		request = NewDescribeNetflowBorderUsedRequest()
	}
	response = NewDescribeNetflowBorderUsedResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEdgeNameRequest() (request *ModifyEdgeNameRequest) {
	request = &ModifyEdgeNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyEdgeName")
	return
}

func NewModifyEdgeNameResponse() (response *ModifyEdgeNameResponse) {
	response = &ModifyEdgeNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改vpc 防火墙 edge的名称
func (c *Client) ModifyEdgeName(request *ModifyEdgeNameRequest) (response *ModifyEdgeNameResponse, err error) {
	if request == nil {
		request = NewModifyEdgeNameRequest()
	}
	response = NewModifyEdgeNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEdgeListRequest() (request *DescribeVpcEdgeListRequest) {
	request = &DescribeVpcEdgeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcEdgeList")
	return
}

func NewDescribeVpcEdgeListResponse() (response *DescribeVpcEdgeListResponse) {
	response = &DescribeVpcEdgeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看vpc防火墙边状态
func (c *Client) DescribeVpcEdgeList(request *DescribeVpcEdgeListRequest) (response *DescribeVpcEdgeListResponse, err error) {
	if request == nil {
		request = NewDescribeVpcEdgeListRequest()
	}
	response = NewDescribeVpcEdgeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFlowCenterLogsV1Request() (request *DescribeVpcFlowCenterLogsV1Request) {
	request = &DescribeVpcFlowCenterLogsV1Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFlowCenterLogsV1")
	return
}

func NewDescribeVpcFlowCenterLogsV1Response() (response *DescribeVpcFlowCenterLogsV1Response) {
	response = &DescribeVpcFlowCenterLogsV1Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Vpc流量中心列表查询
func (c *Client) DescribeVpcFlowCenterLogsV1(request *DescribeVpcFlowCenterLogsV1Request) (response *DescribeVpcFlowCenterLogsV1Response, err error) {
	if request == nil {
		request = NewDescribeVpcFlowCenterLogsV1Request()
	}
	response = NewDescribeVpcFlowCenterLogsV1Response()
	err = c.Send(request, response)
	return
}

func NewRemoveNatFwTcRuleRequest() (request *RemoveNatFwTcRuleRequest) {
	request = &RemoveNatFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "RemoveNatFwTcRule")
	return
}

func NewRemoveNatFwTcRuleResponse() (response *RemoveNatFwTcRuleResponse) {
	response = &RemoveNatFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除流量控制策略
func (c *Client) RemoveNatFwTcRule(request *RemoveNatFwTcRuleRequest) (response *RemoveNatFwTcRuleResponse, err error) {
	if request == nil {
		request = NewRemoveNatFwTcRuleRequest()
	}
	response = NewRemoveNatFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewStartUpdateResourceTaskRequest() (request *StartUpdateResourceTaskRequest) {
	request = &StartUpdateResourceTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "StartUpdateResourceTask")
	return
}

func NewStartUpdateResourceTaskResponse() (response *StartUpdateResourceTaskResponse) {
	response = &StartUpdateResourceTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动资产更新任务，任务为异步任务，没有返回值
func (c *Client) StartUpdateResourceTask(request *StartUpdateResourceTaskRequest) (response *StartUpdateResourceTaskResponse, err error) {
	if request == nil {
		request = NewStartUpdateResourceTaskRequest()
	}
	response = NewStartUpdateResourceTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGuideIdpsInfoRequest() (request *DescribeGuideIdpsInfoRequest) {
	request = &DescribeGuideIdpsInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeGuideIdpsInfo")
	return
}

func NewDescribeGuideIdpsInfoResponse() (response *DescribeGuideIdpsInfoResponse) {
	response = &DescribeGuideIdpsInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeGuideIdpsInfo 新手引导拦截模式信息查询接口
func (c *Client) DescribeGuideIdpsInfo(request *DescribeGuideIdpsInfoRequest) (response *DescribeGuideIdpsInfoResponse, err error) {
	if request == nil {
		request = NewDescribeGuideIdpsInfoRequest()
	}
	response = NewDescribeGuideIdpsInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMessageCenterSwitchRequest() (request *DescribeMessageCenterSwitchRequest) {
	request = &DescribeMessageCenterSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeMessageCenterSwitch")
	return
}

func NewDescribeMessageCenterSwitchResponse() (response *DescribeMessageCenterSwitchResponse) {
	response = &DescribeMessageCenterSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询使用消息中心订阅推送的开关
func (c *Client) DescribeMessageCenterSwitch(request *DescribeMessageCenterSwitchRequest) (response *DescribeMessageCenterSwitchResponse, err error) {
	if request == nil {
		request = NewDescribeMessageCenterSwitchRequest()
	}
	response = NewDescribeMessageCenterSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockListRequest() (request *DescribeBlockListRequest) {
	request = &DescribeBlockListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockList")
	return
}

func NewDescribeBlockListResponse() (response *DescribeBlockListResponse) {
	response = &DescribeBlockListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeBlockList 告警中心阻断资产视图列表
func (c *Client) DescribeBlockList(request *DescribeBlockListRequest) (response *DescribeBlockListResponse, err error) {
	if request == nil {
		request = NewDescribeBlockListRequest()
	}
	response = NewDescribeBlockListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRunSyncAssetRequest() (request *ModifyRunSyncAssetRequest) {
	request = &ModifyRunSyncAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyRunSyncAsset")
	return
}

func NewModifyRunSyncAssetResponse() (response *ModifyRunSyncAssetResponse) {
	response = &ModifyRunSyncAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步资产-互联网&VPC（新）
func (c *Client) ModifyRunSyncAsset(request *ModifyRunSyncAssetRequest) (response *ModifyRunSyncAssetResponse, err error) {
	if request == nil {
		request = NewModifyRunSyncAssetRequest()
	}
	response = NewModifyRunSyncAssetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVisitTimesAndFlowAssetMaxTopRequest() (request *DescribeVisitTimesAndFlowAssetMaxTopRequest) {
	request = &DescribeVisitTimesAndFlowAssetMaxTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVisitTimesAndFlowAssetMaxTop")
	return
}

func NewDescribeVisitTimesAndFlowAssetMaxTopResponse() (response *DescribeVisitTimesAndFlowAssetMaxTopResponse) {
	response = &DescribeVisitTimesAndFlowAssetMaxTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量中心资产Top查询
func (c *Client) DescribeVisitTimesAndFlowAssetMaxTop(request *DescribeVisitTimesAndFlowAssetMaxTopRequest) (response *DescribeVisitTimesAndFlowAssetMaxTopResponse, err error) {
	if request == nil {
		request = NewDescribeVisitTimesAndFlowAssetMaxTopRequest()
	}
	response = NewDescribeVisitTimesAndFlowAssetMaxTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwViewStatRequest() (request *DescribeVpcFwViewStatRequest) {
	request = &DescribeVpcFwViewStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwViewStat")
	return
}

func NewDescribeVpcFwViewStatResponse() (response *DescribeVpcFwViewStatResponse) {
	response = &DescribeVpcFwViewStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VPC防火墙概览页信息
func (c *Client) DescribeVpcFwViewStat(request *DescribeVpcFwViewStatRequest) (response *DescribeVpcFwViewStatResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwViewStatRequest()
	}
	response = NewDescribeVpcFwViewStatResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIPSProtectModeRequest() (request *ModifyIPSProtectModeRequest) {
	request = &ModifyIPSProtectModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyIPSProtectMode")
	return
}

func NewModifyIPSProtectModeResponse() (response *ModifyIPSProtectModeResponse) {
	response = &ModifyIPSProtectModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyIPSProtectMode
func (c *Client) ModifyIPSProtectMode(request *ModifyIPSProtectModeRequest) (response *ModifyIPSProtectModeResponse, err error) {
	if request == nil {
		request = NewModifyIPSProtectModeRequest()
	}
	response = NewModifyIPSProtectModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogsCountAsyncRequest() (request *DescribeLogsCountAsyncRequest) {
	request = &DescribeLogsCountAsyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeLogsCountAsync")
	return
}

func NewDescribeLogsCountAsyncResponse() (response *DescribeLogsCountAsyncResponse) {
	response = &DescribeLogsCountAsyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志审计日志异步查询结果总数获取
func (c *Client) DescribeLogsCountAsync(request *DescribeLogsCountAsyncRequest) (response *DescribeLogsCountAsyncResponse, err error) {
	if request == nil {
		request = NewDescribeLogsCountAsyncRequest()
	}
	response = NewDescribeLogsCountAsyncResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwSwitchRequest() (request *DescribeVpcFwSwitchRequest) {
	request = &DescribeVpcFwSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwSwitch")
	return
}

func NewDescribeVpcFwSwitchResponse() (response *DescribeVpcFwSwitchResponse) {
	response = &DescribeVpcFwSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeVpcFwSwitch
func (c *Client) DescribeVpcFwSwitch(request *DescribeVpcFwSwitchRequest) (response *DescribeVpcFwSwitchResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwSwitchRequest()
	}
	response = NewDescribeVpcFwSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewModifySwitchStatusRequest() (request *ModifySwitchStatusRequest) {
	request = &ModifySwitchStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifySwitchStatus")
	return
}

func NewModifySwitchStatusResponse() (response *ModifySwitchStatusResponse) {
	response = &ModifySwitchStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防火墙开关状态
func (c *Client) ModifySwitchStatus(request *ModifySwitchStatusRequest) (response *ModifySwitchStatusResponse, err error) {
	if request == nil {
		request = NewModifySwitchStatusRequest()
	}
	response = NewModifySwitchStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQueryNotEmptyRuleListInfoRequest() (request *DescribeQueryNotEmptyRuleListInfoRequest) {
	request = &DescribeQueryNotEmptyRuleListInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeQueryNotEmptyRuleListInfo")
	return
}

func NewDescribeQueryNotEmptyRuleListInfoResponse() (response *DescribeQueryNotEmptyRuleListInfoResponse) {
	response = &DescribeQueryNotEmptyRuleListInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取备份规则
func (c *Client) DescribeQueryNotEmptyRuleListInfo(request *DescribeQueryNotEmptyRuleListInfoRequest) (response *DescribeQueryNotEmptyRuleListInfoResponse, err error) {
	if request == nil {
		request = NewDescribeQueryNotEmptyRuleListInfoRequest()
	}
	response = NewDescribeQueryNotEmptyRuleListInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePcapTaskRequest() (request *DeletePcapTaskRequest) {
	request = &DeletePcapTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeletePcapTask")
	return
}

func NewDeletePcapTaskResponse() (response *DeletePcapTaskResponse) {
	response = &DeletePcapTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除防火墙抓包任务
func (c *Client) DeletePcapTask(request *DeletePcapTaskRequest) (response *DeletePcapTaskResponse, err error) {
	if request == nil {
		request = NewDeletePcapTaskRequest()
	}
	response = NewDeletePcapTaskResponse()
	err = c.Send(request, response)
	return
}

func NewAddNewBindRoleUserRequest() (request *AddNewBindRoleUserRequest) {
	request = &AddNewBindRoleUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddNewBindRoleUser")
	return
}

func NewAddNewBindRoleUserResponse() (response *AddNewBindRoleUserResponse) {
	response = &AddNewBindRoleUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙角色授权绑定接口
func (c *Client) AddNewBindRoleUser(request *AddNewBindRoleUserRequest) (response *AddNewBindRoleUserResponse, err error) {
	if request == nil {
		request = NewAddNewBindRoleUserRequest()
	}
	response = NewAddNewBindRoleUserResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLoginTimeRequest() (request *ModifyLoginTimeRequest) {
	request = &ModifyLoginTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyLoginTime")
	return
}

func NewModifyLoginTimeResponse() (response *ModifyLoginTimeResponse) {
	response = &ModifyLoginTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新登录时间(基于subuin)，用于判断新增资产
func (c *Client) ModifyLoginTime(request *ModifyLoginTimeRequest) (response *ModifyLoginTimeResponse, err error) {
	if request == nil {
		request = NewModifyLoginTimeRequest()
	}
	response = NewModifyLoginTimeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPublicIPSwitchStatusRequest() (request *ModifyPublicIPSwitchStatusRequest) {
	request = &ModifyPublicIPSwitchStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyPublicIPSwitchStatus")
	return
}

func NewModifyPublicIPSwitchStatusResponse() (response *ModifyPublicIPSwitchStatusResponse) {
	response = &ModifyPublicIPSwitchStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 单个修改互联网边界防火墙开关
func (c *Client) ModifyPublicIPSwitchStatus(request *ModifyPublicIPSwitchStatusRequest) (response *ModifyPublicIPSwitchStatusResponse, err error) {
	if request == nil {
		request = NewModifyPublicIPSwitchStatusRequest()
	}
	response = NewModifyPublicIPSwitchStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveNatAcRuleRequest() (request *RemoveNatAcRuleRequest) {
	request = &RemoveNatAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "RemoveNatAcRule")
	return
}

func NewRemoveNatAcRuleResponse() (response *RemoveNatAcRuleResponse) {
	response = &RemoveNatAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除NAT访问控制规则
func (c *Client) RemoveNatAcRule(request *RemoveNatAcRuleRequest) (response *RemoveNatAcRuleResponse, err error) {
	if request == nil {
		request = NewRemoveNatAcRuleRequest()
	}
	response = NewRemoveNatAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoInternetSwitchRequest() (request *DescribeAutoInternetSwitchRequest) {
	request = &DescribeAutoInternetSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAutoInternetSwitch")
	return
}

func NewDescribeAutoInternetSwitchResponse() (response *DescribeAutoInternetSwitchResponse) {
	response = &DescribeAutoInternetSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取自动开启互联网边界开关
func (c *Client) DescribeAutoInternetSwitch(request *DescribeAutoInternetSwitchRequest) (response *DescribeAutoInternetSwitchResponse, err error) {
	if request == nil {
		request = NewDescribeAutoInternetSwitchRequest()
	}
	response = NewDescribeAutoInternetSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShowBakRuleListRequest() (request *DescribeShowBakRuleListRequest) {
	request = &DescribeShowBakRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeShowBakRuleList")
	return
}

func NewDescribeShowBakRuleListResponse() (response *DescribeShowBakRuleListResponse) {
	response = &DescribeShowBakRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取备份规则列表
func (c *Client) DescribeShowBakRuleList(request *DescribeShowBakRuleListRequest) (response *DescribeShowBakRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeShowBakRuleListRequest()
	}
	response = NewDescribeShowBakRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwGroupFlowStatRequest() (request *DescribeVpcFwGroupFlowStatRequest) {
	request = &DescribeVpcFwGroupFlowStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwGroupFlowStat")
	return
}

func NewDescribeVpcFwGroupFlowStatResponse() (response *DescribeVpcFwGroupFlowStatResponse) {
	response = &DescribeVpcFwGroupFlowStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// VPC防火墙状态监控TAB页内容
func (c *Client) DescribeVpcFwGroupFlowStat(request *DescribeVpcFwGroupFlowStatRequest) (response *DescribeVpcFwGroupFlowStatResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwGroupFlowStatRequest()
	}
	response = NewDescribeVpcFwGroupFlowStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpsRuleListNewRequest() (request *DescribeIpsRuleListNewRequest) {
	request = &DescribeIpsRuleListNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeIpsRuleListNew")
	return
}

func NewDescribeIpsRuleListNewResponse() (response *DescribeIpsRuleListNewResponse) {
	response = &DescribeIpsRuleListNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeIpsRuleListNew
func (c *Client) DescribeIpsRuleListNew(request *DescribeIpsRuleListNewRequest) (response *DescribeIpsRuleListNewResponse, err error) {
	if request == nil {
		request = NewDescribeIpsRuleListNewRequest()
	}
	response = NewDescribeIpsRuleListNewResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcFwGroupRequest() (request *DeleteVpcFwGroupRequest) {
	request = &DeleteVpcFwGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteVpcFwGroup")
	return
}

func NewDeleteVpcFwGroupResponse() (response *DeleteVpcFwGroupResponse) {
	response = &DeleteVpcFwGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除防火墙(组)，或者删除其中实例
func (c *Client) DeleteVpcFwGroup(request *DeleteVpcFwGroupRequest) (response *DeleteVpcFwGroupResponse, err error) {
	if request == nil {
		request = NewDeleteVpcFwGroupRequest()
	}
	response = NewDeleteVpcFwGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatFwInstanceWithRegionRequest() (request *DescribeNatFwInstanceWithRegionRequest) {
	request = &DescribeNatFwInstanceWithRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInstanceWithRegion")
	return
}

func NewDescribeNatFwInstanceWithRegionResponse() (response *DescribeNatFwInstanceWithRegionResponse) {
	response = &DescribeNatFwInstanceWithRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetNatFwInstanceWithRegion 获取租户新增运维的NAT实例，带上地域
func (c *Client) DescribeNatFwInstanceWithRegion(request *DescribeNatFwInstanceWithRegionRequest) (response *DescribeNatFwInstanceWithRegionResponse, err error) {
	if request == nil {
		request = NewDescribeNatFwInstanceWithRegionRequest()
	}
	response = NewDescribeNatFwInstanceWithRegionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAclApiDispatchRequest() (request *ModifyAclApiDispatchRequest) {
	request = &ModifyAclApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyAclApiDispatch")
	return
}

func NewModifyAclApiDispatchResponse() (response *ModifyAclApiDispatchResponse) {
	response = &ModifyAclApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 访问控制写接口请求中转
func (c *Client) ModifyAclApiDispatch(request *ModifyAclApiDispatchRequest) (response *ModifyAclApiDispatchResponse, err error) {
	if request == nil {
		request = NewModifyAclApiDispatchRequest()
	}
	response = NewModifyAclApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDbOverviewRequest() (request *DescribeDbOverviewRequest) {
	request = &DescribeDbOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeDbOverview")
	return
}

func NewDescribeDbOverviewResponse() (response *DescribeDbOverviewResponse) {
	response = &DescribeDbOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 3.3.0版本数据库页面的概览数据，包含数据库总数，未防护数，类型比例等
func (c *Client) DescribeDbOverview(request *DescribeDbOverviewRequest) (response *DescribeDbOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeDbOverviewRequest()
	}
	response = NewDescribeDbOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTLogInfoRequest() (request *DescribeTLogInfoRequest) {
	request = &DescribeTLogInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeTLogInfo")
	return
}

func NewDescribeTLogInfoResponse() (response *DescribeTLogInfoResponse) {
	response = &DescribeTLogInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTLogInfo告警中心概况
func (c *Client) DescribeTLogInfo(request *DescribeTLogInfoRequest) (response *DescribeTLogInfoResponse, err error) {
	if request == nil {
		request = NewDescribeTLogInfoRequest()
	}
	response = NewDescribeTLogInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetWebScanRequest() (request *ModifyAssetWebScanRequest) {
	request = &ModifyAssetWebScanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyAssetWebScan")
	return
}

func NewModifyAssetWebScanResponse() (response *ModifyAssetWebScanResponse) {
	response = &ModifyAssetWebScanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产和Web扫描
func (c *Client) ModifyAssetWebScan(request *ModifyAssetWebScanRequest) (response *ModifyAssetWebScanResponse, err error) {
	if request == nil {
		request = NewModifyAssetWebScanRequest()
	}
	response = NewModifyAssetWebScanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeCFWSwitchRequest() (request *DescribeEdgeCFWSwitchRequest) {
	request = &DescribeEdgeCFWSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeEdgeCFWSwitch")
	return
}

func NewDescribeEdgeCFWSwitchResponse() (response *DescribeEdgeCFWSwitchResponse) {
	response = &DescribeEdgeCFWSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看是否允许使用sd-wan云防火墙
func (c *Client) DescribeEdgeCFWSwitch(request *DescribeEdgeCFWSwitchRequest) (response *DescribeEdgeCFWSwitchResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeCFWSwitchRequest()
	}
	response = NewDescribeEdgeCFWSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGuideUserStatusRequest() (request *ModifyGuideUserStatusRequest) {
	request = &ModifyGuideUserStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyGuideUserStatus")
	return
}

func NewModifyGuideUserStatusResponse() (response *ModifyGuideUserStatusResponse) {
	response = &ModifyGuideUserStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyGuideUserStatus修改引导用户完成卡片状态位
func (c *Client) ModifyGuideUserStatus(request *ModifyGuideUserStatusRequest) (response *ModifyGuideUserStatusResponse, err error) {
	if request == nil {
		request = NewModifyGuideUserStatusRequest()
	}
	response = NewModifyGuideUserStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBuyPageInfoRequest() (request *DescribeBuyPageInfoRequest) {
	request = &DescribeBuyPageInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBuyPageInfo")
	return
}

func NewDescribeBuyPageInfoResponse() (response *DescribeBuyPageInfoResponse) {
	response = &DescribeBuyPageInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防火墙事件漏洞情况统计
func (c *Client) DescribeBuyPageInfo(request *DescribeBuyPageInfoRequest) (response *DescribeBuyPageInfoResponse, err error) {
	if request == nil {
		request = NewDescribeBuyPageInfoRequest()
	}
	response = NewDescribeBuyPageInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatFwSwitchRequest() (request *ModifyNatFwSwitchRequest) {
	request = &ModifyNatFwSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwSwitch")
	return
}

func NewModifyNatFwSwitchResponse() (response *ModifyNatFwSwitchResponse) {
	response = &ModifyNatFwSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改NAT防火墙开关
func (c *Client) ModifyNatFwSwitch(request *ModifyNatFwSwitchRequest) (response *ModifyNatFwSwitchResponse, err error) {
	if request == nil {
		request = NewModifyNatFwSwitchRequest()
	}
	response = NewModifyNatFwSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatFwTcRuleRequest() (request *ModifyNatFwTcRuleRequest) {
	request = &ModifyNatFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwTcRule")
	return
}

func NewModifyNatFwTcRuleResponse() (response *ModifyNatFwTcRuleResponse) {
	response = &ModifyNatFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑流量控制策略
func (c *Client) ModifyNatFwTcRule(request *ModifyNatFwTcRuleRequest) (response *ModifyNatFwTcRuleResponse, err error) {
	if request == nil {
		request = NewModifyNatFwTcRuleRequest()
	}
	response = NewModifyNatFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowTimesRequest() (request *DescribeFlowTimesRequest) {
	request = &DescribeFlowTimesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFlowTimes")
	return
}

func NewDescribeFlowTimesResponse() (response *DescribeFlowTimesResponse) {
	response = &DescribeFlowTimesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeFlowTimes 概览页网络访问趋势图
func (c *Client) DescribeFlowTimes(request *DescribeFlowTimesRequest) (response *DescribeFlowTimesResponse, err error) {
	if request == nil {
		request = NewDescribeFlowTimesRequest()
	}
	response = NewDescribeFlowTimesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneRequest() (request *DescribeZoneRequest) {
	request = &DescribeZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeZone")
	return
}

func NewDescribeZoneResponse() (response *DescribeZoneResponse) {
	response = &DescribeZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回指定地域的防火墙可用区列表
func (c *Client) DescribeZone(request *DescribeZoneRequest) (response *DescribeZoneResponse, err error) {
	if request == nil {
		request = NewDescribeZoneRequest()
	}
	response = NewDescribeZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCheckCLSStatusRequest() (request *DescribeCheckCLSStatusRequest) {
	request = &DescribeCheckCLSStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCheckCLSStatus")
	return
}

func NewDescribeCheckCLSStatusResponse() (response *DescribeCheckCLSStatusResponse) {
	response = &DescribeCheckCLSStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查当前用户是否开通CLS服务
func (c *Client) DescribeCheckCLSStatus(request *DescribeCheckCLSStatusRequest) (response *DescribeCheckCLSStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCheckCLSStatusRequest()
	}
	response = NewDescribeCheckCLSStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetflowRuleStatusRequest() (request *ModifyNetflowRuleStatusRequest) {
	request = &ModifyNetflowRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyNetflowRuleStatus")
	return
}

func NewModifyNetflowRuleStatusResponse() (response *ModifyNetflowRuleStatusResponse) {
	response = &ModifyNetflowRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用停用单条互联网边界规则
func (c *Client) ModifyNetflowRuleStatus(request *ModifyNetflowRuleStatusRequest) (response *ModifyNetflowRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyNetflowRuleStatusRequest()
	}
	response = NewModifyNetflowRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteResourceGroupRequest() (request *DeleteResourceGroupRequest) {
	request = &DeleteResourceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteResourceGroup")
	return
}

func NewDeleteResourceGroupResponse() (response *DeleteResourceGroupResponse) {
	response = &DeleteResourceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteResourceGroup-资产中心资产组删除
func (c *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
	if request == nil {
		request = NewDeleteResourceGroupRequest()
	}
	response = NewDeleteResourceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnHandleEventTabListRequest() (request *DescribeUnHandleEventTabListRequest) {
	request = &DescribeUnHandleEventTabListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeUnHandleEventTabList")
	return
}

func NewDescribeUnHandleEventTabListResponse() (response *DescribeUnHandleEventTabListResponse) {
	response = &DescribeUnHandleEventTabListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeUnHandleEventTabList 告警中心伪攻击链事件未处置接口
func (c *Client) DescribeUnHandleEventTabList(request *DescribeUnHandleEventTabListRequest) (response *DescribeUnHandleEventTabListResponse, err error) {
	if request == nil {
		request = NewDescribeUnHandleEventTabListRequest()
	}
	response = NewDescribeUnHandleEventTabListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGuideIdpsInfoRequest() (request *CreateGuideIdpsInfoRequest) {
	request = &CreateGuideIdpsInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateGuideIdpsInfo")
	return
}

func NewCreateGuideIdpsInfoResponse() (response *CreateGuideIdpsInfoResponse) {
	response = &CreateGuideIdpsInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateGuideIdpsInfo 新手引导开启拦截模式接口
func (c *Client) CreateGuideIdpsInfo(request *CreateGuideIdpsInfoRequest) (response *CreateGuideIdpsInfoResponse, err error) {
	if request == nil {
		request = NewCreateGuideIdpsInfoRequest()
	}
	response = NewCreateGuideIdpsInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSelectAssetGroupRequest() (request *DescribeSelectAssetGroupRequest) {
	request = &DescribeSelectAssetGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeSelectAssetGroup")
	return
}

func NewDescribeSelectAssetGroupResponse() (response *DescribeSelectAssetGroupResponse) {
	response = &DescribeSelectAssetGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeSelectAssetGroup资产组下的资产信息查询
func (c *Client) DescribeSelectAssetGroup(request *DescribeSelectAssetGroupRequest) (response *DescribeSelectAssetGroupResponse, err error) {
	if request == nil {
		request = NewDescribeSelectAssetGroupRequest()
	}
	response = NewDescribeSelectAssetGroupResponse()
	err = c.Send(request, response)
	return
}

func NewResetVpcRuleHitTimesRequest() (request *ResetVpcRuleHitTimesRequest) {
	request = &ResetVpcRuleHitTimesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ResetVpcRuleHitTimes")
	return
}

func NewResetVpcRuleHitTimesResponse() (response *ResetVpcRuleHitTimesResponse) {
	response = &ResetVpcRuleHitTimesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置vpc内网规则的命中次数
func (c *Client) ResetVpcRuleHitTimes(request *ResetVpcRuleHitTimesRequest) (response *ResetVpcRuleHitTimesResponse, err error) {
	if request == nil {
		request = NewResetVpcRuleHitTimesRequest()
	}
	response = NewResetVpcRuleHitTimesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatFwVpcDnsLstRequest() (request *DescribeNatFwVpcDnsLstRequest) {
	request = &DescribeNatFwVpcDnsLstRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwVpcDnsLst")
	return
}

func NewDescribeNatFwVpcDnsLstResponse() (response *DescribeNatFwVpcDnsLstResponse) {
	response = &DescribeNatFwVpcDnsLstResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示当前natfw 实例对应的vpc dns开关
func (c *Client) DescribeNatFwVpcDnsLst(request *DescribeNatFwVpcDnsLstRequest) (response *DescribeNatFwVpcDnsLstResponse, err error) {
	if request == nil {
		request = NewDescribeNatFwVpcDnsLstRequest()
	}
	response = NewDescribeNatFwVpcDnsLstResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBlockIgnoreRuleRequest() (request *ModifyBlockIgnoreRuleRequest) {
	request = &ModifyBlockIgnoreRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockIgnoreRule")
	return
}

func NewModifyBlockIgnoreRuleResponse() (response *ModifyBlockIgnoreRuleResponse) {
	response = &ModifyBlockIgnoreRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑单条入侵防御封禁列表、放通列表规则
func (c *Client) ModifyBlockIgnoreRule(request *ModifyBlockIgnoreRuleRequest) (response *ModifyBlockIgnoreRuleResponse, err error) {
	if request == nil {
		request = NewModifyBlockIgnoreRuleRequest()
	}
	response = NewModifyBlockIgnoreRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePcapTaskRequest() (request *DescribePcapTaskRequest) {
	request = &DescribePcapTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribePcapTask")
	return
}

func NewDescribePcapTaskResponse() (response *DescribePcapTaskResponse) {
	response = &DescribePcapTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防火墙抓包任务信息
func (c *Client) DescribePcapTask(request *DescribePcapTaskRequest) (response *DescribePcapTaskResponse, err error) {
	if request == nil {
		request = NewDescribePcapTaskRequest()
	}
	response = NewDescribePcapTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBlockIgnoreRuleListRequest() (request *CreateBlockIgnoreRuleListRequest) {
	request = &CreateBlockIgnoreRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateBlockIgnoreRuleList")
	return
}

func NewCreateBlockIgnoreRuleListResponse() (response *CreateBlockIgnoreRuleListResponse) {
	response = &CreateBlockIgnoreRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量添加入侵防御封禁列表、放通列表规则
func (c *Client) CreateBlockIgnoreRuleList(request *CreateBlockIgnoreRuleListRequest) (response *CreateBlockIgnoreRuleListResponse, err error) {
	if request == nil {
		request = NewCreateBlockIgnoreRuleListRequest()
	}
	response = NewCreateBlockIgnoreRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewAddVpcAcRuleRequest() (request *AddVpcAcRuleRequest) {
	request = &AddVpcAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddVpcAcRule")
	return
}

func NewAddVpcAcRuleResponse() (response *AddVpcAcRuleResponse) {
	response = &AddVpcAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加VPC内网间规则
func (c *Client) AddVpcAcRule(request *AddVpcAcRuleRequest) (response *AddVpcAcRuleResponse, err error) {
	if request == nil {
		request = NewAddVpcAcRuleRequest()
	}
	response = NewAddVpcAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFwGroupSwitchRequest() (request *ModifyFwGroupSwitchRequest) {
	request = &ModifyFwGroupSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyFwGroupSwitch")
	return
}

func NewModifyFwGroupSwitchResponse() (response *ModifyFwGroupSwitchResponse) {
	response = &ModifyFwGroupSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改防火墙(组)开关(支持单点模式、多点模式、全互通模式)
func (c *Client) ModifyFwGroupSwitch(request *ModifyFwGroupSwitchRequest) (response *ModifyFwGroupSwitchResponse, err error) {
	if request == nil {
		request = NewModifyFwGroupSwitchRequest()
	}
	response = NewModifyFwGroupSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveNatFwObjRequest() (request *RemoveNatFwObjRequest) {
	request = &RemoveNatFwObjRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "RemoveNatFwObj")
	return
}

func NewRemoveNatFwObjResponse() (response *RemoveNatFwObjResponse) {
	response = &RemoveNatFwObjResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 移除NAT防火墙vpc或nat接入对象
func (c *Client) RemoveNatFwObj(request *RemoveNatFwObjRequest) (response *RemoveNatFwObjResponse, err error) {
	if request == nil {
		request = NewRemoveNatFwObjRequest()
	}
	response = NewRemoveNatFwObjResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePcapTaskRequest() (request *CreatePcapTaskRequest) {
	request = &CreatePcapTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreatePcapTask")
	return
}

func NewCreatePcapTaskResponse() (response *CreatePcapTaskResponse) {
	response = &CreatePcapTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建防火墙抓包任务
func (c *Client) CreatePcapTask(request *CreatePcapTaskRequest) (response *CreatePcapTaskResponse, err error) {
	if request == nil {
		request = NewCreatePcapTaskRequest()
	}
	response = NewCreatePcapTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAllRuleStatusRequest() (request *ModifyAllRuleStatusRequest) {
	request = &ModifyAllRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllRuleStatus")
	return
}

func NewModifyAllRuleStatusResponse() (response *ModifyAllRuleStatusResponse) {
	response = &ModifyAllRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用停用全部规则
func (c *Client) ModifyAllRuleStatus(request *ModifyAllRuleStatusRequest) (response *ModifyAllRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyAllRuleStatusRequest()
	}
	response = NewModifyAllRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfwLimitsRequest() (request *DescribeCfwLimitsRequest) {
	request = &DescribeCfwLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwLimits")
	return
}

func NewDescribeCfwLimitsResponse() (response *DescribeCfwLimitsResponse) {
	response = &DescribeCfwLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户防火墙相关配额
func (c *Client) DescribeCfwLimits(request *DescribeCfwLimitsRequest) (response *DescribeCfwLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeCfwLimitsRequest()
	}
	response = NewDescribeCfwLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetflowCenterTrendsRequest() (request *DescribeNetflowCenterTrendsRequest) {
	request = &DescribeNetflowCenterTrendsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNetflowCenterTrends")
	return
}

func NewDescribeNetflowCenterTrendsResponse() (response *DescribeNetflowCenterTrendsResponse) {
	response = &DescribeNetflowCenterTrendsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量中心流量折线图查询
func (c *Client) DescribeNetflowCenterTrends(request *DescribeNetflowCenterTrendsRequest) (response *DescribeNetflowCenterTrendsResponse, err error) {
	if request == nil {
		request = NewDescribeNetflowCenterTrendsRequest()
	}
	response = NewDescribeNetflowCenterTrendsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDefenceApiDispatchRequest() (request *DescribeDefenceApiDispatchRequest) {
	request = &DescribeDefenceApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeDefenceApiDispatch")
	return
}

func NewDescribeDefenceApiDispatchResponse() (response *DescribeDefenceApiDispatchResponse) {
	response = &DescribeDefenceApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 入侵防御读接口请求中转
func (c *Client) DescribeDefenceApiDispatch(request *DescribeDefenceApiDispatchRequest) (response *DescribeDefenceApiDispatchResponse, err error) {
	if request == nil {
		request = NewDescribeDefenceApiDispatchRequest()
	}
	response = NewDescribeDefenceApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDefenseSwitchRequest() (request *DescribeDefenseSwitchRequest) {
	request = &DescribeDefenseSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeDefenseSwitch")
	return
}

func NewDescribeDefenseSwitchResponse() (response *DescribeDefenseSwitchResponse) {
	response = &DescribeDefenseSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取入侵防御按钮列表
func (c *Client) DescribeDefenseSwitch(request *DescribeDefenseSwitchRequest) (response *DescribeDefenseSwitchResponse, err error) {
	if request == nil {
		request = NewDescribeDefenseSwitchRequest()
	}
	response = NewDescribeDefenseSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewCheckNatFwTcRuleRequest() (request *CheckNatFwTcRuleRequest) {
	request = &CheckNatFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CheckNatFwTcRule")
	return
}

func NewCheckNatFwTcRuleResponse() (response *CheckNatFwTcRuleResponse) {
	response = &CheckNatFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查流量控制规则是否有冲突
func (c *Client) CheckNatFwTcRule(request *CheckNatFwTcRuleRequest) (response *CheckNatFwTcRuleResponse, err error) {
	if request == nil {
		request = NewCheckNatFwTcRuleRequest()
	}
	response = NewCheckNatFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnWithEdgeRequest() (request *DescribeCcnWithEdgeRequest) {
	request = &DescribeCcnWithEdgeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCcnWithEdge")
	return
}

func NewDescribeCcnWithEdgeResponse() (response *DescribeCcnWithEdgeResponse) {
	response = &DescribeCcnWithEdgeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回接入了sd-wan的CCN列表
func (c *Client) DescribeCcnWithEdge(request *DescribeCcnWithEdgeRequest) (response *DescribeCcnWithEdgeResponse, err error) {
	if request == nil {
		request = NewDescribeCcnWithEdgeRequest()
	}
	response = NewDescribeCcnWithEdgeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetFlowDomainTopRequest() (request *DescribeNetFlowDomainTopRequest) {
	request = &DescribeNetFlowDomainTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNetFlowDomainTop")
	return
}

func NewDescribeNetFlowDomainTopResponse() (response *DescribeNetFlowDomainTopResponse) {
	response = &DescribeNetFlowDomainTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主动外联域名流量top5
func (c *Client) DescribeNetFlowDomainTop(request *DescribeNetFlowDomainTopRequest) (response *DescribeNetFlowDomainTopResponse, err error) {
	if request == nil {
		request = NewDescribeNetFlowDomainTopRequest()
	}
	response = NewDescribeNetFlowDomainTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetInfoRequest() (request *DescribeNetInfoRequest) {
	request = &DescribeNetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNetInfo")
	return
}

func NewDescribeNetInfoResponse() (response *DescribeNetInfoResponse) {
	response = &DescribeNetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeNetInfo 概览网络攻击汇总
func (c *Client) DescribeNetInfo(request *DescribeNetInfoRequest) (response *DescribeNetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeNetInfoRequest()
	}
	response = NewDescribeNetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInternetOutOverviewRequest() (request *DescribeInternetOutOverviewRequest) {
	request = &DescribeInternetOutOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeInternetOutOverview")
	return
}

func NewDescribeInternetOutOverviewResponse() (response *DescribeInternetOutOverviewResponse) {
	response = &DescribeInternetOutOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量中心外联概况查询
func (c *Client) DescribeInternetOutOverview(request *DescribeInternetOutOverviewRequest) (response *DescribeInternetOutOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeInternetOutOverviewRequest()
	}
	response = NewDescribeInternetOutOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeModuleConfigRequest() (request *DescribeModuleConfigRequest) {
	request = &DescribeModuleConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeModuleConfig")
	return
}

func NewDescribeModuleConfigResponse() (response *DescribeModuleConfigResponse) {
	response = &DescribeModuleConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙模块基础配置
func (c *Client) DescribeModuleConfig(request *DescribeModuleConfigRequest) (response *DescribeModuleConfigResponse, err error) {
	if request == nil {
		request = NewDescribeModuleConfigRequest()
	}
	response = NewDescribeModuleConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCosBucketListRequest() (request *DescribeCosBucketListRequest) {
	request = &DescribeCosBucketListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCosBucketList")
	return
}

func NewDescribeCosBucketListResponse() (response *DescribeCosBucketListResponse) {
	response = &DescribeCosBucketListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Cos存储桶列表
func (c *Client) DescribeCosBucketList(request *DescribeCosBucketListRequest) (response *DescribeCosBucketListResponse, err error) {
	if request == nil {
		request = NewDescribeCosBucketListRequest()
	}
	response = NewDescribeCosBucketListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcCfwWidthRequest() (request *ModifyVpcCfwWidthRequest) {
	request = &ModifyVpcCfwWidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcCfwWidth")
	return
}

func NewModifyVpcCfwWidthResponse() (response *ModifyVpcCfwWidthResponse) {
	response = &ModifyVpcCfwWidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// vpc间防火墙垂直扩容
func (c *Client) ModifyVpcCfwWidth(request *ModifyVpcCfwWidthRequest) (response *ModifyVpcCfwWidthResponse, err error) {
	if request == nil {
		request = NewModifyVpcCfwWidthRequest()
	}
	response = NewModifyVpcCfwWidthResponse()
	err = c.Send(request, response)
	return
}

func NewQueryVpcFwSupportSwitchModeRequest() (request *QueryVpcFwSupportSwitchModeRequest) {
	request = &QueryVpcFwSupportSwitchModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "QueryVpcFwSupportSwitchMode")
	return
}

func NewQueryVpcFwSupportSwitchModeResponse() (response *QueryVpcFwSupportSwitchModeResponse) {
	response = &QueryVpcFwSupportSwitchModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VPC间防火墙可支持的开关模式，云联网模式的防火墙实例根据实际情况可能仅支持同开同关模式
func (c *Client) QueryVpcFwSupportSwitchMode(request *QueryVpcFwSupportSwitchModeRequest) (response *QueryVpcFwSupportSwitchModeResponse, err error) {
	if request == nil {
		request = NewQueryVpcFwSupportSwitchModeRequest()
	}
	response = NewQueryVpcFwSupportSwitchModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSwitchCommonRequest() (request *DescribeSwitchCommonRequest) {
	request = &DescribeSwitchCommonRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeSwitchCommon")
	return
}

func NewDescribeSwitchCommonResponse() (response *DescribeSwitchCommonResponse) {
	response = &DescribeSwitchCommonResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙开关公共接口
func (c *Client) DescribeSwitchCommon(request *DescribeSwitchCommonRequest) (response *DescribeSwitchCommonResponse, err error) {
	if request == nil {
		request = NewDescribeSwitchCommonRequest()
	}
	response = NewDescribeSwitchCommonResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportLogsRequest() (request *DescribeExportLogsRequest) {
	request = &DescribeExportLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeExportLogs")
	return
}

func NewDescribeExportLogsResponse() (response *DescribeExportLogsResponse) {
	response = &DescribeExportLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志审计日志导出
func (c *Client) DescribeExportLogs(request *DescribeExportLogsRequest) (response *DescribeExportLogsResponse, err error) {
	if request == nil {
		request = NewDescribeExportLogsRequest()
	}
	response = NewDescribeExportLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowCenterAssetListsRequest() (request *DescribeFlowCenterAssetListsRequest) {
	request = &DescribeFlowCenterAssetListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFlowCenterAssetLists")
	return
}

func NewDescribeFlowCenterAssetListsResponse() (response *DescribeFlowCenterAssetListsResponse) {
	response = &DescribeFlowCenterAssetListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量中心外联资产列表查询
func (c *Client) DescribeFlowCenterAssetLists(request *DescribeFlowCenterAssetListsRequest) (response *DescribeFlowCenterAssetListsResponse, err error) {
	if request == nil {
		request = NewDescribeFlowCenterAssetListsRequest()
	}
	response = NewDescribeFlowCenterAssetListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwInstancesInfoRequest() (request *DescribeVpcFwInstancesInfoRequest) {
	request = &DescribeVpcFwInstancesInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwInstancesInfo")
	return
}

func NewDescribeVpcFwInstancesInfoResponse() (response *DescribeVpcFwInstancesInfoResponse) {
	response = &DescribeVpcFwInstancesInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户所有VPC防火墙实例卡片信息
func (c *Client) DescribeVpcFwInstancesInfo(request *DescribeVpcFwInstancesInfoRequest) (response *DescribeVpcFwInstancesInfoResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwInstancesInfoRequest()
	}
	response = NewDescribeVpcFwInstancesInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfwBandWidthRequest() (request *DescribeCfwBandWidthRequest) {
	request = &DescribeCfwBandWidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwBandWidth")
	return
}

func NewDescribeCfwBandWidthResponse() (response *DescribeCfwBandWidthResponse) {
	response = &DescribeCfwBandWidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 带宽查询
// 互联网边界带宽 最大带宽
// NAT边界带宽 最大带宽
func (c *Client) DescribeCfwBandWidth(request *DescribeCfwBandWidthRequest) (response *DescribeCfwBandWidthResponse, err error) {
	if request == nil {
		request = NewDescribeCfwBandWidthRequest()
	}
	response = NewDescribeCfwBandWidthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebAssetFilterListRequest() (request *DescribeWebAssetFilterListRequest) {
	request = &DescribeWebAssetFilterListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeWebAssetFilterList")
	return
}

func NewDescribeWebAssetFilterListResponse() (response *DescribeWebAssetFilterListResponse) {
	response = &DescribeWebAssetFilterListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeWebAssetFilterList 资产中心资产扫描Web过滤资产列表
func (c *Client) DescribeWebAssetFilterList(request *DescribeWebAssetFilterListRequest) (response *DescribeWebAssetFilterListResponse, err error) {
	if request == nil {
		request = NewDescribeWebAssetFilterListRequest()
	}
	response = NewDescribeWebAssetFilterListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNewNatCheckInfoRequest() (request *DescribeNewNatCheckInfoRequest) {
	request = &DescribeNewNatCheckInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNewNatCheckInfo")
	return
}

func NewDescribeNewNatCheckInfoResponse() (response *DescribeNewNatCheckInfoResponse) {
	response = &DescribeNewNatCheckInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// NAT开关页-创建实例前校验信息
func (c *Client) DescribeNewNatCheckInfo(request *DescribeNewNatCheckInfoRequest) (response *DescribeNewNatCheckInfoResponse, err error) {
	if request == nil {
		request = NewDescribeNewNatCheckInfoRequest()
	}
	response = NewDescribeNewNatCheckInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAllPublicIPSwitchStatusRequest() (request *ModifyAllPublicIPSwitchStatusRequest) {
	request = &ModifyAllPublicIPSwitchStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllPublicIPSwitchStatus")
	return
}

func NewModifyAllPublicIPSwitchStatusResponse() (response *ModifyAllPublicIPSwitchStatusResponse) {
	response = &ModifyAllPublicIPSwitchStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 互联网边界防火墙一键开关
func (c *Client) ModifyAllPublicIPSwitchStatus(request *ModifyAllPublicIPSwitchStatusRequest) (response *ModifyAllPublicIPSwitchStatusResponse, err error) {
	if request == nil {
		request = NewModifyAllPublicIPSwitchStatusRequest()
	}
	response = NewModifyAllPublicIPSwitchStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcFwTcRuleRequest() (request *ModifyVpcFwTcRuleRequest) {
	request = &ModifyVpcFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcFwTcRule")
	return
}

func NewModifyVpcFwTcRuleResponse() (response *ModifyVpcFwTcRuleResponse) {
	response = &ModifyVpcFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑流量控制策略
func (c *Client) ModifyVpcFwTcRule(request *ModifyVpcFwTcRuleRequest) (response *ModifyVpcFwTcRuleResponse, err error) {
	if request == nil {
		request = NewModifyVpcFwTcRuleRequest()
	}
	response = NewModifyVpcFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindEdgeCFWRequest() (request *UnbindEdgeCFWRequest) {
	request = &UnbindEdgeCFWRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "UnbindEdgeCFW")
	return
}

func NewUnbindEdgeCFWResponse() (response *UnbindEdgeCFWResponse) {
	response = &UnbindEdgeCFWResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SD-WAN edge和CFW解绑
func (c *Client) UnbindEdgeCFW(request *UnbindEdgeCFWRequest) (response *UnbindEdgeCFWResponse, err error) {
	if request == nil {
		request = NewUnbindEdgeCFWRequest()
	}
	response = NewUnbindEdgeCFWResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebAssetScanListRequest() (request *DescribeWebAssetScanListRequest) {
	request = &DescribeWebAssetScanListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeWebAssetScanList")
	return
}

func NewDescribeWebAssetScanListResponse() (response *DescribeWebAssetScanListResponse) {
	response = &DescribeWebAssetScanListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeWebAssetScanList 资产中心资产扫描Web资产原资产列表
func (c *Client) DescribeWebAssetScanList(request *DescribeWebAssetScanListRequest) (response *DescribeWebAssetScanListResponse, err error) {
	if request == nil {
		request = NewDescribeWebAssetScanListRequest()
	}
	response = NewDescribeWebAssetScanListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcFwSequenceRulesRequest() (request *ModifyVpcFwSequenceRulesRequest) {
	request = &ModifyVpcFwSequenceRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcFwSequenceRules")
	return
}

func NewModifyVpcFwSequenceRulesResponse() (response *ModifyVpcFwSequenceRulesResponse) {
	response = &ModifyVpcFwSequenceRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// vpc间规则快速排序
func (c *Client) ModifyVpcFwSequenceRules(request *ModifyVpcFwSequenceRulesRequest) (response *ModifyVpcFwSequenceRulesResponse, err error) {
	if request == nil {
		request = NewModifyVpcFwSequenceRulesRequest()
	}
	response = NewModifyVpcFwSequenceRulesResponse()
	err = c.Send(request, response)
	return
}

func NewSaveAutoBackUpSettingRequest() (request *SaveAutoBackUpSettingRequest) {
	request = &SaveAutoBackUpSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "SaveAutoBackUpSetting")
	return
}

func NewSaveAutoBackUpSettingResponse() (response *SaveAutoBackUpSettingResponse) {
	response = &SaveAutoBackUpSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存自动备份设置
func (c *Client) SaveAutoBackUpSetting(request *SaveAutoBackUpSettingRequest) (response *SaveAutoBackUpSettingResponse, err error) {
	if request == nil {
		request = NewSaveAutoBackUpSettingRequest()
	}
	response = NewSaveAutoBackUpSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatRuleOverviewRequest() (request *DescribeNatRuleOverviewRequest) {
	request = &DescribeNatRuleOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatRuleOverview")
	return
}

func NewDescribeNatRuleOverviewResponse() (response *DescribeNatRuleOverviewResponse) {
	response = &DescribeNatRuleOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// nat规则列表概况
func (c *Client) DescribeNatRuleOverview(request *DescribeNatRuleOverviewRequest) (response *DescribeNatRuleOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeNatRuleOverviewRequest()
	}
	response = NewDescribeNatRuleOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagIpListRequest() (request *DescribeTagIpListRequest) {
	request = &DescribeTagIpListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeTagIpList")
	return
}

func NewDescribeTagIpListResponse() (response *DescribeTagIpListResponse) {
	response = &DescribeTagIpListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询标签对应的ip列表
func (c *Client) DescribeTagIpList(request *DescribeTagIpListRequest) (response *DescribeTagIpListResponse, err error) {
	if request == nil {
		request = NewDescribeTagIpListRequest()
	}
	response = NewDescribeTagIpListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCdcIdsRequest() (request *DescribeCdcIdsRequest) {
	request = &DescribeCdcIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCdcIds")
	return
}

func NewDescribeCdcIdsResponse() (response *DescribeCdcIdsResponse) {
	response = &DescribeCdcIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地专用集群cdc列表
func (c *Client) DescribeCdcIds(request *DescribeCdcIdsRequest) (response *DescribeCdcIdsResponse, err error) {
	if request == nil {
		request = NewDescribeCdcIdsRequest()
	}
	response = NewDescribeCdcIdsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCfwIdpsModeRequest() (request *UpdateCfwIdpsModeRequest) {
	request = &UpdateCfwIdpsModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "UpdateCfwIdpsMode")
	return
}

func NewUpdateCfwIdpsModeResponse() (response *UpdateCfwIdpsModeResponse) {
	response = &UpdateCfwIdpsModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 联动立体防护，更新防火墙防护模式
func (c *Client) UpdateCfwIdpsMode(request *UpdateCfwIdpsModeRequest) (response *UpdateCfwIdpsModeResponse, err error) {
	if request == nil {
		request = NewUpdateCfwIdpsModeRequest()
	}
	response = NewUpdateCfwIdpsModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetEventTreeRequest() (request *DescribeAssetEventTreeRequest) {
	request = &DescribeAssetEventTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssetEventTree")
	return
}

func NewDescribeAssetEventTreeResponse() (response *DescribeAssetEventTreeResponse) {
	response = &DescribeAssetEventTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警中心-资产安全统计-安全告警
func (c *Client) DescribeAssetEventTree(request *DescribeAssetEventTreeRequest) (response *DescribeAssetEventTreeResponse, err error) {
	if request == nil {
		request = NewDescribeAssetEventTreeRequest()
	}
	response = NewDescribeAssetEventTreeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowDistributeListRequest() (request *DescribeFlowDistributeListRequest) {
	request = &DescribeFlowDistributeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFlowDistributeList")
	return
}

func NewDescribeFlowDistributeListResponse() (response *DescribeFlowDistributeListResponse) {
	response = &DescribeFlowDistributeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 互联网流量中心流量地图查询
func (c *Client) DescribeFlowDistributeList(request *DescribeFlowDistributeListRequest) (response *DescribeFlowDistributeListResponse, err error) {
	if request == nil {
		request = NewDescribeFlowDistributeListRequest()
	}
	response = NewDescribeFlowDistributeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatFwInfoCountRequest() (request *DescribeNatFwInfoCountRequest) {
	request = &DescribeNatFwInfoCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInfoCount")
	return
}

func NewDescribeNatFwInfoCountResponse() (response *DescribeNatFwInfoCountResponse) {
	response = &DescribeNatFwInfoCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前用户接入nat防火墙的所有子网数及natfw实例个数
func (c *Client) DescribeNatFwInfoCount(request *DescribeNatFwInfoCountRequest) (response *DescribeNatFwInfoCountResponse, err error) {
	if request == nil {
		request = NewDescribeNatFwInfoCountRequest()
	}
	response = NewDescribeNatFwInfoCountResponse()
	err = c.Send(request, response)
	return
}

func NewBindEdgeCFWRequest() (request *BindEdgeCFWRequest) {
	request = &BindEdgeCFWRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "BindEdgeCFW")
	return
}

func NewBindEdgeCFWResponse() (response *BindEdgeCFWResponse) {
	response = &BindEdgeCFWResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定SD-WAN edge和CFW
func (c *Client) BindEdgeCFW(request *BindEdgeCFWRequest) (response *BindEdgeCFWResponse, err error) {
	if request == nil {
		request = NewBindEdgeCFWRequest()
	}
	response = NewBindEdgeCFWResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatFwReSelectRequest() (request *ModifyNatFwReSelectRequest) {
	request = &ModifyNatFwReSelectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwReSelect")
	return
}

func NewModifyNatFwReSelectResponse() (response *ModifyNatFwReSelectResponse) {
	response = &ModifyNatFwReSelectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙实例重新选择vpc或nat
func (c *Client) ModifyNatFwReSelect(request *ModifyNatFwReSelectRequest) (response *ModifyNatFwReSelectResponse, err error) {
	if request == nil {
		request = NewModifyNatFwReSelectRequest()
	}
	response = NewModifyNatFwReSelectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeChooseAssetRequest() (request *DescribeChooseAssetRequest) {
	request = &DescribeChooseAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeChooseAsset")
	return
}

func NewDescribeChooseAssetResponse() (response *DescribeChooseAssetResponse) {
	response = &DescribeChooseAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeChooseAsset资产中心资产组选择资产记录
func (c *Client) DescribeChooseAsset(request *DescribeChooseAssetRequest) (response *DescribeChooseAssetResponse, err error) {
	if request == nil {
		request = NewDescribeChooseAssetRequest()
	}
	response = NewDescribeChooseAssetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockIgnoreListRequest() (request *DescribeBlockIgnoreListRequest) {
	request = &DescribeBlockIgnoreListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockIgnoreList")
	return
}

func NewDescribeBlockIgnoreListResponse() (response *DescribeBlockIgnoreListResponse) {
	response = &DescribeBlockIgnoreListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询入侵防御放通封禁列表
func (c *Client) DescribeBlockIgnoreList(request *DescribeBlockIgnoreListRequest) (response *DescribeBlockIgnoreListResponse, err error) {
	if request == nil {
		request = NewDescribeBlockIgnoreListRequest()
	}
	response = NewDescribeBlockIgnoreListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResourceGroupOrderRequest() (request *ModifyResourceGroupOrderRequest) {
	request = &ModifyResourceGroupOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyResourceGroupOrder")
	return
}

func NewModifyResourceGroupOrderResponse() (response *ModifyResourceGroupOrderResponse) {
	response = &ModifyResourceGroupOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyResourceGroupOrder-资产中心资产组移动
func (c *Client) ModifyResourceGroupOrder(request *ModifyResourceGroupOrderRequest) (response *ModifyResourceGroupOrderResponse, err error) {
	if request == nil {
		request = NewModifyResourceGroupOrderRequest()
	}
	response = NewModifyResourceGroupOrderResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplyTrialRequest() (request *CreateApplyTrialRequest) {
	request = &CreateApplyTrialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateApplyTrial")
	return
}

func NewCreateApplyTrialResponse() (response *CreateApplyTrialResponse) {
	response = &CreateApplyTrialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 申请试用防火墙，默认试用7天
func (c *Client) CreateApplyTrial(request *CreateApplyTrialRequest) (response *CreateApplyTrialResponse, err error) {
	if request == nil {
		request = NewCreateApplyTrialRequest()
	}
	response = NewCreateApplyTrialResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBetaTaskRequest() (request *ModifyBetaTaskRequest) {
	request = &ModifyBetaTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyBetaTask")
	return
}

func NewModifyBetaTaskResponse() (response *ModifyBetaTaskResponse) {
	response = &ModifyBetaTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改beta任务
func (c *Client) ModifyBetaTask(request *ModifyBetaTaskRequest) (response *ModifyBetaTaskResponse, err error) {
	if request == nil {
		request = NewModifyBetaTaskRequest()
	}
	response = NewModifyBetaTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSwitchListsRequest() (request *DescribeSwitchListsRequest) {
	request = &DescribeSwitchListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeSwitchLists")
	return
}

func NewDescribeSwitchListsResponse() (response *DescribeSwitchListsResponse) {
	response = &DescribeSwitchListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙开关列表
func (c *Client) DescribeSwitchLists(request *DescribeSwitchListsRequest) (response *DescribeSwitchListsResponse, err error) {
	if request == nil {
		request = NewDescribeSwitchListsRequest()
	}
	response = NewDescribeSwitchListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventNewsRequest() (request *DescribeEventNewsRequest) {
	request = &DescribeEventNewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeEventNews")
	return
}

func NewDescribeEventNewsResponse() (response *DescribeEventNewsResponse) {
	response = &DescribeEventNewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新闻事件查询
func (c *Client) DescribeEventNews(request *DescribeEventNewsRequest) (response *DescribeEventNewsResponse, err error) {
	if request == nil {
		request = NewDescribeEventNewsRequest()
	}
	response = NewDescribeEventNewsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateLogSelectRequest() (request *DescribeOperateLogSelectRequest) {
	request = &DescribeOperateLogSelectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeOperateLogSelect")
	return
}

func NewDescribeOperateLogSelectResponse() (response *DescribeOperateLogSelectResponse) {
	response = &DescribeOperateLogSelectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取操作日志筛选框数据
func (c *Client) DescribeOperateLogSelect(request *DescribeOperateLogSelectRequest) (response *DescribeOperateLogSelectResponse, err error) {
	if request == nil {
		request = NewDescribeOperateLogSelectRequest()
	}
	response = NewDescribeOperateLogSelectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPStatusListRequest() (request *DescribeIPStatusListRequest) {
	request = &DescribeIPStatusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeIPStatusList")
	return
}

func NewDescribeIPStatusListResponse() (response *DescribeIPStatusListResponse) {
	response = &DescribeIPStatusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ip防护状态查询
func (c *Client) DescribeIPStatusList(request *DescribeIPStatusListRequest) (response *DescribeIPStatusListResponse, err error) {
	if request == nil {
		request = NewDescribeIPStatusListRequest()
	}
	response = NewDescribeIPStatusListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogsResultAsyncRequest() (request *DescribeLogsResultAsyncRequest) {
	request = &DescribeLogsResultAsyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeLogsResultAsync")
	return
}

func NewDescribeLogsResultAsyncResponse() (response *DescribeLogsResultAsyncResponse) {
	response = &DescribeLogsResultAsyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志审计日志异步查询结果列表获取
func (c *Client) DescribeLogsResultAsync(request *DescribeLogsResultAsyncRequest) (response *DescribeLogsResultAsyncResponse, err error) {
	if request == nil {
		request = NewDescribeLogsResultAsyncRequest()
	}
	response = NewDescribeLogsResultAsyncResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserAuthCheckStatusRequest() (request *ModifyUserAuthCheckStatusRequest) {
	request = &ModifyUserAuthCheckStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyUserAuthCheckStatus")
	return
}

func NewModifyUserAuthCheckStatusResponse() (response *ModifyUserAuthCheckStatusResponse) {
	response = &ModifyUserAuthCheckStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产同步授权状态变更
func (c *Client) ModifyUserAuthCheckStatus(request *ModifyUserAuthCheckStatusRequest) (response *ModifyUserAuthCheckStatusResponse, err error) {
	if request == nil {
		request = NewModifyUserAuthCheckStatusRequest()
	}
	response = NewModifyUserAuthCheckStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAnalyzeLogRequest() (request *DescribeAnalyzeLogRequest) {
	request = &DescribeAnalyzeLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAnalyzeLog")
	return
}

func NewDescribeAnalyzeLogResponse() (response *DescribeAnalyzeLogResponse) {
	response = &DescribeAnalyzeLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分析日志查询新接口
func (c *Client) DescribeAnalyzeLog(request *DescribeAnalyzeLogRequest) (response *DescribeAnalyzeLogResponse, err error) {
	if request == nil {
		request = NewDescribeAnalyzeLogRequest()
	}
	response = NewDescribeAnalyzeLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEdgeStatusRequest() (request *DescribeVpcEdgeStatusRequest) {
	request = &DescribeVpcEdgeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcEdgeStatus")
	return
}

func NewDescribeVpcEdgeStatusResponse() (response *DescribeVpcEdgeStatusResponse) {
	response = &DescribeVpcEdgeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看vpc防火墙边开关状态
func (c *Client) DescribeVpcEdgeStatus(request *DescribeVpcEdgeStatusRequest) (response *DescribeVpcEdgeStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVpcEdgeStatusRequest()
	}
	response = NewDescribeVpcEdgeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCheckVpcFwTcRuleRequest() (request *CheckVpcFwTcRuleRequest) {
	request = &CheckVpcFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CheckVpcFwTcRule")
	return
}

func NewCheckVpcFwTcRuleResponse() (response *CheckVpcFwTcRuleResponse) {
	response = &CheckVpcFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查流量控制规则是否有冲突
func (c *Client) CheckVpcFwTcRule(request *CheckVpcFwTcRuleRequest) (response *CheckVpcFwTcRuleResponse, err error) {
	if request == nil {
		request = NewCheckVpcFwTcRuleRequest()
	}
	response = NewCheckVpcFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGuidePortProtectRequest() (request *DescribeGuidePortProtectRequest) {
	request = &DescribeGuidePortProtectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeGuidePortProtect")
	return
}

func NewDescribeGuidePortProtectResponse() (response *DescribeGuidePortProtectResponse) {
	response = &DescribeGuidePortProtectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeGuidePortProtect 新手引导禁封端口信息查询
func (c *Client) DescribeGuidePortProtect(request *DescribeGuidePortProtectRequest) (response *DescribeGuidePortProtectResponse, err error) {
	if request == nil {
		request = NewDescribeGuidePortProtectRequest()
	}
	response = NewDescribeGuidePortProtectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleOverviewRequest() (request *DescribeRuleOverviewRequest) {
	request = &DescribeRuleOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeRuleOverview")
	return
}

func NewDescribeRuleOverviewResponse() (response *DescribeRuleOverviewResponse) {
	response = &DescribeRuleOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询规则列表概况
func (c *Client) DescribeRuleOverview(request *DescribeRuleOverviewRequest) (response *DescribeRuleOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeRuleOverviewRequest()
	}
	response = NewDescribeRuleOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatSequenceRulesRequest() (request *ModifyNatSequenceRulesRequest) {
	request = &ModifyNatSequenceRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatSequenceRules")
	return
}

func NewModifyNatSequenceRulesResponse() (response *ModifyNatSequenceRulesResponse) {
	response = &ModifyNatSequenceRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// NAT防火墙规则快速排序
func (c *Client) ModifyNatSequenceRules(request *ModifyNatSequenceRulesRequest) (response *ModifyNatSequenceRulesResponse, err error) {
	if request == nil {
		request = NewModifyNatSequenceRulesRequest()
	}
	response = NewModifyNatSequenceRulesResponse()
	err = c.Send(request, response)
	return
}

func NewAddNatAcRuleRequest() (request *AddNatAcRuleRequest) {
	request = &AddNatAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddNatAcRule")
	return
}

func NewAddNatAcRuleResponse() (response *AddNatAcRuleResponse) {
	response = &AddNatAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加nat访问控制规则
func (c *Client) AddNatAcRule(request *AddNatAcRuleRequest) (response *AddNatAcRuleResponse, err error) {
	if request == nil {
		request = NewAddNatAcRuleRequest()
	}
	response = NewAddNatAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProtectModeCountRequest() (request *DescribeProtectModeCountRequest) {
	request = &DescribeProtectModeCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeProtectModeCount")
	return
}

func NewDescribeProtectModeCountResponse() (response *DescribeProtectModeCountResponse) {
	response = &DescribeProtectModeCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeProtectModeCount
func (c *Client) DescribeProtectModeCount(request *DescribeProtectModeCountRequest) (response *DescribeProtectModeCountResponse, err error) {
	if request == nil {
		request = NewDescribeProtectModeCountRequest()
	}
	response = NewDescribeProtectModeCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIsolateTableRequest() (request *ModifyIsolateTableRequest) {
	request = &ModifyIsolateTableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyIsolateTable")
	return
}

func NewModifyIsolateTableResponse() (response *ModifyIsolateTableResponse) {
	response = &ModifyIsolateTableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyIsolateTable 隔离列表编辑和删除操作
func (c *Client) ModifyIsolateTable(request *ModifyIsolateTableRequest) (response *ModifyIsolateTableResponse, err error) {
	if request == nil {
		request = NewModifyIsolateTableRequest()
	}
	response = NewModifyIsolateTableResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcFwGroupRequest() (request *ModifyVpcFwGroupRequest) {
	request = &ModifyVpcFwGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcFwGroup")
	return
}

func NewModifyVpcFwGroupResponse() (response *ModifyVpcFwGroupResponse) {
	response = &ModifyVpcFwGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑VPC间防火墙(防火墙组)
func (c *Client) ModifyVpcFwGroup(request *ModifyVpcFwGroupRequest) (response *ModifyVpcFwGroupResponse, err error) {
	if request == nil {
		request = NewModifyVpcFwGroupRequest()
	}
	response = NewModifyVpcFwGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpsRuleListRequest() (request *DescribeIpsRuleListRequest) {
	request = &DescribeIpsRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeIpsRuleList")
	return
}

func NewDescribeIpsRuleListResponse() (response *DescribeIpsRuleListResponse) {
	response = &DescribeIpsRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// IPS规则查询接口
func (c *Client) DescribeIpsRuleList(request *DescribeIpsRuleListRequest) (response *DescribeIpsRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeIpsRuleListRequest()
	}
	response = NewDescribeIpsRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIpsRuleListRequest() (request *ModifyIpsRuleListRequest) {
	request = &ModifyIpsRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyIpsRuleList")
	return
}

func NewModifyIpsRuleListResponse() (response *ModifyIpsRuleListResponse) {
	response = &ModifyIpsRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改IPS规则列表
func (c *Client) ModifyIpsRuleList(request *ModifyIpsRuleListRequest) (response *ModifyIpsRuleListResponse, err error) {
	if request == nil {
		request = NewModifyIpsRuleListRequest()
	}
	response = NewModifyIpsRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfwUserStatusRequest() (request *DescribeCfwUserStatusRequest) {
	request = &DescribeCfwUserStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwUserStatus")
	return
}

func NewDescribeCfwUserStatusResponse() (response *DescribeCfwUserStatusResponse) {
	response = &DescribeCfwUserStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户云防火墙使用情况，打开、关闭、使用中产生的安全事件
func (c *Client) DescribeCfwUserStatus(request *DescribeCfwUserStatusRequest) (response *DescribeCfwUserStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCfwUserStatusRequest()
	}
	response = NewDescribeCfwUserStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcFwInstanceRequest() (request *CreateVpcFwInstanceRequest) {
	request = &CreateVpcFwInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateVpcFwInstance")
	return
}

func NewCreateVpcFwInstanceResponse() (response *CreateVpcFwInstanceResponse) {
	response = &CreateVpcFwInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateVpcFwInstance 建立vpc实例
func (c *Client) CreateVpcFwInstance(request *CreateVpcFwInstanceRequest) (response *CreateVpcFwInstanceResponse, err error) {
	if request == nil {
		request = NewCreateVpcFwInstanceRequest()
	}
	response = NewCreateVpcFwInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBetaTaskRequest() (request *DeleteBetaTaskRequest) {
	request = &DeleteBetaTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteBetaTask")
	return
}

func NewDeleteBetaTaskResponse() (response *DeleteBetaTaskResponse) {
	response = &DeleteBetaTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Beta任务
func (c *Client) DeleteBetaTask(request *DeleteBetaTaskRequest) (response *DeleteBetaTaskResponse, err error) {
	if request == nil {
		request = NewDeleteBetaTaskRequest()
	}
	response = NewDeleteBetaTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStorageLogTypeSettingRequest() (request *DescribeStorageLogTypeSettingRequest) {
	request = &DescribeStorageLogTypeSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeStorageLogTypeSetting")
	return
}

func NewDescribeStorageLogTypeSettingResponse() (response *DescribeStorageLogTypeSettingResponse) {
	response = &DescribeStorageLogTypeSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户日志存储类型配置查询
func (c *Client) DescribeStorageLogTypeSetting(request *DescribeStorageLogTypeSettingRequest) (response *DescribeStorageLogTypeSettingResponse, err error) {
	if request == nil {
		request = NewDescribeStorageLogTypeSettingRequest()
	}
	response = NewDescribeStorageLogTypeSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAllAccessControlRuleRequest() (request *DeleteAllAccessControlRuleRequest) {
	request = &DeleteAllAccessControlRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteAllAccessControlRule")
	return
}

func NewDeleteAllAccessControlRuleResponse() (response *DeleteAllAccessControlRuleResponse) {
	response = &DeleteAllAccessControlRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteAllAccessControlRule-删除全部访问控制规则-vpc，nat
func (c *Client) DeleteAllAccessControlRule(request *DeleteAllAccessControlRuleRequest) (response *DeleteAllAccessControlRuleResponse, err error) {
	if request == nil {
		request = NewDeleteAllAccessControlRuleRequest()
	}
	response = NewDeleteAllAccessControlRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcTopolListsRequest() (request *DescribeVpcTopolListsRequest) {
	request = &DescribeVpcTopolListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcTopolLists")
	return
}

func NewDescribeVpcTopolListsResponse() (response *DescribeVpcTopolListsResponse) {
	response = &DescribeVpcTopolListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙开关-VPC边界防火墙-私有网络VPC拓扑图
func (c *Client) DescribeVpcTopolLists(request *DescribeVpcTopolListsRequest) (response *DescribeVpcTopolListsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcTopolListsRequest()
	}
	response = NewDescribeVpcTopolListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeEdgeRequest() (request *DescribeNodeEdgeRequest) {
	request = &DescribeNodeEdgeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNodeEdge")
	return
}

func NewDescribeNodeEdgeResponse() (response *DescribeNodeEdgeResponse) {
	response = &DescribeNodeEdgeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙开关-VPC边界防火墙-私有网络VPC拓扑图
func (c *Client) DescribeNodeEdge(request *DescribeNodeEdgeRequest) (response *DescribeNodeEdgeResponse, err error) {
	if request == nil {
		request = NewDescribeNodeEdgeRequest()
	}
	response = NewDescribeNodeEdgeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPolicyAuthorityRequest() (request *ModifyPolicyAuthorityRequest) {
	request = &ModifyPolicyAuthorityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyPolicyAuthority")
	return
}

func NewModifyPolicyAuthorityResponse() (response *ModifyPolicyAuthorityResponse) {
	response = &ModifyPolicyAuthorityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上报策略权限
func (c *Client) ModifyPolicyAuthority(request *ModifyPolicyAuthorityRequest) (response *ModifyPolicyAuthorityResponse, err error) {
	if request == nil {
		request = NewModifyPolicyAuthorityRequest()
	}
	response = NewModifyPolicyAuthorityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAcListsRequest() (request *DescribeAcListsRequest) {
	request = &DescribeAcListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAcLists")
	return
}

func NewDescribeAcListsResponse() (response *DescribeAcListsResponse) {
	response = &DescribeAcListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 访问控制列表
func (c *Client) DescribeAcLists(request *DescribeAcListsRequest) (response *DescribeAcListsResponse, err error) {
	if request == nil {
		request = NewDescribeAcListsRequest()
	}
	response = NewDescribeAcListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfwUpdateStatusRequest() (request *DescribeCfwUpdateStatusRequest) {
	request = &DescribeCfwUpdateStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwUpdateStatus")
	return
}

func NewDescribeCfwUpdateStatusResponse() (response *DescribeCfwUpdateStatusResponse) {
	response = &DescribeCfwUpdateStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// cfw引擎升级状态查询
func (c *Client) DescribeCfwUpdateStatus(request *DescribeCfwUpdateStatusRequest) (response *DescribeCfwUpdateStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCfwUpdateStatusRequest()
	}
	response = NewDescribeCfwUpdateStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceGroupChangeImpactRequest() (request *DescribeResourceGroupChangeImpactRequest) {
	request = &DescribeResourceGroupChangeImpactRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeResourceGroupChangeImpact")
	return
}

func NewDescribeResourceGroupChangeImpactResponse() (response *DescribeResourceGroupChangeImpactResponse) {
	response = &DescribeResourceGroupChangeImpactResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源分组变更影响-DescribeResourceGroupChangeImpact
func (c *Client) DescribeResourceGroupChangeImpact(request *DescribeResourceGroupChangeImpactRequest) (response *DescribeResourceGroupChangeImpactResponse, err error) {
	if request == nil {
		request = NewDescribeResourceGroupChangeImpactRequest()
	}
	response = NewDescribeResourceGroupChangeImpactResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMessageCenterSwitchRequest() (request *ModifyMessageCenterSwitchRequest) {
	request = &ModifyMessageCenterSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyMessageCenterSwitch")
	return
}

func NewModifyMessageCenterSwitchResponse() (response *ModifyMessageCenterSwitchResponse) {
	response = &ModifyMessageCenterSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改使用消息中心订阅推送
func (c *Client) ModifyMessageCenterSwitch(request *ModifyMessageCenterSwitchRequest) (response *ModifyMessageCenterSwitchResponse, err error) {
	if request == nil {
		request = NewModifyMessageCenterSwitchRequest()
	}
	response = NewModifyMessageCenterSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNatFwInstanceRequest() (request *DeleteNatFwInstanceRequest) {
	request = &DeleteNatFwInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteNatFwInstance")
	return
}

func NewDeleteNatFwInstanceResponse() (response *DeleteNatFwInstanceResponse) {
	response = &DeleteNatFwInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁防火墙实例
func (c *Client) DeleteNatFwInstance(request *DeleteNatFwInstanceRequest) (response *DeleteNatFwInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteNatFwInstanceRequest()
	}
	response = NewDeleteNatFwInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyStorageLogTypeSettingRequest() (request *ModifyStorageLogTypeSettingRequest) {
	request = &ModifyStorageLogTypeSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyStorageLogTypeSetting")
	return
}

func NewModifyStorageLogTypeSettingResponse() (response *ModifyStorageLogTypeSettingResponse) {
	response = &ModifyStorageLogTypeSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户日志存储类型配置修改
func (c *Client) ModifyStorageLogTypeSetting(request *ModifyStorageLogTypeSettingRequest) (response *ModifyStorageLogTypeSettingResponse, err error) {
	if request == nil {
		request = NewModifyStorageLogTypeSettingRequest()
	}
	response = NewModifyStorageLogTypeSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOfflineExportTemporaryCredentialsRequest() (request *DescribeOfflineExportTemporaryCredentialsRequest) {
	request = &DescribeOfflineExportTemporaryCredentialsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeOfflineExportTemporaryCredentials")
	return
}

func NewDescribeOfflineExportTemporaryCredentialsResponse() (response *DescribeOfflineExportTemporaryCredentialsResponse) {
	response = &DescribeOfflineExportTemporaryCredentialsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志离线导出任务文件下载临时凭证
func (c *Client) DescribeOfflineExportTemporaryCredentials(request *DescribeOfflineExportTemporaryCredentialsRequest) (response *DescribeOfflineExportTemporaryCredentialsResponse, err error) {
	if request == nil {
		request = NewDescribeOfflineExportTemporaryCredentialsRequest()
	}
	response = NewDescribeOfflineExportTemporaryCredentialsResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecurityGroupTableStatusRequest() (request *ModifySecurityGroupTableStatusRequest) {
	request = &ModifySecurityGroupTableStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifySecurityGroupTableStatus")
	return
}

func NewModifySecurityGroupTableStatusResponse() (response *ModifySecurityGroupTableStatusResponse) {
	response = &ModifySecurityGroupTableStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改安全组列表状态
func (c *Client) ModifySecurityGroupTableStatus(request *ModifySecurityGroupTableStatusRequest) (response *ModifySecurityGroupTableStatusResponse, err error) {
	if request == nil {
		request = NewModifySecurityGroupTableStatusRequest()
	}
	response = NewModifySecurityGroupTableStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetStatisticRequest() (request *DescribeAssetStatisticRequest) {
	request = &DescribeAssetStatisticRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssetStatistic")
	return
}

func NewDescribeAssetStatisticResponse() (response *DescribeAssetStatisticResponse) {
	response = &DescribeAssetStatisticResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产中心详情统计图
func (c *Client) DescribeAssetStatistic(request *DescribeAssetStatisticRequest) (response *DescribeAssetStatisticResponse, err error) {
	if request == nil {
		request = NewDescribeAssetStatisticRequest()
	}
	response = NewDescribeAssetStatisticResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePcapDownUrlRequest() (request *DescribePcapDownUrlRequest) {
	request = &DescribePcapDownUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribePcapDownUrl")
	return
}

func NewDescribePcapDownUrlResponse() (response *DescribePcapDownUrlResponse) {
	response = &DescribePcapDownUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询抓包任务下载链接
func (c *Client) DescribePcapDownUrl(request *DescribePcapDownUrlRequest) (response *DescribePcapDownUrlResponse, err error) {
	if request == nil {
		request = NewDescribePcapDownUrlRequest()
	}
	response = NewDescribePcapDownUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGuideScanInfoRequest() (request *DescribeGuideScanInfoRequest) {
	request = &DescribeGuideScanInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeGuideScanInfo")
	return
}

func NewDescribeGuideScanInfoResponse() (response *DescribeGuideScanInfoResponse) {
	response = &DescribeGuideScanInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeGuideScanInfo新手引导扫描接口信息
func (c *Client) DescribeGuideScanInfo(request *DescribeGuideScanInfoRequest) (response *DescribeGuideScanInfoResponse, err error) {
	if request == nil {
		request = NewDescribeGuideScanInfoRequest()
	}
	response = NewDescribeGuideScanInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcAclEdgeRangeRequest() (request *DescribeVpcAclEdgeRangeRequest) {
	request = &DescribeVpcAclEdgeRangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcAclEdgeRange")
	return
}

func NewDescribeVpcAclEdgeRangeResponse() (response *DescribeVpcAclEdgeRangeResponse) {
	response = &DescribeVpcAclEdgeRangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询内网间访问控制规则的生效范围
func (c *Client) DescribeVpcAclEdgeRange(request *DescribeVpcAclEdgeRangeRequest) (response *DescribeVpcAclEdgeRangeResponse, err error) {
	if request == nil {
		request = NewDescribeVpcAclEdgeRangeRequest()
	}
	response = NewDescribeVpcAclEdgeRangeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSwitchStatusRequest() (request *DescribeSwitchStatusRequest) {
	request = &DescribeSwitchStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeSwitchStatus")
	return
}

func NewDescribeSwitchStatusResponse() (response *DescribeSwitchStatusResponse) {
	response = &DescribeSwitchStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防火墙开关状态
func (c *Client) DescribeSwitchStatus(request *DescribeSwitchStatusRequest) (response *DescribeSwitchStatusResponse, err error) {
	if request == nil {
		request = NewDescribeSwitchStatusRequest()
	}
	response = NewDescribeSwitchStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAcRuleRequest() (request *ModifyAcRuleRequest) {
	request = &ModifyAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyAcRule")
	return
}

func NewModifyAcRuleResponse() (response *ModifyAcRuleResponse) {
	response = &ModifyAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改规则
func (c *Client) ModifyAcRule(request *ModifyAcRuleRequest) (response *ModifyAcRuleResponse, err error) {
	if request == nil {
		request = NewModifyAcRuleRequest()
	}
	response = NewModifyAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeErrorDetailRequest() (request *DescribeErrorDetailRequest) {
	request = &DescribeErrorDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeErrorDetail")
	return
}

func NewDescribeErrorDetailResponse() (response *DescribeErrorDetailResponse) {
	response = &DescribeErrorDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 异步错误信息查询接口
func (c *Client) DescribeErrorDetail(request *DescribeErrorDetailRequest) (response *DescribeErrorDetailResponse, err error) {
	if request == nil {
		request = NewDescribeErrorDetailRequest()
	}
	response = NewDescribeErrorDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllZoneListRequest() (request *DescribeAllZoneListRequest) {
	request = &DescribeAllZoneListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAllZoneList")
	return
}

func NewDescribeAllZoneListResponse() (response *DescribeAllZoneListResponse) {
	response = &DescribeAllZoneListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAllZoneList
func (c *Client) DescribeAllZoneList(request *DescribeAllZoneListRequest) (response *DescribeAllZoneListResponse, err error) {
	if request == nil {
		request = NewDescribeAllZoneListRequest()
	}
	response = NewDescribeAllZoneListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAreaStatusRequest() (request *DescribeAreaStatusRequest) {
	request = &DescribeAreaStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAreaStatus")
	return
}

func NewDescribeAreaStatusResponse() (response *DescribeAreaStatusResponse) {
	response = &DescribeAreaStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// NAT开关页-选中地域实例创建状态
func (c *Client) DescribeAreaStatus(request *DescribeAreaStatusRequest) (response *DescribeAreaStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAreaStatusRequest()
	}
	response = NewDescribeAreaStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEWRuleStatusRequest() (request *ModifyEWRuleStatusRequest) {
	request = &ModifyEWRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyEWRuleStatus")
	return
}

func NewModifyEWRuleStatusResponse() (response *ModifyEWRuleStatusResponse) {
	response = &ModifyEWRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用停用VPC间规则或Nat边界规则
// VPC间规则需指定EdgeId。Nat边界规则需指定地域Region与Direction。
func (c *Client) ModifyEWRuleStatus(request *ModifyEWRuleStatusRequest) (response *ModifyEWRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyEWRuleStatusRequest()
	}
	response = NewModifyEWRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwCrossStatusRequest() (request *DescribeVpcFwCrossStatusRequest) {
	request = &DescribeVpcFwCrossStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwCrossStatus")
	return
}

func NewDescribeVpcFwCrossStatusResponse() (response *DescribeVpcFwCrossStatusResponse) {
	response = &DescribeVpcFwCrossStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VPC防火墙跨租户edge或者vpc的状态
func (c *Client) DescribeVpcFwCrossStatus(request *DescribeVpcFwCrossStatusRequest) (response *DescribeVpcFwCrossStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwCrossStatusRequest()
	}
	response = NewDescribeVpcFwCrossStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwStatusBarRequest() (request *DescribeVpcFwStatusBarRequest) {
	request = &DescribeVpcFwStatusBarRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwStatusBar")
	return
}

func NewDescribeVpcFwStatusBarResponse() (response *DescribeVpcFwStatusBarResponse) {
	response = &DescribeVpcFwStatusBarResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VPC防火墙状态栏数据
func (c *Client) DescribeVpcFwStatusBar(request *DescribeVpcFwStatusBarRequest) (response *DescribeVpcFwStatusBarResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwStatusBarRequest()
	}
	response = NewDescribeVpcFwStatusBarResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatExistRegionsRequest() (request *DescribeNatExistRegionsRequest) {
	request = &DescribeNatExistRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatExistRegions")
	return
}

func NewDescribeNatExistRegionsResponse() (response *DescribeNatExistRegionsResponse) {
	response = &DescribeNatExistRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// NAT已创建实例地域查询
func (c *Client) DescribeNatExistRegions(request *DescribeNatExistRegionsRequest) (response *DescribeNatExistRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeNatExistRegionsRequest()
	}
	response = NewDescribeNatExistRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceGroupRequest() (request *DescribeResourceGroupRequest) {
	request = &DescribeResourceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeResourceGroup")
	return
}

func NewDescribeResourceGroupResponse() (response *DescribeResourceGroupResponse) {
	response = &DescribeResourceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeResourceGroup资产中心资产树信息
func (c *Client) DescribeResourceGroup(request *DescribeResourceGroupRequest) (response *DescribeResourceGroupResponse, err error) {
	if request == nil {
		request = NewDescribeResourceGroupRequest()
	}
	response = NewDescribeResourceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVPCSwitchStatusRequest() (request *ModifyVPCSwitchStatusRequest) {
	request = &ModifyVPCSwitchStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVPCSwitchStatus")
	return
}

func NewModifyVPCSwitchStatusResponse() (response *ModifyVPCSwitchStatusResponse) {
	response = &ModifyVPCSwitchStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 单个修改VPC火墙开关
func (c *Client) ModifyVPCSwitchStatus(request *ModifyVPCSwitchStatusRequest) (response *ModifyVPCSwitchStatusResponse, err error) {
	if request == nil {
		request = NewModifyVPCSwitchStatusRequest()
	}
	response = NewModifyVPCSwitchStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveAcRuleRequest() (request *RemoveAcRuleRequest) {
	request = &RemoveAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "RemoveAcRule")
	return
}

func NewRemoveAcRuleResponse() (response *RemoveAcRuleResponse) {
	response = &RemoveAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除互联网边界规则
func (c *Client) RemoveAcRule(request *RemoveAcRuleRequest) (response *RemoveAcRuleResponse, err error) {
	if request == nil {
		request = NewRemoveAcRuleRequest()
	}
	response = NewRemoveAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDynamicCheckRequest() (request *ModifyDynamicCheckRequest) {
	request = &ModifyDynamicCheckRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyDynamicCheck")
	return
}

func NewModifyDynamicCheckResponse() (response *ModifyDynamicCheckResponse) {
	response = &ModifyDynamicCheckResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyDynamicCheck产品动态取消角标数量接口
func (c *Client) ModifyDynamicCheck(request *ModifyDynamicCheckRequest) (response *ModifyDynamicCheckResponse, err error) {
	if request == nil {
		request = NewModifyDynamicCheckRequest()
	}
	response = NewModifyDynamicCheckResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowCenterLogsV1Request() (request *DescribeFlowCenterLogsV1Request) {
	request = &DescribeFlowCenterLogsV1Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFlowCenterLogsV1")
	return
}

func NewDescribeFlowCenterLogsV1Response() (response *DescribeFlowCenterLogsV1Response) {
	response = &DescribeFlowCenterLogsV1Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 互联网流量中心列表查询
func (c *Client) DescribeFlowCenterLogsV1(request *DescribeFlowCenterLogsV1Request) (response *DescribeFlowCenterLogsV1Response, err error) {
	if request == nil {
		request = NewDescribeFlowCenterLogsV1Request()
	}
	response = NewDescribeFlowCenterLogsV1Response()
	err = c.Send(request, response)
	return
}

func NewClearResourceGroupRequest() (request *ClearResourceGroupRequest) {
	request = &ClearResourceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ClearResourceGroup")
	return
}

func NewClearResourceGroupResponse() (response *ClearResourceGroupResponse) {
	response = &ClearResourceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清空资产分组
func (c *Client) ClearResourceGroup(request *ClearResourceGroupRequest) (response *ClearResourceGroupResponse, err error) {
	if request == nil {
		request = NewClearResourceGroupRequest()
	}
	response = NewClearResourceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveVpcAcRuleRequest() (request *RemoveVpcAcRuleRequest) {
	request = &RemoveVpcAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "RemoveVpcAcRule")
	return
}

func NewRemoveVpcAcRuleResponse() (response *RemoveVpcAcRuleResponse) {
	response = &RemoveVpcAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除VPC间规则
func (c *Client) RemoveVpcAcRule(request *RemoveVpcAcRuleRequest) (response *RemoveVpcAcRuleResponse, err error) {
	if request == nil {
		request = NewRemoveVpcAcRuleRequest()
	}
	response = NewRemoveVpcAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatFwInstanceRequest() (request *DescribeNatFwInstanceRequest) {
	request = &DescribeNatFwInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInstance")
	return
}

func NewDescribeNatFwInstanceResponse() (response *DescribeNatFwInstanceResponse) {
	response = &DescribeNatFwInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeNatFwInstance 获取租户所有NAT实例
func (c *Client) DescribeNatFwInstance(request *DescribeNatFwInstanceRequest) (response *DescribeNatFwInstanceResponse, err error) {
	if request == nil {
		request = NewDescribeNatFwInstanceRequest()
	}
	response = NewDescribeNatFwInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDefenseSwitchRequest() (request *ModifyDefenseSwitchRequest) {
	request = &ModifyDefenseSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyDefenseSwitch")
	return
}

func NewModifyDefenseSwitchResponse() (response *ModifyDefenseSwitchResponse) {
	response = &ModifyDefenseSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改入侵防御开关
func (c *Client) ModifyDefenseSwitch(request *ModifyDefenseSwitchRequest) (response *ModifyDefenseSwitchResponse, err error) {
	if request == nil {
		request = NewModifyDefenseSwitchRequest()
	}
	response = NewModifyDefenseSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcFwSwitchRequest() (request *ModifyVpcFwSwitchRequest) {
	request = &ModifyVpcFwSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcFwSwitch")
	return
}

func NewModifyVpcFwSwitchResponse() (response *ModifyVpcFwSwitchResponse) {
	response = &ModifyVpcFwSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改VPC间防火墙开关
func (c *Client) ModifyVpcFwSwitch(request *ModifyVpcFwSwitchRequest) (response *ModifyVpcFwSwitchResponse, err error) {
	if request == nil {
		request = NewModifyVpcFwSwitchRequest()
	}
	response = NewModifyVpcFwSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfwCidrRequest() (request *DescribeCfwCidrRequest) {
	request = &DescribeCfwCidrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwCidr")
	return
}

func NewDescribeCfwCidrResponse() (response *DescribeCfwCidrResponse) {
	response = &DescribeCfwCidrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防火墙网段信息
func (c *Client) DescribeCfwCidr(request *DescribeCfwCidrRequest) (response *DescribeCfwCidrResponse, err error) {
	if request == nil {
		request = NewDescribeCfwCidrRequest()
	}
	response = NewDescribeCfwCidrResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetScanListRequest() (request *DescribeAssetScanListRequest) {
	request = &DescribeAssetScanListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssetScanList")
	return
}

func NewDescribeAssetScanListResponse() (response *DescribeAssetScanListResponse) {
	response = &DescribeAssetScanListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAssetScanList资产中心资产扫描过滤资产原资产列表
func (c *Client) DescribeAssetScanList(request *DescribeAssetScanListRequest) (response *DescribeAssetScanListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetScanListRequest()
	}
	response = NewDescribeAssetScanListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetScanStatusRequest() (request *DescribeAssetScanStatusRequest) {
	request = &DescribeAssetScanStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssetScanStatus")
	return
}

func NewDescribeAssetScanStatusResponse() (response *DescribeAssetScanStatusResponse) {
	response = &DescribeAssetScanStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产扫描状态
func (c *Client) DescribeAssetScanStatus(request *DescribeAssetScanStatusRequest) (response *DescribeAssetScanStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAssetScanStatusRequest()
	}
	response = NewDescribeAssetScanStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFuncDynamicsRequest() (request *DescribeFuncDynamicsRequest) {
	request = &DescribeFuncDynamicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFuncDynamics")
	return
}

func NewDescribeFuncDynamicsResponse() (response *DescribeFuncDynamicsResponse) {
	response = &DescribeFuncDynamicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 功能动态查询
func (c *Client) DescribeFuncDynamics(request *DescribeFuncDynamicsRequest) (response *DescribeFuncDynamicsResponse, err error) {
	if request == nil {
		request = NewDescribeFuncDynamicsRequest()
	}
	response = NewDescribeFuncDynamicsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetFilterListRequest() (request *DescribeAssetFilterListRequest) {
	request = &DescribeAssetFilterListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssetFilterList")
	return
}

func NewDescribeAssetFilterListResponse() (response *DescribeAssetFilterListResponse) {
	response = &DescribeAssetFilterListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAssetFilterList资产中心资产扫描过滤资产过滤资产列表
func (c *Client) DescribeAssetFilterList(request *DescribeAssetFilterListRequest) (response *DescribeAssetFilterListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetFilterListRequest()
	}
	response = NewDescribeAssetFilterListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatGwListRequest() (request *DescribeNatGwListRequest) {
	request = &DescribeNatGwListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatGwList")
	return
}

func NewDescribeNatGwListResponse() (response *DescribeNatGwListResponse) {
	response = &DescribeNatGwListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户nat网关列表
func (c *Client) DescribeNatGwList(request *DescribeNatGwListRequest) (response *DescribeNatGwListResponse, err error) {
	if request == nil {
		request = NewDescribeNatGwListRequest()
	}
	response = NewDescribeNatGwListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwGroupSwitchRequest() (request *DescribeVpcFwGroupSwitchRequest) {
	request = &DescribeVpcFwGroupSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwGroupSwitch")
	return
}

func NewDescribeVpcFwGroupSwitchResponse() (response *DescribeVpcFwGroupSwitchResponse) {
	response = &DescribeVpcFwGroupSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// VPC防火墙(组)开关列表
func (c *Client) DescribeVpcFwGroupSwitch(request *DescribeVpcFwGroupSwitchRequest) (response *DescribeVpcFwGroupSwitchResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwGroupSwitchRequest()
	}
	response = NewDescribeVpcFwGroupSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwGroupInsRequest() (request *DescribeVpcFwGroupInsRequest) {
	request = &DescribeVpcFwGroupInsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwGroupIns")
	return
}

func NewDescribeVpcFwGroupInsResponse() (response *DescribeVpcFwGroupInsResponse) {
	response = &DescribeVpcFwGroupInsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防火墙(组)ID名称及实例下对应关系
func (c *Client) DescribeVpcFwGroupIns(request *DescribeVpcFwGroupInsRequest) (response *DescribeVpcFwGroupInsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwGroupInsRequest()
	}
	response = NewDescribeVpcFwGroupInsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBlockTopRequest() (request *ModifyBlockTopRequest) {
	request = &ModifyBlockTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockTop")
	return
}

func NewModifyBlockTopResponse() (response *ModifyBlockTopResponse) {
	response = &ModifyBlockTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyBlockTop取消置顶接口
func (c *Client) ModifyBlockTop(request *ModifyBlockTopRequest) (response *ModifyBlockTopResponse, err error) {
	if request == nil {
		request = NewModifyBlockTopRequest()
	}
	response = NewModifyBlockTopResponse()
	err = c.Send(request, response)
	return
}

func NewSetNatFwDnatRuleRequest() (request *SetNatFwDnatRuleRequest) {
	request = &SetNatFwDnatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "SetNatFwDnatRule")
	return
}

func NewSetNatFwDnatRuleResponse() (response *SetNatFwDnatRuleResponse) {
	response = &SetNatFwDnatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置防火墙Dnat规则
func (c *Client) SetNatFwDnatRule(request *SetNatFwDnatRuleRequest) (response *SetNatFwDnatRuleResponse, err error) {
	if request == nil {
		request = NewSetNatFwDnatRuleRequest()
	}
	response = NewSetNatFwDnatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityGroupListRequest() (request *DescribeSecurityGroupListRequest) {
	request = &DescribeSecurityGroupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeSecurityGroupList")
	return
}

func NewDescribeSecurityGroupListResponse() (response *DescribeSecurityGroupListResponse) {
	response = &DescribeSecurityGroupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全组规则列表
func (c *Client) DescribeSecurityGroupList(request *DescribeSecurityGroupListRequest) (response *DescribeSecurityGroupListResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupListRequest()
	}
	response = NewDescribeSecurityGroupListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserInfoOverviewRequest() (request *DescribeUserInfoOverviewRequest) {
	request = &DescribeUserInfoOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeUserInfoOverview")
	return
}

func NewDescribeUserInfoOverviewResponse() (response *DescribeUserInfoOverviewResponse) {
	response = &DescribeUserInfoOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户信息概览
func (c *Client) DescribeUserInfoOverview(request *DescribeUserInfoOverviewRequest) (response *DescribeUserInfoOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeUserInfoOverviewRequest()
	}
	response = NewDescribeUserInfoOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcRuleStatusRequest() (request *DescribeVpcRuleStatusRequest) {
	request = &DescribeVpcRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcRuleStatus")
	return
}

func NewDescribeVpcRuleStatusResponse() (response *DescribeVpcRuleStatusResponse) {
	response = &DescribeVpcRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询内网间规则的配额和使用情况
func (c *Client) DescribeVpcRuleStatus(request *DescribeVpcRuleStatusRequest) (response *DescribeVpcRuleStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVpcRuleStatusRequest()
	}
	response = NewDescribeVpcRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAcRulesRequest() (request *CreateAcRulesRequest) {
	request = &CreateAcRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateAcRules")
	return
}

func NewCreateAcRulesResponse() (response *CreateAcRulesResponse) {
	response = &CreateAcRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建规则
func (c *Client) CreateAcRules(request *CreateAcRulesRequest) (response *CreateAcRulesResponse, err error) {
	if request == nil {
		request = NewCreateAcRulesRequest()
	}
	response = NewCreateAcRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetOverviewRequest() (request *DescribeAssetOverviewRequest) {
	request = &DescribeAssetOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssetOverview")
	return
}

func NewDescribeAssetOverviewResponse() (response *DescribeAssetOverviewResponse) {
	response = &DescribeAssetOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAssetOverview 资产防护概览
func (c *Client) DescribeAssetOverview(request *DescribeAssetOverviewRequest) (response *DescribeAssetOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeAssetOverviewRequest()
	}
	response = NewDescribeAssetOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewSyncFwOperateRequest() (request *SyncFwOperateRequest) {
	request = &SyncFwOperateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "SyncFwOperate")
	return
}

func NewSyncFwOperateResponse() (response *SyncFwOperateResponse) {
	response = &SyncFwOperateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步防火墙操作，包括同步防火墙路由（若vpc，专线网关等增加了Cidr，需要手动同步一下路由使之在防火墙上生效）等。
func (c *Client) SyncFwOperate(request *SyncFwOperateRequest) (response *SyncFwOperateResponse, err error) {
	if request == nil {
		request = NewSyncFwOperateRequest()
	}
	response = NewSyncFwOperateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnWithRouteRequest() (request *DescribeCcnWithRouteRequest) {
	request = &DescribeCcnWithRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCcnWithRoute")
	return
}

func NewDescribeCcnWithRouteResponse() (response *DescribeCcnWithRouteResponse) {
	response = &DescribeCcnWithRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回开白路由表的CCN列表
func (c *Client) DescribeCcnWithRoute(request *DescribeCcnWithRouteRequest) (response *DescribeCcnWithRouteResponse, err error) {
	if request == nil {
		request = NewDescribeCcnWithRouteRequest()
	}
	response = NewDescribeCcnWithRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSecurityGroupAllRuleRequest() (request *DeleteSecurityGroupAllRuleRequest) {
	request = &DeleteSecurityGroupAllRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteSecurityGroupAllRule")
	return
}

func NewDeleteSecurityGroupAllRuleResponse() (response *DeleteSecurityGroupAllRuleResponse) {
	response = &DeleteSecurityGroupAllRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除全部规则
func (c *Client) DeleteSecurityGroupAllRule(request *DeleteSecurityGroupAllRuleRequest) (response *DeleteSecurityGroupAllRuleResponse, err error) {
	if request == nil {
		request = NewDeleteSecurityGroupAllRuleRequest()
	}
	response = NewDeleteSecurityGroupAllRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcInstanceRequest() (request *DeleteVpcInstanceRequest) {
	request = &DeleteVpcInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteVpcInstance")
	return
}

func NewDeleteVpcInstanceResponse() (response *DeleteVpcInstanceResponse) {
	response = &DeleteVpcInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除防火墙实例
func (c *Client) DeleteVpcInstance(request *DeleteVpcInstanceRequest) (response *DeleteVpcInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteVpcInstanceRequest()
	}
	response = NewDeleteVpcInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcAcRuleRequest() (request *DescribeVpcAcRuleRequest) {
	request = &DescribeVpcAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcAcRule")
	return
}

func NewDescribeVpcAcRuleResponse() (response *DescribeVpcAcRuleResponse) {
	response = &DescribeVpcAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询内网间访问控制列表
func (c *Client) DescribeVpcAcRule(request *DescribeVpcAcRuleRequest) (response *DescribeVpcAcRuleResponse, err error) {
	if request == nil {
		request = NewDescribeVpcAcRuleRequest()
	}
	response = NewDescribeVpcAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTLogTopIpRequest() (request *DescribeTLogTopIpRequest) {
	request = &DescribeTLogTopIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeTLogTopIp")
	return
}

func NewDescribeTLogTopIpResponse() (response *DescribeTLogTopIpResponse) {
	response = &DescribeTLogTopIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTiTopIp 概览页IP攻击top柱状图
func (c *Client) DescribeTLogTopIp(request *DescribeTLogTopIpRequest) (response *DescribeTLogTopIpResponse, err error) {
	if request == nil {
		request = NewDescribeTLogTopIpRequest()
	}
	response = NewDescribeTLogTopIpResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGuideNetFlowSwitchAPIRequest() (request *DescribeGuideNetFlowSwitchAPIRequest) {
	request = &DescribeGuideNetFlowSwitchAPIRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeGuideNetFlowSwitchAPI")
	return
}

func NewDescribeGuideNetFlowSwitchAPIResponse() (response *DescribeGuideNetFlowSwitchAPIResponse) {
	response = &DescribeGuideNetFlowSwitchAPIResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新手引导互联网边界开关信息查询
func (c *Client) DescribeGuideNetFlowSwitchAPI(request *DescribeGuideNetFlowSwitchAPIRequest) (response *DescribeGuideNetFlowSwitchAPIResponse, err error) {
	if request == nil {
		request = NewDescribeGuideNetFlowSwitchAPIRequest()
	}
	response = NewDescribeGuideNetFlowSwitchAPIResponse()
	err = c.Send(request, response)
	return
}

func NewModifySequenceRulesRequest() (request *ModifySequenceRulesRequest) {
	request = &ModifySequenceRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifySequenceRules")
	return
}

func NewModifySequenceRulesResponse() (response *ModifySequenceRulesResponse) {
	response = &ModifySequenceRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改规则执行顺序
func (c *Client) ModifySequenceRules(request *ModifySequenceRulesRequest) (response *ModifySequenceRulesResponse, err error) {
	if request == nil {
		request = NewModifySequenceRulesRequest()
	}
	response = NewModifySequenceRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcFwReSelectRequest() (request *ModifyVpcFwReSelectRequest) {
	request = &ModifyVpcFwReSelectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcFwReSelect")
	return
}

func NewModifyVpcFwReSelectResponse() (response *ModifyVpcFwReSelectResponse) {
	response = &ModifyVpcFwReSelectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// vpc间防火墙重新选择vpc
func (c *Client) ModifyVpcFwReSelect(request *ModifyVpcFwReSelectRequest) (response *ModifyVpcFwReSelectResponse, err error) {
	if request == nil {
		request = NewModifyVpcFwReSelectRequest()
	}
	response = NewModifyVpcFwReSelectResponse()
	err = c.Send(request, response)
	return
}

func NewQueryUpdateResourceTaskStatusRequest() (request *QueryUpdateResourceTaskStatusRequest) {
	request = &QueryUpdateResourceTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "QueryUpdateResourceTaskStatus")
	return
}

func NewQueryUpdateResourceTaskStatusResponse() (response *QueryUpdateResourceTaskStatusResponse) {
	response = &QueryUpdateResourceTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户的资产同步任务状态
func (c *Client) QueryUpdateResourceTaskStatus(request *QueryUpdateResourceTaskStatusRequest) (response *QueryUpdateResourceTaskStatusResponse, err error) {
	if request == nil {
		request = NewQueryUpdateResourceTaskStatusRequest()
	}
	response = NewQueryUpdateResourceTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveOfflineExportTaskRequest() (request *RemoveOfflineExportTaskRequest) {
	request = &RemoveOfflineExportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "RemoveOfflineExportTask")
	return
}

func NewRemoveOfflineExportTaskResponse() (response *RemoveOfflineExportTaskResponse) {
	response = &RemoveOfflineExportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志离线导出任务
func (c *Client) RemoveOfflineExportTask(request *RemoveOfflineExportTaskRequest) (response *RemoveOfflineExportTaskResponse, err error) {
	if request == nil {
		request = NewRemoveOfflineExportTaskRequest()
	}
	response = NewRemoveOfflineExportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwRegionRequest() (request *DescribeVpcFwRegionRequest) {
	request = &DescribeVpcFwRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwRegion")
	return
}

func NewDescribeVpcFwRegionResponse() (response *DescribeVpcFwRegionResponse) {
	response = &DescribeVpcFwRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回防火墙所有的部署地域
func (c *Client) DescribeVpcFwRegion(request *DescribeVpcFwRegionRequest) (response *DescribeVpcFwRegionResponse, err error) {
	if request == nil {
		request = NewDescribeVpcFwRegionRequest()
	}
	response = NewDescribeVpcFwRegionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcAcRuleSwitchRequest() (request *ModifyVpcAcRuleSwitchRequest) {
	request = &ModifyVpcAcRuleSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcAcRuleSwitch")
	return
}

func NewModifyVpcAcRuleSwitchResponse() (response *ModifyVpcAcRuleSwitchResponse) {
	response = &ModifyVpcAcRuleSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量开启或者关闭内网间规则开关
func (c *Client) ModifyVpcAcRuleSwitch(request *ModifyVpcAcRuleSwitchRequest) (response *ModifyVpcAcRuleSwitchResponse, err error) {
	if request == nil {
		request = NewModifyVpcAcRuleSwitchRequest()
	}
	response = NewModifyVpcAcRuleSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlertApiDispatchRequest() (request *ModifyAlertApiDispatchRequest) {
	request = &ModifyAlertApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyAlertApiDispatch")
	return
}

func NewModifyAlertApiDispatchResponse() (response *ModifyAlertApiDispatchResponse) {
	response = &ModifyAlertApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警中心写接口请求中转
func (c *Client) ModifyAlertApiDispatch(request *ModifyAlertApiDispatchRequest) (response *ModifyAlertApiDispatchResponse, err error) {
	if request == nil {
		request = NewModifyAlertApiDispatchRequest()
	}
	response = NewModifyAlertApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFlowApiDispatchRequest() (request *ModifyFlowApiDispatchRequest) {
	request = &ModifyFlowApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyFlowApiDispatch")
	return
}

func NewModifyFlowApiDispatchResponse() (response *ModifyFlowApiDispatchResponse) {
	response = &ModifyFlowApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量中心写接口请求中转
func (c *Client) ModifyFlowApiDispatch(request *ModifyFlowApiDispatchRequest) (response *ModifyFlowApiDispatchResponse, err error) {
	if request == nil {
		request = NewModifyFlowApiDispatchRequest()
	}
	response = NewModifyFlowApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyStorageSettingRequest() (request *ModifyStorageSettingRequest) {
	request = &ModifyStorageSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyStorageSetting")
	return
}

func NewModifyStorageSettingResponse() (response *ModifyStorageSettingResponse) {
	response = &ModifyStorageSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志存储设置，可以修改存储时间和清空日志
func (c *Client) ModifyStorageSetting(request *ModifyStorageSettingRequest) (response *ModifyStorageSettingResponse, err error) {
	if request == nil {
		request = NewModifyStorageSettingRequest()
	}
	response = NewModifyStorageSettingResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperateLogsRequest() (request *CreateOperateLogsRequest) {
	request = &CreateOperateLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateOperateLogs")
	return
}

func NewCreateOperateLogsResponse() (response *CreateOperateLogsResponse) {
	response = &CreateOperateLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志上报
func (c *Client) CreateOperateLogs(request *CreateOperateLogsRequest) (response *CreateOperateLogsResponse, err error) {
	if request == nil {
		request = NewCreateOperateLogsRequest()
	}
	response = NewCreateOperateLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceGroupNewRequest() (request *DescribeResourceGroupNewRequest) {
	request = &DescribeResourceGroupNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeResourceGroupNew")
	return
}

func NewDescribeResourceGroupNewResponse() (response *DescribeResourceGroupNewResponse) {
	response = &DescribeResourceGroupNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeResourceGroupNew资产中心资产树信息
func (c *Client) DescribeResourceGroupNew(request *DescribeResourceGroupNewRequest) (response *DescribeResourceGroupNewResponse, err error) {
	if request == nil {
		request = NewDescribeResourceGroupNewRequest()
	}
	response = NewDescribeResourceGroupNewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoInternetSwitchRequest() (request *ModifyAutoInternetSwitchRequest) {
	request = &ModifyAutoInternetSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyAutoInternetSwitch")
	return
}

func NewModifyAutoInternetSwitchResponse() (response *ModifyAutoInternetSwitchResponse) {
	response = &ModifyAutoInternetSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改自动开启互联网边界开关
func (c *Client) ModifyAutoInternetSwitch(request *ModifyAutoInternetSwitchRequest) (response *ModifyAutoInternetSwitchResponse, err error) {
	if request == nil {
		request = NewModifyAutoInternetSwitchRequest()
	}
	response = NewModifyAutoInternetSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAcRuleRequest() (request *DeleteAcRuleRequest) {
	request = &DeleteAcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteAcRule")
	return
}

func NewDeleteAcRuleResponse() (response *DeleteAcRuleResponse) {
	response = &DeleteAcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除规则
func (c *Client) DeleteAcRule(request *DeleteAcRuleRequest) (response *DeleteAcRuleResponse, err error) {
	if request == nil {
		request = NewDeleteAcRuleRequest()
	}
	response = NewDeleteAcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebServicesRequest() (request *DescribeWebServicesRequest) {
	request = &DescribeWebServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeWebServices")
	return
}

func NewDescribeWebServicesResponse() (response *DescribeWebServicesResponse) {
	response = &DescribeWebServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产中心web服务查询列表
func (c *Client) DescribeWebServices(request *DescribeWebServicesRequest) (response *DescribeWebServicesResponse, err error) {
	if request == nil {
		request = NewDescribeWebServicesRequest()
	}
	response = NewDescribeWebServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcSwitchListsRequest() (request *DescribeVpcSwitchListsRequest) {
	request = &DescribeVpcSwitchListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcSwitchLists")
	return
}

func NewDescribeVpcSwitchListsResponse() (response *DescribeVpcSwitchListsResponse) {
	response = &DescribeVpcSwitchListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙开关-VPC边界防火墙-防火墙开关列表
func (c *Client) DescribeVpcSwitchLists(request *DescribeVpcSwitchListsRequest) (response *DescribeVpcSwitchListsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcSwitchListsRequest()
	}
	response = NewDescribeVpcSwitchListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIsolateNetCardRequest() (request *DescribeIsolateNetCardRequest) {
	request = &DescribeIsolateNetCardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeIsolateNetCard")
	return
}

func NewDescribeIsolateNetCardResponse() (response *DescribeIsolateNetCardResponse) {
	response = &DescribeIsolateNetCardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询隔离资产所有网卡
func (c *Client) DescribeIsolateNetCard(request *DescribeIsolateNetCardRequest) (response *DescribeIsolateNetCardResponse, err error) {
	if request == nil {
		request = NewDescribeIsolateNetCardRequest()
	}
	response = NewDescribeIsolateNetCardResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcFwGroupRequest() (request *CreateVpcFwGroupRequest) {
	request = &CreateVpcFwGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateVpcFwGroup")
	return
}

func NewCreateVpcFwGroupResponse() (response *CreateVpcFwGroupResponse) {
	response = &CreateVpcFwGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建VPC间防火墙(防火墙组)
func (c *Client) CreateVpcFwGroup(request *CreateVpcFwGroupRequest) (response *CreateVpcFwGroupResponse, err error) {
	if request == nil {
		request = NewCreateVpcFwGroupRequest()
	}
	response = NewCreateVpcFwGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNewAuthInfoRequest() (request *DescribeNewAuthInfoRequest) {
	request = &DescribeNewAuthInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeNewAuthInfo")
	return
}

func NewDescribeNewAuthInfoResponse() (response *DescribeNewAuthInfoResponse) {
	response = &DescribeNewAuthInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeNewAuthInfo 授权页复选框
func (c *Client) DescribeNewAuthInfo(request *DescribeNewAuthInfoRequest) (response *DescribeNewAuthInfoResponse, err error) {
	if request == nil {
		request = NewDescribeNewAuthInfoRequest()
	}
	response = NewDescribeNewAuthInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebCosUrlRequest() (request *DescribeWebCosUrlRequest) {
	request = &DescribeWebCosUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeWebCosUrl")
	return
}

func NewDescribeWebCosUrlResponse() (response *DescribeWebCosUrlResponse) {
	response = &DescribeWebCosUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询使用消息中心订阅推送的开关
func (c *Client) DescribeWebCosUrl(request *DescribeWebCosUrlRequest) (response *DescribeWebCosUrlResponse, err error) {
	if request == nil {
		request = NewDescribeWebCosUrlRequest()
	}
	response = NewDescribeWebCosUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGlobalSettingRequest() (request *DescribeGlobalSettingRequest) {
	request = &DescribeGlobalSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeGlobalSetting")
	return
}

func NewDescribeGlobalSettingResponse() (response *DescribeGlobalSettingResponse) {
	response = &DescribeGlobalSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全局配置，包含全产品或单用户灰度策略
func (c *Client) DescribeGlobalSetting(request *DescribeGlobalSettingRequest) (response *DescribeGlobalSettingResponse, err error) {
	if request == nil {
		request = NewDescribeGlobalSettingRequest()
	}
	response = NewDescribeGlobalSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcDetailRequest() (request *DescribeVpcDetailRequest) {
	request = &DescribeVpcDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcDetail")
	return
}

func NewDescribeVpcDetailResponse() (response *DescribeVpcDetailResponse) {
	response = &DescribeVpcDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户Vpc详情信息
func (c *Client) DescribeVpcDetail(request *DescribeVpcDetailRequest) (response *DescribeVpcDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVpcDetailRequest()
	}
	response = NewDescribeVpcDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiDispatchRequest() (request *ModifyApiDispatchRequest) {
	request = &ModifyApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyApiDispatch")
	return
}

func NewModifyApiDispatchResponse() (response *ModifyApiDispatchResponse) {
	response = &ModifyApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 写接口请求中转
func (c *Client) ModifyApiDispatch(request *ModifyApiDispatchRequest) (response *ModifyApiDispatchResponse, err error) {
	if request == nil {
		request = NewModifyApiDispatchRequest()
	}
	response = NewModifyApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpTcRuleRequest() (request *DescribeIpTcRuleRequest) {
	request = &DescribeIpTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeIpTcRule")
	return
}

func NewDescribeIpTcRuleResponse() (response *DescribeIpTcRuleResponse) {
	response = &DescribeIpTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询IP的流控策略
func (c *Client) DescribeIpTcRule(request *DescribeIpTcRuleRequest) (response *DescribeIpTcRuleResponse, err error) {
	if request == nil {
		request = NewDescribeIpTcRuleRequest()
	}
	response = NewDescribeIpTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCheckIpsRuleSwitchRequest() (request *CheckIpsRuleSwitchRequest) {
	request = &CheckIpsRuleSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CheckIpsRuleSwitch")
	return
}

func NewCheckIpsRuleSwitchResponse() (response *CheckIpsRuleSwitchResponse) {
	response = &CheckIpsRuleSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查是否支持IPS规则开关和动作
func (c *Client) CheckIpsRuleSwitch(request *CheckIpsRuleSwitchRequest) (response *CheckIpsRuleSwitchResponse, err error) {
	if request == nil {
		request = NewCheckIpsRuleSwitchRequest()
	}
	response = NewCheckIpsRuleSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBetaTaskRequest() (request *CreateBetaTaskRequest) {
	request = &CreateBetaTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateBetaTask")
	return
}

func NewCreateBetaTaskResponse() (response *CreateBetaTaskResponse) {
	response = &CreateBetaTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 建立Beta自动化任务
func (c *Client) CreateBetaTask(request *CreateBetaTaskRequest) (response *CreateBetaTaskResponse, err error) {
	if request == nil {
		request = NewCreateBetaTaskRequest()
	}
	response = NewCreateBetaTaskResponse()
	err = c.Send(request, response)
	return
}

func NewAddVpcFwTcRuleRequest() (request *AddVpcFwTcRuleRequest) {
	request = &AddVpcFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddVpcFwTcRule")
	return
}

func NewAddVpcFwTcRuleResponse() (response *AddVpcFwTcRuleResponse) {
	response = &AddVpcFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加流量控制策略
func (c *Client) AddVpcFwTcRule(request *AddVpcFwTcRuleRequest) (response *AddVpcFwTcRuleResponse, err error) {
	if request == nil {
		request = NewAddVpcFwTcRuleRequest()
	}
	response = NewAddVpcFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfwInsStatusRequest() (request *DescribeCfwInsStatusRequest) {
	request = &DescribeCfwInsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwInsStatus")
	return
}

func NewDescribeCfwInsStatusResponse() (response *DescribeCfwInsStatusResponse) {
	response = &DescribeCfwInsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// cfw实例运行状态查询
func (c *Client) DescribeCfwInsStatus(request *DescribeCfwInsStatusRequest) (response *DescribeCfwInsStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCfwInsStatusRequest()
	}
	response = NewDescribeCfwInsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAddNatFwTcRuleRequest() (request *AddNatFwTcRuleRequest) {
	request = &AddNatFwTcRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddNatFwTcRule")
	return
}

func NewAddNatFwTcRuleResponse() (response *AddNatFwTcRuleResponse) {
	response = &AddNatFwTcRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加流量控制策略
func (c *Client) AddNatFwTcRule(request *AddNatFwTcRuleRequest) (response *AddNatFwTcRuleResponse, err error) {
	if request == nil {
		request = NewAddNatFwTcRuleRequest()
	}
	response = NewAddNatFwTcRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
	request = &DescribeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeConfig")
	return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
	response = &DescribeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置信息
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
	if request == nil {
		request = NewDescribeConfigRequest()
	}
	response = NewDescribeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBlockIPBySGSwitchRequest() (request *ModifyBlockIPBySGSwitchRequest) {
	request = &ModifyBlockIPBySGSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockIPBySGSwitch")
	return
}

func NewModifyBlockIPBySGSwitchResponse() (response *ModifyBlockIPBySGSwitchResponse) {
	response = &ModifyBlockIPBySGSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改使用安全组封禁ip的开关
func (c *Client) ModifyBlockIPBySGSwitch(request *ModifyBlockIPBySGSwitchRequest) (response *ModifyBlockIPBySGSwitchResponse, err error) {
	if request == nil {
		request = NewModifyBlockIPBySGSwitchRequest()
	}
	response = NewModifyBlockIPBySGSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReplayUrlRequest() (request *DescribeReplayUrlRequest) {
	request = &DescribeReplayUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeReplayUrl")
	return
}

func NewDescribeReplayUrlResponse() (response *DescribeReplayUrlResponse) {
	response = &DescribeReplayUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取回放链接
func (c *Client) DescribeReplayUrl(request *DescribeReplayUrlRequest) (response *DescribeReplayUrlResponse, err error) {
	if request == nil {
		request = NewDescribeReplayUrlRequest()
	}
	response = NewDescribeReplayUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGuideUserStatusRequest() (request *DescribeGuideUserStatusRequest) {
	request = &DescribeGuideUserStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeGuideUserStatus")
	return
}

func NewDescribeGuideUserStatusResponse() (response *DescribeGuideUserStatusResponse) {
	response = &DescribeGuideUserStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeGuideUserStatus  新手引导完成标志位查询
func (c *Client) DescribeGuideUserStatus(request *DescribeGuideUserStatusRequest) (response *DescribeGuideUserStatusResponse, err error) {
	if request == nil {
		request = NewDescribeGuideUserStatusRequest()
	}
	response = NewDescribeGuideUserStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTLogIpListRequest() (request *DescribeTLogIpListRequest) {
	request = &DescribeTLogIpListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeTLogIpList")
	return
}

func NewDescribeTLogIpListResponse() (response *DescribeTLogIpListResponse) {
	response = &DescribeTLogIpListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTLogIpList告警中心IP柱形图
func (c *Client) DescribeTLogIpList(request *DescribeTLogIpListRequest) (response *DescribeTLogIpListResponse, err error) {
	if request == nil {
		request = NewDescribeTLogIpListRequest()
	}
	response = NewDescribeTLogIpListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeThreatInfoNewsRequest() (request *DescribeThreatInfoNewsRequest) {
	request = &DescribeThreatInfoNewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeThreatInfoNews")
	return
}

func NewDescribeThreatInfoNewsResponse() (response *DescribeThreatInfoNewsResponse) {
	response = &DescribeThreatInfoNewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 情报新闻
func (c *Client) DescribeThreatInfoNews(request *DescribeThreatInfoNewsRequest) (response *DescribeThreatInfoNewsResponse, err error) {
	if request == nil {
		request = NewDescribeThreatInfoNewsRequest()
	}
	response = NewDescribeThreatInfoNewsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCommonStatusRequest() (request *ModifyCommonStatusRequest) {
	request = &ModifyCommonStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyCommonStatus")
	return
}

func NewModifyCommonStatusResponse() (response *ModifyCommonStatusResponse) {
	response = &ModifyCommonStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通用异步改表接口
func (c *Client) ModifyCommonStatus(request *ModifyCommonStatusRequest) (response *ModifyCommonStatusResponse, err error) {
	if request == nil {
		request = NewModifyCommonStatusRequest()
	}
	response = NewModifyCommonStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGuideOneClickScanRequest() (request *CreateGuideOneClickScanRequest) {
	request = &CreateGuideOneClickScanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateGuideOneClickScan")
	return
}

func NewCreateGuideOneClickScanResponse() (response *CreateGuideOneClickScanResponse) {
	response = &CreateGuideOneClickScanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新手引导一键扫描
func (c *Client) CreateGuideOneClickScan(request *CreateGuideOneClickScanRequest) (response *CreateGuideOneClickScanResponse, err error) {
	if request == nil {
		request = NewCreateGuideOneClickScanRequest()
	}
	response = NewCreateGuideOneClickScanResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcFwInstanceRequest() (request *DeleteVpcFwInstanceRequest) {
	request = &DeleteVpcFwInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteVpcFwInstance")
	return
}

func NewDeleteVpcFwInstanceResponse() (response *DeleteVpcFwInstanceResponse) {
	response = &DeleteVpcFwInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁VPC间防火墙实例
func (c *Client) DeleteVpcFwInstance(request *DeleteVpcFwInstanceRequest) (response *DeleteVpcFwInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteVpcFwInstanceRequest()
	}
	response = NewDeleteVpcFwInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOnceClickScanRequest() (request *DescribeOnceClickScanRequest) {
	request = &DescribeOnceClickScanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeOnceClickScan")
	return
}

func NewDescribeOnceClickScanResponse() (response *DescribeOnceClickScanResponse) {
	response = &DescribeOnceClickScanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取一键体检用户信息
func (c *Client) DescribeOnceClickScan(request *DescribeOnceClickScanRequest) (response *DescribeOnceClickScanResponse, err error) {
	if request == nil {
		request = NewDescribeOnceClickScanRequest()
	}
	response = NewDescribeOnceClickScanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlackWhiteQuotaRequest() (request *DescribeBlackWhiteQuotaRequest) {
	request = &DescribeBlackWhiteQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlackWhiteQuota")
	return
}

func NewDescribeBlackWhiteQuotaResponse() (response *DescribeBlackWhiteQuotaResponse) {
	response = &DescribeBlackWhiteQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询黑白名单配额总接口
func (c *Client) DescribeBlackWhiteQuota(request *DescribeBlackWhiteQuotaRequest) (response *DescribeBlackWhiteQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeBlackWhiteQuotaRequest()
	}
	response = NewDescribeBlackWhiteQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAddressTemplateRequest() (request *DeleteAddressTemplateRequest) {
	request = &DeleteAddressTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteAddressTemplate")
	return
}

func NewDeleteAddressTemplateResponse() (response *DeleteAddressTemplateResponse) {
	response = &DeleteAddressTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除地址模板规则
func (c *Client) DeleteAddressTemplate(request *DeleteAddressTemplateRequest) (response *DeleteAddressTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteAddressTemplateRequest()
	}
	response = NewDeleteAddressTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiDispatchRequest() (request *DescribeApiDispatchRequest) {
	request = &DescribeApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeApiDispatch")
	return
}

func NewDescribeApiDispatchResponse() (response *DescribeApiDispatchResponse) {
	response = &DescribeApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口请求中转
func (c *Client) DescribeApiDispatch(request *DescribeApiDispatchRequest) (response *DescribeApiDispatchResponse, err error) {
	if request == nil {
		request = NewDescribeApiDispatchRequest()
	}
	response = NewDescribeApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowCenterAddressListsRequest() (request *DescribeFlowCenterAddressListsRequest) {
	request = &DescribeFlowCenterAddressListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFlowCenterAddressLists")
	return
}

func NewDescribeFlowCenterAddressListsResponse() (response *DescribeFlowCenterAddressListsResponse) {
	response = &DescribeFlowCenterAddressListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量中心外联地址列表查询
func (c *Client) DescribeFlowCenterAddressLists(request *DescribeFlowCenterAddressListsRequest) (response *DescribeFlowCenterAddressListsResponse, err error) {
	if request == nil {
		request = NewDescribeFlowCenterAddressListsRequest()
	}
	response = NewDescribeFlowCenterAddressListsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAddressTemplateRequest() (request *CreateAddressTemplateRequest) {
	request = &CreateAddressTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "CreateAddressTemplate")
	return
}

func NewCreateAddressTemplateResponse() (response *CreateAddressTemplateResponse) {
	response = &CreateAddressTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建地址模板规则
func (c *Client) CreateAddressTemplate(request *CreateAddressTemplateRequest) (response *CreateAddressTemplateResponse, err error) {
	if request == nil {
		request = NewCreateAddressTemplateRequest()
	}
	response = NewCreateAddressTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoBackUpSettingListRequest() (request *DescribeAutoBackUpSettingListRequest) {
	request = &DescribeAutoBackUpSettingListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAutoBackUpSettingList")
	return
}

func NewDescribeAutoBackUpSettingListResponse() (response *DescribeAutoBackUpSettingListResponse) {
	response = &DescribeAutoBackUpSettingListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取自动备份设置列表
func (c *Client) DescribeAutoBackUpSettingList(request *DescribeAutoBackUpSettingListRequest) (response *DescribeAutoBackUpSettingListResponse, err error) {
	if request == nil {
		request = NewDescribeAutoBackUpSettingListRequest()
	}
	response = NewDescribeAutoBackUpSettingListResponse()
	err = c.Send(request, response)
	return
}

func NewResetNatRuleHitTimesRequest() (request *ResetNatRuleHitTimesRequest) {
	request = &ResetNatRuleHitTimesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ResetNatRuleHitTimes")
	return
}

func NewResetNatRuleHitTimesResponse() (response *ResetNatRuleHitTimesResponse) {
	response = &ResetNatRuleHitTimesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置NAT防火墙访问控制规则的命中次数
func (c *Client) ResetNatRuleHitTimes(request *ResetNatRuleHitTimesRequest) (response *ResetNatRuleHitTimesResponse, err error) {
	if request == nil {
		request = NewResetNatRuleHitTimesRequest()
	}
	response = NewResetNatRuleHitTimesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAclApiDispatchRequest() (request *DescribeAclApiDispatchRequest) {
	request = &DescribeAclApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAclApiDispatch")
	return
}

func NewDescribeAclApiDispatchResponse() (response *DescribeAclApiDispatchResponse) {
	response = &DescribeAclApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 访问控制读接口请求中转
func (c *Client) DescribeAclApiDispatch(request *DescribeAclApiDispatchRequest) (response *DescribeAclApiDispatchResponse, err error) {
	if request == nil {
		request = NewDescribeAclApiDispatchRequest()
	}
	response = NewDescribeAclApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcLogEdgeRequest() (request *DescribeVpcLogEdgeRequest) {
	request = &DescribeVpcLogEdgeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcLogEdge")
	return
}

func NewDescribeVpcLogEdgeResponse() (response *DescribeVpcLogEdgeResponse) {
	response = &DescribeVpcLogEdgeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回可日志查询得VPC防火墙开关列表
func (c *Client) DescribeVpcLogEdge(request *DescribeVpcLogEdgeRequest) (response *DescribeVpcLogEdgeResponse, err error) {
	if request == nil {
		request = NewDescribeVpcLogEdgeRequest()
	}
	response = NewDescribeVpcLogEdgeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowStatisticsOverviewRequest() (request *DescribeFlowStatisticsOverviewRequest) {
	request = &DescribeFlowStatisticsOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFlowStatisticsOverview")
	return
}

func NewDescribeFlowStatisticsOverviewResponse() (response *DescribeFlowStatisticsOverviewResponse) {
	response = &DescribeFlowStatisticsOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量统计概览
func (c *Client) DescribeFlowStatisticsOverview(request *DescribeFlowStatisticsOverviewRequest) (response *DescribeFlowStatisticsOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeFlowStatisticsOverviewRequest()
	}
	response = NewDescribeFlowStatisticsOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockByIpTimesListRequest() (request *DescribeBlockByIpTimesListRequest) {
	request = &DescribeBlockByIpTimesListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockByIpTimesList")
	return
}

func NewDescribeBlockByIpTimesListResponse() (response *DescribeBlockByIpTimesListResponse) {
	response = &DescribeBlockByIpTimesListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeBlockByIpTimesList 告警中心阻断IP折线图
func (c *Client) DescribeBlockByIpTimesList(request *DescribeBlockByIpTimesListRequest) (response *DescribeBlockByIpTimesListResponse, err error) {
	if request == nil {
		request = NewDescribeBlockByIpTimesListRequest()
	}
	response = NewDescribeBlockByIpTimesListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSecurityGroupRuleRequest() (request *DeleteSecurityGroupRuleRequest) {
	request = &DeleteSecurityGroupRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteSecurityGroupRule")
	return
}

func NewDeleteSecurityGroupRuleResponse() (response *DeleteSecurityGroupRuleResponse) {
	response = &DeleteSecurityGroupRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除规则
func (c *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) (response *DeleteSecurityGroupRuleResponse, err error) {
	if request == nil {
		request = NewDeleteSecurityGroupRuleRequest()
	}
	response = NewDeleteSecurityGroupRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssociatedInstanceListRequest() (request *DescribeAssociatedInstanceListRequest) {
	request = &DescribeAssociatedInstanceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssociatedInstanceList")
	return
}

func NewDescribeAssociatedInstanceListResponse() (response *DescribeAssociatedInstanceListResponse) {
	response = &DescribeAssociatedInstanceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全组关联实例列表
func (c *Client) DescribeAssociatedInstanceList(request *DescribeAssociatedInstanceListRequest) (response *DescribeAssociatedInstanceListResponse, err error) {
	if request == nil {
		request = NewDescribeAssociatedInstanceListRequest()
	}
	response = NewDescribeAssociatedInstanceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlertApiDispatchRequest() (request *DescribeAlertApiDispatchRequest) {
	request = &DescribeAlertApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeAlertApiDispatch")
	return
}

func NewDescribeAlertApiDispatchResponse() (response *DescribeAlertApiDispatchResponse) {
	response = &DescribeAlertApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警中心读接口请求中转
func (c *Client) DescribeAlertApiDispatch(request *DescribeAlertApiDispatchRequest) (response *DescribeAlertApiDispatchResponse, err error) {
	if request == nil {
		request = NewDescribeAlertApiDispatchRequest()
	}
	response = NewDescribeAlertApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewExpandCfwVerticalRequest() (request *ExpandCfwVerticalRequest) {
	request = &ExpandCfwVerticalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ExpandCfwVertical")
	return
}

func NewExpandCfwVerticalResponse() (response *ExpandCfwVerticalResponse) {
	response = &ExpandCfwVerticalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防火墙垂直扩容
func (c *Client) ExpandCfwVertical(request *ExpandCfwVerticalRequest) (response *ExpandCfwVerticalResponse, err error) {
	if request == nil {
		request = NewExpandCfwVerticalRequest()
	}
	response = NewExpandCfwVerticalResponse()
	err = c.Send(request, response)
	return
}

func NewRunSyncAssetRequest() (request *RunSyncAssetRequest) {
	request = &RunSyncAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "RunSyncAsset")
	return
}

func NewRunSyncAssetResponse() (response *RunSyncAssetResponse) {
	response = &RunSyncAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步资产-互联网&VPC
func (c *Client) RunSyncAsset(request *RunSyncAssetRequest) (response *RunSyncAssetResponse, err error) {
	if request == nil {
		request = NewRunSyncAssetRequest()
	}
	response = NewRunSyncAssetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatFwInstanceRequest() (request *ModifyNatFwInstanceRequest) {
	request = &ModifyNatFwInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwInstance")
	return
}

func NewModifyNatFwInstanceResponse() (response *ModifyNatFwInstanceResponse) {
	response = &ModifyNatFwInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNatFwInstance(request *ModifyNatFwInstanceRequest) (response *ModifyNatFwInstanceResponse, err error) {
	if request == nil {
		request = NewModifyNatFwInstanceRequest()
	}
	response = NewModifyNatFwInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNatInstanceRequest() (request *DeleteNatInstanceRequest) {
	request = &DeleteNatInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteNatFwInstance")
	return
}

func NewDeleteNatInstanceResponse() (response *DeleteNatInstanceResponse) {
	response = &DeleteNatInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNatInstance(request *DeleteNatInstanceRequest) (response *DeleteNatInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteNatInstanceRequest()
	}
	response = NewDeleteNatInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcFwGroupInfoRequest() (request *DescribeFwGroupInstanceInfoRequest) {
	request = &DescribeFwGroupInstanceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeFwGroupInstanceInfo")
	return
}

func NewDescribeVpcFwGroupInfoResponse() (response *DescribeFwGroupInstanceInfoResponse) {
	response = &DescribeFwGroupInstanceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func NewDeleteBlockIgnoreRuleListRequest() (request *DeleteBlockIgnoreRuleListRequest) {
	request = &DeleteBlockIgnoreRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cfw", APIVersion, "DeleteBlockIgnoreRuleList")

	return
}

func NewDeleteBlockIgnoreRuleListResponse() (response *DeleteBlockIgnoreRuleListResponse) {
	response = &DeleteBlockIgnoreRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

func (c *Client) DeleteBlockIgnoreRuleList(request *DeleteBlockIgnoreRuleListRequest) (response *DeleteBlockIgnoreRuleListResponse, err error) {
	if request == nil {
		request = NewDeleteBlockIgnoreRuleListRequest()
	}

	response = NewDeleteBlockIgnoreRuleListResponse()
	err = c.Send(request, response)
	return
}
