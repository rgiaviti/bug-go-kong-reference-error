package main

import (
	"context"
	"fmt"
	"github.com/kong/go-kong/kong"
)

func main() {
	fmt.Println("Reproducing bug...")
	CallingCorrectKongPort()
	CallingIncorrectKongPort()
}

func CallingCorrectKongPort() {
	fmt.Println("Calling AdminAPI in Port 8001")

	correctKongHost := "http://localhost:8001"
	kongAdmClient := CreateKongAdminClient(&correctKongHost)

	ctx := context.Background()
	plugin := GetSamplePlugin()
	result, msg, err := kongAdmClient.Plugins.Validate(ctx, plugin)
	if err != nil {
		panic("error validating plugin")
	}

	fmt.Println(fmt.Sprintf("validation of %s", *plugin.Name))
	fmt.Println(fmt.Sprintf("result: %v", result))
	fmt.Println(fmt.Sprintf("message: %s", msg))
}

func CallingIncorrectKongPort() {
	fmt.Println("Calling AdminAPI in Port 9000")

	incorrectKongHost := "http://localhost:9001"
	kongAdmClient := CreateKongAdminClient(&incorrectKongHost)

	ctx := context.Background()
	plugin := GetSamplePlugin()
	result, msg, err := kongAdmClient.Plugins.Validate(ctx, plugin)
	if err != nil {
		panic("error validating plugin")
	}

	// code not reached
	fmt.Println(fmt.Sprintf("validation of %s", *plugin.Name))
	fmt.Println(fmt.Sprintf("result: %v", result))
	fmt.Println(fmt.Sprintf("message: %s", msg))
}

func CreateKongAdminClient(host *string) *kong.Client {
	kongAdmClient, err := kong.NewClient(host, nil)
	if err != nil {
		panic("error creating kong admin client")
	}

	return kongAdmClient
}

func GetSamplePlugin() *kong.Plugin {
	name := "correlation-id"
	enabled := true
	configuration :=  make(map[string]interface{})

	return &kong.Plugin{
		Name:     &name,
		Config:    configuration,
		Enabled:   &enabled,
	}
}