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

package v20201230

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组规则

		Groups *SecurityGroup `json:"Groups,omitempty" name:"Groups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {

	// 节点总个数，要创建实例的某类型节点的总个数（节点组数与每组节点个数的乘积）

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 设置节点的复制类型，目前只有dn设置，可选参数为sync（同步），async（异步），sync2async（异步可退化），默认值为sync2async

	SyncType *string `json:"SyncType,omitempty" name:"SyncType"`
	// 单节点磁盘大小

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 节点总组数

	SetCount *int64 `json:"SetCount,omitempty" name:"SetCount"`
	// 节点规格代码

	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`
}

type ParamDetail struct {

	// 当前运行值。

	RunningValue *string `json:"RunningValue,omitempty" name:"RunningValue"`
	// 参数默认值。

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 修改此参数是否需要重启。

	NeedRestart *bool `json:"NeedRestart,omitempty" name:"NeedRestart"`
	// 参数取值范围。

	ValueRange *ParamConstraint `json:"ValueRange,omitempty" name:"ValueRange"`
	// 参数英文简介。

	ShortDesc *string `json:"ShortDesc,omitempty" name:"ShortDesc"`
	// 参数名。

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// 参数单位。

	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type SellSpec struct {

	// 节点类型: cn（计算节点）, dn（存储节点）。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 规格码，唯一标识一个规格，创建实例时传该字段用于标识一个规格。

	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`
	// 磁盘最小规格，单位：GB。

	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`
	// 磁盘调整步长，单位：GB。

	StorageStep *int64 `json:"StorageStep,omitempty" name:"StorageStep"`
	// 磁盘最大规格，单位：GB。

	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 内存大小，单位：GB。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 磁盘调整步长，单位：GB。

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// cpu核心数，单位：核。

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
}

type DescribeAvailableCharsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持的字符集列表。

		AvailableCharsets []*string `json:"AvailableCharsets,omitempty" name:"AvailableCharsets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableCharsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableCharsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 备份信息数组。

		BackupList []*Backup `json:"BackupList,omitempty" name:"BackupList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseIsolatedInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseIsolatedInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库账号名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为8~32位。

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBackupRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id。

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBackupRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBackupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 排序字段，支持 InstanceId, InstanceName, CreatedAt。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 实例状态信息，支持 creating, online, isolate。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 根据计费类型过滤，支持postpaid后付费，prepaid预付费两种。

	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，enum：InstanceName、Vip、InstanceId、TagKey。

	Filters []*ConditionFilter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，支持ASC和DESC，默认为ASC。

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 按一个或多个实例Id进行查询，每次最多请求100个。指定该参数后，Filters参数失效。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务详情列表

		TaskList []*TaskListOrDetail `json:"TaskList,omitempty" name:"TaskList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamItem struct {

	// 参数个数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 节点类型。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点名。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 节点的各个参数信息。

	Details []*ParamDetail `json:"Details,omitempty" name:"Details"`
}

type ConfigHistory struct {

	// 参数修改状态：success, failed, doing

	Status *string `json:"Status,omitempty" name:"Status"`
	// 节点类型。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 修改前的参数值。

	ParamOldValue *string `json:"ParamOldValue,omitempty" name:"ParamOldValue"`
	// 修改后的参数值。

	ParamNewValue *string `json:"ParamNewValue,omitempty" name:"ParamNewValue"`
	// 参数名。

	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
	// 参数修改的结束时间。

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 参数修改历史记录唯一id。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 参数修改的时间。

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 新实例名。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceUsage struct {

	// 总量

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 已使用量

	Used *int64 `json:"Used,omitempty" name:"Used"`
}

type DescribeAvailableEngineVersionsRequest struct {
	*tchttp.BaseRequest

	// 表示产品类型 (tbase,tdapg,tpostgres)

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeAvailableEngineVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableEngineVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeResourceOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineVersion struct {

	// 引擎版本。

	Version *string `json:"Version,omitempty" name:"Version"`
	// 引擎类型。

	Type *string `json:"Type,omitempty" name:"Type"`
}

type InstanceInfo struct {

	// 付费类型，postpaid - 后付费，prepaid - 预付费。

	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
	// 数据节点列表。

	DNNodes []*InstanceNode `json:"DNNodes,omitempty" name:"DNNodes"`
	// 所属大区。

	Area *string `json:"Area,omitempty" name:"Area"`
	// 数据库版本。

	Version *string `json:"Version,omitempty" name:"Version"`
	// 私有网络子网名称。

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 实例类型，TbaseV2或TbaseV3

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// GTM节点列表，在TbaseV2返回真实数据，在TbaseV3返回空数组

	GTMNodes []*InstanceNode `json:"GTMNodes,omitempty" name:"GTMNodes"`
	// 实例状态。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 产品名。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 私有网络Id。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例隔离者。

	IsolateOperator *string `json:"IsolateOperator,omitempty" name:"IsolateOperator"`
	// 私有网络名称。

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 任务类型。

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 到期时间。

	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
	// 地域名称。

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 预计销毁时间。

	DestroyTime *string `json:"DestroyTime,omitempty" name:"DestroyTime"`
	// 是否欠费，1 - 欠费， 0 - 不欠费。

	IsArrears *int64 `json:"IsArrears,omitempty" name:"IsArrears"`
	// 计算节点列表。

	CNNodes []*InstanceNode `json:"CNNodes,omitempty" name:"CNNodes"`
	// 地域。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 私有网络子网Id。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 管理员账号名。

	AdminUserName *string `json:"AdminUserName,omitempty" name:"AdminUserName"`
	// 实例状态描述。

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 集群id。

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟IP。

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 自动续费标记，0表示默认状态， 1表示自动续费，2表示明确不自动续费。在TbaseV2返回的数据真实有效，在TbaseV3下返回0

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// AppId。

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 任务描述。

	TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`
	// 虚拟端口。

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 可用区名称。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 项目ID，为空表示没有所属项目

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名，为空表示没有所属项目

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 如果指定NodeIds，则重启指定节点。如果为空或者包含所有节点id，则重启整个实例。不能和NodeTypes同时使用。

	NodeIds []*string `json:"NodeIds,omitempty" name:"NodeIds"`
	// 如果指定NodeTypes，则重启指定类型的节点。如果为空或者包含实例所有节点类型，则重启整个实例。不能和NodeIds同时使用。

	NodeTypes []*string `json:"NodeTypes,omitempty" name:"NodeTypes"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Inbound struct {

	// 策略，ACCEPT或者DROP。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 地址组id代表的地址集合。

	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`
	// 来源Ip或Ip段，例如192.168.0.0/16。

	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`
	// 描述。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 安全组id代表的地址集合。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 网络协议，支持udp、tcp等。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 端口。

	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`
	// 服务组id代表的协议和端口集合。

	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`
}

type DescribeAvailableEngineVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎版本数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用的引擎版本列表。

		EngineVersions []*EngineVersion `json:"EngineVersions,omitempty" name:"EngineVersions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableEngineVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableEngineVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionSellConf struct {

	// 是否为默认地域。

	IsDefaultRegion *bool `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`
	// 地域英文名。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域中文名。

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区售卖配置。

	AvailableZones []*ZoneSellConf `json:"AvailableZones,omitempty" name:"AvailableZones"`
	// 所属大区。

	Area *string `json:"Area,omitempty" name:"Area"`
	// 是否售罄，true是，false否

	IsSellOut *bool `json:"IsSellOut,omitempty" name:"IsSellOut"`
}

