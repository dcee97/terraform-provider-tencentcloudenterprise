package tencentcloud

type CNDescription struct {
	TerraformTypeCN string `json:"TerraformTypeCN"`
	AttributesCN    map[string]string
	DescriptionCN   string `json:"DescriptionCN"`
}

// ResourceDescriptionProviders resource中文描述提供全局变量
var ResourceDescriptionProviders = map[string]CNDescription{}

// DataDescriptionProviders data中文描述提供全局变量
var DataDescriptionProviders = map[string]CNDescription{}

func registerResourceDescriptionProvider(resourceName string, description CNDescription) {
	ResourceDescriptionProviders[resourceName] = description
}

func registerDataDescriptionProvider(resourceName string, description CNDescription) {
	DataDescriptionProviders[resourceName] = description
}
