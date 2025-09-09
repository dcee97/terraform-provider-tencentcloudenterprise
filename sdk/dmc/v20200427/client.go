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

package v20200427

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-04-27"

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

func NewDeleteUserSQLRequest() (request *DeleteUserSQLRequest) {
	request = &DeleteUserSQLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DeleteUserSQL")
	return
}

func NewDeleteUserSQLResponse() (response *DeleteUserSQLResponse) {
	response = &DeleteUserSQLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DeleteUserSQL)用于删除用户SQL记录。
func (c *Client) DeleteUserSQL(request *DeleteUserSQLRequest) (response *DeleteUserSQLResponse, err error) {
	if request == nil {
		request = NewDeleteUserSQLRequest()
	}
	response = NewDeleteUserSQLResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyDmcResourceUserRequest() (request *DestroyDmcResourceUserRequest) {
	request = &DestroyDmcResourceUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DestroyDmcResourceUser")
	return
}

func NewDestroyDmcResourceUserResponse() (response *DestroyDmcResourceUserResponse) {
	response = &DestroyDmcResourceUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DestroyDmcResourceUser)用于用户销毁dmc资源
func (c *Client) DestroyDmcResourceUser(request *DestroyDmcResourceUserRequest) (response *DestroyDmcResourceUserResponse, err error) {
	if request == nil {
		request = NewDestroyDmcResourceUserRequest()
	}
	response = NewDestroyDmcResourceUserResponse()
	err = c.Send(request, response)
	return
}

func NewEnableAuthUserRequest() (request *EnableAuthUserRequest) {
	request = &EnableAuthUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "EnableAuthUser")
	return
}

func NewEnableAuthUserResponse() (response *EnableAuthUserResponse) {
	response = &EnableAuthUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用用户
func (c *Client) EnableAuthUser(request *EnableAuthUserRequest) (response *EnableAuthUserResponse, err error) {
	if request == nil {
		request = NewEnableAuthUserRequest()
	}
	response = NewEnableAuthUserResponse()
	err = c.Send(request, response)
	return
}

func NewLeaseDmcResourceRouteRequest() (request *LeaseDmcResourceRouteRequest) {
	request = &LeaseDmcResourceRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "LeaseDmcResourceRoute")
	return
}