type TagInfo struct {

	// tag键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// tag值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type ZoneSellConf struct {

	// 是否为当前地域的默认可用区。

	IsDefaultZone *bool `json:"IsDefaultZone,omitempty" name:"IsDefaultZone"`
	// 可选节点规格。

	AvailableSpec []*SellSpec `json:"AvailableSpec,omitempty" name:"AvailableSpec"`
	// 可用区英文名。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区中文名。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 是否售罄，true是，false否

	IsSellOut *bool `json:"IsSellOut,omitempty" name:"IsSellOut"`
}

type IsolateInstanceHourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateInstanceHourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id。

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {

	// 策略，ACCEPT或者DROP。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 地址组id代表的地址集合。

	AddressMod *string `json:"AddressMod,omitempty" name:"AddressMod"`
	// 来源Ip或Ip段，例如192.168.0.0/16。

	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`
	// 描述。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 安全组id代表的地址集合。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 网络协议，支持udp、tcp等。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 端口。

	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`
	// 服务组id代表的协议和端口集合。

	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`
}

type DescribeAvailableZoneConfigRequest struct {
	*tchttp.BaseRequest

	// 引擎类型，支持TbaseV2，TbaseV3和TbaseV5

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
}

func (r *DescribeAvailableZoneConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZoneConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 需修改的节点参数配置。

	NodeConfigParams []*NodeConfigParams `json:"NodeConfigParams,omitempty" name:"NodeConfigParams"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncTypeInfo struct {

	// 节点的原复制类型，可选项(sync:同步不可退化，async:异步，sync2async:同步可退化)

	SourceSyncType *string `json:"SourceSyncType,omitempty" name:"SourceSyncType"`
	// 备节点的节点id。

	SlaveNodeId *string `json:"SlaveNodeId,omitempty" name:"SlaveNodeId"`
	// 节点的目标复制类型，可选项(sync:同步不可退化，async:异步，sync2async:同步可退化)

	TargetSyncType *string `json:"TargetSyncType,omitempty" name:"TargetSyncType"`
	// 主节点的节点id。

	MasterNodeId *string `json:"MasterNodeId,omitempty" name:"MasterNodeId"`
}

type DescribeErrorLogRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 查询错误日志的结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询错误日志的开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeErrorLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeErrorLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateInstanceHourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateInstanceHourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：tdapg 等。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 实例ID，格式如：tdpg-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceOverviewDetail struct {

	// 本产品实例数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品类型

	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
}

type ModifyInstanceSyncTypeRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 需要修改复制类型的实例节点信息（可选修改的复制类型包括：同步不可退化、同步可退化、异步）。

	ModifySyncTypeNodes []*SyncTypeInfo `json:"ModifySyncTypeNodes,omitempty" name:"ModifySyncTypeNodes"`
}

