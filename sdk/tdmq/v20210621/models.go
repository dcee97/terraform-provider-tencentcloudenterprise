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

package v20210621

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type AMQPRouteRelation struct {

	// 路由关系ID

	RouteRelationId *string `json:"RouteRelationId,omitempty" name:"RouteRelationId"`
	// 源Exchange

	SourceExchange *string `json:"SourceExchange,omitempty" name:"SourceExchange"`
	// 目标类型:Queue|Exchange

	DestType *string `json:"DestType,omitempty" name:"DestType"`
	// 目标值

	DestValue *string `json:"DestValue,omitempty" name:"DestValue"`
	// 绑定key

	RoutingKey *string `json:"RoutingKey,omitempty" name:"RoutingKey"`
	// 源路由类型:Direct|Topic|Fanout

	SourceExchangeType *string `json:"SourceExchangeType,omitempty" name:"SourceExchangeType"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 说明信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type AMQPClusterInfo struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 地域信息

	Region *string `json:"Region,omitempty" name:"Region"`
	// 创建时间，毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 集群说明信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 公网接入地址

	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`
	// VPC接入地址

	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`
}

type AMQPClusterRecentStats struct {

	// Queue数量

	QueueNum *uint64 `json:"QueueNum,omitempty" name:"QueueNum"`
	// 消息生产数

	ProducedMsgNum *uint64 `json:"ProducedMsgNum,omitempty" name:"ProducedMsgNum"`
	// 消息堆积数

	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitempty" name:"AccumulativeMsgNum"`
	// Exchange数量

	ExchangeNum *uint64 `json:"ExchangeNum,omitempty" name:"ExchangeNum"`
}

type AMQPClusterDetail struct {

	// 集群基本信息

	Info *AMQPClusterInfo `json:"Info,omitempty" name:"Info"`
	// 集群配置信息

	Config *AMQPClusterConfig `json:"Config,omitempty" name:"Config"`
}

type ConfigBaseInfoOpt struct {

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 程序路径

	RootPath *string `json:"RootPath,omitempty" name:"RootPath"`
	// 数据路径

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// 配置路径

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 配置文件名

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
}

type GetPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceRuleOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 备注信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 规则当前状态，0表示下线，1表示可用

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 协议名

	ProtocolName *string `json:"ProtocolName,omitempty" name:"ProtocolName"`
}

