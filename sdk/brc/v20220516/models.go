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

package v20220516

import (
	"encoding/json"
	"time"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ActivateBackupServiceRequest struct {
	*tchttp.BaseRequest
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *ActivateBackupServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateBackupServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateBackupServiceResponse struct {
	*tchttp.BaseResponse
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	Response  *struct {
		BrcServiceStatus *string `json:"BrcServiceStatus,omitempty" name:"BrcServiceStatus"`
	}
}

func (r *ActivateBackupServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateBackupServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBackupOverviewRequest struct {
	*tchttp.BaseRequest
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeResourceBackupOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBackupOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupOverview struct {
	ResourceType        *string `json:"ResourceType,omitempty" name:"ResourceType"`
	BackupCount         *uint64 `json:"BackupCount,omitempty" name:"BackupCount"`
	BackupResourceCount *uint64 `json:"BackupResourceCount,omitempty" name:"BackupResourceCount"`
	BackupSizeMb        *uint64 `json:"BackupSizeMb,omitempty" name:"BackupSizeMb"`
}

type ResourceOverview struct {
	BackupResourceCount *uint64 `json:"BackupResourceCount,omitempty" name:"BackupResourceCount"`
	BackupCount         *uint64 `json:"BackupCount,omitempty" name:"BackupCount"`
	BackupSizeMb        *uint64 `json:"BackupSizeMb,omitempty" name:"BackupSizeMb"`
}

type ResourceOverviewDetail struct {
	ResourceType     *string           `json:"ResourceType,omitempty" name:"ResourceType"`
	IsOpen           *bool             `json:"IsOpen,omitempty" name:"IsOpen"`
	IsSupport        *bool             `json:"IsSupport,omitempty" name:"IsSupport"`
	Region           *string           `json:"Region,omitempty" name:"Region"`
	ResourceOverview *ResourceOverview `json:"ResourceOverview,omitempty" name:"ResourceOverview"`
}

type DescribeResourceBackupOverviewResponse struct {
	*tchttp.BaseResponse
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	Response  *struct {
		BackupOverview    []*BackupOverview         `json:"BackupOverview,omitempty" name:"BackupOverview"`
		OverviewDetailSet []*ResourceOverviewDetail `json:"OverviewDetailSet,omitempty" name:"OverviewDetailSet"`
	}
}

func (r *DescribeResourceBackupOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupCvmResourceRequest struct {
	*tchttp.BaseRequest
	DiskIds         []*string  `json:"DiskIds,omitempty" name:"DiskIds"`
	BackupGroupName *string    `json:"BackupGroupName,omitempty" name:"BackupGroupName"`
	Deadline        *time.Time `json:"Deadline,omitempty" name:"Deadline"`
	CreateSpeed     *uint64    `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	NeedArchive     *bool      `json:"NeedArchive,omitempty" name:"NeedArchive"`
	BackupClass     *string    `json:"BackupClass,omitempty" name:"BackupClass"`
}

func (r *CreateBackupCvmResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupCvmResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupCvmResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		BackupGroupId *string `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
		RequestId     *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupCvmResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupCvmResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupBackup struct {
	DiskIds     []*string          `json:"DiskIds,omitempty" name:"DiskIds"`
	BackupName  *string            `json:"BackupName,omitempty" name:"BackupName"`
	Deadline    *time.Time         `json:"Deadline,omitempty" name:"Deadline"`
	BackupClass *string            `json:"BackupClass,omitempty" name:"BackupClass"`
	CreateSpeed *uint64            `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	NeedArchive *bool              `json:"NeedArchive,omitempty" name:"NeedArchive"`
	RemoteInfo  *RemoteStorageInfo `json:"RemoteInfo,omitempty" name:"RemoteInfo"`
}

type BackupGroup struct {
	ContainRootBackup             *bool     `json:"ContainRootBackup,omitempty" name:"ContainRootBackup"`
	BackupGroupId                 *string   `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
	BackupGroupType               *string   `json:"BackupGroupType,omitempty" name:"BackupGroupType"`
	Percent                       *uint64   `json:"Percent,omitempty" name:"Percent"`
	BackupIdSet                   []*string `json:"BackupIdSet,omitempty" name:"BackupIdSet"`
	BackupGroupName               *string   `json:"BackupGroupName,omitempty" name:"BackupGroupName"`
	BackupGroupState              *string   `json:"BackupGroupState,omitempty" name:"BackupGroupState"`
	ModifyTime                    *string   `json:"ModifyTime,omitempty" name:"ModifyTime"`
	CreateTime                    *string   `json:"CreateTime,omitempty" name:"CreateTime"`
	AppId                         *uint64   `json:"AppId,omitempty" name:"AppId"`
	IsPermanent                   *bool     `json:"IsPermanent,omitempty" name:"IsPermanent"`
	DeadlineTime                  *string   `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	InstanceId                    *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceDetails               *string   `json:"InstanceDetails,omitempty" name:"InstanceDetails"`
	AccountName                   *string   `json:"AccountName,omitempty" name:"AccountName"`
	AccountUin                    *string   `json:"AccountUin,omitempty" name:"AccountUin"`
	SubAccountUin                 *string   `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	AutoBackupPolicyId            *string   `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	ArchiveStatus                 *string   `json:"ArchiveStatus,omitempty" name:"ArchiveStatus"`
	LatestOperationState          *string   `json:"LatestOperationState,omitempty" name:"LatestOperationState"`
	ErrorPrompt                   *string   `json:"ErrorPrompt,omitempty" name:"ErrorPrompt"`
	LatestOperationRequestId      *string   `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`
	OperationSuccessResourceIdSet []*string `json:"OperationSuccessResourceIdSet,omitempty" name:"OperationSuccessResourceIdSet"`
	OperationFailedResourceIdSet  []*string `json:"OperationFailedResourceIdSet,omitempty" name:"OperationFailedResourceIdSet"`
	LatestOperation               *string   `json:"LatestOperation,omitempty" name:"LatestOperation"`
}

type Filter struct {
	Name   *string   `json:"Name,omitempty" name:"Name"`
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeBackupCvmResourceRequest struct {
	*tchttp.BaseRequest
	Filters    []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset     *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit      *uint64   `json:"Limit,omitempty" name:"Limit"`
	Order      *string   `json:"Order,omitempty" name:"Order"`
	OrderField *string   `json:"OrderField,omitempty" name:"OrderField"`
}

type DescribeBackupCvmResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		TotalCount     *uint64        `json:"TotalCount,omitempty" name:"TotalCount"`
		BackupGroupSet []*BackupGroup `json:"BackupGroupSet,omitempty" name:"BackupGroupSet"`
		RequestId      *string        `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type DeleteBackupCvmResourceRequest struct {
	*tchttp.BaseRequest
	BackupGroupIds []*string `json:"BackupGroupIds,omitempty" name:"BackupGroupIds"`
	OnlyDismiss    *bool     `json:"OnlyDismiss,omitempty" name:"OnlyDismiss"`
}

func (r *DeleteBackupCvmResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}
func (r *DeleteBackupCvmResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupCvmResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupCvmResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupCvmResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// CBS Backup Models
type CreateBackupRequest struct {
	*tchttp.BaseRequest
	DiskId      *string    `json:"DiskId,omitempty" name:"DiskId"`
	BackupName  *string    `json:"BackupName,omitempty" name:"BackupName"`
	Deadline    *time.Time `json:"Deadline,omitempty" name:"Deadline"`
	BackupClass *string    `json:"BackupClass,omitempty" name:"BackupClass"`
	CreateSpeed *uint64    `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	NeedArchive *bool      `json:"NeedArchive,omitempty" name:"NeedArchive"`
}

func (r *CreateBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		BackupId  *string `json:"BackupId,omitempty" name:"BackupId"`
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Backup struct {
	SubAccountUin      *string    `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	Encrypt            *bool      `json:"Encrypt,omitempty" name:"Encrypt"`
	BackupName         *string    `json:"BackupName,omitempty" name:"BackupName"`
	Percent            *uint64    `json:"Percent,omitempty" name:"Percent"`
	ShareReference     *uint64    `json:"ShareReference,omitempty" name:"ShareReference"`
	AccountUin         *string    `json:"AccountUin,omitempty" name:"AccountUin"`
	DiskDetails        *string    `json:"DiskDetails,omitempty" name:"DiskDetails"`
	PlatformProjectId  *string    `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	CopyingToRegions   []*string  `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`
	BackupClass        *string    `json:"BackupClass,omitempty" name:"BackupClass"`
	DiskName           *string    `json:"DiskName,omitempty" name:"DiskName"`
	Tags               []*Tag     `json:"Tags,omitempty" name:"Tags"`
	IsPermanent        *bool      `json:"IsPermanent,omitempty" name:"IsPermanent"`
	BackupGroupId      *string    `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
	DiskId             *string    `json:"DiskId,omitempty" name:"DiskId"`
	Placement          *Placement `json:"Placement,omitempty" name:"Placement"`
	CopyFromRemote     *bool      `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`
	DeadlineTime       *string    `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	ArchiveStatus      *string    `json:"ArchiveStatus,omitempty" name:"ArchiveStatus"`
	AccountName        *string    `json:"AccountName,omitempty" name:"AccountName"`
	BackupType         *string    `json:"BackupType,omitempty" name:"BackupType"`
	DiskSize           *uint64    `json:"DiskSize,omitempty" name:"DiskSize"`
	DiskUsage          *string    `json:"DiskUsage,omitempty" name:"DiskUsage"`
	AutoBackupPolicyId *string    `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	BackupState        *string    `json:"BackupState,omitempty" name:"BackupState"`
	AppId              *uint64    `json:"AppId,omitempty" name:"AppId"`
	BackupId           *string    `json:"BackupId,omitempty" name:"BackupId"`
	CreateTime         *string    `json:"CreateTime,omitempty" name:"CreateTime"`
	CreateSpeed        *uint64    `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	NeedArchive        *bool      `json:"NeedArchive,omitempty" name:"NeedArchive"`
}

type Placement struct {
	CageId             *string `json:"CageId,omitempty" name:"CageId"`
	Zone               *string `json:"Zone,omitempty" name:"Zone"`
	ProjectId          *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	CdcName            *string `json:"CdcName,omitempty" name:"CdcName"`
	CdcId              *string `json:"CdcId,omitempty" name:"CdcId"`
	ProjectName        *string `json:"ProjectName,omitempty" name:"ProjectName"`
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset  *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit   *uint64   `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		TotalCount *uint64   `json:"TotalCount,omitempty" name:"TotalCount"`
		BackupSet  []*Backup `json:"BackupSet,omitempty" name:"BackupSet"`
		RequestId  *string   `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupsRequest struct {
	*tchttp.BaseRequest
	BackupIds []*string `json:"BackupIds,omitempty" name:"BackupIds"`
}

func (r *DeleteBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// CFS Backup Models
type CreateCfsBackupRequest struct {
	*tchttp.BaseRequest
	FileSystemId *string            `json:"FileSystemId,omitempty" name:"FileSystemId"`
	BackupName   *string            `json:"BackupName,omitempty" name:"BackupName"`
	Deadline     *time.Time         `json:"Deadline,omitempty" name:"Deadline"`
	CreateSpeed  *uint64            `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	NeedArchive  *bool              `json:"NeedArchive,omitempty" name:"NeedArchive"`
	BackupClass  *string            `json:"BackupClass,omitempty" name:"BackupClass"`
	RemoteInfo   *RemoteStorageInfo `json:"RemoteInfo,omitempty" name:"RemoteInfo"`
}

func (r *CreateCfsBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCfsBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		BackupId  *string `json:"BackupId,omitempty" name:"BackupId"`
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCfsBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCfsBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteStorageInfo struct {
	SourceRemoteCloud       *string `json:"SourceRemoteCloud,omitempty" name:"SourceRemoteCloud"`
	SourceRemoteRegion      *string `json:"SourceRemoteRegion,omitempty" name:"SourceRemoteRegion"`
	DestinationRemoteCloud  *string `json:"DestinationRemoteCloud,omitempty" name:"DestinationRemoteCloud"`
	DestinationRemoteRegion *string `json:"DestinationRemoteRegion,omitempty" name:"DestinationRemoteRegion"`
	RemoteAppId             *uint64 `json:"RemoteAppId,omitempty" name:"RemoteAppId"`
	CrossRegionName         *string `json:"CrossRegionName,omitempty" name:"CrossRegionName"`
	IsIsolatedStorage       *bool   `json:"IsIsolatedStorage,omitempty" name:"IsIsolatedStorage"`
}

type CfsBackup struct {
	SubAccountUin      *string    `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	Encrypt            *bool      `json:"Encrypt,omitempty" name:"Encrypt"`
	BackupName         *string    `json:"BackupName,omitempty" name:"BackupName"`
	Percent            *uint64    `json:"Percent,omitempty" name:"Percent"`
	ShareReference     *uint64    `json:"ShareReference,omitempty" name:"ShareReference"`
	AccountUin         *string    `json:"AccountUin,omitempty" name:"AccountUin"`
	PlatformProjectId  *string    `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	CopyingToRegions   []*string  `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`
	BackupClass        *string    `json:"BackupClass,omitempty" name:"BackupClass"`
	DiskName           *string    `json:"DiskName,omitempty" name:"DiskName"`
	Tags               []*Tag     `json:"Tags,omitempty" name:"Tags"`
	IsPermanent        *bool      `json:"IsPermanent,omitempty" name:"IsPermanent"`
	BackupGroupId      *string    `json:"BackupGroupId,omitempty" name:"BackupGroupId"`
	DiskId             *string    `json:"DiskId,omitempty" name:"DiskId"`
	Placement          *Placement `json:"Placement,omitempty" name:"Placement"`
	CopyFromRemote     *bool      `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`
	DeadlineTime       *string    `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	ArchiveStatus      *string    `json:"ArchiveStatus,omitempty" name:"ArchiveStatus"`
	AccountName        *string    `json:"AccountName,omitempty" name:"AccountName"`
	BackupType         *string    `json:"BackupType,omitempty" name:"BackupType"`
	FileSystemId       *string    `json:"FileSystemId,omitempty" name:"FileSystemId"`
	DiskSize           *uint64    `json:"DiskSize,omitempty" name:"DiskSize"`
	DiskUsage          *string    `json:"DiskUsage,omitempty" name:"DiskUsage"`
	AutoBackupPolicyId *string    `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	BackupState        *string    `json:"BackupState,omitempty" name:"BackupState"`
	FileSystemDetails  *string    `json:"FileSystemDetails,omitempty" name:"FileSystemDetails"`
	AppId              *uint64    `json:"AppId,omitempty" name:"AppId"`
	BackupId           *string    `json:"BackupId,omitempty" name:"BackupId"`
	CreateTime         *string    `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeCfsBackupsRequest struct {
	*tchttp.BaseRequest
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset  *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit   *uint64   `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCfsBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfsBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		TotalCount *uint64      `json:"TotalCount,omitempty" name:"TotalCount"`
		BackupSet  []*CfsBackup `json:"BackupSet,omitempty" name:"BackupSet"`
		RequestId  *string      `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfsBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfsBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// COS/CSP Backup Models
type CreateResourceBackupRequest struct {
	*tchttp.BaseRequest
	ResourceType *string       `json:"ResourceType,omitempty" name:"ResourceType"`
	ResourceId   *string       `json:"ResourceId,omitempty" name:"ResourceId"`
	BackupName   *string       `json:"BackupName,omitempty" name:"BackupName"`
	Deadline     *time.Time    `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	CreateSpeed  *uint64       `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	BucketDetail *BucketDetail `json:"BucketDetail,omitempty" name:"BucketDetail"`
}

func (r *CreateResourceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		BackupId  *string `json:"BackupId,omitempty" name:"BackupId"`
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateResourceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BucketDetail struct {
	CosRegion  *string `json:"CosRegion,omitempty" name:"CosRegion"`
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	Prefix     *string `json:"Prefix,omitempty" name:"Prefix"`
}

type ResourceBackup struct {
	ResourceId  *string            `json:"ResourceId,omitempty" name:"FileSystemId"`
	BackupName  *string            `json:"BackupName,omitempty" name:"BackupName"`
	Deadline    *time.Time         `json:"Deadline,omitempty" name:"Deadline"`
	CreateSpeed *uint64            `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	NeedArchive *bool              `json:"NeedArchive,omitempty" name:"NeedArchive"`
	BackupClass *string            `json:"BackupClass,omitempty" name:"BackupClass"`
	RemoteInfo  *RemoteStorageInfo `json:"RemoteInfo,omitempty" name:"RemoteInfo"`
}
type ResourceBackupDetail struct {
	SubAccountUin       *string       `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	BackupName          *string       `json:"BackupName,omitempty" name:"BackupName"`
	InstanceDetails     *string       `json:"InstanceDetails,omitempty" name:"InstanceDetails"`
	Percent             *uint64       `json:"Percent,omitempty" name:"Percent"`
	AccountUin          *string       `json:"AccountUin,omitempty" name:"AccountUin"`
	EndTime             *string       `json:"EndTime,omitempty" name:"EndTime"`
	DataObjectNum       *uint64       `json:"DataObjectNum,omitempty" name:"DataObjectNum"`
	Tags                []*string     `json:"Tags,omitempty" name:"Tags"`
	InstanceId          *string       `json:"InstanceId,omitempty" name:"InstanceId"`
	IsPermanent         *bool         `json:"IsPermanent,omitempty" name:"IsPermanent"`
	BackupSize          *float64      `json:"BackupSize,omitempty" name:"BackupSize"`
	BinlogFileNum       *uint64       `json:"BinlogFileNum,omitempty" name:"BinlogFileNum"`
	DataFileNum         *uint64       `json:"DataFileNum,omitempty" name:"DataFileNum"`
	DeadlineTime        *string       `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	ArchiveStatus       *string       `json:"ArchiveStatus,omitempty" name:"ArchiveStatus"`
	AccountName         *string       `json:"AccountName,omitempty" name:"AccountName"`
	BucketDetail        *BucketDetail `json:"BucketDetail,omitempty" name:"BucketDetail"`
	AutoBackupPolicyId  *string       `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	ShardIds            []*string     `json:"ShardIds,omitempty" name:"ShardIds"`
	BackupState         *string       `json:"BackupState,omitempty" name:"BackupState"`
	ProtectionTimeScope *string       `json:"ProtectionTimeScope,omitempty" name:"ProtectionTimeScope"`
	AppId               *uint64       `json:"AppId,omitempty" name:"AppId"`
	BackupId            *string       `json:"BackupId,omitempty" name:"BackupId"`
	CreateTime          *string       `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeResourceBackupsRequest struct {
	*tchttp.BaseRequest
	ResourceType *string   `json:"ResourceType,omitempty" name:"ResourceType"`
	Filters      []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset       *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit        *uint64   `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourceBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		TotalCount *uint64           `json:"TotalCount,omitempty" name:"TotalCount"`
		BackupSet  []*ResourceBackup `json:"BackupSet,omitempty" name:"BackupSet"`
		RequestId  *string           `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceBackupsRequest struct {
	*tchttp.BaseRequest
	BackupIds             []*string `json:"BackupIds,omitempty" name:"BackupIds"`
	DeleteRetreatedBackup *bool     `json:"DeleteRetreatedBackup,omitempty" name:"DeleteRetreatedBackup"`
}

func (r *DeleteResourceBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteResourceBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Database Backup Models
type CreateDatabaseResourceBackupRequest struct {
	*tchttp.BaseRequest
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	BackupName   *string `json:"BackupName,omitempty" name:"BackupName"`
	ResourceId   *string `json:"ResourceId,omitempty" name:"ResourceId"`
	Deadline     *string `json:"Deadline,omitempty" name:"Deadline"`
}

func (r *CreateDatabaseResourceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDatabaseResourceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDatabaseResourceBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		BackupId  *string `json:"BackupId,omitempty" name:"BackupId"`
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDatabaseResourceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDatabaseResourceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Auto Backup Policy Models
type CreateAutoBackupPolicyRequest struct {
	*tchttp.BaseRequest
	Policy                  []*PolicyItem            `json:"Policy,omitempty" name:"Policy"`
	IsPermanent             *bool                    `json:"IsPermanent,omitempty" name:"IsPermanent"`
	AutoBackupPolicyName    *string                  `json:"AutoBackupPolicyName,omitempty" name:"AutoBackupPolicyName"`
	FullBackupInterval      *int64                   `json:"FullBackupInterval,omitempty" name:"FullBackupInterval"`
	RetentionAmount         *uint64                  `json:"RetentionAmount,omitempty" name:"RetentionAmount"`
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitempty" name:"AdvancedRetentionPolicy"`
	DryRun                  *bool                    `json:"DryRun,omitempty" name:"DryRun"`
	ResourceType            *string                  `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *CreateAutoBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoBackupPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAutoBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		AutoBackupPolicyId *string `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
		NextTriggerTime    *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`
		RequestId          *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAutoBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyItem struct {
	Hour         []*uint64 `json:"Hour,omitempty" name:"Hour"`
	IntervalDays *uint64   `json:"IntervalDays,omitempty" name:"IntervalDays"`
	DayOfWeek    []*uint64 `json:"DayOfWeek,omitempty" name:"DayOfWeek"`
	DayOfMonth   []*uint64 `json:"DayOfMonth,omitempty" name:"DayOfMonth"`
}

type AdvancedRetentionPolicy struct {
	Days   *uint64 `json:"Days,omitempty" name:"Days"`
	Weeks  *uint64 `json:"Weeks,omitempty" name:"Weeks"`
	Months *uint64 `json:"Months,omitempty" name:"Months"`
	Years  *uint64 `json:"Years,omitempty" name:"Years"`
}

type DescribeAutoBackupPolicy struct {
	*tchttp.BaseRequest
	Policy                  []*PolicyItem            `json:"Policy,omitempty" name:"Policy"`
	IsPermanent             *bool                    `json:"IsPermanent,omitempty" name:"IsPermanent"`
	AutoBackupPolicyName    *string                  `json:"AutoBackupPolicyName,omitempty" name:"AutoBackupPolicyName"`
	FullBackupInterval      *int64                   `json:"FullBackupInterval,omitempty" name:"FullBackupInterval"`
	RetentionAmount         *uint64                  `json:"RetentionAmount,omitempty" name:"RetentionAmount"`
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitempty" name:"AdvancedRetentionPolicy"`
	DryRun                  *bool                    `json:"DryRun,omitempty" name:"DryRun"`
	ResourceType            *string                  `json:"ResourceType,omitempty" name:"ResourceType"`
}

type AutoBackupPolicy struct {
	SubAccountUin           *string                  `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	DiskIdSet               []*string                `json:"DiskIdSet,omitempty" name:"DiskIdSet"`
	FileSystemIdSet         []*string                `json:"FileSystemIdSet,omitempty" name:"FileSystemIdSet"`
	ResourceType            *string                  `json:"ResourceType,omitempty" name:"ResourceType"`
	AutoBackupPolicyState   *string                  `json:"AutoBackupPolicyState,omitempty" name:"AutoBackupPolicyState"`
	AccountUin              *string                  `json:"AccountUin,omitempty" name:"AccountUin"`
	RetentionAmount         *uint64                  `json:"RetentionAmount,omitempty" name:"RetentionAmount"`
	Policy                  []*PolicyItem            `json:"Policy,omitempty" name:"Policy"`
	IsArchive               *bool                    `json:"IsArchive,omitempty" name:"IsArchive"`
	IsActivated             *bool                    `json:"IsActivated,omitempty" name:"IsActivated"`
	IsPermanent             *bool                    `json:"IsPermanent,omitempty" name:"IsPermanent"`
	BucketDetails           []*BucketDetail          `json:"BucketDetails,omitempty" name:"BucketDetails"`
	AccountName             *string                  `json:"AccountName,omitempty" name:"AccountName"`
	AutoBackupPolicyName    *string                  `json:"AutoBackupPolicyName,omitempty" name:"AutoBackupPolicyName"`
	InstanceIdSet           []*string                `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
	FullBackupInterval      *int64                   `json:"FullBackupInterval,omitempty" name:"FullBackupInterval"`
	AutoBackupPolicyId      *string                  `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitempty" name:"AdvancedRetentionPolicy"`
	ResourceIdSet           []*string                `json:"ResourceIdSet,omitempty" name:"ResourceIdSet"`
	NextTriggerTime         *string                  `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`
	CreateSpeed             *uint64                  `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	AppId                   *uint64                  `json:"AppId,omitempty" name:"AppId"`
	RetentionMonths         *uint64                  `json:"RetentionMonths,omitempty" name:"RetentionMonths"`
	CreateTime              *string                  `json:"CreateTime,omitempty" name:"CreateTime"`
	RetentionDays           *uint64                  `json:"RetentionDays,omitempty" name:"RetentionDays"`
}

type DescribeAutoBackupPoliciesRequest struct {
	*tchttp.BaseRequest
	Limit        *uint64   `json:"Limit,omitempty" name:"Limit"`
	Offset       *uint64   `json:"Offset,omitempty" name:"Offset"`
	Filters      []*Filter `json:"Filters,omitempty" name:"Filters"`
	ResourceType *string   `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeAutoBackupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoBackupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoBackupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		TotalCount          *uint64             `json:"TotalCount,omitempty" name:"TotalCount"`
		AutoBackupPolicySet []*AutoBackupPolicy `json:"AutoBackupPolicySet,omitempty" name:"AutoBackupPolicySet"`
		RequestId           *string             `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoBackupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoBackupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoBackupPolicyAttributeRequest struct {
	*tchttp.BaseRequest
	AutoBackupPolicyId      *string                  `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	Policy                  []*PolicyItem            `json:"Policy,omitempty" name:"Policy"`
	IsPermanent             *bool                    `json:"IsPermanent,omitempty" name:"IsPermanent"`
	AutoBackupPolicyName    *string                  `json:"AutoBackupPolicyName,omitempty" name:"AutoBackupPolicyName"`
	FullBackupInterval      *int64                   `json:"FullBackupInterval,omitempty" name:"FullBackupInterval"`
	RetentionAmount         *uint64                  `json:"RetentionAmount,omitempty" name:"RetentionAmount"`
	IsActivated             *bool                    `json:"IsActivated,omitempty" name:"IsActivated"`
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitempty" name:"AdvancedRetentionPolicy"`
	CreateSpeed             *uint64                  `json:"CreateSpeed,omitempty" name:"CreateSpeed"`
	NeedArchive             *bool                    `json:"NeedArchive,omitempty" name:"NeedArchive"`
	RetentionMonths         *uint64                  `json:"RetentionMonths,omitempty" name:"RetentionMonths"`
	RetentionDays           *uint64                  `json:"RetentionDays,omitempty" name:"RetentionDays"`
}

func (r *ModifyAutoBackupPolicyAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoBackupPolicyAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoBackupPolicyAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoBackupPolicyAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoBackupPolicyAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoBackupPoliciesRequest struct {
	*tchttp.BaseRequest
	AutoBackupPolicyIds []*string `json:"AutoBackupPolicyIds,omitempty" name:"AutoBackupPolicyIds"`
}

func (r *DeleteAutoBackupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoBackupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoBackupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAutoBackupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoBackupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// BindAutoBackupPolicy Models
type BindAutoBackupPolicyRequest struct {
	*tchttp.BaseRequest
	InstanceIds        []*string       `json:"InstanceIds,omitempty" name:"InstanceIds"`
	DiskIds            []*string       `json:"DiskIds,omitempty" name:"DiskIds"`
	FileSystemIds      []*string       `json:"FileSystemIds,omitempty" name:"FileSystemIds"`
	ResourceIds        []*string       `json:"ResourceIds,omitempty" name:"ResourceIds"`
	BucketDetails      []*BucketDetail `json:"BucketDetails,omitempty" name:"BucketDetails"`
	AutoBackupPolicyId *string         `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	ResourceType       *string         `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *BindAutoBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoBackupPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindAutoBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAutoBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// UnbindAutoBackupPolicyRequest UnbindAutoBackupPolicy Models
type UnbindAutoBackupPolicyRequest struct {
	*tchttp.BaseRequest
	InstanceIds        []*string       `json:"InstanceIds,omitempty" name:"InstanceIds"`
	DiskIds            []*string       `json:"DiskIds,omitempty" name:"DiskIds"`
	FileSystemIds      []*string       `json:"FileSystemIds,omitempty" name:"FileSystemIds"`
	ResourceIds        []*string       `json:"ResourceIds,omitempty" name:"ResourceIds"`
	BucketDetails      []*BucketDetail `json:"BucketDetails,omitempty" name:"BucketDetails"`
	AutoBackupPolicyId *string         `json:"AutoBackupPolicyId,omitempty" name:"AutoBackupPolicyId"`
	ResourceType       *string         `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *UnbindAutoBackupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoBackupPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindAutoBackupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindAutoBackupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoBackupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
