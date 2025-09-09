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
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-12-13"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewDescribeTemplateRequest() (request *DescribeTemplateRequest) {
	request = &DescribeTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DescribeTemplate")
	return
}

func NewDescribeTemplateResponse() (response *DescribeTemplateResponse) {
	response = &DescribeTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取模板详情
func (c *Client) DescribeTemplate(request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeTemplateRequest()
	}
	response = NewDescribeTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetStackVersionContentRequest() (request *GetStackVersionContentRequest) {
	request = &GetStackVersionContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetStackVersionContent")
	return
}

func NewGetStackVersionContentResponse() (response *GetStackVersionContentResponse) {
	response = &GetStackVersionContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取版本tf，state文件
func (c *Client) GetStackVersionContent(request *GetStackVersionContentRequest) (response *GetStackVersionContentResponse, err error) {
	if request == nil {
		request = NewGetStackVersionContentRequest()
	}
	response = NewGetStackVersionContentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRequest() (request *ModifyRequest) {
	request = &ModifyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "Modify")
	return
}

func NewModifyResponse() (response *ModifyResponse) {
	response = &ModifyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 探测
func (c *Client) Modify(request *ModifyRequest) (response *ModifyResponse, err error) {
	if request == nil {
		request = NewModifyRequest()
	}
	response = NewModifyResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourceDetailRequest() (request *GetResourceDetailRequest) {
	request = &GetResourceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetResourceDetail")
	return
}

func NewGetResourceDetailResponse() (response *GetResourceDetailResponse) {
	response = &GetResourceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源参数详细信息
func (c *Client) GetResourceDetail(request *GetResourceDetailRequest) (response *GetResourceDetailResponse, err error) {
	if request == nil {
		request = NewGetResourceDetailRequest()
	}
	response = NewGetResourceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProviderConfigRequest() (request *DeleteProviderConfigRequest) {
	request = &DeleteProviderConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DeleteProviderConfig")
	return
}

func NewDeleteProviderConfigResponse() (response *DeleteProviderConfigResponse) {
	response = &DeleteProviderConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Provider配置。
func (c *Client) DeleteProviderConfig(request *DeleteProviderConfigRequest) (response *DeleteProviderConfigResponse, err error) {
	if request == nil {
		request = NewDeleteProviderConfigRequest()
	}
	response = NewDeleteProviderConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteResourceStackRequest() (request *DeleteResourceStackRequest) {
	request = &DeleteResourceStackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DeleteResourceStack")
	return
}

func NewDeleteResourceStackResponse() (response *DeleteResourceStackResponse) {
	response = &DeleteResourceStackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资源栈
func (c *Client) DeleteResourceStack(request *DeleteResourceStackRequest) (response *DeleteResourceStackResponse, err error) {
	if request == nil {
		request = NewDeleteResourceStackRequest()
	}
	response = NewDeleteResourceStackResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourceDirectoryRequest() (request *GetResourceDirectoryRequest) {
	request = &GetResourceDirectoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetResourceDirectory")
	return
}

func NewGetResourceDirectoryResponse() (response *GetResourceDirectoryResponse) {
	response = &GetResourceDirectoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源文档目录
func (c *Client) GetResourceDirectory(request *GetResourceDirectoryRequest) (response *GetResourceDirectoryResponse, err error) {
	if request == nil {
		request = NewGetResourceDirectoryRequest()
	}
	response = NewGetResourceDirectoryResponse()
	err = c.Send(request, response)
	return
}

func NewCopyTemplateRequest() (request *CopyTemplateRequest) {
	request = &CopyTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CopyTemplate")
	return
}

func NewCopyTemplateResponse() (response *CopyTemplateResponse) {
	response = &CopyTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制模板
func (c *Client) CopyTemplate(request *CopyTemplateRequest) (response *CopyTemplateResponse, err error) {
	if request == nil {
		request = NewCopyTemplateRequest()
	}
	response = NewCopyTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateStackEventRequest() (request *CreateStackEventRequest) {
	request = &CreateStackEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CreateStackEvent")
	return
}

func NewCreateStackEventResponse() (response *CreateStackEventResponse) {
	response = &CreateStackEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建事件
func (c *Client) CreateStackEvent(request *CreateStackEventRequest) (response *CreateStackEventResponse, err error) {
	if request == nil {
		request = NewCreateStackEventRequest()
	}
	response = NewCreateStackEventResponse()
	err = c.Send(request, response)
	return
}

func NewSaveStackVersionRequest() (request *SaveStackVersionRequest) {
	request = &SaveStackVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "SaveStackVersion")
	return
}

func NewSaveStackVersionResponse() (response *SaveStackVersionResponse) {
	response = &SaveStackVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存版本信息
func (c *Client) SaveStackVersion(request *SaveStackVersionRequest) (response *SaveStackVersionResponse, err error) {
	if request == nil {
		request = NewSaveStackVersionRequest()
	}
	response = NewSaveStackVersionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateResourceStackRequest() (request *CreateResourceStackRequest) {
	request = &CreateResourceStackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CreateResourceStack")
	return
}

func NewCreateResourceStackResponse() (response *CreateResourceStackResponse) {
	response = &CreateResourceStackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建资源栈
func (c *Client) CreateResourceStack(request *CreateResourceStackRequest) (response *CreateResourceStackResponse, err error) {
	if request == nil {
		request = NewCreateResourceStackRequest()
	}
	response = NewCreateResourceStackResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTemplateRequest() (request *CreateTemplateRequest) {
	request = &CreateTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CreateTemplate")
	return
}

func NewCreateTemplateResponse() (response *CreateTemplateResponse) {
	response = &CreateTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建模板
func (c *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
	if request == nil {
		request = NewCreateTemplateRequest()
	}
	response = NewCreateTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourceListRequest() (request *GetResourceListRequest) {
	request = &GetResourceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetResourceList")
	return
}

func NewGetResourceListResponse() (response *GetResourceListResponse) {
	response = &GetResourceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源列表
func (c *Client) GetResourceList(request *GetResourceListRequest) (response *GetResourceListResponse, err error) {
	if request == nil {
		request = NewGetResourceListRequest()
	}
	response = NewGetResourceListResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceStackRequest() (request *ListResourceStackRequest) {
	request = &ListResourceStackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "ListResourceStack")
	return
}

func NewListResourceStackResponse() (response *ListResourceStackResponse) {
	response = &ListResourceStackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源栈列表
func (c *Client) ListResourceStack(request *ListResourceStackRequest) (response *ListResourceStackResponse, err error) {
	if request == nil {
		request = NewListResourceStackRequest()
	}
	response = NewListResourceStackResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAllProviderConfigRequest() (request *DeleteAllProviderConfigRequest) {
	request = &DeleteAllProviderConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DeleteAllProviderConfig")
	return
}

func NewDeleteAllProviderConfigResponse() (response *DeleteAllProviderConfigResponse) {
	response = &DeleteAllProviderConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清理所有配置
func (c *Client) DeleteAllProviderConfig(request *DeleteAllProviderConfigRequest) (response *DeleteAllProviderConfigResponse, err error) {
	if request == nil {
		request = NewDeleteAllProviderConfigRequest()
	}
	response = NewDeleteAllProviderConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetStackVersionsRequest() (request *GetStackVersionsRequest) {
	request = &GetStackVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetStackVersions")
	return
}

func NewGetStackVersionsResponse() (response *GetStackVersionsResponse) {
	response = &GetStackVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询版本基本信息
func (c *Client) GetStackVersions(request *GetStackVersionsRequest) (response *GetStackVersionsResponse, err error) {
	if request == nil {
		request = NewGetStackVersionsRequest()
	}
	response = NewGetStackVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewListProviderConfigRequest() (request *ListProviderConfigRequest) {
	request = &ListProviderConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "ListProviderConfig")
	return
}

func NewListProviderConfigResponse() (response *ListProviderConfigResponse) {
	response = &ListProviderConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取provider配置列表
func (c *Client) ListProviderConfig(request *ListProviderConfigRequest) (response *ListProviderConfigResponse, err error) {
	if request == nil {
		request = NewListProviderConfigRequest()
	}
	response = NewListProviderConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateStackVersionRequest() (request *CreateStackVersionRequest) {
	request = &CreateStackVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CreateStackVersion")
	return
}

func NewCreateStackVersionResponse() (response *CreateStackVersionResponse) {
	response = &CreateStackVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建资源栈版本
func (c *Client) CreateStackVersion(request *CreateStackVersionRequest) (response *CreateStackVersionResponse, err error) {
	if request == nil {
		request = NewCreateStackVersionRequest()
	}
	response = NewCreateStackVersionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProviderConfigRequest() (request *ModifyProviderConfigRequest) {
	request = &ModifyProviderConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "ModifyProviderConfig")
	return
}

func NewModifyProviderConfigResponse() (response *ModifyProviderConfigResponse) {
	response = &ModifyProviderConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Provider信息
func (c *Client) ModifyProviderConfig(request *ModifyProviderConfigRequest) (response *ModifyProviderConfigResponse, err error) {
	if request == nil {
		request = NewModifyProviderConfigRequest()
	}
	response = NewModifyProviderConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetRegionsRequest() (request *GetRegionsRequest) {
	request = &GetRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetRegions")
	return
}

func NewGetRegionsResponse() (response *GetRegionsResponse) {
	response = &GetRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取地区列表
func (c *Client) GetRegions(request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
	if request == nil {
		request = NewGetRegionsRequest()
	}
	response = NewGetRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewGetTemplateFilesRequest() (request *GetTemplateFilesRequest) {
	request = &GetTemplateFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetTemplateFiles")
	return
}

func NewGetTemplateFilesResponse() (response *GetTemplateFilesResponse) {
	response = &GetTemplateFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 得到压缩的模板文件
func (c *Client) GetTemplateFiles(request *GetTemplateFilesRequest) (response *GetTemplateFilesResponse, err error) {
	if request == nil {
		request = NewGetTemplateFilesRequest()
	}
	response = NewGetTemplateFilesResponse()
	err = c.Send(request, response)
	return
}

func NewGetFilesFromUrlRequest() (request *GetFilesFromUrlRequest) {
	request = &GetFilesFromUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetFilesFromUrl")
	return
}

func NewGetFilesFromUrlResponse() (response *GetFilesFromUrlResponse) {
	response = &GetFilesFromUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从URL中导入文件
func (c *Client) GetFilesFromUrl(request *GetFilesFromUrlRequest) (response *GetFilesFromUrlResponse, err error) {
	if request == nil {
		request = NewGetFilesFromUrlRequest()
	}
	response = NewGetFilesFromUrlResponse()
	err = c.Send(request, response)
	return
}

func NewListSafeMgrScanResultRequest() (request *ListSafeMgrScanResultRequest) {
	request = &ListSafeMgrScanResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "ListSafeMgrScanResult")
	return
}

func NewListSafeMgrScanResultResponse() (response *ListSafeMgrScanResultResponse) {
	response = &ListSafeMgrScanResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示扫描任务结果
func (c *Client) ListSafeMgrScanResult(request *ListSafeMgrScanResultRequest) (response *ListSafeMgrScanResultResponse, err error) {
	if request == nil {
		request = NewListSafeMgrScanResultRequest()
	}
	response = NewListSafeMgrScanResultResponse()
	err = c.Send(request, response)
	return
}

func NewGetLivingLogKeyRequest() (request *GetLivingLogKeyRequest) {
	request = &GetLivingLogKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetLivingLogKey")
	return
}

func NewGetLivingLogKeyResponse() (response *GetLivingLogKeyResponse) {
	response = &GetLivingLogKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实时日志密钥
func (c *Client) GetLivingLogKey(request *GetLivingLogKeyRequest) (response *GetLivingLogKeyResponse, err error) {
	if request == nil {
		request = NewGetLivingLogKeyRequest()
	}
	response = NewGetLivingLogKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceStackRequest() (request *DescribeResourceStackRequest) {
	request = &DescribeResourceStackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DescribeResourceStack")
	return
}

func NewDescribeResourceStackResponse() (response *DescribeResourceStackResponse) {
	response = &DescribeResourceStackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源栈详情
func (c *Client) DescribeResourceStack(request *DescribeResourceStackRequest) (response *DescribeResourceStackResponse, err error) {
	if request == nil {
		request = NewDescribeResourceStackRequest()
	}
	response = NewDescribeResourceStackResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRoleStatusRequest() (request *DescribeRoleStatusRequest) {
	request = &DescribeRoleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DescribeRoleStatus")
	return
}

func NewDescribeRoleStatusResponse() (response *DescribeRoleStatusResponse) {
	response = &DescribeRoleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询角色授权状态
func (c *Client) DescribeRoleStatus(request *DescribeRoleStatusRequest) (response *DescribeRoleStatusResponse, err error) {
	if request == nil {
		request = NewDescribeRoleStatusRequest()
	}
	response = NewDescribeRoleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetEventResultRequest() (request *GetEventResultRequest) {
	request = &GetEventResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetEventResult")
	return
}

func NewGetEventResultResponse() (response *GetEventResultResponse) {
	response = &GetEventResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取事件结果
func (c *Client) GetEventResult(request *GetEventResultRequest) (response *GetEventResultResponse, err error) {
	if request == nil {
		request = NewGetEventResultRequest()
	}
	response = NewGetEventResultResponse()
	err = c.Send(request, response)
	return
}

func NewGetContentFromUrlRequest() (request *GetContentFromUrlRequest) {
	request = &GetContentFromUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetContentFromUrl")
	return
}

func NewGetContentFromUrlResponse() (response *GetContentFromUrlResponse) {
	response = &GetContentFromUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// URL获取tf内容
func (c *Client) GetContentFromUrl(request *GetContentFromUrlRequest) (response *GetContentFromUrlResponse, err error) {
	if request == nil {
		request = NewGetContentFromUrlRequest()
	}
	response = NewGetContentFromUrlResponse()
	err = c.Send(request, response)
	return
}

func NewActiveProviderConfigRequest() (request *ActiveProviderConfigRequest) {
	request = &ActiveProviderConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "ActiveProviderConfig")
	return
}

func NewActiveProviderConfigResponse() (response *ActiveProviderConfigResponse) {
	response = &ActiveProviderConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 激活或者取消激活Provider配置
func (c *Client) ActiveProviderConfig(request *ActiveProviderConfigRequest) (response *ActiveProviderConfigResponse, err error) {
	if request == nil {
		request = NewActiveProviderConfigRequest()
	}
	response = NewActiveProviderConfigResponse()
	err = c.Send(request, response)
	return
}

func NewAddProviderConfigRequest() (request *AddProviderConfigRequest) {
	request = &AddProviderConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "AddProviderConfig")
	return
}

func NewAddProviderConfigResponse() (response *AddProviderConfigResponse) {
	response = &AddProviderConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加Provider配置。
func (c *Client) AddProviderConfig(request *AddProviderConfigRequest) (response *AddProviderConfigResponse, err error) {
	if request == nil {
		request = NewAddProviderConfigRequest()
	}
	response = NewAddProviderConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTemplateRequest() (request *ModifyTemplateRequest) {
	request = &ModifyTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "ModifyTemplate")
	return
}

func NewModifyTemplateResponse() (response *ModifyTemplateResponse) {
	response = &ModifyTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改模板
func (c *Client) ModifyTemplate(request *ModifyTemplateRequest) (response *ModifyTemplateResponse, err error) {
	if request == nil {
		request = NewModifyTemplateRequest()
	}
	response = NewModifyTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShortTemplatesRequest() (request *DescribeShortTemplatesRequest) {
	request = &DescribeShortTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DescribeShortTemplates")
	return
}

func NewDescribeShortTemplatesResponse() (response *DescribeShortTemplatesResponse) {
	response = &DescribeShortTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取全量模板简短信息列表
func (c *Client) DescribeShortTemplates(request *DescribeShortTemplatesRequest) (response *DescribeShortTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeShortTemplatesRequest()
	}
	response = NewDescribeShortTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourcesRequest() (request *GetResourcesRequest) {
	request = &GetResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "GetResources")
	return
}

func NewGetResourcesResponse() (response *GetResourcesResponse) {
	response = &GetResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取资源类型列表
func (c *Client) GetResources(request *GetResourcesRequest) (response *GetResourcesResponse, err error) {
	if request == nil {
		request = NewGetResourcesRequest()
	}
	response = NewGetResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewListStackEventRequest() (request *ListStackEventRequest) {
	request = &ListStackEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "ListStackEvent")
	return
}

func NewListStackEventResponse() (response *ListStackEventResponse) {
	response = &ListStackEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取事件列表
func (c *Client) ListStackEvent(request *ListStackEventRequest) (response *ListStackEventResponse, err error) {
	if request == nil {
		request = NewListStackEventRequest()
	}
	response = NewListStackEventResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRoleStatusRequest() (request *UpdateRoleStatusRequest) {
	request = &UpdateRoleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "UpdateRoleStatus")
	return
}

func NewUpdateRoleStatusResponse() (response *UpdateRoleStatusResponse) {
	response = &UpdateRoleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新角色授权状态
func (c *Client) UpdateRoleStatus(request *UpdateRoleStatusRequest) (response *UpdateRoleStatusResponse, err error) {
	if request == nil {
		request = NewUpdateRoleStatusRequest()
	}
	response = NewUpdateRoleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
	request = &DescribeTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DescribeTemplates")
	return
}

func NewDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
	response = &DescribeTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取模板列表
func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeTemplatesRequest()
	}
	response = NewDescribeTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSafeMgrScanTaskRequest() (request *CreateSafeMgrScanTaskRequest) {
	request = &CreateSafeMgrScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CreateSafeMgrScanTask")
	return
}

func NewCreateSafeMgrScanTaskResponse() (response *CreateSafeMgrScanTaskResponse) {
	response = &CreateSafeMgrScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建安全扫描任务
func (c *Client) CreateSafeMgrScanTask(request *CreateSafeMgrScanTaskRequest) (response *CreateSafeMgrScanTaskResponse, err error) {
	if request == nil {
		request = NewCreateSafeMgrScanTaskRequest()
	}
	response = NewCreateSafeMgrScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTemplateFromStackRequest() (request *CreateTemplateFromStackRequest) {
	request = &CreateTemplateFromStackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CreateTemplateFromStack")
	return
}

func NewCreateTemplateFromStackResponse() (response *CreateTemplateFromStackResponse) {
	response = &CreateTemplateFromStackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从资源栈创建模板
func (c *Client) CreateTemplateFromStack(request *CreateTemplateFromStackRequest) (response *CreateTemplateFromStackResponse, err error) {
	if request == nil {
		request = NewCreateTemplateFromStackRequest()
	}
	response = NewCreateTemplateFromStackResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTemplateRequest() (request *DeleteTemplateRequest) {
	request = &DeleteTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DeleteTemplate")
	return
}

func NewDeleteTemplateResponse() (response *DeleteTemplateResponse) {
	response = &DeleteTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除模板
func (c *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteTemplateRequest()
	}
	response = NewDeleteTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountRequest() (request *DescribeAccountRequest) {
	request = &DescribeAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DescribeAccount")
	return
}

func NewDescribeAccountResponse() (response *DescribeAccountResponse) {
	response = &DescribeAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看账号信息
func (c *Client) DescribeAccount(request *DescribeAccountRequest) (response *DescribeAccountResponse, err error) {
	if request == nil {
		request = NewDescribeAccountRequest()
	}
	response = NewDescribeAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateComplexTemplateRequest() (request *CreateComplexTemplateRequest) {
	request = &CreateComplexTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CreateComplexTemplate")
	return
}

func NewCreateComplexTemplateResponse() (response *CreateComplexTemplateResponse) {
	response = &CreateComplexTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建带内容的模板
func (c *Client) CreateComplexTemplate(request *CreateComplexTemplateRequest) (response *CreateComplexTemplateResponse, err error) {
	if request == nil {
		request = NewCreateComplexTemplateRequest()
	}
	response = NewCreateComplexTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadStackVersionAllFilesRequest() (request *DownloadStackVersionAllFilesRequest) {
	request = &DownloadStackVersionAllFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DownloadStackVersionAllFiles")
	return
}

func NewDownloadStackVersionAllFilesResponse() (response *DownloadStackVersionAllFilesResponse) {
	response = &DownloadStackVersionAllFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载版本所有tf文件
func (c *Client) DownloadStackVersionAllFiles(request *DownloadStackVersionAllFilesRequest) (response *DownloadStackVersionAllFilesResponse, err error) {
	if request == nil {
		request = NewDownloadStackVersionAllFilesRequest()
	}
	response = NewDownloadStackVersionAllFilesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateValidateEventRequest() (request *CreateValidateEventRequest) {
	request = &CreateValidateEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CreateValidateEvent")
	return
}

func NewCreateValidateEventResponse() (response *CreateValidateEventResponse) {
	response = &CreateValidateEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建语法检测事件
func (c *Client) CreateValidateEvent(request *CreateValidateEventRequest) (response *CreateValidateEventResponse, err error) {
	if request == nil {
		request = NewCreateValidateEventRequest()
	}
	response = NewCreateValidateEventResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResourceStackRequest() (request *ModifyResourceStackRequest) {
	request = &ModifyResourceStackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "ModifyResourceStack")
	return
}

func NewModifyResourceStackResponse() (response *ModifyResourceStackResponse) {
	response = &ModifyResourceStackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资源栈
func (c *Client) ModifyResourceStack(request *ModifyResourceStackRequest) (response *ModifyResourceStackResponse, err error) {
	if request == nil {
		request = NewModifyResourceStackRequest()
	}
	response = NewModifyResourceStackResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteStackVersionRequest() (request *DeleteStackVersionRequest) {
	request = &DeleteStackVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "DeleteStackVersion")
	return
}

func NewDeleteStackVersionResponse() (response *DeleteStackVersionResponse) {
	response = &DeleteStackVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资源栈版本
func (c *Client) DeleteStackVersion(request *DeleteStackVersionRequest) (response *DeleteStackVersionResponse, err error) {
	if request == nil {
		request = NewDeleteStackVersionRequest()
	}
	response = NewDeleteStackVersionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDraftRequest() (request *CreateDraftRequest) {
	request = &CreateDraftRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tic", APIVersion, "CreateDraft")
	return
}

func NewCreateDraftResponse() (response *CreateDraftResponse) {
	response = &CreateDraftResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建草稿
func (c *Client) CreateDraft(request *CreateDraftRequest) (response *CreateDraftResponse, err error) {
	if request == nil {
		request = NewCreateDraftRequest()
	}
	response = NewCreateDraftResponse()
	err = c.Send(request, response)
	return
}
