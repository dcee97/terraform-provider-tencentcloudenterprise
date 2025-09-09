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

func NewDeleteAccessWhiteListRulesRequest() (request *DeleteAccessWhiteListRulesRequest) {
	request = &DeleteAccessWhiteListRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteAccessWhiteListRules")
	return
}

func NewDeleteAccessWhiteListRulesResponse() (response *DeleteAccessWhiteListRulesResponse) {
	response = &DeleteAccessWhiteListRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问白名单规则
func (c *Client) DeleteAccessWhiteListRules(request *DeleteAccessWhiteListRulesRequest) (response *DeleteAccessWhiteListRulesResponse, err error) {
	if request == nil {
		request = NewDeleteAccessWhiteListRulesRequest()
	}
	response = NewDeleteAccessWhiteListRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCmdTemplatesRequest() (request *DeleteCmdTemplatesRequest) {
	request = &DeleteCmdTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteCmdTemplates")
	return
}

func NewDeleteCmdTemplatesResponse() (response *DeleteCmdTemplatesResponse) {
	response = &DeleteCmdTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除高危命令模板
func (c *Client) DeleteCmdTemplates(request *DeleteCmdTemplatesRequest) (response *DeleteCmdTemplatesResponse, err error) {
	if request == nil {
		request = NewDeleteCmdTemplatesRequest()
	}
	response = NewDeleteCmdTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessWhiteListRuleRequest() (request *ModifyAccessWhiteListRuleRequest) {
	request = &ModifyAccessWhiteListRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessWhiteListRule")
	return
}

func NewModifyAccessWhiteListRuleResponse() (response *ModifyAccessWhiteListRuleResponse) {
	response = &ModifyAccessWhiteListRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问白名单规则
func (c *Client) ModifyAccessWhiteListRule(request *ModifyAccessWhiteListRuleRequest) (response *ModifyAccessWhiteListRuleResponse, err error) {
	if request == nil {
		request = NewModifyAccessWhiteListRuleRequest()
	}
	response = NewModifyAccessWhiteListRuleResponse()
	err = c.Send(request, response)
	return
}

func NewSearchAuditLogRequest() (request *SearchAuditLogRequest) {
	request = &SearchAuditLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchAuditLog")
	return
}

func NewSearchAuditLogResponse() (response *SearchAuditLogResponse) {
	response = &SearchAuditLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索审计日志
func (c *Client) SearchAuditLog(request *SearchAuditLogRequest) (response *SearchAuditLogResponse, err error) {
	if request == nil {
		request = NewSearchAuditLogRequest()
	}
	response = NewSearchAuditLogResponse()
	err = c.Send(request, response)
	return
}

func NewAddUserGroupMembersRequest() (request *AddUserGroupMembersRequest) {
	request = &AddUserGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "AddUserGroupMembers")
	return
}

func NewAddUserGroupMembersResponse() (response *AddUserGroupMembersResponse) {
	response = &AddUserGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加用户组成员
func (c *Client) AddUserGroupMembers(request *AddUserGroupMembersRequest) (response *AddUserGroupMembersResponse, err error) {
	if request == nil {
		request = NewAddUserGroupMembersRequest()
	}
	response = NewAddUserGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewReplaySessionRequest() (request *ReplaySessionRequest) {
	request = &ReplaySessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ReplaySession")
	return
}

func NewReplaySessionResponse() (response *ReplaySessionResponse) {
	response = &ReplaySessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 会话回放
func (c *Client) ReplaySession(request *ReplaySessionRequest) (response *ReplaySessionResponse, err error) {
	if request == nil {
		request = NewReplaySessionRequest()
	}
	response = NewReplaySessionResponse()
	err = c.Send(request, response)
	return
}

func NewBindDeviceAccountPrivateKeyRequest() (request *BindDeviceAccountPrivateKeyRequest) {
	request = &BindDeviceAccountPrivateKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "BindDeviceAccountPrivateKey")
	return
}

func NewBindDeviceAccountPrivateKeyResponse() (response *BindDeviceAccountPrivateKeyResponse) {
	response = &BindDeviceAccountPrivateKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定主机账号私钥
func (c *Client) BindDeviceAccountPrivateKey(request *BindDeviceAccountPrivateKeyRequest) (response *BindDeviceAccountPrivateKeyResponse, err error) {
	if request == nil {
		request = NewBindDeviceAccountPrivateKeyRequest()
	}
	response = NewBindDeviceAccountPrivateKeyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetSyncFlagRequest() (request *ModifyAssetSyncFlagRequest) {
	request = &ModifyAssetSyncFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAssetSyncFlag")
	return
}

func NewModifyAssetSyncFlagResponse() (response *ModifyAssetSyncFlagResponse) {
	response = &ModifyAssetSyncFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资产自动同步开关
func (c *Client) ModifyAssetSyncFlag(request *ModifyAssetSyncFlagRequest) (response *ModifyAssetSyncFlagResponse, err error) {
	if request == nil {
		request = NewModifyAssetSyncFlagRequest()
	}
	response = NewModifyAssetSyncFlagResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLDAPSettingRequest() (request *ModifyLDAPSettingRequest) {
	request = &ModifyLDAPSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyLDAPSetting")
	return
}

func NewModifyLDAPSettingResponse() (response *ModifyLDAPSettingResponse) {
	response = &ModifyLDAPSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改LDAP配置信息
func (c *Client) ModifyLDAPSetting(request *ModifyLDAPSettingRequest) (response *ModifyLDAPSettingResponse, err error) {
	if request == nil {
		request = NewModifyLDAPSettingRequest()
	}
	response = NewModifyLDAPSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountsWithDeviceCountRequest() (request *DescribeAccountsWithDeviceCountRequest) {
	request = &DescribeAccountsWithDeviceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccountsWithDeviceCount")
	return
}

func NewDescribeAccountsWithDeviceCountResponse() (response *DescribeAccountsWithDeviceCountResponse) {
	response = &DescribeAccountsWithDeviceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账号,带拥有相同账号的主机数，可以按照主机、主机组查询，不分页
func (c *Client) DescribeAccountsWithDeviceCount(request *DescribeAccountsWithDeviceCountRequest) (response *DescribeAccountsWithDeviceCountResponse, err error) {
	if request == nil {
		request = NewDescribeAccountsWithDeviceCountRequest()
	}
	response = NewDescribeAccountsWithDeviceCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSyncStatusRequest() (request *DescribeAssetSyncStatusRequest) {
	request = &DescribeAssetSyncStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAssetSyncStatus")
	return
}

func NewDescribeAssetSyncStatusResponse() (response *DescribeAssetSyncStatusResponse) {
	response = &DescribeAssetSyncStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产同步状态
func (c *Client) DescribeAssetSyncStatus(request *DescribeAssetSyncStatusRequest) (response *DescribeAssetSyncStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSyncStatusRequest()
	}
	response = NewDescribeAssetSyncStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDeviceGroupRequest() (request *ModifyDeviceGroupRequest) {
	request = &ModifyDeviceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDeviceGroup")
	return
}

func NewModifyDeviceGroupResponse() (response *ModifyDeviceGroupResponse) {
	response = &ModifyDeviceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资产组
func (c *Client) ModifyDeviceGroup(request *ModifyDeviceGroupRequest) (response *ModifyDeviceGroupResponse, err error) {
	if request == nil {
		request = NewModifyDeviceGroupRequest()
	}
	response = NewModifyDeviceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepartmentsRequest() (request *DescribeDepartmentsRequest) {
	request = &DescribeDepartmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDepartments")
	return
}

func NewDescribeDepartmentsResponse() (response *DescribeDepartmentsResponse) {
	response = &DescribeDepartmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询部门信息
func (c *Client) DescribeDepartments(request *DescribeDepartmentsRequest) (response *DescribeDepartmentsResponse, err error) {
	if request == nil {
		request = NewDescribeDepartmentsRequest()
	}
	response = NewDescribeDepartmentsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
	request = &DeleteUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteUsers")
	return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
	response = &DeleteUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
	if request == nil {
		request = NewDeleteUsersRequest()
	}
	response = NewDeleteUsersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAclsRequest() (request *DescribeAclsRequest) {
	request = &DescribeAclsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAcls")
	return
}

func NewDescribeAclsResponse() (response *DescribeAclsResponse) {
	response = &DescribeAclsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询访问权限列表
func (c *Client) DescribeAcls(request *DescribeAclsRequest) (response *DescribeAclsResponse, err error) {
	if request == nil {
		request = NewDescribeAclsRequest()
	}
	response = NewDescribeAclsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
	request = &ModifyUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyUser")
	return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
	response = &ModifyUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户信息
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
	if request == nil {
		request = NewModifyUserRequest()
	}
	response = NewModifyUserResponse()
	err = c.Send(request, response)
	return
}

func NewResetDeviceAccountPasswordRequest() (request *ResetDeviceAccountPasswordRequest) {
	request = &ResetDeviceAccountPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ResetDeviceAccountPassword")
	return
}

func NewResetDeviceAccountPasswordResponse() (response *ResetDeviceAccountPasswordResponse) {
	response = &ResetDeviceAccountPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清除设备账号绑定密码
func (c *Client) ResetDeviceAccountPassword(request *ResetDeviceAccountPasswordRequest) (response *ResetDeviceAccountPasswordResponse, err error) {
	if request == nil {
		request = NewResetDeviceAccountPasswordRequest()
	}
	response = NewResetDeviceAccountPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogOutputSettingsRequest() (request *DescribeLogOutputSettingsRequest) {
	request = &DescribeLogOutputSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeLogOutputSettings")
	return
}

func NewDescribeLogOutputSettingsResponse() (response *DescribeLogOutputSettingsResponse) {
	response = &DescribeLogOutputSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志外发配置
func (c *Client) DescribeLogOutputSettings(request *DescribeLogOutputSettingsRequest) (response *DescribeLogOutputSettingsResponse, err error) {
	if request == nil {
		request = NewDescribeLogOutputSettingsRequest()
	}
	response = NewDescribeLogOutputSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAclRequest() (request *CreateAclRequest) {
	request = &CreateAclRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAcl")
	return
}

func NewCreateAclResponse() (response *CreateAclResponse) {
	response = &CreateAclResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建访问权限
func (c *Client) CreateAcl(request *CreateAclRequest) (response *CreateAclResponse, err error) {
	if request == nil {
		request = NewCreateAclRequest()
	}
	response = NewCreateAclResponse()
	err = c.Send(request, response)
	return
}

func NewSearchEventRequest() (request *SearchEventRequest) {
	request = &SearchEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchEvent")
	return
}

func NewSearchEventResponse() (response *SearchEventResponse) {
	response = &SearchEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索高危事件
func (c *Client) SearchEvent(request *SearchEventRequest) (response *SearchEventResponse, err error) {
	if request == nil {
		request = NewSearchEventRequest()
	}
	response = NewSearchEventResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAclsRequest() (request *DeleteAclsRequest) {
	request = &DeleteAclsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteAcls")
	return
}

func NewDeleteAclsResponse() (response *DeleteAclsResponse) {
	response = &DeleteAclsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问权限
func (c *Client) DeleteAcls(request *DeleteAclsRequest) (response *DeleteAclsResponse, err error) {
	if request == nil {
		request = NewDeleteAclsRequest()
	}
	response = NewDeleteAclsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
	request = &DescribeResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeResources")
	return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
	response = &DescribeResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户购买的堡垒机服务信息，包括资源ID、授权点数、VPC、过期时间等。
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeResourcesRequest()
	}
	response = NewDescribeResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessWhiteListRuleRequest() (request *CreateAccessWhiteListRuleRequest) {
	request = &CreateAccessWhiteListRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAccessWhiteListRule")
	return
}

func NewCreateAccessWhiteListRuleResponse() (response *CreateAccessWhiteListRuleResponse) {
	response = &CreateAccessWhiteListRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加访问白名单规则
func (c *Client) CreateAccessWhiteListRule(request *CreateAccessWhiteListRuleRequest) (response *CreateAccessWhiteListRuleResponse, err error) {
	if request == nil {
		request = NewCreateAccessWhiteListRuleRequest()
	}
	response = NewCreateAccessWhiteListRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDeviceRequest() (request *ModifyDeviceRequest) {
	request = &ModifyDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDevice")
	return
}

func NewModifyDeviceResponse() (response *ModifyDeviceResponse) {
	response = &ModifyDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资产信息
func (c *Client) ModifyDevice(request *ModifyDeviceRequest) (response *ModifyDeviceResponse, err error) {
	if request == nil {
		request = NewModifyDeviceRequest()
	}
	response = NewModifyDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDeviceGroupMembersRequest() (request *DeleteDeviceGroupMembersRequest) {
	request = &DeleteDeviceGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDeviceGroupMembers")
	return
}

func NewDeleteDeviceGroupMembersResponse() (response *DeleteDeviceGroupMembersResponse) {
	response = &DeleteDeviceGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资产组成员
func (c *Client) DeleteDeviceGroupMembers(request *DeleteDeviceGroupMembersRequest) (response *DeleteDeviceGroupMembersResponse, err error) {
	if request == nil {
		request = NewDeleteDeviceGroupMembersRequest()
	}
	response = NewDeleteDeviceGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDevicesDepartmentRequest() (request *ModifyDevicesDepartmentRequest) {
	request = &ModifyDevicesDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDevicesDepartment")
	return
}

func NewModifyDevicesDepartmentResponse() (response *ModifyDevicesDepartmentResponse) {
	response = &ModifyDevicesDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改资产部门
func (c *Client) ModifyDevicesDepartment(request *ModifyDevicesDepartmentRequest) (response *ModifyDevicesDepartmentResponse, err error) {
	if request == nil {
		request = NewModifyDevicesDepartmentRequest()
	}
	response = NewModifyDevicesDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewCreatDasbResourceRequest() (request *CreatDasbResourceRequest) {
	request = &CreatDasbResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreatDasbResource")
	return
}

func NewCreatDasbResourceResponse() (response *CreatDasbResourceResponse) {
	response = &CreatDasbResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建堡垒机资源
func (c *Client) CreatDasbResource(request *CreatDasbResourceRequest) (response *CreatDasbResourceResponse, err error) {
	if request == nil {
		request = NewCreatDasbResourceRequest()
	}
	response = NewCreatDasbResourceResponse()
	err = c.Send(request, response)
	return
}

func NewSetLDAPSyncFlagRequest() (request *SetLDAPSyncFlagRequest) {
	request = &SetLDAPSyncFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SetLDAPSyncFlag")
	return
}

func NewSetLDAPSyncFlagResponse() (response *SetLDAPSyncFlagResponse) {
	response = &SetLDAPSyncFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置LDAP 立即同步标记
func (c *Client) SetLDAPSyncFlag(request *SetLDAPSyncFlagRequest) (response *SetLDAPSyncFlagResponse, err error) {
	if request == nil {
		request = NewSetLDAPSyncFlagRequest()
	}
	response = NewSetLDAPSyncFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserGroupMembersRequest() (request *DescribeUserGroupMembersRequest) {
	request = &DescribeUserGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeUserGroupMembers")
	return
}

func NewDescribeUserGroupMembersResponse() (response *DescribeUserGroupMembersResponse) {
	response = &DescribeUserGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户组成员列表
func (c *Client) DescribeUserGroupMembers(request *DescribeUserGroupMembersRequest) (response *DescribeUserGroupMembersResponse, err error) {
	if request == nil {
		request = NewDescribeUserGroupMembersRequest()
	}
	response = NewDescribeUserGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewMonitorSessionRequest() (request *MonitorSessionRequest) {
	request = &MonitorSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "MonitorSession")
	return
}

func NewMonitorSessionResponse() (response *MonitorSessionResponse) {
	response = &MonitorSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 会话实时监控
func (c *Client) MonitorSession(request *MonitorSessionRequest) (response *MonitorSessionResponse, err error) {
	if request == nil {
		request = NewMonitorSessionRequest()
	}
	response = NewMonitorSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationEventRequest() (request *DescribeOperationEventRequest) {
	request = &DescribeOperationEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationEvent")
	return
}

func NewDescribeOperationEventResponse() (response *DescribeOperationEventResponse) {
	response = &DescribeOperationEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作日志
func (c *Client) DescribeOperationEvent(request *DescribeOperationEventRequest) (response *DescribeOperationEventResponse, err error) {
	if request == nil {
		request = NewDescribeOperationEventRequest()
	}
	response = NewDescribeOperationEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationTypeRequest() (request *DescribeOperationTypeRequest) {
	request = &DescribeOperationTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationType")
	return
}

func NewDescribeOperationTypeResponse() (response *DescribeOperationTypeResponse) {
	response = &DescribeOperationTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作类型
func (c *Client) DescribeOperationType(request *DescribeOperationTypeRequest) (response *DescribeOperationTypeResponse, err error) {
	if request == nil {
		request = NewDescribeOperationTypeRequest()
	}
	response = NewDescribeOperationTypeResponse()
	err = c.Send(request, response)
	return
}

func NewInquireCreateDasbResourceRequest() (request *InquireCreateDasbResourceRequest) {
	request = &InquireCreateDasbResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "InquireCreateDasbResource")
	return
}

func NewInquireCreateDasbResourceResponse() (response *InquireCreateDasbResourceResponse) {
	response = &InquireCreateDasbResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// tce后付费模式下,【新购堡垒机询价】和【变配询价】都是使用该接口。
func (c *Client) InquireCreateDasbResource(request *InquireCreateDasbResourceRequest) (response *InquireCreateDasbResourceResponse, err error) {
	if request == nil {
		request = NewInquireCreateDasbResourceRequest()
	}
	response = NewInquireCreateDasbResourceResponse()
	err = c.Send(request, response)
	return
}

func NewResetUserPasswordRequest() (request *ResetUserPasswordRequest) {
	request = &ResetUserPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ResetUserPassword")
	return
}

func NewResetUserPasswordResponse() (response *ResetUserPasswordResponse) {
	response = &ResetUserPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置用户密码
func (c *Client) ResetUserPassword(request *ResetUserPasswordRequest) (response *ResetUserPasswordResponse, err error) {
	if request == nil {
		request = NewResetUserPasswordRequest()
	}
	response = NewResetUserPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSessionRequest() (request *SearchSessionRequest) {
	request = &SearchSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchSession")
	return
}

func NewSearchSessionResponse() (response *SearchSessionResponse) {
	response = &SearchSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索会话
func (c *Client) SearchSession(request *SearchSessionRequest) (response *SearchSessionResponse, err error) {
	if request == nil {
		request = NewSearchSessionRequest()
	}
	response = NewSearchSessionResponse()
	err = c.Send(request, response)
	return
}

func NewImportExternalDeviceRequest() (request *ImportExternalDeviceRequest) {
	request = &ImportExternalDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ImportExternalDevice")
	return
}

func NewImportExternalDeviceResponse() (response *ImportExternalDeviceResponse) {
	response = &ImportExternalDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入外部资产信息
func (c *Client) ImportExternalDevice(request *ImportExternalDeviceRequest) (response *ImportExternalDeviceResponse, err error) {
	if request == nil {
		request = NewImportExternalDeviceRequest()
	}
	response = NewImportExternalDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewSearchFileBySidRequest() (request *SearchFileBySidRequest) {
	request = &SearchFileBySidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchFileBySid")
	return
}

func NewSearchFileBySidResponse() (response *SearchFileBySidResponse) {
	response = &SearchFileBySidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索文件传输会话下文件操作列表
func (c *Client) SearchFileBySid(request *SearchFileBySidRequest) (response *SearchFileBySidResponse, err error) {
	if request == nil {
		request = NewSearchFileBySidRequest()
	}
	response = NewSearchFileBySidResponse()
	err = c.Send(request, response)
	return
}

func NewCheckLDAPConnectionRequest() (request *CheckLDAPConnectionRequest) {
	request = &CheckLDAPConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CheckLDAPConnection")
	return
}

func NewCheckLDAPConnectionResponse() (response *CheckLDAPConnectionResponse) {
	response = &CheckLDAPConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试LDAP连接
func (c *Client) CheckLDAPConnection(request *CheckLDAPConnectionRequest) (response *CheckLDAPConnectionResponse, err error) {
	if request == nil {
		request = NewCheckLDAPConnectionRequest()
	}
	response = NewCheckLDAPConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserGroupsRequest() (request *DescribeUserGroupsRequest) {
	request = &DescribeUserGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeUserGroups")
	return
}

func NewDescribeUserGroupsResponse() (response *DescribeUserGroupsResponse) {
	response = &DescribeUserGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户组列表
func (c *Client) DescribeUserGroups(request *DescribeUserGroupsRequest) (response *DescribeUserGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeUserGroupsRequest()
	}
	response = NewDescribeUserGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceIdsRequest() (request *DescribeInstanceIdsRequest) {
	request = &DescribeInstanceIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeInstanceIds")
	return
}

func NewDescribeInstanceIdsResponse() (response *DescribeInstanceIdsResponse) {
	response = &DescribeInstanceIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定区域所有主机的InstanceId，不分页
func (c *Client) DescribeInstanceIds(request *DescribeInstanceIdsRequest) (response *DescribeInstanceIdsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceIdsRequest()
	}
	response = NewDescribeInstanceIdsResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyDasbResourceRequest() (request *DestroyDasbResourceRequest) {
	request = &DestroyDasbResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DestroyDasbResource")
	return
}

func NewDestroyDasbResourceResponse() (response *DestroyDasbResourceResponse) {
	response = &DestroyDasbResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁堡垒机资源接口
func (c *Client) DestroyDasbResource(request *DestroyDasbResourceRequest) (response *DestroyDasbResourceResponse, err error) {
	if request == nil {
		request = NewDestroyDasbResourceRequest()
	}
	response = NewDestroyDasbResourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessWhiteListStatusRequest() (request *ModifyAccessWhiteListStatusRequest) {
	request = &ModifyAccessWhiteListStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessWhiteListStatus")
	return
}

func NewModifyAccessWhiteListStatusResponse() (response *ModifyAccessWhiteListStatusResponse) {
	response = &ModifyAccessWhiteListStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问白名单状态：开启或关闭放开全部来源IP。
func (c *Client) ModifyAccessWhiteListStatus(request *ModifyAccessWhiteListStatusRequest) (response *ModifyAccessWhiteListStatusResponse, err error) {
	if request == nil {
		request = NewModifyAccessWhiteListStatusRequest()
	}
	response = NewModifyAccessWhiteListStatusResponse()
	err = c.Send(request, response)
	return
}

func NewSearchFileRequest() (request *SearchFileRequest) {
	request = &SearchFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchFile")
	return
}

func NewSearchFileResponse() (response *SearchFileResponse) {
	response = &SearchFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索文件传输操作（用于控制台审计管理->操作检索->文件传输）
func (c *Client) SearchFile(request *SearchFileRequest) (response *SearchFileResponse, err error) {
	if request == nil {
		request = NewSearchFileRequest()
	}
	response = NewSearchFileResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetSyncJobRequest() (request *CreateAssetSyncJobRequest) {
	request = &CreateAssetSyncJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAssetSyncJob")
	return
}

func NewCreateAssetSyncJobResponse() (response *CreateAssetSyncJobResponse) {
	response = &CreateAssetSyncJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建手工资产同步任务
func (c *Client) CreateAssetSyncJob(request *CreateAssetSyncJobRequest) (response *CreateAssetSyncJobResponse, err error) {
	if request == nil {
		request = NewCreateAssetSyncJobRequest()
	}
	response = NewCreateAssetSyncJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserBatchRequest() (request *CreateUserBatchRequest) {
	request = &CreateUserBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateUserBatch")
	return
}

func NewCreateUserBatchResponse() (response *CreateUserBatchResponse) {
	response = &CreateUserBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入用户
func (c *Client) CreateUserBatch(request *CreateUserBatchRequest) (response *CreateUserBatchResponse, err error) {
	if request == nil {
		request = NewCreateUserBatchRequest()
	}
	response = NewCreateUserBatchResponse()
	err = c.Send(request, response)
	return
}

func NewBindDeviceResourceRequest() (request *BindDeviceResourceRequest) {
	request = &BindDeviceResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "BindDeviceResource")
	return
}

func NewBindDeviceResourceResponse() (response *BindDeviceResourceResponse) {
	response = &BindDeviceResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资产绑定的堡垒机服务
func (c *Client) BindDeviceResource(request *BindDeviceResourceRequest) (response *BindDeviceResourceResponse, err error) {
	if request == nil {
		request = NewBindDeviceResourceRequest()
	}
	response = NewBindDeviceResourceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDepartmentRequest() (request *CreateDepartmentRequest) {
	request = &CreateDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateDepartment")
	return
}

func NewCreateDepartmentResponse() (response *CreateDepartmentResponse) {
	response = &CreateDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建部门
func (c *Client) CreateDepartment(request *CreateDepartmentRequest) (response *CreateDepartmentResponse, err error) {
	if request == nil {
		request = NewCreateDepartmentRequest()
	}
	response = NewCreateDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginEventRequest() (request *DescribeLoginEventRequest) {
	request = &DescribeLoginEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeLoginEvent")
	return
}

func NewDescribeLoginEventResponse() (response *DescribeLoginEventResponse) {
	response = &DescribeLoginEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询登录日志
func (c *Client) DescribeLoginEvent(request *DescribeLoginEventRequest) (response *DescribeLoginEventResponse, err error) {
	if request == nil {
		request = NewDescribeLoginEventRequest()
	}
	response = NewDescribeLoginEventResponse()
	err = c.Send(request, response)
	return
}

func NewResetDeviceAccountPrivateKeyRequest() (request *ResetDeviceAccountPrivateKeyRequest) {
	request = &ResetDeviceAccountPrivateKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ResetDeviceAccountPrivateKey")
	return
}

func NewResetDeviceAccountPrivateKeyResponse() (response *ResetDeviceAccountPrivateKeyResponse) {
	response = &ResetDeviceAccountPrivateKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清除设备账号绑定的密钥
func (c *Client) ResetDeviceAccountPrivateKey(request *ResetDeviceAccountPrivateKeyRequest) (response *ResetDeviceAccountPrivateKeyResponse, err error) {
	if request == nil {
		request = NewResetDeviceAccountPrivateKeyRequest()
	}
	response = NewResetDeviceAccountPrivateKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCmdTemplatesRequest() (request *DescribeCmdTemplatesRequest) {
	request = &DescribeCmdTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeCmdTemplates")
	return
}

func NewDescribeCmdTemplatesResponse() (response *DescribeCmdTemplatesResponse) {
	response = &DescribeCmdTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询命令模板列表
func (c *Client) DescribeCmdTemplates(request *DescribeCmdTemplatesRequest) (response *DescribeCmdTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeCmdTemplatesRequest()
	}
	response = NewDescribeCmdTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserGroupRequest() (request *ModifyUserGroupRequest) {
	request = &ModifyUserGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyUserGroup")
	return
}

func NewModifyUserGroupResponse() (response *ModifyUserGroupResponse) {
	response = &ModifyUserGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户组
func (c *Client) ModifyUserGroup(request *ModifyUserGroupRequest) (response *ModifyUserGroupResponse, err error) {
	if request == nil {
		request = NewModifyUserGroupRequest()
	}
	response = NewModifyUserGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceCountRequest() (request *DescribeDeviceCountRequest) {
	request = &DescribeDeviceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDeviceCount")
	return
}

func NewDescribeDeviceCountResponse() (response *DescribeDeviceCountResponse) {
	response = &DescribeDeviceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户导入的主机数
func (c *Client) DescribeDeviceCount(request *DescribeDeviceCountRequest) (response *DescribeDeviceCountResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceCountRequest()
	}
	response = NewDescribeDeviceCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDepartmentRequest() (request *ModifyDepartmentRequest) {
	request = &ModifyDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDepartment")
	return
}

func NewModifyDepartmentResponse() (response *ModifyDepartmentResponse) {
	response = &ModifyDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改部门信息
func (c *Client) ModifyDepartment(request *ModifyDepartmentRequest) (response *ModifyDepartmentResponse, err error) {
	if request == nil {
		request = NewModifyDepartmentRequest()
	}
	response = NewModifyDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewSearchCommandSessionRequest() (request *SearchCommandSessionRequest) {
	request = &SearchCommandSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchCommandSession")
	return
}

func NewSearchCommandSessionResponse() (response *SearchCommandSessionResponse) {
	response = &SearchCommandSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检索命令和所属会话
func (c *Client) SearchCommandSession(request *SearchCommandSessionRequest) (response *SearchCommandSessionResponse, err error) {
	if request == nil {
		request = NewSearchCommandSessionRequest()
	}
	response = NewSearchCommandSessionResponse()
	err = c.Send(request, response)
	return
}

func NewImportDeviceAccountRequest() (request *ImportDeviceAccountRequest) {
	request = &ImportDeviceAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ImportDeviceAccount")
	return
}

func NewImportDeviceAccountResponse() (response *ImportDeviceAccountResponse) {
	response = &ImportDeviceAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入账号
func (c *Client) ImportDeviceAccount(request *ImportDeviceAccountRequest) (response *ImportDeviceAccountResponse, err error) {
	if request == nil {
		request = NewImportDeviceAccountRequest()
	}
	response = NewImportDeviceAccountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUsersDepartmentRequest() (request *ModifyUsersDepartmentRequest) {
	request = &ModifyUsersDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyUsersDepartment")
	return
}

func NewModifyUsersDepartmentResponse() (response *ModifyUsersDepartmentResponse) {
	response = &ModifyUsersDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改用户部门
func (c *Client) ModifyUsersDepartment(request *ModifyUsersDepartmentRequest) (response *ModifyUsersDepartmentResponse, err error) {
	if request == nil {
		request = NewModifyUsersDepartmentRequest()
	}
	response = NewModifyUsersDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewBindDeviceAccountPasswordRequest() (request *BindDeviceAccountPasswordRequest) {
	request = &BindDeviceAccountPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "BindDeviceAccountPassword")
	return
}

func NewBindDeviceAccountPasswordResponse() (response *BindDeviceAccountPasswordResponse) {
	response = &BindDeviceAccountPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定主机账号密码
func (c *Client) BindDeviceAccountPassword(request *BindDeviceAccountPasswordRequest) (response *BindDeviceAccountPasswordResponse, err error) {
	if request == nil {
		request = NewBindDeviceAccountPasswordRequest()
	}
	response = NewBindDeviceAccountPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessTimePolicyRequest() (request *ModifyAccessTimePolicyRequest) {
	request = &ModifyAccessTimePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessTimePolicy")
	return
}

func NewModifyAccessTimePolicyResponse() (response *ModifyAccessTimePolicyResponse) {
	response = &ModifyAccessTimePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改访问时间段限制
func (c *Client) ModifyAccessTimePolicy(request *ModifyAccessTimePolicyRequest) (response *ModifyAccessTimePolicyResponse, err error) {
	if request == nil {
		request = NewModifyAccessTimePolicyRequest()
	}
	response = NewModifyAccessTimePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationTaskStatisticsRequest() (request *DescribeOperationTaskStatisticsRequest) {
	request = &DescribeOperationTaskStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationTaskStatistics")
	return
}

func NewDescribeOperationTaskStatisticsResponse() (response *DescribeOperationTaskStatisticsResponse) {
	response = &DescribeOperationTaskStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维任务统计信息
func (c *Client) DescribeOperationTaskStatistics(request *DescribeOperationTaskStatisticsRequest) (response *DescribeOperationTaskStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeOperationTaskStatisticsRequest()
	}
	response = NewDescribeOperationTaskStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationTaskDetailRequest() (request *DescribeOperationTaskDetailRequest) {
	request = &DescribeOperationTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationTaskDetail")
	return
}

func NewDescribeOperationTaskDetailResponse() (response *DescribeOperationTaskDetailResponse) {
	response = &DescribeOperationTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维任务详情
func (c *Client) DescribeOperationTaskDetail(request *DescribeOperationTaskDetailRequest) (response *DescribeOperationTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeOperationTaskDetailRequest()
	}
	response = NewDescribeOperationTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPasswordSettingRequest() (request *ModifyPasswordSettingRequest) {
	request = &ModifyPasswordSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyPasswordSetting")
	return
}

func NewModifyPasswordSettingResponse() (response *ModifyPasswordSettingResponse) {
	response = &ModifyPasswordSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改密码配置信息
func (c *Client) ModifyPasswordSetting(request *ModifyPasswordSettingRequest) (response *ModifyPasswordSettingResponse, err error) {
	if request == nil {
		request = NewModifyPasswordSettingRequest()
	}
	response = NewModifyPasswordSettingResponse()
	err = c.Send(request, response)
	return
}

func NewSearchTaskResultRequest() (request *SearchTaskResultRequest) {
	request = &SearchTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchTaskResult")
	return
}

func NewSearchTaskResultResponse() (response *SearchTaskResultResponse) {
	response = &SearchTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索运维任务执行结果
func (c *Client) SearchTaskResult(request *SearchTaskResultRequest) (response *SearchTaskResultResponse, err error) {
	if request == nil {
		request = NewSearchTaskResultRequest()
	}
	response = NewSearchTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAclRequest() (request *ModifyAclRequest) {
	request = &ModifyAclRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAcl")
	return
}

func NewModifyAclResponse() (response *ModifyAclResponse) {
	response = &ModifyAclResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问权限
func (c *Client) ModifyAcl(request *ModifyAclRequest) (response *ModifyAclResponse, err error) {
	if request == nil {
		request = NewModifyAclRequest()
	}
	response = NewModifyAclResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLoginSettingRequest() (request *ModifyLoginSettingRequest) {
	request = &ModifyLoginSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyLoginSetting")
	return
}

func NewModifyLoginSettingResponse() (response *ModifyLoginSettingResponse) {
	response = &ModifyLoginSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改登录相关配置信息
func (c *Client) ModifyLoginSetting(request *ModifyLoginSettingRequest) (response *ModifyLoginSettingResponse, err error) {
	if request == nil {
		request = NewModifyLoginSettingRequest()
	}
	response = NewModifyLoginSettingResponse()
	err = c.Send(request, response)
	return
}

func NewSearchTaskResultDetailRequest() (request *SearchTaskResultDetailRequest) {
	request = &SearchTaskResultDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchTaskResultDetail")
	return
}

func NewSearchTaskResultDetailResponse() (response *SearchTaskResultDetailResponse) {
	response = &SearchTaskResultDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索运维任务执行结果详情
func (c *Client) SearchTaskResultDetail(request *SearchTaskResultDetailRequest) (response *SearchTaskResultDetailResponse, err error) {
	if request == nil {
		request = NewSearchTaskResultDetailRequest()
	}
	response = NewSearchTaskResultDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserGroupRequest() (request *CreateUserGroupRequest) {
	request = &CreateUserGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateUserGroup")
	return
}

func NewCreateUserGroupResponse() (response *CreateUserGroupResponse) {
	response = &CreateUserGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建用户组
func (c *Client) CreateUserGroup(request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
	if request == nil {
		request = NewCreateUserGroupRequest()
	}
	response = NewCreateUserGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDevicesRequest() (request *DeleteDevicesRequest) {
	request = &DeleteDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDevices")
	return
}

func NewDeleteDevicesResponse() (response *DeleteDevicesResponse) {
	response = &DeleteDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除主机
func (c *Client) DeleteDevices(request *DeleteDevicesRequest) (response *DeleteDevicesResponse, err error) {
	if request == nil {
		request = NewDeleteDevicesRequest()
	}
	response = NewDeleteDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewShowTopRequest() (request *ShowTopRequest) {
	request = &ShowTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ShowTop")
	return
}

func NewShowTopResponse() (response *ShowTopResponse) {
	response = &ShowTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示在线用户数、活跃用户、活跃主机、高危用户
func (c *Client) ShowTop(request *ShowTopRequest) (response *ShowTopResponse, err error) {
	if request == nil {
		request = NewShowTopRequest()
	}
	response = NewShowTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
	request = &DescribeDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDevices")
	return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
	response = &DescribeDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产列表
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
	if request == nil {
		request = NewDescribeDevicesRequest()
	}
	response = NewDescribeDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewShowGraphRequest() (request *ShowGraphRequest) {
	request = &ShowGraphRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ShowGraph")
	return
}

func NewShowGraphResponse() (response *ShowGraphResponse) {
	response = &ShowGraphResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示折线图
func (c *Client) ShowGraph(request *ShowGraphRequest) (response *ShowGraphResponse, err error) {
	if request == nil {
		request = NewShowGraphRequest()
	}
	response = NewShowGraphResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeviceAccountRequest() (request *CreateDeviceAccountRequest) {
	request = &CreateDeviceAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateDeviceAccount")
	return
}

func NewCreateDeviceAccountResponse() (response *CreateDeviceAccountResponse) {
	response = &CreateDeviceAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建主机账号
func (c *Client) CreateDeviceAccount(request *CreateDeviceAccountRequest) (response *CreateDeviceAccountResponse, err error) {
	if request == nil {
		request = NewCreateDeviceAccountRequest()
	}
	response = NewCreateDeviceAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDeviceGroupsRequest() (request *DeleteDeviceGroupsRequest) {
	request = &DeleteDeviceGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDeviceGroups")
	return
}

func NewDeleteDeviceGroupsResponse() (response *DeleteDeviceGroupsResponse) {
	response = &DeleteDeviceGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资产组
func (c *Client) DeleteDeviceGroups(request *DeleteDeviceGroupsRequest) (response *DeleteDeviceGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteDeviceGroupsRequest()
	}
	response = NewDeleteDeviceGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAuthModeSettingRequest() (request *ModifyAuthModeSettingRequest) {
	request = &ModifyAuthModeSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAuthModeSetting")
	return
}

func NewModifyAuthModeSettingResponse() (response *ModifyAuthModeSettingResponse) {
	response = &ModifyAuthModeSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改认证方式配置信息
func (c *Client) ModifyAuthModeSetting(request *ModifyAuthModeSettingRequest) (response *ModifyAuthModeSettingResponse, err error) {
	if request == nil {
		request = NewModifyAuthModeSettingRequest()
	}
	response = NewModifyAuthModeSettingResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
	request = &CreateUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateUser")
	return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
	response = &CreateUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建用户
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
	if request == nil {
		request = NewCreateUserRequest()
	}
	response = NewCreateUserResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessWhiteListAutoStatusRequest() (request *ModifyAccessWhiteListAutoStatusRequest) {
	request = &ModifyAccessWhiteListAutoStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessWhiteListAutoStatus")
	return
}

func NewModifyAccessWhiteListAutoStatusResponse() (response *ModifyAccessWhiteListAutoStatusResponse) {
	response = &ModifyAccessWhiteListAutoStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问白名单自动添加IP状态：开启或关闭自动添加IP
func (c *Client) ModifyAccessWhiteListAutoStatus(request *ModifyAccessWhiteListAutoStatusRequest) (response *ModifyAccessWhiteListAutoStatusResponse, err error) {
	if request == nil {
		request = NewModifyAccessWhiteListAutoStatusRequest()
	}
	response = NewModifyAccessWhiteListAutoStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecuritySettingRequest() (request *DescribeSecuritySettingRequest) {
	request = &DescribeSecuritySettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeSecuritySetting")
	return
}

func NewDescribeSecuritySettingResponse() (response *DescribeSecuritySettingResponse) {
	response = &DescribeSecuritySettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全配置信息
func (c *Client) DescribeSecuritySetting(request *DescribeSecuritySettingRequest) (response *DescribeSecuritySettingResponse, err error) {
	if request == nil {
		request = NewDescribeSecuritySettingRequest()
	}
	response = NewDescribeSecuritySettingResponse()
	err = c.Send(request, response)
	return
}

func NewResetUserRequest() (request *ResetUserRequest) {
	request = &ResetUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ResetUser")
	return
}

func NewResetUserResponse() (response *ResetUserResponse) {
	response = &ResetUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置用户密码
func (c *Client) ResetUser(request *ResetUserRequest) (response *ResetUserResponse, err error) {
	if request == nil {
		request = NewResetUserRequest()
	}
	response = NewResetUserResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeviceAccountBatchRequest() (request *CreateDeviceAccountBatchRequest) {
	request = &CreateDeviceAccountBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateDeviceAccountBatch")
	return
}

func NewCreateDeviceAccountBatchResponse() (response *CreateDeviceAccountBatchResponse) {
	response = &CreateDeviceAccountBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量添加主机账号
func (c *Client) CreateDeviceAccountBatch(request *CreateDeviceAccountBatchRequest) (response *CreateDeviceAccountBatchResponse, err error) {
	if request == nil {
		request = NewCreateDeviceAccountBatchRequest()
	}
	response = NewCreateDeviceAccountBatchResponse()
	err = c.Send(request, response)
	return
}

func NewSearchCommandBySidRequest() (request *SearchCommandBySidRequest) {
	request = &SearchCommandBySidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchCommandBySid")
	return
}

func NewSearchCommandBySidResponse() (response *SearchCommandBySidResponse) {
	response = &SearchCommandBySidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据会话Id搜索Command
func (c *Client) SearchCommandBySid(request *SearchCommandBySidRequest) (response *SearchCommandBySidResponse, err error) {
	if request == nil {
		request = NewSearchCommandBySidRequest()
	}
	response = NewSearchCommandBySidResponse()
	err = c.Send(request, response)
	return
}

func NewDeployResourceRequest() (request *DeployResourceRequest) {
	request = &DeployResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeployResource")
	return
}

func NewDeployResourceResponse() (response *DeployResourceResponse) {
	response = &DeployResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开通服务，初始化资源，只针对新购资源
func (c *Client) DeployResource(request *DeployResourceRequest) (response *DeployResourceResponse, err error) {
	if request == nil {
		request = NewDeployResourceRequest()
	}
	response = NewDeployResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseUsageRequest() (request *DescribeLicenseUsageRequest) {
	request = &DescribeLicenseUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeLicenseUsage")
	return
}

func NewDescribeLicenseUsageResponse() (response *DescribeLicenseUsageResponse) {
	response = &DescribeLicenseUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供license平台用于查询license用量
func (c *Client) DescribeLicenseUsage(request *DescribeLicenseUsageRequest) (response *DescribeLicenseUsageResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseUsageRequest()
	}
	response = NewDescribeLicenseUsageResponse()
	err = c.Send(request, response)
	return
}

func NewImportDevicesRequest() (request *ImportDevicesRequest) {
	request = &ImportDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ImportDevices")
	return
}

func NewImportDevicesResponse() (response *ImportDevicesResponse) {
	response = &ImportDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入主机信息
func (c *Client) ImportDevices(request *ImportDevicesRequest) (response *ImportDevicesResponse, err error) {
	if request == nil {
		request = NewImportDevicesRequest()
	}
	response = NewImportDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserCountRequest() (request *DescribeUserCountRequest) {
	request = &DescribeUserCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeUserCount")
	return
}

func NewDescribeUserCountResponse() (response *DescribeUserCountResponse) {
	response = &DescribeUserCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户总数
func (c *Client) DescribeUserCount(request *DescribeUserCountRequest) (response *DescribeUserCountResponse, err error) {
	if request == nil {
		request = NewDescribeUserCountRequest()
	}
	response = NewDescribeUserCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDevicesPortRequest() (request *ModifyDevicesPortRequest) {
	request = &ModifyDevicesPortRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDevicesPort")
	return
}

func NewModifyDevicesPortResponse() (response *ModifyDevicesPortResponse) {
	response = &ModifyDevicesPortResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改主机端口
func (c *Client) ModifyDevicesPort(request *ModifyDevicesPortRequest) (response *ModifyDevicesPortResponse, err error) {
	if request == nil {
		request = NewModifyDevicesPortRequest()
	}
	response = NewModifyDevicesPortResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDasbResourceRequest() (request *ModifyDasbResourceRequest) {
	request = &ModifyDasbResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDasbResource")
	return
}

func NewModifyDasbResourceResponse() (response *ModifyDasbResourceResponse) {
	response = &ModifyDasbResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 堡垒机变配接口
func (c *Client) ModifyDasbResource(request *ModifyDasbResourceRequest) (response *ModifyDasbResourceResponse, err error) {
	if request == nil {
		request = NewModifyDasbResourceRequest()
	}
	response = NewModifyDasbResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationTasksRequest() (request *DescribeOperationTasksRequest) {
	request = &DescribeOperationTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationTasks")
	return
}

func NewDescribeOperationTasksResponse() (response *DescribeOperationTasksResponse) {
	response = &DescribeOperationTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维任务列表
func (c *Client) DescribeOperationTasks(request *DescribeOperationTasksRequest) (response *DescribeOperationTasksResponse, err error) {
	if request == nil {
		request = NewDescribeOperationTasksRequest()
	}
	response = NewDescribeOperationTasksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyExternalDeviceRequest() (request *ModifyExternalDeviceRequest) {
	request = &ModifyExternalDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyExternalDevice")
	return
}

func NewModifyExternalDeviceResponse() (response *ModifyExternalDeviceResponse) {
	response = &ModifyExternalDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改外部资产信息
func (c *Client) ModifyExternalDevice(request *ModifyExternalDeviceRequest) (response *ModifyExternalDeviceResponse, err error) {
	if request == nil {
		request = NewModifyExternalDeviceRequest()
	}
	response = NewModifyExternalDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResourceRequest() (request *ModifyResourceRequest) {
	request = &ModifyResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyResource")
	return
}

func NewModifyResourceResponse() (response *ModifyResourceResponse) {
	response = &ModifyResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 针对资源开通或关闭高级特性，如数据库管理等
func (c *Client) ModifyResource(request *ModifyResourceRequest) (response *ModifyResourceResponse, err error) {
	if request == nil {
		request = NewModifyResourceRequest()
	}
	response = NewModifyResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLDAPUnitSetRequest() (request *DescribeLDAPUnitSetRequest) {
	request = &DescribeLDAPUnitSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeLDAPUnitSet")
	return
}

func NewDescribeLDAPUnitSetResponse() (response *DescribeLDAPUnitSetResponse) {
	response = &DescribeLDAPUnitSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取LDAP ou 列表
func (c *Client) DescribeLDAPUnitSet(request *DescribeLDAPUnitSetRequest) (response *DescribeLDAPUnitSetResponse, err error) {
	if request == nil {
		request = NewDescribeLDAPUnitSetRequest()
	}
	response = NewDescribeLDAPUnitSetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsersRequest() (request *DescribeUsersRequest) {
	request = &DescribeUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeUsers")
	return
}

func NewDescribeUsersResponse() (response *DescribeUsersResponse) {
	response = &DescribeUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户列表
func (c *Client) DescribeUsers(request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
	if request == nil {
		request = NewDescribeUsersRequest()
	}
	response = NewDescribeUsersResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSubtaskResultByIdRequest() (request *SearchSubtaskResultByIdRequest) {
	request = &SearchSubtaskResultByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchSubtaskResultById")
	return
}

func NewSearchSubtaskResultByIdResponse() (response *SearchSubtaskResultByIdResponse) {
	response = &SearchSubtaskResultByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维子任务执行结果
func (c *Client) SearchSubtaskResultById(request *SearchSubtaskResultByIdRequest) (response *SearchSubtaskResultByIdResponse, err error) {
	if request == nil {
		request = NewSearchSubtaskResultByIdRequest()
	}
	response = NewSearchSubtaskResultByIdResponse()
	err = c.Send(request, response)
	return
}

func NewSearchFileSessionRequest() (request *SearchFileSessionRequest) {
	request = &SearchFileSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchFileSession")
	return
}

func NewSearchFileSessionResponse() (response *SearchFileSessionResponse) {
	response = &SearchFileSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索传输的文件情况和所属的会话
func (c *Client) SearchFileSession(request *SearchFileSessionRequest) (response *SearchFileSessionResponse, err error) {
	if request == nil {
		request = NewSearchFileSessionRequest()
	}
	response = NewSearchFileSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDepartmentRequest() (request *DeleteDepartmentRequest) {
	request = &DeleteDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDepartment")
	return
}

func NewDeleteDepartmentResponse() (response *DeleteDepartmentResponse) {
	response = &DeleteDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除部门
func (c *Client) DeleteDepartment(request *DeleteDepartmentRequest) (response *DeleteDepartmentResponse, err error) {
	if request == nil {
		request = NewDeleteDepartmentRequest()
	}
	response = NewDeleteDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceAccountsRequest() (request *DescribeDeviceAccountsRequest) {
	request = &DescribeDeviceAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDeviceAccounts")
	return
}

func NewDescribeDeviceAccountsResponse() (response *DescribeDeviceAccountsResponse) {
	response = &DescribeDeviceAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机账号列表
func (c *Client) DescribeDeviceAccounts(request *DescribeDeviceAccountsRequest) (response *DescribeDeviceAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceAccountsRequest()
	}
	response = NewDescribeDeviceAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSyncFlagRequest() (request *DescribeAssetSyncFlagRequest) {
	request = &DescribeAssetSyncFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAssetSyncFlag")
	return
}

func NewDescribeAssetSyncFlagResponse() (response *DescribeAssetSyncFlagResponse) {
	response = &DescribeAssetSyncFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产自动同步开关
func (c *Client) DescribeAssetSyncFlag(request *DescribeAssetSyncFlagRequest) (response *DescribeAssetSyncFlagResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSyncFlagRequest()
	}
	response = NewDescribeAssetSyncFlagResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCmdTemplateRequest() (request *ModifyCmdTemplateRequest) {
	request = &ModifyCmdTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyCmdTemplate")
	return
}

func NewModifyCmdTemplateResponse() (response *ModifyCmdTemplateResponse) {
	response = &ModifyCmdTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改高危命令模板
func (c *Client) ModifyCmdTemplate(request *ModifyCmdTemplateRequest) (response *ModifyCmdTemplateResponse, err error) {
	if request == nil {
		request = NewModifyCmdTemplateRequest()
	}
	response = NewModifyCmdTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceGroupMembersRequest() (request *DescribeDeviceGroupMembersRequest) {
	request = &DescribeDeviceGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDeviceGroupMembers")
	return
}

func NewDescribeDeviceGroupMembersResponse() (response *DescribeDeviceGroupMembersResponse) {
	response = &DescribeDeviceGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产组成员列表
func (c *Client) DescribeDeviceGroupMembers(request *DescribeDeviceGroupMembersRequest) (response *DescribeDeviceGroupMembersResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceGroupMembersRequest()
	}
	response = NewDescribeDeviceGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogOutputSettingsRequest() (request *ModifyLogOutputSettingsRequest) {
	request = &ModifyLogOutputSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyLogOutputSettings")
	return
}

func NewModifyLogOutputSettingsResponse() (response *ModifyLogOutputSettingsResponse) {
	response = &ModifyLogOutputSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志外发配置
func (c *Client) ModifyLogOutputSettings(request *ModifyLogOutputSettingsRequest) (response *ModifyLogOutputSettingsResponse, err error) {
	if request == nil {
		request = NewModifyLogOutputSettingsRequest()
	}
	response = NewModifyLogOutputSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCmdTemplateRequest() (request *CreateCmdTemplateRequest) {
	request = &CreateCmdTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateCmdTemplate")
	return
}

func NewCreateCmdTemplateResponse() (response *CreateCmdTemplateResponse) {
	response = &CreateCmdTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建高危命令模板
func (c *Client) CreateCmdTemplate(request *CreateCmdTemplateRequest) (response *CreateCmdTemplateResponse, err error) {
	if request == nil {
		request = NewCreateCmdTemplateRequest()
	}
	response = NewCreateCmdTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeviceGroupRequest() (request *CreateDeviceGroupRequest) {
	request = &CreateDeviceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateDeviceGroup")
	return
}

func NewCreateDeviceGroupResponse() (response *CreateDeviceGroupResponse) {
	response = &CreateDeviceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建资产组
func (c *Client) CreateDeviceGroup(request *CreateDeviceGroupRequest) (response *CreateDeviceGroupResponse, err error) {
	if request == nil {
		request = NewCreateDeviceGroupRequest()
	}
	response = NewCreateDeviceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserGroupsRequest() (request *DeleteUserGroupsRequest) {
	request = &DeleteUserGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteUserGroups")
	return
}

func NewDeleteUserGroupsResponse() (response *DeleteUserGroupsResponse) {
	response = &DeleteUserGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户组
func (c *Client) DeleteUserGroups(request *DeleteUserGroupsRequest) (response *DeleteUserGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteUserGroupsRequest()
	}
	response = NewDeleteUserGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceGroupsRequest() (request *DescribeDeviceGroupsRequest) {
	request = &DescribeDeviceGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDeviceGroups")
	return
}

func NewDescribeDeviceGroupsResponse() (response *DescribeDeviceGroupsResponse) {
	response = &DescribeDeviceGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产组列表
func (c *Client) DescribeDeviceGroups(request *DescribeDeviceGroupsRequest) (response *DescribeDeviceGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceGroupsRequest()
	}
	response = NewDescribeDeviceGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDeviceAccountsRequest() (request *DeleteDeviceAccountsRequest) {
	request = &DeleteDeviceAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDeviceAccounts")
	return
}

func NewDeleteDeviceAccountsResponse() (response *DeleteDeviceAccountsResponse) {
	response = &DeleteDeviceAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除主机账号
func (c *Client) DeleteDeviceAccounts(request *DeleteDeviceAccountsRequest) (response *DeleteDeviceAccountsResponse, err error) {
	if request == nil {
		request = NewDeleteDeviceAccountsRequest()
	}
	response = NewDeleteDeviceAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserGroupMembersRequest() (request *DeleteUserGroupMembersRequest) {
	request = &DeleteUserGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteUserGroupMembers")
	return
}

func NewDeleteUserGroupMembersResponse() (response *DeleteUserGroupMembersResponse) {
	response = &DeleteUserGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户组成员
func (c *Client) DeleteUserGroupMembers(request *DeleteUserGroupMembersRequest) (response *DeleteUserGroupMembersResponse, err error) {
	if request == nil {
		request = NewDeleteUserGroupMembersRequest()
	}
	response = NewDeleteUserGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessEntryRequest() (request *DescribeAccessEntryRequest) {
	request = &DescribeAccessEntryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccessEntry")
	return
}

func NewDescribeAccessEntryResponse() (response *DescribeAccessEntryResponse) {
	response = &DescribeAccessEntryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维人员WEB访问入口
func (c *Client) DescribeAccessEntry(request *DescribeAccessEntryRequest) (response *DescribeAccessEntryResponse, err error) {
	if request == nil {
		request = NewDescribeAccessEntryRequest()
	}
	response = NewDescribeAccessEntryResponse()
	err = c.Send(request, response)
	return
}

func NewSearchCommandRequest() (request *SearchCommandRequest) {
	request = &SearchCommandRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchCommand")
	return
}

func NewSearchCommandResponse() (response *SearchCommandResponse) {
	response = &SearchCommandResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 命令执行检索
func (c *Client) SearchCommand(request *SearchCommandRequest) (response *SearchCommandResponse, err error) {
	if request == nil {
		request = NewSearchCommandRequest()
	}
	response = NewSearchCommandResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessWhiteListRulesRequest() (request *DescribeAccessWhiteListRulesRequest) {
	request = &DescribeAccessWhiteListRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccessWhiteListRules")
	return
}

func NewDescribeAccessWhiteListRulesResponse() (response *DescribeAccessWhiteListRulesResponse) {
	response = &DescribeAccessWhiteListRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询访问白名单规则列表
func (c *Client) DescribeAccessWhiteListRules(request *DescribeAccessWhiteListRulesRequest) (response *DescribeAccessWhiteListRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAccessWhiteListRulesRequest()
	}
	response = NewDescribeAccessWhiteListRulesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchStatementBySidRequest() (request *SearchStatementBySidRequest) {
	request = &SearchStatementBySidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchStatementBySid")
	return
}

func NewSearchStatementBySidResponse() (response *SearchStatementBySidResponse) {
	response = &SearchStatementBySidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据会话Id搜索数据库操作
func (c *Client) SearchStatementBySid(request *SearchStatementBySidRequest) (response *SearchStatementBySidResponse, err error) {
	if request == nil {
		request = NewSearchStatementBySidRequest()
	}
	response = NewSearchStatementBySidResponse()
	err = c.Send(request, response)
	return
}

func NewAddDeviceGroupMembersRequest() (request *AddDeviceGroupMembersRequest) {
	request = &AddDeviceGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "AddDeviceGroupMembers")
	return
}

func NewAddDeviceGroupMembersResponse() (response *AddDeviceGroupMembersResponse) {
	response = &AddDeviceGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加资产组成员
func (c *Client) AddDeviceGroupMembers(request *AddDeviceGroupMembersRequest) (response *AddDeviceGroupMembersResponse, err error) {
	if request == nil {
		request = NewAddDeviceGroupMembersRequest()
	}
	response = NewAddDeviceGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewKillSessionRequest() (request *KillSessionRequest) {
	request = &KillSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "KillSession")
	return
}

func NewKillSessionResponse() (response *KillSessionResponse) {
	response = &KillSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 会话切断
func (c *Client) KillSession(request *KillSessionRequest) (response *KillSessionResponse, err error) {
	if request == nil {
		request = NewKillSessionRequest()
	}
	response = NewKillSessionResponse()
	err = c.Send(request, response)
	return
}
