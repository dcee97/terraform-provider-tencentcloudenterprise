/*
The TencentCloud provider is used to interact with many resources supported by [TencentCloud](https://intl.cloud.tencent.com).
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.

-> **Note:** From version 1.9.0 (June 18, 2019), the provider start to support Terraform 0.12.x.

Example Usage
```hcl

	terraform {
	  required_providers {
	    tencentcloud = {
	      source = "tencentcloudstack/tencentcloud"
	    }
	  }
	}

# Configure the TencentCloud Provider

	provider "tencentcloud" {
	  secret_id  = var.secret_id
	  secret_key = var.secret_key
	  region     = var.region
	}

#Configure the TencentCloud Provider with STS

	provider "tencentcloud" {
	  secret_id  = var.secret_id
	  secret_key = var.secret_key
	  region     = var.region
	  assume_role {
	    role_arn         = var.assume_role_arn
	    session_name     = var.session_name
	    session_duration = var.session_duration
	    policy           = var.policy
	  }
	}

```

# Resources List

Provider Data Sources

	cloud_availability_regions
	cloud_availability_zones

Auto Scaling(AS)

	  Data Source
	    cloud_as_scaling_configs
	    cloud_as_scaling_groups
	    cloud_as_scaling_policies
		cloud_as_instances
		cloud_as_last_activity

	  Resource
	    cloud_as_scaling_config
	    cloud_as_scaling_group
	    cloud_as_attachment
	    cloud_as_scaling_policy
	    cloud_as_schedule
	    cloud_as_lifecycle_hook
	    cloud_as_notification
		cloud_as_remove_instances
	    cloud_as_protect_instances

Bare Metal Server(BMS)

	  Data Source
		cloud_bms_instances
		cloud_bms_placement_groups
		cloud_bms_flavors

	  Resource
		cloud_bms_instance
		cloud_bms_placement_group

Cloud Kafka(ckafka)

	  Data Source
	    cloud_ckafka_users
	    cloud_ckafka_acls
	    cloud_ckafka_topics
	    cloud_ckafka_instances
		cloud_ckafka_connect_resource
		cloud_ckafka_region
		cloud_ckafka_datahub_topic
		cloud_ckafka_datahub_group_offsets
		cloud_ckafka_datahub_task
		cloud_ckafka_group
		cloud_ckafka_group_offsets
		cloud_ckafka_group_info
		cloud_ckafka_task_status
		cloud_ckafka_topic_flow_ranking
		cloud_ckafka_topic_produce_connection
		cloud_ckafka_topic_subscribe_group
		cloud_ckafka_topic_sync_replica
		cloud_ckafka_zone

	  Resource
		cloud_ckafka_instance
	    cloud_ckafka_user
	    cloud_ckafka_acl
	    cloud_ckafka_topic
		cloud_ckafka_datahub_topic
		cloud_ckafka_connect_resource
		cloud_ckafka_renew_instance
		cloud_ckafka_acl_rule
		cloud_ckafka_consumer_group
		cloud_ckafka_consumer_group_modify_offset
		cloud_ckafka_datahub_task

Cloud Block Storage(CBS)

	  Data Source
	    cloud_cbs_snapshots
	    cloud_cbs_storages
		cloud_cbs_storages_set
	    cloud_cbs_snapshot_policies

	  Resource
	    cloud_cbs_storage
		cloud_cbs_storage_set
	    cloud_cbs_storage_attachment
		cloud_cbs_storage_set_attachment
	    cloud_cbs_snapshot
	    cloud_cbs_snapshot_policy
	    cloud_cbs_snapshot_policy_attachment
		cloud_cbs_snapshot_share_permission

Cloud Load Balancer(CLB)

	  Data Source
	    cloud_clb_attachments
	    cloud_clb_instances
	    cloud_clb_listener_rules
	    cloud_clb_listeners
	    cloud_clb_redirections
		cloud_clb_instance_by_cert_id
		cloud_clb_instance_detail
		cloud_clb_resources
		cloud_clb_target_health
	    cloud_clb_certificates

	  Resource
	    cloud_clb_instance
	    cloud_clb_listener
	    cloud_clb_listener_rule
	    cloud_clb_attachment
	    cloud_clb_redirection
		cloud_clb_customized_config
		cloud_clb_security_group_attachment
		cloud_clb_certificates

Cloud Object Storage(COS)

	Data Source
	  cloud_cos_bucket_object
	  cloud_cos_buckets

	Resource
	  cloud_cos_bucket
	  cloud_cos_bucket_object
	  cloud_cos_bucket_policy

Cloud Object Storage(CSP)

	Data Source
	  cloud_csp_bucket_object
	  cloud_csp_buckets

	Resource
	  cloud_csp_bucket
	  cloud_csp_bucket_object
	  cloud_csp_bucket_policy

Cloud Virtual Machine(CVM)

	  Data Source
	    cloud_cvm_image
	    cloud_cvm_images
	    cloud_cvm_instance_types
	    cloud_cvm_instances
		cloud_cvm_instances_set
	    cloud_cvm_key_pairs
	    cloud_cvm_placement_groups
		cloud_cvm_instances_modification
		cloud_cvm_instance_vnc_url
		cloud_cvm_disaster_recover_group_quota
		cloud_cvm_image_quota
		cloud_cvm_image_share_permission

	  Resource
	    cloud_cvm_instance
		cloud_cvm_instance_set

	    cloud_cvm_key_pair
	    cloud_cvm_placement_group
	    cloud_cvm_image
		cloud_cvm_launch_template
		cloud_cvm_security_group_attachment
		cloud_cvm_reboot_instance
		cloud_cvm_renew_instance
		cloud_cvm_sync_image
		cloud_cvm_image_share_permission

Cloud Elastic IP(EIP)

	  Data Source
		cloud_eips
		cloud_eip_address_quota
	  Resource
	    cloud_eip
	    cloud_eip_association
		cloud_eip_address_transform
		cloud_eip_public_address_adjust
		cloud_eip_normal_address_return

Direct Connect(DC)

	Data Source


	Resource

Elasticsearch Service(ES)

	Data Source

	Resource

Key Management Service(KMS)

	Data Source
	  cloud_kms_keys

	Resource
	  cloud_kms_key
	  cloud_kms_external_key

Corporate Identity Center(CIC)

		  Data Source

		  Resource
	        cloud_cic_external_saml_identity_provider
	        cloud_cic_group
	        cloud_cic_role_assignment
	        cloud_cic_role_configuration
	        cloud_cic_role_configuration_permission_custom_policy_attachment
	        cloud_cic_role_configuration_permission_policy_attachment
	        cloud_cic_scim_credential
	        cloud_cic_scim_synchronization_status
	        cloud_cic_user
	        cloud_cic_user_group_attachment
	        cloud_cic_user_sync_provisioning

Tencent Kubernetes Engine(TKE)

	  Data Source
	    cloud_tke_kubernetes_clusters
	    cloud_tke_kubernetes_charts
	    cloud_tke_kubernetes_cluster_common_names
		cloud_tke_kubernetes_available_cluster_versions

	  Resource
	    cloud_tke_kubernetes_cluster
	    cloud_tke_kubernetes_scale_worker
	    cloud_tke_kubernetes_cluster_attachment
		cloud_tke_kubernetes_cluster_endpoint

TencentDB for PostgreSQL(PostgreSQL)

	  Data Source
		cloud_postgresql_instances
		cloud_postgresql_specinfos
		cloud_postgresql_xlogs
		cloud_postgresql_parameter_templates
		cloud_postgresql_readonly_groups
		cloud_postgresql_base_backups
		cloud_postgresql_log_backups
		cloud_postgresql_backup_download_urls
		cloud_postgresql_db_instance_classes
		cloud_postgresql_default_parameters
		cloud_postgresql_recovery_time
		cloud_postgresql_regions
		cloud_postgresql_db_instance_versions
		cloud_postgresql_zones

	  Resource
		cloud_postgresql_instance
		cloud_postgresql_readonly_instance
		cloud_postgresql_readonly_group
		cloud_postgresql_readonly_attachment
		cloud_postgresql_parameter_template
		cloud_postgresql_backup_plan_config
		cloud_postgresql_security_group_config
		cloud_postgresql_backup_download_restriction_config
		cloud_postgresql_restart_db_instance_operation
		cloud_postgresql_renew_db_instance_operation
		cloud_postgresql_isolate_db_instance_operation
		cloud_postgresql_disisolate_db_instance_operation
		cloud_postgresql_rebalance_readonly_group_operation
		cloud_postgresql_delete_log_backup_operation
		cloud_postgresql_modify_account_remark_operation
		cloud_postgresql_modify_switch_time_period_operation
		cloud_postgresql_base_backup

Virtual Private Cloud(VPC)

	  Data Source
	    cloud_vpc_security_groups
		cloud_vpc_address_templates
		cloud_vpc_address_template_groups
	    cloud_vpc_acls
		cloud_vpc_account_attributes
		cloud_vpc_classic_link_instances
		cloud_vpc_gateway_flow_qos
		cloud_vpc_cvm_instances
		cloud_vpc_net_detect_states
		cloud_vpc_net_detect_state_check
		cloud_vpc_private_ip_addresses
		cloud_vpc_resource_dashboard
		cloud_vpc_security_group_limits
		cloud_vpc_security_group_references
		cloud_vpc_template_limits
		cloud_vpc_limits
	    cloud_vpc_instances
	    cloud_vpc_route_tables
	    cloud_vpc_subnets
	    cloud_vpc_dnats
	    cloud_vpc_enis
	    cloud_vpc_ha_vip_eip_attachments
	    cloud_vpc_ha_vips
	    cloud_vpc_nat_gateways
		cloud_vpc_bandwidth_package_quota

	  Resource
	    cloud_vpc_eni
	    cloud_vpc_eni_attachment
		cloud_vpc_eni_sg_attachment
	    cloud_vpc
		cloud_vpc_acl
		cloud_vpc_acl_attachment
		cloud_vpc_dc_gateway
		cloud_vpc_net_detect
		cloud_vpc_ipv6_cidr_block
		cloud_vpc_ipv6_subnet_cidr_block
		cloud_vpc_ipv6_eni_address
	    cloud_vpc_subnet
	    cloud_vpc_security_group
	    cloud_vpc_security_group_rule
	    cloud_vpc_security_group_rule_set
	    cloud_vpc_security_group_lite_rule
		cloud_vpc_address_template
		cloud_vpc_address_template_group
	    cloud_vpc_route_table
	    cloud_vpc_route_table_entry
	    cloud_vpc_dnat
	    cloud_vpc_nat_gateway
	    cloud_vpc_ha_vip
	    cloud_vpc_ha_vip_eip_attachment
		cloud_vpc_bandwidth_package
		cloud_vpc_bandwidth_package_attachment
		cloud_vpc_ipv6_address_bandwidth

Virtual Private Cloud DNS(VPCDNS)

	  Data Source
		cloud_vpcdns_domains
	    cloud_vpcdns_records

	  Resource
		cloud_vpcdns_domain
	    cloud_vpcdns_record

TDSQL for MySQL(DCDB)

	  Data Source
		cloud_dcdb_instances
		cloud_dcdb_accounts
		cloud_dcdb_databases
		cloud_dcdb_parameters
		cloud_dcdb_shards
		cloud_dcdb_security_groups
		cloud_dcdb_database_objects
		cloud_dcdb_database_tables

	  Resource
		cloud_dcdb_account
		cloud_dcdb_instance
		cloud_dcdb_security_group_attachment
		cloud_dcdb_account_privileges
		cloud_dcdb_db_parameters
		cloud_dcdb_db_sync_mode_config
		cloud_dcdb_encrypt_attributes_config
		cloud_dcdb_instance_config
		cloud_dcdb_cancel_dcn_job_operation
		cloud_dcdb_activate_hour_instance_operation
		cloud_dcdb_isolate_hour_instance_operation
		cloud_dcdb_flush_binlog_operation
		cloud_dcdb_switch_db_instance_ha_operation

TDSQL PostgreSQL (Tbase)

	  Data Source
	    cloud_tbase_instances
	    cloud_tbase_pg_instances

	  Resource
		cloud_tbase_instance
		cloud_tbase_pg_instance
		cloud_tbase_pg_instance_vip

TencentDB for Redis(crs)

	  Data Source
		cloud_redis_zone_config
	    cloud_redis_instances
	    cloud_redis_backup
	    cloud_redis_backup_download_info
	    cloud_redis_param_records
	    cloud_redis_instance_shards
	    cloud_redis_instance_task_list
	    cloud_redis_instance_node_info

	  Resource
	   	cloud_redis_instance
		cloud_redis_backup_config
		cloud_redis_param
		cloud_redis_clear_instance_operation
		cloud_redis_startup_instance_operation
		cloud_redis_replica_readonly

TDMQ for Pulsar(tpulsar)

	  Data Source
		cloud_tdmq_pulsar_clusters
		cloud_tdmq_pulsar_environments

	  Resource
	    cloud_tdmq_pulsar_cluster
		cloud_tdmq_pulsar_route
		cloud_tdmq_pulsar_environment
		cloud_tdmq_pulsar_topic
		cloud_tdmq_pulsar_role
		cloud_tdmq_pulsar_environment_role_attachment

TDMQ for RabbitMQ(trabbit)

	  Data Source
		cloud_tdmq_rabbitmq_node_list
		cloud_tdmq_rabbitmq_vip_instance

	  Resource
		cloud_tdmq_rabbitmq_user
		cloud_tdmq_rabbitmq_vip_instance
		cloud_tdmq_rabbitmq_virtual_host

TDMQ for RocketMQ(trocket)

	  Data Source
		cloud_tdmq_rocketmq_cluster
		cloud_tdmq_rocketmq_namespace
		cloud_tdmq_rocketmq_topic
		cloud_tdmq_rocketmq_role
		cloud_tdmq_rocketmq_group
		cloud_tdmq_rocketmq_messages

	  Resource
		cloud_tdmq_rocketmq_cluster
		cloud_tdmq_rocketmq_namespace
		cloud_tdmq_rocketmq_role
		cloud_tdmq_rocketmq_topic
		cloud_tdmq_rocketmq_group
		cloud_tdmq_rocketmq_environment_role
		cloud_tdmq_send_rocketmq_message

Tencent Service Framework(TSF)

	Data Source
	  cloud_tsf_application
	  cloud_tsf_application_config
	  cloud_tsf_application_file_config
	  cloud_tsf_application_public_config
	  cloud_tsf_cluster
	  cloud_tsf_microservice
	  cloud_tsf_config_summary
	  cloud_tsf_delivery_config_by_group_id
	  cloud_tsf_delivery_configs
	  cloud_tsf_public_config_summary
	  cloud_tsf_api_group
	  cloud_tsf_application_attribute
	  cloud_tsf_business_log_configs
	  cloud_tsf_api_detail
	  cloud_tsf_microservice_api_version
	  cloud_tsf_repository
	  cloud_tsf_pod_instances
	  cloud_tsf_gateway_all_group_apis
	  cloud_tsf_group_gateways
	  cloud_tsf_usable_unit_namespaces
	  cloud_tsf_group_instances
	  cloud_tsf_container_group
	  cloud_tsf_groups
	  cloud_tsf_ms_api_list

	Resource
	  cloud_tsf_cluster
	  cloud_tsf_microservice
	  cloud_tsf_application_config
	  cloud_tsf_api_group
	  cloud_tsf_namespace
	  cloud_tsf_path_rewrite
	  cloud_tsf_unit_rule
	  cloud_tsf_task
	  cloud_tsf_config_template
	  cloud_tsf_api_rate_limit_rule
	  cloud_tsf_application_release_config
	  cloud_tsf_lane
	  cloud_tsf_lane_rule
	  cloud_tsf_group
	  cloud_tsf_application
	  cloud_tsf_application_public_config_release
	  cloud_tsf_application_public_config
	  cloud_tsf_application_file_config_release
	  cloud_tsf_instances_attachment
	  cloud_tsf_bind_api_group
	  cloud_tsf_application_file_config
	  cloud_tsf_enable_unit_rule
	  cloud_tsf_deploy_container_group
	  cloud_tsf_deploy_vm_group
	  cloud_tsf_release_api_group
	  cloud_tsf_operate_container_group
	  cloud_tsf_operate_group
	  cloud_tsf_unit_namespace

Cloud Workload Protection Platform (CWP)

	  Data Source
		cloud_tdmq_rabbitmq_node_list
		cloud_tdmq_rabbitmq_vip_instance

	  Resource
		cloud_tdmq_rabbitmq_user
		cloud_tdmq_rabbitmq_vip_instance
		cloud_tdmq_rabbitmq_virtual_host
*/
package tencentcloud

