/*
Provides a resource to create a turbo file system(TurboFS).

# Example Usage

```hcl

	resource "cloud_turbofs_file_system" "foo" {
	  fs_name           = "test_file_system"
	  zone 				= "ap-beijing-region-jcctest-ops-2"
	  pgroup_id         = "pgroupbasic"
	  capacity          = 2.5
	  protocol          = "NFS"
	  vpc_id            = "vpc-c1tlh29v"
	  subnet_id         = "subnet-ge8g8x0a"
	  cfs_version       = "4.0"
	  net_interface     = "vpc"
	  storage_type      = "TB"
	}

```

# Import

Cloud file system can be imported using the id, e.g.

```
$ terraform import cloud_turbofs_file_system.foo turbofs-6hgquxmj
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	registerResourceDescriptionProvider("cloud_turbofs_file_system", CNDescription{
		TerraformTypeCN: "文件系统",
		DescriptionCN:   "提供TurboFS文件系统资源，用于创建和管理TurboFS文件系统。",
		AttributesCN: map[string]string{
			"zone_id":                   "可用区 ID",
			"zone":                      "可用区名称，例如:ap-zone-1",
			"creation_token":            "用户自定义文件系统名称，优先级低于fs_name",
			"protocol":                  "文件系统协议类型， 值为TURBO",
			"fs_name":                   "用户自定义文件系统名称,与CreationToken 两者必须填一项",
			"availability_zone":         "可用区",
			"pgroup_id":                 "权限组ID",
			"cfs_version":               "CFS文件系统版本",
			"capacity":                  "TURBO文件系统的容量（TiB）",
			"pool_id":                   "资源池ID",
			"snapshot_id":               "快照ID",
			"storage_type":              "文件系统存储类型，值为 TB/TP ；其中 TB为标准型存储, TP为性能型存储",
			"net_interface":             "网络类型，值为 VPC，BASIC；其中 VPC 为私有网络，BASIC 为基础网络",
			"vpc_id":                    "私有网路（VPC） ID;当网络类型值为 VPC时，与UnVpcId 两者必须填一项",
			"un_vpc_id":                 "系统分配的VPC统一ID",
			"subnet_id":                 "子网， 当网络类型值为 VPC时，与UnSubnetId 两者必须填一项",
			"un_subnet_id":              "系统分配的子网统一 ID",
			"mount_ip":                  "指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP",
			"storage_resource_pkg_id":   "文件系统绑定的存储包，每个文件系统只能绑定一个",
			"bandwidth_resource_pkg_id": "文件系统绑定的带宽包，每个文件系统只能绑定一个",
			"encrypted":                 "文件系统是否加密，若留空则默认为不加密",
			"kms_key_id":                "加密密钥 ID",
			"resource_tags":             "资源标签",
			"project_id":                "项目ID",
			"creation_time":             "文件系统创建时间",
			"cfs_id":                    "文件系统ID",
			"tag_id":                    "TurboFS资源池id",
			"file_system_id":            "文件系统 ID",
			"life_cycle_state":          "文件系统状态",
			"size_byte":                 "文件系统已使用容量大小",
			"tag_key":                   "标签键",
			"tag_value":                 "标签值",
		},
	})
}
func resourceTencentCloudTurbofsFileSystem() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a parallel file storage(TurboFS).",
		Create:      resourceTencentCloudTurbofsFileSystemCreate,
		Read:        resourceTencentCloudTurbofsFileSystemRead,
		Update:      resourceTencentCloudTurbofsFileSystemUpdate,
		Delete:      resourceTencentCloudTurbofsFileSystemDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fs_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of a file system.",
			},
			"creation_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User-defined file system name, with lower priority than fs_name.",
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The available zone that the file system locates at.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The available zone id that the file system locates at.",
			},
			"pgroup_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of a permission group.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "File service protocol. Valid value, the default is `TURBO`.",
			},
			"cfs_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "TurboFS file system version.",
			},
			"capacity": {
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Capacity of the TURBO file system in TiB. The value should be an integer multiple of the spec policy",
			},
			"pool_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the resource pool. If not specified, the system will automatically select the default resource pool.",
			},
			"snapshot_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Snapshot ID.",
			},
			"storage_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "File service StorageType. Valid values are `TP` and `TB`.",
			},
			"net_interface": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Network type. Valid value is `vpc`.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of a VPC network.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of a subnet.",
			},
			"un_vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Uniform id of a VPC network.",
			},
			"un_subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Uniform id of a subnet.",
			},
			"mount_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP of mount point.",
			},
			"storage_resource_pkg_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the storage resource package bound to the file system. Each file system can only be bound to one storage resource package.",
			},
			"bandwidth_resource_pkg_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the bandwidth resource package bound to the file system. Each file system can only be bound to one bandwidth resource package.",
			},
			"encrypted": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether the file system is encrypted.",
			},
			"kms_key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the KMS key used for encryption. This field is required when `encrypted` is set to true.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the project.",
			},
			"tag_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "TurboFS pool id.",
			},
			"resource_tag": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Tag key.",
						},
						"tag_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Tag value.",
						},
					},
				},
				Optional:    true,
				Description: "TurboFS resource tag.",
			},
			// computed
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time of the file system.",
			},
			"file_system_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id of the file system.",
			},
			"life_cycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Life cycle state of the file system.",
			},
			"size_byte": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Usage of the file system.",
			},
			"cfs_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Id of the file system",
			},
		},
	}
}

func resourceTencentCloudTurbofsFileSystemCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_file_system.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	request := turbofs.NewCreateCfsFileSystemRequest()

	if v, ok := d.GetOk("zone"); ok {
		request.Zone = helper.String(v.(string))
	}
	if v, ok := d.GetOk("pgroup_id"); ok {
		request.PGroupId = helper.String(v.(string))
	}
	if v, ok := d.GetOk("protocol"); ok {
		request.Protocol = helper.String(v.(string))
	}
	if v, ok := d.GetOk("cfs_version"); ok {
		request.CfsVersion = helper.String(v.(string))
	}
	if v, ok := d.GetOk("storage_type"); ok {
		request.StorageType = helper.String(v.(string))
	}
	if v, ok := d.GetOk("fs_name"); ok {
		request.FsName = helper.String(v.(string))
	}
	if v, ok := d.GetOk("net_interface"); ok {
		request.NetInterface = helper.String(v.(string))
	}

	if v, ok := d.GetOk("un_vpc_id"); ok {
		request.UnVpcId = helper.String(v.(string))
	} else if v, ok := d.GetOk("vpc_id"); ok {
		request.UnVpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("un_subnet_id"); ok {
		request.UnSubnetId = helper.String(v.(string))
	} else if v, ok := d.GetOk("subnet_id"); ok {
		request.UnSubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("capacity"); ok {
		capacity := v.(float64) * 1024
		request.Capacity = helper.IntUint64(int(capacity))
	}

	if v, ok := d.GetOkExists("zone_id"); ok {
		request.ZoneId = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("pool_id"); ok {
		request.PoolId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("snapshot_id"); ok {
		request.SnapshotId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("mount_ip"); ok {
		request.MountIP = helper.String(v.(string))
	}

	if v, ok := d.GetOk("storage_resource_pkg_id"); ok {
		request.StorageResourcePkgId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bandwidth_resource_pkg_id"); ok {
		request.BandwidthResourcePkgId = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("encrypted"); ok {
		request.Encrypted = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("kms_key_id"); ok {
		request.KmsKeyId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("project_id"); ok {
		request.ProjectId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("resource_tags"); ok {
		request.ResourceTags = make([]*turbofs.ResourceTags, 0, len(v.(map[string]interface{})))
		for tagKey, tagValue := range v.(map[string]interface{}) {
			tag := turbofs.ResourceTags{
				TagKey:   helper.String(tagKey),
				TagValue: helper.String(tagValue.(string)),
			}
			request.ResourceTags = append(request.ResourceTags, &tag)
		}
	}

	var fsId *string
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, err := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().CreateCfsFileSystem(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			return retryError(err)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response.Response.FileSystemId == nil {
			err = fmt.Errorf("file system id is nil")
			return resource.NonRetryableError(err)
		}
		fsId = response.Response.FileSystemId
		return nil
	})
	if err != nil {
		return err
	}
	d.SetId(*fsId)

	// wait for success status
	err = resource.Retry(2*readRetryTimeout, func() *resource.RetryError {
		fileSystems, errRet := turbofsService.DescribeFileSystem(ctx, fsId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if len(fileSystems) < 1 {
			return resource.RetryableError(fmt.Errorf("file system %s not exist", fsId))
		}
		if *fileSystems[0].LifeCycleState == CFS_FILE_SYSTEM_STATUS_CREATING {
			return resource.RetryableError(fmt.Errorf("file system status is %s, retry...", *fileSystems[0].LifeCycleState))
		}
		return nil
	})
	if err != nil {
		return err
	}

	return resourceTencentCloudTurbofsFileSystemRead(d, meta)
}

func resourceTencentCloudTurbofsFileSystemRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_file_system.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	fsId := helper.String(d.Id())
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var fileSystem *turbofs.FileSystemInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		fileSystems, errRet := turbofsService.DescribeFileSystem(ctx, fsId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if len(fileSystems) > 0 {
			fileSystem = fileSystems[0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if fileSystem == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("fs_name", fileSystem.FsName)
	_ = d.Set("zone", fileSystem.Zone)
	_ = d.Set("pgroup_id", fileSystem.PGroup.PGroupId)
	_ = d.Set("protocol", fileSystem.Protocol)
	_ = d.Set("snapshot_id", fileSystem.SnapshotId)
	_ = d.Set("create_time", fileSystem.CreationTime)
	_ = d.Set("storage_type", fileSystem.StorageType)

	return nil
}

func resourceTencentCloudTurbofsFileSystemUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_file_system.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	fsId := d.Id()
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	d.Partial(true)

	if d.HasChange("name") {
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			errRet := turbofsService.ModifyFileSystemName(ctx, fsId, d.Get("name").(string))
			if errRet != nil {
				return retryError(errRet)
			}
			return nil
		})
		if err != nil {
			return err
		}

	}

	if d.HasChange("access_group_id") {
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			errRet := turbofsService.ModifyFileSystemAccessGroup(ctx, fsId, d.Get("access_group_id").(string))
			if errRet != nil {
				return retryError(errRet)
			}
			return nil
		})
		if err != nil {
			return err
		}

	}

	d.Partial(false)

	return resourceTencentCloudCfsFileSystemRead(d, meta)
}

func resourceTencentCloudTurbofsFileSystemDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_file_system.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	fsId := d.Id()
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		errRet := turbofsService.DeleteFileSystem(ctx, fsId)
		if errRet != nil {
			return retryError(errRet)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
