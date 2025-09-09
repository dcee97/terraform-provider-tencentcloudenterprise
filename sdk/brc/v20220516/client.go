package v20220516

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

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2022-05-16"

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

// NewActivateBrcServiceRequest 创建 ActivateBrcService 请求
func NewActivateBrcServiceRequest() *ActivateBackupServiceRequest {
	request := &ActivateBackupServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "ActivateBackupService")
	return request
}

// NewActivateBrcServiceResponse 创建 ActivateBrcService 返回
func NewActivateBrcServiceResponse() *ActivateBackupServiceResponse {
	return &ActivateBackupServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// ActivateBackupService 激活备份服务
func (c *Client) ActivateBackupService(request *ActivateBackupServiceRequest) (response *ActivateBackupServiceResponse, err error) {
	if request == nil {
		request = NewActivateBrcServiceRequest()
	}
	response = NewActivateBrcServiceResponse()
	err = c.Send(request, response)
	return
}

// NewDescribeResourceBackupOverviewRequest 创建 DescribeResourceBackupOverview 请求
func NewDescribeResourceBackupOverviewRequest() *DescribeResourceBackupOverviewRequest {
	request := &DescribeResourceBackupOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeResourceBackupOverview")
	return request
}

// NewDescribeResourceBackupOverviewResponse 创建 DescribeResourceBackupOverview 返回
func NewDescribeResourceBackupOverviewResponse() *DescribeResourceBackupOverviewResponse {
	return &DescribeResourceBackupOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// DescribeResourceBackupOverview 获取备份概览
func (c *Client) DescribeResourceBackupOverview(request *DescribeResourceBackupOverviewRequest) (response *DescribeResourceBackupOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBackupOverviewRequest()
	}
	response = NewDescribeResourceBackupOverviewResponse()
	err = c.Send(request, response)
	return
}

// NewCreateBackupCvmResourceRequest 创建 CreateBackupCvmResource 请求
func NewCreateBackupCvmResourceRequest() *CreateBackupCvmResourceRequest {
	request := &CreateBackupCvmResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateBackupGroup")
	return request
}

// NewCreateBackupCvmResourceResponse 创建 CreateBackupCvmResource 返回
func NewCreateBackupCvmResourceResponse() *CreateBackupCvmResourceResponse {
	return &CreateBackupCvmResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// CreateBackupCvmResource 创建备份组
func (c *Client) CreateBackupCvmResource(request *CreateBackupCvmResourceRequest) (response *CreateBackupCvmResourceResponse, err error) {
	if request == nil {
		request = NewCreateBackupCvmResourceRequest()
	}
	response = NewCreateBackupCvmResourceResponse()
	err = c.Send(request, response)
	return
}

// NewDescribeBackupCvmResourceRequest 创建 DescribeBackupCvmResource 请求
func NewDescribeBackupCvmResourceRequest() *DescribeBackupCvmResourceRequest {
	request := &DescribeBackupCvmResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackupGroups")
	return request
}

// NewDescribeBackupCvmResourceResponse 创建 DescribeBackupCvmResource 返回
func NewDescribeBackupCvmResourceResponse() *DescribeBackupCvmResourceResponse {
	return &DescribeBackupCvmResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// DescribeBackupCvmResource 获取备份组
func (c *Client) DescribeBackupCvmResource(request *DescribeBackupCvmResourceRequest) (
	response *DescribeBackupCvmResourceResponse, err error) {
	if request == nil {
		request = NewDescribeBackupCvmResourceRequest()
	}
	response = NewDescribeBackupCvmResourceResponse()
	err = c.Send(request, response)
	return
}

// NewDeleteBackupCvmResourceRequest 创建 DeleteBackupCvmResource 请求
func NewDeleteBackupCvmResourceRequest() *DeleteBackupCvmResourceRequest {
	request := &DeleteBackupCvmResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DeleteBackupGroups")
	return request
}

// NewDeleteBackupCvmResourceResponse 创建 DeleteBackupCvmResource 返回
func NewDeleteBackupCvmResourceResponse() *DeleteBackupCvmResourceResponse {
	return &DeleteBackupCvmResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// DeleteBackupCvmResource 删除备份组
func (c *Client) DeleteBackupCvmResource(request *DeleteBackupCvmResourceRequest) (response *DeleteBackupCvmResourceResponse, err error) {
	if request == nil {
		request = NewDeleteBackupCvmResourceRequest()
	}
	response = NewDeleteBackupCvmResourceResponse()
	err = c.Send(request, response)
	return
}

// NewCreateBackupRequest 创建 CreateBackup 请求
func NewCreateBackupRequest() *CreateBackupRequest {
	request := &CreateBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateBackup")
	return request
}

// NewCreateBackupResponse 创建 CreateBackup 返回
func NewCreateBackupResponse() *CreateBackupResponse {
	return &CreateBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// CreateBackup 创建备份
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
	if request == nil {
		request = NewCreateBackupRequest()
	}
	response = NewCreateBackupResponse()
	err = c.Send(request, response)
	return
}

// NewDescribeBackupsRequest 创建 DescribeBackups 请求
func NewDescribeBackupsRequest() *DescribeBackupsRequest {
	request := &DescribeBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeBackups")
	return request
}

// NewDescribeBackupsResponse 创建 DescribeBackups 返回
func NewDescribeBackupsResponse() *DescribeBackupsResponse {
	return &DescribeBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// DescribeBackups 获取备份列表
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupsRequest()
	}
	response = NewDescribeBackupsResponse()
	err = c.Send(request, response)
	return
}

// NewDeleteBackupsRequest 创建 DeleteBackups 请求
func NewDeleteBackupsRequest() *DeleteBackupsRequest {
	request := &DeleteBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DeleteBackups")
	return request
}

// NewDeleteBackupsResponse 创建 DeleteBackups 返回
func NewDeleteBackupsResponse() *DeleteBackupsResponse {
	return &DeleteBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// DeleteBackups 删除备份
func (c *Client) DeleteBackups(request *DeleteBackupsRequest) (response *DeleteBackupsResponse, err error) {
	if request == nil {
		request = NewDeleteBackupsRequest()
	}
	response = NewDeleteBackupsResponse()
	err = c.Send(request, response)
	return
}

// NewCreateCfsBackupRequest 创建 CreateCfsBackup 请求
func NewCreateCfsBackupRequest() *CreateCfsBackupRequest {
	request := &CreateCfsBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateCfsBackup")
	return request
}

// NewCreateCfsBackupResponse 创建 CreateCfsBackup 返回
func NewCreateCfsBackupResponse() *CreateCfsBackupResponse {
	return &CreateCfsBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// CreateCfsBackup 创建CFS备份
func (c *Client) CreateCfsBackup(request *CreateCfsBackupRequest) (response *CreateCfsBackupResponse, err error) {
	if request == nil {
		request = NewCreateCfsBackupRequest()
	}
	response = NewCreateCfsBackupResponse()
	err = c.Send(request, response)
	return
}

// NewDescribeCfsBackupsRequest 创建 DescribeCfsBackups 请求
func NewDescribeCfsBackupsRequest() *DescribeCfsBackupsRequest {
	request := &DescribeCfsBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeCfsBackups")
	return request
}

// NewDescribeCfsBackupsResponse 创建 DescribeCfsBackups 返回
func NewDescribeCfsBackupsResponse() *DescribeCfsBackupsResponse {
	return &DescribeCfsBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

// DescribeCfsBackups 获取CFS备份列表
func (c *Client) DescribeCfsBackups(request *DescribeCfsBackupsRequest) (response *DescribeCfsBackupsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsBackupsRequest()
	}
	response = NewDescribeCfsBackupsResponse()
	err = c.Send(request, response)
	return
}

// CreateResourceBackup
// 创建COS/CSP资源备份
func (c *Client) CreateResourceBackup(request *CreateResourceBackupRequest) (response *CreateResourceBackupResponse, err error) {
	if request == nil {
		request = NewCreateResourceBackupRequest()
	}
	response = NewCreateResourceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateResourceBackupRequest() (request *CreateResourceBackupRequest) {
	request = &CreateResourceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateResourceBackup")
	return
}

func NewCreateResourceBackupResponse() (response *CreateResourceBackupResponse) {
	response = &CreateResourceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeResourceBackups
// 查询COS/CSP资源备份列表
func (c *Client) DescribeResourceBackups(request *DescribeResourceBackupsRequest) (response *DescribeResourceBackupsResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBackupsRequest()
	}
	response = NewDescribeResourceBackupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBackupsRequest() (request *DescribeResourceBackupsRequest) {
	request = &DescribeResourceBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeResourceBackups")
	return
}

func NewDescribeResourceBackupsResponse() (response *DescribeResourceBackupsResponse) {
	response = &DescribeResourceBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteResourceBackups
// 删除COS/CSP资源备份
func (c *Client) DeleteResourceBackups(request *DeleteResourceBackupsRequest) (response *DeleteResourceBackupsResponse, err error) {
	if request == nil {
		request = NewDeleteResourceBackupsRequest()
	}
	response = NewDeleteResourceBackupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteResourceBackupsRequest() (request *DeleteResourceBackupsRequest) {
	request = &DeleteResourceBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DeleteResourceBackups")
	return
}

func NewDeleteResourceBackupsResponse() (response *DeleteResourceBackupsResponse) {
	response = &DeleteResourceBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateDatabaseResourceBackup
// 创建数据库资源备份
func (c *Client) CreateDatabaseResourceBackup(request *CreateDatabaseResourceBackupRequest) (response *CreateDatabaseResourceBackupResponse, err error) {
	if request == nil {
		request = NewCreateDatabaseResourceBackupRequest()
	}
	response = NewCreateDatabaseResourceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDatabaseResourceBackupRequest() (request *CreateDatabaseResourceBackupRequest) {
	request = &CreateDatabaseResourceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateResourceBackup")
	return
}

func NewCreateDatabaseResourceBackupResponse() (response *CreateDatabaseResourceBackupResponse) {
	response = &CreateDatabaseResourceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Auto Backup Policy Client Methods
func NewCreateAutoBackupPolicyRequest() *CreateAutoBackupPolicyRequest {
	request := &CreateAutoBackupPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "CreateAutoBackupPolicy")
	return request
}

func NewCreateAutoBackupPolicyResponse() *CreateAutoBackupPolicyResponse {
	return &CreateAutoBackupPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

func (c *Client) CreateAutoBackupPolicy(request *CreateAutoBackupPolicyRequest) (response *CreateAutoBackupPolicyResponse, err error) {
	if request == nil {
		request = NewCreateAutoBackupPolicyRequest()
	}
	response = NewCreateAutoBackupPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoBackupPoliciesRequest() *DescribeAutoBackupPoliciesRequest {
	request := &DescribeAutoBackupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DescribeAutoBackupPolicies")
	return request
}

func NewDescribeAutoBackupPoliciesResponse() *DescribeAutoBackupPoliciesResponse {
	return &DescribeAutoBackupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

func (c *Client) DescribeAutoBackupPolicies(request *DescribeAutoBackupPoliciesRequest) (response *DescribeAutoBackupPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeAutoBackupPoliciesRequest()
	}
	response = NewDescribeAutoBackupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoBackupPolicyAttributeRequest() *ModifyAutoBackupPolicyAttributeRequest {
	request := &ModifyAutoBackupPolicyAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "ModifyAutoBackupPolicyAttribute")
	return request
}

func NewModifyAutoBackupPolicyAttributeResponse() *ModifyAutoBackupPolicyAttributeResponse {
	return &ModifyAutoBackupPolicyAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

func (c *Client) ModifyAutoBackupPolicyAttribute(request *ModifyAutoBackupPolicyAttributeRequest) (response *ModifyAutoBackupPolicyAttributeResponse, err error) {
	if request == nil {
		request = NewModifyAutoBackupPolicyAttributeRequest()
	}
	response = NewModifyAutoBackupPolicyAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAutoBackupPoliciesRequest() *DeleteAutoBackupPoliciesRequest {
	request := &DeleteAutoBackupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "DeleteAutoBackupPolicies")
	return request
}

func NewDeleteAutoBackupPoliciesResponse() *DeleteAutoBackupPoliciesResponse {
	return &DeleteAutoBackupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

func (c *Client) DeleteAutoBackupPolicies(request *DeleteAutoBackupPoliciesRequest) (response *DeleteAutoBackupPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteAutoBackupPoliciesRequest()
	}
	response = NewDeleteAutoBackupPoliciesResponse()
	err = c.Send(request, response)
	return
}

// BindAutoBackupPolicy
func NewBindAutoBackupPolicyRequest() *BindAutoBackupPolicyRequest {
	request := &BindAutoBackupPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "BindAutoBackupPolicy")

	return request
}

func NewBindAutoBackupPolicyResponse() *BindAutoBackupPolicyResponse {
	return &BindAutoBackupPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

func (c *Client) BindAutoBackupPolicy(request *BindAutoBackupPolicyRequest) (response *BindAutoBackupPolicyResponse, err error) {
	if request == nil {
		request = NewBindAutoBackupPolicyRequest()
	}
	response = NewBindAutoBackupPolicyResponse()
	err = c.Send(request, response)
	return response, nil
}

// UnbindAutoBackupPolicy
func NewUnbindAutoBackupPolicyRequest() *UnbindAutoBackupPolicyRequest {
	request := &UnbindAutoBackupPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("brc", APIVersion, "UnbindAutoBackupPolicy")

	return request
}

func NewUnbindAutoBackupPolicyResponse() *UnbindAutoBackupPolicyResponse {
	return &UnbindAutoBackupPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
}

func (c *Client) UnbindAutoBackupPolicy(request *UnbindAutoBackupPolicyRequest) (response *UnbindAutoBackupPolicyResponse, err error) {
	if request == nil {
		request = NewUnbindAutoBackupPolicyRequest()
	}
	response = NewUnbindAutoBackupPolicyResponse()
	err = c.Send(request, response)
	return response, nil
}
