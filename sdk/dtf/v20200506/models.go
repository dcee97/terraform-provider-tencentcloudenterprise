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

package v20200506

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeGroupStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事务分组统计信息

		Result *GroupStatistics `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTransactionsRequest struct {
	*tchttp.BaseRequest

	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 事务开始时间查询起始时间戳，UTC，精确到毫秒

	TransactionBeginFrom *int64 `json:"TransactionBeginFrom,omitempty" name:"TransactionBeginFrom"`
	// 事务开始时间查询截止时间戳，UTC，精确到毫秒

	TransactionBeginTo *int64 `json:"TransactionBeginTo,omitempty" name:"TransactionBeginTo"`
	// 仅查询异常状态的事务，true：仅查询异常，false或不传入：查询所有

	SearchError *bool `json:"SearchError,omitempty" name:"SearchError"`
	// 主事务ID，不传入时查询全量，高优先级

	TransactionId *int64 `json:"TransactionId,omitempty" name:"TransactionId"`
	// 主事务ID列表，不传入时查询全量，低优先级

	TransactionIdList []*int64 `json:"TransactionIdList,omitempty" name:"TransactionIdList"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 起始偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTransactionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTransactionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PagedGroup struct {

	// 总条数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 事务分组列表

	Content []*Group `json:"Content,omitempty" name:"Content"`
}

type DescribeBranchStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分支事务统计信息

		Result []*BranchStatistics `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBranchStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBranchStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupStatisticsRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间戳（包含）

	SearchBegin *int64 `json:"SearchBegin,omitempty" name:"SearchBegin"`
	// 查询截止时间戳（不包含）

	SearchEnd *int64 `json:"SearchEnd,omitempty" name:"SearchEnd"`
	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTransactionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主事务分页列表

		Result *PagedTransaction `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTransactionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTransactionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true：更新成功；false：更新失败

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBranchesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分支事务分页列表

		Result *PagedBranch `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBranchesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBranchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllGroupStatistics struct {

	// 主事务总数

	TransactionQuantity *int64 `json:"TransactionQuantity,omitempty" name:"TransactionQuantity"`
	// 分支事务总数

	BranchQuantity *int64 `json:"BranchQuantity,omitempty" name:"BranchQuantity"`
	// 异常主事务数

	ErrorTransactionQuantity *int64 `json:"ErrorTransactionQuantity,omitempty" name:"ErrorTransactionQuantity"`
	// 事务分组数

	GroupQuantity *int64 `json:"GroupQuantity,omitempty" name:"GroupQuantity"`
}

type Group struct {

	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 事务分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 事务分组描述

	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
	// 创建时间戳，UTC，精确到毫秒

	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`
	// 更新时间戳，UTC，精确到毫秒

	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 删除标识，true：可以删除；false：不可删除

	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`
	// 集群列表

	Servers []*string `json:"Servers,omitempty" name:"Servers"`
	// 协调器集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 协调器集群类型

	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`
	// 限流值

	Ratelimit *int64 `json:"Ratelimit,omitempty" name:"Ratelimit"`
}

type DescribeAllGroupStatisticsRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间戳（包含）

	SearchBegin *int64 `json:"SearchBegin,omitempty" name:"SearchBegin"`
	// 查询截止时间戳（不包含）

	SearchEnd *int64 `json:"SearchEnd,omitempty" name:"SearchEnd"`
}

