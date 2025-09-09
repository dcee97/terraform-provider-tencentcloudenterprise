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

package v20221121

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2022-11-21"

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

func NewDescribeRepositoryImageAssetsRequest() (request *DescribeRepositoryImageAssetsRequest) {
	request = &DescribeRepositoryImageAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRepositoryImageAssets")
	return
}

func NewDescribeRepositoryImageAssetsResponse() (response *DescribeRepositoryImageAssetsResponse) {
	response = &DescribeRepositoryImageAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 仓库镜像列表
func (c *Client) DescribeRepositoryImageAssets(request *DescribeRepositoryImageAssetsRequest) (response *DescribeRepositoryImageAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeRepositoryImageAssetsRequest()
	}
	response = NewDescribeRepositoryImageAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcConnectedVpcListsRequest() (request *DescribeVpcConnectedVpcListsRequest) {
	request = &DescribeVpcConnectedVpcListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeVpcConnectedVpcLists")
	return
}

func NewDescribeVpcConnectedVpcListsResponse() (response *DescribeVpcConnectedVpcListsResponse) {
	response = &DescribeVpcConnectedVpcListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Vpc连通其他vpc列表
func (c *Client) DescribeVpcConnectedVpcLists(request *DescribeVpcConnectedVpcListsRequest) (response *DescribeVpcConnectedVpcListsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcConnectedVpcListsRequest()
	}
	response = NewDescribeVpcConnectedVpcListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCenterAssetRiskMaxRequest() (request *DescribeCenterAssetRiskMaxRequest) {
	request = &DescribeCenterAssetRiskMaxRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCenterAssetRiskMax")
	return
}

func NewDescribeCenterAssetRiskMaxResponse() (response *DescribeCenterAssetRiskMaxResponse) {
	response = &DescribeCenterAssetRiskMaxResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询各个资产的最大风险量
func (c *Client) DescribeCenterAssetRiskMax(request *DescribeCenterAssetRiskMaxRequest) (response *DescribeCenterAssetRiskMaxResponse, err error) {
	if request == nil {
		request = NewDescribeCenterAssetRiskMaxRequest()
	}
	response = NewDescribeCenterAssetRiskMaxResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBASTaskRequest() (request *DescribeBASTaskRequest) {
	request = &DescribeBASTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeBASTask")
	return
}

func NewDescribeBASTaskResponse() (response *DescribeBASTaskResponse) {
	response = &DescribeBASTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bas任务列表
func (c *Client) DescribeBASTask(request *DescribeBASTaskRequest) (response *DescribeBASTaskResponse, err error) {
	if request == nil {
		request = NewDescribeBASTaskRequest()
	}
	response = NewDescribeBASTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCVMAssetsRequest() (request *DescribeCVMAssetsRequest) {
	request = &DescribeCVMAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCVMAssets")
	return
}

func NewDescribeCVMAssetsResponse() (response *DescribeCVMAssetsResponse) {
	response = &DescribeCVMAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// cvm列表
func (c *Client) DescribeCVMAssets(request *DescribeCVMAssetsRequest) (response *DescribeCVMAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeCVMAssetsRequest()
	}
	response = NewDescribeCVMAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebServiceAssetsListRequest() (request *DescribeWebServiceAssetsListRequest) {
	request = &DescribeWebServiceAssetsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWebServiceAssetsList")
	return
}

func NewDescribeWebServiceAssetsListResponse() (response *DescribeWebServiceAssetsListResponse) {
	response = &DescribeWebServiceAssetsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取web服务列表
func (c *Client) DescribeWebServiceAssetsList(request *DescribeWebServiceAssetsListRequest) (response *DescribeWebServiceAssetsListResponse, err error) {
	if request == nil {
		request = NewDescribeWebServiceAssetsListRequest()
	}
	response = NewDescribeWebServiceAssetsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuditDiversityRequest() (request *DescribeAuditDiversityRequest) {
	request = &DescribeAuditDiversityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAuditDiversity")
	return
}

func NewDescribeAuditDiversityResponse() (response *DescribeAuditDiversityResponse) {
	response = &DescribeAuditDiversityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志分类
func (c *Client) DescribeAuditDiversity(request *DescribeAuditDiversityRequest) (response *DescribeAuditDiversityResponse, err error) {
	if request == nil {
		request = NewDescribeAuditDiversityRequest()
	}
	response = NewDescribeAuditDiversityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClbLoadBalancersRequest() (request *DescribeClbLoadBalancersRequest) {
	request = &DescribeClbLoadBalancersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeClbLoadBalancers")
	return
}

func NewDescribeClbLoadBalancersResponse() (response *DescribeClbLoadBalancersResponse) {
	response = &DescribeClbLoadBalancersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb接口 获取负载均衡列表
func (c *Client) DescribeClbLoadBalancers(request *DescribeClbLoadBalancersRequest) (response *DescribeClbLoadBalancersResponse, err error) {
	if request == nil {
		request = NewDescribeClbLoadBalancersRequest()
	}
	response = NewDescribeClbLoadBalancersResponse()
	err = c.Send(request, response)
	return
}

func NewInstallBASAssetAgentByCWPRequest() (request *InstallBASAssetAgentByCWPRequest) {
	request = &InstallBASAssetAgentByCWPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "InstallBASAssetAgentByCWP")
	return
}

func NewInstallBASAssetAgentByCWPResponse() (response *InstallBASAssetAgentByCWPResponse) {
	response = &InstallBASAssetAgentByCWPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过主机agent安装模拟攻击工具
func (c *Client) InstallBASAssetAgentByCWP(request *InstallBASAssetAgentByCWPRequest) (response *InstallBASAssetAgentByCWPResponse, err error) {
	if request == nil {
		request = NewInstallBASAssetAgentByCWPRequest()
	}
	response = NewInstallBASAssetAgentByCWPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetStatTopXRequest() (request *DescribeAssetStatTopXRequest) {
	request = &DescribeAssetStatTopXRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetStatTopX")
	return
}

func NewDescribeAssetStatTopXResponse() (response *DescribeAssetStatTopXResponse) {
	response = &DescribeAssetStatTopXResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 风险统计项数据
func (c *Client) DescribeAssetStatTopX(request *DescribeAssetStatTopXRequest) (response *DescribeAssetStatTopXResponse, err error) {
	if request == nil {
		request = NewDescribeAssetStatTopXRequest()
	}
	response = NewDescribeAssetStatTopXResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainAssetsRequest() (request *DescribeDomainAssetsRequest) {
	request = &DescribeDomainAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeDomainAssets")
	return
}

func NewDescribeDomainAssetsResponse() (response *DescribeDomainAssetsResponse) {
	response = &DescribeDomainAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 域名列表
func (c *Client) DescribeDomainAssets(request *DescribeDomainAssetsRequest) (response *DescribeDomainAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeDomainAssetsRequest()
	}
	response = NewDescribeDomainAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskClassInfoRequest() (request *DescribeRiskClassInfoRequest) {
	request = &DescribeRiskClassInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskClassInfo")
	return
}

func NewDescribeRiskClassInfoResponse() (response *DescribeRiskClassInfoResponse) {
	response = &DescribeRiskClassInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取风险分类信息
func (c *Client) DescribeRiskClassInfo(request *DescribeRiskClassInfoRequest) (response *DescribeRiskClassInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRiskClassInfoRequest()
	}
	response = NewDescribeRiskClassInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSocLogStorageSpaceRequest() (request *DescribeSocLogStorageSpaceRequest) {
	request = &DescribeSocLogStorageSpaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSocLogStorageSpace")
	return
}

func NewDescribeSocLogStorageSpaceResponse() (response *DescribeSocLogStorageSpaceResponse) {
	response = &DescribeSocLogStorageSpaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 投递日志空间占比
func (c *Client) DescribeSocLogStorageSpace(request *DescribeSocLogStorageSpaceRequest) (response *DescribeSocLogStorageSpaceResponse, err error) {
	if request == nil {
		request = NewDescribeSocLogStorageSpaceRequest()
	}
	response = NewDescribeSocLogStorageSpaceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserCvmInfoRequest() (request *DescribeUserCvmInfoRequest) {
	request = &DescribeUserCvmInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeUserCvmInfo")
	return
}

func NewDescribeUserCvmInfoResponse() (response *DescribeUserCvmInfoResponse) {
	response = &DescribeUserCvmInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户cvm信息
func (c *Client) DescribeUserCvmInfo(request *DescribeUserCvmInfoRequest) (response *DescribeUserCvmInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserCvmInfoRequest()
	}
	response = NewDescribeUserCvmInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNICAssetsRequest() (request *DescribeNICAssetsRequest) {
	request = &DescribeNICAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeNICAssets")
	return
}

func NewDescribeNICAssetsResponse() (response *DescribeNICAssetsResponse) {
	response = &DescribeNICAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网卡列表
func (c *Client) DescribeNICAssets(request *DescribeNICAssetsRequest) (response *DescribeNICAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeNICAssetsRequest()
	}
	response = NewDescribeNICAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubnetDetailInfoRequest() (request *DescribeSubnetDetailInfoRequest) {
	request = &DescribeSubnetDetailInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSubnetDetailInfo")
	return
}

func NewDescribeSubnetDetailInfoResponse() (response *DescribeSubnetDetailInfoResponse) {
	response = &DescribeSubnetDetailInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 子网详情
func (c *Client) DescribeSubnetDetailInfo(request *DescribeSubnetDetailInfoRequest) (response *DescribeSubnetDetailInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSubnetDetailInfoRequest()
	}
	response = NewDescribeSubnetDetailInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskPredictTimeRequest() (request *DescribeTaskPredictTimeRequest) {
	request = &DescribeTaskPredictTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskPredictTime")
	return
}

func NewDescribeTaskPredictTimeResponse() (response *DescribeTaskPredictTimeResponse) {
	response = &DescribeTaskPredictTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务预估完成时间
func (c *Client) DescribeTaskPredictTime(request *DescribeTaskPredictTimeRequest) (response *DescribeTaskPredictTimeResponse, err error) {
	if request == nil {
		request = NewDescribeTaskPredictTimeRequest()
	}
	response = NewDescribeTaskPredictTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageBuildRequest() (request *DescribeAssetImageBuildRequest) {
	request = &DescribeAssetImageBuildRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetImageBuild")
	return
}

func NewDescribeAssetImageBuildResponse() (response *DescribeAssetImageBuildResponse) {
	response = &DescribeAssetImageBuildResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 构建列表
func (c *Client) DescribeAssetImageBuild(request *DescribeAssetImageBuildRequest) (response *DescribeAssetImageBuildResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageBuildRequest()
	}
	response = NewDescribeAssetImageBuildResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNICAssetInfoRequest() (request *DescribeNICAssetInfoRequest) {
	request = &DescribeNICAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeNICAssetInfo")
	return
}

func NewDescribeNICAssetInfoResponse() (response *DescribeNICAssetInfoResponse) {
	response = &DescribeNICAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网卡资产详情
func (c *Client) DescribeNICAssetInfo(request *DescribeNICAssetInfoRequest) (response *DescribeNICAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeNICAssetInfoRequest()
	}
	response = NewDescribeNICAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSolidProtectionDataRequest() (request *DescribeSolidProtectionDataRequest) {
	request = &DescribeSolidProtectionDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSolidProtectionData")
	return
}

func NewDescribeSolidProtectionDataResponse() (response *DescribeSolidProtectionDataResponse) {
	response = &DescribeSolidProtectionDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询立体防护中心数据接口
func (c *Client) DescribeSolidProtectionData(request *DescribeSolidProtectionDataRequest) (response *DescribeSolidProtectionDataResponse, err error) {
	if request == nil {
		request = NewDescribeSolidProtectionDataRequest()
	}
	response = NewDescribeSolidProtectionDataResponse()
	err = c.Send(request, response)
	return
}

func NewSyncAssetsProtectStatusRequest() (request *SyncAssetsProtectStatusRequest) {
	request = &SyncAssetsProtectStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "SyncAssetsProtectStatus")
	return
}

func NewSyncAssetsProtectStatusResponse() (response *SyncAssetsProtectStatusResponse) {
	response = &SyncAssetsProtectStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步资产防护状态
func (c *Client) SyncAssetsProtectStatus(request *SyncAssetsProtectStatusRequest) (response *SyncAssetsProtectStatusResponse, err error) {
	if request == nil {
		request = NewSyncAssetsProtectStatusRequest()
	}
	response = NewSyncAssetsProtectStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCanFixMachineRequest() (request *DescribeCanFixMachineRequest) {
	request = &DescribeCanFixMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCanFixMachine")
	return
}

func NewDescribeCanFixMachineResponse() (response *DescribeCanFixMachineResponse) {
	response = &DescribeCanFixMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口可以查询漏扫扫描存在漏洞的主机和主机安全可以修复的资产交集
func (c *Client) DescribeCanFixMachine(request *DescribeCanFixMachineRequest) (response *DescribeCanFixMachineResponse, err error) {
	if request == nil {
		request = NewDescribeCanFixMachineRequest()
	}
	response = NewDescribeCanFixMachineResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIgnoreOrRetryScanTaskRequest() (request *ModifyIgnoreOrRetryScanTaskRequest) {
	request = &ModifyIgnoreOrRetryScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyIgnoreOrRetryScanTask")
	return
}

func NewModifyIgnoreOrRetryScanTaskResponse() (response *ModifyIgnoreOrRetryScanTaskResponse) {
	response = &ModifyIgnoreOrRetryScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 异常任务扫描忽略或者重试
func (c *Client) ModifyIgnoreOrRetryScanTask(request *ModifyIgnoreOrRetryScanTaskRequest) (response *ModifyIgnoreOrRetryScanTaskResponse, err error) {
	if request == nil {
		request = NewModifyIgnoreOrRetryScanTaskRequest()
	}
	response = NewModifyIgnoreOrRetryScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTkeAuthRequest() (request *CreateTkeAuthRequest) {
	request = &CreateTkeAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateTkeAuth")
	return
}

func NewCreateTkeAuthResponse() (response *CreateTkeAuthResponse) {
	response = &CreateTkeAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在用户将tke授权完之后调用该接口以保存用户tke授权状态
func (c *Client) CreateTkeAuth(request *CreateTkeAuthRequest) (response *CreateTkeAuthResponse, err error) {
	if request == nil {
		request = NewCreateTkeAuthRequest()
	}
	response = NewCreateTkeAuthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafCertificateVerifyResultRequest() (request *DescribeWafCertificateVerifyResultRequest) {
	request = &DescribeWafCertificateVerifyResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWafCertificateVerifyResult")
	return
}

func NewDescribeWafCertificateVerifyResultResponse() (response *DescribeWafCertificateVerifyResultResponse) {
	response = &DescribeWafCertificateVerifyResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeCertificateVerifyResult waf接口 获取证书状态
func (c *Client) DescribeWafCertificateVerifyResult(request *DescribeWafCertificateVerifyResultRequest) (response *DescribeWafCertificateVerifyResultResponse, err error) {
	if request == nil {
		request = NewDescribeWafCertificateVerifyResultRequest()
	}
	response = NewDescribeWafCertificateVerifyResultResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogAccessSwitchRequest() (request *ModifyLogAccessSwitchRequest) {
	request = &ModifyLogAccessSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyLogAccessSwitch")
	return
}

func NewModifyLogAccessSwitchResponse() (response *ModifyLogAccessSwitchResponse) {
	response = &ModifyLogAccessSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户日志接入开关
func (c *Client) ModifyLogAccessSwitch(request *ModifyLogAccessSwitchRequest) (response *ModifyLogAccessSwitchResponse, err error) {
	if request == nil {
		request = NewModifyLogAccessSwitchRequest()
	}
	response = NewModifyLogAccessSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterCFGViewCFGRiskListRequest() (request *DescribeRiskCenterCFGViewCFGRiskListRequest) {
	request = &DescribeRiskCenterCFGViewCFGRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterCFGViewCFGRiskList")
	return
}

func NewDescribeRiskCenterCFGViewCFGRiskListResponse() (response *DescribeRiskCenterCFGViewCFGRiskListResponse) {
	response = &DescribeRiskCenterCFGViewCFGRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置视角的配置风险列表
func (c *Client) DescribeRiskCenterCFGViewCFGRiskList(request *DescribeRiskCenterCFGViewCFGRiskListRequest) (response *DescribeRiskCenterCFGViewCFGRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterCFGViewCFGRiskListRequest()
	}
	response = NewDescribeRiskCenterCFGViewCFGRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetVpcDetailInfoRequest() (request *DescribeAssetVpcDetailInfoRequest) {
	request = &DescribeAssetVpcDetailInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetVpcDetailInfo")
	return
}

func NewDescribeAssetVpcDetailInfoResponse() (response *DescribeAssetVpcDetailInfoResponse) {
	response = &DescribeAssetVpcDetailInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// vpc资产详情
func (c *Client) DescribeAssetVpcDetailInfo(request *DescribeAssetVpcDetailInfoRequest) (response *DescribeAssetVpcDetailInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetVpcDetailInfoRequest()
	}
	response = NewDescribeAssetVpcDetailInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseBindScheduleRequest() (request *DescribeLicenseBindScheduleRequest) {
	request = &DescribeLicenseBindScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeLicenseBindSchedule")
	return
}

func NewDescribeLicenseBindScheduleResponse() (response *DescribeLicenseBindScheduleResponse) {
	response = &DescribeLicenseBindScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询授权绑定任务的进度
func (c *Client) DescribeLicenseBindSchedule(request *DescribeLicenseBindScheduleRequest) (response *DescribeLicenseBindScheduleResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseBindScheduleRequest()
	}
	response = NewDescribeLicenseBindScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanAssetByTaskIdRequest() (request *DescribeScanAssetByTaskIdRequest) {
	request = &DescribeScanAssetByTaskIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeScanAssetByTaskId")
	return
}

func NewDescribeScanAssetByTaskIdResponse() (response *DescribeScanAssetByTaskIdResponse) {
	response = &DescribeScanAssetByTaskIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据taskId查询资产信息
func (c *Client) DescribeScanAssetByTaskId(request *DescribeScanAssetByTaskIdRequest) (response *DescribeScanAssetByTaskIdResponse, err error) {
	if request == nil {
		request = NewDescribeScanAssetByTaskIdRequest()
	}
	response = NewDescribeScanAssetByTaskIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterServerRiskListRequest() (request *DescribeRiskCenterServerRiskListRequest) {
	request = &DescribeRiskCenterServerRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterServerRiskList")
	return
}

func NewDescribeRiskCenterServerRiskListResponse() (response *DescribeRiskCenterServerRiskListResponse) {
	response = &DescribeRiskCenterServerRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取风险服务列表
func (c *Client) DescribeRiskCenterServerRiskList(request *DescribeRiskCenterServerRiskListRequest) (response *DescribeRiskCenterServerRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterServerRiskListRequest()
	}
	response = NewDescribeRiskCenterServerRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDiversityRequest() (request *DescribeAssetDiversityRequest) {
	request = &DescribeAssetDiversityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetDiversity")
	return
}

func NewDescribeAssetDiversityResponse() (response *DescribeAssetDiversityResponse) {
	response = &DescribeAssetDiversityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产统计分类
func (c *Client) DescribeAssetDiversity(request *DescribeAssetDiversityRequest) (response *DescribeAssetDiversityResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDiversityRequest()
	}
	response = NewDescribeAssetDiversityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationRequest() (request *DescribeOrganizationRequest) {
	request = &DescribeOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeOrganization")
	return
}

func NewDescribeOrganizationResponse() (response *DescribeOrganizationResponse) {
	response = &DescribeOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集团账号开通情况
func (c *Client) DescribeOrganization(request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationRequest()
	}
	response = NewDescribeOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePortProtectStatusRequest() (request *DescribePortProtectStatusRequest) {
	request = &DescribePortProtectStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribePortProtectStatus")
	return
}

func NewDescribePortProtectStatusResponse() (response *DescribePortProtectStatusResponse) {
	response = &DescribePortProtectStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询端口风险cfw端口状态任务
func (c *Client) DescribePortProtectStatus(request *DescribePortProtectStatusRequest) (response *DescribePortProtectStatusResponse, err error) {
	if request == nil {
		request = NewDescribePortProtectStatusRequest()
	}
	response = NewDescribePortProtectStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNeedToSolveListRequest() (request *DescribeNeedToSolveListRequest) {
	request = &DescribeNeedToSolveListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeNeedToSolveList")
	return
}

func NewDescribeNeedToSolveListResponse() (response *DescribeNeedToSolveListResponse) {
	response = &DescribeNeedToSolveListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询展示待办事件
func (c *Client) DescribeNeedToSolveList(request *DescribeNeedToSolveListRequest) (response *DescribeNeedToSolveListResponse, err error) {
	if request == nil {
		request = NewDescribeNeedToSolveListRequest()
	}
	response = NewDescribeNeedToSolveListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
	request = &DescribeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeConfig")
	return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
	response = &DescribeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防火墙配置信息
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
	if request == nil {
		request = NewDescribeConfigRequest()
	}
	response = NewDescribeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogAccessRequest() (request *DescribeLogAccessRequest) {
	request = &DescribeLogAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeLogAccess")
	return
}

func NewDescribeLogAccessResponse() (response *DescribeLogAccessResponse) {
	response = &DescribeLogAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志接入情况接口
func (c *Client) DescribeLogAccess(request *DescribeLogAccessRequest) (response *DescribeLogAccessResponse, err error) {
	if request == nil {
		request = NewDescribeLogAccessRequest()
	}
	response = NewDescribeLogAccessResponse()
	err = c.Send(request, response)
	return
}

func NewStopBASTaskRequest() (request *StopBASTaskRequest) {
	request = &StopBASTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "StopBASTask")
	return
}

func NewStopBASTaskResponse() (response *StopBASTaskResponse) {
	response = &StopBASTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止bas任务
func (c *Client) StopBASTask(request *StopBASTaskRequest) (response *StopBASTaskResponse, err error) {
	if request == nil {
		request = NewStopBASTaskRequest()
	}
	response = NewStopBASTaskResponse()
	err = c.Send(request, response)
	return
}

func NewExportCSVRequest() (request *ExportCSVRequest) {
	request = &ExportCSVRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ExportCSV")
	return
}

func NewExportCSVResponse() (response *ExportCSVResponse) {
	response = &ExportCSVResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取导出文件的临时下载链接
func (c *Client) ExportCSV(request *ExportCSVRequest) (response *ExportCSVResponse, err error) {
	if request == nil {
		request = NewExportCSVRequest()
	}
	response = NewExportCSVResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetRequest() (request *ModifyAssetRequest) {
	request = &ModifyAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyAsset")
	return
}

func NewModifyAssetResponse() (response *ModifyAssetResponse) {
	response = &ModifyAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标记资产为非核心资产
func (c *Client) ModifyAsset(request *ModifyAssetRequest) (response *ModifyAssetResponse, err error) {
	if request == nil {
		request = NewModifyAssetRequest()
	}
	response = NewModifyAssetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafSaasPortsRequest() (request *DescribeWafSaasPortsRequest) {
	request = &DescribeWafSaasPortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWafSaasPorts")
	return
}

func NewDescribeWafSaasPortsResponse() (response *DescribeWafSaasPortsResponse) {
	response = &DescribeWafSaasPortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribePorts waf接口 获取http端口列表
func (c *Client) DescribeWafSaasPorts(request *DescribeWafSaasPortsRequest) (response *DescribeWafSaasPortsResponse, err error) {
	if request == nil {
		request = NewDescribeWafSaasPortsRequest()
	}
	response = NewDescribeWafSaasPortsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCVMAssetInfoRequest() (request *DescribeCVMAssetInfoRequest) {
	request = &DescribeCVMAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCVMAssetInfo")
	return
}

func NewDescribeCVMAssetInfoResponse() (response *DescribeCVMAssetInfoResponse) {
	response = &DescribeCVMAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// cvm详情
func (c *Client) DescribeCVMAssetInfo(request *DescribeCVMAssetInfoRequest) (response *DescribeCVMAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeCVMAssetInfoRequest()
	}
	response = NewDescribeCVMAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLastTaskOverviewRequest() (request *DescribeLastTaskOverviewRequest) {
	request = &DescribeLastTaskOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeLastTaskOverview")
	return
}

func NewDescribeLastTaskOverviewResponse() (response *DescribeLastTaskOverviewResponse) {
	response = &DescribeLastTaskOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取最近任务任务列表概览
func (c *Client) DescribeLastTaskOverview(request *DescribeLastTaskOverviewRequest) (response *DescribeLastTaskOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeLastTaskOverviewRequest()
	}
	response = NewDescribeLastTaskOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskAssetTopFiveListRequest() (request *DescribeRiskAssetTopFiveListRequest) {
	request = &DescribeRiskAssetTopFiveListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskAssetTopFiveList")
	return
}

func NewDescribeRiskAssetTopFiveListResponse() (response *DescribeRiskAssetTopFiveListResponse) {
	response = &DescribeRiskAssetTopFiveListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取风险资产TOP5
func (c *Client) DescribeRiskAssetTopFiveList(request *DescribeRiskAssetTopFiveListRequest) (response *DescribeRiskAssetTopFiveListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskAssetTopFiveListRequest()
	}
	response = NewDescribeRiskAssetTopFiveListResponse()
	err = c.Send(request, response)
	return
}

func NewExitPocScanMissionRequest() (request *ExitPocScanMissionRequest) {
	request = &ExitPocScanMissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ExitPocScanMission")
	return
}

func NewExitPocScanMissionResponse() (response *ExitPocScanMissionResponse) {
	response = &ExitPocScanMissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 退出立体防护漏洞扫描
func (c *Client) ExitPocScanMission(request *ExitPocScanMissionRequest) (response *ExitPocScanMissionResponse, err error) {
	if request == nil {
		request = NewExitPocScanMissionRequest()
	}
	response = NewExitPocScanMissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanTaskListRequest() (request *DescribeScanTaskListRequest) {
	request = &DescribeScanTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeScanTaskList")
	return
}

func NewDescribeScanTaskListResponse() (response *DescribeScanTaskListResponse) {
	response = &DescribeScanTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取扫描任务列表
func (c *Client) DescribeScanTaskList(request *DescribeScanTaskListRequest) (response *DescribeScanTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeScanTaskListRequest()
	}
	response = NewDescribeScanTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskReportOverviewRequest() (request *DescribeTaskReportOverviewRequest) {
	request = &DescribeTaskReportOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskReportOverview")
	return
}

func NewDescribeTaskReportOverviewResponse() (response *DescribeTaskReportOverviewResponse) {
	response = &DescribeTaskReportOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务和报告列表概览数据
func (c *Client) DescribeTaskReportOverview(request *DescribeTaskReportOverviewRequest) (response *DescribeTaskReportOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeTaskReportOverviewRequest()
	}
	response = NewDescribeTaskReportOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubnetAssetsRequest() (request *DescribeSubnetAssetsRequest) {
	request = &DescribeSubnetAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSubnetAssets")
	return
}

func NewDescribeSubnetAssetsResponse() (response *DescribeSubnetAssetsResponse) {
	response = &DescribeSubnetAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子网列表
func (c *Client) DescribeSubnetAssets(request *DescribeSubnetAssetsRequest) (response *DescribeSubnetAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeSubnetAssetsRequest()
	}
	response = NewDescribeSubnetAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClbListenersRequest() (request *DescribeClbListenersRequest) {
	request = &DescribeClbListenersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeClbListeners")
	return
}

func NewDescribeClbListenersResponse() (response *DescribeClbListenersResponse) {
	response = &DescribeClbListenersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeListeners clb接口 获取clb监听器列表
func (c *Client) DescribeClbListeners(request *DescribeClbListenersRequest) (response *DescribeClbListenersResponse, err error) {
	if request == nil {
		request = NewDescribeClbListenersRequest()
	}
	response = NewDescribeClbListenersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProtectTimeRequest() (request *DescribeProtectTimeRequest) {
	request = &DescribeProtectTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeProtectTime")
	return
}

func NewDescribeProtectTimeResponse() (response *DescribeProtectTimeResponse) {
	response = &DescribeProtectTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防护时长和到期时间
func (c *Client) DescribeProtectTime(request *DescribeProtectTimeRequest) (response *DescribeProtectTimeResponse, err error) {
	if request == nil {
		request = NewDescribeProtectTimeRequest()
	}
	response = NewDescribeProtectTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSolidAssetStaticRequest() (request *DescribeSolidAssetStaticRequest) {
	request = &DescribeSolidAssetStaticRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSolidAssetStatic")
	return
}

func NewDescribeSolidAssetStaticResponse() (response *DescribeSolidAssetStaticResponse) {
	response = &DescribeSolidAssetStaticResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 立体防护资产接口
func (c *Client) DescribeSolidAssetStatic(request *DescribeSolidAssetStaticRequest) (response *DescribeSolidAssetStaticResponse, err error) {
	if request == nil {
		request = NewDescribeSolidAssetStaticRequest()
	}
	response = NewDescribeSolidAssetStaticResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCFWAssetStatisticsRequest() (request *DescribeCFWAssetStatisticsRequest) {
	request = &DescribeCFWAssetStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCFWAssetStatistics")
	return
}

func NewDescribeCFWAssetStatisticsResponse() (response *DescribeCFWAssetStatisticsResponse) {
	response = &DescribeCFWAssetStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云防资产中心统计数据
func (c *Client) DescribeCFWAssetStatistics(request *DescribeCFWAssetStatisticsRequest) (response *DescribeCFWAssetStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeCFWAssetStatisticsRequest()
	}
	response = NewDescribeCFWAssetStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewSyncUserAssetInfoRequest() (request *SyncUserAssetInfoRequest) {
	request = &SyncUserAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "SyncUserAssetInfo")
	return
}

func NewSyncUserAssetInfoResponse() (response *SyncUserAssetInfoResponse) {
	response = &SyncUserAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步单个用户信息(同步tag信息）
func (c *Client) SyncUserAssetInfo(request *SyncUserAssetInfoRequest) (response *SyncUserAssetInfoResponse, err error) {
	if request == nil {
		request = NewSyncUserAssetInfoRequest()
	}
	response = NewSyncUserAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagCountsRequest() (request *DescribeTagCountsRequest) {
	request = &DescribeTagCountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeTagCounts")
	return
}

func NewDescribeTagCountsResponse() (response *DescribeTagCountsResponse) {
	response = &DescribeTagCountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询不同产品日志数量接口
func (c *Client) DescribeTagCounts(request *DescribeTagCountsRequest) (response *DescribeTagCountsResponse, err error) {
	if request == nil {
		request = NewDescribeTagCountsRequest()
	}
	response = NewDescribeTagCountsResponse()
	err = c.Send(request, response)
	return
}

func NewGetBASTaskTechRequest() (request *GetBASTaskTechRequest) {
	request = &GetBASTaskTechRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "GetBASTaskTech")
	return
}

func NewGetBASTaskTechResponse() (response *GetBASTaskTechResponse) {
	response = &GetBASTaskTechResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取剧本的战术战技列表
func (c *Client) GetBASTaskTech(request *GetBASTaskTechRequest) (response *GetBASTaskTechResponse, err error) {
	if request == nil {
		request = NewGetBASTaskTechRequest()
	}
	response = NewGetBASTaskTechResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHighLevelRiskTopFiveListRequest() (request *DescribeHighLevelRiskTopFiveListRequest) {
	request = &DescribeHighLevelRiskTopFiveListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeHighLevelRiskTopFiveList")
	return
}

func NewDescribeHighLevelRiskTopFiveListResponse() (response *DescribeHighLevelRiskTopFiveListResponse) {
	response = &DescribeHighLevelRiskTopFiveListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取高危风险TOP5
func (c *Client) DescribeHighLevelRiskTopFiveList(request *DescribeHighLevelRiskTopFiveListRequest) (response *DescribeHighLevelRiskTopFiveListResponse, err error) {
	if request == nil {
		request = NewDescribeHighLevelRiskTopFiveListRequest()
	}
	response = NewDescribeHighLevelRiskTopFiveListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterWebsiteRiskListRequest() (request *DescribeRiskCenterWebsiteRiskListRequest) {
	request = &DescribeRiskCenterWebsiteRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterWebsiteRiskList")
	return
}

func NewDescribeRiskCenterWebsiteRiskListResponse() (response *DescribeRiskCenterWebsiteRiskListResponse) {
	response = &DescribeRiskCenterWebsiteRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网站风险列表
func (c *Client) DescribeRiskCenterWebsiteRiskList(request *DescribeRiskCenterWebsiteRiskListRequest) (response *DescribeRiskCenterWebsiteRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterWebsiteRiskListRequest()
	}
	response = NewDescribeRiskCenterWebsiteRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDomainAndIpRequest() (request *DeleteDomainAndIpRequest) {
	request = &DeleteDomainAndIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DeleteDomainAndIp")
	return
}

func NewDeleteDomainAndIpResponse() (response *DeleteDomainAndIpResponse) {
	response = &DeleteDomainAndIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除域名和ip请求
func (c *Client) DeleteDomainAndIp(request *DeleteDomainAndIpRequest) (response *DeleteDomainAndIpResponse, err error) {
	if request == nil {
		request = NewDeleteDomainAndIpRequest()
	}
	response = NewDeleteDomainAndIpResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanTimeRequest() (request *DescribeScanTimeRequest) {
	request = &DescribeScanTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeScanTime")
	return
}

func NewDescribeScanTimeResponse() (response *DescribeScanTimeResponse) {
	response = &DescribeScanTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计算任务预计扫描时间
func (c *Client) DescribeScanTime(request *DescribeScanTimeRequest) (response *DescribeScanTimeResponse, err error) {
	if request == nil {
		request = NewDescribeScanTimeRequest()
	}
	response = NewDescribeScanTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmTrendRequest() (request *DescribeAlarmTrendRequest) {
	request = &DescribeAlarmTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAlarmTrend")
	return
}

func NewDescribeAlarmTrendResponse() (response *DescribeAlarmTrendResponse) {
	response = &DescribeAlarmTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警概览
func (c *Client) DescribeAlarmTrend(request *DescribeAlarmTrendRequest) (response *DescribeAlarmTrendResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmTrendRequest()
	}
	response = NewDescribeAlarmTrendResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerAssetInfoRequest() (request *DescribeContainerAssetInfoRequest) {
	request = &DescribeContainerAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeContainerAssetInfo")
	return
}

func NewDescribeContainerAssetInfoResponse() (response *DescribeContainerAssetInfoResponse) {
	response = &DescribeContainerAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器详情信息
func (c *Client) DescribeContainerAssetInfo(request *DescribeContainerAssetInfoRequest) (response *DescribeContainerAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeContainerAssetInfoRequest()
	}
	response = NewDescribeContainerAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterAssetViewVULRiskListRequest() (request *DescribeRiskCenterAssetViewVULRiskListRequest) {
	request = &DescribeRiskCenterAssetViewVULRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewVULRiskList")
	return
}

func NewDescribeRiskCenterAssetViewVULRiskListResponse() (response *DescribeRiskCenterAssetViewVULRiskListResponse) {
	response = &DescribeRiskCenterAssetViewVULRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产视角的漏洞风险列表
func (c *Client) DescribeRiskCenterAssetViewVULRiskList(request *DescribeRiskCenterAssetViewVULRiskListRequest) (response *DescribeRiskCenterAssetViewVULRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterAssetViewVULRiskListRequest()
	}
	response = NewDescribeRiskCenterAssetViewVULRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewSyncAssetsProcessRequest() (request *SyncAssetsProcessRequest) {
	request = &SyncAssetsProcessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "SyncAssetsProcess")
	return
}

func NewSyncAssetsProcessResponse() (response *SyncAssetsProcessResponse) {
	response = &SyncAssetsProcessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看资产同步的进度
func (c *Client) SyncAssetsProcess(request *SyncAssetsProcessRequest) (response *SyncAssetsProcessResponse, err error) {
	if request == nil {
		request = NewSyncAssetsProcessRequest()
	}
	response = NewSyncAssetsProcessResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetsStructureRequest() (request *DescribeAssetsStructureRequest) {
	request = &DescribeAssetsStructureRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetsStructure")
	return
}

func NewDescribeAssetsStructureResponse() (response *DescribeAssetsStructureResponse) {
	response = &DescribeAssetsStructureResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产网络结构
func (c *Client) DescribeAssetsStructure(request *DescribeAssetsStructureRequest) (response *DescribeAssetsStructureResponse, err error) {
	if request == nil {
		request = NewDescribeAssetsStructureRequest()
	}
	response = NewDescribeAssetsStructureResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRiskCenterRiskStatusRequest() (request *ModifyRiskCenterRiskStatusRequest) {
	request = &ModifyRiskCenterRiskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyRiskCenterRiskStatus")
	return
}

func NewModifyRiskCenterRiskStatusResponse() (response *ModifyRiskCenterRiskStatusResponse) {
	response = &ModifyRiskCenterRiskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改风险中心风险状态
func (c *Client) ModifyRiskCenterRiskStatus(request *ModifyRiskCenterRiskStatusRequest) (response *ModifyRiskCenterRiskStatusResponse, err error) {
	if request == nil {
		request = NewModifyRiskCenterRiskStatusRequest()
	}
	response = NewModifyRiskCenterRiskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAssetConfigRequest() (request *DeleteAssetConfigRequest) {
	request = &DeleteAssetConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DeleteAssetConfig")
	return
}

func NewDeleteAssetConfigResponse() (response *DeleteAssetConfigResponse) {
	response = &DeleteAssetConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清楚资产配置
func (c *Client) DeleteAssetConfig(request *DeleteAssetConfigRequest) (response *DeleteAssetConfigResponse, err error) {
	if request == nil {
		request = NewDeleteAssetConfigRequest()
	}
	response = NewDeleteAssetConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndeterminacyAssetInfoRequest() (request *DescribeIndeterminacyAssetInfoRequest) {
	request = &DescribeIndeterminacyAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeIndeterminacyAssetInfo")
	return
}

func NewDescribeIndeterminacyAssetInfoResponse() (response *DescribeIndeterminacyAssetInfoResponse) {
	response = &DescribeIndeterminacyAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通用资产详情
func (c *Client) DescribeIndeterminacyAssetInfo(request *DescribeIndeterminacyAssetInfoRequest) (response *DescribeIndeterminacyAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeIndeterminacyAssetInfoRequest()
	}
	response = NewDescribeIndeterminacyAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafCiphersDetailRequest() (request *DescribeWafCiphersDetailRequest) {
	request = &DescribeWafCiphersDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWafCiphersDetail")
	return
}

func NewDescribeWafCiphersDetailResponse() (response *DescribeWafCiphersDetailResponse) {
	response = &DescribeWafCiphersDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeCiphersDetail waf接口 获取加密套件信息
func (c *Client) DescribeWafCiphersDetail(request *DescribeWafCiphersDetailRequest) (response *DescribeWafCiphersDetailResponse, err error) {
	if request == nil {
		request = NewDescribeWafCiphersDetailRequest()
	}
	response = NewDescribeWafCiphersDetailResponse()
	err = c.Send(request, response)
	return
}

func NewSyncScopeAssetsRequest() (request *SyncScopeAssetsRequest) {
	request = &SyncScopeAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "SyncScopeAssets")
	return
}

func NewSyncScopeAssetsResponse() (response *SyncScopeAssetsResponse) {
	response = &SyncScopeAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部分资产同步
func (c *Client) SyncScopeAssets(request *SyncScopeAssetsRequest) (response *SyncScopeAssetsResponse, err error) {
	if request == nil {
		request = NewSyncScopeAssetsRequest()
	}
	response = NewSyncScopeAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLocalImageAssetsRequest() (request *DescribeLocalImageAssetsRequest) {
	request = &DescribeLocalImageAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeLocalImageAssets")
	return
}

func NewDescribeLocalImageAssetsResponse() (response *DescribeLocalImageAssetsResponse) {
	response = &DescribeLocalImageAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本地镜像列表
func (c *Client) DescribeLocalImageAssets(request *DescribeLocalImageAssetsRequest) (response *DescribeLocalImageAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeLocalImageAssetsRequest()
	}
	response = NewDescribeLocalImageAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainAssetInfoRequest() (request *DescribeDomainAssetInfoRequest) {
	request = &DescribeDomainAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeDomainAssetInfo")
	return
}

func NewDescribeDomainAssetInfoResponse() (response *DescribeDomainAssetInfoResponse) {
	response = &DescribeDomainAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 详情
func (c *Client) DescribeDomainAssetInfo(request *DescribeDomainAssetInfoRequest) (response *DescribeDomainAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDomainAssetInfoRequest()
	}
	response = NewDescribeDomainAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicIpAssetsRequest() (request *DescribePublicIpAssetsRequest) {
	request = &DescribePublicIpAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribePublicIpAssets")
	return
}

func NewDescribePublicIpAssetsResponse() (response *DescribePublicIpAssetsResponse) {
	response = &DescribePublicIpAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ip公网列表
func (c *Client) DescribePublicIpAssets(request *DescribePublicIpAssetsRequest) (response *DescribePublicIpAssetsResponse, err error) {
	if request == nil {
		request = NewDescribePublicIpAssetsRequest()
	}
	response = NewDescribePublicIpAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCWPTrialRequest() (request *CreateCWPTrialRequest) {
	request = &CreateCWPTrialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateCWPTrial")
	return
}

func NewCreateCWPTrialResponse() (response *CreateCWPTrialResponse) {
	response = &CreateCWPTrialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建主机试用
func (c *Client) CreateCWPTrial(request *CreateCWPTrialRequest) (response *CreateCWPTrialResponse, err error) {
	if request == nil {
		request = NewCreateCWPTrialRequest()
	}
	response = NewCreateCWPTrialResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBASScriptRequest() (request *DescribeBASScriptRequest) {
	request = &DescribeBASScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeBASScript")
	return
}

func NewDescribeBASScriptResponse() (response *DescribeBASScriptResponse) {
	response = &DescribeBASScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BAS剧本列表
func (c *Client) DescribeBASScript(request *DescribeBASScriptRequest) (response *DescribeBASScriptResponse, err error) {
	if request == nil {
		request = NewDescribeBASScriptRequest()
	}
	response = NewDescribeBASScriptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportTaskRequest() (request *DescribeExportTaskRequest) {
	request = &DescribeExportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeExportTask")
	return
}

func NewDescribeExportTaskResponse() (response *DescribeExportTaskResponse) {
	response = &DescribeExportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日审导出任务信息
func (c *Client) DescribeExportTask(request *DescribeExportTaskRequest) (response *DescribeExportTaskResponse, err error) {
	if request == nil {
		request = NewDescribeExportTaskRequest()
	}
	response = NewDescribeExportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCFWRuleOverviewRequest() (request *DescribeCFWRuleOverviewRequest) {
	request = &DescribeCFWRuleOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCFWRuleOverview")
	return
}

func NewDescribeCFWRuleOverviewResponse() (response *DescribeCFWRuleOverviewResponse) {
	response = &DescribeCFWRuleOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云防火墙规则列表概况
func (c *Client) DescribeCFWRuleOverview(request *DescribeCFWRuleOverviewRequest) (response *DescribeCFWRuleOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeCFWRuleOverviewRequest()
	}
	response = NewDescribeCFWRuleOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewStopRiskCenterTaskRequest() (request *StopRiskCenterTaskRequest) {
	request = &StopRiskCenterTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "StopRiskCenterTask")
	return
}

func NewStopRiskCenterTaskResponse() (response *StopRiskCenterTaskResponse) {
	response = &StopRiskCenterTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止扫风险中心扫描任务
func (c *Client) StopRiskCenterTask(request *StopRiskCenterTaskRequest) (response *StopRiskCenterTaskResponse, err error) {
	if request == nil {
		request = NewStopRiskCenterTaskRequest()
	}
	response = NewStopRiskCenterTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterPortViewPortRiskListRequest() (request *DescribeRiskCenterPortViewPortRiskListRequest) {
	request = &DescribeRiskCenterPortViewPortRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterPortViewPortRiskList")
	return
}

func NewDescribeRiskCenterPortViewPortRiskListResponse() (response *DescribeRiskCenterPortViewPortRiskListResponse) {
	response = &DescribeRiskCenterPortViewPortRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取端口视角的端口风险列表
func (c *Client) DescribeRiskCenterPortViewPortRiskList(request *DescribeRiskCenterPortViewPortRiskListRequest) (response *DescribeRiskCenterPortViewPortRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterPortViewPortRiskListRequest()
	}
	response = NewDescribeRiskCenterPortViewPortRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewApplyTrialRequest() (request *ApplyTrialRequest) {
	request = &ApplyTrialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ApplyTrial")
	return
}

func NewApplyTrialResponse() (response *ApplyTrialResponse) {
	response = &ApplyTrialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云安全中心试用
func (c *Client) ApplyTrial(request *ApplyTrialRequest) (response *ApplyTrialResponse, err error) {
	if request == nil {
		request = NewApplyTrialRequest()
	}
	response = NewApplyTrialResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExaminationPayRequest() (request *DescribeExaminationPayRequest) {
	request = &DescribeExaminationPayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeExaminationPay")
	return
}

func NewDescribeExaminationPayResponse() (response *DescribeExaminationPayResponse) {
	response = &DescribeExaminationPayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏扫各个付费信息
func (c *Client) DescribeExaminationPay(request *DescribeExaminationPayRequest) (response *DescribeExaminationPayResponse, err error) {
	if request == nil {
		request = NewDescribeExaminationPayRequest()
	}
	response = NewDescribeExaminationPayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserPayInfoRequest() (request *DescribeUserPayInfoRequest) {
	request = &DescribeUserPayInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeUserPayInfo")
	return
}

func NewDescribeUserPayInfoResponse() (response *DescribeUserPayInfoResponse) {
	response = &DescribeUserPayInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户购买情况详情
func (c *Client) DescribeUserPayInfo(request *DescribeUserPayInfoRequest) (response *DescribeUserPayInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserPayInfoRequest()
	}
	response = NewDescribeUserPayInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCheckUpResultRequest() (request *ModifyCheckUpResultRequest) {
	request = &ModifyCheckUpResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyCheckUpResult")
	return
}

func NewModifyCheckUpResultResponse() (response *ModifyCheckUpResultResponse) {
	response = &ModifyCheckUpResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新体检得分
func (c *Client) ModifyCheckUpResult(request *ModifyCheckUpResultRequest) (response *ModifyCheckUpResultResponse, err error) {
	if request == nil {
		request = NewModifyCheckUpResultRequest()
	}
	response = NewModifyCheckUpResultResponse()
	err = c.Send(request, response)
	return
}

func NewSaveAssetsConfigRequest() (request *SaveAssetsConfigRequest) {
	request = &SaveAssetsConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "SaveAssetsConfig")
	return
}

func NewSaveAssetsConfigResponse() (response *SaveAssetsConfigResponse) {
	response = &SaveAssetsConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 插入或者更新资产配置（若资产已存在，则自动更新）
func (c *Client) SaveAssetsConfig(request *SaveAssetsConfigRequest) (response *SaveAssetsConfigResponse, err error) {
	if request == nil {
		request = NewSaveAssetsConfigRequest()
	}
	response = NewSaveAssetsConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetBASSummaryDataRequest() (request *GetBASSummaryDataRequest) {
	request = &GetBASSummaryDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "GetBASSummaryData")
	return
}

func NewGetBASSummaryDataResponse() (response *GetBASSummaryDataResponse) {
	response = &GetBASSummaryDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bas概况数据
func (c *Client) GetBASSummaryData(request *GetBASSummaryDataRequest) (response *GetBASSummaryDataResponse, err error) {
	if request == nil {
		request = NewGetBASSummaryDataRequest()
	}
	response = NewGetBASSummaryDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePayInfoRequest() (request *DescribePayInfoRequest) {
	request = &DescribePayInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribePayInfo")
	return
}

func NewDescribePayInfoResponse() (response *DescribePayInfoResponse) {
	response = &DescribePayInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云安全中心支付信息
func (c *Client) DescribePayInfo(request *DescribePayInfoRequest) (response *DescribePayInfoResponse, err error) {
	if request == nil {
		request = NewDescribePayInfoRequest()
	}
	response = NewDescribePayInfoResponse()
	err = c.Send(request, response)
	return
}

func NewAlarmDiversityRequest() (request *AlarmDiversityRequest) {
	request = &AlarmDiversityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "AlarmDiversity")
	return
}

func NewAlarmDiversityResponse() (response *AlarmDiversityResponse) {
	response = &AlarmDiversityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警分类
func (c *Client) AlarmDiversity(request *AlarmDiversityRequest) (response *AlarmDiversityResponse, err error) {
	if request == nil {
		request = NewAlarmDiversityRequest()
	}
	response = NewAlarmDiversityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserOperationPermissionRequest() (request *DescribeUserOperationPermissionRequest) {
	request = &DescribeUserOperationPermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeUserOperationPermission")
	return
}

func NewDescribeUserOperationPermissionResponse() (response *DescribeUserOperationPermissionResponse) {
	response = &DescribeUserOperationPermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 限制用户的操作权限
func (c *Client) DescribeUserOperationPermission(request *DescribeUserOperationPermissionRequest) (response *DescribeUserOperationPermissionResponse, err error) {
	if request == nil {
		request = NewDescribeUserOperationPermissionRequest()
	}
	response = NewDescribeUserOperationPermissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCVEInfoRequest() (request *DescribeCVEInfoRequest) {
	request = &DescribeCVEInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCVEInfo")
	return
}

func NewDescribeCVEInfoResponse() (response *DescribeCVEInfoResponse) {
	response = &DescribeCVEInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用来判断某个CVE漏洞是否支持虚拟补丁
func (c *Client) DescribeCVEInfo(request *DescribeCVEInfoRequest) (response *DescribeCVEInfoResponse, err error) {
	if request == nil {
		request = NewDescribeCVEInfoRequest()
	}
	response = NewDescribeCVEInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterWeakPasswordViewWeakPasswordRiskListRequest() (request *DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListRequest) {
	request = &DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterWeakPasswordViewWeakPasswordRiskList")
	return
}

func NewDescribeRiskCenterWeakPasswordViewWeakPasswordRiskListResponse() (response *DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListResponse) {
	response = &DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弱口令视角的弱口令风险列表
func (c *Client) DescribeRiskCenterWeakPasswordViewWeakPasswordRiskList(request *DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListRequest) (response *DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterWeakPasswordViewWeakPasswordRiskListRequest()
	}
	response = NewDescribeRiskCenterWeakPasswordViewWeakPasswordRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBASATTCKMatrixRequest() (request *DescribeBASATTCKMatrixRequest) {
	request = &DescribeBASATTCKMatrixRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeBASATTCKMatrix")
	return
}

func NewDescribeBASATTCKMatrixResponse() (response *DescribeBASATTCKMatrixResponse) {
	response = &DescribeBASATTCKMatrixResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示战术战技
func (c *Client) DescribeBASATTCKMatrix(request *DescribeBASATTCKMatrixRequest) (response *DescribeBASATTCKMatrixResponse, err error) {
	if request == nil {
		request = NewDescribeBASATTCKMatrixRequest()
	}
	response = NewDescribeBASATTCKMatrixResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDbAssetsRequest() (request *DescribeDbAssetsRequest) {
	request = &DescribeDbAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeDbAssets")
	return
}

func NewDescribeDbAssetsResponse() (response *DescribeDbAssetsResponse) {
	response = &DescribeDbAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产列表
func (c *Client) DescribeDbAssets(request *DescribeDbAssetsRequest) (response *DescribeDbAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeDbAssetsRequest()
	}
	response = NewDescribeDbAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExaminationScanRequest() (request *DescribeExaminationScanRequest) {
	request = &DescribeExaminationScanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeExaminationScan")
	return
}

func NewDescribeExaminationScanResponse() (response *DescribeExaminationScanResponse) {
	response = &DescribeExaminationScanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 是否能支持扫描
func (c *Client) DescribeExaminationScan(request *DescribeExaminationScanRequest) (response *DescribeExaminationScanResponse, err error) {
	if request == nil {
		request = NewDescribeExaminationScanRequest()
	}
	response = NewDescribeExaminationScanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeListenerListRequest() (request *DescribeListenerListRequest) {
	request = &DescribeListenerListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeListenerList")
	return
}

func NewDescribeListenerListResponse() (response *DescribeListenerListResponse) {
	response = &DescribeListenerListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询clb监听器列表
func (c *Client) DescribeListenerList(request *DescribeListenerListRequest) (response *DescribeListenerListResponse, err error) {
	if request == nil {
		request = NewDescribeListenerListRequest()
	}
	response = NewDescribeListenerListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationUserInfoRequest() (request *DescribeOrganizationUserInfoRequest) {
	request = &DescribeOrganizationUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeOrganizationUserInfo")
	return
}

func NewDescribeOrganizationUserInfoResponse() (response *DescribeOrganizationUserInfoResponse) {
	response = &DescribeOrganizationUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集团账号用户列表
func (c *Client) DescribeOrganizationUserInfo(request *DescribeOrganizationUserInfoRequest) (response *DescribeOrganizationUserInfoResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationUserInfoRequest()
	}
	response = NewDescribeOrganizationUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterVULViewVULRiskListRequest() (request *DescribeRiskCenterVULViewVULRiskListRequest) {
	request = &DescribeRiskCenterVULViewVULRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterVULViewVULRiskList")
	return
}

func NewDescribeRiskCenterVULViewVULRiskListResponse() (response *DescribeRiskCenterVULViewVULRiskListResponse) {
	response = &DescribeRiskCenterVULViewVULRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞视角的漏洞风险列表
func (c *Client) DescribeRiskCenterVULViewVULRiskList(request *DescribeRiskCenterVULViewVULRiskListRequest) (response *DescribeRiskCenterVULViewVULRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterVULViewVULRiskListRequest()
	}
	response = NewDescribeRiskCenterVULViewVULRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFreeScanChanceRequest() (request *CreateFreeScanChanceRequest) {
	request = &CreateFreeScanChanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateFreeScanChance")
	return
}

func NewCreateFreeScanChanceResponse() (response *CreateFreeScanChanceResponse) {
	response = &CreateFreeScanChanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 抢购免费扫描机会
func (c *Client) CreateFreeScanChance(request *CreateFreeScanChanceRequest) (response *CreateFreeScanChanceResponse, err error) {
	if request == nil {
		request = NewCreateFreeScanChanceRequest()
	}
	response = NewCreateFreeScanChanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchBugInfoRequest() (request *DescribeSearchBugInfoRequest) {
	request = &DescribeSearchBugInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSearchBugInfo")
	return
}

func NewDescribeSearchBugInfoResponse() (response *DescribeSearchBugInfoResponse) {
	response = &DescribeSearchBugInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 立体防护中心查询漏洞信息
func (c *Client) DescribeSearchBugInfo(request *DescribeSearchBugInfoRequest) (response *DescribeSearchBugInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSearchBugInfoRequest()
	}
	response = NewDescribeSearchBugInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLicenseBindsRequest() (request *ModifyLicenseBindsRequest) {
	request = &ModifyLicenseBindsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyLicenseBinds")
	return
}

func NewModifyLicenseBindsResponse() (response *ModifyLicenseBindsResponse) {
	response = &ModifyLicenseBindsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定订单信息
func (c *Client) ModifyLicenseBinds(request *ModifyLicenseBindsRequest) (response *ModifyLicenseBindsResponse, err error) {
	if request == nil {
		request = NewModifyLicenseBindsRequest()
	}
	response = NewModifyLicenseBindsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDbAssetInfoRequest() (request *DescribeDbAssetInfoRequest) {
	request = &DescribeDbAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeDbAssetInfo")
	return
}

func NewDescribeDbAssetInfoResponse() (response *DescribeDbAssetInfoResponse) {
	response = &DescribeDbAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// db资产详情
func (c *Client) DescribeDbAssetInfo(request *DescribeDbAssetInfoRequest) (response *DescribeDbAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDbAssetInfoRequest()
	}
	response = NewDescribeDbAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebServiceRequest() (request *ModifyWebServiceRequest) {
	request = &ModifyWebServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyWebService")
	return
}

func NewModifyWebServiceResponse() (response *ModifyWebServiceResponse) {
	response = &ModifyWebServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑web服务
func (c *Client) ModifyWebService(request *ModifyWebServiceRequest) (response *ModifyWebServiceResponse, err error) {
	if request == nil {
		request = NewModifyWebServiceRequest()
	}
	response = NewModifyWebServiceResponse()
	err = c.Send(request, response)
	return
}

func NewGetBASScriptDetailRequest() (request *GetBASScriptDetailRequest) {
	request = &GetBASScriptDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "GetBASScriptDetail")
	return
}

func NewGetBASScriptDetailResponse() (response *GetBASScriptDetailResponse) {
	response = &GetBASScriptDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个剧本详情
func (c *Client) GetBASScriptDetail(request *GetBASScriptDetailRequest) (response *GetBASScriptDetailResponse, err error) {
	if request == nil {
		request = NewGetBASScriptDetailRequest()
	}
	response = NewGetBASScriptDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkStructureRequest() (request *DescribeNetworkStructureRequest) {
	request = &DescribeNetworkStructureRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeNetworkStructure")
	return
}

func NewDescribeNetworkStructureResponse() (response *DescribeNetworkStructureResponse) {
	response = &DescribeNetworkStructureResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网络结构-树
func (c *Client) DescribeNetworkStructure(request *DescribeNetworkStructureRequest) (response *DescribeNetworkStructureResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkStructureRequest()
	}
	response = NewDescribeNetworkStructureResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSSLCertificatesRequest() (request *DescribeSSLCertificatesRequest) {
	request = &DescribeSSLCertificatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSSLCertificates")
	return
}

func NewDescribeSSLCertificatesResponse() (response *DescribeSSLCertificatesResponse) {
	response = &DescribeSSLCertificatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeCertificates  ssl接口 获 取证书信息
func (c *Client) DescribeSSLCertificates(request *DescribeSSLCertificatesRequest) (response *DescribeSSLCertificatesResponse, err error) {
	if request == nil {
		request = NewDescribeSSLCertificatesRequest()
	}
	response = NewDescribeSSLCertificatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBASATTCKListRequest() (request *DescribeBASATTCKListRequest) {
	request = &DescribeBASATTCKListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeBASATTCKList")
	return
}

func NewDescribeBASATTCKListResponse() (response *DescribeBASATTCKListResponse) {
	response = &DescribeBASATTCKListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取att&ck战术战技列表
func (c *Client) DescribeBASATTCKList(request *DescribeBASATTCKListRequest) (response *DescribeBASATTCKListResponse, err error) {
	if request == nil {
		request = NewDescribeBASATTCKListRequest()
	}
	response = NewDescribeBASATTCKListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterAssetViewWeakPasswordRiskListRequest() (request *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) {
	request = &DescribeRiskCenterAssetViewWeakPasswordRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewWeakPasswordRiskList")
	return
}

func NewDescribeRiskCenterAssetViewWeakPasswordRiskListResponse() (response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) {
	response = &DescribeRiskCenterAssetViewWeakPasswordRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产视角的弱口令风险列表
func (c *Client) DescribeRiskCenterAssetViewWeakPasswordRiskList(request *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) (response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterAssetViewWeakPasswordRiskListRequest()
	}
	response = NewDescribeRiskCenterAssetViewWeakPasswordRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetLocationListRequest() (request *DescribeAssetLocationListRequest) {
	request = &DescribeAssetLocationListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetLocationList")
	return
}

func NewDescribeAssetLocationListResponse() (response *DescribeAssetLocationListResponse) {
	response = &DescribeAssetLocationListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产分布
func (c *Client) DescribeAssetLocationList(request *DescribeAssetLocationListRequest) (response *DescribeAssetLocationListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetLocationListRequest()
	}
	response = NewDescribeAssetLocationListResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePocScanMissionRequest() (request *CreatePocScanMissionRequest) {
	request = &CreatePocScanMissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreatePocScanMission")
	return
}

func NewCreatePocScanMissionResponse() (response *CreatePocScanMissionResponse) {
	response = &CreatePocScanMissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于立体防护下发漏洞扫描逻辑
func (c *Client) CreatePocScanMission(request *CreatePocScanMissionRequest) (response *CreatePocScanMissionResponse, err error) {
	if request == nil {
		request = NewCreatePocScanMissionRequest()
	}
	response = NewCreatePocScanMissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseListRequest() (request *DescribeLicenseListRequest) {
	request = &DescribeLicenseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeLicenseList")
	return
}

func NewDescribeLicenseListResponse() (response *DescribeLicenseListResponse) {
	response = &DescribeLicenseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取授权订单列表
func (c *Client) DescribeLicenseList(request *DescribeLicenseListRequest) (response *DescribeLicenseListResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseListRequest()
	}
	response = NewDescribeLicenseListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLocalImageComponentRequest() (request *DescribeLocalImageComponentRequest) {
	request = &DescribeLocalImageComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeLocalImageComponent")
	return
}

func NewDescribeLocalImageComponentResponse() (response *DescribeLocalImageComponentResponse) {
	response = &DescribeLocalImageComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本地组件
func (c *Client) DescribeLocalImageComponent(request *DescribeLocalImageComponentRequest) (response *DescribeLocalImageComponentResponse, err error) {
	if request == nil {
		request = NewDescribeLocalImageComponentRequest()
	}
	response = NewDescribeLocalImageComponentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperateLogListRequest() (request *DescribeOperateLogListRequest) {
	request = &DescribeOperateLogListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeOperateLogList")
	return
}

func NewDescribeOperateLogListResponse() (response *DescribeOperateLogListResponse) {
	response = &DescribeOperateLogListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取操作记录列表
func (c *Client) DescribeOperateLogList(request *DescribeOperateLogListRequest) (response *DescribeOperateLogListResponse, err error) {
	if request == nil {
		request = NewDescribeOperateLogListRequest()
	}
	response = NewDescribeOperateLogListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterRiskTrendListRequest() (request *DescribeRiskCenterRiskTrendListRequest) {
	request = &DescribeRiskCenterRiskTrendListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterRiskTrendList")
	return
}

func NewDescribeRiskCenterRiskTrendListResponse() (response *DescribeRiskCenterRiskTrendListResponse) {
	response = &DescribeRiskCenterRiskTrendListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取风险趋势
func (c *Client) DescribeRiskCenterRiskTrendList(request *DescribeRiskCenterRiskTrendListRequest) (response *DescribeRiskCenterRiskTrendListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterRiskTrendListRequest()
	}
	response = NewDescribeRiskCenterRiskTrendListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAuditTrendRequest() (request *DescribeUserAuditTrendRequest) {
	request = &DescribeUserAuditTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeUserAuditTrend")
	return
}

func NewDescribeUserAuditTrendResponse() (response *DescribeUserAuditTrendResponse) {
	response = &DescribeUserAuditTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户日志趋势
func (c *Client) DescribeUserAuditTrend(request *DescribeUserAuditTrendRequest) (response *DescribeUserAuditTrendResponse, err error) {
	if request == nil {
		request = NewDescribeUserAuditTrendRequest()
	}
	response = NewDescribeUserAuditTrendResponse()
	err = c.Send(request, response)
	return
}

func NewSyncAssetsRequest() (request *SyncAssetsRequest) {
	request = &SyncAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "SyncAssets")
	return
}

func NewSyncAssetsResponse() (response *SyncAssetsResponse) {
	response = &SyncAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产同步按钮
func (c *Client) SyncAssets(request *SyncAssetsRequest) (response *SyncAssetsResponse, err error) {
	if request == nil {
		request = NewSyncAssetsRequest()
	}
	response = NewSyncAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskCostQuotaRequest() (request *DescribeTaskCostQuotaRequest) {
	request = &DescribeTaskCostQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskCostQuota")
	return
}

func NewDescribeTaskCostQuotaResponse() (response *DescribeTaskCostQuotaResponse) {
	response = &DescribeTaskCostQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务消耗配额
func (c *Client) DescribeTaskCostQuota(request *DescribeTaskCostQuotaRequest) (response *DescribeTaskCostQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeTaskCostQuotaRequest()
	}
	response = NewDescribeTaskCostQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserStorageRequest() (request *DescribeUserStorageRequest) {
	request = &DescribeUserStorageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeUserStorage")
	return
}

func NewDescribeUserStorageResponse() (response *DescribeUserStorageResponse) {
	response = &DescribeUserStorageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户的日志存储量信息
func (c *Client) DescribeUserStorage(request *DescribeUserStorageRequest) (response *DescribeUserStorageResponse, err error) {
	if request == nil {
		request = NewDescribeUserStorageRequest()
	}
	response = NewDescribeUserStorageResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNeedToSolveListChangeRequest() (request *UpdateNeedToSolveListChangeRequest) {
	request = &UpdateNeedToSolveListChangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "UpdateNeedToSolveListChange")
	return
}

func NewUpdateNeedToSolveListChangeResponse() (response *UpdateNeedToSolveListChangeResponse) {
	response = &UpdateNeedToSolveListChangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新待办的处理信息
func (c *Client) UpdateNeedToSolveListChange(request *UpdateNeedToSolveListChangeRequest) (response *UpdateNeedToSolveListChangeResponse, err error) {
	if request == nil {
		request = NewUpdateNeedToSolveListChangeRequest()
	}
	response = NewUpdateNeedToSolveListChangeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAssetsRequest() (request *DescribeClusterAssetsRequest) {
	request = &DescribeClusterAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeClusterAssets")
	return
}

func NewDescribeClusterAssetsResponse() (response *DescribeClusterAssetsResponse) {
	response = &DescribeClusterAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群列表
func (c *Client) DescribeClusterAssets(request *DescribeClusterAssetsRequest) (response *DescribeClusterAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAssetsRequest()
	}
	response = NewDescribeClusterAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportTaskListRequest() (request *DescribeExportTaskListRequest) {
	request = &DescribeExportTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeExportTaskList")
	return
}

func NewDescribeExportTaskListResponse() (response *DescribeExportTaskListResponse) {
	response = &DescribeExportTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取异步导出文件的的任务列表
func (c *Client) DescribeExportTaskList(request *DescribeExportTaskListRequest) (response *DescribeExportTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeExportTaskListRequest()
	}
	response = NewDescribeExportTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOrganizationAccountStatusRequest() (request *ModifyOrganizationAccountStatusRequest) {
	request = &ModifyOrganizationAccountStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyOrganizationAccountStatus")
	return
}

func NewModifyOrganizationAccountStatusResponse() (response *ModifyOrganizationAccountStatusResponse) {
	response = &ModifyOrganizationAccountStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集团账号状态
func (c *Client) ModifyOrganizationAccountStatus(request *ModifyOrganizationAccountStatusRequest) (response *ModifyOrganizationAccountStatusResponse, err error) {
	if request == nil {
		request = NewModifyOrganizationAccountStatusRequest()
	}
	response = NewModifyOrganizationAccountStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWafClbHostRequest() (request *CreateWafClbHostRequest) {
	request = &CreateWafClbHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateWafClbHost")
	return
}

func NewCreateWafClbHostResponse() (response *CreateWafClbHostResponse) {
	response = &CreateWafClbHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateHost waf接口 clb开启防护
func (c *Client) CreateWafClbHost(request *CreateWafClbHostRequest) (response *CreateWafClbHostResponse, err error) {
	if request == nil {
		request = NewCreateWafClbHostRequest()
	}
	response = NewCreateWafClbHostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecOverviewRequest() (request *DescribeSecOverviewRequest) {
	request = &DescribeSecOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSecOverview")
	return
}

func NewDescribeSecOverviewResponse() (response *DescribeSecOverviewResponse) {
	response = &DescribeSecOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全总览
func (c *Client) DescribeSecOverview(request *DescribeSecOverviewRequest) (response *DescribeSecOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeSecOverviewRequest()
	}
	response = NewDescribeSecOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWeakPWDRiskAdvanceCFGListRequest() (request *DescribeWeakPWDRiskAdvanceCFGListRequest) {
	request = &DescribeWeakPWDRiskAdvanceCFGListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWeakPWDRiskAdvanceCFGList")
	return
}

func NewDescribeWeakPWDRiskAdvanceCFGListResponse() (response *DescribeWeakPWDRiskAdvanceCFGListResponse) {
	response = &DescribeWeakPWDRiskAdvanceCFGListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用来查询展示可用的弱口令检测项
func (c *Client) DescribeWeakPWDRiskAdvanceCFGList(request *DescribeWeakPWDRiskAdvanceCFGListRequest) (response *DescribeWeakPWDRiskAdvanceCFGListResponse, err error) {
	if request == nil {
		request = NewDescribeWeakPWDRiskAdvanceCFGListRequest()
	}
	response = NewDescribeWeakPWDRiskAdvanceCFGListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafProductStatusRequest() (request *DescribeWafProductStatusRequest) {
	request = &DescribeWafProductStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWafProductStatus")
	return
}

func NewDescribeWafProductStatusResponse() (response *DescribeWafProductStatusResponse) {
	response = &DescribeWafProductStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询waf版本及试用状态
func (c *Client) DescribeWafProductStatus(request *DescribeWafProductStatusRequest) (response *DescribeWafProductStatusResponse, err error) {
	if request == nil {
		request = NewDescribeWafProductStatusRequest()
	}
	response = NewDescribeWafProductStatusResponse()
	err = c.Send(request, response)
	return
}

func NewExportAuditLogRequest() (request *ExportAuditLogRequest) {
	request = &ExportAuditLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ExportAuditLog")
	return
}

func NewExportAuditLogResponse() (response *ExportAuditLogResponse) {
	response = &ExportAuditLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出审计日志
func (c *Client) ExportAuditLog(request *ExportAuditLogRequest) (response *ExportAuditLogResponse, err error) {
	if request == nil {
		request = NewExportAuditLogRequest()
	}
	response = NewExportAuditLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyExitExamRequest() (request *ModifyExitExamRequest) {
	request = &ModifyExitExamRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyExitExam")
	return
}

func NewModifyExitExamResponse() (response *ModifyExitExamResponse) {
	response = &ModifyExitExamResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用来退出安全中心体检
func (c *Client) ModifyExitExam(request *ModifyExitExamRequest) (response *ModifyExitExamResponse, err error) {
	if request == nil {
		request = NewModifyExitExamRequest()
	}
	response = NewModifyExitExamResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNeedToSolveListChangeRequest() (request *DescribeNeedToSolveListChangeRequest) {
	request = &DescribeNeedToSolveListChangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeNeedToSolveListChange")
	return
}

func NewDescribeNeedToSolveListChangeResponse() (response *DescribeNeedToSolveListChangeResponse) {
	response = &DescribeNeedToSolveListChangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询待办处理记录
func (c *Client) DescribeNeedToSolveListChange(request *DescribeNeedToSolveListChangeRequest) (response *DescribeNeedToSolveListChangeResponse, err error) {
	if request == nil {
		request = NewDescribeNeedToSolveListChangeRequest()
	}
	response = NewDescribeNeedToSolveListChangeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterAssetViewPortRiskListRequest() (request *DescribeRiskCenterAssetViewPortRiskListRequest) {
	request = &DescribeRiskCenterAssetViewPortRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewPortRiskList")
	return
}

func NewDescribeRiskCenterAssetViewPortRiskListResponse() (response *DescribeRiskCenterAssetViewPortRiskListResponse) {
	response = &DescribeRiskCenterAssetViewPortRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产视角的端口风险列表
func (c *Client) DescribeRiskCenterAssetViewPortRiskList(request *DescribeRiskCenterAssetViewPortRiskListRequest) (response *DescribeRiskCenterAssetViewPortRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterAssetViewPortRiskListRequest()
	}
	response = NewDescribeRiskCenterAssetViewPortRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskReportInfoRequest() (request *DescribeTaskReportInfoRequest) {
	request = &DescribeTaskReportInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskReportInfo")
	return
}

func NewDescribeTaskReportInfoResponse() (response *DescribeTaskReportInfoResponse) {
	response = &DescribeTaskReportInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取概览页任务和报告信息
func (c *Client) DescribeTaskReportInfo(request *DescribeTaskReportInfoRequest) (response *DescribeTaskReportInfoResponse, err error) {
	if request == nil {
		request = NewDescribeTaskReportInfoRequest()
	}
	response = NewDescribeTaskReportInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStorageRateRequest() (request *DescribeStorageRateRequest) {
	request = &DescribeStorageRateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeStorageRate")
	return
}

func NewDescribeStorageRateResponse() (response *DescribeStorageRateResponse) {
	response = &DescribeStorageRateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户日志存储比例
func (c *Client) DescribeStorageRate(request *DescribeStorageRateRequest) (response *DescribeStorageRateResponse, err error) {
	if request == nil {
		request = NewDescribeStorageRateRequest()
	}
	response = NewDescribeStorageRateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductStatusRequest() (request *CreateProductStatusRequest) {
	request = &CreateProductStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateProductStatus")
	return
}

func NewCreateProductStatusResponse() (response *CreateProductStatusResponse) {
	response = &CreateProductStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用来下发安全中心的体检任务
func (c *Client) CreateProductStatus(request *CreateProductStatusRequest) (response *CreateProductStatusResponse, err error) {
	if request == nil {
		request = NewCreateProductStatusRequest()
	}
	response = NewCreateProductStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRiskScanTaskRequest() (request *DeleteRiskScanTaskRequest) {
	request = &DeleteRiskScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DeleteRiskScanTask")
	return
}

func NewDeleteRiskScanTaskResponse() (response *DeleteRiskScanTaskResponse) {
	response = &DeleteRiskScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除风险中心扫描任务
func (c *Client) DeleteRiskScanTask(request *DeleteRiskScanTaskRequest) (response *DeleteRiskScanTaskResponse, err error) {
	if request == nil {
		request = NewDeleteRiskScanTaskRequest()
	}
	response = NewDeleteRiskScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetConfigListRequest() (request *DescribeAssetConfigListRequest) {
	request = &DescribeAssetConfigListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetConfigList")
	return
}

func NewDescribeAssetConfigListResponse() (response *DescribeAssetConfigListResponse) {
	response = &DescribeAssetConfigListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据资产值获取资产的配置
func (c *Client) DescribeAssetConfigList(request *DescribeAssetConfigListRequest) (response *DescribeAssetConfigListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetConfigListRequest()
	}
	response = NewDescribeAssetConfigListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFreeChanceRequest() (request *DescribeFreeChanceRequest) {
	request = &DescribeFreeChanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeFreeChance")
	return
}

func NewDescribeFreeChanceResponse() (response *DescribeFreeChanceResponse) {
	response = &DescribeFreeChanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询立体防护应急漏洞扫描剩余免费次数
func (c *Client) DescribeFreeChance(request *DescribeFreeChanceRequest) (response *DescribeFreeChanceResponse, err error) {
	if request == nil {
		request = NewDescribeFreeChanceRequest()
	}
	response = NewDescribeFreeChanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUpdateBugInfoRequest() (request *ModifyUpdateBugInfoRequest) {
	request = &ModifyUpdateBugInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyUpdateBugInfo")
	return
}

func NewModifyUpdateBugInfoResponse() (response *ModifyUpdateBugInfoResponse) {
	response = &ModifyUpdateBugInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过该接口更新资产表信息，实时展示漏洞信息
func (c *Client) ModifyUpdateBugInfo(request *ModifyUpdateBugInfoRequest) (response *ModifyUpdateBugInfoResponse, err error) {
	if request == nil {
		request = NewModifyUpdateBugInfoRequest()
	}
	response = NewModifyUpdateBugInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVULRiskAdvanceCFGListRequest() (request *DescribeVULRiskAdvanceCFGListRequest) {
	request = &DescribeVULRiskAdvanceCFGListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeVULRiskAdvanceCFGList")
	return
}

func NewDescribeVULRiskAdvanceCFGListResponse() (response *DescribeVULRiskAdvanceCFGListResponse) {
	response = &DescribeVULRiskAdvanceCFGListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞风险高级配置
func (c *Client) DescribeVULRiskAdvanceCFGList(request *DescribeVULRiskAdvanceCFGListRequest) (response *DescribeVULRiskAdvanceCFGListResponse, err error) {
	if request == nil {
		request = NewDescribeVULRiskAdvanceCFGListRequest()
	}
	response = NewDescribeVULRiskAdvanceCFGListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafVersionRequest() (request *DescribeWafVersionRequest) {
	request = &DescribeWafVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWafVersion")
	return
}

func NewDescribeWafVersionResponse() (response *DescribeWafVersionResponse) {
	response = &DescribeWafVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询waf的saas和clb的开通情况以及当前开通最高的版本
// 该接口测试的时候请带SubAccountUin参数，这个参数的值等于Uin字段，否则返回结果是空的
func (c *Client) DescribeWafVersion(request *DescribeWafVersionRequest) (response *DescribeWafVersionResponse, err error) {
	if request == nil {
		request = NewDescribeWafVersionRequest()
	}
	response = NewDescribeWafVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskLogListRequest() (request *DescribeTaskLogListRequest) {
	request = &DescribeTaskLogListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskLogList")
	return
}

func NewDescribeTaskLogListResponse() (response *DescribeTaskLogListResponse) {
	response = &DescribeTaskLogListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务扫描报告列表
func (c *Client) DescribeTaskLogList(request *DescribeTaskLogListRequest) (response *DescribeTaskLogListResponse, err error) {
	if request == nil {
		request = NewDescribeTaskLogListRequest()
	}
	response = NewDescribeTaskLogListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSendVoucherToUserTaskRequest() (request *CreateSendVoucherToUserTaskRequest) {
	request = &CreateSendVoucherToUserTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateSendVoucherToUserTask")
	return
}

func NewCreateSendVoucherToUserTaskResponse() (response *CreateSendVoucherToUserTaskResponse) {
	response = &CreateSendVoucherToUserTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建赠送优惠券任务
func (c *Client) CreateSendVoucherToUserTask(request *CreateSendVoucherToUserTaskRequest) (response *CreateSendVoucherToUserTaskResponse, err error) {
	if request == nil {
		request = NewCreateSendVoucherToUserTaskRequest()
	}
	response = NewCreateSendVoucherToUserTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogRequest() (request *DeleteLogRequest) {
	request = &DeleteLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DeleteLog")
	return
}

func NewDeleteLogResponse() (response *DeleteLogResponse) {
	response = &DeleteLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清除日志接口，支持批量
func (c *Client) DeleteLog(request *DeleteLogRequest) (response *DeleteLogResponse, err error) {
	if request == nil {
		request = NewDeleteLogRequest()
	}
	response = NewDeleteLogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDomainAndIpRequest() (request *CreateDomainAndIpRequest) {
	request = &CreateDomainAndIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateDomainAndIp")
	return
}

func NewCreateDomainAndIpResponse() (response *CreateDomainAndIpResponse) {
	response = &CreateDomainAndIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建域名、ip相关信息
func (c *Client) CreateDomainAndIp(request *CreateDomainAndIpRequest) (response *CreateDomainAndIpResponse, err error) {
	if request == nil {
		request = NewCreateDomainAndIpRequest()
	}
	response = NewCreateDomainAndIpResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSolidProtectionTrialsRequest() (request *CreateSolidProtectionTrialsRequest) {
	request = &CreateSolidProtectionTrialsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateSolidProtectionTrials")
	return
}

func NewCreateSolidProtectionTrialsResponse() (response *CreateSolidProtectionTrialsResponse) {
	response = &CreateSolidProtectionTrialsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 立体防护中心产品试用一键开通接口
func (c *Client) CreateSolidProtectionTrials(request *CreateSolidProtectionTrialsRequest) (response *CreateSolidProtectionTrialsResponse, err error) {
	if request == nil {
		request = NewCreateSolidProtectionTrialsRequest()
	}
	response = NewCreateSolidProtectionTrialsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCFGRiskAdvanceCFGListRequest() (request *DescribeCFGRiskAdvanceCFGListRequest) {
	request = &DescribeCFGRiskAdvanceCFGListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCFGRiskAdvanceCFGList")
	return
}

func NewDescribeCFGRiskAdvanceCFGListResponse() (response *DescribeCFGRiskAdvanceCFGListResponse) {
	response = &DescribeCFGRiskAdvanceCFGListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配置风险高级配置列表
func (c *Client) DescribeCFGRiskAdvanceCFGList(request *DescribeCFGRiskAdvanceCFGListRequest) (response *DescribeCFGRiskAdvanceCFGListResponse, err error) {
	if request == nil {
		request = NewDescribeCFGRiskAdvanceCFGListRequest()
	}
	response = NewDescribeCFGRiskAdvanceCFGListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCFWStatusRequest() (request *DescribeCFWStatusRequest) {
	request = &DescribeCFWStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCFWStatus")
	return
}

func NewDescribeCFWStatusResponse() (response *DescribeCFWStatusResponse) {
	response = &DescribeCFWStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云防用户信息
func (c *Client) DescribeCFWStatus(request *DescribeCFWStatusRequest) (response *DescribeCFWStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCFWStatusRequest()
	}
	response = NewDescribeCFWStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRegisterBASAssetAgentRequest() (request *RegisterBASAssetAgentRequest) {
	request = &RegisterBASAssetAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "RegisterBASAssetAgent")
	return
}

func NewRegisterBASAssetAgentResponse() (response *RegisterBASAssetAgentResponse) {
	response = &RegisterBASAssetAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安装BASagent
func (c *Client) RegisterBASAssetAgent(request *RegisterBASAssetAgentRequest) (response *RegisterBASAssetAgentResponse, err error) {
	if request == nil {
		request = NewRegisterBASAssetAgentRequest()
	}
	response = NewRegisterBASAssetAgentResponse()
	err = c.Send(request, response)
	return
}

func NewExportRiskEvidenceRequest() (request *ExportRiskEvidenceRequest) {
	request = &ExportRiskEvidenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ExportRiskEvidence")
	return
}

func NewExportRiskEvidenceResponse() (response *ExportRiskEvidenceResponse) {
	response = &ExportRiskEvidenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取导出风险中心的风险数据证据的临时下载链接
func (c *Client) ExportRiskEvidence(request *ExportRiskEvidenceRequest) (response *ExportRiskEvidenceResponse, err error) {
	if request == nil {
		request = NewExportRiskEvidenceRequest()
	}
	response = NewExportRiskEvidenceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterServiceAssetsRequest() (request *DescribeClusterServiceAssetsRequest) {
	request = &DescribeClusterServiceAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeClusterServiceAssets")
	return
}

func NewDescribeClusterServiceAssetsResponse() (response *DescribeClusterServiceAssetsResponse) {
	response = &DescribeClusterServiceAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群service列表
func (c *Client) DescribeClusterServiceAssets(request *DescribeClusterServiceAssetsRequest) (response *DescribeClusterServiceAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterServiceAssetsRequest()
	}
	response = NewDescribeClusterServiceAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayAssetInfoRequest() (request *DescribeGatewayAssetInfoRequest) {
	request = &DescribeGatewayAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeGatewayAssetInfo")
	return
}

func NewDescribeGatewayAssetInfoResponse() (response *DescribeGatewayAssetInfoResponse) {
	response = &DescribeGatewayAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网关资产详情
func (c *Client) DescribeGatewayAssetInfo(request *DescribeGatewayAssetInfoRequest) (response *DescribeGatewayAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayAssetInfoRequest()
	}
	response = NewDescribeGatewayAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceTagGroupsRequest() (request *DescribeResourceTagGroupsRequest) {
	request = &DescribeResourceTagGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeResourceTagGroups")
	return
}

func NewDescribeResourceTagGroupsResponse() (response *DescribeResourceTagGroupsResponse) {
	response = &DescribeResourceTagGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产中心查询tag分组信息
func (c *Client) DescribeResourceTagGroups(request *DescribeResourceTagGroupsRequest) (response *DescribeResourceTagGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeResourceTagGroupsRequest()
	}
	response = NewDescribeResourceTagGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogStorageDayRequest() (request *ModifyLogStorageDayRequest) {
	request = &ModifyLogStorageDayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyLogStorageDay")
	return
}

func NewModifyLogStorageDayResponse() (response *ModifyLogStorageDayResponse) {
	response = &ModifyLogStorageDayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户日志存储时长
func (c *Client) ModifyLogStorageDay(request *ModifyLogStorageDayRequest) (response *ModifyLogStorageDayResponse, err error) {
	if request == nil {
		request = NewModifyLogStorageDayRequest()
	}
	response = NewModifyLogStorageDayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRiskCenterScanTaskRequest() (request *ModifyRiskCenterScanTaskRequest) {
	request = &ModifyRiskCenterScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyRiskCenterScanTask")
	return
}

func NewModifyRiskCenterScanTaskResponse() (response *ModifyRiskCenterScanTaskResponse) {
	response = &ModifyRiskCenterScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改风险中心扫描任务
func (c *Client) ModifyRiskCenterScanTask(request *ModifyRiskCenterScanTaskRequest) (response *ModifyRiskCenterScanTaskResponse, err error) {
	if request == nil {
		request = NewModifyRiskCenterScanTaskRequest()
	}
	response = NewModifyRiskCenterScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayAssetsRequest() (request *DescribeGatewayAssetsRequest) {
	request = &DescribeGatewayAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeGatewayAssets")
	return
}

func NewDescribeGatewayAssetsResponse() (response *DescribeGatewayAssetsResponse) {
	response = &DescribeGatewayAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网关列表
func (c *Client) DescribeGatewayAssets(request *DescribeGatewayAssetsRequest) (response *DescribeGatewayAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayAssetsRequest()
	}
	response = NewDescribeGatewayAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagGroupAssetsRequest() (request *DescribeTagGroupAssetsRequest) {
	request = &DescribeTagGroupAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeTagGroupAssets")
	return
}

func NewDescribeTagGroupAssetsResponse() (response *DescribeTagGroupAssetsResponse) {
	response = &DescribeTagGroupAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// tag下资产查询
func (c *Client) DescribeTagGroupAssets(request *DescribeTagGroupAssetsRequest) (response *DescribeTagGroupAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeTagGroupAssetsRequest()
	}
	response = NewDescribeTagGroupAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanReportListRequest() (request *DescribeScanReportListRequest) {
	request = &DescribeScanReportListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeScanReportList")
	return
}

func NewDescribeScanReportListResponse() (response *DescribeScanReportListResponse) {
	response = &DescribeScanReportListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取扫描报告列表
func (c *Client) DescribeScanReportList(request *DescribeScanReportListRequest) (response *DescribeScanReportListResponse, err error) {
	if request == nil {
		request = NewDescribeScanReportListRequest()
	}
	response = NewDescribeScanReportListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationInfoRequest() (request *DescribeOrganizationInfoRequest) {
	request = &DescribeOrganizationInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeOrganizationInfo")
	return
}

func NewDescribeOrganizationInfoResponse() (response *DescribeOrganizationInfoResponse) {
	response = &DescribeOrganizationInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集团账号详情
func (c *Client) DescribeOrganizationInfo(request *DescribeOrganizationInfoRequest) (response *DescribeOrganizationInfoResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationInfoRequest()
	}
	response = NewDescribeOrganizationInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityTrendRequest() (request *DescribeSecurityTrendRequest) {
	request = &DescribeSecurityTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSecurityTrend")
	return
}

func NewDescribeSecurityTrendResponse() (response *DescribeSecurityTrendResponse) {
	response = &DescribeSecurityTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询互联网暴露面和攻击趋势
func (c *Client) DescribeSecurityTrend(request *DescribeSecurityTrendRequest) (response *DescribeSecurityTrendResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityTrendRequest()
	}
	response = NewDescribeSecurityTrendResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskLogURLRequest() (request *DescribeTaskLogURLRequest) {
	request = &DescribeTaskLogURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeTaskLogURL")
	return
}

func NewDescribeTaskLogURLResponse() (response *DescribeTaskLogURLResponse) {
	response = &DescribeTaskLogURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取报告下载的临时链接
func (c *Client) DescribeTaskLogURL(request *DescribeTaskLogURLRequest) (response *DescribeTaskLogURLResponse, err error) {
	if request == nil {
		request = NewDescribeTaskLogURLRequest()
	}
	response = NewDescribeTaskLogURLResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountMapRequest() (request *DescribeAccountMapRequest) {
	request = &DescribeAccountMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAccountMap")
	return
}

func NewDescribeAccountMapResponse() (response *DescribeAccountMapResponse) {
	response = &DescribeAccountMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账号映射
func (c *Client) DescribeAccountMap(request *DescribeAccountMapRequest) (response *DescribeAccountMapResponse, err error) {
	if request == nil {
		request = NewDescribeAccountMapRequest()
	}
	response = NewDescribeAccountMapResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFreeLimitRequest() (request *DescribeFreeLimitRequest) {
	request = &DescribeFreeLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeFreeLimit")
	return
}

func NewDescribeFreeLimitResponse() (response *DescribeFreeLimitResponse) {
	response = &DescribeFreeLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询总览/待办免费额度相关信息
func (c *Client) DescribeFreeLimit(request *DescribeFreeLimitRequest) (response *DescribeFreeLimitResponse, err error) {
	if request == nil {
		request = NewDescribeFreeLimitRequest()
	}
	response = NewDescribeFreeLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePDFDownLoadLogRequest() (request *DescribePDFDownLoadLogRequest) {
	request = &DescribePDFDownLoadLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribePDFDownLoadLog")
	return
}

func NewDescribePDFDownLoadLogResponse() (response *DescribePDFDownLoadLogResponse) {
	response = &DescribePDFDownLoadLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务扫描报告下载记录
func (c *Client) DescribePDFDownLoadLog(request *DescribePDFDownLoadLogRequest) (response *DescribePDFDownLoadLogResponse, err error) {
	if request == nil {
		request = NewDescribePDFDownLoadLogRequest()
	}
	response = NewDescribePDFDownLoadLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicIpAssetInfoRequest() (request *DescribePublicIpAssetInfoRequest) {
	request = &DescribePublicIpAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribePublicIpAssetInfo")
	return
}

func NewDescribePublicIpAssetInfoResponse() (response *DescribePublicIpAssetInfoResponse) {
	response = &DescribePublicIpAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 公网ip资产详情
func (c *Client) DescribePublicIpAssetInfo(request *DescribePublicIpAssetInfoRequest) (response *DescribePublicIpAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribePublicIpAssetInfoRequest()
	}
	response = NewDescribePublicIpAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSecTrendListRequest() (request *DescribeAssetSecTrendListRequest) {
	request = &DescribeAssetSecTrendListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetSecTrendList")
	return
}

func NewDescribeAssetSecTrendListResponse() (response *DescribeAssetSecTrendListResponse) {
	response = &DescribeAssetSecTrendListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产安全趋势
func (c *Client) DescribeAssetSecTrendList(request *DescribeAssetSecTrendListRequest) (response *DescribeAssetSecTrendListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSecTrendListRequest()
	}
	response = NewDescribeAssetSecTrendListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterMachineAssetsRequest() (request *DescribeClusterMachineAssetsRequest) {
	request = &DescribeClusterMachineAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeClusterMachineAssets")
	return
}

func NewDescribeClusterMachineAssetsResponse() (response *DescribeClusterMachineAssetsResponse) {
	response = &DescribeClusterMachineAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 主机节点列表
func (c *Client) DescribeClusterMachineAssets(request *DescribeClusterMachineAssetsRequest) (response *DescribeClusterMachineAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterMachineAssetsRequest()
	}
	response = NewDescribeClusterMachineAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterAssetViewCFGRiskListRequest() (request *DescribeRiskCenterAssetViewCFGRiskListRequest) {
	request = &DescribeRiskCenterAssetViewCFGRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterAssetViewCFGRiskList")
	return
}

func NewDescribeRiskCenterAssetViewCFGRiskListResponse() (response *DescribeRiskCenterAssetViewCFGRiskListResponse) {
	response = &DescribeRiskCenterAssetViewCFGRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产视角的配置风险列表
func (c *Client) DescribeRiskCenterAssetViewCFGRiskList(request *DescribeRiskCenterAssetViewCFGRiskListRequest) (response *DescribeRiskCenterAssetViewCFGRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterAssetViewCFGRiskListRequest()
	}
	response = NewDescribeRiskCenterAssetViewCFGRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafTlsVersionRequest() (request *DescribeWafTlsVersionRequest) {
	request = &DescribeWafTlsVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWafTlsVersion")
	return
}

func NewDescribeWafTlsVersionResponse() (response *DescribeWafTlsVersionResponse) {
	response = &DescribeWafTlsVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTlsVersion waf接口 获取TLS版本信息
func (c *Client) DescribeWafTlsVersion(request *DescribeWafTlsVersionRequest) (response *DescribeWafTlsVersionResponse, err error) {
	if request == nil {
		request = NewDescribeWafTlsVersionRequest()
	}
	response = NewDescribeWafTlsVersionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplyTrialRequest() (request *CreateApplyTrialRequest) {
	request = &CreateApplyTrialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateApplyTrial")
	return
}

func NewCreateApplyTrialResponse() (response *CreateApplyTrialResponse) {
	response = &CreateApplyTrialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 试用防火墙
func (c *Client) CreateApplyTrial(request *CreateApplyTrialRequest) (response *CreateApplyTrialResponse, err error) {
	if request == nil {
		request = NewCreateApplyTrialRequest()
	}
	response = NewCreateApplyTrialResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWebServiceRequest() (request *DeleteWebServiceRequest) {
	request = &DeleteWebServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DeleteWebService")
	return
}

func NewDeleteWebServiceResponse() (response *DeleteWebServiceResponse) {
	response = &DeleteWebServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除web服务
func (c *Client) DeleteWebService(request *DeleteWebServiceRequest) (response *DeleteWebServiceResponse, err error) {
	if request == nil {
		request = NewDeleteWebServiceRequest()
	}
	response = NewDeleteWebServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllAssetListRequest() (request *DescribeAllAssetListRequest) {
	request = &DescribeAllAssetListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAllAssetList")
	return
}

func NewDescribeAllAssetListResponse() (response *DescribeAllAssetListResponse) {
	response = &DescribeAllAssetListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 全部资产下拉
func (c *Client) DescribeAllAssetList(request *DescribeAllAssetListRequest) (response *DescribeAllAssetListResponse, err error) {
	if request == nil {
		request = NewDescribeAllAssetListRequest()
	}
	response = NewDescribeAllAssetListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerAssetsRequest() (request *DescribeContainerAssetsRequest) {
	request = &DescribeContainerAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeContainerAssets")
	return
}

func NewDescribeContainerAssetsResponse() (response *DescribeContainerAssetsResponse) {
	response = &DescribeContainerAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// MemCount
func (c *Client) DescribeContainerAssets(request *DescribeContainerAssetsRequest) (response *DescribeContainerAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeContainerAssetsRequest()
	}
	response = NewDescribeContainerAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAccountInfoRequest() (request *DescribeUserAccountInfoRequest) {
	request = &DescribeUserAccountInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeUserAccountInfo")
	return
}

func NewDescribeUserAccountInfoResponse() (response *DescribeUserAccountInfoResponse) {
	response = &DescribeUserAccountInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户账户信息
func (c *Client) DescribeUserAccountInfo(request *DescribeUserAccountInfoRequest) (response *DescribeUserAccountInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserAccountInfoRequest()
	}
	response = NewDescribeUserAccountInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserCVEStatusRequest() (request *DescribeUserCVEStatusRequest) {
	request = &DescribeUserCVEStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeUserCVEStatus")
	return
}

func NewDescribeUserCVEStatusResponse() (response *DescribeUserCVEStatusResponse) {
	response = &DescribeUserCVEStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 判断接入waf的所有域名是否都开启虚拟补丁
func (c *Client) DescribeUserCVEStatus(request *DescribeUserCVEStatusRequest) (response *DescribeUserCVEStatusResponse, err error) {
	if request == nil {
		request = NewDescribeUserCVEStatusRequest()
	}
	response = NewDescribeUserCVEStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcAssetsRequest() (request *DescribeVpcAssetsRequest) {
	request = &DescribeVpcAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeVpcAssets")
	return
}

func NewDescribeVpcAssetsResponse() (response *DescribeVpcAssetsResponse) {
	response = &DescribeVpcAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取vpc列表
func (c *Client) DescribeVpcAssets(request *DescribeVpcAssetsRequest) (response *DescribeVpcAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcAssetsRequest()
	}
	response = NewDescribeVpcAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewAddNewBindRoleUserRequest() (request *AddNewBindRoleUserRequest) {
	request = &AddNewBindRoleUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "AddNewBindRoleUser")
	return
}

func NewAddNewBindRoleUserResponse() (response *AddNewBindRoleUserResponse) {
	response = &AddNewBindRoleUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// csip角色授权绑定接口
func (c *Client) AddNewBindRoleUser(request *AddNewBindRoleUserRequest) (response *AddNewBindRoleUserResponse, err error) {
	if request == nil {
		request = NewAddNewBindRoleUserRequest()
	}
	response = NewAddNewBindRoleUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
	request = &DescribeMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeMachines")
	return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
	response = &DescribeMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 机器Machines数量
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeMachinesRequest()
	}
	response = NewDescribeMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserLogAuditSettingRequest() (request *DescribeUserLogAuditSettingRequest) {
	request = &DescribeUserLogAuditSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeUserLogAuditSetting")
	return
}

func NewDescribeUserLogAuditSettingResponse() (response *DescribeUserLogAuditSettingResponse) {
	response = &DescribeUserLogAuditSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户日志审计配置
func (c *Client) DescribeUserLogAuditSetting(request *DescribeUserLogAuditSettingRequest) (response *DescribeUserLogAuditSettingResponse, err error) {
	if request == nil {
		request = NewDescribeUserLogAuditSettingRequest()
	}
	response = NewDescribeUserLogAuditSettingResponse()
	err = c.Send(request, response)
	return
}

func NewCreateACRulesRequest() (request *CreateACRulesRequest) {
	request = &CreateACRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateACRules")
	return
}

func NewCreateACRulesResponse() (response *CreateACRulesResponse) {
	response = &CreateACRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建访问控制规则
func (c *Client) CreateACRules(request *CreateACRulesRequest) (response *CreateACRulesResponse, err error) {
	if request == nil {
		request = NewCreateACRulesRequest()
	}
	response = NewCreateACRulesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRiskCenterScanTaskRequest() (request *CreateRiskCenterScanTaskRequest) {
	request = &CreateRiskCenterScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateRiskCenterScanTask")
	return
}

func NewCreateRiskCenterScanTaskResponse() (response *CreateRiskCenterScanTaskResponse) {
	response = &CreateRiskCenterScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建风险中心扫描任务
func (c *Client) CreateRiskCenterScanTask(request *CreateRiskCenterScanTaskRequest) (response *CreateRiskCenterScanTaskResponse, err error) {
	if request == nil {
		request = NewCreateRiskCenterScanTaskRequest()
	}
	response = NewCreateRiskCenterScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerInfoRequest() (request *DescribeContainerInfoRequest) {
	request = &DescribeContainerInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeContainerInfo")
	return
}

func NewDescribeContainerInfoResponse() (response *DescribeContainerInfoResponse) {
	response = &DescribeContainerInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像信息
func (c *Client) DescribeContainerInfo(request *DescribeContainerInfoRequest) (response *DescribeContainerInfoResponse, err error) {
	if request == nil {
		request = NewDescribeContainerInfoRequest()
	}
	response = NewDescribeContainerInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskCenterRiskOverviewRequest() (request *DescribeRiskCenterRiskOverviewRequest) {
	request = &DescribeRiskCenterRiskOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeRiskCenterRiskOverview")
	return
}

func NewDescribeRiskCenterRiskOverviewResponse() (response *DescribeRiskCenterRiskOverviewResponse) {
	response = &DescribeRiskCenterRiskOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取风险中心风险概况
func (c *Client) DescribeRiskCenterRiskOverview(request *DescribeRiskCenterRiskOverviewRequest) (response *DescribeRiskCenterRiskOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeRiskCenterRiskOverviewRequest()
	}
	response = NewDescribeRiskCenterRiskOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIpProtectRequest() (request *ModifyIpProtectRequest) {
	request = &ModifyIpProtectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyIpProtect")
	return
}

func NewModifyIpProtectResponse() (response *ModifyIpProtectResponse) {
	response = &ModifyIpProtectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改公网ip防护
func (c *Client) ModifyIpProtect(request *ModifyIpProtectRequest) (response *ModifyIpProtectResponse, err error) {
	if request == nil {
		request = NewModifyIpProtectRequest()
	}
	response = NewModifyIpProtectResponse()
	err = c.Send(request, response)
	return
}

func NewExportSCCCSVRequest() (request *ExportSCCCSVRequest) {
	request = &ExportSCCCSVRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ExportSCCCSV")
	return
}

func NewExportSCCCSVResponse() (response *ExportSCCCSVResponse) {
	response = &ExportSCCCSVResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取导出文件的临时下载链接
func (c *Client) ExportSCCCSV(request *ExportSCCCSVRequest) (response *ExportSCCCSVResponse, err error) {
	if request == nil {
		request = NewExportSCCCSVRequest()
	}
	response = NewExportSCCCSVResponse()
	err = c.Send(request, response)
	return
}

func NewAddBASTaskRequest() (request *AddBASTaskRequest) {
	request = &AddBASTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "AddBASTask")
	return
}

func NewAddBASTaskResponse() (response *AddBASTaskResponse) {
	response = &AddBASTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行BAS任务
func (c *Client) AddBASTask(request *AddBASTaskRequest) (response *AddBASTaskResponse, err error) {
	if request == nil {
		request = NewAddBASTaskRequest()
	}
	response = NewAddBASTaskResponse()
	err = c.Send(request, response)
	return
}

func NewAddWebServiceRequest() (request *AddWebServiceRequest) {
	request = &AddWebServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "AddWebService")
	return
}

func NewAddWebServiceResponse() (response *AddWebServiceResponse) {
	response = &AddWebServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增web服务
func (c *Client) AddWebService(request *AddWebServiceRequest) (response *AddWebServiceResponse, err error) {
	if request == nil {
		request = NewAddWebServiceRequest()
	}
	response = NewAddWebServiceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNeedToSolveListRequest() (request *UpdateNeedToSolveListRequest) {
	request = &UpdateNeedToSolveListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "UpdateNeedToSolveList")
	return
}

func NewUpdateNeedToSolveListResponse() (response *UpdateNeedToSolveListResponse) {
	response = &UpdateNeedToSolveListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当有用户处理待办之后，通过该接口更新最新的待办信息
func (c *Client) UpdateNeedToSolveList(request *UpdateNeedToSolveListRequest) (response *UpdateNeedToSolveListResponse, err error) {
	if request == nil {
		request = NewUpdateNeedToSolveListRequest()
	}
	response = NewUpdateNeedToSolveListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetTopXRequest() (request *DescribeAssetTopXRequest) {
	request = &DescribeAssetTopXRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetTopX")
	return
}

func NewDescribeAssetTopXResponse() (response *DescribeAssetTopXResponse) {
	response = &DescribeAssetTopXResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// top统计
func (c *Client) DescribeAssetTopX(request *DescribeAssetTopXRequest) (response *DescribeAssetTopXResponse, err error) {
	if request == nil {
		request = NewDescribeAssetTopXRequest()
	}
	response = NewDescribeAssetTopXResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafAllInstancesRequest() (request *DescribeWafAllInstancesRequest) {
	request = &DescribeWafAllInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWafAllInstances")
	return
}

func NewDescribeWafAllInstancesResponse() (response *DescribeWafAllInstancesResponse) {
	response = &DescribeWafAllInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用waf DescribeInstances  查询用户waf试用过期时间
func (c *Client) DescribeWafAllInstances(request *DescribeWafAllInstancesRequest) (response *DescribeWafAllInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeWafAllInstancesRequest()
	}
	response = NewDescribeWafAllInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmStatisticsRequest() (request *DescribeAlarmStatisticsRequest) {
	request = &DescribeAlarmStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAlarmStatistics")
	return
}

func NewDescribeAlarmStatisticsResponse() (response *DescribeAlarmStatisticsResponse) {
	response = &DescribeAlarmStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警概览
func (c *Client) DescribeAlarmStatistics(request *DescribeAlarmStatisticsRequest) (response *DescribeAlarmStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmStatisticsRequest()
	}
	response = NewDescribeAlarmStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCoreAssetRequest() (request *DescribeCoreAssetRequest) {
	request = &DescribeCoreAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeCoreAsset")
	return
}

func NewDescribeCoreAssetResponse() (response *DescribeCoreAssetResponse) {
	response = &DescribeCoreAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 快速体检拉取核心资产用
func (c *Client) DescribeCoreAsset(request *DescribeCoreAssetRequest) (response *DescribeCoreAssetResponse, err error) {
	if request == nil {
		request = NewDescribeCoreAssetRequest()
	}
	response = NewDescribeCoreAssetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWafSpartaProtectionRequest() (request *CreateWafSpartaProtectionRequest) {
	request = &CreateWafSpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateWafSpartaProtection")
	return
}

func NewCreateWafSpartaProtectionResponse() (response *CreateWafSpartaProtectionResponse) {
	response = &CreateWafSpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// AddSpartaProtection waf接口 开启防护
func (c *Client) CreateWafSpartaProtection(request *CreateWafSpartaProtectionRequest) (response *CreateWafSpartaProtectionResponse, err error) {
	if request == nil {
		request = NewCreateWafSpartaProtectionRequest()
	}
	response = NewCreateWafSpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterPodAssetsRequest() (request *DescribeClusterPodAssetsRequest) {
	request = &DescribeClusterPodAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeClusterPodAssets")
	return
}

func NewDescribeClusterPodAssetsResponse() (response *DescribeClusterPodAssetsResponse) {
	response = &DescribeClusterPodAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群pod列表
func (c *Client) DescribeClusterPodAssets(request *DescribeClusterPodAssetsRequest) (response *DescribeClusterPodAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterPodAssetsRequest()
	}
	response = NewDescribeClusterPodAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePortRiskAdvanceCFGListRequest() (request *DescribePortRiskAdvanceCFGListRequest) {
	request = &DescribePortRiskAdvanceCFGListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribePortRiskAdvanceCFGList")
	return
}

func NewDescribePortRiskAdvanceCFGListResponse() (response *DescribePortRiskAdvanceCFGListResponse) {
	response = &DescribePortRiskAdvanceCFGListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询端口风险高级配置
func (c *Client) DescribePortRiskAdvanceCFGList(request *DescribePortRiskAdvanceCFGListRequest) (response *DescribePortRiskAdvanceCFGListResponse, err error) {
	if request == nil {
		request = NewDescribePortRiskAdvanceCFGListRequest()
	}
	response = NewDescribePortRiskAdvanceCFGListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchExamStatusRequest() (request *DescribeSearchExamStatusRequest) {
	request = &DescribeSearchExamStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeSearchExamStatus")
	return
}

func NewDescribeSearchExamStatusResponse() (response *DescribeSearchExamStatusResponse) {
	response = &DescribeSearchExamStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用来查询体检相关信息和状态
func (c *Client) DescribeSearchExamStatus(request *DescribeSearchExamStatusRequest) (response *DescribeSearchExamStatusResponse, err error) {
	if request == nil {
		request = NewDescribeSearchExamStatusRequest()
	}
	response = NewDescribeSearchExamStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateQuickProtectSettingRequest() (request *CreateQuickProtectSettingRequest) {
	request = &CreateQuickProtectSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "CreateQuickProtectSetting")
	return
}

func NewCreateQuickProtectSettingResponse() (response *CreateQuickProtectSettingResponse) {
	response = &CreateQuickProtectSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键开启防护
func (c *Client) CreateQuickProtectSetting(request *CreateQuickProtectSettingRequest) (response *CreateQuickProtectSettingResponse, err error) {
	if request == nil {
		request = NewCreateQuickProtectSettingRequest()
	}
	response = NewCreateQuickProtectSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeABTestConfigRequest() (request *DescribeABTestConfigRequest) {
	request = &DescribeABTestConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeABTestConfig")
	return
}

func NewDescribeABTestConfigResponse() (response *DescribeABTestConfigResponse) {
	response = &DescribeABTestConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户当前灰度配置
func (c *Client) DescribeABTestConfig(request *DescribeABTestConfigRequest) (response *DescribeABTestConfigResponse, err error) {
	if request == nil {
		request = NewDescribeABTestConfigRequest()
	}
	response = NewDescribeABTestConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebAssetListRequest() (request *DescribeWebAssetListRequest) {
	request = &DescribeWebAssetListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeWebAssetList")
	return
}

func NewDescribeWebAssetListResponse() (response *DescribeWebAssetListResponse) {
	response = &DescribeWebAssetListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 全部web可用资产下拉
func (c *Client) DescribeWebAssetList(request *DescribeWebAssetListRequest) (response *DescribeWebAssetListResponse, err error) {
	if request == nil {
		request = NewDescribeWebAssetListRequest()
	}
	response = NewDescribeWebAssetListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetStatisticsRequest() (request *DescribeAssetStatisticsRequest) {
	request = &DescribeAssetStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeAssetStatistics")
	return
}

func NewDescribeAssetStatisticsResponse() (response *DescribeAssetStatisticsResponse) {
	response = &DescribeAssetStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产中心统计概览
func (c *Client) DescribeAssetStatistics(request *DescribeAssetStatisticsRequest) (response *DescribeAssetStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeAssetStatisticsRequest()
	}
	response = NewDescribeAssetStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogAuditRequest() (request *DescribeLogAuditRequest) {
	request = &DescribeLogAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeLogAudit")
	return
}

func NewDescribeLogAuditResponse() (response *DescribeLogAuditResponse) {
	response = &DescribeLogAuditResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志审计查询列表接口
func (c *Client) DescribeLogAudit(request *DescribeLogAuditRequest) (response *DescribeLogAuditResponse, err error) {
	if request == nil {
		request = NewDescribeLogAuditRequest()
	}
	response = NewDescribeLogAuditResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPortProtectStatusRequest() (request *ModifyPortProtectStatusRequest) {
	request = &ModifyPortProtectStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyPortProtectStatus")
	return
}

func NewModifyPortProtectStatusResponse() (response *ModifyPortProtectStatusResponse) {
	response = &ModifyPortProtectStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改端口风险cfw的防护状态
func (c *Client) ModifyPortProtectStatus(request *ModifyPortProtectStatusRequest) (response *ModifyPortProtectStatusResponse, err error) {
	if request == nil {
		request = NewModifyPortProtectStatusRequest()
	}
	response = NewModifyPortProtectStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndeterminacyAssetTrendRequest() (request *DescribeIndeterminacyAssetTrendRequest) {
	request = &DescribeIndeterminacyAssetTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "DescribeIndeterminacyAssetTrend")
	return
}

func NewDescribeIndeterminacyAssetTrendResponse() (response *DescribeIndeterminacyAssetTrendResponse) {
	response = &DescribeIndeterminacyAssetTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通用趋势图
func (c *Client) DescribeIndeterminacyAssetTrend(request *DescribeIndeterminacyAssetTrendRequest) (response *DescribeIndeterminacyAssetTrendResponse, err error) {
	if request == nil {
		request = NewDescribeIndeterminacyAssetTrendRequest()
	}
	response = NewDescribeIndeterminacyAssetTrendResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserCVEStatusRequest() (request *ModifyUserCVEStatusRequest) {
	request = &ModifyUserCVEStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csip", APIVersion, "ModifyUserCVEStatus")
	return
}

func NewModifyUserCVEStatusResponse() (response *ModifyUserCVEStatusResponse) {
	response = &ModifyUserCVEStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启waf中指定漏洞的规则开关
func (c *Client) ModifyUserCVEStatus(request *ModifyUserCVEStatusRequest) (response *ModifyUserCVEStatusResponse, err error) {
	if request == nil {
		request = NewModifyUserCVEStatusRequest()
	}
	response = NewModifyUserCVEStatusResponse()
	err = c.Send(request, response)
	return
}
