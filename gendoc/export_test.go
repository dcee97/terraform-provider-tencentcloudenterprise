package main

import (
	"fmt"
	"testing"
)

func TestExportSchemaAll(t *testing.T) {
	err := ExportJSON("resource")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestExportDataSchema(t *testing.T) {
	err := ExportJSON("data")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestGetSchemaKeys(t *testing.T) {
	res := getSchemaKeys("data")
	fmt.Println(res)
}

func TestExportDataMenu(t *testing.T) {
	err := ExportMenu("data")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestExportResourceMenu(t *testing.T) {
	err := ExportMenu("resource")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestExportAll(t *testing.T) {
	_ = ExportMenu("resource")
	_ = ExportMenu("data")
	_ = ExportJSON("resource")
	_ = ExportJSON("data")
}
