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

package v20220217

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskGroup struct {

	// 任务动作ID

	TaskGroupId *int64 `json:"TaskGroupId,omitempty" name:"TaskGroupId"`
	// 分组标题

	TaskGroupTitle *string `json:"TaskGroupTitle,omitempty" name:"TaskGroupTitle"`
	// 分组描述

	TaskGroupDescription *string `json:"TaskGroupDescription,omitempty" name:"TaskGroupDescription"`
	// 任务分组顺序

	TaskGroupOrder *int64 `json:"TaskGroupOrder,omitempty" name:"TaskGroupOrder"`
	// 对象类型ID

	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
	// 任务分组创建时间

	TaskGroupCreateTime *string `json:"TaskGroupCreateTime,omitempty" name:"TaskGroupCreateTime"`
	// 任务分组更新时间

	TaskGroupUpdateTime *string `json:"TaskGroupUpdateTime,omitempty" name:"TaskGroupUpdateTime"`
	// 动作分组动作列表

	TaskGroupActions []*TaskGroupAction `json:"TaskGroupActions,omitempty" name:"TaskGroupActions"`
	// 实例列表

	TaskGroupInstanceList []*string `json:"TaskGroupInstanceList,omitempty" name:"TaskGroupInstanceList"`
	// 执行模式。1 --- 顺序执行，2 --- 阶段执行

	TaskGroupMode *int64 `json:"TaskGroupMode,omitempty" name:"TaskGroupMode"`
}