func (r *ModifyInstanceSyncTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceSyncTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountInfo struct {

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 账号名。

	Username *string `json:"Username,omitempty" name:"Username"`
}

type ErrorLogDetail struct {

	// 用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 错误详细信息。

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	// 错误日志发生时间。

	ErrorTime *string `json:"ErrorTime,omitempty" name:"ErrorTime"`
	// 数据库名。

	Database *string `json:"Database,omitempty" name:"Database"`
}

type DescribeBackupListsRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBackupListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailsRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest

	// 排序字段，支持CallTimes,CostTime,TotalCallTimesPercent,TotalCostTimePercent。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 数据库名。

	Database *string `json:"Database,omitempty" name:"Database"`
	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 排序方式，支持ASC和DESC。

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 查询慢查询日志的开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询慢查询日志的结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamConstraint struct {

	// type取值为section时的数值范围。

	Range *ConstraintRange `json:"Range,omitempty" name:"Range"`
	// type取值为enum时的枚举值。

	Enum []*string `json:"Enum,omitempty" name:"Enum"`
	// 用来描述取值范围类型。section: 数值范围, enum: 枚举, string: 字符串。

	Type *string `json:"Type,omitempty" name:"Type"`
	// type取值为string时的数值。

	String *string `json:"String,omitempty" name:"String"`
}

type CreateInstanceHourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 冻结流水Id。

		BillId *string `json:"BillId,omitempty" name:"BillId"`
		// 实例Id列表。

		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
		// 流程Id。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 订单号

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceHourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表。

		Items []*InstanceInfo `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ReleaseIsolatedInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseIsolatedInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowLogDetail struct {

	// 花费总时间。

	TotalTime *float64 `json:"TotalTime,omitempty" name:"TotalTime"`
	// 慢查询sql详细信息。

	NormalQuerys []*NormQueryItem `json:"NormalQuerys,omitempty" name:"NormalQuerys"`
	// 调用总次数。

	TotalCallTimes *int64 `json:"TotalCallTimes,omitempty" name:"TotalCallTimes"`
}

type TaskListOrDetail struct {

	// 流程状态， initial 已提交,未执行，running  正在运行中，success 执行成功，failed 执行失败，undoed 执行回滚成功，undoing 执行回滚中

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 流程创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 流程类型，create 创建 ,modify 变配, isolate 隔离, activeIsolate 恢复,destroy 销毁
	// expand 扩容, offline 下线, init 初始化, restart 重启, retreat 回档, backup 备份,modifyParams 修改参数,
	// modifyInstanceNodeSyncType 修改参数, setBackupRules 设置备份规则

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 任务进度

	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
	// 流程结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 任务ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 任务错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type DescribeBackupRulesRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 付费类型，postpaid - 后付费，prepaid - 预付费。

		TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
		// 备份及日志磁盘空间。

		BackupStorage *int64 `json:"BackupStorage,omitempty" name:"BackupStorage"`
		// 数据节点列表。

		DNNodes []*InstanceNode `json:"DNNodes,omitempty" name:"DNNodes"`
		// 所属大区。

		Area *string `json:"Area,omitempty" name:"Area"`
		// 字符集。

		Charset *string `json:"Charset,omitempty" name:"Charset"`
		// 私有网络子网名称。

		SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
		// 实例版本，TbaseV2或TbaseV3

		EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
		// GTM节点列表。

		GTMNodes []*InstanceNode `json:"GTMNodes,omitempty" name:"GTMNodes"`
		// 状态。

		Status *string `json:"Status,omitempty" name:"Status"`
		// 私有网络Id。

		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
		// 转发节点列表

		FNNodes []*InstanceNode `json:"FNNodes,omitempty" name:"FNNodes"`
		// 可用区。

		Zone *string `json:"Zone,omitempty" name:"Zone"`
		// 实例Id。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 私有网络名称。

		VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
		// 任务类型。

		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
		// 私有网络子网Id。

		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
		// 地域名称。

		RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
		// 计算节点列表。

		CNNodes []*InstanceNode `json:"CNNodes,omitempty" name:"CNNodes"`
		// 地域。

		Region *string `json:"Region,omitempty" name:"Region"`
		// 到期时间。

		PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
		// 引擎版本。

		EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
		// 管理员账号名。

		AdminUserName *string `json:"AdminUserName,omitempty" name:"AdminUserName"`
		// 状态描述。

		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
		// 实例名称。

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 集群名称。

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 虚拟IP。

		Vip *string `json:"Vip,omitempty" name:"Vip"`
		// 任务描述。

		TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`
		// 虚拟端口。

		Vport *int64 `json:"Vport,omitempty" name:"Vport"`
		// 创建时间。

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 可用区名称。

		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
		// 项目ID，为空表示没有所属项目

		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
		// 项目名，为空表示没有所属项目

		ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceNode struct {

	// 节点类型。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点名。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 节点已使用存储空间。

	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`
	// 节点内存。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 节点磁盘。

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 节点Id。

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 节点规格。

	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`
	// 节点的复制类型，sync:同步不可退化，async:异步，sync2async:同步可退化。

	SyncType *string `json:"SyncType,omitempty" name:"SyncType"`
	// 该节点在其节点组中的节点类型，例如master表示其为主节点。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点cpu。

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
}

type DescribeBackupDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的记录数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 节点详情。

		BackupDetailList []*NodeBackupDetail `json:"BackupDetailList,omitempty" name:"BackupDetailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：tbase等。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 实例ID，格式如：tbase-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceOverview struct {

	// 所有类型实例总数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 实例资源详情

	Detail []*InstanceOverviewDetail `json:"Detail,omitempty" name:"Detail"`
}

type DescribeDbNameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的db名称数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的db名称列表。

		DbNameList []*string `json:"DbNameList,omitempty" name:"DbNameList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDbNameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 实例节点类型: cn, dn

	NodeTypes []*string `json:"NodeTypes,omitempty" name:"NodeTypes"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductLicenseData struct {

	// 磐石分配的产品 ID，固定为 87870

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 产品描述

	ProductIdDescribe *string `json:"ProductIdDescribe,omitempty" name:"ProductIdDescribe"`
	// 数量单位或节点类型

	UsageUnit *string `json:"UsageUnit,omitempty" name:"UsageUnit"`
	// 机器上架时标记为 ALL 或 DN 的节点数量

	UsageValue *string `json:"UsageValue,omitempty" name:"UsageValue"`
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的节点数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例各节点参数信息。

		Items []*ParamItem `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态。

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBConfigHistoryRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBConfigHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBConfigHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 引擎版本

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 管理员账号的密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 字符集，根据DescribeAvailableCharset查询

	Charset *string `json:"Charset,omitempty" name:"Charset"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 网络类型

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 创建实例cn节点相关的规格信息。

	CnNode *NodeInfo `json:"CnNode,omitempty" name:"CnNode"`
	// 创建实例dn节点相关的规格信息。

	DnNode *NodeInfo `json:"DnNode,omitempty" name:"DnNode"`
	// 购买数量，(0,10]，不传默认创建1个

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 要绑定的安全组

	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`
	// 回档实例必传参数

	Recovery *RecoveryInfo `json:"Recovery,omitempty" name:"Recovery"`
	// 标签

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
	// 引擎类型，当前只支持TbaseV2,TbaseV3

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateInstanceHourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceHourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDbNameListRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// db名称，可进行模糊匹配。

	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *DescribeDbNameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbNameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeConfigParams struct {

	// 需要修改的参数配置。

	ConfigParams []*ConfigParam `json:"ConfigParams,omitempty" name:"ConfigParams"`
	// 节点类型，目前支持：cn、dn。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

type NormQueryItem struct {

	// 用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 话费的最长时间。

	MaxCostTime *float64 `json:"MaxCostTime,omitempty" name:"MaxCostTime"`
	// 花费的最小时间。

	MinCostTime *float64 `json:"MinCostTime,omitempty" name:"MinCostTime"`
	// 花费总时间。

	CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`
	// 在总的慢查询中的访问时间占比。

	TotalCostTimePercent *float64 `json:"TotalCostTimePercent,omitempty" name:"TotalCostTimePercent"`
	// 调用次数。

	CallTimes *int64 `json:"CallTimes,omitempty" name:"CallTimes"`
	// 写io总耗时。

	WriteCostTime *float64 `json:"WriteCostTime,omitempty" name:"WriteCostTime"`
	// 读io总耗时。

	ReadCostTime *float64 `json:"ReadCostTime,omitempty" name:"ReadCostTime"`
	// 脱敏后的sql语句。

	NormalQuery *string `json:"NormalQuery,omitempty" name:"NormalQuery"`
	// 最早一条慢查询的sql时间。

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 在总的慢查询中的访问次数占比。

	TotalCallTimesPercent *float64 `json:"TotalCallTimesPercent,omitempty" name:"TotalCallTimesPercent"`
	// 数据库名。

	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	// 执行时间最长的语句。

	MaxElapsedQuery *string `json:"MaxElapsedQuery,omitempty" name:"MaxElapsedQuery"`
	// 客户端ip。

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 最近一条慢查询的sql时间。

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 读共享内存块数。

	SharedReadBlocks *int64 `json:"SharedReadBlocks,omitempty" name:"SharedReadBlocks"`
	// 写共享内存块数。

	SharedWriteBlocks *int64 `json:"SharedWriteBlocks,omitempty" name:"SharedWriteBlocks"`
}

type DescribeInstanceTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 异步任务id，由异步的任务接口返回。

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeInstanceTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的账号数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 账号信息。

		Accounts []*AccountInfo `json:"Accounts,omitempty" name:"Accounts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConstraintRange struct {

	// 区间的最大值。

	Max *string `json:"Max,omitempty" name:"Max"`
	// 区间的最小值。

	Min *string `json:"Min,omitempty" name:"Min"`
}

type DescribeDBConfigHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的参数操作记录数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置变更的操作记录。

		ConfigHistory []*ConfigHistory `json:"ConfigHistory,omitempty" name:"ConfigHistory"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBConfigHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBConfigHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceSyncTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id。

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceSyncTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceSyncTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 入站规则。

	Inbound *Inbound `json:"Inbound,omitempty" name:"Inbound"`
	// 出站规则。

	Outbound *Outbound `json:"Outbound,omitempty" name:"Outbound"`
	// 项目ID。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 安全组ID。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组名称。

	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	// 安全组备注。

	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

type DescribeBackupDetailsRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 唯一标识本次备份。

	BackupUid *int64 `json:"BackupUid,omitempty" name:"BackupUid"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBackupDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Backup struct {

	// 备份状态。1(备份中)，2（备份成功），-1（备份失败）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例备份信息的唯一标识，查询备份详情时的入参。

	BackupUid *int64 `json:"BackupUid,omitempty" name:"BackupUid"`
	// 备份类型。

	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
	// 备份策略类型。

	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`
	// 实例备份开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 实例备份结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type RecoveryInfo struct {

	// 要回档的实例ID

	RecoveryInstanceId *string `json:"RecoveryInstanceId,omitempty" name:"RecoveryInstanceId"`
	// 回档时间，会在[获取回档时间]接口中返回，需选择回档时间范围中的某一个时刻点

	RecoveryTime *string `json:"RecoveryTime,omitempty" name:"RecoveryTime"`
	// 回档任务ID，会在[获取回档时间]接口中返回

	RecoveryTaskId *int64 `json:"RecoveryTaskId,omitempty" name:"RecoveryTaskId"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableCharsetRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableCharsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableCharsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的错误日志数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 错误日志详细信息。

		ErrorLogDetails []*ErrorLogDetail `json:"ErrorLogDetails,omitempty" name:"ErrorLogDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeErrorLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeErrorLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的慢查询日志数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 慢查询日志的详细信息。

		SlowLogDetails *SlowLogDetail `json:"SlowLogDetails,omitempty" name:"SlowLogDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTasksRequest struct {
	*tchttp.BaseRequest

	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，目前支持StartTime、EndTime、InstanceIds、TaskType、Status、TaskIds

	Filters []*ConditionFilter `json:"Filters,omitempty" name:"Filters"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUserTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dpag-751k9fi1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RenewInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionFilter struct {

	// 过滤字段值。

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 需要过滤的字段，enum：InstanceName，Vip，InstanceId。

	Name *string `json:"Name,omitempty" name:"Name"`
}

type SetBackupRulesRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 最小备份开始时间，形如：02:00:00。取值00:00:00~24:00:00。

	MinBackupStartTime *string `json:"MinBackupStartTime,omitempty" name:"MinBackupStartTime"`
	// 最大备份开始时间，形如：02:00:00。取值00:00:00~24:00:00。

	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitempty" name:"MaxBackupStartTime"`
	// 备份频率，每天备份几次，enum：1。

	BackupFrequency *int64 `json:"BackupFrequency,omitempty" name:"BackupFrequency"`
}

func (r *SetBackupRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBackupRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZoneConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有地域可选配置数目。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可售卖地域配置详情。

		Items []*RegionSellConf `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableZoneConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份保存时间。

		BackupReserveDays *int64 `json:"BackupReserveDays,omitempty" name:"BackupReserveDays"`
		// 最早备份开始时间。

		MinBackupStartTime *string `json:"MinBackupStartTime,omitempty" name:"MinBackupStartTime"`
		// 是否备份xlog。

		IsBackupXlog *int64 `json:"IsBackupXlog,omitempty" name:"IsBackupXlog"`
		// 最迟备份开始时间。

		MaxBackupStartTime *string `json:"MaxBackupStartTime,omitempty" name:"MaxBackupStartTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域英文 ID

		Region *string `json:"Region,omitempty" name:"Region"`
		// region id

		RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
		// 地域名称

		RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
		// 物理服务器数量

		ServerCount *int64 `json:"ServerCount,omitempty" name:"ServerCount"`
		// 实例统计数据

		Instances *InstanceOverview `json:"Instances,omitempty" name:"Instances"`
		// CPU 使用详情

		Cpu *ResourceUsage `json:"Cpu,omitempty" name:"Cpu"`
		// 内存使用详情

		Memory *ResourceUsage `json:"Memory,omitempty" name:"Memory"`
		// 存储使用详情

		Storage *ResourceUsage `json:"Storage,omitempty" name:"Storage"`
		// 产品名称

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// 接口版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// 服务器上架时标记为 ALL 或 DN 节点的数量统计

		Data []*ProductLicenseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigParam struct {

	// 参数名。

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// 参数值。

	ParameterValue *string `json:"ParameterValue,omitempty" name:"ParameterValue"`
}

type NodeBackupDetail struct {

	// 备份状态。1(备份中)，2(备份成功)，-1(备份失败)

	LatestStatus *int64 `json:"LatestStatus,omitempty" name:"LatestStatus"`
	// 节点类型。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点名。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 最近备份结束时间。

	LatestEndTime *string `json:"LatestEndTime,omitempty" name:"LatestEndTime"`
	// 最近备份开始时间。

	LatestStartTime *string `json:"LatestStartTime,omitempty" name:"LatestStartTime"`
}
