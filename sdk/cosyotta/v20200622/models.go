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

package v20200622

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ListOperationTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterStage1Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 错误消息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NewClusterStage1Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterStage1Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetClusterStageRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *ResetClusterStageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetClusterStageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {

	// ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 删除时间

	DeletedAt *string `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群Region

	Region *string `json:"Region,omitempty" name:"Region"`
	// 集群位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// ServerBusi1

	ServerBusi1 *string `json:"ServerBusi1,omitempty" name:"ServerBusi1"`
	// ServerBusi2

	ServerBusi2 *string `json:"ServerBusi2,omitempty" name:"ServerBusi2"`
	// BusiUUID

	BusiUuid *string `json:"BusiUuid,omitempty" name:"BusiUuid"`
	// Center信息

	Center *string `json:"Center,omitempty" name:"Center"`
	// Columbus组件ModId

	ColumbusModId *uint64 `json:"ColumbusModId,omitempty" name:"ColumbusModId"`
	// Columbus组件CmdId

	ColumbusCmdId *uint64 `json:"ColumbusCmdId,omitempty" name:"ColumbusCmdId"`
	// Columbus组件Token

	ColumbusToken *int64 `json:"ColumbusToken,omitempty" name:"ColumbusToken"`
	// NodeMgr组件ModId

	NodeMgrModId *uint64 `json:"NodeMgrModId,omitempty" name:"NodeMgrModId"`
	// NodeMgr组件CmdId

	NodeMgrCmdId *uint64 `json:"NodeMgrCmdId,omitempty" name:"NodeMgrCmdId"`
	// Token

	Token *int64 `json:"Token,omitempty" name:"Token"`
	// buffer三级模块

	Buffer *string `json:"Buffer,omitempty" name:"Buffer"`
	// 自动新增模块时可以删除的模块名称

	DelModule *string `json:"DelModule,omitempty" name:"DelModule"`
	// blobnode oss上报

	OSSModelMarkBlobNode *string `json:"OSSModelMarkBlobNode,omitempty" name:"OSSModelMarkBlobNode"`
	// space oss上报 12个指标，监控系统最大只能有13个，不能和blobnode合并

	OSSModelMarkSpace *string `json:"OSSModelMarkSpace,omitempty" name:"OSSModelMarkSpace"`
	// 开区申请单id

	ClusterFormId *uint64 `json:"ClusterFormId,omitempty" name:"ClusterFormId"`
	// 集群状态： offline 未上线, online 正式上线

	Status *string `json:"Status,omitempty" name:"Status"`
	// AZ个数

	AzCount *uint64 `json:"AzCount,omitempty" name:"AzCount"`
}

type DeployNodesReq struct {

	// IP列表

	IpList []*int64 `json:"IpList,omitempty" name:"IpList"`
	// 软件版本列表

	SoftwareList []*SoftInfo `json:"SoftwareList,omitempty" name:"SoftwareList"`
}

type ListClustersRequest struct {
	*tchttp.BaseRequest
}

func (r *ListClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitHandleErrorDiskTaskRequest struct {
	*tchttp.BaseRequest

	// 磁盘信息(models.DiskInfo)序列化以后的信息

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// Cluster信息序列化以后的信息

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *SubmitHandleErrorDiskTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitHandleErrorDiskTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDeployableHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDeployableHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDeployableHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNodeVletsRequest struct {
	*tchttp.BaseRequest

	// Node Id

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 序列化以后的Cluster信息

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *GetNodeVletsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeVletsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartBalanceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 水位线

	Relaxation *float64 `json:"Relaxation,omitempty" name:"Relaxation"`
}

func (r *StartBalanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartBalanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回信息，为空

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 错误信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUpgradeChangeFormsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListUpgradeChangeFormsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUpgradeChangeFormsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbortBalanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AbortBalanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AbortBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlackNodesRequest struct {
	*tchttp.BaseRequest

	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 屏蔽时长

	Duration *string `json:"Duration,omitempty" name:"Duration"`
	// Node数组

	NodeList []*string `json:"NodeList,omitempty" name:"NodeList"`
}

func (r *BlackNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlackNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterByIdRequest struct {
	*tchttp.BaseRequest

	// 无

	ClusterId *uint64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RedoUpgradeSubTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 返回信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RedoUpgradeSubTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RedoUpgradeSubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostStatsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetHostStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryErrorDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryErrorDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryErrorDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeNodesRequest struct {
	*tchttp.BaseRequest

	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 操作人

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 升级节点请求

	UpgradeNodesRequest *UpgradeNodesReq `json:"UpgradeNodesRequest,omitempty" name:"UpgradeNodesRequest"`
}

func (r *UpgradeNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSpaceReq struct {

	// SpaceId

	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`
	// SpaceSpec

	SpaceSpec *string `json:"SpaceSpec,omitempty" name:"SpaceSpec"`
	// SpaceCoding

	SpaceCoding *string `json:"SpaceCoding,omitempty" name:"SpaceCoding"`
	// SpaceStatus

	SpaceStatus *string `json:"SpaceStatus,omitempty" name:"SpaceStatus"`
	// SpacePermission

	SpacePermission *string `json:"SpacePermission,omitempty" name:"SpacePermission"`
}

