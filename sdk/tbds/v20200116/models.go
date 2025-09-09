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

package v20200116

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type GetProgressRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterExistsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群名称是否存在

		Data *bool `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClusterExistsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClusterExistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClustersRequest struct {
	*tchttp.BaseRequest

	// 从0开始的页面index

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页面大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 关键字，支持|分隔多个关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 类型，填充“cvm”或“bms”

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *GetClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息详情

		Data *Cluster `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodesRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 从0开始的页面index

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页面大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeClusterNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInstallInfo struct {

	// ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// tbds集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 操作类别

	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
	// 步骤状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 其他

	Extra *string `json:"Extra,omitempty" name:"Extra"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type Sort struct {

	// 是否为空

	Empty *bool `json:"Empty,omitempty" name:"Empty"`
	// 是否排序

	Sorted *bool `json:"Sorted,omitempty" name:"Sorted"`
	// 是否无排序

	Unsorted *bool `json:"Unsorted,omitempty" name:"Unsorted"`
}

type GetInstallStepsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群部署进度

		Data *ClusterInstallResponse `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstallStepsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstallStepsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInstallDetail struct {

	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 集群信息ID

	ClusterInfoId *uint64 `json:"ClusterInfoId,omitempty" name:"ClusterInfoId"`
	// 安装步骤

	InstallStep *string `json:"InstallStep,omitempty" name:"InstallStep"`
	// 安装步骤中文名

	InstallStepName *string `json:"InstallStepName,omitempty" name:"InstallStepName"`
	// 步骤状态

	InstallStatus *uint64 `json:"InstallStatus,omitempty" name:"InstallStatus"`
	// 其他

	Extra *string `json:"Extra,omitempty" name:"Extra"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeClusterNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群机器信息分页对象

		Data *PageClusterNode `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInstallInfoRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 记录id

	ClusterInstallId *string `json:"ClusterInstallId,omitempty" name:"ClusterInstallId"`
}

func (r *GetInstallInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstallInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInstallStepsRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 安装id

	ClusterInstallId *string `json:"ClusterInstallId,omitempty" name:"ClusterInstallId"`
}

func (r *GetInstallStepsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstallStepsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// 数据盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同。默认值为0，表示不购买数据盘。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 数据盘是否随子机销毁。TRUE：子机销毁时，销毁数据盘，只支持按小时后付费云盘。FALSE：子机销毁时，保留数据盘。默认取值：TRUE

	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}

type Pageable struct {

	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 页码

	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页显示条数

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否分页

	Paged *bool `json:"Paged,omitempty" name:"Paged"`
	// 排序信息

	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
	// 是否未分页

	Unpaged *bool `json:"Unpaged,omitempty" name:"Unpaged"`
}

type BmsDescribeInstancesByNameRequest struct {
	*tchttp.BaseRequest

	// 实例名称数组

	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`
}

func (r *BmsDescribeInstancesByNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BmsDescribeInstancesByNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BmsInstanceInfo struct {

	// bms机型

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 要申请创建的BMS实例个数

	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// RAID类型存储类型

	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`
}

type Cluster struct {

	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// Oauth相关

	MasterOauthClientId *string `json:"MasterOauthClientId,omitempty" name:"MasterOauthClientId"`
	// Oauth相关

	MasterOauthClientSecret *string `json:"MasterOauthClientSecret,omitempty" name:"MasterOauthClientSecret"`
	// Oauth相关

	MasterOauthDetail *string `json:"MasterOauthDetail,omitempty" name:"MasterOauthDetail"`
	// Oauth相关

	SlaveOauthClientId *string `json:"SlaveOauthClientId,omitempty" name:"SlaveOauthClientId"`
	// Oauth相关

	SlaveOauthClientSecret *string `json:"SlaveOauthClientSecret,omitempty" name:"SlaveOauthClientSecret"`
	// Oauth相关

	SlaveOauthDetail *string `json:"SlaveOauthDetail,omitempty" name:"SlaveOauthDetail"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 结束时间

	CompletionTime *string `json:"CompletionTime,omitempty" name:"CompletionTime"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 创建人ID

	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 租户ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// tbds版本

	TbdsVersion *string `json:"TbdsVersion,omitempty" name:"TbdsVersion"`
	// 架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 实例数

	InstanceCount *string `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// cvm状态

	CvmStatus *string `json:"CvmStatus,omitempty" name:"CvmStatus"`
	// tbds状态

	TbdsStatus *string `json:"TbdsStatus,omitempty" name:"TbdsStatus"`
	// 实例ID

	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// IP信息

	ClusterPrivateIps *string `json:"ClusterPrivateIps,omitempty" name:"ClusterPrivateIps"`
	// EIP信息

	MasterEipAddress *string `json:"MasterEipAddress,omitempty" name:"MasterEipAddress"`
	// EIP信息

	SlaveEipAddress *string `json:"SlaveEipAddress,omitempty" name:"SlaveEipAddress"`
	// 主节点ID

	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`
	// 从节点ID

	SlaveInstanceId *string `json:"SlaveInstanceId,omitempty" name:"SlaveInstanceId"`
	// 数量

	LocalStorageDiskCount *uint64 `json:"LocalStorageDiskCount,omitempty" name:"LocalStorageDiskCount"`
	// ssh信息

	MasterSshAddress *string `json:"MasterSshAddress,omitempty" name:"MasterSshAddress"`
	// ssh信息

	SlaveSshAddress *string `json:"SlaveSshAddress,omitempty" name:"SlaveSshAddress"`
	// 主节点信息

	MasterPortalApiAddress *string `json:"MasterPortalApiAddress,omitempty" name:"MasterPortalApiAddress"`
	// 从节点信息

	SlavePortalApiAddress *string `json:"SlavePortalApiAddress,omitempty" name:"SlavePortalApiAddress"`
	// 创建请求序列化

	ClusterRequest *string `json:"ClusterRequest,omitempty" name:"ClusterRequest"`
	// 销毁时间

	DestroyTime *string `json:"DestroyTime,omitempty" name:"DestroyTime"`
	// jwt

	AllJnsgwAccessPolicyInfoStr *string `json:"AllJnsgwAccessPolicyInfoStr,omitempty" name:"AllJnsgwAccessPolicyInfoStr"`
	// 说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 异常日志

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// 错误值

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
	// EIP状态

	MasterEipStatus *string `json:"MasterEipStatus,omitempty" name:"MasterEipStatus"`
	// EIP状态

	SlaveEipStatus *string `json:"SlaveEipStatus,omitempty" name:"SlaveEipStatus"`
	// 扩容状态

	ExpandCvmStatus *string `json:"ExpandCvmStatus,omitempty" name:"ExpandCvmStatus"`
	// 阶段下标

	InstallStepNum *uint64 `json:"InstallStepNum,omitempty" name:"InstallStepNum"`
	// 阶段信息

	InstallStepMsg *string `json:"InstallStepMsg,omitempty" name:"InstallStepMsg"`
	// 节点数

	NodeSize *uint64 `json:"NodeSize,omitempty" name:"NodeSize"`
	// 扩容中实例

	ExpandingCvmInstances *string `json:"ExpandingCvmInstances,omitempty" name:"ExpandingCvmInstances"`
	// 主url

	MasterLoginUrl *string `json:"MasterLoginUrl,omitempty" name:"MasterLoginUrl"`
	// 从url

	SlaveLoginUrl *string `json:"SlaveLoginUrl,omitempty" name:"SlaveLoginUrl"`
	// portal节点数

	PortalNodeSize *uint64 `json:"PortalNodeSize,omitempty" name:"PortalNodeSize"`
	// master节点数

	MasterNodeSize *uint64 `json:"MasterNodeSize,omitempty" name:"MasterNodeSize"`
	// worker节点数

	WorkerNodeSize *uint64 `json:"WorkerNodeSize,omitempty" name:"WorkerNodeSize"`
	// cvm状态

	CvmStateMachine *string `json:"CvmStateMachine,omitempty" name:"CvmStateMachine"`
}

type GetSystemArchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统架构

		Data *string `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSystemArchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSystemArchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestApplyBmsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestApplyBmsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestApplyBmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BmsCreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// Docker镜像ID

	CvmImageId *string `json:"CvmImageId,omitempty" name:"CvmImageId"`
	// 实例所属可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// Bms子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 集群描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// master 节点参数

	MasterInstanceInfo *BmsInstanceInfo `json:"MasterInstanceInfo,omitempty" name:"MasterInstanceInfo"`
	// portal 节点参数

	PortalInstanceInfo *InstanceInfo `json:"PortalInstanceInfo,omitempty" name:"PortalInstanceInfo"`
	// worker 节点参数

	WorkerInstanceInfo *BmsInstanceInfo `json:"WorkerInstanceInfo,omitempty" name:"WorkerInstanceInfo"`
	// cvm的子网id

	CvmSubnetId *string `json:"CvmSubnetId,omitempty" name:"CvmSubnetId"`
	// 操作系统名称，例如：Centos 7.5 for x86_64

	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// TBDS版本号

	TbdsVersion *string `json:"TbdsVersion,omitempty" name:"TbdsVersion"`
	// 系统架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
}

func (r *BmsCreateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BmsCreateClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BmsCreateClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		Data *string `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BmsCreateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BmsCreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		Data *string `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpandClusterNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果信息

		Data *bool `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExpandClusterNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpandClusterNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {

	// CVM实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 要申请创建的CVM实例个数

	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 系统盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘大小，单位：GB。默认值为 50

	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
	// 本地磁盘的个数

	LocalStorageDiskCount *uint64 `json:"LocalStorageDiskCount,omitempty" name:"LocalStorageDiskCount"`
	// 数据盘

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 置放群组id，仅支持指定一个。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
}

type GetProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群部署进度

		Data *ClusterProgressVO `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterExistsRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ClusterExistsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClusterExistsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpandClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 扩容结果

		Data *string `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExpandClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpandClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterNode struct {

	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例信息Json

	InstanceJson *string `json:"InstanceJson,omitempty" name:"InstanceJson"`
	// 实例分工

	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
	// 实例类型

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 实例信息

	Instance *string `json:"Instance,omitempty" name:"Instance"`
}

type InternetAccessible struct {

	// 是否分配公网ip

	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
	// 带宽大小

	InternetMaxBandwidthOut *string `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 外网供应商

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
}

type ExpandClusterNewRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 扩容master 节点个数

	ExpandMasterCount *uint64 `json:"ExpandMasterCount,omitempty" name:"ExpandMasterCount"`
	// 扩容worker 节点个数

	ExpandWorkerCount *uint64 `json:"ExpandWorkerCount,omitempty" name:"ExpandWorkerCount"`
}

func (r *ExpandClusterNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpandClusterNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群分页对象

		Data *PageCluster `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSystemArchRequest struct {
	*tchttp.BaseRequest
}

func (r *GetSystemArchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSystemArchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PageClusterNode struct {

	// 集群节点列表

	Content []*ClusterNode `json:"Content,omitempty" name:"Content"`
	// 是否为空

	Empty *bool `json:"Empty,omitempty" name:"Empty"`
	// 是否第一页

	First *bool `json:"First,omitempty" name:"First"`
	// 是否最后一页

	Last *bool `json:"Last,omitempty" name:"Last"`
	// 页码

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 列表返回条数

	NumberOfElements *int64 `json:"NumberOfElements,omitempty" name:"NumberOfElements"`
	// 分页信息

	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`
	// 每页显示条数

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// 排序信息

	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
	// 总数据条数

	TotalElements *int64 `json:"TotalElements,omitempty" name:"TotalElements"`
	// 总页数

	TotalPages *int64 `json:"TotalPages,omitempty" name:"TotalPages"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// Docker镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 实例所属可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 集群描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// master 节点参数

	MasterInstanceInfo *InstanceInfo `json:"MasterInstanceInfo,omitempty" name:"MasterInstanceInfo"`
	// portal 节点参数

	PortalInstanceInfo *InstanceInfo `json:"PortalInstanceInfo,omitempty" name:"PortalInstanceInfo"`
	// worker 节点参数

	WorkerInstanceInfo *InstanceInfo `json:"WorkerInstanceInfo,omitempty" name:"WorkerInstanceInfo"`
	// 置放群组id，仅支持指定一个。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
	// TBDS版本号

	TbdsVersion *string `json:"TbdsVersion,omitempty" name:"TbdsVersion"`
	// 系统架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
}

func (r *CreateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyClusterRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DestroyClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInstallInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群部署基本信息

		Data []*ClusterInstallInfo `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstallInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstallInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInstallResponse struct {

	// 当前步骤

	CurrentStepName *string `json:"CurrentStepName,omitempty" name:"CurrentStepName"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 执行步骤信息

	InstallStepMsg *string `json:"InstallStepMsg,omitempty" name:"InstallStepMsg"`
	// 错误原因

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
	// 异常日志

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// 安装步骤详情列表

	ClusterInstallDetail []*ClusterInstallDetail `json:"ClusterInstallDetail,omitempty" name:"ClusterInstallDetail"`
}

type BmsDescribeInstancesByNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群实例

		Data *string `json:"Data,omitempty" name:"Data"`
		// 调用结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 返回

		Ret *int64 `json:"Ret,omitempty" name:"Ret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BmsDescribeInstancesByNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BmsDescribeInstancesByNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterProgressVO struct {

	// 进展

	Progress *string `json:"Progress,omitempty" name:"Progress"`
	// 集群状态

	TbdsStatus *string `json:"TbdsStatus,omitempty" name:"TbdsStatus"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 完成时间

	CompletionTime *string `json:"CompletionTime,omitempty" name:"CompletionTime"`
}

type PageCluster struct {

	// 集群信息列表

	Content []*Cluster `json:"Content,omitempty" name:"Content"`
	// 是否为空

	Empty *bool `json:"Empty,omitempty" name:"Empty"`
	// 是否第一页

	First *bool `json:"First,omitempty" name:"First"`
	// 是否最后一页

	Last *bool `json:"Last,omitempty" name:"Last"`
	// 页码

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 列表返回条数

	NumberOfElements *int64 `json:"NumberOfElements,omitempty" name:"NumberOfElements"`
	// 分页信息

	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`
	// 每页显示条数

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// 排序信息

	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
	// 总数据条数

	TotalElements *int64 `json:"TotalElements,omitempty" name:"TotalElements"`
	// 总页数

	TotalPages *int64 `json:"TotalPages,omitempty" name:"TotalPages"`
}

type TestApplyBmsRequest struct {
	*tchttp.BaseRequest

	// 集群的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 实例所属可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 集群描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 主机密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 服务器内网ip数组

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 公网带宽信息

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 实例个数

	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 套餐id

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// bms,系统类型，例如：Linux

	OperatingSystemType *string `json:"OperatingSystemType,omitempty" name:"OperatingSystemType"`
	// bms,系统名称

	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// bms,开启主机安全服务

	EnhancedService *bool `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// RAID类型存储类型

	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`
}

func (r *TestApplyBmsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestApplyBmsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpandClusterRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 扩容master 节点个数

	ExpandMasterCount *uint64 `json:"ExpandMasterCount,omitempty" name:"ExpandMasterCount"`
	// 扩容worker 节点个数

	ExpandWorkerCount *uint64 `json:"ExpandWorkerCount,omitempty" name:"ExpandWorkerCount"`
}

func (r *ExpandClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpandClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SSLVpnServer struct {

	// vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// ssl vpn server id

	VpnSslServerId *string `json:"VpnSslServerId,omitempty" name:"VpnSslServerId"`
	// vpngw id

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// SSLVpnServerName

	SSLVpnServerName *string `json:"SSLVpnServerName,omitempty" name:"SSLVpnServerName"`
	// 本段地址

	LocalAddress *string `json:"LocalAddress,omitempty" name:"LocalAddress"`
	// 对端地址

	RemoteAddress *string `json:"RemoteAddress,omitempty" name:"RemoteAddress"`
	// 连接数

	MaxConnection *int64 `json:"MaxConnection,omitempty" name:"MaxConnection"`
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}
