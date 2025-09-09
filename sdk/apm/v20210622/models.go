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

package v20210622

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type OpenApmPaidVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *OpenApmPaidVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenApmPaidVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumptionTotal struct {

	// 实例数目

	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
	// 限流实例数目

	LimitedNum *int64 `json:"LimitedNum,omitempty" name:"LimitedNum"`
	// 上报总量

	TotalReport *int64 `json:"TotalReport,omitempty" name:"TotalReport"`
	// 日期列表

	Dates []*string `json:"Dates,omitempty" name:"Dates"`
}

type Line struct {

	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标中文名

	MetricNameCN *string `json:"MetricNameCN,omitempty" name:"MetricNameCN"`
	// 时间序列

	TimeSerial []*int64 `json:"TimeSerial,omitempty" name:"TimeSerial"`
	// 数据序列

	DataSerial []*float64 `json:"DataSerial,omitempty" name:"DataSerial"`
	// 维度列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
}

type SpanReference struct {

	// 关联关系类型

	RefType *string `json:"RefType,omitempty" name:"RefType"`
	// Span ID

	SpanID *string `json:"SpanID,omitempty" name:"SpanID"`
	// Trace ID

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
}

type YAxis struct {

	// 接口列表

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
	// 接口列表

	Data []*string `json:"Data,omitempty" name:"Data"`
}

type DescribeApmPAASDeduplicateSpanListRequest struct {
	*tchttp.BaseRequest

	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// span查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// span查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 实例名

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 通用过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// paas业务特有名称

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
}

func (r *DescribeApmPAASDeduplicateSpanListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmPAASDeduplicateSpanListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralSpanDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Span列表

		Spans []*Span `json:"Spans,omitempty" name:"Spans"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralSpanDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralSpanDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点集合

		Nodes []*TopologyNode `json:"Nodes,omitempty" name:"Nodes"`
		// 边集合

		Edges []*TopologyEdge `json:"Edges,omitempty" name:"Edges"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表项总个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 111

		TreeListView []*TraceItemView `json:"TreeListView,omitempty" name:"TreeListView"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTraceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTraceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComboInfoOverview struct {

	// 总剩余AgentHour个数

	TotalAgentHourCount *int64 `json:"TotalAgentHourCount,omitempty" name:"TotalAgentHourCount"`
	// 总资源个数

	TotalResourceCount *int64 `json:"TotalResourceCount,omitempty" name:"TotalResourceCount"`
}

type DBTag struct {

	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 键值

	Key *string `json:"Key,omitempty" name:"Key"`
	// 父键值

	ParentKey *string `json:"ParentKey,omitempty" name:"ParentKey"`
}

type InstanceConsumption struct {

	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 日期列表

	Dates []*string `json:"Dates,omitempty" name:"Dates"`
	// 上报量列表

	Reports []*int64 `json:"Reports,omitempty" name:"Reports"`
	// 存储量列表

	Stores []*int64 `json:"Stores,omitempty" name:"Stores"`
}

type DescribeServiceNodesRequest struct {
	*tchttp.BaseRequest

	// 排序方式

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeServiceNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmInfoByAppIdRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeApmInfoByAppIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInfoByAppIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogAgentStatus struct {

	// node ip

	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
	// pod name

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 最近心跳时间

	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`
	// agent状态（0=正常，1=异常）

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type AppConfig struct {

	// 应用名

	Service *string `json:"Service,omitempty" name:"Service"`
	// 采集地址

	IncludePaths []*string `json:"IncludePaths,omitempty" name:"IncludePaths"`
}

type LogAgentInstancesData struct {

	// agent列表信息

	List []*LogAgentInstance `json:"List,omitempty" name:"List"`
}

type QueryMetricItem struct {

	// 同比，已弃用，不建议使用

	Compare *string `json:"Compare,omitempty" name:"Compare"`
	// 同比，支持多种同比方式

	Compares []*string `json:"Compares,omitempty" name:"Compares"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type ApmLogCluster struct {

	// 实例应用

	Service *string `json:"Service,omitempty" name:"Service"`
	// 实例集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 实例集群agent版本

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// 实例集群可升级状态

	UpgradeStatus *int64 `json:"UpgradeStatus,omitempty" name:"UpgradeStatus"`
	// 实力集群运行状态

	ClusterStatus *int64 `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 实例集群节点数量

	ClusterAgentNums *int64 `json:"ClusterAgentNums,omitempty" name:"ClusterAgentNums"`
	// 实例集群可采集状态

	ClusterCollectionStatus *int64 `json:"ClusterCollectionStatus,omitempty" name:"ClusterCollectionStatus"`
	// 集群id

	CID *int64 `json:"CID,omitempty" name:"CID"`
	// 可编辑状态

	IsEditable *int64 `json:"IsEditable,omitempty" name:"IsEditable"`
}

type ApmSampleConfig struct {

	// 实例ID

	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 采样名字

	SampleName *string `json:"SampleName,omitempty" name:"SampleName"`
	// 接口名

	OperationName *string `json:"OperationName,omitempty" name:"OperationName"`
	// 采样的span数

	SpanNum *int64 `json:"SpanNum,omitempty" name:"SpanNum"`
	// 采样配置开关 0 关 1 开

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// tags数组

	Tags []*APMKVItem `json:"Tags,omitempty" name:"Tags"`
	// 采样率

	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
}

type LogPipeOperator struct {

	// 算子id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 算子名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 算子类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 算子作用的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 规则定义

	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
	// 解析失败异常处理逻辑

	ExceptionStrategy *string `json:"ExceptionStrategy,omitempty" name:"ExceptionStrategy"`
	// 最近更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 扩展字段

	Tags []*APMKVItem `json:"Tags,omitempty" name:"Tags"`
	// 前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 是否删除

	Delete *string `json:"Delete,omitempty" name:"Delete"`
	// 算子类型

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 解析字符串

	SampleText *string `json:"SampleText,omitempty" name:"SampleText"`
	// 0禁用 1启用

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
}

type DescribeTagValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 维度值列表

		Values []*string `json:"Values,omitempty" name:"Values"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyMetricLineDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标曲线返回数据

		Lines []*Line `json:"Lines,omitempty" name:"Lines"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopologyMetricLineDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyMetricLineDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 上游层级

	UpLevel *int64 `json:"UpLevel,omitempty" name:"UpLevel"`
	// 127.0.0.1

	ServiceInstance *string `json:"ServiceInstance,omitempty" name:"ServiceInstance"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 查询开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 下游层级

	DownLevel *int64 `json:"DownLevel,omitempty" name:"DownLevel"`
}

func (r *DescribeTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceComboInfo struct {

	// 套餐包总览数据

	Overview *ComboInfoOverview `json:"Overview,omitempty" name:"Overview"`
	// 套餐包信息列表

	ComboInfoList []*ComboInfo `json:"ComboInfoList,omitempty" name:"ComboInfoList"`
}

type DescribeApmComboInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeApmComboInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmComboInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricLineDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标曲线返回数据

		Lines []*Line `json:"Lines,omitempty" name:"Lines"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricLineDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricLineDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*ApmMetricsPoint `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSqlMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSqlMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmSampleConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 采样名

	SampleName *string `json:"SampleName,omitempty" name:"SampleName"`
}

func (r *DescribeApmSampleConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmSampleConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlMetricDataRequest struct {
	*tchttp.BaseRequest

	// 需要查询的指标，不可自定义输入。
	// 支持：sql_duration_avg（平均响应时间）、sql_last_error_occur_time（最后发生时间）、slow_sql_count（慢SQL数）、
	// sql_error_request_count（错误数）、sql_request_count（请求数）。

	Metrics []*string `json:"Metrics,omitempty" name:"Metrics"`
	// 起始时间的时间戳，单位为秒，目前只支持查询2天内最多1小时的指标数据。

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间的时间戳，单位为秒，目前只支持查询2天内最多1小时的指标数据。

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库名称

	DbInstance *string `json:"DbInstance,omitempty" name:"DbInstance"`
	// 实例ip

	DbIp *string `json:"DbIp,omitempty" name:"DbIp"`
	// 查询限制条数，不输入时默认查询5条。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSqlMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSqlMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogStandardConfig struct {

	// 配置Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 原子段

	Source *string `json:"Source,omitempty" name:"Source"`
	// 目标字段

	Target *string `json:"Target,omitempty" name:"Target"`
	// 标准化算子

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DeleteApmSampleConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApmSampleConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmAgentRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 接入方式

	AgentType *string `json:"AgentType,omitempty" name:"AgentType"`
	// 环境

	NetworkMode *string `json:"NetworkMode,omitempty" name:"NetworkMode"`
	// 语言

	LanguageEnvironment *string `json:"LanguageEnvironment,omitempty" name:"LanguageEnvironment"`
	// 上报方式

	ReportMethod *string `json:"ReportMethod,omitempty" name:"ReportMethod"`
}

func (r *DescribeApmAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmSampleConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 采样名

	SampleName *string `json:"SampleName,omitempty" name:"SampleName"`
	// 采样率

	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
	// 接口名

	OperationName *string `json:"OperationName,omitempty" name:"OperationName"`
	// 采样tag

	Tags []*APMKVItem `json:"Tags,omitempty" name:"Tags"`
	// 采样开关 0关 1开 2删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyApmSampleConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmSampleConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComboInfo struct {

	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 套餐包名字

	ComboName *string `json:"ComboName,omitempty" name:"ComboName"`
	// 套餐包生效时间

	ProductionTime *int64 `json:"ProductionTime,omitempty" name:"ProductionTime"`
	// 套餐包过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 套餐包状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 套餐包剩余额度

	RemainingQuota *int64 `json:"RemainingQuota,omitempty" name:"RemainingQuota"`
	// 套餐包总额度

	TotalQuota *int64 `json:"TotalQuota,omitempty" name:"TotalQuota"`
}

type LogPath struct {

	// 文件路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 文件名

	File *string `json:"File,omitempty" name:"File"`
}

type SpanProcess struct {

	// 应用服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// Tags 标签数组

	Tags []*SpanTag `json:"Tags,omitempty" name:"Tags"`
}

type DescribeApmRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// apm可用地区列表

		Regions []*ApmRegion `json:"Regions,omitempty" name:"Regions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricLineDataRequest struct {
	*tchttp.BaseRequest

	// 曲线聚合周期（60,3600,86400）

	AggregatePeriod *int64 `json:"AggregatePeriod,omitempty" name:"AggregatePeriod"`
	// 指标查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 指标名列表

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Top N 曲线参数

	TopLevel *int64 `json:"TopLevel,omitempty" name:"TopLevel"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 指标查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度（服务级别，接口级别）
	// service
	// interface_name

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
}

func (r *DescribeMetricLineDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricLineDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type APMKV struct {

	// Key值定义

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value值定义

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type Filter struct {

	// 过滤方式（=, !=, in）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 过滤维度名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 过滤值，in过滤方式用逗号分割多个值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogMetric struct {

	// 维度维度数组

	Dims []*LogDims `json:"Dims,omitempty" name:"Dims"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 过滤器数组

	Filters []*LogFilters `json:"Filters,omitempty" name:"Filters"`
	// 指标值

	MetricValue *MetricValue `json:"MetricValue,omitempty" name:"MetricValue"`
	// 聚合粒度

	Aggregation *int64 `json:"Aggregation,omitempty" name:"Aggregation"`
	// 计算方式

	Calculation *string `json:"Calculation,omitempty" name:"Calculation"`
}

type DescribeSpanTreeByIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 树形图

		SpanTreeView *SpanTreeView `json:"SpanTreeView,omitempty" name:"SpanTreeView"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpanTreeByIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanTreeByIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmAppId struct {

	// 剩余天数

	RemainderNumberDays *int64 `json:"RemainderNumberDays,omitempty" name:"RemainderNumberDays"`
	// 剩余小时

	RemainderNumberHours *int64 `json:"RemainderNumberHours,omitempty" name:"RemainderNumberHours"`
	// 剩余额度

	RemainderNumberSpans *int64 `json:"RemainderNumberSpans,omitempty" name:"RemainderNumberSpans"`
	// 总额度

	TotalNumbersSpans *int64 `json:"TotalNumbersSpans,omitempty" name:"TotalNumbersSpans"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 试用期过期时间戳

	EndTrialPeriodTime *int64 `json:"EndTrialPeriodTime,omitempty" name:"EndTrialPeriodTime"`
}

type SpanLog struct {

	// 日志时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 标签

	Fields []*SpanTag `json:"Fields,omitempty" name:"Fields"`
}

type DescribeExistInstancesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeExistInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExistInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceAnalyzeSpanListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 链路耗时分析总信息

		AnalyzeView *TraceAnalyzeView `json:"AnalyzeView,omitempty" name:"AnalyzeView"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTraceAnalyzeSpanListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTraceAnalyzeSpanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpanTag struct {

	// 标签类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 标签Key
	// 注意：此字段可能返回 null，表示取不到有效值。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ApmRegion struct {

	// 地区全称

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地区ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 中文别名

	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
	// 缩写

	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`
}

