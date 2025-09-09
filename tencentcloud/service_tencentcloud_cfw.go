package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type CfwService struct {
	client *connectivity.TencentCloudClient
}

func (me *CfwService) DescribeNatFwInstancesInfoById(ctx context.Context, instanceId string) (
	natInsInfo *cfw.NatInstanceInfo, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeNatFwInstancesInfoRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	filter := cfw.NatFwFilter{
		FilterType:    helper.String("natinsid"),
		FilterContent: helper.String(instanceId),
	}

	request.Filter = []*cfw.NatFwFilter{&filter}
	request.Limit = helper.IntUint64(20)
	request.Offset = helper.IntUint64(0)
	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeNatFwInstancesInfo(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *CfwService) DescribeNatFwVpcDnsLstById(ctx context.Context, instanceId string) (
	vpcDnsList []*cfw.VpcDnsInfo, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeNatFwVpcDnsLstRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.NatFwInsId = helper.String(instanceId)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeNatFwVpcDnsLst(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	vpcDnsList = response.Response.VpcDnsSwitchLst

	return
}

func (me *CfwService) DescribeCfwEipsById(ctx context.Context, instanceId string) (
	eipList []*cfw.NatFwEipsInfo, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeCfwEipsRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.CfwInstance = helper.String(instanceId)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeCfwEips(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	eipList = response.Response.NatFwEipList

	return
}

func (me *CfwService) ModifyNatInstance(ctx context.Context, instanceId, instanceName string) (
	err error) {

	logId := getLogId(ctx)
	request := cfw.NewModifyNatFwInstanceRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.NatInstanceId = helper.String(instanceId)
	request.InstanceName = helper.String(instanceName)

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().ModifyNatFwInstance(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

func (me *CfwService) DeleteNatFwInstanceById(ctx context.Context, instanceId string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewDeleteNatFwInstanceRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.CfwInstance = helper.String(instanceId)

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().DeleteNatFwInstance(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

func (me *CfwService) DescribeVpcFwGroupInstanceById(ctx context.Context, fwGroupId string) (
	vpcFwInstance *cfw.VpcFwGroupInfo, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeVpcFwGroupInfoRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("FwGroupId"),
			Values:       helper.Strings([]string{fwGroupId}),
			OperatorType: helper.Int64(1),
		},
	}
	request.Limit = helper.IntUint64(20)
	request.Offset = helper.IntUint64(0)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeFwGroupInstanceInfo(request)
	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.VpcFwGroupLst) < 1 {
		return
	}

	vpcFwInstance = response.Response.VpcFwGroupLst[0]
	return
}

// DeleteCfwVpcInstanceById
func (me *CfwService) DeleteCfwVpcInstanceById(ctx context.Context, fwGroupId string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewDeleteVpcFwGroupRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.FwGroupId = helper.String(fwGroupId)

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().DeleteVpcFwGroup(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

// DescribeNatFwPolicyById
func (me *CfwService) DescribeNatFwPolicyById(ctx context.Context, uuid string) (
	natPolicy *cfw.DescAcItem, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeNatAcRuleRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("Id"),
			Values:       helper.Strings([]string{uuid}),
			OperatorType: helper.Int64(1),
		},
	}
	request.Limit = helper.IntUint64(20)
	request.Offset = helper.IntUint64(0)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeNatAcRule(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	natPolicy = response.Response.Data[0]

	return
}

// DeleteNatFwPolicyById ...
func (me *CfwService) DeleteNatFwPolicyById(ctx context.Context, uuid string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewRemoveNatAcRuleRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()
	uuidInt, _ := strconv.ParseInt(uuid, 10, 64)

	request.RuleUuid = []*int64{&uuidInt}

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().RemoveNatAcRule(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

// DescribeVpcFwPolicyById
func (me *CfwService) DescribeVpcFwPolicyById(ctx context.Context, uuid string) (
	vpcPolicy *cfw.VpcRuleItem, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeVpcAcRuleRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("Id"),
			Values:       helper.Strings([]string{uuid}),
			OperatorType: helper.Int64(1),
		},
	}
	request.Limit = helper.IntUint64(20)
	request.Offset = helper.IntUint64(0)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeVpcAcRule(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	vpcPolicy = response.Response.Data[0]

	return
}

// DeleteVpcFwPolicyById
func (me *CfwService) DeleteVpcFwPolicyById(ctx context.Context, uuid string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewRemoveVpcAcRuleRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	uuidInt, _ := strconv.ParseInt(uuid, 10, 64)

	request.RuleUuids = []*int64{&uuidInt}

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().RemoveVpcAcRule(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

// DescribeBlockIgnoreListById
func (me *CfwService) DescribeBlockIgnoreListById(ctx context.Context, iP, domain, direction, ruleType string) (
	blockIgnoreList *cfw.BlockIgnoreRule, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeBlockIgnoreListRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	var searchStr string
	if iP != "" {
		searchStr = fmt.Sprintf(`{"domain":"%s"}`, iP)
	} else {
		searchStr = fmt.Sprintf(`{"domain":"%s"}`, domain)
	}

	request.Limit = helper.IntUint64(20)
	request.Offset = helper.IntUint64(0)
	request.SearchValue = &searchStr
	request.Direction = &direction
	ruleTypeInt, _ := strconv.ParseUint(ruleType, 10, 64)
	request.RuleType = &ruleTypeInt
	request.By = helper.String("EndTime")
	request.Order = helper.String("desc")

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeBlockIgnoreList(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	blockIgnoreList = response.Response.Data[0]

	return
}

// DeleteBlockIgnoreListById ...
func (me *CfwService) DeleteBlockIgnoreListById(ctx context.Context, iP, domain, direction, ruleType string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewDeleteBlockIgnoreRuleListRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	directionInt, _ := strconv.ParseInt(direction, 10, 64)
	if iP != "" {
		request.Rules = []*cfw.IocListData{
			{
				IP:        helper.String(iP),
				Direction: helper.Int64(directionInt),
			},
		}
	} else {
		request.Rules = []*cfw.IocListData{
			{
				Domain:    helper.String(domain),
				Direction: helper.Int64(directionInt),
			},
		}
	}

	ruleTypeInt, _ := strconv.ParseInt(ruleType, 10, 64)
	request.RuleType = helper.Int64(ruleTypeInt)

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().DeleteBlockIgnoreRuleList(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}
