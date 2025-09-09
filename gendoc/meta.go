package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cloud "terraform-provider-tencentcloudenterprise/tencentcloud"
)

const (
	TypeInvalid = "invalid"
	TypeBool    = "bool"
	TypeInt     = "int"
	TypeFloat   = "float"
	TypeString  = "string"
	TypeList    = "list"
	TypeMap     = "map"
	TypeSet     = "set"
)

// ExportField 导出参数结构体
type ExportField struct {
	Name           string        `json:"Name"`
	Computed       bool          `json:"Computed"`
	Description    string        `json:"Description"`
	Default        any           `json:"Default,omitempty"`
	ForceNew       bool          `json:"ForceNew"`
	Optional       bool          `json:"Optional"`
	Required       bool          `json:"Required"`
	Sensitive      bool          `json:"Sensitive"`
	Type           string        `json:"Type"`
	PossibleValues []ExportField `json:"PossibleValues,omitempty"`
	DescriptionCN  string        `json:"DescriptionCN"`
	// 数组子类型
	SubType string `json:"SubType,omitempty"`
}

type SchemaJSON struct {
	TerraformType    string        `json:"TerraformType"`
	Attributes       []ExportField `json:"Attributes"`
	BlockTypes       []ExportField `json:"BlockTypes"`
	ExportAttributes []ExportField `json:"ExportAttributes"`
	ExportBlockTypes []ExportField `json:"ExportBlockTypes"`
	Description      string        `json:"Description"`
	TerraformTypeCN  string        `json:"TerraformTypeCN"`
	Hcl              string        `json:"Hcl"`
}

type ResourcesJSON struct {
	Items []SchemaJSON `json:"Items"`
}

type Item struct {
	TerraformType string `json:"TerraformType"`
	Title         string `json:"Title"`
	IconPath      string `json:"IconPath"`
	Description   string `json:"Description"`
	DescriptionCN string `json:"DescriptionCN"`
}

type Area struct {
	Name  string `json:"Name"`
	Key   string `json:"Key"`
	Items []Item `json:"Items"`
}

type Category struct {
	Title         string `json:"Title"`
	Description   string `json:"Description"`
	IconPath      string `json:"IconPath"`
	Items         []Item `json:"Items"`
	DescriptionCN string `json:"DescriptionCN"`
}

type Menu struct {
	Key        string     `json:"Key"`
	Name       string     `json:"Name"`
	Categories []Category `json:"Categories"`
}

type TencentCloud struct {
	Areas    Area   `json:"Areas"`
	MenuList []Menu `json:"MenuList"`
}

