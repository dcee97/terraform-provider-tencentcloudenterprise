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

package v20201207

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type InstanceTagInfo struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type KongServicePreview struct {

	// 服务名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 后端配置

	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitempty" name:"UpstreamInfo"`
	// 后端类型

	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type BindCLBSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 白名单列表

	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
	// portal

	CLBType *string `json:"CLBType,omitempty" name:"CLBType"`
	// 环境名称

	Env *string `json:"Env,omitempty" name:"Env"`
	// 部署地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *BindCLBSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindCLBSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GovernanceNamespaceInput struct {

	// 命名空间名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述信息。

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 新增的可以操作此命名空间的用户ID列表

	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
	// 新增的可以操作此命名空间的用户组ID列表

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// 移除可以操作此命名空间的用户ID列表

	RemoveUserIds []*string `json:"RemoveUserIds,omitempty" name:"RemoveUserIds"`
	// 移除可以操作此命名空间的用户组ID列表

	RemoveGroupIds []*string `json:"RemoveGroupIds,omitempty" name:"RemoveGroupIds"`
}

type DeleteGovernanceNamespacesRequest struct {
	*tchttp.BaseRequest

	// tse 实例 id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间信息。

	GovernanceNamespaces []*GovernanceNamespaceInput `json:"GovernanceNamespaces,omitempty" name:"GovernanceNamespaces"`
}

