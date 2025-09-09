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
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DeleteLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeABTestConfigRequest struct {
	*tchttp.BaseRequest

	// 灰度项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *DescribeABTestConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeABTestConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCVMAssetInfoRequest struct {
	*tchttp.BaseRequest

	// -

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeCVMAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCVMAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// -

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeDomainAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExaminationScanRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeExaminationScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExaminationScanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskReportOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 输出数据

		Data *TaskReportOverView `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskReportOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskReportOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataOperateLog struct {

	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 操作类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 操作行为

	Action *string `json:"Action,omitempty" name:"Action"`
	// 操作详情

	Content *string `json:"Content,omitempty" name:"Content"`
	// 操作时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 客户端ip

	IP *string `json:"IP,omitempty" name:"IP"`
	// 操作子类型

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 租户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 租户uin

	UIN *string `json:"UIN,omitempty" name:"UIN"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
}

type DescribeSecOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞动态数据

		Data *SecOverviewInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQuickProtectSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快捷防护-下发检测，扫描结果

		PushTaskResult []*PushTaskResult `json:"PushTaskResult,omitempty" name:"PushTaskResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateQuickProtectSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQuickProtectSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetDetailInfoRequest struct {
	*tchttp.BaseRequest

	// 子网id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeSubnetDetailInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetDetailInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseDetail struct {

	// -

	LicenseId *int64 `json:"LicenseId,omitempty" name:"LicenseId"`
	// -

	LicenseType *int64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// -

	LicenseStatus *int64 `json:"LicenseStatus,omitempty" name:"LicenseStatus"`
	// -

	LicenseCnt *int64 `json:"LicenseCnt,omitempty" name:"LicenseCnt"`
	// -

	UsedLicenseCnt *int64 `json:"UsedLicenseCnt,omitempty" name:"UsedLicenseCnt"`
	// -

	OrderStatus *int64 `json:"OrderStatus,omitempty" name:"OrderStatus"`
	// -

	BuyTime *string `json:"BuyTime,omitempty" name:"BuyTime"`
	// -

	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
	// -

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// -

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// -

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// -

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// -

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
	// -

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// -

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

type WafTlsVersionInfo struct {

	// 版本id

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 版本名称

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

type DescribeClusterPodAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*AssetClusterPod `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群pod状态枚举

		PodStatusList []*FilterDataObject `json:"PodStatusList,omitempty" name:"PodStatusList"`
		// 命名空间枚举

		NamespaceList []*FilterDataObject `json:"NamespaceList,omitempty" name:"NamespaceList"`
		// 地域枚举

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 租户枚举

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterPodAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterPodAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkStructureRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 1查所有的region信息 2:查 vpc信息 3:子网信息

	NodeType *int64 `json:"NodeType,omitempty" name:"NodeType"`
	// vpc_id或者region_id

	NodeValue *string `json:"NodeValue,omitempty" name:"NodeValue"`
}

func (r *DescribeNetworkStructureRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkStructureRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserStorageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用的存储量

		UsedStorage *float64 `json:"UsedStorage,omitempty" name:"UsedStorage"`
		// 总的存储量

		TotalStorage *float64 `json:"TotalStorage,omitempty" name:"TotalStorage"`
		// 存储时长

		Duration *int64 `json:"Duration,omitempty" name:"Duration"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserStorageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckItemIdInfo struct {

	// 检测项下发任务用id

	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`
	// 检测项名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检测项危险等级，1低危，2中危，3高危，4严重

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 检查对象的类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 检查项类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 默认规范

	Standard *string `json:"Standard,omitempty" name:"Standard"`
	// 识别来源，1代表cspm，2代表主机

	Come *int64 `json:"Come,omitempty" name:"Come"`
	// 风险说明

	Content *string `json:"Content,omitempty" name:"Content"`
	// 处理建议

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 文档

	Doc *string `json:"Doc,omitempty" name:"Doc"`
	// 是否启用1启用，0禁用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// middle

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
}

type DescribeWafVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Data *DataWafVersion `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerRiskSuggestion struct {

	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 详情

	Body *string `json:"Body,omitempty" name:"Body"`
}

type AddWebServiceRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// ins-xxx

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 1.1.1.1

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 80

	Port *string `json:"Port,omitempty" name:"Port"`
	// www.xxx.com

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 是否进行漏洞扫描: 1：是，0:否

	IsScan *string `json:"IsScan,omitempty" name:"IsScan"`
}

func (r *AddWebServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddWebServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSendVoucherToUserTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 优惠券类型：
	// importantProtect:重保

	VoucherType *string `json:"VoucherType,omitempty" name:"VoucherType"`
}

func (r *CreateSendVoucherToUserTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSendVoucherToUserTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFGRiskAdvanceCFGListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeCFGRiskAdvanceCFGListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFGRiskAdvanceCFGListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总核数

		AllCores *int64 `json:"AllCores,omitempty" name:"AllCores"`
		// 授权核数

		AuthorizedCores *int64 `json:"AuthorizedCores,omitempty" name:"AuthorizedCores"`
		// 是否付费

		IsPay *bool `json:"IsPay,omitempty" name:"IsPay"`
		// tke授权 1:授权 0 未授权

		TKEAuth *int64 `json:"TKEAuth,omitempty" name:"TKEAuth"`
		// 产品当前状态 0可试用，1未防护，2防护中（已购买），3试用中

		TCSSStatus *int64 `json:"TCSSStatus,omitempty" name:"TCSSStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicIpAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关资产详情

		Data *PublicIpAssetInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicIpAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicIpAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIpProtectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIpProtectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpProtectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BASTaskObject struct {

	// bas再次执行时下发的id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 剧本id

	SID *string `json:"SID,omitempty" name:"SID"`
	// 用于停止任务的id

	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
	// 剧本名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资产总数

	AssetCnt *uint64 `json:"AssetCnt,omitempty" name:"AssetCnt"`
	// 已执行成功的资产数

	SucceedAtCnt *uint64 `json:"SucceedAtCnt,omitempty" name:"SucceedAtCnt"`
	// 关联告警的个数

	AlertLogCnt *uint64 `json:"AlertLogCnt,omitempty" name:"AlertLogCnt"`
	// 首次执行时间

	FirstExecTime *string `json:"FirstExecTime,omitempty" name:"FirstExecTime"`
	// 上次执行时间

	LastExecTime *string `json:"LastExecTime,omitempty" name:"LastExecTime"`
	// 执行中止时的结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 执行成功次数

	Success *uint64 `json:"Success,omitempty" name:"Success"`
	// 执行异常次数

	Fail *uint64 `json:"Fail,omitempty" name:"Fail"`
	// 执行中止次数

	Stopped *uint64 `json:"Stopped,omitempty" name:"Stopped"`
	// 最近执行情况，1是正在执行，2是执行完成，3是执行异常，4是执行中止

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 任务关联战术战技

	Tactics []*BASTacticsObject `json:"Tactics,omitempty" name:"Tactics"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 别名

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 描述异常原因

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 返回task对应的主机各类型告警数量

	AlertLog []*BASAlertTypeCnt `json:"AlertLog,omitempty" name:"AlertLog"`
}

type VSSPayInfo struct {

	// 无

	IsCanScan *bool `json:"IsCanScan,omitempty" name:"IsCanScan"`
	// 无

	VSSPay *bool `json:"VSSPay,omitempty" name:"VSSPay"`
}

type DescribeClbLoadBalancersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据列表

		Data []*ClbLoadBalancers `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClbLoadBalancersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClbLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalImageComponentRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// Filter过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeLocalImageComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLocalImageComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCvmInfoRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
}