func getMenu() TencentCloud {
	return TencentCloud{
		Areas: Area{
			Name: "Areas",
			Key:  "Areas",
			Items: []Item{
				{
					TerraformType: "region",
					Title:         "region",
					IconPath:      "/resource/area/cloud_region.svg",
					Description:   "地域",
				},
				{
					TerraformType: "az",
					Title:         "available zone",
					IconPath:      "/resource/area/cloud_az.svg",
					Description:   "可用区",
				},
			},
		},
		MenuList: []Menu{
			{
				Key:  "Compute",
				Name: "计算",
				Categories: []Category{
					{
						Title:         "CVM",
						Description:   "Cloud Virtual Machine",
						DescriptionCN: "云服务器",
					},
					{
						Title:         "BMS",
						Description:   "Cloud Bare Metal Server",
						DescriptionCN: "裸金属服务器",
					},
					{
						Title:         "AS",
						Description:   "Cloud Auto Scaling",
						DescriptionCN: "弹性伸缩",
					},
					{
						Title:         "TKE",
						Description:   "Cloud Kubernetes Engine",
						DescriptionCN: "容器服务(TKE)",
					},
					{
						Title:         "TCR",
						Description:   "Tencent Container Registry",
						DescriptionCN: "容器镜像服务",
					},
				},
			},
			{
				Key:  "Storage",
				Name: "存储",
				Categories: []Category{
					{
						Title:         "CBS",
						Description:   "Cloud Block Storage",
						DescriptionCN: "云硬盘",
					},
					{
						Title:         "CFS",
						Description:   "Cloud File Storage",
						DescriptionCN: "文件存储",
					},
					{
						Title:         "CSP",
						Description:   "Cloud Storage Private",
						DescriptionCN: "对象存储CSP",
					},
					{
						Title:         "COS",
						Description:   "Cloud Object Storage",
						DescriptionCN: "对象存储COS",
					},
					{
						Title:         "TurboFS",
						Description:   "Parallel File Storage",
						DescriptionCN: "并行文件存储",
					},
					{
						Title:         "BRC",
						Description:   "Backup and Recovery Center",
						DescriptionCN: "备份恢复中心",
					},
					//{
					//	Title:         "DRC",
					//	Description:   "Disaster Recovery Center",
					//	DescriptionCN: "容灾中心",
					//},
				},
			},
			{
				Key:  "Network",
				Name: "网络",
				Categories: []Category{
					{
						Title:         "VPC",
						Description:   "Virtual Private Cloud",
						DescriptionCN: "私有网络",
					},
					{
						Title:         "CLB",
						Description:   "Cloud Load Balancer",
						DescriptionCN: "负载均衡",
					},
					{
						Title:         "VPCDNS",
						Description:   "Virtual Private Cloud Domain Name System",
						DescriptionCN: "DNS域名解析",
					},
					{
						Title:         "EIP",
						Description:   "Cloud Elastic IP",
						DescriptionCN: "弹性公网IP",
					},
				},
			},
			{
				Key:  "Database",
				Name: "数据库",
				Categories: []Category{
					{
						Title:         "Redis",
						Description:   "Cloud Redis®",
						DescriptionCN: "云数据库Redis®",
					},
					{
						Title:         "DCDB",
						Description:   "Cloud Distributed SQL For MySQL",
						DescriptionCN: "云数据库TDSQL MySQL",
					},
					{
						Title:         "TBASE",
						Description:   "Cloud Distributed SQL For PostGreSQL",
						DescriptionCN: "云数据库TDSQL PostgreSQL",
					},
				},
			},
			{
				Key:  "Middleware",
				Name: "中间件",
				Categories: []Category{
					{
						Title:         "TDMQ",
						Description:   "Cloud Distributed Message Queue",
						DescriptionCN: "消息队列TDMQ",
					},
					{
						Title:         "TSF",
						Description:   "Cloud Serverless Framework",
						DescriptionCN: "微服务平台TSF",
					},
					{
						Title:         "CKafka",
						Description:   "Cloud Kafka",
						DescriptionCN: "消息队列CKafka",
					},
				},
			},
			{
				Key:  "BigData",
				Name: "大数据",
				Categories: []Category{
					{
						Title:         "CLS",
						Description:   "Cloud Log Service",
						DescriptionCN: "日志服务",
					},
				},
			},
			{
				Key:  "Security",
				Name: "安全",
				Categories: []Category{
					{
						Title:         "CWP",
						Description:   "Cloud Workload Protection",
						DescriptionCN: "主机安全",
					},
					{
						Title:         "CFW",
						Description:   "Cloud Firewall",
						DescriptionCN: "云防火墙",
					},
					{
						Title:         "SSM",
						Description:   "Secrets Manager",
						DescriptionCN: "凭据管理",
					},
					{
						Title:         "KMS",
						Description:   "Key Management Service",
						DescriptionCN: "密钥管理服务",
					},
				},
			},
			{
				Key:  "Identity and Access Management",
				Name: "身份与访问管理",
				Categories: []Category{
					{
						Title:         "Organization",
						Description:   "Corporate Account Management",
						DescriptionCN: "集团账号管理",
					},
					{
						Title:         "CIC",
						Description:   "Corporate Identity Center",
						DescriptionCN: "集团账号身份中心",
					},
				},
			},
		},
	}
}

func valueType(valueType schema.ValueType) string {
	providerTypeMap := map[schema.ValueType]string{
		schema.TypeInvalid: TypeInvalid,
		schema.TypeBool:    TypeBool,
		schema.TypeInt:     TypeInt,
		schema.TypeFloat:   TypeFloat,
		schema.TypeString:  TypeString,
		schema.TypeList:    TypeList,
		schema.TypeMap:     TypeMap,
		schema.TypeSet:     TypeSet,
	}

	if _, ok := providerTypeMap[valueType]; ok {
		return providerTypeMap[valueType]
	}
	return "unknown"
}

func dfs(key string, value *schema.Schema, description map[string]string) ExportField {
	field := ExportField{
		Name:        key,
		Computed:    value.Computed,
		Description: value.Description,
		Default:     value.Default,
		ForceNew:    value.ForceNew,
		Optional:    value.Optional,
		Required:    value.Required,
		Sensitive:   value.Sensitive,
		Type:        valueType(value.Type),
	}
	if description != nil {
		field.DescriptionCN = description[key]
	}
	//if elemResource, ok := value.Elem.(*schema.Resource); ok {
	//	for elemKey, elemValue := range elemResource.Schema {
	//		elemField := dfs(elemKey, elemValue)
	//		field.PossibleValues = append(field.PossibleValues, elemField)
	//	}
	//}

	if value.Elem != nil {
		switch value.Elem.(type) {
		case *schema.Resource:
			for elemKey, elemValue := range value.Elem.(*schema.Resource).Schema {
				elemField := dfs(elemKey, elemValue, description)
				field.PossibleValues = append(field.PossibleValues, elemField)
			}
		default:
			// 数组子类型
			field.SubType = valueType(value.Elem.(*schema.Schema).Type)
			// field.PossibleValues = append(field.PossibleValues, dfs("SubType", value.Elem.(*schema.Schema), description))
		}
	}

	return field
}