import (
	"context"
	"net/url"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"

	common2 "terraform-provider-tencentcloudenterprise/sdk/common"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

const (
	PROVIDER_SECRET_ID                    = "TENCENTCLOUD_SECRET_ID"
	PROVIDER_SECRET_KEY                   = "TENCENTCLOUD_SECRET_KEY"
	PROVIDER_SECURITY_TOKEN               = "TENCENTCLOUD_SECURITY_TOKEN"
	PROVIDER_REGION                       = "TENCENTCLOUD_REGION"
	PROVIDER_PROTOCOL                     = "TENCENTCLOUD_PROTOCOL"
	PROVIDER_DOMAIN                       = "TENCENTCLOUD_DOMAIN"
	PROVIDER_ASSUME_ROLE_ARN              = "TENCENTCLOUD_ASSUME_ROLE_ARN"
	PROVIDER_ASSUME_ROLE_SESSION_NAME     = "TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME"
	PROVIDER_ASSUME_ROLE_SESSION_DURATION = "TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION"
	PROVIDER_CSP_DOMAIN                   = "TENCENTCLOUD_CSP_DOMAIN"
	PROVIDER_COS_DOMAIN                   = "TENCENTCLOUD_COS_DOMAIN"
)

type TencentCloudClient struct {
	apiV3Conn *connectivity.TencentCloudClient
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_ID, nil),
				Description: "This is the TencentCloud access key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_ID` environment variable.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_KEY, nil),
				Description: "This is the TencentCloud secret key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_KEY` environment variable.",
				Sensitive:   true,
			},
			"security_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECURITY_TOKEN, nil),
				Description: "TencentCloud Security Token of temporary access credentials. It can be sourced from the `TENCENTCLOUD_SECURITY_TOKEN` environment variable. Notice: for supported products, please refer to: [temporary key supported products](https://intl.cloud.tencent.com/document/product/598/10588).",
				Sensitive:   true,
			},
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_REGION, nil),
				Description:  "This is the TencentCloud region. It must be provided, but it can also be sourced from the `TENCENTCLOUD_REGION` environment variables. The default input value is ap-guangzhou.",
				InputDefault: "ap-guangzhou",
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_PROTOCOL, "HTTP"),
				ValidateFunc: validateAllowedStringValue([]string{"HTTP", "HTTPS"}),
				Description:  "The protocol of the API request. Valid values: `HTTP` and `HTTPS`. Default is `HTTPS`.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_DOMAIN, nil),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"csp_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_CSP_DOMAIN, ""),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"cos_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_COS_DOMAIN, ""),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"assume_role": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "The `assume_role` block. If provided, terraform will attempt to assume this role using the supplied credentials.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_arn": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_ARN, nil),
							Description: "The ARN of the role to assume. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_ARN`.",
						},
						"session_name": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_SESSION_NAME, nil),
							Description: "The session name to use when making the AssumeRole call. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME`.",
						},
						"session_duration": {
							Type:     schema.TypeInt,
							Required: true,
							DefaultFunc: func() (interface{}, error) {
								if v := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); v != "" {
									return strconv.Atoi(v)
								}
								return 7200, nil
							},
							ValidateFunc: validateIntegerInRange(0, 43200),
							Description:  "The duration of the session when making the AssumeRole call. Its value ranges from 0 to 43200(seconds), and default is 7200 seconds. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION`.",
						},
						"policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A more restrictive policy when making the AssumeRole call. Its content must not contains `principal` elements. Notice: more syntax references, please refer to: [policies syntax logic](https://intl.cloud.tencent.com/document/product/598/10603).",
						},
					},
				},
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"cloud_as_instances":                              dataSourceTencentCloudAsInstances(),
			"cloud_as_last_activity":                          dataSourceTencentCloudAsLastActivity(),
			"cloud_as_scaling_configs":                        dataSourceTencentCloudAsScalingConfigs(),
			"cloud_as_scaling_groups":                         dataSourceTencentCloudAsScalingGroups(),
			"cloud_as_scaling_policies":                       dataSourceTencentCloudAsScalingPolicies(),
			"cloud_bms_instances":                             dataTencentCloudBmsInstances(),
			"cloud_bms_placement_groups":                      dataTencentCloudBmsPlacementGroup(),
			"cloud_bms_flavors":                               dataTencentCloudBmsFlavors(),
			"cloud_brc_autobackup_policies":                   dataSourceTencentCloudBrcAutoBackupPolicies(),
			"cloud_brc_backups":                               dataSourceTencentCloudBrcBackups(),
			"cloud_brc_resource_backup_overview":              dataSourceTencentCloudBrcResourceBackupOverview(),
			"cloud_cbs_snapshot_policies":                     dataSourceTencentCloudCbsSnapshotPolicies(),
			"cloud_cbs_snapshots":                             dataSourceTencentCloudCbsSnapshots(),
			"cloud_cbs_storages":                              dataSourceTencentCloudCbsStorages(),
			"cloud_cbs_storages_set":                          dataSourceTencentCloudCbsStoragesSet(),
			"cloud_cfs_access_groups":                         dataSourceTencentCloudCfsAccessGroups(),
			"cloud_cfs_access_rules":                          dataSourceTencentCloudCfsAccessRules(),
			"cloud_cfs_available_zone":                        dataSourceTencentCloudCfsAvailableZone(),
			"cloud_cfs_file_system_clients":                   dataSourceTencentCloudCfsFileSystemClients(),
			"cloud_cfs_file_systems":                          dataSourceTencentCloudCfsFileSystems(),
			"cloud_cfs_mount_targets":                         dataSourceTencentCloudCfsMountTargets(),
			"cloud_ckafka_acls":                               dataSourceTencentCloudCkafkaAcls(),
			"cloud_ckafka_group":                              dataSourceTencentCloudCkafkaGroup(),
			"cloud_ckafka_group_info":                         dataSourceTencentCloudCkafkaGroupInfo(),
			"cloud_ckafka_group_offsets":                      dataSourceTencentCloudCkafkaGroupOffsets(),
			"cloud_ckafka_instances":                          dataSourceTencentCloudCkafkaInstances(),
			"cloud_ckafka_task_status":                        dataSourceTencentCloudCkafkaTaskStatus(),
			"cloud_ckafka_topics":                             dataSourceTencentCloudCkafkaTopics(),
			"cloud_ckafka_users":                              dataSourceTencentCloudCkafkaUsers(),
			"cloud_ckafka_zone":                               dataSourceTencentCloudCkafkaZone(),
			"cloud_clb_attachments":                           dataSourceTencentCloudClbServerAttachments(),
			"cloud_clb_certificates":                          dataSourceTencentCloudClbCertificates(),
			"cloud_clb_instance_by_cert_id":                   dataSourceTencentCloudClbInstanceByCertId(),
			"cloud_clb_instance_detail":                       dataSourceTencentCloudClbInstanceDetail(),
			"cloud_clb_instances":                             dataSourceTencentCloudClbInstances(),
			"cloud_clb_listener_rules":                        dataSourceTencentCloudClbListenerRules(),
			"cloud_clb_listeners":                             dataSourceTencentCloudClbListeners(),
			"cloud_clb_redirections":                          dataSourceTencentCloudClbRedirections(),
			"cloud_clb_resources":                             dataSourceTencentCloudClbResources(),
			"cloud_clb_target_health":                         dataSourceTencentCloudClbTargetHealth(),
			"cloud_cls_machine_group_configs":                 dataSourceTencentCloudClsMachineGroupConfigs(),
			"cloud_cos_bucket_object":                         dataSourceTencentCloudCosBucketObject(),
			"cloud_cos_buckets":                               dataSourceTencentCloudCosBuckets(),
			"cloud_csp_bucket_object":                         dataSourceTencentCloudCspBucketObject(),
			"cloud_csp_buckets":                               dataSourceTencentCloudCspBuckets(),
			"cloud_cvm_disaster_recover_group_quota":          dataSourceTencentCloudCvmDisasterRecoverGroupQuota(),
			"cloud_cvm_image":                                 dataSourceTencentCloudImage(),
			"cloud_cvm_image_quota":                           dataSourceTencentCloudCvmImageQuota(),
			"cloud_cvm_image_share_permission":                dataSourceTencentCloudCvmImageSharePermission(),
			"cloud_cvm_images":                                dataSourceTencentCloudImages(),
			"cloud_cvm_instance_types":                        dataSourceInstanceTypes(),
			"cloud_cvm_instance_vnc_url":                      dataSourceTencentCloudCvmInstanceVncUrl(),
			"cloud_cvm_instances":                             dataSourceTencentCloudInstances(),
			"cloud_cvm_instances_modification":                dataSourceTencentCloudCvmInstancesModification(),
			"cloud_cvm_instances_set":                         dataSourceTencentCloudInstancesSet(),
			"cloud_cvm_key_pairs":                             dataSourceTencentCloudKeyPairs(),
			"cloud_cvm_placement_groups":                      dataSourceTencentCloudPlacementGroups(),
			"cloud_dc_access_points":                          dataSourceTencentCloudDcAccessPoints(),
			"cloud_dc_instances":                              dataSourceTencentCloudDcInstances(),
			"cloud_dcdb_accounts":                             dataSourceTencentCloudDcdbAccounts(),
			"cloud_dcdb_database_objects":                     dataSourceTencentCloudDcdbDatabaseObjects(),
			"cloud_dcdb_database_tables":                      dataSourceTencentCloudDcdbDatabaseTables(),
			"cloud_dcdb_databases":                            dataSourceTencentCloudDcdbDatabases(),
			"cloud_dcdb_instances":                            dataSourceTencentCloudDcdbInstances(),
			"cloud_dcdb_parameters":                           dataSourceTencentCloudDcdbParameters(),
			"cloud_dcdb_security_groups":                      dataSourceTencentCloudDcdbSecurityGroups(),
			"cloud_dcdb_shards":                               dataSourceTencentCloudDcdbShards(),
			"cloud_eip_address_quota":                         dataSourceTencentCloudEipAddressQuota(),
			"cloud_eips":                                      dataSourceTencentCloudEips(),
			"cloud_kms_keys":                                  dataSourceTencentCloudKmsKeys(),
			"cloud_redis_backup":                              dataSourceTencentCloudRedisBackup(),
			"cloud_redis_backup_download_info":                dataSourceTencentCloudRedisBackupDownloadInfo(),
			"cloud_redis_instance_node_info":                  dataSourceTencentCloudRedisInstanceNodeInfo(),
			"cloud_redis_instance_shards":                     dataSourceTencentCloudRedisInstanceShards(),
			"cloud_redis_instance_task_list":                  dataSourceTencentCloudRedisInstanceTaskList(),
			"cloud_redis_instances":                           dataSourceTencentRedisInstances(),
			"cloud_redis_param_records":                       dataSourceTencentCloudRedisRecordsParam(),
			"cloud_redis_zone_config":                         dataSourceTencentRedisZoneConfig(),
			"cloud_ssm_secret_versions":                       dataSourceTencentCloudSsmSecretVersions(),
			"cloud_ssm_secrets":                               dataSourceTencentCloudSsmSecrets(),
			"cloud_tbase_instances":                           dataSourceTencentCloudTbaseInstances(),
			"cloud_tbase_pg_instances":                        dataSourceTencentCloudTbasePGInstances(),
			"cloud_tdmq_environments":                         dataSourceTencentCloudTdmqEnvironments(),
			"cloud_tdmq_pulsar_environments":                  dataSourceTencentCloudTdmqPulsarEnvironments(),
			"cloud_tdmq_pulsar_clusters":                      dataSourceTencentCloudTdmqPulsarCluster(),
			"cloud_tdmq_rabbitmq_node_list":                   dataSourceTencentCloudTdmqRabbitmqNodeList(),
			"cloud_tdmq_rabbitmq_vip_instance":                dataSourceTencentCloudTdmqRabbitmqVipInstance(),
			"cloud_tdmq_rocketmq_cluster":                     dataSourceTencentCloudTdmqRocketmqCluster(),
			"cloud_tdmq_rocketmq_group":                       dataSourceTencentCloudTdmqRocketmqGroup(),
			"cloud_tdmq_rocketmq_messages":                    dataSourceTencentCloudTdmqRocketmqMessages(),
			"cloud_tdmq_rocketmq_namespace":                   dataSourceTencentCloudTdmqRocketmqNamespace(),
			"cloud_tdmq_rocketmq_role":                        dataSourceTencentCloudTdmqRocketmqRole(),
			"cloud_tdmq_rocketmq_topic":                       dataSourceTencentCloudTdmqRocketmqTopic(),
			"cloud_tke_kubernetes_available_cluster_versions": dataSourceTencentCloudKubernetesAvailableClusterVersions(),
			"cloud_tke_kubernetes_charts":                     dataSourceTencentCloudKubernetesCharts(),
			"cloud_tke_kubernetes_cluster_common_names":       datasourceTencentCloudKubernetesClusterCommonNames(),
			"cloud_tke_kubernetes_clusters":                   dataSourceTencentCloudKubernetesClusters(),
			"cloud_tsf_api_detail":                            dataSourceTencentCloudTsfApiDetail(),
			"cloud_tsf_api_group":                             dataSourceTencentCloudTsfApiGroup(),
			"cloud_tsf_application":                           dataSourceTencentCloudTsfApplication(),
			"cloud_tsf_application_attribute":                 dataSourceTencentCloudTsfApplicationAttribute(),
			"cloud_tsf_application_config":                    dataSourceTencentCloudTsfApplicationConfig(),
			"cloud_tsf_application_file_config":               dataSourceTencentCloudTsfApplicationFileConfig(),
			"cloud_tsf_application_public_config":             dataSourceTencentCloudTsfApplicationPublicConfig(),
			"cloud_tsf_business_log_configs":                  dataSourceTencentCloudTsfBusinessLogConfigs(),
			"cloud_tsf_cluster":                               dataSourceTencentCloudTsfCluster(),
			"cloud_tsf_config_summary":                        dataSourceTencentCloudTsfConfigSummary(),
			"cloud_tsf_delivery_config_by_group_id":           dataSourceTencentCloudTsfDeliveryConfigByGroupId(),
			"cloud_tsf_delivery_configs":                      dataSourceTencentCloudTsfDeliveryConfigs(),
			"cloud_tsf_gateway_all_group_apis":                dataSourceTencentCloudTsfGatewayAllGroupApis(),
			"cloud_tsf_group_gateways":                        dataSourceTencentCloudTsfGroupGateways(),
			"cloud_tsf_group_instances":                       dataSourceTencentCloudTsfGroupInstances(),
			"cloud_tsf_microservice":                          dataSourceTencentCloudTsfMicroservice(),
			"cloud_tsf_microservice_api_version":              dataSourceTencentCloudTsfMicroserviceApiVersion(),
			"cloud_tsf_pod_instances":                         dataSourceTencentCloudTsfPodInstances(),
			"cloud_tsf_public_config_summary":                 dataSourceTencentCloudTsfPublicConfigSummary(),
			"cloud_tsf_repository":                            dataSourceTencentCloudTsfRepository(),
			"cloud_tsf_usable_unit_namespaces":                dataSourceTencentCloudTsfUsableUnitNamespaces(),
			"cloud_turbofs_p_groups":                          dataSourceTencentCloudTurbofsPGroups(),
			"cloud_turbofs_rules":                             dataSourceTencentCloudTurbofsRules(),
			"cloud_turbofs_file_systems":                      dataSourceTencentCloudTurbofsFileSystems(),
			"cloud_turbofs_mount_targets":                     dataSourceTencentCloudTurbofsMountTargets(),
			"cloud_vpc_account_attributes":                    dataSourceTencentCloudVpcAccountAttributes(),
			"cloud_vpc_acls":                                  dataSourceTencentCloudVpcAcls(),
			"cloud_vpc_address_template_groups":               dataSourceTencentCloudAddressTemplateGroups(),
			"cloud_vpc_address_templates":                     dataSourceTencentCloudAddressTemplates(),
			"cloud_vpc_classic_link_instances":                dataSourceTencentCloudVpcClassicLinkInstances(),
			"cloud_vpc_cvm_instances":                         dataSourceTencentCloudVpcCvmInstances(),
			"cloud_vpc_dnats":                                 dataSourceTencentCloudDnats(),
			"cloud_vpc_enis":                                  dataSourceTencentCloudEnis(),
			"cloud_vpc_gateway_flow_qos":                      dataSourceTencentCloudVpcGatewayFlowQos(),
			"cloud_vpc_ha_vip_eip_attachments":                dataSourceTencentCloudHaVipEipAttachments(),
			"cloud_vpc_ha_vips":                               dataSourceTencentCloudHaVips(),
			"cloud_vpc_instances":                             dataSourceTencentCloudVpcInstances(),
			"cloud_vpc_limits":                                dataSourceTencentCloudVpcLimits(),
			"cloud_vpc_nat_gateways":                          dataSourceTencentCloudNatGateways(),
			"cloud_vpc_net_detect_state_check":                dataSourceTencentCloudVpcNetDetectStateCheck(),
			"cloud_vpc_net_detect_states":                     dataSourceTencentCloudVpcNetDetectStates(),
			"cloud_vpc_private_ip_addresses":                  dataSourceTencentCloudVpcPrivateIpAddresses(),
			"cloud_vpc_resource_dashboard":                    dataSourceTencentCloudVpcResourceDashboard(),
			"cloud_vpc_route_tables":                          dataSourceTencentCloudVpcRouteTables(),
			"cloud_vpc_security_group_limits":                 dataSourceTencentCloudVpcSecurityGroupLimits(),
			"cloud_vpc_security_group_references":             dataSourceTencentCloudVpcSecurityGroupReferences(),
			"cloud_vpc_security_groups":                       dataSourceTencentCloudSecurityGroups(),
			"cloud_vpc_subnets":                               dataSourceTencentCloudVpcSubnets(),
			"cloud_vpc_template_limits":                       dataSourceTencentCloudVpcTemplateLimits(),
			"cloud_vpcdns_domains":                            dataSourceTencentCloudVpcDnsDomain(),
			"cloud_vpcdns_records":                            dataSourceTencentCloudVpcDnsRecord(),
			"cloud_vpcdns_forward_rules":                      dataSourceTencentCloudVpcDnsForwardRules(),
			"cloud_cwp_machines_simple":                       dataSourceTencentCloudCwpMachinesSimple(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"cloud_as_attachment":                           resourceTencentCloudAsAttachment(),
			"cloud_as_lifecycle_hook":                       resourceTencentCloudAsLifecycleHook(),
			"cloud_as_notification":                         resourceTencentCloudAsNotification(),
			"cloud_as_protect_instances":                    resourceTencentCloudAsProtectInstances(),
			"cloud_as_remove_instances":                     resourceTencentCloudAsRemoveInstances(),
			"cloud_as_scaling_config":                       resourceTencentCloudAsScalingConfig(),
			"cloud_as_scaling_group":                        resourceTencentCloudAsScalingGroup(),
			"cloud_as_scaling_policy":                       resourceTencentCloudAsScalingPolicy(),
			"cloud_as_schedule":                             resourceTencentCloudAsSchedule(),
			"cloud_bms_instance":                            resourceTencentCloudBmsInstance(),
			"cloud_bms_placement_group":                     resourceTencentCloudBmsPlacementGroup(),
			"cloud_brc_activate_backup_service":             resourceTencentCloudBrcActivateBackupService(),
			"cloud_brc_auto_backup_policy":                  resourceTencentCloudBrcAutoBackupPolicy(),
			"cloud_brc_auto_backup_policy_binding":          resourceTencentCloudBrcAutoBackupPolicyBinding(),
			"cloud_brc_backup_group":                        resourceTencentCloudBrcBackupGroup(),
			"cloud_brc_backup_disk":                         resourceTencentCloudBrcBackupDisk(),
			"cloud_brc_backup_cfs":                          resourceTencentCloudBrcBackupCfs(),
			"cloud_brc_backup_resource":                     resourceTencentCloudBrcBackupResource(),
			"cloud_cbs_snapshot":                            resourceTencentCloudCbsSnapshot(),
			"cloud_cbs_snapshot_policy":                     resourceTencentCloudCbsSnapshotPolicy(),
			"cloud_cbs_snapshot_policy_attachment":          resourceTencentCloudCbsSnapshotPolicyAttachment(),
			"cloud_cbs_snapshot_share_permission":           resourceTencentCloudCbsSnapshotSharePermission(),
			"cloud_cbs_storage":                             resourceTencentCloudCbsStorage(),
			"cloud_cbs_storage_attachment":                  resourceTencentCloudCbsStorageAttachment(),
			"cloud_cbs_storage_set":                         resourceTencentCloudCbsStorageSet(),
			"cloud_cfs_access_group":                        resourceTencentCloudCfsAccessGroup(),
			"cloud_cfs_access_rule":                         resourceTencentCloudCfsAccessRule(),
			"cloud_cfs_auto_snapshot_policy":                resourceTencentCloudCfsAutoSnapshotPolicy(),
			"cloud_cfs_auto_snapshot_policy_attachment":     resourceTencentCloudCfsAutoSnapshotPolicyAttachment(),
			"cloud_cfs_file_system":                         resourceTencentCloudCfsFileSystem(),
			"cloud_cfs_sign_up_cfs_service":                 resourceTencentCloudCfsSignUpCfsService(),
			"cloud_cfs_snapshot":                            resourceTencentCloudCfsSnapshot(),
			"cloud_cfw_nat_instance":                        resourceTencentCloudCfwNatInstance(),
			"cloud_cfw_vpc_instance":                        resourceTencentCloudCfwVpcInstance(),
			"cloud_cfw_nat_policy":                          resourceTencentCloudCfwNatPolicy(),
			"cloud_cfw_vpc_policy":                          resourceTencentCloudCfwVpcPolicy(),
			"cloud_cfw_block_ignore":                        resourceTencentCloudCfwBlockIgnore(),
			"cloud_ckafka_acl":                              resourceTencentCloudCkafkaAcl(),
			"cloud_ckafka_instance":                         resourceTencentCloudCkafkaInstance(),
			"cloud_ckafka_topic":                            resourceTencentCloudCkafkaTopic(),
			"cloud_ckafka_user":                             resourceTencentCloudCkafkaUser(),
			"cloud_clb_attachment":                          resourceTencentCloudClbServerAttachment(),
			"cloud_clb_certificates":                        resourceTencentCloudClbCertificate(),
			"cloud_clb_customized_config":                   resourceTencentCloudClbCustomizedConfig(),
			"cloud_clb_instance":                            resourceTencentCloudClbInstance(),
			"cloud_clb_listener":                            resourceTencentCloudClbListener(),
			"cloud_clb_listener_rule":                       resourceTencentCloudClbListenerRule(),
			"cloud_clb_redirection":                         resourceTencentCloudClbRedirection(),
			"cloud_clb_replace_cert":                        resourceTencentCloudClbReplaceCert(),
			"cloud_clb_security_group_attachment":           resourceTencentCloudClbSecurityGroupAttachment(),
			"cloud_cls_ckafka_consumer":                     resourceTencentCloudClsCkafkaConsumer(),
			"cloud_cls_config":                              resourceTencentCloudClsConfig(),
			"cloud_cls_config_attachment":                   resourceTencentCloudClsConfigAttachment(),
			"cloud_cls_cos_recharge":                        resourceTencentCloudClsCosRecharge(),
			"cloud_cls_cos_shipper":                         resourceTencentCloudClsCosShipper(),
			"cloud_cls_export":                              resourceTencentCloudClsExport(),
			"cloud_cls_index":                               resourceTencentCloudClsIndex(),
			"cloud_cls_logset":                              resourceTencentCloudClsLogset(),
			"cloud_cls_machine_group":                       resourceTencentCloudClsMachineGroup(),
			"cloud_cls_topic":                               resourceTencentCloudClsTopic(),
			"cloud_cos_bucket":                              resourceTencentCloudCosBucket(),
			"cloud_cos_bucket_object":                       resourceTencentCloudCosBucketObject(),
			"cloud_cos_bucket_policy":                       resourceTencentCloudCosBucketPolicy(),
			"cloud_csp_bucket":                              resourceTencentCloudCspBucket(),
			"cloud_csp_bucket_object":                       resourceTencentCloudCspBucketObject(),
			"cloud_csp_bucket_policy":                       resourceTencentCloudCspBucketPolicy(),
			"cloud_cvm_image":                               resourceTencentCloudImage(),
			"cloud_cvm_image_share_permission":              resourceTencentCloudCvmImageSharePermission(),
			"cloud_cvm_instance":                            resourceTencentCloudInstance(),
			"cloud_cvm_instance_set":                        resourceTencentCloudInstanceSet(),
			"cloud_cvm_key_pair":                            resourceTencentCloudKeyPair(),
			"cloud_cvm_launch_template":                     resourceTencentCloudCvmLaunchTemplate(),
			"cloud_cvm_placement_group":                     resourceTencentCloudPlacementGroup(),
			"cloud_cvm_reboot_instance":                     resourceTencentCloudCvmRebootInstance(),
			"cloud_cvm_renew_instance":                      resourceTencentCloudCvmRenewInstance(),
			"cloud_cvm_security_group_attachment":           resourceTencentCloudCvmSecurityGroupAttachment(),
			"cloud_cvm_sync_image":                          resourceTencentCloudCvmSyncImage(),
			"cloud_dc_dcx":                                  resourceTencentCloudDcxInstance(),
			"cloud_dc_instance":                             resourceTencentCloudDcInstance(),
			"cloud_dcdb_account":                            resourceTencentCloudDcdbAccount(),
			"cloud_dcdb_account_privileges":                 resourceTencentCloudDcdbAccountPrivileges(),
			"cloud_dcdb_activate_hour_instance_operation":   resourceTencentCloudDcdbActivateHourInstanceOperation(),
			"cloud_dcdb_cancel_dcn_job_operation":           resourceTencentCloudDcdbCancelDcnJobOperation(),
			"cloud_dcdb_db_parameters":                      resourceTencentCloudDcdbDbParameters(),
			"cloud_dcdb_db_sync_mode_config":                resourceTencentCloudDcdbDbSyncModeConfig(),
			"cloud_dcdb_encrypt_attributes_config":          resourceTencentCloudDcdbEncryptAttributesConfig(),
			"cloud_dcdb_flush_binlog_operation":             resourceTencentCloudDcdbFlushBinlogOperation(),
			"cloud_dcdb_instance":                           resourceTencentCloudDcdbdbInstance(),
			"cloud_dcdb_instance_config":                    resourceTencentCloudDcdbInstanceConfig(),
			"cloud_dcdb_isolate_hour_instance_operation":    resourceTencentCloudDcdbIsolateHourInstanceOperation(),
			"cloud_dcdb_security_group_attachment":          resourceTencentCloudDcdbSecurityGroupAttachment(),
			"cloud_dcdb_switch_db_instance_ha_operation":    resourceTencentCloudDcdbSwitchDbInstanceHaOperation(),
			"cloud_eip":                                     resourceTencentCloudEip(),
			"cloud_eip_address_transform":                   resourceTencentCloudEipAddressTransform(),
			"cloud_eip_association":                         resourceTencentCloudEipAssociation(),
			"cloud_eip_normal_address_return":               resourceTencentCloudEipNormalAddressReturn(),
			"cloud_kms_external_key":                        resourceTencentCloudKmsExternalKey(),
			"cloud_kms_key":                                 resourceTencentCloudKmsKey(),
			"cloud_redis_backup_config":                     resourceTencentCloudRedisBackupConfig(),
			"cloud_redis_clear_instance_operation":          resourceTencentCloudRedisClearInstanceOperation(),
			"cloud_redis_instance":                          resourceTencentCloudRedisInstance(),
			"cloud_redis_param":                             resourceTencentCloudRedisParam(),
			"cloud_redis_replica_readonly":                  resourceTencentCloudRedisReplicaReadonly(),
			"cloud_redis_startup_instance_operation":        resourceTencentCloudRedisStartupInstanceOperation(),
			"cloud_ssm_secret":                              resourceTencentCloudSsmSecret(),
			"cloud_ssm_secret_version":                      resourceTencentCloudSsmSecretVersion(),
			"cloud_tbase_instance":                          resourceTencentCloudTbaseInstance(),
			"cloud_tbase_pg_instance":                       resourceTencentCloudTbasePGInstance(),
			"cloud_tbase_pg_instance_vip":                   resourceTencentCloudTbasePGInstanceVip(),
			"cloud_tdmq_instance":                           resourceTencentCloudTdmqInstance(),
			"cloud_tdmq_pulsar_cluster":                     resourceTencentCloudTdmqPulsarCluster(),
			"cloud_tdmq_pulsar_route":                       resourceTencentCloudTdmqPulsarRoute(),
			"cloud_tdmq_pulsar_environment":                 resourceTencentCloudTdmqPulsarEnvironment(),
			"cloud_tdmq_pulsar_topic":                       resourceTencentCloudTdmqPulsarTopic(),
			"cloud_tdmq_pulsar_role":                        resourceTencentCloudTdmqPulsarRole(),
			"cloud_tdmq_pulsar_environment_role_attachment": resourceTencentCloudTdmqPulsarEnvironmentRoleAttachment(),
			"cloud_tdmq_namespace":                          resourceTencentCloudTdmqNamespace(),
			"cloud_tdmq_namespace_role_attachment":          resourceTencentCloudTdmqNamespaceRoleAttachment(),
			"cloud_tdmq_rabbitmq_user":                      resourceTencentCloudTdmqRabbitmqUser(),
			"cloud_tdmq_rabbitmq_vip_instance":              resourceTencentCloudTdmqRabbitmqVipInstance(),
			"cloud_tdmq_rabbitmq_virtual_host":              resourceTencentCloudTdmqRabbitmqVirtualHost(),
			"cloud_tdmq_rocketmq_cluster":                   resourceTencentCloudTdmqRocketmqCluster(),
			"cloud_tdmq_rocketmq_environment_role":          resourceTencentCloudTdmqRocketmqEnvironmentRole(),
			"cloud_tdmq_rocketmq_group":                     resourceTencentCloudTdmqRocketmqGroup(),
			"cloud_tdmq_rocketmq_namespace":                 resourceTencentCloudTdmqRocketmqNamespace(),
			"cloud_tdmq_rocketmq_role":                      resourceTencentCloudTdmqRocketmqRole(),
			"cloud_tdmq_rocketmq_topic":                     resourceTencentCloudTdmqRocketmqTopic(),
			"cloud_tdmq_role":                               resourceTencentCloudTdmqRole(),
			"cloud_tdmq_route":                              resourceTencentCloudTdmqRoute(),
			"cloud_tdmq_send_rocketmq_message":              resourceTencentCloudTdmqSendRocketmqMessage(),
			"cloud_tdmq_subscription_attachment":            resourceTencentCloudTdmqSubscriptionAttachment(),
			"cloud_tdmq_topic":                              resourceTencentCloudTdmqTopic(),
			"cloud_tke_kubernetes_cluster":                  resourceTencentCloudTkeCluster(),
			"cloud_tke_kubernetes_cluster_namespace":        resourceTencentCloudTkeClusterNamespace(),
			"cloud_tke_kubernetes_cluster_plugin":           resourceTencentCloudTkeClusterPlugin(),
			"cloud_tke_kubernetes_cluster_secret":           resourceTencentCloudTkeClusterSecret(),
			"cloud_tke_kubernetes_cluster_pv":               resourceTencentCloudTkeClusterPv(),
			"cloud_tke_kubernetes_cluster_pvc":              resourceTencentCloudTkeClusterPvc(),
			"cloud_tke_kubernetes_cluster_deploy":           resourceTencentCloudTkeClusterDeploy(),
			"cloud_tke_kubernetes_cluster_affinity":         resourceTencentCloudTkeClusterAffinity(),
			"cloud_tke_kubernetes_cluster_ing":              resourceTencentCloudTkeClusterIng(),
			"cloud_tke_kubernetes_cluster_attachment":       resourceTencentCloudTkeClusterAttachment(),
			"cloud_tke_kubernetes_cluster_endpoint":         resourceTencentCloudTkeClusterEndpoint(),
			"cloud_tke_kubernetes_scale_worker":             resourceTencentCloudTkeScaleWorker(),
			"cloud_tsf_cluster":                             resourceTencentCloudTsfCluster(),
			"cloud_tsf_namespace":                           resourceTencentCloudTsfNamespace(),
			"cloud_tsf_group":                               resourceTencentCloudTsfGroup(),
			"cloud_tsf_application":                         resourceTencentCloudTsfApplication(),
			"cloud_tsf_application_config":                  resourceTencentCloudTsfApplicationConfig(),
			"cloud_tsf_application_release_config":          resourceTencentCloudTsfApplicationReleaseConfig(),
			"cloud_tsf_application_file_config":             resourceTencentCloudTsfApplicationFileConfig(),
			"cloud_tsf_application_file_config_release":     resourceTencentCloudTsfApplicationFileConfigRelease(),
			"cloud_tsf_application_public_config":           resourceTencentCloudTsfApplicationPublicConfig(),
			"cloud_tsf_application_public_config_release":   resourceTencentCloudTsfApplicationPublicConfigRelease(),
			"cloud_tsf_config_template":                     resourceTencentCloudTsfConfigTemplate(),
			"cloud_tsf_lane":                                resourceTencentCloudTsfLane(),
			"cloud_tsf_lane_rule":                           resourceTencentCloudTsfLaneRule(),
			"cloud_tsf_api_group":                           resourceTencentCloudTsfApiGroup(),
			"cloud_tsf_bind_api_group":                      resourceTencentCloudTsfBindApiGroup(),
			"cloud_tsf_microservice":                        resourceTencentCloudTsfMicroservice(),
			"cloud_tsf_api_rate_limit_rule":                 resourceTencentCloudTsfApiRateLimitRule(),
			"cloud_tsf_enable_unit_rule":                    resourceTencentCloudTsfEnableUnitRule(),
			"cloud_tsf_instances_attachment":                resourceTencentCloudTsfInstancesAttachment(),
			"cloud_tsf_path_rewrite":                        resourceTencentCloudTsfPathRewrite(),
			"cloud_tsf_task":                                resourceTencentCloudTsfTask(),
			"cloud_tsf_unit_rule":                           resourceTencentCloudTsfUnitRule(),
			"cloud_turbofs_sign_up_service":                 resourceTencentCloudTurbofsSignUpService(),
			"cloud_turbofs_file_system":                     resourceTencentCloudTurbofsFileSystem(),
			"cloud_turbofs_p_group":                         resourceTencentCloudTurbofsPGroup(),
			"cloud_turbofs_rule":                            resourceTencentCloudTurbofsRule(),
			"cloud_turbofs_auto_snapshot_policy":            resourceTencentCloudTurbofsAutoSnapshotPolicy(),
			"cloud_turbofs_auto_snapshot_policy_attachment": resourceTencentCloudTurbofsAutoSnapshotPolicyAttachment(),
			"cloud_vpc":                                     resourceTencentCloudVpcInstance(),
			"cloud_vpc_acl":                                 resourceTencentCloudVpcACL(),
			"cloud_vpc_acl_attachment":                      resourceTencentCloudVpcAclAttachment(),
			"cloud_vpc_address_template":                    resourceTencentCloudAddressTemplate(),
			"cloud_vpc_address_template_group":              resourceTencentCloudAddressTemplateGroup(),
			"cloud_vpc_classic_link_attachment":             resourceTencentCloudVpcClassicLinkAttachment(),
			"cloud_vpc_dc_gateway":                          resourceTencentCloudDcGatewayInstance(),
			"cloud_vpc_peer_connect_manager":                resourceTencentCloudVpcPeerConnectManager(),
			"cloud_vpc_peer_connect_ex_manager":             resourceTencentCloudVpcPeerConnectExManager(),
			"cloud_vpc_dnat":                                resourceTencentCloudDnat(),
			"cloud_vpc_enable_end_point_connect":            resourceTencentCloudVpcEnableEndPointConnect(),
			"cloud_vpc_end_point":                           resourceTencentCloudVpcEndPoint(),
			"cloud_vpc_end_point_service":                   resourceTencentCloudVpcEndPointService(),
			"cloud_vpc_eni":                                 resourceTencentCloudEni(),
			"cloud_vpc_eni_attachment":                      resourceTencentCloudEniAttachment(),
			"cloud_vpc_eni_sg_attachment":                   resourceTencentCloudEniSgAttachment(),
			"cloud_vpc_ha_vip":                              resourceTencentCloudHaVip(),
			"cloud_vpc_ha_vip_eip_attachment":               resourceTencentCloudHaVipEipAttachment(),
			"cloud_vpc_ipv6_address_bandwidth":              resourceTencentCloudIpv6AddressBandwidth(),
			"cloud_vpc_ipv6_cidr_block":                     resourceTencentCloudVpcIpv6CidrBlock(),
			"cloud_vpc_ipv6_eni_address":                    resourceTencentCloudVpcIpv6EniAddress(),
			"cloud_vpc_ipv6_subnet_cidr_block":              resourceTencentCloudVpcIpv6SubnetCidrBlock(),
			"cloud_vpc_nat_gateway":                         resourceTencentCloudNatGateway(),
			"cloud_vpc_net_detect":                          resourceTencentCloudVpcNetDetect(),
			"cloud_vpc_route_table":                         resourceTencentCloudVpcRouteTable(),
			"cloud_vpc_route_table_entry":                   resourceTencentCloudVpcRouteEntry(),
			"cloud_vpc_security_group":                      resourceTencentCloudSecurityGroup(),
			"cloud_vpc_security_group_lite_rule":            resourceTencentCloudSecurityGroupLiteRule(),
			"cloud_vpc_security_group_rule":                 resourceTencentCloudSecurityGroupRule(),
			"cloud_vpc_security_group_rule_set":             resourceTencentCloudSecurityGroupRuleSet(),
			"cloud_vpc_subnet":                              resourceTencentCloudVpcSubnet(),
			"cloud_vpcdns_domain":                           resourceTencentCloudVpcDnsDomain(),
			"cloud_vpcdns_forward_rule":                     resourceTencentCloudVpcDnsForwardRule(),
			"cloud_vpcdns_record":                           resourceTencentCloudVpcDnsRecord(),
			"cloud_cwp_license_order":                       ResourceTencentCloudCwpLicenseOrder(),
			"cloud_cwp_license_bind_attachment":             ResourceTencentCloudCwpLicenseBindAttachment(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	secretId := d.Get("secret_id").(string)
	secretKey := d.Get("secret_key").(string)
	securityToken := d.Get("security_token").(string)
	region := d.Get("region").(string)
	protocol := d.Get("protocol").(string)
	domain := d.Get("domain").(string)
	cspDomain := d.Get("csp_domain").(string)
	cosDomain := d.Get("cos_domain").(string)
	// standard client
	var tcClient TencentCloudClient
	tcClient.apiV3Conn = &connectivity.TencentCloudClient{
		Credential: common.NewTokenCredential(
			secretId,
			secretKey,
			securityToken,
		),
		CredentialTce: common2.NewTokenCredential(
			secretId,
			secretKey,
			securityToken,
		),
		Region:    region,
		Protocol:  protocol,
		Domain:    domain,
		CspDomain: cspDomain,
		CosDomain: cosDomain,
	}

	envRoleArn := os.Getenv(PROVIDER_ASSUME_ROLE_ARN)
	envSessionName := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_NAME)

	// get assume role from env
	if envRoleArn != "" && envSessionName != "" {
		var assumeRoleSessionDuration int
		if envSessionDuration := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); envSessionDuration != "" {
			var err error
			assumeRoleSessionDuration, err = strconv.Atoi(envSessionDuration)
			if err != nil {
				return nil, diag.FromErr(err)
			}
		}
		if assumeRoleSessionDuration == 0 {
			assumeRoleSessionDuration = 7200
		}

		_ = genClientWithSTS(&tcClient, envRoleArn, envSessionName, assumeRoleSessionDuration, "")
	}

	// get assume role from tf config
	assumeRoleList := d.Get("assume_role").(*schema.Set).List()
	if len(assumeRoleList) == 1 {
		assumeRole := assumeRoleList[0].(map[string]interface{})
		assumeRoleArn := assumeRole["role_arn"].(string)
		assumeRoleSessionName := assumeRole["session_name"].(string)
		assumeRoleSessionDuration := assumeRole["session_duration"].(int)
		assumeRolePolicy := assumeRole["policy"].(string)

		_ = genClientWithSTS(&tcClient, assumeRoleArn, assumeRoleSessionName, assumeRoleSessionDuration, assumeRolePolicy)
	}
	return &tcClient, nil
}

func genClientWithSTS(tcClient *TencentCloudClient, assumeRoleArn, assumeRoleSessionName string,
	assumeRoleSessionDuration int, assumeRolePolicy string) error {
	// applying STS credentials
	request := sts.NewAssumeRoleRequest()
	request.RoleArn = helper.String(assumeRoleArn)
	request.RoleSessionName = helper.String(assumeRoleSessionName)
	request.DurationSeconds = helper.IntUint64(assumeRoleSessionDuration)
	if assumeRolePolicy != "" {
		request.Policy = helper.String(url.QueryEscape(assumeRolePolicy))
	}
	ratelimit.Check(request.GetAction())
	response, err := tcClient.apiV3Conn.UseStsClient().AssumeRole(request)
	if err != nil {
		return err
	}
	// using STS credentials
	tcClient.apiV3Conn.Credential = common.NewTokenCredential(
		*response.Response.Credentials.TmpSecretId,
		*response.Response.Credentials.TmpSecretKey,
		*response.Response.Credentials.Token,
	)
	return nil
}
