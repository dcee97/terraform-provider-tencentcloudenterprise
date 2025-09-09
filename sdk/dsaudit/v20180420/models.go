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

package v20180420

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeProductTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规格

		ProductTypes []*ProductType `json:"ProductTypes,omitempty" name:"ProductTypes"`
		// 时长

		TimeSpan []*TimeSpan `json:"TimeSpan,omitempty" name:"TimeSpan"`
		// 售卖方式,0:计量,1:预付费,2:后付费

		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
		// 应用状态

		AppStatus *string `json:"AppStatus,omitempty" name:"AppStatus"`
		// 磁盘类型

		DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
		// 地域镜像信息

		RegionImageInfos []*RegionImageItem `json:"RegionImageInfos,omitempty" name:"RegionImageInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSaleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSaleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*string `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 服务总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
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

type DescribePriceRequest struct {
	*tchttp.BaseRequest

	// 无

	ResInfo *string `json:"ResInfo,omitempty" name:"ResInfo"`
}

func (r *DescribePriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegisterUrlRequest struct {
	*tchttp.BaseRequest

	// 无

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 无

	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DescribeRegisterUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegisterUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegisterUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// url

		Url *string `json:"Url,omitempty" name:"Url"`
		// token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegisterUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegisterUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAddressRequest struct {
	*tchttp.BaseRequest

	// 无

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 无

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 无

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 无

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 无

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 无

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *ModifyInstanceAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNameInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetNameInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNameInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 价格信息

		Price []*string `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoodsDetail struct {

	// Pid

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 商品数量

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 地区id

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计时单位

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 购买时长

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

type TimeSpan struct {

	// 名称

	TimeName *string `json:"TimeName,omitempty" name:"TimeName"`
	// 月数

	TimeMonth *uint64 `json:"TimeMonth,omitempty" name:"TimeMonth"`
}

type RegionImageItem struct {

	// 地域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域代码

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 地域支持的架构

	MixFlag *uint64 `json:"MixFlag,omitempty" name:"MixFlag"`
	// 镜像列表

	ImageList []*ImageDesc `json:"ImageList,omitempty" name:"ImageList"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 下单参数

	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 分页数量，默认值为1000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 地域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductType struct {

	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// PID

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 规格名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 规格代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// Cpu核数

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 系统盘大小

	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
	// 数据盘大小

	DataDiskSize *uint64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`
}

type DescribeSaleInfoRequest struct {
	*tchttp.BaseRequest

	// 镜像架构

	ImageArch *string `json:"ImageArch,omitempty" name:"ImageArch"`
	// 地域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *DescribeSaleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSaleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNameInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *SetNameInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNameInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInstanceRequest struct {
	*tchttp.BaseRequest

	// 无

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StartInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品规格列表

		ProductTypes []*ProductType `json:"ProductTypes,omitempty" name:"ProductTypes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageDesc struct {

	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像地域

	ImageRegion *string `json:"ImageRegion,omitempty" name:"ImageRegion"`
	// 镜像架构

	ImageArch *string `json:"ImageArch,omitempty" name:"ImageArch"`
}