func resolveSchema(dtype, fpath, resourceName string, schema *schema.Resource,
	descriptions cloud.CNDescription) SchemaJSON {
	schemaJSON := SchemaJSON{
		TerraformType:    resourceName,
		BlockTypes:       []ExportField{},
		ExportBlockTypes: []ExportField{},
		Description:      schema.Description,
		TerraformTypeCN:  descriptions.TerraformTypeCN,
	}
	hcl, err := getHcl(dtype, fpath, resourceName)
	if err != nil {
		message("[FAIL!]get hcl failed: %s", err)
		os.Exit(1)
	}
	schemaJSON.Hcl = hcl

	schemaMap := schema.Schema
	for key, value := range schemaMap {
		field := dfs(key, value, descriptions.AttributesCN)
		// 参数规则
		if field.Required || field.Optional {
			if len(field.PossibleValues) != 0 {
				schemaJSON.BlockTypes = append(schemaJSON.BlockTypes, field)
			} else {
				schemaJSON.Attributes = append(schemaJSON.Attributes, field)
			}
		}
		// computed 为属性
		if field.Computed {
			if len(field.PossibleValues) != 0 {
				schemaJSON.ExportBlockTypes = append(schemaJSON.ExportBlockTypes, field)
			} else {
				schemaJSON.ExportAttributes = append(schemaJSON.ExportAttributes, field)
			}
		}
	}
	return schemaJSON
}

// ExportJSON 导出所有资源或数据源数据
func ExportJSON(itemType string) error {
	provider := cloud.Provider()
	vProvider := runtime.FuncForPC(reflect.ValueOf(cloud.Provider).Pointer())

	filename, _ := vProvider.FileLine(0)
	fpath := filepath.Dir(filename)
	var (
		err                  error
		schemaJSON           any
		fileName             = itemType
		itemMap              map[string]*schema.Resource
		descriptionsProvider map[string]cloud.CNDescription
		dtype                string
	)

	// 根据 itemType 选择资源或数据源
	if itemType == "resource" {
		itemMap = provider.ResourcesMap
		descriptionsProvider = cloud.ResourceDescriptionProviders
		dtype = "resource"
	} else if itemType == "data" {
		itemMap = provider.DataSourcesMap
		descriptionsProvider = cloud.DataDescriptionProviders
		dtype = "data_source"
	} else {
		return fmt.Errorf("invalid item type: %s", itemType)
	}

	fileName += "_all.json"
	temp := ResourcesJSON{}
	for k, v := range itemMap {
		s := resolveSchema(dtype, fpath, k, v, descriptionsProvider[k])
		temp.Items = append(temp.Items, s)
	}
	//var wg sync.WaitGroup
	//var mutex sync.Mutex
	//
	//index := 0
	//for k, v := range itemMap {
	//	wg.Add(1)
	//
	//	go func(k string, v *schema.Resource, i int) {
	//		defer wg.Done()
	//
	//		// 在协程中调用 resolveSchema
	//		s := resolveSchema(dtype, fpath, k, v, descriptionsProvider[k])
	//
	//		// 使用 mutex 锁来同步对 temp.Items 的操作
	//		mutex.Lock()
	//		temp.Items = append(temp.Items, s)
	//		mutex.Unlock()
	//	}(k, v, index)
	//
	//	index++
	//}
	//
	//wg.Wait()

	sort.Slice(temp.Items, func(i, j int) bool {
		return temp.Items[i].TerraformType < temp.Items[j].TerraformType
	})
	schemaJSON = temp

	jsonBytes, err := json.MarshalIndent(schemaJSON, "", "  ")
	if err != nil {
		return err
	}
	// write to file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	defer file.Close()
	file.Write(jsonBytes)
	return nil
}

