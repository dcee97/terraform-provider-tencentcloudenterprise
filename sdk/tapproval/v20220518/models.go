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

package v20220518

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DeleteApprovalFlowRequest struct {
	*tchttp.BaseRequest

	// 审批流id

	FlowID *uint64 `json:"FlowID,omitempty" name:"FlowID"`
}

func (r *DeleteApprovalFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApprovalFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StageParam struct {

	// stage id

	StageID *uint64 `json:"StageID,omitempty" name:"StageID"`
	// 节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// false 会签， true或签

	SingleSeal *bool `json:"SingleSeal,omitempty" name:"SingleSeal"`
	// 审批租户uin列表

	Approvers []*string `json:"Approvers,omitempty" name:"Approvers"`
	// 节点序号

	SerialNumber *uint64 `json:"SerialNumber,omitempty" name:"SerialNumber"`
}

type ApiInfo struct {

	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type CreateCustomerFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomerFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApprovalFlowStatusRequest struct {
	*tchttp.BaseRequest

	// v3接口参数

	ApprovalApiInfo *ApiInfo `json:"ApprovalApiInfo,omitempty" name:"ApprovalApiInfo"`
}

func (r *QueryApprovalFlowStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApprovalFlowStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawApplicationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WithdrawApplicationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WithdrawApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Order struct {

	// 排序字段英文名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 是否降序，true-升序，false-降序

	IsDesc *bool `json:"IsDesc,omitempty" name:"IsDesc"`
}

type OperateFlowStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperateFlowStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperateFlowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActionSetRequest struct {
	*tchttp.BaseRequest

	// 查询参数对象

	Params *QueryActionParma `json:"Params,omitempty" name:"Params"`
}

func (r *QueryActionSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryActionSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomerApprovalDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCustomerApprovalDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomerApprovalDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Scope struct {

	// 业务接口id

	ActionID *uint64 `json:"ActionID,omitempty" name:"ActionID"`
	// 作用用户列表

	Users []*UserScope `json:"Users,omitempty" name:"Users"`
}

type ModifyApprovalFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApprovalFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApprovalFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApprovalFlowStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// data

		Data *IsOpenAttr `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApprovalFlowStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApprovalFlowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApprovalFlowRequest struct {
	*tchttp.BaseRequest

	// 审批流名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 审批流描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 可修改的action schema

	SchemaProps *string `json:"SchemaProps,omitempty" name:"SchemaProps"`
	// 是否激活

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
	// 关联业务id列表

	ActionIDs []*uint64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
	// 节点列表

	Stages []*StageParam `json:"Stages,omitempty" name:"Stages"`
	// 授权范围

	Scopes []*Scope `json:"Scopes,omitempty" name:"Scopes"`
	// 审批流id

	FlowID *uint64 `json:"FlowID,omitempty" name:"FlowID"`
}

func (r *ModifyApprovalFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApprovalFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterAttr struct {

	// 参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作符, 支持equal, like

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetApprovedSetRequest struct {
	*tchttp.BaseRequest

	// 业务接口(中文名)

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 申请人uin

	ApplicantUin *string `json:"ApplicantUin,omitempty" name:"ApplicantUin"`
	// 申请人账号

	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`
	// 申请人所属主账号

	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
	// 理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 审批单状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 审批单id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 排序字段信息

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
}

func (r *GetApprovedSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApprovedSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateFlowStatusRequest struct {
	*tchttp.BaseRequest

	// 审批流id

	FlowID *uint64 `json:"FlowID,omitempty" name:"FlowID"`
	// true启用, false审批

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
}

func (r *OperateFlowStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperateFlowStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlowDetailRequest struct {
	*tchttp.BaseRequest

	// 审批流id

	FlowID *uint64 `json:"FlowID,omitempty" name:"FlowID"`
}

func (r *GetFlowDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlowDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlowDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFlowDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlowDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActionSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryActionSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryActionSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActionParma struct {

	// id 列表

	ActionIDs []*uint64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
	// action名称

	Action *string `json:"Action,omitempty" name:"Action"`
	// 所属模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 业务名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 接口对应模块(中文名)

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 接口对应产品(中文名)

	YunProductName *string `json:"YunProductName,omitempty" name:"YunProductName"`
	// 排序字段

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
}

type GetApprovedSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApprovedSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApprovedSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApplicationFormRequest struct {
	*tchttp.BaseRequest

	// 业务接口(中文名)

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 申请人uin

	ApplicantUin *string `json:"ApplicantUin,omitempty" name:"ApplicantUin"`
	// 申请人账号

	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`
	// 申请人所属主账号

	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
	// 理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 审批单id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 审批单状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 排序字段信息

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
}

func (r *QueryApplicationFormRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApplicationFormRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCurrApprovalDocRequest struct {
	*tchttp.BaseRequest

	// 业务接口(中文名)

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 申请人uin

	ApplicantUin *string `json:"ApplicantUin,omitempty" name:"ApplicantUin"`
	// 申请人账号

	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`
	// 申请人所属主账号

	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
	// 理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 审批单状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 审批单id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 排序字段信息

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
}

func (r *QueryCurrApprovalDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCurrApprovalDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerFlowRequest struct {
	*tchttp.BaseRequest

	// 审批流名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 审批流描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 可修改的action schema

	SchemaProps *string `json:"SchemaProps,omitempty" name:"SchemaProps"`
	// 是否激活

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
	// 关联业务id列表

	ActionIDs []*uint64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
	// 节点列表

	Stages []*StageParam `json:"Stages,omitempty" name:"Stages"`
	// 授权范围

	Scopes []*Scope `json:"Scopes,omitempty" name:"Scopes"`
}

func (r *CreateCustomerFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitApprovalDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitApprovalDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitApprovalDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApprovalFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApprovalFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApprovalFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApplicationFormResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApplicationFormResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApplicationFormResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsOpenAttr struct {

	// 是否存在已开启的审批流， true：存在， false： 不存在

	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
}

type PerformApprovalRequest struct {
	*tchttp.BaseRequest

	// 当前节点序号

	StageSerialNum *uint64 `json:"StageSerialNum,omitempty" name:"StageSerialNum"`
	// 审批单id

	PaperID *uint64 `json:"PaperID,omitempty" name:"PaperID"`
	// 审批结果12: 拒绝，14: 同意

	Operate *uint64 `json:"Operate,omitempty" name:"Operate"`
	// 审批意见

	Opinion *string `json:"Opinion,omitempty" name:"Opinion"`
	// 审批时修改的请求参数

	Request *string `json:"Request,omitempty" name:"Request"`
}

func (r *PerformApprovalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PerformApprovalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomerApprovalDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCustomerApprovalDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomerApprovalDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowSetRequest struct {
	*tchttp.BaseRequest

	// 审批流名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否启用 0-不启用，1-启用

	IsActive *int64 `json:"IsActive,omitempty" name:"IsActive"`
	// 排序字段信息

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
	// 过滤参数列表

	Filter []*FilterAttr `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeFlowSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PerformApprovalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PerformApprovalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PerformApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomerActionScopeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCustomerActionScopeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomerActionScopeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitApprovalDocRequest struct {
	*tchttp.BaseRequest

	// 当前节点id

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 审批单id

	PaperID *uint64 `json:"PaperID,omitempty" name:"PaperID"`
	// 所属主账号用户名

	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
}

func (r *SubmitApprovalDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitApprovalDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCurrApprovalDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCurrApprovalDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCurrApprovalDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomerActionScopeRequest struct {
	*tchttp.BaseRequest

	// 业务id列表

	ActionIDs []*uint64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
}

func (r *QueryCustomerActionScopeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomerActionScopeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomerApprovalDetailRequest struct {
	*tchttp.BaseRequest

	// 审批单id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *QueryCustomerApprovalDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomerApprovalDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomerApprovalDocRequest struct {
	*tchttp.BaseRequest

	// 业务接口(中文名)

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 申请人uin

	ApplicantUin *string `json:"ApplicantUin,omitempty" name:"ApplicantUin"`
	// 申请人账号

	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`
	// 申请人所属主账号

	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
	// 理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 审批单状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 审批单id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 排序字段信息

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
	// 关联的接口信息

	ApprovalApiInfo *ApiInfo `json:"ApprovalApiInfo,omitempty" name:"ApprovalApiInfo"`
}

func (r *QueryCustomerApprovalDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomerApprovalDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserScope struct {

	// -1，全体租户或具体租户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 是否为子账号

	IsSubAccount *bool `json:"IsSubAccount,omitempty" name:"IsSubAccount"`
}

type WithdrawApplicationRequest struct {
	*tchttp.BaseRequest
}

func (r *WithdrawApplicationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WithdrawApplicationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