type TaskMonitor struct {

	// 监控指标ID

	TaskMonitorId *int64 `json:"TaskMonitorId,omitempty" name:"TaskMonitorId"`
	// 监控指标对象类型ID

	TaskMonitorObjectTypeId *int64 `json:"TaskMonitorObjectTypeId,omitempty" name:"TaskMonitorObjectTypeId"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 实例ID列表

	InstancesIds []*string `json:"InstancesIds,omitempty" name:"InstancesIds"`
	// 中文指标

	MetricChineseName *string `json:"MetricChineseName,omitempty" name:"MetricChineseName"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type DescribeActionLibraryListRequest struct {
	*tchttp.BaseRequest

	// 0-100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 默认值0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Keyword取值{"动作名称": "a_title", "描述": "a_desc", "动作类型": "a_type", "创建时间": "a_create_time"}

	Filters []*ActionFilter `json:"Filters,omitempty" name:"Filters"`
	// 对象类型ID

	ObjectType *uint64 `json:"ObjectType,omitempty" name:"ObjectType"`
}

func (r *DescribeActionLibraryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionLibraryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Task struct {

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务标题

	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`
	// 任务描述

	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
	// 自定义标签

	TaskTag *string `json:"TaskTag,omitempty" name:"TaskTag"`
	// 任务状态

	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// 状态类型: 0 -- 无状态，1 -- 成功，2-- 失败，3--终止

	TaskStatusType *int64 `json:"TaskStatusType,omitempty" name:"TaskStatusType"`
	// 保护策略

	TaskProtectStrategy *string `json:"TaskProtectStrategy,omitempty" name:"TaskProtectStrategy"`
	// 任务创建时间

	TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`
	// 任务更新时间

	TaskUpdateTime *string `json:"TaskUpdateTime,omitempty" name:"TaskUpdateTime"`
	// 任务动作组

	TaskGroups []*TaskGroup `json:"TaskGroups,omitempty" name:"TaskGroups"`
	// 开始时间

	TaskStartTime *string `json:"TaskStartTime,omitempty" name:"TaskStartTime"`
	// 结束时间

	TaskEndTime *string `json:"TaskEndTime,omitempty" name:"TaskEndTime"`
	// 是否符合预期。1：符合预期，2：不符合预期

	TaskExpect *int64 `json:"TaskExpect,omitempty" name:"TaskExpect"`
	// 演习记录

	TaskSummary *string `json:"TaskSummary,omitempty" name:"TaskSummary"`
	// 任务模式。1:手工执行，2:自动执行

	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
	// 自动暂停时长。单位分钟

	TaskPauseDuration *int64 `json:"TaskPauseDuration,omitempty" name:"TaskPauseDuration"`
	// 演练创建者Uin

	TaskOwnerUin *string `json:"TaskOwnerUin,omitempty" name:"TaskOwnerUin"`
	// 地域ID

	TaskRegionId *int64 `json:"TaskRegionId,omitempty" name:"TaskRegionId"`
	// 监控指标列表

	TaskMonitors []*TaskMonitor `json:"TaskMonitors,omitempty" name:"TaskMonitors"`
	// 保护策略

	TaskPolicy *DescribePolicy `json:"TaskPolicy,omitempty" name:"TaskPolicy"`
	// 标签列表

	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

type TaskLog struct {

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务标题

	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`
	// 任务开始时间

	TaskStartTime *string `json:"TaskStartTime,omitempty" name:"TaskStartTime"`
	// 任务结束时间

	TaskEndTime *string `json:"TaskEndTime,omitempty" name:"TaskEndTime"`
}

type DescribeObjectTypeSearchListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果集

		Results []*DescribeObjectSearchField `json:"Results,omitempty" name:"Results"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeObjectTypeSearchListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObjectTypeSearchListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteTaskInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteTaskInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteTaskInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNoticeIdRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNoticeIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNoticeIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomActionRequest struct {
	*tchttp.BaseRequest

	// 需要查询的自定义动作ID

	ActionId *uint64 `json:"ActionId,omitempty" name:"ActionId"`
}

func (r *DescribeCustomActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCustomActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改成功后的动作ID

		ActionId *uint64 `json:"ActionId,omitempty" name:"ActionId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCustomActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已经移除的动作ID

		ActionId *uint64 `json:"ActionId,omitempty" name:"ActionId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteTaskRequest struct {
	*tchttp.BaseRequest

	// 需要执行的任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExecuteTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicy struct {

	// 保护策略ID列表

	TaskPolicyIdList []*string `json:"TaskPolicyIdList,omitempty" name:"TaskPolicyIdList"`
	// 保护策略状态

	TaskPolicyStatus *string `json:"TaskPolicyStatus,omitempty" name:"TaskPolicyStatus"`
	// 策略规则

	TaskPolicyRule *string `json:"TaskPolicyRule,omitempty" name:"TaskPolicyRule"`
	// 护栏策略生效处理策略 1:顺序执行，2:暂停

	TaskPolicyDealType *int64 `json:"TaskPolicyDealType,omitempty" name:"TaskPolicyDealType"`
}

type DescribeObjectMetricsRequest struct {
	*tchttp.BaseRequest

	// 对象类型ID

	ObjectTypeId *uint64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
}

func (r *DescribeObjectMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObjectMetricsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogContentRequest struct {
	*tchttp.BaseRequest

	// 演练ID列表

	TaskIdList []*int64 `json:"TaskIdList,omitempty" name:"TaskIdList"`
}

func (r *DescribeTaskLogContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskLogContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskResultRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 总结

	Summary *string `json:"Summary,omitempty" name:"Summary"`
	// 是否符合预期。1--符合预期，2--不符合预期

	Expect *int64 `json:"Expect,omitempty" name:"Expect"`
}

func (r *ModifyTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionFieldConfigResult struct {

	// 动作ID

	ActionId *uint64 `json:"ActionId,omitempty" name:"ActionId"`
	// 动作名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 动作对应的栏位配置详情

	ConfigDetail []*ActionFieldConfigDetail `json:"ConfigDetail,omitempty" name:"ConfigDetail"`
}

type TaskLogContent struct {

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务标题

	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`
	// 任日志内容json结构化数据务开始时间

	LogContentLinesJson *string `json:"LogContentLinesJson,omitempty" name:"LogContentLinesJson"`
	// 日志内容

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
}

type DeleteCustomActionRequest struct {
	*tchttp.BaseRequest

	// 需要删除的动作ID

	ActionId *uint64 `json:"ActionId,omitempty" name:"ActionId"`
}

func (r *DeleteCustomActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateListRequest struct {
	*tchttp.BaseRequest

	// 演练名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 标签键

	Tag []*string `json:"Tag,omitempty" name:"Tag"`
	// 状态

	IsUsed *int64 `json:"IsUsed,omitempty" name:"IsUsed"`
	// 分页Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 标签对

	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeTemplateListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务信息

		Task *Task `json:"Task,omitempty" name:"Task"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateRequest struct {
	*tchttp.BaseRequest

	// 经验库id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTaskStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatisticsBaseResults struct {

	// 标签

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 数据

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type CreateTaskGroupAction struct {

	// 动作ID

	ActionId *int64 `json:"ActionId,omitempty" name:"ActionId"`
	// 动作顺序

	TaskGroupActionOrder *int64 `json:"TaskGroupActionOrder,omitempty" name:"TaskGroupActionOrder"`
	// 动作通用参数

	TaskGroupActionGeneralConfiguration *string `json:"TaskGroupActionGeneralConfiguration,omitempty" name:"TaskGroupActionGeneralConfiguration"`
	// 动作自定义参数

	TaskGroupActionCustomConfiguration *string `json:"TaskGroupActionCustomConfiguration,omitempty" name:"TaskGroupActionCustomConfiguration"`
	// RandomId

	TaskGroupActionRandomId *int64 `json:"TaskGroupActionRandomId,omitempty" name:"TaskGroupActionRandomId"`
	// RecoverId

	TaskGroupActionRecoverId *int64 `json:"TaskGroupActionRecoverId,omitempty" name:"TaskGroupActionRecoverId"`
	// ExecuteId

	TaskGroupActionExecuteId *int64 `json:"TaskGroupActionExecuteId,omitempty" name:"TaskGroupActionExecuteId"`
}

type CreateTemplatePolicy struct {

	// 保护策略ID列表

	TemplatePolicyIdList []*string `json:"TemplatePolicyIdList,omitempty" name:"TemplatePolicyIdList"`
	// 策略规则

	TemplatePolicyRule *string `json:"TemplatePolicyRule,omitempty" name:"TemplatePolicyRule"`
	// 护栏策略生效处理策略 1:顺序执行，2:暂停

	TemplatePolicyDealType *int64 `json:"TemplatePolicyDealType,omitempty" name:"TemplatePolicyDealType"`
}

type Region struct {

	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域中文名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域英文名

	RegionEnglishName *string `json:"RegionEnglishName,omitempty" name:"RegionEnglishName"`
	// 是否是海外。0:国内，1:海外

	RegionIsOversea *int64 `json:"RegionIsOversea,omitempty" name:"RegionIsOversea"`
	// 地域缩写

	RegionAbbreviation *string `json:"RegionAbbreviation,omitempty" name:"RegionAbbreviation"`
	// 地区

	RegionArea *string `json:"RegionArea,omitempty" name:"RegionArea"`
	// 对应地区的资源总数

	TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
}

type TagWithDescribe struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeRegionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		RegionList []*Region `json:"RegionList,omitempty" name:"RegionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateRequest struct {
	*tchttp.BaseRequest

	// 经验库

	Template *CreateTemplate `json:"Template,omitempty" name:"Template"`
}

func (r *CreateTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeObjectTypeListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeObjectTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObjectTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionLibraryListResult struct {

	// 动作名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 动作描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 动作类型。范围：["平台","自定义"]

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 动作风险描述

	RiskDesc *string `json:"RiskDesc,omitempty" name:"RiskDesc"`
	// 动作ID

	ActionId *uint64 `json:"ActionId,omitempty" name:"ActionId"`
	// 动作属性（ 1：故障  2：恢复）

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 关联的动作ID

	RelationActionId *uint64 `json:"RelationActionId,omitempty" name:"RelationActionId"`
	// 操作命令

	ActionCommand *string `json:"ActionCommand,omitempty" name:"ActionCommand"`
	// 动作类型( 0 -- tat   1 -- 云API）

	ActionCommandType *uint64 `json:"ActionCommandType,omitempty" name:"ActionCommandType"`
	// 自定义动作的参数，json string

	ActionContent *string `json:"ActionContent,omitempty" name:"ActionContent"`
}

type CreateTemplateGroupAction struct {

	// 动作ID

	ActionId *int64 `json:"ActionId,omitempty" name:"ActionId"`
	// 动作顺序

	TemplateGroupActionOrder *int64 `json:"TemplateGroupActionOrder,omitempty" name:"TemplateGroupActionOrder"`
	// 动作通用参数

	TemplateGroupActionGeneralConfiguration *string `json:"TemplateGroupActionGeneralConfiguration,omitempty" name:"TemplateGroupActionGeneralConfiguration"`
	// 动作自定义参数

	TemplateGroupActionCustomConfiguration *string `json:"TemplateGroupActionCustomConfiguration,omitempty" name:"TemplateGroupActionCustomConfiguration"`
	// RandomId

	TemplateGroupActionRandomId *int64 `json:"TemplateGroupActionRandomId,omitempty" name:"TemplateGroupActionRandomId"`
	// RecoverId

	TemplateGroupActionRecoverId *int64 `json:"TemplateGroupActionRecoverId,omitempty" name:"TemplateGroupActionRecoverId"`
	// ExecuteId

	TemplateGroupActionExecuteId *int64 `json:"TemplateGroupActionExecuteId,omitempty" name:"TemplateGroupActionExecuteId"`
}

type TaskGroupInstance struct {

	// 实例ID

	TaskGroupInstanceId *int64 `json:"TaskGroupInstanceId,omitempty" name:"TaskGroupInstanceId"`
	// 实例ID

	TaskGroupInstanceObjectId *string `json:"TaskGroupInstanceObjectId,omitempty" name:"TaskGroupInstanceObjectId"`
	// 实例动作执行状态

	TaskGroupInstanceStatus *int64 `json:"TaskGroupInstanceStatus,omitempty" name:"TaskGroupInstanceStatus"`
	// 实例动作执行日志

	TaskGroupInstanceExecuteLog *string `json:"TaskGroupInstanceExecuteLog,omitempty" name:"TaskGroupInstanceExecuteLog"`
	// 实例创建时间

	TaskGroupInstanceCreateTime *string `json:"TaskGroupInstanceCreateTime,omitempty" name:"TaskGroupInstanceCreateTime"`
	// 实例更新时间

	TaskGroupInstanceUpdateTime *string `json:"TaskGroupInstanceUpdateTime,omitempty" name:"TaskGroupInstanceUpdateTime"`
	// 状态类型: 0 -- 无状态，1 -- 成功，2-- 失败，3--终止，4--跳过

	TaskGroupInstanceStatusType *int64 `json:"TaskGroupInstanceStatusType,omitempty" name:"TaskGroupInstanceStatusType"`
	// 执行开始时间

	TaskGroupInstanceStartTime *string `json:"TaskGroupInstanceStartTime,omitempty" name:"TaskGroupInstanceStartTime"`
	// 执行结束时间

	TaskGroupInstanceEndTime *string `json:"TaskGroupInstanceEndTime,omitempty" name:"TaskGroupInstanceEndTime"`
}

type CheckResourceByRoleNameRequest struct {
	*tchttp.BaseRequest
}

func (r *CheckResourceByRoleNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckResourceByRoleNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 演练统计结果

		Result []*TaskStatistics `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 经验库列表

		TemplateList []*TemplateListItem `json:"TemplateList,omitempty" name:"TemplateList"`
		// 列表数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplateListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskMonitor struct {

	// 监控指标对象类型ID

	TaskMonitorObjectTypeId *int64 `json:"TaskMonitorObjectTypeId,omitempty" name:"TaskMonitorObjectTypeId"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 实例列表

	InstancesIds []*string `json:"InstancesIds,omitempty" name:"InstancesIds"`
	// 指标中文名称

	MetricChineseName *string `json:"MetricChineseName,omitempty" name:"MetricChineseName"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type ObjectTypeConfig struct {

	// 主键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 对象类型配置字段列表

	Fields []*ObjectTypeConfigFields `json:"Fields,omitempty" name:"Fields"`
}

type MetricObjectMeaning struct {

	// 指标英文解释

	En *string `json:"En,omitempty" name:"En"`
	// 指标中文解释

	Zh *string `json:"Zh,omitempty" name:"Zh"`
}

type DeleteTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest

	// 演练名称

	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`
	// 标签键

	TaskTag []*string `json:"TaskTag,omitempty" name:"TaskTag"`
	// 状态

	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// 开始时间

	TaskStartTime *string `json:"TaskStartTime,omitempty" name:"TaskStartTime"`
	// 结束时间

	TaskEndTime *string `json:"TaskEndTime,omitempty" name:"TaskEndTime"`
	// 分页Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 标签对

	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateRequest struct {
	*tchttp.BaseRequest

	// 经验库ID

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatisticsLineChartResults struct {

	// 时间轴的下标

	Date *string `json:"Date,omitempty" name:"Date"`
	// 时间内的数据标签内容数组

	Results []*DescribeTaskStatisticsBaseResults `json:"Results,omitempty" name:"Results"`
}

type ModifyTemplateIsUsedRequest struct {
	*tchttp.BaseRequest

	// 经验库ID

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 经验库使用状态: 1 -- 使用中，2 -- 停用

	IsUsed *int64 `json:"IsUsed,omitempty" name:"IsUsed"`
}

func (r *ModifyTemplateIsUsedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTemplateIsUsedRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 经验库id

		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDummyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1:空操作

		DummyId *string `json:"DummyId,omitempty" name:"DummyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDummyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDummyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeObjectTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		ObjectTypeList []*ObjectType `json:"ObjectTypeList,omitempty" name:"ObjectTypeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeObjectTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObjectTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskStatistics struct {

	// 状态

	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// 状态类型

	TaskStatusType *int64 `json:"TaskStatusType,omitempty" name:"TaskStatusType"`
	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type TemplateListItem struct {

	// 经验库ID

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 经验库标题

	TemplateTitle *string `json:"TemplateTitle,omitempty" name:"TemplateTitle"`
	// 经验库描述

	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`
	// 经验库标签

	TemplateTag *string `json:"TemplateTag,omitempty" name:"TemplateTag"`
	// 经验库状态。1 -- 使用中，2 -- 停用

	TemplateIsUsed *int64 `json:"TemplateIsUsed,omitempty" name:"TemplateIsUsed"`
	// 经验库创建时间

	TemplateCreateTime *string `json:"TemplateCreateTime,omitempty" name:"TemplateCreateTime"`
	// 经验库更新时间

	TemplateUpdateTime *string `json:"TemplateUpdateTime,omitempty" name:"TemplateUpdateTime"`
	// 经验库关联的任务数量

	TemplateUsedNum *int64 `json:"TemplateUsedNum,omitempty" name:"TemplateUsedNum"`
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionLibraryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果列表

		Results []*ActionLibraryListResult `json:"Results,omitempty" name:"Results"`
		// 符合记录条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeActionLibraryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionLibraryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatisticsOperateConditionRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
	// 结束时间

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeTaskStatisticsOperateConditionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStatisticsOperateConditionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTemplateRequest struct {
	*tchttp.BaseRequest

	// 经验库

	Template *EditTemplate `json:"Template,omitempty" name:"Template"`
}

func (r *EditTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 动作ID

		ActionId *uint64 `json:"ActionId,omitempty" name:"ActionId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest

	// 创建任务数据

	Task *CreateTask `json:"Task,omitempty" name:"Task"`
}

func (r *CreateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDummyRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDummyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDummyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatisticsOperateConditionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 演练次数统计

		TaskTimesStatistics []*DescribeTaskStatisticsLineChartResults `json:"TaskTimesStatistics,omitempty" name:"TaskTimesStatistics"`
		// 演练结果统计

		TaskResultStatistics []*DescribeTaskStatisticsLineChartResults `json:"TaskResultStatistics,omitempty" name:"TaskResultStatistics"`
		// 动作库使用统计

		ActionUseStatistics []*DescribeTaskStatisticsBaseResults `json:"ActionUseStatistics,omitempty" name:"ActionUseStatistics"`
		// 经验库使用统计

		ExperienceStatistics []*DescribeTaskStatisticsBaseResults `json:"ExperienceStatistics,omitempty" name:"ExperienceStatistics"`
		// 演练对象演练次数统计

		TaskObjectTimesStatistics []*DescribeTaskStatisticsLineChartResults `json:"TaskObjectTimesStatistics,omitempty" name:"TaskObjectTimesStatistics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatisticsOperateConditionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStatisticsOperateConditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCamIdentityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0: Cam已授权， 1：主账号CAM未授权， 2：子账号CAM未授权

		CamIdentity *uint64 `json:"CamIdentity,omitempty" name:"CamIdentity"`
		// CamIdentity对应的信息

		CamMessage *string `json:"CamMessage,omitempty" name:"CamMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCamIdentityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCamIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 经验库详情

		Template *Template `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateMonitor struct {

	// 监控指标ID

	MonitorId *int64 `json:"MonitorId,omitempty" name:"MonitorId"`
	// 监控指标对象类型ID

	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 中文指标

	MetricChineseName *string `json:"MetricChineseName,omitempty" name:"MetricChineseName"`
}

type DescribeTaskExecuteLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志数据

		LogMessage []*string `json:"LogMessage,omitempty" name:"LogMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskExecuteLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskExecuteLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogListRequest struct {
	*tchttp.BaseRequest

	// 分页limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 开始时间 ：2021-01-01 00:00:00

	TaskStartTime *string `json:"TaskStartTime,omitempty" name:"TaskStartTime"`
	// 结束时间：2021-11-08 23:59:59

	TaskEndTime *string `json:"TaskEndTime,omitempty" name:"TaskEndTime"`
	// 演练名称

	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`
	// 排列顺序：desc 或 asc, 默认desc

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 排列字段，默认t_start_time

	OrderByField *string `json:"OrderByField,omitempty" name:"OrderByField"`
}

func (r *DescribeTaskLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricInfomation struct {

	// 命名空间，每个云产品会有一个命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标使用的单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 指标使用的单位

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
	// 指标支持的统计周期，单位是秒，如60、300

	Period []*int64 `json:"Period,omitempty" name:"Period"`
	// 统计指标含义解释

	Meaning *MetricObjectMeaning `json:"Meaning,omitempty" name:"Meaning"`
	// 维度描述信息

	Dimensions []*MetricDimensionsDesc `json:"Dimensions,omitempty" name:"Dimensions"`
	// 统计周期内指标方式

	Periods []*MetricPeriods `json:"Periods,omitempty" name:"Periods"`
	// 指标的中文名称

	MetricZhName *string `json:"MetricZhName,omitempty" name:"MetricZhName"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTaskResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务状态类型

	StatusType *int64 `json:"StatusType,omitempty" name:"StatusType"`
}

func (r *ModifyTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionFieldConfigDetail struct {

	// 组件类型
	// 可选项如下：
	// input  文本框
	// textarea  多行文本框
	// number  数值输入框
	// select   选择器
	// cascader  级联选择器
	// radio  单选
	// time   时间选择

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件label

	Lable *string `json:"Lable,omitempty" name:"Lable"`
	// 组件唯一标识， 传回后端时的key

	Field *string `json:"Field,omitempty" name:"Field"`
	// 默认值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 支持配置项如下,可根据需要选择配置项，不需要配置是设置空{}：
	//
	// {
	//
	//   placeholder: string (占位符)
	//
	//   tooltip: string (提示信息)
	//
	//   reg: RegExp (对输入内容格式进行正则校验的规则)
	//
	//   max: number (对于输入框，限制最大输入字符数，对于数值输入框，设置上限)
	//
	//   min: number (对于数值输入框，设置下限)
	//
	//   step: number (设置数值输入框的步长，默认为1)
	//
	//   format: string (时间选择的格式，如YYYY-MM-DD表示年月日, YYYY-MM-DD HH:mm:ss 表示时分秒)
	//
	//   separator:  string[] (多行输入框的分隔符，不传或者为空时表示不分隔，直接返回用户输入的文本字符串)
	//
	//   multiple: boolean (是否多选,对选择器和级联选择器有效)
	//
	//   options:  选择器的选项【支持以下两种形式】
	//
	// 直接给定选项数组  { value: string; label: string }[]
	// 通过调接口获取选项                                                                                                                                       { api: string(接口地址),                                                                                                                                       params: string[] (接口参数,对应于参数配置的field，前端根据field对应的所有组件的输入值作为参数查询数据， 为空时在组件加载时直接请求数据)                                                                                                    }
	// }

	Config *string `json:"Config,omitempty" name:"Config"`
	// 是否必填 (0 -- 否   1-- 是)

	Required *uint64 `json:"Required,omitempty" name:"Required"`
	// compute配置依赖的其他field满足的条件时通过校验（如：三个表单项中必须至少有一个填写了）
	//
	// [fieldName,
	//
	// { config:  此项保留，等待后面具体场景细化  }
	//
	// ]

	Validate *string `json:"Validate,omitempty" name:"Validate"`
	// 是否可见

	Visible *string `json:"Visible,omitempty" name:"Visible"`
}

type TemplateGroupAction struct {

	// 经验库分组动作ID

	TemplateGroupActionId *int64 `json:"TemplateGroupActionId,omitempty" name:"TemplateGroupActionId"`
	// 动作ID

	ActionId *int64 `json:"ActionId,omitempty" name:"ActionId"`
	// 分组动作顺序

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// 分组动作通用配置

	GeneralConfiguration *string `json:"GeneralConfiguration,omitempty" name:"GeneralConfiguration"`
	// 分组动作自定义配置

	CustomConfiguration *string `json:"CustomConfiguration,omitempty" name:"CustomConfiguration"`
	// 动作分组创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 动作分组更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 动作名称

	ActionTitle *string `json:"ActionTitle,omitempty" name:"ActionTitle"`
	// 自身随机id

	RandomId *int64 `json:"RandomId,omitempty" name:"RandomId"`
	// 恢复动作id

	RecoverId *int64 `json:"RecoverId,omitempty" name:"RecoverId"`
	// 执行动作id

	ExecuteId *int64 `json:"ExecuteId,omitempty" name:"ExecuteId"`
	// 调用api类型，0:tat, 1:云api

	ActionApiType *int64 `json:"ActionApiType,omitempty" name:"ActionApiType"`
	// 1:故障，2:恢复

	ActionAttribute *int64 `json:"ActionAttribute,omitempty" name:"ActionAttribute"`
	// 动作类型：平台和自定义

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

type TemplatePolicy struct {

	// 保护策略ID列表

	TemplatePolicyIdList []*string `json:"TemplatePolicyIdList,omitempty" name:"TemplatePolicyIdList"`
	// 策略规则

	TemplatePolicyRule *string `json:"TemplatePolicyRule,omitempty" name:"TemplatePolicyRule"`
	// 护栏策略生效处理策略 1:顺序执行，2:暂停

	TemplatePolicyDealType *int64 `json:"TemplatePolicyDealType,omitempty" name:"TemplatePolicyDealType"`
}

type DescribeTaskLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表数量总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 演练列表

		TaskLogList []*TaskLog `json:"TaskLogList,omitempty" name:"TaskLogList"`
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

type DescribeCustomActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询到的自定义动作ID

		ActionId *uint64 `json:"ActionId,omitempty" name:"ActionId"`
		// 动作参数

		Content *Content `json:"Content,omitempty" name:"Content"`
		// 动作关联标签

		Tags []*TagWithCreate `json:"Tags,omitempty" name:"Tags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 经验库id

		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateIsUsedResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTemplateIsUsedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTemplateIsUsedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicy struct {

	// 保护策略ID列表

	TaskPolicyIdList []*string `json:"TaskPolicyIdList,omitempty" name:"TaskPolicyIdList"`
	// 策略规则

	TaskPolicyRule *string `json:"TaskPolicyRule,omitempty" name:"TaskPolicyRule"`
	// 护栏策略生效处理策略 1:顺序执行，2:暂停

	TaskPolicyDealType *int64 `json:"TaskPolicyDealType,omitempty" name:"TaskPolicyDealType"`
}

type DescribeObjectTypeSearchListRequest struct {
	*tchttp.BaseRequest

	// 对象类型ID

	ObjectTypeId *uint64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
}

func (r *DescribeObjectTypeSearchListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObjectTypeSearchListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCustomActionRequest struct {
	*tchttp.BaseRequest

	// 修改后的动作参数

	Content *Content `json:"Content,omitempty" name:"Content"`
	// 修改的动作ID

	ActionId *uint64 `json:"ActionId,omitempty" name:"ActionId"`
	// 标签数组

	Tags []*TagWithCreate `json:"Tags,omitempty" name:"Tags"`
}

func (r *UpdateCustomActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTemplate struct {

	// 经验库id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 标题

	TemplateTitle *string `json:"TemplateTitle,omitempty" name:"TemplateTitle"`
	// 描述

	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`
	// 标签

	TemplateTag *string `json:"TemplateTag,omitempty" name:"TemplateTag"`
	// 模式：1-手工/ 2 ---自动

	TemplateMode *int64 `json:"TemplateMode,omitempty" name:"TemplateMode"`
	// 地域ID

	TemplateRegionId *int64 `json:"TemplateRegionId,omitempty" name:"TemplateRegionId"`
	// 监控指标列表

	TemplateMonitors []*CreateTemplateMonitor `json:"TemplateMonitors,omitempty" name:"TemplateMonitors"`
	// 动作组列表

	TemplateGroups []*CreateTemplateGroup `json:"TemplateGroups,omitempty" name:"TemplateGroups"`
	// 保护策略

	TemplatePolicy *CreateTemplatePolicy `json:"TemplatePolicy,omitempty" name:"TemplatePolicy"`
	// 自动暂停时间，单位分钟

	TemplatePauseDuration *int64 `json:"TemplatePauseDuration,omitempty" name:"TemplatePauseDuration"`
	// 标签列表

	Tags []*TagWithCreate `json:"Tags,omitempty" name:"Tags"`
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建任务后的任务ID

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteTaskInstanceRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务动作ID

	TaskActionId *uint64 `json:"TaskActionId,omitempty" name:"TaskActionId"`
	// 任务动作实例ID

	TaskInstanceIds []*uint64 `json:"TaskInstanceIds,omitempty" name:"TaskInstanceIds"`
	// 是否操作整个任务

	IsOperateAll *bool `json:"IsOperateAll,omitempty" name:"IsOperateAll"`
	// 操作类型：（1--启动   2--执行  3--跳过  4--停止  5--重试）

	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`
	// 动作组ID

	TaskGroupId *uint64 `json:"TaskGroupId,omitempty" name:"TaskGroupId"`
}

func (r *ExecuteTaskInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteTaskInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionFilter struct {

	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 搜索内容值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type CreateTaskGroup struct {

	// 动作组名称

	TaskGroupTitle *string `json:"TaskGroupTitle,omitempty" name:"TaskGroupTitle"`
	// 动作组描述

	TaskGroupDescription *string `json:"TaskGroupDescription,omitempty" name:"TaskGroupDescription"`
	// 对象类型ID

	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
	// 实例ID列表

	TaskGroupInstances []*string `json:"TaskGroupInstances,omitempty" name:"TaskGroupInstances"`
	// 任务动作组动作列表

	TaskGroupActions []*CreateTaskGroupAction `json:"TaskGroupActions,omitempty" name:"TaskGroupActions"`
	// 执行模式。1 --- 顺序执行，2 --- 阶段执行

	TaskGroupMode *int64 `json:"TaskGroupMode,omitempty" name:"TaskGroupMode"`
}

type MetricDimensionsDesc struct {

	// 维度名数组

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type TemplateGroup struct {

	// 经验库动作ID

	TemplateGroupId *int64 `json:"TemplateGroupId,omitempty" name:"TemplateGroupId"`
	// 经验库动作分组动作列表

	TemplateGroupActions []*TemplateGroupAction `json:"TemplateGroupActions,omitempty" name:"TemplateGroupActions"`
	// 分组标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 分组描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 分组顺序

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// 执行模式。1 --- 顺序执行，2 --- 阶段执行

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 对象类型ID

	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
	// 分组创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 分组更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeTaskExecuteLogsRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 返回的内容行数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 日志起始的行数。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTaskExecuteLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskExecuteLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskGroupAction struct {

	// 任务分组动作ID

	TaskGroupActionId *int64 `json:"TaskGroupActionId,omitempty" name:"TaskGroupActionId"`
	// 任务分组动作实例列表

	TaskGroupInstances []*TaskGroupInstance `json:"TaskGroupInstances,omitempty" name:"TaskGroupInstances"`
	// 动作ID

	ActionId *int64 `json:"ActionId,omitempty" name:"ActionId"`
	// 分组动作顺序

	TaskGroupActionOrder *int64 `json:"TaskGroupActionOrder,omitempty" name:"TaskGroupActionOrder"`
	// 分组动作通用配置

	TaskGroupActionGeneralConfiguration *string `json:"TaskGroupActionGeneralConfiguration,omitempty" name:"TaskGroupActionGeneralConfiguration"`
	// 分组动作自定义配置

	TaskGroupActionCustomConfiguration *string `json:"TaskGroupActionCustomConfiguration,omitempty" name:"TaskGroupActionCustomConfiguration"`
	// 分组动作状态

	TaskGroupActionStatus *int64 `json:"TaskGroupActionStatus,omitempty" name:"TaskGroupActionStatus"`
	// 动作分组创建时间

	TaskGroupActionCreateTime *string `json:"TaskGroupActionCreateTime,omitempty" name:"TaskGroupActionCreateTime"`
	// 动作分组更新时间

	TaskGroupActionUpdateTime *string `json:"TaskGroupActionUpdateTime,omitempty" name:"TaskGroupActionUpdateTime"`
	// 动作名称

	ActionTitle *string `json:"ActionTitle,omitempty" name:"ActionTitle"`
	// 状态类型: 0 -- 无状态，1 -- 成功，2-- 失败，3--终止，4--跳过

	TaskGroupActionStatusType *int64 `json:"TaskGroupActionStatusType,omitempty" name:"TaskGroupActionStatusType"`
	// RandomId

	TaskGroupActionRandomId *int64 `json:"TaskGroupActionRandomId,omitempty" name:"TaskGroupActionRandomId"`
	// RecoverId

	TaskGroupActionRecoverId *int64 `json:"TaskGroupActionRecoverId,omitempty" name:"TaskGroupActionRecoverId"`
	// ExecuteId

	TaskGroupActionExecuteId *int64 `json:"TaskGroupActionExecuteId,omitempty" name:"TaskGroupActionExecuteId"`
	// 调用api类型，0:tat, 1:云api

	ActionApiType *int64 `json:"ActionApiType,omitempty" name:"ActionApiType"`
	// 1:故障，2:恢复

	ActionAttribute *int64 `json:"ActionAttribute,omitempty" name:"ActionAttribute"`
	// 动作类型：平台、自定义

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

type ExecuteTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricPeriods struct {

	// 周期

	Period *string `json:"Period,omitempty" name:"Period"`
	// 统计方式

	StatType []*string `json:"StatType,omitempty" name:"StatType"`
}

type CreateCustomActionRequest struct {
	*tchttp.BaseRequest

	// 注意对于自定义动作的接口动作参数都是Content, 而创建演练时给演练传的动作参数是Paramters

	Content *Content `json:"Content,omitempty" name:"Content"`
	// 标签列表

	Tags []*TagWithCreate `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateCustomActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Content struct {

	// 自定义动作名

	Title *string `json:"Title,omitempty" name:"Title"`
	// 命令类型，固定为shell

	CommandType *string `json:"CommandType,omitempty" name:"CommandType"`
	// 执行路径，默认值为/root

	WorkingDirectory *string `json:"WorkingDirectory,omitempty" name:"WorkingDirectory"`
	// 执行用户，默认值root

	Username *string `json:"Username,omitempty" name:"Username"`
	// 命令内容

	Command *string `json:"Command,omitempty" name:"Command"`
	// 是否使用参数

	EnableParameter *bool `json:"EnableParameter,omitempty" name:"EnableParameter"`
	// 命令参数，如果使用参数为True，则为必填

	CommandParameters *string `json:"CommandParameters,omitempty" name:"CommandParameters"`
	// 自定义动作描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 动作风险等级，下拉选择["低风险", "中风险", "高风险"]  必填且无默认值

	RiskDescription *string `json:"RiskDescription,omitempty" name:"RiskDescription"`
}

type CreateTemplateGroup struct {

	// 动作组名称

	TemplateGroupTitle *string `json:"TemplateGroupTitle,omitempty" name:"TemplateGroupTitle"`
	// 动作组描述

	TemplateGroupDescription *string `json:"TemplateGroupDescription,omitempty" name:"TemplateGroupDescription"`
	// 对象类型ID

	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
	// 执行模式。1 --- 顺序执行，2 --- 阶段执行

	TemplateGroupMode *int64 `json:"TemplateGroupMode,omitempty" name:"TemplateGroupMode"`
	// 动作组动作列表

	TemplateGroupActions []*CreateTemplateGroupAction `json:"TemplateGroupActions,omitempty" name:"TemplateGroupActions"`
}

type TaskListItem struct {

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务标题

	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`
	// 任务描述

	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
	// 任务标签

	TaskTag *string `json:"TaskTag,omitempty" name:"TaskTag"`
	// 任务状态

	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// 任务创建时间

	TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`
	// 任务更新时间

	TaskUpdateTime *string `json:"TaskUpdateTime,omitempty" name:"TaskUpdateTime"`
}

type DescribeRegionListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否查询地域下的任务总数

	IsQueryTotal *bool `json:"IsQueryTotal,omitempty" name:"IsQueryTotal"`
}

func (r *DescribeRegionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTask struct {

	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务标题

	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`
	// 任务描述

	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
	// 任务标签

	TaskTag *string `json:"TaskTag,omitempty" name:"TaskTag"`
	// 模式：1-手工/ 2 ---自动

	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
	// 地域ID

	TaskRegionId *int64 `json:"TaskRegionId,omitempty" name:"TaskRegionId"`
	// 任务动作组列表

	TaskGroups []*CreateTaskGroup `json:"TaskGroups,omitempty" name:"TaskGroups"`
	// 监控指标列表

	TaskMonitors []*CreateTaskMonitor `json:"TaskMonitors,omitempty" name:"TaskMonitors"`
	// 保护策略

	TaskPolicy *CreatePolicy `json:"TaskPolicy,omitempty" name:"TaskPolicy"`
	// 任务自动暂停时间

	TaskPauseDuration *int64 `json:"TaskPauseDuration,omitempty" name:"TaskPauseDuration"`
	// 标签列表

	Tags []*TagWithCreate `json:"Tags,omitempty" name:"Tags"`
}

type DescribeCamIdentityRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCamIdentityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCamIdentityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObjectType struct {

	// 对象类型ID

	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
	// 对象类型名称

	ObjectTypeTitle *string `json:"ObjectTypeTitle,omitempty" name:"ObjectTypeTitle"`
	// 对象类型第一级

	ObjectTypeLevelOne *string `json:"ObjectTypeLevelOne,omitempty" name:"ObjectTypeLevelOne"`
	// 对象类型参数

	ObjectTypeParams *ObjectTypeConfig `json:"ObjectTypeParams,omitempty" name:"ObjectTypeParams"`
}

type Template struct {

	// 经验库ID

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 经验库标题

	TemplateTitle *string `json:"TemplateTitle,omitempty" name:"TemplateTitle"`
	// 经验库描述

	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`
	// 自定义标签

	TemplateTag *string `json:"TemplateTag,omitempty" name:"TemplateTag"`
	// 使用状态。1 ---- 使用中，2 --- 停用

	TemplateIsUsed *int64 `json:"TemplateIsUsed,omitempty" name:"TemplateIsUsed"`
	// 经验库创建时间

	TemplateCreateTime *string `json:"TemplateCreateTime,omitempty" name:"TemplateCreateTime"`
	// 经验库更新时间

	TemplateUpdateTime *string `json:"TemplateUpdateTime,omitempty" name:"TemplateUpdateTime"`
	// 经验库模式。1:手工执行，2:自动执行

	TemplateMode *int64 `json:"TemplateMode,omitempty" name:"TemplateMode"`
	// 自动暂停时长。单位分钟

	TemplatePauseDuration *int64 `json:"TemplatePauseDuration,omitempty" name:"TemplatePauseDuration"`
	// 演练创建者Uin

	TemplateOwnerUin *string `json:"TemplateOwnerUin,omitempty" name:"TemplateOwnerUin"`
	// 地域ID

	TemplateRegionId *int64 `json:"TemplateRegionId,omitempty" name:"TemplateRegionId"`
	// 动作组

	TemplateGroups []*TemplateGroup `json:"TemplateGroups,omitempty" name:"TemplateGroups"`
	// 监控指标

	TemplateMonitors []*TemplateMonitor `json:"TemplateMonitors,omitempty" name:"TemplateMonitors"`
	// 护栏监控

	TemplatePolicy *TemplatePolicy `json:"TemplatePolicy,omitempty" name:"TemplatePolicy"`
	// 标签列表

	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

type CheckResourceByRoleNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务相关角色使用用户资源列表

		List []*string `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckResourceByRoleNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckResourceByRoleNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNoticeIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 通知模板notice-mdkevu10

		NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNoticeIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNoticeIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskList []*TaskListItem `json:"TaskList,omitempty" name:"TaskList"`
		// 列表数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTaskRequest struct {
	*tchttp.BaseRequest

	// 任务信息

	Task *EditTask `json:"Task,omitempty" name:"Task"`
}

func (r *EditTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionFieldConfigListRequest struct {
	*tchttp.BaseRequest

	// 动作ID列表

	ActionIds []*uint64 `json:"ActionIds,omitempty" name:"ActionIds"`
	// 对象类型ID

	ObjectTypeId *uint64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
}

func (r *DescribeActionFieldConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionFieldConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeObjectMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询得到的指标描述列表

		MetricList []*MetricInfomation `json:"MetricList,omitempty" name:"MetricList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeObjectMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObjectMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskLogContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志内容列表

		TaskLogContentList []*TaskLogContent `json:"TaskLogContentList,omitempty" name:"TaskLogContentList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskLogContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskLogContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTask struct {

	// 任务标题

	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`
	// 任务描述

	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
	// 任务标签

	TaskTag *string `json:"TaskTag,omitempty" name:"TaskTag"`
	// 模式：1-手工/ 2 ---自动

	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
	// 地域ID

	TaskRegionId *int64 `json:"TaskRegionId,omitempty" name:"TaskRegionId"`
	// 任务动作组列表

	TaskGroups []*CreateTaskGroup `json:"TaskGroups,omitempty" name:"TaskGroups"`
	// 监控指标列表

	TaskMonitors []*CreateTaskMonitor `json:"TaskMonitors,omitempty" name:"TaskMonitors"`
	// 演练自动暂停时间，单位分钟

	TaskPauseDuration *int64 `json:"TaskPauseDuration,omitempty" name:"TaskPauseDuration"`
	// 保护策略

	TaskPolicy *CreatePolicy `json:"TaskPolicy,omitempty" name:"TaskPolicy"`
	// 如果从经验库创建，需要传入经验库ID

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 标签列表

	Tags []*TagWithCreate `json:"Tags,omitempty" name:"Tags"`
}

type CreateTemplate struct {

	// 标题

	TemplateTitle *string `json:"TemplateTitle,omitempty" name:"TemplateTitle"`
	// 描述

	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`
	// 标签

	TemplateTag *string `json:"TemplateTag,omitempty" name:"TemplateTag"`
	// 模式：1-手工/ 2 ---自动

	TemplateMode *int64 `json:"TemplateMode,omitempty" name:"TemplateMode"`
	// 地域ID

	TemplateRegionId *int64 `json:"TemplateRegionId,omitempty" name:"TemplateRegionId"`
	// 自动暂停时间，单位分钟

	TemplatePauseDuration *int64 `json:"TemplatePauseDuration,omitempty" name:"TemplatePauseDuration"`
	// 监控指标列表

	TemplateMonitors []*CreateTemplateMonitor `json:"TemplateMonitors,omitempty" name:"TemplateMonitors"`
	// 动作组列表

	TemplateGroups []*CreateTemplateGroup `json:"TemplateGroups,omitempty" name:"TemplateGroups"`
	// 保护策略

	TemplatePolicy *CreateTemplatePolicy `json:"TemplatePolicy,omitempty" name:"TemplatePolicy"`
	// 标签列表

	Tags []*TagWithCreate `json:"Tags,omitempty" name:"Tags"`
}

type CreateTemplateMonitor struct {

	// 监控指标对象类型ID

	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标中文名称

	MetricChineseName *string `json:"MetricChineseName,omitempty" name:"MetricChineseName"`
}

type DescribeObjectSearchField struct {

	// 搜索键参数

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 搜索键类型（input 、textarea、 tag）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 搜索键的标签名

	Label *string `json:"Label,omitempty" name:"Label"`
	// 搜索键的前缀键

	PrefixKeyword *string `json:"PrefixKeyword,omitempty" name:"PrefixKeyword"`
	// 传参类型，目前只有 int 和 string

	ArrayType *string `json:"ArrayType,omitempty" name:"ArrayType"`
}

type ObjectTypeConfigFields struct {

	// instanceId

	Key *string `json:"Key,omitempty" name:"Key"`
	// 实例id

	Header *string `json:"Header,omitempty" name:"Header"`
}

type TagWithCreate struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeActionFieldConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 通用栏位配置列表

		Common []*ActionFieldConfigResult `json:"Common,omitempty" name:"Common"`
		// 动作栏位配置列表

		Results []*ActionFieldConfigResult `json:"Results,omitempty" name:"Results"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeActionFieldConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionFieldConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
