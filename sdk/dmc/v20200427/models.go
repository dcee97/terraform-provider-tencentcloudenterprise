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
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type EnableAuthUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableAuthUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableAuthUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantRoleToUserRequest struct {
	*tchttp.BaseRequest

	// 角色名

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 账号Id

	User *string `json:"User,omitempty" name:"User"`
}

func (r *GrantRoleToUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantRoleToUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIpWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIpWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePermissionApplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePermissionApplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePermissionApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LeaseDmcResourceRouteRequest struct {
	*tchttp.BaseRequest

	// 需要续期的资源ID 格式如：dmc-c1nl9rpv，与DMC控制台页面中显示的资源ID相同。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *LeaseDmcResourceRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LeaseDmcResourceRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthUser struct {

	// appid信息

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 子账号/协作者账号id

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 是否为主账号，1表示是主账号，0表示非主账号

	IsMain *int64 `json:"IsMain,omitempty" name:"IsMain"`
	// 是否管理员，1表示是管理员，0表示普通用户

	IsAdmin *int64 `json:"IsAdmin,omitempty" name:"IsAdmin"`
	// 状态码，0表示正常，1表示被禁用

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 用户角色列表

	Roles []*AuthRole `json:"Roles,omitempty" name:"Roles"`
}

type AddUserSQLRequest struct {
	*tchttp.BaseRequest

	// 用户SQL记录名称。

	UserSQLName *string `json:"UserSQLName,omitempty" name:"UserSQLName"`
	// 所属产品名称: cdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 用户sql内容。

	SQLContext *string `json:"SQLContext,omitempty" name:"SQLContext"`
	// 可见范围：全局，实例， db级别。

	VisualRange *string `json:"VisualRange,omitempty" name:"VisualRange"`
	// 实例ID，当可见范围非全局时候，必须传递。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库名成，当可见范围是db级别时候，必须传递。

	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *AddUserSQLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserSQLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDmcResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的资源总条数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的资源信息 。

		Items []*DmcResource `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDmcResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDmcResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDmcResourceUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dmc 代理IP。

		ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
		// dmc代理端口。

		ProxyPort *uint64 `json:"ProxyPort,omitempty" name:"ProxyPort"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDmcResourceUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDmcResourceUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePermissionsOfUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 此用户的所有权限

		UserPermissions []*UserPermission `json:"UserPermissions,omitempty" name:"UserPermissions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePermissionsOfUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePermissionsOfUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantPermissionToUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantPermissionToUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantPermissionToUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevokePermissionFromUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevokePermissionFromUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RevokePermissionFromUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchIpWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchIpWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchIpWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDmcResourceUsersRequest struct {
	*tchttp.BaseRequest

	// dmc资源id ，列表过滤使用，例如：dmc-12qjsp

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// dmc资源用户，列表过滤使用。

	UserNames []*string `json:"UserNames,omitempty" name:"UserNames"`
	// 列表排序字段，目前仅支持：created_at。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 列表排序顺序，asc,desc,ASC,DESC。

	OrderByType []*string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 批量拉取限制条数,范围：(0, 1000]。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 批量拉取偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDmcResourceUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDmcResourceUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDmcResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyDmcResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyDmcResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAuthUserRequest struct {
	*tchttp.BaseRequest

	// 要被删除的账号（子账号或协作者账号）

	User *string `json:"User,omitempty" name:"User"`
}

func (r *DeleteAuthUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAuthUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePermissionsOfInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的权限数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 当前用户对于此实例的权限列表

		UserPermissions []*UserPermission `json:"UserPermissions,omitempty" name:"UserPermissions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePermissionsOfInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePermissionsOfInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAuthUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuthUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAuthUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDmcResourceUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的数据总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// dmc 资源用户信息。

		Items []*DmcResourceUser `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDmcResourceUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDmcResourceUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserSQLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserSQLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserSQLRequest struct {
	*tchttp.BaseRequest

	// 用户sql名称。

	UserSQLName *string `json:"UserSQLName,omitempty" name:"UserSQLName"`
	// 所属产品名称: cdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 用户SQL 唯一ID。

	UserSQLID *uint64 `json:"UserSQLID,omitempty" name:"UserSQLID"`
}

func (r *DeleteUserSQLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserSQLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDmcResourcesRequest struct {
	*tchttp.BaseRequest

	// dmc 资源类型，当前仅支持mysql。

	DatabaseKinds []*string `json:"DatabaseKinds,omitempty" name:"DatabaseKinds"`
	// dmc 资源IP。

	ResourceIps []*string `json:"ResourceIps,omitempty" name:"ResourceIps"`
	// dmc资源端口。

	ResourcePorts []*uint64 `json:"ResourcePorts,omitempty" name:"ResourcePorts"`
	// dmc资源Id，例如：dmc-jhdqrs。

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// dmc资源类型：cdb云实例，cvm 自建实例， extranet 外网

	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes"`
	// 搜索key，目前仅支持：resource_name。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 搜索值，配合SearchKey 使用。

	SearchKeyWords *string `json:"SearchKeyWords,omitempty" name:"SearchKeyWords"`
	// 排序key， created_at 。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序字段 asc,desc,ASC,DESC 。

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 限制条数，(0, 1000]。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 拉取偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDmcResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDmcResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSQLMateInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 元数据信息。

		Items []*UserSQLMeta `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserSQLMateInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSQLMateInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LeaseDmcResourceRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 代理IP

		ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
		// 代理端口

		ProxyPort *uint64 `json:"ProxyPort,omitempty" name:"ProxyPort"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LeaseDmcResourceRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LeaseDmcResourceRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Permission struct {

	// 子账号/协作者账号id

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// dmc实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 操作

	AuthAction *string `json:"AuthAction,omitempty" name:"AuthAction"`
}

type DestroyDmcResourceUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyDmcResourceUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyDmcResourceUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevokePermissionFromUserRequest struct {
	*tchttp.BaseRequest

	// 要被收回的用户权限列表

	RevokePermissions []*Permission `json:"RevokePermissions,omitempty" name:"RevokePermissions"`
}

func (r *RevokePermissionFromUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RevokePermissionFromUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCurrentUserRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCurrentUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCurrentUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckPermissionApplyRequest struct {
	*tchttp.BaseRequest

	// 权限申请id

	PermissionRequestId *int64 `json:"PermissionRequestId,omitempty" name:"PermissionRequestId"`
	// approve表示通过，reject表示拒绝，枚举范围: approve,reject

	CheckAction *string `json:"CheckAction,omitempty" name:"CheckAction"`
	// 拒绝理由，如果checkAction是reject，可以带上

	RejectReason *string `json:"RejectReason,omitempty" name:"RejectReason"`
}

func (r *CheckPermissionApplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckPermissionApplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllPermissionsOfInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数字范围: [0, INF)

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// (0, INF)

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAllPermissionsOfInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllPermissionsOfInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserSQLInfo struct {

	// dmc用户SQL记录唯一ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// dmc用户SQL记录所属用户应用ID。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// dmc用户SQL记录名称。

	UserSQLName *string `json:"UserSQLName,omitempty" name:"UserSQLName"`
	// dmc用户SQL记录创建时间。

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// dmc用户SQL创建UIN。

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// dmc用户SQL所属产品。

	Product *string `json:"Product,omitempty" name:"Product"`
	// dmc用户SQL 记录详情。

	UserSQLContext *string `json:"UserSQLContext,omitempty" name:"UserSQLContext"`
	// 可见范围: all, instance, db。

	VisualRange *string `json:"VisualRange,omitempty" name:"VisualRange"`
	// 所属实例ID 例如：cdb-9Rhlpq。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 所属实例database名称。

	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

type ModifyDmcResourceVipVportRequest struct {
	*tchttp.BaseRequest

	// 资源Id 。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源IP。

	ResourceIp *string `json:"ResourceIp,omitempty" name:"ResourceIp"`
	// 资源端口。

	ResourcePort *uint64 `json:"ResourcePort,omitempty" name:"ResourcePort"`
	// 资源所属vpc 网络英文Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 资源所属vpc subne 英文id。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *ModifyDmcResourceVipVportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDmcResourceVipVportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthRolesRequest struct {
	*tchttp.BaseRequest

	// 数字范围: [0, INF)

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数字范围: (0, INF)

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAuthRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuthRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpWhitelistsRequest struct {
	*tchttp.BaseRequest

	// 数字范围: [0, INF)

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数字范围: (0, INF)

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIpWhitelistsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpWhitelistsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSQLRequest struct {
	*tchttp.BaseRequest

	// 产品名称： cdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 用户SQL主键ID。

	UserSQLID *uint64 `json:"UserSQLID,omitempty" name:"UserSQLID"`
	// 用户SQL 名称。

	UserSQLName *string `json:"UserSQLName,omitempty" name:"UserSQLName"`
}

func (r *DescribeUserSQLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSQLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantPermissionToUserRequest struct {
	*tchttp.BaseRequest

	// 到期时间，精确到秒

	ExpireAt *string `json:"ExpireAt,omitempty" name:"ExpireAt"`
	// 要授权的用户权限列表

	GrantPermissions []*Permission `json:"GrantPermissions,omitempty" name:"GrantPermissions"`
}

func (r *GrantPermissionToUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantPermissionToUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePermissionsOfUserRequest struct {
	*tchttp.BaseRequest

	// 账号（子账号或协作者账号）

	User *string `json:"User,omitempty" name:"User"`
	// 数字范围: [0, INF)

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数字范围: (0, INF)

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePermissionsOfUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePermissionsOfUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBConnection struct {

	// 实例id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 用户名

	User *string `json:"User,omitempty" name:"User"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 选择的数据库名

	DBName *string `json:"DBName,omitempty" name:"DBName"`
	// 字符集，例如:utf8

	Charset *string `json:"Charset,omitempty" name:"Charset"`
	// 排序规则，例如:utf8_general_ci

	Collation *string `json:"Collation,omitempty" name:"Collation"`
	// Dial timeout,暂不超时设置为0

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
	// I/O read timeout,暂不超时设置为0

	ReadTimeout *int64 `json:"ReadTimeout,omitempty" name:"ReadTimeout"`
	// I/O write timeout,暂不超时设置为0

	WriteTimeout *int64 `json:"WriteTimeout,omitempty" name:"WriteTimeout"`
}

type DeleteIpWhitelistRequest struct {
	*tchttp.BaseRequest

	// 将被删除的白名单IP

	IpDelete *string `json:"IpDelete,omitempty" name:"IpDelete"`
}

func (r *DeleteIpWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIpWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSQLsRequest struct {
	*tchttp.BaseRequest

	// 产品名称： cdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 用户sqlName。

	UserSQLName []*string `json:"UserSQLName,omitempty" name:"UserSQLName"`
	// 搜索条件。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 搜索值。

	SearchKeyWords *string `json:"SearchKeyWords,omitempty" name:"SearchKeyWords"`
	// 分页偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页条数。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUserSQLsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSQLsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIpWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIpWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIpWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableAuthUserRequest struct {
	*tchttp.BaseRequest

	// 要被禁用的账号（子账号或协作者账号）

	User *string `json:"User,omitempty" name:"User"`
}

func (r *DisableAuthUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableAuthUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantRoleToUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantRoleToUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantRoleToUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckDmcResourceRouteRequest struct {
	*tchttp.BaseRequest

	// 需要检测的资源ID 比如修改资源信息之前一般会调用传递 格式如：dmc-c1nl9rpv，与DMC控制台页面中显示的资源ID相同。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 源实例ID 例如：cdb-c1nl9rpv 或者 ins-c1nl9pv,外网实例可能没有源实例Id。

	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`
	// 数据库类型： mysql 等。

	DatabaseKind *string `json:"DatabaseKind,omitempty" name:"DatabaseKind"`
	// 源实例资源类型：cdb云实例，cvm 自建实例， extranet 外网。

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源IP。

	ResourceIp *string `json:"ResourceIp,omitempty" name:"ResourceIp"`
	// 资源Port。

	ResourcePort *uint64 `json:"ResourcePort,omitempty" name:"ResourcePort"`
	// vpc资源ID， 例如：vpc-c1nl9rp。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc资源子网ID 例如：subnet-c1nl9rp。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 资源登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 资源登录用户密码。

	UserPassword *string `json:"UserPassword,omitempty" name:"UserPassword"`
}

func (r *CheckDmcResourceRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckDmcResourceRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableAuthUserRequest struct {
	*tchttp.BaseRequest

	// 要被启用的账号（子账号或协作者账号）

	User *string `json:"User,omitempty" name:"User"`
}

func (r *EnableAuthUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableAuthUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIpWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIpWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIpWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthRole struct {

	// 角色类型ID

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 角色名称

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type SqlResult struct {

	// 列信息

	Columns []*Column `json:"Columns,omitempty" name:"Columns"`
	// 数据行

	Rows []*string `json:"Rows,omitempty" name:"Rows"`
}

type UserSQLMeta struct {

	// dmc用户SQL记录唯一ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// dmc用户SQL记录所属用户应用ID。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// dmc用户SQL记录名称。

	UserSQLName *string `json:"UserSQLName,omitempty" name:"UserSQLName"`
	// dmc用户SQL记录创建时间。

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// dmc用户SQL创建UIN。

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// dmc用户SQL所属产品。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 可见范围: all, instance, db。

	VisualRange *string `json:"VisualRange,omitempty" name:"VisualRange"`
	// 所属实例ID 例如：cdb-9Rhlpq。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 所属实例database名称。

	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

type AddUserSQLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserSQLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpWhitelistsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IP白名单功能是否开启，on开启，off关闭

		Switch *string `json:"Switch,omitempty" name:"Switch"`
		// 白名单IP的数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 当前appid下白名单列表

		IpWhitelist []*IpWhitelistUtil `json:"IpWhitelist,omitempty" name:"IpWhitelist"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpWhitelistsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpWhitelistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PermissionRequestUtil struct {

	// 权限申请ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// appid信息

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 申请单创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// dmc实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 操作

	AuthAction *string `json:"AuthAction,omitempty" name:"AuthAction"`
	// 申请理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 取消/通过/拒绝操作的操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 状态为拒绝时的拒绝理由

	RejectReason *string `json:"RejectReason,omitempty" name:"RejectReason"`
	// 此申请单的状态，initial表示新提交，cancel表示被取消，approve表示通过，reject表示被拒绝

	Status *string `json:"Status,omitempty" name:"Status"`
	// 申请权限的到期时间

	ExpireAt *string `json:"ExpireAt,omitempty" name:"ExpireAt"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type DescribeAuthUsersRequest struct {
	*tchttp.BaseRequest

	// 数字范围: [0, INF

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数字范围: (0, INF)

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAuthUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuthUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDmcResourceNameRequest struct {
	*tchttp.BaseRequest

	// dmc资源Id，例如：dmc-jhdpsw。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// dmc资源修改名称。

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

func (r *ModifyDmcResourceNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDmcResourceNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DmcResourceUser struct {

	// dmc资源用户信息ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// dmc资源ID 例如： dmc-1jdfp7of

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// dmc资源用户信息创建uin

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// dmc资源用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// dmc资源用户创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// dmc资源用户更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// dmc资源用户密码

	UserPassword *string `json:"UserPassword,omitempty" name:"UserPassword"`
}

type DescribeAuthRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 角色列表

		AuthRoles []*AuthRole `json:"AuthRoles,omitempty" name:"AuthRoles"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuthRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCurrentUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户id

		AppId *int64 `json:"AppId,omitempty" name:"AppId"`
		// 子账号/协作者账号id

		Uin *string `json:"Uin,omitempty" name:"Uin"`
		// 是否为主账号，1表示是主账号，0表示非主账号

		IsMain *int64 `json:"IsMain,omitempty" name:"IsMain"`
		// 是否管理员，1表示是管理员，0表示普通用户

		IsAdmin *int64 `json:"IsAdmin,omitempty" name:"IsAdmin"`
		// 状态码，0表示正常，1表示被禁用

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 用户的角色

		Roles []*AuthRole `json:"Roles,omitempty" name:"Roles"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCurrentUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCurrentUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSQLMateInfosRequest struct {
	*tchttp.BaseRequest

	// 产品名称： cdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 用户sqlName。

	UserSQLName []*string `json:"UserSQLName,omitempty" name:"UserSQLName"`
	// 搜索条件。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 搜索值。

	SearchKeyWords *string `json:"SearchKeyWords,omitempty" name:"SearchKeyWords"`
	// 分页偏移量。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页条数。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUserSQLMateInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSQLMateInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDmcResourceVipVportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 代理IP 。

		ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
		// 代理端口 。

		ProxyPort *uint64 `json:"ProxyPort,omitempty" name:"ProxyPort"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDmcResourceVipVportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDmcResourceVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserSQLRequest struct {
	*tchttp.BaseRequest

	// 所属产品名称: cdb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 用户sql主键ID。

	UserSQLID *uint64 `json:"UserSQLID,omitempty" name:"UserSQLID"`
	// 用户sql名称。

	UserSQLName *string `json:"UserSQLName,omitempty" name:"UserSQLName"`
	// 用户sql内容。

	SQLContext *string `json:"SQLContext,omitempty" name:"SQLContext"`
	// 可见范围：全局，实例， db级别。

	VisualRange *string `json:"VisualRange,omitempty" name:"VisualRange"`
	// 实例ID，当可见范围非全局时候，必须传递。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库名成，当可见范围是db级别时候，必须传递。

	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *ModifyUserSQLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserSQLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DmcResource struct {

	// dmc资源唯一整型ID。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// dmc资源应用ID。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// dmc资源创建用户Uin。

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// dmc资源数据库种类。

	DatabaseKind *string `json:"DatabaseKind,omitempty" name:"DatabaseKind"`
	// dmc资源类型：cloud/cvm/other。

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// dmc资源id 例如：dmc-yf8912h 。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// dmc资源名称 。

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// dmc资源源实例ID，例如：cdb-hjdfsd 或者 ins-dfsjlj。

	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`
	// dmc资源依赖源实例的vpc 英文id。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// dmc资源依赖源实例的vpc子网id。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// dmc资源所在区域信息，例如：ap-guangzhou。

	Region *string `json:"Region,omitempty" name:"Region"`
	// dmc资源代理IP，后端绑定路由是ResourceIp。

	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
	// dmc资源代理端口。

	ProxyPort *uint64 `json:"ProxyPort,omitempty" name:"ProxyPort"`
	// dmc资源代理唯一Id。

	ProxyId *uint64 `json:"ProxyId,omitempty" name:"ProxyId"`
	// dmc资源IP。

	ResourceIp *string `json:"ResourceIp,omitempty" name:"ResourceIp"`
	// dmc资源端口。

	ResourcePort *uint64 `json:"ResourcePort,omitempty" name:"ResourcePort"`
	// dmc资源代理有效标示符。

	ProxyIsValid *bool `json:"ProxyIsValid,omitempty" name:"ProxyIsValid"`
	// dmc资源代理续期时间。

	LeaseAt *string `json:"LeaseAt,omitempty" name:"LeaseAt"`
	// dmc资源创建时间。

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// dmc资源更新时间。

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// dmc资源绑定用户信息。

	UserItems []*DmcResourceUsers `json:"UserItems,omitempty" name:"UserItems"`
}

type DescribeAuthUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户列表

		AuthUsers []*AuthUser `json:"AuthUsers,omitempty" name:"AuthUsers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuthUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserPermission struct {

	// appid信息

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 子账号/协作者账号id

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// dmc实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 操作

	AuthAction *string `json:"AuthAction,omitempty" name:"AuthAction"`
	// 权限到期时间

	ExpireAt *string `json:"ExpireAt,omitempty" name:"ExpireAt"`
}

type CheckPermissionApplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckPermissionApplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckPermissionApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableAuthUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableAuthUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableAuthUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDmcResourceRequest struct {
	*tchttp.BaseRequest

	// 需要销毁的资源Id，例如：dmc-jdhrh3

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DestroyDmcResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyDmcResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDmcResourceUserRequest struct {
	*tchttp.BaseRequest

	// dmc资源用户唯一标示ID。

	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
	// dmc资源ID 例如：dmc-jdfqph。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// dmc资源ip。

	ResourceIp *string `json:"ResourceIp,omitempty" name:"ResourceIp"`
	// dmc资源端口。

	ResourcePort *uint64 `json:"ResourcePort,omitempty" name:"ResourcePort"`
	// dmc资源vpc 英文ID, 例如： vpc-jpdhr0。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// dmc资源vpc 子网Id,例如：subnet-spdj19q。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// dmc资源用户名称。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// dmc资源用户密码。

	UserPassword *string `json:"UserPassword,omitempty" name:"UserPassword"`
}

func (r *ModifyDmcResourceUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDmcResourceUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchIpWhitelistRequest struct {
	*tchttp.BaseRequest

	// on表示开，off表示关

	SwitchAction *string `json:"SwitchAction,omitempty" name:"SwitchAction"`
}

func (r *SwitchIpWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchIpWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDmcResourceUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 代理IP。

		ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
		// 代理端口。

		ProxyPort *uint64 `json:"ProxyPort,omitempty" name:"ProxyPort"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDmcResourceUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDmcResourceUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpWhitelistUtil struct {

	// 白名单id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 白名单IP或ICDR

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
}

type CreateDmcResourceUserRequest struct {
	*tchttp.BaseRequest

	// 需要创建用户信息的资源ID 格式如：dmc-c1nl9rpv，与DMC控制台页面中显示的资源ID相同。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源IP。

	ResourceIp *string `json:"ResourceIp,omitempty" name:"ResourceIp"`
	// 资源Port。

	ResourcePort *uint64 `json:"ResourcePort,omitempty" name:"ResourcePort"`
	// vpc资源ID， 例如：vpc-c1nl9rp。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc资源子网ID 例如：subnet-c1nl9rp。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 资源登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 资源登录用户密码。

	UserPassword *string `json:"UserPassword,omitempty" name:"UserPassword"`
}

func (r *CreateDmcResourceUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDmcResourceUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePermissionApplyRequest struct {
	*tchttp.BaseRequest

	// dmc实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 申请的实例操作，如DDL/DML/DQL

	AuthActions []*string `json:"AuthActions,omitempty" name:"AuthActions"`
	// 申请理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 到期时间，精确到秒

	ExpireAt *string `json:"ExpireAt,omitempty" name:"ExpireAt"`
}

func (r *CreatePermissionApplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePermissionApplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllPermissionsOfInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回权限的数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 此实例上的所有权限

		UserPermissions []*UserPermission `json:"UserPermissions,omitempty" name:"UserPermissions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllPermissionsOfInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllPermissionsOfInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {

	// 列名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 列数据库类型

	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
}

type CheckDmcResourceRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// telnet 检测。

		HealthTelnet *bool `json:"HealthTelnet,omitempty" name:"HealthTelnet"`
		// 连接检测。

		HealthConnect *bool `json:"HealthConnect,omitempty" name:"HealthConnect"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckDmcResourceRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckDmcResourceRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserSQLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserSQLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DmcResourceUsers struct {

	// dmc资源用户信息ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// dmc资源ID 例如： dmc-1jdfp7of

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// dmc资源用户信息创建uin

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// dmc资源用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// dmc资源用户创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// dmc资源用户更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// dmc资源用户密码是否存在

	PasswordIsExists *bool `json:"PasswordIsExists,omitempty" name:"PasswordIsExists"`
	// dmc资源用户密码

	UserPassword *string `json:"UserPassword,omitempty" name:"UserPassword"`
}

type DescribePermissionsOfInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePermissionsOfInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePermissionsOfInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteSqlRequest struct {
	*tchttp.BaseRequest

	// 执行的SQL语句

	Sql *string `json:"Sql,omitempty" name:"Sql"`
	// SQL语句是否使用Base64编码，需要为true

	Base64Encode *bool `json:"Base64Encode,omitempty" name:"Base64Encode"`
	// DB连接信息

	DBConnection *DBConnection `json:"DBConnection,omitempty" name:"DBConnection"`
	// 系统offset,用于DQL语句

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 系统limit,用于DQL语句

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ExecuteSqlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteSqlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDmcResourceNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDmcResourceNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDmcResourceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIpWhitelistRequest struct {
	*tchttp.BaseRequest

	// 被修改的IP白名单的Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 修改后的白名单IP

	NewIp *string `json:"NewIp,omitempty" name:"NewIp"`
	// 修改后的白名单备注

	NewNote *string `json:"NewNote,omitempty" name:"NewNote"`
}

func (r *ModifyIpWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAuthUserRequest struct {
	*tchttp.BaseRequest

	// 要被添加的账号（子账号或协作者账号）

	User *string `json:"User,omitempty" name:"User"`
	// 用户类型，subAccount表示子账号，cooperator表示协作者账号

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用户昵称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色Id

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
}

func (r *CreateAuthUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAuthUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIpWhitelistRequest struct {
	*tchttp.BaseRequest

	// ip或ip段列表

	IpAddList []*string `json:"IpAddList,omitempty" name:"IpAddList"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
}

func (r *CreateIpWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIpWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAuthUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAuthUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAuthUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDmcResourceUserRequest struct {
	*tchttp.BaseRequest

	// dmc资源用户唯一标示ID。

	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
	// dmc资源ID，例如：dmc-1qjrpq。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 需要销毁的dmc资源用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DestroyDmcResourceUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyDmcResourceUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteSqlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SQL语句类型

		SqlStmtType *string `json:"SqlStmtType,omitempty" name:"SqlStmtType"`
		// 受影响的行数

		RowsAffected *int64 `json:"RowsAffected,omitempty" name:"RowsAffected"`
		// 最后InsertId

		LastInsertId *int64 `json:"LastInsertId,omitempty" name:"LastInsertId"`
		// 执行耗时，单位s(秒)

		CostTime *string `json:"CostTime,omitempty" name:"CostTime"`
		// SQl执行结果

		SqlResult *SqlResult `json:"SqlResult,omitempty" name:"SqlResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteSqlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSQLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dmc 用户SQL 数据信息。

		UserSQLInfo *UserSQLInfo `json:"UserSQLInfo,omitempty" name:"UserSQLInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserSQLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSQLsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据信息。

		Items []*UserSQLInfo `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserSQLsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserSQLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
