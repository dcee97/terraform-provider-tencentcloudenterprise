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

package v20230512

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2023-05-12"

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

func NewApprovedSystemBatchRequest() (request *ApprovedSystemBatchRequest) {
	request = &ApprovedSystemBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ApprovedSystemBatch")
	return
}

func NewApprovedSystemBatchResponse() (response *ApprovedSystemBatchResponse) {
	response = &ApprovedSystemBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统审批同意批量
func (c *Client) ApprovedSystemBatch(request *ApprovedSystemBatchRequest) (response *ApprovedSystemBatchResponse, err error) {
	if request == nil {
		request = NewApprovedSystemBatchRequest()
	}
	response = NewApprovedSystemBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionRenewApproveLogsRequest() (request *DescribeSubscriptionRenewApproveLogsRequest) {
	request = &DescribeSubscriptionRenewApproveLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSubscriptionRenewApproveLogs")
	return
}

func NewDescribeSubscriptionRenewApproveLogsResponse() (response *DescribeSubscriptionRenewApproveLogsResponse) {
	response = &DescribeSubscriptionRenewApproveLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订阅续期审批历史记录
func (c *Client) DescribeSubscriptionRenewApproveLogs(request *DescribeSubscriptionRenewApproveLogsRequest) (response *DescribeSubscriptionRenewApproveLogsResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionRenewApproveLogsRequest()
	}
	response = NewDescribeSubscriptionRenewApproveLogsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMicroRegisterCenterRequest() (request *ModifyMicroRegisterCenterRequest) {
	request = &ModifyMicroRegisterCenterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyMicroRegisterCenter")
	return
}

func NewModifyMicroRegisterCenterResponse() (response *ModifyMicroRegisterCenterResponse) {
	response = &ModifyMicroRegisterCenterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务注册中心修改
func (c *Client) ModifyMicroRegisterCenter(request *ModifyMicroRegisterCenterRequest) (response *ModifyMicroRegisterCenterResponse, err error) {
	if request == nil {
		request = NewModifyMicroRegisterCenterRequest()
	}
	response = NewModifyMicroRegisterCenterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResetLeftMenuRequest() (request *ModifyResetLeftMenuRequest) {
	request = &ModifyResetLeftMenuRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyResetLeftMenu")
	return
}

func NewModifyResetLeftMenuResponse() (response *ModifyResetLeftMenuResponse) {
	response = &ModifyResetLeftMenuResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置左侧菜单
func (c *Client) ModifyResetLeftMenu(request *ModifyResetLeftMenuRequest) (response *ModifyResetLeftMenuResponse, err error) {
	if request == nil {
		request = NewModifyResetLeftMenuRequest()
	}
	response = NewModifyResetLeftMenuResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAcsRolePermissionRequest() (request *ModifyAcsRolePermissionRequest) {
	request = &ModifyAcsRolePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyAcsRolePermission")
	return
}

func NewModifyAcsRolePermissionResponse() (response *ModifyAcsRolePermissionResponse) {
	response = &ModifyAcsRolePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改角色权限
func (c *Client) ModifyAcsRolePermission(request *ModifyAcsRolePermissionRequest) (response *ModifyAcsRolePermissionResponse, err error) {
	if request == nil {
		request = NewModifyAcsRolePermissionRequest()
	}
	response = NewModifyAcsRolePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
	request = &CreateAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateAccount")
	return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
	response = &CreateAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建帐号
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
	if request == nil {
		request = NewCreateAccountRequest()
	}
	response = NewCreateAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroRegisterCentersRequest() (request *DeleteMicroRegisterCentersRequest) {
	request = &DeleteMicroRegisterCentersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMicroRegisterCenters")
	return
}

func NewDeleteMicroRegisterCentersResponse() (response *DeleteMicroRegisterCentersResponse) {
	response = &DeleteMicroRegisterCentersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务注册中心批量删除
func (c *Client) DeleteMicroRegisterCenters(request *DeleteMicroRegisterCentersRequest) (response *DeleteMicroRegisterCentersResponse, err error) {
	if request == nil {
		request = NewDeleteMicroRegisterCentersRequest()
	}
	response = NewDeleteMicroRegisterCentersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgenciesRequest() (request *DescribeAgenciesRequest) {
	request = &DescribeAgenciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAgencies")
	return
}

func NewDescribeAgenciesResponse() (response *DescribeAgenciesResponse) {
	response = &DescribeAgenciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询机构
func (c *Client) DescribeAgencies(request *DescribeAgenciesRequest) (response *DescribeAgenciesResponse, err error) {
	if request == nil {
		request = NewDescribeAgenciesRequest()
	}
	response = NewDescribeAgenciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseLifeTimeRequest() (request *DescribeLicenseLifeTimeRequest) {
	request = &DescribeLicenseLifeTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeLicenseLifeTime")
	return
}

func NewDescribeLicenseLifeTimeResponse() (response *DescribeLicenseLifeTimeResponse) {
	response = &DescribeLicenseLifeTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询证书有效期
func (c *Client) DescribeLicenseLifeTime(request *DescribeLicenseLifeTimeRequest) (response *DescribeLicenseLifeTimeResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseLifeTimeRequest()
	}
	response = NewDescribeLicenseLifeTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSmartGatewayGroupRequest() (request *DeleteSmartGatewayGroupRequest) {
	request = &DeleteSmartGatewayGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteSmartGatewayGroup")
	return
}

func NewDeleteSmartGatewayGroupResponse() (response *DeleteSmartGatewayGroupResponse) {
	response = &DeleteSmartGatewayGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除节点阵列
func (c *Client) DeleteSmartGatewayGroup(request *DeleteSmartGatewayGroupRequest) (response *DeleteSmartGatewayGroupResponse, err error) {
	if request == nil {
		request = NewDeleteSmartGatewayGroupRequest()
	}
	response = NewDeleteSmartGatewayGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroRegisterCenterRequest() (request *DescribeMicroRegisterCenterRequest) {
	request = &DescribeMicroRegisterCenterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroRegisterCenter")
	return
}

func NewDescribeMicroRegisterCenterResponse() (response *DescribeMicroRegisterCenterResponse) {
	response = &DescribeMicroRegisterCenterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务注册中心查询
func (c *Client) DescribeMicroRegisterCenter(request *DescribeMicroRegisterCenterRequest) (response *DescribeMicroRegisterCenterResponse, err error) {
	if request == nil {
		request = NewDescribeMicroRegisterCenterRequest()
	}
	response = NewDescribeMicroRegisterCenterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTlsCertificatesRequest() (request *DescribeTlsCertificatesRequest) {
	request = &DescribeTlsCertificatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTlsCertificates")
	return
}

func NewDescribeTlsCertificatesResponse() (response *DescribeTlsCertificatesResponse) {
	response = &DescribeTlsCertificatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询证书列表
func (c *Client) DescribeTlsCertificates(request *DescribeTlsCertificatesRequest) (response *DescribeTlsCertificatesResponse, err error) {
	if request == nil {
		request = NewDescribeTlsCertificatesRequest()
	}
	response = NewDescribeTlsCertificatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTSFMicroservicesRequest() (request *DescribeTSFMicroservicesRequest) {
	request = &DescribeTSFMicroservicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTSFMicroservices")
	return
}

func NewDescribeTSFMicroservicesResponse() (response *DescribeTSFMicroservicesResponse) {
	response = &DescribeTSFMicroservicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF 微服务实例列表
func (c *Client) DescribeTSFMicroservices(request *DescribeTSFMicroservicesRequest) (response *DescribeTSFMicroservicesResponse, err error) {
	if request == nil {
		request = NewDescribeTSFMicroservicesRequest()
	}
	response = NewDescribeTSFMicroservicesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogoRequest() (request *ModifyLogoRequest) {
	request = &ModifyLogoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyLogo")
	return
}

func NewModifyLogoResponse() (response *ModifyLogoResponse) {
	response = &ModifyLogoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑logo
func (c *Client) ModifyLogo(request *ModifyLogoRequest) (response *ModifyLogoResponse, err error) {
	if request == nil {
		request = NewModifyLogoRequest()
	}
	response = NewModifyLogoResponse()
	err = c.Send(request, response)
	return
}

func NewStatisticsAgencyRequest() (request *StatisticsAgencyRequest) {
	request = &StatisticsAgencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "StatisticsAgency")
	return
}

func NewStatisticsAgencyResponse() (response *StatisticsAgencyResponse) {
	response = &StatisticsAgencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计机构
func (c *Client) StatisticsAgency(request *StatisticsAgencyRequest) (response *StatisticsAgencyResponse, err error) {
	if request == nil {
		request = NewStatisticsAgencyRequest()
	}
	response = NewStatisticsAgencyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMenuRequest() (request *ModifyMenuRequest) {
	request = &ModifyMenuRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyMenu")
	return
}

func NewModifyMenuResponse() (response *ModifyMenuResponse) {
	response = &ModifyMenuResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑菜单
func (c *Client) ModifyMenu(request *ModifyMenuRequest) (response *ModifyMenuResponse, err error) {
	if request == nil {
		request = NewModifyMenuRequest()
	}
	response = NewModifyMenuResponse()
	err = c.Send(request, response)
	return
}

func NewRejectRenewSubscriptionBatchRequest() (request *RejectRenewSubscriptionBatchRequest) {
	request = &RejectRenewSubscriptionBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RejectRenewSubscriptionBatch")
	return
}

func NewRejectRenewSubscriptionBatchResponse() (response *RejectRenewSubscriptionBatchResponse) {
	response = &RejectRenewSubscriptionBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续期订阅拒绝
func (c *Client) RejectRenewSubscriptionBatch(request *RejectRenewSubscriptionBatchRequest) (response *RejectRenewSubscriptionBatchResponse, err error) {
	if request == nil {
		request = NewRejectRenewSubscriptionBatchRequest()
	}
	response = NewRejectRenewSubscriptionBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAlertRuleRequest() (request *DeleteAlertRuleRequest) {
	request = &DeleteAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteAlertRule")
	return
}

func NewDeleteAlertRuleResponse() (response *DeleteAlertRuleResponse) {
	response = &DeleteAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除告警规则
func (c *Client) DeleteAlertRule(request *DeleteAlertRuleRequest) (response *DeleteAlertRuleResponse, err error) {
	if request == nil {
		request = NewDeleteAlertRuleRequest()
	}
	response = NewDeleteAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroProtoGroupRequest() (request *DeleteMicroProtoGroupRequest) {
	request = &DeleteMicroProtoGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMicroProtoGroup")
	return
}

func NewDeleteMicroProtoGroupResponse() (response *DeleteMicroProtoGroupResponse) {
	response = &DeleteMicroProtoGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件组删除
func (c *Client) DeleteMicroProtoGroup(request *DeleteMicroProtoGroupRequest) (response *DeleteMicroProtoGroupResponse, err error) {
	if request == nil {
		request = NewDeleteMicroProtoGroupRequest()
	}
	response = NewDeleteMicroProtoGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroSyncJobRequest() (request *DescribeMicroSyncJobRequest) {
	request = &DescribeMicroSyncJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroSyncJob")
	return
}

func NewDescribeMicroSyncJobResponse() (response *DescribeMicroSyncJobResponse) {
	response = &DescribeMicroSyncJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务同步记录查询
func (c *Client) DescribeMicroSyncJob(request *DescribeMicroSyncJobRequest) (response *DescribeMicroSyncJobResponse, err error) {
	if request == nil {
		request = NewDescribeMicroSyncJobRequest()
	}
	response = NewDescribeMicroSyncJobResponse()
	err = c.Send(request, response)
	return
}

func NewReviveSubscriptionRequest() (request *ReviveSubscriptionRequest) {
	request = &ReviveSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ReviveSubscription")
	return
}

func NewReviveSubscriptionResponse() (response *ReviveSubscriptionResponse) {
	response = &ReviveSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订阅重新授权
func (c *Client) ReviveSubscription(request *ReviveSubscriptionRequest) (response *ReviveSubscriptionResponse, err error) {
	if request == nil {
		request = NewReviveSubscriptionRequest()
	}
	response = NewReviveSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlertRuleRequest() (request *ModifyAlertRuleRequest) {
	request = &ModifyAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyAlertRule")
	return
}

func NewModifyAlertRuleResponse() (response *ModifyAlertRuleResponse) {
	response = &ModifyAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警规则
func (c *Client) ModifyAlertRule(request *ModifyAlertRuleRequest) (response *ModifyAlertRuleResponse, err error) {
	if request == nil {
		request = NewModifyAlertRuleRequest()
	}
	response = NewModifyAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewApprovedSiteBatchRequest() (request *ApprovedSiteBatchRequest) {
	request = &ApprovedSiteBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ApprovedSiteBatch")
	return
}

func NewApprovedSiteBatchResponse() (response *ApprovedSiteBatchResponse) {
	response = &ApprovedSiteBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 站点审批同意批量
func (c *Client) ApprovedSiteBatch(request *ApprovedSiteBatchRequest) (response *ApprovedSiteBatchResponse, err error) {
	if request == nil {
		request = NewApprovedSiteBatchRequest()
	}
	response = NewApprovedSiteBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDocumentGroupListRequest() (request *DescribeDocumentGroupListRequest) {
	request = &DescribeDocumentGroupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeDocumentGroupList")
	return
}

func NewDescribeDocumentGroupListResponse() (response *DescribeDocumentGroupListResponse) {
	response = &DescribeDocumentGroupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文档分组列表
func (c *Client) DescribeDocumentGroupList(request *DescribeDocumentGroupListRequest) (response *DescribeDocumentGroupListResponse, err error) {
	if request == nil {
		request = NewDescribeDocumentGroupListRequest()
	}
	response = NewDescribeDocumentGroupListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstantMetricsRequest() (request *DescribeInstantMetricsRequest) {
	request = &DescribeInstantMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeInstantMetrics")
	return
}

func NewDescribeInstantMetricsResponse() (response *DescribeInstantMetricsResponse) {
	response = &DescribeInstantMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询瞬时指标
func (c *Client) DescribeInstantMetrics(request *DescribeInstantMetricsRequest) (response *DescribeInstantMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeInstantMetricsRequest()
	}
	response = NewDescribeInstantMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewGetCurrentUserRequest() (request *GetCurrentUserRequest) {
	request = &GetCurrentUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "GetCurrentUser")
	return
}

func NewGetCurrentUserResponse() (response *GetCurrentUserResponse) {
	response = &GetCurrentUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前登录的用户信息
func (c *Client) GetCurrentUser(request *GetCurrentUserRequest) (response *GetCurrentUserResponse, err error) {
	if request == nil {
		request = NewGetCurrentUserRequest()
	}
	response = NewGetCurrentUserResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTlsCertificateRequest() (request *CreateTlsCertificateRequest) {
	request = &CreateTlsCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateTlsCertificate")
	return
}

func NewCreateTlsCertificateResponse() (response *CreateTlsCertificateResponse) {
	response = &CreateTlsCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建证书
func (c *Client) CreateTlsCertificate(request *CreateTlsCertificateRequest) (response *CreateTlsCertificateResponse, err error) {
	if request == nil {
		request = NewCreateTlsCertificateRequest()
	}
	response = NewCreateTlsCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDocumentFileRequest() (request *DescribeDocumentFileRequest) {
	request = &DescribeDocumentFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeDocumentFile")
	return
}

func NewDescribeDocumentFileResponse() (response *DescribeDocumentFileResponse) {
	response = &DescribeDocumentFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取文档文件
func (c *Client) DescribeDocumentFile(request *DescribeDocumentFileRequest) (response *DescribeDocumentFileResponse, err error) {
	if request == nil {
		request = NewDescribeDocumentFileRequest()
	}
	response = NewDescribeDocumentFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSiteApprovesRequest() (request *DescribeSiteApprovesRequest) {
	request = &DescribeSiteApprovesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSiteApproves")
	return
}

func NewDescribeSiteApprovesResponse() (response *DescribeSiteApprovesResponse) {
	response = &DescribeSiteApprovesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询站点审批列表
func (c *Client) DescribeSiteApproves(request *DescribeSiteApprovesRequest) (response *DescribeSiteApprovesResponse, err error) {
	if request == nil {
		request = NewDescribeSiteApprovesRequest()
	}
	response = NewDescribeSiteApprovesResponse()
	err = c.Send(request, response)
	return
}

func NewSyncSmartGatewayGroupRequest() (request *SyncSmartGatewayGroupRequest) {
	request = &SyncSmartGatewayGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "SyncSmartGatewayGroup")
	return
}

func NewSyncSmartGatewayGroupResponse() (response *SyncSmartGatewayGroupResponse) {
	response = &SyncSmartGatewayGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步节点阵列
func (c *Client) SyncSmartGatewayGroup(request *SyncSmartGatewayGroupRequest) (response *SyncSmartGatewayGroupResponse, err error) {
	if request == nil {
		request = NewSyncSmartGatewayGroupRequest()
	}
	response = NewSyncSmartGatewayGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceOfSubAppRequest() (request *DescribeServiceOfSubAppRequest) {
	request = &DescribeServiceOfSubAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeServiceOfSubApp")
	return
}

func NewDescribeServiceOfSubAppResponse() (response *DescribeServiceOfSubAppResponse) {
	response = &DescribeServiceOfSubAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 应用订阅第三方API列表
func (c *Client) DescribeServiceOfSubApp(request *DescribeServiceOfSubAppRequest) (response *DescribeServiceOfSubAppResponse, err error) {
	if request == nil {
		request = NewDescribeServiceOfSubAppRequest()
	}
	response = NewDescribeServiceOfSubAppResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAgencyRequest() (request *DeleteAgencyRequest) {
	request = &DeleteAgencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteAgency")
	return
}

func NewDeleteAgencyResponse() (response *DeleteAgencyResponse) {
	response = &DeleteAgencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机构
func (c *Client) DeleteAgency(request *DeleteAgencyRequest) (response *DeleteAgencyResponse, err error) {
	if request == nil {
		request = NewDeleteAgencyRequest()
	}
	response = NewDeleteAgencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAdvanceTemplateRequest() (request *DescribeAdvanceTemplateRequest) {
	request = &DescribeAdvanceTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAdvanceTemplate")
	return
}

func NewDescribeAdvanceTemplateResponse() (response *DescribeAdvanceTemplateResponse) {
	response = &DescribeAdvanceTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取高级模板列表（同站点）
func (c *Client) DescribeAdvanceTemplate(request *DescribeAdvanceTemplateRequest) (response *DescribeAdvanceTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeAdvanceTemplateRequest()
	}
	response = NewDescribeAdvanceTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAppUniqueRequest() (request *CheckAppUniqueRequest) {
	request = &CheckAppUniqueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CheckAppUnique")
	return
}

func NewCheckAppUniqueResponse() (response *CheckAppUniqueResponse) {
	response = &CheckAppUniqueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CheckAppUnique
func (c *Client) CheckAppUnique(request *CheckAppUniqueRequest) (response *CheckAppUniqueResponse, err error) {
	if request == nil {
		request = NewCheckAppUniqueRequest()
	}
	response = NewCheckAppUniqueResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroRegisterCentersRequest() (request *DescribeMicroRegisterCentersRequest) {
	request = &DescribeMicroRegisterCentersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroRegisterCenters")
	return
}

func NewDescribeMicroRegisterCentersResponse() (response *DescribeMicroRegisterCentersResponse) {
	response = &DescribeMicroRegisterCentersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务注册中心列表
func (c *Client) DescribeMicroRegisterCenters(request *DescribeMicroRegisterCentersRequest) (response *DescribeMicroRegisterCentersResponse, err error) {
	if request == nil {
		request = NewDescribeMicroRegisterCentersRequest()
	}
	response = NewDescribeMicroRegisterCentersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSqlTemplateRequest() (request *CreateSqlTemplateRequest) {
	request = &CreateSqlTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateSqlTemplate")
	return
}

func NewCreateSqlTemplateResponse() (response *CreateSqlTemplateResponse) {
	response = &CreateSqlTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SQL 模板创建
func (c *Client) CreateSqlTemplate(request *CreateSqlTemplateRequest) (response *CreateSqlTemplateResponse, err error) {
	if request == nil {
		request = NewCreateSqlTemplateRequest()
	}
	response = NewCreateSqlTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgencyConfigsRequest() (request *DescribeAgencyConfigsRequest) {
	request = &DescribeAgencyConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAgencyConfigs")
	return
}

func NewDescribeAgencyConfigsResponse() (response *DescribeAgencyConfigsResponse) {
	response = &DescribeAgencyConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机构流程设置
func (c *Client) DescribeAgencyConfigs(request *DescribeAgencyConfigsRequest) (response *DescribeAgencyConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeAgencyConfigsRequest()
	}
	response = NewDescribeAgencyConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewRestoreAppRequest() (request *RestoreAppRequest) {
	request = &RestoreAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RestoreApp")
	return
}

func NewRestoreAppResponse() (response *RestoreAppResponse) {
	response = &RestoreAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恢复应用
func (c *Client) RestoreApp(request *RestoreAppRequest) (response *RestoreAppResponse, err error) {
	if request == nil {
		request = NewRestoreAppRequest()
	}
	response = NewRestoreAppResponse()
	err = c.Send(request, response)
	return
}

func NewUnFreezeRequest() (request *UnFreezeRequest) {
	request = &UnFreezeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UnFreeze")
	return
}

func NewUnFreezeResponse() (response *UnFreezeResponse) {
	response = &UnFreezeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解冻API服务
func (c *Client) UnFreeze(request *UnFreezeRequest) (response *UnFreezeResponse, err error) {
	if request == nil {
		request = NewUnFreezeRequest()
	}
	response = NewUnFreezeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTSFGatewayServiceRequest() (request *CreateTSFGatewayServiceRequest) {
	request = &CreateTSFGatewayServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateTSFGatewayService")
	return
}

func NewCreateTSFGatewayServiceResponse() (response *CreateTSFGatewayServiceResponse) {
	response = &CreateTSFGatewayServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步租户下 TSF 网关指定 API
func (c *Client) CreateTSFGatewayService(request *CreateTSFGatewayServiceRequest) (response *CreateTSFGatewayServiceResponse, err error) {
	if request == nil {
		request = NewCreateTSFGatewayServiceRequest()
	}
	response = NewCreateTSFGatewayServiceResponse()
	err = c.Send(request, response)
	return
}

func NewApproveRenewSubscriptionRequest() (request *ApproveRenewSubscriptionRequest) {
	request = &ApproveRenewSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ApproveRenewSubscription")
	return
}

func NewApproveRenewSubscriptionResponse() (response *ApproveRenewSubscriptionResponse) {
	response = &ApproveRenewSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续期订阅审批
func (c *Client) ApproveRenewSubscription(request *ApproveRenewSubscriptionRequest) (response *ApproveRenewSubscriptionResponse, err error) {
	if request == nil {
		request = NewApproveRenewSubscriptionRequest()
	}
	response = NewApproveRenewSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAgencyRequest() (request *CheckAgencyRequest) {
	request = &CheckAgencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CheckAgency")
	return
}

func NewCheckAgencyResponse() (response *CheckAgencyResponse) {
	response = &CheckAgencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证机构名称标识是否合法
func (c *Client) CheckAgency(request *CheckAgencyRequest) (response *CheckAgencyResponse, err error) {
	if request == nil {
		request = NewCheckAgencyRequest()
	}
	response = NewCheckAgencyResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadAccRequest() (request *DownloadAccRequest) {
	request = &DownloadAccRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadAcc")
	return
}

func NewDownloadAccResponse() (response *DownloadAccResponse) {
	response = &DownloadAccResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导出用户
func (c *Client) DownloadAcc(request *DownloadAccRequest) (response *DownloadAccResponse, err error) {
	if request == nil {
		request = NewDownloadAccRequest()
	}
	response = NewDownloadAccResponse()
	err = c.Send(request, response)
	return
}

func NewRejectedSystemRequest() (request *RejectedSystemRequest) {
	request = &RejectedSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RejectedSystem")
	return
}

func NewRejectedSystemResponse() (response *RejectedSystemResponse) {
	response = &RejectedSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统审批同意
func (c *Client) RejectedSystem(request *RejectedSystemRequest) (response *RejectedSystemResponse, err error) {
	if request == nil {
		request = NewRejectedSystemRequest()
	}
	response = NewRejectedSystemResponse()
	err = c.Send(request, response)
	return
}

func NewCountVendorRequest() (request *CountVendorRequest) {
	request = &CountVendorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CountVendor")
	return
}

func NewCountVendorResponse() (response *CountVendorResponse) {
	response = &CountVendorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计供应商信息
func (c *Client) CountVendor(request *CountVendorRequest) (response *CountVendorResponse, err error) {
	if request == nil {
		request = NewCountVendorRequest()
	}
	response = NewCountVendorResponse()
	err = c.Send(request, response)
	return
}

func NewDisabledSiteRequest() (request *DisabledSiteRequest) {
	request = &DisabledSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DisabledSite")
	return
}

func NewDisabledSiteResponse() (response *DisabledSiteResponse) {
	response = &DisabledSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用站点
func (c *Client) DisabledSite(request *DisabledSiteRequest) (response *DisabledSiteResponse, err error) {
	if request == nil {
		request = NewDisabledSiteRequest()
	}
	response = NewDisabledSiteResponse()
	err = c.Send(request, response)
	return
}

func NewVerifySmsCodeRequest() (request *VerifySmsCodeRequest) {
	request = &VerifySmsCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "VerifySmsCode")
	return
}

func NewVerifySmsCodeResponse() (response *VerifySmsCodeResponse) {
	response = &VerifySmsCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证短信验证码
func (c *Client) VerifySmsCode(request *VerifySmsCodeRequest) (response *VerifySmsCodeResponse, err error) {
	if request == nil {
		request = NewVerifySmsCodeRequest()
	}
	response = NewVerifySmsCodeResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTlsCertificateRequest() (request *UpdateTlsCertificateRequest) {
	request = &UpdateTlsCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UpdateTlsCertificate")
	return
}

func NewUpdateTlsCertificateResponse() (response *UpdateTlsCertificateResponse) {
	response = &UpdateTlsCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改证书
func (c *Client) UpdateTlsCertificate(request *UpdateTlsCertificateRequest) (response *UpdateTlsCertificateResponse, err error) {
	if request == nil {
		request = NewUpdateTlsCertificateRequest()
	}
	response = NewUpdateTlsCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAreaRequest() (request *ModifyAreaRequest) {
	request = &ModifyAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyArea")
	return
}

func NewModifyAreaResponse() (response *ModifyAreaResponse) {
	response = &ModifyAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改区域
func (c *Client) ModifyArea(request *ModifyAreaRequest) (response *ModifyAreaResponse, err error) {
	if request == nil {
		request = NewModifyAreaRequest()
	}
	response = NewModifyAreaResponse()
	err = c.Send(request, response)
	return
}

func NewFreezeRequest() (request *FreezeRequest) {
	request = &FreezeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "Freeze")
	return
}

func NewFreezeResponse() (response *FreezeResponse) {
	response = &FreezeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 冻结API服务
func (c *Client) Freeze(request *FreezeRequest) (response *FreezeResponse, err error) {
	if request == nil {
		request = NewFreezeRequest()
	}
	response = NewFreezeResponse()
	err = c.Send(request, response)
	return
}

func NewAdvanceTemplateBtnFunctionRequest() (request *AdvanceTemplateBtnFunctionRequest) {
	request = &AdvanceTemplateBtnFunctionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "AdvanceTemplateBtnFunction")
	return
}

func NewAdvanceTemplateBtnFunctionResponse() (response *AdvanceTemplateBtnFunctionResponse) {
	response = &AdvanceTemplateBtnFunctionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 高级模板按钮点击功能
func (c *Client) AdvanceTemplateBtnFunction(request *AdvanceTemplateBtnFunctionRequest) (response *AdvanceTemplateBtnFunctionResponse, err error) {
	if request == nil {
		request = NewAdvanceTemplateBtnFunctionRequest()
	}
	response = NewAdvanceTemplateBtnFunctionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProfileRequest() (request *DescribeProfileRequest) {
	request = &DescribeProfileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeProfile")
	return
}

func NewDescribeProfileResponse() (response *DescribeProfileResponse) {
	response = &DescribeProfileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAppProfile
func (c *Client) DescribeProfile(request *DescribeProfileRequest) (response *DescribeProfileResponse, err error) {
	if request == nil {
		request = NewDescribeProfileRequest()
	}
	response = NewDescribeProfileResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEnableAccountRequest() (request *ModifyEnableAccountRequest) {
	request = &ModifyEnableAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyEnableAccount")
	return
}

func NewModifyEnableAccountResponse() (response *ModifyEnableAccountResponse) {
	response = &ModifyEnableAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用帐号
func (c *Client) ModifyEnableAccount(request *ModifyEnableAccountRequest) (response *ModifyEnableAccountResponse, err error) {
	if request == nil {
		request = NewModifyEnableAccountRequest()
	}
	response = NewModifyEnableAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMicroTenantRequest() (request *CreateMicroTenantRequest) {
	request = &CreateMicroTenantRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateMicroTenant")
	return
}

func NewCreateMicroTenantResponse() (response *CreateMicroTenantResponse) {
	response = &CreateMicroTenantResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务租户创建
func (c *Client) CreateMicroTenant(request *CreateMicroTenantRequest) (response *CreateMicroTenantResponse, err error) {
	if request == nil {
		request = NewCreateMicroTenantRequest()
	}
	response = NewCreateMicroTenantResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMemberRequest() (request *UpdateMemberRequest) {
	request = &UpdateMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UpdateMember")
	return
}

func NewUpdateMemberResponse() (response *UpdateMemberResponse) {
	response = &UpdateMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑机构成员
func (c *Client) UpdateMember(request *UpdateMemberRequest) (response *UpdateMemberResponse, err error) {
	if request == nil {
		request = NewUpdateMemberRequest()
	}
	response = NewUpdateMemberResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTSFRegionListsRequest() (request *DescribeTSFRegionListsRequest) {
	request = &DescribeTSFRegionListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTSFRegionLists")
	return
}

func NewDescribeTSFRegionListsResponse() (response *DescribeTSFRegionListsResponse) {
	response = &DescribeTSFRegionListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF 微服务区域列表
func (c *Client) DescribeTSFRegionLists(request *DescribeTSFRegionListsRequest) (response *DescribeTSFRegionListsResponse, err error) {
	if request == nil {
		request = NewDescribeTSFRegionListsRequest()
	}
	response = NewDescribeTSFRegionListsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroRegisterCenterRequest() (request *DeleteMicroRegisterCenterRequest) {
	request = &DeleteMicroRegisterCenterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMicroRegisterCenter")
	return
}

func NewDeleteMicroRegisterCenterResponse() (response *DeleteMicroRegisterCenterResponse) {
	response = &DeleteMicroRegisterCenterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务注册中心删除
func (c *Client) DeleteMicroRegisterCenter(request *DeleteMicroRegisterCenterRequest) (response *DeleteMicroRegisterCenterResponse, err error) {
	if request == nil {
		request = NewDeleteMicroRegisterCenterRequest()
	}
	response = NewDeleteMicroRegisterCenterResponse()
	err = c.Send(request, response)
	return
}

func NewRejectedServiceRequest() (request *RejectedServiceRequest) {
	request = &RejectedServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RejectedService")
	return
}

func NewRejectedServiceResponse() (response *RejectedServiceResponse) {
	response = &RejectedServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务审批驳回
func (c *Client) RejectedService(request *RejectedServiceRequest) (response *RejectedServiceResponse, err error) {
	if request == nil {
		request = NewRejectedServiceRequest()
	}
	response = NewRejectedServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSitesRequest() (request *DescribeSitesRequest) {
	request = &DescribeSitesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSites")
	return
}

func NewDescribeSitesResponse() (response *DescribeSitesResponse) {
	response = &DescribeSitesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询站点列表
func (c *Client) DescribeSites(request *DescribeSitesRequest) (response *DescribeSitesResponse, err error) {
	if request == nil {
		request = NewDescribeSitesRequest()
	}
	response = NewDescribeSitesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLabelRequest() (request *CreateLabelRequest) {
	request = &CreateLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateLabel")
	return
}

func NewCreateLabelResponse() (response *CreateLabelResponse) {
	response = &CreateLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建标签
func (c *Client) CreateLabel(request *CreateLabelRequest) (response *CreateLabelResponse, err error) {
	if request == nil {
		request = NewCreateLabelRequest()
	}
	response = NewCreateLabelResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAgencyConfigRequest() (request *CreateAgencyConfigRequest) {
	request = &CreateAgencyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateAgencyConfig")
	return
}

func NewCreateAgencyConfigResponse() (response *CreateAgencyConfigResponse) {
	response = &CreateAgencyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建机构流程设置
func (c *Client) CreateAgencyConfig(request *CreateAgencyConfigRequest) (response *CreateAgencyConfigResponse, err error) {
	if request == nil {
		request = NewCreateAgencyConfigRequest()
	}
	response = NewCreateAgencyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewRejectedSubscriptionRequest() (request *RejectedSubscriptionRequest) {
	request = &RejectedSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RejectedSubscription")
	return
}

func NewRejectedSubscriptionResponse() (response *RejectedSubscriptionResponse) {
	response = &RejectedSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拒绝订阅
func (c *Client) RejectedSubscription(request *RejectedSubscriptionRequest) (response *RejectedSubscriptionResponse, err error) {
	if request == nil {
		request = NewRejectedSubscriptionRequest()
	}
	response = NewRejectedSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyPhoneRequest() (request *VerifyPhoneRequest) {
	request = &VerifyPhoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "VerifyPhone")
	return
}

func NewVerifyPhoneResponse() (response *VerifyPhoneResponse) {
	response = &VerifyPhoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证手机号
func (c *Client) VerifyPhone(request *VerifyPhoneRequest) (response *VerifyPhoneResponse, err error) {
	if request == nil {
		request = NewVerifyPhoneRequest()
	}
	response = NewVerifyPhoneResponse()
	err = c.Send(request, response)
	return
}

func NewAddMemberToAcsRoleRequest() (request *AddMemberToAcsRoleRequest) {
	request = &AddMemberToAcsRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "AddMemberToAcsRole")
	return
}

func NewAddMemberToAcsRoleResponse() (response *AddMemberToAcsRoleResponse) {
	response = &AddMemberToAcsRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 角色添加成员
func (c *Client) AddMemberToAcsRole(request *AddMemberToAcsRoleRequest) (response *AddMemberToAcsRoleResponse, err error) {
	if request == nil {
		request = NewAddMemberToAcsRoleRequest()
	}
	response = NewAddMemberToAcsRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountPwdRequest() (request *DescribeAccountPwdRequest) {
	request = &DescribeAccountPwdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAccountPwd")
	return
}

func NewDescribeAccountPwdResponse() (response *DescribeAccountPwdResponse) {
	response = &DescribeAccountPwdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 找回密码
func (c *Client) DescribeAccountPwd(request *DescribeAccountPwdRequest) (response *DescribeAccountPwdResponse, err error) {
	if request == nil {
		request = NewDescribeAccountPwdRequest()
	}
	response = NewDescribeAccountPwdResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMicroRegisterCenterRequest() (request *CreateMicroRegisterCenterRequest) {
	request = &CreateMicroRegisterCenterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateMicroRegisterCenter")
	return
}

func NewCreateMicroRegisterCenterResponse() (response *CreateMicroRegisterCenterResponse) {
	response = &CreateMicroRegisterCenterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务注册中心创建
func (c *Client) CreateMicroRegisterCenter(request *CreateMicroRegisterCenterRequest) (response *CreateMicroRegisterCenterResponse, err error) {
	if request == nil {
		request = NewCreateMicroRegisterCenterRequest()
	}
	response = NewCreateMicroRegisterCenterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTSFGatewayApiGroupRequest() (request *CreateTSFGatewayApiGroupRequest) {
	request = &CreateTSFGatewayApiGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateTSFGatewayApiGroup")
	return
}

func NewCreateTSFGatewayApiGroupResponse() (response *CreateTSFGatewayApiGroupResponse) {
	response = &CreateTSFGatewayApiGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步租户下 TSF 网关 API 分组
func (c *Client) CreateTSFGatewayApiGroup(request *CreateTSFGatewayApiGroupRequest) (response *CreateTSFGatewayApiGroupResponse, err error) {
	if request == nil {
		request = NewCreateTSFGatewayApiGroupRequest()
	}
	response = NewCreateTSFGatewayApiGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceApproveLogsRequest() (request *DescribeServiceApproveLogsRequest) {
	request = &DescribeServiceApproveLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeServiceApproveLogs")
	return
}

func NewDescribeServiceApproveLogsResponse() (response *DescribeServiceApproveLogsResponse) {
	response = &DescribeServiceApproveLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务审批历史记录
func (c *Client) DescribeServiceApproveLogs(request *DescribeServiceApproveLogsRequest) (response *DescribeServiceApproveLogsResponse, err error) {
	if request == nil {
		request = NewDescribeServiceApproveLogsRequest()
	}
	response = NewDescribeServiceApproveLogsResponse()
	err = c.Send(request, response)
	return
}

func NewApprovedServiceBatchRequest() (request *ApprovedServiceBatchRequest) {
	request = &ApprovedServiceBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ApprovedServiceBatch")
	return
}

func NewApprovedServiceBatchResponse() (response *ApprovedServiceBatchResponse) {
	response = &ApprovedServiceBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务审批同意批量
func (c *Client) ApprovedServiceBatch(request *ApprovedServiceBatchRequest) (response *ApprovedServiceBatchResponse, err error) {
	if request == nil {
		request = NewApprovedServiceBatchRequest()
	}
	response = NewApprovedServiceBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroProtoGroupRequest() (request *DescribeMicroProtoGroupRequest) {
	request = &DescribeMicroProtoGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroProtoGroup")
	return
}

func NewDescribeMicroProtoGroupResponse() (response *DescribeMicroProtoGroupResponse) {
	response = &DescribeMicroProtoGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件组查询
func (c *Client) DescribeMicroProtoGroup(request *DescribeMicroProtoGroupRequest) (response *DescribeMicroProtoGroupResponse, err error) {
	if request == nil {
		request = NewDescribeMicroProtoGroupRequest()
	}
	response = NewDescribeMicroProtoGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgencyListRequest() (request *DescribeAgencyListRequest) {
	request = &DescribeAgencyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAgencyList")
	return
}

func NewDescribeAgencyListResponse() (response *DescribeAgencyListResponse) {
	response = &DescribeAgencyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机构列表
func (c *Client) DescribeAgencyList(request *DescribeAgencyListRequest) (response *DescribeAgencyListResponse, err error) {
	if request == nil {
		request = NewDescribeAgencyListRequest()
	}
	response = NewDescribeAgencyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataSourceRequest() (request *DescribeDataSourceRequest) {
	request = &DescribeDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeDataSource")
	return
}

func NewDescribeDataSourceResponse() (response *DescribeDataSourceResponse) {
	response = &DescribeDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据源查询
func (c *Client) DescribeDataSource(request *DescribeDataSourceRequest) (response *DescribeDataSourceResponse, err error) {
	if request == nil {
		request = NewDescribeDataSourceRequest()
	}
	response = NewDescribeDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceArchRequest() (request *DescribeServiceArchRequest) {
	request = &DescribeServiceArchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeServiceArch")
	return
}

func NewDescribeServiceArchResponse() (response *DescribeServiceArchResponse) {
	response = &DescribeServiceArchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询API服务详情信息（已归档的处理）
func (c *Client) DescribeServiceArch(request *DescribeServiceArchRequest) (response *DescribeServiceArchResponse, err error) {
	if request == nil {
		request = NewDescribeServiceArchRequest()
	}
	response = NewDescribeServiceArchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemsRequest() (request *DescribeSystemsRequest) {
	request = &DescribeSystemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSystems")
	return
}

func NewDescribeSystemsResponse() (response *DescribeSystemsResponse) {
	response = &DescribeSystemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询系统列表接口
func (c *Client) DescribeSystems(request *DescribeSystemsRequest) (response *DescribeSystemsResponse, err error) {
	if request == nil {
		request = NewDescribeSystemsRequest()
	}
	response = NewDescribeSystemsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDataSourceRequest() (request *DeleteDataSourceRequest) {
	request = &DeleteDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteDataSource")
	return
}

func NewDeleteDataSourceResponse() (response *DeleteDataSourceResponse) {
	response = &DeleteDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据源删除
func (c *Client) DeleteDataSource(request *DeleteDataSourceRequest) (response *DeleteDataSourceResponse, err error) {
	if request == nil {
		request = NewDeleteDataSourceRequest()
	}
	response = NewDeleteDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifySiteRequest() (request *ModifySiteRequest) {
	request = &ModifySiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifySite")
	return
}

func NewModifySiteResponse() (response *ModifySiteResponse) {
	response = &ModifySiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新站点
func (c *Client) ModifySite(request *ModifySiteRequest) (response *ModifySiteResponse, err error) {
	if request == nil {
		request = NewModifySiteRequest()
	}
	response = NewModifySiteResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMenuRequest() (request *DeleteMenuRequest) {
	request = &DeleteMenuRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMenu")
	return
}

func NewDeleteMenuResponse() (response *DeleteMenuResponse) {
	response = &DeleteMenuResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除菜单
func (c *Client) DeleteMenu(request *DeleteMenuRequest) (response *DeleteMenuResponse, err error) {
	if request == nil {
		request = NewDeleteMenuRequest()
	}
	response = NewDeleteMenuResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
	request = &DeleteServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteService")
	return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
	response = &DeleteServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除API服务
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
	if request == nil {
		request = NewDeleteServiceRequest()
	}
	response = NewDeleteServiceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDownloadLoginLogRequest() (request *ModifyDownloadLoginLogRequest) {
	request = &ModifyDownloadLoginLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyDownloadLoginLog")
	return
}

func NewModifyDownloadLoginLogResponse() (response *ModifyDownloadLoginLogResponse) {
	response = &ModifyDownloadLoginLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出登录日志
func (c *Client) ModifyDownloadLoginLog(request *ModifyDownloadLoginLogRequest) (response *ModifyDownloadLoginLogResponse, err error) {
	if request == nil {
		request = NewModifyDownloadLoginLogRequest()
	}
	response = NewModifyDownloadLoginLogResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccountAndProfileRequest() (request *DeleteAccountAndProfileRequest) {
	request = &DeleteAccountAndProfileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteAccountAndProfile")
	return
}

func NewDeleteAccountAndProfileResponse() (response *DeleteAccountAndProfileResponse) {
	response = &DeleteAccountAndProfileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除帐号
func (c *Client) DeleteAccountAndProfile(request *DeleteAccountAndProfileRequest) (response *DeleteAccountAndProfileResponse, err error) {
	if request == nil {
		request = NewDeleteAccountAndProfileRequest()
	}
	response = NewDeleteAccountAndProfileResponse()
	err = c.Send(request, response)
	return
}

func NewSyncSessionSecretRequest() (request *SyncSessionSecretRequest) {
	request = &SyncSessionSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "SyncSessionSecret")
	return
}

func NewSyncSessionSecretResponse() (response *SyncSessionSecretResponse) {
	response = &SyncSessionSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步会话秘钥
func (c *Client) SyncSessionSecret(request *SyncSessionSecretRequest) (response *SyncSessionSecretResponse, err error) {
	if request == nil {
		request = NewSyncSessionSecretRequest()
	}
	response = NewSyncSessionSecretResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAlertRuleRequest() (request *CreateAlertRuleRequest) {
	request = &CreateAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateAlertRule")
	return
}

func NewCreateAlertRuleResponse() (response *CreateAlertRuleResponse) {
	response = &CreateAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增告警规则
func (c *Client) CreateAlertRule(request *CreateAlertRuleRequest) (response *CreateAlertRuleResponse, err error) {
	if request == nil {
		request = NewCreateAlertRuleRequest()
	}
	response = NewCreateAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionApprovesRequest() (request *DescribeSubscriptionApprovesRequest) {
	request = &DescribeSubscriptionApprovesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSubscriptionApproves")
	return
}

func NewDescribeSubscriptionApprovesResponse() (response *DescribeSubscriptionApprovesResponse) {
	response = &DescribeSubscriptionApprovesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订阅审批列表
func (c *Client) DescribeSubscriptionApproves(request *DescribeSubscriptionApprovesRequest) (response *DescribeSubscriptionApprovesResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionApprovesRequest()
	}
	response = NewDescribeSubscriptionApprovesResponse()
	err = c.Send(request, response)
	return
}

func NewChangeStatusAppRequest() (request *ChangeStatusAppRequest) {
	request = &ChangeStatusAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ChangeStatusApp")
	return
}

func NewChangeStatusAppResponse() (response *ChangeStatusAppResponse) {
	response = &ChangeStatusAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用、启用应用
func (c *Client) ChangeStatusApp(request *ChangeStatusAppRequest) (response *ChangeStatusAppResponse, err error) {
	if request == nil {
		request = NewChangeStatusAppRequest()
	}
	response = NewChangeStatusAppResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSystemRequest() (request *DeleteSystemRequest) {
	request = &DeleteSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteSystem")
	return
}

func NewDeleteSystemResponse() (response *DeleteSystemResponse) {
	response = &DeleteSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除系统
func (c *Client) DeleteSystem(request *DeleteSystemRequest) (response *DeleteSystemResponse, err error) {
	if request == nil {
		request = NewDeleteSystemRequest()
	}
	response = NewDeleteSystemResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthNodeOverviewRequest() (request *DescribeAuthNodeOverviewRequest) {
	request = &DescribeAuthNodeOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAuthNodeOverview")
	return
}

func NewDescribeAuthNodeOverviewResponse() (response *DescribeAuthNodeOverviewResponse) {
	response = &DescribeAuthNodeOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询授权节点概览
func (c *Client) DescribeAuthNodeOverview(request *DescribeAuthNodeOverviewRequest) (response *DescribeAuthNodeOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeAuthNodeOverviewRequest()
	}
	response = NewDescribeAuthNodeOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationLogRequest() (request *DescribeOperationLogRequest) {
	request = &DescribeOperationLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeOperationLog")
	return
}

func NewDescribeOperationLogResponse() (response *DescribeOperationLogResponse) {
	response = &DescribeOperationLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作日志
func (c *Client) DescribeOperationLog(request *DescribeOperationLogRequest) (response *DescribeOperationLogResponse, err error) {
	if request == nil {
		request = NewDescribeOperationLogRequest()
	}
	response = NewDescribeOperationLogResponse()
	err = c.Send(request, response)
	return
}

func NewBatchDeleteLabelRequest() (request *BatchDeleteLabelRequest) {
	request = &BatchDeleteLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "BatchDeleteLabel")
	return
}

func NewBatchDeleteLabelResponse() (response *BatchDeleteLabelResponse) {
	response = &BatchDeleteLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除标签
func (c *Client) BatchDeleteLabel(request *BatchDeleteLabelRequest) (response *BatchDeleteLabelResponse, err error) {
	if request == nil {
		request = NewBatchDeleteLabelRequest()
	}
	response = NewBatchDeleteLabelResponse()
	err = c.Send(request, response)
	return
}

func NewExportAgencyMemberRequest() (request *ExportAgencyMemberRequest) {
	request = &ExportAgencyMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ExportAgencyMember")
	return
}

func NewExportAgencyMemberResponse() (response *ExportAgencyMemberResponse) {
	response = &ExportAgencyMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出机构成员
func (c *Client) ExportAgencyMember(request *ExportAgencyMemberRequest) (response *ExportAgencyMemberResponse, err error) {
	if request == nil {
		request = NewExportAgencyMemberRequest()
	}
	response = NewExportAgencyMemberResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSiteRequest() (request *CreateSiteRequest) {
	request = &CreateSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateSite")
	return
}

func NewCreateSiteResponse() (response *CreateSiteResponse) {
	response = &CreateSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建站点
func (c *Client) CreateSite(request *CreateSiteRequest) (response *CreateSiteResponse, err error) {
	if request == nil {
		request = NewCreateSiteRequest()
	}
	response = NewCreateSiteResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSmartGatewayGroupRequest() (request *CreateSmartGatewayGroupRequest) {
	request = &CreateSmartGatewayGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateSmartGatewayGroup")
	return
}

func NewCreateSmartGatewayGroupResponse() (response *CreateSmartGatewayGroupResponse) {
	response = &CreateSmartGatewayGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建节点阵列
func (c *Client) CreateSmartGatewayGroup(request *CreateSmartGatewayGroupRequest) (response *CreateSmartGatewayGroupResponse, err error) {
	if request == nil {
		request = NewCreateSmartGatewayGroupRequest()
	}
	response = NewCreateSmartGatewayGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTSFGatewayServicesRequest() (request *DescribeTSFGatewayServicesRequest) {
	request = &DescribeTSFGatewayServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTSFGatewayServices")
	return
}

func NewDescribeTSFGatewayServicesResponse() (response *DescribeTSFGatewayServicesResponse) {
	response = &DescribeTSFGatewayServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF 网关微服务 API 列表
func (c *Client) DescribeTSFGatewayServices(request *DescribeTSFGatewayServicesRequest) (response *DescribeTSFGatewayServicesResponse, err error) {
	if request == nil {
		request = NewDescribeTSFGatewayServicesRequest()
	}
	response = NewDescribeTSFGatewayServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSubscriptionRequest() (request *DeleteSubscriptionRequest) {
	request = &DeleteSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteSubscription")
	return
}

func NewDeleteSubscriptionResponse() (response *DeleteSubscriptionResponse) {
	response = &DeleteSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除订阅
func (c *Client) DeleteSubscription(request *DeleteSubscriptionRequest) (response *DeleteSubscriptionResponse, err error) {
	if request == nil {
		request = NewDeleteSubscriptionRequest()
	}
	response = NewDeleteSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityConfigRequest() (request *DescribeSecurityConfigRequest) {
	request = &DescribeSecurityConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSecurityConfig")
	return
}

func NewDescribeSecurityConfigResponse() (response *DescribeSecurityConfigResponse) {
	response = &DescribeSecurityConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全设置
func (c *Client) DescribeSecurityConfig(request *DescribeSecurityConfigRequest) (response *DescribeSecurityConfigResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityConfigRequest()
	}
	response = NewDescribeSecurityConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadWebConfigRequest() (request *DownloadWebConfigRequest) {
	request = &DownloadWebConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadWebConfig")
	return
}

func NewDownloadWebConfigResponse() (response *DownloadWebConfigResponse) {
	response = &DownloadWebConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载前端WebConfig
func (c *Client) DownloadWebConfig(request *DownloadWebConfigRequest) (response *DownloadWebConfigResponse, err error) {
	if request == nil {
		request = NewDownloadWebConfigRequest()
	}
	response = NewDownloadWebConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroProtoGroupsRequest() (request *DescribeMicroProtoGroupsRequest) {
	request = &DescribeMicroProtoGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroProtoGroups")
	return
}

func NewDescribeMicroProtoGroupsResponse() (response *DescribeMicroProtoGroupsResponse) {
	response = &DescribeMicroProtoGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件组列表
func (c *Client) DescribeMicroProtoGroups(request *DescribeMicroProtoGroupsRequest) (response *DescribeMicroProtoGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeMicroProtoGroupsRequest()
	}
	response = NewDescribeMicroProtoGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSystemWarnMsgRequest() (request *DeleteSystemWarnMsgRequest) {
	request = &DeleteSystemWarnMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteSystemWarnMsg")
	return
}

func NewDeleteSystemWarnMsgResponse() (response *DeleteSystemWarnMsgResponse) {
	response = &DeleteSystemWarnMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除时提醒内部应用数量
func (c *Client) DeleteSystemWarnMsg(request *DeleteSystemWarnMsgRequest) (response *DeleteSystemWarnMsgResponse, err error) {
	if request == nil {
		request = NewDeleteSystemWarnMsgRequest()
	}
	response = NewDeleteSystemWarnMsgResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSubscriptionRequest() (request *UpdateSubscriptionRequest) {
	request = &UpdateSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UpdateSubscription")
	return
}

func NewUpdateSubscriptionResponse() (response *UpdateSubscriptionResponse) {
	response = &UpdateSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新订阅
func (c *Client) UpdateSubscription(request *UpdateSubscriptionRequest) (response *UpdateSubscriptionResponse, err error) {
	if request == nil {
		request = NewUpdateSubscriptionRequest()
	}
	response = NewUpdateSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCategoryRequest() (request *DeleteCategoryRequest) {
	request = &DeleteCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteCategory")
	return
}

func NewDeleteCategoryResponse() (response *DeleteCategoryResponse) {
	response = &DeleteCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除⽬录
func (c *Client) DeleteCategory(request *DeleteCategoryRequest) (response *DeleteCategoryResponse, err error) {
	if request == nil {
		request = NewDeleteCategoryRequest()
	}
	response = NewDeleteCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLoginPolicyRequest() (request *ModifyLoginPolicyRequest) {
	request = &ModifyLoginPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyLoginPolicy")
	return
}

func NewModifyLoginPolicyResponse() (response *ModifyLoginPolicyResponse) {
	response = &ModifyLoginPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 登录安全策略修改
func (c *Client) ModifyLoginPolicy(request *ModifyLoginPolicyRequest) (response *ModifyLoginPolicyResponse, err error) {
	if request == nil {
		request = NewModifyLoginPolicyRequest()
	}
	response = NewModifyLoginPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAcsRoleListRequest() (request *DescribeAcsRoleListRequest) {
	request = &DescribeAcsRoleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAcsRoleList")
	return
}

func NewDescribeAcsRoleListResponse() (response *DescribeAcsRoleListResponse) {
	response = &DescribeAcsRoleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取角色列表
func (c *Client) DescribeAcsRoleList(request *DescribeAcsRoleListRequest) (response *DescribeAcsRoleListResponse, err error) {
	if request == nil {
		request = NewDescribeAcsRoleListRequest()
	}
	response = NewDescribeAcsRoleListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRestoreSnapshotRequest() (request *ModifyRestoreSnapshotRequest) {
	request = &ModifyRestoreSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyRestoreSnapshot")
	return
}

func NewModifyRestoreSnapshotResponse() (response *ModifyRestoreSnapshotResponse) {
	response = &ModifyRestoreSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恢复存档
func (c *Client) ModifyRestoreSnapshot(request *ModifyRestoreSnapshotRequest) (response *ModifyRestoreSnapshotResponse, err error) {
	if request == nil {
		request = NewModifyRestoreSnapshotRequest()
	}
	response = NewModifyRestoreSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountAndProfileListRequest() (request *DescribeAccountAndProfileListRequest) {
	request = &DescribeAccountAndProfileListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAccountAndProfileList")
	return
}

func NewDescribeAccountAndProfileListResponse() (response *DescribeAccountAndProfileListResponse) {
	response = &DescribeAccountAndProfileListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账户列表
func (c *Client) DescribeAccountAndProfileList(request *DescribeAccountAndProfileListRequest) (response *DescribeAccountAndProfileListResponse, err error) {
	if request == nil {
		request = NewDescribeAccountAndProfileListRequest()
	}
	response = NewDescribeAccountAndProfileListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseInfoRequest() (request *DescribeLicenseInfoRequest) {
	request = &DescribeLicenseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeLicenseInfo")
	return
}

func NewDescribeLicenseInfoResponse() (response *DescribeLicenseInfoResponse) {
	response = &DescribeLicenseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询证书信息
func (c *Client) DescribeLicenseInfo(request *DescribeLicenseInfoRequest) (response *DescribeLicenseInfoResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseInfoRequest()
	}
	response = NewDescribeLicenseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProfileForRoleRequest() (request *DescribeProfileForRoleRequest) {
	request = &DescribeProfileForRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeProfileForRole")
	return
}

func NewDescribeProfileForRoleResponse() (response *DescribeProfileForRoleResponse) {
	response = &DescribeProfileForRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询角色的用户列表，携带是否是成员或管理员
func (c *Client) DescribeProfileForRole(request *DescribeProfileForRoleRequest) (response *DescribeProfileForRoleResponse, err error) {
	if request == nil {
		request = NewDescribeProfileForRoleRequest()
	}
	response = NewDescribeProfileForRoleResponse()
	err = c.Send(request, response)
	return
}

func NewExportServiceRequest() (request *ExportServiceRequest) {
	request = &ExportServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ExportService")
	return
}

func NewExportServiceResponse() (response *ExportServiceResponse) {
	response = &ExportServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// swagger导出API服务
func (c *Client) ExportService(request *ExportServiceRequest) (response *ExportServiceResponse, err error) {
	if request == nil {
		request = NewExportServiceRequest()
	}
	response = NewExportServiceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDownloadOperationLogRequest() (request *ModifyDownloadOperationLogRequest) {
	request = &ModifyDownloadOperationLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyDownloadOperationLog")
	return
}

func NewModifyDownloadOperationLogResponse() (response *ModifyDownloadOperationLogResponse) {
	response = &ModifyDownloadOperationLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出操作日志
func (c *Client) ModifyDownloadOperationLog(request *ModifyDownloadOperationLogRequest) (response *ModifyDownloadOperationLogResponse, err error) {
	if request == nil {
		request = NewModifyDownloadOperationLogRequest()
	}
	response = NewModifyDownloadOperationLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserVendorsRequest() (request *DescribeUserVendorsRequest) {
	request = &DescribeUserVendorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeUserVendors")
	return
}

func NewDescribeUserVendorsResponse() (response *DescribeUserVendorsResponse) {
	response = &DescribeUserVendorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户所属供应商
func (c *Client) DescribeUserVendors(request *DescribeUserVendorsRequest) (response *DescribeUserVendorsResponse, err error) {
	if request == nil {
		request = NewDescribeUserVendorsRequest()
	}
	response = NewDescribeUserVendorsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateZoneRequest() (request *CreateZoneRequest) {
	request = &CreateZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateZone")
	return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
	response = &CreateZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建分区
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
	if request == nil {
		request = NewCreateZoneRequest()
	}
	response = NewCreateZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroProtoFilesRequest() (request *DeleteMicroProtoFilesRequest) {
	request = &DeleteMicroProtoFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMicroProtoFiles")
	return
}

func NewDeleteMicroProtoFilesResponse() (response *DeleteMicroProtoFilesResponse) {
	response = &DeleteMicroProtoFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件批量删除
func (c *Client) DeleteMicroProtoFiles(request *DeleteMicroProtoFilesRequest) (response *DeleteMicroProtoFilesResponse, err error) {
	if request == nil {
		request = NewDeleteMicroProtoFilesRequest()
	}
	response = NewDeleteMicroProtoFilesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSmartGateRaidRoutingRequest() (request *DeleteSmartGateRaidRoutingRequest) {
	request = &DeleteSmartGateRaidRoutingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteSmartGateRaidRouting")
	return
}

func NewDeleteSmartGateRaidRoutingResponse() (response *DeleteSmartGateRaidRoutingResponse) {
	response = &DeleteSmartGateRaidRoutingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除节点阵列路由
func (c *Client) DeleteSmartGateRaidRouting(request *DeleteSmartGateRaidRoutingRequest) (response *DeleteSmartGateRaidRoutingResponse, err error) {
	if request == nil {
		request = NewDeleteSmartGateRaidRoutingRequest()
	}
	response = NewDeleteSmartGateRaidRoutingResponse()
	err = c.Send(request, response)
	return
}

func NewBatchDeleteServiceRequest() (request *BatchDeleteServiceRequest) {
	request = &BatchDeleteServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "BatchDeleteService")
	return
}

func NewBatchDeleteServiceResponse() (response *BatchDeleteServiceResponse) {
	response = &BatchDeleteServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除API服务
func (c *Client) BatchDeleteService(request *BatchDeleteServiceRequest) (response *BatchDeleteServiceResponse, err error) {
	if request == nil {
		request = NewBatchDeleteServiceRequest()
	}
	response = NewBatchDeleteServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDataSourcesRequest() (request *DeleteDataSourcesRequest) {
	request = &DeleteDataSourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteDataSources")
	return
}

func NewDeleteDataSourcesResponse() (response *DeleteDataSourcesResponse) {
	response = &DeleteDataSourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据源批量删除
func (c *Client) DeleteDataSources(request *DeleteDataSourcesRequest) (response *DeleteDataSourcesResponse, err error) {
	if request == nil {
		request = NewDeleteDataSourcesRequest()
	}
	response = NewDeleteDataSourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroProtoGroupsRequest() (request *DeleteMicroProtoGroupsRequest) {
	request = &DeleteMicroProtoGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMicroProtoGroups")
	return
}

func NewDeleteMicroProtoGroupsResponse() (response *DeleteMicroProtoGroupsResponse) {
	response = &DeleteMicroProtoGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件组批量删除
func (c *Client) DeleteMicroProtoGroups(request *DeleteMicroProtoGroupsRequest) (response *DeleteMicroProtoGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteMicroProtoGroupsRequest()
	}
	response = NewDeleteMicroProtoGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataSourceTypesRequest() (request *DescribeDataSourceTypesRequest) {
	request = &DescribeDataSourceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeDataSourceTypes")
	return
}

func NewDescribeDataSourceTypesResponse() (response *DescribeDataSourceTypesResponse) {
	response = &DescribeDataSourceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据源类型列表
func (c *Client) DescribeDataSourceTypes(request *DescribeDataSourceTypesRequest) (response *DescribeDataSourceTypesResponse, err error) {
	if request == nil {
		request = NewDescribeDataSourceTypesRequest()
	}
	response = NewDescribeDataSourceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLabelRequest() (request *ModifyLabelRequest) {
	request = &ModifyLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyLabel")
	return
}

func NewModifyLabelResponse() (response *ModifyLabelResponse) {
	response = &ModifyLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改标签
func (c *Client) ModifyLabel(request *ModifyLabelRequest) (response *ModifyLabelResponse, err error) {
	if request == nil {
		request = NewModifyLabelRequest()
	}
	response = NewModifyLabelResponse()
	err = c.Send(request, response)
	return
}

func NewModifySystemRequest() (request *ModifySystemRequest) {
	request = &ModifySystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifySystem")
	return
}

func NewModifySystemResponse() (response *ModifySystemResponse) {
	response = &ModifySystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改系统
func (c *Client) ModifySystem(request *ModifySystemRequest) (response *ModifySystemResponse, err error) {
	if request == nil {
		request = NewModifySystemRequest()
	}
	response = NewModifySystemResponse()
	err = c.Send(request, response)
	return
}

func NewCheckVendorRequest() (request *CheckVendorRequest) {
	request = &CheckVendorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CheckVendor")
	return
}

func NewCheckVendorResponse() (response *CheckVendorResponse) {
	response = &CheckVendorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验供应商
func (c *Client) CheckVendor(request *CheckVendorRequest) (response *CheckVendorResponse, err error) {
	if request == nil {
		request = NewCheckVendorRequest()
	}
	response = NewCheckVendorResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMicroProtoFileRequest() (request *CreateMicroProtoFileRequest) {
	request = &CreateMicroProtoFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateMicroProtoFile")
	return
}

func NewCreateMicroProtoFileResponse() (response *CreateMicroProtoFileResponse) {
	response = &CreateMicroProtoFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件创建
func (c *Client) CreateMicroProtoFile(request *CreateMicroProtoFileRequest) (response *CreateMicroProtoFileResponse, err error) {
	if request == nil {
		request = NewCreateMicroProtoFileRequest()
	}
	response = NewCreateMicroProtoFileResponse()
	err = c.Send(request, response)
	return
}

func NewRenewSubscriptionRequest() (request *RenewSubscriptionRequest) {
	request = &RenewSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RenewSubscription")
	return
}

func NewRenewSubscriptionResponse() (response *RenewSubscriptionResponse) {
	response = &RenewSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续期订阅
func (c *Client) RenewSubscription(request *RenewSubscriptionRequest) (response *RenewSubscriptionResponse, err error) {
	if request == nil {
		request = NewRenewSubscriptionRequest()
	}
	response = NewRenewSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDataSourceRequest() (request *CreateDataSourceRequest) {
	request = &CreateDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateDataSource")
	return
}

func NewCreateDataSourceResponse() (response *CreateDataSourceResponse) {
	response = &CreateDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据源创建
func (c *Client) CreateDataSource(request *CreateDataSourceRequest) (response *CreateDataSourceResponse, err error) {
	if request == nil {
		request = NewCreateDataSourceRequest()
	}
	response = NewCreateDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConnectorRequest() (request *DeleteConnectorRequest) {
	request = &DeleteConnectorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteConnector")
	return
}

func NewDeleteConnectorResponse() (response *DeleteConnectorResponse) {
	response = &DeleteConnectorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 认证删除
func (c *Client) DeleteConnector(request *DeleteConnectorRequest) (response *DeleteConnectorResponse, err error) {
	if request == nil {
		request = NewDeleteConnectorRequest()
	}
	response = NewDeleteConnectorResponse()
	err = c.Send(request, response)
	return
}

func NewModifySmartGatewayGroupRequest() (request *ModifySmartGatewayGroupRequest) {
	request = &ModifySmartGatewayGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifySmartGatewayGroup")
	return
}

func NewModifySmartGatewayGroupResponse() (response *ModifySmartGatewayGroupResponse) {
	response = &ModifySmartGatewayGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改节点阵列
func (c *Client) ModifySmartGatewayGroup(request *ModifySmartGatewayGroupRequest) (response *ModifySmartGatewayGroupResponse, err error) {
	if request == nil {
		request = NewModifySmartGatewayGroupRequest()
	}
	response = NewModifySmartGatewayGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLabelRequest() (request *DeleteLabelRequest) {
	request = &DeleteLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteLabel")
	return
}

func NewDeleteLabelResponse() (response *DeleteLabelResponse) {
	response = &DeleteLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除标签
func (c *Client) DeleteLabel(request *DeleteLabelRequest) (response *DeleteLabelResponse, err error) {
	if request == nil {
		request = NewDeleteLabelRequest()
	}
	response = NewDeleteLabelResponse()
	err = c.Send(request, response)
	return
}

func NewExportVendorMemberRequest() (request *ExportVendorMemberRequest) {
	request = &ExportVendorMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ExportVendorMember")
	return
}

func NewExportVendorMemberResponse() (response *ExportVendorMemberResponse) {
	response = &ExportVendorMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供应商成员导出
func (c *Client) ExportVendorMember(request *ExportVendorMemberRequest) (response *ExportVendorMemberResponse, err error) {
	if request == nil {
		request = NewExportVendorMemberRequest()
	}
	response = NewExportVendorMemberResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplyListRequest() (request *DescribeApplyListRequest) {
	request = &DescribeApplyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeApplyList")
	return
}

func NewDescribeApplyListResponse() (response *DescribeApplyListResponse) {
	response = &DescribeApplyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询已申请证书信息
func (c *Client) DescribeApplyList(request *DescribeApplyListRequest) (response *DescribeApplyListResponse, err error) {
	if request == nil {
		request = NewDescribeApplyListRequest()
	}
	response = NewDescribeApplyListResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadResultRequest() (request *DownloadResultRequest) {
	request = &DownloadResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadResult")
	return
}

func NewDownloadResultResponse() (response *DownloadResultResponse) {
	response = &DownloadResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载批量导入用户结果文件
func (c *Client) DownloadResult(request *DownloadResultRequest) (response *DownloadResultResponse, err error) {
	if request == nil {
		request = NewDownloadResultRequest()
	}
	response = NewDownloadResultResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTSFServiceRequest() (request *CreateTSFServiceRequest) {
	request = &CreateTSFServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateTSFService")
	return
}

func NewCreateTSFServiceResponse() (response *CreateTSFServiceResponse) {
	response = &CreateTSFServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步租户下 TSF 指定 API
func (c *Client) CreateTSFService(request *CreateTSFServiceRequest) (response *CreateTSFServiceResponse, err error) {
	if request == nil {
		request = NewCreateTSFServiceRequest()
	}
	response = NewCreateTSFServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlertHistoryFileRequest() (request *DescribeAlertHistoryFileRequest) {
	request = &DescribeAlertHistoryFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAlertHistoryFile")
	return
}

func NewDescribeAlertHistoryFileResponse() (response *DescribeAlertHistoryFileResponse) {
	response = &DescribeAlertHistoryFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出告警历史
func (c *Client) DescribeAlertHistoryFile(request *DescribeAlertHistoryFileRequest) (response *DescribeAlertHistoryFileResponse, err error) {
	if request == nil {
		request = NewDescribeAlertHistoryFileRequest()
	}
	response = NewDescribeAlertHistoryFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSiteArchRequest() (request *DescribeSiteArchRequest) {
	request = &DescribeSiteArchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSiteArch")
	return
}

func NewDescribeSiteArchResponse() (response *DescribeSiteArchResponse) {
	response = &DescribeSiteArchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询站点归档详情
func (c *Client) DescribeSiteArch(request *DescribeSiteArchRequest) (response *DescribeSiteArchResponse, err error) {
	if request == nil {
		request = NewDescribeSiteArchRequest()
	}
	response = NewDescribeSiteArchResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTlsCertificateRequest() (request *DeleteTlsCertificateRequest) {
	request = &DeleteTlsCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteTlsCertificate")
	return
}

func NewDeleteTlsCertificateResponse() (response *DeleteTlsCertificateResponse) {
	response = &DeleteTlsCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除证书
func (c *Client) DeleteTlsCertificate(request *DeleteTlsCertificateRequest) (response *DeleteTlsCertificateResponse, err error) {
	if request == nil {
		request = NewDeleteTlsCertificateRequest()
	}
	response = NewDeleteTlsCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogoRequest() (request *DescribeLogoRequest) {
	request = &DescribeLogoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeLogo")
	return
}

func NewDescribeLogoResponse() (response *DescribeLogoResponse) {
	response = &DescribeLogoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询logo
func (c *Client) DescribeLogo(request *DescribeLogoRequest) (response *DescribeLogoResponse, err error) {
	if request == nil {
		request = NewDescribeLogoRequest()
	}
	response = NewDescribeLogoResponse()
	err = c.Send(request, response)
	return
}

func NewBatchDeleteSiteRequest() (request *BatchDeleteSiteRequest) {
	request = &BatchDeleteSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "BatchDeleteSite")
	return
}

func NewBatchDeleteSiteResponse() (response *BatchDeleteSiteResponse) {
	response = &BatchDeleteSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除站点
func (c *Client) BatchDeleteSite(request *BatchDeleteSiteRequest) (response *BatchDeleteSiteResponse, err error) {
	if request == nil {
		request = NewBatchDeleteSiteRequest()
	}
	response = NewBatchDeleteSiteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSiteRequestLogPageRequest() (request *DescribeSiteRequestLogPageRequest) {
	request = &DescribeSiteRequestLogPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSiteRequestLogPage")
	return
}

func NewDescribeSiteRequestLogPageResponse() (response *DescribeSiteRequestLogPageResponse) {
	response = &DescribeSiteRequestLogPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询站点请求日志列表
func (c *Client) DescribeSiteRequestLogPage(request *DescribeSiteRequestLogPageRequest) (response *DescribeSiteRequestLogPageResponse, err error) {
	if request == nil {
		request = NewDescribeSiteRequestLogPageRequest()
	}
	response = NewDescribeSiteRequestLogPageResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateVendorMemberRequest() (request *UpdateVendorMemberRequest) {
	request = &UpdateVendorMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UpdateVendorMember")
	return
}

func NewUpdateVendorMemberResponse() (response *UpdateVendorMemberResponse) {
	response = &UpdateVendorMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑供应商成员
func (c *Client) UpdateVendorMember(request *UpdateVendorMemberRequest) (response *UpdateVendorMemberResponse, err error) {
	if request == nil {
		request = NewUpdateVendorMemberRequest()
	}
	response = NewUpdateVendorMemberResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMenuAllListRequest() (request *DescribeMenuAllListRequest) {
	request = &DescribeMenuAllListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMenuAllList")
	return
}

func NewDescribeMenuAllListResponse() (response *DescribeMenuAllListResponse) {
	response = &DescribeMenuAllListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询菜单
func (c *Client) DescribeMenuAllList(request *DescribeMenuAllListRequest) (response *DescribeMenuAllListResponse, err error) {
	if request == nil {
		request = NewDescribeMenuAllListRequest()
	}
	response = NewDescribeMenuAllListResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadVendorImportResultRequest() (request *DownloadVendorImportResultRequest) {
	request = &DownloadVendorImportResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadVendorImportResult")
	return
}

func NewDownloadVendorImportResultResponse() (response *DownloadVendorImportResultResponse) {
	response = &DownloadVendorImportResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载导入失败结果文件
func (c *Client) DownloadVendorImportResult(request *DownloadVendorImportResultRequest) (response *DownloadVendorImportResultResponse, err error) {
	if request == nil {
		request = NewDownloadVendorImportResultRequest()
	}
	response = NewDownloadVendorImportResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSiteRequest() (request *DescribeSiteRequest) {
	request = &DescribeSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSite")
	return
}

func NewDescribeSiteResponse() (response *DescribeSiteResponse) {
	response = &DescribeSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询站点详情
func (c *Client) DescribeSite(request *DescribeSiteRequest) (response *DescribeSiteResponse, err error) {
	if request == nil {
		request = NewDescribeSiteRequest()
	}
	response = NewDescribeSiteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSqlTemplatesRequest() (request *DescribeSqlTemplatesRequest) {
	request = &DescribeSqlTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSqlTemplates")
	return
}

func NewDescribeSqlTemplatesResponse() (response *DescribeSqlTemplatesResponse) {
	response = &DescribeSqlTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SQL 模板列表
func (c *Client) DescribeSqlTemplates(request *DescribeSqlTemplatesRequest) (response *DescribeSqlTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeSqlTemplatesRequest()
	}
	response = NewDescribeSqlTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewListProfileRequest() (request *ListProfileRequest) {
	request = &ListProfileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ListProfile")
	return
}

func NewListProfileResponse() (response *ListProfileResponse) {
	response = &ListProfileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机构成员用户信息
func (c *Client) ListProfile(request *ListProfileRequest) (response *ListProfileResponse, err error) {
	if request == nil {
		request = NewListProfileRequest()
	}
	response = NewListProfileResponse()
	err = c.Send(request, response)
	return
}

func NewSendSmsCodeRequest() (request *SendSmsCodeRequest) {
	request = &SendSmsCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "SendSmsCode")
	return
}

func NewSendSmsCodeResponse() (response *SendSmsCodeResponse) {
	response = &SendSmsCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送短信
func (c *Client) SendSmsCode(request *SendSmsCodeRequest) (response *SendSmsCodeResponse, err error) {
	if request == nil {
		request = NewSendSmsCodeRequest()
	}
	response = NewSendSmsCodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountInfoRequest() (request *DescribeAccountInfoRequest) {
	request = &DescribeAccountInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAccountInfo")
	return
}

func NewDescribeAccountInfoResponse() (response *DescribeAccountInfoResponse) {
	response = &DescribeAccountInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询帐号信息
func (c *Client) DescribeAccountInfo(request *DescribeAccountInfoRequest) (response *DescribeAccountInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAccountInfoRequest()
	}
	response = NewDescribeAccountInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServiceRequest() (request *ModifyServiceRequest) {
	request = &ModifyServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyService")
	return
}

func NewModifyServiceResponse() (response *ModifyServiceResponse) {
	response = &ModifyServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新API服务
func (c *Client) ModifyService(request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
	if request == nil {
		request = NewModifyServiceRequest()
	}
	response = NewModifyServiceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAreaRequest() (request *CreateAreaRequest) {
	request = &CreateAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateArea")
	return
}

func NewCreateAreaResponse() (response *CreateAreaResponse) {
	response = &CreateAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建区域
func (c *Client) CreateArea(request *CreateAreaRequest) (response *CreateAreaResponse, err error) {
	if request == nil {
		request = NewCreateAreaRequest()
	}
	response = NewCreateAreaResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteMicroSyncJobRequest() (request *ExecuteMicroSyncJobRequest) {
	request = &ExecuteMicroSyncJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ExecuteMicroSyncJob")
	return
}

func NewExecuteMicroSyncJobResponse() (response *ExecuteMicroSyncJobResponse) {
	response = &ExecuteMicroSyncJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 主动发起微服务同步
func (c *Client) ExecuteMicroSyncJob(request *ExecuteMicroSyncJobRequest) (response *ExecuteMicroSyncJobResponse, err error) {
	if request == nil {
		request = NewExecuteMicroSyncJobRequest()
	}
	response = NewExecuteMicroSyncJobResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadAccTemplateRequest() (request *DownloadAccTemplateRequest) {
	request = &DownloadAccTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadAccTemplate")
	return
}

func NewDownloadAccTemplateResponse() (response *DownloadAccTemplateResponse) {
	response = &DownloadAccTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载批量导入用户文件模板
func (c *Client) DownloadAccTemplate(request *DownloadAccTemplateRequest) (response *DownloadAccTemplateResponse, err error) {
	if request == nil {
		request = NewDownloadAccTemplateRequest()
	}
	response = NewDownloadAccTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppsRequest() (request *DescribeAppsRequest) {
	request = &DescribeAppsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeApps")
	return
}

func NewDescribeAppsResponse() (response *DescribeAppsResponse) {
	response = &DescribeAppsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询应用列表
func (c *Client) DescribeApps(request *DescribeAppsRequest) (response *DescribeAppsResponse, err error) {
	if request == nil {
		request = NewDescribeAppsRequest()
	}
	response = NewDescribeAppsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemAndAppRequest() (request *DescribeSystemAndAppRequest) {
	request = &DescribeSystemAndAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSystemAndApp")
	return
}

func NewDescribeSystemAndAppResponse() (response *DescribeSystemAndAppResponse) {
	response = &DescribeSystemAndAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeSystemAndApp
func (c *Client) DescribeSystemAndApp(request *DescribeSystemAndAppRequest) (response *DescribeSystemAndAppResponse, err error) {
	if request == nil {
		request = NewDescribeSystemAndAppRequest()
	}
	response = NewDescribeSystemAndAppResponse()
	err = c.Send(request, response)
	return
}

func NewAddVendorMemberRequest() (request *AddVendorMemberRequest) {
	request = &AddVendorMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "AddVendorMember")
	return
}

func NewAddVendorMemberResponse() (response *AddVendorMemberResponse) {
	response = &AddVendorMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加单个供应商成员
func (c *Client) AddVendorMember(request *AddVendorMemberRequest) (response *AddVendorMemberResponse, err error) {
	if request == nil {
		request = NewAddVendorMemberRequest()
	}
	response = NewAddVendorMemberResponse()
	err = c.Send(request, response)
	return
}

func NewApprovedSystemRequest() (request *ApprovedSystemRequest) {
	request = &ApprovedSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ApprovedSystem")
	return
}

func NewApprovedSystemResponse() (response *ApprovedSystemResponse) {
	response = &ApprovedSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统审批同意
func (c *Client) ApprovedSystem(request *ApprovedSystemRequest) (response *ApprovedSystemResponse, err error) {
	if request == nil {
		request = NewApprovedSystemRequest()
	}
	response = NewApprovedSystemResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCategoryRequest() (request *CreateCategoryRequest) {
	request = &CreateCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateCategory")
	return
}

func NewCreateCategoryResponse() (response *CreateCategoryResponse) {
	response = &CreateCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建⽬录
func (c *Client) CreateCategory(request *CreateCategoryRequest) (response *CreateCategoryResponse, err error) {
	if request == nil {
		request = NewCreateCategoryRequest()
	}
	response = NewCreateCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAreaRequest() (request *DeleteAreaRequest) {
	request = &DeleteAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteArea")
	return
}

func NewDeleteAreaResponse() (response *DeleteAreaResponse) {
	response = &DeleteAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除区域
func (c *Client) DeleteArea(request *DeleteAreaRequest) (response *DeleteAreaResponse, err error) {
	if request == nil {
		request = NewDeleteAreaRequest()
	}
	response = NewDeleteAreaResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCentreNodeInfoRequest() (request *ModifyCentreNodeInfoRequest) {
	request = &ModifyCentreNodeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyCentreNodeInfo")
	return
}

func NewModifyCentreNodeInfoResponse() (response *ModifyCentreNodeInfoResponse) {
	response = &ModifyCentreNodeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 刷新授权中心节点信息
func (c *Client) ModifyCentreNodeInfo(request *ModifyCentreNodeInfoRequest) (response *ModifyCentreNodeInfoResponse, err error) {
	if request == nil {
		request = NewModifyCentreNodeInfoRequest()
	}
	response = NewModifyCentreNodeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDocumentListRequest() (request *DescribeDocumentListRequest) {
	request = &DescribeDocumentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeDocumentList")
	return
}

func NewDescribeDocumentListResponse() (response *DescribeDocumentListResponse) {
	response = &DescribeDocumentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取分组
func (c *Client) DescribeDocumentList(request *DescribeDocumentListRequest) (response *DescribeDocumentListResponse, err error) {
	if request == nil {
		request = NewDescribeDocumentListRequest()
	}
	response = NewDescribeDocumentListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRolesRequest() (request *DescribeRolesRequest) {
	request = &DescribeRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeRoles")
	return
}

func NewDescribeRolesResponse() (response *DescribeRolesResponse) {
	response = &DescribeRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 角色列表
func (c *Client) DescribeRoles(request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
	if request == nil {
		request = NewDescribeRolesRequest()
	}
	response = NewDescribeRolesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSubscriptionRequest() (request *CreateSubscriptionRequest) {
	request = &CreateSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateSubscription")
	return
}

func NewCreateSubscriptionResponse() (response *CreateSubscriptionResponse) {
	response = &CreateSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 申请订阅
func (c *Client) CreateSubscription(request *CreateSubscriptionRequest) (response *CreateSubscriptionResponse, err error) {
	if request == nil {
		request = NewCreateSubscriptionRequest()
	}
	response = NewCreateSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVendorListRequest() (request *DescribeVendorListRequest) {
	request = &DescribeVendorListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeVendorList")
	return
}

func NewDescribeVendorListResponse() (response *DescribeVendorListResponse) {
	response = &DescribeVendorListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供应商列表查询
func (c *Client) DescribeVendorList(request *DescribeVendorListRequest) (response *DescribeVendorListResponse, err error) {
	if request == nil {
		request = NewDescribeVendorListRequest()
	}
	response = NewDescribeVendorListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCheckAccountRequest() (request *DescribeCheckAccountRequest) {
	request = &DescribeCheckAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeCheckAccount")
	return
}

func NewDescribeCheckAccountResponse() (response *DescribeCheckAccountResponse) {
	response = &DescribeCheckAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// check account是否唯一
func (c *Client) DescribeCheckAccount(request *DescribeCheckAccountRequest) (response *DescribeCheckAccountResponse, err error) {
	if request == nil {
		request = NewDescribeCheckAccountRequest()
	}
	response = NewDescribeCheckAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTSFMicroserviceRequest() (request *CreateTSFMicroserviceRequest) {
	request = &CreateTSFMicroserviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateTSFMicroservice")
	return
}

func NewCreateTSFMicroserviceResponse() (response *CreateTSFMicroserviceResponse) {
	response = &CreateTSFMicroserviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步租户下 TSF 微服务实例
func (c *Client) CreateTSFMicroservice(request *CreateTSFMicroserviceRequest) (response *CreateTSFMicroserviceResponse, err error) {
	if request == nil {
		request = NewCreateTSFMicroserviceRequest()
	}
	response = NewCreateTSFMicroserviceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataSourcesRequest() (request *DescribeDataSourcesRequest) {
	request = &DescribeDataSourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeDataSources")
	return
}

func NewDescribeDataSourcesResponse() (response *DescribeDataSourcesResponse) {
	response = &DescribeDataSourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据源列表
func (c *Client) DescribeDataSources(request *DescribeDataSourcesRequest) (response *DescribeDataSourcesResponse, err error) {
	if request == nil {
		request = NewDescribeDataSourcesRequest()
	}
	response = NewDescribeDataSourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAgencyBatchRequest() (request *DeleteAgencyBatchRequest) {
	request = &DeleteAgencyBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteAgencyBatch")
	return
}

func NewDeleteAgencyBatchResponse() (response *DeleteAgencyBatchResponse) {
	response = &DeleteAgencyBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除机构
func (c *Client) DeleteAgencyBatch(request *DeleteAgencyBatchRequest) (response *DeleteAgencyBatchResponse, err error) {
	if request == nil {
		request = NewDeleteAgencyBatchRequest()
	}
	response = NewDeleteAgencyBatchResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSmartGateRaidRoutingRequest() (request *CreateSmartGateRaidRoutingRequest) {
	request = &CreateSmartGateRaidRoutingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateSmartGateRaidRouting")
	return
}

func NewCreateSmartGateRaidRoutingResponse() (response *CreateSmartGateRaidRoutingResponse) {
	response = &CreateSmartGateRaidRoutingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建节点阵列路由
func (c *Client) CreateSmartGateRaidRouting(request *CreateSmartGateRaidRoutingRequest) (response *CreateSmartGateRaidRoutingResponse, err error) {
	if request == nil {
		request = NewCreateSmartGateRaidRoutingRequest()
	}
	response = NewCreateSmartGateRaidRoutingResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDisableAccountRequest() (request *ModifyDisableAccountRequest) {
	request = &ModifyDisableAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyDisableAccount")
	return
}

func NewModifyDisableAccountResponse() (response *ModifyDisableAccountResponse) {
	response = &ModifyDisableAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用帐号
func (c *Client) ModifyDisableAccount(request *ModifyDisableAccountRequest) (response *ModifyDisableAccountResponse, err error) {
	if request == nil {
		request = NewModifyDisableAccountRequest()
	}
	response = NewModifyDisableAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAppRequest() (request *DeleteAppRequest) {
	request = &DeleteAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteApp")
	return
}

func NewDeleteAppResponse() (response *DeleteAppResponse) {
	response = &DeleteAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用
func (c *Client) DeleteApp(request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
	if request == nil {
		request = NewDeleteAppRequest()
	}
	response = NewDeleteAppResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadAgencyImportResultRequest() (request *DownloadAgencyImportResultRequest) {
	request = &DownloadAgencyImportResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadAgencyImportResult")
	return
}

func NewDownloadAgencyImportResultResponse() (response *DownloadAgencyImportResultResponse) {
	response = &DownloadAgencyImportResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载文件导入机构成员结果
func (c *Client) DownloadAgencyImportResult(request *DownloadAgencyImportResultRequest) (response *DownloadAgencyImportResultResponse, err error) {
	if request == nil {
		request = NewDownloadAgencyImportResultRequest()
	}
	response = NewDownloadAgencyImportResultResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadRequest() (request *DownloadRequest) {
	request = &DownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "Download")
	return
}

func NewDownloadResponse() (response *DownloadResponse) {
	response = &DownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通用下载
func (c *Client) Download(request *DownloadRequest) (response *DownloadResponse, err error) {
	if request == nil {
		request = NewDownloadRequest()
	}
	response = NewDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadImportResultRequest() (request *DownloadImportResultRequest) {
	request = &DownloadImportResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadImportResult")
	return
}

func NewDownloadImportResultResponse() (response *DownloadImportResultResponse) {
	response = &DownloadImportResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载swagger导入的报告
func (c *Client) DownloadImportResult(request *DownloadImportResultRequest) (response *DownloadImportResultResponse, err error) {
	if request == nil {
		request = NewDownloadImportResultRequest()
	}
	response = NewDownloadImportResultResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateVendorRequest() (request *UpdateVendorRequest) {
	request = &UpdateVendorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UpdateVendor")
	return
}

func NewUpdateVendorResponse() (response *UpdateVendorResponse) {
	response = &UpdateVendorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新供应商
func (c *Client) UpdateVendor(request *UpdateVendorRequest) (response *UpdateVendorResponse, err error) {
	if request == nil {
		request = NewUpdateVendorRequest()
	}
	response = NewUpdateVendorResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAcsRoleRequest() (request *CreateAcsRoleRequest) {
	request = &CreateAcsRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateAcsRole")
	return
}

func NewCreateAcsRoleResponse() (response *CreateAcsRoleResponse) {
	response = &CreateAcsRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建角色
func (c *Client) CreateAcsRole(request *CreateAcsRoleRequest) (response *CreateAcsRoleResponse, err error) {
	if request == nil {
		request = NewCreateAcsRoleRequest()
	}
	response = NewCreateAcsRoleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVendorRequest() (request *CreateVendorRequest) {
	request = &CreateVendorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateVendor")
	return
}

func NewCreateVendorResponse() (response *CreateVendorResponse) {
	response = &CreateVendorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建供应商
func (c *Client) CreateVendor(request *CreateVendorRequest) (response *CreateVendorResponse, err error) {
	if request == nil {
		request = NewCreateVendorRequest()
	}
	response = NewCreateVendorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroProtoFileRequest() (request *DescribeMicroProtoFileRequest) {
	request = &DescribeMicroProtoFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroProtoFile")
	return
}

func NewDescribeMicroProtoFileResponse() (response *DescribeMicroProtoFileResponse) {
	response = &DescribeMicroProtoFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件查询
func (c *Client) DescribeMicroProtoFile(request *DescribeMicroProtoFileRequest) (response *DescribeMicroProtoFileResponse, err error) {
	if request == nil {
		request = NewDescribeMicroProtoFileRequest()
	}
	response = NewDescribeMicroProtoFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemApproveLogsRequest() (request *DescribeSystemApproveLogsRequest) {
	request = &DescribeSystemApproveLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSystemApproveLogs")
	return
}

func NewDescribeSystemApproveLogsResponse() (response *DescribeSystemApproveLogsResponse) {
	response = &DescribeSystemApproveLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统审批历史记录
func (c *Client) DescribeSystemApproveLogs(request *DescribeSystemApproveLogsRequest) (response *DescribeSystemApproveLogsResponse, err error) {
	if request == nil {
		request = NewDescribeSystemApproveLogsRequest()
	}
	response = NewDescribeSystemApproveLogsResponse()
	err = c.Send(request, response)
	return
}

func NewListAgencyProfileRequest() (request *ListAgencyProfileRequest) {
	request = &ListAgencyProfileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ListAgencyProfile")
	return
}

func NewListAgencyProfileResponse() (response *ListAgencyProfileResponse) {
	response = &ListAgencyProfileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机构成员用户信息
func (c *Client) ListAgencyProfile(request *ListAgencyProfileRequest) (response *ListAgencyProfileResponse, err error) {
	if request == nil {
		request = NewListAgencyProfileRequest()
	}
	response = NewListAgencyProfileResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProfileRequest() (request *UpdateProfileRequest) {
	request = &UpdateProfileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UpdateProfile")
	return
}

func NewUpdateProfileResponse() (response *UpdateProfileResponse) {
	response = &UpdateProfileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新帐号
func (c *Client) UpdateProfile(request *UpdateProfileRequest) (response *UpdateProfileResponse, err error) {
	if request == nil {
		request = NewUpdateProfileRequest()
	}
	response = NewUpdateProfileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTSFServicesRequest() (request *DescribeTSFServicesRequest) {
	request = &DescribeTSFServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTSFServices")
	return
}

func NewDescribeTSFServicesResponse() (response *DescribeTSFServicesResponse) {
	response = &DescribeTSFServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF 微服务 API 列表
func (c *Client) DescribeTSFServices(request *DescribeTSFServicesRequest) (response *DescribeTSFServicesResponse, err error) {
	if request == nil {
		request = NewDescribeTSFServicesRequest()
	}
	response = NewDescribeTSFServicesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAcsRoleRequest() (request *ModifyAcsRoleRequest) {
	request = &ModifyAcsRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyAcsRole")
	return
}

func NewModifyAcsRoleResponse() (response *ModifyAcsRoleResponse) {
	response = &ModifyAcsRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改角色
func (c *Client) ModifyAcsRole(request *ModifyAcsRoleRequest) (response *ModifyAcsRoleResponse, err error) {
	if request == nil {
		request = NewModifyAcsRoleRequest()
	}
	response = NewModifyAcsRoleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMicroProtoGroupRequest() (request *ModifyMicroProtoGroupRequest) {
	request = &ModifyMicroProtoGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyMicroProtoGroup")
	return
}

func NewModifyMicroProtoGroupResponse() (response *ModifyMicroProtoGroupResponse) {
	response = &ModifyMicroProtoGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件组修改
func (c *Client) ModifyMicroProtoGroup(request *ModifyMicroProtoGroupRequest) (response *ModifyMicroProtoGroupResponse, err error) {
	if request == nil {
		request = NewModifyMicroProtoGroupRequest()
	}
	response = NewModifyMicroProtoGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMemberBatchRequest() (request *DeleteMemberBatchRequest) {
	request = &DeleteMemberBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMemberBatch")
	return
}

func NewDeleteMemberBatchResponse() (response *DeleteMemberBatchResponse) {
	response = &DeleteMemberBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除机构成员
func (c *Client) DeleteMemberBatch(request *DeleteMemberBatchRequest) (response *DeleteMemberBatchResponse, err error) {
	if request == nil {
		request = NewDeleteMemberBatchRequest()
	}
	response = NewDeleteMemberBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTSFNamespacesRequest() (request *DescribeTSFNamespacesRequest) {
	request = &DescribeTSFNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTSFNamespaces")
	return
}

func NewDescribeTSFNamespacesResponse() (response *DescribeTSFNamespacesResponse) {
	response = &DescribeTSFNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF 命令空间列表
func (c *Client) DescribeTSFNamespaces(request *DescribeTSFNamespacesRequest) (response *DescribeTSFNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeTSFNamespacesRequest()
	}
	response = NewDescribeTSFNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDisabledRequest() (request *DisabledRequest) {
	request = &DisabledRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "Disabled")
	return
}

func NewDisabledResponse() (response *DisabledResponse) {
	response = &DisabledResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用API服务
func (c *Client) Disabled(request *DisabledRequest) (response *DisabledResponse, err error) {
	if request == nil {
		request = NewDisabledRequest()
	}
	response = NewDisabledResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroProtoFileRequest() (request *DeleteMicroProtoFileRequest) {
	request = &DeleteMicroProtoFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMicroProtoFile")
	return
}

func NewDeleteMicroProtoFileResponse() (response *DeleteMicroProtoFileResponse) {
	response = &DeleteMicroProtoFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件删除
func (c *Client) DeleteMicroProtoFile(request *DeleteMicroProtoFileRequest) (response *DeleteMicroProtoFileResponse, err error) {
	if request == nil {
		request = NewDeleteMicroProtoFileRequest()
	}
	response = NewDeleteMicroProtoFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroSyncJobsRequest() (request *DescribeMicroSyncJobsRequest) {
	request = &DescribeMicroSyncJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroSyncJobs")
	return
}

func NewDescribeMicroSyncJobsResponse() (response *DescribeMicroSyncJobsResponse) {
	response = &DescribeMicroSyncJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务同步记录列表
func (c *Client) DescribeMicroSyncJobs(request *DescribeMicroSyncJobsRequest) (response *DescribeMicroSyncJobsResponse, err error) {
	if request == nil {
		request = NewDescribeMicroSyncJobsRequest()
	}
	response = NewDescribeMicroSyncJobsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServicePageRequest() (request *DescribeServicePageRequest) {
	request = &DescribeServicePageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeServicePage")
	return
}

func NewDescribeServicePageResponse() (response *DescribeServicePageResponse) {
	response = &DescribeServicePageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询API服务列表
func (c *Client) DescribeServicePage(request *DescribeServicePageRequest) (response *DescribeServicePageResponse, err error) {
	if request == nil {
		request = NewDescribeServicePageRequest()
	}
	response = NewDescribeServicePageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSmartGatewayGroupRequest() (request *DescribeSmartGatewayGroupRequest) {
	request = &DescribeSmartGatewayGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSmartGatewayGroup")
	return
}

func NewDescribeSmartGatewayGroupResponse() (response *DescribeSmartGatewayGroupResponse) {
	response = &DescribeSmartGatewayGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点阵列详情
func (c *Client) DescribeSmartGatewayGroup(request *DescribeSmartGatewayGroupRequest) (response *DescribeSmartGatewayGroupResponse, err error) {
	if request == nil {
		request = NewDescribeSmartGatewayGroupRequest()
	}
	response = NewDescribeSmartGatewayGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAgencyRequest() (request *UpdateAgencyRequest) {
	request = &UpdateAgencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UpdateAgency")
	return
}

func NewUpdateAgencyResponse() (response *UpdateAgencyResponse) {
	response = &UpdateAgencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新机构
func (c *Client) UpdateAgency(request *UpdateAgencyRequest) (response *UpdateAgencyResponse, err error) {
	if request == nil {
		request = NewUpdateAgencyRequest()
	}
	response = NewUpdateAgencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemApprovesRequest() (request *DescribeSystemApprovesRequest) {
	request = &DescribeSystemApprovesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSystemApproves")
	return
}

func NewDescribeSystemApprovesResponse() (response *DescribeSystemApprovesResponse) {
	response = &DescribeSystemApprovesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询系统审批列表
func (c *Client) DescribeSystemApproves(request *DescribeSystemApprovesRequest) (response *DescribeSystemApprovesResponse, err error) {
	if request == nil {
		request = NewDescribeSystemApprovesRequest()
	}
	response = NewDescribeSystemApprovesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAcsRolePermissionRequest() (request *DescribeAcsRolePermissionRequest) {
	request = &DescribeAcsRolePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAcsRolePermission")
	return
}

func NewDescribeAcsRolePermissionResponse() (response *DescribeAcsRolePermissionResponse) {
	response = &DescribeAcsRolePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取角色名下权限列表
func (c *Client) DescribeAcsRolePermission(request *DescribeAcsRolePermissionRequest) (response *DescribeAcsRolePermissionResponse, err error) {
	if request == nil {
		request = NewDescribeAcsRolePermissionRequest()
	}
	response = NewDescribeAcsRolePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAcsRoleMembersRequest() (request *DescribeAcsRoleMembersRequest) {
	request = &DescribeAcsRoleMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAcsRoleMembers")
	return
}

func NewDescribeAcsRoleMembersResponse() (response *DescribeAcsRoleMembersResponse) {
	response = &DescribeAcsRoleMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取角色名下用户列表
func (c *Client) DescribeAcsRoleMembers(request *DescribeAcsRoleMembersRequest) (response *DescribeAcsRoleMembersResponse, err error) {
	if request == nil {
		request = NewDescribeAcsRoleMembersRequest()
	}
	response = NewDescribeAcsRoleMembersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConnectorRequest() (request *ModifyConnectorRequest) {
	request = &ModifyConnectorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyConnector")
	return
}

func NewModifyConnectorResponse() (response *ModifyConnectorResponse) {
	response = &ModifyConnectorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 认证修改
func (c *Client) ModifyConnector(request *ModifyConnectorRequest) (response *ModifyConnectorResponse, err error) {
	if request == nil {
		request = NewModifyConnectorRequest()
	}
	response = NewModifyConnectorResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPasswordByOldRequest() (request *ModifyPasswordByOldRequest) {
	request = &ModifyPasswordByOldRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyPasswordByOld")
	return
}

func NewModifyPasswordByOldResponse() (response *ModifyPasswordByOldResponse) {
	response = &ModifyPasswordByOldResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过原密码修改密码
func (c *Client) ModifyPasswordByOld(request *ModifyPasswordByOldRequest) (response *ModifyPasswordByOldResponse, err error) {
	if request == nil {
		request = NewModifyPasswordByOldRequest()
	}
	response = NewModifyPasswordByOldResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVendorRequest() (request *DeleteVendorRequest) {
	request = &DeleteVendorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteVendor")
	return
}

func NewDeleteVendorResponse() (response *DeleteVendorResponse) {
	response = &DeleteVendorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除供应商
func (c *Client) DeleteVendor(request *DeleteVendorRequest) (response *DeleteVendorResponse, err error) {
	if request == nil {
		request = NewDeleteVendorRequest()
	}
	response = NewDeleteVendorResponse()
	err = c.Send(request, response)
	return
}

func NewRejectedSystemBatchRequest() (request *RejectedSystemBatchRequest) {
	request = &RejectedSystemBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RejectedSystemBatch")
	return
}

func NewRejectedSystemBatchResponse() (response *RejectedSystemBatchResponse) {
	response = &RejectedSystemBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统审批驳回批量
func (c *Client) RejectedSystemBatch(request *RejectedSystemBatchRequest) (response *RejectedSystemBatchResponse, err error) {
	if request == nil {
		request = NewRejectedSystemBatchRequest()
	}
	response = NewRejectedSystemBatchResponse()
	err = c.Send(request, response)
	return
}

func NewApprovedServiceRequest() (request *ApprovedServiceRequest) {
	request = &ApprovedServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ApprovedService")
	return
}

func NewApprovedServiceResponse() (response *ApprovedServiceResponse) {
	response = &ApprovedServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务审批同意
func (c *Client) ApprovedService(request *ApprovedServiceRequest) (response *ApprovedServiceResponse, err error) {
	if request == nil {
		request = NewApprovedServiceRequest()
	}
	response = NewApprovedServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroTenantRequest() (request *DeleteMicroTenantRequest) {
	request = &DeleteMicroTenantRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMicroTenant")
	return
}

func NewDeleteMicroTenantResponse() (response *DeleteMicroTenantResponse) {
	response = &DeleteMicroTenantResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务租户删除
func (c *Client) DeleteMicroTenant(request *DeleteMicroTenantRequest) (response *DeleteMicroTenantResponse, err error) {
	if request == nil {
		request = NewDeleteMicroTenantRequest()
	}
	response = NewDeleteMicroTenantResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
	request = &DeleteZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteZone")
	return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
	response = &DeleteZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除分区
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
	if request == nil {
		request = NewDeleteZoneRequest()
	}
	response = NewDeleteZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTSFSyncTypesRequest() (request *DescribeTSFSyncTypesRequest) {
	request = &DescribeTSFSyncTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTSFSyncTypes")
	return
}

func NewDescribeTSFSyncTypesResponse() (response *DescribeTSFSyncTypesResponse) {
	response = &DescribeTSFSyncTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务同步方式列表
func (c *Client) DescribeTSFSyncTypes(request *DescribeTSFSyncTypesRequest) (response *DescribeTSFSyncTypesResponse, err error) {
	if request == nil {
		request = NewDescribeTSFSyncTypesRequest()
	}
	response = NewDescribeTSFSyncTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVendorMemberBatchListRequest() (request *DescribeVendorMemberBatchListRequest) {
	request = &DescribeVendorMemberBatchListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeVendorMemberBatchList")
	return
}

func NewDescribeVendorMemberBatchListResponse() (response *DescribeVendorMemberBatchListResponse) {
	response = &DescribeVendorMemberBatchListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询供应商成员，通过供应商ID列表批量查询
func (c *Client) DescribeVendorMemberBatchList(request *DescribeVendorMemberBatchListRequest) (response *DescribeVendorMemberBatchListResponse, err error) {
	if request == nil {
		request = NewDescribeVendorMemberBatchListRequest()
	}
	response = NewDescribeVendorMemberBatchListResponse()
	err = c.Send(request, response)
	return
}

func NewImportAgencyMemberRequest() (request *ImportAgencyMemberRequest) {
	request = &ImportAgencyMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ImportAgencyMember")
	return
}

func NewImportAgencyMemberResponse() (response *ImportAgencyMemberResponse) {
	response = &ImportAgencyMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 机构用户导入
func (c *Client) ImportAgencyMember(request *ImportAgencyMemberRequest) (response *ImportAgencyMemberResponse, err error) {
	if request == nil {
		request = NewImportAgencyMemberRequest()
	}
	response = NewImportAgencyMemberResponse()
	err = c.Send(request, response)
	return
}

func NewBatchDeleteCategoryRequest() (request *BatchDeleteCategoryRequest) {
	request = &BatchDeleteCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "BatchDeleteCategory")
	return
}

func NewBatchDeleteCategoryResponse() (response *BatchDeleteCategoryResponse) {
	response = &BatchDeleteCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除⽬录
func (c *Client) BatchDeleteCategory(request *BatchDeleteCategoryRequest) (response *BatchDeleteCategoryResponse, err error) {
	if request == nil {
		request = NewBatchDeleteCategoryRequest()
	}
	response = NewBatchDeleteCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceRequestLogBodyDetailRequest() (request *DescribeServiceRequestLogBodyDetailRequest) {
	request = &DescribeServiceRequestLogBodyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeServiceRequestLogBodyDetail")
	return
}

func NewDescribeServiceRequestLogBodyDetailResponse() (response *DescribeServiceRequestLogBodyDetailResponse) {
	response = &DescribeServiceRequestLogBodyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务请求报文详情日志
func (c *Client) DescribeServiceRequestLogBodyDetail(request *DescribeServiceRequestLogBodyDetailRequest) (response *DescribeServiceRequestLogBodyDetailResponse, err error) {
	if request == nil {
		request = NewDescribeServiceRequestLogBodyDetailRequest()
	}
	response = NewDescribeServiceRequestLogBodyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewEnableRequest() (request *EnableRequest) {
	request = &EnableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "Enable")
	return
}

func NewEnableResponse() (response *EnableResponse) {
	response = &EnableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用API服务
func (c *Client) Enable(request *EnableRequest) (response *EnableResponse, err error) {
	if request == nil {
		request = NewEnableRequest()
	}
	response = NewEnableResponse()
	err = c.Send(request, response)
	return
}

func NewModifySmartGateRaidRoutingRequest() (request *ModifySmartGateRaidRoutingRequest) {
	request = &ModifySmartGateRaidRoutingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifySmartGateRaidRouting")
	return
}

func NewModifySmartGateRaidRoutingResponse() (response *ModifySmartGateRaidRoutingResponse) {
	response = &ModifySmartGateRaidRoutingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改节点阵列路由
func (c *Client) ModifySmartGateRaidRouting(request *ModifySmartGateRaidRoutingRequest) (response *ModifySmartGateRaidRoutingResponse, err error) {
	if request == nil {
		request = NewModifySmartGateRaidRoutingRequest()
	}
	response = NewModifySmartGateRaidRoutingResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTSFNamespaceRequest() (request *CreateTSFNamespaceRequest) {
	request = &CreateTSFNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateTSFNamespace")
	return
}

func NewCreateTSFNamespaceResponse() (response *CreateTSFNamespaceResponse) {
	response = &CreateTSFNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步租户下 TSF 命名空间
func (c *Client) CreateTSFNamespace(request *CreateTSFNamespaceRequest) (response *CreateTSFNamespaceResponse, err error) {
	if request == nil {
		request = NewCreateTSFNamespaceRequest()
	}
	response = NewCreateTSFNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceMarketPageRequest() (request *DescribeServiceMarketPageRequest) {
	request = &DescribeServiceMarketPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeServiceMarketPage")
	return
}

func NewDescribeServiceMarketPageResponse() (response *DescribeServiceMarketPageResponse) {
	response = &DescribeServiceMarketPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取API市场列表
func (c *Client) DescribeServiceMarketPage(request *DescribeServiceMarketPageRequest) (response *DescribeServiceMarketPageResponse, err error) {
	if request == nil {
		request = NewDescribeServiceMarketPageRequest()
	}
	response = NewDescribeServiceMarketPageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneRequest() (request *DescribeZoneRequest) {
	request = &DescribeZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeZone")
	return
}

func NewDescribeZoneResponse() (response *DescribeZoneResponse) {
	response = &DescribeZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分区详情
func (c *Client) DescribeZone(request *DescribeZoneRequest) (response *DescribeZoneResponse, err error) {
	if request == nil {
		request = NewDescribeZoneRequest()
	}
	response = NewDescribeZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnSyncPaasRequest() (request *DescribeUnSyncPaasRequest) {
	request = &DescribeUnSyncPaasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeUnSyncPaas")
	return
}

func NewDescribeUnSyncPaasResponse() (response *DescribeUnSyncPaasResponse) {
	response = &DescribeUnSyncPaasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取已同步过的应用列表
func (c *Client) DescribeUnSyncPaas(request *DescribeUnSyncPaasRequest) (response *DescribeUnSyncPaasResponse, err error) {
	if request == nil {
		request = NewDescribeUnSyncPaasRequest()
	}
	response = NewDescribeUnSyncPaasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
	request = &DescribeServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeService")
	return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
	response = &DescribeServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询API服务详情信息（适用于编辑页面）
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
	if request == nil {
		request = NewDescribeServiceRequest()
	}
	response = NewDescribeServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVendorMemberListRequest() (request *DescribeVendorMemberListRequest) {
	request = &DescribeVendorMemberListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeVendorMemberList")
	return
}

func NewDescribeVendorMemberListResponse() (response *DescribeVendorMemberListResponse) {
	response = &DescribeVendorMemberListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询供应商成员列表
func (c *Client) DescribeVendorMemberList(request *DescribeVendorMemberListRequest) (response *DescribeVendorMemberListResponse, err error) {
	if request == nil {
		request = NewDescribeVendorMemberListRequest()
	}
	response = NewDescribeVendorMemberListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceTop10Request() (request *DescribeServiceTop10Request) {
	request = &DescribeServiceTop10Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeServiceTop10")
	return
}

func NewDescribeServiceTop10Response() (response *DescribeServiceTop10Response) {
	response = &DescribeServiceTop10Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Top10服务
func (c *Client) DescribeServiceTop10(request *DescribeServiceTop10Request) (response *DescribeServiceTop10Response, err error) {
	if request == nil {
		request = NewDescribeServiceTop10Request()
	}
	response = NewDescribeServiceTop10Response()
	err = c.Send(request, response)
	return
}

func NewDescribeRangeMetricsRequest() (request *DescribeRangeMetricsRequest) {
	request = &DescribeRangeMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeRangeMetrics")
	return
}

func NewDescribeRangeMetricsResponse() (response *DescribeRangeMetricsResponse) {
	response = &DescribeRangeMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询时间范围指标
func (c *Client) DescribeRangeMetrics(request *DescribeRangeMetricsRequest) (response *DescribeRangeMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeRangeMetricsRequest()
	}
	response = NewDescribeRangeMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemsForNavRequest() (request *DescribeSystemsForNavRequest) {
	request = &DescribeSystemsForNavRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSystemsForNav")
	return
}

func NewDescribeSystemsForNavResponse() (response *DescribeSystemsForNavResponse) {
	response = &DescribeSystemsForNavResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询系统列表接口(导航)
func (c *Client) DescribeSystemsForNav(request *DescribeSystemsForNavRequest) (response *DescribeSystemsForNavResponse, err error) {
	if request == nil {
		request = NewDescribeSystemsForNavRequest()
	}
	response = NewDescribeSystemsForNavResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConnectorRequest() (request *DescribeConnectorRequest) {
	request = &DescribeConnectorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeConnector")
	return
}

func NewDescribeConnectorResponse() (response *DescribeConnectorResponse) {
	response = &DescribeConnectorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 认证详情
func (c *Client) DescribeConnector(request *DescribeConnectorRequest) (response *DescribeConnectorResponse, err error) {
	if request == nil {
		request = NewDescribeConnectorRequest()
	}
	response = NewDescribeConnectorResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDataSourceRequest() (request *ModifyDataSourceRequest) {
	request = &ModifyDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyDataSource")
	return
}

func NewModifyDataSourceResponse() (response *ModifyDataSourceResponse) {
	response = &ModifyDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据源修改
func (c *Client) ModifyDataSource(request *ModifyDataSourceRequest) (response *ModifyDataSourceResponse, err error) {
	if request == nil {
		request = NewModifyDataSourceRequest()
	}
	response = NewModifyDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAcsRoleRequest() (request *DeleteAcsRoleRequest) {
	request = &DeleteAcsRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteAcsRole")
	return
}

func NewDeleteAcsRoleResponse() (response *DeleteAcsRoleResponse) {
	response = &DeleteAcsRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除角色
func (c *Client) DeleteAcsRole(request *DeleteAcsRoleRequest) (response *DeleteAcsRoleResponse, err error) {
	if request == nil {
		request = NewDeleteAcsRoleRequest()
	}
	response = NewDeleteAcsRoleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMicroProtoGroupRequest() (request *CreateMicroProtoGroupRequest) {
	request = &CreateMicroProtoGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateMicroProtoGroup")
	return
}

func NewCreateMicroProtoGroupResponse() (response *CreateMicroProtoGroupResponse) {
	response = &CreateMicroProtoGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件组创建
func (c *Client) CreateMicroProtoGroup(request *CreateMicroProtoGroupRequest) (response *CreateMicroProtoGroupResponse, err error) {
	if request == nil {
		request = NewCreateMicroProtoGroupRequest()
	}
	response = NewCreateMicroProtoGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlertHistoryRequest() (request *DescribeAlertHistoryRequest) {
	request = &DescribeAlertHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAlertHistory")
	return
}

func NewDescribeAlertHistoryResponse() (response *DescribeAlertHistoryResponse) {
	response = &DescribeAlertHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询告警历史
func (c *Client) DescribeAlertHistory(request *DescribeAlertHistoryRequest) (response *DescribeAlertHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeAlertHistoryRequest()
	}
	response = NewDescribeAlertHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMemberRequest() (request *DeleteMemberRequest) {
	request = &DeleteMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteMember")
	return
}

func NewDeleteMemberResponse() (response *DeleteMemberResponse) {
	response = &DeleteMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机构成员
func (c *Client) DeleteMember(request *DeleteMemberRequest) (response *DeleteMemberResponse, err error) {
	if request == nil {
		request = NewDeleteMemberRequest()
	}
	response = NewDeleteMemberResponse()
	err = c.Send(request, response)
	return
}

func NewRejectedServiceBatchRequest() (request *RejectedServiceBatchRequest) {
	request = &RejectedServiceBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RejectedServiceBatch")
	return
}

func NewRejectedServiceBatchResponse() (response *RejectedServiceBatchResponse) {
	response = &RejectedServiceBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务审批驳回批量
func (c *Client) RejectedServiceBatch(request *RejectedServiceBatchRequest) (response *RejectedServiceBatchResponse, err error) {
	if request == nil {
		request = NewRejectedServiceBatchRequest()
	}
	response = NewRejectedServiceBatchResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConnectorRequest() (request *CreateConnectorRequest) {
	request = &CreateConnectorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateConnector")
	return
}

func NewCreateConnectorResponse() (response *CreateConnectorResponse) {
	response = &CreateConnectorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 认证创建
func (c *Client) CreateConnector(request *CreateConnectorRequest) (response *CreateConnectorResponse, err error) {
	if request == nil {
		request = NewCreateConnectorRequest()
	}
	response = NewCreateConnectorResponse()
	err = c.Send(request, response)
	return
}

func NewEnabledSiteRequest() (request *EnabledSiteRequest) {
	request = &EnabledSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "EnabledSite")
	return
}

func NewEnabledSiteResponse() (response *EnabledSiteResponse) {
	response = &EnabledSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用站点
func (c *Client) EnabledSite(request *EnabledSiteRequest) (response *EnabledSiteResponse, err error) {
	if request == nil {
		request = NewEnabledSiteRequest()
	}
	response = NewEnabledSiteResponse()
	err = c.Send(request, response)
	return
}

func NewValidateServicePubPathRequest() (request *ValidateServicePubPathRequest) {
	request = &ValidateServicePubPathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ValidateServicePubPath")
	return
}

func NewValidateServicePubPathResponse() (response *ValidateServicePubPathResponse) {
	response = &ValidateServicePubPathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检测服务公共路径是否冲突
func (c *Client) ValidateServicePubPath(request *ValidateServicePubPathRequest) (response *ValidateServicePubPathResponse, err error) {
	if request == nil {
		request = NewValidateServicePubPathRequest()
	}
	response = NewValidateServicePubPathResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTipsRequest() (request *DescribeTipsRequest) {
	request = &DescribeTipsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTips")
	return
}

func NewDescribeTipsResponse() (response *DescribeTipsResponse) {
	response = &DescribeTipsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询提示语
func (c *Client) DescribeTips(request *DescribeTipsRequest) (response *DescribeTipsResponse, err error) {
	if request == nil {
		request = NewDescribeTipsRequest()
	}
	response = NewDescribeTipsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSnapshotOriginDataRequest() (request *DeleteSnapshotOriginDataRequest) {
	request = &DeleteSnapshotOriginDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteSnapshotOriginData")
	return
}

func NewDeleteSnapshotOriginDataResponse() (response *DeleteSnapshotOriginDataResponse) {
	response = &DeleteSnapshotOriginDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除存档源数据
func (c *Client) DeleteSnapshotOriginData(request *DeleteSnapshotOriginDataRequest) (response *DeleteSnapshotOriginDataResponse, err error) {
	if request == nil {
		request = NewDeleteSnapshotOriginDataRequest()
	}
	response = NewDeleteSnapshotOriginDataResponse()
	err = c.Send(request, response)
	return
}

func NewRejectedSubscriptionBatchRequest() (request *RejectedSubscriptionBatchRequest) {
	request = &RejectedSubscriptionBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RejectedSubscriptionBatch")
	return
}

func NewRejectedSubscriptionBatchResponse() (response *RejectedSubscriptionBatchResponse) {
	response = &RejectedSubscriptionBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拒绝订阅批量
func (c *Client) RejectedSubscriptionBatch(request *RejectedSubscriptionBatchRequest) (response *RejectedSubscriptionBatchResponse, err error) {
	if request == nil {
		request = NewRejectedSubscriptionBatchRequest()
	}
	response = NewRejectedSubscriptionBatchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMicroTenantRequest() (request *ModifyMicroTenantRequest) {
	request = &ModifyMicroTenantRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyMicroTenant")
	return
}

func NewModifyMicroTenantResponse() (response *ModifyMicroTenantResponse) {
	response = &ModifyMicroTenantResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务租户修改
func (c *Client) ModifyMicroTenant(request *ModifyMicroTenantRequest) (response *ModifyMicroTenantResponse, err error) {
	if request == nil {
		request = NewModifyMicroTenantRequest()
	}
	response = NewModifyMicroTenantResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppRequest() (request *ModifyAppRequest) {
	request = &ModifyAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyApp")
	return
}

func NewModifyAppResponse() (response *ModifyAppResponse) {
	response = &ModifyAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改应用
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
	if request == nil {
		request = NewModifyAppRequest()
	}
	response = NewModifyAppResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccountAndProfileBatchRequest() (request *DeleteAccountAndProfileBatchRequest) {
	request = &DeleteAccountAndProfileBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteAccountAndProfileBatch")
	return
}

func NewDeleteAccountAndProfileBatchResponse() (response *DeleteAccountAndProfileBatchResponse) {
	response = &DeleteAccountAndProfileBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除帐号
func (c *Client) DeleteAccountAndProfileBatch(request *DeleteAccountAndProfileBatchRequest) (response *DeleteAccountAndProfileBatchResponse, err error) {
	if request == nil {
		request = NewDeleteAccountAndProfileBatchRequest()
	}
	response = NewDeleteAccountAndProfileBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllSystemAndAppRequest() (request *DescribeAllSystemAndAppRequest) {
	request = &DescribeAllSystemAndAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAllSystemAndApp")
	return
}

func NewDescribeAllSystemAndAppResponse() (response *DescribeAllSystemAndAppResponse) {
	response = &DescribeAllSystemAndAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询全部系统应用列表
func (c *Client) DescribeAllSystemAndApp(request *DescribeAllSystemAndAppRequest) (response *DescribeAllSystemAndAppResponse, err error) {
	if request == nil {
		request = NewDescribeAllSystemAndAppRequest()
	}
	response = NewDescribeAllSystemAndAppResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSiteTop10Request() (request *DescribeSiteTop10Request) {
	request = &DescribeSiteTop10Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSiteTop10")
	return
}

func NewDescribeSiteTop10Response() (response *DescribeSiteTop10Response) {
	response = &DescribeSiteTop10Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Top10站点
func (c *Client) DescribeSiteTop10(request *DescribeSiteTop10Request) (response *DescribeSiteTop10Response, err error) {
	if request == nil {
		request = NewDescribeSiteTop10Request()
	}
	response = NewDescribeSiteTop10Response()
	err = c.Send(request, response)
	return
}

func NewDownloadImportTemplateRequest() (request *DownloadImportTemplateRequest) {
	request = &DownloadImportTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadImportTemplate")
	return
}

func NewDownloadImportTemplateResponse() (response *DownloadImportTemplateResponse) {
	response = &DownloadImportTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供应商下载导入模板
func (c *Client) DownloadImportTemplate(request *DownloadImportTemplateRequest) (response *DownloadImportTemplateResponse, err error) {
	if request == nil {
		request = NewDownloadImportTemplateRequest()
	}
	response = NewDownloadImportTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMemberRequest() (request *CreateMemberRequest) {
	request = &CreateMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateMember")
	return
}

func NewCreateMemberResponse() (response *CreateMemberResponse) {
	response = &CreateMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建机构成员
func (c *Client) CreateMember(request *CreateMemberRequest) (response *CreateMemberResponse, err error) {
	if request == nil {
		request = NewCreateMemberRequest()
	}
	response = NewCreateMemberResponse()
	err = c.Send(request, response)
	return
}

func NewUnFreezeSiteRequest() (request *UnFreezeSiteRequest) {
	request = &UnFreezeSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UnFreezeSite")
	return
}

func NewUnFreezeSiteResponse() (response *UnFreezeSiteResponse) {
	response = &UnFreezeSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解冻站点
func (c *Client) UnFreezeSite(request *UnFreezeSiteRequest) (response *UnFreezeSiteResponse, err error) {
	if request == nil {
		request = NewUnFreezeSiteRequest()
	}
	response = NewUnFreezeSiteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAreasRequest() (request *DescribeAreasRequest) {
	request = &DescribeAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAreas")
	return
}

func NewDescribeAreasResponse() (response *DescribeAreasResponse) {
	response = &DescribeAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 区域列表
func (c *Client) DescribeAreas(request *DescribeAreasRequest) (response *DescribeAreasResponse, err error) {
	if request == nil {
		request = NewDescribeAreasRequest()
	}
	response = NewDescribeAreasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLabelsRequest() (request *DescribeLabelsRequest) {
	request = &DescribeLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeLabels")
	return
}

func NewDescribeLabelsResponse() (response *DescribeLabelsResponse) {
	response = &DescribeLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取标签列表
func (c *Client) DescribeLabels(request *DescribeLabelsRequest) (response *DescribeLabelsResponse, err error) {
	if request == nil {
		request = NewDescribeLabelsRequest()
	}
	response = NewDescribeLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOverviewRequest() (request *DescribeOverviewRequest) {
	request = &DescribeOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeOverview")
	return
}

func NewDescribeOverviewResponse() (response *DescribeOverviewResponse) {
	response = &DescribeOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用概览信息
func (c *Client) DescribeOverview(request *DescribeOverviewRequest) (response *DescribeOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeOverviewRequest()
	}
	response = NewDescribeOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotRequest() (request *DescribeSnapshotRequest) {
	request = &DescribeSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSnapshot")
	return
}

func NewDescribeSnapshotResponse() (response *DescribeSnapshotResponse) {
	response = &DescribeSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询存档
func (c *Client) DescribeSnapshot(request *DescribeSnapshotRequest) (response *DescribeSnapshotResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotRequest()
	}
	response = NewDescribeSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewRevokeSubscriptionRequest() (request *RevokeSubscriptionRequest) {
	request = &RevokeSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RevokeSubscription")
	return
}

func NewRevokeSubscriptionResponse() (response *RevokeSubscriptionResponse) {
	response = &RevokeSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订阅取消授权
func (c *Client) RevokeSubscription(request *RevokeSubscriptionRequest) (response *RevokeSubscriptionResponse, err error) {
	if request == nil {
		request = NewRevokeSubscriptionRequest()
	}
	response = NewRevokeSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAcsRoleUniqueRequest() (request *CheckAcsRoleUniqueRequest) {
	request = &CheckAcsRoleUniqueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CheckAcsRoleUnique")
	return
}

func NewCheckAcsRoleUniqueResponse() (response *CheckAcsRoleUniqueResponse) {
	response = &CheckAcsRoleUniqueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查角色属性唯一
func (c *Client) CheckAcsRoleUnique(request *CheckAcsRoleUniqueRequest) (response *CheckAcsRoleUniqueResponse, err error) {
	if request == nil {
		request = NewCheckAcsRoleUniqueRequest()
	}
	response = NewCheckAcsRoleUniqueResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAgencyRequest() (request *CreateAgencyRequest) {
	request = &CreateAgencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateAgency")
	return
}

func NewCreateAgencyResponse() (response *CreateAgencyResponse) {
	response = &CreateAgencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建机构
func (c *Client) CreateAgency(request *CreateAgencyRequest) (response *CreateAgencyResponse, err error) {
	if request == nil {
		request = NewCreateAgencyRequest()
	}
	response = NewCreateAgencyResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserPermissionsRequest() (request *GetUserPermissionsRequest) {
	request = &GetUserPermissionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "GetUserPermissions")
	return
}

func NewGetUserPermissionsResponse() (response *GetUserPermissionsResponse) {
	response = &GetUserPermissionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户权限数组
func (c *Client) GetUserPermissions(request *GetUserPermissionsRequest) (response *GetUserPermissionsResponse, err error) {
	if request == nil {
		request = NewGetUserPermissionsRequest()
	}
	response = NewGetUserPermissionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVendorMemberRequest() (request *DeleteVendorMemberRequest) {
	request = &DeleteVendorMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteVendorMember")
	return
}

func NewDeleteVendorMemberResponse() (response *DeleteVendorMemberResponse) {
	response = &DeleteVendorMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除单个/批量供应商成员
func (c *Client) DeleteVendorMember(request *DeleteVendorMemberRequest) (response *DeleteVendorMemberResponse, err error) {
	if request == nil {
		request = NewDeleteVendorMemberRequest()
	}
	response = NewDeleteVendorMemberResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppRequest() (request *CreateAppRequest) {
	request = &CreateAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateApp")
	return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
	response = &CreateAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建应用
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
	if request == nil {
		request = NewCreateAppRequest()
	}
	response = NewCreateAppResponse()
	err = c.Send(request, response)
	return
}

func NewModifySqlTemplateRequest() (request *ModifySqlTemplateRequest) {
	request = &ModifySqlTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifySqlTemplate")
	return
}

func NewModifySqlTemplateResponse() (response *ModifySqlTemplateResponse) {
	response = &ModifySqlTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SQL 模板修改
func (c *Client) ModifySqlTemplate(request *ModifySqlTemplateRequest) (response *ModifySqlTemplateResponse, err error) {
	if request == nil {
		request = NewModifySqlTemplateRequest()
	}
	response = NewModifySqlTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroTenantRequest() (request *DescribeMicroTenantRequest) {
	request = &DescribeMicroTenantRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroTenant")
	return
}

func NewDescribeMicroTenantResponse() (response *DescribeMicroTenantResponse) {
	response = &DescribeMicroTenantResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务租户查询
func (c *Client) DescribeMicroTenant(request *DescribeMicroTenantRequest) (response *DescribeMicroTenantResponse, err error) {
	if request == nil {
		request = NewDescribeMicroTenantRequest()
	}
	response = NewDescribeMicroTenantResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPasswordRequest() (request *ModifyPasswordRequest) {
	request = &ModifyPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyPassword")
	return
}

func NewModifyPasswordResponse() (response *ModifyPasswordResponse) {
	response = &ModifyPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新密码
func (c *Client) ModifyPassword(request *ModifyPasswordRequest) (response *ModifyPasswordResponse, err error) {
	if request == nil {
		request = NewModifyPasswordRequest()
	}
	response = NewModifyPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAgencyManagerRequest() (request *DeleteAgencyManagerRequest) {
	request = &DeleteAgencyManagerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteAgencyManager")
	return
}

func NewDeleteAgencyManagerResponse() (response *DeleteAgencyManagerResponse) {
	response = &DeleteAgencyManagerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从机构负责人角色删除用户
func (c *Client) DeleteAgencyManager(request *DeleteAgencyManagerRequest) (response *DeleteAgencyManagerResponse, err error) {
	if request == nil {
		request = NewDeleteAgencyManagerRequest()
	}
	response = NewDeleteAgencyManagerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroTenantsRequest() (request *DescribeMicroTenantsRequest) {
	request = &DescribeMicroTenantsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroTenants")
	return
}

func NewDescribeMicroTenantsResponse() (response *DescribeMicroTenantsResponse) {
	response = &DescribeMicroTenantsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 微服务租户列表
func (c *Client) DescribeMicroTenants(request *DescribeMicroTenantsRequest) (response *DescribeMicroTenantsResponse, err error) {
	if request == nil {
		request = NewDescribeMicroTenantsRequest()
	}
	response = NewDescribeMicroTenantsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSystemRequest() (request *CreateSystemRequest) {
	request = &CreateSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateSystem")
	return
}

func NewCreateSystemResponse() (response *CreateSystemResponse) {
	response = &CreateSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建系统
func (c *Client) CreateSystem(request *CreateSystemRequest) (response *CreateSystemResponse, err error) {
	if request == nil {
		request = NewCreateSystemRequest()
	}
	response = NewCreateSystemResponse()
	err = c.Send(request, response)
	return
}

func NewAddVendorMemberByIDsRequest() (request *AddVendorMemberByIDsRequest) {
	request = &AddVendorMemberByIDsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "AddVendorMemberByIDs")
	return
}

func NewAddVendorMemberByIDsResponse() (response *AddVendorMemberByIDsResponse) {
	response = &AddVendorMemberByIDsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加供应商成员-通过用户ID列表批量添加
func (c *Client) AddVendorMemberByIDs(request *AddVendorMemberByIDsRequest) (response *AddVendorMemberByIDsResponse, err error) {
	if request == nil {
		request = NewAddVendorMemberByIDsRequest()
	}
	response = NewAddVendorMemberByIDsResponse()
	err = c.Send(request, response)
	return
}

func NewValidateSitePubPathRequest() (request *ValidateSitePubPathRequest) {
	request = &ValidateSitePubPathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ValidateSitePubPath")
	return
}

func NewValidateSitePubPathResponse() (response *ValidateSitePubPathResponse) {
	response = &ValidateSitePubPathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检测站点预览路径是否冲突
func (c *Client) ValidateSitePubPath(request *ValidateSitePubPathRequest) (response *ValidateSitePubPathResponse, err error) {
	if request == nil {
		request = NewValidateSitePubPathRequest()
	}
	response = NewValidateSitePubPathResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginPolicyRequest() (request *DescribeLoginPolicyRequest) {
	request = &DescribeLoginPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeLoginPolicy")
	return
}

func NewDescribeLoginPolicyResponse() (response *DescribeLoginPolicyResponse) {
	response = &DescribeLoginPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 登录安全策略查询
func (c *Client) DescribeLoginPolicy(request *DescribeLoginPolicyRequest) (response *DescribeLoginPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeLoginPolicyRequest()
	}
	response = NewDescribeLoginPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceApprovesRequest() (request *DescribeServiceApprovesRequest) {
	request = &DescribeServiceApprovesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeServiceApproves")
	return
}

func NewDescribeServiceApprovesResponse() (response *DescribeServiceApprovesResponse) {
	response = &DescribeServiceApprovesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务审批列表
func (c *Client) DescribeServiceApproves(request *DescribeServiceApprovesRequest) (response *DescribeServiceApprovesResponse, err error) {
	if request == nil {
		request = NewDescribeServiceApprovesRequest()
	}
	response = NewDescribeServiceApprovesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConnectorsRequest() (request *DeleteConnectorsRequest) {
	request = &DeleteConnectorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteConnectors")
	return
}

func NewDeleteConnectorsResponse() (response *DeleteConnectorsResponse) {
	response = &DeleteConnectorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 认证批量删除
func (c *Client) DeleteConnectors(request *DeleteConnectorsRequest) (response *DeleteConnectorsResponse, err error) {
	if request == nil {
		request = NewDeleteConnectorsRequest()
	}
	response = NewDeleteConnectorsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSmartGateRaidRoutingRequest() (request *DescribeSmartGateRaidRoutingRequest) {
	request = &DescribeSmartGateRaidRoutingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSmartGateRaidRouting")
	return
}

func NewDescribeSmartGateRaidRoutingResponse() (response *DescribeSmartGateRaidRoutingResponse) {
	response = &DescribeSmartGateRaidRoutingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点阵列路由详情
func (c *Client) DescribeSmartGateRaidRouting(request *DescribeSmartGateRaidRoutingRequest) (response *DescribeSmartGateRaidRoutingResponse, err error) {
	if request == nil {
		request = NewDescribeSmartGateRaidRoutingRequest()
	}
	response = NewDescribeSmartGateRaidRoutingResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecurityConfigRequest() (request *ModifySecurityConfigRequest) {
	request = &ModifySecurityConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifySecurityConfig")
	return
}

func NewModifySecurityConfigResponse() (response *ModifySecurityConfigResponse) {
	response = &ModifySecurityConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改安全设置
func (c *Client) ModifySecurityConfig(request *ModifySecurityConfigRequest) (response *ModifySecurityConfigResponse, err error) {
	if request == nil {
		request = NewModifySecurityConfigRequest()
	}
	response = NewModifySecurityConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCheckServiceUniqueRequest() (request *CheckServiceUniqueRequest) {
	request = &CheckServiceUniqueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CheckServiceUnique")
	return
}

func NewCheckServiceUniqueResponse() (response *CheckServiceUniqueResponse) {
	response = &CheckServiceUniqueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// name，id 唯一校验
func (c *Client) CheckServiceUnique(request *CheckServiceUniqueRequest) (response *CheckServiceUniqueResponse, err error) {
	if request == nil {
		request = NewCheckServiceUniqueRequest()
	}
	response = NewCheckServiceUniqueResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCategoryRequest() (request *ModifyCategoryRequest) {
	request = &ModifyCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyCategory")
	return
}

func NewModifyCategoryResponse() (response *ModifyCategoryResponse) {
	response = &ModifyCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改⽬录
func (c *Client) ModifyCategory(request *ModifyCategoryRequest) (response *ModifyCategoryResponse, err error) {
	if request == nil {
		request = NewModifyCategoryRequest()
	}
	response = NewModifyCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewUploadRequest() (request *UploadRequest) {
	request = &UploadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "Upload")
	return
}

func NewUploadResponse() (response *UploadResponse) {
	response = &UploadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通用上传
func (c *Client) Upload(request *UploadRequest) (response *UploadResponse, err error) {
	if request == nil {
		request = NewUploadRequest()
	}
	response = NewUploadResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSiteRequest() (request *DeleteSiteRequest) {
	request = &DeleteSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteSite")
	return
}

func NewDeleteSiteResponse() (response *DeleteSiteResponse) {
	response = &DeleteSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除站点
func (c *Client) DeleteSite(request *DeleteSiteRequest) (response *DeleteSiteResponse, err error) {
	if request == nil {
		request = NewDeleteSiteRequest()
	}
	response = NewDeleteSiteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceRequestLogPageRequest() (request *DescribeServiceRequestLogPageRequest) {
	request = &DescribeServiceRequestLogPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeServiceRequestLogPage")
	return
}

func NewDescribeServiceRequestLogPageResponse() (response *DescribeServiceRequestLogPageResponse) {
	response = &DescribeServiceRequestLogPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务请求日志列表
func (c *Client) DescribeServiceRequestLogPage(request *DescribeServiceRequestLogPageRequest) (response *DescribeServiceRequestLogPageResponse, err error) {
	if request == nil {
		request = NewDescribeServiceRequestLogPageRequest()
	}
	response = NewDescribeServiceRequestLogPageResponse()
	err = c.Send(request, response)
	return
}

func NewAddAgencyManagerRequest() (request *AddAgencyManagerRequest) {
	request = &AddAgencyManagerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "AddAgencyManager")
	return
}

func NewAddAgencyManagerResponse() (response *AddAgencyManagerResponse) {
	response = &AddAgencyManagerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 向机构负责人角色添加用户
func (c *Client) AddAgencyManager(request *AddAgencyManagerRequest) (response *AddAgencyManagerResponse, err error) {
	if request == nil {
		request = NewAddAgencyManagerRequest()
	}
	response = NewAddAgencyManagerResponse()
	err = c.Send(request, response)
	return
}

func NewModifyZoneRequest() (request *ModifyZoneRequest) {
	request = &ModifyZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyZone")
	return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
	response = &ModifyZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改分区
func (c *Client) ModifyZone(request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
	if request == nil {
		request = NewModifyZoneRequest()
	}
	response = NewModifyZoneResponse()
	err = c.Send(request, response)
	return
}

func NewRejectedSiteBatchRequest() (request *RejectedSiteBatchRequest) {
	request = &RejectedSiteBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RejectedSiteBatch")
	return
}

func NewRejectedSiteBatchResponse() (response *RejectedSiteBatchResponse) {
	response = &RejectedSiteBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 站点审批驳回批量
func (c *Client) RejectedSiteBatch(request *RejectedSiteBatchRequest) (response *RejectedSiteBatchResponse, err error) {
	if request == nil {
		request = NewRejectedSiteBatchRequest()
	}
	response = NewRejectedSiteBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSnapshotRequest() (request *DeleteSnapshotRequest) {
	request = &DeleteSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteSnapshot")
	return
}

func NewDeleteSnapshotResponse() (response *DeleteSnapshotResponse) {
	response = &DeleteSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除存档
func (c *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (response *DeleteSnapshotResponse, err error) {
	if request == nil {
		request = NewDeleteSnapshotRequest()
	}
	response = NewDeleteSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadMachineInfoRequest() (request *DownloadMachineInfoRequest) {
	request = &DownloadMachineInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadMachineInfo")
	return
}

func NewDownloadMachineInfoResponse() (response *DownloadMachineInfoResponse) {
	response = &DownloadMachineInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载机器部署节点信息
func (c *Client) DownloadMachineInfo(request *DownloadMachineInfoRequest) (response *DownloadMachineInfoResponse, err error) {
	if request == nil {
		request = NewDownloadMachineInfoRequest()
	}
	response = NewDownloadMachineInfoResponse()
	err = c.Send(request, response)
	return
}

func NewVerifySmsCodeAndModifyPhoneRequest() (request *VerifySmsCodeAndModifyPhoneRequest) {
	request = &VerifySmsCodeAndModifyPhoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "VerifySmsCodeAndModifyPhone")
	return
}

func NewVerifySmsCodeAndModifyPhoneResponse() (response *VerifySmsCodeAndModifyPhoneResponse) {
	response = &VerifySmsCodeAndModifyPhoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证密码并更新手机号
func (c *Client) VerifySmsCodeAndModifyPhone(request *VerifySmsCodeAndModifyPhoneRequest) (response *VerifySmsCodeAndModifyPhoneResponse, err error) {
	if request == nil {
		request = NewVerifySmsCodeAndModifyPhoneRequest()
	}
	response = NewVerifySmsCodeAndModifyPhoneResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAgencyConfigRequest() (request *DeleteAgencyConfigRequest) {
	request = &DeleteAgencyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteAgencyConfig")
	return
}

func NewDeleteAgencyConfigResponse() (response *DeleteAgencyConfigResponse) {
	response = &DeleteAgencyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机构流程设置
func (c *Client) DeleteAgencyConfig(request *DeleteAgencyConfigRequest) (response *DeleteAgencyConfigResponse, err error) {
	if request == nil {
		request = NewDeleteAgencyConfigRequest()
	}
	response = NewDeleteAgencyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResetTopMenuRequest() (request *ModifyResetTopMenuRequest) {
	request = &ModifyResetTopMenuRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyResetTopMenu")
	return
}

func NewModifyResetTopMenuResponse() (response *ModifyResetTopMenuResponse) {
	response = &ModifyResetTopMenuResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置顶部菜单
func (c *Client) ModifyResetTopMenu(request *ModifyResetTopMenuRequest) (response *ModifyResetTopMenuResponse, err error) {
	if request == nil {
		request = NewModifyResetTopMenuRequest()
	}
	response = NewModifyResetTopMenuResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAcsRoleRequest() (request *DescribeAcsRoleRequest) {
	request = &DescribeAcsRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAcsRole")
	return
}

func NewDescribeAcsRoleResponse() (response *DescribeAcsRoleResponse) {
	response = &DescribeAcsRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取角色详情
func (c *Client) DescribeAcsRole(request *DescribeAcsRoleRequest) (response *DescribeAcsRoleResponse, err error) {
	if request == nil {
		request = NewDescribeAcsRoleRequest()
	}
	response = NewDescribeAcsRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadImportFileRequest() (request *DownloadImportFileRequest) {
	request = &DownloadImportFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadImportFile")
	return
}

func NewDownloadImportFileResponse() (response *DownloadImportFileResponse) {
	response = &DownloadImportFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载机构文件导入成员模板
func (c *Client) DownloadImportFile(request *DownloadImportFileRequest) (response *DownloadImportFileResponse, err error) {
	if request == nil {
		request = NewDownloadImportFileRequest()
	}
	response = NewDownloadImportFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVendorProfileRequest() (request *DescribeVendorProfileRequest) {
	request = &DescribeVendorProfileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeVendorProfile")
	return
}

func NewDescribeVendorProfileResponse() (response *DescribeVendorProfileResponse) {
	response = &DescribeVendorProfileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户列表，携带是否是供应商成员和管理员字段
func (c *Client) DescribeVendorProfile(request *DescribeVendorProfileRequest) (response *DescribeVendorProfileResponse, err error) {
	if request == nil {
		request = NewDescribeVendorProfileRequest()
	}
	response = NewDescribeVendorProfileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePermissionRequest() (request *DescribePermissionRequest) {
	request = &DescribePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribePermission")
	return
}

func NewDescribePermissionResponse() (response *DescribePermissionResponse) {
	response = &DescribePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取权限列表树形结构
func (c *Client) DescribePermission(request *DescribePermissionRequest) (response *DescribePermissionResponse, err error) {
	if request == nil {
		request = NewDescribePermissionRequest()
	}
	response = NewDescribePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSiteApproveLogsRequest() (request *DescribeSiteApproveLogsRequest) {
	request = &DescribeSiteApproveLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSiteApproveLogs")
	return
}

func NewDescribeSiteApproveLogsResponse() (response *DescribeSiteApproveLogsResponse) {
	response = &DescribeSiteApproveLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 站点审批历史记录
func (c *Client) DescribeSiteApproveLogs(request *DescribeSiteApproveLogsRequest) (response *DescribeSiteApproveLogsResponse, err error) {
	if request == nil {
		request = NewDescribeSiteApproveLogsRequest()
	}
	response = NewDescribeSiteApproveLogsResponse()
	err = c.Send(request, response)
	return
}

func NewImportServiceRequest() (request *ImportServiceRequest) {
	request = &ImportServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ImportService")
	return
}

func NewImportServiceResponse() (response *ImportServiceResponse) {
	response = &ImportServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// swagger导入API服务
func (c *Client) ImportService(request *ImportServiceRequest) (response *ImportServiceResponse, err error) {
	if request == nil {
		request = NewImportServiceRequest()
	}
	response = NewImportServiceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResetPasswordRequest() (request *ModifyResetPasswordRequest) {
	request = &ModifyResetPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyResetPassword")
	return
}

func NewModifyResetPasswordResponse() (response *ModifyResetPasswordResponse) {
	response = &ModifyResetPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置帐密
func (c *Client) ModifyResetPassword(request *ModifyResetPasswordRequest) (response *ModifyResetPasswordResponse, err error) {
	if request == nil {
		request = NewModifyResetPasswordRequest()
	}
	response = NewModifyResetPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionsOfPubAppRequest() (request *DescribeSubscriptionsOfPubAppRequest) {
	request = &DescribeSubscriptionsOfPubAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSubscriptionsOfPubApp")
	return
}

func NewDescribeSubscriptionsOfPubAppResponse() (response *DescribeSubscriptionsOfPubAppResponse) {
	response = &DescribeSubscriptionsOfPubAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务授权记录
func (c *Client) DescribeSubscriptionsOfPubApp(request *DescribeSubscriptionsOfPubAppRequest) (response *DescribeSubscriptionsOfPubAppResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionsOfPubAppRequest()
	}
	response = NewDescribeSubscriptionsOfPubAppResponse()
	err = c.Send(request, response)
	return
}

func NewFreezeSiteRequest() (request *FreezeSiteRequest) {
	request = &FreezeSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "FreezeSite")
	return
}

func NewFreezeSiteResponse() (response *FreezeSiteResponse) {
	response = &FreezeSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 冻结站点
func (c *Client) FreezeSite(request *FreezeSiteRequest) (response *FreezeSiteResponse, err error) {
	if request == nil {
		request = NewFreezeSiteRequest()
	}
	response = NewFreezeSiteResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserAgenciesRequest() (request *GetUserAgenciesRequest) {
	request = &GetUserAgenciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "GetUserAgencies")
	return
}

func NewGetUserAgenciesResponse() (response *GetUserAgenciesResponse) {
	response = &GetUserAgenciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetUserAgencies
func (c *Client) GetUserAgencies(request *GetUserAgenciesRequest) (response *GetUserAgenciesResponse, err error) {
	if request == nil {
		request = NewGetUserAgenciesRequest()
	}
	response = NewGetUserAgenciesResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSiteUniqueRequest() (request *CheckSiteUniqueRequest) {
	request = &CheckSiteUniqueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CheckSiteUnique")
	return
}

func NewCheckSiteUniqueResponse() (response *CheckSiteUniqueResponse) {
	response = &CheckSiteUniqueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检测siteName，siteID唯一性
func (c *Client) CheckSiteUnique(request *CheckSiteUniqueRequest) (response *CheckSiteUniqueResponse, err error) {
	if request == nil {
		request = NewCheckSiteUniqueRequest()
	}
	response = NewCheckSiteUniqueResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSqlTemplateRequest() (request *DescribeSqlTemplateRequest) {
	request = &DescribeSqlTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSqlTemplate")
	return
}

func NewDescribeSqlTemplateResponse() (response *DescribeSqlTemplateResponse) {
	response = &DescribeSqlTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SQL 模板查询
func (c *Client) DescribeSqlTemplate(request *DescribeSqlTemplateRequest) (response *DescribeSqlTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeSqlTemplateRequest()
	}
	response = NewDescribeSqlTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTSFGatewayInstanceRequest() (request *CreateTSFGatewayInstanceRequest) {
	request = &CreateTSFGatewayInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateTSFGatewayInstance")
	return
}

func NewCreateTSFGatewayInstanceResponse() (response *CreateTSFGatewayInstanceResponse) {
	response = &CreateTSFGatewayInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步租户下 TSF 网关实例
func (c *Client) CreateTSFGatewayInstance(request *CreateTSFGatewayInstanceRequest) (response *CreateTSFGatewayInstanceResponse, err error) {
	if request == nil {
		request = NewCreateTSFGatewayInstanceRequest()
	}
	response = NewCreateTSFGatewayInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadServiceDataRequest() (request *DownloadServiceDataRequest) {
	request = &DownloadServiceDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadServiceData")
	return
}

func NewDownloadServiceDataResponse() (response *DownloadServiceDataResponse) {
	response = &DownloadServiceDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务数据
func (c *Client) DownloadServiceData(request *DownloadServiceDataRequest) (response *DownloadServiceDataResponse, err error) {
	if request == nil {
		request = NewDownloadServiceDataRequest()
	}
	response = NewDownloadServiceDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppTop10Request() (request *DescribeAppTop10Request) {
	request = &DescribeAppTop10Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAppTop10")
	return
}

func NewDescribeAppTop10Response() (response *DescribeAppTop10Response) {
	response = &DescribeAppTop10Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAppTop10
func (c *Client) DescribeAppTop10(request *DescribeAppTop10Request) (response *DescribeAppTop10Response, err error) {
	if request == nil {
		request = NewDescribeAppTop10Request()
	}
	response = NewDescribeAppTop10Response()
	err = c.Send(request, response)
	return
}

func NewDescribeCheckProfileRequest() (request *DescribeCheckProfileRequest) {
	request = &DescribeCheckProfileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeCheckProfile")
	return
}

func NewDescribeCheckProfileResponse() (response *DescribeCheckProfileResponse) {
	response = &DescribeCheckProfileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询profile校验字段
func (c *Client) DescribeCheckProfile(request *DescribeCheckProfileRequest) (response *DescribeCheckProfileResponse, err error) {
	if request == nil {
		request = NewDescribeCheckProfileRequest()
	}
	response = NewDescribeCheckProfileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
	request = &DescribeZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeZones")
	return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
	response = &DescribeZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分区列表
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
	if request == nil {
		request = NewDescribeZonesRequest()
	}
	response = NewDescribeZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppListByAgencyIdRequest() (request *DescribeAppListByAgencyIdRequest) {
	request = &DescribeAppListByAgencyIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAppListByAgencyId")
	return
}

func NewDescribeAppListByAgencyIdResponse() (response *DescribeAppListByAgencyIdResponse) {
	response = &DescribeAppListByAgencyIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据机构ID获取应用列表
func (c *Client) DescribeAppListByAgencyId(request *DescribeAppListByAgencyIdRequest) (response *DescribeAppListByAgencyIdResponse, err error) {
	if request == nil {
		request = NewDescribeAppListByAgencyIdRequest()
	}
	response = NewDescribeAppListByAgencyIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTSFGatewayApiGroupsRequest() (request *DescribeTSFGatewayApiGroupsRequest) {
	request = &DescribeTSFGatewayApiGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTSFGatewayApiGroups")
	return
}

func NewDescribeTSFGatewayApiGroupsResponse() (response *DescribeTSFGatewayApiGroupsResponse) {
	response = &DescribeTSFGatewayApiGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF 网关 API 分组列表
func (c *Client) DescribeTSFGatewayApiGroups(request *DescribeTSFGatewayApiGroupsRequest) (response *DescribeTSFGatewayApiGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeTSFGatewayApiGroupsRequest()
	}
	response = NewDescribeTSFGatewayApiGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewBindMemberRequest() (request *BindMemberRequest) {
	request = &BindMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "BindMember")
	return
}

func NewBindMemberResponse() (response *BindMemberResponse) {
	response = &BindMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 机构绑定已有成员
func (c *Client) BindMember(request *BindMemberRequest) (response *BindMemberResponse, err error) {
	if request == nil {
		request = NewBindMemberRequest()
	}
	response = NewBindMemberResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAreaZonesRequest() (request *DescribeAreaZonesRequest) {
	request = &DescribeAreaZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAreaZones")
	return
}

func NewDescribeAreaZonesResponse() (response *DescribeAreaZonesResponse) {
	response = &DescribeAreaZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 区域分区列表
func (c *Client) DescribeAreaZones(request *DescribeAreaZonesRequest) (response *DescribeAreaZonesResponse, err error) {
	if request == nil {
		request = NewDescribeAreaZonesRequest()
	}
	response = NewDescribeAreaZonesResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertSqlTemplateRequest() (request *UpsertSqlTemplateRequest) {
	request = &UpsertSqlTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UpsertSqlTemplate")
	return
}

func NewUpsertSqlTemplateResponse() (response *UpsertSqlTemplateResponse) {
	response = &UpsertSqlTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SQL 模板新增或修改
func (c *Client) UpsertSqlTemplate(request *UpsertSqlTemplateRequest) (response *UpsertSqlTemplateResponse, err error) {
	if request == nil {
		request = NewUpsertSqlTemplateRequest()
	}
	response = NewUpsertSqlTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSmartGateRaidRoutingsRequest() (request *DescribeSmartGateRaidRoutingsRequest) {
	request = &DescribeSmartGateRaidRoutingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSmartGateRaidRoutings")
	return
}

func NewDescribeSmartGateRaidRoutingsResponse() (response *DescribeSmartGateRaidRoutingsResponse) {
	response = &DescribeSmartGateRaidRoutingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点阵列路由列表
func (c *Client) DescribeSmartGateRaidRoutings(request *DescribeSmartGateRaidRoutingsRequest) (response *DescribeSmartGateRaidRoutingsResponse, err error) {
	if request == nil {
		request = NewDescribeSmartGateRaidRoutingsRequest()
	}
	response = NewDescribeSmartGateRaidRoutingsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionRequest() (request *DescribeSubscriptionRequest) {
	request = &DescribeSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSubscription")
	return
}

func NewDescribeSubscriptionResponse() (response *DescribeSubscriptionResponse) {
	response = &DescribeSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订阅详情
func (c *Client) DescribeSubscription(request *DescribeSubscriptionRequest) (response *DescribeSubscriptionResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionRequest()
	}
	response = NewDescribeSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadSystemDataRequest() (request *DownloadSystemDataRequest) {
	request = &DownloadSystemDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadSystemData")
	return
}

func NewDownloadSystemDataResponse() (response *DownloadSystemDataResponse) {
	response = &DownloadSystemDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统数据
func (c *Client) DownloadSystemData(request *DownloadSystemDataRequest) (response *DownloadSystemDataResponse, err error) {
	if request == nil {
		request = NewDownloadSystemDataRequest()
	}
	response = NewDownloadSystemDataResponse()
	err = c.Send(request, response)
	return
}

func NewRejectRenewSubscriptionRequest() (request *RejectRenewSubscriptionRequest) {
	request = &RejectRenewSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RejectRenewSubscription")
	return
}

func NewRejectRenewSubscriptionResponse() (response *RejectRenewSubscriptionResponse) {
	response = &RejectRenewSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续期订阅拒绝
func (c *Client) RejectRenewSubscription(request *RejectRenewSubscriptionRequest) (response *RejectRenewSubscriptionResponse, err error) {
	if request == nil {
		request = NewRejectRenewSubscriptionRequest()
	}
	response = NewRejectRenewSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppRequest() (request *DescribeAppRequest) {
	request = &DescribeAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeApp")
	return
}

func NewDescribeAppResponse() (response *DescribeAppResponse) {
	response = &DescribeAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 应用详情
func (c *Client) DescribeApp(request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
	if request == nil {
		request = NewDescribeAppRequest()
	}
	response = NewDescribeAppResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTlsCertificateRequest() (request *DescribeTlsCertificateRequest) {
	request = &DescribeTlsCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTlsCertificate")
	return
}

func NewDescribeTlsCertificateResponse() (response *DescribeTlsCertificateResponse) {
	response = &DescribeTlsCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询证书详情
func (c *Client) DescribeTlsCertificate(request *DescribeTlsCertificateRequest) (response *DescribeTlsCertificateResponse, err error) {
	if request == nil {
		request = NewDescribeTlsCertificateRequest()
	}
	response = NewDescribeTlsCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlertRuleRequest() (request *DescribeAlertRuleRequest) {
	request = &DescribeAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAlertRule")
	return
}

func NewDescribeAlertRuleResponse() (response *DescribeAlertRuleResponse) {
	response = &DescribeAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询告警规则
func (c *Client) DescribeAlertRule(request *DescribeAlertRuleRequest) (response *DescribeAlertRuleResponse, err error) {
	if request == nil {
		request = NewDescribeAlertRuleRequest()
	}
	response = NewDescribeAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConnectorsRequest() (request *DescribeConnectorsRequest) {
	request = &DescribeConnectorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeConnectors")
	return
}

func NewDescribeConnectorsResponse() (response *DescribeConnectorsResponse) {
	response = &DescribeConnectorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 认证列表
func (c *Client) DescribeConnectors(request *DescribeConnectorsRequest) (response *DescribeConnectorsResponse, err error) {
	if request == nil {
		request = NewDescribeConnectorsRequest()
	}
	response = NewDescribeConnectorsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginLogRequest() (request *DescribeLoginLogRequest) {
	request = &DescribeLoginLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeLoginLog")
	return
}

func NewDescribeLoginLogResponse() (response *DescribeLoginLogResponse) {
	response = &DescribeLoginLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询登录日志
func (c *Client) DescribeLoginLog(request *DescribeLoginLogRequest) (response *DescribeLoginLogResponse, err error) {
	if request == nil {
		request = NewDescribeLoginLogRequest()
	}
	response = NewDescribeLoginLogResponse()
	err = c.Send(request, response)
	return
}

func NewUploadLicenseRequest() (request *UploadLicenseRequest) {
	request = &UploadLicenseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UploadLicense")
	return
}

func NewUploadLicenseResponse() (response *UploadLicenseResponse) {
	response = &UploadLicenseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上传证书
func (c *Client) UploadLicense(request *UploadLicenseRequest) (response *UploadLicenseResponse, err error) {
	if request == nil {
		request = NewUploadLicenseRequest()
	}
	response = NewUploadLicenseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneAreasRequest() (request *DescribeZoneAreasRequest) {
	request = &DescribeZoneAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeZoneAreas")
	return
}

func NewDescribeZoneAreasResponse() (response *DescribeZoneAreasResponse) {
	response = &DescribeZoneAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分区区域列表
func (c *Client) DescribeZoneAreas(request *DescribeZoneAreasRequest) (response *DescribeZoneAreasResponse, err error) {
	if request == nil {
		request = NewDescribeZoneAreasRequest()
	}
	response = NewDescribeZoneAreasResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveMemberFromAcsRoleRequest() (request *RemoveMemberFromAcsRoleRequest) {
	request = &RemoveMemberFromAcsRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "RemoveMemberFromAcsRole")
	return
}

func NewRemoveMemberFromAcsRoleResponse() (response *RemoveMemberFromAcsRoleResponse) {
	response = &RemoveMemberFromAcsRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 角色删除成员
func (c *Client) RemoveMemberFromAcsRole(request *RemoveMemberFromAcsRoleRequest) (response *RemoveMemberFromAcsRoleResponse, err error) {
	if request == nil {
		request = NewRemoveMemberFromAcsRoleRequest()
	}
	response = NewRemoveMemberFromAcsRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemRequest() (request *DescribeSystemRequest) {
	request = &DescribeSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSystem")
	return
}

func NewDescribeSystemResponse() (response *DescribeSystemResponse) {
	response = &DescribeSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取系统详情
func (c *Client) DescribeSystem(request *DescribeSystemRequest) (response *DescribeSystemResponse, err error) {
	if request == nil {
		request = NewDescribeSystemRequest()
	}
	response = NewDescribeSystemResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSqlTemplateRequest() (request *DeleteSqlTemplateRequest) {
	request = &DeleteSqlTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DeleteSqlTemplate")
	return
}

func NewDeleteSqlTemplateResponse() (response *DeleteSqlTemplateResponse) {
	response = &DeleteSqlTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SQL 模板删除
func (c *Client) DeleteSqlTemplate(request *DeleteSqlTemplateRequest) (response *DeleteSqlTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteSqlTemplateRequest()
	}
	response = NewDeleteSqlTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTSFGatewayInstancesRequest() (request *DescribeTSFGatewayInstancesRequest) {
	request = &DescribeTSFGatewayInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTSFGatewayInstances")
	return
}

func NewDescribeTSFGatewayInstancesResponse() (response *DescribeTSFGatewayInstancesResponse) {
	response = &DescribeTSFGatewayInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF 网关实例列表
func (c *Client) DescribeTSFGatewayInstances(request *DescribeTSFGatewayInstancesRequest) (response *DescribeTSFGatewayInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeTSFGatewayInstancesRequest()
	}
	response = NewDescribeTSFGatewayInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgencyConfigRequest() (request *DescribeAgencyConfigRequest) {
	request = &DescribeAgencyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAgencyConfig")
	return
}

func NewDescribeAgencyConfigResponse() (response *DescribeAgencyConfigResponse) {
	response = &DescribeAgencyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机构流程设置
func (c *Client) DescribeAgencyConfig(request *DescribeAgencyConfigRequest) (response *DescribeAgencyConfigResponse, err error) {
	if request == nil {
		request = NewDescribeAgencyConfigRequest()
	}
	response = NewDescribeAgencyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSmartGatewayGroupsRequest() (request *DescribeSmartGatewayGroupsRequest) {
	request = &DescribeSmartGatewayGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSmartGatewayGroups")
	return
}

func NewDescribeSmartGatewayGroupsResponse() (response *DescribeSmartGatewayGroupsResponse) {
	response = &DescribeSmartGatewayGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点阵列列表
func (c *Client) DescribeSmartGatewayGroups(request *DescribeSmartGatewayGroupsRequest) (response *DescribeSmartGatewayGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeSmartGatewayGroupsRequest()
	}
	response = NewDescribeSmartGatewayGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTodoStatusRequest() (request *DescribeTodoStatusRequest) {
	request = &DescribeTodoStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeTodoStatus")
	return
}

func NewDescribeTodoStatusResponse() (response *DescribeTodoStatusResponse) {
	response = &DescribeTodoStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询是否有待办
func (c *Client) DescribeTodoStatus(request *DescribeTodoStatusRequest) (response *DescribeTodoStatusResponse, err error) {
	if request == nil {
		request = NewDescribeTodoStatusRequest()
	}
	response = NewDescribeTodoStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTipRequest() (request *ModifyTipRequest) {
	request = &ModifyTipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ModifyTip")
	return
}

func NewModifyTipResponse() (response *ModifyTipResponse) {
	response = &ModifyTipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改提示语
func (c *Client) ModifyTip(request *ModifyTipRequest) (response *ModifyTipResponse, err error) {
	if request == nil {
		request = NewModifyTipRequest()
	}
	response = NewModifyTipResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSnapshotRequest() (request *CreateSnapshotRequest) {
	request = &CreateSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateSnapshot")
	return
}

func NewCreateSnapshotResponse() (response *CreateSnapshotResponse) {
	response = &CreateSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建存档
func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
	if request == nil {
		request = NewCreateSnapshotRequest()
	}
	response = NewCreateSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionApproveLogsRequest() (request *DescribeSubscriptionApproveLogsRequest) {
	request = &DescribeSubscriptionApproveLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSubscriptionApproveLogs")
	return
}

func NewDescribeSubscriptionApproveLogsResponse() (response *DescribeSubscriptionApproveLogsResponse) {
	response = &DescribeSubscriptionApproveLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订阅审批历史记录
func (c *Client) DescribeSubscriptionApproveLogs(request *DescribeSubscriptionApproveLogsRequest) (response *DescribeSubscriptionApproveLogsResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionApproveLogsRequest()
	}
	response = NewDescribeSubscriptionApproveLogsResponse()
	err = c.Send(request, response)
	return
}

func NewUploadAccRequest() (request *UploadAccRequest) {
	request = &UploadAccRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "UploadAcc")
	return
}

func NewUploadAccResponse() (response *UploadAccResponse) {
	response = &UploadAccResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入用户
func (c *Client) UploadAcc(request *UploadAccRequest) (response *UploadAccResponse, err error) {
	if request == nil {
		request = NewUploadAccRequest()
	}
	response = NewUploadAccResponse()
	err = c.Send(request, response)
	return
}

func NewApprovedSiteRequest() (request *ApprovedSiteRequest) {
	request = &ApprovedSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ApprovedSite")
	return
}

func NewApprovedSiteResponse() (response *ApprovedSiteResponse) {
	response = &ApprovedSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 站点审批同意
func (c *Client) ApprovedSite(request *ApprovedSiteRequest) (response *ApprovedSiteResponse, err error) {
	if request == nil {
		request = NewApprovedSiteRequest()
	}
	response = NewApprovedSiteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAreaRequest() (request *DescribeAreaRequest) {
	request = &DescribeAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeArea")
	return
}

func NewDescribeAreaResponse() (response *DescribeAreaResponse) {
	response = &DescribeAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 区域详情
func (c *Client) DescribeArea(request *DescribeAreaRequest) (response *DescribeAreaResponse, err error) {
	if request == nil {
		request = NewDescribeAreaRequest()
	}
	response = NewDescribeAreaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLabelRequest() (request *DescribeLabelRequest) {
	request = &DescribeLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeLabel")
	return
}

func NewDescribeLabelResponse() (response *DescribeLabelResponse) {
	response = &DescribeLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取标签详情
func (c *Client) DescribeLabel(request *DescribeLabelRequest) (response *DescribeLabelResponse, err error) {
	if request == nil {
		request = NewDescribeLabelRequest()
	}
	response = NewDescribeLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionsOfSubAppRequest() (request *DescribeSubscriptionsOfSubAppRequest) {
	request = &DescribeSubscriptionsOfSubAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSubscriptionsOfSubApp")
	return
}

func NewDescribeSubscriptionsOfSubAppResponse() (response *DescribeSubscriptionsOfSubAppResponse) {
	response = &DescribeSubscriptionsOfSubAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 第三方服务订阅列表
func (c *Client) DescribeSubscriptionsOfSubApp(request *DescribeSubscriptionsOfSubAppRequest) (response *DescribeSubscriptionsOfSubAppResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionsOfSubAppRequest()
	}
	response = NewDescribeSubscriptionsOfSubAppResponse()
	err = c.Send(request, response)
	return
}

func NewApprovedSubscriptionRequest() (request *ApprovedSubscriptionRequest) {
	request = &ApprovedSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ApprovedSubscription")
	return
}

func NewApprovedSubscriptionResponse() (response *ApprovedSubscriptionResponse) {
	response = &ApprovedSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 审批订阅
func (c *Client) ApprovedSubscription(request *ApprovedSubscriptionRequest) (response *ApprovedSubscriptionResponse, err error) {
	if request == nil {
		request = NewApprovedSubscriptionRequest()
	}
	response = NewApprovedSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMembersRequest() (request *DescribeMembersRequest) {
	request = &DescribeMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMembers")
	return
}

func NewDescribeMembersResponse() (response *DescribeMembersResponse) {
	response = &DescribeMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机构成员
func (c *Client) DescribeMembers(request *DescribeMembersRequest) (response *DescribeMembersResponse, err error) {
	if request == nil {
		request = NewDescribeMembersRequest()
	}
	response = NewDescribeMembersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
	request = &CreateServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "CreateService")
	return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
	response = &CreateServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建API服务
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
	if request == nil {
		request = NewCreateServiceRequest()
	}
	response = NewCreateServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCategoryRequest() (request *DescribeCategoryRequest) {
	request = &DescribeCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeCategory")
	return
}

func NewDescribeCategoryResponse() (response *DescribeCategoryResponse) {
	response = &DescribeCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取⽬录详情
func (c *Client) DescribeCategory(request *DescribeCategoryRequest) (response *DescribeCategoryResponse, err error) {
	if request == nil {
		request = NewDescribeCategoryRequest()
	}
	response = NewDescribeCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroProtoFilesRequest() (request *DescribeMicroProtoFilesRequest) {
	request = &DescribeMicroProtoFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeMicroProtoFiles")
	return
}

func NewDescribeMicroProtoFilesResponse() (response *DescribeMicroProtoFilesResponse) {
	response = &DescribeMicroProtoFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// proto 文件列表
func (c *Client) DescribeMicroProtoFiles(request *DescribeMicroProtoFilesRequest) (response *DescribeMicroProtoFilesResponse, err error) {
	if request == nil {
		request = NewDescribeMicroProtoFilesRequest()
	}
	response = NewDescribeMicroProtoFilesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthInfoRequest() (request *DescribeAuthInfoRequest) {
	request = &DescribeAuthInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeAuthInfo")
	return
}

func NewDescribeAuthInfoResponse() (response *DescribeAuthInfoResponse) {
	response = &DescribeAuthInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询授权节点信息
func (c *Client) DescribeAuthInfo(request *DescribeAuthInfoRequest) (response *DescribeAuthInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAuthInfoRequest()
	}
	response = NewDescribeAuthInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionRenewApprovesRequest() (request *DescribeSubscriptionRenewApprovesRequest) {
	request = &DescribeSubscriptionRenewApprovesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeSubscriptionRenewApproves")
	return
}

func NewDescribeSubscriptionRenewApprovesResponse() (response *DescribeSubscriptionRenewApprovesResponse) {
	response = &DescribeSubscriptionRenewApprovesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeSubscriptionRenewApproves
func (c *Client) DescribeSubscriptionRenewApproves(request *DescribeSubscriptionRenewApprovesRequest) (response *DescribeSubscriptionRenewApprovesResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionRenewApprovesRequest()
	}
	response = NewDescribeSubscriptionRenewApprovesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCategorysRequest() (request *DescribeCategorysRequest) {
	request = &DescribeCategorysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DescribeCategorys")
	return
}

func NewDescribeCategorysResponse() (response *DescribeCategorysResponse) {
	response = &DescribeCategorysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取⽬录列表
func (c *Client) DescribeCategorys(request *DescribeCategorysRequest) (response *DescribeCategorysResponse, err error) {
	if request == nil {
		request = NewDescribeCategorysRequest()
	}
	response = NewDescribeCategorysResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadAppDataRequest() (request *DownloadAppDataRequest) {
	request = &DownloadAppDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "DownloadAppData")
	return
}

func NewDownloadAppDataResponse() (response *DownloadAppDataResponse) {
	response = &DownloadAppDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 应用数据
func (c *Client) DownloadAppData(request *DownloadAppDataRequest) (response *DownloadAppDataResponse, err error) {
	if request == nil {
		request = NewDownloadAppDataRequest()
	}
	response = NewDownloadAppDataResponse()
	err = c.Send(request, response)
	return
}

func NewImportVendorMemberRequest() (request *ImportVendorMemberRequest) {
	request = &ImportVendorMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("rio", APIVersion, "ImportVendorMember")
	return
}

func NewImportVendorMemberResponse() (response *ImportVendorMemberResponse) {
	response = &ImportVendorMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入供应商成员
func (c *Client) ImportVendorMember(request *ImportVendorMemberRequest) (response *ImportVendorMemberResponse, err error) {
	if request == nil {
		request = NewImportVendorMemberRequest()
	}
	response = NewImportVendorMemberResponse()
	err = c.Send(request, response)
	return
}
