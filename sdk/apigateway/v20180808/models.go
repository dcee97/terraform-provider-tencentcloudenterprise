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

package v20180808

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type CreateFreePkgResourceRequest struct {
	*tchttp.BaseRequest

	// 操作类型，传值： CREATE

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 付费类型，传值： PREPAID

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *CreateFreePkgResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFreePkgResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAppBindApisStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用绑定的Api列表。

		Result *ApiAppApiInfos `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiAppBindApisStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAppBindApisStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPStrategysStatus struct {

	// 策略数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 策略列表。

	StrategySet []*IPStrategy `json:"StrategySet,omitempty" name:"StrategySet"`
}

type DescribeApiKeyRequest struct {
	*tchttp.BaseRequest

	// API 密钥 ID。

	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DescribeApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceEnvironmentSet struct {

	// 服务绑定环境总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 服务绑定环境列表。

	EnvironmentList []*Environment `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainMappingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义路径映射列表。

		Result *ServiceSubDomainMappings `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceSubDomainMappingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceSubDomainMappingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindSubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待解绑的自定义的域名。

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *UnBindSubDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindSubDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayResult struct {

	// 云原生API网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生网关状态。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeleteServiceSubDomainMappingRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务绑定的自定义域名。

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 待删除映射的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DeleteServiceSubDomainMappingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceSubDomainMappingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSearchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取更多检索结果的游标，值为""表示无后续结果

		ConText *string `json:"ConText,omitempty" name:"ConText"`
		// 由0或多条日志组成，每条日志格式如下：
		// '[$app_id][$env_name][$service_id][$http_host][$api_id][$uri][$scheme][rsp_st:$status][ups_st:$upstream_status]'
		// '[cip:$remote_addr][uip:$upstream_addr][vip:$server_addr][rsp_len:$bytes_sent][req_len:$request_length]'
		// '[req_t:$request_time][ups_rsp_t:$upstream_response_time][ups_conn_t:$upstream_connect_time][ups_head_t:$upstream_header_time]’
		// '[err_msg:$err_msg][tcp_rtt:$tcpinfo_rtt][$pid][$time_local][req_id:$request_id]';
		//
		// 说明：
		// app_id： 用户 ID。
		// env_name：环境名称。
		// service_id： 服务 ID。
		// http_host： 域名。
		// api_id： API 的 ID。
		// uri：请求的路径。
		// scheme： HTTP/HTTPS 协议。
		// rsp_st： 请求响应状态码。
		// ups_st： 后端业务服务器的响应状态码（如果请求透传到后端，改变量不为空。如果请求在 APIGW 就被拦截了，那么该变量显示为 -）。
		// cip： 客户端 IP。
		// uip： 后端业务服务（upstream）的 IP。
		// vip： 请求访问的 VIP。
		// rsp_len： 响应长度。
		// req_len： 请求长度。
		// req_t： 请求响应的总时间。
		// ups_rsp_t： 后端响应的总时间（apigw 建立连接到接收到后端响应的时间）。
		// ups_conn_t: 与后端业务服务器连接建立成功时间。
		// ups_head_t：后端响应的头部到达时间。
		// err_msg： 错误信息。
		// tcp_rtt： 客户端 TCP 连接信息，RTT（Round Trip Time）由三部分组成：链路的传播时间（propagation delay）、末端系统的处理时间、路由器缓存中的排队和处理时间（queuing delay）。
		// req_id：请求id。

		LogSet []*string `json:"LogSet,omitempty" name:"LogSet"`
		// 单次搜索返回的日志条数，TotalCount <= Limit

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogSearchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApisAuthRelationApiRequest struct {
	*tchttp.BaseRequest

	// 服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// api列表

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
	// 关联授权api

	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`
}

func (r *ModifyApisAuthRelationApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApisAuthRelationApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUpstreamResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建返回的唯一id

		UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUpstreamResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePackStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *ResourcePackResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcePackStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcePackStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 待查询的使用计划唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DescribeUsagePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsagePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceReleaseVersion struct {

	// 发布版本总数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 发布版本列表。

	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList"`
}

type CreateApiAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增的应用详情。

		Result *ApiAppInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyApisStatusRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 策略唯一ID。

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略所在环境。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。支持 ApiPath、ApiName、KeyWord（模糊查询Path 和Name）。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPStrategyApisStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPStrategyApisStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePluginsByApiRequest struct {
	*tchttp.BaseRequest

	// 要查询的API ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 环境信息。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。预留字段，目前不支持过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序依据，默认按照插件创建时间排序。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序方式，默认倒序排序。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 要查询的服务ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DescribePluginsByApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginsByApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待修改路径映射的自定义的域名。

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 证书 ID，协议包含 HTTPS 的时候要传该字段。

	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
	// 是否修改为使用默认路径映射。为 true，表示使用默认路径映射，为 false，表示使用自定义路径映射。

	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`
	// 修改后的自定义域名协议类型。（http，https 或 http&https)

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 修改后的路径映射列表。

	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`
	// 网络类型 （'INNER' 或 'OUTER'）

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 是否将HTTP请求强制跳转 HTTPS，默认为false。参数为 true时，API网关会将所有使用该自定义域名的 HTTP 协议的请求重定向至 HTTPS 协议进行转发。

	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
}

func (r *ModifySubDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindTriggerRequest struct {
	*tchttp.BaseRequest

	// 六段式资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 触发器类型，默认为apigw。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 函数命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 函数名称

	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`
	// 函数的版本名称或别名

	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`
}

func (r *UnbindTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindTriggerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询服务自定义域名列表。

		Result *DomainSets `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceSubDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceSubDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTcbScfApisRequest struct {
	*tchttp.BaseRequest

	// 环境ID。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// SCF方法名称。

	ScfFunctionName *string `json:"ScfFunctionName,omitempty" name:"ScfFunctionName"`
	// SCF方法命名空间。

	ScfFunctionNamespace *string `json:"ScfFunctionNamespace,omitempty" name:"ScfFunctionNamespace"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud..com/document/api/213/15688)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud..com/document/api/213/15688)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTcbScfApisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcbScfApisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlansStatusRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 使用计划过滤条件。支持UsagePlanId、UsagePlanName、NotServiceId、NotApiId、Environment。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUsagePlansStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsagePlansStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApiKeyRequest struct {
	*tchttp.BaseRequest

	// 待更换的密钥 ID。

	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
	// 待更换的密钥 Key，更新自定义密钥时，该字段为必传。长度10 - 50字符，包括字母、数字、英文下划线。

	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *UpdateApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckCloneApisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回格式校验

		Result *CheckSet `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckCloneApisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCloneApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailablePluginApisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 插件可绑定的API列表信息。

		Result *ApiInfoSummary `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailablePluginApisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailablePluginApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务列表查询结果。

		Result *ServicesStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServicesStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServicesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunApiRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API 唯一 ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// API 的前端请求头部，是 json_dump 后的数据。

	RequestHeader *string `json:"RequestHeader,omitempty" name:"RequestHeader"`
	// API 的前端请求 Query，是 json_dump 后的数据。

	RequestQuery *string `json:"RequestQuery,omitempty" name:"RequestQuery"`
	// API 的请求 Path，是 json_dump 后的数据。

	RequestPath *string `json:"RequestPath,omitempty" name:"RequestPath"`
	// API 的请求方法。只支持 HEAD、GET、POST、PUT、PATCH 和 DELETE。

	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`
	// API 的请求 Body。

	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
	// 调试请求的内容类型。当前只支持 application/json 和 application/x-www-form-urlencoded，不传的话，默认为 application/x-www-form-urlencoded。

	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	// API 的请求 Body，当 API 有设置 Body 类型入参时，用数组格式传入。

	RequestBodyDict *string `json:"RequestBodyDict,omitempty" name:"RequestBodyDict"`
	// http直通的场景下，用户可自定义的path

	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *RunApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiEnvironmentStrategy struct {

	// API唯一ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 用户自定义API名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// API的路径。如/path。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API的方法。如GET。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 环境的限流信息。

	EnvironmentStrategySet []*EnvironmentStrategy `json:"EnvironmentStrategySet,omitempty" name:"EnvironmentStrategySet"`
}

type EnvironmentList struct {

	// 总数

	TotolCount *int64 `json:"TotolCount,omitempty" name:"TotolCount"`
	// 服务的环境列表

	EnvironmentSet []*EnvironmentUpload `json:"EnvironmentSet,omitempty" name:"EnvironmentSet"`
}

type ServiceConfig struct {

	// 后端类型。启用vpc时生效，目前支持的类型为clb和vpc通道

	Product *string `json:"Product,omitempty" name:"Product"`
	// vpc 的唯一ID。

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// API 的后端服务url。如果ServiceType是HTTP，则此参数必传。

	Url *string `json:"Url,omitempty" name:"Url"`
	// API 的后端服务路径，如 /path。如果 ServiceType 是 HTTP，则此参数必传。前后端路径可不同。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API的后端服务请求方法，如 GET。如果 ServiceType 是 HTTP，则此参数必传。前后端方法可不同。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 当绑定vpc通道才需要

	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
}

type CreateUsagePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 使用计划详情。

		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUsagePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceReleaseVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务发布版本列表。

		Result *ServiceReleaseVersion `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceReleaseVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceReleaseVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签的 key。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 便签的 value。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteLogRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeysStatusRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。支持AccessKeyId、AccessKeySecret、SecretName、NotUsagePlanId、Status、KeyWord（ 可以匹配name或者path）。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiKeysStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiKeysStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindApiAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解除绑定操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindApiAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReqParameter struct {

	// API 的前端参数名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// API 的前端参数位置，如 header。目前支持 header、query、path。

	Position *string `json:"Position,omitempty" name:"Position"`
	// API 的前端参数类型，如 String、int。

	Type *string `json:"Type,omitempty" name:"Type"`
	// API 的前端参数默认值。

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// API 的前端参数是否必填，True：表示必填，False：表示可选。

	Required *bool `json:"Required,omitempty" name:"Required"`
	// API 的前端参数备注。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type DeleteIPStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIPStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiBindApiAppsStatusRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。支持ApiAppId、Environment、KeyWord（ 可以匹配name或者ID）。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 服务ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// Api的ID的数组

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *DescribeApiBindApiAppsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiBindApiAppsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type APIDocs struct {

	// API文档数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// API文档基本信息

	APIDocSet []*APIDoc `json:"APIDocSet,omitempty" name:"APIDocSet"`
}

type DocumentSDK struct {

	// 生成的 document 会存放到 COS 中，此出参返回产生文件的下载链接。

	DocumentURL *string `json:"DocumentURL,omitempty" name:"DocumentURL"`
	// 生成的 SDK 会存放到 COS 中，此出参返回产生 SDK 文件的下载链接。

	SdkURL *string `json:"SdkURL,omitempty" name:"SdkURL"`
}

type TargetLoadBalanceConfReq struct {

	// 方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 是否会话保持

	SessionStickRequired *bool `json:"SessionStickRequired,omitempty" name:"SessionStickRequired"`
	// 会话保持超时时间

	SessionStickTimeout *int64 `json:"SessionStickTimeout,omitempty" name:"SessionStickTimeout"`
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务唯一ID。

		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
		// 用户自定义服务名称。

		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
		// 用户自定义服务描述。

		ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`
		// 外网默认域名。

		OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`
		// vpc内网默认域名。

		InnerSubDomain *string `json:"InnerSubDomain,omitempty" name:"InnerSubDomain"`
		// 服务创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
		// 网络类型列表，INNER为内网访问，OUTER为外网访问。

		NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
		// IP版本号。

		IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAPIDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAPIDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesStatusRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。支持ServiceId、ServiceName、NotUsagePlanId、Environment、IpVersion。InstanceId

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeServicesStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServicesStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportOpenApiRequest struct {
	*tchttp.BaseRequest

	// API所在的服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// Content格式，只能是YAML或者JSON，默认是YAML。

	EncodeType *string `json:"EncodeType,omitempty" name:"EncodeType"`
	// API 接口唯一 ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *ExportOpenApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportOpenApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayResult struct {

	// 云原生网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生网关状态。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type BindApiAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindApiAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BuildAPIDocRequest struct {
	*tchttp.BaseRequest

	// API文档ID

	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *BuildAPIDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BuildAPIDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAppRequest struct {
	*tchttp.BaseRequest

	// 应用ID。

	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

func (r *DescribeApiAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiBindApiAppsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用绑定的Api列表。

		Result *ApiAppApiInfos `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiBindApiAppsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiBindApiAppsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PluginSummary struct {

	// 插件个数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 插件详情。

	PluginSet []*Plugin `json:"PluginSet,omitempty" name:"PluginSet"`
}

type ServiceEnvironmentStrategyStatus struct {

	// 限流策略数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 限流策略列表。

	EnvironmentList []*ServiceEnvironmentStrategy `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
}

type CreateLogRuleRequest struct {
	*tchttp.BaseRequest

	// 日志规则名称。

	LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`
	// 服务唯一ID，表示所需上报日志的服务ID，不传代表所有服务。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 环境名称，表示所需要上报日志的环境名称，不传代表所有环境。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 日志集ID。

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 日志主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *CreateLogRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApiKeyRequest struct {
	*tchttp.BaseRequest

	// 待删除的密钥 ID。

	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DeleteApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DescribeServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 禁用密钥操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsagePlanEnvironment struct {

	// 绑定的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API 的唯一ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// API 的名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// API 的路径。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API 的方法。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 已经绑定的环境名称。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 已经使用的配额。

	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`
	// 最大请求量。

	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`
	// 每秒最大请求次数。

	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// 服务名称。

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type AttachPluginRequest struct {
	*tchttp.BaseRequest

	// 绑定的API网关插件ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
	// 要操作的服务ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 要操作API的环境。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 要绑定的API列表。

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *AttachPluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachPluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportOpenApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// export返回参数

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportOpenApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportOpenApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 待修改的策略所属服务的唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待修改的策略唯一ID。

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 待修改的策略详情。

	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *ModifyIPStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIPStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 使用计划唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 修改后的用户自定义的使用计划名称。

	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`
	// 修改后的用户自定义的使用计划描述。

	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`
	// 请求配额总数，取值范围为-1或者[1, 99999999]，默认为-1，表示不开启。

	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`
	// 每秒请求限制数，取值范围为-1或者[1, 2000]，默认-1，表示不开启。

	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *ModifyUsagePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUsagePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAPIDocRequest struct {
	*tchttp.BaseRequest

	// API文档名称

	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`
	// 服务名称

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 环境名称

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 生成文档的API列表

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *CreateAPIDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAPIDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePluginsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 插件详情。

		Result *PluginSummary `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePluginRequest struct {
	*tchttp.BaseRequest

	// 用户自定义的插件名称。最长50个字符，最短2个字符，支持 a-z,A-Z,0-9,_, 必须字母开头，字母或者数字结尾。

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 插件类型。目前支持IPControl, TrafficControl, Cors, CustomReq, CustomAuth。

	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`
	// 插件描述，限定200字以内。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 插件定义语句，支持json。

	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`
	// 参数属性，System：系统插件，Custom：自定义插件，默认为System。

	PluginAttribute *string `json:"PluginAttribute,omitempty" name:"PluginAttribute"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 项目id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreatePluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanSecretIdsRequest struct {
	*tchttp.BaseRequest

	// 绑定的使用计划唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanSecretIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsagePlanSecretIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 限流值。

	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
	// 环境名。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// API列表。

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *ModifyApiEnvironmentStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiEnvironmentStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindSubDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解绑自定义域名操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindSubDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindSubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceParameter struct {

	// API的后端服务参数名称。只有ServiceType是HTTP才会用到此参数。前后端参数名称可不同。

	Name *string `json:"Name,omitempty" name:"Name"`
	// API 的后端服务参数位置，如 head。只有 ServiceType 是 HTTP 才会用到此参数。前后端参数位置可配置不同。

	Position *string `json:"Position,omitempty" name:"Position"`
	// API 的后端服务参数对应的前端参数位置，如 head。只有 ServiceType 是 HTTP 才会用到此参数。

	RelevantRequestParameterPosition *string `json:"RelevantRequestParameterPosition,omitempty" name:"RelevantRequestParameterPosition"`
	// API 的后端服务参数对应的前端参数名称。只有 ServiceType 是 HTTP 才会用到此参数。

	RelevantRequestParameterName *string `json:"RelevantRequestParameterName,omitempty" name:"RelevantRequestParameterName"`
	// API 的后端服务参数默认值。只有 ServiceType 是 HTTP 才会用到此参数。

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// API 的后端服务参数备注。只有 ServiceType 是 HTTP 才会用到此参数。

	RelevantRequestParameterDesc *string `json:"RelevantRequestParameterDesc,omitempty" name:"RelevantRequestParameterDesc"`
	// API 的后端服务参数类型。只有 ServiceType 是 HTTP 才会用到此参数。

	RelevantRequestParameterType *string `json:"RelevantRequestParameterType,omitempty" name:"RelevantRequestParameterType"`
}

type CheckServiceConfigRequest struct {
	*tchttp.BaseRequest

	// 服务id。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务类型，如：HTTP、WEBSOCKET。若检测后端地址，ServiceType、BackendUrl必传，前端和后端不能同时检测。

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 后端服务地址，若检测后端地址，ServiceType、BackendUrl必传，前端和后端不能同时检测。

	BackendUrl *string `json:"BackendUrl,omitempty" name:"BackendUrl"`
	// 前端路径配置，若检测前端路径，RequestConfig与ServiceType、BackendUrl不可同时存在。

	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`
}

func (r *CheckServiceConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckServiceConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiInfo struct {

	// API 所在的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API 所在的服务的名称。

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// API 所在的服务的描述。

	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`
	// API 接口唯一 ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// API 接口的描述。

	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`
	// 创建时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 最后修改时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// API 接口的名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// API 类型。可取值为NORMAL（普通API）、TSF（微服务API）。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
	// API 的前端请求类型，如 HTTP 或 HTTPS 或者 HTTP 和 HTTPS。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// API 鉴权类型。可取值为 SECRET（密钥对鉴权）、NONE（免鉴权）、OAUTH。

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// OAUTH API的类型。可取值为NORMAL（业务API）、OAUTH（授权API）。

	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`
	// OAUTH 业务API 关联的授权API 唯一 ID。

	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`
	// OAUTH配置。

	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`
	// 是否购买后调试（云市场预留参数）。

	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`
	// 请求的前端配置。

	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`
	// 返回类型。

	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`
	// 自定义响应配置成功响应示例。

	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`
	// 自定义响应配置失败响应示例。

	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`
	// 用户自定义错误码配置。

	ResponseErrorCodes []*ErrorCodes `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`
	// 前端请求参数。

	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`
	// API 的后端服务超时时间，单位是秒。

	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`
	// API 的后端服务类型。可取值为 HTTP、MOCK、TSF、CLB、SCF、WEBSOCKET、TARGET（内测）。

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// API 的后端服务配置。

	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`
	// API的后端服务参数。

	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`
	// 常量参数。

	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`
	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。

	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`
	// scf 函数名称。当后端类型是SCF时生效。

	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`
	// scf 函数命名空间。当后端类型是SCF时生效。

	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`
	// scf函数版本。当后端类型是SCF时生效。

	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`
	// 是否开启集成响应。

	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`
	// scf websocket注册函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效

	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`
	// scf websocket注册函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`
	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`
	// scf websocket清理函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`
	// scf websocket清理函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`
	// scf websocket清理函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`
	// WEBSOCKET 回推地址。

	InternalDomain *string `json:"InternalDomain,omitempty" name:"InternalDomain"`
	// scf websocket传输函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`
	// scf websocket传输函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`
	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`
	// API绑定微服务服务列表。

	MicroServices []*MicroService `json:"MicroServices,omitempty" name:"MicroServices"`
	// 微服务信息详情。

	MicroServicesInfo []*int64 `json:"MicroServicesInfo,omitempty" name:"MicroServicesInfo"`
	// 微服务的负载均衡配置。

	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`
	// 微服务的健康检查配置。

	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`
	// 是否开启跨域。

	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`
	// API绑定的tag信息。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// API已发布的环境信息。

	Environments []*string `json:"Environments,omitempty" name:"Environments"`
	// 是否开启Base64编码，只有后端为scf时才会生效。

	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`
	// 是否开启Base64编码的header触发，只有后端为scf时才会生效。

	IsBase64Trigger *bool `json:"IsBase64Trigger,omitempty" name:"IsBase64Trigger"`
	// Header触发规则，总规则数量不超过10。

	Base64EncodedTriggerRules []*Base64EncodedTriggerRule `json:"Base64EncodedTriggerRules,omitempty" name:"Base64EncodedTriggerRules"`
	// 事件总线ID。

	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`
	// scf函数类型。当后端类型是SCF时生效。支持事件触发(EVENT)，http直通云函数(HTTP)。

	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`
	// EIAM应用ID。

	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`
	// EIAM应用类型。

	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`
	// EIAM认证类型。

	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`
	// Token有效时间。

	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`
}

type BindApiInfo struct {

	// api唯一id

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// Service唯一id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// api名字

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 绑定时间

	BindTime *string `json:"BindTime,omitempty" name:"BindTime"`
}

type BindApiAppRequest struct {
	*tchttp.BaseRequest

	// 待绑定的应用唯一 ID 。

	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
	// 待绑定的环境。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 待绑定的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待绑定的API唯一ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *BindApiAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindApiAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpstreamInfo struct {

	// 查询总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 查询列表

	UpstreamSet []*UpstreamInfo `json:"UpstreamSet,omitempty" name:"UpstreamSet"`
}

type DescribeIPStrategyApisStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境绑定API列表。

		Result *IPStrategyApiStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStrategyApisStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPStrategyApisStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateApiDocumentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api文档&sdk链接。

		Result *DocumentSDK `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateApiDocumentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateApiDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConstantParameter struct {

	// 常量参数名称。只有 ServiceType 是 HTTP 才会用到此参数。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 常量参数描述。只有 ServiceType 是 HTTP 才会用到此参数。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 常量参数位置。只有 ServiceType 是 HTTP 才会用到此参数。

	Position *string `json:"Position,omitempty" name:"Position"`
	// 常量参数默认值。只有 ServiceType 是 HTTP 才会用到此参数。

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
}

type Service struct {

	// 内网访问https端口。

	InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`
	// 用户自定义的服务描述。

	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`
	// 服务的前端请求类型。如http、https 或者 http&https。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// 服务支持的网络类型。

	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
	// 独占集群名称。

	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`
	// 服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// IP版本。

	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
	// 已经发布的环境列表。如test、prepub、release。

	AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments"`
	// 用户自定义的服务名称。

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 系统为该服务分配的外网域名。

	OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`
	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 内网访问http端口。

	InnerHttpPort *uint64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`
	// 系统为该服务自动分配的内网域名。

	InnerSubDomain *string `json:"InnerSubDomain,omitempty" name:"InnerSubDomain"`
	// 服务的计费状态。

	TradeIsolateStatus *int64 `json:"TradeIsolateStatus,omitempty" name:"TradeIsolateStatus"`
	// 服务绑定的标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 独享实例

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群类型

	SetType *string `json:"SetType,omitempty" name:"SetType"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// CPU架构，可选项[arm, x86]

	TceArch *string `json:"TceArch,omitempty" name:"TceArch"`
}

type ApiAppApiInfos struct {

	// 数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 应用绑定的Api信息数组

	ApiAppApiSet []*ApiAppApiInfo `json:"ApiAppApiSet,omitempty" name:"ApiAppApiSet"`
}

type RunApiReturn struct {

	// API 接口的响应头部。

	ReturnHeader *string `json:"ReturnHeader,omitempty" name:"ReturnHeader"`
	// API 接口的响应包体。

	ReturnBody *string `json:"ReturnBody,omitempty" name:"ReturnBody"`
	// API 接口的响应码。

	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
	// API 接口的响应延迟，单位 ms。

	Delay *int64 `json:"Delay,omitempty" name:"Delay"`
}

type CheckCloneApisRequest struct {
	*tchttp.BaseRequest

	// 源 serviceid

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 目的 serviceid

	DestServiceId *string `json:"DestServiceId,omitempty" name:"DestServiceId"`
	// 目的地域  gz hk sh bj 等缩写

	DestRegion *string `json:"DestRegion,omitempty" name:"DestRegion"`
	// 源Apiid 列表

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *CheckCloneApisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCloneApisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountSettingsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAccountSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建的插件详情。

		Result *Plugin `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExclusiveSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *ExclusiveSetList `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExclusiveSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExclusiveSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPStrategyApi struct {

	// API 唯一 ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 用户自定义的 API 名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// API 类型。取值为NORMAL（普通API）和TSF （微服务API）。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
	// API 的路径。如 /path。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API 的请求方法。如 GET。

	Method *string `json:"Method,omitempty" name:"Method"`
	// API 已经绑定的其他策略唯一ID。

	OtherIPStrategyId *string `json:"OtherIPStrategyId,omitempty" name:"OtherIPStrategyId"`
	// API 已经绑定的环境。

	OtherEnvironmentName *string `json:"OtherEnvironmentName,omitempty" name:"OtherEnvironmentName"`
}

type DescribeTcbScfApisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 满足条件查询结果总数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// TCB-SCF-HTTP触发器API列表。

		TcbScfApiSet []*TcbScfApi `json:"TcbScfApiSet,omitempty" name:"TcbScfApiSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTcbScfApisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcbScfApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 绑定类型，取值为 API、SERVICE，默认值为 SERVICE。

	BindType *string `json:"BindType,omitempty" name:"BindType"`
	// 待绑定的使用计划唯一 ID 列表。

	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`
	// 待解绑的服务环境。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 待解绑的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API 唯一 ID 数组，当 BindType=API 时，需要传入此参数。

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *UnBindEnvironmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindEnvironmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountUsageSecretIdCount struct {

	// 密钥使用计数

	SecretIdCount *int64 `json:"SecretIdCount,omitempty" name:"SecretIdCount"`
}

type Version struct {

	// version id

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
	// version 描述

	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 环境列表

	Environments []*string `json:"Environments,omitempty" name:"Environments"`
}

type DeleteAPIDocRequest struct {
	*tchttp.BaseRequest

	// API文档ID

	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *DeleteAPIDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAPIDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名状态

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiIdStatus struct {

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API唯一ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// API描述

	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`
	// API PATH。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API METHOD。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 服务创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 服务修改时间。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// API名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// VPC唯一ID。

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// API类型。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
	// API协议。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 是否买后调试。

	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`
	// 授权类型。

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// API业务类型。

	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`
	// 关联授权API唯一ID。

	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`
	// 授权API关联的业务API列表。

	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds"`
	// oauth配置信息。

	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`
	// oauth2.0API请求，token存放位置。

	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`
}

type CloudNativeAPIGatewayConfig struct {

	// 控制台类型。

	ConsoleType *string `json:"ConsoleType,omitempty" name:"ConsoleType"`
	// HTTP链接地址。

	HttpUrl *string `json:"HttpUrl,omitempty" name:"HttpUrl"`
	// HTTPS链接地址。

	HttpsUrl *string `json:"HttpsUrl,omitempty" name:"HttpsUrl"`
	// 网络类型, Open|Internal。

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 管理员用户名。

	AdminUser *string `json:"AdminUser,omitempty" name:"AdminUser"`
	// 管理员密码。

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
}

type Usage struct {

	// 服务使用量

	ServiceCount *AccountUsageServiceCount `json:"ServiceCount,omitempty" name:"ServiceCount"`
	// 使用计划使用量

	UsagePlanCount *AccountUsageUsagePlanCount `json:"UsagePlanCount,omitempty" name:"UsagePlanCount"`
	// 密钥使用量

	SecretIdCount *AccountUsageSecretIdCount `json:"SecretIdCount,omitempty" name:"SecretIdCount"`
	// 日志使用量

	LogRuleClount *AccountUsageLogRuleClount `json:"LogRuleClount,omitempty" name:"LogRuleClount"`
}

type UsagePlanBindSecretStatus struct {

	// 使用计划绑定密钥的数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 密钥详情列表。

	AccessKeyList []*UsagePlanBindSecret `json:"AccessKeyList,omitempty" name:"AccessKeyList"`
}

type BindSubDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSubDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewaySystemParameter struct {

	// 参数名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数内容。

	Value *string `json:"Value,omitempty" name:"Value"`
	// 参数默认值。

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 参数可修改值。

	ModifiableValue *string `json:"ModifiableValue,omitempty" name:"ModifiableValue"`
	// 是否需要重启生效。

	NeedReload *bool `json:"NeedReload,omitempty" name:"NeedReload"`
}

type EnvironmentStrategy struct {

	// 环境名

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 限流值

	Quota *int64 `json:"Quota,omitempty" name:"Quota"`
	// 限流最大值

	MaxQuota *int64 `json:"MaxQuota,omitempty" name:"MaxQuota"`
}

type IPStrategy struct {

	// 策略唯一ID。

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 用户自定义策略名称。

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略类型。支持WHITE（白名单）和BLACK（黑名单）。

	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`
	// IP列表。

	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 修改时间。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 策略绑定的API数量。

	BindApiTotalCount *int64 `json:"BindApiTotalCount,omitempty" name:"BindApiTotalCount"`
	// 绑定的API详情。

	BindApis []*DesApisStatus `json:"BindApis,omitempty" name:"BindApis"`
}

type ServiceReleaseHistory struct {

	// 发布版本总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 历史版本列表。

	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList"`
}

type CloudNativeAPIGatewayClsConfig struct {

	// CLS日志集ID。

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// CLS主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志类型, KongErrorLog|KongAccessLog。

	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

type CreateUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 用户自定义的使用计划名称。

	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`
	// 用户自定义的使用计划描述。

	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`
	// 请求配额总数，取值范围为-1或者[1, 99999999]，默认为-1，表示不开启。

	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`
	// 每秒请求限制数，取值范围为-1或者[1, 2000]，默认-1，表示不开启。

	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *CreateUsagePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUsagePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllPluginApisRequest struct {
	*tchttp.BaseRequest

	// 要查询的服务ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 要查询的插件ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
	// 环境信息。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。预留字段，目前不支持过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序依据，默认按照插件创建时间排序。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序方式，默认倒序排序。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeAllPluginApisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllPluginApisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindTriggerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解绑操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEIAMApisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API 详情列表。

		Result *ApisStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEIAMApisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEIAMApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiKeyFilter struct {

	// 密钥id

	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
	// 密钥key

	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
	// 未绑定的使用计划

	NotUsagePlanId *string `json:"NotUsagePlanId,omitempty" name:"NotUsagePlanId"`
	// 密钥状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 密钥名称

	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

type RequestConfig struct {

	// API 的路径，如 /path。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API 的请求方法，如 GET。

	Method *string `json:"Method,omitempty" name:"Method"`
}

type DemoteServiceUsagePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 降级操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DemoteServiceUsagePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DemoteServiceUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountLimitLogRuleClount struct {

	// 日志规则

	LogRuleClount *int64 `json:"LogRuleClount,omitempty" name:"LogRuleClount"`
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest

	// 待删除服务的唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 跳过删除前置条件校验（仅支持独享实例上的服务）

	SkipVerification *int64 `json:"SkipVerification,omitempty" name:"SkipVerification"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAppBindApisStatusRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。支持ApiId、ApiName、ServiceId、Environment 、KeyWord（ 可以匹配name或者ID）。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 应用ID

	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

func (r *DescribeApiAppBindApisStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAppBindApisStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindTriggersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解绑操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindTriggersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindTriggersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiUsagePlanSet struct {

	// API 绑定的使用计划总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// API 绑定使用计划列表。

	ApiUsagePlanList []*ApiUsagePlan `json:"ApiUsagePlanList,omitempty" name:"ApiUsagePlanList"`
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest

	// 用户自定义的服务名称。

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 用户自定义的服务描述。

	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`
	// 服务的前端请求类型。如 http、https、http&https。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 独立集群名称，用于指定创建服务所在的独立集群。

	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`
	// 网络类型列表，用于指定支持的访问类型，INNER为内网访问，OUTER为外网访问。默认为OUTER。

	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
	// IP版本号，支持IPv4和IPv6，默认为IPv4。

	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
	// 集群名称。保留字段，tsf serverlss类型使用。

	SetServerName *string `json:"SetServerName,omitempty" name:"SetServerName"`
	// 用户类型。保留类型，serverless用户使用。

	AppIdType *string `json:"AppIdType,omitempty" name:"AppIdType"`
	// 标签。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 独享实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// CPU架构，可选项[arm, x86]

	TceArch *string `json:"TceArch,omitempty" name:"TceArch"`
}

func (r *CreateServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogRulesRequest struct {
	*tchttp.BaseRequest

	// 日志规则唯一ID数组。

	LogRuleIds []*string `json:"LogRuleIds,omitempty" name:"LogRuleIds"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud..com/document/api/213/15688)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud..com/document/api/213/15688)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 日志规则名称，支持模糊匹配。

	LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`
	// 环境名称。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
}

func (r *DescribeLogRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApiAppKeyRequest struct {
	*tchttp.BaseRequest

	// 应用唯一 ID。

	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
	// 应用的Key。

	ApiAppKey *string `json:"ApiAppKey,omitempty" name:"ApiAppKey"`
	// 应用的Secret。

	ApiAppSecret *string `json:"ApiAppSecret,omitempty" name:"ApiAppSecret"`
}

func (r *UpdateApiAppKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApiAppKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Plugin struct {

	// 插件ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
	// 插件名称。

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 插件类型。

	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`
	// 插件定义语句。

	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`
	// 插件描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 插件创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 插件修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// 插件绑定的API总数。

	AttachedApiTotalCount *int64 `json:"AttachedApiTotalCount,omitempty" name:"AttachedApiTotalCount"`
	// 插件绑定的API信息。

	AttachedApis []*AttachedApiInfo `json:"AttachedApis,omitempty" name:"AttachedApis"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 插件属性

	PluginAttribute *string `json:"PluginAttribute,omitempty" name:"PluginAttribute"`
	// 原始插件定义语句

	OriginPluginData *string `json:"OriginPluginData,omitempty" name:"OriginPluginData"`
}

type UpstreamNode struct {

	// vm实例id

	VmInstanceId *string `json:"VmInstanceId,omitempty" name:"VmInstanceId"`
	// IP（domain）

	Host *string `json:"Host,omitempty" name:"Host"`
	// 端口[0, 65535]

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 权重[0, 100], 0为禁用

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 染色标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type DescribeApiAppsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用列表。

		Result *ApiAppInfos `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiAppsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAppsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// API所属服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 环境列表。

	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`
	// API唯一ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiEnvironmentStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiEnvironmentStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增的密钥详情。

		Result *ApiKey `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceRequest struct {
	*tchttp.BaseRequest

	// 待切换服务的唯一 Id。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待切换的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 切换的版本号。

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
	// 本次的切换描述。

	UpdateDesc *string `json:"UpdateDesc,omitempty" name:"UpdateDesc"`
}

func (r *UpdateServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待刷新的自定义的域名，ServiceId和SubDomain同时传或者不传

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *UpdateSubDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindIPStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindIPStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiAppInfos struct {

	// 应用数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 应用信息数组

	ApiAppSet []*ApiAppInfo `json:"ApiAppSet,omitempty" name:"ApiAppSet"`
}

type ServiceEnvironmentStrategy struct {

	// 环境名。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 访问服务对应环境的url。

	Url *string `json:"Url,omitempty" name:"Url"`
	// 发布状态。

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 发布的版本号。

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
	// 限流值。

	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
	// 最大限流值

	MaxStrategy *int64 `json:"MaxStrategy,omitempty" name:"MaxStrategy"`
}

type ServicesStatus struct {

	// 服务列表总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 服务列表详情。

	ServiceSet []*Service `json:"ServiceSet,omitempty" name:"ServiceSet"`
}

type UsagePlanBindEnvironment struct {

	// 环境名。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

type CreateLogRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志规则名称。

		LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`
		// 服务唯一ID，表示所需上报日志的服务ID，不传代表所有服务。

		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
		// 环境名称，表示所需要上报日志的环境名称，不传代表所有环境。

		EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
		// 日志集ID。

		LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
		// 日志主题ID。

		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
		// 日志规则唯一ID。

		LogRuleId *string `json:"LogRuleId,omitempty" name:"LogRuleId"`
		// 创建时间。

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsagePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsagePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindSecretIdsRequest struct {
	*tchttp.BaseRequest

	// 待解绑的使用计划唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 待解绑的密钥 ID 数组。

	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

func (r *UnBindSecretIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindSecretIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceParameterInput struct {

	// ServiceRequestNumPreSec，ApiRequestNumPreSec

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyUpstreamRequest struct {
	*tchttp.BaseRequest

	// VPC通道名字

	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`
	// VPC通道描述

	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`
	// 后端协议，HTTP, HTTPS其中之一

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 后端访问类型

	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// 负载均衡算法目前支持ROUND_ROBIN

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
	// VPC唯一ID

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 请求重试次数，默认3次

	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`
	// 请求到后端的，host头

	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`
	// VPC通道唯一ID

	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
	// 后端节点列表

	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`
	// 健康检查配置

	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`
}

func (r *ModifyUpstreamRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpstreamRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiRequestConfig struct {

	// path

	Path *string `json:"Path,omitempty" name:"Path"`
	// 方法

	Method *string `json:"Method,omitempty" name:"Method"`
}

type ReleaseServiceRequest struct {
	*tchttp.BaseRequest

	// 待发布服务的唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// apiId列表，预留字段，默认全量api发布。

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
	// 待发布的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 本次的发布描述。

	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *ReleaseServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApisStatus struct {

	// 符合条件的 API 接口数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// API 接口列表。

	ApiIdStatusSet []*DesApisStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`
}

type DescribeUpstreamBindApisRequest struct {
	*tchttp.BaseRequest

	// 分页

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// vpc通道Id

	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
	// ServiceId和ApiId过滤查询

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUpstreamBindApisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpstreamBindApisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApiAppKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更新操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApiAppKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApiAppKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiResponseTime struct {

	// api id

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// api 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 响应时间（毫秒）

	ResponseTime *uint64 `json:"ResponseTime,omitempty" name:"ResponseTime"`
	// 用户可读的服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// api类型。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
}

type DomainSetList struct {

	// 域名名称。

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
	// 域名解析状态。1 表示正常解析，0 表示解析失败。

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 证书ID。

	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
	// 是否使用默认路径映射。

	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`
	// 自定义域名协议类型。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 网络类型（'INNER' 或 'OUTER'）。

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 是否将HTTP请求强制跳转 HTTPS，默认为false。参数为 true时，API网关会将所有使用该自定义域名的 HTTP 协议的请求重定向至 HTTPS 协议进行转发。

	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
	// 域名备案注册状态

	RegistrationStatus *bool `json:"RegistrationStatus,omitempty" name:"RegistrationStatus"`
}

type Environment struct {

	// 环境名称。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 访问路径。

	Url *string `json:"Url,omitempty" name:"Url"`
	// 发布状态，1 表示已发布，0 表示未发布。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 运行版本。

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

type CreateIPStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建的IP策略详情。

		Result *IPStrategy `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIPStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllPluginApisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 插件相关API的列表。

		Result *ApiInfoSummary `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllPluginApisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllPluginApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountUsageServiceCount struct {

	// 服务数量

	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
}

type KV struct {

	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务的唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 用户自定义的策略名称。

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略类型。支持WHITE（白名单）和BLACK（黑名单）。

	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`
	// 策略详情，多个ip 使用\n 分隔符分开。

	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *CreateIPStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIPStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApiAppRequest struct {
	*tchttp.BaseRequest

	// 应用唯一 ID。

	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

func (r *DeleteApiAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePluginsRequest struct {
	*tchttp.BaseRequest

	// 要查询的插件列表。

	PluginIds []*string `json:"PluginIds,omitempty" name:"PluginIds"`
	// 要查询的插件名称。

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 要查询的插件类型。

	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。预留字段，目前不支持过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序依据，默认按照插件创建时间排序。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序方式，默认倒序排序。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 插件属性，System：系统插件，Custom：自定义插件，默认值：System

	PluginAttribute *string `json:"PluginAttribute,omitempty" name:"PluginAttribute"`
}

func (r *DescribePluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableApiKeyRequest struct {
	*tchttp.BaseRequest

	// 待启用的密钥 ID。

	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *EnableApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TsfLoadBalanceConfReq struct {

	// 方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 是否会话保持

	SessionStickRequired *bool `json:"SessionStickRequired,omitempty" name:"SessionStickRequired"`
	// 会话保持超时时间

	SessionStickTimeout *int64 `json:"SessionStickTimeout,omitempty" name:"SessionStickTimeout"`
}

type DescribeApiEnvironmentApiKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API绑定密钥列表。

		Result *ApiEnvironmentApiKeys `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiEnvironmentApiKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiEnvironmentApiKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExclusiveInstancesRequest struct {
	*tchttp.BaseRequest

	// 分页查询，limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页查询，offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeExclusiveInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExclusiveInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunApiForMarketRequest struct {
	*tchttp.BaseRequest

	// 服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 接口id

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 请求的header

	RequestHeader *string `json:"RequestHeader,omitempty" name:"RequestHeader"`
	// 请求的query

	RequestQuery *string `json:"RequestQuery,omitempty" name:"RequestQuery"`
	// 请求的path

	RequestPath *string `json:"RequestPath,omitempty" name:"RequestPath"`
	// 请求的method

	RequestMethod *string `json:"RequestMethod,omitempty" name:"RequestMethod"`
	// 请求的body

	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
	// 请求的body（字典类型）

	RequestBodyDict *string `json:"RequestBodyDict,omitempty" name:"RequestBodyDict"`
	// 可选值（'application/json'， 'application/x-www-form-urlencoded'）

	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	// 私钥id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *RunApiForMarketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunApiForMarketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindSubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待绑定的自定义的域名。

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 服务支持协议，可选值为http、https、http&https。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 网络类型，可选值为OUTER、INNER。

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 是否使用默认路径映射，默认为 true。为 false 时，表示自定义路径映射，此时 PathMappingSet 必填。

	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`
	// 默认域名。

	NetSubDomain *string `json:"NetSubDomain,omitempty" name:"NetSubDomain"`
	// 待绑定自定义域名的证书唯一 ID。针对Protocol 为https或http&https可以选择上传。

	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
	// 自定义域名路径映射，最多输入三个Environment，并且只能分别取值“test”、 ”prepub“、”release“。

	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`
	// 是否将HTTP请求强制跳转 HTTPS，默认为false。参数为 true时，API网关会将所有使用该自定义域名的 HTTP 协议的请求重定向至 HTTPS 协议进行转发。

	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
}

func (r *BindSubDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSubDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentKeyMonitorUploadRequest struct {
	*tchttp.BaseRequest

	// 服务Id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DescribeServiceEnvironmentKeyMonitorUploadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceEnvironmentKeyMonitorUploadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiEnvironmentStrategyStataus struct {

	// API绑定的限流策略数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// API绑定的限流策略列表。

	ApiEnvironmentStrategySet []*ApiEnvironmentStrategy `json:"ApiEnvironmentStrategySet,omitempty" name:"ApiEnvironmentStrategySet"`
}

type AttachedPluginInfo struct {

	// 插件ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
	// 环境信息。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 绑定时间。

	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`
	// 插件名称。

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 插件类型。

	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`
	// 插件描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 插件定义语句。

	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`
}

type ModifyLogRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExclusiveSetList struct {

	// 返回总数

	TotolCount *int64 `json:"TotolCount,omitempty" name:"TotolCount"`
	// 独占集群列表

	ExclusiveSet []*ExclusiveSet `json:"ExclusiveSet,omitempty" name:"ExclusiveSet"`
}

type ReleaseService struct {

	// 发布时的备注信息填写。

	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
	// 发布的版本id。

	ReleaseVersion *string `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
}

type DescribeApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api绑定策略详情

		Result *ApiEnvironmentStrategyStataus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiEnvironmentStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExclusiveInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// 独享实例唯一id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeExclusiveInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExclusiveInstanceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiAppRequest struct {
	*tchttp.BaseRequest

	// 应用唯一 ID。

	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
	// 修改的应用名称

	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`
	// 修改的应用描述

	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyApiAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnReleaseServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下线操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnReleaseServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnReleaseServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayResult struct {

	// 云原生API网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关状态。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 云原生API网关所属可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 云原生API网关名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云原生API网关类型。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 云原生API网关版本。

	GatewayVersion *string `json:"GatewayVersion,omitempty" name:"GatewayVersion"`
	// 云原生API网关节点信息。

	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitempty" name:"NodeConfig"`
	// 云原生API网关vpc配置。

	VpcConfig *CloudNativeAPIGatewayVpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`
	// 云原生API网关描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 云原生API网关创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeUpstreamBindApis struct {

	// 总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 绑定的api信息

	BindApiSet []*BindApiInfo `json:"BindApiSet,omitempty" name:"BindApiSet"`
}

type ResponseErrorCodeReq struct {

	// 自定义响应配置错误码。

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 自定义响应配置错误信息。

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 自定义响应配置错误码备注。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 自定义错误码转换。

	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`
	// 是否需要开启错误码转换。

	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type BindSecretIdsRequest struct {
	*tchttp.BaseRequest

	// 待绑定的使用计划唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 待绑定的密钥 ID 数组。

	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

func (r *BindSecretIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSecretIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePluginsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApisStatusRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为 20，最大值为 100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// API过滤条件。支持ApiId、ApiName、ApiPath、ApiType、AuthRelationApiId、AuthType、ApiBuniessType、NotUsagePlanId、Environment、Tags (values为 $tag_key:tag_value的列表)、TagKeys （values 为 tag key的列表）。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApisStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApisStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloneApisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 克隆结果。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneApisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 待查询的使用计划唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 定类型，取值为 API、SERVICE，默认值为 SERVICE。

	BindType *string `json:"BindType,omitempty" name:"BindType"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanEnvironmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsagePlanEnvironmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnReleaseServiceRequest struct {
	*tchttp.BaseRequest

	// 待下线服务的唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 保留字段，待下线的API列表。

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
	// 待下线的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
}

func (r *UnReleaseServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnReleaseServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudNativeAPIGatewaySystemParametersResult struct {

	// 云原生API网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关系统参数数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 云原生API网关系统参数列表。

	SystemParameterList []*CloudNativeAPIGatewaySystemParameter `json:"SystemParameterList,omitempty" name:"SystemParameterList"`
}

type ResourcePackSetList struct {

	// 资源包id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域，默认是广州，表示全地域通用

	Region *string `json:"Region,omitempty" name:"Region"`
	// 资源包来源, free：免费包，activity：运营包

	Origin *string `json:"Origin,omitempty" name:"Origin"`
	// 资源包已使用量，保留4位小数，succ_req的单位是万次，out_traffic的单位是GB

	Used *string `json:"Used,omitempty" name:"Used"`
	// 资源包总量，保留4位小数，单位同上

	Total *string `json:"Total,omitempty" name:"Total"`
	// 资源包类型，succ_req：调用次数，out_traffic：出流量，不传表示所有类型

	GoodsType *string `json:"GoodsType,omitempty" name:"GoodsType"`
	// 资源包状态，valid：有效，exhaust：资源耗尽，expire：资源包到期，isolated: 被隔离，后面3种都是失效状态，不传表示所有类型

	GoodsStatus *string `json:"GoodsStatus,omitempty" name:"GoodsStatus"`
	// 资源包创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 资源包生效时间

	EnableTime *string `json:"EnableTime,omitempty" name:"EnableTime"`
	// 资源包到期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 是否允许退费

	AllowReturn *bool `json:"AllowReturn,omitempty" name:"AllowReturn"`
	// 订单号

	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

type DeleteApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailableApiInfo struct {

	// API ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// API名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// API类型。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
	// API路径。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API方法。

	Method *string `json:"Method,omitempty" name:"Method"`
	// API是否绑定其他插件。

	AttachedOtherPlugin *bool `json:"AttachedOtherPlugin,omitempty" name:"AttachedOtherPlugin"`
	// API是否绑定当前插件。

	IsAttached *bool `json:"IsAttached,omitempty" name:"IsAttached"`
}

type LogRule struct {

	// 日志规则唯一ID。

	LogRuleId *string `json:"LogRuleId,omitempty" name:"LogRuleId"`
	// 服务唯一ID，ALL表示全部服务。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 环境名称，ALL表示全部环境。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 日志集ID。

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 日志主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志规则名称

	LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 服务名称。

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type BindEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 待绑定的使用计划唯一 ID 列表。

	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`
	// 绑定类型，取值为API、SERVICE，默认值为SERVICE。

	BindType *string `json:"BindType,omitempty" name:"BindType"`
	// 待绑定的环境。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 待绑定的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API唯一ID数组，当bindType=API时，需要传入此参数。

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *BindEnvironmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindEnvironmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type APIDocInfo struct {

	// API文档ID

	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
	// API文档名称

	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`
	// API文档构建状态

	ApiDocStatus *string `json:"ApiDocStatus,omitempty" name:"ApiDocStatus"`
	// API文档API数量

	ApiCount *int64 `json:"ApiCount,omitempty" name:"ApiCount"`
	// API文档查看次数

	ViewCount *int64 `json:"ViewCount,omitempty" name:"ViewCount"`
	// API文档发布次数

	ReleaseCount *int64 `json:"ReleaseCount,omitempty" name:"ReleaseCount"`
	// API文档访问URI

	ApiDocUri *string `json:"ApiDocUri,omitempty" name:"ApiDocUri"`
	// API文档分享密码

	SharePassword *string `json:"SharePassword,omitempty" name:"SharePassword"`
	// API文档更新时间

	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
	// 服务ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 环境信息

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 生成API文档的API ID

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 生成API文档的API名称

	ApiNames []*string `json:"ApiNames,omitempty" name:"ApiNames"`
}

type DescribeApiAppsStatusRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。支持ApiAppId、ApiAppName、KeyWord（ 可以匹配name或者ID）。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiAppsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAppsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeysStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥列表。

		Result *ApiKeysStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiKeysStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiKeysStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 使用计划详情。

		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Limit struct {

	// 日志规则配额

	LogRuleClount *AccountLimitLogRuleClount `json:"LogRuleClount,omitempty" name:"LogRuleClount"`
	// 服务配额

	ServiceCount *AccountLimitServiceCount `json:"ServiceCount,omitempty" name:"ServiceCount"`
	// 使用计划配额

	UsagePlanCount *AccountLimitUsagePlanCount `json:"UsagePlanCount,omitempty" name:"UsagePlanCount"`
	// 密钥配额

	SecretIdCount *AccountLimitSecretIdCount `json:"SecretIdCount,omitempty" name:"SecretIdCount"`
}

type UsagePlan struct {

	// 环境名称。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 使用计划唯一ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 使用计划名称。

	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`
	// 使用计划描述。

	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`
	// 使用计划qps，-1表示没有限制。

	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
	// 使用计划时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 使用计划修改时间。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type BindSecretIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSecretIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSecretIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePluginApisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 插件绑定的API列表信息。

		Result *AttachedApiSummary `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePluginApisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAPIDocPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API文档基本信息

		Result *APIDoc `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAPIDocPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAPIDocPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PathMapping struct {

	// 路径。

	Path *string `json:"Path,omitempty" name:"Path"`
	// 发布环境，可选值为“test”、 ”prepub“、”release“。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type DescribeApiForApiAppRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API 接口唯一 ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// Api所属地域

	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

func (r *DescribeApiForApiAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiForApiAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportOpenApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// import返回参数

		Result *CreateApiRspSet `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportOpenApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportOpenApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkConfig struct {

	// 最大出带宽

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// EnableInternetInbound信息

	EnableInternetInbound *bool `json:"EnableInternetInbound,omitempty" name:"EnableInternetInbound"`
	// EnableInternetOutbound信息

	EnableInternetOutbound *bool `json:"EnableInternetOutbound,omitempty" name:"EnableInternetOutbound"`
	// InboundIpAddresses信息

	InboundIpAddresses []*string `json:"InboundIpAddresses,omitempty" name:"InboundIpAddresses"`
	// OutboundIpAddresses信息

	OutboundIpAddresses []*string `json:"OutboundIpAddresses,omitempty" name:"OutboundIpAddresses"`
}

type DescribeServiceForApiAppRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务所属的地域

	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

func (r *DescribeServiceForApiAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceForApiAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlansStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 使用计划列表。

		Result *UsagePlansStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlansStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsagePlansStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAPIDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API文档基本信息

		Result *APIDoc `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAPIDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPStrategyFilter struct {

	// 策略名称

	IPStrategyName *string `json:"IPStrategyName,omitempty" name:"IPStrategyName"`
}

type UnBindSecretIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解绑操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindSecretIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindSecretIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountLimitSecretIdCount struct {

	// 密钥计数

	SecretIdCount *int64 `json:"SecretIdCount,omitempty" name:"SecretIdCount"`
	// 密钥的使用计划

	UsagePlanCountForSecretId *int64 `json:"UsagePlanCountForSecretId,omitempty" name:"UsagePlanCountForSecretId"`
}

type DesApisStatus struct {

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API唯一ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 用户自定义的 API 接口描述。

	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`
	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// API 接口的名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// VPCID。

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// VPC唯一ID。

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// API类型。取值为NORMAL（普通API）和TSF（微服务API）。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
	// API协议。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 是否买后调试。（云市场预留字段）

	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`
	// API 鉴权类型。取值为SECRET（密钥对鉴权）、NONE（免鉴权）、OAUTH、EIAM。

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// OAUTH API的类型。当AuthType 为 OAUTH时该字段有效， 取值为NORMAL（业务API）和 OAUTH（授权API）。

	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`
	// 关联的授权API 唯一 ID，当AuthType为OAUTH且ApiBusinessType为NORMAL时生效。标示业务API绑定的oauth2.0授权API唯一ID。

	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`
	// OAUTH 配置信息。当AuthType是OAUTH时生效。

	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`
	// 授权API关联的业务API列表。

	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds"`
	// API关联的标签信息。

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// API 的路径，如 /path。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API 的请求方法，如 GET。

	Method *string `json:"Method,omitempty" name:"Method"`
	// EIAM 应用ID。

	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`
}

type DisableApiKeyRequest struct {
	*tchttp.BaseRequest

	// 待禁用的密钥 ID。

	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DisableApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改自定义域名操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpstreamHealthChecker struct {

	// 标识是否开启主动健康检查。

	EnableActiveCheck *bool `json:"EnableActiveCheck,omitempty" name:"EnableActiveCheck"`
	// 标识是否开启被动健康检查。

	EnablePassiveCheck *bool `json:"EnablePassiveCheck,omitempty" name:"EnablePassiveCheck"`
	// 主动健康检查时探测请求的路径。默认为"/"。

	ActiveCheckHttpPath *string `json:"ActiveCheckHttpPath,omitempty" name:"ActiveCheckHttpPath"`
	// 主动健康检查的探测请求超时，单位秒。默认为5秒。

	ActiveCheckTimeout *uint64 `json:"ActiveCheckTimeout,omitempty" name:"ActiveCheckTimeout"`
	// 主动健康检查的时间间隔，默认5秒。

	ActiveCheckInterval *uint64 `json:"ActiveCheckInterval,omitempty" name:"ActiveCheckInterval"`
	// 主动健康检查时探测请求的的请求头。

	ActiveRequestHeader []*UpstreamHealthCheckerReqHeaders `json:"ActiveRequestHeader,omitempty" name:"ActiveRequestHeader"`
	// 异常节点的状态自动恢复时间，单位秒。当只开启被动检查的话，必须设置为 > 0 的值，否则被动异常节点将无法恢复。默认30秒。

	UnhealthyTimeout *uint64 `json:"UnhealthyTimeout,omitempty" name:"UnhealthyTimeout"`
	// 健康检查时，判断为成功请求的 HTTP 状态码。

	HealthyHttpStatus *string `json:"HealthyHttpStatus,omitempty" name:"HealthyHttpStatus"`
	// 健康检查时，判断为失败请求的 HTTP 状态码。

	UnhealthyHttpStatus *string `json:"UnhealthyHttpStatus,omitempty" name:"UnhealthyHttpStatus"`
	// TCP连续错误阈值。0 表示禁用 TCP 检查。取值范围：[0, 254]。

	TcpFailureThreshold *uint64 `json:"TcpFailureThreshold,omitempty" name:"TcpFailureThreshold"`
	// 连续超时阈值。0 表示禁用超时检查。取值范围：[0, 254]。

	TimeoutThreshold *uint64 `json:"TimeoutThreshold,omitempty" name:"TimeoutThreshold"`
	// HTTP连续错误阈值。0 表示禁用HTTP检查。取值范围：[0, 254]。

	HttpFailureThreshold *uint64 `json:"HttpFailureThreshold,omitempty" name:"HttpFailureThreshold"`
}

type BindEnvironmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindEnvironmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachPluginRequest struct {
	*tchttp.BaseRequest

	// 要解绑的API网关插件ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
	// 要操作的服务ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 要操作API的环境。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 要解绑的API ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DetachPluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachPluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 待删除的使用计划唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DeleteUsagePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsagePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiEnvironmentApiKeysRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API唯一ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 环境名。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiEnvironmentApiKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiEnvironmentApiKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailablePluginApisRequest struct {
	*tchttp.BaseRequest

	// 要查询的插件ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
	// 环境信息。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。预留字段，目前不支持过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序依据，默认按照插件创建时间排序。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序方式，默认倒序排序。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 要查询的服务ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DescribeAvailablePluginApisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailablePluginApisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindTriggersRequest struct {
	*tchttp.BaseRequest

	// 六段式资源ID数组

	Triggers []*TriggerInfo `json:"Triggers,omitempty" name:"Triggers"`
	// 触发器类型，默认为apigw。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 函数名称

	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`
	// 函数的命名空间。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *UnbindTriggersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindTriggersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type APIDoc struct {

	// API文档ID

	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
	// API文档名称

	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`
	// API文档构建状态

	ApiDocStatus *string `json:"ApiDocStatus,omitempty" name:"ApiDocStatus"`
}

type Filter struct {

	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InstanceInfo struct {

	// 独享实例唯一id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 独享实例name

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 独享实例描述

	InstanceDescription *string `json:"InstanceDescription,omitempty" name:"InstanceDescription"`
	// 独享实例计费类型

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 独享实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 独享实例状态

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 独享实例创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 订单号

	DealName *string `json:"DealName,omitempty" name:"DealName"`
	// 资源ID同唯一id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type DemoteServiceUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 使用计划ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 待降级的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 环境名称。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DemoteServiceUsagePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DemoteServiceUsagePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API 详情。

		Result *ApiInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentKeyMonitorUploadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceEnvironmentKeyMonitorUploadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceEnvironmentKeyMonitorUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewayNodeConfig struct {

	// 节点配置, 1c2g|2c4g|4c8g|8c16g。

	Specification *string `json:"Specification,omitempty" name:"Specification"`
	// 节点数量，2-9。

	Number *int64 `json:"Number,omitempty" name:"Number"`
}

type DescribeServiceEnvironmentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务绑定环境详情。

		Result *ServiceEnvironmentSet `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceEnvironmentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApiRspSet struct {

	// 个数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 返回的数组

	ApiSet []*CreateApiRsp `json:"ApiSet,omitempty" name:"ApiSet"`
}

type ListCloudNativeAPIGatewayResult struct {

	// 总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 云原生API网关实例列表。

	GatewayList []*DescribeCloudNativeAPIGatewayResult `json:"GatewayList,omitempty" name:"GatewayList"`
}

type DeletePluginRequest struct {
	*tchttp.BaseRequest

	// 要删除的API网关插件的ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
}

func (r *DeletePluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudNativeAPIGatewaySystemParameterReq struct {

	// 参数名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type OauthConfig struct {

	// 公钥，用于验证用户token。

	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
	// token传递位置。

	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`
	// 重定向地址，用于引导用户登录操作。

	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

type DeleteUpstreamRequest struct {
	*tchttp.BaseRequest

	// 待删除的VPC通道唯一ID

	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
}

func (r *DeleteUpstreamRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpstreamRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSearchRequest struct {
	*tchttp.BaseRequest

	// 日志开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 日志结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 保留字段

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 单次要返回的日志条数，单次返回的最大条数为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据上次返回的ConText，获取后续的内容，最多可获取10000条

	ConText *string `json:"ConText,omitempty" name:"ConText"`
	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 保留字段

	Query *string `json:"Query,omitempty" name:"Query"`
	// 检索条件,支持的检索条件如下：
	// req_id：“=”
	// api_id：“=”
	// cip：“=”
	// uip：“:”
	// err_msg：“:”
	// rsp_st：“=” 、“!=” 、 “:” 、 “>” 、 “<”
	// req_t：”>=“ 、 ”<=“
	//
	// 说明：
	// “:”表示包含，“!=”表示不等于，字段含义见输出参数的LogSet说明

	LogQuerys []*LogQuery `json:"LogQuerys,omitempty" name:"LogQuerys"`
}

func (r *DescribeLogSearchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogSearchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentReleaseHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务发布历史。

		Result *ServiceReleaseHistory `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountLimitServiceCount struct {

	// 服务数量

	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
	// 服务的api数量

	ApiCountInService *int64 `json:"ApiCountInService,omitempty" name:"ApiCountInService"`
	// 服务的自定义域名数量

	DomainCountInService *int64 `json:"DomainCountInService,omitempty" name:"DomainCountInService"`
}

type ApiFilter struct {

	// api路径

	ApiPath *string `json:"ApiPath,omitempty" name:"ApiPath"`
	// api名称

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
}

type ServiceReleaseHistoryInfo struct {

	// 版本号。

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
	// 版本描述。

	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`
	// 版本发布时间。

	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
}

type ServiceReqCount struct {

	// 用户服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 请求数

	ReqCount *uint64 `json:"ReqCount,omitempty" name:"ReqCount"`
}

type BindIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 待绑定的IP策略所属的服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待绑定的IP策略唯一ID。

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// IP策略待绑定的环境。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// IP策略待绑定的API列表。

	BindApiIds []*string `json:"BindApiIds,omitempty" name:"BindApiIds"`
}

func (r *BindIPStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindIPStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApiRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 用户自定义的 API 名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 用户自定义的 API 接口描述。

	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`
	// API 类型，支持NORMAL（普通API）和TSF（微服务API），默认为NORMAL。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
	// API 鉴权类型。支持SECRET（密钥对鉴权）、NONE（免鉴权）、OAUTH、APP（应用认证）。默认为NONE。

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// API 的后端服务类型。支持HTTP、MOCK、TSF、SCF、WEBSOCKET、TARGET（内测）。

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// API 的后端服务超时时间，单位是秒。

	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`
	// API 的前端请求协议，支持HTTP和WEBSOCKET。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 请求的前端配置。

	RequestConfig *ApiRequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`
	// 是否开启跨域。

	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`
	// 常量参数。

	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`
	// 前端请求参数。

	RequestParameters []*RequestParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`
	// 当AuthType 为 OAUTH时，该字段有效， NORMAL：业务api OAUTH：授权API。

	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`
	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。

	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`
	// API绑定微服务服务列表。

	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`
	// 微服务的负载均衡配置。

	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`
	// 微服务的健康检查配置。

	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`
	// target类型后端资源信息。（内测阶段）

	TargetServices []*TargetServicesReq `json:"TargetServices,omitempty" name:"TargetServices"`
	// target类型负载均衡配置。（内测阶段）

	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`
	// target健康检查配置。（内测阶段）

	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`
	// scf 函数名称。当后端类型是SCF时生效。

	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`
	// scf websocket注册函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`
	// scf websocket清理函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`
	// scf websocket传输函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`
	// scf 函数命名空间。当后端类型是SCF时生效。

	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`
	// scf函数版本。当后端类型是SCF时生效。

	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`
	// scf websocket注册函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`
	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`
	// scf websocket传输函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`
	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`
	// scf websocket清理函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`
	// scf websocket清理函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`
	// 是否开启响应集成。当后端类型是SCF时生效。

	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`
	// 开始调试后计费。（云市场预留字段）

	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`
	// 是否删除自定义响应配置错误码，如果不传或者传 False，不删除，当传 True 时，则删除此 API 所有自定义响应配置错误码。

	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`
	// 返回类型。

	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`
	// 自定义响应配置成功响应示例。

	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`
	// 自定义响应配置失败响应示例。

	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`
	// API 的后端服务配置。

	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`
	// 关联的授权API 唯一 ID，当AuthType为OAUTH且ApiBusinessType为NORMAL时生效。标示业务API绑定的oauth2.0授权API唯一ID。

	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`
	// API的后端服务参数。

	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`
	// oauth配置。当AuthType是OAUTH时生效。

	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`
	// 用户自定义错误码配置。

	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`
	// tsf serverless 命名空间ID。（内测中）

	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`
	// 用户类型。

	UserType *string `json:"UserType,omitempty" name:"UserType"`
	// 是否打开Base64编码，只有后端是scf时才会生效。

	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`
	// 事件总线ID。

	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`
	// 内部参数，标示请求来源

	InternalReqSource *string `json:"InternalReqSource,omitempty" name:"InternalReqSource"`
	// scf函数类型。当后端类型是SCF时生效。支持事件触发(EVENT)，http直通云函数(HTTP)。

	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`
	// EIAM应用类型。

	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`
	// EIAM应用认证类型，支持仅认证（AuthenticationOnly）、认证和鉴权（Authorization）。

	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`
	// EIAM应用Token 有效时间，单位为秒，默认为7200秒。

	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`
	// EIAM应用ID。

	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`
}

func (r *CreateApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFreePkgResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开通状态，0 正常，1 已经存在免费额度包，2 未实名认证，3 db创建记录失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFreePkgResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFreePkgResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 限流策略列表。

		Result *ServiceEnvironmentStrategyStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindApiAppRequest struct {
	*tchttp.BaseRequest

	// 待绑定的应用唯一 ID 。

	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
	// 待绑定的环境。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 待绑定的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待绑定的API唯一ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *UnbindApiAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindApiAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 待解绑的服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待解绑的IP策略唯一ID。

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 待解绑的环境。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 待解绑的 API 列表。

	UnBindApiIds []*string `json:"UnBindApiIds,omitempty" name:"UnBindApiIds"`
}

func (r *UnBindIPStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindIPStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RequestParameter struct {

	// 请求参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 参数位置

	Position *string `json:"Position,omitempty" name:"Position"`
	// 参数类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 默认值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 是否必须

	Required *bool `json:"Required,omitempty" name:"Required"`
}

type CreateApiAppRequest struct {
	*tchttp.BaseRequest

	// 用户自定义应用名称。

	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`
	// 应用描述

	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateApiAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApiAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥详情。

		Result *ApiKey `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTopsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// top 10 访问数的服务列表。

		Top10ReqService []*ServiceReqCount `json:"Top10ReqService,omitempty" name:"Top10ReqService"`
		// top 10 访问数的api列表。

		Top10ReqApi []*ApiReqCount `json:"Top10ReqApi,omitempty" name:"Top10ReqApi"`
		// top 10 响应时间的api列表。

		Top10ResponseTimeApi []*ApiResponseTime `json:"Top10ResponseTimeApi,omitempty" name:"Top10ResponseTimeApi"`
		// top 10 错误率的api列表。

		Top10ErrorRateApi []*ApiErrorRate `json:"Top10ErrorRateApi,omitempty" name:"Top10ErrorRateApi"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MonitorTopsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MonitorTopsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务的唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 限流值。

	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
	// 环境列表。

	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`
}

func (r *ModifyServiceEnvironmentStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsagePlanInfo struct {

	// 使用计划唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 使用计划名称。

	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`
	// 使用计划描述。

	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`
	// 初始化调用次数。

	InitQuota *int64 `json:"InitQuota,omitempty" name:"InitQuota"`
	// 每秒请求限制数。

	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
	// 最大调用次数。

	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`
	// 是否隐藏。

	IsHide *int64 `json:"IsHide,omitempty" name:"IsHide"`
	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// 绑定密钥的数量。

	BindSecretIdTotalCount *int64 `json:"BindSecretIdTotalCount,omitempty" name:"BindSecretIdTotalCount"`
	// 绑定密钥的详情。

	BindSecretIds []*string `json:"BindSecretIds,omitempty" name:"BindSecretIds"`
	// 绑定环境数量。

	BindEnvironmentTotalCount *int64 `json:"BindEnvironmentTotalCount,omitempty" name:"BindEnvironmentTotalCount"`
	// 绑定环境详情。

	BindEnvironments []*UsagePlanBindEnvironment `json:"BindEnvironments,omitempty" name:"BindEnvironments"`
}

type ModifyServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsagePlanFilter struct {

	// 使用计划id

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 使用计划名称

	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`
	// 非service id

	NotServiceId *string `json:"NotServiceId,omitempty" name:"NotServiceId"`
	// 非api id

	NotApiId *string `json:"NotApiId,omitempty" name:"NotApiId"`
	// 环境

	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CheckActionToInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 独享实例id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 操作

		ActionToInstance *string `json:"ActionToInstance,omitempty" name:"ActionToInstance"`
		// True 允许，False 不允许

		Allowable *bool `json:"Allowable,omitempty" name:"Allowable"`
		// 错误原因码

		ReasonCode *int64 `json:"ReasonCode,omitempty" name:"ReasonCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckActionToInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckActionToInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExclusiveInstanceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 独享实例详情

		Result *InstanceDetail `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExclusiveInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExclusiveInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 用户自定义的 API 名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 用户自定义的 API 接口描述。

	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`
	// API 类型，支持NORMAL和TSF，默认为NORMAL。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
	// API 鉴权类型。支持SECRET、NONE、OAUTH、APP。默认为NONE。

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// API 的后端服务类型。支持HTTP、MOCK、TSF、CLB、SCF、WEBSOCKET、TARGET（内测）。

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 是否需要签名认证，True 表示需要，False 表示不需要。待废弃。

	AuthRequired *bool `json:"AuthRequired,omitempty" name:"AuthRequired"`
	// API 的后端服务超时时间，单位是秒。

	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`
	// API 的前端请求类型，如 HTTP 或 HTTPS 或者 HTTP 和 HTTPS。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 请求的前端配置。

	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`
	// 是否需要开启跨域，Ture 表示需要，False 表示不需要。

	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`
	// 常量参数。

	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`
	// 前端请求参数。

	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`
	// 当AuthType 为 OAUTH时，该字段有效， NORMAL：业务api   OAUTH：授权API。

	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`
	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。

	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`
	// API绑定微服务服务列表。

	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`
	// 微服务的负载均衡配置。

	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`
	// 微服务的健康检查配置。

	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`
	// target类型负载均衡配置。（内测阶段）

	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`
	// target健康检查配置。（内测阶段）

	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`
	// scf 函数名称。当后端类型是SCF时生效。

	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`
	// scf websocket注册函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`
	// scf websocket清理函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`
	// scf websocket传输函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`
	// scf 函数命名空间。当后端类型是SCF时生效。

	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`
	// scf函数版本。当后端类型是SCF时生效。

	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`
	// scf websocket注册函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`
	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`
	// scf websocket传输函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`
	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`
	// scf websocket清理函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`
	// scf websocket清理函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。

	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`
	// 是否开启响应集成。当后端类型是SCF时生效。

	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`
	// 开始调试后计费。（云市场预留字段）

	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`
	// 标签。

	TagSpecifications *Tag `json:"TagSpecifications,omitempty" name:"TagSpecifications"`
	// 是否删除自定义响应配置错误码，如果不传或者传 False，不删除，当传 True 时，则删除此 API 所有自定义响应配置错误码。

	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`
	// 返回类型。

	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`
	// 自定义响应配置成功响应示例。

	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`
	// 自定义响应配置失败响应示例。

	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`
	// API 的后端服务配置。

	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`
	// 关联的授权API 唯一 ID，当AuthType为OAUTH且ApiBusinessType为NORMAL时生效。标示业务API绑定的oauth2.0授权API唯一ID。

	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`
	// API的后端服务参数。

	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`
	// oauth配置。当AuthType是OAUTH时生效。

	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`
	// API 接口唯一 ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 用户自定义错误码配置。

	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`
	// 是否开启Base64编码，只有后端为scf时才会生效。

	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`
	// 是否开启Base64编码的header触发，只有后端为scf时才会生效。

	IsBase64Trigger *bool `json:"IsBase64Trigger,omitempty" name:"IsBase64Trigger"`
	// Header触发规则，总规则数不能超过10。

	Base64EncodedTriggerRules []*Base64EncodedTriggerRule `json:"Base64EncodedTriggerRules,omitempty" name:"Base64EncodedTriggerRules"`
	// 事件总线ID。

	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`
	// 内部参数，标识请求来源

	InternalReqSource *string `json:"InternalReqSource,omitempty" name:"InternalReqSource"`
	// scf函数类型。当后端类型是SCF时生效。支持事件触发(EVENT)，http直通云函数(HTTP)。

	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`
	// EIAM应用类型。

	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`
	// EIAM应用认证类型，支持仅认证（AuthenticationOnly）、认证和鉴权（Authorization）。

	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`
	// EIAM应用Token 有效时间，单位为秒，默认为7200秒。

	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`
	// EIAM应用ID。

	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`
}

func (r *ModifyApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiIncrementResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiIncrementResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiIncrementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiKeysStatus struct {

	// 符合条件的 API 密钥数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// API 密钥列表。

	ApiKeySet []*ApiKey `json:"ApiKeySet,omitempty" name:"ApiKeySet"`
}

type Base64EncodedTriggerRule struct {

	// 进行编码触发的header，可选值 "Accept"和"Content_Type" 对应实际数据流请求header中的Accept和 Content-Type。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 进行编码触发的header的可选值数组, 数组元素的字符串最大长度为40，元素可以包括数字，英文字母以及特殊字符，特殊字符的可选值为： `.`    `+`    `*`   `-`   `/`  `_`
	//
	// 例如 [
	//     "application/x-vpeg005",
	//     "application/xhtml+xml",
	//     "application/vnd.ms-project",
	//     "application/vnd.rn-rn_music_package"
	// ] 等都是合法的。

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribePluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 插件详情。

		Result *Plugin `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePluginApisRequest struct {
	*tchttp.BaseRequest

	// 查询的插件ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。预留字段，目前不支持过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序依据，默认按照插件绑定时间排序。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序方式，默认倒序排序。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribePluginApisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginApisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpstreamsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果

		Result *DescribeUpstreamInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpstreamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpstreamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindEnvironmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解绑操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindEnvironmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApisStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API 详情列表。

		Result *ApisStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApisStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApisStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTrendRequest struct {
	*tchttp.BaseRequest

	// 日期格式如"2018-12-28 09:50:04"。

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 日期格式如"2018-12-28 09:50:04"。

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 曲线粒度（可选60，3600，86400）

	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *MonitorTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MonitorTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedPluginSummary struct {

	// 已绑定的插件总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 已绑定的插件信息。

	PluginSummary []*AttachedPluginInfo `json:"PluginSummary,omitempty" name:"PluginSummary"`
}

type ImportOpenApiRequest struct {
	*tchttp.BaseRequest

	// API所在的服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// Content格式，只能是YAML或者JSON，默认是YAML。

	EncodeType *string `json:"EncodeType,omitempty" name:"EncodeType"`
	// openAPI正文内容。

	Content *string `json:"Content,omitempty" name:"Content"`
	// Content版本，默认是openAPI，目前只支持openAPI。

	ContentVersion *string `json:"ContentVersion,omitempty" name:"ContentVersion"`
}

func (r *ImportOpenApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportOpenApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateExclusiveInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateExclusiveInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateExclusiveInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExclusiveSet struct {

	// 集群名字

	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`
	// 集群状态

	ExclusiveStatus *string `json:"ExclusiveStatus,omitempty" name:"ExclusiveStatus"`
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务唯一ID。

		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
		// 服务 环境列表。

		AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments"`
		// 服务名称。

		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
		// 服务描述。

		ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`
		// 服务支持协议，可选值为http、https、http&https。

		Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
		// 服务创建时间。

		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
		// 服务修改时间。

		ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
		// 独立集群名称。

		ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`
		// 网络类型列表，INNER为内网访问，OUTER为外网访问。

		NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
		// 内网访问子域名。

		InternalSubDomain *string `json:"InternalSubDomain,omitempty" name:"InternalSubDomain"`
		// 外网访问子域名。

		OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`
		// 内网访问http服务端口号。

		InnerHttpPort *int64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`
		// 内网访问https端口号。

		InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`
		// API总数。

		ApiTotalCount *int64 `json:"ApiTotalCount,omitempty" name:"ApiTotalCount"`
		// API列表。

		ApiIdStatusSet []*ApiIdStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`
		// 使用计划总数量。

		UsagePlanTotalCount *int64 `json:"UsagePlanTotalCount,omitempty" name:"UsagePlanTotalCount"`
		// 使用计划数组。

		UsagePlanList []*UsagePlan `json:"UsagePlanList,omitempty" name:"UsagePlanList"`
		// IP版本。

		IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
		// 此服务的用户类型。

		UserType *string `json:"UserType,omitempty" name:"UserType"`
		// 预留字段。

		SetId *int64 `json:"SetId,omitempty" name:"SetId"`
		// 服务绑定的标签。

		Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
		// 独享实例id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 独享实例name

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 集群类型

		SetType *string `json:"SetType,omitempty" name:"SetType"`
		// 项目ID

		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
		// CPU架构

		TceArch *string `json:"TceArch,omitempty" name:"TceArch"`
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

type ReleaseServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发布信息。

		Result *ReleaseService `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateExclusiveInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TerminateExclusiveInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateExclusiveInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheckConf struct {

	// 是否开启健康检查。

	IsHealthCheck *bool `json:"IsHealthCheck,omitempty" name:"IsHealthCheck"`
	// 健康检查阈值。

	RequestVolumeThreshold *int64 `json:"RequestVolumeThreshold,omitempty" name:"RequestVolumeThreshold"`
	// 窗口大小。

	SleepWindowInMilliseconds *int64 `json:"SleepWindowInMilliseconds,omitempty" name:"SleepWindowInMilliseconds"`
	// 阈值百分比。

	ErrorThresholdPercentage *int64 `json:"ErrorThresholdPercentage,omitempty" name:"ErrorThresholdPercentage"`
}

type CreateTcbScfApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境ID。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 请求PATH。

		Path *string `json:"Path,omitempty" name:"Path"`
		// 请求子域名。

		SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTcbScfApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTcbScfApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAPIDocsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API文档列表信息

		Result *APIDocs `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAPIDocsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAPIDocsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayConfigResult struct {

	// 云原生API网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 云原生API网关配置列表。

	ConfigList []*CloudNativeAPIGatewayConfig `json:"ConfigList,omitempty" name:"ConfigList"`
}

type BuildAPIDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BuildAPIDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BuildAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApiKeyRequest struct {
	*tchttp.BaseRequest

	// 用户自定义密钥名称。

	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
	// 密钥类型，支持 auto 和 manual（自定义密钥），默认为 auto。

	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`
	// 用户自定义密钥 ID，AccessKeyType 为 manual 时必传。长度为5 - 50字符，由字母、数字、英文下划线组成。

	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
	// 用户自定义密钥 Key，AccessKeyType 为 manual 时必传。长度为10 - 50字符，由字母、数字、英文下划线。

	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *CreateApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePluginsByApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 插件可绑定的API列表信息。

		Result *AttachedPluginSummary `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePluginsByApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginsByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainSets struct {

	// 服务下的自定义域名数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 自定义服务域名列表。

	DomainSet []*DomainSetList `json:"DomainSet,omitempty" name:"DomainSet"`
}

type EnvironmentUpload struct {

	// 环境

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 是否上报

	IsUpdate *bool `json:"IsUpdate,omitempty" name:"IsUpdate"`
}

type DeleteServiceSubDomainMappingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除自定义域名的路径映射操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceSubDomainMappingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceSubDomainMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiForApiAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API 详情。

		Result *ApiInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiForApiAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiForApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExclusiveInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 独享实例详情信息

		Result *InstanceDetail `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyExclusiveInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExclusiveInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceReleaseVersionRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceReleaseVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceReleaseVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIPStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIPStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountLimitUsagePlanCount struct {

	// 使用计划数量

	UsagePlanCount *int64 `json:"UsagePlanCount,omitempty" name:"UsagePlanCount"`
	// 密钥使用计划

	SecretIdCountInUsagePlan *int64 `json:"SecretIdCountInUsagePlan,omitempty" name:"SecretIdCountInUsagePlan"`
	// QPS使用计划

	MaxQPSInUsagePlan *int64 `json:"MaxQPSInUsagePlan,omitempty" name:"MaxQPSInUsagePlan"`
	// 阶段计数使用计划

	StageCountInUsagePlan *int64 `json:"StageCountInUsagePlan,omitempty" name:"StageCountInUsagePlan"`
}

type ApiErrorRate struct {

	// 用户可读的api id

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// api路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// api 错误率

	ErrorRate *float64 `json:"ErrorRate,omitempty" name:"ErrorRate"`
	// 用户可读的服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// api类型

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
}

type UpstreamInfo struct {

	// VPC通道唯一ID

	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
	// VPC通道名字

	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`
	// VPC通道描述

	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`
	// 写意

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 负载均衡算法

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
	// vpc唯一ID

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 请求重拾次数

	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`
	// 后端节点

	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 健康检查配置

	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`
}

type CloneApisRequest struct {
	*tchttp.BaseRequest

	// 复制的源服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 复制的目的服务唯一ID。

	DestServiceId *string `json:"DestServiceId,omitempty" name:"DestServiceId"`
	// 复制目的地域。  gz cq sh 等缩写。

	DestRegion *string `json:"DestRegion,omitempty" name:"DestRegion"`
	// 复制的源API列表。

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *CloneApisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneApisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api信息

		Result *CreateApiRsp `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcConfig struct {

	// vpcid

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// subnetid

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
}

type DescribeServiceSubDomainMappingsRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务绑定的自定义域名。

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *DescribeServiceSubDomainMappingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceSubDomainMappingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanEnvironmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 使用计划绑定详情。

		Result *UsagePlanEnvironmentStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanEnvironmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsagePlanEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApisRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待删除的API列表。

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
	// 内部参数，标识请求来源。

	InternalReqSource *string `json:"InternalReqSource,omitempty" name:"InternalReqSource"`
}

func (r *DeleteApisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApiAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceFilter struct {

	// 服务唯一id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 没有绑定的使用计划

	NotUsagePlanId *string `json:"NotUsagePlanId,omitempty" name:"NotUsagePlanId"`
	// 环境

	Environments *string `json:"Environments,omitempty" name:"Environments"`
	// ip版本

	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
}

type TsfLoadBalanceConfResp struct {

	// 是否开启负载均衡。

	IsLoadBalance *bool `json:"IsLoadBalance,omitempty" name:"IsLoadBalance"`
	// 负载均衡方式。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 是否开启会话保持。

	SessionStickRequired *bool `json:"SessionStickRequired,omitempty" name:"SessionStickRequired"`
	// 会话保持超时时间。

	SessionStickTimeout *int64 `json:"SessionStickTimeout,omitempty" name:"SessionStickTimeout"`
}

type DeleteApiRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API 接口唯一 ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 内部参数，标识请求来源。

	InternalReqSource *string `json:"InternalReqSource,omitempty" name:"InternalReqSource"`
}

func (r *DeleteApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiUsagePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api绑定使用计划列表。

		Result *ApiUsagePlanSet `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiUsagePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceUsagePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务绑定使用计划列表。

		Result *ServiceUsagePlanSet `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceUsagePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUpstreamResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功删除的vpc通道id

		UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUpstreamResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MicroService struct {

	// 微服务集群ID。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 微服务命名空间ID。

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 微服务名称。

	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

type DescribeAPIDocDetailRequest struct {
	*tchttp.BaseRequest

	// API文档ID

	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *DescribeAPIDocDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAPIDocDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceForApiAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务唯一ID。

		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
		// 服务 环境列表。

		AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments"`
		// 服务名称。

		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
		// 服务描述。

		ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`
		// 服务支持协议，可选值为http、https、http&https。

		Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
		// 服务创建时间。

		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
		// 服务修改时间。

		ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
		// 独立集群名称。

		ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`
		// 网络类型列表，INNER为内网访问，OUTER为外网访问。

		NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
		// 内网访问子域名。

		InternalSubDomain *string `json:"InternalSubDomain,omitempty" name:"InternalSubDomain"`
		// 外网访问子域名。

		OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`
		// 内网访问http服务端口号。

		InnerHttpPort *int64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`
		// 内网访问https端口号。

		InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`
		// API总数。

		ApiTotalCount *int64 `json:"ApiTotalCount,omitempty" name:"ApiTotalCount"`
		// API列表。

		ApiIdStatusSet []*ApiIdStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`
		// 使用计划总数量。

		UsagePlanTotalCount *int64 `json:"UsagePlanTotalCount,omitempty" name:"UsagePlanTotalCount"`
		// 使用计划数组。

		UsagePlanList []*UsagePlan `json:"UsagePlanList,omitempty" name:"UsagePlanList"`
		// IP版本。

		IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
		// 此服务的用户类型。

		UserType *string `json:"UserType,omitempty" name:"UserType"`
		// 预留字段。

		SetId *int64 `json:"SetId,omitempty" name:"SetId"`
		// 服务绑定的标签。

		Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceForApiAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceForApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiEnvironmentStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiTags struct {

	// 标签的key。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签的values。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ErrorCodes struct {

	// 自定义响应配置错误码。

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 自定义响应配置错误信息。

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 自定义响应配置错误码备注。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 自定义错误码转换。

	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`
	// 是否需要开启错误码转换。

	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type Price struct {

	// 请求数计费询价

	RequestPrice *PriceInfos `json:"RequestPrice,omitempty" name:"RequestPrice"`
	// 流量计费询价

	BandwidthPrice *PriceInfos `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

type ServiceUsagePlanSet struct {

	// 服务上绑定的使用计划总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 服务上绑定的使用计划列表。

	ServiceUsagePlanList []*ApiUsagePlan `json:"ServiceUsagePlanList,omitempty" name:"ServiceUsagePlanList"`
}

type DescribeIPStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IP策略详情。

		Result *IPStrategy `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindIPStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解绑操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindIPStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentKeyMonitorUploadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *EnvironmentList `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentKeyMonitorUploadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceEnvironmentKeyMonitorUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSubDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSubDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceDetail struct {

	// 独享实例唯一id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 独享实例名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 独享实例描述

	InstanceDescription *string `json:"InstanceDescription,omitempty" name:"InstanceDescription"`
	// 独享实例计费类型

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 独享实例状态

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 独享实例预付费类型

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 独享实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 独享实例网络类型

	NetworkConfig *NetworkConfig `json:"NetworkConfig,omitempty" name:"NetworkConfig"`
	// 独享实例vpc配置

	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`
	// 独享实例参数配置

	Parameters []*ParameterInfo `json:"Parameters,omitempty" name:"Parameters"`
	// 独享实例隔离时间

	IsolationStartedTime *string `json:"IsolationStartedTime,omitempty" name:"IsolationStartedTime"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 订单号

	DealName *string `json:"DealName,omitempty" name:"DealName"`
	// 可用区列表

	Zones *string `json:"Zones,omitempty" name:"Zones"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type AdminInfo struct {

	// vip组id

	VipgroupId *uint64 `json:"VipgroupId,omitempty" name:"VipgroupId"`
	// 集群id

	SetId *uint64 `json:"SetId,omitempty" name:"SetId"`
	// 服务的内外网信息

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 服务对应的ld信息

	LdList []*string `json:"LdList,omitempty" name:"LdList"`
	// 服务对应的vip信息

	VipList []*string `json:"VipList,omitempty" name:"VipList"`
}

type CheckActionToInstanceRequest struct {
	*tchttp.BaseRequest

	// 独享实例

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 操作名称。默认 销毁独享实例

	ActionToInstance *string `json:"ActionToInstance,omitempty" name:"ActionToInstance"`
}

func (r *CheckActionToInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckActionToInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentListRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceEnvironmentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// IP 策略唯一ID。

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略关联的环境。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。预留字段，目前不支持过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategysStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的策略列表。

		Result *IPStrategysStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStrategysStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPStrategysStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsagePlanEnvironmentStatus struct {

	// 使用计划绑定的服务的环境数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 使用计划已经绑定的各个服务的环境状态。

	EnvironmentList []*UsagePlanEnvironment `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
}

type DeleteLogRulesRequest struct {
	*tchttp.BaseRequest

	// 日志规则ID列表。

	LogRuleIds []*string `json:"LogRuleIds,omitempty" name:"LogRuleIds"`
}

func (r *DeleteLogRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceUsagePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceUsagePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApisAuthRelationApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果，成功返回true

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApisAuthRelationApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApisAuthRelationApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceSubDomainMappings struct {

	// 是否使用默认路径映射，为 True 表示使用默认路径映射；为 False 的话，表示使用自定义路径映射，此时 PathMappingSet 不为空。

	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`
	// 自定义路径映射列表。

	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`
}

type InquiryPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价数据列表

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiGatewayLog struct {

	// log_format
	// '[$app_id][$env_name][$service_id][$http_host][$api_id][$uri][$scheme][rsp_st:$status][ups_st:$upstream_status]'
	// '[cip:$remote_addr][uip:$upstream_addr][vip:$server_addr][rsp_len:$bytes_sent][req_len:$request_length]'
	// '[req_t:$request_time][ups_rsp_t:$upstream_response_time][ups_conn_t:$upstream_connect_time][ups_head_t:$upstream_header_time]'
	// '[err_msg:$err_msg][tcp_rtt:$tcpinfo_rtt][$pid][$time_local]';
	//
	// app_id： 用户 ID。
	// env_name：环境名称。
	// service_id： 服务 ID。
	// http_host： 域名。
	// api_id： API 的 ID。
	// uri：请求的路径。
	// scheme： HTTP/HTTPS 协议。
	// rsp_st： 请求响应状态码。
	// ups_st： 后端业务服务器的响应状态码（如果请求透传到后端，改变量不为空。如果请求在 APIGW 就被拦截了，那么该变量显示为 -）。
	// cip： 客户端 IP。
	// uip： 后端业务服务（upstream）的 IP。
	// vip： 请求访问的 VIP。
	// rsp_len： 响应长度。
	// req_len： 请求长度。
	// req_t： 请求响应的总时间。
	// ups_rsp_t： 后端响应的总时间（apigw 建立连接到接收到后端响应的时间）。
	// ups_conn_t: 与后端业务服务器连接建立成功时间。
	// ups_head_t：后端响应的头部到达时间。
	// err_msg： 错误信息。
	// tcp_rtt： 客户端 TCP 连接信息，RTT（Round Trip Time）由三部分组成：链路的传播时间（propagation delay）、末端系统的处理时间、路由器缓存中的排队和处理时间（queuing delay）。

	LogData *string `json:"LogData,omitempty" name:"LogData"`
}

type UsagePlanBindSecret struct {

	// 密钥ID。

	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
	// 密钥名称。

	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
	// 密钥状态，0表示已禁用，1表示启用中。

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DeleteApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Result *AccountSettings `json:"Result,omitempty" name:"Result"`
		// 请求标识

		GetAccountSettings *string `json:"GetAccountSettings,omitempty" name:"GetAccountSettings"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAccountSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSet struct {

	// 返回码 0 ok 非0 失败

	CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 错误信息

	CheckMessageList []*string `json:"CheckMessageList,omitempty" name:"CheckMessageList"`
}

type TargetServicesReq struct {

	// docker ip

	DockerIp *string `json:"DockerIp,omitempty" name:"DockerIp"`
	// vm ip

	VmIp *string `json:"VmIp,omitempty" name:"VmIp"`
	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vm port

	VmPort *int64 `json:"VmPort,omitempty" name:"VmPort"`
	// cvm所在宿主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
}

type UsagePlanStatusInfo struct {

	// 使用计划唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 用户自定义的使用计划名称。

	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`
	// 用户自定义的使用计划描述。

	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`
	// 每秒最大请求次数。

	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
	// 请求配额总量，-1表示没有限制。

	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`
	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type DeleteTcbScfApiRequest struct {
	*tchttp.BaseRequest

	// 环境ID。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// SCF方法名称。

	ScfFunctionName *string `json:"ScfFunctionName,omitempty" name:"ScfFunctionName"`
	// SCF方法命名空间。

	ScfFunctionNamespace *string `json:"ScfFunctionNamespace,omitempty" name:"ScfFunctionNamespace"`
}

func (r *DeleteTcbScfApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTcbScfApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateApiDocumentRequest struct {
	*tchttp.BaseRequest

	// 待创建文档的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待创建 SDK 的服务所在环境。

	GenEnvironment *string `json:"GenEnvironment,omitempty" name:"GenEnvironment"`
	// 待创建 SDK 的语言。当前只支持 Python 和 JavaScript。

	GenLanguage *string `json:"GenLanguage,omitempty" name:"GenLanguage"`
}

func (r *GenerateApiDocumentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateApiDocumentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogRuleRequest struct {
	*tchttp.BaseRequest

	// 日志规则唯一ID。

	LogRuleId *string `json:"LogRuleId,omitempty" name:"LogRuleId"`
	// 服务唯一ID，表示所需上报日志的服务ID，ALL代表所有服务。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 环境名称，表示所需要上报日志的环境名称，ALL代表所有环境。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 日志集ID。

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 日志主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志规则名称。

	LogRuleName *string `json:"LogRuleName,omitempty" name:"LogRuleName"`
}

func (r *ModifyLogRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountUsageUsagePlanCount struct {

	// 使用计划计数

	UsagePlanCount *int64 `json:"UsagePlanCount,omitempty" name:"UsagePlanCount"`
}

type ApiAppApiInfo struct {

	// 应用名称

	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`
	// 应用ID

	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
	// Api的ID

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// Api名称

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 服务ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 授权绑定时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	AuthorizedTime *string `json:"AuthorizedTime,omitempty" name:"AuthorizedTime"`
	// Api所属地域

	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
	// 授权绑定的环境

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
}

type DescribeLogRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志规则列表。

		LogRuleSet []*LogRule `json:"LogRuleSet,omitempty" name:"LogRuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePluginRequest struct {
	*tchttp.BaseRequest

	// 要查询的插件ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。预留字段，目前不支持过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序依据，默认按照插件创建时间排序。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序方式，默认倒序排序。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribePluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TriggerInfo struct {

	// API网关资源描述六段式

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 云函数版本信息

	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`
}

type DeleteTcbScfApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTcbScfApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTcbScfApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUpstreamResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回修改后的vpc通道信息

		Result *UpstreamInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUpstreamResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePackStatusRequest struct {
	*tchttp.BaseRequest

	// 资源包来源, free：免费包，activity：运营包

	Origin *string `json:"Origin,omitempty" name:"Origin"`
	// 资源包类型，succ_req：调用次数，out_traffic：出流量，不传表示所有类型

	GoodsType []*string `json:"GoodsType,omitempty" name:"GoodsType"`
	// 资源包状态，valid：有效，exhaust：资源耗尽，expire：资源包到期，isolated: 被隔离，后面3种都是失效状态，不传表示所有类型

	GoodsStatus []*string `json:"GoodsStatus,omitempty" name:"GoodsStatus"`
	// 分页参数，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段，默认是：createTime

	Orderby *string `json:"Orderby,omitempty" name:"Orderby"`
	// 排序方式，desc：倒序，asc：正序，默认是 desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeResourcePackStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcePackStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanSecretIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 使用计划绑定的密钥列表。

		Result *UsagePlanBindSecretStatus `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanSecretIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsagePlanSecretIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更换后的密钥详情。

		Result *ApiKey `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用详情。

		Result *ApiAppInfos `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayClsConfigResult struct {

	// 云原生网关实例ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// CLS配置信息列表。

	ClsConfigList []*CloudNativeAPIGatewayClsConfig `json:"ClsConfigList,omitempty" name:"ClsConfigList"`
}

type ModifyUsagePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 使用计划详情。

		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUsagePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 切换版本操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountUsageLogRuleClount struct {

	// 日志使用计数

	LogRuleClount *int64 `json:"LogRuleClount,omitempty" name:"LogRuleClount"`
}

type CreateUpstreamRequest struct {
	*tchttp.BaseRequest

	// VPC通道名字

	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`
	// VPC通道描述

	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`
	// 后端协议，HTTP, HTTPS其中之一

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 后端访问类型L5, IP_PORT

	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`
	// 负载均衡算法目前支持ROUND_ROBIN

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
	// VPC唯一ID

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 请求重试次数，默认3次

	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`
	// 请求到后端的，host头

	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`
	// 后端节点

	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 健康检查配置

	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`
}

func (r *CreateUpstreamRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUpstreamRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MicroServiceReq struct {

	// 微服务集群。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 微服务命名空间。

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 微服务名称。

	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

type TcbScfApi struct {

	// 环境ID。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// SCF方法名称。

	ScfFunctionName *string `json:"ScfFunctionName,omitempty" name:"ScfFunctionName"`
	// SCF方法命名空间。

	ScfFunctionNamespace *string `json:"ScfFunctionNamespace,omitempty" name:"ScfFunctionNamespace"`
	// 自定义子域名。

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 请求Path。

	Path *string `json:"Path,omitempty" name:"Path"`
}

type DescribeUpstreamBindApisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果

		Result *DescribeUpstreamBindApis `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpstreamBindApisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpstreamBindApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRequest struct {
	*tchttp.BaseRequest

	// 计费模式。其中0：按小时计费。

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *InquiryPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachPluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachPluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckServiceConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// True:验证通过；False:验证不通过。

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 验证结果的详细描述信息。

		VerifyMessage *string `json:"VerifyMessage,omitempty" name:"VerifyMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckServiceConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckServiceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApiRsp struct {

	// api id

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// path

	Path *string `json:"Path,omitempty" name:"Path"`
	// method

	Method *string `json:"Method,omitempty" name:"Method"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type DescribeIPStrategysStatusRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 过滤条件。支持StrategyName。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPStrategysStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPStrategysStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAPIDocPasswordRequest struct {
	*tchttp.BaseRequest

	// API文档ID

	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *ResetAPIDocPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAPIDocPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedApiInfo struct {

	// API所在服务ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API所在服务名称。

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// API所在服务描述信息。

	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`
	// API ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// API名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// API描述。

	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`
	// 插件绑定API的环境。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 插件和API绑定时间。

	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`
}

type DeleteIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 待删除的IP策略所属的服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 待删除的IP策略唯一ID。

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DeleteIPStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIPStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExclusiveInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 独享实例列表查询结果

		Result *InstanceInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExclusiveInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExclusiveInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentReleaseHistoryRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 环境名称。

	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceEnvironmentReleaseHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiReqCount struct {

	// 用户的api id

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// api的路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 请求数

	ReqCount *uint64 `json:"ReqCount,omitempty" name:"ReqCount"`
	// 用户可读的服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// api类型。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
}

type LogQuery struct {

	// 检索字段

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 检索值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateAPIDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API文档基本信息

		Result *APIDoc `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAPIDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiEnvironmentApiKeys struct {

	// API绑定密钥数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// API绑定密钥列表。

	AccessKeyIdList []*string `json:"AccessKeyIdList,omitempty" name:"AccessKeyIdList"`
}

type DescribeExclusiveSetRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeExclusiveSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExclusiveSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachPluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解绑操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachPluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAPIDocRequest struct {
	*tchttp.BaseRequest

	// API文档名称

	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`
	// 服务名称

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 环境名称

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 生成文档的API列表

	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
	// API文档ID

	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *ModifyAPIDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAPIDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Api struct {

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API唯一ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// API描述

	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`
	// API PATH。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API METHOD。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 服务创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 服务修改时间。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// API名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// VPCID。

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// VPC唯一ID。

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// API类型。

	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
	// API协议。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 是否买后调试。

	IsDebugAfterCharge *string `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`
	// 授权类型。

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// API业务类型。

	ApiBuniessType *string `json:"ApiBuniessType,omitempty" name:"ApiBuniessType"`
	// 关联授权API唯一ID。

	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`
	// 是否需要鉴权。

	AuthRequired *string `json:"AuthRequired,omitempty" name:"AuthRequired"`
	// oauth配置信息。

	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`
	// oauth2.0API请求，token存放位置。

	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`
	// 授权API关联的业务API列表。

	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds"`
}

type CloudNativeAPIGatewayVpcConfig struct {

	// 私有网络ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type IPStrategyApiStatus struct {

	// 环境绑定API数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 环境绑定API详情。

	ApiIdStatusSet []*IPStrategyApi `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`
}

type DescribeAPIDocDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API文档详细信息

		Result *APIDocInfo `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAPIDocDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAPIDocDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {

	// 自动续费标示

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 预付费到期时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}

type ModifyServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceEnvironmentStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunApiForMarketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的header。

		ReturnHeader *string `json:"ReturnHeader,omitempty" name:"ReturnHeader"`
		// 返回的Body。

		ReturnBody *string `json:"ReturnBody,omitempty" name:"ReturnBody"`
		// 返回码。

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 延迟。

		Delay *int64 `json:"Delay,omitempty" name:"Delay"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunApiForMarketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunApiForMarketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiInfoSummary struct {

	// 插件相关的API总数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 插件相关的API信息。

	ApiSet []*AvailableApiInfo `json:"ApiSet,omitempty" name:"ApiSet"`
}

type ApiKey struct {

	// 创建的 API 密钥 ID 。

	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
	// 创建的 API 密钥 Key。

	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
	// 密钥类型，auto 或者 manual。

	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`
	// 用户自定义密钥名称。

	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
	// 最后一次修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// 密钥状态。0表示禁用，1表示启用。

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ParameterInfo struct {

	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 当前值

	Value *int64 `json:"Value,omitempty" name:"Value"`
	// 默认值

	Default *int64 `json:"Default,omitempty" name:"Default"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 最小

	Minimum *int64 `json:"Minimum,omitempty" name:"Minimum"`
	// 最大

	Maximum *int64 `json:"Maximum,omitempty" name:"Maximum"`
	// 修改时间

	ModifedTime *string `json:"ModifedTime,omitempty" name:"ModifedTime"`
}

type CreateTcbScfApiRequest struct {
	*tchttp.BaseRequest

	// 环境ID。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// SCF方法名称。

	ScfFunctionName *string `json:"ScfFunctionName,omitempty" name:"ScfFunctionName"`
	// SCF方法命名空间。

	ScfFunctionNamespace *string `json:"ScfFunctionNamespace,omitempty" name:"ScfFunctionNamespace"`
}

func (r *CreateTcbScfApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTcbScfApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEIAMApisRequest struct {
	*tchttp.BaseRequest

	// EIAM应用ID。

	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。支持ApiId和ApiName。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEIAMApisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEIAMApisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiIncrementRequest struct {
	*tchttp.BaseRequest

	// 服务ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 接口ID

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// 需要修改的API auth类型(可选择OAUTH-授权API)

	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`
	// oauth接口需要修改的公钥值

	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
	// oauth接口重定向地址

	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

func (r *ModifyApiIncrementRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiIncrementRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PriceInfos struct {

	// 计费单位价格（元）

	UnitPrice *string `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 计费单位周期

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
}

type DescribeApiRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API 接口唯一 ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DescribeApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainsRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceSubDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceSubDomainsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentKeyMonitorUploadRequest struct {
	*tchttp.BaseRequest

	// 服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 需要修改的服务环境列表

	EnvironmentList []*EnvironmentUpload `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
}

func (r *ModifyServiceEnvironmentKeyMonitorUploadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceEnvironmentKeyMonitorUploadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiUsagePlan struct {

	// 服务唯一ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// API 唯一 ID。

	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
	// API 名称。

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// API 路径。

	Path *string `json:"Path,omitempty" name:"Path"`
	// API 方法。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 使用计划的唯一 ID。

	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
	// 使用计划的名称。

	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`
	// 使用计划的描述。

	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`
	// 使用计划绑定的服务环境。

	Environment *string `json:"Environment,omitempty" name:"Environment"`
	// 已经使用的配额。

	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`
	// 请求配额总量，-1表示没有限制。

	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`
	// 请求 QPS 上限，-1 表示没有限制。

	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
	// 使用计划创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 使用计划最后修改时间。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// 服务名称。

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type ModifyPluginRequest struct {
	*tchttp.BaseRequest

	// 要修改的插件ID。

	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
	// 要修改的API网关插件名称。最长50个字符，支持 a-z,A-Z,0-9,_, 必须字母开头，字母或者数字结尾。

	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`
	// 要修改的插件描述，限定200字以内。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 要修改的插件定义语句，支持json。

	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyPluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedApiSummary struct {

	// 插件绑定的API数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 插件绑定的API信息。

	AttachedApis []*AttachedApiInfo `json:"AttachedApis,omitempty" name:"AttachedApis"`
}

type UpstreamHealthCheckerReqHeaders struct {

	// 请求头名称

	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
	// 请求头值

	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`
}

type DescribeApiUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiUsagePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiUsagePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 调试API响应。

		Result *RunApiReturn `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountSettings struct {

	// 用户配额

	AccountLimit *Limit `json:"AccountLimit,omitempty" name:"AccountLimit"`
	// 使用量

	AccountUsage *Usage `json:"AccountUsage,omitempty" name:"AccountUsage"`
}

type ResourcePackResult struct {

	// 资源包数量

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 资源包具体信息

	ResourcePackSet []*ResourcePackSetList `json:"ResourcePackSet,omitempty" name:"ResourcePackSet"`
}

type DescribeUpstreamsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUpstreamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpstreamsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 启动密钥操作是否成功。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsagePlansStatus struct {

	// 符合条件的使用计划数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 使用计划列表。

	UsagePlanStatusSet []*UsagePlanStatusInfo `json:"UsagePlanStatusSet,omitempty" name:"UsagePlanStatusSet"`
}

type MonitorTopsRequest struct {
	*tchttp.BaseRequest

	// 环境标识

	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *MonitorTopsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MonitorTopsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源利用率趋势图。

		ResourceUsageTrend []*float64 `json:"ResourceUsageTrend,omitempty" name:"ResourceUsageTrend"`
		// 用户的独占集群数量。

		ExclusiveLdCount *uint64 `json:"ExclusiveLdCount,omitempty" name:"ExclusiveLdCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MonitorTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MonitorTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePluginsRequest struct {
	*tchttp.BaseRequest

	// 批量删除的API网关插件的ID列表。

	PluginIds []*string `json:"PluginIds,omitempty" name:"PluginIds"`
}

func (r *DeletePluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePluginsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAPIDocsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。预留字段，目前不支持过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序依据，默认按照插件创建时间排序。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序方式，默认倒序排序。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeAPIDocsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAPIDocsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExclusiveInstanceRequest struct {
	*tchttp.BaseRequest

	// 独享实例唯一id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 独享实例name

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 独享实例描述

	InstanceDescription *string `json:"InstanceDescription,omitempty" name:"InstanceDescription"`
	// 独享实例参数配置

	Parameters []*InstanceParameterInput `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *ModifyExclusiveInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExclusiveInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceRequest struct {
	*tchttp.BaseRequest

	// 待修改服务的唯一 Id。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 修改后的服务名称。

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 修改后的服务描述。

	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`
	// 修改后的服务前端请求类型，如 http、https和 http&https。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 网络类型列表，用于指定支持的访问类型，INNER为内网访问，OUTER为外网访问。默认为OUTER。

	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiAppInfo struct {

	// 应用名称

	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`
	// 应用ID

	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
	// 应用SECRET

	ApiAppSecret *string `json:"ApiAppSecret,omitempty" name:"ApiAppSecret"`
	// 应用描述

	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
	// 创建时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 修改时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// 应用KEY

	ApiAppKey *string `json:"ApiAppKey,omitempty" name:"ApiAppKey"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}