func (r *DeleteGovernanceNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表

		List []*OperateInstance `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnablePolarisCLSResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnablePolarisCLSResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnablePolarisCLSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSREInstanceSubnetRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 子网ID

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

func (r *UpdateSREInstanceSubnetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSREInstanceSubnetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGovernanceStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGovernanceStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewaySourceRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务来源ID

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 服务来源类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteCloudNativeAPIGatewaySourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewaySourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NacosConfigInfo struct {

	// 参数名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 当前运行参数值

	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`
	// 配置的类型说明,支持option、range

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 配置类型的限制描述

	ConfigTypeLimit *string `json:"ConfigTypeLimit,omitempty" name:"ConfigTypeLimit"`
	// 默认参数值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 参数描述说明

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// true表示只读，不支持修改

	ReadOnly *string `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 对应的分组

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 依赖的配置项，需要同时满足

	DependencyConfigs *string `json:"DependencyConfigs,omitempty" name:"DependencyConfigs"`
	// 分组中文显示名称

	GroupNameDesc *string `json:"GroupNameDesc,omitempty" name:"GroupNameDesc"`
}

type NginxIngressVersion struct {

	// 版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// 兼容的 k8s 版本列表

	CompatibleK8sVersions []*string `json:"CompatibleK8sVersions,omitempty" name:"CompatibleK8sVersions"`
}

type DescribeCloudNativeAPIGatewayIngressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Ingress列表

		IngressList []*Ingress `json:"IngressList,omitempty" name:"IngressList"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayIngressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayIngressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperMigrateMergeStatusRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeZookeeperMigrateMergeStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperMigrateMergeStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartZookeeperMigrateImportDataRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StartZookeeperMigrateImportDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartZookeeperMigrateImportDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApolloEnvInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApolloEnvInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApolloEnvInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadCloudNativeAPIGatewayPluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 插件信息

		Result *CloudNativeAPIGatewayPluginInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadCloudNativeAPIGatewayPluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadCloudNativeAPIGatewayPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigFileInput struct {

	// 配置文件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 配置文件命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 配置文件组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 配置文件内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 配置文件格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 配置文件注释

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 配置文件标签数组

	Tags []*ConfigFileTag `json:"Tags,omitempty" name:"Tags"`
}

type PolarisServerInterface struct {

	// 接口名

	Interface *string `json:"Interface,omitempty" name:"Interface"`
}

type DescribeGovernanceInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所在的服务名。

	Service *string `json:"Service,omitempty" name:"Service"`
	// 实例所在命名空间名。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 根据实例ip过滤，多个ip使用英文逗号分隔。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 根据实例版本过滤。

	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`
	// 根据实例协议过滤。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 根据实例健康状态过滤。false：表示不健康，true：表示健康。

	HealthStatus *bool `json:"HealthStatus,omitempty" name:"HealthStatus"`
	// 根据实例隔离状态过滤。false：表示非隔离，true：表示隔离中。

	Isolate *bool `json:"Isolate,omitempty" name:"Isolate"`
	// 根据元数据信息过滤。目前只支持一组元数据键值，若传了多个键值对，只会以第一个过滤。

	Metadatas []*Metadata `json:"Metadatas,omitempty" name:"Metadatas"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGovernanceInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateAgentDetailRequest struct {
	*tchttp.BaseRequest

	// 镜像地址

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// China

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateAgentDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateAgentDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnablePolarisCLSRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

func (r *EnablePolarisCLSRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnablePolarisCLSRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargeInfo struct {

	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 计费方式

	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
	// 数据存储方式

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 每小时每G费用

	Charge *float64 `json:"Charge,omitempty" name:"Charge"`
}

type ModifyGovernanceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTagInfosRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 修改后的实例标签集合

	TagInfos []*InstanceTagInfo `json:"TagInfos,omitempty" name:"TagInfos"`
}

func (r *ModifyInstanceTagInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTagInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartZookeeperMigrateImportDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartZookeeperMigrateImportDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartZookeeperMigrateImportDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateEksRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	EksClusterId *string `json:"EksClusterId,omitempty" name:"EksClusterId"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DeleteOperateEksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateEksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceStrategyRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 策略修改请求数组

	Strategy []*ModifyAuthStrategy `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *ModifyGovernanceStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Ingress的json字符串

		IngressData *string `json:"IngressData,omitempty" name:"IngressData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallCloudNativeAPIGatewaySystemPluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallCloudNativeAPIGatewaySystemPluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallCloudNativeAPIGatewaySystemPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Metadata struct {

	// 元数据键名。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 元数据键值。不填则默认为空字符串。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeZookeeperService struct {

	// 注册服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 所有实例个数

	TotalServiceInstanceCount *int64 `json:"TotalServiceInstanceCount,omitempty" name:"TotalServiceInstanceCount"`
}

type NodeAccessControl struct {

	// 端口控制信息

	PortInfos []*PortInfos `json:"PortInfos,omitempty" name:"PortInfos"`
}

type TagInstanceDescription struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 引擎实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type ZookeeperRootDescription struct {

	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 创建的事务ID

	CreateZxid *string `json:"CreateZxid,omitempty" name:"CreateZxid"`
	// znode最近一次修改的事物ID

	ModifiedZxid *string `json:"ModifiedZxid,omitempty" name:"ModifiedZxid"`
	// 创建时间时间戳，单位毫秒

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间时间戳，单位毫秒

	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// znode的数据域变化次数

	ZNodeVersion *int64 `json:"ZNodeVersion,omitempty" name:"ZNodeVersion"`
	// znode的节点的变化次数

	ChildVersion *int64 `json:"ChildVersion,omitempty" name:"ChildVersion"`
	// znode的acl变化次数

	ACLVersion *int64 `json:"ACLVersion,omitempty" name:"ACLVersion"`
	// 临时节点的会话id，如果不是临时节点，则为0

	EphemeralOwner *int64 `json:"EphemeralOwner,omitempty" name:"EphemeralOwner"`
	// 数据长度

	DataLength *int64 `json:"DataLength,omitempty" name:"DataLength"`
	// 子节点个数

	NumChildren *int64 `json:"NumChildren,omitempty" name:"NumChildren"`
	// 最近修改的子节点

	PZxid *string `json:"PZxid,omitempty" name:"PZxid"`
	// 子节点描述

	ChildDesc []*ZookeeperNodeDescription `json:"ChildDesc,omitempty" name:"ChildDesc"`
	// 节点名称

	Key *string `json:"Key,omitempty" name:"Key"`
}

type DescribeEnginePodAccessAddressRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 地域信息

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *DescribeEnginePodAccessAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnginePodAccessAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayNode struct {

	// 云原生网关节点 id

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 节点 ip

	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`
}

type DescribeGovernanceNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 治理中心命名空间实例列表。

		Content []*GovernanceNamespace `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateInstanceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例详情

		Data *OperateInstance `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApolloDBSvTag struct {

	// 规格ID

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// DB计费标签

	SvTag *string `json:"SvTag,omitempty" name:"SvTag"`
}

type DescribeGovernanceGroupDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组详细

		Group *UserGroup `json:"Group,omitempty" name:"Group"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceGroupDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceGroupDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroySREInstanceRequest struct {
	*tchttp.BaseRequest

	// 微服务注册引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroySREInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroySREInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageZookeeperConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对配置中心操作配置之后的返回值

		Result *string `json:"Result,omitempty" name:"Result"`
		// 操作是否成功

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageZookeeperConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageZookeeperConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigFileReleaseHistory struct {

	// 配置文件发布历史记录id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 配置文件发布历史名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 配置文件发布历史命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 配置文件发布历史组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 配置文件发布历史名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 配置文件发布历史内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 配置文件发布历史格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 配置文件发布历史注释

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 配置文件发布历史Md5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 配置文件发布历史类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 配置文件发布历史状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 配置文件发布历史标签组

	Tags []*ConfigFileTag `json:"Tags,omitempty" name:"Tags"`
	// 配置文件发布创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 配置文件发布创建者

	CreateBy *string `json:"CreateBy,omitempty" name:"CreateBy"`
	// 配置文件发布修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 配置文件发布修改者

	ModifyBy *string `json:"ModifyBy,omitempty" name:"ModifyBy"`
}

type ModifyNativeGatewayServiceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNativeGatewayServiceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNativeGatewayServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BreakZookeeperMigrateClusterRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *BreakZookeeperMigrateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BreakZookeeperMigrateClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSREInstanceFeatureResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前实例是否有请求的特性

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 功能特性提示

		Tips *string `json:"Tips,omitempty" name:"Tips"`
		// 当前引擎是否支持升级

		UpgradeEnable *bool `json:"UpgradeEnable,omitempty" name:"UpgradeEnable"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSREInstanceFeatureResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSREInstanceFeatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAPIGatewayCanaryRuleList struct {

	// 灰度规则

	CanaryRuleList []*CloudNativeAPIGatewayCanaryRule `json:"CanaryRuleList,omitempty" name:"CanaryRuleList"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type GovernanceAlias struct {

	// 服务别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 服务别名命名空间

	AliasNamespace *string `json:"AliasNamespace,omitempty" name:"AliasNamespace"`
	// 服务别名指向的服务名

	Service *string `json:"Service,omitempty" name:"Service"`
	// 服务别名指向的服务命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务别名的描述信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 服务别名创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 服务别名修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 服务别名ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 该服务别名是否可以编辑

	Editable *bool `json:"Editable,omitempty" name:"Editable"`
}

type Principals struct {

	// 用户ID列表

	Users []*PrincipalEntry `json:"Users,omitempty" name:"Users"`
	// 用户组ID列表

	Groups []*PrincipalEntry `json:"Groups,omitempty" name:"Groups"`
}

type DescribeNativeGatewayServiceSourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总实例数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 服务来源实例列表

		List []*NativeGatewayServiceSourceItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNativeGatewayServiceSourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNativeGatewayServiceSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateEngineSpecListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎规格列表

		List []*OperateEngineSpec `json:"List,omitempty" name:"List"`
		// 列表总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateEngineSpecListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateEngineSpecListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayCertificateDetailsRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 证书Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *ListCloudNativeAPIGatewayCertificateDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayCertificateDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewaySystemParameterReq struct {

	// 参数名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeCloudNativeAPIGatewayBasicLogResult struct {

	// 获取基础日志的返回内容

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
}

type KongIngressControllerK8sVersion struct {

	// Kong Ingress Controller的版本

	KongIngressControllerVersion *string `json:"KongIngressControllerVersion,omitempty" name:"KongIngressControllerVersion"`
	// Kong Ingress Controller对应的K8s版本信息

	K8sVersions []*string `json:"K8sVersions,omitempty" name:"K8sVersions"`
}

type CheckSREInstanceLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码：PASS/EXCEED_COUNT_LIMIT/BALANCE_INSUFFICIENT

		ResultCode *string `json:"ResultCode,omitempty" name:"ResultCode"`
		// 检查结果的信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSREInstanceLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSREInstanceLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewaySourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewaySourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewaySourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcInfo struct {

	// Vpc Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 内网访问地址

	IntranetAddress *string `json:"IntranetAddress,omitempty" name:"IntranetAddress"`
}

type DescribeEngineQuotasRequest struct {
	*tchttp.BaseRequest

	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeEngineQuotasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineQuotasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperMigrateClusterInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 源节点信息

		SrcNodeInfo []*DescribeZookeeperMigrateUserNodeInfo `json:"SrcNodeInfo,omitempty" name:"SrcNodeInfo"`
		// 目标节点信息

		DstNodeInfo *DescribeZookeeperMigrateTseNodeInfo `json:"DstNodeInfo,omitempty" name:"DstNodeInfo"`
		// 目标节点信息列表

		DstNodeInfoList []*DescribeZookeeperMigrateTseNodeInfo `json:"DstNodeInfoList,omitempty" name:"DstNodeInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperMigrateClusterInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperMigrateClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceStrategiesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 分页查询偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 策略名称，如果需要模糊搜索的话，最后加上一个 *

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户 ID｜用户组 ID

	PrincipalId *string `json:"PrincipalId,omitempty" name:"PrincipalId"`
	// 1 为用户，2 为用户组
	// 如果设置了PrincipalId，则必须设置PrincipalType信息
	// 如果只设置PrincipalId而不设置PrincipalType的话，默认PrincipalType = 1

	PrincipalType *int64 `json:"PrincipalType,omitempty" name:"PrincipalType"`
	// 2为返回结果只显示默认策略，1为返回结果不带默认策略, 0为返回结果为同时包含满足条件的默认以及非默认策略,
	// 这个主要用在授权，在给某个用户授权时，应该设置这个值为 1

	Default *int64 `json:"Default,omitempty" name:"Default"`
	// 资源ID

	ResId *string `json:"ResId,omitempty" name:"ResId"`
	// 资源类型，namespace | service | config_group
	// 如果设置了ResId，必须设置ResType信息
	// 如果只设置了ResId而不设置ResType的话，默认ResType=namespace

	ResType *string `json:"ResType,omitempty" name:"ResType"`
	// 是否显示该策略的详细信息（显示策略下的Resource以及Principal信息），需要显示的话，应该设置值为 true

	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`
}

func (r *DescribeGovernanceStrategiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceStrategiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户列表

		Content []*User `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServiceInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务实例记录

		Content []*DescribeZookeeperServiceInstance `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperServiceInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServiceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KubernetesInfo struct {

	// 容器集群ID

	BoundClusterId *string `json:"BoundClusterId,omitempty" name:"BoundClusterId"`
	// 容器集群类型。tke/eks

	BoundClusterType *string `json:"BoundClusterType,omitempty" name:"BoundClusterType"`
	// 同步类型。all/demand

	SyncMode *string `json:"SyncMode,omitempty" name:"SyncMode"`
	// 健康状态

	Health *bool `json:"Health,omitempty" name:"Health"`
}

type UpdateApolloEnvInfoRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 新的环境信息（目前支持修改环境描述）

	EnvInfo *EnvInfo `json:"EnvInfo,omitempty" name:"EnvInfo"`
}

func (r *UpdateApolloEnvInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApolloEnvInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConsulConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重启集群的任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateConsulConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateConsulConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEurekaConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置项Key

	ConfigKey *string `json:"ConfigKey,omitempty" name:"ConfigKey"`
	// 配置项Value，json字符串格式

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
}

func (r *UpdateEurekaConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateEurekaConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSREInstanceBasicInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSREInstanceBasicInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSREInstanceBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SRETaskPhase struct {

	// 任务执行的阶段

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 任务所处阶段执行的状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 处于当前状态的原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 处于当前状态的详细信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type CreateAutoScalerResourceStrategyRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 弹性伸缩配置列表

	Config *CloudNativeAPIGatewayStrategyAutoScalerConfig `json:"Config,omitempty" name:"Config"`
}

func (r *CreateAutoScalerResourceStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoScalerResourceStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置文件

		ConfigFile *ConfigFile `json:"ConfigFile,omitempty" name:"ConfigFile"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceUsersRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户ID列表

	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *DeleteGovernanceUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudAlarmPolarisPodsRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回列表中元素个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回列表起始元素的偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCloudAlarmPolarisPodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudAlarmPolarisPodsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewaySecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 证书详细信息

		Secret *CloudNativeAPIGatewaySecret `json:"Secret,omitempty" name:"Secret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewaySecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddGovernedK8SRequest struct {
	*tchttp.BaseRequest

	// 服务治理引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务治理引擎所在地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 服务治理引擎所在地域待绑定的kubernetes集群的信息

	BoundK8SInfos []*BoundK8SInfo `json:"BoundK8SInfos,omitempty" name:"BoundK8SInfos"`
}

func (r *AddGovernedK8SRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGovernedK8SRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域数组

		List []*OperateRegion `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置项或路径key

		Key *string `json:"Key,omitempty" name:"Key"`
		// 配置项的值

		Value *string `json:"Value,omitempty" name:"Value"`
		// 当前key是否为路径

		IsDir *bool `json:"IsDir,omitempty" name:"IsDir"`
		// 当前key下的子路径

		List []*string `json:"List,omitempty" name:"List"`
		// zookeeper节点的描述

		ZKDesc *ZookeeperRootDescription `json:"ZKDesc,omitempty" name:"ZKDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateEngineVersion struct {

	// 引擎版本号

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type OpenGovernanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开通是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenGovernanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenGovernanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GovernanceNamespace struct {

	// 命名空间名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间描述信息。

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 命名空间下总服务数据量

	TotalServiceCount *int64 `json:"TotalServiceCount,omitempty" name:"TotalServiceCount"`
	// 命名空间下总健康实例数量

	TotalHealthInstanceCount *int64 `json:"TotalHealthInstanceCount,omitempty" name:"TotalHealthInstanceCount"`
	// 命名空间下总实例数量

	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" name:"TotalInstanceCount"`
	// 命名空间ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否可以编辑

	Editable *bool `json:"Editable,omitempty" name:"Editable"`
}

type CheckZookeeperMigrateConnectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 执行结果解释信息

		ResultMessage *string `json:"ResultMessage,omitempty" name:"ResultMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckZookeeperMigrateConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckZookeeperMigrateConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNativeGatewayServiceRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 命名空间ID

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 服务名称

	Service *string `json:"Service,omitempty" name:"Service"`
}

func (r *DeleteNativeGatewayServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNativeGatewayServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNativeGatewayServiceGroupRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 命名空间

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 服务名称

	Service *string `json:"Service,omitempty" name:"Service"`
	// 实例分组名称

	Group *string `json:"Group,omitempty" name:"Group"`
}

func (r *DeleteNativeGatewayServiceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNativeGatewayServiceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolarisCLSTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修复结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolarisCLSTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolarisCLSTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetZookeeperMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetZookeeperMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetZookeeperMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvAddressInfo struct {

	// 环境名

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
	// 是否开启config公网

	EnableConfigInternet *bool `json:"EnableConfigInternet,omitempty" name:"EnableConfigInternet"`
	// config公网ip

	ConfigInternetServiceIp *string `json:"ConfigInternetServiceIp,omitempty" name:"ConfigInternetServiceIp"`
	// config内网访问地址

	ConfigIntranetAddress *string `json:"ConfigIntranetAddress,omitempty" name:"ConfigIntranetAddress"`
	// 是否开启config内网clb

	EnableConfigIntranet *bool `json:"EnableConfigIntranet,omitempty" name:"EnableConfigIntranet"`
	// 客户端公网带宽

	InternetBandWidth *int64 `json:"InternetBandWidth,omitempty" name:"InternetBandWidth"`
}

type CheckUserCloudProductRequest struct {
	*tchttp.BaseRequest

	// 云产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 引擎地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *CheckUserCloudProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckUserCloudProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEngineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务执行ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEngineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NginxIngressConfigMapParam struct {

	// 配置项的键

	Name *string `json:"Name,omitempty" name:"Name"`
	// 默认值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 配置项类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 配置项的描述信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type CloseKongIngressRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *CloseKongIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKongIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernedK8SRequest struct {
	*tchttp.BaseRequest

	// 服务治理引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务治理引擎所在地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 服务治理引擎所在地域待绑定的kubernetes集群的信息

	BoundK8SInfos []*BoundK8SInfo `json:"BoundK8SInfos,omitempty" name:"BoundK8SInfos"`
}

func (r *DeleteGovernedK8SRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernedK8SRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageApolloEnvRequest struct {
	*tchttp.BaseRequest

	// 命令（create、delete）

	Command *string `json:"Command,omitempty" name:"Command"`
	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 引擎环境信息

	EnvInfos []*EnvInfo `json:"EnvInfos,omitempty" name:"EnvInfos"`
}

func (r *ManageApolloEnvRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageApolloEnvRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceInstance struct {

	// 实例IP

	ServiceInstanceIp *string `json:"ServiceInstanceIp,omitempty" name:"ServiceInstanceIp"`
	// 实例端口

	ServiceInstancePort *int64 `json:"ServiceInstancePort,omitempty" name:"ServiceInstancePort"`
	// 实例状态

	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`
	// 服务ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 节点状态

	NodeStatus *string `json:"NodeStatus,omitempty" name:"NodeStatus"`
}

type GovernanceInstance struct {

	// 实例id。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 实例所在服务名。

	Service *string `json:"Service,omitempty" name:"Service"`
	// 实例所在命名空间名。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 实例ip地址。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 实例端口信息。

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 通信协议。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 版本信息。

	Version *string `json:"Version,omitempty" name:"Version"`
	// 负载均衡权重。

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 是否开启健康检查。

	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`
	// 实例是否健康。

	Healthy *bool `json:"Healthy,omitempty" name:"Healthy"`
	// 实例是否隔离。

	Isolate *bool `json:"Isolate,omitempty" name:"Isolate"`
	// 实例创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例修改时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 元数据数组。

	Metadatas []*Metadata `json:"Metadatas,omitempty" name:"Metadatas"`
	// 上报心跳间隔。

	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`
}

type DescribeGovernanceDiscoverEventsRequest struct {
	*tchttp.BaseRequest

	// 服务事件类型/InstanceTurnUnHealth/InstanceTurnHealth/InstanceOpenIsolate/InstanceCloseIsolate/InstanceOffline

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// tse实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开始时间时间戳，单位为秒，默认为七天前这个时间点的时间戳。

	TimeStart *uint64 `json:"TimeStart,omitempty" name:"TimeStart"`
	// 结束时间时间戳，单位为秒，默认为当前时间戳

	TimeEnd *uint64 `json:"TimeEnd,omitempty" name:"TimeEnd"`
	// 服务命名空间，默认值为default，后台默认模糊搜索

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务名称，默认为 ""，后台默认模糊搜索

	Service *string `json:"Service,omitempty" name:"Service"`
	// 影响对象，搜索条件格式为ip，后台默认模糊搜索

	AffectTarget *string `json:"AffectTarget,omitempty" name:"AffectTarget"`
}

func (r *DescribeGovernanceDiscoverEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceDiscoverEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInternetRequest struct {
	*tchttp.BaseRequest

	// tse 实例 id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 公网带宽

	BandWidth *uint64 `json:"BandWidth,omitempty" name:"BandWidth"`
	// 环境名，修改 apollo 公网带宽时必填

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
	// 部署地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *ModifyInternetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInternetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewaySecretRequest struct {
	*tchttp.BaseRequest

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 查询的Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 证书名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeCloudNativeAPIGatewaySecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceAuthStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true为开启

		Open *bool `json:"Open,omitempty" name:"Open"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceAuthStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceAuthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSREInstanceRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 注册引擎类型，支持consul

	Type *string `json:"Type,omitempty" name:"Type"`
	// 注册引擎版本，例如1.8.3

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 注册引擎实例副本数

	Replica *int64 `json:"Replica,omitempty" name:"Replica"`
	// 注册引擎实例规格ID

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// Vpc Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 是否开启持久化存储

	EnableStorage *bool `json:"EnableStorage,omitempty" name:"EnableStorage"`
	// 数据存储方式

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 云硬盘容量

	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`
	// 计费方式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 环境配置信息列表

	EnvInfos []*EnvInfo `json:"EnvInfos,omitempty" name:"EnvInfos"`
	// 是否开启公网

	EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`
	// 引擎所在的区域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 绑定的kubernetes集群信息列表

	BoundK8SInfos []*BoundK8SInfo `json:"BoundK8SInfos,omitempty" name:"BoundK8SInfos"`
	// 功能版本名。
	// - TRIAL：体验版
	// - STANDARD： 标准版
	// - PROFESSIONAL：专业版

	FeatureVersion *string `json:"FeatureVersion,omitempty" name:"FeatureVersion"`
	// 标签列表

	TagInfos []*InstanceTagInfo `json:"TagInfos,omitempty" name:"TagInfos"`
	// 是否开启北极星引擎日志上报到CLS

	ReportPolarisLogToCLS *bool `json:"ReportPolarisLogToCLS,omitempty" name:"ReportPolarisLogToCLS"`
	// 跨地域部署的引擎地域配置详情

	EngineRegionInfos []*EngineRegionInfo `json:"EngineRegionInfos,omitempty" name:"EngineRegionInfos"`
}

func (r *CreateSREInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSREInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernedK8SResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行任务的任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGovernedK8SResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernedK8SResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGovernanceAliasRequest struct {
	*tchttp.BaseRequest

	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 服务别名命名空间

	AliasNamespace *string `json:"AliasNamespace,omitempty" name:"AliasNamespace"`
	// 服务别名所指向的服务名

	Service *string `json:"Service,omitempty" name:"Service"`
	// 服务别名所指向的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务别名描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateGovernanceAliasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceAliasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperateRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadCloudNativeAPIGatewayPluginRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 插件代码压缩包zip格式转base64

	ZipFile *string `json:"ZipFile,omitempty" name:"ZipFile"`
	// 插件的COS存放信息

	CosConfig *CloudNativeAPIGatewayCosConfig `json:"CosConfig,omitempty" name:"CosConfig"`
	// 插件来源,ZipFile|Cos

	Source *string `json:"Source,omitempty" name:"Source"`
	// 插件版本号

	PluginVersion *string `json:"PluginVersion,omitempty" name:"PluginVersion"`
	// 版本描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UploadCloudNativeAPIGatewayPluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadCloudNativeAPIGatewayPluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEngineInternetAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateEngineInternetAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateEngineInternetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstancesRequest struct {
	*tchttp.BaseRequest

	// 请求过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 翻页单页查询限制数量[0,1000], 默认值0

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 翻页单页偏移量，默认值0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询类型

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// 调用方来源

	QuerySource *string `json:"QuerySource,omitempty" name:"QuerySource"`
	// 是否不加载envInfos信息

	NoEnvInfos *bool `json:"NoEnvInfos,omitempty" name:"NoEnvInfos"`
}

func (r *DescribeSREInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSREInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateFeatureVersionRequest struct {
	*tchttp.BaseRequest

	// 中国区

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 功能版本

	FeatureVersion *string `json:"FeatureVersion,omitempty" name:"FeatureVersion"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 支持的副本数列表

	Replicas *string `json:"Replicas,omitempty" name:"Replicas"`
	// 支持的规格列表

	Specs *string `json:"Specs,omitempty" name:"Specs"`
	// 功能版本名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyOperateFeatureVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateFeatureVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateFeatureVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperateFeatureVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateFeatureVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateFeatureVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperateFeatureVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateFeatureVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolarisCLSTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 北极星CLS日志主题列表

		Topics []*PolarisCLSTopic `json:"Topics,omitempty" name:"Topics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolarisCLSTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolarisCLSTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableEngineInternetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableEngineInternetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableEngineInternetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateEngineSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateEngineSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateEngineSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OptTaskBatchInfo struct {

	// 序列号

	Seq *int64 `json:"Seq,omitempty" name:"Seq"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 批次索引

	BatchId *int64 `json:"BatchId,omitempty" name:"BatchId"`
	// ins-xxx,ins-yyy

	InstanceListStr *string `json:"InstanceListStr,omitempty" name:"InstanceListStr"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type KongCertificate struct {

	// 无

	Cert *KongCertificatesPreview `json:"Cert,omitempty" name:"Cert"`
}

type DescribeSREInstanceHealthStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总副本数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 运行中的副本数量

		RunningCount *int64 `json:"RunningCount,omitempty" name:"RunningCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSREInstanceHealthStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSREInstanceHealthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInternetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInternetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInternetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayVpcConfig struct {

	// 私有网络ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DescribeEurekaReplicasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 注册引擎实例副本信息

		Replicas []*EurekaReplica `json:"Replicas,omitempty" name:"Replicas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEurekaReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEurekaReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebuildPolarisCLSTopicRequest struct {
	*tchttp.BaseRequest

	// 北极星引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// LogType

	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

func (r *RebuildPolarisCLSTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebuildPolarisCLSTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseRuleInfo struct {

	// 资源标示key，实例数量：instance_count

	Key *string `json:"Key,omitempty" name:"Key"`
	// 计算符号。Exp_Lt  Exp = "<"
	// Exp_Gt  Exp = ">"
	// Exp_Eq  Exp = "=="
	// Exp_Let Exp = "<="
	// Exp_Get Exp = ">="
	// Exp_Neq Exp = "!="

	Exp *string `json:"Exp,omitempty" name:"Exp"`
	// 约束值

	Val *string `json:"Val,omitempty" name:"Val"`
	// Val的类型，目前仅支持 int

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeEngineClbAccessAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客户端内网CLB访问地址

		ClientIntranetAddress *string `json:"ClientIntranetAddress,omitempty" name:"ClientIntranetAddress"`
		// 控制台内网CLB访问地址

		ConsoleIntranetAddress *string `json:"ConsoleIntranetAddress,omitempty" name:"ConsoleIntranetAddress"`
		// apollo多环境公网ip

		EnvAddressInfos []*EnvAddressInfo `json:"EnvAddressInfos,omitempty" name:"EnvAddressInfos"`
		// 监控数据上报内网CLB地址

		ReportIntranetAddress *string `json:"ReportIntranetAddress,omitempty" name:"ReportIntranetAddress"`
		// 北极星限流服务实例CLB地址

		LimiterAddressInfos []*PolarisLimiterAddress `json:"LimiterAddressInfos,omitempty" name:"LimiterAddressInfos"`
		// 客户端内网是否已开启

		EnableClientIntranet *bool `json:"EnableClientIntranet,omitempty" name:"EnableClientIntranet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEngineClbAccessAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineClbAccessAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateConsoleInstancesRequest struct {
	*tchttp.BaseRequest

	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 主控地域

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateConsoleInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateConsoleInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateConsoleInstanceItem struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 集群ID

	EKSClusterID *string `json:"EKSClusterID,omitempty" name:"EKSClusterID"`
	// pod列表

	PodList []*OperatePodItem `json:"PodList,omitempty" name:"PodList"`
	// service列表

	ServiceList []*OperateServiceItem `json:"ServiceList,omitempty" name:"ServiceList"`
}

type ModifyAutoScalerResourceStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoScalerResourceStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoScalerResourceStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceLabel struct {

	// 标签名称

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type UpdateZookeeperMigrateUploadMetaInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更新结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateZookeeperMigrateUploadMetaInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateZookeeperMigrateUploadMetaInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayCodeNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCloudNativeAPIGatewayCodeNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayCodeNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayNodeConfig struct {

	// 节点配置, 1c2g|2c4g|4c8g|8c16g。

	Specification *string `json:"Specification,omitempty" name:"Specification"`
	// 节点数量，2-9。

	Number *int64 `json:"Number,omitempty" name:"Number"`
}

type NetworkConfig struct {

	// 网络控制类型，比如 Kong, Konga

	ConsoleType *string `json:"ConsoleType,omitempty" name:"ConsoleType"`
	// 网络类型：Open|Internal

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 是否开启

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 访问控制策略

	AccessControl *NetworkAccessControl `json:"AccessControl,omitempty" name:"AccessControl"`
	// 节点级别控制信息

	NodeAccessControl *NodeAccessControl `json:"NodeAccessControl,omitempty" name:"NodeAccessControl"`
}

type DescribeCloudNativeAPIGatewayUserStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCloudNativeAPIGatewayUserStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayUserStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户组列表

		Content []*UserGroup `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListCloudNativeAPIGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NativeGatewayServiceItem struct {

	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务来源实例类型

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 服务来源实例ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 服务来源实例名称

	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
	// 命名空间名称

	Namespace *NamespaceInfo `json:"Namespace,omitempty" name:"Namespace"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 服务的唯一标识

	ID *string `json:"ID,omitempty" name:"ID"`
}

type DescribeZookeeperMigrateClusterInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeZookeeperMigrateClusterInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperMigrateClusterInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigFileTemplate struct {

	// 配置文件模板id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 配置文件模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 配置文件模板内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 配置文件模板格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 配置文件模板注释

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 配置文件模板创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 配置文件模板创建者

	CreateBy *string `json:"CreateBy,omitempty" name:"CreateBy"`
	// 配置文件模板修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 配置文件模板修改者

	ModifyBy *string `json:"ModifyBy,omitempty" name:"ModifyBy"`
}

type OperateEksCluster struct {

	// EKS集群ID

	EksClusterId *string `json:"EksClusterId,omitempty" name:"EksClusterId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// vpc子网id

	ClbVpcSubnetId *string `json:"ClbVpcSubnetId,omitempty" name:"ClbVpcSubnetId"`
	// 集群apiserver地址

	EksApiServerAddress *string `json:"EksApiServerAddress,omitempty" name:"EksApiServerAddress"`
	// 集群token

	EksApiServerToken *string `json:"EksApiServerToken,omitempty" name:"EksApiServerToken"`
	// vpc绑定id

	EksVpcBindId *string `json:"EksVpcBindId,omitempty" name:"EksVpcBindId"`
	// 已用ip数量

	UseCount *int64 `json:"UseCount,omitempty" name:"UseCount"`
	// 全部ip数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 引擎类型 混用/注册中心/配置中心/北极星

	EngineTypeText *string `json:"EngineTypeText,omitempty" name:"EngineTypeText"`
	// 是否隔离

	StatusText *string `json:"StatusText,omitempty" name:"StatusText"`
	// EKS集群子网列表

	SubnetList []*OperateSubnetInfo `json:"SubnetList,omitempty" name:"SubnetList"`
	// 服务网段

	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`
	// CLBIP可用数

	ClbAvailableCount *int64 `json:"ClbAvailableCount,omitempty" name:"ClbAvailableCount"`
	// PODIP可用数

	PodAvailableCount *int64 `json:"PodAvailableCount,omitempty" name:"PodAvailableCount"`
	// 集群当前POD数

	PodNumber *int64 `json:"PodNumber,omitempty" name:"PodNumber"`
	// 集群当前POD配额数

	PodQuota *int64 `json:"PodQuota,omitempty" name:"PodQuota"`
}

type DeleteEngineSyncJobRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据同步任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteEngineSyncJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEngineSyncJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperateRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceStatus struct {

	// 服务状态的中文名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务状态的编码

	Code *string `json:"Code,omitempty" name:"Code"`
}

type DescribeNativeGatewayServerGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分组列表信息

		Result *NativeGatewayServerGroups `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNativeGatewayServerGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNativeGatewayServerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关名字, 最多支持60个字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云原生API网关描述信息, 最多支持120个字符。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否开启 CLS 日志。暂时取值只能是 true，即只能从关闭状态变成开启状态。

	EnableCls *bool `json:"EnableCls,omitempty" name:"EnableCls"`
	// 公网计费模式。可选取值 BANDWIDTH | TRAFFIC ，表示按带宽和按流量计费。

	InternetPayMode *string `json:"InternetPayMode,omitempty" name:"InternetPayMode"`
}

func (r *ModifyCloudNativeAPIGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigFilesRequest struct {
	*tchttp.BaseRequest

	// ins-df344df5

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置文件列表

	ConfigFile *ConfigFile `json:"ConfigFile,omitempty" name:"ConfigFile"`
}

func (r *ModifyConfigFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateZooKeeperConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateZooKeeperConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateZooKeeperConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSREInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 注册引擎实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSREInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSREInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNginxIngressNginxRawConfRequest struct {
	*tchttp.BaseRequest

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *GetNginxIngressNginxRawConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNginxIngressNginxRawConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenKongIngressRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群的类型：tke或eks

	Type *string `json:"Type,omitempty" name:"Type"`
	// Ingress版本

	IngressVersion *string `json:"IngressVersion,omitempty" name:"IngressVersion"`
	// K8s集群版本

	K8sClusterVersion *string `json:"K8sClusterVersion,omitempty" name:"K8sClusterVersion"`
}

func (r *OpenKongIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKongIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KongIngressStatus struct {

	// 是否开启了KongIngress

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 开启了KongIngress对接的集群

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群的类型：tke或eks

	Type *string `json:"Type,omitempty" name:"Type"`
	// Kong Ingress版本号

	IngressVersion *string `json:"IngressVersion,omitempty" name:"IngressVersion"`
}

type StepInfo struct {

	// 步骤编号

	StepNo *int64 `json:"StepNo,omitempty" name:"StepNo"`
	// 步骤名称

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 步骤开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type DescribeGovernanceInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务实例总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务里实例列表。

		Content []*GovernanceInstance `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServiceListRequest struct {
	*tchttp.BaseRequest

	// 请求过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 注册引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeZookeeperServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartZookeeperMigrateMergeClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartZookeeperMigrateMergeClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartZookeeperMigrateMergeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZooKeeperConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeZooKeeperConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZooKeeperConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplyLicenseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplyLicenseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplyLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGovernanceNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCredentialRequest struct {
	*tchttp.BaseRequest

	// 用户凭证

	TenancyKey *string `json:"TenancyKey,omitempty" name:"TenancyKey"`
}

func (r *DescribeCredentialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCredentialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayPluginsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 按照插件名称模糊匹配。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *ListCloudNativeAPIGatewayPluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayPluginsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeOperateApolloDeployTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeOperateApolloDeployTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeOperateApolloDeployTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateConsoleInstanceLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志内容

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateConsoleInstanceLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateConsoleInstanceLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServerInterfacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口总个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 接口列表

		Content []*ZookeeperServerInterface `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperServerInterfacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServerInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateInstancesRequest struct {
	*tchttp.BaseRequest

	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEurekaConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重启集群的任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateEurekaConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateEurekaConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindCLBSecurityGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定成功的白名单列表

		WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindCLBSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindCLBSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigFileReleaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置文件发布详情

		ConfigFileRelease *ConfigFileRelease `json:"ConfigFileRelease,omitempty" name:"ConfigFileRelease"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigFileReleaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFileReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEurekaServiceInstancesRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 注册服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeEurekaServiceInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEurekaServiceInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSimpleInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表

		InstanceList []*TagInstanceDescription `json:"InstanceList,omitempty" name:"InstanceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceSimpleInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceSimpleInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolarisReplicasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 北极星引擎实例副本信息

		Replicas []*PolarisReplica `json:"Replicas,omitempty" name:"Replicas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolarisReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolarisReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineQuotasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎额度列表

		EngineQuotas []*EngineQuota `json:"EngineQuotas,omitempty" name:"EngineQuotas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEngineQuotasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayDomain struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 关联的tke集群信息

	BindClusters []*CloudNativeAPIGatewayDomainCluster `json:"BindClusters,omitempty" name:"BindClusters"`
	// 关联的证书信息

	BindSecrets []*CloudNativeAPIGatewaySecret `json:"BindSecrets,omitempty" name:"BindSecrets"`
}

type DescribeConfigFileReleaseHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置文件发布历史列表

		ConfigFileReleaseHistories []*ConfigFileReleaseHistory `json:"ConfigFileReleaseHistories,omitempty" name:"ConfigFileReleaseHistories"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigFileReleaseHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFileReleaseHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternetConfig struct {

	// 公网地址版本，可选："IPV4" | "IPV6-NAT64" 。不填默认 IPV4 。

	InternetAddressVersion *string `json:"InternetAddressVersion,omitempty" name:"InternetAddressVersion"`
	// 公网付费类型，当前仅可选："BANDWIDTH"。不填默认为 "BANDWIDTH"

	InternetPayMode *string `json:"InternetPayMode,omitempty" name:"InternetPayMode"`
	// 公网带宽。

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 负载均衡描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 负载均衡的规格类型，传 "SLA" 表示性能容量型，不传为共享型。

	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`
}

type DeleteOperateEngineSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperateEngineSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateEngineSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigFileReleaseHistoriesRequest struct {
	*tchttp.BaseRequest

	// TSE实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 发布历史记录id，用于分页优化，一般指定 EndId，就不用指定 Offset，否则分页可能不连续

	EndId *uint64 `json:"EndId,omitempty" name:"EndId"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConfigFileReleaseHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFileReleaseHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayNetworkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolarisCLSTopicRequest struct {
	*tchttp.BaseRequest

	// 北极星引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Event/History/Monitor

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 日志保存天数，-1为永久保存

	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *ModifyPolarisCLSTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolarisCLSTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateZookeeperMigrateUploadMetaInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据MD5

	DataMD5 *string `json:"DataMD5,omitempty" name:"DataMD5"`
}

func (r *UpdateZookeeperMigrateUploadMetaInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateZookeeperMigrateUploadMetaInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperMigrateUserNodeInfo struct {

	// 源ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 客户端端口

	ClientPort *int64 `json:"ClientPort,omitempty" name:"ClientPort"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 事务id

	ZXID *string `json:"ZXID,omitempty" name:"ZXID"`
	// 与Leader事务ID 的差值

	Diff *int64 `json:"Diff,omitempty" name:"Diff"`
	// zookeeper的myid

	MYID *string `json:"MYID,omitempty" name:"MYID"`
	// 选举端口

	ElectionPort *int64 `json:"ElectionPort,omitempty" name:"ElectionPort"`
	// 集群端口

	ClusterPort *int64 `json:"ClusterPort,omitempty" name:"ClusterPort"`
}

type FeatureVersionInfo struct {

	// 功能版本，可选以下值：
	//
	// - TRIAL：体验版
	// - STANDARD： 标准版
	// - PROFESSIONAL：专业版

	FeatureVersion *string `json:"FeatureVersion,omitempty" name:"FeatureVersion"`
	// 该功能版本支持的节点数的列表。

	Replicas []*uint64 `json:"Replicas,omitempty" name:"Replicas"`
	// 该功能版本支持的规格列表。

	Specs []*string `json:"Specs,omitempty" name:"Specs"`
	// 该版本是否可售

	Available *bool `json:"Available,omitempty" name:"Available"`
}

type CreateCloudNativeAPIGatewayCertificateRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 证书名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 证书私钥

	Key *string `json:"Key,omitempty" name:"Key"`
	// 证书pem格式

	Crt *string `json:"Crt,omitempty" name:"Crt"`
	// 绑定的域名

	BindDomains []*string `json:"BindDomains,omitempty" name:"BindDomains"`
	// ssl平台证书 Id

	CertId *string `json:"CertId,omitempty" name:"CertId"`
}

func (r *CreateCloudNativeAPIGatewayCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceUserTokenRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户ID

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeGovernanceUserTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceUserTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperateApolloInstancesRequest struct {
	*tchttp.BaseRequest

	// 更新的实例列表

	InstanceIDList []*string `json:"InstanceIDList,omitempty" name:"InstanceIDList"`
	// Portal 的 yaml 参数信息

	PortalService *OperateApolloDeploymentParams `json:"PortalService,omitempty" name:"PortalService"`
	// ConfigService 的 yaml 参数信息

	ConfigService *OperateApolloDeploymentParams `json:"ConfigService,omitempty" name:"ConfigService"`
	// AdminService 的 yaml 参数信息

	AdminService *OperateApolloDeploymentParams `json:"AdminService,omitempty" name:"AdminService"`
}

func (r *UpdateOperateApolloInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperateApolloInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewaySystemPluginsRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 搜索关键字

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeCloudNativeAPIGatewaySystemPluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySystemPluginsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayCanaryRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayCanaryRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayCanaryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceGroupTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceGroupTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceGroupTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNginxIngressConfigMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNginxIngressConfigMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNginxIngressConfigMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorDataPoint struct {

	// 顺序索引

	Index *uint64 `json:"Index,omitempty" name:"Index"`
	// 图例名称

	Legend *string `json:"Legend,omitempty" name:"Legend"`
	// 计量单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 时间戳数组

	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps"`
	// 监控值数组

	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type DeleteConfigFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayVersionRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例类型，目前支持 kong

	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`
}

func (r *DescribeCloudNativeAPIGatewayVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayConfig struct {

	// 控制台类型。

	ConsoleType *string `json:"ConsoleType,omitempty" name:"ConsoleType"`
	// HTTP链接地址。

	HttpUrl *string `json:"HttpUrl,omitempty" name:"HttpUrl"`
	// HTTPS链接地址。

	HttpsUrl *string `json:"HttpsUrl,omitempty" name:"HttpsUrl"`
	// 网络类型, Open|Internal。

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 管理员用户名。

	AdminUser *string `json:"AdminUser,omitempty" name:"AdminUser"`
	// 管理员密码。

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 网络状态, Open|Closed|Updating

	Status *string `json:"Status,omitempty" name:"Status"`
	// 网络访问策略

	AccessControl *NetworkAccessControl `json:"AccessControl,omitempty" name:"AccessControl"`
	// 内网子网 ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 内网VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 节点级别的访问控制

	NodeAccessControl *NodeAccessControl `json:"NodeAccessControl,omitempty" name:"NodeAccessControl"`
	// 负载均衡的描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 负载均衡的规格类型，传 "SLA" 表示性能容量型，返回空为共享型

	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`
}

type CreateCloudNativeAPIGatewayServerGroupResult struct {

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 分组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type PayModeInfo struct {

	// 计费方式类型

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 计费方式名称

	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
}

type DescribeFeatureVersionInfosRequest struct {
	*tchttp.BaseRequest

	// 引擎类型。
	// - zookeeper

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeFeatureVersionInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFeatureVersionInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunExecTaskBatchJobRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务批次索引

	BatchIndex *uint64 `json:"BatchIndex,omitempty" name:"BatchIndex"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *RunExecTaskBatchJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunExecTaskBatchJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncOperateAgentRequest struct {
	*tchttp.BaseRequest

	// 引擎ID数组

	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
	// 镜像地址

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 环境变量数组

	EnvList []*KVPair `json:"EnvList,omitempty" name:"EnvList"`
	// China

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *SyncOperateAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncOperateAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCloudNativeAPIGatewaySpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更新云原生API网关实例规格的响应结果。

		Result *UpdateCloudNativeAPIGatewayResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCloudNativeAPIGatewaySpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCloudNativeAPIGatewaySpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GovernanceServiceInput struct {

	// 服务名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务所属命名空间。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务描述信息。

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 服务元数据。

	Metadatas []*Metadata `json:"Metadatas,omitempty" name:"Metadatas"`
	// 服务所属部门。

	Department *string `json:"Department,omitempty" name:"Department"`
	// 服务所属业务。

	Business *string `json:"Business,omitempty" name:"Business"`
	// 被添加进来可以操作此命名空间的用户ID列表

	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
	// 被添加进来可以操作此命名空间的用户组ID列表

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// 从操作此命名空间的用户组ID列表被移除的ID列表

	RemoveUserIds []*string `json:"RemoveUserIds,omitempty" name:"RemoveUserIds"`
	// 从可以操作此命名空间的用户组ID列表中被移除的ID列表

	RemoveGroupIds []*string `json:"RemoveGroupIds,omitempty" name:"RemoveGroupIds"`
}

type DescribeOperateConsoleInstanceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 平台层实例详情

		Data *OperateConsoleInstanceItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateConsoleInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateConsoleInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrometheusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Result *PrometheusStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrometheusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrometheusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadLicenseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件名

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// license的内容

		CipherText *string `json:"CipherText,omitempty" name:"CipherText"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadLicenseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetZookeeperMigrateRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ResetZookeeperMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetZookeeperMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindAutoScalerResourceStrategyFromGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindAutoScalerResourceStrategyFromGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoScalerResourceStrategyFromGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGovernanceInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateRegionRequest struct {
	*tchttp.BaseRequest

	// 地域

	RegionField *string `json:"RegionField,omitempty" name:"RegionField"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DeleteOperateRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExecTaskRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，支持taskId的过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 条件个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 条件偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeExecTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExecTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadZookeeperMigrateDataRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *LoadZookeeperMigrateDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoadZookeeperMigrateDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 地域

	RegionField *string `json:"RegionField,omitempty" name:"RegionField"`
	// 用户uin，非必选

	UinField *string `json:"UinField,omitempty" name:"UinField"`
	// 开放的端口，多个以英文逗号分割

	Ports *string `json:"Ports,omitempty" name:"Ports"`
	// 主控地域

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *ModifyOperateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayUserAuthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否有TSE_QCSRole的角色

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayUserAuthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayUserAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServiceInstancesRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 注册服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeZookeeperServiceInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServiceInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayResult struct {

	// 总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 云原生API网关实例列表。

	GatewayList []*DescribeCloudNativeAPIGatewayResult `json:"GatewayList,omitempty" name:"GatewayList"`
}

type CreateCloudNativeAPIGatewaySecretRequest struct {
	*tchttp.BaseRequest

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// tke集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 证书所属的 Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// Secret名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务端证书

	Crt *string `json:"Crt,omitempty" name:"Crt"`
	// 服务端私钥

	Key *string `json:"Key,omitempty" name:"Key"`
	// 根证书，只有在双向认证时才需要

	RootCrt *string `json:"RootCrt,omitempty" name:"RootCrt"`
}

func (r *CreateCloudNativeAPIGatewaySecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewaySecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNativeGatewayServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNativeGatewayServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNativeGatewayServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置项或路径key

		Key *string `json:"Key,omitempty" name:"Key"`
		// 配置项的值

		Value *string `json:"Value,omitempty" name:"Value"`
		// 当前key是否为路径

		IsDir *bool `json:"IsDir,omitempty" name:"IsDir"`
		// 当前key下的子路径

		List []*string `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayPluginVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关插件版本信息列表响应结果。

		Result *ListCloudNativeAPIGatewayPluginVersionsResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewayPluginVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayPluginVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZNodeOption struct {

	// 节点类型

	NodeType *int64 `json:"NodeType,omitempty" name:"NodeType"`
}

type CreateCloudNativeAPIGatewayCanaryRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayCanaryRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayCanaryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEngineRequest struct {
	*tchttp.BaseRequest

	// 引擎类型。参考值：
	// - zookeeper
	// - nacos
	// - consul
	// - apollo
	// - eureka
	// - polaris

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// 引擎的开源版本。每种引擎支持的开源版本不同，请参考产品文档或者控制台购买页

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 引擎的产品版本。参考值：
	// - STANDARD： 标准版
	//
	// 引擎各版本及可选择的规格、节点数说明：
	// apollo - STANDARD版本
	// 规格列表：1C2G、2C4G、4C8G、8C16G、16C32G
	// 节点数：1，2，3，4，5
	//
	// eureka - STANDARD版本
	// 规格列表：1C2G、2C4G、4C8G、8C16G、16C32G
	// 节点数：3，4，5
	//
	// polarismesh - STANDARD版本
	// 规格列表：NUM50、NUM100、NUM200、NUM500、NUM1000、NUM5000、NUM10000、NUM50000
	//
	// 兼容原spec-xxxxxx形式的规格ID

	EngineProductVersion *string `json:"EngineProductVersion,omitempty" name:"EngineProductVersion"`
	// 引擎所在地域。参考值说明：
	// 中国区 参考值：
	// - ap-guangzhou：广州
	// - ap-beijing：北京
	// - ap-chengdu：成都
	// - ap-chongqing：重庆
	// - ap-nanjing：南京
	// - ap-shanghai：上海
	// - ap-hongkong：香港
	// - ap-taipei：台北
	// 亚太区 参考值：
	// - ap-jakarta：雅加达
	// - ap-singapore：新加坡
	// 北美区 参考值
	// - na-toronto：多伦多
	// 金融专区 参考值
	// - ap-beijing-fsi：北京金融
	// - ap-shanghai-fsi：上海金融
	// - ap-shenzhen-fsi：深圳金融

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 引擎的节点规格 ID。参见EngineProductVersion字段说明

	EngineResourceSpec *string `json:"EngineResourceSpec,omitempty" name:"EngineResourceSpec"`
	// 引擎的节点数量。参见EngineProductVersion字段说明

	EngineNodeNum *int64 `json:"EngineNodeNum,omitempty" name:"EngineNodeNum"`
	// VPC ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - vpc-conz6aix

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网 ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - subnet-ahde9me9

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// Apollo 环境配置参数列表。参数说明：
	// 如果创建Apollo类型，此参数为必填的环境信息列表，最多可选4个环境。环境信息参数说明：
	// - Name：环境名。参考值：prod, dev, fat, uat
	// - EngineResourceSpec：环境内引擎的节点规格ID。参见EngineProductVersion参数说明
	// - EngineNodeNum：环境内引擎的节点数量。参见EngineProductVersion参数说明，其中prod环境支持的节点数为2，3，4，5
	// - StorageCapacity：配置存储空间大小，以GB为单位，步长为5.参考值：35
	// - VpcId：VPC ID。参考值：vpc-conz6aix
	// - SubnetId：子网 ID。参考值：subnet-ahde9me9

	ApolloEnvParams []*ApolloEnvParam `json:"ApolloEnvParams,omitempty" name:"ApolloEnvParams"`
	// 引擎名称。参考值：
	// - eurek-test

	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`
	// 引擎的标签列表。用户自定义的key/value形式，无参考值

	EngineTags []*InstanceTagInfo `json:"EngineTags,omitempty" name:"EngineTags"`
	// 引擎的初始帐号信息。可设置参数：
	// - Name：控制台初始用户名
	// - Password：控制台初始密码
	// - Token：引擎接口的管理员 Token

	EngineAdmin *EngineAdmin `json:"EngineAdmin,omitempty" name:"EngineAdmin"`
	// 付费类型。参考值：
	// - 0：后付费
	// - 1：预付费（接口暂不支持创建预付费实例）

	TradeType *int64 `json:"TradeType,omitempty" name:"TradeType"`
	// 预付费时长，以月为单位

	PrepaidPeriod *int64 `json:"PrepaidPeriod,omitempty" name:"PrepaidPeriod"`
	// 自动续费标记，仅预付费使用。参考值：
	// - 0：不自动续费
	// - 1：自动续费

	PrepaidRenewFlag *int64 `json:"PrepaidRenewFlag,omitempty" name:"PrepaidRenewFlag"`
	// 请求来源。用户自定义的请求来源标识，无参考值

	Source *string `json:"Source,omitempty" name:"Source"`
	// 是否开启北极星的日志上报到CLS

	ReportPolarisLogToCLS *bool `json:"ReportPolarisLogToCLS,omitempty" name:"ReportPolarisLogToCLS"`
	// 跨地域部署的引擎地域配置详情

	EngineRegionInfos []*EngineRegionInfo `json:"EngineRegionInfos,omitempty" name:"EngineRegionInfos"`
}

func (r *CreateEngineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEngineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartZookeeperMigratePrepareResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartZookeeperMigratePrepareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartZookeeperMigratePrepareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayCanaryRuleRequest struct {
	*tchttp.BaseRequest

	// 网关 ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务 ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *DeleteCloudNativeAPIGatewayCanaryRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayCanaryRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNacosMonitorDashboardRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 监控类型
	// system:系统概览
	// business:业务概览

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeNacosMonitorDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosMonitorDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagCloudNativeAPIGatewayDescription struct {

	// 云原生网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生网关类型

	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`
	// 云原生网关名称

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
	// 租户ID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type CloudNativeAPIGatewaySystemParameter struct {

	// 参数名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数内容。

	Value *string `json:"Value,omitempty" name:"Value"`
	// 参数默认值。

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 参数可修改值。

	ModifiableValue *string `json:"ModifiableValue,omitempty" name:"ModifiableValue"`
	// 是否需要重启生效。

	NeedReload *bool `json:"NeedReload,omitempty" name:"NeedReload"`
}

type UpdateInstanceSpecForApolloResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInstanceSpecForApolloResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceSpecForApolloResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApolloEnvParam struct {

	// 环境名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 环境内引擎的节点规格 ID
	// -1C2G
	// -2C4G
	// 兼容原spec-xxxxxx形式的规格ID

	EngineResourceSpec *string `json:"EngineResourceSpec,omitempty" name:"EngineResourceSpec"`
	// 环境内引擎的节点数量

	EngineNodeNum *int64 `json:"EngineNodeNum,omitempty" name:"EngineNodeNum"`
	// 配置存储空间大小，以GB为单位

	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`
	// VPC ID。在 VPC 的子网内分配一个 IP 作为 ConfigServer 的访问地址

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网 ID。在 VPC 的子网内分配一个 IP 作为 ConfigServer 的访问地址

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 环境描述

	EnvDesc *string `json:"EnvDesc,omitempty" name:"EnvDesc"`
}

type BreakZookeeperMigrateClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BreakZookeeperMigrateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BreakZookeeperMigrateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEurekaServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务记录

		Content []*DescribeEurekaService `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEurekaServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEurekaServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApolloAdminInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApolloAdminInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApolloAdminInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableInternetRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// true开 false关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// apollo env name

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
	// 客户端公网白名单

	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
	// 公网带宽

	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`
	// 部署地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *EnableInternetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableInternetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindAutoScalerResourceStrategyToGroupsRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 策略ID

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 网关分组ID列表

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

func (r *BindAutoScalerResourceStrategyToGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoScalerResourceStrategyToGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeNacosAuthStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeNacosAuthStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeNacosAuthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckZookeeperMigrateConnectionRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群信息

	ClusterInfoList []*ZookeeperMigrateMetaClusterInfo `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`
}

func (r *CheckZookeeperMigrateConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckZookeeperMigrateConnectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateEksListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域数组

		List []*OperateEksCluster `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateEksListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateEksListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEngineInternetAccessRequest struct {
	*tchttp.BaseRequest

	// 引擎ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否开启客户端公网访问，true开 false关

	EnableClientInternetAccess *bool `json:"EnableClientInternetAccess,omitempty" name:"EnableClientInternetAccess"`
	// 引擎类型

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
}

func (r *UpdateEngineInternetAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateEngineInternetAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeOperateApolloDeployTypeRequest struct {
	*tchttp.BaseRequest

	// 需要扩容的引擎实例ID集

	InstanceIDList []*string `json:"InstanceIDList,omitempty" name:"InstanceIDList"`
}

func (r *ChangeOperateApolloDeployTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeOperateApolloDeployTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayCertificatesRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 列表数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 列表offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件，多个过滤条件之间是与的关系，支持BindDomain ，Name

	Filters []*ListFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListCloudNativeAPIGatewayCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayCertificatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceUserTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceUserTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceUserTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateEngineVersionListRequest struct {
	*tchttp.BaseRequest

	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 总量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperateEngineVersionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateEngineVersionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateSpecListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规格数组

		List []*SRESpec `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateSpecListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateSpecListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayCanaryRulesRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务 ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 列表数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 列表offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListCloudNativeAPIGatewayCanaryRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayCanaryRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewaySourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云原生网关ID

		GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
		// 服务来源列表

		SourceList []*Source `json:"SourceList,omitempty" name:"SourceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewaySourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsulServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务记录

		Content []*DescribeService `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsulServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 区域相关信息列表

		RegionInfos []*RegionInfo `json:"RegionInfos,omitempty" name:"RegionInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrModifyCloudNativeAPIGatewayIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrModifyCloudNativeAPIGatewayIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrModifyCloudNativeAPIGatewayIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayPluginVersionsRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 插件名称。

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCloudNativeAPIGatewayPluginVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayPluginVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateEngineVersionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎版本信息

		List []*OperateEngineVersion `json:"List,omitempty" name:"List"`
		// 列表总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateEngineVersionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateEngineVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageApolloEnvResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageApolloEnvResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageApolloEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsulConfigInfo struct {

	// 配置项Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Gossip加密

	Title *string `json:"Title,omitempty" name:"Title"`
	// 配置项内容

	Value *string `json:"Value,omitempty" name:"Value"`
	// 配置描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 配置说明文档

	DocLink *string `json:"DocLink,omitempty" name:"DocLink"`
	// 配置变更提示

	ChangeAlert *string `json:"ChangeAlert,omitempty" name:"ChangeAlert"`
	// 配置项说明列表

	DescList []*ConsulConfigDesc `json:"DescList,omitempty" name:"DescList"`
}

type LogResult struct {

	// 时间戳

	Time *uint64 `json:"Time,omitempty" name:"Time"`
	// 日志内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreateExecTaskRequest struct {
	*tchttp.BaseRequest

	// 运营平台任务描述

	TaskComment *string `json:"TaskComment,omitempty" name:"TaskComment"`
	// 运营平台任务命令

	ExeCommand *string `json:"ExeCommand,omitempty" name:"ExeCommand"`
	// 可执行的实例信息列表

	ExecInstanceInfoList []*ExecInstanceInfo `json:"ExecInstanceInfoList,omitempty" name:"ExecInstanceInfoList"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 操作人名称

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
}

func (r *CreateExecTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExecTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineSyncJobCheckResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 校验结果：failed | success

		Status *string `json:"Status,omitempty" name:"Status"`
		// 步骤总数

		StepCount *uint64 `json:"StepCount,omitempty" name:"StepCount"`
		// 当前所在步骤

		StepCur *uint64 `json:"StepCur,omitempty" name:"StepCur"`
		// 总体进度

		Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
		// 步骤信息

		StepInfos []*EngineSyncJobCheckStepInfo `json:"StepInfos,omitempty" name:"StepInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEngineSyncJobCheckResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineSyncJobCheckResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Result *KongServiceList `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewayServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsulConfigDesc struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否支持编辑

	EditEnable *bool `json:"EditEnable,omitempty" name:"EditEnable"`
	// 配置项默认值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 可选值

	OptionValue *string `json:"OptionValue,omitempty" name:"OptionValue"`
	// 配置项提示

	Tips *string `json:"Tips,omitempty" name:"Tips"`
}

type UpdateInstanceSpecForApolloRequest struct {
	*tchttp.BaseRequest

	// 微服务引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 修改类型（规格/节点数）

	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`
	// Apollo实例环境信息

	EnvInfo *EnvInfo `json:"EnvInfo,omitempty" name:"EnvInfo"`
}

func (r *UpdateInstanceSpecForApolloRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceSpecForApolloRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsulReplica struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 可用区ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type CreateAutoScalerResourceStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAutoScalerResourceStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoScalerResourceStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayUpstreamRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeCloudNativeAPIGatewayUpstreamRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayUpstreamRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallCloudNativeAPIGatewaySystemPluginRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 插件名称

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 插件版本号

	PluginVersion *string `json:"PluginVersion,omitempty" name:"PluginVersion"`
}

func (r *InstallCloudNativeAPIGatewaySystemPluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallCloudNativeAPIGatewaySystemPluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayServicesRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 列表数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 列表offset

	Offset *string `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListCloudNativeAPIGatewayServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceAliasRequest struct {
	*tchttp.BaseRequest

	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 服务别名命名空间

	AliasNamespace *string `json:"AliasNamespace,omitempty" name:"AliasNamespace"`
	// 服务别名所指向的服务名

	Service *string `json:"Service,omitempty" name:"Service"`
	// 服务别名所指向的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务别名描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *ModifyGovernanceAliasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceAliasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineAdmin struct {

	// 控制台初始用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 控制台初始密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 引擎接口的管理员 Token

	Token *string `json:"Token,omitempty" name:"Token"`
}

type PolarisCLSTopic struct {

	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 日志集ID

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 日志集名称

	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 保存天数，-1为永久保存

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 日志主题的标签

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 日志类型描述

	LogTypeInfo *string `json:"LogTypeInfo,omitempty" name:"LogTypeInfo"`
}

type CreateCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest

	// 可用区名称, 例如，ap-guangzhou-3。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 云原生API网关名字, 最多支持60个字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云原生API网关类型, 目前只支持Kong。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 云原生API网关版本, 目前只支持2.4.1。

	GatewayVersion *string `json:"GatewayVersion,omitempty" name:"GatewayVersion"`
	// 云原生API网关节点配置。

	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitempty" name:"NodeConfig"`
	// 云原生API网关vpc配置。

	VpcConfig *CloudNativeAPIGatewayVpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`
	// 云原生API网关描述信息, 最多支持120个字符。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 标签列表

	Tags []*InstanceTagInfo `json:"Tags,omitempty" name:"Tags"`
	// 是否开启 CLS 日志

	EnableCls *bool `json:"EnableCls,omitempty" name:"EnableCls"`
	// 实例版本，当前支持开发版和标准版【TRIAL、STANDARD】

	FeatureVersion *string `json:"FeatureVersion,omitempty" name:"FeatureVersion"`
	// 公网出流量带宽，[1,2048]Mbps

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 实例实际的地域信息

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// ingress Class名称

	IngressClassName *string `json:"IngressClassName,omitempty" name:"IngressClassName"`
	// 公网相关配置

	InternetConfig *InternetConfig `json:"InternetConfig,omitempty" name:"InternetConfig"`
}

func (r *CreateCloudNativeAPIGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateApolloDeploymentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperateApolloDeploymentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateApolloDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApolloAdminInfoRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 修改类型（password：管理员密码；token：管理员token）

	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`
	// 自定义密码

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
	// 自定义token

	NewToken *string `json:"NewToken,omitempty" name:"NewToken"`
}

func (r *ModifyApolloAdminInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApolloAdminInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayCanaryRuleRequest struct {
	*tchttp.BaseRequest

	// 网关 ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务 ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 灰度规则配置

	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitempty" name:"CanaryRule"`
}

func (r *CreateCloudNativeAPIGatewayCanaryRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayCanaryRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateInstanceLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志内容

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateInstanceLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateInstanceLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateAgentInfo struct {

	// 引擎ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 引擎类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 镜像地址

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 镜像tag

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 更新时间

	UpdateAt *string `json:"UpdateAt,omitempty" name:"UpdateAt"`
	// 创建时间

	CreateAt *string `json:"CreateAt,omitempty" name:"CreateAt"`
	// 集群名称

	EKSClusterID *string `json:"EKSClusterID,omitempty" name:"EKSClusterID"`
	// 实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 用户id

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type SourceInstanceAuth struct {

	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 账户密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 访问凭据 token

	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`
}

type DescribeCloudNativeAPIGatewayPluginsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关插件列表响应结果。

		Result *ListCloudNativeAPIGatewayPluginsResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayPluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsulConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参数配置列表

		ConfigList []*ConsulConfigInfo `json:"ConfigList,omitempty" name:"ConfigList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsulConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEurekaConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参数配置列表

		ConfigList []*EurekaConfigInfo `json:"ConfigList,omitempty" name:"ConfigList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEurekaConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEurekaConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageConsulConsoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应数据

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageConsulConsoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageConsulConsoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCloudNativeAPIGatewayDomainsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名信息列表

		Domains []*CloudNativeAPIGatewayDomain `json:"Domains,omitempty" name:"Domains"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCloudNativeAPIGatewayDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCloudNativeAPIGatewayDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayCanaryRuleCondition struct {

	// 条件类型，支持 path, method, query, header, cookie, body 和 system。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 参数名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 操作符，支持 "le", "eq", "lt", "ne", "ge", "gt", "regex", "exists", "prefix" ,"exact", "regex" 等

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 目标参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type KongCertificatesList struct {

	// 证书列表总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 无

	CertificatesList []*KongCertificatesPreview `json:"CertificatesList,omitempty" name:"CertificatesList"`
	// 证书列表总页数

	Pages *int64 `json:"Pages,omitempty" name:"Pages"`
}

type CreateOperateWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperateWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayClsConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关实例关联的CLS配置信息响应结果。

		Result *DescribeCloudNativeAPIGatewayClsConfigResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayClsConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayClsConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayClsConfigRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关实例CLS配置信息。

	ClsConfig *CloudNativeAPIGatewayClsConfig `json:"ClsConfig,omitempty" name:"ClsConfig"`
}

func (r *ModifyCloudNativeAPIGatewayClsConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayClsConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkAccessControl struct {

	// 访问模式：Whitelist|Blacklist

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 白名单列表

	CidrWhiteList []*string `json:"CidrWhiteList,omitempty" name:"CidrWhiteList"`
	// 黑名单列表

	CidrBlackList []*string `json:"CidrBlackList,omitempty" name:"CidrBlackList"`
}

type DescribeInstanceTagInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例标签集合

		TagInfos []*InstanceTagInfo `json:"TagInfos,omitempty" name:"TagInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTagInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTagInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMainControlRegionInfosRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMainControlRegionInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMainControlRegionInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateConsoleInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	EKSClusterID *string `json:"EKSClusterID,omitempty" name:"EKSClusterID"`
	// 主控

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateConsoleInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateConsoleInstanceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExecTaskBatchJobExeInfoRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 批次索引

	BatchIndex *int64 `json:"BatchIndex,omitempty" name:"BatchIndex"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeExecTaskBatchJobExeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExecTaskBatchJobExeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KongUpstreamList struct {

	// 无

	UpstreamList []*KongUpstreamPreview `json:"UpstreamList,omitempty" name:"UpstreamList"`
}

type DescribeEurekaReplicasRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 副本列表Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 副本列表Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeEurekaReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEurekaReplicasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishConfigFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置文件发布是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishConfigFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishConfigFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KongUpstreamInfo struct {

	// IP或域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 服务来源ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 名字空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务（注册中心或Kubernetes中的服务）名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务后端类型是IPList时提供

	Targets []*KongTarget `json:"Targets,omitempty" name:"Targets"`
	// 服务来源类型

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// SCF函数类型

	ScfType *string `json:"ScfType,omitempty" name:"ScfType"`
	// SCF函数命名空间

	ScfNamespace *string `json:"ScfNamespace,omitempty" name:"ScfNamespace"`
	// SCF函数名

	ScfLambdaName *string `json:"ScfLambdaName,omitempty" name:"ScfLambdaName"`
	// SCF函数版本

	ScfLambdaQualifier *string `json:"ScfLambdaQualifier,omitempty" name:"ScfLambdaQualifier"`
}

type DescribeFeatureVersionInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 功能版本的规格信息。

		FeatureVersionInfos []*FeatureVersionInfo `json:"FeatureVersionInfos,omitempty" name:"FeatureVersionInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFeatureVersionInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFeatureVersionInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperMigratePhaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例迁移元数据

		MetaInfo *ZookeeperMigrateMetaInfo `json:"MetaInfo,omitempty" name:"MetaInfo"`
		// 各阶段执行信息

		MigratePhaseInfoList []*DescribeZookeeperMigratePhaseInfo `json:"MigratePhaseInfoList,omitempty" name:"MigratePhaseInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperMigratePhaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperMigratePhaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallCloudNativeAPIGatewayPluginRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 插件名称

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 插件版本号

	PluginVersion *string `json:"PluginVersion,omitempty" name:"PluginVersion"`
}

func (r *InstallCloudNativeAPIGatewayPluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallCloudNativeAPIGatewayPluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncOperateAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncOperateAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncOperateAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigFileRequest struct {
	*tchttp.BaseRequest

	// TSE实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeConfigFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolarisCLSTopicRequest struct {
	*tchttp.BaseRequest

	// 北极星引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePolarisCLSTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolarisCLSTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNginxIngressGlobalConfigRequest struct {
	*tchttp.BaseRequest

	// 实例版本号

	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`
}

func (r *GetNginxIngressGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNginxIngressGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceAliasesRequest struct {
	*tchttp.BaseRequest

	// 服务别名所指向的服务名。

	Service *string `json:"Service,omitempty" name:"Service"`
	// 服务别名所指向的命名空间名。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务别名。

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 服务别名命名空间。

	AliasNamespace *string `json:"AliasNamespace,omitempty" name:"AliasNamespace"`
	// 服务别名描述。

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGovernanceAliasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceAliasesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolarisAuthStatusRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// true开 false关

	PolarisAuth *bool `json:"PolarisAuth,omitempty" name:"PolarisAuth"`
}

func (r *ModifyPolarisAuthStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolarisAuthStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperateApolloInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperateApolloInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperateApolloInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsulServerInterfacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口总个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 接口列表

		Content []*NacosServerInterface `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsulServerInterfacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulServerInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayStrategyAutoScalerConfig struct {

	// 最大副本数

	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`
	// 指标列表

	Metrics []*CloudNativeAPIGatewayStrategyAutoScalerConfigMetric `json:"Metrics,omitempty" name:"Metrics"`
}

type UpdateCloudNativeAPIGatewayResult struct {

	// 云原生API网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生网关状态。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyNacosAdminPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNacosAdminPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNacosAdminPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObtainApolloAdminTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// apollo管理员token

		AdminToken *string `json:"AdminToken,omitempty" name:"AdminToken"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ObtainApolloAdminTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ObtainApolloAdminTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewaySecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewaySecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewaySecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGovernanceNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGovernanceNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateEngineSpecRequest struct {
	*tchttp.BaseRequest

	// 要删除的引擎规格id

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 主控地域

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DeleteOperateEngineSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateEngineSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateEngineSpecRequest struct {
	*tchttp.BaseRequest

	// 引擎规格id

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 引擎规格详情，使用 json 存储。

	SpecInfo *string `json:"SpecInfo,omitempty" name:"SpecInfo"`
}

func (r *ModifyOperateEngineSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateEngineSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetGovernanceUserTokenRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户ID

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *ResetGovernanceUserTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetGovernanceUserTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddGovernedK8SResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行任务的任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddGovernedK8SResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGovernedK8SResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEngineSyncJobRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据同步任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 数据同步任务运行模式(Immediate：立即运行)

	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`
	// 数据同步任务接入类型(extranet：公网类型接入)

	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`
	// 源端数据库IP

	SrcDbIp *string `json:"SrcDbIp,omitempty" name:"SrcDbIp"`
	// 源端数据库Port

	SrcDbPort *int64 `json:"SrcDbPort,omitempty" name:"SrcDbPort"`
	// 源端数据库所在地域

	SrcDbRegion *string `json:"SrcDbRegion,omitempty" name:"SrcDbRegion"`
	// 源端数据库类型

	SrcDbType *string `json:"SrcDbType,omitempty" name:"SrcDbType"`
	// 源端数据库名称

	SrcDbName *string `json:"SrcDbName,omitempty" name:"SrcDbName"`
	// 源端数据库用户名

	SrcDbUser *string `json:"SrcDbUser,omitempty" name:"SrcDbUser"`
	// 源端数据库密码

	SrcDbPassword *string `json:"SrcDbPassword,omitempty" name:"SrcDbPassword"`
	// 数据库名称映射关系

	DbMappings []*KVPair `json:"DbMappings,omitempty" name:"DbMappings"`
}

func (r *CreateEngineSyncJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEngineSyncJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNacosConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeNacosConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceUserTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户

		User *User `json:"User,omitempty" name:"User"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceUserTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceUserTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateEventListRequest struct {
	*tchttp.BaseRequest

	// 模块类型：engine ｜ console

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
}

func (r *DescribeOperateEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceAliasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceAliasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGovernanceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGovernanceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceStrategiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGovernanceStrategiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayConfigRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 分组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeCloudNativeAPIGatewayConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteCloudNativeAPIGatewayServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MergeZookeeperMigrateClusterRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *MergeZookeeperMigrateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MergeZookeeperMigrateClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExecTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务描述列表

		List []*OptTaskInfo `json:"List,omitempty" name:"List"`
		// 总个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExecTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExecTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置中心类型（consul、zookeeper、apollo等）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 配置项的节点路径key

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *DescribeZookeeperConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigFileGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigFileGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigFileGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNativeGatewayServiceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNativeGatewayServiceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNativeGatewayServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayOriginalPluginsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生网关插件信息响应结果。

		Result *ListCloudNativeAPIGatewaySystemPluginInfoResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayOriginalPluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayOriginalPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupRelation struct {

	// 用户组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 用户数组

	Users []*User `json:"Users,omitempty" name:"Users"`
}

type ModifyAuthStrategy struct {

	// 策略ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 新增关联的用户、用户组信息

	AddPrincipals *Principals `json:"AddPrincipals,omitempty" name:"AddPrincipals"`
	// 移除关联的用户、用户组信息

	RemovePrincipals *Principals `json:"RemovePrincipals,omitempty" name:"RemovePrincipals"`
	// 新增关联的资源信息

	AddResources *StrategyResource `json:"AddResources,omitempty" name:"AddResources"`
	// 移除关联的资源信息

	RemoveResources *StrategyResource `json:"RemoveResources,omitempty" name:"RemoveResources"`
	// 鉴权策略动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 简单描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 腾讯云主账户ID

	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

type ModifyGovernanceGroupTokenRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户组Token信息

	Group *UserGroup `json:"Group,omitempty" name:"Group"`
}

func (r *ModifyGovernanceGroupTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceGroupTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EurekaConfigInfo struct {

	// 配置项Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Gossip加密

	Title *string `json:"Title,omitempty" name:"Title"`
	// 配置项内容

	Value *string `json:"Value,omitempty" name:"Value"`
	// 配置描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 配置说明文档

	DocLink *string `json:"DocLink,omitempty" name:"DocLink"`
	// 配置变更提示

	ChangeAlert *string `json:"ChangeAlert,omitempty" name:"ChangeAlert"`
	// 配置项说明列表

	DescList []*EurekaConfigDesc `json:"DescList,omitempty" name:"DescList"`
}

type DeleteCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 是否删除实例关联的 CLS 日志主题。

	DeleteClsTopic *bool `json:"DeleteClsTopic,omitempty" name:"DeleteClsTopic"`
}

func (r *DeleteCloudNativeAPIGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNativeGatewayServiceGroupRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 命名空间ID，Consul 的可以不填写

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 服务名称

	Service *string `json:"Service,omitempty" name:"Service"`
}

func (r *DescribeNativeGatewayServiceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNativeGatewayServiceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServiceInstanceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperServiceInstanceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServiceInstanceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEurekaService struct {

	// 注册服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 正常服务实例个数

	NormalServiceInstanceCount *int64 `json:"NormalServiceInstanceCount,omitempty" name:"NormalServiceInstanceCount"`
	// 所有实例个数

	TotalServiceInstanceCount *int64 `json:"TotalServiceInstanceCount,omitempty" name:"TotalServiceInstanceCount"`
	// 服务状态

	ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`
}

type DescribeZookeeperMigratePhaseInfo struct {

	// 迁移阶段名称，分为三种阶段：ZookeeperJobPhaseMigratePrepare、ZookeeperJobPhaseMigrateImportData、ZookeeperJobPhaseMigrateMerge

	Name *string `json:"Name,omitempty" name:"Name"`
	// 迁移阶段状态，分为：waiting，doing，success

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeCloudNativeAPIGatewaySimpleInfosRequest struct {
	*tchttp.BaseRequest

	// 列表个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 列表偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 云原生网关ID列表

	GatewayIds []*string `json:"GatewayIds,omitempty" name:"GatewayIds"`
	// 云原生网关类型

	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`
	// 云原生网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DescribeCloudNativeAPIGatewaySimpleInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySimpleInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKongIngressControllerK8sVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeKongIngressControllerK8sVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKongIngressControllerK8sVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKongIngressControllerK8sVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Kong Ingress Controller 与 K8s 版本映射表

		Versions []*KongIngressControllerK8sVersion `json:"Versions,omitempty" name:"Versions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKongIngressControllerK8sVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKongIngressControllerK8sVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateWhiteListRequest struct {
	*tchttp.BaseRequest

	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南
	// - NorthAmerica，北美

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineQuota struct {

	// 引擎所在区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 剩余额度

	RemainQuota *int64 `json:"RemainQuota,omitempty" name:"RemainQuota"`
	// 大区

	Area *string `json:"Area,omitempty" name:"Area"`
	// 区域中文名称

	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
}

type ModifySREInstanceSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySREInstanceSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySREInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetGovernanceGroupTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetGovernanceGroupTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetGovernanceGroupTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GovernanceService struct {

	// 服务名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间名称。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 元数据信息数组。

	Metadatas []*Metadata `json:"Metadatas,omitempty" name:"Metadatas"`
	// 描述信息。

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 服务所属部门。

	Department *string `json:"Department,omitempty" name:"Department"`
	// 服务所属业务。

	Business *string `json:"Business,omitempty" name:"Business"`
	// 健康服务实例数

	HealthyInstanceCount *uint64 `json:"HealthyInstanceCount,omitempty" name:"HealthyInstanceCount"`
	// 服务实例总数

	TotalInstanceCount *uint64 `json:"TotalInstanceCount,omitempty" name:"TotalInstanceCount"`
	// 服务ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否可以编辑

	Editable *bool `json:"Editable,omitempty" name:"Editable"`
}

type NativeGatewayServerGroups struct {

	// 总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 分组信息数组。

	GatewayGroupList []*NativeGatewayServerGroup `json:"GatewayGroupList,omitempty" name:"GatewayGroupList"`
}

type OperateServiceItem struct {

	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 集群内IP

	ClusterIP *string `json:"ClusterIP,omitempty" name:"ClusterIP"`
	// CLB的id

	ClbId *string `json:"ClbId,omitempty" name:"ClbId"`
	// 服务IP

	ServiceIP *string `json:"ServiceIP,omitempty" name:"ServiceIP"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeGovernanceUsersRequest struct {
	*tchttp.BaseRequest

	// 用户名称，模糊搜索最后加上 * 字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 账户来源，QCloud | Polaris

	Source *string `json:"Source,omitempty" name:"Source"`
	// tse实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeGovernanceUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateInstanceLogsRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// pod名

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 日志行数

	Line *int64 `json:"Line,omitempty" name:"Line"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateInstanceLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateInstanceLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceStatusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 注册服务状态列表

		ServiceStatusList []*ServiceStatus `json:"ServiceStatusList,omitempty" name:"ServiceStatusList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceStatusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceNamespacesRequest struct {
	*tchttp.BaseRequest

	// 根据命名空间名称过滤。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// tse实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGovernanceNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZooKeeperDefaultConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统参数列表

		ZooKeeperDefaultConfigInfoList []*ZooKeeperDefaultConfigInfo `json:"ZooKeeperDefaultConfigInfoList,omitempty" name:"ZooKeeperDefaultConfigInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZooKeeperDefaultConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZooKeeperDefaultConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayTaskPhaseName struct {

	// 执行阶段

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 中文名称

	ChineseName *string `json:"ChineseName,omitempty" name:"ChineseName"`
}

type OptTaskInfo struct {

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务描述

	TaskComment *string `json:"TaskComment,omitempty" name:"TaskComment"`
	// 任务执行的命令

	TaskCommand *string `json:"TaskCommand,omitempty" name:"TaskCommand"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeleteOperateEksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperateEksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateEksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineIntranetAccessAddressListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络打通地址列表

		Content []*IntranetAccessAddress `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEngineIntranetAccessAddressListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineIntranetAccessAddressListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApolloEnvReplica struct {

	// 所属环境

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
	// 节点信息

	Replicas []*ApolloReplicaPodDetail `json:"Replicas,omitempty" name:"Replicas"`
}

type ListCloudNativeAPIGatewayRoutesRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 列表数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 列表 offset，若值可转成数字，该值代表列表索引，返回值包含总数；若值为非数字的字符串，则代表游标模式，返回值下不含总数；

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 服务的名字，精确匹配

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 路由的名字，精确匹配

	RouteName *string `json:"RouteName,omitempty" name:"RouteName"`
	// 过滤条件，多个过滤条件之间是与的关系，支持 name, path, host, method, service, protocol

	Filters []*ListFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListCloudNativeAPIGatewayRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenPrometheusRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// Prometheus集群的ID

	PromId *string `json:"PromId,omitempty" name:"PromId"`
}

func (r *OpenPrometheusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenPrometheusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNacosDefaultConfigInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNacosDefaultConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosDefaultConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigFilesByGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置文件列表

		ConfigFiles []*ConfigFile `json:"ConfigFiles,omitempty" name:"ConfigFiles"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigFilesByGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFilesByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id数组

	InstanceIDList []*string `json:"InstanceIDList,omitempty" name:"InstanceIDList"`
	// 实例tag数组

	TagList []*KVPair `json:"TagList,omitempty" name:"TagList"`
	// 主控地域

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *ModifyOperateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateSecurityGroupItem struct {

	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 安全组地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 安全组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 安全组状态

	GroupStatus *string `json:"GroupStatus,omitempty" name:"GroupStatus"`
	// 同步操作状态

	OperateStatus *string `json:"OperateStatus,omitempty" name:"OperateStatus"`
	// 开启端口

	OpenPorts *string `json:"OpenPorts,omitempty" name:"OpenPorts"`
	// 安全组版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ModifyGovernanceUserTokenRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户Token信息

	User *User `json:"User,omitempty" name:"User"`
}

func (r *ModifyGovernanceUserTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceUserTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetGovernanceUserTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetGovernanceUserTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetGovernanceUserTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayResult struct {

	// 云原生API网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关状态。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 云原生API网关所属可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 云原生API网关名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云原生API网关类型。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 云原生API网关版本。

	GatewayVersion *string `json:"GatewayVersion,omitempty" name:"GatewayVersion"`
	// 云原生API网关节点信息。

	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitempty" name:"NodeConfig"`
	// 云原生API网关vpc配置。

	VpcConfig *CloudNativeAPIGatewayVpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`
	// 云原生API网关描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 云原生API网关创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例的标签信息

	Tags []*InstanceTagInfo `json:"Tags,omitempty" name:"Tags"`
	// 是否开启 cls 日志

	EnableCls *bool `json:"EnableCls,omitempty" name:"EnableCls"`
	// 付费模式，0表示后付费，1预付费

	TradeType *int64 `json:"TradeType,omitempty" name:"TradeType"`
	// 实例版本，当前支持开发版和标准版【TRIAL、STANDARD】

	FeatureVersion *string `json:"FeatureVersion,omitempty" name:"FeatureVersion"`
	// 公网出流量带宽，[1,2048]Mbps

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 自动续费标记，0表示默认状态(用户未设置，即初始状态)；
	// 1表示自动续费，2表示明确不自动续费(用户设置)，若业务无续费概念或无需自动续费，需要设置为0

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 到期时间，预付费时使用

	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
	// 隔离时间，实例隔离时使用

	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
	// 是否开启客户端公网。

	EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`
	// 实例实际的地域信息

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// Ingress class名称

	IngressClassName *string `json:"IngressClassName,omitempty" name:"IngressClassName"`
}

type ExecInstanceInfo struct {

	// 实例id列表

	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

type DeleteCloudNativeAPIGatewayPublicNetworkRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 分组id。

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// ip 地址版本。

	InternetAddressVersion *string `json:"InternetAddressVersion,omitempty" name:"InternetAddressVersion"`
	// 公网ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *DeleteCloudNativeAPIGatewayPublicNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayPublicNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNacosReplicasRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 副本列表Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 副本列表Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNacosReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosReplicasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZookeeperMigrateLoadDataStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据导入的当前状态，有initiate、success、fail、doing四种状态。

		Status *string `json:"Status,omitempty" name:"Status"`
		// 状态信息

		StatusMessage *string `json:"StatusMessage,omitempty" name:"StatusMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetZookeeperMigrateLoadDataStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZookeeperMigrateLoadDataStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZookeeperMigrateMetaClusterInfo struct {

	// 源ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 客户端端口

	ClientPort *int64 `json:"ClientPort,omitempty" name:"ClientPort"`
	// 选举端口

	ElectionPort *int64 `json:"ElectionPort,omitempty" name:"ElectionPort"`
	// 集群端口

	ClusterPort *int64 `json:"ClusterPort,omitempty" name:"ClusterPort"`
}

type CheckUserCloudProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true为具备，false为不具备

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckUserCloudProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckUserCloudProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsulConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeConsulConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNativeGatewayServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 服务列表

		List []*NativeGatewayServiceItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNativeGatewayServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNativeGatewayServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNginxIngressConfigMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ConfigMap 列表

		ConfigMaps []*NginxIngressConfigMap `json:"ConfigMaps,omitempty" name:"ConfigMaps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNginxIngressConfigMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNginxIngressConfigMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageVpcEndPointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageVpcEndPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageVpcEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNativeGatewayServerGroupRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关名字, 最多支持60个字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云原生API网关描述信息, 最多支持120个字符。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 网关分组 id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ModifyNativeGatewayServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNativeGatewayServerGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayNodesResult struct {

	// 获取云原生API网关节点列表响应结果。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 云原生API网关节点列表。

	NodeList []*CloudNativeAPIGatewayNode `json:"NodeList,omitempty" name:"NodeList"`
}

type DeleteEngineRequest struct {
	*tchttp.BaseRequest

	// 引擎实例 ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteEngineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEngineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewaySystemParametersRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DescribeCloudNativeAPIGatewaySystemParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySystemParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrometheusRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DescribePrometheusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrometheusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZookeeperMigrateLoadDataStatusRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetZookeeperMigrateLoadDataStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZookeeperMigrateLoadDataStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayClsConfigResult struct {

	// 云原生网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// CLS配置信息列表。

	ClsConfigList []*CloudNativeAPIGatewayClsConfig `json:"ClsConfigList,omitempty" name:"ClsConfigList"`
}

type KongUpstreamPreview struct {

	// 服务名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 后端配置

	Target []*KongTarget `json:"Target,omitempty" name:"Target"`
}

type CheckEngineSyncJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 数据同步任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckEngineSyncJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckEngineSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayNodesRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 实例分组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 翻页获取多少个

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 翻页从第几个开始获取

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCloudNativeAPIGatewayNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsulServiceInstancesRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 注册服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeConsulServiceInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulServiceInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServerInterfacesRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回的列表个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回的列表起始偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeZookeeperServerInterfacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServerInterfacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSREInstanceLimitRequest struct {
	*tchttp.BaseRequest

	// 检查类型（分为注册中心和配置中心等类型）

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
}

func (r *CheckSREInstanceLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSREInstanceLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OptTaskBatchExeInfo struct {

	// 序列号

	Seq *int64 `json:"Seq,omitempty" name:"Seq"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 批次索引id

	BatchId *int64 `json:"BatchId,omitempty" name:"BatchId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 执行的状态，0（创建）、1（运行）、2（完成）、3（执行错误）、4（执行失败）

	ExeStatus *int64 `json:"ExeStatus,omitempty" name:"ExeStatus"`
	// 执行描述

	ExeMessage *string `json:"ExeMessage,omitempty" name:"ExeMessage"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeleteCloudNativeAPIGatewaySourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewaySourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewaySourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceAuthStatusRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGovernanceAuthStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceAuthStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务记录

		Content []*DescribeZookeeperService `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNativeGatewayServiceGroupRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 命名空间

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 服务列表

	Service *string `json:"Service,omitempty" name:"Service"`
	// 实例分组名称

	Group *string `json:"Group,omitempty" name:"Group"`
	// 实例标签

	Labels []*InstanceLabel `json:"Labels,omitempty" name:"Labels"`
}

func (r *ModifyNativeGatewayServiceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNativeGatewayServiceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoScalerResourceStrategyRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 策略ID

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DeleteAutoScalerResourceStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoScalerResourceStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateApolloDeploymentRequest struct {
	*tchttp.BaseRequest

	// 需要缩容的引擎实例ID集

	InstanceIDList []*string `json:"InstanceIDList,omitempty" name:"InstanceIDList"`
}

func (r *DeleteOperateApolloDeploymentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateApolloDeploymentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayLatestTaskPhasesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务执行阶段列表

		Phases []*CloudNativeAPIGatewayTaskPhase `json:"Phases,omitempty" name:"Phases"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayLatestTaskPhasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayLatestTaskPhasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEngineSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEngineSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEngineSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FinishMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FinishMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FinishMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNacosConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功更新nacos配置

		UpdateResult *bool `json:"UpdateResult,omitempty" name:"UpdateResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNacosConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNacosConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateSecurityGroupListRequest struct {
	*tchttp.BaseRequest

	// 查询条件: region | uin

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateSecurityGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateSecurityGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckZookeeperMigrateTargetStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 执行结果解释信息

		ResultMessage *string `json:"ResultMessage,omitempty" name:"ResultMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckZookeeperMigrateTargetStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckZookeeperMigrateTargetStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZoneRequest struct {
	*tchttp.BaseRequest

	// 部署地域

	DeployRegion *string `json:"DeployRegion,omitempty" name:"DeployRegion"`
	// 引擎类型

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
}

func (r *DescribeAvailableZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关实例列表响应结果。

		Result *ListCloudNativeAPIGatewayResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GovernanceInstanceInput struct {

	// 实例所在服务名。

	Service *string `json:"Service,omitempty" name:"Service"`
	// 实例服务所在命名空间。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 实例负载均衡权重信息。不填默认为100。

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 实例默认健康信息。不填默认为健康。

	Healthy *bool `json:"Healthy,omitempty" name:"Healthy"`
	// 实例隔离信息。不填默认为非隔离。

	Isolate *bool `json:"Isolate,omitempty" name:"Isolate"`
	// 实例ip。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 实例监听端口。

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 实例使用协议。不填默认为空。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 实例版本。不填默认为空。

	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`
	// 是否启用健康检查。不填默认不启用。

	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`
	// 上报心跳时间间隔。若 EnableHealthCheck 为不启用，则此参数不生效；若 EnableHealthCheck 启用，此参数不填，则默认 ttl 为 5s。

	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`
	// 元数据信息。

	Metadatas []*Metadata `json:"Metadatas,omitempty" name:"Metadatas"`
}

type OperateEngineSpec struct {

	// 引擎规格id

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 引擎规格信息

	SpecInfo *string `json:"SpecInfo,omitempty" name:"SpecInfo"`
}

type OperateWorkloadItem struct {

	// 负载名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 负载类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ZooKeeperConfigInfo struct {

	// 参数名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 当前运行参数值

	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`
	// 配置的类型说明

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 配置类型的限制描述

	ConfigTypeLimit *string `json:"ConfigTypeLimit,omitempty" name:"ConfigTypeLimit"`
	// 默认运行参数值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 参数描述说明

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// true表示只读，不支持修改

	ReadOnly *string `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 对应的分组:zoo、sasl、env

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 依赖的配置项，需要同时满足

	DependencyConfigs *string `json:"DependencyConfigs,omitempty" name:"DependencyConfigs"`
	// 分组中文显示名称

	GroupNameDesc *string `json:"GroupNameDesc,omitempty" name:"GroupNameDesc"`
}

type CreateNativeGatewayServiceSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNativeGatewayServiceSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNativeGatewayServiceSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigFileTag struct {

	// key-value 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// key-value 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ConfigFileGroupInput struct {

	// 配置文件组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 文件数

	FileCount *uint64 `json:"FileCount,omitempty" name:"FileCount"`
	// 关联用户，link_users

	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
	// 组id，link_groups

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// remove_link_users

	RemoveUserIds []*string `json:"RemoveUserIds,omitempty" name:"RemoveUserIds"`
	// remove_link_groups

	RemoveGroupIds []*string `json:"RemoveGroupIds,omitempty" name:"RemoveGroupIds"`
}

type DescribeCloudNativeAPIGatewayPluginVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关插件版本信息列表响应结果。

		Result *ListCloudNativeAPIGatewayPluginVersionsResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayPluginVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayPluginVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREGlobalConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规格列表。

		Specs []*SRESpec `json:"Specs,omitempty" name:"Specs"`
		// 版本列表

		Editions []*SREVersion `json:"Editions,omitempty" name:"Editions"`
		// 数据存储方式

		StorageTypes []*StorageTypeInfo `json:"StorageTypes,omitempty" name:"StorageTypes"`
		// 计费方式

		PayModes []*PayModeInfo `json:"PayModes,omitempty" name:"PayModes"`
		// 收费标准

		Charges []*ChargeInfo `json:"Charges,omitempty" name:"Charges"`
		// 配置中心版本列表

		ConfigEditions []*SREVersion `json:"ConfigEditions,omitempty" name:"ConfigEditions"`
		// 配置中心规格列表

		ConfigSpecs []*ConfigSpec `json:"ConfigSpecs,omitempty" name:"ConfigSpecs"`
		// 配置中心环境信息列表

		EnvNameInfos []*EnvNameInfo `json:"EnvNameInfos,omitempty" name:"EnvNameInfos"`
		// 服务治理引擎的规格列表

		ServiceGovernanceSpecs []*ServiceGovernanceSpec `json:"ServiceGovernanceSpecs,omitempty" name:"ServiceGovernanceSpecs"`
		// 服务治理引擎的规格信息列表

		ServiceGovernanceEditions []*ServiceGovernanceEdition `json:"ServiceGovernanceEditions,omitempty" name:"ServiceGovernanceEditions"`
		// Nacos规格列表

		NacosSpecs []*SRESpec `json:"NacosSpecs,omitempty" name:"NacosSpecs"`
		// Nacos版本列表

		NacosEditions []*SREVersion `json:"NacosEditions,omitempty" name:"NacosEditions"`
		// Apollo规格列表

		ApolloSpecs []*ApolloSpec `json:"ApolloSpecs,omitempty" name:"ApolloSpecs"`
		// Apollo DB计费标签列表

		ApolloDBSvTags []*ApolloDBSvTag `json:"ApolloDBSvTags,omitempty" name:"ApolloDBSvTags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSREGlobalConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSREGlobalConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateEksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateEksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateEksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayUpstreamResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Result *KongUpstreamList `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayUpstreamResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除云原生API网关实例响应结果。

		Result *DeleteCloudNativeAPIGatewayResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewaySystemPluginInfoResult struct {

	// 总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 系统插件信息列表。

	PluginList []*SystemPluginInfo `json:"PluginList,omitempty" name:"PluginList"`
}

type ServiceGovernanceInfo struct {

	// 引擎所在的地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 服务治理引擎绑定的kubernetes集群信息

	BoundK8SInfos []*BoundK8SInfo `json:"BoundK8SInfos,omitempty" name:"BoundK8SInfos"`
	// 服务治理引擎绑定的网络信息

	VpcInfos []*VpcInfo `json:"VpcInfos,omitempty" name:"VpcInfos"`
	// 当前实例鉴权是否开启

	AuthOpen *bool `json:"AuthOpen,omitempty" name:"AuthOpen"`
	// 该实例支持的功能，鉴权就是 Auth

	Features []*string `json:"Features,omitempty" name:"Features"`
	// 主账户名默认为 polaris，该值为主账户的默认密码

	MainPassword *string `json:"MainPassword,omitempty" name:"MainPassword"`
	// 服务治理pushgateway引擎绑定的网络信息

	PgwVpcInfos []*VpcInfo `json:"PgwVpcInfos,omitempty" name:"PgwVpcInfos"`
	// 服务治理限流server引擎绑定的网络信息

	LimiterVpcInfos []*VpcInfo `json:"LimiterVpcInfos,omitempty" name:"LimiterVpcInfos"`
}

type ListCloudNativeAPIGatewayRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Result *KongRouteList `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewayRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Query struct {

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// tke metric GroupBy

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// tke metric MetricName

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// tke metric Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// tke metricQueryVersion

	QueryVersion *string `json:"QueryVersion,omitempty" name:"QueryVersion"`
	// tke metricStartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// tke metric Conditions

	Conditions []*Condition `json:"Conditions,omitempty" name:"Conditions"`
	// tke metric Datasource

	Datasource *string `json:"Datasource,omitempty" name:"Datasource"`
}

type SourceInstanceVpcInfo struct {

	// 微服务引擎VPC信息

	VpcID *string `json:"VpcID,omitempty" name:"VpcID"`
	// 微服务引擎子网信息

	SubnetID *string `json:"SubnetID,omitempty" name:"SubnetID"`
}

type CreateCloudNativeAPIGatewaySourceRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务来源ID

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 服务来源类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateCloudNativeAPIGatewaySourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewaySourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsulServiceInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务实例记录

		Content []*DescribeServiceInstance `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsulServiceInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulServiceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstanceAccessAddressRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 引擎其他组件名称（pushgateway、polaris-limiter）

	Workload *string `json:"Workload,omitempty" name:"Workload"`
	// 部署地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *DescribeSREInstanceAccessAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSREInstanceAccessAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthStrategy struct {

	// 策略唯一ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 涉及的用户 or 用户组

	Principals *Principals `json:"Principals,omitempty" name:"Principals"`
	// 资源操作权限

	Action *string `json:"Action,omitempty" name:"Action"`
	// 简单描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 主账户的UIN

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 策略关联的资源

	Resources *StrategyResource `json:"Resources,omitempty" name:"Resources"`
	// 该策略是否是默认策略

	Default *bool `json:"Default,omitempty" name:"Default"`
}

type Data struct {

	// 维度组合

	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 命名空间名

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标统计持续时间

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 指标数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeKongIngressRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DescribeKongIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKongIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateEngineSpecListRequest struct {
	*tchttp.BaseRequest

	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 便宜

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 总量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperateEngineSpecListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateEngineSpecListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTagInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否修改成功

		ModifiedResult *bool `json:"ModifiedResult,omitempty" name:"ModifiedResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceTagInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTagInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewaySourcesRequest struct {
	*tchttp.BaseRequest

	// 云原生网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DescribeCloudNativeAPIGatewaySourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开通了服务治理中心。

		GovernanceOpened *bool `json:"GovernanceOpened,omitempty" name:"GovernanceOpened"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServicePort struct {

	// 端口名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 端口号

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 接口分类key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 接口分类名

	Title *string `json:"Title,omitempty" name:"Title"`
}

type DescribeInstanceLogRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 引擎服务实例名

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 日志类型，可选值：stdout（标准输出），runtime（运行日志），monitor（审计日志），默认为运行日志

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 部署地域，默认值是部署主地域

	DeployRegion *string `json:"DeployRegion,omitempty" name:"DeployRegion"`
	// 起始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 终止时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 单页请求配置数量，取值范围[1, 200]，默认值为50

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 加载更多日志时使用，透传上次返回的Context值，获取后续的日志内容

	Context *string `json:"Context,omitempty" name:"Context"`
	// Apollo实例请求日志时，需要按照组件分类

	ApolloContainer *string `json:"ApolloContainer,omitempty" name:"ApolloContainer"`
}

func (r *DescribeInstanceLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateAgentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// agent分布数组

		Result []*OperateAgentVersionInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateAgentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZookeeperNodeDescription struct {

	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 创建的事务ID

	CreateZxid *string `json:"CreateZxid,omitempty" name:"CreateZxid"`
	// znode最近一次修改的事物ID

	ModifiedZxid *string `json:"ModifiedZxid,omitempty" name:"ModifiedZxid"`
	// 创建时间时间戳，单位毫秒

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间时间戳，单位毫秒

	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// znode的数据域变化次数

	ZNodeVersion *int64 `json:"ZNodeVersion,omitempty" name:"ZNodeVersion"`
	// znode的节点的变化次数

	ChildVersion *int64 `json:"ChildVersion,omitempty" name:"ChildVersion"`
	// znode的acl变化次数

	ACLVersion *int64 `json:"ACLVersion,omitempty" name:"ACLVersion"`
	// 临时节点的会话id，如果不是临时节点，则为0

	EphemeralOwner *int64 `json:"EphemeralOwner,omitempty" name:"EphemeralOwner"`
	// 数据长度

	DataLength *int64 `json:"DataLength,omitempty" name:"DataLength"`
	// 子节点个数

	NumChildren *int64 `json:"NumChildren,omitempty" name:"NumChildren"`
	// 最近修改的子节点

	PZxid *string `json:"PZxid,omitempty" name:"PZxid"`
	// 节点名称

	Key *string `json:"Key,omitempty" name:"Key"`
}

type CreateGovernanceNamespacesRequest struct {
	*tchttp.BaseRequest

	// tse 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间信息。

	GovernanceNamespaces []*GovernanceNamespaceInput `json:"GovernanceNamespaces,omitempty" name:"GovernanceNamespaces"`
}

func (r *CreateGovernanceNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateFeatureVersionListRequest struct {
	*tchttp.BaseRequest

	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 总量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperateFeatureVersionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateFeatureVersionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopEngineSyncJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopEngineSyncJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopEngineSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Label struct {

	// 标签键名称

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeCloudNativeAPIGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关基础信息响应结果。

		Result *DescribeCloudNativeAPIGatewayResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsulServerInterfacesRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回的列表个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回的列表起始偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConsulServerInterfacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulServerInterfacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceStrategiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 策略列表

		Content *AuthStrategy `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceStrategiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNacosDefaultConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统参数列表

		NacosDefaultConfigInfoList []*NacosDefaultConfigInfo `json:"NacosDefaultConfigInfoList,omitempty" name:"NacosDefaultConfigInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNacosDefaultConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosDefaultConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceInstancesRequest struct {
	*tchttp.BaseRequest

	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务实例信息。

	GovernanceInstances []*GovernanceInstanceUpdate `json:"GovernanceInstances,omitempty" name:"GovernanceInstances"`
}

func (r *ModifyGovernanceInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否创建成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGovernanceUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewaySystemPluginsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生网关系统插件信息响应结果。

		Result *ListCloudNativeAPIGatewaySystemPluginInfoResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewaySystemPluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySystemPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SREVersion struct {

	// 版本名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本号。

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 可升级版本列表

	UpgradeVersionList []*string `json:"UpgradeVersionList,omitempty" name:"UpgradeVersionList"`
}

type DeleteOperateWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperateWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNativeGatewayServiceSourcesRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例名称，模糊搜索

	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
	// 微服务引擎类型：TSE-Nacos｜TSE-Consul｜TSE-PolarisMesh｜Customer-Nacos｜Customer-Consul｜Customer-PolarisMesh

	SourceTypes []*string `json:"SourceTypes,omitempty" name:"SourceTypes"`
	// 排序字段类型，当前仅支持SourceName

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序类型，AES/DESC

	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
	// 单页条数，最大100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNativeGatewayServiceSourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNativeGatewayServiceSourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewaySecret struct {

	// secret所属的集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// secret所属的namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 证书名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 证书创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 证书过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// Secret原始的json对象

	Raw *string `json:"Raw,omitempty" name:"Raw"`
	// 是否包含根证书

	HasCaCrt *bool `json:"HasCaCrt,omitempty" name:"HasCaCrt"`
}

type OpenKongIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenKongIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKongIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateFeatureVersion struct {

	// STANDARD

	FeatureVersion *string `json:"FeatureVersion,omitempty" name:"FeatureVersion"`
	// 功能版本名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 可选副本数列表

	ReplicaList *string `json:"ReplicaList,omitempty" name:"ReplicaList"`
	// 可选引擎规格列表

	SpecList *string `json:"SpecList,omitempty" name:"SpecList"`
}

type PolarisReplica struct {

	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type ZookeeperReplica struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 可用区ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 别名

	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeCloudAlarmPolarisPodsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点信息

		Replicas []*PolarisReplica `json:"Replicas,omitempty" name:"Replicas"`
		// 节点总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudAlarmPolarisPodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudAlarmPolarisPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceUserPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceUserPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceUserPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayVersion struct {

	// 版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// 版本状态，OPEN为可用，CLOSE为不可用。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 可兼容的版本，使用于版本升级场景。

	CompatibleVersion []*string `json:"CompatibleVersion,omitempty" name:"CompatibleVersion"`
}

type ZookeeperMigrateMetaInfo struct {

	// 源集群元数据

	ClusterInfoList []*ZookeeperMigrateMetaClusterInfo `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`
	// 管理员用户名

	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`
	// 管理员密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 数据MD5

	DataMD5 *string `json:"DataMD5,omitempty" name:"DataMD5"`
	// 数据名称

	DataFileName *string `json:"DataFileName,omitempty" name:"DataFileName"`
}

type EnvNameInfo struct {

	// 环境名称

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
	// 环境中文描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type CreateGovernanceInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGovernanceInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NacosDefaultConfigInfo struct {

	// 参数名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 默认运行参数值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 参数描述说明

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// true表示只读，不支持修改

	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 配置的类型说明

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 配置类型的限制描述

	ConfigTypeLimit *string `json:"ConfigTypeLimit,omitempty" name:"ConfigTypeLimit"`
}

type PrincipalEntry struct {

	// 用户ID｜用户组ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 用户名｜用户组名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateGovernanceGroupRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户组

	Group *UserGroup `json:"Group,omitempty" name:"Group"`
}

func (r *CreateGovernanceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSREInstanceBasicInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *UpdateSREInstanceBasicInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSREInstanceBasicInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableConsoleAccessAddressRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// true开 false关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// internet公网  intranet内网

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 公网白名单

	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
	// apollo开启控制台内网时选择vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// apollo开启控制台内网时选择subnet

	SubnetIds *string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 公网带宽

	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`
	// 指定打通后的vip地址

	EndpointVip *string `json:"EndpointVip,omitempty" name:"EndpointVip"`
}

func (r *EnableConsoleAccessAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableConsoleAccessAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstanceHealthStatusRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例部署地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *DescribeSREInstanceHealthStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSREInstanceHealthStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishConfigFilesRequest struct {
	*tchttp.BaseRequest

	// TSE实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置文件发布

	ConfigFileReleases *ConfigFileRelease `json:"ConfigFileReleases,omitempty" name:"ConfigFileReleases"`
}

func (r *PublishConfigFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishConfigFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateRegion struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域中文

	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
	// 大区

	Area *string `json:"Area,omitempty" name:"Area"`
	// 排序

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// TSE镜像仓库地址

	TSEPublicRepoUrl *string `json:"TSEPublicRepoUrl,omitempty" name:"TSEPublicRepoUrl"`
	// agent镜像地址

	AgentImageUrl *string `json:"AgentImageUrl,omitempty" name:"AgentImageUrl"`
	// 巴拉多地址

	BaradHost *string `json:"BaradHost,omitempty" name:"BaradHost"`
	// EKS地址

	EksHost *string `json:"EksHost,omitempty" name:"EksHost"`
	// VPC服务地址

	InnerVpcHost *string `json:"InnerVpcHost,omitempty" name:"InnerVpcHost"`
	// 表示该地域的发布状态。有以下可选值：
	//
	// - PRE：表示发布中，外部客户不可见。内部账号可见，用于验证。
	// - PROD：表示已发布，外部客户可见。

	ReleaseStatus *string `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`
}

type DeleteCloudNativeAPIGatewayPublicNetworkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayPublicNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayPublicNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshOperateAgentRequest struct {
	*tchttp.BaseRequest

	// China

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *RefreshOperateAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshOperateAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayCosConfig struct {

	// Cos存储通名称。

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// Cos对象路径。

	ObjectName *string `json:"ObjectName,omitempty" name:"ObjectName"`
	// 地域。

	Region *string `json:"Region,omitempty" name:"Region"`
}

type ConfigFileRelease struct {

	// 配置文件发布id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 配置文件发布名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 配置文件发布命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 配置文件发布组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 配置文件发布文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 配置文件发布内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 配置文件发布注释

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 配置文件发布Md5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 配置文件发布版本

	Version *uint64 `json:"Version,omitempty" name:"Version"`
	// 配置文件发布创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 配置文件发布创建者

	CreateBy *string `json:"CreateBy,omitempty" name:"CreateBy"`
	// 配置文件发布修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 配置文件发布修改者

	ModifyBy *string `json:"ModifyBy,omitempty" name:"ModifyBy"`
}

type EurekaConfigDesc struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否支持编辑

	EditEnable *bool `json:"EditEnable,omitempty" name:"EditEnable"`
	// 配置项默认值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 可选值

	OptionValue *string `json:"OptionValue,omitempty" name:"OptionValue"`
	// 配置项提示

	Tips *string `json:"Tips,omitempty" name:"Tips"`
}

type KongRoutePreview struct {

	// 服务名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 无

	Methods []*string `json:"Methods,omitempty" name:"Methods"`
	// 无

	Paths []*string `json:"Paths,omitempty" name:"Paths"`
	// 无

	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
	// 无

	Protocols []*string `json:"Protocols,omitempty" name:"Protocols"`
	// 无

	PreserveHost *bool `json:"PreserveHost,omitempty" name:"PreserveHost"`
	// 无

	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitempty" name:"HttpsRedirectStatusCode"`
	// 无

	StripPath *bool `json:"StripPath,omitempty" name:"StripPath"`
	// 无

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 是否开启了强制HTTPS

	ForceHttps *bool `json:"ForceHttps,omitempty" name:"ForceHttps"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 目的端口

	DestinationPorts []*uint64 `json:"DestinationPorts,omitempty" name:"DestinationPorts"`
}

type DescribeEngineIntranetAccessAddressListRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 部署地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *DescribeEngineIntranetAccessAddressListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineIntranetAccessAddressListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicensesRequest struct {
	*tchttp.BaseRequest

	// 分页 offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 一页的条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLicensesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicensesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateSecurityGroupResult struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 开放的端口

	Ports *string `json:"Ports,omitempty" name:"Ports"`
	// 操作总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 成功数

	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
	// 失败数

	FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
}

type DescribeCloudNativeAPIGatewayClsConfigRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DescribeCloudNativeAPIGatewayClsConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayClsConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置中心类型（consul、zookeeper、apollo等）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 配置项的节点路径key

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *DescribeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineExportPortsRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeEngineExportPortsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineExportPortsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayCertificateRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 证书名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 证书私钥

	Key *string `json:"Key,omitempty" name:"Key"`
	// 证书pem格式

	Crt *string `json:"Crt,omitempty" name:"Crt"`
	// 绑定的域名

	BindDomains []*string `json:"BindDomains,omitempty" name:"BindDomains"`
	// ssl平台证书 Id

	CertId *string `json:"CertId,omitempty" name:"CertId"`
	// 证书id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyCloudNativeAPIGatewayCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceGovernanceSpec struct {

	// 规格ID

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 规格对应的cpu核数

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 规格对应的内存容量，单位MB

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 规格名称

	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`
	// 节点数

	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`
}

type DescribeConfigFileGroupsRequest struct {
	*tchttp.BaseRequest

	// tse实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 根据命名空间过滤

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 根据配置文件组名过滤

	Group *string `json:"Group,omitempty" name:"Group"`
	// 根据配置文件组名过滤

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConfigFileGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFileGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewaySystemParametersRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *ListCloudNativeAPIGatewaySystemParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewaySystemParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayBasicLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志内容

		Result *DescribeCloudNativeAPIGatewayBasicLogResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayBasicLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayBasicLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigFileReleaseRequest struct {
	*tchttp.BaseRequest

	// TSE实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 配置分组名称

	Group *string `json:"Group,omitempty" name:"Group"`
	// 配置文件名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeConfigFileReleaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFileReleaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSRELatestTaskPhasesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSRELatestTaskPhasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSRELatestTaskPhasesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayNetworkRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关网络配置列表

	NetworkConfigList []*NetworkConfig `json:"NetworkConfigList,omitempty" name:"NetworkConfigList"`
	// 公网出流量带宽[1,2048]Mbps

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 分组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 公网ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 负载均衡描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyCloudNativeAPIGatewayNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplyLicenseRequest struct {
	*tchttp.BaseRequest

	// 北极星集群ID

	ResourceID *string `json:"ResourceID,omitempty" name:"ResourceID"`
	// 合同ID

	ContractNo *string `json:"ContractNo,omitempty" name:"ContractNo"`
	// 客户uin

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 客户名称

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 销售方uin

	SellerUin *string `json:"SellerUin,omitempty" name:"SellerUin"`
	// 销售方公司名称

	SellerName *string `json:"SellerName,omitempty" name:"SellerName"`
	// 代理商uin

	AgencyUin *string `json:"AgencyUin,omitempty" name:"AgencyUin"`
	// 代理商名称

	AgencyName *string `json:"AgencyName,omitempty" name:"AgencyName"`
	// （北极星-独立版本｜北极星-TCS版本），前端传入的值为 tse_private_polaris_tcs 以及 tse_private_polaris_vm

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 控制类型，H ｜ S，H为硬限制，S 为软限制

	LicenseControlType *string `json:"LicenseControlType,omitempty" name:"LicenseControlType"`
	// License何时生效，时间戳，秒级别【重要】

	LicenseEffectiveTime *int64 `json:"LicenseEffectiveTime,omitempty" name:"LicenseEffectiveTime"`
	// license 的有效天数，天级别

	LicenseActiveDays *int64 `json:"LicenseActiveDays,omitempty" name:"LicenseActiveDays"`
	// License中对资源的约束表达式

	LicenseRuleList []*LicenseRuleInfo `json:"LicenseRuleList,omitempty" name:"LicenseRuleList"`
}

func (r *CreateApplyLicenseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplyLicenseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGovernanceStrategyRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 策略

	Strategy *AuthStrategy `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *CreateGovernanceStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateAgentVersionInfo struct {

	// 引擎类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// agent分布列表

	List []*OperateAgentVersion `json:"List,omitempty" name:"List"`
}

type ZookeeperServerInterface struct {

	// 接口名

	Interface *string `json:"Interface,omitempty" name:"Interface"`
}

type DescribeCloudNativeAPIGatewayVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云原生网关实例版本列表

		Result *DescribeCloudNativeAPIGatewayVersionResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分页总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置文件列表

		ConfigFiles []*ConfigFile `json:"ConfigFiles,omitempty" name:"ConfigFiles"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNginxIngressConfigMapRequest struct {
	*tchttp.BaseRequest

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// ConfigMap 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// json 格式的 base64编码

	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *ModifyNginxIngressConfigMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNginxIngressConfigMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayCertificateRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 证书Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteCloudNativeAPIGatewayCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayStrategyAutoScalerConfigMetric struct {

	// 指标类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 指标资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 指标目标类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 指标目标值

	TargetValue *int64 `json:"TargetValue,omitempty" name:"TargetValue"`
}

type DescribeAvailableZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用子网列表

		AvailableSubnet []*string `json:"AvailableSubnet,omitempty" name:"AvailableSubnet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApolloReplicasRequest struct {
	*tchttp.BaseRequest

	// 引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据大小限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 起始页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 是否需要获取admin类型的pods

	AdminPodFlag *bool `json:"AdminPodFlag,omitempty" name:"AdminPodFlag"`
}

func (r *DescribeApolloReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApolloReplicasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNginxIngressConfigMapRequest struct {
	*tchttp.BaseRequest

	// 网关实例Id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// ConfigMap名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *GetNginxIngressConfigMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNginxIngressConfigMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateAgentVersion struct {

	// 镜像地址

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 镜像标签

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 数量

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeConsulReplicasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 注册引擎实例副本信息

		Replicas []*ConsulReplica `json:"Replicas,omitempty" name:"Replicas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsulReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperMigratePhaseRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeZookeeperMigratePhaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperMigratePhaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewaySystemParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关系统参数列表响应结果。

		Result *ListCloudNativeAPIGatewaySystemParametersResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewaySystemParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewaySystemParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolarisServerInterfacesRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回的列表个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回的列表起始偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePolarisServerInterfacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolarisServerInterfacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewaySystemParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewaySystemParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewaySystemParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshOperateAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshOperateAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshOperateAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartZookeeperMigrateMergeClusterRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StartZookeeperMigrateMergeClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartZookeeperMigrateMergeClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEngineSyncJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 数据同步任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEngineSyncJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEngineSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorResult struct {

	// 监控类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 监控标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 监控数据

	DataPoints []*MonitorDataPoint `json:"DataPoints,omitempty" name:"DataPoints"`
}

type StrategyResourceEntry struct {

	// 资源Id，如果是全部的话，那么ID就是 *

	Id *string `json:"Id,omitempty" name:"Id"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务名｜配置分组名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type NginxIngressConfigMap struct {

	// configmap 文本内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// ConfigMap 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeCloudNativeAPIGatewayPluginsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 按照插件名称模糊匹配。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeCloudNativeAPIGatewayPluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayPluginsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenGovernanceRequest struct {
	*tchttp.BaseRequest
}

func (r *OpenGovernanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenGovernanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayPlugin struct {

	// 插件名称

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 插件状态, NotInstalled|Installing|Installed

	Status *string `json:"Status,omitempty" name:"Status"`
	// 插件版本

	PluginVersion *string `json:"PluginVersion,omitempty" name:"PluginVersion"`
	// 插件下载地址

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

type UpdateZooKeeperConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例系统参数列表

	ZooKeeperConfigInfoList []*ZooKeeperConfigInfo `json:"ZooKeeperConfigInfoList,omitempty" name:"ZooKeeperConfigInfoList"`
}

func (r *UpdateZooKeeperConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateZooKeeperConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayRouteRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 路由的ID或名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteCloudNativeAPIGatewayRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateAgentListRequest struct {
	*tchttp.BaseRequest

	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// China

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperateAgentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateAgentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayPluginsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关插件列表响应结果。

		Result *ListCloudNativeAPIGatewayPluginsResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewayPluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateEngineVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperateEngineVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateEngineVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {

	// 维度名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 维度值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeZookeeperServiceInstance struct {

	// 实例IP

	ServiceInstanceIp *string `json:"ServiceInstanceIp,omitempty" name:"ServiceInstanceIp"`
	// 实例端口

	ServiceInstancePort *int64 `json:"ServiceInstancePort,omitempty" name:"ServiceInstancePort"`
}

type DescribeNacosServerInterfacesRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回的列表个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回的列表起始偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNacosServerInterfacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosServerInterfacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolarisBindK8sClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定的Kubernetes集群列表

		KubernetesInfos []*KubernetesInfo `json:"KubernetesInfos,omitempty" name:"KubernetesInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolarisBindK8sClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolarisBindK8sClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayCertificatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Result *KongCertificatesList `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewayCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayCanaryRuleRequest struct {
	*tchttp.BaseRequest

	// 网关 ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务 ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 优先级，同一个服务的灰度规则优先级是唯一的

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 灰度规则配置

	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitempty" name:"CanaryRule"`
}

func (r *ModifyCloudNativeAPIGatewayCanaryRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayCanaryRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayStrategyBindingGroupInfoResult struct {

	// 数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 云原生API网关实例策略绑定网关分组列表

	GroupInfos []*CloudNativeAPIGatewayStrategyBindingGroupInfo `json:"GroupInfos,omitempty" name:"GroupInfos"`
}

type DeleteOperateFeatureVersionRequest struct {
	*tchttp.BaseRequest

	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 要删除的功能版本

	FeatureVersion *string `json:"FeatureVersion,omitempty" name:"FeatureVersion"`
	// 要删除的功能版本引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteOperateFeatureVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateFeatureVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNginxIngressNginxRawConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// nginx.conf 文本内容

		NginxConf *string `json:"NginxConf,omitempty" name:"NginxConf"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNginxIngressNginxRawConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNginxIngressNginxRawConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableIntranetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableIntranetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableIntranetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceGroupRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户组数组

	Groups []*ModifyUserGroup `json:"Groups,omitempty" name:"Groups"`
}

func (r *ModifyGovernanceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayCanaryRule struct {

	// 优先级，值范围为 0 到 100；值越大，优先级越高；不同规则间优先级不可重复

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 是否启用规则

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 参数匹配条件

	ConditionList []*CloudNativeAPIGatewayCanaryRuleCondition `json:"ConditionList,omitempty" name:"ConditionList"`
	// 服务的流量百分比配置

	BalancedServiceList []*CloudNativeAPIGatewayBalancedService `json:"BalancedServiceList,omitempty" name:"BalancedServiceList"`
}

type CreateCloudNativeAPIGatewayResult struct {

	// 云原生API网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生网关状态。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type CreateOrModifyCloudNativeAPIGatewayIngressRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// k8s集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 无

	IngressData *string `json:"IngressData,omitempty" name:"IngressData"`
}

func (r *CreateOrModifyCloudNativeAPIGatewayIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrModifyCloudNativeAPIGatewayIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnginePodAccessAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点信息

		EnginePodInfoList []*EnginePodInfo `json:"EnginePodInfoList,omitempty" name:"EnginePodInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnginePodAccessAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnginePodAccessAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayStrategy struct {

	// 策略ID

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 弹性伸缩配置

	Config *CloudNativeAPIGatewayStrategyAutoScalerConfig `json:"Config,omitempty" name:"Config"`
	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

type DeleteCloudNativeAPIGatewayServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEngineSyncJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEngineSyncJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEngineSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateInstanceUpgradeRequest struct {
	*tchttp.BaseRequest

	// 引擎id数组

	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
	// 升级辅助内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *ModifyOperateInstanceUpgradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateInstanceUpgradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateSpecListRequest struct {
	*tchttp.BaseRequest

	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperateSpecListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateSpecListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigFileGroupRequest struct {
	*tchttp.BaseRequest

	// tse 实例 id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置文件组实体

	ConfigFileGroup *ConfigFileGroup `json:"ConfigFileGroup,omitempty" name:"ConfigFileGroup"`
}

func (r *CreateConfigFileGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigFileGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGovernanceGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayPublicNetworkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayPublicNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayPublicNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceDiscoverEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务实例事件列表

		Content []*GovernanceDiscoverEvent `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceDiscoverEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceDiscoverEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateEksListRequest struct {
	*tchttp.BaseRequest

	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 集群类型:engine | console

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// 主控地域参数，支持的值如下。（QueryType 为 console 时可不填）
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateEksListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateEksListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayDomainCluster struct {

	// tke集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// tke Namespace信息

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type UnbindAutoScalerResourceStrategyFromGroupsRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 策略ID

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 网关分组ID列表

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

func (r *UnbindAutoScalerResourceStrategyFromGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoScalerResourceStrategyFromGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableInternetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableInternetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableInternetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEngineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务 ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 引擎实例 ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEngineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceGroupsRequest struct {
	*tchttp.BaseRequest

	// 用户名称，模糊搜索最后加上 * 字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 账户来源，QCloud | Polaris

	Source *string `json:"Source,omitempty" name:"Source"`
	// tse实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户ID，填写用户ID时为查询该用户下的所有group

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeGovernanceGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateEngineSpecRequest struct {
	*tchttp.BaseRequest

	// 引擎规格id

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 引擎规格详情，使用 json 存储。

	SpecInfo *string `json:"SpecInfo,omitempty" name:"SpecInfo"`
	// 无

	IsModify *string `json:"IsModify,omitempty" name:"IsModify"`
}

func (r *CreateOperateEngineSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateEngineSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateEngineVersionRequest struct {
	*tchttp.BaseRequest

	// 引擎版本号

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DeleteOperateEngineVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateEngineVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZookeeperMigrateUploadMetaInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 临时密钥 Id，可用于计算签名

		TmpSecretID *string `json:"TmpSecretID,omitempty" name:"TmpSecretID"`
		// 临时密钥 Key，可用于计算签名

		TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
		// 请求时需要用的 token 字符串

		SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
		// cos对应的aapid

		TmpAppId *string `json:"TmpAppId,omitempty" name:"TmpAppId"`
		// 密钥的失效时间，是 UNIX 时间戳

		ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// cos地址

		Domain *string `json:"Domain,omitempty" name:"Domain"`
		// cas存储的文件名

		SaveKey *string `json:"SaveKey,omitempty" name:"SaveKey"`
		// 桶的名称

		BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetZookeeperMigrateUploadMetaInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZookeeperMigrateUploadMetaInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateEngineSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperateEngineSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateEngineSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableConsoleAccessAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableConsoleAccessAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableConsoleAccessAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayVersionResult struct {

	// 云原生网关实例版本列表

	VersionList []*CloudNativeAPIGatewayVersion `json:"VersionList,omitempty" name:"VersionList"`
}

type DeleteNativeGatewayServiceSourceRequest struct {
	*tchttp.BaseRequest

	// 网关实例 ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例 ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
}

func (r *DeleteNativeGatewayServiceSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNativeGatewayServiceSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsulServicesRequest struct {
	*tchttp.BaseRequest

	// 请求过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 注册引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeConsulServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceGroupDetailRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// tse实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGovernanceGroupDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceGroupDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperateWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南
	// - NorthAmerica，北美

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DeleteOperateWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEurekaConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeEurekaConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEurekaConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayCertificatesDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Result *KongCertificatesDomainList `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewayCertificatesDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayCertificatesDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateEksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperateEksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateEksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserGroup struct {

	// 用户组ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 简单描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 添加的用户ID列表

	AddRelation *GroupRelation `json:"AddRelation,omitempty" name:"AddRelation"`
	// 移除的用户ID列表

	RemoveRelation *GroupRelation `json:"RemoveRelation,omitempty" name:"RemoveRelation"`
}

type ManageZookeeperConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置中心类型（consul、zookeeper、apollo等）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 请求命名 PUT GET DELETE

	Command *string `json:"Command,omitempty" name:"Command"`
	// 配置的Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 配置的Value

	Value *string `json:"Value,omitempty" name:"Value"`
	// ZK配置的选项

	ZKConfigOption *ZNodeOption `json:"ZKConfigOption,omitempty" name:"ZKConfigOption"`
}

func (r *ManageZookeeperConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageZookeeperConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayResult struct {

	// 云原生网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生网关状态。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type Condition struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 查询维度

	Dimension []*string `json:"Dimension,omitempty" name:"Dimension"`
}

type CosTokenInfo struct {

	// 桶的名称

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// 请求时需要用的 token 字符串

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// cos对应的aapid

	TempAppId *string `json:"TempAppId,omitempty" name:"TempAppId"`
	// 临时密钥 Id，可用于计算签名

	TempSecretID *string `json:"TempSecretID,omitempty" name:"TempSecretID"`
	// 临时密钥 Key，可用于计算签名

	TempSecretKey *string `json:"TempSecretKey,omitempty" name:"TempSecretKey"`
	// 密钥的失效时间，是 UNIX 时间戳

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// cos地址

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// cos存储的文件名

	SaveKey *string `json:"SaveKey,omitempty" name:"SaveKey"`
	// cos地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type SRESpec struct {

	// 规格ID

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 规格名称，例如：标准型

	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`
	// CPU核数

	CpuNum *uint64 `json:"CpuNum,omitempty" name:"CpuNum"`
	// 内存大小,单位MB

	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`
}

type DescribeCloudNativeAPIGatewaySystemParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关系统参数列表响应结果。

		Result *ListCloudNativeAPIGatewaySystemParametersResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewaySystemParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySystemParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAuthRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserAuthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAuthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrategyResource struct {

	// 鉴权策略ID

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 命名空间ID列表

	Namespaces []*StrategyResourceEntry `json:"Namespaces,omitempty" name:"Namespaces"`
	// 服务ID列表

	Services []*StrategyResourceEntry `json:"Services,omitempty" name:"Services"`
	// 配置组ID列表

	ConfigGroups []*StrategyResourceEntry `json:"ConfigGroups,omitempty" name:"ConfigGroups"`
}

type SystemPluginInfo struct {

	// 插件名称

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 功能类型

	FeatureType *string `json:"FeatureType,omitempty" name:"FeatureType"`
	// 当前版本

	CurrentVersion *string `json:"CurrentVersion,omitempty" name:"CurrentVersion"`
	// 最新版本

	LatestVersion *string `json:"LatestVersion,omitempty" name:"LatestVersion"`
	// 是否需要安装

	CanInstall *bool `json:"CanInstall,omitempty" name:"CanInstall"`
	// 是否需要升级

	CanUpdate *bool `json:"CanUpdate,omitempty" name:"CanUpdate"`
	// 插件描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 插件描述链接

	DescriptionLink *string `json:"DescriptionLink,omitempty" name:"DescriptionLink"`
}

type DeleteNativeGatewayServiceSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNativeGatewayServiceSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNativeGatewayServiceSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicensesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		List []*LicenseView `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicensesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicensesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallCloudNativeAPIGatewayPluginRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 插件名称

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
}

func (r *UninstallCloudNativeAPIGatewayPluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallCloudNativeAPIGatewayPluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFilter struct {

	// 过滤字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 过滤值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type NamespaceInfo struct {

	// 命名空间ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 命名空间名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeCloudNativeAPIGatewayConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关响应结果。

		Result *DescribeCloudNativeAPIGatewayConfigResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEurekaServiceInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务实例记录

		Content []*DescribeEurekaServiceInstance `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEurekaServiceInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEurekaServiceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateInstanceUpgradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		OpResult *string `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateInstanceUpgradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateInstanceUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigFileGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否删除成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigFileGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigFileGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNativeGatewayServerGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除信息

		Result *DeleteNativeGatewayServerGroupResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNativeGatewayServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNativeGatewayServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroySREInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务执行ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroySREInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroySREInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateRegionRequest struct {
	*tchttp.BaseRequest

	// 地域

	RegionField *string `json:"RegionField,omitempty" name:"RegionField"`
	// 地域中文

	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
	// 大区

	Area *string `json:"Area,omitempty" name:"Area"`
	// agent镜像地址

	AgentImageUrl *string `json:"AgentImageUrl,omitempty" name:"AgentImageUrl"`
	// 巴拉多地址

	BaradHost *string `json:"BaradHost,omitempty" name:"BaradHost"`
	// EKS地址

	EksHost *string `json:"EksHost,omitempty" name:"EksHost"`
	// VPC服务地址

	InnerVpcHost *string `json:"InnerVpcHost,omitempty" name:"InnerVpcHost"`
	// 顺序

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// TSE仓库地址

	TSEPublicRepoUrl *string `json:"TSEPublicRepoUrl,omitempty" name:"TSEPublicRepoUrl"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 地域发布状态。
	//
	// -PRE：发布中，外部不可见
	// -PROD：线上，外不可见

	ReleaseStatus *string `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`
}

func (r *ModifyOperateRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceGroupsRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户组ID列表

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

func (r *DeleteGovernanceGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayLatestTaskPhasesRequest struct {
	*tchttp.BaseRequest

	// 云原生网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DescribeCloudNativeAPIGatewayLatestTaskPhasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayLatestTaskPhasesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateConsoleInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 平台层实例列表

		List []*OperateConsoleInstanceItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateConsoleInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateConsoleInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigFileRequest struct {
	*tchttp.BaseRequest

	// TSE 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置文件列表详情

	ConfigFile *ConfigFile `json:"ConfigFile,omitempty" name:"ConfigFile"`
}

func (r *CreateConfigFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineSyncJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据同步任务数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据同步任务列表

		Content []*EngineSyncJob `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEngineSyncJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineSyncJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageTypeInfo struct {

	// 数据存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 数据存储类型名称

	StorageTypeName *string `json:"StorageTypeName,omitempty" name:"StorageTypeName"`
}

type DeleteOperateEngineVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperateEngineVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperateEngineVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCloudNativeAPIGatewaySecretsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Secret总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// secret 列表

		Secrets []*CloudNativeAPIGatewaySecret `json:"Secrets,omitempty" name:"Secrets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCloudNativeAPIGatewaySecretsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCloudNativeAPIGatewaySecretsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetGovernanceGroupTokenRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ResetGovernanceGroupTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetGovernanceGroupTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GovernanceInstanceUpdate struct {

	// 实例所在服务名。

	Service *string `json:"Service,omitempty" name:"Service"`
	// 实例服务所在命名空间。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 实例负载均衡权重信息。不填默认为100。

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 实例默认健康信息。不填默认为健康。

	Healthy *bool `json:"Healthy,omitempty" name:"Healthy"`
	// 实例隔离信息。不填默认为非隔离。

	Isolate *bool `json:"Isolate,omitempty" name:"Isolate"`
	// 实例ip。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 实例监听端口。

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 实例使用协议。不填默认为空。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 实例版本。不填默认为空。

	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`
	// 是否启用健康检查。不填默认不启用。

	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`
	// 上报心跳时间间隔。若 EnableHealthCheck 为不启用，则此参数不生效；若 EnableHealthCheck 启用，此参数不填，则默认 ttl 为 5s。

	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`
	// 治理中心服务实例id。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 元数据信息。

	Metadatas []*Metadata `json:"Metadatas,omitempty" name:"Metadatas"`
}

type DeleteGovernanceServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除服务结果。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGovernanceServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateEksRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	EksClusterId *string `json:"EksClusterId,omitempty" name:"EksClusterId"`
	// 地域

	RegionField *string `json:"RegionField,omitempty" name:"RegionField"`
	// vpc子网ID

	ClbVpcSubnetId *string `json:"ClbVpcSubnetId,omitempty" name:"ClbVpcSubnetId"`
	// 集群apiserver地址

	EksApiServerAddress *string `json:"EksApiServerAddress,omitempty" name:"EksApiServerAddress"`
	// 集群token

	EksApiServerToken *string `json:"EksApiServerToken,omitempty" name:"EksApiServerToken"`
	// vpc绑定id

	EksVpcBindId *string `json:"EksVpcBindId,omitempty" name:"EksVpcBindId"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 支持的引擎类型，用英文逗号分隔

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// 集群状态，支持以下参数：
	//
	// - enable : 可使用
	// - disable : 隔离中，不能使用

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyOperateEksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateEksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOneCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeOneCloudNativeAPIGatewayServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOneCloudNativeAPIGatewayServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceUserLoginInfoRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 腾讯云帐户Uin

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeGovernanceUserLoginInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceUserLoginInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateSubnetInfo struct {

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网IP

	CIDR *string `json:"CIDR,omitempty" name:"CIDR"`
	// 可用IP数

	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" name:"AvailableIpAddressCount"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 总IP数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 可用区状态：online正常 soldout售罄

	Status *string `json:"Status,omitempty" name:"Status"`
}

type Source struct {

	// 服务来源对应资源ID

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 服务来源的类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 只有在Type为容器集群时，才有此参数

	IngressClass *string `json:"IngressClass,omitempty" name:"IngressClass"`
}

type NativeGatewayServerGroup struct {

	// 云原生网关分组唯一id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 分组名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 节点规格、节点数信息

	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitempty" name:"NodeConfig"`
	// 网关分组状态。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否是默认分组。
	// 0：否。
	// 1：是。

	IsFirstGroup *int64 `json:"IsFirstGroup,omitempty" name:"IsFirstGroup"`
	// 关联策略信息

	BindingStrategy *CloudNativeAPIGatewayStrategy `json:"BindingStrategy,omitempty" name:"BindingStrategy"`
	// 网关实例 id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 带宽

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 子网id

	SubnetIds *string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

type DescribeAllConfigFileTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置文件模板列表

		ConfigFileTemplates []*ConfigFileTemplate `json:"ConfigFileTemplates,omitempty" name:"ConfigFileTemplates"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllConfigFileTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllConfigFileTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVersionRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 升级的版本号

	UpgradeVersion *string `json:"UpgradeVersion,omitempty" name:"UpgradeVersion"`
}

func (r *ModifyInstanceVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineExportPortsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口数组，北极星自有协议

		ServerPorts []*ServicePort `json:"ServerPorts,omitempty" name:"ServerPorts"`
		// 端口数组，OpenAPI

		HttpPorts []*ServicePort `json:"HttpPorts,omitempty" name:"HttpPorts"`
		// 端口数组，兼容协议

		CompatiblePorts []*ServicePort `json:"CompatiblePorts,omitempty" name:"CompatiblePorts"`
		// 端口数组，（废弃）

		ServicePorts []*ServicePort `json:"ServicePorts,omitempty" name:"ServicePorts"`
		// 端口数组，（废弃）

		Ports []*ServicePort `json:"Ports,omitempty" name:"Ports"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEngineExportPortsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineExportPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMainControlRegionInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主控区域相关信息列表

		RegionInfos []*RegionInfo `json:"RegionInfos,omitempty" name:"RegionInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMainControlRegionInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMainControlRegionInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNacosMonitorDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控内容详情

		Result *MonitorResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNacosMonitorDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosMonitorDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClosePrometheusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClosePrometheusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClosePrometheusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNativeGatewayServiceSourceRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 服务来源实例名称

	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
	// 微服务引擎类型，TSE Nacos｜TSE Consul｜TSE PolarisMesh｜Nacos｜Consul｜PolarisMesh｜TEM｜TKE｜EKS

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 服务来源实例额外信息

	SourceInfo *SourceInfo `json:"SourceInfo,omitempty" name:"SourceInfo"`
}

func (r *CreateNativeGatewayServiceSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNativeGatewayServiceSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateEngineVersionRequest struct {
	*tchttp.BaseRequest

	// 引擎版本号

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 无

	IsModify *string `json:"IsModify,omitempty" name:"IsModify"`
}

func (r *CreateOperateEngineVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateEngineVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSRECodeNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务执行阶段文案列表

		TaskPhases []*TaskPhaseName `json:"TaskPhases,omitempty" name:"TaskPhases"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSRECodeNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSRECodeNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FinishMigrateRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *FinishMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FinishMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayPluginVersionsRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 插件名称。

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListCloudNativeAPIGatewayPluginVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayPluginVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayClsConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayClsConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayClsConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCloudNativeAPIGatewaySpecRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 网关分组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 云原生API网关节点配置。

	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitempty" name:"NodeConfig"`
}

func (r *UpdateCloudNativeAPIGatewaySpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCloudNativeAPIGatewaySpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperReplicasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 注册引擎实例副本信息

		Replicas []*ZookeeperReplica `json:"Replicas,omitempty" name:"Replicas"`
		// 副本个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineClbAccessAddressRequest struct {
	*tchttp.BaseRequest

	// 引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 部署地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *DescribeEngineClbAccessAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineClbAccessAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSRECodeNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSRECodeNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSRECodeNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigFilesRequest struct {
	*tchttp.BaseRequest

	// TSE实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 配置分组名称

	Group *string `json:"Group,omitempty" name:"Group"`
	// 配置文件名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteConfigFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNativeGatewayServicesRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例名称，模糊搜素

	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
	// 服务名称，模糊搜索

	Service *string `json:"Service,omitempty" name:"Service"`
	// 排序字段，仅支持 Service

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序类型，AES/DESC

	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
	// 单页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 服务唯一ID标识

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 服务来源类型

	SourceTypes []*string `json:"SourceTypes,omitempty" name:"SourceTypes"`
}

func (r *DescribeNativeGatewayServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNativeGatewayServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayPluginsResult struct {

	// 总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 云原生API网关插件列表。

	PluginList []*CloudNativeAPIGatewayPlugin `json:"PluginList,omitempty" name:"PluginList"`
}

type NginxIngressVersions struct {

	// 版本描述信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 版本列表

	Versions []*NginxIngressVersion `json:"Versions,omitempty" name:"Versions"`
}

type DescribeCloudNativeAPIGatewayIngressRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// k8s集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// Ingress的名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeCloudNativeAPIGatewayIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineSyncJobCheckResultRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeEngineSyncJobCheckResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineSyncJobCheckResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadLicenseRequest struct {
	*tchttp.BaseRequest

	// License申请单的ID

	RecordID *string `json:"RecordID,omitempty" name:"RecordID"`
}

func (r *DownloadLicenseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadLicenseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadZookeeperMigrateDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LoadZookeeperMigrateDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoadZookeeperMigrateDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayPublicNetworkRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 分组id。

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 公网负载均衡配置。

	InternetConfig *InternetConfig `json:"InternetConfig,omitempty" name:"InternetConfig"`
}

func (r *CreateCloudNativeAPIGatewayPublicNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayPublicNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKongIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Result *KongIngressStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKongIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKongIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置中心类型（consul、zookeeper、apollo等）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 请求命名 PUT GET DELETE

	Command *string `json:"Command,omitempty" name:"Command"`
	// 配置的Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 配置的Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ManageConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartEngineSyncJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartEngineSyncJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartEngineSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckZookeeperMigrateTargetStatusRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群信息

	ClusterInfoList []*ZookeeperMigrateMetaClusterInfo `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`
	// 管理员名称

	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`
	// 管理员密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
}

func (r *CheckZookeeperMigrateTargetStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckZookeeperMigrateTargetStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGovernanceServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGovernanceServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApolloEnvReplicasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点信息列表

		Replicas []*ApolloReplicaPodDetail `json:"Replicas,omitempty" name:"Replicas"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApolloEnvReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApolloEnvReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolarisReplicasRequest struct {
	*tchttp.BaseRequest

	// 北极星引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePolarisReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolarisReplicasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayCertificateDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Result *KongCertificate `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewayCertificateDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayCertificateDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NativeGatewayServiceSourceItem struct {

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 服务来源名称

	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
	// 服务来源类型

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 服务来源额外信息

	SourceInfo *SourceInfo `json:"SourceInfo,omitempty" name:"SourceInfo"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type CreateNativeGatewayServiceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNativeGatewayServiceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNativeGatewayServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineSyncJobCheckStepTip struct {

	// 错误码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 解决方式

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 帮助文档

	HelpDoc *string `json:"HelpDoc,omitempty" name:"HelpDoc"`
}

type GetZookeeperMigrateUploadMetaInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetZookeeperMigrateUploadMetaInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZookeeperMigrateUploadMetaInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallCloudNativeAPIGatewayPluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallCloudNativeAPIGatewayPluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallCloudNativeAPIGatewayPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KongTarget struct {

	// Host

	Host *string `json:"Host,omitempty" name:"Host"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 权重

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 健康状态

	Health *string `json:"Health,omitempty" name:"Health"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type PolarisLimiterAddress struct {

	// VPC接入IP列表

	IntranetAddress *string `json:"IntranetAddress,omitempty" name:"IntranetAddress"`
}

type DescribeAllConfigFileTemplatesRequest struct {
	*tchttp.BaseRequest

	// TSE实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAllConfigFileTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllConfigFileTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCloudNativeAPIGatewayDomainsRequest struct {
	*tchttp.BaseRequest

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *GetCloudNativeAPIGatewayDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCloudNativeAPIGatewayDomainsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件仪表盘URL

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNativeGatewayServerGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNativeGatewayServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNativeGatewayServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayCodeNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务执行阶段文案列表

		TaskPhases []*CloudNativeAPIGatewayTaskPhaseName `json:"TaskPhases,omitempty" name:"TaskPhases"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayCodeNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayCodeNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KVPair struct {

	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeGovernanceServicesRequest struct {
	*tchttp.BaseRequest

	// 按照服务名过滤，精确匹配。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 按照命名空间过滤，精确匹配。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 使用元数据过滤，目前只支持一组元组数，若传了多条，只会使用第一条元数据过滤。

	Metadatas []*Metadata `json:"Metadatas,omitempty" name:"Metadatas"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// tse 实例 id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务所属部门。

	Department *string `json:"Department,omitempty" name:"Department"`
	// 服务所属业务。

	Business *string `json:"Business,omitempty" name:"Business"`
	// 服务中实例的ip，用来过滤服务。

	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DescribeGovernanceServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 用户uin

	UinField *string `json:"UinField,omitempty" name:"UinField"`
	// 白名单类型：apollo | region

	Type *string `json:"Type,omitempty" name:"Type"`
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南
	// - NorthAmerica，北美

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *ModifyOperateWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalerResourceStrategiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关实例弹性伸缩策略列表响应结果。

		Result *ListCloudNativeAPIGatewayStrategyResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalerResourceStrategiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoScalerResourceStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateFeatureVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateFeatureVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateFeatureVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {

	// 用户ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 用户名称，对应云上的话应该是UIN

	Name *string `json:"Name,omitempty" name:"Name"`
	// 对应主账户的UIN

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 用户来源，QCloud时不需要设置Password

	Source *string `json:"Source,omitempty" name:"Source"`
	// 用户鉴权Token

	AuthToken *string `json:"AuthToken,omitempty" name:"AuthToken"`
	// 该token是否被禁用

	TokenEnable *bool `json:"TokenEnable,omitempty" name:"TokenEnable"`
	// 该账户的简单描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 该账户的创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 该账户的最近一次修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 帐户密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 帐户手机号

	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
	// 帐户邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
}

type DescribeCloudNativeAPIGatewaysRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 请求过滤参数，支持按照实例名称、ID和标签键值（Name、GatewayId、Tag）筛选

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCloudNativeAPIGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntranetAccessAddress struct {

	// 打通网络的vpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 打通网络的subnetId

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 打通网络的访问地址

	IntranetAddress *string `json:"IntranetAddress,omitempty" name:"IntranetAddress"`
	// 网络打通状态：normal正常 ｜initializing初始化中 ｜ failed 失败

	Status *string `json:"Status,omitempty" name:"Status"`
}

type ListCloudNativeAPIGatewaySystemParametersResult struct {

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关系统参数数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 云原生API网关系统参数列表。

	SystemParameterList []*CloudNativeAPIGatewaySystemParameter `json:"SystemParameterList,omitempty" name:"SystemParameterList"`
	// 云原生API网关系统参数JSON格式字符串

	Value *string `json:"Value,omitempty" name:"Value"`
	// 配置文件的描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 配置文件的名字

	Key *string `json:"Key,omitempty" name:"Key"`
	// 配置文件的Title

	Title *string `json:"Title,omitempty" name:"Title"`
	// 文档链接

	DocLink *string `json:"DocLink,omitempty" name:"DocLink"`
	// 修改系统参数提示

	ChangeAlert *string `json:"ChangeAlert,omitempty" name:"ChangeAlert"`
}

type CloseKongIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseKongIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKongIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceAliasesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGovernanceAliasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceAliasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EurekaReplica struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 可用区ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type ListCloudNativeAPIGatewayStrategyResult struct {

	// 总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 云原生API网关实例策略列表。

	StrategyList []*CloudNativeAPIGatewayStrategy `json:"StrategyList,omitempty" name:"StrategyList"`
}

type DeleteNativeGatewayServerGroupResult struct {

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 网关分组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 删除状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeleteCloudNativeAPIGatewayPluginRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 插件名称

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 插件版本号

	PluginVersion *string `json:"PluginVersion,omitempty" name:"PluginVersion"`
}

func (r *DeleteCloudNativeAPIGatewayPluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayPluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeService struct {

	// 注册服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 正常服务实例个数

	NormalServiceInstanceCount *int64 `json:"NormalServiceInstanceCount,omitempty" name:"NormalServiceInstanceCount"`
	// 所有实例个数

	TotalServiceInstanceCount *int64 `json:"TotalServiceInstanceCount,omitempty" name:"TotalServiceInstanceCount"`
	// 服务状态

	ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`
}

type EngineSyncJobCheckStepInfo struct {

	// 步骤编号

	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`
	// 步骤名称

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 步骤标号

	StepId *string `json:"StepId,omitempty" name:"StepId"`
	// 当前状态，failed | finished

	Status *string `json:"Status,omitempty" name:"Status"`
	// 当前步骤进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 错误信息

	Errors *EngineSyncJobCheckStepTip `json:"Errors,omitempty" name:"Errors"`
	// 警告信息

	Warnings *EngineSyncJobCheckStepTip `json:"Warnings,omitempty" name:"Warnings"`
}

type GovernanceDiscoverEvent struct {

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务名称

	Service *string `json:"Service,omitempty" name:"Service"`
	// 实例IP

	Host *string `json:"Host,omitempty" name:"Host"`
	// 实例端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 事件类型

	EType *string `json:"EType,omitempty" name:"EType"`
	// 事件发生事件戳，单位秒

	CreateTimeSec *uint64 `json:"CreateTimeSec,omitempty" name:"CreateTimeSec"`
	// 服务实例事件

	EBType *string `json:"EBType,omitempty" name:"EBType"`
	// 服务实例事件中文描述

	EventTypeDesc *string `json:"EventTypeDesc,omitempty" name:"EventTypeDesc"`
}

type DescribeCloudNativeAPIGatewayUserAuthRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCloudNativeAPIGatewayUserAuthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayUserAuthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEurekaServicesRequest struct {
	*tchttp.BaseRequest

	// 请求过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 注册引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeEurekaServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEurekaServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperReplicasRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 副本列表Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 副本列表Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeZookeeperReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperReplicasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrometheusStatus struct {

	// 是否开启了Prometheus指标推送

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 开启了Prometheus指标推送集群

	PromId *string `json:"PromId,omitempty" name:"PromId"`
}

type ZooKeeperDefaultConfigInfo struct {

	// 参数名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 默认运行参数值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 参数描述说明

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// true表示只读，不支持修改

	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 配置的类型说明

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 配置类型的限制描述

	ConfigTypeLimit *string `json:"ConfigTypeLimit,omitempty" name:"ConfigTypeLimit"`
	// 此配置所属zk版本范围

	VersionRange *string `json:"VersionRange,omitempty" name:"VersionRange"`
	// 是否不可见

	Invisible *bool `json:"Invisible,omitempty" name:"Invisible"`
	// 对应的分组:zoo、sasl、env

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 依赖的配置项，需要同时满足

	DependencyConfigs *string `json:"DependencyConfigs,omitempty" name:"DependencyConfigs"`
}

type ModifyOperateEngineVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateEngineVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateEngineVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNginxIngressGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一个实例最大的节点数

		InstanceMaxNodesSize *uint64 `json:"InstanceMaxNodesSize,omitempty" name:"InstanceMaxNodesSize"`
		// Nginx Ingress 模板

		Template *string `json:"Template,omitempty" name:"Template"`
		// Nginx Ingress ConfigMap 参数信息

		ConfigMapParams *NginxIngressConfigMapParams `json:"ConfigMapParams,omitempty" name:"ConfigMapParams"`
		// Nginx Ingress 版本信息

		NginxIngressVersions *NginxIngressVersions `json:"NginxIngressVersions,omitempty" name:"NginxIngressVersions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNginxIngressGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNginxIngressGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceUserRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户

	User *User `json:"User,omitempty" name:"User"`
}

func (r *ModifyGovernanceUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSimpleInfosRequest struct {
	*tchttp.BaseRequest

	// 列表个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 列表偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 展示实例列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstanceSimpleInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceSimpleInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewaySecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewaySecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewaySecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCredentialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 凭证key

		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
		// 凭证ID

		SecretID *string `json:"SecretID,omitempty" name:"SecretID"`
		// 临时token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 失效时间，单位秒

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCredentialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineSyncJob struct {

	// 引擎数据同步任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 引擎数据同步任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 引擎数据同步任务运行模式

	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`
	// 引擎数据同步任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 引擎数据同步任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 引擎数据同步任务运行时间

	RunTime *string `json:"RunTime,omitempty" name:"RunTime"`
	// 引擎数据同步任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 总步骤数

	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`
	// 当前步骤

	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`
	// 总体描述信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 引擎数据同步任务接入类型

	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`
	// 源端数据库IP

	SrcDbIp *string `json:"SrcDbIp,omitempty" name:"SrcDbIp"`
	// 源端数据库Port

	SrcDbPort *int64 `json:"SrcDbPort,omitempty" name:"SrcDbPort"`
	// 源端数据库类型

	SrcDbType *string `json:"SrcDbType,omitempty" name:"SrcDbType"`
	// 源端数据库名称

	SrcDbName *string `json:"SrcDbName,omitempty" name:"SrcDbName"`
	// 源端数据库用户名

	SrcDbUser *string `json:"SrcDbUser,omitempty" name:"SrcDbUser"`
	// 引擎数据同步任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 引擎数据同步任务详细步骤信息

	StepInfos []*StepInfo `json:"StepInfos,omitempty" name:"StepInfos"`
	// 同步两端数据量差距

	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`
	// 同步两端时间差距

	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`
	// 数据库名称映射关系

	DbMappings []*KVPair `json:"DbMappings,omitempty" name:"DbMappings"`
	// 源数据库所在地域

	SrcDbRegion *string `json:"SrcDbRegion,omitempty" name:"SrcDbRegion"`
}

type DeleteCloudNativeAPIGatewayCanaryRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayCanaryRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayCanaryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生网关节点列表结果。

		Result *DescribeCloudNativeAPIGatewayNodesResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewaySimpleInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云原生网关列表

		GatewayList []*TagCloudNativeAPIGatewayDescription `json:"GatewayList,omitempty" name:"GatewayList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewaySimpleInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaySimpleInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunExecTaskBatchJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务批次索引

		BatchIndex *uint64 `json:"BatchIndex,omitempty" name:"BatchIndex"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunExecTaskBatchJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunExecTaskBatchJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NativeGatewayServiceGroupItem struct {

	// 分组名称

	Group *string `json:"Group,omitempty" name:"Group"`
	// 标签数组

	Labels []*InstanceLabel `json:"Labels,omitempty" name:"Labels"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 实例所占百分比

	Percent *string `json:"Percent,omitempty" name:"Percent"`
}

type EnableIntranetRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// true开 false关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 内网类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 实例为apollo时必填，apollo环境名称

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 部署地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *EnableIntranetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableIntranetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateSecurityGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作统计结果

		Data *OperateSecurityGroupResult `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolarisAuthStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolarisAuthStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolarisAuthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageConsulConsoleRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 模块路由

	ReqAction *string `json:"ReqAction,omitempty" name:"ReqAction"`
	// 请求数据

	ReqBody *string `json:"ReqBody,omitempty" name:"ReqBody"`
}

func (r *ManageConsulConsoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageConsulConsoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebuildPolarisCLSTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修复结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebuildPolarisCLSTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebuildPolarisCLSTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigFilesByGroupRequest struct {
	*tchttp.BaseRequest

	// TSE实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间名

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 组名

	Group *string `json:"Group,omitempty" name:"Group"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConfigFilesByGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFilesByGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KongRouteList struct {

	// 无

	RouteList []*KongRoutePreview `json:"RouteList,omitempty" name:"RouteList"`
	// 剩余列表的offset

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ChangeNacosAuthStatusRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// console_auth_enable：控制台，sdk_auth_enable：sdk

	NacosAuthType *string `json:"NacosAuthType,omitempty" name:"NacosAuthType"`
	// true开 false关

	NacosAuthStatus *bool `json:"NacosAuthStatus,omitempty" name:"NacosAuthStatus"`
}

func (r *ChangeNacosAuthStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeNacosAuthStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigFilesRequest struct {
	*tchttp.BaseRequest

	// TSE实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 组名

	Group *string `json:"Group,omitempty" name:"Group"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标签列表

	Tags []*ConfigFileTag `json:"Tags,omitempty" name:"Tags"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConfigFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnginePodInfo struct {

	// 节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网卡ip

	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`
	// 节点ip

	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`
	// 节点地址绑定的vpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 节点地址绑定的subnetId

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type CreateConfigFileGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否创建成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigFileGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigFileGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceGroupTokenRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGovernanceGroupTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceGroupTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayCanaryRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 灰度规则列表

		Result *CloudAPIGatewayCanaryRuleList `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudNativeAPIGatewayCanaryRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayCanaryRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayBasicLogRequest struct {
	*tchttp.BaseRequest

	// 网关 id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 节点名字。例如 kong 网关一般为 kong-0, kong-1......kong-n，其中 n 为 购买的副本数-1

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 日志名字。如 nginx 网关取值为 AccessLog, ErrorLog

	LogName *string `json:"LogName,omitempty" name:"LogName"`
	// 取最新的 Limit 条数据。缺省值为 100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 网关日志类型

	GatewayLogType *string `json:"GatewayLogType,omitempty" name:"GatewayLogType"`
}

func (r *DescribeCloudNativeAPIGatewayBasicLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayBasicLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务数总量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务信息详情。

		Content []*GovernanceService `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperDigestTokenRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户名

	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`
	// 密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
}

func (r *DescribeZookeeperDigestTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperDigestTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServicesRequest struct {
	*tchttp.BaseRequest

	// 请求过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 注册引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeZookeeperServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayPluginVersionsResult struct {

	// 总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 插件信息列表。

	PluginInfoList []*CloudNativeAPIGatewayPluginInfo `json:"PluginInfoList,omitempty" name:"PluginInfoList"`
}

type BindAutoScalerResourceStrategyToGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAutoScalerResourceStrategyToGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoScalerResourceStrategyToGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSREInstanceFeatureRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 特性名称

	FeatureKey *string `json:"FeatureKey,omitempty" name:"FeatureKey"`
}

func (r *CheckSREInstanceFeatureRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSREInstanceFeatureRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsulReplicasRequest struct {
	*tchttp.BaseRequest

	// 注册引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 分页参数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConsulReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsulReplicasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTagInfosRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceTagInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTagInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KongCertificatesDomainPreview struct {

	// 证书域名

	BindDomain *string `json:"BindDomain,omitempty" name:"BindDomain"`
	// 证书名称

	Names []*string `json:"Names,omitempty" name:"Names"`
}

type CloudNativeAPIGatewayStrategyBindingGroupInfo struct {

	// 网关分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 节点配置

	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitempty" name:"NodeConfig"`
	// 绑定时间

	BindTime *string `json:"BindTime,omitempty" name:"BindTime"`
	// 网关分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 绑定状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type ServiceGovernanceEdition struct {

	// 服务治理引擎的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务治理引擎的版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

type CreateGovernanceUsersRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户列表

	Users []*User `json:"Users,omitempty" name:"Users"`
}

func (r *CreateGovernanceUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateEksRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	EksClusterId *string `json:"EksClusterId,omitempty" name:"EksClusterId"`
	// 地域

	RegionField *string `json:"RegionField,omitempty" name:"RegionField"`
	// vpc子网ID

	ClbVpcSubnetId *string `json:"ClbVpcSubnetId,omitempty" name:"ClbVpcSubnetId"`
	// 集群apiserver地址

	EksApiServerAddress *string `json:"EksApiServerAddress,omitempty" name:"EksApiServerAddress"`
	// 集群token

	EksApiServerToken *string `json:"EksApiServerToken,omitempty" name:"EksApiServerToken"`
	// vpc绑定id

	EksVpcBindId *string `json:"EksVpcBindId,omitempty" name:"EksVpcBindId"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 支持的引擎类型，用英文逗号分隔

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// 集群状态，支持以下参数：
	//
	// - enable : 可使用
	// - disable : 隔离中，不能使用

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *CreateOperateEksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateEksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperateEngineVersionRequest struct {
	*tchttp.BaseRequest

	// 引擎版本号

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *ModifyOperateEngineVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperateEngineVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateInstance struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本号

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 状态, 枚举值:creating/create_fail/running/updating/update_fail/restarting/restart_fail/destroying/destroy_fail

	Status *string `json:"Status,omitempty" name:"Status"`
	// 规格ID

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 副本数

	Replica *int64 `json:"Replica,omitempty" name:"Replica"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Vpc iD

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 是否开启持久化存储

	EnableStorage *bool `json:"EnableStorage,omitempty" name:"EnableStorage"`
	// 数据存储方式

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 云硬盘容量

	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`
	// 计费方式

	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
	// EKS集群的ID

	EKSClusterID *string `json:"EKSClusterID,omitempty" name:"EKSClusterID"`
	// 集群创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 环境配置信息列表

	EnvInfos []*EnvInfo `json:"EnvInfos,omitempty" name:"EnvInfos"`
	// 引擎所在的区域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 注册引擎是否开启公网

	EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`
	// pod列表

	PodList []*OperatePodItem `json:"PodList,omitempty" name:"PodList"`
	// service列表

	ServiceList []*OperateServiceItem `json:"ServiceList,omitempty" name:"ServiceList"`
	// 用户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 负载列表

	WorkloadList []*OperateWorkloadItem `json:"WorkloadList,omitempty" name:"WorkloadList"`
	// 实例所属的安全组ID

	UserSecretGroupName *string `json:"UserSecretGroupName,omitempty" name:"UserSecretGroupName"`
	// 实例tag列表

	TagList []*KVPair `json:"TagList,omitempty" name:"TagList"`
}

type DescribeNacosConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统参数列表

		NacosConfigInfoList []*NacosConfigInfo `json:"NacosConfigInfoList,omitempty" name:"NacosConfigInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNacosConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZooKeeperDefaultConfigInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZooKeeperDefaultConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZooKeeperDefaultConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartEngineSyncJobRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据同步任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StartEngineSyncJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartEngineSyncJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserGroup struct {

	// 用户组ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 用户组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 对应主账户UIN

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 该用户组的授权Token

	AuthToken *string `json:"AuthToken,omitempty" name:"AuthToken"`
	// 该用户组的授权Token是否可用

	TokenEnable *bool `json:"TokenEnable,omitempty" name:"TokenEnable"`
	// 简单描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 该用户组下的用户ID列表信息

	Relation *GroupRelation `json:"Relation,omitempty" name:"Relation"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 该用户组下的用户数量

	UserCount *uint64 `json:"UserCount,omitempty" name:"UserCount"`
}

type DeleteGovernanceAliasesRequest struct {
	*tchttp.BaseRequest

	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务别名列表

	GovernanceAliases []*GovernanceAlias `json:"GovernanceAliases,omitempty" name:"GovernanceAliases"`
}

func (r *DeleteGovernanceAliasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceAliasesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalerResourceStrategyBindingGroupsRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 策略ID

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DescribeAutoScalerResourceStrategyBindingGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoScalerResourceStrategyBindingGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOneCloudNativeAPIGatewayServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Result *KongServiceDetail `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOneCloudNativeAPIGatewayServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOneCloudNativeAPIGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExecTaskBatchJobInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 批次信息列表

		List []*OptTaskBatchInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExecTaskBatchJobInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExecTaskBatchJobInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 白名单数组

		List []*OperateWhiteListItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableEngineInternetRequest struct {
	*tchttp.BaseRequest

	// 引擎ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// true开 false关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *EnableEngineInternetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableEngineInternetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NacosReplica struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 可用区ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeGovernanceUserLoginInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 帐户名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 帐户密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceUserLoginInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceUserLoginInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperMigrateMergeStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分为initiate、mergeSuccess、mergeFail、mergeDoing、breakDoing、breakSuccess、breakFail

		Status *string `json:"Status,omitempty" name:"Status"`
		// 状态描述

		StatusMessage *string `json:"StatusMessage,omitempty" name:"StatusMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperMigrateMergeStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperMigrateMergeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApolloReplicasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// portal 节点信息

		PortalReplicas []*ApolloReplicaPodDetail `json:"PortalReplicas,omitempty" name:"PortalReplicas"`
		// 环境列表

		EnvReplicas []*ApolloEnvReplica `json:"EnvReplicas,omitempty" name:"EnvReplicas"`
		// 4

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApolloReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApolloReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 注册服务服务名列表

		Content []*string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApolloSpec struct {

	// 规格id

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 规格名称

	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`
	// cpu核数

	CpuNum *int64 `json:"CpuNum,omitempty" name:"CpuNum"`
	// 内存大小

	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`
	// 最小限制

	MinLimit *int64 `json:"MinLimit,omitempty" name:"MinLimit"`
	// 最大限制

	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
	// 是否在白名单

	InWhiteList *bool `json:"InWhiteList,omitempty" name:"InWhiteList"`
}

type KongCertificatesDomainList struct {

	// 证书域名列表总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 无

	DomainsList []*KongCertificatesDomainPreview `json:"DomainsList,omitempty" name:"DomainsList"`
	// 证书域名列表总页数

	Pages *int64 `json:"Pages,omitempty" name:"Pages"`
}

type KVMapping struct {

	// key

	Key *string `json:"Key,omitempty" name:"Key"`
	// value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type SourceInfo struct {

	// 微服务引擎接入IP地址信息

	Addresses []*string `json:"Addresses,omitempty" name:"Addresses"`
	// 微服务引擎VPC信息

	VpcInfo *SourceInstanceVpcInfo `json:"VpcInfo,omitempty" name:"VpcInfo"`
	// 微服务引擎鉴权信息

	Auth *SourceInstanceAuth `json:"Auth,omitempty" name:"Auth"`
}

type DescribeOperateRegionsRequest struct {
	*tchttp.BaseRequest

	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperateRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSRELatestTaskPhasesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务执行阶段列表

		Phases []*SRETaskPhase `json:"Phases,omitempty" name:"Phases"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSRELatestTaskPhasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSRELatestTaskPhasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceNamespacesRequest struct {
	*tchttp.BaseRequest

	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间信息。

	GovernanceNamespaces []*GovernanceNamespaceInput `json:"GovernanceNamespaces,omitempty" name:"GovernanceNamespaces"`
}

func (r *ModifyGovernanceNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceServicesRequest struct {
	*tchttp.BaseRequest

	// tse 实例 id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务信息。

	GovernanceServices []*GovernanceServiceInput `json:"GovernanceServices,omitempty" name:"GovernanceServices"`
}

func (r *ModifyGovernanceServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceInstancesRequest struct {
	*tchttp.BaseRequest

	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 要删除的服务实例信息。

	GovernanceInstances []*GovernanceInstanceUpdate `json:"GovernanceInstances,omitempty" name:"GovernanceInstances"`
}

func (r *DeleteGovernanceInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigFileGroup struct {

	// 配置文件组id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 配置文件组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者

	CreateBy *string `json:"CreateBy,omitempty" name:"CreateBy"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 修改者

	ModifyBy *string `json:"ModifyBy,omitempty" name:"ModifyBy"`
	// 文件数

	FileCount *uint64 `json:"FileCount,omitempty" name:"FileCount"`
	// 关联用户，link_users

	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
	// 组id，link_groups

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// remove_link_users

	RemoveUserIds []*string `json:"RemoveUserIds,omitempty" name:"RemoveUserIds"`
	// remove_link_groups

	RemoveGroupIds []*string `json:"RemoveGroupIds,omitempty" name:"RemoveGroupIds"`
	// 是否可编辑

	Editable *bool `json:"Editable,omitempty" name:"Editable"`
	// 归属者

	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

type DeleteCloudNativeAPIGatewayPluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayPluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewaySecretRequest struct {
	*tchttp.BaseRequest

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 集群id，如果不填则查询网关所有关联的集群

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 查询的Namespace，如果不填则查询所有 Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 证书名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteCloudNativeAPIGatewaySecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewaySecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigFileGroupRequest struct {
	*tchttp.BaseRequest

	// tse 实例 id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 组

	Group *string `json:"Group,omitempty" name:"Group"`
}

func (r *DeleteConfigFileGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigFileGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取云原生API网关实例列表响应结果。

		Result *ListCloudNativeAPIGatewayResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateAgentDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// agent下的实例数组

		List []*OperateAgentInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateAgentDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateAgentDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 查询类型：info/workload/pod/service

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateInstanceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZooKeeperConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统参数列表

		ZooKeeperConfigInfoList []*ZooKeeperConfigInfo `json:"ZooKeeperConfigInfoList,omitempty" name:"ZooKeeperConfigInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZooKeeperConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZooKeeperConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartZookeeperMigratePrepareRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StartZookeeperMigratePrepareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartZookeeperMigratePrepareRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KongServiceList struct {

	// 无

	ServiceList []*KongServicePreview `json:"ServiceList,omitempty" name:"ServiceList"`
	// 剩余列表的offset

	Offset *string `json:"Offset,omitempty" name:"Offset"`
}

type CreateNativeGatewayServerGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关分组创建信息

		Result *CreateCloudNativeAPIGatewayServerGroupResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNativeGatewayServerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNativeGatewayServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateFeatureVersionRequest struct {
	*tchttp.BaseRequest

	// 中国区

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 功能版本

	FeatureVersion *string `json:"FeatureVersion,omitempty" name:"FeatureVersion"`
	// 功能版本名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 引擎类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 支持的副本数列表

	Replicas *string `json:"Replicas,omitempty" name:"Replicas"`
	// 支持的规格列表

	Specs *string `json:"Specs,omitempty" name:"Specs"`
	// 无

	IsModify *string `json:"IsModify,omitempty" name:"IsModify"`
}

func (r *CreateOperateFeatureVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateFeatureVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 访问后端协议，http或https

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问后端的地址

	Path *string `json:"Path,omitempty" name:"Path"`
	// 后端超时时间，单位ms

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 后端重试次数

	Retries *int64 `json:"Retries,omitempty" name:"Retries"`
	// 后端类型

	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// 后端配置

	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitempty" name:"UpstreamInfo"`
	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *ModifyCloudNativeAPIGatewayServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGovernanceServicesRequest struct {
	*tchttp.BaseRequest

	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务信息。

	GovernanceServices []*GovernanceServiceInput `json:"GovernanceServices,omitempty" name:"GovernanceServices"`
}

func (r *DeleteGovernanceServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExecTaskBatchJobExeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 批次执行描述列表

		List []*OptTaskBatchExeInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExecTaskBatchJobExeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExecTaskBatchJobExeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAuthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否有TSE_QCSRole的角色

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserAuthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineRegionInfo struct {

	// 引擎节点所在地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 此地域节点分配数量

	Replica *int64 `json:"Replica,omitempty" name:"Replica"`
	// 集群网络信息

	VpcInfos []*VpcInfo `json:"VpcInfos,omitempty" name:"VpcInfos"`
}

type DescribeCloudNativeAPIGatewayOriginalPluginsRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 搜索关键字

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeCloudNativeAPIGatewayOriginalPluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayOriginalPluginsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateSecurityGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 安全组数组

		List []*OperateSecurityGroupItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateSecurityGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateSecurityGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例记录

		Content []*SREInstance `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSREInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSREInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEurekaServiceInstance struct {

	// 实例IP

	ServiceInstanceIp *string `json:"ServiceInstanceIp,omitempty" name:"ServiceInstanceIp"`
	// 实例端口

	ServiceInstancePort *int64 `json:"ServiceInstancePort,omitempty" name:"ServiceInstancePort"`
	// 实例状态

	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`
}

type LicenseView struct {

	// 记录ID

	RecordID *string `json:"RecordID,omitempty" name:"RecordID"`
	// 合同ID

	ContractNO *string `json:"ContractNO,omitempty" name:"ContractNO"`
	// 客户名称

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 销售方公司名称

	SellerName *string `json:"SellerName,omitempty" name:"SellerName"`
	// 代理商名称

	AgencyName *string `json:"AgencyName,omitempty" name:"AgencyName"`
	// 可用时间

	ActiveDays *int64 `json:"ActiveDays,omitempty" name:"ActiveDays"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 校验列表

	LicenseRuleList []*LicenseRuleInfo `json:"LicenseRuleList,omitempty" name:"LicenseRuleList"`
	// 生效时间

	EffectTime *string `json:"EffectTime,omitempty" name:"EffectTime"`
	// 加密License字符串

	LicenseCipher *string `json:"LicenseCipher,omitempty" name:"LicenseCipher"`
	// License的Key。也就是文件名

	LicenseKey *string `json:"LicenseKey,omitempty" name:"LicenseKey"`
}

type DescribeAutoScalerResourceStrategyBindingGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云原生API网关实例策略绑定网关分组列表响应结果

		Result *ListCloudNativeAPIGatewayStrategyBindingGroupInfoResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalerResourceStrategyBindingGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoScalerResourceStrategyBindingGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineSyncJobsRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeEngineSyncJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineSyncJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PortInfos struct {

	// 操作的端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 开启或者关闭

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type DeleteNativeGatewayServerGroupRequest struct {
	*tchttp.BaseRequest

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 网关分组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteNativeGatewayServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNativeGatewayServerGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateConsoleInstanceLogsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	EKSClusterID *string `json:"EKSClusterID,omitempty" name:"EKSClusterID"`
	// pod名

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 日志行数

	Line *int64 `json:"Line,omitempty" name:"Line"`
	// 主控

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeOperateConsoleInstanceLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateConsoleInstanceLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayRouteRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 所属服务的ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 路由的名字，实例级别唯一，可以不提供

	RouteName *string `json:"RouteName,omitempty" name:"RouteName"`
	// 路由的方法

	Methods []*string `json:"Methods,omitempty" name:"Methods"`
	// 路由的域名

	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
	// 路由的路径

	Paths []*string `json:"Paths,omitempty" name:"Paths"`
	// 路由的协议，可选https, http

	Protocols []*string `json:"Protocols,omitempty" name:"Protocols"`
	// 转发到后端时是否保留Host

	PreserveHost *bool `json:"PreserveHost,omitempty" name:"PreserveHost"`
	// https重定向状态码

	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitempty" name:"HttpsRedirectStatusCode"`
	// 转发到后端时是否StripPath

	StripPath *bool `json:"StripPath,omitempty" name:"StripPath"`
	// 是否开启强制HTTPS

	ForceHttps *bool `json:"ForceHttps,omitempty" name:"ForceHttps"`
	// 路由的ID，实例级别唯一

	RouteID *string `json:"RouteID,omitempty" name:"RouteID"`
	// 四层匹配的目的端口

	DestinationPorts []*uint64 `json:"DestinationPorts,omitempty" name:"DestinationPorts"`
}

func (r *ModifyCloudNativeAPIGatewayRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ingress struct {

	// k8s集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// Ingress的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// Ingress的数据

	IngressData *string `json:"IngressData,omitempty" name:"IngressData"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type KongServiceDetail struct {

	// 服务名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 后端协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 后端路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 后端延时，单位ms

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 重试次数

	Retries *int64 `json:"Retries,omitempty" name:"Retries"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 后端配置

	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitempty" name:"UpstreamInfo"`
	// 后端类型

	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// 是否可编辑

	Editable *bool `json:"Editable,omitempty" name:"Editable"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CheckEngineSyncJobRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据同步任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 数据同步任务运行模式(Immediate：立即运行)

	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`
	// 数据同步任务接入类型(extranet：公网类型接入)

	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`
	// 源端数据库IP

	SrcDbIp *string `json:"SrcDbIp,omitempty" name:"SrcDbIp"`
	// 源端数据库Port

	SrcDbPort *int64 `json:"SrcDbPort,omitempty" name:"SrcDbPort"`
	// 源端数据库所在地域

	SrcDbRegion *string `json:"SrcDbRegion,omitempty" name:"SrcDbRegion"`
	// 源端数据库类型

	SrcDbType *string `json:"SrcDbType,omitempty" name:"SrcDbType"`
	// 源端数据库名称

	SrcDbName *string `json:"SrcDbName,omitempty" name:"SrcDbName"`
	// 源端数据库用户名

	SrcDbUser *string `json:"SrcDbUser,omitempty" name:"SrcDbUser"`
	// 源端数据库密码

	SrcDbPassword *string `json:"SrcDbPassword,omitempty" name:"SrcDbPassword"`
	// 数据库名称映射关系

	DbMappings []*KVMapping `json:"DbMappings,omitempty" name:"DbMappings"`
	// 任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CheckEngineSyncJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckEngineSyncJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNativeGatewayServiceGroupRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 命名空间

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 服务列表

	Service *string `json:"Service,omitempty" name:"Service"`
	// 实例分组名称

	Group *string `json:"Group,omitempty" name:"Group"`
	// 实例标签

	Labels []*InstanceLabel `json:"Labels,omitempty" name:"Labels"`
}

func (r *CreateNativeGatewayServiceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNativeGatewayServiceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateWhiteListItem struct {

	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 白名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeRegionInfosRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewaySystemParametersRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关实例系统参数请求信息列表。

	SystemParameterList []*CloudNativeAPIGatewaySystemParameterReq `json:"SystemParameterList,omitempty" name:"SystemParameterList"`
	// 目前只支持kong.conf

	ConfigKey *string `json:"ConfigKey,omitempty" name:"ConfigKey"`
	// 系统参数的JSON格式字符串

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
}

func (r *ModifyCloudNativeAPIGatewaySystemParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewaySystemParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGovernanceAliasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGovernanceAliasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApolloEnvReplicasRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据大小限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 第几页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 环境名

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
}

func (r *DescribeApolloEnvReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApolloEnvReplicasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigSpec struct {

	// 规格ID

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 规格名称，例如：标准型

	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`
	// 规则升级优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 规格具体描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 存储容量下限

	MinLimit *int64 `json:"MinLimit,omitempty" name:"MinLimit"`
	// 存储容量上限

	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
	// 是否在白名单内

	InWhiteList *bool `json:"InWhiteList,omitempty" name:"InWhiteList"`
}

type DescribeCloudNativeAPIGatewayUserStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否是自研用户。

		IsYunti *bool `json:"IsYunti,omitempty" name:"IsYunti"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayUserStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartSREInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartSREInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartSREInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BoundK8SInfo struct {

	// 绑定的kubernetes集群ID

	BoundClusterId *string `json:"BoundClusterId,omitempty" name:"BoundClusterId"`
	// 绑定的kubernetes的集群类型，分tke和eks两种

	BoundClusterType *string `json:"BoundClusterType,omitempty" name:"BoundClusterType"`
	// 服务同步模式，all为全量同步，demand为按需同步

	SyncMode *string `json:"SyncMode,omitempty" name:"SyncMode"`
}

type CreateCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 服务名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 访问后端协议，http或https

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问后端的地址

	Path *string `json:"Path,omitempty" name:"Path"`
	// 后端超时时间，单位ms

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 后端重试次数

	Retries *int64 `json:"Retries,omitempty" name:"Retries"`
	// 后端类型

	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// 后端配置

	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitempty" name:"UpstreamInfo"`
}

func (r *CreateCloudNativeAPIGatewayServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNativeGatewayServicesRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayID *string `json:"GatewayID,omitempty" name:"GatewayID"`
	// 服务来源实例ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 命名空间信息

	Namespace *NamespaceInfo `json:"Namespace,omitempty" name:"Namespace"`
	// 服务列表

	Services []*string `json:"Services,omitempty" name:"Services"`
}

func (r *CreateNativeGatewayServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNativeGatewayServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopEngineSyncJobRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据同步任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StopEngineSyncJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopEngineSyncJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREInstanceAccessAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内网访问地址

		IntranetAddress *string `json:"IntranetAddress,omitempty" name:"IntranetAddress"`
		// 公网访问地址

		InternetAddress *string `json:"InternetAddress,omitempty" name:"InternetAddress"`
		// apollo多环境公网ip

		EnvAddressInfos []*EnvAddressInfo `json:"EnvAddressInfos,omitempty" name:"EnvAddressInfos"`
		// 控制台公网访问地址

		ConsoleInternetAddress *string `json:"ConsoleInternetAddress,omitempty" name:"ConsoleInternetAddress"`
		// 控制台内网访问地址

		ConsoleIntranetAddress *string `json:"ConsoleIntranetAddress,omitempty" name:"ConsoleIntranetAddress"`
		// 客户端公网带宽

		InternetBandWidth *int64 `json:"InternetBandWidth,omitempty" name:"InternetBandWidth"`
		// 控制台公网带宽

		ConsoleInternetBandWidth *int64 `json:"ConsoleInternetBandWidth,omitempty" name:"ConsoleInternetBandWidth"`
		// 北极星限流server节点接入IP

		LimiterAddressInfos []*PolarisLimiterAddress `json:"LimiterAddressInfos,omitempty" name:"LimiterAddressInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSREInstanceAccessAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSREInstanceAccessAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApolloReplicaPodDetail struct {

	// 节点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 节点所属的 zone 信息

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type Filter struct {

	// 过滤参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤参数值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type OperatePodItem struct {

	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// pod运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// xx

	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`
	// pod启动时间

	PodStartTime *string `json:"PodStartTime,omitempty" name:"PodStartTime"`
	// pod重启次数

	PodRestartCount *int64 `json:"PodRestartCount,omitempty" name:"PodRestartCount"`
	// 安全组ID

	SecretGroupName *string `json:"SecretGroupName,omitempty" name:"SecretGroupName"`
}

type CreateCloudNativeAPIGatewayRouteRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 所属服务的ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 路由的名字，实例级别唯一，可以不提供

	RouteName *string `json:"RouteName,omitempty" name:"RouteName"`
	// 路由的方法

	Methods []*string `json:"Methods,omitempty" name:"Methods"`
	// 路由的域名

	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
	// 路由的路径

	Paths []*string `json:"Paths,omitempty" name:"Paths"`
	// 路由的协议，可选https, http

	Protocols []*string `json:"Protocols,omitempty" name:"Protocols"`
	// 转发到后端时是否保留Host

	PreserveHost *bool `json:"PreserveHost,omitempty" name:"PreserveHost"`
	// https重定向状态码

	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitempty" name:"HttpsRedirectStatusCode"`
	// 转发到后端时是否StripPath

	StripPath *bool `json:"StripPath,omitempty" name:"StripPath"`
	// 是否开启强制HTTPS

	ForceHttps *bool `json:"ForceHttps,omitempty" name:"ForceHttps"`
	// 四层匹配的目的端口

	DestinationPorts []*uint64 `json:"DestinationPorts,omitempty" name:"DestinationPorts"`
}

func (r *CreateCloudNativeAPIGatewayRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalerResourceStrategiesRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 策略ID

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DescribeAutoScalerResourceStrategiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoScalerResourceStrategiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateApolloDeploymentParams struct {

	// 更新的镜像地址

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

type DescribeNativeGatewayServiceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 网关服务实例分组列表

		List []*NativeGatewayServiceGroupItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNativeGatewayServiceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNativeGatewayServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceRegionInfo struct {

	// 引擎部署地域信息

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 引擎在该地域的副本数

	Replica *int64 `json:"Replica,omitempty" name:"Replica"`
	// 引擎在该地域的规格id

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 内网的网络信息

	IntranetVpcInfos []*VpcInfo `json:"IntranetVpcInfos,omitempty" name:"IntranetVpcInfos"`
	// 是否开公网

	EnableClientInternet *bool `json:"EnableClientInternet,omitempty" name:"EnableClientInternet"`
}

type ManageVpcEndPointRequest struct {
	*tchttp.BaseRequest

	// 命令（create、delete）

	Command *string `json:"Command,omitempty" name:"Command"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Apollo环境名称

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
	// 引擎相关组件名称（pushgateway）

	Workload *string `json:"Workload,omitempty" name:"Workload"`
	// 部署地域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 打通的服务类型。
	// apollo控制台内网：console-intranet

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 指定打通后的vip地址

	EndpointVip *string `json:"EndpointVip,omitempty" name:"EndpointVip"`
}

func (r *ManageVpcEndPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageVpcEndPointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEngineSpecRequest struct {
	*tchttp.BaseRequest

	// 引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 引擎规格，暂不支持同时变更规格和节点数
	// nacos、zookeeper、apollo、consul、eureka参考值：
	// - 1C2G
	// - 2C4G
	// polarismesh参考值：
	// - NUM100
	// - NUM500
	//
	// 兼容原spec-xxxxxx形式的规格ID
	//

	EngineResourceSpec *string `json:"EngineResourceSpec,omitempty" name:"EngineResourceSpec"`
	// 引擎实例节点数，暂不支持同时变更规格和节点数
	// 北极星实例只有配额数，此参数无需填写

	EngineNodeNum *int64 `json:"EngineNodeNum,omitempty" name:"EngineNodeNum"`
	// apollo实例需要指定环境名

	ApolloEnvName *string `json:"ApolloEnvName,omitempty" name:"ApolloEnvName"`
}

func (r *ModifyEngineSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEngineSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSREInstanceSubnetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSREInstanceSubnetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSREInstanceSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayBalancedService struct {

	// 服务 ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// Upstream 名称

	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`
	// 百分比，10 即 10%

	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
}

type CreateNativeGatewayServerGroupRequest struct {
	*tchttp.BaseRequest

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 网关分组名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点配置

	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitempty" name:"NodeConfig"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 公网带宽信息

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 公网配置。

	InternetConfig *InternetConfig `json:"InternetConfig,omitempty" name:"InternetConfig"`
}

func (r *CreateNativeGatewayServerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNativeGatewayServerGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateRegionRequest struct {
	*tchttp.BaseRequest

	// 地域

	RegionField *string `json:"RegionField,omitempty" name:"RegionField"`
	// 地域中文

	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
	// 大区

	Area *string `json:"Area,omitempty" name:"Area"`
	// agent镜像地址

	AgentImageUrl *string `json:"AgentImageUrl,omitempty" name:"AgentImageUrl"`
	// 巴拉多地址

	BaradHost *string `json:"BaradHost,omitempty" name:"BaradHost"`
	// EKS地址

	EksHost *string `json:"EksHost,omitempty" name:"EksHost"`
	// VPC服务地址

	InnerVpcHost *string `json:"InnerVpcHost,omitempty" name:"InnerVpcHost"`
	// 顺序

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// TSE仓库地址

	TSEPublicRepoUrl *string `json:"TSEPublicRepoUrl,omitempty" name:"TSEPublicRepoUrl"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
	// 地域发布状态。
	//
	// -PRE：发布中，外部不可见
	// -PROD：线上，外不可见

	ReleaseStatus *string `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`
}

func (r *CreateOperateRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceStrategyDetailRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 鉴权策略ID

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DescribeGovernanceStrategyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceStrategyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SREInstance struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本号

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 状态, 枚举值:creating/create_fail/running/updating/update_fail/restarting/restart_fail/destroying/destroy_fail

	Status *string `json:"Status,omitempty" name:"Status"`
	// 规格ID

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 副本数

	Replica *int64 `json:"Replica,omitempty" name:"Replica"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Vpc iD

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 是否开启持久化存储

	EnableStorage *bool `json:"EnableStorage,omitempty" name:"EnableStorage"`
	// 数据存储方式

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 云硬盘容量

	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`
	// 计费方式

	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
	// EKS集群的ID

	EKSClusterID *string `json:"EKSClusterID,omitempty" name:"EKSClusterID"`
	// 集群创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 环境配置信息列表

	EnvInfos []*EnvInfo `json:"EnvInfos,omitempty" name:"EnvInfos"`
	// 引擎所在的区域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 注册引擎是否开启公网

	EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`
	// 私有网络列表信息

	VpcInfos []*VpcInfo `json:"VpcInfos,omitempty" name:"VpcInfos"`
	// 服务治理相关信息列表

	ServiceGovernanceInfos []*ServiceGovernanceInfo `json:"ServiceGovernanceInfos,omitempty" name:"ServiceGovernanceInfos"`
	// 实例的标签信息

	Tags []*KVPair `json:"Tags,omitempty" name:"Tags"`
	// 引擎实例是否开启控制台公网访问地址

	EnableConsoleInternet *bool `json:"EnableConsoleInternet,omitempty" name:"EnableConsoleInternet"`
	// 引擎实例是否开启控制台内网访问地址

	EnableConsoleIntranet *bool `json:"EnableConsoleIntranet,omitempty" name:"EnableConsoleIntranet"`
	// 引擎实例是否展示参数配置页面

	ConfigInfoVisible *bool `json:"ConfigInfoVisible,omitempty" name:"ConfigInfoVisible"`
	// 引擎实例控制台默认密码

	ConsoleDefaultPwd *string `json:"ConsoleDefaultPwd,omitempty" name:"ConsoleDefaultPwd"`
	// 交易付费类型，0后付费/1预付费

	TradeType *int64 `json:"TradeType,omitempty" name:"TradeType"`
	// 自动续费标记：0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 预付费到期时间

	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
	// 隔离开始时间

	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
	// 实例地域相关的描述信息

	RegionInfos []*DescribeInstanceRegionInfo `json:"RegionInfos,omitempty" name:"RegionInfos"`
	// 所在EKS环境，分为common和yunti

	EKSType *string `json:"EKSType,omitempty" name:"EKSType"`
	// 引擎的产品版本

	FeatureVersion *string `json:"FeatureVersion,omitempty" name:"FeatureVersion"`
	// 引擎实例是否开启客户端内网访问地址

	EnableClientIntranet *bool `json:"EnableClientIntranet,omitempty" name:"EnableClientIntranet"`
}

type CreateGovernanceUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGovernanceUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenPrometheusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenPrometheusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenPrometheusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolarisServerInterfacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口总个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 接口列表

		Content []*PolarisServerInterface `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolarisServerInterfacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolarisServerInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayPluginInfo struct {

	// 插件名称

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 插件版本

	PluginVersion *string `json:"PluginVersion,omitempty" name:"PluginVersion"`
	// 插件下载地址

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// 插件状态, NotInstalled|Installing|Installed

	Status *string `json:"Status,omitempty" name:"Status"`
	// 插件上传时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 版本描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type KongCertificatesPreview struct {

	// 证书名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 绑定的域名

	BindDomains []*string `json:"BindDomains,omitempty" name:"BindDomains"`
	// 证书状态：expired(已过期)
	//                    active(生效中)

	Status *string `json:"Status,omitempty" name:"Status"`
	// 证书pem格式

	Crt *string `json:"Crt,omitempty" name:"Crt"`
	// 证书私钥

	Key *string `json:"Key,omitempty" name:"Key"`
	// 证书过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 证书上传时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 证书签发时间

	IssueTime *string `json:"IssueTime,omitempty" name:"IssueTime"`
	// 证书来源：native(kong自定义证书)
	//                     ssl(ssl平台证书)

	CertSource *string `json:"CertSource,omitempty" name:"CertSource"`
	// ssl平台证书Id

	CertId *string `json:"CertId,omitempty" name:"CertId"`
}

type GetCloudNativeAPIGatewaySecretsRequest struct {
	*tchttp.BaseRequest

	// 网关实例id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 集群id，如果不填则查询网关所有关联的集群

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 查询的Namespace，如果不填则查询所有 Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 翻页起始索引，从 0 开始

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetCloudNativeAPIGatewaySecretsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCloudNativeAPIGatewaySecretsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigFile struct {

	// 配置文件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 配置文件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 配置文件命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 配置文件组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 配置文件内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 配置文件格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 配置文件注释

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 配置文件状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 配置文件标签数组

	Tags []*ConfigFileTag `json:"Tags,omitempty" name:"Tags"`
	// 配置文件创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 配置文件创建者

	CreateBy *string `json:"CreateBy,omitempty" name:"CreateBy"`
	// 配置文件修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 配置文件修改者

	ModifyBy *string `json:"ModifyBy,omitempty" name:"ModifyBy"`
	// 配置文件发布时间

	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
	// 配置文件发布者

	ReleaseBy *string `json:"ReleaseBy,omitempty" name:"ReleaseBy"`
}

type DeleteGovernanceStrategiesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 鉴权策略ID列表

	StrategyIds []*string `json:"StrategyIds,omitempty" name:"StrategyIds"`
}

func (r *DeleteGovernanceStrategiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGovernanceStrategiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNativeGatewayServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNativeGatewayServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNativeGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperServiceInstanceCountRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeZookeeperServiceInstanceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperServiceInstanceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayTaskPhase struct {

	// 任务执行的阶段

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 任务所处阶段执行的状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 处于当前状态的原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 处于当前状态的详细信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeCLBSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// clb类型，例如 portal

	CLBType *string `json:"CLBType,omitempty" name:"CLBType"`
	// 环境信息

	Env *string `json:"Env,omitempty" name:"Env"`
	// 地域信息

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
}

func (r *DescribeCLBSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCLBSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayIngressesRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// k8s集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 分页其实索引，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 也分大小，默认10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCloudNativeAPIGatewayIngressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayIngressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加载更多日志时使用，透传上次返回的Context值，获取后续的日志内容

		Context *string `json:"Context,omitempty" name:"Context"`
		// 日志查询结果是否全部返回

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 日志的内容

		LogResults []*LogResult `json:"LogResults,omitempty" name:"LogResults"`
		// 实时日志日志内容

		LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNacosServerInterfacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口总个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 接口列表

		Content []*NacosServerInterface `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNacosServerInterfacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosServerInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallCloudNativeAPIGatewayPluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallCloudNativeAPIGatewayPluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallCloudNativeAPIGatewayPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateWhiteListRequest struct {
	*tchttp.BaseRequest

	// 用户uin

	UinField *string `json:"UinField,omitempty" name:"UinField"`
	// 白名单类型：apollo | region

	Type *string `json:"Type,omitempty" name:"Type"`
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南
	// - NorthAmerica，北美

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *CreateOperateWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCLBSecurityGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定的白名单信息

		WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCLBSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCLBSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperMigrateTseNodeInfo struct {

	// 节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 客户端端口

	ClientPort *int64 `json:"ClientPort,omitempty" name:"ClientPort"`
	// 节点状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 事务ID

	ZXID *string `json:"ZXID,omitempty" name:"ZXID"`
	// 差值

	Diff *int64 `json:"Diff,omitempty" name:"Diff"`
	// zookeeper的myid

	MYID *string `json:"MYID,omitempty" name:"MYID"`
	// 选举端口

	ElectionPort *int64 `json:"ElectionPort,omitempty" name:"ElectionPort"`
	// 集群端口

	ClusterPort *int64 `json:"ClusterPort,omitempty" name:"ClusterPort"`
}

type DeleteAutoScalerResourceStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAutoScalerResourceStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoScalerResourceStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceStrategyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 鉴权策略详细

		Strategy *AuthStrategy `json:"Strategy,omitempty" name:"Strategy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceStrategyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceStrategyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCloudNativeAPIGatewayServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCloudNativeAPIGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建云原生API网关实例响应结果。

		Result *CreateCloudNativeAPIGatewayResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCloudNativeAPIGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNativeGatewayServerGroupsRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 翻页从第几个开始获取

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 翻页获取多少个

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNativeGatewayServerGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNativeGatewayServerGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对配置中心操作配置之后的返回值

		Result *string `json:"Result,omitempty" name:"Result"`
		// 操作是否成功

		OpResult *bool `json:"OpResult,omitempty" name:"OpResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNacosConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例系统参数列表

	NacosConfigInfoList []*NacosConfigInfo `json:"NacosConfigInfoList,omitempty" name:"NacosConfigInfoList"`
}

func (r *UpdateNacosConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNacosConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DescribeCloudNativeAPIGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudNativeAPIGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEksMetricsRequest struct {
	*tchttp.BaseRequest

	// 空间唯一id

	SpaceUUID *string `json:"SpaceUUID,omitempty" name:"SpaceUUID"`
	// 查询详情

	Query []*Query `json:"Query,omitempty" name:"Query"`
	// 引擎所在区域

	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeEksMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEksMetricsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolarisBindK8sClustersRequest struct {
	*tchttp.BaseRequest

	// 引擎实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePolarisBindK8sClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolarisBindK8sClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNacosAdminPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否使用初始密码

	UseInitPassword *bool `json:"UseInitPassword,omitempty" name:"UseInitPassword"`
	// 自定义密码（与恢复初始密码二选一）

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
}

func (r *ModifyNacosAdminPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNacosAdminPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayConfigResult struct {

	// 云原生API网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关配置列表。

	ConfigList []*CloudNativeAPIGatewayConfig `json:"ConfigList,omitempty" name:"ConfigList"`
}

type NginxIngressConfigMapParams struct {

	// 文档地址

	Doc *string `json:"Doc,omitempty" name:"Doc"`
	// 生效的参数列表

	EnabledParams []*NginxIngressConfigMapParam `json:"EnabledParams,omitempty" name:"EnabledParams"`
	// 不生效的参数列表

	DisabledParams []*NginxIngressConfigMapParam `json:"DisabledParams,omitempty" name:"DisabledParams"`
}

type DescribeNacosReplicasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎实例副本信息

		Replicas []*NacosReplica `json:"Replicas,omitempty" name:"Replicas"`
		// 副本个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNacosReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNacosReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartSREInstanceRequest struct {
	*tchttp.BaseRequest

	// 微服务引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 重启的环境类型（PROD，DEV，UAT等）

	EnvTypes []*string `json:"EnvTypes,omitempty" name:"EnvTypes"`
}

func (r *RestartSREInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartSREInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvInfo struct {

	// 环境名称

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
	// 环境对应的网络信息

	VpcInfos []*VpcInfo `json:"VpcInfos,omitempty" name:"VpcInfos"`
	// 运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 云硬盘容量

	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`
	// Admin service 访问地址

	AdminServiceIp *string `json:"AdminServiceIp,omitempty" name:"AdminServiceIp"`
	// Config service访问地址

	ConfigServiceIp *string `json:"ConfigServiceIp,omitempty" name:"ConfigServiceIp"`
	// 是否开启config-server公网

	EnableConfigInternet *bool `json:"EnableConfigInternet,omitempty" name:"EnableConfigInternet"`
	// config-server公网访问地址

	ConfigInternetServiceIp *string `json:"ConfigInternetServiceIp,omitempty" name:"ConfigInternetServiceIp"`
	// 规格ID

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 环境的节点数

	EnvReplica *int64 `json:"EnvReplica,omitempty" name:"EnvReplica"`
	// 环境运行的节点数

	RunningCount *int64 `json:"RunningCount,omitempty" name:"RunningCount"`
	// 环境别名

	AliasEnvName *string `json:"AliasEnvName,omitempty" name:"AliasEnvName"`
	// 环境描述

	EnvDesc *string `json:"EnvDesc,omitempty" name:"EnvDesc"`
	// 客户端带宽

	ClientBandWidth *uint64 `json:"ClientBandWidth,omitempty" name:"ClientBandWidth"`
}

type CreateExecTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExecTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExecTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MergeZookeeperMigrateClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MergeZookeeperMigrateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MergeZookeeperMigrateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGovernanceUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceGroupTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组

		Group *UserGroup `json:"Group,omitempty" name:"Group"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceGroupTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceGroupTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGovernanceAliasesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务别名总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务别名列表。

		Content []*GovernanceAlias `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGovernanceAliasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGovernanceAliasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZookeeperDigestTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 摘要信息

		Digest *string `json:"Digest,omitempty" name:"Digest"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZookeeperDigestTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZookeeperDigestTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigFileReleaseInput struct {

	// 配置文件发布名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 配置文件发布命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 配置文件发布组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 配置文件发布文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 配置文件发布内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 配置文件发布注释

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 配置文件发布版本

	Version *uint64 `json:"Version,omitempty" name:"Version"`
}

type RegionInfo struct {

	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 大区

	Area *string `json:"Area,omitempty" name:"Area"`
	// 区域中文名称

	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
}

type CreateGovernanceServicesRequest struct {
	*tchttp.BaseRequest

	// tse 实例 id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务信息。

	GovernanceServices []*GovernanceServiceInput `json:"GovernanceServices,omitempty" name:"GovernanceServices"`
}

func (r *CreateGovernanceServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayIngressRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// k8s集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// Ingress的名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteCloudNativeAPIGatewayIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCloudNativeAPIGatewayIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGovernanceInstancesRequest struct {
	*tchttp.BaseRequest

	// tse实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务实例信息。

	GovernanceInstances []*GovernanceInstanceInput `json:"GovernanceInstances,omitempty" name:"GovernanceInstances"`
}

func (r *CreateGovernanceInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGovernanceInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceStatusListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeServiceStatusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceStatusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayClsConfig struct {

	// CLS日志集ID。

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// CLS主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志类型, KongErrorLog|KongAccessLog。

	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

type TaskPhaseName struct {

	// 执行阶段

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 中文名称

	ChineseName *string `json:"ChineseName,omitempty" name:"ChineseName"`
}

type DescribeConfigFileGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置文件组列表

		ConfigFileGroups []*ConfigFileGroup `json:"ConfigFileGroups,omitempty" name:"ConfigFileGroups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigFileGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigFileGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateFeatureVersionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 功能版本详情

		List []*OperateFeatureVersion `json:"List,omitempty" name:"List"`
		// 列表总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateFeatureVersionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateFeatureVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCloudNativeAPIGatewayPluginUploadInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Cos临时秘钥信息

		Result *CosTokenInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCloudNativeAPIGatewayPluginUploadInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCloudNativeAPIGatewayPluginUploadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigFileGroupRequest struct {
	*tchttp.BaseRequest

	// tse实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置文件组

	ConfigFileGroup *ConfigFileGroup `json:"ConfigFileGroup,omitempty" name:"ConfigFileGroup"`
}

func (r *ModifyConfigFileGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigFileGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGovernanceUserPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户Id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 旧密码

	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`
	// 新密码

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
}

func (r *ModifyGovernanceUserPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGovernanceUserPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObtainApolloAdminTokenRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ObtainApolloAdminTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ObtainApolloAdminTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCloudNativeAPIGatewayPluginUploadInfoRequest struct {
	*tchttp.BaseRequest

	// 云原生API网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 插件名称

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
}

func (r *GetCloudNativeAPIGatewayPluginUploadInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCloudNativeAPIGatewayPluginUploadInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoScalerResourceStrategyRequest struct {
	*tchttp.BaseRequest

	// 网关实例ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 策略ID

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 弹性伸缩配置

	Config *CloudNativeAPIGatewayStrategyAutoScalerConfig `json:"Config,omitempty" name:"Config"`
}

func (r *ModifyAutoScalerResourceStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoScalerResourceStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConsulConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 配置项Key

	ConfigKey *string `json:"ConfigKey,omitempty" name:"ConfigKey"`
	// 配置项Value，json字符串格式

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
}

func (r *UpdateConsulConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateConsulConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NacosServerInterface struct {

	// 接口名

	Interface *string `json:"Interface,omitempty" name:"Interface"`
}

type DescribeExecTaskBatchJobInfoRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 批次索引

	BatchIndex *uint64 `json:"BatchIndex,omitempty" name:"BatchIndex"`
	// 主控地域参数，支持的值如下。
	// - China，中国区
	// - Europe，欧洲
	// - SoutheastAsia，亚太东南

	MainRegion *string `json:"MainRegion,omitempty" name:"MainRegion"`
}

func (r *DescribeExecTaskBatchJobInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExecTaskBatchJobInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSREGlobalConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSREGlobalConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSREGlobalConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewayCertificatesDomainRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 列表数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 列表offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListCloudNativeAPIGatewayCertificatesDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudNativeAPIGatewayCertificatesDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClosePrometheusRequest struct {
	*tchttp.BaseRequest

	// 网关ID

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *ClosePrometheusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClosePrometheusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEksMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标内容详情

		Data []*Data `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEksMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEksMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySREInstanceSpecRequest struct {
	*tchttp.BaseRequest

	// 微服务引擎实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 微服务引擎规格Id

	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`
	// 微服务引擎实例副本数

	Replica *int64 `json:"Replica,omitempty" name:"Replica"`
}

func (r *ModifySREInstanceSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySREInstanceSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