// ExportMenu 导出菜单数据
func ExportMenu(name string) error {
	provider := cloud.Provider()
	if name != "resource" && name != "data" {
		return errors.New("invalid name")
	}
	var descriptionMap map[string]cloud.CNDescription
	var data any
	menu := getMenu()
	if name == "resource" {
		data = provider.ResourcesMap
		descriptionMap = cloud.ResourceDescriptionProviders
	} else {
		data = provider.DataSourcesMap
		descriptionMap = cloud.DataDescriptionProviders
	}
	mp := getSchemaKeys(name)

	for i, v := range menu.MenuList {
		for j, c := range v.Categories {
			// 当前产品
			menu.MenuList[i].Categories[j].IconPath = fmt.Sprintf("/%s/%s/%s.svg",
				name, strings.ToLower(c.Title), strings.ToLower(c.Title))
			key := strings.ToLower(c.Title)

			// 当前产品下的所有资源
			keys := mp[key]
			if len(keys) == 0 {
				continue
			}

			for _, r := range keys {
				resource := data.(map[string]*schema.Resource)[r]
				descriptions := descriptionMap[r]
				descriptionCN := descriptions.DescriptionCN
				//if descriptionCN == "" {
				//	// 如果没有中文描述，使用英文描述或生成默认中文描述
				//	descriptionCN = generateDefaultDescriptionCN(resource.Description, descriptions.TerraformTypeCN)
				//}
				item := Item{
					TerraformType: r,
					Title:         descriptions.TerraformTypeCN,
					IconPath:      fmt.Sprintf("/%s/%s/%s.svg", name, strings.ToLower(c.Title), r),
					Description:   resource.Description,
					DescriptionCN: descriptionCN,
				}
				menu.MenuList[i].Categories[j].Items = append(menu.MenuList[i].Categories[j].Items, item)
			}
			// 产品内资源排序
			sort.Slice(menu.MenuList[i].Categories[j].Items, func(a, b int) bool {
				return menu.MenuList[i].Categories[j].Items[a].TerraformType < menu.MenuList[i].Categories[j].
					Items[b].TerraformType
			})
			// redis修改商标
			if strings.ToLower(c.Title) == "redis" {
				menu.MenuList[i].Categories[j].Title = "Redis®"
			}
		}
	}

	fileName := fmt.Sprintf("%s_menu.json", name)
	jsonBytes, err := json.MarshalIndent(menu, "", "  ")
	if err != nil {
		return err
	}
	// write to file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	defer file.Close()
	file.Write(jsonBytes)
	return nil
}

/*
getSchemaKeys 获取所有产品下的资源
返回格式：
map[cloud_as:[cloud_as_scaling_policies
cloud_as_instances cloud_as_scaling_groups
cloud_as_scaling_configs cloud_as_last_activity]
*/
func getSchemaKeys(name string) map[string][]string {
	provider := cloud.Provider()
	var data any
	menu := getMenu()
	if name == "resource" {
		data = provider.ResourcesMap

	} else if name == "data" {
		data = provider.DataSourcesMap
	} else {
		return nil
	}
	products := map[string][]string{}
	for _, v := range menu.MenuList {
		for _, c := range v.Categories {
			key := strings.ToLower(c.Title)
			products[key] = []string{}
		}
	}

	for k := range data.(map[string]*schema.Resource) {
		for key := range products {
			ss := strings.Split(k, "_")
			if ss[1] == key {
				products[key] = append(products[key], k)
			}
		}
	}
	return products
}

func getHcl(dtype, fpath, name string) (string, error) {
	name = strings.TrimPrefix(name, cloudPrefix)

	filename := fmt.Sprintf("%s_%s_%s.go", dtype, cloudMarkShort, name)
	message("[START]get description from file: %s\n", filename)

	description, err := getFileDescription(filepath.Join(fpath, filename))
	if err != nil {
		message("[FAIL!]get description failed: %s", err)
		return "", err
	}

	description = strings.TrimSpace(description)
	if description == "" {
		message("[FAIL!]description empty: %s\n", filename)
		return "", err
	}

	importPos := strings.Index(description, "\nImport\n")
	if importPos != -1 {
		description = strings.TrimSpace(description[:importPos])
	}

	pos := strings.Index(description, "\nExample Usage\n")
	hcl := ""
	if pos != -1 {
		hcl = formatHCL(description[pos+15:])
		description = strings.TrimSpace(description[:pos])
	} else {
		message("[FAIL!]example usage missing: %s\n", filename)
		os.Exit(1)
	}
	return hcl, nil
}

// generateDefaultDescriptionCN 为没有中文描述的资源生成默认中文描述
//func generateDefaultDescriptionCN(englishDescription, resourceTypeCN string) string {
//	if englishDescription == "" && resourceTypeCN == "" {
//		return "提供资源管理功能。"
//	}
//
//	if resourceTypeCN != "" {
//		return fmt.Sprintf("提供%s资源，用于创建和管理相关资源。", resourceTypeCN)
//	}
//
//	// 如果只有英文描述，尝试简单的翻译映射
//	englishLower := strings.ToLower(englishDescription)
//	if strings.Contains(englishLower, "provide") && strings.Contains(englishLower, "resource") {
//		return "提供资源管理功能。"
//	}
//
//	return "提供资源管理功能。"
//}