type StatsView struct {

	// 当前分钟日志数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 当前分钟

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type LogCollectConfig struct {

	// 采集路径

	IncludePaths []*LogPath `json:"IncludePaths,omitempty" name:"IncludePaths"`
	// 排查路径

	ExcludeFiles []*string `json:"ExcludeFiles,omitempty" name:"ExcludeFiles"`
	// 编码

	Encoding *string `json:"Encoding,omitempty" name:"Encoding"`
	// 尾部采集

	TailFiles *bool `json:"TailFiles,omitempty" name:"TailFiles"`
	// 多行采集配置

	MultiLine *string `json:"MultiLine,omitempty" name:"MultiLine"`
	// 采集行

	IncludeLines []*string `json:"IncludeLines,omitempty" name:"IncludeLines"`
	// 排除行

	ExcludeLines []*string `json:"ExcludeLines,omitempty" name:"ExcludeLines"`
	// tke路径配置

	PathConfig *PathConfig `json:"PathConfig,omitempty" name:"PathConfig"`
	// 额外字段

	ExtraFields []*ExtraField `json:"ExtraFields,omitempty" name:"ExtraFields"`
}

type DescribeApmApplicationConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Apm应用配置

		ApmAppConfig *ApmAppConfig `json:"ApmAppConfig,omitempty" name:"ApmAppConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmApplicationConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmMetricRecord struct {

	// field数组

	Fields []*ApmField `json:"Fields,omitempty" name:"Fields"`
	// tag数组

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
}

type TraceItemView struct {

	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// trace唯一标识符

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
	// 入口接口名

	InterfaceName *string `json:"InterfaceName,omitempty" name:"InterfaceName"`
	// 开始采样时间

	SamplingTime *string `json:"SamplingTime,omitempty" name:"SamplingTime"`
	// 入口应用

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 耗时（单位毫秒）

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// span唯一标识符

	SpanID *string `json:"SpanID,omitempty" name:"SpanID"`
}

type DescribeGeneralSpanListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Span分页列表

		Spans []*Span `json:"Spans,omitempty" name:"Spans"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralSpanListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralSpanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePAASTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拓扑节点数组

		Data []*ApmTopologyNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePAASTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePAASTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopologySnapshotRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例ID的节点及位置信息

	Snapshot []*Snapshot `json:"Snapshot,omitempty" name:"Snapshot"`
}

func (r *CreateTopologySnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopologySnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmPAASDeduplicateSpanListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Span分页列表

		Spans []*Span `json:"Spans,omitempty" name:"Spans"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmPAASDeduplicateSpanListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmPAASDeduplicateSpanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpanTagListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表项

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpanTagListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanTagListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceAnalyzeSpanListRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// traceID

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// span查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// span查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// spanID

	SpanID *string `json:"SpanID,omitempty" name:"SpanID"`
}

func (r *DescribeTraceAnalyzeSpanListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTraceAnalyzeSpanListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmField struct {

	// 昨日同比指标值，已弃用，不建议使用

	CompareVal *string `json:"CompareVal,omitempty" name:"CompareVal"`
	// Compare值结果数组，推荐使用

	CompareVals []*APMKVItem `json:"CompareVals,omitempty" name:"CompareVals"`
	// 指标值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 指标所对应的单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 请求数

	Key *string `json:"Key,omitempty" name:"Key"`
	// 同环比上周期具体数值

	LastPeriodValue []*APMKV `json:"LastPeriodValue,omitempty" name:"LastPeriodValue"`
}

type DescribeServiceNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricPointDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 多指标返回数据

		Points []*ApmMetricsPoint `json:"Points,omitempty" name:"Points"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricPointDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricPointDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmApplicationConfigRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 异常过滤正则规则，逗号分隔

	ExceptionFilter *string `json:"ExceptionFilter,omitempty" name:"ExceptionFilter"`
	// URL收敛正则规则，逗号分隔

	UrlConvergence *string `json:"UrlConvergence,omitempty" name:"UrlConvergence"`
	// 错误码过滤，逗号分隔

	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitempty" name:"ErrorCodeFilter"`
	// URL收敛开关,0 关 | 1 开

	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitempty" name:"UrlConvergenceSwitch"`
	// URL收敛阈值

	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitempty" name:"UrlConvergenceThreshold"`
	// URL排除正则规则，逗号分隔

	UrlExclude *string `json:"UrlExclude,omitempty" name:"UrlExclude"`
	// 日志开关 0 关 1 开

	IsRelatedLog *int64 `json:"IsRelatedLog,omitempty" name:"IsRelatedLog"`
	// 日志地域

	LogRegion *string `json:"LogRegion,omitempty" name:"LogRegion"`
	// 日志主题ID

	LogTopicID *string `json:"LogTopicID,omitempty" name:"LogTopicID"`
	// CLS 日志集 | ES 集群ID

	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`
	// 日志来源 CLS | ES

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 需过滤的接口

	IgnoreOperationName *string `json:"IgnoreOperationName,omitempty" name:"IgnoreOperationName"`
}

func (r *ModifyApmApplicationConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmApplicationConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Consumption struct {

	// 用量总量

	Total *ConsumptionTotal `json:"Total,omitempty" name:"Total"`
	// 实例用量数组

	Instances []*InstanceConsumption `json:"Instances,omitempty" name:"Instances"`
	// 柱状图数据

	Lines []*ConsumptionLine `json:"Lines,omitempty" name:"Lines"`
}

type LogQueryData struct {

	// 日志起始索引

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 日志页码大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 日志总条数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 日志具体内容

	List []*SearchData `json:"List,omitempty" name:"List"`
}

type Snapshot struct {

	// 节点id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 节点位置信息

	Position *Position `json:"Position,omitempty" name:"Position"`
}

type SpanAnalyzeView struct {

	// 应用名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 接口名称

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 总耗时

	DurationTotal *uint64 `json:"DurationTotal,omitempty" name:"DurationTotal"`
	// 平均耗时

	DurationAvg *uint64 `json:"DurationAvg,omitempty" name:"DurationAvg"`
	// 请求次数

	ReqCount *uint64 `json:"ReqCount,omitempty" name:"ReqCount"`
	// ip

	IP *string `json:"IP,omitempty" name:"IP"`
	// traceID

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
}

type DescribeTagCountValuesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 维度名

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 聚合维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeTagCountValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagCountValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点集合

		Nodes []*TopologyNode `json:"Nodes,omitempty" name:"Nodes"`
		// 边集合

		Edges []*TopologyEdgeNew `json:"Edges,omitempty" name:"Edges"`
		// 拓扑图是否有修改

		TopologyModifyFlag *int64 `json:"TopologyModifyFlag,omitempty" name:"TopologyModifyFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopologyNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TraceAnalyzeView struct {

	// 主span调用耗时分析数据

	Analyze *SpanAnalyzeView `json:"Analyze,omitempty" name:"Analyze"`
	// 链路耗时分析子span耗时信息

	Children []*ChildrenSpanView `json:"Children,omitempty" name:"Children"`
	// 链路耗时分析调用类型耗时信息

	Components []*ComponentCostView `json:"Components,omitempty" name:"Components"`
}