func (r *DescribeUserCvmInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCvmInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CFWConfig struct {

	// 是否防火墙用户

	IsFWUser *bool `json:"IsFWUser,omitempty" name:"IsFWUser"`
	// 防护状态，1未防护，2已购买，3试用中，4已过期

	FWUserStatus *int64 `json:"FWUserStatus,omitempty" name:"FWUserStatus"`
	// 是否可以申请试用，true可以申请

	CanApplyTrial *bool `json:"CanApplyTrial,omitempty" name:"CanApplyTrial"`
	// 不是企业用户，不符合防火墙试用条件"  无法试用原因，可试用为空

	CanNotApplyReason *string `json:"CanNotApplyReason,omitempty" name:"CanNotApplyReason"`
	// 最近登录时间

	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
	// 试用结束时间

	ApplyEndTime *string `json:"ApplyEndTime,omitempty" name:"ApplyEndTime"`
	// ip防护配额

	IpProtectLimit *int64 `json:"IpProtectLimit,omitempty" name:"IpProtectLimit"`
	// ip防护已使用配额

	IpProtectUsed *int64 `json:"IpProtectUsed,omitempty" name:"IpProtectUsed"`
	// 用户版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type ContainerAssetInfo struct {

	// 容器id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 容器名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// pod id

	PodId *string `json:"PodId,omitempty" name:"PodId"`
	// pod ip

	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`
	// pod name

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 主机id

	MachineId *string `json:"MachineId,omitempty" name:"MachineId"`
	// 主机名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 容器创建时间

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
	// 最近体检时间

	CheckTime *string `json:"CheckTime,omitempty" name:"CheckTime"`
	// 网络状态描述

	NetworkStatusReason *string `json:"NetworkStatusReason,omitempty" name:"NetworkStatusReason"`
	// 网络状态

	NetworkStatus *string `json:"NetworkStatus,omitempty" name:"NetworkStatus"`
	// 是否核心资产，1核心，2非核心

	IsCore *int64 `json:"IsCore,omitempty" name:"IsCore"`
	// 容器状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 核心数

	CpuCount *float64 `json:"CpuCount,omitempty" name:"CpuCount"`
	// 内存大小,kb

	MemCount *int64 `json:"MemCount,omitempty" name:"MemCount"`
	// 公网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 私网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 漏洞风险数

	VulRisk *int64 `json:"VulRisk,omitempty" name:"VulRisk"`
	// 配置风险数

	CFGRisk *int64 `json:"CFGRisk,omitempty" name:"CFGRisk"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 来源地域

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 原始容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 复杂名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 用户是否标记核心，0未标记，1已标记

	UserAction *string `json:"UserAction,omitempty" name:"UserAction"`
}

type CreateProductStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 该部分返回体检下发是否成功

		Data *CreateExamMissionResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSolidProtectionTrialsRequest struct {
	*tchttp.BaseRequest

	// 选择开通试用产品的名称 CFW云防火墙 WAFWeb应用防火墙 CWP主机安全 TCSS容器安全 SOC安全运营中心

	TrialProducts []*string `json:"TrialProducts,omitempty" name:"TrialProducts"`
	// Waf 创建实例的参数

	WafParamsData *WafParams `json:"WafParamsData,omitempty" name:"WafParamsData"`
	// 云防创建试用的参数

	CFWParamsData *CFWParams `json:"CFWParamsData,omitempty" name:"CFWParamsData"`
	// 主机安全创建试用所需的参数

	CWPParamsData *CWPParams `json:"CWPParamsData,omitempty" name:"CWPParamsData"`
	// 试用来源， dns、clb 等 默认为soc

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *CreateSolidProtectionTrialsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolidProtectionTrialsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSecTrendListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeAssetSecTrendListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSecTrendListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNeedToSolveListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口请求是否成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 接口状态

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 更新的信息，空代表成功

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNeedToSolveListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNeedToSolveListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallBASAssetAgentByCWPRequest struct {
	*tchttp.BaseRequest

	// 资产ID

	AssetID *string `json:"AssetID,omitempty" name:"AssetID"`
	// 资产对应的appid

	AssetAppID *int64 `json:"AssetAppID,omitempty" name:"AssetAppID"`
}

func (r *InstallBASAssetAgentByCWPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallBASAssetAgentByCWPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafInstances struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 字段id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 域名个数

	DomainCount *int64 `json:"DomainCount,omitempty" name:"DomainCount"`
	// 域名上限

	SubDomainLimit *int64 `json:"SubDomainLimit,omitempty" name:"SubDomainLimit"`
	// 域名类型

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// 套餐

	Level *uint64 `json:"Level,omitempty" name:"Level"`
}

type StartWafPatch struct {

	// 1设置成功，0设置失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeCFWStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否为云防付费用户

		IsFWUser *bool `json:"IsFWUser,omitempty" name:"IsFWUser"`
		// 防护状态，1未防护，2已购买，3试用中，4已过期

		FWUserStatus *int64 `json:"FWUserStatus,omitempty" name:"FWUserStatus"`
		// 是否可以申请试用，true可以申请

		CanApplyTrial *bool `json:"CanApplyTrial,omitempty" name:"CanApplyTrial"`
		// -

		CanNotApplyReason *string `json:"CanNotApplyReason,omitempty" name:"CanNotApplyReason"`
		// 最后登陆时间

		LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
		// 试用结束时间

		ApplyEndTime *string `json:"ApplyEndTime,omitempty" name:"ApplyEndTime"`
		// ip防护配额

		IpProtectLimit *int64 `json:"IpProtectLimit,omitempty" name:"IpProtectLimit"`
		// -

		IpProtectUsed *int64 `json:"IpProtectUsed,omitempty" name:"IpProtectUsed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCFWStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogAccessSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogAccessSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogAccessSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BASScriptObject struct {

	// 剧本名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 剧本ID

	SID *string `json:"SID,omitempty" name:"SID"`
	// 剧本类型，1：默认剧本，2：自定义剧本

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 剧本描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 剧本战术

	Tactics []*BASTacticsObject `json:"Tactics,omitempty" name:"Tactics"`
	// 执行成功次数

	Success *uint64 `json:"Success,omitempty" name:"Success"`
	// 执行异常次数

	Fail *uint64 `json:"Fail,omitempty" name:"Fail"`
	// 执行中止次数

	Stopped *uint64 `json:"Stopped,omitempty" name:"Stopped"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 首次执行时间

	FirstExecTime *string `json:"FirstExecTime,omitempty" name:"FirstExecTime"`
	// 上次执行时间

	LastExecTime *string `json:"LastExecTime,omitempty" name:"LastExecTime"`
	// 关联告警数

	AlertLogCnt *uint64 `json:"AlertLogCnt,omitempty" name:"AlertLogCnt"`
	// 主机告警类型详情

	AlertLog []*BASAlertTypeCnt `json:"AlertLog,omitempty" name:"AlertLog"`
}

type ExportAuditLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID

		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAuditLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAuditLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyTrialRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 用户等级

	UserLevel *string `json:"UserLevel,omitempty" name:"UserLevel"`
}

func (r *ApplyTrialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyTrialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLastTaskOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 输出概览数据

		Data []*RecentTaskList `json:"Data,omitempty" name:"Data"`
		// 总任务数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 主账户ID列表

		UINList []*string `json:"UINList,omitempty" name:"UINList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLastTaskOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLastTaskOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseBindScheduleRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 限制条数,默认10.

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Filter过滤条件

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLicenseBindScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseBindScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskPredictTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预计扫描120分钟

		Data *int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskPredictTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskPredictTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAuditLogRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 导出条数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 导出日志类型

	LogType []*string `json:"LogType,omitempty" name:"LogType"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询逻辑表达式

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 按时间升序/降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 集群 0:国内 1:国际

	Cluster *int64 `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *ExportAuditLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAuditLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserCVEStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Data *StartWafPatch `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserCVEStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserCVEStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogAuditData struct {

	// 单条数据

	SingleData []*KeyValue `json:"SingleData,omitempty" name:"SingleData"`
}

type TaskLogSimpleInfo struct {

	// 任务报告名称

	TaskLogName *string `json:"TaskLogName,omitempty" name:"TaskLogName"`
	// 任务报告id

	TaskLogId *string `json:"TaskLogId,omitempty" name:"TaskLogId"`
	// 任务报告生成时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// APP ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type AssetInfoDetail struct {

	// 用户appid

	AppID *string `json:"AppID,omitempty" name:"AppID"`
	// CVE编号

	CVEId *string `json:"CVEId,omitempty" name:"CVEId"`
	// 是扫描，0默认未扫描，1正在扫描，2扫描完成，3扫描出错

	IsScan *int64 `json:"IsScan,omitempty" name:"IsScan"`
	// 影响资产数目

	InfluenceAsset *int64 `json:"InfluenceAsset,omitempty" name:"InfluenceAsset"`
	// 未修复资产数目

	NotRepairAsset *int64 `json:"NotRepairAsset,omitempty" name:"NotRepairAsset"`
	// 未防护资产数目

	NotProtectAsset *int64 `json:"NotProtectAsset,omitempty" name:"NotProtectAsset"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务百分比

	TaskPercent *int64 `json:"TaskPercent,omitempty" name:"TaskPercent"`
	// 任务时间

	TaskTime *int64 `json:"TaskTime,omitempty" name:"TaskTime"`
	// 扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data *CFWConfig `json:"Data,omitempty" name:"Data"`
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

type DescribeVpcAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*Vpc `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// vpc列表

		VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`
		// 地域列表

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// appid列表

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafAllInstancesRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 查询参数

	Filters []*WafInstancesFilter `json:"Filters,omitempty" name:"Filters"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeWafAllInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafAllInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSCCCSVRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 筛选过滤参数

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// 筛选过滤的接口

	ExportAction *string `json:"ExportAction,omitempty" name:"ExportAction"`
	// 需要导出的列

	ExportFields []*string `json:"ExportFields,omitempty" name:"ExportFields"`
	// 是否异步下载，1: 是，其他为否

	Async *int64 `json:"Async,omitempty" name:"Async"`
}

func (r *ExportSCCCSVRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSCCCSVRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllProductsInfo struct {

	// 无

	IsVSSScan *bool `json:"IsVSSScan,omitempty" name:"IsVSSScan"`
	// 无

	UserAssetNum *uint64 `json:"UserAssetNum,omitempty" name:"UserAssetNum"`
	// 无

	ProductVersions []*ProductVersionInfos `json:"ProductVersions,omitempty" name:"ProductVersions"`
}

type BASSubTechniquesObject struct {

	// 子战技名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 子战技描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 子战技ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

type AllInfo struct {

	// 攻击趋势

	Bug []*BugInfoTrend `json:"Bug,omitempty" name:"Bug"`
	// 暴露面趋势

	Internet []*AssetTrend `json:"Internet,omitempty" name:"Internet"`
}

type CreatePocScanMissionRequest struct {
	*tchttp.BaseRequest

	// CVE-2012-12121

	CVEId *string `json:"CVEId,omitempty" name:"CVEId"`
	// sql注入漏洞

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

func (r *CreatePocScanMissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePocScanMissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetTopXRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 时间参数 1:24小时 2:7天 3:30天

	TimeRange *int64 `json:"TimeRange,omitempty" name:"TimeRange"`
}

func (r *DescribeAssetTopXRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTopXRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerAssetInfoRequest struct {
	*tchttp.BaseRequest

	// -

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeContainerAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取导出的异步任务信息

		Data []*AsyncExportTaskResult `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PDFDownLoadRecord struct {

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务报告id

	TaskLogId *string `json:"TaskLogId,omitempty" name:"TaskLogId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务报告名称

	TaskLogName *string `json:"TaskLogName,omitempty" name:"TaskLogName"`
	// 下载时间

	DownTime *string `json:"DownTime,omitempty" name:"DownTime"`
	// APP ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type CreateWafClbHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态 0为成功 其他值失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 消息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 报错代码

		Code *string `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWafClbHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWafClbHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAccountInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户账户信息

		Data *UserAccountInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetWorkInfo struct {

	// 总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 节点信息

	Node []*GroupNodeInfo `json:"Node,omitempty" name:"Node"`
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否集团账号用户

		IsOrganizationUser *bool `json:"IsOrganizationUser,omitempty" name:"IsOrganizationUser"`
		// 是否管理员

		IsManager *bool `json:"IsManager,omitempty" name:"IsManager"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// success

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 0

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 互联网暴露面和攻击趋势数据

		Data *AllInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogAccessConfig struct {

	// 日志源

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 累计存储条数

	StoreCnt *uint64 `json:"StoreCnt,omitempty" name:"StoreCnt"`
	// 接入开关

	Switch *bool `json:"Switch,omitempty" name:"Switch"`
	// 日志类型说明数量

	LogType []*LogAccessTypeDescCount `json:"LogType,omitempty" name:"LogType"`
}

type AsyncExportTaskResult struct {

	// 任务处理成功时的导出的链接

	Url *string `json:"Url,omitempty" name:"Url"`
	// 导出的状态: -1 没有该任务，1: 任务正在处理中，2:任务已处理成功，3:任务处理失败

	Status *string `json:"Status,omitempty" name:"Status"`
	// 如果任务处理失败，这里显示错误信息

	Error *string `json:"Error,omitempty" name:"Error"`
	// -

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// -

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type CreateApplyTrialRequest struct {
	*tchttp.BaseRequest

	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *CreateApplyTrialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplyTrialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFreeLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口请求结果

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 接口请求状态

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 相关信息

		Data *ExamFreeLimitResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFreeLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFreeLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IndeterminacyAssetInfo struct {

	// vpc-id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc-name

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// 最近扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
}

type PortViewPortRisk struct {

	// 影响资产

	NoHandleCount *int64 `json:"NoHandleCount,omitempty" name:"NoHandleCount"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 处置建议,0保持现状、1限制访问、2封禁端口

	Suggestion *uint64 `json:"Suggestion,omitempty" name:"Suggestion"`
	// 状态，0未处理、1已处置、2已忽略

	AffectAssetCount *string `json:"AffectAssetCount,omitempty" name:"AffectAssetCount"`
	// 资产唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 资产子类型

	From *string `json:"From,omitempty" name:"From"`
	// 前端索引

	Index *string `json:"Index,omitempty" name:"Index"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 服务

	Service *string `json:"Service,omitempty" name:"Service"`
}

type ModifyUpdateBugInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 选择开通试用产品状态返回

		Data *UpdateBugInfoResp `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUpdateBugInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpdateBugInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRiskCenterScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRiskCenterScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskCenterScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 冒号前后为映射关系；冒号左边为测试环境，冒号右边为现网

		AppIdMap []*string `json:"AppIdMap,omitempty" name:"AppIdMap"`
		// 冒号前后为映射关系；冒号左边为测试环境，冒号右边为现网

		UinMap []*string `json:"UinMap,omitempty" name:"UinMap"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出信息

		Data *ExportInfo `json:"Data,omitempty" name:"Data"`
		// 返回代码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回消息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageRateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集团成员的数量

		Num *int64 `json:"Num,omitempty" name:"Num"`
		// 每个成员的信息和比例

		NodeInfo []*NodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStorageRateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStorageRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserLogAuditSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 多个业务类型的日志审计配置详情

		Data []*LogAuditSetting `json:"Data,omitempty" name:"Data"`
		// 返回状态码 0 成功 非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息 success 成功 其他 不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserLogAuditSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserLogAuditSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcConnectedVpcListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*VpcConnectedVpcInfo `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域列表

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 连接类型列表

		ConnTypeList []*FilterDataObject `json:"ConnTypeList,omitempty" name:"ConnTypeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcConnectedVpcListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcConnectedVpcListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *SyncAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {

	// 英文名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 中文名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type TaskAdvanceCFG struct {

	// 端口风险高级配置

	PortRisk []*PortRiskAdvanceCFGParamItem `json:"PortRisk,omitempty" name:"PortRisk"`
	// 漏洞风险高级配置

	VulRisk []*TaskCenterVulRiskInputParam `json:"VulRisk,omitempty" name:"VulRisk"`
	// 弱口令风险高级配置

	WeakPwdRisk []*TaskCenterWeakPwdRiskInputParam `json:"WeakPwdRisk,omitempty" name:"WeakPwdRisk"`
	// 配置风险高级配置

	CFGRisk []*TaskCenterCFGRiskInputParam `json:"CFGRisk,omitempty" name:"CFGRisk"`
}

type VULViewVULRisk struct {

	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 影响资产

	NoHandleCount *int64 `json:"NoHandleCount,omitempty" name:"NoHandleCount"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 状态，0未处理、1已处置、2已忽略

	AffectAssetCount *uint64 `json:"AffectAssetCount,omitempty" name:"AffectAssetCount"`
	// 资产唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 资产子类型

	From *string `json:"From,omitempty" name:"From"`
	// 前端索引

	Index *string `json:"Index,omitempty" name:"Index"`
	// 漏洞类型

	VULType *string `json:"VULType,omitempty" name:"VULType"`
	// 漏洞名

	VULName *string `json:"VULName,omitempty" name:"VULName"`
	// cve

	CVE *string `json:"CVE,omitempty" name:"CVE"`
	// 描述

	Describe *string `json:"Describe,omitempty" name:"Describe"`
	// 负载

	Payload *string `json:"Payload,omitempty" name:"Payload"`
	// 版本名

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 相关引用

	References *string `json:"References,omitempty" name:"References"`
	// 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 漏洞链接

	VULURL *string `json:"VULURL,omitempty" name:"VULURL"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 修复建议

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞

	EMGCVulType *int64 `json:"EMGCVulType,omitempty" name:"EMGCVulType"`
}

type DeleteId struct {

	// 被删除的资源的appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 被删除的资源的id

	Id *string `json:"Id,omitempty" name:"Id"`
}

type CreateACRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0:操作成功

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateACRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateACRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterWebsiteRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterWebsiteRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterWebsiteRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVULRiskAdvanceCFGListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeVULRiskAdvanceCFGListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVULRiskAdvanceCFGListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportURL struct {

	// 导出文件的临时链接

	URL *string `json:"URL,omitempty" name:"URL"`
}

type RiskCenterAssetObj struct {

	// 公网IP/域名

	PublicIPDomain *string `json:"PublicIPDomain,omitempty" name:"PublicIPDomain"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeExportTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 导出任务id

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
	// 集群 0:国内 1:国际

	Cluster *int64 `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DescribeExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalImageComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件

		Data []*AssetLocalImageComponent `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLocalImageComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLocalImageComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBASSummaryDataRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *GetBASSummaryDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBASSummaryDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskCenterScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 0，修改成功，其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRiskCenterScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskCenterScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataSearchBug struct {

	// 返回查询状态

	StateCode *string `json:"StateCode,omitempty" name:"StateCode"`
	// 无

	DataBug []*BugInfoDetail `json:"DataBug,omitempty" name:"DataBug"`
	// 无

	DataAsset []*AssetInfoDetail `json:"DataAsset,omitempty" name:"DataAsset"`
	// true支持扫描。false不支持扫描

	VSSScan *bool `json:"VSSScan,omitempty" name:"VSSScan"`
	// 0不支持，1支持

	CWPScan *string `json:"CWPScan,omitempty" name:"CWPScan"`
	// 1支持虚拟补丁，0或空不支持

	CFWPatch *string `json:"CFWPatch,omitempty" name:"CFWPatch"`
	// 0不支持，1支持

	WafPatch *int64 `json:"WafPatch,omitempty" name:"WafPatch"`
	// 0不支持，1支持

	CWPFix *int64 `json:"CWPFix,omitempty" name:"CWPFix"`
}

type DescribeAuditDiversityRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAuditDiversityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuditDiversityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNICAssetInfoRequest struct {
	*tchttp.BaseRequest

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeNICAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNICAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagGroupAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// 分组信息

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeTagGroupAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagGroupAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBASScriptDetailRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 剧本id

	SID *string `json:"SID,omitempty" name:"SID"`
}

func (r *GetBASScriptDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBASScriptDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncScopeAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *SyncScopeAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncScopeAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecurityTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClbLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 查询内容

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 负载均衡网络类型   OPEN/INTERNAL

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 负载均衡地域

	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`
}

func (r *DescribeClbLoadBalancersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClbLoadBalancersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeakPWDRiskAdvanceCFGListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测项详细信息

		Data []*CheckWeakPassWordItemInfo `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 风险枚举

		RiskLevelLists []*FilterDataObject `json:"RiskLevelLists,omitempty" name:"RiskLevelLists"`
		// 种类枚举

		TypeLists []*FilterDataObject `json:"TypeLists,omitempty" name:"TypeLists"`
		// 检测来源枚举

		CheckFromLists []*FilterDataObject `json:"CheckFromLists,omitempty" name:"CheckFromLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeakPWDRiskAdvanceCFGListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeakPWDRiskAdvanceCFGListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBAssetVO struct {

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc标签

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// 最近扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 配置风险

	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitempty" name:"ConfigurationRisk"`
	// 网络攻击

	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`
	// 网络访问

	Access *uint64 `json:"Access,omitempty" name:"Access"`
	// 扫描任务

	ScanTask *uint64 `json:"ScanTask,omitempty" name:"ScanTask"`
	// 用户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 昵称别名

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
}

type TaskLogURL struct {

	// 报告下载临时链接

	URL *string `json:"URL,omitempty" name:"URL"`
	// 任务报告id

	LogId *string `json:"LogId,omitempty" name:"LogId"`
	// 任务报告名称

	TaskLogName *string `json:"TaskLogName,omitempty" name:"TaskLogName"`
	// APP ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type DeleteAssetConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 被删除的数量

		Data *int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAssetConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAssetConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPortProtectStatusRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 是否强制 1是  0 不是定时更新

	Force *int64 `json:"Force,omitempty" name:"Force"`
	// 来源  漏扫不传或者vss  新的风险端口界面itg

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *ModifyPortProtectStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPortProtectStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetImageBuild struct {

	// 用户appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 资产Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 标签

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// Cmd命令

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 资产

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
}

type CheckWeakPassWordItemInfo struct {

	// 检测项id

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 检测项名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 1低危，2中危，3高危，4严重

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 1主机

	Come *int64 `json:"Come,omitempty" name:"Come"`
	// 描述

	Content *string `json:"Content,omitempty" name:"Content"`
	// 1启用0禁用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// middle

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
}

type DbAssetInfo struct {

	// 云防状态

	CFWStatus *uint64 `json:"CFWStatus,omitempty" name:"CFWStatus"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// vpc信息

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 私网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// vpc信息

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 云防保护版本

	CFWProtectLevel *uint64 `json:"CFWProtectLevel,omitempty" name:"CFWProtectLevel"`
	// tag信息

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
}

type DescribeAlarmTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势列表

		AlarmTrend []*ObjectTrend `json:"AlarmTrend,omitempty" name:"AlarmTrend"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集团用户列表

		Data []*OrganizationInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePortProtectStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1为正在更新 前端轮询，2为更新成功  ，3为更新失败  无需轮询

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePortProtectStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePortProtectStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPayInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0未购买，1已购买

		VSSStatus *int64 `json:"VSSStatus,omitempty" name:"VSSStatus"`
		// 0未购买，1已购买

		CWPStatus *int64 `json:"CWPStatus,omitempty" name:"CWPStatus"`
		// 0未购买，1已购买

		SSAStatus *int64 `json:"SSAStatus,omitempty" name:"SSAStatus"`
		// 漏扫剩余次数

		VSSRemain *int64 `json:"VSSRemain,omitempty" name:"VSSRemain"`
		// 0未购买，1已购买

		CFWStatus *int64 `json:"CFWStatus,omitempty" name:"CFWStatus"`
		// 0未购买，1已购买

		CSIPStatus *int64 `json:"CSIPStatus,omitempty" name:"CSIPStatus"`
		// 漏扫总配额

		VSSQuotaTotal *int64 `json:"VSSQuotaTotal,omitempty" name:"VSSQuotaTotal"`
		// 漏扫已经使用的配额

		VSSScanUsed *int64 `json:"VSSScanUsed,omitempty" name:"VSSScanUsed"`
		// 云安全中心总配额

		CSIPQuotaTotal *int64 `json:"CSIPQuotaTotal,omitempty" name:"CSIPQuotaTotal"`
		// 云安全中心已经使用的配额

		CSIPScanUsed *int64 `json:"CSIPScanUsed,omitempty" name:"CSIPScanUsed"`
		// 体检任务总配额

		TaskQuotaTotal *int64 `json:"TaskQuotaTotal,omitempty" name:"TaskQuotaTotal"`
		// 体检任务已经使用的配额

		TaskScanUsed *int64 `json:"TaskScanUsed,omitempty" name:"TaskScanUsed"`
		// 0未购买，1已购买

		TCSSStatus *int64 `json:"TCSSStatus,omitempty" name:"TCSSStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserPayInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPayInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScanResp struct {

	// 下发结果

	Info *string `json:"Info,omitempty" name:"Info"`
	// 状态，0失败，1成功

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type CreateRiskCenterScanTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 扫描资产信息列表

	Assets []*TaskAssetObject `json:"Assets,omitempty" name:"Assets"`
	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填

	ScanAssetType *int64 `json:"ScanAssetType,omitempty" name:"ScanAssetType"`
	// 扫描项目；port/poc/weakpass/webcontent/configrisk

	ScanItem []*string `json:"ScanItem,omitempty" name:"ScanItem"`
	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填

	ScanPlanType *int64 `json:"ScanPlanType,omitempty" name:"ScanPlanType"`
	// 扫描计划详情

	ScanPlanContent *string `json:"ScanPlanContent,omitempty" name:"ScanPlanContent"`
	// ip/域名/url数组

	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitempty" name:"SelfDefiningAssets"`
	// 请求发起源

	ScanFrom *string `json:"ScanFrom,omitempty" name:"ScanFrom"`
	// 是否本月第一次

	IsNew *bool `json:"IsNew,omitempty" name:"IsNew"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
	// 高级配置

	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitempty" name:"TaskAdvanceCFG"`
	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式

	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
}

func (r *CreateRiskCenterScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskCenterScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*WebAssetObject `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetCluster struct {

	// 租户id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 租户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 租户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 集群id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 集群名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 集群类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 集群创建时间

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// kubernetes版本

	KubernetesVersion *string `json:"KubernetesVersion,omitempty" name:"KubernetesVersion"`
	// 运行时组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 运行时组件版本

	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`
	// 组件状态

	ComponentStatus *string `json:"ComponentStatus,omitempty" name:"ComponentStatus"`
	// 体检时间

	CheckTime *string `json:"CheckTime,omitempty" name:"CheckTime"`
	// 关联主机数

	MachineCount *int64 `json:"MachineCount,omitempty" name:"MachineCount"`
	// 关联pod数

	PodCount *int64 `json:"PodCount,omitempty" name:"PodCount"`
	// 关联service数

	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
	// 漏洞风险

	VulRisk *int64 `json:"VulRisk,omitempty" name:"VulRisk"`
	// 配置风险

	CFGRisk *int64 `json:"CFGRisk,omitempty" name:"CFGRisk"`
	// 体检数

	CheckCount *int64 `json:"CheckCount,omitempty" name:"CheckCount"`
	// 是否核心：1:核心，2:非核心

	IsCore *int64 `json:"IsCore,omitempty" name:"IsCore"`
}

type AssetClusterPod struct {

	// 租户id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 租户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 租户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// pod id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// pod名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// pod创建时间

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 主机id

	MachineId *string `json:"MachineId,omitempty" name:"MachineId"`
	// 主机名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// pod ip

	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`
	// 关联service数

	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
	// 关联容器数

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 是否核心：1:核心，2:非核心

	IsCore *int64 `json:"IsCore,omitempty" name:"IsCore"`
}

type AlarmDiversityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警产品分类

		AlarmDiversity []*ObjectDiversity `json:"AlarmDiversity,omitempty" name:"AlarmDiversity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmDiversityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmDiversityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNeedToSolveListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNeedToSolveListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNeedToSolveListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlagShip struct {

	// 无

	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`
	// 无

	IsApplyFor *bool `json:"IsApplyFor,omitempty" name:"IsApplyFor"`
	// 无

	LicenseNum *int64 `json:"LicenseNum,omitempty" name:"LicenseNum"`
	// 无

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
}

type ServerRisk struct {

	// 测绘标签

	ServiceTag *string `json:"ServiceTag,omitempty" name:"ServiceTag"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 影响资产

	AffectAsset *string `json:"AffectAsset,omitempty" name:"AffectAsset"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 资产类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 服务

	Service *string `json:"Service,omitempty" name:"Service"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 风险详情

	RiskDetails *string `json:"RiskDetails,omitempty" name:"RiskDetails"`
	// 处置建议

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 状态，0未处理、1已处置、2已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 资产唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 服务快照

	ServiceSnapshot *string `json:"ServiceSnapshot,omitempty" name:"ServiceSnapshot"`
	// 服务访问的url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 列表索引值

	Index *string `json:"Index,omitempty" name:"Index"`
	// 风险列表

	RiskList []*ServerRiskSuggestion `json:"RiskList,omitempty" name:"RiskList"`
	// 建议列表

	SuggestionList []*ServerRiskSuggestion `json:"SuggestionList,omitempty" name:"SuggestionList"`
}

type OrganizationUserInfo struct {

	// 成员账号Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 成员账号名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 部门节点名称，账号所属部门

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 资产数量

	AssetCount *int64 `json:"AssetCount,omitempty" name:"AssetCount"`
	// 风险数量

	RiskCount *int64 `json:"RiskCount,omitempty" name:"RiskCount"`
	// 攻击数量

	AttackCount *int64 `json:"AttackCount,omitempty" name:"AttackCount"`
	// Member/Admin;成员或者管理员

	Role *string `json:"Role,omitempty" name:"Role"`
	// 成员账号id

	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
	// 成员账号Appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 账号加入方式,create/invite

	JoinType *string `json:"JoinType,omitempty" name:"JoinType"`
	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本

	CFWProtect *string `json:"CFWProtect,omitempty" name:"CFWProtect"`
	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本

	WAFProtect *string `json:"WAFProtect,omitempty" name:"WAFProtect"`
	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本

	CWPProtect *string `json:"CWPProtect,omitempty" name:"CWPProtect"`
	// 1启用，0未启用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// "Free"       //免费版  "Advanced"   //高级版 "Enterprise" //企业版 "Ultimate"   //旗舰版

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeExaminationScanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 无

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 无

		Data *VSSPayInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExaminationScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExaminationScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePortRiskAdvanceCFGListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribePortRiskAdvanceCFGListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePortRiskAdvanceCFGListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicIpAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*IpAssetListVO `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 资产归属地

		AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitempty" name:"AssetLocationList"`
		// ip列表枚举

		IpTypeList []*FilterDataObject `json:"IpTypeList,omitempty" name:"IpTypeList"`
		// 地域列表枚举

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 防护枚举

		DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitempty" name:"DefenseStatusList"`
		// 资产类型枚举

		AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`
		// AppId枚举

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicIpAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicIpAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeTaskLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCvmInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		AvailableProVersionLicenseCnt *int64 `json:"AvailableProVersionLicenseCnt,omitempty" name:"AvailableProVersionLicenseCnt"`
		// 无

		AvailableFlagshipVersionLicenseCnt *int64 `json:"AvailableFlagshipVersionLicenseCnt,omitempty" name:"AvailableFlagshipVersionLicenseCnt"`
		// 无

		ProVersionLicenseCnt *int64 `json:"ProVersionLicenseCnt,omitempty" name:"ProVersionLicenseCnt"`
		// 无

		FlagshipVersionLicenseCnt *int64 `json:"FlagshipVersionLicenseCnt,omitempty" name:"FlagshipVersionLicenseCnt"`
		// 无

		CanApplyTrial *bool `json:"CanApplyTrial,omitempty" name:"CanApplyTrial"`
		// 无

		LastTrialTime *string `json:"LastTrialTime,omitempty" name:"LastTrialTime"`
		// 无

		FlagShip *FlagShip `json:"FlagShip,omitempty" name:"FlagShip"`
		// 无

		ProfessionalShip *Professional `json:"ProfessionalShip,omitempty" name:"ProfessionalShip"`
		// 阻断状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserCvmInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCvmInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallBASAssetAgentByCWPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命令执行状态码，0代表成功，其他代表失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallBASAssetAgentByCWPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallBASAssetAgentByCWPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PortRiskAdvanceCFGParamItem struct {

	// 端口集合,以逗号分隔

	PortSets *string `json:"PortSets,omitempty" name:"PortSets"`
	// 检测项类型，0-系统定义，1-用户自定义

	CheckType *int64 `json:"CheckType,omitempty" name:"CheckType"`
	// 检测项描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 是否启用，1-启用，0-禁用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type RiskCenterStatusKey struct {

	// 风险ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 公网IP/域名

	PublicIPDomain *string `json:"PublicIPDomain,omitempty" name:"PublicIPDomain"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// APP ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type AssetsApiConfig struct {

	// api子路径

	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`
	// 1: get方法，2: post方法，3： delete方法， 4： patch方法

	Method *int64 `json:"Method,omitempty" name:"Method"`
	// 携带的参数：json形式post/petch模式下是http body的数据，不一定是json形式 ；get/delete模式下是http params的参数，形式如： a=1&b=2&c=3

	Args *string `json:"Args,omitempty" name:"Args"`
}

type CreateWafSpartaProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0表示成功 其他值表示失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 消息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 报错代码

		Code *string `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWafSpartaProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWafSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBASATTCKMatrixRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBASATTCKMatrixRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBASATTCKMatrixRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskPredictTimeRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 扫描资产信息列表

	Assets []*TaskAssetObject `json:"Assets,omitempty" name:"Assets"`
	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填

	ScanAssetType *int64 `json:"ScanAssetType,omitempty" name:"ScanAssetType"`
	// 扫描项目；port/poc/weakpass/webcontent/configrisk

	ScanItem []*string `json:"ScanItem,omitempty" name:"ScanItem"`
	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填

	ScanPlanType *int64 `json:"ScanPlanType,omitempty" name:"ScanPlanType"`
	// 扫描计划详情

	ScanPlanContent *string `json:"ScanPlanContent,omitempty" name:"ScanPlanContent"`
	// ip/域名/url数组

	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitempty" name:"SelfDefiningAssets"`
	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式

	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
	// 请求发起源 vss/csip

	ScanFrom *string `json:"ScanFrom,omitempty" name:"ScanFrom"`
}

func (r *DescribeTaskPredictTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskPredictTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetTrend struct {

	// 时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 互联网暴露面数量

	Internet *int64 `json:"Internet,omitempty" name:"Internet"`
}

type BugInfoDetail struct {

	// 漏洞编号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 漏洞对应pocId

	PatchId *string `json:"PatchId,omitempty" name:"PatchId"`
	// 漏洞名称

	VULName *string `json:"VULName,omitempty" name:"VULName"`
	// 漏洞严重性：high,middle，low，info

	Level *string `json:"Level,omitempty" name:"Level"`
	// cvss评分

	CVSSScore *string `json:"CVSSScore,omitempty" name:"CVSSScore"`
	// cve编号

	CVEId *string `json:"CVEId,omitempty" name:"CVEId"`
	// 漏洞标签

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 漏洞种类，1:web应用，2:系统组件漏洞，3:配置风险

	VULCategory *uint64 `json:"VULCategory,omitempty" name:"VULCategory"`
	// 漏洞影响系统

	ImpactOs *string `json:"ImpactOs,omitempty" name:"ImpactOs"`
	// 漏洞影响组件

	ImpactCOMPENT *string `json:"ImpactCOMPENT,omitempty" name:"ImpactCOMPENT"`
	// 漏洞影响版本

	ImpactVersion *string `json:"ImpactVersion,omitempty" name:"ImpactVersion"`
	// 链接

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 漏洞描述

	VULDescribe *string `json:"VULDescribe,omitempty" name:"VULDescribe"`
	// 修复建议

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// 产品支持状态，实时返回

	ProSupport *uint64 `json:"ProSupport,omitempty" name:"ProSupport"`
	// 是否公开，0为未发布，1为发布

	IsPublish *uint64 `json:"IsPublish,omitempty" name:"IsPublish"`
	// 释放时间

	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 漏洞子类别

	SubCategory *string `json:"SubCategory,omitempty" name:"SubCategory"`
}

type ExportInfo struct {

	// 任务id

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
	// 日志下载状态。Processing:导出正在进行中，Completed:导出完成，Failed:导出失败，Expired:日志导出已过期(三天有效期), Queuing 排队中

	Status *string `json:"Status,omitempty" name:"Status"`
	// 日志导出路径

	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`
	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件大小

	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
}

type SyncAssetsProtectStatusRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 更新指定资产状态，全部更新为空

	Assets []*string `json:"Assets,omitempty" name:"Assets"`
	// 更新的资产类型

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
}

func (r *SyncAssetsProtectStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetsProtectStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCanFixMachine struct {

	// 可修复资产数量

	CanFixMachineNum *int64 `json:"CanFixMachineNum,omitempty" name:"CanFixMachineNum"`
	// 展示可修复资产的quuid用来下发修复任务

	CanFixMachineUuid []*string `json:"CanFixMachineUuid,omitempty" name:"CanFixMachineUuid"`
	// 展示该cveid在主机安全中的VulId用来下发修复任务

	VULId *int64 `json:"VULId,omitempty" name:"VULId"`
	// 修复支持情况：0-windows/linux均不支持修复 ;1-windows/linux 均支持修复 ;2-仅linux支持修复;3-仅windows支持修复

	FixSwitch *int64 `json:"FixSwitch,omitempty" name:"FixSwitch"`
}

type VpcConnectedVpcInfo struct {

	// 租户

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// uin账号

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// cidr信息

	VpcCidr *string `json:"VpcCidr,omitempty" name:"VpcCidr"`
	// 互通方式类型，0：对等连接，1：云联网

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 连接id

	ConnectId *string `json:"ConnectId,omitempty" name:"ConnectId"`
	// 连接名称

	ConnectName *string `json:"ConnectName,omitempty" name:"ConnectName"`
	// 互通id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 互通名称

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 显示edge

	ShowEdge *int64 `json:"ShowEdge,omitempty" name:"ShowEdge"`
}

type WebAssetObject struct {

	// 租户id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 实例内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
}

type DescribeCVMAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Data *AssetBaseInfoResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCVMAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCVMAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportCSVResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回到处文件下载的临时url

		Data *ExportURL `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportCSVResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportCSVResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllExamInfoAndStatus struct {

	// 用户个人信息

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 用户个人信息

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 是否缓存数据，1是缓存，0是正式数据

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// Cfw开启状态，0是关闭，1是开启

	CFW *int64 `json:"CFW,omitempty" name:"CFW"`
	// 版本信息

	CFWVersion *string `json:"CFWVersion,omitempty" name:"CFWVersion"`
	// waf开启状态，0是关闭，1是开启

	Waf *int64 `json:"Waf,omitempty" name:"Waf"`
	// 版本信息

	WafVersion *string `json:"WafVersion,omitempty" name:"WafVersion"`
	// 主机安全的开启状态，0是关闭，1是开启

	CWP *int64 `json:"CWP,omitempty" name:"CWP"`
	// 版本信息

	CWPVersion *string `json:"CWPVersion,omitempty" name:"CWPVersion"`
	// 第一步体检的分数

	StepOneScore *int64 `json:"StepOneScore,omitempty" name:"StepOneScore"`
	// 第一步体检的状态，0未开始，1体检中，2体检完成

	StepOneStatus *int64 `json:"StepOneStatus,omitempty" name:"StepOneStatus"`
	// 公网ip数量

	IpSum *int64 `json:"IpSum,omitempty" name:"IpSum"`
	// 未防护公网ip数量

	UNProtectIpSum *int64 `json:"UNProtectIpSum,omitempty" name:"UNProtectIpSum"`
	// 未防护核心公网ip资产数量

	UNProtectCoreIpSum *int64 `json:"UNProtectCoreIpSum,omitempty" name:"UNProtectCoreIpSum"`
	// 域名资产数量

	DomainSum *int64 `json:"DomainSum,omitempty" name:"DomainSum"`
	// 未防护域名资产数量

	UNProtectDomainSum *int64 `json:"UNProtectDomainSum,omitempty" name:"UNProtectDomainSum"`
	// 未防护核心域名资产数量

	UNProtectCoreDomainSum *int64 `json:"UNProtectCoreDomainSum,omitempty" name:"UNProtectCoreDomainSum"`
	// 主机资产数量

	HostSum *int64 `json:"HostSum,omitempty" name:"HostSum"`
	// 未防护主机资产数量

	UNProtectHostSum *int64 `json:"UNProtectHostSum,omitempty" name:"UNProtectHostSum"`
	// 未防护核心主机资产数量

	UNProtectCoreHostSum *int64 `json:"UNProtectCoreHostSum,omitempty" name:"UNProtectCoreHostSum"`
	// 第二步体检得分

	StepTwoScore *int64 `json:"StepTwoScore,omitempty" name:"StepTwoScore"`
	// 第二步体检状态，0未开始，1体检中，2体检结束

	StepTwoStatus *int64 `json:"StepTwoStatus,omitempty" name:"StepTwoStatus"`
	// 风险端口数量

	PortRiskSum *int64 `json:"PortRiskSum,omitempty" name:"PortRiskSum"`
	// 未处理高危端口数量

	UNDealHighRiskPortSum *int64 `json:"UNDealHighRiskPortSum,omitempty" name:"UNDealHighRiskPortSum"`
	// 存在端口风险的资产数量

	PortRiskAssetSum *int64 `json:"PortRiskAssetSum,omitempty" name:"PortRiskAssetSum"`
	// 存在高危端口的资产数量

	HighRiskPortAssetSum *int64 `json:"HighRiskPortAssetSum,omitempty" name:"HighRiskPortAssetSum"`
	// 	漏洞风险数量

	BugRiskSum *int64 `json:"BugRiskSum,omitempty" name:"BugRiskSum"`
	// 未处理高危漏洞风险数量

	UNDealHighRiskBug *int64 `json:"UNDealHighRiskBug,omitempty" name:"UNDealHighRiskBug"`
	// 存在漏洞风险的资产数量

	BugRiskAssetSum *int64 `json:"BugRiskAssetSum,omitempty" name:"BugRiskAssetSum"`
	// 存在高危漏洞风险的资产数量

	HighRiskBugAssetSum *int64 `json:"HighRiskBugAssetSum,omitempty" name:"HighRiskBugAssetSum"`
	// 弱口令风险数量

	WeakPasswordSum *int64 `json:"WeakPasswordSum,omitempty" name:"WeakPasswordSum"`
	// 存在弱口令风险的资产数量

	WeakPasswordAsset *int64 `json:"WeakPasswordAsset,omitempty" name:"WeakPasswordAsset"`
	// 第三步体检得分

	StepThreeScore *int64 `json:"StepThreeScore,omitempty" name:"StepThreeScore"`
	// 第三步体检状态，0未开始，1体检中，2体检结束

	StepThreeStatus *int64 `json:"StepThreeStatus,omitempty" name:"StepThreeStatus"`
	// 攻击告警数量

	AlertSum *int64 `json:"AlertSum,omitempty" name:"AlertSum"`
	// 未处理攻击告警数量

	UNDealAlertSum *int64 `json:"UNDealAlertSum,omitempty" name:"UNDealAlertSum"`
	// 有高危告警的资产数量

	UNDealAlertAsset *int64 `json:"UNDealAlertAsset,omitempty" name:"UNDealAlertAsset"`
	// 云上资产数量

	AssetSum *int64 `json:"AssetSum,omitempty" name:"AssetSum"`
	// 第四步体检得分

	StepFourScore *int64 `json:"StepFourScore,omitempty" name:"StepFourScore"`
	// 第四步状态，0未开始，1体检中，2体检结束

	StepFourStatus *int64 `json:"StepFourStatus,omitempty" name:"StepFourStatus"`
	// 最终得分

	LastScore *int64 `json:"LastScore,omitempty" name:"LastScore"`
	// 体检总状态，0未开始，1体检中，2体检结束

	LastStatus *int64 `json:"LastStatus,omitempty" name:"LastStatus"`
	// 体检百分比进度

	Percentage *int64 `json:"Percentage,omitempty" name:"Percentage"`
	// 预估体检时间

	EstimatedTime *int64 `json:"EstimatedTime,omitempty" name:"EstimatedTime"`
	// 数据插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 漏扫查询用taskid

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 用户超越多少数量，小数点后一位

	Rank *float64 `json:"Rank,omitempty" name:"Rank"`
	// 弱口令总数

	WeakPasswordAllSum *int64 `json:"WeakPasswordAllSum,omitempty" name:"WeakPasswordAllSum"`
}

type ToDoListReq struct {

	// 结构参数

	Type []*int64 `json:"Type,omitempty" name:"Type"`
}

type DescribeAssetConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api配置

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 扫描类型枚举

		TypeLists []*string `json:"TypeLists,omitempty" name:"TypeLists"`
		// 扫描类型枚举

		AuthMethodLists []*string `json:"AuthMethodLists,omitempty" name:"AuthMethodLists"`
		// 所有配置的端口合集

		PortSets *string `json:"PortSets,omitempty" name:"PortSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFWAssetStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCFWAssetStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWAssetStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchBugInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞信息和资产信息

		Data *DataSearchBug `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSearchBugInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchBugInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPayInfoRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// vss/csip,默认vss

	ScanFrom *string `json:"ScanFrom,omitempty" name:"ScanFrom"`
}

func (r *DescribeUserPayInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPayInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreOrRetryScanTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// Ignore忽略  Retry重试

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
}

func (r *ModifyIgnoreOrRetryScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreOrRetryScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNeedToSolveListRequest struct {
	*tchttp.BaseRequest

	// 1代表操作了资产防护，2代表操作了风险处理，3代表操作了告警处置

	Type []*int64 `json:"Type,omitempty" name:"Type"`
}

func (r *UpdateNeedToSolveListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNeedToSolveListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tags struct {

	// 无

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 无

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type AssetViewWeakPassRisk struct {

	// 影响资产

	AffectAsset *string `json:"AffectAsset,omitempty" name:"AffectAsset"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 资产类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 服务

	Service *string `json:"Service,omitempty" name:"Service"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 状态，0未处理、1已处置、2已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 资产唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 前端索引

	Index *string `json:"Index,omitempty" name:"Index"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 弱口令类型

	PasswordType *string `json:"PasswordType,omitempty" name:"PasswordType"`
	// 来源

	From *string `json:"From,omitempty" name:"From"`
	// 漏洞类型

	VULType *string `json:"VULType,omitempty" name:"VULType"`
	// 漏洞url

	VULURL *string `json:"VULURL,omitempty" name:"VULURL"`
	// 修复建议

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// 负载

	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

type RecentTaskList struct {

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义

	ScanPlanType *int64 `json:"ScanPlanType,omitempty" name:"ScanPlanType"`
	// corn

	ScanPlanContent *string `json:"ScanPlanContent,omitempty" name:"ScanPlanContent"`
	// 状态

	ScanStatus *int64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 任务进度

	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
	// 任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 异常码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 异常信息

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// APP ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// UIN

	UIN *string `json:"UIN,omitempty" name:"UIN"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type TaskCenterVulRiskInputParam struct {

	// 风险ID

	RiskId *string `json:"RiskId,omitempty" name:"RiskId"`
	// 是否开启，0-不开启，1-开启

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type DescribeAlarmStatisticsRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAlarmStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*AssetObject `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产分类列表

		AssetTypeLists []*FilterDataObject `json:"AssetTypeLists,omitempty" name:"AssetTypeLists"`
		// 资产类型列表

		InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`
		// 地域列表

		RegionLists []*FilterDataObject `json:"RegionLists,omitempty" name:"RegionLists"`
		// 主机总数

		InstanceTotal *int64 `json:"InstanceTotal,omitempty" name:"InstanceTotal"`
		// 数据库总数

		DbTotal *int64 `json:"DbTotal,omitempty" name:"DbTotal"`
		// 域名总数

		DomainTotal *int64 `json:"DomainTotal,omitempty" name:"DomainTotal"`
		// 公网ip总数

		PublicIPTotal *int64 `json:"PublicIPTotal,omitempty" name:"PublicIPTotal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关资产详情

		Data *GatewayDetailInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndeterminacyAssetInfoRequest struct {
	*tchttp.BaseRequest

	// -

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeIndeterminacyAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndeterminacyAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterRiskOverviewRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterRiskOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterRiskOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskCenterRiskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRiskCenterRiskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskCenterRiskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataWafVersion struct {

	// 0未开通，1已开通

	CLBWaf *int64 `json:"CLBWaf,omitempty" name:"CLBWaf"`
	// 0未开通，1已开通

	SAASWaf *int64 `json:"SAASWaf,omitempty" name:"SAASWaf"`
	// 返回版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type CreateSolidProtectionTrialsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 选择开通试用产品状态返回

		Data []*SolidProtectionSingleApplyTrial `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSolidProtectionTrialsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolidProtectionTrialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CWPParams struct {

	// 主机选取设置，系统推荐选1，用户自选绑定主机选2，选2必须在HostListQuuid中带有quuid

	HostOption *int64 `json:"HostOption,omitempty" name:"HostOption"`
	// 自选则传入需要开通的主机的quuid数组形式

	HostListUID []*string `json:"HostListUID,omitempty" name:"HostListUID"`
}

type List struct {

	// 云服务器UUID

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 0 执行中, 1 成功,2失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 修复建议

	FixMessage *string `json:"FixMessage,omitempty" name:"FixMessage"`
}

type DescribeCanFixMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 输出vulid,quuid,和可修复资产数量

		Data *SearchCanFixMachine `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCanFixMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCanFixMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebServiceAssetInfo struct {

	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// Nick

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// IP_port

	IPPort *string `json:"IPPort,omitempty" name:"IPPort"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 服务类型

	NetworkType *int64 `json:"NetworkType,omitempty" name:"NetworkType"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 资产ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 组件

	Module *string `json:"Module,omitempty" name:"Module"`
	// 网络访问

	NetworkAccess *int64 `json:"NetworkAccess,omitempty" name:"NetworkAccess"`
	// BOT访问

	BotVisit *int64 `json:"BotVisit,omitempty" name:"BotVisit"`
	// 网络攻击

	NetworkAttack *int64 `json:"NetworkAttack,omitempty" name:"NetworkAttack"`
	// 弱口令风险

	WeekPasswordCount *int64 `json:"WeekPasswordCount,omitempty" name:"WeekPasswordCount"`
	// 内容风险

	WebContentCount *int64 `json:"WebContentCount,omitempty" name:"WebContentCount"`
	// waf/防火墙防护

	ProtectType *string `json:"ProtectType,omitempty" name:"ProtectType"`
	// 扫描任务

	ScanTask *int64 `json:"ScanTask,omitempty" name:"ScanTask"`
	// 最近扫描

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 资产发现

	DiscoverTime *string `json:"DiscoverTime,omitempty" name:"DiscoverTime"`
	// 资产状态：1，有效；2，无效

	InstanceStatus *int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
	// 云防防护类型：-1 未防护，0:基础防护，1:高级防护

	WebProtectType *int64 `json:"WebProtectType,omitempty" name:"WebProtectType"`
	// web服务唯一id

	WebId *int64 `json:"WebId,omitempty" name:"WebId"`
	// 资产

	Asset *string `json:"Asset,omitempty" name:"Asset"`
	// Manual手动等

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
}

type ACRuleItem struct {

	// APP ID

	TargetAppid *string `json:"TargetAppid,omitempty" name:"TargetAppid"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 策略, 0：观察，1：阻断，2：放行

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 访问源类型，1是IP，3是域名，4是IP地址模版，5是域名地址模版

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 访问目的类型，1是IP，3是域名，4是IP地址模版，5是域名地址模版

	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`
	// id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 日志id，从告警处创建必传，其它为空

	LogId *string `json:"LogId,omitempty" name:"LogId"`
	// 城市Code

	City *uint64 `json:"City,omitempty" name:"City"`
	// 国家Code

	Country *uint64 `json:"Country,omitempty" name:"Country"`
	// 云厂商，支持多个，以逗号分隔， 1:腾讯云（仅中国香港及海外）,2:阿里云,3:亚马逊云,4:华为云,5:微软云

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// 是否为地域

	IsRegion *uint64 `json:"IsRegion,omitempty" name:"IsRegion"`
	// 城市名

	CityName *string `json:"CityName,omitempty" name:"CityName"`
	// 国家名

	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
}

type GroupNodeInfo struct {

	// 节点id

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 资产数量

	InsNum *int64 `json:"InsNum,omitempty" name:"InsNum"`
	// 节点其他信息

	NodeInfo *string `json:"NodeInfo,omitempty" name:"NodeInfo"`
}

type PublicIpAssetInfo struct {

	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 绑定资产id

	BoundAssetId *string `json:"BoundAssetId,omitempty" name:"BoundAssetId"`
	// 绑定资产名

	BoundAssetName *string `json:"BoundAssetName,omitempty" name:"BoundAssetName"`
	// 绑定资产类别

	BoundAssetType *string `json:"BoundAssetType,omitempty" name:"BoundAssetType"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// 最近扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 公网ip类型 0 未知,1 公网,2 弹性,3 弹性ipv6,4 anycastIP

	PublicIpType *uint64 `json:"PublicIpType,omitempty" name:"PublicIpType"`
	// 防护版本名称 0、未知 1、旗舰版 2、企业版 3、高级版 4、IPS版 5、普惠版

	CfwProtectLevel *int64 `json:"CfwProtectLevel,omitempty" name:"CfwProtectLevel"`
	// 云防防火墙防护状态 1 防护中 0 未防护

	IsProtected *int64 `json:"IsProtected,omitempty" name:"IsProtected"`
	// 地域名称

	Region *string `json:"Region,omitempty" name:"Region"`
	// ip地理位置信息

	IpLocation *string `json:"IpLocation,omitempty" name:"IpLocation"`
	// 云防火墙防护状态

	CFWStatus *int64 `json:"CFWStatus,omitempty" name:"CFWStatus"`
	// 弹性网卡id

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
}

type ModifyCheckUpResultRequest struct {
	*tchttp.BaseRequest

	// 步骤，1 安全防线建设 2 业务资产梳理 3 资产风险发现 4 高危告警处理

	Step *uint64 `json:"Step,omitempty" name:"Step"`
}

func (r *ModifyCheckUpResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCheckUpResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTkeAuthRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *CreateTkeAuthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTkeAuthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBASATTCKMatrixResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 战术战技矩阵

		Tactics []*BASTacticsObject `json:"Tactics,omitempty" name:"Tactics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBASATTCKMatrixResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBASATTCKMatrixResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServiceAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeClusterServiceAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFreeChanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data *FreeLimitResp `json:"Data,omitempty" name:"Data"`
		// 接口请求信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 接口请求状态

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFreeChanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFreeChanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchExamStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSearchExamStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchExamStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafSaasPortsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeWafSaasPortsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafSaasPortsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetLocation struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 域名资产数量

	DomainAssetCount *uint64 `json:"DomainAssetCount,omitempty" name:"DomainAssetCount"`
	// 公网IP资产数量

	PublicIPAssetCount *uint64 `json:"PublicIPAssetCount,omitempty" name:"PublicIPAssetCount"`
}

type BASIDNamePair struct {

	// 战术/战技ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 战术战技名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type RiskCenterOverviewTrend struct {

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 端口数

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 漏洞数

	VUL *int64 `json:"VUL,omitempty" name:"VUL"`
	// 弱口令数

	WeakPassword *int64 `json:"WeakPassword,omitempty" name:"WeakPassword"`
	// 网站数

	Website *int64 `json:"Website,omitempty" name:"Website"`
	// 配置数

	CFG *int64 `json:"CFG,omitempty" name:"CFG"`
	// 测绘风险数

	Server *int64 `json:"Server,omitempty" name:"Server"`
}

type Filters struct {

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 模糊匹配

	ExactMatch *string `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type DescribeUserCVEStatusRequest struct {
	*tchttp.BaseRequest

	// 漏洞CVEid编号

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
}

func (r *DescribeUserCVEStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCVEStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCVEStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Data *SearchWafPatch `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserCVEStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCVEStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBASTaskTechRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *GetBASTaskTechRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBASTaskTechRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetObject struct {

	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 	资产类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 资产分类

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// ip/域名/资产id，数据库id等

	Asset *string `json:"Asset,omitempty" name:"Asset"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// bas安装状态
	// 0: 未安装
	// 1: 已安装
	// -1: 不支持

	BASAgentStatus *int64 `json:"BASAgentStatus,omitempty" name:"BASAgentStatus"`
	// 镜像授权状态,0:未授权，1:已授权

	ImageAuth *int64 `json:"ImageAuth,omitempty" name:"ImageAuth"`
	// 容器组件状态

	ClusterComponentStatus *string `json:"ClusterComponentStatus,omitempty" name:"ClusterComponentStatus"`
}

type ObjectDiversity struct {

	// 父信息

	ObjectInfo *NodeInfo `json:"ObjectInfo,omitempty" name:"ObjectInfo"`
	// 子信息

	SubNodeInfo []*NodeInfo `json:"SubNodeInfo,omitempty" name:"SubNodeInfo"`
}

type WeakPasswordViewWeakPassRisk struct {

	// 影响资产

	NoHandleCount *int64 `json:"NoHandleCount,omitempty" name:"NoHandleCount"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 服务

	Service *string `json:"Service,omitempty" name:"Service"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 状态，0未处理、1已处置、2已忽略

	AffectAssetCount *uint64 `json:"AffectAssetCount,omitempty" name:"AffectAssetCount"`
	// 资产唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 资产子类型

	From *string `json:"From,omitempty" name:"From"`
	// 前端索引

	Index *string `json:"Index,omitempty" name:"Index"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 弱口令类型

	PasswordType *string `json:"PasswordType,omitempty" name:"PasswordType"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type Professional struct {

	// 无

	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`
	// 无

	IsApplyFor *bool `json:"IsApplyFor,omitempty" name:"IsApplyFor"`
	// 无

	LicenseNum *int64 `json:"LicenseNum,omitempty" name:"LicenseNum"`
	// 无

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
}

type TaskOverView struct {

	// 总任务数

	TaskNumber *int64 `json:"TaskNumber,omitempty" name:"TaskNumber"`
	// 待开始扫描任务数

	WaitTaskNumber *int64 `json:"WaitTaskNumber,omitempty" name:"WaitTaskNumber"`
	// 正在扫描任务数

	TaskNowNumber *int64 `json:"TaskNowNumber,omitempty" name:"TaskNowNumber"`
	// 报告数量

	TaskLogNumber *int64 `json:"TaskLogNumber,omitempty" name:"TaskLogNumber"`
	// 任务信息

	TaskScanList []*RecentTaskList `json:"TaskScanList,omitempty" name:"TaskScanList"`
	// 扫描报告信息

	TaskLogList []*TaskLogSimpleInfo `json:"TaskLogList,omitempty" name:"TaskLogList"`
	// 异常任务总数

	ErrTaskNumber *int64 `json:"ErrTaskNumber,omitempty" name:"ErrTaskNumber"`
}

type WebContentEvidenceInfo struct {

	// 文件url

	URL *string `json:"URL,omitempty" name:"URL"`
	// 文件格式

	FileFormat *string `json:"FileFormat,omitempty" name:"FileFormat"`
}

type AddNewBindRoleUserRequest struct {
	*tchttp.BaseRequest

	// 是否发起扫描，0不发起，1发起

	IsScan *int64 `json:"IsScan,omitempty" name:"IsScan"`
	// 用户来源,waf/cwp/cfw/ssa/vss

	From *string `json:"From,omitempty" name:"From"`
}

func (r *AddNewBindRoleUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNewBindRoleUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateACRulesRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 创建规则数据

	Data []*ACRuleItem `json:"Data,omitempty" name:"Data"`
	// 0：添加（默认），1：插入

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 边id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 访问控制规则状态

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 0：添加，1：覆盖

	Overwrite *uint64 `json:"Overwrite,omitempty" name:"Overwrite"`
	// NAT实例ID, 参数Area存在的时候这个必传

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// portScan: 来自于端口扫描, patchImport: 来自于批量导入

	From *string `json:"From,omitempty" name:"From"`
	// NAT地域

	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *CreateACRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateACRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterMachineAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeClusterMachineAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterMachineAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FreeLimitUserType struct {

	// 1是云防用户，2是尊贵的appid用户，3是当月抢到免费体检资格的用户

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// 该类型用户是否已使用对应的免费体检额度

	IsUseScan *bool `json:"IsUseScan,omitempty" name:"IsUseScan"`
	// 是否是该类型用户

	IsThisType *bool `json:"IsThisType,omitempty" name:"IsThisType"`
}

type AddWebServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddWebServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddWebServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBASScriptRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤filter

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeBASScriptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBASScriptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFWRuleOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`
		// 阻断策略规则数量

		StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`
		// 启用规则数量

		StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`
		// 停用规则数量

		StopRuleNum *uint64 `json:"StopRuleNum,omitempty" name:"StopRuleNum"`
		// 剩余配额

		RemainingNum *int64 `json:"RemainingNum,omitempty" name:"RemainingNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCFWRuleOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCoreAssetRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeCoreAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCoreAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllAsset struct {

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 私网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 私网名

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 端口风险

	PortRisk *int64 `json:"PortRisk,omitempty" name:"PortRisk"`
	// 漏洞风险

	VulnerabilityRisk *int64 `json:"VulnerabilityRisk,omitempty" name:"VulnerabilityRisk"`
	// vpc-id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc-name

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 私网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 2:公网 3:私网

	PublicIpType *int64 `json:"PublicIpType,omitempty" name:"PublicIpType"`
	// 租户昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

type DomainAssetInfo struct {

	// 子域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 资产创建时间

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
	// waf防护版本

	ProtectLevel *int64 `json:"ProtectLevel,omitempty" name:"ProtectLevel"`
	// WAF防护状态  1 防护中 0 未防护

	IsProtected *uint64 `json:"IsProtected,omitempty" name:"IsProtected"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// waf接入状态

	WAFState *uint64 `json:"WAFState,omitempty" name:"WAFState"`
	// waf状态

	WAFStatus *uint64 `json:"WAFStatus,omitempty" name:"WAFStatus"`
	// 解析ip

	ServerIp []*string `json:"ServerIp,omitempty" name:"ServerIp"`
}

type DescribeAssetStatisticsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeAssetStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterMachineAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群主机节点

		Data []*ClusterMachineVO `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 资产类型

		AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`
		// 操作系统

		OsList []*FilterDataObject `json:"OsList,omitempty" name:"OsList"`
		// vpc列表

		VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterMachineAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterMachineAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Data *DomainAssetInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogAuditResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据列表

		Data []*LogAuditData `json:"Data,omitempty" name:"Data"`
		// 返回数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回代码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 结果信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogAuditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标签内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeWafTlsVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本信息

		Data []*WafTlsVersionInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafTlsVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafTlsVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWebServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopBASTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 停止的taskid数组

	TaskID []*string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *StopBASTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopBASTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GatewayDetailInfo struct {

	// 资产Id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 私有Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// 资产最后扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// tag标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
}

type DescribeResourceTagGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// tag信息列表

		Data []*GroupLastNumThreeNodeInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceTagGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTagGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskAssetTopFiveListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeRiskAssetTopFiveListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskAssetTopFiveListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogURLRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务报告Id 列表

	ReportItemKeyList []*ReportItemKey `json:"ReportItemKeyList,omitempty" name:"ReportItemKeyList"`
	// 0: 预览， 1: 下载

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeTaskLogURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskLogURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LocalImageVO struct {

	// 用户appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// nicke名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 镜像id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 镜像名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 资产创建时间

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
	// 资产大小-kb

	InstanceSize *string `json:"InstanceSize,omitempty" name:"InstanceSize"`
	// 主机数量

	MachineCount *int64 `json:"MachineCount,omitempty" name:"MachineCount"`
	// 容器数量

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 构建数量

	BuildCount *int64 `json:"BuildCount,omitempty" name:"BuildCount"`
	// 组件数量

	ComponentCount *int64 `json:"ComponentCount,omitempty" name:"ComponentCount"`
	// 镜像授权状态

	AuthStatus *int64 `json:"AuthStatus,omitempty" name:"AuthStatus"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
	// 漏洞风险

	VulRisk *int64 `json:"VulRisk,omitempty" name:"VulRisk"`
	// 配置风险

	CfgRisk *int64 `json:"CfgRisk,omitempty" name:"CfgRisk"`
	// 体检任务数

	CheckCount *int64 `json:"CheckCount,omitempty" name:"CheckCount"`
	// 体检时间

	CheckTime *string `json:"CheckTime,omitempty" name:"CheckTime"`
}

type WhereFilter struct {

	// 过滤的项

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤的值

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 精确匹配填 7 模糊匹配填9 ， 兼容 中台定的结构

	OperatorType *int64 `json:"OperatorType,omitempty" name:"OperatorType"`
}

type DescribeHighLevelRiskTopFiveListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 高危TOP5列表

		Data []*HighLevelRiskData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHighLevelRiskTopFiveListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHighLevelRiskTopFiveListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWafVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NeedToSolveList struct {

	// 该类型枚举值为1，2，3

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 当主类型是1时，枚举值是ip，domain，host，当主类型是2时，枚举值是port，bug，weakpasswd，当主类型是3时，这个值是告警名称，方便前端查询锚定

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 详细的待办信息，前端直接拿来展示即可，前两个数据是方便联动

	Message *string `json:"Message,omitempty" name:"Message"`
}

type WebsiteRisk struct {

	// 影响资产

	AffectAsset *string `json:"AffectAsset,omitempty" name:"AffectAsset"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 状态，0未处理、1已处置、2已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 资产唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 前端索引

	Index *string `json:"Index,omitempty" name:"Index"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 风险链接

	URL *string `json:"URL,omitempty" name:"URL"`
	// 风险文件地址

	URLPath *string `json:"URLPath,omitempty" name:"URLPath"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 类型

	DetectEngine *string `json:"DetectEngine,omitempty" name:"DetectEngine"`
	// 结果描述

	ResultDescribe *string `json:"ResultDescribe,omitempty" name:"ResultDescribe"`
	// 源地址url

	SourceURL *string `json:"SourceURL,omitempty" name:"SourceURL"`
	// 源文件地址

	SourceURLPath *string `json:"SourceURLPath,omitempty" name:"SourceURLPath"`
}

type DeleteAssetConfigRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 资产配置的id列表

	ConfigIds []*string `json:"ConfigIds,omitempty" name:"ConfigIds"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
}

func (r *DeleteAssetConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAssetConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBASTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeBASTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBASTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskReportOverviewRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// vss/csip,默认vss

	ScanFrom *string `json:"ScanFrom,omitempty" name:"ScanFrom"`
}

func (r *DescribeTaskReportOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskReportOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateProductStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSendVoucherToUserTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// code 0为成功 其他值为失败

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 消息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSendVoucherToUserTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSendVoucherToUserTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePDFDownLoadLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 报告下载记录列表

		Data []*PDFDownLoadRecord `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePDFDownLoadLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePDFDownLoadLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanAssetByTaskIdRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
}

func (r *DescribeScanAssetByTaskIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanAssetByTaskIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ABTestConfig struct {

	// 灰度项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// true：正在灰度，false：不在灰度

	Status *bool `json:"Status,omitempty" name:"Status"`
}

type ClusterMachineVO struct {

	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 用户昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 主机节点id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 主机节点资产类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 主机节点操作系统

	InstanceOs *string `json:"InstanceOs,omitempty" name:"InstanceOs"`
	// VPC

	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`
	// VPC自定义名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 资产创建时间

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
	// 容器数量

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 镜像数量

	ImageCount *int64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 关联集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 关联集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// docker版本

	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
	// 容器版本

	ContainerVersion *string `json:"ContainerVersion,omitempty" name:"ContainerVersion"`
	// 用户uin信息

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 私网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
	// 最近上线时间

	LastOnlineTime *string `json:"LastOnlineTime,omitempty" name:"LastOnlineTime"`
}

type DescribeCVEInfoRequest struct {
	*tchttp.BaseRequest

	// 漏洞CVE编号

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
}

func (r *DescribeCVEInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCVEInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集团用户列表

		Data []*OrganizationUserInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BugInfoTrend struct {

	// 时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 数据

	Bug *int64 `json:"Bug,omitempty" name:"Bug"`
}

type DescribeResourceTagGroupsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeResourceTagGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTagGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤参数

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeSubnetAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤参数

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeVpcAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机器列表

		Machines []*CWPMachine `json:"Machines,omitempty" name:"Machines"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterRiskTrendListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势列表

		Data []*RiskCenterOverviewTrend `json:"Data,omitempty" name:"Data"`
		// 日期类型列表

		DateTypeLists []*FilterDataObject `json:"DateTypeLists,omitempty" name:"DateTypeLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterRiskTrendListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterRiskTrendListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskLogInfo struct {

	// 报告名称

	TaskLogName *string `json:"TaskLogName,omitempty" name:"TaskLogName"`
	// 报告ID

	TaskLogId *string `json:"TaskLogId,omitempty" name:"TaskLogId"`
	// 关联资产个数

	AssetsNumber *int64 `json:"AssetsNumber,omitempty" name:"AssetsNumber"`
	// 安全风险数量

	RiskNumber *int64 `json:"RiskNumber,omitempty" name:"RiskNumber"`
	// 报告生成时间,任务结束时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 任务状态码：0 初始值  1正在扫描  2扫描完成  3扫描出错，4停止，5暂停，6该任务已被重启过

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 关联任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 扫描开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务中心扫描任务ID

	TaskCenterTaskId *string `json:"TaskCenterTaskId,omitempty" name:"TaskCenterTaskId"`
	// 租户ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 主账户ID

	UIN *string `json:"UIN,omitempty" name:"UIN"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DescribeCVEInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Data *SearchWafPatch `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCVEInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCVEInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskCostQuotaRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 扫描资产信息列表

	Assets []*TaskAssetObject `json:"Assets,omitempty" name:"Assets"`
	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填

	ScanAssetType *int64 `json:"ScanAssetType,omitempty" name:"ScanAssetType"`
	// 扫描项目；port/poc/weakpass/webcontent/configrisk

	ScanItem []*string `json:"ScanItem,omitempty" name:"ScanItem"`
	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填

	ScanPlanType *int64 `json:"ScanPlanType,omitempty" name:"ScanPlanType"`
	// 扫描计划详情

	ScanPlanContent *string `json:"ScanPlanContent,omitempty" name:"ScanPlanContent"`
	// ip/域名/url数组

	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitempty" name:"SelfDefiningAssets"`
	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式

	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
	// 扫描来源 vss/csip

	ScanFrom *string `json:"ScanFrom,omitempty" name:"ScanFrom"`
}

func (r *DescribeTaskCostQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskCostQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderQuotaInfo struct {

	// 配额键

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 配额总量

	QuotaNum *int64 `json:"QuotaNum,omitempty" name:"QuotaNum"`
	// 配额已使用量

	QuotaUsed *int64 `json:"QuotaUsed,omitempty" name:"QuotaUsed"`
}

type DescribeClusterServiceAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*AssetClusterService `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群服务类型枚举

		ServiceTypeList []*FilterDataObject `json:"ServiceTypeList,omitempty" name:"ServiceTypeList"`
		// 命名枚举

		NamespaceList []*FilterDataObject `json:"NamespaceList,omitempty" name:"NamespaceList"`
		// 地域枚举

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 租户枚举

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterServiceAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNICAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网卡详情

		Data *NICDetailInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNICAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNICAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务日志列表

		Data []*ScanTaskInfoList `json:"Data,omitempty" name:"Data"`
		// 主账户ID列表

		UINList []*string `json:"UINList,omitempty" name:"UINList"`
		// 体检模式过滤列表

		TaskModeList []*FilterDataObject `json:"TaskModeList,omitempty" name:"TaskModeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupLastNumFourNodeInfo struct {

	// 节点信息

	NodeInfo *GroupNodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`
	// 第三层节点信息

	SubNodeList []*GroupLastNumThreeNodeInfo `json:"SubNodeList,omitempty" name:"SubNodeList"`
}

type PublicIpDomainListKey struct {

	// 资产值

	Asset *string `json:"Asset,omitempty" name:"Asset"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
}

type DescribeBASATTCKListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBASATTCKListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBASATTCKListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePocScanMissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最终结果

		Data *CreateScanResp `json:"Data,omitempty" name:"Data"`
		// 接口请求状态

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口请求信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePocScanMissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePocScanMissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 业务标签: 1:云防火墙，2:WAF，3:主机安全，4:SOC

	Tags []*int64 `json:"Tags,omitempty" name:"Tags"`
}

func (r *DeleteLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRiskScanTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务id 列表

	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitempty" name:"TaskIdList"`
}

func (r *DeleteRiskScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRiskScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRiskScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRiskScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDiversityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产分类信息

		ObjectDiversity *ObjectDiversity `json:"ObjectDiversity,omitempty" name:"ObjectDiversity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDiversityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDiversityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetVpcDetailInfoRequest struct {
	*tchttp.BaseRequest

	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeAssetVpcDetailInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetVpcDetailInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndeterminacyAssetTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// -

		Data []*AssetDetailTrendInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndeterminacyAssetTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndeterminacyAssetTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogAuditRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 查询类型

	QueryType *int64 `json:"QueryType,omitempty" name:"QueryType"`
	// 查询日志分类, 前一个数字代表产品，后一个数字代表日志类型

	LogType []*string `json:"LogType,omitempty" name:"LogType"`
	// 查询字段名称, 为空查全部

	Field *string `json:"Field,omitempty" name:"Field"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 返回页高

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回页刻度

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 逻辑表达式

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 柱状图间隔

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 集群 0:国内 1:国际 2:全部

	Cluster *int64 `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DescribeLogAuditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogAuditRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBASScriptDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 剧本战术详细

		Tactics []*BASTacticsObject `json:"Tactics,omitempty" name:"Tactics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBASScriptDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBASScriptDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPortProtectStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPortProtectStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPortProtectStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {

	// id信息

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 无

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 无

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 无

	LicenseType *int64 `json:"LicenseType,omitempty" name:"LicenseType"`
}

type DescribePortRiskAdvanceCFGListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置项列表

		Data []*PortRiskAdvanceCFGList `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 检测项过滤列表

		CheckTypeLists []*FilterDataObject `json:"CheckTypeLists,omitempty" name:"CheckTypeLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePortRiskAdvanceCFGListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePortRiskAdvanceCFGListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAuditTrendRequest struct {
	*tchttp.BaseRequest

	// 开始时间 粒度为天的格式yyyy-mm-dd

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间 粒度为天的格式yyyy-mm-dd
	//

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeUserAuditTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAuditTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserOperationPermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*OperationPermission `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserOperationPermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserOperationPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLicenseBindsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLicenseBindsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLicenseBindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNeedToSolveListChangeRequest struct {
	*tchttp.BaseRequest

	// 插入新的待办操作记录所需参数

	Param []*ParamInfo `json:"Param,omitempty" name:"Param"`
	// 步骤

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *UpdateNeedToSolveListChangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNeedToSolveListChangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProtectTimeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProtectTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterAssetViewCFGRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产视角的配置风险列表

		Data []*AssetViewCFGRisk `json:"Data,omitempty" name:"Data"`
		// 状态列表

		StatusLists []*FilterDataObject `json:"StatusLists,omitempty" name:"StatusLists"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 配置名列表

		CFGNameLists []*FilterDataObject `json:"CFGNameLists,omitempty" name:"CFGNameLists"`
		// 检查类型列表

		CheckTypeLists []*FilterDataObject `json:"CheckTypeLists,omitempty" name:"CheckTypeLists"`
		// 资产类型列表

		InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`
		// 来源列表

		FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewCFGRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterAssetViewCFGRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSolidAssetStaticResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 立体防护资产数据

		Data *SolidProtectionAssetData `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSolidAssetStaticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSolidAssetStaticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageRateRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeStorageRateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStorageRateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafCiphersDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本信息

		Data []*WafTLSCipherInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafCiphersDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafCiphersDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBASScriptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模拟攻击剧本总数

		TotalCnt *uint64 `json:"TotalCnt,omitempty" name:"TotalCnt"`
		// 默认剧本总数

		TotalDefaultCnt *uint64 `json:"TotalDefaultCnt,omitempty" name:"TotalDefaultCnt"`
		// 自定义剧本总数

		TotalUserCnt *uint64 `json:"TotalUserCnt,omitempty" name:"TotalUserCnt"`
		// 剧本描述

		Scripts []*BASScriptObject `json:"Scripts,omitempty" name:"Scripts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBASScriptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBASScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeakPWDRiskAdvanceCFGListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 筛选内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeWeakPWDRiskAdvanceCFGListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeakPWDRiskAdvanceCFGListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetBaseInfoResponse struct {

	// vpc-id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc-name

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 操作系统

	Os *string `json:"Os,omitempty" name:"Os"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 账号数量

	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`
	// 端口数量

	PortNum *uint64 `json:"PortNum,omitempty" name:"PortNum"`
	// 进程数量

	ProcessNum *uint64 `json:"ProcessNum,omitempty" name:"ProcessNum"`
	// 软件应用数量

	SoftApplicationNum *uint64 `json:"SoftApplicationNum,omitempty" name:"SoftApplicationNum"`
	// 数据库数量

	DatabaseNum *uint64 `json:"DatabaseNum,omitempty" name:"DatabaseNum"`
	// Web应用数量

	WebApplicationNum *uint64 `json:"WebApplicationNum,omitempty" name:"WebApplicationNum"`
	// 服务数量

	ServiceNum *uint64 `json:"ServiceNum,omitempty" name:"ServiceNum"`
	// web框架数量

	WebFrameworkNum *uint64 `json:"WebFrameworkNum,omitempty" name:"WebFrameworkNum"`
	// Web站点数量

	WebSiteNum *uint64 `json:"WebSiteNum,omitempty" name:"WebSiteNum"`
	// Jar包数量

	JarPackageNum *uint64 `json:"JarPackageNum,omitempty" name:"JarPackageNum"`
	// 启动服务数量

	StartServiceNum *uint64 `json:"StartServiceNum,omitempty" name:"StartServiceNum"`
	// 计划任务数量

	ScheduledTaskNum *uint64 `json:"ScheduledTaskNum,omitempty" name:"ScheduledTaskNum"`
	// 环境变量数量

	EnvironmentVariableNum *uint64 `json:"EnvironmentVariableNum,omitempty" name:"EnvironmentVariableNum"`
	// 内核模块数量

	KernelModuleNum *uint64 `json:"KernelModuleNum,omitempty" name:"KernelModuleNum"`
	// 系统安装包数量

	SystemInstallationPackageNum *uint64 `json:"SystemInstallationPackageNum,omitempty" name:"SystemInstallationPackageNum"`
	// 剩余防护时长

	SurplusProtectDay *uint64 `json:"SurplusProtectDay,omitempty" name:"SurplusProtectDay"`
	// 客户端是否安装  1 已安装 0 未安装

	CWPStatus *uint64 `json:"CWPStatus,omitempty" name:"CWPStatus"`
	// 标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// 防护等级

	ProtectLevel *string `json:"ProtectLevel,omitempty" name:"ProtectLevel"`
	// 防护时长

	ProtectedDay *uint64 `json:"ProtectedDay,omitempty" name:"ProtectedDay"`
}

type TaskIdListKey struct {

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
}

type ApplyTrial struct {

	// 是否试用

	TrialStatus *bool `json:"TrialStatus,omitempty" name:"TrialStatus"`
	// 失败原因

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
}

type CreateFreeScanChanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// success该接口可查看结果

		Data *string `json:"Data,omitempty" name:"Data"`
		// Success接口请求结果

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 0接口请求结果

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFreeScanChanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFreeScanChanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationInfoRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeOrganizationInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyValue struct {

	// 字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type SubnetDetailInfo struct {

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// 可用ip数

	IpAvailable *uint64 `json:"IpAvailable,omitempty" name:"IpAvailable"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// Cidr地址

	CIDR *string `json:"CIDR,omitempty" name:"CIDR"`
	// vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc名字

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// Tag标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
}

type AddBASTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 资产列表

	AssetIDs []*string `json:"AssetIDs,omitempty" name:"AssetIDs"`
	// 资产选择类型，0为正选，1为反选

	AssetType *uint64 `json:"AssetType,omitempty" name:"AssetType"`
	// 执行的剧本id列表

	ScriptIDs []*string `json:"ScriptIDs,omitempty" name:"ScriptIDs"`
	// 再次执行传入的taskID

	TaskConfigID []*string `json:"TaskConfigID,omitempty" name:"TaskConfigID"`
	// 执行壳

	Shell *string `json:"Shell,omitempty" name:"Shell"`
	// 周期执行时间

	Periodic *string `json:"Periodic,omitempty" name:"Periodic"`
	// 超时时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 账号别名

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 执行任务的uin

	TaskUin *string `json:"TaskUin,omitempty" name:"TaskUin"`
}

func (r *AddBASTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBASTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCanFixMachineRequest struct {
	*tchttp.BaseRequest

	// 漏洞cveid

	CVEId *string `json:"CVEId,omitempty" name:"CVEId"`
	// 漏扫下发扫描任务的任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeCanFixMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCanFixMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器列表

		Data []*ContainerVO `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回报告临时下载url

		Data []*TaskLogURL `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskLogURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskLogURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CFWParams struct {

	// //互联网边界开关  1：开启 0：关闭

	InternetEdgeSwitch *int64 `json:"InternetEdgeSwitch,omitempty" name:"InternetEdgeSwitch"`
	// // 0观察，1拦截，2严格

	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`
}

type RiskCenterRiskOverview struct {

	// 端口风险总数

	PortTotal *uint64 `json:"PortTotal,omitempty" name:"PortTotal"`
	// 端口风险高危数量

	PortHighLevel *uint64 `json:"PortHighLevel,omitempty" name:"PortHighLevel"`
	// 	弱口令风险总数

	WeakPasswordTotal *uint64 `json:"WeakPasswordTotal,omitempty" name:"WeakPasswordTotal"`
	// 弱口令风险高危数量

	WeakPasswordHighLevel *uint64 `json:"WeakPasswordHighLevel,omitempty" name:"WeakPasswordHighLevel"`
	// 网站风险数量

	WebsiteTotal *uint64 `json:"WebsiteTotal,omitempty" name:"WebsiteTotal"`
	// 网站高危风险数量

	WebsiteHighLevel *uint64 `json:"WebsiteHighLevel,omitempty" name:"WebsiteHighLevel"`
	// 最新的扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 漏洞风险数

	VULTotal *int64 `json:"VULTotal,omitempty" name:"VULTotal"`
	// 高危漏洞风险数

	VULHighLevel *int64 `json:"VULHighLevel,omitempty" name:"VULHighLevel"`
	// 配置项风险数量

	CFGTotal *int64 `json:"CFGTotal,omitempty" name:"CFGTotal"`
	// 高危配置项风险数量

	CFGHighLevel *int64 `json:"CFGHighLevel,omitempty" name:"CFGHighLevel"`
	// 测绘服务风险数量

	ServerTotal *int64 `json:"ServerTotal,omitempty" name:"ServerTotal"`
	// 测绘服务高危数量

	ServerHighLevel *int64 `json:"ServerHighLevel,omitempty" name:"ServerHighLevel"`
}

type DescribeAssetSecTrendListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产安全趋势数据

		Data []*AssetSecTrendObject `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSecTrendListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSecTrendListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFreeChanceRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeFreeChanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFreeChanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CVMAssetVO struct {

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 防护状态

	CWPStatus *uint64 `json:"CWPStatus,omitempty" name:"CWPStatus"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 私网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc 名

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// appid信息

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 可用区

	AvailableArea *string `json:"AvailableArea,omitempty" name:"AvailableArea"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// uuid

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// qquid

	InstanceQUuid *string `json:"InstanceQUuid,omitempty" name:"InstanceQUuid"`
	// os名

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 分区

	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
	// cpu信息

	CPUInfo *string `json:"CPUInfo,omitempty" name:"CPUInfo"`
	// cpu大小

	CPUSize *uint64 `json:"CPUSize,omitempty" name:"CPUSize"`
	// cpu负载

	CPULoad *string `json:"CPULoad,omitempty" name:"CPULoad"`
	// 内存大小

	MemorySize *string `json:"MemorySize,omitempty" name:"MemorySize"`
	// 内存负载

	MemoryLoad *string `json:"MemoryLoad,omitempty" name:"MemoryLoad"`
	// 硬盘大小

	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`
	// 硬盘负载

	DiskLoad *string `json:"DiskLoad,omitempty" name:"DiskLoad"`
	// 账号数

	AccountCount *string `json:"AccountCount,omitempty" name:"AccountCount"`
	// 进程数

	ProcessCount *string `json:"ProcessCount,omitempty" name:"ProcessCount"`
	// 软件应用

	AppCount *string `json:"AppCount,omitempty" name:"AppCount"`
	// 监听端口

	PortCount *uint64 `json:"PortCount,omitempty" name:"PortCount"`
	// 网络攻击

	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`
	// 网络访问

	Access *uint64 `json:"Access,omitempty" name:"Access"`
	// 网络拦截

	Intercept *uint64 `json:"Intercept,omitempty" name:"Intercept"`
	// 入向峰值带宽

	InBandwidth *string `json:"InBandwidth,omitempty" name:"InBandwidth"`
	// 出向峰值带宽

	OutBandwidth *string `json:"OutBandwidth,omitempty" name:"OutBandwidth"`
	// 入向累计流量

	InFlow *string `json:"InFlow,omitempty" name:"InFlow"`
	// 出向累计流量

	OutFlow *string `json:"OutFlow,omitempty" name:"OutFlow"`
	// 最近扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 恶意主动外联

	NetWorkOut *uint64 `json:"NetWorkOut,omitempty" name:"NetWorkOut"`
	// 端口风险

	PortRisk *uint64 `json:"PortRisk,omitempty" name:"PortRisk"`
	// 漏洞风险

	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitempty" name:"VulnerabilityRisk"`
	// 配置风险

	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitempty" name:"ConfigurationRisk"`
	// 扫描任务数

	ScanTask *uint64 `json:"ScanTask,omitempty" name:"ScanTask"`
	// 标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// memberId

	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
	// os全称

	Os *string `json:"Os,omitempty" name:"Os"`
	// 风险服务暴露

	RiskExposure *int64 `json:"RiskExposure,omitempty" name:"RiskExposure"`
	// 模拟攻击工具状态。0代表未安装，1代表已安装，2代表已离线

	BASAgentStatus *int64 `json:"BASAgentStatus,omitempty" name:"BASAgentStatus"`
}

type Vpc struct {

	// 子网(只支持32位)

	Subnet *uint64 `json:"Subnet,omitempty" name:"Subnet"`
	// 互通vpc(只支持32位)

	ConnectedVpc *uint64 `json:"ConnectedVpc,omitempty" name:"ConnectedVpc"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// region区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 云服务器(只支持32位)

	CVM *uint64 `json:"CVM,omitempty" name:"CVM"`
	// 标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// dns域名

	DNS []*string `json:"DNS,omitempty" name:"DNS"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// cidr网段

	CIDR *string `json:"CIDR,omitempty" name:"CIDR"`
	// 资产创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
}

type CreateCWPTrialRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 无

	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`
	// 无

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
	// 无

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	LicenseNum *int64 `json:"LicenseNum,omitempty" name:"LicenseNum"`
	// 无

	LicenseType *int64 `json:"LicenseType,omitempty" name:"LicenseType"`
}

func (r *CreateCWPTrialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCWPTrialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWafSpartaProtectionRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 需要防御的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 证书类型，0表示没有证书，CertType=1表示自有证书,2 为托管证书

	CertType *int64 `json:"CertType,omitempty" name:"CertType"`
	// CertType=1时，需要填次参数，表示证书内容

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// CertType=1时，需要填次参数，表示证书的私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// CertType=2时，需要填次参数，表示证书的ID

	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`
	// Waf的资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 表示是否开启了CDN代理，1：有部署CDN，0：未部署CDN

	IsCDN *int64 `json:"IsCDN,omitempty" name:"IsCDN"`
	// HTTPS回源协议，填http或者https

	UpstreamScheme *string `json:"UpstreamScheme,omitempty" name:"UpstreamScheme"`
	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段

	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitempty" name:"HttpsUpstreamPort"`
	// 是否开启灰度，0表示不开启灰度

	IsGray *int64 `json:"IsGray,omitempty" name:"IsGray"`
	// 灰度的地区 ["gz","bj"]

	GrayAreas []*string `json:"GrayAreas,omitempty" name:"GrayAreas"`
	// 回源类型，0表示通过IP回源,1 表示通过域名回源

	UpstreamType *int64 `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// UpstreamType=1时，填次字段表示回源域名

	UpstreamDomain *string `json:"UpstreamDomain,omitempty" name:"UpstreamDomain"`
	// UpstreamType=0时，填次字段表示回源IP

	SrcList []*string `json:"SrcList,omitempty" name:"SrcList"`
	// 是否开启HTTP2,开启HTTP2需要HTTPS支持

	IsHttp2 *int64 `json:"IsHttp2,omitempty" name:"IsHttp2"`
	// 是否开启WebSocket支持，1表示开启，0不开启 必填

	IsWebsocket *int64 `json:"IsWebsocket,omitempty" name:"IsWebsocket"`
	// 负载均衡策略，0表示轮徇，1表示IP hash

	LoadBalance *string `json:"LoadBalance,omitempty" name:"LoadBalance"`
	// 表示是否强制跳转到HTTPS，1强制跳转Https，0不强制跳转

	HttpsRewrite *int64 `json:"HttpsRewrite,omitempty" name:"HttpsRewrite"`
	// 服务有多端口需要设置此字段

	Ports []*WafSpartaPortItem `json:"Ports,omitempty" name:"Ports"`
	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF，cdn-waf表示CDN上的Web防护能力

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 是否开启长连接，仅IP回源时可以用填次参数，域名回源时这个参数无效

	IsKeepAlive *string `json:"IsKeepAlive,omitempty" name:"IsKeepAlive"`
	// 实例id，上线之后带上此字段

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// anycast IP类型开关： 0 普通IP 1 Anycast IP

	Anycast *int64 `json:"Anycast,omitempty" name:"Anycast"`
	// 	src权重

	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`
	// 是否开启主动健康检测，1表示开启，0表示不开启

	ActiveCheck *int64 `json:"ActiveCheck,omitempty" name:"ActiveCheck"`
	// TLS版本信息

	TLSVersion *int64 `json:"TLSVersion,omitempty" name:"TLSVersion"`
	// 加密套件信息

	Ciphers []*int64 `json:"Ciphers,omitempty" name:"Ciphers"`
	// 0:不支持选择：默认模版 1:通用型模版 2:安全型模版 3:自定义

	CipherTemplate *int64 `json:"CipherTemplate,omitempty" name:"CipherTemplate"`
	// 300s

	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitempty" name:"ProxyReadTimeout"`
	// 300s

	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitempty" name:"ProxySendTimeout"`
	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；

	SniType *int64 `json:"SniType,omitempty" name:"SniType"`
	// SniType=3时，需要填此参数，表示自定义的host；

	SniHost *string `json:"SniHost,omitempty" name:"SniHost"`
	// is_cdn=3时，需要填此参数，表示自定义header

	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`
}

func (r *CreateWafSpartaProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWafSpartaProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExitPocScanMissionRequest struct {
	*tchttp.BaseRequest

	// 任务id用来退出该任务

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExitPocScanMissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExitPocScanMissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogAccessSwitchRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 业务标签: 1:云防火墙，2:WAF，3:主机安全，4:SOC

	Tags []*int64 `json:"Tags,omitempty" name:"Tags"`
	// 日志接入开关：0关；1开

	AccessSwitch *int64 `json:"AccessSwitch,omitempty" name:"AccessSwitch"`
}

func (r *ModifyLogAccessSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogAccessSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchWafPatch struct {

	// 1支持虚拟补丁/全部开启，0未全部开启 2不支持该cve

	State *int64 `json:"State,omitempty" name:"State"`
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest

	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBASTaskTechResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 战术列表

		Tactics []*BASIDNamePair `json:"Tactics,omitempty" name:"Tactics"`
		// 战技列表

		Techniques []*BASIDNamePair `json:"Techniques,omitempty" name:"Techniques"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBASTaskTechResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBASTaskTechResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserCVEStatusRequest struct {
	*tchttp.BaseRequest

	// 漏洞cve编号

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 开启开关，1开启该规则，0关闭该规则

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyUserCVEStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserCVEStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbLoadBalancers struct {

	// 负载均衡id

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 网络类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 地域

	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`
	// 负载均衡ip

	Vip []*string `json:"Vip,omitempty" name:"Vip"`
	// 所在区域

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 网络id

	NumericalVpcId *uint64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 集群id

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// IP版本，ipv4 | ipv6

	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`
	// IP地址版本为ipv6时此字段有意义， IPv6Nat64 | IPv6FullChain

	IPV6Mode *string `json:"IPV6Mode,omitempty" name:"IPV6Mode"`
	// 负载均衡实例的IPv6地址

	AddressIPV6 *string `json:"AddressIPV6,omitempty" name:"AddressIPV6"`
}

type DescribeAssetConfigListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 资产值

	TargetAsset *string `json:"TargetAsset,omitempty" name:"TargetAsset"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeAssetConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsStructureRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeAssetsStructureRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsStructureRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterPortViewPortRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterPortViewPortRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterPortViewPortRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSolidProtectionDataRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSolidProtectionDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSolidProtectionDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUpdateBugInfoRequest struct {
	*tchttp.BaseRequest

	// 漏洞CVE编号

	CVEId *string `json:"CVEId,omitempty" name:"CVEId"`
	// 扫描任务下发之后的taskid

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 1和2，1在漏扫下发任务之后更新，2在主机修复结束之后更新

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyUpdateBugInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpdateBugInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HighLevelRiskData struct {

	// 风险名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 风险类型
	// port: 端口风险
	// loophole: 漏洞风险
	// weakPassword: 弱口令风险
	// webContent: 内容风险
	// cfg: 配置风险

	Type *string `json:"Type,omitempty" name:"Type"`
	// 影响资产数量

	AffectAssetCount *uint64 `json:"AffectAssetCount,omitempty" name:"AffectAssetCount"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 风险ID

	RiskId *string `json:"RiskId,omitempty" name:"RiskId"`
	// 额外信息
	// Type为webContent时，这里Extra为以下几个枚举：sensitive-敏感信息/horse-挂马/darklink-篡改

	Extra *string `json:"Extra,omitempty" name:"Extra"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
}

type CreateCWPTrialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Resource *Resource `json:"Resource,omitempty" name:"Resource"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCWPTrialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCWPTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 筛选列表

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeContainerAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCoreAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 核心资产详细信息

		Data []*AssetObject `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCoreAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCoreAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExitExamResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 暂停体检相关信息

		Data *CreateExamMissionResp `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyExitExamResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExitExamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAssetVO struct {

	// 资产id

	AssetId []*string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名

	AssetName []*string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产类型

	AssetType []*string `json:"AssetType,omitempty" name:"AssetType"`
	// 地域

	Region []*string `json:"Region,omitempty" name:"Region"`
	// Waf状态

	WAFStatus *uint64 `json:"WAFStatus,omitempty" name:"WAFStatus"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// Appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 账号id

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 账号名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
	// 是否云上资产

	IsCloud *uint64 `json:"IsCloud,omitempty" name:"IsCloud"`
	// 网络攻击

	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`
	// 网络访问

	Access *uint64 `json:"Access,omitempty" name:"Access"`
	// 网络拦截

	Intercept *uint64 `json:"Intercept,omitempty" name:"Intercept"`
	// 入站峰值带宽

	InBandwidth *string `json:"InBandwidth,omitempty" name:"InBandwidth"`
	// 出站峰值带宽

	OutBandwidth *string `json:"OutBandwidth,omitempty" name:"OutBandwidth"`
	// 入站累计流量

	InFlow *string `json:"InFlow,omitempty" name:"InFlow"`
	// 出站累计流量

	OutFlow *string `json:"OutFlow,omitempty" name:"OutFlow"`
	// 最近扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 端口风险

	PortRisk *uint64 `json:"PortRisk,omitempty" name:"PortRisk"`
	// 漏洞风险

	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitempty" name:"VulnerabilityRisk"`
	// 配置风险

	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitempty" name:"ConfigurationRisk"`
	// 扫描任务

	ScanTask *uint64 `json:"ScanTask,omitempty" name:"ScanTask"`
	// 域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 解析ip

	SeverIp []*string `json:"SeverIp,omitempty" name:"SeverIp"`
	// boi访问数量

	BotCount *uint64 `json:"BotCount,omitempty" name:"BotCount"`
	// 弱口令风险

	WeakPassword *uint64 `json:"WeakPassword,omitempty" name:"WeakPassword"`
	// 内容风险

	WebContentRisk *uint64 `json:"WebContentRisk,omitempty" name:"WebContentRisk"`
	// tag标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// 关联实例类型

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// memberiD

	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
	// cc攻击

	CCAttack *int64 `json:"CCAttack,omitempty" name:"CCAttack"`
	// web攻击

	WebAttack *int64 `json:"WebAttack,omitempty" name:"WebAttack"`
}

type ProductStatusData struct {

	// 产品当前状态 0可试用，1未防护，2防护中（已购买），3试用中，4已过期

	ProductStatus *int64 `json:"ProductStatus,omitempty" name:"ProductStatus"`
	// 返回消息，一般为不可试用原因

	TrialMessage *string `json:"TrialMessage,omitempty" name:"TrialMessage"`
	// 产品名称 CFW：云防火墙 WAF：Web应用防火墙 CWP：主机安全 SOC：安全运营中心 TCSS：容器安全

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 是否提供试用功能 0不提供 1提供

	TrialStatus *int64 `json:"TrialStatus,omitempty" name:"TrialStatus"`
	// 产品对应分层 1: 第一道防线 2: 第二道防线 3: 第三道防线 4: 加固防线

	ProductLayer *int64 `json:"ProductLayer,omitempty" name:"ProductLayer"`
	// 产品是否已上线 0: 未上线 1: 已上线

	ProductOnline *int64 `json:"ProductOnline,omitempty" name:"ProductOnline"`
	// 产品最高版本

	MostHighVersion *string `json:"MostHighVersion,omitempty" name:"MostHighVersion"`
}

type TodoListRecordAllInfo struct {

	// 用户id

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 详细信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type DescribeLicenseListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 数据量多少

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 资源id

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 标签

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
	// 过滤条件

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLicenseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterCFGViewCFGRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产视角的配置风险列表

		Data []*CFGViewCFGRisk `json:"Data,omitempty" name:"Data"`
		// 状态列表

		StatusLists []*FilterDataObject `json:"StatusLists,omitempty" name:"StatusLists"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 配置名列表

		CFGNameLists []*FilterDataObject `json:"CFGNameLists,omitempty" name:"CFGNameLists"`
		// 检查类型列表

		CheckTypeLists []*FilterDataObject `json:"CheckTypeLists,omitempty" name:"CheckTypeLists"`
		// 资产类型列表

		InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`
		// 来源列表

		FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterCFGViewCFGRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterCFGViewCFGRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeClusterAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterPodAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeClusterPodAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterPodAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayAssetInfoRequest struct {
	*tchttp.BaseRequest

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeGatewayAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterRiskTrendListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterRiskTrendListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterRiskTrendListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafProductStatusRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
}

func (r *DescribeWafProductStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafProductStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogStorageDayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogStorageDayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogStorageDayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryImageAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// filter过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRepositoryImageAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryImageAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanTimeRequest struct {
	*tchttp.BaseRequest

	// 扫描的资产个数

	AssetNumber *int64 `json:"AssetNumber,omitempty" name:"AssetNumber"`
}

func (r *DescribeScanTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExitPocScanMissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// success

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 0

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 退出成功

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExitPocScanMissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExitPocScanMissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetDetailTrendInfo struct {

	// 时间坐标

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 告警数量

	Alarm *int64 `json:"Alarm,omitempty" name:"Alarm"`
	// 入站流量

	InFlow *float64 `json:"InFlow,omitempty" name:"InFlow"`
	// 出站流量

	OutFlow *float64 `json:"OutFlow,omitempty" name:"OutFlow"`
	// 网络拦截

	Intercept *int64 `json:"Intercept,omitempty" name:"Intercept"`
	// 网络访问

	Access *int64 `json:"Access,omitempty" name:"Access"`
}

type BASAlertTypeCnt struct {

	// 主机告警类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 告警个数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 告警优先级

	Index *uint64 `json:"Index,omitempty" name:"Index"`
}

type DescribeNetworkStructureResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 层级结构

		Data *NetWorkInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkStructureResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkStructureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetViewVULRisk struct {

	// 影响资产

	AffectAsset *string `json:"AffectAsset,omitempty" name:"AffectAsset"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 资产类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 服务

	Service *string `json:"Service,omitempty" name:"Service"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 状态，0未处理、1已处置、2已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 资产唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 前端索引

	Index *string `json:"Index,omitempty" name:"Index"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 漏洞类型

	VULType *string `json:"VULType,omitempty" name:"VULType"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 描述

	Describe *string `json:"Describe,omitempty" name:"Describe"`
	// 版本名

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 相关信息

	References *string `json:"References,omitempty" name:"References"`
	// 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 漏洞url

	VULURL *string `json:"VULURL,omitempty" name:"VULURL"`
	// 漏洞名称

	VULName *string `json:"VULName,omitempty" name:"VULName"`
	// cve

	CVE *string `json:"CVE,omitempty" name:"CVE"`
	// 修复建议

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// pocid

	POCId *string `json:"POCId,omitempty" name:"POCId"`
	// 来源

	From *string `json:"From,omitempty" name:"From"`
	// 主机版本

	CWPVersion *int64 `json:"CWPVersion,omitempty" name:"CWPVersion"`
	// 是否支持修复

	IsSupportRepair *bool `json:"IsSupportRepair,omitempty" name:"IsSupportRepair"`
	// 是否支持扫描

	IsSupportDetect *bool `json:"IsSupportDetect,omitempty" name:"IsSupportDetect"`
	// 实例uuid

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 负载

	Payload *string `json:"Payload,omitempty" name:"Payload"`
	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞

	EMGCVulType *int64 `json:"EMGCVulType,omitempty" name:"EMGCVulType"`
}

type FreeLimitResp struct {

	// 剩余2次可扫描

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type LogAccessTypeDescCount struct {

	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 日志类型说明

	TypeDesc *string `json:"TypeDesc,omitempty" name:"TypeDesc"`
	// 单类型日志条数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type RiskOverviewObject struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
	// 高风险的个数

	HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
}

type DescribeExaminationPayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 无

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 无

		Data *AllProductsInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExaminationPayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExaminationPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicIpAssetInfoRequest struct {
	*tchttp.BaseRequest

	// 公网Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
}

func (r *DescribePublicIpAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicIpAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSSLCertificatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本信息

		Data []*SSLCertificate `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSSLCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSSLCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebServiceAssetsListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤参数

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeWebServiceAssetsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebServiceAssetsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogStorageDayRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 业务标签: 1:云防火墙，2:WAF，3:主机安全，4:SOC

	Tags []*int64 `json:"Tags,omitempty" name:"Tags"`
	// 日志存储时长，只允许特定值，值参考返回的时长列表

	StorageDay *int64 `json:"StorageDay,omitempty" name:"StorageDay"`
}

func (r *ModifyLogStorageDayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogStorageDayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskClassData struct {

	// 端口风险

	PortRiskCount *uint64 `json:"PortRiskCount,omitempty" name:"PortRiskCount"`
	// 漏洞风险

	VULRiskCount *uint64 `json:"VULRiskCount,omitempty" name:"VULRiskCount"`
	// 弱口令风险

	WeakPasswordRiskCount *uint64 `json:"WeakPasswordRiskCount,omitempty" name:"WeakPasswordRiskCount"`
	// 网页内容

	WebsiteCount *uint64 `json:"WebsiteCount,omitempty" name:"WebsiteCount"`
	// 总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
	// 配置风险

	CFGCount *uint64 `json:"CFGCount,omitempty" name:"CFGCount"`
	// 服务测绘风险

	ServerCount *uint64 `json:"ServerCount,omitempty" name:"ServerCount"`
}

type DescribeRepositoryImageAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仓库镜像列表

		Data []*RepositoryImageVO `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// region列表

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryImageAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryImageAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 数据库limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 数据库offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 机器地域

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 机器类型

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 过滤条件

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNICAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*NICAsset `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域列表

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 资产类型列表

		AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`
		// vpc列表

		VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`
		// appid列表

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNICAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNICAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TimeData struct {

	// 开通时长

	ProtectTime *int64 `json:"ProtectTime,omitempty" name:"ProtectTime"`
	// 到期时间

	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`
}

type WafClbHostRecordInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 主域名

	MainDomain *string `json:"MainDomain,omitempty" name:"MainDomain"`
	// waf模式，同saas waf保持一致

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// waf和LD的绑定，0：没有绑定，1：绑定

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 域名状态，0：正常，1：未检测到流量，2：即将过期，3：过期

	State *uint64 `json:"State,omitempty" name:"State"`
	// 使用的规则，同saas waf保持一致

	Engine *uint64 `json:"Engine,omitempty" name:"Engine"`
	// 是否开启代理，0：不开启，1：开启

	IsCDN *uint64 `json:"IsCDN,omitempty" name:"IsCDN"`
	// 绑定的LB列表

	LoadBalancerSet []*WafLoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
	// 域名绑定的LB的地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 产品分类，取值为：sparta-waf、clb-waf、cdn-waf

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// WAF的流量模式，1：清洗模式，0：镜像模式

	FlowMode *uint64 `json:"FlowMode,omitempty" name:"FlowMode"`
	// 是否开启访问日志，1：开启，0：关闭

	CLSStatus *uint64 `json:"CLSStatus,omitempty" name:"CLSStatus"`
	// 防护等级，可选值100,200,300

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 域名需要下发到的cdc集群列表

	CdcClusters []*string `json:"CdcClusters,omitempty" name:"CdcClusters"`
	// 应用型负载均衡类型: clb或者apisix，默认clb

	AlbType *string `json:"AlbType,omitempty" name:"AlbType"`
	// IsCdn=3时，需要填此参数，表示自定义header

	IpHeaders []*string `json:"IpHeaders,omitempty" name:"IpHeaders"`
	// 规则引擎类型， 1: menshen, 2:tiga

	EngineType *int64 `json:"EngineType,omitempty" name:"EngineType"`
}

type DescribeAuditDiversityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计elment信息

		ObjectTopElement []*NodeInfo `json:"ObjectTopElement,omitempty" name:"ObjectTopElement"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditDiversityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuditDiversityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤参数

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeGatewayAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 操作日志列表

		Data []*DataOperateLog `json:"Data,omitempty" name:"Data"`
		// 操作行为列表

		ActionLists []*FilterDataObject `json:"ActionLists,omitempty" name:"ActionLists"`
		// 操作类型列表

		TypeLists []*FilterDataObject `json:"TypeLists,omitempty" name:"TypeLists"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopX struct {

	// 类型

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 元素

	Element []*Element `json:"Element,omitempty" name:"Element"`
}

type AlarmDiversityRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *AlarmDiversityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmDiversityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetLocationListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeAssetLocationListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetLocationListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebServiceAssetsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*WebServiceAssetInfo `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// appid列表

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 地域列表

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 服务类型

		NetworkTypeList []*FilterDataObject `json:"NetworkTypeList,omitempty" name:"NetworkTypeList"`
		// 防护类型

		ProtectTypeList []*FilterDataObject `json:"ProtectTypeList,omitempty" name:"ProtectTypeList"`
		// 云防防护类型

		WebProtectTypeList []*FilterDataObject `json:"WebProtectTypeList,omitempty" name:"WebProtectTypeList"`
		// 协议列表

		ProtocolList []*FilterDataObject `json:"ProtocolList,omitempty" name:"ProtocolList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebServiceAssetsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebServiceAssetsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetSecTrendObject struct {

	// 资产安全趋势图横坐标,表示最近30天的日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 总资产数量

	AssetTotal *uint64 `json:"AssetTotal,omitempty" name:"AssetTotal"`
	// 暴露面总数，漏洞风险+端口风险

	ExposeTotal *uint64 `json:"ExposeTotal,omitempty" name:"ExposeTotal"`
	// 所有风险类型总和

	RiskTotal *uint64 `json:"RiskTotal,omitempty" name:"RiskTotal"`
}

type WafTLSCipherInfo struct {

	// 版本id

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 套件id

	CipherId *int64 `json:"CipherId,omitempty" name:"CipherId"`
	// 套件名称

	CipherName *string `json:"CipherName,omitempty" name:"CipherName"`
}

type NICAsset struct {

	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 资产ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 私有ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// 出向峰值带宽

	OutboundPeakBandwidth *string `json:"OutboundPeakBandwidth,omitempty" name:"OutboundPeakBandwidth"`
	// 入向峰值带宽

	InboundPeakBandwidth *string `json:"InboundPeakBandwidth,omitempty" name:"InboundPeakBandwidth"`
	// 出站累计流量

	OutboundCumulativeFlow *string `json:"OutboundCumulativeFlow,omitempty" name:"OutboundCumulativeFlow"`
	// 入站累计流量

	InboundCumulativeFlow *string `json:"InboundCumulativeFlow,omitempty" name:"InboundCumulativeFlow"`
	// 网络攻击

	NetworkAttack *int64 `json:"NetworkAttack,omitempty" name:"NetworkAttack"`
	// 暴露端口

	ExposedPort *int64 `json:"ExposedPort,omitempty" name:"ExposedPort"`
	// 暴露漏洞

	ExposedVUL *int64 `json:"ExposedVUL,omitempty" name:"ExposedVUL"`
	// 配置风险

	ConfigureRisk *int64 `json:"ConfigureRisk,omitempty" name:"ConfigureRisk"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务数

	ScanTask *int64 `json:"ScanTask,omitempty" name:"ScanTask"`
	// 最后扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
}

type CreateWafClbHostRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 防护域名配置信息

	Host *WafClbHostRecordInfo `json:"Host,omitempty" name:"Host"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *CreateWafClbHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWafClbHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskClassInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险分类信息

		Data *RiskClassData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskClassInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskClassInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafTlsVersionRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
}

func (r *DescribeWafTlsVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafTlsVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Element struct {

	// 统计类型

	Key *string `json:"Key,omitempty" name:"Key"`
	// 统计对象

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GroupLastNumThreeNodeInfo struct {

	// 节点信息

	NodeInfo *GroupNodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`
	// 子节点信息

	SubNodeList []*GroupLastNumTwoNodeInfo `json:"SubNodeList,omitempty" name:"SubNodeList"`
}

type ToDoListRecordResp struct {

	// 详细信息

	TodoListInfo []*TodoListRecordAllInfo `json:"TodoListInfo,omitempty" name:"TodoListInfo"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DescribeRiskCenterServerRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterServerRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterServerRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainAndIpRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// -

	Content []*string `json:"Content,omitempty" name:"Content"`
}

func (r *CreateDomainAndIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDomainAndIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFGRiskAdvanceCFGListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 检测项内容

		Data []*CheckItemIdInfo `json:"Data,omitempty" name:"Data"`
		// 风险枚举

		RiskLevelLists []*FilterDataObject `json:"RiskLevelLists,omitempty" name:"RiskLevelLists"`
		// 资产类型枚举

		ResourceTypeLists []*FilterDataObject `json:"ResourceTypeLists,omitempty" name:"ResourceTypeLists"`
		// 检测项类型枚举

		TypeLists []*FilterDataObject `json:"TypeLists,omitempty" name:"TypeLists"`
		// 标准枚举

		StandardLists []*FilterDataObject `json:"StandardLists,omitempty" name:"StandardLists"`
		// 检测来源枚举

		CheckFromLists []*FilterDataObject `json:"CheckFromLists,omitempty" name:"CheckFromLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCFGRiskAdvanceCFGListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFGRiskAdvanceCFGListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogAuditSetting struct {

	// 日志总存储条数

	StorageCount *int64 `json:"StorageCount,omitempty" name:"StorageCount"`
	// 多个对应日志类型的存储信息

	LogTypeStorages []*LogTypeStorage `json:"LogTypeStorages,omitempty" name:"LogTypeStorages"`
	// 用户时间参数可选列表

	StorageDayList []*int64 `json:"StorageDayList,omitempty" name:"StorageDayList"`
	// 用户当前存储时间

	StorageDay *int64 `json:"StorageDay,omitempty" name:"StorageDay"`
	// 用户剩余清空日志次数

	CleanUpCount *int64 `json:"CleanUpCount,omitempty" name:"CleanUpCount"`
	// 上次设置时间

	LastStorageDate *string `json:"LastStorageDate,omitempty" name:"LastStorageDate"`
	// 上次清空时间

	LastCleanUpDate *string `json:"LastCleanUpDate,omitempty" name:"LastCleanUpDate"`
	// 上次设置存储类型时间

	LastLogTypeDate *string `json:"LastLogTypeDate,omitempty" name:"LastLogTypeDate"`
	// 业务标签

	BizTag *int64 `json:"BizTag,omitempty" name:"BizTag"`
	// 业务名称

	BizName *string `json:"BizName,omitempty" name:"BizName"`
	// 产品当前状态 0可试用，1未防护，2防护中（已购买），3试用中，4已过期

	ProductStatus *int64 `json:"ProductStatus,omitempty" name:"ProductStatus"`
	// 产品最高版本

	MostHighVersion *string `json:"MostHighVersion,omitempty" name:"MostHighVersion"`
	// 日志接入开关

	LogAccessSwitch *bool `json:"LogAccessSwitch,omitempty" name:"LogAccessSwitch"`
}

type AddBASTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回

		TaskID []*string `json:"TaskID,omitempty" name:"TaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddBASTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBASTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllAssetListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeAllAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFWRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DescribeCFWRuleOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWRuleOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskCenterWeakPwdRiskInputParam struct {

	// 检测项ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 是否开启，0-不开启，1-开启

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type DescribeIndeterminacyAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Data *IndeterminacyAssetInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndeterminacyAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndeterminacyAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationUserInfoRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// home为列表主页请求

	From *string `json:"From,omitempty" name:"From"`
}

func (r *DescribeOrganizationUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePDFDownLoadLogRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribePDFDownLoadLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePDFDownLoadLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcConnectedVpcListsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤参数

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeVpcConnectedVpcListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcConnectedVpcListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagGroupAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产类型总数

		Data []*AllAsset `json:"Data,omitempty" name:"Data"`
		// 资产总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 地域

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 资产类型

		AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagGroupAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagGroupAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PortRiskAdvanceCFGList struct {

	// 端口集合,以逗号分隔

	PortSets *string `json:"PortSets,omitempty" name:"PortSets"`
	// 检测项类型，0-系统定义，1-用户自定义

	CheckType *int64 `json:"CheckType,omitempty" name:"CheckType"`
	// 检测项描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 识别来源

	CheckFrom *string `json:"CheckFrom,omitempty" name:"CheckFrom"`
	// 是否启用，1-启用，0-禁用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type ScanTaskInfoList struct {

	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// corn

	ScanPlanContent *string `json:"ScanPlanContent,omitempty" name:"ScanPlanContent"`
	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义

	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
	// 创建时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 排除扫描资产信息

	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitempty" name:"SelfDefiningAssets"`
	// 预估时间

	PredictTime *int64 `json:"PredictTime,omitempty" name:"PredictTime"`
	// 预估完成时间

	PredictEndTime *string `json:"PredictEndTime,omitempty" name:"PredictEndTime"`
	// 报告数量

	ReportNumber *int64 `json:"ReportNumber,omitempty" name:"ReportNumber"`
	// 资产数量

	AssetNumber *int64 `json:"AssetNumber,omitempty" name:"AssetNumber"`
	// 扫描状态 0 初始值  1正在扫描  2扫描完成  3扫描出错

	ScanStatus *int64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 任务进度

	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
	// port/poc/weakpass/webcontent/configrisk

	ScanItem *string `json:"ScanItem,omitempty" name:"ScanItem"`
	// 0-全扫，1-指定资产扫，2-排除资产扫

	ScanAssetType *int64 `json:"ScanAssetType,omitempty" name:"ScanAssetType"`
	// vss子任务id

	VSSTaskId *string `json:"VSSTaskId,omitempty" name:"VSSTaskId"`
	// cspm子任务id

	CSPMTaskId *string `json:"CSPMTaskId,omitempty" name:"CSPMTaskId"`
	// 主机漏扫子任务id

	CWPPOCId *string `json:"CWPPOCId,omitempty" name:"CWPPOCId"`
	// 主机基线子任务id

	CWPBlId *string `json:"CWPBlId,omitempty" name:"CWPBlId"`
	// vss子任务进度

	VSSTaskProcess *int64 `json:"VSSTaskProcess,omitempty" name:"VSSTaskProcess"`
	// cspm子任务进度

	CSPMTaskProcess *uint64 `json:"CSPMTaskProcess,omitempty" name:"CSPMTaskProcess"`
	// 主机漏扫子任务进度

	CWPPOCProcess *int64 `json:"CWPPOCProcess,omitempty" name:"CWPPOCProcess"`
	// 主机基线子任务进度

	CWPBlProcess *uint64 `json:"CWPBlProcess,omitempty" name:"CWPBlProcess"`
	// 异常状态码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 异常信息

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// 周期任务开始的天数

	StartDay *int64 `json:"StartDay,omitempty" name:"StartDay"`
	// 扫描频率,单位天,1-每天,7-每周,30-月,0-扫描一次

	Frequency *int64 `json:"Frequency,omitempty" name:"Frequency"`
	// 完成次数

	CompleteNumber *int64 `json:"CompleteNumber,omitempty" name:"CompleteNumber"`
	// 已完成资产个数

	CompleteAssetNumber *int64 `json:"CompleteAssetNumber,omitempty" name:"CompleteAssetNumber"`
	// 风险数

	RiskCount *int64 `json:"RiskCount,omitempty" name:"RiskCount"`
	// 资产

	Assets []*TaskAssetObject `json:"Assets,omitempty" name:"Assets"`
	// 用户Appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户主账户ID

	UIN *string `json:"UIN,omitempty" name:"UIN"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 体检模式，0-标准模式，1-快速模式，2-高级模式

	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
	// 扫描来源

	ScanFrom *string `json:"ScanFrom,omitempty" name:"ScanFrom"`
	// 是否可以编辑，1-可以，0-不可以，对应多账户管理使用

	IsEdit *int64 `json:"IsEdit,omitempty" name:"IsEdit"`
	// 是否限免体检0不是，1是

	IsFree *int64 `json:"IsFree,omitempty" name:"IsFree"`
}

type DescribeDbAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// db资产详情

		Data *DbAssetInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDbAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSocLogStorageSpaceRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
}

func (r *DescribeSocLogStorageSpaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSocLogStorageSpaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObjectTopElement struct {

	// 节点信息

	NodeInfo []*NodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`
}

type TagCount struct {

	// 产品名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志条数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type DescribeRiskCenterCFGViewCFGRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterCFGViewCFGRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterCFGViewCFGRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagCountsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 查询日志分类, 前一个数字代表产品，后一个数字代表日志类型

	LogType []*string `json:"LogType,omitempty" name:"LogType"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 逻辑表达式

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 集群 0:国内 1:国际 2:全部

	Cluster *int64 `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DescribeTagCountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagCountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VULRiskAdvanceCFGList struct {

	// 风险ID

	RiskId *string `json:"RiskId,omitempty" name:"RiskId"`
	// 漏洞名称

	VULName *string `json:"VULName,omitempty" name:"VULName"`
	// 风险等级

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 识别来源

	CheckFrom *string `json:"CheckFrom,omitempty" name:"CheckFrom"`
	// 是否启用，1-启用，0-禁用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 风险类型

	VULType *string `json:"VULType,omitempty" name:"VULType"`
	// 影响版本

	ImpactVersion *string `json:"ImpactVersion,omitempty" name:"ImpactVersion"`
	// CVE

	CVE *string `json:"CVE,omitempty" name:"CVE"`
	// 漏洞标签

	VULTag []*string `json:"VULTag,omitempty" name:"VULTag"`
	// 修复方式

	FixMethod []*string `json:"FixMethod,omitempty" name:"FixMethod"`
	// 披露时间

	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞

	EMGCVulType *int64 `json:"EMGCVulType,omitempty" name:"EMGCVulType"`
}

type DescribeFreeLimitRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeFreeLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFreeLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafParams struct {

	// 1:saas waf 2:clb waf

	WafType *int64 `json:"WafType,omitempty" name:"WafType"`
	// 可选的 广州（控制台的地域）：1（RealRegion参数）、上海：4、成都：16、北京：8

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeAssetStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机资产总数

		CVMAssetTotal *uint64 `json:"CVMAssetTotal,omitempty" name:"CVMAssetTotal"`
		// 未防护的主机资产总数

		UnprotectedCVMTotal *uint64 `json:"UnprotectedCVMTotal,omitempty" name:"UnprotectedCVMTotal"`
		// 主机风险资产总数

		RiskCVMTotal *uint64 `json:"RiskCVMTotal,omitempty" name:"RiskCVMTotal"`
		// 未防护的核心主机总数

		NonCoreCVMTotal *uint64 `json:"NonCoreCVMTotal,omitempty" name:"NonCoreCVMTotal"`
		// 公网ip资产总数

		PublicAssetTotal *uint64 `json:"PublicAssetTotal,omitempty" name:"PublicAssetTotal"`
		// 未防护的公网资产总数

		UnprotectedPublicTotal *uint64 `json:"UnprotectedPublicTotal,omitempty" name:"UnprotectedPublicTotal"`
		// 未防护的核心公网资产总数

		NonCorePublicTotal *uint64 `json:"NonCorePublicTotal,omitempty" name:"NonCorePublicTotal"`
		// Ip公网风险资产总数

		RiskPublicTotal *uint64 `json:"RiskPublicTotal,omitempty" name:"RiskPublicTotal"`
		// 域名总数

		DomainTotal *uint64 `json:"DomainTotal,omitempty" name:"DomainTotal"`
		// 未防护的域名总数

		UnprotectedDomainTotal *uint64 `json:"UnprotectedDomainTotal,omitempty" name:"UnprotectedDomainTotal"`
		// 域名风险总数

		RiskDomainTotal *uint64 `json:"RiskDomainTotal,omitempty" name:"RiskDomainTotal"`
		// 未防护的核心域名资产数

		NonCoreDomainTotal *uint64 `json:"NonCoreDomainTotal,omitempty" name:"NonCoreDomainTotal"`
		// web总数

		WebTotal *uint64 `json:"WebTotal,omitempty" name:"WebTotal"`
		// 网关资产总数

		GatewayTotal *uint64 `json:"GatewayTotal,omitempty" name:"GatewayTotal"`
		// myysql资产总数

		DbTotal *uint64 `json:"DbTotal,omitempty" name:"DbTotal"`
		// 资产总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 容器总数

		ContainerTotal *int64 `json:"ContainerTotal,omitempty" name:"ContainerTotal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCVMAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// -

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeCVMAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCVMAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterAssetViewVULRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产视角的漏洞风险列表

		Data []*AssetViewVULRisk `json:"Data,omitempty" name:"Data"`
		// 状态列表

		StatusLists []*FilterDataObject `json:"StatusLists,omitempty" name:"StatusLists"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 来源列表

		FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`
		// 漏洞类型列表

		VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitempty" name:"VULTypeLists"`
		// 资产类型列表

		InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewVULRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterAssetViewVULRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预计扫描120分钟

		Data *int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 需要修改核心标记资产id、ip、domain集合

	AssetId []*string `json:"AssetId,omitempty" name:"AssetId"`
	// 是否核心

	IsNotCore *uint64 `json:"IsNotCore,omitempty" name:"IsNotCore"`
	// 资产类型，1 ip, 2 domain， 3 主机

	AssetType *uint64 `json:"AssetType,omitempty" name:"AssetType"`
}

func (r *ModifyAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BASAlertStatistic struct {

	// 剧本名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 剧本关联的告警数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 剧本关联告警分类

	CWPAlertLog []*BASAlertTypeCnt `json:"CWPAlertLog,omitempty" name:"CWPAlertLog"`
}

type DescribeUserAccountInfoRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeUserAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAccountInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWebServiceRequest struct {
	*tchttp.BaseRequest

	// xx

	WebIds []*DeleteId `json:"WebIds,omitempty" name:"WebIds"`
}

func (r *DeleteWebServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWebServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetLocationListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产分布数据

		Data []*AssetLocation `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetLocationListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetLocationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCVMAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// -

		Data []*CVMAssetVO `json:"Data,omitempty" name:"Data"`
		// 地域列表

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 防护状态

		DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitempty" name:"DefenseStatusList"`
		// vpc枚举

		VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`
		// 资产类型枚举

		AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`
		// 操作系统枚举

		SystemTypeList []*FilterDataObject `json:"SystemTypeList,omitempty" name:"SystemTypeList"`
		// ip列表

		IpTypeList []*FilterDataObject `json:"IpTypeList,omitempty" name:"IpTypeList"`
		// appid列表

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 可用区列表

		ZoneList []*FilterDataObject `json:"ZoneList,omitempty" name:"ZoneList"`
		// os列表

		OsList []*FilterDataObject `json:"OsList,omitempty" name:"OsList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCVMAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCVMAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBASSummaryDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已安装模拟工具资产数

		AgentTotalCount *uint64 `json:"AgentTotalCount,omitempty" name:"AgentTotalCount"`
		// 未安装模拟工具资产数

		UnRegisterAgentTotalCount *uint64 `json:"UnRegisterAgentTotalCount,omitempty" name:"UnRegisterAgentTotalCount"`
		// 用户上次执行模拟攻击的时间

		LastExecTime *string `json:"LastExecTime,omitempty" name:"LastExecTime"`
		// 模拟攻击执行成功次数

		Success *uint64 `json:"Success,omitempty" name:"Success"`
		// 模拟攻击执行异常次数

		Fail *uint64 `json:"Fail,omitempty" name:"Fail"`
		// 模拟攻击执行中止次数

		Stopped *uint64 `json:"Stopped,omitempty" name:"Stopped"`
		// 模拟攻击关联总告警数

		AlertLogTotalCount *uint64 `json:"AlertLogTotalCount,omitempty" name:"AlertLogTotalCount"`
		// 告警关联剧本情况

		AlertLog []*BASAlertStatistic `json:"AlertLog,omitempty" name:"AlertLog"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBASSummaryDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBASSummaryDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NICDetailInfo struct {

	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc name

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// 最近扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// tag标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
}

type DescribeContainerAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Data *ContainerAssetInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSolidProtectionDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 立体防护中心概览数据

		Data *SolidProtectionDataResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSolidProtectionDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSolidProtectionDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafProductStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0为成功 其他值为失败

		Succ *int64 `json:"Succ,omitempty" name:"Succ"`
		// 产品状态 0 可试用 1 不可试用 2防护中 3 试用中 4 已到期

		ProductStatus *int64 `json:"ProductStatus,omitempty" name:"ProductStatus"`
		// 版本

		MostHighVersion *string `json:"MostHighVersion,omitempty" name:"MostHighVersion"`
		// 产品名称

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// 是否可试用 是否可试用  1可试用 0不可试用

		TrialStatus *int64 `json:"TrialStatus,omitempty" name:"TrialStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafProductStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafProductStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetLocalImageComponent struct {

	// 用户id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像组件id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 镜像名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// 组件路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 组件类型：系统组件，应用组件

	Type *string `json:"Type,omitempty" name:"Type"`
	// 漏洞风险

	VulRisk *int64 `json:"VulRisk,omitempty" name:"VulRisk"`
}

type DescribeGatewayAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*GateWayAsset `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域列表

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 资产类型列表

		AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`
		// vpc列表

		VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`
		// appid列表

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogAccessRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
}

func (r *DescribeLogAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRiskEvidenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回到处文件下载的临时url和文件格式

		Data *WebContentEvidenceInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRiskEvidenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRiskEvidenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSCCCSVResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 如果参数Async为1:这里返回任务ID，后面通过任务id去获取任务信息，否则：返回到处文件下载的临时url

		Data *ExportURL `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportSCCCSVResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSCCCSVResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebServiceRequest struct {
	*tchttp.BaseRequest

	// 1111

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 1.1.1.1

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 80

	Port *string `json:"Port,omitempty" name:"Port"`
	// www.xxx.com

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// web服务的AppId

	WebAppId *string `json:"WebAppId,omitempty" name:"WebAppId"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyWebServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncUserAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1：成功 0：失败

		Success *int64 `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncUserAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncUserAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamChangeToDoList struct {

	// 类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 具体参数

	Param []*ParamInfo `json:"Param,omitempty" name:"Param"`
}

type DescribeNeedToSolveListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 返回值

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 待办详细

		Data []*NeedToSolveList `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNeedToSolveListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNeedToSolveListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterAssetViewWeakPasswordRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 风险列表

		Data []*AssetViewWeakPassRisk `json:"Data,omitempty" name:"Data"`
		// 状态列表

		StatusLists []*FilterDataObject `json:"StatusLists,omitempty" name:"StatusLists"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 来源列表

		FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`
		// 资产类型列表

		InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`
		// 弱口令类型列表

		PasswordTypeLists []*FilterDataObject `json:"PasswordTypeLists,omitempty" name:"PasswordTypeLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterServerRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 风险服务列表

		Data []*ServerRisk `json:"Data,omitempty" name:"Data"`
		// 资产类型枚举

		InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterServerRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterServerRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafAllInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*DescribeWafInstances `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// waf域名接入状态 0为未接入  1为已接入

		State *int64 `json:"State,omitempty" name:"State"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafAllInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafAllInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndeterminacyAssetTrendRequest struct {
	*tchttp.BaseRequest

	// -

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// -

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// -

	PublicIp []*string `json:"PublicIp,omitempty" name:"PublicIp"`
	// -

	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeIndeterminacyAssetTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndeterminacyAssetTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCheckUpResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功，0失败 1成功

		Success *uint64 `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCheckUpResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCheckUpResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBASAssetAgentRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 资产名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资产id

	AssetID *string `json:"AssetID,omitempty" name:"AssetID"`
	// 资产类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 公网ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 私网ip

	PrivateIP *string `json:"PrivateIP,omitempty" name:"PrivateIP"`
	// 操作系统

	OsSystem *string `json:"OsSystem,omitempty" name:"OsSystem"`
}

func (r *RegisterBASAssetAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterBASAssetAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductVersionInfos struct {

	// 无

	Product *uint64 `json:"Product,omitempty" name:"Product"`
	// 无

	Version *uint64 `json:"Version,omitempty" name:"Version"`
}

type ReportItemKey struct {

	// 租户ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 日志Id列表

	TaskLogList []*string `json:"TaskLogList,omitempty" name:"TaskLogList"`
}

type DescribeRiskCenterVULViewVULRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterVULViewVULRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterVULViewVULRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskClassInfoRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeRiskClassInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskClassInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebAssetListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeWebAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRiskEvidenceRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 风险项ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 风险类型，敏感信息-sentive、挂马暗链-horse、网页篡改-tamper

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
	// 链接类型，来源页面-source、目标页面-target

	URLType *string `json:"URLType,omitempty" name:"URLType"`
	// path路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
}

func (r *ExportRiskEvidenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRiskEvidenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportTaskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 需要定量返回的任务信息

	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`
}

func (r *DescribeExportTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFreeScanChanceRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateFreeScanChanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFreeScanChanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 急需关注

		EmergencyNum *int64 `json:"EmergencyNum,omitempty" name:"EmergencyNum"`
		// 急需关注高危、严重

		EmergencyHighNum *int64 `json:"EmergencyHighNum,omitempty" name:"EmergencyHighNum"`
		// 全量告警

		AllNum *int64 `json:"AllNum,omitempty" name:"AllNum"`
		// 所有高危告警

		AllHighNum *int64 `json:"AllHighNum,omitempty" name:"AllHighNum"`
		// 已生成调查事件

		DealNum *int64 `json:"DealNum,omitempty" name:"DealNum"`
		// 失陷资产

		LostAssets *int64 `json:"LostAssets,omitempty" name:"LostAssets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageBuildResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件

		Data []*AssetImageBuild `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageBuildResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageBuildResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SolidProtectionAssetData struct {

	// 主机资产总数

	CWPAssetTotal *int64 `json:"CWPAssetTotal,omitempty" name:"CWPAssetTotal"`
	// 未防护主机资产总数

	CWPUNDefenseTotal *int64 `json:"CWPUNDefenseTotal,omitempty" name:"CWPUNDefenseTotal"`
	// 公网IP资产总数

	PublicAssetTotal *int64 `json:"PublicAssetTotal,omitempty" name:"PublicAssetTotal"`
	// 未防护公网IP资产总数

	PublicUNDefenseTotal *int64 `json:"PublicUNDefenseTotal,omitempty" name:"PublicUNDefenseTotal"`
	// Web服务资产总数

	WebTotal *int64 `json:"WebTotal,omitempty" name:"WebTotal"`
	// 未防护Web服务资产总数

	WebUNDefenseTotal *int64 `json:"WebUNDefenseTotal,omitempty" name:"WebUNDefenseTotal"`
}

type TaskAssetObject struct {

	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 	资产类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 资产分类

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// ip/域名/资产id，数据库id等

	Asset *string `json:"Asset,omitempty" name:"Asset"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeListenerListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 监听器列表

		Data []*ClbListenerListInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListenerListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeListenerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalImageAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本地镜像

		Data []*LocalImageVO `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLocalImageAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLocalImageAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterWebsiteRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产视角的端口风险列表

		Data []*WebsiteRisk `json:"Data,omitempty" name:"Data"`
		// 状态列表

		StatusLists []*FilterDataObject `json:"StatusLists,omitempty" name:"StatusLists"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 资产类型列表

		InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`
		// 风险类型列表

		DetectEngineLists []*FilterDataObject `json:"DetectEngineLists,omitempty" name:"DetectEngineLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterWebsiteRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterWebsiteRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetDetailInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 子网详情

		Data *SubnetDetailInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetDetailInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafSaasPortsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// http端口列表

		HttpPorts []*string `json:"HttpPorts,omitempty" name:"HttpPorts"`
		// https端口列表

		HttpsPorts []*string `json:"HttpsPorts,omitempty" name:"HttpsPorts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafSaasPortsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafSaasPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanTaskInfo struct {

	// 任务日志Id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务日志名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务状态码：1等待开始  2正在扫描  3扫描出错 4扫描完成

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务进度

	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
	// 对应的展示时间

	TaskTime *string `json:"TaskTime,omitempty" name:"TaskTime"`
	// 报表id

	ReportId *string `json:"ReportId,omitempty" name:"ReportId"`
	// 报表名称

	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
	// 扫描计划，0-周期任务,1-立即扫描,2-定时扫描,3-自定义

	ScanPlan *int64 `json:"ScanPlan,omitempty" name:"ScanPlan"`
	// 关联的资产数

	AssetCount *int64 `json:"AssetCount,omitempty" name:"AssetCount"`
	// APP ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户主账户ID

	UIN *string `json:"UIN,omitempty" name:"UIN"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type TimeRange struct {

	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeAssetDiversityRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 资产类型

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
	// 统计类型

	StatType *int64 `json:"StatType,omitempty" name:"StatType"`
}

func (r *DescribeAssetDiversityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDiversityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseBindScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 进度

		Schedule *int64 `json:"Schedule,omitempty" name:"Schedule"`
		// 绑定任务详情

		List []*List `json:"List,omitempty" name:"List"`
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseBindScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseBindScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNeedToSolveListChangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 请求状态数据

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 具体信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNeedToSolveListChangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNeedToSolveListChangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetRiskMax struct {

	// web服务风险量

	WEB *int64 `json:"WEB,omitempty" name:"WEB"`
	// 主机风险量

	CVM *int64 `json:"CVM,omitempty" name:"CVM"`
	// 数据库风险量

	DB *int64 `json:"DB,omitempty" name:"DB"`
	// 网关风险量

	GateWay *int64 `json:"GateWay,omitempty" name:"GateWay"`
	// 网卡风险量

	NIC *int64 `json:"NIC,omitempty" name:"NIC"`
	// 域名风险量

	Domain *int64 `json:"Domain,omitempty" name:"Domain"`
	// 公网ip风险量

	PublicIP *int64 `json:"PublicIP,omitempty" name:"PublicIP"`
	// 容貌风险量

	Container *int64 `json:"Container,omitempty" name:"Container"`
	// 本地镜像风险量

	LocalImage *int64 `json:"LocalImage,omitempty" name:"LocalImage"`
	// 远程镜像风险量

	RepositoryImage *int64 `json:"RepositoryImage,omitempty" name:"RepositoryImage"`
	// 集群风险量

	Cluster *int64 `json:"Cluster,omitempty" name:"Cluster"`
}

type AssetViewCFGRisk struct {

	// 唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 配置名

	CFGName *string `json:"CFGName,omitempty" name:"CFGName"`
	// 检查类型

	CheckType *string `json:"CheckType,omitempty" name:"CheckType"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 影响资产

	AffectAsset *string `json:"AffectAsset,omitempty" name:"AffectAsset"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 来源

	From *string `json:"From,omitempty" name:"From"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// -

	CFGSTD *string `json:"CFGSTD,omitempty" name:"CFGSTD"`
	// 配置详情

	CFGDescribe *string `json:"CFGDescribe,omitempty" name:"CFGDescribe"`
	// 修复建议

	CFGFix *string `json:"CFGFix,omitempty" name:"CFGFix"`
	// 帮助文档链接

	CFGHelpURL *string `json:"CFGHelpURL,omitempty" name:"CFGHelpURL"`
	// 前端使用索引

	Index *string `json:"Index,omitempty" name:"Index"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type UserAccountInfo struct {

	// 用户账户到期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 资产路径配额

	URLQuota *uint64 `json:"URLQuota,omitempty" name:"URLQuota"`
	// 用户计费周期剩余扫描次数

	LeftScanCount *uint64 `json:"LeftScanCount,omitempty" name:"LeftScanCount"`
	// 用户试用：0 普通用户，1：试用用户

	IsWhiteCustomer *int64 `json:"IsWhiteCustomer,omitempty" name:"IsWhiteCustomer"`
	// 版本剩余天数

	SurplusDays *int64 `json:"SurplusDays,omitempty" name:"SurplusDays"`
	// 用户昵称

	UserNick *string `json:"UserNick,omitempty" name:"UserNick"`
	// 免费赠送次数

	FreeTryTimes *int64 `json:"FreeTryTimes,omitempty" name:"FreeTryTimes"`
	// 账号所属环境；DEVLOPMENT | PRODUCTION

	Envision *string `json:"Envision,omitempty" name:"Envision"`
	// 下次扫描次数刷新时间

	NextQuotaFlushTime *string `json:"NextQuotaFlushTime,omitempty" name:"NextQuotaFlushTime"`
	// 新计费购买类型

	BuyCompareValue *string `json:"BuyCompareValue,omitempty" name:"BuyCompareValue"`
	// 0不存在 1正常， 2过期，3销毁

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 0-新用户，1-老用户，2-老用户转成了新用户

	IsOld *int64 `json:"IsOld,omitempty" name:"IsOld"`
	// 是否展示水印,0-不展示,1-展示

	IsShowWaterMark *int64 `json:"IsShowWaterMark,omitempty" name:"IsShowWaterMark"`
	// 资产总数

	TotalAssets *int64 `json:"TotalAssets,omitempty" name:"TotalAssets"`
}

type CreateDomainAndIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建成功的数量

		Data *int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainAndIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDomainAndIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicIpAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// filte过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribePublicIpAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePublicIpAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSSLCertificatesRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。

	CertificateStatus []*uint64 `json:"CertificateStatus,omitempty" name:"CertificateStatus"`
}

func (r *DescribeSSLCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSSLCertificatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExamFreeLimitResp struct {

	// 还剩余可抢的免费体检名额

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 用户类型以及是否使用体检资格

	UserType []*FreeLimitUserType `json:"UserType,omitempty" name:"UserType"`
}

type GroupLastNumTwoNodeInfo struct {

	// 节点信息

	NodeInfo *GroupNodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`
	// 子节点信息

	SubNodeList []*GroupLastNumOneNodeInfo `json:"SubNodeList,omitempty" name:"SubNodeList"`
}

type SolidProtectionSingleApplyTrial struct {

	// 试用状态 0失败 1成功

	TrialStatus *int64 `json:"TrialStatus,omitempty" name:"TrialStatus"`
	// 试用失败原因，成功为空，主机成功则带有taskid

	TrialMessage *string `json:"TrialMessage,omitempty" name:"TrialMessage"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type TaskCenterCFGRiskInputParam struct {

	// 检测项ID

	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`
	// 是否开启，0-不开启，1-开启

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type DescribeAccountMapRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFWAssetStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络资产总数

		NetworkTotal *int64 `json:"NetworkTotal,omitempty" name:"NetworkTotal"`
		// 资产clb数量

		ClbTotal *int64 `json:"ClbTotal,omitempty" name:"ClbTotal"`
		// nat数量

		NatTotal *int64 `json:"NatTotal,omitempty" name:"NatTotal"`
		// 公网ip数量

		PublicAssetTotal *int64 `json:"PublicAssetTotal,omitempty" name:"PublicAssetTotal"`
		// 主机数量

		CVMAssetTotal *int64 `json:"CVMAssetTotal,omitempty" name:"CVMAssetTotal"`
		// 配置风险

		CFGTotal *int64 `json:"CFGTotal,omitempty" name:"CFGTotal"`
		// 端口风险

		PortTotal *int64 `json:"PortTotal,omitempty" name:"PortTotal"`
		// 内容风险

		WebsiteTotal *int64 `json:"WebsiteTotal,omitempty" name:"WebsiteTotal"`
		// 风险服务暴露

		ServerTotal *int64 `json:"ServerTotal,omitempty" name:"ServerTotal"`
		// 弱口令风险

		WeakPasswordTotal *int64 `json:"WeakPasswordTotal,omitempty" name:"WeakPasswordTotal"`
		// 漏洞风险

		VULTotal *int64 `json:"VULTotal,omitempty" name:"VULTotal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCFWAssetStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWAssetStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDbAssetInfoRequest struct {
	*tchttp.BaseRequest

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeDbAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterPortViewPortRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产视角的端口风险列表

		Data []*PortViewPortRisk `json:"Data,omitempty" name:"Data"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 处置建议列表

		SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitempty" name:"SuggestionLists"`
		// 来源列表

		FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterPortViewPortRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterPortViewPortRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterVULViewVULRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞产视角的漏洞风险列表

		Data []*VULViewVULRisk `json:"Data,omitempty" name:"Data"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 来源列表

		FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`
		// 漏洞类型列表

		VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitempty" name:"VULTypeLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterVULViewVULRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterVULViewVULRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreOrRetryScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态 0操作成功 1 失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 失败 错误码  0默认值 1 配额不足 目前就支持一种   后续扩展

		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIgnoreOrRetryScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreOrRetryScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBASAssetAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载命令

		RegisterDownloadCmd *string `json:"RegisterDownloadCmd,omitempty" name:"RegisterDownloadCmd"`
		// 执行命令

		RegisterExecCmd *string `json:"RegisterExecCmd,omitempty" name:"RegisterExecCmd"`
		// cwp执行命令

		RegisterCWPCmd *string `json:"RegisterCWPCmd,omitempty" name:"RegisterCWPCmd"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterBASAssetAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterBASAssetAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchExamStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该部分展示体检信息和状态

		Data []*AllExamInfoAndStatus `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSearchExamStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchExamStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVULRiskAdvanceCFGListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置项列表

		Data []*VULRiskAdvanceCFGList `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 风险等级过滤列表

		RiskLevelLists []*FilterDataObject `json:"RiskLevelLists,omitempty" name:"RiskLevelLists"`
		// 漏洞类型过滤列表

		VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitempty" name:"VULTypeLists"`
		// 识别来源过滤列表

		CheckFromLists []*FilterDataObject `json:"CheckFromLists,omitempty" name:"CheckFromLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVULRiskAdvanceCFGListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVULRiskAdvanceCFGListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOrganizationAccountStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值为0，则修改成功

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOrganizationAccountStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrganizationAccountStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubnetAsset struct {

	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 资产ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// 昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// cidr

	CIDR *string `json:"CIDR,omitempty" name:"CIDR"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// cvm数

	CVM *int64 `json:"CVM,omitempty" name:"CVM"`
	// 可用ip数

	AvailableIp *int64 `json:"AvailableIp,omitempty" name:"AvailableIp"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 配置风险

	ConfigureRisk *int64 `json:"ConfigureRisk,omitempty" name:"ConfigureRisk"`
	// 任务数

	ScanTask *int64 `json:"ScanTask,omitempty" name:"ScanTask"`
	// 最后扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
}

type SecOverviewInfo struct {

	// 公网IP资产总数

	IpAssets *uint64 `json:"IpAssets,omitempty" name:"IpAssets"`
	// 公网IP风险资产总数

	IPAssetsRisk *uint64 `json:"IPAssetsRisk,omitempty" name:"IPAssetsRisk"`
	// 域名资产总数

	DomainAssets *uint64 `json:"DomainAssets,omitempty" name:"DomainAssets"`
	// 域名风险资产总数

	DomainAssetsRisk *uint64 `json:"DomainAssetsRisk,omitempty" name:"DomainAssetsRisk"`
	// 漏洞风险总数

	VULRisk *uint64 `json:"VULRisk,omitempty" name:"VULRisk"`
	// 漏洞高危风险总数

	VULRiskHigh *uint64 `json:"VULRiskHigh,omitempty" name:"VULRiskHigh"`
	// 端口风险

	PortRisk *uint64 `json:"PortRisk,omitempty" name:"PortRisk"`
	// 端口高危风险

	PortRiskHigh *uint64 `json:"PortRiskHigh,omitempty" name:"PortRiskHigh"`
	// 网站风险

	WebRisk *uint64 `json:"WebRisk,omitempty" name:"WebRisk"`
	// 网站高危风险

	WebRiskHigh *uint64 `json:"WebRiskHigh,omitempty" name:"WebRiskHigh"`
	// 弱口令数量

	WeakPasswordTotal *int64 `json:"WeakPasswordTotal,omitempty" name:"WeakPasswordTotal"`
	// 高危弱口令数量

	WeakPasswordHighLevel *int64 `json:"WeakPasswordHighLevel,omitempty" name:"WeakPasswordHighLevel"`
	// 配置项风险数量

	CFGTotal *int64 `json:"CFGTotal,omitempty" name:"CFGTotal"`
	// 高危配置项风险数量

	CFGHighLevel *int64 `json:"CFGHighLevel,omitempty" name:"CFGHighLevel"`
}

type DescribeWafCiphersDetailRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
}

func (r *DescribeWafCiphersDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafCiphersDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskCenterRiskStatusRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 风险资产相关数据

	RiskStatusKeys []*RiskCenterStatusKey `json:"RiskStatusKeys,omitempty" name:"RiskStatusKeys"`
	// 处置状态，1为已处置、2为已忽略，3为取消已处置，4为取消已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 风险类型，0-端口风险， 1-漏洞风险，2-弱口令风险， 3-网站内容风险，4-配置风险

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyRiskCenterRiskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskCenterRiskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopBASTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopBASTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopBASTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GateWayAsset struct {

	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 资产ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 私有ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// 出向峰值带宽

	OutboundPeakBandwidth *string `json:"OutboundPeakBandwidth,omitempty" name:"OutboundPeakBandwidth"`
	// 入向峰值带宽

	InboundPeakBandwidth *string `json:"InboundPeakBandwidth,omitempty" name:"InboundPeakBandwidth"`
	// 出站累计流量

	OutboundCumulativeFlow *string `json:"OutboundCumulativeFlow,omitempty" name:"OutboundCumulativeFlow"`
	// 入站累计流量

	InboundCumulativeFlow *string `json:"InboundCumulativeFlow,omitempty" name:"InboundCumulativeFlow"`
	// 网络攻击

	NetworkAttack *int64 `json:"NetworkAttack,omitempty" name:"NetworkAttack"`
	// 暴露端口

	ExposedPort *int64 `json:"ExposedPort,omitempty" name:"ExposedPort"`
	// 暴露漏洞

	ExposedVUL *int64 `json:"ExposedVUL,omitempty" name:"ExposedVUL"`
	// 配置风险

	ConfigureRisk *int64 `json:"ConfigureRisk,omitempty" name:"ConfigureRisk"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务数

	ScanTask *int64 `json:"ScanTask,omitempty" name:"ScanTask"`
	// 最后扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// ipv6地址

	AddressIPV6 *string `json:"AddressIPV6,omitempty" name:"AddressIPV6"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
	// 风险服务暴露

	RiskExposure *int64 `json:"RiskExposure,omitempty" name:"RiskExposure"`
}

type GroupLastNumOneNodeInfo struct {

	// 节点信息

	NodeInfo *GroupNodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`
}

type DescribeAssetsStructureResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产类型总数

		Data []*AllAsset `json:"Data,omitempty" name:"Data"`
		// 资产总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 地域

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 资产类型

		InstanceList []*FilterDataObject `json:"InstanceList,omitempty" name:"InstanceList"`
		// 公网私网枚举

		PublicTypeList []*FilterDataObject `json:"PublicTypeList,omitempty" name:"PublicTypeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetsStructureResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsStructureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenerListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// -

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeListenerListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeListenerListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNICAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤参数

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeNICAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNICAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BASTacticsObject struct {

	// 战术名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 战技

	Techniques []*BASTechniquesObject `json:"Techniques,omitempty" name:"Techniques"`
	// 战术下子战技的总数

	SubTechniqueCnt *uint64 `json:"SubTechniqueCnt,omitempty" name:"SubTechniqueCnt"`
	// 战术的id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 战术icon的链接（灰）

	IconGray *string `json:"IconGray,omitempty" name:"IconGray"`
	// 战术icon的链接（蓝）

	IconBlue *string `json:"IconBlue,omitempty" name:"IconBlue"`
	// 战术icon的链接（选择）

	IconSelect *string `json:"IconSelect,omitempty" name:"IconSelect"`
	// 战术描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeAssetStatTopXResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// object统计数据

		ObjectTopX []*ObjectTopElement `json:"ObjectTopX,omitempty" name:"ObjectTopX"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetStatTopXResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetStatTopXResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateLogListRequest struct {
	*tchttp.BaseRequest

	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeOperateLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePortProtectStatusRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 来源  漏扫不传或者vss  新的风险端口界面itg

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *DescribePortProtectStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePortProtectStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanReportListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 列表过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeScanReportListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanReportListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbListenerInfo struct {

	// 监听器的ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器的名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 监听器协议，http、https

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听器端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 转发规则信息

	Rules []*ClbRuleInfo `json:"Rules,omitempty" name:"Rules"`
}

type DeleteDomainAndIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除的资产数量

		Data *int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDomainAndIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDomainAndIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePayInfoRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 查询来源 curr/order 查询当前信息还是订单信息 不传默认查当前

	QueryFrom *string `json:"QueryFrom,omitempty" name:"QueryFrom"`
}

func (r *DescribePayInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePayInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetsProtectStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1成功，其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncAssetsProtectStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetsProtectStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CFGViewCFGRisk struct {

	// 影响资产

	NoHandleCount *int64 `json:"NoHandleCount,omitempty" name:"NoHandleCount"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 状态，0未处理、1已处置、2已忽略

	AffectAssetCount *uint64 `json:"AffectAssetCount,omitempty" name:"AffectAssetCount"`
	// 资产唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 资产子类型

	From *string `json:"From,omitempty" name:"From"`
	// 前端索引

	Index *string `json:"Index,omitempty" name:"Index"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 配置名

	CFGName *string `json:"CFGName,omitempty" name:"CFGName"`
	// 检查类型

	CheckType *string `json:"CheckType,omitempty" name:"CheckType"`
	// -

	CFGSTD *string `json:"CFGSTD,omitempty" name:"CFGSTD"`
	// 描述

	CFGDescribe *string `json:"CFGDescribe,omitempty" name:"CFGDescribe"`
	// 修复建议

	CFGFix *string `json:"CFGFix,omitempty" name:"CFGFix"`
	// 帮助文档

	CFGHelpURL *string `json:"CFGHelpURL,omitempty" name:"CFGHelpURL"`
}

type CreateExamMissionResp struct {

	// 1成功，其他失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 相关信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DescribeClbListenersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据列表

		Data []*ClbListenerInfo `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClbListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClbListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SolidProtectionDataResp struct {

	// 超过多少XX选择三道防线

	ProductPercentage *float64 `json:"ProductPercentage,omitempty" name:"ProductPercentage"`
	// 产品状态数组

	ProductStatusData []*ProductStatusData `json:"ProductStatusData,omitempty" name:"ProductStatusData"`
	// 云防火墙防护公网IP数

	CFWProtectedIpCount *int64 `json:"CFWProtectedIpCount,omitempty" name:"CFWProtectedIpCount"`
	// 云防火墙待处理告警数

	CFWAlertCount *int64 `json:"CFWAlertCount,omitempty" name:"CFWAlertCount"`
	// 主机安全待处理风险数

	CWPRiskCount *int64 `json:"CWPRiskCount,omitempty" name:"CWPRiskCount"`
	// 主机安全已授权绑定主机数量

	CWPUsedLicenseCount *int64 `json:"CWPUsedLicenseCount,omitempty" name:"CWPUsedLicenseCount"`
	// waf已接入域名数量

	WAFDomainCount *int64 `json:"WAFDomainCount,omitempty" name:"WAFDomainCount"`
	// waf七天内攻击数量汇总

	WAFWebAttackCount *int64 `json:"WAFWebAttackCount,omitempty" name:"WAFWebAttackCount"`
	// soc七天内总告警数量

	SOCTotalAlert *int64 `json:"SOCTotalAlert,omitempty" name:"SOCTotalAlert"`
	// 云安全中心云上总资产数量

	SOCTotalAsset *int64 `json:"SOCTotalAsset,omitempty" name:"SOCTotalAsset"`
	// 容器安全待处理风险数

	TCSSRiskCount *int64 `json:"TCSSRiskCount,omitempty" name:"TCSSRiskCount"`
	// 容器安全防护实例数量

	TCSSCvmCount *int64 `json:"TCSSCvmCount,omitempty" name:"TCSSCvmCount"`
	// 恶意情报IOC数

	TICSBlackIOCNum *int64 `json:"TICSBlackIOCNum,omitempty" name:"TICSBlackIOCNum"`
	// 暴露风险数

	TICSExposeRiskNum *int64 `json:"TICSExposeRiskNum,omitempty" name:"TICSExposeRiskNum"`
	// Ddos待处理事件

	DDOSTotal *int64 `json:"DDOSTotal,omitempty" name:"DDOSTotal"`
	// Ddos防护ip+域名数量

	DDOSAllProtect *int64 `json:"DDOSAllProtect,omitempty" name:"DDOSAllProtect"`
	// 互联网边界防火墙是否开通 0 未开通  1开通

	CFWNatIsOpen *int64 `json:"CFWNatIsOpen,omitempty" name:"CFWNatIsOpen"`
	// NAT边界防火墙是否开通  0 未开通  1开通

	CFWInternetIsOpen *int64 `json:"CFWInternetIsOpen,omitempty" name:"CFWInternetIsOpen"`
	// Web应用防火墙-Saas型是否开通 0 未开通  1开通

	WAFSAASIsOpen *int64 `json:"WAFSAASIsOpen,omitempty" name:"WAFSAASIsOpen"`
	// Web应用防火墙-Clb型是否开通 0 未开通  1开通

	WAFCLBIsOpen *int64 `json:"WAFCLBIsOpen,omitempty" name:"WAFCLBIsOpen"`
	// 主机安全是否开通 0 未开通  1开通

	CWPIsOpen *int64 `json:"CWPIsOpen,omitempty" name:"CWPIsOpen"`
	// 容器安全是否开通 0 未开通  1开通

	TCSSIsOpen *int64 `json:"TCSSIsOpen,omitempty" name:"TCSSIsOpen"`
	// 安全运营中心是否开通 0 未开通  1开通

	SOCIsOpen *int64 `json:"SOCIsOpen,omitempty" name:"SOCIsOpen"`
	// 威胁情报是否开通 0未开通   1开通

	TICSIsOpen *int64 `json:"TICSIsOpen,omitempty" name:"TICSIsOpen"`
	// Ddos是否开通，0未开通，1开通

	DDOSIsOpen *int64 `json:"DDOSIsOpen,omitempty" name:"DDOSIsOpen"`
}

type TaskReportOverView struct {

	// 总任务数

	TaskNumber *int64 `json:"TaskNumber,omitempty" name:"TaskNumber"`
	// 周期扫描任务数

	CycleTaskNumber *int64 `json:"CycleTaskNumber,omitempty" name:"CycleTaskNumber"`
	// 正在扫描任务数

	TaskNowNumber *int64 `json:"TaskNowNumber,omitempty" name:"TaskNowNumber"`
	// 扫描次数，报告数量

	TaskLogNumber *int64 `json:"TaskLogNumber,omitempty" name:"TaskLogNumber"`
	// 扫描次数总配额

	ScanQuotaTotal *int64 `json:"ScanQuotaTotal,omitempty" name:"ScanQuotaTotal"`
	// 扫描次数使用的配额

	ScanQuotaUsed *int64 `json:"ScanQuotaUsed,omitempty" name:"ScanQuotaUsed"`
	// 体检任务总配额

	TaskQuotaTotal *int64 `json:"TaskQuotaTotal,omitempty" name:"TaskQuotaTotal"`
	// 体检任务已经使用的配额

	TaskScanUsed *int64 `json:"TaskScanUsed,omitempty" name:"TaskScanUsed"`
}

type WafInstancesFilter struct {

	// 字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数值

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 是否精确匹配 true/false

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type DescribeDomainAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// -

		Data []*DomainAssetVO `json:"Data,omitempty" name:"Data"`
		// 防护状态列表

		DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitempty" name:"DefenseStatusList"`
		// 资产归属地列表

		AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitempty" name:"AssetLocationList"`
		// 资产类型列表

		SourceTypeList []*FilterDataObject `json:"SourceTypeList,omitempty" name:"SourceTypeList"`
		// 地域列表

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePayInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本 Free/Advanced/Enterprise/Ultimate  分别是免费版/高级版/企业版/旗舰版

		Version *string `json:"Version,omitempty" name:"Version"`
		// 订单状态 0/1/2/3/6/7
		// 0免费版/1正常支付/2(隔离)/3 销毁/6 试用中/7订单到期

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 过期时间 yyyy-MM-dd hh:mm:ss

		ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
		// 配额信息列表

		QuotaList []*OrderQuotaInfo `json:"QuotaList,omitempty" name:"QuotaList"`
		// 时长

		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
		// 时间单位 y/m/d 年/月/日

		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
		// 创建时间  yyyy-MM-dd hh:mm:ss

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 自动续费 0不开启自动续费 1开启自动续费

		AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
		// 资源id

		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
		// 支付模式 1预付费   csip目前只支持预付费

		PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
		// 试用状态 0未试用 1试用中

		TrailStatus *int64 `json:"TrailStatus,omitempty" name:"TrailStatus"`
		// 试用次数

		TrailNum *int64 `json:"TrailNum,omitempty" name:"TrailNum"`
		// 用户收入等级

		UserLevel *int64 `json:"UserLevel,omitempty" name:"UserLevel"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePayInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePayInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanTaskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeScanTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLastTaskOverviewRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// vss/csip,默认vss

	ScanFrom *string `json:"ScanFrom,omitempty" name:"ScanFrom"`
}

func (r *DescribeLastTaskOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLastTaskOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncScopeAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步是否成功 1：成功 0：失败

		SyncResult *int64 `json:"SyncResult,omitempty" name:"SyncResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncScopeAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncScopeAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterDataObject struct {

	// 英文翻译

	Value *string `json:"Value,omitempty" name:"Value"`
	// 中文翻译

	Text *string `json:"Text,omitempty" name:"Text"`
}

type DescribeAlarmTrendRequest struct {
	*tchttp.BaseRequest

	// 16位时间戳函数

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 16位时间戳函数

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAlarmTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExaminationPayRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeExaminationPayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExaminationPayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchBugInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// id=3时传入该参数

	CVEId *string `json:"CVEId,omitempty" name:"CVEId"`
}

func (r *DescribeSearchBugInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchBugInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagCountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各大类日志数量

		TagCounts []*TagCount `json:"TagCounts,omitempty" name:"TagCounts"`
		// 返回数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回代码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 结果信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagCountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagCountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskReportInfoRequest struct {
	*tchttp.BaseRequest

	// vss/csip,默认vss

	ScanFrom *string `json:"ScanFrom,omitempty" name:"ScanFrom"`
}

func (r *DescribeTaskReportInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskReportInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RepositoryImageVO struct {

	// 用户appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 镜像id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 镜像名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 镜像创建时间

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
	// 镜像大小带单位

	InstanceSize *string `json:"InstanceSize,omitempty" name:"InstanceSize"`
	// 构建次数

	BuildCount *int64 `json:"BuildCount,omitempty" name:"BuildCount"`
	// 镜像类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 授权状态

	AuthStatus *int64 `json:"AuthStatus,omitempty" name:"AuthStatus"`
	// 镜像版本

	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 仓库地址

	RepositoryUrl *string `json:"RepositoryUrl,omitempty" name:"RepositoryUrl"`
	// 仓库名称

	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
	// 漏洞风险

	VulRisk *int64 `json:"VulRisk,omitempty" name:"VulRisk"`
	// 检查任务

	CheckCount *int64 `json:"CheckCount,omitempty" name:"CheckCount"`
	// 体检时间

	CheckTime *string `json:"CheckTime,omitempty" name:"CheckTime"`
}

type StopRiskCenterTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Status为0， 停止成功

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopRiskCenterTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopRiskCenterTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNeedToSolveListChangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 请求结果

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 需要的数据

		Data *ToDoListRecordResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNeedToSolveListChangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNeedToSolveListChangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskAssetTopFiveListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险资产TOP5列表

		Data []*RiskOverviewObject `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskAssetTopFiveListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskAssetTopFiveListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterRiskOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产概况数据

		Data *RiskCenterRiskOverview `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterRiskOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterRiskOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLicenseBindsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 无

	LicenseType *int64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 无

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 无

	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`
	// 无

	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`
}

func (r *ModifyLicenseBindsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLicenseBindsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogTypeStorage struct {

	// 日志类型

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 日志总存储量，单位Byte

	StorageUsed *int64 `json:"StorageUsed,omitempty" name:"StorageUsed"`
	// 日志总存储条数

	StorageCount *int64 `json:"StorageCount,omitempty" name:"StorageCount"`
}

type DescribeBASTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总个数

		TotalCnt *uint64 `json:"TotalCnt,omitempty" name:"TotalCnt"`
		// 任务列表

		Tasks []*BASTaskObject `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBASTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBASTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表

		List []*LicenseDetail `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNeedToSolveListChangeRequest struct {
	*tchttp.BaseRequest

	// 第几页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数字

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询结构

	Search *ToDoListReq `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeNeedToSolveListChangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNeedToSolveListChangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterAssetViewPortRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产视角的配置风险列表

		Data []*AssetViewPortRisk `json:"Data,omitempty" name:"Data"`
		// 状态列表

		StatusLists []*FilterDataObject `json:"StatusLists,omitempty" name:"StatusLists"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 建议列表

		SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitempty" name:"SuggestionLists"`
		// 资产类型列表

		InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`
		// 来源列表

		FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewPortRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterAssetViewPortRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserOperationPermissionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserOperationPermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserOperationPermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOrganizationAccountStatusRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 修改集团账号状态，1 开启， 2关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyOrganizationAccountStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOrganizationAccountStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetViewPortRisk struct {

	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 影响资产

	AffectAsset *string `json:"AffectAsset,omitempty" name:"AffectAsset"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 资产类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 服务

	Service *string `json:"Service,omitempty" name:"Service"`
	// 最近识别时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// 首次识别时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 处置建议,0保持现状、1限制访问、2封禁端口

	Suggestion *uint64 `json:"Suggestion,omitempty" name:"Suggestion"`
	// 状态，0未处理、1已处置、2已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 资产唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 前端索引

	Index *string `json:"Index,omitempty" name:"Index"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 用户appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 用户昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 来源

	From *string `json:"From,omitempty" name:"From"`
}

type DescribeAssetVpcDetailInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// vpc详情

		Data *VpcDetailInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetVpcDetailInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetVpcDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetsProcessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1：已完成 0:未完成

		SyncProcess *int64 `json:"SyncProcess,omitempty" name:"SyncProcess"`
		// 同步时间

		SyncTime *string `json:"SyncTime,omitempty" name:"SyncTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncAssetsProcessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetsProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationPermission struct {

	// 限制名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否启用 1为启用 0为未启用

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
}

type ParamInfo struct {

	// 子类型

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 数量

	Sum *int64 `json:"Sum,omitempty" name:"Sum"`
	// 处置动作

	Action *string `json:"Action,omitempty" name:"Action"`
}

type DescribeSocLogStorageSpaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		UsedStorage *float64 `json:"UsedStorage,omitempty" name:"UsedStorage"`
		// 无

		TotalStorage *float64 `json:"TotalStorage,omitempty" name:"TotalStorage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSocLogStorageSpaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSocLogStorageSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskReportInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 输出数据

		Data *TaskOverView `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskReportInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskReportInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetClusterService struct {

	// 租户id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 服务 id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 服务名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 服务类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 服务创建时间

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// selectors

	Selectors *string `json:"Selectors,omitempty" name:"Selectors"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
}

type WafSpartaPortItem struct {

	// 监听端口配置

	Port *string `json:"Port,omitempty" name:"Port"`
	// 与Port一一对应，表示端口对应的协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 与Port一一对应, 表示回源端口

	UpstreamPort *string `json:"UpstreamPort,omitempty" name:"UpstreamPort"`
	// 与Port一一对应, 表示回源协议

	UpstreamProtocol *string `json:"UpstreamProtocol,omitempty" name:"UpstreamProtocol"`
	// Nginx的服务器ID

	NginxServerId *string `json:"NginxServerId,omitempty" name:"NginxServerId"`
}

type DescribeUserLogAuditSettingRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeUserLogAuditSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserLogAuditSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNewBindRoleUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0成功，其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNewBindRoleUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNewBindRoleUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHighLevelRiskTopFiveListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeHighLevelRiskTopFiveListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHighLevelRiskTopFiveListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterAssetViewCFGRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterAssetViewCFGRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterAssetViewCFGRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterAssetViewVULRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterAssetViewVULRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterAssetViewVULRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskCostQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消耗配额

		Data *int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskCostQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskCostQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafCertificateVerifyResultRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 证书类型

	CertType *int64 `json:"CertType,omitempty" name:"CertType"`
	// 证书

	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`
	// 证书Id

	CertID *string `json:"CertID,omitempty" name:"CertID"`
	// 私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

func (r *DescribeWafCertificateVerifyResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafCertificateVerifyResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafLoadBalancer struct {

	// 负载均衡LD的ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡LD的名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡监听器的ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 负载均衡监听器的名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 负载均衡实例的IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 负载均衡实例的端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 负载均衡LD的地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 监听器协议，http、https

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 负载均衡监听器所在的zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 负载均衡的VPCID，公网为-1，内网按实际填写

	NumericalVpcId *int64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`
	// 负载均衡的网络类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
}

type DescribeClbListenersRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 负载均衡实例id

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡地域

	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`
}

func (r *DescribeClbListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClbListenersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainAssetInfoRequest struct {
	*tchttp.BaseRequest

	// -

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeDomainAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterAssetViewWeakPasswordRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CWPMachine struct {

	// 主机qquid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 机器名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 机器ip

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 机器状态

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 资产信息

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 资产iD

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 资产uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type VpcDetailInfo struct {

	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc name

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// cidr地址

	CIDR *string `json:"CIDR,omitempty" name:"CIDR"`
	// dns地址

	DNS *string `json:"DNS,omitempty" name:"DNS"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// Tag标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// 是否默认vpc

	IsDefaultVpc *int64 `json:"IsDefaultVpc,omitempty" name:"IsDefaultVpc"`
	// hvip数量

	HAVIPNum *int64 `json:"HAVIPNum,omitempty" name:"HAVIPNum"`
	// nat数量

	NATFWNum *int64 `json:"NATFWNum,omitempty" name:"NATFWNum"`
	// clb数量

	CLBNum *int64 `json:"CLBNum,omitempty" name:"CLBNum"`
	// eni数量

	ENINum *int64 `json:"ENINum,omitempty" name:"ENINum"`
	// 网关数量

	GatewayNum *int64 `json:"GatewayNum,omitempty" name:"GatewayNum"`
	// 网卡数量

	NetworkCardNum *int64 `json:"NetworkCardNum,omitempty" name:"NetworkCardNum"`
	// 服务器数量

	ServerNum *int64 `json:"ServerNum,omitempty" name:"ServerNum"`
	// 子网数量

	SubnetNum *int64 `json:"SubnetNum,omitempty" name:"SubnetNum"`
}

type ApplyTrialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 状态 为0成功 其他值失败

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyTrialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerInfoRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeContainerInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDbAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// -

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeDbAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafCertificateVerifyResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 详细错误信息

		Detail []*string `json:"Detail,omitempty" name:"Detail"`
		// 过期时间

		NotAfter *string `json:"NotAfter,omitempty" name:"NotAfter"`
		// 证书是否改变:1有改变，0没有改变

		Changed *int64 `json:"Changed,omitempty" name:"Changed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWafCertificateVerifyResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWafCertificateVerifyResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveAssetsConfigRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 配置类型：1: web， 2: api

	ConfigType *int64 `json:"ConfigType,omitempty" name:"ConfigType"`
	// 扫描目标：域名或者ip

	Target *string `json:"Target,omitempty" name:"Target"`
	// 扫描端口：多个端口用英文逗号隔开

	Port *string `json:"Port,omitempty" name:"Port"`
	// 扫描路径：默认路径为/

	Path *string `json:"Path,omitempty" name:"Path"`
	// 扫描描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 鉴权类型：0:无鉴权，1: apikey鉴权，2:basic鉴权，3:cookie鉴权

	AuthMethod *int64 `json:"AuthMethod,omitempty" name:"AuthMethod"`
	// Cookie值，AuthMethod为3时有效

	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
	// apikey的位置，AuthMethod为1时有效，1: head上添加，2: url上添加

	AddTo *int64 `json:"AddTo,omitempty" name:"AddTo"`
	// apikey的位置，AuthMethod为1时有效, 添加到head或者url的参数，跟AddTo配对使用，格式为json字符串

	AddToValue *string `json:"AddToValue,omitempty" name:"AddToValue"`
	// 鉴权用户名，AuthMethod为2时有效

	User *string `json:"User,omitempty" name:"User"`
	// 鉴权密码，AuthMethod为2时有效

	Password *string `json:"Password,omitempty" name:"Password"`
	// api参数

	APIS []*AssetsApiConfig `json:"APIS,omitempty" name:"APIS"`
	// 配置id

	Id *string `json:"Id,omitempty" name:"Id"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
}

func (r *SaveAssetsConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveAssetsConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 报告列表

		Data []*TaskLogInfo `json:"Data,omitempty" name:"Data"`
		// 待查看数量

		NotViewNumber *int64 `json:"NotViewNumber,omitempty" name:"NotViewNumber"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveAssetsConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveAssetsConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveAssetsConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BASTechniquesObject struct {

	// 战技名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 战技描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 子战技

	SubTechniques []*BASSubTechniquesObject `json:"SubTechniques,omitempty" name:"SubTechniques"`
	// 战技的id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 战技关联剧本的个数

	ScriptCnt *uint64 `json:"ScriptCnt,omitempty" name:"ScriptCnt"`
}

type DescribeAssetImageBuildRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeAssetImageBuildRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageBuildRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterAssetViewPortRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterAssetViewPortRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterAssetViewPortRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功，0失败 1成功

		Success *uint64 `json:"Success,omitempty" name:"Success"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopRiskCenterTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务id 列表

	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitempty" name:"TaskIdList"`
}

func (r *StopRiskCenterTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopRiskCenterTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenterAssetRiskMaxRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeCenterAssetRiskMaxRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenterAssetRiskMaxRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSolidAssetStaticRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSolidAssetStaticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSolidAssetStaticRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIpProtectRequest struct {
	*tchttp.BaseRequest

	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 防护状态，0， 1

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 拦截模式，0，1，2，3

	IDPSAction *int64 `json:"IDPSAction,omitempty" name:"IDPSAction"`
	// 公网ip，如果没有，则表示一键开启防护

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 如果是一键开启，这里值为1，如果是单个公网ip，这里值为0

	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

func (r *ModifyIpProtectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpProtectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTkeAuthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTkeAuthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTkeAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 弱口令视角的弱口令风险列表

		Data []*WeakPasswordViewWeakPassRisk `json:"Data,omitempty" name:"Data"`
		// 危险等级列表

		LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`
		// 来源列表

		FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`
		// 弱口令类型列表

		PasswordTypeLists []*FilterDataObject `json:"PasswordTypeLists,omitempty" name:"PasswordTypeLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessSwitch struct {

	// 接入产品的tag

	Tag *uint64 `json:"Tag,omitempty" name:"Tag"`
}

type Filter struct {

	// 查询数量限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询偏移位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序采用升序还是降序 升:asc 降 desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 需排序的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤的列及内容

	Filters []*WhereFilter `json:"Filters,omitempty" name:"Filters"`
	// 可填无， 日志使用查询时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 可填无， 日志使用查询时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type SyncAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步是否成功 1：成功 0：失败

		SyncResult *int64 `json:"SyncResult,omitempty" name:"SyncResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbRuleInfo struct {

	// 转发规则的域名。

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 转发规则的路径。

	Url *string `json:"Url,omitempty" name:"Url"`
	// 转发规则的 ID

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type IpAssetListVO struct {

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产name

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 云防状态

	CFWStatus *uint64 `json:"CFWStatus,omitempty" name:"CFWStatus"`
	// 资产创建时间

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 公网ip类型

	PublicIpType *uint64 `json:"PublicIpType,omitempty" name:"PublicIpType"`
	// vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc名

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
	// 云上

	IsCloud *uint64 `json:"IsCloud,omitempty" name:"IsCloud"`
	// 网络攻击

	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`
	// 网络访问

	Access *uint64 `json:"Access,omitempty" name:"Access"`
	// 网络拦截

	Intercept *uint64 `json:"Intercept,omitempty" name:"Intercept"`
	// 入向带宽

	InBandwidth *string `json:"InBandwidth,omitempty" name:"InBandwidth"`
	// 出向带宽

	OutBandwidth *string `json:"OutBandwidth,omitempty" name:"OutBandwidth"`
	// 入向流量

	InFlow *string `json:"InFlow,omitempty" name:"InFlow"`
	// 出向流量

	OutFlow *string `json:"OutFlow,omitempty" name:"OutFlow"`
	// 最近扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 端口风险

	PortRisk *uint64 `json:"PortRisk,omitempty" name:"PortRisk"`
	// 漏洞风险

	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitempty" name:"VulnerabilityRisk"`
	// 配置风险

	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitempty" name:"ConfigurationRisk"`
	// 扫描任务

	ScanTask *uint64 `json:"ScanTask,omitempty" name:"ScanTask"`
	// 弱口令

	WeakPassword *uint64 `json:"WeakPassword,omitempty" name:"WeakPassword"`
	// 内容风险

	WebContentRisk *uint64 `json:"WebContentRisk,omitempty" name:"WebContentRisk"`
	// 标签

	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
	// eip主键

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// memberid信息

	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
	// 风险服务暴露

	RiskExposure *int64 `json:"RiskExposure,omitempty" name:"RiskExposure"`
}

type ObjectTrend struct {

	// 当前时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 统计数据

	NodeInfo []*NodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`
}

type OrganizationInfo struct {

	// 成员账号名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 部门节点名称，账号所属部门

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// Member/Admin;成员或者管理员

	Role *string `json:"Role,omitempty" name:"Role"`
	// 成员账号id

	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
	// 账号加入方式,create/invite

	JoinType *string `json:"JoinType,omitempty" name:"JoinType"`
	// 集团名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 管理员账号名称

	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`
	// 管理员Uin

	AdminUin *string `json:"AdminUin,omitempty" name:"AdminUin"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 部门数

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 成员数

	MemberCount *int64 `json:"MemberCount,omitempty" name:"MemberCount"`
	// 子账号数

	SubAccountCount *int64 `json:"SubAccountCount,omitempty" name:"SubAccountCount"`
	// 集团关系策略权限

	GroupPermission []*string `json:"GroupPermission,omitempty" name:"GroupPermission"`
	// 成员关系策略权限

	MemberPermission []*string `json:"MemberPermission,omitempty" name:"MemberPermission"`
	// 集团付费模式；0/自付费，1/代付费

	GroupPayMode *int64 `json:"GroupPayMode,omitempty" name:"GroupPayMode"`
	// 个人付费模式；0/自付费，1/代付费

	MemberPayMode *int64 `json:"MemberPayMode,omitempty" name:"MemberPayMode"`
	// 所有部门的集合

	Department *string `json:"Department,omitempty" name:"Department"`
	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本

	CFWProtect *string `json:"CFWProtect,omitempty" name:"CFWProtect"`
	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本

	WAFProtect *string `json:"WAFProtect,omitempty" name:"WAFProtect"`
	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本

	CWPProtect *string `json:"CWPProtect,omitempty" name:"CWPProtect"`
	// 所有部门的集合数组

	Departments []*string `json:"Departments,omitempty" name:"Departments"`
	// 成员创建时间

	MemberCreateTime *string `json:"MemberCreateTime,omitempty" name:"MemberCreateTime"`
	// 集团关系策略权限

	GroupPermissionPlain *string `json:"GroupPermissionPlain,omitempty" name:"GroupPermissionPlain"`
	// 成员关系策略权限

	MemberPermissionPlain *string `json:"MemberPermissionPlain,omitempty" name:"MemberPermissionPlain"`
	// 启用的成员数量

	EnableMemberCount *int64 `json:"EnableMemberCount,omitempty" name:"EnableMemberCount"`
	// 购买的成员配额

	MemberQuota *int64 `json:"MemberQuota,omitempty" name:"MemberQuota"`
}

type DescribeABTestConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 灰度项目配置

		Config []*ABTestConfig `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeABTestConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeABTestConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetsProcessRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *SyncAssetsProcessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetsProcessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*SubnetAsset `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域列表

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// vpc列表

		VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`
		// appid列表

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 可用区列表

		ZoneList []*FilterDataObject `json:"ZoneList,omitempty" name:"ZoneList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRiskCenterScanTaskRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 扫描资产信息列表

	Assets []*TaskAssetObject `json:"Assets,omitempty" name:"Assets"`
	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填

	ScanAssetType *int64 `json:"ScanAssetType,omitempty" name:"ScanAssetType"`
	// 扫描项目；port/poc/weakpass/webcontent/configrisk

	ScanItem []*string `json:"ScanItem,omitempty" name:"ScanItem"`
	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填

	ScanPlanType *int64 `json:"ScanPlanType,omitempty" name:"ScanPlanType"`
	// 扫描计划详情

	ScanPlanContent *string `json:"ScanPlanContent,omitempty" name:"ScanPlanContent"`
	// ip/域名/url数组

	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitempty" name:"SelfDefiningAssets"`
	// 要修改的任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// APP ID

	TargetAppId *string `json:"TargetAppId,omitempty" name:"TargetAppId"`
	// vss/csip

	ScanFrom *string `json:"ScanFrom,omitempty" name:"ScanFrom"`
	// 高级配置

	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitempty" name:"TaskAdvanceCFG"`
	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式

	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
}

func (r *ModifyRiskCenterScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRiskCenterScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerVO struct {

	// appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 资产名

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 私有ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 资产名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// pod id

	PodId *string `json:"PodId,omitempty" name:"PodId"`
	// pod ip

	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 机器id

	MachineId *string `json:"MachineId,omitempty" name:"MachineId"`
	// 机器名

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 资产创建时间

	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`
	// 网络状态

	NetworkStatus *string `json:"NetworkStatus,omitempty" name:"NetworkStatus"`
	// cpu使用率

	CpuRate *string `json:"CpuRate,omitempty" name:"CpuRate"`
	// 内存大小

	MemCount *string `json:"MemCount,omitempty" name:"MemCount"`
	// 网络状态原因

	NetworkStatusReason *string `json:"NetworkStatusReason,omitempty" name:"NetworkStatusReason"`
	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 是否核心

	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
	// 漏洞风险

	VulRisk *int64 `json:"VulRisk,omitempty" name:"VulRisk"`
	// 配置风险

	CfgRisk *int64 `json:"CfgRisk,omitempty" name:"CfgRisk"`
	// 体检任务数

	CheckCount *int64 `json:"CheckCount,omitempty" name:"CheckCount"`
	// 体检时间

	CheckTime *string `json:"CheckTime,omitempty" name:"CheckTime"`
	// 内存比率

	MemRate *string `json:"MemRate,omitempty" name:"MemRate"`
	// cpu核数

	CpuCount *int64 `json:"CpuCount,omitempty" name:"CpuCount"`
}

type UpdateBugInfoResp struct {

	// 2代表更新成功，3代表更新出错，4代表主机修复后更新成功

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeUserAuditTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志趋势图

		AuditTrend []*ObjectTrend `json:"AuditTrend,omitempty" name:"AuditTrend"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserAuditTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAuditTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQuickProtectSettingRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 被调用的集团账号的成员id

	OperatedMemberId []*string `json:"OperatedMemberId,omitempty" name:"OperatedMemberId"`
	// 	快捷防护类型：Vul-漏洞检测、Malware-文件查杀、Baseline-基线检测、AutoBan-自动阻断

	Type []*string `json:"Type,omitempty" name:"Type"`
	// 	Quuid数组，云服务器uuid，最大100条

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *CreateQuickProtectSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQuickProtectSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainAndIpRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// -

	Content []*PublicIpDomainListKey `json:"Content,omitempty" name:"Content"`
	// 是否保留路径配置，1：保留，其他：不保留，默认不传为不保留

	RetainPath *int64 `json:"RetainPath,omitempty" name:"RetainPath"`
}

func (r *DeleteDomainAndIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDomainAndIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCenterAssetRiskMaxResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data *AssetRiskMax `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCenterAssetRiskMaxResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCenterAssetRiskMaxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDbAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 资产总数

		Data []*DBAssetVO `json:"Data,omitempty" name:"Data"`
		// 地域枚举

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 资产类型枚举

		AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`
		// Vpc枚举

		VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`
		// Appid枚举

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDbAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalImageAssetsRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// filter过滤条件

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeLocalImageAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLocalImageAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 过滤内容

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskCenterWeakPasswordViewWeakPasswordRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserStorageRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserStorageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserStorageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportCSVRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 筛选过滤参数

	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
	// 筛选过滤的接口

	ExportAction *string `json:"ExportAction,omitempty" name:"ExportAction"`
	// 需要导出的列

	ExportFields []*string `json:"ExportFields,omitempty" name:"ExportFields"`
}

func (r *ExportCSVRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportCSVRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetTopXResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// top统计

		Data []*TopX `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetTopXResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTopXResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*AccessSwitch `json:"Data,omitempty" name:"Data"`
		// 返回代码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProtectTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询成功状态

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 需要的时长和过期时间

		Data *TimeData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProtectTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SSLCertificate struct {

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。

	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`
	// 域名。

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 状态。0：审核中，1：已通过，2：审核失败，3：已过期，4：验证方式为 DNS_AUTO 类型的证书， 已添加DNS记录，5：企业证书，待提交，6：订单取消中，7：已取消，8：已提交资料， 待上传确认函，9：证书吊销中，10：已吊销，11：重颁发中，12：待上传吊销确认函，13：免费证书待提交资料状态，14：已退款，
	//

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 证书生效时间。

	CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`
	// 证书过期时间。

	CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`
	// 证书有效期，单位（月）。

	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`
	// 证书 ID。

	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
	// 证书套餐类型：1 = GeoTrust DV SSL CA - G3， 2 = TrustAsia TLS RSA CA， 3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书。

	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

type SyncUserAssetInfoRequest struct {
	*tchttp.BaseRequest

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产类型 CVM等

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 资产地域

	Location *string `json:"Location,omitempty" name:"Location"`
}

func (r *SyncUserAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncUserAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbListenerListInfo struct {

	// 监听器id

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 负载均衡Id

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 负载均衡ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 端口

	VPort *int64 `json:"VPort,omitempty" name:"VPort"`
	// 区域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 私有网络id

	NumericalVpcId *int64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`
	// 负载均衡类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 监听器域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type CreateApplyTrialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data *ApplyTrial `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplyTrialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplyTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetStatTopXRequest struct {
	*tchttp.BaseRequest

	// 集团账号的成员id

	MemberId []*string `json:"MemberId,omitempty" name:"MemberId"`
	// 资产类型 1:公网ip 2:网卡
	// 3:域名

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
	// 统计类别

	StatType *int64 `json:"StatType,omitempty" name:"StatType"`
	// 时间戳

	TimeRange *int64 `json:"TimeRange,omitempty" name:"TimeRange"`
}

func (r *DescribeAssetStatTopXRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetStatTopXRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBASATTCKListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 战术个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 战技个数

		TotalTechCount *uint64 `json:"TotalTechCount,omitempty" name:"TotalTechCount"`
		// 战术详情

		Tactics []*BASTacticsObject `json:"Tactics,omitempty" name:"Tactics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBASATTCKListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBASATTCKListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFWStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCFWStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*AssetCluster `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群类型枚举

		ClusterTypeList []*FilterDataObject `json:"ClusterTypeList,omitempty" name:"ClusterTypeList"`
		// 集群状态枚举

		ClusterStatusList []*FilterDataObject `json:"ClusterStatusList,omitempty" name:"ClusterStatusList"`
		// 组件状态枚举

		ComponentStatusList []*FilterDataObject `json:"ComponentStatusList,omitempty" name:"ComponentStatusList"`
		// 私有网络枚举

		VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`
		// 地域枚举

		RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`
		// 租户枚举

		AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanAssetByTaskIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产信息

		Assets []*TaskAssetObject `json:"Assets,omitempty" name:"Assets"`
		// 0-全扫，1-指定资产扫，2-排除资产扫

		ScanAssetType *int64 `json:"ScanAssetType,omitempty" name:"ScanAssetType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanAssetByTaskIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanAssetByTaskIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExitExamRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyExitExamRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExitExamRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PushTaskResult struct {

	// 无

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 无

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 无

	IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`
	// 无

	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`
}

type Intelligence struct {

	// 威胁情报标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 威胁情报ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 威胁等级：info 提示- low低危- middle中危-high高危

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 发布时间

	PubTime *string `json:"PubTime,omitempty" name:"PubTime"`
	// 威胁情报来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 关键词

	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`
	// 该威胁情报用于针对扫描可能需要的资产ID类别：-1为不扫描，0为网站，1为主机，2为网站+主机

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
	// 该威胁情报用于针对扫描可能需要的poc_id值

	POCId *string `json:"POCId,omitempty" name:"POCId"`
	// 情报危险详情

	Content *string `json:"Content,omitempty" name:"Content"`
}

type DeleteWebServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWebServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWebServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanReportListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务日志列表

		Data []*ScanTaskInfo `json:"Data,omitempty" name:"Data"`
		// 主账户ID列表

		UINList []*string `json:"UINList,omitempty" name:"UINList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanReportListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
