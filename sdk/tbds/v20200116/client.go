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

package v20200116

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-01-16"

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

func NewCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "CreateCluster")
	return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建TBDS集群，申请CVM机器，并在此机器上部署TBDS集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	response = NewCreateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewTestApplyBmsRequest() (request *TestApplyBmsRequest) {
	request = &TestApplyBmsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "TestApplyBms")
	return
}

func NewTestApplyBmsResponse() (response *TestApplyBmsResponse) {
	response = &TestApplyBmsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bms创建实例接口测试
func (c *Client) TestApplyBms(request *TestApplyBmsRequest) (response *TestApplyBmsResponse, err error) {
	if request == nil {
		request = NewTestApplyBmsRequest()
	}
	response = NewTestApplyBmsResponse()
	err = c.Send(request, response)
	return
}

func NewExpandClusterRequest() (request *ExpandClusterRequest) {
	request = &ExpandClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "ExpandCluster")
	return
}

func NewExpandClusterResponse() (response *ExpandClusterResponse) {
	response = &ExpandClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扩容tbds master 和worker 集群实例节点
func (c *Client) ExpandCluster(request *ExpandClusterRequest) (response *ExpandClusterResponse, err error) {
	if request == nil {
		request = NewExpandClusterRequest()
	}
	response = NewExpandClusterResponse()
	err = c.Send(request, response)
	return
}

func NewClusterExistsRequest() (request *ClusterExistsRequest) {
	request = &ClusterExistsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "ClusterExists")
	return
}

func NewClusterExistsResponse() (response *ClusterExistsResponse) {
	response = &ClusterExistsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) ClusterExists(request *ClusterExistsRequest) (response *ClusterExistsResponse, err error) {
	if request == nil {
		request = NewClusterExistsRequest()
	}
	response = NewClusterExistsResponse()
	err = c.Send(request, response)
	return
}

func NewGetInstallInfoRequest() (request *GetInstallInfoRequest) {
	request = &GetInstallInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "GetInstallInfo")
	return
}

func NewGetInstallInfoResponse() (response *GetInstallInfoResponse) {
	response = &GetInstallInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群变更
func (c *Client) GetInstallInfo(request *GetInstallInfoRequest) (response *GetInstallInfoResponse, err error) {
	if request == nil {
		request = NewGetInstallInfoRequest()
	}
	response = NewGetInstallInfoResponse()
	err = c.Send(request, response)
	return
}

func NewBmsCreateClusterRequest() (request *BmsCreateClusterRequest) {
	request = &BmsCreateClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "BmsCreateCluster")
	return
}

func NewBmsCreateClusterResponse() (response *BmsCreateClusterResponse) {
	response = &BmsCreateClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建TBDS集群，申请BMS机器，并在此机器上部署TBDS集群.
func (c *Client) BmsCreateCluster(request *BmsCreateClusterRequest) (response *BmsCreateClusterResponse, err error) {
	if request == nil {
		request = NewBmsCreateClusterRequest()
	}
	response = NewBmsCreateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
	request = &DescribeClusterDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "DescribeClusterDetail")
	return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
	response = &DescribeClusterDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群详情
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
	if request == nil {
		request = NewDescribeClusterDetailRequest()
	}
	response = NewDescribeClusterDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNodesRequest() (request *DescribeClusterNodesRequest) {
	request = &DescribeClusterNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "DescribeClusterNodes")
	return
}

func NewDescribeClusterNodesResponse() (response *DescribeClusterNodesResponse) {
	response = &DescribeClusterNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群节点.
func (c *Client) DescribeClusterNodes(request *DescribeClusterNodesRequest) (response *DescribeClusterNodesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNodesRequest()
	}
	response = NewDescribeClusterNodesResponse()
	err = c.Send(request, response)
	return
}

func NewGetSystemArchRequest() (request *GetSystemArchRequest) {
	request = &GetSystemArchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "GetSystemArch")
	return
}

func NewGetSystemArchResponse() (response *GetSystemArchResponse) {
	response = &GetSystemArchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取一云多芯时所选的架构信息
func (c *Client) GetSystemArch(request *GetSystemArchRequest) (response *GetSystemArchResponse, err error) {
	if request == nil {
		request = NewGetSystemArchRequest()
	}
	response = NewGetSystemArchResponse()
	err = c.Send(request, response)
	return
}

func NewBmsDescribeInstancesByNameRequest() (request *BmsDescribeInstancesByNameRequest) {
	request = &BmsDescribeInstancesByNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "BmsDescribeInstancesByName")
	return
}

func NewBmsDescribeInstancesByNameResponse() (response *BmsDescribeInstancesByNameResponse) {
	response = &BmsDescribeInstancesByNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据实例名称获取bms实例信息
func (c *Client) BmsDescribeInstancesByName(request *BmsDescribeInstancesByNameRequest) (response *BmsDescribeInstancesByNameResponse, err error) {
	if request == nil {
		request = NewBmsDescribeInstancesByNameRequest()
	}
	response = NewBmsDescribeInstancesByNameResponse()
	err = c.Send(request, response)
	return
}

func NewExpandClusterNewRequest() (request *ExpandClusterNewRequest) {
	request = &ExpandClusterNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "ExpandClusterNew")
	return
}

func NewExpandClusterNewResponse() (response *ExpandClusterNewResponse) {
	response = &ExpandClusterNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扩容tbds master 和worker 集群实例节点
func (c *Client) ExpandClusterNew(request *ExpandClusterNewRequest) (response *ExpandClusterNewResponse, err error) {
	if request == nil {
		request = NewExpandClusterNewRequest()
	}
	response = NewExpandClusterNewResponse()
	err = c.Send(request, response)
	return
}

func NewGetInstallStepsRequest() (request *GetInstallStepsRequest) {
	request = &GetInstallStepsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "GetInstallSteps")
	return
}

func NewGetInstallStepsResponse() (response *GetInstallStepsResponse) {
	response = &GetInstallStepsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安装已完成步骤
func (c *Client) GetInstallSteps(request *GetInstallStepsRequest) (response *GetInstallStepsResponse, err error) {
	if request == nil {
		request = NewGetInstallStepsRequest()
	}
	response = NewGetInstallStepsResponse()
	err = c.Send(request, response)
	return
}

func NewGetProgressRequest() (request *GetProgressRequest) {
	request = &GetProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "GetProgress")
	return
}

func NewGetProgressResponse() (response *GetProgressResponse) {
	response = &GetProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) GetProgress(request *GetProgressRequest) (response *GetProgressResponse, err error) {
	if request == nil {
		request = NewGetProgressRequest()
	}
	response = NewGetProgressResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyClusterRequest() (request *DestroyClusterRequest) {
	request = &DestroyClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "DestroyCluster")
	return
}

func NewDestroyClusterResponse() (response *DestroyClusterResponse) {
	response = &DestroyClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁集群
func (c *Client) DestroyCluster(request *DestroyClusterRequest) (response *DestroyClusterResponse, err error) {
	if request == nil {
		request = NewDestroyClusterRequest()
	}
	response = NewDestroyClusterResponse()
	err = c.Send(request, response)
	return
}

func NewGetClustersRequest() (request *GetClustersRequest) {
	request = &GetClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbds", APIVersion, "GetClusters")
	return
}

func NewGetClustersResponse() (response *GetClustersResponse) {
	response = &GetClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) GetClusters(request *GetClustersRequest) (response *GetClustersResponse, err error) {
	if request == nil {
		request = NewGetClustersRequest()
	}
	response = NewGetClustersResponse()
	err = c.Send(request, response)
	return
}