type LogInstanceInfo struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否异步查询，0=同步，1=异步

	Async *int64 `json:"Async,omitempty" name:"Async"`
	// 存储时长

	SaveDays *int64 `json:"SaveDays,omitempty" name:"SaveDays"`
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmInstanceBrief struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeGeneralSpanListRequest struct {
	*tchttp.BaseRequest

	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// span查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 实例名

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 列表项个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 通用过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 业务自身服务名

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
	// span查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeGeneralSpanListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralSpanListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineExportTaskData struct {

	// 离线导出任务数组

	List []*OfflineExportTaskView `json:"List,omitempty" name:"List"`
	// 页码索引

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页码大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 总数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DescribeMetricDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Metric Point

		Point *ApmMetricsPoint `json:"Point,omitempty" name:"Point"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// 查询指标结果集条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogPipelineParam struct {

	// pipeline id

	Id *string `json:"Id,omitempty" name:"Id"`
	// pipeline fid

	Fid *string `json:"Fid,omitempty" name:"Fid"`
	// pipeline名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤条件

	Filters []*SearchFilter `json:"Filters,omitempty" name:"Filters"`
	// 启用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type SpanTreeView struct {

	// 树结构json串

	TreeData *string `json:"TreeData,omitempty" name:"TreeData"`
	// 列表

	ChartData *ChartData `json:"ChartData,omitempty" name:"ChartData"`
	// Trace概览信息

	TraceOverview *TraceOverview `json:"TraceOverview,omitempty" name:"TraceOverview"`
	// 压缩编码后的树结构json串

	CompressedTreeData *string `json:"CompressedTreeData,omitempty" name:"CompressedTreeData"`
}

type TopologyEdge struct {

	// 源节点

	Source *string `json:"Source,omitempty" name:"Source"`
	// 边ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 边权重

	Weight *float64 `json:"Weight,omitempty" name:"Weight"`
	// 目标节点

	Target *string `json:"Target,omitempty" name:"Target"`
}

type DescribeInstanceBriefsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// apm实例列表

		Instances []*ApmInstanceBrief `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceBriefsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceBriefsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBTags struct {

	// 地域标签

	RegionTags []*RegionTag `json:"RegionTags,omitempty" name:"RegionTags"`
	// 实例标签

	TapmInstanceKeyTags []*DBTag `json:"TapmInstanceKeyTags,omitempty" name:"TapmInstanceKeyTags"`
	// 数据库实例标签

	DBInstanceTags []*DBTag `json:"DBInstanceTags,omitempty" name:"DBInstanceTags"`
	// 数据库ip标签

	DBIpTags []*DBTag `json:"DBIpTags,omitempty" name:"DBIpTags"`
}

type InstanceRegion struct {

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type LogPipeline struct {

	// Pipeline id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 上游pipline ID

	Fid *string `json:"Fid,omitempty" name:"Fid"`
	// pipeline类型

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// Pipeline名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤条件

	Filters []*SearchFilter `json:"Filters,omitempty" name:"Filters"`
	// 算子数组

	Operators []*LogPipeOperator `json:"Operators,omitempty" name:"Operators"`
	// 最近修改时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 0 禁用 1 启用

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
}

type TerminateApmInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateApmInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmAppConfig struct {

	// 实例ID

	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// URL收敛开关

	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitempty" name:"UrlConvergenceSwitch"`
	// URL收敛阈值

	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitempty" name:"UrlConvergenceThreshold"`
	// URL收敛正则

	UrlConvergence *string `json:"UrlConvergence,omitempty" name:"UrlConvergence"`
	// 异常过滤正则

	ExceptionFilter *string `json:"ExceptionFilter,omitempty" name:"ExceptionFilter"`
	// 错误码过滤

	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitempty" name:"ErrorCodeFilter"`
	// 服务组件类型

	Components *string `json:"Components,omitempty" name:"Components"`
	// URL排除正则

	UrlExclude *string `json:"UrlExclude,omitempty" name:"UrlExclude"`
	// 日志来源

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 日志所在地域

	LogRegion *string `json:"LogRegion,omitempty" name:"LogRegion"`
	// 是否开启日志 0 关 1 开

	IsRelatedLog *int64 `json:"IsRelatedLog,omitempty" name:"IsRelatedLog"`
	// 日志主题ID

	LogTopicID *string `json:"LogTopicID,omitempty" name:"LogTopicID"`
	// 需过滤的接口名

	IgnoreOperationName *string `json:"IgnoreOperationName,omitempty" name:"IgnoreOperationName"`
}

type PodLabel struct {

	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 表达式

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type SearchFilter struct {

	// 过滤的键的名称

	Key *string `json:"Key,omitempty" name:"Key"`
	// 过滤的方式

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 过滤的键的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeMetricDetailRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 指标名列表

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
}

func (r *DescribeMetricDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogClusterInfo struct {

	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// agent版本

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// agent最新版本

	LastVersion *string `json:"LastVersion,omitempty" name:"LastVersion"`
	// 是否可升级

	Upgrade *bool `json:"Upgrade,omitempty" name:"Upgrade"`
	// 集群绑定应用

	Services []*string `json:"Services,omitempty" name:"Services"`
	// 异常节点数量

	AbnormalNum *int64 `json:"AbnormalNum,omitempty" name:"AbnormalNum"`
	// 节点数量

	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`
	// 正常节点数量

	NormalNum *int64 `json:"NormalNum,omitempty" name:"NormalNum"`
	// helm chart版本

	ChartVersion *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
}

type DescribeGeneralMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*Line `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmTag struct {

	// 维度Key(列名，标签Key)

	Key *string `json:"Key,omitempty" name:"Key"`
	// 维度值（标签值）

	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogDims struct {

	// 维度名

	DimName *string `json:"DimName,omitempty" name:"DimName"`
	// 日志字段

	LogDim *string `json:"LogDim,omitempty" name:"LogDim"`
}

type Span struct {

	// Trace Id

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
	// 日志

	Logs []*SpanLog `json:"Logs,omitempty" name:"Logs"`
	// 标签

	Tags []*SpanTag `json:"Tags,omitempty" name:"Tags"`
	// 上报应用服务信息

	Process *SpanProcess `json:"Process,omitempty" name:"Process"`
	// 产生时间戳(毫秒)

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// Span名称

	OperationName *string `json:"OperationName,omitempty" name:"OperationName"`
	// 关联关系

	References []*SpanReference `json:"References,omitempty" name:"References"`
	// 产生时间戳(微秒)

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 持续耗时(微妙)

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// Span Id

	SpanID *string `json:"SpanID,omitempty" name:"SpanID"`
	// 产生时间戳(毫秒)

	StartTimeMillis *int64 `json:"StartTimeMillis,omitempty" name:"StartTimeMillis"`
}

type DescribeApmRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeApmRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmAgentInfo struct {

	// Agent下载地址

	AgentDownloadURL *string `json:"AgentDownloadURL,omitempty" name:"AgentDownloadURL"`
	// Collector上报地址

	CollectorURL *string `json:"CollectorURL,omitempty" name:"CollectorURL"`
	// Token信息

	Token *string `json:"Token,omitempty" name:"Token"`
	// 外网上报地址

	PublicCollectorURL *string `json:"PublicCollectorURL,omitempty" name:"PublicCollectorURL"`
	// 自研VPC上报地址

	InnerCollectorURL *string `json:"InnerCollectorURL,omitempty" name:"InnerCollectorURL"`
	// 内网上报地址(Private Link上报地址)

	PrivateLinkCollectorURL *string `json:"PrivateLinkCollectorURL,omitempty" name:"PrivateLinkCollectorURL"`
}

type ChartData struct {

	// 接口参数id列表

	Series []*Series `json:"Series,omitempty" name:"Series"`
	// 接口参数列表

	YAxis []*YAxis `json:"YAxis,omitempty" name:"YAxis"`
	// 接口参数列表

	YAxisEx *YAxis `json:"YAxisEx,omitempty" name:"YAxisEx"`
}

type ComponentCostView struct {

	// 调用类型

	Component *string `json:"Component,omitempty" name:"Component"`
	// 总执行时间

	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
	// 时间占比

	Proportion *float64 `json:"Proportion,omitempty" name:"Proportion"`
}

type RegionTag struct {

	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 键值

	Key *string `json:"Key,omitempty" name:"Key"`
	// 地域id

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域

	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`
	// 缩写

	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`
}

type DescribeApmAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Agent信息

		ApmAgent *ApmAgentInfo `json:"ApmAgent,omitempty" name:"ApmAgent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmApplicationConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApmApplicationConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenApmPaidVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenApmPaidVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenApmPaidVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApmSampleConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采样配置参数

		ApmSampleConfig *ApmSampleConfig `json:"ApmSampleConfig,omitempty" name:"ApmSampleConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApmSampleConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogClusterStatus struct {

	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 正常数量

	NormalNum *int64 `json:"NormalNum,omitempty" name:"NormalNum"`
	// 异常数量

	AbnormalNum *int64 `json:"AbnormalNum,omitempty" name:"AbnormalNum"`
	// 节点数量

	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`
	// agent状态列表

	Agents []*LogAgentStatus `json:"Agents,omitempty" name:"Agents"`
}

type LogPipelineDetail struct {

	// Pipeline节点

	Pipelines []*LogPipeline `json:"Pipelines,omitempty" name:"Pipelines"`
	// Pipeline 关系

	Chains []*LogPipelineChain `json:"Chains,omitempty" name:"Chains"`
}

type DescribeExistInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true为该用户存在APM实例，反之

		Exist *bool `json:"Exist,omitempty" name:"Exist"`
		// true为该用户上报过数据，反之

		IsReportData *bool `json:"IsReportData,omitempty" name:"IsReportData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExistInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExistInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumptionLine struct {

	// 上报量

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 存储量

	MetricNameCN *string `json:"MetricNameCN,omitempty" name:"MetricNameCN"`
	// 时间序列

	TimeSerial []*int64 `json:"TimeSerial,omitempty" name:"TimeSerial"`
	// 数值数组

	DataSerial []*int64 `json:"DataSerial,omitempty" name:"DataSerial"`
}

type DescribeInstanceBriefsRequest struct {
	*tchttp.BaseRequest

	// Tag列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 是否查询官方demo实例

	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitempty" name:"DemoInstanceFlag"`
}

func (r *DescribeInstanceBriefsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceBriefsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogInstanceInfoData struct {

	// 指标数据

	Metrics []*ApmMetric `json:"Metrics,omitempty" name:"Metrics"`
	// 维度数据

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 曲线数据

	Lines []*Line `json:"Lines,omitempty" name:"Lines"`
	// 实例集群信息数据

	Clusters []*ApmLogCluster `json:"Clusters,omitempty" name:"Clusters"`
}

type ApmTopologyEdge struct {

	// 上游节点ID

	SourceID *string `json:"SourceID,omitempty" name:"SourceID"`
	// 服务端侧指标（Target Metric）

	ServerMetrics []*ApmMetric `json:"ServerMetrics,omitempty" name:"ServerMetrics"`
	// 下游节点ID

	TargetID *string `json:"TargetID,omitempty" name:"TargetID"`
	// 客户端侧指标（Source Metric）

	ClientMetrics []*ApmMetric `json:"ClientMetrics,omitempty" name:"ClientMetrics"`
	// 维度信息

	NodeDimensions []*ApmTag `json:"NodeDimensions,omitempty" name:"NodeDimensions"`
}

type Series struct {

	// 耗时列表信息

	Data []*int64 `json:"Data,omitempty" name:"Data"`
	// 描述

	Stack *string `json:"Stack,omitempty" name:"Stack"`
	// 字段名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeServiceLinkRequest struct {
	*tchttp.BaseRequest

	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指标列表

	Metrics []*QueryMetricItem `json:"Metrics,omitempty" name:"Metrics"`
	// 每页大小

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 分页起始点

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 业务名称（默认值：taw）

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
	// 页码

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页长

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeServiceLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceLinkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmInstanceDetail struct {

	// 存储使用量(MB)

	AmountOfUsedStorage *float64 `json:"AmountOfUsedStorage,omitempty" name:"AmountOfUsedStorage"`
	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例所属tag列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 创建人Uin

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 该实例已上报的服务端应用数量

	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
	// 日均上报Span数

	CountOfReportSpanPerDay *int64 `json:"CountOfReportSpanPerDay,omitempty" name:"CountOfReportSpanPerDay"`
	// AppId信息

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// Trace数据保存时长

	TraceDuration *int64 `json:"TraceDuration,omitempty" name:"TraceDuration"`
	// 实例描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 实例状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例上报额度

	SpanDailyCounters *int64 `json:"SpanDailyCounters,omitempty" name:"SpanDailyCounters"`
	// 实例是否开通计费

	BillingInstance *int64 `json:"BillingInstance,omitempty" name:"BillingInstance"`
	// 错误率阈值

	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitempty" name:"ErrRateThreshold"`
	// 采样率阈值

	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
	// 是否开启错误采样 0  关 1 开

	ErrorSample *int64 `json:"ErrorSample,omitempty" name:"ErrorSample"`
	// 慢调用保存阈值

	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitempty" name:"SlowRequestSavedThreshold"`
	// cls日志所在地域

	LogRegion *string `json:"LogRegion,omitempty" name:"LogRegion"`
	// 日志来源

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 日志功能开关 0 关 | 1 开

	IsRelatedLog *int64 `json:"IsRelatedLog,omitempty" name:"IsRelatedLog"`
	// 日志主题ID

	LogTopicID *string `json:"LogTopicID,omitempty" name:"LogTopicID"`
	// 该实例已上报的客户端应用数量

	ClientCount *int64 `json:"ClientCount,omitempty" name:"ClientCount"`
	// 该实例已上报的总应用数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type LogClusterConfig struct {

	// 配置id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 日志源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 服务名

	Service *string `json:"Service,omitempty" name:"Service"`
	// 接入方式

	Access *string `json:"Access,omitempty" name:"Access"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 采集配置

	Configs []*LogCollectConfig `json:"Configs,omitempty" name:"Configs"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 状态（0=正常，1=禁用，2=删除）

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type CreateApmSampleConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 采样名

	SampleName *string `json:"SampleName,omitempty" name:"SampleName"`
	// 采样率

	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
	// 采样Tags

	Tags []*APMKVItem `json:"Tags,omitempty" name:"Tags"`
	// 接口名

	OperationName *string `json:"OperationName,omitempty" name:"OperationName"`
}

func (r *CreateApmSampleConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApmSampleConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmComboInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 套餐包页面相关信息

		ResourceComboInfo *ResourceComboInfo `json:"ResourceComboInfo,omitempty" name:"ResourceComboInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmComboInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmComboInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Position struct {

	// 节点位置横坐标

	X *float64 `json:"X,omitempty" name:"X"`
	// 节点位置纵坐标

	Y *float64 `json:"Y,omitempty" name:"Y"`
}

type DescribeApmInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// apm实例列表

		Instances []*ApmInstanceDetail `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineExportTaskView struct {

	// 离线任务id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 离线任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 接入点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务导出信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 任务导出状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 任务导出时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 日志下载url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 实例名称

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 日志导出状态

	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`
}

type ApmTopologyNode struct {

	// 拓扑名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件类型

	Component *string `json:"Component,omitempty" name:"Component"`
	// 拓扑节点ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 当前节点指标

	Metrics []*ApmMetric `json:"Metrics,omitempty" name:"Metrics"`
	// 上游关系边数据

	EdgeList []*ApmTopologyEdge `json:"EdgeList,omitempty" name:"EdgeList"`
	// 服务端测指标

	ServerMetrics []*ApmMetric `json:"ServerMetrics,omitempty" name:"ServerMetrics"`
	// 客户端测指标

	ClientMetrics []*ApmMetric `json:"ClientMetrics,omitempty" name:"ClientMetrics"`
	// 维度信息

	Dimensions []*ApmTag `json:"Dimensions,omitempty" name:"Dimensions"`
}

type DescribeApmInstanceOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmInstanceOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInstanceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceListRequest struct {
	*tchttp.BaseRequest

	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// sql通用过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 列表项个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// trace查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// trace查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeTraceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTraceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmInstancesRequest struct {
	*tchttp.BaseRequest

	// Tag列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 搜索实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 过滤实例ID

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否查询官方demo实例

	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitempty" name:"DemoInstanceFlag"`
}

func (r *DescribeApmInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type APMKVItem struct {

	// Key值定义

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value值定义

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeSpanTagListRequest struct {
	*tchttp.BaseRequest

	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 启始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指标列表

	Metrics []*QueryMetricItem `json:"Metrics,omitempty" name:"Metrics"`
	// 翻页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 列表页需要展示的Tag字段

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 翻页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 查询过滤条件

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
}

func (r *DescribeSpanTagListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanTagListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApmInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmConsumptionRequest struct {
	*tchttp.BaseRequest

	// "all"获取用量总量统计，"report"获取上报量柱状图数据，"store"获取存储量柱状图数据

	Metric *string `json:"Metric,omitempty" name:"Metric"`
}

func (r *DescribeApmConsumptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmConsumptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogAgentInstance struct {

	// agent版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// agent ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// agent状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// agent状态码

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 唯一识别码

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type GeneralFilter struct {

	// 过滤维度名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 过滤值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeDBTagValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 标签组

		Values *DBTags `json:"Values,omitempty" name:"Values"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBTagValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpanTreeByIDRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// traceID

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// span查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// span查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 代表请求的是否为完整瀑布图

	IsFold *bool `json:"IsFold,omitempty" name:"IsFold"`
}

func (r *DescribeSpanTreeByIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanTreeByIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 标签列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 实例详情

	Description *string `json:"Description,omitempty" name:"Description"`
	// Trace数据保存时长

	TraceDuration *int64 `json:"TraceDuration,omitempty" name:"TraceDuration"`
	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否开启计费

	OpenBilling *bool `json:"OpenBilling,omitempty" name:"OpenBilling"`
	// 实例上报额度

	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitempty" name:"SpanDailyCounters"`
	// 错误率阈值

	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitempty" name:"ErrRateThreshold"`
	// 采样率

	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
	// 是否开启错误采样 0 关 1 开

	ErrorSample *int64 `json:"ErrorSample,omitempty" name:"ErrorSample"`
	// 慢请求阈值

	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitempty" name:"SlowRequestSavedThreshold"`
	// 是否开启日志功能 0 关 1 开

	IsRelatedLog *int64 `json:"IsRelatedLog,omitempty" name:"IsRelatedLog"`
	// 日志地域

	LogRegion *string `json:"LogRegion,omitempty" name:"LogRegion"`
	// CLS日志主题ID | ES 索引名

	LogTopicID *string `json:"LogTopicID,omitempty" name:"LogTopicID"`
	// CLS日志集 | ES集群ID

	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`
	// CLS | ES

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
}

func (r *ModifyApmInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateApmInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TerminateApmInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateApmInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmSampleConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采样配置列表

		ApmSampleConfigs []*ApmSampleConfig `json:"ApmSampleConfigs,omitempty" name:"ApmSampleConfigs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmSampleConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogAgentConfig struct {

	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 服务

	Service *string `json:"Service,omitempty" name:"Service"`
	// 采集路径

	IncludePaths []*string `json:"IncludePaths,omitempty" name:"IncludePaths"`
	// 排除路径

	ExcludePaths []*string `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 编码方式

	Encoding *string `json:"Encoding,omitempty" name:"Encoding"`
	// 是否尾部采集  true尾部采集，false 头部采集

	TailFiles *bool `json:"TailFiles,omitempty" name:"TailFiles"`
	// 多行分隔符

	MultiLine *string `json:"MultiLine,omitempty" name:"MultiLine"`
	// 包含行

	IncludeLines []*string `json:"IncludeLines,omitempty" name:"IncludeLines"`
	// 排除行

	ExcludeLines []*string `json:"ExcludeLines,omitempty" name:"ExcludeLines"`
	// 网络模式

	NetworkMode *string `json:"NetworkMode,omitempty" name:"NetworkMode"`
	// 配置Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type CreateTopologySnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopologySnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopologySnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmConsumptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用量统计

		Consumption *Consumption `json:"Consumption,omitempty" name:"Consumption"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmConsumptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmConsumptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmMetricsPoint struct {

	// 指标列表

	Metrics []*ApmMetric `json:"Metrics,omitempty" name:"Metrics"`
	// 维度列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
}

type MetricValue struct {

	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type TopologyEdgeNew struct {

	// 源节点

	Source *string `json:"Source,omitempty" name:"Source"`
	// 边ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 边权重

	Weight *float64 `json:"Weight,omitempty" name:"Weight"`
	// 目标节点

	Target *string `json:"Target,omitempty" name:"Target"`
	// 响应时间

	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
	// 错误率

	ErrRate *float64 `json:"ErrRate,omitempty" name:"ErrRate"`
	// 吞吐量

	Qps *float64 `json:"Qps,omitempty" name:"Qps"`
	// 边类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 边颜色

	Color *string `json:"Color,omitempty" name:"Color"`
	// Sql调用数

	SqlRequestCount *float64 `json:"SqlRequestCount,omitempty" name:"SqlRequestCount"`
	// Sql调用错误数

	SqlErrorRequestCount *float64 `json:"SqlErrorRequestCount,omitempty" name:"SqlErrorRequestCount"`
}

type DescribeApmApplicationConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeApmApplicationConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmApplicationConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFieldMappingOperation struct {

	// 操作

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 索引

	Fields []*LogFieldMapping `json:"Fields,omitempty" name:"Fields"`
}

type PathConfig struct {

	// 路径类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// k8s namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// pod标签

	Labels []*PodLabel `json:"Labels,omitempty" name:"Labels"`
	// 工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 工作负载名

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
}

type DescribeLogMonitorDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控数据

		Lines []*Line `json:"Lines,omitempty" name:"Lines"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogMonitorDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopologyNode struct {

	// 错误率

	ErrRate *float64 `json:"ErrRate,omitempty" name:"ErrRate"`
	// 节点类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 节点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点权重

	Weight *float64 `json:"Weight,omitempty" name:"Weight"`
	// 节点颜色

	Color *string `json:"Color,omitempty" name:"Color"`
	// 响应时间

	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
	// 吞吐量

	Qps *float64 `json:"Qps,omitempty" name:"Qps"`
	// 节点类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 节点大小

	Size *string `json:"Size,omitempty" name:"Size"`
	// 节点是否为组件类型

	IsModule *bool `json:"IsModule,omitempty" name:"IsModule"`
	// 节点位置信息

	Position *Position `json:"Position,omitempty" name:"Position"`
}

type DescribeServiceOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFilters struct {

	// 字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 正则表达式

	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
	// 日志处理类型

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeLogMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 筛选条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLogMonitorDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogMonitorDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogAsyncQueryData struct {

	// 日志起始索引

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 日志页码大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 日志总条数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 日志具体内容

	List []*SearchData `json:"List,omitempty" name:"List"`
	// 查询任务状态(0=初始化，1=运行中，2=成功，3=失败)

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeMetricPointDataRequest struct {
	*tchttp.BaseRequest

	// 指标查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 指标名列表

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 指标查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度（服务级别，接口级别）
	// service
	// interface_name

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
}

func (r *DescribeMetricPointDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricPointDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricRecordsRequest struct {
	*tchttp.BaseRequest

	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指标列表

	Metrics []*QueryMetricItem `json:"Metrics,omitempty" name:"Metrics"`
	// 每页大小

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 分页起始点

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 业务名称（默认值：taw）

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
	// 页码

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页长

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeMetricRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceOverviewRequest struct {
	*tchttp.BaseRequest

	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指标列表

	Metrics []*QueryMetricItem `json:"Metrics,omitempty" name:"Metrics"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 分页起始点

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
}

func (r *DescribeServiceOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagValuesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 维度名

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeTagValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyNewRequest struct {
	*tchttp.BaseRequest

	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 上游层级

	UpLevel *int64 `json:"UpLevel,omitempty" name:"UpLevel"`
	// 实例信息

	ServiceInstance *string `json:"ServiceInstance,omitempty" name:"ServiceInstance"`
	// 查询开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 下游层级

	DownLevel *int64 `json:"DownLevel,omitempty" name:"DownLevel"`
	// 视角

	View *string `json:"View,omitempty" name:"View"`
	// 过滤器

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 表示Topic（MQ拓扑图用）

	Topic *string `json:"Topic,omitempty" name:"Topic"`
}

func (r *DescribeTopologyNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderBy struct {

	// 需要排序的字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 顺序排序/倒序排序

	Value *string `json:"Value,omitempty" name:"Value"`
}

type PAASInstance struct {

	// Token值

	Token *string `json:"Token,omitempty" name:"Token"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 实例 ID

	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateApmInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// Trace数据保存时长

	TraceDuration *int64 `json:"TraceDuration,omitempty" name:"TraceDuration"`
	// 标签列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 实例上报额度值

	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitempty" name:"SpanDailyCounters"`
}

func (r *CreateApmInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApmInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagCountValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TagCount的键值对

		TagCountKV []*TagCountKV `json:"TagCountKV,omitempty" name:"TagCountKV"`
		// Tag维度分组查询总数

		TotalTagCount *int64 `json:"TotalTagCount,omitempty" name:"TotalTagCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagCountValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagCountValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChildrenSpanView struct {

	// 接口名称

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 调用类型

	Component *string `json:"Component,omitempty" name:"Component"`
	// 调用次数

	ReqCount *uint64 `json:"ReqCount,omitempty" name:"ReqCount"`
	// 总执行时间

	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
	// 时间占比

	Proportion *float64 `json:"Proportion,omitempty" name:"Proportion"`
}

type CreateApmInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApmInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePAASTopologyRequest struct {
	*tchttp.BaseRequest

	// 本地节点ID字段数组，取值用于标识节点唯一性并识别调用关系

	LocalNodeIdFields []*string `json:"LocalNodeIdFields,omitempty" name:"LocalNodeIdFields"`
	// 下游追溯层数（相对于FocusNodeId标记的节点）

	DownstreamLayer *int64 `json:"DownstreamLayer,omitempty" name:"DownstreamLayer"`
	// 查询过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询视图

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 上游追溯层数（相对于FocusNodeId标记的节点）

	UpstreamLayer *int64 `json:"UpstreamLayer,omitempty" name:"UpstreamLayer"`
	// 实例ID（团队ID/数据隔离ID）

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 节点维度

	NodeDimensions []*string `json:"NodeDimensions,omitempty" name:"NodeDimensions"`
	// 组件类型字段名

	ComponentKey *string `json:"ComponentKey,omitempty" name:"ComponentKey"`
	// 节点指标

	NodeMetrics []*string `json:"NodeMetrics,omitempty" name:"NodeMetrics"`
	// 标记节点ID，配合UpstreamLayer和DownstreamLayer查询指定层数上下游

	FocusNodeId *string `json:"FocusNodeId,omitempty" name:"FocusNodeId"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 业务名称（默认值：taw）

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
	// Span 类型字段名

	SpanKindKey *string `json:"SpanKindKey,omitempty" name:"SpanKindKey"`
	// 边指标名称数组

	EdgeMetrics []*string `json:"EdgeMetrics,omitempty" name:"EdgeMetrics"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 远程节点ID字段数组，取值用于标识节点唯一性并识别调用关系

	PeerNodeIdFields []*string `json:"PeerNodeIdFields,omitempty" name:"PeerNodeIdFields"`
	// 远程节点组件类型字段名

	PeerComponentKey *string `json:"PeerComponentKey,omitempty" name:"PeerComponentKey"`
}

func (r *DescribePAASTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePAASTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EtlResult struct {

	// 解析状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 解析结果map

	Fields []*LogFieldMapping `json:"Fields,omitempty" name:"Fields"`
}

type ExtraField struct {

	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeServiceLinkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBTagValuesRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 数据库IP

	DBIp *string `json:"DBIp,omitempty" name:"DBIp"`
	// 业务名称

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
}

func (r *DescribeDBTagValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBTagValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogInstanceView struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 索引配置

	FieldMapping []*LogFieldMapping `json:"FieldMapping,omitempty" name:"FieldMapping"`
	// pipeline详情

	Pipeline *LogPipelineDetail `json:"Pipeline,omitempty" name:"Pipeline"`
	// 标准化配置

	StandardConfig []*LogStandardConfig `json:"StandardConfig,omitempty" name:"StandardConfig"`
}

type DescribeTopologyMetricLineDataRequest struct {
	*tchttp.BaseRequest

	// 曲线聚合周期（60,3600,86400）

	AggregatePeriod *int64 `json:"AggregatePeriod,omitempty" name:"AggregatePeriod"`
	// 指标查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 指标名列表

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Top N 曲线参数

	TopLevel *int64 `json:"TopLevel,omitempty" name:"TopLevel"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 指标查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度（服务级别，接口级别）

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
}

func (r *DescribeTopologyMetricLineDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyMetricLineDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchData struct {

	// 日志时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 日志偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 日志所属实例

	Host *string `json:"Host,omitempty" name:"Host"`
	// 日志内容

	Message *string `json:"Message,omitempty" name:"Message"`
	// 日志级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 上报源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 服务名

	Service *string `json:"Service,omitempty" name:"Service"`
	// 索引字段

	IndexFields *string `json:"IndexFields,omitempty" name:"IndexFields"`
	// 扩展字段

	ExtFields *string `json:"ExtFields,omitempty" name:"ExtFields"`
}

type TagCountKV struct {

	// 维度Key值

	Key *string `json:"Key,omitempty" name:"Key"`
	// 维度对应的count数

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DeleteApmSampleConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 采样名

	SampleName *string `json:"SampleName,omitempty" name:"SampleName"`
}

func (r *DeleteApmSampleConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApmSampleConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogPipelineChain struct {

	// 父pipeline id

	Fid *string `json:"Fid,omitempty" name:"Fid"`
	// pipeline id

	Id *string `json:"Id,omitempty" name:"Id"`
}

type ApmMetric struct {

	// 指标中文名

	MetricNameCN *string `json:"MetricNameCN,omitempty" name:"MetricNameCN"`
	// 指标值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type FieldView struct {

	// 字段类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 可搜索字段名称

	Key *string `json:"Key,omitempty" name:"Key"`
}

type LogPipeOperatorParam struct {

	// 算子id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 算子名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 算子类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 算子作用的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 示例日志

	SampleText *string `json:"SampleText,omitempty" name:"SampleText"`
	// 模式

	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
	// 解析失败策略

	ExceptionStrategy *string `json:"ExceptionStrategy,omitempty" name:"ExceptionStrategy"`
	// 扩展字段

	Tags []*APMKVItem `json:"Tags,omitempty" name:"Tags"`
	// 是否删除

	Delete *string `json:"Delete,omitempty" name:"Delete"`
	// 前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 是否启用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type TraceOverview struct {

	// traceID调用状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// TraceId查询

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
	// 错误信息

	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 耗时

	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
	// 接口名

	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type DescribeGeneralSpanDetailRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 查询条件列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeGeneralSpanDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralSpanDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFieldMapping struct {

	// 字段名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 字段类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 索引的ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 0表示预设
	// 1表示自定义

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
}

type DescribeApmInfoByAppIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// appId相关apm返回具体信息

		Data *ApmAppId `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmInfoByAppIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInfoByAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmInstanceOverviewRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指标列表

	Metrics []*QueryMetricItem `json:"Metrics,omitempty" name:"Metrics"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 业务名称（默认值：taw）

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
}

func (r *DescribeApmInstanceOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInstanceOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogMetricData struct {

	// 指标数据

	Metrics []*ApmMetric `json:"Metrics,omitempty" name:"Metrics"`
	// 维度数据

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 曲线数据

	Lines []*Line `json:"Lines,omitempty" name:"Lines"`
}

type DescribeGeneralMetricDataRequest struct {
	*tchttp.BaseRequest

	// 要过滤的维度信息：
	// service_metric视图支持：service.name（服务名）、span.kind（客户端/服务端视角）为维度进行过滤，service.name（服务名）必填。
	// span.kind:
	// 	server:服务端视角
	// 	client:客户端视角
	// 默认为服务端视角进行查询。
	// runtime_metric视图支持：service.name（服务名）维度进行过滤，service.name（服务名）必填。
	// sql_metric视图支持：service.name（服务名）、db.instance（数据库名称）、db.ip（数据库实例ip）维度进行过滤，查询service_slow_sql_count（慢sql）指标时service.name必填，查询sql_duration_avg（耗时）指标时db.instance（数据库名称）必填。

	Filters []*GeneralFilter `json:"Filters,omitempty" name:"Filters"`
	// 聚合维度：
	// service_metric视图支持：service.name（服务名）、span.kind （客户端/服务端视角）维度进行聚合，service.name（服务名）必填。
	// runtime_metric视图支持：service.name（服务名）维度进行聚合，service.name（服务名）必填。
	// sql_metric视图支持：service.name（服务名）、db.statement（sql语句）维度进行聚合，查询service_slow_sql_count（慢sql）时service.name（服务名）必填，查询sql_duration_avg（耗时）指标时service.name（服务名）、db.statement（sql语句）必填。

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 需要查询的指标，不可自定义输入。
	// service_metric视图支持：service_request_count（总请求）、service_duration（平均响应时间）、service_error_req_rate（平均错误率）、service_slow_call_count（慢调用）、service_error_request_count（异常数量）。
	// runtime_metric视图支持：service_gc_full_count（Full GC）。
	// sql_metric视图支持：service_slow_sql_count（慢sql）、sql_duration_avg（耗时）。

	Metrics []*string `json:"Metrics,omitempty" name:"Metrics"`
	// 起始时间的时间戳，单位为秒，只支持查询2天内最多1小时的指标数据。

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间的时间戳，单位为秒，只支持查询2天内最多1小时的指标数据。

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 聚合粒度，单位为秒，最小为60s，即一分钟的聚合粒度；如果为空或0则计算开始时间到截止时间的指标数据，上报其他值会报错。

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 视图名称，不可自定义输入。支持：service_metric、runtime_metric、sql_metric。

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 对查询指标进行排序：
	// service_metric视图支持：service_request_count（总请求）、service_duration（平均响应时间）、service_error_req_rate（平均错误率）、service_slow_call_count（慢调用）、service_error_request_count（异常数量）。
	// runtime_metric视图支持：service_gc_full_count（Full GC）。
	// sql_metric视图支持：service_slow_sql_count（慢sql）、sql_duration_avg（耗时）。
	// asc:对查询指标进行升序排序
	// desc：对查询指标进行降序排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 查询指标的限制条数，目前最多展示50条数据，PageSize取值为1-50，上送PageSize则根据PageSize的值展示限制条数。

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeGeneralMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmSampleConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApmSampleConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
