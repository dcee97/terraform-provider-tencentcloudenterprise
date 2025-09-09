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

package v20220627

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type QueryProductPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySubProductConsoleRequest struct {
	*tchttp.BaseRequest
}

func (r *QuerySubProductConsoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySubProductConsoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductConfigRequest struct {
	*tchttp.BaseRequest

	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品编码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项标签

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项标签

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 配置项keys

	ConfigKeys []*string `json:"ConfigKeys,omitempty" name:"ConfigKeys"`
	// 配置项values

	ConfigValues []*string `json:"ConfigValues,omitempty" name:"ConfigValues"`
	// 查询包含的配置项value

	ConfigValueIn *string `json:"ConfigValueIn,omitempty" name:"ConfigValueIn"`
}

func (r *QueryProductConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomDiscountRequest struct {
	*tchttp.BaseRequest

	// 折扣类型，user-用户折扣, common-官网折扣, activity-活动折扣

	DiscountType *string `json:"DiscountType,omitempty" name:"DiscountType"`
	// 专区账号uin

	TceUin *string `json:"TceUin,omitempty" name:"TceUin"`
	//  优惠状态，0初始化，1生效，2废弃，3未开始，4已结束

	Status *string `json:"Status,omitempty" name:"Status"`
	// 一层产品code

	ProductDisplayCode *string `json:"ProductDisplayCode,omitempty" name:"ProductDisplayCode"`
	// 二层产品code

	SubProductDisplayCode *string `json:"SubProductDisplayCode,omitempty" name:"SubProductDisplayCode"`
	// 有效期时间范围

	ValidityTime *ValidityTime `json:"ValidityTime,omitempty" name:"ValidityTime"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 页数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryCustomDiscountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomDiscountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品配置

		Data []*ProductConfig `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySubProductConsoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询一二层产品结果

		Data []*ProductInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySubProductConsoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySubProductConsoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAccountDescRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryAccountDescRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAccountDescRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductConfig struct {

	// 自增id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品编码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项编码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项编码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 配置项key

	ConfigKey *string `json:"ConfigKey,omitempty" name:"ConfigKey"`
	// 配置项value

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
	// 创建人

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 修改人

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// 修改时间

	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
}

type QueryCustomDiscountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 折扣信息列表

		DiscountSet []*Discount `json:"DiscountSet,omitempty" name:"DiscountSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCustomDiscountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySubBillingItemsRequest struct {
	*tchttp.BaseRequest

	// 二层子产品code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
}

func (r *QuerySubBillingItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySubBillingItemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubProductInfo struct {

	// 子产品Code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 子产品展示标识

	SubProductDisplayCode *string `json:"SubProductDisplayCode,omitempty" name:"SubProductDisplayCode"`
	// 子产品展示名称

	SubProductDisplayName *string `json:"SubProductDisplayName,omitempty" name:"SubProductDisplayName"`
	// 二层可询价， 0不可询价 1可询价

	SubProductShow *uint64 `json:"SubProductShow,omitempty" name:"SubProductShow"`
}

type QueryBalanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 共有云侧uin

		CloudUin *string `json:"CloudUin,omitempty" name:"CloudUin"`
		// 可用余额

		Balance *string `json:"Balance,omitempty" name:"Balance"`
		// 更新时间

		UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBalanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryInstructionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品配置

		Data *ProductConfig `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryInstructionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryInstructionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBalanceRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryBalanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBalanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Discount struct {

	// 活动id，折扣类型为活动时有效

	ActivityID *string `json:"ActivityID,omitempty" name:"ActivityID"`
	// 折扣生效时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 计费项

	BillingItemDisplayCode *string `json:"BillingItemDisplayCode,omitempty" name:"BillingItemDisplayCode"`
	// 计费项展示名称

	BillingItemDisplayName *string `json:"BillingItemDisplayName,omitempty" name:"BillingItemDisplayName"`
	// 共有云侧账号名称

	CloudName *string `json:"CloudName,omitempty" name:"CloudName"`
	// 共有云侧账号uin

	CloudUin *string `json:"CloudUin,omitempty" name:"CloudUin"`
	// 折扣条件列表

	ConditionList []*DiscountCondition `json:"ConditionList,omitempty" name:"ConditionList"`
	// 同步公有云折扣操作者，前端可忽略

	CreateBy *string `json:"CreateBy,omitempty" name:"CreateBy"`
	// 折扣创建者

	DiscountCreateUser *string `json:"DiscountCreateUser,omitempty" name:"DiscountCreateUser"`
	// 优惠id

	DiscountID *int64 `json:"DiscountID,omitempty" name:"DiscountID"`
	// 折扣来源0无来源1产品手配置2审批流3管理端

	DiscountOrigin *int64 `json:"DiscountOrigin,omitempty" name:"DiscountOrigin"`
	// 折扣类型，user-用户折扣, common-官网折扣, activity-活动折扣

	DiscountType *string `json:"DiscountType,omitempty" name:"DiscountType"`
	// 折扣优先级，按最新时间排序

	DiscountUpdateTime *string `json:"DiscountUpdateTime,omitempty" name:"DiscountUpdateTime"`
	// 折扣用户，等同于CloudUin

	DiscountUser *string `json:"DiscountUser,omitempty" name:"DiscountUser"`
	// 折扣值

	DiscountValue *string `json:"DiscountValue,omitempty" name:"DiscountValue"`
	// 折扣失效时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 主键id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 是否自动续期 0-否，1-是

	IsAutoRenew *int64 `json:"IsAutoRenew,omitempty" name:"IsAutoRenew"`
	// 是否和官网折扣冲突，当type为user时有效，0-不冲突，1-冲突，冲突时代表享受官网折扣

	IsConflict *int64 `json:"IsConflict,omitempty" name:"IsConflict"`
	// 是否为高优折扣,0-否，1-是

	IsHighPriority *int64 `json:"IsHighPriority,omitempty" name:"IsHighPriority"`
	// 是否继承,0-否, 1-是

	IsInherit *int64 `json:"IsInherit,omitempty" name:"IsInherit"`
	// 是否内部用户折扣，0否，1是

	IsInternal *int64 `json:"IsInternal,omitempty" name:"IsInternal"`
	// 付费模式，0-后付费，1-预付费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 优惠类型,0-折扣值, 1-指定价

	PreferentialType *int64 `json:"PreferentialType,omitempty" name:"PreferentialType"`
	// 指定价详情

	SpecifyPriceDetail *specifyPrice `json:"SpecifyPriceDetail,omitempty" name:"SpecifyPriceDetail"`
	// 一层产品code

	ProductDisplayCode *string `json:"ProductDisplayCode,omitempty" name:"ProductDisplayCode"`
	// 一层产品展示名称

	ProductCodeDisplayName *string `json:"ProductCodeDisplayName,omitempty" name:"ProductCodeDisplayName"`
	// 折扣备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	//  优惠状态，0初始化，1生效，2废弃，3未开始，4已结束

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 计费细项code

	SubBillingItemDisplayCode *string `json:"SubBillingItemDisplayCode,omitempty" name:"SubBillingItemDisplayCode"`
	// 计费细项展示名称，可以为空

	SubBillingItemDisplayName *string `json:"SubBillingItemDisplayName,omitempty" name:"SubBillingItemDisplayName"`
	// 二层产品code

	SubProductDisplayCode *string `json:"SubProductDisplayCode,omitempty" name:"SubProductDisplayCode"`
	// 二层产品展示名称

	SubProductDisplayName *string `json:"SubProductDisplayName,omitempty" name:"SubProductDisplayName"`
	// 同步折扣时间

	SyncTime *string `json:"SyncTime,omitempty" name:"SyncTime"`
	// 专区账号名

	TceName *string `json:"TceName,omitempty" name:"TceName"`
	// 专区账号uin

	TceUin *string `json:"TceUin,omitempty" name:"TceUin"`
	// 折扣更新操作者

	UpdateBy *string `json:"UpdateBy,omitempty" name:"UpdateBy"`
}

type PriceDetail struct {

	// 区间的高位，阶梯价：开区间，累进阶梯：闭区间,线性、指定不需要

	High *string `json:"High,omitempty" name:"High"`
	// 区间的低位，阶梯价：闭区间，累进阶梯：开区间,线性、指定不需要

	Low *string `json:"Low,omitempty" name:"Low"`
	// 指定价格需要,指定每几个基本单位单价多少, 其他类型都传:1

	Price *string `json:"Price,omitempty" name:"Price"`
	// 指定价规格

	Cnt *string `json:"Cnt,omitempty" name:"Cnt"`
	// 时间单位, 1:年 2:月 3:日 4:小时 5:分钟 6:秒 7:打包价与时间无关

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 平台, 云平: 1、开平: 0

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 币种:CNY:人民币 USD:美元

	Currency *string `json:"Currency,omitempty" name:"Currency"`
	// 指定价类型，1:指定价格，2:线性价格，3:阶梯价格，5:累进阶梯 6. 区间指定价

	PriceType *string `json:"PriceType,omitempty" name:"PriceType"`
	// 付费模式, 	0: 后付费 1:预付费

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	//  商品计量单位描述(核，G...)只用于展示

	ProductUnitDes *string `json:"ProductUnitDes,omitempty" name:"ProductUnitDes"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 区间指定价需要,步长

	Size *string `json:"Size,omitempty" name:"Size"`
	// 区间指定价需要,是否添加步长策略

	Add *string `json:"Add,omitempty" name:"Add"`
	// 区间指定价需要,每个步长的价格

	StepPrice *string `json:"StepPrice,omitempty" name:"StepPrice"`
}

type QueryAccountDescResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAccountDescResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAccountDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryInstructionRequest struct {
	*tchttp.BaseRequest

	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品编码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
}

func (r *QueryInstructionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryInstructionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySubBillingItemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySubBillingItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySubBillingItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiscountCondition struct {

	// 条件创建时间

	FCreateTime *string `json:"FCreateTime,omitempty" name:"FCreateTime"`
	// 记录唯一id

	FId *int64 `json:"FId,omitempty" name:"FId"`
	// 优惠条件Key

	FKey *string `json:"FKey,omitempty" name:"FKey"`
	// 优惠条件名称

	FKeyName *string `json:"FKeyName,omitempty" name:"FKeyName"`
	// 条件字段类型(string字符串,integer整数,json四层黑名单)，若传FValue，FKeyType必传

	FKeyType *string `json:"FKeyType,omitempty" name:"FKeyType"`
	// 条件操作符

	FOper *string `json:"FOper,omitempty" name:"FOper"`
	// 条件备注

	FRemark *string `json:"FRemark,omitempty" name:"FRemark"`
	// 状态（0：失效，1：生效）

	FStatus *int64 `json:"FStatus,omitempty" name:"FStatus"`
	// 更新时间

	FUpdateTime *string `json:"FUpdateTime,omitempty" name:"FUpdateTime"`
	// 条件value，如果FValue为数字，FKeyType必须传integer|如果是组合优惠，多个value用###&分隔，FKeyType为string|

	FValue *string `json:"FValue,omitempty" name:"FValue"`
	// 条件value名称

	FValueName *string `json:"FValueName,omitempty" name:"FValueName"`
}

type ProductInfo struct {

	// 产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品展示标识

	ProductDisplayCode *string `json:"ProductDisplayCode,omitempty" name:"ProductDisplayCode"`
	// 产品展示名称

	ProductDisplayName *string `json:"ProductDisplayName,omitempty" name:"ProductDisplayName"`
	// 一层可询价，0不可询价 1可询价

	ProductShow *uint64 `json:"ProductShow,omitempty" name:"ProductShow"`
	// 二层信息

	SubProducts []*SubProductInfo `json:"SubProducts,omitempty" name:"SubProducts"`
}

type QueryProductPriceRequest struct {
	*tchttp.BaseRequest

	// 一层，产品code

	ProductDisplayCode *string `json:"ProductDisplayCode,omitempty" name:"ProductDisplayCode"`
	// 三层，计费项code

	BillingItemDisplayCodes []*string `json:"BillingItemDisplayCodes,omitempty" name:"BillingItemDisplayCodes"`
	// 四层，计费细项code

	SubBillingItemDisplayCodes []*string `json:"SubBillingItemDisplayCodes,omitempty" name:"SubBillingItemDisplayCodes"`
	// 分页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 分页页码

	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`
	// 顺序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 二层，子产品code

	SubProductDisplayCode *string `json:"SubProductDisplayCode,omitempty" name:"SubProductDisplayCode"`
}

func (r *QueryProductPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiscountPriceType struct {

	// // 指定价类型，1:指定价格，2:线性价格，3:阶梯价格，5:累进阶梯 6. 区间指定价

	PriceType *string `json:"PriceType,omitempty" name:"PriceType"`
	// 折扣创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
}

type specifyPrice struct {

	// 折扣指定价类型信息

	Product *DiscountPriceType `json:"Product,omitempty" name:"Product"`
	// 折扣指定价详情列表

	Prices []*PriceDetail `json:"Prices,omitempty" name:"Prices"`
}

type ValidityTime struct {

	// 生效时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 失效时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}