func (r *UpdateResourceRuleOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResourceRuleOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPQueueConsumer struct {

	// 消费者实例的地址和端口

	ClientAddr *string `json:"ClientAddr,omitempty" name:"ClientAddr"`
	// 消费者标签

	ClientTag *string `json:"ClientTag,omitempty" name:"ClientTag"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type AMQPVHostConnection struct {

	// 客户端实例的地址和端口

	ClientAddr *string `json:"ClientAddr,omitempty" name:"ClientAddr"`
	// 客户端协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 是否支持SSL/TLS

	Security *bool `json:"Security,omitempty" name:"Security"`
	// 信道数

	ChannelNum *uint64 `json:"ChannelNum,omitempty" name:"ChannelNum"`
	// 接受速率

	ReceiveRate *uint64 `json:"ReceiveRate,omitempty" name:"ReceiveRate"`
	// 发送速率

	SendRate *uint64 `json:"SendRate,omitempty" name:"SendRate"`
}

type AMQPClusterConfig struct {

	// 单Vhost TPS上限

	MaxTpsPerVHost *uint64 `json:"MaxTpsPerVHost,omitempty" name:"MaxTpsPerVHost"`
	// 单Vhost客户端连接数上限

	MaxConnNumPerVHost *uint64 `json:"MaxConnNumPerVHost,omitempty" name:"MaxConnNumPerVHost"`
	// 最大Vhost数量

	MaxVHostNum *uint64 `json:"MaxVHostNum,omitempty" name:"MaxVHostNum"`
	// 最大exchange数量

	MaxExchangeNum *uint64 `json:"MaxExchangeNum,omitempty" name:"MaxExchangeNum"`
	// 最大Queue数量

	MaxQueueNum *uint64 `json:"MaxQueueNum,omitempty" name:"MaxQueueNum"`
	// 消息最大保留时间，以毫秒为单位

	MaxRetentionTime *uint64 `json:"MaxRetentionTime,omitempty" name:"MaxRetentionTime"`
	// 已使用Vhost数量

	UsedVHostNum *uint64 `json:"UsedVHostNum,omitempty" name:"UsedVHostNum"`
	// 已使用exchange数量

	UsedExchangeNum *uint64 `json:"UsedExchangeNum,omitempty" name:"UsedExchangeNum"`
	// 已使用queue数量

	UsedQueueNum *uint64 `json:"UsedQueueNum,omitempty" name:"UsedQueueNum"`
}

type ConfigTypeInfoOpt struct {

	// 模块组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 模块配置文件名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 配置模块名

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type AMQPExchange struct {

	// Exchange名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Exchange的类别，为枚举类型:Direct, Fanout, Topic

	Type *string `json:"Type,omitempty" name:"Type"`
	// 主绑定数

	SourceBindedNum *uint64 `json:"SourceBindedNum,omitempty" name:"SourceBindedNum"`
	// 说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 被绑定数

	DestBindedNum *uint64 `json:"DestBindedNum,omitempty" name:"DestBindedNum"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否为内部Exchange(以amq.前缀开头的)

	Internal *bool `json:"Internal,omitempty" name:"Internal"`
}

type UpdateResourceRuleOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateResourceRuleOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResourceRuleOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPVHost struct {

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 未消费消息的保留时间，以毫秒单位，范围60秒到15天

	MsgTtl *uint64 `json:"MsgTtl,omitempty" name:"MsgTtl"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type DeliveryInfoOpt struct {

	// 代理节点IP

	DeliveryIp *string `json:"DeliveryIp,omitempty" name:"DeliveryIp"`
	// 下发时间

	DeliveryTimes *int64 `json:"DeliveryTimes,omitempty" name:"DeliveryTimes"`
	// 下发状态（实际大小1字节） 0 下发中 1 陈工 2失败

	DeliveryStatus *int64 `json:"DeliveryStatus,omitempty" name:"DeliveryStatus"`
	// 版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 新增时间

	AddTime *int64 `json:"AddTime,omitempty" name:"AddTime"`
}

type InternalNamespace struct {

	// 命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 客户UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 客户的APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 虚拟集群ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 虚拟集群名称

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// 创建时间，毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建时间，毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type GetPriceRequest struct {
	*tchttp.BaseRequest

	// 产品数

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 产品类型（sv_tdmq_pulsar_partition/sv_tdmq_pulsar_message_count/sv_tdmq_pulsar_storage）

	GoodType *string `json:"GoodType,omitempty" name:"GoodType"`
}

func (r *GetPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendConfigInfoOpt struct {

	// 节点ip地址

	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`
	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 下发状态

	Status *bool `json:"Status,omitempty" name:"Status"`
}

type AMQPQueueDetail struct {

	// Queue名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 被绑定数

	DestBindedNum *uint64 `json:"DestBindedNum,omitempty" name:"DestBindedNum"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 在线消费者数

	OnlineConsumerNum *uint64 `json:"OnlineConsumerNum,omitempty" name:"OnlineConsumerNum"`
	// 每秒钟的事务数

	Tps *uint64 `json:"Tps,omitempty" name:"Tps"`
	// 消息堆积数

	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitempty" name:"AccumulativeMsgNum"`
	// 是否自动删除

	AutoDelete *bool `json:"AutoDelete,omitempty" name:"AutoDelete"`
	// 死信交换机

	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" name:"DeadLetterExchange"`
	// 死信交换机路由键

	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" name:"DeadLetterRoutingKey"`
}

type ConfigTemplateItemOpt struct {

	// 配置属性

	ItemKey *string `json:"ItemKey,omitempty" name:"ItemKey"`
	// 属性值

	ItemValue *string `json:"ItemValue,omitempty" name:"ItemValue"`
	// 配置项备注

	ItemDesc *string `json:"ItemDesc,omitempty" name:"ItemDesc"`
	// 是否动态配置

	IsDynamic *bool `json:"IsDynamic,omitempty" name:"IsDynamic"`
}
