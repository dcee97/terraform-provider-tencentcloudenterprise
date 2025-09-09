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

package v20181210

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type BlockConfigInterfaceIEditApiConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIEditApiConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIEditApiConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelManagerInterfaceIAddSourcePlatformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LabelManagerInterfaceIAddSourcePlatformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LabelManagerInterfaceIAddSourcePlatformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportIpFlowInterfaceIGetIpFlowTableRequest struct {
	*tchttp.BaseRequest

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 当前页数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 默认按in方向的流量峰值 大到小，JSON格式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询标记，多余参数

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 搜索内容

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *ReportIpFlowInterfaceIGetIpFlowTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReportIpFlowInterfaceIGetIpFlowTableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetThreatRawLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetThreatRawLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetThreatRawLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBatchChangeStatus struct {

	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// "stop"表示暂停，"start"表示启动

	NetSwitch *string `json:"NetSwitch,omitempty" name:"NetSwitch"`
	// "WHITE"或者"BLACK",表示黑白

	Type *string `json:"Type,omitempty" name:"Type"`
}

type AlarmInterfaceIGetSystemRawLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetSystemRawLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetSystemRawLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetThresholdConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetThresholdConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetThresholdConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceGetStatisticResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceGetStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceGetStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetTOPStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetTOPStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetTOPStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIGetRuleDetailStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SystemRuleInterfaceIGetRuleDetailStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIGetRuleDetailStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceAdvSerachContent struct {

	// src_ip

	SrcIp []*string `json:"SrcIp,omitempty" name:"SrcIp"`
	// dst_ip

	DstIp []*string `json:"DstIp,omitempty" name:"DstIp"`
}

type IpFlowTable struct {

	// 排序字段。可选值为 max_in_flow/max_out_flow

	Field *string `json:"Field,omitempty" name:"Field"`
	// 是否倒叙。0：正叙，1：倒叙

	IsInvert *uint64 `json:"IsInvert,omitempty" name:"IsInvert"`
}

type AlarmInterfaceIGetSystemConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetSystemConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetSystemConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetIdsInterceptConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetIdsInterceptConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetIdsInterceptConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceIGetMsgStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *MsgInterfaceIGetMsgStatisticInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MsgInterfaceIGetMsgStatisticInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmTrendStatInfoRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *OverviewInterfaceIGetAlarmTrendStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetAlarmTrendStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Order struct {

	// direc

	Direc []*string `json:"Direc,omitempty" name:"Direc"`
	// field

	Field *string `json:"Field,omitempty" name:"Field"`
}

type BlockConfigInterfaceIGetOrResetSecretKeyRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 无

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 无

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *BlockConfigInterfaceIGetOrResetSecretKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetOrResetSecretKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIStopFlowRequest struct {
	*tchttp.BaseRequest

	// 操作

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *TceManagerInterfaceIStopFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIStopFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetRedirectURLRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIGetRedirectURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetRedirectURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetAlarmReasonCntResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetAlarmReasonCntResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetAlarmReasonCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchInsertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIBatchInsertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchInsertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIListRuleRequest struct {
	*tchttp.BaseRequest

	// 时间

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 当前页数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 执行标志

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 输入内容

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 规则状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 黑白名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RuleConfigManagerInterfaceIListRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIListRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIDeleteStopNetworkConfigRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIDeleteStopNetworkConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIDeleteStopNetworkConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceISetUnreadMsgToReadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MsgInterfaceISetUnreadMsgToReadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MsgInterfaceISetUnreadMsgToReadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchDeleteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIBatchDeleteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchDeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetSystemRawLogStatisticInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetSystemRawLogStatisticInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetSystemRawLogStatisticInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetThreatRawLogDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIGetFilterItemsListInfoRequest struct {
	*tchttp.BaseRequest

	// 查询标识

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 时间

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 查询内容

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 当前页数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RuleConfigManagerInterfaceIGetFilterItemsListInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIGetFilterItemsListInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIGetRuleDetailStatRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 无

	FilterAz []*string `json:"FilterAz,omitempty" name:"FilterAz"`
	// 无

	Search *string `json:"Search,omitempty" name:"Search"`
	// 无

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 无

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 无

	SourcePlatform []*string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// 无

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// AddTime

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// SearchContent

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
	// SourceId

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
}

func (r *SystemRuleInterfaceIGetRuleDetailStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIGetRuleDetailStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIChangeBlockStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIChangeBlockStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIChangeBlockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiConfigInfoListRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 无

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *BlockConfigInterfaceIGetApiConfigInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetApiConfigInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetNetoffConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetNetoffConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetNetoffConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetAlarmLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetAlarmLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetAlarmLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetStatisticsRequest struct {
	*tchttp.BaseRequest

	// region_id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 用户选择的az列表，数据类型为数组

	RuleAz []*string `json:"RuleAz,omitempty" name:"RuleAz"`
	// 高级查询下拉项

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// status

	Status []*string `json:"Status,omitempty" name:"Status"`
	// valid_duration

	ValidDuration []*string `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// hit_history

	HitHistory []*string `json:"HitHistory,omitempty" name:"HitHistory"`
	// hit_week

	HitWeek []*string `json:"HitWeek,omitempty" name:"HitWeek"`
	// hit_day

	HitDay []*string `json:"HitDay,omitempty" name:"HitDay"`
	// match_operation

	MatchOperation []*string `json:"MatchOperation,omitempty" name:"MatchOperation"`
	// ip_type

	IpType []*string `json:"IpType,omitempty" name:"IpType"`
	// source_platform

	SourcePlatform []*string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// ip_version

	IpVersion []*string `json:"IpVersion,omitempty" name:"IpVersion"`
	// BeginTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// AddTime

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// SearchContent

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *GlobalAccessInterfaceIGetStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIGetStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RawlogSearchInterfaceIQueryRwLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RawlogSearchInterfaceIQueryRwLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RawlogSearchInterfaceIQueryRwLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIEditRuleRequest struct {
	*tchttp.BaseRequest

	// 见接口描述

	ValidDuration *uint64 `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// 见接口描述

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 见接口描述

	Type *string `json:"Type,omitempty" name:"Type"`
	// 见接口描述

	Name *string `json:"Name,omitempty" name:"Name"`
	// 见接口描述

	MatchOperation *string `json:"MatchOperation,omitempty" name:"MatchOperation"`
	// 见接口描述

	SourcePlatform *string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// 见接口描述

	IsOverwrite *string `json:"IsOverwrite,omitempty" name:"IsOverwrite"`
	// 见接口描述

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// 见接口描述

	IpData *string `json:"IpData,omitempty" name:"IpData"`
	// 见接口描述

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// AzName

	AzName []*string `json:"AzName,omitempty" name:"AzName"`
	// BatchId

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// DstIp

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// DstPort

	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`
	// HitDay

	HitDay *string `json:"HitDay,omitempty" name:"HitDay"`
	// HitHistory

	HitHistory *int64 `json:"HitHistory,omitempty" name:"HitHistory"`
	// HitWeek

	HitWeek *string `json:"HitWeek,omitempty" name:"HitWeek"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// Reboot

	Reboot *bool `json:"Reboot,omitempty" name:"Reboot"`
	// RuleId

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// SendStatus

	SendStatus *string `json:"SendStatus,omitempty" name:"SendStatus"`
	// SrcIp

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// UpdateTime

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// ValidEndTime

	ValidEndTime *string `json:"ValidEndTime,omitempty" name:"ValidEndTime"`
	// ValidStartTime

	ValidStartTime *string `json:"ValidStartTime,omitempty" name:"ValidStartTime"`
}

func (r *GlobalAccessInterfaceIEditRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIEditRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmStatInfoRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *OverviewInterfaceIGetAlarmStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetAlarmStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshConfigInterfaceIUpdateRefreshConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshConfigInterfaceIUpdateRefreshConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshConfigInterfaceIUpdateRefreshConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportNetFlowInterfaceIGetNetFlowViewDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReportNetFlowInterfaceIGetNetFlowViewDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReportNetFlowInterfaceIGetNetFlowViewDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIInsertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIInsertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIInsertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceGetBlockLogHistogramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceGetBlockLogHistogramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceGetBlockLogHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIDeleteStopNetworkConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIDeleteStopNetworkConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIDeleteStopNetworkConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetAvailableCountRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIGetAvailableCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetAvailableCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateAzIdsInterceptConfigRequest struct {
	*tchttp.BaseRequest

	// {
	//     "region_id": "region_1",
	//     "az": ["az_1"],
	//     "black_duration": 30,    # 拉黑时长
	//     }

	Data *string `json:"Data,omitempty" name:"Data"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIUpdateAzIdsInterceptConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIUpdateAzIdsInterceptConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetSecretInfoRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Passwd

	Passwd *string `json:"Passwd,omitempty" name:"Passwd"`
	// Time

	Time *string `json:"Time,omitempty" name:"Time"`
	// User

	User *string `json:"User,omitempty" name:"User"`
}

func (r *BlockConfigInterfaceIGetSecretInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetSecretInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIGetStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIGetStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceIGetMsgListRequest struct {
	*tchttp.BaseRequest

	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
}

func (r *MsgInterfaceIGetMsgListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MsgInterfaceIGetMsgListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportIpFlowInterfaceIGetIpFlowOverviewRequest struct {
	*tchttp.BaseRequest

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *ReportIpFlowInterfaceIGetIpFlowOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReportIpFlowInterfaceIGetIpFlowOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIDeleteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 动作

	NetSwitch *string `json:"NetSwitch,omitempty" name:"NetSwitch"`
	// 名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RuleConfigManagerInterfaceIDeleteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIDeleteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetBlockLogStatRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 高级查询

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// 告警来源

	AlarmSource []*string `json:"AlarmSource,omitempty" name:"AlarmSource"`
	// 状态

	Status []*string `json:"Status,omitempty" name:"Status"`
	// BeginTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// FlowType

	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// OccurTime

	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// Reason

	Reason []*string `json:"Reason,omitempty" name:"Reason"`
	// ThreatLevel

	ThreatLevel []*string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`
	// ThreatType

	ThreatType []*string `json:"ThreatType,omitempty" name:"ThreatType"`
	// IpVersion

	IpVersion []*string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *AlarmInterfaceIGetBlockLogStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetBlockLogStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIConfigBlockTimeRequest struct {
	*tchttp.BaseRequest

	// 分钟为单位

	BlackDuration *uint64 `json:"BlackDuration,omitempty" name:"BlackDuration"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIConfigBlockTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIConfigBlockTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceISetApiBlockStatusRequest struct {
	*tchttp.BaseRequest

	// 0为关闭，1为开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceISetApiBlockStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceISetApiBlockStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceICommitCheckTaskRuleRequest struct {
	*tchttp.BaseRequest

	// CheckStatus

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *GlobalAccessInterfaceICommitCheckTaskRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceICommitCheckTaskRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetThreatRawLogListRequest struct {
	*tchttp.BaseRequest

	// 查询内容

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询时间

	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime"`
	// 应用层，传输层

	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`
	// 威胁等级，低危，中危，高危，严重，致命

	ThreatLevel *string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`
	// webshell检测

	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`
	// 黑名单，黑客攻击

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 拦截，未拦截

	Status *string `json:"Status,omitempty" name:"Status"`
	// 执行标志

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 当前页

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *TceManagerInterfaceIGetThreatRawLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetThreatRawLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceISetConfigTaskSendStatusRequest struct {
	*tchttp.BaseRequest

	// 无

	TaskID *uint64 `json:"TaskID,omitempty" name:"TaskID"`
	// 是否暂停下发：1-暂停下发，0-重新下发

	Stop *uint64 `json:"Stop,omitempty" name:"Stop"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *AccessRuleInterfaceISetConfigTaskSendStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessRuleInterfaceISetConfigTaskSendStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest struct {
	*tchttp.BaseRequest

	// 置信度阀值

	Threshold *float64 `json:"Threshold,omitempty" name:"Threshold"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIUpdateApiBlockTrustThresholdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIAddRuleRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// Content

	Content *string `json:"Content,omitempty" name:"Content"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// IsOverwrite

	IsOverwrite *string `json:"IsOverwrite,omitempty" name:"IsOverwrite"`
	// MatchOperation

	MatchOperation *string `json:"MatchOperation,omitempty" name:"MatchOperation"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// SourcePlatform

	SourcePlatform *string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// ValidDuration

	ValidDuration *int64 `json:"ValidDuration,omitempty" name:"ValidDuration"`
}

func (r *PreciseAccessInterfaceIAddRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIAddRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIGetRuleStatInfoRequest struct {
	*tchttp.BaseRequest

	// AddTime

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// AdvSerachContent

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// BeginTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// HitCount

	HitCount []*string `json:"HitCount,omitempty" name:"HitCount"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// SearchContent

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
	// Source

	Source []*string `json:"Source,omitempty" name:"Source"`
	// Status

	Status []*string `json:"Status,omitempty" name:"Status"`
	// ValidDuration

	ValidDuration []*string `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// IpVersion

	IpVersion []*string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *SystemRuleInterfaceIGetRuleStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIGetRuleStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAzIdsInterceptStatus struct {

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// AZ

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 状态

	InterceptStatus *string `json:"InterceptStatus,omitempty" name:"InterceptStatus"`
}

type TceManagerInterfaceIGetSystemRawLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetSystemRawLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetSystemRawLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetBlockLogHistogramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetBlockLogHistogramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetBlockLogHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetBlockLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetBlockLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetBlockLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetComplianceInfoRequest struct {
	*tchttp.BaseRequest

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *DayOverviewInterfaceIGetComplianceInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetComplianceInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchChangeWorkStatusRequest struct {
	*tchttp.BaseRequest

	// 规则列表

	Data *string `json:"Data,omitempty" name:"Data"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RuleConfigManagerInterfaceIBatchChangeWorkStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchChangeWorkStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvSerachContent struct {

	// 无

	Host []*string `json:"Host,omitempty" name:"Host"`
	// 无

	Cgi []*string `json:"Cgi,omitempty" name:"Cgi"`
	// 无

	SrcIp []*string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 无

	DstIp []*string `json:"DstIp,omitempty" name:"DstIp"`
	// 无

	DstPort []*string `json:"DstPort,omitempty" name:"DstPort"`
	// 无

	RuleId []*string `json:"RuleId,omitempty" name:"RuleId"`
	// 无

	CreateTime []*string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	StartTime []*string `json:"StartTime,omitempty" name:"StartTime"`
	// 无

	EndTime []*string `json:"EndTime,omitempty" name:"EndTime"`
	// 无

	AlarmTime []*string `json:"AlarmTime,omitempty" name:"AlarmTime"`
	// 无

	EffectIp []*string `json:"EffectIp,omitempty" name:"EffectIp"`
	// 无

	TypeName []*string `json:"TypeName,omitempty" name:"TypeName"`
	// 无

	AlarmId []*string `json:"AlarmId,omitempty" name:"AlarmId"`
	// 无

	ReqSrcIp []*string `json:"ReqSrcIp,omitempty" name:"ReqSrcIp"`
	// 无

	ReqDstIp []*string `json:"ReqDstIp,omitempty" name:"ReqDstIp"`
	// 无

	ReqDstPort []*string `json:"ReqDstPort,omitempty" name:"ReqDstPort"`
	// 无

	ReqSrcPort []*string `json:"ReqSrcPort,omitempty" name:"ReqSrcPort"`
	// 无

	ReqHost []*string `json:"ReqHost,omitempty" name:"ReqHost"`
	// 无

	ReqCgi []*string `json:"ReqCgi,omitempty" name:"ReqCgi"`
	// 无

	AlarmStartTime []*string `json:"AlarmStartTime,omitempty" name:"AlarmStartTime"`
	// 无

	AlarmEndTime []*string `json:"AlarmEndTime,omitempty" name:"AlarmEndTime"`
	// 无

	BatchId []*string `json:"BatchId,omitempty" name:"BatchId"`
	// 无

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 无

	Domain []*string `json:"Domain,omitempty" name:"Domain"`
	// 无

	Url []*string `json:"Url,omitempty" name:"Url"`
	// 无

	UpdateTime []*string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	ValidStartTime []*string `json:"ValidStartTime,omitempty" name:"ValidStartTime"`
	// 无

	ValidEndTime []*string `json:"ValidEndTime,omitempty" name:"ValidEndTime"`
	// 无

	SendTime []*string `json:"SendTime,omitempty" name:"SendTime"`
	// 无

	FlowId []*string `json:"FlowId,omitempty" name:"FlowId"`
}

type AlarmInterfaceIGetThreatRawLogStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// 查询内容

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 威胁等级

	ThreatLevel []*string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`
	// 威胁类型

	ThreatType []*string `json:"ThreatType,omitempty" name:"ThreatType"`
	// 应用层，传输层

	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`
	// 原因

	Reason []*string `json:"Reason,omitempty" name:"Reason"`
	// 拦截，未拦截

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 多余的参数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 多余的参数

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 多余的参数

	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime"`
	// 多余的参数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 高级查询

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// ipVersion

	IpVersion []*string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *AlarmInterfaceIGetThreatRawLogStatisticInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogStatisticInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIAddApiConfigRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// Company

	Company *string `json:"Company,omitempty" name:"Company"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Platform

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// SecretIds

	SecretIds *string `json:"SecretIds,omitempty" name:"SecretIds"`
	// SecretKey

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// WindowCount

	WindowCount *string `json:"WindowCount,omitempty" name:"WindowCount"`
	// WindowTime

	WindowTime *string `json:"WindowTime,omitempty" name:"WindowTime"`
}

func (r *BlockConfigInterfaceIAddApiConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIAddApiConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIConfigRedirectURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIConfigRedirectURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIConfigRedirectURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIStopNetworkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIStopNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIStopNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiBlockStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetApiBlockStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetApiBlockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceICommitCheckTaskRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceICommitCheckTaskRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceICommitCheckTaskRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetBlockTrendStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetBlockTrendStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetBlockTrendStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIGetStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreciseAccessInterfaceIGetStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIGetStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIStopRulesRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Rule

	Rule *string `json:"Rule,omitempty" name:"Rule"`
}

func (r *SystemRuleInterfaceIStopRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIStopRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIConfigRedirectURLRequest struct {
	*tchttp.BaseRequest

	// 跳转地址

	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIConfigRedirectURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIConfigRedirectURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetThresholdConfigRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIGetThresholdConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetThresholdConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogDetailRequest struct {
	*tchttp.BaseRequest

	// ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *AlarmInterfaceIGetThreatRawLogDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetIdsInterceptConfigListRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *BlockConfigInterfaceIGetIdsInterceptConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetIdsInterceptConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIChangeWorkStatusRequest struct {
	*tchttp.BaseRequest

	// 规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 动作

	NetSwitch *string `json:"NetSwitch,omitempty" name:"NetSwitch"`
	// 名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RuleConfigManagerInterfaceIChangeWorkStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIChangeWorkStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIEditRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreciseAccessInterfaceIEditRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIEditRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SrcIpDstIp struct {

	// 源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 目的IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
}

type DayOverviewInterfaceIGetComplianceInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetComplianceInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetComplianceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmTrendStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetAlarmTrendStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetAlarmTrendStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIGetRuleDetailListRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 无

	FilterAz []*string `json:"FilterAz,omitempty" name:"FilterAz"`
	// 无

	Search *string `json:"Search,omitempty" name:"Search"`
	// 无

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 无

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 无

	SourcePlatform []*string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// 无

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// 无

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 无

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// AddTime

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// SearchContent

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
	// SourceId

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
	// ExcuteFlag

	ExcuteFlag *int64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
}

func (r *SystemRuleInterfaceIGetRuleDetailListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIGetRuleDetailListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetCheckTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelManagerInterfaceIGetSourcePlatformRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// SearchContent

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
}

func (r *LabelManagerInterfaceIGetSourcePlatformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LabelManagerInterfaceIGetSourcePlatformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportIpFlowInterfaceIGetIpFlowOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReportIpFlowInterfaceIGetIpFlowOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReportIpFlowInterfaceIGetIpFlowOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportNetFlowInterfaceIGetNetFlowViewDetailRequest struct {
	*tchttp.BaseRequest

	// 趋势 trend, 对比 compare

	Type *string `json:"Type,omitempty" name:"Type"`
	// 时间

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 查看单日，可以为空或者与start_date相等，否则必须是有效的时间区间

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *ReportNetFlowInterfaceIGetNetFlowViewDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReportNetFlowInterfaceIGetNetFlowViewDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceISetConfigTaskSendStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AccessRuleInterfaceISetConfigTaskSendStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessRuleInterfaceISetConfigTaskSendStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetAlarmLogListRequest struct {
	*tchttp.BaseRequest

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *DayOverviewInterfaceIGetAlarmLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetAlarmLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetSystemRawLogListRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 时间

	AlarmTime []*string `json:"AlarmTime,omitempty" name:"AlarmTime"`
	// 名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 查询标识

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 当前页数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *TceManagerInterfaceIGetSystemRawLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetSystemRawLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateAzIdsInterceptConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIUpdateAzIdsInterceptConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIUpdateAzIdsInterceptConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIStopFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIStopFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIStopFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiBlockStatusRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIGetApiBlockStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetApiBlockStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceISetUnreadMsgToReadRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *MsgInterfaceISetUnreadMsgToReadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MsgInterfaceISetUnreadMsgToReadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RawlogSearchInterfaceIQueryRwLogRequest struct {
	*tchttp.BaseRequest

	// 组之间的操作符,JSON格式

	GroupBetweenCond *string `json:"GroupBetweenCond,omitempty" name:"GroupBetweenCond"`
	// 来源字段

	DataSrc *string `json:"DataSrc,omitempty" name:"DataSrc"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询标识

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 当前页数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 柱状图按什么单位聚合

	AggreUnit *string `json:"AggreUnit,omitempty" name:"AggreUnit"`
	// JSON格式

	GroupVal *string `json:"GroupVal,omitempty" name:"GroupVal"`
	// 多余参数

	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RawlogSearchInterfaceIQueryRwLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RawlogSearchInterfaceIQueryRwLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchChangeWorkStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIBatchChangeWorkStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchChangeWorkStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIStartNetworkRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIStartNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIStartNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIEditRuleRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// Content

	Content *string `json:"Content,omitempty" name:"Content"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// IsOverwrite

	IsOverwrite *string `json:"IsOverwrite,omitempty" name:"IsOverwrite"`
	// MatchOperation

	MatchOperation *string `json:"MatchOperation,omitempty" name:"MatchOperation"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Reboot

	Reboot *bool `json:"Reboot,omitempty" name:"Reboot"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// SourcePlatform

	SourcePlatform *string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// ValidDuration

	ValidDuration *int64 `json:"ValidDuration,omitempty" name:"ValidDuration"`
}

func (r *PreciseAccessInterfaceIEditRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIEditRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceIUpdateRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AccessRuleInterfaceIUpdateRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessRuleInterfaceIUpdateRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogHistogramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetThreatRawLogHistogramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIDeleteApiConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIDeleteApiConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIDeleteApiConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetBlockTimeRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *BlockConfigInterfaceIGetBlockTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetBlockTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelManagerInterfaceIAddSourcePlatformRequest struct {
	*tchttp.BaseRequest

	// Label

	Label *string `json:"Label,omitempty" name:"Label"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *LabelManagerInterfaceIAddSourcePlatformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LabelManagerInterfaceIAddSourcePlatformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIChangeBlockStatusRequest struct {
	*tchttp.BaseRequest

	// "off/on"  off表示关，on表示开

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIChangeBlockStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIChangeBlockStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetRedirectURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetRedirectURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetRedirectURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIAddRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIAddRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIAddRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetBlockTrendStatInfoRequest struct {
	*tchttp.BaseRequest

	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *OverviewInterfaceIGetBlockTrendStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetBlockTrendStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetSystemRawLogStatisticInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetSystemRawLogStatisticInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetSystemRawLogStatisticInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAzIdsInterceptConfigData struct {

	// RegionId

	RegionId []*string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 拉黑时长

	BlackDuration *uint64 `json:"BlackDuration,omitempty" name:"BlackDuration"`
}

type BlockConfigInterfaceIGetSecretInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetSecretInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetSecretInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetBlockStatInfoRequest struct {
	*tchttp.BaseRequest

	// start_time

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// end_time

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *OverviewInterfaceIGetBlockStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetBlockStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIAddCheckTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreciseAccessInterfaceIAddCheckTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIAddCheckTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIListRuleRequest struct {
	*tchttp.BaseRequest

	// AddTime

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// AdvSerachContent

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// BeginTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// HitCount

	HitCount []*string `json:"HitCount,omitempty" name:"HitCount"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// SearchContent

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
	// Source

	Source []*string `json:"Source,omitempty" name:"Source"`
	// Status

	Status []*string `json:"Status,omitempty" name:"Status"`
	// ValidDuration

	ValidDuration []*string `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// ipVersion

	IpVersion []*string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *SystemRuleInterfaceIListRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIListRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogListRequest struct {
	*tchttp.BaseRequest

	// 精准匹配

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 威胁等级

	ThreatLevel []*string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`
	// 威胁类型

	ThreatType []*string `json:"ThreatType,omitempty" name:"ThreatType"`
	// 应用层，传输层

	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`
	// 原因

	Reason []*string `json:"Reason,omitempty" name:"Reason"`
	// 状态

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 当前页面

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 查询标志

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 多余的参数

	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 用户选择的az列表

	AlarmAz []*string `json:"AlarmAz,omitempty" name:"AlarmAz"`
	// 高级查询

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// IpVersion

	IpVersion []*string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *AlarmInterfaceIGetThreatRawLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIAddApiConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIAddApiConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIAddApiConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIDeleteApiConfigRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIDeleteApiConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIDeleteApiConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIAddCheckTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIAddCheckTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIAddCheckTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIChangeWorkStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIChangeWorkStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIChangeWorkStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetOrResetSecretKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetOrResetSecretKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetOrResetSecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIStopRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SystemRuleInterfaceIStopRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIStopRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetThreatRawLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RawlogSearchInterfaceIGetDataFieldsRequest struct {
	*tchttp.BaseRequest

	// 数据字段

	DataSrc *string `json:"DataSrc,omitempty" name:"DataSrc"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RawlogSearchInterfaceIGetDataFieldsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RawlogSearchInterfaceIGetDataFieldsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIListRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SystemRuleInterfaceIListRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIListRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetFlowStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetFlowStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetFlowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBatchInsertRule struct {

	// 标题名称

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 类型，IP/URL/域名

	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`
	// 黑白名单内容

	MatchDetail []*string `json:"MatchDetail,omitempty" name:"MatchDetail"`
	// 方式,阻断\告警\放行

	MatchOperate *string `json:"MatchOperate,omitempty" name:"MatchOperate"`
	// 有效期

	ValidDuration *float64 `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// 黑白名单类型,

	Type *string `json:"Type,omitempty" name:"Type"`
}

type BlockConfigInterfaceIStartNetworkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIStartNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIStartNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetPlatformControlInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetPlatformControlInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetPlatformControlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIGetRuleStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SystemRuleInterfaceIGetRuleStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIGetRuleStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceGetRuleStatInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *OverviewInterfaceGetRuleStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceGetRuleStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshConfigInterfaceIGetRefreshConfigRequest struct {
	*tchttp.BaseRequest

	// Page

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RefreshConfigInterfaceIGetRefreshConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshConfigInterfaceIGetRefreshConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetOverviewImportantInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetOverviewImportantInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetOverviewImportantInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetRuleTrendStatInfoRequest struct {
	*tchttp.BaseRequest

	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *OverviewInterfaceIGetRuleTrendStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetRuleTrendStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIInsertRuleRequest struct {
	*tchttp.BaseRequest

	// 标题名称

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 白名单类型，ip...

	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`
	// 白名单内容

	MatchDetail []*string `json:"MatchDetail,omitempty" name:"MatchDetail"`
	// 阻断方式

	MatchOperate *string `json:"MatchOperate,omitempty" name:"MatchOperate"`
	// 有效期

	ValidDuration *float64 `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// 黑白名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 操作

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RuleConfigManagerInterfaceIInsertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIInsertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceGetStatisticRequest struct {
	*tchttp.BaseRequest

	// 见接口描述

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 见接口描述

	RuleAz []*string `json:"RuleAz,omitempty" name:"RuleAz"`
	// 见接口描述

	AdvSerachContent []*AdvSerachContent `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// 见接口描述

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 见接口描述

	ValidDuration []*string `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// 见接口描述

	HitHistory []*string `json:"HitHistory,omitempty" name:"HitHistory"`
	// 见接口描述

	HitWeek []*string `json:"HitWeek,omitempty" name:"HitWeek"`
	// 见接口描述

	HitDay []*string `json:"HitDay,omitempty" name:"HitDay"`
	// 见接口描述

	MatchOperation []*string `json:"MatchOperation,omitempty" name:"MatchOperation"`
	// 见接口描述

	IpType []*string `json:"IpType,omitempty" name:"IpType"`
	// 见接口描述

	SourcePlatform []*string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
}

func (r *GlobalAccessInterfaceGetStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceGetStatisticRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelManagerInterfaceIGetSourcePlatformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LabelManagerInterfaceIGetSourcePlatformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LabelManagerInterfaceIGetSourcePlatformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceIGetMsgStatisticInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MsgInterfaceIGetMsgStatisticInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MsgInterfaceIGetMsgStatisticInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceGetRuleStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceGetRuleStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceGetRuleStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetRuleStatInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *OverviewInterfaceIGetRuleStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetRuleStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RawlogSearchInterfaceIGetDataFieldsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RawlogSearchInterfaceIGetDataFieldsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RawlogSearchInterfaceIGetDataFieldsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetBlockLogListRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 高级查询下拉项

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// 告警来源

	AlarmSource []*string `json:"AlarmSource,omitempty" name:"AlarmSource"`
	// 状态

	Status []*string `json:"Status,omitempty" name:"Status"`
	// BeginTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// FlowType

	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// OccurTime

	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// Reason

	Reason []*string `json:"Reason,omitempty" name:"Reason"`
	// ThreatLevel

	ThreatLevel []*string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`
	// ThreatType

	ThreatType []*string `json:"ThreatType,omitempty" name:"ThreatType"`
	// IpVersion

	IpVersion []*string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *AlarmInterfaceIGetBlockLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetBlockLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackendTaskInterfaceIGetDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BackendTaskInterfaceIGetDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BackendTaskInterfaceIGetDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetAvailableCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetAvailableCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetAvailableCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetNetFlowRequest struct {
	*tchttp.BaseRequest

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *DayOverviewInterfaceIGetNetFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetNetFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIAddCheckTaskRequest struct {
	*tchttp.BaseRequest

	// Data

	Data *string `json:"Data,omitempty" name:"Data"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *GlobalAccessInterfaceIAddCheckTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIAddCheckTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportIpFlowInterfaceIGetIpFlowTableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReportIpFlowInterfaceIGetIpFlowTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReportIpFlowInterfaceIGetIpFlowTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchDeleteRuleRequest struct {
	*tchttp.BaseRequest

	// 删除的规则数组，JSON格式

	Data *string `json:"Data,omitempty" name:"Data"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RuleConfigManagerInterfaceIBatchDeleteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchDeleteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceIDeleteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AccessRuleInterfaceIDeleteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessRuleInterfaceIDeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetSystemConfigRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// AZ

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *AlarmInterfaceIGetSystemConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetSystemConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceISetApiConfigStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceISetApiConfigStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceISetApiConfigStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetAlarmReasonCntRequest struct {
	*tchttp.BaseRequest

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *DayOverviewInterfaceIGetAlarmReasonCntRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetAlarmReasonCntRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetPlatformControlInfoRequest struct {
	*tchttp.BaseRequest

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *DayOverviewInterfaceIGetPlatformControlInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetPlatformControlInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIGetRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreciseAccessInterfaceIGetRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIGetRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBatchDeleteRule struct {

	// 删除的规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 黑白名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type RuleConfigManagerInterfaceIDeleteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIDeleteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIDeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetRuleStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetRuleStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetRuleStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetSystemRawLogListRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 类型

	Type []*string `json:"Type,omitempty" name:"Type"`
	// 页数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 查询标志

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 时间

	AlarmTime []*string `json:"AlarmTime,omitempty" name:"AlarmTime"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// AdvSerachContent

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
}

func (r *AlarmInterfaceIGetSystemRawLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetSystemRawLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiConfigInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetApiConfigInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetApiConfigInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetOverviewImportantInfoRequest struct {
	*tchttp.BaseRequest

	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *DayOverviewInterfaceIGetOverviewImportantInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetOverviewImportantInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetNetoffConfigRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIGetNetoffConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetNetoffConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIAddRuleRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 无

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 无

	Type *string `json:"Type,omitempty" name:"Type"`
	// 无无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	ValidDuration *uint64 `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// 无

	MatchOperation *string `json:"MatchOperation,omitempty" name:"MatchOperation"`
	// 无

	SourcePlatform *string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// 无

	IsOverwrite *string `json:"IsOverwrite,omitempty" name:"IsOverwrite"`
	// 无

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// 无

	IpData *string `json:"IpData,omitempty" name:"IpData"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GlobalAccessInterfaceIAddRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIAddRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetRuleTrendStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetRuleTrendStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetRuleTrendStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIAddCheckTaskRequest struct {
	*tchttp.BaseRequest

	// Data

	Data *string `json:"Data,omitempty" name:"Data"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *PreciseAccessInterfaceIAddCheckTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIAddCheckTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIGetStatisticsRequest struct {
	*tchttp.BaseRequest

	// 见接口描述

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 见接口描述

	RuleAz []*string `json:"RuleAz,omitempty" name:"RuleAz"`
	// 见接口描述

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// 见接口描述

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 见接口描述

	ValidDuration []*string `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// 见接口描述

	HitHistory []*string `json:"HitHistory,omitempty" name:"HitHistory"`
	// 见接口描述

	HitWeek []*string `json:"HitWeek,omitempty" name:"HitWeek"`
	// 见接口描述

	HitDay []*string `json:"HitDay,omitempty" name:"HitDay"`
	// 见接口描述

	MatchOperation []*string `json:"MatchOperation,omitempty" name:"MatchOperation"`
	// 见接口描述

	IpType []*string `json:"IpType,omitempty" name:"IpType"`
	// 见接口描述

	SourcePlatform []*string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// AddTime

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// BeginTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// SearchContent

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *PreciseAccessInterfaceIGetStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIGetStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshConfigInterfaceIUpdateRefreshConfigRequest struct {
	*tchttp.BaseRequest

	// AutoRefresh

	AutoRefresh *uint64 `json:"AutoRefresh,omitempty" name:"AutoRefresh"`
	// Frequency

	Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`
	// Page

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RefreshConfigInterfaceIUpdateRefreshConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshConfigInterfaceIUpdateRefreshConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIListRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIListRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIListRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetBlockLogStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetBlockLogStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetBlockLogStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiBlockTrustThresholdRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIGetApiBlockTrustThresholdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetApiBlockTrustThresholdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIEditRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIEditRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIEditRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceIDeleteRuleRequest struct {
	*tchttp.BaseRequest

	// AccessType

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// Data

	Data *string `json:"Data,omitempty" name:"Data"`
	// NetSwitch

	NetSwitch *string `json:"NetSwitch,omitempty" name:"NetSwitch"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// InfoType

	InfoType *string `json:"InfoType,omitempty" name:"InfoType"`
	// InfoValue

	InfoValue *string `json:"InfoValue,omitempty" name:"InfoValue"`
	// Reboot

	Reboot *bool `json:"Reboot,omitempty" name:"Reboot"`
}

func (r *AccessRuleInterfaceIDeleteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessRuleInterfaceIDeleteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayOverviewInterfaceIGetNetFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DayOverviewInterfaceIGetNetFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DayOverviewInterfaceIGetNetFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIGetRuleListRequest struct {
	*tchttp.BaseRequest

	// 见接口描述

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 见接口描述

	RuleAz []*string `json:"RuleAz,omitempty" name:"RuleAz"`
	// 见接口描述

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// 见接口描述

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 见接口描述

	ValidDuration []*string `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// 见接口描述

	HitHistory []*string `json:"HitHistory,omitempty" name:"HitHistory"`
	// 见接口描述

	HitWeek []*string `json:"HitWeek,omitempty" name:"HitWeek"`
	// 见接口描述

	HitDay []*string `json:"HitDay,omitempty" name:"HitDay"`
	// 见接口描述

	MatchOperation []*string `json:"MatchOperation,omitempty" name:"MatchOperation"`
	// 见接口描述

	IpType []*string `json:"IpType,omitempty" name:"IpType"`
	// 见接口描述

	SourcePlatform []*string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// AddTime

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// BeginTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// SearchContent

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *PreciseAccessInterfaceIGetRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIGetRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetCheckTaskResultListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetAlarmStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetAlarmStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetAlarmStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetTOPStatInfoRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *OverviewInterfaceIGetTOPStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetTOPStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIBatchInsertRuleRequest struct {
	*tchttp.BaseRequest

	// 插入的规则，格式是JSON

	Data *string `json:"Data,omitempty" name:"Data"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RuleConfigManagerInterfaceIBatchInsertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIBatchInsertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetIdsInterceptStatusRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIGetIdsInterceptStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetIdsInterceptStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceISetApiConfigStatusRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *BlockConfigInterfaceISetApiConfigStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceISetApiConfigStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIStopNetworkRequest struct {
	*tchttp.BaseRequest

	// 断网原因

	BlockReason *string `json:"BlockReason,omitempty" name:"BlockReason"`
	// 断网原因html

	BlockReasonHtml *string `json:"BlockReasonHtml,omitempty" name:"BlockReasonHtml"`
	// 断网开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 断网结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIStopNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIStopNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIUpdateRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIUpdateRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIUpdateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetFlowStatusRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *TceManagerInterfaceIGetFlowStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetFlowStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackendTaskInterfaceIGetDataRequest struct {
	*tchttp.BaseRequest

	// BackendTaskId

	BackendTaskId *uint64 `json:"BackendTaskId,omitempty" name:"BackendTaskId"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BackendTaskInterfaceIGetDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BackendTaskInterfaceIGetDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIConfigBlockTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIConfigBlockTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIConfigBlockTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiConfigDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetApiConfigDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetApiConfigDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetBlockTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetBlockTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetBlockTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceISetApiBlockStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceISetApiBlockStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceISetApiBlockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIGetFilterItemsListInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RuleConfigManagerInterfaceIGetFilterItemsListInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIGetFilterItemsListInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetThreatRawLogStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// 查询内容

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询时间

	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime"`
	// 应用层，传输层

	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`
	// 威胁等级，低危，中危，高危，严重，致命

	ThreatLevel *string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`
	// 黑名单，黑客攻击

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 拦截，未拦截

	Status *string `json:"Status,omitempty" name:"Status"`
	// 执行标志

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 当前页

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// webshell检测

	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *TceManagerInterfaceIGetThreatRawLogStatisticInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetThreatRawLogStatisticInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogHistogramRequest struct {
	*tchttp.BaseRequest

	// 根据ip，端口，cgi，host等信息搜索，精准匹配

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 威胁等级，低危，中危，高危，严重，致命

	ThreatLevel []*string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`
	// 威胁类型

	ThreatType []*string `json:"ThreatType,omitempty" name:"ThreatType"`
	// 层级

	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`
	// 原因

	Reason []*string `json:"Reason,omitempty" name:"Reason"`
	// 拦截，未拦截

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 柱状图的时间类型

	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`
	// 多余的参数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 多余的参数

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 多余的参数

	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime"`
	// 多余的参数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 高级查询

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// 用户选择的az列表，数据类型为数组

	AlarmAz []*string `json:"AlarmAz,omitempty" name:"AlarmAz"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *AlarmInterfaceIGetThreatRawLogHistogramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogHistogramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiConfigDetailRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIGetApiConfigDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetApiConfigDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetCheckDistribStatInfoRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *OverviewInterfaceIGetCheckDistribStatInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetCheckDistribStatInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GlobalAccessInterfaceIGetRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIGetRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceIGetConfigTaskInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AccessRuleInterfaceIGetConfigTaskInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessRuleInterfaceIGetConfigTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetThreatRawLogStatisticInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AlarmInterfaceIGetThreatRawLogStatisticInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetThreatRawLogStatisticInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIUpdateApiBlockTrustThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateAzIdsInterceptStatusRequest struct {
	*tchttp.BaseRequest

	// {
	//     "region_id": "region_1",
	//     "az": ["az_1"],
	//     "intercept_status": "0", # 状态："0"-开，"1"-关
	//     }

	Data *string `json:"Data,omitempty" name:"Data"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIUpdateAzIdsInterceptStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIUpdateAzIdsInterceptStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetCheckTaskResultRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetSystemRawLogStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 时间

	AlarmTime []*string `json:"AlarmTime,omitempty" name:"AlarmTime"`
	// 名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 查询标识

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 当前页数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *TceManagerInterfaceIGetSystemRawLogStatisticInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetSystemRawLogStatisticInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TceManagerInterfaceIGetThreatRawLogStatisticInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TceManagerInterfaceIGetThreatRawLogStatisticInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TceManagerInterfaceIGetThreatRawLogStatisticInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetCurrentBlockStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetCurrentBlockStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetCurrentBlockStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateAzIdsInterceptStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIUpdateAzIdsInterceptStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIUpdateAzIdsInterceptStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgInterfaceIGetMsgListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MsgInterfaceIGetMsgListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MsgInterfaceIGetMsgListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateIdsInterceptStatusRequest struct {
	*tchttp.BaseRequest

	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 拦截状态: "1"->关闭, "0"->打开

	Status *string `json:"Status,omitempty" name:"Status"`
	// InterceptStatus

	InterceptStatus *string `json:"InterceptStatus,omitempty" name:"InterceptStatus"`
}

func (r *BlockConfigInterfaceIUpdateIdsInterceptStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIUpdateIdsInterceptStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIUpdateIdsInterceptStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIUpdateIdsInterceptStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIUpdateIdsInterceptStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetBlockStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetBlockStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetBlockStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIEditApiConfigRequest struct {
	*tchttp.BaseRequest

	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// WindowCount

	WindowCount *uint64 `json:"WindowCount,omitempty" name:"WindowCount"`
	// WindowTime

	WindowTime *string `json:"WindowTime,omitempty" name:"WindowTime"`
}

func (r *BlockConfigInterfaceIEditApiConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIEditApiConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemRuleInterfaceIGetRuleDetailListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SystemRuleInterfaceIGetRuleDetailListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SystemRuleInterfaceIGetRuleDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceGetBlockLogHistogramRequest struct {
	*tchttp.BaseRequest

	// 接口描述

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 接口描述

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 接口描述

	AdvSerachContent []*AdvSerachContent `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// 接口描述

	AlarmSource []*string `json:"AlarmSource,omitempty" name:"AlarmSource"`
	// 接口描述

	Status []*string `json:"Status,omitempty" name:"Status"`
}

func (r *AlarmInterfaceGetBlockLogHistogramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceGetBlockLogHistogramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetSystemRawLogStatisticInfoRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 类型

	Type []*string `json:"Type,omitempty" name:"Type"`
	// 时间

	AlarmTime []*string `json:"AlarmTime,omitempty" name:"AlarmTime"`
	// 查询标识

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// 每页条数

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// 当前页数

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// AdvSerachContent

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
}

func (r *AlarmInterfaceIGetSystemRawLogStatisticInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetSystemRawLogStatisticInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreciseAccessInterfaceIAddRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreciseAccessInterfaceIAddRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreciseAccessInterfaceIAddRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleConfigManagerInterfaceIUpdateRuleRequest struct {
	*tchttp.BaseRequest

	// 修改后的标题名称

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 需要修改得规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 白名单类型，ip...

	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`
	// 白名单内容

	MatchDetail []*string `json:"MatchDetail,omitempty" name:"MatchDetail"`
	// 阻断方式

	MatchOperate *string `json:"MatchOperate,omitempty" name:"MatchOperate"`
	// 有效期(整数)，有效时长，**注意永久用-1表示**

	ValidDuration *float64 `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// 黑白名单类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 操作

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 添加时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *RuleConfigManagerInterfaceIUpdateRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RuleConfigManagerInterfaceIUpdateRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceIGetConfigTaskInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// 无

	Type *string `json:"Type,omitempty" name:"Type"`
	// 无

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *AccessRuleInterfaceIGetConfigTaskInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessRuleInterfaceIGetConfigTaskInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRuleInterfaceIUpdateRuleStatusRequest struct {
	*tchttp.BaseRequest

	// AccessType

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// Data

	Data *string `json:"Data,omitempty" name:"Data"`
	// InfoType

	InfoType *string `json:"InfoType,omitempty" name:"InfoType"`
	// InfoValue

	InfoValue *string `json:"InfoValue,omitempty" name:"InfoValue"`
	// NetSwitch

	NetSwitch *string `json:"NetSwitch,omitempty" name:"NetSwitch"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Reboot

	Reboot *bool `json:"Reboot,omitempty" name:"Reboot"`
}

func (r *AccessRuleInterfaceIUpdateRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessRuleInterfaceIUpdateRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInterfaceIGetBlockLogHistogramRequest struct {
	*tchttp.BaseRequest

	// BeginTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// FlowType

	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// OccurTime

	OccurTime []*string `json:"OccurTime,omitempty" name:"OccurTime"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// Reason

	Reason []*string `json:"Reason,omitempty" name:"Reason"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// Status

	Status []*string `json:"Status,omitempty" name:"Status"`
	// ThreatLevel

	ThreatLevel []*string `json:"ThreatLevel,omitempty" name:"ThreatLevel"`
	// ThreatType

	ThreatType []*string `json:"ThreatType,omitempty" name:"ThreatType"`
	// TimeType

	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`
	// AdvSerachContent

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// AlarmSource

	AlarmSource []*string `json:"AlarmSource,omitempty" name:"AlarmSource"`
	// IpVersion

	IpVersion []*string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *AlarmInterfaceIGetBlockLogHistogramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AlarmInterfaceIGetBlockLogHistogramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewInterfaceIGetCheckDistribStatInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverviewInterfaceIGetCheckDistribStatInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverviewInterfaceIGetCheckDistribStatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetApiBlockTrustThresholdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetApiBlockTrustThresholdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetApiBlockTrustThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetCheckTaskResultListRequest struct {
	*tchttp.BaseRequest

	// CheckStatus

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// RegionId

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIGetCheckTaskResultListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalAccessInterfaceIGetRuleListRequest struct {
	*tchttp.BaseRequest

	// 见接口描述

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 见接口描述

	RuleAz []*string `json:"RuleAz,omitempty" name:"RuleAz"`
	// 见接口描述

	AdvSerachContent *string `json:"AdvSerachContent,omitempty" name:"AdvSerachContent"`
	// 见接口描述

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 见接口描述

	ValidDuration []*string `json:"ValidDuration,omitempty" name:"ValidDuration"`
	// 见接口描述

	HitHistory []*string `json:"HitHistory,omitempty" name:"HitHistory"`
	// 见接口描述

	HitWeek []*string `json:"HitWeek,omitempty" name:"HitWeek"`
	// 见接口描述

	HitDay []*string `json:"HitDay,omitempty" name:"HitDay"`
	// 见接口描述

	MatchOperation []*string `json:"MatchOperation,omitempty" name:"MatchOperation"`
	// 见接口描述

	IpType []*string `json:"IpType,omitempty" name:"IpType"`
	// 见接口描述

	SourcePlatform []*string `json:"SourcePlatform,omitempty" name:"SourcePlatform"`
	// AddTime

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
	// BeginTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// CurrentPage

	CurrentPage *uint64 `json:"CurrentPage,omitempty" name:"CurrentPage"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ExcuteFlag

	ExcuteFlag *uint64 `json:"ExcuteFlag,omitempty" name:"ExcuteFlag"`
	// InputCond

	InputCond *string `json:"InputCond,omitempty" name:"InputCond"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
	// PerPage

	PerPage *uint64 `json:"PerPage,omitempty" name:"PerPage"`
	// SearchContent

	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Az

	Az []*string `json:"Az,omitempty" name:"Az"`
	// IpVersion

	IpVersion []*string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *GlobalAccessInterfaceIGetRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GlobalAccessInterfaceIGetRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshConfigInterfaceIGetRefreshConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshConfigInterfaceIGetRefreshConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshConfigInterfaceIGetRefreshConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetCurrentBlockStateRequest struct {
	*tchttp.BaseRequest

	// 无

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *BlockConfigInterfaceIGetCurrentBlockStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetCurrentBlockStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockConfigInterfaceIGetIdsInterceptStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockConfigInterfaceIGetIdsInterceptStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockConfigInterfaceIGetIdsInterceptStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