func NewLeaseDmcResourceRouteResponse() (response *LeaseDmcResourceRouteResponse) {
	response = &LeaseDmcResourceRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(LeaseDmcResourceRoute)用于资源路由续期
func (c *Client) LeaseDmcResourceRoute(request *LeaseDmcResourceRouteRequest) (response *LeaseDmcResourceRouteResponse, err error) {
	if request == nil {
		request = NewLeaseDmcResourceRouteRequest()
	}
	response = NewLeaseDmcResourceRouteResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIpWhitelistRequest() (request *CreateIpWhitelistRequest) {
	request = &CreateIpWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "CreateIpWhitelist")
	return
}

func NewCreateIpWhitelistResponse() (response *CreateIpWhitelistResponse) {
	response = &CreateIpWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加一个白名单IP
func (c *Client) CreateIpWhitelist(request *CreateIpWhitelistRequest) (response *CreateIpWhitelistResponse, err error) {
	if request == nil {
		request = NewCreateIpWhitelistRequest()
	}
	response = NewCreateIpWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAuthUserRequest() (request *DeleteAuthUserRequest) {
	request = &DeleteAuthUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DeleteAuthUser")
	return
}

func NewDeleteAuthUserResponse() (response *DeleteAuthUserResponse) {
	response = &DeleteAuthUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户
func (c *Client) DeleteAuthUser(request *DeleteAuthUserRequest) (response *DeleteAuthUserResponse, err error) {
	if request == nil {
		request = NewDeleteAuthUserRequest()
	}
	response = NewDeleteAuthUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpWhitelistsRequest() (request *DescribeIpWhitelistsRequest) {
	request = &DescribeIpWhitelistsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeIpWhitelists")
	return
}

func NewDescribeIpWhitelistsResponse() (response *DescribeIpWhitelistsResponse) {
	response = &DescribeIpWhitelistsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询符合条件的白名单IP
func (c *Client) DescribeIpWhitelists(request *DescribeIpWhitelistsRequest) (response *DescribeIpWhitelistsResponse, err error) {
	if request == nil {
		request = NewDescribeIpWhitelistsRequest()
	}
	response = NewDescribeIpWhitelistsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllPermissionsOfInstanceRequest() (request *DescribeAllPermissionsOfInstanceRequest) {
	request = &DescribeAllPermissionsOfInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeAllPermissionsOfInstance")
	return
}

func NewDescribeAllPermissionsOfInstanceResponse() (response *DescribeAllPermissionsOfInstanceResponse) {
	response = &DescribeAllPermissionsOfInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个实例关联的全部权限
func (c *Client) DescribeAllPermissionsOfInstance(request *DescribeAllPermissionsOfInstanceRequest) (response *DescribeAllPermissionsOfInstanceResponse, err error) {
	if request == nil {
		request = NewDescribeAllPermissionsOfInstanceRequest()
	}
	response = NewDescribeAllPermissionsOfInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewRevokePermissionFromUserRequest() (request *RevokePermissionFromUserRequest) {
	request = &RevokePermissionFromUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "RevokePermissionFromUser")
	return
}

func NewRevokePermissionFromUserResponse() (response *RevokePermissionFromUserResponse) {
	response = &RevokePermissionFromUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 收回用户权限
func (c *Client) RevokePermissionFromUser(request *RevokePermissionFromUserRequest) (response *RevokePermissionFromUserResponse, err error) {
	if request == nil {
		request = NewRevokePermissionFromUserRequest()
	}
	response = NewRevokePermissionFromUserResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyDmcResourceRequest() (request *DestroyDmcResourceRequest) {
	request = &DestroyDmcResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DestroyDmcResource")
	return
}

func NewDestroyDmcResourceResponse() (response *DestroyDmcResourceResponse) {
	response = &DestroyDmcResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DestroyDmcResource)用于dmc资源销毁。
func (c *Client) DestroyDmcResource(request *DestroyDmcResourceRequest) (response *DestroyDmcResourceResponse, err error) {
	if request == nil {
		request = NewDestroyDmcResourceRequest()
	}
	response = NewDestroyDmcResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDmcResourcesRequest() (request *DescribeDmcResourcesRequest) {
	request = &DescribeDmcResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeDmcResources")
	return
}

func NewDescribeDmcResourcesResponse() (response *DescribeDmcResourcesResponse) {
	response = &DescribeDmcResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDmcResources)用于拉取dmc资源列表。
func (c *Client) DescribeDmcResources(request *DescribeDmcResourcesRequest) (response *DescribeDmcResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeDmcResourcesRequest()
	}
	response = NewDescribeDmcResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDmcResourceUsersRequest() (request *DescribeDmcResourceUsersRequest) {
	request = &DescribeDmcResourceUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeDmcResourceUsers")
	return
}

func NewDescribeDmcResourceUsersResponse() (response *DescribeDmcResourceUsersResponse) {
	response = &DescribeDmcResourceUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDmcResourceUsers)用于拉取dmc资源用户列表
func (c *Client) DescribeDmcResourceUsers(request *DescribeDmcResourceUsersRequest) (response *DescribeDmcResourceUsersResponse, err error) {
	if request == nil {
		request = NewDescribeDmcResourceUsersRequest()
	}
	response = NewDescribeDmcResourceUsersResponse()
	err = c.Send(request, response)
	return
}

func NewAddUserSQLRequest() (request *AddUserSQLRequest) {
	request = &AddUserSQLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "AddUserSQL")
	return
}

func NewAddUserSQLResponse() (response *AddUserSQLResponse) {
	response = &AddUserSQLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(AddUserSQL)用于添加用户SQL记录。
func (c *Client) AddUserSQL(request *AddUserSQLRequest) (response *AddUserSQLResponse, err error) {
	if request == nil {
		request = NewAddUserSQLRequest()
	}
	response = NewAddUserSQLResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthRolesRequest() (request *DescribeAuthRolesRequest) {
	request = &DescribeAuthRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeAuthRoles")
	return
}

func NewDescribeAuthRolesResponse() (response *DescribeAuthRolesResponse) {
	response = &DescribeAuthRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询角色类型
func (c *Client) DescribeAuthRoles(request *DescribeAuthRolesRequest) (response *DescribeAuthRolesResponse, err error) {
	if request == nil {
		request = NewDescribeAuthRolesRequest()
	}
	response = NewDescribeAuthRolesResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchIpWhitelistRequest() (request *SwitchIpWhitelistRequest) {
	request = &SwitchIpWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "SwitchIpWhitelist")
	return
}

func NewSwitchIpWhitelistResponse() (response *SwitchIpWhitelistResponse) {
	response = &SwitchIpWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启或者关闭IP白名单
func (c *Client) SwitchIpWhitelist(request *SwitchIpWhitelistRequest) (response *SwitchIpWhitelistResponse, err error) {
	if request == nil {
		request = NewSwitchIpWhitelistRequest()
	}
	response = NewSwitchIpWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDmcResourceVipVportRequest() (request *ModifyDmcResourceVipVportRequest) {
	request = &ModifyDmcResourceVipVportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "ModifyDmcResourceVipVport")
	return
}

func NewModifyDmcResourceVipVportResponse() (response *ModifyDmcResourceVipVportResponse) {
	response = &ModifyDmcResourceVipVportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyDmcResourceVipVport)用于修改dmc资源vip/vport 。
func (c *Client) ModifyDmcResourceVipVport(request *ModifyDmcResourceVipVportRequest) (response *ModifyDmcResourceVipVportResponse, err error) {
	if request == nil {
		request = NewModifyDmcResourceVipVportRequest()
	}
	response = NewModifyDmcResourceVipVportResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePermissionApplyRequest() (request *CreatePermissionApplyRequest) {
	request = &CreatePermissionApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "CreatePermissionApply")
	return
}

func NewCreatePermissionApplyResponse() (response *CreatePermissionApplyResponse) {
	response = &CreatePermissionApplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建权限申请单
func (c *Client) CreatePermissionApply(request *CreatePermissionApplyRequest) (response *CreatePermissionApplyResponse, err error) {
	if request == nil {
		request = NewCreatePermissionApplyRequest()
	}
	response = NewCreatePermissionApplyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDmcResourceUserRequest() (request *CreateDmcResourceUserRequest) {
	request = &CreateDmcResourceUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "CreateDmcResourceUser")
	return
}

func NewCreateDmcResourceUserResponse() (response *CreateDmcResourceUserResponse) {
	response = &CreateDmcResourceUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateDmcResourceUser)用于创建dmc资源用户信息
func (c *Client) CreateDmcResourceUser(request *CreateDmcResourceUserRequest) (response *CreateDmcResourceUserResponse, err error) {
	if request == nil {
		request = NewCreateDmcResourceUserRequest()
	}
	response = NewCreateDmcResourceUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePermissionsOfInstanceRequest() (request *DescribePermissionsOfInstanceRequest) {
	request = &DescribePermissionsOfInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribePermissionsOfInstance")
	return
}

func NewDescribePermissionsOfInstanceResponse() (response *DescribePermissionsOfInstanceResponse) {
	response = &DescribePermissionsOfInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户对某个实例的权限
func (c *Client) DescribePermissionsOfInstance(request *DescribePermissionsOfInstanceRequest) (response *DescribePermissionsOfInstanceResponse, err error) {
	if request == nil {
		request = NewDescribePermissionsOfInstanceRequest()
	}
	response = NewDescribePermissionsOfInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewGrantPermissionToUserRequest() (request *GrantPermissionToUserRequest) {
	request = &GrantPermissionToUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "GrantPermissionToUser")
	return
}

func NewGrantPermissionToUserResponse() (response *GrantPermissionToUserResponse) {
	response = &GrantPermissionToUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给用户授权
func (c *Client) GrantPermissionToUser(request *GrantPermissionToUserRequest) (response *GrantPermissionToUserResponse, err error) {
	if request == nil {
		request = NewGrantPermissionToUserRequest()
	}
	response = NewGrantPermissionToUserResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAuthUserRequest() (request *CreateAuthUserRequest) {
	request = &CreateAuthUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "CreateAuthUser")
	return
}

func NewCreateAuthUserResponse() (response *CreateAuthUserResponse) {
	response = &CreateAuthUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建用户
func (c *Client) CreateAuthUser(request *CreateAuthUserRequest) (response *CreateAuthUserResponse, err error) {
	if request == nil {
		request = NewCreateAuthUserRequest()
	}
	response = NewCreateAuthUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCurrentUserRequest() (request *DescribeCurrentUserRequest) {
	request = &DescribeCurrentUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeCurrentUser")
	return
}

func NewDescribeCurrentUserResponse() (response *DescribeCurrentUserResponse) {
	response = &DescribeCurrentUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前用户信息
func (c *Client) DescribeCurrentUser(request *DescribeCurrentUserRequest) (response *DescribeCurrentUserResponse, err error) {
	if request == nil {
		request = NewDescribeCurrentUserRequest()
	}
	response = NewDescribeCurrentUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthUsersRequest() (request *DescribeAuthUsersRequest) {
	request = &DescribeAuthUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeAuthUsers")
	return
}

func NewDescribeAuthUsersResponse() (response *DescribeAuthUsersResponse) {
	response = &DescribeAuthUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询访问控制服务的所有用户
func (c *Client) DescribeAuthUsers(request *DescribeAuthUsersRequest) (response *DescribeAuthUsersResponse, err error) {
	if request == nil {
		request = NewDescribeAuthUsersRequest()
	}
	response = NewDescribeAuthUsersResponse()
	err = c.Send(request, response)
	return
}

func NewCheckDmcResourceRouteRequest() (request *CheckDmcResourceRouteRequest) {
	request = &CheckDmcResourceRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "CheckDmcResourceRoute")
	return
}

func NewCheckDmcResourceRouteResponse() (response *CheckDmcResourceRouteResponse) {
	response = &CheckDmcResourceRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CheckDmcResourceRoute)用于检查资源连通性
func (c *Client) CheckDmcResourceRoute(request *CheckDmcResourceRouteRequest) (response *CheckDmcResourceRouteResponse, err error) {
	if request == nil {
		request = NewCheckDmcResourceRouteRequest()
	}
	response = NewCheckDmcResourceRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserSQLMateInfosRequest() (request *DescribeUserSQLMateInfosRequest) {
	request = &DescribeUserSQLMateInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeUserSQLMateInfos")
	return
}

func NewDescribeUserSQLMateInfosResponse() (response *DescribeUserSQLMateInfosResponse) {
	response = &DescribeUserSQLMateInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeUserSQLMateInfos)用于dmc拉取某条用户SQL元数据信息。
func (c *Client) DescribeUserSQLMateInfos(request *DescribeUserSQLMateInfosRequest) (response *DescribeUserSQLMateInfosResponse, err error) {
	if request == nil {
		request = NewDescribeUserSQLMateInfosRequest()
	}
	response = NewDescribeUserSQLMateInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserSQLsRequest() (request *DescribeUserSQLsRequest) {
	request = &DescribeUserSQLsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeUserSQLs")
	return
}

func NewDescribeUserSQLsResponse() (response *DescribeUserSQLsResponse) {
	response = &DescribeUserSQLsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeUserSQLs)用于dmc拉取某条用户SQL数据列表。
func (c *Client) DescribeUserSQLs(request *DescribeUserSQLsRequest) (response *DescribeUserSQLsResponse, err error) {
	if request == nil {
		request = NewDescribeUserSQLsRequest()
	}
	response = NewDescribeUserSQLsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDmcResourceUserRequest() (request *ModifyDmcResourceUserRequest) {
	request = &ModifyDmcResourceUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "ModifyDmcResourceUser")
	return
}

func NewModifyDmcResourceUserResponse() (response *ModifyDmcResourceUserResponse) {
	response = &ModifyDmcResourceUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyDmcResourceUser)用于dmc资源用户信息修改
func (c *Client) ModifyDmcResourceUser(request *ModifyDmcResourceUserRequest) (response *ModifyDmcResourceUserResponse, err error) {
	if request == nil {
		request = NewModifyDmcResourceUserRequest()
	}
	response = NewModifyDmcResourceUserResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIpWhitelistRequest() (request *ModifyIpWhitelistRequest) {
	request = &ModifyIpWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "ModifyIpWhitelist")
	return
}

func NewModifyIpWhitelistResponse() (response *ModifyIpWhitelistResponse) {
	response = &ModifyIpWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改IP白名单
func (c *Client) ModifyIpWhitelist(request *ModifyIpWhitelistRequest) (response *ModifyIpWhitelistResponse, err error) {
	if request == nil {
		request = NewModifyIpWhitelistRequest()
	}
	response = NewModifyIpWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserSQLRequest() (request *ModifyUserSQLRequest) {
	request = &ModifyUserSQLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "ModifyUserSQL")
	return
}

func NewModifyUserSQLResponse() (response *ModifyUserSQLResponse) {
	response = &ModifyUserSQLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyUserSQL)用于dmc修改用户SQL信息。
func (c *Client) ModifyUserSQL(request *ModifyUserSQLRequest) (response *ModifyUserSQLResponse, err error) {
	if request == nil {
		request = NewModifyUserSQLRequest()
	}
	response = NewModifyUserSQLResponse()
	err = c.Send(request, response)
	return
}

func NewCheckPermissionApplyRequest() (request *CheckPermissionApplyRequest) {
	request = &CheckPermissionApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "CheckPermissionApply")
	return
}

func NewCheckPermissionApplyResponse() (response *CheckPermissionApplyResponse) {
	response = &CheckPermissionApplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 处理权限申请单，通过或拒绝
func (c *Client) CheckPermissionApply(request *CheckPermissionApplyRequest) (response *CheckPermissionApplyResponse, err error) {
	if request == nil {
		request = NewCheckPermissionApplyRequest()
	}
	response = NewCheckPermissionApplyResponse()
	err = c.Send(request, response)
	return
}

func NewDisableAuthUserRequest() (request *DisableAuthUserRequest) {
	request = &DisableAuthUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DisableAuthUser")
	return
}

func NewDisableAuthUserResponse() (response *DisableAuthUserResponse) {
	response = &DisableAuthUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用用户
func (c *Client) DisableAuthUser(request *DisableAuthUserRequest) (response *DisableAuthUserResponse, err error) {
	if request == nil {
		request = NewDisableAuthUserRequest()
	}
	response = NewDisableAuthUserResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteSqlRequest() (request *ExecuteSqlRequest) {
	request = &ExecuteSqlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "ExecuteSql")
	return
}

func NewExecuteSqlResponse() (response *ExecuteSqlResponse) {
	response = &ExecuteSqlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ExecuteSql)用于执行用户SQL。
func (c *Client) ExecuteSql(request *ExecuteSqlRequest) (response *ExecuteSqlResponse, err error) {
	if request == nil {
		request = NewExecuteSqlRequest()
	}
	response = NewExecuteSqlResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDmcResourceNameRequest() (request *ModifyDmcResourceNameRequest) {
	request = &ModifyDmcResourceNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "ModifyDmcResourceName")
	return
}

func NewModifyDmcResourceNameResponse() (response *ModifyDmcResourceNameResponse) {
	response = &ModifyDmcResourceNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyDmcResourceName)用于修改dmc资源名称。
func (c *Client) ModifyDmcResourceName(request *ModifyDmcResourceNameRequest) (response *ModifyDmcResourceNameResponse, err error) {
	if request == nil {
		request = NewModifyDmcResourceNameRequest()
	}
	response = NewModifyDmcResourceNameResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIpWhitelistRequest() (request *DeleteIpWhitelistRequest) {
	request = &DeleteIpWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DeleteIpWhitelist")
	return
}

func NewDeleteIpWhitelistResponse() (response *DeleteIpWhitelistResponse) {
	response = &DeleteIpWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从IP白名单中删除一个IP
func (c *Client) DeleteIpWhitelist(request *DeleteIpWhitelistRequest) (response *DeleteIpWhitelistResponse, err error) {
	if request == nil {
		request = NewDeleteIpWhitelistRequest()
	}
	response = NewDeleteIpWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePermissionsOfUserRequest() (request *DescribePermissionsOfUserRequest) {
	request = &DescribePermissionsOfUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribePermissionsOfUser")
	return
}

func NewDescribePermissionsOfUserResponse() (response *DescribePermissionsOfUserResponse) {
	response = &DescribePermissionsOfUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个用户的所有权限
func (c *Client) DescribePermissionsOfUser(request *DescribePermissionsOfUserRequest) (response *DescribePermissionsOfUserResponse, err error) {
	if request == nil {
		request = NewDescribePermissionsOfUserRequest()
	}
	response = NewDescribePermissionsOfUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserSQLRequest() (request *DescribeUserSQLRequest) {
	request = &DescribeUserSQLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "DescribeUserSQL")
	return
}

func NewDescribeUserSQLResponse() (response *DescribeUserSQLResponse) {
	response = &DescribeUserSQLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeUserSQL)用于dmc拉取某条用户SQL信息详情。
func (c *Client) DescribeUserSQL(request *DescribeUserSQLRequest) (response *DescribeUserSQLResponse, err error) {
	if request == nil {
		request = NewDescribeUserSQLRequest()
	}
	response = NewDescribeUserSQLResponse()
	err = c.Send(request, response)
	return
}

func NewGrantRoleToUserRequest() (request *GrantRoleToUserRequest) {
	request = &GrantRoleToUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dmc", APIVersion, "GrantRoleToUser")
	return
}

func NewGrantRoleToUserResponse() (response *GrantRoleToUserResponse) {
	response = &GrantRoleToUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将角色授予用户
func (c *Client) GrantRoleToUser(request *GrantRoleToUserRequest) (response *GrantRoleToUserResponse, err error) {
	if request == nil {
		request = NewGrantRoleToUserRequest()
	}
	response = NewGrantRoleToUserResponse()
	err = c.Send(request, response)
	return
}
