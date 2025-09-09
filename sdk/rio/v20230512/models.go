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

package v20230512

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type CreateConnectorRequest struct {
	*tchttp.BaseRequest

	// 认证 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 认证名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 认证类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 认证配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// 认证状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 认证标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// AgencyID

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
	// Usage

	Usage *string `json:"Usage,omitempty" name:"Usage"`
}

func (r *CreateConnectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConnectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTop10Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务列表，参考附录 《对象说明》

		Services []*ServiceVO `json:"Services,omitempty" name:"Services"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTop10Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceTop10Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedServiceBatchRequest struct {
	*tchttp.BaseRequest

	// 系统 ID 列表

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
	// 系统 ID 列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *ApprovedServiceBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedServiceBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticsAgencyRequest struct {
	*tchttp.BaseRequest

	// 关键字机构名称等

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *StatisticsAgencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StatisticsAgencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckVendorRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标识

	Badge *string `json:"Badge,omitempty" name:"Badge"`
}

func (r *CheckVendorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckVendorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSitesRequest struct {
	*tchttp.BaseRequest

	// 应用ID数组

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 状态列表。 0 草稿 1 待审批 2 已审批 5 已禁用 6 已冻结

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 归档状态（0：全部，3：已拒绝，7：已删除）默认取0

	ArchStatus *int64 `json:"ArchStatus,omitempty" name:"ArchStatus"`
	// 限定分页条目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序搜索

	Sort *DescribeSiteSortDTO `json:"Sort,omitempty" name:"Sort"`
	// Keyword

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeSitesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSitesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAppUniqueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *CheckUniqueResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAppUniqueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAppUniqueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConnectorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录 《对象说明》

		Items []*Connector `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConnectorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTlsCertificatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *string `json:"Total,omitempty" name:"Total"`
		// 当前页面的证书数组

		Items []*TlsCertificateVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTlsCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTlsCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpsertSqlTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpsertSqlTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertSqlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSqlTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSqlTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSqlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroRegisterCentersRequest struct {
	*tchttp.BaseRequest

	// ID

	ID []*string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteMicroRegisterCentersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroRegisterCentersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroTenantsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*MicroTenant `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroTenantsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroTenantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProfileForRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户列表 参考附录 《对象说明》

		Items []*UserProfileWithExist `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProfileForRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProfileForRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *SiteVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VendorOfPagedApps struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateTSFGatewayInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTSFGatewayInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFGatewayInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCheckAccountRequest struct {
	*tchttp.BaseRequest

	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
}

func (r *DescribeCheckAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFSyncTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录 《对象说明》

		Items []*SyncType `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTSFSyncTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFSyncTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneAreasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// ZoneArea

		Items *ZoneArea `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneAreasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVendorMemberRequest struct {
	*tchttp.BaseRequest

	// VendorID

	VendorID *string `json:"VendorID,omitempty" name:"VendorID"`
}

func (r *ExportVendorMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVendorMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProfileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Items

		Items []*UserProfileWithExist `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProfileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationLog struct {

	// 日志id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 账号id

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
	// 账号名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 模块名称

	OpModule *string `json:"OpModule,omitempty" name:"OpModule"`
	// UserGroup

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// 模块对象

	OpObject *string `json:"OpObject,omitempty" name:"OpObject"`
	// 模块动作

	OpAction *string `json:"OpAction,omitempty" name:"OpAction"`
	// OpData

	OpData *string `json:"OpData,omitempty" name:"OpData"`
	// 时间

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// ua

	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
	// 访问url

	URL *string `json:"URL,omitempty" name:"URL"`
	// 返回结果

	Result *string `json:"Result,omitempty" name:"Result"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type TSFNamespace struct {

	// 命名空间标识

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 命名空间描述

	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`
}

type DescribeAgenciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机构总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 机构列表

		Items []*AgencyVo `json:"Items,omitempty" name:"Items"`
		// 机构成员总数

		MemberCount *int64 `json:"MemberCount,omitempty" name:"MemberCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgenciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgenciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceArchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用ID

		PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
		// 服务ID

		SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
		// 服务名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 描述信息

		Description *string `json:"Description,omitempty" name:"Description"`
		// 公开路径

		PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
		// 可见性，公开还是私有

		IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
		// 仅适用于里约自有服务初始化时用的，页面不用管这个字段

		Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
		// 负载方式：false随机负载，true会话保持

		KeepSession *string `json:"KeepSession,omitempty" name:"KeepSession"`
		// 目标服务类型

		TargetServiceType *string `json:"TargetServiceType,omitempty" name:"TargetServiceType"`
		// 全局设置

		ServiceAllSet *ServiceAllSet `json:"ServiceAllSet,omitempty" name:"ServiceAllSet"`
		// 分区信息，数组

		Regions *RegionsOfService `json:"Regions,omitempty" name:"Regions"`
		// 高级配置

		AdvanceSetting *AdvanceSettingOfService `json:"AdvanceSetting,omitempty" name:"AdvanceSetting"`
		// 标签

		LabelIDs []*string `json:"LabelIDs,omitempty" name:"LabelIDs"`
		// 服务目录，也叫服务分类

		CategoryIDs []*string `json:"CategoryIDs,omitempty" name:"CategoryIDs"`
		// 服务类别目录

		CategoryIds []*string `json:"CategoryIds,omitempty" name:"CategoryIds"`
		// 接口地址

		DocUrl *string `json:"DocUrl,omitempty" name:"DocUrl"`
		// 接口文档ID

		DocIds []*string `json:"DocIds,omitempty" name:"DocIds"`
		// 应用

		App *Summary `json:"App,omitempty" name:"App"`
		// 系统

		Sys *Summary `json:"Sys,omitempty" name:"Sys"`
		// 机构

		Agency *Summary `json:"Agency,omitempty" name:"Agency"`
		// 状态（草稿0，待审批1，已审批2，禁用5，冻结6）

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 被订阅的应用数

		AuthAppNum *int64 `json:"AuthAppNum,omitempty" name:"AuthAppNum"`
		// 创建者

		CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
		// CreateTime

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// LastUpdateTime

		LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
		// 审批者

		ApprovalUser *CreateUser `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
		// 归档ID

		ServiceArchID *string `json:"ServiceArchID,omitempty" name:"ServiceArchID"`
		// ApprovalTime

		ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceArchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceArchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProfileForRoleRequest struct {
	*tchttp.BaseRequest

	// 机构ID列表. 第一个传角色 code 第二个传 agencyID。第一个 角色code 必填, 第二个 agencyID 全局角色时必填 ""

	AgencyIDs []*string `json:"AgencyIDs,omitempty" name:"AgencyIDs"`
	// 供应商ID列表, 无供应商填写 []

	VendorIDs []*string `json:"VendorIDs,omitempty" name:"VendorIDs"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段， 无查询填写 ""

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeProfileForRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProfileForRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTSFGatewayServiceRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 网关实例 ID

	GatewayInstanceID *string `json:"GatewayInstanceID,omitempty" name:"GatewayInstanceID"`
	// api 分组 ID

	GroupID *string `json:"GroupID,omitempty" name:"GroupID"`
	// api 分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// api 分组路径前缀

	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`
	// 参考附录 《对象说明》

	Apis []*MicroApi `json:"Apis,omitempty" name:"Apis"`
}

func (r *CreateTSFGatewayServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFGatewayServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountAndProfileListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录 《对象说明》

		Items []*AccountAndProfileDTO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountAndProfileListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountAndProfileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectedSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroTenantRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeMicroTenantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroTenantRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileDownloadResponse struct {

	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 内容的base64数据

	Content *string `json:"Content,omitempty" name:"Content"`
}

type AddVendorMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户帐号

		Account *string `json:"Account,omitempty" name:"Account"`
		// 初始密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddVendorMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVendorMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest

	// 0：表示继续保存草稿，1：表示之前的草稿现在提交审批，2：已审批通过的服务编辑（包含已审批之后的其他状态，）

	DraftFlag *int64 `json:"DraftFlag,omitempty" name:"DraftFlag"`
	// 分区信息，数组

	Regions []*RegionsOfService `json:"Regions,omitempty" name:"Regions"`
	// 接口地址

	DocUrl *string `json:"DocUrl,omitempty" name:"DocUrl"`
	// 接口文档ID

	DocIDs []*string `json:"DocIDs,omitempty" name:"DocIDs"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 服务ID

	SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
	// 服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 公开路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 负载方式：false随机负载，true会话保持

	KeepSession *bool `json:"KeepSession,omitempty" name:"KeepSession"`
	// 可见性，公开还是私有

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 仅适用于里约自有服务初始化时用的

	Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
	// 目标服务类型

	TargetServiceType *string `json:"TargetServiceType,omitempty" name:"TargetServiceType"`
	// 全局设置

	ServiceAllSet *ServiceAllSet `json:"ServiceAllSet,omitempty" name:"ServiceAllSet"`
	// 高级设置

	AdvanceSetting *AdvanceSettingOfService `json:"AdvanceSetting,omitempty" name:"AdvanceSetting"`
	// 标签

	LabelIDs []*string `json:"LabelIDs,omitempty" name:"LabelIDs"`
	// 服务目录

	CategoryIDs []*string `json:"CategoryIDs,omitempty" name:"CategoryIDs"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
}

func (r *CreateServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCurrentUserRequest struct {
	*tchttp.BaseRequest
}

func (r *GetCurrentUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCurrentUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResetTopMenuResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResetTopMenuResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResetTopMenuResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemApproveLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批列表，参考附录 《对象说明》

		Items []*SystemApproveLogVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemApproveLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemApproveLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Register struct {

	// 注册中心ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 注册中心名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 注册中心类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 后端服务协议

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 注册中心地址

	Address []*string `json:"Address,omitempty" name:"Address"`
	// 注册中心自定义配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type CreateMicroProtoGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMicroProtoGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMicroProtoGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteSiteRequest struct {
	*tchttp.BaseRequest

	// 站点ID

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
}

func (r *BatchDeleteSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Subscription struct {

	// 订阅ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// SubscriberApp

	SubscriberApp *string `json:"SubscriberApp,omitempty" name:"SubscriberApp"`
	// PublisherService

	PublisherService *string `json:"PublisherService,omitempty" name:"PublisherService"`
	// 发起调用区域

	AccessRegions []*string `json:"AccessRegions,omitempty" name:"AccessRegions"`
	// 运行被调用区域

	DeployRegions *int64 `json:"DeployRegions,omitempty" name:"DeployRegions"`
	// CreateUser

	CreateUser *string `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ApprovalUser

	ApprovalUser *string `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// ApprovalTime

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// 鉴权方式

	SignType []*string `json:"SignType,omitempty" name:"SignType"`
	// 签名方法

	Signature *string `json:"Signature,omitempty" name:"Signature"`
	// 高级规则

	AdvancedRules *string `json:"AdvancedRules,omitempty" name:"AdvancedRules"`
	// ip白名单

	IpWhiteList *int64 `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 0:待审批,1:已审批，2：已拒绝，3：已取消，5：已删除，6：已撤销

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CurRenewStatus

	CurRenewStatus *int64 `json:"CurRenewStatus,omitempty" name:"CurRenewStatus"`
	// Builtin

	Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
	// 过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 限流设置类型，0：默认服务最大限流，1：用户自定义

	InvokeLimitType *int64 `json:"InvokeLimitType,omitempty" name:"InvokeLimitType"`
	// 限流大小

	InvokeLimit *int64 `json:"InvokeLimit,omitempty" name:"InvokeLimit"`
	// 限流统计时间间隔

	InvokeLimitInterval *int64 `json:"InvokeLimitInterval,omitempty" name:"InvokeLimitInterval"`
}

type DeleteZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceOfSubAppRequest struct {
	*tchttp.BaseRequest

	// 订阅应用ID

	SubscriberID *string `json:"SubscriberID,omitempty" name:"SubscriberID"`
	// 系统ID列表

	SysIDs []*string `json:"SysIDs,omitempty" name:"SysIDs"`
	// 应用ID列表

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 关键词搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 分页查询数据起始点

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页查询数据数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeServiceOfSubAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceOfSubAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedServiceBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败数量

		FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
		// 成功数量

		SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectedServiceBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedServiceBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SysOfSiteVO struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeMicroRegisterCenterRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeMicroRegisterCenterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroRegisterCenterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceApprovesRequest struct {
	*tchttp.BaseRequest

	// 应用id列表

	AppIDs []*string `json:"AppIDs,omitempty" name:"AppIDs"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// PaasIDs

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
}

func (r *DescribeServiceApprovesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceApprovesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFGatewayInstancesRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 网关实例模糊搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeTSFGatewayInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFGatewayInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemApproveLogVO struct {

	// 系统id

	SysId *string `json:"SysId,omitempty" name:"SysId"`
	// 系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属机构

	Org *OrgOfSystem `json:"Org,omitempty" name:"Org"`
	// 负责人

	Operators []*CreateUser `json:"Operators,omitempty" name:"Operators"`
	// 申请人

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// 申请时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 系统描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 供应商

	Vendors []*Summary `json:"Vendors,omitempty" name:"Vendors"`
	// 审批记录id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 审批人

	ApprovalUser *string `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// 审批时间

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 驳回原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type CreateTSFGatewayApiGroupRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 网关实例 ID

	GatewayInstanceID *string `json:"GatewayInstanceID,omitempty" name:"GatewayInstanceID"`
	// 参考附录 《对象说明》

	ApiGroups []*TSFApiGroup `json:"ApiGroups,omitempty" name:"ApiGroups"`
}

func (r *CreateTSFGatewayApiGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFGatewayApiGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteRequestLogPageRequest struct {
	*tchttp.BaseRequest

	// 限定分页条目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 结束时间（时间戳，秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间（时间戳，秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 站点ID

	SiteID *string `json:"SiteID,omitempty" name:"SiteID"`
	// 应用ID数组

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 站点请求路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 响应状态码

	StatusCode *string `json:"StatusCode,omitempty" name:"StatusCode"`
	// 源IP

	OriginIP *string `json:"OriginIP,omitempty" name:"OriginIP"`
	// 访问者

	UserID *string `json:"UserID,omitempty" name:"UserID"`
	// 排序搜索

	Sort *DescribeRequestLogSortDTO `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeSiteRequestLogPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteRequestLogPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmartGateRaidRoutingRequest struct {
	*tchttp.BaseRequest

	// 节点阵列路由 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeSmartGateRaidRoutingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSmartGateRaidRoutingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectRenewSubscriptionBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectRenewSubscriptionBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectRenewSubscriptionBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRequest struct {

	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 手机区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
}

type DescribeSitesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *string `json:"Total,omitempty" name:"Total"`
		// 当前页面的站点数组, 参考附录 《对象说明》

		Items []*SiteVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSitesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemAndAppRequest struct {
	*tchttp.BaseRequest

	// 应用关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 状态列表

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSystemAndAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemAndAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedSiteBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败数量

		FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
		// 成功数量

		SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApprovedSiteBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSiteBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsOfPubAppRequest struct {
	*tchttp.BaseRequest

	// 被订阅应用ID

	PublisherAppID *string `json:"PublisherAppID,omitempty" name:"PublisherAppID"`
	// 服务名称或标识

	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
	// 分页查询数据起始点

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 状态

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 分页查询数据数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubscriptionsOfPubAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionsOfPubAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteMicroSyncJobRequest struct {
	*tchttp.BaseRequest

	// 同步记录 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *ExecuteMicroSyncJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteMicroSyncJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroProtoFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*MicroProtoFile `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroProtoFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroProtoFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFGatewayApiGroupsRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 网关实例 ID

	GatewayInstanceID *string `json:"GatewayInstanceID,omitempty" name:"GatewayInstanceID"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 网关分组模糊搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeTSFGatewayApiGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFGatewayApiGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 续期时间，时间戳，单位秒，默认0，永久有效

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *RenewSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DocNodeVO struct {

	// 文档名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件目录

	Dirs []*DocNodeVO `json:"Dirs,omitempty" name:"Dirs"`
	// 文件结构

	Files []*FileVO `json:"Files,omitempty" name:"Files"`
}

type CreateZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTipsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items *Tips `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTipsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMenuResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMenuResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMenuResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectedSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProfileInfo struct {

	// 用户ID 更新时送

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 证件号码

	CID *string `json:"CID,omitempty" name:"CID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 用户显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 性别

	Gender *string `json:"Gender,omitempty" name:"Gender"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
}

type DeleteAccountAndProfileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountAndProfileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountAndProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAdvanceTemplateRequest struct {
	*tchttp.BaseRequest

	// 站点：site，服务：service

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAdvanceTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAdvanceTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportServiceRequest struct {
	*tchttp.BaseRequest

	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 内容的base64数据

	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *ImportServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgencyMember struct {

	// 用户账号id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 证件号码

	CID *string `json:"CID,omitempty" name:"CID"`
	// 帐号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 姓名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 性别

	Gender *string `json:"Gender,omitempty" name:"Gender"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 绑定时间

	BindTime *uint64 `json:"BindTime,omitempty" name:"BindTime"`
	// 机构ID

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

type Link struct {

	// 链接类型 1 跳转外链 2 初始化数据的链接类型 3 嵌入页面

	LinkType *int64 `json:"LinkType,omitempty" name:"LinkType"`
	// url

	LinkURL *string `json:"LinkURL,omitempty" name:"LinkURL"`
}

type ApprovedSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApprovedSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data []*DescribeGrafanaAlertRuleResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *AppOfSystemAppInfoVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLabelRequest struct {
	*tchttp.BaseRequest

	// 标签ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSubscriptionBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败数量

		FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
		// 成功数量

		SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectedSubscriptionBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSubscriptionBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MicroApi struct {

	// 机构标识

	Org *string `json:"Org,omitempty" name:"Org"`
	// 系统标识

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// 系统名称

	SysName *string `json:"SysName,omitempty" name:"SysName"`
	// 系统描述

	SysDesc *string `json:"SysDesc,omitempty" name:"SysDesc"`
	// 应用标识

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 应用名称

	PaasName *string `json:"PaasName,omitempty" name:"PaasName"`
	// 应用描述

	PaasDesc *string `json:"PaasDesc,omitempty" name:"PaasDesc"`
	// service 描述

	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`
	// 目标主机

	Host []*string `json:"Host,omitempty" name:"Host"`
	// 方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 原始路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 导入路径

	RioPath *string `json:"RioPath,omitempty" name:"RioPath"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type Receivers struct {

	// Channel

	Channel *string `json:"Channel,omitempty" name:"Channel"`
	// qyapi.weixin.qq.com/cgi-bin/webhook/send?key=8c1e844b-7bb1-4d9b-91ba-75d83a47db8c"

	Address *string `json:"Address,omitempty" name:"Address"`
}

type ApprovedSiteRequest struct {
	*tchttp.BaseRequest

	// ID 列表

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *ApprovedSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDownloadLoginLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDownloadLoginLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDownloadLoginLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroTenantRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteMicroTenantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroTenantRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest

	// 服务标识

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResetPasswordRequest struct {
	*tchttp.BaseRequest

	// 帐号

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
}

func (r *ModifyResetPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResetPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlTemplateRequest struct {
	*tchttp.BaseRequest

	// ServiceID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
}

func (r *DescribeSqlTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSqlTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTSFNamespaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTSFNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录 《对象说明》

		AtlasCaptcha *AtlasCaptcha `json:"AtlasCaptcha,omitempty" name:"AtlasCaptcha"`
		// 参考附录 《对象说明》

		AtlasLockSetting *AtlasLockSetting `json:"AtlasLockSetting,omitempty" name:"AtlasLockSetting"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcsRolePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAcsRolePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAcsRolePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLabelRequest struct {
	*tchttp.BaseRequest

	// 标签id

	LabelID *string `json:"LabelID,omitempty" name:"LabelID"`
	// 标签名称

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
}

func (r *ModifyLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAgencyRequest struct {
	*tchttp.BaseRequest

	// 机构ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 机构名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 机构标识

	Badge *string `json:"Badge,omitempty" name:"Badge"`
	// 机构描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 机构负责人

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
}

func (r *UpdateAgencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAgencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateVendorMemberRequest struct {
	*tchttp.BaseRequest

	// 供应商ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 用户id

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 性别

	Gender *string `json:"Gender,omitempty" name:"Gender"`
}

func (r *UpdateVendorMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateVendorMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cert struct {

	// 文件名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 文件上传时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateConnectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConnectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConnectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadAppDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadAppDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadAppDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAcsRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAcsRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAcsRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionApprovesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统审批列表，参考附录 《对象说明》

		Items []*SubscriptionApproveVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionApprovesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionApprovesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordByOldRequest struct {
	*tchttp.BaseRequest

	// 应用ID

	AppID *string `json:"AppID,omitempty" name:"AppID"`
	// 账户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 原密码

	CurrentPassword *string `json:"CurrentPassword,omitempty" name:"CurrentPassword"`
	// 新密码

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
	// 时间戳

	Nonce *string `json:"Nonce,omitempty" name:"Nonce"`
}

func (r *ModifyPasswordByOldRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordByOldRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigOfDescribeAdvanceTemplate struct {

	// Function

	Function *string `json:"Function,omitempty" name:"Function"`
	// CacheByAppId

	CacheByAppId *bool `json:"CacheByAppId,omitempty" name:"CacheByAppId"`
	// CacheByParameters

	CacheByParameters []*string `json:"CacheByParameters,omitempty" name:"CacheByParameters"`
	// RequestMethod

	RequestMethod []*string `json:"RequestMethod,omitempty" name:"RequestMethod"`
	// ResponseCode

	ResponseCode []*string `json:"ResponseCode,omitempty" name:"ResponseCode"`
	// ClientCacheControl

	ClientCacheControl *string `json:"ClientCacheControl,omitempty" name:"ClientCacheControl"`
	// ExpireTimeSec

	ExpireTimeSec *int64 `json:"ExpireTimeSec,omitempty" name:"ExpireTimeSec"`
	// CacheMaxSize

	CacheMaxSize *int64 `json:"CacheMaxSize,omitempty" name:"CacheMaxSize"`
}

type SmartGatewayGroup struct {

	// SggID

	SggID *string `json:"SggID,omitempty" name:"SggID"`
	// SggName

	SggName *string `json:"SggName,omitempty" name:"SggName"`
	// 新增编辑修改不填

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// BelongZoneID

	BelongZoneID *string `json:"BelongZoneID,omitempty" name:"BelongZoneID"`
	// 新增编辑修改不填

	BelongZoneName *string `json:"BelongZoneName,omitempty" name:"BelongZoneName"`
	// BelongAreaID

	BelongAreaID *string `json:"BelongAreaID,omitempty" name:"BelongAreaID"`
	// 新增编辑修改不填

	BelongAreaName *string `json:"BelongAreaName,omitempty" name:"BelongAreaName"`
	// Domain

	Domain []*string `json:"Domain,omitempty" name:"Domain"`
	// 查找客户端真实ip时过滤内部ip

	IPs []*string `json:"IPs,omitempty" name:"IPs"`
	// RoleID

	RoleID *string `json:"RoleID,omitempty" name:"RoleID"`
	// 新增编辑修改不填

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 127.0.0.1:8006",

	SgAdminBaseUrl *string `json:"SgAdminBaseUrl,omitempty" name:"SgAdminBaseUrl"`
	// 新增编辑修改不填

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 新增编辑修改不填

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type ModifySiteRequest struct {
	*tchttp.BaseRequest

	// 0表示草稿，1表示创建

	DraftFlag *int64 `json:"DraftFlag,omitempty" name:"DraftFlag"`
	// 站点ID

	SiteID *string `json:"SiteID,omitempty" name:"SiteID"`
	// paasID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 站点名称

	SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
	// 公开路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 负载方式：false，随机负载，true：会话保持

	KeepSession *bool `json:"KeepSession,omitempty" name:"KeepSession"`
	// 分区配置信息

	Regions []*RegionsOfSite `json:"Regions,omitempty" name:"Regions"`
	// AdvanceSetting

	AdvanceSetting *AdvanceSetting `json:"AdvanceSetting,omitempty" name:"AdvanceSetting"`
}

func (r *ModifySiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMicroProtoFileRequest struct {
	*tchttp.BaseRequest

	// 文件组 ID

	ProtoGroupId *string `json:"ProtoGroupId,omitempty" name:"ProtoGroupId"`
	// 文件名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件内容

	Text *string `json:"Text,omitempty" name:"Text"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

func (r *CreateMicroProtoFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMicroProtoFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvanceSettingOfService struct {

	// 是否开启健康检查

	HealthCheckStatus *bool `json:"HealthCheckStatus,omitempty" name:"HealthCheckStatus"`
	// 健康检查地址

	HealthCheckPath *string `json:"HealthCheckPath,omitempty" name:"HealthCheckPath"`
	// 健康检查超时

	HealthCheckTimeout *int64 `json:"HealthCheckTimeout,omitempty" name:"HealthCheckTimeout"`
	// 健康检查码

	ValidHealthCheckStatusCode []*int64 `json:"ValidHealthCheckStatusCode,omitempty" name:"ValidHealthCheckStatusCode"`
	// 是否开启格式转换

	SourceTypeStatus *bool `json:"SourceTypeStatus,omitempty" name:"SourceTypeStatus"`
	// 返回资源类型

	ResSourceType *string `json:"ResSourceType,omitempty" name:"ResSourceType"`
	// 返回目标类型

	ResTargetType *string `json:"ResTargetType,omitempty" name:"ResTargetType"`
	// 请求资源类型

	ReqSourceType *string `json:"ReqSourceType,omitempty" name:"ReqSourceType"`
	// 请求目标类型

	ReqTargetType *string `json:"ReqTargetType,omitempty" name:"ReqTargetType"`
	// 是否自定义请求响应

	FilterStatus *bool `json:"FilterStatus,omitempty" name:"FilterStatus"`
	// 目标主机头

	TargetHostHeader *string `json:"TargetHostHeader,omitempty" name:"TargetHostHeader"`
	// 请求头

	RequestHeadersFilter *string `json:"RequestHeadersFilter,omitempty" name:"RequestHeadersFilter"`
	// 请求内容,

	RequestBodyFilter *string `json:"RequestBodyFilter,omitempty" name:"RequestBodyFilter"`
	// 请求参数

	RequestParameterProcesser *string `json:"RequestParameterProcesser,omitempty" name:"RequestParameterProcesser"`
	// 响应头

	ResponseHeadersFilter *string `json:"ResponseHeadersFilter,omitempty" name:"ResponseHeadersFilter"`
	// 响应内容

	ResponseBodyFilter *string `json:"ResponseBodyFilter,omitempty" name:"ResponseBodyFilter"`
	// 响应参数

	ResponseParameterProcesser *string `json:"ResponseParameterProcesser,omitempty" name:"ResponseParameterProcesser"`
	// 是否开启IP白名单，注意黑白名单之间互斥（前端不好判断的话没事，直接提交，后端会校验的）

	IpWhiteStatus *bool `json:"IpWhiteStatus,omitempty" name:"IpWhiteStatus"`
	// IP白名单，

	IpWhiteList []*IpWhiteList `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// 是否开启IP黑名单，注意黑白名单之间互斥

	IpBlackStatus *bool `json:"IpBlackStatus,omitempty" name:"IpBlackStatus"`
	// IP黑名单

	IpBlackList []*IpWhiteList `json:"IpBlackList,omitempty" name:"IpBlackList"`
	// 是否开启API水印

	DataWaterStatus *bool `json:"DataWaterStatus,omitempty" name:"DataWaterStatus"`
	// 数据水印添加的字段，如果开启了API水印，该项为必填项,逗号分隔

	DataWaterFields *string `json:"DataWaterFields,omitempty" name:"DataWaterFields"`
	// TemplateAdvanceSetList

	TemplateAdvanceSetList []*TemplateAdvanceSetList `json:"TemplateAdvanceSetList,omitempty" name:"TemplateAdvanceSetList"`
	// OpenMessageLogStatus

	OpenMessageLogStatus *bool `json:"OpenMessageLogStatus,omitempty" name:"OpenMessageLogStatus"`
}

type AddMemberToAcsRoleRequest struct {
	*tchttp.BaseRequest

	// 0:普通角色成员 10:机构负责人 11:系统负责人 12: 应用负责人

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 空：角色关系 非空：负责人关系

	RelationshipBelongTo *string `json:"RelationshipBelongTo,omitempty" name:"RelationshipBelongTo"`
	// 编码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 批量添加用户

	AccountIDs []*string `json:"AccountIDs,omitempty" name:"AccountIDs"`
}

func (r *AddMemberToAcsRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMemberToAcsRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticsAgencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机构数量

		AgencyCount *int64 `json:"AgencyCount,omitempty" name:"AgencyCount"`
		// 成员数量

		MemberCount *int64 `json:"MemberCount,omitempty" name:"MemberCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StatisticsAgencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StatisticsAgencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTSFNamespaceRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 命名空间 ID

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 微服务实例列表 参考附录 《对象说明》

	Microservices []*MicroserviceData `json:"Microservices,omitempty" name:"Microservices"`
}

func (r *CreateTSFNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgencyConfigRequest struct {
	*tchttp.BaseRequest

	// 机构id

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
}

func (r *DeleteAgencyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgencyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*AreaZone `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAreaZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 偏移量

		Offset *int64 `json:"Offset,omitempty" name:"Offset"`
		// 条数

		Limit *int64 `json:"Limit,omitempty" name:"Limit"`
		// 开始时间

		StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
		// 操作日志列表

		Logs []*OperationLog `json:"Logs,omitempty" name:"Logs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceApproveLogVO struct {

	// 所属应用

	App *Summary `json:"App,omitempty" name:"App"`
	// 申请时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 申请人

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// API名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 访问路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// API ID

	SvcId *string `json:"SvcId,omitempty" name:"SvcId"`
	// 审批记录id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 审批人

	ApprovalUser *string `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// 审批时间

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 驳回原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type AdvanceTemplateBtnFunctionRequest struct {
	*tchttp.BaseRequest

	// 站点：site，服务：service

	Type *string `json:"Type,omitempty" name:"Type"`
	// 服务或者站点的ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 模板ID

	TemplateID *string `json:"TemplateID,omitempty" name:"TemplateID"`
	// 高级设置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 按钮的提交功能

	Submit *string `json:"Submit,omitempty" name:"Submit"`
}

func (r *AdvanceTemplateBtnFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AdvanceTemplateBtnFunctionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 授权中心节点名称

		AuthCenterName *string `json:"AuthCenterName,omitempty" name:"AuthCenterName"`
		// 参考附录 《对象说明》

		AuthNodesInfo []*AuthNodeInfo `json:"AuthNodesInfo,omitempty" name:"AuthNodesInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuthInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroRegisterCenterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *Register `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroRegisterCenterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroRegisterCenterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindMemberRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserIDs []*string `json:"UserIDs,omitempty" name:"UserIDs"`
	// 机构id

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

func (r *BindMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataSourceRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmartGatewayGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*SmartGatewayGroup `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmartGatewayGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSmartGatewayGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MicroserviceData struct {

	// MicroserviceID

	MicroserviceID *string `json:"MicroserviceID,omitempty" name:"MicroserviceID"`
	// PaasID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
}

type Tips struct {

	// 提示语码号 tip_code应为>=100000整型字符串且在有效范围内

	TipCode *string `json:"TipCode,omitempty" name:"TipCode"`
	// 当前提示语

	CurrentTip *string `json:"CurrentTip,omitempty" name:"CurrentTip"`
	// 说明

	Describe *string `json:"Describe,omitempty" name:"Describe"`
	// 默认提示语

	DefaultTip *string `json:"DefaultTip,omitempty" name:"DefaultTip"`
	// 路径

	URLPath *string `json:"URLPath,omitempty" name:"URLPath"`
	// 提示语类型 1: 登录模块 tip

	TipType *int64 `json:"TipType,omitempty" name:"TipType"`
	// 提示语子类型 10: 登录模块--帐号密码登录子模块 11: 登录模块--短信验证码登录 12: 登录模块--双因子登录

	TipSubType *int64 `json:"TipSubType,omitempty" name:"TipSubType"`
	// 是否为手机号格式

	PhoneRelated *bool `json:"PhoneRelated,omitempty" name:"PhoneRelated"`
	// 手机号脱敏方式  "0" 表示不显示手机号 "1" 表示脱敏保留前三位，后四位：176****7096 "2" 表示脱敏前八位：********7096

	PhoneMaskFormat *string `json:"PhoneMaskFormat,omitempty" name:"PhoneMaskFormat"`
}

type DescribeRolesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricData struct {

	// Timestamp

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateDataSourceRequest struct {
	*tchttp.BaseRequest

	// 数据源所属应用

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 数据源 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 数据源名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据源类型, 枚举值参考 DescribeDataSourceTypes 接口返回

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据源连接串

	Dsn *string `json:"Dsn,omitempty" name:"Dsn"`
	// 最大返回记录数

	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
}

func (r *CreateDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConnectorsRequest struct {
	*tchttp.BaseRequest

	// 认证 ID

	ID []*string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteConnectorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConnectorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站点数量

		SiteNum *int64 `json:"SiteNum,omitempty" name:"SiteNum"`
		// 服务数量

		ServiceNum *int64 `json:"ServiceNum,omitempty" name:"ServiceNum"`
		// AreaID

		AreaID *string `json:"AreaID,omitempty" name:"AreaID"`
		// AreaName

		AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
		// DefaultArea

		DefaultArea *bool `json:"DefaultArea,omitempty" name:"DefaultArea"`
		// Status

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// CreateTime

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// LastUpdateTime

		LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequestLogPageRequest struct {
	*tchttp.BaseRequest

	// 限定分页条目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 结束时间（时间戳，秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间（时间戳，秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 服务ID

	SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
	// 应用ID数组

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 服务请求路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 响应状态码

	StatusCode *string `json:"StatusCode,omitempty" name:"StatusCode"`
	// 源IP

	OriginIP *string `json:"OriginIP,omitempty" name:"OriginIP"`
	// 排序搜索

	Sort *DescribeRequestLogSortDTO `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeServiceRequestLogPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequestLogPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVendorListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录《对象说明》

		Items []*Vendor `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 成员数

		TotalMemberCount *int64 `json:"TotalMemberCount,omitempty" name:"TotalMemberCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVendorListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVendorListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Snapshot struct {

	// 仓库名称

	Snapshot *string `json:"Snapshot,omitempty" name:"Snapshot"`
	// 唯一id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 索引名称

	Indices []*string `json:"Indices,omitempty" name:"Indices"`
	// 存档状态码

	State *string `json:"State,omitempty" name:"State"`
	// 存档状态

	StateZh *string `json:"StateZh,omitempty" name:"StateZh"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建账号名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// DetailDesc

	DetailDesc *string `json:"DetailDesc,omitempty" name:"DetailDesc"`
}

type DescribeAllSystemAndAppRequest struct {
	*tchttp.BaseRequest

	// 应用关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 状态列表

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// AgencyIDs

	AgencyIDs []*string `json:"AgencyIDs,omitempty" name:"AgencyIDs"`
}

func (r *DescribeAllSystemAndAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllSystemAndAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConnectorRequest struct {
	*tchttp.BaseRequest

	// 认证 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 认证名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 认证类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 认证配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// 认证状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 认证标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// AgencyID

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
	// Usage

	Usage *string `json:"Usage,omitempty" name:"Usage"`
}

func (r *ModifyConnectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConnectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgOfSystem struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeTlsCertificateRequest struct {
	*tchttp.BaseRequest

	// 证书ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeTlsCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTlsCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFGatewayServicesRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// api 分组 ID

	GroupID *string `json:"GroupID,omitempty" name:"GroupID"`
	// api 分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// api 分组路径前缀

	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// API模糊搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeTSFGatewayServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFGatewayServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectRenewSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *RejectRenewSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectRenewSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCategoryRequest struct {
	*tchttp.BaseRequest

	// ⽬录id

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入报告文件

		ResultPath *string `json:"ResultPath,omitempty" name:"ResultPath"`
		// 导入成功数

		Success *int64 `json:"Success,omitempty" name:"Success"`
		// 导入失败数

		Fail *int64 `json:"Fail,omitempty" name:"Fail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Trigger struct {

	// Metric

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 范围值

	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
	// Condition

	Condition *string `json:"Condition,omitempty" name:"Condition"`
}

type CreateAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPwdRequest struct {
	*tchttp.BaseRequest

	// 账号

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
}

func (r *DescribeAccountPwdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountPwdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePermissionRequest struct {
	*tchttp.BaseRequest

	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProfileRequest struct {
	*tchttp.BaseRequest

	// SysID

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// PaasID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// Keyword

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// AgencyIDs

	AgencyIDs []*string `json:"AgencyIDs,omitempty" name:"AgencyIDs"`
	// VendorIDs

	VendorIDs []*string `json:"VendorIDs,omitempty" name:"VendorIDs"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProfileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProfileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDownloadLoginLogRequest struct {
	*tchttp.BaseRequest

	// 账号名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 登录类型

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 开始时间 秒级时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间 秒级时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyDownloadLoginLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDownloadLoginLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcsRolePermissionRequest struct {
	*tchttp.BaseRequest

	// 编码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 权限项

	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
}

func (r *ModifyAcsRolePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAcsRolePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTSFGatewayInstanceRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 网关实例 ID

	GatewayInstanceID *string `json:"GatewayInstanceID,omitempty" name:"GatewayInstanceID"`
	// 参考附录 《对象说明》

	ApiGroups []*TSFApiGroup `json:"ApiGroups,omitempty" name:"ApiGroups"`
}

func (r *CreateTSFGatewayInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFGatewayInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstantMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *DescribeRangeMetricsResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstantMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstantMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogoRequest struct {
	*tchttp.BaseRequest

	// 文件base64内容

	Logo *string `json:"Logo,omitempty" name:"Logo"`
	// logo文件名称

	LogoName *string `json:"LogoName,omitempty" name:"LogoName"`
	// logo名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否隐藏名称

	NameDisplay *bool `json:"NameDisplay,omitempty" name:"NameDisplay"`
}

func (r *ModifyLogoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMemberRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserID *string `json:"UserID,omitempty" name:"UserID"`
	// 机构id

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DeleteMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDocumentGroupListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDocumentGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDocumentGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFNamespacesRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 命令空间名称模糊搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeTSFNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSqlTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSqlTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSqlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Apps

		Apps []*AppOfPagedApps `json:"Apps,omitempty" name:"Apps"`
		// Total

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// ApprovingService

		ApprovingService *int64 `json:"ApprovingService,omitempty" name:"ApprovingService"`
		// ApprovingSite

		ApprovingSite *int64 `json:"ApprovingSite,omitempty" name:"ApprovingSite"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApprovedServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServicePageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *string `json:"Total,omitempty" name:"Total"`
		// 当前页面的服务数组

		Items []*ServiceVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServicePageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServicePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserOfCreateApp struct {

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
}

type Zone struct {

	// ZoneID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 新增编辑修改不填

	DefaultZone *bool `json:"DefaultZone,omitempty" name:"DefaultZone"`
	// 新增编辑修改不填

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 新增编辑修改不填

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 新增编辑修改不填

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type DescribeAccountInfoRequest struct {
	*tchttp.BaseRequest

	// 帐号

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
}

func (r *DescribeAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResetTopMenuRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyResetTopMenuRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResetTopMenuRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAgencyRequest struct {
	*tchttp.BaseRequest

	// 机构名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 机构标识

	Badge *string `json:"Badge,omitempty" name:"Badge"`
}

func (r *CheckAgencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAgencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppsRequest struct {
	*tchttp.BaseRequest

	// 系统ID

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// 状态列表。 0 已启用 1 已禁用 2 已删除

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 认证id。例如 rio-atlas

	Connectors []*string `json:"Connectors,omitempty" name:"Connectors"`
	// 分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 机构ID

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

func (r *DescribeAppsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteTop10Request struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSiteTop10Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteTop10Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemRequest struct {
	*tchttp.BaseRequest

	// 系统ID

	SysIDs []*string `json:"SysIDs,omitempty" name:"SysIDs"`
}

func (r *DescribeSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserVendorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录《对象说明》

		Items []*Vendor `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserVendorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserVendorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSiteBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败数量

		FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
		// 成功数量

		SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectedSiteBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSiteBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviveSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *ReviveSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReviveSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgencyRequest struct {
	*tchttp.BaseRequest

	// 机构名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 机构标识

	Badge *string `json:"Badge,omitempty" name:"Badge"`
	// 机构描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 机构负责人

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
}

func (r *CreateAgencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSmartGatewayGroupRequest struct {
	*tchttp.BaseRequest

	// SggID

	SggID *string `json:"SggID,omitempty" name:"SggID"`
	// SggName

	SggName *string `json:"SggName,omitempty" name:"SggName"`
	// 新增编辑修改不填

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// BelongZoneID

	BelongZoneID *string `json:"BelongZoneID,omitempty" name:"BelongZoneID"`
	// 新增编辑修改不填

	BelongZoneName *string `json:"BelongZoneName,omitempty" name:"BelongZoneName"`
	// BelongAreaID

	BelongAreaID *string `json:"BelongAreaID,omitempty" name:"BelongAreaID"`
	// 新增编辑修改不填

	BelongAreaName *string `json:"BelongAreaName,omitempty" name:"BelongAreaName"`
	// Domain

	Domain []*string `json:"Domain,omitempty" name:"Domain"`
	// 查找客户端真实ip时过滤内部ip

	IPs []*string `json:"IPs,omitempty" name:"IPs"`
	// RoleID

	RoleID *string `json:"RoleID,omitempty" name:"RoleID"`
	// 新增编辑修改不填

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 127.0.0.1:8006",

	SgAdminBaseUrl *string `json:"SgAdminBaseUrl,omitempty" name:"SgAdminBaseUrl"`
	// 新增编辑修改不填

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 新增编辑修改不填

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *CreateSmartGatewayGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSmartGatewayGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVendorMemberBatchListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 供应商ID列表

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
}

func (r *DescribeVendorMemberBatchListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVendorMemberBatchListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateVendorMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户帐号

		Account *string `json:"Account,omitempty" name:"Account"`
		// 初始密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateVendorMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateVendorMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRequestLogSortDTO struct {

	// Datestamp

	Datestamp *int64 `json:"Datestamp,omitempty" name:"Datestamp"`
	// Duration

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// StatusCode

	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`
	// ReqLength

	ReqLength *int64 `json:"ReqLength,omitempty" name:"ReqLength"`
	// ResLength

	ResLength *int64 `json:"ResLength,omitempty" name:"ResLength"`
}

type Summary struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*SqlTemplate `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSqlTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSqlTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneRequest struct {
	*tchttp.BaseRequest

	// 分区 ID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
}

func (r *DescribeZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckVendorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 校验结果

		Valid *bool `json:"Valid,omitempty" name:"Valid"`
		// 校验提示信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckVendorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckVendorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCategorysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条⽬

		Total *string `json:"Total,omitempty" name:"Total"`
		// 当前分⻚的⽬录列表

		Items []*CategoryVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCategorysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCategorysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*Role `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVendorMemberListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录《对象说明》

		Items []*VendorUserVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVendorMemberListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVendorMemberListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserPermissionsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserPermissionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserPermissionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncSessionSecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncSessionSecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncSessionSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncType struct {

	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Desc

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type CreateLabelRequest struct {
	*tchttp.BaseRequest

	// 标签名称

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
}

func (r *CreateLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcsRolePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 角色权限信息

		Items []*AcsRolePermission `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAcsRolePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcsRolePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMicroTenantResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMicroTenantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMicroTenantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroProtoFilesRequest struct {
	*tchttp.BaseRequest

	// 文件组 ID

	ProtoGroupId *string `json:"ProtoGroupId,omitempty" name:"ProtoGroupId"`
	// 文件名

	Name []*string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteMicroProtoFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroProtoFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseLifeTimeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLicenseLifeTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseLifeTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否 强制签名校验

		IsForceSign *bool `json:"IsForceSign,omitempty" name:"IsForceSign"`
		// IP白名单个数限制

		IpWhiteListLimit *int64 `json:"IpWhiteListLimit,omitempty" name:"IpWhiteListLimit"`
		// IP黑名单 列表

		IpBlackList []*string `json:"IpBlackList,omitempty" name:"IpBlackList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTodoStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTodoStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTodoStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGrafanaAlertRuleResponse struct {

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
	// Updated

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// State

	State *string `json:"State,omitempty" name:"State"`
	// Metadata

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
}

type ModifySecurityConfigRequest struct {
	*tchttp.BaseRequest

	// 是否 强制签名校验

	IsForceSign *bool `json:"IsForceSign,omitempty" name:"IsForceSign"`
	// IP白名单个数限制

	IpWhiteListLimit *int64 `json:"IpWhiteListLimit,omitempty" name:"IpWhiteListLimit"`
	// IP黑名单 列表

	IpBlackList []*string `json:"IpBlackList,omitempty" name:"IpBlackList"`
}

func (r *ModifySecurityConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVendorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVendorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVendorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFMicroservicesRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 命名空间 ID

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 微服务实例名称模糊搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeTSFMicroservicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFMicroservicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySqlTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySqlTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySqlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyItem struct {

	// key值

	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`
	// value值

	AuthValue *string `json:"AuthValue,omitempty" name:"AuthValue"`
}

type VersionsOfService struct {

	// 版本值

	Version *string `json:"Version,omitempty" name:"Version"`
	// 版本转发路径

	Path *string `json:"Path,omitempty" name:"Path"`
}

type DescribeTodoStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true 有待办 false无待办

		HaveTodos *bool `json:"HaveTodos,omitempty" name:"HaveTodos"`
		// true 有待办 false无待办-系统

		HaveSystemTodos *bool `json:"HaveSystemTodos,omitempty" name:"HaveSystemTodos"`
		// true 有待办 false无待办-服务

		HaveServiceTodos *bool `json:"HaveServiceTodos,omitempty" name:"HaveServiceTodos"`
		// true 有待办 false无待办-站点

		HaveSiteTodos *bool `json:"HaveSiteTodos,omitempty" name:"HaveSiteTodos"`
		// true 有待办 false无待办-订阅

		HaveSubscriptionTodos *bool `json:"HaveSubscriptionTodos,omitempty" name:"HaveSubscriptionTodos"`
		// true 有待办 false无待办-订阅续期

		HaveRenewSubscriptionTodos *bool `json:"HaveRenewSubscriptionTodos,omitempty" name:"HaveRenewSubscriptionTodos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTodoStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTodoStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*MicroApi `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTSFServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tsf struct {

	// 机构信息

	Org *Summary `json:"Org,omitempty" name:"Org"`
	// TSF 租户 secretID

	SecretID *string `json:"SecretID,omitempty" name:"SecretID"`
	// TSF 租户 secretKey

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// 管理端是否 https

	IsHttps *bool `json:"IsHttps,omitempty" name:"IsHttps"`
	// 目标api是否 https

	IsApiHttps *bool `json:"IsApiHttps,omitempty" name:"IsApiHttps"`
	// TSF 微服务管理域名

	EndPoint *string `json:"EndPoint,omitempty" name:"EndPoint"`
	// TSF 微服务 api 域名, 多个域名 “;" 分隔, 根据第一个域名设置 schema

	ApiEndPoint *string `json:"ApiEndPoint,omitempty" name:"ApiEndPoint"`
	// TSF 微服务区域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type BatchDeleteServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVendorMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVendorMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVendorMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequestLogBodyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 服务请求body详情

		Items []*DescribeServiceRequestLogBodyDetailVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceRequestLogBodyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequestLogBodyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadWebConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DownloadWebConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadWebConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendSmsCodeRequest struct {
	*tchttp.BaseRequest

	// 手机号 sm4加密

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 账户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 图形验证码id

	CaptchaID *string `json:"CaptchaID,omitempty" name:"CaptchaID"`
	// 图形验证码

	CaptchaValue *string `json:"CaptchaValue,omitempty" name:"CaptchaValue"`
}

func (r *SendSmsCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendSmsCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifySmsCodeAndModifyPhoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifySmsCodeAndModifyPhoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifySmsCodeAndModifyPhoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscriptionVO struct {

	// ExpireTime

	PublisherService *ServiceVO `json:"PublisherService,omitempty" name:"PublisherService"`
	// CreateTime

	AccessRegions []*RegionOfSubscription `json:"AccessRegions,omitempty" name:"AccessRegions"`
	// UpdateTime

	DeployRegions []*RegionOfSubscription `json:"DeployRegions,omitempty" name:"DeployRegions"`
	// 订阅ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// SubscriberApp

	SubscriberApp *string `json:"SubscriberApp,omitempty" name:"SubscriberApp"`
	// CreateUser

	CreateUser *string `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ApprovalUser

	ApprovalUser *string `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// ApprovalTime

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// 鉴权方式

	SignType []*string `json:"SignType,omitempty" name:"SignType"`
	// 签名方法

	Signature *string `json:"Signature,omitempty" name:"Signature"`
	// 高级规则

	AdvancedRules *string `json:"AdvancedRules,omitempty" name:"AdvancedRules"`
	// ip白名单

	IpWhiteList *int64 `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 0:待审批,1:已审批，2：已拒绝，3：已取消，5：已删除，6：已撤销

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CurRenewStatus

	CurRenewStatus *int64 `json:"CurRenewStatus,omitempty" name:"CurRenewStatus"`
	// Builtin

	Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
	// 过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 限流设置类型，0：默认服务最大限流，1：用户自定义

	InvokeLimitType *int64 `json:"InvokeLimitType,omitempty" name:"InvokeLimitType"`
	// 限流大小

	InvokeLimit *int64 `json:"InvokeLimit,omitempty" name:"InvokeLimit"`
	// 限流统计时间间隔

	InvokeLimitInterval *int64 `json:"InvokeLimitInterval,omitempty" name:"InvokeLimitInterval"`
}

type DescribeSubscriptionApproveLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批列表，参考附录 《对象说明》

		Items []*SubscriptionApproveLogVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionApproveLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionApproveLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVendorMemberBatchListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录《对象说明》

		Items []*VendorUserVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVendorMemberBatchListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVendorMemberBatchListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ItemOfValidateSitePubPath struct {

	// Sys

	Sys *Summary `json:"Sys,omitempty" name:"Sys"`
	// App

	App *Summary `json:"App,omitempty" name:"App"`
	// Site

	Site *Summary `json:"Site,omitempty" name:"Site"`
	// Path

	Path *string `json:"Path,omitempty" name:"Path"`
}

type CheckAcsRoleUniqueRequest struct {
	*tchttp.BaseRequest

	// 过滤器，同bson，例如 {"code": "user_manager"}

	FilterT *string `json:"FilterT,omitempty" name:"FilterT"`
	// 取值：role

	Object *string `json:"Object,omitempty" name:"Object"`
}

func (r *CheckAcsRoleUniqueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAcsRoleUniqueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemAndAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data []*SystemAppInfoVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemAndAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemAndAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthNodeInfo struct {

	// 地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// mac地址

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// cpu核数

	CpuCount *int64 `json:"CpuCount,omitempty" name:"CpuCount"`
}

type ApprovedServiceBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败数量

		FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
		// 成功数量

		SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApprovedServiceBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedServiceBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountAndProfileBatchRequest struct {
	*tchttp.BaseRequest

	// 帐号列表

	AccountIDs []*string `json:"AccountIDs,omitempty" name:"AccountIDs"`
}

func (r *DeleteAccountAndProfileBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountAndProfileBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppRequest struct {
	*tchttp.BaseRequest

	// 应用ID

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
}

func (r *DeleteAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataSource struct {

	// 数据源 id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 数据源名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属应用

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 数据源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据源连接串 dsn

	Dsn *string `json:"Dsn,omitempty" name:"Dsn"`
	// 列表返回最大记录数

	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
	// 数据源限定表/集合名 [BOLA (Broken Object Level Authorization) 被破解的对象级]

	Tables []*string `json:"Tables,omitempty" name:"Tables"`
	// 数据源查询屏蔽字段 [EDE (Excessive Data Exposure) 过度数据暴露]

	QueryIgnoreFields []*string `json:"QueryIgnoreFields,omitempty" name:"QueryIgnoreFields"`
	// 数据源修改屏蔽字段 [MA (Mass Assignment) 批量赋值]

	ModifyIgnoreFields []*string `json:"ModifyIgnoreFields,omitempty" name:"ModifyIgnoreFields"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type CreateMicroRegisterCenterRequest struct {
	*tchttp.BaseRequest

	// 注册中心ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 注册中心名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 注册中心类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 后端服务协议

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 注册中心地址

	Address []*string `json:"Address,omitempty" name:"Address"`
	// 注册中心自定义配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *CreateMicroRegisterCenterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMicroRegisterCenterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroRegisterCentersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMicroRegisterCentersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroRegisterCentersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSiteBatchRequest struct {
	*tchttp.BaseRequest

	// 系统 ID 列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *RejectedSiteBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSiteBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSystemBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败数量

		FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
		// 成功数量

		SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectedSystemBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSystemBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Menu struct {

	// 菜单id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 父级id 为0是主菜单

	ParentID *string `json:"ParentID,omitempty" name:"ParentID"`
	// 菜单类型 1为父菜单 2为子菜单

	MenuType *int64 `json:"MenuType,omitempty" name:"MenuType"`
	// 菜单名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 排序

	Sort *int64 `json:"Sort,omitempty" name:"Sort"`
	// 是否默认菜单

	Protected *bool `json:"Protected,omitempty" name:"Protected"`
	// 是否展示

	Display *bool `json:"Display,omitempty" name:"Display"`
	// 点击事件 1 跳转链接  2 显示子菜单

	ClickEvent *int64 `json:"ClickEvent,omitempty" name:"ClickEvent"`
	// 菜单位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 链接

	Link *Link `json:"Link,omitempty" name:"Link"`
	// icon

	IconURL *string `json:"IconURL,omitempty" name:"IconURL"`
	// icon类型 1默认icon 2自定义icon

	IconType *int64 `json:"IconType,omitempty" name:"IconType"`
	// 打开方式  1 本标签页打开  2 新标签页打开

	OpenMode *int64 `json:"OpenMode,omitempty" name:"OpenMode"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建人

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// GroupName

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 为true不可隐藏该菜单

	NoDisplay *bool `json:"NoDisplay,omitempty" name:"NoDisplay"`
	// 是否隐藏菜单，该属性不能在页面修改

	IsHiddenMenu *bool `json:"IsHiddenMenu,omitempty" name:"IsHiddenMenu"`
	// 子菜单，与外层结构一样

	SubMenus []*Menu `json:"SubMenus,omitempty" name:"SubMenus"`
}

type DescribeConnectorRequest struct {
	*tchttp.BaseRequest

	// 认证 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeConnectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *SqlTemplate `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSqlTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSqlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTlsCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTlsCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTlsCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMemberBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMemberBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMemberBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountAndProfileListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 用户状态 0:已启用, 9: 已禁用

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 排序字段 如{"Account": 1,"CreateTime": -1},1正序，-1 倒序

	Sort *SortOfDescribeAccountAndProfileList `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeAccountAndProfileListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountAndProfileListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySmartGatewayGroupRequest struct {
	*tchttp.BaseRequest

	// SggID

	SggID *string `json:"SggID,omitempty" name:"SggID"`
	// SggName

	SggName *string `json:"SggName,omitempty" name:"SggName"`
	// 新增编辑修改不填

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// BelongZoneID

	BelongZoneID *string `json:"BelongZoneID,omitempty" name:"BelongZoneID"`
	// 新增编辑修改不填

	BelongZoneName *string `json:"BelongZoneName,omitempty" name:"BelongZoneName"`
	// BelongAreaID

	BelongAreaID *string `json:"BelongAreaID,omitempty" name:"BelongAreaID"`
	// 新增编辑修改不填

	BelongAreaName *string `json:"BelongAreaName,omitempty" name:"BelongAreaName"`
	// Domain

	Domain []*string `json:"Domain,omitempty" name:"Domain"`
	// 查找客户端真实ip时过滤内部ip

	IPs []*string `json:"IPs,omitempty" name:"IPs"`
	// RoleID

	RoleID *string `json:"RoleID,omitempty" name:"RoleID"`
	// 新增编辑修改不填

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 127.0.0.1:8006",

	SgAdminBaseUrl *string `json:"SgAdminBaseUrl,omitempty" name:"SgAdminBaseUrl"`
	// 新增编辑修改不填

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 新增编辑修改不填

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *ModifySmartGatewayGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySmartGatewayGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TSFApiGroup struct {

	// api 分组 ID

	GroupID *string `json:"GroupID,omitempty" name:"GroupID"`
	// api 分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// api 分组路径前缀

	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`
}

type DescribeUserVendorsRequest struct {
	*tchttp.BaseRequest

	// 用户id

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
}

func (r *DescribeUserVendorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserVendorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅的ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 允许调用的区域，单分区默认情况下填写：{AreaID: "default_area", ZoneID: "default_zone"}，多分区根据实际情况填写

	DeployRegions []*RegionOfSubscription `json:"DeployRegions,omitempty" name:"DeployRegions"`
	// 限流设置类型，0：默认服务最大限流，1：用户自定义，默认为0

	InvokeLimitType *int64 `json:"InvokeLimitType,omitempty" name:"InvokeLimitType"`
	// 限流大小，大于0

	InvokeLimit *int64 `json:"InvokeLimit,omitempty" name:"InvokeLimit"`
	// 限流统计时间间隔

	InvokeLimitInterval *int64 `json:"InvokeLimitInterval,omitempty" name:"InvokeLimitInterval"`
	// ip白名单

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// 鉴权方式：signature、ipWhiteList，最少选择其中一种

	SignType []*string `json:"SignType,omitempty" name:"SignType"`
	// 签名方法：sha256、sm3、hmac.sm3，选择其中一种

	SignatureType *string `json:"SignatureType,omitempty" name:"SignatureType"`
	// 过期时间，事件戳，单位秒，不填时为申请订阅时填写的过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 高级规则，一般用于在初始化网关规则是修改规则

	AdvancedRules *string `json:"AdvancedRules,omitempty" name:"AdvancedRules"`
}

func (r *UpdateSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SiteWaterConfig struct {

	// 启用范围，1：所有，2：仅根路径

	Range *int64 `json:"Range,omitempty" name:"Range"`
	// 水印内容，1：账号，2：ip

	Content []*int64 `json:"Content,omitempty" name:"Content"`
	// 水印透明度 1-100

	Transparency *int64 `json:"Transparency,omitempty" name:"Transparency"`
	// 间距，1-2000

	Space *int64 `json:"Space,omitempty" name:"Space"`
}

type ZoneArea struct {

	// ZoneID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// DefaultZone

	DefaultZone *bool `json:"DefaultZone,omitempty" name:"DefaultZone"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// Areas

	Areas []*Area `json:"Areas,omitempty" name:"Areas"`
}

type PublisherService struct {

	// 被订阅服务ID

	Svcid *string `json:"Svcid,omitempty" name:"Svcid"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// 被订阅服务所属APP

	App *Summary `json:"App,omitempty" name:"App"`
	// 被订阅服务所属系统

	Sys *Summary `json:"Sys,omitempty" name:"Sys"`
	// 标签

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 机构信息

	Agency *Summary `json:"Agency,omitempty" name:"Agency"`
}

type CreateSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 申请订阅的应用ID

	SubscriberID *string `json:"SubscriberID,omitempty" name:"SubscriberID"`
	// 被订阅的服务ID

	PublisherID *string `json:"PublisherID,omitempty" name:"PublisherID"`
	// 发起调用的区域，单分区默认情况下填写：{AreaID: "default_area", ZoneID: "default_zone"}，多分区根据实际情况填写

	AccessRegions []*RegionOfSubscription `json:"AccessRegions,omitempty" name:"AccessRegions"`
	// ip白名单

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// 鉴权方式：signature、ipWhiteList，最少选择其中一种

	SignType []*string `json:"SignType,omitempty" name:"SignType"`
	// 签名方法：sha256、sm3、hmac.sm3，选择其中一种

	SignatureType *string `json:"SignatureType,omitempty" name:"SignatureType"`
	// 过期时间，时间戳，单位秒，默认0，永久有效

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 高级规则，一般用于在初始化网关规则是修改规则

	AdvancedRules *string `json:"AdvancedRules,omitempty" name:"AdvancedRules"`
}

func (r *CreateSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProfileRequest struct {
	*tchttp.BaseRequest

	// 用户ID 更新时送

	AccountInfo *AccountInfo `json:"AccountInfo,omitempty" name:"AccountInfo"`
	// 用户名称

	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`
}

func (r *UpdateProfileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProfileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroTenantResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMicroTenantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroTenantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppTop10Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// Apps

		Apps []*AppOfSystemAppInfoVO `json:"Apps,omitempty" name:"Apps"`
		// Increment

		Increment *uint64 `json:"Increment,omitempty" name:"Increment"`
		// Total

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppTop10Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppTop10Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDocumentFileRequest struct {
	*tchttp.BaseRequest

	// 文件路径

	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *DescribeDocumentFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDocumentFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAreaRequest struct {
	*tchttp.BaseRequest

	// AreaID

	AreaID *string `json:"AreaID,omitempty" name:"AreaID"`
	// AreaName

	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
	// DefaultArea

	DefaultArea *bool `json:"DefaultArea,omitempty" name:"DefaultArea"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *ModifyAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AgencyVO

		Data *AgencyVo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAgencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAreaRequest struct {
	*tchttp.BaseRequest

	// AreaID

	AreaID *string `json:"AreaID,omitempty" name:"AreaID"`
	// AreaName

	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
	// DefaultArea

	DefaultArea *bool `json:"DefaultArea,omitempty" name:"DefaultArea"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *CreateAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreasRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAreasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FreezeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FreezeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FreezeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApproveRenewSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApproveRenewSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApproveRenewSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroProtoGroupRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteMicroProtoGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroProtoGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsOfSubAppRequest struct {
	*tchttp.BaseRequest

	// 订阅应用ID

	SubscriberID *string `json:"SubscriberID,omitempty" name:"SubscriberID"`
	// 服务名称或标识

	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
	// 订阅状态，不填查询待审批、已审批、续期中、已撤销状态的订阅

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 分页查询数据起始点

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页查询数据数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubscriptionsOfSubAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionsOfSubAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDownloadOperationLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDownloadOperationLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDownloadOperationLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncSmartGatewayGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncSmartGatewayGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncSmartGatewayGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TlsCertificateVO struct {

	// 数据库id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 证书名称

	TlsName *string `json:"TlsName,omitempty" name:"TlsName"`
	// TlsCrt

	TlsCrt *string `json:"TlsCrt,omitempty" name:"TlsCrt"`
	// TlsKey

	TlsKey *string `json:"TlsKey,omitempty" name:"TlsKey"`
	// 域名列表

	TlsDomains []*string `json:"TlsDomains,omitempty" name:"TlsDomains"`
	// 证书的生效时间

	TlsStartData *string `json:"TlsStartData,omitempty" name:"TlsStartData"`
	// 证书的过期时间

	TlsEndDate *string `json:"TlsEndDate,omitempty" name:"TlsEndDate"`
	// 证书有效期（天）

	TlsDuration *string `json:"TlsDuration,omitempty" name:"TlsDuration"`
	// 被使用次数

	UseNum *int64 `json:"UseNum,omitempty" name:"UseNum"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// UpdateUser

	UpdateUser *CreateUser `json:"UpdateUser,omitempty" name:"UpdateUser"`
	// 类型，0：里约自有证书，1:独立域名证书

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 状态(保留字段，过期用TlsDuration或者TlsEndDate判断即可)

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Agency

	Agency *Summary `json:"Agency,omitempty" name:"Agency"`
}

type DescribeCheckProfileRequest struct {
	*tchttp.BaseRequest

	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 证件号码

	CID *string `json:"CID,omitempty" name:"CID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 用户显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
}

func (r *DescribeCheckProfileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckProfileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecurityConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountAndProfileRequest struct {
	*tchttp.BaseRequest

	// 帐号

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
}

func (r *DeleteAccountAndProfileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountAndProfileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVendorProfileRequest struct {
	*tchttp.BaseRequest

	// 机构ID列表

	AgencyIDs []*string `json:"AgencyIDs,omitempty" name:"AgencyIDs"`
	// 供应商ID列表

	VendorIDs []*string `json:"VendorIDs,omitempty" name:"VendorIDs"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeVendorProfileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVendorProfileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyPhoneRequest struct {
	*tchttp.BaseRequest

	// 账号

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 手机密文

	PhoneCipher *string `json:"PhoneCipher,omitempty" name:"PhoneCipher"`
	// 手机区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 应用id

	AppID *string `json:"AppID,omitempty" name:"AppID"`
}

func (r *VerifyPhoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyPhoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataSourceType struct {

	// 数据源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据源连接串示例

	ExampleDsn *string `json:"ExampleDsn,omitempty" name:"ExampleDsn"`
	// 密码脱敏表达式

	DesensitizationDsn *string `json:"DesensitizationDsn,omitempty" name:"DesensitizationDsn"`
}

type SiteUpdateDTO struct {

	// 站点ID

	SiteID *string `json:"SiteID,omitempty" name:"SiteID"`
	// paasID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 站点名称

	SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
	// 公开路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 负载方式：false，随机负载，true：会话保持

	KeepSession *bool `json:"KeepSession,omitempty" name:"KeepSession"`
	// 归档ID

	SiteArchID *string `json:"SiteArchID,omitempty" name:"SiteArchID"`
	// 分区配置信息

	Regions []*RegionsOfSite `json:"Regions,omitempty" name:"Regions"`
	// AdvanceSetting

	AdvanceSetting *AdvanceSetting `json:"AdvanceSetting,omitempty" name:"AdvanceSetting"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Sys

	Sys *SysOfSiteVO `json:"Sys,omitempty" name:"Sys"`
	// App

	App *AppOfSiteVO `json:"App,omitempty" name:"App"`
}

type CreateVendorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 供应商ID

		ID *string `json:"ID,omitempty" name:"ID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVendorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVendorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSiteRequest struct {
	*tchttp.BaseRequest

	// 站点ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmartGateRaidRoutingsRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSmartGateRaidRoutingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSmartGateRaidRoutingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAgencyMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAgencyMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAgencyMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadAccRequest struct {
	*tchttp.BaseRequest

	// form-data文件

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// Content

	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *UploadAccRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadAccRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisabledSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceRequestLogVO struct {

	// 请求时间

	Datestamp *string `json:"Datestamp,omitempty" name:"Datestamp"`
	// 索引后缀

	SourcePrefix *string `json:"SourcePrefix,omitempty" name:"SourcePrefix"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 服务ID

	SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
	// 请求原始路径

	RequestPath *string `json:"RequestPath,omitempty" name:"RequestPath"`
	// 源IP

	OriginIP *string `json:"OriginIP,omitempty" name:"OriginIP"`
	// 请求耗时

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// 目标服务响应

	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`
	// 错误信息

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// 需要查询详情才可见

	ReqBody *string `json:"ReqBody,omitempty" name:"ReqBody"`
	// 请求body长度

	ReqLength *int64 `json:"ReqLength,omitempty" name:"ReqLength"`
	// 需要查询详情才可见

	ResBody *string `json:"ResBody,omitempty" name:"ResBody"`
	// 响应body长度

	ResLength *int64 `json:"ResLength,omitempty" name:"ResLength"`
	// 日志唯一标识

	ReqID *string `json:"ReqID,omitempty" name:"ReqID"`
}

type TargetRankHosts struct {

	//  Host

	Host *string `json:"Host,omitempty" name:"Host"`
	//  Rank

	Rank *uint64 `json:"Rank,omitempty" name:"Rank"`
}

type DownloadResultRequest struct {
	*tchttp.BaseRequest

	// 结果文件下载路径,导入接口返回

	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *DownloadResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemsForNavRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 搜索字段

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
}

func (r *DescribeSystemsForNavRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemsForNavRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseInfoRequest struct {
	*tchttp.BaseRequest

	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeLicenseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVendorProfileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *DescribeProfileResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVendorProfileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVendorProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSystemRequest struct {
	*tchttp.BaseRequest

	// 系统ID

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 机构，见 Org

	Org *OrgOfSystem `json:"Org,omitempty" name:"Org"`
	// 供应商，见 Vendor

	Vendors []*VendorOfPagedApps `json:"Vendors,omitempty" name:"Vendors"`
	// 创建用户，见 User

	Operators []*CreateUser `json:"Operators,omitempty" name:"Operators"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConnectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *Connector `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConnectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledRequest struct {
	*tchttp.BaseRequest

	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DisabledRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestoreAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroProtoFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMicroProtoFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroProtoFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcsRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAcsRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAcsRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRangeMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标类型

		ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
		// 指标数据

		Result []*ResultMetric `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRangeMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRangeMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAgenciesRequest struct {
	*tchttp.BaseRequest

	// UserID

	UserID *string `json:"UserID,omitempty" name:"UserID"`
}

func (r *GetUserAgenciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAgenciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SqlTemplate struct {

	// 服务 id

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 数据源 id

	DataSourceID *string `json:"DataSourceID,omitempty" name:"DataSourceID"`
	// sql 模板

	Sql *string `json:"Sql,omitempty" name:"Sql"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type DescribeTlsCertificateSortDTO struct {

	// 1升序，-1降序

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 1升序，-1降序

	UseNum *string `json:"UseNum,omitempty" name:"UseNum"`
	// 1升序，-1降序

	TlsEndDate *string `json:"TlsEndDate,omitempty" name:"TlsEndDate"`
}

type SubscriptionRenewApproveVO struct {

	// SubscriptionID

	SubscriptionID *string `json:"SubscriptionID,omitempty" name:"SubscriptionID"`
	// Subscription

	Subscription *SubscriptionVO `json:"Subscription,omitempty" name:"Subscription"`
	// ExpireTime

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// UpdateTime

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// UpdateUser

	UpdateUser *CreateUser `json:"UpdateUser,omitempty" name:"UpdateUser"`
}

type DescribeAgencyListRequest struct {
	*tchttp.BaseRequest

	// 限定分页条目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAgencyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgencyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataSourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*DataSource `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataSourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceApprovesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统审批列表，参考附录 《对象说明》

		Items []*ServiceApproveVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceApprovesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceApprovesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionRenewApprovesRequest struct {
	*tchttp.BaseRequest

	// Keyword

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// PublisherAppIDs

	PublisherAppIDs []*string `json:"PublisherAppIDs,omitempty" name:"PublisherAppIDs"`
	// SubscriberIDs

	SubscriberIDs []*string `json:"SubscriberIDs,omitempty" name:"SubscriberIDs"`
}

func (r *DescribeSubscriptionRenewApprovesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionRenewApprovesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedServiceRequest struct {
	*tchttp.BaseRequest

	// 系统 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *ApprovedServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSmartGatewayGroupRequest struct {
	*tchttp.BaseRequest

	// 节点阵列 ID

	SggID *string `json:"SggID,omitempty" name:"SggID"`
}

func (r *DeleteSmartGatewayGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSmartGatewayGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnSyncPaasRequest struct {
	*tchttp.BaseRequest

	// 机构ID

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
	// 页数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUnSyncPaasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnSyncPaasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCurrentUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号ID

		AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
		// 账号名称

		AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
		// 展示名称

		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCurrentUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCurrentUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest

	// 系统ID

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 应用 token 见 Token

	Tokens []*TokenOfCreateApp `json:"Tokens,omitempty" name:"Tokens"`
	// 创建用户，见 User

	Operators []*UserOfCreateApp `json:"Operators,omitempty" name:"Operators"`
	// 系统ID

	Sysid *string `json:"Sysid,omitempty" name:"Sysid"`
	// 应用描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Vendors

	Vendors []*Summary `json:"Vendors,omitempty" name:"Vendors"`
}

func (r *ModifyAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Manager struct {

	// 用户账号id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 证件号码

	CID *string `json:"CID,omitempty" name:"CID"`
	// 帐号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 姓名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 性别

	Gender *string `json:"Gender,omitempty" name:"Gender"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 绑定时间

	BindTime *uint64 `json:"BindTime,omitempty" name:"BindTime"`
	// 机构ID

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

type ChangeStatusAppRequest struct {
	*tchttp.BaseRequest

	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 取值：0: "已启用",2: "已禁用"

	Statuses *int64 `json:"Statuses,omitempty" name:"Statuses"`
}

func (r *ChangeStatusAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeStatusAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCategoryRequest struct {
	*tchttp.BaseRequest

	// ⽬录ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadAccTemplateRequest struct {
	*tchttp.BaseRequest
}

func (r *DownloadAccTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadAccTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySystemRequest struct {
	*tchttp.BaseRequest

	// 系统ID

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 机构，见 Org

	Org *OrgOfSystem `json:"Org,omitempty" name:"Org"`
	// 供应商，见 Vendor

	Vendors []*VendorOfPagedApps `json:"Vendors,omitempty" name:"Vendors"`
	// 创建用户，见 User

	Operators []*CreateUser `json:"Operators,omitempty" name:"Operators"`
}

func (r *ModifySystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAcsRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAcsRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAcsRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTlsCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *TlsCertificateVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTlsCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTlsCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadImportFileRequest struct {
	*tchttp.BaseRequest
}

func (r *DownloadImportFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadImportFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceMarketPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 当前分页的服务列表

		Items []*ServiceVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceMarketPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceMarketPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpWhiteList struct {

	// Address

	Address *string `json:"Address,omitempty" name:"Address"`
}

type DescribeAccountPwdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用ID

		AppID *string `json:"AppID,omitempty" name:"AppID"`
		// 手机区号

		PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
		// 脱敏手机号

		PhoneMask *string `json:"PhoneMask,omitempty" name:"PhoneMask"`
		// 发送标志

		SendFlag *bool `json:"SendFlag,omitempty" name:"SendFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPwdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountPwdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeStatusAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeStatusAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeStatusAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemApprovesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统审批列表

		Items []*SystemApproveVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemApprovesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemApprovesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserPermissionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 扁平化的权限项列表

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserPermissionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserPermissionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionOfSubscription struct {

	// AreaID

	AreaID *string `json:"AreaID,omitempty" name:"AreaID"`
	// AreaName

	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
	// ZoneID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type AddVendorMemberRequest struct {
	*tchttp.BaseRequest

	// 供应商ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 用户帐号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 证件号

	CID *string `json:"CID,omitempty" name:"CID"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 性别

	Gender *string `json:"Gender,omitempty" name:"Gender"`
}

func (r *AddVendorMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVendorMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadMachineInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DownloadMachineInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadMachineInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMicroProtoGroupRequest struct {
	*tchttp.BaseRequest

	// 文件组 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 文件组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件组描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *ModifyMicroProtoGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMicroProtoGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTlsCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTlsCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTlsCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequestLogPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 服务器请求日志数组

		Items []*ServiceRequestLogVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceRequestLogPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequestLogPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMicroRegisterCenterRequest struct {
	*tchttp.BaseRequest

	// 注册中心ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 注册中心名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 注册中心类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 后端服务协议

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 注册中心地址

	Address []*string `json:"Address,omitempty" name:"Address"`
	// 注册中心自定义配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *ModifyMicroRegisterCenterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMicroRegisterCenterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResetLeftMenuResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResetLeftMenuResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResetLeftMenuResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadAgencyImportResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadAgencyImportResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadAgencyImportResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnabledSiteRequest struct {
	*tchttp.BaseRequest

	// 站点ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *EnabledSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnabledSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSmartGateRaidRoutingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSmartGateRaidRoutingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSmartGateRaidRoutingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUser struct {

	// 创建人id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// DisplayName

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
}

type DescribeAreaZonesRequest struct {
	*tchttp.BaseRequest

	// 类型 [site/service]

	Role *string `json:"Role,omitempty" name:"Role"`
}

func (r *DescribeAreaZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisabledResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountsOfAccessRange struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Account

	Account *string `json:"Account,omitempty" name:"Account"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// DisplayName

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
}

type CreateAlertRuleRequest struct {
	*tchttp.BaseRequest

	// 告警规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 告警类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 告警触发条件

	Triggers []*Triggers `json:"Triggers,omitempty" name:"Triggers"`
	// 指标计算规则

	Evaluate *Evaluate `json:"Evaluate,omitempty" name:"Evaluate"`
	// 告警周期

	NotifyInterval *string `json:"NotifyInterval,omitempty" name:"NotifyInterval"`
	// 告警通知渠道

	Receivers []*Receivers `json:"Receivers,omitempty" name:"Receivers"`
	// 告警启用状态

	State *string `json:"State,omitempty" name:"State"`
}

func (r *CreateAlertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgencyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAgencyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgencyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginLogRequest struct {
	*tchttp.BaseRequest

	// 账号名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 登录类型

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 开始时间 秒级时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间 秒级时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeLoginLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用ID

		PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
		// 服务ID

		SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
		// 服务名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 描述信息

		Description *string `json:"Description,omitempty" name:"Description"`
		// 公开路径

		PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
		// 可见性，公开还是私有

		IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
		// 仅适用于里约自有服务初始化时用的，页面不用管这个字段

		Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
		// 负载方式：false随机负载，true会话保持

		KeepSession *string `json:"KeepSession,omitempty" name:"KeepSession"`
		// 目标服务类型

		TargetServiceType *string `json:"TargetServiceType,omitempty" name:"TargetServiceType"`
		// 全局设置

		ServiceAllSet *ServiceAllSet `json:"ServiceAllSet,omitempty" name:"ServiceAllSet"`
		// 分区信息，数组

		Regions *RegionsOfService `json:"Regions,omitempty" name:"Regions"`
		// 高级配置

		AdvanceSetting *AdvanceSettingOfService `json:"AdvanceSetting,omitempty" name:"AdvanceSetting"`
		// 标签

		LabelIDs []*string `json:"LabelIDs,omitempty" name:"LabelIDs"`
		// 服务目录，也叫服务分类

		CategoryIDs []*string `json:"CategoryIDs,omitempty" name:"CategoryIDs"`
		// 服务类别目录

		CategoryIds []*string `json:"CategoryIds,omitempty" name:"CategoryIds"`
		// 接口地址

		DocUrl *string `json:"DocUrl,omitempty" name:"DocUrl"`
		// 接口文档ID

		DocIds []*string `json:"DocIds,omitempty" name:"DocIds"`
		// 应用

		App *Summary `json:"App,omitempty" name:"App"`
		// 系统

		Sys *Summary `json:"Sys,omitempty" name:"Sys"`
		// 机构

		Agency *Summary `json:"Agency,omitempty" name:"Agency"`
		// 状态（草稿0，待审批1，已审批2，禁用5，冻结6）

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 被订阅的应用数

		AuthAppNum *int64 `json:"AuthAppNum,omitempty" name:"AuthAppNum"`
		// 创建者

		CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
		// CreateTime

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// LastUpdateTime

		LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
		// 审批者

		ApprovalUser *CreateUser `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
		// 归档ID

		ServiceArchID *string `json:"ServiceArchID,omitempty" name:"ServiceArchID"`
		// ApprovalTime

		ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
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

type VerifyRequest struct {

	// 账号名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 短信验证码

	SmsCode *string `json:"SmsCode,omitempty" name:"SmsCode"`
	// 验证类型 edit_phone 修改手机号 edit_password 修改密码 login 登录

	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`
}

type CreateDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据源 ID

		ID *string `json:"ID,omitempty" name:"ID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTlsCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTlsCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTlsCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertHistoryRequest struct {
	*tchttp.BaseRequest

	// 时间起点

	From *string `json:"From,omitempty" name:"From"`
	// 时间终点

	To *string `json:"To,omitempty" name:"To"`
	// 限制返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlertHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroTenantsRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMicroTenantsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroTenantsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceApproveVO struct {

	// 所属应用

	App *Summary `json:"App,omitempty" name:"App"`
	// 申请时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 申请人

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// API名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 访问路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// API ID

	SvcId *string `json:"SvcId,omitempty" name:"SvcId"`
}

type DescribeApplyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		ApplyInfos []*ApplyInfo `json:"ApplyInfos,omitempty" name:"ApplyInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApproveRenewSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 续期时间，默认为续期申请时填写续期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *ApproveRenewSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApproveRenewSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号

		Account *string `json:"Account,omitempty" name:"Account"`
		// 密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTSFGatewayServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTSFGatewayServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadMachineInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadMachineInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Evaluate struct {

	// 统计周期

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 超过多少次才告警

	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
}

type DeleteSnapshotOriginDataRequest struct {
	*tchttp.BaseRequest

	// 仓库名 枚举见RepoType

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// 存档名称

	Snapshot *string `json:"Snapshot,omitempty" name:"Snapshot"`
}

func (r *DeleteSnapshotOriginDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotOriginDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteArchRequest struct {
	*tchttp.BaseRequest

	// 站点归档ID（SiteArchID）

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeSiteArchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteArchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendSmsCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否发送成功

		SendSuccess *bool `json:"SendSuccess,omitempty" name:"SendSuccess"`
		// 提示信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 打码手机号

		MaskPhone *string `json:"MaskPhone,omitempty" name:"MaskPhone"`
		// 发送间隔

		Remaining *int64 `json:"Remaining,omitempty" name:"Remaining"`
		// 图形验证码 参考附录 《对象说明》

		Captcha *Captcha `json:"Captcha,omitempty" name:"Captcha"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendSmsCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendSmsCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Area struct {

	// AreaID

	AreaID *string `json:"AreaID,omitempty" name:"AreaID"`
	// AreaName

	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
	// DefaultArea

	DefaultArea *bool `json:"DefaultArea,omitempty" name:"DefaultArea"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type DownloadAppDataRequest struct {
	*tchttp.BaseRequest
}

func (r *DownloadAppDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadAppDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FreezeRequest struct {
	*tchttp.BaseRequest

	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *FreezeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FreezeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EvaluateConf struct {

	// 统计周期

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 超过多少次才告警

	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
}

type CountVendorRequest struct {
	*tchttp.BaseRequest
}

func (r *CountVendorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CountVendorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMemberRequest struct {
	*tchttp.BaseRequest

	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 机构ID

	AgencyId *string `json:"AgencyId,omitempty" name:"AgencyId"`
	// 头像地址

	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
	// 证件号

	CID *string `json:"CID,omitempty" name:"CID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 性别

	Gender *string `json:"Gender,omitempty" name:"Gender"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 手机区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
}

func (r *CreateMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteRequest struct {
	*tchttp.BaseRequest

	// 站点ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportServiceRequest struct {
	*tchttp.BaseRequest

	// 应用ID数组

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 服务目录数组

	CategoryIDs []*string `json:"CategoryIDs,omitempty" name:"CategoryIDs"`
	// 服务标签数组

	LabelIDs []*string `json:"LabelIDs,omitempty" name:"LabelIDs"`
	// 关键词搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 状态数组（草稿0，待审批1，已审批2，禁用5，冻结6）

	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`
	// 限定分页条目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序搜索

	Sort *DescribeServicePageSortDTO `json:"Sort,omitempty" name:"Sort"`
}

func (r *ExportServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateVendorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateVendorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateVendorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCheckAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否有效

		IsValid *bool `json:"IsValid,omitempty" name:"IsValid"`
		// 错误信息

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCheckAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDisableAccountRequest struct {
	*tchttp.BaseRequest

	// 帐号

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
}

func (r *ModifyDisableAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisableAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceRequest struct {
	*tchttp.BaseRequest

	// 0：表示继续保存草稿，1：表示之前的草稿现在提交审批，2：已审批通过的服务编辑（包含已审批之后的其他状态，）

	DraftFlag *int64 `json:"DraftFlag,omitempty" name:"DraftFlag"`
	// 分区信息，数组

	Regions []*RegionsOfService `json:"Regions,omitempty" name:"Regions"`
	// 接口地址

	DocUrl *string `json:"DocUrl,omitempty" name:"DocUrl"`
	// 接口文档ID

	DocIDs []*string `json:"DocIDs,omitempty" name:"DocIDs"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 服务ID

	SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
	// 服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 公开路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 负载方式：false随机负载，true会话保持

	KeepSession *bool `json:"KeepSession,omitempty" name:"KeepSession"`
	// 可见性，公开还是私有

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 仅适用于里约自有服务初始化时用的

	Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
	// 目标服务类型

	TargetServiceType *string `json:"TargetServiceType,omitempty" name:"TargetServiceType"`
	// 全局设置

	ServiceAllSet *ServiceAllSet `json:"ServiceAllSet,omitempty" name:"ServiceAllSet"`
	// 高级设置

	AdvanceSetting *AdvanceSettingOfService `json:"AdvanceSetting,omitempty" name:"AdvanceSetting"`
	// 标签

	LabelIDs []*string `json:"LabelIDs,omitempty" name:"LabelIDs"`
	// 服务目录

	CategoryIDs []*string `json:"CategoryIDs,omitempty" name:"CategoryIDs"`
}

func (r *ModifyServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TokenOfCreateApp struct {

	// zone

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// 应用ID

	Token *string `json:"Token,omitempty" name:"Token"`
	// 被弃用的 token

	DiscardingToken *string `json:"DiscardingToken,omitempty" name:"DiscardingToken"`
}

type CreateSqlTemplateRequest struct {
	*tchttp.BaseRequest

	// 服务 id

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 数据源 id

	DataSourceID *string `json:"DataSourceID,omitempty" name:"DataSourceID"`
	// sql 模板

	Sql *string `json:"Sql,omitempty" name:"Sql"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *CreateSqlTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSqlTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTSFGatewayApiGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTSFGatewayApiGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFGatewayApiGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TSFMicroservice struct {

	// 命名空间标识

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 命名空间描述

	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`
	// 微服务标识

	MicroserviceID *string `json:"MicroserviceID,omitempty" name:"MicroserviceID"`
	// 微服务名称

	MicroserivceName *string `json:"MicroserivceName,omitempty" name:"MicroserivceName"`
	// 微服务描述

	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
	// 微服务实例清单

	MicroserviceHost []*string `json:"MicroserviceHost,omitempty" name:"MicroserviceHost"`
}

type DescribeConnectorsRequest struct {
	*tchttp.BaseRequest

	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 查询类型 manage 管理 use 使用

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
}

func (r *DescribeConnectorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSqlTemplateRequest struct {
	*tchttp.BaseRequest

	// ServiceID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
}

func (r *DeleteSqlTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSqlTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteServiceRequest struct {
	*tchttp.BaseRequest

	// 服务ID数组

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
}

func (r *BatchDeleteServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CountVendorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 供应商总数(当前用户)

		TotalCountVendors *int64 `json:"TotalCountVendors,omitempty" name:"TotalCountVendors"`
		// 供应商成员总数(当前用户)

		TotalCountVendorsUsers *int64 `json:"TotalCountVendorsUsers,omitempty" name:"TotalCountVendorsUsers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CountVendorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CountVendorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroProtoGroupsRequest struct {
	*tchttp.BaseRequest

	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// ID 模糊搜索

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name 模糊搜索

	Name *string `json:"Name,omitempty" name:"Name"`
	// Description 模糊搜索

	Description *string `json:"Description,omitempty" name:"Description"`
	// ID/Name/Description 模糊搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeMicroProtoGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroProtoGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplyListRequest struct {
	*tchttp.BaseRequest

	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApplyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroProtoFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *MicroProtoFile `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroProtoFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroProtoFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFGatewayServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*MicroApi `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTSFGatewayServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFGatewayServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadRequest struct {
	*tchttp.BaseRequest

	// FileName

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *DownloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadImportFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//  Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadImportFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadImportFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcsRoleRequest struct {
	*tchttp.BaseRequest

	// 编码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 角色名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色范围描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAcsRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAcsRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateBtn struct {

	// Info

	Info *string `json:"Info,omitempty" name:"Info"`
	// Submit

	Submit *string `json:"Submit,omitempty" name:"Submit"`
}

type DeleteMicroProtoGroupsRequest struct {
	*tchttp.BaseRequest

	// ID

	ID []*string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteMicroProtoGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroProtoGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProfileRequest struct {
	*tchttp.BaseRequest

	// 机构id

	AgencyIDs []*string `json:"AgencyIDs,omitempty" name:"AgencyIDs"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListProfileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProfileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据源 ID

		ID *string `json:"ID,omitempty" name:"ID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadLicenseRequest struct {
	*tchttp.BaseRequest

	// .lic后缀名的文件

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// Content

	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *UploadLicenseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadLicenseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountAndProfileBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountAndProfileBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountAndProfileBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceApproveLogsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 应用列表

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// Statuses

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
}

func (r *DescribeServiceApproveLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceApproveLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRestoreSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRestoreSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRestoreSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyPhoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyPhoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyPhoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Connector struct {

	// 认证 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 认证名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 认证类型 oidc/cas/saml/ldap

	Type *string `json:"Type,omitempty" name:"Type"`
	// 认证配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// 认证状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 认证标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type RegionsOfSite struct {

	//  分区ID

	DeployZoneID *string `json:"DeployZoneID,omitempty" name:"DeployZoneID"`
	//  区域ID

	DeployAreaID *string `json:"DeployAreaID,omitempty" name:"DeployAreaID"`
	//  目标服务器列表,

	TargetRankHosts []*TargetRankHosts `json:"TargetRankHosts,omitempty" name:"TargetRankHosts"`
	//  目标路径

	TargetPath *string `json:"TargetPath,omitempty" name:"TargetPath"`
	// 协议是否支持https

	TargetIsHttps *bool `json:"TargetIsHttps,omitempty" name:"TargetIsHttps"`
	//  协议是否支持http2，这两个需配合使用，如果是https（HTTP/2）的话，二者都为true

	TargetIsHttp2 *bool `json:"TargetIsHttp2,omitempty" name:"TargetIsHttp2"`
	//  选择https后，是否需要https证书检查

	IgnoreTargetCertsError *bool `json:"IgnoreTargetCertsError,omitempty" name:"IgnoreTargetCertsError"`
	//  是否需要登录，

	NeedLogin *bool `json:"NeedLogin,omitempty" name:"NeedLogin"`
	//  登录方式（多选） 如果填写的NeedLogin = true，此项必填

	ConnectorIDs []*string `json:"ConnectorIDs,omitempty" name:"ConnectorIDs"`
	//  访问范围（候补）

	AccessRange *AccessRange `json:"AccessRange,omitempty" name:"AccessRange"`
	//  限流

	InvokeLimit *int64 `json:"InvokeLimit,omitempty" name:"InvokeLimit"`
	//  限流时间（跟服务一样）

	InvokeLimitInterval *int64 `json:"InvokeLimitInterval,omitempty" name:"InvokeLimitInterval"`
	//  超时（跟服务一样）

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// PubAreaConfig

	PubAreaConfig *PubAreaConfig `json:"PubAreaConfig,omitempty" name:"PubAreaConfig"`
	//  子路径路由规则

	AdvanceRouting []*AdvanceRouting `json:"AdvanceRouting,omitempty" name:"AdvanceRouting"`
}

type DescribeSubscriptionRenewApprovesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Total

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Items

		Items *SubscriptionRenewVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionRenewApprovesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionRenewApprovesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServicePageSortDTO struct {

	// 1升序，-1降序

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 1升序，-1降序

	Name *int64 `json:"Name,omitempty" name:"Name"`
	// 1升序，-1降序

	AuthAppNum *int64 `json:"AuthAppNum,omitempty" name:"AuthAppNum"`
}

type CreateTSFMicroserviceRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 命名空间 ID

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 微服务实例列表 参考附录 《对象说明》

	Microservices []*MicroserviceData `json:"Microservices,omitempty" name:"Microservices"`
}

func (r *CreateTSFMicroserviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFMicroserviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnableAccountRequest struct {
	*tchttp.BaseRequest

	// 帐号

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
}

func (r *ModifyEnableAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEnableAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteApproveLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批列表，参考附录 《对象说明》

		Items []*SiteApproveLogVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSiteApproveLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteApproveLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadImportResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadImportResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadImportResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CategoryVO struct {

	// CategoryID

	CategoryID *string `json:"CategoryID,omitempty" name:"CategoryID"`
	// CategoryName

	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`
	// 绑定

	BindServiceNum *int64 `json:"BindServiceNum,omitempty" name:"BindServiceNum"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type CreateTlsCertificateRequest struct {
	*tchttp.BaseRequest

	// 证书名称

	TlsName *string `json:"TlsName,omitempty" name:"TlsName"`
	// 证书

	TlsCrt *string `json:"TlsCrt,omitempty" name:"TlsCrt"`
	// 私钥

	TlsKey *int64 `json:"TlsKey,omitempty" name:"TlsKey"`
	// 证书类型（0：里约自有证书，1:独立域名证书）

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 机构

	Agency *Summary `json:"Agency,omitempty" name:"Agency"`
}

func (r *CreateTlsCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTlsCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroProtoGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMicroProtoGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroProtoGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgencyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 当前分页的机构列表

		Items []*AgencyVo `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgencyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgencyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadAccResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadAccResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadAccResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadAccRequest struct {
	*tchttp.BaseRequest

	// 导出字段

	Field []*string `json:"Field,omitempty" name:"Field"`
}

func (r *DownloadAccRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadAccRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadImportTemplateRequest struct {
	*tchttp.BaseRequest
}

func (r *DownloadImportTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadImportTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppListByAgencyIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data []*Summary `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppListByAgencyIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppListByAgencyIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroProtoFilesRequest struct {
	*tchttp.BaseRequest

	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 文件组 ID

	ProtoGroupId *string `json:"ProtoGroupId,omitempty" name:"ProtoGroupId"`
	// 文件名模糊搜索

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeMicroProtoFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroProtoFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncSessionSecretRequest struct {
	*tchttp.BaseRequest

	// 会话秘钥

	SessionSecret *string `json:"SessionSecret,omitempty" name:"SessionSecret"`
}

func (r *SyncSessionSecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncSessionSecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateVendorRequest struct {
	*tchttp.BaseRequest

	// 供应商ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标识

	Badge *string `json:"Badge,omitempty" name:"Badge"`
	// 分页大小

	Description *string `json:"Description,omitempty" name:"Description"`
	// 参考附录《对象说明》

	Certs *Cert `json:"Certs,omitempty" name:"Certs"`
	// 管理员ID列表

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
	// 公司地点

	Addr *string `json:"Addr,omitempty" name:"Addr"`
	// 成立年限

	Founded *string `json:"Founded,omitempty" name:"Founded"`
}

func (r *UpdateVendorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateVendorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMenuRequest struct {
	*tchttp.BaseRequest

	// 菜单id

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
}

func (r *DeleteMenuRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMenuRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthNodeOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AuthNodeNum

		AuthNodeNum *int64 `json:"AuthNodeNum,omitempty" name:"AuthNodeNum"`
		// CpuCount

		CpuCount *int64 `json:"CpuCount,omitempty" name:"CpuCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthNodeOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuthNodeOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgencyConfig struct {

	// 机构id

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
	// 配置编码

	ConfigCode *string `json:"ConfigCode,omitempty" name:"ConfigCode"`
	// 配置值

	ConfigVal *string `json:"ConfigVal,omitempty" name:"ConfigVal"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人id

	CreateBy *string `json:"CreateBy,omitempty" name:"CreateBy"`
	// 更新时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 更新人id

	UpdateBy *string `json:"UpdateBy,omitempty" name:"UpdateBy"`
	// 机构名称

	AgencyName *string `json:"AgencyName,omitempty" name:"AgencyName"`
}

type DeleteMemberBatchRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserIDs []*string `json:"UserIDs,omitempty" name:"UserIDs"`
	// 机构id

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

func (r *DeleteMemberBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMemberBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅的ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 允许调用的区域，单分区默认情况下填写：{AreaID: "default_area", ZoneID: "default_zone"}，多分区根据实际情况填写

	DeployRegions []*RegionOfSubscription `json:"DeployRegions,omitempty" name:"DeployRegions"`
	// 限流设置类型，0：默认服务最大限流，1：用户自定义，默认为0

	InvokeLimitType *int64 `json:"InvokeLimitType,omitempty" name:"InvokeLimitType"`
	// 限流大小，大于0

	InvokeLimit *int64 `json:"InvokeLimit,omitempty" name:"InvokeLimit"`
	// 限流统计时间间隔

	InvokeLimitInterval *int64 `json:"InvokeLimitInterval,omitempty" name:"InvokeLimitInterval"`
	// ip白名单

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// 鉴权方式：signature、ipWhiteList，最少选择其中一种

	SignType []*string `json:"SignType,omitempty" name:"SignType"`
	// 签名方法：sha256、sm3、hmac.sm3，选择其中一种

	SignatureType *string `json:"SignatureType,omitempty" name:"SignatureType"`
	// 过期时间，事件戳，单位秒，不填时为申请订阅时填写的过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 高级规则，一般用于在初始化网关规则是修改规则

	AdvancedRules *string `json:"AdvancedRules,omitempty" name:"AdvancedRules"`
}

func (r *ApprovedSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSiteUniqueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true表示校验通过，false则校验不通过

		IsValid *bool `json:"IsValid,omitempty" name:"IsValid"`
		// ValidMsg

		ValidMsg *string `json:"ValidMsg,omitempty" name:"ValidMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSiteUniqueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSiteUniqueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgencyBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAgencyBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgencyBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgencyConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 条数

		Limit *int64 `json:"Limit,omitempty" name:"Limit"`
		// 偏移量

		Offset *int64 `json:"Offset,omitempty" name:"Offset"`
		// 参考附录 《对象说明》

		Items []*AgencyConfig `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgencyConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgencyConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationLogRequest struct {
	*tchttp.BaseRequest

	// 账号名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 模块名称 参考OpModule

	OpModule *string `json:"OpModule,omitempty" name:"OpModule"`
	// 关键词

	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 开始时间 秒级时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间 秒级时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeOperationLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionApproveLogsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 应用列表

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// Statuses

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// SubscriberIDs

	SubscriberIDs []*string `json:"SubscriberIDs,omitempty" name:"SubscriberIDs"`
}

func (r *DescribeSubscriptionApproveLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionApproveLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionApprovesRequest struct {
	*tchttp.BaseRequest

	// 发布服务的应用ID

	PublisherAppIDs []*string `json:"PublisherAppIDs,omitempty" name:"PublisherAppIDs"`
	// 订阅服务的应用ID

	SubscriberIDs []*string `json:"SubscriberIDs,omitempty" name:"SubscriberIDs"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubscriptionApprovesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionApprovesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequestLogBodyDetailVO struct {

	// SourcePrefix

	SourcePrefix *string `json:"SourcePrefix,omitempty" name:"SourcePrefix"`
	// 是否有记录到报文日志

	OpenMessageLog *bool `json:"OpenMessageLog,omitempty" name:"OpenMessageLog"`
	// 请求头

	RequestHeaders *string `json:"RequestHeaders,omitempty" name:"RequestHeaders"`
	// 请求body

	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
	// 响应头

	ResponseHeaders *string `json:"ResponseHeaders,omitempty" name:"ResponseHeaders"`
	// 响应body

	ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`
}

type DeleteMicroRegisterCenterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMicroRegisterCenterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroRegisterCenterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcsUser struct {

	// true

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
	// true

	Account *string `json:"Account,omitempty" name:"Account"`
	// true

	Protected *bool `json:"Protected,omitempty" name:"Protected"`
	// true

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// true

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// true

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// true

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// true

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// true

	Name *string `json:"Name,omitempty" name:"Name"`
	// true

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// true

	CID *string `json:"CID,omitempty" name:"CID"`
	// true

	CType *string `json:"CType,omitempty" name:"CType"`
	// true

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// true

	Email *string `json:"Email,omitempty" name:"Email"`
	// true

	Gender *string `json:"Gender,omitempty" name:"Gender"`
	// true

	Phone *string `json:"Phone,omitempty" name:"Phone"`
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
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

type ModifySiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectRenewSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectRenewSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectRenewSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstantMetricsRequest struct {
	*tchttp.BaseRequest

	// PromQL语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 时间戳

	Time *string `json:"Time,omitempty" name:"Time"`
	// 时间间隔

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *DescribeInstantMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstantMetricsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAcsRoleRequest struct {
	*tchttp.BaseRequest

	// 编码

	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *DeleteAcsRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAcsRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用总数

		AppTotal *int64 `json:"AppTotal,omitempty" name:"AppTotal"`
		// 系统总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 系统列表

		Systems []*System `json:"Systems,omitempty" name:"Systems"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAgencyProfileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *ListProfileResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAgencyProfileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAgencyProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneRequest struct {
	*tchttp.BaseRequest

	// 分区 ID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
}

func (r *DeleteZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTop10Request struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeServiceTop10Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceTop10Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddVendorMemberByIDsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddVendorMemberByIDsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVendorMemberByIDsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotRequest struct {
	*tchttp.BaseRequest

	// 仓库名 枚举见RepoType

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// 存档名称

	Snapshot *string `json:"Snapshot,omitempty" name:"Snapshot"`
}

func (r *DeleteSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppRequest struct {
	*tchttp.BaseRequest

	// 应用id

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
}

func (r *DescribeAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppListByAgencyIdRequest struct {
	*tchttp.BaseRequest

	// 机构ID数组

	AgencyIDs []*string `json:"AgencyIDs,omitempty" name:"AgencyIDs"`
}

func (r *DescribeAppListByAgencyIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppListByAgencyIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmartGateRaidRoutingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *SmartGateRaidRouting `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmartGateRaidRoutingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSmartGateRaidRoutingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotRequest struct {
	*tchttp.BaseRequest

	// 仓库名 枚举见RepoType

	Repo *string `json:"Repo,omitempty" name:"Repo"`
}

func (r *DescribeSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadVendorImportResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadVendorImportResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadVendorImportResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVendorRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标识

	Badge *string `json:"Badge,omitempty" name:"Badge"`
	// 分页大小

	Description *string `json:"Description,omitempty" name:"Description"`
	// 参考附录《对象说明》

	Certs *Cert `json:"Certs,omitempty" name:"Certs"`
	// 管理员ID列表

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
	// 公司地点

	Addr *string `json:"Addr,omitempty" name:"Addr"`
	// 成立年限

	Founded *string `json:"Founded,omitempty" name:"Founded"`
}

func (r *CreateVendorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVendorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAreaRequest struct {
	*tchttp.BaseRequest

	// 区域 ID

	AreaID *string `json:"AreaID,omitempty" name:"AreaID"`
}

func (r *DeleteAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsOfSubAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *string `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*Subscription `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionsOfSubAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionsOfSubAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcsRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一标识

		Code *string `json:"Code,omitempty" name:"Code"`
		// 角色名

		Name *string `json:"Name,omitempty" name:"Name"`
		// 描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 全局或者机构

		Scope *string `json:"Scope,omitempty" name:"Scope"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAcsRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcsRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCheckProfileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *CheckUniqueResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCheckProfileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTlsCertificateRequest struct {
	*tchttp.BaseRequest

	// 证书ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 证书名称

	TlsName *string `json:"TlsName,omitempty" name:"TlsName"`
	// 证书

	TlsCrt *string `json:"TlsCrt,omitempty" name:"TlsCrt"`
	// 私钥

	TlsKey *int64 `json:"TlsKey,omitempty" name:"TlsKey"`
	// 机构

	Agency *Summary `json:"Agency,omitempty" name:"Agency"`
}

func (r *UpdateTlsCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTlsCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MicroSyncJob struct {

	// 唯一标识

	ID *string `json:"ID,omitempty" name:"ID"`
	// 租户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 微服务类型

	TenantType *string `json:"TenantType,omitempty" name:"TenantType"`
	// 同步方式, namespace/microservice/service

	SyncType *string `json:"SyncType,omitempty" name:"SyncType"`
	// 系统标识

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// 系统名称

	SysName *string `json:"SysName,omitempty" name:"SysName"`
	// 应用标识

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 应用名称

	PaasName *string `json:"PaasName,omitempty" name:"PaasName"`
	// 目标路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 状态, 0: 初始, 1: 同步中, 2: 同步成功 3: 同步失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 初次同步时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 上次同步时间

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type SubscriberApp struct {

	// 订阅应用ID

	Paasid *string `json:"Paasid,omitempty" name:"Paasid"`
	// 订阅应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 订阅应用所属系统

	Sys *Summary `json:"Sys,omitempty" name:"Sys"`
}

type SystemApproveVO struct {

	// 系统id

	SysId *string `json:"SysId,omitempty" name:"SysId"`
	// 系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属机构

	Org *OrgOfSystem `json:"Org,omitempty" name:"Org"`
	// 负责人

	Operators []*CreateUser `json:"Operators,omitempty" name:"Operators"`
	// 申请人

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// 申请时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 系统描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 供应商

	Vendors []*Summary `json:"Vendors,omitempty" name:"Vendors"`
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
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

type DeleteConnectorRequest struct {
	*tchttp.BaseRequest

	// 认证 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteConnectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConnectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServicePageRequest struct {
	*tchttp.BaseRequest

	// 应用ID数组

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 服务目录数组

	CategoryIDs []*string `json:"CategoryIDs,omitempty" name:"CategoryIDs"`
	// 服务标签数组

	LabelIDs []*string `json:"LabelIDs,omitempty" name:"LabelIDs"`
	// 关键词搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 归档状态（0：全部，3：已拒绝，7：已删除）默认取0

	ArchStatus *int64 `json:"ArchStatus,omitempty" name:"ArchStatus"`
	// 状态数组（草稿0，待审批1，已审批2，禁用5，冻结6）

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 限定分页条目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序搜索

	Sort *DescribeServicePageSortDTO `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeServicePageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServicePageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceDetailRespVO struct {

	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 服务ID

	SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
	// 服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 公开路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 可见性，公开还是私有

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 仅适用于里约自有服务初始化时用的，页面不用管这个字段

	Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
	// 负载方式：false随机负载，true会话保持

	KeepSession *string `json:"KeepSession,omitempty" name:"KeepSession"`
	// 目标服务类型

	TargetServiceType *string `json:"TargetServiceType,omitempty" name:"TargetServiceType"`
	// 全局设置

	ServiceAllSet *ServiceAllSet `json:"ServiceAllSet,omitempty" name:"ServiceAllSet"`
	// 分区信息，数组

	Regions *RegionsOfService `json:"Regions,omitempty" name:"Regions"`
	// 高级配置

	AdvanceSetting *AdvanceSettingOfService `json:"AdvanceSetting,omitempty" name:"AdvanceSetting"`
	// 标签

	LabelIDs []*string `json:"LabelIDs,omitempty" name:"LabelIDs"`
	// 服务目录，也叫服务分类

	CategoryIDs []*string `json:"CategoryIDs,omitempty" name:"CategoryIDs"`
	// 服务类别目录

	CategoryIds []*string `json:"CategoryIds,omitempty" name:"CategoryIds"`
	// 接口地址

	DocUrl *string `json:"DocUrl,omitempty" name:"DocUrl"`
	// 接口文档ID

	DocIds []*string `json:"DocIds,omitempty" name:"DocIds"`
	// 应用

	App *Summary `json:"App,omitempty" name:"App"`
	// 系统

	Sys *Summary `json:"Sys,omitempty" name:"Sys"`
	// 机构

	Agency *Summary `json:"Agency,omitempty" name:"Agency"`
	// 状态（草稿0，待审批1，已审批2，禁用5，冻结6）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 被订阅的应用数

	AuthAppNum *int64 `json:"AuthAppNum,omitempty" name:"AuthAppNum"`
	// 创建者

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 审批者

	ApprovalUser *CreateUser `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// 归档ID

	ServiceArchID *string `json:"ServiceArchID,omitempty" name:"ServiceArchID"`
	// ApprovalTime

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
}

type DescribeDocumentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *DocNodeVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDocumentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDocumentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroProtoGroupRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeMicroProtoGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroProtoGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroTenantResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *MicroTenant `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroTenantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroTenantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnableAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnableAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEnableAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest

	// 应用id

	AppID *string `json:"AppID,omitempty" name:"AppID"`
	// 账号

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 当前密码

	CurrentPassword *string `json:"CurrentPassword,omitempty" name:"CurrentPassword"`
	// 新密码

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
	// 加密随机值

	Nonce *string `json:"Nonce,omitempty" name:"Nonce"`
	// 重置密码记录code

	ResetVerifyCode *string `json:"ResetVerifyCode,omitempty" name:"ResetVerifyCode"`
}

func (r *ModifyPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSystemWarnMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AppEnabled

		AppEnabled []*int64 `json:"AppEnabled,omitempty" name:"AppEnabled"`
		// Approving

		Approving []*int64 `json:"Approving,omitempty" name:"Approving"`
		// ConfirmingSubscription

		ConfirmingSubscription []*int64 `json:"ConfirmingSubscription,omitempty" name:"ConfirmingSubscription"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSystemWarnMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSystemWarnMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportAgencyMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ErrorFile

		ErrorFile *string `json:"ErrorFile,omitempty" name:"ErrorFile"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportAgencyMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportAgencyMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgencyConfigsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAgencyConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgencyConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequestLogBodyDetailRequest struct {
	*tchttp.BaseRequest

	// 索引后缀

	SourcePrefix *string `json:"SourcePrefix,omitempty" name:"SourcePrefix"`
	// 服务ID

	SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
	// 日志唯一标识

	ReqID *string `json:"ReqID,omitempty" name:"ReqID"`
}

func (r *DescribeServiceRequestLogBodyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequestLogBodyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFServicesRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 微服务实例 ID

	MicroserviceID *string `json:"MicroserviceID,omitempty" name:"MicroserviceID"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 微服务实例名称模糊搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeTSFServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnabledSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnabledSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnabledSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MicroProtoFile struct {

	// 文件组 ID

	ProtoGroupId *string `json:"ProtoGroupId,omitempty" name:"ProtoGroupId"`
	// 文件名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件内容

	Text *string `json:"Text,omitempty" name:"Text"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DeleteMenuResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMenuResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMenuResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroProtoGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*MicroProtoGroup `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroProtoGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroProtoGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionRenewApproveLogsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 被订阅应用列表

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 订阅应用列表

	SubscriberIDs []*string `json:"SubscriberIDs,omitempty" name:"SubscriberIDs"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// Statuses

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
}

func (r *DescribeSubscriptionRenewApproveLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionRenewApproveLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnFreezeSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnFreezeSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnFreezeSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemAppInfoVO struct {

	// System

	System *SystemOfSystemAppInfoVO `json:"System,omitempty" name:"System"`
	// 应用总数量

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// Apps

	Apps []*AppOfSystemAppInfoVO `json:"Apps,omitempty" name:"Apps"`
}

type AddAgencyManagerRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserIDs []*string `json:"UserIDs,omitempty" name:"UserIDs"`
	// 机构id

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

func (r *AddAgencyManagerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAgencyManagerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataSourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDataSourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 偏移量

		Offset *int64 `json:"Offset,omitempty" name:"Offset"`
		// 条数

		Limit *int64 `json:"Limit,omitempty" name:"Limit"`
		// 开始时间

		StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
		// 登录日志列表

		Logs []*LoginLog `json:"Logs,omitempty" name:"Logs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResetPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 帐号

		Account *string `json:"Account,omitempty" name:"Account"`
		// 密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResetPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifySmsCodeRequest struct {
	*tchttp.BaseRequest

	// 账户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 验证码

	SmsCode *string `json:"SmsCode,omitempty" name:"SmsCode"`
	// 验证类型

	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`
}

func (r *VerifySmsCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifySmsCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgencyVo struct {

	// 机构id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 机构名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 机构标识

	Badge *string `json:"Badge,omitempty" name:"Badge"`
	// 机构描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 机构负责人列表

	Managers []*Manager `json:"Managers,omitempty" name:"Managers"`
	// CreateTime

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
}

type SortOfDescribeAccountAndProfileList struct {

	// 1正序，-1 倒序

	Account *int64 `json:"Account,omitempty" name:"Account"`
	// 1正序，-1 倒序

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// 参考附录 《对象说明》

	AccountInfo *AccountInfo `json:"AccountInfo,omitempty" name:"AccountInfo"`
	// 参考附录 《对象说明》

	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`
}

func (r *CreateAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgencyConfigRequest struct {
	*tchttp.BaseRequest

	// 机构id

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
}

func (r *CreateAgencyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgencyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCategorysRequest struct {
	*tchttp.BaseRequest

	// 限定分⻚条⽬

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分⻚

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 关键词查询

	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
	// 排序

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeCategorysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCategorysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMenuRequest struct {
	*tchttp.BaseRequest

	// 位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 菜单 参考附录 《对象说明》

	Menus []*Menu `json:"Menus,omitempty" name:"Menus"`
}

func (r *ModifyMenuRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMenuRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemApproveLogsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 机构id

	OrgIds []*string `json:"OrgIds,omitempty" name:"OrgIds"`
	// Statuses

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
}

func (r *DescribeSystemApproveLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemApproveLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordByOldResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPasswordByOldResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordByOldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotRequest struct {
	*tchttp.BaseRequest

	// 仓库名 枚举见RepoType

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// 开始时间 秒级时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间 秒级时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *CreateSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroSyncJobRequest struct {
	*tchttp.BaseRequest

	// 同步记录 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeMicroSyncJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroSyncJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FreezeSiteRequest struct {
	*tchttp.BaseRequest

	// 站点ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *FreezeSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FreezeSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAgenciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data []*AgencyVo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserAgenciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAgenciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApproveLogVO struct {

	// 审批记录id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 审批人

	ApprovalUser *string `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// 审批时间

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 驳回原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type AreaZone struct {

	// AreaID

	AreaID *string `json:"AreaID,omitempty" name:"AreaID"`
	// AreaName

	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
	// Domain

	Domain []*string `json:"Domain,omitempty" name:"Domain"`
	// DefaultArea

	DefaultArea *bool `json:"DefaultArea,omitempty" name:"DefaultArea"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// Zones

	Zones []*Zone `json:"Zones,omitempty" name:"Zones"`
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest

	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条⽬

		Total *string `json:"Total,omitempty" name:"Total"`
		// Items

		Items *LabelVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFRegionListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Items

		Items []*string `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTSFRegionListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFRegionListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFSyncTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTSFSyncTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFSyncTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTSFServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTSFServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceOfSubAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *string `json:"Total,omitempty" name:"Total"`
		// 当前页面的服务数组

		Items []*ServiceSubscriptionVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceOfSubAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceOfSubAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCentreNodeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *DescribeAuthInfoResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCentreNodeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCentreNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmartGatewayGroupsRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSmartGatewayGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSmartGatewayGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelVO struct {

	// LabelID

	LabelID *string `json:"LabelID,omitempty" name:"LabelID"`
	// LabelName

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
	// 绑定

	BindServiceNum *int64 `json:"BindServiceNum,omitempty" name:"BindServiceNum"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type DescribeAlertRuleRequest struct {
	*tchttp.BaseRequest

	// 告警规则名称，不填返回所有规则

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAlertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAgencyProfileRequest struct {
	*tchttp.BaseRequest

	// 机构id

	AgencyIDs []*string `json:"AgencyIDs,omitempty" name:"AgencyIDs"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListAgencyProfileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAgencyProfileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroSyncJobsRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	TenantID *string `json:"TenantID,omitempty" name:"TenantID"`
	// 导入粒度

	SyncType *string `json:"SyncType,omitempty" name:"SyncType"`
	// 导入源关键字

	SourceKeyword *string `json:"SourceKeyword,omitempty" name:"SourceKeyword"`
	// 导入应用数组

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 同步状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMicroSyncJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroSyncJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnFreezeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnFreezeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnFreezeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppOfPagedApps struct {

	// PaasID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// SysID

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// SysName

	SysName *string `json:"SysName,omitempty" name:"SysName"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
	// Tokens

	Tokens []*TokenOfPagedApps `json:"Tokens,omitempty" name:"Tokens"`
	// Operators

	Operators []*CreateUser `json:"Operators,omitempty" name:"Operators"`
	// Vendors

	Vendors []*string `json:"Vendors,omitempty" name:"Vendors"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Role struct {

	// RoleID

	RoleID *string `json:"RoleID,omitempty" name:"RoleID"`
	// RoleName

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type TokenOfPagedApps struct {

	// ZoneID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// Token

	Token *string `json:"Token,omitempty" name:"Token"`
	// DiscardingToken

	DiscardingToken *string `json:"DiscardingToken,omitempty" name:"DiscardingToken"`
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest

	// ZoneID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 新增编辑修改不填

	DefaultZone *bool `json:"DefaultZone,omitempty" name:"DefaultZone"`
	// 新增编辑修改不填

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 新增编辑修改不填

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 新增编辑修改不填

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *CreateZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcsRoleMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true

		Total *string `json:"Total,omitempty" name:"Total"`
		// true

		Items []*AcsUser `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAcsRoleMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcsRoleMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAdvanceTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api缓存模板

		Name *string `json:"Name,omitempty" name:"Name"`
		// ID

		ID *string `json:"ID,omitempty" name:"ID"`
		// AllowCount

		AllowCount *int64 `json:"AllowCount,omitempty" name:"AllowCount"`
		// Btn

		Btn *BtnOfDescribeAdvanceTemplate `json:"Btn,omitempty" name:"Btn"`
		// ConfigStr

		ConfigStr *string `json:"ConfigStr,omitempty" name:"ConfigStr"`
		// Config

		Config *ConfigOfDescribeAdvanceTemplate `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAdvanceTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAdvanceTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifySmsCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重置密码验证ID

		ResetPwdVerificationID *string `json:"ResetPwdVerificationID,omitempty" name:"ResetPwdVerificationID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifySmsCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifySmsCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSqlTemplatesRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSqlTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSqlTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertHistoryFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 文件内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertHistoryFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertHistoryFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 权限信息

		Items []*ParentPermission `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemApprovesRequest struct {
	*tchttp.BaseRequest

	// 机构id列表

	OrgIds []*string `json:"OrgIds,omitempty" name:"OrgIds"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 应用状态

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
}

func (r *DescribeSystemApprovesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemApprovesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAgencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseLifeTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最大过期时间

		LifeTime *uint64 `json:"LifeTime,omitempty" name:"LifeTime"`
		// 提示信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 是否有效

		IsValid *bool `json:"IsValid,omitempty" name:"IsValid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseLifeTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseLifeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVendorListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeVendorListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVendorListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyInfo struct {

	// 节点id

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 节点ip地址

	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`
	// 节点host名称

	NodeHostname *string `json:"NodeHostname,omitempty" name:"NodeHostname"`
	// 申请编号id

	ApplyNo *string `json:"ApplyNo,omitempty" name:"ApplyNo"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 申请时间

	ApplyTime *int64 `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 刷新时间

	RefreshTime *int64 `json:"RefreshTime,omitempty" name:"RefreshTime"`
	// 申请状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// ApplyItem

	ApplyItem []*ApplyItem `json:"ApplyItem,omitempty" name:"ApplyItem"`
}

type BatchDeleteCategoryRequest struct {
	*tchttp.BaseRequest

	// ⽬录id数组

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
}

func (r *BatchDeleteCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BtnOfDescribeAdvanceTemplate struct {

	// Info

	Info *string `json:"Info,omitempty" name:"Info"`
	// Submit

	Submit *string `json:"Submit,omitempty" name:"Submit"`
}

type CheckUniqueResponse struct {

	// IsValid

	IsValid *bool `json:"IsValid,omitempty" name:"IsValid"`
	// ValidMsg

	ValidMsg *string `json:"ValidMsg,omitempty" name:"ValidMsg"`
}

type TemplateAdvanceSetList struct {

	// 自定义高级配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 自定义高级配置描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 优先级，数值越低，优先级越高，默认递增

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 开关与否，是否启用它，必填

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 下拉框选择，必填

	TemplateID *string `json:"TemplateID,omitempty" name:"TemplateID"`
	// 模板的按钮列表，显示用，无需提交

	TemplateBtn []*TemplateBtn `json:"TemplateBtn,omitempty" name:"TemplateBtn"`
	// 模板脚本yaml数据格式

	TemplateScript *string `json:"TemplateScript,omitempty" name:"TemplateScript"`
}

type CreateSmartGateRaidRoutingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSmartGateRaidRoutingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSmartGateRaidRoutingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroProtoGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *MicroProtoGroup `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroProtoGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroProtoGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportVendorMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误文件

		ErrorFile *string `json:"ErrorFile,omitempty" name:"ErrorFile"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportVendorMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportVendorMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteLabelRequest struct {
	*tchttp.BaseRequest

	// 标签id数组

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
}

func (r *BatchDeleteLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSystemRequest struct {
	*tchttp.BaseRequest

	// 系统ID列表

	SysIDs []*string `json:"SysIDs,omitempty" name:"SysIDs"`
}

func (r *DeleteSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgenciesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeAgenciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgenciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledSiteRequest struct {
	*tchttp.BaseRequest

	// 站点ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DisabledSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceAllSet struct {

	// 请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 是否开启入参校验

	RequestParamsValidatorStatus *bool `json:"RequestParamsValidatorStatus,omitempty" name:"RequestParamsValidatorStatus"`
	// 入参校验配置

	RequestParamsValidatorJsonInfoT *string `json:"RequestParamsValidatorJsonInfoT,omitempty" name:"RequestParamsValidatorJsonInfoT"`
	// 是否开启出参校验

	ResponseParamsValidatorStatus *bool `json:"ResponseParamsValidatorStatus,omitempty" name:"ResponseParamsValidatorStatus"`
	// 出参校验配置，没有query

	ResponseParamsValidatorJsonInfoT *string `json:"ResponseParamsValidatorJsonInfoT,omitempty" name:"ResponseParamsValidatorJsonInfoT"`
}

type ResultMetric struct {

	// Metric

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// Data

	Data []*MetricData `json:"Data,omitempty" name:"Data"`
}

type DescribeSystemsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 搜索字段

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 查询单个系统，Keyword和Statuses将失效

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// 机构ID列表

	OrgIds []*string `json:"OrgIds,omitempty" name:"OrgIds"`
}

func (r *DescribeSystemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FreezeSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FreezeSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FreezeSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgencyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAgencyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgencyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataSourceRequest struct {
	*tchttp.BaseRequest

	// 数据源 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcsRolePermission struct {

	// 角色名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色编码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 依赖项

	UIDep []*string `json:"UIDep,omitempty" name:"UIDep"`
}

type UserProfileWithExist struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// AccountName

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// DisplayName

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 是否是成员

	ExistInMember *bool `json:"ExistInMember,omitempty" name:"ExistInMember"`
	// 是否是管理员

	ExistInManager *bool `json:"ExistInManager,omitempty" name:"ExistInManager"`
	// Account

	Account *string `json:"Account,omitempty" name:"Account"`
}

type ModifyMicroProtoGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMicroProtoGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMicroProtoGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type System struct {

	// SysID

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Org

	Org *OrgOfSystem `json:"Org,omitempty" name:"Org"`
	// Operators

	Operators []*CreateUser `json:"Operators,omitempty" name:"Operators"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// Vendors

	Vendors []*VendorOfPagedApps `json:"Vendors,omitempty" name:"Vendors"`
}

type DescribeSystemsForNavResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统ID

		SysID *string `json:"SysID,omitempty" name:"SysID"`
		// 系统名

		Name *string `json:"Name,omitempty" name:"Name"`
		// app 数量

		Appnum *int64 `json:"Appnum,omitempty" name:"Appnum"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemsForNavResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemsForNavResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceVO struct {

	// API标识

	SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
	// 所属应用

	App *Summary `json:"App,omitempty" name:"App"`
	// 所属系统

	Sys *Summary `json:"Sys,omitempty" name:"Sys"`
	// API名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// API描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// 服务可见性，true公开，false私有

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 状态（草稿0，待审批1，已审批2，禁用5，冻结6）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// jiekl.dizhi",

	DocUrl *string `json:"DocUrl,omitempty" name:"DocUrl"`
	// 创建用户，下同

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ApprovalUser

	ApprovalUser *CreateUser `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// 目标服务类型

	TargetServiceType *string `json:"TargetServiceType,omitempty" name:"TargetServiceType"`
	// ApprovalTime

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 授权app数

	SubscriptionCount *int64 `json:"SubscriptionCount,omitempty" name:"SubscriptionCount"`
	// 归档ID

	ServiceArchID *string `json:"ServiceArchID,omitempty" name:"ServiceArchID"`
}

type DeleteVendorMemberRequest struct {
	*tchttp.BaseRequest

	// 供应商ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 用户ID列表

	UserIDs []*string `json:"UserIDs,omitempty" name:"UserIDs"`
}

func (r *DeleteVendorMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVendorMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllSystemAndAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data []*SystemAppInfoVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllSystemAndAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllSystemAndAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *RejectedSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRestoreSnapshotRequest struct {
	*tchttp.BaseRequest

	// 仓库名 枚举见RepoType

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// 存档名称

	Snapshot *string `json:"Snapshot,omitempty" name:"Snapshot"`
}

func (r *ModifyRestoreSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRestoreSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceArchRequest struct {
	*tchttp.BaseRequest

	// 服务标识

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeServiceArchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceArchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedServiceBatchRequest struct {
	*tchttp.BaseRequest

	// 系统 ID 列表

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
	// 系统 ID 列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *RejectedServiceBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedServiceBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PubArea struct {

	//  区域ID

	PubAreaID *string `json:"PubAreaID,omitempty" name:"PubAreaID"`
	//  0:共享域名（沿用区域的domain），1:独立域名（自行填写PubHost）

	HostType *int64 `json:"HostType,omitempty" name:"HostType"`
	//  如果HostType = 0，区域domain填充，不可编辑。如果HostType = 1，该处为自定义host输入表单

	PubHost *string `json:"PubHost,omitempty" name:"PubHost"`
	//  预览访问地址：如果HostType = 0，则为当前PubHost /paasId/ 最顶层的那个PubPath，如果HostType = 1，则为当前PubHost / 最顶层的那个PubPath。皆不可编辑

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	//  是否开启静态资源访问

	StaticAbsolute *bool `json:"StaticAbsolute,omitempty" name:"StaticAbsolute"`
	//  协议，独立域名时有效，（-1：http，0：https，1：http & https）

	NeedProtocol *int64 `json:"NeedProtocol,omitempty" name:"NeedProtocol"`
	//  证书ID，独立域名时有效

	TlsID *string `json:"TlsID,omitempty" name:"TlsID"`
}

type CreateMicroRegisterCenterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMicroRegisterCenterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMicroRegisterCenterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreAppRequest struct {
	*tchttp.BaseRequest

	// 应用ID

	PaasID []*string `json:"PaasID,omitempty" name:"PaasID"`
}

func (r *RestoreAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestoreAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviveSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReviveSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReviveSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAgencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// sss

		Test *string `json:"Test,omitempty" name:"Test"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAgencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAgencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DocGroupVO struct {

	// 分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// Nodes

	Nodes []*DocNodeVO `json:"Nodes,omitempty" name:"Nodes"`
}

type RegionsOfService struct {

	// 分区ID

	DeployZoneID *string `json:"DeployZoneID,omitempty" name:"DeployZoneID"`
	// DeployZoneName

	DeployZoneName *string `json:"DeployZoneName,omitempty" name:"DeployZoneName"`
	// 区域ID

	DeployAreaID *string `json:"DeployAreaID,omitempty" name:"DeployAreaID"`
	// DeployAreaName

	DeployAreaName *string `json:"DeployAreaName,omitempty" name:"DeployAreaName"`
	//  TargetHostID

	TargetHostID *string `json:"TargetHostID,omitempty" name:"TargetHostID"`
	// 允许用户访问的验证方式

	SecCerts []*string `json:"SecCerts,omitempty" name:"SecCerts"`
	// 限流

	InvokeLimit *int64 `json:"InvokeLimit,omitempty" name:"InvokeLimit"`
	// 限流时间

	InvokeLimitInterval *int64 `json:"InvokeLimitInterval,omitempty" name:"InvokeLimitInterval"`
	// 是否允许用户访问

	UserAccessible *bool `json:"UserAccessible,omitempty" name:"UserAccessible"`
	// 签名

	Signature *string `json:"Signature,omitempty" name:"Signature"`
	// 超时

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 目标服务器

	TargetRankHosts []*TargetRankHosts `json:"TargetRankHosts,omitempty" name:"TargetRankHosts"`
	// 目标路径，也称（原始路径）

	TargetPath *string `json:"TargetPath,omitempty" name:"TargetPath"`
	// 版本

	Versions []*VersionsOfService `json:"Versions,omitempty" name:"Versions"`
	// 协议是否支持https

	TargetIsHttps *bool `json:"TargetIsHttps,omitempty" name:"TargetIsHttps"`
	// 协议是否支持http2，这两个需配合使用，如果是https（HTTP/2）的话，二者都为true

	TargetIsHttp2 *bool `json:"TargetIsHttp2,omitempty" name:"TargetIsHttp2"`
	// 选择https后，是否需要https证书检查

	IgnoreTargetCertsError *bool `json:"IgnoreTargetCertsError,omitempty" name:"IgnoreTargetCertsError"`
	// dubbo注册中心ID

	DubboRegisterCenterID *string `json:"DubboRegisterCenterID,omitempty" name:"DubboRegisterCenterID"`
	// dubbo接口全类名

	DubboInterfaceClassName *string `json:"DubboInterfaceClassName,omitempty" name:"DubboInterfaceClassName"`
	// dubbo接口目标方法名

	DubboInterfaceMethodName *string `json:"DubboInterfaceMethodName,omitempty" name:"DubboInterfaceMethodName"`
	// dubbo接口版本号

	DubboInterfaceVersion *string `json:"DubboInterfaceVersion,omitempty" name:"DubboInterfaceVersion"`
	// dubbo接口分组名

	DubboInterfaceGroup *string `json:"DubboInterfaceGroup,omitempty" name:"DubboInterfaceGroup"`
	// dubbo接口令牌

	DubboInterfaceToken *string `json:"DubboInterfaceToken,omitempty" name:"DubboInterfaceToken"`
	// dubbo接口调用协议

	DubboInterfaceProtocol *string `json:"DubboInterfaceProtocol,omitempty" name:"DubboInterfaceProtocol"`
	// 超时配置

	DubboInterfaceTimeout *int64 `json:"DubboInterfaceTimeout,omitempty" name:"DubboInterfaceTimeout"`
	// grpc服务实例，，直连实例hosts，注册中心registerCenter，二选一

	GrpcServiceSource *string `json:"GrpcServiceSource,omitempty" name:"GrpcServiceSource"`
	// grpc服务实例列表，若GrpcServiceSource == hosts的话

	GrpcServerHosts *string `json:"GrpcServerHosts,omitempty" name:"GrpcServerHosts"`
	// grpc服务注册中心，若GrpcServiceSource == registerCenter的话

	GrpcRegisterCenterID *string `json:"GrpcRegisterCenterID,omitempty" name:"GrpcRegisterCenterID"`
	// grpc服务注册的服务名，若GrpcServiceSource == registerCenter的话

	GrpcRegisterServiceName *string `json:"GrpcRegisterServiceName,omitempty" name:"GrpcRegisterServiceName"`
	// grpc文件组ID

	GrpcProtoFilesID *string `json:"GrpcProtoFilesID,omitempty" name:"GrpcProtoFilesID"`
	// grpc定义的包名

	GrpcPackageName *string `json:"GrpcPackageName,omitempty" name:"GrpcPackageName"`
	// grpc定义的服务名

	GrpcServiceName *string `json:"GrpcServiceName,omitempty" name:"GrpcServiceName"`
	// grpc目标方法民

	GrpcMethodName *string `json:"GrpcMethodName,omitempty" name:"GrpcMethodName"`
	// grpc配置的CA内容

	GrpcTlsCa *string `json:"GrpcTlsCa,omitempty" name:"GrpcTlsCa"`
	// grpc配置的Key

	GrpcTlsKey *string `json:"GrpcTlsKey,omitempty" name:"GrpcTlsKey"`
	// grpc配置的Csr

	GrpcTlsCsr *string `json:"GrpcTlsCsr,omitempty" name:"GrpcTlsCsr"`
	// 自定义认证字段

	GrpcMetaMap *string `json:"GrpcMetaMap,omitempty" name:"GrpcMetaMap"`
	// 数据源ID，服务类型为Database时必填

	DataSourceID *string `json:"DataSourceID,omitempty" name:"DataSourceID"`
	// Sql，服务类型为Database时必填

	Sql *string `json:"Sql,omitempty" name:"Sql"`
}

type SubscriptionApproveLogVO struct {

	// AccessRegions

	AccessRegions []*RegionOfSubscription `json:"AccessRegions,omitempty" name:"AccessRegions"`
	// AdvancedRules

	AdvancedRules *string `json:"AdvancedRules,omitempty" name:"AdvancedRules"`
	// ApprovalTime

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// Builtin

	Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CurRenewStatus

	CurRenewStatus *int64 `json:"CurRenewStatus,omitempty" name:"CurRenewStatus"`
	// DeployRegions

	DeployRegions []*RegionOfSubscription `json:"DeployRegions,omitempty" name:"DeployRegions"`
	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// InvokeLimit

	InvokeLimit *int64 `json:"InvokeLimit,omitempty" name:"InvokeLimit"`
	// InvokeLimitInterval

	InvokeLimitInterval *int64 `json:"InvokeLimitInterval,omitempty" name:"InvokeLimitInterval"`
	// IpWhiteList

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// PublisherService

	PublisherService *PublisherService `json:"PublisherService,omitempty" name:"PublisherService"`
	// SignType

	SignType []*string `json:"SignType,omitempty" name:"SignType"`
	// Signature

	Signature *string `json:"Signature,omitempty" name:"Signature"`
	// SubscriberApp

	SubscriberApp *SubscriberApp `json:"SubscriberApp,omitempty" name:"SubscriberApp"`
	// 审批人

	ApprovalUser *string `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 驳回原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type DescribeLogoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// logo id

		ID *string `json:"ID,omitempty" name:"ID"`
		// 文件base64内容

		Logo *string `json:"Logo,omitempty" name:"Logo"`
		// logo文件名称

		LogoName *string `json:"LogoName,omitempty" name:"LogoName"`
		// logo名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 是否展示名称

		NameDisplay *bool `json:"NameDisplay,omitempty" name:"NameDisplay"`
		// 是否是默认logo

		Protected *bool `json:"Protected,omitempty" name:"Protected"`
		// 创建时间

		CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
		// 最后更新时间

		LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
		// 创建人

		CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnSyncPaasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用列表

		Items []*AppOfPagedApps `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnSyncPaasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnSyncPaasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMicroRegisterCenterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMicroRegisterCenterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMicroRegisterCenterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceMarketPageRequest struct {
	*tchttp.BaseRequest

	// 限定分页条目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// true：里约自有服务，false：开发者服务

	Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
	// 机构ID列表

	AgencyIDs []*string `json:"AgencyIDs,omitempty" name:"AgencyIDs"`
	// 应用ID列表

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 服务目录列表

	DocIDs []*string `json:"DocIDs,omitempty" name:"DocIDs"`
	// 服务标签列表

	LabelIDs []*string `json:"LabelIDs,omitempty" name:"LabelIDs"`
	// API目录

	CategoryIDs []*string `json:"CategoryIDs,omitempty" name:"CategoryIDs"`
	// 关键词搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeServiceMarketPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceMarketPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAndProfileDTO struct {

	// 用户ID

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
	// 用户帐号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 是否修改过密码

	NeedChangePassword *bool `json:"NeedChangePassword,omitempty" name:"NeedChangePassword"`
	// 是否保护帐号

	Protected *bool `json:"Protected,omitempty" name:"Protected"`
	// 用户状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// appid

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 证件号码

	CID *string `json:"CID,omitempty" name:"CID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 用户显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 性别

	Gender *string `json:"Gender,omitempty" name:"Gender"`
	// 上次更新时间

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
}

type ApprovedSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApprovedSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteApproveLogsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 应用列表

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// Statuses

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
}

func (r *DescribeSiteApproveLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteApproveLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTlsCertificatesRequest struct {
	*tchttp.BaseRequest

	// 关键词搜索（证书名称）

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 机构ID

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
	// 限定分页条目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序搜索

	Sort *DescribeTlsCertificateSortDTO `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeTlsCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTlsCertificatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AtlasCaptcha struct {

	// 应用 paasID, 默认 rio

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 第三方验证验证码地址

	VerifyAddress *string `json:"VerifyAddress,omitempty" name:"VerifyAddress"`
	// 是否每次登录均出现

	Always *bool `json:"Always,omitempty" name:"Always"`
	// 登录失败连续时长

	Minutes *int64 `json:"Minutes,omitempty" name:"Minutes"`
	// 登录失败次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 启用场景 -- 界面图形验证码出现时机

	EnabledOn []*string `json:"EnabledOn,omitempty" name:"EnabledOn"`
	// 是否启用图形验证码功能

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type DescribeMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Items

		Items []*AgencyMember `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmartGatewayGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站点数量

		SiteNum *int64 `json:"SiteNum,omitempty" name:"SiteNum"`
		// 服务数量

		ServiceNum *int64 `json:"ServiceNum,omitempty" name:"ServiceNum"`
		// SggID

		SggID *string `json:"SggID,omitempty" name:"SggID"`
		// SggName

		SggName *string `json:"SggName,omitempty" name:"SggName"`
		// 新增编辑修改不填

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// BelongZoneID

		BelongZoneID *string `json:"BelongZoneID,omitempty" name:"BelongZoneID"`
		// 新增编辑修改不填

		BelongZoneName *string `json:"BelongZoneName,omitempty" name:"BelongZoneName"`
		// BelongAreaID

		BelongAreaID *string `json:"BelongAreaID,omitempty" name:"BelongAreaID"`
		// 新增编辑修改不填

		BelongAreaName *string `json:"BelongAreaName,omitempty" name:"BelongAreaName"`
		// Domain

		Domain []*string `json:"Domain,omitempty" name:"Domain"`
		// 查找客户端真实ip时过滤内部ip

		IPs []*string `json:"IPs,omitempty" name:"IPs"`
		// RoleID

		RoleID *string `json:"RoleID,omitempty" name:"RoleID"`
		// 新增编辑修改不填

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 127.0.0.1:8006",

		SgAdminBaseUrl *string `json:"SgAdminBaseUrl,omitempty" name:"SgAdminBaseUrl"`
		// 新增编辑修改不填

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 新增编辑修改不填

		LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmartGatewayGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSmartGatewayGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertRuleRequest struct {
	*tchttp.BaseRequest

	// 告警规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 告警类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 告警触发条件

	Triggers []*Triggers `json:"Triggers,omitempty" name:"Triggers"`
	// 指标计算规则

	Evaluate *Evaluate `json:"Evaluate,omitempty" name:"Evaluate"`
	// 告警周期

	NotifyInterval *string `json:"NotifyInterval,omitempty" name:"NotifyInterval"`
	// 告警通知渠道

	Receivers []*Receivers `json:"Receivers,omitempty" name:"Receivers"`
	// 告警启用状态

	State *string `json:"State,omitempty" name:"State"`
}

func (r *ModifyAlertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConnectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConnectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConnectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcsRoleVO struct {

	// 唯一标识

	Code *string `json:"Code,omitempty" name:"Code"`
	// 角色名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色成员数量

	MemberCount *int64 `json:"MemberCount,omitempty" name:"MemberCount"`
	// 作用范围

	Scope *string `json:"Scope,omitempty" name:"Scope"`
}

type AdvanceSetting struct {

	//  是否开启异常响应设置

	ExceptionStatus *bool `json:"ExceptionStatus,omitempty" name:"ExceptionStatus"`
	//  502异常

	GatewayErrorContent *string `json:"GatewayErrorContent,omitempty" name:"GatewayErrorContent"`
	//  不在访问范围之内

	PermissionDeniedMessage *string `json:"PermissionDeniedMessage,omitempty" name:"PermissionDeniedMessage"`
	//  是否自定义请求响应

	FilterStatus *bool `json:"FilterStatus,omitempty" name:"FilterStatus"`
	//  目标主机头

	TargetHostHeader *string `json:"TargetHostHeader,omitempty" name:"TargetHostHeader"`
	//  请求头

	RequestHeadersFilter *string `json:"RequestHeadersFilter,omitempty" name:"RequestHeadersFilter"`
	//  请求内容,

	RequestBodyFilter *string `json:"RequestBodyFilter,omitempty" name:"RequestBodyFilter"`
	//  请求参数

	RequestParameterProcesser *string `json:"RequestParameterProcesser,omitempty" name:"RequestParameterProcesser"`
	//  响应头

	ResponseHeadersFilter *string `json:"ResponseHeadersFilter,omitempty" name:"ResponseHeadersFilter"`
	//  响应内容

	ResponseBodyFilter *string `json:"ResponseBodyFilter,omitempty" name:"ResponseBodyFilter"`
	//  响应参数

	ResponseParameterProcesser *string `json:"ResponseParameterProcesser,omitempty" name:"ResponseParameterProcesser"`
	//  是否开启健康检查

	HealthCheckStatus *bool `json:"HealthCheckStatus,omitempty" name:"HealthCheckStatus"`
	//  健康检查地址

	HealthCheckPath *string `json:"HealthCheckPath,omitempty" name:"HealthCheckPath"`
	//  健康检查超时

	HealthCheckTimeout *uint64 `json:"HealthCheckTimeout,omitempty" name:"HealthCheckTimeout"`
	//  健康检查码

	ValidHealthCheckStatusCode []*uint64 `json:"ValidHealthCheckStatusCode,omitempty" name:"ValidHealthCheckStatusCode"`
	// 是否开启站点水印

	SiteWaterStatus *bool `json:"SiteWaterStatus,omitempty" name:"SiteWaterStatus"`
	// SiteWaterConfig

	SiteWaterConfig *SiteWaterConfig `json:"SiteWaterConfig,omitempty" name:"SiteWaterConfig"`
	// TemplateAdvanceSetList

	TemplateAdvanceSetList []*TemplateAdvanceSetList `json:"TemplateAdvanceSetList,omitempty" name:"TemplateAdvanceSetList"`
}

type PubAreaConfig struct {

	// 0:网络区域统一配置，1:各个网络区域单独配置

	PubAreaConfigType *int64 `json:"PubAreaConfigType,omitempty" name:"PubAreaConfigType"`
	// PubArea

	PubArea []*PubArea `json:"PubArea,omitempty" name:"PubArea"`
}

type DescribeSmartGateRaidRoutingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*SmartGateRaidRouting `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmartGateRaidRoutingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSmartGateRaidRoutingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncSmartGatewayGroupRequest struct {
	*tchttp.BaseRequest

	// 节点阵列 ID

	SggID *string `json:"SggID,omitempty" name:"SggID"`
}

func (r *SyncSmartGatewayGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncSmartGatewayGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SiteApproveVo struct {

	// 所属应用

	App *AppOfSiteApproveVo `json:"App,omitempty" name:"App"`
	// 申请时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 申请人

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// 站点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 访问路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 站点Id

	SiteID *string `json:"SiteID,omitempty" name:"SiteID"`
}

type DeleteAgencyManagerRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserIDs []*string `json:"UserIDs,omitempty" name:"UserIDs"`
	// 机构id

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

func (r *DeleteAgencyManagerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgencyManagerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataSourceTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDataSourceTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataSourceTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemOfSystemAppInfoVO struct {

	// 系统id

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
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

type ModifySmartGateRaidRoutingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySmartGateRaidRoutingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySmartGateRaidRoutingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckProfileRequest struct {

	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 证件号码

	CID *string `json:"CID,omitempty" name:"CID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 用户显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
}

type Triggers struct {

	// Metric

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 范围值

	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
	// Condition

	Condition *string `json:"Condition,omitempty" name:"Condition"`
}

type CreateSmartGatewayGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSmartGatewayGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSmartGatewayGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcsRoleRequest struct {
	*tchttp.BaseRequest

	// 唯一标识

	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *DescribeAcsRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcsRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewRequest struct {
	*tchttp.BaseRequest

	// 应用id

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
}

func (r *DescribeOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRangeMetricsRequest struct {
	*tchttp.BaseRequest

	// PromQL语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 开始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 时间间隔

	Step *string `json:"Step,omitempty" name:"Step"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
}

func (r *DescribeRangeMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRangeMetricsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrafanaAlertRecordVO struct {

	// AlertID

	AlertID *int64 `json:"AlertID,omitempty" name:"AlertID"`
	// TimeStart

	TimeStart *int64 `json:"TimeStart,omitempty" name:"TimeStart"`
	// TimeEnd

	TimeEnd *int64 `json:"TimeEnd,omitempty" name:"TimeEnd"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
	// Metadata

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	// State

	State *string `json:"State,omitempty" name:"State"`
}

type DownloadWebConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadWebConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadWebConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroSyncJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *MicroSyncJob `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroSyncJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySmartGatewayGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySmartGatewayGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySmartGatewayGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSubscriptionBatchRequest struct {
	*tchttp.BaseRequest

	// 系统 ID 列表

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
	// 系统 ID 列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *RejectedSubscriptionBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSubscriptionBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMicroProtoGroupRequest struct {
	*tchttp.BaseRequest

	// 文件组 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 文件组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件组描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *CreateMicroProtoGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMicroProtoGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTSFServiceRequest struct {
	*tchttp.BaseRequest

	// 租户 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 命名空间 ID

	NamespaceID *string `json:"NamespaceID,omitempty" name:"NamespaceID"`
	// 微服务实例 ID

	MicroserviceID *string `json:"MicroserviceID,omitempty" name:"MicroserviceID"`
	// 参考附录 《对象说明》

	Apis []*MicroApi `json:"Apis,omitempty" name:"Apis"`
}

func (r *CreateTSFServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSystemWarnMsgRequest struct {
	*tchttp.BaseRequest

	// 系统ID

	SysID *string `json:"SysID,omitempty" name:"SysID"`
}

func (r *DeleteSystemWarnMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSystemWarnMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcsRoleListRequest struct {
	*tchttp.BaseRequest

	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 角色名. code

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 机构id下的角色

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

func (r *DescribeAcsRoleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcsRoleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDocumentGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data []*DocGroupVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDocumentGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDocumentGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAgencyMemberRequest struct {
	*tchttp.BaseRequest

	// 机构ID

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

func (r *ExportAgencyMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAgencyMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDataSourceRequest struct {
	*tchttp.BaseRequest

	// 数据源 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 数据源连接串

	Dsn *string `json:"Dsn,omitempty" name:"Dsn"`
	// 最大返回记录数

	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
}

func (r *ModifyDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveMemberFromAcsRoleRequest struct {
	*tchttp.BaseRequest

	// 0:普通角色成员 10:机构负责人 11:系统负责人 12: 应用负责人

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 空：角色关系 非空：负责人关系

	RelationshipBelongTo *string `json:"RelationshipBelongTo,omitempty" name:"RelationshipBelongTo"`
	// 编码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 批量添加用户

	AccountIDs []*string `json:"AccountIDs,omitempty" name:"AccountIDs"`
}

func (r *RemoveMemberFromAcsRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveMemberFromAcsRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedSiteBatchRequest struct {
	*tchttp.BaseRequest

	// 系统 ID 列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *ApprovedSiteBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSiteBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *Subscription `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroProtoFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMicroProtoFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroProtoFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSmartGateRaidRoutingRequest struct {
	*tchttp.BaseRequest

	// 节点阵列路由 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteSmartGateRaidRoutingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSmartGateRaidRoutingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveMemberFromAcsRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveMemberFromAcsRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveMemberFromAcsRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroRegisterCenterRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteMicroRegisterCenterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroRegisterCenterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *Subscription `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneAreasRequest struct {
	*tchttp.BaseRequest

	// 类型 [site/service]

	Role *string `json:"Role,omitempty" name:"Role"`
}

func (r *DescribeZoneAreasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneAreasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCategoryRequest struct {
	*tchttp.BaseRequest

	// ⽬录id

	CategoryID *string `json:"CategoryID,omitempty" name:"CategoryID"`
	// ⽬录名称

	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`
}

func (r *ModifyCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppOfSystemAppInfoVO struct {

	// 应用id

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type MicroTenant struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 租户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 微服务类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Tsf

	Tsf *Tsf `json:"Tsf,omitempty" name:"Tsf"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type DescribeMicroProtoFileRequest struct {
	*tchttp.BaseRequest

	// 文件组 ID

	ProtoGroupId *string `json:"ProtoGroupId,omitempty" name:"ProtoGroupId"`
	// 文件名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeMicroProtoFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroProtoFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessRange struct {

	// 访问范围

	Accounts []*AccountsOfAccessRange `json:"Accounts,omitempty" name:"Accounts"`
}

type AddMemberToAcsRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMemberToAcsRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMemberToAcsRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgencyConfigRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 条数

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAgencyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgencyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataSourceTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录 《对象说明》

		Items []*DataSourceType `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataSourceTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataSourceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDocumentFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 文件内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDocumentFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDocumentFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportAgencyMemberRequest struct {
	*tchttp.BaseRequest

	// FileName

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// Content

	Content *string `json:"Content,omitempty" name:"Content"`
	// AgencyID

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
}

func (r *ImportAgencyMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportAgencyMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppOfSiteApproveVo struct {

	// 应用ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 应用名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type SubscriptionRenewVO struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// SubscriptionID

	SubscriptionID *string `json:"SubscriptionID,omitempty" name:"SubscriptionID"`
	// Subscription

	Subscription *SubscriptionVO `json:"Subscription,omitempty" name:"Subscription"`
	// ExpireTime

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// UpdateTime

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// UpdateUser

	UpdateUser *CreateUser `json:"UpdateUser,omitempty" name:"UpdateUser"`
}

type DescribeAlertHistoryFileRequest struct {
	*tchttp.BaseRequest

	// 时间起点

	From *string `json:"From,omitempty" name:"From"`
	// 时间终点

	To *string `json:"To,omitempty" name:"To"`
	// 限制返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlertHistoryFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertHistoryFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroSyncJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*MicroSyncJob `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroSyncJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroSyncJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*TSFNamespace `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTSFNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvanceTemplateBtnFunctionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AdvanceTemplateBtnFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AdvanceTemplateBtnFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsOfPubAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *string `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*Subscription `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionsOfPubAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionsOfPubAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthInfoRequest struct {
	*tchttp.BaseRequest

	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAuthInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuthInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvanceRouting struct {

	//  规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	//  访问路径，（前缀：上层的PubPath/）

	AccessPath *string `json:"AccessPath,omitempty" name:"AccessPath"`
	//  是否需要登录

	NeedLogin *bool `json:"NeedLogin,omitempty" name:"NeedLogin"`
	//  访问范围（候补）,如果NeedLogin = false，此项取消

	AccessRange *AccessRange `json:"AccessRange,omitempty" name:"AccessRange"`
	//  0:与目标服务器保持一致（一致在提交的时候，也要沿用填充），1:自定义

	TargetRankHostType *int64 `json:"TargetRankHostType,omitempty" name:"TargetRankHostType"`
	// TargetRankHosts

	TargetRankHosts []*TargetRankHosts `json:"TargetRankHosts,omitempty" name:"TargetRankHosts"`
	//  目标路径。（前缀：上层的TargetPath/）

	TargetPath *string `json:"TargetPath,omitempty" name:"TargetPath"`
}

type DescribeDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *DataSource `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*Zone `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevokeSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 订阅 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *RevokeSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RevokeSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddVendorMemberByIDsRequest struct {
	*tchttp.BaseRequest

	// 供应商ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 用户ID列表

	UserIDs []*string `json:"UserIDs,omitempty" name:"UserIDs"`
}

func (r *AddVendorMemberByIDsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVendorMemberByIDsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedSystemRequest struct {
	*tchttp.BaseRequest

	// 系统 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *ApprovedSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMembersRequest struct {
	*tchttp.BaseRequest

	// 机构id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteArchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *SiteVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSiteArchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteArchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTipsRequest struct {
	*tchttp.BaseRequest

	// 提示语类型 1: 登录模块 tip

	TipType *int64 `json:"TipType,omitempty" name:"TipType"`
	// 提示语子类型 10: 登录模块--帐号密码登录子模块 11: 登录模块--短信验证码登录 12: 登录模块--双因子登录

	TipSubType *int64 `json:"TipSubType,omitempty" name:"TipSubType"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTipsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTipsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadAccTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadAccTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadAccTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppOfSiteVO struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type AddAgencyManagerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAgencyManagerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAgencyManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSmartGateRaidRoutingRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// 新增编辑修改不填

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 新增编辑修改不填

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// StartSgg

	StartSgg *Summary `json:"StartSgg,omitempty" name:"StartSgg"`
	// ForwardSgg

	ForwardSgg []*string `json:"ForwardSgg,omitempty" name:"ForwardSgg"`
	// TargetSgg

	TargetSgg *Summary `json:"TargetSgg,omitempty" name:"TargetSgg"`
}

func (r *CreateSmartGateRaidRoutingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSmartGateRaidRoutingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMenuAllListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 顶部菜单

		TopMenus []*Menu `json:"TopMenus,omitempty" name:"TopMenus"`
		// 左侧菜单

		LeftMenus []*Menu `json:"LeftMenus,omitempty" name:"LeftMenus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMenuAllListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMenuAllListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存档列表 参考附录 《对象说明》

		Snapshots []*Snapshot `json:"Snapshots,omitempty" name:"Snapshots"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadRequest struct {
	*tchttp.BaseRequest

	// FileName

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// Content

	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *UploadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifySmsCodeAndModifyPhoneRequest struct {
	*tchttp.BaseRequest

	// 短信验证码验证

	VerifyRequest *VerifyRequest `json:"VerifyRequest,omitempty" name:"VerifyRequest"`
	// 更新手机号入参

	UpdateRequest *UpdateRequest `json:"UpdateRequest,omitempty" name:"UpdateRequest"`
}

func (r *VerifySmsCodeAndModifyPhoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifySmsCodeAndModifyPhoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgencyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 条数

		Limit *int64 `json:"Limit,omitempty" name:"Limit"`
		// 偏移量

		Offset *int64 `json:"Offset,omitempty" name:"Offset"`
		// AgencyConfig

		Items []*AgencyConfig `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgencyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgencyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionRenewApproveLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批列表，参考附录 《对象说明》

		Items []*string `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionRenewApproveLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionRenewApproveLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectedServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Captcha struct {

	// 图形验证码ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 图形验证码URL

	Pic *string `json:"Pic,omitempty" name:"Pic"`
}

type DeleteConnectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConnectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConnectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMicroTenantRequest struct {
	*tchttp.BaseRequest

	// 租户ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 参考附录 《对象说明》

	Tsf *Tsf `json:"Tsf,omitempty" name:"Tsf"`
}

func (r *ModifyMicroTenantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMicroTenantRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedSystemBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败数量

		FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
		// 成功数量

		SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApprovedSystemBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSystemBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAcsRoleUniqueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否有效

		IsValid *bool `json:"IsValid,omitempty" name:"IsValid"`
		// 错误信息

		ValidMsg *string `json:"ValidMsg,omitempty" name:"ValidMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAcsRoleUniqueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAcsRoleUniqueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SiteRequestLogVO struct {

	// 请求时间

	Datestamp *string `json:"Datestamp,omitempty" name:"Datestamp"`
	// 索引后缀

	SourcePrefix *string `json:"SourcePrefix,omitempty" name:"SourcePrefix"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 站点ID

	SiteID *string `json:"SiteID,omitempty" name:"SiteID"`
	// 请求域名（host）

	RequestHost *string `json:"RequestHost,omitempty" name:"RequestHost"`
	// 请求路径

	RequestPath *string `json:"RequestPath,omitempty" name:"RequestPath"`
	// 源IP

	OriginIP *string `json:"OriginIP,omitempty" name:"OriginIP"`
	// userID

	UserID *string `json:"UserID,omitempty" name:"UserID"`
	// 耗时

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// 目标服务响应码

	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`
	// 错误信息

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// 日志唯一标识

	ReqID *string `json:"ReqID,omitempty" name:"ReqID"`
}

type DeleteLabelRequest struct {
	*tchttp.BaseRequest

	// 标签ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroRegisterCentersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*Register `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroRegisterCentersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroRegisterCentersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVendorMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVendorMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVendorMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckServiceUniqueRequest struct {
	*tchttp.BaseRequest

	// 定值查询，不变

	Object *string `json:"Object,omitempty" name:"Object"`
	// 过滤条件

	FilterT *string `json:"FilterT,omitempty" name:"FilterT"`
}

func (r *CheckServiceUniqueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckServiceUniqueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAcsRoleRequest struct {
	*tchttp.BaseRequest

	// 编码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 角色名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色属于机构/全局/内部

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 角色范围描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 属于机构, 为空表示全局机构

	AgencyID *string `json:"AgencyID,omitempty" name:"AgencyID"`
	// 权限项

	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
	// AccountIDs

	AccountIDs []*string `json:"AccountIDs,omitempty" name:"AccountIDs"`
}

func (r *CreateAcsRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAcsRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVendorRequest struct {
	*tchttp.BaseRequest

	// 供应商ID列表

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
}

func (r *DeleteVendorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVendorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppTop10Request struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAppTop10Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppTop10Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceApproveLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批列表，参考附录 《对象说明》

		Items []*ServiceApproveLogVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceApproveLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceApproveLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableRequest struct {
	*tchttp.BaseRequest

	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *EnableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Vendor struct {

	// 供应商ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 供应商名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 供应商成立时间

	Badge *string `json:"Badge,omitempty" name:"Badge"`
	// 供应商成立时间

	Description *string `json:"Description,omitempty" name:"Description"`
	// 供应商成立时间

	Founded *string `json:"Founded,omitempty" name:"Founded"`
	// 供应商地址

	Addr *string `json:"Addr,omitempty" name:"Addr"`
	// 供应商管理员列表

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 供应商资质证明

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// 供应商成员数量

	MemberCount *uint64 `json:"MemberCount,omitempty" name:"MemberCount"`
}

type DeleteSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySmartGateRaidRoutingRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// 新增编辑修改不填

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 新增编辑修改不填

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// StartSgg

	StartSgg *Summary `json:"StartSgg,omitempty" name:"StartSgg"`
	// ForwardSgg

	ForwardSgg []*string `json:"ForwardSgg,omitempty" name:"ForwardSgg"`
	// TargetSgg

	TargetSgg *Summary `json:"TargetSgg,omitempty" name:"TargetSgg"`
}

func (r *ModifySmartGateRaidRoutingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySmartGateRaidRoutingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectRenewSubscriptionBatchRequest struct {
	*tchttp.BaseRequest

	// 订阅IDs

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
	// 订阅IDs

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *RejectRenewSubscriptionBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectRenewSubscriptionBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMicroTenantResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户id

		ID *string `json:"ID,omitempty" name:"ID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMicroTenantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMicroTenantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConnectorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConnectorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConnectorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*Area `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAreasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcsRolePermissionRequest struct {
	*tchttp.BaseRequest

	// 编码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAcsRolePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcsRolePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLabelsRequest struct {
	*tchttp.BaseRequest

	// 限定分⻚条⽬

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分⻚

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 关键词查询

	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
	// 排序

	Sort *LabelSort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadLicenseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadLicenseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscriptionApproveVO struct {

	// AccessRegions

	AccessRegions []*RegionOfSubscription `json:"AccessRegions,omitempty" name:"AccessRegions"`
	// AdvancedRules

	AdvancedRules *string `json:"AdvancedRules,omitempty" name:"AdvancedRules"`
	// ApprovalTime

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// Builtin

	Builtin *bool `json:"Builtin,omitempty" name:"Builtin"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CurRenewStatus

	CurRenewStatus *int64 `json:"CurRenewStatus,omitempty" name:"CurRenewStatus"`
	// DeployRegions

	DeployRegions []*RegionOfSubscription `json:"DeployRegions,omitempty" name:"DeployRegions"`
	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// InvokeLimit

	InvokeLimit *int64 `json:"InvokeLimit,omitempty" name:"InvokeLimit"`
	// InvokeLimitInterval

	InvokeLimitInterval *int64 `json:"InvokeLimitInterval,omitempty" name:"InvokeLimitInterval"`
	// IpWhiteList

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// PublisherService

	PublisherService *PublisherService `json:"PublisherService,omitempty" name:"PublisherService"`
	// SignType

	SignType []*string `json:"SignType,omitempty" name:"SignType"`
	// Signature

	Signature *string `json:"Signature,omitempty" name:"Signature"`
	// SubscriberApp

	SubscriberApp *SubscriberApp `json:"SubscriberApp,omitempty" name:"SubscriberApp"`
}

type DescribeDocumentListRequest struct {
	*tchttp.BaseRequest

	// 分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 文档名称

	DocName *string `json:"DocName,omitempty" name:"DocName"`
}

func (r *DescribeDocumentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDocumentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteTop10Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站点列表，参考附录 《对象说明》

		Items []*SiteVO `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSiteTop10Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteTop10Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedServiceRequest struct {
	*tchttp.BaseRequest

	// 系统 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *RejectedServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseData struct {

	// 客户名称

	Customer *string `json:"Customer,omitempty" name:"Customer"`
	// 序列号

	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`
	// 证书类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// Products

	Products []*Products `json:"Products,omitempty" name:"Products"`
}

type CreateCategoryRequest struct {
	*tchttp.BaseRequest

	// ⽬录名称

	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`
}

func (r *CreateCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelSort struct {

	// 生序1，降序-1

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 生序1，降序-1

	LabelName *int64 `json:"LabelName,omitempty" name:"LabelName"`
}

type CheckSiteUniqueRequest struct {
	*tchttp.BaseRequest

	// 定值查询，不变

	Object *string `json:"Object,omitempty" name:"Object"`
	// 过滤条件

	FilterT *string `json:"FilterT,omitempty" name:"FilterT"`
}

func (r *CheckSiteUniqueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSiteUniqueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 帐号ID

		AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
		// 密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataSourcesRequest struct {
	*tchttp.BaseRequest

	// 数据源 ID

	ID []*string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteDataSourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataSourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTlsCertificateRequest struct {
	*tchttp.BaseRequest

	// 证书ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteTlsCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTlsCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginLog struct {

	// 日志id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 账号id

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
	// 账号名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 操作类型

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// ip地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 登录类型

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 时间

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// ua

	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
	// UserGroup

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// VendorAgency

	VendorAgency *int64 `json:"VendorAgency,omitempty" name:"VendorAgency"`
	// Errcode

	Errcode *string `json:"Errcode,omitempty" name:"Errcode"`
	// ErrMsg

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// Result

	Result *string `json:"Result,omitempty" name:"Result"`
}

type CreateSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *AccountAndProfileDTO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcsRoleMembersRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Code

	Code *string `json:"Code,omitempty" name:"Code"`
	// Types

	Types []*int64 `json:"Types,omitempty" name:"Types"`
	// RelationshipBelongTo

	RelationshipBelongTo *string `json:"RelationshipBelongTo,omitempty" name:"RelationshipBelongTo"`
}

func (r *DescribeAcsRoleMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcsRoleMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Receiver struct {

	// Channel

	Channel *string `json:"Channel,omitempty" name:"Channel"`
	// qyapi.weixin.qq.com/cgi-bin/webhook/send?key=8c1e844b-7bb1-4d9b-91ba-75d83a47db8c"

	Address *string `json:"Address,omitempty" name:"Address"`
}

type DescribeTSFGatewayInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 网关实例列表

		Items []*string `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTSFGatewayInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFGatewayInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVendorMemberListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 供应商ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeVendorMemberListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVendorMemberListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// path

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTSFMicroserviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTSFMicroserviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTSFMicroserviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSmartGatewayGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSmartGatewayGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSmartGatewayGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFRegionListsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTSFRegionListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFRegionListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDownloadOperationLogRequest struct {
	*tchttp.BaseRequest

	// 账号名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 模块名称

	OpModule *string `json:"OpModule,omitempty" name:"OpModule"`
	// 关键词

	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
	// 开始时间 秒级时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间 秒级时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyDownloadOperationLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDownloadOperationLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VendorUserVO struct {

	// 用户账号id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 证件号码

	CID *string `json:"CID,omitempty" name:"CID"`
	// 帐号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 姓名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 性别

	Gender *string `json:"Gender,omitempty" name:"Gender"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 绑定时间

	BindTime *uint64 `json:"BindTime,omitempty" name:"BindTime"`
	// 供应商ID

	VendorID *string `json:"VendorID,omitempty" name:"VendorID"`
}

type DownloadVendorImportResultRequest struct {
	*tchttp.BaseRequest

	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *DownloadVendorImportResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadVendorImportResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevokeSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevokeSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RevokeSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteSortDTO struct {

	// 1升序，-1降序

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 1升序，-1降序

	Name *int64 `json:"Name,omitempty" name:"Name"`
	// 1升序，-1降序

	SiteName *int64 `json:"SiteName,omitempty" name:"SiteName"`
}

type CheckAgencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否合法

		IsValid *bool `json:"IsValid,omitempty" name:"IsValid"`
		// 不合法时的提示信息

		ValidMsg *string `json:"ValidMsg,omitempty" name:"ValidMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAgencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAgencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFMicroservicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*TSFMicroservice `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTSFMicroservicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFMicroservicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadSystemDataRequest struct {
	*tchttp.BaseRequest
}

func (r *DownloadSystemDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadSystemDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlertRuleRequest struct {
	*tchttp.BaseRequest

	// 告警规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteAlertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnFreezeRequest struct {
	*tchttp.BaseRequest

	// 服务ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *UnFreezeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnFreezeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckServiceUniqueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *CheckSiteUniqueResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckServiceUniqueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckServiceUniqueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginPolicyRequest struct {
	*tchttp.BaseRequest

	// 应用 paasID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
}

func (r *DescribeLoginPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadServiceDataRequest struct {
	*tchttp.BaseRequest
}

func (r *DownloadServiceDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadServiceDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileVO struct {

	// 文件路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 文件名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteMicroProtoGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMicroProtoGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroProtoGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *LabelVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteMicroSyncJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteMicroSyncJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteMicroSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteRequest struct {
	*tchttp.BaseRequest

	// 0表示草稿，1表示创建

	DraftFlag *int64 `json:"DraftFlag,omitempty" name:"DraftFlag"`
	// 站点ID

	SiteID *string `json:"SiteID,omitempty" name:"SiteID"`
	// paasID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 站点名称

	SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
	// 公开路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 负载方式：false，随机负载，true：会话保持

	KeepSession *bool `json:"KeepSession,omitempty" name:"KeepSession"`
	// 分区配置信息

	Regions []*RegionsOfSite `json:"Regions,omitempty" name:"Regions"`
	// AdvanceSetting

	AdvanceSetting *AdvanceSetting `json:"AdvanceSetting,omitempty" name:"AdvanceSetting"`
}

func (r *CreateSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SiteVO struct {

	// 站点ID

	SiteID *string `json:"SiteID,omitempty" name:"SiteID"`
	// paasID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 站点名称

	SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
	// 公开路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 负载方式：false，随机负载，true：会话保持

	KeepSession *bool `json:"KeepSession,omitempty" name:"KeepSession"`
	// 归档ID

	SiteArchID *string `json:"SiteArchID,omitempty" name:"SiteArchID"`
	// 分区配置信息

	Regions []*RegionsOfSite `json:"Regions,omitempty" name:"Regions"`
	// AdvanceSetting

	AdvanceSetting *AdvanceSetting `json:"AdvanceSetting,omitempty" name:"AdvanceSetting"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Sys

	Sys *SysOfSiteVO `json:"Sys,omitempty" name:"Sys"`
	// App

	App *AppOfSiteVO `json:"App,omitempty" name:"App"`
}

type DescribeAlertHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data []*GrafanaAlertRecordVO `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySqlTemplateRequest struct {
	*tchttp.BaseRequest

	// 服务 id

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 数据源 id

	DataSourceID *string `json:"DataSourceID,omitempty" name:"DataSourceID"`
	// sql 模板

	Sql *string `json:"Sql,omitempty" name:"Sql"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *ModifySqlTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySqlTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSystemBatchRequest struct {
	*tchttp.BaseRequest

	// 系统 ID 列表

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
	// 系统 ID 列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *RejectedSystemBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSystemBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AtlasLockSetting struct {

	// 应用 paasID, 默认 rio

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 重试次数 -- 界面错误次数上限

	LockRetryTimes *int64 `json:"LockRetryTimes,omitempty" name:"LockRetryTimes"`
	// 重试窗口 -- 界面时间段

	LockRetryPeriod *int64 `json:"LockRetryPeriod,omitempty" name:"LockRetryPeriod"`
	// 锁定时长(分钟) -- 界面自动解锁时间

	LockPeriod *int64 `json:"LockPeriod,omitempty" name:"LockPeriod"`
	// 是否启用锁定功能

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type DeleteAgencyRequest struct {
	*tchttp.BaseRequest

	// 机构ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteAgencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataSourcesRequest struct {
	*tchttp.BaseRequest

	// 应用 ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 关键词, 根据ID/Name模糊搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeDataSourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataSourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteApprovesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统审批列表，参考附录 《对象说明》

		Items []*SiteApproveVo `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSiteApprovesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteApprovesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDisableAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDisableAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisableAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SmartGateRaidRouting struct {

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// 新增编辑修改不填

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 新增编辑修改不填

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// StartSgg

	StartSgg *Summary `json:"StartSgg,omitempty" name:"StartSgg"`
	// ForwardSgg

	ForwardSgg []*string `json:"ForwardSgg,omitempty" name:"ForwardSgg"`
	// TargetSgg

	TargetSgg *Summary `json:"TargetSgg,omitempty" name:"TargetSgg"`
}

type DeleteLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSFGatewayApiGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 参考附录 《对象说明》

		Items []*TSFApiGroup `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTSFGatewayApiGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSFGatewayApiGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadServiceDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadServiceDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadServiceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnFreezeSiteRequest struct {
	*tchttp.BaseRequest

	// 站点ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *UnFreezeSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnFreezeSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParentPermission struct {

	// 父权限名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 父权限编码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 子权限

	Child []*AcsRolePermission `json:"Child,omitempty" name:"Child"`
}

type DescribeProfileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参考附录《对象说明》

		Items []*UserProfileWithExist `json:"Items,omitempty" name:"Items"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProfileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadAgencyImportResultRequest struct {
	*tchttp.BaseRequest

	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *DownloadAgencyImportResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadAgencyImportResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceSubscriptionVO struct {

	// API标识

	SvcID *string `json:"SvcID,omitempty" name:"SvcID"`
	// 所属应用

	App *Summary `json:"App,omitempty" name:"App"`
	// 所属系统

	Sys *Summary `json:"Sys,omitempty" name:"Sys"`
	// API名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// API描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// 服务可见性，true公开，false私有

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 状态（草稿0，待审批1，已审批2，禁用5，冻结6）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// jiekl.dizhi",

	DocUrl *string `json:"DocUrl,omitempty" name:"DocUrl"`
	// 创建用户，下同

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ApprovalUser

	ApprovalUser *CreateUser `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// 目标服务类型

	TargetServiceType *string `json:"TargetServiceType,omitempty" name:"TargetServiceType"`
	// ApprovalTime

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 授权app数

	SubscriptionCount *int64 `json:"SubscriptionCount,omitempty" name:"SubscriptionCount"`
	// 归档ID

	ServiceArchID *string `json:"ServiceArchID,omitempty" name:"ServiceArchID"`
	// -1代表未订阅，其他值代表已订阅，值为订阅的状态

	SubscribeStatus *string `json:"SubscribeStatus,omitempty" name:"SubscribeStatus"`
	// 标签列表

	LabelNames []*string `json:"LabelNames,omitempty" name:"LabelNames"`
	// 服务目录列表

	CategoryNames []*string `json:"CategoryNames,omitempty" name:"CategoryNames"`
}

type CreateMicroTenantRequest struct {
	*tchttp.BaseRequest

	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 租户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 微服务类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Tsf

	Tsf *Tsf `json:"Tsf,omitempty" name:"Tsf"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *CreateMicroTenantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMicroTenantRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResetLeftMenuRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyResetLeftMenuRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResetLeftMenuRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadAccResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入结果文件下载路径

		Path *string `json:"Path,omitempty" name:"Path"`
		// 导入结果码 "0" 成功，其他失败

		Code *string `json:"Code,omitempty" name:"Code"`
		// 导入结果信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadAccResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadAccResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateServicePubPathRequest struct {
	*tchttp.BaseRequest

	// 服务公开路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
}

func (r *ValidateServicePubPathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateServicePubPathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateServicePubPathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *ValidateSitePubPathResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateServicePubPathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateServicePubPathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Products struct {

	// ProductCode

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// AuthItems

	AuthItems []*string `json:"AuthItems,omitempty" name:"AuthItems"`
}

type DescribeOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用id

		PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
		// 系统名称

		SysName *string `json:"SysName,omitempty" name:"SysName"`
		// 应用名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 站点数量

		SiteNum *int64 `json:"SiteNum,omitempty" name:"SiteNum"`
		// API数量

		ServiceNum *int64 `json:"ServiceNum,omitempty" name:"ServiceNum"`
		// 订阅API数量

		SubscriptionNum *int64 `json:"SubscriptionNum,omitempty" name:"SubscriptionNum"`
		// API授权申请数量

		ServiceApplicationNum *int64 `json:"ServiceApplicationNum,omitempty" name:"ServiceApplicationNum"`
		// 历史授权记录数量

		HistoryAuthorizeNum *int64 `json:"HistoryAuthorizeNum,omitempty" name:"HistoryAuthorizeNum"`
		// 自定义数据源数量

		DataSourceNum *int64 `json:"DataSourceNum,omitempty" name:"DataSourceNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteApprovesRequest struct {
	*tchttp.BaseRequest

	// 应用id列表

	AppIDs []*string `json:"AppIDs,omitempty" name:"AppIDs"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索字段

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// PaasIDs

	PaasIDs []*string `json:"PaasIDs,omitempty" name:"PaasIDs"`
}

func (r *DescribeSiteApprovesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteApprovesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountInfo struct {

	// 用户帐号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 是否保护帐号

	Protected *bool `json:"Protected,omitempty" name:"Protected"`
	// 用户显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 用户ID

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
}

type DownloadImportTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadImportTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadImportTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest

	// ZoneID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 新增编辑修改不填

	DefaultZone *bool `json:"DefaultZone,omitempty" name:"DefaultZone"`
	// 新增编辑修改不填

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 新增编辑修改不填

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 新增编辑修改不填

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *ModifyZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMemberRequest struct {
	*tchttp.BaseRequest

	// 用户id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 机构ID

	AgencyId *string `json:"AgencyId,omitempty" name:"AgencyId"`
	// 头像地址

	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
	// 证件号

	CID *string `json:"CID,omitempty" name:"CID"`
	// 证件类型

	CType *string `json:"CType,omitempty" name:"CType"`
	// 显示名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 性别

	Gender *string `json:"Gender,omitempty" name:"Gender"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 手机区号

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
}

func (r *UpdateMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgencyManagerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAgencyManagerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgencyManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaRequest struct {
	*tchttp.BaseRequest

	// 区域 ID

	AreaID *string `json:"AreaID,omitempty" name:"AreaID"`
}

func (r *DescribeAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroRegisterCentersRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Keyword

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeMicroRegisterCentersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMicroRegisterCentersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadImportResultRequest struct {
	*tchttp.BaseRequest

	// swagger导入的报告文件路径

	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *DownloadImportResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadImportResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站点数量

		SiteNum *int64 `json:"SiteNum,omitempty" name:"SiteNum"`
		// 服务数量

		ServiceNum *int64 `json:"ServiceNum,omitempty" name:"ServiceNum"`
		// ZoneID

		ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
		// ZoneName

		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
		// DefaultZone

		DefaultZone *bool `json:"DefaultZone,omitempty" name:"DefaultZone"`
		// Status

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// CreateTime

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// LastUpdateTime

		LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportVendorMemberRequest struct {
	*tchttp.BaseRequest

	// FileName

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// Content

	Content *string `json:"Content,omitempty" name:"Content"`
	// VendorID

	VendorID *string `json:"VendorID,omitempty" name:"VendorID"`
}

func (r *ImportVendorMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportVendorMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotOriginDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotOriginDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotOriginDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data []*LicenseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppRequest struct {
	*tchttp.BaseRequest

	// 系统ID

	Sysid *string `json:"Sysid,omitempty" name:"Sysid"`
	// 应用ID

	PaasID *string `json:"PaasID,omitempty" name:"PaasID"`
	// 应用 token 见 Token

	Tokens []*TokenOfCreateApp `json:"Tokens,omitempty" name:"Tokens"`
	// 创建用户，见 User

	Operators []*UserOfCreateApp `json:"Operators,omitempty" name:"Operators"`
	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 系统ID

	SysID *string `json:"SysID,omitempty" name:"SysID"`
	// 应用描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Vendors

	Vendors []*Summary `json:"Vendors,omitempty" name:"Vendors"`
}

func (r *CreateAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroProtoFileRequest struct {
	*tchttp.BaseRequest

	// 文件组 ID

	ProtoGroupId *string `json:"ProtoGroupId,omitempty" name:"ProtoGroupId"`
	// 文件名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteMicroProtoFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMicroProtoFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedSystemRequest struct {
	*tchttp.BaseRequest

	// 系统 ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *RejectedSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectedSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProfileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProfileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApprovedSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovedSystemBatchRequest struct {
	*tchttp.BaseRequest

	// 系统 ID 列表

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
	// 系统 ID 列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *ApprovedSystemBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApprovedSystemBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMicroProtoFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMicroProtoFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMicroProtoFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthNodeOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAuthNodeOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuthNodeOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginPolicyRequest struct {
	*tchttp.BaseRequest

	// 参考附录 《对象说明》

	AtlasCaptcha *AtlasCaptcha `json:"AtlasCaptcha,omitempty" name:"AtlasCaptcha"`
	// 参考附录 《对象说明》

	AtlasLockSetting *AtlasLockSetting `json:"AtlasLockSetting,omitempty" name:"AtlasLockSetting"`
}

func (r *ModifyLoginPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SiteApproveLogVO struct {

	// 所属应用

	App *AppOfSiteApproveVo `json:"App,omitempty" name:"App"`
	// 申请时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 申请人

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// 站点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 访问路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 站点Id

	SiteID *string `json:"SiteID,omitempty" name:"SiteID"`
	// 审批记录id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 审批人

	ApprovalUser *string `json:"ApprovalUser,omitempty" name:"ApprovalUser"`
	// 审批时间

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 驳回原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type DescribeSiteRequestLogPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 站点请求日志数组

		Items []*SiteRequestLogVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSiteRequestLogPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteRequestLogPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSmartGatewayGroupRequest struct {
	*tchttp.BaseRequest

	// 节点阵列 ID

	SggID *string `json:"SggID,omitempty" name:"SggID"`
}

func (r *DescribeSmartGatewayGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSmartGatewayGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMenuAllListRequest struct {
	*tchttp.BaseRequest

	// 支持left、top、all入参

	Location *string `json:"Location,omitempty" name:"Location"`
}

func (r *DescribeMenuAllListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMenuAllListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgencyBatchRequest struct {
	*tchttp.BaseRequest

	// 机构ID数组

	IDs []*string `json:"IDs,omitempty" name:"IDs"`
}

func (r *DeleteAgencyBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgencyBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadSystemDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *FileDownloadResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadSystemDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadSystemDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCentreNodeInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyCentreNodeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCentreNodeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateSitePubPathRequest struct {
	*tchttp.BaseRequest

	// 站点公开路径

	PubPath *string `json:"PubPath,omitempty" name:"PubPath"`
	// 站点ID（如果是编辑状态，必须携带，新增状态不用携带）

	SiteID *string `json:"SiteID,omitempty" name:"SiteID"`
}

func (r *ValidateSitePubPathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateSitePubPathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateSitePubPathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发生冲突，true冲突，false未冲突

		IsConflict *bool `json:"IsConflict,omitempty" name:"IsConflict"`
		// 冲突的项

		Items []*ItemOfValidateSitePubPath `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateSitePubPathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateSitePubPathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcsRoleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 角色列表

		Items []*AcsRoleVO `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAcsRoleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcsRoleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MicroProtoGroup struct {

	// 文件组 ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 文件组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件组描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// CreateUser

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type CheckAppUniqueRequest struct {
	*tchttp.BaseRequest

	// 过滤器，字符串，同bson，例如 {"name": "amp"}

	FilterT *string `json:"FilterT,omitempty" name:"FilterT"`
	// 取值：app, sys

	Object *string `json:"Object,omitempty" name:"Object"`
}

func (r *CheckAppUniqueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAppUniqueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTipRequest struct {
	*tchttp.BaseRequest

	// 当前提示语

	CurrentTip *string `json:"CurrentTip,omitempty" name:"CurrentTip"`
	// 提示语码号 tip_code应为>=100000整型字符串且在有效范围内

	TipCode *string `json:"TipCode,omitempty" name:"TipCode"`
	// 是否为手机号格式

	PhoneRelated *bool `json:"PhoneRelated,omitempty" name:"PhoneRelated"`
	// 手机号脱敏方式 "0" 表示不显示手机号 "1" 表示脱敏保留前三位，后四位：1767096 "2" 表示脱敏前八位：****7096

	PhoneMaskFormat *string `json:"PhoneMaskFormat,omitempty" name:"PhoneMaskFormat"`
}

func (r *ModifyTipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpsertSqlTemplateRequest struct {
	*tchttp.BaseRequest

	// 服务 id

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 数据源 id

	DataSourceID *string `json:"DataSourceID,omitempty" name:"DataSourceID"`
	// sql 模板

	Sql *string `json:"Sql,omitempty" name:"Sql"`
	// 创建用户信息

	CreateUser *CreateUser `json:"CreateUser,omitempty" name:"CreateUser"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateTime

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

func (r *UpsertSqlTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpsertSqlTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