type ListHostsReq struct {

	// 集群Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 所在Region

	Region *string `json:"Region,omitempty" name:"Region"`
	// Ip列表

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 所在IDC

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// module

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 交换机

	SwitchName *string `json:"SwitchName,omitempty" name:"SwitchName"`
	// 机架

	RackName *string `json:"RackName,omitempty" name:"RackName"`
	// 排序顺序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序

	DescOrder *bool `json:"DescOrder,omitempty" name:"DescOrder"`
	// 机型名称

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type GetAvailabilityTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAvailabilityTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAvailabilityTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KickDiskRequest struct {
	*tchttp.BaseRequest

	// 硬盘信息序列化后的字符串

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// Cluster信息序列化后的字符串

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 备注信息

	Note *string `json:"Note,omitempty" name:"Note"`
	// 是否强制踢盘

	ForceDrop *bool `json:"ForceDrop,omitempty" name:"ForceDrop"`
}

func (r *KickDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KickDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KickDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KickDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KickDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterConfirmUpdateByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NewClusterConfirmUpdateByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterConfirmUpdateByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetChainTaskByUuidRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *GetChainTaskByUuidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetChainTaskByUuidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GinMakePackageInfoRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Module Id数组

	ModuleIds []*string `json:"ModuleIds,omitempty" name:"ModuleIds"`
}

func (r *GinMakePackageInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GinMakePackageInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollBackSubTaskRequest struct {
	*tchttp.BaseRequest

	// 升级变更表单id

	UpgradeConfFormId *uint64 `json:"UpgradeConfFormId,omitempty" name:"UpgradeConfFormId"`
	// 升级子任务表单id

	SubTaskInfoFormIds []*uint64 `json:"SubTaskInfoFormIds,omitempty" name:"SubTaskInfoFormIds"`
}

func (r *RollBackSubTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollBackSubTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLibraSummaryRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 父任务Job Id

	ParentJobId *int64 `json:"ParentJobId,omitempty" name:"ParentJobId"`
	// 任务类型

	JobType *int64 `json:"JobType,omitempty" name:"JobType"`
}

func (r *ListLibraSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLibraSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBlackNodesRequest struct {
	*tchttp.BaseRequest

	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// NodeId数组

	NodeList []*string `json:"NodeList,omitempty" name:"NodeList"`
}

func (r *UnBlackNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBlackNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTasksRequest struct {
	*tchttp.BaseRequest

	// 任务创建时间区间

	Begin *string `json:"Begin,omitempty" name:"Begin"`
	// 任务创建时间区间

	End *string `json:"End,omitempty" name:"End"`
	// 任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 任务是否已完成

	Completed *string `json:"Completed,omitempty" name:"Completed"`
	// 筛选标签,格式如: {'k':'v','k2':'v2'}

	Meta *string `json:"Meta,omitempty" name:"Meta"`
}

func (r *GetTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBalanceJobsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// CreateTime

	CreateTime []*string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Job ID

	JobId *int64 `json:"JobId,omitempty" name:"JobId"`
	// 任务状态

	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// Page页

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// Page页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否降序

	DescOrder *bool `json:"DescOrder,omitempty" name:"DescOrder"`
	// 排序方式

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListBalanceJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBalanceJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubChangeTaskRequest struct {
	*tchttp.BaseRequest

	// 升级变更任务名称

	ConfName *string `json:"ConfName,omitempty" name:"ConfName"`
}

func (r *GetSubChangeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubChangeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GinDoPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GinDoPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GinDoPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Cluster信息字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 错误消息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSpacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSpacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSpacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableNodesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// module字符串

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 均衡水位线，默认直接填0

	Relaxation *float64 `json:"Relaxation,omitempty" name:"Relaxation"`
}

func (r *EnableNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubChangeTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 返回信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 子任务表单，返回数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubChangeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubChangeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeChangesRequest struct {
	*tchttp.BaseRequest

	// 升级变更名称

	ConfName *string `json:"ConfName,omitempty" name:"ConfName"`
}

func (r *GetUpgradeChangesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeChangesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSpaceTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSpaceTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSpaceTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterConfirmUpdateByIdRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *NewClusterConfirmUpdateByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterConfirmUpdateByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RedoUpgradeSubTaskRequest struct {
	*tchttp.BaseRequest

	// 升级变更表单id

	UpgradeConfFormId *uint64 `json:"UpgradeConfFormId,omitempty" name:"UpgradeConfFormId"`
	// 升级子任务表单id

	SubTaskInfoFormIds []*uint64 `json:"SubTaskInfoFormIds,omitempty" name:"SubTaskInfoFormIds"`
}

func (r *RedoUpgradeSubTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RedoUpgradeSubTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNodeStatsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 服务ID

	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *GetNodeStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置字段（公共配置、自定义配置）返回数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 返回码

		RetCode *string `json:"RetCode,omitempty" name:"RetCode"`
		// 错误信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetColumbusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetColumbusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetColumbusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterStage3Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误消息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NewClusterStage3Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterStage3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitClusterStageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 错误消息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitClusterStageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitClusterStageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryClusterFormResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryClusterFormResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryClusterFormResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSpaceThresholdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSpaceThresholdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSpaceThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllClustersOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *AllClustersOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllClustersOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 升级分组信息，返回数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 返回码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 返回信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHostsRequest struct {
	*tchttp.BaseRequest

	// 获取Host请求

	ListHostsRequest *ListHostsReq `json:"ListHostsRequest,omitempty" name:"ListHostsRequest"`
	// 分页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 当前页

	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
}

func (r *ListHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StableVersion struct {

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 删除时间

	DeletedAt *string `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// PlatName

	PlatName *string `json:"PlatName,omitempty" name:"PlatName"`
	// ServiceName

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// ClusterId

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// VersionType

	VersionType *string `json:"VersionType,omitempty" name:"VersionType"`
	// Version

	Version *string `json:"Version,omitempty" name:"Version"`
	// VersionNumber

	VersionNumber *int64 `json:"VersionNumber,omitempty" name:"VersionNumber"`
	// NodeCount

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// MatchNodeCount

	MatchNodeCount *int64 `json:"MatchNodeCount,omitempty" name:"MatchNodeCount"`
	// Consist

	Consist *bool `json:"Consist,omitempty" name:"Consist"`
	// Note

	Note *string `json:"Note,omitempty" name:"Note"`
}

type QueryClusterFormByIdRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *QueryClusterFormByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryClusterFormByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSpaceThresholdRequest struct {
	*tchttp.BaseRequest

	// 编码策略

	CodingSchemaM *int64 `json:"CodingSchemaM,omitempty" name:"CodingSchemaM"`
	// Put阀值

	PutThreshold *int64 `json:"PutThreshold,omitempty" name:"PutThreshold"`
	// Get阀值

	GetThreshold *int64 `json:"GetThreshold,omitempty" name:"GetThreshold"`
	// Delete阀值

	DeleteThreshold *int64 `json:"DeleteThreshold,omitempty" name:"DeleteThreshold"`
	// Put失败阀值

	BadPutThreshold *int64 `json:"BadPutThreshold,omitempty" name:"BadPutThreshold"`
	// Tmp Put阀值

	TmpPutThresholdValue *int64 `json:"TmpPutThresholdValue,omitempty" name:"TmpPutThresholdValue"`
	// Tmp Put结束时间

	TmpPutThresholdEndTime *string `json:"TmpPutThresholdEndTime,omitempty" name:"TmpPutThresholdEndTime"`
}

func (r *UpdateSpaceThresholdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSpaceThresholdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterConfirmRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *NewClusterConfirmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterConfirmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLibraJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLibraJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLibraJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLibraSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLibraSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLibraSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationTasksRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 任务类型

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 排序方式

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序

	DescOrder *bool `json:"DescOrder,omitempty" name:"DescOrder"`
}

func (r *ListOperationTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnregisteredNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnregisteredNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnregisteredNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSpacePermissionRequest struct {
	*tchttp.BaseRequest

	// ClusterId

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Space Id

	SpaceId *uint64 `json:"SpaceId,omitempty" name:"SpaceId"`
	// 权限

	Permission *string `json:"Permission,omitempty" name:"Permission"`
}

func (r *UpdateSpacePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSpacePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateHandleDiskTaskRequest struct {
	*tchttp.BaseRequest

	// 硬盘信息序列化后的字符串

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// Cluster信息序列化后的字符串

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 备注信息

	Note *string `json:"Note,omitempty" name:"Note"`
}

func (r *GenerateHandleDiskTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateHandleDiskTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubTaskInfoRequest struct {
	*tchttp.BaseRequest

	// 子任务表单id

	SubTaskFormId *uint64 `json:"SubTaskFormId,omitempty" name:"SubTaskFormId"`
}

func (r *GetSubTaskInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubTaskInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeChangesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 升级变更表内容，返回数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 返回码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 返回信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeChangesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeChangesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSpacePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSpacePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSpacePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterConfirmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NewClusterConfirmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterConfirmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartBalanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartBalanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterByIdRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	ClusterInfo *string `json:"ClusterInfo,omitempty" name:"ClusterInfo"`
}

func (r *UpdateClusterByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterByIdRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *GetClusterByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLibraRelatedJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLibraRelatedJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLibraRelatedJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStorageTrendRequest struct {
	*tchttp.BaseRequest

	// 无

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 开始日期，不填则默认前15天

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束日期，不填则默认前一天

	End *string `json:"End,omitempty" name:"End"`
	// 采样间隔，不填默认一天为间隔

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *GetStorageTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStorageTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDiskByIpsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Ip列表

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *ListDiskByIpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDiskByIpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterApproveRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 是否通过

	Pass *string `json:"Pass,omitempty" name:"Pass"`
	// 审批意见

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 审批人

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
}

func (r *NewClusterApproveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterApproveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbortBalanceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Job ID

	JobId *int64 `json:"JobId,omitempty" name:"JobId"`
	// Job主题

	JobTheme *int64 `json:"JobTheme,omitempty" name:"JobTheme"`
}

func (r *AbortBalanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AbortBalanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSpaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSpaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGlobalDashboardRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetGlobalDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGlobalDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRegionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubTaskInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 返回信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 子任务详情，返回数据序列化字符串

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubTaskInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetVersionInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 错误消息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetVersionInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVersionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveHostsRequest struct {
	*tchttp.BaseRequest

	// IP列表

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 操作人

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
}

func (r *RemoveHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ControlNodeStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ControlNodeStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ControlNodeStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSpaceByIdRequest struct {
	*tchttp.BaseRequest

	// ClusterId

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// SpaceId

	SpaceId *uint64 `json:"SpaceId,omitempty" name:"SpaceId"`
}

func (r *GetSpaceByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSpaceByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterStage0Request struct {
	*tchttp.BaseRequest

	// 用户名字

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// AZ个数

	AzCount *string `json:"AzCount,omitempty" name:"AzCount"`
}

func (r *NewClusterStage0Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterStage0Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterStage2Request struct {
	*tchttp.BaseRequest

	// ClusterForm信息

	ClusterForm *string `json:"ClusterForm,omitempty" name:"ClusterForm"`
}

func (r *NewClusterStage2Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterStage2Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewColumbusRequest struct {
	*tchttp.BaseRequest

	// columbus服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// ServerNic

	ServerNic *string `json:"ServerNic,omitempty" name:"ServerNic"`
	// ScakNic

	ScalNic *string `json:"ScalNic,omitempty" name:"ScalNic"`
	// tokenseed

	TokenSeed *string `json:"TokenSeed,omitempty" name:"TokenSeed"`
	// ip列表

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 用户名

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 二进制版本信息

	BinaryVersion *string `json:"BinaryVersion,omitempty" name:"BinaryVersion"`
	// 配置版本信息

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 关联一级模块

	ServerBusi1 *string `json:"ServerBusi1,omitempty" name:"ServerBusi1"`
	// 关联二级模块

	ServerBusi2 *string `json:"ServerBusi2,omitempty" name:"ServerBusi2"`
	// Buffer三级模块

	BufferBusi *string `json:"BufferBusi,omitempty" name:"BufferBusi"`
}

func (r *NewColumbusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewColumbusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlackNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlackNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlackNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceConfRequest struct {
	*tchttp.BaseRequest

	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetServiceConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTaskByUuidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTaskByUuidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTaskByUuidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ControlNodeStatsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ControlNodeStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ControlNodeStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTaskByUuidRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *GetTaskByUuidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTaskByUuidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNodeStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNodeStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryChainTaskByUuidRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *RetryChainTaskByUuidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryChainTaskByUuidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOnlineVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOnlineVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOnlineVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetVersionInfoRequest struct {
	*tchttp.BaseRequest

	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetVersionInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVersionInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUpgradeChangeRequest struct {
	*tchttp.BaseRequest

	// 升级变更配置名称

	ConfName *string `json:"ConfName,omitempty" name:"ConfName"`
}

func (r *DeleteUpgradeChangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpgradeChangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSpaceByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSpaceByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSpaceByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Clusters信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误消息说明

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmHandleErrorDiskCompleteRequest struct {
	*tchttp.BaseRequest

	// 磁盘信息(models.DiskInfo)序列化以后的信息

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// Cluster信息序列化以后的信息

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
}

func (r *ConfirmHandleErrorDiskCompleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmHandleErrorDiskCompleteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSpaceTemplateByIdRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *GetSpaceTemplateByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSpaceTemplateByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DropSpaceRequest struct {
	*tchttp.BaseRequest

	// 集群信息

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// Space Id

	SpaceId *uint64 `json:"SpaceId,omitempty" name:"SpaceId"`
	// SpaceAction

	SpaceAction *string `json:"SpaceAction,omitempty" name:"SpaceAction"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DropSpaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DropSpaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDiskStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDiskStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDiskStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GinMakePackageInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GinMakePackageInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GinMakePackageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewColumbusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NewColumbusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewColumbusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLibraRelatedJobsRequest struct {
	*tchttp.BaseRequest

	// Volume ID

	VolumeId *int64 `json:"VolumeId,omitempty" name:"VolumeId"`
	// Node ID

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// DIsk ID

	DiskId *int64 `json:"DiskId,omitempty" name:"DiskId"`
	// 序列号

	Seq *int64 `json:"Seq,omitempty" name:"Seq"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime []*string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 排序

	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
	// 任务类型

	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
	// 源或者目的地

	SrcOrDst *int64 `json:"SrcOrDst,omitempty" name:"SrcOrDst"`
	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Page页

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// Page大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListLibraRelatedJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLibraRelatedJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitHandleErrorDiskTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitHandleErrorDiskTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitHandleErrorDiskTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitUpgradeChangesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id，用来在任务管理展示(cluster_id=-1过滤)

		Data *string `json:"Data,omitempty" name:"Data"`
		// 返回码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 返回信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitUpgradeChangesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitUpgradeChangesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNodeFilter struct {

	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Node Id数组

	NodeIdList []*string `json:"NodeIdList,omitempty" name:"NodeIdList"`
	// Module Id

	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
	// AZ名称

	AzName *string `json:"AzName,omitempty" name:"AzName"`
	// MZ ID

	MzId *int64 `json:"MzId,omitempty" name:"MzId"`
	// 服务状态

	ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`
	// 健康状态

	HealthStatus *string `json:"HealthStatus,omitempty" name:"HealthStatus"`
	// OnlyBlack

	OnlyBlack *bool `json:"OnlyBlack,omitempty" name:"OnlyBlack"`
	// IP列表

	IpSet []*string `json:"IpSet,omitempty" name:"IpSet"`
	// 排序顺序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序

	DescOrder *bool `json:"DescOrder,omitempty" name:"DescOrder"`
	// 部署状态

	DeployStatus *string `json:"DeployStatus,omitempty" name:"DeployStatus"`
}

type AddErrorDiskRequest struct {
	*tchttp.BaseRequest

	// 磁盘信息(models.DiskInfo)序列化以后的信息

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// Cluster信息序列化以后的信息

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 是否强制添加坏盘

	ForceDrop *bool `json:"ForceDrop,omitempty" name:"ForceDrop"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *AddErrorDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddErrorDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllClustersOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllClustersOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllClustersOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDeployableHostRequest struct {
	*tchttp.BaseRequest

	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 获取可部署Host请求

	DeployableHostRequest *DeployableHostReq `json:"DeployableHostRequest,omitempty" name:"DeployableHostRequest"`
}

func (r *GetDeployableHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDeployableHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStorageTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStorageTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStorageTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryChainTaskByUuidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryChainTaskByUuidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryChainTaskByUuidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollBackSubTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 返回信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollBackSubTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollBackSubTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetScalRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Cluster信息

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// Module ID

	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
	// MZ ID

	MzId *int64 `json:"MzId,omitempty" name:"MzId"`
	// Az Name

	AzName *string `json:"AzName,omitempty" name:"AzName"`
}

func (r *GetScalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetScalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterDashboardRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 具体时间，不填的话默认过去24小时的数据

	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *GetClusterDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetErrorDiskChangeLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetErrorDiskChangeLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetErrorDiskChangeLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListErrorDiskRequest struct {
	*tchttp.BaseRequest

	// 页面大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 当前页面

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 硬盘序列号

	DiskSn *string `json:"DiskSn,omitempty" name:"DiskSn"`
	// 错误类型

	ErrorType *string `json:"ErrorType,omitempty" name:"ErrorType"`
	// 任务ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 任务结果

	Result *string `json:"Result,omitempty" name:"Result"`
	// 当前状态

	State *string `json:"State,omitempty" name:"State"`
	// 盘符

	Symbol *string `json:"Symbol,omitempty" name:"Symbol"`
	// 是否降序

	DescOrder *bool `json:"DescOrder,omitempty" name:"DescOrder"`
	// 排序方式

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListErrorDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListErrorDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDiskStatsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetDiskStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDiskStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VueOption struct {

	// Lable

	Label *string `json:"Label,omitempty" name:"Label"`
	// Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetColumbusListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetColumbusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetColumbusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterStage2Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误消息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NewClusterStage2Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterStage2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateClusterByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PackageInfo struct {

	// StartIndex

	StartIndex *int64 `json:"StartIndex,omitempty" name:"StartIndex"`
	// StableBinaryVersion

	StableBinaryVersion *StableVersion `json:"StableBinaryVersion,omitempty" name:"StableBinaryVersion"`
	// StableConfigVersion

	StableConfigVersion *StableVersion `json:"StableConfigVersion,omitempty" name:"StableConfigVersion"`
	// ConfigVersion

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// BinaryVersion

	BinaryVersion *string `json:"BinaryVersion,omitempty" name:"BinaryVersion"`
	// ModuleId

	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
	// Module

	Module *string `json:"Module,omitempty" name:"Module"`
	// BinaryVersions

	BinaryVersions []*VueOption `json:"BinaryVersions,omitempty" name:"BinaryVersions"`
	// ConfigVersions

	ConfigVersions []*VueOption `json:"ConfigVersions,omitempty" name:"ConfigVersions"`
}

type GetGroupInfosRequest struct {
	*tchttp.BaseRequest

	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 集群ID

	ClusterId []*int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetGroupInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSpacesRequest struct {
	*tchttp.BaseRequest

	// 无

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 无

	ListSpacesRequest *ListSpaceReq `json:"ListSpacesRequest,omitempty" name:"ListSpacesRequest"`
}

func (r *ListSpacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSpacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskFilterItem struct {

	// Metric

	Metric *int64 `json:"Metric,omitempty" name:"Metric"`
	// Operator

	Operator *int64 `json:"Operator,omitempty" name:"Operator"`
	// Value

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type GetSpaceTemplateByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSpaceTemplateByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSpaceTemplateByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineNodesRequest struct {
	*tchttp.BaseRequest

	// 操作人

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Node数组

	NodeList []*string `json:"NodeList,omitempty" name:"NodeList"`
}

func (r *OfflineNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OfflineNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetClusterStageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误消息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetClusterStageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetClusterStageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSpaceTemplateRequest struct {
	*tchttp.BaseRequest
}

func (r *ListSpaceTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSpaceTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRegionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDiskByIpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListDiskByIpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDiskByIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SoftInfo struct {

	// Module Id

	Module *int64 `json:"Module,omitempty" name:"Module"`
	// Module Name

	SoftName *string `json:"SoftName,omitempty" name:"SoftName"`
	// 二进制版本

	BinaryVersion *string `json:"BinaryVersion,omitempty" name:"BinaryVersion"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 可用二进制版本

	AvailableBinaryVersion []*string `json:"AvailableBinaryVersion,omitempty" name:"AvailableBinaryVersion"`
	// 可用配置版本

	AvailableConfigVersion []*string `json:"AvailableConfigVersion,omitempty" name:"AvailableConfigVersion"`
}

type CapacityReportRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 获取某一天的容量数据，不填则默认获取当天的

	Date *string `json:"Date,omitempty" name:"Date"`
	// 获取容量趋势的起始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 获取容量趋势的结束时间

	End *string `json:"End,omitempty" name:"End"`
}

func (r *CapacityReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CapacityReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBlobSpaceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetBlobSpaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBlobSpaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListErrorDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListErrorDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListErrorDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryChainTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *QueryChainTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryChainTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitClusterStageRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *SubmitClusterStageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitClusterStageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddHostsRequest struct {
	*tchttp.BaseRequest

	// IP列表

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 操作人

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 前置检查

	SysRegister *string `json:"SysRegister,omitempty" name:"SysRegister"`
	// 是否严格校验

	Strict *string `json:"Strict,omitempty" name:"Strict"`
	// Cluster信息

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *AddHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUpgradeChangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 返回信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUpgradeChangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpgradeChangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAvailabilityTrendRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 开始日期，不填默认第前十五天

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束日期，不填默认第前一天

	End *string `json:"End,omitempty" name:"End"`
}

func (r *GetAvailabilityTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAvailabilityTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDiskRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Az

	Az *string `json:"Az,omitempty" name:"Az"`
	// Mz

	Mz *int64 `json:"Mz,omitempty" name:"Mz"`
	// Ip列表

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// Node Id

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 是否降序排序

	DescOrder *bool `json:"DescOrder,omitempty" name:"DescOrder"`
	// 排序规则

	OrderBy *int64 `json:"OrderBy,omitempty" name:"OrderBy"`
	// Filter数组

	Filters []*DiskFilterItem `json:"Filters,omitempty" name:"Filters"`
	// 页面Size

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 第几页

	Page *int64 `json:"Page,omitempty" name:"Page"`
}

func (r *ListDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBlobSpaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBlobSpaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBlobSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetErrorDiskChangeLogRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *GetErrorDiskChangeLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetErrorDiskChangeLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRequestTrendRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 开始日期，不填默认第前十五天

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束日期，不填默认第前一天

	End *string `json:"End,omitempty" name:"End"`
}

func (r *GetRequestTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBalanceJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListBalanceJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBalanceJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGlobalDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGlobalDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGlobalDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GinDoPlanRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// IP列表

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 安装包信息

	PackageInfos *string `json:"PackageInfos,omitempty" name:"PackageInfos"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *GinDoPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GinDoPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleVersionRequest struct {
	*tchttp.BaseRequest

	// 模块名

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *GetModuleVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitUpgradeChangesRequest struct {
	*tchttp.BaseRequest

	// 升级配置信息，数据序列化字符串

	UpgradeInfo *string `json:"UpgradeInfo,omitempty" name:"UpgradeInfo"`
}

func (r *SubmitUpgradeChangesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitUpgradeChangesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVolumesRequest struct {
	*tchttp.BaseRequest

	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Space Id

	SpaceId *uint64 `json:"SpaceId,omitempty" name:"SpaceId"`
	// Volume个数

	VolumeCount *int64 `json:"VolumeCount,omitempty" name:"VolumeCount"`
}

func (r *CreateVolumesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVolumesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetScalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetScalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetScalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnregisteredNodesRequest struct {
	*tchttp.BaseRequest

	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Node ID数组

	NodeList []*string `json:"NodeList,omitempty" name:"NodeList"`
}

func (r *UnregisteredNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnregisteredNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployableHostReq struct {

	// Module数组

	ModuleList []*string `json:"ModuleList,omitempty" name:"ModuleList"`
	// IP数组

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 限制返回结果

	LimitSize *int64 `json:"LimitSize,omitempty" name:"LimitSize"`
	// Host类型数组

	HostType []*string `json:"HostType,omitempty" name:"HostType"`
}

type ConfirmHandleErrorDiskCompleteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmHandleErrorDiskCompleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmHandleErrorDiskCompleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRequestTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRequestTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLibraJobsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// Page大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// Page编号

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 任务类型

	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
	// Job类型

	JobType *int64 `json:"JobType,omitempty" name:"JobType"`
	// Task状态

	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// Job ID

	JobId *int64 `json:"JobId,omitempty" name:"JobId"`
	// 是否包含子任务

	IncludeChild *bool `json:"IncludeChild,omitempty" name:"IncludeChild"`
	// 创建时间

	CreateTime []*string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 排序方式

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序

	DescOrder *bool `json:"DescOrder,omitempty" name:"DescOrder"`
}

func (r *ListLibraJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLibraJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryChainTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryChainTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryChainTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GinMakeDeployPlanRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// IP列表

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 版本信息

	PackageInfos *string `json:"PackageInfos,omitempty" name:"PackageInfos"`
}

func (r *GinMakeDeployPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GinMakeDeployPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryErrorDiskRequest struct {
	*tchttp.BaseRequest

	// 坏盘任务Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *RetryErrorDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryErrorDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeNodesReq struct {

	// NodeId

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// Ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Module名称

	Module *string `json:"Module,omitempty" name:"Module"`
	// 二进制版本

	BinaryVersion *string `json:"BinaryVersion,omitempty" name:"BinaryVersion"`
	// 配置文件版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

type GetNodeVletsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNodeVletsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeVletsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryClusterFormRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryClusterFormRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryClusterFormRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBlackNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBlackNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBlackNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterStage1Request struct {
	*tchttp.BaseRequest

	// ClusterForm信息序列化

	ClusterForm *string `json:"ClusterForm,omitempty" name:"ClusterForm"`
	// sys_register

	SysRegister *string `json:"SysRegister,omitempty" name:"SysRegister"`
}

func (r *NewClusterStage1Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterStage1Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddErrorDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddErrorDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddErrorDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVolumesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVolumesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVolumesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployNodesRequest struct {
	*tchttp.BaseRequest

	// Cluster Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 操作人

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 部署请求

	DeployNodesRequest *DeployNodesReq `json:"DeployNodesRequest,omitempty" name:"DeployNodesRequest"`
}

func (r *DeployNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOnlineVersionRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetOnlineVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOnlineVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GinMakeDeployPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GinMakeDeployPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GinMakeDeployPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUpgradeChangeFormsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 表信息，返回数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 返回码

		RetCode *string `json:"RetCode,omitempty" name:"RetCode"`
		// 返回信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUpgradeChangeFormsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUpgradeChangeFormsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DropSpaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DropSpaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DropSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OfflineNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModuleVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterApproveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NewClusterApproveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterApproveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetChainTaskByUuidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChainTaskByUuidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetChainTaskByUuidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNodesRequest struct {
	*tchttp.BaseRequest

	// 服务Node状态类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 过滤器

	Filter *ListNodeFilter `json:"Filter,omitempty" name:"Filter"`
	// 当前页

	CurrentPage *int64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 分页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterStage3Request struct {
	*tchttp.BaseRequest

	// 是否检查数据库

	CheckMysql *string `json:"CheckMysql,omitempty" name:"CheckMysql"`
	// ClusterForm信息

	ClusterForm *string `json:"ClusterForm,omitempty" name:"ClusterForm"`
}

func (r *NewClusterStage3Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterStage3Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CapacityReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CapacityReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CapacityReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSpaceRequest struct {
	*tchttp.BaseRequest

	// ClusterId

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 模版ID

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// Space个数

	SpaceCount *int64 `json:"SpaceCount,omitempty" name:"SpaceCount"`
	// StaffName

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
}

func (r *CreateSpaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSpaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateHandleDiskTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateHandleDiskTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateHandleDiskTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryClusterFormByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryClusterFormByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryClusterFormByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewClusterStage0Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据序列化字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 错误码

		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
		// 错误消息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *NewClusterStage0Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *NewClusterStage0Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