func (r *DescribeAllGroupStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllGroupStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryTransactionError struct {

	// 主事务ID

	TransactionId *int64 `json:"TransactionId,omitempty" name:"TransactionId"`
	// 异常信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type DescribeBranchesRequest struct {
	*tchttp.BaseRequest

	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 主事务ID

	TransactionId *int64 `json:"TransactionId,omitempty" name:"TransactionId"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 起始偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBranchesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBranchesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事务分组分页列表

		Result *PagedGroup `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 事务分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 事务分组描述

	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
	// 所属vpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *CreateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Transaction struct {

	// 主事务ID

	TransactionId *int64 `json:"TransactionId,omitempty" name:"TransactionId"`
	// 主事务开始时间戳，UTC，精确到毫秒

	TransactionBegin *int64 `json:"TransactionBegin,omitempty" name:"TransactionBegin"`
	// 主事务结束时间戳，UTC，精确到毫秒

	TransactionEnd *int64 `json:"TransactionEnd,omitempty" name:"TransactionEnd"`
	// 主事务提交时间戳，UTC，精确到毫秒

	TransactionCommit *int64 `json:"TransactionCommit,omitempty" name:"TransactionCommit"`
	// 主事务回滚时间戳，UTC，精确到毫秒

	TransactionRollback *int64 `json:"TransactionRollback,omitempty" name:"TransactionRollback"`
	// 主事务异常停止时间戳，UTC，精确到毫秒

	TransactionError *int64 `json:"TransactionError,omitempty" name:"TransactionError"`
	// 主事务超时时长，单位毫秒

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 主事务状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 主事务结束标识

	EndFlag *int64 `json:"EndFlag,omitempty" name:"EndFlag"`
	// 主事务超时标识

	TimeoutFlag *int64 `json:"TimeoutFlag,omitempty" name:"TimeoutFlag"`
	// 异常信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 主事务来源服务标识

	Server *string `json:"Server,omitempty" name:"Server"`
	// 分支事务数量

	BranchQuantity *int64 `json:"BranchQuantity,omitempty" name:"BranchQuantity"`
	// 重试标识：true：可以重试；false：不可重试

	RetryFlag *bool `json:"RetryFlag,omitempty" name:"RetryFlag"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PagedBranch struct {

	// 总条数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 分支事务分组列表

	Content []*Branch `json:"Content,omitempty" name:"Content"`
}

type Branch struct {

	// 分支事务ID

	BranchId *int64 `json:"BranchId,omitempty" name:"BranchId"`
	// 父级分支事务ID，顶级时返回null

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// 分支事务开始时间戳，UTC，精确到毫秒

	BranchBegin *int64 `json:"BranchBegin,omitempty" name:"BranchBegin"`
	// 分支事务结束时间戳，UTC，精确到毫秒

	BranchEnd *int64 `json:"BranchEnd,omitempty" name:"BranchEnd"`
	// 分支事务类型

	BranchType *int64 `json:"BranchType,omitempty" name:"BranchType"`
	// 分支事务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分支事务状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 分支事务来源服务标识

	Server *string `json:"Server,omitempty" name:"Server"`
	// 分支运行参数，Json格式字符串

	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
	// 主事务ID

	TransactionId *int64 `json:"TransactionId,omitempty" name:"TransactionId"`
}

type PagedTransaction struct {

	// 总条数，特定在该接口中总是会返回null

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 主事务分组列表

	Content []*Transaction `json:"Content,omitempty" name:"Content"`
}

type RetryAllTransactionsRequest struct {
	*tchttp.BaseRequest

	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *RetryAllTransactionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryAllTransactionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryTransactionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 触发结果

		Result []*RetryTransactionError `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryTransactionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryTransactionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BranchStatistics struct {

	// 分支事务类型

	BranchType *int64 `json:"BranchType,omitempty" name:"BranchType"`
	// 分支事务数

	Quantity *int64 `json:"Quantity,omitempty" name:"Quantity"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true：删除成功；false：删除失败

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBranchStatisticsRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间戳（包含）

	SearchBegin *int64 `json:"SearchBegin,omitempty" name:"SearchBegin"`
	// 查询截止时间戳（不包含）

	SearchEnd *int64 `json:"SearchEnd,omitempty" name:"SearchEnd"`
	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeBranchStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBranchStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest

	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 事务分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 事务分组描述

	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true：创建成功；false：创建失败

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryTransactionsRequest struct {
	*tchttp.BaseRequest

	// 主事务ID列表

	TransactionIdList []*int64 `json:"TransactionIdList,omitempty" name:"TransactionIdList"`
	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *RetryTransactionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryTransactionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllGroupStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有事务分组统计信息

		Result *AllGroupStatistics `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllGroupStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllGroupStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsRequest struct {
	*tchttp.BaseRequest

	// 事务分组ID，不传入时查询全量，高优先级

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 事务分组ID列表，不传入时查询全量，低优先级

	GroupIdList []*string `json:"GroupIdList,omitempty" name:"GroupIdList"`
	// 模糊查询，对事务分组ID，事务分组名称，事务分组描述字段进行模糊查询，不传入时查询全量

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 起始偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryAllTransactionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true：触发重试成功；false：触发重试失败

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryAllTransactionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryAllTransactionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupStatistics struct {

	// 事务分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 主事务总数

	TransactionQuantity *int64 `json:"TransactionQuantity,omitempty" name:"TransactionQuantity"`
	// 分支事务总数

	BranchQuantity *int64 `json:"BranchQuantity,omitempty" name:"BranchQuantity"`
	// 异常主事务数

	ErrorTransactionQuantity *int64 `json:"ErrorTransactionQuantity,omitempty" name:"ErrorTransactionQuantity"`
}
