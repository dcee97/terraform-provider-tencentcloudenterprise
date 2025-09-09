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

package v20200721

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ResultCondition struct {

	// 命中的条件的ID

	ConditionId *int64 `json:"ConditionId,omitempty" name:"ConditionId"`
	// 命中此条件的实例的列表

	UnsafeInstances []*Instance `json:"UnsafeInstances,omitempty" name:"UnsafeInstances"`
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsRequest struct {
	*tchttp.BaseRequest

	// 返回数量限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件,支持name, id 过滤

	Filters *Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSubscriptionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreStrategiesRequest struct {
	*tchttp.BaseRequest

	// 增加`add`或者删除`delete`

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 操作的策略ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *ModifyIgnoreStrategiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreStrategiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResultStrategy struct {

	// 策略包含的条件列表

	Conditions []*ResultCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 策略扫描的实例的列表

	AllInstances []*Instance `json:"AllInstances,omitempty" name:"AllInstances"`
	// 策略的ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略执行情况: init 初始化 running 进行中 success 成功 fail失败

	Code *string `json:"Code,omitempty" name:"Code"`
	// 如果策略执行失败,此字段为测试的失败详情

	Message *string `json:"Message,omitempty" name:"Message"`
}

type ListAllIgnoreInstancesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListAllIgnoreInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAllIgnoreInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleStatusRequest struct {
	*tchttp.BaseRequest

	// 角色授权状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateRoleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscriptionTimeInfo struct {

	// 订阅周期，可为Daily（每天）或Weekly（每周）

	Period *string `json:"Period,omitempty" name:"Period"`
	// 仅在Period为Weekly时可用，取值为1-7，对应周一至周日

	DayOfWeek *int64 `json:"DayOfWeek,omitempty" name:"DayOfWeek"`
	// 订阅时间（24小时制），格式为“00:00"，只能选择小时整点时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type DescribeSubscriptionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订阅信息，包括订阅时间和订阅邮箱等信息

		SubscriptionInfos []*SubscriptionInfo `json:"SubscriptionInfos,omitempty" name:"SubscriptionInfos"`
		// 符合条件的数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStrategyIgnoredDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 该策略下被忽略的实例列表。

		IgnoredDetails []*IgnoredDetail `json:"IgnoredDetails,omitempty" name:"IgnoredDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStrategyIgnoredDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStrategyIgnoredDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStrategyUnsafeDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 该策略下风险实例列表。

		UnsafeDetails []*UnsafeDetail `json:"UnsafeDetails,omitempty" name:"UnsafeDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStrategyUnsafeDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStrategyUnsafeDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreInstanceRequest struct {
	*tchttp.BaseRequest

	// 被忽略的策略的ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 增加`add`或者删除`delete`

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 被忽略的实例的信息

	Instance *IgnoreInstance `json:"Instance,omitempty" name:"Instance"`
}

func (r *ModifyIgnoreInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例的ID,例如:ins-e4frg13

	Id *string `json:"Id,omitempty" name:"Id"`
	// 实例的地区,例如:ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例的附属信息,例如:cvm的名字

	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

type DescribeDownloadTaskRequest struct {
	*tchttp.BaseRequest

	// excel下载结果的唯一id，由DownloadReportFileAsync接口生成。

	ResultId *string `json:"ResultId,omitempty" name:"ResultId"`
}

func (r *DescribeDownloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStrategyIgnoresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 根据此配置，匹配忽略实例列表（Risks）对应字段，例如:
		// {"Response":{"RequestId":"111","RiskFieldsDesc":[{"Field":"InstanceId","FieldName":"ID","FieldType":"string","FieldDict":{}},{"Field":"InstanceName","FieldName":"名称","FieldType":"string","FieldDict":{}},{"Field":"InstanceState","FieldName":"状态","FieldType":"string","FieldDict":{"LAUNCH_FAILED":"创建失败","PENDING":"创建中","REBOOTING":"重启中","RUNNING":"运行中","SHUTDOWN":"停止待销毁","STARTING":"开机中","STOPPED":"关机","STOPPING":"关机中","TERMINATING":"销毁中"}},{"Field":"Zone","FieldName":"可用区","FieldType":"string","FieldDict":{}},{"Field":"PrivateIPAddresses","FieldName":"IP地址(内)","FieldType":"stringSlice","FieldDict":{}},{"Field":"PublicIPAddresses","FieldName":"IP地址(公)","FieldType":"stringSlice","FieldDict":{}},{"Field":"Region","FieldName":"地域","FieldType":"string","FieldDict":{}},{"Field":"Tags","FieldName":"标签","FieldType":"tags","FieldDict":{}}],"RiskTotalCount":3,"Risks":"[{\"InstanceId\":\"ins-xxx1\",\"InstanceName\":\"xxx1\",\"InstanceState\":\"RUNNING\",\"PrivateIPAddresses\":[\"1.17.64.2\"],\"PublicIPAddresses\":null,\"Region\":\"ap-shanghai\",\"Tags\":null,\"Zone\":\"ap-shanghai-2\"},{\"InstanceId\":\"ins-xxx2\",\"InstanceName\":\"xxx2\",\"InstanceState\":\"RUNNING\",\"PrivateIPAddresses\":[\"1.17.64.11\"],\"PublicIPAddresses\":null,\"Region\":\"ap-shanghai\",\"Tags\":null,\"Zone\":\"ap-shanghai-2\"}]","StrategyId":9}}

		RiskFieldsDesc []*RiskFieldsDesc `json:"RiskFieldsDesc,omitempty" name:"RiskFieldsDesc"`
		// 评估项ID

		StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
		// 忽略实例个数

		RiskTotalCount *uint64 `json:"RiskTotalCount,omitempty" name:"RiskTotalCount"`
		// 忽略实例详情列表，需要json decode

		Risks *string `json:"Risks,omitempty" name:"Risks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStrategyIgnoresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStrategyIgnoresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskSummaryRequest struct {
	*tchttp.BaseRequest

	// 任务的ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// filter

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTaskSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIgnoreInstancesRequest struct {
	*tchttp.BaseRequest

	// 策略的ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *ListIgnoreInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIgnoreInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateReportStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateReportStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateReportStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务的开始时间戳

		InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
		// 任务的结束时间戳,如果为0则表示还未执行完

		FinishTime *int64 `json:"FinishTime,omitempty" name:"FinishTime"`
		// 任务执行的结果

		Strategies []*ResultStrategy `json:"Strategies,omitempty" name:"Strategies"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStrategiesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeStrategiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStrategiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIgnoreStrategiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 忽略的策略的ID列表

		StrategyIds []*int64 `json:"StrategyIds,omitempty" name:"StrategyIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListIgnoreStrategiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIgnoreStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscriptionTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 模版名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 订阅的产品列表

	Products []*string `json:"Products,omitempty" name:"Products"`
	// 订阅的维度ID

	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds"`
	// 订阅的TAG列表

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifySubscriptionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscriptionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filters struct {

	// 过滤名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DeleteSubscriptionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubscriptionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubscriptionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品名

		Products []*string `json:"Products,omitempty" name:"Products"`
		// groupId和名称

		Groups *Groups `json:"Groups,omitempty" name:"Groups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigStrategy struct {

	// 策略的ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 所属于那个产品,例如:cvm,cos

	Product *string `json:"Product,omitempty" name:"Product"`
	// 策略的名字,例如:COS存储桶权限

	Name *string `json:"Name,omitempty" name:"Name"`
	// 策略的描述,例如:检查COS中具有开放访问权限或者任何已通过身份认证的CAM用户访问的存储桶

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 策略的修复方式,例如:使用"COS阻止公共访问"来控制

	Repair *string `json:"Repair,omitempty" name:"Repair"`
	// 策略出现问题时候的综述,例如:%d个存储桶授予全局访问权限

	Notice *string `json:"Notice,omitempty" name:"Notice"`
	// 策略被忽略时候的描述,例如:%d个存储桶被忽略

	Ignore *string `json:"Ignore,omitempty" name:"Ignore"`
	// 策略包含的条件列表

	Conditions []*ConfigCondition `json:"Conditions,omitempty" name:"Conditions"`
}

type DeleteGlobalIgnoreTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGlobalIgnoreTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGlobalIgnoreTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStrategyIgnoredDetailRequest struct {
	*tchttp.BaseRequest

	// 任务Id。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 策略Id。

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 过滤条件。

	Filter *string `json:"Filter,omitempty" name:"Filter"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTaskStrategyIgnoredDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStrategyIgnoredDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStrategie struct {

	// 评估项ID

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 评估项名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 评估项描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 忽略模版

	Ignore *string `json:"Ignore,omitempty" name:"Ignore"`
	// 警告模版

	Notice *string `json:"Notice,omitempty" name:"Notice"`
	// 评估项对应产品ID

	Product *string `json:"Product,omitempty" name:"Product"`
	// 评估项对应产品名称

	ProductDesc *string `json:"ProductDesc,omitempty" name:"ProductDesc"`
	// 评估项优化建议

	Repair *string `json:"Repair,omitempty" name:"Repair"`
	// 评估项类别ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 评估项类别名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 评估项风险列表

	Conditions []*DescribeStrategiesCondition `json:"Conditions,omitempty" name:"Conditions"`
}

type SubscriptionEmailInfo struct {

	// 订阅邮箱

	EmailAddrs []*string `json:"EmailAddrs,omitempty" name:"EmailAddrs"`
}

type DescribeTaskStrategyRisksRequest struct {
	*tchttp.BaseRequest

	// 评估项ID

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 限制数量,默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 任务ID

	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
	// 过滤条件，支持:
	// 1. level 等级，2：中风险，3：高风险
	// 2. fuzzy 模糊搜索，用作ID或名称过滤

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 是否包含所有字段

	IncludeAllFields *bool `json:"IncludeAllFields,omitempty" name:"IncludeAllFields"`
}

func (r *DescribeTaskStrategyRisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStrategyRisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLastTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLastTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLastTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadReportFileAsyncResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 根据上述的参数，返回的唯一id，用于后续excel下载。

		ResultId *string `json:"ResultId,omitempty" name:"ResultId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadReportFileAsyncResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadReportFileAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGlobalIgnoreTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGlobalIgnoreTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGlobalIgnoreTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIgnoreInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateReportStatusRequest struct {
	*tchttp.BaseRequest

	// 修改的报告授权状态，可选值为0（不授权），1（授权）

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateReportStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateReportStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGlobalIgnoreTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// tag总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// tag结果

		Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGlobalIgnoreTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalIgnoreTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略扫描执行的结果（只包含未忽略的策略信息）。注意：此字段可能返回 null，表示取不到有效值。

		StrategySummaries []*StrategySummary `json:"StrategySummaries,omitempty" name:"StrategySummaries"`
		// 分组扫描执行的结果。注意：此字段可能返回 null，表示取不到有效值。

		GroupSummaries []*GroupSummary `json:"GroupSummaries,omitempty" name:"GroupSummaries"`
		// 任务的开始时间，按照ISO8601标准表示，格式为"YYYY-MM-DDThh:mm:ssZ"。任务尚未结束，返回"1970-01-01T00:00:00+00:00"。

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 任务的结束时间，按照ISO8601标准表示，格式为"YYYY-MM-DDThh:mm:ssZ"。为"1970-01-01T00:00:00+00:00"时表示任务尚未结束。

		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Groups struct {

	// 类型Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 类型名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type IgnoreInstance struct {

	// 实例的ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 实例的地区

	Region *string `json:"Region,omitempty" name:"Region"`
}

type RiskDay struct {

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type DescribeSubscriptionTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// template set

		TplSets []*TplSets `json:"TplSets,omitempty" name:"TplSets"`
		// template count

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportLanguageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持的语言列表。

		Languages []*string `json:"Languages,omitempty" name:"Languages"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSupportLanguageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportLanguageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskFieldsDesc struct {

	// 字段ID

	Field *string `json:"Field,omitempty" name:"Field"`
	// 字段名称

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 字段类型,
	// string: 字符串类型，例如"aa"
	// int: 整形，例如 111
	// stringSlice : 字符串数组类型，例如["a", "b"]
	// tags: 标签类型, 例如: [{"Key":"kkk","Value":"vvv"},{"Key":"kkk2","Value":"vvv2"}]

	FieldType *string `json:"FieldType,omitempty" name:"FieldType"`
	// 字段值对应字典

	FieldDict []*KeyValue `json:"FieldDict,omitempty" name:"FieldDict"`
}

type DeleteSubscriptionTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版ID列表

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteSubscriptionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubscriptionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 创建或修改的订阅信息

	SubscriptionInfos []*SubscriptionInfo `json:"SubscriptionInfos,omitempty" name:"SubscriptionInfos"`
}

func (r *UpdateSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRoleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadReportFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 评估结果文件名

		Name *string `json:"Name,omitempty" name:"Name"`
		// 经过base64编码的评估结果

		Content *string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadReportFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadReportFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrategySummary struct {

	// 策略Id

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略的扫描状态。策略执行情况: init 初始化 running 进行中 success 成功 fail失败。

	Code *string `json:"Code,omitempty" name:"Code"`
	// 策略对应的分类Id。

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 该策略下高风险的实例数量(不包括忽略实例数量)。

	HighRiskCount *int64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
	// 该策略下中风险的实例数量(不包括忽略实例数量)

	MediumRiskCount *int64 `json:"MediumRiskCount,omitempty" name:"MediumRiskCount"`
	// 该策略下低风险的实例数量(不包括忽略实例数量)。

	LowRiskCount *int64 `json:"LowRiskCount,omitempty" name:"LowRiskCount"`
	// 该策略下无风险的实例数量(不包括忽略实例数量)。

	NoRiskCount *int64 `json:"NoRiskCount,omitempty" name:"NoRiskCount"`
	// 该策略下被忽略的实例数量。

	IgnoredInstanceCount *int64 `json:"IgnoredInstanceCount,omitempty" name:"IgnoredInstanceCount"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 高风险+中风险+无风险

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type DeleteSubscriptionRequest struct {
	*tchttp.BaseRequest

	// subscriptionId

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGlobalIgnoreTagsRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 过滤参数：
	//
	// 支持的Name：
	//
	// TagKey
	// TagValues

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 1:taskID对应的镜像配置，taskID不为空，则查询对应的taskID下的镜像的tag全局忽略配置，此时Filters无效。
	//
	// 0: 原始的配置,此时Filters生效。

	Choose *int64 `json:"Choose,omitempty" name:"Choose"`
}

func (r *DescribeGlobalIgnoreTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalIgnoreTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLastTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最新任务的ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 是否存在

		Has *bool `json:"Has,omitempty" name:"Has"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLastTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLastTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest

	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIgnoreStrategiesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListIgnoreStrategiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIgnoreStrategiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品列表

	Products []*string `json:"Products,omitempty" name:"Products"`
	// group id

	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds"`
	// tag

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateSubscriptionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscriptionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscriptionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscriptionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDisplayRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRiskDisplayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDisplayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscriptionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscriptionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubscriptionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStrategiesCondition struct {

	// 警告条件ID

	ConditionId *uint64 `json:"ConditionId,omitempty" name:"ConditionId"`
	// 警告级别，2:中风险，3:高风险

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 警告级别描述

	LevelDesc *string `json:"LevelDesc,omitempty" name:"LevelDesc"`
	// 警告条件描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type TplSets struct {

	// group ids

	GroupIds []*uint64 `json:"GroupIds,omitempty" name:"GroupIds"`
	// template id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// template products

	Products []*string `json:"Products,omitempty" name:"Products"`
	// template name

	Name *string `json:"Name,omitempty" name:"Name"`
	// template tag

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

type DescribeStrategiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 评估项列表

		Strategies []*DescribeStrategie `json:"Strategies,omitempty" name:"Strategies"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStrategiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest

	// 任务的ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreStrategiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIgnoreStrategiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务的ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGlobalIgnoreTagsRequest struct {
	*tchttp.BaseRequest

	// 需要删除的tag key的值

	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

func (r *DeleteGlobalIgnoreTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGlobalIgnoreTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAllIgnoreInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 忽略的列表

		Ignores []*IgnoreStrategyInstance `json:"Ignores,omitempty" name:"Ignores"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAllIgnoreInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAllIgnoreInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IgnoreStrategyInstance struct {

	// 实例的ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 实例的地区编号 eg:ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 忽略的策略的ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

type Progress struct {

	// 已扫描实例数量。

	ScannedCount *int64 `json:"ScannedCount,omitempty" name:"ScannedCount"`
	// 总实例数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type SubscriptionInfo struct {

	// 订阅消息ID，-1表示新建，其余非负值表示将对应消息更新

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 消息订阅状态，false表示取消订阅，true表示订阅

	IsSubscribed *bool `json:"IsSubscribed,omitempty" name:"IsSubscribed"`
	// 订阅时间信息

	TimeInfo *SubscriptionTimeInfo `json:"TimeInfo,omitempty" name:"TimeInfo"`
	// 订阅邮箱信息

	EmailInfo *SubscriptionEmailInfo `json:"EmailInfo,omitempty" name:"EmailInfo"`
	// 订阅信息使用的语言。目前支持 zh-CN : 中文简体，en-US : 英文。默认值为 zh-CN

	Language *string `json:"Language,omitempty" name:"Language"`
	// 订阅名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 对应的订阅模版ID

	TplId *uint64 `json:"TplId,omitempty" name:"TplId"`
	// 对应的订阅模版名称

	TplName *string `json:"TplName,omitempty" name:"TplName"`
}

type DescribeRoleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色授权状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// CAM对应角色是否存在，0：不存在，1：存在

		RoleExist *int64 `json:"RoleExist,omitempty" name:"RoleExist"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyValue struct {

	// 键名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 键名对应值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type RegionCode struct {

	// 地区 , 如: ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地区的编码, 如:1

	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type RiskFieldsDescDetail struct {

	// 巡检项ID

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 巡检项实例对应的URL

	Url *string `json:"Url,omitempty" name:"Url"`
	// 巡检项对应配置

	RiskFieldsDesc []*RiskFieldsDesc `json:"RiskFieldsDesc,omitempty" name:"RiskFieldsDesc"`
	// 链接渲染字段ID

	LinkId *string `json:"LinkId,omitempty" name:"LinkId"`
	// 设置忽略字段ID

	PriId *string `json:"PriId,omitempty" name:"PriId"`
}

type Tags struct {

	// Tag对应的key

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// tag对应的values

	TagValues []*string `json:"TagValues,omitempty" name:"TagValues"`
}

type DescribeDownloadTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储扫描结果的Cos链接。

		CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`
		// 生成Excel任务执行情况: init：初始化，running：进行中，success：成功，fail：失败。

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductSummary struct {

	// 该分组下高风险的策略数量(不包括忽略策略数量)。

	HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
	// 该分组下中风险的策略数量(不包括忽略策略数量)。

	MediumRiskCount *uint64 `json:"MediumRiskCount,omitempty" name:"MediumRiskCount"`
	// 该分组下无风险的策略数量(不包括忽略策略数量)。

	NoRiskCount *uint64 `json:"NoRiskCount,omitempty" name:"NoRiskCount"`
	// 该分组下（高风险+中风险+无风险）的策略数量(不包括忽略策略数量)。

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 风险率=当前风险数/(资源数*已开启资源型巡检项)

	RiskRate *float64 `json:"RiskRate,omitempty" name:"RiskRate"`
	// 巡检项数量

	StrategyCount *uint64 `json:"StrategyCount,omitempty" name:"StrategyCount"`
	// 该分组下（高风险+中风险）的策略数量(不包括忽略策略数量)。

	RiskCount *uint64 `json:"RiskCount,omitempty" name:"RiskCount"`
	// 产品简称

	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略组的配置列表

		Groups []*ConfigGroup `json:"Groups,omitempty" name:"Groups"`
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

type DescribeOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务开始时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 任务结束时间

		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
		// 是否有成功的任务

		Has *bool `json:"Has,omitempty" name:"Has"`
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 按组划分数据

		GroupSummaries []*GroupSummary `json:"GroupSummaries,omitempty" name:"GroupSummaries"`
		// 按产品划分数据

		ProductSummaries []*ProductSummary `json:"ProductSummaries,omitempty" name:"ProductSummaries"`
		// 风险趋势数据

		RiskDays []*RiskDay `json:"RiskDays,omitempty" name:"RiskDays"`
		// TOPN 数据

		StrategyTopSummaries []*StrategySummary `json:"StrategyTopSummaries,omitempty" name:"StrategyTopSummaries"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 报告授权状态，0（不授权），1（授权）

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReportStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// tag总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// tag结果

		Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStrategyIgnoresRequest struct {
	*tchttp.BaseRequest

	// 评估项ID

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 限制数量,默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 任务ID

	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
	// 过滤条件，支持:
	// 1. level 等级，2：中风险，3：高风险
	// 2. fuzzy 模糊搜索，用作ID或名称过滤

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 是否包含所有字段

	IncludeAllFields *bool `json:"IncludeAllFields,omitempty" name:"IncludeAllFields"`
}

func (r *DescribeTaskStrategyIgnoresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStrategyIgnoresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStrategyRisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 根据此配置，匹配风险实例列表（Risks）对应字段，例如:
		// {"Response":{"RequestId":"111","RiskFieldsDesc":[{"Field":"InstanceId","FieldName":"ID","FieldType":"string","FieldDict":{}},{"Field":"InstanceName","FieldName":"名称","FieldType":"string","FieldDict":{}},{"Field":"InstanceState","FieldName":"状态","FieldType":"string","FieldDict":{"LAUNCH_FAILED":"创建失败","PENDING":"创建中","REBOOTING":"重启中","RUNNING":"运行中","SHUTDOWN":"停止待销毁","STARTING":"开机中","STOPPED":"关机","STOPPING":"关机中","TERMINATING":"销毁中"}},{"Field":"Zone","FieldName":"可用区","FieldType":"string","FieldDict":{}},{"Field":"PrivateIPAddresses","FieldName":"IP地址(内)","FieldType":"stringSlice","FieldDict":{}},{"Field":"PublicIPAddresses","FieldName":"IP地址(公)","FieldType":"stringSlice","FieldDict":{}},{"Field":"Region","FieldName":"地域","FieldType":"string","FieldDict":{}},{"Field":"Tags","FieldName":"标签","FieldType":"tags","FieldDict":{}}],"RiskTotalCount":3,"Risks":"[{\"InstanceId\":\"ins-xxx1\",\"InstanceName\":\"xxx1\",\"InstanceState\":\"RUNNING\",\"PrivateIPAddresses\":[\"1.17.64.2\"],\"PublicIPAddresses\":null,\"Region\":\"ap-shanghai\",\"Tags\":null,\"Zone\":\"ap-shanghai-2\"},{\"InstanceId\":\"ins-xxx2\",\"InstanceName\":\"xxx2\",\"InstanceState\":\"RUNNING\",\"PrivateIPAddresses\":[\"1.17.64.11\"],\"PublicIPAddresses\":null,\"Region\":\"ap-shanghai\",\"Tags\":null,\"Zone\":\"ap-shanghai-2\"}]","StrategyId":9}}

		RiskFieldsDesc []*RiskFieldsDesc `json:"RiskFieldsDesc,omitempty" name:"RiskFieldsDesc"`
		// 评估项ID

		StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
		// 风险实例个数

		RiskTotalCount *uint64 `json:"RiskTotalCount,omitempty" name:"RiskTotalCount"`
		// 风险实例详情列表，需要json decode

		Risks *string `json:"Risks,omitempty" name:"Risks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStrategyRisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStrategyRisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScanTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeReportStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIgnoreInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 忽略的实例列表

		Ignores []*IgnoreInstance `json:"Ignores,omitempty" name:"Ignores"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListIgnoreInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIgnoreInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIgnoreInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IgnoredDetail struct {

	// 实例的忽略状态，true表示被忽略。

	IgnoredStatus *bool `json:"IgnoredStatus,omitempty" name:"IgnoredStatus"`
	// 实例信息。

	Instance *Instance `json:"Instance,omitempty" name:"Instance"`
}

type ProductProgress struct {

	// 产品名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品扫描进度。

	Progress *Progress `json:"Progress,omitempty" name:"Progress"`
}

type DescribeRiskDisplayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 显示字段详情

		DisplayFields []*RiskFieldsDescDetail `json:"DisplayFields,omitempty" name:"DisplayFields"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDisplayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDisplayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportLanguageRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSupportLanguageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportLanguageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRegionCodesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListRegionCodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRegionCodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreInstancesRequest struct {
	*tchttp.BaseRequest

	// 忽略扫描的策略的ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 忽略的实例的列表

	IgnoreInstances []*IgnoreInstance `json:"IgnoreInstances,omitempty" name:"IgnoreInstances"`
	// 增加: add，或者删除: delete。默认值为add

	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

func (r *ModifyIgnoreInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRoleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadReportFileRequest struct {
	*tchttp.BaseRequest

	// 评估结果对应产品的Id，与Type字段配合使用。当Type为Group时，表示类别Id（-1：总览，1：安全，2：可靠，3：服务限制），当Type为Strategy时，表示策略Id。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 产品类型，可为Group（类别）和Strategy（策略）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 任务Id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DownloadReportFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadReportFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadReportFileAsyncRequest struct {
	*tchttp.BaseRequest

	// 评估结果对应产品的Id，与Type字段配合使用。当Type为Group时，表示类别Id（-1：总览，1：安全，2：可靠，3：服务限制），当Type为Strategy时，表示策略Id。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 产品类型，可为Group（类别）和Strategy（策略）。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 任务Id。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 需要下载的产品列表

	Products []*string `json:"Products,omitempty" name:"Products"`
	// 需要下载的维度列表

	GroupIds []*uint64 `json:"GroupIds,omitempty" name:"GroupIds"`
	// 需要下载的巡检项列表

	StrategyIds []*uint64 `json:"StrategyIds,omitempty" name:"StrategyIds"`
	// 环境

	Env *string `json:"Env,omitempty" name:"Env"`
	// 标签

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

func (r *DownloadReportFileAsyncRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadReportFileAsyncRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigGroup struct {

	// 策略组的ID

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略组的名字,例如:安全

	Name *string `json:"Name,omitempty" name:"Name"`
	// 策略组包含的策略组列表

	Strategies []*ConfigStrategy `json:"Strategies,omitempty" name:"Strategies"`
}

type DescribeTaskProgressRequest struct {
	*tchttp.BaseRequest

	// 扫描任务ID。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGlobalIgnoreTagsRequest struct {
	*tchttp.BaseRequest

	// tags值

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyGlobalIgnoreTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGlobalIgnoreTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigCondition struct {

	// 条件ID

	ConditionId *int64 `json:"ConditionId,omitempty" name:"ConditionId"`
	// 此条件的安全等级,1低风险 2中风险 3高风险

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 此条件的描述 eg:对所有人允许上传和删除访问权限

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type DescribeOverviewRequest struct {
	*tchttp.BaseRequest

	// 环境

	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnsafeDetail struct {

	// 该实例命中的条件ID。

	ConditionId *int64 `json:"ConditionId,omitempty" name:"ConditionId"`
	// 实例的忽略状态，true表示被忽略。

	IgnoredStatus *bool `json:"IgnoredStatus,omitempty" name:"IgnoredStatus"`
	// 实例信息。

	Instance *Instance `json:"Instance,omitempty" name:"Instance"`
}

type DescribeSubscriptionTemplatesRequest struct {
	*tchttp.BaseRequest

	// 返回限制个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件,支持name, id 过滤

	Filters *Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSubscriptionTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否完成，true：是，false：否。

		Finished *bool `json:"Finished,omitempty" name:"Finished"`
		// 任务的进度信息（仅在Finished为false时有意义）。如果数据库中不存在进度记录时返回值为{0, 0}。

		Overview *Progress `json:"Overview,omitempty" name:"Overview"`
		// 预估的任务剩余时间（单位：秒，仅在Finished为false时有意义，-1表示评估任务初始化中）。

		EstimatedTime *int64 `json:"EstimatedTime,omitempty" name:"EstimatedTime"`
		// 任务下每个产品的进度详情（仅在Finished为false时有意义）。

		Details []*ProductProgress `json:"Details,omitempty" name:"Details"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStrategyUnsafeDetailRequest struct {
	*tchttp.BaseRequest

	// 任务Id。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 策略Id。

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 过滤条件。

	Filter *string `json:"Filter,omitempty" name:"Filter"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTaskStrategyUnsafeDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStrategyUnsafeDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRegionCodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地区的编码列表

		Codes []*RegionCode `json:"Codes,omitempty" name:"Codes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListRegionCodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRegionCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupSummary struct {

	// 分组Id，-1：总览，1：安全，2：可靠，3：服务限制，4：成本，5：性能。

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 该分组下高风险的策略数量(不包括忽略策略数量)。

	HighRiskCount *int64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
	// 该分组下中风险的策略数量(不包括忽略策略数量)。

	MediumRiskCount *int64 `json:"MediumRiskCount,omitempty" name:"MediumRiskCount"`
	// 该分组下低风险的策略数量(不包括忽略策略数量)。

	LowRiskCount *int64 `json:"LowRiskCount,omitempty" name:"LowRiskCount"`
	// 该分组下无风险的策略数量(不包括忽略策略数量)。

	NoRiskCount *int64 `json:"NoRiskCount,omitempty" name:"NoRiskCount"`
}
