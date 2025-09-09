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

package v20191213

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ActiveProviderConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActiveProviderConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveProviderConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStackVersionRequest struct {
	*tchttp.BaseRequest

	// 资源栈id

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
}

func (r *CreateStackVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStackVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAllProviderConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteAllProviderConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAllProviderConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分页大小

		Limit *int64 `json:"Limit,omitempty" name:"Limit"`
		// 偏移

		Offset *int64 `json:"Offset,omitempty" name:"Offset"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 模板列表(没有TfTree信息)

		Templates []*TemplateInfo `json:"Templates,omitempty" name:"Templates"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetContentFromUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// tf内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetContentFromUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetContentFromUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StackIdInfo struct {

	// 资源栈ID

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 版本ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
}

type ModifyProviderConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProviderConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProviderConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLivingLogKeyRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
}

func (r *GetLivingLogKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLivingLogKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateValidateEventRequest struct {
	*tchttp.BaseRequest

	// 版本ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 模板ID

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateValidateEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateValidateEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStackEventRequest struct {
	*tchttp.BaseRequest

	// 版本ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 事件类型,可选值：plan, apply, destroy

	EventType *string `json:"EventType,omitempty" name:"EventType"`
}

func (r *CreateStackEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStackEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板详情(带有TfTree信息)

		Template *TemplateInfo `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFilesFromUrlRequest struct {
	*tchttp.BaseRequest

	// 文件Url

	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *GetFilesFromUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFilesFromUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceDetailRequest struct {
	*tchttp.BaseRequest

	// 查询的资源ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *GetResourceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的模板信息

		Template *TemplateShortInfo `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除模板的Id

		Template *TemplateShortInfo `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesRequest struct {
	*tchttp.BaseRequest

	// 模板类型(默认0是私有模板,1是共有模板)

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 查询名字或者描述过滤条件

	Filter *string `json:"Filter,omitempty" name:"Filter"`
	// 排序,可选择字段InsertTime\UpdateTime,排序方式有+(asc),-(desc),eg:+InsertTime表示按照InsertTime正序排列

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsRequest struct {
	*tchttp.BaseRequest

	// Provider的名字  tencentcloud/alicloud/aws

	Provider *string `json:"Provider,omitempty" name:"Provider"`
}

func (r *GetRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceDirectoryRequest struct {
	*tchttp.BaseRequest

	// 云厂商名称，现在仅支持tencentcloud

	Provider *string `json:"Provider,omitempty" name:"Provider"`
	// 搜索关键字

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *GetResourceDirectoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceDirectoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测字段

		GG *ModifyTest `json:"GG,omitempty" name:"GG"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTemplateFilesRequest struct {
	*tchttp.BaseRequest

	// 模板id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *GetTemplateFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTemplateFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStackVersionRequest struct {
	*tchttp.BaseRequest

	// 资源栈id

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 要删除的版本id列表

	VersionIds []*int64 `json:"VersionIds,omitempty" name:"VersionIds"`
}

func (r *DeleteStackVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStackVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStackVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除成功的id

		Success *int64 `json:"Success,omitempty" name:"Success"`
		// 删除失败的资源栈版本列表

		Failed []*DeleteStackVersionFailedInfo `json:"Failed,omitempty" name:"Failed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStackVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStackVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板的id

		Template *TemplateShortInfo `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddProviderConfigRequest struct {
	*tchttp.BaseRequest

	// 支持的Provider类型。

	Provider *string `json:"Provider,omitempty" name:"Provider"`
	// Provider描述。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// Provider名字。

	Name *string `json:"Name,omitempty" name:"Name"`
	// Provider具体的配置。

	ProviderConfig *string `json:"ProviderConfig,omitempty" name:"ProviderConfig"`
}

func (r *AddProviderConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProviderConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateValidateEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 语法是否正确

		Success *bool `json:"Success,omitempty" name:"Success"`
		// 输出结果

		Console *string `json:"Console,omitempty" name:"Console"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateValidateEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateValidateEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// provider列表

		Providers []*ProviderResourceInfo `json:"Providers,omitempty" name:"Providers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResourceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProviderResourceInfo struct {

	// provider名字，目前只有tencentcloud

	ProviderName *string `json:"ProviderName,omitempty" name:"ProviderName"`
	// resource列表

	ResourceList []*ResourceInfo `json:"ResourceList,omitempty" name:"ResourceList"`
	// data source列表

	DataList []*ResourceInfo `json:"DataList,omitempty" name:"DataList"`
}

type GetEventResultRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
}

func (r *GetEventResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEventResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceStackRequest struct {
	*tchttp.BaseRequest

	// 资源栈ID

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 资源栈名称

	StackName *string `json:"StackName,omitempty" name:"StackName"`
	// 资源栈描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyResourceStackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceStackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountInfo struct {

	// appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
}

type CreateStackEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件ID

		EventId *int64 `json:"EventId,omitempty" name:"EventId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStackEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceListRequest struct {
	*tchttp.BaseRequest

	// 版本ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 搜索关键字

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *GetResourceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRoleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateShortInfo struct {

	// 模板Id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeRoleStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRoleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShortTemplates struct {

	// 模板Id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 模板的名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DownloadStackVersionAllFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// zip压缩包，包含所有tf文件，经过了base64编码

		ZipFile *string `json:"ZipFile,omitempty" name:"ZipFile"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadStackVersionAllFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadStackVersionAllFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStackVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共有多少个版本

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资源栈版本信息列表

		Versions []*StackVersionInfo `json:"Versions,omitempty" name:"Versions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStackVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStackVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProviderConfigRequest struct {
	*tchttp.BaseRequest

	// 修改provider的ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 修改provider的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 修改provider的描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 修改provider的具体配置

	ProviderConfig *string `json:"ProviderConfig,omitempty" name:"ProviderConfig"`
}

func (r *ModifyProviderConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProviderConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShortTemplatesRequest struct {
	*tchttp.BaseRequest

	// 模版类型

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模版名字过滤条件

	Filter *string `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeShortTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShortTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceStackRequest struct {
	*tchttp.BaseRequest

	// 过滤的地区列表

	Regions []*string `json:"Regions,omitempty" name:"Regions"`
	// 过滤的状态列表

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 搜索关键字

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页数据条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListResourceStackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceStackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStackVersionsRequest struct {
	*tchttp.BaseRequest

	// 资源栈id

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 根据版本id进行过滤

	VersionIds []*int64 `json:"VersionIds,omitempty" name:"VersionIds"`
	// 根据关键字进行过滤

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 根据版本状态进行过滤

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 排序，可选择字段CreateTime，排序方式有+(asc)，-(desc)。eg: +CreateTime表示按照CreateTime正序排列。默认使用CreateTime倒序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 查询偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最多查询多少条版本信息，默认值10，最小为1，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetStackVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStackVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceStackRequest struct {
	*tchttp.BaseRequest

	// 资源栈ID

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 版本ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 资源栈名称

	StackName *string `json:"StackName,omitempty" name:"StackName"`
	// 资源栈描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateResourceStackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceStackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSafeMgrScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 扫描异步任务的id

		AsyncID *string `json:"AsyncID,omitempty" name:"AsyncID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSafeMgrScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSafeMgrScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板的id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadStackVersionAllFilesRequest struct {
	*tchttp.BaseRequest

	// 资源栈id

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 要下载的版本id

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *DownloadStackVersionAllFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadStackVersionAllFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStackVersionContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本名字

		VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
		// tf树内容，经过了gzip压缩并且使用base64编码

		TfTree *string `json:"TfTree,omitempty" name:"TfTree"`
		// 资源栈state文件，经过了gzip压缩并且使用base64编码

		TfState *string `json:"TfState,omitempty" name:"TfState"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStackVersionContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStackVersionContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProviderConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 偏移

		Offset *int64 `json:"Offset,omitempty" name:"Offset"`
		// 分页大小

		Limit *int64 `json:"Limit,omitempty" name:"Limit"`
		// provider配置列表

		List []*ProviderConfigDetail `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProviderConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProviderConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProviderConfigRequest struct {
	*tchttp.BaseRequest

	// 排序,可选择字段InsertTime\UpdateTime\Active,排序方式有+(asc),-(desc),eg:+InsertTime表示按照InsertTime正序排列

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否拉取激活的provider列表.0表示拉取所有,1表示拉取激活的列表

	Active *int64 `json:"Active,omitempty" name:"Active"`
	// 分页偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小,默认10个,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListProviderConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProviderConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceStackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源栈总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资源栈列表

		ResourceStacks []*ResourceStackInfo `json:"ResourceStacks,omitempty" name:"ResourceStacks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceStackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateComplexTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模板的描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 模板的配置文件

	TfTree *string `json:"TfTree,omitempty" name:"TfTree"`
}

func (r *CreateComplexTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComplexTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceDirectoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 产品文档列表

		Products []*ProductDocInfo `json:"Products,omitempty" name:"Products"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResourceDirectoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceDirectoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStackVersionContentRequest struct {
	*tchttp.BaseRequest

	// 资源栈id

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 资源栈版本id

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *GetStackVersionContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStackVersionContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地区的编码(mysql使用)

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
}

type ResourceTypeInfo struct {

	// 资源ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 资源名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源所属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 资源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 资源描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type GetEventResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件类型

		EventType *string `json:"EventType,omitempty" name:"EventType"`
		// 终端输出结果(base64)

		Console *string `json:"Console,omitempty" name:"Console"`
		// 日志 (base64)

		Log *string `json:"Log,omitempty" name:"Log"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEventResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEventResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSafeMgrScanResultRequest struct {
	*tchttp.BaseRequest

	// 扫描的异步任务ID

	AsyncID *string `json:"AsyncID,omitempty" name:"AsyncID"`
	// 扫描的产品名字,例如cvm.cos,redis等

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *ListSafeMgrScanResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSafeMgrScanResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceInfo struct {

	// 实例ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 实例名称

	AttrName *string `json:"AttrName,omitempty" name:"AttrName"`
	// 资源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 资源名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源编号

	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type ResourceStackInfo struct {

	// 资源栈ID

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 资源栈名称

	StackName *string `json:"StackName,omitempty" name:"StackName"`
	// 资源栈描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 当前版本ID

	CurrentVersionId *int64 `json:"CurrentVersionId,omitempty" name:"CurrentVersionId"`
	// 当前版本名称

	CurrentVersionName *string `json:"CurrentVersionName,omitempty" name:"CurrentVersionName"`
	// 资源栈状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// 云厂商

	Provider *string `json:"Provider,omitempty" name:"Provider"`
	// 资源栈来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 栈ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否使用了TIC授权

	TokenFlag *bool `json:"TokenFlag,omitempty" name:"TokenFlag"`
	// 密钥配置

	Credential *string `json:"Credential,omitempty" name:"Credential"`
}

type SafeItem struct {

	// 条目名字,eg:InstanceNumber 产品里边唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 提醒,eg:未开启登录保护的风险

	Hint *string `json:"Hint,omitempty" name:"Hint"`
	// 安全等级 0无风险 1低风险 2中风险 3高风险

	RiskClass *int64 `json:"RiskClass,omitempty" name:"RiskClass"`
	// 本条要减去多少分

	Sub *float64 `json:"Sub,omitempty" name:"Sub"`
	// 描述信息,如"需要获取安全组的访问规则，并对规则进行分析高危端口：主要包括TCP端口（1-1024）、UDP端口( 1-1024)、MySQL端口（3306）、SQL Server端口(1433、1434)、SSH端口（22）和Windows管理端口(3389)"

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 怎么修复,如:关闭高危端口,具体关闭操作,请参考安全实例

	Repair *string `json:"Repair,omitempty" name:"Repair"`
	// 修复的Url

	RepairUrl *string `json:"RepairUrl,omitempty" name:"RepairUrl"`
	// 如果是综合信息  如"15", "100G"等

	Sum *string `json:"Sum,omitempty" name:"Sum"`
	// 不安全的实例I

	UnsafeIds []*Instance `json:"UnsafeIds,omitempty" name:"UnsafeIds"`
}

type TemplateInfo struct {

	// 模板Id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 模板的类型,0为私有,1为公有

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 创建者的appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 模板的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模板的描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 模版创建的时间戳

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 模板更新的时间戳

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 模版的配置信息

	TfTree *string `json:"TfTree,omitempty" name:"TfTree"`
}

type AddProviderConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Provider具体的配置

		Provider *ProviderConfigDetail `json:"Provider,omitempty" name:"Provider"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddProviderConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProviderConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateComplexTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板ID

		Template *TemplateShortInfo `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateComplexTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComplexTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceStackRequest struct {
	*tchttp.BaseRequest

	// 资源栈ID

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
}

func (r *DeleteResourceStackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceStackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceStackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源栈详情

		ResourceStack *ResourceStackInfo `json:"ResourceStack,omitempty" name:"ResourceStack"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceStackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源ID

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 资源所属provider

		Provider *string `json:"Provider,omitempty" name:"Provider"`
		// 资源名字

		Name *string `json:"Name,omitempty" name:"Name"`
		// 资源所属产品

		Product *string `json:"Product,omitempty" name:"Product"`
		// 资源类型

		Type *string `json:"Type,omitempty" name:"Type"`
		// 输入参数列表，示例一，子参数为简单类型：[{"Name":"egress","Type":"list","Description":"test","Deprecated":"","Required":false,"Optional":true,"Computed":false,"ForceNew":false,"SubArguments":[{"Name":"","Type":"string","Description":"","Deprecated":"","Required":false,"Optional":false,"Computed":false,"ForceNew":false,"SubArguments":null}]}]。示例二，子参数为复杂类型：[{"Name":"data_disk","Type":"list","Description":"test.","Deprecated":"","Required":false,"Optional":true,"Computed":false,"ForceNew":true,"SubArguments":[{"Name":"disk_type","Type":"string","Description":"test","Deprecated":"","Required":false,"Optional":true,"Computed":false,"ForceNew":true,"SubArguments":null},{"Name":"snapshot_id","Type":"string","Description":"test","Deprecated":"","Required":false,"Optional":true,"Computed":false,"ForceNew":true,"SubArguments":null}]}]

		InputArguments *string `json:"InputArguments,omitempty" name:"InputArguments"`
		// 输出参数列表，示例一，子参数为简单类型：{"Name":"assigned_eip_set","Type":"list","Description":"test","Deprecated":"","Required":false,"Optional":false,"Computed":true,"ForceNew":false,"SubArguments":[{"Name":"","Type":"string","Description":"","Deprecated":"","Required":false,"Optional":false,"Computed":false,"ForceNew":false,"SubArguments":null}]}。示例二，子参数为复杂类型：[{"Name":"nats","Type":"list","Description":"test","Deprecated":"","Required":false,"Optional":false,"Computed":true,"ForceNew":false,"SubArguments":[{"Name":"bandwidth","Type":"int","Description":"test","Deprecated":"","Required":false,"Optional":false,"Computed":true,"ForceNew":false,"SubArguments":null},{"Name":"create_time","Type":"string","Description":"test","Deprecated":"","Required":false,"Optional":false,"Computed":true,"ForceNew":false,"SubArguments":null}]}]

		OutputArguments *string `json:"OutputArguments,omitempty" name:"OutputArguments"`
		// 资源简介

		Description *string `json:"Description,omitempty" name:"Description"`
		// 资源使用样例

		Example *string `json:"Example,omitempty" name:"Example"`
		// 资源的国际站文档链接

		IntlDocUrl *string `json:"IntlDocUrl,omitempty" name:"IntlDocUrl"`
		// 资源的中国站文档链接

		ChinaDocUrl *string `json:"ChinaDocUrl,omitempty" name:"ChinaDocUrl"`
		// 资源的国际站api文档链接

		IntlApiUrl *string `json:"IntlApiUrl,omitempty" name:"IntlApiUrl"`
		// 资源的中国站api文档链接

		ChinaApiUrl *string `json:"ChinaApiUrl,omitempty" name:"ChinaApiUrl"`
		// terraform的文档链接

		TerraformDocUrl *string `json:"TerraformDocUrl,omitempty" name:"TerraformDocUrl"`
		// 资源对应的provider版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// 资源支持的平台，TIC或Terraform

		Support *string `json:"Support,omitempty" name:"Support"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResourceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceDocInfo struct {

	// 文档ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 资源名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ActiveProviderConfigRequest struct {
	*tchttp.BaseRequest

	// provider的id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 0表示反激活,1表示激活

	Active *int64 `json:"Active,omitempty" name:"Active"`
}

func (r *ActiveProviderConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveProviderConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetContentFromUrlRequest struct {
	*tchttp.BaseRequest

	// tf内容URL

	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *GetContentFromUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetContentFromUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListStackEventRequest struct {
	*tchttp.BaseRequest

	// 资源栈ID

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 版本ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 事件ID

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 事件类型，可选值：plan, apply, destroy

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 过滤的状态列表

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页数据条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索关键字

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListStackEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStackEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StackVersionInfo struct {

	// 资源栈id

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 资源栈版本id

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 创建版本时的appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建版本时的uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 创建版本时的SubAccountUin

	SubAccountUin *int64 `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	// 资源栈版本名称

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
	// 资源栈版本状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 资源栈版本状态信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 资源栈版本最近的一次事件

	LastEventId *int64 `json:"LastEventId,omitempty" name:"LastEventId"`
	// 插入时间戳

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间戳

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreateDraftRequest struct {
	*tchttp.BaseRequest

	// 创建来源，可选值(`url`, `direct`, `public_template`, `private_template`)

	Source *string `json:"Source,omitempty" name:"Source"`
	// provider的id

	ProviderIds []*int64 `json:"ProviderIds,omitempty" name:"ProviderIds"`
	// tf文件内容(gzip压缩后base64)

	Content *string `json:"Content,omitempty" name:"Content"`
	// 地域

	StackRegion *string `json:"StackRegion,omitempty" name:"StackRegion"`
	// 云厂商

	Provider *string `json:"Provider,omitempty" name:"Provider"`
}

func (r *CreateDraftRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDraftRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSafeMgrScanTaskRequest struct {
	*tchttp.BaseRequest

	// 与扫描的产品,例如:mysql,redis,cvm,cbs,cos,cam,mongodb,es,emr等

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateSafeMgrScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSafeMgrScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveStackVersionRequest struct {
	*tchttp.BaseRequest

	// 资源栈id

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 要保存的版本id

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 要保存的版本名字，如果此参数为空字符串，则不修改版本名字

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
	// 要保存的版本描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// tf树内容，经过了gzip压缩并且使用base64编码

	TfTree *string `json:"TfTree,omitempty" name:"TfTree"`
}

func (r *SaveStackVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveStackVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceStackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源栈ID

		StackId *int64 `json:"StackId,omitempty" name:"StackId"`
		// 版本ID

		VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
		// 事件ID

		EventId *int64 `json:"EventId,omitempty" name:"EventId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateResourceStackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProviderConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProviderConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProviderConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateRequest struct {
	*tchttp.BaseRequest

	// 修改模板的id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// 修改模板的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 修改模板的描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 修改模板的配置文件

	TfTree *string `json:"TfTree,omitempty" name:"TfTree"`
}

func (r *ModifyTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceStackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteResourceStackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShortTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全量模板简短信息列表

		Templates []*ShortTemplates `json:"Templates,omitempty" name:"Templates"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShortTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShortTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTest struct {

	// 测试

	TemplateId []*bool `json:"TemplateId,omitempty" name:"TemplateId"`
}

type StackSetMemberInfo struct {

	// 资源栈ID

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 资源栈名称

	StackName *string `json:"StackName,omitempty" name:"StackName"`
	// 资源栈描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 云厂商

	Provider *string `json:"Provider,omitempty" name:"Provider"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 资源栈状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeleteTemplateRequest struct {
	*tchttp.BaseRequest

	// 删除模板的id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资源列表

		Resources []*ResourceTypeInfo `json:"Resources,omitempty" name:"Resources"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色授权状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFilesFromUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置文件内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFilesFromUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFilesFromUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveStackVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveStackVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveStackVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStackVersionFailedInfo struct {

	// 删除失败的资源栈版本id

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 删除失败的原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type ProviderConfigDetail struct {

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Provider类型

	Provider *string `json:"Provider,omitempty" name:"Provider"`
	// Provider的描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// Provider的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// Provider的具体配置

	ProviderConfig *string `json:"ProviderConfig,omitempty" name:"ProviderConfig"`
	// 是否被激活

	Active *int64 `json:"Active,omitempty" name:"Active"`
	// 是否已经被使用

	Using *int64 `json:"Using,omitempty" name:"Using"`
	// 插入时间戳

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 更新时间戳

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
}

type CreateTemplateRequest struct {
	*tchttp.BaseRequest

	// 创建模板的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建模板的描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAllProviderConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAllProviderConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAllProviderConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// appid

		AppId *int64 `json:"AppId,omitempty" name:"AppId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceStackRequest struct {
	*tchttp.BaseRequest

	// 资源栈ID

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
}

func (r *DescribeResourceStackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceStackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourcesRequest struct {
	*tchttp.BaseRequest

	// provider名字，现在仅支持tencentcloud

	Provider *string `json:"Provider,omitempty" name:"Provider"`
	// 根据关键字进行过滤

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 查询偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最多查询多少个资源，默认值10，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleStatusRequest struct {
	*tchttp.BaseRequest

	// 角色授权状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateRoleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRoleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StackEventInfo struct {

	// 资源栈ID

	StackId *int64 `json:"StackId,omitempty" name:"StackId"`
	// 版本ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 版本名称

	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
	// 事件ID

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 事件描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateStackVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源栈版本id

		VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
		// 资源栈版本名字

		VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStackVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStackVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSafeMgrScanResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务的状态  0未开始 1运行中 2运行完成  <0表示执行失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 任务的插入时间戳

		InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
		// 任务的完成呢时间戳

		FinishTime *int64 `json:"FinishTime,omitempty" name:"FinishTime"`
		// 任务的结果

		Result []*SafeItem `json:"Result,omitempty" name:"Result"`
		// 任务如果出错,为原因描述

		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`
		// 任务出错的错误代码

		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`
		// 任务的异步ID

		AsyncID *string `json:"AsyncID,omitempty" name:"AsyncID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSafeMgrScanResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSafeMgrScanResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateFromStackRequest struct {
	*tchttp.BaseRequest

	// 版本ID

	VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
	// 模板名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 模板描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateTemplateFromStackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateFromStackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceStackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceStackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 复制生成的模板的id

		Template *TemplateShortInfo `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDraftResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的资源栈ID

		StackId *int64 `json:"StackId,omitempty" name:"StackId"`
		// 创建的版本ID

		VersionId *int64 `json:"VersionId,omitempty" name:"VersionId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDraftResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDraftResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateFromStackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板ID

		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTemplateFromStackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateFromStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTemplateFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板文件的zip压缩后base64

		ZipFile *string `json:"ZipFile,omitempty" name:"ZipFile"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTemplateFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTemplateFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceInputArgumentInfo struct {

	// 参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 参数是否为必填项

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 参数是否不可简单更新

	ForceNew *bool `json:"ForceNew,omitempty" name:"ForceNew"`
	// 参数简介

	Description *string `json:"Description,omitempty" name:"Description"`
	// 参数取值范围

	ValueRange *string `json:"ValueRange,omitempty" name:"ValueRange"`
}

type CopyTemplateRequest struct {
	*tchttp.BaseRequest

	// 要复制模板的id

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CopyTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductDocInfo struct {

	// 产品名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Data Source列表

	DataSources []*ResourceDocInfo `json:"DataSources,omitempty" name:"DataSources"`
	// Resource列表

	Resources []*ResourceDocInfo `json:"Resources,omitempty" name:"Resources"`
}

type DeleteProviderConfigRequest struct {
	*tchttp.BaseRequest

	// 删除的provider的id列表

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteProviderConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProviderConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLivingLogKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Key ID

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// Token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 时间戳

		Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLivingLogKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLivingLogKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地区列表

		RegionList []*RegionInfo `json:"RegionList,omitempty" name:"RegionList"`
		// 地区总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListStackEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 事件列表

		StackEvents []*StackEventInfo `json:"StackEvents,omitempty" name:"StackEvents"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListStackEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地区, 例如ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地区名称, 例如广州地区

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 英文地区名称

	EnglishRegionName *string `json:"EnglishRegionName,omitempty" name:"EnglishRegionName"`
}

type StackSetInfo struct {

	// 资源栈集ID

	StackSetId *int64 `json:"StackSetId,omitempty" name:"StackSetId"`
	// 资源栈集名称

	StackSetName *string `json:"StackSetName,omitempty" name:"StackSetName"`
	// 资源栈集描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 包含的资源栈数目

	StackCount *int64 `json:"StackCount,omitempty" name:"StackCount"`
	// 云厂商

	Providers *string `json:"Providers,omitempty" name:"Providers"`
}
