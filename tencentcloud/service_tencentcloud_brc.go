package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"time"

	brc "terraform-provider-tencentcloudenterprise/sdk/brc/v20220516"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type BRCService struct {
	client *connectivity.TencentCloudClient
}

// DescribeResourceBackupStatusByResourceType 查询资源备份概览
func (me *BRCService) DescribeResourceBackupStatusByResourceType(ctx context.Context, resourceType string) (
	detail *brc.ResourceOverviewDetail, err error) {
	logId := getLogId(ctx)
	request := brc.NewDescribeResourceBackupOverviewRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()
	request.ResourceType = helper.String(resourceType)
	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBrcClient().DescribeResourceBackupOverview(request)

	if err != nil {
		return nil, err
	}

	for _, resourceOverviewDetail := range response.Response.OverviewDetailSet {
		if resourceType == *resourceOverviewDetail.ResourceType {
			return resourceOverviewDetail, nil
		}
	}

	return nil, nil
}

// DescribeBackupCvmResourceById 查询备份资源详情
func (me *BRCService) DescribeBackupCvmResourceById(ctx context.Context, backupGroupId string) (backupGroup *brc.BackupGroup, err error) {
	logId := getLogId(ctx)
	request := brc.NewDescribeBackupCvmResourceRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()
	request.Filters = []*brc.Filter{
		{
			Name:   helper.String(BACKUP_GROUP_ID),
			Values: []*string{helper.String(backupGroupId)},
		},
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBrcClient().DescribeBackupCvmResource(request)

	if err != nil {
		return nil, err
	}

	for _, backupGroup := range response.Response.BackupGroupSet {
		if backupGroupId == *backupGroup.BackupGroupId {
			return backupGroup, nil
		}
	}

	return nil, nil
}

// DeleteBackupCvmResourceById 删除备份资源
func (me *BRCService) DeleteBackupCvmResourceById(ctx context.Context, backupGroupId string, onlyDismiss *bool) (err error) {
	logId := getLogId(ctx)
	request := brc.NewDeleteBackupCvmResourceRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()
	request.BackupGroupIds = []*string{helper.String(backupGroupId)}
	request.OnlyDismiss = onlyDismiss
	ratelimit.Check(request.GetAction())

	_, err = me.client.UseBrcClient().DeleteBackupCvmResource(request)

	return err
}

// DescribeBackupById 查询备份详情
func (me *BRCService) DescribeBackupById(ctx context.Context, backupId string) (backup *brc.Backup, err error) {
	logId := getLogId(ctx)
	request := brc.NewDescribeBackupsRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()

	// 创建过滤器来查询特定的备份ID
	filter := &brc.Filter{
		Name:   helper.String(BACKUP_ID),
		Values: []*string{helper.String(backupId)},
	}
	request.Filters = []*brc.Filter{filter}
	request.Limit = helper.Uint64(1)
	request.Offset = helper.Uint64(0)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBrcClient().DescribeBackups(request)

	if err != nil {
		return nil, err
	}

	if len(response.Response.BackupSet) > 0 {
		return response.Response.BackupSet[0], nil
	}

	return nil, nil
}

// DeleteBackupById 删除备份
func (me *BRCService) DeleteBackupById(ctx context.Context, backupId string) (err error) {
	logId := getLogId(ctx)
	request := brc.NewDeleteBackupsRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()
	request.BackupIds = []*string{helper.String(backupId)}
	ratelimit.Check(request.GetAction())

	_, err = me.client.UseBrcClient().DeleteBackups(request)

	return err
}

// DescribeCfsBackupById 查询CFS备份详情
func (me *BRCService) DescribeCfsBackupById(ctx context.Context, backupId string) (backup *brc.CfsBackup, err error) {
	logId := getLogId(ctx)
	request := brc.NewDescribeCfsBackupsRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()

	// 创建过滤器来查询特定的备份ID
	filter := &brc.Filter{
		Name:   helper.String("backup-id"),
		Values: []*string{helper.String(backupId)},
	}
	request.Filters = []*brc.Filter{filter}
	request.Limit = helper.Uint64(1)
	request.Offset = helper.Uint64(0)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBrcClient().DescribeCfsBackups(request)

	if err != nil {
		return nil, err
	}

	if len(response.Response.BackupSet) > 0 {
		return response.Response.BackupSet[0], nil
	}

	return nil, nil
}

// CreateResourceBackup
func (me *BRCService) CreateResourceBackup(ctx context.Context, resourceType, resourceId, backupName *string, deadline *time.Time,
	createSpeed *uint64, bucketDetail *brc.BucketDetail) (backupId string, err error) {
	logId := getLogId(ctx)
	request := brc.NewCreateResourceBackupRequest()
	request.ResourceType = resourceType
	request.ResourceId = resourceId
	request.BackupName = backupName
	request.Deadline = deadline
	request.CreateSpeed = createSpeed
	request.BucketDetail = bucketDetail

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBrcClient().CreateResourceBackup(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return "", err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil || response.Response.BackupId == nil {
		return "", fmt.Errorf("create resource backup failed: backup id is nil")
	}

	return *response.Response.BackupId, nil
}

// DescribeResourceBackups 查询COS/CSP资源备份
func (me *BRCService) DescribeResourceBackups(ctx context.Context, resourceType, backupId string) (backup *brc.ResourceBackup, err error) {
	logId := getLogId(ctx)
	request := brc.NewDescribeResourceBackupsRequest()
	request.ResourceType = &resourceType
	request.Limit = helper.Uint64(20)
	request.Offset = helper.Uint64(0)

	if backupId != "" {
		request.Filters = []*brc.Filter{
			{
				Name:   helper.String("backup-id"),
				Values: []*string{&backupId},
			},
		}
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBrcClient().DescribeResourceBackups(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return nil, err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil || response.Response.TotalCount == nil || *response.Response.TotalCount == 0 {
		return nil, fmt.Errorf("resource backup not found")
	}

	if len(response.Response.BackupSet) == 0 {
		return nil, fmt.Errorf("resource backup not found")
	}

	return response.Response.BackupSet[0], nil
}

// DeleteResourceBackups
func (me *BRCService) DeleteResourceBackups(ctx context.Context, backupIds []*string, deleteRetreatedBackup *bool) error {
	logId := getLogId(ctx)
	request := brc.NewDeleteResourceBackupsRequest()
	request.BackupIds = backupIds
	request.DeleteRetreatedBackup = deleteRetreatedBackup

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBrcClient().DeleteResourceBackups(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

// CreateDatabaseResourceBackup 创建数据库资源备份
func (me *BRCService) CreateDatabaseResourceBackup(ctx context.Context, resourceType, backupName, resourceId, deadline string) (backupId string, err error) {
	logId := getLogId(ctx)
	request := brc.NewCreateDatabaseResourceBackupRequest()
	request.ResourceType = &resourceType
	request.BackupName = &backupName
	request.ResourceId = &resourceId
	request.Deadline = &deadline

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBrcClient().CreateDatabaseResourceBackup(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return "", err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil || response.Response.BackupId == nil {
		return "", fmt.Errorf("create database resource backup failed: backup id is nil")
	}

	return *response.Response.BackupId, nil
}

// CreateAutoBackupPolicy 创建自动备份策略
func (me *BRCService) CreateAutoBackupPolicy(ctx context.Context, policy *brc.PolicyItem, isPermanent bool,
	autoBackupPolicyName string, fullBackupInterval int64, retentionAmount uint64,
	advancedRetentionPolicy *brc.AdvancedRetentionPolicy, dryRun bool, resourceType string) (autoBackupPolicyId string, err error) {
	logId := getLogId(ctx)
	request := brc.NewCreateAutoBackupPolicyRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()

	request.Policy = []*brc.PolicyItem{policy}
	request.IsPermanent = helper.Bool(isPermanent)
	request.AutoBackupPolicyName = helper.String(autoBackupPolicyName)
	request.FullBackupInterval = helper.Int64(fullBackupInterval)
	// 当isPermanent为true时，不设置RetentionAmount参数
	if !isPermanent {
		request.RetentionAmount = helper.Uint64(retentionAmount)
	}
	// 只有当advancedRetentionPolicy不为nil时才设置
	if advancedRetentionPolicy != nil {
		request.AdvancedRetentionPolicy = advancedRetentionPolicy
	}
	request.DryRun = helper.Bool(dryRun)
	request.ResourceType = helper.String(resourceType)

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBrcClient().CreateAutoBackupPolicy(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return "", err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	// 检查响应
	if response.Response == nil {
		return "", fmt.Errorf("create auto backup policy failed: response is nil")
	}

	// 如果是DryRun模式，AutoBackupPolicyId可能为空，这是正常的
	if dryRun {
		if response.Response.AutoBackupPolicyId == nil || *response.Response.AutoBackupPolicyId == "" {
			// 在DryRun模式下，我们可以使用一个临时的ID或者返回错误
			return "", fmt.Errorf("dry run mode: auto backup policy ID is empty, this is expected for dry run")
		}
	} else {
		// 非DryRun模式下，AutoBackupPolicyId必须存在
		if response.Response.AutoBackupPolicyId == nil || *response.Response.AutoBackupPolicyId == "" {
			return "", fmt.Errorf("create auto backup policy failed: auto backup policy ID is empty")
		}
	}

	return *response.Response.AutoBackupPolicyId, nil
}

// DescribeAutoBackupPolicyById 根据ID查询自动备份策略
func (me *BRCService) DescribeAutoBackupPolicyById(ctx context.Context, autoBackupPolicyId *string) (
	policy *brc.AutoBackupPolicy, err error) {
	logId := getLogId(ctx)
	request := brc.NewDescribeAutoBackupPoliciesRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()

	request.Filters = []*brc.Filter{
		{
			Name:   helper.String("auto-backup-policy-id"),
			Values: []*string{autoBackupPolicyId},
		},
	}
	request.Limit = helper.Uint64(20)
	request.Offset = helper.Uint64(0)

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBrcClient().DescribeAutoBackupPolicies(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return nil, err
	}

	if len(response.Response.AutoBackupPolicySet) > 0 {
		return response.Response.AutoBackupPolicySet[0], nil
	}

	return nil, nil
}

// ModifyAutoBackupPolicyAttribute 修改自动备份策略属性
func (me *BRCService) ModifyAutoBackupPolicyAttribute(ctx context.Context, request *brc.ModifyAutoBackupPolicyAttributeRequest) (err error) {
	logId := getLogId(ctx)

	// 使用NewModifyAutoBackupPolicyAttributeRequest()来正确初始化请求
	apiRequest := brc.NewModifyAutoBackupPolicyAttributeRequest()

	// 复制字段到正确初始化的请求中
	if request.AutoBackupPolicyId != nil {
		apiRequest.AutoBackupPolicyId = request.AutoBackupPolicyId
	}
	if request.Policy != nil {
		apiRequest.Policy = request.Policy
	}
	if request.IsPermanent != nil {
		apiRequest.IsPermanent = request.IsPermanent
	}
	if request.AutoBackupPolicyName != nil {
		apiRequest.AutoBackupPolicyName = request.AutoBackupPolicyName
	}
	if request.FullBackupInterval != nil {
		apiRequest.FullBackupInterval = request.FullBackupInterval
	}
	if request.RetentionAmount != nil {
		apiRequest.RetentionAmount = request.RetentionAmount
	}
	if request.IsActivated != nil {
		apiRequest.IsActivated = request.IsActivated
	}

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, apiRequest.GetAction(), err.Error())
		}
	}()

	ratelimit.Check(apiRequest.GetAction())
	_, err = me.client.UseBrcClient().ModifyAutoBackupPolicyAttribute(apiRequest)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, apiRequest.GetAction(), apiRequest.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n",
		logId, apiRequest.GetAction(), apiRequest.ToJsonString())

	return nil
}

// DeleteAutoBackupPolicies 删除自动备份策略
func (me *BRCService) DeleteAutoBackupPolicies(ctx context.Context, autoBackupPolicyIds []string) (err error) {
	logId := getLogId(ctx)
	request := brc.NewDeleteAutoBackupPoliciesRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()

	request.AutoBackupPolicyIds = helper.Strings(autoBackupPolicyIds)

	ratelimit.Check(request.GetAction())
	_, err = me.client.UseBrcClient().DeleteAutoBackupPolicies(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}

	return nil
}

// BindAutoBackupPolicy 绑定自动备份策略到资源
func (me *BRCService) BindAutoBackupPolicy(ctx context.Context, instanceIds, diskIds, fileSystemIds, resourceIds []*string,
	bucketDetails []*brc.BucketDetail, autoBackupPolicyId, resourceType *string) (err error) {
	logId := getLogId(ctx)
	request := brc.NewBindAutoBackupPolicyRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()

	request.AutoBackupPolicyId = autoBackupPolicyId
	request.ResourceType = resourceType
	request.InstanceIds = instanceIds
	request.DiskIds = diskIds
	request.FileSystemIds = fileSystemIds
	request.ResourceIds = resourceIds
	request.BucketDetails = bucketDetails

	ratelimit.Check(request.GetAction())
	_, err = me.client.UseBrcClient().BindAutoBackupPolicy(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n",
		logId, request.GetAction(), request.ToJsonString())

	return nil
}

// UnbindAutoBackupPolicy 解绑自动备份策略
func (me *BRCService) UnbindAutoBackupPolicy(ctx context.Context, instanceIds, diskIds, fileSystemIds, resourceIds []*string,
	bucketDetails []*brc.BucketDetail, autoBackupPolicyId, resourceType *string) (err error) {
	logId := getLogId(ctx)
	request := brc.NewUnbindAutoBackupPolicyRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()

	request.AutoBackupPolicyId = autoBackupPolicyId
	request.ResourceType = resourceType
	request.AutoBackupPolicyId = autoBackupPolicyId
	request.ResourceType = resourceType
	request.InstanceIds = instanceIds
	request.DiskIds = diskIds
	request.FileSystemIds = fileSystemIds
	request.ResourceIds = resourceIds
	request.BucketDetails = bucketDetails

	ratelimit.Check(request.GetAction())
	_, err = me.client.UseBrcClient().UnbindAutoBackupPolicy(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n",
		logId, request.GetAction(), request.ToJsonString())

	return nil
}
