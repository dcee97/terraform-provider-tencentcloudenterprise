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

package v20190919

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type Instance struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 私有网络子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 私有网络子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 任务类型

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 任务描述

	TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`
	// 虚拟IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 虚拟端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 付费类型，postpaid - 后付费，prepaid - 预付费

	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
	// 计算节点列表

	CNNodes []*InstanceNode `json:"CNNodes,omitempty" name:"CNNodes"`
	// 数据节点列表

	DNNodes []*InstanceNode `json:"DNNodes,omitempty" name:"DNNodes"`
	// 项目ID，为空表示没有所属项目

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名，为空表示没有所属项目

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表。

		Items []*Instance `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 按一个或多个实例Id进行查询，每次最多请求100个。指定该参数后，Filters参数失效。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 过滤条件，目前支持InstanceName、Vip、InstanceId、TagKey。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段，支持 InstanceId,InstanceName,CreateTime,PeriodEndTime。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序方式，支持ASC和DESC，默认为ASC。

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 偏移量，默认为0。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回数量，默认为20，最大值为100。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 实例状态信息，支持 creating, online, isolate

	Status *string `json:"Status,omitempty" name:"Status"`
	// 根据计费类型过滤，支持postpaid后付费，prepaid预付费两种。

	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceNode struct {

	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点cpu

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 节点内存

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 节点数据盘大小

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 节点唯一ID

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 已用数据存储空间

	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`
	// 节点在所在节点组中的类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type Filter struct {

	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}
