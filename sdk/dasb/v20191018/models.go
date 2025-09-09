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

package v20191018

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeDasbCvmInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Dasb主机信息

		InstanceSet []*DasbCvmConfig `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbCvmInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbCvmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表

		ServiceList []*DasbInstance `json:"ServiceList,omitempty" name:"ServiceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitDasbInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例唯一标识串

	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
	// 配置内容

	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`
}

func (r *InitDasbInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitDasbInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstanceRequest struct {
	*tchttp.BaseRequest

	// 堡垒机基础配置信息

	BaseInfo *string `json:"BaseInfo,omitempty" name:"BaseInfo"`
	// CVM配置信息

	CvmInfo *string `json:"CvmInfo,omitempty" name:"CvmInfo"`
	// 数据盘配置信息

	CbsInfo *string `json:"CbsInfo,omitempty" name:"CbsInfo"`
}

func (r *RunInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DasbCvmConfig struct {

	// 主机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主机核数

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 外网带宽

	NetBand *uint64 `json:"NetBand,omitempty" name:"NetBand"`
	// 系统盘大小

	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
	// 数据盘大小

	DataDiskSize *uint64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`
	// 购买月份

	MonthSpan *uint64 `json:"MonthSpan,omitempty" name:"MonthSpan"`
	// 所属商品模型ID

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名

	Region *string `json:"Region,omitempty" name:"Region"`
	// 使用的镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type DasbCvmInstanceType struct {

	// 主机实例类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 主机实例名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 系统盘类型

	SystemDiskType *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`
	// 数据盘类型

	DataDiskType *string `json:"DataDiskType,omitempty" name:"DataDiskType"`
	// 所属区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type SendMessageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMessageRequest struct {
	*tchttp.BaseRequest

	// 手机号码

	Telphone *string `json:"Telphone,omitempty" name:"Telphone"`
	// 验证码

	SmsCode *string `json:"SmsCode,omitempty" name:"SmsCode"`
}

func (r *SendMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDasbSecretInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDasbSecretInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDasbSecretInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbInstanceAlarmRequest struct {
	*tchttp.BaseRequest

	// 实例唯一Key

	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
	// OwnerUin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *DescribeDasbInstanceAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbInstanceAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DasbInstanceLog struct {

	// 日志ID

	LogId *uint64 `json:"LogId,omitempty" name:"LogId"`
	// 日志内容

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
	// 日志记录时间

	LogTime *uint64 `json:"LogTime,omitempty" name:"LogTime"`
}

type DescribeDasbImageIdsRequest struct {
	*tchttp.BaseRequest

	// 镜像架构

	ImageArch *string `json:"ImageArch,omitempty" name:"ImageArch"`
	// 地域id

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *DescribeDasbImageIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbImageIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDasbSecretInfosRequest struct {
	*tchttp.BaseRequest

	// 密钥ID

	UserSecretId *string `json:"UserSecretId,omitempty" name:"UserSecretId"`
	// 密钥KEY

	UserSecretKey *string `json:"UserSecretKey,omitempty" name:"UserSecretKey"`
}

func (r *SetDasbSecretInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDasbSecretInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type PwdLog struct {

	// 用户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 订单关联标识

	IncKey *string `json:"IncKey,omitempty" name:"IncKey"`
}

type DescribeDasbCvmConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Dasb主机配置信息

		ConfigSet []*DasbCvmConfig `json:"ConfigSet,omitempty" name:"ConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbCvmConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbCvmConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbSystemVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDasbSystemVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbSystemVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageDesc struct {

	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像架构

	ImageArch *string `json:"ImageArch,omitempty" name:"ImageArch"`
}

type DescribeDasbImageIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域镜像信息

		RegionImageInfos []*RegionImageItem `json:"RegionImageInfos,omitempty" name:"RegionImageInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbImageIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbImageIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDasbUserInfosRequest struct {
	*tchttp.BaseRequest
}

func (r *GetDasbUserInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDasbUserInfosRequest) FromJsonString(s string) error {
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

type DescribeDasbCvmConfigRequest struct {
	*tchttp.BaseRequest

	// 商品定价ID

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

func (r *DescribeDasbCvmConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbCvmConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbCvmInstanceRequest struct {
	*tchttp.BaseRequest

	// DASB产品实例唯一KEY

	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
}

func (r *DescribeDasbCvmInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbCvmInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbSystemConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置

		Config []*string `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbSystemConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbSystemConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbSystemConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDasbSystemConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbSystemConfigRequest) FromJsonString(s string) error {
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

type AsyncInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AsyncInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AsyncInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbInstanceAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志列表

		LogSet []*DasbInstanceLog `json:"LogSet,omitempty" name:"LogSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbInstanceAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbInstanceAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDasbUserInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否已设置Secret，0：未设置；1：已设置

		IsSetSecret *int64 `json:"IsSetSecret,omitempty" name:"IsSetSecret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDasbUserInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDasbUserInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitDasbInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDasbInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitDasbInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsyncInstanceRequest struct {
	*tchttp.BaseRequest

	// 地域

	AsyncRegion *string `json:"AsyncRegion,omitempty" name:"AsyncRegion"`
	// uin

	AsyncUin *string `json:"AsyncUin,omitempty" name:"AsyncUin"`
}

func (r *AsyncInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AsyncInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDasbSystemVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDasbSystemVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbSystemVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNameInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InsKey *string `json:"InsKey,omitempty" name:"InsKey"`
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

type DasbInstance struct {

	// 实例key

	PincKey *string `json:"PincKey,omitempty" name:"PincKey"`
	// 实例名称

	PsysName *string `json:"PsysName,omitempty" name:"PsysName"`
	// 公网ip

	MpublicIp *string `json:"MpublicIp,omitempty" name:"MpublicIp"`
	// 产品型号

	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`
	// 购买时间

	PaddTime *uint64 `json:"PaddTime,omitempty" name:"PaddTime"`
	// 内网IP

	MprivateIp *string `json:"MprivateIp,omitempty" name:"MprivateIp"`
	// 实例状态

	PsysStatus *uint64 `json:"PsysStatus,omitempty" name:"PsysStatus"`
	// 实例开始时间

	PstartTime *uint64 `json:"PstartTime,omitempty" name:"PstartTime"`
	// 实例结束时间

	PendTime *uint64 `json:"PendTime,omitempty" name:"PendTime"`
	// 实例所在地域ID

	Pregion *uint64 `json:"Pregion,omitempty" name:"Pregion"`
	// 实例所在地区ID

	PzoneId *uint64 `json:"PzoneId,omitempty" name:"PzoneId"`
	// CVM实例名列表

	CvmInstanceNames []*string `json:"CvmInstanceNames,omitempty" name:"CvmInstanceNames"`
	// CVM实例ID列表

	CvmInstanceIds []*string `json:"CvmInstanceIds,omitempty" name:"CvmInstanceIds"`
	// 实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 购买者UIN

	Puin *string `json:"Puin,omitempty" name:"Puin"`
	// 系统管理员初始密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type DescribeDasbServiceRequest struct {
	*tchttp.BaseRequest

	// 分页数量，默认值为1000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDasbServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDasbServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
