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
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type SearchTaskResultDetailRequest struct {
	*tchttp.BaseRequest

	// 运维任务日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *SearchTaskResultDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchTaskResultDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShowTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 活跃用户、主机、高危用户和在线用户

		Top *TopResult `json:"Top,omitempty" name:"Top"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShowTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShowTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSyncStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产同步结果

		Status *AssetSyncStatus `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSyncStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest

	// 地域码, 如: ap-guangzhou

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 按照堡垒机开通的 VPC 实例ID查询

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 资源ID集合，当传入ID集合时忽略 ApCode 和 VpcId

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetSyncFlagRequest struct {
	*tchttp.BaseRequest

	// 是否开启资产自动同步，false-不开启，true-开启

	AutoSync *bool `json:"AutoSync,omitempty" name:"AutoSync"`
}

func (r *ModifyAssetSyncFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetSyncFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 文件和所属会话

		FileSessionSet []*SessionFile `json:"FileSessionSet,omitempty" name:"FileSessionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFileSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTaskResultDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维任务执行结果详情

		TaskResultDetail *TaskResultDetail `json:"TaskResultDetail,omitempty" name:"TaskResultDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchTaskResultDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchTaskResultDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetSyncFlags struct {

	// 是否已完成角色授权

	RoleGranted *bool `json:"RoleGranted,omitempty" name:"RoleGranted"`
	// 是否已开启自动资产同步

	AutoSync *bool `json:"AutoSync,omitempty" name:"AutoSync"`
}

type AuditLogResult struct {

	// 被审计会话的Sid

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 审计者的编号

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 审计动作发生的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 审计者的Ip

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 审计动作类型，1--回放、2--中断、3--监控

	Operation *int64 `json:"Operation,omitempty" name:"Operation"`
	// 被审计主机的Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 被审计主机的主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 被审计会话所属的类型，如字符会话

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 被审计主机的内部Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 被审计主机的外部Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 审计者的子账号

	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type Resource struct {

	// 服务实例ID，如bh-saas-s3ed4r5e

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域编码

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 服务实例规格信息

	SvArgs *string `json:"SvArgs,omitempty" name:"SvArgs"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 服务规格对应的资产数

	Nodes *uint64 `json:"Nodes,omitempty" name:"Nodes"`
	// 自动续费标记，0 - 表示默认状态，1 - 表示自动续费，2 - 表示明确不自动续费

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 资源状态，0 - 未初始化，1 - 正常，2 - 隔离，3 - 销毁，4 - 初始化失败，5 - 初始化中

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 服务实例名，如T-Sec-堡垒机（SaaS型）

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 定价模型ID

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 资源创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 商品码, p_cds_dasb

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码, sp_cds_dasb_bh_saas

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 是否过期，true-过期，false-未过期

	Expired *bool `json:"Expired,omitempty" name:"Expired"`
	// 是否开通，true-开通，false-未开通

	Deployed *bool `json:"Deployed,omitempty" name:"Deployed"`
	// 开通服务的 VPC 名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 开通服务的 VPC 对应的网段

	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`
	// 开通服务的子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 开通服务的子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 开通服务的子网网段

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 外部IP

	PublicIpSet []*string `json:"PublicIpSet,omitempty" name:"PublicIpSet"`
	// 内部IP

	PrivateIpSet []*string `json:"PrivateIpSet,omitempty" name:"PrivateIpSet"`
	// 服务开通的高级功能列表，如:[DB]

	ModuleSet []*string `json:"ModuleSet,omitempty" name:"ModuleSet"`
	// 已使用的授权点数

	UsedNodes *uint64 `json:"UsedNodes,omitempty" name:"UsedNodes"`
	// 扩展点数

	ExtendPoints *uint64 `json:"ExtendPoints,omitempty" name:"ExtendPoints"`
	// 带宽扩展包个数(4M)

	PackageBandwidth *uint64 `json:"PackageBandwidth,omitempty" name:"PackageBandwidth"`
	// 授权点数扩展包个数(50点)

	PackageNode *uint64 `json:"PackageNode,omitempty" name:"PackageNode"`
}

type ModifyLogOutputSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogOutputSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogOutputSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagFilter struct {

	// 标签key

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签value

	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type ModifyLoginSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileSessionRequest struct {
	*tchttp.BaseRequest

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 主机内网Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 主机外网Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 来源Ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 文件名

	File *string `json:"File,omitempty" name:"File"`
	// 对文件的动作，1为上传，2为下载，3为删除，4为重命名

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 只能取 "SFTP"、"RDP"、"rz/sz"

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页中一页内的记录数，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchFileSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthModeSetting struct {

	// 双因子认证，0-不开启，1-OTP，2-短信

	AuthMode *uint64 `json:"AuthMode,omitempty" name:"AuthMode"`
}

type EventResult struct {

	// 会话Id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 地域信息

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 主机id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 高危命令或文件传输风险操作的文件名

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 高危命令执行时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 事件类型，如shell命令、文件传输等

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 各Type所属子类型，如:Type为文件传输时，此字段代表文件操作类型

	Method *int64 `json:"Method,omitempty" name:"Method"`
}

type DeleteDeviceGroupsRequest struct {
	*tchttp.BaseRequest

	// 待删除的资产组ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteDeviceGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetSyncFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetSyncFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetSyncFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAclResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功的访问权限ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginEvent struct {

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 操作时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 来源IP

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 登录入口：1-字符界面,2-图形界面，3-web页面, 4-API

	Entry *uint64 `json:"Entry,omitempty" name:"Entry"`
	// 操作结果，1-成功，2-失败

	Result *uint64 `json:"Result,omitempty" name:"Result"`
}

type DescribeDeviceGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产组成员总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产组成员列表

		DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSyncStatusRequest struct {
	*tchttp.BaseRequest

	// 查询的资产同步类型。1 -主机资产， 2 - 数据库资产

	Category *uint64 `json:"Category,omitempty" name:"Category"`
}

func (r *DescribeAssetSyncStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginEventRequest struct {
	*tchttp.BaseRequest

	// 用户名，如果不包含其他条件时对user_name or real_name两个字段模糊查询

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名，模糊查询

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 查询时间范围，起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询时间范围，结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 来源IP，模糊查询

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 登录入口：1-字符界面,2-图形界面，3-web页面, 4-API

	Entry *uint64 `json:"Entry,omitempty" name:"Entry"`
	// 操作结果，1-成功，2-失败

	Result *uint64 `json:"Result,omitempty" name:"Result"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页每页记录数，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLoginEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceAccountPrivateKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetDeviceAccountPrivateKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDeviceAccountPrivateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaySessionRequest struct {
	*tchttp.BaseRequest

	// 会话Sid

	Sid *string `json:"Sid,omitempty" name:"Sid"`
}

func (r *ReplaySessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaySessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileInformation struct {

	// 文件名

	File *string `json:"File,omitempty" name:"File"`
	// 对文件的动作，1为上传，2为下载，3为删除，4为重命名

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 操作文件的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 文件的大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 操作文件时所用的协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type KeyCount struct {

	// 用户名或日期

	Key *string `json:"Key,omitempty" name:"Key"`
	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type KillSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 终端会话所需信息

		ReplayInfo *ReplayInformation `json:"ReplayInfo,omitempty" name:"ReplayInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KillSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KillSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDepartmentRequest struct {
	*tchttp.BaseRequest

	// 部门名称，1 - 256个字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 上级部门的ID

	Parent *string `json:"Parent,omitempty" name:"Parent"`
	// 部门管理员的账号ID

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
}

func (r *CreateDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KillSessionRequest struct {
	*tchttp.BaseRequest

	// 会话Sid

	Sid *string `json:"Sid,omitempty" name:"Sid"`
}

func (r *KillSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KillSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplayInformation struct {

	// 令牌

	Token *string `json:"Token,omitempty" name:"Token"`
	// 会话开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 回放链接

	Address *string `json:"Address,omitempty" name:"Address"`
}

type ImportDeviceAccountRequest struct {
	*tchttp.BaseRequest

	// 账号列表

	DeviceAccountSet []*ParamDeviceAccount `json:"DeviceAccountSet,omitempty" name:"DeviceAccountSet"`
}

func (r *ImportDeviceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDeviceAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseUsageRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLicenseUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseUsageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmdTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命令模板列表

		CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitempty" name:"CmdTemplateSet"`
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmdTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmdTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicesPortRequest struct {
	*tchttp.BaseRequest

	// 主机记录ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

func (r *ModifyDevicesPortRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesPortRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListStatusRequest struct {
	*tchttp.BaseRequest

	// true：放开全部来源IP；false：不放开全部来源IP

	AllowAny *bool `json:"AllowAny,omitempty" name:"AllowAny"`
}

func (r *ModifyAccessWhiteListStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskResultDetail struct {

	// 运维任务结果日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务类型 1 - 手工任务， 2 - 周期性任务

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 运维任务来源IP

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 运维任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运维任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 运维任务创建人员用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 运维任务创建人员姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 运维任务执行账户

	Account *string `json:"Account,omitempty" name:"Account"`
	// 运维任务超时时间, 单位秒

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 运维任务命令

	Script *string `json:"Script,omitempty" name:"Script"`
}

type SearchAuditLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审计日志

		AuditLogSet []*AuditLogResult `json:"AuditLogSet,omitempty" name:"AuditLogSet"`
		// 日志总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchAuditLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAuditLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 资产公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 资产内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 操作类型：1 - 文件上传，2 - 文件下载，3 - 文件删除，4 - 文件(夹)移动，5 - 文件(夹)重命名，6 - 新建文件夹，9 - 删除文件夹

	Method []*uint64 `json:"Method,omitempty" name:"Method"`
	// 可填写路径名或文件（夹）名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 1-已执行， 2-被阻断

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 登录日志列表

		LoginEventSet []*LoginEvent `json:"LoginEventSet,omitempty" name:"LoginEventSet"`
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 主机账号ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机账号密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *BindDeviceAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 用户组名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户组所属的部门ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyUserGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLDAPSyncFlagRequest struct {
	*tchttp.BaseRequest
}

func (r *SetLDAPSyncFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLDAPSyncFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功后返回的记录ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessWhiteListRulesRequest struct {
	*tchttp.BaseRequest

	// 用户ID集合，非必需，如果使用IdSet参数则忽略Name参数

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 来源IP或网段，模糊查询，最大长度64字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页，偏移位置，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccessWhiteListRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessWhiteListRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户列表

		UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetUserPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetUserPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetUserPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceRequest struct {
	*tchttp.BaseRequest

	// 需要开通服务的资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 1 - 开通， 2 - 关闭

	Status *string `json:"Status,omitempty" name:"Status"`
	// 开通或关闭的高级特性列表

	ModuleSet []*string `json:"ModuleSet,omitempty" name:"ModuleSet"`
}

func (r *ModifyResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Command struct {

	// 命令

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 命令输入的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 命令执行时间相对于所属会话开始时间的偏移量，单位ms

	TimeOffset *uint64 `json:"TimeOffset,omitempty" name:"TimeOffset"`
	// 命令执行情况，1--允许，2--拒绝，3--确认

	Action *int64 `json:"Action,omitempty" name:"Action"`
}

type PwdLog struct {

	// 用户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 订单关联标识

	IncKey *string `json:"IncKey,omitempty" name:"IncKey"`
}

type DescribeCmdTemplatesRequest struct {
	*tchttp.BaseRequest

	// 命令模板ID集合，非必需

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 命令模板名，模糊查询，最大长度64字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页，偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCmdTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmdTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationTaskDetail struct {

	// 运维任务主键ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 任务类型，1 - 手工执行任务， 2 - 周期性任务

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 周期性任务执行间隔，单位天

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 首次执行时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 执行账户

	Account *string `json:"Account,omitempty" name:"Account"`
	// 任务超时时间，单位秒

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 任务命令

	Script *string `json:"Script,omitempty" name:"Script"`
	// 执行任务的资产信息

	Devices []*Device `json:"Devices,omitempty" name:"Devices"`
	// 下次任务执行时间

	NextTime *string `json:"NextTime,omitempty" name:"NextTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
}

type ModifyDasbResourceRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例id bh-saas-xxxx

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源类型。标准版(standard)或者专业版(pro)

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 资源默认节点数。取值：10/20/50/100/200/500/1000

	ResourceNode *uint64 `json:"ResourceNode,omitempty" name:"ResourceNode"`
	// 带宽扩展包数量

	MbpExp *uint64 `json:"MbpExp,omitempty" name:"MbpExp"`
	// 节点扩展包数量

	NodeExp *uint64 `json:"NodeExp,omitempty" name:"NodeExp"`
}

func (r *ModifyDasbResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDasbResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 会话信息列表

		SessionSet []*SessionResult `json:"SessionSet,omitempty" name:"SessionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceGroupRequest struct {
	*tchttp.BaseRequest

	// 资产组名，最大长度32字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资产组所属部门ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreateDeviceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDepartmentRequest struct {
	*tchttp.BaseRequest

	// 待删除的部门ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAclsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAclsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchStatementBySidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据库操作列表

		StatementSet []*DatabaseStatement `json:"StatementSet,omitempty" name:"StatementSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchStatementBySidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchStatementBySidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceIdsRequest struct {
	*tchttp.BaseRequest

	// 地域码集合

	ApCodeSet []*string `json:"ApCodeSet,omitempty" name:"ApCodeSet"`
}

func (r *DescribeInstanceIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordSettingRequest struct {
	*tchttp.BaseRequest

	// 密码最小长度，8-20，默认8

	MinLength *uint64 `json:"MinLength,omitempty" name:"MinLength"`
	// 密码复杂度，0不限制，1包含字母和数字，2至少包括大写字母、小写字母、数字、特殊符号，默认2

	Complexity *uint64 `json:"Complexity,omitempty" name:"Complexity"`
	// 密码有效期，0不限制，30天，90天，180天

	ValidTerm *uint64 `json:"ValidTerm,omitempty" name:"ValidTerm"`
	// 检查最近n次密码设置是否存在相同密码，2-10，默认5

	CheckHistory *uint64 `json:"CheckHistory,omitempty" name:"CheckHistory"`
}

func (r *ModifyPasswordSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUsersDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUsersDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUsersDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SessionResult struct {

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 主机账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 会话大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 设备ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 设备名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 内部Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 外部Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 来源Ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 会话持续时长

	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
	// 该会话内命令数量 ，搜索图形会话时该字段无意义

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 该会话内高危命令数，搜索图形时该字段无意义

	DangerCount *uint64 `json:"DangerCount,omitempty" name:"DangerCount"`
	// 会话状态，如1会话活跃  2会话结束  3强制离线  4其他错误

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 会话Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 设备所属的地域

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 会话协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type AddDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 资产组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 需要添加到资产组的资产ID集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *AddDeviceGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDeviceGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOperationTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 需删除的成员用户ID集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *DeleteUserGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessEntryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccessEntryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessEntryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 资产组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 资产名或资产IP，模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// true - 查询已在该资产组的资产，false - 查询未在该资产组的资产

	Bound *bool `json:"Bound,omitempty" name:"Bound"`
	// 分页偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数，默认20, 最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 资产类型，1 - Linux，2 - Windows，3 - MySQL，4 - SQLServer

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 所属部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 	过滤条件，可按照标签键、标签进行过滤

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *DescribeDeviceGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicesPortResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDevicesPortResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {

	// 用户ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 用户名, 3-20个字符 必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户姓名， 最大20个字符，不能包含空白字符

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 手机号码， 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 电子邮件

	Email *string `json:"Email,omitempty" name:"Email"`
	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 用户失效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 所属用户组列表

	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`
	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问

	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
	// 用户所属部门（用于出参）

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 用户所属部门（用于入参）

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type CheckLDAPConnectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckLDAPConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLDAPConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicesDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDevicesDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组名，最大长度32字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户组所属部门的ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreateUserGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLDAPSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLDAPSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLDAPSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessWhiteListRuleRequest struct {
	*tchttp.BaseRequest

	// ip 10.10.10.1或者网段10.10.10.0/24，最小长度4字节，最大长度40字节。

	Source *string `json:"Source,omitempty" name:"Source"`
	// 备注信息，最小长度0字符，最大长度40字符。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateAccessWhiteListRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessWhiteListRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Departments struct {

	// 部门列表

	DepartmentSet []*Department `json:"DepartmentSet,omitempty" name:"DepartmentSet"`
	// 是否开启了部门管理 true - 已开启, false - 未开启

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 当前操作UIN是否是根部门管理员

	RootManager *bool `json:"RootManager,omitempty" name:"RootManager"`
}

type DescribeOperationTaskStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维任务统计信息

		OperationTaskStatistics *OperationTaskStatistics `json:"OperationTaskStatistics,omitempty" name:"OperationTaskStatistics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationTaskStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTaskStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportDevicesRequest struct {
	*tchttp.BaseRequest

	// 地域参数

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// cvm参数列表

	CvmSet []*Cvm `json:"CvmSet,omitempty" name:"CvmSet"`
	// 用户主机所在的VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 堡垒机服务ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *ImportDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessTimePolicyRequest struct {
	*tchttp.BaseRequest

	// 用户ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 生效时间段, 0、1组成的字符串，长度168(7*24), 代表该用户的生效时间. 0 - 未生效，1 - 生效

	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
}

func (r *ModifyAccessTimePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessTimePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest

	// ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *ResetDeviceAccountPrivateKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDeviceAccountPrivateKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandRequest struct {
	*tchttp.BaseRequest

	// 搜索区间的开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 搜索区间的结束时间，不填默认为开始时间到现在为止

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 资产的公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 资产的内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 执行的命令

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 根据拦截状态进行过滤：1 - 已执行，2 - 被阻断

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
	// 每页容量，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchCommandRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserBatchRequest struct {
	*tchttp.BaseRequest

	// 用户信息

	UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`
}

func (r *CreateUserBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作日志列表

		OperationEventSet []*OperationEvent `json:"OperationEventSet,omitempty" name:"OperationEventSet"`
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维任务列表

		OperationTasks []*OperationTask `json:"OperationTasks,omitempty" name:"OperationTasks"`
		// 任务总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 用户名或用户姓名，最长64个字符，模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// true - 查询已添加到该用户组的用户，false - 查询未添加到该用户组的用户

	Bound *bool `json:"Bound,omitempty" name:"Bound"`
	// 分页偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20, 最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 所属部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeUserGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命令列表

		Commands []*SearchCommandResult `json:"Commands,omitempty" name:"Commands"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchCommandResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GraphResult struct {

	// 日期和当日新开会话数量

	Sessions []*KeyCount `json:"Sessions,omitempty" name:"Sessions"`
	// 日期和当日命令总数

	Commands []*KeyCount `json:"Commands,omitempty" name:"Commands"`
	// 日期和高危命令总数

	Dangers []*KeyCount `json:"Dangers,omitempty" name:"Dangers"`
}

type TaskResult struct {

	// 运维任务结果日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 执行任务来源IP

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 运维任务所属用户

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 运维任务所属用户的姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 运维任务执行状态 1 - 执行中，2 - 成功，3 - 失败，4 - 部分失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 运维任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运维任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 堡垒机资源列表

		ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileBySidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 某会话的文件操作列表

		SearchFileBySidResult []*SearchFileBySidResult `json:"SearchFileBySidResult,omitempty" name:"SearchFileBySidResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFileBySidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileBySidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessWhiteListRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 访问白名单规则列表

		AccessWhiteListRuleSet []*AccessWhiteListRule `json:"AccessWhiteListRuleSet,omitempty" name:"AccessWhiteListRuleSet"`
		// 是否放开全部来源IP，如果为true，TotalCount为0，AccessWhiteListRuleSet为空

		AllowAny *bool `json:"AllowAny,omitempty" name:"AllowAny"`
		// 是否开启自动添加来源IP, 如果为true, 在开启访问白名单的情况下将自动添加来源IP

		AllowAuto *bool `json:"AllowAuto,omitempty" name:"AllowAuto"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessWhiteListRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessWhiteListRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// 用户名, 3-20个字符, 必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户姓名，最大长度20个字符，不能包含空白字符

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 大陆手机号直接填写，如果是其他国家、地区号码， 按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 电子邮件

	Email *string `json:"Email,omitempty" name:"Email"`
	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 所属用户组ID集合

	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`
	// 认证方式，0 - 本地， 1 - LDAP， 2 - OAuth 不传则默认为0

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问

	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
	// 所属部门ID，如：“1.2.3”

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindDeviceResourceRequest struct {
	*tchttp.BaseRequest

	// 资产ID集合

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 堡垒机服务ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *BindDeviceResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest

	// 待删除的用户ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户组列表

		GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeviceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogOutputSettings struct {

	// 是否已开启日志外发

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 是否已开启产品登录日志发送

	ProductLogin *bool `json:"ProductLogin,omitempty" name:"ProductLogin"`
	// 是否已开启资产登录日志发送

	DeviceLogin *bool `json:"DeviceLogin,omitempty" name:"DeviceLogin"`
}

type InquireCreateDasbResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品价格

		ProductCost *ProductCost `json:"ProductCost,omitempty" name:"ProductCost"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquireCreateDasbResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquireCreateDasbResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDasbResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDasbResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDasbResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessWhiteListRule struct {

	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// IP或者网段

	Source *string `json:"Source,omitempty" name:"Source"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type DasbCvmConfig struct {

	// 主机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主机核数

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 外网带宽

	NetBand *uint64 `json:"NetBand,omitempty" name:"NetBand"`
	// 系统盘大小

	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
	// 数据盘大小

	DataDiskSize *uint64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`
	// 购买月份

	MonthSpan *uint64 `json:"MonthSpan,omitempty" name:"MonthSpan"`
	// 所属商品模型ID

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名

	Region *string `json:"Region,omitempty" name:"Region"`
	// 使用的镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type SearchAuditLogRequest struct {
	*tchttp.BaseRequest

	// 开始时间，不得早于当前时间的180天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchAuditLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAuditLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SessionFile struct {

	// 会话Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 主机账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 设备ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 主机内网Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 主机外网Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 来源Ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 属于该会话的文件集合

	Files []*FileInformation `json:"Files,omitempty" name:"Files"`
}

type BindDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest

	// 主机账号ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机账号私钥，最新长度128字节，最大长度8192字节

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 主机账号私钥口令，最大长度256字节

	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
}

func (r *BindDeviceAccountPrivateKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceAccountPrivateKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAclsRequest struct {
	*tchttp.BaseRequest

	// 待删除的权限ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteAclsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAclsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmdTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmdTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmdTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTasksRequest struct {
	*tchttp.BaseRequest

	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运维任务类型，1 - 手工执行任务， 2 - 周期性任务

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 分页偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperationTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDepartmentRequest struct {
	*tchttp.BaseRequest

	// 部门ID，如：“1.2.3”

	Id *string `json:"Id,omitempty" name:"Id"`
	// 部门名称，1 - 256个字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 部门管理员的账号ID

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
}

func (r *ModifyDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationTask struct {

	// 运维任务主键ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建用户

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 运维人员姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 任务类型，1 - 手工执行任务， 2 - 周期性任务

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 周期性任务执行间隔，单位天

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 执行账户

	NextTime *string `json:"NextTime,omitempty" name:"NextTime"`
	// 下一次执行时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
}

type CreateUserGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功的用户组ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListAutoStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessWhiteListAutoStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListAutoStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExternalDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyExternalDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExternalDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindDeviceResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindDeviceResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationEventRequest struct {
	*tchttp.BaseRequest

	// 用户名，如果不包含其他条件时对user_name or real_name两个字段模糊查询

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名，模糊查询

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 查询时间范围，起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询时间范围，结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 来源IP，模糊查询

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 操作类型，参考DescribeOperationType返回结果

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 操作结果，1-成功，2-失败

	Result *uint64 `json:"Result,omitempty" name:"Result"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页每页记录数，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperationEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLDAPSettingRequest struct {
	*tchttp.BaseRequest

	// 是否开启LDAP认证，false-不开启，true-开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 服务器地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用服务器地址

	IpBackup *string `json:"IpBackup,omitempty" name:"IpBackup"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否开启SSL，false-不开启，true-开启

	EnableSSL *bool `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// Base DN

	BaseDN *string `json:"BaseDN,omitempty" name:"BaseDN"`
	// 管理员账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理员密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 用户属性

	AttributeUser *string `json:"AttributeUser,omitempty" name:"AttributeUser"`
	// 用户名属性

	AttributeUserName *string `json:"AttributeUserName,omitempty" name:"AttributeUserName"`
	// 自动同步，false-不开启，true-开启

	AutoSync *bool `json:"AutoSync,omitempty" name:"AutoSync"`
	// 覆盖用户信息，false-不开启，true-开启

	Overwrite *bool `json:"Overwrite,omitempty" name:"Overwrite"`
	// 同步周期，30～60000之间的整数

	SyncPeriod *uint64 `json:"SyncPeriod,omitempty" name:"SyncPeriod"`
	// 是否同步全部，false-不开启，true-开启

	SyncAll *bool `json:"SyncAll,omitempty" name:"SyncAll"`
	// 同步OU列表，SyncAll为false时必传

	SyncUnitSet []*string `json:"SyncUnitSet,omitempty" name:"SyncUnitSet"`
	// 组织单元属性

	AttributeUnit *string `json:"AttributeUnit,omitempty" name:"AttributeUnit"`
	// 用户姓名属性

	AttributeRealName *string `json:"AttributeRealName,omitempty" name:"AttributeRealName"`
	// 手机号属性

	AttributePhone *string `json:"AttributePhone,omitempty" name:"AttributePhone"`
	// 邮箱属性

	AttributeEmail *string `json:"AttributeEmail,omitempty" name:"AttributeEmail"`
}

func (r *ModifyLDAPSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLDAPSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 成员用户ID集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *AddUserGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportExternalDeviceRequest struct {
	*tchttp.BaseRequest

	// 资产参数列表

	DeviceSet []*ExternalDevice `json:"DeviceSet,omitempty" name:"DeviceSet"`
}

func (r *ImportExternalDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportExternalDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUsersDepartmentRequest struct {
	*tchttp.BaseRequest

	// 用户记录ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 用户所属部门的ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyUsersDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUsersDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命令和所属会话

		CommandSessionSet []*SessionCommand `json:"CommandSessionSet,omitempty" name:"CommandSessionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchCommandSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsWithDeviceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 携带拥有该同名账号主机数量的账号信息列表

		AccountWithDeviceCountSet []*AccountWithDeviceCount `json:"AccountWithDeviceCountSet,omitempty" name:"AccountWithDeviceCountSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsWithDeviceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsWithDeviceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecuritySetting struct {

	// 认证方式设置

	AuthMode *AuthModeSetting `json:"AuthMode,omitempty" name:"AuthMode"`
	// 密码安全设置

	Password *PasswordSetting `json:"Password,omitempty" name:"Password"`
	// 登录安全设置

	Login *LoginSetting `json:"Login,omitempty" name:"Login"`
	// LDAP配置信息

	LDAP *LDAPSetting `json:"LDAP,omitempty" name:"LDAP"`
}

type CreateCmdTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名，最大长度32字符，不能包含空白字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命令列表，\n分隔，最大长度32768字节

	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
}

func (r *CreateCmdTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmdTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSyncFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产同步标志

		AssetSyncFlags *AssetSyncFlags `json:"AssetSyncFlags,omitempty" name:"AssetSyncFlags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSyncFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAclRequest struct {
	*tchttp.BaseRequest

	// 权限名称，最大32字符，不能包含空白字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否开启磁盘映射

	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`
	// 是否开启剪贴板文件上行

	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`
	// 是否开启剪贴板文件下行

	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`
	// 是否开启剪贴板文本（含图片）上行

	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`
	// 是否开启剪贴板文本（含图片）下行

	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`
	// 是否开启 SFTP 文件上传

	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`
	// 文件传输上传大小限制（预留参数，目前暂未使用）

	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`
	// 是否开启 SFTP 文件下载

	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`
	// 文件传输下载大小限制（预留参数，目前暂未使用）

	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`
	// 是否允许任意账号登录

	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`
	// 关联的用户ID集合

	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`
	// 关联的用户组ID

	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`
	// 关联的资产ID集合

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 关联的资产组ID

	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`
	// 关联的账号

	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`
	// 关联的高危命令模板ID

	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitempty" name:"CmdTemplateIdSet"`
	// 是否开启rdp磁盘映射文件上传

	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`
	// 是否开启rdp磁盘映射文件下载

	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`
	// 是否开启rz sz文件上传

	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`
	// 是否开启rz sz文件下载

	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`
	// 是否开启 SFTP 文件删除

	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`
	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 访问权限所属部门的ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreateAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAclRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTaskResultRequest struct {
	*tchttp.BaseRequest

	// 搜索区间的开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 搜索区间的结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户名，长度不超过20

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名，长度不超过20

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 查询偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Map struct {

	// 对应map的key

	Name *string `json:"Name,omitempty" name:"Name"`
	// 对应map的value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type SearchCommandResult struct {

	// 命令输入的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 资产公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 资产内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 命令

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 命令执行情况，1--允许，2--拒绝

	Action *uint64 `json:"Action,omitempty" name:"Action"`
	// 命令所属的会话ID

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 命令执行时间相对于所属会话开始时间的偏移量，单位ms

	TimeOffset *uint64 `json:"TimeOffset,omitempty" name:"TimeOffset"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest

	// 如果IdSet不为空，则忽略其他参数

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页，偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20, 最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 精确查询，IdSet为空时才生效

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 精确查询，IdSet、UserName为空时才生效。
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 查询具有指定资产ID访问权限的用户

	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`
	// 认证方式，0 - 本地, 1 - LDAP, 2 - OAuth, 不传为全部

	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitempty" name:"AuthTypeSet"`
	// 部门ID，用于过滤属于某个部门的用户

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShowGraphRequest struct {
	*tchttp.BaseRequest

	// 开始时间，不得早于当前的180天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ShowGraphRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShowGraphRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseUsage struct {

	// 磐石中计费项id

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 用量

	UsageValue *uint64 `json:"UsageValue,omitempty" name:"UsageValue"`
}

type CreateDeviceAccountBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceAccountBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceAccountBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDasbResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyDasbResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyDasbResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetUserPasswordRequest struct {
	*tchttp.BaseRequest

	// 用户ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *ResetUserPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetUserPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveDevice struct {

	// 主机ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 地域信息

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 活跃主机的活跃次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type AddDeviceGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDeviceGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDeviceGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExternalDeviceRequest struct {
	*tchttp.BaseRequest

	// 资产记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 资产名，不能为空

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 资产IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 资产组ID集合

	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`
	// 资产所属部门ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyExternalDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExternalDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountWithDeviceCount struct {

	// 账号名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 拥有该同名账号的主机数量

	DeviceCount *uint64 `json:"DeviceCount,omitempty" name:"DeviceCount"`
}

type TopResult struct {

	// 在线用户数

	OnlineUser *int64 `json:"OnlineUser,omitempty" name:"OnlineUser"`
	// 高危用户和高危操作次数

	DangerUsers []*KeyCount `json:"DangerUsers,omitempty" name:"DangerUsers"`
	// 活跃用户和活跃次数

	ActiveUsers []*KeyCount `json:"ActiveUsers,omitempty" name:"ActiveUsers"`
	// 活跃主机和活跃次数及主机信息

	ActiveDevices []*ActiveDevice `json:"ActiveDevices,omitempty" name:"ActiveDevices"`
}

type SearchCommandSessionRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 默认值为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 检索的目标命令，为模糊搜索

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 开始时间，不得早于当前时间的180天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *SearchCommandSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DasbInstance struct {

	// 实例key

	PincKey *string `json:"PincKey,omitempty" name:"PincKey"`
	// 实例名称

	PsysName *string `json:"PsysName,omitempty" name:"PsysName"`
	// 公网ip

	MpublicIp *string `json:"MpublicIp,omitempty" name:"MpublicIp"`
	// 产品型号

	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`
	// 购买时间

	PaddTime *uint64 `json:"PaddTime,omitempty" name:"PaddTime"`
	// 内网IP

	MprivateIp *string `json:"MprivateIp,omitempty" name:"MprivateIp"`
	// 实例状态

	PsysStatus *uint64 `json:"PsysStatus,omitempty" name:"PsysStatus"`
	// 实例开始时间

	PstartTime *uint64 `json:"PstartTime,omitempty" name:"PstartTime"`
	// 实例结束时间

	PendTime *uint64 `json:"PendTime,omitempty" name:"PendTime"`
	// 实例所在地域ID

	Pregion *uint64 `json:"Pregion,omitempty" name:"Pregion"`
	// 实例所在地区ID

	PzoneId *uint64 `json:"PzoneId,omitempty" name:"PzoneId"`
	// CVM实例名列表

	CvmInstanceNames []*string `json:"CvmInstanceNames,omitempty" name:"CvmInstanceNames"`
	// CVM实例ID列表

	CvmInstanceIds []*string `json:"CvmInstanceIds,omitempty" name:"CvmInstanceIds"`
}

type SessionCommand struct {

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 账号

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 设备名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 内部Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 外部Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 命令列表

	Commands []*Command `json:"Commands,omitempty" name:"Commands"`
	// 记录数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 会话Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 设备id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 设备所属的地域

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
}

type DescribeAclsRequest struct {
	*tchttp.BaseRequest

	// 访问权限ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 访问权限名称，模糊查询，最长64字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20，最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 是否根据Name进行精确查询，默认值false

	Exact *bool `json:"Exact,omitempty" name:"Exact"`
	// 有访问权限的用户ID集合

	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`
	// 有访问权限的资产ID集合

	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`
	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 部门ID，用于过滤属于某个部门的访问权限

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeAclsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *ResetDeviceAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDeviceAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 文件操作列表

		Files *SearchFileResult `json:"Files,omitempty" name:"Files"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupsRequest struct {
	*tchttp.BaseRequest

	// 用户组ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 用户组名，模糊查询,长度：0-64字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，缺省20，最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 部门ID，用于过滤属于某个部门的用户组

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeUserGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckLDAPConnectionRequest struct {
	*tchttp.BaseRequest

	// 是否开启LDAP认证，必须为true

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 服务器地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用服务器地址

	IpBackup *string `json:"IpBackup,omitempty" name:"IpBackup"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否开启SSL，false-不开启，true-开启

	EnableSSL *bool `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// Base DN

	BaseDN *string `json:"BaseDN,omitempty" name:"BaseDN"`
	// 管理员账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理员密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
}

func (r *CheckLDAPConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLDAPConnectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetSyncStatus struct {

	// 上一次同步完成的时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 上一次同步的结果。 0 - 从未进行, 1 - 成功， 2 - 失败

	LastStatus *uint64 `json:"LastStatus,omitempty" name:"LastStatus"`
	// 同步任务是否正在进行中

	InProcess *bool `json:"InProcess,omitempty" name:"InProcess"`
}

type ImportDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLDAPUnitSetRequest struct {
	*tchttp.BaseRequest

	// 是否开启LDAP认证，true-开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 服务器地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用服务器地址

	IpBackup *string `json:"IpBackup,omitempty" name:"IpBackup"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否开启SSL，false-不开启，true-开启

	EnableSSL *bool `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// Base DN

	BaseDN *string `json:"BaseDN,omitempty" name:"BaseDN"`
	// 管理员账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理员密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 用户名映射属性

	AttributeUserName *string `json:"AttributeUserName,omitempty" name:"AttributeUserName"`
	// 部门过滤

	AttributeUnit *string `json:"AttributeUnit,omitempty" name:"AttributeUnit"`
}

func (r *DescribeLDAPUnitSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLDAPUnitSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceRequest struct {
	*tchttp.BaseRequest

	// 资产记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 资产所属组ID集合

	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`
	// 资产所属部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 高危事件

		EventSet []*EventResult `json:"EventSet,omitempty" name:"EventSet"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSessionRequest struct {
	*tchttp.BaseRequest

	// 内部Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 外部Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 用户名，长度不超过20

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 账号，长度不超过64

	Account *string `json:"Account,omitempty" name:"Account"`
	// 来源Ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 搜索区间的开始时间。若入参是Id，则非必传，否则为必传。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 搜索区间的结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 会话协议类型，只能是1、2、3或4 对应关系为1-tui 2-gui 3-file 4-数据库。若入参是Id，则非必传，否则为必传。

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 姓名，长度不超过20

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 主机名，长度不超过64

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 状态，1为活跃，2为结束，3为强制离线，4为其他错误

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *SearchSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PasswordSetting struct {

	// 密码最小长度，8-20，默认8。

	MinLength *uint64 `json:"MinLength,omitempty" name:"MinLength"`
	// 密码复杂度，0不限制，1包含字母和数字，2至少包括大写字母、小写字母、数字、特殊符号，默认2。

	Complexity *uint64 `json:"Complexity,omitempty" name:"Complexity"`
	// 密码有效期，0不限制，30天，90天，180天。

	ValidTerm *uint64 `json:"ValidTerm,omitempty" name:"ValidTerm"`
	// 检查最近n次密码设置是否存在相同密码，2-10，默认5。

	CheckHistory *uint64 `json:"CheckHistory,omitempty" name:"CheckHistory"`
}

type DescribeDeviceGroupsRequest struct {
	*tchttp.BaseRequest

	// 资产组ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 资产组名，最长64个字符，模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，缺省20，最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 部门ID，用于过滤属于某个部门的资产组

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeDeviceGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseUsageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// license用量列表

		LicenseUsageList []*LicenseUsage `json:"LicenseUsageList,omitempty" name:"LicenseUsageList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 会话监控所需信息

		ReplayInfo *ReplayInformation `json:"ReplayInfo,omitempty" name:"ReplayInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MonitorSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MonitorSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubtaskResult struct {

	// 执行日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 执行主机实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 执行主机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 执行主机地域

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 执行主机外网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 执行主机内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 运维任务状态 1 - 执行中，2 - 成功， 3 - 失败，4 - 超时

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 运维任务失败原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 运维任务命令退出码

	ExitCode *int64 `json:"ExitCode,omitempty" name:"ExitCode"`
	// 运维任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运维任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 运维任务执行结果输出

	StdOut *string `json:"StdOut,omitempty" name:"StdOut"`
	// 运维任务执行结果错误

	StdErr *string `json:"StdErr,omitempty" name:"StdErr"`
}

type DeleteDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAclRequest struct {
	*tchttp.BaseRequest

	// 访问权限名称，最大32字符，不能包含空白字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否开启磁盘映射

	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`
	// 是否开启剪贴板文件上行

	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`
	// 是否开启剪贴板文件下行

	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`
	// 是否开启剪贴板文本（含图片）上行

	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`
	// 是否开启剪贴板文本（含图片）下行

	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`
	// 是否开启文件传输上传

	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`
	// 文件传输上传大小限制（预留参数，目前暂未使用）

	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`
	// 是否开启文件传输下载

	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`
	// 文件传输下载大小限制（预留参数，目前暂未使用）

	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`
	// 是否允许任意账号登录

	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`
	// 关联的用户ID

	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`
	// 关联的用户组ID

	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`
	// 关联的资产ID

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 关联的资产组ID

	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`
	// 关联的账号

	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`
	// 关联的高危命令模板ID

	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitempty" name:"CmdTemplateIdSet"`
	// 访问权限ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否开启 RDP 磁盘映射文件上传

	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`
	// 是否开启 RDP 磁盘映射文件下载

	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`
	// 是否开启rz sz文件上传

	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`
	// 是否开启rz sz文件下载

	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`
	// 是否开启 SFTP 文件删除

	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`
	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 权限所属部门的ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseStatement struct {

	// 数据库操作

	Statement *string `json:"Statement,omitempty" name:"Statement"`
	// 操作时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 执行情况，1--允许，2--拒绝，3--确认

	Action *int64 `json:"Action,omitempty" name:"Action"`
}

type LoginSetting struct {

	// 登录会话超时，10分钟，20分钟，30分钟，默认20分钟

	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`
	// 连续密码错误次数，超过锁定账号，3-5

	LockThreshold *uint64 `json:"LockThreshold,omitempty" name:"LockThreshold"`
	// 账号锁定时长，10分钟，20分钟，30分钟

	LockTime *uint64 `json:"LockTime,omitempty" name:"LockTime"`
}

type CreateAssetSyncJobRequest struct {
	*tchttp.BaseRequest

	// 同步资产类别，1 - 主机资产, 2 - 数据库资产

	Category *uint64 `json:"Category,omitempty" name:"Category"`
}

func (r *CreateAssetSyncJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetSyncJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessWhiteListRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLDAPSyncFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLDAPSyncFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLDAPSyncFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessWhiteListRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功后返回的记录ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessWhiteListRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessWhiteListRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmdTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名，最长32字符，不能包含空白字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命令列表，\n分隔，最长32768字节

	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
	// 命令模板ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyCmdTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmdTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLDAPUnitSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ou 列表

		UnitSet []*string `json:"UnitSet,omitempty" name:"UnitSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLDAPUnitSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLDAPUnitSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceAccountsRequest struct {
	*tchttp.BaseRequest

	// 主机账号ID集合，非必需，如果使用IdSet则忽略其他过滤参数

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 主机账号名，模糊查询，不能单独出现，必须于DeviceId一起提交

	Account *string `json:"Account,omitempty" name:"Account"`
	// 主机ID，未使用IdSet时必须携带

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 分页偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDeviceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogOutputSettingsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogOutputSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogOutputSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicesDepartmentRequest struct {
	*tchttp.BaseRequest

	// 资产记录ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 资产所属部门的ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyDevicesDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetSyncJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetSyncJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDevicesRequest struct {
	*tchttp.BaseRequest

	// 待删除的ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatDasbResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatDasbResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatDasbResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileBySidResult struct {

	// 文件操作时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 1-上传文件 2-下载文件 3-删除文件 4-移动文件 5-重命名文件 6-新建文件夹 7-移动文件夹 8-重命名文件夹 9-删除文件夹

	Method *int64 `json:"Method,omitempty" name:"Method"`
	// 文件传输协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// method为上传、下载、删除时文件在服务器上的位置, 或重命名、移动文件前文件的位置

	FileCurr *string `json:"FileCurr,omitempty" name:"FileCurr"`
	// method为重命名、移动文件时代表移动后的新位置.其他情况为null

	FileNew *string `json:"FileNew,omitempty" name:"FileNew"`
	// method为上传文件、下载文件、删除文件时显示文件大小。其他情况为null

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// 堡垒机拦截情况, 1-已执行，  2-被阻断

	Action *int64 `json:"Action,omitempty" name:"Action"`
}

type BindDeviceAccountPrivateKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindDeviceAccountPrivateKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceAccountPrivateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportExternalDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportExternalDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportExternalDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmdTemplate struct {

	// 高危命令模板ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 高危命令模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命令列表，命令之间用换行符（"\n"）分隔

	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
}

type DasbCvmInstanceType struct {

	// 主机实例类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 主机实例名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 系统盘类型

	SystemDiskType *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`
	// 数据盘类型

	DataDiskType *string `json:"DataDiskType,omitempty" name:"DataDiskType"`
	// 所属区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type Group struct {

	// 组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属部门信息

	Department *Department `json:"Department,omitempty" name:"Department"`
}

type GroupWithCount struct {

	// 组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 成员数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type DescribeUserGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组成员总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户组成员列表

		UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Device struct {

	// 资产ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例ID，对应CVM、CDB等实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 地域编码

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 操作系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 资产类型 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 所属资产组列表

	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`
	// 资产绑定的账号数

	AccountCount *uint64 `json:"AccountCount,omitempty" name:"AccountCount"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 堡垒机服务信息，注意没有绑定服务时为null

	Resource *Resource `json:"Resource,omitempty" name:"Resource"`
	// 资产所属部门

	Department *Department `json:"Department,omitempty" name:"Department"`
}

type LDAPSetting struct {

	// 是否开启LDAP认证，false-不开启，true-开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 服务器地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用服务器地址

	IpBackup *string `json:"IpBackup,omitempty" name:"IpBackup"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否开启SSL，false-不开启，true-开启

	EnableSSL *bool `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// Base DN

	BaseDN *string `json:"BaseDN,omitempty" name:"BaseDN"`
	// 管理员账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 用户属性

	AttributeUser *string `json:"AttributeUser,omitempty" name:"AttributeUser"`
	// 用户名属性

	AttributeUserName *string `json:"AttributeUserName,omitempty" name:"AttributeUserName"`
	// 自动同步，false-不开启，true-开启

	AutoSync *bool `json:"AutoSync,omitempty" name:"AutoSync"`
	// 覆盖用户信息，false-不开启，true-开启

	Overwrite *bool `json:"Overwrite,omitempty" name:"Overwrite"`
	// 同步周期，30～60000之间的整数

	SyncPeriod *uint64 `json:"SyncPeriod,omitempty" name:"SyncPeriod"`
	// 是否同步全部，false-不开启，true-开启

	SyncAll *bool `json:"SyncAll,omitempty" name:"SyncAll"`
	// 同步OU列表

	SyncUnitSet []*string `json:"SyncUnitSet,omitempty" name:"SyncUnitSet"`
	// 组织单元属性

	AttributeUnit *string `json:"AttributeUnit,omitempty" name:"AttributeUnit"`
	// 用户姓名属性

	AttributeRealName *string `json:"AttributeRealName,omitempty" name:"AttributeRealName"`
	// 手机号属性

	AttributePhone *string `json:"AttributePhone,omitempty" name:"AttributePhone"`
	// 邮箱属性

	AttributeEmail *string `json:"AttributeEmail,omitempty" name:"AttributeEmail"`
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产信息列表

		DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthModeSettingRequest struct {
	*tchttp.BaseRequest

	// 双因子认证，0-不开启，1-OTP，2-短信

	AuthMode *uint64 `json:"AuthMode,omitempty" name:"AuthMode"`
}

func (r *ModifyAuthModeSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAuthModeSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthModeSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAuthModeSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAuthModeSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceAccountPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetDeviceAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDeviceAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 运维任务执行结果

		TaskResult []*TaskResult `json:"TaskResult,omitempty" name:"TaskResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchTaskResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmdTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功后返回的记录ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmdTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmdTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandBySidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命令列表

		CommandSet []*Command `json:"CommandSet,omitempty" name:"CommandSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchCommandBySidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandBySidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问权限总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 访问权限列表

		AclSet []*Acl `json:"AclSet,omitempty" name:"AclSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAclsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShowTopRequest struct {
	*tchttp.BaseRequest
}

func (r *ShowTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShowTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建部门的ID，如：“1.2.3”

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmdTemplatesRequest struct {
	*tchttp.BaseRequest

	// 待删除的ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteCmdTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmdTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalDevice struct {

	// 主机名，可为空

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统名称，只能是Linux、Windows或MySQL

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 资产所属的部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type CreateDeviceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功的资产组ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchEventRequest struct {
	*tchttp.BaseRequest

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 高危事件类型

	EventStyle *int64 `json:"EventStyle,omitempty" name:"EventStyle"`
	// 开始时间，若传该参数，则不得早于当前的180天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量，默认为5，最大200

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsWithDeviceCountRequest struct {
	*tchttp.BaseRequest

	// 主机ID集合

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 主机组ID集合

	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`
}

func (r *DescribeAccountsWithDeviceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsWithDeviceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileBySidRequest struct {
	*tchttp.BaseRequest

	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 可填写路径名或文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 分页用偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否创建审计日志,通过查看按钮调用时为true,其他为false

	AuditLog *bool `json:"AuditLog,omitempty" name:"AuditLog"`
	// 1-已执行，  2-被阻断

	AuditAction *int64 `json:"AuditAction,omitempty" name:"AuditAction"`
	// 以Protocol和Method为条件查询

	TypeFilters []*SearchFileTypeFilter `json:"TypeFilters,omitempty" name:"TypeFilters"`
}

func (r *SearchFileBySidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileBySidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceCountRequest struct {
	*tchttp.BaseRequest

	// 地域码

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 用户VPC实例ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 堡垒机服务ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资产类型,1-Linux, 2-Windows,3-MySQL,4-SqlServer 不传-全部

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 是否绑定服务,1-已绑定, 2-未绑定， 不传-全部

	BindResource *uint64 `json:"BindResource,omitempty" name:"BindResource"`
}

func (r *DescribeDeviceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaySessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回放所需信息

		ReplayInfo *ReplayInformation `json:"ReplayInfo,omitempty" name:"ReplayInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaySessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaySessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductCost struct {

	// 总费用

	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// RealCost

	RealCost *float64 `json:"RealCost,omitempty" name:"RealCost"`
}

type SearchFileResult struct {

	// 文件传输的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 资产公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 资产内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 操作结果：1 - 已执行，2 - 已阻断

	Action *uint64 `json:"Action,omitempty" name:"Action"`
	// 操作类型：1 - 文件上传，2 - 文件下载，3 - 文件删除，4 - 文件(夹)移动，5 - 文件(夹)重命名，6 - 新建文件夹，9 - 删除文件夹

	Method *uint64 `json:"Method,omitempty" name:"Method"`
	// 下载的文件（夹）路径及名称

	FileCurr *string `json:"FileCurr,omitempty" name:"FileCurr"`
	// 上传或新建文件（夹）路径及名称

	FileNew *string `json:"FileNew,omitempty" name:"FileNew"`
}

type CreateDeviceAccountBatchRequest struct {
	*tchttp.BaseRequest

	// 主机ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 主机账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 账号密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 私钥口令

	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
}

func (r *CreateDeviceAccountBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceAccountBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogOutputSettingsRequest struct {
	*tchttp.BaseRequest

	// 是否开启日志外发，false-不开启，true-开启

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 是否开启产品登录日志外发，false-不开启，true-开启

	ProductLogin *bool `json:"ProductLogin,omitempty" name:"ProductLogin"`
	// 是否开启资产登录日志外发，false-不开启，true-开启

	DeviceLogin *bool `json:"DeviceLogin,omitempty" name:"DeviceLogin"`
}

func (r *ModifyLogOutputSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogOutputSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceAccount struct {

	// 账号ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机ID

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 账号名

	Account *string `json:"Account,omitempty" name:"Account"`
	// true-已托管密码，false-未托管密码

	BoundPassword *bool `json:"BoundPassword,omitempty" name:"BoundPassword"`
	// true-已托管私钥，false-未托管私钥

	BoundPrivateKey *bool `json:"BoundPrivateKey,omitempty" name:"BoundPrivateKey"`
}

type OperationEvent struct {

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 操作时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 来源IP

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 操作类型

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 具体操作内容

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 操作结果，1-成功，2-失败

	Result *uint64 `json:"Result,omitempty" name:"Result"`
}

type ParamDeviceAccount struct {

	// 资产ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 账号名称，1-64，不能包含空白字符

	Account *string `json:"Account,omitempty" name:"Account"`
	// 账号密码，1-64字符

	Password *string `json:"Password,omitempty" name:"Password"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 账号密钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 账号密钥的密码

	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
}

type DescribeAssetSyncFlagRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetSyncFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceAccountsRequest struct {
	*tchttp.BaseRequest

	// 待删除的ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteDeviceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest

	// 资产ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 资产名或资产IP，模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// 暂未使用

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 地域码集合

	ApCodeSet []*string `json:"ApCodeSet,omitempty" name:"ApCodeSet"`
	// 操作系统类型, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 分页，偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 有该资产访问权限的用户ID集合

	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`
	// 过滤条件，资产绑定的堡垒机服务ID集合

	ResourceIdSet []*string `json:"ResourceIdSet,omitempty" name:"ResourceIdSet"`
	// 可提供按照多种类型过滤, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer

	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`
	// 过滤条件，可按照部门ID进行过滤

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 过滤条件，可按照标签键、标签进行过滤

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// InstanceId集合

		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Acl struct {

	// 访问权限ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 访问权限名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否开启磁盘映射

	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`
	// 是否开启剪贴板文件上行

	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`
	// 是否开启剪贴板文件下行

	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`
	// 是否开启剪贴板文本（目前含图片）上行

	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`
	// 是否开启剪贴板文本（目前含图片）下行

	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`
	// 是否开启文件传输上传

	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`
	// 文件传输上传大小限制（预留参数，暂未启用）

	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`
	// 是否开启文件传输下载

	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`
	// 文件传输下载大小限制（预留参数，暂未启用）

	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`
	// 是否允许任意账号登录

	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`
	// 关联的用户列表

	UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`
	// 关联的用户组列表

	UserGroupSet []*Group `json:"UserGroupSet,omitempty" name:"UserGroupSet"`
	// 关联的资产列表

	DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`
	// 关联的资产组列表

	DeviceGroupSet []*Group `json:"DeviceGroupSet,omitempty" name:"DeviceGroupSet"`
	// 关联的账号列表

	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`
	// 关联的高危命令模板列表

	CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitempty" name:"CmdTemplateSet"`
	// 是否开启 RDP 磁盘映射文件上传

	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`
	// 是否开启 RDP 磁盘映射文件下载

	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`
	// 是否开启 rz sz 文件上传

	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`
	// 是否开启 rz sz 文件下载

	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`
	// 是否开启 SFTP 文件删除

	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`
	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 所属部门的信息

	Department *Department `json:"Department,omitempty" name:"Department"`
}

type CreateDeviceAccountRequest struct {
	*tchttp.BaseRequest

	// 主机记录ID

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 账号名

	Account *string `json:"Account,omitempty" name:"Account"`
}

func (r *CreateDeviceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作类型名称与值的对应关系

		OperationTypeSet []*Map `json:"OperationTypeSet,omitempty" name:"OperationTypeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSubtaskResultByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 运维子任务执行结果

		SubtaskResult []*SubtaskResult `json:"SubtaskResult,omitempty" name:"SubtaskResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSubtaskResultByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSubtaskResultByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 要查询的运维任务ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 分页每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页起始位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 资产名称或资产IP

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资产地域过滤

	ApCode []*string `json:"ApCode,omitempty" name:"ApCode"`
}

func (r *DescribeOperationTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPasswordSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationTaskStatistics struct {

	// 当前允许配置的运维任务最大数量

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 当前已配置的运维任务数量

	Used *int64 `json:"Used,omitempty" name:"Used"`
	// 当前已配置的周期性触发运维任务数量

	Periodic *int64 `json:"Periodic,omitempty" name:"Periodic"`
	// 当前已配置的手工触发运维任务数量

	Manual *int64 `json:"Manual,omitempty" name:"Manual"`
	// 当前正在执行中的运维任务数量

	Running *int64 `json:"Running,omitempty" name:"Running"`
}

type DeleteDeviceGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessWhiteListStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest

	// 用户ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 用户姓名，最大长度20个字符，不能包含空格

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 电子邮件

	Email *string `json:"Email,omitempty" name:"Email"`
	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 所属用户组ID集合

	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`
	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth 不传则默认为0

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问

	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
	// 用户所属部门的ID，如1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSubtaskResultByIdRequest struct {
	*tchttp.BaseRequest

	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 查询偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 运维父任务执行日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 运维父任务执行状态

	Status []*uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SearchSubtaskResultByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSubtaskResultByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cvm struct {

	// cvm实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 操作系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 外网IP

	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 内网IP

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// vpc ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type BindDeviceAccountPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindDeviceAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 账号信息列表

		DeviceAccountSet []*DeviceAccount `json:"DeviceAccountSet,omitempty" name:"DeviceAccountSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListAutoStatusRequest struct {
	*tchttp.BaseRequest

	// true：放开自动添加IP；false：不放开自动添加IP

	AllowAuto *bool `json:"AllowAuto,omitempty" name:"AllowAuto"`
}

func (r *ModifyAccessWhiteListAutoStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListAutoStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTaskStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOperationTaskStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTaskStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetUserRequest struct {
	*tchttp.BaseRequest

	// 用户ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *ResetUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建用户的ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Department struct {

	// 部门ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 部门名称，1 - 256个字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 部门管理员账号ID

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
}

type ModifyAccessTimePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessTimePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessTimePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAclResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileTypeFilter struct {

	// 需要查询的文件传输类型，如SFTP/CLIP/RDP/RZSZ

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 在当前指定的protocol下进一步过滤具体操作类型,如剪贴板文件上传，剪贴板文件下载等

	Method []*int64 `json:"Method,omitempty" name:"Method"`
}

type DeleteUserGroupsRequest struct {
	*tchttp.BaseRequest

	// 待删除的用户组ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteUserGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DasbInstanceLog struct {

	// 日志ID

	LogId *uint64 `json:"LogId,omitempty" name:"LogId"`
	// 日志内容

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
	// 日志记录时间

	LogTime *uint64 `json:"LogTime,omitempty" name:"LogTime"`
}

type DescribeSecuritySettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecuritySettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecuritySettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogOutputSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志外发配置

		LogOutputSettings *LogOutputSettings `json:"LogOutputSettings,omitempty" name:"LogOutputSettings"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogOutputSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogOutputSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 资产组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 需要删除的资产ID集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *DeleteDeviceGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessEntryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维人员WEB访问入口

		AccessEntrySet []*string `json:"AccessEntrySet,omitempty" name:"AccessEntrySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessEntryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessEntryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepartmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 部门列表

		Departments *Departments `json:"Departments,omitempty" name:"Departments"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepartmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepartmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessWhiteListRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessWhiteListRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessWhiteListRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceGroupRequest struct {
	*tchttp.BaseRequest

	// 资产组名，最大长度32字符，不能为空

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资产组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 资产组所属部门ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyDeviceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchStatementBySidRequest struct {
	*tchttp.BaseRequest

	// 会话Id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// SQL语句

	Statement *string `json:"Statement,omitempty" name:"Statement"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据拦截状态进行过滤

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
}

func (r *SearchStatementBySidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchStatementBySidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShowGraphResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 折线图

		Graph *GraphResult `json:"Graph,omitempty" name:"Graph"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShowGraphResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShowGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDasbResourceRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例id bh-saas-xxxx

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DestroyDasbResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyDasbResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产组总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产组列表

		GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquireCreateDasbResourceRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 私有网络vpcId  vpc-xxxx

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网subnetId subnet-xxxx

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// lb vip运营商。默认为空，内网访问。 CMCC 移动 CTCC 电信 CUCC 联通 BGP 外网CAP

	LbVipIsp *string `json:"LbVipIsp,omitempty" name:"LbVipIsp"`
	// 资源类型。标准版(standard)或者专业版(pro)

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 资源默认节点数。取值：10/20/50/100/200/500/1000

	ResourceNode *uint64 `json:"ResourceNode,omitempty" name:"ResourceNode"`
	// 带宽扩展包

	MbpExp *uint64 `json:"MbpExp,omitempty" name:"MbpExp"`
	// 节点扩展包

	NodeExp *uint64 `json:"NodeExp,omitempty" name:"NodeExp"`
	// 计费周期 只支持 m

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 计费时长

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 付费方式。0表示按需计费/后付费，1表示预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *InquireCreateDasbResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquireCreateDasbResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginSettingRequest struct {
	*tchttp.BaseRequest

	// 登录会话超时，10分钟，20分钟，30分钟, 60分钟，默认30分钟

	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`
	// 连续密码错误次数，超过锁定账号，3-5

	LockThreshold *uint64 `json:"LockThreshold,omitempty" name:"LockThreshold"`
	// 账号锁定时长，10分钟，20分钟，30分钟

	LockTime *uint64 `json:"LockTime,omitempty" name:"LockTime"`
}

func (r *ModifyLoginSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmdTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmdTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmdTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepartmentsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDepartmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepartmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维任务详情

		OperationTaskDetail *OperationTaskDetail `json:"OperationTaskDetail,omitempty" name:"OperationTaskDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandBySidRequest struct {
	*tchttp.BaseRequest

	// 会话Id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 命令，可模糊搜索

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据拦截状态进行过滤

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
}

func (r *SearchCommandBySidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandBySidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployResourceRequest struct {
	*tchttp.BaseRequest

	// 需要开通服务的资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 需要开通服务的地域

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 子网所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 需要开通服务的VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 需要开通服务的子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 需要开通服务的子网网段

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 需要开通服务的VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 需要开通服务的VPC对应的网段

	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`
	// 需要开通服务的子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
}

func (r *DeployResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecuritySettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		SecuritySetting *SecuritySetting `json:"SecuritySetting,omitempty" name:"SecuritySetting"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecuritySettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecuritySettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatDasbResourceRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// vpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// subnetId

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// lb vip运营商。默认为空，内网访问。 CMCC 移动 CTCC 电信 CUCC 联通 BGP 外网CAP

	LbVipIsp *string `json:"LbVipIsp,omitempty" name:"LbVipIsp"`
	// 资源类型。标准版(standard)或者专业版(pro)

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 资源默认节点数。取值：10/20/50/100/200/500/1000

	ResourceNode *uint64 `json:"ResourceNode,omitempty" name:"ResourceNode"`
	// 带宽扩展包

	MbpExp *uint64 `json:"MbpExp,omitempty" name:"MbpExp"`
	// 节点扩展包

	NodeExp *uint64 `json:"NodeExp,omitempty" name:"NodeExp"`
	// 计费周期。只支持m

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 计费时长

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 付费方式。0表示按需计费/后付费，1表示预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *CreatDasbResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatDasbResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListRuleRequest struct {
	*tchttp.BaseRequest

	// 白名单规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// ip或网段信息，如10.10.10.1或10.10.10.0/24，最大长度40字节

	Source *string `json:"Source,omitempty" name:"Source"`
	// 备注信息，最大长度64字符。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAccessWhiteListRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorSessionRequest struct {
	*tchttp.BaseRequest

	// 会话Sid

	Sid *string `json:"Sid,omitempty" name:"Sid"`
}

func (r *MonitorSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MonitorSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportDeviceAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportDeviceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDeviceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessWhiteListRulesRequest struct {
	*tchttp.BaseRequest

	// 待删除的ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteAccessWhiteListRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessWhiteListRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
