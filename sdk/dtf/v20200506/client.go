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

package v20200506

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-05-06"

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

func NewDescribeBranchesRequest() (request *DescribeBranchesRequest) {
	request = &DescribeBranchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "DescribeBranches")
	return
}

func NewDescribeBranchesResponse() (response *DescribeBranchesResponse) {
	response = &DescribeBranchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子事务列表
func (c *Client) DescribeBranches(request *DescribeBranchesRequest) (response *DescribeBranchesResponse, err error) {
	if request == nil {
		request = NewDescribeBranchesRequest()
	}
	response = NewDescribeBranchesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBranchStatisticsRequest() (request *DescribeBranchStatisticsRequest) {
	request = &DescribeBranchStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "DescribeBranchStatistics")
	return
}

func NewDescribeBranchStatisticsResponse() (response *DescribeBranchStatisticsResponse) {
	response = &DescribeBranchStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询分支事务统计信息
func (c *Client) DescribeBranchStatistics(request *DescribeBranchStatisticsRequest) (response *DescribeBranchStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeBranchStatisticsRequest()
	}
	response = NewDescribeBranchStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewRetryAllTransactionsRequest() (request *RetryAllTransactionsRequest) {
	request = &RetryAllTransactionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "RetryAllTransactions")
	return
}

func NewRetryAllTransactionsResponse() (response *RetryAllTransactionsResponse) {
	response = &RetryAllTransactionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 触发重试全部异常主事务
func (c *Client) RetryAllTransactions(request *RetryAllTransactionsRequest) (response *RetryAllTransactionsResponse, err error) {
	if request == nil {
		request = NewRetryAllTransactionsRequest()
	}
	response = NewRetryAllTransactionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTransactionsRequest() (request *DescribeTransactionsRequest) {
	request = &DescribeTransactionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "DescribeTransactions")
	return
}

func NewDescribeTransactionsResponse() (response *DescribeTransactionsResponse) {
	response = &DescribeTransactionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主事务列表
func (c *Client) DescribeTransactions(request *DescribeTransactionsRequest) (response *DescribeTransactionsResponse, err error) {
	if request == nil {
		request = NewDescribeTransactionsRequest()
	}
	response = NewDescribeTransactionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllGroupStatisticsRequest() (request *DescribeAllGroupStatisticsRequest) {
	request = &DescribeAllGroupStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "DescribeAllGroupStatistics")
	return
}

func NewDescribeAllGroupStatisticsResponse() (response *DescribeAllGroupStatisticsResponse) {
	response = &DescribeAllGroupStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有事务分组统计信息
func (c *Client) DescribeAllGroupStatistics(request *DescribeAllGroupStatisticsRequest) (response *DescribeAllGroupStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeAllGroupStatisticsRequest()
	}
	response = NewDescribeAllGroupStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
	request = &DeleteGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "DeleteGroup")
	return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
	response = &DeleteGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除事务分组
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
	if request == nil {
		request = NewDeleteGroupRequest()
	}
	response = NewDeleteGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
	request = &ModifyGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "ModifyGroup")
	return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
	response = &ModifyGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新事务分组
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
	if request == nil {
		request = NewModifyGroupRequest()
	}
	response = NewModifyGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupsRequest() (request *DescribeGroupsRequest) {
	request = &DescribeGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "DescribeGroups")
	return
}

func NewDescribeGroupsResponse() (response *DescribeGroupsResponse) {
	response = &DescribeGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询事务分组
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeGroupsRequest()
	}
	response = NewDescribeGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
	request = &CreateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "CreateGroup")
	return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
	response = &CreateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建事务分组
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
	if request == nil {
		request = NewCreateGroupRequest()
	}
	response = NewCreateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupStatisticsRequest() (request *DescribeGroupStatisticsRequest) {
	request = &DescribeGroupStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "DescribeGroupStatistics")
	return
}

func NewDescribeGroupStatisticsResponse() (response *DescribeGroupStatisticsResponse) {
	response = &DescribeGroupStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询事务分组统计信息
func (c *Client) DescribeGroupStatistics(request *DescribeGroupStatisticsRequest) (response *DescribeGroupStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeGroupStatisticsRequest()
	}
	response = NewDescribeGroupStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewRetryTransactionsRequest() (request *RetryTransactionsRequest) {
	request = &RetryTransactionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dtf", APIVersion, "RetryTransactions")
	return
}

func NewRetryTransactionsResponse() (response *RetryTransactionsResponse) {
	response = &RetryTransactionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 触发重试异常主事务
func (c *Client) RetryTransactions(request *RetryTransactionsRequest) (response *RetryTransactionsResponse, err error) {
	if request == nil {
		request = NewRetryTransactionsRequest()
	}
	response = NewRetryTransactionsResponse()
	err = c.Send(request, response)
	return
}
