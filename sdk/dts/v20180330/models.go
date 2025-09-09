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

package v20180330

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ModifyConsumerGroupDescriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅任务ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 消费组账户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 消费组名称

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
}

func (r *ModifyConsumerGroupDescriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerGroupDescriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternParams struct {
}

type SyncInstanceInfo struct {

	// 地域英文名，如：ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例短ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 是否强制检查

	ForceCheck *bool `json:"ForceCheck,omitempty" name:"ForceCheck"`
}

func (r *CreateMigrateCheckJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMigrateCheckJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeConsumeTimeRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 消费时间起点，也即是指定订阅数据的时间起点，时间格式如：Y-m-d h:m:s，取值范围为过去24小时之内

	ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`
}

func (r *ModifySubscribeConsumeTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeConsumeTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DstInfo struct {

	// 目标实例ID，如cdb-jd92ijd8

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 目标实例vip。已废弃，无需填写

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 目标实例vport。已废弃，无需填写

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 目标实例地域，如ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 目前只对MySQL有效。当为整实例迁移时，1-只读，0-可读写。

	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 目标数据库账号

	User *string `json:"User,omitempty" name:"User"`
	// 目标数据库密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type DescribeConnectTestResultRequest struct {
	*tchttp.BaseRequest

	// 连接性检查任务ID

	TaskIds []*int64 `json:"TaskIds,omitempty" name:"TaskIds"`
}

func (r *DescribeConnectTestResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectTestResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaSubscribeConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订阅实例ID

		SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
		// 订阅实例名称

		SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`
		// 订阅数据库类型

		Product *string `json:"Product,omitempty" name:"Product"`
		// 被订阅的实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 被订阅的实例的状态，可能的值有running,offline,isolate

		InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
		// 订阅实例状态，可能的值有unconfigure-未配置，configuring-配置中，configured-已配置

		SubsStatus *string `json:"SubsStatus,omitempty" name:"SubsStatus"`
		// 订阅实例修改时间

		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
		// 订阅实例创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 订阅实例被隔离时间

		IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
		// 订阅实例到期时间

		ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
		// 订阅实例下线时间

		OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
		// 自动续费标识,0-不自动续费，1-自动续费

		AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
		// 地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 订阅实例计费类型，1-小时计费，0-包年包月

		PayType *int64 `json:"PayType,omitempty" name:"PayType"`
		// 订阅实例生命周期状态，可能的值有：normal-正常，isolating-隔离中，isolated-已隔离，offlining-下线中

		Status *string `json:"Status,omitempty" name:"Status"`
		// 订阅实例的标签

		Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`
		// 订阅实例版本;txdts-旧版数据订阅，kafka-kafka版本数据订阅

		SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`
		// 订阅实例ChannelId

		Topic *string `json:"Topic,omitempty" name:"Topic"`
		// 订阅实例kafka的Broker信息

		Broker *string `json:"Broker,omitempty" name:"Broker"`
		// 订阅对象类型0-全实例订阅，1-DDL数据订阅，2-DML结构订阅，3-DDL数据订阅+DML结构订阅

		SubscribeObjectType *string `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`
		// 配置的订阅对象信息

		SubscribeObjects []*SubsObject `json:"SubscribeObjects,omitempty" name:"SubscribeObjects"`
		// 订阅错误信息

		Errors []*SubsErr `json:"Errors,omitempty" name:"Errors"`
		// kafka配置信息

		KafkaConfig *KafkaConfig `json:"KafkaConfig,omitempty" name:"KafkaConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaSubscribeConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaSubscribeConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateCheckJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeMigrateCheckJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateCheckJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopMigrateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMigrateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthenticateSubscribeSDKRequest struct {
	*tchttp.BaseRequest

	// 订阅Channel

	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
	// 订阅接口魔数

	MagicKey *string `json:"MagicKey,omitempty" name:"MagicKey"`
	// 订阅服务IP地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 订阅服务Port

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 订阅SDK版本号

	SdkVersion *string `json:"SdkVersion,omitempty" name:"SdkVersion"`
}

func (r *AuthenticateSubscribeSDKRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthenticateSubscribeSDKRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeConfRequest struct {
	*tchttp.BaseRequest

	// 订阅实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *DescribeSubscribeConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据订阅实例的ID数组

		SubscribeIds []*string `json:"SubscribeIds,omitempty" name:"SubscribeIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeRunningJobRequest struct {
	*tchttp.BaseRequest

	// 订阅实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *DescribeSubscribeRunningJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeRunningJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JobAttr struct {

	// 实例在后台数据库中的ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 归属用户Uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 当前用户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 任务ID,dts-开头短ID

	UJobId *string `json:"UJobId,omitempty" name:"UJobId"`
	// 任务可被用户识别别名

	Name *string `json:"Name,omitempty" name:"Name"`
	// immediate: 立即执行，timed: 定时执行

	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`
	// 定时任务设定定时启动时间

	ExpectRunTime *string `json:"ExpectRunTime,omitempty" name:"ExpectRunTime"`
	// 任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 启动时间

	StartedAt *string `json:"StartedAt,omitempty" name:"StartedAt"`
	// 变更时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 完成时间

	FinishedAt *string `json:"FinishedAt,omitempty" name:"FinishedAt"`
}

type CreateConnectTestJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 校验任务的Id，后续使用这些ID查询校验结果

		TaskIds []*int64 `json:"TaskIds,omitempty" name:"TaskIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConnectTestJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConnectTestJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsumerGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConsumerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncOption struct {

	// 同步对象，1-整个实例，2-指定库表

	SyncObject *uint64 `json:"SyncObject,omitempty" name:"SyncObject"`
	// 同步开始设置，1-立即开始

	RunMode *uint64 `json:"RunMode,omitempty" name:"RunMode"`
	// 同步模式， 3-全量且增量同步

	SyncType *uint64 `json:"SyncType,omitempty" name:"SyncType"`
	// 数据一致性检测， 1-无需配置

	ConsistencyType *uint64 `json:"ConsistencyType,omitempty" name:"ConsistencyType"`
}

type SyncStepDetailInfo struct {

	// 步骤编号

	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`
	// 步骤名

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 能否中止

	CanStop *int64 `json:"CanStop,omitempty" name:"CanStop"`
	// 步骤号

	StepId *int64 `json:"StepId,omitempty" name:"StepId"`
}

type DescribeJobStatusRequest struct {
	*tchttp.BaseRequest

	// 指定任务 ID 进行过滤

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 指定任务名称进行过滤

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 任务开始时间，时间格式：`yyyy-mm-dd hh:mm:ss`

	UpdateTimeStart *string `json:"UpdateTimeStart,omitempty" name:"UpdateTimeStart"`
	// 任务结束时间，时间格式：`yyyy-mm-dd hh:mm:ss`

	UpdateTimeEnd *string `json:"UpdateTimeEnd,omitempty" name:"UpdateTimeEnd"`
	// 任务状态，可选值有：1 - 创建中 - 2 - 创建完成 3 - 校验中4 - 校验通过5 - 校验不通过6 - 准备运行7 - 任务运行8 - 准备完成9 - 任务成功10 - 任务失败11 - 撤消中12 - 完成中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 偏移量，默认为 0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的记录数量，默认 20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeJobStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeDBTablesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 表数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 表对象

		Items []*string `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeDBTablesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeDBTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerGroupPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConsumerGroupPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerGroupPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscribeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {

	// 账户

	Account *string `json:"Account,omitempty" name:"Account"`
	// 消费组名

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 消费组位点

	ConsumerGroupOffset *int64 `json:"ConsumerGroupOffset,omitempty" name:"ConsumerGroupOffset"`
	// 消费组积压

	ConsumerGroupLag *int64 `json:"ConsumerGroupLag,omitempty" name:"ConsumerGroupLag"`
	// 消费组延迟

	Latency *int64 `json:"Latency,omitempty" name:"Latency"`
	// 分区状态

	StateOfPartition []*MonitorInfo `json:"StateOfPartition,omitempty" name:"StateOfPartition"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest

	// 任务 ID

	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 消费组信息

		Items []*GroupInfo `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumerGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeObjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务的ID

		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscribeObjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConnectTestResult struct {

	// 校验任务执行状态，状态有：starting运行中，finished失败

	Status *string `json:"Status,omitempty" name:"Status"`
	// 校验项是否全通过,0否，1是

	IsPass *int64 `json:"IsPass,omitempty" name:"IsPass"`
	// 代理IP，用户的防火墙需要允许该IP连接实例

	SNatIp *string `json:"SNatIp,omitempty" name:"SNatIp"`
	// 具体检查项,以及结果情况

	TestItems []*TestItems `json:"TestItems,omitempty" name:"TestItems"`
}

type SubscribeObject struct {

	// 数据订阅对象的类型，0-数据库，1-数据库内的表

	ObjectsType *int64 `json:"ObjectsType,omitempty" name:"ObjectsType"`
	// 订阅数据库的名称

	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	// 订阅数据库中表名称数组

	TableNames []*string `json:"TableNames,omitempty" name:"TableNames"`
}

type CreateConsumerGroupRequest struct {
	*tchttp.BaseRequest

	// 数据订阅任务ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 消费组名称

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
	// 用户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateConsumerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConsumerGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateCheckJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 校验任务状态：unavailable(当前不可用), starting(开始中)，running(校验中)，finished(校验完成)

		Status *string `json:"Status,omitempty" name:"Status"`
		// 任务的错误码

		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
		// 任务的错误信息

		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
		// Check任务总进度,如："30"表示30%

		Progress *string `json:"Progress,omitempty" name:"Progress"`
		// 校验是否通过,0-未通过，1-校验通过, 3-未校验

		CheckFlag *int64 `json:"CheckFlag,omitempty" name:"CheckFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateCheckJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxySupportRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Proxy 区域 列表信息

		Items []*RegionsOpt `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxySupportRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProxySupportRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合查询条件的总记录条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可迁移的实例列表

		Instances []*ExternParams `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOffsetByTimeRequest struct {
	*tchttp.BaseRequest

	// 订阅实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 时间字符串，格式如：2006-01-02 15:04:05

	Time *string `json:"Time,omitempty" name:"Time"`
}

func (r *DescribeOffsetByTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOffsetByTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeSubscribeRequest struct {
	*tchttp.BaseRequest

	// 订阅任务ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *ResumeSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeAutoRenewFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscribeAutoRenewFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步请求ID

		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateJobs4DevOpsRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 数据迁移任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 返回实例数量，默认20，有效区间[1,100]

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 数据迁移任务状态集合，可取值包括：任务状态,1-创建中(Creating),2-创建完成(Created),3-校验中(Checking)4-校验通过(CheckPass),5-校验不通过（CheckNotPass）,6-准备运行(ReadyRun),7-任务运行(Running),8-准备完成（ReadyComplete）,9-任务成功（Success）,10-任务失败（Failed）,11-中止中（Stoping）,12-完成中（Completing）,13-任务删除（Removed）

	Status []*int64 `json:"Status,omitempty" name:"Status"`
	// 源实例短ID，格式如：cdb-c1nl9rpv，与控制台页面中显示的实例ID相同 (**json里查询，效率可能比较低**)

	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`
	// 目标实例数据库类型（为异构迁移预留）,mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb

	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`
	// 目标实例短ID，格式如：cdb-c1nl9rpv，与控制台页面中显示的实例ID相同(**json里查询，效率可能比较低**)

	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMigrateJobs4DevOpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateJobs4DevOpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConnectTestJobRequest struct {
	*tchttp.BaseRequest

	// 数据库类型，如mysql,mariadb

	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
	// 数据库实例连接参数

	Endpoints []*EndpointItem `json:"Endpoints,omitempty" name:"Endpoints"`
}

func (r *CreateConnectTestJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConnectTestJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 单次返回的记录数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回记录的起始偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConsumerGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateSubscribeRequest struct {
	*tchttp.BaseRequest

	// 订阅实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *IsolateSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeConsumeTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscribeConsumeTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeConsumeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigureSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfigureSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigureSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConsumerGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConsumerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合查询条件的实例总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据订阅实例的信息列表

		Items []*SubscribeInfo `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedSubscribeRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *OfflineIsolatedSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OfflineIsolatedSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBEndpointInfo struct {

	// NA

	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`
	// NA

	Region *string `json:"Region,omitempty" name:"Region"`
	// NA

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// NA

	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
	// NA

	Info []*MigrateDBInfo `json:"Info,omitempty" name:"Info"`
	// NA

	Proxy []*MigrateProxyEndPoint `json:"Proxy,omitempty" name:"Proxy"`
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务执行结果信息

		Info *string `json:"Info,omitempty" name:"Info"`
		// 任务执行状态，可能的值有：success，failed，running

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateSurveyRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMigrateSurveyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateSurveyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsistencyParams struct {

	// 数据内容检测参数。表中选出用来数据对比的行，占表的总行数的百分比。取值范围是整数[1-100]

	SelectRowsPerTable *int64 `json:"SelectRowsPerTable,omitempty" name:"SelectRowsPerTable"`
	// 数据内容检测参数。迁移库表中，要进行数据内容检测的表，占所有表的百分比。取值范围是整数[1-100]

	TablesSelectAll *int64 `json:"TablesSelectAll,omitempty" name:"TablesSelectAll"`
	// 数据数量检测，检测表行数是否一致。迁移库表中，要进行数据数量检测的表，占所有表的百分比。取值范围是整数[1-100]

	TablesSelectCount *int64 `json:"TablesSelectCount,omitempty" name:"TablesSelectCount"`
}

type CreateMigrateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据迁移任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMigrateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StartMigrateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartMigrateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JobStatusInfo struct {

	// 迁移任务编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 迁移任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 迁移任务OwnerUin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 迁移任务Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 迁移任务AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 迁移任务状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type EndpointItem struct {

	// 网络类型：extranet,cvm,dcg,vpncloud,cdb,vpnselfbuild,intranet

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// 数据库提供商

	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`
	// AccessKey

	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`
	// Rds实例ID

	RdsInstance *string `json:"RdsInstance,omitempty" name:"RdsInstance"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// Cvm实例ID

	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`
	// 专线网络ID

	UniqDcgId *string `json:"UniqDcgId,omitempty" name:"UniqDcgId"`
	// 云vpn接入的实例ID

	UniqVpnGwId *string `json:"UniqVpnGwId,omitempty" name:"UniqVpnGwId"`
	// 云联网实例ID

	CcnGwId *string `json:"CcnGwId,omitempty" name:"CcnGwId"`
	// VPC实例ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC实例子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// CDB实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库实例连接地址

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 数据库实例连接端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 数据库实例连接用户名

	User *string `json:"User,omitempty" name:"User"`
	// 数据库实例连接用户密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 扩展参数对

	ExtraAttr []*KeyValuePairOption `json:"ExtraAttr,omitempty" name:"ExtraAttr"`
}

type MigrateDBInfo struct {

	// Supplier

	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`
	// DatabaseType

	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
	// AccessType

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// Region

	Region *string `json:"Region,omitempty" name:"Region"`
	// Host

	Host *string `json:"Host,omitempty" name:"Host"`
	// NA

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// NA

	User *string `json:"User,omitempty" name:"User"`
	// NA

	Password *string `json:"Password,omitempty" name:"Password"`
	// NA

	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`
	// NA

	VpnGwId *string `json:"VpnGwId,omitempty" name:"VpnGwId"`
	// NA

	DcgGwId *string `json:"DcgGwId,omitempty" name:"DcgGwId"`
	// NA

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// NA

	CcnGwId *string `json:"CcnGwId,omitempty" name:"CcnGwId"`
	// NA

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// NA

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// NA

	RdsInstanceId *string `json:"RdsInstanceId,omitempty" name:"RdsInstanceId"`
	// NA

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// NA

	ExtraAttr []*KeyValuePairOption `json:"ExtraAttr,omitempty" name:"ExtraAttr"`
}

type DescribeMigrateJobsRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 数据迁移任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 排序字段，可以取值为JobId、Status、JobName、MigrateType、RunMode、CreateTime

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序方式，升序为ASC，降序为DESC

	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回实例数量，默认20，有效区间[1,100]

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 数据迁移任务状态集合，可取值包括：任务状态,1-创建中(Creating),2-创建完成(Created),3-校验中(Checking)4-校验通过(CheckPass),5-校验不通过（CheckNotPass）,6-准备运行(ReadyRun),7-任务运行(Running),8-准备完成（ReadyComplete）,9-任务成功（Success）,10-任务失败（Failed）,11-中止中（Stoping）,12-完成中（Completing）,13-任务删除（Removed）

	Status []*int64 `json:"Status,omitempty" name:"Status"`
	// 源实例短ID，格式如：cdb-c1nl9rpv，与控制台页面中显示的实例ID相同 (**json里查询，效率可能比较低**)

	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`
	// 源地域英文名，如:ap-guangzhou, 和原来的数字regionId:1对应(**json里查询，效率可能比较低**)

	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`
	// 源实例数据库类型，mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb

	SrcDatabaseType []*string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`
	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)

	SrcAccessType []*string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`
	// 目标实例短ID，格式如：cdb-c1nl9rpv，与控制台页面中显示的实例ID相同(**json里查询，效率可能比较低**)

	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`
	// 目标地域英文名，如:ap-guangzhou, 和原来的数字regionId:1对应(**json里查询，效率可能比较低**)

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
	// 目标实例数据库类型（为异构迁移预留）,mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb

	DstDatabaseType []*string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`
	// 目标接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)

	DstAccessType []*string `json:"DstAccessType,omitempty" name:"DstAccessType"`
	// 任务运行模式，值包括：1-立即执行，2-定时执行

	RunMode *int64 `json:"RunMode,omitempty" name:"RunMode"`
	// 要查询的任务提交时间起始点，格式为 yyyy-mm-dd hh:mm:ss

	CreateTimeStart *string `json:"CreateTimeStart,omitempty" name:"CreateTimeStart"`
	// 要查询的任务提交时间结束点，格式为 yyyy-mm-dd hh:mm:ss

	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" name:"CreateTimeEnd"`
	// 根据 `目标实例信息` 模糊匹配搜索

	DstInfoPattern *string `json:"DstInfoPattern,omitempty" name:"DstInfoPattern"`
	// 根据 `源实例信息` 模糊匹配搜索

	SrcInfoPattern *string `json:"SrcInfoPattern,omitempty" name:"SrcInfoPattern"`
}

func (r *DescribeMigrateJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeDBDatabasesRequest struct {
	*tchttp.BaseRequest

	// 云数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云数据库类型，目前支持mysql

	Product *string `json:"Product,omitempty" name:"Product"`
	// 库正则过滤条件

	DatabaseRegexp *string `json:"DatabaseRegexp,omitempty" name:"DatabaseRegexp"`
	// 返回列表最大长度,默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页起始点，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 订阅实例id

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *DescribeSubscribeDBDatabasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeDBDatabasesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeReturnableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例是否支持退还/退货

		IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`
		// 不支持退还的原因

		ReturnFailMessage *string `json:"ReturnFailMessage,omitempty" name:"ReturnFailMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeReturnableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeReturnableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMigrateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateDetailInfo struct {

	// 总步骤数

	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`
	// 当前步骤

	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`
	// 总进度,如："10"

	Progress *string `json:"Progress,omitempty" name:"Progress"`
	// 当前步骤进度,如:"1"

	CurrentStepProgress *string `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`
	// 主从差距，MB；在增量同步阶段有效，目前支持产品为：redis和mysql

	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`
	// 主从差距，秒；在增量同步阶段有效，目前支持产品为：mysql

	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`
	// 步骤信息

	StepInfo []*MigrateStepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo"`
}

type MigrateProxyEndPoint struct {

	// 数据库类型:mariadb

	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
	// 数据库节点实例

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 路由信息在库表中的编号

	TrsfRouteInfoId *int64 `json:"TrsfRouteInfoId,omitempty" name:"TrsfRouteInfoId"`
	// 代理IP

	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
	// 代理Port

	ProxyPort *int64 `json:"ProxyPort,omitempty" name:"ProxyPort"`
	// nat ip

	SnatIp *string `json:"SnatIp,omitempty" name:"SnatIp"`
	// 代理路由所在区域

	ProxyRegion *string `json:"ProxyRegion,omitempty" name:"ProxyRegion"`
	// 任务Id

	ForUJobId *string `json:"ForUJobId,omitempty" name:"ForUJobId"`
	// User

	User *string `json:"User,omitempty" name:"User"`
	// Password

	Password *string `json:"Password,omitempty" name:"Password"`
	// 添加时间

	AddTimeStamp *string `json:"AddTimeStamp,omitempty" name:"AddTimeStamp"`
	// 是否是vpc支持的

	IsVpc *int64 `json:"IsVpc,omitempty" name:"IsVpc"`
	// 链接字段扩展信息

	Attr *string `json:"Attr,omitempty" name:"Attr"`
	// EndPointId 内部编号

	EndPointId *string `json:"EndPointId,omitempty" name:"EndPointId"`
	// NeedReturn

	NeedReturn *bool `json:"NeedReturn,omitempty" name:"NeedReturn"`
	// HasReturned

	HasReturned *bool `json:"HasReturned,omitempty" name:"HasReturned"`
	// ReturnAt

	ReturnAt *string `json:"ReturnAt,omitempty" name:"ReturnAt"`
	// 扩展字段

	ExtraAttr []*KeyValuePairOption `json:"ExtraAttr,omitempty" name:"ExtraAttr"`
}

type DescribeSubscribeCheckJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订阅实例ID

		SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
		// 异步任务当前状态信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 结果状态，取值：success,failed,running

		Status *string `json:"Status,omitempty" name:"Status"`
		// 进度百分比

		Progress *int64 `json:"Progress,omitempty" name:"Progress"`
		// 总步骤数

		StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`
		// 当前步骤

		StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`
		// 具体步骤详细信息

		Steps []*StepInfo `json:"Steps,omitempty" name:"Steps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeCheckJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeRunningJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订阅实例ID

		SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
		// 异步任务当前状态信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 结果状态，取值：success,failed,running

		Status *string `json:"Status,omitempty" name:"Status"`
		// 进度百分比

		Progress *int64 `json:"Progress,omitempty" name:"Progress"`
		// 总步骤数

		StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`
		// 当前步骤

		StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`
		// 具体步骤详细信息

		Steps []*StepInfo `json:"Steps,omitempty" name:"Steps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeRunningJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeRunningJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineIsolatedSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OfflineIsolatedSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartMigrateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMigrateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SupportSubscribeDBItem struct {

	// 数据库实例所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例版本

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 实例所在VPC

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 实例所在subnet

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// cdb实例的Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// cdb实例的Vport

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 实例是否可以作为被订阅对象

	Usable *bool `json:"Usable,omitempty" name:"Usable"`
	// 实例不能被作为订阅对象的原因

	Hint *string `json:"Hint,omitempty" name:"Hint"`
}

type KafkaConfig struct {

	// kafka分区数，只有 1，4，8，16几个选项

	NumberOfPartitions *int64 `json:"NumberOfPartitions,omitempty" name:"NumberOfPartitions"`
	// 分区规则

	DistributeRules []*DistributeRule `json:"DistributeRules,omitempty" name:"DistributeRules"`
}

type DescribeRdsSupportRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRdsSupportRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRdsSupportRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribesRequest struct {
	*tchttp.BaseRequest

	// 数据订阅的实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 数据订阅的实例名称

	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`
	// 绑定数据库实例的ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据订阅实例的通道ID

	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
	// 计费模式筛选，可能的值：0-包年包月，1-按量计费

	PayType *string `json:"PayType,omitempty" name:"PayType"`
	// 数据订阅实例的状态，creating - 创建中，normal - 正常运行，isolating - 隔离中，isolated - 已隔离，offlining - 下线中

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 数据订阅实例的配置状态，unconfigure - 未配置， configuring - 配置中，configured - 已配置

	SubsStatus []*string `json:"SubsStatus,omitempty" name:"SubsStatus"`
	// 返回记录的起始偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 单次返回的记录数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序方向，可选的值为"DESC"和"ASC"，默认为"DESC"，按创建时间逆序排序

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 标签过滤条件

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
	// 订阅实例版本;txdts-旧版数据订阅，kafka-kafka版本数据订阅

	SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`
	// Token信息

	QToken *string `json:"QToken,omitempty" name:"QToken"`
	// 客户端IP

	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`
}

func (r *DescribeSubscribesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerGroupDescriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConsumerGroupDescriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerGroupDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeNameRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 数据订阅实例的名称，长度限制为[1,60]

	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`
}

func (r *ModifySubscribeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetSubscribeRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *ResetSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OffsetTimeMap struct {

	// 时间字符串

	Time *string `json:"Time,omitempty" name:"Time"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分区编号

	PartitionNo *int64 `json:"PartitionNo,omitempty" name:"PartitionNo"`
}

type RegionsOpt struct {

	// 参数名称

	SupportType *string `json:"SupportType,omitempty" name:"SupportType"`
	// 区域列表

	SupportRegion []*string `json:"SupportRegion,omitempty" name:"SupportRegion"`
}

type DescribeJobStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志列表

		List []*JobStatusInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestItems struct {

	// 校验项，可能的值为Telnet,Connect

	TestName *string `json:"TestName,omitempty" name:"TestName"`
	// 错误码，其中-2表示未开始，-1表示测试未通过，0表示通过

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 校验结果信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type MonitorInfo struct {

	// 分区编号

	PartitionNo *int64 `json:"PartitionNo,omitempty" name:"PartitionNo"`
	// 消费组位点

	ConsumerGroupOffset *int64 `json:"ConsumerGroupOffset,omitempty" name:"ConsumerGroupOffset"`
	// 消费组积压

	ConsumerGroupLag *int64 `json:"ConsumerGroupLag,omitempty" name:"ConsumerGroupLag"`
	// 消费组延迟

	Latency *int64 `json:"Latency,omitempty" name:"Latency"`
}

type ModifySubscribeVipVportRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 指定目的子网，如果传此参数，DstIp必须在目的子网内

	DstUniqSubnetId *string `json:"DstUniqSubnetId,omitempty" name:"DstUniqSubnetId"`
	// 目标IP，与DstPort至少传一个

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 目标PORT，支持范围为：[1025-65535]

	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`
}

func (r *ModifySubscribeVipVportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeVipVportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerGroupPasswordRequest struct {
	*tchttp.BaseRequest

	// 订阅任务ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 消费组账户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 消费组名称

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
	// 旧的密码

	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`
	// 新密码

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
}

func (r *ModifyConsumerGroupPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerGroupPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscribeInfo struct {

	// 数据订阅的实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 数据订阅实例的名称

	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`
	// 数据订阅实例绑定的通道ID

	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
	// 数据订阅绑定实例对应的产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 数据订阅实例绑定的数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据订阅实例绑定的数据库实例状态

	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
	// 数据订阅实例的配置状态，unconfigure - 未配置， configuring - 配置中，configured - 已配置

	SubsStatus *string `json:"SubsStatus,omitempty" name:"SubsStatus"`
	// 上次修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 隔离时间

	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
	// 到期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 下线时间

	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
	// 最近一次修改的消费时间起点，如果从未修改则为零值

	ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`
	// 数据订阅实例所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 计费方式，0 - 包年包月，1 - 按量计费

	PayType *int64 `json:"PayType,omitempty" name:"PayType"`
	// 数据订阅实例的Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 数据订阅实例的Vport

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 数据订阅实例Vip所在VPC的唯一ID

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 数据订阅实例Vip所在子网的唯一ID

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// 数据订阅实例的状态，creating - 创建中，normal - 正常运行，isolating - 隔离中，isolated - 已隔离，offlining - 下线中，offline - 已下线

	Status *string `json:"Status,omitempty" name:"Status"`
	// SDK最后一条确认消息的时间戳，如果SDK一直消费，也可以作为SDK当前消费时间点

	SdkConsumedTime *string `json:"SdkConsumedTime,omitempty" name:"SdkConsumedTime"`
	// 标签

	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`
	// 自动续费标识。0-不自动续费，1-自动续费

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 订阅实例版本；txdts-旧版数据订阅,kafka-kafka版本数据订阅

	SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`
}

type DescribeConnectTestResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 连通性测试结果

		Items []*ConnectTestResult `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConnectTestResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectTestResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务详情数组

		JobList []*MigrateJobInfo `json:"JobList,omitempty" name:"JobList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可作为订阅的云数据库数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表

		Items []*SupportSubscribeDBItem `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeDBTablesRequest struct {
	*tchttp.BaseRequest

	// 云数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云数据库类型，目前支持mysql

	Product *string `json:"Product,omitempty" name:"Product"`
	// 库名

	Database *string `json:"Database,omitempty" name:"Database"`
	// 正则过滤条件

	TableRegexp *string `json:"TableRegexp,omitempty" name:"TableRegexp"`
	// 最大显示数量，默认为20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页起始点，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 订阅实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *DescribeSubscribeDBTablesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeDBTablesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步请求ID

		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateJobBasicItem struct {

	// 迁移任务基础配置信息

	BasicInfo *JobAttr `json:"BasicInfo,omitempty" name:"BasicInfo"`
	// 迁移任务源实例信息

	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`
	// 迁移任务目标实例信息

	DstInfo *DBEndpointInfo `json:"DstInfo,omitempty" name:"DstInfo"`
}

type ConfigureSubscribeRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 订阅模式，取值有：all，dml，ddl, dml,dmlAndDdl

	SubscribeMode *string `json:"SubscribeMode,omitempty" name:"SubscribeMode"`
	// 配置订阅对象

	SubscribeObjects []*SubsObject `json:"SubscribeObjects,omitempty" name:"SubscribeObjects"`
	// 接入类型，默认为cdb

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// 数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库实例账户

	DbAccount *string `json:"DbAccount,omitempty" name:"DbAccount"`
	// 数据库实例密码

	DbPassword *string `json:"DbPassword,omitempty" name:"DbPassword"`
	// single,cluster 单节点或者集群

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 数据库实例节点信息

	Endpoints []*Endpoint `json:"Endpoints,omitempty" name:"Endpoints"`
	// 订阅kafka配置

	KafkaConfig *KafkaConfig `json:"KafkaConfig,omitempty" name:"KafkaConfig"`
}

func (r *ConfigureSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigureSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncDetailInfo struct {

	// 总步骤数

	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`
	// 当前步骤

	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`
	// 总进度

	Progress *string `json:"Progress,omitempty" name:"Progress"`
	// 当前步骤进度

	CurrentStepProgress *string `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`
	// 主从差距，MB

	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`
	// 主从差距，秒

	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`
	// 步骤信息

	StepInfo []*SyncStepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo"`
}

type DescribeDBSupportRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBSupportRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSupportRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProxySupportRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProxySupportRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProxySupportRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncSrcDstInfo struct {

	// 地域英文名，如：ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例短Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CreateMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 迁移任务配置选项

	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`
	// 源实例数据库类型，目前支持：mysql，redis，mongodb，postgresql，mariadb，percona。不同地域数据库类型的具体支持情况，请参考控制台创建迁移页面。

	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`
	// 源实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例),cdb(腾讯云数据库实例),ccn(云联网实例)

	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`
	// 源实例信息，具体内容跟迁移任务类型相关

	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`
	// 目标实例数据库类型，目前支持：mysql，redis，mongodb，postgresql，mariadb，percona。不同地域数据库类型的具体支持情况，请参考控制台创建迁移页面。

	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`
	// 目标实例接入类型，目前支持：cdb（腾讯云数据库实例）

	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`
	// 目标实例信息

	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`
	// 需要迁移的源数据库表信息，用json格式的字符串描述。当MigrateOption.MigrateObject配置为2（指定库表迁移）时必填。
	// 对于database-table两级结构的数据库：
	// [{Database:db1,Table:[table1,table2]},{Database:db2}]
	// 对于database-schema-table三级结构：
	// [{Database:db1,Schema:s1
	// Table:[table1,table2]},{Database:db1,Schema:s2
	// Table:[table1,table2]},{Database:db2,Schema:s1
	// Table:[table1,table2]},{Database:db3},{Database:db4
	// Schema:s1}]

	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
	// 迁移后台版本:v1.0 v2,0

	MigrateEngineVersion *string `json:"MigrateEngineVersion,omitempty" name:"MigrateEngineVersion"`
	// redis 迁移，单节点传 "simple"，集群迁移传 "cluster“，其他情况统一传的 ""

	SrcNodeType *string `json:"SrcNodeType,omitempty" name:"SrcNodeType"`
	// SrcInfo数组，用户redis集群信息

	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitempty" name:"SrcInfoMulti"`
}

func (r *CreateMigrateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMigrateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeCheckJobRequest struct {
	*tchttp.BaseRequest

	// 数据订阅任务ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *CreateSubscribeCheckJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscribeCheckJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 待修改的数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 数据迁移任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 迁移任务配置选项

	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`
	// 源实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例),cdb(云上CDB实例)

	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`
	// 源实例信息，具体内容跟迁移任务类型相关

	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`
	// 目标实例接入类型，值包括：extranet(外网),cvm(CVM自建实例),dcg(专线接入的实例),vpncloud(云VPN接入的实例)，cdb(云上CDB实例). 目前只支持cdb.

	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`
	// 目标实例信息, 其中目标实例地域不允许修改.

	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`
	// 当选择'指定库表'迁移的时候, 需要设置待迁移的源数据库表信息,用符合json数组格式的字符串描述, 如下所例。
	//
	// 对于database-table两级结构的数据库：
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// 对于database-schema-table三级结构：
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	//
	// 如果是'整个实例'的迁移模式,不需设置该字段

	DatabaseInfo []*DatabaseItem `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
	// redis 迁移，单节点传 "simple"，集群迁移传 "cluster“，其他情况统一传的 ""

	SrcNodeType *string `json:"SrcNodeType,omitempty" name:"SrcNodeType"`
	// Redis集群版传入多个源节点信息

	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitempty" name:"SrcInfoMulti"`
	// 源数据库类型

	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`
	// 目标数据库类型

	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`
}

func (r *ModifyMigrateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMigrateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSubscribeRequest struct {
	*tchttp.BaseRequest

	// 订阅任务ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *StartSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StepTip struct {

	// 提示信息内容

	Message *string `json:"Message,omitempty" name:"Message"`
	// 帮助文档

	HelpDoc *string `json:"HelpDoc,omitempty" name:"HelpDoc"`
}

type MigrateStepDetailInfo struct {

	// 步骤序列

	StepNo *int64 `json:"StepNo,omitempty" name:"StepNo"`
	// 步骤展现名称

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 步骤英文标识

	StepId *string `json:"StepId,omitempty" name:"StepId"`
	// 步骤状态:0-默认值,1-成功,2-失败,3-执行中,4-未执行

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 当前步骤开始的时间，格式为"yyyy-mm-dd hh:mm:ss"，该字段不存在或者为空是无意义

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type ActivateSubscribeRequest struct {
	*tchttp.BaseRequest

	// 订阅实例ID。

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据订阅类型0-全实例订阅，1数据订阅，2结构订阅，3数据订阅与结构订阅

	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`
	// 订阅对象

	Objects *SubscribeObject `json:"Objects,omitempty" name:"Objects"`
	// 数据订阅服务所在子网。默认为数据库实例所在的子网内。

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// 订阅服务端口；默认为7507

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 订阅实例名称

	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`
	// 数据库账户

	DbAccount *string `json:"DbAccount,omitempty" name:"DbAccount"`
	// 数据库密码

	DbPassword *string `json:"DbPassword,omitempty" name:"DbPassword"`
}

func (r *ActivateSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrateNameRequest struct {
	*tchttp.BaseRequest

	// 迁移任务Id

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 新名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

func (r *ModifyMigrateNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMigrateNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyValuePairOption struct {

	// 扩展参数名称

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// 扩展参数内容

	ParameterValue *string `json:"ParameterValue,omitempty" name:"ParameterValue"`
}

type ActivateSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置数据订阅任务ID。

		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActivateSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateJobInfo struct {

	// 数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 数据迁移任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 迁移任务配置选项

	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`
	// 源实例数据库类型:mysql，redis，mongodb，postgresql，mariadb，percona

	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`
	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),cdb(腾讯云数据库实例),ccn(云联网实例)

	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`
	// 源实例信息，具体内容跟迁移任务类型相关

	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`
	// 目标实例数据库类型:mysql，redis，mongodb，postgresql，mariadb，percona

	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`
	// 目标实例接入类型，目前支持：cdb(腾讯云数据库实例)

	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`
	// 目标实例信息

	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`
	// 需要迁移的源数据库表信息，如果需要迁移的是整个实例，该字段为[]

	DatabaseInfo *MigrateObjectItem `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
	// 任务创建(提交)时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务开始执行时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务执行结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 任务状态,取值为：1-创建中(Creating),3-校验中(Checking)4-校验通过(CheckPass),5-校验不通过（CheckNotPass）,7-任务运行(Running),8-准备完成（ReadyComplete）,9-任务成功（Success）,10-任务失败（Failed）,11-撤销中（Stopping）,12-完成中（Completing）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务详情

	Detail *MigrateDetailInfo `json:"Detail,omitempty" name:"Detail"`
	// 任务错误信息提示，当任务发生错误时，不为null或者空值

	ErrorInfo []*ErrorInfo `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
}

type CreateSubscribeRequest struct {
	*tchttp.BaseRequest

	// 订阅的数据库类型，目前支持的有 mysql

	Product *string `json:"Product,omitempty" name:"Product"`
	// 实例付费类型，1小时计费，0包年包月

	PayType *int64 `json:"PayType,omitempty" name:"PayType"`
	// 购买时长。PayType为0时必填。单位为月，最大支持120

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// 购买数量,默认为1，最大为10

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 是否自动续费，默认为0，1表示自动续费。小时计费实例设置该标识无效。

	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
	// 实例资源标签

	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`
	// 订阅实例版本;txdts-旧版数据订阅，kafka-kafka版本数据订阅

	SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`
	// 数据订阅的实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteMigrateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMigrateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaSubscribeConfRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *DescribeKafkaSubscribeConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaSubscribeConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 订阅实例ID，例如：subs-8uey736k

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 自动续费标识。1-自动续费，0-不自动续费

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *ModifySubscribeAutoRenewFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeAutoRenewFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSupportRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DB 区域 列表信息

		Items []*RegionsOpt `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSupportRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSupportRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubsErr struct {

	// 订阅错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type SyncJobInfo struct {

	// 灾备任务id

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 灾备任务名

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 任务同步

	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`
	// 源接入类型

	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`
	// 源数据类型

	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`
	// 源实例信息

	SrcInfo *SyncInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`
	// 灾备接入类型

	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`
	// 灾备数据类型

	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`
	// 灾备实例信息

	DstInfo *SyncInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`
	// 任务信息

	Detail *SyncDetailInfo `json:"Detail,omitempty" name:"Detail"`
	// 任务状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 迁移库表

	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeMigrateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 数据库类型

	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
	// 分页偏移量 从0开始

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页容量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 迁移任务Role: "src", "dst"

	MigrateRole *string `json:"MigrateRole,omitempty" name:"MigrateRole"`
	// 数据引擎版本

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

func (r *DescribeMigrateDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateSurveyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// appId下总的迁移实例

		TotalInstances *int64 `json:"TotalInstances,omitempty" name:"TotalInstances"`
		// 分地域的迁移实例情况

		RegionInfos []*RegionInfo `json:"RegionInfos,omitempty" name:"RegionInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateSurveyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateSurveyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrateNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMigrateNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMigrateNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubsObject struct {

	// 订阅对象类型，0-database,1-table,默认database

	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
	// 数据库名

	Database *string `json:"Database,omitempty" name:"Database"`
	// 订阅表集合

	Tables []*string `json:"Tables,omitempty" name:"Tables"`
}

type DescribeSubscribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 云数据库类型，目前支持mysql

	Product *string `json:"Product,omitempty" name:"Product"`
	// 云数据库版本

	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`
	// 云数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 最大显示长度，默认为20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页起始点，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 订阅实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *DescribeSubscribeDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeObjectsRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 数据订阅的类型，可选的值有：0 - 全实例订阅；1 - 数据订阅；2 - 结构订阅；3 - 数据订阅+结构订阅

	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`
	// 订阅的数据库表信息

	Objects []*SubscribeObject `json:"Objects,omitempty" name:"Objects"`
	// 分区规则

	DistributeRules []*DistributeRule `json:"DistributeRules,omitempty" name:"DistributeRules"`
}

func (r *ModifySubscribeObjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeObjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DbInstances struct {

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例Vport

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// 实例是否可用，1为可用，0为不可用

	Usable *int64 `json:"Usable,omitempty" name:"Usable"`
	// 前端文本索引key，格式是：“msg-产品缩写-数字编号”，以mysql为例：“msg-mysql-1”

	Hint *string `json:"Hint,omitempty" name:"Hint"`
}

type SrcInfo struct {

	// 阿里云AccessKey。源库是阿里云RDS5.6适用

	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`
	// 实例的IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 实例的端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 实例的用户名

	User *string `json:"User,omitempty" name:"User"`
	// 实例的密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 阿里云RDS实例ID。源库是阿里云RDS5.6/5.6适用

	RdsInstanceId *string `json:"RdsInstanceId,omitempty" name:"RdsInstanceId"`
	// CVM实例短ID，格式如：ins-olgl39y8，与云服务器控制台页面显示的实例ID相同。如果是CVM自建实例，需要传递此字段

	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`
	// 专线网关ID，格式如：dcg-0rxtqqxb

	UniqDcgId *string `json:"UniqDcgId,omitempty" name:"UniqDcgId"`
	// 私有网络ID，格式如：vpc-92jblxto

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络下的子网ID，格式如：subnet-3paxmkdz

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// VPN网关ID，格式如：vpngw-9ghexg7q

	UniqVpnGwId *string `json:"UniqVpnGwId,omitempty" name:"UniqVpnGwId"`
	// 数据库实例ID，格式如：cdb-powiqx8q

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 地域英文名，如：ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 当实例为RDS实例时，填写为aliyun, 其他情况均填写others

	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`
	// 云联网ID，如：ccn-afp6kltc

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 数据库版本，当实例为RDS实例时才有效，格式如：5.6或者5.7，默认为5.6

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

type DescribeMigrateJobs4DevOpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务详情数组

		Items []*MigrateJobBasicItem `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateJobs4DevOpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateJobs4DevOpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeVipVportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscribeVipVportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscribeVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步请求ID

		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscribeRegionConf struct {

	// 地域名称，如广州

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地区标识，如ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域名称，如华南地区

	Area *string `json:"Area,omitempty" name:"Area"`
	// 是否为默认地域，0 - 不是，1 - 是的

	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`
	// 当前地域的售卖情况，1 - 正常， 2-灰度，3 - 停售

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeMigrateObjectRequest struct {
	*tchttp.BaseRequest

	// 迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// Object名字过滤

	ObjectFilter *string `json:"ObjectFilter,omitempty" name:"ObjectFilter"`
	// Object路径

	ObjectPath []*string `json:"ObjectPath,omitempty" name:"ObjectPath"`
	// 新增::类型，比如 database,table,view

	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
	// 兼容原有 root:dbName:tableName

	CurrentPath *string `json:"CurrentPath,omitempty" name:"CurrentPath"`
}

func (r *DescribeMigrateObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否是叶子对象

		IsLeaf *int64 `json:"IsLeaf,omitempty" name:"IsLeaf"`
		// 对象列表

		ObjectList []*string `json:"ObjectList,omitempty" name:"ObjectList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可售卖地域的数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可售卖地域详情

		Items []*SubscribeRegionConf `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeCheckJobRequest struct {
	*tchttp.BaseRequest

	// 订阅实例ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *DescribeSubscribeCheckJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeCheckJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseItem struct {

	// Database 名称

	Database *string `json:"Database,omitempty" name:"Database"`
	// Schema名称

	Schema *string `json:"Schema,omitempty" name:"Schema"`
	// Table名称 列表

	Table []*string `json:"Table,omitempty" name:"Table"`
	// View名称 列表

	View []*string `json:"View,omitempty" name:"View"`
}

type TagFilter struct {

	// 标签键值

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type RegionInfo struct {

	// 此地区总的迁移实例

	TotalInstances *int64 `json:"TotalInstances,omitempty" name:"TotalInstances"`
	// 正常实例

	NormalInstances *int64 `json:"NormalInstances,omitempty" name:"NormalInstances"`
	// 异常实例

	AbnormalInstances *int64 `json:"AbnormalInstances,omitempty" name:"AbnormalInstances"`
	// \u76ee\u6807\u5b9e\u4f8bId

	Region *string `json:"Region,omitempty" name:"Region"`
}

type StepInfo struct {

	// 步骤名称

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 步骤ID

	StepId *string `json:"StepId,omitempty" name:"StepId"`
	// 步骤编号

	StepNo *int64 `json:"StepNo,omitempty" name:"StepNo"`
	// 步骤状态：notStarted,running,failed,success

	Status *string `json:"Status,omitempty" name:"Status"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 当前进度百分比

	Percent *int64 `json:"Percent,omitempty" name:"Percent"`
	// 错误提示信息

	Errors []*StepTip `json:"Errors,omitempty" name:"Errors"`
	// 警告提示信息

	Warnings []*StepTip `json:"Warnings,omitempty" name:"Warnings"`
}

type DescribeSubscribeConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订阅实例ID

		SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
		// 订阅实例名称

		SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`
		// 订阅通道

		ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
		// 订阅数据库类型

		Product *string `json:"Product,omitempty" name:"Product"`
		// 被订阅的实例

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 被订阅的实例的状态，可能的值有running,offline,isolate

		InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
		// 订阅实例状态，可能的值有unconfigure-未配置，configuring-配置中，configured-已配置

		SubsStatus *string `json:"SubsStatus,omitempty" name:"SubsStatus"`
		// 订阅实例生命周期状态，可能的值有：normal-正常，isolating-隔离中，isolated-已隔离，offlining-下线中

		Status *string `json:"Status,omitempty" name:"Status"`
		// 订阅实例创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 订阅实例被隔离时间

		IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
		// 订阅实例到期时间

		ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
		// 订阅实例下线时间

		OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
		// 订阅实例消费时间起点。

		ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`
		// 订阅实例计费类型，1-小时计费，0-包年包月

		PayType *int64 `json:"PayType,omitempty" name:"PayType"`
		// 订阅通道Vip

		Vip *string `json:"Vip,omitempty" name:"Vip"`
		// 订阅通道Port

		Vport *int64 `json:"Vport,omitempty" name:"Vport"`
		// 订阅通道所在VpcId

		UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
		// 订阅通道所在SubnetId

		UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
		// 当前SDK消费时间位点

		SdkConsumedTime *string `json:"SdkConsumedTime,omitempty" name:"SdkConsumedTime"`
		// 订阅SDK IP地址

		SdkHost *string `json:"SdkHost,omitempty" name:"SdkHost"`
		// 订阅对象类型0-全实例订阅，1-DDL数据订阅，2-DML结构订阅，3-DDL数据订阅+DML结构订阅

		SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`
		// 订阅对象，当SubscribeObjectType 为0时，此字段为空数组

		SubscribeObjects []*SubscribeObject `json:"SubscribeObjects,omitempty" name:"SubscribeObjects"`
		// 修改时间

		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
		// 地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 订阅实例的标签

		Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`
		// 自动续费标识,0-不自动续费，1-自动续费

		AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorInfo struct {

	// 具体的报错日志, 包含错误码和错误信息

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// 报错对应的帮助文档Ur

	HelpDoc *string `json:"HelpDoc,omitempty" name:"HelpDoc"`
}

type DescribeOffsetByTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 偏移量信息

		Items []*OffsetTimeMap `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOffsetByTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOffsetByTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DistributeRule struct {

	// 规则类型table,pk,cols

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 分区模式

	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
	// 匹配DB名称

	DbPattern *string `json:"DbPattern,omitempty" name:"DbPattern"`
	// 匹配Table名称

	TablePattern *string `json:"TablePattern,omitempty" name:"TablePattern"`
	// 匹配列名称

	Columns []*string `json:"Columns,omitempty" name:"Columns"`
}

type TagItem struct {

	// 标签键值

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type AuthenticateSubscribeSDKResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 鉴权有效期时间

		ExpirePeriod *uint64 `json:"ExpirePeriod,omitempty" name:"ExpirePeriod"`
		// 加密存储的密钥

		Token *string `json:"Token,omitempty" name:"Token"`
		// 订阅服务IP地址

		VIp *string `json:"VIp,omitempty" name:"VIp"`
		// 订阅服务端口

		VPort *uint64 `json:"VPort,omitempty" name:"VPort"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuthenticateSubscribeSDKResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthenticateSubscribeSDKResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditLog struct {

	// 影响行数

	AffectRows *int64 `json:"AffectRows,omitempty" name:"AffectRows"`
	// SQL类型

	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 审计策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 数据库名称

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// SQL语句详情

	Sql *string `json:"Sql,omitempty" name:"Sql"`
	// 客户端地址

	Host *string `json:"Host,omitempty" name:"Host"`
	// 执行时间，单位为微秒

	ExecTime *int64 `json:"ExecTime,omitempty" name:"ExecTime"`
	// 账号名称

	User *string `json:"User,omitempty" name:"User"`
	// 表名称

	TableName *string `json:"TableName,omitempty" name:"TableName"`
	// 时间戳

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// 执行结果

	Result *string `json:"Result,omitempty" name:"Result"`
}

type DescribeSubscribeReturnableRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *DescribeSubscribeReturnableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeReturnableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncCheckStepInfo struct {

	// 步骤序列

	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`
	// 步骤展现名称

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 步骤执行结果代码

	StepCode *int64 `json:"StepCode,omitempty" name:"StepCode"`
	// 步骤执行结果提示

	StepMessage *string `json:"StepMessage,omitempty" name:"StepMessage"`
}

type DescribeRdsSupportRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RDS 区域 列表信息

		Items []*RegionsOpt `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRdsSupportRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRdsSupportRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeDBDatabasesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例库数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例库

		Items []*string `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeDBDatabasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeDBDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateObjectItem struct {

	// 库名

	Database *string `json:"Database,omitempty" name:"Database"`
	// 表信息

	Table []*string `json:"Table,omitempty" name:"Table"`
	// 视图信息

	View []*string `json:"View,omitempty" name:"View"`
}

type CompleteMigrateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompleteMigrateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CompleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeCheckJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscribeCheckJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscribeCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Endpoint struct {

	// 数据库类型

	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
	// CDB实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Cvm实例ID

	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`
	// 专线网络ID

	UniqDcgId *string `json:"UniqDcgId,omitempty" name:"UniqDcgId"`
	// 云vpn接入的实例ID

	UniqVpnGwId *string `json:"UniqVpnGwId,omitempty" name:"UniqVpnGwId"`
	// 云联网实例ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// VPC实例ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC实例子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 数据库实例连接地址

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 数据库实例连接端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 数据库实例连接用户名

	DbAccount *string `json:"DbAccount,omitempty" name:"DbAccount"`
	// 数据库实例连接用户密码

	DbPassword *string `json:"DbPassword,omitempty" name:"DbPassword"`
}

type CompleteMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CompleteMigrateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CompleteMigrateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateOption struct {

	// 任务运行模式，值包括：1-立即执行，2-定时执行

	RunMode *int64 `json:"RunMode,omitempty" name:"RunMode"`
	// 期望执行时间，当runMode=2时，该字段必填，时间格式：yyyy-mm-dd hh:mm:ss

	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`
	// 数据迁移类型，值包括：1-结构迁移,2-全量迁移,3-全量+增量迁移

	MigrateType *int64 `json:"MigrateType,omitempty" name:"MigrateType"`
	// 迁移对象，1-整个实例，2-指定库表

	MigrateObject *int64 `json:"MigrateObject,omitempty" name:"MigrateObject"`
	// 抽样数据一致性检测参数，1-未配置,2-全量检测,3-抽样检测, 4-仅校验不一致表,5-不检测

	ConsistencyType *int64 `json:"ConsistencyType,omitempty" name:"ConsistencyType"`
	// 是否用源库Root账户覆盖目标库，值包括：0-不覆盖，1-覆盖，选择库表或者结构迁移时应该为0

	IsOverrideRoot *int64 `json:"IsOverrideRoot,omitempty" name:"IsOverrideRoot"`
	// 不同数据库用到的额外参数.以JSON格式描述.
	// Redis可定义如下的参数:
	// {
	// 	"ClientOutputBufferHardLimit":512, 	从机缓冲区的硬性容量限制(MB)
	// 	"ClientOutputBufferSoftLimit":512, 	从机缓冲区的软性容量限制(MB)
	// 	"ClientOutputBufferPersistTime":60, 从机缓冲区的软性限制持续时间(秒)
	// 	"ReplBacklogSize":512, 	环形缓冲区容量限制(MB)
	// 	"ReplTimeout":120，		复制超时时间(秒)
	// }
	// MongoDB可定义如下的参数:
	// {
	// 	'SrcAuthDatabase':'admin',
	// 	'SrcAuthFlag': "1",
	// 	'SrcAuthMechanism':"SCRAM-SHA-1"
	// }
	// MySQL暂不支持额外参数设置。

	ExternParams *string `json:"ExternParams,omitempty" name:"ExternParams"`
	// 仅用于“抽样数据一致性检测”，ConsistencyType配置为抽样检测时，必选

	ConsistencyParams *ConsistencyParams `json:"ConsistencyParams,omitempty" name:"ConsistencyParams"`
}

type DeleteConsumerGroupRequest struct {
	*tchttp.BaseRequest

	// 数据订阅实例的ID

	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
	// 消费组账户名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 消费组名称

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
}

func (r *DeleteConsumerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsumerGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMigrateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionConfRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncSrcInfo struct {

	// 地域英文名，如：ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例短Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CreateMigrateCheckJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMigrateCheckJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopMigrateJobRequest struct {
	*tchttp.BaseRequest

	// 数据迁移任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *StopMigrateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopMigrateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
