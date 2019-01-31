// Copyright (c) Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package lambdroid

import (
    "github.com/stack360/go-lambdroid-tower"
)

// LambdroidTowerClient acts like a client that triggers actions
type LambdroidTowerClient interface {
    ListServices() ([]string, error)
    GetServiceByName(serviceName string) (map[string]interface{}, error)
    AddService(serviceSpec map[string]interface{}) (map[string]interface{}, error)
    DeleteService(serviceName string) error
    UpdateService(serviceName string, updatedServiceSpec map[string]interface{}) (map[string]interface{}, error)
    InvokeService(serviceName string, serviceParams []byte) (string, error)
}

// Client is the REST Client type
type Client struct {
    towerClient *client.Client
    config      *Config
}

func NewTowerClientFromConfig(config *Config) (LambdroidTowerClient, error) {
    c := client.NewClient(config.LambdroidTowerURL)
    client := Client {
        towerClient:   c,
        config:        config,
    }
    return &client, nil
}

// TODO: Most of these functions are placeholders for now. Implement these after Lambdroid master app has the functionalities.
func (c *Client) ListServices() ([]string, error) {
    services := []string{"labelmaker2"}
    return services, nil
}

func (c *Client) GetServiceByName(serviceName string) (map[string]interface{}, error) {
    service := map[string]interface{} {}
    return service, nil
}

func (c *Client) AddService(serviceSpec map[string]interface{}) (map[string]interface{}, error) {
    service := map[string]interface{} {}
    return service, nil
}

func (c *Client) DeleteService(serviceName string) (error) {
    return nil
}

func (c *Client) UpdateService(serviceName string, updatedServiceSpec map[string]interface{}) (map[string]interface{}, error) {
    service := map[string]interface{} {}
    return service, nil
}

func (c *Client) InvokeService(serviceName string, serviceParams []byte) (string, error) {
    status, err := c.towerClient.RunFunction(serviceName, serviceParams)
    if err != nil {
        return "Error when invoking service", err
    }
    return status, nil
}
