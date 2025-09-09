package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const successDoc = `
Provider Data Sources
  cloud_availability_zones

Auto Scaling(AS)
  Data Source
    cloud_as_scaling_configs
    cloud_as_scaling_groups


  Resource
    cloud_as_scaling_config
    cloud_as_scaling_group


Cloud Access Management(CAM)
  Data Source
    cloud_cam_group_memberships
    cloud_cam_group_policy_attachments


  Resource
    cloud_cam_role
    cloud_cam_role_policy_attachment

`

func TestGetIndex(t *testing.T) {
	type args struct {
		doc string
	}
	tests := []struct {
		name    string
		args    args
		want    []Product
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				doc: successDoc,
			},
			want: []Product{
				{
					Name: "Provider Data Sources",
					DataSources: []string{
						"cloud_availability_zones",
					},
				},
				{
					Name: "Auto Scaling(AS)",
					DataSources: []string{
						"cloud_as_scaling_configs",
						"cloud_as_scaling_groups",
					},
					Resources: []string{
						"cloud_as_scaling_config",
						"cloud_as_scaling_group",
					},
				},
				{
					Name: "Cloud Access Management(CAM)",
					DataSources: []string{
						"cloud_cam_group_memberships",
						"cloud_cam_group_policy_attachments",
					},
					Resources: []string{
						"cloud_cam_role",
						"cloud_cam_role_policy_attachment",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIndex(tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndexFromURL(t *testing.T) {
	dtype := "data_resource"
	fmt.Println(strings.Replace(dtype, "_", "", -1))
}
