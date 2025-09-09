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

package v20201016

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DeleteAsyncContextTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAsyncContextTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAsyncContextTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFunctionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFunctionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFunctionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// CLS对应topic标识

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *DescribeKafkaConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCosRechargeResponseParams struct {
	Id        *string `json:"Id,omitempty" name:"Id"`
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCosRechargeResponse struct {
	*tchttp.BaseResponse

	Response *CreateCosRechargeResponseParams `json:"Response"`
}

func (r *CreateCosRechargeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigRequest struct {
	*tchttp.BaseRequest

	// 采集规则配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 采集规则配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 日志采集路径，包含文件名

	Path *string `json:"Path,omitempty" name:"Path"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 采集配置关联的日志主题（TopicId）

	Output *string `json:"Output,omitempty" name:"Output"`
	// 用户自定义解析字符串，Json格式序列化的字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// config_extra表主键ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集配置标

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：  ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时 ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数 ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false 样例： {\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true} 控制台默认占位值：{\"ClsAgentDefault\":0}
	AdvancedConfig *string `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`

	// 采集配置来源，0：默认来源，1: TKE
	Source *uint64 `json:"Source,omitempty" name:"Source"`
}

func (r *ModifyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataTransformRequest struct {
	*tchttp.BaseRequest

	// 函数类型. DSL:1 SQL:2

	FuncType *int64 `json:"FuncType,omitempty" name:"FuncType"`
	// 源日志主题

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 加工任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 加工逻辑函数

	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`
	// 任务启动状态.   默认为1，正常开启,  2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 加工任务目的topic_id以及别名

	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`
	// 任务类型.  以SrcTopicId为数据源建立预览任务:1，以PreviewLogStatistics为数据源建立预览任务:2  真实任务:3

	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
	// 测试数据

	PreviewLogStatistics []*PreviewLogStatistic `json:"PreviewLogStatistics,omitempty" name:"PreviewLogStatistics"`
}

func (r *CreateDataTransformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataTransformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContentInfo struct {

	// 内容格式，支持json、csv

	Format *string `json:"Format,omitempty" name:"Format"`
	// csv格式内容描述

	Csv *CsvInfo `json:"Csv,omitempty" name:"Csv"`
	// json格式内容描述

	Json *JsonInfo `json:"Json,omitempty" name:"Json"`
	// parquet格式内容描述

	Parquet *ParquetInfo `json:"Parquet,omitempty" name:"Parquet"`
}

type KeyRegexInfo struct {

	// 需要过滤日志的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// key对应的过滤规则regex

	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

type TopicIndexInfo struct {

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 索引是否生效

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 索引配置信息 注意：此字段可能返回 null，表示取不到有效值。

	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
	// 索引修改时间，初始值为索引创建时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 全文索引系统预置字段标记，默认false。  false:不包含系统预置字段， true:包含系统预置字段

	IncludeInternalFields *bool `json:"IncludeInternalFields,omitempty" name:"IncludeInternalFields"`
	// 元数据相关标志位，默认为0。 0：全文索引包括开启键值索引的元数据字段， 1：全文索引包括所有元数据字段，2：全文索引不包括元数据字段。

	MetadataFlag *uint64 `json:"MetadataFlag,omitempty" name:"MetadataFlag"`
	// 自定义日志解析异常存储字段。

	CoverageField *string `json:"CoverageField,omitempty" name:"CoverageField"`
}

type CreateIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmNoticeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardRequest struct {
	*tchttp.BaseRequest

	// 仪表盘id

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
	// 仪表盘名称

	DashboardName *string `json:"DashboardName,omitempty" name:"DashboardName"`
	// 仪表盘配置数据

	Data *string `json:"Data,omitempty" name:"Data"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAlarmChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 测试结果

		ChannelTestResults []*ChannelTestResult `json:"ChannelTestResults,omitempty" name:"ChannelTestResults"`
		// 是否成功进行测试

		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
		// 错误原因

		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAlarmChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAlarmChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserConfigRequest struct {
	*tchttp.BaseRequest

	// 用户配置内容

	Data []*UserConfigInfo `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyUserConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskLogStatistic struct {

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 读取的源日志主题的行数

	ReadLines *int64 `json:"ReadLines,omitempty" name:"ReadLines"`
	// 加工后的行数

	WriteLines *int64 `json:"WriteLines,omitempty" name:"WriteLines"`
	// 加工失败的行数

	FailedLines *int64 `json:"FailedLines,omitempty" name:"FailedLines"`
	// 输出到目标日志主题的总体统计数据

	DstTopicLogStatistics []*TopicIdLogStatistic `json:"DstTopicLogStatistics,omitempty" name:"DstTopicLogStatistics"`
	// 加工过滤的行数

	FilterLines *uint64 `json:"FilterLines,omitempty" name:"FilterLines"`
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机器状态信息组

		Machines []*MachineInfo `json:"Machines,omitempty" name:"Machines"`
		// 机器组是否开启自动升级功能

		AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
		// 机器组自动升级功能预设开始时间

		UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
		// 机器组自动升级功能预设结束时间

		UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
		// 当前用户可用最新的Loglistener版本

		LatestAgentVersion *string `json:"LatestAgentVersion,omitempty" name:"LatestAgentVersion"`
		// 是否开启服务日志

		ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
		// 默认值""，"label_k8s"

		Flag *string `json:"Flag,omitempty" name:"Flag"`
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

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警通知模板列表。

		AlarmNotices []*AlarmNotice `json:"AlarmNotices,omitempty" name:"AlarmNotices"`
		// 符合条件的告警通知模板总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigExtrasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集配置列表

		Configs []*ConfigExtraInfo `json:"Configs,omitempty" name:"Configs"`
		// 过滤到的总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigExtrasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigExtrasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsyncSearchTask struct {

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 状态，0表示待开始，1表示运行中，2表示已完成，-1表示失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 异步检索任务ID

	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`
	// 查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 日志扫描顺序，可选值：asc(升序)、desc(降序)

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 任务失败的错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	// 异步检索任务匹配的总日志条数

	LogCount *int64 `json:"LogCount,omitempty" name:"LogCount"`
	// 任务完成时间

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
}

type ContainerWorkLoadInfo struct {

	// 容器名

	Container *string `json:"Container,omitempty" name:"Container"`
	// 工作负载的类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 工作负载的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type QueryTemplateItem struct {

	// 检索语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 检索语句名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteAsyncSearchTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAsyncSearchTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAsyncSearchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigExtraResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集配置扩展信息ID

		ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigExtraResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HeartBeatRequest struct {
	*tchttp.BaseRequest

	// agent的版本号

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// agent采集的数据大小

	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// agent序列号

	AgentSeq *string `json:"AgentSeq,omitempty" name:"AgentSeq"`
	// 标签列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// Agent是否开启自动升级功能

	AutoUpdate *string `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 是否需要检查自动升级任务

	CheckUpdate *string `json:"CheckUpdate,omitempty" name:"CheckUpdate"`
}

func (r *HeartBeatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HeartBeatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourcesInfo struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 日志集数

	Logsets *uint64 `json:"Logsets,omitempty" name:"Logsets"`
	// 日志主题数

	Topics *uint64 `json:"Topics,omitempty" name:"Topics"`
	// 分区数

	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`
	// 机器数

	Machines *uint64 `json:"Machines,omitempty" name:"Machines"`
	// 心跳正常机器数

	HeartMachines *uint64 `json:"HeartMachines,omitempty" name:"HeartMachines"`
}

type RemoveMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalysisDimensional struct {

	// 分析名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分析类型：query，field

	Type *string `json:"Type,omitempty" name:"Type"`
	// 分析内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 配置

	ConfigInfo []*AlarmAnalysisConfig `json:"ConfigInfo,omitempty" name:"ConfigInfo"`
}

type ChannelTestResult struct {

	// 序号

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 错误码，0是正确

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	// 发送结果

	SendTotal *SendDetail `json:"SendTotal,omitempty" name:"SendTotal"`
}

type DescribeAsyncContextResultRequest struct {
	*tchttp.BaseRequest

	// 异步检索任务ID

	AsyncContextTaskId *string `json:"AsyncContextTaskId,omitempty" name:"AsyncContextTaskId"`
	// 日志包序号

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 日志在日志包内的序号

	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 上文日志条数，默认值10

	PrevLogs *int64 `json:"PrevLogs,omitempty" name:"PrevLogs"`
	// 下文日志条数，默认值10

	NextLogs *int64 `json:"NextLogs,omitempty" name:"NextLogs"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeAsyncContextResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncContextResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志上下文信息集合

		LogContextInfos []*LogContextInfo `json:"LogContextInfos,omitempty" name:"LogContextInfos"`
		// 上文日志是否已经返回

		PrevOver *bool `json:"PrevOver,omitempty" name:"PrevOver"`
		// 下文日志是否已经返回

		NextOver *bool `json:"NextOver,omitempty" name:"NextOver"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogContextResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogItem struct {

	// 日志Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 日志Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CheckAlarmRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控对象测试结果

		AlarmRuleTestResults []*AlarmRuleTestResult `json:"AlarmRuleTestResults,omitempty" name:"AlarmRuleTestResults"`
		// 触发条件测试结果

		ConditionTestResult *ConditionTestResult `json:"ConditionTestResult,omitempty" name:"ConditionTestResult"`
		// 是否成功检测

		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
		// 错误原因

		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAlarmRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组名字，不能重复

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 创建机器组类型，Type为ip，Values中为Ip字符串列表创建机器组，Type为label， Values中为标签字符串列表创建机器组

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的机器组。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否开启机器组自动更新

	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 升级开始时间，建议业务低峰期升级LogListener

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 升级结束时间，建议业务低峰期升级LogListener

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费

	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
	// 默认值""，"label_k8s"

	Flag *string `json:"Flag,omitempty" name:"Flag"`

	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitempty" name:"DelayCleanupTime"`

	OSType *uint64 `json:"OSType,omitempty" name:"OSType"`

	MetaTags []*MetaTagInfo `json:"MetaTags,omitempty" name:"MetaTags"`
}

func (r *CreateMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardRequest struct {
	*tchttp.BaseRequest

	// 仪表盘id

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
}

func (r *DeleteDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集配置ID

		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineGroupTypeInfo struct {

	// 机器组类型，ip表示该机器组Values中存的是采集机器的IP地址，label表示该机器组Values中存储的是机器的标签

	Type *string `json:"Type,omitempty" name:"Type"`
	// 机器描述列表

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type TemplateItem struct {

	// 模版项ID

	TemplateItemId *string `json:"TemplateItemId,omitempty" name:"TemplateItemId"`
	// 模版项名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源类型：'LOGSET' | 'TOPIC' | 'DASHBOARD' | 'ALARM' | 'ALARM_NOTICE'

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 模版数据

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteDemonstrationRequest struct {
	*tchttp.BaseRequest

	// 演示示例类型：'CLB'

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteDemonstrationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDemonstrationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformPreviewInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1: 任务完成， 2: 任务处理中，3: 任务处理失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 错误信息

		FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformPreviewInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformPreviewInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题的索引信息列表

		TopicIndexInfos []*TopicIndexInfo `json:"TopicIndexInfos,omitempty" name:"TopicIndexInfos"`
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndexsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDataTransformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDataTransformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertHistoryNotice struct {

	// 通知名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 通知ID

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
}

type DescribeDeliverCloudFunctionRequest struct {
	*tchttp.BaseRequest

	// 投递规则属于的 topic id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeDeliverCloudFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeliverCloudFunctionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigExtraRequest struct {
	*tchttp.BaseRequest

	// 采集配置扩展信息id

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 节点文件配置信息

	HostFile *HostFileInfo `json:"HostFile,omitempty" name:"HostFile"`
	// 容器文件路径信息

	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitempty" name:"ContainerFile"`
	// 容器标准输出信息

	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitempty" name:"ContainerStdout"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 用户自定义采集规则，Json格式序列化的字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// 类型：container_stdout、container_file、host_file

	Type *string `json:"Type,omitempty" name:"Type"`
	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 自建采集配置标

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集name

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题name

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *ModifyConfigExtraRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigExtraRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenKVRegexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正则表达式

		Regex *string `json:"Regex,omitempty" name:"Regex"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenKVRegexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenKVRegexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ckafka struct {

	// Ckafka 的 Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Ckafka 的 Vport

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// Ckafka 的 InstanceId

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Ckafka 的 InstanceName

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// Ckafka 的 TopicId

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// Ckafka 的 TopicName

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type Demonstration struct {

	// 演示示例资源所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 演示示例类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 演示示例资源

	Resources []*DemonstrationResource `json:"Resources,omitempty" name:"Resources"`
	// 演示示例状态：CREATING, FAILED, SUCCESS, DELETING

	Status *string `json:"Status,omitempty" name:"Status"`
}

type SendDetailItem struct {

	// Uin或者Gin

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Uin或者Group或者Http或者Wecom

	Type *string `json:"Type,omitempty" name:"Type"`
	// 成功数

	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
}

type ModifyDataTransformRequest struct {
	*tchttp.BaseRequest

	// 加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 加工任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 加工逻辑函数

	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`
	// 任务启动状态. 默认为1，正常开启,  2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 加工任务目的topic_id以及别名

	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`
}

func (r *ModifyDataTransformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataTransformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseKafkaConsumeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseKafkaConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKafkaConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QcloudGoodsInfoListInfo struct {

	// 付费模式，0:后付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 业务产品录入的商品码，业务名称，categoryid中的商品码

	Type *string `json:"Type,omitempty" name:"Type"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 区域ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 商品实例的个数

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 业务参数，用户询价和透传给业务

	GoodsDetail *QcloudGoodsDetailInfo `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
}

type CreateConfigExtraRequest struct {
	*tchttp.BaseRequest

	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 类型：container_stdout、container_file、host_file

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点文件配置信息

	HostFile *HostFileInfo `json:"HostFile,omitempty" name:"HostFile"`
	// 容器文件路径信息

	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitempty" name:"ContainerFile"`
	// 容器标准输出信息

	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitempty" name:"ContainerStdout"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 用户自定义采集规则，Json格式序列化的字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// 绑定的机器组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 采集配置标

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集name

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 绑定的机器组id列表

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

func (r *CreateConfigExtraRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigExtraRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加载后续内容的Context，过期时间1小时

		Context *string `json:"Context,omitempty" name:"Context"`
		// 原始日志查询结果是否全部返回。查询语句(Query)包含SQL时该参数无意义

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 返回的是否为分析结果

		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`
		// 如果Analysis为True，则返回分析结果的列名，否则为空

		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`
		// 日志查询结果；当Analysis为True时，可能返回为null

		Results []*LogInfo `json:"Results,omitempty" name:"Results"`
		// 日志分析结果；当Analysis为False时，可能返回为null

		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`
		// 新的日志分析结果; UseNewAnalysis为true有效

		AnalysisRecords []*string `json:"AnalysisRecords,omitempty" name:"AnalysisRecords"`
		// 日志分析的列属性; UseNewAnalysis为true有效

		Columns []*Column `json:"Columns,omitempty" name:"Columns"`
		// 返回语法优化后的语句（QueryOptimize 为 1 时返回，其他情况返回空字符串）

		Query *string `json:"Query,omitempty" name:"Query"`
		// 当前系统使用的采样率SamplingRate（主要是当客户输入0时，返回真实后台的AutoSamplingRate值）

		SamplingRate *float64 `json:"SamplingRate,omitempty" name:"SamplingRate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JsonLogInfo struct {

	// K-V形式日志信息

	Log []*KeyLogInfo `json:"Log,omitempty" name:"Log"`
	// K-V形式标签信息

	Tag []*KeyLogInfo `json:"Tag,omitempty" name:"Tag"`
	// 时间戳

	Time *uint64 `json:"Time,omitempty" name:"Time"`
}

type DescribeAccountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAgentStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAgentStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAgentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbTopicExtendConfig struct {

	// LB关键信息，VIP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// LB关键信息，VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// clb服务端的公共topic

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// clb用户的topic

	UserTopicId *string `json:"UserTopicId,omitempty" name:"UserTopicId"`
	// clb用户的uin信息

	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`
	// clb用户的appid信息

	UserAppId *uint64 `json:"UserAppId,omitempty" name:"UserAppId"`
	// 临时证书加密密钥ID。最长不超过1024字节。

	UserTmpSecretId *string `json:"UserTmpSecretId,omitempty" name:"UserTmpSecretId"`
	// 临时证书加密密钥Key。最长不超过1024字节。

	UserTmpSecretKey *string `json:"UserTmpSecretKey,omitempty" name:"UserTmpSecretKey"`
	// token, 最长不超过4096字节。

	UserToken *string `json:"UserToken,omitempty" name:"UserToken"`
	// 临时证书有效的时间，返回 Unix 时间戳，精确到秒

	TmpKeyExpired *uint64 `json:"TmpKeyExpired,omitempty" name:"TmpKeyExpired"`
	// 唯一标识clb的一种业务

	LbKey *string `json:"LbKey,omitempty" name:"LbKey"`
	// 公共topic的采样比

	LogSample *string `json:"LogSample,omitempty" name:"LogSample"`
	// 用户topic的采样比

	UserSample *string `json:"UserSample,omitempty" name:"UserSample"`
	// LB健康检查日志 Topic ID, 和topicId属于另外一种公共的topic

	UserHealthTopicId *string `json:"UserHealthTopicId,omitempty" name:"UserHealthTopicId"`
	// topic的采集配置是否生效,true为生效

	UserSampleStatus *bool `json:"UserSampleStatus,omitempty" name:"UserSampleStatus"`
	// 1代表用户topic已删除，0代表用户topic未删除。

	UserTopicStatus *uint64 `json:"UserTopicStatus,omitempty" name:"UserTopicStatus"`
	// lbkey是否要采集到公共topic, true为要采集， false为不采集，默认为false

	Collection *bool `json:"Collection,omitempty" name:"Collection"`
}

type DeleteIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncContextResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上文日志是否已经返回

		PrevOver *bool `json:"PrevOver,omitempty" name:"PrevOver"`
		// 下文日志是否已经返回

		NextOver *bool `json:"NextOver,omitempty" name:"NextOver"`
		// 日志内容

		Results []*LogInfo `json:"Results,omitempty" name:"Results"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncContextResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncContextResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncSearchTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步检索任务列表

		AsyncSearchTasks []*AsyncSearchTask `json:"AsyncSearchTasks,omitempty" name:"AsyncSearchTasks"`
		// 异步检索任务的总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncSearchTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncSearchTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户配置数据

		Data []*UserConfigInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CsvInfo struct {

	// csv首行是否打印key

	PrintKey *bool `json:"PrintKey,omitempty" name:"PrintKey"`
	// 每列key的名字

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
	// 各字段间的分隔符

	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`
	// 若字段内容中包含分隔符，则使用该转义符包裹改字段，只能填写单引号、双引号、空字符串

	EscapeChar *string `json:"EscapeChar,omitempty" name:"EscapeChar"`
	// 对于上面指定的不存在字段使用该内容填充

	NonExistingField *string `json:"NonExistingField,omitempty" name:"NonExistingField"`
}

type LogContextInfo struct {

	// 日志来源设备

	Source *string `json:"Source,omitempty" name:"Source"`
	// 采集路径

	Filename *string `json:"Filename,omitempty" name:"Filename"`
	// 日志内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 日志包序号

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 日志包内一条日志的序号

	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 日志时间戳

	BTime *int64 `json:"BTime,omitempty" name:"BTime"`
	// 日志来源主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

type DeleteMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNonBillingRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 类型含义， 2：topic_region， 目前只支持2

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *CreateNonBillingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNonBillingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// CLS控制台创建的TopicId

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *OpenKafkaConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKafkaConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志集ID

		LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShipperTasksRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// 查询的开始时间戳，支持最近3天的查询， 毫秒

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 查询的结束时间戳， 毫秒

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeShipperTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShipperTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogItems struct {

	// 分析结果返回的KV数据对

	Data []*LogItem `json:"Data,omitempty" name:"Data"`
}

type DescribeIndexsRequest struct {
	*tchttp.BaseRequest

	// <br><li> topicName按照【日志主题名称】进行过滤。类型：String必选：否<br><li> topicId按照【日志主题ID】进行过滤。类型：String必选：否<br><li> logsetId按照【日志集ID】进行过滤，可通过调用DescribeLogsets查询已创建的日志集列表或登录控制台进行查看；也可以调用CreateLogset创建新的日志集。类型：String必选：否<br><li> tagKey按照【标签键】进行过滤。类型：String必选：否<br><li> tag:tagKey按照【标签键值对】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例2。类型：String必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIndexsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题列表

		Topics []*TopicInfo `json:"Topics,omitempty" name:"Topics"`
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigMachineGroupsRequest struct {
	*tchttp.BaseRequest

	// 采集配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeConfigMachineGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigMachineGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserConfigInfo struct {

	// 附加配置key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 附加配置内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题 ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardTopicInfo struct {

	// 主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// topic所在的地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeDataTransformProcessInfoRequest struct {
	*tchttp.BaseRequest

	// 数据加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 要查询的起始时间，Unix时间戳，单位ms

	From *uint64 `json:"From,omitempty" name:"From"`
	// 要查询的结束时间，Unix时间戳，单位ms

	To *uint64 `json:"To,omitempty" name:"To"`
	// 是否需要分成多个时间段获取

	NeedMultTimePeriod *bool `json:"NeedMultTimePeriod,omitempty" name:"NeedMultTimePeriod"`
}

func (r *DescribeDataTransformProcessInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformProcessInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicExtendConfigRequest struct {
	*tchttp.BaseRequest

	// clb的topic业务配置,数组大小不可以超过100

	ClbTopicExtendConfigs []*ClbTopicExtendConfig `json:"ClbTopicExtendConfigs,omitempty" name:"ClbTopicExtendConfigs"`
}

func (r *CreateTopicExtendConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicExtendConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFunctionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可加工的函数列表

		FunctionInfos []*FunctionInfo `json:"FunctionInfos,omitempty" name:"FunctionInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFunctionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFunctionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 需要过滤的值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type HighLightItem struct {

	// 高亮的日志Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 高亮的语法

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ConsumerContent struct {

	// 是否投递 TAG 信息

	EnableTag *bool `json:"EnableTag,omitempty" name:"EnableTag"`
	// 需要投递的元数据列表，目前仅支持：\_\_SOURCE\_\_，\_\_FILENAME\_\_和\_\_TIMESTAMP\_\_

	MetaFields []*string `json:"MetaFields,omitempty" name:"MetaFields"`
	// 当EnableTag为true时，必须填写TagJsonNotTiled字段，TagJsonNotTiled用于标识tag信息是否json平铺，TagJsonNotTiled为true时不平铺，false时平铺

	TagJsonNotTiled *bool `json:"TagJsonNotTiled,omitempty" name:"TagJsonNotTiled"`
	// 投递时间戳精度，可选项 [1:秒；2:毫秒] ，默认是秒

	TimestampAccuracy *int64 `json:"TimestampAccuracy,omitempty" name:"TimestampAccuracy"`
}

type CreateMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机器组ID

		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题 ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogHistogramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计周期： 单位ms

		Interval *int64 `json:"Interval,omitempty" name:"Interval"`
		// 命中关键字的日志总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 周期内统计结果详情

		HistogramInfos []*HistogramInfo `json:"HistogramInfos,omitempty" name:"HistogramInfos"`
		// 返回语法优化后的语句（QueryOptimize 为 1 时返回，其他情况返回空字符串）

		Query *string `json:"Query,omitempty" name:"Query"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogHistogramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeliverCloudFunctionRequest struct {
	*tchttp.BaseRequest

	// 投递规则属于的 topic id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递状态： 0：关闭投递， 1：开启投递

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDeliverCloudFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeliverCloudFunctionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAgentStatusRequest struct {
	*tchttp.BaseRequest

	// Agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// Agent的Instance ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Agent的请求序列号

	AgentSeq *string `json:"AgentSeq,omitempty" name:"AgentSeq"`
	// Agent当前版本

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// Agent升级目标版本

	UpdateVersion *string `json:"UpdateVersion,omitempty" name:"UpdateVersion"`
	// Agent升级状态信息

	AgentStatus *AgentUpdateStatus `json:"AgentStatus,omitempty" name:"AgentStatus"`
}

func (r *UpdateAgentStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAgentStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentUpdateInfo struct {

	// 是否需要升级

	NeedUpdate *bool `json:"NeedUpdate,omitempty" name:"NeedUpdate"`
	// 升级类型:0-null,1-manual,2-auto,3-force

	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 升级动作:0-null,1-update,2-revert

	UpdateAction *int64 `json:"UpdateAction,omitempty" name:"UpdateAction"`
	// 重试次数,最大3次

	RetryCount *int64 `json:"RetryCount,omitempty" name:"RetryCount"`
	// 目标版本

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
	// 安装包下载链接1

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// 安装包下载链接2

	DownloadUrlSecond *string `json:"DownloadUrlSecond,omitempty" name:"DownloadUrlSecond"`
	// 安装包文件MD5值

	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
}

type CreateExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志导出ID。

		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParquetInfo struct {

	// ParquetKeyInfo数组

	ParquetKeyInfo []*ParquetKeyInfo `json:"ParquetKeyInfo,omitempty" name:"ParquetKeyInfo"`
}

type ModifyAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExportsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopicIdLogFailureInfo struct {

	// 日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志失败信息

	LogFailureInfos []*DataTransformFailureInfo `json:"LogFailureInfos,omitempty" name:"LogFailureInfos"`
}

type ApplyConfigToMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyConfigToMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyConfigToMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertRecordHistoryRequest struct {
	*tchttp.BaseRequest

	// key是：alertId,topicId,status

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 启始时间

	From *uint64 `json:"From,omitempty" name:"From"`
	// 终止时间

	To *uint64 `json:"To,omitempty" name:"To"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlertRecordHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertRecordHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RealtimeProducerRequest struct {
	*tchttp.BaseRequest
}

func (r *RealtimeProducerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RealtimeProducerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QcloudDataInfo struct {

	// 后付费订单号，每个物品对应一个dealName

	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
	// 异步发货产品查询发货状态标识

	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
	// 多物品发货对应的flowId，与dealNames一一对应

	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`
	// 每个订单号对应的发货资源id列表:{"20200929112744":["ins-kjs4jvkj"],"20200929112745":["disk-8ijw00wy"]}

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 冻结流水，一次开通一个冻结流水

	BillId *string `json:"BillId,omitempty" name:"BillId"`
}

type DeleteAlarmNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmNoticeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Template struct {

	// 演示示例类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 模版子类型信息

	SubTypes []*SubTypeItem `json:"SubTypes,omitempty" name:"SubTypes"`
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警模板ID

		AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmNoticeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogsetRequest struct {
	*tchttp.BaseRequest

	// 日志集名字，不能重名

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 标签描述列表。最大支持10个标签键值对，并且不能有重复的键值对

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 生命周期，单位天；可取值范围1～366

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *CreateLogsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataTransformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDataTransformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsetsRequest struct {
	*tchttp.BaseRequest

	// <br><li> logsetName
	//
	// 按照【日志集名称】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> logsetId
	//
	// 按照【日志集ID】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> tagKey
	//
	// 按照【标签键】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> tag:tagKey
	//
	// 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。
	// 类型：String
	//
	// 必选：否
	//
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLogsetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseKafkaConsumeRequest struct {
	*tchttp.BaseRequest

	// CLS对应的topic标识

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *CloseKafkaConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKafkaConsumeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmNotice struct {

	// 告警通知模板名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 告警模板的类型。可选值：
	// <br><li> Trigger - 告警触发
	// <br><li> Recovery - 告警恢复
	// <br><li> All - 告警触发和告警恢复

	Type *string `json:"Type,omitempty" name:"Type"`
	// 告警通知模板接收者信息。

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 告警通知模板回调信息。

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
	// 告警通知模板ID。

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AlarmRuleTestResult struct {

	// 位序

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 错误码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type ParquetKeyInfo struct {

	// 键值名称

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 数据类型，目前支持6种类型：string、boolean、int32、int64、float、double

	KeyType *string `json:"KeyType,omitempty" name:"KeyType"`
	// 解析失败赋值信息

	KeyNonExistingField *string `json:"KeyNonExistingField,omitempty" name:"KeyNonExistingField"`
}

type DescribeMachineGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机器组信息列表

		MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitempty" name:"MachineGroups"`
		// 分页的总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyShipperRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// 投递规则投递的新的bucket

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 投递规则投递的新的目录前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 投递规则的开关状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 投递规则的名字

	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`
	// 投递的时间间隔，单位 秒，默认300，范围 300-900

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 投递的文件的最大值，单位 MB，默认256，范围 100-256

	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`
	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递

	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`
	// 投递日志的分区规则，支持strftime的时间格式表示

	Partition *string `json:"Partition,omitempty" name:"Partition"`
	// 投递日志的压缩配置

	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`
	// 投递日志的内容格式配置

	Content *ContentInfo `json:"Content,omitempty" name:"Content"`

	// 投递文件命名配置，0：随机数命名，1：投递时间命名，默认0（随机数命名）
	FilenameMode *uint64 `json:"FilenameMode,omitempty" name:"FilenameMode"`

	// 跨账户uin，用于支持跨账户bucket投递
	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`

	// 投递数据范围的开始时间点，不能超出日志主题的生命周期起点。如果用户不填写，默认为用户新建投递任务的时间。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 投递数据范围的结束时间点，不能填写未来时间。如果用户不填写，默认为持续投递，即无限。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyShipperRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyShipperRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncContextTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步上下文任务ID

		AsyncContextTaskId *string `json:"AsyncContextTaskId,omitempty" name:"AsyncContextTaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAsyncContextTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAsyncContextTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicExtendConfigRequest struct {
	*tchttp.BaseRequest

	// cls的业务标识字段

	LbKeys []*string `json:"LbKeys,omitempty" name:"LbKeys"`
}

func (r *DescribeTopicExtendConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicExtendConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesRequest struct {
	*tchttp.BaseRequest

	// <br><li> Type

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadServiceLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadServiceLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadServiceLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncSearchTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步检索任务ID

		AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAsyncSearchTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAsyncSearchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicExtendConfigRequest struct {
	*tchttp.BaseRequest

	// clb的topic业务配置

	ClbTopicExtendConfigs []*ClbTopicExtendConfig `json:"ClbTopicExtendConfigs,omitempty" name:"ClbTopicExtendConfigs"`
	// 修改模式。 mode取值为1：则用lbkey更新ClbTopicExtendConfig；取值为2：则用user_uin更新FederationToken

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyTopicExtendConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicExtendConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题 ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递任务是否生效，默认不生效

	Effective *bool `json:"Effective,omitempty" name:"Effective"`
	// 是否投递日志的元数据信息，默认为 false

	NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`
	// 如果需要投递元数据信息，元数据信息的描述

	Content *ConsumerContent `json:"Content,omitempty" name:"Content"`
	// CKafka的描述

	Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`
	// 压缩方式，0：NONE，2：SNAPPY，3：LZ4

	Compression *int64 `json:"Compression,omitempty" name:"Compression"`
}

func (r *ModifyConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigFromMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigFromMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigFromMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformPreviewDataInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预览数据详细信息

		PreviewLogStatistics []*PreviewLogStatistic `json:"PreviewLogStatistics,omitempty" name:"PreviewLogStatistics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformPreviewDataInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformPreviewDataInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRebuildIndexTasksRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 索引重建任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 索引重建任务状态，不填返回所有状态任务列表，多种状态之间用逗号分隔，0:索引重建任务已创建，1:已创建索引重建资源，2:重建中，3:重建完成，4:重建成功（可检索），5:任务取消，6:元数据和索引已删除

	Status *string `json:"Status,omitempty" name:"Status"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为10，最大值20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRebuildIndexTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRebuildIndexTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDemonstrationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDemonstrationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDemonstrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebCallback struct {

	// 回调地址。

	Url *string `json:"Url,omitempty" name:"Url"`
	// 回调方法。可选值：
	// <br><li> POST
	// <br><li> PUT
	// 默认值为POST。CallbackType为Http时为必选。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 请求头。

	Headers []*string `json:"Headers,omitempty" name:"Headers"`
	// 请求内容。CallbackType为Http时为必选。

	Body *string `json:"Body,omitempty" name:"Body"`
	// 回调的类型。可选值：
	// <br><li> WeCom
	// <br><li> Http

	CallbackType *string `json:"CallbackType,omitempty" name:"CallbackType"`
	// 序号

	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type DescribeLogHistogramRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 时间间隔: 单位ms

	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
	// 0（默认值）：不执行语法优化；1：执行语法优化。

	QueryOptimize *uint64 `json:"QueryOptimize,omitempty" name:"QueryOptimize"`
}

func (r *DescribeLogHistogramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogHistogramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformFailLogInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据加工任务失败日志详请

		LogFailureInfos []*DataTransformFailureInfo `json:"LogFailureInfos,omitempty" name:"LogFailureInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformFailLogInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformFailLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestJsonLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json格式日志内容

		LogData *JsonLogInfo `json:"LogData,omitempty" name:"LogData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLatestJsonLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestJsonLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenBeginRegexRequest struct {
	*tchttp.BaseRequest

	// 日志样例

	LogSample *string `json:"LogSample,omitempty" name:"LogSample"`
}

func (r *GenBeginRegexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenBeginRegexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceLogConfigInfo struct {

	// 服务日志的logset信息

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 服务日志的Topic ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 服务日志的Topic Name

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，并且不能有重复的键值对。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 该日志主题是否开始采集

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 是否开启自动分裂

	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`
	// 若开启最大分裂，该主题能够能够允许的最大分区数

	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`
	// 生命周期，单位天；可取值范围1~366

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 用户自定义抽样配置

	UserSample *string `json:"UserSample,omitempty" name:"UserSample"`
	// 用户抽样配置开关：默认关闭

	UserSampleStatus *bool `json:"UserSampleStatus,omitempty" name:"UserSampleStatus"`
	// 存储类型：low 低频

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 日志主题描述

	Describes *string `json:"Describes,omitempty" name:"Describes"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataTransformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDataTransformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDemonstrationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDemonstrationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDemonstrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShippersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递规则列表

		Shippers []*ShipperInfo `json:"Shippers,omitempty" name:"Shippers"`
		// 本次查询获取到的总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShippersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShippersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleKeyValueInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 需要建立索引的键值对信息；最大只能配置100个键值对

	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
	// 索引是否开启动态模板；若开启，则会根据上报的键值对配置索引，但是所有字段类型都是text，大小写敏感，不支持分析，采用默认分词符

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	TopicInfo
}

func (r *CreateTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigExtraResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigExtraResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsyncContextTask struct {

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 创建时间，时间戳，精确到毫秒

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 状态，0表示待开始，1表示运行中，2表示已完成，-1表示失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 异步上下文任务ID

	AsyncContextTaskId *string `json:"AsyncContextTaskId,omitempty" name:"AsyncContextTaskId"`
	// 任务失败的错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	// 日志包序号

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 日志包内一条日志的序号

	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 日志时间

	Time *int64 `json:"Time,omitempty" name:"Time"`
	// 任务完成时间，时间戳，精确到毫秒

	FinishTime *int64 `json:"FinishTime,omitempty" name:"FinishTime"`
	// 相关联的异步检索ID

	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`
	// 相关联的异步检索任务的查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 相关联的异步检索任务的查询起始时间

	From *int64 `json:"From,omitempty" name:"From"`
	// 相关联的异步检索任务的查询结束时间

	To *int64 `json:"To,omitempty" name:"To"`
}

type ExcludePathInfo struct {

	// 类型，选填File或Path

	Type *string `json:"Type,omitempty" name:"Type"`
	// Type对应的具体内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeAsyncSearchTasksRequest struct {
	*tchttp.BaseRequest

	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// <br><li> topicId
	//
	// 按照【日志主题ID】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> logsetId
	//
	// 按照【日志集ID】进行过滤，可通过调用DescribeLogsets查询已创建的日志集列表或登录控制台进行查看；也可以调用CreateLogset创建新的日志集。
	//
	// 类型：String
	//
	// 必选：否
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为5

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAsyncSearchTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncSearchTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformPreviewDataInfoRequest struct {
	*tchttp.BaseRequest

	// 任务id， 获取加工后数据有效

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 1: 获取原始数据。   2: 获取加工后数据

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 源日志主题id,  获取加工前数据有效

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeDataTransformPreviewDataInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformPreviewDataInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostFileInfo struct {

	// 日志文件夹

	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
	// 日志文件名

	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`
	// metadata信息

	CustomLabels []*string `json:"CustomLabels,omitempty" name:"CustomLabels"`
}

type CancelRebuildIndexTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelRebuildIndexTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelRebuildIndexTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogFastAnalysisRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 字段名

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 查询语句，语句长度最大为4096

	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeLogFastAnalysisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogFastAnalysisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MergePartitionRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 合并的PartitionId

	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

func (r *MergePartitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MergePartitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRebuildIndexTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 索引重建任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRebuildIndexTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRebuildIndexTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分区列表

		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePartitionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HeartBeatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// agent升级任务信息

		UpdateInfo *AgentUpdateInfo `json:"UpdateInfo,omitempty" name:"UpdateInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HeartBeatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HeartBeatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenKafkaConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待消费TopicId

		TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenKafkaConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集日志配置

		LogConfigs []*LogConfigInfo `json:"LogConfigs,omitempty" name:"LogConfigs"`
		// 弃用

		NeedUpdate *bool `json:"NeedUpdate,omitempty" name:"NeedUpdate"`
		// 弃用

		LastVersion *string `json:"LastVersion,omitempty" name:"LastVersion"`
		// 弃用

		URL *string `json:"URL,omitempty" name:"URL"`
		// 弃用

		FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
		// 服务日志的配置信息

		ServiceLogConfigs *ServiceLogConfigInfo `json:"ServiceLogConfigs,omitempty" name:"ServiceLogConfigs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据加工任务列表信息

		DataTransformTaskInfos []*DataTransformTaskInfo `json:"DataTransformTaskInfos,omitempty" name:"DataTransformTaskInfos"`
		// 任务总次数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeliverCloudFunctionRequest struct {
	*tchttp.BaseRequest

	// 投递规则属于的 topic id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递的云函数名字

	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 函数版本

	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`
	// 投递最长等待时间，单位：秒

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 投递最大消息数

	MaxMsgNum *uint64 `json:"MaxMsgNum,omitempty" name:"MaxMsgNum"`
}

func (r *CreateDeliverCloudFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeliverCloudFunctionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionInfo struct {

	// 函数名称

	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`
	// 函数描述

	FuncNameDesc *string `json:"FuncNameDesc,omitempty" name:"FuncNameDesc"`
	// 用来说明函数功能

	FuncUseDesc *string `json:"FuncUseDesc,omitempty" name:"FuncUseDesc"`
	// 语法描述

	FuncSyntaxDesc *string `json:"FuncSyntaxDesc,omitempty" name:"FuncSyntaxDesc"`
	// demo展示

	FuncDemo *string `json:"FuncDemo,omitempty" name:"FuncDemo"`
	// 函数类型

	FuncType *string `json:"FuncType,omitempty" name:"FuncType"`
	// true是可变长度参数的加工函数

	IsVariadic *bool `json:"IsVariadic,omitempty" name:"IsVariadic"`
	// 可变参的参数个数最大限制

	MaxArgumentSize *int64 `json:"MaxArgumentSize,omitempty" name:"MaxArgumentSize"`
	// 函数返回结果类型，用来校验嵌套函数中，返回结果是否和函数参数类型匹配。
	//
	// 不同的func_type返回的对象类型不同 string/int/bool/condition/func

	ReturnType *string `json:"ReturnType,omitempty" name:"ReturnType"`
	// 函数参数描述

	Arguments []*FunctionArgument `json:"Arguments,omitempty" name:"Arguments"`
}

type LogsetInfo struct {

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 若AssumerUin非空，则表示创建该日志集的服务方Uin

	AssumerUin *uint64 `json:"AssumerUin,omitempty" name:"AssumerUin"`
	// 若AssumerUin非空，则表示创建该日志集的服务方名称

	AssumerName *string `json:"AssumerName,omitempty" name:"AssumerName"`
	// 若AssumerUin非空，则表示非改服务方的调用者对于日志集的修改权限

	LogsetModifyAcl *int64 `json:"LogsetModifyAcl,omitempty" name:"LogsetModifyAcl"`
	// 日志集绑定的标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 日志集下日志主题的数目

	TopicCount *int64 `json:"TopicCount,omitempty" name:"TopicCount"`
	// 若AssumerUin非空，则表示创建该日志集的服务方角色

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 生命周期，单位为天

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 若AssumerUin非空，则表示非改服务方的调用者对于日志集的删除权限

	LogsetDelACL *int64 `json:"LogsetDelACL,omitempty" name:"LogsetDelACL"`
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 通知渠道组名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 通知类型。可选值： <li> Trigger - 告警触发 <li> Recovery - 告警恢复 <li> All - 告警触发和告警恢复

	Type *string `json:"Type,omitempty" name:"Type"`
	// 通知接收对象。

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 接口回调信息（包括企业微信）。

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncContextTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志时间，单位ms

	Time *int64 `json:"Time,omitempty" name:"Time"`
	// 日志包序号

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 日志包内一条日志的序号

	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 异步检索任务ID

	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`
}

func (r *CreateAsyncContextTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAsyncContextTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicExtendConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicExtendConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicExtendConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDemonstrationResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDemonstrationResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDemonstrationResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKafkaConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyKafkaConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {

	// 全文索引配置

	FullText *FullTextInfo `json:"FullText,omitempty" name:"FullText"`
	// 键值索引配置

	KeyValue *RuleKeyValueInfo `json:"KeyValue,omitempty" name:"KeyValue"`
	// 元字段索引配置

	Tag *RuleTagInfo `json:"Tag,omitempty" name:"Tag"`

	DynamicIndex *DynamicIndex `json:"DynamicIndex,omitempty" name:"DynamicIndex"`
}

type CheckAlarmChannelRequest struct {
	*tchttp.BaseRequest

	// 告警通知接收者数组

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 回调数组

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
	// 通知类型（Trigger,Recovery,All）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckAlarmChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAlarmChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDemonstrationsRequest struct {
	*tchttp.BaseRequest

	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// DemonstrationId

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDemonstrationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDemonstrationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmAnalysisConfig struct {

	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogInfo struct {

	// 日志时间，单位ms

	Time *int64 `json:"Time,omitempty" name:"Time"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 日志来源IP

	Source *string `json:"Source,omitempty" name:"Source"`
	// 日志文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 日志上报请求包的ID

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 请求包内日志的ID

	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 日志内容，由多个LogItem (KV结构）组成

	Logs []*LogItem `json:"Logs,omitempty" name:"Logs"`
	// 日志内容的高亮描述信息

	HighLights []*HighLightItem `json:"HighLights,omitempty" name:"HighLights"`
	// 日志内容的Json序列化字符串

	LogJson *string `json:"LogJson,omitempty" name:"LogJson"`
	// 日志来源主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

type CancelRebuildIndexTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 索引重建任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CancelRebuildIndexTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelRebuildIndexTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShipperRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
}

func (r *DeleteShipperRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteShipperRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集配置列表

		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`
		// 过滤到的总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模版数组

		Templates []*Template `json:"Templates,omitempty" name:"Templates"`
		// 演示示例地域

		DemonstrationRegion *string `json:"DemonstrationRegion,omitempty" name:"DemonstrationRegion"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeliverCloudFunctionRequest struct {
	*tchttp.BaseRequest

	// 投递规则属于的 topic id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteDeliverCloudFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeliverCloudFunctionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerStdoutInfo struct {

	// 是否所有容器

	AllContainers *bool `json:"AllContainers,omitempty" name:"AllContainers"`
	// container为空表所有的，不为空采集指定的容器

	Container *string `json:"Container,omitempty" name:"Container"`
	// namespace可以多个，用分隔号分割,例如A,B；为空或者没有这个字段，表示所有namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// pod标签信息

	IncludeLabels []*string `json:"IncludeLabels,omitempty" name:"IncludeLabels"`
	// 工作负载信息

	WorkLoads []*ContainerWorkLoadInfo `json:"WorkLoads,omitempty" name:"WorkLoads"`
	// 需要排除的namespace可以多个，用分隔号分割,例如A,B

	ExcludeNamespace *string `json:"ExcludeNamespace,omitempty" name:"ExcludeNamespace"`
	// 需要排除的pod标签信息

	ExcludeLabels []*string `json:"ExcludeLabels,omitempty" name:"ExcludeLabels"`
}

type CreateShipperResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递规则ID

		ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateShipperResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// kafka消费用户名

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompressInfo struct {

	// 压缩格式，支持gzip、lzop和none不压缩

	Format *string `json:"Format,omitempty" name:"Format"`
}

type CloseKafkaConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseKafkaConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserConfigRequest struct {
	*tchttp.BaseRequest

	// 要查询配置key列表

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
}

func (r *DescribeUserConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInfo struct {

	// 告警策略名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 监控对象列表。

	AlarmTargets []*AlarmTargetInfo `json:"AlarmTargets,omitempty" name:"AlarmTargets"`
	// 监控任务运行时间点。

	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`
	// 触发条件。

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 告警重复的周期。单位是min。取值范围是0~1440。

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 关联的告警通知模板列表。

	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`
	// 开启状态。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 告警策略ID。

	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 开启状态

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 自定义通知模板

	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`
	// 自定义回调模板

	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`
	// 多维分析设置

	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
}

type DescribeKafkaConsumeRequest struct {
	*tchttp.BaseRequest

	// CLS对应topic标识

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *DescribeKafkaConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardInfo struct {

	// 仪表盘id

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
	// 仪表盘名字

	DashboardName *string `json:"DashboardName,omitempty" name:"DashboardName"`
	// 仪表盘数据

	Data *string `json:"Data,omitempty" name:"Data"`
	// 创建仪表盘的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// AssumerUin非空则表示创建该日志主题的服务方Uin

	AssumerUin *uint64 `json:"AssumerUin,omitempty" name:"AssumerUin"`
	// RoleName非空则表示创建该日志主题的服务方使用的角色

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// AssumerName非空则表示创建该日志主题的服务方名称

	AssumerName *string `json:"AssumerName,omitempty" name:"AssumerName"`
	// 日志主题绑定的标签信息

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 仪表盘所在地域： 为了兼容老的地域。

	DashboardRegion *string `json:"DashboardRegion,omitempty" name:"DashboardRegion"`
	// 修改仪表盘的时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 仪表盘对应的topic相关信息

	DashboardTopicInfos []*DashboardTopicInfo `json:"DashboardTopicInfos,omitempty" name:"DashboardTopicInfos"`
}

type DataTransformTaskInfo struct {

	// 数据加工任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务启用状态，默认为1，正常开启,  2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 加工任务类型，1： DSL， 2：SQL

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 源日志主题

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 当前加工任务状态（1准备中/2运行中/3停止中/4已停止）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 加工任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 最后启用时间，如果需要重建集群，修改该时间

	LastEnableTime *string `json:"LastEnableTime,omitempty" name:"LastEnableTime"`
	// 日志主题名称

	SrcTopicName *string `json:"SrcTopicName,omitempty" name:"SrcTopicName"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 加工任务目的topic_id以及别名

	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`
	// 加工逻辑函数

	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`
}

type DescribeConfigExtrasRequest struct {
	*tchttp.BaseRequest

	// 支持的key： topicId,name, configExtraId, machineGroupId

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigExtrasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigExtrasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigMachineGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集规则配置绑定的机器组列表

		MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitempty" name:"MachineGroups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigMachineGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenKafkaConsumeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待消费kafka信息

		Kafka *KafkaInfo `json:"Kafka,omitempty" name:"Kafka"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenKafkaConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKafkaConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserConfigRequest struct {
	*tchttp.BaseRequest

	// 要删除key列表

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
}

func (r *DeleteUserConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeAgentNormalRequest struct {
	*tchttp.BaseRequest

	// 需要升级的机器IP列表

	AgentIps []*string `json:"AgentIps,omitempty" name:"AgentIps"`
	// 升级类型:0-disable,1-manual,2-auto

	UpdateMode *int64 `json:"UpdateMode,omitempty" name:"UpdateMode"`
	// 升级开始时间，如：22:00:00，晚上10点开始

	UpdateStart *string `json:"UpdateStart,omitempty" name:"UpdateStart"`
	// 升级结束时间，如：23:00:00，晚上11点结束

	UpdateStop *string `json:"UpdateStop,omitempty" name:"UpdateStop"`
	// 升级目标版本

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
}

func (r *UpgradeAgentNormalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeAgentNormalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportInfo struct {

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志导出任务ID

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
	// 日志导出查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 日志导出文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 日志文件大小

	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
	// 日志导出时间排序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 日志导出格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 日志导出数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 日志下载状态。Processing:导出正在进行中，Complete:导出完成，Failed:导出失败，Expired:日志导出已过期（三天有效期）。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 日志导出起始时间

	From *int64 `json:"From,omitempty" name:"From"`
	// 日志导出结束时间

	To *int64 `json:"To,omitempty" name:"To"`
	// 日志导出路径

	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`
	// 日志导出创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略ID。

		AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分页的总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志集列表

		Logsets []*LogsetInfo `json:"Logsets,omitempty" name:"Logsets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID。

	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
	// 告警策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 监控任务运行时间点。

	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`
	// 触发条件。

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 告警重复的周期。单位是分钟。取值范围是0~1440。

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 关联的告警通知模板列表。

	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`
	// 监控对象列表。

	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`
	// 是否开启告警策略。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 是否开启告警策略。默认值为true

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 用户自定义告警内容

	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`
	// 用户自定义回调

	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`
	// 多维分析

	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
}

func (r *ModifyAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest

	// <br><li> topicName
	//
	// 按照【日志主题名称】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> topicId
	//
	// 按照【日志主题ID】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> logsetId
	//
	// 按照【日志集ID】进行过滤，可通过调用DescribeLogsets查询已创建的日志集列表或登录控制台进行查看；也可以调用CreateLogset创建新的日志集。
	//
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> tagKey
	//
	// 按照【标签键】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> tag:tagKey
	//
	// 按照【标签键值对】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例2。
	//
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> storageType
	//
	// 按照【日志主题的存储类型】进行过滤。可选值 hot（实时存储），cold（离线存储）
	// 类型：String
	//
	// 必选：否
	//
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 控制filter相关字段是否精确查询。默认（0）模糊查询, 1: topicName精确查询, 2:logsetName精确查询, 3：topicName和logsetName都精确查询

	PreciseSearch *uint64 `json:"PreciseSearch,omitempty" name:"PreciseSearch"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldValueRatioInfos struct {

	// 字段值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 字段值所占的数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 字段值所占的比例

	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
}

type KeyLogInfo struct {

	// 日志key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 日志内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type TopicIdAndRegion struct {

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题id

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源列表

		Resources []*ResourcesInfo `json:"Resources,omitempty" name:"Resources"`
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

type DescribeShippersRequest struct {
	*tchttp.BaseRequest

	// <br><li> shipperName
	//
	// 按照【投递规则名称】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> shipperId
	//
	// 按照【投递规则ID】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> topicId
	//
	// 按照【日志主题】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeShippersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShippersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionArgument struct {

	// 参数序号，根据参数顺序定义

	ArgIndex *int64 `json:"ArgIndex,omitempty" name:"ArgIndex"`
	// 参数名称

	ArgName *string `json:"ArgName,omitempty" name:"ArgName"`
	// 参数描述

	ArgDesc *string `json:"ArgDesc,omitempty" name:"ArgDesc"`
	// 可接受的参数类型列表，包括字面常量、数组、条件表达式、函数表达式等任意一种或多种

	ArgType *string `json:"ArgType,omitempty" name:"ArgType"`
	// 参数默认值

	ArgValueDefault *string `json:"ArgValueDefault,omitempty" name:"ArgValueDefault"`
	// 范围、枚举类型

	ArgValueType *string `json:"ArgValueType,omitempty" name:"ArgValueType"`
	// 参数值域校验范围，这里仅针对常量进行校验，如果arg_value_type是scope类型，则此数组表示前闭后开区间，否则表示枚举的值类型.
	//
	// 如果此值为空，或者空数组，则不进行参数值校验

	ArgValueScope []*string `json:"ArgValueScope,omitempty" name:"ArgValueScope"`
	// 是否必须

	IsNecessary *bool `json:"IsNecessary,omitempty" name:"IsNecessary"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest

	// 获取数据地域类型，all:全地域，ap-region指定地域

	DataRegion *string `json:"DataRegion,omitempty" name:"DataRegion"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentUpdateStatus struct {

	// 升级类型:0-null,1-manual,2-auto,3-force

	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 升级动作:0-null,1-update,2-revert

	UpdateAction *int64 `json:"UpdateAction,omitempty" name:"UpdateAction"`
	// 重试次数,最大三次

	RetryCount *int64 `json:"RetryCount,omitempty" name:"RetryCount"`
	// Agent升级状态:0-querying,1-updating,2-reverting,-1-updatefail,-2-revertfail, -10-notsupport

	UpdateStatus *int64 `json:"UpdateStatus,omitempty" name:"UpdateStatus"`
	// 错误码

	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest

	// <br><li> configName
	//
	// 按照【采集配置名称】进行模糊匹配过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> configId
	//
	// 按照【采集配置ID】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> topicId
	//
	// 按照【日志主题】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegexIndexInfo struct {

	// 起始位置

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 结束位置

	End *int64 `json:"End,omitempty" name:"End"`
}

type DescribeAsyncSearchResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加载后续内容的Context

		Context *string `json:"Context,omitempty" name:"Context"`
		// 日志查询结果是否全部返回

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 日志内容

		Results []*LogInfo `json:"Results,omitempty" name:"Results"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncSearchResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncSearchResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlarmLogRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 查询语句，语句长度最大为1024

	Query *string `json:"Query,omitempty" name:"Query"`
	// 单次查询返回的日志条数，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 加载更多日志时使用，透传上次返回的Context值，获取后续的日志内容

	Context *string `json:"Context,omitempty" name:"Context"`
	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 为true代表使用新检索,响应参数AnalysisRecords和Columns有效， 为false时代表使用老检索方式, AnalysisResults和ColNames有效

	UseNewAnalysis *bool `json:"UseNewAnalysis,omitempty" name:"UseNewAnalysis"`
}

func (r *GetAlarmLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlarmLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaTagInfo struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ExtractRuleInfo struct {

	// 时间字段的key名字，time_key和time_format必须成对出现

	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`
	// 时间字段的格式，参考c语言的strftime函数对于时间的格式说明输出参数

	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
	// 分隔符类型日志的分隔符，只有log_type为delimiter_log时有效

	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`
	// 整条日志匹配规则，只有log_type为fullregex_log时有效

	LogRegex *string `json:"LogRegex,omitempty" name:"LogRegex"`
	// 行首匹配规则，只有log_type为multiline_log或fullregex_log时有效

	BeginRegex *string `json:"BeginRegex,omitempty" name:"BeginRegex"`
	// 取的每个字段的key名字，为空的key代表丢弃这个字段，只有log_type为delimiter_log时有效，json_log的日志使用json本身的key

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
	// 需要过滤日志的key，及其对应的regex

	FilterKeyRegex []*KeyRegexInfo `json:"FilterKeyRegex,omitempty" name:"FilterKeyRegex"`
	// 解析失败日志是否上传，true表示上传，false表示不上传

	UnMatchUpLoadSwitch *bool `json:"UnMatchUpLoadSwitch,omitempty" name:"UnMatchUpLoadSwitch"`
	// 失败日志的key

	UnMatchLogKey *string `json:"UnMatchLogKey,omitempty" name:"UnMatchLogKey"`
	// 增量采集模式下的回溯数据量，默认-1（全量采集）

	Backtracking *int64 `json:"Backtracking,omitempty" name:"Backtracking"`
	// 是否为Gbk编码.   0: 否, 1: 是

	IsGBK *int64 `json:"IsGBK,omitempty" name:"IsGBK"`
	// 是否为标准json.   0: 否, 1: 是

	JsonStandard *int64 `json:"JsonStandard,omitempty" name:"JsonStandard"`

	// syslog协议，tcp或者udp
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// syslog系统日志采集指定采集器监听的地址和端口
	Address *string `json:"Address,omitempty" name:"Address"`

	// rfc3164或者rfc5424
	ParseProtocol *string `json:"ParseProtocol,omitempty" name:"ParseProtocol"`

	// 元数据类型，0: 不使用元数据，1:使用机器组元数据，2:使用用户自定义元数据，3:使用采集路径规定元数据
	MetadataType *int64 `json:"MetadataType,omitempty" name:"MetadataType"`

	// 采集路径正则表达式
	PathRegex *string `json:"PathRegex,omitempty" name:"PathRegex"`

	// 元数据信息
	MetaTags []*MetaTagInfo `json:"MetaTags,omitempty" name:"MetaTags"`

	//
	EventLogRules []*EventLogRuleInfo `json:"EventLogRules,omitempty" name:"EventLogRules"`
}

type EventLogRuleInfo struct {

	// 事件通道，支持Application，Security，Setup，System，ALL
	EventChannel *string `json:"EventChannel,omitempty" name:"EventChannel"`

	// 时间类型，1:用户自定义，2:当前时间
	TimeType *uint64 `json:"TimeType,omitempty" name:"TimeType"`

	// 时间，用户选择自定义时间类型时，需要指定时间
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 事件ID过滤列表  选填，为空表示不做过滤 支持正向过滤单个值（例：20）或范围（例：0-20），也支持反向过滤单个值(例：-20) 多个过滤项之间可由逗号隔开，例：1-200,-100表示采集1-200范围内除了100以外的事件日志
	EventIDs []*string `json:"EventIDs,omitempty" name:"EventIDs"`
}

type KeyValueInfo struct {

	// 需要配置键值或者元字段索引的字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 字段的索引描述信息

	Value *ValueInfo `json:"Value,omitempty" name:"Value"`
}

type DescribeDataTransformProcessInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据加工任务进度详请

		TaskLogStatistics []*TaskLogStatistic `json:"TaskLogStatistics,omitempty" name:"TaskLogStatistics"`
		// 读取的源日志主题的总行数

		ReadLineSum *uint64 `json:"ReadLineSum,omitempty" name:"ReadLineSum"`
		// 加工后的总行数

		WriteLineSum *uint64 `json:"WriteLineSum,omitempty" name:"WriteLineSum"`
		// 加工失败的总行数

		FailedLineSum *uint64 `json:"FailedLineSum,omitempty" name:"FailedLineSum"`
		// 加工过滤的总行数

		FilterLineSum *uint64 `json:"FilterLineSum,omitempty" name:"FilterLineSum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformProcessInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformProcessInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenBeginRegexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正则表达式

		Regex *string `json:"Regex,omitempty" name:"Regex"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenBeginRegexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenBeginRegexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardRequest struct {
	*tchttp.BaseRequest

	// 仪表盘名称

	DashboardName *string `json:"DashboardName,omitempty" name:"DashboardName"`
	// 仪表盘配置数据

	Data *string `json:"Data,omitempty" name:"Data"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDemonstrationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDemonstrationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDemonstrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmTargetInfo struct {

	// 日志集ID。

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集名称。

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 查询语句。

	Query *string `json:"Query,omitempty" name:"Query"`
	// 告警对象序号。

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 查询范围起始时间相对当前的历史时间，取值为非正，最大值为0，最小值为-1440。

	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`
	// 查询范围终止时间相对当前的历史时间，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。

	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type JsonInfo struct {

	// 启用标志

	EnableTag *bool `json:"EnableTag,omitempty" name:"EnableTag"`
	// 元数据信息列表

	MetaFields []*string `json:"MetaFields,omitempty" name:"MetaFields"`

	//
	JsonType *int64 `json:"JsonType,omitempty" name:"JsonType"`
}

type AddMachineGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMachineGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMachineGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCosRechargesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 见: CosRechargeInfo 结构描述

		Data []*CosRechargeInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCosRechargesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosRechargesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘的数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 仪表盘详细明细

		DashboardInfos []*DashboardInfo `json:"DashboardInfos,omitempty" name:"DashboardInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubTypeItem struct {

	// 演示示例子类型

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 检索语句模版

	Queries []*QueryTemplateItem `json:"Queries,omitempty" name:"Queries"`
	// 模版项

	TemplateItems []*TemplateItem `json:"TemplateItems,omitempty" name:"TemplateItems"`
}

type ApplyConfigToMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 采集配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ApplyConfigToMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyConfigToMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaConsumeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Kafka消费信息

		Kafka *KafkaInfo `json:"Kafka,omitempty" name:"Kafka"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribePartitionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePartitionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigExtraResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigExtraResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 监控对象列表。

	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`
	// 监控任务运行时间点。

	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`
	// 触发条件。

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 告警重复的周期。单位是分钟。取值范围是0~1440。

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 关联的告警通知模板列表。

	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`
	// 是否开启告警策略。默认值为true

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 是否开启告警策略。默认值为true

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 用户自定义告警内容

	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`
	// 用户自定义回调

	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`
	// 多维分析

	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
}

func (r *CreateAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 机器组类型

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
	// 标签列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否开启机器组自动更新

	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 升级开始时间，建议业务低峰期升级LogListener

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 升级结束时间，建议业务低峰期升级LogListener

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费

	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
	// 默认值""，"label_k8s"

	// 机器组中机器定期离线清理时间
	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitempty" name:"DelayCleanupTime"`

	// 机器组元数据信息列表
	MetaTags []*MetaTagInfo `json:"MetaTags,omitempty" name:"MetaTags"`

	Flag *string `json:"Flag,omitempty" name:"Flag"`
}

func (r *ModifyMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckFunctionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败错误码

		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
		// 失败错误信息

		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RealtimeProducerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RealtimeProducerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RealtimeProducerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmTarget struct {

	// 日志主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 查询语句。

	Query *string `json:"Query,omitempty" name:"Query"`
	// 告警对象序号；从1开始递增。

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 查询范围起始时间相对当前的历史时间，单位非分钟，取值为非正，最大值为0，最小值为-1440。

	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`
	// 查询范围终止时间相对当前的历史时间，单位非分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。

	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
	// 日志集ID。

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

type HistogramInfo struct {

	// 统计周期内的日志条数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 按 period 取整后的 unix timestamp： 单位毫秒

	BTime *int64 `json:"BTime,omitempty" name:"BTime"`
}

type DescribeDataTransformInfoRequest struct {
	*tchttp.BaseRequest

	// <br><li> taskName
	//
	// 按照【加工任务名称】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> taskId
	//
	// 按照【加工任务id】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 默认值为2.   1: 获取单个任务的详细信息 2：获取任务列表

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// Type为1， 此参数必填

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDataTransformInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MergePartitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 合并结果集

		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MergePartitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MergePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionTestResult struct {

	// 错误码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type CreateDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘id

		DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDemonstrationRequest struct {
	*tchttp.BaseRequest

	// 演示示例类型：'CLB'

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateDemonstrationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDemonstrationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigRequest struct {
	*tchttp.BaseRequest

	// 采集规则配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集规则配置列表

		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGroupConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGroupConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCosRechargeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的存储桶下的某个文件的前几行数据

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 匹配到的存储桶下的文件个数

		Sum *uint64 `json:"Sum,omitempty" name:"Sum"`
		// 当前预览文件路径

		Path *string `json:"Path,omitempty" name:"Path"`
		// 预览获取数据失败原因

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchCosRechargeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCosRechargeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMachineGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组类型 目前type支持 ip 和 label

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
}

func (r *AddMachineGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMachineGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigExtraRequest struct {
	*tchttp.BaseRequest

	// 采集规则扩展配置ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
}

func (r *DeleteConfigExtraRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigExtraRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAsyncContextTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 异步上下文任务ID

	AsyncContextTaskId *string `json:"AsyncContextTaskId,omitempty" name:"AsyncContextTaskId"`
}

func (r *DeleteAsyncContextTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAsyncContextTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteLogsetRequest struct {
	*tchttp.BaseRequest

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *DeleteLogsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShipperResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteShipperResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentConfigsRequest struct {
	*tchttp.BaseRequest

	// agent的版本号

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// 标签列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// agent的instance id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAgentConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportRequest struct {
	*tchttp.BaseRequest

	// 日志导出ID

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
}

func (r *DeleteExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerResponse struct {
	*tchttp.BaseResponse

	Response *DescribeConsumerResponseParams `json:"Response"`
}

type DescribeConsumerResponseParams struct {
	// 投递任务是否生效

	Effective *bool `json:"Effective,omitempty" name:"Effective"`
	// 是否投递日志的元数据信息

	NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`
	// 如果需要投递元数据信息，元数据信息的描述

	Content *ConsumerContent `json:"Content,omitempty" name:"Content"`
	// CKafka的描述

	Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`

	Compression *int `json:"Compression,omitempty" name:"Compression"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

func (r *DescribeConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardsRequest struct {
	*tchttp.BaseRequest

	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// <br><li> dashboardId
	//
	// 按照【仪表盘id】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> dashboardName
	//
	// 按照【仪表盘名字】进行模糊搜索过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> dashboardRegion
	//
	// 按照【仪表盘地域】进行过滤，为了兼容老的仪表盘，通过云API创建的仪表盘没有地域属性
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> tagKey
	//
	// 按照【标签键】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> tag:tagKey
	//
	// 按照【标签键值对】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例2。
	//
	// 类型：String
	//
	// 必选：否
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 按照topicId和regionId过滤。

	TopicIdRegionFilter []*TopicIdAndRegion `json:"TopicIdRegionFilter,omitempty" name:"TopicIdRegionFilter"`
}

func (r *DescribeDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryShipperTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryShipperTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryShipperTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosRechargeInfo struct {

	// 主键ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// cos导入任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// cos存储桶

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// cos存储桶地域

	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`
	// cos存储桶前缀地址

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志； 默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 状态   status 0: created, 1: running, 2: pause, 3: finished, 4: failed。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 是否启用:   0： 未启用  ， 1：启用

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 进度条百分值

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// supported: "", "gzip", "lzop", "snappy”; 默认空

	Compress *string `json:"Compress,omitempty" name:"Compress"`
	// 见： ExtractRuleInfo 结构描述

	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitempty" name:"ExtractRuleInfo"`
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 是否生效，默认为true

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 索引规则

	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
	// 全文索引系统预置字段标记，默认false。  false:不包含系统预置字段， true:包含系统预置字段

	IncludeInternalFields *bool `json:"IncludeInternalFields,omitempty" name:"IncludeInternalFields"`
	// 元数据相关标志位，默认为0。 0：全文索引包括开启键值索引的元数据字段， 1：全文索引包括所有元数据字段，2：全文索引不包括元数据字段。

	MetadataFlag *uint64 `json:"MetadataFlag,omitempty" name:"MetadataFlag"`
	// 自定义日志解析异常存储字段。

	CoverageField *string `json:"CoverageField,omitempty" name:"CoverageField"`
}

func (r *CreateIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组类型 目前type支持 ip 和 label

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
}

func (r *DeleteMachineGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerFileInfo struct {

	// namespace可以多个，用分隔号分割,例如A,B

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 容器名称

	Container *string `json:"Container,omitempty" name:"Container"`
	// 日志文件夹

	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
	// 日志名称

	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`
	// pod标签信息

	IncludeLabels []*string `json:"IncludeLabels,omitempty" name:"IncludeLabels"`
	// 工作负载信息

	WorkLoad *ContainerWorkLoadInfo `json:"WorkLoad,omitempty" name:"WorkLoad"`
	// 需要排除的namespace可以多个，用分隔号分割,例如A,B

	ExcludeNamespace *string `json:"ExcludeNamespace,omitempty" name:"ExcludeNamespace"`
	// 需要排除的pod标签信息

	ExcludeLabels []*string `json:"ExcludeLabels,omitempty" name:"ExcludeLabels"`
}

type RuleTagInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 标签索引配置中的字段信息

	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题ID

		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebuildIndexTaskInfo struct {

	// 索引重建任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 索引重建任务当前状态，0:索引重建任务已创建，1:创建索引重建资源，2:索引重建资源创建完成，3:重建中，4:暂停，5:重建索引成功，6:重建成功（可检索），7:重建失败，8:撤销，9:删除元数据和索引

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 重建任务开始时间戳

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 重建任务结束时间戳

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 重投预估剩余时间，单位秒

	RemainTime *int64 `json:"RemainTime,omitempty" name:"RemainTime"`
	// 重建任务创建时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 重投完成度，百分比

	Progress *float64 `json:"Progress,omitempty" name:"Progress"`
	// 重建任务更新时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 附加状态描述信息（目前仅描述失败时失败原因）

	StatusMessage *string `json:"StatusMessage,omitempty" name:"StatusMessage"`
}

type DeleteTopicExtendConfigRequest struct {
	*tchttp.BaseRequest

	// clb的业务配置

	LbKeys []*string `json:"LbKeys,omitempty" name:"LbKeys"`
}

func (r *DeleteTopicExtendConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicExtendConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAsyncSearchTaskRequest struct {
	*tchttp.BaseRequest

	// 异步检索任务ID

	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteAsyncSearchTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAsyncSearchTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterRuleInfo struct {

	// 过滤规则Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 过滤规则

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 过滤规则Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateDeliverCloudFunctionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeliverCloudFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeliverCloudFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataTransformResouceInfo struct {

	// 目标主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 主账号Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type DescribeDemonstrationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 演示示例数组

		Demonstrations []*Demonstration `json:"Demonstrations,omitempty" name:"Demonstrations"`
		// 符合条件的演示示例数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDemonstrationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDemonstrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckFunctionRequest struct {
	*tchttp.BaseRequest

	// 用户输入的加工语句

	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`
	// 加工任务目的topic_id以及别名

	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`
}

func (r *CheckFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFunctionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 告警通知模板

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
}

func (r *DeleteAlarmNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsRequest struct {
	*tchttp.BaseRequest

	// <br><li> name
	//
	// 按照【告警策略名称】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> alarmId
	//
	// 按照【告警策略ID】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> topicId
	//
	// 按照【监控对象的日志主题ID】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> enable
	//
	// 按照【启用状态】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogFastAnalysisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 字段取值的个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 字段取值的占比信息

		FieldValueRatioInfos []*FieldValueRatioInfos `json:"FieldValueRatioInfos,omitempty" name:"FieldValueRatioInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogFastAnalysisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogFastAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicExtendConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicExtendConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicExtendConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveMachineRequest struct {
	*tchttp.BaseRequest

	// 机器组 ID。

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 剔除的Ip数组。

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

func (r *RemoveMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShipperTaskInfo struct {

	// 投递任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 投递信息ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 本批投递的日志的开始时间戳，毫秒

	RangeStart *int64 `json:"RangeStart,omitempty" name:"RangeStart"`
	// 本批投递的日志的结束时间戳， 毫秒

	RangeEnd *int64 `json:"RangeEnd,omitempty" name:"RangeEnd"`
	// 本次投递任务的开始时间戳， 毫秒

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 本次投递任务的结束时间戳， 毫秒

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 本次投递的结果，"success","running","failed"

	Status *string `json:"Status,omitempty" name:"Status"`
	// 结果的详细信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DescribeCosRechargesRequest struct {
	*tchttp.BaseRequest

	// 日志主题 ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 状态   status 0: created, 1: running, 2: pause, 3: finished, 4: failed。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 是否启用:   0： 未启用  ， 1：启用

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *DescribeCosRechargesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosRechargesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadLogRequest struct {
	*tchttp.BaseRequest

	// 根据 hashkey 写入相应范围的主题分区

	HashKey *string `json:"HashKey,omitempty" name:"HashKey"`
	// 主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 压缩方法

	CompressType *string `json:"CompressType,omitempty" name:"CompressType"`
	// Agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// Agent请求的序列号

	AgentSeq *string `json:"AgentSeq,omitempty" name:"AgentSeq"`
	// Agent版本号

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// Agent请求的Unique Id

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
}

func (r *UploadLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShipperInfo struct {

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递的bucket地址

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 投递的前缀目录

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 投递规则的名字

	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`
	// 投递的时间间隔，单位 秒

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 投递的文件的最大值，单位 MB

	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`
	// 是否生效

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 投递日志的过滤规则

	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`
	// 投递日志的分区规则，支持strftime的时间格式表示

	Partition *string `json:"Partition,omitempty" name:"Partition"`
	// 投递日志的压缩配置

	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`
	// 投递日志的内容格式配置

	Content *ContentInfo `json:"Content,omitempty" name:"Content"`
	// 投递日志的创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 投递文件命名配置，0：随机数命名，1：投递时间命名，默认0（随机数命名）
	FilenameMode *uint64 `json:"FilenameMode,omitempty" name:"FilenameMode"`

	// 跨账户uin，用于支持跨账户bucket投递
	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`

	// 投递数据范围的开始时间点，不能超出日志主题的生命周期起点。如果用户不填写，默认为用户新建投递任务的时间。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 投递数据范围的结束时间点，不能填写未来时间。如果用户不填写，默认为持续投递，即无限。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type UnbindDeliverCloudFunctionRequest struct {
	*tchttp.BaseRequest

	// 投递规则属于的 topic id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *UnbindDeliverCloudFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindDeliverCloudFunctionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeliverCloudFunctionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeliverCloudFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeliverCloudFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCosRechargeRequest struct {
	*tchttp.BaseRequest

	// 主键Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 日志主题Id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// cos导入任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否启用:   0： 未启用  ， 1：启用

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyCosRechargeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCosRechargeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeliverCloudFunctionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeliverCloudFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeliverCloudFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitPartitionRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 待分裂分区ID

	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
	// 分区切分的哈希key的位置，只在Number=2时有意义

	SplitKey *string `json:"SplitKey,omitempty" name:"SplitKey"`
	// 分区分裂个数(可选)，默认等于2

	Number *int64 `json:"Number,omitempty" name:"Number"`
}

func (r *SplitPartitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SplitPartitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreviewLogStatistic struct {

	// 目标日志主题

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
	// 日志内容

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
	// 失败错误码， 空字符串""表示正常

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
	// 行号

	LineNum *int64 `json:"LineNum,omitempty" name:"LineNum"`
	// 日志时间戳

	Time *string `json:"Time,omitempty" name:"Time"`
	// 目标topic-name

	DstTopicName *string `json:"DstTopicName,omitempty" name:"DstTopicName"`
}

type CreateCosRechargeRequest struct {
	*tchttp.BaseRequest

	// 日志主题 ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 投递任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 存储桶

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 存储桶所在地域

	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`
	// cos文件所在文件夹的前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志； 默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// supported: "", "gzip", "lzop", "snappy”; 默认空

	Compress *string `json:"Compress,omitempty" name:"Compress"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitempty" name:"ExtractRuleInfo"`
}

func (r *CreateCosRechargeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCosRechargeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// 查询的机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryShipperTaskRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// 投递任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RetryShipperTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryShipperTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCosRechargeInfoRequest struct {
	*tchttp.BaseRequest

	// 日志主题 ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 投递任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 存储桶

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 存储桶所在地域

	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`
	// cos文件所在文件夹的前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 压缩模式:   "", "gzip", "lzop", "snappy”;   默认""

	Compress *string `json:"Compress,omitempty" name:"Compress"`
}

func (r *SearchCosRechargeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCosRechargeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQcloudHourRequest struct {
	*tchttp.BaseRequest

	//  开通接口信息

	Param *QcloudInterfacePara `json:"Param,omitempty" name:"Param"`
}

func (r *CreateQcloudHourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQcloudHourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataTransformFailureInfo struct {

	// 源日志

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
	// 加工失败原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type CreateDemonstrationsRequest struct {
	*tchttp.BaseRequest

	// 演示示例类型：'CLB', 'NGINX'

	Type *string `json:"Type,omitempty" name:"Type"`
	// 日志主题ID。传入表示对已有日志主题生成演示示例资源，如仪表盘、监控告警；不传入表示生成包括日志集、日志主题在内的演示示例资源。只支持新建时传入，不支持与DemonstrationIds同时传入。

	TopicIds []*string `json:"TopicIds,omitempty" name:"TopicIds"`
	// 演示示例ID，重置时传入。不支持与TopicIds同时传入。

	DemonstrationIds []*string `json:"DemonstrationIds,omitempty" name:"DemonstrationIds"`
}

func (r *CreateDemonstrationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDemonstrationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataTransformRequest struct {
	*tchttp.BaseRequest

	// 数据加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteDataTransformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataTransformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 默认不生效

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 索引规则，Rule和Effective两个必须有一个参数存在

	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
	// 全文索引系统预置字段标记，默认false。  false:不包含系统预置字段， true:包含系统预置字段

	IncludeInternalFields *bool `json:"IncludeInternalFields,omitempty" name:"IncludeInternalFields"`
	// 元数据相关标志位，默认为0。 0：全文索引包括开启键值索引的元数据字段， 1：全文索引包括所有元数据字段，2：全文索引不包括元数据字段。

	MetadataFlag *uint64 `json:"MetadataFlag,omitempty" name:"MetadataFlag"`
	// 自定义日志解析异常存储字段。

	CoverageField *string `json:"CoverageField,omitempty" name:"CoverageField"`
}

func (r *ModifyIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {

	// 列的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 列的属性

	Type *string `json:"Type,omitempty" name:"Type"`
}

type UploadServiceLogRequest struct {
	*tchttp.BaseRequest

	// 主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 根据 hashkey 写入相应范围的主题分区

	HashKey *string `json:"HashKey,omitempty" name:"HashKey"`
	// agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// agent的Instance ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// agent是否开启自动升级功能

	AutoUpdate *string `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// agent请求序列号

	AgentSeq *string `json:"AgentSeq,omitempty" name:"AgentSeq"`
	// agent版本

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// 请求Unique ID

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
	// 服务日志类型

	MetricType *string `json:"MetricType,omitempty" name:"MetricType"`
	// agent的label信息

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 压缩方法

	CompressType *string `json:"CompressType,omitempty" name:"CompressType"`
}

func (r *UploadServiceLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadServiceLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCosRechargeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCosRechargeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyShipperResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyShipperResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindDeliverCloudFunctionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindDeliverCloudFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindDeliverCloudFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PartitionInfo struct {

	// 分区ID

	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
	// 分区的状态（readwrite或者是readonly）

	Status *string `json:"Status,omitempty" name:"Status"`
	// 分区哈希键起始key

	InclusiveBeginKey *string `json:"InclusiveBeginKey,omitempty" name:"InclusiveBeginKey"`
	// 分区哈希键结束key

	ExclusiveEndKey *string `json:"ExclusiveEndKey,omitempty" name:"ExclusiveEndKey"`
	// 分区创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 只读分区数据停止写入时间

	LastWriteTime *string `json:"LastWriteTime,omitempty" name:"LastWriteTime"`
}

type DescribeAsyncContextTasksRequest struct {
	*tchttp.BaseRequest

	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// <br><li> topicId
	//
	// 按照【日志主题ID】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> logsetId
	//
	// 按照【日志集ID】进行过滤，可通过调用DescribeLogsets查询已创建的日志集列表或登录控制台进行查看；也可以调用CreateLogset创建新的日志集。
	//
	// 类型：String
	//
	// 必选：否
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为5

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAsyncContextTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncContextTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeliverCloudFunctionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递规则属于的 topic id

		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
		// 投递的云函数名字

		FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`
		// 命名空间

		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
		// 函数版本

		Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`
		// 投递最长等待时间，单位：秒

		Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
		// 投递最大消息数

		MaxMsgNum *uint64 `json:"MaxMsgNum,omitempty" name:"MaxMsgNum"`
		// 投递状态： 0：关闭投递， 1：开启投递

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeliverCloudFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeliverCloudFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志导出列表

		Exports []*ExportInfo `json:"Exports,omitempty" name:"Exports"`
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志时间,  格式: YYYY-mm-dd HH:MM:SS

	BTime *string `json:"BTime,omitempty" name:"BTime"`
	// 日志包序号

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 日志包内一条日志的序号

	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 上文日志条数,  默认值10

	PrevLogs *int64 `json:"PrevLogs,omitempty" name:"PrevLogs"`
	// 下文日志条数,  默认值10

	NextLogs *int64 `json:"NextLogs,omitempty" name:"NextLogs"`
}

func (r *DescribeLogContextRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogContextRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTime struct {

	// 可选值：
	// <br><li> Period - 周期执行
	// <br><li> Fixed - 定期执行

	Type *string `json:"Type,omitempty" name:"Type"`
	// 执行的周期，或者定制执行的时间节点。单位为分钟，取值范围为1~1440。

	Time *int64 `json:"Time,omitempty" name:"Time"`
}

type CallBackInfo struct {

	// 回调时的Body

	Body *string `json:"Body,omitempty" name:"Body"`
	// 回调时的Headers

	Headers []*string `json:"Headers,omitempty" name:"Headers"`
}

type CreateRebuildIndexTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 重建起始时间戳，毫秒

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 重建结束时间戳，毫秒

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *CreateRebuildIndexTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRebuildIndexTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID。

	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
}

func (r *DeleteAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户状态，0:未开通，1:正常，2: 欠费， 3: 销毁

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 告警模板名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 告警模板的类型。可选值：
	// <br><li> Trigger - 告警触发
	// <br><li> Recovery - 告警恢复
	// <br><li> All - 告警触发和告警恢复

	Type *string `json:"Type,omitempty" name:"Type"`
	// 告警模板接收者信息。

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 告警模板回调信息。

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
	// 告警通知模板ID。

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionRequest struct {
	*tchttp.BaseRequest

	// dfsf

	Tsed *string `json:"Tsed,omitempty" name:"Tsed"`
}

func (r *ActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigInfo struct {

	// 采集规则配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 采集规则配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 日志采集路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 采集配置所属日志主题ID即TopicId

	Output *string `json:"Output,omitempty" name:"Output"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 用户自定义解析字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// config_extra主键ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集配置标签

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：  ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时 ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数 ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false 样例： {\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true} 控制台默认占位值：{\"ClsAgentDefault\":0}
	AdvancedConfig *string `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`

	// 采集配置来源，0：默认来源，1: TKE
	Source *uint64 `json:"Source,omitempty" name:"Source"`

	// 日志输入类型，支持file、window_event、syslog、k8s_stdout、k8s_file
	//InputType *string `json:"InputType,omitempty" name:"InputType"`
}

type ModifyDemonstrationResourcesRequest struct {
	*tchttp.BaseRequest

	// 演示示例资源

	Resources []*DemonstrationResource `json:"Resources,omitempty" name:"Resources"`
}

func (r *ModifyDemonstrationResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDemonstrationResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitPartitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分裂结果集

		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SplitPartitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SplitPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertRecordHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 记录

		Records []*AlertHistoryRecord `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertRecordHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertRecordHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformPreviewInfoRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDataTransformPreviewInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformPreviewInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DemonstrationResource struct {

	// 资源类型：'LOGSET' | 'TOPIC' | 'DASHBOARD' | 'ALARM' | 'ALARM_NOTICE'

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 是否启用。目前只用于表示Topic是否开启日志自动写入。

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 资源所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 演示示例子类型

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 资源类型为TOPIC的上传状态：'INITIAL' | 'UPLOADING' | 'UPLOAD_FAILED' | 'STOPPED'

	UploadStatus *string `json:"UploadStatus,omitempty" name:"UploadStatus"`
	// 模版项ID

	TemplateItemId *string `json:"TemplateItemId,omitempty" name:"TemplateItemId"`
	// 创建者：CLS, USER

	Creator *string `json:"Creator,omitempty" name:"Creator"`
}

type DescribeAlarmsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略列表。

		Alarms []*AlarmInfo `json:"Alarms,omitempty" name:"Alarms"`
		// 符合查询条件的告警策略数目。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncSearchResultRequest struct {
	*tchttp.BaseRequest

	// 异步检索任务ID

	AsyncSearchTaskId *string `json:"AsyncSearchTaskId,omitempty" name:"AsyncSearchTaskId"`
	// 加载更多日志时使用，透传上次返回的Context值，获取后续的日志内容

	Context *string `json:"Context,omitempty" name:"Context"`
	// 单次调用返回的日志条数，默认为20，最大为500

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 日志集ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeAsyncSearchResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncSearchResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicExtendConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// clb的topic业务配置

		ClbTopicExtendConfigs []*ClbTopicExtendConfig `json:"ClbTopicExtendConfigs,omitempty" name:"ClbTopicExtendConfigs"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicExtendConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicExtendConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogsetRequest struct {
	*tchttp.BaseRequest

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志集的绑定的标签键值对。最大支持10个标签键值对，同一个资源只能同时绑定一个标签键。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 生命周期，单位为天；可取值范围1~366

	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *ModifyLogsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Kafka协议消费打开状态

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 待消费TopicId

		TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenKafkaConsumeRequest struct {
	*tchttp.BaseRequest

	// CLS相关TopicId

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *OpenKafkaConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKafkaConsumeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// CLS对应的topic标识

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *CloseKafkaConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKafkaConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FullTextInfo struct {

	// 是否大小写敏感

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 全文索引的分词符，字符串中每个字符代表一个分词符

	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`
	// 是否包含中文

	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type MachineInfo struct {

	// 机器的IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 机器状态，0:异常，1:正常

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 机器离线时间，空为正常，异常返回具体时间

	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
	// 机器是否开启自动升级。0:关闭，1:开启

	AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 机器当前版本号。

	Version *string `json:"Version,omitempty" name:"Version"`
	// 机器升级功能状态。

	UpdateStatus *int64 `json:"UpdateStatus,omitempty" name:"UpdateStatus"`
	// 机器升级结果标识。

	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`
	// 机器升级结果信息。

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type DeleteConfigFromMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 采集配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigFromMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigFromMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineGroupInfo struct {

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 机器组类型

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 机器组绑定的标签列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否开启机器组自动更新

	AutoUpdate *string `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 升级开始时间，建议业务低峰期升级LogListener

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 升级结束时间，建议业务低峰期升级LogListener

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费

	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
	// 默认值""， "label_k8s"

	Flag *string `json:"Flag,omitempty" name:"Flag"`

	// 机器组中机器离线清理时间。单位：天
	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitempty" name:"DelayCleanupTime"`

	// 机器组元数据信息列表
	OSType *uint64 `json:"OSType,omitempty" name:"OSType"`

	// 系统类型，取值如下：  0：Linux （默认值） 1：Windows
	MetaTags []*MetaTagInfo `json:"MetaTags,omitempty" name:"MetaTags"`
}

type SendDetail struct {

	// 发送次数

	Sms []*SendDetailItem `json:"Sms,omitempty" name:"Sms"`
	// 发送次数

	Email []*SendDetailItem `json:"Email,omitempty" name:"Email"`
	// 发送次数

	WeChat []*SendDetailItem `json:"WeChat,omitempty" name:"WeChat"`
	// 发送次数

	Phone []*SendDetailItem `json:"Phone,omitempty" name:"Phone"`
	// 发送次数

	Callback []*SendDetailItem `json:"Callback,omitempty" name:"Callback"`
}

type DeleteDemonstrationsRequest struct {
	*tchttp.BaseRequest

	// 演示示例ID

	DemonstrationIds []*string `json:"DemonstrationIds,omitempty" name:"DemonstrationIds"`
}

func (r *DeleteDemonstrationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDemonstrationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// <br><li> name
	//
	// 按照【告警通知模板名称】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> alarmNoticeId
	//
	// 按照【告警通知模板ID】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> uid
	//
	// 按照【接收用户ID】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> groupId
	//
	// 按照【用户组ID】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 查询语句，语句长度最大为4096

	Query *string `json:"Query,omitempty" name:"Query"`
	// 单次查询返回的原始日志条数，最大值为100。查询语句(Query)包含SQL时，针对SQL的结果条数需在Query中指定，参考https://cloud.tencent.com/document/product/614/58977

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 加载更多日志时使用，透传上次返回的Context值，获取后续的日志内容。过期时间1小时

	Context *string `json:"Context,omitempty" name:"Context"`
	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 是否返回检索的高亮结果

	HighLight *bool `json:"HighLight,omitempty" name:"HighLight"`
	// 为true代表使用新检索,响应参数AnalysisRecords和Columns有效， 为false时代表使用老检索方式, AnalysisResults和ColNames有效

	UseNewAnalysis *bool `json:"UseNewAnalysis,omitempty" name:"UseNewAnalysis"`
	// 0：不执行语法优化；1：执行语法优化

	QueryOptimize *uint64 `json:"QueryOptimize,omitempty" name:"QueryOptimize"`
	// 0：表示客户选择auto自动采样率; 0～1之间：表示客户指定的采样率（例如0.02; 1（默认值）：表示不采样。

	SamplingRate *float64 `json:"SamplingRate,omitempty" name:"SamplingRate"`
}

func (r *SearchLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigRequest struct {
	*tchttp.BaseRequest

	// 采集配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 日志采集路径,包含文件名

	Path *string `json:"Path,omitempty" name:"Path"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 采集配置所属日志主题ID即TopicId

	Output *string `json:"Output,omitempty" name:"Output"`
	// 用户自定义采集规则，Json格式序列化的字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// config_extra表主键ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集配置标签

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：  ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时 ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数 ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false 样例： {\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true} 控制台默认占位值：{\"ClsAgentDefault\":0}
	AdvancedConfig *string `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`

	// 采集配置来源，0：默认来源，1: TKE
	Source *uint64 `json:"Source,omitempty" name:"Source"`

	// 日志输入类型，支持file、window_event、syslog、k8s_stdout、k8s_file
	//InputType *string `json:"InputType,omitempty" name:"InputType"`
}

func (r *CreateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// CLS控制台创建的TopicId

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *ModifyKafkaConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKafkaConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncSearchTaskRequest struct {
	*tchttp.BaseRequest

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志主题ID，目前仅支持StorageType为cold的日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 查询语句，语句长度最大为1024

	Query *string `json:"Query,omitempty" name:"Query"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 日志扫描顺序；可选值：asc(升序)、desc(降序)，默认为 desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *CreateAsyncSearchTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAsyncSearchTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigExtraInfo struct {

	// 采集规则扩展配置ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 类型：container_stdout、container_file、host_file

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点文件配置信息

	HostFile *HostFileInfo `json:"HostFile,omitempty" name:"HostFile"`
	// 容器文件路径信息

	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitempty" name:"ContainerFile"`
	// 容器标准输出信息

	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitempty" name:"ContainerStdout"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 用户自定义解析字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 自建采集配置标

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集name

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题name

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type KafkaInfo struct {

	// 可消费topic名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// ACL模式用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// ACL模式密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 可消费kafka实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// SASL接入点信息

	BootstrapServers *string `json:"BootstrapServers,omitempty" name:"BootstrapServers"`
}

type CreateConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题 ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 是否投递日志的元数据信息，默认为 true

	NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`
	// 如果需要投递元数据信息，元数据信息的描述

	Content *ConsumerContent `json:"Content,omitempty" name:"Content"`
	// CKafka的描述

	Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`
	// 压缩方式，0：NONE，2：SNAPPY，3：LZ4

	Compression *int64 `json:"Compression,omitempty" name:"Compression"`
}

func (r *CreateConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestJsonLogRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeLatestJsonLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestJsonLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertHistoryRecord struct {

	// 记录ID

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
	// 报警ID

	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
	// 告警名

	AlarmName *string `json:"AlarmName,omitempty" name:"AlarmName"`
	// topic的id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// topic名

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 触发条件

	Trigger *string `json:"Trigger,omitempty" name:"Trigger"`
	// 报警发生阈值

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 连续多少次报警，进行通知

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 报警对象

	Notices []*AlertHistoryNotice `json:"Notices,omitempty" name:"Notices"`
	// 连续报警时间

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 发生时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type TopicInfo struct {

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 主题分区个数

	PartitionCount *int64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
	// 是否开启索引

	Index *bool `json:"Index,omitempty" name:"Index"`
	// AssumerUin非空则表示创建该日志主题的服务方Uin

	AssumerUin *uint64 `json:"AssumerUin,omitempty" name:"AssumerUin"`
	// AssumerName非空则表示创建该日志主题的服务方名称

	AssumerName *string `json:"AssumerName,omitempty" name:"AssumerName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 若AssumerUin非空，则表示除服务方外其余调用者修改日志主题的权限

	TopicModifyAcl *int64 `json:"TopicModifyAcl,omitempty" name:"TopicModifyAcl"`
	// 若AssumerUin非空，则表示除服务方外其余调用者展示日志主题的权限

	TopicShowAcl *int64 `json:"TopicShowAcl,omitempty" name:"TopicShowAcl"`
	// 日主主题是否开启采集

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 日志主题绑定的标签信息

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// RoleName非空则表示创建该日志主题的服务方使用的角色

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 该主题是否开启自动分裂

	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`
	// 若开启自动分裂的话，该主题能够允许的最大分区数

	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`
	// 日主题的存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 生命周期，单位为天

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 二级产品标识

	SubAssumerName *string `json:"SubAssumerName,omitempty" name:"SubAssumerName"`
	// topic对应的日志集信息

	LogsetInfo *LogsetInfo `json:"LogsetInfo,omitempty" name:"LogsetInfo"`
	// 日志主题描述

	Describes *string `json:"Describes,omitempty" name:"Describes"`

	// 0：关闭日志沉降。 非0：开启日志沉降后标准存储的天数。HotPeriod需要大于等于7，且小于Period。仅在StorageType为 hot 时生效
	HotPeriod *uint64 `json:"HotPeriod,omitempty" name:"HotPeriod"`

	// 加密相关参数。 支持加密地域并且开白用户可以传此参数，其他场景不能传递该参数。  0或者不传： 不加密 1：kms-cls 云产品秘钥加密
	Encryption *uint64 `json:"Encryption,omitempty" name:"Encryption"`

	// 0: 日志主题topic 1: 时序主题topic
	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`

	// webtracking开关； false: 关闭 true： 开启
	IsWebTracking *bool `json:"IsWebTracking,omitempty" name:"IsWebTracking"`
}

type ValueInfo struct {

	// 字段类型，目前支持的类型有：long、text、double

	Type *string `json:"Type,omitempty" name:"Type"`
	// 字段的分词符，只有当字段类型为text时才有意义；输入字符串中的每个字符代表一个分词符

	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`
	// 字段是否开启分析功能

	SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`
	// 是否包含中文

	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type DescribeKafkaUserRequest struct {
	*tchttp.BaseRequest

	// kafka消费用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DescribeKafkaUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformFailLogInfoRequest struct {
	*tchttp.BaseRequest

	// 数据加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 要查询的起始时间，Unix时间戳，单位ms

	From *uint64 `json:"From,omitempty" name:"From"`
	// 要查询的结束时间，Unix时间戳，单位ms

	To *uint64 `json:"To,omitempty" name:"To"`
	// 目标日志主题id

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
}

func (r *DescribeDataTransformFailLogInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformFailLogInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupConfigsRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeMachineGroupConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGroupConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QcloudGoodsDetailInfo struct {

	// 价格模型

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 商品的时间单位：y:年；m:月；d:日；h:小时；M:分钟；s:秒; p:与计费周期无关,一次性购买的产品传p

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 商品的时间大小（一次性售卖固定传1）

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 放入goodsDetail内:子产品标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 新增参数: 产品标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 相关计费项个数

	SvClsPartitionCount *uint64 `json:"SvClsPartitionCount,omitempty" name:"SvClsPartitionCount"`
}

type CreateNonBillingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNonBillingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNonBillingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportRequest struct {
	*tchttp.BaseRequest

	// 日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志导出检索语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 日志导出数量,  最大值1000万

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 日志导出时间排序。desc，asc，默认为desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 日志导出数据格式。json，csv，默认为json

	Format *string `json:"Format,omitempty" name:"Format"`
	// 日志导出起始时间，毫秒时间戳

	From *int64 `json:"From,omitempty" name:"From"`
	// 日志导出结束时间，毫秒时间戳

	To *int64 `json:"To,omitempty" name:"To"`
}

func (r *CreateExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeAgentNormalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeAgentNormalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeAgentNormalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogConfigInfo struct {

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 采集日志路径列表

	Path *string `json:"Path,omitempty" name:"Path"`
	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 日志格式化格式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 黑名单path列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 用户自定义解析字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
}

type TopicIdLogStatistic struct {

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 读取的源日志主题的行数

	ReadLines *int64 `json:"ReadLines,omitempty" name:"ReadLines"`
	// 加工后输出到目标日志主题的行数

	WriteLines *int64 `json:"WriteLines,omitempty" name:"WriteLines"`
	// 加工失败的行数

	FailedLines *int64 `json:"FailedLines,omitempty" name:"FailedLines"`
	// 加工过滤的行数

	FilterLines *uint64 `json:"FilterLines,omitempty" name:"FilterLines"`
}

type GetAlarmLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加载后续内容的Context

		Context *string `json:"Context,omitempty" name:"Context"`
		// 日志查询结果是否全部返回

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 返回的是否为分析结果

		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`
		// 如果Analysis为True，则返回分析结果的列名，否则为空

		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`
		// 日志查询结果；当Analysis为True时，可能返回为null

		Results []*LogInfo `json:"Results,omitempty" name:"Results"`
		// 日志分析结果；当Analysis为False时，可能返回为null

		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`
		// 新的日志分析结果; UseNewAnalysis为true有效

		AnalysisRecords []*string `json:"AnalysisRecords,omitempty" name:"AnalysisRecords"`
		// 日志分析的列属性; UseNewAnalysis为true有效

		Columns []*Column `json:"Columns,omitempty" name:"Columns"`
		// 返回语法优化后的语句（入参如有QueryOptimize 且为 1 时返回，其他情况返回空字符串）

		Query *string `json:"Query,omitempty" name:"Query"`
		// 当前系统使用的采样率，入参如有SamplingRate时生效（主要是当客户输入0时，返回真实后台的AutoSamplingRate值）

		SamplingRate *float64 `json:"SamplingRate,omitempty" name:"SamplingRate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlarmLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlarmLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicExtendConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicExtendConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicExtendConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QcloudInterfacePara struct {

	// 恒传1，该字段若传值表示是否开通多个物品，需把开通各项物品传入goodsInfoList

	Multi *uint64 `json:"Multi,omitempty" name:"Multi"`
	// 开通的物品列表，以下参数为每个物品相关的参数

	GoodsInfoList []*QcloudGoodsInfoListInfo `json:"GoodsInfoList,omitempty" name:"GoodsInfoList"`
}

type DescribeShipperTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递任务列表

		Tasks []*ShipperTaskInfo `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShipperTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShipperTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeReceiver struct {

	// 接受者类型。可选值：
	// <br><li> Uin - 用户ID
	// <br><li> Group - 用户组ID
	// 暂不支持其余接收者类型。

	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 接收者。

	ReceiverIds []*int64 `json:"ReceiverIds,omitempty" name:"ReceiverIds"`
	// 通知接收渠道。
	// <br><li> Email - 邮件
	// <br><li> Sms - 短信
	// <br><li> WeChat - 微信
	// <br><li> Phone - 电话

	ReceiverChannels []*string `json:"ReceiverChannels,omitempty" name:"ReceiverChannels"`
	// 允许接收信息的开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 允许接收信息的结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 位序

	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type DescribeAsyncContextTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步上下文任务列表

		AsyncContextTasks []*AsyncContextTask `json:"AsyncContextTasks,omitempty" name:"AsyncContextTasks"`
		// 异步上下文任务的总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncContextTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncContextTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQcloudHourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本

		Version *uint64 `json:"Version,omitempty" name:"Version"`
		// 组件名称

		ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
		// 事件Id

		EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
		// 时间戳

		Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
		// 返回值

		ReturnValue *int64 `json:"ReturnValue,omitempty" name:"ReturnValue"`
		// 返回code

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 返回信息

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// 返回数据

		Data *QcloudDataInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateQcloudHourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQcloudHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenKVRegexRequest struct {
	*tchttp.BaseRequest

	// 日志样例

	LogSample *string `json:"LogSample,omitempty" name:"LogSample"`
	// 样例索引列表

	Indexes []*RegexIndexInfo `json:"Indexes,omitempty" name:"Indexes"`
	// 是否为多行全文，默认单行

	Multiline *bool `json:"Multiline,omitempty" name:"Multiline"`
}

func (r *GenKVRegexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenKVRegexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAlarmRuleRequest struct {
	*tchttp.BaseRequest

	// 触发条件

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 触发次数

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 报警周期

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 报警时间

	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`
	// 监控对象

	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`
}

func (r *CheckAlarmRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAlarmRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DynamicIndex struct {
	Status *bool `json:"Status,omitempty" name:"Status"`
}

type DescribeIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题ID

		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
		// 是否生效

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 索引配置信息

		Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
		// 索引修改时间，初始值为索引创建时间。

		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
		// 全文索引系统预置字段标记，默认false。  false:不包含系统预置字段， true:包含系统预置字段

		IncludeInternalFields *bool `json:"IncludeInternalFields,omitempty" name:"IncludeInternalFields"`
		// 元数据相关标志位，默认为0。 0：全文索引包括开启键值索引的元数据字段， 1：全文索引包括所有元数据字段，2：全文索引不包括元数据字段。

		MetadataFlag *uint64 `json:"MetadataFlag,omitempty" name:"MetadataFlag"`
		// 自定义日志解析异常存储字段。

		CoverageField *string `json:"CoverageField,omitempty" name:"CoverageField"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateShipperRequest struct {
	*tchttp.BaseRequest

	// 创建的投递规则所属的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 创建的投递规则投递的bucket

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 创建的投递规则投递目录的前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 投递规则的名字

	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`
	// 投递的时间间隔，单位 秒，默认300，范围 300-900

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 投递的文件的最大值，单位 MB，默认256，范围 100-256

	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`
	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递

	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`
	// 投递日志的分区规则，支持strftime的时间格式表示

	Partition *string `json:"Partition,omitempty" name:"Partition"`
	// 投递日志的压缩配置

	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`
	// 投递日志的内容格式配置

	Content *ContentInfo `json:"Content,omitempty" name:"Content"`

	// 投递文件命名配置，0：随机数命名，1：投递时间命名，默认0（随机数命名）
	FilenameMode *uint64 `json:"FilenameMode,omitempty" name:"FilenameMode"`

	// 跨账户uin，用于支持跨账户bucket投递
	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`

	// 投递数据范围的开始时间点，不能超出日志主题的生命周期起点。如果用户不填写，默认为用户新建投递任务的时间。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 投递数据范围的结束时间点，不能填写未来时间。如果用户不填写，默认为持续投递，即无限。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *CreateShipperRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateShipperRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDemonstrationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDemonstrationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDemonstrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupsRequest struct {
	*tchttp.BaseRequest

	// <br><li> machineGroupName
	//
	// 按照【机器组名称】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> machineGroupId
	//
	// 按照【机器组ID】进行过滤。
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> tagKey
	//
	// 按照【标签键】进行过滤。
	//
	// 类型：String
	//
	// 必选：否
	//
	// <br><li> tag:tagKey
	//
	// 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。
	// 类型：String
	//
	// 必选：否
	//
	//
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMachineGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRebuildIndexTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 索引重建任务列表

		RebuildTasks []*RebuildIndexTaskInfo `json:"RebuildTasks,omitempty" name:"RebuildTasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRebuildIndexTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRebuildIndexTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
